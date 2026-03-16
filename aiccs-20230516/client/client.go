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
	client.Endpoint, _err = client.GetEndpoint(dara.String("aiccs"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 创建坐席
//
// @param request - AddAgentRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddAgentResponse
func (client *Client) AddAgentWithOptions(request *AddAgentRequest, runtime *dara.RuntimeOptions) (_result *AddAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Account) {
		query["Account"] = request.Account
	}

	if !dara.IsNil(request.AgentTag) {
		query["AgentTag"] = request.AgentTag
	}

	if !dara.IsNil(request.ExtensionPwd) {
		query["ExtensionPwd"] = request.ExtensionPwd
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddAgent"),
		Version:     dara.String("2023-05-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddAgentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建坐席
//
// @param request - AddAgentRequest
//
// @return AddAgentResponse
func (client *Client) AddAgent(request *AddAgentRequest) (_result *AddAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddAgentResponse{}
	_body, _err := client.AddAgentWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建坐席组接口
//
// @param request - AddAgentGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddAgentGroupResponse
func (client *Client) AddAgentGroupWithOptions(request *AddAgentGroupRequest, runtime *dara.RuntimeOptions) (_result *AddAgentGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentGroupName) {
		query["AgentGroupName"] = request.AgentGroupName
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
		Action:      dara.String("AddAgentGroup"),
		Version:     dara.String("2023-05-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddAgentGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建坐席组接口
//
// @param request - AddAgentGroupRequest
//
// @return AddAgentGroupResponse
func (client *Client) AddAgentGroup(request *AddAgentGroupRequest) (_result *AddAgentGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddAgentGroupResponse{}
	_body, _err := client.AddAgentGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加黑名单接口
//
// @param tmpReq - AddBlacklistRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddBlacklistResponse
func (client *Client) AddBlacklistWithOptions(tmpReq *AddBlacklistRequest, runtime *dara.RuntimeOptions) (_result *AddBlacklistResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddBlacklistShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Numbers) {
		request.NumbersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Numbers, dara.String("Numbers"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ExpiredDay) {
		query["ExpiredDay"] = request.ExpiredDay
	}

	if !dara.IsNil(request.NumbersShrink) {
		query["Numbers"] = request.NumbersShrink
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Remark) {
		query["Remark"] = request.Remark
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
		Action:      dara.String("AddBlacklist"),
		Version:     dara.String("2023-05-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddBlacklistResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加黑名单接口
//
// @param request - AddBlacklistRequest
//
// @return AddBlacklistResponse
func (client *Client) AddBlacklist(request *AddBlacklistRequest) (_result *AddBlacklistResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddBlacklistResponse{}
	_body, _err := client.AddBlacklistWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建任务接口
//
// @param tmpReq - AddTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTaskResponse
func (client *Client) AddTaskWithOptions(tmpReq *AddTaskRequest, runtime *dara.RuntimeOptions) (_result *AddTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AddTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CallTimeList) {
		request.CallTimeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CallTimeList, dara.String("CallTimeList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CallTimeStrList) {
		request.CallTimeStrListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CallTimeStrList, dara.String("CallTimeStrList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RepeatReason) {
		request.RepeatReasonShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RepeatReason, dara.String("RepeatReason"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RepeatTimes) {
		request.RepeatTimesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RepeatTimes, dara.String("RepeatTimes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SendSmsPlan) {
		request.SendSmsPlanShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SendSmsPlan, dara.String("SendSmsPlan"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CallTimeListShrink) {
		query["CallTimeList"] = request.CallTimeListShrink
	}

	if !dara.IsNil(request.CallTimeStrListShrink) {
		query["CallTimeStrList"] = request.CallTimeStrListShrink
	}

	if !dara.IsNil(request.CallbackUrl) {
		query["CallbackUrl"] = request.CallbackUrl
	}

	if !dara.IsNil(request.FlashSmsTemplateId) {
		query["FlashSmsTemplateId"] = request.FlashSmsTemplateId
	}

	if !dara.IsNil(request.FlashSmsType) {
		query["FlashSmsType"] = request.FlashSmsType
	}

	if !dara.IsNil(request.MaxConcurrency) {
		query["MaxConcurrency"] = request.MaxConcurrency
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PlaySleepVal) {
		query["PlaySleepVal"] = request.PlaySleepVal
	}

	if !dara.IsNil(request.PlayTimes) {
		query["PlayTimes"] = request.PlayTimes
	}

	if !dara.IsNil(request.RecallType) {
		query["RecallType"] = request.RecallType
	}

	if !dara.IsNil(request.RecordPath) {
		query["RecordPath"] = request.RecordPath
	}

	if !dara.IsNil(request.RepeatCount) {
		query["RepeatCount"] = request.RepeatCount
	}

	if !dara.IsNil(request.RepeatInterval) {
		query["RepeatInterval"] = request.RepeatInterval
	}

	if !dara.IsNil(request.RepeatReasonShrink) {
		query["RepeatReason"] = request.RepeatReasonShrink
	}

	if !dara.IsNil(request.RepeatTimesShrink) {
		query["RepeatTimes"] = request.RepeatTimesShrink
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SendSmsPlanShrink) {
		query["SendSmsPlan"] = request.SendSmsPlanShrink
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddTask"),
		Version:     dara.String("2023-05-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建任务接口
//
// @param request - AddTaskRequest
//
// @return AddTaskResponse
func (client *Client) AddTask(request *AddTaskRequest) (_result *AddTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddTaskResponse{}
	_body, _err := client.AddTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 坐席外呼查询外呼记录
//
// @param tmpReq - AgentCallListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AgentCallListResponse
func (client *Client) AgentCallListWithOptions(tmpReq *AgentCallListRequest, runtime *dara.RuntimeOptions) (_result *AgentCallListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AgentCallListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NumberMD5s) {
		request.NumberMD5sShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NumberMD5s, dara.String("NumberMD5s"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Numbers) {
		request.NumbersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Numbers, dara.String("Numbers"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.AgentTag) {
		query["AgentTag"] = request.AgentTag
	}

	if !dara.IsNil(request.CallDate) {
		query["CallDate"] = request.CallDate
	}

	if !dara.IsNil(request.EndCallDate) {
		query["EndCallDate"] = request.EndCallDate
	}

	if !dara.IsNil(request.NumberMD5sShrink) {
		query["NumberMD5s"] = request.NumberMD5sShrink
	}

	if !dara.IsNil(request.NumbersShrink) {
		query["Numbers"] = request.NumbersShrink
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Page) {
		query["Page"] = request.Page
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

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AgentCallList"),
		Version:     dara.String("2023-05-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AgentCallListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 坐席外呼查询外呼记录
//
// @param request - AgentCallListRequest
//
// @return AgentCallListResponse
func (client *Client) AgentCallList(request *AgentCallListRequest) (_result *AgentCallListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AgentCallListResponse{}
	_body, _err := client.AgentCallListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 坐席取消号码外呼
//
// @param tmpReq - AgentCancelCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AgentCancelCallResponse
func (client *Client) AgentCancelCallWithOptions(tmpReq *AgentCancelCallRequest, runtime *dara.RuntimeOptions) (_result *AgentCancelCallResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AgentCancelCallShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Numbers) {
		request.NumbersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Numbers, dara.String("Numbers"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.AgentTag) {
		query["AgentTag"] = request.AgentTag
	}

	if !dara.IsNil(request.NumbersShrink) {
		query["Numbers"] = request.NumbersShrink
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

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AgentCancelCall"),
		Version:     dara.String("2023-05-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AgentCancelCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 坐席取消号码外呼
//
// @param request - AgentCancelCallRequest
//
// @return AgentCancelCallResponse
func (client *Client) AgentCancelCall(request *AgentCancelCallRequest) (_result *AgentCancelCallResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AgentCancelCallResponse{}
	_body, _err := client.AgentCancelCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 坐席组分页查询
//
// @param request - AgentGroupPageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AgentGroupPageResponse
func (client *Client) AgentGroupPageWithOptions(request *AgentGroupPageRequest, runtime *dara.RuntimeOptions) (_result *AgentGroupPageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentGroupId) {
		query["AgentGroupId"] = request.AgentGroupId
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PageNum) {
		query["PageNum"] = request.PageNum
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AgentGroupPage"),
		Version:     dara.String("2023-05-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AgentGroupPageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 坐席组分页查询
//
// @param request - AgentGroupPageRequest
//
// @return AgentGroupPageResponse
func (client *Client) AgentGroupPage(request *AgentGroupPageRequest) (_result *AgentGroupPageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AgentGroupPageResponse{}
	_body, _err := client.AgentGroupPageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 坐席外呼导入号码
//
// @param tmpReq - AgentImportNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AgentImportNumberResponse
func (client *Client) AgentImportNumberWithOptions(tmpReq *AgentImportNumberRequest, runtime *dara.RuntimeOptions) (_result *AgentImportNumberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AgentImportNumberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Customers) {
		request.CustomersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Customers, dara.String("Customers"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.AgentTag) {
		query["AgentTag"] = request.AgentTag
	}

	if !dara.IsNil(request.CallType) {
		query["CallType"] = request.CallType
	}

	if !dara.IsNil(request.CustomersShrink) {
		query["Customers"] = request.CustomersShrink
	}

	if !dara.IsNil(request.GatewayId) {
		query["GatewayId"] = request.GatewayId
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

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AgentImportNumber"),
		Version:     dara.String("2023-05-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AgentImportNumberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 坐席外呼导入号码
//
// @param request - AgentImportNumberRequest
//
// @return AgentImportNumberResponse
func (client *Client) AgentImportNumber(request *AgentImportNumberRequest) (_result *AgentImportNumberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AgentImportNumberResponse{}
	_body, _err := client.AgentImportNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 坐席任务恢复号码
//
// @param tmpReq - AgentRecoverCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AgentRecoverCallResponse
func (client *Client) AgentRecoverCallWithOptions(tmpReq *AgentRecoverCallRequest, runtime *dara.RuntimeOptions) (_result *AgentRecoverCallResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &AgentRecoverCallShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Numbers) {
		request.NumbersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Numbers, dara.String("Numbers"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.AgentTag) {
		query["AgentTag"] = request.AgentTag
	}

	if !dara.IsNil(request.BeginImportTime) {
		query["BeginImportTime"] = request.BeginImportTime
	}

	if !dara.IsNil(request.EndImportTime) {
		query["EndImportTime"] = request.EndImportTime
	}

	if !dara.IsNil(request.NumbersShrink) {
		query["Numbers"] = request.NumbersShrink
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

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AgentRecoverCall"),
		Version:     dara.String("2023-05-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AgentRecoverCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 坐席任务恢复号码
//
// @param request - AgentRecoverCallRequest
//
// @return AgentRecoverCallResponse
func (client *Client) AgentRecoverCall(request *AgentRecoverCallRequest) (_result *AgentRecoverCallResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AgentRecoverCallResponse{}
	_body, _err := client.AgentRecoverCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取聊天内容
//
// @param request - CallChatListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CallChatListResponse
func (client *Client) CallChatListWithOptions(request *CallChatListRequest, runtime *dara.RuntimeOptions) (_result *CallChatListResponse, _err error) {
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

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CallChatList"),
		Version:     dara.String("2023-05-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CallChatListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取聊天内容
//
// @param request - CallChatListRequest
//
// @return CallChatListResponse
func (client *Client) CallChatList(request *CallChatListRequest) (_result *CallChatListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CallChatListResponse{}
	_body, _err := client.CallChatListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取号码外呼详情
//
// @param request - CallNumberDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CallNumberDetailResponse
func (client *Client) CallNumberDetailWithOptions(request *CallNumberDetailRequest, runtime *dara.RuntimeOptions) (_result *CallNumberDetailResponse, _err error) {
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

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CallNumberDetail"),
		Version:     dara.String("2023-05-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CallNumberDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取号码外呼详情
//
// @param request - CallNumberDetailRequest
//
// @return CallNumberDetailResponse
func (client *Client) CallNumberDetail(request *CallNumberDetailRequest) (_result *CallNumberDetailResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CallNumberDetailResponse{}
	_body, _err := client.CallNumberDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # AI批量任务查询号码状态接口
//
// @param tmpReq - DetailsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetailsResponse
func (client *Client) DetailsWithOptions(tmpReq *DetailsRequest, runtime *dara.RuntimeOptions) (_result *DetailsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DetailsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Numbers) {
		request.NumbersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Numbers, dara.String("Numbers"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BatchId) {
		query["BatchId"] = request.BatchId
	}

	if !dara.IsNil(request.EndTime) {
		query["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.NumberStatus) {
		query["NumberStatus"] = request.NumberStatus
	}

	if !dara.IsNil(request.NumbersShrink) {
		query["Numbers"] = request.NumbersShrink
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

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Details"),
		Version:     dara.String("2023-05-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # AI批量任务查询号码状态接口
//
// @param request - DetailsRequest
//
// @return DetailsResponse
func (client *Client) Details(request *DetailsRequest) (_result *DetailsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DetailsResponse{}
	_body, _err := client.DetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 编辑任务接口
//
// @param tmpReq - EditTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EditTaskResponse
func (client *Client) EditTaskWithOptions(tmpReq *EditTaskRequest, runtime *dara.RuntimeOptions) (_result *EditTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &EditTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CallTimeList) {
		request.CallTimeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CallTimeList, dara.String("CallTimeList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CallTimeStrList) {
		request.CallTimeStrListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CallTimeStrList, dara.String("CallTimeStrList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RepeatReason) {
		request.RepeatReasonShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RepeatReason, dara.String("RepeatReason"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.RepeatTimes) {
		request.RepeatTimesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RepeatTimes, dara.String("RepeatTimes"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.SendSmsPlan) {
		request.SendSmsPlanShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SendSmsPlan, dara.String("SendSmsPlan"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CallTimeListShrink) {
		query["CallTimeList"] = request.CallTimeListShrink
	}

	if !dara.IsNil(request.CallTimeStrListShrink) {
		query["CallTimeStrList"] = request.CallTimeStrListShrink
	}

	if !dara.IsNil(request.CallbackUrl) {
		query["CallbackUrl"] = request.CallbackUrl
	}

	if !dara.IsNil(request.FlashSmsTemplateId) {
		query["FlashSmsTemplateId"] = request.FlashSmsTemplateId
	}

	if !dara.IsNil(request.FlashSmsType) {
		query["FlashSmsType"] = request.FlashSmsType
	}

	if !dara.IsNil(request.MaxConcurrency) {
		query["MaxConcurrency"] = request.MaxConcurrency
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.PlaySleepVal) {
		query["PlaySleepVal"] = request.PlaySleepVal
	}

	if !dara.IsNil(request.PlayTimes) {
		query["PlayTimes"] = request.PlayTimes
	}

	if !dara.IsNil(request.RecallType) {
		query["RecallType"] = request.RecallType
	}

	if !dara.IsNil(request.RecordPath) {
		query["RecordPath"] = request.RecordPath
	}

	if !dara.IsNil(request.RepeatCount) {
		query["RepeatCount"] = request.RepeatCount
	}

	if !dara.IsNil(request.RepeatInterval) {
		query["RepeatInterval"] = request.RepeatInterval
	}

	if !dara.IsNil(request.RepeatReasonShrink) {
		query["RepeatReason"] = request.RepeatReasonShrink
	}

	if !dara.IsNil(request.RepeatTimesShrink) {
		query["RepeatTimes"] = request.RepeatTimesShrink
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SendSmsPlanShrink) {
		query["SendSmsPlan"] = request.SendSmsPlanShrink
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EditTask"),
		Version:     dara.String("2023-05-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &EditTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑任务接口
//
// @param request - EditTaskRequest
//
// @return EditTaskResponse
func (client *Client) EditTask(request *EditTaskRequest) (_result *EditTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &EditTaskResponse{}
	_body, _err := client.EditTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 导入号码
//
// @param tmpReq - ImportNumberRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportNumberResponse
func (client *Client) ImportNumberWithOptions(tmpReq *ImportNumberRequest, runtime *dara.RuntimeOptions) (_result *ImportNumberResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ImportNumberShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Customers) {
		request.CustomersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Customers, dara.String("Customers"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomersShrink) {
		query["Customers"] = request.CustomersShrink
	}

	if !dara.IsNil(request.FailReturn) {
		query["FailReturn"] = request.FailReturn
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

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportNumber"),
		Version:     dara.String("2023-05-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportNumberResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导入号码
//
// @param request - ImportNumberRequest
//
// @return ImportNumberResponse
func (client *Client) ImportNumber(request *ImportNumberRequest) (_result *ImportNumberResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ImportNumberResponse{}
	_body, _err := client.ImportNumberWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 导入号码
//
// @param tmpReq - ImportNumberV2Request
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportNumberV2Response
func (client *Client) ImportNumberV2WithOptions(tmpReq *ImportNumberV2Request, runtime *dara.RuntimeOptions) (_result *ImportNumberV2Response, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ImportNumberV2ShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Customers) {
		request.CustomersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Customers, dara.String("Customers"), dara.String("json"))
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.CustomersShrink) {
		body["Customers"] = request.CustomersShrink
	}

	if !dara.IsNil(request.FailReturn) {
		body["FailReturn"] = request.FailReturn
	}

	if !dara.IsNil(request.OutId) {
		body["OutId"] = request.OutId
	}

	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ImportNumberV2"),
		Version:     dara.String("2023-05-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ImportNumberV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 导入号码
//
// @param request - ImportNumberV2Request
//
// @return ImportNumberV2Response
func (client *Client) ImportNumberV2(request *ImportNumberV2Request) (_result *ImportNumberV2Response, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ImportNumberV2Response{}
	_body, _err := client.ImportNumberV2WithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 绑定坐席组
//
// @param tmpReq - JoinAgentGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return JoinAgentGroupResponse
func (client *Client) JoinAgentGroupWithOptions(tmpReq *JoinAgentGroupRequest, runtime *dara.RuntimeOptions) (_result *JoinAgentGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &JoinAgentGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AgentIds) {
		request.AgentIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AgentIds, dara.String("AgentIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentGroupId) {
		query["AgentGroupId"] = request.AgentGroupId
	}

	if !dara.IsNil(request.AgentIdsShrink) {
		query["AgentIds"] = request.AgentIdsShrink
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
		Action:      dara.String("JoinAgentGroup"),
		Version:     dara.String("2023-05-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &JoinAgentGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 绑定坐席组
//
// @param request - JoinAgentGroupRequest
//
// @return JoinAgentGroupResponse
func (client *Client) JoinAgentGroup(request *JoinAgentGroupRequest) (_result *JoinAgentGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &JoinAgentGroupResponse{}
	_body, _err := client.JoinAgentGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询企业黑名单
//
// @param tmpReq - PageRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PageResponse
func (client *Client) PageWithOptions(tmpReq *PageRequest, runtime *dara.RuntimeOptions) (_result *PageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &PageShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Numbers) {
		request.NumbersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Numbers, dara.String("Numbers"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.NumbersShrink) {
		query["Numbers"] = request.NumbersShrink
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

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Page"),
		Version:     dara.String("2023-05-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询企业黑名单
//
// @param request - PageRequest
//
// @return PageResponse
func (client *Client) Page(request *PageRequest) (_result *PageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PageResponse{}
	_body, _err := client.PageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询坐席具体信息
//
// @param request - QueryAgentInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAgentInfoResponse
func (client *Client) QueryAgentInfoWithOptions(request *QueryAgentInfoRequest, runtime *dara.RuntimeOptions) (_result *QueryAgentInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.AgentTag) {
		query["AgentTag"] = request.AgentTag
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
		Action:      dara.String("QueryAgentInfo"),
		Version:     dara.String("2023-05-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAgentInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询坐席具体信息
//
// @param request - QueryAgentInfoRequest
//
// @return QueryAgentInfoResponse
func (client *Client) QueryAgentInfo(request *QueryAgentInfoRequest) (_result *QueryAgentInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryAgentInfoResponse{}
	_body, _err := client.QueryAgentInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 快速创建任务接口
//
// @param tmpReq - QuickAddTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuickAddTaskResponse
func (client *Client) QuickAddTaskWithOptions(tmpReq *QuickAddTaskRequest, runtime *dara.RuntimeOptions) (_result *QuickAddTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QuickAddTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CallTimeList) {
		request.CallTimeListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CallTimeList, dara.String("CallTimeList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.CallTimeStrList) {
		request.CallTimeStrListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CallTimeStrList, dara.String("CallTimeStrList"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentGroupId) {
		query["AgentGroupId"] = request.AgentGroupId
	}

	if !dara.IsNil(request.CallTimeListShrink) {
		query["CallTimeList"] = request.CallTimeListShrink
	}

	if !dara.IsNil(request.CallTimeStrListShrink) {
		query["CallTimeStrList"] = request.CallTimeStrListShrink
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.ReferenceTaskId) {
		query["ReferenceTaskId"] = request.ReferenceTaskId
	}

	if !dara.IsNil(request.ResourceOwnerAccount) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !dara.IsNil(request.ResourceOwnerId) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !dara.IsNil(request.SmsTemplateId) {
		query["SmsTemplateId"] = request.SmsTemplateId
	}

	if !dara.IsNil(request.StartTime) {
		query["StartTime"] = request.StartTime
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuickAddTask"),
		Version:     dara.String("2023-05-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuickAddTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 快速创建任务接口
//
// @param request - QuickAddTaskRequest
//
// @return QuickAddTaskResponse
func (client *Client) QuickAddTask(request *QuickAddTaskRequest) (_result *QuickAddTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuickAddTaskResponse{}
	_body, _err := client.QuickAddTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 解绑坐席组
//
// @param tmpReq - QuitAgentGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuitAgentGroupResponse
func (client *Client) QuitAgentGroupWithOptions(tmpReq *QuitAgentGroupRequest, runtime *dara.RuntimeOptions) (_result *QuitAgentGroupResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QuitAgentGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AgentIds) {
		request.AgentIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AgentIds, dara.String("AgentIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentGroupId) {
		query["AgentGroupId"] = request.AgentGroupId
	}

	if !dara.IsNil(request.AgentIdsShrink) {
		query["AgentIds"] = request.AgentIdsShrink
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
		Action:      dara.String("QuitAgentGroup"),
		Version:     dara.String("2023-05-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QuitAgentGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解绑坐席组
//
// @param request - QuitAgentGroupRequest
//
// @return QuitAgentGroupResponse
func (client *Client) QuitAgentGroup(request *QuitAgentGroupRequest) (_result *QuitAgentGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QuitAgentGroupResponse{}
	_body, _err := client.QuitAgentGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 短信模板创建
//
// @param request - SmsTemplateCreateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SmsTemplateCreateResponse
func (client *Client) SmsTemplateCreateWithOptions(request *SmsTemplateCreateRequest, runtime *dara.RuntimeOptions) (_result *SmsTemplateCreateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		query["Content"] = request.Content
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

	if !dara.IsNil(request.Sign) {
		query["Sign"] = request.Sign
	}

	if !dara.IsNil(request.SmsType) {
		query["SmsType"] = request.SmsType
	}

	if !dara.IsNil(request.TemplateName) {
		query["TemplateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SmsTemplateCreate"),
		Version:     dara.String("2023-05-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SmsTemplateCreateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 短信模板创建
//
// @param request - SmsTemplateCreateRequest
//
// @return SmsTemplateCreateResponse
func (client *Client) SmsTemplateCreate(request *SmsTemplateCreateRequest) (_result *SmsTemplateCreateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SmsTemplateCreateResponse{}
	_body, _err := client.SmsTemplateCreateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 短信模板列表查询
//
// @param request - SmsTemplatePageListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SmsTemplatePageListResponse
func (client *Client) SmsTemplatePageListWithOptions(request *SmsTemplatePageListRequest, runtime *dara.RuntimeOptions) (_result *SmsTemplatePageListResponse, _err error) {
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

	if !dara.IsNil(request.Sign) {
		query["Sign"] = request.Sign
	}

	if !dara.IsNil(request.SmsType) {
		query["SmsType"] = request.SmsType
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	if !dara.IsNil(request.TemplateType) {
		query["TemplateType"] = request.TemplateType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SmsTemplatePageList"),
		Version:     dara.String("2023-05-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SmsTemplatePageListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 短信模板列表查询
//
// @param request - SmsTemplatePageListRequest
//
// @return SmsTemplatePageListResponse
func (client *Client) SmsTemplatePageList(request *SmsTemplatePageListRequest) (_result *SmsTemplatePageListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SmsTemplatePageListResponse{}
	_body, _err := client.SmsTemplatePageListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询聊天记录接口
//
// @param request - TaskCallChatsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TaskCallChatsResponse
func (client *Client) TaskCallChatsWithOptions(request *TaskCallChatsRequest, runtime *dara.RuntimeOptions) (_result *TaskCallChatsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.AgentTag) {
		query["AgentTag"] = request.AgentTag
	}

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

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TaskCallChats"),
		Version:     dara.String("2023-05-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TaskCallChatsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询聊天记录接口
//
// @param request - TaskCallChatsRequest
//
// @return TaskCallChatsResponse
func (client *Client) TaskCallChats(request *TaskCallChatsRequest) (_result *TaskCallChatsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TaskCallChatsResponse{}
	_body, _err := client.TaskCallChatsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取任务外呼情况接口
//
// @param request - TaskCallInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TaskCallInfoResponse
func (client *Client) TaskCallInfoWithOptions(request *TaskCallInfoRequest, runtime *dara.RuntimeOptions) (_result *TaskCallInfoResponse, _err error) {
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
		Action:      dara.String("TaskCallInfo"),
		Version:     dara.String("2023-05-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TaskCallInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务外呼情况接口
//
// @param request - TaskCallInfoRequest
//
// @return TaskCallInfoResponse
func (client *Client) TaskCallInfo(request *TaskCallInfoRequest) (_result *TaskCallInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TaskCallInfoResponse{}
	_body, _err := client.TaskCallInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # AI批量任务查询外呼记录接口
//
// @param tmpReq - TaskCallListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TaskCallListResponse
func (client *Client) TaskCallListWithOptions(tmpReq *TaskCallListRequest, runtime *dara.RuntimeOptions) (_result *TaskCallListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &TaskCallListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.IntentTags) {
		request.IntentTagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IntentTags, dara.String("IntentTags"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Numbers) {
		request.NumbersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Numbers, dara.String("Numbers"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BatchId) {
		query["BatchId"] = request.BatchId
	}

	if !dara.IsNil(request.CallDate) {
		query["CallDate"] = request.CallDate
	}

	if !dara.IsNil(request.EndCallDate) {
		query["EndCallDate"] = request.EndCallDate
	}

	if !dara.IsNil(request.IntentTagsShrink) {
		query["IntentTags"] = request.IntentTagsShrink
	}

	if !dara.IsNil(request.NumbersShrink) {
		query["Numbers"] = request.NumbersShrink
	}

	if !dara.IsNil(request.OwnerId) {
		query["OwnerId"] = request.OwnerId
	}

	if !dara.IsNil(request.Page) {
		query["Page"] = request.Page
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
		Action:      dara.String("TaskCallList"),
		Version:     dara.String("2023-05-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TaskCallListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # AI批量任务查询外呼记录接口
//
// @param request - TaskCallListRequest
//
// @return TaskCallListResponse
func (client *Client) TaskCallList(request *TaskCallListRequest) (_result *TaskCallListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TaskCallListResponse{}
	_body, _err := client.TaskCallListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量任务取消号码外呼
//
// @param tmpReq - TaskCancelCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TaskCancelCallResponse
func (client *Client) TaskCancelCallWithOptions(tmpReq *TaskCancelCallRequest, runtime *dara.RuntimeOptions) (_result *TaskCancelCallResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &TaskCancelCallShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Numbers) {
		request.NumbersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Numbers, dara.String("Numbers"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.NumbersShrink) {
		query["Numbers"] = request.NumbersShrink
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

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TaskCancelCall"),
		Version:     dara.String("2023-05-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TaskCancelCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量任务取消号码外呼
//
// @param request - TaskCancelCallRequest
//
// @return TaskCancelCallResponse
func (client *Client) TaskCancelCall(request *TaskCancelCallRequest) (_result *TaskCancelCallResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TaskCancelCallResponse{}
	_body, _err := client.TaskCancelCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询任务列表接口
//
// @param request - TaskListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TaskListResponse
func (client *Client) TaskListWithOptions(request *TaskListRequest, runtime *dara.RuntimeOptions) (_result *TaskListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreateTime) {
		query["CreateTime"] = request.CreateTime
	}

	if !dara.IsNil(request.LastCallTime) {
		query["LastCallTime"] = request.LastCallTime
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
		Action:      dara.String("TaskList"),
		Version:     dara.String("2023-05-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TaskListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询任务列表接口
//
// @param request - TaskListRequest
//
// @return TaskListResponse
func (client *Client) TaskList(request *TaskListRequest) (_result *TaskListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TaskListResponse{}
	_body, _err := client.TaskListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 批量任务恢复号码
//
// @param tmpReq - TaskRecoverCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TaskRecoverCallResponse
func (client *Client) TaskRecoverCallWithOptions(tmpReq *TaskRecoverCallRequest, runtime *dara.RuntimeOptions) (_result *TaskRecoverCallResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &TaskRecoverCallShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Numbers) {
		request.NumbersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Numbers, dara.String("Numbers"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Tags) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, dara.String("Tags"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.BeginImportTime) {
		query["BeginImportTime"] = request.BeginImportTime
	}

	if !dara.IsNil(request.EndImportTime) {
		query["EndImportTime"] = request.EndImportTime
	}

	if !dara.IsNil(request.NumbersShrink) {
		query["Numbers"] = request.NumbersShrink
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

	if !dara.IsNil(request.TagsShrink) {
		query["Tags"] = request.TagsShrink
	}

	if !dara.IsNil(request.TaskId) {
		query["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TaskRecoverCall"),
		Version:     dara.String("2023-05-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TaskRecoverCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量任务恢复号码
//
// @param request - TaskRecoverCallRequest
//
// @return TaskRecoverCallResponse
func (client *Client) TaskRecoverCall(request *TaskRecoverCallRequest) (_result *TaskRecoverCallResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TaskRecoverCallResponse{}
	_body, _err := client.TaskRecoverCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 话术模板列表查询接口
//
// @param request - TemplateListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TemplateListResponse
func (client *Client) TemplateListWithOptions(request *TemplateListRequest, runtime *dara.RuntimeOptions) (_result *TemplateListResponse, _err error) {
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

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TemplateList"),
		Version:     dara.String("2023-05-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &TemplateListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 话术模板列表查询接口
//
// @param request - TemplateListRequest
//
// @return TemplateListResponse
func (client *Client) TemplateList(request *TemplateListRequest) (_result *TemplateListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &TemplateListResponse{}
	_body, _err := client.TemplateListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改坐席状态
//
// @param request - UpdateAgentStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAgentStatusResponse
func (client *Client) UpdateAgentStatusWithOptions(request *UpdateAgentStatusRequest, runtime *dara.RuntimeOptions) (_result *UpdateAgentStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["AgentId"] = request.AgentId
	}

	if !dara.IsNil(request.AgentStatus) {
		query["AgentStatus"] = request.AgentStatus
	}

	if !dara.IsNil(request.AgentTag) {
		query["AgentTag"] = request.AgentTag
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
		Action:      dara.String("UpdateAgentStatus"),
		Version:     dara.String("2023-05-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAgentStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改坐席状态
//
// @param request - UpdateAgentStatusRequest
//
// @return UpdateAgentStatusResponse
func (client *Client) UpdateAgentStatus(request *UpdateAgentStatusRequest) (_result *UpdateAgentStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateAgentStatusResponse{}
	_body, _err := client.UpdateAgentStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新当天导入的号码
//
// @param tmpReq - UpdateTaskCustomerRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTaskCustomerResponse
func (client *Client) UpdateTaskCustomerWithOptions(tmpReq *UpdateTaskCustomerRequest, runtime *dara.RuntimeOptions) (_result *UpdateTaskCustomerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UpdateTaskCustomerShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Customers) {
		request.CustomersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Customers, dara.String("Customers"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.CustomersShrink) {
		query["Customers"] = request.CustomersShrink
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
		Action:      dara.String("UpdateTaskCustomer"),
		Version:     dara.String("2023-05-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTaskCustomerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新当天导入的号码
//
// @param request - UpdateTaskCustomerRequest
//
// @return UpdateTaskCustomerResponse
func (client *Client) UpdateTaskCustomer(request *UpdateTaskCustomerRequest) (_result *UpdateTaskCustomerResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UpdateTaskCustomerResponse{}
	_body, _err := client.UpdateTaskCustomerWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
