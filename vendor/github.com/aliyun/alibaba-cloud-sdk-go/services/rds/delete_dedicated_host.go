package rds

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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DeleteDedicatedHost invokes the rds.DeleteDedicatedHost API synchronously
// api document: https://help.aliyun.com/api/rds/deletededicatedhost.html
func (client *Client) DeleteDedicatedHost(request *DeleteDedicatedHostRequest) (response *DeleteDedicatedHostResponse, err error) {
	response = CreateDeleteDedicatedHostResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDedicatedHostWithChan invokes the rds.DeleteDedicatedHost API asynchronously
// api document: https://help.aliyun.com/api/rds/deletededicatedhost.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteDedicatedHostWithChan(request *DeleteDedicatedHostRequest) (<-chan *DeleteDedicatedHostResponse, <-chan error) {
	responseChan := make(chan *DeleteDedicatedHostResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDedicatedHost(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DeleteDedicatedHostWithCallback invokes the rds.DeleteDedicatedHost API asynchronously
// api document: https://help.aliyun.com/api/rds/deletededicatedhost.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteDedicatedHostWithCallback(request *DeleteDedicatedHostRequest, callback func(response *DeleteDedicatedHostResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDedicatedHostResponse
		var err error
		defer close(result)
		response, err = client.DeleteDedicatedHost(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DeleteDedicatedHostRequest is the request struct for api DeleteDedicatedHost
type DeleteDedicatedHostRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DedicatedHostId      string           `position:"Query" name:"DedicatedHostId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteDedicatedHostResponse is the response struct for api DeleteDedicatedHost
type DeleteDedicatedHostResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteDedicatedHostRequest creates a request to invoke DeleteDedicatedHost API
func CreateDeleteDedicatedHostRequest() (request *DeleteDedicatedHostRequest) {
	request = &DeleteDedicatedHostRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DeleteDedicatedHost", "rds", "openAPI")
	return
}

// CreateDeleteDedicatedHostResponse creates a response to parse from DeleteDedicatedHost response
func CreateDeleteDedicatedHostResponse() (response *DeleteDedicatedHostResponse) {
	response = &DeleteDedicatedHostResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}