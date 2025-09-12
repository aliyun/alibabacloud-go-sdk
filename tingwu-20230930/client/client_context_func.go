// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 创建听悟任务
//
// @param request - CreateTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTaskResponse
func (client *Client) CreateTaskWithContext(ctx context.Context, request *CreateTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Operation) {
		query["operation"] = request.Operation
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AppKey) {
		body["AppKey"] = request.AppKey
	}

	if !dara.IsNil(request.Input) {
		body["Input"] = request.Input
	}

	if !dara.IsNil(request.Parameters) {
		body["Parameters"] = request.Parameters
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTask"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/tingwu/v2/tasks"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建热词词表
//
// @param request - CreateTranscriptionPhrasesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTranscriptionPhrasesResponse
func (client *Client) CreateTranscriptionPhrasesWithContext(ctx context.Context, request *CreateTranscriptionPhrasesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTranscriptionPhrasesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.WordWeights) {
		body["WordWeights"] = request.WordWeights
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTranscriptionPhrases"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/tingwu/v2/resources/phrases"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTranscriptionPhrasesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除词表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTranscriptionPhrasesResponse
func (client *Client) DeleteTranscriptionPhrasesWithContext(ctx context.Context, PhraseId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteTranscriptionPhrasesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteTranscriptionPhrases"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/tingwu/v2/resources/phrases/" + dara.PercentEncode(dara.StringValue(PhraseId))),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteTranscriptionPhrasesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询听悟任务信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskInfoResponse
func (client *Client) GetTaskInfoWithContext(ctx context.Context, TaskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTaskInfoResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTaskInfo"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/tingwu/v2/tasks/" + dara.PercentEncode(dara.StringValue(TaskId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTaskInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询热词词表信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTranscriptionPhrasesResponse
func (client *Client) GetTranscriptionPhrasesWithContext(ctx context.Context, PhraseId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTranscriptionPhrasesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTranscriptionPhrases"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/tingwu/v2/resources/phrases/" + dara.PercentEncode(dara.StringValue(PhraseId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTranscriptionPhrasesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举用户所有热词词表信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTranscriptionPhrasesResponse
func (client *Client) ListTranscriptionPhrasesWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTranscriptionPhrasesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTranscriptionPhrases"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/tingwu/v2/resources/phrases"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTranscriptionPhrasesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新热词词表
//
// @param request - UpdateTranscriptionPhrasesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTranscriptionPhrasesResponse
func (client *Client) UpdateTranscriptionPhrasesWithContext(ctx context.Context, PhraseId *string, request *UpdateTranscriptionPhrasesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateTranscriptionPhrasesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.WordWeights) {
		body["WordWeights"] = request.WordWeights
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateTranscriptionPhrases"),
		Version:     dara.String("2023-09-30"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/openapi/tingwu/v2/resources/phrases/" + dara.PercentEncode(dara.StringValue(PhraseId))),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateTranscriptionPhrasesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}
