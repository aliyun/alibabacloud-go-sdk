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
	client.Endpoint, _err = client.GetEndpoint(dara.String("agentrun"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// # Create an agent runtime
//
// Description:
//
// 创建一个新的智能体运行时实例，用于执行AI代理任务。智能体运行时是AgentRun服务的核心组件，提供代码执行、浏览器操作、内存管理等能力。
//
// @param request - CreateAgentRuntimeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAgentRuntimeResponse
func (client *Client) CreateAgentRuntimeWithOptions(request *CreateAgentRuntimeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAgentRuntimeResponse, _err error) {
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
		Action:      dara.String("CreateAgentRuntime"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/runtimes"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAgentRuntimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create an agent runtime
//
// Description:
//
// 创建一个新的智能体运行时实例，用于执行AI代理任务。智能体运行时是AgentRun服务的核心组件，提供代码执行、浏览器操作、内存管理等能力。
//
// @param request - CreateAgentRuntimeRequest
//
// @return CreateAgentRuntimeResponse
func (client *Client) CreateAgentRuntime(request *CreateAgentRuntimeRequest) (_result *CreateAgentRuntimeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAgentRuntimeResponse{}
	_body, _err := client.CreateAgentRuntimeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建智能体运行时端点
//
// Description:
//
// 为指定的智能体运行时创建新的端点，用于外部访问和调用。端点是智能体运行时对外提供服务的入口。
//
// @param request - CreateAgentRuntimeEndpointRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAgentRuntimeEndpointResponse
func (client *Client) CreateAgentRuntimeEndpointWithOptions(agentRuntimeId *string, request *CreateAgentRuntimeEndpointRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAgentRuntimeEndpointResponse, _err error) {
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
		Action:      dara.String("CreateAgentRuntimeEndpoint"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/runtimes/" + dara.PercentEncode(dara.StringValue(agentRuntimeId)) + "/endpoints"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAgentRuntimeEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建智能体运行时端点
//
// Description:
//
// 为指定的智能体运行时创建新的端点，用于外部访问和调用。端点是智能体运行时对外提供服务的入口。
//
// @param request - CreateAgentRuntimeEndpointRequest
//
// @return CreateAgentRuntimeEndpointResponse
func (client *Client) CreateAgentRuntimeEndpoint(agentRuntimeId *string, request *CreateAgentRuntimeEndpointRequest) (_result *CreateAgentRuntimeEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAgentRuntimeEndpointResponse{}
	_body, _err := client.CreateAgentRuntimeEndpointWithOptions(agentRuntimeId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建浏览器实例
//
// Description:
//
// 创建一个新的浏览器实例，用于执行网页自动化任务。浏览器实例提供网页浏览、元素操作、截图录制等功能。
//
// @param request - CreateBrowserRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBrowserResponse
func (client *Client) CreateBrowserWithOptions(request *CreateBrowserRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateBrowserResponse, _err error) {
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
		Action:      dara.String("CreateBrowser"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/browsers"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateBrowserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建浏览器实例
//
// Description:
//
// 创建一个新的浏览器实例，用于执行网页自动化任务。浏览器实例提供网页浏览、元素操作、截图录制等功能。
//
// @param request - CreateBrowserRequest
//
// @return CreateBrowserResponse
func (client *Client) CreateBrowser(request *CreateBrowserRequest) (_result *CreateBrowserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateBrowserResponse{}
	_body, _err := client.CreateBrowserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建代码解释器
//
// Description:
//
// 创建一个新的代码解释器实例，用于执行代码解释和运行任务。代码解释器提供Python代码执行、数据处理、机器学习等功能。
//
// @param request - CreateCodeInterpreterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCodeInterpreterResponse
func (client *Client) CreateCodeInterpreterWithOptions(request *CreateCodeInterpreterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateCodeInterpreterResponse, _err error) {
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
		Action:      dara.String("CreateCodeInterpreter"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/code-interpreters"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCodeInterpreterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建代码解释器
//
// Description:
//
// 创建一个新的代码解释器实例，用于执行代码解释和运行任务。代码解释器提供Python代码执行、数据处理、机器学习等功能。
//
// @param request - CreateCodeInterpreterRequest
//
// @return CreateCodeInterpreterResponse
func (client *Client) CreateCodeInterpreter(request *CreateCodeInterpreterRequest) (_result *CreateCodeInterpreterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateCodeInterpreterResponse{}
	_body, _err := client.CreateCodeInterpreterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// create memory store
//
// @param request - CreateMemoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMemoryResponse
func (client *Client) CreateMemoryWithOptions(request *CreateMemoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateMemoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.LongTtl) {
		body["longTtl"] = request.LongTtl
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.ShortTtl) {
		body["shortTtl"] = request.ShortTtl
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMemory"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/memories"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMemoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// create memory store
//
// @param request - CreateMemoryRequest
//
// @return CreateMemoryResponse
func (client *Client) CreateMemory(request *CreateMemoryRequest) (_result *CreateMemoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateMemoryResponse{}
	_body, _err := client.CreateMemoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// create event
//
// @param request - CreateMemoryEventRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMemoryEventResponse
func (client *Client) CreateMemoryEventWithOptions(memoryName *string, request *CreateMemoryEventRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateMemoryEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Events) {
		body["events"] = request.Events
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateMemoryEvent"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/memories/" + dara.PercentEncode(dara.StringValue(memoryName)) + "/events"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMemoryEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// create event
//
// @param request - CreateMemoryEventRequest
//
// @return CreateMemoryEventResponse
func (client *Client) CreateMemoryEvent(memoryName *string, request *CreateMemoryEventRequest) (_result *CreateMemoryEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateMemoryEventResponse{}
	_body, _err := client.CreateMemoryEventWithOptions(memoryName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除智能体运行时
//
// Description:
//
// 删除指定的智能体运行时实例，包括其所有相关资源和数据。删除操作不可逆，请谨慎操作。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAgentRuntimeResponse
func (client *Client) DeleteAgentRuntimeWithOptions(agentRuntimeId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAgentRuntimeResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAgentRuntime"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/runtimes/" + dara.PercentEncode(dara.StringValue(agentRuntimeId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAgentRuntimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除智能体运行时
//
// Description:
//
// 删除指定的智能体运行时实例，包括其所有相关资源和数据。删除操作不可逆，请谨慎操作。
//
// @return DeleteAgentRuntimeResponse
func (client *Client) DeleteAgentRuntime(agentRuntimeId *string) (_result *DeleteAgentRuntimeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAgentRuntimeResponse{}
	_body, _err := client.DeleteAgentRuntimeWithOptions(agentRuntimeId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete an agent runtime endpoint
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAgentRuntimeEndpointResponse
func (client *Client) DeleteAgentRuntimeEndpointWithOptions(agentRuntimeId *string, agentRuntimeEndpointId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAgentRuntimeEndpointResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAgentRuntimeEndpoint"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/runtimes/" + dara.PercentEncode(dara.StringValue(agentRuntimeId)) + "/endpoints/" + dara.PercentEncode(dara.StringValue(agentRuntimeEndpointId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAgentRuntimeEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete an agent runtime endpoint
//
// @return DeleteAgentRuntimeEndpointResponse
func (client *Client) DeleteAgentRuntimeEndpoint(agentRuntimeId *string, agentRuntimeEndpointId *string) (_result *DeleteAgentRuntimeEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAgentRuntimeEndpointResponse{}
	_body, _err := client.DeleteAgentRuntimeEndpointWithOptions(agentRuntimeId, agentRuntimeEndpointId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除浏览器实例
//
// Description:
//
// 删除指定的浏览器实例，包括其所有相关资源和数据。删除操作不可逆，请谨慎操作。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBrowserResponse
func (client *Client) DeleteBrowserWithOptions(browserId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteBrowserResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteBrowser"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/browsers/" + dara.PercentEncode(dara.StringValue(browserId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteBrowserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除浏览器实例
//
// Description:
//
// 删除指定的浏览器实例，包括其所有相关资源和数据。删除操作不可逆，请谨慎操作。
//
// @return DeleteBrowserResponse
func (client *Client) DeleteBrowser(browserId *string) (_result *DeleteBrowserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteBrowserResponse{}
	_body, _err := client.DeleteBrowserWithOptions(browserId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除代码解释器
//
// Description:
//
// 删除指定的代码解释器实例，包括其所有相关资源和数据。删除操作不可逆，请谨慎操作。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCodeInterpreterResponse
func (client *Client) DeleteCodeInterpreterWithOptions(codeInterpreterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteCodeInterpreterResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCodeInterpreter"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/code-interpreters/" + dara.PercentEncode(dara.StringValue(codeInterpreterId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCodeInterpreterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除代码解释器
//
// Description:
//
// 删除指定的代码解释器实例，包括其所有相关资源和数据。删除操作不可逆，请谨慎操作。
//
// @return DeleteCodeInterpreterResponse
func (client *Client) DeleteCodeInterpreter(codeInterpreterId *string) (_result *DeleteCodeInterpreterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteCodeInterpreterResponse{}
	_body, _err := client.DeleteCodeInterpreterWithOptions(codeInterpreterId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// delete memory store
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMemoryResponse
func (client *Client) DeleteMemoryWithOptions(memoryName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteMemoryResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMemory"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/memories/" + dara.PercentEncode(dara.StringValue(memoryName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMemoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// delete memory store
//
// @return DeleteMemoryResponse
func (client *Client) DeleteMemory(memoryName *string) (_result *DeleteMemoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteMemoryResponse{}
	_body, _err := client.DeleteMemoryWithOptions(memoryName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取智能体运行时详情
//
// Description:
//
// 根据智能体运行时ID获取指定智能体运行时的详细信息，包括配置、状态、资源使用情况等。
//
// @param request - GetAgentRuntimeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentRuntimeResponse
func (client *Client) GetAgentRuntimeWithOptions(agentRuntimeId *string, request *GetAgentRuntimeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAgentRuntimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentRuntimeVersion) {
		query["agentRuntimeVersion"] = request.AgentRuntimeVersion
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAgentRuntime"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/runtimes/" + dara.PercentEncode(dara.StringValue(agentRuntimeId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgentRuntimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取智能体运行时详情
//
// Description:
//
// 根据智能体运行时ID获取指定智能体运行时的详细信息，包括配置、状态、资源使用情况等。
//
// @param request - GetAgentRuntimeRequest
//
// @return GetAgentRuntimeResponse
func (client *Client) GetAgentRuntime(agentRuntimeId *string, request *GetAgentRuntimeRequest) (_result *GetAgentRuntimeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAgentRuntimeResponse{}
	_body, _err := client.GetAgentRuntimeWithOptions(agentRuntimeId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get an agent runtime endpoint
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentRuntimeEndpointResponse
func (client *Client) GetAgentRuntimeEndpointWithOptions(agentRuntimeId *string, agentRuntimeEndpointId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAgentRuntimeEndpointResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAgentRuntimeEndpoint"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/runtimes/" + dara.PercentEncode(dara.StringValue(agentRuntimeId)) + "/endpoints/" + dara.PercentEncode(dara.StringValue(agentRuntimeEndpointId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAgentRuntimeEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get an agent runtime endpoint
//
// @return GetAgentRuntimeEndpointResponse
func (client *Client) GetAgentRuntimeEndpoint(agentRuntimeId *string, agentRuntimeEndpointId *string) (_result *GetAgentRuntimeEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAgentRuntimeEndpointResponse{}
	_body, _err := client.GetAgentRuntimeEndpointWithOptions(agentRuntimeId, agentRuntimeEndpointId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取浏览器实例详情
//
// Description:
//
// 根据浏览器ID获取指定浏览器实例的详细信息，包括配置、状态、资源使用情况等。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBrowserResponse
func (client *Client) GetBrowserWithOptions(browserId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetBrowserResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetBrowser"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/browsers/" + dara.PercentEncode(dara.StringValue(browserId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetBrowserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取浏览器实例详情
//
// Description:
//
// 根据浏览器ID获取指定浏览器实例的详细信息，包括配置、状态、资源使用情况等。
//
// @return GetBrowserResponse
func (client *Client) GetBrowser(browserId *string) (_result *GetBrowserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetBrowserResponse{}
	_body, _err := client.GetBrowserWithOptions(browserId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取代码解释器详情
//
// Description:
//
// 根据代码解释器ID获取指定代码解释器实例的详细信息，包括配置、状态、资源使用情况等。
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCodeInterpreterResponse
func (client *Client) GetCodeInterpreterWithOptions(codeInterpreterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCodeInterpreterResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCodeInterpreter"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/code-interpreters/" + dara.PercentEncode(dara.StringValue(codeInterpreterId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCodeInterpreterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取代码解释器详情
//
// Description:
//
// 根据代码解释器ID获取指定代码解释器实例的详细信息，包括配置、状态、资源使用情况等。
//
// @return GetCodeInterpreterResponse
func (client *Client) GetCodeInterpreter(codeInterpreterId *string) (_result *GetCodeInterpreterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCodeInterpreterResponse{}
	_body, _err := client.GetCodeInterpreterWithOptions(codeInterpreterId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # GetMemory
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMemoryResponse
func (client *Client) GetMemoryWithOptions(memoryName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMemoryResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMemory"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/memories/" + dara.PercentEncode(dara.StringValue(memoryName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMemoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # GetMemory
//
// @return GetMemoryResponse
func (client *Client) GetMemory(memoryName *string) (_result *GetMemoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMemoryResponse{}
	_body, _err := client.GetMemoryWithOptions(memoryName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// get event
//
// @param request - GetMemoryEventRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMemoryEventResponse
func (client *Client) GetMemoryEventWithOptions(memoryName *string, sessionId *string, eventId *string, request *GetMemoryEventRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMemoryEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.From) {
		query["from"] = request.From
	}

	if !dara.IsNil(request.To) {
		query["to"] = request.To
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMemoryEvent"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/memories/" + dara.PercentEncode(dara.StringValue(memoryName)) + "/sessions/" + dara.PercentEncode(dara.StringValue(sessionId)) + "/events/" + dara.PercentEncode(dara.StringValue(eventId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMemoryEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// get event
//
// @param request - GetMemoryEventRequest
//
// @return GetMemoryEventResponse
func (client *Client) GetMemoryEvent(memoryName *string, sessionId *string, eventId *string, request *GetMemoryEventRequest) (_result *GetMemoryEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMemoryEventResponse{}
	_body, _err := client.GetMemoryEventWithOptions(memoryName, sessionId, eventId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取内存会话详情
//
// Description:
//
// 根据会话ID获取指定内存会话的详细信息，包括会话中的事件记录、时间戳等。用于查看和管理对话历史。
//
// @param request - GetMemorySessionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMemorySessionResponse
func (client *Client) GetMemorySessionWithOptions(memoryName *string, sessionId *string, request *GetMemorySessionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMemorySessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.From) {
		query["from"] = request.From
	}

	if !dara.IsNil(request.Size) {
		query["size"] = request.Size
	}

	if !dara.IsNil(request.To) {
		query["to"] = request.To
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMemorySession"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/memories/" + dara.PercentEncode(dara.StringValue(memoryName)) + "/sessions/" + dara.PercentEncode(dara.StringValue(sessionId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMemorySessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取内存会话详情
//
// Description:
//
// 根据会话ID获取指定内存会话的详细信息，包括会话中的事件记录、时间戳等。用于查看和管理对话历史。
//
// @param request - GetMemorySessionRequest
//
// @return GetMemorySessionResponse
func (client *Client) GetMemorySession(memoryName *string, sessionId *string, request *GetMemorySessionRequest) (_result *GetMemorySessionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMemorySessionResponse{}
	_body, _err := client.GetMemorySessionWithOptions(memoryName, sessionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出智能体运行时端点
//
// Description:
//
// 获取指定智能体运行时的所有端点列表，支持按名称过滤和分页查询。端点用于外部系统访问智能体运行时服务。
//
// @param request - ListAgentRuntimeEndpointsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAgentRuntimeEndpointsResponse
func (client *Client) ListAgentRuntimeEndpointsWithOptions(agentRuntimeId *string, request *ListAgentRuntimeEndpointsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAgentRuntimeEndpointsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndpointName) {
		query["endpointName"] = request.EndpointName
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
		Action:      dara.String("ListAgentRuntimeEndpoints"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/runtimes/" + dara.PercentEncode(dara.StringValue(agentRuntimeId)) + "/endpoints"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAgentRuntimeEndpointsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出智能体运行时端点
//
// Description:
//
// 获取指定智能体运行时的所有端点列表，支持按名称过滤和分页查询。端点用于外部系统访问智能体运行时服务。
//
// @param request - ListAgentRuntimeEndpointsRequest
//
// @return ListAgentRuntimeEndpointsResponse
func (client *Client) ListAgentRuntimeEndpoints(agentRuntimeId *string, request *ListAgentRuntimeEndpointsRequest) (_result *ListAgentRuntimeEndpointsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAgentRuntimeEndpointsResponse{}
	_body, _err := client.ListAgentRuntimeEndpointsWithOptions(agentRuntimeId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # List agent runtime versions
//
// @param request - ListAgentRuntimeVersionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAgentRuntimeVersionsResponse
func (client *Client) ListAgentRuntimeVersionsWithOptions(agentRuntimeId *string, request *ListAgentRuntimeVersionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAgentRuntimeVersionsResponse, _err error) {
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
		Action:      dara.String("ListAgentRuntimeVersions"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/runtimes/" + dara.PercentEncode(dara.StringValue(agentRuntimeId)) + "/versions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAgentRuntimeVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # List agent runtime versions
//
// @param request - ListAgentRuntimeVersionsRequest
//
// @return ListAgentRuntimeVersionsResponse
func (client *Client) ListAgentRuntimeVersions(agentRuntimeId *string, request *ListAgentRuntimeVersionsRequest) (_result *ListAgentRuntimeVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAgentRuntimeVersionsResponse{}
	_body, _err := client.ListAgentRuntimeVersionsWithOptions(agentRuntimeId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出智能体运行时
//
// Description:
//
// 获取当前用户的所有智能体运行时列表，支持按名称、标签等条件过滤，支持分页查询。
//
// @param request - ListAgentRuntimesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAgentRuntimesResponse
func (client *Client) ListAgentRuntimesWithOptions(request *ListAgentRuntimesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAgentRuntimesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentRuntimeName) {
		query["agentRuntimeName"] = request.AgentRuntimeName
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
		Action:      dara.String("ListAgentRuntimes"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/runtimes"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAgentRuntimesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出智能体运行时
//
// Description:
//
// 获取当前用户的所有智能体运行时列表，支持按名称、标签等条件过滤，支持分页查询。
//
// @param request - ListAgentRuntimesRequest
//
// @return ListAgentRuntimesResponse
func (client *Client) ListAgentRuntimes(request *ListAgentRuntimesRequest) (_result *ListAgentRuntimesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAgentRuntimesResponse{}
	_body, _err := client.ListAgentRuntimesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出浏览器实例
//
// Description:
//
// 获取当前用户的所有浏览器实例列表，支持按名称、状态等条件过滤，支持分页查询。
//
// @param request - ListBrowsersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBrowsersResponse
func (client *Client) ListBrowsersWithOptions(request *ListBrowsersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListBrowsersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BrowserName) {
		query["browserName"] = request.BrowserName
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListBrowsers"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/browsers"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListBrowsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出浏览器实例
//
// Description:
//
// 获取当前用户的所有浏览器实例列表，支持按名称、状态等条件过滤，支持分页查询。
//
// @param request - ListBrowsersRequest
//
// @return ListBrowsersResponse
func (client *Client) ListBrowsers(request *ListBrowsersRequest) (_result *ListBrowsersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListBrowsersResponse{}
	_body, _err := client.ListBrowsersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出代码解释器
//
// Description:
//
// 获取当前用户的所有代码解释器实例列表，支持按名称等条件过滤，支持分页查询。
//
// @param request - ListCodeInterpretersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCodeInterpretersResponse
func (client *Client) ListCodeInterpretersWithOptions(request *ListCodeInterpretersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListCodeInterpretersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CodeInterpreterName) {
		query["codeInterpreterName"] = request.CodeInterpreterName
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
		Action:      dara.String("ListCodeInterpreters"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/code-interpreters"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCodeInterpretersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出代码解释器
//
// Description:
//
// 获取当前用户的所有代码解释器实例列表，支持按名称等条件过滤，支持分页查询。
//
// @param request - ListCodeInterpretersRequest
//
// @return ListCodeInterpretersResponse
func (client *Client) ListCodeInterpreters(request *ListCodeInterpretersRequest) (_result *ListCodeInterpretersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListCodeInterpretersResponse{}
	_body, _err := client.ListCodeInterpretersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # ListMemory
//
// @param request - ListMemoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMemoryResponse
func (client *Client) ListMemoryWithOptions(request *ListMemoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMemoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NamePrefix) {
		query["namePrefix"] = request.NamePrefix
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
		Action:      dara.String("ListMemory"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/memories"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMemoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ListMemory
//
// @param request - ListMemoryRequest
//
// @return ListMemoryResponse
func (client *Client) ListMemory(request *ListMemoryRequest) (_result *ListMemoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMemoryResponse{}
	_body, _err := client.ListMemoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// list events
//
// @param request - ListMemoryEventRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMemoryEventResponse
func (client *Client) ListMemoryEventWithOptions(memoryName *string, sessionId *string, request *ListMemoryEventRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMemoryEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.From) {
		query["from"] = request.From
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.To) {
		query["to"] = request.To
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMemoryEvent"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/memories/" + dara.PercentEncode(dara.StringValue(memoryName)) + "/sessions/" + dara.PercentEncode(dara.StringValue(sessionId)) + "/events"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMemoryEventResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// list events
//
// @param request - ListMemoryEventRequest
//
// @return ListMemoryEventResponse
func (client *Client) ListMemoryEvent(memoryName *string, sessionId *string, request *ListMemoryEventRequest) (_result *ListMemoryEventResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMemoryEventResponse{}
	_body, _err := client.ListMemoryEventWithOptions(memoryName, sessionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列出内存会话
//
// Description:
//
// 获取指定内存实例的所有会话列表，支持按时间范围过滤和分页查询。会话是AgentRun中用于存储对话历史和管理上下文的重要组件。
//
// @param request - ListMemorySessionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMemorySessionsResponse
func (client *Client) ListMemorySessionsWithOptions(memoryName *string, request *ListMemorySessionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMemorySessionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.From) {
		query["from"] = request.From
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.To) {
		query["to"] = request.To
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMemorySessions"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/memories/" + dara.PercentEncode(dara.StringValue(memoryName)) + "/sessions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMemorySessionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列出内存会话
//
// Description:
//
// 获取指定内存实例的所有会话列表，支持按时间范围过滤和分页查询。会话是AgentRun中用于存储对话历史和管理上下文的重要组件。
//
// @param request - ListMemorySessionsRequest
//
// @return ListMemorySessionsResponse
func (client *Client) ListMemorySessions(memoryName *string, request *ListMemorySessionsRequest) (_result *ListMemorySessionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMemorySessionsResponse{}
	_body, _err := client.ListMemorySessionsWithOptions(memoryName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 发布运行时版本
//
// Description:
//
// 为指定的智能体运行时发布新版本，用于版本管理和部署。新版本可以包含代码更新、配置变更等内容。
//
// @param request - PublishRuntimeVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishRuntimeVersionResponse
func (client *Client) PublishRuntimeVersionWithOptions(agentRuntimeId *string, request *PublishRuntimeVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PublishRuntimeVersionResponse, _err error) {
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
		Action:      dara.String("PublishRuntimeVersion"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/runtimes/" + dara.PercentEncode(dara.StringValue(agentRuntimeId)) + "/versions"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishRuntimeVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发布运行时版本
//
// Description:
//
// 为指定的智能体运行时发布新版本，用于版本管理和部署。新版本可以包含代码更新、配置变更等内容。
//
// @param request - PublishRuntimeVersionRequest
//
// @return PublishRuntimeVersionResponse
func (client *Client) PublishRuntimeVersion(agentRuntimeId *string, request *PublishRuntimeVersionRequest) (_result *PublishRuntimeVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PublishRuntimeVersionResponse{}
	_body, _err := client.PublishRuntimeVersionWithOptions(agentRuntimeId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # RetrieveMemory
//
// @param request - RetrieveMemoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetrieveMemoryResponse
func (client *Client) RetrieveMemoryWithOptions(memoryName *string, request *RetrieveMemoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RetrieveMemoryResponse, _err error) {
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

	if !dara.IsNil(request.Query) {
		body["query"] = request.Query
	}

	if !dara.IsNil(request.Store) {
		body["store"] = request.Store
	}

	if !dara.IsNil(request.To) {
		body["to"] = request.To
	}

	if !dara.IsNil(request.Topk) {
		body["topk"] = request.Topk
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RetrieveMemory"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/memories/" + dara.PercentEncode(dara.StringValue(memoryName)) + "/records"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RetrieveMemoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # RetrieveMemory
//
// @param request - RetrieveMemoryRequest
//
// @return RetrieveMemoryResponse
func (client *Client) RetrieveMemory(memoryName *string, request *RetrieveMemoryRequest) (_result *RetrieveMemoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RetrieveMemoryResponse{}
	_body, _err := client.RetrieveMemoryWithOptions(memoryName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新智能体运行时
//
// Description:
//
// 更新指定智能体运行时的配置信息，包括资源分配、网络配置、环境变量等。更新操作会触发运行时重启。
//
// @param request - UpdateAgentRuntimeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAgentRuntimeResponse
func (client *Client) UpdateAgentRuntimeWithOptions(agentRuntimeId *string, request *UpdateAgentRuntimeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAgentRuntimeResponse, _err error) {
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
		Action:      dara.String("UpdateAgentRuntime"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/runtimes/" + dara.PercentEncode(dara.StringValue(agentRuntimeId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAgentRuntimeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新智能体运行时
//
// Description:
//
// 更新指定智能体运行时的配置信息，包括资源分配、网络配置、环境变量等。更新操作会触发运行时重启。
//
// @param request - UpdateAgentRuntimeRequest
//
// @return UpdateAgentRuntimeResponse
func (client *Client) UpdateAgentRuntime(agentRuntimeId *string, request *UpdateAgentRuntimeRequest) (_result *UpdateAgentRuntimeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAgentRuntimeResponse{}
	_body, _err := client.UpdateAgentRuntimeWithOptions(agentRuntimeId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Update an agent runtime endpoint
//
// @param request - UpdateAgentRuntimeEndpointRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAgentRuntimeEndpointResponse
func (client *Client) UpdateAgentRuntimeEndpointWithOptions(agentRuntimeId *string, agentRuntimeEndpointId *string, request *UpdateAgentRuntimeEndpointRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAgentRuntimeEndpointResponse, _err error) {
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
		Action:      dara.String("UpdateAgentRuntimeEndpoint"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/runtimes/" + dara.PercentEncode(dara.StringValue(agentRuntimeId)) + "/endpoints/" + dara.PercentEncode(dara.StringValue(agentRuntimeEndpointId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAgentRuntimeEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update an agent runtime endpoint
//
// @param request - UpdateAgentRuntimeEndpointRequest
//
// @return UpdateAgentRuntimeEndpointResponse
func (client *Client) UpdateAgentRuntimeEndpoint(agentRuntimeId *string, agentRuntimeEndpointId *string, request *UpdateAgentRuntimeEndpointRequest) (_result *UpdateAgentRuntimeEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAgentRuntimeEndpointResponse{}
	_body, _err := client.UpdateAgentRuntimeEndpointWithOptions(agentRuntimeId, agentRuntimeEndpointId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Update Memory
//
// @param request - UpdateMemoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMemoryResponse
func (client *Client) UpdateMemoryWithOptions(memoryName *string, request *UpdateMemoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateMemoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LongTtl) {
		query["longTtl"] = request.LongTtl
	}

	if !dara.IsNil(request.ShortTtl) {
		query["shortTtl"] = request.ShortTtl
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateMemory"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/memories/" + dara.PercentEncode(dara.StringValue(memoryName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMemoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update Memory
//
// @param request - UpdateMemoryRequest
//
// @return UpdateMemoryResponse
func (client *Client) UpdateMemory(memoryName *string, request *UpdateMemoryRequest) (_result *UpdateMemoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateMemoryResponse{}
	_body, _err := client.UpdateMemoryWithOptions(memoryName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
