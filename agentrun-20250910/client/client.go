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
// Enable the TemplateMCP service.
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
// Enable the TemplateMCP service.
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
// # CreateAgentRuntime
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
// # CreateAgentRuntime
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
// Create a template.
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
// Create a template.
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
// 创建工具
//
// Description:
//
// 创建一个新的工具，支持创建 MCP、函数调用和技能等多种类型的工具。工具创建后可以被 Agent 调用以扩展其能力。
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
// 创建工具
//
// Description:
//
// 创建一个新的工具，支持创建 MCP、函数调用和技能等多种类型的工具。工具创建后可以被 Agent 调用以扩展其能力。
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
// 创建工作空间
//
// Description:
//
// 创建工作空间
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
// 创建工作空间
//
// Description:
//
// 创建工作空间
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
// 删除智能体运行时
//
// Description:
//
// 删除指定的智能体运行时实例，包括其所有相关资源和数据。删除操作不可逆，请谨慎操作。
//
// @param request - DeleteAgentRuntimeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAgentRuntimeResponse
func (client *Client) DeleteAgentRuntimeWithOptions(agentRuntimeId *string, request *DeleteAgentRuntimeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAgentRuntimeResponse, _err error) {
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
// @param request - DeleteAgentRuntimeRequest
//
// @return DeleteAgentRuntimeResponse
func (client *Client) DeleteAgentRuntime(agentRuntimeId *string, request *DeleteAgentRuntimeRequest) (_result *DeleteAgentRuntimeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAgentRuntimeResponse{}
	_body, _err := client.DeleteAgentRuntimeWithOptions(agentRuntimeId, request, headers, runtime)
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
// @param request - DeleteAgentRuntimeEndpointRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAgentRuntimeEndpointResponse
func (client *Client) DeleteAgentRuntimeEndpointWithOptions(agentRuntimeId *string, agentRuntimeEndpointId *string, request *DeleteAgentRuntimeEndpointRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAgentRuntimeEndpointResponse, _err error) {
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
// @param request - DeleteAgentRuntimeEndpointRequest
//
// @return DeleteAgentRuntimeEndpointResponse
func (client *Client) DeleteAgentRuntimeEndpoint(agentRuntimeId *string, agentRuntimeEndpointId *string, request *DeleteAgentRuntimeEndpointRequest) (_result *DeleteAgentRuntimeEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAgentRuntimeEndpointResponse{}
	_body, _err := client.DeleteAgentRuntimeEndpointWithOptions(agentRuntimeId, agentRuntimeEndpointId, request, headers, runtime)
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
// @param request - DeleteBrowserRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBrowserResponse
func (client *Client) DeleteBrowserWithOptions(browserId *string, request *DeleteBrowserRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteBrowserResponse, _err error) {
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
// @param request - DeleteBrowserRequest
//
// @return DeleteBrowserResponse
func (client *Client) DeleteBrowser(browserId *string, request *DeleteBrowserRequest) (_result *DeleteBrowserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteBrowserResponse{}
	_body, _err := client.DeleteBrowserWithOptions(browserId, request, headers, runtime)
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
// @param request - DeleteCodeInterpreterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCodeInterpreterResponse
func (client *Client) DeleteCodeInterpreterWithOptions(codeInterpreterId *string, request *DeleteCodeInterpreterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteCodeInterpreterResponse, _err error) {
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
// @param request - DeleteCodeInterpreterRequest
//
// @return DeleteCodeInterpreterResponse
func (client *Client) DeleteCodeInterpreter(codeInterpreterId *string, request *DeleteCodeInterpreterRequest) (_result *DeleteCodeInterpreterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteCodeInterpreterResponse{}
	_body, _err := client.DeleteCodeInterpreterWithOptions(codeInterpreterId, request, headers, runtime)
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
// @param request - DeleteCredentialRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCredentialResponse
func (client *Client) DeleteCredentialWithOptions(credentialName *string, request *DeleteCredentialRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteCredentialResponse, _err error) {
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
// @param request - DeleteCredentialRequest
//
// @return DeleteCredentialResponse
func (client *Client) DeleteCredential(credentialName *string, request *DeleteCredentialRequest) (_result *DeleteCredentialResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteCredentialResponse{}
	_body, _err := client.DeleteCredentialWithOptions(credentialName, request, headers, runtime)
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
// @param request - DeleteCustomDomainRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomDomainResponse
func (client *Client) DeleteCustomDomainWithOptions(domainName *string, request *DeleteCustomDomainRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteCustomDomainResponse, _err error) {
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
// @param request - DeleteCustomDomainRequest
//
// @return DeleteCustomDomainResponse
func (client *Client) DeleteCustomDomain(domainName *string, request *DeleteCustomDomainRequest) (_result *DeleteCustomDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteCustomDomainResponse{}
	_body, _err := client.DeleteCustomDomainWithOptions(domainName, request, headers, runtime)
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
func (client *Client) DeleteKnowledgeBaseWithOptions(knowledgeBaseName *string, request *DeleteKnowledgeBaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteKnowledgeBaseResponse, _err error) {
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
// @param request - DeleteKnowledgeBaseRequest
//
// @return DeleteKnowledgeBaseResponse
func (client *Client) DeleteKnowledgeBase(knowledgeBaseName *string, request *DeleteKnowledgeBaseRequest) (_result *DeleteKnowledgeBaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteKnowledgeBaseResponse{}
	_body, _err := client.DeleteKnowledgeBaseWithOptions(knowledgeBaseName, request, headers, runtime)
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
// @param request - DeleteMemoryCollectionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMemoryCollectionResponse
func (client *Client) DeleteMemoryCollectionWithOptions(memoryCollectionName *string, request *DeleteMemoryCollectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteMemoryCollectionResponse, _err error) {
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
// @param request - DeleteMemoryCollectionRequest
//
// @return DeleteMemoryCollectionResponse
func (client *Client) DeleteMemoryCollection(memoryCollectionName *string, request *DeleteMemoryCollectionRequest) (_result *DeleteMemoryCollectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteMemoryCollectionResponse{}
	_body, _err := client.DeleteMemoryCollectionWithOptions(memoryCollectionName, request, headers, runtime)
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
// @param request - DeleteModelProxyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteModelProxyResponse
func (client *Client) DeleteModelProxyWithOptions(modelProxyName *string, request *DeleteModelProxyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteModelProxyResponse, _err error) {
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
// @param request - DeleteModelProxyRequest
//
// @return DeleteModelProxyResponse
func (client *Client) DeleteModelProxy(modelProxyName *string, request *DeleteModelProxyRequest) (_result *DeleteModelProxyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteModelProxyResponse{}
	_body, _err := client.DeleteModelProxyWithOptions(modelProxyName, request, headers, runtime)
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
// @param request - DeleteModelServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteModelServiceResponse
func (client *Client) DeleteModelServiceWithOptions(modelServiceName *string, request *DeleteModelServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteModelServiceResponse, _err error) {
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
// @param request - DeleteModelServiceRequest
//
// @return DeleteModelServiceResponse
func (client *Client) DeleteModelService(modelServiceName *string, request *DeleteModelServiceRequest) (_result *DeleteModelServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteModelServiceResponse{}
	_body, _err := client.DeleteModelServiceWithOptions(modelServiceName, request, headers, runtime)
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
// @param request - DeleteSandboxRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSandboxResponse
func (client *Client) DeleteSandboxWithOptions(sandboxId *string, request *DeleteSandboxRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteSandboxResponse, _err error) {
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
// @param request - DeleteSandboxRequest
//
// @return DeleteSandboxResponse
func (client *Client) DeleteSandbox(sandboxId *string, request *DeleteSandboxRequest) (_result *DeleteSandboxResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteSandboxResponse{}
	_body, _err := client.DeleteSandboxWithOptions(sandboxId, request, headers, runtime)
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
// @param request - DeleteTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTemplateResponse
func (client *Client) DeleteTemplateWithOptions(templateName *string, request *DeleteTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteTemplateResponse, _err error) {
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
// @param request - DeleteTemplateRequest
//
// @return DeleteTemplateResponse
func (client *Client) DeleteTemplate(templateName *string, request *DeleteTemplateRequest) (_result *DeleteTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTemplateResponse{}
	_body, _err := client.DeleteTemplateWithOptions(templateName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除工具
//
// Description:
//
// 删除指定的工具。删除操作不可逆，请谨慎操作。删除工具后，所有引用该工具的 Agent 将无法继续使用该工具。
//
// @param request - DeleteToolRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteToolResponse
func (client *Client) DeleteToolWithOptions(toolName *string, request *DeleteToolRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteToolResponse, _err error) {
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
// 删除工具
//
// Description:
//
// 删除指定的工具。删除操作不可逆，请谨慎操作。删除工具后，所有引用该工具的 Agent 将无法继续使用该工具。
//
// @param request - DeleteToolRequest
//
// @return DeleteToolResponse
func (client *Client) DeleteTool(toolName *string, request *DeleteToolRequest) (_result *DeleteToolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteToolResponse{}
	_body, _err := client.DeleteToolWithOptions(toolName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除工作空间
//
// Description:
//
// 删除工作空间
//
// @param request - DeleteWorkspaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteWorkspaceResponse
func (client *Client) DeleteWorkspaceWithOptions(workspaceId *string, request *DeleteWorkspaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteWorkspaceResponse, _err error) {
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
// 删除工作空间
//
// Description:
//
// 删除工作空间
//
// @param request - DeleteWorkspaceRequest
//
// @return DeleteWorkspaceResponse
func (client *Client) DeleteWorkspace(workspaceId *string, request *DeleteWorkspaceRequest) (_result *DeleteWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteWorkspaceResponse{}
	_body, _err := client.DeleteWorkspaceWithOptions(workspaceId, request, headers, runtime)
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
// @param request - GetAgentRuntimeEndpointRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentRuntimeEndpointResponse
func (client *Client) GetAgentRuntimeEndpointWithOptions(agentRuntimeId *string, agentRuntimeEndpointId *string, request *GetAgentRuntimeEndpointRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAgentRuntimeEndpointResponse, _err error) {
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
// @param request - GetAgentRuntimeEndpointRequest
//
// @return GetAgentRuntimeEndpointResponse
func (client *Client) GetAgentRuntimeEndpoint(agentRuntimeId *string, agentRuntimeEndpointId *string, request *GetAgentRuntimeEndpointRequest) (_result *GetAgentRuntimeEndpointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAgentRuntimeEndpointResponse{}
	_body, _err := client.GetAgentRuntimeEndpointWithOptions(agentRuntimeId, agentRuntimeEndpointId, request, headers, runtime)
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
// @param request - GetBrowserRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBrowserResponse
func (client *Client) GetBrowserWithOptions(browserId *string, request *GetBrowserRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetBrowserResponse, _err error) {
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
// @param request - GetBrowserRequest
//
// @return GetBrowserResponse
func (client *Client) GetBrowser(browserId *string, request *GetBrowserRequest) (_result *GetBrowserResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetBrowserResponse{}
	_body, _err := client.GetBrowserWithOptions(browserId, request, headers, runtime)
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
// @param request - GetCodeInterpreterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCodeInterpreterResponse
func (client *Client) GetCodeInterpreterWithOptions(codeInterpreterId *string, request *GetCodeInterpreterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCodeInterpreterResponse, _err error) {
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
// @param request - GetCodeInterpreterRequest
//
// @return GetCodeInterpreterResponse
func (client *Client) GetCodeInterpreter(codeInterpreterId *string, request *GetCodeInterpreterRequest) (_result *GetCodeInterpreterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCodeInterpreterResponse{}
	_body, _err := client.GetCodeInterpreterWithOptions(codeInterpreterId, request, headers, runtime)
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
// @param request - GetCredentialRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCredentialResponse
func (client *Client) GetCredentialWithOptions(credentialName *string, request *GetCredentialRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCredentialResponse, _err error) {
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
// @param request - GetCredentialRequest
//
// @return GetCredentialResponse
func (client *Client) GetCredential(credentialName *string, request *GetCredentialRequest) (_result *GetCredentialResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCredentialResponse{}
	_body, _err := client.GetCredentialWithOptions(credentialName, request, headers, runtime)
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
// @param request - GetCustomDomainRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomDomainResponse
func (client *Client) GetCustomDomainWithOptions(domainName *string, request *GetCustomDomainRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCustomDomainResponse, _err error) {
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
// @param request - GetCustomDomainRequest
//
// @return GetCustomDomainResponse
func (client *Client) GetCustomDomain(domainName *string, request *GetCustomDomainRequest) (_result *GetCustomDomainResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCustomDomainResponse{}
	_body, _err := client.GetCustomDomainWithOptions(domainName, request, headers, runtime)
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
// @param request - GetKnowledgeBaseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetKnowledgeBaseResponse
func (client *Client) GetKnowledgeBaseWithOptions(knowledgeBaseName *string, request *GetKnowledgeBaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetKnowledgeBaseResponse, _err error) {
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
// @param request - GetKnowledgeBaseRequest
//
// @return GetKnowledgeBaseResponse
func (client *Client) GetKnowledgeBase(knowledgeBaseName *string, request *GetKnowledgeBaseRequest) (_result *GetKnowledgeBaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetKnowledgeBaseResponse{}
	_body, _err := client.GetKnowledgeBaseWithOptions(knowledgeBaseName, request, headers, runtime)
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
// @param request - GetMemoryCollectionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMemoryCollectionResponse
func (client *Client) GetMemoryCollectionWithOptions(memoryCollectionName *string, request *GetMemoryCollectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMemoryCollectionResponse, _err error) {
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
// @param request - GetMemoryCollectionRequest
//
// @return GetMemoryCollectionResponse
func (client *Client) GetMemoryCollection(memoryCollectionName *string, request *GetMemoryCollectionRequest) (_result *GetMemoryCollectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetMemoryCollectionResponse{}
	_body, _err := client.GetMemoryCollectionWithOptions(memoryCollectionName, request, headers, runtime)
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
// @param request - GetModelProxyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetModelProxyResponse
func (client *Client) GetModelProxyWithOptions(modelProxyName *string, request *GetModelProxyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetModelProxyResponse, _err error) {
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
// @param request - GetModelProxyRequest
//
// @return GetModelProxyResponse
func (client *Client) GetModelProxy(modelProxyName *string, request *GetModelProxyRequest) (_result *GetModelProxyResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetModelProxyResponse{}
	_body, _err := client.GetModelProxyWithOptions(modelProxyName, request, headers, runtime)
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
// @param request - GetModelServiceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetModelServiceResponse
func (client *Client) GetModelServiceWithOptions(modelServiceName *string, request *GetModelServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetModelServiceResponse, _err error) {
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
// @param request - GetModelServiceRequest
//
// @return GetModelServiceResponse
func (client *Client) GetModelService(modelServiceName *string, request *GetModelServiceRequest) (_result *GetModelServiceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetModelServiceResponse{}
	_body, _err := client.GetModelServiceWithOptions(modelServiceName, request, headers, runtime)
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
// @param request - GetSandboxRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSandboxResponse
func (client *Client) GetSandboxWithOptions(sandboxId *string, request *GetSandboxRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSandboxResponse, _err error) {
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
// @param request - GetSandboxRequest
//
// @return GetSandboxResponse
func (client *Client) GetSandbox(sandboxId *string, request *GetSandboxRequest) (_result *GetSandboxResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSandboxResponse{}
	_body, _err := client.GetSandboxWithOptions(sandboxId, request, headers, runtime)
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
// @param request - GetTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTemplateResponse
func (client *Client) GetTemplateWithOptions(templateName *string, request *GetTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTemplateResponse, _err error) {
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
// @param request - GetTemplateRequest
//
// @return GetTemplateResponse
func (client *Client) GetTemplate(templateName *string, request *GetTemplateRequest) (_result *GetTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTemplateResponse{}
	_body, _err := client.GetTemplateWithOptions(templateName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取工具详情
//
// Description:
//
// 根据工具名称获取工具的完整配置信息，包括工具的基本信息、资源配置、网络配置、运行状态等所有详细信息。
//
// @param request - GetToolRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetToolResponse
func (client *Client) GetToolWithOptions(toolName *string, request *GetToolRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetToolResponse, _err error) {
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
// 获取工具详情
//
// Description:
//
// 根据工具名称获取工具的完整配置信息，包括工具的基本信息、资源配置、网络配置、运行状态等所有详细信息。
//
// @param request - GetToolRequest
//
// @return GetToolResponse
func (client *Client) GetTool(toolName *string, request *GetToolRequest) (_result *GetToolResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetToolResponse{}
	_body, _err := client.GetToolWithOptions(toolName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看工作空间
//
// Description:
//
// 查看工作空间
//
// @param request - GetWorkspaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkspaceResponse
func (client *Client) GetWorkspaceWithOptions(workspaceId *string, request *GetWorkspaceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetWorkspaceResponse, _err error) {
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
// 查看工作空间
//
// Description:
//
// 查看工作空间
//
// @param request - GetWorkspaceRequest
//
// @return GetWorkspaceResponse
func (client *Client) GetWorkspace(workspaceId *string, request *GetWorkspaceRequest) (_result *GetWorkspaceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetWorkspaceResponse{}
	_body, _err := client.GetWorkspaceWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取工作空间下的发现端点
//
// @param request - GetWorkspaceDiscoveryEndpointsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetWorkspaceDiscoveryEndpointsResponse
func (client *Client) GetWorkspaceDiscoveryEndpointsWithOptions(workspaceId *string, request *GetWorkspaceDiscoveryEndpointsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetWorkspaceDiscoveryEndpointsResponse, _err error) {
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
// 获取工作空间下的发现端点
//
// @param request - GetWorkspaceDiscoveryEndpointsRequest
//
// @return GetWorkspaceDiscoveryEndpointsResponse
func (client *Client) GetWorkspaceDiscoveryEndpoints(workspaceId *string, request *GetWorkspaceDiscoveryEndpointsRequest) (_result *GetWorkspaceDiscoveryEndpointsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetWorkspaceDiscoveryEndpointsResponse{}
	_body, _err := client.GetWorkspaceDiscoveryEndpointsWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Retrieve the list of access endpoints for an agent runtime
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
// # Retrieve the list of access endpoints for an agent runtime
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
// 工具列表
//
// Description:
//
// 查询工具列表，支持分页查询和按工具类型、工作空间等条件过滤。返回符合条件的工具列表及分页信息。
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
// 工具列表
//
// Description:
//
// 查询工具列表，支持分页查询和按工具类型、工作空间等条件过滤。返回符合条件的工具列表及分页信息。
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
// 获取工作空间列表
//
// Description:
//
// 获取工作空间列表
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
// 获取工作空间列表
//
// Description:
//
// 获取工作空间列表
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
// 暂停沙箱
//
// Description:
//
// 停止指定的沙箱实例。停止后，沙箱将进入TERMINATED状态。
//
// @param request - PauseSandboxRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PauseSandboxResponse
func (client *Client) PauseSandboxWithOptions(sandboxId *string, request *PauseSandboxRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PauseSandboxResponse, _err error) {
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
// 暂停沙箱
//
// Description:
//
// 停止指定的沙箱实例。停止后，沙箱将进入TERMINATED状态。
//
// @param request - PauseSandboxRequest
//
// @return PauseSandboxResponse
func (client *Client) PauseSandbox(sandboxId *string, request *PauseSandboxRequest) (_result *PauseSandboxResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PauseSandboxResponse{}
	_body, _err := client.PauseSandboxWithOptions(sandboxId, request, headers, runtime)
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
// 恢复沙箱
//
// @param request - ResumeSandboxRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ResumeSandboxResponse
func (client *Client) ResumeSandboxWithOptions(sandboxId *string, request *ResumeSandboxRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ResumeSandboxResponse, _err error) {
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
// 恢复沙箱
//
// @param request - ResumeSandboxRequest
//
// @return ResumeSandboxResponse
func (client *Client) ResumeSandbox(sandboxId *string, request *ResumeSandboxRequest) (_result *ResumeSandboxResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ResumeSandboxResponse{}
	_body, _err := client.ResumeSandboxWithOptions(sandboxId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 停止沙箱
//
// Description:
//
// 停止指定的沙箱实例。停止后，沙箱将进入TERMINATED状态。
//
// @param request - StopSandboxRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopSandboxResponse
func (client *Client) StopSandboxWithOptions(sandboxId *string, request *StopSandboxRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopSandboxResponse, _err error) {
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
// 停止沙箱
//
// Description:
//
// 停止指定的沙箱实例。停止后，沙箱将进入TERMINATED状态。
//
// @param request - StopSandboxRequest
//
// @return StopSandboxResponse
func (client *Client) StopSandbox(sandboxId *string, request *StopSandboxRequest) (_result *StopSandboxResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopSandboxResponse{}
	_body, _err := client.StopSandboxWithOptions(sandboxId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Stop the TemplateMCP service.
//
// @param request - StopTemplateMCPRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopTemplateMCPResponse
func (client *Client) StopTemplateMCPWithOptions(templateName *string, request *StopTemplateMCPRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopTemplateMCPResponse, _err error) {
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
// Stop the TemplateMCP service.
//
// @param request - StopTemplateMCPRequest
//
// @return StopTemplateMCPResponse
func (client *Client) StopTemplateMCP(templateName *string, request *StopTemplateMCPRequest) (_result *StopTemplateMCPResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopTemplateMCPResponse{}
	_body, _err := client.StopTemplateMCPWithOptions(templateName, request, headers, runtime)
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

// Summary:
//
// 更新工具
//
// Description:
//
// 更新现有工具的配置信息，可以修改工具的描述、资源配置、网络配置等。更新操作支持部分更新，只需提供需要修改的字段。
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
// 更新工具
//
// Description:
//
// 更新现有工具的配置信息，可以修改工具的描述、资源配置、网络配置等。更新操作支持部分更新，只需提供需要修改的字段。
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
// 更新工作空间
//
// Description:
//
// 更新工作空间
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
// 更新工作空间
//
// Description:
//
// 更新工作空间
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
// 获取工作空间下的发现端点
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
// 获取工作空间下的发现端点
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
