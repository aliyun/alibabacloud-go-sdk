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
	client.Endpoint, _err = client.GetEndpoint(dara.String("bsn"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

// @param request - GetBsnByResourceRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBsnByResourceResponse
func (client *Client) GetBsnByResourceWithOptions(request *GetBsnByResourceRequest, runtime *dara.RuntimeOptions) (_result *GetBsnByResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Aliuid) {
		query["aliuid"] = request.Aliuid
	}

	if !dara.IsNil(request.ResourceId) {
		query["resourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["resourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBsnByResource"),
		Version:     dara.String("2015-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBsnByResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetBsnByResourceRequest
//
// @return GetBsnByResourceResponse
func (client *Client) GetBsnByResource(request *GetBsnByResourceRequest) (_result *GetBsnByResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GetBsnByResourceResponse{}
	_body, _err := client.GetBsnByResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 授权服务码，供虚商使用
//
// @param request - GrantBsnCodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GrantBsnCodeResponse
func (client *Client) GrantBsnCodeWithOptions(request *GrantBsnCodeRequest, runtime *dara.RuntimeOptions) (_result *GrantBsnCodeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AliUid) {
		query["AliUid"] = request.AliUid
	}

	if !dara.IsNil(request.GrantToAliuid) {
		query["GrantToAliuid"] = request.GrantToAliuid
	}

	if !dara.IsNil(request.Notes) {
		query["Notes"] = request.Notes
	}

	if !dara.IsNil(request.Sn) {
		query["Sn"] = request.Sn
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GrantBsnCode"),
		Version:     dara.String("2015-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GrantBsnCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 授权服务码，供虚商使用
//
// @param request - GrantBsnCodeRequest
//
// @return GrantBsnCodeResponse
func (client *Client) GrantBsnCode(request *GrantBsnCodeRequest) (_result *GrantBsnCodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &GrantBsnCodeResponse{}
	_body, _err := client.GrantBsnCodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ProductBindBsnRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ProductBindBsnResponse
func (client *Client) ProductBindBsnWithOptions(request *ProductBindBsnRequest, runtime *dara.RuntimeOptions) (_result *ProductBindBsnResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Aliuid) {
		query["aliuid"] = request.Aliuid
	}

	if !dara.IsNil(request.Num) {
		query["num"] = request.Num
	}

	if !dara.IsNil(request.ResourceId) {
		query["resourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["resourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ProductBindBsn"),
		Version:     dara.String("2015-05-12"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ProductBindBsnResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - ProductBindBsnRequest
//
// @return ProductBindBsnResponse
func (client *Client) ProductBindBsn(request *ProductBindBsnRequest) (_result *ProductBindBsnResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &ProductBindBsnResponse{}
	_body, _err := client.ProductBindBsnWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
