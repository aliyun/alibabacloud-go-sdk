// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 主动交互消息传递
//
// @param request - ActiveInteractionCreateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActiveInteractionCreateResponse
func (client *Client) ActiveInteractionCreateWithContext(ctx context.Context, request *ActiveInteractionCreateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ActiveInteractionCreateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Image) {
		body["image"] = request.Image
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ActiveInteractionCreate"),
		Version:     dara.String("2024-08-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/open/api/v1/active/interaction/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ActiveInteractionCreateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设备注册
//
// @param request - DeviceRegisterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeviceRegisterResponse
func (client *Client) DeviceRegisterWithContext(ctx context.Context, request *DeviceRegisterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeviceRegisterResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["appId"] = request.AppId
	}

	if !dara.IsNil(request.Nonce) {
		body["nonce"] = request.Nonce
	}

	if !dara.IsNil(request.RequestTime) {
		body["requestTime"] = request.RequestTime
	}

	if !dara.IsNil(request.Signature) {
		body["signature"] = request.Signature
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeviceRegister"),
		Version:     dara.String("2024-08-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/open/api/device/v1/register"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &DeviceRegisterResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取网关校验Token
//
// @param request - GetTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTokenResponse
func (client *Client) GetTokenWithContext(ctx context.Context, request *GetTokenRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppId) {
		body["appId"] = request.AppId
	}

	if !dara.IsNil(request.DeviceName) {
		body["deviceName"] = request.DeviceName
	}

	if !dara.IsNil(request.Nonce) {
		body["nonce"] = request.Nonce
	}

	if !dara.IsNil(request.RequestTime) {
		body["requestTime"] = request.RequestTime
	}

	if !dara.IsNil(request.Signature) {
		body["signature"] = request.Signature
	}

	if !dara.IsNil(request.TokenKey) {
		body["tokenKey"] = request.TokenKey
	}

	if !dara.IsNil(request.TokenType) {
		body["tokenType"] = request.TokenType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetToken"),
		Version:     dara.String("2024-08-16"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/open/api/auth/v1/token/get"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
