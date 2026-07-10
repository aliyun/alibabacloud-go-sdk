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
	client.Endpoint, _err = client.GetEndpoint(dara.String("fcsandbox"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 创建 ApiKey
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
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApiKey"),
		Version:     dara.String("2026-05-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/2026-05-09/api-keys"),
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
// 创建 ApiKey
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
// 创建 Team
//
// @param request - CreateTeamRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTeamResponse
func (client *Client) CreateTeamWithOptions(request *CreateTeamRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTeamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTeam"),
		Version:     dara.String("2026-05-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/2026-05-09/teams"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTeamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建 Team
//
// @param request - CreateTeamRequest
//
// @return CreateTeamResponse
func (client *Client) CreateTeam(request *CreateTeamRequest) (_result *CreateTeamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTeamResponse{}
	_body, _err := client.CreateTeamWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除 ApiKey
//
// @param request - DeleteApiKeyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApiKeyResponse
func (client *Client) DeleteApiKeyWithOptions(apiKeyID *string, request *DeleteApiKeyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteApiKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApiKey"),
		Version:     dara.String("2026-05-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/2026-05-09/api-keys/" + dara.PercentEncode(dara.StringValue(apiKeyID))),
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
// 删除 ApiKey
//
// @param request - DeleteApiKeyRequest
//
// @return DeleteApiKeyResponse
func (client *Client) DeleteApiKey(apiKeyID *string, request *DeleteApiKeyRequest) (_result *DeleteApiKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteApiKeyResponse{}
	_body, _err := client.DeleteApiKeyWithOptions(apiKeyID, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除 quota 配置
//
// @param request - DeleteQuotaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteQuotaResponse
func (client *Client) DeleteQuotaWithOptions(request *DeleteQuotaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TagValue) {
		query["tagValue"] = request.TagValue
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteQuota"),
		Version:     dara.String("2026-05-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/2026-05-09/quotas/tag"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除 quota 配置
//
// @param request - DeleteQuotaRequest
//
// @return DeleteQuotaResponse
func (client *Client) DeleteQuota(request *DeleteQuotaRequest) (_result *DeleteQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteQuotaResponse{}
	_body, _err := client.DeleteQuotaWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除 Team
//
// @param request - DeleteTeamRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTeamResponse
func (client *Client) DeleteTeamWithOptions(teamID *string, request *DeleteTeamRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteTeamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTeam"),
		Version:     dara.String("2026-05-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/2026-05-09/teams/" + dara.PercentEncode(dara.StringValue(teamID))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTeamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除 Team
//
// @param request - DeleteTeamRequest
//
// @return DeleteTeamResponse
func (client *Client) DeleteTeam(teamID *string, request *DeleteTeamRequest) (_result *DeleteTeamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTeamResponse{}
	_body, _err := client.DeleteTeamWithOptions(teamID, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看 ApiKey
//
// @param request - DescribeApiKeyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeApiKeyResponse
func (client *Client) DescribeApiKeyWithOptions(apiKeyID *string, request *DescribeApiKeyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeApiKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeApiKey"),
		Version:     dara.String("2026-05-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/2026-05-09/api-keys/" + dara.PercentEncode(dara.StringValue(apiKeyID))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeApiKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看 ApiKey
//
// @param request - DescribeApiKeyRequest
//
// @return DescribeApiKeyResponse
func (client *Client) DescribeApiKey(apiKeyID *string, request *DescribeApiKeyRequest) (_result *DescribeApiKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeApiKeyResponse{}
	_body, _err := client.DescribeApiKeyWithOptions(apiKeyID, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取 quota 配置
//
// @param request - DescribeQuotaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeQuotaResponse
func (client *Client) DescribeQuotaWithOptions(request *DescribeQuotaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TagValue) {
		query["tagValue"] = request.TagValue
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeQuota"),
		Version:     dara.String("2026-05-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/2026-05-09/quotas/tag"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取 quota 配置
//
// @param request - DescribeQuotaRequest
//
// @return DescribeQuotaResponse
func (client *Client) DescribeQuota(request *DescribeQuotaRequest) (_result *DescribeQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeQuotaResponse{}
	_body, _err := client.DescribeQuotaWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取Team详情
//
// @param request - GetTeamRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTeamResponse
func (client *Client) GetTeamWithOptions(teamID *string, request *GetTeamRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTeamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTeam"),
		Version:     dara.String("2026-05-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/2026-05-09/teams/" + dara.PercentEncode(dara.StringValue(teamID))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTeamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Team详情
//
// @param request - GetTeamRequest
//
// @return GetTeamResponse
func (client *Client) GetTeam(teamID *string, request *GetTeamRequest) (_result *GetTeamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTeamResponse{}
	_body, _err := client.GetTeamWithOptions(teamID, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 分页查询 ApiKey
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
	if !dara.IsNil(request.ApiKeyName) {
		query["apiKeyName"] = request.ApiKeyName
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupID) {
		query["resourceGroupID"] = request.ResourceGroupID
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	if !dara.IsNil(request.TeamID) {
		query["teamID"] = request.TeamID
	}

	if !dara.IsNil(request.UserID) {
		query["userID"] = request.UserID
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApiKeys"),
		Version:     dara.String("2026-05-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/2026-05-09/api-keys"),
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
// 分页查询 ApiKey
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
// 查询 quota 配置
//
// @param request - ListQuotaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQuotaResponse
func (client *Client) ListQuotaWithOptions(request *ListQuotaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListQuota"),
		Version:     dara.String("2026-05-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/2026-05-09/quotas"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询 quota 配置
//
// @param request - ListQuotaRequest
//
// @return ListQuotaResponse
func (client *Client) ListQuota(request *ListQuotaRequest) (_result *ListQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListQuotaResponse{}
	_body, _err := client.ListQuotaWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询 Team 列表
//
// @param request - ListTeamsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTeamsResponse
func (client *Client) ListTeamsWithOptions(request *ListTeamsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTeamsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupID) {
		query["resourceGroupID"] = request.ResourceGroupID
	}

	if !dara.IsNil(request.TeamName) {
		query["teamName"] = request.TeamName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTeams"),
		Version:     dara.String("2026-05-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/2026-05-09/teams"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTeamsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询 Team 列表
//
// @param request - ListTeamsRequest
//
// @return ListTeamsResponse
func (client *Client) ListTeams(request *ListTeamsRequest) (_result *ListTeamsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTeamsResponse{}
	_body, _err := client.ListTeamsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 重置 ApiKey
//
// @param request - ResetApiKeyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResetApiKeyResponse
func (client *Client) ResetApiKeyWithOptions(apiKeyID *string, request *ResetApiKeyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ResetApiKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResetApiKey"),
		Version:     dara.String("2026-05-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/2026-05-09/api-keys/" + dara.PercentEncode(dara.StringValue(apiKeyID)) + "/reset"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ResetApiKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重置 ApiKey
//
// @param request - ResetApiKeyRequest
//
// @return ResetApiKeyResponse
func (client *Client) ResetApiKey(apiKeyID *string, request *ResetApiKeyRequest) (_result *ResetApiKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ResetApiKeyResponse{}
	_body, _err := client.ResetApiKeyWithOptions(apiKeyID, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新 ApiKey
//
// @param request - UpdateApiKeyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApiKeyResponse
func (client *Client) UpdateApiKeyWithOptions(apiKeyID *string, request *UpdateApiKeyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateApiKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateApiKey"),
		Version:     dara.String("2026-05-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/2026-05-09/api-keys/" + dara.PercentEncode(dara.StringValue(apiKeyID))),
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
// 更新 ApiKey
//
// @param request - UpdateApiKeyRequest
//
// @return UpdateApiKeyResponse
func (client *Client) UpdateApiKey(apiKeyID *string, request *UpdateApiKeyRequest) (_result *UpdateApiKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateApiKeyResponse{}
	_body, _err := client.UpdateApiKeyWithOptions(apiKeyID, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新 quota 配置
//
// @param request - UpdateQuotaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateQuotaResponse
func (client *Client) UpdateQuotaWithOptions(request *UpdateQuotaRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateQuotaResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateQuota"),
		Version:     dara.String("2026-05-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/2026-05-09/quotas/tag"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新 quota 配置
//
// @param request - UpdateQuotaRequest
//
// @return UpdateQuotaResponse
func (client *Client) UpdateQuota(request *UpdateQuotaRequest) (_result *UpdateQuotaResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateQuotaResponse{}
	_body, _err := client.UpdateQuotaWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新 Team
//
// @param request - UpdateTeamRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTeamResponse
func (client *Client) UpdateTeamWithOptions(teamID *string, request *UpdateTeamRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateTeamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTeam"),
		Version:     dara.String("2026-05-09"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pop/2026-05-09/teams/" + dara.PercentEncode(dara.StringValue(teamID))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTeamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新 Team
//
// @param request - UpdateTeamRequest
//
// @return UpdateTeamResponse
func (client *Client) UpdateTeam(teamID *string, request *UpdateTeamRequest) (_result *UpdateTeamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateTeamResponse{}
	_body, _err := client.UpdateTeamWithOptions(teamID, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
