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
	client.EndpointRule = dara.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("ecs-workbench"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 获取实例录屏配置
//
// @param request - GetInstanceRecordConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInstanceRecordConfigResponse
func (client *Client) GetInstanceRecordConfigWithOptions(request *GetInstanceRecordConfigRequest, runtime *dara.RuntimeOptions) (_result *GetInstanceRecordConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInstanceRecordConfig"),
		Version:     dara.String("2022-02-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInstanceRecordConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实例录屏配置
//
// @param request - GetInstanceRecordConfigRequest
//
// @return GetInstanceRecordConfigResponse
func (client *Client) GetInstanceRecordConfig(request *GetInstanceRecordConfigRequest) (_result *GetInstanceRecordConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetInstanceRecordConfigResponse{}
	_body, _err := client.GetInstanceRecordConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取实例录屏记录列表
//
// @param request - ListInstanceRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInstanceRecordsResponse
func (client *Client) ListInstanceRecordsWithOptions(request *ListInstanceRecordsRequest, runtime *dara.RuntimeOptions) (_result *ListInstanceRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListInstanceRecords"),
		Version:     dara.String("2022-02-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListInstanceRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取实例录屏记录列表
//
// @param request - ListInstanceRecordsRequest
//
// @return ListInstanceRecordsResponse
func (client *Client) ListInstanceRecords(request *ListInstanceRecordsRequest) (_result *ListInstanceRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListInstanceRecordsResponse{}
	_body, _err := client.ListInstanceRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看实例Workbench登录后执行命令的历史列表。
//
// @param request - ListTerminalCommandsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTerminalCommandsResponse
func (client *Client) ListTerminalCommandsWithOptions(request *ListTerminalCommandsRequest, runtime *dara.RuntimeOptions) (_result *ListTerminalCommandsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		body["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		body["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TerminalSessionToken) {
		body["TerminalSessionToken"] = request.TerminalSessionToken
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTerminalCommands"),
		Version:     dara.String("2022-02-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTerminalCommandsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看实例Workbench登录后执行命令的历史列表。
//
// @param request - ListTerminalCommandsRequest
//
// @return ListTerminalCommandsResponse
func (client *Client) ListTerminalCommands(request *ListTerminalCommandsRequest) (_result *ListTerminalCommandsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTerminalCommandsResponse{}
	_body, _err := client.ListTerminalCommandsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 登录实例
//
// @param request - LoginInstanceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LoginInstanceResponse
func (client *Client) LoginInstanceWithOptions(request *LoginInstanceRequest, runtime *dara.RuntimeOptions) (_result *LoginInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceLoginInfo) {
		query["InstanceLoginInfo"] = request.InstanceLoginInfo
	}

	if !dara.IsNil(request.PartnerInfo) {
		query["PartnerInfo"] = request.PartnerInfo
	}

	if !dara.IsNil(request.UserAccount) {
		query["UserAccount"] = request.UserAccount
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("LoginInstance"),
		Version:     dara.String("2022-02-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &LoginInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 登录实例
//
// @param request - LoginInstanceRequest
//
// @return LoginInstanceResponse
func (client *Client) LoginInstance(request *LoginInstanceRequest) (_result *LoginInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &LoginInstanceResponse{}
	_body, _err := client.LoginInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 设置实例录屏配置
//
// @param request - SetInstanceRecordConfigRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetInstanceRecordConfigResponse
func (client *Client) SetInstanceRecordConfigWithOptions(request *SetInstanceRecordConfigRequest, runtime *dara.RuntimeOptions) (_result *SetInstanceRecordConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Enabled) {
		body["Enabled"] = request.Enabled
	}

	if !dara.IsNil(request.ExpirationDays) {
		body["ExpirationDays"] = request.ExpirationDays
	}

	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RecordStorageTarget) {
		body["RecordStorageTarget"] = request.RecordStorageTarget
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceRegionId) {
		body["ResourceRegionId"] = request.ResourceRegionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetInstanceRecordConfig"),
		Version:     dara.String("2022-02-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SetInstanceRecordConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置实例录屏配置
//
// @param request - SetInstanceRecordConfigRequest
//
// @return SetInstanceRecordConfigResponse
func (client *Client) SetInstanceRecordConfig(request *SetInstanceRecordConfigRequest) (_result *SetInstanceRecordConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SetInstanceRecordConfigResponse{}
	_body, _err := client.SetInstanceRecordConfigWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看实例录屏内容
//
// @param request - ViewInstanceRecordsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ViewInstanceRecordsResponse
func (client *Client) ViewInstanceRecordsWithOptions(request *ViewInstanceRecordsRequest, runtime *dara.RuntimeOptions) (_result *ViewInstanceRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		body["InstanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.RegionId) {
		body["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.TerminalSessionToken) {
		body["TerminalSessionToken"] = request.TerminalSessionToken
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ViewInstanceRecords"),
		Version:     dara.String("2022-02-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ViewInstanceRecordsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看实例录屏内容
//
// @param request - ViewInstanceRecordsRequest
//
// @return ViewInstanceRecordsResponse
func (client *Client) ViewInstanceRecords(request *ViewInstanceRecordsRequest) (_result *ViewInstanceRecordsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ViewInstanceRecordsResponse{}
	_body, _err := client.ViewInstanceRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
