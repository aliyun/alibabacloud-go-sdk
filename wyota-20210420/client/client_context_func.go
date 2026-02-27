// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 添加终端
//
// @param request - AddTerminalRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTerminalResponse
func (client *Client) AddTerminalWithContext(ctx context.Context, request *AddTerminalRequest, runtime *dara.RuntimeOptions) (_result *AddTerminalResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTerminalsResponse
func (client *Client) AddTerminalsWithContext(ctx context.Context, request *AddTerminalsRequest, runtime *dara.RuntimeOptions) (_result *AddTerminalsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindAccountLessLoginUserResponse
func (client *Client) BindAccountLessLoginUserWithContext(ctx context.Context, request *BindAccountLessLoginUserRequest, runtime *dara.RuntimeOptions) (_result *BindAccountLessLoginUserResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindPasswordFreeLoginUserResponse
func (client *Client) BindPasswordFreeLoginUserWithContext(ctx context.Context, request *BindPasswordFreeLoginUserRequest, runtime *dara.RuntimeOptions) (_result *BindPasswordFreeLoginUserResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询设备座位
//
// @param request - DescribeDeviceSeatsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDeviceSeatsResponse
func (client *Client) DescribeDeviceSeatsWithContext(ctx context.Context, request *DescribeDeviceSeatsRequest, runtime *dara.RuntimeOptions) (_result *DescribeDeviceSeatsResponse, _err error) {
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

	if !dara.IsNil(request.SerialNo) {
		body["SerialNo"] = request.SerialNo
	}

	if !dara.IsNil(request.SerialNoList) {
		body["SerialNoList"] = request.SerialNoList
	}

	if !dara.IsNil(request.SiteId) {
		body["SiteId"] = request.SiteId
	}

	if !dara.IsNil(request.TenantId) {
		body["TenantId"] = request.TenantId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDeviceSeats"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDeviceSeatsResponse{}
	_body, _err := client.DoRPCRequestWithCtx(ctx, params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTerminalResponse
func (client *Client) ListTerminalWithContext(ctx context.Context, request *ListTerminalRequest, runtime *dara.RuntimeOptions) (_result *ListTerminalResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendOpsMessageToTerminalsResponse
func (client *Client) SendOpsMessageToTerminalsWithContext(ctx context.Context, request *SendOpsMessageToTerminalsRequest, runtime *dara.RuntimeOptions) (_result *SendOpsMessageToTerminalsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindAccountLessLoginUserResponse
func (client *Client) UnbindAccountLessLoginUserWithContext(ctx context.Context, request *UnbindAccountLessLoginUserRequest, runtime *dara.RuntimeOptions) (_result *UnbindAccountLessLoginUserResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 解绑设备座位
//
// @param tmpReq - UnbindDeviceSeatsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindDeviceSeatsResponse
func (client *Client) UnbindDeviceSeatsWithContext(ctx context.Context, tmpReq *UnbindDeviceSeatsRequest, runtime *dara.RuntimeOptions) (_result *UnbindDeviceSeatsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &UnbindDeviceSeatsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SerialNoList) {
		request.SerialNoListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SerialNoList, dara.String("SerialNoList"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SerialNoListShrink) {
		body["SerialNoList"] = request.SerialNoListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UnbindDeviceSeats"),
		Version:     dara.String("2021-04-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &UnbindDeviceSeatsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindPasswordFreeLoginUserResponse
func (client *Client) UnbindPasswordFreeLoginUserWithContext(ctx context.Context, request *UnbindPasswordFreeLoginUserRequest, runtime *dara.RuntimeOptions) (_result *UnbindPasswordFreeLoginUserResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
