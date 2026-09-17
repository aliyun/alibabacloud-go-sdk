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
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("hologram"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 用于创建一个新的Agent会话并返回会话ID。
//
// Description:
//
// ## 请求说明
//
// - 该接口用于创建一个新的 Agent 会话。
//
// - 通过 `_meta.agent.agentName` 指定绑定的 Agent 名称，这是必填项。
//
// - 可以通过 `_meta.config.sessionSource` 透传会话来源标识，便于后续按来源检索。
//
// - 支持通过 `_meta.config.sessionTags[].sessionTagCode` 传入会话标签。
//
// @param tmpReq - CreateAgentSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAgentSessionResponse
func (client *Client) CreateAgentSessionWithOptions(tmpReq *CreateAgentSessionRequest, runtime *dara.RuntimeOptions) (_result *CreateAgentSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateAgentSessionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Params) {
		request.ParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Params, dara.String("Params"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Jsonrpc) {
		body["Jsonrpc"] = request.Jsonrpc
	}

	if !dara.IsNil(request.ParamsShrink) {
		body["Params"] = request.ParamsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAgentSession"),
		Version:     dara.String("2022-06-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAgentSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 用于创建一个新的Agent会话并返回会话ID。
//
// Description:
//
// ## 请求说明
//
// - 该接口用于创建一个新的 Agent 会话。
//
// - 通过 `_meta.agent.agentName` 指定绑定的 Agent 名称，这是必填项。
//
// - 可以通过 `_meta.config.sessionSource` 透传会话来源标识，便于后续按来源检索。
//
// - 支持通过 `_meta.config.sessionTags[].sessionTagCode` 传入会话标签。
//
// @param request - CreateAgentSessionRequest
//
// @return CreateAgentSessionResponse
func (client *Client) CreateAgentSession(request *CreateAgentSessionRequest) (_result *CreateAgentSessionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &CreateAgentSessionResponse{}
	_body, _err := client.CreateAgentSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 向指定会话发送用户Prompt，并以SSE方式流式接收Agent响应。
//
// Description:
//
// ## 请求说明
//
// - 该 API 用于向指定的会话 ID 发送用户的 Prompt，并以 SSE（Server-Sent Events）流式方式接收来自 Agent 的响应。
//
// - 响应可能包括消息分片、思考过程、工具调用状态更新等信息。
//
// - 如果指定的会话不存在，将通过 SSE 错误帧返回 400 错误。
//
// - `stopReason`字段指示了 Agent 停止本轮对话的原因。
//
// - 可选地提供额外元信息`Meta`来传递更多上下文给服务端。
//
// - 返回的内容符合开源协议 Agent Client Protocol (ACP) 的规范。
//
// @param tmpReq - PromptAgentSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PromptAgentSessionResponse
func (client *Client) PromptAgentSessionWithSSE(tmpReq *PromptAgentSessionRequest, runtime *dara.RuntimeOptions, _yield chan *PromptAgentSessionResponse, _yieldErr chan error) {
	defer close(_yield)
	client.promptAgentSessionWithSSE_opYieldFunc(_yield, _yieldErr, tmpReq, runtime)
	return
}

// Summary:
//
// 向指定会话发送用户Prompt，并以SSE方式流式接收Agent响应。
//
// Description:
//
// ## 请求说明
//
// - 该 API 用于向指定的会话 ID 发送用户的 Prompt，并以 SSE（Server-Sent Events）流式方式接收来自 Agent 的响应。
//
// - 响应可能包括消息分片、思考过程、工具调用状态更新等信息。
//
// - 如果指定的会话不存在，将通过 SSE 错误帧返回 400 错误。
//
// - `stopReason`字段指示了 Agent 停止本轮对话的原因。
//
// - 可选地提供额外元信息`Meta`来传递更多上下文给服务端。
//
// - 返回的内容符合开源协议 Agent Client Protocol (ACP) 的规范。
//
// @param tmpReq - PromptAgentSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PromptAgentSessionResponse
func (client *Client) PromptAgentSessionWithOptions(tmpReq *PromptAgentSessionRequest, runtime *dara.RuntimeOptions) (_result *PromptAgentSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &PromptAgentSessionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Params) {
		request.ParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Params, dara.String("Params"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CallerContext) {
		body["Caller-Context"] = request.CallerContext
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Jsonrpc) {
		body["Jsonrpc"] = request.Jsonrpc
	}

	if !dara.IsNil(request.ParamsShrink) {
		body["Params"] = request.ParamsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PromptAgentSession"),
		Version:     dara.String("2022-06-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &PromptAgentSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 向指定会话发送用户Prompt，并以SSE方式流式接收Agent响应。
//
// Description:
//
// ## 请求说明
//
// - 该 API 用于向指定的会话 ID 发送用户的 Prompt，并以 SSE（Server-Sent Events）流式方式接收来自 Agent 的响应。
//
// - 响应可能包括消息分片、思考过程、工具调用状态更新等信息。
//
// - 如果指定的会话不存在，将通过 SSE 错误帧返回 400 错误。
//
// - `stopReason`字段指示了 Agent 停止本轮对话的原因。
//
// - 可选地提供额外元信息`Meta`来传递更多上下文给服务端。
//
// - 返回的内容符合开源协议 Agent Client Protocol (ACP) 的规范。
//
// @param request - PromptAgentSessionRequest
//
// @return PromptAgentSessionResponse
func (client *Client) PromptAgentSession(request *PromptAgentSessionRequest) (_result *PromptAgentSessionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_result = &PromptAgentSessionResponse{}
	_body, _err := client.PromptAgentSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) promptAgentSessionWithSSE_opYieldFunc(_yield chan *PromptAgentSessionResponse, _yieldErr chan error, tmpReq *PromptAgentSessionRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &PromptAgentSessionShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Params) {
		request.ParamsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Params, dara.String("Params"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.CallerContext) {
		body["Caller-Context"] = request.CallerContext
	}

	if !dara.IsNil(request.Id) {
		body["Id"] = request.Id
	}

	if !dara.IsNil(request.Jsonrpc) {
		body["Jsonrpc"] = request.Jsonrpc
	}

	if !dara.IsNil(request.ParamsShrink) {
		body["Params"] = request.ParamsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PromptAgentSession"),
		Version:     dara.String("2022-06-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		if !dara.IsNil(resp.Event) && !dara.IsNil(resp.Event.Data) {
			data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
			_err := dara.ConvertChan(map[string]interface{}{
				"statusCode": dara.IntValue(resp.StatusCode),
				"headers":    resp.Headers,
				"id":         dara.StringValue(resp.Event.Id),
				"event":      dara.StringValue(resp.Event.Event),
				"body":       data,
			}, _yield)
			if _err != nil {
				_yieldErr <- _err
				return
			}
		}

	}
}
