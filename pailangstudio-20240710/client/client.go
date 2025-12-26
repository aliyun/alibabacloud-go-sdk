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
	client.Endpoint, _err = client.GetEndpoint(dara.String("pailangstudio"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 创建知识库
//
// @param request - CreateKnowledgeBaseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateKnowledgeBaseResponse
func (client *Client) CreateKnowledgeBaseWithOptions(request *CreateKnowledgeBaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateKnowledgeBaseResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建知识库
//
// @param request - CreateKnowledgeBaseRequest
//
// @return CreateKnowledgeBaseResponse
func (client *Client) CreateKnowledgeBase(request *CreateKnowledgeBaseRequest) (_result *CreateKnowledgeBaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateKnowledgeBaseResponse{}
	_body, _err := client.CreateKnowledgeBaseWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateKnowledgeBaseJobWithOptions(KnowledgeBaseId *string, request *CreateKnowledgeBaseJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateKnowledgeBaseJobResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateKnowledgeBaseJobResponse
func (client *Client) CreateKnowledgeBaseJob(KnowledgeBaseId *string, request *CreateKnowledgeBaseJobRequest) (_result *CreateKnowledgeBaseJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateKnowledgeBaseJobResponse{}
	_body, _err := client.CreateKnowledgeBaseJobWithOptions(KnowledgeBaseId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteKnowledgeBaseWithOptions(KnowledgeBaseId *string, request *DeleteKnowledgeBaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteKnowledgeBaseResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteKnowledgeBaseResponse
func (client *Client) DeleteKnowledgeBase(KnowledgeBaseId *string, request *DeleteKnowledgeBaseRequest) (_result *DeleteKnowledgeBaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteKnowledgeBaseResponse{}
	_body, _err := client.DeleteKnowledgeBaseWithOptions(KnowledgeBaseId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteKnowledgeBaseJobWithOptions(KnowledgeBaseId *string, KnowledgeBaseJobId *string, request *DeleteKnowledgeBaseJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteKnowledgeBaseJobResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteKnowledgeBaseJobResponse
func (client *Client) DeleteKnowledgeBaseJob(KnowledgeBaseId *string, KnowledgeBaseJobId *string, request *DeleteKnowledgeBaseJobRequest) (_result *DeleteKnowledgeBaseJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteKnowledgeBaseJobResponse{}
	_body, _err := client.DeleteKnowledgeBaseJobWithOptions(KnowledgeBaseId, KnowledgeBaseJobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetKnowledgeBaseWithOptions(KnowledgeBaseId *string, request *GetKnowledgeBaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetKnowledgeBaseResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetKnowledgeBaseResponse
func (client *Client) GetKnowledgeBase(KnowledgeBaseId *string, request *GetKnowledgeBaseRequest) (_result *GetKnowledgeBaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetKnowledgeBaseResponse{}
	_body, _err := client.GetKnowledgeBaseWithOptions(KnowledgeBaseId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetKnowledgeBaseJobWithOptions(KnowledgeBaseId *string, KnowledgeBaseJobId *string, request *GetKnowledgeBaseJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetKnowledgeBaseJobResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetKnowledgeBaseJobResponse
func (client *Client) GetKnowledgeBaseJob(KnowledgeBaseId *string, KnowledgeBaseJobId *string, request *GetKnowledgeBaseJobRequest) (_result *GetKnowledgeBaseJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetKnowledgeBaseJobResponse{}
	_body, _err := client.GetKnowledgeBaseJobWithOptions(KnowledgeBaseId, KnowledgeBaseJobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListKnowledgeBaseJobsWithOptions(KnowledgeBaseId *string, request *ListKnowledgeBaseJobsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListKnowledgeBaseJobsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListKnowledgeBaseJobsResponse
func (client *Client) ListKnowledgeBaseJobs(KnowledgeBaseId *string, request *ListKnowledgeBaseJobsRequest) (_result *ListKnowledgeBaseJobsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListKnowledgeBaseJobsResponse{}
	_body, _err := client.ListKnowledgeBaseJobsWithOptions(KnowledgeBaseId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListKnowledgeBasesWithOptions(request *ListKnowledgeBasesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListKnowledgeBasesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListKnowledgeBasesResponse
func (client *Client) ListKnowledgeBases(request *ListKnowledgeBasesRequest) (_result *ListKnowledgeBasesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListKnowledgeBasesResponse{}
	_body, _err := client.ListKnowledgeBasesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RetrieveKnowledgeBaseWithOptions(KnowledgeBaseId *string, request *RetrieveKnowledgeBaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RetrieveKnowledgeBaseResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RetrieveKnowledgeBaseResponse
func (client *Client) RetrieveKnowledgeBase(KnowledgeBaseId *string, request *RetrieveKnowledgeBaseRequest) (_result *RetrieveKnowledgeBaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RetrieveKnowledgeBaseResponse{}
	_body, _err := client.RetrieveKnowledgeBaseWithOptions(KnowledgeBaseId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateKnowledgeBaseWithOptions(KnowledgeBaseId *string, request *UpdateKnowledgeBaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateKnowledgeBaseResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateKnowledgeBaseResponse
func (client *Client) UpdateKnowledgeBase(KnowledgeBaseId *string, request *UpdateKnowledgeBaseRequest) (_result *UpdateKnowledgeBaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateKnowledgeBaseResponse{}
	_body, _err := client.UpdateKnowledgeBaseWithOptions(KnowledgeBaseId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateKnowledgeBaseJobWithOptions(KnowledgeBaseId *string, KnowledgeBaseJobId *string, request *UpdateKnowledgeBaseJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateKnowledgeBaseJobResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateKnowledgeBaseJobResponse
func (client *Client) UpdateKnowledgeBaseJob(KnowledgeBaseId *string, KnowledgeBaseJobId *string, request *UpdateKnowledgeBaseJobRequest) (_result *UpdateKnowledgeBaseJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateKnowledgeBaseJobResponse{}
	_body, _err := client.UpdateKnowledgeBaseJobWithOptions(KnowledgeBaseId, KnowledgeBaseJobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
