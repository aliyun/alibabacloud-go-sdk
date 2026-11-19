// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	websocketutils "github.com/alibabacloud-go/darabonba-openapi/v2/websocketUtils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// Sends text or multimodal messages to initiate conversations with web search capabilities, supporting deep thinking, session history management, and both complete JSON and Server-Sent Events (SSE) streaming responses. Specific agent and media capabilities are subject to the activated scope.
//
// @param request - ChatRequest
//
// @param headers - ChatHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatResponse
func (client *Client) ChatWithSSECtx(ctx context.Context, request *ChatRequest, headers *ChatHeaders, runtime *dara.RuntimeOptions, _yield chan *ChatResponse, _yieldErr chan error) {
	defer close(_yield)
	client.chatWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, headers, runtime)
	return
}

// Summary:
//
// Sends text or multimodal messages to initiate conversations with web search capabilities, supporting deep thinking, session history management, and both complete JSON and Server-Sent Events (SSE) streaming responses. Specific agent and media capabilities are subject to the activated scope.
//
// @param request - ChatRequest
//
// @param headers - ChatHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatResponse
func (client *Client) ChatWithContext(ctx context.Context, request *ChatRequest, headers *ChatHeaders, runtime *dara.RuntimeOptions) (_result *ChatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Communication) {
		body["communication"] = request.Communication
	}

	if !dara.IsNil(request.Data) {
		body["data"] = request.Data
	}

	if !dara.IsNil(request.Params) {
		body["params"] = request.Params
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Accept) {
		realHeaders["Accept"] = dara.String(dara.ToString(dara.StringValue(headers.Accept)))
	}

	if !dara.IsNil(headers.XDashScopeSSE) {
		realHeaders["X-DashScope-SSE"] = dara.String(dara.ToString(dara.StringValue(headers.XDashScopeSSE)))
	}

	if !dara.IsNil(headers.XGWAccess) {
		realHeaders["X-GW-Access"] = dara.String(dara.ToString(dara.StringValue(headers.XGWAccess)))
	}

	if !dara.IsNil(headers.XGWNonce) {
		realHeaders["X-GW-Nonce"] = dara.String(dara.ToString(dara.StringValue(headers.XGWNonce)))
	}

	if !dara.IsNil(headers.XGWProtocol) {
		realHeaders["X-GW-Protocol"] = dara.String(dara.ToString(dara.StringValue(headers.XGWProtocol)))
	}

	if !dara.IsNil(headers.XGWSignature) {
		realHeaders["X-GW-Signature"] = dara.String(dara.ToString(dara.StringValue(headers.XGWSignature)))
	}

	if !dara.IsNil(headers.XGWTimestamp) {
		realHeaders["X-GW-Timestamp"] = dara.String(dara.ToString(dara.StringValue(headers.XGWTimestamp)))
	}

	if !dara.IsNil(headers.XRequestId) {
		realHeaders["X-Request-Id"] = dara.String(dara.ToString(dara.StringValue(headers.XRequestId)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Chat"),
		Version:     dara.String("2026-11-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/api/chat"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ChatResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Validates the service activation API or initiates service activation for the current account, and returns the current activation status along with the available order or instance ID.
//
// @param request - OpenServiceRequest
//
// @param headers - OpenServiceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OpenServiceResponse
func (client *Client) OpenServiceWithContext(ctx context.Context, request *OpenServiceRequest, headers *OpenServiceHeaders, runtime *dara.RuntimeOptions) (_result *OpenServiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Communication) {
		body["communication"] = request.Communication
	}

	if !dara.IsNil(request.OpenBenefits) {
		body["openBenefits"] = request.OpenBenefits
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XGWAccess) {
		realHeaders["X-GW-Access"] = dara.String(dara.ToString(dara.StringValue(headers.XGWAccess)))
	}

	if !dara.IsNil(headers.XGWNonce) {
		realHeaders["X-GW-Nonce"] = dara.String(dara.ToString(dara.StringValue(headers.XGWNonce)))
	}

	if !dara.IsNil(headers.XGWSignature) {
		realHeaders["X-GW-Signature"] = dara.String(dara.ToString(dara.StringValue(headers.XGWSignature)))
	}

	if !dara.IsNil(headers.XGWTimestamp) {
		realHeaders["X-GW-Timestamp"] = dara.String(dara.ToString(dara.StringValue(headers.XGWTimestamp)))
	}

	if !dara.IsNil(headers.XRequestId) {
		realHeaders["X-Request-Id"] = dara.String(dara.ToString(dara.StringValue(headers.XRequestId)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OpenService"),
		Version:     dara.String("2026-11-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/api/openService"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &OpenServiceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Uploads audio and receives recognized text, agent responses, and synthesized audio over a WebSocket connection. Supports image and text auxiliary input, recognition context, and conversation history. The operation description covers the complete handshake, control frames, audio frames, return events, and error handling protocols.
//
// @param request - StreamVoiceRequest
//
// @param headers - StreamVoiceHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StreamVoiceResponse
func (client *Client) StreamVoiceWithContext(ctx context.Context, request *StreamVoiceRequest, headers *StreamVoiceHeaders, runtime *dara.RuntimeOptions) (_result *StreamVoiceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XGWAccess) {
		realHeaders["X-GW-Access"] = dara.String(dara.ToString(dara.StringValue(headers.XGWAccess)))
	}

	if !dara.IsNil(headers.XGWNonce) {
		realHeaders["X-GW-Nonce"] = dara.String(dara.ToString(dara.StringValue(headers.XGWNonce)))
	}

	if !dara.IsNil(headers.XGWProtocol) {
		realHeaders["X-GW-Protocol"] = dara.String(dara.ToString(dara.StringValue(headers.XGWProtocol)))
	}

	if !dara.IsNil(headers.XGWSignature) {
		realHeaders["X-GW-Signature"] = dara.String(dara.ToString(dara.StringValue(headers.XGWSignature)))
	}

	if !dara.IsNil(headers.XGWTimestamp) {
		realHeaders["X-GW-Timestamp"] = dara.String(dara.ToString(dara.StringValue(headers.XGWTimestamp)))
	}

	if !dara.IsNil(headers.XRequestId) {
		realHeaders["X-Request-Id"] = dara.String(dara.ToString(dara.StringValue(headers.XRequestId)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
	}
	params := &openapiutil.Params{
		Action:               dara.String("StreamVoice"),
		Version:              dara.String("2026-11-05"),
		Protocol:             dara.String("wss"),
		Pathname:             dara.String("/v1/api/voice"),
		Method:               dara.String("GET"),
		AuthType:             dara.String("AK"),
		Style:                dara.String("ROA"),
		ReqBodyType:          dara.String("json"),
		BodyType:             dara.String("json"),
		WebsocketSubProtocol: dara.String("general"),
	}
	res := &StreamVoiceResponse{}
	callApiTmp, err := client.CallApiWithCtx(ctx, params, req, runtime)
	if err != nil {
		_err = err
		return _result, _err
	}
	tmp := dara.ToMap(callApiTmp)
	if !dara.IsNil(tmp["webSocketClient"]) {
		res.WebSocketClient = websocketutils.CreateWebSocketClient(tmp["webSocketClient"])
	}

	_result = res
	return _result, _err
}

func (client *Client) chatWithSSECtx_opYieldFunc(_yield chan *ChatResponse, _yieldErr chan error, ctx context.Context, request *ChatRequest, headers *ChatHeaders, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Communication) {
		body["communication"] = request.Communication
	}

	if !dara.IsNil(request.Data) {
		body["data"] = request.Data
	}

	if !dara.IsNil(request.Params) {
		body["params"] = request.Params
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.Accept) {
		realHeaders["Accept"] = dara.String(dara.ToString(dara.StringValue(headers.Accept)))
	}

	if !dara.IsNil(headers.XDashScopeSSE) {
		realHeaders["X-DashScope-SSE"] = dara.String(dara.ToString(dara.StringValue(headers.XDashScopeSSE)))
	}

	if !dara.IsNil(headers.XGWAccess) {
		realHeaders["X-GW-Access"] = dara.String(dara.ToString(dara.StringValue(headers.XGWAccess)))
	}

	if !dara.IsNil(headers.XGWNonce) {
		realHeaders["X-GW-Nonce"] = dara.String(dara.ToString(dara.StringValue(headers.XGWNonce)))
	}

	if !dara.IsNil(headers.XGWProtocol) {
		realHeaders["X-GW-Protocol"] = dara.String(dara.ToString(dara.StringValue(headers.XGWProtocol)))
	}

	if !dara.IsNil(headers.XGWSignature) {
		realHeaders["X-GW-Signature"] = dara.String(dara.ToString(dara.StringValue(headers.XGWSignature)))
	}

	if !dara.IsNil(headers.XGWTimestamp) {
		realHeaders["X-GW-Timestamp"] = dara.String(dara.ToString(dara.StringValue(headers.XGWTimestamp)))
	}

	if !dara.IsNil(headers.XRequestId) {
		realHeaders["X-Request-Id"] = dara.String(dara.ToString(dara.StringValue(headers.XRequestId)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Chat"),
		Version:     dara.String("2026-11-05"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/api/chat"),
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
