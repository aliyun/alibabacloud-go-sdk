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
	client.Endpoint, _err = client.GetEndpoint(dara.String("captcha"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 验证码验证
//
// @param request - VerifyCaptchaRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VerifyCaptchaResponse
func (client *Client) VerifyCaptchaWithOptions(request *VerifyCaptchaRequest, runtime *dara.RuntimeOptions) (_result *VerifyCaptchaResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - VerifyCaptchaRequest
//
// @return VerifyCaptchaResponse
func (client *Client) VerifyCaptcha(request *VerifyCaptchaRequest) (_result *VerifyCaptchaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &VerifyCaptchaResponse{}
	_body, _err := client.VerifyCaptchaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) VerifyIntelligentCaptchaWithOptions(request *VerifyIntelligentCaptchaRequest, runtime *dara.RuntimeOptions) (_result *VerifyIntelligentCaptchaResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return VerifyIntelligentCaptchaResponse
func (client *Client) VerifyIntelligentCaptcha(request *VerifyIntelligentCaptchaRequest) (_result *VerifyIntelligentCaptchaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &VerifyIntelligentCaptchaResponse{}
	_body, _err := client.VerifyIntelligentCaptchaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
