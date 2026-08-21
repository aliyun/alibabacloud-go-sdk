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
// # A2A接口
//
// @param request - A2aRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return A2aResponse
func (client *Client) A2aWithSSECtx(ctx context.Context, request *A2aRequest, runtime *dara.RuntimeOptions, _yield chan *A2aResponse, _yieldErr chan error) {
	defer close(_yield)
	client.a2aWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, runtime)
	return
}

// Summary:
//
// # A2A接口
//
// @param request - A2aRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return A2aResponse
func (client *Client) A2aWithContext(ctx context.Context, request *A2aRequest, runtime *dara.RuntimeOptions) (_result *A2aResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["id"] = request.Id
	}

	if !dara.IsNil(request.Jsonrpc) {
		body["jsonrpc"] = request.Jsonrpc
	}

	if !dara.IsNil(request.Method) {
		body["method"] = request.Method
	}

	if !dara.IsNil(request.Params) {
		body["params"] = request.Params
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("a2a"),
		Version:     dara.String("2026-08-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("any"),
	}
	_result = &A2aResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取Agent_Card
//
// @param request - AgentCardRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AgentCardResponse
func (client *Client) Agent_cardWithContext(ctx context.Context, request *AgentCardRequest, runtime *dara.RuntimeOptions) (_result *AgentCardResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := openapiutil.Query(dara.ToMap(request))
	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("agent_card"),
		Version:     dara.String("2026-08-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("any"),
	}
	_result = &AgentCardResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) a2aWithSSECtx_opYieldFunc(_yield chan *A2aResponse, _yieldErr chan error, ctx context.Context, request *A2aRequest, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Env) {
		query["Env"] = request.Env
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["id"] = request.Id
	}

	if !dara.IsNil(request.Jsonrpc) {
		body["jsonrpc"] = request.Jsonrpc
	}

	if !dara.IsNil(request.Method) {
		body["method"] = request.Method
	}

	if !dara.IsNil(request.Params) {
		body["params"] = request.Params
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("a2a"),
		Version:     dara.String("2026-08-10"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("any"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		if !dara.IsNil(resp.Event) && !dara.IsNil(resp.Event.Data) {
			data := dara.StringValue(resp.Event.Data)
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
