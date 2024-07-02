// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type RunCompletionRequest struct {
	// This parameter is required.
	Dialogue *RunCompletionRequestDialogue `json:"Dialogue,omitempty" xml:"Dialogue,omitempty" type:"Struct"`
	Fields   []*RunCompletionRequestFields `json:"Fields,omitempty" xml:"Fields,omitempty" type:"Repeated"`
	// example:
	//
	// ccai-14b
	ModelCode         *string                                `json:"ModelCode,omitempty" xml:"ModelCode,omitempty"`
	ServiceInspection *RunCompletionRequestServiceInspection `json:"ServiceInspection,omitempty" xml:"ServiceInspection,omitempty" type:"Struct"`
	// example:
	//
	// false
	Stream *bool `json:"Stream,omitempty" xml:"Stream,omitempty"`
	// This parameter is required.
	TemplateIds []*int64 `json:"TemplateIds,omitempty" xml:"TemplateIds,omitempty" type:"Repeated"`
}

func (s RunCompletionRequest) String() string {
	return tea.Prettify(s)
}

func (s RunCompletionRequest) GoString() string {
	return s.String()
}

func (s *RunCompletionRequest) SetDialogue(v *RunCompletionRequestDialogue) *RunCompletionRequest {
	s.Dialogue = v
	return s
}

func (s *RunCompletionRequest) SetFields(v []*RunCompletionRequestFields) *RunCompletionRequest {
	s.Fields = v
	return s
}

func (s *RunCompletionRequest) SetModelCode(v string) *RunCompletionRequest {
	s.ModelCode = &v
	return s
}

func (s *RunCompletionRequest) SetServiceInspection(v *RunCompletionRequestServiceInspection) *RunCompletionRequest {
	s.ServiceInspection = v
	return s
}

func (s *RunCompletionRequest) SetStream(v bool) *RunCompletionRequest {
	s.Stream = &v
	return s
}

func (s *RunCompletionRequest) SetTemplateIds(v []*int64) *RunCompletionRequest {
	s.TemplateIds = v
	return s
}

type RunCompletionRequestDialogue struct {
	Sentences []*RunCompletionRequestDialogueSentences `json:"Sentences,omitempty" xml:"Sentences,omitempty" type:"Repeated"`
	// example:
	//
	// d25zc9c7004f8dad2b454d
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s RunCompletionRequestDialogue) String() string {
	return tea.Prettify(s)
}

func (s RunCompletionRequestDialogue) GoString() string {
	return s.String()
}

func (s *RunCompletionRequestDialogue) SetSentences(v []*RunCompletionRequestDialogueSentences) *RunCompletionRequestDialogue {
	s.Sentences = v
	return s
}

func (s *RunCompletionRequestDialogue) SetSessionId(v string) *RunCompletionRequestDialogue {
	s.SessionId = &v
	return s
}

type RunCompletionRequestDialogueSentences struct {
	ChatId *string `json:"ChatId,omitempty" xml:"ChatId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// This parameter is required.
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RunCompletionRequestDialogueSentences) String() string {
	return tea.Prettify(s)
}

func (s RunCompletionRequestDialogueSentences) GoString() string {
	return s.String()
}

func (s *RunCompletionRequestDialogueSentences) SetChatId(v string) *RunCompletionRequestDialogueSentences {
	s.ChatId = &v
	return s
}

func (s *RunCompletionRequestDialogueSentences) SetRole(v string) *RunCompletionRequestDialogueSentences {
	s.Role = &v
	return s
}

func (s *RunCompletionRequestDialogueSentences) SetText(v string) *RunCompletionRequestDialogueSentences {
	s.Text = &v
	return s
}

type RunCompletionRequestFields struct {
	// example:
	//
	// phoneNumber
	Code       *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Desc       *string                                 `json:"Desc,omitempty" xml:"Desc,omitempty"`
	EnumValues []*RunCompletionRequestFieldsEnumValues `json:"EnumValues,omitempty" xml:"EnumValues,omitempty" type:"Repeated"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s RunCompletionRequestFields) String() string {
	return tea.Prettify(s)
}

func (s RunCompletionRequestFields) GoString() string {
	return s.String()
}

func (s *RunCompletionRequestFields) SetCode(v string) *RunCompletionRequestFields {
	s.Code = &v
	return s
}

func (s *RunCompletionRequestFields) SetDesc(v string) *RunCompletionRequestFields {
	s.Desc = &v
	return s
}

func (s *RunCompletionRequestFields) SetEnumValues(v []*RunCompletionRequestFieldsEnumValues) *RunCompletionRequestFields {
	s.EnumValues = v
	return s
}

func (s *RunCompletionRequestFields) SetName(v string) *RunCompletionRequestFields {
	s.Name = &v
	return s
}

type RunCompletionRequestFieldsEnumValues struct {
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// This parameter is required.
	EnumValue *string `json:"EnumValue,omitempty" xml:"EnumValue,omitempty"`
}

func (s RunCompletionRequestFieldsEnumValues) String() string {
	return tea.Prettify(s)
}

func (s RunCompletionRequestFieldsEnumValues) GoString() string {
	return s.String()
}

func (s *RunCompletionRequestFieldsEnumValues) SetDesc(v string) *RunCompletionRequestFieldsEnumValues {
	s.Desc = &v
	return s
}

func (s *RunCompletionRequestFieldsEnumValues) SetEnumValue(v string) *RunCompletionRequestFieldsEnumValues {
	s.EnumValue = &v
	return s
}

type RunCompletionRequestServiceInspection struct {
	InspectionContents     []*RunCompletionRequestServiceInspectionInspectionContents `json:"InspectionContents,omitempty" xml:"InspectionContents,omitempty" type:"Repeated"`
	InspectionIntroduction *string                                                    `json:"InspectionIntroduction,omitempty" xml:"InspectionIntroduction,omitempty"`
	SceneIntroduction      *string                                                    `json:"SceneIntroduction,omitempty" xml:"SceneIntroduction,omitempty"`
}

func (s RunCompletionRequestServiceInspection) String() string {
	return tea.Prettify(s)
}

func (s RunCompletionRequestServiceInspection) GoString() string {
	return s.String()
}

func (s *RunCompletionRequestServiceInspection) SetInspectionContents(v []*RunCompletionRequestServiceInspectionInspectionContents) *RunCompletionRequestServiceInspection {
	s.InspectionContents = v
	return s
}

func (s *RunCompletionRequestServiceInspection) SetInspectionIntroduction(v string) *RunCompletionRequestServiceInspection {
	s.InspectionIntroduction = &v
	return s
}

func (s *RunCompletionRequestServiceInspection) SetSceneIntroduction(v string) *RunCompletionRequestServiceInspection {
	s.SceneIntroduction = &v
	return s
}

type RunCompletionRequestServiceInspectionInspectionContents struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s RunCompletionRequestServiceInspectionInspectionContents) String() string {
	return tea.Prettify(s)
}

func (s RunCompletionRequestServiceInspectionInspectionContents) GoString() string {
	return s.String()
}

func (s *RunCompletionRequestServiceInspectionInspectionContents) SetContent(v string) *RunCompletionRequestServiceInspectionInspectionContents {
	s.Content = &v
	return s
}

func (s *RunCompletionRequestServiceInspectionInspectionContents) SetTitle(v string) *RunCompletionRequestServiceInspectionInspectionContents {
	s.Title = &v
	return s
}

type RunCompletionResponseBody struct {
	FinishReason *string `json:"FinishReason,omitempty" xml:"FinishReason,omitempty"`
	// example:
	//
	// 17204B98-xxxx-4F9A-8464-2446A84821CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Text      *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RunCompletionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunCompletionResponseBody) GoString() string {
	return s.String()
}

func (s *RunCompletionResponseBody) SetFinishReason(v string) *RunCompletionResponseBody {
	s.FinishReason = &v
	return s
}

func (s *RunCompletionResponseBody) SetRequestId(v string) *RunCompletionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunCompletionResponseBody) SetText(v string) *RunCompletionResponseBody {
	s.Text = &v
	return s
}

type RunCompletionResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunCompletionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunCompletionResponse) String() string {
	return tea.Prettify(s)
}

func (s RunCompletionResponse) GoString() string {
	return s.String()
}

func (s *RunCompletionResponse) SetHeaders(v map[string]*string) *RunCompletionResponse {
	s.Headers = v
	return s
}

func (s *RunCompletionResponse) SetStatusCode(v int32) *RunCompletionResponse {
	s.StatusCode = &v
	return s
}

func (s *RunCompletionResponse) SetBody(v *RunCompletionResponseBody) *RunCompletionResponse {
	s.Body = v
	return s
}

type RunCompletionMessageRequest struct {
	// This parameter is required.
	Messages []*RunCompletionMessageRequestMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// example:
	//
	// ccai-14b
	ModelCode *string `json:"ModelCode,omitempty" xml:"ModelCode,omitempty"`
	// example:
	//
	// false
	Stream *bool `json:"Stream,omitempty" xml:"Stream,omitempty"`
}

func (s RunCompletionMessageRequest) String() string {
	return tea.Prettify(s)
}

func (s RunCompletionMessageRequest) GoString() string {
	return s.String()
}

func (s *RunCompletionMessageRequest) SetMessages(v []*RunCompletionMessageRequestMessages) *RunCompletionMessageRequest {
	s.Messages = v
	return s
}

func (s *RunCompletionMessageRequest) SetModelCode(v string) *RunCompletionMessageRequest {
	s.ModelCode = &v
	return s
}

func (s *RunCompletionMessageRequest) SetStream(v bool) *RunCompletionMessageRequest {
	s.Stream = &v
	return s
}

type RunCompletionMessageRequestMessages struct {
	// This parameter is required.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s RunCompletionMessageRequestMessages) String() string {
	return tea.Prettify(s)
}

func (s RunCompletionMessageRequestMessages) GoString() string {
	return s.String()
}

func (s *RunCompletionMessageRequestMessages) SetContent(v string) *RunCompletionMessageRequestMessages {
	s.Content = &v
	return s
}

func (s *RunCompletionMessageRequestMessages) SetRole(v string) *RunCompletionMessageRequestMessages {
	s.Role = &v
	return s
}

type RunCompletionMessageResponseBody struct {
	FinishReason *string `json:"FinishReason,omitempty" xml:"FinishReason,omitempty"`
	// example:
	//
	// 17204B98-xxxx-4F9A-8464-2446A84821CA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Text      *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RunCompletionMessageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunCompletionMessageResponseBody) GoString() string {
	return s.String()
}

func (s *RunCompletionMessageResponseBody) SetFinishReason(v string) *RunCompletionMessageResponseBody {
	s.FinishReason = &v
	return s
}

func (s *RunCompletionMessageResponseBody) SetRequestId(v string) *RunCompletionMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunCompletionMessageResponseBody) SetText(v string) *RunCompletionMessageResponseBody {
	s.Text = &v
	return s
}

type RunCompletionMessageResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunCompletionMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunCompletionMessageResponse) String() string {
	return tea.Prettify(s)
}

func (s RunCompletionMessageResponse) GoString() string {
	return s.String()
}

func (s *RunCompletionMessageResponse) SetHeaders(v map[string]*string) *RunCompletionMessageResponse {
	s.Headers = v
	return s
}

func (s *RunCompletionMessageResponse) SetStatusCode(v int32) *RunCompletionMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *RunCompletionMessageResponse) SetBody(v *RunCompletionMessageResponseBody) *RunCompletionMessageResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("contactcenterai"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// CCAI服务面API
//
// @param request - RunCompletionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunCompletionResponse
func (client *Client) RunCompletionWithOptions(workspaceId *string, appId *string, request *RunCompletionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RunCompletionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Dialogue)) {
		body["Dialogue"] = request.Dialogue
	}

	if !tea.BoolValue(util.IsUnset(request.Fields)) {
		body["Fields"] = request.Fields
	}

	if !tea.BoolValue(util.IsUnset(request.ModelCode)) {
		body["ModelCode"] = request.ModelCode
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInspection)) {
		body["ServiceInspection"] = request.ServiceInspection
	}

	if !tea.BoolValue(util.IsUnset(request.Stream)) {
		body["Stream"] = request.Stream
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateIds)) {
		body["TemplateIds"] = request.TemplateIds
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunCompletion"),
		Version:     tea.String("2024-06-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/ccai/app/" + tea.StringValue(openapiutil.GetEncodeParam(appId)) + "/completion"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RunCompletionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// CCAI服务面API
//
// @param request - RunCompletionRequest
//
// @return RunCompletionResponse
func (client *Client) RunCompletion(workspaceId *string, appId *string, request *RunCompletionRequest) (_result *RunCompletionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
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
// CCAI服务面API
//
// @param request - RunCompletionMessageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunCompletionMessageResponse
func (client *Client) RunCompletionMessageWithOptions(workspaceId *string, appId *string, request *RunCompletionMessageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RunCompletionMessageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Messages)) {
		body["Messages"] = request.Messages
	}

	if !tea.BoolValue(util.IsUnset(request.ModelCode)) {
		body["ModelCode"] = request.ModelCode
	}

	if !tea.BoolValue(util.IsUnset(request.Stream)) {
		body["Stream"] = request.Stream
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunCompletionMessage"),
		Version:     tea.String("2024-06-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/ccai/app/" + tea.StringValue(openapiutil.GetEncodeParam(appId)) + "/completion_message"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RunCompletionMessageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// CCAI服务面API
//
// @param request - RunCompletionMessageRequest
//
// @return RunCompletionMessageResponse
func (client *Client) RunCompletionMessage(workspaceId *string, appId *string, request *RunCompletionMessageRequest) (_result *RunCompletionMessageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunCompletionMessageResponse{}
	_body, _err := client.RunCompletionMessageWithOptions(workspaceId, appId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
