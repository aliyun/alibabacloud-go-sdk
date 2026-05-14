// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

type Client struct {
	openapi.Client
	DisableSDKError *bool
	EnableValidate  *bool
}

func NewClient(config *openapiutil.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapiutil.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = dara.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("dyvmsapi"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !dara.IsNil(endpoint) {
		_result = endpoint
		return _result, _err
	}

	if !dara.IsNil(endpointMap) && !dara.IsNil(endpointMap[dara.StringValue(regionId)]) {
		_result = endpointMap[dara.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := openapiutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Binds multiple real numbers to a service instance at a time.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 200 times per second per account.
//
// @param request - AddVirtualNumberRelationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddVirtualNumberRelationResponse
func (client *Client) AddVirtualNumberRelationWithOptions(request *AddVirtualNumberRelationRequest, runtime *dara.RuntimeOptions) (_result *AddVirtualNumberRelationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CorpNameList) {
		query["CorpNameList"] = request.CorpNameList
	}

	if !dara.IsNil(request.NumberList) {
		query["NumberList"] = request.NumberList
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNum) {
		query["PhoneNum"] = request.PhoneNum
	}

	if !dara.IsNil(request.ProdCode) {
		query["ProdCode"] = request.ProdCode
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RouteType) {
		query["RouteType"] = request.RouteType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddVirtualNumberRelation"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddVirtualNumberRelationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Binds multiple real numbers to a service instance at a time.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 200 times per second per account.
//
// @param request - AddVirtualNumberRelationRequest
//
// @return AddVirtualNumberRelationResponse
func (client *Client) AddVirtualNumberRelation(request *AddVirtualNumberRelationRequest) (_result *AddVirtualNumberRelationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddVirtualNumberRelationResponse{}
	_body, _err := client.AddVirtualNumberRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Initiates outbound robocall tasks in a batch. You can set up to 100 numbers in a task.
//
// Description:
//
//	  In an intelligent speech interaction task, you can use the robot communication scripts preset in the Voice Messaging Service console, or invoke the callback function to return the response mode configured by the business party in each call.
//
//		- The BatchRobotSmartCall operation is used to initiate an outbound robocall task by using the robot communication scripts preset in the Voice Messaging Service console.
//
// ## Prerequisites
//
//   - You have passed the real-name verification for an enterprise user and passed the enterprise qualification review.
//
//   - You have purchased numbers in the [Voice Messaging Service console](https://dyvms.console.aliyun.com/dyvms.htm#/number/normal).
//
//   - You have added communication scripts on the [Communication script management](https://dyvms.console.aliyun.com/dyvms.htm#/smart-call/saas/robot/list) page, and the communication scripts have been approved.
//
// > Before you call this operation, make sure that you are familiar with the [billing](https://www.aliyun.com/price/product#/vms/detail) of Voice Messaging Service (VMS).
//
// @param request - BatchRobotSmartCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchRobotSmartCallResponse
func (client *Client) BatchRobotSmartCallWithOptions(request *BatchRobotSmartCallRequest, runtime *dara.RuntimeOptions) (_result *BatchRobotSmartCallResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CalledNumber) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !dara.IsNil(request.CalledShowNumber) {
		query["CalledShowNumber"] = request.CalledShowNumber
	}

	if !dara.IsNil(request.CorpName) {
		query["CorpName"] = request.CorpName
	}

	if !dara.IsNil(request.DialogId) {
		query["DialogId"] = request.DialogId
	}

	if !dara.IsNil(request.EarlyMediaAsr) {
		query["EarlyMediaAsr"] = request.EarlyMediaAsr
	}

	if !dara.IsNil(request.IsSelfLine) {
		query["IsSelfLine"] = request.IsSelfLine
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ScheduleCall) {
		query["ScheduleCall"] = request.ScheduleCall
	}

	if !dara.IsNil(request.ScheduleTime) {
		query["ScheduleTime"] = request.ScheduleTime
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.TtsParam) {
		query["TtsParam"] = request.TtsParam
	}

	if !dara.IsNil(request.TtsParamHead) {
		query["TtsParamHead"] = request.TtsParamHead
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchRobotSmartCall"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchRobotSmartCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiates outbound robocall tasks in a batch. You can set up to 100 numbers in a task.
//
// Description:
//
//	  In an intelligent speech interaction task, you can use the robot communication scripts preset in the Voice Messaging Service console, or invoke the callback function to return the response mode configured by the business party in each call.
//
//		- The BatchRobotSmartCall operation is used to initiate an outbound robocall task by using the robot communication scripts preset in the Voice Messaging Service console.
//
// ## Prerequisites
//
//   - You have passed the real-name verification for an enterprise user and passed the enterprise qualification review.
//
//   - You have purchased numbers in the [Voice Messaging Service console](https://dyvms.console.aliyun.com/dyvms.htm#/number/normal).
//
//   - You have added communication scripts on the [Communication script management](https://dyvms.console.aliyun.com/dyvms.htm#/smart-call/saas/robot/list) page, and the communication scripts have been approved.
//
// > Before you call this operation, make sure that you are familiar with the [billing](https://www.aliyun.com/price/product#/vms/detail) of Voice Messaging Service (VMS).
//
// @param request - BatchRobotSmartCallRequest
//
// @return BatchRobotSmartCallResponse
func (client *Client) BatchRobotSmartCall(request *BatchRobotSmartCallRequest) (_result *BatchRobotSmartCallResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BatchRobotSmartCallResponse{}
	_body, _err := client.BatchRobotSmartCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels the two-way call that is initiated by calling the ClickToDial operation.
//
// @param request - CancelCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelCallResponse
func (client *Client) CancelCallWithOptions(request *CancelCallRequest, runtime *dara.RuntimeOptions) (_result *CancelCallResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallId) {
		query["CallId"] = request.CallId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelCall"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels the two-way call that is initiated by calling the ClickToDial operation.
//
// @param request - CancelCallRequest
//
// @return CancelCallResponse
func (client *Client) CancelCall(request *CancelCallRequest) (_result *CancelCallResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelCallResponse{}
	_body, _err := client.CancelCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels a robocall task that has not been started.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - CancelOrderRobotTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelOrderRobotTaskResponse
func (client *Client) CancelOrderRobotTaskWithOptions(request *CancelOrderRobotTaskRequest, runtime *dara.RuntimeOptions) (_result *CancelOrderRobotTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelOrderRobotTask"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelOrderRobotTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels a robocall task that has not been started.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - CancelOrderRobotTaskRequest
//
// @return CancelOrderRobotTaskResponse
func (client *Client) CancelOrderRobotTask(request *CancelOrderRobotTaskRequest) (_result *CancelOrderRobotTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelOrderRobotTaskResponse{}
	_body, _err := client.CancelOrderRobotTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Terminates a robocall task.
//
// Description:
//
// Only a task in progress can be terminated by calling the CancelRobotTask operation, and the task cannot be resumed after it is terminated.
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - CancelRobotTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelRobotTaskResponse
func (client *Client) CancelRobotTaskWithOptions(request *CancelRobotTaskRequest, runtime *dara.RuntimeOptions) (_result *CancelRobotTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CancelRobotTask"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelRobotTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Terminates a robocall task.
//
// Description:
//
// Only a task in progress can be terminated by calling the CancelRobotTask operation, and the task cannot be resumed after it is terminated.
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - CancelRobotTaskRequest
//
// @return CancelRobotTaskResponse
func (client *Client) CancelRobotTask(request *CancelRobotTaskRequest) (_result *CancelRobotTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CancelRobotTaskResponse{}
	_body, _err := client.CancelRobotTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # ChangeMediaType
//
// @param request - ChangeMediaTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeMediaTypeResponse
func (client *Client) ChangeMediaTypeWithOptions(request *ChangeMediaTypeRequest, runtime *dara.RuntimeOptions) (_result *ChangeMediaTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallId) {
		query["CallId"] = request.CallId
	}

	if !dara.IsNil(request.CalledNum) {
		query["CalledNum"] = request.CalledNum
	}

	if !dara.IsNil(request.MediaType) {
		query["MediaType"] = request.MediaType
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeMediaType"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChangeMediaTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ChangeMediaType
//
// @param request - ChangeMediaTypeRequest
//
// @return ChangeMediaTypeResponse
func (client *Client) ChangeMediaType(request *ChangeMediaTypeRequest) (_result *ChangeMediaTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ChangeMediaTypeResponse{}
	_body, _err := client.ChangeMediaTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Agent status monitoring.
//
// @param request - ClinkAgentStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClinkAgentStatusResponse
func (client *Client) ClinkAgentStatusWithOptions(request *ClinkAgentStatusRequest, runtime *dara.RuntimeOptions) (_result *ClinkAgentStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.Offset) {
		query["Offset"] = request.Offset
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClinkAgentStatus"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClinkAgentStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Agent status monitoring.
//
// @param request - ClinkAgentStatusRequest
//
// @return ClinkAgentStatusResponse
func (client *Client) ClinkAgentStatus(request *ClinkAgentStatusRequest) (_result *ClinkAgentStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClinkAgentStatusResponse{}
	_body, _err := client.ClinkAgentStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the status details of an agent.
//
// @param request - ClinkAgentStatusDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClinkAgentStatusDetailResponse
func (client *Client) ClinkAgentStatusDetailWithOptions(request *ClinkAgentStatusDetailRequest, runtime *dara.RuntimeOptions) (_result *ClinkAgentStatusDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cno) {
		query["Cno"] = request.Cno
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClinkAgentStatusDetail"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClinkAgentStatusDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the status details of an agent.
//
// @param request - ClinkAgentStatusDetailRequest
//
// @return ClinkAgentStatusDetailResponse
func (client *Client) ClinkAgentStatusDetail(request *ClinkAgentStatusDetailRequest) (_result *ClinkAgentStatusDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClinkAgentStatusDetailResponse{}
	_body, _err := client.ClinkAgentStatusDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Attach an agent phone.
//
// @param request - ClinkBindClientTelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClinkBindClientTelResponse
func (client *Client) ClinkBindClientTelWithOptions(request *ClinkBindClientTelRequest, runtime *dara.RuntimeOptions) (_result *ClinkBindClientTelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cno) {
		query["Cno"] = request.Cno
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.IsBind) {
		query["IsBind"] = request.IsBind
	}

	if !dara.IsNil(request.IsReserveTel) {
		query["IsReserveTel"] = request.IsReserveTel
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tel) {
		query["Tel"] = request.Tel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClinkBindClientTel"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClinkBindClientTelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Attach an agent phone.
//
// @param request - ClinkBindClientTelRequest
//
// @return ClinkBindClientTelResponse
func (client *Client) ClinkBindClientTel(request *ClinkBindClientTelRequest) (_result *ClinkBindClientTelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClinkBindClientTelResponse{}
	_body, _err := client.ClinkBindClientTelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an outbound call record.
//
// @param request - ClinkCdrObDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClinkCdrObDetailsResponse
func (client *Client) ClinkCdrObDetailsWithOptions(request *ClinkCdrObDetailsRequest, runtime *dara.RuntimeOptions) (_result *ClinkCdrObDetailsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.HiddenType) {
		query["HiddenType"] = request.HiddenType
	}

	if !dara.IsNil(request.MainUniqueId) {
		query["MainUniqueId"] = request.MainUniqueId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClinkCdrObDetails"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClinkCdrObDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an outbound call record.
//
// @param request - ClinkCdrObDetailsRequest
//
// @return ClinkCdrObDetailsResponse
func (client *Client) ClinkCdrObDetails(request *ClinkCdrObDetailsRequest) (_result *ClinkCdrObDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClinkCdrObDetailsResponse{}
	_body, _err := client.ClinkCdrObDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Add an agent.
//
// @param tmpReq - ClinkCreateClientRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClinkCreateClientResponse
func (client *Client) ClinkCreateClientWithOptions(tmpReq *ClinkCreateClientRequest, runtime *dara.RuntimeOptions) (_result *ClinkCreateClientResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ClinkCreateClientShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Clid) {
		request.ClidShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Clid, dara.String("Clid"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ClidArea) {
		request.ClidAreaShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ClidArea, dara.String("ClidArea"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ClidDefault) {
		request.ClidDefaultShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ClidDefault, dara.String("ClidDefault"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CloudNumberModes) {
		request.CloudNumberModesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CloudNumberModes, dara.String("CloudNumberModes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Permission) {
		request.PermissionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Permission, dara.String("Permission"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Qnos) {
		request.QnosShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Qnos, dara.String("Qnos"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ServeArea) {
		request.ServeAreaShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ServeArea, dara.String("ServeArea"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Active) {
		query["Active"] = request.Active
	}

	if !dara.IsNil(request.AreaCode) {
		query["AreaCode"] = request.AreaCode
	}

	if !dara.IsNil(request.AssignType) {
		query["AssignType"] = request.AssignType
	}

	if !dara.IsNil(request.ClidShrink) {
		query["Clid"] = request.ClidShrink
	}

	if !dara.IsNil(request.ClidAreaShrink) {
		query["ClidArea"] = request.ClidAreaShrink
	}

	if !dara.IsNil(request.ClidDefaultShrink) {
		query["ClidDefault"] = request.ClidDefaultShrink
	}

	if !dara.IsNil(request.ClidRule) {
		query["ClidRule"] = request.ClidRule
	}

	if !dara.IsNil(request.ClidType) {
		query["ClidType"] = request.ClidType
	}

	if !dara.IsNil(request.CloudNumberEnabled) {
		query["CloudNumberEnabled"] = request.CloudNumberEnabled
	}

	if !dara.IsNil(request.CloudNumberModesShrink) {
		query["CloudNumberModes"] = request.CloudNumberModesShrink
	}

	if !dara.IsNil(request.Cno) {
		query["Cno"] = request.Cno
	}

	if !dara.IsNil(request.CrmId) {
		query["CrmId"] = request.CrmId
	}

	if !dara.IsNil(request.DynamicTelGroupIdDefault) {
		query["DynamicTelGroupIdDefault"] = request.DynamicTelGroupIdDefault
	}

	if !dara.IsNil(request.DynamicTelGroupName) {
		query["DynamicTelGroupName"] = request.DynamicTelGroupName
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.HiddenTel) {
		query["HiddenTel"] = request.HiddenTel
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ObClidDefaultType) {
		query["ObClidDefaultType"] = request.ObClidDefaultType
	}

	if !dara.IsNil(request.ObHangupSms) {
		query["ObHangupSms"] = request.ObHangupSms
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.PauseLogin) {
		query["PauseLogin"] = request.PauseLogin
	}

	if !dara.IsNil(request.PermissionShrink) {
		query["Permission"] = request.PermissionShrink
	}

	if !dara.IsNil(request.QnosShrink) {
		query["Qnos"] = request.QnosShrink
	}

	if !dara.IsNil(request.RecurrentselectionType) {
		query["RecurrentselectionType"] = request.RecurrentselectionType
	}

	if !dara.IsNil(request.RecurrentselectionValue) {
		query["RecurrentselectionValue"] = request.RecurrentselectionValue
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Role) {
		query["Role"] = request.Role
	}

	if !dara.IsNil(request.ServeAreaShrink) {
		query["ServeArea"] = request.ServeAreaShrink
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.WechatMiniProgramRtc) {
		query["WechatMiniProgramRtc"] = request.WechatMiniProgramRtc
	}

	if !dara.IsNil(request.WrapupTime) {
		query["WrapupTime"] = request.WrapupTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClinkCreateClient"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClinkCreateClientResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Add an agent.
//
// @param request - ClinkCreateClientRequest
//
// @return ClinkCreateClientResponse
func (client *Client) ClinkCreateClient(request *ClinkCreateClientRequest) (_result *ClinkCreateClientResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClinkCreateClientResponse{}
	_body, _err := client.ClinkCreateClientWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Add a busy status.
//
// @param request - ClinkCreateEnterprisePauseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClinkCreateEnterprisePauseResponse
func (client *Client) ClinkCreateEnterprisePauseWithOptions(request *ClinkCreateEnterprisePauseRequest, runtime *dara.RuntimeOptions) (_result *ClinkCreateEnterprisePauseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.IsDefault) {
		query["IsDefault"] = request.IsDefault
	}

	if !dara.IsNil(request.IsRest) {
		query["IsRest"] = request.IsRest
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PauseStatus) {
		query["PauseStatus"] = request.PauseStatus
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClinkCreateEnterprisePause"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClinkCreateEnterprisePauseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Add a busy status.
//
// @param request - ClinkCreateEnterprisePauseRequest
//
// @return ClinkCreateEnterprisePauseResponse
func (client *Client) ClinkCreateEnterprisePause(request *ClinkCreateEnterprisePauseRequest) (_result *ClinkCreateEnterprisePauseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClinkCreateEnterprisePauseResponse{}
	_body, _err := client.ClinkCreateEnterprisePauseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Add a phone.
//
// @param request - ClinkCreateExtenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClinkCreateExtenResponse
func (client *Client) ClinkCreateExtenWithOptions(request *ClinkCreateExtenRequest, runtime *dara.RuntimeOptions) (_result *ClinkCreateExtenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Allow) {
		query["Allow"] = request.Allow
	}

	if !dara.IsNil(request.AreaCode) {
		query["AreaCode"] = request.AreaCode
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.ExtenNumber) {
		query["ExtenNumber"] = request.ExtenNumber
	}

	if !dara.IsNil(request.IsDirect) {
		query["IsDirect"] = request.IsDirect
	}

	if !dara.IsNil(request.JittBuffer) {
		query["JittBuffer"] = request.JittBuffer
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClinkCreateExten"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClinkCreateExtenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Add a phone.
//
// @param request - ClinkCreateExtenRequest
//
// @return ClinkCreateExtenResponse
func (client *Client) ClinkCreateExten(request *ClinkCreateExtenRequest) (_result *ClinkCreateExtenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClinkCreateExtenResponse{}
	_body, _err := client.ClinkCreateExtenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Add a queue.
//
// @param tmpReq - ClinkCreateQueueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClinkCreateQueueResponse
func (client *Client) ClinkCreateQueueWithOptions(tmpReq *ClinkCreateQueueRequest, runtime *dara.RuntimeOptions) (_result *ClinkCreateQueueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ClinkCreateQueueShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.QueueMembers) {
		request.QueueMembersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QueueMembers, dara.String("QueueMembers"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ChatLocation) {
		query["ChatLocation"] = request.ChatLocation
	}

	if !dara.IsNil(request.ChatMaxWait) {
		query["ChatMaxWait"] = request.ChatMaxWait
	}

	if !dara.IsNil(request.ChatStrategy) {
		query["ChatStrategy"] = request.ChatStrategy
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.IbAllowed) {
		query["IbAllowed"] = request.IbAllowed
	}

	if !dara.IsNil(request.JoinEmpty) {
		query["JoinEmpty"] = request.JoinEmpty
	}

	if !dara.IsNil(request.MaxPauseClientFlag) {
		query["MaxPauseClientFlag"] = request.MaxPauseClientFlag
	}

	if !dara.IsNil(request.MaxPauseClientType) {
		query["MaxPauseClientType"] = request.MaxPauseClientType
	}

	if !dara.IsNil(request.MaxPauseClientValue) {
		query["MaxPauseClientValue"] = request.MaxPauseClientValue
	}

	if !dara.IsNil(request.MaxWait) {
		query["MaxWait"] = request.MaxWait
	}

	if !dara.IsNil(request.MemberTimeout) {
		query["MemberTimeout"] = request.MemberTimeout
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Qno) {
		query["Qno"] = request.Qno
	}

	if !dara.IsNil(request.QueueMembersShrink) {
		query["QueueMembers"] = request.QueueMembersShrink
	}

	if !dara.IsNil(request.QueueTimeout) {
		query["QueueTimeout"] = request.QueueTimeout
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SayCno) {
		query["SayCno"] = request.SayCno
	}

	if !dara.IsNil(request.ServiceLevel) {
		query["ServiceLevel"] = request.ServiceLevel
	}

	if !dara.IsNil(request.Strategy) {
		query["Strategy"] = request.Strategy
	}

	if !dara.IsNil(request.VipSupport) {
		query["VipSupport"] = request.VipSupport
	}

	if !dara.IsNil(request.Weight) {
		query["Weight"] = request.Weight
	}

	if !dara.IsNil(request.WrapupTime) {
		query["WrapupTime"] = request.WrapupTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClinkCreateQueue"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClinkCreateQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Add a queue.
//
// @param request - ClinkCreateQueueRequest
//
// @return ClinkCreateQueueResponse
func (client *Client) ClinkCreateQueue(request *ClinkCreateQueueRequest) (_result *ClinkCreateQueueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClinkCreateQueueResponse{}
	_body, _err := client.ClinkCreateQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an agent.
//
// @param request - ClinkDeleteClientRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClinkDeleteClientResponse
func (client *Client) ClinkDeleteClientWithOptions(request *ClinkDeleteClientRequest, runtime *dara.RuntimeOptions) (_result *ClinkDeleteClientResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cno) {
		query["Cno"] = request.Cno
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClinkDeleteClient"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClinkDeleteClientResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an agent.
//
// @param request - ClinkDeleteClientRequest
//
// @return ClinkDeleteClientResponse
func (client *Client) ClinkDeleteClient(request *ClinkDeleteClientRequest) (_result *ClinkDeleteClientResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClinkDeleteClientResponse{}
	_body, _err := client.ClinkDeleteClientWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete the phone.
//
// @param request - ClinkDeleteExtenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClinkDeleteExtenResponse
func (client *Client) ClinkDeleteExtenWithOptions(request *ClinkDeleteExtenRequest, runtime *dara.RuntimeOptions) (_result *ClinkDeleteExtenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.ExtenNumber) {
		query["ExtenNumber"] = request.ExtenNumber
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClinkDeleteExten"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClinkDeleteExtenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete the phone.
//
// @param request - ClinkDeleteExtenRequest
//
// @return ClinkDeleteExtenResponse
func (client *Client) ClinkDeleteExten(request *ClinkDeleteExtenRequest) (_result *ClinkDeleteExtenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClinkDeleteExtenResponse{}
	_body, _err := client.ClinkDeleteExtenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// View incoming call records.
//
// @param request - ClinkDescribeCdrIbRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClinkDescribeCdrIbResponse
func (client *Client) ClinkDescribeCdrIbWithOptions(request *ClinkDescribeCdrIbRequest, runtime *dara.RuntimeOptions) (_result *ClinkDescribeCdrIbResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.HiddenType) {
		query["HiddenType"] = request.HiddenType
	}

	if !dara.IsNil(request.MainUniqueId) {
		query["MainUniqueId"] = request.MainUniqueId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClinkDescribeCdrIb"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClinkDescribeCdrIbResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// View incoming call records.
//
// @param request - ClinkDescribeCdrIbRequest
//
// @return ClinkDescribeCdrIbResponse
func (client *Client) ClinkDescribeCdrIb(request *ClinkDescribeCdrIbRequest) (_result *ClinkDescribeCdrIbResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClinkDescribeCdrIbResponse{}
	_body, _err := client.ClinkDescribeCdrIbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// View incoming call record details.
//
// @param request - ClinkDescribeCdrIbDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClinkDescribeCdrIbDetailsResponse
func (client *Client) ClinkDescribeCdrIbDetailsWithOptions(request *ClinkDescribeCdrIbDetailsRequest, runtime *dara.RuntimeOptions) (_result *ClinkDescribeCdrIbDetailsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.MainUniqueId) {
		query["MainUniqueId"] = request.MainUniqueId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClinkDescribeCdrIbDetails"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClinkDescribeCdrIbDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// View incoming call record details.
//
// @param request - ClinkDescribeCdrIbDetailsRequest
//
// @return ClinkDescribeCdrIbDetailsResponse
func (client *Client) ClinkDescribeCdrIbDetails(request *ClinkDescribeCdrIbDetailsRequest) (_result *ClinkDescribeCdrIbDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClinkDescribeCdrIbDetailsResponse{}
	_body, _err := client.ClinkDescribeCdrIbDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// View outbound call records.
//
// @param request - ClinkDescribeCdrObRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClinkDescribeCdrObResponse
func (client *Client) ClinkDescribeCdrObWithOptions(request *ClinkDescribeCdrObRequest, runtime *dara.RuntimeOptions) (_result *ClinkDescribeCdrObResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.HiddenType) {
		query["HiddenType"] = request.HiddenType
	}

	if !dara.IsNil(request.MainUniqueId) {
		query["MainUniqueId"] = request.MainUniqueId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClinkDescribeCdrOb"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClinkDescribeCdrObResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// View outbound call records.
//
// @param request - ClinkDescribeCdrObRequest
//
// @return ClinkDescribeCdrObResponse
func (client *Client) ClinkDescribeCdrOb(request *ClinkDescribeCdrObRequest) (_result *ClinkDescribeCdrObResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClinkDescribeCdrObResponse{}
	_body, _err := client.ClinkDescribeCdrObWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// View outbound call record details.
//
// @param request - ClinkDescribeCdrObDetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClinkDescribeCdrObDetailsResponse
func (client *Client) ClinkDescribeCdrObDetailsWithOptions(request *ClinkDescribeCdrObDetailsRequest, runtime *dara.RuntimeOptions) (_result *ClinkDescribeCdrObDetailsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.MainUniqueId) {
		query["MainUniqueId"] = request.MainUniqueId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClinkDescribeCdrObDetails"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClinkDescribeCdrObDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// View outbound call record details.
//
// @param request - ClinkDescribeCdrObDetailsRequest
//
// @return ClinkDescribeCdrObDetailsResponse
func (client *Client) ClinkDescribeCdrObDetails(request *ClinkDescribeCdrObDetailsRequest) (_result *ClinkDescribeCdrObDetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClinkDescribeCdrObDetailsResponse{}
	_body, _err := client.ClinkDescribeCdrObDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// View the agent details.
//
// @param request - ClinkDescribeClientRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClinkDescribeClientResponse
func (client *Client) ClinkDescribeClientWithOptions(request *ClinkDescribeClientRequest, runtime *dara.RuntimeOptions) (_result *ClinkDescribeClientResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cno) {
		query["Cno"] = request.Cno
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClinkDescribeClient"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClinkDescribeClientResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// View the agent details.
//
// @param request - ClinkDescribeClientRequest
//
// @return ClinkDescribeClientResponse
func (client *Client) ClinkDescribeClient(request *ClinkDescribeClientRequest) (_result *ClinkDescribeClientResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClinkDescribeClientResponse{}
	_body, _err := client.ClinkDescribeClientWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// View queue details.
//
// @param request - ClinkDescribeQueueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClinkDescribeQueueResponse
func (client *Client) ClinkDescribeQueueWithOptions(request *ClinkDescribeQueueRequest, runtime *dara.RuntimeOptions) (_result *ClinkDescribeQueueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Qno) {
		query["Qno"] = request.Qno
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClinkDescribeQueue"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClinkDescribeQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// View queue details.
//
// @param request - ClinkDescribeQueueRequest
//
// @return ClinkDescribeQueueResponse
func (client *Client) ClinkDescribeQueue(request *ClinkDescribeQueueRequest) (_result *ClinkDescribeQueueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClinkDescribeQueueResponse{}
	_body, _err := client.ClinkDescribeQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// View call recording address.
//
// @param request - ClinkDescribeRecordFileUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClinkDescribeRecordFileUrlResponse
func (client *Client) ClinkDescribeRecordFileUrlWithOptions(request *ClinkDescribeRecordFileUrlRequest, runtime *dara.RuntimeOptions) (_result *ClinkDescribeRecordFileUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Download) {
		query["Download"] = request.Download
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.MainUniqueId) {
		query["MainUniqueId"] = request.MainUniqueId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RecordSide) {
		query["RecordSide"] = request.RecordSide
	}

	if !dara.IsNil(request.RecordType) {
		query["RecordType"] = request.RecordType
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClinkDescribeRecordFileUrl"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClinkDescribeRecordFileUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// View call recording address.
//
// @param request - ClinkDescribeRecordFileUrlRequest
//
// @return ClinkDescribeRecordFileUrlResponse
func (client *Client) ClinkDescribeRecordFileUrl(request *ClinkDescribeRecordFileUrlRequest) (_result *ClinkDescribeRecordFileUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClinkDescribeRecordFileUrlResponse{}
	_body, _err := client.ClinkDescribeRecordFileUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query the details of customer call records.
//
// @param request - ClinkDetailCdrIbRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClinkDetailCdrIbResponse
func (client *Client) ClinkDetailCdrIbWithOptions(request *ClinkDetailCdrIbRequest, runtime *dara.RuntimeOptions) (_result *ClinkDetailCdrIbResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.HiddenType) {
		query["HiddenType"] = request.HiddenType
	}

	if !dara.IsNil(request.MainUniqueId) {
		query["MainUniqueId"] = request.MainUniqueId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClinkDetailCdrIb"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClinkDetailCdrIbResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the details of customer call records.
//
// @param request - ClinkDetailCdrIbRequest
//
// @return ClinkDetailCdrIbResponse
func (client *Client) ClinkDetailCdrIb(request *ClinkDetailCdrIbRequest) (_result *ClinkDetailCdrIbResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClinkDetailCdrIbResponse{}
	_body, _err := client.ClinkDetailCdrIbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query the customer call record list.
//
// @param request - ClinkListCdrIbRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClinkListCdrIbResponse
func (client *Client) ClinkListCdrIbWithOptions(request *ClinkListCdrIbRequest, runtime *dara.RuntimeOptions) (_result *ClinkListCdrIbResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BridgeDuration) {
		query["BridgeDuration"] = request.BridgeDuration
	}

	if !dara.IsNil(request.BridgeDurationEnd) {
		query["BridgeDurationEnd"] = request.BridgeDurationEnd
	}

	if !dara.IsNil(request.BridgeTime) {
		query["BridgeTime"] = request.BridgeTime
	}

	if !dara.IsNil(request.BridgeTimeEnd) {
		query["BridgeTimeEnd"] = request.BridgeTimeEnd
	}

	if !dara.IsNil(request.ClientNumber) {
		query["ClientNumber"] = request.ClientNumber
	}

	if !dara.IsNil(request.Cnos) {
		query["Cnos"] = request.Cnos
	}

	if !dara.IsNil(request.CustomerNumber) {
		query["CustomerNumber"] = request.CustomerNumber
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EndTimeEnd) {
		query["EndTimeEnd"] = request.EndTimeEnd
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.FirstCallCno) {
		query["FirstCallCno"] = request.FirstCallCno
	}

	if !dara.IsNil(request.FirstCallNumber) {
		query["FirstCallNumber"] = request.FirstCallNumber
	}

	if !dara.IsNil(request.FirstCallQno) {
		query["FirstCallQno"] = request.FirstCallQno
	}

	if !dara.IsNil(request.HiddenType) {
		query["HiddenType"] = request.HiddenType
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.MainUniqueId) {
		query["MainUniqueId"] = request.MainUniqueId
	}

	if !dara.IsNil(request.Offset) {
		query["Offset"] = request.Offset
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Qnos) {
		query["Qnos"] = request.Qnos
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ScrollId) {
		query["ScrollId"] = request.ScrollId
	}

	if !dara.IsNil(request.ScrollSearch) {
		query["ScrollSearch"] = request.ScrollSearch
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StartTimeEnd) {
		query["StartTimeEnd"] = request.StartTimeEnd
	}

	if !dara.IsNil(request.Statuses) {
		query["Statuses"] = request.Statuses
	}

	if !dara.IsNil(request.TotalDuration) {
		query["TotalDuration"] = request.TotalDuration
	}

	if !dara.IsNil(request.TotalDurationEnd) {
		query["TotalDurationEnd"] = request.TotalDurationEnd
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClinkListCdrIb"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClinkListCdrIbResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the customer call record list.
//
// @param request - ClinkListCdrIbRequest
//
// @return ClinkListCdrIbResponse
func (client *Client) ClinkListCdrIb(request *ClinkListCdrIbRequest) (_result *ClinkListCdrIbResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClinkListCdrIbResponse{}
	_body, _err := client.ClinkListCdrIbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of agent answering records.
//
// @param request - ClinkListCdrIbAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClinkListCdrIbAgentResponse
func (client *Client) ClinkListCdrIbAgentWithOptions(request *ClinkListCdrIbAgentRequest, runtime *dara.RuntimeOptions) (_result *ClinkListCdrIbAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentAnswerInTime) {
		query["AgentAnswerInTime"] = request.AgentAnswerInTime
	}

	if !dara.IsNil(request.BridgeDuration) {
		query["BridgeDuration"] = request.BridgeDuration
	}

	if !dara.IsNil(request.BridgeDurationEnd) {
		query["BridgeDurationEnd"] = request.BridgeDurationEnd
	}

	if !dara.IsNil(request.BridgeTime) {
		query["BridgeTime"] = request.BridgeTime
	}

	if !dara.IsNil(request.BridgeTimeEnd) {
		query["BridgeTimeEnd"] = request.BridgeTimeEnd
	}

	if !dara.IsNil(request.ClientNumber) {
		query["ClientNumber"] = request.ClientNumber
	}

	if !dara.IsNil(request.Cno) {
		query["Cno"] = request.Cno
	}

	if !dara.IsNil(request.CustomerNumber) {
		query["CustomerNumber"] = request.CustomerNumber
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EndTimeEnd) {
		query["EndTimeEnd"] = request.EndTimeEnd
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.HiddenType) {
		query["HiddenType"] = request.HiddenType
	}

	if !dara.IsNil(request.HotlineName) {
		query["HotlineName"] = request.HotlineName
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.MainUniqueId) {
		query["MainUniqueId"] = request.MainUniqueId
	}

	if !dara.IsNil(request.Offset) {
		query["Offset"] = request.Offset
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Qno) {
		query["Qno"] = request.Qno
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ScrollId) {
		query["ScrollId"] = request.ScrollId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StartTimeEnd) {
		query["StartTimeEnd"] = request.StartTimeEnd
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClinkListCdrIbAgent"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClinkListCdrIbAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of agent answering records.
//
// @param request - ClinkListCdrIbAgentRequest
//
// @return ClinkListCdrIbAgentResponse
func (client *Client) ClinkListCdrIbAgent(request *ClinkListCdrIbAgentRequest) (_result *ClinkListCdrIbAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClinkListCdrIbAgentResponse{}
	_body, _err := client.ClinkListCdrIbAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query the list of inbound call records.
//
// @param request - ClinkListCdrIbsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClinkListCdrIbsResponse
func (client *Client) ClinkListCdrIbsWithOptions(request *ClinkListCdrIbsRequest, runtime *dara.RuntimeOptions) (_result *ClinkListCdrIbsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientNumber) {
		query["ClientNumber"] = request.ClientNumber
	}

	if !dara.IsNil(request.Cno) {
		query["Cno"] = request.Cno
	}

	if !dara.IsNil(request.CustomerNumber) {
		query["CustomerNumber"] = request.CustomerNumber
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.HiddenType) {
		query["HiddenType"] = request.HiddenType
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.MainUniqueId) {
		query["MainUniqueId"] = request.MainUniqueId
	}

	if !dara.IsNil(request.Offset) {
		query["Offset"] = request.Offset
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Qno) {
		query["Qno"] = request.Qno
	}

	if !dara.IsNil(request.QueueAnswerInTime) {
		query["QueueAnswerInTime"] = request.QueueAnswerInTime
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClinkListCdrIbs"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClinkListCdrIbsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the list of inbound call records.
//
// @param request - ClinkListCdrIbsRequest
//
// @return ClinkListCdrIbsResponse
func (client *Client) ClinkListCdrIbs(request *ClinkListCdrIbsRequest) (_result *ClinkListCdrIbsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClinkListCdrIbsResponse{}
	_body, _err := client.ClinkListCdrIbsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of outbound call records.
//
// @param request - ClinkListCdrObRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClinkListCdrObResponse
func (client *Client) ClinkListCdrObWithOptions(request *ClinkListCdrObRequest, runtime *dara.RuntimeOptions) (_result *ClinkListCdrObResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BridgeDuration) {
		query["BridgeDuration"] = request.BridgeDuration
	}

	if !dara.IsNil(request.BridgeDurationEnd) {
		query["BridgeDurationEnd"] = request.BridgeDurationEnd
	}

	if !dara.IsNil(request.BridgeTime) {
		query["BridgeTime"] = request.BridgeTime
	}

	if !dara.IsNil(request.BridgeTimeEnd) {
		query["BridgeTimeEnd"] = request.BridgeTimeEnd
	}

	if !dara.IsNil(request.CallType) {
		query["CallType"] = request.CallType
	}

	if !dara.IsNil(request.ClientNumber) {
		query["ClientNumber"] = request.ClientNumber
	}

	if !dara.IsNil(request.Cnos) {
		query["Cnos"] = request.Cnos
	}

	if !dara.IsNil(request.CustomerNumber) {
		query["CustomerNumber"] = request.CustomerNumber
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EndTimeEnd) {
		query["EndTimeEnd"] = request.EndTimeEnd
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.HiddenType) {
		query["HiddenType"] = request.HiddenType
	}

	if !dara.IsNil(request.LeftClid) {
		query["LeftClid"] = request.LeftClid
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.MainUniqueId) {
		query["MainUniqueId"] = request.MainUniqueId
	}

	if !dara.IsNil(request.Offset) {
		query["Offset"] = request.Offset
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Qnos) {
		query["Qnos"] = request.Qnos
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ScrollId) {
		query["ScrollId"] = request.ScrollId
	}

	if !dara.IsNil(request.ScrollSearch) {
		query["ScrollSearch"] = request.ScrollSearch
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StartTimeEnd) {
		query["StartTimeEnd"] = request.StartTimeEnd
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.TotalDuration) {
		query["TotalDuration"] = request.TotalDuration
	}

	if !dara.IsNil(request.TotalDurationEnd) {
		query["TotalDurationEnd"] = request.TotalDurationEnd
	}

	if !dara.IsNil(request.Xnumber) {
		query["Xnumber"] = request.Xnumber
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClinkListCdrOb"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClinkListCdrObResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of outbound call records.
//
// @param request - ClinkListCdrObRequest
//
// @return ClinkListCdrObResponse
func (client *Client) ClinkListCdrOb(request *ClinkListCdrObRequest) (_result *ClinkListCdrObResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClinkListCdrObResponse{}
	_body, _err := client.ClinkListCdrObWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the outbound call record list.
//
// @param request - ClinkListCdrObsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClinkListCdrObsResponse
func (client *Client) ClinkListCdrObsWithOptions(request *ClinkListCdrObsRequest, runtime *dara.RuntimeOptions) (_result *ClinkListCdrObsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AssociatedId) {
		query["AssociatedId"] = request.AssociatedId
	}

	if !dara.IsNil(request.City) {
		query["City"] = request.City
	}

	if !dara.IsNil(request.ClientNumber) {
		query["ClientNumber"] = request.ClientNumber
	}

	if !dara.IsNil(request.Cno) {
		query["Cno"] = request.Cno
	}

	if !dara.IsNil(request.CustomerNumber) {
		query["CustomerNumber"] = request.CustomerNumber
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Evaluation) {
		query["Evaluation"] = request.Evaluation
	}

	if !dara.IsNil(request.HiddenType) {
		query["HiddenType"] = request.HiddenType
	}

	if !dara.IsNil(request.Hotline) {
		query["Hotline"] = request.Hotline
	}

	if !dara.IsNil(request.IdType) {
		query["IdType"] = request.IdType
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.MainUniqueId) {
		query["MainUniqueId"] = request.MainUniqueId
	}

	if !dara.IsNil(request.Mark) {
		query["Mark"] = request.Mark
	}

	if !dara.IsNil(request.Offset) {
		query["Offset"] = request.Offset
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Province) {
		query["Province"] = request.Province
	}

	if !dara.IsNil(request.QueueAnswerInTime) {
		query["QueueAnswerInTime"] = request.QueueAnswerInTime
	}

	if !dara.IsNil(request.RequestUniqueId) {
		query["RequestUniqueId"] = request.RequestUniqueId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClinkListCdrObs"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClinkListCdrObsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the outbound call record list.
//
// @param request - ClinkListCdrObsRequest
//
// @return ClinkListCdrObsResponse
func (client *Client) ClinkListCdrObs(request *ClinkListCdrObsRequest) (_result *ClinkListCdrObsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClinkListCdrObsResponse{}
	_body, _err := client.ClinkListCdrObsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of agents.
//
// @param request - ClinkListClientsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClinkListClientsResponse
func (client *Client) ClinkListClientsWithOptions(request *ClinkListClientsRequest, runtime *dara.RuntimeOptions) (_result *ClinkListClientsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Active) {
		query["Active"] = request.Active
	}

	if !dara.IsNil(request.BindTel) {
		query["BindTel"] = request.BindTel
	}

	if !dara.IsNil(request.Clid) {
		query["Clid"] = request.Clid
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.Offset) {
		query["Offset"] = request.Offset
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Qno) {
		query["Qno"] = request.Qno
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.UpdateEndTime) {
		query["UpdateEndTime"] = request.UpdateEndTime
	}

	if !dara.IsNil(request.UpdateStartTime) {
		query["UpdateStartTime"] = request.UpdateStartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClinkListClients"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClinkListClientsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of agents.
//
// @param request - ClinkListClientsRequest
//
// @return ClinkListClientsResponse
func (client *Client) ClinkListClients(request *ClinkListClientsRequest) (_result *ClinkListClientsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClinkListClientsResponse{}
	_body, _err := client.ClinkListClientsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the busy status list.
//
// @param request - ClinkListEnterprisePausesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClinkListEnterprisePausesResponse
func (client *Client) ClinkListEnterprisePausesWithOptions(request *ClinkListEnterprisePausesRequest, runtime *dara.RuntimeOptions) (_result *ClinkListEnterprisePausesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClinkListEnterprisePauses"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClinkListEnterprisePausesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the busy status list.
//
// @param request - ClinkListEnterprisePausesRequest
//
// @return ClinkListEnterprisePausesResponse
func (client *Client) ClinkListEnterprisePauses(request *ClinkListEnterprisePausesRequest) (_result *ClinkListEnterprisePausesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClinkListEnterprisePausesResponse{}
	_body, _err := client.ClinkListEnterprisePausesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query the phone list.
//
// @param request - ClinkListExtensRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClinkListExtensResponse
func (client *Client) ClinkListExtensWithOptions(request *ClinkListExtensRequest, runtime *dara.RuntimeOptions) (_result *ClinkListExtensResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AreaCode) {
		query["AreaCode"] = request.AreaCode
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.Offset) {
		query["Offset"] = request.Offset
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClinkListExtens"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClinkListExtensResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the phone list.
//
// @param request - ClinkListExtensRequest
//
// @return ClinkListExtensResponse
func (client *Client) ClinkListExtens(request *ClinkListExtensRequest) (_result *ClinkListExtensResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClinkListExtensResponse{}
	_body, _err := client.ClinkListExtensWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the voice navigation edge zone list.
//
// @param request - ClinkListIvrNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClinkListIvrNodesResponse
func (client *Client) ClinkListIvrNodesWithOptions(request *ClinkListIvrNodesRequest, runtime *dara.RuntimeOptions) (_result *ClinkListIvrNodesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.IvrName) {
		query["IvrName"] = request.IvrName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClinkListIvrNodes"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClinkListIvrNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the voice navigation edge zone list.
//
// @param request - ClinkListIvrNodesRequest
//
// @return ClinkListIvrNodesResponse
func (client *Client) ClinkListIvrNodes(request *ClinkListIvrNodesRequest) (_result *ClinkListIvrNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClinkListIvrNodesResponse{}
	_body, _err := client.ClinkListIvrNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the voice navigation list.
//
// @param request - ClinkListIvrsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClinkListIvrsResponse
func (client *Client) ClinkListIvrsWithOptions(request *ClinkListIvrsRequest, runtime *dara.RuntimeOptions) (_result *ClinkListIvrsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClinkListIvrs"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClinkListIvrsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the voice navigation list.
//
// @param request - ClinkListIvrsRequest
//
// @return ClinkListIvrsResponse
func (client *Client) ClinkListIvrs(request *ClinkListIvrsRequest) (_result *ClinkListIvrsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClinkListIvrsResponse{}
	_body, _err := client.ClinkListIvrsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the queue list.
//
// @param request - ClinkListQueuesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClinkListQueuesResponse
func (client *Client) ClinkListQueuesWithOptions(request *ClinkListQueuesRequest, runtime *dara.RuntimeOptions) (_result *ClinkListQueuesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.Offset) {
		query["Offset"] = request.Offset
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClinkListQueues"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClinkListQueuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the queue list.
//
// @param request - ClinkListQueuesRequest
//
// @return ClinkListQueuesResponse
func (client *Client) ClinkListQueues(request *ClinkListQueuesRequest) (_result *ClinkListQueuesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClinkListQueuesResponse{}
	_body, _err := client.ClinkListQueuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Inbound report - call statistics.
//
// @param request - ClinkStatIbRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClinkStatIbResponse
func (client *Client) ClinkStatIbWithOptions(request *ClinkStatIbRequest, runtime *dara.RuntimeOptions) (_result *ClinkStatIbResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Date) {
		query["Date"] = request.Date
	}

	if !dara.IsNil(request.DateEnd) {
		query["DateEnd"] = request.DateEnd
	}

	if !dara.IsNil(request.EndHour) {
		query["EndHour"] = request.EndHour
	}

	if !dara.IsNil(request.EndMinute) {
		query["EndMinute"] = request.EndMinute
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Fields) {
		query["Fields"] = request.Fields
	}

	if !dara.IsNil(request.Hotlines) {
		query["Hotlines"] = request.Hotlines
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.StartHour) {
		query["StartHour"] = request.StartHour
	}

	if !dara.IsNil(request.StartMinute) {
		query["StartMinute"] = request.StartMinute
	}

	if !dara.IsNil(request.StatisticMethod) {
		query["StatisticMethod"] = request.StatisticMethod
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClinkStatIb"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClinkStatIbResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Inbound report - call statistics.
//
// @param request - ClinkStatIbRequest
//
// @return ClinkStatIbResponse
func (client *Client) ClinkStatIb(request *ClinkStatIbRequest) (_result *ClinkStatIbResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClinkStatIbResponse{}
	_body, _err := client.ClinkStatIbWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Update an agent.
//
// @param tmpReq - ClinkUpdateClientRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClinkUpdateClientResponse
func (client *Client) ClinkUpdateClientWithOptions(tmpReq *ClinkUpdateClientRequest, runtime *dara.RuntimeOptions) (_result *ClinkUpdateClientResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ClinkUpdateClientShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Clid) {
		request.ClidShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Clid, dara.String("Clid"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ClidArea) {
		request.ClidAreaShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ClidArea, dara.String("ClidArea"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ClidDefault) {
		request.ClidDefaultShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ClidDefault, dara.String("ClidDefault"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CloudNumberModes) {
		request.CloudNumberModesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CloudNumberModes, dara.String("CloudNumberModes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Permission) {
		request.PermissionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Permission, dara.String("Permission"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Qnos) {
		request.QnosShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Qnos, dara.String("Qnos"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.ServeArea) {
		request.ServeAreaShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ServeArea, dara.String("ServeArea"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Active) {
		query["Active"] = request.Active
	}

	if !dara.IsNil(request.AreaCode) {
		query["AreaCode"] = request.AreaCode
	}

	if !dara.IsNil(request.AssignType) {
		query["AssignType"] = request.AssignType
	}

	if !dara.IsNil(request.ClidShrink) {
		query["Clid"] = request.ClidShrink
	}

	if !dara.IsNil(request.ClidAreaShrink) {
		query["ClidArea"] = request.ClidAreaShrink
	}

	if !dara.IsNil(request.ClidDefaultShrink) {
		query["ClidDefault"] = request.ClidDefaultShrink
	}

	if !dara.IsNil(request.ClidRule) {
		query["ClidRule"] = request.ClidRule
	}

	if !dara.IsNil(request.ClidType) {
		query["ClidType"] = request.ClidType
	}

	if !dara.IsNil(request.CloudNumberEnabled) {
		query["CloudNumberEnabled"] = request.CloudNumberEnabled
	}

	if !dara.IsNil(request.CloudNumberModesShrink) {
		query["CloudNumberModes"] = request.CloudNumberModesShrink
	}

	if !dara.IsNil(request.Cno) {
		query["Cno"] = request.Cno
	}

	if !dara.IsNil(request.CrmId) {
		query["CrmId"] = request.CrmId
	}

	if !dara.IsNil(request.DynamicTelGroupIdDefault) {
		query["DynamicTelGroupIdDefault"] = request.DynamicTelGroupIdDefault
	}

	if !dara.IsNil(request.DynamicTelGroupName) {
		query["DynamicTelGroupName"] = request.DynamicTelGroupName
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.HiddenTel) {
		query["HiddenTel"] = request.HiddenTel
	}

	if !dara.IsNil(request.IbWrapupTime) {
		query["IbWrapupTime"] = request.IbWrapupTime
	}

	if !dara.IsNil(request.IbWrapupType) {
		query["IbWrapupType"] = request.IbWrapupType
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ObClidDefaultType) {
		query["ObClidDefaultType"] = request.ObClidDefaultType
	}

	if !dara.IsNil(request.ObHangupSms) {
		query["ObHangupSms"] = request.ObHangupSms
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.PauseLogin) {
		query["PauseLogin"] = request.PauseLogin
	}

	if !dara.IsNil(request.PermissionShrink) {
		query["Permission"] = request.PermissionShrink
	}

	if !dara.IsNil(request.QnosShrink) {
		query["Qnos"] = request.QnosShrink
	}

	if !dara.IsNil(request.RecurrentselectionType) {
		query["RecurrentselectionType"] = request.RecurrentselectionType
	}

	if !dara.IsNil(request.RecurrentselectionValue) {
		query["RecurrentselectionValue"] = request.RecurrentselectionValue
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Role) {
		query["Role"] = request.Role
	}

	if !dara.IsNil(request.ServeAreaShrink) {
		query["ServeArea"] = request.ServeAreaShrink
	}

	if !dara.IsNil(request.WechatMiniProgramRtc) {
		query["WechatMiniProgramRtc"] = request.WechatMiniProgramRtc
	}

	if !dara.IsNil(request.WrapupTime) {
		query["WrapupTime"] = request.WrapupTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ClinkUpdateClient"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ClinkUpdateClientResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update an agent.
//
// @param request - ClinkUpdateClientRequest
//
// @return ClinkUpdateClientResponse
func (client *Client) ClinkUpdateClient(request *ClinkUpdateClientRequest) (_result *ClinkUpdateClientResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ClinkUpdateClientResponse{}
	_body, _err := client.ClinkUpdateClientWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Publish the agent online through this interface.
//
// @param request - CloudAgentLoginRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudAgentLoginResponse
func (client *Client) CloudAgentLoginWithOptions(request *CloudAgentLoginRequest, runtime *dara.RuntimeOptions) (_result *CloudAgentLoginResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BindTel) {
		query["BindTel"] = request.BindTel
	}

	if !dara.IsNil(request.BindType) {
		query["BindType"] = request.BindType
	}

	if !dara.IsNil(request.Cno) {
		query["Cno"] = request.Cno
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.LoginStatus) {
		query["LoginStatus"] = request.LoginStatus
	}

	if !dara.IsNil(request.PauseDescription) {
		query["PauseDescription"] = request.PauseDescription
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudAgentLogin"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudAgentLoginResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Publish the agent online through this interface.
//
// @param request - CloudAgentLoginRequest
//
// @return CloudAgentLoginResponse
func (client *Client) CloudAgentLogin(request *CloudAgentLoginRequest) (_result *CloudAgentLoginResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudAgentLoginResponse{}
	_body, _err := client.CloudAgentLoginWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unpublish an agent.
//
// @param request - CloudAgentLogoutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudAgentLogoutResponse
func (client *Client) CloudAgentLogoutWithOptions(request *CloudAgentLogoutRequest, runtime *dara.RuntimeOptions) (_result *CloudAgentLogoutResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cno) {
		query["Cno"] = request.Cno
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.IgnoreOffline) {
		query["IgnoreOffline"] = request.IgnoreOffline
	}

	if !dara.IsNil(request.IsKickout) {
		query["IsKickout"] = request.IsKickout
	}

	if !dara.IsNil(request.RemoveBinding) {
		query["RemoveBinding"] = request.RemoveBinding
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudAgentLogout"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudAgentLogoutResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unpublish an agent.
//
// @param request - CloudAgentLogoutRequest
//
// @return CloudAgentLogoutResponse
func (client *Client) CloudAgentLogout(request *CloudAgentLogoutRequest) (_result *CloudAgentLogoutResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudAgentLogoutResponse{}
	_body, _err := client.CloudAgentLogoutWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Real-time statistics on agent call data.
//
// @param request - CloudAgentMonitorStatisticsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudAgentMonitorStatisticsResponse
func (client *Client) CloudAgentMonitorStatisticsWithOptions(request *CloudAgentMonitorStatisticsRequest, runtime *dara.RuntimeOptions) (_result *CloudAgentMonitorStatisticsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cnos) {
		query["Cnos"] = request.Cnos
	}

	if !dara.IsNil(request.EndHour) {
		query["EndHour"] = request.EndHour
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Fields) {
		query["Fields"] = request.Fields
	}

	if !dara.IsNil(request.Gno) {
		query["Gno"] = request.Gno
	}

	if !dara.IsNil(request.IsNeedQueueName) {
		query["IsNeedQueueName"] = request.IsNeedQueueName
	}

	if !dara.IsNil(request.IsUseGno) {
		query["IsUseGno"] = request.IsUseGno
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.Offset) {
		query["Offset"] = request.Offset
	}

	if !dara.IsNil(request.StartHour) {
		query["StartHour"] = request.StartHour
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudAgentMonitorStatistics"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudAgentMonitorStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Real-time statistics on agent call data.
//
// @param request - CloudAgentMonitorStatisticsRequest
//
// @return CloudAgentMonitorStatisticsResponse
func (client *Client) CloudAgentMonitorStatistics(request *CloudAgentMonitorStatisticsRequest) (_result *CloudAgentMonitorStatisticsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudAgentMonitorStatisticsResponse{}
	_body, _err := client.CloudAgentMonitorStatisticsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Use this interface to set agent associated data.
//
// @param request - CloudAgentSetUserDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudAgentSetUserDataResponse
func (client *Client) CloudAgentSetUserDataWithOptions(request *CloudAgentSetUserDataRequest, runtime *dara.RuntimeOptions) (_result *CloudAgentSetUserDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cno) {
		query["Cno"] = request.Cno
	}

	if !dara.IsNil(request.Direction) {
		query["Direction"] = request.Direction
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.UserData) {
		query["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudAgentSetUserData"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudAgentSetUserDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Use this interface to set agent associated data.
//
// @param request - CloudAgentSetUserDataRequest
//
// @return CloudAgentSetUserDataResponse
func (client *Client) CloudAgentSetUserData(request *CloudAgentSetUserDataRequest) (_result *CloudAgentSetUserDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudAgentSetUserDataResponse{}
	_body, _err := client.CloudAgentSetUserDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// The agent hangs up.
//
// @param request - CloudAgentUnlinkRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudAgentUnlinkResponse
func (client *Client) CloudAgentUnlinkWithOptions(request *CloudAgentUnlinkRequest, runtime *dara.RuntimeOptions) (_result *CloudAgentUnlinkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cno) {
		query["Cno"] = request.Cno
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.RequestUniqueId) {
		query["RequestUniqueId"] = request.RequestUniqueId
	}

	if !dara.IsNil(request.Side) {
		query["Side"] = request.Side
	}

	if !dara.IsNil(request.UniqueId) {
		query["UniqueId"] = request.UniqueId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudAgentUnlink"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudAgentUnlinkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// The agent hangs up.
//
// @param request - CloudAgentUnlinkRequest
//
// @return CloudAgentUnlinkResponse
func (client *Client) CloudAgentUnlink(request *CloudAgentUnlinkRequest) (_result *CloudAgentUnlinkResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudAgentUnlinkResponse{}
	_body, _err := client.CloudAgentUnlinkWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Gets agent workload report data.
//
// @param request - CloudAgentWorkloadReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudAgentWorkloadReportResponse
func (client *Client) CloudAgentWorkloadReportWithOptions(request *CloudAgentWorkloadReportRequest, runtime *dara.RuntimeOptions) (_result *CloudAgentWorkloadReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cnos) {
		query["Cnos"] = request.Cnos
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Gnos) {
		query["Gnos"] = request.Gnos
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Start) {
		query["Start"] = request.Start
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StatisticMethod) {
		query["StatisticMethod"] = request.StatisticMethod
	}

	if !dara.IsNil(request.TimeRangeType) {
		query["TimeRangeType"] = request.TimeRangeType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudAgentWorkloadReport"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudAgentWorkloadReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets agent workload report data.
//
// @param request - CloudAgentWorkloadReportRequest
//
// @return CloudAgentWorkloadReportResponse
func (client *Client) CloudAgentWorkloadReport(request *CloudAgentWorkloadReportRequest) (_result *CloudAgentWorkloadReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudAgentWorkloadReportResponse{}
	_body, _err := client.CloudAgentWorkloadReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Assigns an agent to an outbound group.
//
// @param request - CloudAssignAgentGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudAssignAgentGroupResponse
func (client *Client) CloudAssignAgentGroupWithOptions(request *CloudAssignAgentGroupRequest, runtime *dara.RuntimeOptions) (_result *CloudAssignAgentGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cnos) {
		query["Cnos"] = request.Cnos
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Gno) {
		query["Gno"] = request.Gno
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudAssignAgentGroup"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudAssignAgentGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Assigns an agent to an outbound group.
//
// @param request - CloudAssignAgentGroupRequest
//
// @return CloudAssignAgentGroupResponse
func (client *Client) CloudAssignAgentGroup(request *CloudAssignAgentGroupRequest) (_result *CloudAssignAgentGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudAssignAgentGroupResponse{}
	_body, _err := client.CloudAssignAgentGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Batch add agents. You can create up to 100 agents in a single batch.
//
// @param request - CloudBatchCreateAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudBatchCreateAgentResponse
func (client *Client) CloudBatchCreateAgentWithOptions(request *CloudBatchCreateAgentRequest, runtime *dara.RuntimeOptions) (_result *CloudBatchCreateAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Active) {
		query["Active"] = request.Active
	}

	if !dara.IsNil(request.AreaCode) {
		query["AreaCode"] = request.AreaCode
	}

	if !dara.IsNil(request.CallPower) {
		query["CallPower"] = request.CallPower
	}

	if !dara.IsNil(request.Cno) {
		query["Cno"] = request.Cno
	}

	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.EndCno) {
		query["EndCno"] = request.EndCno
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.IbRecord) {
		query["IbRecord"] = request.IbRecord
	}

	if !dara.IsNil(request.IsAsr) {
		query["IsAsr"] = request.IsAsr
	}

	if !dara.IsNil(request.IsOb) {
		query["IsOb"] = request.IsOb
	}

	if !dara.IsNil(request.IsQualityCheck) {
		query["IsQualityCheck"] = request.IsQualityCheck
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ObClid) {
		query["ObClid"] = request.ObClid
	}

	if !dara.IsNil(request.ObClidProperty) {
		query["ObClidProperty"] = request.ObClidProperty
	}

	if !dara.IsNil(request.ObClidType) {
		query["ObClidType"] = request.ObClidType
	}

	if !dara.IsNil(request.ObRecord) {
		query["ObRecord"] = request.ObRecord
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Power) {
		query["Power"] = request.Power
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SkillIds) {
		query["SkillIds"] = request.SkillIds
	}

	if !dara.IsNil(request.SkillLevels) {
		query["SkillLevels"] = request.SkillLevels
	}

	if !dara.IsNil(request.WebrtcUrlType) {
		query["WebrtcUrlType"] = request.WebrtcUrlType
	}

	if !dara.IsNil(request.Wrapup) {
		query["Wrapup"] = request.Wrapup
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudBatchCreateAgent"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudBatchCreateAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Batch add agents. You can create up to 100 agents in a single batch.
//
// @param request - CloudBatchCreateAgentRequest
//
// @return CloudBatchCreateAgentResponse
func (client *Client) CloudBatchCreateAgent(request *CloudBatchCreateAgentRequest) (_result *CloudBatchCreateAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudBatchCreateAgentResponse{}
	_body, _err := client.CloudBatchCreateAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the real-time status of agents in batches based on their job numbers.
//
// @param request - CloudBatchGetAgentStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudBatchGetAgentStatusResponse
func (client *Client) CloudBatchGetAgentStatusWithOptions(request *CloudBatchGetAgentStatusRequest, runtime *dara.RuntimeOptions) (_result *CloudBatchGetAgentStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cnos) {
		query["Cnos"] = request.Cnos
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudBatchGetAgentStatus"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudBatchGetAgentStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the real-time status of agents in batches based on their job numbers.
//
// @param request - CloudBatchGetAgentStatusRequest
//
// @return CloudBatchGetAgentStatusResponse
func (client *Client) CloudBatchGetAgentStatus(request *CloudBatchGetAgentStatusRequest) (_result *CloudBatchGetAgentStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudBatchGetAgentStatusResponse{}
	_body, _err := client.CloudBatchGetAgentStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Batch update the basic info of agents, excluding the update of skill info attached to agents.
//
// @param request - CloudBatchUpdateAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudBatchUpdateAgentResponse
func (client *Client) CloudBatchUpdateAgentWithOptions(request *CloudBatchUpdateAgentRequest, runtime *dara.RuntimeOptions) (_result *CloudBatchUpdateAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Active) {
		query["Active"] = request.Active
	}

	if !dara.IsNil(request.AgentType) {
		query["AgentType"] = request.AgentType
	}

	if !dara.IsNil(request.AreaCode) {
		query["AreaCode"] = request.AreaCode
	}

	if !dara.IsNil(request.CallPower) {
		query["CallPower"] = request.CallPower
	}

	if !dara.IsNil(request.Cnos) {
		query["Cnos"] = request.Cnos
	}

	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.IbRecord) {
		query["IbRecord"] = request.IbRecord
	}

	if !dara.IsNil(request.IsAsr) {
		query["IsAsr"] = request.IsAsr
	}

	if !dara.IsNil(request.IsOb) {
		query["IsOb"] = request.IsOb
	}

	if !dara.IsNil(request.IsObRemember) {
		query["IsObRemember"] = request.IsObRemember
	}

	if !dara.IsNil(request.IsQualityCheck) {
		query["IsQualityCheck"] = request.IsQualityCheck
	}

	if !dara.IsNil(request.IsRandom) {
		query["IsRandom"] = request.IsRandom
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ObClid) {
		query["ObClid"] = request.ObClid
	}

	if !dara.IsNil(request.ObClidProperty) {
		query["ObClidProperty"] = request.ObClidProperty
	}

	if !dara.IsNil(request.ObClidType) {
		query["ObClidType"] = request.ObClidType
	}

	if !dara.IsNil(request.ObRecord) {
		query["ObRecord"] = request.ObRecord
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PermitObPreviewTime) {
		query["PermitObPreviewTime"] = request.PermitObPreviewTime
	}

	if !dara.IsNil(request.Power) {
		query["Power"] = request.Power
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.WebrtcUrlType) {
		query["WebrtcUrlType"] = request.WebrtcUrlType
	}

	if !dara.IsNil(request.Wrapup) {
		query["Wrapup"] = request.Wrapup
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudBatchUpdateAgent"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudBatchUpdateAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Batch update the basic info of agents, excluding the update of skill info attached to agents.
//
// @param request - CloudBatchUpdateAgentRequest
//
// @return CloudBatchUpdateAgentResponse
func (client *Client) CloudBatchUpdateAgent(request *CloudBatchUpdateAgentRequest) (_result *CloudBatchUpdateAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudBatchUpdateAgentResponse{}
	_body, _err := client.CloudBatchUpdateAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an agent.
//
// @param request - CloudCreateAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudCreateAgentResponse
func (client *Client) CloudCreateAgentWithOptions(request *CloudCreateAgentRequest, runtime *dara.RuntimeOptions) (_result *CloudCreateAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Active) {
		query["Active"] = request.Active
	}

	if !dara.IsNil(request.AreaCode) {
		query["AreaCode"] = request.AreaCode
	}

	if !dara.IsNil(request.CallPower) {
		query["CallPower"] = request.CallPower
	}

	if !dara.IsNil(request.Cno) {
		query["Cno"] = request.Cno
	}

	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.IbRecord) {
		query["IbRecord"] = request.IbRecord
	}

	if !dara.IsNil(request.IsAsr) {
		query["IsAsr"] = request.IsAsr
	}

	if !dara.IsNil(request.IsOb) {
		query["IsOb"] = request.IsOb
	}

	if !dara.IsNil(request.IsObRemember) {
		query["IsObRemember"] = request.IsObRemember
	}

	if !dara.IsNil(request.IsQualityCheck) {
		query["IsQualityCheck"] = request.IsQualityCheck
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ObClid) {
		query["ObClid"] = request.ObClid
	}

	if !dara.IsNil(request.ObClidProperty) {
		query["ObClidProperty"] = request.ObClidProperty
	}

	if !dara.IsNil(request.ObClidType) {
		query["ObClidType"] = request.ObClidType
	}

	if !dara.IsNil(request.ObRecord) {
		query["ObRecord"] = request.ObRecord
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PermitObPreviewTime) {
		query["PermitObPreviewTime"] = request.PermitObPreviewTime
	}

	if !dara.IsNil(request.Power) {
		query["Power"] = request.Power
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SkillIds) {
		query["SkillIds"] = request.SkillIds
	}

	if !dara.IsNil(request.SkillLevels) {
		query["SkillLevels"] = request.SkillLevels
	}

	if !dara.IsNil(request.WebrtcUrlType) {
		query["WebrtcUrlType"] = request.WebrtcUrlType
	}

	if !dara.IsNil(request.Wrapup) {
		query["Wrapup"] = request.Wrapup
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudCreateAgent"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudCreateAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an agent.
//
// @param request - CloudCreateAgentRequest
//
// @return CloudCreateAgentResponse
func (client *Client) CloudCreateAgent(request *CloudCreateAgentRequest) (_result *CloudCreateAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudCreateAgentResponse{}
	_body, _err := client.CloudCreateAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call this operation to add an outbound group.
//
// @param request - CloudCreateAgentGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudCreateAgentGroupResponse
func (client *Client) CloudCreateAgentGroupWithOptions(request *CloudCreateAgentGroupRequest, runtime *dara.RuntimeOptions) (_result *CloudCreateAgentGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Gno) {
		query["Gno"] = request.Gno
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudCreateAgentGroup"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudCreateAgentGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call this operation to add an outbound group.
//
// @param request - CloudCreateAgentGroupRequest
//
// @return CloudCreateAgentGroupResponse
func (client *Client) CloudCreateAgentGroup(request *CloudCreateAgentGroupRequest) (_result *CloudCreateAgentGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudCreateAgentGroupResponse{}
	_body, _err := client.CloudCreateAgentGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Upload a recording file and create an ASR transform job.
//
// @param request - CloudCreateAsrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudCreateAsrResponse
func (client *Client) CloudCreateAsrWithOptions(request *CloudCreateAsrRequest, runtime *dara.RuntimeOptions) (_result *CloudCreateAsrResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallType) {
		query["CallType"] = request.CallType
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.MainUniqueId) {
		query["MainUniqueId"] = request.MainUniqueId
	}

	if !dara.IsNil(request.RecordFile) {
		query["RecordFile"] = request.RecordFile
	}

	if !dara.IsNil(request.RecordSide) {
		query["RecordSide"] = request.RecordSide
	}

	if !dara.IsNil(request.RecordType) {
		query["RecordType"] = request.RecordType
	}

	if !dara.IsNil(request.SupportMp3) {
		query["SupportMp3"] = request.SupportMp3
	}

	if !dara.IsNil(request.UniqueId) {
		query["UniqueId"] = request.UniqueId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudCreateAsr"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudCreateAsrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upload a recording file and create an ASR transform job.
//
// @param request - CloudCreateAsrRequest
//
// @return CloudCreateAsrResponse
func (client *Client) CloudCreateAsr(request *CloudCreateAsrRequest) (_result *CloudCreateAsrResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudCreateAsrResponse{}
	_body, _err := client.CloudCreateAsrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a time condition configuration.
//
// @param request - CloudCreateEnterpriseTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudCreateEnterpriseTimeResponse
func (client *Client) CloudCreateEnterpriseTimeWithOptions(request *CloudCreateEnterpriseTimeRequest, runtime *dara.RuntimeOptions) (_result *CloudCreateEnterpriseTimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DayOfWeek) {
		query["DayOfWeek"] = request.DayOfWeek
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.FromDay) {
		query["FromDay"] = request.FromDay
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TimeType) {
		query["TimeType"] = request.TimeType
	}

	if !dara.IsNil(request.ToDay) {
		query["ToDay"] = request.ToDay
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudCreateEnterpriseTime"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudCreateEnterpriseTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a time condition configuration.
//
// @param request - CloudCreateEnterpriseTimeRequest
//
// @return CloudCreateEnterpriseTimeResponse
func (client *Client) CloudCreateEnterpriseTime(request *CloudCreateEnterpriseTimeRequest) (_result *CloudCreateEnterpriseTimeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudCreateEnterpriseTimeResponse{}
	_body, _err := client.CloudCreateEnterpriseTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Add an extension through this interface.
//
// @param request - CloudCreateExtenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudCreateExtenResponse
func (client *Client) CloudCreateExtenWithOptions(request *CloudCreateExtenRequest, runtime *dara.RuntimeOptions) (_result *CloudCreateExtenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Allow) {
		query["Allow"] = request.Allow
	}

	if !dara.IsNil(request.AreaCode) {
		query["AreaCode"] = request.AreaCode
	}

	if !dara.IsNil(request.CallPower) {
		query["CallPower"] = request.CallPower
	}

	if !dara.IsNil(request.Denoise) {
		query["Denoise"] = request.Denoise
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Exten) {
		query["Exten"] = request.Exten
	}

	if !dara.IsNil(request.IadName) {
		query["IadName"] = request.IadName
	}

	if !dara.IsNil(request.IbRecord) {
		query["IbRecord"] = request.IbRecord
	}

	if !dara.IsNil(request.IsDirect) {
		query["IsDirect"] = request.IsDirect
	}

	if !dara.IsNil(request.IsOb) {
		query["IsOb"] = request.IsOb
	}

	if !dara.IsNil(request.JitterBuffer) {
		query["JitterBuffer"] = request.JitterBuffer
	}

	if !dara.IsNil(request.ObRecord) {
		query["ObRecord"] = request.ObRecord
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Password) {
		query["Password"] = request.Password
	}

	if !dara.IsNil(request.Property) {
		query["Property"] = request.Property
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudCreateExten"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudCreateExtenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Add an extension through this interface.
//
// @param request - CloudCreateExtenRequest
//
// @return CloudCreateExtenResponse
func (client *Client) CloudCreateExten(request *CloudCreateExtenRequest) (_result *CloudCreateExtenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudCreateExtenResponse{}
	_body, _err := client.CloudCreateExtenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Add a queue through this interface.
//
// @param tmpReq - CloudCreateQueueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudCreateQueueResponse
func (client *Client) CloudCreateQueueWithOptions(tmpReq *CloudCreateQueueRequest, runtime *dara.RuntimeOptions) (_result *CloudCreateQueueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CloudCreateQueueShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Queue) {
		request.QueueShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Queue, dara.String("Queue"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.QueueSkills) {
		request.QueueSkillsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.QueueSkills, dara.String("QueueSkills"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.QueueShrink) {
		query["Queue"] = request.QueueShrink
	}

	if !dara.IsNil(request.QueueSkillsShrink) {
		query["QueueSkills"] = request.QueueSkillsShrink
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudCreateQueue"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudCreateQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Add a queue through this interface.
//
// @param request - CloudCreateQueueRequest
//
// @return CloudCreateQueueResponse
func (client *Client) CloudCreateQueue(request *CloudCreateQueueRequest) (_result *CloudCreateQueueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudCreateQueueResponse{}
	_body, _err := client.CloudCreateQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Add skills to a queue.
//
// @param request - CloudCreateQueueSkillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudCreateQueueSkillResponse
func (client *Client) CloudCreateQueueSkillWithOptions(request *CloudCreateQueueSkillRequest, runtime *dara.RuntimeOptions) (_result *CloudCreateQueueSkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Qno) {
		query["Qno"] = request.Qno
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SkillId) {
		query["SkillId"] = request.SkillId
	}

	if !dara.IsNil(request.SkillLevel) {
		query["SkillLevel"] = request.SkillLevel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudCreateQueueSkill"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudCreateQueueSkillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Add skills to a queue.
//
// @param request - CloudCreateQueueSkillRequest
//
// @return CloudCreateQueueSkillResponse
func (client *Client) CloudCreateQueueSkill(request *CloudCreateQueueSkillRequest) (_result *CloudCreateQueueSkillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudCreateQueueSkillResponse{}
	_body, _err := client.CloudCreateQueueSkillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds a skill by calling this interface.
//
// @param request - CloudCreateSkillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudCreateSkillResponse
func (client *Client) CloudCreateSkillWithOptions(request *CloudCreateSkillRequest, runtime *dara.RuntimeOptions) (_result *CloudCreateSkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudCreateSkill"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudCreateSkillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a skill by calling this interface.
//
// @param request - CloudCreateSkillRequest
//
// @return CloudCreateSkillResponse
func (client *Client) CloudCreateSkill(request *CloudCreateSkillRequest) (_result *CloudCreateSkillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudCreateSkillResponse{}
	_body, _err := client.CloudCreateSkillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an outbound call job.
//
// @param request - CloudCreateTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudCreateTaskResponse
func (client *Client) CloudCreateTaskWithOptions(request *CloudCreateTaskRequest, runtime *dara.RuntimeOptions) (_result *CloudCreateTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentGroup) {
		query["AgentGroup"] = request.AgentGroup
	}

	if !dara.IsNil(request.AgentTimeout) {
		query["AgentTimeout"] = request.AgentTimeout
	}

	if !dara.IsNil(request.AnswerRate) {
		query["AnswerRate"] = request.AnswerRate
	}

	if !dara.IsNil(request.AutoComplete) {
		query["AutoComplete"] = request.AutoComplete
	}

	if !dara.IsNil(request.AutoDelete) {
		query["AutoDelete"] = request.AutoDelete
	}

	if !dara.IsNil(request.AutoStart) {
		query["AutoStart"] = request.AutoStart
	}

	if !dara.IsNil(request.AutoStartDay) {
		query["AutoStartDay"] = request.AutoStartDay
	}

	if !dara.IsNil(request.AutoStartTime) {
		query["AutoStartTime"] = request.AutoStartTime
	}

	if !dara.IsNil(request.AutoStop) {
		query["AutoStop"] = request.AutoStop
	}

	if !dara.IsNil(request.AutoStopDay) {
		query["AutoStopDay"] = request.AutoStopDay
	}

	if !dara.IsNil(request.AutoStopTime) {
		query["AutoStopTime"] = request.AutoStopTime
	}

	if !dara.IsNil(request.AutoTaskType) {
		query["AutoTaskType"] = request.AutoTaskType
	}

	if !dara.IsNil(request.AutoTriggerTimeStrategy) {
		query["AutoTriggerTimeStrategy"] = request.AutoTriggerTimeStrategy
	}

	if !dara.IsNil(request.CallGroupType) {
		query["CallGroupType"] = request.CallGroupType
	}

	if !dara.IsNil(request.CallLimitStrategy) {
		query["CallLimitStrategy"] = request.CallLimitStrategy
	}

	if !dara.IsNil(request.CallPriorityStrategy) {
		query["CallPriorityStrategy"] = request.CallPriorityStrategy
	}

	if !dara.IsNil(request.CallRouteStrategy) {
		query["CallRouteStrategy"] = request.CallRouteStrategy
	}

	if !dara.IsNil(request.CallStrategy) {
		query["CallStrategy"] = request.CallStrategy
	}

	if !dara.IsNil(request.CallVariables) {
		query["CallVariables"] = request.CallVariables
	}

	if !dara.IsNil(request.ClidProperty) {
		query["ClidProperty"] = request.ClidProperty
	}

	if !dara.IsNil(request.Cnos) {
		query["Cnos"] = request.Cnos
	}

	if !dara.IsNil(request.Concurrency) {
		query["Concurrency"] = request.Concurrency
	}

	if !dara.IsNil(request.CustomerClidType) {
		query["CustomerClidType"] = request.CustomerClidType
	}

	if !dara.IsNil(request.CustomerClidWeight) {
		query["CustomerClidWeight"] = request.CustomerClidWeight
	}

	if !dara.IsNil(request.CustomerClidWeightFlag) {
		query["CustomerClidWeightFlag"] = request.CustomerClidWeightFlag
	}

	if !dara.IsNil(request.CustomerClids) {
		query["CustomerClids"] = request.CustomerClids
	}

	if !dara.IsNil(request.CustomerClidsCategory) {
		query["CustomerClidsCategory"] = request.CustomerClidsCategory
	}

	if !dara.IsNil(request.CustomerClidsGroup) {
		query["CustomerClidsGroup"] = request.CustomerClidsGroup
	}

	if !dara.IsNil(request.CustomerMoh) {
		query["CustomerMoh"] = request.CustomerMoh
	}

	if !dara.IsNil(request.CustomerTimeout) {
		query["CustomerTimeout"] = request.CustomerTimeout
	}

	if !dara.IsNil(request.CustomerVoice) {
		query["CustomerVoice"] = request.CustomerVoice
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.ForceEndFlag) {
		query["ForceEndFlag"] = request.ForceEndFlag
	}

	if !dara.IsNil(request.IsRewarm) {
		query["IsRewarm"] = request.IsRewarm
	}

	if !dara.IsNil(request.IvrId) {
		query["IvrId"] = request.IvrId
	}

	if !dara.IsNil(request.IvrName) {
		query["IvrName"] = request.IvrName
	}

	if !dara.IsNil(request.MaxWaitTime) {
		query["MaxWaitTime"] = request.MaxWaitTime
	}

	if !dara.IsNil(request.MinAvailableAgentCount) {
		query["MinAvailableAgentCount"] = request.MinAvailableAgentCount
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PredictAdjust) {
		query["PredictAdjust"] = request.PredictAdjust
	}

	if !dara.IsNil(request.Quotiety) {
		query["Quotiety"] = request.Quotiety
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RetryStrategy) {
		query["RetryStrategy"] = request.RetryStrategy
	}

	if !dara.IsNil(request.RetryStrategyOnlyToday) {
		query["RetryStrategyOnlyToday"] = request.RetryStrategyOnlyToday
	}

	if !dara.IsNil(request.RetryStrategyTimeType) {
		query["RetryStrategyTimeType"] = request.RetryStrategyTimeType
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TimeStrategy) {
		query["TimeStrategy"] = request.TimeStrategy
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.UserFields) {
		query["UserFields"] = request.UserFields
	}

	if !dara.IsNil(request.WarmUpDuration) {
		query["WarmUpDuration"] = request.WarmUpDuration
	}

	if !dara.IsNil(request.Wrapup) {
		query["Wrapup"] = request.Wrapup
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudCreateTask"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudCreateTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an outbound call job.
//
// @param request - CloudCreateTaskRequest
//
// @return CloudCreateTaskResponse
func (client *Client) CloudCreateTask(request *CloudCreateTaskRequest) (_result *CloudCreateTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudCreateTaskResponse{}
	_body, _err := client.CloudCreateTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete an agent based on the agent number.
//
// @param request - CloudDeleteAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudDeleteAgentResponse
func (client *Client) CloudDeleteAgentWithOptions(request *CloudDeleteAgentRequest, runtime *dara.RuntimeOptions) (_result *CloudDeleteAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cno) {
		query["Cno"] = request.Cno
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudDeleteAgent"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudDeleteAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete an agent based on the agent number.
//
// @param request - CloudDeleteAgentRequest
//
// @return CloudDeleteAgentResponse
func (client *Client) CloudDeleteAgent(request *CloudDeleteAgentRequest) (_result *CloudDeleteAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudDeleteAgentResponse{}
	_body, _err := client.CloudDeleteAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete an outbound call group.
//
// @param request - CloudDeleteAgentGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudDeleteAgentGroupResponse
func (client *Client) CloudDeleteAgentGroupWithOptions(request *CloudDeleteAgentGroupRequest, runtime *dara.RuntimeOptions) (_result *CloudDeleteAgentGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Gno) {
		query["Gno"] = request.Gno
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudDeleteAgentGroup"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudDeleteAgentGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete an outbound call group.
//
// @param request - CloudDeleteAgentGroupRequest
//
// @return CloudDeleteAgentGroupResponse
func (client *Client) CloudDeleteAgentGroup(request *CloudDeleteAgentGroupRequest) (_result *CloudDeleteAgentGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudDeleteAgentGroupResponse{}
	_body, _err := client.CloudDeleteAgentGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete the phone under the agent.
//
// @param request - CloudDeleteAgentTelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudDeleteAgentTelResponse
func (client *Client) CloudDeleteAgentTelWithOptions(request *CloudDeleteAgentTelRequest, runtime *dara.RuntimeOptions) (_result *CloudDeleteAgentTelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cno) {
		query["Cno"] = request.Cno
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tel) {
		query["Tel"] = request.Tel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudDeleteAgentTel"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudDeleteAgentTelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete the phone under the agent.
//
// @param request - CloudDeleteAgentTelRequest
//
// @return CloudDeleteAgentTelResponse
func (client *Client) CloudDeleteAgentTel(request *CloudDeleteAgentTelRequest) (_result *CloudDeleteAgentTelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudDeleteAgentTelResponse{}
	_body, _err := client.CloudDeleteAgentTelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete a time condition setting.
//
// @param request - CloudDeleteEnterpriseTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudDeleteEnterpriseTimeResponse
func (client *Client) CloudDeleteEnterpriseTimeWithOptions(request *CloudDeleteEnterpriseTimeRequest, runtime *dara.RuntimeOptions) (_result *CloudDeleteEnterpriseTimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudDeleteEnterpriseTime"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudDeleteEnterpriseTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete a time condition setting.
//
// @param request - CloudDeleteEnterpriseTimeRequest
//
// @return CloudDeleteEnterpriseTimeResponse
func (client *Client) CloudDeleteEnterpriseTime(request *CloudDeleteEnterpriseTimeRequest) (_result *CloudDeleteEnterpriseTimeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudDeleteEnterpriseTimeResponse{}
	_body, _err := client.CloudDeleteEnterpriseTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an extension based on the extension number.
//
// @param request - CloudDeleteExtenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudDeleteExtenResponse
func (client *Client) CloudDeleteExtenWithOptions(request *CloudDeleteExtenRequest, runtime *dara.RuntimeOptions) (_result *CloudDeleteExtenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Exten) {
		query["Exten"] = request.Exten
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudDeleteExten"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudDeleteExtenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an extension based on the extension number.
//
// @param request - CloudDeleteExtenRequest
//
// @return CloudDeleteExtenResponse
func (client *Client) CloudDeleteExten(request *CloudDeleteExtenRequest) (_result *CloudDeleteExtenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudDeleteExtenResponse{}
	_body, _err := client.CloudDeleteExtenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete a queue.
//
// @param request - CloudDeleteQueueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudDeleteQueueResponse
func (client *Client) CloudDeleteQueueWithOptions(request *CloudDeleteQueueRequest, runtime *dara.RuntimeOptions) (_result *CloudDeleteQueueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Qno) {
		query["Qno"] = request.Qno
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudDeleteQueue"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudDeleteQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete a queue.
//
// @param request - CloudDeleteQueueRequest
//
// @return CloudDeleteQueueResponse
func (client *Client) CloudDeleteQueue(request *CloudDeleteQueueRequest) (_result *CloudDeleteQueueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudDeleteQueueResponse{}
	_body, _err := client.CloudDeleteQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queue delete skills.
//
// @param request - CloudDeleteQueueSkillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudDeleteQueueSkillResponse
func (client *Client) CloudDeleteQueueSkillWithOptions(request *CloudDeleteQueueSkillRequest, runtime *dara.RuntimeOptions) (_result *CloudDeleteQueueSkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Qno) {
		query["Qno"] = request.Qno
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SkillId) {
		query["SkillId"] = request.SkillId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudDeleteQueueSkill"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudDeleteQueueSkillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queue delete skills.
//
// @param request - CloudDeleteQueueSkillRequest
//
// @return CloudDeleteQueueSkillResponse
func (client *Client) CloudDeleteQueueSkill(request *CloudDeleteQueueSkillRequest) (_result *CloudDeleteQueueSkillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudDeleteQueueSkillResponse{}
	_body, _err := client.CloudDeleteQueueSkillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a skill.
//
// @param request - CloudDeleteSkillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudDeleteSkillResponse
func (client *Client) CloudDeleteSkillWithOptions(request *CloudDeleteSkillRequest, runtime *dara.RuntimeOptions) (_result *CloudDeleteSkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudDeleteSkill"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudDeleteSkillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a skill.
//
// @param request - CloudDeleteSkillRequest
//
// @return CloudDeleteSkillResponse
func (client *Client) CloudDeleteSkill(request *CloudDeleteSkillRequest) (_result *CloudDeleteSkillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudDeleteSkillResponse{}
	_body, _err := client.CloudDeleteSkillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Supports deleting predictive outbound call and automatic outbound call jobs. Only jobs in the initial or completed status can be deleted. When a job is deleted, the associated numbers are also deleted.
//
// @param request - CloudDeleteTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudDeleteTaskResponse
func (client *Client) CloudDeleteTaskWithOptions(request *CloudDeleteTaskRequest, runtime *dara.RuntimeOptions) (_result *CloudDeleteTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudDeleteTask"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudDeleteTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Supports deleting predictive outbound call and automatic outbound call jobs. Only jobs in the initial or completed status can be deleted. When a job is deleted, the associated numbers are also deleted.
//
// @param request - CloudDeleteTaskRequest
//
// @return CloudDeleteTaskResponse
func (client *Client) CloudDeleteTask(request *CloudDeleteTaskRequest) (_result *CloudDeleteTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudDeleteTaskResponse{}
	_body, _err := client.CloudDeleteTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes the number of a call job through this interface.
//
// @param request - CloudDeleteTaskTelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudDeleteTaskTelResponse
func (client *Client) CloudDeleteTaskTelWithOptions(request *CloudDeleteTaskTelRequest, runtime *dara.RuntimeOptions) (_result *CloudDeleteTaskTelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	if !dara.IsNil(request.FileName) {
		query["FileName"] = request.FileName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.Tels) {
		query["Tels"] = request.Tels
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudDeleteTaskTel"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudDeleteTaskTelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the number of a call job through this interface.
//
// @param request - CloudDeleteTaskTelRequest
//
// @return CloudDeleteTaskTelResponse
func (client *Client) CloudDeleteTaskTel(request *CloudDeleteTaskTelRequest) (_result *CloudDeleteTaskTelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudDeleteTaskTelResponse{}
	_body, _err := client.CloudDeleteTaskTelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of an agent.
//
// @param request - CloudGetAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudGetAgentResponse
func (client *Client) CloudGetAgentWithOptions(request *CloudGetAgentRequest, runtime *dara.RuntimeOptions) (_result *CloudGetAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cno) {
		query["Cno"] = request.Cno
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudGetAgent"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudGetAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of an agent.
//
// @param request - CloudGetAgentRequest
//
// @return CloudGetAgentResponse
func (client *Client) CloudGetAgent(request *CloudGetAgentRequest) (_result *CloudGetAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudGetAgentResponse{}
	_body, _err := client.CloudGetAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the real-time status info of agents.
//
// @param request - CloudGetAgentStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudGetAgentStatusResponse
func (client *Client) CloudGetAgentStatusWithOptions(request *CloudGetAgentStatusRequest, runtime *dara.RuntimeOptions) (_result *CloudGetAgentStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cno) {
		query["Cno"] = request.Cno
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudGetAgentStatus"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudGetAgentStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the real-time status info of agents.
//
// @param request - CloudGetAgentStatusRequest
//
// @return CloudGetAgentStatusResponse
func (client *Client) CloudGetAgentStatus(request *CloudGetAgentStatusRequest) (_result *CloudGetAgentStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudGetAgentStatusResponse{}
	_body, _err := client.CloudGetAgentStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query phone number attribution.
//
// @param request - CloudGetAreaCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudGetAreaCodeResponse
func (client *Client) CloudGetAreaCodeWithOptions(request *CloudGetAreaCodeRequest, runtime *dara.RuntimeOptions) (_result *CloudGetAreaCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Tel) {
		query["Tel"] = request.Tel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudGetAreaCode"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudGetAreaCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query phone number attribution.
//
// @param request - CloudGetAreaCodeRequest
//
// @return CloudGetAreaCodeResponse
func (client *Client) CloudGetAreaCode(request *CloudGetAreaCodeRequest) (_result *CloudGetAreaCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudGetAreaCodeResponse{}
	_body, _err := client.CloudGetAreaCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query ASR job task results and obtain data.
//
// @param request - CloudGetAsrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudGetAsrResponse
func (client *Client) CloudGetAsrWithOptions(request *CloudGetAsrRequest, runtime *dara.RuntimeOptions) (_result *CloudGetAsrResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["All"] = request.All
	}

	if !dara.IsNil(request.CallType) {
		query["CallType"] = request.CallType
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.MainUniqueId) {
		query["MainUniqueId"] = request.MainUniqueId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudGetAsr"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudGetAsrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query ASR job task results and obtain data.
//
// @param request - CloudGetAsrRequest
//
// @return CloudGetAsrResponse
func (client *Client) CloudGetAsr(request *CloudGetAsrRequest) (_result *CloudGetAsrResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudGetAsrResponse{}
	_body, _err := client.CloudGetAsrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the info about a specified extension.
//
// @param request - CloudGetExtenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudGetExtenResponse
func (client *Client) CloudGetExtenWithOptions(request *CloudGetExtenRequest, runtime *dara.RuntimeOptions) (_result *CloudGetExtenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Exten) {
		query["Exten"] = request.Exten
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudGetExten"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudGetExtenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the info about a specified extension.
//
// @param request - CloudGetExtenRequest
//
// @return CloudGetExtenResponse
func (client *Client) CloudGetExten(request *CloudGetExtenRequest) (_result *CloudGetExtenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudGetExtenResponse{}
	_body, _err := client.CloudGetExtenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of outbound call records of a specified agent based on the unique phone identity.
//
// @param request - CloudGetObCdrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudGetObCdrResponse
func (client *Client) CloudGetObCdrWithOptions(request *CloudGetObCdrRequest, runtime *dara.RuntimeOptions) (_result *CloudGetObCdrResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.UniqueId) {
		query["UniqueId"] = request.UniqueId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudGetObCdr"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudGetObCdrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of outbound call records of a specified agent based on the unique phone identity.
//
// @param request - CloudGetObCdrRequest
//
// @return CloudGetObCdrResponse
func (client *Client) CloudGetObCdr(request *CloudGetObCdrRequest) (_result *CloudGetObCdrResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudGetObCdrResponse{}
	_body, _err := client.CloudGetObCdrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query queue info.
//
// @param request - CloudGetQueueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudGetQueueResponse
func (client *Client) CloudGetQueueWithOptions(request *CloudGetQueueRequest, runtime *dara.RuntimeOptions) (_result *CloudGetQueueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Qno) {
		query["Qno"] = request.Qno
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudGetQueue"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudGetQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query queue info.
//
// @param request - CloudGetQueueRequest
//
// @return CloudGetQueueResponse
func (client *Client) CloudGetQueue(request *CloudGetQueueRequest) (_result *CloudGetQueueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudGetQueueResponse{}
	_body, _err := client.CloudGetQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the audition or download URL of a recording based on the recording file name.
//
// @param request - CloudGetRecordUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudGetRecordUrlResponse
func (client *Client) CloudGetRecordUrlWithOptions(request *CloudGetRecordUrlRequest, runtime *dara.RuntimeOptions) (_result *CloudGetRecordUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallType) {
		query["CallType"] = request.CallType
	}

	if !dara.IsNil(request.Download) {
		query["Download"] = request.Download
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.RecordFile) {
		query["RecordFile"] = request.RecordFile
	}

	if !dara.IsNil(request.RecordFormat) {
		query["RecordFormat"] = request.RecordFormat
	}

	if !dara.IsNil(request.RecordSide) {
		query["RecordSide"] = request.RecordSide
	}

	if !dara.IsNil(request.RecordType) {
		query["RecordType"] = request.RecordType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudGetRecordUrl"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudGetRecordUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the audition or download URL of a recording based on the recording file name.
//
// @param request - CloudGetRecordUrlRequest
//
// @return CloudGetRecordUrlResponse
func (client *Client) CloudGetRecordUrl(request *CloudGetRecordUrlRequest) (_result *CloudGetRecordUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudGetRecordUrlResponse{}
	_body, _err := client.CloudGetRecordUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the configuration info of a single job based on the job ID.
//
// @param request - CloudGetTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudGetTaskResponse
func (client *Client) CloudGetTaskWithOptions(request *CloudGetTaskRequest, runtime *dara.RuntimeOptions) (_result *CloudGetTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudGetTask"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudGetTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the configuration info of a single job based on the job ID.
//
// @param request - CloudGetTaskRequest
//
// @return CloudGetTaskResponse
func (client *Client) CloudGetTask(request *CloudGetTaskRequest) (_result *CloudGetTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudGetTaskResponse{}
	_body, _err := client.CloudGetTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Imports phone numbers for the outbound call task.
//
// @param tmpReq - CloudImportTaskTelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudImportTaskTelResponse
func (client *Client) CloudImportTaskTelWithOptions(tmpReq *CloudImportTaskTelRequest, runtime *dara.RuntimeOptions) (_result *CloudImportTaskTelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CloudImportTaskTelShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TaskTelList) {
		request.TaskTelListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskTelList, dara.String("TaskTelList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BridgeVoicePath) {
		query["BridgeVoicePath"] = request.BridgeVoicePath
	}

	if !dara.IsNil(request.BridgeVoiceType) {
		query["BridgeVoiceType"] = request.BridgeVoiceType
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.FileId) {
		query["FileId"] = request.FileId
	}

	if !dara.IsNil(request.ImportTelAutoStart) {
		query["ImportTelAutoStart"] = request.ImportTelAutoStart
	}

	if !dara.IsNil(request.IsRepeat) {
		query["IsRepeat"] = request.IsRepeat
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskTelListShrink) {
		query["TaskTelList"] = request.TaskTelListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudImportTaskTel"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudImportTaskTelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Imports phone numbers for the outbound call task.
//
// @param request - CloudImportTaskTelRequest
//
// @return CloudImportTaskTelResponse
func (client *Client) CloudImportTaskTel(request *CloudImportTaskTelRequest) (_result *CloudImportTaskTelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudImportTaskTelResponse{}
	_body, _err := client.CloudImportTaskTelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// If the current call is at an await edge zone in voice navigation, this interface can interrupt the wait and execute to the next hop.
//
// @param request - CloudInterruptIvrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudInterruptIvrResponse
func (client *Client) CloudInterruptIvrWithOptions(request *CloudInterruptIvrRequest, runtime *dara.RuntimeOptions) (_result *CloudInterruptIvrResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CheckName) {
		query["CheckName"] = request.CheckName
	}

	if !dara.IsNil(request.CheckValue) {
		query["CheckValue"] = request.CheckValue
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.UniqueId) {
		query["UniqueId"] = request.UniqueId
	}

	if !dara.IsNil(request.UserField) {
		query["UserField"] = request.UserField
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudInterruptIvr"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudInterruptIvrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// If the current call is at an await edge zone in voice navigation, this interface can interrupt the wait and execute to the next hop.
//
// @param request - CloudInterruptIvrRequest
//
// @return CloudInterruptIvrResponse
func (client *Client) CloudInterruptIvr(request *CloudInterruptIvrRequest) (_result *CloudInterruptIvrResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudInterruptIvrResponse{}
	_body, _err := client.CloudInterruptIvrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call this operation to query outbound groups.
//
// @param request - CloudListAgentGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudListAgentGroupResponse
func (client *Client) CloudListAgentGroupWithOptions(request *CloudListAgentGroupRequest, runtime *dara.RuntimeOptions) (_result *CloudListAgentGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Gno) {
		query["Gno"] = request.Gno
	}

	if !dara.IsNil(request.GroupName) {
		query["GroupName"] = request.GroupName
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Start) {
		query["Start"] = request.Start
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudListAgentGroup"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudListAgentGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to query outbound groups.
//
// @param request - CloudListAgentGroupRequest
//
// @return CloudListAgentGroupResponse
func (client *Client) CloudListAgentGroup(request *CloudListAgentGroupRequest) (_result *CloudListAgentGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudListAgentGroupResponse{}
	_body, _err := client.CloudListAgentGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtain agent phone info by agent number.
//
// @param request - CloudListAgentTelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudListAgentTelResponse
func (client *Client) CloudListAgentTelWithOptions(request *CloudListAgentTelRequest, runtime *dara.RuntimeOptions) (_result *CloudListAgentTelResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cno) {
		query["Cno"] = request.Cno
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Tel) {
		query["Tel"] = request.Tel
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudListAgentTel"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudListAgentTelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain agent phone info by agent number.
//
// @param request - CloudListAgentTelRequest
//
// @return CloudListAgentTelResponse
func (client *Client) CloudListAgentTel(request *CloudListAgentTelRequest) (_result *CloudListAgentTelResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudListAgentTelResponse{}
	_body, _err := client.CloudListAgentTelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// You can call this operation to query the list of agents in an outbound group.
//
// @param request - CloudListAssignedAgentGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudListAssignedAgentGroupResponse
func (client *Client) CloudListAssignedAgentGroupWithOptions(request *CloudListAssignedAgentGroupRequest, runtime *dara.RuntimeOptions) (_result *CloudListAssignedAgentGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cname) {
		query["Cname"] = request.Cname
	}

	if !dara.IsNil(request.Cno) {
		query["Cno"] = request.Cno
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Gno) {
		query["Gno"] = request.Gno
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudListAssignedAgentGroup"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudListAssignedAgentGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// You can call this operation to query the list of agents in an outbound group.
//
// @param request - CloudListAssignedAgentGroupRequest
//
// @return CloudListAssignedAgentGroupResponse
func (client *Client) CloudListAssignedAgentGroup(request *CloudListAssignedAgentGroupRequest) (_result *CloudListAssignedAgentGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudListAssignedAgentGroupResponse{}
	_body, _err := client.CloudListAssignedAgentGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries user push logs.
//
// @param request - CloudListCurlLogRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudListCurlLogResponse
func (client *Client) CloudListCurlLogWithOptions(request *CloudListCurlLogRequest, runtime *dara.RuntimeOptions) (_result *CloudListCurlLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.Result) {
		query["Result"] = request.Result
	}

	if !dara.IsNil(request.Retry) {
		query["Retry"] = request.Retry
	}

	if !dara.IsNil(request.Start) {
		query["Start"] = request.Start
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	if !dara.IsNil(request.Url) {
		query["Url"] = request.Url
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudListCurlLog"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudListCurlLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries user push logs.
//
// @param request - CloudListCurlLogRequest
//
// @return CloudListCurlLogResponse
func (client *Client) CloudListCurlLog(request *CloudListCurlLogRequest) (_result *CloudListCurlLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudListCurlLogResponse{}
	_body, _err := client.CloudListCurlLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of time condition settings.
//
// @param request - CloudListEnterpriseTimeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudListEnterpriseTimeResponse
func (client *Client) CloudListEnterpriseTimeWithOptions(request *CloudListEnterpriseTimeRequest, runtime *dara.RuntimeOptions) (_result *CloudListEnterpriseTimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Priority) {
		query["Priority"] = request.Priority
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudListEnterpriseTime"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudListEnterpriseTimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of time condition settings.
//
// @param request - CloudListEnterpriseTimeRequest
//
// @return CloudListEnterpriseTimeResponse
func (client *Client) CloudListEnterpriseTime(request *CloudListEnterpriseTimeRequest) (_result *CloudListEnterpriseTimeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudListEnterpriseTimeResponse{}
	_body, _err := client.CloudListEnterpriseTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the extension info list.
//
// @param request - CloudListExtenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudListExtenResponse
func (client *Client) CloudListExtenWithOptions(request *CloudListExtenRequest, runtime *dara.RuntimeOptions) (_result *CloudListExtenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AreaCode) {
		query["AreaCode"] = request.AreaCode
	}

	if !dara.IsNil(request.CallPower) {
		query["CallPower"] = request.CallPower
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Exten) {
		query["Exten"] = request.Exten
	}

	if !dara.IsNil(request.IbRecord) {
		query["IbRecord"] = request.IbRecord
	}

	if !dara.IsNil(request.IsBind) {
		query["IsBind"] = request.IsBind
	}

	if !dara.IsNil(request.IsOb) {
		query["IsOb"] = request.IsOb
	}

	if !dara.IsNil(request.JitterBuffer) {
		query["JitterBuffer"] = request.JitterBuffer
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.ObRecord) {
		query["ObRecord"] = request.ObRecord
	}

	if !dara.IsNil(request.Offset) {
		query["Offset"] = request.Offset
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudListExten"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudListExtenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the extension info list.
//
// @param request - CloudListExtenRequest
//
// @return CloudListExtenResponse
func (client *Client) CloudListExten(request *CloudListExtenRequest) (_result *CloudListExtenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudListExtenResponse{}
	_body, _err := client.CloudListExtenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of online agents.
//
// @param request - CloudListFreeAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudListFreeAgentResponse
func (client *Client) CloudListFreeAgentWithOptions(request *CloudListFreeAgentRequest, runtime *dara.RuntimeOptions) (_result *CloudListFreeAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudListFreeAgent"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudListFreeAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of online agents.
//
// @param request - CloudListFreeAgentRequest
//
// @return CloudListFreeAgentResponse
func (client *Client) CloudListFreeAgent(request *CloudListFreeAgentRequest) (_result *CloudListFreeAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudListFreeAgentResponse{}
	_body, _err := client.CloudListFreeAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of online agent info.
//
// @param request - CloudListOnlineAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudListOnlineAgentResponse
func (client *Client) CloudListOnlineAgentWithOptions(request *CloudListOnlineAgentRequest, runtime *dara.RuntimeOptions) (_result *CloudListOnlineAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cnos) {
		query["Cnos"] = request.Cnos
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PauseDescription) {
		query["PauseDescription"] = request.PauseDescription
	}

	if !dara.IsNil(request.PauseType) {
		query["PauseType"] = request.PauseType
	}

	if !dara.IsNil(request.Qnos) {
		query["Qnos"] = request.Qnos
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.StateCode) {
		query["StateCode"] = request.StateCode
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudListOnlineAgent"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudListOnlineAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of online agent info.
//
// @param request - CloudListOnlineAgentRequest
//
// @return CloudListOnlineAgentResponse
func (client *Client) CloudListOnlineAgent(request *CloudListOnlineAgentRequest) (_result *CloudListOnlineAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudListOnlineAgentResponse{}
	_body, _err := client.CloudListOnlineAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the queue info list.
//
// @param request - CloudListQueueRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudListQueueResponse
func (client *Client) CloudListQueueWithOptions(request *CloudListQueueRequest, runtime *dara.RuntimeOptions) (_result *CloudListQueueResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.Offset) {
		query["Offset"] = request.Offset
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Qno) {
		query["Qno"] = request.Qno
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudListQueue"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudListQueueResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the queue info list.
//
// @param request - CloudListQueueRequest
//
// @return CloudListQueueResponse
func (client *Client) CloudListQueue(request *CloudListQueueRequest) (_result *CloudListQueueResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudListQueueResponse{}
	_body, _err := client.CloudListQueueWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains queue skill info.
//
// @param request - CloudListQueueSkillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudListQueueSkillResponse
func (client *Client) CloudListQueueSkillWithOptions(request *CloudListQueueSkillRequest, runtime *dara.RuntimeOptions) (_result *CloudListQueueSkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Qno) {
		query["Qno"] = request.Qno
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudListQueueSkill"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudListQueueSkillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains queue skill info.
//
// @param request - CloudListQueueSkillRequest
//
// @return CloudListQueueSkillResponse
func (client *Client) CloudListQueueSkill(request *CloudListQueueSkillRequest) (_result *CloudListQueueSkillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudListQueueSkillResponse{}
	_body, _err := client.CloudListQueueSkillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the skill info list.
//
// @param request - CloudListSkillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudListSkillResponse
func (client *Client) CloudListSkillWithOptions(request *CloudListSkillRequest, runtime *dara.RuntimeOptions) (_result *CloudListSkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Offset) {
		query["Offset"] = request.Offset
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudListSkill"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudListSkillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the skill info list.
//
// @param request - CloudListSkillRequest
//
// @return CloudListSkillResponse
func (client *Client) CloudListSkill(request *CloudListSkillRequest) (_result *CloudListSkillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudListSkillResponse{}
	_body, _err := client.CloudListSkillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query the call job batch list info.
//
// @param request - CloudListTaskFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudListTaskFileResponse
func (client *Client) CloudListTaskFileWithOptions(request *CloudListTaskFileRequest, runtime *dara.RuntimeOptions) (_result *CloudListTaskFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Start) {
		query["Start"] = request.Start
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudListTaskFile"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudListTaskFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the call job batch list info.
//
// @param request - CloudListTaskFileRequest
//
// @return CloudListTaskFileResponse
func (client *Client) CloudListTaskFile(request *CloudListTaskFileRequest) (_result *CloudListTaskFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudListTaskFileResponse{}
	_body, _err := client.CloudListTaskFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Call job monitoring.
//
// @param request - CloudMonitorTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudMonitorTaskResponse
func (client *Client) CloudMonitorTaskWithOptions(request *CloudMonitorTaskRequest, runtime *dara.RuntimeOptions) (_result *CloudMonitorTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudMonitorTask"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudMonitorTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Call job monitoring.
//
// @param request - CloudMonitorTaskRequest
//
// @return CloudMonitorTaskResponse
func (client *Client) CloudMonitorTask(request *CloudMonitorTaskRequest) (_result *CloudMonitorTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudMonitorTaskResponse{}
	_body, _err := client.CloudMonitorTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains preview outbound caller report data.
//
// @param request - CloudOutboundObClidReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudOutboundObClidReportResponse
func (client *Client) CloudOutboundObClidReportWithOptions(request *CloudOutboundObClidReportRequest, runtime *dara.RuntimeOptions) (_result *CloudOutboundObClidReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AreaCodes) {
		query["AreaCodes"] = request.AreaCodes
	}

	if !dara.IsNil(request.EndHour) {
		query["EndHour"] = request.EndHour
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.Start) {
		query["Start"] = request.Start
	}

	if !dara.IsNil(request.StartHour) {
		query["StartHour"] = request.StartHour
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StatisticMethod) {
		query["StatisticMethod"] = request.StatisticMethod
	}

	if !dara.IsNil(request.TimeRangeType) {
		query["TimeRangeType"] = request.TimeRangeType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudOutboundObClidReport"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudOutboundObClidReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains preview outbound caller report data.
//
// @param request - CloudOutboundObClidReportRequest
//
// @return CloudOutboundObClidReportResponse
func (client *Client) CloudOutboundObClidReport(request *CloudOutboundObClidReportRequest) (_result *CloudOutboundObClidReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudOutboundObClidReportResponse{}
	_body, _err := client.CloudOutboundObClidReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains a preview outbound report.
//
// @param request - CloudOutboundPreviewObReportRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudOutboundPreviewObReportResponse
func (client *Client) CloudOutboundPreviewObReportWithOptions(request *CloudOutboundPreviewObReportRequest, runtime *dara.RuntimeOptions) (_result *CloudOutboundPreviewObReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cnos) {
		query["Cnos"] = request.Cnos
	}

	if !dara.IsNil(request.EndHour) {
		query["EndHour"] = request.EndHour
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.Start) {
		query["Start"] = request.Start
	}

	if !dara.IsNil(request.StartHour) {
		query["StartHour"] = request.StartHour
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.StatisticMethod) {
		query["StatisticMethod"] = request.StatisticMethod
	}

	if !dara.IsNil(request.TimeRangeType) {
		query["TimeRangeType"] = request.TimeRangeType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudOutboundPreviewObReport"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudOutboundPreviewObReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a preview outbound report.
//
// @param request - CloudOutboundPreviewObReportRequest
//
// @return CloudOutboundPreviewObReportResponse
func (client *Client) CloudOutboundPreviewObReport(request *CloudOutboundPreviewObReportRequest) (_result *CloudOutboundPreviewObReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudOutboundPreviewObReportResponse{}
	_body, _err := client.CloudOutboundPreviewObReportWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Suspends a call job through this interface.
//
// @param request - CloudPauseTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudPauseTaskResponse
func (client *Client) CloudPauseTaskWithOptions(request *CloudPauseTaskRequest, runtime *dara.RuntimeOptions) (_result *CloudPauseTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PauseDuration) {
		query["PauseDuration"] = request.PauseDuration
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudPauseTask"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudPauseTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Suspends a call job through this interface.
//
// @param request - CloudPauseTaskRequest
//
// @return CloudPauseTaskResponse
func (client *Client) CloudPauseTask(request *CloudPauseTaskRequest) (_result *CloudPauseTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudPauseTaskResponse{}
	_body, _err := client.CloudPauseTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// When the agent is online, make a call through this interface.
//
// @param request - CloudPreviewoutcallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudPreviewoutcallResponse
func (client *Client) CloudPreviewoutcallWithOptions(request *CloudPreviewoutcallRequest, runtime *dara.RuntimeOptions) (_result *CloudPreviewoutcallResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BackupTels) {
		query["BackupTels"] = request.BackupTels
	}

	if !dara.IsNil(request.CallVariables) {
		query["CallVariables"] = request.CallVariables
	}

	if !dara.IsNil(request.CdrIsAsr) {
		query["CdrIsAsr"] = request.CdrIsAsr
	}

	if !dara.IsNil(request.ClidList) {
		query["ClidList"] = request.ClidList
	}

	if !dara.IsNil(request.Cno) {
		query["Cno"] = request.Cno
	}

	if !dara.IsNil(request.CrnId) {
		query["CrnId"] = request.CrnId
	}

	if !dara.IsNil(request.DialTelTimeout) {
		query["DialTelTimeout"] = request.DialTelTimeout
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.IsInvestigation) {
		query["IsInvestigation"] = request.IsInvestigation
	}

	if !dara.IsNil(request.ObClid) {
		query["ObClid"] = request.ObClid
	}

	if !dara.IsNil(request.ObClidAreaCode) {
		query["ObClidAreaCode"] = request.ObClidAreaCode
	}

	if !dara.IsNil(request.ObClidGroup) {
		query["ObClidGroup"] = request.ObClidGroup
	}

	if !dara.IsNil(request.RequestUniqueId) {
		query["RequestUniqueId"] = request.RequestUniqueId
	}

	if !dara.IsNil(request.Tel) {
		query["Tel"] = request.Tel
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudPreviewoutcall"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudPreviewoutcallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// When the agent is online, make a call through this interface.
//
// @param request - CloudPreviewoutcallRequest
//
// @return CloudPreviewoutcallResponse
func (client *Client) CloudPreviewoutcall(request *CloudPreviewoutcallRequest) (_result *CloudPreviewoutcallResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudPreviewoutcallResponse{}
	_body, _err := client.CloudPreviewoutcallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of agent details.
//
// @param request - CloudQueryAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudQueryAgentResponse
func (client *Client) CloudQueryAgentWithOptions(request *CloudQueryAgentRequest, runtime *dara.RuntimeOptions) (_result *CloudQueryAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Active) {
		query["Active"] = request.Active
	}

	if !dara.IsNil(request.Cnos) {
		query["Cnos"] = request.Cnos
	}

	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Qno) {
		query["Qno"] = request.Qno
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ReturnQueue) {
		query["ReturnQueue"] = request.ReturnQueue
	}

	if !dara.IsNil(request.Start) {
		query["Start"] = request.Start
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.WebrtcUrlType) {
		query["WebrtcUrlType"] = request.WebrtcUrlType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudQueryAgent"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudQueryAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of agent details.
//
// @param request - CloudQueryAgentRequest
//
// @return CloudQueryAgentResponse
func (client *Client) CloudQueryAgent(request *CloudQueryAgentRequest) (_result *CloudQueryAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudQueryAgentResponse{}
	_body, _err := client.CloudQueryAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get all agent numbers and parameter messages.
//
// @param request - CloudQueryAgentCnoAndNameRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudQueryAgentCnoAndNameResponse
func (client *Client) CloudQueryAgentCnoAndNameWithOptions(request *CloudQueryAgentCnoAndNameRequest, runtime *dara.RuntimeOptions) (_result *CloudQueryAgentCnoAndNameResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cnos) {
		query["Cnos"] = request.Cnos
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudQueryAgentCnoAndName"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudQueryAgentCnoAndNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get all agent numbers and parameter messages.
//
// @param request - CloudQueryAgentCnoAndNameRequest
//
// @return CloudQueryAgentCnoAndNameResponse
func (client *Client) CloudQueryAgentCnoAndName(request *CloudQueryAgentCnoAndNameRequest) (_result *CloudQueryAgentCnoAndNameResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudQueryAgentCnoAndNameResponse{}
	_body, _err := client.CloudQueryAgentCnoAndNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the info about the outbound group to which the agent belongs.
//
// @param request - CloudQueryAgentGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudQueryAgentGroupResponse
func (client *Client) CloudQueryAgentGroupWithOptions(request *CloudQueryAgentGroupRequest, runtime *dara.RuntimeOptions) (_result *CloudQueryAgentGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cno) {
		query["Cno"] = request.Cno
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudQueryAgentGroup"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudQueryAgentGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the info about the outbound group to which the agent belongs.
//
// @param request - CloudQueryAgentGroupRequest
//
// @return CloudQueryAgentGroupResponse
func (client *Client) CloudQueryAgentGroup(request *CloudQueryAgentGroupRequest) (_result *CloudQueryAgentGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudQueryAgentGroupResponse{}
	_body, _err := client.CloudQueryAgentGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query agent skills by agent number.
//
// @param request - CloudQueryAgentSkillRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudQueryAgentSkillResponse
func (client *Client) CloudQueryAgentSkillWithOptions(request *CloudQueryAgentSkillRequest, runtime *dara.RuntimeOptions) (_result *CloudQueryAgentSkillResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cnos) {
		query["Cnos"] = request.Cnos
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudQueryAgentSkill"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudQueryAgentSkillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query agent skills by agent number.
//
// @param request - CloudQueryAgentSkillRequest
//
// @return CloudQueryAgentSkillResponse
func (client *Client) CloudQueryAgentSkill(request *CloudQueryAgentSkillRequest) (_result *CloudQueryAgentSkillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudQueryAgentSkillResponse{}
	_body, _err := client.CloudQueryAgentSkillWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query incoming call records based on specified conditions.
//
// @param request - CloudQueryIbCdrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudQueryIbCdrResponse
func (client *Client) CloudQueryIbCdrWithOptions(request *CloudQueryIbCdrRequest, runtime *dara.RuntimeOptions) (_result *CloudQueryIbCdrResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CalleeNumber) {
		query["CalleeNumber"] = request.CalleeNumber
	}

	if !dara.IsNil(request.City) {
		query["City"] = request.City
	}

	if !dara.IsNil(request.Cno) {
		query["Cno"] = request.Cno
	}

	if !dara.IsNil(request.CustomerNumber) {
		query["CustomerNumber"] = request.CustomerNumber
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Hotline) {
		query["Hotline"] = request.Hotline
	}

	if !dara.IsNil(request.JoinQueueCode) {
		query["JoinQueueCode"] = request.JoinQueueCode
	}

	if !dara.IsNil(request.LeaveQueueCode) {
		query["LeaveQueueCode"] = request.LeaveQueueCode
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.Province) {
		query["Province"] = request.Province
	}

	if !dara.IsNil(request.Qno) {
		query["Qno"] = request.Qno
	}

	if !dara.IsNil(request.Start) {
		query["Start"] = request.Start
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TimeRangeType) {
		query["TimeRangeType"] = request.TimeRangeType
	}

	if !dara.IsNil(request.UniqueId) {
		query["UniqueId"] = request.UniqueId
	}

	if !dara.IsNil(request.UserFieldValue) {
		query["UserFieldValue"] = request.UserFieldValue
	}

	if !dara.IsNil(request.UserFieldkey) {
		query["UserFieldkey"] = request.UserFieldkey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudQueryIbCdr"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudQueryIbCdrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query incoming call records based on specified conditions.
//
// @param request - CloudQueryIbCdrRequest
//
// @return CloudQueryIbCdrResponse
func (client *Client) CloudQueryIbCdr(request *CloudQueryIbCdrRequest) (_result *CloudQueryIbCdrResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudQueryIbCdrResponse{}
	_body, _err := client.CloudQueryIbCdrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query agent outbound call records based on conditions.
//
// @param request - CloudQueryObCdrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudQueryObCdrResponse
func (client *Client) CloudQueryObCdrWithOptions(request *CloudQueryObCdrRequest, runtime *dara.RuntimeOptions) (_result *CloudQueryObCdrResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentName) {
		query["AgentName"] = request.AgentName
	}

	if !dara.IsNil(request.AgentNumber) {
		query["AgentNumber"] = request.AgentNumber
	}

	if !dara.IsNil(request.CallType) {
		query["CallType"] = request.CallType
	}

	if !dara.IsNil(request.City) {
		query["City"] = request.City
	}

	if !dara.IsNil(request.Clid) {
		query["Clid"] = request.Clid
	}

	if !dara.IsNil(request.Cno) {
		query["Cno"] = request.Cno
	}

	if !dara.IsNil(request.CustomerNumber) {
		query["CustomerNumber"] = request.CustomerNumber
	}

	if !dara.IsNil(request.DownGrade) {
		query["DownGrade"] = request.DownGrade
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Gno) {
		query["Gno"] = request.Gno
	}

	if !dara.IsNil(request.IsRealAnswer) {
		query["IsRealAnswer"] = request.IsRealAnswer
	}

	if !dara.IsNil(request.LeftDisplayNumber) {
		query["LeftDisplayNumber"] = request.LeftDisplayNumber
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.Province) {
		query["Province"] = request.Province
	}

	if !dara.IsNil(request.RequestUniqueId) {
		query["RequestUniqueId"] = request.RequestUniqueId
	}

	if !dara.IsNil(request.ReturnCount) {
		query["ReturnCount"] = request.ReturnCount
	}

	if !dara.IsNil(request.Start) {
		query["Start"] = request.Start
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TimeRangeType) {
		query["TimeRangeType"] = request.TimeRangeType
	}

	if !dara.IsNil(request.UniqueId) {
		query["UniqueId"] = request.UniqueId
	}

	if !dara.IsNil(request.UserFieldValue) {
		query["UserFieldValue"] = request.UserFieldValue
	}

	if !dara.IsNil(request.UserFieldkey) {
		query["UserFieldkey"] = request.UserFieldkey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudQueryObCdr"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudQueryObCdrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query agent outbound call records based on conditions.
//
// @param request - CloudQueryObCdrRequest
//
// @return CloudQueryObCdrResponse
func (client *Client) CloudQueryObCdr(request *CloudQueryObCdrRequest) (_result *CloudQueryObCdrResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudQueryObCdrResponse{}
	_body, _err := client.CloudQueryObCdrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query the call record info of predictive outbound calls.
//
// @param request - CloudQueryPredictiveCallCdrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudQueryPredictiveCallCdrResponse
func (client *Client) CloudQueryPredictiveCallCdrWithOptions(request *CloudQueryPredictiveCallCdrRequest, runtime *dara.RuntimeOptions) (_result *CloudQueryPredictiveCallCdrResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentName) {
		query["AgentName"] = request.AgentName
	}

	if !dara.IsNil(request.City) {
		query["City"] = request.City
	}

	if !dara.IsNil(request.Clid) {
		query["Clid"] = request.Clid
	}

	if !dara.IsNil(request.Cno) {
		query["Cno"] = request.Cno
	}

	if !dara.IsNil(request.CustomerNumber) {
		query["CustomerNumber"] = request.CustomerNumber
	}

	if !dara.IsNil(request.DisplayNumber) {
		query["DisplayNumber"] = request.DisplayNumber
	}

	if !dara.IsNil(request.DownGrade) {
		query["DownGrade"] = request.DownGrade
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Gno) {
		query["Gno"] = request.Gno
	}

	if !dara.IsNil(request.IsRealAnswer) {
		query["IsRealAnswer"] = request.IsRealAnswer
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.Province) {
		query["Province"] = request.Province
	}

	if !dara.IsNil(request.RequestUniqueId) {
		query["RequestUniqueId"] = request.RequestUniqueId
	}

	if !dara.IsNil(request.Start) {
		query["Start"] = request.Start
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskFileId) {
		query["TaskFileId"] = request.TaskFileId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TimeRangeType) {
		query["TimeRangeType"] = request.TimeRangeType
	}

	if !dara.IsNil(request.UniqueId) {
		query["UniqueId"] = request.UniqueId
	}

	if !dara.IsNil(request.UserFieldValue) {
		query["UserFieldValue"] = request.UserFieldValue
	}

	if !dara.IsNil(request.UserFieldkey) {
		query["UserFieldkey"] = request.UserFieldkey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudQueryPredictiveCallCdr"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudQueryPredictiveCallCdrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the call record info of predictive outbound calls.
//
// @param request - CloudQueryPredictiveCallCdrRequest
//
// @return CloudQueryPredictiveCallCdrResponse
func (client *Client) CloudQueryPredictiveCallCdr(request *CloudQueryPredictiveCallCdrRequest) (_result *CloudQueryPredictiveCallCdrResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudQueryPredictiveCallCdrResponse{}
	_body, _err := client.CloudQueryPredictiveCallCdrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtain the rasr info based on the uniqueId.
//
// @param request - CloudQueryRasrEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudQueryRasrEventResponse
func (client *Client) CloudQueryRasrEventWithOptions(request *CloudQueryRasrEventRequest, runtime *dara.RuntimeOptions) (_result *CloudQueryRasrEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.UniqueId) {
		query["UniqueId"] = request.UniqueId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudQueryRasrEvent"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudQueryRasrEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the rasr info based on the uniqueId.
//
// @param request - CloudQueryRasrEventRequest
//
// @return CloudQueryRasrEventResponse
func (client *Client) CloudQueryRasrEvent(request *CloudQueryRasrEventRequest) (_result *CloudQueryRasrEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudQueryRasrEventResponse{}
	_body, _err := client.CloudQueryRasrEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries call jobs. Supports conditional query.
//
// @param request - CloudQueryTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudQueryTaskResponse
func (client *Client) CloudQueryTaskWithOptions(request *CloudQueryTaskRequest, runtime *dara.RuntimeOptions) (_result *CloudQueryTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AutoStart) {
		query["AutoStart"] = request.AutoStart
	}

	if !dara.IsNil(request.AutoStop) {
		query["AutoStop"] = request.AutoStop
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Start) {
		query["Start"] = request.Start
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TimeType) {
		query["TimeType"] = request.TimeType
	}

	if !dara.IsNil(request.Type) {
		query["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudQueryTask"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudQueryTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries call jobs. Supports conditional query.
//
// @param request - CloudQueryTaskRequest
//
// @return CloudQueryTaskResponse
func (client *Client) CloudQueryTask(request *CloudQueryTaskRequest) (_result *CloudQueryTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudQueryTaskResponse{}
	_body, _err := client.CloudQueryTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query webCall call record info through this interface.
//
// @param request - CloudQueryWebcallCdrRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudQueryWebcallCdrResponse
func (client *Client) CloudQueryWebcallCdrWithOptions(request *CloudQueryWebcallCdrRequest, runtime *dara.RuntimeOptions) (_result *CloudQueryWebcallCdrResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CalleeClid) {
		query["CalleeClid"] = request.CalleeClid
	}

	if !dara.IsNil(request.CalleeDisplayNumber) {
		query["CalleeDisplayNumber"] = request.CalleeDisplayNumber
	}

	if !dara.IsNil(request.CalleeNumber) {
		query["CalleeNumber"] = request.CalleeNumber
	}

	if !dara.IsNil(request.City) {
		query["City"] = request.City
	}

	if !dara.IsNil(request.Clid) {
		query["Clid"] = request.Clid
	}

	if !dara.IsNil(request.Cno) {
		query["Cno"] = request.Cno
	}

	if !dara.IsNil(request.CustomerNumber) {
		query["CustomerNumber"] = request.CustomerNumber
	}

	if !dara.IsNil(request.DisplayNumber) {
		query["DisplayNumber"] = request.DisplayNumber
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Gno) {
		query["Gno"] = request.Gno
	}

	if !dara.IsNil(request.IsRealAnswer) {
		query["IsRealAnswer"] = request.IsRealAnswer
	}

	if !dara.IsNil(request.Limit) {
		query["Limit"] = request.Limit
	}

	if !dara.IsNil(request.Province) {
		query["Province"] = request.Province
	}

	if !dara.IsNil(request.RequestUniqueId) {
		query["RequestUniqueId"] = request.RequestUniqueId
	}

	if !dara.IsNil(request.ReturnCount) {
		query["ReturnCount"] = request.ReturnCount
	}

	if !dara.IsNil(request.Start) {
		query["Start"] = request.Start
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TimeRangeType) {
		query["TimeRangeType"] = request.TimeRangeType
	}

	if !dara.IsNil(request.UniqueId) {
		query["UniqueId"] = request.UniqueId
	}

	if !dara.IsNil(request.UserFieldValue) {
		query["UserFieldValue"] = request.UserFieldValue
	}

	if !dara.IsNil(request.UserFieldkey) {
		query["UserFieldkey"] = request.UserFieldkey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudQueryWebcallCdr"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudQueryWebcallCdrResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query webCall call record info through this interface.
//
// @param request - CloudQueryWebcallCdrRequest
//
// @return CloudQueryWebcallCdrResponse
func (client *Client) CloudQueryWebcallCdr(request *CloudQueryWebcallCdrRequest) (_result *CloudQueryWebcallCdrResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudQueryWebcallCdrResponse{}
	_body, _err := client.CloudQueryWebcallCdrWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts an outbound call task.
//
// @param request - CloudStartTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudStartTaskResponse
func (client *Client) CloudStartTaskWithOptions(request *CloudStartTaskRequest, runtime *dara.RuntimeOptions) (_result *CloudStartTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudStartTask"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudStartTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts an outbound call task.
//
// @param request - CloudStartTaskRequest
//
// @return CloudStartTaskResponse
func (client *Client) CloudStartTask(request *CloudStartTaskRequest) (_result *CloudStartTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudStartTaskResponse{}
	_body, _err := client.CloudStartTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Purges the attachment between an outbound call group and agents.
//
// @param request - CloudUnassignAgentGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudUnassignAgentGroupResponse
func (client *Client) CloudUnassignAgentGroupWithOptions(request *CloudUnassignAgentGroupRequest, runtime *dara.RuntimeOptions) (_result *CloudUnassignAgentGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Cno) {
		query["Cno"] = request.Cno
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.Gno) {
		query["Gno"] = request.Gno
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudUnassignAgentGroup"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudUnassignAgentGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Purges the attachment between an outbound call group and agents.
//
// @param request - CloudUnassignAgentGroupRequest
//
// @return CloudUnassignAgentGroupResponse
func (client *Client) CloudUnassignAgentGroup(request *CloudUnassignAgentGroupRequest) (_result *CloudUnassignAgentGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudUnassignAgentGroupResponse{}
	_body, _err := client.CloudUnassignAgentGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Update agent basic info.
//
// @param request - CloudUpdateAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudUpdateAgentResponse
func (client *Client) CloudUpdateAgentWithOptions(request *CloudUpdateAgentRequest, runtime *dara.RuntimeOptions) (_result *CloudUpdateAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Active) {
		query["Active"] = request.Active
	}

	if !dara.IsNil(request.AgentType) {
		query["AgentType"] = request.AgentType
	}

	if !dara.IsNil(request.AreaCode) {
		query["AreaCode"] = request.AreaCode
	}

	if !dara.IsNil(request.CallPower) {
		query["CallPower"] = request.CallPower
	}

	if !dara.IsNil(request.Cno) {
		query["Cno"] = request.Cno
	}

	if !dara.IsNil(request.Comment) {
		query["Comment"] = request.Comment
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.IbRecord) {
		query["IbRecord"] = request.IbRecord
	}

	if !dara.IsNil(request.IsAsr) {
		query["IsAsr"] = request.IsAsr
	}

	if !dara.IsNil(request.IsOb) {
		query["IsOb"] = request.IsOb
	}

	if !dara.IsNil(request.IsObRemember) {
		query["IsObRemember"] = request.IsObRemember
	}

	if !dara.IsNil(request.IsQualityCheck) {
		query["IsQualityCheck"] = request.IsQualityCheck
	}

	if !dara.IsNil(request.IsRandom) {
		query["IsRandom"] = request.IsRandom
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.ObClid) {
		query["ObClid"] = request.ObClid
	}

	if !dara.IsNil(request.ObClidProperty) {
		query["ObClidProperty"] = request.ObClidProperty
	}

	if !dara.IsNil(request.ObClidType) {
		query["ObClidType"] = request.ObClidType
	}

	if !dara.IsNil(request.ObRecord) {
		query["ObRecord"] = request.ObRecord
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PermitObPreviewTime) {
		query["PermitObPreviewTime"] = request.PermitObPreviewTime
	}

	if !dara.IsNil(request.Power) {
		query["Power"] = request.Power
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SkillIds) {
		query["SkillIds"] = request.SkillIds
	}

	if !dara.IsNil(request.SkillLevels) {
		query["SkillLevels"] = request.SkillLevels
	}

	if !dara.IsNil(request.WebrtcUrlType) {
		query["WebrtcUrlType"] = request.WebrtcUrlType
	}

	if !dara.IsNil(request.Wrapup) {
		query["Wrapup"] = request.Wrapup
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudUpdateAgent"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudUpdateAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update agent basic info.
//
// @param request - CloudUpdateAgentRequest
//
// @return CloudUpdateAgentResponse
func (client *Client) CloudUpdateAgent(request *CloudUpdateAgentRequest) (_result *CloudUpdateAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudUpdateAgentResponse{}
	_body, _err := client.CloudUpdateAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the basic information of an outbound call job.
//
// @param request - CloudUpdateTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudUpdateTaskResponse
func (client *Client) CloudUpdateTaskWithOptions(request *CloudUpdateTaskRequest, runtime *dara.RuntimeOptions) (_result *CloudUpdateTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentGroup) {
		query["AgentGroup"] = request.AgentGroup
	}

	if !dara.IsNil(request.AgentTimeout) {
		query["AgentTimeout"] = request.AgentTimeout
	}

	if !dara.IsNil(request.AnswerRate) {
		query["AnswerRate"] = request.AnswerRate
	}

	if !dara.IsNil(request.AutoComplete) {
		query["AutoComplete"] = request.AutoComplete
	}

	if !dara.IsNil(request.AutoStart) {
		query["AutoStart"] = request.AutoStart
	}

	if !dara.IsNil(request.AutoStartDay) {
		query["AutoStartDay"] = request.AutoStartDay
	}

	if !dara.IsNil(request.AutoStartTime) {
		query["AutoStartTime"] = request.AutoStartTime
	}

	if !dara.IsNil(request.AutoStop) {
		query["AutoStop"] = request.AutoStop
	}

	if !dara.IsNil(request.AutoStopDay) {
		query["AutoStopDay"] = request.AutoStopDay
	}

	if !dara.IsNil(request.AutoStopTime) {
		query["AutoStopTime"] = request.AutoStopTime
	}

	if !dara.IsNil(request.AutoTaskType) {
		query["AutoTaskType"] = request.AutoTaskType
	}

	if !dara.IsNil(request.AutoTriggerTimeStrategy) {
		query["AutoTriggerTimeStrategy"] = request.AutoTriggerTimeStrategy
	}

	if !dara.IsNil(request.CallLimitStrategy) {
		query["CallLimitStrategy"] = request.CallLimitStrategy
	}

	if !dara.IsNil(request.CallPriorityStrategy) {
		query["CallPriorityStrategy"] = request.CallPriorityStrategy
	}

	if !dara.IsNil(request.CallRouteStrategy) {
		query["CallRouteStrategy"] = request.CallRouteStrategy
	}

	if !dara.IsNil(request.CallStrategy) {
		query["CallStrategy"] = request.CallStrategy
	}

	if !dara.IsNil(request.CallVariables) {
		query["CallVariables"] = request.CallVariables
	}

	if !dara.IsNil(request.ClidProperty) {
		query["ClidProperty"] = request.ClidProperty
	}

	if !dara.IsNil(request.Cnos) {
		query["Cnos"] = request.Cnos
	}

	if !dara.IsNil(request.Concurrency) {
		query["Concurrency"] = request.Concurrency
	}

	if !dara.IsNil(request.CustomerClidType) {
		query["CustomerClidType"] = request.CustomerClidType
	}

	if !dara.IsNil(request.CustomerClidWeight) {
		query["CustomerClidWeight"] = request.CustomerClidWeight
	}

	if !dara.IsNil(request.CustomerClidWeightFlag) {
		query["CustomerClidWeightFlag"] = request.CustomerClidWeightFlag
	}

	if !dara.IsNil(request.CustomerClids) {
		query["CustomerClids"] = request.CustomerClids
	}

	if !dara.IsNil(request.CustomerClidsCategory) {
		query["CustomerClidsCategory"] = request.CustomerClidsCategory
	}

	if !dara.IsNil(request.CustomerClidsGroup) {
		query["CustomerClidsGroup"] = request.CustomerClidsGroup
	}

	if !dara.IsNil(request.CustomerMoh) {
		query["CustomerMoh"] = request.CustomerMoh
	}

	if !dara.IsNil(request.CustomerTimeout) {
		query["CustomerTimeout"] = request.CustomerTimeout
	}

	if !dara.IsNil(request.CustomerVoice) {
		query["CustomerVoice"] = request.CustomerVoice
	}

	if !dara.IsNil(request.Description) {
		query["Description"] = request.Description
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.ForceEndFlag) {
		query["ForceEndFlag"] = request.ForceEndFlag
	}

	if !dara.IsNil(request.IsRewarm) {
		query["IsRewarm"] = request.IsRewarm
	}

	if !dara.IsNil(request.IvrId) {
		query["IvrId"] = request.IvrId
	}

	if !dara.IsNil(request.IvrName) {
		query["IvrName"] = request.IvrName
	}

	if !dara.IsNil(request.MaxWaitTime) {
		query["MaxWaitTime"] = request.MaxWaitTime
	}

	if !dara.IsNil(request.MinAvailableAgentCount) {
		query["MinAvailableAgentCount"] = request.MinAvailableAgentCount
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PredictAdjust) {
		query["PredictAdjust"] = request.PredictAdjust
	}

	if !dara.IsNil(request.Quotiety) {
		query["Quotiety"] = request.Quotiety
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RetryStrategy) {
		query["RetryStrategy"] = request.RetryStrategy
	}

	if !dara.IsNil(request.RetryStrategyOnlyToday) {
		query["RetryStrategyOnlyToday"] = request.RetryStrategyOnlyToday
	}

	if !dara.IsNil(request.RetryStrategyTimeType) {
		query["RetryStrategyTimeType"] = request.RetryStrategyTimeType
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TimeStrategy) {
		query["TimeStrategy"] = request.TimeStrategy
	}

	if !dara.IsNil(request.UserFields) {
		query["UserFields"] = request.UserFields
	}

	if !dara.IsNil(request.WarmUpDuration) {
		query["WarmUpDuration"] = request.WarmUpDuration
	}

	if !dara.IsNil(request.Wrapup) {
		query["Wrapup"] = request.Wrapup
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudUpdateTask"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudUpdateTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the basic information of an outbound call job.
//
// @param request - CloudUpdateTaskRequest
//
// @return CloudUpdateTaskResponse
func (client *Client) CloudUpdateTask(request *CloudUpdateTaskRequest) (_result *CloudUpdateTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudUpdateTaskResponse{}
	_body, _err := client.CloudUpdateTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Through this interface, a third-party platform can send a call request to a hosted call center. The system calls the customer first and then the agent, and connects both parties.
//
// @param request - CloudWebcallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloudWebcallResponse
func (client *Client) CloudWebcallWithOptions(request *CloudWebcallRequest, runtime *dara.RuntimeOptions) (_result *CloudWebcallResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Amd) {
		query["Amd"] = request.Amd
	}

	if !dara.IsNil(request.Clid) {
		query["Clid"] = request.Clid
	}

	if !dara.IsNil(request.ClidAreaCode) {
		query["ClidAreaCode"] = request.ClidAreaCode
	}

	if !dara.IsNil(request.ClidGroup) {
		query["ClidGroup"] = request.ClidGroup
	}

	if !dara.IsNil(request.ClidList) {
		query["ClidList"] = request.ClidList
	}

	if !dara.IsNil(request.CrnId) {
		query["CrnId"] = request.CrnId
	}

	if !dara.IsNil(request.Delay) {
		query["Delay"] = request.Delay
	}

	if !dara.IsNil(request.EnterpriseId) {
		query["EnterpriseId"] = request.EnterpriseId
	}

	if !dara.IsNil(request.ExpirTime) {
		query["ExpirTime"] = request.ExpirTime
	}

	if !dara.IsNil(request.IvrId) {
		query["IvrId"] = request.IvrId
	}

	if !dara.IsNil(request.MultiTelDelay) {
		query["MultiTelDelay"] = request.MultiTelDelay
	}

	if !dara.IsNil(request.MultiTelPush) {
		query["MultiTelPush"] = request.MultiTelPush
	}

	if !dara.IsNil(request.RequestUniqueId) {
		query["RequestUniqueId"] = request.RequestUniqueId
	}

	if !dara.IsNil(request.Retry) {
		query["Retry"] = request.Retry
	}

	if !dara.IsNil(request.RetryInterval) {
		query["RetryInterval"] = request.RetryInterval
	}

	if !dara.IsNil(request.Tel) {
		query["Tel"] = request.Tel
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	if !dara.IsNil(request.UserField) {
		query["UserField"] = request.UserField
	}

	if !dara.IsNil(request.Vid) {
		query["Vid"] = request.Vid
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloudWebcall"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloudWebcallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Through this interface, a third-party platform can send a call request to a hosted call center. The system calls the customer first and then the agent, and connects both parties.
//
// @param request - CloudWebcallRequest
//
// @return CloudWebcallResponse
func (client *Client) CloudWebcall(request *CloudWebcallRequest) (_result *CloudWebcallResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CloudWebcallResponse{}
	_body, _err := client.CloudWebcallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uses a service instance to create a text-to-speech (TTS) task, a voice notification task, or a voice verification code task for multiple called numbers.
//
// Description:
//
// You can create up to 1,000 voice notifications for each task.
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - CreateCallTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCallTaskResponse
func (client *Client) CreateCallTaskWithOptions(request *CreateCallTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateCallTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.Data) {
		query["Data"] = request.Data
	}

	if !dara.IsNil(request.DataType) {
		query["DataType"] = request.DataType
	}

	if !dara.IsNil(request.FireTime) {
		query["FireTime"] = request.FireTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Resource) {
		query["Resource"] = request.Resource
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.ScheduleType) {
		query["ScheduleType"] = request.ScheduleType
	}

	if !dara.IsNil(request.StopTime) {
		query["StopTime"] = request.StopTime
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.TemplateCode) {
		query["TemplateCode"] = request.TemplateCode
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCallTask"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCallTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uses a service instance to create a text-to-speech (TTS) task, a voice notification task, or a voice verification code task for multiple called numbers.
//
// Description:
//
// You can create up to 1,000 voice notifications for each task.
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - CreateCallTaskRequest
//
// @return CreateCallTaskResponse
func (client *Client) CreateCallTask(request *CreateCallTaskRequest) (_result *CreateCallTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateCallTaskResponse{}
	_body, _err := client.CreateCallTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Initiates an outbound robocall task.
//
// Description:
//
// You can call this operation to initiate an outbound robocall task by using the robot communication scripts preset in the Voice Messaging Service console. In an intelligent speech interaction task, you can use the robot communication scripts preset in the Voice Messaging Service console, or invoke the callback function to return the response mode configured by the business party in each call.
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - CreateRobotTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRobotTaskResponse
func (client *Client) CreateRobotTaskWithOptions(request *CreateRobotTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateRobotTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Caller) {
		query["Caller"] = request.Caller
	}

	if !dara.IsNil(request.CorpName) {
		query["CorpName"] = request.CorpName
	}

	if !dara.IsNil(request.DialogId) {
		query["DialogId"] = request.DialogId
	}

	if !dara.IsNil(request.IsSelfLine) {
		query["IsSelfLine"] = request.IsSelfLine
	}

	if !dara.IsNil(request.NumberStatusIdent) {
		query["NumberStatusIdent"] = request.NumberStatusIdent
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RecallInterval) {
		query["RecallInterval"] = request.RecallInterval
	}

	if !dara.IsNil(request.RecallStateCodes) {
		query["RecallStateCodes"] = request.RecallStateCodes
	}

	if !dara.IsNil(request.RecallTimes) {
		query["RecallTimes"] = request.RecallTimes
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RetryType) {
		query["RetryType"] = request.RetryType
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRobotTask"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRobotTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiates an outbound robocall task.
//
// Description:
//
// You can call this operation to initiate an outbound robocall task by using the robot communication scripts preset in the Voice Messaging Service console. In an intelligent speech interaction task, you can use the robot communication scripts preset in the Voice Messaging Service console, or invoke the callback function to return the response mode configured by the business party in each call.
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - CreateRobotTaskRequest
//
// @return CreateRobotTaskResponse
func (client *Client) CreateRobotTask(request *CreateRobotTaskRequest) (_result *CreateRobotTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateRobotTaskResponse{}
	_body, _err := client.CreateRobotTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Downgrades from a video call to a voice call.
//
// @param request - DegradeVideoFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DegradeVideoFileResponse
func (client *Client) DegradeVideoFileWithOptions(request *DegradeVideoFileRequest, runtime *dara.RuntimeOptions) (_result *DegradeVideoFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallId) {
		query["CallId"] = request.CallId
	}

	if !dara.IsNil(request.CalledNumber) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !dara.IsNil(request.MediaType) {
		query["MediaType"] = request.MediaType
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DegradeVideoFile"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DegradeVideoFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Downgrades from a video call to a voice call.
//
// @param request - DegradeVideoFileRequest
//
// @return DegradeVideoFileResponse
func (client *Client) DegradeVideoFile(request *DegradeVideoFileRequest) (_result *DegradeVideoFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DegradeVideoFileResponse{}
	_body, _err := client.DegradeVideoFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a robocall task.
//
// Description:
//
// You can call this operation to delete only tasks that are not started, that are completed, and that are terminated.
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - DeleteRobotTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRobotTaskResponse
func (client *Client) DeleteRobotTaskWithOptions(request *DeleteRobotTaskRequest, runtime *dara.RuntimeOptions) (_result *DeleteRobotTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteRobotTask"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRobotTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a robocall task.
//
// Description:
//
// You can call this operation to delete only tasks that are not started, that are completed, and that are terminated.
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - DeleteRobotTaskRequest
//
// @return DeleteRobotTaskResponse
func (client *Client) DeleteRobotTask(request *DeleteRobotTaskRequest) (_result *DeleteRobotTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteRobotTaskResponse{}
	_body, _err := client.DeleteRobotTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Executes a call task.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - ExecuteCallTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteCallTaskResponse
func (client *Client) ExecuteCallTaskWithOptions(request *ExecuteCallTaskRequest, runtime *dara.RuntimeOptions) (_result *ExecuteCallTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FireTime) {
		query["FireTime"] = request.FireTime
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteCallTask"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteCallTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Executes a call task.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - ExecuteCallTaskRequest
//
// @return ExecuteCallTaskResponse
func (client *Client) ExecuteCallTask(request *ExecuteCallTaskRequest) (_result *ExecuteCallTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExecuteCallTaskResponse{}
	_body, _err := client.ExecuteCallTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the call type during a call.
//
// @param request - GetCallMediaTypeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCallMediaTypeResponse
func (client *Client) GetCallMediaTypeWithOptions(request *GetCallMediaTypeRequest, runtime *dara.RuntimeOptions) (_result *GetCallMediaTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallId) {
		query["CallId"] = request.CallId
	}

	if !dara.IsNil(request.CalledNumber) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCallMediaType"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCallMediaTypeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the call type during a call.
//
// @param request - GetCallMediaTypeRequest
//
// @return GetCallMediaTypeResponse
func (client *Client) GetCallMediaType(request *GetCallMediaTypeRequest) (_result *GetCallMediaTypeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCallMediaTypeResponse{}
	_body, _err := client.GetCallMediaTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # GetCallProgress
//
// @param request - GetCallProgressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCallProgressResponse
func (client *Client) GetCallProgressWithOptions(request *GetCallProgressRequest, runtime *dara.RuntimeOptions) (_result *GetCallProgressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallId) {
		query["CallId"] = request.CallId
	}

	if !dara.IsNil(request.CalledNum) {
		query["CalledNum"] = request.CalledNum
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCallProgress"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCallProgressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetCallProgress
//
// @param request - GetCallProgressRequest
//
// @return GetCallProgressResponse
func (client *Client) GetCallProgress(request *GetCallProgressRequest) (_result *GetCallProgressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetCallProgressResponse{}
	_body, _err := client.GetCallProgressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the qualification ID based on the ID of a qualification application ticket.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - GetHotlineQualificationByOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHotlineQualificationByOrderResponse
func (client *Client) GetHotlineQualificationByOrderWithOptions(request *GetHotlineQualificationByOrderRequest, runtime *dara.RuntimeOptions) (_result *GetHotlineQualificationByOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OrderId) {
		query["OrderId"] = request.OrderId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHotlineQualificationByOrder"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHotlineQualificationByOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the qualification ID based on the ID of a qualification application ticket.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - GetHotlineQualificationByOrderRequest
//
// @return GetHotlineQualificationByOrderResponse
func (client *Client) GetHotlineQualificationByOrder(request *GetHotlineQualificationByOrderRequest) (_result *GetHotlineQualificationByOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetHotlineQualificationByOrderResponse{}
	_body, _err := client.GetHotlineQualificationByOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains a temporary URL of a video or audio file. You can view the video or audio file immediately by using this temporary URL.
//
// @param request - GetTemporaryFileUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTemporaryFileUrlResponse
func (client *Client) GetTemporaryFileUrlWithOptions(request *GetTemporaryFileUrlRequest, runtime *dara.RuntimeOptions) (_result *GetTemporaryFileUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.VideoId) {
		query["VideoId"] = request.VideoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTemporaryFileUrl"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTemporaryFileUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains a temporary URL of a video or audio file. You can view the video or audio file immediately by using this temporary URL.
//
// @param request - GetTemporaryFileUrlRequest
//
// @return GetTemporaryFileUrlResponse
func (client *Client) GetTemporaryFileUrl(request *GetTemporaryFileUrlRequest) (_result *GetTemporaryFileUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTemporaryFileUrlResponse{}
	_body, _err := client.GetTemporaryFileUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains the token for authentication.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to five times per second per account.
//
// @param request - GetTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTokenResponse
func (client *Client) GetTokenWithOptions(request *GetTokenRequest, runtime *dara.RuntimeOptions) (_result *GetTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TokenType) {
		query["TokenType"] = request.TokenType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetToken"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the token for authentication.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to five times per second per account.
//
// @param request - GetTokenRequest
//
// @return GetTokenResponse
func (client *Client) GetToken(request *GetTokenRequest) (_result *GetTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetTokenResponse{}
	_body, _err := client.GetTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # GetVideoFieldUrl
//
// @param request - GetVideoFieldUrlRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVideoFieldUrlResponse
func (client *Client) GetVideoFieldUrlWithOptions(request *GetVideoFieldUrlRequest, runtime *dara.RuntimeOptions) (_result *GetVideoFieldUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.VideoFile) {
		query["VideoFile"] = request.VideoFile
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVideoFieldUrl"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVideoFieldUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetVideoFieldUrl
//
// @param request - GetVideoFieldUrlRequest
//
// @return GetVideoFieldUrlResponse
func (client *Client) GetVideoFieldUrl(request *GetVideoFieldUrlRequest) (_result *GetVideoFieldUrlResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetVideoFieldUrlResponse{}
	_body, _err := client.GetVideoFieldUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Initiates an interactive voice response (IVR) call to a specified number.
//
// Description:
//
//	  Your enterprise qualification is approved. For more information, see [Submit enterprise qualifications](https://help.aliyun.com/document_detail/149795.html).
//
//		- Voice numbers are purchased. For more information, see [Purchase numbers](https://help.aliyun.com/document_detail/149794.html).
//
//		- When the subscriber answers the call, the subscriber hears a voice that instructs the subscriber to press a key as needed. If the [message receipt](https://help.aliyun.com/document_detail/112503.html) feature is enabled, the Voice Messaging Service (VMS) platform returns the information about the key pressed by the subscriber to the business system. The key information includes the order confirmation, questionnaire survey, and satisfaction survey completed by the subscriber.
//
// ## QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - IvrCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IvrCallResponse
func (client *Client) IvrCallWithOptions(request *IvrCallRequest, runtime *dara.RuntimeOptions) (_result *IvrCallResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ByeCode) {
		query["ByeCode"] = request.ByeCode
	}

	if !dara.IsNil(request.ByeTtsParams) {
		query["ByeTtsParams"] = request.ByeTtsParams
	}

	if !dara.IsNil(request.CalledNumber) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !dara.IsNil(request.CalledShowNumber) {
		query["CalledShowNumber"] = request.CalledShowNumber
	}

	if !dara.IsNil(request.MenuKeyMap) {
		query["MenuKeyMap"] = request.MenuKeyMap
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PlayTimes) {
		query["PlayTimes"] = request.PlayTimes
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.StartCode) {
		query["StartCode"] = request.StartCode
	}

	if !dara.IsNil(request.StartTtsParams) {
		query["StartTtsParams"] = request.StartTtsParams
	}

	if !dara.IsNil(request.Timeout) {
		query["Timeout"] = request.Timeout
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IvrCall"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &IvrCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiates an interactive voice response (IVR) call to a specified number.
//
// Description:
//
//	  Your enterprise qualification is approved. For more information, see [Submit enterprise qualifications](https://help.aliyun.com/document_detail/149795.html).
//
//		- Voice numbers are purchased. For more information, see [Purchase numbers](https://help.aliyun.com/document_detail/149794.html).
//
//		- When the subscriber answers the call, the subscriber hears a voice that instructs the subscriber to press a key as needed. If the [message receipt](https://help.aliyun.com/document_detail/112503.html) feature is enabled, the Voice Messaging Service (VMS) platform returns the information about the key pressed by the subscriber to the business system. The key information includes the order confirmation, questionnaire survey, and satisfaction survey completed by the subscriber.
//
// ## QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - IvrCallRequest
//
// @return IvrCallResponse
func (client *Client) IvrCall(request *IvrCallRequest) (_result *IvrCallResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &IvrCallResponse{}
	_body, _err := client.IvrCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a voice call task after the task is created, including the task ID, task status, and templates used by the task.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - ListCallTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCallTaskResponse
func (client *Client) ListCallTaskWithOptions(request *ListCallTaskRequest, runtime *dara.RuntimeOptions) (_result *ListCallTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCallTask"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCallTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a voice call task after the task is created, including the task ID, task status, and templates used by the task.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - ListCallTaskRequest
//
// @return ListCallTaskResponse
func (client *Client) ListCallTask(request *ListCallTaskRequest) (_result *ListCallTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCallTaskResponse{}
	_body, _err := client.ListCallTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of call tasks based on task IDs after call tasks are complete.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - ListCallTaskDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCallTaskDetailResponse
func (client *Client) ListCallTaskDetailWithOptions(request *ListCallTaskDetailRequest, runtime *dara.RuntimeOptions) (_result *ListCallTaskDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CalledNum) {
		query["CalledNum"] = request.CalledNum
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCallTaskDetail"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCallTaskDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of call tasks based on task IDs after call tasks are complete.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - ListCallTaskDetailRequest
//
// @return ListCallTaskDetailResponse
func (client *Client) ListCallTaskDetail(request *ListCallTaskDetailRequest) (_result *ListCallTaskDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListCallTaskDetailResponse{}
	_body, _err := client.ListCallTaskDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the registration information about a China 400 number.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - ListHotlineTransferRegisterFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListHotlineTransferRegisterFileResponse
func (client *Client) ListHotlineTransferRegisterFileWithOptions(request *ListHotlineTransferRegisterFileRequest, runtime *dara.RuntimeOptions) (_result *ListHotlineTransferRegisterFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.HotlineNumber) {
		query["HotlineNumber"] = request.HotlineNumber
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.QualificationId) {
		query["QualificationId"] = request.QualificationId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListHotlineTransferRegisterFile"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListHotlineTransferRegisterFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the registration information about a China 400 number.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - ListHotlineTransferRegisterFileRequest
//
// @return ListHotlineTransferRegisterFileResponse
func (client *Client) ListHotlineTransferRegisterFile(request *ListHotlineTransferRegisterFileRequest) (_result *ListHotlineTransferRegisterFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListHotlineTransferRegisterFileResponse{}
	_body, _err := client.ListHotlineTransferRegisterFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询服务实例列表
//
// @param tmpReq - ListServiceInstanceForPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServiceInstanceForPageResponse
func (client *Client) ListServiceInstanceForPageWithOptions(tmpReq *ListServiceInstanceForPageRequest, runtime *dara.RuntimeOptions) (_result *ListServiceInstanceForPageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListServiceInstanceForPageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Pager) {
		request.PagerShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Pager, dara.String("Pager"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		query["Code"] = request.Code
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PagerShrink) {
		query["Pager"] = request.PagerShrink
	}

	if !dara.IsNil(request.RelationNumber) {
		query["RelationNumber"] = request.RelationNumber
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SceneId) {
		query["SceneId"] = request.SceneId
	}

	if !dara.IsNil(request.UsageId) {
		query["UsageId"] = request.UsageId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServiceInstanceForPage"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServiceInstanceForPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询服务实例列表
//
// @param request - ListServiceInstanceForPageRequest
//
// @return ListServiceInstanceForPageResponse
func (client *Client) ListServiceInstanceForPage(request *ListServiceInstanceForPageRequest) (_result *ListServiceInstanceForPageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListServiceInstanceForPageResponse{}
	_body, _err := client.ListServiceInstanceForPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Pauses video playback when a video file is played back during a voice call.
//
// @param request - PauseVideoFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PauseVideoFileResponse
func (client *Client) PauseVideoFileWithOptions(request *PauseVideoFileRequest, runtime *dara.RuntimeOptions) (_result *PauseVideoFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallId) {
		query["CallId"] = request.CallId
	}

	if !dara.IsNil(request.CalledNumber) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PauseVideoFile"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PauseVideoFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Pauses video playback when a video file is played back during a voice call.
//
// @param request - PauseVideoFileRequest
//
// @return PauseVideoFileResponse
func (client *Client) PauseVideoFile(request *PauseVideoFileRequest) (_result *PauseVideoFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PauseVideoFileResponse{}
	_body, _err := client.PauseVideoFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Plays back a video file during a voice call.
//
// @param request - PlayVideoFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PlayVideoFileResponse
func (client *Client) PlayVideoFileWithOptions(request *PlayVideoFileRequest, runtime *dara.RuntimeOptions) (_result *PlayVideoFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallId) {
		query["CallId"] = request.CallId
	}

	if !dara.IsNil(request.CalledNumber) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !dara.IsNil(request.OnlyPhone) {
		query["OnlyPhone"] = request.OnlyPhone
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.VideoId) {
		query["VideoId"] = request.VideoId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PlayVideoFile"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PlayVideoFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Plays back a video file during a voice call.
//
// @param request - PlayVideoFileRequest
//
// @return PlayVideoFileResponse
func (client *Client) PlayVideoFile(request *PlayVideoFileRequest) (_result *PlayVideoFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PlayVideoFileResponse{}
	_body, _err := client.PlayVideoFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a call.
//
// Description:
//
// QueryCallDetailByCallId is a common query operation. You can call this operation to query the details of a voice notification, voice verification code, interactive voice response (IVR), intelligent inbound voice call, intelligent outbound voice call, or intelligent robocall.
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - QueryCallDetailByCallIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCallDetailByCallIdResponse
func (client *Client) QueryCallDetailByCallIdWithOptions(request *QueryCallDetailByCallIdRequest, runtime *dara.RuntimeOptions) (_result *QueryCallDetailByCallIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallId) {
		query["CallId"] = request.CallId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ProdId) {
		query["ProdId"] = request.ProdId
	}

	if !dara.IsNil(request.QueryDate) {
		query["QueryDate"] = request.QueryDate
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCallDetailByCallId"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCallDetailByCallIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a call.
//
// Description:
//
// QueryCallDetailByCallId is a common query operation. You can call this operation to query the details of a voice notification, voice verification code, interactive voice response (IVR), intelligent inbound voice call, intelligent outbound voice call, or intelligent robocall.
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - QueryCallDetailByCallIdRequest
//
// @return QueryCallDetailByCallIdResponse
func (client *Client) QueryCallDetailByCallId(request *QueryCallDetailByCallIdRequest) (_result *QueryCallDetailByCallIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryCallDetailByCallIdResponse{}
	_body, _err := client.QueryCallDetailByCallIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the call details of an outbound robocall task.
//
// @param request - QueryCallDetailByTaskIdRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCallDetailByTaskIdResponse
func (client *Client) QueryCallDetailByTaskIdWithOptions(request *QueryCallDetailByTaskIdRequest, runtime *dara.RuntimeOptions) (_result *QueryCallDetailByTaskIdResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Callee) {
		query["Callee"] = request.Callee
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.QueryDate) {
		query["QueryDate"] = request.QueryDate
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCallDetailByTaskId"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCallDetailByTaskIdResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the call details of an outbound robocall task.
//
// @param request - QueryCallDetailByTaskIdRequest
//
// @return QueryCallDetailByTaskIdResponse
func (client *Client) QueryCallDetailByTaskId(request *QueryCallDetailByTaskIdRequest) (_result *QueryCallDetailByTaskIdResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryCallDetailByTaskIdResponse{}
	_body, _err := client.QueryCallDetailByTaskIdWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the configuration of the phone number used to transfer a call.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - QueryCallInPoolTransferConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCallInPoolTransferConfigResponse
func (client *Client) QueryCallInPoolTransferConfigWithOptions(request *QueryCallInPoolTransferConfigRequest, runtime *dara.RuntimeOptions) (_result *QueryCallInPoolTransferConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCallInPoolTransferConfig"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCallInPoolTransferConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the configuration of the phone number used to transfer a call.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - QueryCallInPoolTransferConfigRequest
//
// @return QueryCallInPoolTransferConfigResponse
func (client *Client) QueryCallInPoolTransferConfig(request *QueryCallInPoolTransferConfigRequest) (_result *QueryCallInPoolTransferConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryCallInPoolTransferConfigResponse{}
	_body, _err := client.QueryCallInPoolTransferConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries call transfer records.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - QueryCallInTransferRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryCallInTransferRecordResponse
func (client *Client) QueryCallInTransferRecordWithOptions(request *QueryCallInTransferRecordRequest, runtime *dara.RuntimeOptions) (_result *QueryCallInTransferRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallInCaller) {
		query["CallInCaller"] = request.CallInCaller
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.QueryDate) {
		query["QueryDate"] = request.QueryDate
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryCallInTransferRecord"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryCallInTransferRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries call transfer records.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - QueryCallInTransferRecordRequest
//
// @return QueryCallInTransferRecordResponse
func (client *Client) QueryCallInTransferRecord(request *QueryCallInTransferRecordRequest) (_result *QueryCallInTransferRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryCallInTransferRecordResponse{}
	_body, _err := client.QueryCallInTransferRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of robots to obtain their details.
//
// @param request - QueryRobotInfoListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRobotInfoListResponse
func (client *Client) QueryRobotInfoListWithOptions(request *QueryRobotInfoListRequest, runtime *dara.RuntimeOptions) (_result *QueryRobotInfoListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AuditStatus) {
		query["AuditStatus"] = request.AuditStatus
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRobotInfoList"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRobotInfoListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of robots to obtain their details.
//
// @param request - QueryRobotInfoListRequest
//
// @return QueryRobotInfoListResponse
func (client *Client) QueryRobotInfoList(request *QueryRobotInfoListRequest) (_result *QueryRobotInfoListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryRobotInfoListResponse{}
	_body, _err := client.QueryRobotInfoListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the call details of a called number.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - QueryRobotTaskCallDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRobotTaskCallDetailResponse
func (client *Client) QueryRobotTaskCallDetailWithOptions(request *QueryRobotTaskCallDetailRequest, runtime *dara.RuntimeOptions) (_result *QueryRobotTaskCallDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Callee) {
		query["Callee"] = request.Callee
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.QueryDate) {
		query["QueryDate"] = request.QueryDate
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRobotTaskCallDetail"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRobotTaskCallDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the call details of a called number.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - QueryRobotTaskCallDetailRequest
//
// @return QueryRobotTaskCallDetailResponse
func (client *Client) QueryRobotTaskCallDetail(request *QueryRobotTaskCallDetailRequest) (_result *QueryRobotTaskCallDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryRobotTaskCallDetailResponse{}
	_body, _err := client.QueryRobotTaskCallDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a robocall task.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - QueryRobotTaskCallListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRobotTaskCallListResponse
func (client *Client) QueryRobotTaskCallListWithOptions(request *QueryRobotTaskCallListRequest, runtime *dara.RuntimeOptions) (_result *QueryRobotTaskCallListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallResult) {
		query["CallResult"] = request.CallResult
	}

	if !dara.IsNil(request.Called) {
		query["Called"] = request.Called
	}

	if !dara.IsNil(request.DialogCountFrom) {
		query["DialogCountFrom"] = request.DialogCountFrom
	}

	if !dara.IsNil(request.DialogCountTo) {
		query["DialogCountTo"] = request.DialogCountTo
	}

	if !dara.IsNil(request.DurationFrom) {
		query["DurationFrom"] = request.DurationFrom
	}

	if !dara.IsNil(request.DurationTo) {
		query["DurationTo"] = request.DurationTo
	}

	if !dara.IsNil(request.HangupDirection) {
		query["HangupDirection"] = request.HangupDirection
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRobotTaskCallList"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRobotTaskCallListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a robocall task.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - QueryRobotTaskCallListRequest
//
// @return QueryRobotTaskCallListResponse
func (client *Client) QueryRobotTaskCallList(request *QueryRobotTaskCallListRequest) (_result *QueryRobotTaskCallListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryRobotTaskCallListResponse{}
	_body, _err := client.QueryRobotTaskCallListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a robocall task.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - QueryRobotTaskDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRobotTaskDetailResponse
func (client *Client) QueryRobotTaskDetailWithOptions(request *QueryRobotTaskDetailRequest, runtime *dara.RuntimeOptions) (_result *QueryRobotTaskDetailResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRobotTaskDetail"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRobotTaskDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a robocall task.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - QueryRobotTaskDetailRequest
//
// @return QueryRobotTaskDetailResponse
func (client *Client) QueryRobotTaskDetail(request *QueryRobotTaskDetailRequest) (_result *QueryRobotTaskDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryRobotTaskDetailResponse{}
	_body, _err := client.QueryRobotTaskDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about all robocall tasks.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - QueryRobotTaskListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRobotTaskListResponse
func (client *Client) QueryRobotTaskListWithOptions(request *QueryRobotTaskListRequest, runtime *dara.RuntimeOptions) (_result *QueryRobotTaskListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskName) {
		query["TaskName"] = request.TaskName
	}

	if !dara.IsNil(request.Time) {
		query["Time"] = request.Time
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRobotTaskList"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRobotTaskListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about all robocall tasks.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - QueryRobotTaskListRequest
//
// @return QueryRobotTaskListResponse
func (client *Client) QueryRobotTaskList(request *QueryRobotTaskListRequest) (_result *QueryRobotTaskListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryRobotTaskListResponse{}
	_body, _err := client.QueryRobotTaskListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of robot communication scripts.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - QueryRobotv2AllListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRobotv2AllListResponse
func (client *Client) QueryRobotv2AllListWithOptions(request *QueryRobotv2AllListRequest, runtime *dara.RuntimeOptions) (_result *QueryRobotv2AllListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRobotv2AllList"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRobotv2AllListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of robot communication scripts.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - QueryRobotv2AllListRequest
//
// @return QueryRobotv2AllListResponse
func (client *Client) QueryRobotv2AllList(request *QueryRobotv2AllListRequest) (_result *QueryRobotv2AllListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryRobotv2AllListResponse{}
	_body, _err := client.QueryRobotv2AllListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the video playback progress after you play a video file during a voice call.
//
// @param request - QueryVideoPlayProgressRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryVideoPlayProgressResponse
func (client *Client) QueryVideoPlayProgressWithOptions(request *QueryVideoPlayProgressRequest, runtime *dara.RuntimeOptions) (_result *QueryVideoPlayProgressResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallId) {
		query["CallId"] = request.CallId
	}

	if !dara.IsNil(request.CalledNumber) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryVideoPlayProgress"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryVideoPlayProgressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the video playback progress after you play a video file during a voice call.
//
// @param request - QueryVideoPlayProgressRequest
//
// @return QueryVideoPlayProgressResponse
func (client *Client) QueryVideoPlayProgress(request *QueryVideoPlayProgressRequest) (_result *QueryVideoPlayProgressResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryVideoPlayProgressResponse{}
	_body, _err := client.QueryVideoPlayProgressWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists real numbers bound to service instances. The returned data includes the binding time, the number activation time, and the number of real numbers bound to a service instance.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 200 times per second per account.
//
// @param request - QueryVirtualNumberRelationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryVirtualNumberRelationResponse
func (client *Client) QueryVirtualNumberRelationWithOptions(request *QueryVirtualNumberRelationRequest, runtime *dara.RuntimeOptions) (_result *QueryVirtualNumberRelationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PhoneNum) {
		query["PhoneNum"] = request.PhoneNum
	}

	if !dara.IsNil(request.ProdCode) {
		query["ProdCode"] = request.ProdCode
	}

	if !dara.IsNil(request.QualificationId) {
		query["QualificationId"] = request.QualificationId
	}

	if !dara.IsNil(request.RegionNameCity) {
		query["RegionNameCity"] = request.RegionNameCity
	}

	if !dara.IsNil(request.RelatedNum) {
		query["RelatedNum"] = request.RelatedNum
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.RouteType) {
		query["RouteType"] = request.RouteType
	}

	if !dara.IsNil(request.SpecId) {
		query["SpecId"] = request.SpecId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryVirtualNumberRelation"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryVirtualNumberRelationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists real numbers bound to service instances. The returned data includes the binding time, the number activation time, and the number of real numbers bound to a service instance.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 200 times per second per account.
//
// @param request - QueryVirtualNumberRelationRequest
//
// @return QueryVirtualNumberRelationResponse
func (client *Client) QueryVirtualNumberRelation(request *QueryVirtualNumberRelationRequest) (_result *QueryVirtualNumberRelationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryVirtualNumberRelationResponse{}
	_body, _err := client.QueryVirtualNumberRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询真实号接通率
//
// @param request - QueryVmsRealNumberCallConnectionRateInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryVmsRealNumberCallConnectionRateInfoResponse
func (client *Client) QueryVmsRealNumberCallConnectionRateInfoWithOptions(request *QueryVmsRealNumberCallConnectionRateInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryVmsRealNumberCallConnectionRateInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.RealNumber) {
		query["RealNumber"] = request.RealNumber
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TimePeriod) {
		query["TimePeriod"] = request.TimePeriod
	}

	if !dara.IsNil(request.VirtualNumber) {
		query["VirtualNumber"] = request.VirtualNumber
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryVmsRealNumberCallConnectionRateInfo"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryVmsRealNumberCallConnectionRateInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询真实号接通率
//
// @param request - QueryVmsRealNumberCallConnectionRateInfoRequest
//
// @return QueryVmsRealNumberCallConnectionRateInfoResponse
func (client *Client) QueryVmsRealNumberCallConnectionRateInfo(request *QueryVmsRealNumberCallConnectionRateInfoRequest) (_result *QueryVmsRealNumberCallConnectionRateInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryVmsRealNumberCallConnectionRateInfoResponse{}
	_body, _err := client.QueryVmsRealNumberCallConnectionRateInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询虚拟号码与真实号码绑定关系列表
//
// @param request - QueryVmsVirtualNumberRelationByPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryVmsVirtualNumberRelationByPageResponse
func (client *Client) QueryVmsVirtualNumberRelationByPageWithOptions(request *QueryVmsVirtualNumberRelationByPageRequest, runtime *dara.RuntimeOptions) (_result *QueryVmsVirtualNumberRelationByPageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NumberCity) {
		query["NumberCity"] = request.NumberCity
	}

	if !dara.IsNil(request.NumberProvince) {
		query["NumberProvince"] = request.NumberProvince
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNo) {
		query["PageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RealNumber) {
		query["RealNumber"] = request.RealNumber
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.State) {
		query["State"] = request.State
	}

	if !dara.IsNil(request.VirtualNumber) {
		query["VirtualNumber"] = request.VirtualNumber
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryVmsVirtualNumberRelationByPage"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryVmsVirtualNumberRelationByPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询虚拟号码与真实号码绑定关系列表
//
// @param request - QueryVmsVirtualNumberRelationByPageRequest
//
// @return QueryVmsVirtualNumberRelationByPageResponse
func (client *Client) QueryVmsVirtualNumberRelationByPage(request *QueryVmsVirtualNumberRelationByPageRequest) (_result *QueryVmsVirtualNumberRelationByPageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryVmsVirtualNumberRelationByPageResponse{}
	_body, _err := client.QueryVmsVirtualNumberRelationByPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the review state of a voice file.
//
// @param request - QueryVoiceFileAuditInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryVoiceFileAuditInfoResponse
func (client *Client) QueryVoiceFileAuditInfoWithOptions(request *QueryVoiceFileAuditInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryVoiceFileAuditInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BusinessType) {
		query["BusinessType"] = request.BusinessType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.VoiceCodes) {
		query["VoiceCodes"] = request.VoiceCodes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryVoiceFileAuditInfo"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryVoiceFileAuditInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the review state of a voice file.
//
// @param request - QueryVoiceFileAuditInfoRequest
//
// @return QueryVoiceFileAuditInfoResponse
func (client *Client) QueryVoiceFileAuditInfo(request *QueryVoiceFileAuditInfoRequest) (_result *QueryVoiceFileAuditInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryVoiceFileAuditInfoResponse{}
	_body, _err := client.QueryVoiceFileAuditInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resumes the inbound call that is transferred by using a China 400 number.
//
// @param request - RecoverCallInConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecoverCallInConfigResponse
func (client *Client) RecoverCallInConfigWithOptions(request *RecoverCallInConfigRequest, runtime *dara.RuntimeOptions) (_result *RecoverCallInConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Number) {
		query["Number"] = request.Number
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RecoverCallInConfig"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RecoverCallInConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resumes the inbound call that is transferred by using a China 400 number.
//
// @param request - RecoverCallInConfigRequest
//
// @return RecoverCallInConfigResponse
func (client *Client) RecoverCallInConfig(request *RecoverCallInConfigRequest) (_result *RecoverCallInConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RecoverCallInConfigResponse{}
	_body, _err := client.RecoverCallInConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resumes video playback after you pause video playback during a voice call.
//
// @param request - ResumeVideoFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeVideoFileResponse
func (client *Client) ResumeVideoFileWithOptions(request *ResumeVideoFileRequest, runtime *dara.RuntimeOptions) (_result *ResumeVideoFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallId) {
		query["CallId"] = request.CallId
	}

	if !dara.IsNil(request.CalledNumber) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResumeVideoFile"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ResumeVideoFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resumes video playback after you pause video playback during a voice call.
//
// @param request - ResumeVideoFileRequest
//
// @return ResumeVideoFileResponse
func (client *Client) ResumeVideoFile(request *ResumeVideoFileRequest) (_result *ResumeVideoFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ResumeVideoFileResponse{}
	_body, _err := client.ResumeVideoFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # SeekVideoFile
//
// @param request - SeekVideoFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SeekVideoFileResponse
func (client *Client) SeekVideoFileWithOptions(request *SeekVideoFileRequest, runtime *dara.RuntimeOptions) (_result *SeekVideoFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallId) {
		query["CallId"] = request.CallId
	}

	if !dara.IsNil(request.CalledNumber) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SeekTimes) {
		query["SeekTimes"] = request.SeekTimes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SeekVideoFile"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SeekVideoFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # SeekVideoFile
//
// @param request - SeekVideoFileRequest
//
// @return SeekVideoFileResponse
func (client *Client) SeekVideoFile(request *SeekVideoFileRequest) (_result *SeekVideoFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SeekVideoFileResponse{}
	_body, _err := client.SeekVideoFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sends an SMS verification code.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - SendVerificationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendVerificationResponse
func (client *Client) SendVerificationWithOptions(request *SendVerificationRequest, runtime *dara.RuntimeOptions) (_result *SendVerificationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizType) {
		query["BizType"] = request.BizType
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Target) {
		query["Target"] = request.Target
	}

	if !dara.IsNil(request.VerifyType) {
		query["VerifyType"] = request.VerifyType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendVerification"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendVerificationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends an SMS verification code.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - SendVerificationRequest
//
// @return SendVerificationResponse
func (client *Client) SendVerification(request *SendVerificationRequest) (_result *SendVerificationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SendVerificationResponse{}
	_body, _err := client.SendVerificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sets the phone numbers for transferring a call.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - SetTransferCalleePoolConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetTransferCalleePoolConfigResponse
func (client *Client) SetTransferCalleePoolConfigWithOptions(request *SetTransferCalleePoolConfigRequest, runtime *dara.RuntimeOptions) (_result *SetTransferCalleePoolConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CalledRouteMode) {
		query["CalledRouteMode"] = request.CalledRouteMode
	}

	if !dara.IsNil(request.Details) {
		query["Details"] = request.Details
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PhoneNumber) {
		query["PhoneNumber"] = request.PhoneNumber
	}

	if !dara.IsNil(request.QualificationId) {
		query["QualificationId"] = request.QualificationId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetTransferCalleePoolConfig"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetTransferCalleePoolConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sets the phone numbers for transferring a call.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - SetTransferCalleePoolConfigRequest
//
// @return SetTransferCalleePoolConfigResponse
func (client *Client) SetTransferCalleePoolConfig(request *SetTransferCalleePoolConfigRequest) (_result *SetTransferCalleePoolConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetTransferCalleePoolConfigResponse{}
	_body, _err := client.SetTransferCalleePoolConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sends a voice verification code or a voice notification with variables to a specified phone number.
//
// Description:
//
//	  Due to business adjustments, the updates of the voice notification and voice verification code services have been stopped in regions outside the Chinese mainland and the services have been discontinued since March 2022. Only qualified customers can continue using the voice notification and voice verification code services.
//
//		- For more information about voice plans or voice service billing, see [Pricing of VMS on China site (aliyun.com)](https://help.aliyun.com/document_detail/150083.html).
//
// ### QPS limits
//
// You can call this operation up to 1,000 times per second per account.
//
// @param request - SingleCallByTtsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SingleCallByTtsResponse
func (client *Client) SingleCallByTtsWithOptions(request *SingleCallByTtsRequest, runtime *dara.RuntimeOptions) (_result *SingleCallByTtsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CalledNumber) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !dara.IsNil(request.CalledShowNumber) {
		query["CalledShowNumber"] = request.CalledShowNumber
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PlayTimes) {
		query["PlayTimes"] = request.PlayTimes
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Speed) {
		query["Speed"] = request.Speed
	}

	if !dara.IsNil(request.TtsCode) {
		query["TtsCode"] = request.TtsCode
	}

	if !dara.IsNil(request.TtsParam) {
		query["TtsParam"] = request.TtsParam
	}

	if !dara.IsNil(request.Volume) {
		query["Volume"] = request.Volume
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SingleCallByTts"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SingleCallByTtsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends a voice verification code or a voice notification with variables to a specified phone number.
//
// Description:
//
//	  Due to business adjustments, the updates of the voice notification and voice verification code services have been stopped in regions outside the Chinese mainland and the services have been discontinued since March 2022. Only qualified customers can continue using the voice notification and voice verification code services.
//
//		- For more information about voice plans or voice service billing, see [Pricing of VMS on China site (aliyun.com)](https://help.aliyun.com/document_detail/150083.html).
//
// ### QPS limits
//
// You can call this operation up to 1,000 times per second per account.
//
// @param request - SingleCallByTtsRequest
//
// @return SingleCallByTtsResponse
func (client *Client) SingleCallByTts(request *SingleCallByTtsRequest) (_result *SingleCallByTtsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SingleCallByTtsResponse{}
	_body, _err := client.SingleCallByTtsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sends voice file notifications or video file notifications to a single called number.
//
// @param request - SingleCallByVideoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SingleCallByVideoResponse
func (client *Client) SingleCallByVideoWithOptions(request *SingleCallByVideoRequest, runtime *dara.RuntimeOptions) (_result *SingleCallByVideoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CalledNumber) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !dara.IsNil(request.CalledShowNumber) {
		query["CalledShowNumber"] = request.CalledShowNumber
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PlayTimes) {
		query["PlayTimes"] = request.PlayTimes
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Speed) {
		query["Speed"] = request.Speed
	}

	if !dara.IsNil(request.VideoCode) {
		query["VideoCode"] = request.VideoCode
	}

	if !dara.IsNil(request.VoiceCode) {
		query["VoiceCode"] = request.VoiceCode
	}

	if !dara.IsNil(request.Volume) {
		query["Volume"] = request.Volume
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SingleCallByVideo"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SingleCallByVideoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends voice file notifications or video file notifications to a single called number.
//
// @param request - SingleCallByVideoRequest
//
// @return SingleCallByVideoResponse
func (client *Client) SingleCallByVideo(request *SingleCallByVideoRequest) (_result *SingleCallByVideoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SingleCallByVideoResponse{}
	_body, _err := client.SingleCallByVideoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Sends a voice notification to a phone number by using a voice notification file.
//
// Description:
//
// > Due to business adjustments, the updates of the voice notification and voice verification code services have been stopped in regions outside the Chinese mainland and the services have been discontinued since March 2022. Only qualified customers can continue using the voice notification and voice verification code services.
//
// You can call the [SingleCallByTts](https://help.aliyun.com/document_detail/393519.html) operation to send voice notifications with variables.
//
// ### QPS limits
//
// You can call this operation up to 1,200 times per second per account.
//
// @param request - SingleCallByVoiceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SingleCallByVoiceResponse
func (client *Client) SingleCallByVoiceWithOptions(request *SingleCallByVoiceRequest, runtime *dara.RuntimeOptions) (_result *SingleCallByVoiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CalledNumber) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !dara.IsNil(request.CalledShowNumber) {
		query["CalledShowNumber"] = request.CalledShowNumber
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PlayTimes) {
		query["PlayTimes"] = request.PlayTimes
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.Speed) {
		query["Speed"] = request.Speed
	}

	if !dara.IsNil(request.VoiceCode) {
		query["VoiceCode"] = request.VoiceCode
	}

	if !dara.IsNil(request.Volume) {
		query["Volume"] = request.Volume
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SingleCallByVoice"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SingleCallByVoiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Sends a voice notification to a phone number by using a voice notification file.
//
// Description:
//
// > Due to business adjustments, the updates of the voice notification and voice verification code services have been stopped in regions outside the Chinese mainland and the services have been discontinued since March 2022. Only qualified customers can continue using the voice notification and voice verification code services.
//
// You can call the [SingleCallByTts](https://help.aliyun.com/document_detail/393519.html) operation to send voice notifications with variables.
//
// ### QPS limits
//
// You can call this operation up to 1,200 times per second per account.
//
// @param request - SingleCallByVoiceRequest
//
// @return SingleCallByVoiceResponse
func (client *Client) SingleCallByVoice(request *SingleCallByVoiceRequest) (_result *SingleCallByVoiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SingleCallByVoiceResponse{}
	_body, _err := client.SingleCallByVoiceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Fast forwards or rewinds a video when you play the video.
//
// @param request - SkipVideoFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SkipVideoFileResponse
func (client *Client) SkipVideoFileWithOptions(request *SkipVideoFileRequest, runtime *dara.RuntimeOptions) (_result *SkipVideoFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallId) {
		query["CallId"] = request.CallId
	}

	if !dara.IsNil(request.CalledNumber) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SkipTimes) {
		query["SkipTimes"] = request.SkipTimes
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SkipVideoFile"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SkipVideoFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Fast forwards or rewinds a video when you play the video.
//
// @param request - SkipVideoFileRequest
//
// @return SkipVideoFileResponse
func (client *Client) SkipVideoFile(request *SkipVideoFileRequest) (_result *SkipVideoFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SkipVideoFileResponse{}
	_body, _err := client.SkipVideoFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Initiates an intelligent voice call.
//
// Description:
//
//	  The SmartCall operation must be used together with the [intelligent outbound HTTP operation](https://help.aliyun.com/document_detail/112703.html). After the call initiated by the Voice Messaging Service (VMS) platform is connected, the VMS platform sends the text converted from speech back to the business side, and the business side then returns the follow-up action to the VMS platform.
//
//		- The SmartCall operation does not support the following characters: `@ = : "" $ { } ^ 	- ￥`.
//
// ### QPS limits
//
// You can call this operation up to 1,000 times per second per account.
//
// @param request - SmartCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SmartCallResponse
func (client *Client) SmartCallWithOptions(request *SmartCallRequest, runtime *dara.RuntimeOptions) (_result *SmartCallResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ActionCodeBreak) {
		query["ActionCodeBreak"] = request.ActionCodeBreak
	}

	if !dara.IsNil(request.ActionCodeTimeBreak) {
		query["ActionCodeTimeBreak"] = request.ActionCodeTimeBreak
	}

	if !dara.IsNil(request.AsrBaseId) {
		query["AsrBaseId"] = request.AsrBaseId
	}

	if !dara.IsNil(request.AsrModelId) {
		query["AsrModelId"] = request.AsrModelId
	}

	if !dara.IsNil(request.BackgroundFileCode) {
		query["BackgroundFileCode"] = request.BackgroundFileCode
	}

	if !dara.IsNil(request.BackgroundSpeed) {
		query["BackgroundSpeed"] = request.BackgroundSpeed
	}

	if !dara.IsNil(request.BackgroundVolume) {
		query["BackgroundVolume"] = request.BackgroundVolume
	}

	if !dara.IsNil(request.CalledNumber) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !dara.IsNil(request.CalledShowNumber) {
		query["CalledShowNumber"] = request.CalledShowNumber
	}

	if !dara.IsNil(request.DynamicId) {
		query["DynamicId"] = request.DynamicId
	}

	if !dara.IsNil(request.EarlyMediaAsr) {
		query["EarlyMediaAsr"] = request.EarlyMediaAsr
	}

	if !dara.IsNil(request.EnableITN) {
		query["EnableITN"] = request.EnableITN
	}

	if !dara.IsNil(request.MuteTime) {
		query["MuteTime"] = request.MuteTime
	}

	if !dara.IsNil(request.NoiseThreshold) {
		query["NoiseThreshold"] = request.NoiseThreshold
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PauseTime) {
		query["PauseTime"] = request.PauseTime
	}

	if !dara.IsNil(request.RecordFlag) {
		query["RecordFlag"] = request.RecordFlag
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SessionTimeout) {
		query["SessionTimeout"] = request.SessionTimeout
	}

	if !dara.IsNil(request.Speed) {
		query["Speed"] = request.Speed
	}

	if !dara.IsNil(request.StreamAsr) {
		query["StreamAsr"] = request.StreamAsr
	}

	if !dara.IsNil(request.TtsConf) {
		query["TtsConf"] = request.TtsConf
	}

	if !dara.IsNil(request.TtsSpeed) {
		query["TtsSpeed"] = request.TtsSpeed
	}

	if !dara.IsNil(request.TtsStyle) {
		query["TtsStyle"] = request.TtsStyle
	}

	if !dara.IsNil(request.TtsVolume) {
		query["TtsVolume"] = request.TtsVolume
	}

	if !dara.IsNil(request.VoiceCode) {
		query["VoiceCode"] = request.VoiceCode
	}

	if !dara.IsNil(request.VoiceCodeParam) {
		query["VoiceCodeParam"] = request.VoiceCodeParam
	}

	if !dara.IsNil(request.Volume) {
		query["Volume"] = request.Volume
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SmartCall"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SmartCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiates an intelligent voice call.
//
// Description:
//
//	  The SmartCall operation must be used together with the [intelligent outbound HTTP operation](https://help.aliyun.com/document_detail/112703.html). After the call initiated by the Voice Messaging Service (VMS) platform is connected, the VMS platform sends the text converted from speech back to the business side, and the business side then returns the follow-up action to the VMS platform.
//
//		- The SmartCall operation does not support the following characters: `@ = : "" $ { } ^ 	- ￥`.
//
// ### QPS limits
//
// You can call this operation up to 1,000 times per second per account.
//
// @param request - SmartCallRequest
//
// @return SmartCallResponse
func (client *Client) SmartCall(request *SmartCallRequest) (_result *SmartCallResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SmartCallResponse{}
	_body, _err := client.SmartCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Initiates an action in an outbound robocall. This operation is applicable only when the robocall is transferred to an agent or an agent is listening in on the conversation between the robot and the user.
//
// Description:
//
// You can call this operation to initiate a specified action on the called number of an outbound robocall when the call is transferred to an agent of the call center.
//
// > You can only initiate the action of bridging a called number and an agent of the call center.
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - SmartCallOperateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SmartCallOperateResponse
func (client *Client) SmartCallOperateWithOptions(request *SmartCallOperateRequest, runtime *dara.RuntimeOptions) (_result *SmartCallOperateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallId) {
		query["CallId"] = request.CallId
	}

	if !dara.IsNil(request.Command) {
		query["Command"] = request.Command
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Param) {
		query["Param"] = request.Param
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SmartCallOperate"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SmartCallOperateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiates an action in an outbound robocall. This operation is applicable only when the robocall is transferred to an agent or an agent is listening in on the conversation between the robot and the user.
//
// Description:
//
// You can call this operation to initiate a specified action on the called number of an outbound robocall when the call is transferred to an agent of the call center.
//
// > You can only initiate the action of bridging a called number and an agent of the call center.
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - SmartCallOperateRequest
//
// @return SmartCallOperateResponse
func (client *Client) SmartCallOperate(request *SmartCallOperateRequest) (_result *SmartCallOperateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SmartCallOperateResponse{}
	_body, _err := client.SmartCallOperateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts a robocall task immediately or at a scheduled time.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - StartRobotTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartRobotTaskResponse
func (client *Client) StartRobotTaskWithOptions(request *StartRobotTaskRequest, runtime *dara.RuntimeOptions) (_result *StartRobotTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.ScheduleTime) {
		query["ScheduleTime"] = request.ScheduleTime
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartRobotTask"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StartRobotTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a robocall task immediately or at a scheduled time.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - StartRobotTaskRequest
//
// @return StartRobotTaskResponse
func (client *Client) StartRobotTask(request *StartRobotTaskRequest) (_result *StartRobotTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StartRobotTaskResponse{}
	_body, _err := client.StartRobotTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops the inbound call that is transferred from a China 400 number.
//
// @param request - StopCallInConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopCallInConfigResponse
func (client *Client) StopCallInConfigWithOptions(request *StopCallInConfigRequest, runtime *dara.RuntimeOptions) (_result *StopCallInConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Number) {
		query["Number"] = request.Number
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopCallInConfig"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopCallInConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops the inbound call that is transferred from a China 400 number.
//
// @param request - StopCallInConfigRequest
//
// @return StopCallInConfigResponse
func (client *Client) StopCallInConfig(request *StopCallInConfigRequest) (_result *StopCallInConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopCallInConfigResponse{}
	_body, _err := client.StopCallInConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a robocall task that is in progress.
//
// Description:
//
// After you stop a robocall task, you can call the [StartRobotTask](~~StartRobotTask~~) operation to start it again.
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - StopRobotTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopRobotTaskResponse
func (client *Client) StopRobotTaskWithOptions(request *StopRobotTaskRequest, runtime *dara.RuntimeOptions) (_result *StopRobotTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopRobotTask"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &StopRobotTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a robocall task that is in progress.
//
// Description:
//
// After you stop a robocall task, you can call the [StartRobotTask](~~StartRobotTask~~) operation to start it again.
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - StopRobotTaskRequest
//
// @return StopRobotTaskResponse
func (client *Client) StopRobotTask(request *StopRobotTaskRequest) (_result *StopRobotTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &StopRobotTaskResponse{}
	_body, _err := client.StopRobotTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Submits a 400 number for registration.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - SubmitHotlineTransferRegisterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitHotlineTransferRegisterResponse
func (client *Client) SubmitHotlineTransferRegisterWithOptions(request *SubmitHotlineTransferRegisterRequest, runtime *dara.RuntimeOptions) (_result *SubmitHotlineTransferRegisterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Agreement) {
		query["Agreement"] = request.Agreement
	}

	if !dara.IsNil(request.HotlineNumber) {
		query["HotlineNumber"] = request.HotlineNumber
	}

	if !dara.IsNil(request.OperatorIdentityCard) {
		query["OperatorIdentityCard"] = request.OperatorIdentityCard
	}

	if !dara.IsNil(request.OperatorMail) {
		query["OperatorMail"] = request.OperatorMail
	}

	if !dara.IsNil(request.OperatorMailVerifyCode) {
		query["OperatorMailVerifyCode"] = request.OperatorMailVerifyCode
	}

	if !dara.IsNil(request.OperatorMobile) {
		query["OperatorMobile"] = request.OperatorMobile
	}

	if !dara.IsNil(request.OperatorMobileVerifyCode) {
		query["OperatorMobileVerifyCode"] = request.OperatorMobileVerifyCode
	}

	if !dara.IsNil(request.OperatorName) {
		query["OperatorName"] = request.OperatorName
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.QualificationId) {
		query["QualificationId"] = request.QualificationId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TransferPhoneNumberInfos) {
		query["TransferPhoneNumberInfos"] = request.TransferPhoneNumberInfos
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitHotlineTransferRegister"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitHotlineTransferRegisterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Submits a 400 number for registration.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - SubmitHotlineTransferRegisterRequest
//
// @return SubmitHotlineTransferRegisterResponse
func (client *Client) SubmitHotlineTransferRegister(request *SubmitHotlineTransferRegisterRequest) (_result *SubmitHotlineTransferRegisterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SubmitHotlineTransferRegisterResponse{}
	_body, _err := client.SubmitHotlineTransferRegisterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Upgrades from a voice call to a video call.
//
// @param request - UpgradeVideoFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpgradeVideoFileResponse
func (client *Client) UpgradeVideoFileWithOptions(request *UpgradeVideoFileRequest, runtime *dara.RuntimeOptions) (_result *UpgradeVideoFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallId) {
		query["CallId"] = request.CallId
	}

	if !dara.IsNil(request.CalledNumber) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !dara.IsNil(request.MediaType) {
		query["MediaType"] = request.MediaType
	}

	if !dara.IsNil(request.OutId) {
		query["OutId"] = request.OutId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpgradeVideoFile"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpgradeVideoFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Upgrades from a voice call to a video call.
//
// @param request - UpgradeVideoFileRequest
//
// @return UpgradeVideoFileResponse
func (client *Client) UpgradeVideoFile(request *UpgradeVideoFileRequest) (_result *UpgradeVideoFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpgradeVideoFileResponse{}
	_body, _err := client.UpgradeVideoFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Uploads the called numbers of a robocall task.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - UploadRobotTaskCalledFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadRobotTaskCalledFileResponse
func (client *Client) UploadRobotTaskCalledFileWithOptions(request *UploadRobotTaskCalledFileRequest, runtime *dara.RuntimeOptions) (_result *UploadRobotTaskCalledFileResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CalledNumber) {
		query["CalledNumber"] = request.CalledNumber
	}

	if !dara.IsNil(request.Id) {
		query["Id"] = request.Id
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.TtsParam) {
		query["TtsParam"] = request.TtsParam
	}

	if !dara.IsNil(request.TtsParamHead) {
		query["TtsParamHead"] = request.TtsParamHead
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadRobotTaskCalledFile"),
		Version:     dara.String("2017-05-25"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadRobotTaskCalledFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads the called numbers of a robocall task.
//
// Description:
//
// ### QPS limits
//
// You can call this operation up to 100 times per second per account.
//
// @param request - UploadRobotTaskCalledFileRequest
//
// @return UploadRobotTaskCalledFileResponse
func (client *Client) UploadRobotTaskCalledFile(request *UploadRobotTaskCalledFileRequest) (_result *UploadRobotTaskCalledFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UploadRobotTaskCalledFileResponse{}
	_body, _err := client.UploadRobotTaskCalledFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
