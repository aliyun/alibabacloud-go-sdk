// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnalyzeConversationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryTags(v []*AnalyzeConversationRequestCategoryTags) *AnalyzeConversationRequest
	GetCategoryTags() []*AnalyzeConversationRequestCategoryTags
	SetCustomPrompt(v string) *AnalyzeConversationRequest
	GetCustomPrompt() *string
	SetDialogue(v *AnalyzeConversationRequestDialogue) *AnalyzeConversationRequest
	GetDialogue() *AnalyzeConversationRequestDialogue
	SetExamples(v []*AnalyzeConversationRequestExamples) *AnalyzeConversationRequest
	GetExamples() []*AnalyzeConversationRequestExamples
	SetFields(v []*AnalyzeConversationRequestFields) *AnalyzeConversationRequest
	GetFields() []*AnalyzeConversationRequestFields
	SetModelCode(v string) *AnalyzeConversationRequest
	GetModelCode() *string
	SetResponseFormatType(v string) *AnalyzeConversationRequest
	GetResponseFormatType() *string
	SetResultTypes(v []*string) *AnalyzeConversationRequest
	GetResultTypes() []*string
	SetSceneName(v string) *AnalyzeConversationRequest
	GetSceneName() *string
	SetServiceInspection(v *AnalyzeConversationRequestServiceInspection) *AnalyzeConversationRequest
	GetServiceInspection() *AnalyzeConversationRequestServiceInspection
	SetSourceCallerUid(v string) *AnalyzeConversationRequest
	GetSourceCallerUid() *string
	SetStream(v bool) *AnalyzeConversationRequest
	GetStream() *bool
	SetTimeConstraintList(v []*string) *AnalyzeConversationRequest
	GetTimeConstraintList() []*string
	SetUserProfiles(v []*AnalyzeConversationRequestUserProfiles) *AnalyzeConversationRequest
	GetUserProfiles() []*AnalyzeConversationRequestUserProfiles
}

type AnalyzeConversationRequest struct {
	CategoryTags []*AnalyzeConversationRequestCategoryTags `json:"categoryTags,omitempty" xml:"categoryTags,omitempty" type:"Repeated"`
	CustomPrompt *string                                   `json:"customPrompt,omitempty" xml:"customPrompt,omitempty"`
	Dialogue     *AnalyzeConversationRequestDialogue       `json:"dialogue,omitempty" xml:"dialogue,omitempty" type:"Struct"`
	Examples     []*AnalyzeConversationRequestExamples     `json:"examples,omitempty" xml:"examples,omitempty" type:"Repeated"`
	Fields       []*AnalyzeConversationRequestFields       `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	// example:
	//
	// tyxmTurbo
	ModelCode          *string `json:"modelCode,omitempty" xml:"modelCode,omitempty"`
	ResponseFormatType *string `json:"responseFormatType,omitempty" xml:"responseFormatType,omitempty"`
	// This parameter is required.
	ResultTypes       []*string                                    `json:"resultTypes,omitempty" xml:"resultTypes,omitempty" type:"Repeated"`
	SceneName         *string                                      `json:"sceneName,omitempty" xml:"sceneName,omitempty"`
	ServiceInspection *AnalyzeConversationRequestServiceInspection `json:"serviceInspection,omitempty" xml:"serviceInspection,omitempty" type:"Struct"`
	SourceCallerUid   *string                                      `json:"sourceCallerUid,omitempty" xml:"sourceCallerUid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	Stream             *bool                                     `json:"stream,omitempty" xml:"stream,omitempty"`
	TimeConstraintList []*string                                 `json:"timeConstraintList,omitempty" xml:"timeConstraintList,omitempty" type:"Repeated"`
	UserProfiles       []*AnalyzeConversationRequestUserProfiles `json:"userProfiles,omitempty" xml:"userProfiles,omitempty" type:"Repeated"`
}

func (s AnalyzeConversationRequest) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeConversationRequest) GoString() string {
	return s.String()
}

func (s *AnalyzeConversationRequest) GetCategoryTags() []*AnalyzeConversationRequestCategoryTags {
	return s.CategoryTags
}

func (s *AnalyzeConversationRequest) GetCustomPrompt() *string {
	return s.CustomPrompt
}

func (s *AnalyzeConversationRequest) GetDialogue() *AnalyzeConversationRequestDialogue {
	return s.Dialogue
}

func (s *AnalyzeConversationRequest) GetExamples() []*AnalyzeConversationRequestExamples {
	return s.Examples
}

func (s *AnalyzeConversationRequest) GetFields() []*AnalyzeConversationRequestFields {
	return s.Fields
}

func (s *AnalyzeConversationRequest) GetModelCode() *string {
	return s.ModelCode
}

func (s *AnalyzeConversationRequest) GetResponseFormatType() *string {
	return s.ResponseFormatType
}

func (s *AnalyzeConversationRequest) GetResultTypes() []*string {
	return s.ResultTypes
}

func (s *AnalyzeConversationRequest) GetSceneName() *string {
	return s.SceneName
}

func (s *AnalyzeConversationRequest) GetServiceInspection() *AnalyzeConversationRequestServiceInspection {
	return s.ServiceInspection
}

func (s *AnalyzeConversationRequest) GetSourceCallerUid() *string {
	return s.SourceCallerUid
}

func (s *AnalyzeConversationRequest) GetStream() *bool {
	return s.Stream
}

func (s *AnalyzeConversationRequest) GetTimeConstraintList() []*string {
	return s.TimeConstraintList
}

func (s *AnalyzeConversationRequest) GetUserProfiles() []*AnalyzeConversationRequestUserProfiles {
	return s.UserProfiles
}

func (s *AnalyzeConversationRequest) SetCategoryTags(v []*AnalyzeConversationRequestCategoryTags) *AnalyzeConversationRequest {
	s.CategoryTags = v
	return s
}

func (s *AnalyzeConversationRequest) SetCustomPrompt(v string) *AnalyzeConversationRequest {
	s.CustomPrompt = &v
	return s
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

func (s *AnalyzeConversationRequest) SetResponseFormatType(v string) *AnalyzeConversationRequest {
	s.ResponseFormatType = &v
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

func (s *AnalyzeConversationRequest) SetSourceCallerUid(v string) *AnalyzeConversationRequest {
	s.SourceCallerUid = &v
	return s
}

func (s *AnalyzeConversationRequest) SetStream(v bool) *AnalyzeConversationRequest {
	s.Stream = &v
	return s
}

func (s *AnalyzeConversationRequest) SetTimeConstraintList(v []*string) *AnalyzeConversationRequest {
	s.TimeConstraintList = v
	return s
}

func (s *AnalyzeConversationRequest) SetUserProfiles(v []*AnalyzeConversationRequestUserProfiles) *AnalyzeConversationRequest {
	s.UserProfiles = v
	return s
}

func (s *AnalyzeConversationRequest) Validate() error {
	return dara.Validate(s)
}

type AnalyzeConversationRequestCategoryTags struct {
	TagDesc *string `json:"tagDesc,omitempty" xml:"tagDesc,omitempty"`
	TagName *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
}

func (s AnalyzeConversationRequestCategoryTags) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeConversationRequestCategoryTags) GoString() string {
	return s.String()
}

func (s *AnalyzeConversationRequestCategoryTags) GetTagDesc() *string {
	return s.TagDesc
}

func (s *AnalyzeConversationRequestCategoryTags) GetTagName() *string {
	return s.TagName
}

func (s *AnalyzeConversationRequestCategoryTags) SetTagDesc(v string) *AnalyzeConversationRequestCategoryTags {
	s.TagDesc = &v
	return s
}

func (s *AnalyzeConversationRequestCategoryTags) SetTagName(v string) *AnalyzeConversationRequestCategoryTags {
	s.TagName = &v
	return s
}

func (s *AnalyzeConversationRequestCategoryTags) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s AnalyzeConversationRequestDialogue) GoString() string {
	return s.String()
}

func (s *AnalyzeConversationRequestDialogue) GetSentences() []*AnalyzeConversationRequestDialogueSentences {
	return s.Sentences
}

func (s *AnalyzeConversationRequestDialogue) GetSessionId() *string {
	return s.SessionId
}

func (s *AnalyzeConversationRequestDialogue) SetSentences(v []*AnalyzeConversationRequestDialogueSentences) *AnalyzeConversationRequestDialogue {
	s.Sentences = v
	return s
}

func (s *AnalyzeConversationRequestDialogue) SetSessionId(v string) *AnalyzeConversationRequestDialogue {
	s.SessionId = &v
	return s
}

func (s *AnalyzeConversationRequestDialogue) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s AnalyzeConversationRequestDialogueSentences) GoString() string {
	return s.String()
}

func (s *AnalyzeConversationRequestDialogueSentences) GetRole() *string {
	return s.Role
}

func (s *AnalyzeConversationRequestDialogueSentences) GetText() *string {
	return s.Text
}

func (s *AnalyzeConversationRequestDialogueSentences) SetRole(v string) *AnalyzeConversationRequestDialogueSentences {
	s.Role = &v
	return s
}

func (s *AnalyzeConversationRequestDialogueSentences) SetText(v string) *AnalyzeConversationRequestDialogueSentences {
	s.Text = &v
	return s
}

func (s *AnalyzeConversationRequestDialogueSentences) Validate() error {
	return dara.Validate(s)
}

type AnalyzeConversationRequestExamples struct {
	// This parameter is required.
	Output *string `json:"output,omitempty" xml:"output,omitempty"`
	// This parameter is required.
	Sentences []*AnalyzeConversationRequestExamplesSentences `json:"sentences,omitempty" xml:"sentences,omitempty" type:"Repeated"`
}

func (s AnalyzeConversationRequestExamples) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeConversationRequestExamples) GoString() string {
	return s.String()
}

func (s *AnalyzeConversationRequestExamples) GetOutput() *string {
	return s.Output
}

func (s *AnalyzeConversationRequestExamples) GetSentences() []*AnalyzeConversationRequestExamplesSentences {
	return s.Sentences
}

func (s *AnalyzeConversationRequestExamples) SetOutput(v string) *AnalyzeConversationRequestExamples {
	s.Output = &v
	return s
}

func (s *AnalyzeConversationRequestExamples) SetSentences(v []*AnalyzeConversationRequestExamplesSentences) *AnalyzeConversationRequestExamples {
	s.Sentences = v
	return s
}

func (s *AnalyzeConversationRequestExamples) Validate() error {
	return dara.Validate(s)
}

type AnalyzeConversationRequestExamplesSentences struct {
	ChatId *string `json:"chatId,omitempty" xml:"chatId,omitempty"`
	// This parameter is required.
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// This parameter is required.
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s AnalyzeConversationRequestExamplesSentences) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeConversationRequestExamplesSentences) GoString() string {
	return s.String()
}

func (s *AnalyzeConversationRequestExamplesSentences) GetChatId() *string {
	return s.ChatId
}

func (s *AnalyzeConversationRequestExamplesSentences) GetRole() *string {
	return s.Role
}

func (s *AnalyzeConversationRequestExamplesSentences) GetText() *string {
	return s.Text
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

func (s *AnalyzeConversationRequestExamplesSentences) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s AnalyzeConversationRequestFields) GoString() string {
	return s.String()
}

func (s *AnalyzeConversationRequestFields) GetCode() *string {
	return s.Code
}

func (s *AnalyzeConversationRequestFields) GetDesc() *string {
	return s.Desc
}

func (s *AnalyzeConversationRequestFields) GetEnumValues() []*AnalyzeConversationRequestFieldsEnumValues {
	return s.EnumValues
}

func (s *AnalyzeConversationRequestFields) GetName() *string {
	return s.Name
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

func (s *AnalyzeConversationRequestFields) Validate() error {
	return dara.Validate(s)
}

type AnalyzeConversationRequestFieldsEnumValues struct {
	// This parameter is required.
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// This parameter is required.
	EnumValue *string `json:"enumValue,omitempty" xml:"enumValue,omitempty"`
}

func (s AnalyzeConversationRequestFieldsEnumValues) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeConversationRequestFieldsEnumValues) GoString() string {
	return s.String()
}

func (s *AnalyzeConversationRequestFieldsEnumValues) GetDesc() *string {
	return s.Desc
}

func (s *AnalyzeConversationRequestFieldsEnumValues) GetEnumValue() *string {
	return s.EnumValue
}

func (s *AnalyzeConversationRequestFieldsEnumValues) SetDesc(v string) *AnalyzeConversationRequestFieldsEnumValues {
	s.Desc = &v
	return s
}

func (s *AnalyzeConversationRequestFieldsEnumValues) SetEnumValue(v string) *AnalyzeConversationRequestFieldsEnumValues {
	s.EnumValue = &v
	return s
}

func (s *AnalyzeConversationRequestFieldsEnumValues) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s AnalyzeConversationRequestServiceInspection) GoString() string {
	return s.String()
}

func (s *AnalyzeConversationRequestServiceInspection) GetInspectionContents() []*AnalyzeConversationRequestServiceInspectionInspectionContents {
	return s.InspectionContents
}

func (s *AnalyzeConversationRequestServiceInspection) GetInspectionIntroduction() *string {
	return s.InspectionIntroduction
}

func (s *AnalyzeConversationRequestServiceInspection) GetSceneIntroduction() *string {
	return s.SceneIntroduction
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

func (s *AnalyzeConversationRequestServiceInspection) Validate() error {
	return dara.Validate(s)
}

type AnalyzeConversationRequestServiceInspectionInspectionContents struct {
	// This parameter is required.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// This parameter is required.
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s AnalyzeConversationRequestServiceInspectionInspectionContents) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeConversationRequestServiceInspectionInspectionContents) GoString() string {
	return s.String()
}

func (s *AnalyzeConversationRequestServiceInspectionInspectionContents) GetContent() *string {
	return s.Content
}

func (s *AnalyzeConversationRequestServiceInspectionInspectionContents) GetTitle() *string {
	return s.Title
}

func (s *AnalyzeConversationRequestServiceInspectionInspectionContents) SetContent(v string) *AnalyzeConversationRequestServiceInspectionInspectionContents {
	s.Content = &v
	return s
}

func (s *AnalyzeConversationRequestServiceInspectionInspectionContents) SetTitle(v string) *AnalyzeConversationRequestServiceInspectionInspectionContents {
	s.Title = &v
	return s
}

func (s *AnalyzeConversationRequestServiceInspectionInspectionContents) Validate() error {
	return dara.Validate(s)
}

type AnalyzeConversationRequestUserProfiles struct {
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s AnalyzeConversationRequestUserProfiles) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeConversationRequestUserProfiles) GoString() string {
	return s.String()
}

func (s *AnalyzeConversationRequestUserProfiles) GetName() *string {
	return s.Name
}

func (s *AnalyzeConversationRequestUserProfiles) GetValue() *string {
	return s.Value
}

func (s *AnalyzeConversationRequestUserProfiles) SetName(v string) *AnalyzeConversationRequestUserProfiles {
	s.Name = &v
	return s
}

func (s *AnalyzeConversationRequestUserProfiles) SetValue(v string) *AnalyzeConversationRequestUserProfiles {
	s.Value = &v
	return s
}

func (s *AnalyzeConversationRequestUserProfiles) Validate() error {
	return dara.Validate(s)
}
