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
	client.Endpoint, _err = client.GetEndpoint(dara.String("maasqiservice"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 以 SSE 流式调用已发布 AIGC Agent；支持 AIGCLite 文生图/图生图与 AIGCStandard Planner，兼容 OpenAI Chat Completions。
//
// @param request - AigcChatCompletionStreamRequest
//
// @param headers - AigcChatCompletionStreamHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AigcChatCompletionStreamResponse
func (client *Client) AigcChatCompletionStreamWithSSE(request *AigcChatCompletionStreamRequest, headers *AigcChatCompletionStreamHeaders, runtime *dara.RuntimeOptions, _yield chan *AigcChatCompletionStreamResponse, _yieldErr chan error) {
	defer close(_yield)
	client.aigcChatCompletionStreamWithSSE_opYieldFunc(_yield, _yieldErr, request, headers, runtime)
	return
}

// Summary:
//
// 以 SSE 流式调用已发布 AIGC Agent；支持 AIGCLite 文生图/图生图与 AIGCStandard Planner，兼容 OpenAI Chat Completions。
//
// @param request - AigcChatCompletionStreamRequest
//
// @param headers - AigcChatCompletionStreamHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AigcChatCompletionStreamResponse
func (client *Client) AigcChatCompletionStreamWithOptions(request *AigcChatCompletionStreamRequest, headers *AigcChatCompletionStreamHeaders, runtime *dara.RuntimeOptions) (_result *AigcChatCompletionStreamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Messages) {
		body["messages"] = request.Messages
	}

	if !dara.IsNil(request.Metadata) {
		body["metadata"] = request.Metadata
	}

	if !dara.IsNil(request.Model) {
		body["model"] = request.Model
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	if !dara.IsNil(request.StreamOptions) {
		body["streamOptions"] = request.StreamOptions
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XQIAgentApiKey) {
		realHeaders["X-QI-Agent-Api-Key"] = dara.String(dara.ToString(dara.StringValue(headers.XQIAgentApiKey)))
	}

	if !dara.IsNil(headers.XQIInstanceId) {
		realHeaders["X-QI-Instance-Id"] = dara.String(dara.ToString(dara.StringValue(headers.XQIInstanceId)))
	}

	if !dara.IsNil(headers.XQISessionId) {
		realHeaders["X-QI-Session-Id"] = dara.String(dara.ToString(dara.StringValue(headers.XQISessionId)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AigcChatCompletionStream"),
		Version:     dara.String("2026-08-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/aigc/v1/chat/completions"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AigcChatCompletionStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 以 SSE 流式调用已发布 AIGC Agent；支持 AIGCLite 文生图/图生图与 AIGCStandard Planner，兼容 OpenAI Chat Completions。
//
// @param request - AigcChatCompletionStreamRequest
//
// @return AigcChatCompletionStreamResponse
func (client *Client) AigcChatCompletionStream(request *AigcChatCompletionStreamRequest) (_result *AigcChatCompletionStreamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &AigcChatCompletionStreamHeaders{}
	_result = &AigcChatCompletionStreamResponse{}
	_body, _err := client.AigcChatCompletionStreamWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 以 SSE 流式调用已发布 GUI Agent；兼容 OpenAI Chat Completions，输入屏幕截图与任务文本，返回下一步 GUI 操作。
//
// @param request - GuiChatCompletionStreamRequest
//
// @param headers - GuiChatCompletionStreamHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GuiChatCompletionStreamResponse
func (client *Client) GuiChatCompletionStreamWithSSE(request *GuiChatCompletionStreamRequest, headers *GuiChatCompletionStreamHeaders, runtime *dara.RuntimeOptions, _yield chan *GuiChatCompletionStreamResponse, _yieldErr chan error) {
	defer close(_yield)
	client.guiChatCompletionStreamWithSSE_opYieldFunc(_yield, _yieldErr, request, headers, runtime)
	return
}

// Summary:
//
// 以 SSE 流式调用已发布 GUI Agent；兼容 OpenAI Chat Completions，输入屏幕截图与任务文本，返回下一步 GUI 操作。
//
// @param request - GuiChatCompletionStreamRequest
//
// @param headers - GuiChatCompletionStreamHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GuiChatCompletionStreamResponse
func (client *Client) GuiChatCompletionStreamWithOptions(request *GuiChatCompletionStreamRequest, headers *GuiChatCompletionStreamHeaders, runtime *dara.RuntimeOptions) (_result *GuiChatCompletionStreamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AllowedTokenIds) {
		body["allowedTokenIds"] = request.AllowedTokenIds
	}

	if !dara.IsNil(request.BadWords) {
		body["badWords"] = request.BadWords
	}

	if !dara.IsNil(request.ChatTemplateKwargs) {
		body["chatTemplateKwargs"] = request.ChatTemplateKwargs
	}

	if !dara.IsNil(request.FrequencyPenalty) {
		body["frequencyPenalty"] = request.FrequencyPenalty
	}

	if !dara.IsNil(request.IgnoreEos) {
		body["ignoreEos"] = request.IgnoreEos
	}

	if !dara.IsNil(request.IncludeReasoning) {
		body["includeReasoning"] = request.IncludeReasoning
	}

	if !dara.IsNil(request.Logprobs) {
		body["logprobs"] = request.Logprobs
	}

	if !dara.IsNil(request.MaxCompletionTokens) {
		body["maxCompletionTokens"] = request.MaxCompletionTokens
	}

	if !dara.IsNil(request.MaxTokens) {
		body["maxTokens"] = request.MaxTokens
	}

	if !dara.IsNil(request.Messages) {
		body["messages"] = request.Messages
	}

	if !dara.IsNil(request.Metadata) {
		body["metadata"] = request.Metadata
	}

	if !dara.IsNil(request.MinP) {
		body["minP"] = request.MinP
	}

	if !dara.IsNil(request.MinTokens) {
		body["minTokens"] = request.MinTokens
	}

	if !dara.IsNil(request.MmProcessorKwargs) {
		body["mmProcessorKwargs"] = request.MmProcessorKwargs
	}

	if !dara.IsNil(request.Model) {
		body["model"] = request.Model
	}

	if !dara.IsNil(request.N) {
		body["n"] = request.N
	}

	if !dara.IsNil(request.ParallelToolCalls) {
		body["parallelToolCalls"] = request.ParallelToolCalls
	}

	if !dara.IsNil(request.PresencePenalty) {
		body["presencePenalty"] = request.PresencePenalty
	}

	if !dara.IsNil(request.PromptLogprobs) {
		body["promptLogprobs"] = request.PromptLogprobs
	}

	if !dara.IsNil(request.ReasoningEffort) {
		body["reasoningEffort"] = request.ReasoningEffort
	}

	if !dara.IsNil(request.RepetitionPenalty) {
		body["repetitionPenalty"] = request.RepetitionPenalty
	}

	if !dara.IsNil(request.ResponseFormat) {
		body["responseFormat"] = request.ResponseFormat
	}

	if !dara.IsNil(request.Seed) {
		body["seed"] = request.Seed
	}

	if !dara.IsNil(request.SkipSpecialTokens) {
		body["skipSpecialTokens"] = request.SkipSpecialTokens
	}

	if !dara.IsNil(request.Stop) {
		body["stop"] = request.Stop
	}

	if !dara.IsNil(request.StopTokenIds) {
		body["stopTokenIds"] = request.StopTokenIds
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	if !dara.IsNil(request.StreamOptions) {
		body["streamOptions"] = request.StreamOptions
	}

	if !dara.IsNil(request.StructuredOutputs) {
		body["structuredOutputs"] = request.StructuredOutputs
	}

	if !dara.IsNil(request.Temperature) {
		body["temperature"] = request.Temperature
	}

	if !dara.IsNil(request.TopK) {
		body["topK"] = request.TopK
	}

	if !dara.IsNil(request.TopLogprobs) {
		body["topLogprobs"] = request.TopLogprobs
	}

	if !dara.IsNil(request.TopP) {
		body["topP"] = request.TopP
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XQIAgentApiKey) {
		realHeaders["X-QI-Agent-Api-Key"] = dara.String(dara.ToString(dara.StringValue(headers.XQIAgentApiKey)))
	}

	if !dara.IsNil(headers.XQIInstanceId) {
		realHeaders["X-QI-Instance-Id"] = dara.String(dara.ToString(dara.StringValue(headers.XQIInstanceId)))
	}

	if !dara.IsNil(headers.XQISessionId) {
		realHeaders["X-QI-Session-Id"] = dara.String(dara.ToString(dara.StringValue(headers.XQISessionId)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GuiChatCompletionStream"),
		Version:     dara.String("2026-08-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/gui/v1/chat/completions"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GuiChatCompletionStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 以 SSE 流式调用已发布 GUI Agent；兼容 OpenAI Chat Completions，输入屏幕截图与任务文本，返回下一步 GUI 操作。
//
// @param request - GuiChatCompletionStreamRequest
//
// @return GuiChatCompletionStreamResponse
func (client *Client) GuiChatCompletionStream(request *GuiChatCompletionStreamRequest) (_result *GuiChatCompletionStreamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &GuiChatCompletionStreamHeaders{}
	_result = &GuiChatCompletionStreamResponse{}
	_body, _err := client.GuiChatCompletionStreamWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 以 SSE 流式调用已发布 PA Agent；兼容 OpenAI Chat Completions，支持多轮消息、工具调用、多模态输入与思考内容。
//
// @param request - PaChatCompletionStreamRequest
//
// @param headers - PaChatCompletionStreamHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PaChatCompletionStreamResponse
func (client *Client) PaChatCompletionStreamWithSSE(request *PaChatCompletionStreamRequest, headers *PaChatCompletionStreamHeaders, runtime *dara.RuntimeOptions, _yield chan *PaChatCompletionStreamResponse, _yieldErr chan error) {
	defer close(_yield)
	client.paChatCompletionStreamWithSSE_opYieldFunc(_yield, _yieldErr, request, headers, runtime)
	return
}

// Summary:
//
// 以 SSE 流式调用已发布 PA Agent；兼容 OpenAI Chat Completions，支持多轮消息、工具调用、多模态输入与思考内容。
//
// @param request - PaChatCompletionStreamRequest
//
// @param headers - PaChatCompletionStreamHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PaChatCompletionStreamResponse
func (client *Client) PaChatCompletionStreamWithOptions(request *PaChatCompletionStreamRequest, headers *PaChatCompletionStreamHeaders, runtime *dara.RuntimeOptions) (_result *PaChatCompletionStreamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AllowedTokenIds) {
		body["allowedTokenIds"] = request.AllowedTokenIds
	}

	if !dara.IsNil(request.BadWords) {
		body["badWords"] = request.BadWords
	}

	if !dara.IsNil(request.ChatTemplateKwargs) {
		body["chatTemplateKwargs"] = request.ChatTemplateKwargs
	}

	if !dara.IsNil(request.FrequencyPenalty) {
		body["frequencyPenalty"] = request.FrequencyPenalty
	}

	if !dara.IsNil(request.IgnoreEos) {
		body["ignoreEos"] = request.IgnoreEos
	}

	if !dara.IsNil(request.IncludeReasoning) {
		body["includeReasoning"] = request.IncludeReasoning
	}

	if !dara.IsNil(request.Logprobs) {
		body["logprobs"] = request.Logprobs
	}

	if !dara.IsNil(request.MaxCompletionTokens) {
		body["maxCompletionTokens"] = request.MaxCompletionTokens
	}

	if !dara.IsNil(request.MaxTokens) {
		body["maxTokens"] = request.MaxTokens
	}

	if !dara.IsNil(request.Messages) {
		body["messages"] = request.Messages
	}

	if !dara.IsNil(request.MinP) {
		body["minP"] = request.MinP
	}

	if !dara.IsNil(request.MinTokens) {
		body["minTokens"] = request.MinTokens
	}

	if !dara.IsNil(request.MmProcessorKwargs) {
		body["mmProcessorKwargs"] = request.MmProcessorKwargs
	}

	if !dara.IsNil(request.Model) {
		body["model"] = request.Model
	}

	if !dara.IsNil(request.N) {
		body["n"] = request.N
	}

	if !dara.IsNil(request.ParallelToolCalls) {
		body["parallelToolCalls"] = request.ParallelToolCalls
	}

	if !dara.IsNil(request.PresencePenalty) {
		body["presencePenalty"] = request.PresencePenalty
	}

	if !dara.IsNil(request.PromptLogprobs) {
		body["promptLogprobs"] = request.PromptLogprobs
	}

	if !dara.IsNil(request.ReasoningEffort) {
		body["reasoningEffort"] = request.ReasoningEffort
	}

	if !dara.IsNil(request.RepetitionPenalty) {
		body["repetitionPenalty"] = request.RepetitionPenalty
	}

	if !dara.IsNil(request.ResponseFormat) {
		body["responseFormat"] = request.ResponseFormat
	}

	if !dara.IsNil(request.Seed) {
		body["seed"] = request.Seed
	}

	if !dara.IsNil(request.SkipSpecialTokens) {
		body["skipSpecialTokens"] = request.SkipSpecialTokens
	}

	if !dara.IsNil(request.Stop) {
		body["stop"] = request.Stop
	}

	if !dara.IsNil(request.StopTokenIds) {
		body["stopTokenIds"] = request.StopTokenIds
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	if !dara.IsNil(request.StreamOptions) {
		body["streamOptions"] = request.StreamOptions
	}

	if !dara.IsNil(request.StructuredOutputs) {
		body["structuredOutputs"] = request.StructuredOutputs
	}

	if !dara.IsNil(request.Temperature) {
		body["temperature"] = request.Temperature
	}

	if !dara.IsNil(request.ToolChoice) {
		body["toolChoice"] = request.ToolChoice
	}

	if !dara.IsNil(request.Tools) {
		body["tools"] = request.Tools
	}

	if !dara.IsNil(request.TopK) {
		body["topK"] = request.TopK
	}

	if !dara.IsNil(request.TopLogprobs) {
		body["topLogprobs"] = request.TopLogprobs
	}

	if !dara.IsNil(request.TopP) {
		body["topP"] = request.TopP
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XQIAgentApiKey) {
		realHeaders["X-QI-Agent-Api-Key"] = dara.String(dara.ToString(dara.StringValue(headers.XQIAgentApiKey)))
	}

	if !dara.IsNil(headers.XQIInstanceId) {
		realHeaders["X-QI-Instance-Id"] = dara.String(dara.ToString(dara.StringValue(headers.XQIInstanceId)))
	}

	if !dara.IsNil(headers.XQISessionId) {
		realHeaders["X-QI-Session-Id"] = dara.String(dara.ToString(dara.StringValue(headers.XQISessionId)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PaChatCompletionStream"),
		Version:     dara.String("2026-08-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pa/v1/chat/completions"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PaChatCompletionStreamResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 以 SSE 流式调用已发布 PA Agent；兼容 OpenAI Chat Completions，支持多轮消息、工具调用、多模态输入与思考内容。
//
// @param request - PaChatCompletionStreamRequest
//
// @return PaChatCompletionStreamResponse
func (client *Client) PaChatCompletionStream(request *PaChatCompletionStreamRequest) (_result *PaChatCompletionStreamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := &PaChatCompletionStreamHeaders{}
	_result = &PaChatCompletionStreamResponse{}
	_body, _err := client.PaChatCompletionStreamWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) aigcChatCompletionStreamWithSSE_opYieldFunc(_yield chan *AigcChatCompletionStreamResponse, _yieldErr chan error, request *AigcChatCompletionStreamRequest, headers *AigcChatCompletionStreamHeaders, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Messages) {
		body["messages"] = request.Messages
	}

	if !dara.IsNil(request.Metadata) {
		body["metadata"] = request.Metadata
	}

	if !dara.IsNil(request.Model) {
		body["model"] = request.Model
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	if !dara.IsNil(request.StreamOptions) {
		body["streamOptions"] = request.StreamOptions
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XQIAgentApiKey) {
		realHeaders["X-QI-Agent-Api-Key"] = dara.String(dara.ToString(dara.StringValue(headers.XQIAgentApiKey)))
	}

	if !dara.IsNil(headers.XQIInstanceId) {
		realHeaders["X-QI-Instance-Id"] = dara.String(dara.ToString(dara.StringValue(headers.XQIInstanceId)))
	}

	if !dara.IsNil(headers.XQISessionId) {
		realHeaders["X-QI-Session-Id"] = dara.String(dara.ToString(dara.StringValue(headers.XQISessionId)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AigcChatCompletionStream"),
		Version:     dara.String("2026-08-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/aigc/v1/chat/completions"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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

func (client *Client) guiChatCompletionStreamWithSSE_opYieldFunc(_yield chan *GuiChatCompletionStreamResponse, _yieldErr chan error, request *GuiChatCompletionStreamRequest, headers *GuiChatCompletionStreamHeaders, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AllowedTokenIds) {
		body["allowedTokenIds"] = request.AllowedTokenIds
	}

	if !dara.IsNil(request.BadWords) {
		body["badWords"] = request.BadWords
	}

	if !dara.IsNil(request.ChatTemplateKwargs) {
		body["chatTemplateKwargs"] = request.ChatTemplateKwargs
	}

	if !dara.IsNil(request.FrequencyPenalty) {
		body["frequencyPenalty"] = request.FrequencyPenalty
	}

	if !dara.IsNil(request.IgnoreEos) {
		body["ignoreEos"] = request.IgnoreEos
	}

	if !dara.IsNil(request.IncludeReasoning) {
		body["includeReasoning"] = request.IncludeReasoning
	}

	if !dara.IsNil(request.Logprobs) {
		body["logprobs"] = request.Logprobs
	}

	if !dara.IsNil(request.MaxCompletionTokens) {
		body["maxCompletionTokens"] = request.MaxCompletionTokens
	}

	if !dara.IsNil(request.MaxTokens) {
		body["maxTokens"] = request.MaxTokens
	}

	if !dara.IsNil(request.Messages) {
		body["messages"] = request.Messages
	}

	if !dara.IsNil(request.Metadata) {
		body["metadata"] = request.Metadata
	}

	if !dara.IsNil(request.MinP) {
		body["minP"] = request.MinP
	}

	if !dara.IsNil(request.MinTokens) {
		body["minTokens"] = request.MinTokens
	}

	if !dara.IsNil(request.MmProcessorKwargs) {
		body["mmProcessorKwargs"] = request.MmProcessorKwargs
	}

	if !dara.IsNil(request.Model) {
		body["model"] = request.Model
	}

	if !dara.IsNil(request.N) {
		body["n"] = request.N
	}

	if !dara.IsNil(request.ParallelToolCalls) {
		body["parallelToolCalls"] = request.ParallelToolCalls
	}

	if !dara.IsNil(request.PresencePenalty) {
		body["presencePenalty"] = request.PresencePenalty
	}

	if !dara.IsNil(request.PromptLogprobs) {
		body["promptLogprobs"] = request.PromptLogprobs
	}

	if !dara.IsNil(request.ReasoningEffort) {
		body["reasoningEffort"] = request.ReasoningEffort
	}

	if !dara.IsNil(request.RepetitionPenalty) {
		body["repetitionPenalty"] = request.RepetitionPenalty
	}

	if !dara.IsNil(request.ResponseFormat) {
		body["responseFormat"] = request.ResponseFormat
	}

	if !dara.IsNil(request.Seed) {
		body["seed"] = request.Seed
	}

	if !dara.IsNil(request.SkipSpecialTokens) {
		body["skipSpecialTokens"] = request.SkipSpecialTokens
	}

	if !dara.IsNil(request.Stop) {
		body["stop"] = request.Stop
	}

	if !dara.IsNil(request.StopTokenIds) {
		body["stopTokenIds"] = request.StopTokenIds
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	if !dara.IsNil(request.StreamOptions) {
		body["streamOptions"] = request.StreamOptions
	}

	if !dara.IsNil(request.StructuredOutputs) {
		body["structuredOutputs"] = request.StructuredOutputs
	}

	if !dara.IsNil(request.Temperature) {
		body["temperature"] = request.Temperature
	}

	if !dara.IsNil(request.TopK) {
		body["topK"] = request.TopK
	}

	if !dara.IsNil(request.TopLogprobs) {
		body["topLogprobs"] = request.TopLogprobs
	}

	if !dara.IsNil(request.TopP) {
		body["topP"] = request.TopP
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XQIAgentApiKey) {
		realHeaders["X-QI-Agent-Api-Key"] = dara.String(dara.ToString(dara.StringValue(headers.XQIAgentApiKey)))
	}

	if !dara.IsNil(headers.XQIInstanceId) {
		realHeaders["X-QI-Instance-Id"] = dara.String(dara.ToString(dara.StringValue(headers.XQIInstanceId)))
	}

	if !dara.IsNil(headers.XQISessionId) {
		realHeaders["X-QI-Session-Id"] = dara.String(dara.ToString(dara.StringValue(headers.XQISessionId)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GuiChatCompletionStream"),
		Version:     dara.String("2026-08-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/gui/v1/chat/completions"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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

func (client *Client) paChatCompletionStreamWithSSE_opYieldFunc(_yield chan *PaChatCompletionStreamResponse, _yieldErr chan error, request *PaChatCompletionStreamRequest, headers *PaChatCompletionStreamHeaders, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AllowedTokenIds) {
		body["allowedTokenIds"] = request.AllowedTokenIds
	}

	if !dara.IsNil(request.BadWords) {
		body["badWords"] = request.BadWords
	}

	if !dara.IsNil(request.ChatTemplateKwargs) {
		body["chatTemplateKwargs"] = request.ChatTemplateKwargs
	}

	if !dara.IsNil(request.FrequencyPenalty) {
		body["frequencyPenalty"] = request.FrequencyPenalty
	}

	if !dara.IsNil(request.IgnoreEos) {
		body["ignoreEos"] = request.IgnoreEos
	}

	if !dara.IsNil(request.IncludeReasoning) {
		body["includeReasoning"] = request.IncludeReasoning
	}

	if !dara.IsNil(request.Logprobs) {
		body["logprobs"] = request.Logprobs
	}

	if !dara.IsNil(request.MaxCompletionTokens) {
		body["maxCompletionTokens"] = request.MaxCompletionTokens
	}

	if !dara.IsNil(request.MaxTokens) {
		body["maxTokens"] = request.MaxTokens
	}

	if !dara.IsNil(request.Messages) {
		body["messages"] = request.Messages
	}

	if !dara.IsNil(request.MinP) {
		body["minP"] = request.MinP
	}

	if !dara.IsNil(request.MinTokens) {
		body["minTokens"] = request.MinTokens
	}

	if !dara.IsNil(request.MmProcessorKwargs) {
		body["mmProcessorKwargs"] = request.MmProcessorKwargs
	}

	if !dara.IsNil(request.Model) {
		body["model"] = request.Model
	}

	if !dara.IsNil(request.N) {
		body["n"] = request.N
	}

	if !dara.IsNil(request.ParallelToolCalls) {
		body["parallelToolCalls"] = request.ParallelToolCalls
	}

	if !dara.IsNil(request.PresencePenalty) {
		body["presencePenalty"] = request.PresencePenalty
	}

	if !dara.IsNil(request.PromptLogprobs) {
		body["promptLogprobs"] = request.PromptLogprobs
	}

	if !dara.IsNil(request.ReasoningEffort) {
		body["reasoningEffort"] = request.ReasoningEffort
	}

	if !dara.IsNil(request.RepetitionPenalty) {
		body["repetitionPenalty"] = request.RepetitionPenalty
	}

	if !dara.IsNil(request.ResponseFormat) {
		body["responseFormat"] = request.ResponseFormat
	}

	if !dara.IsNil(request.Seed) {
		body["seed"] = request.Seed
	}

	if !dara.IsNil(request.SkipSpecialTokens) {
		body["skipSpecialTokens"] = request.SkipSpecialTokens
	}

	if !dara.IsNil(request.Stop) {
		body["stop"] = request.Stop
	}

	if !dara.IsNil(request.StopTokenIds) {
		body["stopTokenIds"] = request.StopTokenIds
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	if !dara.IsNil(request.StreamOptions) {
		body["streamOptions"] = request.StreamOptions
	}

	if !dara.IsNil(request.StructuredOutputs) {
		body["structuredOutputs"] = request.StructuredOutputs
	}

	if !dara.IsNil(request.Temperature) {
		body["temperature"] = request.Temperature
	}

	if !dara.IsNil(request.ToolChoice) {
		body["toolChoice"] = request.ToolChoice
	}

	if !dara.IsNil(request.Tools) {
		body["tools"] = request.Tools
	}

	if !dara.IsNil(request.TopK) {
		body["topK"] = request.TopK
	}

	if !dara.IsNil(request.TopLogprobs) {
		body["topLogprobs"] = request.TopLogprobs
	}

	if !dara.IsNil(request.TopP) {
		body["topP"] = request.TopP
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XQIAgentApiKey) {
		realHeaders["X-QI-Agent-Api-Key"] = dara.String(dara.ToString(dara.StringValue(headers.XQIAgentApiKey)))
	}

	if !dara.IsNil(headers.XQIInstanceId) {
		realHeaders["X-QI-Instance-Id"] = dara.String(dara.ToString(dara.StringValue(headers.XQIInstanceId)))
	}

	if !dara.IsNil(headers.XQISessionId) {
		realHeaders["X-QI-Session-Id"] = dara.String(dara.ToString(dara.StringValue(headers.XQISessionId)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PaChatCompletionStream"),
		Version:     dara.String("2026-08-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/pa/v1/chat/completions"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
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
