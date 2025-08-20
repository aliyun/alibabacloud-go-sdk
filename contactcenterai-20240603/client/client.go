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
	client.Endpoint, _err = client.GetEndpoint(dara.String("contactcenterai"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 语音文件分析任务极速版
//
// @param request - AnalyzeAudioSyncRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AnalyzeAudioSyncResponse
func (client *Client) AnalyzeAudioSyncWithSSE(workspaceId *string, appId *string, request *AnalyzeAudioSyncRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *AnalyzeAudioSyncResponse, _yieldErr chan error) {
	defer close(_yield)
	client.analyzeAudioSyncWithSSE_opYieldFunc(_yield, _yieldErr, workspaceId, appId, request, headers, runtime)
	return
}

// Summary:
//
// 语音文件分析任务极速版
//
// @param request - AnalyzeAudioSyncRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AnalyzeAudioSyncResponse
func (client *Client) AnalyzeAudioSyncWithOptions(workspaceId *string, appId *string, request *AnalyzeAudioSyncRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AnalyzeAudioSyncResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CategoryTags) {
		body["categoryTags"] = request.CategoryTags
	}

	if !dara.IsNil(request.CustomPrompt) {
		body["customPrompt"] = request.CustomPrompt
	}

	if !dara.IsNil(request.Fields) {
		body["fields"] = request.Fields
	}

	if !dara.IsNil(request.ModelCode) {
		body["modelCode"] = request.ModelCode
	}

	if !dara.IsNil(request.ResponseFormatType) {
		body["responseFormatType"] = request.ResponseFormatType
	}

	if !dara.IsNil(request.ResultTypes) {
		body["resultTypes"] = request.ResultTypes
	}

	if !dara.IsNil(request.ServiceInspection) {
		body["serviceInspection"] = request.ServiceInspection
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	if !dara.IsNil(request.TemplateIds) {
		body["templateIds"] = request.TemplateIds
	}

	if !dara.IsNil(request.Transcription) {
		body["transcription"] = request.Transcription
	}

	if !dara.IsNil(request.Variables) {
		body["variables"] = request.Variables
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AnalyzeAudioSync"),
		Version:     dara.String("2024-06-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/ccai/app/" + dara.PercentEncode(dara.StringValue(appId)) + "/analyzeAudioSync"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AnalyzeAudioSyncResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 语音文件分析任务极速版
//
// @param request - AnalyzeAudioSyncRequest
//
// @return AnalyzeAudioSyncResponse
func (client *Client) AnalyzeAudioSync(workspaceId *string, appId *string, request *AnalyzeAudioSyncRequest) (_result *AnalyzeAudioSyncResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AnalyzeAudioSyncResponse{}
	_body, _err := client.AnalyzeAudioSyncWithOptions(workspaceId, appId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 根据类型调用大模型
//
// @param request - AnalyzeConversationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AnalyzeConversationResponse
func (client *Client) AnalyzeConversationWithSSE(workspaceId *string, appId *string, request *AnalyzeConversationRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *AnalyzeConversationResponse, _yieldErr chan error) {
	defer close(_yield)
	client.analyzeConversationWithSSE_opYieldFunc(_yield, _yieldErr, workspaceId, appId, request, headers, runtime)
	return
}

// Summary:
//
// 根据类型调用大模型
//
// @param request - AnalyzeConversationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AnalyzeConversationResponse
func (client *Client) AnalyzeConversationWithOptions(workspaceId *string, appId *string, request *AnalyzeConversationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AnalyzeConversationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CategoryTags) {
		body["categoryTags"] = request.CategoryTags
	}

	if !dara.IsNil(request.CustomPrompt) {
		body["customPrompt"] = request.CustomPrompt
	}

	if !dara.IsNil(request.Dialogue) {
		body["dialogue"] = request.Dialogue
	}

	if !dara.IsNil(request.Examples) {
		body["examples"] = request.Examples
	}

	if !dara.IsNil(request.Fields) {
		body["fields"] = request.Fields
	}

	if !dara.IsNil(request.ModelCode) {
		body["modelCode"] = request.ModelCode
	}

	if !dara.IsNil(request.ResponseFormatType) {
		body["responseFormatType"] = request.ResponseFormatType
	}

	if !dara.IsNil(request.ResultTypes) {
		body["resultTypes"] = request.ResultTypes
	}

	if !dara.IsNil(request.SceneName) {
		body["sceneName"] = request.SceneName
	}

	if !dara.IsNil(request.ServiceInspection) {
		body["serviceInspection"] = request.ServiceInspection
	}

	if !dara.IsNil(request.SourceCallerUid) {
		body["sourceCallerUid"] = request.SourceCallerUid
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	if !dara.IsNil(request.TimeConstraintList) {
		body["timeConstraintList"] = request.TimeConstraintList
	}

	if !dara.IsNil(request.UserProfiles) {
		body["userProfiles"] = request.UserProfiles
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AnalyzeConversation"),
		Version:     dara.String("2024-06-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/ccai/app/" + dara.PercentEncode(dara.StringValue(appId)) + "/analyze_conversation"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AnalyzeConversationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 根据类型调用大模型
//
// @param request - AnalyzeConversationRequest
//
// @return AnalyzeConversationResponse
func (client *Client) AnalyzeConversation(workspaceId *string, appId *string, request *AnalyzeConversationRequest) (_result *AnalyzeConversationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AnalyzeConversationResponse{}
	_body, _err := client.AnalyzeConversationWithOptions(workspaceId, appId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 图片分析
//
// @param request - AnalyzeImageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AnalyzeImageResponse
func (client *Client) AnalyzeImageWithSSE(workspaceId *string, appId *string, request *AnalyzeImageRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *AnalyzeImageResponse, _yieldErr chan error) {
	defer close(_yield)
	client.analyzeImageWithSSE_opYieldFunc(_yield, _yieldErr, workspaceId, appId, request, headers, runtime)
	return
}

// Summary:
//
// 图片分析
//
// @param request - AnalyzeImageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AnalyzeImageResponse
func (client *Client) AnalyzeImageWithOptions(workspaceId *string, appId *string, request *AnalyzeImageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AnalyzeImageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageUrls) {
		body["imageUrls"] = request.ImageUrls
	}

	if !dara.IsNil(request.ResponseFormatType) {
		body["responseFormatType"] = request.ResponseFormatType
	}

	if !dara.IsNil(request.ResultTypes) {
		body["resultTypes"] = request.ResultTypes
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AnalyzeImage"),
		Version:     dara.String("2024-06-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/ccai/app/" + dara.PercentEncode(dara.StringValue(appId)) + "/analyzeImage"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &AnalyzeImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 图片分析
//
// @param request - AnalyzeImageRequest
//
// @return AnalyzeImageResponse
func (client *Client) AnalyzeImage(workspaceId *string, appId *string, request *AnalyzeImageRequest) (_result *AnalyzeImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AnalyzeImageResponse{}
	_body, _err := client.AnalyzeImageWithOptions(workspaceId, appId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建语音文件调用llm任务
//
// @param request - CreateTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTaskResponse
func (client *Client) CreateTaskWithOptions(workspaceId *string, appId *string, request *CreateTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CallBackUrl) {
		body["callBackUrl"] = request.CallBackUrl
	}

	if !dara.IsNil(request.CategoryTags) {
		body["categoryTags"] = request.CategoryTags
	}

	if !dara.IsNil(request.CustomPrompt) {
		body["customPrompt"] = request.CustomPrompt
	}

	if !dara.IsNil(request.Dialogue) {
		body["dialogue"] = request.Dialogue
	}

	if !dara.IsNil(request.Examples) {
		body["examples"] = request.Examples
	}

	if !dara.IsNil(request.Fields) {
		body["fields"] = request.Fields
	}

	if !dara.IsNil(request.ModelCode) {
		body["modelCode"] = request.ModelCode
	}

	if !dara.IsNil(request.ResponseFormatType) {
		body["responseFormatType"] = request.ResponseFormatType
	}

	if !dara.IsNil(request.ResultTypes) {
		body["resultTypes"] = request.ResultTypes
	}

	if !dara.IsNil(request.ServiceInspection) {
		body["serviceInspection"] = request.ServiceInspection
	}

	if !dara.IsNil(request.TaskType) {
		body["taskType"] = request.TaskType
	}

	if !dara.IsNil(request.TemplateIds) {
		body["templateIds"] = request.TemplateIds
	}

	if !dara.IsNil(request.Transcription) {
		body["transcription"] = request.Transcription
	}

	if !dara.IsNil(request.Variables) {
		body["variables"] = request.Variables
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateTask"),
		Version:     dara.String("2024-06-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/ccai/app/" + dara.PercentEncode(dara.StringValue(appId)) + "/createTask"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建语音文件调用llm任务
//
// @param request - CreateTaskRequest
//
// @return CreateTaskResponse
func (client *Client) CreateTask(workspaceId *string, appId *string, request *CreateTaskRequest) (_result *CreateTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTaskResponse{}
	_body, _err := client.CreateTaskWithOptions(workspaceId, appId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建热词
//
// @param request - CreateVocabRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVocabResponse
func (client *Client) CreateVocabWithOptions(request *CreateVocabRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateVocabResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AudioModelCode) {
		body["audioModelCode"] = request.AudioModelCode
	}

	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.WordWeightList) {
		body["wordWeightList"] = request.WordWeightList
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVocab"),
		Version:     dara.String("2024-06-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/vocab/createVocab"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &CreateVocabResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建热词
//
// @param request - CreateVocabRequest
//
// @return CreateVocabResponse
func (client *Client) CreateVocab(request *CreateVocabRequest) (_result *CreateVocabResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateVocabResponse{}
	_body, _err := client.CreateVocabWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删删除热词
//
// @param request - DeleteVocabRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVocabResponse
func (client *Client) DeleteVocabWithOptions(request *DeleteVocabRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteVocabResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.VocabularyId) {
		body["vocabularyId"] = request.VocabularyId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVocab"),
		Version:     dara.String("2024-06-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/vocab/deleteVocab"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &DeleteVocabResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删删除热词
//
// @param request - DeleteVocabRequest
//
// @return DeleteVocabResponse
func (client *Client) DeleteVocab(request *DeleteVocabRequest) (_result *DeleteVocabResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteVocabResponse{}
	_body, _err := client.DeleteVocabWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 通用图片分析
//
// @param request - GeneralAnalyzeImageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GeneralAnalyzeImageResponse
func (client *Client) GeneralAnalyzeImageWithSSE(workspaceId *string, appId *string, request *GeneralAnalyzeImageRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *GeneralAnalyzeImageResponse, _yieldErr chan error) {
	defer close(_yield)
	client.generalAnalyzeImageWithSSE_opYieldFunc(_yield, _yieldErr, workspaceId, appId, request, headers, runtime)
	return
}

// Summary:
//
// 通用图片分析
//
// @param request - GeneralAnalyzeImageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GeneralAnalyzeImageResponse
func (client *Client) GeneralAnalyzeImageWithOptions(workspaceId *string, appId *string, request *GeneralAnalyzeImageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GeneralAnalyzeImageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CustomPrompt) {
		body["customPrompt"] = request.CustomPrompt
	}

	if !dara.IsNil(request.ImageUrls) {
		body["imageUrls"] = request.ImageUrls
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	if !dara.IsNil(request.TemplateIds) {
		body["templateIds"] = request.TemplateIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GeneralAnalyzeImage"),
		Version:     dara.String("2024-06-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/ccai/app/" + dara.PercentEncode(dara.StringValue(appId)) + "/generalanalyzeImage"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GeneralAnalyzeImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 通用图片分析
//
// @param request - GeneralAnalyzeImageRequest
//
// @return GeneralAnalyzeImageResponse
func (client *Client) GeneralAnalyzeImage(workspaceId *string, appId *string, request *GeneralAnalyzeImageRequest) (_result *GeneralAnalyzeImageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GeneralAnalyzeImageResponse{}
	_body, _err := client.GeneralAnalyzeImageWithOptions(workspaceId, appId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 语音文件调用大模型获取结果
//
// @param tmpReq - GetTaskResultRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskResultResponse
func (client *Client) GetTaskResultWithOptions(tmpReq *GetTaskResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetTaskResultResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &GetTaskResultShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.RequiredFieldList) {
		request.RequiredFieldListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.RequiredFieldList, dara.String("requiredFieldList"), dara.String("simple"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.RequiredFieldListShrink) {
		query["requiredFieldList"] = request.RequiredFieldListShrink
	}

	if !dara.IsNil(request.TaskId) {
		query["taskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetTaskResult"),
		Version:     dara.String("2024-06-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/ccai/app/getTaskResult"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetTaskResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 语音文件调用大模型获取结果
//
// @param request - GetTaskResultRequest
//
// @return GetTaskResultResponse
func (client *Client) GetTaskResult(request *GetTaskResultRequest) (_result *GetTaskResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTaskResultResponse{}
	_body, _err := client.GetTaskResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取热词
//
// @param request - GetVocabRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVocabResponse
func (client *Client) GetVocabWithOptions(request *GetVocabRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetVocabResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.VocabularyId) {
		body["vocabularyId"] = request.VocabularyId
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GetVocab"),
		Version:     dara.String("2024-06-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/vocab/getVocab"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &GetVocabResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取热词
//
// @param request - GetVocabRequest
//
// @return GetVocabResponse
func (client *Client) GetVocab(request *GetVocabRequest) (_result *GetVocabResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetVocabResponse{}
	_body, _err := client.GetVocabWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 热词列表
//
// @param request - ListVocabRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVocabResponse
func (client *Client) ListVocabWithOptions(request *ListVocabRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ListVocabResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.WorkspaceId) {
		body["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVocab"),
		Version:     dara.String("2024-06-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/vocab/listVocab"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &ListVocabResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 热词列表
//
// @param request - ListVocabRequest
//
// @return ListVocabResponse
func (client *Client) ListVocab(request *ListVocabRequest) (_result *ListVocabResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListVocabResponse{}
	_body, _err := client.ListVocabWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # CCAI服务面API
//
// @param request - RunCompletionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunCompletionResponse
func (client *Client) RunCompletionWithSSE(workspaceId *string, appId *string, request *RunCompletionRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunCompletionResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runCompletionWithSSE_opYieldFunc(_yield, _yieldErr, workspaceId, appId, request, headers, runtime)
	return
}

// Summary:
//
// # CCAI服务面API
//
// @param request - RunCompletionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunCompletionResponse
func (client *Client) RunCompletionWithOptions(workspaceId *string, appId *string, request *RunCompletionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunCompletionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Dialogue) {
		body["Dialogue"] = request.Dialogue
	}

	if !dara.IsNil(request.Fields) {
		body["Fields"] = request.Fields
	}

	if !dara.IsNil(request.ModelCode) {
		body["ModelCode"] = request.ModelCode
	}

	if !dara.IsNil(request.ServiceInspection) {
		body["ServiceInspection"] = request.ServiceInspection
	}

	if !dara.IsNil(request.Stream) {
		body["Stream"] = request.Stream
	}

	if !dara.IsNil(request.TemplateIds) {
		body["TemplateIds"] = request.TemplateIds
	}

	if !dara.IsNil(request.ResponseFormatType) {
		body["responseFormatType"] = request.ResponseFormatType
	}

	if !dara.IsNil(request.Variables) {
		body["variables"] = request.Variables
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunCompletion"),
		Version:     dara.String("2024-06-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/ccai/app/" + dara.PercentEncode(dara.StringValue(appId)) + "/completion"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RunCompletionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # CCAI服务面API
//
// @param request - RunCompletionRequest
//
// @return RunCompletionResponse
func (client *Client) RunCompletion(workspaceId *string, appId *string, request *RunCompletionRequest) (_result *RunCompletionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunCompletionResponse{}
	_body, _err := client.RunCompletionWithOptions(workspaceId, appId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # CCAI服务面API
//
// @param request - RunCompletionMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunCompletionMessageResponse
func (client *Client) RunCompletionMessageWithSSE(workspaceId *string, appId *string, request *RunCompletionMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *RunCompletionMessageResponse, _yieldErr chan error) {
	defer close(_yield)
	client.runCompletionMessageWithSSE_opYieldFunc(_yield, _yieldErr, workspaceId, appId, request, headers, runtime)
	return
}

// Summary:
//
// # CCAI服务面API
//
// @param request - RunCompletionMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunCompletionMessageResponse
func (client *Client) RunCompletionMessageWithOptions(workspaceId *string, appId *string, request *RunCompletionMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RunCompletionMessageResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Messages) {
		body["Messages"] = request.Messages
	}

	if !dara.IsNil(request.ModelCode) {
		body["ModelCode"] = request.ModelCode
	}

	if !dara.IsNil(request.Stream) {
		body["Stream"] = request.Stream
	}

	if !dara.IsNil(request.ResponseFormatType) {
		body["responseFormatType"] = request.ResponseFormatType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunCompletionMessage"),
		Version:     dara.String("2024-06-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/ccai/app/" + dara.PercentEncode(dara.StringValue(appId)) + "/completion_message"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &RunCompletionMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # CCAI服务面API
//
// @param request - RunCompletionMessageRequest
//
// @return RunCompletionMessageResponse
func (client *Client) RunCompletionMessage(workspaceId *string, appId *string, request *RunCompletionMessageRequest) (_result *RunCompletionMessageResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunCompletionMessageResponse{}
	_body, _err := client.RunCompletionMessageWithOptions(workspaceId, appId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 修改热词
//
// @param request - UpdateVocabRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateVocabResponse
func (client *Client) UpdateVocabWithOptions(request *UpdateVocabRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *UpdateVocabResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["description"] = request.Description
	}

	if !dara.IsNil(request.Name) {
		body["name"] = request.Name
	}

	if !dara.IsNil(request.VocabularyId) {
		body["vocabularyId"] = request.VocabularyId
	}

	if !dara.IsNil(request.WordWeightList) {
		body["wordWeightList"] = request.WordWeightList
	}

	if !dara.IsNil(request.WorkspaceId) {
		body["workspaceId"] = request.WorkspaceId
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateVocab"),
		Version:     dara.String("2024-06-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/vocab/updateVocab"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	_result = &UpdateVocabResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改热词
//
// @param request - UpdateVocabRequest
//
// @return UpdateVocabResponse
func (client *Client) UpdateVocab(request *UpdateVocabRequest) (_result *UpdateVocabResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateVocabResponse{}
	_body, _err := client.UpdateVocabWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) analyzeAudioSyncWithSSE_opYieldFunc(_yield chan *AnalyzeAudioSyncResponse, _yieldErr chan error, workspaceId *string, appId *string, request *AnalyzeAudioSyncRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	_err := request.Validate()
	if _err != nil {
		_yieldErr <- _err
		return
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CategoryTags) {
		body["categoryTags"] = request.CategoryTags
	}

	if !dara.IsNil(request.CustomPrompt) {
		body["customPrompt"] = request.CustomPrompt
	}

	if !dara.IsNil(request.Fields) {
		body["fields"] = request.Fields
	}

	if !dara.IsNil(request.ModelCode) {
		body["modelCode"] = request.ModelCode
	}

	if !dara.IsNil(request.ResponseFormatType) {
		body["responseFormatType"] = request.ResponseFormatType
	}

	if !dara.IsNil(request.ResultTypes) {
		body["resultTypes"] = request.ResultTypes
	}

	if !dara.IsNil(request.ServiceInspection) {
		body["serviceInspection"] = request.ServiceInspection
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	if !dara.IsNil(request.TemplateIds) {
		body["templateIds"] = request.TemplateIds
	}

	if !dara.IsNil(request.Transcription) {
		body["transcription"] = request.Transcription
	}

	if !dara.IsNil(request.Variables) {
		body["variables"] = request.Variables
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AnalyzeAudioSync"),
		Version:     dara.String("2024-06-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/ccai/app/" + dara.PercentEncode(dara.StringValue(appId)) + "/analyzeAudioSync"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) analyzeConversationWithSSE_opYieldFunc(_yield chan *AnalyzeConversationResponse, _yieldErr chan error, workspaceId *string, appId *string, request *AnalyzeConversationRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	_err := request.Validate()
	if _err != nil {
		_yieldErr <- _err
		return
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CategoryTags) {
		body["categoryTags"] = request.CategoryTags
	}

	if !dara.IsNil(request.CustomPrompt) {
		body["customPrompt"] = request.CustomPrompt
	}

	if !dara.IsNil(request.Dialogue) {
		body["dialogue"] = request.Dialogue
	}

	if !dara.IsNil(request.Examples) {
		body["examples"] = request.Examples
	}

	if !dara.IsNil(request.Fields) {
		body["fields"] = request.Fields
	}

	if !dara.IsNil(request.ModelCode) {
		body["modelCode"] = request.ModelCode
	}

	if !dara.IsNil(request.ResponseFormatType) {
		body["responseFormatType"] = request.ResponseFormatType
	}

	if !dara.IsNil(request.ResultTypes) {
		body["resultTypes"] = request.ResultTypes
	}

	if !dara.IsNil(request.SceneName) {
		body["sceneName"] = request.SceneName
	}

	if !dara.IsNil(request.ServiceInspection) {
		body["serviceInspection"] = request.ServiceInspection
	}

	if !dara.IsNil(request.SourceCallerUid) {
		body["sourceCallerUid"] = request.SourceCallerUid
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	if !dara.IsNil(request.TimeConstraintList) {
		body["timeConstraintList"] = request.TimeConstraintList
	}

	if !dara.IsNil(request.UserProfiles) {
		body["userProfiles"] = request.UserProfiles
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AnalyzeConversation"),
		Version:     dara.String("2024-06-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/ccai/app/" + dara.PercentEncode(dara.StringValue(appId)) + "/analyze_conversation"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) analyzeImageWithSSE_opYieldFunc(_yield chan *AnalyzeImageResponse, _yieldErr chan error, workspaceId *string, appId *string, request *AnalyzeImageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	_err := request.Validate()
	if _err != nil {
		_yieldErr <- _err
		return
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ImageUrls) {
		body["imageUrls"] = request.ImageUrls
	}

	if !dara.IsNil(request.ResponseFormatType) {
		body["responseFormatType"] = request.ResponseFormatType
	}

	if !dara.IsNil(request.ResultTypes) {
		body["resultTypes"] = request.ResultTypes
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("AnalyzeImage"),
		Version:     dara.String("2024-06-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/ccai/app/" + dara.PercentEncode(dara.StringValue(appId)) + "/analyzeImage"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) generalAnalyzeImageWithSSE_opYieldFunc(_yield chan *GeneralAnalyzeImageResponse, _yieldErr chan error, workspaceId *string, appId *string, request *GeneralAnalyzeImageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	_err := request.Validate()
	if _err != nil {
		_yieldErr <- _err
		return
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.CustomPrompt) {
		body["customPrompt"] = request.CustomPrompt
	}

	if !dara.IsNil(request.ImageUrls) {
		body["imageUrls"] = request.ImageUrls
	}

	if !dara.IsNil(request.Stream) {
		body["stream"] = request.Stream
	}

	if !dara.IsNil(request.TemplateIds) {
		body["templateIds"] = request.TemplateIds
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("GeneralAnalyzeImage"),
		Version:     dara.String("2024-06-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/ccai/app/" + dara.PercentEncode(dara.StringValue(appId)) + "/generalanalyzeImage"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runCompletionWithSSE_opYieldFunc(_yield chan *RunCompletionResponse, _yieldErr chan error, workspaceId *string, appId *string, request *RunCompletionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	_err := request.Validate()
	if _err != nil {
		_yieldErr <- _err
		return
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Dialogue) {
		body["Dialogue"] = request.Dialogue
	}

	if !dara.IsNil(request.Fields) {
		body["Fields"] = request.Fields
	}

	if !dara.IsNil(request.ModelCode) {
		body["ModelCode"] = request.ModelCode
	}

	if !dara.IsNil(request.ServiceInspection) {
		body["ServiceInspection"] = request.ServiceInspection
	}

	if !dara.IsNil(request.Stream) {
		body["Stream"] = request.Stream
	}

	if !dara.IsNil(request.TemplateIds) {
		body["TemplateIds"] = request.TemplateIds
	}

	if !dara.IsNil(request.ResponseFormatType) {
		body["responseFormatType"] = request.ResponseFormatType
	}

	if !dara.IsNil(request.Variables) {
		body["variables"] = request.Variables
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunCompletion"),
		Version:     dara.String("2024-06-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/ccai/app/" + dara.PercentEncode(dara.StringValue(appId)) + "/completion"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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

func (client *Client) runCompletionMessageWithSSE_opYieldFunc(_yield chan *RunCompletionMessageResponse, _yieldErr chan error, workspaceId *string, appId *string, request *RunCompletionMessageRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	_err := request.Validate()
	if _err != nil {
		_yieldErr <- _err
		return
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Messages) {
		body["Messages"] = request.Messages
	}

	if !dara.IsNil(request.ModelCode) {
		body["ModelCode"] = request.ModelCode
	}

	if !dara.IsNil(request.Stream) {
		body["Stream"] = request.Stream
	}

	if !dara.IsNil(request.ResponseFormatType) {
		body["responseFormatType"] = request.ResponseFormatType
	}

	req := &openapiutil.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunCompletionMessage"),
		Version:     dara.String("2024-06-03"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/" + dara.PercentEncode(dara.StringValue(workspaceId)) + "/ccai/app/" + dara.PercentEncode(dara.StringValue(appId)) + "/completion_message"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("ROA"),
		ReqBodyType: dara.String("json"),
		BodyType:    dara.String("json"),
	}
	sseResp := make(chan *openapi.SSEResponse, 1)
	go client.CallSSEApi(params, req, runtime, sseResp, _yieldErr)
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
