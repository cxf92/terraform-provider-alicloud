package gpdb

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

// ModifyDBInstanceConnectionString invokes the gpdb.ModifyDBInstanceConnectionString API synchronously
// api document: https://help.aliyun.com/api/gpdb/modifydbinstanceconnectionstring.html
func (client *Client) ModifyDBInstanceConnectionString(request *ModifyDBInstanceConnectionStringRequest) (response *ModifyDBInstanceConnectionStringResponse, err error) {
	response = CreateModifyDBInstanceConnectionStringResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDBInstanceConnectionStringWithChan invokes the gpdb.ModifyDBInstanceConnectionString API asynchronously
// api document: https://help.aliyun.com/api/gpdb/modifydbinstanceconnectionstring.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDBInstanceConnectionStringWithChan(request *ModifyDBInstanceConnectionStringRequest) (<-chan *ModifyDBInstanceConnectionStringResponse, <-chan error) {
	responseChan := make(chan *ModifyDBInstanceConnectionStringResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBInstanceConnectionString(request)
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

// ModifyDBInstanceConnectionStringWithCallback invokes the gpdb.ModifyDBInstanceConnectionString API asynchronously
// api document: https://help.aliyun.com/api/gpdb/modifydbinstanceconnectionstring.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDBInstanceConnectionStringWithCallback(request *ModifyDBInstanceConnectionStringRequest, callback func(response *ModifyDBInstanceConnectionStringResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBInstanceConnectionStringResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBInstanceConnectionString(request)
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

// ModifyDBInstanceConnectionStringRequest is the request struct for api ModifyDBInstanceConnectionString
type ModifyDBInstanceConnectionStringRequest struct {
	*requests.RpcRequest
	ConnectionStringPrefix  string `position:"Query" name:"ConnectionStringPrefix"`
	Port                    string `position:"Query" name:"Port"`
	DBInstanceId            string `position:"Query" name:"DBInstanceId"`
	CurrentConnectionString string `position:"Query" name:"CurrentConnectionString"`
}

// ModifyDBInstanceConnectionStringResponse is the response struct for api ModifyDBInstanceConnectionString
type ModifyDBInstanceConnectionStringResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDBInstanceConnectionStringRequest creates a request to invoke ModifyDBInstanceConnectionString API
func CreateModifyDBInstanceConnectionStringRequest() (request *ModifyDBInstanceConnectionStringRequest) {
	request = &ModifyDBInstanceConnectionStringRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("gpdb", "2016-05-03", "ModifyDBInstanceConnectionString", "gpdb", "openAPI")
	return
}

// CreateModifyDBInstanceConnectionStringResponse creates a response to parse from ModifyDBInstanceConnectionString response
func CreateModifyDBInstanceConnectionStringResponse() (response *ModifyDBInstanceConnectionStringResponse) {
	response = &ModifyDBInstanceConnectionStringResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
