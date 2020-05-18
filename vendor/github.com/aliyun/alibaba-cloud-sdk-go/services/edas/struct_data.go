package edas

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// Data is a nested struct in edas response
type Data struct {
	ChangeOrderId     string                    `json:"ChangeOrderId" xml:"ChangeOrderId"`
	ExtSlbIp          string                    `json:"ExtSlbIp" xml:"ExtSlbIp"`
	UpdateTime        int64                     `json:"UpdateTime" xml:"UpdateTime"`
	ExtSlbId          string                    `json:"ExtSlbId" xml:"ExtSlbId"`
	ExtSlbName        string                    `json:"ExtSlbName" xml:"ExtSlbName"`
	OversoldFactor    int                       `json:"OversoldFactor" xml:"OversoldFactor"`
	SlbName           string                    `json:"SlbName" xml:"SlbName"`
	VServerGroupId    string                    `json:"VServerGroupId" xml:"VServerGroupId"`
	ExtVServerGroupId string                    `json:"ExtVServerGroupId" xml:"ExtVServerGroupId"`
	SlbPort           int                       `json:"SlbPort" xml:"SlbPort"`
	SlbId             string                    `json:"SlbId" xml:"SlbId"`
	ClusterType       int                       `json:"ClusterType" xml:"ClusterType"`
	VpcId             string                    `json:"VpcId" xml:"VpcId"`
	SlbIp             string                    `json:"SlbIp" xml:"SlbIp"`
	RuleList          RuleListInGetScalingRules `json:"RuleList" xml:"RuleList"`
}