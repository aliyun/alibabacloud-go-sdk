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
	client.Endpoint, _err = client.GetEndpoint(dara.String("maas"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 创建ApiKey
//
// @param request - CreateApiKeyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApiKeyResponse
func (client *Client) CreateApiKeyWithOptions(request *CreateApiKeyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateApiKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		query["description"] = request.Description
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApiKey"),
		Version:     dara.String("2026-03-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/bailianControl/apiKey/createApiKey"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateApiKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建ApiKey
//
// @param request - CreateApiKeyRequest
//
// @return CreateApiKeyResponse
func (client *Client) CreateApiKey(request *CreateApiKeyRequest) (_result *CreateApiKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateApiKeyResponse{}
	_body, _err := client.CreateApiKeyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除apiKey
//
// @param request - DeleteApiKeyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApiKeyResponse
func (client *Client) DeleteApiKeyWithOptions(request *DeleteApiKeyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteApiKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKeyId) {
		query["apiKeyId"] = request.ApiKeyId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApiKey"),
		Version:     dara.String("2026-03-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/bailianControl/apiKey/deleteApiKey"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApiKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除apiKey
//
// @param request - DeleteApiKeyRequest
//
// @return DeleteApiKeyResponse
func (client *Client) DeleteApiKey(request *DeleteApiKeyRequest) (_result *DeleteApiKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteApiKeyResponse{}
	_body, _err := client.DeleteApiKeyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询ApiKey详情
//
// @param request - GetApiKeyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApiKeyResponse
func (client *Client) GetApiKeyWithOptions(request *GetApiKeyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetApiKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKeyId) {
		query["apiKeyId"] = request.ApiKeyId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApiKey"),
		Version:     dara.String("2026-03-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/bailianControl/apiKey/getApiKey"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApiKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询ApiKey详情
//
// @param request - GetApiKeyRequest
//
// @return GetApiKeyResponse
func (client *Client) GetApiKey(request *GetApiKeyRequest) (_result *GetApiKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetApiKeyResponse{}
	_body, _err := client.GetApiKeyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取ApiKey列表
//
// @param request - ListApiKeysRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApiKeysResponse
func (client *Client) ListApiKeysWithOptions(request *ListApiKeysRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListApiKeysResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKeyId) {
		query["apiKeyId"] = request.ApiKeyId
	}

	if !dara.IsNil(request.Description) {
		query["description"] = request.Description
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApiKeys"),
		Version:     dara.String("2026-03-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/bailianControl/apiKeys"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApiKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取ApiKey列表
//
// @param request - ListApiKeysRequest
//
// @return ListApiKeysResponse
func (client *Client) ListApiKeys(request *ListApiKeysRequest) (_result *ListApiKeysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListApiKeysResponse{}
	_body, _err := client.ListApiKeysWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 编辑apiKey的描述
//
// @param request - UpdateApiKeyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApiKeyResponse
func (client *Client) UpdateApiKeyWithOptions(request *UpdateApiKeyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateApiKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiKeyId) {
		query["apiKeyId"] = request.ApiKeyId
	}

	if !dara.IsNil(request.Description) {
		query["description"] = request.Description
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateApiKey"),
		Version:     dara.String("2026-03-18"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/bailianControl/apiKey/updateApiKey"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateApiKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑apiKey的描述
//
// @param request - UpdateApiKeyRequest
//
// @return UpdateApiKeyResponse
func (client *Client) UpdateApiKey(request *UpdateApiKeyRequest) (_result *UpdateApiKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateApiKeyResponse{}
	_body, _err := client.UpdateApiKeyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
