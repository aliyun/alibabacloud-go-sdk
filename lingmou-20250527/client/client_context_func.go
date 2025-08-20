// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 关闭会话实例session
//
// @param tmpReq - CloseChatInstanceSessionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseChatInstanceSessionsResponse
func (client *Client) CloseChatInstanceSessionsWithContext(ctx context.Context, instanceId *string, tmpReq *CloseChatInstanceSessionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CloseChatInstanceSessionsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CloseChatInstanceSessionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SessionIds) {
		request.SessionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SessionIds, dara.String("sessionIds"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.SessionIdsShrink) {
		body["sessionIds"] = request.SessionIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloseChatInstanceSessions"),
		Version:     dara.String("2025-05-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/chat/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/sessions/close"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CloseChatInstanceSessionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建数字人会话
//
// @param request - CreateChatSessionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateChatSessionResponse
func (client *Client) CreateChatSessionWithContext(ctx context.Context, id *string, request *CreateChatSessionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateChatSessionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.InstanceId) {
		query["instanceId"] = request.InstanceId
	}

	if !dara.IsNil(request.License) {
		query["license"] = request.License
	}

	if !dara.IsNil(request.Platform) {
		query["platform"] = request.Platform
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateChatSession"),
		Version:     dara.String("2025-05-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/chat/init/" + dara.PercentEncode(dara.StringValue(id))),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateChatSessionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询会话实例session
//
// @param tmpReq - QueryChatInstanceSessionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryChatInstanceSessionsResponse
func (client *Client) QueryChatInstanceSessionsWithContext(ctx context.Context, instanceId *string, tmpReq *QueryChatInstanceSessionsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryChatInstanceSessionsResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &QueryChatInstanceSessionsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.SessionIds) {
		request.SessionIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SessionIds, dara.String("sessionIds"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.SessionIdsShrink) {
		query["sessionIds"] = request.SessionIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryChatInstanceSessions"),
		Version:     dara.String("2025-05-27"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/chat/instances/" + dara.PercentEncode(dara.StringValue(instanceId)) + "/sessions"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryChatInstanceSessionsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
