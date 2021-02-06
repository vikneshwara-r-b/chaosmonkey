// Copyright 2016 Netflix, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package tracker provides an entry point for instantiating Trackers
package tracker

import (
	"github.com/Netflix/chaosmonkey"
	"github.com/Netflix/chaosmonkey/config"
	"github.com/Netflix/chaosmonkey/deps"
	"github.com/pkg/errors"
	"github.com/Netflix/chaosmonkey/helpers"
	"fmt"
	"strconv"
)

func init() {
	deps.GetTrackers = getTrackers
}

// getTrackers returns a list of trackers specified in the configuration
func getTrackers(cfg *config.Monkey) ([]chaosmonkey.Tracker, error) {
	var result []chaosmonkey.Tracker

	kinds, err := cfg.Trackers()
	if err != nil {
		return nil, err
	}

	for _, kind := range kinds {
		tr, err := getTracker(kind, cfg)
		if err != nil {
			return nil, err
		}
		result = append(result, tr)
	}
	return result, nil
}

// Posting instance termination message to Slack through webhook
func postToSlack(t chaosmonkey.Termination,cfg *config.Monkey) {
	fmt.Printf("Posting to slack")
	webhookUrl := cfg.GetWebHookUrl()
	fmt.Printf("Webhook URL:%s",webhookUrl)
	attachment1 := slack.Attachment {}
	termination_time := t.Time.String()
	fmt.Printf("Termination time:%s",termination_time)
	leashed := strconv.FormatBool(t.Leashed)
	fmt.Printf("Leashed status:%s",leashed)
	instance_data := t.Instance
	app_name := instance_data.AppName()
	fmt.Printf("Application name:%s",app_name)
	account_name := instance_data.AccountName()
	fmt.Printf("Account name:%s",account_name)
	region_name := instance_data.RegionName()
	fmt.Printf("Region name:%s",region_name)
	stack_name := instance_data.StackName()
	fmt.Printf("Stack name:%s",stack_name)
	cluster_name := instance_data.ClusterName()
	fmt.Printf("Cluster name:%s",cluster_name)
	asg_name := instance_data.ASGName()
	fmt.Printf("ASG name:%s",asg_name)
	instance_id :=  instance_data.ID()
	fmt.Printf("Instance ID:%s",instance_id)
	cloud_provider := instance_data.CloudProvider()
	fmt.Printf("Cloud Provider:%s",cloud_provider)
	message_format = `Termination time:%s
	Leashed status:%s
	----------- Instance details are given below: ------------
	Application name: %s 
	Account name: %s
	Region name: %s 
	Stack name: %s 
	Cluster name: %s  
	Auto Scaling Group name: %s  
	Instance-ID: %s
	Cloud Provider:%s`
    message_text := fmt.Sprintf(message_format,termination_time,leashed,app_name,account_name,region_name,stack_name,cluster_name,asg_name,instance_id,cloud_provider)
    payload := slack.Payload {
      Text: message_text,
      Attachments: []slack.Attachment{attachment1},
    }
    slack.Send(webhookUrl, "", payload)
}

// getTracker returns a tracker by name
// No trackers have been implemented yet
func getTracker(kind string, cfg *config.Monkey) (chaosmonkey.Tracker, error) {
	var termination chaosmonkey.Termination
	switch kind {
	// As trackers are contributed to the open source project, they should
	// be instantiated here
	case "notify_slack":
		fmt.Printf("Choosing notification through slack")
		postToSlack(termination,cfg)
	default:
		return nil, errors.Errorf("unsupported tracker: %s", kind)
	}
}
