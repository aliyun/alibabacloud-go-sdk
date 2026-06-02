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
	client.Endpoint, _err = client.GetEndpoint(dara.String("wyota"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 添加终端
//
// @param request - AddTerminalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTerminalResponse
func (client *Client) AddTerminalWithOptions(request *AddTerminalRequest, runtime *dara.RuntimeOptions) (_result *AddTerminalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Alias) {
		body["Alias"] = request.Alias
	}

	if !dara.IsNil(request.ClientType) {
		body["ClientType"] = request.ClientType
	}

	if !dara.IsNil(request.MainBizType) {
		body["MainBizType"] = request.MainBizType
	}

	if !dara.IsNil(request.SerialNumber) {
		body["SerialNumber"] = request.SerialNumber
	}

	if !dara.IsNil(request.TerminalGroupId) {
		body["TerminalGroupId"] = request.TerminalGroupId
	}

	if !dara.IsNil(request.Uuid) {
		body["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddTerminal"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddTerminalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加终端
//
// @param request - AddTerminalRequest
//
// @return AddTerminalResponse
func (client *Client) AddTerminal(request *AddTerminalRequest) (_result *AddTerminalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddTerminalResponse{}
	_body, _err := client.AddTerminalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 添加终端
//
// @param request - AddTerminalsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTerminalsResponse
func (client *Client) AddTerminalsWithOptions(request *AddTerminalsRequest, runtime *dara.RuntimeOptions) (_result *AddTerminalsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.AddTerminalParams) {
		bodyFlat["AddTerminalParams"] = request.AddTerminalParams
	}

	if !dara.IsNil(request.MainBizType) {
		body["MainBizType"] = request.MainBizType
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddTerminals"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AddTerminalsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加终端
//
// @param request - AddTerminalsRequest
//
// @return AddTerminalsResponse
func (client *Client) AddTerminals(request *AddTerminalsRequest) (_result *AddTerminalsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AddTerminalsResponse{}
	_body, _err := client.AddTerminalsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 绑定免账号登录用户
//
// @param request - BindAccountLessLoginUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAccountLessLoginUserResponse
func (client *Client) BindAccountLessLoginUserWithOptions(request *BindAccountLessLoginUserRequest, runtime *dara.RuntimeOptions) (_result *BindAccountLessLoginUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndUserId) {
		body["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.SerialNumber) {
		body["SerialNumber"] = request.SerialNumber
	}

	if !dara.IsNil(request.Uuid) {
		body["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindAccountLessLoginUser"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindAccountLessLoginUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 绑定免账号登录用户
//
// @param request - BindAccountLessLoginUserRequest
//
// @return BindAccountLessLoginUserResponse
func (client *Client) BindAccountLessLoginUser(request *BindAccountLessLoginUserRequest) (_result *BindAccountLessLoginUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindAccountLessLoginUserResponse{}
	_body, _err := client.BindAccountLessLoginUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 绑定免账号登录用户
//
// @param request - BindPasswordFreeLoginUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindPasswordFreeLoginUserResponse
func (client *Client) BindPasswordFreeLoginUserWithOptions(request *BindPasswordFreeLoginUserRequest, runtime *dara.RuntimeOptions) (_result *BindPasswordFreeLoginUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EndUserId) {
		body["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.MainBizType) {
		body["MainBizType"] = request.MainBizType
	}

	if !dara.IsNil(request.SerialNumber) {
		body["SerialNumber"] = request.SerialNumber
	}

	if !dara.IsNil(request.Uuid) {
		body["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BindPasswordFreeLoginUser"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &BindPasswordFreeLoginUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 绑定免账号登录用户
//
// @param request - BindPasswordFreeLoginUserRequest
//
// @return BindPasswordFreeLoginUserResponse
func (client *Client) BindPasswordFreeLoginUser(request *BindPasswordFreeLoginUserRequest) (_result *BindPasswordFreeLoginUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &BindPasswordFreeLoginUserResponse{}
	_body, _err := client.BindPasswordFreeLoginUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 解除桌面端、移动端纳管
//
// @param request - DeleteClientsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteClientsResponse
func (client *Client) DeleteClientsWithOptions(request *DeleteClientsRequest, runtime *dara.RuntimeOptions) (_result *DeleteClientsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallerAliUid) {
		query["CallerAliUid"] = request.CallerAliUid
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.InManage) {
		body["InManage"] = request.InManage
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.Uuids) {
		bodyFlat["Uuids"] = request.Uuids
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteClients"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteClientsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解除桌面端、移动端纳管
//
// @param request - DeleteClientsRequest
//
// @return DeleteClientsResponse
func (client *Client) DeleteClients(request *DeleteClientsRequest) (_result *DeleteClientsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DeleteClientsResponse{}
	_body, _err := client.DeleteClientsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询桌面端、移动端详细信息
//
// @param request - DescribeClientsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClientsResponse
func (client *Client) DescribeClientsWithOptions(request *DescribeClientsRequest, runtime *dara.RuntimeOptions) (_result *DescribeClientsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CallerAliUid) {
		query["CallerAliUid"] = request.CallerAliUid
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientType) {
		body["ClientType"] = request.ClientType
	}

	if !dara.IsNil(request.CustomResourceId) {
		body["CustomResourceId"] = request.CustomResourceId
	}

	if !dara.IsNil(request.CustomResourceStatus) {
		body["CustomResourceStatus"] = request.CustomResourceStatus
	}

	if !dara.IsNil(request.InManage) {
		body["InManage"] = request.InManage
	}

	if !dara.IsNil(request.IncludeSubGroups) {
		body["IncludeSubGroups"] = request.IncludeSubGroups
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.Model) {
		body["Model"] = request.Model
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.OnlineStatus) {
		body["OnlineStatus"] = request.OnlineStatus
	}

	if !dara.IsNil(request.Platform) {
		body["Platform"] = request.Platform
	}

	if !dara.IsNil(request.SearchKeyword) {
		body["SearchKeyword"] = request.SearchKeyword
	}

	if !dara.IsNil(request.TerminalGroupId) {
		body["TerminalGroupId"] = request.TerminalGroupId
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.Uuids) {
		bodyFlat["Uuids"] = request.Uuids
	}

	if !dara.IsNil(request.WithBindUser) {
		body["WithBindUser"] = request.WithBindUser
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeClients"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeClientsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询桌面端、移动端详细信息
//
// @param request - DescribeClientsRequest
//
// @return DescribeClientsResponse
func (client *Client) DescribeClients(request *DescribeClientsRequest) (_result *DescribeClientsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &DescribeClientsResponse{}
	_body, _err := client.DescribeClientsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取桌面端纳管邀请码
//
// @param request - GetOrCreateInvitationCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOrCreateInvitationCodeResponse
func (client *Client) GetOrCreateInvitationCodeWithOptions(request *GetOrCreateInvitationCodeRequest, runtime *dara.RuntimeOptions) (_result *GetOrCreateInvitationCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ExpireDays) {
		body["ExpireDays"] = request.ExpireDays
	}

	if !dara.IsNil(request.ExpireMinutes) {
		body["ExpireMinutes"] = request.ExpireMinutes
	}

	if !dara.IsNil(request.TerminalGroupId) {
		body["TerminalGroupId"] = request.TerminalGroupId
	}

	if !dara.IsNil(request.Type) {
		body["Type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOrCreateInvitationCode"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOrCreateInvitationCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取桌面端纳管邀请码
//
// @param request - GetOrCreateInvitationCodeRequest
//
// @return GetOrCreateInvitationCodeResponse
func (client *Client) GetOrCreateInvitationCode(request *GetOrCreateInvitationCodeRequest) (_result *GetOrCreateInvitationCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetOrCreateInvitationCodeResponse{}
	_body, _err := client.GetOrCreateInvitationCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询终端列表
//
// @param request - ListTerminalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTerminalResponse
func (client *Client) ListTerminalWithOptions(request *ListTerminalRequest, runtime *dara.RuntimeOptions) (_result *ListTerminalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Alias) {
		body["Alias"] = request.Alias
	}

	if !dara.IsNil(request.BuildId) {
		body["BuildId"] = request.BuildId
	}

	if !dara.IsNil(request.ClientType) {
		body["ClientType"] = request.ClientType
	}

	if !dara.IsNil(request.InManage) {
		body["InManage"] = request.InManage
	}

	if !dara.IsNil(request.Ipv4) {
		body["Ipv4"] = request.Ipv4
	}

	if !dara.IsNil(request.LocationInfo) {
		body["LocationInfo"] = request.LocationInfo
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.Model) {
		body["Model"] = request.Model
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.SearchKeyword) {
		body["SearchKeyword"] = request.SearchKeyword
	}

	if !dara.IsNil(request.SerialNumber) {
		body["SerialNumber"] = request.SerialNumber
	}

	if !dara.IsNil(request.TerminalGroupId) {
		body["TerminalGroupId"] = request.TerminalGroupId
	}

	if !dara.IsNil(request.Uuid) {
		body["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTerminal"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTerminalResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询终端列表
//
// @param request - ListTerminalRequest
//
// @return ListTerminalResponse
func (client *Client) ListTerminal(request *ListTerminalRequest) (_result *ListTerminalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ListTerminalResponse{}
	_body, _err := client.ListTerminalWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 向终端发送运维命令
//
// @param request - SendOpsMessageToTerminalsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendOpsMessageToTerminalsResponse
func (client *Client) SendOpsMessageToTerminalsWithOptions(request *SendOpsMessageToTerminalsRequest, runtime *dara.RuntimeOptions) (_result *SendOpsMessageToTerminalsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Delay) {
		query["Delay"] = request.Delay
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Msg) {
		body["Msg"] = request.Msg
	}

	if !dara.IsNil(request.OpsAction) {
		body["OpsAction"] = request.OpsAction
	}

	bodyFlat := map[string]interface{}{}
	if !dara.IsNil(request.Uuids) {
		bodyFlat["Uuids"] = request.Uuids
	}

	if !dara.IsNil(request.WaitForAck) {
		body["WaitForAck"] = request.WaitForAck
	}

	body = dara.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendOpsMessageToTerminals"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &SendOpsMessageToTerminalsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 向终端发送运维命令
//
// @param request - SendOpsMessageToTerminalsRequest
//
// @return SendOpsMessageToTerminalsResponse
func (client *Client) SendOpsMessageToTerminals(request *SendOpsMessageToTerminalsRequest) (_result *SendOpsMessageToTerminalsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &SendOpsMessageToTerminalsResponse{}
	_body, _err := client.SendOpsMessageToTerminalsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 解绑免账号登录用户
//
// @param request - UnbindAccountLessLoginUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindAccountLessLoginUserResponse
func (client *Client) UnbindAccountLessLoginUserWithOptions(request *UnbindAccountLessLoginUserRequest, runtime *dara.RuntimeOptions) (_result *UnbindAccountLessLoginUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.SerialNumber) {
		body["SerialNumber"] = request.SerialNumber
	}

	if !dara.IsNil(request.Uuid) {
		body["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindAccountLessLoginUser"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindAccountLessLoginUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解绑免账号登录用户
//
// @param request - UnbindAccountLessLoginUserRequest
//
// @return UnbindAccountLessLoginUserResponse
func (client *Client) UnbindAccountLessLoginUser(request *UnbindAccountLessLoginUserRequest) (_result *UnbindAccountLessLoginUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnbindAccountLessLoginUserResponse{}
	_body, _err := client.UnbindAccountLessLoginUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 解绑免密登录用户
//
// @param request - UnbindPasswordFreeLoginUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindPasswordFreeLoginUserResponse
func (client *Client) UnbindPasswordFreeLoginUserWithOptions(request *UnbindPasswordFreeLoginUserRequest, runtime *dara.RuntimeOptions) (_result *UnbindPasswordFreeLoginUserResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MainBizType) {
		body["MainBizType"] = request.MainBizType
	}

	if !dara.IsNil(request.SerialNumber) {
		body["SerialNumber"] = request.SerialNumber
	}

	if !dara.IsNil(request.Uuid) {
		body["Uuid"] = request.Uuid
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindPasswordFreeLoginUser"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindPasswordFreeLoginUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解绑免密登录用户
//
// @param request - UnbindPasswordFreeLoginUserRequest
//
// @return UnbindPasswordFreeLoginUserResponse
func (client *Client) UnbindPasswordFreeLoginUser(request *UnbindPasswordFreeLoginUserRequest) (_result *UnbindPasswordFreeLoginUserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &UnbindPasswordFreeLoginUserResponse{}
	_body, _err := client.UnbindPasswordFreeLoginUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
