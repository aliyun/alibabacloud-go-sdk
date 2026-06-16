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
		"cn-shenzhen":    dara.String("agentrun.cn-shenzhen.aliyuncs.com"),
		"cn-shanghai":    dara.String("agentrun.cn-shanghai.aliyuncs.com"),
		"cn-hangzhou":    dara.String("agentrun.cn-hangzhou.aliyuncs.com"),
		"cn-beijing":     dara.String("agentrun.cn-beijing.aliyuncs.com"),
		"ap-southeast-1": dara.String("agentrun.ap-southeast-1.aliyuncs.com"),
	}
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
// Activates the `MCP service` for a `sandbox` `template`. This enables a client to access the `sandbox` using the MCP protocol.
//
// Description:
//
// After activation, the platform automatically deploys the `MCP service` `function` for the specified `sandbox` `template`. The `MCP service` ensures a unique mapping between an `mcp-session-id` and a `SandboxID`. When an MCP `client` invokes a `tool`, the `MCP service` automatically creates a `sandbox`.
//
// @param request - ActivateTemplateMCPRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActivateTemplateMCPResponse
func (client *Client) ActivateTemplateMCPWithOptions(templateName *string, request *ActivateTemplateMCPRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ActivateTemplateMCPResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.EnabledTools) {
		body["enabledTools"] = request.EnabledTools
	}

	if !dara.IsNil(request.Transport) {
		body["transport"] = request.Transport
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ActivateTemplateMCP"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/templates/" + dara.PercentEncode(dara.StringValue(templateName)) + "/mcp/activate"),
		Method:      dara.String("PATCH"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ActivateTemplateMCPResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Activates the `MCP service` for a `sandbox` `template`. This enables a client to access the `sandbox` using the MCP protocol.
//
// Description:
//
// After activation, the platform automatically deploys the `MCP service` `function` for the specified `sandbox` `template`. The `MCP service` ensures a unique mapping between an `mcp-session-id` and a `SandboxID`. When an MCP `client` invokes a `tool`, the `MCP service` automatically creates a `sandbox`.
//
// @param request - ActivateTemplateMCPRequest
//
// @return ActivateTemplateMCPResponse
func (client *Client) ActivateTemplateMCP(templateName *string, request *ActivateTemplateMCPRequest) (_result *ActivateTemplateMCPResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ActivateTemplateMCPResponse{}
	_body, _err := client.ActivateTemplateMCPWithOptions(templateName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Converts a Flow DSL.
//
// Description:
//
// This operation converts a third-party workflow DSL, such as Dify or n8n, into an AgentRun Flow definition. It performs compatibility checks, identifies plugins, and extracts metadata. The operation runs in dry-run mode, returning the converted Flow configuration, a compatibility analysis report, and the required Toolset installation configuration without creating a Flow resource.
//
// @param request - ConvertFlowDSLRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ConvertFlowDSLResponse
func (client *Client) ConvertFlowDSLWithOptions(request *ConvertFlowDSLRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ConvertFlowDSLResponse, _err error) {
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
		Action:      dara.String("ConvertFlowDSL"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/flows/action/convertDsl"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ConvertFlowDSLResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Converts a Flow DSL.
//
// Description:
//
// This operation converts a third-party workflow DSL, such as Dify or n8n, into an AgentRun Flow definition. It performs compatibility checks, identifies plugins, and extracts metadata. The operation runs in dry-run mode, returning the converted Flow configuration, a compatibility analysis report, and the required Toolset installation configuration without creating a Flow resource.
//
// @param request - ConvertFlowDSLRequest
//
// @return ConvertFlowDSLResponse
func (client *Client) ConvertFlowDSL(request *ConvertFlowDSLRequest) (_result *ConvertFlowDSLResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ConvertFlowDSLResponse{}
	_body, _err := client.ConvertFlowDSLWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an agent runtime.
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
// Creates an agent runtime.
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
// # Create an access endpoint for an agent runtime
//
// Description:
//
// Creates a new endpoint for the specified agent runtime, used for external access and invocation. An endpoint serves as the entry point through which an agent runtime provides services externally.
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
// # Create an access endpoint for an agent runtime
//
// Description:
//
// Creates a new endpoint for the specified agent runtime, used for external access and invocation. An endpoint serves as the entry point through which an agent runtime provides services externally.
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
// # Create Browser Sandbox
//
// Description:
//
// Create a new browser instance for executing web automation tasks. The browser instance provides features such as web browsing, element manipulation, and screenshot recording.
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
// # Create Browser Sandbox
//
// Description:
//
// Create a new browser instance for executing web automation tasks. The browser instance provides features such as web browsing, element manipulation, and screenshot recording.
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
// Creates a code interpreter.
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
// Creates a code interpreter.
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
// Creates a new credential.
//
// Description:
//
// This operation creates a credential for an agent.
//
// @param request - CreateCredentialRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCredentialResponse
func (client *Client) CreateCredentialWithOptions(request *CreateCredentialRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateCredentialResponse, _err error) {
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
		Action:      dara.String("CreateCredential"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/credentials"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCredentialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a new credential.
//
// Description:
//
// This operation creates a credential for an agent.
//
// @param request - CreateCredentialRequest
//
// @return CreateCredentialResponse
func (client *Client) CreateCredential(request *CreateCredentialRequest) (_result *CreateCredentialResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateCredentialResponse{}
	_body, _err := client.CreateCredentialWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a custom domain.
//
// @param request - CreateCustomDomainRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomDomainResponse
func (client *Client) CreateCustomDomainWithOptions(request *CreateCustomDomainRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateCustomDomainResponse, _err error) {
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
		Action:      dara.String("CreateCustomDomain"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/custom-domains"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateCustomDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a custom domain.
//
// @param request - CreateCustomDomainRequest
//
// @return CreateCustomDomainResponse
func (client *Client) CreateCustomDomain(request *CreateCustomDomainRequest) (_result *CreateCustomDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateCustomDomainResponse{}
	_body, _err := client.CreateCustomDomainWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a workflow.
//
// Description:
//
// Creates a flow orchestration agent. Flow orchestration is a core component of the AgentRun service that supports visual orchestration and version management.
//
// @param request - CreateFlowRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFlowResponse
func (client *Client) CreateFlowWithOptions(request *CreateFlowRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateFlowResponse, _err error) {
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
		Action:      dara.String("CreateFlow"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/flows"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a workflow.
//
// Description:
//
// Creates a flow orchestration agent. Flow orchestration is a core component of the AgentRun service that supports visual orchestration and version management.
//
// @param request - CreateFlowRequest
//
// @return CreateFlowResponse
func (client *Client) CreateFlow(request *CreateFlowRequest) (_result *CreateFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateFlowResponse{}
	_body, _err := client.CreateFlowWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a flow endpoint.
//
// Description:
//
// 为指定工作流创建一个新的端点，用于对外提供服务访问。
//
// @param request - CreateFlowEndpointRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFlowEndpointResponse
func (client *Client) CreateFlowEndpointWithOptions(flowName *string, request *CreateFlowEndpointRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateFlowEndpointResponse, _err error) {
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
		Action:      dara.String("CreateFlowEndpoint"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/flows/" + dara.PercentEncode(dara.StringValue(flowName)) + "/endpoints"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFlowEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a flow endpoint.
//
// Description:
//
// 为指定工作流创建一个新的端点，用于对外提供服务访问。
//
// @param request - CreateFlowEndpointRequest
//
// @return CreateFlowEndpointResponse
func (client *Client) CreateFlowEndpoint(flowName *string, request *CreateFlowEndpointRequest) (_result *CreateFlowEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateFlowEndpointResponse{}
	_body, _err := client.CreateFlowEndpointWithOptions(flowName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates an IM Bot.
//
// Description:
//
// A successful request returns an HTTP 201 status code. Once created, an IM Bot\\"s status is always `running`. The response is in a standard format, and its `data` field contains an `IMBotInfo` object.
//
// @param request - CreateIMBotRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIMBotResponse
func (client *Client) CreateIMBotWithOptions(request *CreateIMBotRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateIMBotResponse, _err error) {
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
		Action:      dara.String("CreateIMBot"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/im-bots"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIMBotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an IM Bot.
//
// Description:
//
// A successful request returns an HTTP 201 status code. Once created, an IM Bot\\"s status is always `running`. The response is in a standard format, and its `data` field contains an `IMBotInfo` object.
//
// @param request - CreateIMBotRequest
//
// @return CreateIMBotResponse
func (client *Client) CreateIMBot(request *CreateIMBotRequest) (_result *CreateIMBotResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateIMBotResponse{}
	_body, _err := client.CreateIMBotWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a knowledge base.
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
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateKnowledgeBase"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/knowledgebases"),
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
// Creates a knowledge base.
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
// Creates a memory collection.
//
// @param request - CreateMemoryCollectionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMemoryCollectionResponse
func (client *Client) CreateMemoryCollectionWithOptions(request *CreateMemoryCollectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateMemoryCollectionResponse, _err error) {
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
		Action:      dara.String("CreateMemoryCollection"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/memory-collections"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateMemoryCollectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a memory collection.
//
// @param request - CreateMemoryCollectionRequest
//
// @return CreateMemoryCollectionResponse
func (client *Client) CreateMemoryCollection(request *CreateMemoryCollectionRequest) (_result *CreateMemoryCollectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateMemoryCollectionResponse{}
	_body, _err := client.CreateMemoryCollectionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a Model Proxy.
//
// Description:
//
// This operation creates a Model Proxy based on the specified configuration.
//
// @param request - CreateModelProxyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateModelProxyResponse
func (client *Client) CreateModelProxyWithOptions(request *CreateModelProxyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateModelProxyResponse, _err error) {
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
		Action:      dara.String("CreateModelProxy"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/model-proxies"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateModelProxyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a Model Proxy.
//
// Description:
//
// This operation creates a Model Proxy based on the specified configuration.
//
// @param request - CreateModelProxyRequest
//
// @return CreateModelProxyResponse
func (client *Client) CreateModelProxy(request *CreateModelProxyRequest) (_result *CreateModelProxyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateModelProxyResponse{}
	_body, _err := client.CreateModelProxyWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a model service.
//
// Description:
//
// This operation creates a model service, such as a code interpreter, based on the specified configuration.
//
// @param request - CreateModelServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateModelServiceResponse
func (client *Client) CreateModelServiceWithOptions(request *CreateModelServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateModelServiceResponse, _err error) {
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
		Action:      dara.String("CreateModelService"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/model-services"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateModelServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a model service.
//
// Description:
//
// This operation creates a model service, such as a code interpreter, based on the specified configuration.
//
// @param request - CreateModelServiceRequest
//
// @return CreateModelServiceResponse
func (client *Client) CreateModelService(request *CreateModelServiceRequest) (_result *CreateModelServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateModelServiceResponse{}
	_body, _err := client.CreateModelServiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a sandbox.
//
// Description:
//
// Creates a new sandbox instance from a specified template. A sandbox provides an isolated execution environment to run code or launch a browser.
//
// @param request - CreateSandboxRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSandboxResponse
func (client *Client) CreateSandboxWithOptions(request *CreateSandboxRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateSandboxResponse, _err error) {
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
		Action:      dara.String("CreateSandbox"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/sandboxes"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateSandboxResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a sandbox.
//
// Description:
//
// Creates a new sandbox instance from a specified template. A sandbox provides an isolated execution environment to run code or launch a browser.
//
// @param request - CreateSandboxRequest
//
// @return CreateSandboxResponse
func (client *Client) CreateSandbox(request *CreateSandboxRequest) (_result *CreateSandboxResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSandboxResponse{}
	_body, _err := client.CreateSandboxWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a template.
//
// Description:
//
// Creates a template for launching sandboxes. A template defines the runtime environment, resource configuration, and other settings for a sandbox.
//
// @param request - CreateTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTemplateResponse
func (client *Client) CreateTemplateWithOptions(request *CreateTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTemplateResponse, _err error) {
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
		Action:      dara.String("CreateTemplate"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/templates"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a template.
//
// Description:
//
// Creates a template for launching sandboxes. A template defines the runtime environment, resource configuration, and other settings for a sandbox.
//
// @param request - CreateTemplateRequest
//
// @return CreateTemplateResponse
func (client *Client) CreateTemplate(request *CreateTemplateRequest) (_result *CreateTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTemplateResponse{}
	_body, _err := client.CreateTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a tool.
//
// Description:
//
// This operation creates various types of tools, such as MCP, function call, and skill. An Agent can then call a tool to extend its capabilities.
//
// @param request - CreateToolRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateToolResponse
func (client *Client) CreateToolWithOptions(request *CreateToolRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateToolResponse, _err error) {
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
		Action:      dara.String("CreateTool"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/tools"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateToolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a tool.
//
// Description:
//
// This operation creates various types of tools, such as MCP, function call, and skill. An Agent can then call a tool to extend its capabilities.
//
// @param request - CreateToolRequest
//
// @return CreateToolResponse
func (client *Client) CreateTool(request *CreateToolRequest) (_result *CreateToolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateToolResponse{}
	_body, _err := client.CreateToolWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a workspace.
//
// Description:
//
// Creates a workspace.
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
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateWorkspace"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/workspaces"),
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
// Creates a workspace.
//
// Description:
//
// Creates a workspace.
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
// # Delete Agent Runtime
//
// Description:
//
// Deletes a specified agent runtime instance, including all associated resources and data. This operation is irreversible. Proceed with caution.
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
// # Delete Agent Runtime
//
// Description:
//
// Deletes a specified agent runtime instance, including all associated resources and data. This operation is irreversible. Proceed with caution.
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
// # Delete Agent Runtime Endpoint
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
// # Delete Agent Runtime Endpoint
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
// # Delete Browser Sandbox
//
// Description:
//
// Delete the specified browser instance, including all its associated resources and data. The deletion is irreversible. Please proceed with caution.
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
// # Delete Browser Sandbox
//
// Description:
//
// Delete the specified browser instance, including all its associated resources and data. The deletion is irreversible. Please proceed with caution.
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
// # Delete Code Interpreter
//
// Description:
//
// Delete a specified code interpreter instance, including all its associated resources and data. This operation is irreversible. Please proceed with caution.
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
// # Delete Code Interpreter
//
// Description:
//
// Delete a specified code interpreter instance, including all its associated resources and data. This operation is irreversible. Please proceed with caution.
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
// Deletes the specified credential.
//
// Description:
//
// This operation deletes the specified credential. This action cannot be undone.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCredentialResponse
func (client *Client) DeleteCredentialWithOptions(credentialName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteCredentialResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCredential"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/credentials/" + dara.PercentEncode(dara.StringValue(credentialName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCredentialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes the specified credential.
//
// Description:
//
// This operation deletes the specified credential. This action cannot be undone.
//
// @return DeleteCredentialResponse
func (client *Client) DeleteCredential(credentialName *string) (_result *DeleteCredentialResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteCredentialResponse{}
	_body, _err := client.DeleteCredentialWithOptions(credentialName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a custom domain.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomDomainResponse
func (client *Client) DeleteCustomDomainWithOptions(domainName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteCustomDomainResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCustomDomain"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/custom-domains/" + dara.PercentEncode(dara.StringValue(domainName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteCustomDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a custom domain.
//
// @return DeleteCustomDomainResponse
func (client *Client) DeleteCustomDomain(domainName *string) (_result *DeleteCustomDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteCustomDomainResponse{}
	_body, _err := client.DeleteCustomDomainWithOptions(domainName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a flow.
//
// Description:
//
// Deletes a specified flow instance, along with all its related resources and data. This operation is irreversible and cannot be undone. Use with caution.
//
// @param request - DeleteFlowRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFlowResponse
func (client *Client) DeleteFlowWithOptions(flowName *string, request *DeleteFlowRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteFlowResponse, _err error) {
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
		Action:      dara.String("DeleteFlow"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/flows/" + dara.PercentEncode(dara.StringValue(flowName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a flow.
//
// Description:
//
// Deletes a specified flow instance, along with all its related resources and data. This operation is irreversible and cannot be undone. Use with caution.
//
// @param request - DeleteFlowRequest
//
// @return DeleteFlowResponse
func (client *Client) DeleteFlow(flowName *string, request *DeleteFlowRequest) (_result *DeleteFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteFlowResponse{}
	_body, _err := client.DeleteFlowWithOptions(flowName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a flow endpoint.
//
// Description:
//
// Deletes the specified flow endpoint. This operation is irreversible. Proceed with caution.
//
// @param request - DeleteFlowEndpointRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFlowEndpointResponse
func (client *Client) DeleteFlowEndpointWithOptions(flowName *string, flowEndpointName *string, request *DeleteFlowEndpointRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteFlowEndpointResponse, _err error) {
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
		Action:      dara.String("DeleteFlowEndpoint"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/flows/" + dara.PercentEncode(dara.StringValue(flowName)) + "/endpoints/" + dara.PercentEncode(dara.StringValue(flowEndpointName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFlowEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a flow endpoint.
//
// Description:
//
// Deletes the specified flow endpoint. This operation is irreversible. Proceed with caution.
//
// @param request - DeleteFlowEndpointRequest
//
// @return DeleteFlowEndpointResponse
func (client *Client) DeleteFlowEndpoint(flowName *string, flowEndpointName *string, request *DeleteFlowEndpointRequest) (_result *DeleteFlowEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteFlowEndpointResponse{}
	_body, _err := client.DeleteFlowEndpointWithOptions(flowName, flowEndpointName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a workflow version.
//
// Description:
//
// Deletes a specified version of a workflow. This operation is irreversible. Proceed with caution.
//
// @param request - DeleteFlowVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFlowVersionResponse
func (client *Client) DeleteFlowVersionWithOptions(flowName *string, flowVersion *string, request *DeleteFlowVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteFlowVersionResponse, _err error) {
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
		Action:      dara.String("DeleteFlowVersion"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/flows/" + dara.PercentEncode(dara.StringValue(flowName)) + "/versions/" + dara.PercentEncode(dara.StringValue(flowVersion))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFlowVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a workflow version.
//
// Description:
//
// Deletes a specified version of a workflow. This operation is irreversible. Proceed with caution.
//
// @param request - DeleteFlowVersionRequest
//
// @return DeleteFlowVersionResponse
func (client *Client) DeleteFlowVersion(flowName *string, flowVersion *string, request *DeleteFlowVersionRequest) (_result *DeleteFlowVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteFlowVersionResponse{}
	_body, _err := client.DeleteFlowVersionWithOptions(flowName, flowVersion, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes an IM bot.
//
// Description:
//
// This operation deletes an IM bot via a `DELETE` request to the `/2025-09-10/agents/im-bots/{imBotId}` endpoint. A successful request returns an HTTP `204 No Content` status code and an empty response body.
//
// @param request - DeleteIMBotRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIMBotResponse
func (client *Client) DeleteIMBotWithOptions(imBotId *string, request *DeleteIMBotRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteIMBotResponse, _err error) {
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
		Action:      dara.String("DeleteIMBot"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/im-bots/" + dara.PercentEncode(dara.StringValue(imBotId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteIMBotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an IM bot.
//
// Description:
//
// This operation deletes an IM bot via a `DELETE` request to the `/2025-09-10/agents/im-bots/{imBotId}` endpoint. A successful request returns an HTTP `204 No Content` status code and an empty response body.
//
// @param request - DeleteIMBotRequest
//
// @return DeleteIMBotResponse
func (client *Client) DeleteIMBot(imBotId *string, request *DeleteIMBotRequest) (_result *DeleteIMBotResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteIMBotResponse{}
	_body, _err := client.DeleteIMBotWithOptions(imBotId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a knowledge base.
//
// > This operation is permanent and cannot be undone.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteKnowledgeBaseResponse
func (client *Client) DeleteKnowledgeBaseWithOptions(knowledgeBaseName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteKnowledgeBaseResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteKnowledgeBase"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/knowledgebases/" + dara.PercentEncode(dara.StringValue(knowledgeBaseName))),
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
// Deletes a knowledge base.
//
// > This operation is permanent and cannot be undone.
//
// @return DeleteKnowledgeBaseResponse
func (client *Client) DeleteKnowledgeBase(knowledgeBaseName *string) (_result *DeleteKnowledgeBaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteKnowledgeBaseResponse{}
	_body, _err := client.DeleteKnowledgeBaseWithOptions(knowledgeBaseName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a memory collection.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMemoryCollectionResponse
func (client *Client) DeleteMemoryCollectionWithOptions(memoryCollectionName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteMemoryCollectionResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteMemoryCollection"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/memory-collections/" + dara.PercentEncode(dara.StringValue(memoryCollectionName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteMemoryCollectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a memory collection.
//
// @return DeleteMemoryCollectionResponse
func (client *Client) DeleteMemoryCollection(memoryCollectionName *string) (_result *DeleteMemoryCollectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteMemoryCollectionResponse{}
	_body, _err := client.DeleteMemoryCollectionWithOptions(memoryCollectionName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a model proxy.
//
// Description:
//
// This operation deletes the specified model proxy configuration without deleting the underlying models or related resources.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteModelProxyResponse
func (client *Client) DeleteModelProxyWithOptions(modelProxyName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteModelProxyResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteModelProxy"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/model-proxies/" + dara.PercentEncode(dara.StringValue(modelProxyName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteModelProxyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a model proxy.
//
// Description:
//
// This operation deletes the specified model proxy configuration without deleting the underlying models or related resources.
//
// @return DeleteModelProxyResponse
func (client *Client) DeleteModelProxy(modelProxyName *string) (_result *DeleteModelProxyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteModelProxyResponse{}
	_body, _err := client.DeleteModelProxyWithOptions(modelProxyName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a model service.
//
// Description:
//
// This operation deletes a model service. You must specify the name of the service to delete.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteModelServiceResponse
func (client *Client) DeleteModelServiceWithOptions(modelServiceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteModelServiceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteModelService"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/model-services/" + dara.PercentEncode(dara.StringValue(modelServiceName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteModelServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a model service.
//
// Description:
//
// This operation deletes a model service. You must specify the name of the service to delete.
//
// @return DeleteModelServiceResponse
func (client *Client) DeleteModelService(modelServiceName *string) (_result *DeleteModelServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteModelServiceResponse{}
	_body, _err := client.DeleteModelServiceWithOptions(modelServiceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a sandbox instance.
//
// Description:
//
// Deletes a sandbox instance.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSandboxResponse
func (client *Client) DeleteSandboxWithOptions(sandboxId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteSandboxResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteSandbox"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/sandboxes/" + dara.PercentEncode(dara.StringValue(sandboxId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteSandboxResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a sandbox instance.
//
// Description:
//
// Deletes a sandbox instance.
//
// @return DeleteSandboxResponse
func (client *Client) DeleteSandbox(sandboxId *string) (_result *DeleteSandboxResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteSandboxResponse{}
	_body, _err := client.DeleteSandboxWithOptions(sandboxId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a template.
//
// Description:
//
// Deletes the specified template. After you delete a template, you can no longer use it to create new sandboxes.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTemplateResponse
func (client *Client) DeleteTemplateWithOptions(templateName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteTemplateResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTemplate"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/templates/" + dara.PercentEncode(dara.StringValue(templateName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a template.
//
// Description:
//
// Deletes the specified template. After you delete a template, you can no longer use it to create new sandboxes.
//
// @return DeleteTemplateResponse
func (client *Client) DeleteTemplate(templateName *string) (_result *DeleteTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTemplateResponse{}
	_body, _err := client.DeleteTemplateWithOptions(templateName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete a tool
//
// Description:
//
// Delete the specified tool. The delete operation is irreversible. Proceed with caution. After the tool is deleted, all Agents that reference this tool will no longer be able to use it.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteToolResponse
func (client *Client) DeleteToolWithOptions(toolName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteToolResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTool"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/tools/" + dara.PercentEncode(dara.StringValue(toolName))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteToolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete a tool
//
// Description:
//
// Delete the specified tool. The delete operation is irreversible. Proceed with caution. After the tool is deleted, all Agents that reference this tool will no longer be able to use it.
//
// @return DeleteToolResponse
func (client *Client) DeleteTool(toolName *string) (_result *DeleteToolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteToolResponse{}
	_body, _err := client.DeleteToolWithOptions(toolName, headers, runtime)
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
// Description:
//
// Deletes the specified workspace.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkspaceResponse
func (client *Client) DeleteWorkspaceWithOptions(workspaceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteWorkspaceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteWorkspace"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId))),
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
// Description:
//
// Deletes the specified workspace.
//
// @return DeleteWorkspaceResponse
func (client *Client) DeleteWorkspace(workspaceId *string) (_result *DeleteWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteWorkspaceResponse{}
	_body, _err := client.DeleteWorkspaceWithOptions(workspaceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Obtains an access token.
//
// Description:
//
// Obtains a temporary accessToken that is used to authenticate subsequent API requests.
//
// @param request - GetAccessTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccessTokenResponse
func (client *Client) GetAccessTokenWithOptions(request *GetAccessTokenRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAccessTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceId) {
		query["resourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceName) {
		query["resourceName"] = request.ResourceName
	}

	if !dara.IsNil(request.ResourceType) {
		query["resourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAccessToken"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/accessToken"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAccessTokenResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains an access token.
//
// Description:
//
// Obtains a temporary accessToken that is used to authenticate subsequent API requests.
//
// @param request - GetAccessTokenRequest
//
// @return GetAccessTokenResponse
func (client *Client) GetAccessToken(request *GetAccessTokenRequest) (_result *GetAccessTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAccessTokenResponse{}
	_body, _err := client.GetAccessTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get Agent Runtime
//
// Description:
//
// Retrieves detailed information about a specified agent runtime by its agent runtime ID, including configuration, status, resource usage, and more.
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
// # Get Agent Runtime
//
// Description:
//
// Retrieves detailed information about a specified agent runtime by its agent runtime ID, including configuration, status, resource usage, and more.
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
// # Get Agent Runtime Access Endpoint
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
// # Get Agent Runtime Access Endpoint
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
// # GetBrowserSandbox
//
// Description:
//
// Retrieves detailed information about a specified browser instance by browser ID, including configuration, status, resource usage, and more.
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
// # GetBrowserSandbox
//
// Description:
//
// Retrieves detailed information about a specified browser instance by browser ID, including configuration, status, resource usage, and more.
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
// Retrieves an interpreter.
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
// Retrieves an interpreter.
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
// Retrieves information about a specific credential.
//
// Description:
//
// Retrieves detailed information about a specified credential, including its configuration, metadata, and related resources.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCredentialResponse
func (client *Client) GetCredentialWithOptions(credentialName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCredentialResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCredential"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/credentials/" + dara.PercentEncode(dara.StringValue(credentialName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCredentialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves information about a specific credential.
//
// Description:
//
// Retrieves detailed information about a specified credential, including its configuration, metadata, and related resources.
//
// @return GetCredentialResponse
func (client *Client) GetCredential(credentialName *string) (_result *GetCredentialResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCredentialResponse{}
	_body, _err := client.GetCredentialWithOptions(credentialName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the configuration of a custom domain.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomDomainResponse
func (client *Client) GetCustomDomainWithOptions(domainName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCustomDomainResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCustomDomain"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/custom-domains/" + dara.PercentEncode(dara.StringValue(domainName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCustomDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the configuration of a custom domain.
//
// @return GetCustomDomainResponse
func (client *Client) GetCustomDomain(domainName *string) (_result *GetCustomDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCustomDomainResponse{}
	_body, _err := client.GetCustomDomainWithOptions(domainName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get flow details
//
// Description:
//
// 根据工作流ID获取指定工作流的详细信息，包括配置、定义、版本信息等。
//
// @param request - GetFlowRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFlowResponse
func (client *Client) GetFlowWithOptions(flowName *string, request *GetFlowRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetFlowResponse, _err error) {
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
		Action:      dara.String("GetFlow"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/flows/" + dara.PercentEncode(dara.StringValue(flowName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get flow details
//
// Description:
//
// 根据工作流ID获取指定工作流的详细信息，包括配置、定义、版本信息等。
//
// @param request - GetFlowRequest
//
// @return GetFlowResponse
func (client *Client) GetFlow(flowName *string, request *GetFlowRequest) (_result *GetFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFlowResponse{}
	_body, _err := client.GetFlowWithOptions(flowName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get the workflow draft.
//
// Description:
//
// 获取指定工作流的草稿版本，返回草稿中的配置信息。
//
// @param request - GetFlowDraftRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFlowDraftResponse
func (client *Client) GetFlowDraftWithOptions(flowName *string, request *GetFlowDraftRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetFlowDraftResponse, _err error) {
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
		Action:      dara.String("GetFlowDraft"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/flows/" + dara.PercentEncode(dara.StringValue(flowName)) + "/draft"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFlowDraftResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get the workflow draft.
//
// Description:
//
// 获取指定工作流的草稿版本，返回草稿中的配置信息。
//
// @param request - GetFlowDraftRequest
//
// @return GetFlowDraftResponse
func (client *Client) GetFlowDraft(flowName *string, request *GetFlowDraftRequest) (_result *GetFlowDraftResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFlowDraftResponse{}
	_body, _err := client.GetFlowDraftWithOptions(flowName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a workflow endpoint.
//
// Description:
//
// 根据工作流ID和端点ID获取指定工作流端点的详细信息。
//
// @param request - GetFlowEndpointRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFlowEndpointResponse
func (client *Client) GetFlowEndpointWithOptions(flowName *string, flowEndpointName *string, request *GetFlowEndpointRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetFlowEndpointResponse, _err error) {
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
		Action:      dara.String("GetFlowEndpoint"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/flows/" + dara.PercentEncode(dara.StringValue(flowName)) + "/endpoints/" + dara.PercentEncode(dara.StringValue(flowEndpointName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFlowEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a workflow endpoint.
//
// Description:
//
// 根据工作流ID和端点ID获取指定工作流端点的详细信息。
//
// @param request - GetFlowEndpointRequest
//
// @return GetFlowEndpointResponse
func (client *Client) GetFlowEndpoint(flowName *string, flowEndpointName *string, request *GetFlowEndpointRequest) (_result *GetFlowEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFlowEndpointResponse{}
	_body, _err := client.GetFlowEndpointWithOptions(flowName, flowEndpointName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a workflow version.
//
// Description:
//
// Retrieves the details of a specific workflow version, including a complete configuration snapshot of its definition, environment variables, tracing configuration, and logging configuration.
//
// @param request - GetFlowVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFlowVersionResponse
func (client *Client) GetFlowVersionWithOptions(flowName *string, flowVersion *string, request *GetFlowVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetFlowVersionResponse, _err error) {
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
		Action:      dara.String("GetFlowVersion"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/flows/" + dara.PercentEncode(dara.StringValue(flowName)) + "/versions/" + dara.PercentEncode(dara.StringValue(flowVersion))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFlowVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a workflow version.
//
// Description:
//
// Retrieves the details of a specific workflow version, including a complete configuration snapshot of its definition, environment variables, tracing configuration, and logging configuration.
//
// @param request - GetFlowVersionRequest
//
// @return GetFlowVersionResponse
func (client *Client) GetFlowVersion(flowName *string, flowVersion *string, request *GetFlowVersionRequest) (_result *GetFlowVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFlowVersionResponse{}
	_body, _err := client.GetFlowVersionWithOptions(flowName, flowVersion, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Gets the details of a specific IM Bot.
//
// Description:
//
// GET /2025-09-10/agents/im-bots/{imBotId}；200 OK，Body 标准包装，data 为 IMBotInfo
//
// @param request - GetIMBotRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIMBotResponse
func (client *Client) GetIMBotWithOptions(imBotId *string, request *GetIMBotRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetIMBotResponse, _err error) {
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
		Action:      dara.String("GetIMBot"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/im-bots/" + dara.PercentEncode(dara.StringValue(imBotId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIMBotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets the details of a specific IM Bot.
//
// Description:
//
// GET /2025-09-10/agents/im-bots/{imBotId}；200 OK，Body 标准包装，data 为 IMBotInfo
//
// @param request - GetIMBotRequest
//
// @return GetIMBotResponse
func (client *Client) GetIMBot(imBotId *string, request *GetIMBotRequest) (_result *GetIMBotResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetIMBotResponse{}
	_body, _err := client.GetIMBotWithOptions(imBotId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Gets information about a knowledge base.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetKnowledgeBaseResponse
func (client *Client) GetKnowledgeBaseWithOptions(knowledgeBaseName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetKnowledgeBaseResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetKnowledgeBase"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/knowledgebases/" + dara.PercentEncode(dara.StringValue(knowledgeBaseName))),
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
// Gets information about a knowledge base.
//
// @return GetKnowledgeBaseResponse
func (client *Client) GetKnowledgeBase(knowledgeBaseName *string) (_result *GetKnowledgeBaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetKnowledgeBaseResponse{}
	_body, _err := client.GetKnowledgeBaseWithOptions(knowledgeBaseName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves details for a specific memory collection.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMemoryCollectionResponse
func (client *Client) GetMemoryCollectionWithOptions(memoryCollectionName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMemoryCollectionResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetMemoryCollection"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/memory-collections/" + dara.PercentEncode(dara.StringValue(memoryCollectionName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetMemoryCollectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves details for a specific memory collection.
//
// @return GetMemoryCollectionResponse
func (client *Client) GetMemoryCollection(memoryCollectionName *string) (_result *GetMemoryCollectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMemoryCollectionResponse{}
	_body, _err := client.GetMemoryCollectionWithOptions(memoryCollectionName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves model governance information.
//
// Description:
//
// This operation retrieves the configuration details of a specific model proxy.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetModelProxyResponse
func (client *Client) GetModelProxyWithOptions(modelProxyName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetModelProxyResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetModelProxy"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/model-proxies/" + dara.PercentEncode(dara.StringValue(modelProxyName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetModelProxyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves model governance information.
//
// Description:
//
// This operation retrieves the configuration details of a specific model proxy.
//
// @return GetModelProxyResponse
func (client *Client) GetModelProxy(modelProxyName *string) (_result *GetModelProxyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetModelProxyResponse{}
	_body, _err := client.GetModelProxyWithOptions(modelProxyName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified model service.
//
// Description:
//
// Retrieves the details of a specified model service.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetModelServiceResponse
func (client *Client) GetModelServiceWithOptions(modelServiceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetModelServiceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetModelService"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/model-services/" + dara.PercentEncode(dara.StringValue(modelServiceName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetModelServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specified model service.
//
// Description:
//
// Retrieves the details of a specified model service.
//
// @return GetModelServiceResponse
func (client *Client) GetModelService(modelServiceName *string) (_result *GetModelServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetModelServiceResponse{}
	_body, _err := client.GetModelServiceWithOptions(modelServiceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specific sandbox.
//
// Description:
//
// Retrieves the details of a specific sandbox by its `sandboxId`, including its status and configuration.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSandboxResponse
func (client *Client) GetSandboxWithOptions(sandboxId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSandboxResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetSandbox"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/sandboxes/" + dara.PercentEncode(dara.StringValue(sandboxId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSandboxResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specific sandbox.
//
// Description:
//
// Retrieves the details of a specific sandbox by its `sandboxId`, including its status and configuration.
//
// @return GetSandboxResponse
func (client *Client) GetSandbox(sandboxId *string) (_result *GetSandboxResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSandboxResponse{}
	_body, _err := client.GetSandboxWithOptions(sandboxId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a template.
//
// Description:
//
// Retrieves the details of a template by its name. The response includes the template\\"s configuration and status.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTemplateResponse
func (client *Client) GetTemplateWithOptions(templateName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTemplateResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTemplate"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/templates/" + dara.PercentEncode(dara.StringValue(templateName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a template.
//
// Description:
//
// Retrieves the details of a template by its name. The response includes the template\\"s configuration and status.
//
// @return GetTemplateResponse
func (client *Client) GetTemplate(templateName *string) (_result *GetTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTemplateResponse{}
	_body, _err := client.GetTemplateWithOptions(templateName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get a tool
//
// Description:
//
// Obtain the complete configuration information of a tool by its name, including basic information, resource configuration, network configuration, running status, and all other detailed information.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetToolResponse
func (client *Client) GetToolWithOptions(toolName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetToolResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTool"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/tools/" + dara.PercentEncode(dara.StringValue(toolName))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetToolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get a tool
//
// Description:
//
// Obtain the complete configuration information of a tool by its name, including basic information, resource configuration, network configuration, running status, and all other detailed information.
//
// @return GetToolResponse
func (client *Client) GetTool(toolName *string) (_result *GetToolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetToolResponse{}
	_body, _err := client.GetToolWithOptions(toolName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves the details of a specific workspace.
//
// Description:
//
// Retrieves the details of a specific workspace.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkspaceResponse
func (client *Client) GetWorkspaceWithOptions(workspaceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetWorkspaceResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWorkspace"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId))),
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
// Retrieves the details of a specific workspace.
//
// Description:
//
// Retrieves the details of a specific workspace.
//
// @return GetWorkspaceResponse
func (client *Client) GetWorkspace(workspaceId *string) (_result *GetWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetWorkspaceResponse{}
	_body, _err := client.GetWorkspaceWithOptions(workspaceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Gets the discovery endpoints for a workspace.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkspaceDiscoveryEndpointsResponse
func (client *Client) GetWorkspaceDiscoveryEndpointsWithOptions(workspaceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetWorkspaceDiscoveryEndpointsResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetWorkspaceDiscoveryEndpoints"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/discovery/endpoints"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetWorkspaceDiscoveryEndpointsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets the discovery endpoints for a workspace.
//
// @return GetWorkspaceDiscoveryEndpointsResponse
func (client *Client) GetWorkspaceDiscoveryEndpoints(workspaceId *string) (_result *GetWorkspaceDiscoveryEndpointsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetWorkspaceDiscoveryEndpointsResponse{}
	_body, _err := client.GetWorkspaceDiscoveryEndpointsWithOptions(workspaceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # List Agent Runtime Endpoints
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

	if !dara.IsNil(request.SearchMode) {
		query["searchMode"] = request.SearchMode
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
// # List Agent Runtime Endpoints
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
// Retrieves agent runtime versions.
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
// Retrieves agent runtime versions.
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
// Retrieves a list of agent runtimes.
//
// Description:
//
// Retrieves a list of agent runtimes for the current user. You can filter the results based on criteria such as name and tags. This operation supports pagination.
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

	if !dara.IsNil(request.DiscoveryResourceGroupId) {
		query["discoveryResourceGroupId"] = request.DiscoveryResourceGroupId
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

	if !dara.IsNil(request.SearchMode) {
		query["searchMode"] = request.SearchMode
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	if !dara.IsNil(request.SystemTags) {
		query["systemTags"] = request.SystemTags
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	if !dara.IsNil(request.WorkspaceIds) {
		query["workspaceIds"] = request.WorkspaceIds
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
// Retrieves a list of agent runtimes.
//
// Description:
//
// Retrieves a list of agent runtimes for the current user. You can filter the results based on criteria such as name and tags. This operation supports pagination.
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
// # List Browser Sandboxes
//
// Description:
//
// Retrieves a list of all browser instances for the current user. Supports filtering by conditions such as name and status, and supports paginated queries.
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
// # List Browser Sandboxes
//
// Description:
//
// Retrieves a list of all browser instances for the current user. Supports filtering by conditions such as name and status, and supports paginated queries.
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
// # List Code Interpreters
//
// Description:
//
// Retrieve a list of all code interpreter instances for the current user, with support for filtering by name and pagination.
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
// # List Code Interpreters
//
// Description:
//
// Retrieve a list of all code interpreter instances for the current user, with support for filtering by name and pagination.
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
// Lists all credentials.
//
// Description:
//
// Lists the credentials in your account. This operation supports filtering and pagination.
//
// @param request - ListCredentialsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCredentialsResponse
func (client *Client) ListCredentialsWithOptions(request *ListCredentialsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListCredentialsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.CredentialAuthType) {
		query["credentialAuthType"] = request.CredentialAuthType
	}

	if !dara.IsNil(request.CredentialName) {
		query["credentialName"] = request.CredentialName
	}

	if !dara.IsNil(request.CredentialSourceType) {
		query["credentialSourceType"] = request.CredentialSourceType
	}

	if !dara.IsNil(request.Enabled) {
		query["enabled"] = request.Enabled
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Provider) {
		query["provider"] = request.Provider
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	if !dara.IsNil(request.WorkspaceIds) {
		query["workspaceIds"] = request.WorkspaceIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCredentials"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/credentials"),
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
// Lists all credentials.
//
// Description:
//
// Lists the credentials in your account. This operation supports filtering and pagination.
//
// @param request - ListCredentialsRequest
//
// @return ListCredentialsResponse
func (client *Client) ListCredentials(request *ListCredentialsRequest) (_result *ListCredentialsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListCredentialsResponse{}
	_body, _err := client.ListCredentialsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists custom domains.
//
// @param request - ListCustomDomainsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomDomainsResponse
func (client *Client) ListCustomDomainsWithOptions(request *ListCustomDomainsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListCustomDomainsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DomainName) {
		query["domainName"] = request.DomainName
	}

	if !dara.IsNil(request.DomainType) {
		query["domainType"] = request.DomainType
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResourceName) {
		query["resourceName"] = request.ResourceName
	}

	if !dara.IsNil(request.ResourceType) {
		query["resourceType"] = request.ResourceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListCustomDomains"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/custom-domains"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListCustomDomainsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists custom domains.
//
// @param request - ListCustomDomainsRequest
//
// @return ListCustomDomainsResponse
func (client *Client) ListCustomDomains(request *ListCustomDomainsRequest) (_result *ListCustomDomainsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListCustomDomainsResponse{}
	_body, _err := client.ListCustomDomainsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// List workflow endpoints.
//
// Description:
//
// Retrieve all endpoints for a specified workflow, with pagination support.
//
// @param request - ListFlowEndpointsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFlowEndpointsResponse
func (client *Client) ListFlowEndpointsWithOptions(flowId *string, request *ListFlowEndpointsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListFlowEndpointsResponse, _err error) {
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
		Action:      dara.String("ListFlowEndpoints"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/flows/" + dara.PercentEncode(dara.StringValue(flowId)) + "/endpoints"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFlowEndpointsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// List workflow endpoints.
//
// Description:
//
// Retrieve all endpoints for a specified workflow, with pagination support.
//
// @param request - ListFlowEndpointsRequest
//
// @return ListFlowEndpointsResponse
func (client *Client) ListFlowEndpoints(flowId *string, request *ListFlowEndpointsRequest) (_result *ListFlowEndpointsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFlowEndpointsResponse{}
	_body, _err := client.ListFlowEndpointsWithOptions(flowId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the versions of a flow.
//
// Description:
//
// Returns a paginated list of all versions for a specified flow.
//
// @param request - ListFlowVersionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFlowVersionsResponse
func (client *Client) ListFlowVersionsWithOptions(flowName *string, request *ListFlowVersionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListFlowVersionsResponse, _err error) {
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
		Action:      dara.String("ListFlowVersions"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/flows/" + dara.PercentEncode(dara.StringValue(flowName)) + "/versions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFlowVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists the versions of a flow.
//
// Description:
//
// Returns a paginated list of all versions for a specified flow.
//
// @param request - ListFlowVersionsRequest
//
// @return ListFlowVersionsResponse
func (client *Client) ListFlowVersions(flowName *string, request *ListFlowVersionsRequest) (_result *ListFlowVersionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFlowVersionsResponse{}
	_body, _err := client.ListFlowVersionsWithOptions(flowName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # List flows
//
// Description:
//
// 获取当前用户的工作流列表，支持按名称、工作空间等条件过滤，支持分页查询。
//
// @param request - ListFlowsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFlowsResponse
func (client *Client) ListFlowsWithOptions(request *ListFlowsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListFlowsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FlowName) {
		query["flowName"] = request.FlowName
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	if !dara.IsNil(request.WorkspaceIds) {
		query["workspaceIds"] = request.WorkspaceIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFlows"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/flows"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListFlowsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # List flows
//
// Description:
//
// 获取当前用户的工作流列表，支持按名称、工作空间等条件过滤，支持分页查询。
//
// @param request - ListFlowsRequest
//
// @return ListFlowsResponse
func (client *Client) ListFlows(request *ListFlowsRequest) (_result *ListFlowsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFlowsResponse{}
	_body, _err := client.ListFlowsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Retrieves a paginated list of IM bots.
//
// Description:
//
// Send a GET request to the `/2025-09-10/agents/im-bots` endpoint to retrieve a paginated list of IM bots. Use the `botName`, `agentRuntimeId`, `botBizType`, and `status` query parameters to filter the results. For pagination, the `pageNumber` defaults to 1 and the `pageSize` defaults to 20, with a maximum of 100.
//
// @param request - ListIMBotsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIMBotsResponse
func (client *Client) ListIMBotsWithOptions(request *ListIMBotsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListIMBotsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentRuntimeId) {
		query["agentRuntimeId"] = request.AgentRuntimeId
	}

	if !dara.IsNil(request.BotBizType) {
		query["botBizType"] = request.BotBizType
	}

	if !dara.IsNil(request.BotName) {
		query["botName"] = request.BotName
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
		Action:      dara.String("ListIMBots"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/im-bots"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListIMBotsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves a paginated list of IM bots.
//
// Description:
//
// Send a GET request to the `/2025-09-10/agents/im-bots` endpoint to retrieve a paginated list of IM bots. Use the `botName`, `agentRuntimeId`, `botBizType`, and `status` query parameters to filter the results. For pagination, the `pageNumber` defaults to 1 and the `pageSize` defaults to 20, with a maximum of 100.
//
// @param request - ListIMBotsRequest
//
// @return ListIMBotsResponse
func (client *Client) ListIMBots(request *ListIMBotsRequest) (_result *ListIMBotsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListIMBotsResponse{}
	_body, _err := client.ListIMBotsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists the knowledge bases in your account.
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
	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Provider) {
		query["provider"] = request.Provider
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	if !dara.IsNil(request.WorkspaceIds) {
		query["workspaceIds"] = request.WorkspaceIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListKnowledgeBases"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/knowledgebases"),
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
// Lists the knowledge bases in your account.
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
// Returns a paginated list of memory collections.
//
// @param request - ListMemoryCollectionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMemoryCollectionsResponse
func (client *Client) ListMemoryCollectionsWithOptions(request *ListMemoryCollectionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMemoryCollectionsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.MemoryCollectionName) {
		query["memoryCollectionName"] = request.MemoryCollectionName
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

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	if !dara.IsNil(request.WorkspaceIds) {
		query["workspaceIds"] = request.WorkspaceIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMemoryCollections"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/memory-collections"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListMemoryCollectionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Returns a paginated list of memory collections.
//
// @param request - ListMemoryCollectionsRequest
//
// @return ListMemoryCollectionsResponse
func (client *Client) ListMemoryCollections(request *ListMemoryCollectionsRequest) (_result *ListMemoryCollectionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMemoryCollectionsResponse{}
	_body, _err := client.ListMemoryCollectionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists all model providers.
//
// Description:
//
// Lists the available model providers. This operation supports filtering and pagination.
//
// @param request - ListModelProvidersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListModelProvidersResponse
func (client *Client) ListModelProvidersWithOptions(request *ListModelProvidersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListModelProvidersResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ModelName) {
		query["modelName"] = request.ModelName
	}

	if !dara.IsNil(request.ModelType) {
		query["modelType"] = request.ModelType
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Provider) {
		query["provider"] = request.Provider
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListModelProviders"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/model-providers"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListModelProvidersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists all model providers.
//
// Description:
//
// Lists the available model providers. This operation supports filtering and pagination.
//
// @param request - ListModelProvidersRequest
//
// @return ListModelProvidersResponse
func (client *Client) ListModelProviders(request *ListModelProvidersRequest) (_result *ListModelProvidersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListModelProvidersResponse{}
	_body, _err := client.ListModelProvidersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists all Model Proxies.
//
// Description:
//
// Retrieves a paginated list of all Model Proxies for the current user. You can filter the list by status.
//
// @param request - ListModelProxiesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListModelProxiesResponse
func (client *Client) ListModelProxiesWithOptions(request *ListModelProxiesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListModelProxiesResponse, _err error) {
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

	if !dara.IsNil(request.ProxyMode) {
		query["proxyMode"] = request.ProxyMode
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	if !dara.IsNil(request.WorkspaceIds) {
		query["workspaceIds"] = request.WorkspaceIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListModelProxies"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/model-proxies"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListModelProxiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists all Model Proxies.
//
// Description:
//
// Retrieves a paginated list of all Model Proxies for the current user. You can filter the list by status.
//
// @param request - ListModelProxiesRequest
//
// @return ListModelProxiesResponse
func (client *Client) ListModelProxies(request *ListModelProxiesRequest) (_result *ListModelProxiesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListModelProxiesResponse{}
	_body, _err := client.ListModelProxiesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists all models.
//
// Description:
//
// Retrieves a list of all models for the current user. You can filter the models by type and paginate the results.
//
// @param request - ListModelServicesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListModelServicesResponse
func (client *Client) ListModelServicesWithOptions(request *ListModelServicesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListModelServicesResponse, _err error) {
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

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Provider) {
		query["provider"] = request.Provider
	}

	if !dara.IsNil(request.ProviderType) {
		query["providerType"] = request.ProviderType
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	if !dara.IsNil(request.WorkspaceIds) {
		query["workspaceIds"] = request.WorkspaceIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListModelServices"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/model-services"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListModelServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists all models.
//
// Description:
//
// Retrieves a list of all models for the current user. You can filter the models by type and paginate the results.
//
// @param request - ListModelServicesRequest
//
// @return ListModelServicesResponse
func (client *Client) ListModelServices(request *ListModelServicesRequest) (_result *ListModelServicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListModelServicesResponse{}
	_body, _err := client.ListModelServicesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Returns a list of sandboxes.
//
// Description:
//
// Retrieves a list of sandboxes for the current user. You can filter the results by template name, template type, or status. This operation supports pagination.
//
// @param request - ListSandboxesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSandboxesResponse
func (client *Client) ListSandboxesWithOptions(request *ListSandboxesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSandboxesResponse, _err error) {
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

	if !dara.IsNil(request.SandboxId) {
		query["sandboxId"] = request.SandboxId
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	if !dara.IsNil(request.TemplateName) {
		query["templateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TemplateType) {
		query["templateType"] = request.TemplateType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSandboxes"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/sandboxes"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListSandboxesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Returns a list of sandboxes.
//
// Description:
//
// Retrieves a list of sandboxes for the current user. You can filter the results by template name, template type, or status. This operation supports pagination.
//
// @param request - ListSandboxesRequest
//
// @return ListSandboxesResponse
func (client *Client) ListSandboxes(request *ListSandboxesRequest) (_result *ListSandboxesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSandboxesResponse{}
	_body, _err := client.ListSandboxesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists templates.
//
// Description:
//
// Lists all templates for the current user. You can filter the results by template type and use pagination.
//
// @param request - ListTemplatesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTemplatesResponse
func (client *Client) ListTemplatesWithOptions(request *ListTemplatesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTemplatesResponse, _err error) {
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

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	if !dara.IsNil(request.TemplateName) {
		query["templateName"] = request.TemplateName
	}

	if !dara.IsNil(request.TemplateType) {
		query["templateType"] = request.TemplateType
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	if !dara.IsNil(request.WorkspaceIds) {
		query["workspaceIds"] = request.WorkspaceIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTemplates"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/templates"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Lists templates.
//
// Description:
//
// Lists all templates for the current user. You can filter the results by template type and use pagination.
//
// @param request - ListTemplatesRequest
//
// @return ListTemplatesResponse
func (client *Client) ListTemplates(request *ListTemplatesRequest) (_result *ListTemplatesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTemplatesResponse{}
	_body, _err := client.ListTemplatesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # List tools
//
// Description:
//
// Query the tool list. Support paged query and conditional filtering by tool type, workspace, and other criteria. Return the list of tools that meet the conditions and paging information.
//
// @param request - ListToolsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListToolsResponse
func (client *Client) ListToolsWithOptions(request *ListToolsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListToolsResponse, _err error) {
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

	if !dara.IsNil(request.ToolName) {
		query["toolName"] = request.ToolName
	}

	if !dara.IsNil(request.ToolType) {
		query["toolType"] = request.ToolType
	}

	if !dara.IsNil(request.WorkspaceId) {
		query["workspaceId"] = request.WorkspaceId
	}

	if !dara.IsNil(request.WorkspaceIds) {
		query["workspaceIds"] = request.WorkspaceIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTools"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/tools"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListToolsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # List tools
//
// Description:
//
// Query the tool list. Support paged query and conditional filtering by tool type, workspace, and other criteria. Return the list of tools that meet the conditions and paging information.
//
// @param request - ListToolsRequest
//
// @return ListToolsResponse
func (client *Client) ListTools(request *ListToolsRequest) (_result *ListToolsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListToolsResponse{}
	_body, _err := client.ListToolsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Lists all workspaces in your account.
//
// Description:
//
// Lists all workspaces.
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListWorkspaces"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/workspaces"),
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
// Lists all workspaces in your account.
//
// Description:
//
// Lists all workspaces.
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
// Pause the sandbox and retain snapshots of its memory and file system. The sandbox enters the PAUSED state so that it can be recovered later.
//
// Description:
//
// This API is used to pause a sandbox. When invoked, the system takes a snapshot of the sandbox, capturing and persisting the memory and disk states. The user can recover the sandbox at a later time.
//
// Note that sandbox snapshots are retained for only 30 days. After 30 days, recovery becomes unavailable.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PauseSandboxResponse
func (client *Client) PauseSandboxWithOptions(sandboxId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PauseSandboxResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("PauseSandbox"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/sandboxes/" + dara.PercentEncode(dara.StringValue(sandboxId)) + "/pause"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PauseSandboxResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Pause the sandbox and retain snapshots of its memory and file system. The sandbox enters the PAUSED state so that it can be recovered later.
//
// Description:
//
// This API is used to pause a sandbox. When invoked, the system takes a snapshot of the sandbox, capturing and persisting the memory and disk states. The user can recover the sandbox at a later time.
//
// Note that sandbox snapshots are retained for only 30 days. After 30 days, recovery becomes unavailable.
//
// @return PauseSandboxResponse
func (client *Client) PauseSandbox(sandboxId *string) (_result *PauseSandboxResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PauseSandboxResponse{}
	_body, _err := client.PauseSandboxWithOptions(sandboxId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Publishes a workflow version.
//
// Description:
//
// Publishes a new version of a specified workflow. This operation supports version management and rollbacks.
//
// @param request - PublishFlowVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishFlowVersionResponse
func (client *Client) PublishFlowVersionWithOptions(flowName *string, request *PublishFlowVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PublishFlowVersionResponse, _err error) {
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
		Action:      dara.String("PublishFlowVersion"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/flows/" + dara.PercentEncode(dara.StringValue(flowName)) + "/versions"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PublishFlowVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Publishes a workflow version.
//
// Description:
//
// Publishes a new version of a specified workflow. This operation supports version management and rollbacks.
//
// @param request - PublishFlowVersionRequest
//
// @return PublishFlowVersionResponse
func (client *Client) PublishFlowVersion(flowName *string, request *PublishFlowVersionRequest) (_result *PublishFlowVersionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PublishFlowVersionResponse{}
	_body, _err := client.PublishFlowVersionWithOptions(flowName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Publishes a new version of an agent runtime.
//
// Description:
//
// Publishes a new version for a specified agent runtime for version management and deployment. The new version can include code updates, configuration changes, and other content.
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
// Publishes a new version of an agent runtime.
//
// Description:
//
// Publishes a new version for a specified agent runtime for version management and deployment. The new version can include code updates, configuration changes, and other content.
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
// Resume a paused sandbox instance to restore it from the PAUSED state to the READY (running) state.
//
// Description:
//
// This API resumes a sandbox instance from the paused state to the ready state, allowing the user to invoke it and restore it to the state it was in before being paused.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeSandboxResponse
func (client *Client) ResumeSandboxWithOptions(sandboxId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ResumeSandboxResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ResumeSandbox"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/sandboxes/" + dara.PercentEncode(dara.StringValue(sandboxId)) + "/resume"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ResumeSandboxResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Resume a paused sandbox instance to restore it from the PAUSED state to the READY (running) state.
//
// Description:
//
// This API resumes a sandbox instance from the paused state to the ready state, allowing the user to invoke it and restore it to the state it was in before being paused.
//
// @return ResumeSandboxResponse
func (client *Client) ResumeSandbox(sandboxId *string) (_result *ResumeSandboxResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ResumeSandboxResponse{}
	_body, _err := client.ResumeSandboxWithOptions(sandboxId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops a sandbox.
//
// Description:
//
// Stops the specified sandbox instance. After the operation, the sandbox enters the TERMINATED state.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopSandboxResponse
func (client *Client) StopSandboxWithOptions(sandboxId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopSandboxResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopSandbox"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/sandboxes/" + dara.PercentEncode(dara.StringValue(sandboxId)) + "/stop"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopSandboxResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops a sandbox.
//
// Description:
//
// Stops the specified sandbox instance. After the operation, the sandbox enters the TERMINATED state.
//
// @return StopSandboxResponse
func (client *Client) StopSandbox(sandboxId *string) (_result *StopSandboxResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopSandboxResponse{}
	_body, _err := client.StopSandboxWithOptions(sandboxId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stops the TemplateMCP service.
//
// Description:
//
// Stopping the MCP service deletes the associated MCP resources and makes the endpoint inaccessible.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopTemplateMCPResponse
func (client *Client) StopTemplateMCPWithOptions(templateName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopTemplateMCPResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopTemplateMCP"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/templates/" + dara.PercentEncode(dara.StringValue(templateName)) + "/mcp/stop"),
		Method:      dara.String("PATCH"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopTemplateMCPResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Stops the TemplateMCP service.
//
// Description:
//
// Stopping the MCP service deletes the associated MCP resources and makes the endpoint inaccessible.
//
// @return StopTemplateMCPResponse
func (client *Client) StopTemplateMCP(templateName *string) (_result *StopTemplateMCPResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopTemplateMCPResponse{}
	_body, _err := client.StopTemplateMCPWithOptions(templateName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # UpdateAgentRuntime
//
// Description:
//
// Updates the configuration of a specified agent runtime, including resource allocation, network configuration, and environment variables. The update operation triggers a runtime restart.
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
// # UpdateAgentRuntime
//
// Description:
//
// Updates the configuration of a specified agent runtime, including resource allocation, network configuration, and environment variables. The update operation triggers a runtime restart.
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
// # Update Agent Runtime Endpoint
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
// # Update Agent Runtime Endpoint
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
// Updates a credential.
//
// Description:
//
// Updates the configuration of a specified credential.
//
// @param request - UpdateCredentialRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCredentialResponse
func (client *Client) UpdateCredentialWithOptions(credentialName *string, request *UpdateCredentialRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateCredentialResponse, _err error) {
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
		Action:      dara.String("UpdateCredential"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/credentials/" + dara.PercentEncode(dara.StringValue(credentialName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCredentialResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a credential.
//
// Description:
//
// Updates the configuration of a specified credential.
//
// @param request - UpdateCredentialRequest
//
// @return UpdateCredentialResponse
func (client *Client) UpdateCredential(credentialName *string, request *UpdateCredentialRequest) (_result *UpdateCredentialResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateCredentialResponse{}
	_body, _err := client.UpdateCredentialWithOptions(credentialName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a custom domain.
//
// @param request - UpdateCustomDomainRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCustomDomainResponse
func (client *Client) UpdateCustomDomainWithOptions(domainName *string, request *UpdateCustomDomainRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateCustomDomainResponse, _err error) {
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
		Action:      dara.String("UpdateCustomDomain"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/custom-domains/" + dara.PercentEncode(dara.StringValue(domainName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateCustomDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a custom domain.
//
// @param request - UpdateCustomDomainRequest
//
// @return UpdateCustomDomainResponse
func (client *Client) UpdateCustomDomain(domainName *string, request *UpdateCustomDomainRequest) (_result *UpdateCustomDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateCustomDomainResponse{}
	_body, _err := client.UpdateCustomDomainWithOptions(domainName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a workflow.
//
// Description:
//
// Updates the configuration of a specified workflow, including the definition, execution mode, and environment variables.
//
// @param request - UpdateFlowRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFlowResponse
func (client *Client) UpdateFlowWithOptions(flowName *string, request *UpdateFlowRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateFlowResponse, _err error) {
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
		Action:      dara.String("UpdateFlow"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/flows/" + dara.PercentEncode(dara.StringValue(flowName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFlowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a workflow.
//
// Description:
//
// Updates the configuration of a specified workflow, including the definition, execution mode, and environment variables.
//
// @param request - UpdateFlowRequest
//
// @return UpdateFlowResponse
func (client *Client) UpdateFlow(flowName *string, request *UpdateFlowRequest) (_result *UpdateFlowResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateFlowResponse{}
	_body, _err := client.UpdateFlowWithOptions(flowName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Update a flow draft.
//
// Description:
//
// 更新指定工作流的草稿版本，草稿更新不影响已发布的工作流版本。
//
// @param request - UpdateFlowDraftRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFlowDraftResponse
func (client *Client) UpdateFlowDraftWithOptions(flowName *string, request *UpdateFlowDraftRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateFlowDraftResponse, _err error) {
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
		Action:      dara.String("UpdateFlowDraft"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/flows/" + dara.PercentEncode(dara.StringValue(flowName)) + "/draft"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFlowDraftResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update a flow draft.
//
// Description:
//
// 更新指定工作流的草稿版本，草稿更新不影响已发布的工作流版本。
//
// @param request - UpdateFlowDraftRequest
//
// @return UpdateFlowDraftResponse
func (client *Client) UpdateFlowDraft(flowName *string, request *UpdateFlowDraftRequest) (_result *UpdateFlowDraftResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateFlowDraftResponse{}
	_body, _err := client.UpdateFlowDraftWithOptions(flowName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Update Flow Endpoint
//
// Description:
//
// 更新指定工作流端点的配置信息，包括目标版本、路由配置等。
//
// @param request - UpdateFlowEndpointRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFlowEndpointResponse
func (client *Client) UpdateFlowEndpointWithOptions(flowName *string, flowEndpointName *string, request *UpdateFlowEndpointRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateFlowEndpointResponse, _err error) {
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
		Action:      dara.String("UpdateFlowEndpoint"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/flows/" + dara.PercentEncode(dara.StringValue(flowName)) + "/endpoints/" + dara.PercentEncode(dara.StringValue(flowEndpointName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateFlowEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update Flow Endpoint
//
// Description:
//
// 更新指定工作流端点的配置信息，包括目标版本、路由配置等。
//
// @param request - UpdateFlowEndpointRequest
//
// @return UpdateFlowEndpointResponse
func (client *Client) UpdateFlowEndpoint(flowName *string, flowEndpointName *string, request *UpdateFlowEndpointRequest) (_result *UpdateFlowEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateFlowEndpointResponse{}
	_body, _err := client.UpdateFlowEndpointWithOptions(flowName, flowEndpointName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates an IM bot.
//
// Description:
//
// PUT /2025-09-10/agents/im-bots/{imBotId}；成功建议 HTTP 202，Body 标准包装，data 为更新后 IMBotInfo
//
// @param request - UpdateIMBotRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIMBotResponse
func (client *Client) UpdateIMBotWithOptions(imBotId *string, request *UpdateIMBotRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateIMBotResponse, _err error) {
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
		Action:      dara.String("UpdateIMBot"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/im-bots/" + dara.PercentEncode(dara.StringValue(imBotId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateIMBotResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an IM bot.
//
// Description:
//
// PUT /2025-09-10/agents/im-bots/{imBotId}；成功建议 HTTP 202，Body 标准包装，data 为更新后 IMBotInfo
//
// @param request - UpdateIMBotRequest
//
// @return UpdateIMBotResponse
func (client *Client) UpdateIMBot(imBotId *string, request *UpdateIMBotRequest) (_result *UpdateIMBotResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateIMBotResponse{}
	_body, _err := client.UpdateIMBotWithOptions(imBotId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configuration of a knowledge base.
//
// @param request - UpdateKnowledgeBaseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateKnowledgeBaseResponse
func (client *Client) UpdateKnowledgeBaseWithOptions(knowledgeBaseName *string, request *UpdateKnowledgeBaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateKnowledgeBaseResponse, _err error) {
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
		Action:      dara.String("UpdateKnowledgeBase"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/knowledgebases/" + dara.PercentEncode(dara.StringValue(knowledgeBaseName))),
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
// Updates the configuration of a knowledge base.
//
// @param request - UpdateKnowledgeBaseRequest
//
// @return UpdateKnowledgeBaseResponse
func (client *Client) UpdateKnowledgeBase(knowledgeBaseName *string, request *UpdateKnowledgeBaseRequest) (_result *UpdateKnowledgeBaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateKnowledgeBaseResponse{}
	_body, _err := client.UpdateKnowledgeBaseWithOptions(knowledgeBaseName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a memory collection.
//
// @param request - UpdateMemoryCollectionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMemoryCollectionResponse
func (client *Client) UpdateMemoryCollectionWithOptions(memoryCollectionName *string, request *UpdateMemoryCollectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateMemoryCollectionResponse, _err error) {
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
		Action:      dara.String("UpdateMemoryCollection"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/memory-collections/" + dara.PercentEncode(dara.StringValue(memoryCollectionName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateMemoryCollectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a memory collection.
//
// @param request - UpdateMemoryCollectionRequest
//
// @return UpdateMemoryCollectionResponse
func (client *Client) UpdateMemoryCollection(memoryCollectionName *string, request *UpdateMemoryCollectionRequest) (_result *UpdateMemoryCollectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateMemoryCollectionResponse{}
	_body, _err := client.UpdateMemoryCollectionWithOptions(memoryCollectionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the configuration of a model proxy.
//
// Description:
//
// Use this operation to update the configuration of a specific model proxy.
//
// @param request - UpdateModelProxyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateModelProxyResponse
func (client *Client) UpdateModelProxyWithOptions(modelProxyName *string, request *UpdateModelProxyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateModelProxyResponse, _err error) {
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
		Action:      dara.String("UpdateModelProxy"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/model-proxies/" + dara.PercentEncode(dara.StringValue(modelProxyName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateModelProxyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the configuration of a model proxy.
//
// Description:
//
// Use this operation to update the configuration of a specific model proxy.
//
// @param request - UpdateModelProxyRequest
//
// @return UpdateModelProxyResponse
func (client *Client) UpdateModelProxy(modelProxyName *string, request *UpdateModelProxyRequest) (_result *UpdateModelProxyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateModelProxyResponse{}
	_body, _err := client.UpdateModelProxyWithOptions(modelProxyName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a model service.
//
// Description:
//
// This operation modifies the configuration of an existing model service.
//
// @param request - UpdateModelServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateModelServiceResponse
func (client *Client) UpdateModelServiceWithOptions(modelServiceName *string, request *UpdateModelServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateModelServiceResponse, _err error) {
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
		Action:      dara.String("UpdateModelService"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/model-services/" + dara.PercentEncode(dara.StringValue(modelServiceName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateModelServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a model service.
//
// Description:
//
// This operation modifies the configuration of an existing model service.
//
// @param request - UpdateModelServiceRequest
//
// @return UpdateModelServiceResponse
func (client *Client) UpdateModelService(modelServiceName *string, request *UpdateModelServiceRequest) (_result *UpdateModelServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateModelServiceResponse{}
	_body, _err := client.UpdateModelServiceWithOptions(modelServiceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a template.
//
// Description:
//
// Updates a template\\"s configuration, including its resource configuration, network configuration, and environment variables.
//
// @param request - UpdateTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTemplateResponse
func (client *Client) UpdateTemplateWithOptions(templateName *string, request *UpdateTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateTemplateResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTemplate"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/templates/" + dara.PercentEncode(dara.StringValue(templateName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a template.
//
// Description:
//
// Updates a template\\"s configuration, including its resource configuration, network configuration, and environment variables.
//
// @param request - UpdateTemplateRequest
//
// @return UpdateTemplateResponse
func (client *Client) UpdateTemplate(templateName *string, request *UpdateTemplateRequest) (_result *UpdateTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateTemplateResponse{}
	_body, _err := client.UpdateTemplateWithOptions(templateName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a tool.
//
// Description:
//
// Updates the configuration of an existing tool. You can modify its description, resource configuration, network configuration, and more. This operation supports partial updates. You only need to specify the fields that you want to modify.
//
// @param request - UpdateToolRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateToolResponse
func (client *Client) UpdateToolWithOptions(toolName *string, request *UpdateToolRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateToolResponse, _err error) {
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
		Action:      dara.String("UpdateTool"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/agents/tools/" + dara.PercentEncode(dara.StringValue(toolName))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateToolResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a tool.
//
// Description:
//
// Updates the configuration of an existing tool. You can modify its description, resource configuration, network configuration, and more. This operation supports partial updates. You only need to specify the fields that you want to modify.
//
// @param request - UpdateToolRequest
//
// @return UpdateToolResponse
func (client *Client) UpdateTool(toolName *string, request *UpdateToolRequest) (_result *UpdateToolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateToolResponse{}
	_body, _err := client.UpdateToolWithOptions(toolName, request, headers, runtime)
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
// Description:
//
// Updates the properties of a workspace.
//
// @param request - UpdateWorkspaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkspaceResponse
func (client *Client) UpdateWorkspaceWithOptions(workspaceId *string, request *UpdateWorkspaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateWorkspaceResponse, _err error) {
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
		Action:      dara.String("UpdateWorkspace"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId))),
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
// Description:
//
// Updates the properties of a workspace.
//
// @param request - UpdateWorkspaceRequest
//
// @return UpdateWorkspaceResponse
func (client *Client) UpdateWorkspace(workspaceId *string, request *UpdateWorkspaceRequest) (_result *UpdateWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateWorkspaceResponse{}
	_body, _err := client.UpdateWorkspaceWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the discovery endpoint configuration for a specified workspace.
//
// @param request - UpdateWorkspaceDiscoveryEndpointsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateWorkspaceDiscoveryEndpointsResponse
func (client *Client) UpdateWorkspaceDiscoveryEndpointsWithOptions(workspaceId *string, request *UpdateWorkspaceDiscoveryEndpointsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateWorkspaceDiscoveryEndpointsResponse, _err error) {
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
		Action:      dara.String("UpdateWorkspaceDiscoveryEndpoints"),
		Version:     dara.String("2025-09-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/2025-09-10/workspaces/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/discovery/endpoints"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateWorkspaceDiscoveryEndpointsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the discovery endpoint configuration for a specified workspace.
//
// @param request - UpdateWorkspaceDiscoveryEndpointsRequest
//
// @return UpdateWorkspaceDiscoveryEndpointsResponse
func (client *Client) UpdateWorkspaceDiscoveryEndpoints(workspaceId *string, request *UpdateWorkspaceDiscoveryEndpointsRequest) (_result *UpdateWorkspaceDiscoveryEndpointsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateWorkspaceDiscoveryEndpointsResponse{}
	_body, _err := client.UpdateWorkspaceDiscoveryEndpointsWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
