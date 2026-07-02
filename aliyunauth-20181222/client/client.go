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
	client.Endpoint, _err = client.GetEndpoint(dara.String("aliyunauth"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// @param request - AuthenticateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthenticateResponse
func (client *Client) AuthenticateWithOptions(request *AuthenticateRequest, runtime *dara.RuntimeOptions) (_result *AuthenticateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Authenticate"),
		Version:     dara.String("2018-12-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &AuthenticateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - AuthenticateRequest
//
// @return AuthenticateResponse
func (client *Client) Authenticate(request *AuthenticateRequest) (_result *AuthenticateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &AuthenticateResponse{}
	_body, _err := client.AuthenticateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetDetailByOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDetailByOrderResponse
func (client *Client) GetDetailByOrderWithOptions(request *GetDetailByOrderRequest, runtime *dara.RuntimeOptions) (_result *GetDetailByOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Acceptor) {
		query["Acceptor"] = request.Acceptor
	}

	if !dara.IsNil(request.BizNo) {
		query["BizNo"] = request.BizNo
	}

	if !dara.IsNil(request.Channel) {
		query["Channel"] = request.Channel
	}

	if !dara.IsNil(request.CheckAuthItems) {
		query["CheckAuthItems"] = request.CheckAuthItems
	}

	if !dara.IsNil(request.EmpId) {
		query["EmpId"] = request.EmpId
	}

	if !dara.IsNil(request.Language) {
		query["Language"] = request.Language
	}

	if !dara.IsNil(request.OptSource) {
		query["OptSource"] = request.OptSource
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDetailByOrder"),
		Version:     dara.String("2018-12-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDetailByOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetDetailByOrderRequest
//
// @return GetDetailByOrderResponse
func (client *Client) GetDetailByOrder(request *GetDetailByOrderRequest) (_result *GetDetailByOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetDetailByOrderResponse{}
	_body, _err := client.GetDetailByOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryAuthRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAuthResponse
func (client *Client) QueryAuthWithOptions(request *QueryAuthRequest, runtime *dara.RuntimeOptions) (_result *QueryAuthResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAuth"),
		Version:     dara.String("2018-12-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAuthResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryAuthRequest
//
// @return QueryAuthResponse
func (client *Client) QueryAuth(request *QueryAuthRequest) (_result *QueryAuthResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryAuthResponse{}
	_body, _err := client.QueryAuthWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - QueryInEffectQuthOrderRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryInEffectQuthOrderResponse
func (client *Client) QueryInEffectQuthOrderWithOptions(request *QueryInEffectQuthOrderRequest, runtime *dara.RuntimeOptions) (_result *QueryInEffectQuthOrderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryInEffectQuthOrder"),
		Version:     dara.String("2018-12-22"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryInEffectQuthOrderResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - QueryInEffectQuthOrderRequest
//
// @return QueryInEffectQuthOrderResponse
func (client *Client) QueryInEffectQuthOrder(request *QueryInEffectQuthOrderRequest) (_result *QueryInEffectQuthOrderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &QueryInEffectQuthOrderResponse{}
	_body, _err := client.QueryInEffectQuthOrderWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
