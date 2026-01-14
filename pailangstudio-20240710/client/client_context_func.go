// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 创建知识库
//
// @param request - CreateKnowledgeBaseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateKnowledgeBaseResponse
func (client *Client) CreateKnowledgeBaseWithContext(ctx context.Context, request *CreateKnowledgeBaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateKnowledgeBaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Accessibility) {
		body["Accessibility"] = request.Accessibility
	}

	if !dara.IsNil(request.ChunkConfig) {
		body["ChunkConfig"] = request.ChunkConfig
	}

	if !dara.IsNil(request.DataSources) {
		body["DataSources"] = request.DataSources
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.EmbeddingConfig) {
		body["EmbeddingConfig"] = request.EmbeddingConfig
	}

	if !dara.IsNil(request.IndexColumnConfig) {
		body["IndexColumnConfig"] = request.IndexColumnConfig
	}

	if !dara.IsNil(request.KnowledgeBaseType) {
		body["KnowledgeBaseType"] = request.KnowledgeBaseType
	}

	if !dara.IsNil(request.MetaDataConfig) {
		body["MetaDataConfig"] = request.MetaDataConfig
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.OutputDir) {
		body["OutputDir"] = request.OutputDir
	}

	if !dara.IsNil(request.RuntimeId) {
		body["RuntimeId"] = request.RuntimeId
	}

	if !dara.IsNil(request.VectorDBConfig) {
		body["VectorDBConfig"] = request.VectorDBConfig
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateKnowledgeBase"),
		Version:     dara.String("2024-07-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/langstudio/knowledgebases"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateKnowledgeBaseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建知识库任务
//
// @param request - CreateKnowledgeBaseJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateKnowledgeBaseJobResponse
func (client *Client) CreateKnowledgeBaseJobWithContext(ctx context.Context, KnowledgeBaseId *string, request *CreateKnowledgeBaseJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateKnowledgeBaseJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Accessibility) {
		body["Accessibility"] = request.Accessibility
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.EcsSpecs) {
		body["EcsSpecs"] = request.EcsSpecs
	}

	if !dara.IsNil(request.EmbeddingConfig) {
		body["EmbeddingConfig"] = request.EmbeddingConfig
	}

	if !dara.IsNil(request.JobAction) {
		body["JobAction"] = request.JobAction
	}

	if !dara.IsNil(request.MaxRunningTimeInSeconds) {
		body["MaxRunningTimeInSeconds"] = request.MaxRunningTimeInSeconds
	}

	if !dara.IsNil(request.ResourceId) {
		body["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.UserVpc) {
		body["UserVpc"] = request.UserVpc
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateKnowledgeBaseJob"),
		Version:     dara.String("2024-07-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/langstudio/knowledgebases/" + dara.PercentEncode(dara.StringValue(KnowledgeBaseId)) + "/knowledgebasejobs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateKnowledgeBaseJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除知识库
//
// @param request - DeleteKnowledgeBaseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteKnowledgeBaseResponse
func (client *Client) DeleteKnowledgeBaseWithContext(ctx context.Context, KnowledgeBaseId *string, request *DeleteKnowledgeBaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteKnowledgeBaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteKnowledgeBase"),
		Version:     dara.String("2024-07-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/langstudio/knowledgebases/" + dara.PercentEncode(dara.StringValue(KnowledgeBaseId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteKnowledgeBaseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除知识库任务
//
// @param request - DeleteKnowledgeBaseJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteKnowledgeBaseJobResponse
func (client *Client) DeleteKnowledgeBaseJobWithContext(ctx context.Context, KnowledgeBaseId *string, KnowledgeBaseJobId *string, request *DeleteKnowledgeBaseJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteKnowledgeBaseJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteKnowledgeBaseJob"),
		Version:     dara.String("2024-07-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/langstudio/knowledgebases/" + dara.PercentEncode(dara.StringValue(KnowledgeBaseId)) + "/knowledgebasejobs/" + dara.PercentEncode(dara.StringValue(KnowledgeBaseJobId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteKnowledgeBaseJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看知识库
//
// @param request - GetKnowledgeBaseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetKnowledgeBaseResponse
func (client *Client) GetKnowledgeBaseWithContext(ctx context.Context, KnowledgeBaseId *string, request *GetKnowledgeBaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetKnowledgeBaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.VersionName) {
		query["VersionName"] = request.VersionName
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetKnowledgeBase"),
		Version:     dara.String("2024-07-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/langstudio/knowledgebases/" + dara.PercentEncode(dara.StringValue(KnowledgeBaseId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetKnowledgeBaseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看知识库任务
//
// @param request - GetKnowledgeBaseJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetKnowledgeBaseJobResponse
func (client *Client) GetKnowledgeBaseJobWithContext(ctx context.Context, KnowledgeBaseId *string, KnowledgeBaseJobId *string, request *GetKnowledgeBaseJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetKnowledgeBaseJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetKnowledgeBaseJob"),
		Version:     dara.String("2024-07-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/langstudio/knowledgebases/" + dara.PercentEncode(dara.StringValue(KnowledgeBaseId)) + "/knowledgebasejobs/" + dara.PercentEncode(dara.StringValue(KnowledgeBaseJobId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetKnowledgeBaseJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取知识库切片列表
//
// @param request - ListKnowledgeBaseChunksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListKnowledgeBaseChunksResponse
func (client *Client) ListKnowledgeBaseChunksWithContext(ctx context.Context, KnowledgeBaseId *string, request *ListKnowledgeBaseChunksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListKnowledgeBaseChunksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ChunkStatus) {
		query["ChunkStatus"] = request.ChunkStatus
	}

	if !dara.IsNil(request.MetaData) {
		query["MetaData"] = request.MetaData
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.VersionName) {
		query["VersionName"] = request.VersionName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListKnowledgeBaseChunks"),
		Version:     dara.String("2024-07-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/langstudio/knowledgebases/" + dara.PercentEncode(dara.StringValue(KnowledgeBaseId)) + "/knowledgebasechunks"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListKnowledgeBaseChunksResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取知识库任务列表
//
// @param request - ListKnowledgeBaseJobsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListKnowledgeBaseJobsResponse
func (client *Client) ListKnowledgeBaseJobsWithContext(ctx context.Context, KnowledgeBaseId *string, request *ListKnowledgeBaseJobsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListKnowledgeBaseJobsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.JobAction) {
		query["JobAction"] = request.JobAction
	}

	if !dara.IsNil(request.KnowledgeBaseJobId) {
		query["KnowledgeBaseJobId"] = request.KnowledgeBaseJobId
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Status) {
		query["Status"] = request.Status
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListKnowledgeBaseJobs"),
		Version:     dara.String("2024-07-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/langstudio/knowledgebases/" + dara.PercentEncode(dara.StringValue(KnowledgeBaseId)) + "/knowledgebasejobs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListKnowledgeBaseJobsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取知识库列表
//
// @param request - ListKnowledgeBasesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListKnowledgeBasesResponse
func (client *Client) ListKnowledgeBasesWithContext(ctx context.Context, request *ListKnowledgeBasesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListKnowledgeBasesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Creator) {
		query["Creator"] = request.Creator
	}

	if !dara.IsNil(request.KnowledgeBaseId) {
		query["KnowledgeBaseId"] = request.KnowledgeBaseId
	}

	if !dara.IsNil(request.KnowledgeBaseType) {
		query["KnowledgeBaseType"] = request.KnowledgeBaseType
	}

	if !dara.IsNil(request.MaxResults) {
		query["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.Name) {
		query["Name"] = request.Name
	}

	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Order) {
		query["Order"] = request.Order
	}

	if !dara.IsNil(request.PageNumber) {
		query["PageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["PageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SortBy) {
		query["SortBy"] = request.SortBy
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListKnowledgeBases"),
		Version:     dara.String("2024-07-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/langstudio/knowledgebases"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListKnowledgeBasesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 搜索知识库
//
// @param request - RetrieveKnowledgeBaseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetrieveKnowledgeBaseResponse
func (client *Client) RetrieveKnowledgeBaseWithContext(ctx context.Context, KnowledgeBaseId *string, request *RetrieveKnowledgeBaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RetrieveKnowledgeBaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.HybridStrategyConfig) {
		body["HybridStrategyConfig"] = request.HybridStrategyConfig
	}

	if !dara.IsNil(request.MetaDataFilterConditions) {
		body["MetaDataFilterConditions"] = request.MetaDataFilterConditions
	}

	if !dara.IsNil(request.Query) {
		body["Query"] = request.Query
	}

	if !dara.IsNil(request.QueryMode) {
		body["QueryMode"] = request.QueryMode
	}

	if !dara.IsNil(request.RerankConfig) {
		body["RerankConfig"] = request.RerankConfig
	}

	if !dara.IsNil(request.RewriteConfig) {
		body["RewriteConfig"] = request.RewriteConfig
	}

	if !dara.IsNil(request.ScoreThreshold) {
		body["ScoreThreshold"] = request.ScoreThreshold
	}

	if !dara.IsNil(request.TopK) {
		body["TopK"] = request.TopK
	}

	if !dara.IsNil(request.VersionName) {
		body["VersionName"] = request.VersionName
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RetrieveKnowledgeBase"),
		Version:     dara.String("2024-07-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/langstudio/knowledgebases/" + dara.PercentEncode(dara.StringValue(KnowledgeBaseId)) + "/action/retrieve"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RetrieveKnowledgeBaseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新知识库
//
// @param request - UpdateKnowledgeBaseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateKnowledgeBaseResponse
func (client *Client) UpdateKnowledgeBaseWithContext(ctx context.Context, KnowledgeBaseId *string, request *UpdateKnowledgeBaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateKnowledgeBaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AutoUpdateConfig) {
		body["AutoUpdateConfig"] = request.AutoUpdateConfig
	}

	if !dara.IsNil(request.BindRuntime) {
		body["BindRuntime"] = request.BindRuntime
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.MetaDataConfig) {
		body["MetaDataConfig"] = request.MetaDataConfig
	}

	if !dara.IsNil(request.RuntimeId) {
		body["RuntimeId"] = request.RuntimeId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateKnowledgeBase"),
		Version:     dara.String("2024-07-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/langstudio/knowledgebases/" + dara.PercentEncode(dara.StringValue(KnowledgeBaseId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateKnowledgeBaseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新知识库切片
//
// @param request - UpdateKnowledgeBaseChunkRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateKnowledgeBaseChunkResponse
func (client *Client) UpdateKnowledgeBaseChunkWithContext(ctx context.Context, KnowledgeBaseId *string, KnowledgeBaseChunkId *string, request *UpdateKnowledgeBaseChunkRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateKnowledgeBaseChunkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ChunkContent) {
		body["ChunkContent"] = request.ChunkContent
	}

	if !dara.IsNil(request.ChunkStatus) {
		body["ChunkStatus"] = request.ChunkStatus
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateKnowledgeBaseChunk"),
		Version:     dara.String("2024-07-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/langstudio/knowledgebases/" + dara.PercentEncode(dara.StringValue(KnowledgeBaseId)) + "/knowledgebasechunks/" + dara.PercentEncode(dara.StringValue(KnowledgeBaseChunkId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateKnowledgeBaseChunkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新知识库任务
//
// @param request - UpdateKnowledgeBaseJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateKnowledgeBaseJobResponse
func (client *Client) UpdateKnowledgeBaseJobWithContext(ctx context.Context, KnowledgeBaseId *string, KnowledgeBaseJobId *string, request *UpdateKnowledgeBaseJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateKnowledgeBaseJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateKnowledgeBaseJob"),
		Version:     dara.String("2024-07-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/langstudio/knowledgebases/" + dara.PercentEncode(dara.StringValue(KnowledgeBaseId)) + "/knowledgebasejobs/" + dara.PercentEncode(dara.StringValue(KnowledgeBaseJobId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateKnowledgeBaseJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
