// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AnalyzeConversationRequest struct {
	// This parameter is required.
	Dialogue *AnalyzeConversationRequestDialogue   `json:"dialogue,omitempty" xml:"dialogue,omitempty" type:"Struct"`
	Examples []*AnalyzeConversationRequestExamples `json:"examples,omitempty" xml:"examples,omitempty" type:"Repeated"`
	Fields   []*AnalyzeConversationRequestFields   `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	// example:
	//
	// tyxmTurbo
	ModelCode *string `json:"modelCode,omitempty" xml:"modelCode,omitempty"`
	// This parameter is required.
	ResultTypes       []*string                                    `json:"resultTypes,omitempty" xml:"resultTypes,omitempty" type:"Repeated"`
	SceneName         *string                                      `json:"sceneName,omitempty" xml:"sceneName,omitempty"`
	ServiceInspection *AnalyzeConversationRequestServiceInspection `json:"serviceInspection,omitempty" xml:"serviceInspection,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// false
	Stream *bool `json:"stream,omitempty" xml:"stream,omitempty"`
}

func (s AnalyzeConversationRequest) String() string {
	return tea.Prettify(s)
}

func (s AnalyzeConversationRequest) GoString() string {
	return s.String()
}

func (s *AnalyzeConversationRequest) SetDialogue(v *AnalyzeConversationRequestDialogue) *AnalyzeConversationRequest {
	s.Dialogue = v
	return s
}

func (s *AnalyzeConversationRequest) SetExamples(v []*AnalyzeConversationRequestExamples) *AnalyzeConversationRequest {
	s.Examples = v
	return s
}

func (s *AnalyzeConversationRequest) SetFields(v []*AnalyzeConversationRequestFields) *AnalyzeConversationRequest {
	s.Fields = v
	return s
}

func (s *AnalyzeConversationRequest) SetModelCode(v string) *AnalyzeConversationRequest {
	s.ModelCode = &v
	return s
}

func (s *AnalyzeConversationRequest) SetResultTypes(v []*string) *AnalyzeConversationRequest {
	s.ResultTypes = v
	return s
}

func (s *AnalyzeConversationRequest) SetSceneName(v string) *AnalyzeConversationRequest {
	s.SceneName = &v
	return s
}

func (s *AnalyzeConversationRequest) SetServiceInspection(v *AnalyzeConversationRequestServiceInspection) *AnalyzeConversationRequest {
	s.ServiceInspection = v
	return s
}

func (s *AnalyzeConversationRequest) SetStream(v bool) *AnalyzeConversationRequest {
	s.Stream = &v
	return s
}

type AnalyzeConversationRequestDialogue struct {
	// This parameter is required.
	Sentences []*AnalyzeConversationRequestDialogueSentences `json:"sentences,omitempty" xml:"sentences,omitempty" type:"Repeated"`
	// example:
	//
	// session-01
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s AnalyzeConversationRequestDialogue) String() string {
	return tea.Prettify(s)
}

func (s AnalyzeConversationRequestDialogue) GoString() string {
	return s.String()
}

func (s *AnalyzeConversationRequestDialogue) SetSentences(v []*AnalyzeConversationRequestDialogueSentences) *AnalyzeConversationRequestDialogue {
	s.Sentences = v
	return s
}

func (s *AnalyzeConversationRequestDialogue) SetSessionId(v string) *AnalyzeConversationRequestDialogue {
	s.SessionId = &v
	return s
}

type AnalyzeConversationRequestDialogueSentences struct {
	// This parameter is required.
	//
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// This parameter is required.
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s AnalyzeConversationRequestDialogueSentences) String() string {
	return tea.Prettify(s)
}

func (s AnalyzeConversationRequestDialogueSentences) GoString() string {
	return s.String()
}

func (s *AnalyzeConversationRequestDialogueSentences) SetRole(v string) *AnalyzeConversationRequestDialogueSentences {
	s.Role = &v
	return s
}

func (s *AnalyzeConversationRequestDialogueSentences) SetText(v string) *AnalyzeConversationRequestDialogueSentences {
	s.Text = &v
	return s
}

type AnalyzeConversationRequestExamples struct {
	// This parameter is required.
	Output *string `json:"output,omitempty" xml:"output,omitempty"`
	// This parameter is required.
	Sentences []*AnalyzeConversationRequestExamplesSentences `json:"sentences,omitempty" xml:"sentences,omitempty" type:"Repeated"`
}

func (s AnalyzeConversationRequestExamples) String() string {
	return tea.Prettify(s)
}

func (s AnalyzeConversationRequestExamples) GoString() string {
	return s.String()
}

func (s *AnalyzeConversationRequestExamples) SetOutput(v string) *AnalyzeConversationRequestExamples {
	s.Output = &v
	return s
}

func (s *AnalyzeConversationRequestExamples) SetSentences(v []*AnalyzeConversationRequestExamplesSentences) *AnalyzeConversationRequestExamples {
	s.Sentences = v
	return s
}

type AnalyzeConversationRequestExamplesSentences struct {
	ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
	// This parameter is required.
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// This parameter is required.
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s AnalyzeConversationRequestExamplesSentences) String() string {
	return tea.Prettify(s)
}

func (s AnalyzeConversationRequestExamplesSentences) GoString() string {
	return s.String()
}

func (s *AnalyzeConversationRequestExamplesSentences) SetChatId(v string) *AnalyzeConversationRequestExamplesSentences {
	s.ChatId = &v
	return s
}

func (s *AnalyzeConversationRequestExamplesSentences) SetRole(v string) *AnalyzeConversationRequestExamplesSentences {
	s.Role = &v
	return s
}

func (s *AnalyzeConversationRequestExamplesSentences) SetText(v string) *AnalyzeConversationRequestExamplesSentences {
	s.Text = &v
	return s
}

type AnalyzeConversationRequestFields struct {
	// example:
	//
	// phoneNumber
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// This parameter is required.
	Desc       *string                                       `json:"desc,omitempty" xml:"desc,omitempty"`
	EnumValues []*AnalyzeConversationRequestFieldsEnumValues `json:"enumValues,omitempty" xml:"enumValues,omitempty" type:"Repeated"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s AnalyzeConversationRequestFields) String() string {
	return tea.Prettify(s)
}

func (s AnalyzeConversationRequestFields) GoString() string {
	return s.String()
}

func (s *AnalyzeConversationRequestFields) SetCode(v string) *AnalyzeConversationRequestFields {
	s.Code = &v
	return s
}

func (s *AnalyzeConversationRequestFields) SetDesc(v string) *AnalyzeConversationRequestFields {
	s.Desc = &v
	return s
}

func (s *AnalyzeConversationRequestFields) SetEnumValues(v []*AnalyzeConversationRequestFieldsEnumValues) *AnalyzeConversationRequestFields {
	s.EnumValues = v
	return s
}

func (s *AnalyzeConversationRequestFields) SetName(v string) *AnalyzeConversationRequestFields {
	s.Name = &v
	return s
}

type AnalyzeConversationRequestFieldsEnumValues struct {
	// This parameter is required.
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// This parameter is required.
	EnumValue *string `json:"enumValue,omitempty" xml:"enumValue,omitempty"`
}

func (s AnalyzeConversationRequestFieldsEnumValues) String() string {
	return tea.Prettify(s)
}

func (s AnalyzeConversationRequestFieldsEnumValues) GoString() string {
	return s.String()
}

func (s *AnalyzeConversationRequestFieldsEnumValues) SetDesc(v string) *AnalyzeConversationRequestFieldsEnumValues {
	s.Desc = &v
	return s
}

func (s *AnalyzeConversationRequestFieldsEnumValues) SetEnumValue(v string) *AnalyzeConversationRequestFieldsEnumValues {
	s.EnumValue = &v
	return s
}

type AnalyzeConversationRequestServiceInspection struct {
	// This parameter is required.
	InspectionContents []*AnalyzeConversationRequestServiceInspectionInspectionContents `json:"inspectionContents,omitempty" xml:"inspectionContents,omitempty" type:"Repeated"`
	// This parameter is required.
	InspectionIntroduction *string `json:"inspectionIntroduction,omitempty" xml:"inspectionIntroduction,omitempty"`
	// This parameter is required.
	SceneIntroduction *string `json:"sceneIntroduction,omitempty" xml:"sceneIntroduction,omitempty"`
}

func (s AnalyzeConversationRequestServiceInspection) String() string {
	return tea.Prettify(s)
}

func (s AnalyzeConversationRequestServiceInspection) GoString() string {
	return s.String()
}

func (s *AnalyzeConversationRequestServiceInspection) SetInspectionContents(v []*AnalyzeConversationRequestServiceInspectionInspectionContents) *AnalyzeConversationRequestServiceInspection {
	s.InspectionContents = v
	return s
}

func (s *AnalyzeConversationRequestServiceInspection) SetInspectionIntroduction(v string) *AnalyzeConversationRequestServiceInspection {
	s.InspectionIntroduction = &v
	return s
}

func (s *AnalyzeConversationRequestServiceInspection) SetSceneIntroduction(v string) *AnalyzeConversationRequestServiceInspection {
	s.SceneIntroduction = &v
	return s
}

type AnalyzeConversationRequestServiceInspectionInspectionContents struct {
	// This parameter is required.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s AnalyzeConversationRequestServiceInspectionInspectionContents) String() string {
	return tea.Prettify(s)
}

func (s AnalyzeConversationRequestServiceInspectionInspectionContents) GoString() string {
	return s.String()
}

func (s *AnalyzeConversationRequestServiceInspectionInspectionContents) SetContent(v string) *AnalyzeConversationRequestServiceInspectionInspectionContents {
	s.Content = &v
	return s
}

func (s *AnalyzeConversationRequestServiceInspectionInspectionContents) SetTitle(v string) *AnalyzeConversationRequestServiceInspectionInspectionContents {
	s.Title = &v
	return s
}

type AnalyzeConversationResponseBody struct {
	// example:
	//
	// InvalidUser.NotFound
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorInfo *string `json:"errorInfo,omitempty" xml:"errorInfo,omitempty"`
	// example:
	//
	// stop
	FinishReason *string `json:"finishReason,omitempty" xml:"finishReason,omitempty"`
	// example:
	//
	// 968A8634-FA2C-5381-9B3E-C552DED7E8BF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
	Text    *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s AnalyzeConversationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AnalyzeConversationResponseBody) GoString() string {
	return s.String()
}

func (s *AnalyzeConversationResponseBody) SetErrorCode(v string) *AnalyzeConversationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AnalyzeConversationResponseBody) SetErrorInfo(v string) *AnalyzeConversationResponseBody {
	s.ErrorInfo = &v
	return s
}

func (s *AnalyzeConversationResponseBody) SetFinishReason(v string) *AnalyzeConversationResponseBody {
	s.FinishReason = &v
	return s
}

func (s *AnalyzeConversationResponseBody) SetRequestId(v string) *AnalyzeConversationResponseBody {
	s.RequestId = &v
	return s
}

func (s *AnalyzeConversationResponseBody) SetSuccess(v bool) *AnalyzeConversationResponseBody {
	s.Success = &v
	return s
}

func (s *AnalyzeConversationResponseBody) SetText(v string) *AnalyzeConversationResponseBody {
	s.Text = &v
	return s
}

type AnalyzeConversationResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AnalyzeConversationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AnalyzeConversationResponse) String() string {
	return tea.Prettify(s)
}

func (s AnalyzeConversationResponse) GoString() string {
	return s.String()
}

func (s *AnalyzeConversationResponse) SetHeaders(v map[string]*string) *AnalyzeConversationResponse {
	s.Headers = v
	return s
}

func (s *AnalyzeConversationResponse) SetStatusCode(v int32) *AnalyzeConversationResponse {
	s.StatusCode = &v
	return s
}

func (s *AnalyzeConversationResponse) SetBody(v *AnalyzeConversationResponseBody) *AnalyzeConversationResponse {
	s.Body = v
	return s
}

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
// 根据类型调用大模型
//
// @param request - AnalyzeConversationRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AnalyzeConversationResponse
func (client *Client) AnalyzeConversationWithOptions(workspaceId *string, appId *string, request *AnalyzeConversationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AnalyzeConversationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Dialogue)) {
		body["dialogue"] = request.Dialogue
	}

	if !tea.BoolValue(util.IsUnset(request.Examples)) {
		body["examples"] = request.Examples
	}

	if !tea.BoolValue(util.IsUnset(request.Fields)) {
		body["fields"] = request.Fields
	}

	if !tea.BoolValue(util.IsUnset(request.ModelCode)) {
		body["modelCode"] = request.ModelCode
	}

	if !tea.BoolValue(util.IsUnset(request.ResultTypes)) {
		body["resultTypes"] = request.ResultTypes
	}

	if !tea.BoolValue(util.IsUnset(request.SceneName)) {
		body["sceneName"] = request.SceneName
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceInspection)) {
		body["serviceInspection"] = request.ServiceInspection
	}

	if !tea.BoolValue(util.IsUnset(request.Stream)) {
		body["stream"] = request.Stream
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AnalyzeConversation"),
		Version:     tea.String("2024-06-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(openapiutil.GetEncodeParam(workspaceId)) + "/ccai/app/" + tea.StringValue(openapiutil.GetEncodeParam(appId)) + "/analyze_conversation"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AnalyzeConversationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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
	runtime := &util.RuntimeOptions{}
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
