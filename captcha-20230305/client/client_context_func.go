// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 验证码验证
//
// @param request - VerifyCaptchaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyCaptchaResponse
func (client *Client) VerifyCaptchaWithContext(ctx context.Context, request *VerifyCaptchaRequest, runtime *dara.RuntimeOptions) (_result *VerifyCaptchaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CaptchaVerifyParam) {
		query["CaptchaVerifyParam"] = request.CaptchaVerifyParam
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifyCaptcha"),
		Version:     dara.String("2023-03-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyCaptchaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 验证码验证
//
// @param request - VerifyIntelligentCaptchaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyIntelligentCaptchaResponse
func (client *Client) VerifyIntelligentCaptchaWithContext(ctx context.Context, request *VerifyIntelligentCaptchaRequest, runtime *dara.RuntimeOptions) (_result *VerifyIntelligentCaptchaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CaptchaVerifyParam) {
		body["CaptchaVerifyParam"] = request.CaptchaVerifyParam
	}

	if !dara.IsNil(request.SceneId) {
		body["SceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("VerifyIntelligentCaptcha"),
		Version:     dara.String("2023-03-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &VerifyIntelligentCaptchaResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
