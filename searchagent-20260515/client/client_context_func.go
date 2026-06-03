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
// search agent acp 接口
//
// @param request - ExecuteAcpCommandRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExecuteAcpCommandResponse
func (client *Client) ExecuteAcpCommandWithContext(ctx context.Context, workspaceName *string, request *ExecuteAcpCommandRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ExecuteAcpCommandResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentName) {
		body["agentName"] = request.AgentName
	}

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
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExecuteAcpCommand"),
		Version:     dara.String("2026-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/agent/acp"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExecuteAcpCommandResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 生成 acp stream
//
// @param request - GenerateAcpCompletionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateAcpCompletionResponse
func (client *Client) GenerateAcpCompletionWithSSECtx(ctx context.Context, workspaceName *string, request *GenerateAcpCompletionRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *GenerateAcpCompletionResponse, _yieldErr chan error) {
	defer close(_yield)
	client.generateAcpCompletionWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, workspaceName, request, headers, runtime)
	return
}

// Summary:
//
// 生成 acp stream
//
// @param request - GenerateAcpCompletionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateAcpCompletionResponse
func (client *Client) GenerateAcpCompletionWithContext(ctx context.Context, workspaceName *string, request *GenerateAcpCompletionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GenerateAcpCompletionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentName) {
		body["agentName"] = request.AgentName
	}

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
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateAcpCompletion"),
		Version:     dara.String("2026-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/agent/acp"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("string"),
	}
	_result = &GenerateAcpCompletionResponse{}
	_body, _err := client.ExecuteWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) generateAcpCompletionWithSSECtx_opYieldFunc(_yield chan *GenerateAcpCompletionResponse, _yieldErr chan error, ctx context.Context, workspaceName *string, request *GenerateAcpCompletionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentName) {
		body["agentName"] = request.AgentName
	}

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
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateAcpCompletion"),
		Version:     dara.String("2026-05-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v3/openapi/workspaces/" + dara.StringValue(workspaceName) + "/agent/acp"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("string"),
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
