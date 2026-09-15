// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

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
func (client *Client) AigcChatCompletionStreamWithSSECtx(ctx context.Context, request *AigcChatCompletionStreamRequest, headers *AigcChatCompletionStreamHeaders, runtime *dara.RuntimeOptions, _yield chan *AigcChatCompletionStreamResponse, _yieldErr chan error) {
	defer close(_yield)
	client.aigcChatCompletionStreamWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, headers, runtime)
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
func (client *Client) AigcChatCompletionStreamWithContext(ctx context.Context, request *AigcChatCompletionStreamRequest, headers *AigcChatCompletionStreamHeaders, runtime *dara.RuntimeOptions) (_result *AigcChatCompletionStreamResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - GuiChatCompletionStreamHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GuiChatCompletionStreamResponse
func (client *Client) GuiChatCompletionStreamWithSSECtx(ctx context.Context, request *GuiChatCompletionStreamRequest, headers *GuiChatCompletionStreamHeaders, runtime *dara.RuntimeOptions, _yield chan *GuiChatCompletionStreamResponse, _yieldErr chan error) {
	defer close(_yield)
	client.guiChatCompletionStreamWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, headers, runtime)
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
func (client *Client) GuiChatCompletionStreamWithContext(ctx context.Context, request *GuiChatCompletionStreamRequest, headers *GuiChatCompletionStreamHeaders, runtime *dara.RuntimeOptions) (_result *GuiChatCompletionStreamResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
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
// @param headers - PaChatCompletionStreamHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PaChatCompletionStreamResponse
func (client *Client) PaChatCompletionStreamWithSSECtx(ctx context.Context, request *PaChatCompletionStreamRequest, headers *PaChatCompletionStreamHeaders, runtime *dara.RuntimeOptions, _yield chan *PaChatCompletionStreamResponse, _yieldErr chan error) {
	defer close(_yield)
	client.paChatCompletionStreamWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, headers, runtime)
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
func (client *Client) PaChatCompletionStreamWithContext(ctx context.Context, request *PaChatCompletionStreamRequest, headers *PaChatCompletionStreamHeaders, runtime *dara.RuntimeOptions) (_result *PaChatCompletionStreamResponse, _err error) {
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
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) aigcChatCompletionStreamWithSSECtx_opYieldFunc(_yield chan *AigcChatCompletionStreamResponse, _yieldErr chan error, ctx context.Context, request *AigcChatCompletionStreamRequest, headers *AigcChatCompletionStreamHeaders, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) guiChatCompletionStreamWithSSECtx_opYieldFunc(_yield chan *GuiChatCompletionStreamResponse, _yieldErr chan error, ctx context.Context, request *GuiChatCompletionStreamRequest, headers *GuiChatCompletionStreamHeaders, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) paChatCompletionStreamWithSSECtx_opYieldFunc(_yield chan *PaChatCompletionStreamResponse, _yieldErr chan error, ctx context.Context, request *PaChatCompletionStreamRequest, headers *PaChatCompletionStreamHeaders, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
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
