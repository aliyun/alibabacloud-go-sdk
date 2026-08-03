// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

type Client struct {
	openapi.Client
	DisableSDKError *bool
	EnableValidate  *bool
}

func NewClient(config *openapiutil.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapiutil.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = dara.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("intelligentcreation"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !dara.IsNil(endpoint) {
		_result = endpoint
		return _result, _err
	}

	if !dara.IsNil(endpointMap) && !dara.IsNil(endpointMap[dara.StringValue(regionId)]) {
		_result = endpointMap[dara.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := openapiutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

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
func (client *Client) AddTextFeedbackWithOptions(request *AddTextFeedbackRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddTextFeedbackResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加文案反馈
//
// @param request - AddTextFeedbackRequest
//
// @return AddTextFeedbackResponse
func (client *Client) AddTextFeedback(request *AddTextFeedbackRequest) (_result *AddTextFeedbackResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddTextFeedbackResponse{}
	_body, _err := client.AddTextFeedbackWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) BatchAddDocumentWithOptions(knowledgeBaseId *string, request *BatchAddDocumentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchAddDocumentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return BatchAddDocumentResponse
func (client *Client) BatchAddDocument(knowledgeBaseId *string, request *BatchAddDocumentRequest) (_result *BatchAddDocumentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchAddDocumentResponse{}
	_body, _err := client.BatchAddDocumentWithOptions(knowledgeBaseId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) BatchCreateAICoachTaskWithOptions(request *BatchCreateAICoachTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchCreateAICoachTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return BatchCreateAICoachTaskResponse
func (client *Client) BatchCreateAICoachTask(request *BatchCreateAICoachTaskRequest) (_result *BatchCreateAICoachTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchCreateAICoachTaskResponse{}
	_body, _err := client.BatchCreateAICoachTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) BatchDeletePracticeTaskWithOptions(tmpReq *BatchDeletePracticeTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchDeletePracticeTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - BatchDeletePracticeTaskRequest
//
// @return BatchDeletePracticeTaskResponse
func (client *Client) BatchDeletePracticeTask(request *BatchDeletePracticeTaskRequest) (_result *BatchDeletePracticeTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchDeletePracticeTaskResponse{}
	_body, _err := client.BatchDeletePracticeTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) BatchGetProjectTaskWithOptions(tmpReq *BatchGetProjectTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchGetProjectTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - BatchGetProjectTaskRequest
//
// @return BatchGetProjectTaskResponse
func (client *Client) BatchGetProjectTask(request *BatchGetProjectTaskRequest) (_result *BatchGetProjectTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchGetProjectTaskResponse{}
	_body, _err := client.BatchGetProjectTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) BatchGetTrainTaskWithOptions(tmpReq *BatchGetTrainTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchGetTrainTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - BatchGetTrainTaskRequest
//
// @return BatchGetTrainTaskResponse
func (client *Client) BatchGetTrainTask(request *BatchGetTrainTaskRequest) (_result *BatchGetTrainTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchGetTrainTaskResponse{}
	_body, _err := client.BatchGetTrainTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) BatchGetVideoClipTaskWithOptions(tmpReq *BatchGetVideoClipTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchGetVideoClipTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - BatchGetVideoClipTaskRequest
//
// @return BatchGetVideoClipTaskResponse
func (client *Client) BatchGetVideoClipTask(request *BatchGetVideoClipTaskRequest) (_result *BatchGetVideoClipTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchGetVideoClipTaskResponse{}
	_body, _err := client.BatchGetVideoClipTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) BatchQueryIndividuationTextWithOptions(tmpReq *BatchQueryIndividuationTextRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchQueryIndividuationTextResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - BatchQueryIndividuationTextRequest
//
// @return BatchQueryIndividuationTextResponse
func (client *Client) BatchQueryIndividuationText(request *BatchQueryIndividuationTextRequest) (_result *BatchQueryIndividuationTextResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchQueryIndividuationTextResponse{}
	_body, _err := client.BatchQueryIndividuationTextWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) BuildAICoachScriptRecordWithOptions(request *BuildAICoachScriptRecordRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BuildAICoachScriptRecordResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return BuildAICoachScriptRecordResponse
func (client *Client) BuildAICoachScriptRecord(request *BuildAICoachScriptRecordRequest) (_result *BuildAICoachScriptRecordResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BuildAICoachScriptRecordResponse{}
	_body, _err := client.BuildAICoachScriptRecordWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CheckSessionWithOptions(request *CheckSessionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CheckSessionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CheckSessionResponse
func (client *Client) CheckSession(request *CheckSessionRequest) (_result *CheckSessionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CheckSessionResponse{}
	_body, _err := client.CheckSessionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CloseAICoachTaskSessionWithOptions(request *CloseAICoachTaskSessionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CloseAICoachTaskSessionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CloseAICoachTaskSessionResponse
func (client *Client) CloseAICoachTaskSession(request *CloseAICoachTaskSessionRequest) (_result *CloseAICoachTaskSessionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CloseAICoachTaskSessionResponse{}
	_body, _err := client.CloseAICoachTaskSessionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CountTextWithOptions(request *CountTextRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CountTextResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CountTextResponse
func (client *Client) CountText(request *CountTextRequest) (_result *CountTextResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CountTextResponse{}
	_body, _err := client.CountTextWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateAICoachTaskWithOptions(request *CreateAICoachTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAICoachTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateAICoachTaskResponse
func (client *Client) CreateAICoachTask(request *CreateAICoachTaskRequest) (_result *CreateAICoachTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAICoachTaskResponse{}
	_body, _err := client.CreateAICoachTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateAICoachTaskReportWithOptions(request *CreateAICoachTaskReportRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAICoachTaskReportResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateAICoachTaskReportResponse
func (client *Client) CreateAICoachTaskReport(request *CreateAICoachTaskReportRequest) (_result *CreateAICoachTaskReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAICoachTaskReportResponse{}
	_body, _err := client.CreateAICoachTaskReportWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateAICoachTaskSessionWithOptions(request *CreateAICoachTaskSessionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAICoachTaskSessionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateAICoachTaskSessionResponse
func (client *Client) CreateAICoachTaskSession(request *CreateAICoachTaskSessionRequest) (_result *CreateAICoachTaskSessionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAICoachTaskSessionResponse{}
	_body, _err := client.CreateAICoachTaskSessionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateAgentWithOptions(request *CreateAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAgentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateAgentResponse
func (client *Client) CreateAgent(request *CreateAgentRequest) (_result *CreateAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAgentResponse{}
	_body, _err := client.CreateAgentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateAnchorWithOptions(request *CreateAnchorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAnchorResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateAnchorResponse
func (client *Client) CreateAnchor(request *CreateAnchorRequest) (_result *CreateAnchorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAnchorResponse{}
	_body, _err := client.CreateAnchorWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建文档生成剧本任务
//
// @param request - CreateGenerateAICoachScriptTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGenerateAICoachScriptTaskResponse
func (client *Client) CreateGenerateAICoachScriptTaskWithOptions(request *CreateGenerateAICoachScriptTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateGenerateAICoachScriptTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AssessmentPoint) {
		body["assessmentPoint"] = request.AssessmentPoint
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.DialogueKey) {
		body["dialogueKey"] = request.DialogueKey
	}

	if !dara.IsNil(request.DialogueUrl) {
		body["dialogueUrl"] = request.DialogueUrl
	}

	if !dara.IsNil(request.DocList) {
		body["docList"] = request.DocList
	}

	if !dara.IsNil(request.DocUrlList) {
		body["docUrlList"] = request.DocUrlList
	}

	if !dara.IsNil(request.ScriptName) {
		body["scriptName"] = request.ScriptName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateGenerateAICoachScriptTask"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/aicoach/scriptGenerateTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateGenerateAICoachScriptTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建文档生成剧本任务
//
// @param request - CreateGenerateAICoachScriptTaskRequest
//
// @return CreateGenerateAICoachScriptTaskResponse
func (client *Client) CreateGenerateAICoachScriptTask(request *CreateGenerateAICoachScriptTaskRequest) (_result *CreateGenerateAICoachScriptTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateGenerateAICoachScriptTaskResponse{}
	_body, _err := client.CreateGenerateAICoachScriptTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateIllustrationTaskWithOptions(textId *string, request *CreateIllustrationTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateIllustrationTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateIllustrationTaskResponse
func (client *Client) CreateIllustrationTask(textId *string, request *CreateIllustrationTaskRequest) (_result *CreateIllustrationTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateIllustrationTaskResponse{}
	_body, _err := client.CreateIllustrationTaskWithOptions(textId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateIndividuationProjectWithOptions(request *CreateIndividuationProjectRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateIndividuationProjectResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateIndividuationProjectResponse
func (client *Client) CreateIndividuationProject(request *CreateIndividuationProjectRequest) (_result *CreateIndividuationProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateIndividuationProjectResponse{}
	_body, _err := client.CreateIndividuationProjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateIndividuationTextTaskWithOptions(request *CreateIndividuationTextTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateIndividuationTextTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateIndividuationTextTaskResponse
func (client *Client) CreateIndividuationTextTask(request *CreateIndividuationTextTaskRequest) (_result *CreateIndividuationTextTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateIndividuationTextTaskResponse{}
	_body, _err := client.CreateIndividuationTextTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateProductImageWithOptions(request *CreateProductImageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateProductImageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateProductImageResponse
func (client *Client) CreateProductImage(request *CreateProductImageRequest) (_result *CreateProductImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateProductImageResponse{}
	_body, _err := client.CreateProductImageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateRealisticPortraitWithOptions(request *CreateRealisticPortraitRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateRealisticPortraitResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateRealisticPortraitResponse
func (client *Client) CreateRealisticPortrait(request *CreateRealisticPortraitRequest) (_result *CreateRealisticPortraitResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateRealisticPortraitResponse{}
	_body, _err := client.CreateRealisticPortraitWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateTextTaskWithOptions(request *CreateTextTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTextTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateTextTaskResponse
func (client *Client) CreateTextTask(request *CreateTextTaskRequest) (_result *CreateTextTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTextTaskResponse{}
	_body, _err := client.CreateTextTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateTrainTaskWithOptions(request *CreateTrainTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTrainTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateTrainTaskResponse
func (client *Client) CreateTrainTask(request *CreateTrainTaskRequest) (_result *CreateTrainTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTrainTaskResponse{}
	_body, _err := client.CreateTrainTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateVideoClipTaskWithOptions(request *CreateVideoClipTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateVideoClipTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateVideoClipTaskResponse
func (client *Client) CreateVideoClipTask(request *CreateVideoClipTaskRequest) (_result *CreateVideoClipTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateVideoClipTaskResponse{}
	_body, _err := client.CreateVideoClipTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteAICoachScriptWithOptions(request *DeleteAICoachScriptRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAICoachScriptResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteAICoachScriptResponse
func (client *Client) DeleteAICoachScript(request *DeleteAICoachScriptRequest) (_result *DeleteAICoachScriptResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAICoachScriptResponse{}
	_body, _err := client.DeleteAICoachScriptWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteAgentWithOptions(request *DeleteAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteAgentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteAgentResponse
func (client *Client) DeleteAgent(request *DeleteAgentRequest) (_result *DeleteAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAgentResponse{}
	_body, _err := client.DeleteAgentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteIndividuationProjectWithOptions(request *DeleteIndividuationProjectRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteIndividuationProjectResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteIndividuationProjectResponse
func (client *Client) DeleteIndividuationProject(request *DeleteIndividuationProjectRequest) (_result *DeleteIndividuationProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteIndividuationProjectResponse{}
	_body, _err := client.DeleteIndividuationProjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteIndividuationTextWithOptions(request *DeleteIndividuationTextRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteIndividuationTextResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteIndividuationTextResponse
func (client *Client) DeleteIndividuationText(request *DeleteIndividuationTextRequest) (_result *DeleteIndividuationTextResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteIndividuationTextResponse{}
	_body, _err := client.DeleteIndividuationTextWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DescribeDocumentWithOptions(knowledgeBaseId *string, documentId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DescribeDocumentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DescribeDocumentResponse
func (client *Client) DescribeDocument(knowledgeBaseId *string, documentId *string) (_result *DescribeDocumentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeDocumentResponse{}
	_body, _err := client.DescribeDocumentWithOptions(knowledgeBaseId, documentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) FinishAICoachTaskSessionWithOptions(request *FinishAICoachTaskSessionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *FinishAICoachTaskSessionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return FinishAICoachTaskSessionResponse
func (client *Client) FinishAICoachTaskSession(request *FinishAICoachTaskSessionRequest) (_result *FinishAICoachTaskSessionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &FinishAICoachTaskSessionResponse{}
	_body, _err := client.FinishAICoachTaskSessionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetAICoachAssessmentPointWithOptions(request *GetAICoachAssessmentPointRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAICoachAssessmentPointResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetAICoachAssessmentPointResponse
func (client *Client) GetAICoachAssessmentPoint(request *GetAICoachAssessmentPointRequest) (_result *GetAICoachAssessmentPointResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAICoachAssessmentPointResponse{}
	_body, _err := client.GetAICoachAssessmentPointWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetAICoachCheatDetectionWithOptions(request *GetAICoachCheatDetectionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAICoachCheatDetectionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetAICoachCheatDetectionResponse
func (client *Client) GetAICoachCheatDetection(request *GetAICoachCheatDetectionRequest) (_result *GetAICoachCheatDetectionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAICoachCheatDetectionResponse{}
	_body, _err := client.GetAICoachCheatDetectionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查看剧本调试详情
//
// @param request - GetAICoachDebugResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAICoachDebugResultResponse
func (client *Client) GetAICoachDebugResultWithOptions(request *GetAICoachDebugResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAICoachDebugResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DataId) {
		query["dataId"] = request.DataId
	}

	if !dara.IsNil(request.DataType) {
		query["dataType"] = request.DataType
	}

	if !dara.IsNil(request.ScriptDebugId) {
		query["scriptDebugId"] = request.ScriptDebugId
	}

	if !dara.IsNil(request.ScriptRecordId) {
		query["scriptRecordId"] = request.ScriptRecordId
	}

	if !dara.IsNil(request.ScriptSnapshotId) {
		query["scriptSnapshotId"] = request.ScriptSnapshotId
	}

	if !dara.IsNil(request.TaskId) {
		query["taskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAICoachDebugResult"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/aicoach/getDebugResult"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAICoachDebugResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看剧本调试详情
//
// @param request - GetAICoachDebugResultRequest
//
// @return GetAICoachDebugResultResponse
func (client *Client) GetAICoachDebugResult(request *GetAICoachDebugResultRequest) (_result *GetAICoachDebugResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAICoachDebugResultResponse{}
	_body, _err := client.GetAICoachDebugResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetAICoachScriptWithOptions(request *GetAICoachScriptRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAICoachScriptResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetAICoachScriptResponse
func (client *Client) GetAICoachScript(request *GetAICoachScriptRequest) (_result *GetAICoachScriptResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAICoachScriptResponse{}
	_body, _err := client.GetAICoachScriptWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询文档生成剧本任务结果
//
// @param request - GetAICoachScriptGenerateTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAICoachScriptGenerateTaskResponse
func (client *Client) GetAICoachScriptGenerateTaskWithOptions(request *GetAICoachScriptGenerateTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAICoachScriptGenerateTaskResponse, _err error) {
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
		Action:      dara.String("GetAICoachScriptGenerateTask"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/aicoach/scriptGenerateTask"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAICoachScriptGenerateTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询文档生成剧本任务结果
//
// @param request - GetAICoachScriptGenerateTaskRequest
//
// @return GetAICoachScriptGenerateTaskResponse
func (client *Client) GetAICoachScriptGenerateTask(request *GetAICoachScriptGenerateTaskRequest) (_result *GetAICoachScriptGenerateTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAICoachScriptGenerateTaskResponse{}
	_body, _err := client.GetAICoachScriptGenerateTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetAICoachTaskSessionHistoryWithOptions(request *GetAICoachTaskSessionHistoryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAICoachTaskSessionHistoryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetAICoachTaskSessionHistoryResponse
func (client *Client) GetAICoachTaskSessionHistory(request *GetAICoachTaskSessionHistoryRequest) (_result *GetAICoachTaskSessionHistoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAICoachTaskSessionHistoryResponse{}
	_body, _err := client.GetAICoachTaskSessionHistoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetAICoachTaskSessionReportWithOptions(request *GetAICoachTaskSessionReportRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAICoachTaskSessionReportResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetAICoachTaskSessionReportResponse
func (client *Client) GetAICoachTaskSessionReport(request *GetAICoachTaskSessionReportRequest) (_result *GetAICoachTaskSessionReportResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAICoachTaskSessionReportResponse{}
	_body, _err := client.GetAICoachTaskSessionReportWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetAICoachTaskSessionResourceUsageWithOptions(request *GetAICoachTaskSessionResourceUsageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAICoachTaskSessionResourceUsageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetAICoachTaskSessionResourceUsageResponse
func (client *Client) GetAICoachTaskSessionResourceUsage(request *GetAICoachTaskSessionResourceUsageRequest) (_result *GetAICoachTaskSessionResourceUsageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAICoachTaskSessionResourceUsageResponse{}
	_body, _err := client.GetAICoachTaskSessionResourceUsageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetIllustrationWithOptions(textId *string, illustrationId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetIllustrationResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetIllustrationResponse
func (client *Client) GetIllustration(textId *string, illustrationId *string) (_result *GetIllustrationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetIllustrationResponse{}
	_body, _err := client.GetIllustrationWithOptions(textId, illustrationId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetIllustrationTaskWithOptions(textId *string, illustrationTaskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetIllustrationTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetIllustrationTaskResponse
func (client *Client) GetIllustrationTask(textId *string, illustrationTaskId *string) (_result *GetIllustrationTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetIllustrationTaskResponse{}
	_body, _err := client.GetIllustrationTaskWithOptions(textId, illustrationTaskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetOssUploadTokenWithOptions(request *GetOssUploadTokenRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetOssUploadTokenResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetOssUploadTokenResponse
func (client *Client) GetOssUploadToken(request *GetOssUploadTokenRequest) (_result *GetOssUploadTokenResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetOssUploadTokenResponse{}
	_body, _err := client.GetOssUploadTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetProjectTaskWithOptions(request *GetProjectTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetProjectTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetProjectTaskResponse
func (client *Client) GetProjectTask(request *GetProjectTaskRequest) (_result *GetProjectTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetProjectTaskResponse{}
	_body, _err := client.GetProjectTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetTextWithOptions(textId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTextResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetTextResponse
func (client *Client) GetText(textId *string) (_result *GetTextResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTextResponse{}
	_body, _err := client.GetTextWithOptions(textId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetTextTaskWithOptions(textTaskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTextTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetTextTaskResponse
func (client *Client) GetTextTask(textTaskId *string) (_result *GetTextTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTextTaskResponse{}
	_body, _err := client.GetTextTaskWithOptions(textTaskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetTextTemplateWithOptions(request *GetTextTemplateRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTextTemplateResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetTextTemplateResponse
func (client *Client) GetTextTemplate(request *GetTextTemplateRequest) (_result *GetTextTemplateResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTextTemplateResponse{}
	_body, _err := client.GetTextTemplateWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) InteractTextWithSSE(request *InteractTextRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *InteractTextResponse, _yieldErr chan error) {
	defer close(_yield)
	client.interactTextWithSSE_opYieldFunc(_yield, _yieldErr, request, headers, runtime)
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
func (client *Client) InteractTextWithOptions(request *InteractTextRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InteractTextResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return InteractTextResponse
func (client *Client) InteractText(request *InteractTextRequest) (_result *InteractTextResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InteractTextResponse{}
	_body, _err := client.InteractTextWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListAICoachScriptPageWithOptions(request *ListAICoachScriptPageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAICoachScriptPageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListAICoachScriptPageResponse
func (client *Client) ListAICoachScriptPage(request *ListAICoachScriptPageRequest) (_result *ListAICoachScriptPageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAICoachScriptPageResponse{}
	_body, _err := client.ListAICoachScriptPageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListAICoachTaskPageWithOptions(request *ListAICoachTaskPageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAICoachTaskPageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListAICoachTaskPageResponse
func (client *Client) ListAICoachTaskPage(request *ListAICoachTaskPageRequest) (_result *ListAICoachTaskPageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAICoachTaskPageResponse{}
	_body, _err := client.ListAICoachTaskPageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据剧本对练任务查询会话历史
//
// @param request - ListAICoachTaskSessionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAICoachTaskSessionResponse
func (client *Client) ListAICoachTaskSessionWithOptions(request *ListAICoachTaskSessionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAICoachTaskSessionResponse, _err error) {
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

	if !dara.IsNil(request.TaskId) {
		query["taskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListAICoachTaskSession"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/aicoach/listTaskSession"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListAICoachTaskSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据剧本对练任务查询会话历史
//
// @param request - ListAICoachTaskSessionRequest
//
// @return ListAICoachTaskSessionResponse
func (client *Client) ListAICoachTaskSession(request *ListAICoachTaskSessionRequest) (_result *ListAICoachTaskSessionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAICoachTaskSessionResponse{}
	_body, _err := client.ListAICoachTaskSessionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListAgentsWithOptions(request *ListAgentsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAgentsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListAgentsResponse
func (client *Client) ListAgents(request *ListAgentsRequest) (_result *ListAgentsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAgentsResponse{}
	_body, _err := client.ListAgentsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListAnchorWithOptions(request *ListAnchorRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAnchorResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListAnchorResponse
func (client *Client) ListAnchor(request *ListAnchorRequest) (_result *ListAnchorResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAnchorResponse{}
	_body, _err := client.ListAnchorWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListAvatarProjectWithOptions(tmpReq *ListAvatarProjectRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListAvatarProjectResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - ListAvatarProjectRequest
//
// @return ListAvatarProjectResponse
func (client *Client) ListAvatarProject(request *ListAvatarProjectRequest) (_result *ListAvatarProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAvatarProjectResponse{}
	_body, _err := client.ListAvatarProjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListKnowledgeBaseWithOptions(request *ListKnowledgeBaseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListKnowledgeBaseResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListKnowledgeBaseResponse
func (client *Client) ListKnowledgeBase(request *ListKnowledgeBaseRequest) (_result *ListKnowledgeBaseResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListKnowledgeBaseResponse{}
	_body, _err := client.ListKnowledgeBaseWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListTextThemesWithOptions(request *ListTextThemesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTextThemesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListTextThemesResponse
func (client *Client) ListTextThemes(request *ListTextThemesRequest) (_result *ListTextThemesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTextThemesResponse{}
	_body, _err := client.ListTextThemesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListTextsWithOptions(request *ListTextsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListTextsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListTextsResponse
func (client *Client) ListTexts(request *ListTextsRequest) (_result *ListTextsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTextsResponse{}
	_body, _err := client.ListTextsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ListVoiceModelsWithOptions(request *ListVoiceModelsRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListVoiceModelsResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ListVoiceModelsResponse
func (client *Client) ListVoiceModels(request *ListVoiceModelsRequest) (_result *ListVoiceModelsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListVoiceModelsResponse{}
	_body, _err := client.ListVoiceModelsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) OfflineAICoachScriptWithOptions(request *OfflineAICoachScriptRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OfflineAICoachScriptResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return OfflineAICoachScriptResponse
func (client *Client) OfflineAICoachScript(request *OfflineAICoachScriptRequest) (_result *OfflineAICoachScriptResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OfflineAICoachScriptResponse{}
	_body, _err := client.OfflineAICoachScriptWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) OperateAvatarProjectWithOptions(request *OperateAvatarProjectRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *OperateAvatarProjectResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return OperateAvatarProjectResponse
func (client *Client) OperateAvatarProject(request *OperateAvatarProjectRequest) (_result *OperateAvatarProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OperateAvatarProjectResponse{}
	_body, _err := client.OperateAvatarProjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryAvatarProjectWithOptions(request *QueryAvatarProjectRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryAvatarProjectResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryAvatarProjectResponse
func (client *Client) QueryAvatarProject(request *QueryAvatarProjectRequest) (_result *QueryAvatarProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryAvatarProjectResponse{}
	_body, _err := client.QueryAvatarProjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryAvatarResourceWithOptions(request *QueryAvatarResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryAvatarResourceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryAvatarResourceResponse
func (client *Client) QueryAvatarResource(request *QueryAvatarResourceRequest) (_result *QueryAvatarResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryAvatarResourceResponse{}
	_body, _err := client.QueryAvatarResourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryImageToVideoTaskWithOptions(request *QueryImageToVideoTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryImageToVideoTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryImageToVideoTaskResponse
func (client *Client) QueryImageToVideoTask(request *QueryImageToVideoTaskRequest) (_result *QueryImageToVideoTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryImageToVideoTaskResponse{}
	_body, _err := client.QueryImageToVideoTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryIndividuationTextTaskWithOptions(request *QueryIndividuationTextTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryIndividuationTextTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryIndividuationTextTaskResponse
func (client *Client) QueryIndividuationTextTask(request *QueryIndividuationTextTaskRequest) (_result *QueryIndividuationTextTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryIndividuationTextTaskResponse{}
	_body, _err := client.QueryIndividuationTextTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QuerySessionInfoWithOptions(tmpReq *QuerySessionInfoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QuerySessionInfoResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @param request - QuerySessionInfoRequest
//
// @return QuerySessionInfoResponse
func (client *Client) QuerySessionInfo(request *QuerySessionInfoRequest) (_result *QuerySessionInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QuerySessionInfoResponse{}
	_body, _err := client.QuerySessionInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) QueryTextStreamWithSSE(textId *string, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *QueryTextStreamResponse, _yieldErr chan error) {
	defer close(_yield)
	client.queryTextStreamWithSSE_opYieldFunc(_yield, _yieldErr, textId, headers, runtime)
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
func (client *Client) QueryTextStreamWithOptions(textId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *QueryTextStreamResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return QueryTextStreamResponse
func (client *Client) QueryTextStream(textId *string) (_result *QueryTextStreamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryTextStreamResponse{}
	_body, _err := client.QueryTextStreamWithOptions(textId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ReleaseAgentWithOptions(request *ReleaseAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ReleaseAgentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ReleaseAgentResponse
func (client *Client) ReleaseAgent(request *ReleaseAgentRequest) (_result *ReleaseAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ReleaseAgentResponse{}
	_body, _err := client.ReleaseAgentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SaveAvatarProjectWithOptions(request *SaveAvatarProjectRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SaveAvatarProjectResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SaveAvatarProjectResponse
func (client *Client) SaveAvatarProject(request *SaveAvatarProjectRequest) (_result *SaveAvatarProjectResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SaveAvatarProjectResponse{}
	_body, _err := client.SaveAvatarProjectWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SelectImageTaskWithOptions(taskId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SelectImageTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SelectImageTaskResponse
func (client *Client) SelectImageTask(taskId *string) (_result *SelectImageTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SelectImageTaskResponse{}
	_body, _err := client.SelectImageTaskWithOptions(taskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SelectResourceWithOptions(request *SelectResourceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SelectResourceResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SelectResourceResponse
func (client *Client) SelectResource(request *SelectResourceRequest) (_result *SelectResourceResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SelectResourceResponse{}
	_body, _err := client.SelectResourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SendSdkMessageWithOptions(request *SendSdkMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SendSdkMessageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SendSdkMessageResponse
func (client *Client) SendSdkMessage(request *SendSdkMessageRequest) (_result *SendSdkMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SendSdkMessageResponse{}
	_body, _err := client.SendSdkMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SendSdkStreamMessageWithSSE(request *SendSdkStreamMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *SendSdkStreamMessageResponse, _yieldErr chan error) {
	defer close(_yield)
	client.sendSdkStreamMessageWithSSE_opYieldFunc(_yield, _yieldErr, request, headers, runtime)
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
func (client *Client) SendSdkStreamMessageWithOptions(request *SendSdkStreamMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SendSdkStreamMessageResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SendSdkStreamMessageResponse
func (client *Client) SendSdkStreamMessage(request *SendSdkStreamMessageRequest) (_result *SendSdkStreamMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SendSdkStreamMessageResponse{}
	_body, _err := client.SendSdkStreamMessageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SendTextMsgWithOptions(request *SendTextMsgRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SendTextMsgResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SendTextMsgResponse
func (client *Client) SendTextMsg(request *SendTextMsgRequest) (_result *SendTextMsgResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SendTextMsgResponse{}
	_body, _err := client.SendTextMsgWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) StartAvatarSessionWithOptions(request *StartAvatarSessionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StartAvatarSessionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return StartAvatarSessionResponse
func (client *Client) StartAvatarSession(request *StartAvatarSessionRequest) (_result *StartAvatarSessionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartAvatarSessionResponse{}
	_body, _err := client.StartAvatarSessionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) StopAvatarSessionWithOptions(request *StopAvatarSessionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopAvatarSessionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return StopAvatarSessionResponse
func (client *Client) StopAvatarSession(request *StopAvatarSessionRequest) (_result *StopAvatarSessionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopAvatarSessionResponse{}
	_body, _err := client.StopAvatarSessionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) StopProjectTaskWithOptions(request *StopProjectTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *StopProjectTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return StopProjectTaskResponse
func (client *Client) StopProjectTask(request *StopProjectTaskRequest) (_result *StopProjectTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopProjectTaskResponse{}
	_body, _err := client.StopProjectTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 提交剧本考核点调试
//
// @param request - SubmitAICoachDebugRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitAICoachDebugResponse
func (client *Client) SubmitAICoachDebugWithOptions(request *SubmitAICoachDebugRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitAICoachDebugResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DataId) {
		body["dataId"] = request.DataId
	}

	if !dara.IsNil(request.DataType) {
		body["dataType"] = request.DataType
	}

	if !dara.IsNil(request.DeductionRule) {
		body["deductionRule"] = request.DeductionRule
	}

	if !dara.IsNil(request.DialogueList) {
		body["dialogueList"] = request.DialogueList
	}

	if !dara.IsNil(request.Expressiveness) {
		body["expressiveness"] = request.Expressiveness
	}

	if !dara.IsNil(request.Point) {
		body["point"] = request.Point
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitAICoachDebug"),
		Version:     dara.String("2024-03-13"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/yic/yic-console/openService/v1/aicoach/saveDebug"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitAICoachDebugResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交剧本考核点调试
//
// @param request - SubmitAICoachDebugRequest
//
// @return SubmitAICoachDebugResponse
func (client *Client) SubmitAICoachDebug(request *SubmitAICoachDebugRequest) (_result *SubmitAICoachDebugResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitAICoachDebugResponse{}
	_body, _err := client.SubmitAICoachDebugWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SubmitImageToVideoTaskWithOptions(request *SubmitImageToVideoTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitImageToVideoTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SubmitImageToVideoTaskResponse
func (client *Client) SubmitImageToVideoTask(request *SubmitImageToVideoTaskRequest) (_result *SubmitImageToVideoTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitImageToVideoTaskResponse{}
	_body, _err := client.SubmitImageToVideoTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SubmitProjectTaskWithOptions(request *SubmitProjectTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitProjectTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SubmitProjectTaskResponse
func (client *Client) SubmitProjectTask(request *SubmitProjectTaskRequest) (_result *SubmitProjectTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitProjectTaskResponse{}
	_body, _err := client.SubmitProjectTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) TransferPortraitStyleWithOptions(request *TransferPortraitStyleRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *TransferPortraitStyleResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return TransferPortraitStyleResponse
func (client *Client) TransferPortraitStyle(request *TransferPortraitStyleRequest) (_result *TransferPortraitStyleResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TransferPortraitStyleResponse{}
	_body, _err := client.TransferPortraitStyleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) UpdateAgentWithOptions(request *UpdateAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateAgentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return UpdateAgentResponse
func (client *Client) UpdateAgent(request *UpdateAgentRequest) (_result *UpdateAgentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateAgentResponse{}
	_body, _err := client.UpdateAgentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) interactTextWithSSE_opYieldFunc(_yield chan *InteractTextResponse, _yieldErr chan error, request *InteractTextRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) queryTextStreamWithSSE_opYieldFunc(_yield chan *QueryTextStreamResponse, _yieldErr chan error, textId *string, headers map[string]*string, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) sendSdkStreamMessageWithSSE_opYieldFunc(_yield chan *SendSdkStreamMessageResponse, _yieldErr chan error, request *SendSdkStreamMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
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
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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
