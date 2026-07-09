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
		"cn-zhangjiakou": dara.String("agentloop.cn-zhangjiakou.aliyuncs.com"),
		"cn-shenzhen":    dara.String("agentloop.cn-shenzhen.aliyuncs.com"),
		"cn-shanghai":    dara.String("agentloop.cn-shanghai.aliyuncs.com"),
		"cn-hongkong":    dara.String("agentloop.cn-hongkong.aliyuncs.com"),
		"cn-hangzhou":    dara.String("agentloop.cn-hangzhou.aliyuncs.com"),
		"cn-guangzhou":   dara.String("agentloop.cn-guangzhou.aliyuncs.com"),
		"cn-chengdu":     dara.String("agentloop.cn-chengdu.aliyuncs.com"),
		"cn-beijing":     dara.String("agentloop.cn-beijing.aliyuncs.com"),
		"ap-southeast-1": dara.String("agentloop.ap-southeast-1.aliyuncs.com"),
	}
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
// Appends structured data rows to a specified dataset without requiring the client to construct SQL statements.
//
// @param request - AddDatasetDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddDatasetDataResponse
func (client *Client) AddDatasetDataWithOptions(agentSpace *string, datasetName *string, request *AddDatasetDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddDatasetDataResponse, _err error) {
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
	if !dara.IsNil(request.DataArray) {
		body["dataArray"] = request.DataArray
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddDatasetData"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/agentspace/" + dara.PercentEncode(dara.StringValue(agentSpace)) + "/dataset/" + dara.PercentEncode(dara.StringValue(datasetName)) + "/rows"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AddDatasetDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Appends structured data rows to a specified dataset without requiring the client to construct SQL statements.
//
// @param request - AddDatasetDataRequest
//
// @return AddDatasetDataResponse
func (client *Client) AddDatasetData(agentSpace *string, datasetName *string, request *AddDatasetDataRequest) (_result *AddDatasetDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddDatasetDataResponse{}
	_body, _err := client.AddDatasetDataWithOptions(agentSpace, datasetName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels a pipeline run.
//
// @param request - CancelPipelineRunRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelPipelineRunResponse
func (client *Client) CancelPipelineRunWithOptions(agentSpace *string, pipelineName *string, runId *string, request *CancelPipelineRunRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CancelPipelineRunResponse, _err error) {
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
		Action:      dara.String("CancelPipelineRun"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/agentspace/" + dara.PercentEncode(dara.StringValue(agentSpace)) + "/pipeline/" + dara.PercentEncode(dara.StringValue(pipelineName)) + "/runs/" + dara.PercentEncode(dara.StringValue(runId)) + "/cancel"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CancelPipelineRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels a pipeline run.
//
// @param request - CancelPipelineRunRequest
//
// @return CancelPipelineRunResponse
func (client *Client) CancelPipelineRun(agentSpace *string, pipelineName *string, runId *string, request *CancelPipelineRunRequest) (_result *CancelPipelineRunResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelPipelineRunResponse{}
	_body, _err := client.CancelPipelineRunWithOptions(agentSpace, pipelineName, runId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an AgentSpace.
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

	if !dara.IsNil(request.TrajectoryStoreEnabled) {
		body["trajectoryStoreEnabled"] = request.TrajectoryStoreEnabled
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
// Creates an AgentSpace.
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
// Creates a context store.
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
// Creates a context store.
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
// Creates an API key.
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
// Creates an API key.
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
// Creates a dataset.
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
// Creates a dataset.
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
// Creates an evaluation task.
//
// Description:
//
// Calls the CreateEvaluationTask operation to create an evaluation task in a specified AgentSpace. The server verifies AgentSpace permissions, initializes evaluation result storage, checks the uniqueness of the task name, and asynchronously creates and executes an EvaluationRun based on `taskMode` and `runStrategies`.
//
// This operation is applicable to running built-in or custom evaluators on Trace, Dataset, or SLS Log data. It supports two execution strategies: historical backfill and continuous evaluation.
//
// @param request - CreateEvaluationTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEvaluationTaskResponse
func (client *Client) CreateEvaluationTaskWithOptions(agentSpace *string, request *CreateEvaluationTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateEvaluationTaskResponse, _err error) {
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
	if !dara.IsNil(request.Channel) {
		body["channel"] = request.Channel
	}

	if !dara.IsNil(request.Config) {
		body["config"] = request.Config
	}

	if !dara.IsNil(request.DataFilter) {
		body["dataFilter"] = request.DataFilter
	}

	if !dara.IsNil(request.DataType) {
		body["dataType"] = request.DataType
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Evaluators) {
		body["evaluators"] = request.Evaluators
	}

	if !dara.IsNil(request.RunStrategies) {
		body["runStrategies"] = request.RunStrategies
	}

	if !dara.IsNil(request.Tags) {
		body["tags"] = request.Tags
	}

	if !dara.IsNil(request.TaskMode) {
		body["taskMode"] = request.TaskMode
	}

	if !dara.IsNil(request.TaskName) {
		body["taskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEvaluationTask"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/evaluation-task/" + dara.PercentEncode(dara.StringValue(agentSpace))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEvaluationTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an evaluation task.
//
// Description:
//
// Calls the CreateEvaluationTask operation to create an evaluation task in a specified AgentSpace. The server verifies AgentSpace permissions, initializes evaluation result storage, checks the uniqueness of the task name, and asynchronously creates and executes an EvaluationRun based on `taskMode` and `runStrategies`.
//
// This operation is applicable to running built-in or custom evaluators on Trace, Dataset, or SLS Log data. It supports two execution strategies: historical backfill and continuous evaluation.
//
// @param request - CreateEvaluationTaskRequest
//
// @return CreateEvaluationTaskResponse
func (client *Client) CreateEvaluationTask(agentSpace *string, request *CreateEvaluationTaskRequest) (_result *CreateEvaluationTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateEvaluationTaskResponse{}
	_body, _err := client.CreateEvaluationTaskWithOptions(agentSpace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an evaluator.
//
// @param request - CreateEvaluatorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEvaluatorResponse
func (client *Client) CreateEvaluatorWithOptions(agentSpace *string, request *CreateEvaluatorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateEvaluatorResponse, _err error) {
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
	if !dara.IsNil(request.Annotations) {
		body["annotations"] = request.Annotations
	}

	if !dara.IsNil(request.Config) {
		body["config"] = request.Config
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DisplayName) {
		body["displayName"] = request.DisplayName
	}

	if !dara.IsNil(request.MetricName) {
		body["metricName"] = request.MetricName
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Properties) {
		body["properties"] = request.Properties
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	if !dara.IsNil(request.Version) {
		body["version"] = request.Version
	}

	if !dara.IsNil(request.VersionDescription) {
		body["versionDescription"] = request.VersionDescription
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEvaluator"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/evaluators/" + dara.PercentEncode(dara.StringValue(agentSpace))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEvaluatorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an evaluator.
//
// @param request - CreateEvaluatorRequest
//
// @return CreateEvaluatorResponse
func (client *Client) CreateEvaluator(agentSpace *string, request *CreateEvaluatorRequest) (_result *CreateEvaluatorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateEvaluatorResponse{}
	_body, _err := client.CreateEvaluatorWithOptions(agentSpace, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an evaluator skill.
//
// @param request - CreateEvaluatorSkillRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEvaluatorSkillResponse
func (client *Client) CreateEvaluatorSkillWithOptions(name *string, request *CreateEvaluatorSkillRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateEvaluatorSkillResponse, _err error) {
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

	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DisplayName) {
		body["displayName"] = request.DisplayName
	}

	if !dara.IsNil(request.Enable) {
		body["enable"] = request.Enable
	}

	if !dara.IsNil(request.Files) {
		body["files"] = request.Files
	}

	if !dara.IsNil(request.SkillName) {
		body["skillName"] = request.SkillName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEvaluatorSkill"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/evaluator/" + dara.PercentEncode(dara.StringValue(name)) + "/skill"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEvaluatorSkillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an evaluator skill.
//
// @param request - CreateEvaluatorSkillRequest
//
// @return CreateEvaluatorSkillResponse
func (client *Client) CreateEvaluatorSkill(name *string, request *CreateEvaluatorSkillRequest) (_result *CreateEvaluatorSkillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateEvaluatorSkillResponse{}
	_body, _err := client.CreateEvaluatorSkillWithOptions(name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an AgentSpace.
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
// Deletes an AgentSpace.
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
// Deletes a context store.
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
// Deletes a context store.
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
// Deletes an API key.
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
// Deletes an API key.
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
// Deletes a dataset.
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
// Deletes a dataset.
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
// Deletes an evaluation run.
//
// @param request - DeleteEvaluationRunRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEvaluationRunResponse
func (client *Client) DeleteEvaluationRunWithOptions(agentSpace *string, taskId *string, runId *string, request *DeleteEvaluationRunRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteEvaluationRunResponse, _err error) {
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
		Action:      dara.String("DeleteEvaluationRun"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/evaluation-task/" + dara.PercentEncode(dara.StringValue(agentSpace)) + "/" + dara.PercentEncode(dara.StringValue(taskId)) + "/run/" + dara.PercentEncode(dara.StringValue(runId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEvaluationRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an evaluation run.
//
// @param request - DeleteEvaluationRunRequest
//
// @return DeleteEvaluationRunResponse
func (client *Client) DeleteEvaluationRun(agentSpace *string, taskId *string, runId *string, request *DeleteEvaluationRunRequest) (_result *DeleteEvaluationRunResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteEvaluationRunResponse{}
	_body, _err := client.DeleteEvaluationRunWithOptions(agentSpace, taskId, runId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an evaluation task.
//
// @param request - DeleteEvaluationTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEvaluationTaskResponse
func (client *Client) DeleteEvaluationTaskWithOptions(agentSpace *string, taskId *string, request *DeleteEvaluationTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteEvaluationTaskResponse, _err error) {
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
		Action:      dara.String("DeleteEvaluationTask"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/evaluation-task/" + dara.PercentEncode(dara.StringValue(agentSpace)) + "/" + dara.PercentEncode(dara.StringValue(taskId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEvaluationTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an evaluation task.
//
// @param request - DeleteEvaluationTaskRequest
//
// @return DeleteEvaluationTaskResponse
func (client *Client) DeleteEvaluationTask(agentSpace *string, taskId *string, request *DeleteEvaluationTaskRequest) (_result *DeleteEvaluationTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteEvaluationTaskResponse{}
	_body, _err := client.DeleteEvaluationTaskWithOptions(agentSpace, taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an evaluator.
//
// @param request - DeleteEvaluatorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEvaluatorResponse
func (client *Client) DeleteEvaluatorWithOptions(agentSpace *string, name *string, request *DeleteEvaluatorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteEvaluatorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteEvaluator"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/evaluators/" + dara.PercentEncode(dara.StringValue(agentSpace)) + "/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEvaluatorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an evaluator.
//
// @param request - DeleteEvaluatorRequest
//
// @return DeleteEvaluatorResponse
func (client *Client) DeleteEvaluator(agentSpace *string, name *string, request *DeleteEvaluatorRequest) (_result *DeleteEvaluatorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteEvaluatorResponse{}
	_body, _err := client.DeleteEvaluatorWithOptions(agentSpace, name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an evaluator skill.
//
// @param request - DeleteEvaluatorSkillRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEvaluatorSkillResponse
func (client *Client) DeleteEvaluatorSkillWithOptions(name *string, skillName *string, request *DeleteEvaluatorSkillRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteEvaluatorSkillResponse, _err error) {
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
		Action:      dara.String("DeleteEvaluatorSkill"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/evaluator/" + dara.PercentEncode(dara.StringValue(name)) + "/skill/" + dara.PercentEncode(dara.StringValue(skillName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteEvaluatorSkillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an evaluator skill.
//
// @param request - DeleteEvaluatorSkillRequest
//
// @return DeleteEvaluatorSkillResponse
func (client *Client) DeleteEvaluatorSkill(name *string, skillName *string, request *DeleteEvaluatorSkillRequest) (_result *DeleteEvaluatorSkillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteEvaluatorSkillResponse{}
	_body, _err := client.DeleteEvaluatorSkillWithOptions(name, skillName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a pipeline.
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
// Deletes a pipeline.
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
// Queries regions.
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
// Queries regions.
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
// Executes a query statement.
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
	if !dara.IsNil(request.From) {
		body["from"] = request.From
	}

	if !dara.IsNil(request.Length) {
		body["length"] = request.Length
	}

	if !dara.IsNil(request.MaxOutputLength) {
		body["maxOutputLength"] = request.MaxOutputLength
	}

	if !dara.IsNil(request.Offset) {
		body["offset"] = request.Offset
	}

	if !dara.IsNil(request.Query) {
		body["query"] = request.Query
	}

	if !dara.IsNil(request.To) {
		body["to"] = request.To
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
// Executes a query statement.
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
// Queries an AgentSpace.
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
// Queries an AgentSpace.
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
// Queries a context store.
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
// Queries a context store.
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
// Retrieves an API key.
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
// Retrieves an API key.
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
// Queries a dataset.
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
// Queries a dataset.
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
// Retrieves the details of an evaluation run.
//
// @param request - GetEvaluationRunRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEvaluationRunResponse
func (client *Client) GetEvaluationRunWithOptions(agentSpace *string, taskId *string, runId *string, request *GetEvaluationRunRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEvaluationRunResponse, _err error) {
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
		Action:      dara.String("GetEvaluationRun"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/evaluation-task/" + dara.PercentEncode(dara.StringValue(agentSpace)) + "/" + dara.PercentEncode(dara.StringValue(taskId)) + "/run/" + dara.PercentEncode(dara.StringValue(runId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEvaluationRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of an evaluation run.
//
// @param request - GetEvaluationRunRequest
//
// @return GetEvaluationRunResponse
func (client *Client) GetEvaluationRun(agentSpace *string, taskId *string, runId *string, request *GetEvaluationRunRequest) (_result *GetEvaluationRunResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEvaluationRunResponse{}
	_body, _err := client.GetEvaluationRunWithOptions(agentSpace, taskId, runId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of an evaluation task.
//
// @param request - GetEvaluationTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEvaluationTaskResponse
func (client *Client) GetEvaluationTaskWithOptions(agentSpace *string, taskId *string, request *GetEvaluationTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEvaluationTaskResponse, _err error) {
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
		Action:      dara.String("GetEvaluationTask"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/evaluation-task/" + dara.PercentEncode(dara.StringValue(agentSpace)) + "/" + dara.PercentEncode(dara.StringValue(taskId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEvaluationTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of an evaluation task.
//
// @param request - GetEvaluationTaskRequest
//
// @return GetEvaluationTaskResponse
func (client *Client) GetEvaluationTask(agentSpace *string, taskId *string, request *GetEvaluationTaskRequest) (_result *GetEvaluationTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEvaluationTaskResponse{}
	_body, _err := client.GetEvaluationTaskWithOptions(agentSpace, taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of an evaluator.
//
// @param request - GetEvaluatorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEvaluatorResponse
func (client *Client) GetEvaluatorWithOptions(agentSpace *string, name *string, request *GetEvaluatorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEvaluatorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Version) {
		query["version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEvaluator"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/evaluators/" + dara.PercentEncode(dara.StringValue(agentSpace)) + "/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEvaluatorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of an evaluator.
//
// @param request - GetEvaluatorRequest
//
// @return GetEvaluatorResponse
func (client *Client) GetEvaluator(agentSpace *string, name *string, request *GetEvaluatorRequest) (_result *GetEvaluatorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEvaluatorResponse{}
	_body, _err := client.GetEvaluatorWithOptions(agentSpace, name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of an evaluator skill.
//
// @param request - GetEvaluatorSkillRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEvaluatorSkillResponse
func (client *Client) GetEvaluatorSkillWithOptions(name *string, skillName *string, request *GetEvaluatorSkillRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEvaluatorSkillResponse, _err error) {
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

	if !dara.IsNil(request.Version) {
		query["version"] = request.Version
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEvaluatorSkill"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/evaluator/" + dara.PercentEncode(dara.StringValue(name)) + "/skill/" + dara.PercentEncode(dara.StringValue(skillName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEvaluatorSkillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of an evaluator skill.
//
// @param request - GetEvaluatorSkillRequest
//
// @return GetEvaluatorSkillResponse
func (client *Client) GetEvaluatorSkill(name *string, skillName *string, request *GetEvaluatorSkillRequest) (_result *GetEvaluatorSkillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEvaluatorSkillResponse{}
	_body, _err := client.GetEvaluatorSkillWithOptions(name, skillName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a CI/CD pipeline.
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
// Queries a CI/CD pipeline.
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
// Queries the details of a single pipeline run.
//
// @param request - GetPipelineRunRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPipelineRunResponse
func (client *Client) GetPipelineRunWithOptions(agentSpace *string, pipelineName *string, runId *string, request *GetPipelineRunRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPipelineRunResponse, _err error) {
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
		Action:      dara.String("GetPipelineRun"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/agentspace/" + dara.PercentEncode(dara.StringValue(agentSpace)) + "/pipeline/" + dara.PercentEncode(dara.StringValue(pipelineName)) + "/runs/" + dara.PercentEncode(dara.StringValue(runId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPipelineRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a single pipeline run.
//
// @param request - GetPipelineRunRequest
//
// @return GetPipelineRunResponse
func (client *Client) GetPipelineRun(agentSpace *string, pipelineName *string, runId *string, request *GetPipelineRunRequest) (_result *GetPipelineRunResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPipelineRunResponse{}
	_body, _err := client.GetPipelineRunWithOptions(agentSpace, pipelineName, runId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries pipeline run statistics.
//
// @param request - GetPipelineStatsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPipelineStatsResponse
func (client *Client) GetPipelineStatsWithOptions(agentSpace *string, pipelineName *string, request *GetPipelineStatsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPipelineStatsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.Granularity) {
		query["granularity"] = request.Granularity
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPipelineStats"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/agentspace/" + dara.PercentEncode(dara.StringValue(agentSpace)) + "/pipeline/" + dara.PercentEncode(dara.StringValue(pipelineName)) + "/stats"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPipelineStatsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries pipeline run statistics.
//
// @param request - GetPipelineStatsRequest
//
// @return GetPipelineStatsResponse
func (client *Client) GetPipelineStats(agentSpace *string, pipelineName *string, request *GetPipelineStatsRequest) (_result *GetPipelineStatsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPipelineStatsResponse{}
	_body, _err := client.GetPipelineStatsWithOptions(agentSpace, pipelineName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of AgentSpaces.
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

	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
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
// Queries a list of AgentSpaces.
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
// Retrieves a list of API keys.
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
// Retrieves a list of API keys.
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
// Queries a list of context stores.
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
// Queries a list of context stores.
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
// Queries a list of datasets.
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
// Queries a list of datasets.
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
// Queries the list of evaluation runs.
//
// @param request - ListEvaluationRunsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEvaluationRunsResponse
func (client *Client) ListEvaluationRunsWithOptions(agentSpace *string, taskId *string, request *ListEvaluationRunsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListEvaluationRunsResponse, _err error) {
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

	if !dara.IsNil(request.RunType) {
		query["runType"] = request.RunType
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEvaluationRuns"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/evaluation-task/" + dara.PercentEncode(dara.StringValue(agentSpace)) + "/" + dara.PercentEncode(dara.StringValue(taskId)) + "/runs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEvaluationRunsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of evaluation runs.
//
// @param request - ListEvaluationRunsRequest
//
// @return ListEvaluationRunsResponse
func (client *Client) ListEvaluationRuns(agentSpace *string, taskId *string, request *ListEvaluationRunsRequest) (_result *ListEvaluationRunsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEvaluationRunsResponse{}
	_body, _err := client.ListEvaluationRunsWithOptions(agentSpace, taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of evaluation tasks.
//
// @param request - ListEvaluationTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEvaluationTasksResponse
func (client *Client) ListEvaluationTasksWithOptions(request *ListEvaluationTasksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListEvaluationTasksResponse, _err error) {
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

	if !dara.IsNil(request.Channel) {
		query["channel"] = request.Channel
	}

	if !dara.IsNil(request.DataType) {
		query["dataType"] = request.DataType
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	if !dara.IsNil(request.TaskMode) {
		query["taskMode"] = request.TaskMode
	}

	if !dara.IsNil(request.TaskName) {
		query["taskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEvaluationTasks"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/evaluation-tasks"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEvaluationTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of evaluation tasks.
//
// @param request - ListEvaluationTasksRequest
//
// @return ListEvaluationTasksResponse
func (client *Client) ListEvaluationTasks(request *ListEvaluationTasksRequest) (_result *ListEvaluationTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEvaluationTasksResponse{}
	_body, _err := client.ListEvaluationTasksWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the skill list of an evaluator.
//
// @param request - ListEvaluatorSkillsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEvaluatorSkillsResponse
func (client *Client) ListEvaluatorSkillsWithOptions(name *string, request *ListEvaluatorSkillsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListEvaluatorSkillsResponse, _err error) {
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
		Action:      dara.String("ListEvaluatorSkills"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/evaluator/" + dara.PercentEncode(dara.StringValue(name)) + "/skills"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEvaluatorSkillsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the skill list of an evaluator.
//
// @param request - ListEvaluatorSkillsRequest
//
// @return ListEvaluatorSkillsResponse
func (client *Client) ListEvaluatorSkills(name *string, request *ListEvaluatorSkillsRequest) (_result *ListEvaluatorSkillsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEvaluatorSkillsResponse{}
	_body, _err := client.ListEvaluatorSkillsWithOptions(name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of evaluators.
//
// @param request - ListEvaluatorsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListEvaluatorsResponse
func (client *Client) ListEvaluatorsWithOptions(request *ListEvaluatorsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListEvaluatorsResponse, _err error) {
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

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.Source) {
		query["source"] = request.Source
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListEvaluators"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/evaluators"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListEvaluatorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of evaluators.
//
// @param request - ListEvaluatorsRequest
//
// @return ListEvaluatorsResponse
func (client *Client) ListEvaluators(request *ListEvaluatorsRequest) (_result *ListEvaluatorsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEvaluatorsResponse{}
	_body, _err := client.ListEvaluatorsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the execution history list of a pipeline.
//
// @param request - ListPipelineRunsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListPipelineRunsResponse
func (client *Client) ListPipelineRunsWithOptions(agentSpace *string, pipelineName *string, request *ListPipelineRunsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListPipelineRunsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.MaxResults) {
		query["maxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		query["nextToken"] = request.NextToken
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	if !dara.IsNil(request.TriggerType) {
		query["triggerType"] = request.TriggerType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListPipelineRuns"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/agentspace/" + dara.PercentEncode(dara.StringValue(agentSpace)) + "/pipeline/" + dara.PercentEncode(dara.StringValue(pipelineName)) + "/runs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListPipelineRunsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution history list of a pipeline.
//
// @param request - ListPipelineRunsRequest
//
// @return ListPipelineRunsResponse
func (client *Client) ListPipelineRuns(agentSpace *string, pipelineName *string, request *ListPipelineRunsRequest) (_result *ListPipelineRunsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPipelineRunsResponse{}
	_body, _err := client.ListPipelineRunsWithOptions(agentSpace, pipelineName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists CI/CD pipelines.
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

	if !dara.IsNil(request.ScheduleStatus) {
		query["scheduleStatus"] = request.ScheduleStatus
	}

	if !dara.IsNil(request.ScheduleType) {
		query["scheduleType"] = request.ScheduleType
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
// Lists CI/CD pipelines.
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
// Pauses a pipeline.
//
// @param request - PausePipelineRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PausePipelineResponse
func (client *Client) PausePipelineWithOptions(agentSpace *string, pipelineName *string, request *PausePipelineRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PausePipelineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Reason) {
		body["reason"] = request.Reason
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PausePipeline"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/agentspace/" + dara.PercentEncode(dara.StringValue(agentSpace)) + "/pipeline/" + dara.PercentEncode(dara.StringValue(pipelineName)) + "/pause"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PausePipelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Pauses a pipeline.
//
// @param request - PausePipelineRequest
//
// @return PausePipelineResponse
func (client *Client) PausePipeline(agentSpace *string, pipelineName *string, request *PausePipelineRequest) (_result *PausePipelineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PausePipelineResponse{}
	_body, _err := client.PausePipelineWithOptions(agentSpace, pipelineName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Resumes a pipeline.
//
// @param request - ResumePipelineRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumePipelineResponse
func (client *Client) ResumePipelineWithOptions(agentSpace *string, pipelineName *string, request *ResumePipelineRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ResumePipelineResponse, _err error) {
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
		Action:      dara.String("ResumePipeline"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/agentspace/" + dara.PercentEncode(dara.StringValue(agentSpace)) + "/pipeline/" + dara.PercentEncode(dara.StringValue(pipelineName)) + "/resume"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ResumePipelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resumes a pipeline.
//
// @param request - ResumePipelineRequest
//
// @return ResumePipelineResponse
func (client *Client) ResumePipeline(agentSpace *string, pipelineName *string, request *ResumePipelineRequest) (_result *ResumePipelineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ResumePipelineResponse{}
	_body, _err := client.ResumePipelineWithOptions(agentSpace, pipelineName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Manually triggers a pipeline execution.
//
// @param request - RunPipelineRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunPipelineResponse
func (client *Client) RunPipelineWithOptions(agentSpace *string, pipelineName *string, request *RunPipelineRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunPipelineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FromTime) {
		body["fromTime"] = request.FromTime
	}

	if !dara.IsNil(request.Output) {
		body["output"] = request.Output
	}

	if !dara.IsNil(request.ToTime) {
		body["toTime"] = request.ToTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunPipeline"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/agentspace/" + dara.PercentEncode(dara.StringValue(agentSpace)) + "/pipeline/" + dara.PercentEncode(dara.StringValue(pipelineName)) + "/run"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RunPipelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Manually triggers a pipeline execution.
//
// @param request - RunPipelineRequest
//
// @return RunPipelineResponse
func (client *Client) RunPipeline(agentSpace *string, pipelineName *string, request *RunPipelineRequest) (_result *RunPipelineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunPipelineResponse{}
	_body, _err := client.RunPipelineWithOptions(agentSpace, pipelineName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Searches contexts.
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
// Searches contexts.
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
// Stops a pipeline.
//
// @param request - TerminatePipelineRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TerminatePipelineResponse
func (client *Client) TerminatePipelineWithOptions(agentSpace *string, pipelineName *string, request *TerminatePipelineRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TerminatePipelineResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Reason) {
		body["reason"] = request.Reason
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TerminatePipeline"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/agentspace/" + dara.PercentEncode(dara.StringValue(agentSpace)) + "/pipeline/" + dara.PercentEncode(dara.StringValue(pipelineName)) + "/terminate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TerminatePipelineResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a pipeline.
//
// @param request - TerminatePipelineRequest
//
// @return TerminatePipelineResponse
func (client *Client) TerminatePipeline(agentSpace *string, pipelineName *string, request *TerminatePipelineRequest) (_result *TerminatePipelineResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TerminatePipelineResponse{}
	_body, _err := client.TerminatePipelineWithOptions(agentSpace, pipelineName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an AgentSpace.
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
// Updates an AgentSpace.
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
// Modifies the configuration of a context store.
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
// Modifies the configuration of a context store.
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
// Updates a dataset.
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
// Updates a dataset.
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
// Updates an evaluation run.
//
// @param request - UpdateEvaluationRunRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEvaluationRunResponse
func (client *Client) UpdateEvaluationRunWithOptions(agentSpace *string, taskId *string, runId *string, request *UpdateEvaluationRunRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateEvaluationRunResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEvaluationRun"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/evaluation-task/" + dara.PercentEncode(dara.StringValue(agentSpace)) + "/" + dara.PercentEncode(dara.StringValue(taskId)) + "/run/" + dara.PercentEncode(dara.StringValue(runId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEvaluationRunResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an evaluation run.
//
// @param request - UpdateEvaluationRunRequest
//
// @return UpdateEvaluationRunResponse
func (client *Client) UpdateEvaluationRun(agentSpace *string, taskId *string, runId *string, request *UpdateEvaluationRunRequest) (_result *UpdateEvaluationRunResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateEvaluationRunResponse{}
	_body, _err := client.UpdateEvaluationRunWithOptions(agentSpace, taskId, runId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an evaluation task.
//
// @param request - UpdateEvaluationTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEvaluationTaskResponse
func (client *Client) UpdateEvaluationTaskWithOptions(agentSpace *string, taskId *string, request *UpdateEvaluationTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateEvaluationTaskResponse, _err error) {
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

	if !dara.IsNil(request.DataFilter) {
		body["dataFilter"] = request.DataFilter
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Evaluators) {
		body["evaluators"] = request.Evaluators
	}

	if !dara.IsNil(request.RunStrategies) {
		body["runStrategies"] = request.RunStrategies
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	if !dara.IsNil(request.Tags) {
		body["tags"] = request.Tags
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEvaluationTask"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/evaluation-task/" + dara.PercentEncode(dara.StringValue(agentSpace)) + "/" + dara.PercentEncode(dara.StringValue(taskId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEvaluationTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an evaluation task.
//
// @param request - UpdateEvaluationTaskRequest
//
// @return UpdateEvaluationTaskResponse
func (client *Client) UpdateEvaluationTask(agentSpace *string, taskId *string, request *UpdateEvaluationTaskRequest) (_result *UpdateEvaluationTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateEvaluationTaskResponse{}
	_body, _err := client.UpdateEvaluationTaskWithOptions(agentSpace, taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an evaluator.
//
// @param request - UpdateEvaluatorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEvaluatorResponse
func (client *Client) UpdateEvaluatorWithOptions(agentSpace *string, name *string, request *UpdateEvaluatorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateEvaluatorResponse, _err error) {
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
	if !dara.IsNil(request.Annotations) {
		body["annotations"] = request.Annotations
	}

	if !dara.IsNil(request.Config) {
		body["config"] = request.Config
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DisplayName) {
		body["displayName"] = request.DisplayName
	}

	if !dara.IsNil(request.Properties) {
		body["properties"] = request.Properties
	}

	if !dara.IsNil(request.Version) {
		body["version"] = request.Version
	}

	if !dara.IsNil(request.VersionDescription) {
		body["versionDescription"] = request.VersionDescription
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEvaluator"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/evaluators/" + dara.PercentEncode(dara.StringValue(agentSpace)) + "/" + dara.PercentEncode(dara.StringValue(name))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEvaluatorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an evaluator.
//
// @param request - UpdateEvaluatorRequest
//
// @return UpdateEvaluatorResponse
func (client *Client) UpdateEvaluator(agentSpace *string, name *string, request *UpdateEvaluatorRequest) (_result *UpdateEvaluatorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateEvaluatorResponse{}
	_body, _err := client.UpdateEvaluatorWithOptions(agentSpace, name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an evaluator skill.
//
// @param request - UpdateEvaluatorSkillRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateEvaluatorSkillResponse
func (client *Client) UpdateEvaluatorSkillWithOptions(name *string, skillName *string, request *UpdateEvaluatorSkillRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateEvaluatorSkillResponse, _err error) {
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

	if !dara.IsNil(request.ClientToken) {
		query["clientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DisplayName) {
		body["displayName"] = request.DisplayName
	}

	if !dara.IsNil(request.Enable) {
		body["enable"] = request.Enable
	}

	if !dara.IsNil(request.Files) {
		body["files"] = request.Files
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateEvaluatorSkill"),
		Version:     dara.String("2026-05-20"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/evaluator/" + dara.PercentEncode(dara.StringValue(name)) + "/skill/" + dara.PercentEncode(dara.StringValue(skillName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateEvaluatorSkillResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an evaluator skill.
//
// @param request - UpdateEvaluatorSkillRequest
//
// @return UpdateEvaluatorSkillResponse
func (client *Client) UpdateEvaluatorSkill(name *string, skillName *string, request *UpdateEvaluatorSkillRequest) (_result *UpdateEvaluatorSkillResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateEvaluatorSkillResponse{}
	_body, _err := client.UpdateEvaluatorSkillWithOptions(name, skillName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a pipeline.
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
// Updates a pipeline.
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
