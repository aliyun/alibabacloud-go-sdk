// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunCompletionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDialogue(v *RunCompletionRequestDialogue) *RunCompletionRequest
	GetDialogue() *RunCompletionRequestDialogue
	SetFields(v []*RunCompletionRequestFields) *RunCompletionRequest
	GetFields() []*RunCompletionRequestFields
	SetModelCode(v string) *RunCompletionRequest
	GetModelCode() *string
	SetServiceInspection(v *RunCompletionRequestServiceInspection) *RunCompletionRequest
	GetServiceInspection() *RunCompletionRequestServiceInspection
	SetStream(v bool) *RunCompletionRequest
	GetStream() *bool
	SetTemplateIds(v []*int64) *RunCompletionRequest
	GetTemplateIds() []*int64
	SetResponseFormatType(v string) *RunCompletionRequest
	GetResponseFormatType() *string
	SetVariables(v []*RunCompletionRequestVariables) *RunCompletionRequest
	GetVariables() []*RunCompletionRequestVariables
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
	TemplateIds        []*int64                         `json:"TemplateIds,omitempty" xml:"TemplateIds,omitempty" type:"Repeated"`
	ResponseFormatType *string                          `json:"responseFormatType,omitempty" xml:"responseFormatType,omitempty"`
	Variables          []*RunCompletionRequestVariables `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
}

func (s RunCompletionRequest) String() string {
	return dara.Prettify(s)
}

func (s RunCompletionRequest) GoString() string {
	return s.String()
}

func (s *RunCompletionRequest) GetDialogue() *RunCompletionRequestDialogue {
	return s.Dialogue
}

func (s *RunCompletionRequest) GetFields() []*RunCompletionRequestFields {
	return s.Fields
}

func (s *RunCompletionRequest) GetModelCode() *string {
	return s.ModelCode
}

func (s *RunCompletionRequest) GetServiceInspection() *RunCompletionRequestServiceInspection {
	return s.ServiceInspection
}

func (s *RunCompletionRequest) GetStream() *bool {
	return s.Stream
}

func (s *RunCompletionRequest) GetTemplateIds() []*int64 {
	return s.TemplateIds
}

func (s *RunCompletionRequest) GetResponseFormatType() *string {
	return s.ResponseFormatType
}

func (s *RunCompletionRequest) GetVariables() []*RunCompletionRequestVariables {
	return s.Variables
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

func (s *RunCompletionRequest) SetResponseFormatType(v string) *RunCompletionRequest {
	s.ResponseFormatType = &v
	return s
}

func (s *RunCompletionRequest) SetVariables(v []*RunCompletionRequestVariables) *RunCompletionRequest {
	s.Variables = v
	return s
}

func (s *RunCompletionRequest) Validate() error {
	return dara.Validate(s)
}

type RunCompletionRequestDialogue struct {
	Sentences []*RunCompletionRequestDialogueSentences `json:"Sentences,omitempty" xml:"Sentences,omitempty" type:"Repeated"`
	// example:
	//
	// d25zc9c7004f8dad2b454d
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s RunCompletionRequestDialogue) String() string {
	return dara.Prettify(s)
}

func (s RunCompletionRequestDialogue) GoString() string {
	return s.String()
}

func (s *RunCompletionRequestDialogue) GetSentences() []*RunCompletionRequestDialogueSentences {
	return s.Sentences
}

func (s *RunCompletionRequestDialogue) GetSessionId() *string {
	return s.SessionId
}

func (s *RunCompletionRequestDialogue) SetSentences(v []*RunCompletionRequestDialogueSentences) *RunCompletionRequestDialogue {
	s.Sentences = v
	return s
}

func (s *RunCompletionRequestDialogue) SetSessionId(v string) *RunCompletionRequestDialogue {
	s.SessionId = &v
	return s
}

func (s *RunCompletionRequestDialogue) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s RunCompletionRequestDialogueSentences) GoString() string {
	return s.String()
}

func (s *RunCompletionRequestDialogueSentences) GetChatId() *string {
	return s.ChatId
}

func (s *RunCompletionRequestDialogueSentences) GetRole() *string {
	return s.Role
}

func (s *RunCompletionRequestDialogueSentences) GetText() *string {
	return s.Text
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

func (s *RunCompletionRequestDialogueSentences) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s RunCompletionRequestFields) GoString() string {
	return s.String()
}

func (s *RunCompletionRequestFields) GetCode() *string {
	return s.Code
}

func (s *RunCompletionRequestFields) GetDesc() *string {
	return s.Desc
}

func (s *RunCompletionRequestFields) GetEnumValues() []*RunCompletionRequestFieldsEnumValues {
	return s.EnumValues
}

func (s *RunCompletionRequestFields) GetName() *string {
	return s.Name
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

func (s *RunCompletionRequestFields) Validate() error {
	return dara.Validate(s)
}

type RunCompletionRequestFieldsEnumValues struct {
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// This parameter is required.
	EnumValue *string `json:"EnumValue,omitempty" xml:"EnumValue,omitempty"`
}

func (s RunCompletionRequestFieldsEnumValues) String() string {
	return dara.Prettify(s)
}

func (s RunCompletionRequestFieldsEnumValues) GoString() string {
	return s.String()
}

func (s *RunCompletionRequestFieldsEnumValues) GetDesc() *string {
	return s.Desc
}

func (s *RunCompletionRequestFieldsEnumValues) GetEnumValue() *string {
	return s.EnumValue
}

func (s *RunCompletionRequestFieldsEnumValues) SetDesc(v string) *RunCompletionRequestFieldsEnumValues {
	s.Desc = &v
	return s
}

func (s *RunCompletionRequestFieldsEnumValues) SetEnumValue(v string) *RunCompletionRequestFieldsEnumValues {
	s.EnumValue = &v
	return s
}

func (s *RunCompletionRequestFieldsEnumValues) Validate() error {
	return dara.Validate(s)
}

type RunCompletionRequestServiceInspection struct {
	InspectionContents     []*RunCompletionRequestServiceInspectionInspectionContents `json:"InspectionContents,omitempty" xml:"InspectionContents,omitempty" type:"Repeated"`
	InspectionIntroduction *string                                                    `json:"InspectionIntroduction,omitempty" xml:"InspectionIntroduction,omitempty"`
	SceneIntroduction      *string                                                    `json:"SceneIntroduction,omitempty" xml:"SceneIntroduction,omitempty"`
}

func (s RunCompletionRequestServiceInspection) String() string {
	return dara.Prettify(s)
}

func (s RunCompletionRequestServiceInspection) GoString() string {
	return s.String()
}

func (s *RunCompletionRequestServiceInspection) GetInspectionContents() []*RunCompletionRequestServiceInspectionInspectionContents {
	return s.InspectionContents
}

func (s *RunCompletionRequestServiceInspection) GetInspectionIntroduction() *string {
	return s.InspectionIntroduction
}

func (s *RunCompletionRequestServiceInspection) GetSceneIntroduction() *string {
	return s.SceneIntroduction
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

func (s *RunCompletionRequestServiceInspection) Validate() error {
	return dara.Validate(s)
}

type RunCompletionRequestServiceInspectionInspectionContents struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s RunCompletionRequestServiceInspectionInspectionContents) String() string {
	return dara.Prettify(s)
}

func (s RunCompletionRequestServiceInspectionInspectionContents) GoString() string {
	return s.String()
}

func (s *RunCompletionRequestServiceInspectionInspectionContents) GetContent() *string {
	return s.Content
}

func (s *RunCompletionRequestServiceInspectionInspectionContents) GetTitle() *string {
	return s.Title
}

func (s *RunCompletionRequestServiceInspectionInspectionContents) SetContent(v string) *RunCompletionRequestServiceInspectionInspectionContents {
	s.Content = &v
	return s
}

func (s *RunCompletionRequestServiceInspectionInspectionContents) SetTitle(v string) *RunCompletionRequestServiceInspectionInspectionContents {
	s.Title = &v
	return s
}

func (s *RunCompletionRequestServiceInspectionInspectionContents) Validate() error {
	return dara.Validate(s)
}

type RunCompletionRequestVariables struct {
	VariableCode  *string `json:"variableCode,omitempty" xml:"variableCode,omitempty"`
	VariableValue *string `json:"variableValue,omitempty" xml:"variableValue,omitempty"`
}

func (s RunCompletionRequestVariables) String() string {
	return dara.Prettify(s)
}

func (s RunCompletionRequestVariables) GoString() string {
	return s.String()
}

func (s *RunCompletionRequestVariables) GetVariableCode() *string {
	return s.VariableCode
}

func (s *RunCompletionRequestVariables) GetVariableValue() *string {
	return s.VariableValue
}

func (s *RunCompletionRequestVariables) SetVariableCode(v string) *RunCompletionRequestVariables {
	s.VariableCode = &v
	return s
}

func (s *RunCompletionRequestVariables) SetVariableValue(v string) *RunCompletionRequestVariables {
	s.VariableValue = &v
	return s
}

func (s *RunCompletionRequestVariables) Validate() error {
	return dara.Validate(s)
}
