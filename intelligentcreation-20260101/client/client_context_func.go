// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 检查Turing任务
//
// @param request - CheckTuringTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckTuringTaskResponse
func (client *Client) CheckTuringTaskWithContext(ctx context.Context, request *CheckTuringTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CheckTuringTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["taskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckTuringTask"),
		Version:     dara.String("2026-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/aigc-turing/openService/v1/task/checkTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckTuringTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 任务提交接口
//
// @param request - SubmitTuringTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitTuringTaskResponse
func (client *Client) SubmitTuringTaskWithContext(ctx context.Context, request *SubmitTuringTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitTuringTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Duration) {
		body["duration"] = request.Duration
	}

	if !dara.IsNil(request.IdempotentKey) {
		body["idempotentKey"] = request.IdempotentKey
	}

	if !dara.IsNil(request.ImgUrl) {
		body["imgUrl"] = request.ImgUrl
	}

	if !dara.IsNil(request.Resolution) {
		body["resolution"] = request.Resolution
	}

	if !dara.IsNil(request.ResourceType) {
		body["resourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.TaskType) {
		body["taskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitTuringTask"),
		Version:     dara.String("2026-01-01"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/aigc-turing/openService/v1/task/submitTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitTuringTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
