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
	client.Endpoint, _err = client.GetEndpoint(dara.String("energyexpertexternal"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) _postOSSObject(bucketName *string, form map[string]interface{}) (_result map[string]interface{}, _err error) {
	request_ := dara.NewRequest()
	boundary := dara.GetBoundary()
	request_.Protocol = dara.String("HTTPS")
	request_.Method = dara.String("POST")
	request_.Pathname = dara.String("/")
	request_.Headers = map[string]*string{
		"host":       dara.String(dara.ToString(form["host"])),
		"date":       openapiutil.GetDateUTCString(),
		"user-agent": openapiutil.GetUserAgent(dara.String("")),
	}
	request_.Headers["content-type"] = dara.String("multipart/form-data; boundary=" + boundary)
	request_.Body = dara.ToFileForm(form, boundary)
	response_, _err := dara.DoRequest(request_, nil)
	if _err != nil {
		return nil, _err
	}

	_result, _err = _postOSSObject_opResponse(response_)
	if _err != nil {
		return nil, _err
	}

	return _result, nil
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
// 创建文件夹
//
// @param request - AddFolderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddFolderResponse
func (client *Client) AddFolderWithOptions(request *AddFolderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AddFolderResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = dara.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建文件夹
//
// @param request - AddFolderRequest
//
// @return AddFolderResponse
func (client *Client) AddFolder(request *AddFolderRequest) (_result *AddFolderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddFolderResponse{}
	_body, _err := client.AddFolderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) AnalyzeVlRealtimeWithOptions(request *AnalyzeVlRealtimeRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AnalyzeVlRealtimeResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return AnalyzeVlRealtimeResponse
func (client *Client) AnalyzeVlRealtime(request *AnalyzeVlRealtimeRequest) (_result *AnalyzeVlRealtimeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AnalyzeVlRealtimeResponse{}
	_body, _err := client.AnalyzeVlRealtimeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AnalyzeVlRealtimeAdvance(request *AnalyzeVlRealtimeAdvanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *AnalyzeVlRealtimeResponse, _err error) {
	// Step 0: init client
	if dara.IsNil(client.Credential) {
		_err = &openapi.ClientError{
			Code:    dara.String("InvalidCredentials"),
			Message: dara.String("Please set up the credentials correctly. If you are setting them through environment variables, please ensure that ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set correctly. See https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems for more details."),
		}
		return _result, _err
	}

	credentialModel, _err := client.Credential.GetCredential()
	if _err != nil {
		return _result, _err
	}

	accessKeyId := dara.StringValue(credentialModel.AccessKeyId)
	accessKeySecret := dara.StringValue(credentialModel.AccessKeySecret)
	securityToken := dara.StringValue(credentialModel.SecurityToken)
	credentialType := dara.StringValue(credentialModel.Type)
	openPlatformEndpoint := dara.StringValue(client.OpenPlatformEndpoint)
	if dara.IsNil(dara.String(openPlatformEndpoint)) || openPlatformEndpoint == "" {
		openPlatformEndpoint = "openplatform.aliyuncs.com"
	}

	if dara.IsNil(dara.String(credentialType)) {
		credentialType = "access_key"
	}

	authConfig := &openapiutil.Config{
		AccessKeyId:     dara.String(accessKeyId),
		AccessKeySecret: dara.String(accessKeySecret),
		SecurityToken:   dara.String(securityToken),
		Type:            dara.String(credentialType),
		Endpoint:        dara.String(openPlatformEndpoint),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openapi.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := map[string]*string{
		"Product":  dara.String("energyExpertExternal"),
		"RegionId": client.RegionId,
	}
	authReq := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(authRequest),
	}
	authParams := &openapiutil.Params{
		Action:      dara.String("AuthorizeFileUpload"),
		Version:     dara.String("2019-12-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	authResponse := map[string]interface{}{}
	fileObj := &dara.FileField{}
	ossHeader := map[string]interface{}{}
	tmpBody := map[string]interface{}{}
	useAccelerate := false
	authResponseBody := make(map[string]*string)
	analyzeVlRealtimeReq := &AnalyzeVlRealtimeRequest{}
	openapiutil.Convert(request, analyzeVlRealtimeReq)
	if !dara.IsNil(request.FileUrlObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.FileUrlObject,
			ContentType: dara.String(""),
		}
		ossHeader = map[string]interface{}{
			"host":                  dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(openapiutil.GetEndpoint(authResponseBody["Endpoint"], dara.Bool(useAccelerate), client.EndpointType)),
			"OSSAccessKeyId":        dara.StringValue(authResponseBody["AccessKeyId"]),
			"policy":                dara.StringValue(authResponseBody["EncodedPolicy"]),
			"Signature":             dara.StringValue(authResponseBody["Signature"]),
			"key":                   dara.StringValue(authResponseBody["ObjectKey"]),
			"file":                  fileObj,
			"success_action_status": "201",
		}
		_, _err = client._postOSSObject(authResponseBody["Bucket"], ossHeader)
		if _err != nil {
			return _result, _err
		}
		analyzeVlRealtimeReq.FileUrl = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	analyzeVlRealtimeResp, _err := client.AnalyzeVlRealtimeWithOptions(analyzeVlRealtimeReq, headers, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = analyzeVlRealtimeResp
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
func (client *Client) BatchSaveInstructionStatusWithOptions(request *BatchSaveInstructionStatusRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchSaveInstructionStatusResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return BatchSaveInstructionStatusResponse
func (client *Client) BatchSaveInstructionStatus(request *BatchSaveInstructionStatusRequest) (_result *BatchSaveInstructionStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchSaveInstructionStatusResponse{}
	_body, _err := client.BatchSaveInstructionStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) BatchUpdateSystemRunningPlanWithOptions(request *BatchUpdateSystemRunningPlanRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *BatchUpdateSystemRunningPlanResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return BatchUpdateSystemRunningPlanResponse
func (client *Client) BatchUpdateSystemRunningPlan(request *BatchUpdateSystemRunningPlanRequest) (_result *BatchUpdateSystemRunningPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchUpdateSystemRunningPlanResponse{}
	_body, _err := client.BatchUpdateSystemRunningPlanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ChatWithOptions(request *ChatRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ChatResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ChatResponse
func (client *Client) Chat(request *ChatRequest) (_result *ChatResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ChatResponse{}
	_body, _err := client.ChatWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) ChatStreamWithSSE(request *ChatStreamRequest, headers map[string]*string, runtime *dara.RuntimeOptions, _yield chan *ChatStreamResponse, _yieldErr chan error) {
	defer close(_yield)
	client.chatStreamWithSSE_opYieldFunc(_yield, _yieldErr, request, headers, runtime)
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
func (client *Client) ChatStreamWithOptions(request *ChatStreamRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *ChatStreamResponse, _err error) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err = request.Validate()
		if _err != nil {
			return _result, _err
		}
	}
	body := map[string]interface{}{}
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return ChatStreamResponse
func (client *Client) ChatStream(request *ChatStreamRequest) (_result *ChatStreamResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ChatStreamResponse{}
	_body, _err := client.ChatStreamWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) CreateChatSessionWithOptions(request *CreateChatSessionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *CreateChatSessionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return CreateChatSessionResponse
func (client *Client) CreateChatSession(request *CreateChatSessionRequest) (_result *CreateChatSessionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateChatSessionResponse{}
	_body, _err := client.CreateChatSessionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteDocumentWithOptions(request *DeleteDocumentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteDocumentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteDocumentResponse
func (client *Client) DeleteDocument(request *DeleteDocumentRequest) (_result *DeleteDocumentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDocumentResponse{}
	_body, _err := client.DeleteDocumentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DeleteFolderWithOptions(request *DeleteFolderRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DeleteFolderResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DeleteFolderResponse
func (client *Client) DeleteFolder(request *DeleteFolderRequest) (_result *DeleteFolderResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteFolderResponse{}
	_body, _err := client.DeleteFolderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) DetailDocumentWithOptions(request *DetailDocumentRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *DetailDocumentResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return DetailDocumentResponse
func (client *Client) DetailDocument(request *DetailDocumentRequest) (_result *DetailDocumentResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DetailDocumentResponse{}
	_body, _err := client.DetailDocumentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) EditProhibitedDevicesWithOptions(request *EditProhibitedDevicesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EditProhibitedDevicesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return EditProhibitedDevicesResponse
func (client *Client) EditProhibitedDevices(request *EditProhibitedDevicesRequest) (_result *EditProhibitedDevicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EditProhibitedDevicesResponse{}
	_body, _err := client.EditProhibitedDevicesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) EditUnfavorableAreaDevicesWithOptions(request *EditUnfavorableAreaDevicesRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *EditUnfavorableAreaDevicesResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return EditUnfavorableAreaDevicesResponse
func (client *Client) EditUnfavorableAreaDevices(request *EditUnfavorableAreaDevicesRequest) (_result *EditUnfavorableAreaDevicesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EditUnfavorableAreaDevicesResponse{}
	_body, _err := client.EditUnfavorableAreaDevicesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GenerateResultWithOptions(request *GenerateResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GenerateResultResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GenerateResultResponse
func (client *Client) GenerateResult(request *GenerateResultRequest) (_result *GenerateResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GenerateResultResponse{}
	_body, _err := client.GenerateResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetAreaElecConstituteWithOptions(request *GetAreaElecConstituteRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetAreaElecConstituteResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetAreaElecConstituteResponse
func (client *Client) GetAreaElecConstitute(request *GetAreaElecConstituteRequest) (_result *GetAreaElecConstituteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAreaElecConstituteResponse{}
	_body, _err := client.GetAreaElecConstituteWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetCarbonEmissionTrendWithOptions(request *GetCarbonEmissionTrendRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetCarbonEmissionTrendResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetCarbonEmissionTrendResponse
func (client *Client) GetCarbonEmissionTrend(request *GetCarbonEmissionTrendRequest) (_result *GetCarbonEmissionTrendResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCarbonEmissionTrendResponse{}
	_body, _err := client.GetCarbonEmissionTrendWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetChatFolderListWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetChatFolderListResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetChatFolderListResponse
func (client *Client) GetChatFolderList() (_result *GetChatFolderListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetChatFolderListResponse{}
	_body, _err := client.GetChatFolderListWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetChatListWithOptions(request *GetChatListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetChatListResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetChatListResponse
func (client *Client) GetChatList(request *GetChatListRequest) (_result *GetChatListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetChatListResponse{}
	_body, _err := client.GetChatListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetChatSessionListWithOptions(request *GetChatSessionListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetChatSessionListResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetChatSessionListResponse
func (client *Client) GetChatSessionList(request *GetChatSessionListRequest) (_result *GetChatSessionListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetChatSessionListResponse{}
	_body, _err := client.GetChatSessionListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetDataItemListWithOptions(request *GetDataItemListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDataItemListResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetDataItemListResponse
func (client *Client) GetDataItemList(request *GetDataItemListRequest) (_result *GetDataItemListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDataItemListResponse{}
	_body, _err := client.GetDataItemListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetDataQualityAnalysisWithOptions(request *GetDataQualityAnalysisRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDataQualityAnalysisResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetDataQualityAnalysisResponse
func (client *Client) GetDataQualityAnalysis(request *GetDataQualityAnalysisRequest) (_result *GetDataQualityAnalysisResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDataQualityAnalysisResponse{}
	_body, _err := client.GetDataQualityAnalysisWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetDeviceInfoWithOptions(request *GetDeviceInfoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDeviceInfoResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetDeviceInfoResponse
func (client *Client) GetDeviceInfo(request *GetDeviceInfoRequest) (_result *GetDeviceInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDeviceInfoResponse{}
	_body, _err := client.GetDeviceInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetDeviceListWithOptions(request *GetDeviceListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDeviceListResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetDeviceListResponse
func (client *Client) GetDeviceList(request *GetDeviceListRequest) (_result *GetDeviceListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDeviceListResponse{}
	_body, _err := client.GetDeviceListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetDocExtractionResultWithOptions(request *GetDocExtractionResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDocExtractionResultResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetDocExtractionResultResponse
func (client *Client) GetDocExtractionResult(request *GetDocExtractionResultRequest) (_result *GetDocExtractionResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDocExtractionResultResponse{}
	_body, _err := client.GetDocExtractionResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetDocParsingResultWithOptions(request *GetDocParsingResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDocParsingResultResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetDocParsingResultResponse
func (client *Client) GetDocParsingResult(request *GetDocParsingResultRequest) (_result *GetDocParsingResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDocParsingResultResponse{}
	_body, _err := client.GetDocParsingResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetDocumentAnalyzeResultWithOptions(request *GetDocumentAnalyzeResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetDocumentAnalyzeResultResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetDocumentAnalyzeResultResponse
func (client *Client) GetDocumentAnalyzeResult(request *GetDocumentAnalyzeResultRequest) (_result *GetDocumentAnalyzeResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDocumentAnalyzeResultResponse{}
	_body, _err := client.GetDocumentAnalyzeResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetElecConstituteWithOptions(request *GetElecConstituteRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetElecConstituteResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetElecConstituteResponse
func (client *Client) GetElecConstitute(request *GetElecConstituteRequest) (_result *GetElecConstituteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetElecConstituteResponse{}
	_body, _err := client.GetElecConstituteWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetElecTrendWithOptions(request *GetElecTrendRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetElecTrendResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetElecTrendResponse
func (client *Client) GetElecTrend(request *GetElecTrendRequest) (_result *GetElecTrendResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetElecTrendResponse{}
	_body, _err := client.GetElecTrendWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetEmissionSourceConstituteWithOptions(request *GetEmissionSourceConstituteRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEmissionSourceConstituteResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetEmissionSourceConstituteResponse
func (client *Client) GetEmissionSourceConstitute(request *GetEmissionSourceConstituteRequest) (_result *GetEmissionSourceConstituteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEmissionSourceConstituteResponse{}
	_body, _err := client.GetEmissionSourceConstituteWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetEmissionSummaryWithOptions(request *GetEmissionSummaryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEmissionSummaryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetEmissionSummaryResponse
func (client *Client) GetEmissionSummary(request *GetEmissionSummaryRequest) (_result *GetEmissionSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEmissionSummaryResponse{}
	_body, _err := client.GetEmissionSummaryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetEpdInventoryConstituteWithOptions(request *GetEpdInventoryConstituteRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEpdInventoryConstituteResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetEpdInventoryConstituteResponse
func (client *Client) GetEpdInventoryConstitute(request *GetEpdInventoryConstituteRequest) (_result *GetEpdInventoryConstituteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEpdInventoryConstituteResponse{}
	_body, _err := client.GetEpdInventoryConstituteWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetEpdSummaryWithOptions(request *GetEpdSummaryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetEpdSummaryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetEpdSummaryResponse
func (client *Client) GetEpdSummary(request *GetEpdSummaryRequest) (_result *GetEpdSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetEpdSummaryResponse{}
	_body, _err := client.GetEpdSummaryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetFootprintListWithOptions(request *GetFootprintListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetFootprintListResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetFootprintListResponse
func (client *Client) GetFootprintList(request *GetFootprintListRequest) (_result *GetFootprintListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFootprintListResponse{}
	_body, _err := client.GetFootprintListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetGasConstituteWithOptions(request *GetGasConstituteRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetGasConstituteResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetGasConstituteResponse
func (client *Client) GetGasConstitute(request *GetGasConstituteRequest) (_result *GetGasConstituteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetGasConstituteResponse{}
	_body, _err := client.GetGasConstituteWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetGwpBenchmarkListWithOptions(request *GetGwpBenchmarkListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetGwpBenchmarkListResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetGwpBenchmarkListResponse
func (client *Client) GetGwpBenchmarkList(request *GetGwpBenchmarkListRequest) (_result *GetGwpBenchmarkListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetGwpBenchmarkListResponse{}
	_body, _err := client.GetGwpBenchmarkListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetGwpBenchmarkSummaryWithOptions(request *GetGwpBenchmarkSummaryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetGwpBenchmarkSummaryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetGwpBenchmarkSummaryResponse
func (client *Client) GetGwpBenchmarkSummary(request *GetGwpBenchmarkSummaryRequest) (_result *GetGwpBenchmarkSummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetGwpBenchmarkSummaryResponse{}
	_body, _err := client.GetGwpBenchmarkSummaryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetGwpInventoryConstituteWithOptions(request *GetGwpInventoryConstituteRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetGwpInventoryConstituteResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetGwpInventoryConstituteResponse
func (client *Client) GetGwpInventoryConstitute(request *GetGwpInventoryConstituteRequest) (_result *GetGwpInventoryConstituteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetGwpInventoryConstituteResponse{}
	_body, _err := client.GetGwpInventoryConstituteWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetGwpInventorySummaryWithOptions(request *GetGwpInventorySummaryRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetGwpInventorySummaryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetGwpInventorySummaryResponse
func (client *Client) GetGwpInventorySummary(request *GetGwpInventorySummaryRequest) (_result *GetGwpInventorySummaryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetGwpInventorySummaryResponse{}
	_body, _err := client.GetGwpInventorySummaryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetInventoryListWithOptions(request *GetInventoryListRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetInventoryListResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetInventoryListResponse
func (client *Client) GetInventoryList(request *GetInventoryListRequest) (_result *GetInventoryListResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInventoryListResponse{}
	_body, _err := client.GetInventoryListWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetOrgAndFactoryWithOptions(headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetOrgAndFactoryResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetOrgAndFactoryResponse
func (client *Client) GetOrgAndFactory() (_result *GetOrgAndFactoryResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetOrgAndFactoryResponse{}
	_body, _err := client.GetOrgAndFactoryWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetOrgConstituteWithOptions(request *GetOrgConstituteRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetOrgConstituteResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetOrgConstituteResponse
func (client *Client) GetOrgConstitute(request *GetOrgConstituteRequest) (_result *GetOrgConstituteResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetOrgConstituteResponse{}
	_body, _err := client.GetOrgConstituteWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetPcrInfoWithOptions(request *GetPcrInfoRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetPcrInfoResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetPcrInfoResponse
func (client *Client) GetPcrInfo(request *GetPcrInfoRequest) (_result *GetPcrInfoResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPcrInfoResponse{}
	_body, _err := client.GetPcrInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetReductionProposalWithOptions(request *GetReductionProposalRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetReductionProposalResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetReductionProposalResponse
func (client *Client) GetReductionProposal(request *GetReductionProposalRequest) (_result *GetReductionProposalResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetReductionProposalResponse{}
	_body, _err := client.GetReductionProposalWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) GetVLExtractionResultWithOptions(request *GetVLExtractionResultRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *GetVLExtractionResultResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return GetVLExtractionResultResponse
func (client *Client) GetVLExtractionResult(request *GetVLExtractionResultRequest) (_result *GetVLExtractionResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetVLExtractionResultResponse{}
	_body, _err := client.GetVLExtractionResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) IsCompletedWithOptions(request *IsCompletedRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *IsCompletedResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return IsCompletedResponse
func (client *Client) IsCompleted(request *IsCompletedRequest) (_result *IsCompletedResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &IsCompletedResponse{}
	_body, _err := client.IsCompletedWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) PushDeviceDataWithOptions(request *PushDeviceDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PushDeviceDataResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return PushDeviceDataResponse
func (client *Client) PushDeviceData(request *PushDeviceDataRequest) (_result *PushDeviceDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PushDeviceDataResponse{}
	_body, _err := client.PushDeviceDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) PushItemDataWithOptions(request *PushItemDataRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *PushItemDataResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return PushItemDataResponse
func (client *Client) PushItemData(request *PushItemDataRequest) (_result *PushItemDataResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PushItemDataResponse{}
	_body, _err := client.PushItemDataWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) RecalculateCarbonEmissionWithOptions(request *RecalculateCarbonEmissionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *RecalculateCarbonEmissionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return RecalculateCarbonEmissionResponse
func (client *Client) RecalculateCarbonEmission(request *RecalculateCarbonEmissionRequest) (_result *RecalculateCarbonEmissionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RecalculateCarbonEmissionResponse{}
	_body, _err := client.RecalculateCarbonEmissionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SendDocumentAskQuestionWithOptions(request *SendDocumentAskQuestionRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SendDocumentAskQuestionResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SendDocumentAskQuestionResponse
func (client *Client) SendDocumentAskQuestion(request *SendDocumentAskQuestionRequest) (_result *SendDocumentAskQuestionResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SendDocumentAskQuestionResponse{}
	_body, _err := client.SendDocumentAskQuestionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SetRunningPlanWithOptions(request *SetRunningPlanRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SetRunningPlanResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SetRunningPlanResponse
func (client *Client) SetRunningPlan(request *SetRunningPlanRequest) (_result *SetRunningPlanResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SetRunningPlanResponse{}
	_body, _err := client.SetRunningPlanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
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
func (client *Client) SubmitDocExtractionTaskWithOptions(request *SubmitDocExtractionTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitDocExtractionTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SubmitDocExtractionTaskResponse
func (client *Client) SubmitDocExtractionTask(request *SubmitDocExtractionTaskRequest) (_result *SubmitDocExtractionTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitDocExtractionTaskResponse{}
	_body, _err := client.SubmitDocExtractionTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitDocExtractionTaskAdvance(request *SubmitDocExtractionTaskAdvanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitDocExtractionTaskResponse, _err error) {
	// Step 0: init client
	if dara.IsNil(client.Credential) {
		_err = &openapi.ClientError{
			Code:    dara.String("InvalidCredentials"),
			Message: dara.String("Please set up the credentials correctly. If you are setting them through environment variables, please ensure that ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set correctly. See https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems for more details."),
		}
		return _result, _err
	}

	credentialModel, _err := client.Credential.GetCredential()
	if _err != nil {
		return _result, _err
	}

	accessKeyId := dara.StringValue(credentialModel.AccessKeyId)
	accessKeySecret := dara.StringValue(credentialModel.AccessKeySecret)
	securityToken := dara.StringValue(credentialModel.SecurityToken)
	credentialType := dara.StringValue(credentialModel.Type)
	openPlatformEndpoint := dara.StringValue(client.OpenPlatformEndpoint)
	if dara.IsNil(dara.String(openPlatformEndpoint)) || openPlatformEndpoint == "" {
		openPlatformEndpoint = "openplatform.aliyuncs.com"
	}

	if dara.IsNil(dara.String(credentialType)) {
		credentialType = "access_key"
	}

	authConfig := &openapiutil.Config{
		AccessKeyId:     dara.String(accessKeyId),
		AccessKeySecret: dara.String(accessKeySecret),
		SecurityToken:   dara.String(securityToken),
		Type:            dara.String(credentialType),
		Endpoint:        dara.String(openPlatformEndpoint),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openapi.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := map[string]*string{
		"Product":  dara.String("energyExpertExternal"),
		"RegionId": client.RegionId,
	}
	authReq := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(authRequest),
	}
	authParams := &openapiutil.Params{
		Action:      dara.String("AuthorizeFileUpload"),
		Version:     dara.String("2019-12-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	authResponse := map[string]interface{}{}
	fileObj := &dara.FileField{}
	ossHeader := map[string]interface{}{}
	tmpBody := map[string]interface{}{}
	useAccelerate := false
	authResponseBody := make(map[string]*string)
	submitDocExtractionTaskReq := &SubmitDocExtractionTaskRequest{}
	openapiutil.Convert(request, submitDocExtractionTaskReq)
	if !dara.IsNil(request.FileUrlObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.FileUrlObject,
			ContentType: dara.String(""),
		}
		ossHeader = map[string]interface{}{
			"host":                  dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(openapiutil.GetEndpoint(authResponseBody["Endpoint"], dara.Bool(useAccelerate), client.EndpointType)),
			"OSSAccessKeyId":        dara.StringValue(authResponseBody["AccessKeyId"]),
			"policy":                dara.StringValue(authResponseBody["EncodedPolicy"]),
			"Signature":             dara.StringValue(authResponseBody["Signature"]),
			"key":                   dara.StringValue(authResponseBody["ObjectKey"]),
			"file":                  fileObj,
			"success_action_status": "201",
		}
		_, _err = client._postOSSObject(authResponseBody["Bucket"], ossHeader)
		if _err != nil {
			return _result, _err
		}
		submitDocExtractionTaskReq.FileUrl = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	submitDocExtractionTaskResp, _err := client.SubmitDocExtractionTaskWithOptions(submitDocExtractionTaskReq, headers, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = submitDocExtractionTaskResp
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
func (client *Client) SubmitDocParsingTaskWithOptions(request *SubmitDocParsingTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitDocParsingTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SubmitDocParsingTaskResponse
func (client *Client) SubmitDocParsingTask(request *SubmitDocParsingTaskRequest) (_result *SubmitDocParsingTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitDocParsingTaskResponse{}
	_body, _err := client.SubmitDocParsingTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitDocParsingTaskAdvance(request *SubmitDocParsingTaskAdvanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitDocParsingTaskResponse, _err error) {
	// Step 0: init client
	if dara.IsNil(client.Credential) {
		_err = &openapi.ClientError{
			Code:    dara.String("InvalidCredentials"),
			Message: dara.String("Please set up the credentials correctly. If you are setting them through environment variables, please ensure that ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set correctly. See https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems for more details."),
		}
		return _result, _err
	}

	credentialModel, _err := client.Credential.GetCredential()
	if _err != nil {
		return _result, _err
	}

	accessKeyId := dara.StringValue(credentialModel.AccessKeyId)
	accessKeySecret := dara.StringValue(credentialModel.AccessKeySecret)
	securityToken := dara.StringValue(credentialModel.SecurityToken)
	credentialType := dara.StringValue(credentialModel.Type)
	openPlatformEndpoint := dara.StringValue(client.OpenPlatformEndpoint)
	if dara.IsNil(dara.String(openPlatformEndpoint)) || openPlatformEndpoint == "" {
		openPlatformEndpoint = "openplatform.aliyuncs.com"
	}

	if dara.IsNil(dara.String(credentialType)) {
		credentialType = "access_key"
	}

	authConfig := &openapiutil.Config{
		AccessKeyId:     dara.String(accessKeyId),
		AccessKeySecret: dara.String(accessKeySecret),
		SecurityToken:   dara.String(securityToken),
		Type:            dara.String(credentialType),
		Endpoint:        dara.String(openPlatformEndpoint),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openapi.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := map[string]*string{
		"Product":  dara.String("energyExpertExternal"),
		"RegionId": client.RegionId,
	}
	authReq := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(authRequest),
	}
	authParams := &openapiutil.Params{
		Action:      dara.String("AuthorizeFileUpload"),
		Version:     dara.String("2019-12-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	authResponse := map[string]interface{}{}
	fileObj := &dara.FileField{}
	ossHeader := map[string]interface{}{}
	tmpBody := map[string]interface{}{}
	useAccelerate := false
	authResponseBody := make(map[string]*string)
	submitDocParsingTaskReq := &SubmitDocParsingTaskRequest{}
	openapiutil.Convert(request, submitDocParsingTaskReq)
	if !dara.IsNil(request.FileUrlObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.FileUrlObject,
			ContentType: dara.String(""),
		}
		ossHeader = map[string]interface{}{
			"host":                  dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(openapiutil.GetEndpoint(authResponseBody["Endpoint"], dara.Bool(useAccelerate), client.EndpointType)),
			"OSSAccessKeyId":        dara.StringValue(authResponseBody["AccessKeyId"]),
			"policy":                dara.StringValue(authResponseBody["EncodedPolicy"]),
			"Signature":             dara.StringValue(authResponseBody["Signature"]),
			"key":                   dara.StringValue(authResponseBody["ObjectKey"]),
			"file":                  fileObj,
			"success_action_status": "201",
		}
		_, _err = client._postOSSObject(authResponseBody["Bucket"], ossHeader)
		if _err != nil {
			return _result, _err
		}
		submitDocParsingTaskReq.FileUrl = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	submitDocParsingTaskResp, _err := client.SubmitDocParsingTaskWithOptions(submitDocParsingTaskReq, headers, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = submitDocParsingTaskResp
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
func (client *Client) SubmitDocumentAnalyzeJobWithOptions(request *SubmitDocumentAnalyzeJobRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitDocumentAnalyzeJobResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SubmitDocumentAnalyzeJobResponse
func (client *Client) SubmitDocumentAnalyzeJob(request *SubmitDocumentAnalyzeJobRequest) (_result *SubmitDocumentAnalyzeJobResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitDocumentAnalyzeJobResponse{}
	_body, _err := client.SubmitDocumentAnalyzeJobWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitDocumentAnalyzeJobAdvance(request *SubmitDocumentAnalyzeJobAdvanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitDocumentAnalyzeJobResponse, _err error) {
	// Step 0: init client
	if dara.IsNil(client.Credential) {
		_err = &openapi.ClientError{
			Code:    dara.String("InvalidCredentials"),
			Message: dara.String("Please set up the credentials correctly. If you are setting them through environment variables, please ensure that ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set correctly. See https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems for more details."),
		}
		return _result, _err
	}

	credentialModel, _err := client.Credential.GetCredential()
	if _err != nil {
		return _result, _err
	}

	accessKeyId := dara.StringValue(credentialModel.AccessKeyId)
	accessKeySecret := dara.StringValue(credentialModel.AccessKeySecret)
	securityToken := dara.StringValue(credentialModel.SecurityToken)
	credentialType := dara.StringValue(credentialModel.Type)
	openPlatformEndpoint := dara.StringValue(client.OpenPlatformEndpoint)
	if dara.IsNil(dara.String(openPlatformEndpoint)) || openPlatformEndpoint == "" {
		openPlatformEndpoint = "openplatform.aliyuncs.com"
	}

	if dara.IsNil(dara.String(credentialType)) {
		credentialType = "access_key"
	}

	authConfig := &openapiutil.Config{
		AccessKeyId:     dara.String(accessKeyId),
		AccessKeySecret: dara.String(accessKeySecret),
		SecurityToken:   dara.String(securityToken),
		Type:            dara.String(credentialType),
		Endpoint:        dara.String(openPlatformEndpoint),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openapi.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := map[string]*string{
		"Product":  dara.String("energyExpertExternal"),
		"RegionId": client.RegionId,
	}
	authReq := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(authRequest),
	}
	authParams := &openapiutil.Params{
		Action:      dara.String("AuthorizeFileUpload"),
		Version:     dara.String("2019-12-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	authResponse := map[string]interface{}{}
	fileObj := &dara.FileField{}
	ossHeader := map[string]interface{}{}
	tmpBody := map[string]interface{}{}
	useAccelerate := false
	authResponseBody := make(map[string]*string)
	submitDocumentAnalyzeJobReq := &SubmitDocumentAnalyzeJobRequest{}
	openapiutil.Convert(request, submitDocumentAnalyzeJobReq)
	if !dara.IsNil(request.FileUrlObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.FileUrlObject,
			ContentType: dara.String(""),
		}
		ossHeader = map[string]interface{}{
			"host":                  dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(openapiutil.GetEndpoint(authResponseBody["Endpoint"], dara.Bool(useAccelerate), client.EndpointType)),
			"OSSAccessKeyId":        dara.StringValue(authResponseBody["AccessKeyId"]),
			"policy":                dara.StringValue(authResponseBody["EncodedPolicy"]),
			"Signature":             dara.StringValue(authResponseBody["Signature"]),
			"key":                   dara.StringValue(authResponseBody["ObjectKey"]),
			"file":                  fileObj,
			"success_action_status": "201",
		}
		_, _err = client._postOSSObject(authResponseBody["Bucket"], ossHeader)
		if _err != nil {
			return _result, _err
		}
		submitDocumentAnalyzeJobReq.FileUrl = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	submitDocumentAnalyzeJobResp, _err := client.SubmitDocumentAnalyzeJobWithOptions(submitDocumentAnalyzeJobReq, headers, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = submitDocumentAnalyzeJobResp
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
func (client *Client) SubmitVLExtractionTaskWithOptions(request *SubmitVLExtractionTaskRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitVLExtractionTaskResponse, _err error) {
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
	_body, _err := client.CallApi(params, req, runtime)
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
// @return SubmitVLExtractionTaskResponse
func (client *Client) SubmitVLExtractionTask(request *SubmitVLExtractionTaskRequest) (_result *SubmitVLExtractionTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitVLExtractionTaskResponse{}
	_body, _err := client.SubmitVLExtractionTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitVLExtractionTaskAdvance(request *SubmitVLExtractionTaskAdvanceRequest, headers map[string]*string, runtime *dara.RuntimeOptions) (_result *SubmitVLExtractionTaskResponse, _err error) {
	// Step 0: init client
	if dara.IsNil(client.Credential) {
		_err = &openapi.ClientError{
			Code:    dara.String("InvalidCredentials"),
			Message: dara.String("Please set up the credentials correctly. If you are setting them through environment variables, please ensure that ALIBABA_CLOUD_ACCESS_KEY_ID and ALIBABA_CLOUD_ACCESS_KEY_SECRET are set correctly. See https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems for more details."),
		}
		return _result, _err
	}

	credentialModel, _err := client.Credential.GetCredential()
	if _err != nil {
		return _result, _err
	}

	accessKeyId := dara.StringValue(credentialModel.AccessKeyId)
	accessKeySecret := dara.StringValue(credentialModel.AccessKeySecret)
	securityToken := dara.StringValue(credentialModel.SecurityToken)
	credentialType := dara.StringValue(credentialModel.Type)
	openPlatformEndpoint := dara.StringValue(client.OpenPlatformEndpoint)
	if dara.IsNil(dara.String(openPlatformEndpoint)) || openPlatformEndpoint == "" {
		openPlatformEndpoint = "openplatform.aliyuncs.com"
	}

	if dara.IsNil(dara.String(credentialType)) {
		credentialType = "access_key"
	}

	authConfig := &openapiutil.Config{
		AccessKeyId:     dara.String(accessKeyId),
		AccessKeySecret: dara.String(accessKeySecret),
		SecurityToken:   dara.String(securityToken),
		Type:            dara.String(credentialType),
		Endpoint:        dara.String(openPlatformEndpoint),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openapi.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := map[string]*string{
		"Product":  dara.String("energyExpertExternal"),
		"RegionId": client.RegionId,
	}
	authReq := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(authRequest),
	}
	authParams := &openapiutil.Params{
		Action:      dara.String("AuthorizeFileUpload"),
		Version:     dara.String("2019-12-19"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("GET"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	authResponse := map[string]interface{}{}
	fileObj := &dara.FileField{}
	ossHeader := map[string]interface{}{}
	tmpBody := map[string]interface{}{}
	useAccelerate := false
	authResponseBody := make(map[string]*string)
	submitVLExtractionTaskReq := &SubmitVLExtractionTaskRequest{}
	openapiutil.Convert(request, submitVLExtractionTaskReq)
	if !dara.IsNil(request.FileUrlObject) {
		authResponse, _err = authClient.CallApi(authParams, authReq, runtime)
		if _err != nil {
			return _result, _err
		}

		tmpBody = dara.ToMap(authResponse["body"])
		useAccelerate = dara.ForceBoolean(tmpBody["UseAccelerate"])
		authResponseBody = openapiutil.StringifyMapValue(tmpBody)
		fileObj = &dara.FileField{
			Filename:    authResponseBody["ObjectKey"],
			Content:     request.FileUrlObject,
			ContentType: dara.String(""),
		}
		ossHeader = map[string]interface{}{
			"host":                  dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(openapiutil.GetEndpoint(authResponseBody["Endpoint"], dara.Bool(useAccelerate), client.EndpointType)),
			"OSSAccessKeyId":        dara.StringValue(authResponseBody["AccessKeyId"]),
			"policy":                dara.StringValue(authResponseBody["EncodedPolicy"]),
			"Signature":             dara.StringValue(authResponseBody["Signature"]),
			"key":                   dara.StringValue(authResponseBody["ObjectKey"]),
			"file":                  fileObj,
			"success_action_status": "201",
		}
		_, _err = client._postOSSObject(authResponseBody["Bucket"], ossHeader)
		if _err != nil {
			return _result, _err
		}
		submitVLExtractionTaskReq.FileUrl = dara.String("http://" + dara.StringValue(authResponseBody["Bucket"]) + "." + dara.StringValue(authResponseBody["Endpoint"]) + "/" + dara.StringValue(authResponseBody["ObjectKey"]))
	}

	submitVLExtractionTaskResp, _err := client.SubmitVLExtractionTaskWithOptions(submitVLExtractionTaskReq, headers, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = submitVLExtractionTaskResp
	return _result, _err
}

func _postOSSObject_opResponse(response_ *dara.Response) (_result map[string]interface{}, _err error) {
	var respMap map[string]interface{}
	bodyStr, _err := dara.ReadAsString(response_.Body)
	if _err != nil {
		return _result, _err
	}

	if (dara.IntValue(response_.StatusCode) >= 400) && (dara.IntValue(response_.StatusCode) < 600) {
		respMap = dara.ParseXml(bodyStr, nil)
		err := dara.ToMap(respMap["Error"])
		_err = &openapi.ClientError{
			Code:    dara.String(dara.ToString(err["Code"])),
			Message: dara.String(dara.ToString(err["Message"])),
			Data: map[string]interface{}{
				"httpCode":  dara.IntValue(response_.StatusCode),
				"requestId": dara.ToString(err["RequestId"]),
				"hostId":    dara.ToString(err["HostId"]),
			},
		}
		return _result, _err
	}

	respMap = dara.ParseXml(bodyStr, nil)
	_result = make(map[string]interface{})
	_err = dara.Convert(dara.ToMap(respMap), &_result)

	return _result, _err
}

func (client *Client) chatStreamWithSSE_opYieldFunc(_yield chan *ChatStreamResponse, _yieldErr chan error, request *ChatStreamRequest, headers map[string]*string, runtime *dara.RuntimeOptions) {
	if dara.BoolValue(client.EnableValidate) == true {
		_err := request.Validate()
		if _err != nil {
			_yieldErr <- _err
			return
		}
	}
	body := map[string]interface{}{}
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
