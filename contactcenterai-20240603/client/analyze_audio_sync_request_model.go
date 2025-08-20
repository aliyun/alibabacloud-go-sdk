// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnalyzeAudioSyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryTags(v []*AnalyzeAudioSyncRequestCategoryTags) *AnalyzeAudioSyncRequest
	GetCategoryTags() []*AnalyzeAudioSyncRequestCategoryTags
	SetCustomPrompt(v string) *AnalyzeAudioSyncRequest
	GetCustomPrompt() *string
	SetFields(v []*AnalyzeAudioSyncRequestFields) *AnalyzeAudioSyncRequest
	GetFields() []*AnalyzeAudioSyncRequestFields
	SetModelCode(v string) *AnalyzeAudioSyncRequest
	GetModelCode() *string
	SetResponseFormatType(v string) *AnalyzeAudioSyncRequest
	GetResponseFormatType() *string
	SetResultTypes(v []*string) *AnalyzeAudioSyncRequest
	GetResultTypes() []*string
	SetServiceInspection(v *AnalyzeAudioSyncRequestServiceInspection) *AnalyzeAudioSyncRequest
	GetServiceInspection() *AnalyzeAudioSyncRequestServiceInspection
	SetStream(v bool) *AnalyzeAudioSyncRequest
	GetStream() *bool
	SetTemplateIds(v []*string) *AnalyzeAudioSyncRequest
	GetTemplateIds() []*string
	SetTranscription(v *AnalyzeAudioSyncRequestTranscription) *AnalyzeAudioSyncRequest
	GetTranscription() *AnalyzeAudioSyncRequestTranscription
	SetVariables(v []*AnalyzeAudioSyncRequestVariables) *AnalyzeAudioSyncRequest
	GetVariables() []*AnalyzeAudioSyncRequestVariables
}

type AnalyzeAudioSyncRequest struct {
	CategoryTags []*AnalyzeAudioSyncRequestCategoryTags `json:"categoryTags,omitempty" xml:"categoryTags,omitempty" type:"Repeated"`
	CustomPrompt *string                                `json:"customPrompt,omitempty" xml:"customPrompt,omitempty"`
	Fields       []*AnalyzeAudioSyncRequestFields       `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	// example:
	//
	// tyxmTurbo
	ModelCode *string `json:"modelCode,omitempty" xml:"modelCode,omitempty"`
	// example:
	//
	// jsonObject
	ResponseFormatType *string                                   `json:"responseFormatType,omitempty" xml:"responseFormatType,omitempty"`
	ResultTypes        []*string                                 `json:"resultTypes,omitempty" xml:"resultTypes,omitempty" type:"Repeated"`
	ServiceInspection  *AnalyzeAudioSyncRequestServiceInspection `json:"serviceInspection,omitempty" xml:"serviceInspection,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// false
	Stream        *bool                                 `json:"stream,omitempty" xml:"stream,omitempty"`
	TemplateIds   []*string                             `json:"templateIds,omitempty" xml:"templateIds,omitempty" type:"Repeated"`
	Transcription *AnalyzeAudioSyncRequestTranscription `json:"transcription,omitempty" xml:"transcription,omitempty" type:"Struct"`
	Variables     []*AnalyzeAudioSyncRequestVariables   `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
}

func (s AnalyzeAudioSyncRequest) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeAudioSyncRequest) GoString() string {
	return s.String()
}

func (s *AnalyzeAudioSyncRequest) GetCategoryTags() []*AnalyzeAudioSyncRequestCategoryTags {
	return s.CategoryTags
}

func (s *AnalyzeAudioSyncRequest) GetCustomPrompt() *string {
	return s.CustomPrompt
}

func (s *AnalyzeAudioSyncRequest) GetFields() []*AnalyzeAudioSyncRequestFields {
	return s.Fields
}

func (s *AnalyzeAudioSyncRequest) GetModelCode() *string {
	return s.ModelCode
}

func (s *AnalyzeAudioSyncRequest) GetResponseFormatType() *string {
	return s.ResponseFormatType
}

func (s *AnalyzeAudioSyncRequest) GetResultTypes() []*string {
	return s.ResultTypes
}

func (s *AnalyzeAudioSyncRequest) GetServiceInspection() *AnalyzeAudioSyncRequestServiceInspection {
	return s.ServiceInspection
}

func (s *AnalyzeAudioSyncRequest) GetStream() *bool {
	return s.Stream
}

func (s *AnalyzeAudioSyncRequest) GetTemplateIds() []*string {
	return s.TemplateIds
}

func (s *AnalyzeAudioSyncRequest) GetTranscription() *AnalyzeAudioSyncRequestTranscription {
	return s.Transcription
}

func (s *AnalyzeAudioSyncRequest) GetVariables() []*AnalyzeAudioSyncRequestVariables {
	return s.Variables
}

func (s *AnalyzeAudioSyncRequest) SetCategoryTags(v []*AnalyzeAudioSyncRequestCategoryTags) *AnalyzeAudioSyncRequest {
	s.CategoryTags = v
	return s
}

func (s *AnalyzeAudioSyncRequest) SetCustomPrompt(v string) *AnalyzeAudioSyncRequest {
	s.CustomPrompt = &v
	return s
}

func (s *AnalyzeAudioSyncRequest) SetFields(v []*AnalyzeAudioSyncRequestFields) *AnalyzeAudioSyncRequest {
	s.Fields = v
	return s
}

func (s *AnalyzeAudioSyncRequest) SetModelCode(v string) *AnalyzeAudioSyncRequest {
	s.ModelCode = &v
	return s
}

func (s *AnalyzeAudioSyncRequest) SetResponseFormatType(v string) *AnalyzeAudioSyncRequest {
	s.ResponseFormatType = &v
	return s
}

func (s *AnalyzeAudioSyncRequest) SetResultTypes(v []*string) *AnalyzeAudioSyncRequest {
	s.ResultTypes = v
	return s
}

func (s *AnalyzeAudioSyncRequest) SetServiceInspection(v *AnalyzeAudioSyncRequestServiceInspection) *AnalyzeAudioSyncRequest {
	s.ServiceInspection = v
	return s
}

func (s *AnalyzeAudioSyncRequest) SetStream(v bool) *AnalyzeAudioSyncRequest {
	s.Stream = &v
	return s
}

func (s *AnalyzeAudioSyncRequest) SetTemplateIds(v []*string) *AnalyzeAudioSyncRequest {
	s.TemplateIds = v
	return s
}

func (s *AnalyzeAudioSyncRequest) SetTranscription(v *AnalyzeAudioSyncRequestTranscription) *AnalyzeAudioSyncRequest {
	s.Transcription = v
	return s
}

func (s *AnalyzeAudioSyncRequest) SetVariables(v []*AnalyzeAudioSyncRequestVariables) *AnalyzeAudioSyncRequest {
	s.Variables = v
	return s
}

func (s *AnalyzeAudioSyncRequest) Validate() error {
	return dara.Validate(s)
}

type AnalyzeAudioSyncRequestCategoryTags struct {
	TagDesc *string `json:"tagDesc,omitempty" xml:"tagDesc,omitempty"`
	TagName *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
}

func (s AnalyzeAudioSyncRequestCategoryTags) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeAudioSyncRequestCategoryTags) GoString() string {
	return s.String()
}

func (s *AnalyzeAudioSyncRequestCategoryTags) GetTagDesc() *string {
	return s.TagDesc
}

func (s *AnalyzeAudioSyncRequestCategoryTags) GetTagName() *string {
	return s.TagName
}

func (s *AnalyzeAudioSyncRequestCategoryTags) SetTagDesc(v string) *AnalyzeAudioSyncRequestCategoryTags {
	s.TagDesc = &v
	return s
}

func (s *AnalyzeAudioSyncRequestCategoryTags) SetTagName(v string) *AnalyzeAudioSyncRequestCategoryTags {
	s.TagName = &v
	return s
}

func (s *AnalyzeAudioSyncRequestCategoryTags) Validate() error {
	return dara.Validate(s)
}

type AnalyzeAudioSyncRequestFields struct {
	// example:
	//
	// phoneNumber
	Code       *string                                    `json:"code,omitempty" xml:"code,omitempty"`
	Desc       *string                                    `json:"desc,omitempty" xml:"desc,omitempty"`
	EnumValues []*AnalyzeAudioSyncRequestFieldsEnumValues `json:"enumValues,omitempty" xml:"enumValues,omitempty" type:"Repeated"`
	Name       *string                                    `json:"name,omitempty" xml:"name,omitempty"`
}

func (s AnalyzeAudioSyncRequestFields) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeAudioSyncRequestFields) GoString() string {
	return s.String()
}

func (s *AnalyzeAudioSyncRequestFields) GetCode() *string {
	return s.Code
}

func (s *AnalyzeAudioSyncRequestFields) GetDesc() *string {
	return s.Desc
}

func (s *AnalyzeAudioSyncRequestFields) GetEnumValues() []*AnalyzeAudioSyncRequestFieldsEnumValues {
	return s.EnumValues
}

func (s *AnalyzeAudioSyncRequestFields) GetName() *string {
	return s.Name
}

func (s *AnalyzeAudioSyncRequestFields) SetCode(v string) *AnalyzeAudioSyncRequestFields {
	s.Code = &v
	return s
}

func (s *AnalyzeAudioSyncRequestFields) SetDesc(v string) *AnalyzeAudioSyncRequestFields {
	s.Desc = &v
	return s
}

func (s *AnalyzeAudioSyncRequestFields) SetEnumValues(v []*AnalyzeAudioSyncRequestFieldsEnumValues) *AnalyzeAudioSyncRequestFields {
	s.EnumValues = v
	return s
}

func (s *AnalyzeAudioSyncRequestFields) SetName(v string) *AnalyzeAudioSyncRequestFields {
	s.Name = &v
	return s
}

func (s *AnalyzeAudioSyncRequestFields) Validate() error {
	return dara.Validate(s)
}

type AnalyzeAudioSyncRequestFieldsEnumValues struct {
	Desc      *string `json:"desc,omitempty" xml:"desc,omitempty"`
	EnumValue *string `json:"enumValue,omitempty" xml:"enumValue,omitempty"`
}

func (s AnalyzeAudioSyncRequestFieldsEnumValues) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeAudioSyncRequestFieldsEnumValues) GoString() string {
	return s.String()
}

func (s *AnalyzeAudioSyncRequestFieldsEnumValues) GetDesc() *string {
	return s.Desc
}

func (s *AnalyzeAudioSyncRequestFieldsEnumValues) GetEnumValue() *string {
	return s.EnumValue
}

func (s *AnalyzeAudioSyncRequestFieldsEnumValues) SetDesc(v string) *AnalyzeAudioSyncRequestFieldsEnumValues {
	s.Desc = &v
	return s
}

func (s *AnalyzeAudioSyncRequestFieldsEnumValues) SetEnumValue(v string) *AnalyzeAudioSyncRequestFieldsEnumValues {
	s.EnumValue = &v
	return s
}

func (s *AnalyzeAudioSyncRequestFieldsEnumValues) Validate() error {
	return dara.Validate(s)
}

type AnalyzeAudioSyncRequestServiceInspection struct {
	InspectionContents     []*AnalyzeAudioSyncRequestServiceInspectionInspectionContents `json:"inspectionContents,omitempty" xml:"inspectionContents,omitempty" type:"Repeated"`
	InspectionIntroduction *string                                                       `json:"inspectionIntroduction,omitempty" xml:"inspectionIntroduction,omitempty"`
	SceneIntroduction      *string                                                       `json:"sceneIntroduction,omitempty" xml:"sceneIntroduction,omitempty"`
}

func (s AnalyzeAudioSyncRequestServiceInspection) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeAudioSyncRequestServiceInspection) GoString() string {
	return s.String()
}

func (s *AnalyzeAudioSyncRequestServiceInspection) GetInspectionContents() []*AnalyzeAudioSyncRequestServiceInspectionInspectionContents {
	return s.InspectionContents
}

func (s *AnalyzeAudioSyncRequestServiceInspection) GetInspectionIntroduction() *string {
	return s.InspectionIntroduction
}

func (s *AnalyzeAudioSyncRequestServiceInspection) GetSceneIntroduction() *string {
	return s.SceneIntroduction
}

func (s *AnalyzeAudioSyncRequestServiceInspection) SetInspectionContents(v []*AnalyzeAudioSyncRequestServiceInspectionInspectionContents) *AnalyzeAudioSyncRequestServiceInspection {
	s.InspectionContents = v
	return s
}

func (s *AnalyzeAudioSyncRequestServiceInspection) SetInspectionIntroduction(v string) *AnalyzeAudioSyncRequestServiceInspection {
	s.InspectionIntroduction = &v
	return s
}

func (s *AnalyzeAudioSyncRequestServiceInspection) SetSceneIntroduction(v string) *AnalyzeAudioSyncRequestServiceInspection {
	s.SceneIntroduction = &v
	return s
}

func (s *AnalyzeAudioSyncRequestServiceInspection) Validate() error {
	return dara.Validate(s)
}

type AnalyzeAudioSyncRequestServiceInspectionInspectionContents struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	Title   *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s AnalyzeAudioSyncRequestServiceInspectionInspectionContents) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeAudioSyncRequestServiceInspectionInspectionContents) GoString() string {
	return s.String()
}

func (s *AnalyzeAudioSyncRequestServiceInspectionInspectionContents) GetContent() *string {
	return s.Content
}

func (s *AnalyzeAudioSyncRequestServiceInspectionInspectionContents) GetTitle() *string {
	return s.Title
}

func (s *AnalyzeAudioSyncRequestServiceInspectionInspectionContents) SetContent(v string) *AnalyzeAudioSyncRequestServiceInspectionInspectionContents {
	s.Content = &v
	return s
}

func (s *AnalyzeAudioSyncRequestServiceInspectionInspectionContents) SetTitle(v string) *AnalyzeAudioSyncRequestServiceInspectionInspectionContents {
	s.Title = &v
	return s
}

func (s *AnalyzeAudioSyncRequestServiceInspectionInspectionContents) Validate() error {
	return dara.Validate(s)
}

type AnalyzeAudioSyncRequestTranscription struct {
	// example:
	//
	// nls
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
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// example:
	//
	// low
	Level *string `json:"level,omitempty" xml:"level,omitempty"`
	// example:
	//
	// 1
	ServiceChannel         *int32    `json:"serviceChannel,omitempty" xml:"serviceChannel,omitempty"`
	ServiceChannelKeywords []*string `json:"serviceChannelKeywords,omitempty" xml:"serviceChannelKeywords,omitempty" type:"Repeated"`
	// example:
	//
	// esnvknv*****skdnvjksd
	VocabularyId *string `json:"vocabularyId,omitempty" xml:"vocabularyId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://1111.com/sss.mp3
	VoiceFileUrl *string `json:"voiceFileUrl,omitempty" xml:"voiceFileUrl,omitempty"`
}

func (s AnalyzeAudioSyncRequestTranscription) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeAudioSyncRequestTranscription) GoString() string {
	return s.String()
}

func (s *AnalyzeAudioSyncRequestTranscription) GetAsrModelCode() *string {
	return s.AsrModelCode
}

func (s *AnalyzeAudioSyncRequestTranscription) GetAutoSplit() *int32 {
	return s.AutoSplit
}

func (s *AnalyzeAudioSyncRequestTranscription) GetClientChannel() *int32 {
	return s.ClientChannel
}

func (s *AnalyzeAudioSyncRequestTranscription) GetFileName() *string {
	return s.FileName
}

func (s *AnalyzeAudioSyncRequestTranscription) GetLevel() *string {
	return s.Level
}

func (s *AnalyzeAudioSyncRequestTranscription) GetServiceChannel() *int32 {
	return s.ServiceChannel
}

func (s *AnalyzeAudioSyncRequestTranscription) GetServiceChannelKeywords() []*string {
	return s.ServiceChannelKeywords
}

func (s *AnalyzeAudioSyncRequestTranscription) GetVocabularyId() *string {
	return s.VocabularyId
}

func (s *AnalyzeAudioSyncRequestTranscription) GetVoiceFileUrl() *string {
	return s.VoiceFileUrl
}

func (s *AnalyzeAudioSyncRequestTranscription) SetAsrModelCode(v string) *AnalyzeAudioSyncRequestTranscription {
	s.AsrModelCode = &v
	return s
}

func (s *AnalyzeAudioSyncRequestTranscription) SetAutoSplit(v int32) *AnalyzeAudioSyncRequestTranscription {
	s.AutoSplit = &v
	return s
}

func (s *AnalyzeAudioSyncRequestTranscription) SetClientChannel(v int32) *AnalyzeAudioSyncRequestTranscription {
	s.ClientChannel = &v
	return s
}

func (s *AnalyzeAudioSyncRequestTranscription) SetFileName(v string) *AnalyzeAudioSyncRequestTranscription {
	s.FileName = &v
	return s
}

func (s *AnalyzeAudioSyncRequestTranscription) SetLevel(v string) *AnalyzeAudioSyncRequestTranscription {
	s.Level = &v
	return s
}

func (s *AnalyzeAudioSyncRequestTranscription) SetServiceChannel(v int32) *AnalyzeAudioSyncRequestTranscription {
	s.ServiceChannel = &v
	return s
}

func (s *AnalyzeAudioSyncRequestTranscription) SetServiceChannelKeywords(v []*string) *AnalyzeAudioSyncRequestTranscription {
	s.ServiceChannelKeywords = v
	return s
}

func (s *AnalyzeAudioSyncRequestTranscription) SetVocabularyId(v string) *AnalyzeAudioSyncRequestTranscription {
	s.VocabularyId = &v
	return s
}

func (s *AnalyzeAudioSyncRequestTranscription) SetVoiceFileUrl(v string) *AnalyzeAudioSyncRequestTranscription {
	s.VoiceFileUrl = &v
	return s
}

func (s *AnalyzeAudioSyncRequestTranscription) Validate() error {
	return dara.Validate(s)
}

type AnalyzeAudioSyncRequestVariables struct {
	// example:
	//
	// name
	VariableCode  *string `json:"variableCode,omitempty" xml:"variableCode,omitempty"`
	VariableValue *string `json:"variableValue,omitempty" xml:"variableValue,omitempty"`
}

func (s AnalyzeAudioSyncRequestVariables) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeAudioSyncRequestVariables) GoString() string {
	return s.String()
}

func (s *AnalyzeAudioSyncRequestVariables) GetVariableCode() *string {
	return s.VariableCode
}

func (s *AnalyzeAudioSyncRequestVariables) GetVariableValue() *string {
	return s.VariableValue
}

func (s *AnalyzeAudioSyncRequestVariables) SetVariableCode(v string) *AnalyzeAudioSyncRequestVariables {
	s.VariableCode = &v
	return s
}

func (s *AnalyzeAudioSyncRequestVariables) SetVariableValue(v string) *AnalyzeAudioSyncRequestVariables {
	s.VariableValue = &v
	return s
}

func (s *AnalyzeAudioSyncRequestVariables) Validate() error {
	return dara.Validate(s)
}
