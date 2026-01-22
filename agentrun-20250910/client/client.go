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
// 启动模板的MCP服务器
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
// 启动模板的MCP服务器
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
// # Create a credential
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
// # Create a credential
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
// 创建自定义域名
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
// 创建自定义域名
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
// 添加记忆存储
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
// 添加记忆存储
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
// 新增模型
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
// 新增模型
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
// 新增模型
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
// 新增模型
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
// 创建沙箱
//
// Description:
//
// 根据模板创建一个新的沙箱实例。沙箱是运行时的执行环境，可以执行代码或运行浏览器。
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
// 创建沙箱
//
// Description:
//
// 根据模板创建一个新的沙箱实例。沙箱是运行时的执行环境，可以执行代码或运行浏览器。
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
// 创建模板
//
// Description:
//
// 创建一个新的模板，用于后续创建沙箱。模板定义了沙箱的运行时环境、资源配置等。
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
// 创建模板
//
// Description:
//
// 创建一个新的模板，用于后续创建沙箱。模板定义了沙箱的运行时环境、资源配置等。
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
// # Delete a credential
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
// # Delete a credential
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
// # Delete a custom domain
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
// # Delete a custom domain
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
// 删除知识库
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
// 删除知识库
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
// 删除记忆存储
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
// 删除记忆存储
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
// 删除模型
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
// 删除模型
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
// 删除模型
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
// 删除模型
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
// 删除Sandbox
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
// 删除Sandbox
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
// 删除模板
//
// Description:
//
// 删除指定的模板。删除后，该模板将无法再用于创建新的沙箱。
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
// 删除模板
//
// Description:
//
// 删除指定的模板。删除后，该模板将无法再用于创建新的沙箱。
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
// # Get access token for a resource
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
// # Get access token for a resource
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
// # Get a credential
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
// # Get a credential
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
// 获取自定义域名详情
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
// 获取自定义域名详情
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
// 获取知识库
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
// 获取知识库
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
// 查询记忆存储详情
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
// 查询记忆存储详情
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
// 查看model
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
// 查看model
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
// 查看model
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
// 查看model
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
// 获取沙箱
//
// Description:
//
// 根据沙箱ID获取指定沙箱的详细信息，包括状态、配置等。
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
// 获取沙箱
//
// Description:
//
// 根据沙箱ID获取指定沙箱的详细信息，包括状态、配置等。
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
// 获取模板
//
// Description:
//
// 根据模板名称获取指定模板的详细信息，包括配置、状态等。
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
// 获取模板
//
// Description:
//
// 根据模板名称获取指定模板的详细信息，包括配置、状态等。
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

	if !dara.IsNil(request.ResourceGroupId) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.SearchMode) {
		query["searchMode"] = request.SearchMode
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
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
// # List credentials
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
// # List credentials
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
// 自定义域名列表
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
// 自定义域名列表
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
// 列出知识库
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
// 列出知识库
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
// 查询记忆存储列表
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
// 查询记忆存储列表
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
// 查询支持的模型提供商及其模型
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
// 查询支持的模型提供商及其模型
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
// model列表
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
// model列表
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
// model列表
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
// model列表
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
// 列出沙箱
//
// Description:
//
// 获取当前用户的所有沙箱列表，支持按模板名称过滤，支持分页查询。
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
// 列出沙箱
//
// Description:
//
// 获取当前用户的所有沙箱列表，支持按模板名称过滤，支持分页查询。
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
// 列出模板
//
// Description:
//
// 获取当前用户的所有模板列表，支持按模板类型过滤，支持分页查询。
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
// 列出模板
//
// Description:
//
// 获取当前用户的所有模板列表，支持按模板类型过滤，支持分页查询。
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
// 删除沙箱
//
// Description:
//
// 停止指定的沙箱实例。停止后，沙箱将进入TERMINATED状态。
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
// 删除沙箱
//
// Description:
//
// 停止指定的沙箱实例。停止后，沙箱将进入TERMINATED状态。
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
// 停止模板的MCP服务器
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
// 停止模板的MCP服务器
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
// # Update a credential
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
// # Update a credential
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
// 更新自定义域名
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
// 更新自定义域名
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
// 更新知识库
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
// 更新知识库
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
// 修改记忆存储信息
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
// 修改记忆存储信息
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
// 更新模型
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
// 更新模型
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
// 更新模型
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
// 更新模型
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
// 更新模板
//
// Description:
//
// 更新指定模板的配置信息，包括资源配置、网络配置、环境变量等。
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
// 更新模板
//
// Description:
//
// 更新指定模板的配置信息，包括资源配置、网络配置、环境变量等。
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
