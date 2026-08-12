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
	client.EndpointRule = dara.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-zhangjiakou":        dara.String("aisc.cn-shanghai.aliyuncs.com"),
		"cn-wulanchabu":         dara.String("aisc.cn-shanghai.aliyuncs.com"),
		"cn-shanghai":           dara.String("aisc.cn-shanghai.aliyuncs.com"),
		"cn-qingdao":            dara.String("aisc.cn-shanghai.aliyuncs.com"),
		"cn-nanjing":            dara.String("aisc.cn-shanghai.aliyuncs.com"),
		"cn-huhehaote":          dara.String("aisc.cn-shanghai.aliyuncs.com"),
		"cn-hangzhou":           dara.String("aisc.cn-shanghai.aliyuncs.com"),
		"cn-guangzhou":          dara.String("aisc.cn-shanghai.aliyuncs.com"),
		"cn-beijing":            dara.String("aisc.cn-shanghai.aliyuncs.com"),
		"ap-southeast-7":        dara.String("aisc.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-6":        dara.String("aisc.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-5":        dara.String("aisc.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-1":        dara.String("aisc.ap-southeast-1.aliyuncs.com"),
		"ap-northeast-2":        dara.String("aisc.ap-southeast-1.aliyuncs.com"),
		"ap-northeast-1":        dara.String("aisc.ap-southeast-1.aliyuncs.com"),
		"eu-central-1":          dara.String("aisc.ap-southeast-1.aliyuncs.com"),
		"eu-west-1":             dara.String("aisc.ap-southeast-1.aliyuncs.com"),
		"us-east-1":             dara.String("aisc.ap-southeast-1.aliyuncs.com"),
		"us-west-1":             dara.String("aisc.ap-southeast-1.aliyuncs.com"),
		"me-east-1":             dara.String("aisc.ap-southeast-1.aliyuncs.com"),
		"cn-beijing-finance-1":  dara.String("aisc.cn-shanghai.aliyuncs.com"),
		"cn-hangzhou-finance":   dara.String("aisc.cn-shanghai.aliyuncs.com"),
		"cn-heyuan-acdr-1":      dara.String("aisc.cn-shanghai.aliyuncs.com"),
		"cn-shanghai-finance-1": dara.String("aisc.cn-shanghai.aliyuncs.com"),
		"cn-shenzhen-finance-1": dara.String("aisc.cn-shanghai.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("aisc"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Initiates batch detection for user-defined skills.
//
// @param request - CreateSkillFileCheckRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSkillFileCheckResponse
func (client *Client) CreateSkillFileCheckWithOptions(request *CreateSkillFileCheckRequest, runtime *dara.RuntimeOptions) (_result *CreateSkillFileCheckResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Files) {
		query["Files"] = request.Files
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSkillFileCheck"),
		Version:     dara.String("2026-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSkillFileCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Initiates batch detection for user-defined skills.
//
// @param request - CreateSkillFileCheckRequest
//
// @return CreateSkillFileCheckResponse
func (client *Client) CreateSkillFileCheck(request *CreateSkillFileCheckRequest) (_result *CreateSkillFileCheckResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateSkillFileCheckResponse{}
	_body, _err := client.CreateSkillFileCheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of agent risk events.
//
// @param request - ListAIAgentEventRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAIAgentEventResponse
func (client *Client) ListAIAgentEventWithOptions(request *ListAIAgentEventRequest, runtime *dara.RuntimeOptions) (_result *ListAIAgentEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		query["AppId"] = request.AppId
	}

	if !dara.IsNil(request.AppName) {
		query["AppName"] = request.AppName
	}

	if !dara.IsNil(request.AssetName) {
		query["AssetName"] = request.AssetName
	}

	if !dara.IsNil(request.AssetType) {
		query["AssetType"] = request.AssetType
	}

	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.InfraInstanceId) {
		query["InfraInstanceId"] = request.InfraInstanceId
	}

	if !dara.IsNil(request.InfraName) {
		query["InfraName"] = request.InfraName
	}

	if !dara.IsNil(request.InfraRegionId) {
		query["InfraRegionId"] = request.InfraRegionId
	}

	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RiskLevel) {
		query["RiskLevel"] = request.RiskLevel
	}

	if !dara.IsNil(request.RiskName) {
		query["RiskName"] = request.RiskName
	}

	if !dara.IsNil(request.Source) {
		query["Source"] = request.Source
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.StatusList) {
		query["StatusList"] = request.StatusList
	}

	if !dara.IsNil(request.Vendor) {
		query["Vendor"] = request.Vendor
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAIAgentEvent"),
		Version:     dara.String("2026-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAIAgentEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of agent risk events.
//
// @param request - ListAIAgentEventRequest
//
// @return ListAIAgentEventResponse
func (client *Client) ListAIAgentEvent(request *ListAIAgentEventRequest) (_result *ListAIAgentEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListAIAgentEventResponse{}
	_body, _err := client.ListAIAgentEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get subtask information.
//
// @param request - ListSubTasksRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSubTasksResponse
func (client *Client) ListSubTasksWithOptions(request *ListSubTasksRequest, runtime *dara.RuntimeOptions) (_result *ListSubTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		query["CurrentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RootTaskId) {
		query["RootTaskId"] = request.RootTaskId
	}

	if !dara.IsNil(request.TaskType) {
		query["TaskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSubTasks"),
		Version:     dara.String("2026-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSubTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get subtask information.
//
// @param request - ListSubTasksRequest
//
// @return ListSubTasksResponse
func (client *Client) ListSubTasks(request *ListSubTasksRequest) (_result *ListSubTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListSubTasksResponse{}
	_body, _err := client.ListSubTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
