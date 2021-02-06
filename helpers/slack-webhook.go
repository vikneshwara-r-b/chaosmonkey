package slack
import (
	"log"
  "net/http"
  "io/ioutil"
  "encoding/json"
   "bytes"
)
type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

type Action struct {
	Type	string   `json:"type"`
	Text	string   `json:"text"`
	Url 	string   `json:"url"`
	Style 	string   `json:"style"`
}

type Attachment struct {
	Fallback     *string   `json:"fallback"`
	Color        *string   `json:"color"`
	PreText      *string   `json:"pretext"`
	AuthorName   *string   `json:"author_name"`
	AuthorLink   *string   `json:"author_link"`
	AuthorIcon   *string   `json:"author_icon"`
	Title        *string   `json:"title"`
	TitleLink    *string   `json:"title_link"`
	Text         *string   `json:"text"`
	ImageUrl     *string   `json:"image_url"`
	Fields       []*Field  `json:"fields"`
	Footer       *string   `json:"footer"`
	FooterIcon   *string   `json:"footer_icon"`
	Timestamp    *int64    `json:"ts"`
	MarkdownIn   *[]string `json:"mrkdwn_in"`
	Actions      []*Action `json:"actions"`
	CallbackID   *string   `json:"callback_id"`
	ThumbnailUrl *string   `json:"thumb_url"`
}

type Payload struct {
	Parse       string       `json:"parse,omitempty"`
	Username    string       `json:"username,omitempty"`
	IconUrl     string       `json:"icon_url,omitempty"`
	IconEmoji   string       `json:"icon_emoji,omitempty"`
	Channel     string       `json:"channel,omitempty"`
	Text        string       `json:"text,omitempty"`
	LinkNames   string       `json:"link_names,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
	UnfurlLinks bool         `json:"unfurl_links,omitempty"`
	UnfurlMedia bool         `json:"unfurl_media,omitempty"`
	Markdown    bool         `json:"mrkdwn,omitempty"`
}

func (attachment *Attachment) AddField(field Field) *Attachment {
	attachment.Fields = append(attachment.Fields, &field)
	return attachment
}

func (attachment *Attachment) AddAction(action Action) *Attachment {
	attachment.Actions = append(attachment.Actions, &action)
	return attachment
}

func Send(webhookUrl string, proxy string, payload Payload) interface{} {
  json_payload,err := json.Marshal(payload)
  if json_payload != nil {
    resp, err := http.Post(webhookUrl, "application/json",  bytes.NewBuffer(json_payload))
    if err != nil {
      log.Fatalf("Error Occured after hittng Slack WebHook", err)
      return err
    } else {
      body, err := ioutil.ReadAll(resp.Body)
      if err != nil {
        log.Fatalln("Error in reading response body",err)
        return err
       }
       sb := string(body)
       log.Println("Response body is given below:")
       log.Printf(sb)
       return sb
    }
  } else {
    log.Fatalf("Error in parsing payload data to JSON",err)
    return err
  }
  return nil
}
