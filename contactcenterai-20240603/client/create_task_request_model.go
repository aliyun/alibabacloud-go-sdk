// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallBackUrl(v string) *CreateTaskRequest
	GetCallBackUrl() *string
	SetCategoryTags(v []*CreateTaskRequestCategoryTags) *CreateTaskRequest
	GetCategoryTags() []*CreateTaskRequestCategoryTags
	SetCustomPrompt(v string) *CreateTaskRequest
	GetCustomPrompt() *string
	SetDialogue(v *CreateTaskRequestDialogue) *CreateTaskRequest
	GetDialogue() *CreateTaskRequestDialogue
	SetExamples(v *CreateTaskRequestExamples) *CreateTaskRequest
	GetExamples() *CreateTaskRequestExamples
	SetFields(v []*CreateTaskRequestFields) *CreateTaskRequest
	GetFields() []*CreateTaskRequestFields
	SetModelCode(v string) *CreateTaskRequest
	GetModelCode() *string
	SetResponseFormatType(v string) *CreateTaskRequest
	GetResponseFormatType() *string
	SetResultTypes(v []*string) *CreateTaskRequest
	GetResultTypes() []*string
	SetServiceInspection(v *CreateTaskRequestServiceInspection) *CreateTaskRequest
	GetServiceInspection() *CreateTaskRequestServiceInspection
	SetTaskType(v string) *CreateTaskRequest
	GetTaskType() *string
	SetTemplateIds(v []*string) *CreateTaskRequest
	GetTemplateIds() []*string
	SetTranscription(v *CreateTaskRequestTranscription) *CreateTaskRequest
	GetTranscription() *CreateTaskRequestTranscription
	SetVariables(v []*CreateTaskRequestVariables) *CreateTaskRequest
	GetVariables() []*CreateTaskRequestVariables
}

type CreateTaskRequest struct {
	CallBackUrl  *string                          `json:"callBackUrl,omitempty" xml:"callBackUrl,omitempty"`
	CategoryTags []*CreateTaskRequestCategoryTags `json:"categoryTags,omitempty" xml:"categoryTags,omitempty" type:"Repeated"`
	CustomPrompt *string                          `json:"customPrompt,omitempty" xml:"customPrompt,omitempty"`
	Dialogue     *CreateTaskRequestDialogue       `json:"dialogue,omitempty" xml:"dialogue,omitempty" type:"Struct"`
	Examples     *CreateTaskRequestExamples       `json:"examples,omitempty" xml:"examples,omitempty" type:"Struct"`
	Fields       []*CreateTaskRequestFields       `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// tyxmTurbo
	ModelCode          *string                             `json:"modelCode,omitempty" xml:"modelCode,omitempty"`
	ResponseFormatType *string                             `json:"responseFormatType,omitempty" xml:"responseFormatType,omitempty"`
	ResultTypes        []*string                           `json:"resultTypes,omitempty" xml:"resultTypes,omitempty" type:"Repeated"`
	ServiceInspection  *CreateTaskRequestServiceInspection `json:"serviceInspection,omitempty" xml:"serviceInspection,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// text
	TaskType      *string                         `json:"taskType,omitempty" xml:"taskType,omitempty"`
	TemplateIds   []*string                       `json:"templateIds,omitempty" xml:"templateIds,omitempty" type:"Repeated"`
	Transcription *CreateTaskRequestTranscription `json:"transcription,omitempty" xml:"transcription,omitempty" type:"Struct"`
	Variables     []*CreateTaskRequestVariables   `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
}

func (s CreateTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateTaskRequest) GetCallBackUrl() *string {
	return s.CallBackUrl
}

func (s *CreateTaskRequest) GetCategoryTags() []*CreateTaskRequestCategoryTags {
	return s.CategoryTags
}

func (s *CreateTaskRequest) GetCustomPrompt() *string {
	return s.CustomPrompt
}

func (s *CreateTaskRequest) GetDialogue() *CreateTaskRequestDialogue {
	return s.Dialogue
}

func (s *CreateTaskRequest) GetExamples() *CreateTaskRequestExamples {
	return s.Examples
}

func (s *CreateTaskRequest) GetFields() []*CreateTaskRequestFields {
	return s.Fields
}

func (s *CreateTaskRequest) GetModelCode() *string {
	return s.ModelCode
}

func (s *CreateTaskRequest) GetResponseFormatType() *string {
	return s.ResponseFormatType
}

func (s *CreateTaskRequest) GetResultTypes() []*string {
	return s.ResultTypes
}

func (s *CreateTaskRequest) GetServiceInspection() *CreateTaskRequestServiceInspection {
	return s.ServiceInspection
}

func (s *CreateTaskRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *CreateTaskRequest) GetTemplateIds() []*string {
	return s.TemplateIds
}

func (s *CreateTaskRequest) GetTranscription() *CreateTaskRequestTranscription {
	return s.Transcription
}

func (s *CreateTaskRequest) GetVariables() []*CreateTaskRequestVariables {
	return s.Variables
}

func (s *CreateTaskRequest) SetCallBackUrl(v string) *CreateTaskRequest {
	s.CallBackUrl = &v
	return s
}

func (s *CreateTaskRequest) SetCategoryTags(v []*CreateTaskRequestCategoryTags) *CreateTaskRequest {
	s.CategoryTags = v
	return s
}

func (s *CreateTaskRequest) SetCustomPrompt(v string) *CreateTaskRequest {
	s.CustomPrompt = &v
	return s
}

func (s *CreateTaskRequest) SetDialogue(v *CreateTaskRequestDialogue) *CreateTaskRequest {
	s.Dialogue = v
	return s
}

func (s *CreateTaskRequest) SetExamples(v *CreateTaskRequestExamples) *CreateTaskRequest {
	s.Examples = v
	return s
}

func (s *CreateTaskRequest) SetFields(v []*CreateTaskRequestFields) *CreateTaskRequest {
	s.Fields = v
	return s
}

func (s *CreateTaskRequest) SetModelCode(v string) *CreateTaskRequest {
	s.ModelCode = &v
	return s
}

func (s *CreateTaskRequest) SetResponseFormatType(v string) *CreateTaskRequest {
	s.ResponseFormatType = &v
	return s
}

func (s *CreateTaskRequest) SetResultTypes(v []*string) *CreateTaskRequest {
	s.ResultTypes = v
	return s
}

func (s *CreateTaskRequest) SetServiceInspection(v *CreateTaskRequestServiceInspection) *CreateTaskRequest {
	s.ServiceInspection = v
	return s
}

func (s *CreateTaskRequest) SetTaskType(v string) *CreateTaskRequest {
	s.TaskType = &v
	return s
}

func (s *CreateTaskRequest) SetTemplateIds(v []*string) *CreateTaskRequest {
	s.TemplateIds = v
	return s
}

func (s *CreateTaskRequest) SetTranscription(v *CreateTaskRequestTranscription) *CreateTaskRequest {
	s.Transcription = v
	return s
}

func (s *CreateTaskRequest) SetVariables(v []*CreateTaskRequestVariables) *CreateTaskRequest {
	s.Variables = v
	return s
}

func (s *CreateTaskRequest) Validate() error {
	return dara.Validate(s)
}

type CreateTaskRequestCategoryTags struct {
	TagDesc *string `json:"tagDesc,omitempty" xml:"tagDesc,omitempty"`
	TagName *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
}

func (s CreateTaskRequestCategoryTags) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequestCategoryTags) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestCategoryTags) GetTagDesc() *string {
	return s.TagDesc
}

func (s *CreateTaskRequestCategoryTags) GetTagName() *string {
	return s.TagName
}

func (s *CreateTaskRequestCategoryTags) SetTagDesc(v string) *CreateTaskRequestCategoryTags {
	s.TagDesc = &v
	return s
}

func (s *CreateTaskRequestCategoryTags) SetTagName(v string) *CreateTaskRequestCategoryTags {
	s.TagName = &v
	return s
}

func (s *CreateTaskRequestCategoryTags) Validate() error {
	return dara.Validate(s)
}

type CreateTaskRequestDialogue struct {
	// This parameter is required.
	Sentences []*CreateTaskRequestDialogueSentences `json:"sentences,omitempty" xml:"sentences,omitempty" type:"Repeated"`
	// example:
	//
	// session-01
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s CreateTaskRequestDialogue) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequestDialogue) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestDialogue) GetSentences() []*CreateTaskRequestDialogueSentences {
	return s.Sentences
}

func (s *CreateTaskRequestDialogue) GetSessionId() *string {
	return s.SessionId
}

func (s *CreateTaskRequestDialogue) SetSentences(v []*CreateTaskRequestDialogueSentences) *CreateTaskRequestDialogue {
	s.Sentences = v
	return s
}

func (s *CreateTaskRequestDialogue) SetSessionId(v string) *CreateTaskRequestDialogue {
	s.SessionId = &v
	return s
}

func (s *CreateTaskRequestDialogue) Validate() error {
	return dara.Validate(s)
}

type CreateTaskRequestDialogueSentences struct {
	// This parameter is required.
	//
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// This parameter is required.
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s CreateTaskRequestDialogueSentences) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequestDialogueSentences) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestDialogueSentences) GetRole() *string {
	return s.Role
}

func (s *CreateTaskRequestDialogueSentences) GetText() *string {
	return s.Text
}

func (s *CreateTaskRequestDialogueSentences) SetRole(v string) *CreateTaskRequestDialogueSentences {
	s.Role = &v
	return s
}

func (s *CreateTaskRequestDialogueSentences) SetText(v string) *CreateTaskRequestDialogueSentences {
	s.Text = &v
	return s
}

func (s *CreateTaskRequestDialogueSentences) Validate() error {
	return dara.Validate(s)
}

type CreateTaskRequestExamples struct {
	Output *string `json:"output,omitempty" xml:"output,omitempty"`
	// This parameter is required.
	Sentences []*CreateTaskRequestExamplesSentences `json:"sentences,omitempty" xml:"sentences,omitempty" type:"Repeated"`
}

func (s CreateTaskRequestExamples) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequestExamples) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestExamples) GetOutput() *string {
	return s.Output
}

func (s *CreateTaskRequestExamples) GetSentences() []*CreateTaskRequestExamplesSentences {
	return s.Sentences
}

func (s *CreateTaskRequestExamples) SetOutput(v string) *CreateTaskRequestExamples {
	s.Output = &v
	return s
}

func (s *CreateTaskRequestExamples) SetSentences(v []*CreateTaskRequestExamplesSentences) *CreateTaskRequestExamples {
	s.Sentences = v
	return s
}

func (s *CreateTaskRequestExamples) Validate() error {
	return dara.Validate(s)
}

type CreateTaskRequestExamplesSentences struct {
	// This parameter is required.
	//
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// This parameter is required.
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s CreateTaskRequestExamplesSentences) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequestExamplesSentences) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestExamplesSentences) GetRole() *string {
	return s.Role
}

func (s *CreateTaskRequestExamplesSentences) GetText() *string {
	return s.Text
}

func (s *CreateTaskRequestExamplesSentences) SetRole(v string) *CreateTaskRequestExamplesSentences {
	s.Role = &v
	return s
}

func (s *CreateTaskRequestExamplesSentences) SetText(v string) *CreateTaskRequestExamplesSentences {
	s.Text = &v
	return s
}

func (s *CreateTaskRequestExamplesSentences) Validate() error {
	return dara.Validate(s)
}

type CreateTaskRequestFields struct {
	// example:
	//
	// phoneNumber
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// This parameter is required.
	Desc       *string                              `json:"desc,omitempty" xml:"desc,omitempty"`
	EnumValues []*CreateTaskRequestFieldsEnumValues `json:"enumValues,omitempty" xml:"enumValues,omitempty" type:"Repeated"`
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateTaskRequestFields) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequestFields) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestFields) GetCode() *string {
	return s.Code
}

func (s *CreateTaskRequestFields) GetDesc() *string {
	return s.Desc
}

func (s *CreateTaskRequestFields) GetEnumValues() []*CreateTaskRequestFieldsEnumValues {
	return s.EnumValues
}

func (s *CreateTaskRequestFields) GetName() *string {
	return s.Name
}

func (s *CreateTaskRequestFields) SetCode(v string) *CreateTaskRequestFields {
	s.Code = &v
	return s
}

func (s *CreateTaskRequestFields) SetDesc(v string) *CreateTaskRequestFields {
	s.Desc = &v
	return s
}

func (s *CreateTaskRequestFields) SetEnumValues(v []*CreateTaskRequestFieldsEnumValues) *CreateTaskRequestFields {
	s.EnumValues = v
	return s
}

func (s *CreateTaskRequestFields) SetName(v string) *CreateTaskRequestFields {
	s.Name = &v
	return s
}

func (s *CreateTaskRequestFields) Validate() error {
	return dara.Validate(s)
}

type CreateTaskRequestFieldsEnumValues struct {
	// This parameter is required.
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// This parameter is required.
	EnumValue *string `json:"enumValue,omitempty" xml:"enumValue,omitempty"`
}

func (s CreateTaskRequestFieldsEnumValues) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequestFieldsEnumValues) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestFieldsEnumValues) GetDesc() *string {
	return s.Desc
}

func (s *CreateTaskRequestFieldsEnumValues) GetEnumValue() *string {
	return s.EnumValue
}

func (s *CreateTaskRequestFieldsEnumValues) SetDesc(v string) *CreateTaskRequestFieldsEnumValues {
	s.Desc = &v
	return s
}

func (s *CreateTaskRequestFieldsEnumValues) SetEnumValue(v string) *CreateTaskRequestFieldsEnumValues {
	s.EnumValue = &v
	return s
}

func (s *CreateTaskRequestFieldsEnumValues) Validate() error {
	return dara.Validate(s)
}

type CreateTaskRequestServiceInspection struct {
	// This parameter is required.
	InspectionContents []*CreateTaskRequestServiceInspectionInspectionContents `json:"inspectionContents,omitempty" xml:"inspectionContents,omitempty" type:"Repeated"`
	// This parameter is required.
	InspectionIntroduction *string `json:"inspectionIntroduction,omitempty" xml:"inspectionIntroduction,omitempty"`
	// This parameter is required.
	SceneIntroduction *string `json:"sceneIntroduction,omitempty" xml:"sceneIntroduction,omitempty"`
}

func (s CreateTaskRequestServiceInspection) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequestServiceInspection) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestServiceInspection) GetInspectionContents() []*CreateTaskRequestServiceInspectionInspectionContents {
	return s.InspectionContents
}

func (s *CreateTaskRequestServiceInspection) GetInspectionIntroduction() *string {
	return s.InspectionIntroduction
}

func (s *CreateTaskRequestServiceInspection) GetSceneIntroduction() *string {
	return s.SceneIntroduction
}

func (s *CreateTaskRequestServiceInspection) SetInspectionContents(v []*CreateTaskRequestServiceInspectionInspectionContents) *CreateTaskRequestServiceInspection {
	s.InspectionContents = v
	return s
}

func (s *CreateTaskRequestServiceInspection) SetInspectionIntroduction(v string) *CreateTaskRequestServiceInspection {
	s.InspectionIntroduction = &v
	return s
}

func (s *CreateTaskRequestServiceInspection) SetSceneIntroduction(v string) *CreateTaskRequestServiceInspection {
	s.SceneIntroduction = &v
	return s
}

func (s *CreateTaskRequestServiceInspection) Validate() error {
	return dara.Validate(s)
}

type CreateTaskRequestServiceInspectionInspectionContents struct {
	// This parameter is required.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateTaskRequestServiceInspectionInspectionContents) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequestServiceInspectionInspectionContents) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestServiceInspectionInspectionContents) GetContent() *string {
	return s.Content
}

func (s *CreateTaskRequestServiceInspectionInspectionContents) GetTitle() *string {
	return s.Title
}

func (s *CreateTaskRequestServiceInspectionInspectionContents) SetContent(v string) *CreateTaskRequestServiceInspectionInspectionContents {
	s.Content = &v
	return s
}

func (s *CreateTaskRequestServiceInspectionInspectionContents) SetTitle(v string) *CreateTaskRequestServiceInspectionInspectionContents {
	s.Title = &v
	return s
}

func (s *CreateTaskRequestServiceInspectionInspectionContents) Validate() error {
	return dara.Validate(s)
}

type CreateTaskRequestTranscription struct {
	AsrModelCode *string `json:"asrModelCode,omitempty" xml:"asrModelCode,omitempty"`
	// example:
	//
	// 1
	AutoSplit *int32 `json:"autoSplit,omitempty" xml:"autoSplit,omitempty"`
	// example:
	//
	// 1
	ClientChannel *int32 `json:"clientChannel,omitempty" xml:"clientChannel,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sss.mp3
	FileName           *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	LanguageHints      *string `json:"languageHints,omitempty" xml:"languageHints,omitempty"`
	Level              *string `json:"level,omitempty" xml:"level,omitempty"`
	RoleIdentification *bool   `json:"roleIdentification,omitempty" xml:"roleIdentification,omitempty"`
	// example:
	//
	// 1
	ServiceChannel         *int32    `json:"serviceChannel,omitempty" xml:"serviceChannel,omitempty"`
	ServiceChannelKeywords []*string `json:"serviceChannelKeywords,omitempty" xml:"serviceChannelKeywords,omitempty" type:"Repeated"`
	VocabularyId           *string   `json:"vocabularyId,omitempty" xml:"vocabularyId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://1111.com/sss.mp3
	VoiceFileUrl *string `json:"voiceFileUrl,omitempty" xml:"voiceFileUrl,omitempty"`
}

func (s CreateTaskRequestTranscription) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequestTranscription) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestTranscription) GetAsrModelCode() *string {
	return s.AsrModelCode
}

func (s *CreateTaskRequestTranscription) GetAutoSplit() *int32 {
	return s.AutoSplit
}

func (s *CreateTaskRequestTranscription) GetClientChannel() *int32 {
	return s.ClientChannel
}

func (s *CreateTaskRequestTranscription) GetFileName() *string {
	return s.FileName
}

func (s *CreateTaskRequestTranscription) GetLanguageHints() *string {
	return s.LanguageHints
}

func (s *CreateTaskRequestTranscription) GetLevel() *string {
	return s.Level
}

func (s *CreateTaskRequestTranscription) GetRoleIdentification() *bool {
	return s.RoleIdentification
}

func (s *CreateTaskRequestTranscription) GetServiceChannel() *int32 {
	return s.ServiceChannel
}

func (s *CreateTaskRequestTranscription) GetServiceChannelKeywords() []*string {
	return s.ServiceChannelKeywords
}

func (s *CreateTaskRequestTranscription) GetVocabularyId() *string {
	return s.VocabularyId
}

func (s *CreateTaskRequestTranscription) GetVoiceFileUrl() *string {
	return s.VoiceFileUrl
}

func (s *CreateTaskRequestTranscription) SetAsrModelCode(v string) *CreateTaskRequestTranscription {
	s.AsrModelCode = &v
	return s
}

func (s *CreateTaskRequestTranscription) SetAutoSplit(v int32) *CreateTaskRequestTranscription {
	s.AutoSplit = &v
	return s
}

func (s *CreateTaskRequestTranscription) SetClientChannel(v int32) *CreateTaskRequestTranscription {
	s.ClientChannel = &v
	return s
}

func (s *CreateTaskRequestTranscription) SetFileName(v string) *CreateTaskRequestTranscription {
	s.FileName = &v
	return s
}

func (s *CreateTaskRequestTranscription) SetLanguageHints(v string) *CreateTaskRequestTranscription {
	s.LanguageHints = &v
	return s
}

func (s *CreateTaskRequestTranscription) SetLevel(v string) *CreateTaskRequestTranscription {
	s.Level = &v
	return s
}

func (s *CreateTaskRequestTranscription) SetRoleIdentification(v bool) *CreateTaskRequestTranscription {
	s.RoleIdentification = &v
	return s
}

func (s *CreateTaskRequestTranscription) SetServiceChannel(v int32) *CreateTaskRequestTranscription {
	s.ServiceChannel = &v
	return s
}

func (s *CreateTaskRequestTranscription) SetServiceChannelKeywords(v []*string) *CreateTaskRequestTranscription {
	s.ServiceChannelKeywords = v
	return s
}

func (s *CreateTaskRequestTranscription) SetVocabularyId(v string) *CreateTaskRequestTranscription {
	s.VocabularyId = &v
	return s
}

func (s *CreateTaskRequestTranscription) SetVoiceFileUrl(v string) *CreateTaskRequestTranscription {
	s.VoiceFileUrl = &v
	return s
}

func (s *CreateTaskRequestTranscription) Validate() error {
	return dara.Validate(s)
}

type CreateTaskRequestVariables struct {
	VariableCode  *string `json:"variableCode,omitempty" xml:"variableCode,omitempty"`
	VariableValue *string `json:"variableValue,omitempty" xml:"variableValue,omitempty"`
}

func (s CreateTaskRequestVariables) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequestVariables) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestVariables) GetVariableCode() *string {
	return s.VariableCode
}

func (s *CreateTaskRequestVariables) GetVariableValue() *string {
	return s.VariableValue
}

func (s *CreateTaskRequestVariables) SetVariableCode(v string) *CreateTaskRequestVariables {
	s.VariableCode = &v
	return s
}

func (s *CreateTaskRequestVariables) SetVariableValue(v string) *CreateTaskRequestVariables {
	s.VariableValue = &v
	return s
}

func (s *CreateTaskRequestVariables) Validate() error {
	return dara.Validate(s)
}
