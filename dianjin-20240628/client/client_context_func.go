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
// 创建按年文档总结任务
//
// @param request - CreateAnnualDocSummaryTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAnnualDocSummaryTaskResponse
func (client *Client) CreateAnnualDocSummaryTaskWithContext(ctx context.Context, workspaceId *string, request *CreateAnnualDocSummaryTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateAnnualDocSummaryTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AnaYears) {
		body["anaYears"] = request.AnaYears
	}

	if !dara.IsNil(request.DocInfos) {
		body["docInfos"] = request.DocInfos
	}

	if !dara.IsNil(request.EnableTable) {
		body["enableTable"] = request.EnableTable
	}

	if !dara.IsNil(request.Instruction) {
		body["instruction"] = request.Instruction
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateAnnualDocSummaryTask"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/task/summary/doc/annual"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateAnnualDocSummaryTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建外呼会话
//
// @param request - CreateDialogRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDialogResponse
func (client *Client) CreateDialogWithContext(ctx context.Context, workspaceId *string, request *CreateDialogRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDialogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Channel) {
		body["channel"] = request.Channel
	}

	if !dara.IsNil(request.EnableLibrary) {
		body["enableLibrary"] = request.EnableLibrary
	}

	if !dara.IsNil(request.MetaData) {
		body["metaData"] = request.MetaData
	}

	if !dara.IsNil(request.PlayCode) {
		body["playCode"] = request.PlayCode
	}

	if !dara.IsNil(request.QaLibraryList) {
		body["qaLibraryList"] = request.QaLibraryList
	}

	if !dara.IsNil(request.RequestId) {
		body["requestId"] = request.RequestId
	}

	if !dara.IsNil(request.SelfDirected) {
		body["selfDirected"] = request.SelfDirected
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDialog"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/virtualHuman/dialog/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDialogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建会话分析任务
//
// @param request - CreateDialogAnalysisTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDialogAnalysisTaskResponse
func (client *Client) CreateDialogAnalysisTaskWithContext(ctx context.Context, workspaceId *string, request *CreateDialogAnalysisTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDialogAnalysisTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AnalysisNodes) {
		body["analysisNodes"] = request.AnalysisNodes
	}

	if !dara.IsNil(request.ConversationList) {
		body["conversationList"] = request.ConversationList
	}

	if !dara.IsNil(request.MetaData) {
		body["metaData"] = request.MetaData
	}

	if !dara.IsNil(request.PlayCode) {
		body["playCode"] = request.PlayCode
	}

	if !dara.IsNil(request.RequestId) {
		body["requestId"] = request.RequestId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDialogAnalysisTask"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/virtualHuman/dialog/analysis/submit"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDialogAnalysisTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建财报总结任务
//
// @param request - CreateDocsSummaryTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDocsSummaryTaskResponse
func (client *Client) CreateDocsSummaryTaskWithContext(ctx context.Context, workspaceId *string, request *CreateDocsSummaryTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateDocsSummaryTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocInfos) {
		body["docInfos"] = request.DocInfos
	}

	if !dara.IsNil(request.EnableTable) {
		body["enableTable"] = request.EnableTable
	}

	if !dara.IsNil(request.Instruction) {
		body["instruction"] = request.Instruction
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDocsSummaryTask"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/task/summary/docs"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateDocsSummaryTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建财报总结任务
//
// @param request - CreateFinReportSummaryTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFinReportSummaryTaskResponse
func (client *Client) CreateFinReportSummaryTaskWithContext(ctx context.Context, workspaceId *string, request *CreateFinReportSummaryTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateFinReportSummaryTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocId) {
		body["docId"] = request.DocId
	}

	if !dara.IsNil(request.EnableTable) {
		body["enableTable"] = request.EnableTable
	}

	if !dara.IsNil(request.EndPage) {
		body["endPage"] = request.EndPage
	}

	if !dara.IsNil(request.Instruction) {
		body["instruction"] = request.Instruction
	}

	if !dara.IsNil(request.LibraryId) {
		body["libraryId"] = request.LibraryId
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.StartPage) {
		body["startPage"] = request.StartPage
	}

	if !dara.IsNil(request.TaskType) {
		body["taskType"] = request.TaskType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateFinReportSummaryTask"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/task/summary"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateFinReportSummaryTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建图片检测任务
//
// @param request - CreateImageDetectionTaskRequest
//
// @param headers - CreateImageDetectionTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateImageDetectionTaskResponse
func (client *Client) CreateImageDetectionTaskWithContext(ctx context.Context, workspaceId *string, request *CreateImageDetectionTaskRequest, headers *CreateImageDetectionTaskHeaders, runtime *dara.RuntimeOptions) (_result *CreateImageDetectionTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FileInfo) {
		body["fileInfo"] = request.FileInfo
	}

	if !dara.IsNil(request.FileUrl) {
		body["fileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.RequestId) {
		body["requestId"] = request.RequestId
	}

	if !dara.IsNil(request.UserId) {
		body["userId"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XLoadTest) {
		realHeaders["X-Load-Test"] = dara.String(dara.Stringify(dara.BoolValue(headers.XLoadTest)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateImageDetectionTask"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/imageDetect/task/submit"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateImageDetectionTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建文档库
//
// @param request - CreateLibraryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLibraryResponse
func (client *Client) CreateLibraryWithContext(ctx context.Context, workspaceId *string, request *CreateLibraryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateLibraryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.IndexSetting) {
		body["indexSetting"] = request.IndexSetting
	}

	if !dara.IsNil(request.LibraryName) {
		body["libraryName"] = request.LibraryName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateLibrary"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/library/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateLibraryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建PDF翻译任务
//
// @param request - CreatePdfTranslateTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePdfTranslateTaskResponse
func (client *Client) CreatePdfTranslateTaskWithContext(ctx context.Context, workspaceId *string, request *CreatePdfTranslateTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePdfTranslateTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocId) {
		body["docId"] = request.DocId
	}

	if !dara.IsNil(request.Knowledge) {
		body["knowledge"] = request.Knowledge
	}

	if !dara.IsNil(request.LibraryId) {
		body["libraryId"] = request.LibraryId
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.TranslateTo) {
		body["translateTo"] = request.TranslateTo
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePdfTranslateTask"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/task/pdfTranslate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePdfTranslateTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建预定义文档
//
// @param request - CreatePredefinedDocumentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePredefinedDocumentResponse
func (client *Client) CreatePredefinedDocumentWithContext(ctx context.Context, workspaceId *string, request *CreatePredefinedDocumentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreatePredefinedDocumentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Chunks) {
		body["chunks"] = request.Chunks
	}

	if !dara.IsNil(request.LibraryId) {
		body["libraryId"] = request.LibraryId
	}

	if !dara.IsNil(request.Metadata) {
		body["metadata"] = request.Metadata
	}

	if !dara.IsNil(request.Title) {
		body["title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreatePredefinedDocument"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/library/document/createPredefinedDocument"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreatePredefinedDocumentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建财报总结的任务
//
// @param request - CreateQualityCheckTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQualityCheckTaskResponse
func (client *Client) CreateQualityCheckTaskWithContext(ctx context.Context, workspaceId *string, request *CreateQualityCheckTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateQualityCheckTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ConversationList) {
		body["conversationList"] = request.ConversationList
	}

	if !dara.IsNil(request.GmtService) {
		body["gmtService"] = request.GmtService
	}

	if !dara.IsNil(request.MetaData) {
		body["metaData"] = request.MetaData
	}

	if !dara.IsNil(request.QualityGroup) {
		body["qualityGroup"] = request.QualityGroup
	}

	if !dara.IsNil(request.RequestId) {
		body["requestId"] = request.RequestId
	}

	if !dara.IsNil(request.SceneCode) {
		body["sceneCode"] = request.SceneCode
	}

	if !dara.IsNil(request.Type) {
		body["type"] = request.Type
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateQualityCheckTask"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/qualitycheck/task/submit"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateQualityCheckTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建视频生成任务
//
// @param request - CreateVideoCreationTaskRequest
//
// @param headers - CreateVideoCreationTaskHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVideoCreationTaskResponse
func (client *Client) CreateVideoCreationTaskWithContext(ctx context.Context, workspaceId *string, request *CreateVideoCreationTaskRequest, headers *CreateVideoCreationTaskHeaders, runtime *dara.RuntimeOptions) (_result *CreateVideoCreationTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CreationInstruction) {
		body["creationInstruction"] = request.CreationInstruction
	}

	if !dara.IsNil(request.FileInfo) {
		body["fileInfo"] = request.FileInfo
	}

	if !dara.IsNil(request.ImageDetectionTaskId) {
		body["imageDetectionTaskId"] = request.ImageDetectionTaskId
	}

	if !dara.IsNil(request.RequestId) {
		body["requestId"] = request.RequestId
	}

	if !dara.IsNil(request.UserId) {
		body["userId"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XLoadTest) {
		realHeaders["X-Load-Test"] = dara.String(dara.Stringify(dara.BoolValue(headers.XLoadTest)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVideoCreationTask"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/videoCreation/task/create"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVideoCreationTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除文档
//
// @param request - DeleteDocumentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDocumentResponse
func (client *Client) DeleteDocumentWithContext(ctx context.Context, workspaceId *string, request *DeleteDocumentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDocumentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocIds) {
		body["docIds"] = request.DocIds
	}

	if !dara.IsNil(request.LibraryId) {
		body["libraryId"] = request.LibraryId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteDocument"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/library/document/delete"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteDocumentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除文档库
//
// @param request - DeleteLibraryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteLibraryResponse
func (client *Client) DeleteLibraryWithContext(ctx context.Context, workspaceId *string, request *DeleteLibraryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteLibraryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LibraryId) {
		query["libraryId"] = request.LibraryId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteLibrary"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/library/delete"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteLibraryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 端到端实时对话
//
// @param request - EndToEndRealTimeDialogRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EndToEndRealTimeDialogResponse
func (client *Client) EndToEndRealTimeDialogWithContext(ctx context.Context, workspaceId *string, request *EndToEndRealTimeDialogRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EndToEndRealTimeDialogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AsrModelId) {
		query["asrModelId"] = request.AsrModelId
	}

	if !dara.IsNil(request.InputFormat) {
		query["inputFormat"] = request.InputFormat
	}

	if !dara.IsNil(request.OutputFormat) {
		query["outputFormat"] = request.OutputFormat
	}

	if !dara.IsNil(request.PitchRate) {
		query["pitchRate"] = request.PitchRate
	}

	if !dara.IsNil(request.SampleRate) {
		query["sampleRate"] = request.SampleRate
	}

	if !dara.IsNil(request.SpeechRate) {
		query["speechRate"] = request.SpeechRate
	}

	if !dara.IsNil(request.TtsModelId) {
		query["ttsModelId"] = request.TtsModelId
	}

	if !dara.IsNil(request.VoiceCode) {
		query["voiceCode"] = request.VoiceCode
	}

	if !dara.IsNil(request.Volume) {
		query["volume"] = request.Volume
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EndToEndRealTimeDialog"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/ws/realtime/dialog"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &EndToEndRealTimeDialogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 中断任务
//
// @param request - EvictTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EvictTaskResponse
func (client *Client) EvictTaskWithContext(ctx context.Context, workspaceId *string, request *EvictTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EvictTaskResponse, _err error) {
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
		Action:      dara.String("EvictTask"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/task/evict"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &EvictTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据文档解析问答QA
//
// @param request - GenDocQaResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenDocQaResultResponse
func (client *Client) GenDocQaResultWithContext(ctx context.Context, workspaceId *string, request *GenDocQaResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GenDocQaResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocId) {
		body["docId"] = request.DocId
	}

	if !dara.IsNil(request.LibraryId) {
		body["libraryId"] = request.LibraryId
	}

	if !dara.IsNil(request.RequestId) {
		body["requestId"] = request.RequestId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenDocQaResult"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/virtualHuman/qa/parse"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GenDocQaResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取app配置
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppConfigResponse
func (client *Client) GetAppConfigWithContext(ctx context.Context, workspaceId *string, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAppConfigResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAppConfig"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/app/config"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAppConfigResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取问答结果
//
// @param request - GetChatQuestionRespRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChatQuestionRespResponse
func (client *Client) GetChatQuestionRespWithContext(ctx context.Context, workspaceId *string, request *GetChatQuestionRespRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetChatQuestionRespResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BatchId) {
		body["batchId"] = request.BatchId
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetChatQuestionResp"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/virtualHuman/chat/query"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetChatQuestionRespResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取外呼会话分析结果
//
// @param request - GetDialogAnalysisResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDialogAnalysisResultResponse
func (client *Client) GetDialogAnalysisResultWithContext(ctx context.Context, workspaceId *string, request *GetDialogAnalysisResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDialogAnalysisResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Asc) {
		body["asc"] = request.Asc
	}

	if !dara.IsNil(request.EndTime) {
		body["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.SessionIds) {
		body["sessionIds"] = request.SessionIds
	}

	if !dara.IsNil(request.StartTime) {
		body["startTime"] = request.StartTime
	}

	if !dara.IsNil(request.UseUrl) {
		body["useUrl"] = request.UseUrl
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDialogAnalysisResult"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/virtualHuman/dialog/analysis"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDialogAnalysisResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取异步任务的结果
//
// @param request - GetDialogDetailRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDialogDetailResponse
func (client *Client) GetDialogDetailWithContext(ctx context.Context, workspaceId *string, request *GetDialogDetailRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDialogDetailResponse, _err error) {
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
		Action:      dara.String("GetDialogDetail"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/virtualHuman/dialog/detail"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDialogDetailResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询会话日志
//
// @param request - GetDialogLogRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDialogLogResponse
func (client *Client) GetDialogLogWithContext(ctx context.Context, workspaceId *string, request *GetDialogLogRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDialogLogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Id) {
		body["id"] = request.Id
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDialogLog"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/dialog/log"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDialogLogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文档的chunk列表
//
// @param request - GetDocumentChunkListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentChunkListResponse
func (client *Client) GetDocumentChunkListWithContext(ctx context.Context, workspaceId *string, request *GetDocumentChunkListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDocumentChunkListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ChunkIdList) {
		body["chunkIdList"] = request.ChunkIdList
	}

	if !dara.IsNil(request.DocId) {
		body["docId"] = request.DocId
	}

	if !dara.IsNil(request.LibraryId) {
		body["libraryId"] = request.LibraryId
	}

	if !dara.IsNil(request.Order) {
		body["order"] = request.Order
	}

	if !dara.IsNil(request.OrderBy) {
		body["orderBy"] = request.OrderBy
	}

	if !dara.IsNil(request.Page) {
		body["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		body["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SearchQuery) {
		body["searchQuery"] = request.SearchQuery
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDocumentChunkList"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/library/getDocumentChunk"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDocumentChunkListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询文档库的文档列表
//
// @param request - GetDocumentListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentListResponse
func (client *Client) GetDocumentListWithContext(ctx context.Context, workspaceId *string, request *GetDocumentListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDocumentListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LibraryId) {
		query["libraryId"] = request.LibraryId
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
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
		Action:      dara.String("GetDocumentList"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/library/listDocument"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDocumentListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文档URL
//
// @param request - GetDocumentUrlRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentUrlResponse
func (client *Client) GetDocumentUrlWithContext(ctx context.Context, workspaceId *string, request *GetDocumentUrlRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDocumentUrlResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DocumentId) {
		query["documentId"] = request.DocumentId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDocumentUrl"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/library/document/url"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDocumentUrlResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 带条件的分页查询文档库的文档列表
//
// @param request - GetFilterDocumentListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFilterDocumentListResponse
func (client *Client) GetFilterDocumentListWithContext(ctx context.Context, workspaceId *string, request *GetFilterDocumentListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetFilterDocumentListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.And) {
		body["and"] = request.And
	}

	if !dara.IsNil(request.DocIdList) {
		body["docIdList"] = request.DocIdList
	}

	if !dara.IsNil(request.LibraryId) {
		body["libraryId"] = request.LibraryId
	}

	if !dara.IsNil(request.Or) {
		body["or"] = request.Or
	}

	if !dara.IsNil(request.Page) {
		body["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		body["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Status) {
		body["status"] = request.Status
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFilterDocumentList"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/library/filterDocument"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFilterDocumentListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询文档库列表
//
// @param request - GetHistoryListByBizTypeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetHistoryListByBizTypeResponse
func (client *Client) GetHistoryListByBizTypeWithContext(ctx context.Context, workspaceId *string, request *GetHistoryListByBizTypeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetHistoryListByBizTypeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.BizId) {
		query["bizId"] = request.BizId
	}

	if !dara.IsNil(request.BizType) {
		query["bizType"] = request.BizType
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetHistoryListByBizType"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/history/list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetHistoryListByBizTypeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取检测结果
//
// @param request - GetImageDetectionTaskResultRequest
//
// @param headers - GetImageDetectionTaskResultHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetImageDetectionTaskResultResponse
func (client *Client) GetImageDetectionTaskResultWithContext(ctx context.Context, workspaceId *string, request *GetImageDetectionTaskResultRequest, headers *GetImageDetectionTaskResultHeaders, runtime *dara.RuntimeOptions) (_result *GetImageDetectionTaskResultResponse, _err error) {
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

	if !dara.IsNil(request.UserId) {
		body["userId"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XLoadTest) {
		realHeaders["X-Load-Test"] = dara.String(dara.Stringify(dara.BoolValue(headers.XLoadTest)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetImageDetectionTaskResult"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/imageDetect/task/query"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetImageDetectionTaskResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文档库配置详情
//
// @param request - GetLibraryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLibraryResponse
func (client *Client) GetLibraryWithContext(ctx context.Context, workspaceId *string, request *GetLibraryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetLibraryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.LibraryId) {
		query["libraryId"] = request.LibraryId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLibrary"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/library/get"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLibraryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 分页查询文档库列表
//
// @param request - GetLibraryListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLibraryListResponse
func (client *Client) GetLibraryListWithContext(ctx context.Context, workspaceId *string, request *GetLibraryListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetLibraryListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.Query) {
		query["query"] = request.Query
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetLibraryList"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/library/list"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetLibraryListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取解析结果
//
// @param request - GetParseResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetParseResultResponse
func (client *Client) GetParseResultWithContext(ctx context.Context, workspaceId *string, request *GetParseResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetParseResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocId) {
		body["docId"] = request.DocId
	}

	if !dara.IsNil(request.LibraryId) {
		body["libraryId"] = request.LibraryId
	}

	if !dara.IsNil(request.UseUrlResult) {
		body["useUrlResult"] = request.UseUrlResult
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetParseResult"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/library/document/getParseResult"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetParseResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取异步任务的结果
//
// @param request - GetQualityCheckTaskResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetQualityCheckTaskResultResponse
func (client *Client) GetQualityCheckTaskResultWithContext(ctx context.Context, workspaceId *string, request *GetQualityCheckTaskResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetQualityCheckTaskResultResponse, _err error) {
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
		Action:      dara.String("GetQualityCheckTaskResult"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/qualitycheck/task/query"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetQualityCheckTaskResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取财报总结任务结果
//
// @param request - GetSummaryTaskResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSummaryTaskResultResponse
func (client *Client) GetSummaryTaskResultWithContext(ctx context.Context, workspaceId *string, request *GetSummaryTaskResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetSummaryTaskResultResponse, _err error) {
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
		Action:      dara.String("GetSummaryTaskResult"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/task/summary/result"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetSummaryTaskResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取异步任务结果
//
// @param request - GetTaskResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskResultResponse
func (client *Client) GetTaskResultWithContext(ctx context.Context, workspaceId *string, request *GetTaskResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTaskResultResponse, _err error) {
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
		Action:      dara.String("GetTaskResult"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/task/result"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTaskResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取财报总结任务结果
//
// @param request - GetTaskStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskStatusResponse
func (client *Client) GetTaskStatusWithContext(ctx context.Context, workspaceId *string, request *GetTaskStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTaskStatusResponse, _err error) {
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
		Action:      dara.String("GetTaskStatus"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/task/status"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTaskStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取视频生成任务结果
//
// @param request - GetVideoCreationTaskResultRequest
//
// @param headers - GetVideoCreationTaskResultHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVideoCreationTaskResultResponse
func (client *Client) GetVideoCreationTaskResultWithContext(ctx context.Context, workspaceId *string, request *GetVideoCreationTaskResultRequest, headers *GetVideoCreationTaskResultHeaders, runtime *dara.RuntimeOptions) (_result *GetVideoCreationTaskResultResponse, _err error) {
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

	if !dara.IsNil(request.UserId) {
		body["userId"] = request.UserId
	}

	realHeaders := make(map[string]*string)
	if !dara.IsNil(headers.CommonHeaders) {
		realHeaders = headers.CommonHeaders
	}

	if !dara.IsNil(headers.XLoadTest) {
		realHeaders["X-Load-Test"] = dara.String(dara.Stringify(dara.BoolValue(headers.XLoadTest)))
	}

	req := &openapiutil.OpenApiRequest{
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVideoCreationTaskResult"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/videoCreation/task/query"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVideoCreationTaskResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 插件调试接口
//
// @param request - InvokePluginRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvokePluginResponse
func (client *Client) InvokePluginWithContext(ctx context.Context, workspaceId *string, request *InvokePluginRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *InvokePluginResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Params) {
		body["params"] = request.Params
	}

	if !dara.IsNil(request.PluginId) {
		body["pluginId"] = request.PluginId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("InvokePlugin"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/plugin/invoke"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &InvokePluginResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文档预览
//
// @param request - PreviewDocumentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PreviewDocumentResponse
func (client *Client) PreviewDocumentWithContext(ctx context.Context, workspaceId *string, request *PreviewDocumentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PreviewDocumentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DocumentId) {
		query["documentId"] = request.DocumentId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PreviewDocument"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/library/document/preview"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PreviewDocumentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重新索引
//
// @param request - ReIndexRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReIndexResponse
func (client *Client) ReIndexWithContext(ctx context.Context, workspaceId *string, request *ReIndexRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ReIndexResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DocumentId) {
		query["documentId"] = request.DocumentId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReIndex"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/library/document/reIndex"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ReIndexResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实时对话
//
// @param request - RealTimeDialogRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RealTimeDialogResponse
func (client *Client) RealTimeDialogWithSSECtx(ctx context.Context, workspaceId *string, request *RealTimeDialogRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RealTimeDialogResponse, _yieldErr chan error) {
	defer close(_yield)
	client.realTimeDialogWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, workspaceId, request, headers, runtime)
	return
}

// Summary:
//
// 实时对话
//
// @param request - RealTimeDialogRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RealTimeDialogResponse
func (client *Client) RealTimeDialogWithContext(ctx context.Context, workspaceId *string, request *RealTimeDialogRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RealTimeDialogResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Analysis) {
		body["analysis"] = request.Analysis
	}

	if !dara.IsNil(request.BizType) {
		body["bizType"] = request.BizType
	}

	if !dara.IsNil(request.ConversationModel) {
		body["conversationModel"] = request.ConversationModel
	}

	if !dara.IsNil(request.DialogMemoryTurns) {
		body["dialogMemoryTurns"] = request.DialogMemoryTurns
	}

	if !dara.IsNil(request.MetaData) {
		body["metaData"] = request.MetaData
	}

	if !dara.IsNil(request.OpType) {
		body["opType"] = request.OpType
	}

	if !dara.IsNil(request.Recommend) {
		body["recommend"] = request.Recommend
	}

	if !dara.IsNil(request.ScriptContentPlayed) {
		body["scriptContentPlayed"] = request.ScriptContentPlayed
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	if !dara.IsNil(request.UserVad) {
		body["userVad"] = request.UserVad
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RealTimeDialog"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/realtime/dialog/chat"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RealTimeDialogResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 实时会话辅助
//
// @param request - RealtimeDialogAssistRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RealtimeDialogAssistResponse
func (client *Client) RealtimeDialogAssistWithContext(ctx context.Context, workspaceId *string, request *RealtimeDialogAssistRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RealtimeDialogAssistResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Analysis) {
		body["analysis"] = request.Analysis
	}

	if !dara.IsNil(request.BizType) {
		body["bizType"] = request.BizType
	}

	if !dara.IsNil(request.ConversationModel) {
		body["conversationModel"] = request.ConversationModel
	}

	if !dara.IsNil(request.DialogMemoryTurns) {
		body["dialogMemoryTurns"] = request.DialogMemoryTurns
	}

	if !dara.IsNil(request.HangUpDialog) {
		body["hangUpDialog"] = request.HangUpDialog
	}

	if !dara.IsNil(request.MetaData) {
		body["metaData"] = request.MetaData
	}

	if !dara.IsNil(request.RequestId) {
		body["requestId"] = request.RequestId
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RealtimeDialogAssist"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/realtime/dialog/assist"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RealtimeDialogAssistResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 重建任务
//
// @param request - RebuildTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RebuildTaskResponse
func (client *Client) RebuildTaskWithContext(ctx context.Context, workspaceId *string, request *RebuildTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RebuildTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskIds) {
		body["taskIds"] = request.TaskIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RebuildTask"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/task/rebuild"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RebuildTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文档召回。
//
// @param request - RecallDocumentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecallDocumentResponse
func (client *Client) RecallDocumentWithContext(ctx context.Context, workspaceId *string, request *RecallDocumentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RecallDocumentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Filters) {
		body["filters"] = request.Filters
	}

	if !dara.IsNil(request.Query) {
		body["query"] = request.Query
	}

	if !dara.IsNil(request.Rearrangement) {
		body["rearrangement"] = request.Rearrangement
	}

	if !dara.IsNil(request.TopK) {
		body["topK"] = request.TopK
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RecallDocument"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/library/recallDocument"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RecallDocumentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 意图识别
//
// @param request - RecognizeIntentionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecognizeIntentionResponse
func (client *Client) RecognizeIntentionWithContext(ctx context.Context, workspaceId *string, request *RecognizeIntentionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RecognizeIntentionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Analysis) {
		body["analysis"] = request.Analysis
	}

	if !dara.IsNil(request.BizType) {
		body["bizType"] = request.BizType
	}

	if !dara.IsNil(request.Conversation) {
		body["conversation"] = request.Conversation
	}

	if !dara.IsNil(request.GlobalIntentionList) {
		body["globalIntentionList"] = request.GlobalIntentionList
	}

	if !dara.IsNil(request.HierarchicalIntentionList) {
		body["hierarchicalIntentionList"] = request.HierarchicalIntentionList
	}

	if !dara.IsNil(request.IntentionDomainCode) {
		body["intentionDomainCode"] = request.IntentionDomainCode
	}

	if !dara.IsNil(request.IntentionList) {
		body["intentionList"] = request.IntentionList
	}

	if !dara.IsNil(request.OpType) {
		body["opType"] = request.OpType
	}

	if !dara.IsNil(request.Recommend) {
		body["recommend"] = request.Recommend
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RecognizeIntention"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/recog/intent"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RecognizeIntentionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 运行智能体
//
// @param request - RunAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunAgentResponse
func (client *Client) RunAgentWithSSECtx(ctx context.Context, workspaceId *string, request *RunAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunAgentResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runAgentWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, workspaceId, request, headers, runtime)
	return
}

// Summary:
//
// 运行智能体
//
// @param request - RunAgentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunAgentResponse
func (client *Client) RunAgentWithContext(ctx context.Context, workspaceId *string, request *RunAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunAgentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BotId) {
		body["botId"] = request.BotId
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	if !dara.IsNil(request.ThreadId) {
		body["threadId"] = request.ThreadId
	}

	if !dara.IsNil(request.UseDraft) {
		body["useDraft"] = request.UseDraft
	}

	if !dara.IsNil(request.UserContent) {
		body["userContent"] = request.UserContent
	}

	if !dara.IsNil(request.UserInputs) {
		body["userInputs"] = request.UserInputs
	}

	if !dara.IsNil(request.VersionId) {
		body["versionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunAgent"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/bot/thread/run"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RunAgentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取生成式对话结果
//
// @param request - RunChatResultGenerationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunChatResultGenerationResponse
func (client *Client) RunChatResultGenerationWithSSECtx(ctx context.Context, workspaceId *string, request *RunChatResultGenerationRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunChatResultGenerationResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runChatResultGenerationWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, workspaceId, request, headers, runtime)
	return
}

// Summary:
//
// 获取生成式对话结果
//
// @param request - RunChatResultGenerationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunChatResultGenerationResponse
func (client *Client) RunChatResultGenerationWithContext(ctx context.Context, workspaceId *string, request *RunChatResultGenerationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunChatResultGenerationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InferenceParameters) {
		body["inferenceParameters"] = request.InferenceParameters
	}

	if !dara.IsNil(request.Messages) {
		body["messages"] = request.Messages
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	if !dara.IsNil(request.Tools) {
		body["tools"] = request.Tools
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunChatResultGeneration"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/run/chat/generation"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RunChatResultGenerationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 流式获取外呼会话分析结果
//
// @param request - RunDialogAnalysisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunDialogAnalysisResponse
func (client *Client) RunDialogAnalysisWithSSECtx(ctx context.Context, workspaceId *string, request *RunDialogAnalysisRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunDialogAnalysisResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runDialogAnalysisWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, workspaceId, request, headers, runtime)
	return
}

// Summary:
//
// 流式获取外呼会话分析结果
//
// @param request - RunDialogAnalysisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunDialogAnalysisResponse
func (client *Client) RunDialogAnalysisWithContext(ctx context.Context, workspaceId *string, request *RunDialogAnalysisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunDialogAnalysisResponse, _err error) {
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

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunDialogAnalysis"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/virtualHuman/dialog/stream/analysis"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RunDialogAnalysisResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取生成式对话结果
//
// @param request - RunLibraryChatGenerationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunLibraryChatGenerationResponse
func (client *Client) RunLibraryChatGenerationWithSSECtx(ctx context.Context, workspaceId *string, request *RunLibraryChatGenerationRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunLibraryChatGenerationResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runLibraryChatGenerationWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, workspaceId, request, headers, runtime)
	return
}

// Summary:
//
// 获取生成式对话结果
//
// @param request - RunLibraryChatGenerationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunLibraryChatGenerationResponse
func (client *Client) RunLibraryChatGenerationWithContext(ctx context.Context, workspaceId *string, request *RunLibraryChatGenerationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunLibraryChatGenerationResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocIdList) {
		body["docIdList"] = request.DocIdList
	}

	if !dara.IsNil(request.EnableFollowUp) {
		body["enableFollowUp"] = request.EnableFollowUp
	}

	if !dara.IsNil(request.EnableMultiQuery) {
		body["enableMultiQuery"] = request.EnableMultiQuery
	}

	if !dara.IsNil(request.EnableOpenQa) {
		body["enableOpenQa"] = request.EnableOpenQa
	}

	if !dara.IsNil(request.FollowUpLlm) {
		body["followUpLlm"] = request.FollowUpLlm
	}

	if !dara.IsNil(request.LibraryId) {
		body["libraryId"] = request.LibraryId
	}

	if !dara.IsNil(request.LlmType) {
		body["llmType"] = request.LlmType
	}

	if !dara.IsNil(request.MultiQueryLlm) {
		body["multiQueryLlm"] = request.MultiQueryLlm
	}

	if !dara.IsNil(request.Query) {
		body["query"] = request.Query
	}

	if !dara.IsNil(request.QueryCriteria) {
		body["queryCriteria"] = request.QueryCriteria
	}

	if !dara.IsNil(request.RerankType) {
		body["rerankType"] = request.RerankType
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	if !dara.IsNil(request.SubQueryList) {
		body["subQueryList"] = request.SubQueryList
	}

	if !dara.IsNil(request.TextSearchParameter) {
		body["textSearchParameter"] = request.TextSearchParameter
	}

	if !dara.IsNil(request.TopK) {
		body["topK"] = request.TopK
	}

	if !dara.IsNil(request.VectorSearchParameter) {
		body["vectorSearchParameter"] = request.VectorSearchParameter
	}

	if !dara.IsNil(request.WithDocumentReference) {
		body["withDocumentReference"] = request.WithDocumentReference
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunLibraryChatGeneration"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/run/library/chat/generation"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RunLibraryChatGenerationResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 提交问题列表
//
// @param request - SubmitChatQuestionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitChatQuestionResponse
func (client *Client) SubmitChatQuestionWithContext(ctx context.Context, workspaceId *string, request *SubmitChatQuestionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitChatQuestionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.GmtService) {
		body["gmtService"] = request.GmtService
	}

	if !dara.IsNil(request.LiveScriptContent) {
		body["liveScriptContent"] = request.LiveScriptContent
	}

	if !dara.IsNil(request.OpenSmallTalk) {
		body["openSmallTalk"] = request.OpenSmallTalk
	}

	if !dara.IsNil(request.QuestionList) {
		body["questionList"] = request.QuestionList
	}

	if !dara.IsNil(request.RequestId) {
		body["requestId"] = request.RequestId
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitChatQuestion"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/virtualHuman/chat/submit"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitChatQuestionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新文档
//
// @param request - UpdateDocumentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDocumentResponse
func (client *Client) UpdateDocumentWithContext(ctx context.Context, workspaceId *string, request *UpdateDocumentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDocumentResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocId) {
		body["docId"] = request.DocId
	}

	if !dara.IsNil(request.LibraryId) {
		body["libraryId"] = request.LibraryId
	}

	if !dara.IsNil(request.Meta) {
		body["meta"] = request.Meta
	}

	if !dara.IsNil(request.Title) {
		body["title"] = request.Title
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDocument"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/library/document/updateDocument"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDocumentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新文档的chunk
//
// @param request - UpdateDocumentChunkRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDocumentChunkResponse
func (client *Client) UpdateDocumentChunkWithContext(ctx context.Context, workspaceId *string, request *UpdateDocumentChunkRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateDocumentChunkResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Chunks) {
		body["chunks"] = request.Chunks
	}

	if !dara.IsNil(request.LibraryId) {
		body["libraryId"] = request.LibraryId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateDocumentChunk"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/library/updateDocumentChunk"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateDocumentChunkResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新文档库配置
//
// @param request - UpdateLibraryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLibraryResponse
func (client *Client) UpdateLibraryWithContext(ctx context.Context, workspaceId *string, request *UpdateLibraryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateLibraryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.IndexSetting) {
		body["indexSetting"] = request.IndexSetting
	}

	if !dara.IsNil(request.LibraryId) {
		body["libraryId"] = request.LibraryId
	}

	if !dara.IsNil(request.LibraryName) {
		body["libraryName"] = request.LibraryName
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateLibrary"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/library/update"),
		Method:      dara.String("PUT"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateLibraryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新QA问答库
//
// @param request - UpdateQaLibraryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateQaLibraryResponse
func (client *Client) UpdateQaLibraryWithContext(ctx context.Context, workspaceId *string, request *UpdateQaLibraryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateQaLibraryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ParseQaResults) {
		body["parseQaResults"] = request.ParseQaResults
	}

	if !dara.IsNil(request.QaLibraryId) {
		body["qaLibraryId"] = request.QaLibraryId
	}

	if !dara.IsNil(request.RequestId) {
		body["requestId"] = request.RequestId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateQaLibrary"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/virtualHuman/qa/upload"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateQaLibraryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上传文档到文档库
//
// @param request - UploadDocumentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadDocumentResponse
func (client *Client) UploadDocumentWithContext(ctx context.Context, workspaceId *string, request *UploadDocumentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UploadDocumentResponse, _err error) {
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

	if !dara.IsNil(request.FileName) {
		body["fileName"] = request.FileName
	}

	if !dara.IsNil(request.FileUrl) {
		body["fileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.LibraryId) {
		body["libraryId"] = request.LibraryId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UploadDocument"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/library/document/upload"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UploadDocumentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) realTimeDialogWithSSECtx_opYieldFunc(_yield chan *RealTimeDialogResponse, _yieldErr chan error, ctx context.Context, workspaceId *string, request *RealTimeDialogRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Analysis) {
		body["analysis"] = request.Analysis
	}

	if !dara.IsNil(request.BizType) {
		body["bizType"] = request.BizType
	}

	if !dara.IsNil(request.ConversationModel) {
		body["conversationModel"] = request.ConversationModel
	}

	if !dara.IsNil(request.DialogMemoryTurns) {
		body["dialogMemoryTurns"] = request.DialogMemoryTurns
	}

	if !dara.IsNil(request.MetaData) {
		body["metaData"] = request.MetaData
	}

	if !dara.IsNil(request.OpType) {
		body["opType"] = request.OpType
	}

	if !dara.IsNil(request.Recommend) {
		body["recommend"] = request.Recommend
	}

	if !dara.IsNil(request.ScriptContentPlayed) {
		body["scriptContentPlayed"] = request.ScriptContentPlayed
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	if !dara.IsNil(request.UserVad) {
		body["userVad"] = request.UserVad
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RealTimeDialog"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/realtime/dialog/chat"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runAgentWithSSECtx_opYieldFunc(_yield chan *RunAgentResponse, _yieldErr chan error, ctx context.Context, workspaceId *string, request *RunAgentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.BotId) {
		body["botId"] = request.BotId
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	if !dara.IsNil(request.ThreadId) {
		body["threadId"] = request.ThreadId
	}

	if !dara.IsNil(request.UseDraft) {
		body["useDraft"] = request.UseDraft
	}

	if !dara.IsNil(request.UserContent) {
		body["userContent"] = request.UserContent
	}

	if !dara.IsNil(request.UserInputs) {
		body["userInputs"] = request.UserInputs
	}

	if !dara.IsNil(request.VersionId) {
		body["versionId"] = request.VersionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunAgent"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/bot/thread/run"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runChatResultGenerationWithSSECtx_opYieldFunc(_yield chan *RunChatResultGenerationResponse, _yieldErr chan error, ctx context.Context, workspaceId *string, request *RunChatResultGenerationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InferenceParameters) {
		body["inferenceParameters"] = request.InferenceParameters
	}

	if !dara.IsNil(request.Messages) {
		body["messages"] = request.Messages
	}

	if !dara.IsNil(request.ModelId) {
		body["modelId"] = request.ModelId
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	if !dara.IsNil(request.Tools) {
		body["tools"] = request.Tools
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunChatResultGeneration"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/run/chat/generation"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runDialogAnalysisWithSSECtx_opYieldFunc(_yield chan *RunDialogAnalysisResponse, _yieldErr chan error, ctx context.Context, workspaceId *string, request *RunDialogAnalysisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunDialogAnalysis"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/virtualHuman/dialog/stream/analysis"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}

func (client *Client) runLibraryChatGenerationWithSSECtx_opYieldFunc(_yield chan *RunLibraryChatGenerationResponse, _yieldErr chan error, ctx context.Context, workspaceId *string, request *RunLibraryChatGenerationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocIdList) {
		body["docIdList"] = request.DocIdList
	}

	if !dara.IsNil(request.EnableFollowUp) {
		body["enableFollowUp"] = request.EnableFollowUp
	}

	if !dara.IsNil(request.EnableMultiQuery) {
		body["enableMultiQuery"] = request.EnableMultiQuery
	}

	if !dara.IsNil(request.EnableOpenQa) {
		body["enableOpenQa"] = request.EnableOpenQa
	}

	if !dara.IsNil(request.FollowUpLlm) {
		body["followUpLlm"] = request.FollowUpLlm
	}

	if !dara.IsNil(request.LibraryId) {
		body["libraryId"] = request.LibraryId
	}

	if !dara.IsNil(request.LlmType) {
		body["llmType"] = request.LlmType
	}

	if !dara.IsNil(request.MultiQueryLlm) {
		body["multiQueryLlm"] = request.MultiQueryLlm
	}

	if !dara.IsNil(request.Query) {
		body["query"] = request.Query
	}

	if !dara.IsNil(request.QueryCriteria) {
		body["queryCriteria"] = request.QueryCriteria
	}

	if !dara.IsNil(request.RerankType) {
		body["rerankType"] = request.RerankType
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	if !dara.IsNil(request.SubQueryList) {
		body["subQueryList"] = request.SubQueryList
	}

	if !dara.IsNil(request.TextSearchParameter) {
		body["textSearchParameter"] = request.TextSearchParameter
	}

	if !dara.IsNil(request.TopK) {
		body["topK"] = request.TopK
	}

	if !dara.IsNil(request.VectorSearchParameter) {
		body["vectorSearchParameter"] = request.VectorSearchParameter
	}

	if !dara.IsNil(request.WithDocumentReference) {
		body["withDocumentReference"] = request.WithDocumentReference
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunLibraryChatGeneration"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/run/library/chat/generation"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApiWithCtx(ctx, params, req, runtime, sseResp, _yieldErr)
	for resp := range sseResp {
		data := dara.ToMap(dara.ParseJSON(dara.StringValue(resp.Event.Data)))
		_err := dara.ConvertChan(map[string]interface{}{
			"statusCode": dara.IntValue(resp.StatusCode),
			"headers":    resp.Headers,
			"body": dara.ToMap(map[string]interface{}{
				"RequestId": dara.StringValue(resp.Event.Id),
				"Message":   dara.StringValue(resp.Event.Event),
			}, data),
		}, _yield)
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
}
