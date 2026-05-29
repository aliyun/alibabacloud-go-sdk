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
// 与 JVS Crew 进行流式对话，采用 Server-Sent Events (SSE) 协议实时推送对话内容。
//
// @param tmpReq - ChatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatResponse
func (client *Client) ChatWithSSECtx(ctx context.Context, tmpReq *ChatRequest, runtime *dara.RuntimeOptions, _yield chan *ChatResponse, _yieldErr chan error) {
	defer close(_yield)
	client.chatWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, tmpReq, runtime)
	return
}

// Summary:
//
// 与 JVS Crew 进行流式对话，采用 Server-Sent Events (SSE) 协议实时推送对话内容。
//
// @param tmpReq - ChatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatResponse
func (client *Client) ChatWithContext(ctx context.Context, tmpReq *ChatRequest, runtime *dara.RuntimeOptions) (_result *ChatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ChatShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Input) {
		request.InputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Input, dara.String("Input"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Settings) {
		request.SettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Settings, dara.String("Settings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.StreamOptions) {
		request.StreamOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StreamOptions, dara.String("StreamOptions"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Authorization) {
		query["Authorization"] = request.Authorization
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ExternalUserId) {
		body["ExternalUserId"] = request.ExternalUserId
	}

	if !dara.IsNil(request.InputShrink) {
		body["Input"] = request.InputShrink
	}

	if !dara.IsNil(request.RoutingKey) {
		body["RoutingKey"] = request.RoutingKey
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.SettingsShrink) {
		body["Settings"] = request.SettingsShrink
	}

	if !dara.IsNil(request.StreamOptionsShrink) {
		body["StreamOptions"] = request.StreamOptionsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Chat"),
		Version:     dara.String("2026-03-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &ChatResponse{}
	_body, _err := client.DoRPCRequestWithCtx(ctx, params.Action, params.Version, params.Protocol, params.Method, params.AuthType, params.BodyType, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户进行对话所需的访问令牌（AccessToken），用于后续调用 Chat 接口进行身份验证。
//
// @param request - GetAccessTokenRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAccessTokenResponse
func (client *Client) GetAccessTokenWithContext(ctx context.Context, request *GetAccessTokenRequest, runtime *dara.RuntimeOptions) (_result *GetAccessTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExternalUserId) {
		query["ExternalUserId"] = request.ExternalUserId
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAccessToken"),
		Version:     dara.String("2026-03-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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

func (client *Client) chatWithSSECtx_opYieldFunc(_yield chan *ChatResponse, _yieldErr chan error, ctx context.Context, tmpReq *ChatRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &ChatShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Input) {
		request.InputShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Input, dara.String("Input"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Settings) {
		request.SettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Settings, dara.String("Settings"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.StreamOptions) {
		request.StreamOptionsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StreamOptions, dara.String("StreamOptions"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Authorization) {
		query["Authorization"] = request.Authorization
	}

	if !dara.IsNil(request.TemplateId) {
		query["TemplateId"] = request.TemplateId
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ExternalUserId) {
		body["ExternalUserId"] = request.ExternalUserId
	}

	if !dara.IsNil(request.InputShrink) {
		body["Input"] = request.InputShrink
	}

	if !dara.IsNil(request.RoutingKey) {
		body["RoutingKey"] = request.RoutingKey
	}

	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.SettingsShrink) {
		body["Settings"] = request.SettingsShrink
	}

	if !dara.IsNil(request.StreamOptionsShrink) {
		body["StreamOptions"] = request.StreamOptionsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Chat"),
		Version:     dara.String("2026-03-11"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("Anonymous"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
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
