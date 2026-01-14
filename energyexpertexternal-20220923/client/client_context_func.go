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
// 创建文件夹
//
// @param request - AddFolderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddFolderResponse
func (client *Client) AddFolderWithContext(ctx context.Context, request *AddFolderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddFolderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FolderName) {
		body["folderName"] = request.FolderName
	}

	if !dara.IsNil(request.ParentFolderId) {
		body["parentFolderId"] = request.ParentFolderId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AddFolder"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/aidoc/folder/add"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AddFolderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get Document Results
//
// Description:
//
// Users obtain real-time VL results by uploading a document URL.
//
// @param request - AnalyzeVlRealtimeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AnalyzeVlRealtimeResponse
func (client *Client) AnalyzeVlRealtimeWithContext(ctx context.Context, request *AnalyzeVlRealtimeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AnalyzeVlRealtimeResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FileUrl) {
		query["fileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.Language) {
		query["language"] = request.Language
	}

	if !dara.IsNil(request.TemplateId) {
		query["templateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AnalyzeVlRealtime"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/aidoc/document/analyzeVlRealtime"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AnalyzeVlRealtimeResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 策略执行状态反馈
//
// @param request - BatchSaveInstructionStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchSaveInstructionStatusResponse
func (client *Client) BatchSaveInstructionStatusWithContext(ctx context.Context, request *BatchSaveInstructionStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchSaveInstructionStatusResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FactoryId) {
		body["factoryId"] = request.FactoryId
	}

	if !dara.IsNil(request.PKey) {
		body["pKey"] = request.PKey
	}

	if !dara.IsNil(request.StatusList) {
		body["statusList"] = request.StatusList
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchSaveInstructionStatus"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/carbon/hvac/batchSaveInstructionStatus"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchSaveInstructionStatusResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 批量设置空调站点运行计划
//
// @param request - BatchUpdateSystemRunningPlanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchUpdateSystemRunningPlanResponse
func (client *Client) BatchUpdateSystemRunningPlanWithContext(ctx context.Context, request *BatchUpdateSystemRunningPlanRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchUpdateSystemRunningPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ControlType) {
		body["controlType"] = request.ControlType
	}

	if !dara.IsNil(request.DateType) {
		body["dateType"] = request.DateType
	}

	if !dara.IsNil(request.EarliestStartupTime) {
		body["earliestStartupTime"] = request.EarliestStartupTime
	}

	if !dara.IsNil(request.EndTime) {
		body["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.FactoryId) {
		body["factoryId"] = request.FactoryId
	}

	if !dara.IsNil(request.LatestShutdownTime) {
		body["latestShutdownTime"] = request.LatestShutdownTime
	}

	if !dara.IsNil(request.MaxCarbonDioxide) {
		body["maxCarbonDioxide"] = request.MaxCarbonDioxide
	}

	if !dara.IsNil(request.MaxTem) {
		body["maxTem"] = request.MaxTem
	}

	if !dara.IsNil(request.MinTem) {
		body["minTem"] = request.MinTem
	}

	if !dara.IsNil(request.SeasonMode) {
		body["seasonMode"] = request.SeasonMode
	}

	if !dara.IsNil(request.StartTime) {
		body["startTime"] = request.StartTime
	}

	if !dara.IsNil(request.SystemId) {
		body["systemId"] = request.SystemId
	}

	if !dara.IsNil(request.WorkingEndTime) {
		body["workingEndTime"] = request.WorkingEndTime
	}

	if !dara.IsNil(request.WorkingStartTime) {
		body["workingStartTime"] = request.WorkingStartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("BatchUpdateSystemRunningPlan"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/carbon/hvac/batchUpdateSystemRunningPlan"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &BatchUpdateSystemRunningPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Knowledge Base Q\\&A
//
// Description:
//
// - The interface provides Q&A services within the scope of the selected directory in the session.
//
// - The sessionId information is obtained through GetChatSessionList.
//
// - You can also create a new session via the CreateChatSession interface.
//
// @param request - ChatRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatResponse
func (client *Client) ChatWithContext(ctx context.Context, request *ChatRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ChatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocumentIds) {
		body["documentIds"] = request.DocumentIds
	}

	if !dara.IsNil(request.Question) {
		body["question"] = request.Question
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Chat"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/aidoc/document/chat"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ChatResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Knowledge Base Q&A
//
// Description:
//
// - The interface provides Q&A services within the scope of the selected directory in the session.
//
// - The sessionId information is obtained through GetChatSessionList.
//
// - You can also create a new session via the CreateChatSession interface.
//
// @param request - ChatStreamRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatStreamResponse
func (client *Client) ChatStreamWithSSECtx(ctx context.Context, request *ChatStreamRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *ChatStreamResponse, _yieldErr chan error) {
	defer close(_yield)
	client.chatStreamWithSSECtx_opYieldFunc(_yield, _yieldErr, ctx, request, headers, runtime)
	return
}

// Summary:
//
// Knowledge Base Q&A
//
// Description:
//
// - The interface provides Q&A services within the scope of the selected directory in the session.
//
// - The sessionId information is obtained through GetChatSessionList.
//
// - You can also create a new session via the CreateChatSession interface.
//
// @param request - ChatStreamRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChatStreamResponse
func (client *Client) ChatStreamWithContext(ctx context.Context, request *ChatStreamRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ChatStreamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocumentIds) {
		body["documentIds"] = request.DocumentIds
	}

	if !dara.IsNil(request.Question) {
		body["question"] = request.Question
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatStream"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/aidoc/document/chat/stream"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ChatStreamResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Create Q&A Window
//
// @param request - CreateChatSessionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateChatSessionResponse
func (client *Client) CreateChatSessionWithContext(ctx context.Context, request *CreateChatSessionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateChatSessionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FolderId) {
		body["folderId"] = request.FolderId
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.UserId) {
		body["userId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateChatSession"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/aidoc/document/chat/session/create"),
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
// 删除解析过的文件
//
// @param request - DeleteDocumentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDocumentResponse
func (client *Client) DeleteDocumentWithContext(ctx context.Context, request *DeleteDocumentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDocumentResponse, _err error) {
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
		Action:      dara.String("DeleteDocument"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/aidoc/document/delete"),
		Method:      dara.String("DELETE"),
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
// 删除文件夹
//
// @param request - DeleteFolderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFolderResponse
func (client *Client) DeleteFolderWithContext(ctx context.Context, request *DeleteFolderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteFolderResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FolderId) {
		query["folderId"] = request.FolderId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteFolder"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/aidoc/folder/delete"),
		Method:      dara.String("DELETE"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteFolderResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文档detail
//
// @param request - DetailDocumentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetailDocumentResponse
func (client *Client) DetailDocumentWithContext(ctx context.Context, request *DetailDocumentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DetailDocumentResponse, _err error) {
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
		Action:      dara.String("DetailDocument"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/aidoc/document/detail"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DetailDocumentResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑禁用设备
//
// @param request - EditProhibitedDevicesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EditProhibitedDevicesResponse
func (client *Client) EditProhibitedDevicesWithContext(ctx context.Context, request *EditProhibitedDevicesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EditProhibitedDevicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FactoryId) {
		body["factoryId"] = request.FactoryId
	}

	if !dara.IsNil(request.HvacDeviceConfigVOList) {
		body["hvacDeviceConfigVOList"] = request.HvacDeviceConfigVOList
	}

	if !dara.IsNil(request.SystemId) {
		body["systemId"] = request.SystemId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EditProhibitedDevices"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/carbon/hvac/editProhibitedDevices"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &EditProhibitedDevicesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 编辑不利区设备
//
// @param request - EditUnfavorableAreaDevicesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EditUnfavorableAreaDevicesResponse
func (client *Client) EditUnfavorableAreaDevicesWithContext(ctx context.Context, request *EditUnfavorableAreaDevicesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EditUnfavorableAreaDevicesResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FactoryId) {
		body["factoryId"] = request.FactoryId
	}

	if !dara.IsNil(request.HvacDeviceConfigVOList) {
		body["hvacDeviceConfigVOList"] = request.HvacDeviceConfigVOList
	}

	if !dara.IsNil(request.SystemId) {
		body["systemId"] = request.SystemId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("EditUnfavorableAreaDevices"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/carbon/hvac/editUnfavorableAreaDevices"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &EditUnfavorableAreaDevicesResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generate a report of the specified carbon footprint.
//
// Description:
//
// Given a product ID, this API initiates a task to calculate the carbon footprint result for the corresponding product. The task\\"s status can be checked using the `IsCompleted` API. Following the generation of results, other result inquiry APIs can be accessed for display content.
//
// @param request - GenerateResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateResultResponse
func (client *Client) GenerateResultWithContext(ctx context.Context, request *GenerateResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GenerateResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		body["code"] = request.Code
	}

	if !dara.IsNil(request.ProductId) {
		body["productId"] = request.ProductId
	}

	if !dara.IsNil(request.ProductType) {
		body["productType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GenerateResult"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/carbon/footprint/result/generate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GenerateResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This interface is used to obtain electrical constitute analysis data.
//
// @param request - GetAreaElecConstituteRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAreaElecConstituteResponse
func (client *Client) GetAreaElecConstituteWithContext(ctx context.Context, request *GetAreaElecConstituteRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAreaElecConstituteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		body["code"] = request.Code
	}

	if !dara.IsNil(request.Year) {
		body["year"] = request.Year
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetAreaElecConstitute"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/carbon/emission/analysis/elec/area"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetAreaElecConstituteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get trends in carbon emissions.
//
// @param request - GetCarbonEmissionTrendRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetCarbonEmissionTrendResponse
func (client *Client) GetCarbonEmissionTrendWithContext(ctx context.Context, request *GetCarbonEmissionTrendRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCarbonEmissionTrendResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		body["code"] = request.Code
	}

	if !dara.IsNil(request.ModuleCode) {
		body["moduleCode"] = request.ModuleCode
	}

	if !dara.IsNil(request.ModuleType) {
		body["moduleType"] = request.ModuleType
	}

	if !dara.IsNil(request.TrendType) {
		body["trendType"] = request.TrendType
	}

	if !dara.IsNil(request.YearList) {
		body["yearList"] = request.YearList
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetCarbonEmissionTrend"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/carbon/emission/analysis/trend"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetCarbonEmissionTrendResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get Q&A folder List
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChatFolderListResponse
func (client *Client) GetChatFolderListWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetChatFolderListResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetChatFolderList"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/aidoc/document/chat/folder/list"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetChatFolderListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Retrieve the historical documents of a session
//
// Description:
//
// - This API retrieves the list of historical documents within a session by passing in the session ID.
//
// - The sessionId information is obtained through GetChatSessionList.
//
// - A new session can also be created using the CreateChatSession interface.
//
// @param request - GetChatListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChatListResponse
func (client *Client) GetChatListWithContext(ctx context.Context, request *GetChatListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetChatListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		body["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		body["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetChatList"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/aidoc/document/chat/list"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetChatListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get Q&A Session List
//
// @param request - GetChatSessionListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetChatSessionListResponse
func (client *Client) GetChatSessionListWithContext(ctx context.Context, request *GetChatSessionListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetChatSessionListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CurrentPage) {
		body["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.PageSize) {
		body["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.UserId) {
		body["userId"] = request.UserId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetChatSessionList"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/aidoc/document/chat/session/list"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetChatSessionListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This interface is used to obtain the details category of a data item.
//
// Description:
//
// - obtain data item detail list under the current enterprise.
//
// @param request - GetDataItemListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataItemListResponse
func (client *Client) GetDataItemListWithContext(ctx context.Context, request *GetDataItemListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDataItemListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		body["code"] = request.Code
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataItemList"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/carbon/emission/data/item/list"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataItemListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the data quality evaluation results DQR and DQI.
//
// Description:
//
// This API returns the data quality evaluation results based on the user-provided product ID. It\\"s useful for understanding the data quality of the carbon emission factors for each inventory of the product.
//
// @param request - GetDataQualityAnalysisRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDataQualityAnalysisResponse
func (client *Client) GetDataQualityAnalysisWithContext(ctx context.Context, request *GetDataQualityAnalysisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDataQualityAnalysisResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		body["code"] = request.Code
	}

	if !dara.IsNil(request.DataQualityEvaluationType) {
		body["dataQualityEvaluationType"] = request.DataQualityEvaluationType
	}

	if !dara.IsNil(request.ProductId) {
		body["productId"] = request.ProductId
	}

	if !dara.IsNil(request.ProductType) {
		body["productType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDataQualityAnalysis"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/carbon/footprint/data/quality/analysis"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDataQualityAnalysisResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a device at a site that is activated by using an Alibaba Cloud account.
//
// Description:
//
//	  You can call this operation to query the parameters of a data collection device based on the device ID. If the verification is passed, the device parameters are returned. If the verification fails, a null value is returned.
//
//		- You can query the parameters of a single device by day. If data of the device does not exist, a null value is returned.
//
// - By current, endpoint only supports Hangzhou: `energyexpertexternal.cn-hangzhou.aliyuncs.com`.
//
// - To use this API, you need to be added to the whitelist. Please contact us through  <props="china">[official website](https://energy.aliyun.com/ifa/web/defaultLoginPage?adapter=aliyun#/consult?source=%E8%83%BD%E8%80%97%E5%AE%9D%E7%99%BB%E5%BD%95%E9%A1%B5%EF%BC%88WEB%EF%BC%89)
//
// <props="intl">[official website](https://energy.alibabacloud.com/common?adapter=aliyun&lang=en-US#/home/en) to apply for whitelist activation.
//
// @param request - GetDeviceInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeviceInfoResponse
func (client *Client) GetDeviceInfoWithContext(ctx context.Context, request *GetDeviceInfoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDeviceInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.DeviceId) {
		query["deviceId"] = request.DeviceId
	}

	if !dara.IsNil(request.Ds) {
		query["ds"] = request.Ds
	}

	if !dara.IsNil(request.FactoryId) {
		query["factoryId"] = request.FactoryId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeviceInfo"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/external/getDeviceInfo"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeviceInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the devices of a site that is activated by using an Alibaba Cloud account.
//
// Description:
//
//	  You can query the information about data collection devices of a site based on the ID of the site. If the verification is passed, the information about the devices of the site is returned. If the verification fails, a null value is returned.
//
//		- Virtual meters at the site are not returned.
//
// - By current, endpoint only supports Hangzhou: `energyexpertexternal.cn-hangzhou.aliyuncs.com`.
//
// - To use this API, you need to be added to the whitelist. Please contact us through  <props="china">[official website](https://energy.aliyun.com/ifa/web/defaultLoginPage?adapter=aliyun#/consult?source=%E8%83%BD%E8%80%97%E5%AE%9D%E7%99%BB%E5%BD%95%E9%A1%B5%EF%BC%88WEB%EF%BC%89)
//
// <props="intl">[official website](https://energy.alibabacloud.com/common?adapter=aliyun&lang=en-US#/home/en) to apply for whitelist activation.
//
// @param request - GetDeviceListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDeviceListResponse
func (client *Client) GetDeviceListWithContext(ctx context.Context, request *GetDeviceListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDeviceListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.FactoryId) {
		query["factoryId"] = request.FactoryId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDeviceList"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/external/getDeviceList"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDeviceListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// For Querying Information Extraction Result.
//
// The input parameter taskId is obtained from the taskId returned by the interfaces SubmitDocExtractionTaskAdvance or SubmitDocExtractionTask.
//
// The query results can reflect one of three statuses: processing, successfully completed, or failed.
//
// @param request - GetDocExtractionResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocExtractionResultResponse
func (client *Client) GetDocExtractionResultWithContext(ctx context.Context, request *GetDocExtractionResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDocExtractionResultResponse, _err error) {
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
		Action:      dara.String("GetDocExtractionResult"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/aidoc/document/getDocExtractionResult"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDocExtractionResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// For Querying Document Parsing Results.
//
// The input parameter taskId is obtained from the taskId returned by the interfaces SubmitDocParsingTaskAdvance or SubmitDocParsingTask.
//
// The query results can be one of three statuses: processing, successful, or failed.
//
// @param request - GetDocParsingResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocParsingResultResponse
func (client *Client) GetDocParsingResultWithContext(ctx context.Context, request *GetDocParsingResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDocParsingResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ReturnFormat) {
		body["returnFormat"] = request.ReturnFormat
	}

	if !dara.IsNil(request.TaskId) {
		body["taskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDocParsingResult"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/aidoc/document/getDocParsingResult"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDocParsingResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// [Important] The api is no longer maintained, please use GetDocExtractionResult, GetVLExtractionResult to get the extraction results.
//
// @param request - GetDocumentAnalyzeResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDocumentAnalyzeResultResponse
func (client *Client) GetDocumentAnalyzeResultWithContext(ctx context.Context, request *GetDocumentAnalyzeResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDocumentAnalyzeResultResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.JobId) {
		body["jobId"] = request.JobId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetDocumentAnalyzeResult"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/aidoc/document/getDocumentAnalyzeResult"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetDocumentAnalyzeResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This interface is used to obtain power composition analysis data.
//
// @param request - GetElecConstituteRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetElecConstituteResponse
func (client *Client) GetElecConstituteWithContext(ctx context.Context, request *GetElecConstituteRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetElecConstituteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		body["code"] = request.Code
	}

	if !dara.IsNil(request.Year) {
		body["year"] = request.Year
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetElecConstitute"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/carbon/emission/analysis/elec/constitute"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetElecConstituteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This interface is used to obtain power trend analysis data.
//
// @param request - GetElecTrendRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetElecTrendResponse
func (client *Client) GetElecTrendWithContext(ctx context.Context, request *GetElecTrendRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetElecTrendResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		body["code"] = request.Code
	}

	if !dara.IsNil(request.YearList) {
		body["yearList"] = request.YearList
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetElecTrend"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/carbon/emission/analysis/elec/trend"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetElecTrendResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the emission source composition.
//
// @param request - GetEmissionSourceConstituteRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEmissionSourceConstituteResponse
func (client *Client) GetEmissionSourceConstituteWithContext(ctx context.Context, request *GetEmissionSourceConstituteRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEmissionSourceConstituteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		body["code"] = request.Code
	}

	if !dara.IsNil(request.ModuleCode) {
		body["moduleCode"] = request.ModuleCode
	}

	if !dara.IsNil(request.ModuleType) {
		body["moduleType"] = request.ModuleType
	}

	if !dara.IsNil(request.Year) {
		body["year"] = request.Year
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEmissionSourceConstitute"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/carbon/emission/analysis/constitute"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEmissionSourceConstituteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get a summary of carbon emissions.
//
// @param request - GetEmissionSummaryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEmissionSummaryResponse
func (client *Client) GetEmissionSummaryWithContext(ctx context.Context, request *GetEmissionSummaryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEmissionSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		body["code"] = request.Code
	}

	if !dara.IsNil(request.ModuleCode) {
		body["moduleCode"] = request.ModuleCode
	}

	if !dara.IsNil(request.ModuleType) {
		body["moduleType"] = request.ModuleType
	}

	if !dara.IsNil(request.Year) {
		body["year"] = request.Year
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEmissionSummary"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/carbon/emission/analysis/summary"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEmissionSummaryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Gets the result details of the environmental impact category.
//
// Description:
//
// This API returns the emission amounts for various environmental impact categories at different levels for the given product ID. It helps understand the emission quantities for different environmental impact categories and inventories of the product.
//
// @param request - GetEpdInventoryConstituteRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEpdInventoryConstituteResponse
func (client *Client) GetEpdInventoryConstituteWithContext(ctx context.Context, request *GetEpdInventoryConstituteRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEpdInventoryConstituteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		body["code"] = request.Code
	}

	if !dara.IsNil(request.ProductId) {
		body["productId"] = request.ProductId
	}

	if !dara.IsNil(request.ProductType) {
		body["productType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEpdInventoryConstitute"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/carbon/footprint/result/epd/inventory/constitute"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEpdInventoryConstituteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtain the total amount of emissions for various environmental impacts.
//
// Description:
//
// This API takes a product ID from the user and returns the summary of environmental impact generated for the product. This info helps understand the overall emissions for different environmental impact categories of the product.
//
// @param request - GetEpdSummaryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetEpdSummaryResponse
func (client *Client) GetEpdSummaryWithContext(ctx context.Context, request *GetEpdSummaryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEpdSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		body["code"] = request.Code
	}

	if !dara.IsNil(request.ProductId) {
		body["productId"] = request.ProductId
	}

	if !dara.IsNil(request.ProductType) {
		body["productType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetEpdSummary"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/carbon/footprint/result/epd/summary"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetEpdSummaryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get the list of product carbon footprints.
//
// Description:
//
// With user-specified parameters such as enterprise code, current page, and page size, this API returns a list of matching product carbon footprints (or supply chain carbon footprints), including product names and product IDs. The product ID can be used as input parameters in other APIs to get the corresponding product\\"s detailed information.
//
// @param request - GetFootprintListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFootprintListResponse
func (client *Client) GetFootprintListWithContext(ctx context.Context, request *GetFootprintListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetFootprintListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		body["code"] = request.Code
	}

	if !dara.IsNil(request.CurrentPage) {
		body["currentPage"] = request.CurrentPage
	}

	if !dara.IsNil(request.PageSize) {
		body["pageSize"] = request.PageSize
	}

	if !dara.IsNil(request.ProductType) {
		body["productType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetFootprintList"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/carbon/footprint/product/list"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetFootprintListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This interface is used to obtain gas composition analysis.
//
// @param request - GetGasConstituteRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGasConstituteResponse
func (client *Client) GetGasConstituteWithContext(ctx context.Context, request *GetGasConstituteRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetGasConstituteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		body["code"] = request.Code
	}

	if !dara.IsNil(request.ModuleCode) {
		body["moduleCode"] = request.ModuleCode
	}

	if !dara.IsNil(request.ModuleType) {
		body["moduleType"] = request.ModuleType
	}

	if !dara.IsNil(request.Year) {
		body["year"] = request.Year
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGasConstitute"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/carbon/emission/analysis/gas/constitute"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGasConstituteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// obtain the active carbon reduction ranking list.
//
// Description:
//
// This interface returns a list of proactive carbon reduction information given product ID. It\\"s used to understand the carbon reduction efforts at various levels of the product.
//
// @param request - GetGwpBenchmarkListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGwpBenchmarkListResponse
func (client *Client) GetGwpBenchmarkListWithContext(ctx context.Context, request *GetGwpBenchmarkListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetGwpBenchmarkListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		body["code"] = request.Code
	}

	if !dara.IsNil(request.ProductId) {
		body["productId"] = request.ProductId
	}

	if !dara.IsNil(request.ProductType) {
		body["productType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGwpBenchmarkList"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/carbon/footprint/result/gwp/benchmark/list"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGwpBenchmarkListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This API is to obtain the total amount of active carbon reduction.
//
// Description:
//
// The API takes a product ID and returns data on the carbon emissions reduction along with a list of the top four contributors to carbon reduction. This info helps understand the total carbon reduction of the product and its main sources.
//
// @param request - GetGwpBenchmarkSummaryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGwpBenchmarkSummaryResponse
func (client *Client) GetGwpBenchmarkSummaryWithContext(ctx context.Context, request *GetGwpBenchmarkSummaryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetGwpBenchmarkSummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		body["code"] = request.Code
	}

	if !dara.IsNil(request.ProductId) {
		body["productId"] = request.ProductId
	}

	if !dara.IsNil(request.ProductType) {
		body["productType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGwpBenchmarkSummary"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/carbon/footprint/result/gwp/benchmark/summary"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGwpBenchmarkSummaryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Used to obtain the carbon emission composition analysis of a specified product. Carbon emission composition analysis includes two analysis dimensions: inventory and type. In the rendering effect, including a hierarchical list and pie chart.
//
// Description:
//
// Used to obtain the carbon emission composition analysis of a specified product. Carbon emission composition analysis includes two analysis dimensions: inventory and type. In the rendering effect, including a hierarchical list and pie chart.
//
// @param request - GetGwpInventoryConstituteRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGwpInventoryConstituteResponse
func (client *Client) GetGwpInventoryConstituteWithContext(ctx context.Context, request *GetGwpInventoryConstituteRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetGwpInventoryConstituteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		body["code"] = request.Code
	}

	if !dara.IsNil(request.ProductId) {
		body["productId"] = request.ProductId
	}

	if !dara.IsNil(request.ProductType) {
		body["productType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGwpInventoryConstitute"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/carbon/footprint/result/gwp/inventory/constitute"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGwpInventoryConstituteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This API is used to obtain the total carbon footprint of a product and the top four types of carbon footprint contribution.
//
// Description:
//
// Returns the total carbon footprint data for the user-specified product ID, along with details on the top four contributors to the carbon footprint, helping to understand the overall carbon footprint and its main components.
//
// @param request - GetGwpInventorySummaryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGwpInventorySummaryResponse
func (client *Client) GetGwpInventorySummaryWithContext(ctx context.Context, request *GetGwpInventorySummaryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetGwpInventorySummaryResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		body["code"] = request.Code
	}

	if !dara.IsNil(request.ProductId) {
		body["productId"] = request.ProductId
	}

	if !dara.IsNil(request.ProductType) {
		body["productType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetGwpInventorySummary"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/carbon/footprint/result/gwp/inventory/summary"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetGwpInventorySummaryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get the list of emissions in descending order under the specified environmental impact (methodType), specified aggregate level (group), and specified calculation mode (emissionType).
//
// Description:
//
// This interface retrieves a descending order list of emissions for a specified product ID, environmental impact method, group level, and calculation method. It\\"s used to understand various environmental impact emission scenarios.
//
// @param request - GetInventoryListRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetInventoryListResponse
func (client *Client) GetInventoryListWithContext(ctx context.Context, request *GetInventoryListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetInventoryListResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		body["code"] = request.Code
	}

	if !dara.IsNil(request.EmissionType) {
		body["emissionType"] = request.EmissionType
	}

	if !dara.IsNil(request.Group) {
		body["group"] = request.Group
	}

	if !dara.IsNil(request.MethodType) {
		body["methodType"] = request.MethodType
	}

	if !dara.IsNil(request.ProductId) {
		body["productId"] = request.ProductId
	}

	if !dara.IsNil(request.ProductType) {
		body["productType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetInventoryList"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/carbon/footprint/result/inventory/list"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetInventoryListResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the organizations and sites that are activated by using an Alibaba Cloud account. You cannot call this operation to query the organizations or sites that have not been activated in the console.
//
// Description:
//
//	If an activated site exists, the information about the site and the organization to which the site belongs is returned. If no activated site exists, null is returned.
//
// - By current, endpoint only supports Hangzhou: `energyexpertexternal.cn-hangzhou.aliyuncs.com`.
//
// - To use this API, you need to be added to the whitelist. Please contact us through  <props="china">[official website](https://energy.aliyun.com/ifa/web/defaultLoginPage?adapter=aliyun#/consult?source=%E8%83%BD%E8%80%97%E5%AE%9D%E7%99%BB%E5%BD%95%E9%A1%B5%EF%BC%88WEB%EF%BC%89)
//
// <props="intl">[official website](https://energy.alibabacloud.com/common?adapter=aliyun&lang=en-US#/home/en) to apply for whitelist activation.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOrgAndFactoryResponse
func (client *Client) GetOrgAndFactoryWithContext(ctx context.Context, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetOrgAndFactoryResponse, _err error) {
	req := &openapiutil.OpenApiRequest{
		Headers: headers,
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOrgAndFactory"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/external/getOrgAndFactory"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOrgAndFactoryResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This interface is used to obtain carbon inventory organization analysis data.
//
// @param request - GetOrgConstituteRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetOrgConstituteResponse
func (client *Client) GetOrgConstituteWithContext(ctx context.Context, request *GetOrgConstituteRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetOrgConstituteResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		body["code"] = request.Code
	}

	if !dara.IsNil(request.ModuleCode) {
		body["moduleCode"] = request.ModuleCode
	}

	if !dara.IsNil(request.ModuleType) {
		body["moduleType"] = request.ModuleType
	}

	if !dara.IsNil(request.Year) {
		body["year"] = request.Year
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetOrgConstitute"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/carbon/emission/analysis/org"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetOrgConstituteResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Obtains the oss address of the Product Carbon footprint Report.
//
// Description:
//
// With the user-specified product ID, this interface retrieves detailed information and download links for previously generated PCR reports. To use it, two conditions must be met: 1) the result has already been generated; 2) the PCR report has been created.
//
// @param request - GetPcrInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPcrInfoResponse
func (client *Client) GetPcrInfoWithContext(ctx context.Context, request *GetPcrInfoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPcrInfoResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		body["code"] = request.Code
	}

	if !dara.IsNil(request.ProductId) {
		body["productId"] = request.ProductId
	}

	if !dara.IsNil(request.ProductType) {
		body["productType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetPcrInfo"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/carbon/footprint/result/pcr/detail"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetPcrInfoResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get carbon reduction recommendations.
//
// Description:
//
// This API returns carbon reduction proposals based on the product ID. It\\"s useful for understanding optimization tips to reduce the carbon emissions associated with a product.
//
// @param request - GetReductionProposalRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetReductionProposalResponse
func (client *Client) GetReductionProposalWithContext(ctx context.Context, request *GetReductionProposalRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetReductionProposalResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		body["code"] = request.Code
	}

	if !dara.IsNil(request.DataQualityEvaluationType) {
		body["dataQualityEvaluationType"] = request.DataQualityEvaluationType
	}

	if !dara.IsNil(request.ProductId) {
		body["productId"] = request.ProductId
	}

	if !dara.IsNil(request.ProductType) {
		body["productType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetReductionProposal"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/carbon/footprint/result/reduction/proposal"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetReductionProposalResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// For Querying Qwen-VL Model Information Extraction Results.
//
// The input parameter taskId is obtained from the taskId returned by the interfaces SubmitVLExtractionTask or SubmitVLExtractionTaskAdvance.
//
// The query results can be in one of three statuses: processing, successfully completed, or failed.
//
// @param request - GetVLExtractionResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVLExtractionResultResponse
func (client *Client) GetVLExtractionResultWithContext(ctx context.Context, request *GetVLExtractionResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetVLExtractionResultResponse, _err error) {
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
		Action:      dara.String("GetVLExtractionResult"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/aidoc/document/getVLExtractionResult"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVLExtractionResultResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Check if the result generation is complete.
//
// Description:
//
// This API checks the completion status of generating a report. It should be used before calling other result APIs, as they will only display content once the report generation is complete.
//
// @param request - IsCompletedRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return IsCompletedResponse
func (client *Client) IsCompletedWithContext(ctx context.Context, request *IsCompletedRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *IsCompletedResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		body["code"] = request.Code
	}

	if !dara.IsNil(request.ProductId) {
		body["productId"] = request.ProductId
	}

	if !dara.IsNil(request.ProductType) {
		body["productType"] = request.ProductType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("IsCompleted"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/carbon/footprint/result/completed"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &IsCompletedResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This interface is used to push device measuring point data, such as power meter voltage and other data.
//
// @param request - PushDeviceDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushDeviceDataResponse
func (client *Client) PushDeviceDataWithContext(ctx context.Context, request *PushDeviceDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PushDeviceDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DeviceType) {
		body["deviceType"] = request.DeviceType
	}

	if !dara.IsNil(request.Devices) {
		body["devices"] = request.Devices
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushDeviceData"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/data/increment/push"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PushDeviceDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// This interface is used to push data items.
//
// Description:
//
// - This interface is used for individual data item data.
//
// - Data items can link data to services such as carbon footprints and carbon inventories.
//
// - Depending on the platform configuration, active data on a yearly and monthly basis is supported.
//
// @param request - PushItemDataRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushItemDataResponse
func (client *Client) PushItemDataWithContext(ctx context.Context, request *PushItemDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PushItemDataResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		body["code"] = request.Code
	}

	if !dara.IsNil(request.Items) {
		body["items"] = request.Items
	}

	if !dara.IsNil(request.Year) {
		body["year"] = request.Year
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("PushItemData"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/carbon/emission/data/item/push"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &PushItemDataResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Recalculate carbon emissions.
//
// Description:
//
// - After uploading the data items, you need to call this interface to recalculate the carbon inventory data.
//
// @param request - RecalculateCarbonEmissionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RecalculateCarbonEmissionResponse
func (client *Client) RecalculateCarbonEmissionWithContext(ctx context.Context, request *RecalculateCarbonEmissionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RecalculateCarbonEmissionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Code) {
		body["code"] = request.Code
	}

	if !dara.IsNil(request.Year) {
		body["year"] = request.Year
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RecalculateCarbonEmission"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/carbon/emission/data/item/recalculate"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RecalculateCarbonEmissionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 支持多文件夹ID或文件ID检索的RAG结果获取接口，供客户端自行加工结果并嵌入业务逻辑。
//
// @param request - RetrieveRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RetrieveResponse
func (client *Client) RetrieveWithContext(ctx context.Context, request *RetrieveRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RetrieveResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocumentIds) {
		body["documentIds"] = request.DocumentIds
	}

	if !dara.IsNil(request.FolderIds) {
		body["folderIds"] = request.FolderIds
	}

	if !dara.IsNil(request.PreRetrieveTopK) {
		body["preRetrieveTopK"] = request.PreRetrieveTopK
	}

	if !dara.IsNil(request.Query) {
		body["query"] = request.Query
	}

	if !dara.IsNil(request.RerankerThreshold) {
		body["rerankerThreshold"] = request.RerankerThreshold
	}

	if !dara.IsNil(request.TopK) {
		body["topK"] = request.TopK
	}

	if !dara.IsNil(request.UseReranker) {
		body["useReranker"] = request.UseReranker
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("Retrieve"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/aidoc/knowledgebase/retrieve"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RetrieveResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// [Important] This api is no longer maintained, please use the Chat api.
//
// @param request - SendDocumentAskQuestionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendDocumentAskQuestionResponse
func (client *Client) SendDocumentAskQuestionWithContext(ctx context.Context, request *SendDocumentAskQuestionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SendDocumentAskQuestionResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FolderId) {
		body["folderId"] = request.FolderId
	}

	if !dara.IsNil(request.Prompt) {
		body["prompt"] = request.Prompt
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendDocumentAskQuestion"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/aidoc/document/sendDocumentAskQuestion"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SendDocumentAskQuestionResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置运行计划
//
// @param request - SetRunningPlanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetRunningPlanResponse
func (client *Client) SetRunningPlanWithContext(ctx context.Context, request *SetRunningPlanRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SetRunningPlanResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ControlType) {
		body["controlType"] = request.ControlType
	}

	if !dara.IsNil(request.DateType) {
		body["dateType"] = request.DateType
	}

	if !dara.IsNil(request.EarliestStartupTime) {
		body["earliestStartupTime"] = request.EarliestStartupTime
	}

	if !dara.IsNil(request.EndTime) {
		body["endTime"] = request.EndTime
	}

	if !dara.IsNil(request.FactoryId) {
		body["factoryId"] = request.FactoryId
	}

	if !dara.IsNil(request.LatestShutdownTime) {
		body["latestShutdownTime"] = request.LatestShutdownTime
	}

	if !dara.IsNil(request.MaxCarbonDioxide) {
		body["maxCarbonDioxide"] = request.MaxCarbonDioxide
	}

	if !dara.IsNil(request.MaxTem) {
		body["maxTem"] = request.MaxTem
	}

	if !dara.IsNil(request.MinTem) {
		body["minTem"] = request.MinTem
	}

	if !dara.IsNil(request.PKey) {
		body["pKey"] = request.PKey
	}

	if !dara.IsNil(request.SeasonMode) {
		body["seasonMode"] = request.SeasonMode
	}

	if !dara.IsNil(request.StartTime) {
		body["startTime"] = request.StartTime
	}

	if !dara.IsNil(request.StatisticsTime) {
		body["statisticsTime"] = request.StatisticsTime
	}

	if !dara.IsNil(request.SystemId) {
		body["systemId"] = request.SystemId
	}

	if !dara.IsNil(request.WorkingEndTime) {
		body["workingEndTime"] = request.WorkingEndTime
	}

	if !dara.IsNil(request.WorkingStartTime) {
		body["workingStartTime"] = request.WorkingStartTime
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SetRunningPlan"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/carbon/hvac/setRunningPlan"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SetRunningPlanResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Extracts key information from documents using user-defined Key-Value or prompt templates. A taskId is returned upon successful execution for retrieving extraction results via GetDocExtractionResult.
//
// Supports:
//
// URL Upload: SubmitDocExtractionTask
//
// Local File Upload: SubmitDocExtractionTask
//
// @param request - SubmitDocExtractionTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitDocExtractionTaskResponse
func (client *Client) SubmitDocExtractionTaskWithContext(ctx context.Context, request *SubmitDocExtractionTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitDocExtractionTaskResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ExtractType) {
		query["extractType"] = request.ExtractType
	}

	if !dara.IsNil(request.FileName) {
		query["fileName"] = request.FileName
	}

	if !dara.IsNil(request.FileUrl) {
		query["fileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.FolderId) {
		query["folderId"] = request.FolderId
	}

	if !dara.IsNil(request.TemplateId) {
		query["templateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitDocExtractionTask"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/aidoc/document/submitDocExtractionTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitDocExtractionTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Parses text, tables, images, and more from documents. After execution, a taskId is returned for retrieving document parsing results via GetDocParsingResult.
//
// Supports:
//
// URL Upload: SubmitDocParsingTask
//
// Local File Upload: SubmitDocParsingTaskAdvance
//
// @param request - SubmitDocParsingTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitDocParsingTaskResponse
func (client *Client) SubmitDocParsingTaskWithContext(ctx context.Context, request *SubmitDocParsingTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitDocParsingTaskResponse, _err error) {
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

	if !dara.IsNil(request.FileUrl) {
		query["fileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.FolderId) {
		query["folderId"] = request.FolderId
	}

	if !dara.IsNil(request.NeedAnalyzeImg) {
		query["needAnalyzeImg"] = request.NeedAnalyzeImg
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitDocParsingTask"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/aidoc/document/submitDocParsingTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitDocParsingTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// [Important] The api is no longer maintained, please use the following api:
//
// Document parsing using SubmitDocParsingTask.
//
// Document extraction using SubmitVLExtractionTask, SubmitDocExtractionTask.
//
// @param request - SubmitDocumentAnalyzeJobRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitDocumentAnalyzeJobResponse
func (client *Client) SubmitDocumentAnalyzeJobWithContext(ctx context.Context, request *SubmitDocumentAnalyzeJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitDocumentAnalyzeJobResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.AnalysisType) {
		query["analysisType"] = request.AnalysisType
	}

	if !dara.IsNil(request.FileName) {
		query["fileName"] = request.FileName
	}

	if !dara.IsNil(request.FileUrl) {
		query["fileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.FolderId) {
		query["folderId"] = request.FolderId
	}

	if !dara.IsNil(request.TemplateId) {
		query["templateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitDocumentAnalyzeJob"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v1/aidoc/document/submitDocumentAnalyzeJob"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitDocumentAnalyzeJobResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Extracts key information from documents using KV templates or prompts with the Qwen-VL model, ideal for image extraction.
//
// Supports:
//
// URL Upload: SubmitVLExtractionTask.
//
// Local File Upload: SubmitVLExtractionTaskAdvance
//
// @param request - SubmitVLExtractionTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SubmitVLExtractionTaskResponse
func (client *Client) SubmitVLExtractionTaskWithContext(ctx context.Context, request *SubmitVLExtractionTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitVLExtractionTaskResponse, _err error) {
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

	if !dara.IsNil(request.FileUrl) {
		query["fileUrl"] = request.FileUrl
	}

	if !dara.IsNil(request.FolderId) {
		query["folderId"] = request.FolderId
	}

	if !dara.IsNil(request.TemplateId) {
		query["templateId"] = request.TemplateId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SubmitVLExtractionTask"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/aidoc/document/submitVLExtractionTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &SubmitVLExtractionTaskResponse{}
	_body, _err := client.CallApiWithCtx(ctx, params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) chatStreamWithSSECtx_opYieldFunc(_yield chan *ChatStreamResponse, _yieldErr chan error, ctx context.Context, request *ChatStreamRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DocumentIds) {
		body["documentIds"] = request.DocumentIds
	}

	if !dara.IsNil(request.Question) {
		body["question"] = request.Question
	}

	if !dara.IsNil(request.SessionId) {
		body["sessionId"] = request.SessionId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChatStream"),
		Version:     dara.String("2022-09-23"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/api/v2/aidoc/document/chat/stream"),
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
