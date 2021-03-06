package vpc

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

// ModifyPhysicalConnectionAttribute invokes the vpc.ModifyPhysicalConnectionAttribute API synchronously
// api document: https://help.aliyun.com/api/vpc/modifyphysicalconnectionattribute.html
func (client *Client) ModifyPhysicalConnectionAttribute(request *ModifyPhysicalConnectionAttributeRequest) (response *ModifyPhysicalConnectionAttributeResponse, err error) {
	response = CreateModifyPhysicalConnectionAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyPhysicalConnectionAttributeWithChan invokes the vpc.ModifyPhysicalConnectionAttribute API asynchronously
// api document: https://help.aliyun.com/api/vpc/modifyphysicalconnectionattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyPhysicalConnectionAttributeWithChan(request *ModifyPhysicalConnectionAttributeRequest) (<-chan *ModifyPhysicalConnectionAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyPhysicalConnectionAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyPhysicalConnectionAttribute(request)
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

// ModifyPhysicalConnectionAttributeWithCallback invokes the vpc.ModifyPhysicalConnectionAttribute API asynchronously
// api document: https://help.aliyun.com/api/vpc/modifyphysicalconnectionattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyPhysicalConnectionAttributeWithCallback(request *ModifyPhysicalConnectionAttributeRequest, callback func(response *ModifyPhysicalConnectionAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyPhysicalConnectionAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyPhysicalConnectionAttribute(request)
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

// ModifyPhysicalConnectionAttributeRequest is the request struct for api ModifyPhysicalConnectionAttribute
type ModifyPhysicalConnectionAttributeRequest struct {
	*requests.RpcRequest
	PhysicalConnectionId          string           `position:"Query" name:"PhysicalConnectionId"`
	LineOperator                  string           `position:"Query" name:"LineOperator"`
	Bandwidth                     requests.Integer `position:"Query" name:"bandwidth"`
	PeerLocation                  string           `position:"Query" name:"PeerLocation"`
	PortType                      string           `position:"Query" name:"PortType"`
	RedundantPhysicalConnectionId string           `position:"Query" name:"RedundantPhysicalConnectionId"`
	Description                   string           `position:"Query" name:"Description"`
	Name                          string           `position:"Query" name:"Name"`
	ClientToken                   string           `position:"Query" name:"ClientToken"`
	OwnerId                       requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount          string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId               requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount                  string           `position:"Query" name:"OwnerAccount"`
	CircuitCode                   string           `position:"Query" name:"CircuitCode"`
}

// ModifyPhysicalConnectionAttributeResponse is the response struct for api ModifyPhysicalConnectionAttribute
type ModifyPhysicalConnectionAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyPhysicalConnectionAttributeRequest creates a request to invoke ModifyPhysicalConnectionAttribute API
func CreateModifyPhysicalConnectionAttributeRequest() (request *ModifyPhysicalConnectionAttributeRequest) {
	request = &ModifyPhysicalConnectionAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ModifyPhysicalConnectionAttribute", "vpc", "openAPI")
	return
}

// CreateModifyPhysicalConnectionAttributeResponse creates a response to parse from ModifyPhysicalConnectionAttribute response
func CreateModifyPhysicalConnectionAttributeResponse() (response *ModifyPhysicalConnectionAttributeResponse) {
	response = &ModifyPhysicalConnectionAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
