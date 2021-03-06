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

// MigrateToOtherRegion invokes the rds.MigrateToOtherRegion API synchronously
// api document: https://help.aliyun.com/api/rds/migratetootherregion.html
func (client *Client) MigrateToOtherRegion(request *MigrateToOtherRegionRequest) (response *MigrateToOtherRegionResponse, err error) {
	response = CreateMigrateToOtherRegionResponse()
	err = client.DoAction(request, response)
	return
}

// MigrateToOtherRegionWithChan invokes the rds.MigrateToOtherRegion API asynchronously
// api document: https://help.aliyun.com/api/rds/migratetootherregion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MigrateToOtherRegionWithChan(request *MigrateToOtherRegionRequest) (<-chan *MigrateToOtherRegionResponse, <-chan error) {
	responseChan := make(chan *MigrateToOtherRegionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.MigrateToOtherRegion(request)
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

// MigrateToOtherRegionWithCallback invokes the rds.MigrateToOtherRegion API asynchronously
// api document: https://help.aliyun.com/api/rds/migratetootherregion.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MigrateToOtherRegionWithCallback(request *MigrateToOtherRegionRequest, callback func(response *MigrateToOtherRegionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *MigrateToOtherRegionResponse
		var err error
		defer close(result)
		response, err = client.MigrateToOtherRegion(request)
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

// MigrateToOtherRegionRequest is the request struct for api MigrateToOtherRegion
type MigrateToOtherRegionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	TargetVSwitchId      string           `position:"Query" name:"TargetVSwitchId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	TargetVpcId          string           `position:"Query" name:"TargetVpcId"`
	TargetZoneId         string           `position:"Query" name:"TargetZoneId"`
	EffectiveTime        string           `position:"Query" name:"EffectiveTime"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	TargetRegionId       string           `position:"Query" name:"TargetRegionId"`
	SwitchTime           string           `position:"Query" name:"SwitchTime"`
}

// MigrateToOtherRegionResponse is the response struct for api MigrateToOtherRegion
type MigrateToOtherRegionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateMigrateToOtherRegionRequest creates a request to invoke MigrateToOtherRegion API
func CreateMigrateToOtherRegionRequest() (request *MigrateToOtherRegionRequest) {
	request = &MigrateToOtherRegionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "MigrateToOtherRegion", "rds", "openAPI")
	return
}

// CreateMigrateToOtherRegionResponse creates a response to parse from MigrateToOtherRegion response
func CreateMigrateToOtherRegionResponse() (response *MigrateToOtherRegionResponse) {
	response = &MigrateToOtherRegionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
