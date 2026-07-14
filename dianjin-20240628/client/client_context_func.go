// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"context"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	websocketutils "github.com/alibabacloud-go/darabonba-openapi/v2/websocketUtils"
	"github.com/alibabacloud-go/tea/dara"
)

// Summary:
//
// 申请取数
//
// @param request - CommercializeFetchRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CommercializeFetchResponse
func (client *Client) CommercializeFetchWithContext(ctx context.Context, workspaceId *string, cjfCode *string, zjfCode *string, request *CommercializeFetchRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CommercializeFetchResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ChannelId) {
		body["channelId"] = request.ChannelId
	}

	if !dara.IsNil(request.Data) {
		body["data"] = request.Data
	}

	if !dara.IsNil(request.EncryptType) {
		body["encryptType"] = request.EncryptType
	}

	if !dara.IsNil(request.Env) {
		body["env"] = request.Env
	}

	if !dara.IsNil(request.ProductId) {
		body["productId"] = request.ProductId
	}

	if !dara.IsNil(request.RequestId) {
		body["requestId"] = request.RequestId
	}

	if !dara.IsNil(request.SecretKey) {
		body["secretKey"] = request.SecretKey
	}

	if !dara.IsNil(request.Sign) {
		body["sign"] = request.Sign
	}

	if !dara.IsNil(request.SignType) {
		body["signType"] = request.SignType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CommercializeFetch"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/spi/path/" + dara.PercentEncode(dara.StringValue(cjfCode)) + "/api/support/" + dara.PercentEncode(dara.StringValue(zjfCode)) + "/firefly/commercializeFetch"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CommercializeFetchResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create a task to summarize documents by year.
//
// Description:
//
// Before you use this operation, review the billing methods and pricing for Alibaba Cloud Tongyi Dianjin.
//
// # Prerequisites
//
// You have activated Alibaba Cloud Model Studio and Tongyi Dianjin.
//
// Obtain your [workspace ID](https://help.aliyun.com/zh/model-studio/developer-reference/get-app-id-and-workspace?spm=openapi-amp.newDocPublishment.0.0.2eb8281foUVd15#2612f896detsz:~:text=%E6%9F%A5%E7%9C%8BAPI%2DKEY%E3%80%82-,%E8%8E%B7%E5%8F%96APP%2DID%E5%92%8CWORKSPACE,-%E8%BF%9B%E5%85%A5%E6%88%91%E7%9A%84%E5%BA%94%E7%94%A8).
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
// Creates an outbound call session.
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
// Creates a session analysis task. After the task is created, use the session ID with GetDialogAnalysisResult to retrieve the results.
//
// Description:
//
// Before you use this API, review the billing methods and pricing for DianJin.
//
// # Prerequisites
//
// You have activated Alibaba Cloud Model Studio and DianJin.
//
// Obtain a workspace ID. For more information, see [Get a workspace ID](https://help.aliyun.com/zh/model-studio/developer-reference/get-app-id-and-workspace?spm=openapi-amp.newDocPublishment.0.0.2eb8281foUVd15#2612f896detsz:~:text=%E6%9F%A5%E7%9C%8BAPI%2DKEY%E3%80%82-,%E8%8E%B7%E5%8F%96APP%2DID%E5%92%8CWORKSPACE,-%E8%BF%9B%E5%85%A5%E6%88%91%E7%9A%84%E5%BA%94%E7%94%A8).
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
// Creates a multi-document summary task.
//
// Description:
//
// Before you use this API, review the billing methods and pricing for Alibaba Cloud Tongyi Dianjin.
//
// # Prerequisites
//
// You have activated Alibaba Cloud Model Studio and Tongyi Dianjin.
//
// You have obtained a [workspace identifier](https://help.aliyun.com/zh/model-studio/developer-reference/get-app-id-and-workspace?spm=openapi-amp.newDocPublishment.0.0.2eb8281foUVd15#2612f896detsz:~:text=%E6%9F%A5%E7%9C%8BAPI%2DKEY%E3%80%82-,%E8%8E%B7%E5%8F%96APP%2DID%E5%92%8CWORKSPACE,-%E8%BF%9B%E5%85%A5%E6%88%91%E7%9A%84%E5%BA%94%E7%94%A8).
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
// Creates a financial report summary.
//
// Description:
//
// Before using this API, review the pricing and billing methods for Alibaba Cloud Gold products.
//
// **Prerequisites**
//
// - Enable Alibaba Cloud Model Studio and Alibaba Cloud Gold services.
//
// - Obtain the workspace ID. For more information, see [workspace identity](https://help.aliyun.com/zh/model-studio/developer-reference/get-app-id-and-workspace?spm=openapi-amp.newDocPublishment.0.0.2eb8281foUVd15#2612f896detsz:~:text=%E6%9F%A5%E7%9C%8BAPI%2DKEY%E3%80%82-,%E8%8E%B7%E5%8F%96APP%2DID%E5%92%8CWORKSPACE,-%E8%BF%9B%E5%85%A5%E6%88%91%E7%9A%84%E5%BA%94%E7%94%A8).
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
// Creates a document library. A document library isolates document and index data. If your use case requires frequent natural language search by category, create multiple libraries to isolate different data types. You can customize vector and text indexes by format.
//
// Description:
//
// *Prerequisites**
//
// - You have activated Alibaba Cloud Model Studio and Tongyi Dianjin.
//
// - Obtain the workspace ID: Retrieve the [workspace identifier](https://help.aliyun.com/zh/model-studio/developer-reference/get-app-id-and-workspace?spm=openapi-amp.newDocPublishment.0.0.2eb8281foUVd15#2612f896detsz:~:text=%E6%9F%A5%E7%9C%8BAPI%2DKEY%E3%80%82-,%E8%8E%B7%E5%8F%96APP%2DID%E5%92%8CWORKSPACE,-%E8%BF%9B%E5%85%A5%E6%88%91%E7%9A%84%E5%BA%94%E7%94%A8).
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
// Create a PDF document translation task. Submit the task to start asynchronous translation.
//
// Description:
//
// Before you use this operation, review the billing methods and pricing for Alibaba Cloud Tongyi Dianjin.
//
// **Prerequisites**
//
// - You have activated Alibaba Cloud Model Studio and Tongyi Dianjin.
//
// - You have obtained a workspace ID. To obtain your [workspace ID](https://help.aliyun.com/zh/model-studio/developer-reference/get-app-id-and-workspace?spm=openapi-amp.newDocPublishment.0.0.2eb8281foUVd15#2612f896detsz:~:text=%E6%9F%A5%E7%9C%8BAPI%2DKEY%E3%80%82-,%E8%8E%B7%E5%8F%96APP%2DID%E5%92%8CWORKSPACE,-%E8%BF%9B%E5%85%A5%E6%88%91%E7%9A%84%E5%BA%94%E7%94%A8).
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
// Creates document chunks based on your business scenarios.
//
// Description:
//
// Before using this API, review the billing methods and pricing for Tongyi Dianjin.
//
// **Prerequisites**
//
// - Activate Alibaba Cloud Model Studio and Tongyi Dianjin.
//
// - Obtain a workspace ID. For more information, see [Get a workspace identity](https://help.aliyun.com/zh/model-studio/developer-reference/get-app-id-and-workspace?spm=openapi-amp.newDocPublishment.0.0.2eb8281foUVd15#2612f896detsz:~:text=%E6%9F%A5%E7%9C%8BAPI%2DKEY%E3%80%82-,%E8%8E%B7%E5%8F%96APP%2DID%E5%92%8CWORKSPACE,-%E8%BF%9B%E5%85%A5%E6%88%91%E7%9A%84%E5%BA%94%E7%94%A8).
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
// Creates a quality check task.
//
// Description:
//
// Before using this API, review the pricing and billing methods for the Tongyi Dianjin product.
//
// # Prerequisites
//
// Activate Alibaba Cloud Model Studio and Tongyi Dianjin services.
//
// Obtain the workspaceId: Retrieve the [workspace identity](https://help.aliyun.com/zh/model-studio/developer-reference/get-app-id-and-workspace?spm=openapi-amp.newDocPublishment.0.0.2eb8281foUVd15#2612f896detsz:~:text=%E6%9F%A5%E7%9C%8BAPI%2DKEY%E3%80%82-,%E8%8E%B7%E5%8F%96APP%2DID%E5%92%8CWORKSPACE,-%E8%BF%9B%E5%85%A5%E6%88%91%E7%9A%84%E5%BA%94%E7%94%A8).
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
// A callback event that indicates the completion of a Dashscope asynchronous task.
//
// @param request - DashscopeAsyncTaskFinishEventRequest
//
// @param headers - DashscopeAsyncTaskFinishEventHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DashscopeAsyncTaskFinishEventResponse
func (client *Client) DashscopeAsyncTaskFinishEventWithContext(ctx context.Context, workspaceId *string, request *DashscopeAsyncTaskFinishEventRequest, headers *DashscopeAsyncTaskFinishEventHeaders, runtime *dara.RuntimeOptions) (_result *DashscopeAsyncTaskFinishEventResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Body) {
		body["body"] = request.Body
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
		Action:      dara.String("DashscopeAsyncTaskFinishEvent"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/event/dashscopeAsyncTaskFinish"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DashscopeAsyncTaskFinishEventResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a document. After deletion, you cannot view the original document or recall it.
//
// Description:
//
// *Prerequisites**
//
// - You have activated Alibaba Cloud Model Studio and Tongyi Dianjin services.
//
// - Obtain your workspace ID: retrieve your [workspace identifier](https://help.aliyun.com/zh/model-studio/developer-reference/get-app-id-and-workspace?spm=openapi-amp.newDocPublishment.0.0.2eb8281foUVd15#2612f896detsz:~:text=%E6%9F%A5%E7%9C%8BAPI%2DKEY%E3%80%82-,%E8%8E%B7%E5%8F%96APP%2DID%E5%92%8CWORKSPACE,-%E8%BF%9B%E5%85%A5%E6%88%91%E7%9A%84%E5%BA%94%E7%94%A8).
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
// Delete a document library. ⚠️ This operation deletes the library and all its associated documents.
//
// Description:
//
// *Prerequisites**
//
// - Activate Alibaba Cloud Model Studio and Tongyi Dianjin services.
//
// - Obtain your workspaceId. For more information, refer to the [workspace identifier](https://help.aliyun.com/zh/model-studio/developer-reference/get-app-id-and-workspace?spm=openapi-amp.newDocPublishment.0.0.2eb8281foUVd15#2612f896detsz:~:text=%E6%9F%A5%E7%9C%8BAPI%2DKEY%E3%80%82-,%E8%8E%B7%E5%8F%96APP%2DID%E5%92%8CWORKSPACE,-%E8%BF%9B%E5%85%A5%E6%88%91%E7%9A%84%E5%BA%94%E7%94%A8).
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
// This API uses the WebSocket protocol to perform real-time conversational transcription, intent recognition, and speech synthesis. It supports various audio formats for both input and output to ensure real-time performance and high compatibility.
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
		Action:               dara.String("EndToEndRealTimeDialog"),
		Version:              dara.String("2024-06-28"),
		Protocol:             dara.String("wss"),
		Pathname:             dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/ws/realtime/dialog"),
		Method:               dara.String("GET"),
		AuthType:             dara.String("AK"),
		Style:                dara.String("ROA"),
		ReqBodyType:          dara.String("json"),
		BodyType:             dara.String("json"),
		WebsocketSubProtocol: dara.String("awap"),
	}
	res := &EndToEndRealTimeDialogResponse{}
	callApiTmp, err := client.CallApiWithCtx(ctx, params, req, runtime)
	if err != nil {
		_err = err
		return _result, _err
	}
	tmp := dara.ToMap(callApiTmp)
	if !dara.IsNil(tmp["webSocketClient"]) {
		res.WebSocketClient = websocketutils.CreateWebSocketClient(tmp["webSocketClient"])
	}

	_result = res
	return _result, _err
}

// Summary:
//
// Terminate the job.
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
// 兑换权益
//
// @param request - ExchangeEntitlementRequest
//
// @param headers - ExchangeEntitlementHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExchangeEntitlementResponse
func (client *Client) ExchangeEntitlementWithContext(ctx context.Context, workspaceId *string, tenantId *string, request *ExchangeEntitlementRequest, headers *ExchangeEntitlementHeaders, runtime *dara.RuntimeOptions) (_result *ExchangeEntitlementResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ExternalUserId) {
		body["externalUserId"] = request.ExternalUserId
	}

	if !dara.IsNil(request.KeyHash) {
		body["keyHash"] = request.KeyHash
	}

	if !dara.IsNil(request.RequestId) {
		body["requestId"] = request.RequestId
	}

	if !dara.IsNil(request.TemplateId) {
		body["templateId"] = request.TemplateId
	}

	if !dara.IsNil(request.UserName) {
		body["userName"] = request.UserName
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
		Action:      dara.String("ExchangeEntitlement"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(tenantId)) + "/redeem"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ExchangeEntitlementResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Parses question and answer (Q&A) pairs from a document. You can use the UpdateQaLibrary API to update the Q&A pairs.
//
// Description:
//
// Before you use this API, make sure you understand the billing methods and pricing of the Tongyi Dianjin product.
//
// # Prerequisites
//
// Activate Alibaba Cloud Model Studio and the Tongyi Dianjin service.
//
// Obtain a workspace ID. For more information, see [Get a workspace ID](https://help.aliyun.com/zh/model-studio/developer-reference/get-app-id-and-workspace?spm=openapi-amp.newDocPublishment.0.0.2eb8281foUVd15#2612f896detsz:~:text=%E6%9F%A5%E7%9C%8BAPI%2DKEY%E3%80%82-,%E8%8E%B7%E5%8F%96APP%2DID%E5%92%8CWORKSPACE,-%E8%BF%9B%E5%85%A5%E6%88%91%E7%9A%84%E5%BA%94%E7%94%A8).
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
// Retrieve app configuration.
//
// Description:
//
// *Prerequisites**
//
// - You have activated Alibaba Cloud Model Studio and Tongyi Gold services.
//
// - You can obtain the workspace ID. For details, see [workspace identity](https://help.aliyun.com/zh/model-studio/developer-reference/get-app-id-and-workspace?spm=openapi-amp.newDocPublishment.0.0.2eb8281foUVd15#2612f896detsz:~:text=%E6%9F%A5%E7%9C%8BAPI%2DKEY%E3%80%82-,%E8%8E%B7%E5%8F%96APP%E3%80%9DID%E5%92%8CWORKSPACE,-%E8%BF%9B%E5%85%A5%E6%88%91%E7%9A%84%E5%BA%94%E7%94%A8).
//
// @param request - GetAppConfigRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAppConfigResponse
func (client *Client) GetAppConfigWithContext(ctx context.Context, workspaceId *string, request *GetAppConfigRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAppConfigResponse, _err error) {
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
// Retrieves the Q&A results generated by the SubmitChatQuestion API.
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
// Retrieve session analysis results. You can retrieve results in batches by specifying a list of session IDs or a time range.
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
// Retrieve session details.
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
// Retrieves records of real-time conversations and the results of intent analysis.
//
// Description:
//
// ## Request description
//
// This API retrieves conversation records between customers and service agents, along with intent analysis results generated by the model.
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
// Retrieve a list of document chunks. You can filter them by query conditions.
//
// Description:
//
// *Prerequisites**
//
// - Activate Alibaba Cloud Model Studio and Tongyi Gold Service.
//
// - Obtain the workspaceId and the [workspace identity](https://help.aliyun.com/zh/model-studio/developer-reference/get-app-id-and-workspace?spm=openapi-amp.newDocPublishment.0.0.2eb8281foUVd15#2612f896detsz:~:text=%E6%9F%A5%E7%9C%8BAPI%2DKEY%E3%80%82-,%E8%8E%B7%E5%8F%96APP%2DID%E5%92%8CWORKSPACE,-%E8%BF%9B%E5%85%A5%E6%88%91%E7%9A%84%E5%BA%94%E7%94%B3).
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
// Retrieves a list of documents from a document library. This operation supports paged queries and filtering by document status.
//
// Description:
//
// *Prerequisites**
//
// - You have activated Alibaba Cloud Model Studio and the Tongyi Gold Point service.
//
// - You have obtained a workspace ID. For more information, see [Get a workspace ID](https://help.aliyun.com/zh/model-studio/developer-reference/get-app-id-and-workspace?spm=openapi-amp.newDocPublishment.0.0.2eb8281foUVd15#2612f896detsz:~:text=%E6%9F%A5%E7%9C%8BAPI%2DKEY%E3%80%82-,%E8%8E%B7%E5%8F%96APP%2DID%E5%92%8CWORKSPACE,-%E8%BF%9B%E5%85%A5%E6%88%91%E7%9A%84%E5%BA%94%E7%94%A8).
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
// Retrieve a download URL for a document. The URL expires after 1 hour.
//
// Description:
//
// *Prerequisites**
//
// - You have activated Alibaba Cloud Model Studio and Tongyi Dianjin.
//
// - Obtain the workspace ID. For more information, see [Get the workspace ID](https://help.aliyun.com/zh/model-studio/developer-reference/get-app-id-and-workspace?spm=openapi-amp.newDocPublishment.0.0.2eb8281foUVd15#2612f896detsz:~:text=%E6%9F%A5%E7%9C%8BAPI%2DKEY%E3%80%82-,%E8%8E%B7%E5%8F%96APP%2DID%E5%92%8CWORKSPACE,-%E8%BF%9B%E5%85%A5%E6%88%91%E7%9A%84%E5%BA%94%E7%94%A8).
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
// This operation retrieves a list of documents. You can filter documents by metadata or use paging.
//
// Description:
//
// # Prerequisites
//
// You have activated Alibaba Cloud Model Studio and Tongyi Dianjin.
//
// To obtain the workspace ID, see [the document about obtaining the workspace ID](https://help.aliyun.com/zh/model-studio/developer-reference/get-app-id-and-workspace?spm=openapi-amp.newDocPublishment.0.0.2eb8281foUVd15#2612f896detsz:~:text=%E6%9F%A5%E7%9C%8BAPI%2DKEY%E3%80%82-,%E8%8E%B7%E5%8F%96APP%2DID%E5%92%8CWORKSPACE,-%E8%BF%9B%E5%85%A5%E6%88%91%E7%9A%84%E5%BA%94%E7%94%A8).
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
// Retrieve conversation history records by business type.
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
// Retrieves the detailed configuration of a document library, including its name, description, and index settings.
//
// Description:
//
// *Prerequisites**
//
// - Activate the Alibaba Cloud Model Studio and Tongyi Dianjin services.
//
// - Obtain a [workspace identifier](https://help.aliyun.com/zh/model-studio/developer-reference/get-app-id-and-workspace?spm=openapi-amp.newDocPublishment.0.0.2eb8281foUVd15#2612f896detsz:~:text=%E6%9F%A5%E7%9C%8BAPI%2DKEY%E3%80%82-,%E8%8E%B7%E5%8F%96APP%2DID%E5%92%8CWORKSPACE,-%E8%BF%9B%E5%85%A5%E6%88%91%E7%9A%84%E5%BA%94%E7%94%A8).
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
// Retrieve the document library list. The list includes document names, descriptions, and unique identifiers.
//
// Description:
//
// # Prerequisites
//
// - You must activate Alibaba Cloud Model Studio and Tongyi Dianjin services.
//
// - Obtain the workspaceId. For more information, see [workspace identity](https://help.aliyun.com/zh/model-studio/developer-reference/get-app-id-and-workspace?spm=openapi-amp.newDocPublishment.0.0.2eb8281foUVd15#2612f896detsz:~:text=%E6%9F%A5%E7%9C%8BAPI%2DKEY%E3%80%82-,%E8%8E%B7%E5%8F%96APP%2DID%E5%92%8CWORKSPACE,-%E8%BF%9B%E5%85%A5%E6%88%91%E7%9A%84%E5%BA%94%E7%94%A8).
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
// Retrieve document parsing results. You can query the document\\"s parsing status and obtain the parsing results.
//
// Description:
//
// *Prerequisites**
//
// - You have activated Alibaba Cloud Model Studio and Tongyi Gold services.
//
// - Obtain the workspace ID. For more information, see [the workspace identity document](https://help.aliyun.com/zh/model-studio/developer-reference/get-app-id-and-workspace?spm=openapi-amp.newDocPublishment.0.0.2eb8281foUVd15#2612f896detsz:~:text=%E6%9F%A5%E7%9C%8BAPI%2DKEY%E3%80%82-,%E8%8E%B7%E5%8F%96APP%2DID%E5%92%8CWORKSPACE,-%E8%BF%9B%E5%85%A5%E6%88%91%E7%9A%84%E5%BA%94%E7%94%A8).
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
// Retrieve quality check results.
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
// 获取报告结果
//
// @param request - GetReportResponseRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetReportResponseResponse
func (client *Client) GetReportResponseWithContext(ctx context.Context, workspaceId *string, sceneCode *string, fundProduct *string, outRequestNo *string, request *GetReportResponseRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetReportResponseResponse, _err error) {
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
		Action:      dara.String("GetReportResponse"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/firefly/v1/" + dara.PercentEncode(dara.StringValue(sceneCode)) + "/" + dara.PercentEncode(dara.StringValue(fundProduct)) + "/tasks/" + dara.PercentEncode(dara.StringValue(outRequestNo)) + "/report"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetReportResponseResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取任务状态
//
// @param request - GetReportTaskStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetReportTaskStatusResponse
func (client *Client) GetReportTaskStatusWithContext(ctx context.Context, workspaceId *string, sceneCode *string, fundProduct *string, outRequestNo *string, request *GetReportTaskStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetReportTaskStatusResponse, _err error) {
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
		Action:      dara.String("GetReportTaskStatus"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/firefly/v1/" + dara.PercentEncode(dara.StringValue(sceneCode)) + "/" + dara.PercentEncode(dara.StringValue(fundProduct)) + "/tasks/" + dara.PercentEncode(dara.StringValue(outRequestNo))),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetReportTaskStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Retrieves the result of a financial report summary task.
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
// Retrieves the result of an asynchronous task.
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
// Gets the status of a task.
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
// 查询用量明细
//
// @param request - GetUsageRequest
//
// @param headers - GetUsageHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUsageResponse
func (client *Client) GetUsageWithContext(ctx context.Context, workspaceId *string, tenantId *string, request *GetUsageRequest, headers *GetUsageHeaders, runtime *dara.RuntimeOptions) (_result *GetUsageResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExternalUserId) {
		query["externalUserId"] = request.ExternalUserId
	}

	if !dara.IsNil(request.RedemptionOrderNo) {
		query["redemptionOrderNo"] = request.RedemptionOrderNo
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
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetUsage"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(tenantId)) + "/usage"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetUsageResponse{}
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
// Call a plugin and retrieve its response.
//
// Description:
//
// *Prerequisites**
//
// - You have activated Alibaba Cloud Model Studio and Tongyi Dianjin services.
//
// - Obtain the workspace ID. For more information, see [Get the workspace ID](https://help.aliyun.com/zh/model-studio/developer-reference/get-app-id-and-workspace?spm=openapi-amp.newDocPublishment.0.0.2eb8281foUVd15#2612f896detsz:~:text=%E6%9F%A5%E7%9C%8BAPI%2DKEY%E3%80%82-,%E8%8E%B7%E5%8F%96APP%2DID%E5%92%8CWORKSPACE,-%E8%BF%9B%E5%85%A5%E6%88%91%E7%9A%84%E5%BA%94%E7%94%A8).
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
// Preview documents. Retrieve document download links, types, and titles. Use this operation to preview documents.
//
// Description:
//
// *Prerequisites**
//
// - You must activate Alibaba Cloud Model Studio and Tongyi Gold services.
//
// - Obtain the workspace ID. For more information, see [workspace identity](https://help.aliyun.com/zh/model-studio/developer-reference/get-app-id-and-workspace?spm=openapi-amp.newDocPublishment.0.0.2eb8281foUVd15#2612f896detsz:~:text=%E6%9F%A5%E7%9C%8BAPI%2DKEY%E3%80%82-,%E8%8E%B7%E5%8F%96APP%2DID%E5%92%8CWORKSPACE,-%E8%BF%9B%E5%85%A5%E6%88%91%E7%9A%84%E5%BA%94%E7%94%A8).
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
// 查询兑换记录
//
// @param request - QueryApiKeysRequest
//
// @param headers - QueryApiKeysHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryApiKeysResponse
func (client *Client) QueryApiKeysWithContext(ctx context.Context, workspaceId *string, tenantId *string, request *QueryApiKeysRequest, headers *QueryApiKeysHeaders, runtime *dara.RuntimeOptions) (_result *QueryApiKeysResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExternalUserId) {
		query["externalUserId"] = request.ExternalUserId
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
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryApiKeys"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(tenantId)) + "/apikeys"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryApiKeysResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询兑换记录
//
// @param request - QueryRedemptionRecordsRequest
//
// @param headers - QueryRedemptionRecordsHeaders
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryRedemptionRecordsResponse
func (client *Client) QueryRedemptionRecordsWithContext(ctx context.Context, workspaceId *string, tenantId *string, request *QueryRedemptionRecordsRequest, headers *QueryRedemptionRecordsHeaders, runtime *dara.RuntimeOptions) (_result *QueryRedemptionRecordsResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExternalUserId) {
		query["externalUserId"] = request.ExternalUserId
	}

	if !dara.IsNil(request.Page) {
		query["page"] = request.Page
	}

	if !dara.IsNil(request.PageSize) {
		query["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.RedemptionOrderNo) {
		query["redemptionOrderNo"] = request.RedemptionOrderNo
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
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("QueryRedemptionRecords"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/v1/tenants/" + dara.PercentEncode(dara.StringValue(tenantId)) + "/redemption-records"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &QueryRedemptionRecordsResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Reindexing reprocesses the specified document by parsing it, splitting it into chunks, and building a new index.
//
// Description:
//
// Before you use this operation, review the billing method and pricing for Tongyi Dianjin.
//
// **Prerequisites**
//
// - You have activated Alibaba Cloud Model Studio and Tongyi Dianjin.
//
// - You have obtained a workspace ID: Retrieve the [workspace identifier](https://help.aliyun.com/zh/model-studio/developer-reference/get-app-id-and-workspace?spm=openapi-amp.newDocPublishment.0.0.2eb8281foUVd15#2612f896detsz:~:text=%E6%9F%A5%E7%9C%8BAPI%2DKEY%E3%80%82-,%E8%8E%B7%E5%8F%96APP%2DID%E5%92%8CWORKSPACE,-%E8%BF%9B%E5%85%A5%E6%88%91%E7%9A%84%E5%BA%94%E7%94%A8).
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
// Starts a real-time session. After you create a session by calling the CreateDialog API, use this API to conduct the real-time interaction.
//
// Description:
//
// Before using this API, make sure you understand the billing methods and pricing of the Tongyi Gold service.
//
// # Prerequisites
//
// Alibaba Cloud Model Studio and the Tongyi Gold service are activated.
//
// Obtain a workspaceId. For more information, see [Get an app ID and workspace](https://help.aliyun.com/zh/model-studio/developer-reference/get-app-id-and-workspace?spm=openapi-amp.newDocPublishment.0.0.2eb8281foUVd15#2612f896detsz:~:text=%E6%9F%A5%E7%9C%8BAPI%2DKEY%E3%80%82-,%E8%8E%B7%E5%8F%96APP%2DID%E5%92%8CWORKSPACE,-%E8%BF%9B%E5%85%A5%E6%88%91%E7%9A%84%E5%BA%94%E7%94%A8).
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
// Starts a real-time session. After you create a session by calling the CreateDialog API, use this API to conduct the real-time interaction.
//
// Description:
//
// Before using this API, make sure you understand the billing methods and pricing of the Tongyi Gold service.
//
// # Prerequisites
//
// Alibaba Cloud Model Studio and the Tongyi Gold service are activated.
//
// Obtain a workspaceId. For more information, see [Get an app ID and workspace](https://help.aliyun.com/zh/model-studio/developer-reference/get-app-id-and-workspace?spm=openapi-amp.newDocPublishment.0.0.2eb8281foUVd15#2612f896detsz:~:text=%E6%9F%A5%E7%9C%8BAPI%2DKEY%E3%80%82-,%E8%8E%B7%E5%8F%96APP%2DID%E5%92%8CWORKSPACE,-%E8%BF%9B%E5%85%A5%E6%88%91%E7%9A%84%E5%BA%94%E7%94%A8).
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
// Provides real-time dialog assistance after a session is created using CreateDialog. Note: This operation returns multiple intents, tags, and SOP flows. Unlike real-time sessions, it does not support streaming responses.
//
// Description:
//
// Before you use this API, make sure that you understand the billing methods and [pricing](https://help.aliyun.com/zh/model-studio/tongyi-dianjin-overview?spm=a2c4g.11186623.help-menu-2400256.d_1_6_6_0.15e77499sSMTGb) of Alibaba Cloud Model Studio DianJin.
//
// # Prerequisites
//
// Activate the Alibaba Cloud Model Studio and Model Studio DianJin services.
//
// Obtain the [workspace identity](https://help.aliyun.com/zh/model-studio/developer-reference/get-app-id-and-workspace?spm=openapi-amp.newDocPublishment.0.0.2eb8281foUVd15#2612f896detsz:~:text=%E6%9F%A5%E7%9C%8BAPI%2DKEY%E3%80%82-,%E8%8E%B7%E5%8F%96APP%2DID%E5%92%8CWORKSPACE,-%E8%BF%9B%E5%85%A5%E6%88%91%E7%9A%84%E5%BA%94%E7%94%A8) to use as your workspaceId.
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

	if !dara.IsNil(request.ScriptContentPlayed) {
		body["scriptContentPlayed"] = request.ScriptContentPlayed
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	if !dara.IsNil(request.UserVad) {
		body["userVad"] = request.UserVad
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
// Rebuilds an existing task. You cannot rebuild tasks that are queued or currently executing.
//
// Description:
//
// Before you use this API, review the billing methods and pricing for the Dianjin service.
//
// # Prerequisites
//
// You have activated Alibaba Cloud Model Studio and the Dianjin service.
//
// You have obtained a workspace ID. For more information, see [Get an app ID and workspace](https://help.aliyun.com/zh/model-studio/developer-reference/get-app-id-and-workspace?spm=openapi-amp.newDocPublishment.0.0.2eb8281foUVd15#2612f896detsz:~:text=%E6%9F%A5%E7%9C%8BAPI%2DKEY%E3%80%82-,%E8%8E%B7%E5%8F%96APP%2DID%E5%92%8CWORKSPACE,-%E8%BF%9B%E5%85%A5%E6%88%91%E7%9A%84%E5%BA%94%E7%94%A8).
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
// You can retrieve document chunks from a document library using text. You can specify the number of chunks to retrieve, filter them by metadata conditions, and choose whether to complete the document chunks.
//
// Description:
//
// *Prerequisites**
//
// - Activate the Alibaba Cloud Model Studio service and the Tongyi Dianjin service.
//
// - Obtain the workspace ID. Obtain the [workspace identity](https://help.aliyun.com/zh/model-studio/developer-reference/get-app-id-and-workspace?spm=openapi-amp.newDocPublishment.0.0.2eb8281foUVd15#2612f896detsz:~:text=%E6%9F%93%E7%9C%8BAPI%2DKEY%E3%80%82-,%E8%8E%B7%E5%8F%96APP%2DID%E5%92%8CWORKSPACE,-%E8%BF%9B%E5%85%A5%E6%88%91%E7%9A%84%E5%BA%94%E7%94%A8).
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
// Detects intents. This operation supports global and hierarchical intent detection, attitude detection, and enterprise detection.
//
// Description:
//
// Before you use this operation, review the billing methods and pricing of Alibaba Cloud Tongyi Dianjin.
//
// **Prerequisites**
//
// - You have activated Alibaba Cloud Model Studio and Tongyi Dianjin.
//
// - You have obtained a [workspace ID](https://help.aliyun.com/zh/model-studio/developer-reference/get-app-id-and-workspace?spm=openapi-amp.newDocPublishment.0.0.2eb8281foUVd15#2612f896detsz:~:text=%E6%9F%A5%E7%9C%8BAPI%2DKEY%E3%80%82-,%E8%8E%B7%E5%8F%96APP%2DID%E5%92%8CWORKSPACE,-%E8%BF%9B%E5%85%A5%E6%88%91%E7%9A%84%E5%BA%94%E7%94%A8).
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
// 重试任务
//
// @param request - RetryReportTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetryReportTaskResponse
func (client *Client) RetryReportTaskWithContext(ctx context.Context, workspaceId *string, sceneCode *string, fundProduct *string, outRequestNo *string, request *RetryReportTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RetryReportTaskResponse, _err error) {
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
		Action:      dara.String("RetryReportTask"),
		Version:     dara.String("2024-06-28"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/api/firefly/v1/" + dara.PercentEncode(dara.StringValue(sceneCode)) + "/" + dara.PercentEncode(dara.StringValue(fundProduct)) + "/tasks/" + dara.PercentEncode(dara.StringValue(outRequestNo)) + "/retry"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RetryReportTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Run an agent. This API supports both streaming and non-streaming responses.
//
// Description:
//
// Before you use this API, review the billing model and pricing for Tongyi Dianjin.
//
// # Prerequisites
//
// You have activated Alibaba Cloud Model Studio and Tongyi Dianjin.
//
// Obtain a workspace ID: [Obtain your workspace ID](https://help.aliyun.com/zh/model-studio/developer-reference/get-app-id-and-workspace?spm=openapi-amp.newDocPublishment.0.0.2eb8281foUVd15#2612f896detsz:~:text=%E6%9F%A5%E7%9C%8BAPI%2DKEY%E3%80%82-,%E8%8E%B7%E5%8F%96APP%2DID%E5%92%8CWORKSPACE,-%E8%BF%9B%E5%85%A5%E6%88%91%E7%9A%84%E5%BA%94%E7%94%A8).
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
// Run an agent. This API supports both streaming and non-streaming responses.
//
// Description:
//
// Before you use this API, review the billing model and pricing for Tongyi Dianjin.
//
// # Prerequisites
//
// You have activated Alibaba Cloud Model Studio and Tongyi Dianjin.
//
// Obtain a workspace ID: [Obtain your workspace ID](https://help.aliyun.com/zh/model-studio/developer-reference/get-app-id-and-workspace?spm=openapi-amp.newDocPublishment.0.0.2eb8281foUVd15#2612f896detsz:~:text=%E6%9F%A5%E7%9C%8BAPI%2DKEY%E3%80%82-,%E8%8E%B7%E5%8F%96APP%2DID%E5%92%8CWORKSPACE,-%E8%BF%9B%E5%85%A5%E6%88%91%E7%9A%84%E5%BA%94%E7%94%A8).
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
// Generates chat responses. You can select a model for the conversation and choose streaming or non-streaming output.
//
// Description:
//
// Before you use this API, review the billing method and pricing for Alibaba Cloud Tongyi Dianjin.
//
// **Prerequisites**
//
// - Activate Alibaba Cloud Model Studio and Tongyi Dianjin.
//
// - Obtain your workspace ID: retrieve your [workspace identifier](https://help.aliyun.com/zh/model-studio/developer-reference/get-app-id-and-workspace?spm=openapi-amp.newDocPublishment.0.0.2eb8281foUVd15#2612f896detsz:~:text=%E6%9F%A5%E7%9C%8BAPI%2DKEY%E3%80%82-,%E8%8E%B7%E5%8F%96APP%2DID%E5%92%8CWORKSPACE,-%E8%BF%9B%E5%85%A5%E6%88%91%E7%9A%84%E5%BA%94%E7%94%A8).
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
// Generates chat responses. You can select a model for the conversation and choose streaming or non-streaming output.
//
// Description:
//
// Before you use this API, review the billing method and pricing for Alibaba Cloud Tongyi Dianjin.
//
// **Prerequisites**
//
// - Activate Alibaba Cloud Model Studio and Tongyi Dianjin.
//
// - Obtain your workspace ID: retrieve your [workspace identifier](https://help.aliyun.com/zh/model-studio/developer-reference/get-app-id-and-workspace?spm=openapi-amp.newDocPublishment.0.0.2eb8281foUVd15#2612f896detsz:~:text=%E6%9F%A5%E7%9C%8BAPI%2DKEY%E3%80%82-,%E8%8E%B7%E5%8F%96APP%2DID%E5%92%8CWORKSPACE,-%E8%BF%9B%E5%85%A5%E6%88%91%E7%9A%84%E5%BA%94%E7%94%A8).
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
// Retrieves session analysis results through a streaming API.
//
// Description:
//
// Before you use this API, make sure that you understand the billing methods and pricing for Tongyi Gold.
//
// # Prerequisites
//
// You have activated Alibaba Cloud Model Studio and Tongyi Gold.
//
// You must have a workspace ID. For more information, see [workspace identity](https://help.aliyun.com/zh/model-studio/developer-reference/get-app-id-and-workspace?spm=openapi-amp.newDocPublishment.0.0.2eb8281foUVd15#2612f896detsz:~:text=%E6%9F%A5%E7%9C%8BAPI%2DKEY%E3%80%82-,%E8%8E%B7%E5%8F%96APP%2DID%E5%92%8CWORKSPACE,-%E8%BF%9B%E5%85%A5%E6%88%91%E7%9A%84%E5%BA%94%E7%94%A8).
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
// Retrieves session analysis results through a streaming API.
//
// Description:
//
// Before you use this API, make sure that you understand the billing methods and pricing for Tongyi Gold.
//
// # Prerequisites
//
// You have activated Alibaba Cloud Model Studio and Tongyi Gold.
//
// You must have a workspace ID. For more information, see [workspace identity](https://help.aliyun.com/zh/model-studio/developer-reference/get-app-id-and-workspace?spm=openapi-amp.newDocPublishment.0.0.2eb8281foUVd15#2612f896detsz:~:text=%E6%9F%A5%E7%9C%8BAPI%2DKEY%E3%80%82-,%E8%8E%B7%E5%8F%96APP%2DID%E5%92%8CWORKSPACE,-%E8%BF%9B%E5%85%A5%E6%88%91%E7%9A%84%E5%BA%94%E7%94%A8).
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
// Generates a chat response from a document library. You can ask questions in natural language, and the system retrieves relevant information to provide a summarized answer.
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
// Generates a chat response from a document library. You can ask questions in natural language, and the system retrieves relevant information to provide a summarized answer.
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
// Submit a list of questions and retrieve results by calling the `GetChatQuestionResp` API.
//
// Description:
//
// Review the pricing and billing details for the Tongyi Dianjin product before you use this API.
//
// # Prerequisites
//
// Activate Alibaba Cloud Model Studio and Tongyi Dianjin.
//
// Obtain the \\`workspaceId\\`: Retrieve the [workspace identity](https://help.aliyun.com/zh/model-studio/developer-reference/get-app-id-and-workspace?spm=openapi-amp.newDocPublishment.0.0.2eb8281foUVd15#2612f896detsz:~:text=%E6%9F%A5%E7%9C%8BAPI%2DKEY%E3%80%82-,%E8%8E%B7%E5%8F%96APP%2DID%E5%92%8CWORKSPACE,-%E8%BF%9B%E5%85%A5%E6%88%91%E7%9A%84%E5%BA%94%E7%94%A8).
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
// Updates the title, metadata, and other information of a document.
//
// Description:
//
// *Prerequisites**
//
// - Activate the Alibaba Cloud Model Studio service and the Tongyi Gold service.
//
// - Obtain your workspace ID. For more information, see [workspace identity](https://help.aliyun.com/zh/model-studio/developer-reference/get-app-id-and-workspace?spm=openapi-amp.newDocPublishment.0.0.2eb8281foUVd15#2612f896detsz:~:text=%E6%9F%A5%E7%9C%8BAPI%2DKEY%E3%80%82-,%E8%8E%B7%E5%8F%96APP%2DID%E5%92%8CWORKSPACE,-%E8%BF%9B%E5%85%A5%E6%88%91%E7%9A%84%E5%BA%94%E7%94%A8).
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
// Update the text content of a document chunk in a document.
//
// Description:
//
// Before using this API, ensure you understand the billing methods and pricing of the Tongyi Gold product.
//
// # Prerequisites
//
// You must activate Alibaba Cloud Model Studio and Tongyi Gold services.
//
// Obtain the workspaceId. For more information, see the [workspace identity](https://help.aliyun.com/zh/model-studio/developer-reference/get-app-id-and-workspace?spm=openapi-amp.newDocPublishment.0.0.2eb8281foUVd15#2612f896detsz:~:text=%E6%9F%A5%E7%9C%8BAPI%2DKEY%E3%80%82-,%E8%8E%B7%E5%8F%96APP%2DID%E5%92%8CWORKSPACE,-%E8%BF%9B%E5%85%A5%E6%88%91%E7%9A%84%E5%BA%94%E7%94%A8).
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
// Updates a document library. You can update the library name, description, and index configuration.
//
// Description:
//
// *Prerequisites**
//
// - You have activated Alibaba Cloud Model Studio and Tongyi Dianjin.
//
// - Obtain the workspace ID. For more information, see [Get the workspace ID](https://help.aliyun.com/zh/model-studio/developer-reference/get-app-id-and-workspace?spm=openapi-amp.newDocPublishment.0.0.2eb8281foUVd15#2612f896detsz:~:text=%E6%9F%A5%E7%9C%8BAPI%2DKEY%E3%80%82-,%E8%8E%B7%E5%8F%96APP%2DID%E5%92%8CWORKSPACE,-%E8%BF%9B%E5%85%A5%E6%88%91%E7%9A%84%E5%BA%94%E7%94%A8).
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
// Updates a Q&A library. After the update, use the GenDocQaResult API to parse the Q&A pairs.
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
// Upload a document to a document library. The system parses the document, splits it into chunks, and builds an index.
//
// Description:
//
// Before you use this operation, review the Tongyi Dianjin pricing details.
//
// **Prerequisites**
//
// - You have activated Alibaba Cloud Model Studio and Tongyi Dianjin.
//
// - You have obtained a workspace ID. For more information, see [Get the workspace ID](https://help.aliyun.com/zh/model-studio/developer-reference/get-app-id-and-workspace?spm=openapi-amp.newDocPublishment.0.0.2eb8281foUVd15#2612f896detsz:~:text=%E6%9F%A5%E7%9C%8BAPI%2DKEY%E3%80%82-,%E8%8E%B7%E5%8F%96APP%2DID%E5%92%8CWORKSPACE,-%E8%BF%9B%E5%85%A5%E6%88%91%E7%9A%84%E5%BA%94%E7%94%A8).
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
