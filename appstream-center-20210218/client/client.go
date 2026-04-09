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
	client.Endpoint, _err = client.GetEndpoint(dara.String("appstream-center"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Manually invalidates a logon token.
//
// @param request - ExpireLoginTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExpireLoginTokenResponse
func (client *Client) ExpireLoginTokenWithOptions(request *ExpireLoginTokenRequest, runtime *dara.RuntimeOptions) (_result *ExpireLoginTokenResponse, _err error) {
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

	if !dara.IsNil(request.LoginToken) {
		body["LoginToken"] = request.LoginToken
	}

	if !dara.IsNil(request.OfficeSiteId) {
		body["OfficeSiteId"] = request.OfficeSiteId
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExpireLoginToken"),
		Version:     dara.String("2021-02-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExpireLoginTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Manually invalidates a logon token.
//
// @param request - ExpireLoginTokenRequest
//
// @return ExpireLoginTokenResponse
func (client *Client) ExpireLoginToken(request *ExpireLoginTokenRequest) (_result *ExpireLoginTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExpireLoginTokenResponse{}
	_body, _err := client.ExpireLoginTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains an authorization code that includes the identity and permission information of a user. You can use the code to launch cloud apps in integration scenarios.
//
// @param request - GetAuthCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAuthCodeResponse
func (client *Client) GetAuthCodeWithOptions(request *GetAuthCodeRequest, runtime *dara.RuntimeOptions) (_result *GetAuthCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TokenType) {
		query["TokenType"] = request.TokenType
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoCreateUser) {
		body["AutoCreateUser"] = request.AutoCreateUser
	}

	if !dara.IsNil(request.EndUserId) {
		body["EndUserId"] = request.EndUserId
	}

	if !dara.IsNil(request.ExternalUserId) {
		body["ExternalUserId"] = request.ExternalUserId
	}

	if !dara.IsNil(request.Policy) {
		body["Policy"] = request.Policy
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAuthCode"),
		Version:     dara.String("2021-02-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAuthCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains an authorization code that includes the identity and permission information of a user. You can use the code to launch cloud apps in integration scenarios.
//
// @param request - GetAuthCodeRequest
//
// @return GetAuthCodeResponse
func (client *Client) GetAuthCode(request *GetAuthCodeRequest) (_result *GetAuthCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetAuthCodeResponse{}
	_body, _err := client.GetAuthCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取stsToken
//
// @param request - GetStsTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStsTokenResponse
func (client *Client) GetStsTokenWithOptions(request *GetStsTokenRequest, runtime *dara.RuntimeOptions) (_result *GetStsTokenResponse, _err error) {
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

	if !dara.IsNil(request.Expiration) {
		body["Expiration"] = request.Expiration
	}

	if !dara.IsNil(request.ExternalId) {
		body["ExternalId"] = request.ExternalId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetStsToken"),
		Version:     dara.String("2021-02-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetStsTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取stsToken
//
// @param request - GetStsTokenRequest
//
// @return GetStsTokenResponse
func (client *Client) GetStsToken(request *GetStsTokenRequest) (_result *GetStsTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetStsTokenResponse{}
	_body, _err := client.GetStsTokenWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
