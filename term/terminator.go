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

package term

import (
	"fmt"

	"github.com/vikneshwara-r-b/chaosmonkey"
)

// fake is a fake implementation of a terminator that just prints termination info but does nothing
type fake struct{}

// Fake returns a "fake" terminator that just outputs a message upon instance termination
func Fake() chaosmonkey.Terminator {
	return fake{}
}

// Kill implements Terminator.kill, pretends to terminate an instance
func (t fake) Execute(trm chaosmonkey.Termination) error {
	ins := trm.Instance
	fmt.Printf("fakeTerminator fake-terminating: account=%s region=%s id=%s\n", ins.AccountName(), ins.RegionName(), ins.ID())
	return nil
}
