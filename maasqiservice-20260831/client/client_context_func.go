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
// Calls a published Creative Agent with support for Lite and Dedicated editions, and returns execution progress and image artifacts by using Server-Sent Events (SSE). The Lite edition performs text-to-image or image-to-image generation depending on whether a reference image is provided. The Dedicated edition uses a Planner to plan and execute generation and editing tasks with multi-turn conversation support. Pass generation count, dimensions, and other parameters through metadata.parameters. Image URLs are included in the artifacts of the result.
//
// Description:
//
// ### Before you begin
//
// - Install the SDK from the SDK Center.
//
// - Activate the Creative Agent service and prepare the published instance ID and the corresponding instance access key.
//
// - Creative Agent provides Lite and Dedicated editions. The Lite edition generates or edits images based on whether a reference image is provided. The Dedicated edition uses a Planner to plan and execute creative tasks.
//
// - Complete AccessKey authentication by using the SDK. The instance access key is placed in X-QI-Agent-Api-Key and cannot replace the AccessKey signature.
//
// - Set Content-Type: application/json and Accept: text/event-stream. The stream parameter must be set to true.
//
// ### How to call
//
// - Method and path: `POST /aigc/v1/chat/completions`.
//
// - Specify the instance by using `X-QI-Instance-Id`. The request is executed based on the published configuration of the specified instance. The model field in the request body is a compatibility placeholder and cannot be used to switch models.
//
// - This operation returns incremental events instead of a complete JSON response. Continue reading until the stream ends and check for in-stream errors.
//
// ### Session management
//
// You can omit `X-QI-Session-Id` for the first turn. Save the session identifier returned by the server from the `qi_session_id` field in the SSE body, and pass the same value back in subsequent turns. Do not rely on the gateway to pass through custom response headers. Use different identifiers for different sessions and send requests sequentially within the same session.
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
// Calls a published Creative Agent with support for Lite and Dedicated editions, and returns execution progress and image artifacts by using Server-Sent Events (SSE). The Lite edition performs text-to-image or image-to-image generation depending on whether a reference image is provided. The Dedicated edition uses a Planner to plan and execute generation and editing tasks with multi-turn conversation support. Pass generation count, dimensions, and other parameters through metadata.parameters. Image URLs are included in the artifacts of the result.
//
// Description:
//
// ### Before you begin
//
// - Install the SDK from the SDK Center.
//
// - Activate the Creative Agent service and prepare the published instance ID and the corresponding instance access key.
//
// - Creative Agent provides Lite and Dedicated editions. The Lite edition generates or edits images based on whether a reference image is provided. The Dedicated edition uses a Planner to plan and execute creative tasks.
//
// - Complete AccessKey authentication by using the SDK. The instance access key is placed in X-QI-Agent-Api-Key and cannot replace the AccessKey signature.
//
// - Set Content-Type: application/json and Accept: text/event-stream. The stream parameter must be set to true.
//
// ### How to call
//
// - Method and path: `POST /aigc/v1/chat/completions`.
//
// - Specify the instance by using `X-QI-Instance-Id`. The request is executed based on the published configuration of the specified instance. The model field in the request body is a compatibility placeholder and cannot be used to switch models.
//
// - This operation returns incremental events instead of a complete JSON response. Continue reading until the stream ends and check for in-stream errors.
//
// ### Session management
//
// You can omit `X-QI-Session-Id` for the first turn. Save the session identifier returned by the server from the `qi_session_id` field in the SSE body, and pass the same value back in subsequent turns. Do not rely on the gateway to pass through custom response headers. Use different identifiers for different sessions and send requests sequentially within the same session.
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
// 调用已发布 CUAStandard 实例，根据任务和桌面截图生成电脑操作单步决策，以 SSE 返回 computer_use 工具调用
//
// @param request - CuaChatCompletionStreamRequest
//
// @param headers - CuaChatCompletionStreamHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CuaChatCompletionStreamResponse
func (client *Client) CuaChatCompletionStreamWithSSECtx(ctx context.Context, request *CuaChatCompletionStreamRequest, headers *CuaChatCompletionStreamHeaders, runtime *dara.RuntimeOptions, _yield chan *CuaChatCompletionStreamResponse, _yieldErr chan error) {
	defer close(_yield)
	client.cuaChatCompletionStreamWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, headers, runtime)
	return
}

// Summary:
//
// 调用已发布 CUAStandard 实例，根据任务和桌面截图生成电脑操作单步决策，以 SSE 返回 computer_use 工具调用
//
// @param request - CuaChatCompletionStreamRequest
//
// @param headers - CuaChatCompletionStreamHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CuaChatCompletionStreamResponse
func (client *Client) CuaChatCompletionStreamWithContext(ctx context.Context, request *CuaChatCompletionStreamRequest, headers *CuaChatCompletionStreamHeaders, runtime *dara.RuntimeOptions) (_result *CuaChatCompletionStreamResponse, _err error) {
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

	if !dara.IsNil(request.Tools) {
		body["tools"] = request.Tools
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
		Action:      dara.String("CuaChatCompletionStream"),
		Version:     dara.String("2026-08-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/cua/v1/chat/completions"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CuaChatCompletionStreamResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls a published Mobile-Use Agent instance, which supports the Standard Edition and Dedicated Edition. Submits the task text and the current screen screenshot, passes the device and application context through metadata, and returns the next mobile UI operation by using Server-Sent Events (SSE). The caller executes the UI operation and then continues to call the API with the session ID and the latest screenshot.
//
// Description:
//
// ### Prerequisites
//
// - Install the SDK from the SDK Center.
//
// - Activate the Mobile-Use Agent service and prepare the ID and access key of a published Mobile-Use Agent instance.
//
// - Mobile-Use Agent provides the Standard Edition and Dedicated Edition. Call the API based on the edition of the selected instance.
//
// - Use the SDK to complete AccessKey authentication. Place the instance access key in the X-QI-Agent-Api-Key header, which cannot replace the AccessKey signature.
//
// - Set Content-Type to application/json and Accept to text/event-stream. The stream parameter must be set to true.
//
// ### Call methods
//
// - Method and path: POST /gui/v1/chat/completions.
//
// - Use X-QI-Instance-Id to specify the instance and execute based on the published configurations. The model field in the request body is a compatibility placeholder and cannot be used to switch models.
//
// - The response consists of SSE incremental events rather than a complete JSON object. Continuously read the stream until it ends and check for errors within the stream.
//
// ### Session management
//
// You can omit X-QI-Session-Id in the first round. Save the qi_session_id from the SSE body and pass it in the request header with the same name for subsequent calls. Make sequential calls within the same session and use different identifiers for different sessions. Do not rely on the gateway to pass through custom response headers. The actual identifier is subject to the SSE body.
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
// Calls a published Mobile-Use Agent instance, which supports the Standard Edition and Dedicated Edition. Submits the task text and the current screen screenshot, passes the device and application context through metadata, and returns the next mobile UI operation by using Server-Sent Events (SSE). The caller executes the UI operation and then continues to call the API with the session ID and the latest screenshot.
//
// Description:
//
// ### Prerequisites
//
// - Install the SDK from the SDK Center.
//
// - Activate the Mobile-Use Agent service and prepare the ID and access key of a published Mobile-Use Agent instance.
//
// - Mobile-Use Agent provides the Standard Edition and Dedicated Edition. Call the API based on the edition of the selected instance.
//
// - Use the SDK to complete AccessKey authentication. Place the instance access key in the X-QI-Agent-Api-Key header, which cannot replace the AccessKey signature.
//
// - Set Content-Type to application/json and Accept to text/event-stream. The stream parameter must be set to true.
//
// ### Call methods
//
// - Method and path: POST /gui/v1/chat/completions.
//
// - Use X-QI-Instance-Id to specify the instance and execute based on the published configurations. The model field in the request body is a compatibility placeholder and cannot be used to switch models.
//
// - The response consists of SSE incremental events rather than a complete JSON object. Continuously read the stream until it ends and check for errors within the stream.
//
// ### Session management
//
// You can omit X-QI-Session-Id in the first round. Save the qi_session_id from the SSE body and pass it in the request header with the same name for subsequent calls. Make sequential calls within the same session and use different identifiers for different sessions. Do not rely on the gateway to pass through custom response headers. The actual identifier is subject to the SSE body.
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
// Invokes a published Mobile Planner Agent using OpenAI-Compatible Chat Completions messages to initiate a conversation or node planning, and returns text, model-provided reasoning content, and structured tool calling via SSE. Supports multi-turn sessions and model-supported image input. Business tools are executed by the caller.
//
// Description:
//
// ### Before you begin
//
// - Install the SDK from the SDK center.
//
// - Activate the Mobile Planner Agent service and prepare the published Mobile Planner Agent instance ID and its access key.
//
// - Use the SDK to complete AccessKey authentication. The instance access key is placed in X-QI-Agent-Api-Key and cannot replace the AccessKey signature.
//
// - Set Content-Type: application/json and Accept: text/event-stream. The stream parameter must be set to true.
//
// ### How to call
//
// - Method and path: `POST /pa/v1/chat/completions`.
//
// - X-QI-Instance-Id specifies the instance and executes based on the published configuration. The model field in the request body is a compatibility placeholder and cannot switch models.
//
// - The response consists of SSE incremental events, not a single complete JSON object. Continue reading until the stream ends and check for errors within the stream.
//
// ### Session management
//
// You can omit X-QI-Session-Id for the first turn. Save qi_session_id from the SSE body and pass it back through the request header of the same name in subsequent turns. Call sequentially within the same session. Use different identifiers for different sessions. Do not rely on the gateway to pass through custom response headers. The actual identifier is based on the SSE body.
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
// Invokes a published Mobile Planner Agent using OpenAI-Compatible Chat Completions messages to initiate a conversation or node planning, and returns text, model-provided reasoning content, and structured tool calling via SSE. Supports multi-turn sessions and model-supported image input. Business tools are executed by the caller.
//
// Description:
//
// ### Before you begin
//
// - Install the SDK from the SDK center.
//
// - Activate the Mobile Planner Agent service and prepare the published Mobile Planner Agent instance ID and its access key.
//
// - Use the SDK to complete AccessKey authentication. The instance access key is placed in X-QI-Agent-Api-Key and cannot replace the AccessKey signature.
//
// - Set Content-Type: application/json and Accept: text/event-stream. The stream parameter must be set to true.
//
// ### How to call
//
// - Method and path: `POST /pa/v1/chat/completions`.
//
// - X-QI-Instance-Id specifies the instance and executes based on the published configuration. The model field in the request body is a compatibility placeholder and cannot switch models.
//
// - The response consists of SSE incremental events, not a single complete JSON object. Continue reading until the stream ends and check for errors within the stream.
//
// ### Session management
//
// You can omit X-QI-Session-Id for the first turn. Save qi_session_id from the SSE body and pass it back through the request header of the same name in subsequent turns. Call sequentially within the same session. Use different identifiers for different sessions. Do not rely on the gateway to pass through custom response headers. The actual identifier is based on the SSE body.
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

func (client *Client) cuaChatCompletionStreamWithSSECtx_opYieldFunc(_yield chan *CuaChatCompletionStreamResponse, _yieldErr chan error, ctx context.Context, request *CuaChatCompletionStreamRequest, headers *CuaChatCompletionStreamHeaders, runtime *dara.RuntimeOptions) {
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

	if !dara.IsNil(request.Tools) {
		body["tools"] = request.Tools
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
		Action:      dara.String("CuaChatCompletionStream"),
		Version:     dara.String("2026-08-31"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/cua/v1/chat/completions"),
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
