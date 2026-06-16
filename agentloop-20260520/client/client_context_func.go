// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 给记忆库中增加数据
//
// @param request - AddMem0MemoriesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddMem0MemoriesResponse
func (client *Client) AddMem0MemoriesWithContext(ctx context.Context, request *AddMem0MemoriesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddMem0MemoriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentSpace) {
		query["agentSpace"] = request.AgentSpace
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Body) {
		body["body"] = request.Body
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddMem0Memories"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/memories"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AddMem0MemoriesResponse{}
	_body, _err := client.DoROARequestWithCtx(ctx, params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建AgentSpace
//
// @param request - CreateAgentSpaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAgentSpaceResponse
func (client *Client) CreateAgentSpaceWithContext(ctx context.Context, request *CreateAgentSpaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAgentSpaceResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentSpace) {
		body["agentSpace"] = request.AgentSpace
	}

	if !dara.IsNil(request.CmsWorkspace) {
		body["cmsWorkspace"] = request.CmsWorkspace
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAgentSpace"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/agentspace"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAgentSpaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建上下文库
//
// @param request - CreateContextStoreRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateContextStoreResponse
func (client *Client) CreateContextStoreWithContext(ctx context.Context, agentSpace *string, request *CreateContextStoreRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateContextStoreResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["config"] = request.Config
	}

	if !dara.IsNil(request.ContextStoreName) {
		body["contextStoreName"] = request.ContextStoreName
	}

	if !dara.IsNil(request.ContextType) {
		body["contextType"] = request.ContextType
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateContextStore"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/agentspace/" + dara.PercentEncode(dara.StringValue(agentSpace)) + "/contextstore"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateContextStoreResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建 API Key
//
// @param request - CreateContextStoreAPIKeyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateContextStoreAPIKeyResponse
func (client *Client) CreateContextStoreAPIKeyWithContext(ctx context.Context, agentSpace *string, contextStoreName *string, request *CreateContextStoreAPIKeyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateContextStoreAPIKeyResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateContextStoreAPIKey"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/agentspace/" + dara.PercentEncode(dara.StringValue(agentSpace)) + "/contextstore/" + dara.PercentEncode(dara.StringValue(contextStoreName)) + "/apikey"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateContextStoreAPIKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建数据集
//
// @param request - CreateDatasetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDatasetResponse
func (client *Client) CreateDatasetWithContext(ctx context.Context, agentSpace *string, request *CreateDatasetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDatasetResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.DatasetName) {
		body["datasetName"] = request.DatasetName
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Schema) {
		body["schema"] = request.Schema
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDataset"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/agentspace/" + dara.PercentEncode(dara.StringValue(agentSpace)) + "/dataset"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDatasetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除AgentSpace
//
// @param request - DeleteAgentSpaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAgentSpaceResponse
func (client *Client) DeleteAgentSpaceWithContext(ctx context.Context, agentSpace *string, request *DeleteAgentSpaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAgentSpaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeleteCmsWorkspace) {
		query["deleteCmsWorkspace"] = request.DeleteCmsWorkspace
	}

	if !dara.IsNil(request.DeleteMseNamespace) {
		query["deleteMseNamespace"] = request.DeleteMseNamespace
	}

	if !dara.IsNil(request.DeleteSlsProject) {
		query["deleteSlsProject"] = request.DeleteSlsProject
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAgentSpace"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/agentspace/" + dara.PercentEncode(dara.StringValue(agentSpace))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAgentSpaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除上下文库
//
// @param request - DeleteContextStoreRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteContextStoreResponse
func (client *Client) DeleteContextStoreWithContext(ctx context.Context, agentSpace *string, contextStoreName *string, request *DeleteContextStoreRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteContextStoreResponse, _err error) {
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
		Action:      dara.String("DeleteContextStore"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/agentspace/" + dara.PercentEncode(dara.StringValue(agentSpace)) + "/contextstore/" + dara.PercentEncode(dara.StringValue(contextStoreName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteContextStoreResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除 API Key
//
// @param request - DeleteContextStoreAPIKeyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteContextStoreAPIKeyResponse
func (client *Client) DeleteContextStoreAPIKeyWithContext(ctx context.Context, agentSpace *string, contextStoreName *string, name *string, request *DeleteContextStoreAPIKeyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteContextStoreAPIKeyResponse, _err error) {
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
		Action:      dara.String("DeleteContextStoreAPIKey"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/agentspace/" + dara.PercentEncode(dara.StringValue(agentSpace)) + "/contextstore/" + dara.PercentEncode(dara.StringValue(contextStoreName)) + "/apikey/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteContextStoreAPIKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除数据集
//
// @param request - DeleteDatasetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDatasetResponse
func (client *Client) DeleteDatasetWithContext(ctx context.Context, agentSpace *string, datasetName *string, request *DeleteDatasetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDatasetResponse, _err error) {
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
		Action:      dara.String("DeleteDataset"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/agentspace/" + dara.PercentEncode(dara.StringValue(agentSpace)) + "/dataset/" + dara.PercentEncode(dara.StringValue(datasetName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDatasetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量删除记忆
//
// @param tmpReq - DeleteMem0MemoriesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMem0MemoriesResponse
func (client *Client) DeleteMem0MemoriesWithContext(ctx context.Context, tmpReq *DeleteMem0MemoriesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteMem0MemoriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &DeleteMem0MemoriesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Metadata) {
		request.MetadataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Metadata, dara.String("metadata"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentSpace) {
		query["agentSpace"] = request.AgentSpace
	}

	if !dara.IsNil(request.AgentId) {
		query["agent_id"] = request.AgentId
	}

	if !dara.IsNil(request.AppId) {
		query["app_id"] = request.AppId
	}

	if !dara.IsNil(request.ContextStoreId) {
		query["context_store_id"] = request.ContextStoreId
	}

	if !dara.IsNil(request.MetadataShrink) {
		query["metadata"] = request.MetadataShrink
	}

	if !dara.IsNil(request.OrgId) {
		query["org_id"] = request.OrgId
	}

	if !dara.IsNil(request.ProjectId) {
		query["project_id"] = request.ProjectId
	}

	if !dara.IsNil(request.RunId) {
		query["run_id"] = request.RunId
	}

	if !dara.IsNil(request.UserId) {
		query["user_id"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMem0Memories"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/memories"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMem0MemoriesResponse{}
	_body, _err := client.DoROARequestWithCtx(ctx, params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除记忆
//
// @param request - DeleteMem0MemoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMem0MemoryResponse
func (client *Client) DeleteMem0MemoryWithContext(ctx context.Context, memoryId *string, request *DeleteMem0MemoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteMem0MemoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentSpace) {
		query["agentSpace"] = request.AgentSpace
	}

	if !dara.IsNil(request.ContextStoreId) {
		query["context_store_id"] = request.ContextStoreId
	}

	if !dara.IsNil(request.OrgId) {
		query["org_id"] = request.OrgId
	}

	if !dara.IsNil(request.ProjectId) {
		query["project_id"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMem0Memory"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/memories/" + dara.PercentEncode(dara.StringValue(memoryId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMem0MemoryResponse{}
	_body, _err := client.DoROARequestWithCtx(ctx, params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除流水线
//
// @param request - DeletePipelineRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePipelineResponse
func (client *Client) DeletePipelineWithContext(ctx context.Context, agentSpace *string, pipelineName *string, request *DeletePipelineRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePipelineResponse, _err error) {
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
		Action:      dara.String("DeletePipeline"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/agentspace/" + dara.PercentEncode(dara.StringValue(agentSpace)) + "/pipeline/" + dara.PercentEncode(dara.StringValue(pipelineName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeletePipelineResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Regions
//
// @param request - DescribeRegionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithContext(ctx context.Context, request *DescribeRegionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Language) {
		query["language"] = request.Language
	}

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
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/regions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 执行查询语句
//
// @param request - ExecuteQueryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteQueryResponse
func (client *Client) ExecuteQueryWithContext(ctx context.Context, agentSpace *string, datasetName *string, request *ExecuteQueryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecuteQueryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Query) {
		body["query"] = request.Query
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteQuery"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/agentspace/" + dara.PercentEncode(dara.StringValue(agentSpace)) + "/dataset/" + dara.PercentEncode(dara.StringValue(datasetName)) + "/query"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteQueryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询AgentSpace
//
// @param request - GetAgentSpaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentSpaceResponse
func (client *Client) GetAgentSpaceWithContext(ctx context.Context, agentSpace *string, request *GetAgentSpaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAgentSpaceResponse, _err error) {
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
		Action:      dara.String("GetAgentSpace"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/agentspace/" + dara.PercentEncode(dara.StringValue(agentSpace))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgentSpaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询上下文库
//
// @param request - GetContextStoreRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetContextStoreResponse
func (client *Client) GetContextStoreWithContext(ctx context.Context, agentSpace *string, contextStoreName *string, request *GetContextStoreRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetContextStoreResponse, _err error) {
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
		Action:      dara.String("GetContextStore"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/agentspace/" + dara.PercentEncode(dara.StringValue(agentSpace)) + "/contextstore/" + dara.PercentEncode(dara.StringValue(contextStoreName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetContextStoreResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取 API Key
//
// @param request - GetContextStoreAPIKeyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetContextStoreAPIKeyResponse
func (client *Client) GetContextStoreAPIKeyWithContext(ctx context.Context, agentSpace *string, contextStoreName *string, name *string, request *GetContextStoreAPIKeyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetContextStoreAPIKeyResponse, _err error) {
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
		Action:      dara.String("GetContextStoreAPIKey"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/agentspace/" + dara.PercentEncode(dara.StringValue(agentSpace)) + "/contextstore/" + dara.PercentEncode(dara.StringValue(contextStoreName)) + "/apikey/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetContextStoreAPIKeyResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询数据集
//
// @param request - GetDatasetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDatasetResponse
func (client *Client) GetDatasetWithContext(ctx context.Context, agentSpace *string, datasetName *string, request *GetDatasetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDatasetResponse, _err error) {
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
		Action:      dara.String("GetDataset"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/agentspace/" + dara.PercentEncode(dara.StringValue(agentSpace)) + "/dataset/" + dara.PercentEncode(dara.StringValue(datasetName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDatasetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询记忆库数据
//
// @param request - GetMem0MemoriesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMem0MemoriesResponse
func (client *Client) GetMem0MemoriesWithContext(ctx context.Context, request *GetMem0MemoriesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMem0MemoriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentSpace) {
		query["agentSpace"] = request.AgentSpace
	}

	if !dara.IsNil(request.ContextStoreId) {
		query["context_store_id"] = request.ContextStoreId
	}

	if !dara.IsNil(request.EnableGraph) {
		query["enable_graph"] = request.EnableGraph
	}

	if !dara.IsNil(request.OrgId) {
		query["org_id"] = request.OrgId
	}

	if !dara.IsNil(request.ProjectId) {
		query["project_id"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Body) {
		body["body"] = request.Body
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMem0Memories"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/memories"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("array"),
	}
	_result = &GetMem0MemoriesResponse{}
	_body, _err := client.DoROARequestWithCtx(ctx, params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询记忆
//
// @param request - GetMem0MemoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMem0MemoryResponse
func (client *Client) GetMem0MemoryWithContext(ctx context.Context, memoryId *string, request *GetMem0MemoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMem0MemoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentSpace) {
		query["agentSpace"] = request.AgentSpace
	}

	if !dara.IsNil(request.ContextStoreId) {
		query["context_store_id"] = request.ContextStoreId
	}

	if !dara.IsNil(request.OrgId) {
		query["org_id"] = request.OrgId
	}

	if !dara.IsNil(request.ProjectId) {
		query["project_id"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMem0Memory"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/memories/" + dara.PercentEncode(dara.StringValue(memoryId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMem0MemoryResponse{}
	_body, _err := client.DoROARequestWithCtx(ctx, params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询流水线
//
// @param request - GetPipelineRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPipelineResponse
func (client *Client) GetPipelineWithContext(ctx context.Context, agentSpace *string, pipelineName *string, request *GetPipelineRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPipelineResponse, _err error) {
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
		Action:      dara.String("GetPipeline"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/agentspace/" + dara.PercentEncode(dara.StringValue(agentSpace)) + "/pipeline/" + dara.PercentEncode(dara.StringValue(pipelineName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPipelineResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询AgentSpace列表
//
// @param request - ListAgentSpacesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAgentSpacesResponse
func (client *Client) ListAgentSpacesWithContext(ctx context.Context, request *ListAgentSpacesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAgentSpacesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentSpace) {
		query["agentSpace"] = request.AgentSpace
	}

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
		Action:      dara.String("ListAgentSpaces"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/agentspace"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAgentSpacesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取 API Key 列表
//
// @param request - ListContextStoreAPIKeysRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListContextStoreAPIKeysResponse
func (client *Client) ListContextStoreAPIKeysWithContext(ctx context.Context, agentSpace *string, contextStoreName *string, request *ListContextStoreAPIKeysRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListContextStoreAPIKeysResponse, _err error) {
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
		Action:      dara.String("ListContextStoreAPIKeys"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/agentspace/" + dara.PercentEncode(dara.StringValue(agentSpace)) + "/contextstore/" + dara.PercentEncode(dara.StringValue(contextStoreName)) + "/apikey"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListContextStoreAPIKeysResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询上下文库列表
//
// @param request - ListContextStoresRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListContextStoresResponse
func (client *Client) ListContextStoresWithContext(ctx context.Context, agentSpace *string, request *ListContextStoresRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListContextStoresResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ContextStoreName) {
		query["contextStoreName"] = request.ContextStoreName
	}

	if !dara.IsNil(request.ContextType) {
		query["contextType"] = request.ContextType
	}

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
		Action:      dara.String("ListContextStores"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/agentspace/" + dara.PercentEncode(dara.StringValue(agentSpace)) + "/contextstore"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListContextStoresResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询数据集列表
//
// @param request - ListDatasetsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDatasetsResponse
func (client *Client) ListDatasetsWithContext(ctx context.Context, agentSpace *string, request *ListDatasetsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDatasetsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DatasetName) {
		query["datasetName"] = request.DatasetName
	}

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
		Action:      dara.String("ListDatasets"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/agentspace/" + dara.PercentEncode(dara.StringValue(agentSpace)) + "/dataset"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListDatasetsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询流水线列表
//
// @param request - ListPipelinesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPipelinesResponse
func (client *Client) ListPipelinesWithContext(ctx context.Context, agentSpace *string, request *ListPipelinesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPipelinesResponse, _err error) {
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

	if !dara.IsNil(request.PipelineName) {
		query["pipelineName"] = request.PipelineName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPipelines"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/agentspace/" + dara.PercentEncode(dara.StringValue(agentSpace)) + "/pipeline"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPipelinesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 搜索上下文
//
// @param request - SearchContextRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchContextResponse
func (client *Client) SearchContextWithContext(ctx context.Context, agentSpace *string, contextStoreName *string, request *SearchContextRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SearchContextResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		body["filter"] = request.Filter
	}

	if !dara.IsNil(request.Formatted) {
		body["formatted"] = request.Formatted
	}

	if !dara.IsNil(request.Limit) {
		body["limit"] = request.Limit
	}

	if !dara.IsNil(request.Query) {
		body["query"] = request.Query
	}

	if !dara.IsNil(request.RetrievalOption) {
		body["retrievalOption"] = request.RetrievalOption
	}

	if !dara.IsNil(request.Threshold) {
		body["threshold"] = request.Threshold
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchContext"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/agentspace/" + dara.PercentEncode(dara.StringValue(agentSpace)) + "/contextstore/" + dara.PercentEncode(dara.StringValue(contextStoreName)) + "/context/search"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SearchContextResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询记忆库内容
//
// @param request - SearchMem0MemoriesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchMem0MemoriesResponse
func (client *Client) SearchMem0MemoriesWithContext(ctx context.Context, request *SearchMem0MemoriesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SearchMem0MemoriesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentSpace) {
		query["agentSpace"] = request.AgentSpace
	}

	if !dara.IsNil(request.ContextStoreId) {
		query["context_store_id"] = request.ContextStoreId
	}

	if !dara.IsNil(request.EnableGraph) {
		query["enable_graph"] = request.EnableGraph
	}

	if !dara.IsNil(request.OrgId) {
		query["org_id"] = request.OrgId
	}

	if !dara.IsNil(request.ProjectId) {
		query["project_id"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Body) {
		body["body"] = request.Body
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SearchMem0Memories"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v2/memories/search"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("array"),
	}
	_result = &SearchMem0MemoriesResponse{}
	_body, _err := client.DoROARequestWithCtx(ctx, params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新AgentSpace
//
// @param request - UpdateAgentSpaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAgentSpaceResponse
func (client *Client) UpdateAgentSpaceWithContext(ctx context.Context, agentSpace *string, request *UpdateAgentSpaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAgentSpaceResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.CmsWorkspace) {
		body["cmsWorkspace"] = request.CmsWorkspace
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAgentSpace"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/agentspace/" + dara.PercentEncode(dara.StringValue(agentSpace))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAgentSpaceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改上下文库配置
//
// @param request - UpdateContextStoreRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateContextStoreResponse
func (client *Client) UpdateContextStoreWithContext(ctx context.Context, agentSpace *string, contextStoreName *string, request *UpdateContextStoreRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateContextStoreResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.Config) {
		body["config"] = request.Config
	}

	if !dara.IsNil(request.ContextType) {
		body["contextType"] = request.ContextType
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateContextStore"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/agentspace/" + dara.PercentEncode(dara.StringValue(agentSpace)) + "/contextstore/" + dara.PercentEncode(dara.StringValue(contextStoreName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateContextStoreResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新数据集
//
// @param request - UpdateDatasetRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDatasetResponse
func (client *Client) UpdateDatasetWithContext(ctx context.Context, agentSpace *string, datasetName *string, request *UpdateDatasetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDatasetResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Schema) {
		body["schema"] = request.Schema
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDataset"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/agentspace/" + dara.PercentEncode(dara.StringValue(agentSpace)) + "/dataset/" + dara.PercentEncode(dara.StringValue(datasetName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDatasetResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改记忆
//
// @param request - UpdateMem0MemoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMem0MemoryResponse
func (client *Client) UpdateMem0MemoryWithContext(ctx context.Context, memoryId *string, request *UpdateMem0MemoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateMem0MemoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentSpace) {
		query["agentSpace"] = request.AgentSpace
	}

	if !dara.IsNil(request.ContextStoreId) {
		query["context_store_id"] = request.ContextStoreId
	}

	if !dara.IsNil(request.OrgId) {
		query["org_id"] = request.OrgId
	}

	if !dara.IsNil(request.ProjectId) {
		query["project_id"] = request.ProjectId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Body) {
		body["body"] = request.Body
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMem0Memory"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/memories/" + dara.PercentEncode(dara.StringValue(memoryId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMem0MemoryResponse{}
	_body, _err := client.DoROARequestWithCtx(ctx, params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新流水线
//
// @param request - UpdatePipelineRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdatePipelineResponse
func (client *Client) UpdatePipelineWithContext(ctx context.Context, agentSpace *string, pipelineName *string, request *UpdatePipelineRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePipelineResponse, _err error) {
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

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.ExecutePolicy) {
		body["executePolicy"] = request.ExecutePolicy
	}

	if !dara.IsNil(request.Pipeline) {
		body["pipeline"] = request.Pipeline
	}

	if !dara.IsNil(request.Sink) {
		body["sink"] = request.Sink
	}

	if !dara.IsNil(request.Source) {
		body["source"] = request.Source
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdatePipeline"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/agentspace/" + dara.PercentEncode(dara.StringValue(agentSpace)) + "/pipeline/" + dara.PercentEncode(dara.StringValue(pipelineName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdatePipelineResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 校验 Mem0 / ContextStore API Key
//
// @param request - ValidateMem0APIKeyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateMem0APIKeyResponse
func (client *Client) ValidateMem0APIKeyWithContext(ctx context.Context, request *ValidateMem0APIKeyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ValidateMem0APIKeyResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentSpace) {
		query["agentSpace"] = request.AgentSpace
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ValidateMem0APIKey"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/ping"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ValidateMem0APIKeyResponse{}
	_body, _err := client.DoROARequestWithCtx(ctx, params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
