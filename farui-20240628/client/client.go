// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	openplatform "github.com/alibabacloud-go/openplatform-20191219/v2/client"
	fileform "github.com/alibabacloud-go/tea-fileform/service"
	oss "github.com/alibabacloud-go/tea-oss-sdk/client"
	ossutil "github.com/alibabacloud-go/tea-oss-utils/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type CreateTextFileRequest struct {
	// example:
	//
	// e9a93201-7e96-4dc1-9678-2832fc132d08
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// 1714476549
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	TextFileName *string `json:"TextFileName,omitempty" xml:"TextFileName,omitempty"`
	TextFileUrl  *string `json:"TextFileUrl,omitempty" xml:"TextFileUrl,omitempty"`
}

func (s CreateTextFileRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTextFileRequest) GoString() string {
	return s.String()
}

func (s *CreateTextFileRequest) SetClientToken(v string) *CreateTextFileRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTextFileRequest) SetCreateTime(v string) *CreateTextFileRequest {
	s.CreateTime = &v
	return s
}

func (s *CreateTextFileRequest) SetTextFileName(v string) *CreateTextFileRequest {
	s.TextFileName = &v
	return s
}

func (s *CreateTextFileRequest) SetTextFileUrl(v string) *CreateTextFileRequest {
	s.TextFileUrl = &v
	return s
}

type CreateTextFileAdvanceRequest struct {
	// example:
	//
	// e9a93201-7e96-4dc1-9678-2832fc132d08
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// 1714476549
	CreateTime        *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	TextFileName      *string   `json:"TextFileName,omitempty" xml:"TextFileName,omitempty"`
	TextFileUrlObject io.Reader `json:"TextFileUrl,omitempty" xml:"TextFileUrl,omitempty"`
}

func (s CreateTextFileAdvanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTextFileAdvanceRequest) GoString() string {
	return s.String()
}

func (s *CreateTextFileAdvanceRequest) SetClientToken(v string) *CreateTextFileAdvanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTextFileAdvanceRequest) SetCreateTime(v string) *CreateTextFileAdvanceRequest {
	s.CreateTime = &v
	return s
}

func (s *CreateTextFileAdvanceRequest) SetTextFileName(v string) *CreateTextFileAdvanceRequest {
	s.TextFileName = &v
	return s
}

func (s *CreateTextFileAdvanceRequest) SetTextFileUrlObject(v io.Reader) *CreateTextFileAdvanceRequest {
	s.TextFileUrlObject = v
	return s
}

type CreateTextFileResponseBody struct {
	// example:
	//
	// Request.Signature.Error
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateTextFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int64  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 81E6F6D2-8ACB-5BDA-9C7C-4D6268CD9652
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTextFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTextFileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTextFileResponseBody) SetCode(v string) *CreateTextFileResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTextFileResponseBody) SetData(v *CreateTextFileResponseBodyData) *CreateTextFileResponseBody {
	s.Data = v
	return s
}

func (s *CreateTextFileResponseBody) SetHttpStatusCode(v int64) *CreateTextFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateTextFileResponseBody) SetMessage(v string) *CreateTextFileResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTextFileResponseBody) SetRequestId(v string) *CreateTextFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTextFileResponseBody) SetSuccess(v bool) *CreateTextFileResponseBody {
	s.Success = &v
	return s
}

type CreateTextFileResponseBodyData struct {
	// example:
	//
	// 36d6447d277c4a1c9fd0def1d16341f1
	TextFileId   *string `json:"TextFileId,omitempty" xml:"TextFileId,omitempty"`
	TextFileName *string `json:"TextFileName,omitempty" xml:"TextFileName,omitempty"`
	TextFileUrl  *string `json:"TextFileUrl,omitempty" xml:"TextFileUrl,omitempty"`
}

func (s CreateTextFileResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateTextFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateTextFileResponseBodyData) SetTextFileId(v string) *CreateTextFileResponseBodyData {
	s.TextFileId = &v
	return s
}

func (s *CreateTextFileResponseBodyData) SetTextFileName(v string) *CreateTextFileResponseBodyData {
	s.TextFileName = &v
	return s
}

func (s *CreateTextFileResponseBodyData) SetTextFileUrl(v string) *CreateTextFileResponseBodyData {
	s.TextFileUrl = &v
	return s
}

type CreateTextFileResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTextFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTextFileResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTextFileResponse) GoString() string {
	return s.String()
}

func (s *CreateTextFileResponse) SetHeaders(v map[string]*string) *CreateTextFileResponse {
	s.Headers = v
	return s
}

func (s *CreateTextFileResponse) SetStatusCode(v int32) *CreateTextFileResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTextFileResponse) SetBody(v *CreateTextFileResponseBody) *CreateTextFileResponse {
	s.Body = v
	return s
}

type RunLegalAdviceConsultationRequest struct {
	// example:
	//
	// farui
	AppId     *string                                     `json:"appId,omitempty" xml:"appId,omitempty"`
	Assistant *RunLegalAdviceConsultationRequestAssistant `json:"assistant,omitempty" xml:"assistant,omitempty" type:"Struct"`
	// example:
	//
	// true
	Stream *bool                                    `json:"stream,omitempty" xml:"stream,omitempty"`
	Thread *RunLegalAdviceConsultationRequestThread `json:"thread,omitempty" xml:"thread,omitempty" type:"Struct"`
}

func (s RunLegalAdviceConsultationRequest) String() string {
	return tea.Prettify(s)
}

func (s RunLegalAdviceConsultationRequest) GoString() string {
	return s.String()
}

func (s *RunLegalAdviceConsultationRequest) SetAppId(v string) *RunLegalAdviceConsultationRequest {
	s.AppId = &v
	return s
}

func (s *RunLegalAdviceConsultationRequest) SetAssistant(v *RunLegalAdviceConsultationRequestAssistant) *RunLegalAdviceConsultationRequest {
	s.Assistant = v
	return s
}

func (s *RunLegalAdviceConsultationRequest) SetStream(v bool) *RunLegalAdviceConsultationRequest {
	s.Stream = &v
	return s
}

func (s *RunLegalAdviceConsultationRequest) SetThread(v *RunLegalAdviceConsultationRequestThread) *RunLegalAdviceConsultationRequest {
	s.Thread = v
	return s
}

type RunLegalAdviceConsultationRequestAssistant struct {
	// example:
	//
	// assitant_abc_123
	Id       *string            `json:"id,omitempty" xml:"id,omitempty"`
	MetaData map[string]*string `json:"metaData,omitempty" xml:"metaData,omitempty"`
	// example:
	//
	// legal_advice_consult
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s RunLegalAdviceConsultationRequestAssistant) String() string {
	return tea.Prettify(s)
}

func (s RunLegalAdviceConsultationRequestAssistant) GoString() string {
	return s.String()
}

func (s *RunLegalAdviceConsultationRequestAssistant) SetId(v string) *RunLegalAdviceConsultationRequestAssistant {
	s.Id = &v
	return s
}

func (s *RunLegalAdviceConsultationRequestAssistant) SetMetaData(v map[string]*string) *RunLegalAdviceConsultationRequestAssistant {
	s.MetaData = v
	return s
}

func (s *RunLegalAdviceConsultationRequestAssistant) SetType(v string) *RunLegalAdviceConsultationRequestAssistant {
	s.Type = &v
	return s
}

func (s *RunLegalAdviceConsultationRequestAssistant) SetVersion(v string) *RunLegalAdviceConsultationRequestAssistant {
	s.Version = &v
	return s
}

type RunLegalAdviceConsultationRequestThread struct {
	Messages []*RunLegalAdviceConsultationRequestThreadMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
}

func (s RunLegalAdviceConsultationRequestThread) String() string {
	return tea.Prettify(s)
}

func (s RunLegalAdviceConsultationRequestThread) GoString() string {
	return s.String()
}

func (s *RunLegalAdviceConsultationRequestThread) SetMessages(v []*RunLegalAdviceConsultationRequestThreadMessages) *RunLegalAdviceConsultationRequestThread {
	s.Messages = v
	return s
}

type RunLegalAdviceConsultationRequestThreadMessages struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s RunLegalAdviceConsultationRequestThreadMessages) String() string {
	return tea.Prettify(s)
}

func (s RunLegalAdviceConsultationRequestThreadMessages) GoString() string {
	return s.String()
}

func (s *RunLegalAdviceConsultationRequestThreadMessages) SetContent(v string) *RunLegalAdviceConsultationRequestThreadMessages {
	s.Content = &v
	return s
}

func (s *RunLegalAdviceConsultationRequestThreadMessages) SetRole(v string) *RunLegalAdviceConsultationRequestThreadMessages {
	s.Role = &v
	return s
}

type RunLegalAdviceConsultationShrinkRequest struct {
	// example:
	//
	// farui
	AppId           *string `json:"appId,omitempty" xml:"appId,omitempty"`
	AssistantShrink *string `json:"assistant,omitempty" xml:"assistant,omitempty"`
	// example:
	//
	// true
	Stream       *bool   `json:"stream,omitempty" xml:"stream,omitempty"`
	ThreadShrink *string `json:"thread,omitempty" xml:"thread,omitempty"`
}

func (s RunLegalAdviceConsultationShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RunLegalAdviceConsultationShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunLegalAdviceConsultationShrinkRequest) SetAppId(v string) *RunLegalAdviceConsultationShrinkRequest {
	s.AppId = &v
	return s
}

func (s *RunLegalAdviceConsultationShrinkRequest) SetAssistantShrink(v string) *RunLegalAdviceConsultationShrinkRequest {
	s.AssistantShrink = &v
	return s
}

func (s *RunLegalAdviceConsultationShrinkRequest) SetStream(v bool) *RunLegalAdviceConsultationShrinkRequest {
	s.Stream = &v
	return s
}

func (s *RunLegalAdviceConsultationShrinkRequest) SetThreadShrink(v string) *RunLegalAdviceConsultationShrinkRequest {
	s.ThreadShrink = &v
	return s
}

type RunLegalAdviceConsultationResponseBody struct {
	// example:
	//
	// Request.Signature.Error
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 744419D0-671A-5997-9840-E8AE48356194
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResponseMarkdown *string `json:"ResponseMarkdown,omitempty" xml:"ResponseMarkdown,omitempty"`
	// example:
	//
	// 1
	Round  *int32  `json:"Round,omitempty" xml:"Round,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// True
	Success *bool                                        `json:"Success,omitempty" xml:"Success,omitempty"`
	Usage   *RunLegalAdviceConsultationResponseBodyUsage `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
}

func (s RunLegalAdviceConsultationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunLegalAdviceConsultationResponseBody) GoString() string {
	return s.String()
}

func (s *RunLegalAdviceConsultationResponseBody) SetCode(v string) *RunLegalAdviceConsultationResponseBody {
	s.Code = &v
	return s
}

func (s *RunLegalAdviceConsultationResponseBody) SetMessage(v string) *RunLegalAdviceConsultationResponseBody {
	s.Message = &v
	return s
}

func (s *RunLegalAdviceConsultationResponseBody) SetRequestId(v string) *RunLegalAdviceConsultationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunLegalAdviceConsultationResponseBody) SetResponseMarkdown(v string) *RunLegalAdviceConsultationResponseBody {
	s.ResponseMarkdown = &v
	return s
}

func (s *RunLegalAdviceConsultationResponseBody) SetRound(v int32) *RunLegalAdviceConsultationResponseBody {
	s.Round = &v
	return s
}

func (s *RunLegalAdviceConsultationResponseBody) SetStatus(v string) *RunLegalAdviceConsultationResponseBody {
	s.Status = &v
	return s
}

func (s *RunLegalAdviceConsultationResponseBody) SetSuccess(v bool) *RunLegalAdviceConsultationResponseBody {
	s.Success = &v
	return s
}

func (s *RunLegalAdviceConsultationResponseBody) SetUsage(v *RunLegalAdviceConsultationResponseBodyUsage) *RunLegalAdviceConsultationResponseBody {
	s.Usage = v
	return s
}

func (s *RunLegalAdviceConsultationResponseBody) SetHttpStatusCode(v string) *RunLegalAdviceConsultationResponseBody {
	s.HttpStatusCode = &v
	return s
}

type RunLegalAdviceConsultationResponseBodyUsage struct {
	// example:
	//
	// 500
	InputTokens *int32 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// example:
	//
	// 700
	OutputTokens *int32 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// example:
	//
	// 1200
	TotalTokens *int32 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s RunLegalAdviceConsultationResponseBodyUsage) String() string {
	return tea.Prettify(s)
}

func (s RunLegalAdviceConsultationResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *RunLegalAdviceConsultationResponseBodyUsage) SetInputTokens(v int32) *RunLegalAdviceConsultationResponseBodyUsage {
	s.InputTokens = &v
	return s
}

func (s *RunLegalAdviceConsultationResponseBodyUsage) SetOutputTokens(v int32) *RunLegalAdviceConsultationResponseBodyUsage {
	s.OutputTokens = &v
	return s
}

func (s *RunLegalAdviceConsultationResponseBodyUsage) SetTotalTokens(v int32) *RunLegalAdviceConsultationResponseBodyUsage {
	s.TotalTokens = &v
	return s
}

type RunLegalAdviceConsultationResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunLegalAdviceConsultationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunLegalAdviceConsultationResponse) String() string {
	return tea.Prettify(s)
}

func (s RunLegalAdviceConsultationResponse) GoString() string {
	return s.String()
}

func (s *RunLegalAdviceConsultationResponse) SetHeaders(v map[string]*string) *RunLegalAdviceConsultationResponse {
	s.Headers = v
	return s
}

func (s *RunLegalAdviceConsultationResponse) SetStatusCode(v int32) *RunLegalAdviceConsultationResponse {
	s.StatusCode = &v
	return s
}

func (s *RunLegalAdviceConsultationResponse) SetBody(v *RunLegalAdviceConsultationResponseBody) *RunLegalAdviceConsultationResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("farui"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 上传合同文件
//
// @param request - CreateTextFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTextFileResponse
func (client *Client) CreateTextFileWithOptions(WorkspaceId *string, request *CreateTextFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateTextFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTime)) {
		body["CreateTime"] = request.CreateTime
	}

	if !tea.BoolValue(util.IsUnset(request.TextFileName)) {
		body["TextFileName"] = request.TextFileName
	}

	if !tea.BoolValue(util.IsUnset(request.TextFileUrl)) {
		body["TextFileUrl"] = request.TextFileUrl
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTextFile"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(WorkspaceId)) + "/data/textFile"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTextFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 上传合同文件
//
// @param request - CreateTextFileRequest
//
// @return CreateTextFileResponse
func (client *Client) CreateTextFile(WorkspaceId *string, request *CreateTextFileRequest) (_result *CreateTextFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTextFileResponse{}
	_body, _err := client.CreateTextFileWithOptions(WorkspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTextFileAdvance(WorkspaceId *string, request *CreateTextFileAdvanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateTextFileResponse, _err error) {
	// Step 0: init client
	accessKeyId, _err := client.Credential.GetAccessKeyId()
	if _err != nil {
		return _result, _err
	}

	accessKeySecret, _err := client.Credential.GetAccessKeySecret()
	if _err != nil {
		return _result, _err
	}

	securityToken, _err := client.Credential.GetSecurityToken()
	if _err != nil {
		return _result, _err
	}

	credentialType := client.Credential.GetType()
	openPlatformEndpoint := client.OpenPlatformEndpoint
	if tea.BoolValue(util.Empty(openPlatformEndpoint)) {
		openPlatformEndpoint = tea.String("openplatform.aliyuncs.com")
	}

	if tea.BoolValue(util.IsUnset(credentialType)) {
		credentialType = tea.String("access_key")
	}

	authConfig := &openapi.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		SecurityToken:   securityToken,
		Type:            credentialType,
		Endpoint:        openPlatformEndpoint,
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	authClient, _err := openplatform.NewClient(authConfig)
	if _err != nil {
		return _result, _err
	}

	authRequest := &openplatform.AuthorizeFileUploadRequest{
		Product:  tea.String("FaRui"),
		RegionId: client.RegionId,
	}
	authResponse := &openplatform.AuthorizeFileUploadResponse{}
	ossConfig := &oss.Config{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		Type:            tea.String("access_key"),
		Protocol:        client.Protocol,
		RegionId:        client.RegionId,
	}
	ossClient, _err := oss.NewClient(ossConfig)
	if _err != nil {
		return _result, _err
	}

	fileObj := &fileform.FileField{}
	ossHeader := &oss.PostObjectRequestHeader{}
	uploadRequest := &oss.PostObjectRequest{}
	ossRuntime := &ossutil.RuntimeOptions{}
	openapiutil.Convert(runtime, ossRuntime)
	createTextFileReq := &CreateTextFileRequest{}
	openapiutil.Convert(request, createTextFileReq)
	if !tea.BoolValue(util.IsUnset(request.TextFileUrlObject)) {
		authResponse, _err = authClient.AuthorizeFileUploadWithOptions(authRequest, runtime)
		if _err != nil {
			return _result, _err
		}

		ossConfig.AccessKeyId = authResponse.Body.AccessKeyId
		ossConfig.Endpoint = openapiutil.GetEndpoint(authResponse.Body.Endpoint, authResponse.Body.UseAccelerate, client.EndpointType)
		ossClient, _err = oss.NewClient(ossConfig)
		if _err != nil {
			return _result, _err
		}

		fileObj = &fileform.FileField{
			Filename:    authResponse.Body.ObjectKey,
			Content:     request.TextFileUrlObject,
			ContentType: tea.String(""),
		}
		ossHeader = &oss.PostObjectRequestHeader{
			AccessKeyId:         authResponse.Body.AccessKeyId,
			Policy:              authResponse.Body.EncodedPolicy,
			Signature:           authResponse.Body.Signature,
			Key:                 authResponse.Body.ObjectKey,
			File:                fileObj,
			SuccessActionStatus: tea.String("201"),
		}
		uploadRequest = &oss.PostObjectRequest{
			BucketName: authResponse.Body.Bucket,
			Header:     ossHeader,
		}
		_, _err = ossClient.PostObject(uploadRequest, ossRuntime)
		if _err != nil {
			return _result, _err
		}
		createTextFileReq.TextFileUrl = tea.String("http://" + tea.StringValue(authResponse.Body.Bucket) + "." + tea.StringValue(authResponse.Body.Endpoint) + "/" + tea.StringValue(authResponse.Body.ObjectKey))
	}

	createTextFileResp, _err := client.CreateTextFileWithOptions(WorkspaceId, createTextFileReq, headers, runtime)
	if _err != nil {
		return _result, _err
	}

	_result = createTextFileResp
	return _result, _err
}

// Summary:
//
// 法律咨询
//
// @param tmpReq - RunLegalAdviceConsultationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunLegalAdviceConsultationResponse
func (client *Client) RunLegalAdviceConsultationWithOptions(workspaceId *string, tmpReq *RunLegalAdviceConsultationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RunLegalAdviceConsultationResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RunLegalAdviceConsultationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Assistant)) {
		request.AssistantShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Assistant, tea.String("assistant"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Thread)) {
		request.ThreadShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Thread, tea.String("thread"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		body["appId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AssistantShrink)) {
		body["assistant"] = request.AssistantShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Stream)) {
		body["stream"] = request.Stream
	}

	if !tea.BoolValue(util.IsUnset(request.ThreadShrink)) {
		body["thread"] = request.ThreadShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunLegalAdviceConsultation"),
		Version:     tea.String("2024-06-28"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/farui/legalAdvice/consult"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunLegalAdviceConsultationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 法律咨询
//
// @param request - RunLegalAdviceConsultationRequest
//
// @return RunLegalAdviceConsultationResponse
func (client *Client) RunLegalAdviceConsultation(workspaceId *string, request *RunLegalAdviceConsultationRequest) (_result *RunLegalAdviceConsultationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunLegalAdviceConsultationResponse{}
	_body, _err := client.RunLegalAdviceConsultationWithOptions(workspaceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
