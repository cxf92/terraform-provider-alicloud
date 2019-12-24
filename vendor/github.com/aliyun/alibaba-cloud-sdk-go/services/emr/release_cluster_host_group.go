package emr

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

// ReleaseClusterHostGroup invokes the emr.ReleaseClusterHostGroup API synchronously
// api document: https://help.aliyun.com/api/emr/releaseclusterhostgroup.html
func (client *Client) ReleaseClusterHostGroup(request *ReleaseClusterHostGroupRequest) (response *ReleaseClusterHostGroupResponse, err error) {
	response = CreateReleaseClusterHostGroupResponse()
	err = client.DoAction(request, response)
	return
}

// ReleaseClusterHostGroupWithChan invokes the emr.ReleaseClusterHostGroup API asynchronously
// api document: https://help.aliyun.com/api/emr/releaseclusterhostgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ReleaseClusterHostGroupWithChan(request *ReleaseClusterHostGroupRequest) (<-chan *ReleaseClusterHostGroupResponse, <-chan error) {
	responseChan := make(chan *ReleaseClusterHostGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReleaseClusterHostGroup(request)
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

// ReleaseClusterHostGroupWithCallback invokes the emr.ReleaseClusterHostGroup API asynchronously
// api document: https://help.aliyun.com/api/emr/releaseclusterhostgroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ReleaseClusterHostGroupWithCallback(request *ReleaseClusterHostGroupRequest, callback func(response *ReleaseClusterHostGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReleaseClusterHostGroupResponse
		var err error
		defer close(result)
		response, err = client.ReleaseClusterHostGroup(request)
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

// ReleaseClusterHostGroupRequest is the request struct for api ReleaseClusterHostGroup
type ReleaseClusterHostGroupRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string           `position:"Query" name:"ClusterId"`
	HostGroupId     string           `position:"Query" name:"HostGroupId"`
	InstanceIdList  string           `position:"Query" name:"InstanceIdList"`
}

// ReleaseClusterHostGroupResponse is the response struct for api ReleaseClusterHostGroup
type ReleaseClusterHostGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateReleaseClusterHostGroupRequest creates a request to invoke ReleaseClusterHostGroup API
func CreateReleaseClusterHostGroupRequest() (request *ReleaseClusterHostGroupRequest) {
	request = &ReleaseClusterHostGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ReleaseClusterHostGroup", "emr", "openAPI")
	return
}

// CreateReleaseClusterHostGroupResponse creates a response to parse from ReleaseClusterHostGroup response
func CreateReleaseClusterHostGroupResponse() (response *ReleaseClusterHostGroupResponse) {
	response = &ReleaseClusterHostGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
