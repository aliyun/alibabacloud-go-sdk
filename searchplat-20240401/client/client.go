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
		"eu-central-1": dara.String("searchplat.eu-central-1.aliyuncs.com"),
		"cn-shanghai":  dara.String("searchplat.cn-shanghai.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("searchplat"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// Stops a service.
//
// @param request - CeaseFunctionInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CeaseFunctionInstanceResponse
func (client *Client) CeaseFunctionInstanceWithOptions(workspaceName *string, functionName *string, instanceName *string, request *CeaseFunctionInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CeaseFunctionInstanceResponse, _err error) {
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
		Action:      dara.String("CeaseFunctionInstance"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/instances/" + dara.PercentEncode(dara.StringValue(instanceName)) + "/actions/cease"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CeaseFunctionInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a service.
//
// @param request - CeaseFunctionInstanceRequest
//
// @return CeaseFunctionInstanceResponse
func (client *Client) CeaseFunctionInstance(workspaceName *string, functionName *string, instanceName *string, request *CeaseFunctionInstanceRequest) (_result *CeaseFunctionInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CeaseFunctionInstanceResponse{}
	_body, _err := client.CeaseFunctionInstanceWithOptions(workspaceName, functionName, instanceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an asynchronous task.
//
// @param request - CreateAsyncTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAsyncTaskResponse
func (client *Client) CreateAsyncTaskWithOptions(workspaceName *string, request *CreateAsyncTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAsyncTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.DataId) {
		body["dataId"] = request.DataId
	}

	if !dara.IsNil(request.Id) {
		body["id"] = request.Id
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.ServiceId) {
		body["serviceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ServiceType) {
		body["serviceType"] = request.ServiceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAsyncTask"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/async-tasks"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAsyncTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an asynchronous task.
//
// @param request - CreateAsyncTaskRequest
//
// @return CreateAsyncTaskResponse
func (client *Client) CreateAsyncTask(workspaceName *string, request *CreateAsyncTaskRequest) (_result *CreateAsyncTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAsyncTaskResponse{}
	_body, _err := client.CreateAsyncTaskWithOptions(workspaceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a knowledge base-related configuration.
//
// @param request - CreateCapabilityRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCapabilityResponse
func (client *Client) CreateCapabilityWithOptions(workspaceName *string, itemCategory *string, request *CreateCapabilityRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateCapabilityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ItemDesc) {
		body["itemDesc"] = request.ItemDesc
	}

	if !dara.IsNil(request.ItemName) {
		body["itemName"] = request.ItemName
	}

	if !dara.IsNil(request.ItemValue) {
		body["itemValue"] = request.ItemValue
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCapability"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/capabilities/" + dara.PercentEncode(dara.StringValue(itemCategory)) + "/items"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCapabilityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a knowledge base-related configuration.
//
// @param request - CreateCapabilityRequest
//
// @return CreateCapabilityResponse
func (client *Client) CreateCapability(workspaceName *string, itemCategory *string, request *CreateCapabilityRequest) (_result *CreateCapabilityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateCapabilityResponse{}
	_body, _err := client.CreateCapabilityWithOptions(workspaceName, itemCategory, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a configuration item in a specified workspace. The prompt and lark types are supported.
//
// Description:
//
// ## Operation description
//
// - This API operation allows you to create a configuration for a specific workspace.
//
// - The `configType` parameter specifies the type of configuration to create. Valid values: `prompt` and `lark`.
//
// - When `dryRun` is set to `true`, the API operation only validates the request without actually performing the creation.
//
// - The `configData` field varies depending on the value of `configType`. Refer to the examples for the specific structure to construct the request body.
//
// @param request - CreateConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateConfigResponse
func (client *Client) CreateConfigWithOptions(workspaceName *string, configType *string, request *CreateConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ConfigData) {
		body["configData"] = request.ConfigData
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateConfig"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/configs/" + dara.PercentEncode(dara.StringValue(configType))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a configuration item in a specified workspace. The prompt and lark types are supported.
//
// Description:
//
// ## Operation description
//
// - This API operation allows you to create a configuration for a specific workspace.
//
// - The `configType` parameter specifies the type of configuration to create. Valid values: `prompt` and `lark`.
//
// - When `dryRun` is set to `true`, the API operation only validates the request without actually performing the creation.
//
// - The `configData` field varies depending on the value of `configType`. Refer to the examples for the specific structure to construct the request body.
//
// @param request - CreateConfigRequest
//
// @return CreateConfigResponse
func (client *Client) CreateConfig(workspaceName *string, configType *string, request *CreateConfigRequest) (_result *CreateConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateConfigResponse{}
	_body, _err := client.CreateConfigWithOptions(workspaceName, configType, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates access credentials.
//
// @param request - CreateCredentialsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCredentialsResponse
func (client *Client) CreateCredentialsWithOptions(workspaceName *string, request *CreateCredentialsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateCredentialsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCredentials"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/credentials"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCredentialsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates access credentials.
//
// @param request - CreateCredentialsRequest
//
// @return CreateCredentialsResponse
func (client *Client) CreateCredentials(workspaceName *string, request *CreateCredentialsRequest) (_result *CreateCredentialsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateCredentialsResponse{}
	_body, _err := client.CreateCredentialsWithOptions(workspaceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates experience data.
//
// @param request - CreateExperienceDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateExperienceDataResponse
func (client *Client) CreateExperienceDataWithOptions(workspaceName *string, request *CreateExperienceDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateExperienceDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ContentType) {
		body["contentType"] = request.ContentType
	}

	if !dara.IsNil(request.DataSize) {
		body["dataSize"] = request.DataSize
	}

	if !dara.IsNil(request.DataType) {
		body["dataType"] = request.DataType
	}

	if !dara.IsNil(request.DataValue) {
		body["dataValue"] = request.DataValue
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.ServiceType) {
		body["serviceType"] = request.ServiceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateExperienceData"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/experience-data"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateExperienceDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates experience data.
//
// @param request - CreateExperienceDataRequest
//
// @return CreateExperienceDataResponse
func (client *Client) CreateExperienceData(workspaceName *string, request *CreateExperienceDataRequest) (_result *CreateExperienceDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateExperienceDataResponse{}
	_body, _err := client.CreateExperienceDataWithOptions(workspaceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a service configuration.
//
// @param request - CreateFunctionInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFunctionInstanceResponse
func (client *Client) CreateFunctionInstanceWithOptions(workspaceName *string, functionName *string, request *CreateFunctionInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateFunctionInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateParameters) {
		body["createParameters"] = request.CreateParameters
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.FunctionType) {
		body["functionType"] = request.FunctionType
	}

	if !dara.IsNil(request.InstanceName) {
		body["instanceName"] = request.InstanceName
	}

	if !dara.IsNil(request.ModelType) {
		body["modelType"] = request.ModelType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFunctionInstance"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/instances"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFunctionInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a service configuration.
//
// @param request - CreateFunctionInstanceRequest
//
// @return CreateFunctionInstanceResponse
func (client *Client) CreateFunctionInstance(workspaceName *string, functionName *string, request *CreateFunctionInstanceRequest) (_result *CreateFunctionInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateFunctionInstanceResponse{}
	_body, _err := client.CreateFunctionInstanceWithOptions(workspaceName, functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a service configuration task.
//
// @param request - CreateFunctionTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFunctionTaskResponse
func (client *Client) CreateFunctionTaskWithOptions(workspaceName *string, functionName *string, instanceName *string, request *CreateFunctionTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateFunctionTaskResponse, _err error) {
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
		Action:      dara.String("CreateFunctionTask"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/instances/" + dara.PercentEncode(dara.StringValue(instanceName)) + "/tasks"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFunctionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a service configuration task.
//
// @param request - CreateFunctionTaskRequest
//
// @return CreateFunctionTaskResponse
func (client *Client) CreateFunctionTask(workspaceName *string, functionName *string, instanceName *string, request *CreateFunctionTaskRequest) (_result *CreateFunctionTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateFunctionTaskResponse{}
	_body, _err := client.CreateFunctionTaskWithOptions(workspaceName, functionName, instanceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an offline processing task for video retrieval. You can configure the data source, processing parameters, and output destination.
//
// Description:
//
// ## Operation description.
//
// @param request - CreateOfflineTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOfflineTaskResponse
func (client *Client) CreateOfflineTaskWithOptions(workspaceName *string, _type *string, request *CreateOfflineTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateOfflineTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Draft) {
		query["draft"] = request.Draft
	}

	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Meta) {
		body["meta"] = request.Meta
	}

	if !dara.IsNil(request.Parameters) {
		body["parameters"] = request.Parameters
	}

	if !dara.IsNil(request.Processors) {
		body["processors"] = request.Processors
	}

	if !dara.IsNil(request.Sink) {
		body["sink"] = request.Sink
	}

	if !dara.IsNil(request.Source) {
		body["source"] = request.Source
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateOfflineTask"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/offline-tasks/" + dara.PercentEncode(dara.StringValue(_type))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateOfflineTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an offline processing task for video retrieval. You can configure the data source, processing parameters, and output destination.
//
// Description:
//
// ## Operation description.
//
// @param request - CreateOfflineTaskRequest
//
// @return CreateOfflineTaskResponse
func (client *Client) CreateOfflineTask(workspaceName *string, _type *string, request *CreateOfflineTaskRequest) (_result *CreateOfflineTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateOfflineTaskResponse{}
	_body, _err := client.CreateOfflineTaskWithOptions(workspaceName, _type, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an evaluation task for the RAG edition.
//
// @param request - CreateRagEvaluatorTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRagEvaluatorTaskResponse
func (client *Client) CreateRagEvaluatorTaskWithOptions(workspaceName *string, request *CreateRagEvaluatorTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateRagEvaluatorTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AppName) {
		body["app_name"] = request.AppName
	}

	if !dara.IsNil(request.Data) {
		body["data"] = request.Data
	}

	if !dara.IsNil(request.DataSourceConfig) {
		body["data_source_config"] = request.DataSourceConfig
	}

	if !dara.IsNil(request.Emails) {
		body["emails"] = request.Emails
	}

	if !dara.IsNil(request.EvaluateConfig) {
		body["evaluate_config"] = request.EvaluateConfig
	}

	if !dara.IsNil(request.HasDataSource) {
		body["has_data_source"] = request.HasDataSource
	}

	if !dara.IsNil(request.Metrics) {
		body["metrics"] = request.Metrics
	}

	if !dara.IsNil(request.TaskName) {
		body["task_name"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRagEvaluatorTask"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/rag-evaluator/v1/api/task"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRagEvaluatorTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an evaluation task for the RAG edition.
//
// @param request - CreateRagEvaluatorTaskRequest
//
// @return CreateRagEvaluatorTaskResponse
func (client *Client) CreateRagEvaluatorTask(workspaceName *string, request *CreateRagEvaluatorTaskRequest) (_result *CreateRagEvaluatorTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateRagEvaluatorTaskResponse{}
	_body, _err := client.CreateRagEvaluatorTaskWithOptions(workspaceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Workspace
//
// @param request - CreateWorkspaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateWorkspaceResponse
func (client *Client) CreateWorkspaceWithOptions(request *CreateWorkspaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateWorkspaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ChargeType) {
		body["chargeType"] = request.ChargeType
	}

	if !dara.IsNil(request.EngineType) {
		body["engineType"] = request.EngineType
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.Quota) {
		body["quota"] = request.Quota
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWorkspace"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Workspace
//
// @param request - CreateWorkspaceRequest
//
// @return CreateWorkspaceResponse
func (client *Client) CreateWorkspace(request *CreateWorkspaceRequest) (_result *CreateWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateWorkspaceResponse{}
	_body, _err := client.CreateWorkspaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specific configuration item from a specified workspace.
//
// Description:
//
// ## Operation description
//
// This API operation allows you to delete a specific configuration item by specifying the workspace name, configuration category, and configuration name. Before calling this operation, ensure that you have sufficient permissions (such as the `DeleteCapability` action in a RAM policy). After a configuration item is deleted, all related data and services may be affected.
//
// @param request - DeleteCapabilityRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCapabilityResponse
func (client *Client) DeleteCapabilityWithOptions(workspaceName *string, itemCategory *string, itemName *string, request *DeleteCapabilityRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteCapabilityResponse, _err error) {
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
		Action:      dara.String("DeleteCapability"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/capabilities/" + dara.PercentEncode(dara.StringValue(itemCategory)) + "/items/" + dara.PercentEncode(dara.StringValue(itemName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCapabilityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specific configuration item from a specified workspace.
//
// Description:
//
// ## Operation description
//
// This API operation allows you to delete a specific configuration item by specifying the workspace name, configuration category, and configuration name. Before calling this operation, ensure that you have sufficient permissions (such as the `DeleteCapability` action in a RAM policy). After a configuration item is deleted, all related data and services may be affected.
//
// @param request - DeleteCapabilityRequest
//
// @return DeleteCapabilityResponse
func (client *Client) DeleteCapability(workspaceName *string, itemCategory *string, itemName *string, request *DeleteCapabilityRequest) (_result *DeleteCapabilityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteCapabilityResponse{}
	_body, _err := client.DeleteCapabilityWithOptions(workspaceName, itemCategory, itemName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a specific type of configuration from a specified workspace.
//
// Description:
//
// ## Request description.
//
// @param request - DeleteConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteConfigResponse
func (client *Client) DeleteConfigWithOptions(workspaceName *string, configType *string, id *string, request *DeleteConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteConfigResponse, _err error) {
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
		Action:      dara.String("DeleteConfig"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/configs/" + dara.PercentEncode(dara.StringValue(configType)) + "/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a specific type of configuration from a specified workspace.
//
// Description:
//
// ## Request description.
//
// @param request - DeleteConfigRequest
//
// @return DeleteConfigResponse
func (client *Client) DeleteConfig(workspaceName *string, configType *string, id *string, request *DeleteConfigRequest) (_result *DeleteConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteConfigResponse{}
	_body, _err := client.DeleteConfigWithOptions(workspaceName, configType, id, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an access credential.
//
// @param request - DeleteCredentialsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCredentialsResponse
func (client *Client) DeleteCredentialsWithOptions(token *string, workspaceName *string, request *DeleteCredentialsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteCredentialsResponse, _err error) {
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
		Action:      dara.String("DeleteCredentials"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/credentials/" + dara.PercentEncode(dara.StringValue(token))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCredentialsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an access credential.
//
// @param request - DeleteCredentialsRequest
//
// @return DeleteCredentialsResponse
func (client *Client) DeleteCredentials(token *string, workspaceName *string, request *DeleteCredentialsRequest) (_result *DeleteCredentialsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteCredentialsResponse{}
	_body, _err := client.DeleteCredentialsWithOptions(token, workspaceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete experience data
//
// @param request - DeleteExperienceDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteExperienceDataResponse
func (client *Client) DeleteExperienceDataWithOptions(id *string, workspaceName *string, request *DeleteExperienceDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteExperienceDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteExperienceData"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/experience-data/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteExperienceDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete experience data
//
// @param request - DeleteExperienceDataRequest
//
// @return DeleteExperienceDataResponse
func (client *Client) DeleteExperienceData(id *string, workspaceName *string, request *DeleteExperienceDataRequest) (_result *DeleteExperienceDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteExperienceDataResponse{}
	_body, _err := client.DeleteExperienceDataWithOptions(id, workspaceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a service configuration.
//
// @param request - DeleteFunctionInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFunctionInstanceResponse
func (client *Client) DeleteFunctionInstanceWithOptions(workspaceName *string, functionName *string, instanceName *string, request *DeleteFunctionInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteFunctionInstanceResponse, _err error) {
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
		Action:      dara.String("DeleteFunctionInstance"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/instances/" + dara.PercentEncode(dara.StringValue(instanceName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFunctionInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a service configuration.
//
// @param request - DeleteFunctionInstanceRequest
//
// @return DeleteFunctionInstanceResponse
func (client *Client) DeleteFunctionInstance(workspaceName *string, functionName *string, instanceName *string, request *DeleteFunctionInstanceRequest) (_result *DeleteFunctionInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteFunctionInstanceResponse{}
	_body, _err := client.DeleteFunctionInstanceWithOptions(workspaceName, functionName, instanceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a batch task.
//
// @param request - DeleteOfflineTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteOfflineTaskResponse
func (client *Client) DeleteOfflineTaskWithOptions(workspaceName *string, _type *string, taskName *string, request *DeleteOfflineTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteOfflineTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteOfflineTask"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/offline-tasks/" + dara.PercentEncode(dara.StringValue(_type)) + "/" + dara.PercentEncode(dara.StringValue(taskName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteOfflineTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a batch task.
//
// @param request - DeleteOfflineTaskRequest
//
// @return DeleteOfflineTaskResponse
func (client *Client) DeleteOfflineTask(workspaceName *string, _type *string, taskName *string, request *DeleteOfflineTaskRequest) (_result *DeleteOfflineTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteOfflineTaskResponse{}
	_body, _err := client.DeleteOfflineTaskWithOptions(workspaceName, _type, taskName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a RAG evaluation task.
//
// @param request - DeleteRagEvaluatorTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRagEvaluatorTaskResponse
func (client *Client) DeleteRagEvaluatorTaskWithOptions(workspaceName *string, taskId *string, request *DeleteRagEvaluatorTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteRagEvaluatorTaskResponse, _err error) {
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
		Action:      dara.String("DeleteRagEvaluatorTask"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/rag-evaluator/v1/api/task/" + dara.PercentEncode(dara.StringValue(taskId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteRagEvaluatorTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a RAG evaluation task.
//
// @param request - DeleteRagEvaluatorTaskRequest
//
// @return DeleteRagEvaluatorTaskResponse
func (client *Client) DeleteRagEvaluatorTask(workspaceName *string, taskId *string, request *DeleteRagEvaluatorTaskRequest) (_result *DeleteRagEvaluatorTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteRagEvaluatorTaskResponse{}
	_body, _err := client.DeleteRagEvaluatorTaskWithOptions(workspaceName, taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a workspace.
//
// @param request - DeleteWorkspaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkspaceResponse
func (client *Client) DeleteWorkspaceWithOptions(workspaceName *string, request *DeleteWorkspaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteWorkspaceResponse, _err error) {
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
		Action:      dara.String("DeleteWorkspace"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a workspace.
//
// @param request - DeleteWorkspaceRequest
//
// @return DeleteWorkspaceResponse
func (client *Client) DeleteWorkspace(workspaceName *string, request *DeleteWorkspaceRequest) (_result *DeleteWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteWorkspaceResponse{}
	_body, _err := client.DeleteWorkspaceWithOptions(workspaceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Query the details of a configuration item of a specific category within a specified workspace.
//
// Description:
//
// ## Request Description
//
// This API is used to retrieve specific configuration information based on the provided workspace name, configuration category, and configuration name. Please ensure the parameters in the request path are accurate, especially the three required fields: `workspaceName`, `itemCategory`, and `itemName`. Additionally, please note that `itemCategory` currently only supports the `ai_search_agent` category.
//
// @param request - DescribeCapabilityRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCapabilityResponse
func (client *Client) DescribeCapabilityWithOptions(workspaceName *string, itemCategory *string, itemName *string, request *DescribeCapabilityRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeCapabilityResponse, _err error) {
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
		Action:      dara.String("DescribeCapability"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/capabilities/" + dara.PercentEncode(dara.StringValue(itemCategory)) + "/items/" + dara.PercentEncode(dara.StringValue(itemName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeCapabilityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Query the details of a configuration item of a specific category within a specified workspace.
//
// Description:
//
// ## Request Description
//
// This API is used to retrieve specific configuration information based on the provided workspace name, configuration category, and configuration name. Please ensure the parameters in the request path are accurate, especially the three required fields: `workspaceName`, `itemCategory`, and `itemName`. Additionally, please note that `itemCategory` currently only supports the `ai_search_agent` category.
//
// @param request - DescribeCapabilityRequest
//
// @return DescribeCapabilityResponse
func (client *Client) DescribeCapability(workspaceName *string, itemCategory *string, itemName *string, request *DescribeCapabilityRequest) (_result *DescribeCapabilityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeCapabilityResponse{}
	_body, _err := client.DescribeCapabilityWithOptions(workspaceName, itemCategory, itemName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// DescribeRegions.
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
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/regions"),
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
// DescribeRegions.
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
// Retrieves the details of an asynchronous task.
//
// @param request - GetAsyncTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAsyncTaskResponse
func (client *Client) GetAsyncTaskWithOptions(workspaceName *string, id *string, request *GetAsyncTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAsyncTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAsyncTask"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/async-tasks/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAsyncTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of an asynchronous task.
//
// @param request - GetAsyncTaskRequest
//
// @return GetAsyncTaskResponse
func (client *Client) GetAsyncTask(workspaceName *string, id *string, request *GetAsyncTaskRequest) (_result *GetAsyncTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAsyncTaskResponse{}
	_body, _err := client.GetAsyncTaskWithOptions(workspaceName, id, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a specific type of configuration from a specified workspace.
//
// Description:
//
// ## Request description.
//
// @param request - GetConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetConfigResponse
func (client *Client) GetConfigWithOptions(workspaceName *string, configType *string, id *string, request *GetConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetConfigResponse, _err error) {
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
		Action:      dara.String("GetConfig"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/configs/" + dara.PercentEncode(dara.StringValue(configType)) + "/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a specific type of configuration from a specified workspace.
//
// Description:
//
// ## Request description.
//
// @param request - GetConfigRequest
//
// @return GetConfigResponse
func (client *Client) GetConfig(workspaceName *string, configType *string, id *string, request *GetConfigRequest) (_result *GetConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetConfigResponse{}
	_body, _err := client.GetConfigWithOptions(workspaceName, configType, id, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of an access credential.
//
// @param request - GetCredentialsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCredentialsResponse
func (client *Client) GetCredentialsWithOptions(token *string, workspaceName *string, request *GetCredentialsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCredentialsResponse, _err error) {
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
		Action:      dara.String("GetCredentials"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/credentials/" + dara.PercentEncode(dara.StringValue(token))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCredentialsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of an access credential.
//
// @param request - GetCredentialsRequest
//
// @return GetCredentialsResponse
func (client *Client) GetCredentials(token *string, workspaceName *string, request *GetCredentialsRequest) (_result *GetCredentialsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCredentialsResponse{}
	_body, _err := client.GetCredentialsWithOptions(token, workspaceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of experience data.
//
// @param request - GetExperienceDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetExperienceDataResponse
func (client *Client) GetExperienceDataWithOptions(workspaceName *string, id *string, request *GetExperienceDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetExperienceDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetExperienceData"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/experience-data/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetExperienceDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of experience data.
//
// @param request - GetExperienceDataRequest
//
// @return GetExperienceDataResponse
func (client *Client) GetExperienceData(workspaceName *string, id *string, request *GetExperienceDataRequest) (_result *GetExperienceDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetExperienceDataResponse{}
	_body, _err := client.GetExperienceDataWithOptions(workspaceName, id, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a specific feature instance in a specified workspace.
//
// @param request - GetFunctionInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFunctionInstanceResponse
func (client *Client) GetFunctionInstanceWithOptions(workspaceName *string, functionName *string, instanceName *string, request *GetFunctionInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetFunctionInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Output) {
		query["output"] = request.Output
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFunctionInstance"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/instances/" + dara.PercentEncode(dara.StringValue(instanceName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFunctionInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of a specific feature instance in a specified workspace.
//
// @param request - GetFunctionInstanceRequest
//
// @return GetFunctionInstanceResponse
func (client *Client) GetFunctionInstance(workspaceName *string, functionName *string, instanceName *string, request *GetFunctionInstanceRequest) (_result *GetFunctionInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFunctionInstanceResponse{}
	_body, _err := client.GetFunctionInstanceWithOptions(workspaceName, functionName, instanceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves information about an offline node.
//
// @param request - GetOfflineTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOfflineTaskResponse
func (client *Client) GetOfflineTaskWithOptions(workspaceName *string, _type *string, taskName *string, request *GetOfflineTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetOfflineTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOfflineTask"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/offline-tasks/" + dara.PercentEncode(dara.StringValue(_type)) + "/" + dara.PercentEncode(dara.StringValue(taskName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOfflineTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves information about an offline node.
//
// @param request - GetOfflineTaskRequest
//
// @return GetOfflineTaskResponse
func (client *Client) GetOfflineTask(workspaceName *string, _type *string, taskName *string, request *GetOfflineTaskRequest) (_result *GetOfflineTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetOfflineTaskResponse{}
	_body, _err := client.GetOfflineTaskWithOptions(workspaceName, _type, taskName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of offline task logs in a specified workspace.
//
// Description:
//
// ## Operation description
//
// - This operation allows you to retrieve offline task logs information by specifying the workspace name, node type, and node name.
//
// - Provide a valid `regionId` as one of the query parameters to specify the area for the request.
//
// - The returned information includes but is not limited to network configurations (private ES and public ES) and their enabling status, domain names, and IP whitelist groups.
//
// - Note: Ensure that you have sufficient permissions (such as the `GetLog` action in the RAM policy) to invoke this operation.
//
// @param request - GetOfflineTaskLogRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOfflineTaskLogResponse
func (client *Client) GetOfflineTaskLogWithOptions(workspaceName *string, _type *string, taskName *string, request *GetOfflineTaskLogRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetOfflineTaskLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOfflineTaskLog"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/offline-tasks/" + dara.PercentEncode(dara.StringValue(_type)) + "/" + dara.PercentEncode(dara.StringValue(taskName)) + "/log"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOfflineTaskLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the details of offline task logs in a specified workspace.
//
// Description:
//
// ## Operation description
//
// - This operation allows you to retrieve offline task logs information by specifying the workspace name, node type, and node name.
//
// - Provide a valid `regionId` as one of the query parameters to specify the area for the request.
//
// - The returned information includes but is not limited to network configurations (private ES and public ES) and their enabling status, domain names, and IP whitelist groups.
//
// - Note: Ensure that you have sufficient permissions (such as the `GetLog` action in the RAM policy) to invoke this operation.
//
// @param request - GetOfflineTaskLogRequest
//
// @return GetOfflineTaskLogResponse
func (client *Client) GetOfflineTaskLog(workspaceName *string, _type *string, taskName *string, request *GetOfflineTaskLogRequest) (_result *GetOfflineTaskLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetOfflineTaskLogResponse{}
	_body, _err := client.GetOfflineTaskLogWithOptions(workspaceName, _type, taskName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a RAG evaluation task.
//
// @param request - GetRagEvaluatorTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRagEvaluatorTaskResponse
func (client *Client) GetRagEvaluatorTaskWithOptions(workspaceName *string, taskId *string, request *GetRagEvaluatorTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetRagEvaluatorTaskResponse, _err error) {
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
		Action:      dara.String("GetRagEvaluatorTask"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/rag-evaluator/v1/api/task/" + dara.PercentEncode(dara.StringValue(taskId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetRagEvaluatorTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a RAG evaluation task.
//
// @param request - GetRagEvaluatorTaskRequest
//
// @return GetRagEvaluatorTaskResponse
func (client *Client) GetRagEvaluatorTask(workspaceName *string, taskId *string, request *GetRagEvaluatorTaskRequest) (_result *GetRagEvaluatorTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRagEvaluatorTaskResponse{}
	_body, _err := client.GetRagEvaluatorTaskWithOptions(workspaceName, taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves data table field information.
//
// @param request - GetTableColumnsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableColumnsResponse
func (client *Client) GetTableColumnsWithOptions(workspaceName *string, dataSourceType *string, request *GetTableColumnsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTableColumnsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Params) {
		query["params"] = request.Params
	}

	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTableColumns"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/data-sources/" + dara.PercentEncode(dara.StringValue(dataSourceType)) + "/columns"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTableColumnsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves data table field information.
//
// @param request - GetTableColumnsRequest
//
// @return GetTableColumnsResponse
func (client *Client) GetTableColumns(workspaceName *string, dataSourceType *string, request *GetTableColumnsRequest) (_result *GetTableColumnsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTableColumnsResponse{}
	_body, _err := client.GetTableColumnsWithOptions(workspaceName, dataSourceType, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - GetTableFieldsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTableFieldsResponse
func (client *Client) GetTableFieldsWithOptions(workspaceName *string, dataSourceType *string, request *GetTableFieldsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTableFieldsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Params) {
		query["params"] = request.Params
	}

	if !dara.IsNil(request.RawType) {
		query["rawType"] = request.RawType
	}

	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTableFields"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/data-sources/" + dara.PercentEncode(dara.StringValue(dataSourceType)) + "/fields"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTableFieldsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetTableFieldsRequest
//
// @return GetTableFieldsResponse
func (client *Client) GetTableFields(workspaceName *string, dataSourceType *string, request *GetTableFieldsRequest) (_result *GetTableFieldsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTableFieldsResponse{}
	_body, _err := client.GetTableFieldsWithOptions(workspaceName, dataSourceType, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves data tables.
//
// @param request - GetTablesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTablesResponse
func (client *Client) GetTablesWithOptions(workspaceName *string, dataSourceType *string, request *GetTablesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTablesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Params) {
		query["params"] = request.Params
	}

	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTables"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/data-sources/" + dara.PercentEncode(dara.StringValue(dataSourceType)) + "/tables"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves data tables.
//
// @param request - GetTablesRequest
//
// @return GetTablesResponse
func (client *Client) GetTables(workspaceName *string, dataSourceType *string, request *GetTablesRequest) (_result *GetTablesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTablesResponse{}
	_body, _err := client.GetTablesWithOptions(workspaceName, dataSourceType, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a workspace.
//
// @param request - GetWorkspaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkspaceResponse
func (client *Client) GetWorkspaceWithOptions(workspaceName *string, request *GetWorkspaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetWorkspaceResponse, _err error) {
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
		Action:      dara.String("GetWorkspace"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a workspace.
//
// @param request - GetWorkspaceRequest
//
// @return GetWorkspaceResponse
func (client *Client) GetWorkspace(workspaceName *string, request *GetWorkspaceRequest) (_result *GetWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetWorkspaceResponse{}
	_body, _err := client.GetWorkspaceWithOptions(workspaceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of asynchronous tasks.
//
// @param request - ListAsyncTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAsyncTasksResponse
func (client *Client) ListAsyncTasksWithOptions(workspaceName *string, request *ListAsyncTasksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAsyncTasksResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataId) {
		query["dataId"] = request.DataId
	}

	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	if !dara.IsNil(request.ServiceType) {
		query["serviceType"] = request.ServiceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAsyncTasks"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/async-tasks"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAsyncTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of asynchronous tasks.
//
// @param request - ListAsyncTasksRequest
//
// @return ListAsyncTasksResponse
func (client *Client) ListAsyncTasks(workspaceName *string, request *ListAsyncTasksRequest) (_result *ListAsyncTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAsyncTasksResponse{}
	_body, _err := client.ListAsyncTasksWithOptions(workspaceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of configuration items of a specific category in a specified workspace.
//
// Description:
//
// ## Operation description
//
// You can use this API operation to retrieve the list of configuration items based on the specified workspace name and configuration category. Paged query is supported. Use the `pageNumber` and `pageSize` parameters to control the number of results and the page number. The `nextToken` and `maxResults` parameters are also provided for paged query when handling large amounts of data.
//
// @param request - ListCapabilitiesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCapabilitiesResponse
func (client *Client) ListCapabilitiesWithOptions(workspaceName *string, itemCategory *string, request *ListCapabilitiesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListCapabilitiesResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCapabilities"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/capabilities/" + dara.PercentEncode(dara.StringValue(itemCategory)) + "/items"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCapabilitiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of configuration items of a specific category in a specified workspace.
//
// Description:
//
// ## Operation description
//
// You can use this API operation to retrieve the list of configuration items based on the specified workspace name and configuration category. Paged query is supported. Use the `pageNumber` and `pageSize` parameters to control the number of results and the page number. The `nextToken` and `maxResults` parameters are also provided for paged query when handling large amounts of data.
//
// @param request - ListCapabilitiesRequest
//
// @return ListCapabilitiesResponse
func (client *Client) ListCapabilities(workspaceName *string, itemCategory *string, request *ListCapabilitiesRequest) (_result *ListCapabilitiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListCapabilitiesResponse{}
	_body, _err := client.ListCapabilitiesWithOptions(workspaceName, itemCategory, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the list of configurations of a specific type in a specified workspace.
//
// Description:
//
// ## Request description.
//
// @param request - ListConfigsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListConfigsResponse
func (client *Client) ListConfigsWithOptions(workspaceName *string, configType *string, request *ListConfigsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListConfigsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Filter) {
		query["filter"] = request.Filter
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListConfigs"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/configs/" + dara.PercentEncode(dara.StringValue(configType))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListConfigsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the list of configurations of a specific type in a specified workspace.
//
// Description:
//
// ## Request description.
//
// @param request - ListConfigsRequest
//
// @return ListConfigsResponse
func (client *Client) ListConfigs(workspaceName *string, configType *string, request *ListConfigsRequest) (_result *ListConfigsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListConfigsResponse{}
	_body, _err := client.ListConfigsWithOptions(workspaceName, configType, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of access credentials.
//
// @param request - ListCredentialsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCredentialsResponse
func (client *Client) ListCredentialsWithOptions(workspaceName *string, request *ListCredentialsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListCredentialsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCredentials"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/credentials"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCredentialsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of access credentials.
//
// @param request - ListCredentialsRequest
//
// @return ListCredentialsResponse
func (client *Client) ListCredentials(workspaceName *string, request *ListCredentialsRequest) (_result *ListCredentialsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListCredentialsResponse{}
	_body, _err := client.ListCredentialsWithOptions(workspaceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves all experience data in a specified workspace, with support for filtering by service type and data type.
//
// Description:
//
// ## Operation description
//
// - This API operation queries all experience data of a user in a specific workspace. The results are sorted by creation time in descending order by default.
//
// - Pagination is not supported. However, you can filter data by using the serviceType and dataType parameters.
//
// - workspaceName is a path parameter and must be specified to indicate the workspace to query.
//
// @param request - ListExperienceDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListExperienceDataResponse
func (client *Client) ListExperienceDataWithOptions(workspaceName *string, request *ListExperienceDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListExperienceDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataType) {
		query["dataType"] = request.DataType
	}

	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	if !dara.IsNil(request.ServiceType) {
		query["serviceType"] = request.ServiceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListExperienceData"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/experience-data"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListExperienceDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves all experience data in a specified workspace, with support for filtering by service type and data type.
//
// Description:
//
// ## Operation description
//
// - This API operation queries all experience data of a user in a specific workspace. The results are sorted by creation time in descending order by default.
//
// - Pagination is not supported. However, you can filter data by using the serviceType and dataType parameters.
//
// - workspaceName is a path parameter and must be specified to indicate the workspace to query.
//
// @param request - ListExperienceDataRequest
//
// @return ListExperienceDataResponse
func (client *Client) ListExperienceData(workspaceName *string, request *ListExperienceDataRequest) (_result *ListExperienceDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListExperienceDataResponse{}
	_body, _err := client.ListExperienceDataWithOptions(workspaceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of service configurations.
//
// @param request - ListFunctionInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFunctionInstancesResponse
func (client *Client) ListFunctionInstancesWithOptions(workspaceName *string, functionName *string, request *ListFunctionInstancesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListFunctionInstancesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FunctionType) {
		query["functionType"] = request.FunctionType
	}

	if !dara.IsNil(request.ModelType) {
		query["modelType"] = request.ModelType
	}

	if !dara.IsNil(request.Output) {
		query["output"] = request.Output
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Source) {
		query["source"] = request.Source
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFunctionInstances"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/instances"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFunctionInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of service configurations.
//
// @param request - ListFunctionInstancesRequest
//
// @return ListFunctionInstancesResponse
func (client *Client) ListFunctionInstances(workspaceName *string, functionName *string, request *ListFunctionInstancesRequest) (_result *ListFunctionInstancesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFunctionInstancesResponse{}
	_body, _err := client.ListFunctionInstancesWithOptions(workspaceName, functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the restriction items of a feature.
//
// @param request - ListFunctionRestrictionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFunctionRestrictionsResponse
func (client *Client) ListFunctionRestrictionsWithOptions(workspaceName *string, functionName *string, restrictionName *string, request *ListFunctionRestrictionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListFunctionRestrictionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ModelType) {
		query["modelType"] = request.ModelType
	}

	if !dara.IsNil(request.Region) {
		query["region"] = request.Region
	}

	if !dara.IsNil(request.Source) {
		query["source"] = request.Source
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFunctionRestrictions"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/restrictions/" + dara.PercentEncode(dara.StringValue(restrictionName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFunctionRestrictionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the restriction items of a feature.
//
// @param request - ListFunctionRestrictionsRequest
//
// @return ListFunctionRestrictionsResponse
func (client *Client) ListFunctionRestrictions(workspaceName *string, functionName *string, restrictionName *string, request *ListFunctionRestrictionsRequest) (_result *ListFunctionRestrictionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFunctionRestrictionsResponse{}
	_body, _err := client.ListFunctionRestrictionsWithOptions(workspaceName, functionName, restrictionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Offline Task Information List
//
// @param tmpReq - ListOfflineTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOfflineTaskResponse
func (client *Client) ListOfflineTaskWithOptions(workspaceName *string, _type *string, tmpReq *ListOfflineTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListOfflineTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListOfflineTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Labels) {
		request.LabelsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Labels, dara.String("labels"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TaskStatus) {
		request.TaskStatusShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskStatus, dara.String("taskStatus"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.LabelsShrink) {
		query["labels"] = request.LabelsShrink
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	if !dara.IsNil(request.TaskName) {
		query["taskName"] = request.TaskName
	}

	if !dara.IsNil(request.TaskStatusShrink) {
		query["taskStatus"] = request.TaskStatusShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOfflineTask"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/offline-tasks/" + dara.PercentEncode(dara.StringValue(_type))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOfflineTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Offline Task Information List
//
// @param request - ListOfflineTaskRequest
//
// @return ListOfflineTaskResponse
func (client *Client) ListOfflineTask(workspaceName *string, _type *string, request *ListOfflineTaskRequest) (_result *ListOfflineTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListOfflineTaskResponse{}
	_body, _err := client.ListOfflineTaskWithOptions(workspaceName, _type, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the error log list of batch tasks in a specified workspace.
//
// Description:
//
// ## Operation description
//
// - This API operation retrieves error logs for a specific workspace, node type, and node name.
//
// - The `startTime` and `endTime` parameters allow you to define a custom query time range. If not provided, data from the past hour is queried by default.
//
// - The paging parameters `pageNum` and `pageSize` help control the number of returned results and page navigation. They represent the requested page number and the number of log entries per page, with default values of 1 and 10 respectively.
//
// - Note: Ensure that you have obtained the required RAM permissions (Action: ListErrorLogs) before you invoke this operation.
//
// @param request - ListOfflineTaskErrorLogsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListOfflineTaskErrorLogsResponse
func (client *Client) ListOfflineTaskErrorLogsWithOptions(workspaceName *string, _type *string, taskName *string, request *ListOfflineTaskErrorLogsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListOfflineTaskErrorLogsResponse, _err error) {
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

	if !dara.IsNil(request.PageNum) {
		query["pageNum"] = request.PageNum
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListOfflineTaskErrorLogs"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/offline-tasks/" + dara.PercentEncode(dara.StringValue(_type)) + "/" + dara.PercentEncode(dara.StringValue(taskName)) + "/error-logs"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListOfflineTaskErrorLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the error log list of batch tasks in a specified workspace.
//
// Description:
//
// ## Operation description
//
// - This API operation retrieves error logs for a specific workspace, node type, and node name.
//
// - The `startTime` and `endTime` parameters allow you to define a custom query time range. If not provided, data from the past hour is queried by default.
//
// - The paging parameters `pageNum` and `pageSize` help control the number of returned results and page navigation. They represent the requested page number and the number of log entries per page, with default values of 1 and 10 respectively.
//
// - Note: Ensure that you have obtained the required RAM permissions (Action: ListErrorLogs) before you invoke this operation.
//
// @param request - ListOfflineTaskErrorLogsRequest
//
// @return ListOfflineTaskErrorLogsResponse
func (client *Client) ListOfflineTaskErrorLogs(workspaceName *string, _type *string, taskName *string, request *ListOfflineTaskErrorLogsRequest) (_result *ListOfflineTaskErrorLogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListOfflineTaskErrorLogsResponse{}
	_body, _err := client.ListOfflineTaskErrorLogsWithOptions(workspaceName, _type, taskName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the list of RAG evaluation tasks.
//
// @param request - ListRagEvaluatorTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRagEvaluatorTasksResponse
func (client *Client) ListRagEvaluatorTasksWithOptions(workspaceName *string, request *ListRagEvaluatorTasksRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListRagEvaluatorTasksResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListRagEvaluatorTasks"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/rag-evaluator/v1/api/tasks"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListRagEvaluatorTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of RAG evaluation tasks.
//
// @param request - ListRagEvaluatorTasksRequest
//
// @return ListRagEvaluatorTasksResponse
func (client *Client) ListRagEvaluatorTasks(workspaceName *string, request *ListRagEvaluatorTasksRequest) (_result *ListRagEvaluatorTasksResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRagEvaluatorTasksResponse{}
	_body, _err := client.ListRagEvaluatorTasksWithOptions(workspaceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the list of services.
//
// @param request - ListServicesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListServicesResponse
func (client *Client) ListServicesWithOptions(workspaceName *string, request *ListServicesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListServicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ModelType) {
		query["modelType"] = request.ModelType
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.ServiceId) {
		query["serviceId"] = request.ServiceId
	}

	if !dara.IsNil(request.ServiceType) {
		query["serviceType"] = request.ServiceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListServices"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/services"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the list of services.
//
// @param request - ListServicesRequest
//
// @return ListServicesResponse
func (client *Client) ListServices(workspaceName *string, request *ListServicesRequest) (_result *ListServicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListServicesResponse{}
	_body, _err := client.ListServicesWithOptions(workspaceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a list of workspaces.
//
// @param request - ListWorkspacesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListWorkspacesResponse
func (client *Client) ListWorkspacesWithOptions(request *ListWorkspacesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListWorkspacesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["instanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceGroupId) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SortBy) {
		query["sortBy"] = request.SortBy
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkspaces"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListWorkspacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a list of workspaces.
//
// @param request - ListWorkspacesRequest
//
// @return ListWorkspacesResponse
func (client *Client) ListWorkspaces(request *ListWorkspacesRequest) (_result *ListWorkspacesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListWorkspacesResponse{}
	_body, _err := client.ListWorkspacesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a batch task.
//
// @param request - ModifyOfflineTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyOfflineTaskResponse
func (client *Client) ModifyOfflineTaskWithOptions(workspaceName *string, _type *string, taskName *string, request *ModifyOfflineTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyOfflineTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Meta) {
		body["meta"] = request.Meta
	}

	if !dara.IsNil(request.Parameters) {
		body["parameters"] = request.Parameters
	}

	if !dara.IsNil(request.Processors) {
		body["processors"] = request.Processors
	}

	if !dara.IsNil(request.Sink) {
		body["sink"] = request.Sink
	}

	if !dara.IsNil(request.Source) {
		body["source"] = request.Source
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyOfflineTask"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/offline-tasks/" + dara.PercentEncode(dara.StringValue(_type)) + "/" + dara.PercentEncode(dara.StringValue(taskName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyOfflineTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a batch task.
//
// @param request - ModifyOfflineTaskRequest
//
// @return ModifyOfflineTaskResponse
func (client *Client) ModifyOfflineTask(workspaceName *string, _type *string, taskName *string, request *ModifyOfflineTaskRequest) (_result *ModifyOfflineTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyOfflineTaskResponse{}
	_body, _err := client.ModifyOfflineTaskWithOptions(workspaceName, _type, taskName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the network configuration for batch task log scenarios, including enabling or disabling public and private network access and setting IP whitelists.
//
// Description:
//
// ## Operation description
//
// - This API allows you to adjust network-related configurations for a specific type of batch node within a specified workspace.
//
// - Use this operation to control public or private network access permissions for the ES service and set the corresponding IP whitelists.
//
// - When you need to change any network settings (such as enabling or shutting down public network access or updating IP whitelists), ensure that the `network` object contains the correct parameters.
//
// - Note: Executing this operation may affect currently running nodes. Proceed with caution.
//
// @param request - ModifyOfflineTaskLogRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyOfflineTaskLogResponse
func (client *Client) ModifyOfflineTaskLogWithOptions(workspaceName *string, _type *string, taskName *string, request *ModifyOfflineTaskLogRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ModifyOfflineTaskLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Network) {
		body["network"] = request.Network
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ModifyOfflineTaskLog"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/offline-tasks/" + dara.PercentEncode(dara.StringValue(_type)) + "/" + dara.PercentEncode(dara.StringValue(taskName)) + "/log"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ModifyOfflineTaskLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the network configuration for batch task log scenarios, including enabling or disabling public and private network access and setting IP whitelists.
//
// Description:
//
// ## Operation description
//
// - This API allows you to adjust network-related configurations for a specific type of batch node within a specified workspace.
//
// - Use this operation to control public or private network access permissions for the ES service and set the corresponding IP whitelists.
//
// - When you need to change any network settings (such as enabling or shutting down public network access or updating IP whitelists), ensure that the `network` object contains the correct parameters.
//
// - Note: Executing this operation may affect currently running nodes. Proceed with caution.
//
// @param request - ModifyOfflineTaskLogRequest
//
// @return ModifyOfflineTaskLogResponse
func (client *Client) ModifyOfflineTaskLog(workspaceName *string, _type *string, taskName *string, request *ModifyOfflineTaskLogRequest) (_result *ModifyOfflineTaskLogResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyOfflineTaskLogResponse{}
	_body, _err := client.ModifyOfflineTaskLogWithOptions(workspaceName, _type, taskName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restarts the service.
//
// @param request - ResumeFunctionInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeFunctionInstanceResponse
func (client *Client) ResumeFunctionInstanceWithOptions(workspaceName *string, functionName *string, instanceName *string, request *ResumeFunctionInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ResumeFunctionInstanceResponse, _err error) {
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
		Action:      dara.String("ResumeFunctionInstance"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/instances/" + dara.PercentEncode(dara.StringValue(instanceName)) + "/actions/resume"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ResumeFunctionInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restarts the service.
//
// @param request - ResumeFunctionInstanceRequest
//
// @return ResumeFunctionInstanceResponse
func (client *Client) ResumeFunctionInstance(workspaceName *string, functionName *string, instanceName *string, request *ResumeFunctionInstanceRequest) (_result *ResumeFunctionInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ResumeFunctionInstanceResponse{}
	_body, _err := client.ResumeFunctionInstanceWithOptions(workspaceName, functionName, instanceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts a batch task.
//
// Description:
//
// ## Operation description.
//
// @param request - StartOfflineTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartOfflineTaskResponse
func (client *Client) StartOfflineTaskWithOptions(workspaceName *string, _type *string, taskName *string, request *StartOfflineTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StartOfflineTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Parallelism) {
		body["parallelism"] = request.Parallelism
	}

	if !dara.IsNil(request.Timestamp) {
		body["timestamp"] = request.Timestamp
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartOfflineTask"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/offline-tasks/" + dara.PercentEncode(dara.StringValue(_type)) + "/" + dara.PercentEncode(dara.StringValue(taskName)) + "/actions/start"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StartOfflineTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Starts a batch task.
//
// Description:
//
// ## Operation description.
//
// @param request - StartOfflineTaskRequest
//
// @return StartOfflineTaskResponse
func (client *Client) StartOfflineTask(workspaceName *string, _type *string, taskName *string, request *StartOfflineTaskRequest) (_result *StartOfflineTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartOfflineTaskResponse{}
	_body, _err := client.StartOfflineTaskWithOptions(workspaceName, _type, taskName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a batch task.
//
// @param request - StopOfflineTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopOfflineTaskResponse
func (client *Client) StopOfflineTaskWithOptions(workspaceName *string, _type *string, taskName *string, request *StopOfflineTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopOfflineTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["regionId"] = request.RegionId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Parallelism) {
		body["parallelism"] = request.Parallelism
	}

	if !dara.IsNil(request.Timestamp) {
		body["timestamp"] = request.Timestamp
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopOfflineTask"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/offline-tasks/" + dara.PercentEncode(dara.StringValue(_type)) + "/" + dara.PercentEncode(dara.StringValue(taskName)) + "/actions/stop"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopOfflineTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a batch task.
//
// @param request - StopOfflineTaskRequest
//
// @return StopOfflineTaskResponse
func (client *Client) StopOfflineTask(workspaceName *string, _type *string, taskName *string, request *StopOfflineTaskRequest) (_result *StopOfflineTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopOfflineTaskResponse{}
	_body, _err := client.StopOfflineTaskWithOptions(workspaceName, _type, taskName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a specific configuration item in a specified workspace.
//
// Description:
//
// ## Request description
//
// This API operation allows you to update a specific configuration item (`itemName`) under a category (`itemCategory`) in a specified workspace (`workspaceName`). By setting the `dryRun` parameter, you can preview changes without actually applying them. The request body can contain a new configuration description (`itemDesc`) and configuration content (`itemValue`). The structure of `itemValue` must conform to the requirements of the target configuration item.
//
// @param request - UpdateCapabilityRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCapabilityResponse
func (client *Client) UpdateCapabilityWithOptions(workspaceName *string, itemCategory *string, itemName *string, request *UpdateCapabilityRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateCapabilityResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ItemDesc) {
		body["itemDesc"] = request.ItemDesc
	}

	if !dara.IsNil(request.ItemValue) {
		body["itemValue"] = request.ItemValue
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCapability"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/capabilities/" + dara.PercentEncode(dara.StringValue(itemCategory)) + "/items/" + dara.PercentEncode(dara.StringValue(itemName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCapabilityResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a specific configuration item in a specified workspace.
//
// Description:
//
// ## Request description
//
// This API operation allows you to update a specific configuration item (`itemName`) under a category (`itemCategory`) in a specified workspace (`workspaceName`). By setting the `dryRun` parameter, you can preview changes without actually applying them. The request body can contain a new configuration description (`itemDesc`) and configuration content (`itemValue`). The structure of `itemValue` must conform to the requirements of the target configuration item.
//
// @param request - UpdateCapabilityRequest
//
// @return UpdateCapabilityResponse
func (client *Client) UpdateCapability(workspaceName *string, itemCategory *string, itemName *string, request *UpdateCapabilityRequest) (_result *UpdateCapabilityResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateCapabilityResponse{}
	_body, _err := client.UpdateCapabilityWithOptions(workspaceName, itemCategory, itemName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configuration of a specific type in a specified workspace.
//
// Description:
//
// ## Request description.
//
// @param request - UpdateConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateConfigResponse
func (client *Client) UpdateConfigWithOptions(workspaceName *string, configType *string, request *UpdateConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateConfigResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ConfigData) {
		body["configData"] = request.ConfigData
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateConfig"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/configs/" + dara.PercentEncode(dara.StringValue(configType))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateConfigResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration of a specific type in a specified workspace.
//
// Description:
//
// ## Request description.
//
// @param request - UpdateConfigRequest
//
// @return UpdateConfigResponse
func (client *Client) UpdateConfig(workspaceName *string, configType *string, request *UpdateConfigRequest) (_result *UpdateConfigResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateConfigResponse{}
	_body, _err := client.UpdateConfigWithOptions(workspaceName, configType, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies an access credential.
//
// @param request - UpdateCredentialsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCredentialsResponse
func (client *Client) UpdateCredentialsWithOptions(token *string, workspaceName *string, request *UpdateCredentialsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateCredentialsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DryRun) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Enabled) {
		body["enabled"] = request.Enabled
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateCredentials"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/credentials/" + dara.PercentEncode(dara.StringValue(token))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCredentialsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies an access credential.
//
// @param request - UpdateCredentialsRequest
//
// @return UpdateCredentialsResponse
func (client *Client) UpdateCredentials(token *string, workspaceName *string, request *UpdateCredentialsRequest) (_result *UpdateCredentialsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateCredentialsResponse{}
	_body, _err := client.UpdateCredentialsWithOptions(token, workspaceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the service configuration.
//
// @param request - UpdateFunctionInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFunctionInstanceResponse
func (client *Client) UpdateFunctionInstanceWithOptions(workspaceName *string, functionName *string, instanceName *string, request *UpdateFunctionInstanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateFunctionInstanceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CreateParameters) {
		body["createParameters"] = request.CreateParameters
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateFunctionInstance"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName)) + "/functions/" + dara.PercentEncode(dara.StringValue(functionName)) + "/instances/" + dara.PercentEncode(dara.StringValue(instanceName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFunctionInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the service configuration.
//
// @param request - UpdateFunctionInstanceRequest
//
// @return UpdateFunctionInstanceResponse
func (client *Client) UpdateFunctionInstance(workspaceName *string, functionName *string, instanceName *string, request *UpdateFunctionInstanceRequest) (_result *UpdateFunctionInstanceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateFunctionInstanceResponse{}
	_body, _err := client.UpdateFunctionInstanceWithOptions(workspaceName, functionName, instanceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a workspace.
//
// @param request - UpdateWorkspaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkspaceResponse
func (client *Client) UpdateWorkspaceWithOptions(workspaceName *string, request *UpdateWorkspaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateWorkspaceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["resourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateWorkspace"),
		Version:     dara.String("2024-04-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/platform/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWorkspaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a workspace.
//
// @param request - UpdateWorkspaceRequest
//
// @return UpdateWorkspaceResponse
func (client *Client) UpdateWorkspace(workspaceName *string, request *UpdateWorkspaceRequest) (_result *UpdateWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateWorkspaceResponse{}
	_body, _err := client.UpdateWorkspaceWithOptions(workspaceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
