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
		"cn-hangzhou": dara.String("saf.cn-shanghai.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("saf"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Detects risks. This API operation supports public service invocation and allows you to invoke custom events.
//
// @param request - ExecuteRequestRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteRequestResponse
func (client *Client) ExecuteRequestWithOptions(request *ExecuteRequestRequest, runtime *dara.RuntimeOptions) (_result *ExecuteRequestResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Service) {
		query["Service"] = request.Service
	}

	if !dara.IsNil(request.ServiceParameters) {
		query["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteRequest"),
		Version:     dara.String("2019-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteRequestResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detects risks. This API operation supports public service invocation and allows you to invoke custom events.
//
// @param request - ExecuteRequestRequest
//
// @return ExecuteRequestResponse
func (client *Client) ExecuteRequest(request *ExecuteRequestRequest) (_result *ExecuteRequestResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExecuteRequestResponse{}
	_body, _err := client.ExecuteRequestWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Detects risks. This operation supports only custom event calls and is available only in the Malaysia region.
//
// @param request - ExecuteRequestMLRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteRequestMLResponse
func (client *Client) ExecuteRequestMLWithOptions(request *ExecuteRequestMLRequest, runtime *dara.RuntimeOptions) (_result *ExecuteRequestMLResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Service) {
		query["Service"] = request.Service
	}

	if !dara.IsNil(request.ServiceParameters) {
		query["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteRequestML"),
		Version:     dara.String("2019-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteRequestMLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detects risks. This operation supports only custom event calls and is available only in the Malaysia region.
//
// @param request - ExecuteRequestMLRequest
//
// @return ExecuteRequestMLResponse
func (client *Client) ExecuteRequestML(request *ExecuteRequestMLRequest) (_result *ExecuteRequestMLResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExecuteRequestMLResponse{}
	_body, _err := client.ExecuteRequestMLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Detects risks. This API allows you to invoke public services and custom events.
//
// @param request - ExecuteRequestSGRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteRequestSGResponse
func (client *Client) ExecuteRequestSGWithOptions(request *ExecuteRequestSGRequest, runtime *dara.RuntimeOptions) (_result *ExecuteRequestSGResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Lang) {
		query["Lang"] = request.Lang
	}

	if !dara.IsNil(request.Service) {
		query["Service"] = request.Service
	}

	if !dara.IsNil(request.ServiceParameters) {
		query["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteRequestSG"),
		Version:     dara.String("2019-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteRequestSGResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detects risks. This API allows you to invoke public services and custom events.
//
// @param request - ExecuteRequestSGRequest
//
// @return ExecuteRequestSGResponse
func (client *Client) ExecuteRequestSG(request *ExecuteRequestSGRequest) (_result *ExecuteRequestSGResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ExecuteRequestSGResponse{}
	_body, _err := client.ExecuteRequestSGWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Detects risks. This operation supports only custom event calls.
//
// @param request - RequestDecisionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RequestDecisionResponse
func (client *Client) RequestDecisionWithOptions(request *RequestDecisionRequest, runtime *dara.RuntimeOptions) (_result *RequestDecisionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EventCode) {
		query["EventCode"] = request.EventCode
	}

	if !dara.IsNil(request.ServiceParameters) {
		query["ServiceParameters"] = request.ServiceParameters
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RequestDecision"),
		Version:     dara.String("2019-05-21"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &RequestDecisionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Detects risks. This operation supports only custom event calls.
//
// @param request - RequestDecisionRequest
//
// @return RequestDecisionResponse
func (client *Client) RequestDecision(request *RequestDecisionRequest) (_result *RequestDecisionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &RequestDecisionResponse{}
	_body, _err := client.RequestDecisionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
