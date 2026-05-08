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
// 添加文案反馈
//
// @param request - AddTextFeedbackRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddTextFeedbackResponse
func (client *Client) AddTextFeedbackWithContext(ctx context.Context, request *AddTextFeedbackRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddTextFeedbackResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	if !dara.IsNil(request.Quality) {
		body["quality"] = request.Quality
	}

	if !dara.IsNil(request.TextId) {
		body["textId"] = request.TextId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddTextFeedback"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/addTextFeedback"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AddTextFeedbackResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量添加知识文档
//
// @param request - BatchAddDocumentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchAddDocumentResponse
func (client *Client) BatchAddDocumentWithContext(ctx context.Context, knowledgeBaseId *string, request *BatchAddDocumentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchAddDocumentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AddDocumentInfos) {
		body["addDocumentInfos"] = request.AddDocumentInfos
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchAddDocument"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/knowledge-base/" + dara.PercentEncode(dara.StringValue(knowledgeBaseId)) + "/documents"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchAddDocumentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量发布剧本任务
//
// @param request - BatchCreateAICoachTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchCreateAICoachTaskResponse
func (client *Client) BatchCreateAICoachTaskWithContext(ctx context.Context, request *BatchCreateAICoachTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchCreateAICoachTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RequestId) {
		body["requestId"] = request.RequestId
	}

	if !dara.IsNil(request.ScriptRecordId) {
		body["scriptRecordId"] = request.ScriptRecordId
	}

	if !dara.IsNil(request.StudentIds) {
		body["studentIds"] = request.StudentIds
	}

	if !dara.IsNil(request.StudentList) {
		body["studentList"] = request.StudentList
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchCreateAICoachTask"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/aicoach/batchCreateTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchCreateAICoachTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量删除对练任务
//
// @param tmpReq - BatchDeletePracticeTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchDeletePracticeTaskResponse
func (client *Client) BatchDeletePracticeTaskWithContext(ctx context.Context, tmpReq *BatchDeletePracticeTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchDeletePracticeTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchDeletePracticeTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TaskIds) {
		request.TaskIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskIds, dara.String("taskIds"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.IdempotentId) {
		query["idempotentId"] = request.IdempotentId
	}

	if !dara.IsNil(request.TaskIdsShrink) {
		query["taskIds"] = request.TaskIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchDeletePracticeTask"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/aicoach/batchDeletePracticeTask"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchDeletePracticeTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询项目信息
//
// @param tmpReq - BatchGetProjectTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchGetProjectTaskResponse
func (client *Client) BatchGetProjectTaskWithContext(ctx context.Context, tmpReq *BatchGetProjectTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchGetProjectTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchGetProjectTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TaskIdList) {
		request.TaskIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskIdList, dara.String("taskIdList"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskIdListShrink) {
		query["taskIdList"] = request.TaskIdListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchGetProjectTask"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/digitalHuman/project/batchGetProjectTask"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchGetProjectTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询声音复刻任务
//
// @param tmpReq - BatchGetTrainTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchGetTrainTaskResponse
func (client *Client) BatchGetTrainTaskWithContext(ctx context.Context, tmpReq *BatchGetTrainTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchGetTrainTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchGetTrainTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TaskIdList) {
		request.TaskIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskIdList, dara.String("taskIdList"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.AliyunMainId) {
		query["aliyunMainId"] = request.AliyunMainId
	}

	if !dara.IsNil(request.TaskIdListShrink) {
		query["taskIdList"] = request.TaskIdListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchGetTrainTask"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/train/task/batchGetTrainTaskInfo"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchGetTrainTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询视频切片任务信息
//
// @param tmpReq - BatchGetVideoClipTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchGetVideoClipTaskResponse
func (client *Client) BatchGetVideoClipTaskWithContext(ctx context.Context, tmpReq *BatchGetVideoClipTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchGetVideoClipTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchGetVideoClipTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TaskIdList) {
		request.TaskIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TaskIdList, dara.String("taskIdList"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TaskIdListShrink) {
		query["taskIdList"] = request.TaskIdListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchGetVideoClipTask"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/video/clip/batchGetVideoClipTask"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchGetVideoClipTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询文案
//
// @param tmpReq - BatchQueryIndividuationTextRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchQueryIndividuationTextResponse
func (client *Client) BatchQueryIndividuationTextWithContext(ctx context.Context, tmpReq *BatchQueryIndividuationTextRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchQueryIndividuationTextResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &BatchQueryIndividuationTextShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.TextIdList) {
		request.TextIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TextIdList, dara.String("textIdList"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.TextIdListShrink) {
		query["textIdList"] = request.TextIdListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchQueryIndividuationText"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/individuationText/batchQueryText"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchQueryIndividuationTextResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 快速发布剧本
//
// @param request - BuildAICoachScriptRecordRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BuildAICoachScriptRecordResponse
func (client *Client) BuildAICoachScriptRecordWithContext(ctx context.Context, request *BuildAICoachScriptRecordRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BuildAICoachScriptRecordResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ScriptJsonUrl) {
		body["scriptJsonUrl"] = request.ScriptJsonUrl
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BuildAICoachScriptRecord"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/aicoach/buildScriptRecord"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BuildAICoachScriptRecordResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 检查会话状态
//
// @param request - CheckSessionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CheckSessionResponse
func (client *Client) CheckSessionWithContext(ctx context.Context, request *CheckSessionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CheckSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		body["projectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CheckSession"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/avatar/project/checkSession"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CheckSessionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 学员关闭会话
//
// @param request - CloseAICoachTaskSessionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseAICoachTaskSessionResponse
func (client *Client) CloseAICoachTaskSessionWithContext(ctx context.Context, request *CloseAICoachTaskSessionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CloseAICoachTaskSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Uid) {
		body["uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloseAICoachTaskSession"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/aicoach/closeSession"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CloseAICoachTaskSessionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文本数量统计
//
// @param request - CountTextRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CountTextResponse
func (client *Client) CountTextWithContext(ctx context.Context, request *CountTextRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CountTextResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GenerationSource) {
		query["generationSource"] = request.GenerationSource
	}

	if !dara.IsNil(request.Industry) {
		query["industry"] = request.Industry
	}

	if !dara.IsNil(request.PublishStatus) {
		query["publishStatus"] = request.PublishStatus
	}

	if !dara.IsNil(request.Style) {
		query["style"] = request.Style
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CountText"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/countText"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CountTextResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询剧本列表
//
// @param request - CreateAICoachTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAICoachTaskResponse
func (client *Client) CreateAICoachTaskWithContext(ctx context.Context, request *CreateAICoachTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAICoachTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.RequestId) {
		body["requestId"] = request.RequestId
	}

	if !dara.IsNil(request.ScriptRecordId) {
		body["scriptRecordId"] = request.ScriptRecordId
	}

	if !dara.IsNil(request.StudentAudioUrl) {
		body["studentAudioUrl"] = request.StudentAudioUrl
	}

	if !dara.IsNil(request.StudentId) {
		body["studentId"] = request.StudentId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAICoachTask"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/aicoach/createTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAICoachTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建离线评测任务
//
// @param request - CreateAICoachTaskReportRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAICoachTaskReportResponse
func (client *Client) CreateAICoachTaskReportWithContext(ctx context.Context, request *CreateAICoachTaskReportRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAICoachTaskReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DialogueList) {
		body["dialogueList"] = request.DialogueList
	}

	if !dara.IsNil(request.IdempotentId) {
		body["idempotentId"] = request.IdempotentId
	}

	if !dara.IsNil(request.TaskId) {
		body["taskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAICoachTaskReport"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/aicoach/startSessionReport"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAICoachTaskReportResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 学员开启对练会话
//
// @param request - CreateAICoachTaskSessionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAICoachTaskSessionResponse
func (client *Client) CreateAICoachTaskSessionWithContext(ctx context.Context, request *CreateAICoachTaskSessionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAICoachTaskSessionResponse, _err error) {
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

	if !dara.IsNil(request.Uid) {
		body["uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAICoachTaskSession"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/aicoach/startSession"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAICoachTaskSessionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # CreateAgent
//
// @param request - CreateAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAgentResponse
func (client *Client) CreateAgentWithContext(ctx context.Context, request *CreateAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentIconUrl) {
		body["agentIconUrl"] = request.AgentIconUrl
	}

	if !dara.IsNil(request.AgentName) {
		body["agentName"] = request.AgentName
	}

	if !dara.IsNil(request.AgentScene) {
		body["agentScene"] = request.AgentScene
	}

	if !dara.IsNil(request.CharacterAgeStage) {
		body["characterAgeStage"] = request.CharacterAgeStage
	}

	if !dara.IsNil(request.CharacterGender) {
		body["characterGender"] = request.CharacterGender
	}

	if !dara.IsNil(request.CharacterName) {
		body["characterName"] = request.CharacterName
	}

	if !dara.IsNil(request.ExtraDescription) {
		body["extraDescription"] = request.ExtraDescription
	}

	if !dara.IsNil(request.Industry) {
		body["industry"] = request.Industry
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAgent"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/agent/createAgent"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAgentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建照片数字人
//
// @param request - CreateAnchorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAnchorResponse
func (client *Client) CreateAnchorWithContext(ctx context.Context, request *CreateAnchorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAnchorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AnchorCategory) {
		body["anchorCategory"] = request.AnchorCategory
	}

	if !dara.IsNil(request.AnchorMaterialName) {
		body["anchorMaterialName"] = request.AnchorMaterialName
	}

	if !dara.IsNil(request.CoverUrl) {
		body["coverUrl"] = request.CoverUrl
	}

	if !dara.IsNil(request.DigitalHumanType) {
		body["digitalHumanType"] = request.DigitalHumanType
	}

	if !dara.IsNil(request.Gender) {
		body["gender"] = request.Gender
	}

	if !dara.IsNil(request.UseScene) {
		body["useScene"] = request.UseScene
	}

	if !dara.IsNil(request.VideoOssKey) {
		body["videoOssKey"] = request.VideoOssKey
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAnchor"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/digitalHuman/anchorOpen/createAnchor"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAnchorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建配图生成任务
//
// @param request - CreateIllustrationTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIllustrationTaskResponse
func (client *Client) CreateIllustrationTaskWithContext(ctx context.Context, textId *string, request *CreateIllustrationTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateIllustrationTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateIllustrationTask"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/texts/" + dara.PercentEncode(dara.StringValue(textId)) + "/illustrationTasks"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIllustrationTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建个性化文案项目
//
// @param request - CreateIndividuationProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIndividuationProjectResponse
func (client *Client) CreateIndividuationProjectWithContext(ctx context.Context, request *CreateIndividuationProjectRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateIndividuationProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ProjectInfo) {
		body["projectInfo"] = request.ProjectInfo
	}

	if !dara.IsNil(request.ProjectName) {
		body["projectName"] = request.ProjectName
	}

	if !dara.IsNil(request.Purpose) {
		body["purpose"] = request.Purpose
	}

	if !dara.IsNil(request.SceneId) {
		body["sceneId"] = request.SceneId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateIndividuationProject"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/individuationText/createProject"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIndividuationProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建个性化文案任务
//
// @param request - CreateIndividuationTextTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIndividuationTextTaskResponse
func (client *Client) CreateIndividuationTextTaskWithContext(ctx context.Context, request *CreateIndividuationTextTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateIndividuationTextTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CrowdPack) {
		body["crowdPack"] = request.CrowdPack
	}

	if !dara.IsNil(request.ProjectId) {
		body["projectId"] = request.ProjectId
	}

	if !dara.IsNil(request.TaskName) {
		body["taskName"] = request.TaskName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateIndividuationTextTask"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/individuationText/createTextTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateIndividuationTextTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建产品图
//
// @param request - CreateProductImageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateProductImageResponse
func (client *Client) CreateProductImageWithContext(ctx context.Context, request *CreateProductImageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateProductImageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BackgroundDescription) {
		body["backgroundDescription"] = request.BackgroundDescription
	}

	if !dara.IsNil(request.BackgroundPriority) {
		body["backgroundPriority"] = request.BackgroundPriority
	}

	if !dara.IsNil(request.BackgroundUrl) {
		body["backgroundUrl"] = request.BackgroundUrl
	}

	if !dara.IsNil(request.HighlightText) {
		body["highlightText"] = request.HighlightText
	}

	if !dara.IsNil(request.ImageCount) {
		body["imageCount"] = request.ImageCount
	}

	if !dara.IsNil(request.ImageUrl) {
		body["imageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.SubTitle) {
		body["subTitle"] = request.SubTitle
	}

	if !dara.IsNil(request.Title) {
		body["title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateProductImage"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/images/products"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateProductImageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 写实人像创作
//
// @param request - CreateRealisticPortraitRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateRealisticPortraitResponse
func (client *Client) CreateRealisticPortraitWithContext(ctx context.Context, request *CreateRealisticPortraitRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateRealisticPortraitResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Ages) {
		body["ages"] = request.Ages
	}

	if !dara.IsNil(request.Cloth) {
		body["cloth"] = request.Cloth
	}

	if !dara.IsNil(request.Color) {
		body["color"] = request.Color
	}

	if !dara.IsNil(request.Custom) {
		body["custom"] = request.Custom
	}

	if !dara.IsNil(request.Face) {
		body["face"] = request.Face
	}

	if !dara.IsNil(request.Figure) {
		body["figure"] = request.Figure
	}

	if !dara.IsNil(request.Gender) {
		body["gender"] = request.Gender
	}

	if !dara.IsNil(request.HairColor) {
		body["hairColor"] = request.HairColor
	}

	if !dara.IsNil(request.Hairstyle) {
		body["hairstyle"] = request.Hairstyle
	}

	if !dara.IsNil(request.Height) {
		body["height"] = request.Height
	}

	if !dara.IsNil(request.ImageUrl) {
		body["imageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.Numbers) {
		body["numbers"] = request.Numbers
	}

	if !dara.IsNil(request.Ratio) {
		body["ratio"] = request.Ratio
	}

	if !dara.IsNil(request.Width) {
		body["width"] = request.Width
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateRealisticPortrait"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/images/portrait/realistic"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateRealisticPortraitResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建文案生成任务
//
// @param request - CreateTextTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTextTaskResponse
func (client *Client) CreateTextTaskWithContext(ctx context.Context, request *CreateTextTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTextTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTextTask"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/textTasks"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTextTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交声音复刻任务
//
// @param request - CreateTrainTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTrainTaskResponse
func (client *Client) CreateTrainTaskWithContext(ctx context.Context, request *CreateTrainTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTrainTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AliyunMainId) {
		body["aliyunMainId"] = request.AliyunMainId
	}

	if !dara.IsNil(request.ResSpecType) {
		body["resSpecType"] = request.ResSpecType
	}

	if !dara.IsNil(request.TaskType) {
		body["taskType"] = request.TaskType
	}

	if !dara.IsNil(request.UseScene) {
		body["useScene"] = request.UseScene
	}

	if !dara.IsNil(request.VoiceGender) {
		body["voiceGender"] = request.VoiceGender
	}

	if !dara.IsNil(request.VoiceLanguage) {
		body["voiceLanguage"] = request.VoiceLanguage
	}

	if !dara.IsNil(request.VoiceName) {
		body["voiceName"] = request.VoiceName
	}

	if !dara.IsNil(request.VoicePath) {
		body["voicePath"] = request.VoicePath
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTrainTask"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/train/task/createTrainTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTrainTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交视频切片任务
//
// @param request - CreateVideoClipTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVideoClipTaskResponse
func (client *Client) CreateVideoClipTaskWithContext(ctx context.Context, request *CreateVideoClipTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateVideoClipTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AliyunMainId) {
		body["aliyunMainId"] = request.AliyunMainId
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.OssKeys) {
		body["ossKeys"] = request.OssKeys
	}

	if !dara.IsNil(request.Requirement) {
		body["requirement"] = request.Requirement
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVideoClipTask"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/video/clip/createVideoClipTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVideoClipTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DeleteAICoachScript
//
// @param request - DeleteAICoachScriptRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAICoachScriptResponse
func (client *Client) DeleteAICoachScriptWithContext(ctx context.Context, request *DeleteAICoachScriptRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAICoachScriptResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ScriptId) {
		body["scriptId"] = request.ScriptId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAICoachScript"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/aicoach/deleteAICoachScript"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAICoachScriptResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # DeleteAgent
//
// @param request - DeleteAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAgentResponse
func (client *Client) DeleteAgentWithContext(ctx context.Context, request *DeleteAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		body["agentId"] = request.AgentId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteAgent"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/agent/deleteAgent"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteAgentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除个性化文案项目
//
// @param request - DeleteIndividuationProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIndividuationProjectResponse
func (client *Client) DeleteIndividuationProjectWithContext(ctx context.Context, request *DeleteIndividuationProjectRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteIndividuationProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		body["projectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteIndividuationProject"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/individuationText/deleteProject"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteIndividuationProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除个性化文案
//
// @param request - DeleteIndividuationTextRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteIndividuationTextResponse
func (client *Client) DeleteIndividuationTextWithContext(ctx context.Context, request *DeleteIndividuationTextRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteIndividuationTextResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TextIdList) {
		body["textIdList"] = request.TextIdList
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteIndividuationText"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/individuationText/deleteText"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteIndividuationTextResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询文档信息与状态
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDocumentResponse
func (client *Client) DescribeDocumentWithContext(ctx context.Context, knowledgeBaseId *string, documentId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeDocumentResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDocument"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/knowledge-base/" + dara.PercentEncode(dara.StringValue(knowledgeBaseId)) + "/documents/" + dara.PercentEncode(dara.StringValue(documentId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DescribeDocumentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 学员完成会话
//
// @param request - FinishAICoachTaskSessionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FinishAICoachTaskSessionResponse
func (client *Client) FinishAICoachTaskSessionWithContext(ctx context.Context, request *FinishAICoachTaskSessionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *FinishAICoachTaskSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Uid) {
		body["uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("FinishAICoachTaskSession"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/aicoach/finishSession"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &FinishAICoachTaskSessionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取考核点详情
//
// @param request - GetAICoachAssessmentPointRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAICoachAssessmentPointResponse
func (client *Client) GetAICoachAssessmentPointWithContext(ctx context.Context, request *GetAICoachAssessmentPointRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAICoachAssessmentPointResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PointId) {
		query["pointId"] = request.PointId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAICoachAssessmentPoint"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/aicoach/getAssessmentPoint"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAICoachAssessmentPointResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询作弊检测详情
//
// @param request - GetAICoachCheatDetectionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAICoachCheatDetectionResponse
func (client *Client) GetAICoachCheatDetectionWithContext(ctx context.Context, request *GetAICoachCheatDetectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAICoachCheatDetectionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SessionId) {
		query["sessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAICoachCheatDetection"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/aicoach/getCheatDetection"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAICoachCheatDetectionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询剧本详情
//
// @param request - GetAICoachScriptRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAICoachScriptResponse
func (client *Client) GetAICoachScriptWithContext(ctx context.Context, request *GetAICoachScriptRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAICoachScriptResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ScriptRecordId) {
		query["scriptRecordId"] = request.ScriptRecordId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAICoachScript"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/aicoach/getScript"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAICoachScriptResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 学员查询会话历史
//
// @param request - GetAICoachTaskSessionHistoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAICoachTaskSessionHistoryResponse
func (client *Client) GetAICoachTaskSessionHistoryWithContext(ctx context.Context, request *GetAICoachTaskSessionHistoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAICoachTaskSessionHistoryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SessionId) {
		query["sessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Uid) {
		query["uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAICoachTaskSessionHistory"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/aicoach/querySessionHistory"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAICoachTaskSessionHistoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 学员查询会话评测报告
//
// @param request - GetAICoachTaskSessionReportRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAICoachTaskSessionReportResponse
func (client *Client) GetAICoachTaskSessionReportWithContext(ctx context.Context, request *GetAICoachTaskSessionReportRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAICoachTaskSessionReportResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SessionId) {
		query["sessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Uid) {
		query["uid"] = request.Uid
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAICoachTaskSessionReport"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/aicoach/queryTaskSessionReport"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAICoachTaskSessionReportResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取对练会话资源使用情况
//
// @param request - GetAICoachTaskSessionResourceUsageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAICoachTaskSessionResourceUsageResponse
func (client *Client) GetAICoachTaskSessionResourceUsageWithContext(ctx context.Context, request *GetAICoachTaskSessionResourceUsageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAICoachTaskSessionResourceUsageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.SessionId) {
		query["sessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAICoachTaskSessionResourceUsage"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/aicoach/getSessionResourceUsage"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAICoachTaskSessionResourceUsageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询配图
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIllustrationResponse
func (client *Client) GetIllustrationWithContext(ctx context.Context, textId *string, illustrationId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetIllustrationResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIllustration"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/texts/" + dara.PercentEncode(dara.StringValue(textId)) + "/illustrations/" + dara.PercentEncode(dara.StringValue(illustrationId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIllustrationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询配图任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIllustrationTaskResponse
func (client *Client) GetIllustrationTaskWithContext(ctx context.Context, textId *string, illustrationTaskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetIllustrationTaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetIllustrationTask"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/texts/" + dara.PercentEncode(dara.StringValue(textId)) + "/illustrationTasks/" + dara.PercentEncode(dara.StringValue(illustrationTaskId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetIllustrationTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取图片上传oss token
//
// @param request - GetOssUploadTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOssUploadTokenResponse
func (client *Client) GetOssUploadTokenWithContext(ctx context.Context, request *GetOssUploadTokenRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetOssUploadTokenResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileName) {
		query["fileName"] = request.FileName
	}

	if !dara.IsNil(request.FileType) {
		query["fileType"] = request.FileType
	}

	if !dara.IsNil(request.UploadType) {
		query["uploadType"] = request.UploadType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOssUploadToken"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/uploadToken"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOssUploadTokenResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数据人合成信息
//
// @param request - GetProjectTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetProjectTaskResponse
func (client *Client) GetProjectTaskWithContext(ctx context.Context, request *GetProjectTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetProjectTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IdempotentId) {
		query["IdempotentId"] = request.IdempotentId
	}

	if !dara.IsNil(request.TaskId) {
		query["taskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetProjectTask"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/digitalHuman/project/getProjectTask"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetProjectTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询文案
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTextResponse
func (client *Client) GetTextWithContext(ctx context.Context, textId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTextResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetText"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/texts/" + dara.PercentEncode(dara.StringValue(textId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTextResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询文案任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTextTaskResponse
func (client *Client) GetTextTaskWithContext(ctx context.Context, textTaskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTextTaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTextTask"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/textTasks/" + dara.PercentEncode(dara.StringValue(textTaskId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTextTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询表单配置
//
// @param request - GetTextTemplateRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTextTemplateResponse
func (client *Client) GetTextTemplateWithContext(ctx context.Context, request *GetTextTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTextTemplateResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Industry) {
		query["industry"] = request.Industry
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTextTemplate"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/texts/commands/getTextTemplate"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTextTemplateResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 营销文案互动问答
//
// @param request - InteractTextRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InteractTextResponse
func (client *Client) InteractTextWithSSECtx(ctx context.Context, request *InteractTextRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *InteractTextResponse, _yieldErr chan error) {
	defer close(_yield)
	client.interactTextWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, headers, runtime)
	return
}

// Summary:
//
// 营销文案互动问答
//
// @param request - InteractTextRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InteractTextResponse
func (client *Client) InteractTextWithContext(ctx context.Context, request *InteractTextRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InteractTextResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		body["agentId"] = request.AgentId
	}

	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InteractText"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/stream/interactText"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InteractTextResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询剧本列表
//
// @param request - ListAICoachScriptPageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAICoachScriptPageResponse
func (client *Client) ListAICoachScriptPageWithContext(ctx context.Context, request *ListAICoachScriptPageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAICoachScriptPageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		query["name"] = request.Name
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	if !dara.IsNil(request.Type) {
		query["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAICoachScriptPage"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/aicoach/pageScript"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAICoachScriptPageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询任务列表
//
// @param request - ListAICoachTaskPageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAICoachTaskPageResponse
func (client *Client) ListAICoachTaskPageWithContext(ctx context.Context, request *ListAICoachTaskPageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAICoachTaskPageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.EndTime) {
		query["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.StartTime) {
		query["startTime"] = request.StartTime
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	if !dara.IsNil(request.StudentId) {
		query["studentId"] = request.StudentId
	}

	if !dara.IsNil(request.TaskId) {
		query["taskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAICoachTaskPage"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/aicoach/listTaskPage"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAICoachTaskPageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询智能体
//
// @param request - ListAgentsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAgentsResponse
func (client *Client) ListAgentsWithContext(ctx context.Context, request *ListAgentsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAgentsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		query["agentId"] = request.AgentId
	}

	if !dara.IsNil(request.AgentScene) {
		query["agentScene"] = request.AgentScene
	}

	if !dara.IsNil(request.Owner) {
		query["owner"] = request.Owner
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		query["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAgents"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/agent/listAgents"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAgentsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取数字人模特列表
//
// @param request - ListAnchorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAnchorResponse
func (client *Client) ListAnchorWithContext(ctx context.Context, request *ListAnchorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAnchorResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AnchorCategory) {
		query["anchorCategory"] = request.AnchorCategory
	}

	if !dara.IsNil(request.AnchorId) {
		query["anchorId"] = request.AnchorId
	}

	if !dara.IsNil(request.AnchorType) {
		query["anchorType"] = request.AnchorType
	}

	if !dara.IsNil(request.CoverRate) {
		query["coverRate"] = request.CoverRate
	}

	if !dara.IsNil(request.DigitalHumanType) {
		query["digitalHumanType"] = request.DigitalHumanType
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResSpecType) {
		query["resSpecType"] = request.ResSpecType
	}

	if !dara.IsNil(request.UseScene) {
		query["useScene"] = request.UseScene
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAnchor"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/digitalHuman/anchorOpen/listAnchor"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAnchorResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量查询数字人项目启动结果
//
// @param tmpReq - ListAvatarProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAvatarProjectResponse
func (client *Client) ListAvatarProjectWithContext(ctx context.Context, tmpReq *ListAvatarProjectRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAvatarProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &ListAvatarProjectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.ProjectIdList) {
		request.ProjectIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ProjectIdList, dara.String("projectIdList"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectIdListShrink) {
		query["projectIdList"] = request.ProjectIdListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAvatarProject"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/avatar/project/listAvatarProject"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAvatarProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询知识库
//
// @param request - ListKnowledgeBaseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListKnowledgeBaseResponse
func (client *Client) ListKnowledgeBaseWithContext(ctx context.Context, request *ListKnowledgeBaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListKnowledgeBaseResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.KnowledgeBaseId) {
		query["knowledgeBaseId"] = request.KnowledgeBaseId
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListKnowledgeBase"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/knowledge-base"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListKnowledgeBaseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询文案主题列表
//
// @param request - ListTextThemesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTextThemesResponse
func (client *Client) ListTextThemesWithContext(ctx context.Context, request *ListTextThemesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTextThemesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Industry) {
		query["industry"] = request.Industry
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTextThemes"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/textThemes"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTextThemesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举文案
//
// @param request - ListTextsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTextsResponse
func (client *Client) ListTextsWithContext(ctx context.Context, request *ListTextsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTextsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.GenerationSource) {
		query["generationSource"] = request.GenerationSource
	}

	if !dara.IsNil(request.Industry) {
		query["industry"] = request.Industry
	}

	if !dara.IsNil(request.Keyword) {
		query["keyword"] = request.Keyword
	}

	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.PublishStatus) {
		query["publishStatus"] = request.PublishStatus
	}

	if !dara.IsNil(request.TextStyleType) {
		query["textStyleType"] = request.TextStyleType
	}

	if !dara.IsNil(request.TextTheme) {
		query["textTheme"] = request.TextTheme
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTexts"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/texts"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListTextsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取声音模版列表
//
// @param request - ListVoiceModelsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVoiceModelsResponse
func (client *Client) ListVoiceModelsWithContext(ctx context.Context, request *ListVoiceModelsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListVoiceModelsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNumber) {
		query["pageNumber"] = request.PageNumber
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ResSpecType) {
		query["resSpecType"] = request.ResSpecType
	}

	if !dara.IsNil(request.UseScene) {
		query["useScene"] = request.UseScene
	}

	if !dara.IsNil(request.VoiceLanguage) {
		query["voiceLanguage"] = request.VoiceLanguage
	}

	if !dara.IsNil(request.VoiceType) {
		query["voiceType"] = request.VoiceType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVoiceModels"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/digitalHuman/voiceOpen/listVoiceModels"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVoiceModelsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 下线剧本
//
// @param request - OfflineAICoachScriptRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OfflineAICoachScriptResponse
func (client *Client) OfflineAICoachScriptWithContext(ctx context.Context, request *OfflineAICoachScriptRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OfflineAICoachScriptResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ScriptId) {
		body["scriptId"] = request.ScriptId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OfflineAICoachScript"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/aicoach/offlineAICoachScript"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &OfflineAICoachScriptResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 操作实时数字人项目
//
// @param request - OperateAvatarProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return OperateAvatarProjectResponse
func (client *Client) OperateAvatarProjectWithContext(ctx context.Context, request *OperateAvatarProjectRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OperateAvatarProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.OperateType) {
		body["operateType"] = request.OperateType
	}

	if !dara.IsNil(request.ProjectId) {
		body["projectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ResChannelNumber) {
		body["resChannelNumber"] = request.ResChannelNumber
	}

	if !dara.IsNil(request.ResType) {
		body["resType"] = request.ResType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("OperateAvatarProject"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/avatar/project/operateProjectAvatar"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &OperateAvatarProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询数字人项目信息
//
// @param request - QueryAvatarProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAvatarProjectResponse
func (client *Client) QueryAvatarProjectWithContext(ctx context.Context, request *QueryAvatarProjectRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryAvatarProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		query["projectId"] = request.ProjectId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAvatarProject"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/avatar/project/queryAvatarProject"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAvatarProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查找资源
//
// @param request - QueryAvatarResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryAvatarResourceResponse
func (client *Client) QueryAvatarResourceWithContext(ctx context.Context, request *QueryAvatarResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryAvatarResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IdempotentId) {
		query["idempotentId"] = request.IdempotentId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryAvatarResource"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/avatar/project/queryResource"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryAvatarResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询图片转视频任务
//
// @param request - QueryImageToVideoTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryImageToVideoTaskResponse
func (client *Client) QueryImageToVideoTaskWithContext(ctx context.Context, request *QueryImageToVideoTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryImageToVideoTaskResponse, _err error) {
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
		Action:      dara.String("QueryImageToVideoTask"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/video/imageToVideo/task"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryImageToVideoTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询个性化文案任务
//
// @param request - QueryIndividuationTextTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryIndividuationTextTaskResponse
func (client *Client) QueryIndividuationTextTaskWithContext(ctx context.Context, request *QueryIndividuationTextTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryIndividuationTextTaskResponse, _err error) {
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
		Action:      dara.String("QueryIndividuationTextTask"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/individuationText/queryTextTask"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryIndividuationTextTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询会话信息
//
// @param tmpReq - QuerySessionInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QuerySessionInfoResponse
func (client *Client) QuerySessionInfoWithContext(ctx context.Context, tmpReq *QuerySessionInfoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QuerySessionInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = tmpReq.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	request := &QuerySessionInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.StatusList) {
		request.StatusListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StatusList, dara.String("statusList"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.PageNo) {
		query["pageNo"] = request.PageNo
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProjectId) {
		query["projectId"] = request.ProjectId
	}

	if !dara.IsNil(request.StatusListShrink) {
		query["statusList"] = request.StatusListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QuerySessionInfo"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/avatar/project/querySessionInfo"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QuerySessionInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 流式输出文案
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTextStreamResponse
func (client *Client) QueryTextStreamWithSSECtx(ctx context.Context, textId *string, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *QueryTextStreamResponse, _yieldErr chan error) {
	defer close(_yield)
	client.queryTextStreamWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, textId, headers, runtime)
	return
}

// Summary:
//
// 流式输出文案
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryTextStreamResponse
func (client *Client) QueryTextStreamWithContext(ctx context.Context, textId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryTextStreamResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryTextStream"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/stream/queryTextStream/" + dara.PercentEncode(dara.StringValue(textId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryTextStreamResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # ReleaseAgent
//
// @param request - ReleaseAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseAgentResponse
func (client *Client) ReleaseAgentWithContext(ctx context.Context, request *ReleaseAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ReleaseAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		body["agentId"] = request.AgentId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReleaseAgent"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/agent/releaseAgent"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ReleaseAgentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 保存实时数字人项目
//
// @param request - SaveAvatarProjectRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveAvatarProjectResponse
func (client *Client) SaveAvatarProjectWithContext(ctx context.Context, request *SaveAvatarProjectRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SaveAvatarProjectResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		body["agentId"] = request.AgentId
	}

	if !dara.IsNil(request.BitRate) {
		body["bitRate"] = request.BitRate
	}

	if !dara.IsNil(request.FrameRate) {
		body["frameRate"] = request.FrameRate
	}

	if !dara.IsNil(request.Frames) {
		body["frames"] = request.Frames
	}

	if !dara.IsNil(request.OperateType) {
		body["operateType"] = request.OperateType
	}

	if !dara.IsNil(request.ProjectId) {
		body["projectId"] = request.ProjectId
	}

	if !dara.IsNil(request.ProjectName) {
		body["projectName"] = request.ProjectName
	}

	if !dara.IsNil(request.ResSpecType) {
		body["resSpecType"] = request.ResSpecType
	}

	if !dara.IsNil(request.Resolution) {
		body["resolution"] = request.Resolution
	}

	if !dara.IsNil(request.ScaleType) {
		body["scaleType"] = request.ScaleType
	}

	if !dara.IsNil(request.ScriptModelTag) {
		body["scriptModelTag"] = request.ScriptModelTag
	}

	if !dara.IsNil(request.SynchronizedDisplay) {
		body["synchronizedDisplay"] = request.SynchronizedDisplay
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SaveAvatarProject"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/avatar/project/saveAvatarProject"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SaveAvatarProjectResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询图片任务
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SelectImageTaskResponse
func (client *Client) SelectImageTaskWithContext(ctx context.Context, taskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SelectImageTaskResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("SelectImageTask"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/images/portrait/select/" + dara.PercentEncode(dara.StringValue(taskId))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SelectImageTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询离线数字人剩余资源
//
// @param request - SelectResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SelectResourceResponse
func (client *Client) SelectResourceWithContext(ctx context.Context, request *SelectResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SelectResourceResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.IdempotentId) {
		query["idempotentId"] = request.IdempotentId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SelectResource"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/digitalHuman/project/commands/overview"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SelectResourceResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发送sdk消息
//
// @param request - SendSdkMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendSdkMessageResponse
func (client *Client) SendSdkMessageWithContext(ctx context.Context, request *SendSdkMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SendSdkMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Data) {
		body["data"] = request.Data
	}

	if !dara.IsNil(request.Header) {
		body["header"] = request.Header
	}

	if !dara.IsNil(request.ModuleName) {
		body["moduleName"] = request.ModuleName
	}

	if !dara.IsNil(request.OperationName) {
		body["operationName"] = request.OperationName
	}

	if !dara.IsNil(request.UserId) {
		body["userId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendSdkMessage"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/sdk/sendMessage"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SendSdkMessageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发送sdk流式消息
//
// @param request - SendSdkStreamMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendSdkStreamMessageResponse
func (client *Client) SendSdkStreamMessageWithSSECtx(ctx context.Context, request *SendSdkStreamMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *SendSdkStreamMessageResponse, _yieldErr chan error) {
	defer close(_yield)
	client.sendSdkStreamMessageWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, headers, runtime)
	return
}

// Summary:
//
// 发送sdk流式消息
//
// @param request - SendSdkStreamMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendSdkStreamMessageResponse
func (client *Client) SendSdkStreamMessageWithContext(ctx context.Context, request *SendSdkStreamMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SendSdkStreamMessageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Data) {
		body["data"] = request.Data
	}

	if !dara.IsNil(request.Header) {
		body["header"] = request.Header
	}

	if !dara.IsNil(request.ModuleName) {
		body["moduleName"] = request.ModuleName
	}

	if !dara.IsNil(request.OperationName) {
		body["operationName"] = request.OperationName
	}

	if !dara.IsNil(request.UserId) {
		body["userId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendSdkStreamMessage"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/sdk/stream/sendMessage"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SendSdkStreamMessageResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 发送文本消息
//
// @param request - SendTextMsgRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendTextMsgResponse
func (client *Client) SendTextMsgWithContext(ctx context.Context, request *SendTextMsgRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SendTextMsgResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		body["projectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RequestId) {
		body["requestId"] = request.RequestId
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Text) {
		body["text"] = request.Text
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendTextMsg"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/avatar/project/sendTextMsg"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SendTextMsgResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 启动会话
//
// @param request - StartAvatarSessionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartAvatarSessionResponse
func (client *Client) StartAvatarSessionWithContext(ctx context.Context, request *StartAvatarSessionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StartAvatarSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ChannelToken) {
		body["channelToken"] = request.ChannelToken
	}

	if !dara.IsNil(request.CustomPushUrl) {
		body["customPushUrl"] = request.CustomPushUrl
	}

	if !dara.IsNil(request.CustomUserId) {
		body["customUserId"] = request.CustomUserId
	}

	if !dara.IsNil(request.ProjectId) {
		body["projectId"] = request.ProjectId
	}

	if !dara.IsNil(request.RequestId) {
		body["requestId"] = request.RequestId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StartAvatarSession"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/avatar/project/startAvatarSession"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StartAvatarSessionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 停止会话
//
// @param request - StopAvatarSessionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopAvatarSessionResponse
func (client *Client) StopAvatarSessionWithContext(ctx context.Context, request *StopAvatarSessionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopAvatarSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ProjectId) {
		body["projectId"] = request.ProjectId
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopAvatarSession"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/avatar/project/stopAvatarSession"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopAvatarSessionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视频合成任务停止
//
// @param request - StopProjectTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopProjectTaskResponse
func (client *Client) StopProjectTaskWithContext(ctx context.Context, request *StopProjectTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopProjectTaskResponse, _err error) {
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
		Action:      dara.String("StopProjectTask"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/digitalHuman/project/stop"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &StopProjectTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交图片转视频任务
//
// @param request - SubmitImageToVideoTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitImageToVideoTaskResponse
func (client *Client) SubmitImageToVideoTaskWithContext(ctx context.Context, request *SubmitImageToVideoTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitImageToVideoTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageUrl) {
		body["imageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.PosPrompt) {
		body["posPrompt"] = request.PosPrompt
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitImageToVideoTask"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/video/imageToVideo/task"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitImageToVideoTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交离线数字人合成任务
//
// @param request - SubmitProjectTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitProjectTaskResponse
func (client *Client) SubmitProjectTaskWithContext(ctx context.Context, request *SubmitProjectTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitProjectTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Frames) {
		body["frames"] = request.Frames
	}

	if !dara.IsNil(request.ScaleType) {
		body["scaleType"] = request.ScaleType
	}

	if !dara.IsNil(request.SubtitleTag) {
		body["subtitleTag"] = request.SubtitleTag
	}

	if !dara.IsNil(request.TransparentBackground) {
		body["transparentBackground"] = request.TransparentBackground
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitProjectTask"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/digitalHuman/project/submitProjectTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitProjectTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 人像风格变化
//
// @param request - TransferPortraitStyleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TransferPortraitStyleResponse
func (client *Client) TransferPortraitStyleWithContext(ctx context.Context, request *TransferPortraitStyleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TransferPortraitStyleResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Height) {
		body["height"] = request.Height
	}

	if !dara.IsNil(request.ImageUrl) {
		body["imageUrl"] = request.ImageUrl
	}

	if !dara.IsNil(request.Numbers) {
		body["numbers"] = request.Numbers
	}

	if !dara.IsNil(request.RedrawAmplitude) {
		body["redrawAmplitude"] = request.RedrawAmplitude
	}

	if !dara.IsNil(request.Style) {
		body["style"] = request.Style
	}

	if !dara.IsNil(request.Width) {
		body["width"] = request.Width
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TransferPortraitStyle"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/images/portrait/transferPortraitStyle"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &TransferPortraitStyleResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # UpdateAgent
//
// @param request - UpdateAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAgentResponse
func (client *Client) UpdateAgentWithContext(ctx context.Context, request *UpdateAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentIconUrl) {
		body["agentIconUrl"] = request.AgentIconUrl
	}

	if !dara.IsNil(request.AgentId) {
		body["agentId"] = request.AgentId
	}

	if !dara.IsNil(request.AgentName) {
		body["agentName"] = request.AgentName
	}

	if !dara.IsNil(request.CharacterAgeStage) {
		body["characterAgeStage"] = request.CharacterAgeStage
	}

	if !dara.IsNil(request.CharacterGender) {
		body["characterGender"] = request.CharacterGender
	}

	if !dara.IsNil(request.CharacterName) {
		body["characterName"] = request.CharacterName
	}

	if !dara.IsNil(request.ExtraDescription) {
		body["extraDescription"] = request.ExtraDescription
	}

	if !dara.IsNil(request.Industry) {
		body["industry"] = request.Industry
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateAgent"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/agent/updateAgent"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateAgentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) interactTextWithSSECtx_opYieldFunc(_yield chan *InteractTextResponse, _yieldErr chan error, ctx context.Context, request *InteractTextRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AgentId) {
		body["agentId"] = request.AgentId
	}

	if !dara.IsNil(request.Content) {
		body["content"] = request.Content
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InteractText"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/stream/interactText"),
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

func (client *Client) queryTextStreamWithSSECtx_opYieldFunc(_yield chan *QueryTextStreamResponse, _yieldErr chan error, ctx context.Context, textId *string, headers map[string]*string, runtime *dara.RuntimeOptions) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryTextStream"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/stream/queryTextStream/" + dara.PercentEncode(dara.StringValue(textId))),
		Method:      dara.String("GET"),
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

func (client *Client) sendSdkStreamMessageWithSSECtx_opYieldFunc(_yield chan *SendSdkStreamMessageResponse, _yieldErr chan error, ctx context.Context, request *SendSdkStreamMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Data) {
		body["data"] = request.Data
	}

	if !dara.IsNil(request.Header) {
		body["header"] = request.Header
	}

	if !dara.IsNil(request.ModuleName) {
		body["moduleName"] = request.ModuleName
	}

	if !dara.IsNil(request.OperationName) {
		body["operationName"] = request.OperationName
	}

	if !dara.IsNil(request.UserId) {
		body["userId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendSdkStreamMessage"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/sdk/stream/sendMessage"),
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
