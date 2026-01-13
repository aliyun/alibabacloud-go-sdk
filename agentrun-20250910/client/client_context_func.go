// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

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
func (client *Client) ActivateTemplateMCPWithContext(ctx context.Context, templateName *string, request *ActivateTemplateMCPRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ActivateTemplateMCPResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAgentRuntimeResponse
func (client *Client) CreateAgentRuntimeWithContext(ctx context.Context, request *CreateAgentRuntimeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAgentRuntimeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAgentRuntimeEndpointResponse
func (client *Client) CreateAgentRuntimeEndpointWithContext(ctx context.Context, agentRuntimeId *string, request *CreateAgentRuntimeEndpointRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAgentRuntimeEndpointResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateBrowserResponse
func (client *Client) CreateBrowserWithContext(ctx context.Context, request *CreateBrowserRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateBrowserResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCodeInterpreterResponse
func (client *Client) CreateCodeInterpreterWithContext(ctx context.Context, request *CreateCodeInterpreterRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateCodeInterpreterResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCredentialResponse
func (client *Client) CreateCredentialWithContext(ctx context.Context, request *CreateCredentialRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateCredentialResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomDomainResponse
func (client *Client) CreateCustomDomainWithContext(ctx context.Context, request *CreateCustomDomainRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateCustomDomainResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMemoryCollectionResponse
func (client *Client) CreateMemoryCollectionWithContext(ctx context.Context, request *CreateMemoryCollectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateMemoryCollectionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateModelProxyResponse
func (client *Client) CreateModelProxyWithContext(ctx context.Context, request *CreateModelProxyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateModelProxyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateModelServiceResponse
func (client *Client) CreateModelServiceWithContext(ctx context.Context, request *CreateModelServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateModelServiceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSandboxResponse
func (client *Client) CreateSandboxWithContext(ctx context.Context, request *CreateSandboxRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateSandboxResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTemplateResponse
func (client *Client) CreateTemplateWithContext(ctx context.Context, request *CreateTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTemplateResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAgentRuntimeResponse
func (client *Client) DeleteAgentRuntimeWithContext(ctx context.Context, agentRuntimeId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAgentRuntimeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAgentRuntimeEndpointResponse
func (client *Client) DeleteAgentRuntimeEndpointWithContext(ctx context.Context, agentRuntimeId *string, agentRuntimeEndpointId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAgentRuntimeEndpointResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteBrowserResponse
func (client *Client) DeleteBrowserWithContext(ctx context.Context, browserId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteBrowserResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCodeInterpreterResponse
func (client *Client) DeleteCodeInterpreterWithContext(ctx context.Context, codeInterpreterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteCodeInterpreterResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCredentialResponse
func (client *Client) DeleteCredentialWithContext(ctx context.Context, credentialName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteCredentialResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteCustomDomainResponse
func (client *Client) DeleteCustomDomainWithContext(ctx context.Context, domainName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteCustomDomainResponse, _err error) {
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteKnowledgeBaseResponse
func (client *Client) DeleteKnowledgeBaseWithContext(ctx context.Context, knowledgeBaseName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteKnowledgeBaseResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMemoryCollectionResponse
func (client *Client) DeleteMemoryCollectionWithContext(ctx context.Context, memoryCollectionName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteMemoryCollectionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteModelProxyResponse
func (client *Client) DeleteModelProxyWithContext(ctx context.Context, modelProxyName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteModelProxyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteModelServiceResponse
func (client *Client) DeleteModelServiceWithContext(ctx context.Context, modelServiceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteModelServiceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSandboxResponse
func (client *Client) DeleteSandboxWithContext(ctx context.Context, sandboxId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteSandboxResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTemplateResponse
func (client *Client) DeleteTemplateWithContext(ctx context.Context, templateName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteTemplateResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccessTokenResponse
func (client *Client) GetAccessTokenWithContext(ctx context.Context, request *GetAccessTokenRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAccessTokenResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentRuntimeResponse
func (client *Client) GetAgentRuntimeWithContext(ctx context.Context, agentRuntimeId *string, request *GetAgentRuntimeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAgentRuntimeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAgentRuntimeEndpointResponse
func (client *Client) GetAgentRuntimeEndpointWithContext(ctx context.Context, agentRuntimeId *string, agentRuntimeEndpointId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAgentRuntimeEndpointResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetBrowserResponse
func (client *Client) GetBrowserWithContext(ctx context.Context, browserId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetBrowserResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCodeInterpreterResponse
func (client *Client) GetCodeInterpreterWithContext(ctx context.Context, codeInterpreterId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCodeInterpreterResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCredentialResponse
func (client *Client) GetCredentialWithContext(ctx context.Context, credentialName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCredentialResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCustomDomainResponse
func (client *Client) GetCustomDomainWithContext(ctx context.Context, domainName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCustomDomainResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetKnowledgeBaseResponse
func (client *Client) GetKnowledgeBaseWithContext(ctx context.Context, knowledgeBaseName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetKnowledgeBaseResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMemoryCollectionResponse
func (client *Client) GetMemoryCollectionWithContext(ctx context.Context, memoryCollectionName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetMemoryCollectionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetModelProxyResponse
func (client *Client) GetModelProxyWithContext(ctx context.Context, modelProxyName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetModelProxyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetModelServiceResponse
func (client *Client) GetModelServiceWithContext(ctx context.Context, modelServiceName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetModelServiceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSandboxResponse
func (client *Client) GetSandboxWithContext(ctx context.Context, sandboxId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSandboxResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTemplateResponse
func (client *Client) GetTemplateWithContext(ctx context.Context, templateName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTemplateResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAgentRuntimeEndpointsResponse
func (client *Client) ListAgentRuntimeEndpointsWithContext(ctx context.Context, agentRuntimeId *string, request *ListAgentRuntimeEndpointsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAgentRuntimeEndpointsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAgentRuntimeVersionsResponse
func (client *Client) ListAgentRuntimeVersionsWithContext(ctx context.Context, agentRuntimeId *string, request *ListAgentRuntimeVersionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAgentRuntimeVersionsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAgentRuntimesResponse
func (client *Client) ListAgentRuntimesWithContext(ctx context.Context, request *ListAgentRuntimesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAgentRuntimesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListBrowsersResponse
func (client *Client) ListBrowsersWithContext(ctx context.Context, request *ListBrowsersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListBrowsersResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCodeInterpretersResponse
func (client *Client) ListCodeInterpretersWithContext(ctx context.Context, request *ListCodeInterpretersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListCodeInterpretersResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCredentialsResponse
func (client *Client) ListCredentialsWithContext(ctx context.Context, request *ListCredentialsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListCredentialsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListCustomDomainsResponse
func (client *Client) ListCustomDomainsWithContext(ctx context.Context, request *ListCustomDomainsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListCustomDomainsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMemoryCollectionsResponse
func (client *Client) ListMemoryCollectionsWithContext(ctx context.Context, request *ListMemoryCollectionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListMemoryCollectionsResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListModelProvidersResponse
func (client *Client) ListModelProvidersWithContext(ctx context.Context, request *ListModelProvidersRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListModelProvidersResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListModelProxiesResponse
func (client *Client) ListModelProxiesWithContext(ctx context.Context, request *ListModelProxiesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListModelProxiesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListModelServicesResponse
func (client *Client) ListModelServicesWithContext(ctx context.Context, request *ListModelServicesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListModelServicesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSandboxesResponse
func (client *Client) ListSandboxesWithContext(ctx context.Context, request *ListSandboxesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListSandboxesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTemplatesResponse
func (client *Client) ListTemplatesWithContext(ctx context.Context, request *ListTemplatesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTemplatesResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PublishRuntimeVersionResponse
func (client *Client) PublishRuntimeVersionWithContext(ctx context.Context, agentRuntimeId *string, request *PublishRuntimeVersionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PublishRuntimeVersionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopSandboxResponse
func (client *Client) StopSandboxWithContext(ctx context.Context, sandboxId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopSandboxResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopTemplateMCPResponse
func (client *Client) StopTemplateMCPWithContext(ctx context.Context, templateName *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopTemplateMCPResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAgentRuntimeResponse
func (client *Client) UpdateAgentRuntimeWithContext(ctx context.Context, agentRuntimeId *string, request *UpdateAgentRuntimeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAgentRuntimeResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAgentRuntimeEndpointResponse
func (client *Client) UpdateAgentRuntimeEndpointWithContext(ctx context.Context, agentRuntimeId *string, agentRuntimeEndpointId *string, request *UpdateAgentRuntimeEndpointRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAgentRuntimeEndpointResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCredentialResponse
func (client *Client) UpdateCredentialWithContext(ctx context.Context, credentialName *string, request *UpdateCredentialRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateCredentialResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCustomDomainResponse
func (client *Client) UpdateCustomDomainWithContext(ctx context.Context, domainName *string, request *UpdateCustomDomainRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateCustomDomainResponse, _err error) {
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
func (client *Client) UpdateKnowledgeBaseWithContext(ctx context.Context, knowledgeBaseName *string, request *UpdateKnowledgeBaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateKnowledgeBaseResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMemoryCollectionResponse
func (client *Client) UpdateMemoryCollectionWithContext(ctx context.Context, memoryCollectionName *string, request *UpdateMemoryCollectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateMemoryCollectionResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateModelProxyResponse
func (client *Client) UpdateModelProxyWithContext(ctx context.Context, modelProxyName *string, request *UpdateModelProxyRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateModelProxyResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateModelServiceResponse
func (client *Client) UpdateModelServiceWithContext(ctx context.Context, modelServiceName *string, request *UpdateModelServiceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateModelServiceResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTemplateResponse
func (client *Client) UpdateTemplateWithContext(ctx context.Context, templateName *string, request *UpdateTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateTemplateResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
