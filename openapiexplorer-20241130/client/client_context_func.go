// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 验证 Terraform HCL 语法
//
// @param request - ApiMcpServerValidateHclRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApiMcpServerValidateHclResponse
func (client *Client) ApiMcpServerValidateHclWithContext(ctx context.Context, request *ApiMcpServerValidateHclRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ApiMcpServerValidateHclResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		body["code"] = request.Code
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApiMcpServerValidateHcl"),
		Version:     dara.String("2024-11-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/terraform/validate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ApiMcpServerValidateHclResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建ApiMcpServer
//
// @param request - CreateApiMcpServerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateApiMcpServerResponse
func (client *Client) CreateApiMcpServerWithContext(ctx context.Context, request *CreateApiMcpServerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateApiMcpServerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AdditionalApiDescriptions) {
		body["additionalApiDescriptions"] = request.AdditionalApiDescriptions
	}

	if !dara.IsNil(request.Apis) {
		body["apis"] = request.Apis
	}

	if !dara.IsNil(request.AssumeRoleExtraPolicy) {
		body["assumeRoleExtraPolicy"] = request.AssumeRoleExtraPolicy
	}

	if !dara.IsNil(request.AssumeRoleName) {
		body["assumeRoleName"] = request.AssumeRoleName
	}

	if !dara.IsNil(request.ClientToken) {
		body["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.EnableAssumeRole) {
		body["enableAssumeRole"] = request.EnableAssumeRole
	}

	if !dara.IsNil(request.EnableCustomVpcWhitelist) {
		body["enableCustomVpcWhitelist"] = request.EnableCustomVpcWhitelist
	}

	if !dara.IsNil(request.Instructions) {
		body["instructions"] = request.Instructions
	}

	if !dara.IsNil(request.Language) {
		body["language"] = request.Language
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.OauthClientId) {
		body["oauthClientId"] = request.OauthClientId
	}

	if !dara.IsNil(request.Prompts) {
		body["prompts"] = request.Prompts
	}

	if !dara.IsNil(request.PublicAccess) {
		body["publicAccess"] = request.PublicAccess
	}

	if !dara.IsNil(request.SystemTools) {
		body["systemTools"] = request.SystemTools
	}

	if !dara.IsNil(request.TerraformTools) {
		body["terraformTools"] = request.TerraformTools
	}

	if !dara.IsNil(request.VpcWhitelists) {
		body["vpcWhitelists"] = request.VpcWhitelists
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateApiMcpServer"),
		Version:     dara.String("2024-11-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/apimcpserver"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateApiMcpServerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除ApiMcpServer
//
// @param request - DeleteApiMcpServerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteApiMcpServerResponse
func (client *Client) DeleteApiMcpServerWithContext(ctx context.Context, request *DeleteApiMcpServerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteApiMcpServerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteApiMcpServer"),
		Version:     dara.String("2024-11-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/apimcpserver"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteApiMcpServerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 动态生成Aliyun CLI命令
//
// @param tmpReq - GenerateCLICommandRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateCLICommandResponse
func (client *Client) GenerateCLICommandWithContext(ctx context.Context, tmpReq *GenerateCLICommandRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GenerateCLICommandResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &GenerateCLICommandShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ApiParams) {
		request.ApiParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ApiParams, dara.String("apiParams"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Api) {
		body["api"] = request.Api
	}

	if !dara.IsNil(request.ApiParamsShrink) {
		body["apiParams"] = request.ApiParamsShrink
	}

	if !dara.IsNil(request.ApiVersion) {
		body["apiVersion"] = request.ApiVersion
	}

	if !dara.IsNil(request.JsonApiParams) {
		body["jsonApiParams"] = request.JsonApiParams
	}

	if !dara.IsNil(request.Product) {
		body["product"] = request.Product
	}

	if !dara.IsNil(request.RegionId) {
		body["regionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateCLICommand"),
		Version:     dara.String("2024-11-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/cli/makeCode"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateCLICommandResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取产品相关接口的开放元数据
//
// @param request - GetApiDefinitionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApiDefinitionResponse
func (client *Client) GetApiDefinitionWithContext(ctx context.Context, request *GetApiDefinitionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetApiDefinitionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Api) {
		query["api"] = request.Api
	}

	if !dara.IsNil(request.ApiVersion) {
		query["apiVersion"] = request.ApiVersion
	}

	if !dara.IsNil(request.Product) {
		query["product"] = request.Product
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApiDefinition"),
		Version:     dara.String("2024-11-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/definition"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApiDefinitionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询 ApiMcpServer
//
// @param request - GetApiMcpServerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApiMcpServerResponse
func (client *Client) GetApiMcpServerWithContext(ctx context.Context, request *GetApiMcpServerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetApiMcpServerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApiMcpServer"),
		Version:     dara.String("2024-11-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/apimcpserver"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApiMcpServerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询用户全局API MCP Server配置
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetApiMcpServerUserConfigResponse
func (client *Client) GetApiMcpServerUserConfigWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetApiMcpServerUserConfigResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetApiMcpServerUserConfig"),
		Version:     dara.String("2024-11-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/userconfig/get"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetApiMcpServerUserConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries an error solution by error code.
//
// Description:
//
// You can call this API operation to query public information instead of special information, such as the account ownership. Permissions on this API operation cannot be granted to other members.
//
// @param request - GetErrorCodeSolutionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetErrorCodeSolutionsResponse
func (client *Client) GetErrorCodeSolutionsWithContext(ctx context.Context, request *GetErrorCodeSolutionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetErrorCodeSolutionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		query["acceptLanguage"] = request.AcceptLanguage
	}

	if !dara.IsNil(request.ErrorCode) {
		query["errorCode"] = request.ErrorCode
	}

	if !dara.IsNil(request.ErrorMessage) {
		query["errorMessage"] = request.ErrorMessage
	}

	if !dara.IsNil(request.Product) {
		query["product"] = request.Product
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetErrorCodeSolutions"),
		Version:     dara.String("2024-11-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/getErrorCodeSolutions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetErrorCodeSolutionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the log of an API call performed by using the current account based on the returned request ID of the API to troubleshoot issues.
//
// Description:
//
// Permissions on this API cannot be granted to other members.
//
// @param request - GetOwnRequestLogRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOwnRequestLogResponse
func (client *Client) GetOwnRequestLogWithContext(ctx context.Context, request *GetOwnRequestLogRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetOwnRequestLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LogRequestId) {
		query["logRequestId"] = request.LogRequestId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOwnRequestLog"),
		Version:     dara.String("2024-11-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/getOwnRequestLog"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOwnRequestLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取产品的接入点信息
//
// @param request - GetProductEndpointsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProductEndpointsResponse
func (client *Client) GetProductEndpointsWithContext(ctx context.Context, request *GetProductEndpointsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetProductEndpointsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Product) {
		query["product"] = request.Product
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProductEndpoints"),
		Version:     dara.String("2024-11-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/product/endpoints"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProductEndpointsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the log of an API call based on the returned request ID of the API to troubleshoot issues.
//
// Description:
//
// You can grant permissions to a Resource Access Management (RAM) user or assume a role to query the log of an API call across RAM users or Alibaba Cloud accounts. For more information, see [Grant permissions to troubleshoot API errors across accounts](https://help.aliyun.com/document_detail/2868101.html).
//
// @param request - GetRequestLogRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRequestLogResponse
func (client *Client) GetRequestLogWithContext(ctx context.Context, request *GetRequestLogRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRequestLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LogRequestId) {
		query["logRequestId"] = request.LogRequestId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetRequestLog"),
		Version:     dara.String("2024-11-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/getRequestLog"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRequestLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取产品的开放元数据
//
// @param request - ListApiDefinitionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApiDefinitionsResponse
func (client *Client) ListApiDefinitionsWithContext(ctx context.Context, request *ListApiDefinitionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListApiDefinitionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ApiVersion) {
		query["apiVersion"] = request.ApiVersion
	}

	if !dara.IsNil(request.Product) {
		query["product"] = request.Product
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApiDefinitions"),
		Version:     dara.String("2024-11-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/definitions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApiDefinitionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询系统工具列表
//
// @param request - ListApiMcpServerSystemToolsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApiMcpServerSystemToolsResponse
func (client *Client) ListApiMcpServerSystemToolsWithContext(ctx context.Context, request *ListApiMcpServerSystemToolsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListApiMcpServerSystemToolsResponse, _err error) {
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

	if !dara.IsNil(request.Skip) {
		query["skip"] = request.Skip
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApiMcpServerSystemTools"),
		Version:     dara.String("2024-11-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/mcpSystemTools"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApiMcpServerSystemToolsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出资源ApiMcpServer
//
// @param request - ListApiMcpServersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListApiMcpServersResponse
func (client *Client) ListApiMcpServersWithContext(ctx context.Context, request *ListApiMcpServersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListApiMcpServersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CreateTime) {
		query["createTime"] = request.CreateTime
	}

	if !dara.IsNil(request.Description) {
		query["description"] = request.Description
	}

	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.Language) {
		query["language"] = request.Language
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Skip) {
		query["skip"] = request.Skip
	}

	if !dara.IsNil(request.SourceType) {
		query["sourceType"] = request.SourceType
	}

	if !dara.IsNil(request.UpdateTime) {
		query["updateTime"] = request.UpdateTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListApiMcpServers"),
		Version:     dara.String("2024-11-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/apimcpservers"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListApiMcpServersResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新UpdateApiMcpServer
//
// @param request - UpdateApiMcpServerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApiMcpServerResponse
func (client *Client) UpdateApiMcpServerWithContext(ctx context.Context, request *UpdateApiMcpServerRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateApiMcpServerResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.Id) {
		query["id"] = request.Id
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AdditionalApiDescriptions) {
		body["additionalApiDescriptions"] = request.AdditionalApiDescriptions
	}

	if !dara.IsNil(request.Apis) {
		body["apis"] = request.Apis
	}

	if !dara.IsNil(request.AssumeRoleExtraPolicy) {
		body["assumeRoleExtraPolicy"] = request.AssumeRoleExtraPolicy
	}

	if !dara.IsNil(request.AssumeRoleName) {
		body["assumeRoleName"] = request.AssumeRoleName
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.EnableAssumeRole) {
		body["enableAssumeRole"] = request.EnableAssumeRole
	}

	if !dara.IsNil(request.EnableCustomVpcWhitelist) {
		body["enableCustomVpcWhitelist"] = request.EnableCustomVpcWhitelist
	}

	if !dara.IsNil(request.Instructions) {
		body["instructions"] = request.Instructions
	}

	if !dara.IsNil(request.Language) {
		body["language"] = request.Language
	}

	if !dara.IsNil(request.OauthClientId) {
		body["oauthClientId"] = request.OauthClientId
	}

	if !dara.IsNil(request.Prompts) {
		body["prompts"] = request.Prompts
	}

	if !dara.IsNil(request.PublicAccess) {
		body["publicAccess"] = request.PublicAccess
	}

	if !dara.IsNil(request.SystemTools) {
		body["systemTools"] = request.SystemTools
	}

	if !dara.IsNil(request.TerraformTools) {
		body["terraformTools"] = request.TerraformTools
	}

	if !dara.IsNil(request.VpcWhitelists) {
		body["vpcWhitelists"] = request.VpcWhitelists
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateApiMcpServer"),
		Version:     dara.String("2024-11-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/apimcpserver"),
		Method:      dara.String("PATCH"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateApiMcpServerResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改用户全局API MCP Server配置
//
// @param request - UpdateApiMcpServerUserConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateApiMcpServerUserConfigResponse
func (client *Client) UpdateApiMcpServerUserConfigWithContext(ctx context.Context, request *UpdateApiMcpServerUserConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateApiMcpServerUserConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EnablePublicAccess) {
		body["enablePublicAccess"] = request.EnablePublicAccess
	}

	if !dara.IsNil(request.VpcWhitelists) {
		body["vpcWhitelists"] = request.VpcWhitelists
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateApiMcpServerUserConfig"),
		Version:     dara.String("2024-11-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/userconfig/update"),
		Method:      dara.String("PATCH"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateApiMcpServerUserConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
