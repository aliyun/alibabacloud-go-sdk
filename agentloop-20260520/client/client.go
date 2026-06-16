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
	client.Endpoint, _err = client.GetEndpoint(dara.String("agentloop"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 给记忆库中增加数据
//
// @param request - AddMem0MemoriesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddMem0MemoriesResponse
func (client *Client) AddMem0MemoriesWithOptions(request *AddMem0MemoriesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddMem0MemoriesResponse, _err error) {
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
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 给记忆库中增加数据
//
// @param request - AddMem0MemoriesRequest
//
// @return AddMem0MemoriesResponse
func (client *Client) AddMem0Memories(request *AddMem0MemoriesRequest) (_result *AddMem0MemoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddMem0MemoriesResponse{}
	_body, _err := client.AddMem0MemoriesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateAgentSpaceWithOptions(request *CreateAgentSpaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAgentSpaceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateAgentSpaceResponse
func (client *Client) CreateAgentSpace(request *CreateAgentSpaceRequest) (_result *CreateAgentSpaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAgentSpaceResponse{}
	_body, _err := client.CreateAgentSpaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateContextStoreWithOptions(agentSpace *string, request *CreateContextStoreRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateContextStoreResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateContextStoreResponse
func (client *Client) CreateContextStore(agentSpace *string, request *CreateContextStoreRequest) (_result *CreateContextStoreResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateContextStoreResponse{}
	_body, _err := client.CreateContextStoreWithOptions(agentSpace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateContextStoreAPIKeyWithOptions(agentSpace *string, contextStoreName *string, request *CreateContextStoreAPIKeyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateContextStoreAPIKeyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateContextStoreAPIKeyResponse
func (client *Client) CreateContextStoreAPIKey(agentSpace *string, contextStoreName *string, request *CreateContextStoreAPIKeyRequest) (_result *CreateContextStoreAPIKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateContextStoreAPIKeyResponse{}
	_body, _err := client.CreateContextStoreAPIKeyWithOptions(agentSpace, contextStoreName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateDatasetWithOptions(agentSpace *string, request *CreateDatasetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDatasetResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateDatasetResponse
func (client *Client) CreateDataset(agentSpace *string, request *CreateDatasetRequest) (_result *CreateDatasetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDatasetResponse{}
	_body, _err := client.CreateDatasetWithOptions(agentSpace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteAgentSpaceWithOptions(agentSpace *string, request *DeleteAgentSpaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAgentSpaceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteAgentSpaceResponse
func (client *Client) DeleteAgentSpace(agentSpace *string, request *DeleteAgentSpaceRequest) (_result *DeleteAgentSpaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAgentSpaceResponse{}
	_body, _err := client.DeleteAgentSpaceWithOptions(agentSpace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteContextStoreWithOptions(agentSpace *string, contextStoreName *string, request *DeleteContextStoreRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteContextStoreResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteContextStoreResponse
func (client *Client) DeleteContextStore(agentSpace *string, contextStoreName *string, request *DeleteContextStoreRequest) (_result *DeleteContextStoreResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteContextStoreResponse{}
	_body, _err := client.DeleteContextStoreWithOptions(agentSpace, contextStoreName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteContextStoreAPIKeyWithOptions(agentSpace *string, contextStoreName *string, name *string, request *DeleteContextStoreAPIKeyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteContextStoreAPIKeyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteContextStoreAPIKeyResponse
func (client *Client) DeleteContextStoreAPIKey(agentSpace *string, contextStoreName *string, name *string, request *DeleteContextStoreAPIKeyRequest) (_result *DeleteContextStoreAPIKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteContextStoreAPIKeyResponse{}
	_body, _err := client.DeleteContextStoreAPIKeyWithOptions(agentSpace, contextStoreName, name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteDatasetWithOptions(agentSpace *string, datasetName *string, request *DeleteDatasetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDatasetResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteDatasetResponse
func (client *Client) DeleteDataset(agentSpace *string, datasetName *string, request *DeleteDatasetRequest) (_result *DeleteDatasetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDatasetResponse{}
	_body, _err := client.DeleteDatasetWithOptions(agentSpace, datasetName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteMem0MemoriesWithOptions(tmpReq *DeleteMem0MemoriesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteMem0MemoriesResponse, _err error) {
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
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
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
// @param request - DeleteMem0MemoriesRequest
//
// @return DeleteMem0MemoriesResponse
func (client *Client) DeleteMem0Memories(request *DeleteMem0MemoriesRequest) (_result *DeleteMem0MemoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteMem0MemoriesResponse{}
	_body, _err := client.DeleteMem0MemoriesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteMem0MemoryWithOptions(memoryId *string, request *DeleteMem0MemoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteMem0MemoryResponse, _err error) {
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
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
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
// @return DeleteMem0MemoryResponse
func (client *Client) DeleteMem0Memory(memoryId *string, request *DeleteMem0MemoryRequest) (_result *DeleteMem0MemoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteMem0MemoryResponse{}
	_body, _err := client.DeleteMem0MemoryWithOptions(memoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeletePipelineWithOptions(agentSpace *string, pipelineName *string, request *DeletePipelineRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeletePipelineResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeletePipelineResponse
func (client *Client) DeletePipeline(agentSpace *string, pipelineName *string, request *DeletePipelineRequest) (_result *DeletePipelineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeletePipelineResponse{}
	_body, _err := client.DeletePipelineWithOptions(agentSpace, pipelineName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeRegionsResponse
func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ExecuteQueryWithOptions(agentSpace *string, datasetName *string, request *ExecuteQueryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecuteQueryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ExecuteQueryResponse
func (client *Client) ExecuteQuery(agentSpace *string, datasetName *string, request *ExecuteQueryRequest) (_result *ExecuteQueryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExecuteQueryResponse{}
	_body, _err := client.ExecuteQueryWithOptions(agentSpace, datasetName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetAgentSpaceWithOptions(agentSpace *string, request *GetAgentSpaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAgentSpaceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetAgentSpaceResponse
func (client *Client) GetAgentSpace(agentSpace *string, request *GetAgentSpaceRequest) (_result *GetAgentSpaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAgentSpaceResponse{}
	_body, _err := client.GetAgentSpaceWithOptions(agentSpace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetContextStoreWithOptions(agentSpace *string, contextStoreName *string, request *GetContextStoreRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetContextStoreResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetContextStoreResponse
func (client *Client) GetContextStore(agentSpace *string, contextStoreName *string, request *GetContextStoreRequest) (_result *GetContextStoreResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetContextStoreResponse{}
	_body, _err := client.GetContextStoreWithOptions(agentSpace, contextStoreName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetContextStoreAPIKeyWithOptions(agentSpace *string, contextStoreName *string, name *string, request *GetContextStoreAPIKeyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetContextStoreAPIKeyResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetContextStoreAPIKeyResponse
func (client *Client) GetContextStoreAPIKey(agentSpace *string, contextStoreName *string, name *string, request *GetContextStoreAPIKeyRequest) (_result *GetContextStoreAPIKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetContextStoreAPIKeyResponse{}
	_body, _err := client.GetContextStoreAPIKeyWithOptions(agentSpace, contextStoreName, name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetDatasetWithOptions(agentSpace *string, datasetName *string, request *GetDatasetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDatasetResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetDatasetResponse
func (client *Client) GetDataset(agentSpace *string, datasetName *string, request *GetDatasetRequest) (_result *GetDatasetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDatasetResponse{}
	_body, _err := client.GetDatasetWithOptions(agentSpace, datasetName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetMem0MemoriesWithOptions(request *GetMem0MemoriesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMem0MemoriesResponse, _err error) {
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
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
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
// @return GetMem0MemoriesResponse
func (client *Client) GetMem0Memories(request *GetMem0MemoriesRequest) (_result *GetMem0MemoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMem0MemoriesResponse{}
	_body, _err := client.GetMem0MemoriesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetMem0MemoryWithOptions(memoryId *string, request *GetMem0MemoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMem0MemoryResponse, _err error) {
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
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
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
// @return GetMem0MemoryResponse
func (client *Client) GetMem0Memory(memoryId *string, request *GetMem0MemoryRequest) (_result *GetMem0MemoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMem0MemoryResponse{}
	_body, _err := client.GetMem0MemoryWithOptions(memoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetPipelineWithOptions(agentSpace *string, pipelineName *string, request *GetPipelineRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPipelineResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetPipelineResponse
func (client *Client) GetPipeline(agentSpace *string, pipelineName *string, request *GetPipelineRequest) (_result *GetPipelineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPipelineResponse{}
	_body, _err := client.GetPipelineWithOptions(agentSpace, pipelineName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListAgentSpacesWithOptions(request *ListAgentSpacesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAgentSpacesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListAgentSpacesResponse
func (client *Client) ListAgentSpaces(request *ListAgentSpacesRequest) (_result *ListAgentSpacesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAgentSpacesResponse{}
	_body, _err := client.ListAgentSpacesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListContextStoreAPIKeysWithOptions(agentSpace *string, contextStoreName *string, request *ListContextStoreAPIKeysRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListContextStoreAPIKeysResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListContextStoreAPIKeysResponse
func (client *Client) ListContextStoreAPIKeys(agentSpace *string, contextStoreName *string, request *ListContextStoreAPIKeysRequest) (_result *ListContextStoreAPIKeysResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListContextStoreAPIKeysResponse{}
	_body, _err := client.ListContextStoreAPIKeysWithOptions(agentSpace, contextStoreName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListContextStoresWithOptions(agentSpace *string, request *ListContextStoresRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListContextStoresResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListContextStoresResponse
func (client *Client) ListContextStores(agentSpace *string, request *ListContextStoresRequest) (_result *ListContextStoresResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListContextStoresResponse{}
	_body, _err := client.ListContextStoresWithOptions(agentSpace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListDatasetsWithOptions(agentSpace *string, request *ListDatasetsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListDatasetsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListDatasetsResponse
func (client *Client) ListDatasets(agentSpace *string, request *ListDatasetsRequest) (_result *ListDatasetsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDatasetsResponse{}
	_body, _err := client.ListDatasetsWithOptions(agentSpace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListPipelinesWithOptions(agentSpace *string, request *ListPipelinesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPipelinesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListPipelinesResponse
func (client *Client) ListPipelines(agentSpace *string, request *ListPipelinesRequest) (_result *ListPipelinesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPipelinesResponse{}
	_body, _err := client.ListPipelinesWithOptions(agentSpace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SearchContextWithOptions(agentSpace *string, contextStoreName *string, request *SearchContextRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SearchContextResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SearchContextResponse
func (client *Client) SearchContext(agentSpace *string, contextStoreName *string, request *SearchContextRequest) (_result *SearchContextResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SearchContextResponse{}
	_body, _err := client.SearchContextWithOptions(agentSpace, contextStoreName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SearchMem0MemoriesWithOptions(request *SearchMem0MemoriesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SearchMem0MemoriesResponse, _err error) {
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
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
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
// @return SearchMem0MemoriesResponse
func (client *Client) SearchMem0Memories(request *SearchMem0MemoriesRequest) (_result *SearchMem0MemoriesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SearchMem0MemoriesResponse{}
	_body, _err := client.SearchMem0MemoriesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateAgentSpaceWithOptions(agentSpace *string, request *UpdateAgentSpaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAgentSpaceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateAgentSpaceResponse
func (client *Client) UpdateAgentSpace(agentSpace *string, request *UpdateAgentSpaceRequest) (_result *UpdateAgentSpaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAgentSpaceResponse{}
	_body, _err := client.UpdateAgentSpaceWithOptions(agentSpace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateContextStoreWithOptions(agentSpace *string, contextStoreName *string, request *UpdateContextStoreRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateContextStoreResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateContextStoreResponse
func (client *Client) UpdateContextStore(agentSpace *string, contextStoreName *string, request *UpdateContextStoreRequest) (_result *UpdateContextStoreResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateContextStoreResponse{}
	_body, _err := client.UpdateContextStoreWithOptions(agentSpace, contextStoreName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateDatasetWithOptions(agentSpace *string, datasetName *string, request *UpdateDatasetRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDatasetResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateDatasetResponse
func (client *Client) UpdateDataset(agentSpace *string, datasetName *string, request *UpdateDatasetRequest) (_result *UpdateDatasetResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDatasetResponse{}
	_body, _err := client.UpdateDatasetWithOptions(agentSpace, datasetName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateMem0MemoryWithOptions(memoryId *string, request *UpdateMem0MemoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateMem0MemoryResponse, _err error) {
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
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
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
// @return UpdateMem0MemoryResponse
func (client *Client) UpdateMem0Memory(memoryId *string, request *UpdateMem0MemoryRequest) (_result *UpdateMem0MemoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateMem0MemoryResponse{}
	_body, _err := client.UpdateMem0MemoryWithOptions(memoryId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdatePipelineWithOptions(agentSpace *string, pipelineName *string, request *UpdatePipelineRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdatePipelineResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdatePipelineResponse
func (client *Client) UpdatePipeline(agentSpace *string, pipelineName *string, request *UpdatePipelineRequest) (_result *UpdatePipelineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdatePipelineResponse{}
	_body, _err := client.UpdatePipelineWithOptions(agentSpace, pipelineName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ValidateMem0APIKeyWithOptions(request *ValidateMem0APIKeyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ValidateMem0APIKeyResponse, _err error) {
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
	_body, _err := client.DoROARequest(params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.Pathname, params.BodyType, req, runtime)
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
// @return ValidateMem0APIKeyResponse
func (client *Client) ValidateMem0APIKey(request *ValidateMem0APIKeyRequest) (_result *ValidateMem0APIKeyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ValidateMem0APIKeyResponse{}
	_body, _err := client.ValidateMem0APIKeyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
