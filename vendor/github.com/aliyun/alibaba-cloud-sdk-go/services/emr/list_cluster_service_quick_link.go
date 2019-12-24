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

// ListClusterServiceQuickLink invokes the emr.ListClusterServiceQuickLink API synchronously
// api document: https://help.aliyun.com/api/emr/listclusterservicequicklink.html
func (client *Client) ListClusterServiceQuickLink(request *ListClusterServiceQuickLinkRequest) (response *ListClusterServiceQuickLinkResponse, err error) {
	response = CreateListClusterServiceQuickLinkResponse()
	err = client.DoAction(request, response)
	return
}

// ListClusterServiceQuickLinkWithChan invokes the emr.ListClusterServiceQuickLink API asynchronously
// api document: https://help.aliyun.com/api/emr/listclusterservicequicklink.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListClusterServiceQuickLinkWithChan(request *ListClusterServiceQuickLinkRequest) (<-chan *ListClusterServiceQuickLinkResponse, <-chan error) {
	responseChan := make(chan *ListClusterServiceQuickLinkResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListClusterServiceQuickLink(request)
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

// ListClusterServiceQuickLinkWithCallback invokes the emr.ListClusterServiceQuickLink API asynchronously
// api document: https://help.aliyun.com/api/emr/listclusterservicequicklink.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListClusterServiceQuickLinkWithCallback(request *ListClusterServiceQuickLinkRequest, callback func(response *ListClusterServiceQuickLinkResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListClusterServiceQuickLinkResponse
		var err error
		defer close(result)
		response, err = client.ListClusterServiceQuickLink(request)
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

// ListClusterServiceQuickLinkRequest is the request struct for api ListClusterServiceQuickLink
type ListClusterServiceQuickLinkRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string           `position:"Query" name:"ClusterId"`
	ServiceName     string           `position:"Query" name:"ServiceName"`
}

// ListClusterServiceQuickLinkResponse is the response struct for api ListClusterServiceQuickLink
type ListClusterServiceQuickLinkResponse struct {
	*responses.BaseResponse
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	QuickLinkList QuickLinkList `json:"QuickLinkList" xml:"QuickLinkList"`
}

// CreateListClusterServiceQuickLinkRequest creates a request to invoke ListClusterServiceQuickLink API
func CreateListClusterServiceQuickLinkRequest() (request *ListClusterServiceQuickLinkRequest) {
	request = &ListClusterServiceQuickLinkRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ListClusterServiceQuickLink", "emr", "openAPI")
	return
}

// CreateListClusterServiceQuickLinkResponse creates a response to parse from ListClusterServiceQuickLink response
func CreateListClusterServiceQuickLinkResponse() (response *ListClusterServiceQuickLinkResponse) {
	response = &ListClusterServiceQuickLinkResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
