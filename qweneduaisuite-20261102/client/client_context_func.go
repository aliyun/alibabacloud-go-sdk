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
// 开通千问教育API权限
//
// Description:
//
// 开通千问教育API权限
//
// @param request - ActivateQwenEduApiRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActivateQwenEduApiResponse
func (client *Client) ActivateQwenEduApiWithContext(ctx context.Context, request *ActivateQwenEduApiRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ActivateQwenEduApiResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ActivateQwenEduApi"),
		Version:     dara.String("2026-11-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/service/activate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ActivateQwenEduApiResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交单页作业图片，返回任务ID及初步检测结果。
//
// Description:
//
// ## 请求说明
//
// - 该接口用于提交单页作业图片进行批改，支持通过图片URL或Base64编码的图片数据。
//
// - 提交后会立即返回任务ID、矫正图及题目检测结果，但不包含具体的批改结论。
//
// - 需要通过查询接口轮询获取最终的批改结果。
//
// - 支持对图片进行透视矫正，默认开启。
//
// - 图片要求：支持多种格式（JPEG, PNG, GIF等），推荐体积不超过1MB，最大支持15MB（Base64编码后）。
//
// - 每次请求仅处理一张图片，多张图片需分别提交。
//
// - `request_id`字段建议由调用方生成并传递，以便于问题排查和对账。
//
// @param request - CreateEduHomeworkCorrectSubmitRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEduHomeworkCorrectSubmitResponse
func (client *Client) CreateEduHomeworkCorrectSubmitWithContext(ctx context.Context, request *CreateEduHomeworkCorrectSubmitRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateEduHomeworkCorrectSubmitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Dewarp) {
		body["dewarp"] = request.Dewarp
	}

	if !dara.IsNil(request.ImageBase64) {
		body["imageBase64"] = request.ImageBase64
	}

	if !dara.IsNil(request.ImageUrl) {
		body["imageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.RequestId) {
		body["requestId"] = request.RequestId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEduHomeworkCorrectSubmit"),
		Version:     dara.String("2026-11-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/homework/correct/submit"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEduHomeworkCorrectSubmitResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通过图片URL或Base64编码接收题目，返回解题过程和结果。
//
// Description:
//
// ## 请求说明
//
// - 本接口支持通过单张题目图片的URL或Base64编码进行解题请求。
//
// - 支持处理JPEG、JPG、PNG格式的图片，且图片大小不超过10MB，像素总数不超过260万。
//
// - 请求必须携带`request_id`字段，用于链路关联和对账；建议使用UUID以保证唯一性。
//
// - `Authorization`头部需包含有效的`appKey`，采用`Bearer <appKey>`格式。
//
// - 接口响应为SSE流式数据，客户端应逐事件读取并解析，不得等待连接关闭后整体读取。
//
// - 成功响应包括`message`（过程态全量解题快照）和`complete`（最终完整结果）两种事件类型。
//
// - 错误处理分为建流前与建流后两个阶段，根据错误码及`retryable`字段决定是否重试。
//
// - 客户端超时设置建议不低于3分钟，并确保代理层不会缓存SSE响应或提前关闭长连接。
//
// @param tmpReq - CreateEduProblemSolveRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEduProblemSolveResponse
func (client *Client) CreateEduProblemSolveWithSSECtx(ctx context.Context, tmpReq *CreateEduProblemSolveRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *CreateEduProblemSolveResponse, _yieldErr chan error) {
	defer close(_yield)
	client.createEduProblemSolveWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, tmpReq, headers, runtime)
	return
}

// Summary:
//
// 通过图片URL或Base64编码接收题目，返回解题过程和结果。
//
// Description:
//
// ## 请求说明
//
// - 本接口支持通过单张题目图片的URL或Base64编码进行解题请求。
//
// - 支持处理JPEG、JPG、PNG格式的图片，且图片大小不超过10MB，像素总数不超过260万。
//
// - 请求必须携带`request_id`字段，用于链路关联和对账；建议使用UUID以保证唯一性。
//
// - `Authorization`头部需包含有效的`appKey`，采用`Bearer <appKey>`格式。
//
// - 接口响应为SSE流式数据，客户端应逐事件读取并解析，不得等待连接关闭后整体读取。
//
// - 成功响应包括`message`（过程态全量解题快照）和`complete`（最终完整结果）两种事件类型。
//
// - 错误处理分为建流前与建流后两个阶段，根据错误码及`retryable`字段决定是否重试。
//
// - 客户端超时设置建议不低于3分钟，并确保代理层不会缓存SSE响应或提前关闭长连接。
//
// @param tmpReq - CreateEduProblemSolveRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEduProblemSolveResponse
func (client *Client) CreateEduProblemSolveWithContext(ctx context.Context, tmpReq *CreateEduProblemSolveRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateEduProblemSolveResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &CreateEduProblemSolveShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Query) {
		request.QueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Query, dara.String("query"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.QueryShrink) {
		body["query"] = request.QueryShrink
	}

	if !dara.IsNil(request.RequestId) {
		body["requestId"] = request.RequestId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEduProblemSolve"),
		Version:     dara.String("2026-11-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/problem/solve/stream"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateEduProblemSolveResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询批改进度、增量结果、终态结果及 Token 用量
//
// @param request - GetEduHomeworkCorrectionResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEduHomeworkCorrectionResultResponse
func (client *Client) GetEduHomeworkCorrectionResultWithContext(ctx context.Context, request *GetEduHomeworkCorrectionResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEduHomeworkCorrectionResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		query["taskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEduHomeworkCorrectionResult"),
		Version:     dara.String("2026-11-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/homework/correct/result"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEduHomeworkCorrectionResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 千问教育API开通状态查询
//
// @param request - GetQwenEduApiActivationStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQwenEduApiActivationStatusResponse
func (client *Client) GetQwenEduApiActivationStatusWithContext(ctx context.Context, request *GetQwenEduApiActivationStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetQwenEduApiActivationStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetQwenEduApiActivationStatus"),
		Version:     dara.String("2026-11-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/service/activation/status"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQwenEduApiActivationStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) createEduProblemSolveWithSSECtx_opYieldFunc(_yield chan *CreateEduProblemSolveResponse, _yieldErr chan error, ctx context.Context, tmpReq *CreateEduProblemSolveRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := tmpReq.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	request := &CreateEduProblemSolveShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Query) {
		request.QueryShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Query, dara.String("query"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.QueryShrink) {
		body["query"] = request.QueryShrink
	}

	if !dara.IsNil(request.RequestId) {
		body["requestId"] = request.RequestId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateEduProblemSolve"),
		Version:     dara.String("2026-11-02"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/v1/problem/solve/stream"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
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
