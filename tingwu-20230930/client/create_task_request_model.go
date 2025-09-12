// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v string) *CreateTaskRequest
	GetAppKey() *string
	SetInput(v *CreateTaskRequestInput) *CreateTaskRequest
	GetInput() *CreateTaskRequestInput
	SetParameters(v *CreateTaskRequestParameters) *CreateTaskRequest
	GetParameters() *CreateTaskRequestParameters
	SetOperation(v string) *CreateTaskRequest
	GetOperation() *string
	SetType(v string) *CreateTaskRequest
	GetType() *string
}

type CreateTaskRequest struct {
	// example:
	//
	// JV1sRTisRMi****
	AppKey     *string                      `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	Input      *CreateTaskRequestInput      `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	Parameters *CreateTaskRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Struct"`
	// example:
	//
	// stop
	Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// offline
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateTaskRequest) GetAppKey() *string {
	return s.AppKey
}

func (s *CreateTaskRequest) GetInput() *CreateTaskRequestInput {
	return s.Input
}

func (s *CreateTaskRequest) GetParameters() *CreateTaskRequestParameters {
	return s.Parameters
}

func (s *CreateTaskRequest) GetOperation() *string {
	return s.Operation
}

func (s *CreateTaskRequest) GetType() *string {
	return s.Type
}

func (s *CreateTaskRequest) SetAppKey(v string) *CreateTaskRequest {
	s.AppKey = &v
	return s
}

func (s *CreateTaskRequest) SetInput(v *CreateTaskRequestInput) *CreateTaskRequest {
	s.Input = v
	return s
}

func (s *CreateTaskRequest) SetParameters(v *CreateTaskRequestParameters) *CreateTaskRequest {
	s.Parameters = v
	return s
}

func (s *CreateTaskRequest) SetOperation(v string) *CreateTaskRequest {
	s.Operation = &v
	return s
}

func (s *CreateTaskRequest) SetType(v string) *CreateTaskRequest {
	s.Type = &v
	return s
}

func (s *CreateTaskRequest) Validate() error {
	return dara.Validate(s)
}

type CreateTaskRequestInput struct {
	AudioChannelMode *string `json:"AudioChannelMode,omitempty" xml:"AudioChannelMode,omitempty"`
	// example:
	//
	// http://xxx.com/zzz/1.wav
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// pcm
	Format                      *string   `json:"Format,omitempty" xml:"Format,omitempty"`
	LanguageHints               []*string `json:"LanguageHints,omitempty" xml:"LanguageHints,omitempty" type:"Repeated"`
	MultipleStreamsEnabled      *bool     `json:"MultipleStreamsEnabled,omitempty" xml:"MultipleStreamsEnabled,omitempty"`
	OutputPath                  *string   `json:"OutputPath,omitempty" xml:"OutputPath,omitempty"`
	ProgressiveCallbacksEnabled *bool     `json:"ProgressiveCallbacksEnabled,omitempty" xml:"ProgressiveCallbacksEnabled,omitempty"`
	// example:
	//
	// 16000
	SampleRate *int32 `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	TaskId         *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// task_tingwu_123
	TaskKey *string `json:"TaskKey,omitempty" xml:"TaskKey,omitempty"`
}

func (s CreateTaskRequestInput) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequestInput) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestInput) GetAudioChannelMode() *string {
	return s.AudioChannelMode
}

func (s *CreateTaskRequestInput) GetFileUrl() *string {
	return s.FileUrl
}

func (s *CreateTaskRequestInput) GetFormat() *string {
	return s.Format
}

func (s *CreateTaskRequestInput) GetLanguageHints() []*string {
	return s.LanguageHints
}

func (s *CreateTaskRequestInput) GetMultipleStreamsEnabled() *bool {
	return s.MultipleStreamsEnabled
}

func (s *CreateTaskRequestInput) GetOutputPath() *string {
	return s.OutputPath
}

func (s *CreateTaskRequestInput) GetProgressiveCallbacksEnabled() *bool {
	return s.ProgressiveCallbacksEnabled
}

func (s *CreateTaskRequestInput) GetSampleRate() *int32 {
	return s.SampleRate
}

func (s *CreateTaskRequestInput) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *CreateTaskRequestInput) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateTaskRequestInput) GetTaskKey() *string {
	return s.TaskKey
}

func (s *CreateTaskRequestInput) SetAudioChannelMode(v string) *CreateTaskRequestInput {
	s.AudioChannelMode = &v
	return s
}

func (s *CreateTaskRequestInput) SetFileUrl(v string) *CreateTaskRequestInput {
	s.FileUrl = &v
	return s
}

func (s *CreateTaskRequestInput) SetFormat(v string) *CreateTaskRequestInput {
	s.Format = &v
	return s
}

func (s *CreateTaskRequestInput) SetLanguageHints(v []*string) *CreateTaskRequestInput {
	s.LanguageHints = v
	return s
}

func (s *CreateTaskRequestInput) SetMultipleStreamsEnabled(v bool) *CreateTaskRequestInput {
	s.MultipleStreamsEnabled = &v
	return s
}

func (s *CreateTaskRequestInput) SetOutputPath(v string) *CreateTaskRequestInput {
	s.OutputPath = &v
	return s
}

func (s *CreateTaskRequestInput) SetProgressiveCallbacksEnabled(v bool) *CreateTaskRequestInput {
	s.ProgressiveCallbacksEnabled = &v
	return s
}

func (s *CreateTaskRequestInput) SetSampleRate(v int32) *CreateTaskRequestInput {
	s.SampleRate = &v
	return s
}

func (s *CreateTaskRequestInput) SetSourceLanguage(v string) *CreateTaskRequestInput {
	s.SourceLanguage = &v
	return s
}

func (s *CreateTaskRequestInput) SetTaskId(v string) *CreateTaskRequestInput {
	s.TaskId = &v
	return s
}

func (s *CreateTaskRequestInput) SetTaskKey(v string) *CreateTaskRequestInput {
	s.TaskKey = &v
	return s
}

func (s *CreateTaskRequestInput) Validate() error {
	return dara.Validate(s)
}

type CreateTaskRequestParameters struct {
	AutoChapters *CreateTaskRequestParametersAutoChapters `json:"AutoChapters,omitempty" xml:"AutoChapters,omitempty" type:"Struct"`
	// example:
	//
	// true
	AutoChaptersEnabled        *bool                                           `json:"AutoChaptersEnabled,omitempty" xml:"AutoChaptersEnabled,omitempty"`
	ContentExtraction          *CreateTaskRequestParametersContentExtraction   `json:"ContentExtraction,omitempty" xml:"ContentExtraction,omitempty" type:"Struct"`
	ContentExtractionEnabled   *bool                                           `json:"ContentExtractionEnabled,omitempty" xml:"ContentExtractionEnabled,omitempty"`
	CustomPrompt               *CreateTaskRequestParametersCustomPrompt        `json:"CustomPrompt,omitempty" xml:"CustomPrompt,omitempty" type:"Struct"`
	CustomPromptEnabled        *bool                                           `json:"CustomPromptEnabled,omitempty" xml:"CustomPromptEnabled,omitempty"`
	ExtraParams                *CreateTaskRequestParametersExtraParams         `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty" type:"Struct"`
	IdentityRecognition        *CreateTaskRequestParametersIdentityRecognition `json:"IdentityRecognition,omitempty" xml:"IdentityRecognition,omitempty" type:"Struct"`
	IdentityRecognitionEnabled *bool                                           `json:"IdentityRecognitionEnabled,omitempty" xml:"IdentityRecognitionEnabled,omitempty"`
	LlmOutputLanguage          *string                                         `json:"LlmOutputLanguage,omitempty" xml:"LlmOutputLanguage,omitempty"`
	MeetingAssistance          *CreateTaskRequestParametersMeetingAssistance   `json:"MeetingAssistance,omitempty" xml:"MeetingAssistance,omitempty" type:"Struct"`
	// example:
	//
	// false
	MeetingAssistanceEnabled *bool                                         `json:"MeetingAssistanceEnabled,omitempty" xml:"MeetingAssistanceEnabled,omitempty"`
	Model                    *string                                       `json:"Model,omitempty" xml:"Model,omitempty"`
	PptExtractionEnabled     *bool                                         `json:"PptExtractionEnabled,omitempty" xml:"PptExtractionEnabled,omitempty"`
	ServiceInspection        *CreateTaskRequestParametersServiceInspection `json:"ServiceInspection,omitempty" xml:"ServiceInspection,omitempty" type:"Struct"`
	ServiceInspectionEnabled *bool                                         `json:"ServiceInspectionEnabled,omitempty" xml:"ServiceInspectionEnabled,omitempty"`
	Summarization            *CreateTaskRequestParametersSummarization     `json:"Summarization,omitempty" xml:"Summarization,omitempty" type:"Struct"`
	// example:
	//
	// false
	SummarizationEnabled *bool                                     `json:"SummarizationEnabled,omitempty" xml:"SummarizationEnabled,omitempty"`
	TextPolishEnabled    *bool                                     `json:"TextPolishEnabled,omitempty" xml:"TextPolishEnabled,omitempty"`
	Transcoding          *CreateTaskRequestParametersTranscoding   `json:"Transcoding,omitempty" xml:"Transcoding,omitempty" type:"Struct"`
	Transcription        *CreateTaskRequestParametersTranscription `json:"Transcription,omitempty" xml:"Transcription,omitempty" type:"Struct"`
	Translation          *CreateTaskRequestParametersTranslation   `json:"Translation,omitempty" xml:"Translation,omitempty" type:"Struct"`
	// example:
	//
	// false
	TranslationEnabled *bool `json:"TranslationEnabled,omitempty" xml:"TranslationEnabled,omitempty"`
}

func (s CreateTaskRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequestParameters) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestParameters) GetAutoChapters() *CreateTaskRequestParametersAutoChapters {
	return s.AutoChapters
}

func (s *CreateTaskRequestParameters) GetAutoChaptersEnabled() *bool {
	return s.AutoChaptersEnabled
}

func (s *CreateTaskRequestParameters) GetContentExtraction() *CreateTaskRequestParametersContentExtraction {
	return s.ContentExtraction
}

func (s *CreateTaskRequestParameters) GetContentExtractionEnabled() *bool {
	return s.ContentExtractionEnabled
}

func (s *CreateTaskRequestParameters) GetCustomPrompt() *CreateTaskRequestParametersCustomPrompt {
	return s.CustomPrompt
}

func (s *CreateTaskRequestParameters) GetCustomPromptEnabled() *bool {
	return s.CustomPromptEnabled
}

func (s *CreateTaskRequestParameters) GetExtraParams() *CreateTaskRequestParametersExtraParams {
	return s.ExtraParams
}

func (s *CreateTaskRequestParameters) GetIdentityRecognition() *CreateTaskRequestParametersIdentityRecognition {
	return s.IdentityRecognition
}

func (s *CreateTaskRequestParameters) GetIdentityRecognitionEnabled() *bool {
	return s.IdentityRecognitionEnabled
}

func (s *CreateTaskRequestParameters) GetLlmOutputLanguage() *string {
	return s.LlmOutputLanguage
}

func (s *CreateTaskRequestParameters) GetMeetingAssistance() *CreateTaskRequestParametersMeetingAssistance {
	return s.MeetingAssistance
}

func (s *CreateTaskRequestParameters) GetMeetingAssistanceEnabled() *bool {
	return s.MeetingAssistanceEnabled
}

func (s *CreateTaskRequestParameters) GetModel() *string {
	return s.Model
}

func (s *CreateTaskRequestParameters) GetPptExtractionEnabled() *bool {
	return s.PptExtractionEnabled
}

func (s *CreateTaskRequestParameters) GetServiceInspection() *CreateTaskRequestParametersServiceInspection {
	return s.ServiceInspection
}

func (s *CreateTaskRequestParameters) GetServiceInspectionEnabled() *bool {
	return s.ServiceInspectionEnabled
}

func (s *CreateTaskRequestParameters) GetSummarization() *CreateTaskRequestParametersSummarization {
	return s.Summarization
}

func (s *CreateTaskRequestParameters) GetSummarizationEnabled() *bool {
	return s.SummarizationEnabled
}

func (s *CreateTaskRequestParameters) GetTextPolishEnabled() *bool {
	return s.TextPolishEnabled
}

func (s *CreateTaskRequestParameters) GetTranscoding() *CreateTaskRequestParametersTranscoding {
	return s.Transcoding
}

func (s *CreateTaskRequestParameters) GetTranscription() *CreateTaskRequestParametersTranscription {
	return s.Transcription
}

func (s *CreateTaskRequestParameters) GetTranslation() *CreateTaskRequestParametersTranslation {
	return s.Translation
}

func (s *CreateTaskRequestParameters) GetTranslationEnabled() *bool {
	return s.TranslationEnabled
}

func (s *CreateTaskRequestParameters) SetAutoChapters(v *CreateTaskRequestParametersAutoChapters) *CreateTaskRequestParameters {
	s.AutoChapters = v
	return s
}

func (s *CreateTaskRequestParameters) SetAutoChaptersEnabled(v bool) *CreateTaskRequestParameters {
	s.AutoChaptersEnabled = &v
	return s
}

func (s *CreateTaskRequestParameters) SetContentExtraction(v *CreateTaskRequestParametersContentExtraction) *CreateTaskRequestParameters {
	s.ContentExtraction = v
	return s
}

func (s *CreateTaskRequestParameters) SetContentExtractionEnabled(v bool) *CreateTaskRequestParameters {
	s.ContentExtractionEnabled = &v
	return s
}

func (s *CreateTaskRequestParameters) SetCustomPrompt(v *CreateTaskRequestParametersCustomPrompt) *CreateTaskRequestParameters {
	s.CustomPrompt = v
	return s
}

func (s *CreateTaskRequestParameters) SetCustomPromptEnabled(v bool) *CreateTaskRequestParameters {
	s.CustomPromptEnabled = &v
	return s
}

func (s *CreateTaskRequestParameters) SetExtraParams(v *CreateTaskRequestParametersExtraParams) *CreateTaskRequestParameters {
	s.ExtraParams = v
	return s
}

func (s *CreateTaskRequestParameters) SetIdentityRecognition(v *CreateTaskRequestParametersIdentityRecognition) *CreateTaskRequestParameters {
	s.IdentityRecognition = v
	return s
}

func (s *CreateTaskRequestParameters) SetIdentityRecognitionEnabled(v bool) *CreateTaskRequestParameters {
	s.IdentityRecognitionEnabled = &v
	return s
}

func (s *CreateTaskRequestParameters) SetLlmOutputLanguage(v string) *CreateTaskRequestParameters {
	s.LlmOutputLanguage = &v
	return s
}

func (s *CreateTaskRequestParameters) SetMeetingAssistance(v *CreateTaskRequestParametersMeetingAssistance) *CreateTaskRequestParameters {
	s.MeetingAssistance = v
	return s
}

func (s *CreateTaskRequestParameters) SetMeetingAssistanceEnabled(v bool) *CreateTaskRequestParameters {
	s.MeetingAssistanceEnabled = &v
	return s
}

func (s *CreateTaskRequestParameters) SetModel(v string) *CreateTaskRequestParameters {
	s.Model = &v
	return s
}

func (s *CreateTaskRequestParameters) SetPptExtractionEnabled(v bool) *CreateTaskRequestParameters {
	s.PptExtractionEnabled = &v
	return s
}

func (s *CreateTaskRequestParameters) SetServiceInspection(v *CreateTaskRequestParametersServiceInspection) *CreateTaskRequestParameters {
	s.ServiceInspection = v
	return s
}

func (s *CreateTaskRequestParameters) SetServiceInspectionEnabled(v bool) *CreateTaskRequestParameters {
	s.ServiceInspectionEnabled = &v
	return s
}

func (s *CreateTaskRequestParameters) SetSummarization(v *CreateTaskRequestParametersSummarization) *CreateTaskRequestParameters {
	s.Summarization = v
	return s
}

func (s *CreateTaskRequestParameters) SetSummarizationEnabled(v bool) *CreateTaskRequestParameters {
	s.SummarizationEnabled = &v
	return s
}

func (s *CreateTaskRequestParameters) SetTextPolishEnabled(v bool) *CreateTaskRequestParameters {
	s.TextPolishEnabled = &v
	return s
}

func (s *CreateTaskRequestParameters) SetTranscoding(v *CreateTaskRequestParametersTranscoding) *CreateTaskRequestParameters {
	s.Transcoding = v
	return s
}

func (s *CreateTaskRequestParameters) SetTranscription(v *CreateTaskRequestParametersTranscription) *CreateTaskRequestParameters {
	s.Transcription = v
	return s
}

func (s *CreateTaskRequestParameters) SetTranslation(v *CreateTaskRequestParametersTranslation) *CreateTaskRequestParameters {
	s.Translation = v
	return s
}

func (s *CreateTaskRequestParameters) SetTranslationEnabled(v bool) *CreateTaskRequestParameters {
	s.TranslationEnabled = &v
	return s
}

func (s *CreateTaskRequestParameters) Validate() error {
	return dara.Validate(s)
}

type CreateTaskRequestParametersAutoChapters struct {
	ChapterGranularity *string `json:"ChapterGranularity,omitempty" xml:"ChapterGranularity,omitempty"`
}

func (s CreateTaskRequestParametersAutoChapters) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequestParametersAutoChapters) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestParametersAutoChapters) GetChapterGranularity() *string {
	return s.ChapterGranularity
}

func (s *CreateTaskRequestParametersAutoChapters) SetChapterGranularity(v string) *CreateTaskRequestParametersAutoChapters {
	s.ChapterGranularity = &v
	return s
}

func (s *CreateTaskRequestParametersAutoChapters) Validate() error {
	return dara.Validate(s)
}

type CreateTaskRequestParametersContentExtraction struct {
	ExtractionContents []*CreateTaskRequestParametersContentExtractionExtractionContents `json:"ExtractionContents,omitempty" xml:"ExtractionContents,omitempty" type:"Repeated"`
	SceneIntroduction  *string                                                           `json:"SceneIntroduction,omitempty" xml:"SceneIntroduction,omitempty"`
	SpeakerMap         map[string]interface{}                                            `json:"SpeakerMap,omitempty" xml:"SpeakerMap,omitempty"`
}

func (s CreateTaskRequestParametersContentExtraction) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequestParametersContentExtraction) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestParametersContentExtraction) GetExtractionContents() []*CreateTaskRequestParametersContentExtractionExtractionContents {
	return s.ExtractionContents
}

func (s *CreateTaskRequestParametersContentExtraction) GetSceneIntroduction() *string {
	return s.SceneIntroduction
}

func (s *CreateTaskRequestParametersContentExtraction) GetSpeakerMap() map[string]interface{} {
	return s.SpeakerMap
}

func (s *CreateTaskRequestParametersContentExtraction) SetExtractionContents(v []*CreateTaskRequestParametersContentExtractionExtractionContents) *CreateTaskRequestParametersContentExtraction {
	s.ExtractionContents = v
	return s
}

func (s *CreateTaskRequestParametersContentExtraction) SetSceneIntroduction(v string) *CreateTaskRequestParametersContentExtraction {
	s.SceneIntroduction = &v
	return s
}

func (s *CreateTaskRequestParametersContentExtraction) SetSpeakerMap(v map[string]interface{}) *CreateTaskRequestParametersContentExtraction {
	s.SpeakerMap = v
	return s
}

func (s *CreateTaskRequestParametersContentExtraction) Validate() error {
	return dara.Validate(s)
}

type CreateTaskRequestParametersContentExtractionExtractionContents struct {
	Content  *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Identity *string `json:"Identity,omitempty" xml:"Identity,omitempty"`
	Title    *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateTaskRequestParametersContentExtractionExtractionContents) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequestParametersContentExtractionExtractionContents) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestParametersContentExtractionExtractionContents) GetContent() *string {
	return s.Content
}

func (s *CreateTaskRequestParametersContentExtractionExtractionContents) GetIdentity() *string {
	return s.Identity
}

func (s *CreateTaskRequestParametersContentExtractionExtractionContents) GetTitle() *string {
	return s.Title
}

func (s *CreateTaskRequestParametersContentExtractionExtractionContents) SetContent(v string) *CreateTaskRequestParametersContentExtractionExtractionContents {
	s.Content = &v
	return s
}

func (s *CreateTaskRequestParametersContentExtractionExtractionContents) SetIdentity(v string) *CreateTaskRequestParametersContentExtractionExtractionContents {
	s.Identity = &v
	return s
}

func (s *CreateTaskRequestParametersContentExtractionExtractionContents) SetTitle(v string) *CreateTaskRequestParametersContentExtractionExtractionContents {
	s.Title = &v
	return s
}

func (s *CreateTaskRequestParametersContentExtractionExtractionContents) Validate() error {
	return dara.Validate(s)
}

type CreateTaskRequestParametersCustomPrompt struct {
	Contents []*CreateTaskRequestParametersCustomPromptContents `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
}

func (s CreateTaskRequestParametersCustomPrompt) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequestParametersCustomPrompt) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestParametersCustomPrompt) GetContents() []*CreateTaskRequestParametersCustomPromptContents {
	return s.Contents
}

func (s *CreateTaskRequestParametersCustomPrompt) SetContents(v []*CreateTaskRequestParametersCustomPromptContents) *CreateTaskRequestParametersCustomPrompt {
	s.Contents = v
	return s
}

func (s *CreateTaskRequestParametersCustomPrompt) Validate() error {
	return dara.Validate(s)
}

type CreateTaskRequestParametersCustomPromptContents struct {
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	Prompt    *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	TransType *string `json:"TransType,omitempty" xml:"TransType,omitempty"`
}

func (s CreateTaskRequestParametersCustomPromptContents) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequestParametersCustomPromptContents) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestParametersCustomPromptContents) GetModel() *string {
	return s.Model
}

func (s *CreateTaskRequestParametersCustomPromptContents) GetName() *string {
	return s.Name
}

func (s *CreateTaskRequestParametersCustomPromptContents) GetPrompt() *string {
	return s.Prompt
}

func (s *CreateTaskRequestParametersCustomPromptContents) GetTransType() *string {
	return s.TransType
}

func (s *CreateTaskRequestParametersCustomPromptContents) SetModel(v string) *CreateTaskRequestParametersCustomPromptContents {
	s.Model = &v
	return s
}

func (s *CreateTaskRequestParametersCustomPromptContents) SetName(v string) *CreateTaskRequestParametersCustomPromptContents {
	s.Name = &v
	return s
}

func (s *CreateTaskRequestParametersCustomPromptContents) SetPrompt(v string) *CreateTaskRequestParametersCustomPromptContents {
	s.Prompt = &v
	return s
}

func (s *CreateTaskRequestParametersCustomPromptContents) SetTransType(v string) *CreateTaskRequestParametersCustomPromptContents {
	s.TransType = &v
	return s
}

func (s *CreateTaskRequestParametersCustomPromptContents) Validate() error {
	return dara.Validate(s)
}

type CreateTaskRequestParametersExtraParams struct {
	DomainEducationEnabled   *bool   `json:"DomainEducationEnabled,omitempty" xml:"DomainEducationEnabled,omitempty"`
	FullTextSummaryFormat    *string `json:"FullTextSummaryFormat,omitempty" xml:"FullTextSummaryFormat,omitempty"`
	MaxKeywords              *int32  `json:"MaxKeywords,omitempty" xml:"MaxKeywords,omitempty"`
	NfixEnabled              *bool   `json:"NfixEnabled,omitempty" xml:"NfixEnabled,omitempty"`
	OcrAuxiliaryEnabled      *bool   `json:"OcrAuxiliaryEnabled,omitempty" xml:"OcrAuxiliaryEnabled,omitempty"`
	TranslateLlmSceneEnabled *bool   `json:"TranslateLlmSceneEnabled,omitempty" xml:"TranslateLlmSceneEnabled,omitempty"`
}

func (s CreateTaskRequestParametersExtraParams) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequestParametersExtraParams) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestParametersExtraParams) GetDomainEducationEnabled() *bool {
	return s.DomainEducationEnabled
}

func (s *CreateTaskRequestParametersExtraParams) GetFullTextSummaryFormat() *string {
	return s.FullTextSummaryFormat
}

func (s *CreateTaskRequestParametersExtraParams) GetMaxKeywords() *int32 {
	return s.MaxKeywords
}

func (s *CreateTaskRequestParametersExtraParams) GetNfixEnabled() *bool {
	return s.NfixEnabled
}

func (s *CreateTaskRequestParametersExtraParams) GetOcrAuxiliaryEnabled() *bool {
	return s.OcrAuxiliaryEnabled
}

func (s *CreateTaskRequestParametersExtraParams) GetTranslateLlmSceneEnabled() *bool {
	return s.TranslateLlmSceneEnabled
}

func (s *CreateTaskRequestParametersExtraParams) SetDomainEducationEnabled(v bool) *CreateTaskRequestParametersExtraParams {
	s.DomainEducationEnabled = &v
	return s
}

func (s *CreateTaskRequestParametersExtraParams) SetFullTextSummaryFormat(v string) *CreateTaskRequestParametersExtraParams {
	s.FullTextSummaryFormat = &v
	return s
}

func (s *CreateTaskRequestParametersExtraParams) SetMaxKeywords(v int32) *CreateTaskRequestParametersExtraParams {
	s.MaxKeywords = &v
	return s
}

func (s *CreateTaskRequestParametersExtraParams) SetNfixEnabled(v bool) *CreateTaskRequestParametersExtraParams {
	s.NfixEnabled = &v
	return s
}

func (s *CreateTaskRequestParametersExtraParams) SetOcrAuxiliaryEnabled(v bool) *CreateTaskRequestParametersExtraParams {
	s.OcrAuxiliaryEnabled = &v
	return s
}

func (s *CreateTaskRequestParametersExtraParams) SetTranslateLlmSceneEnabled(v bool) *CreateTaskRequestParametersExtraParams {
	s.TranslateLlmSceneEnabled = &v
	return s
}

func (s *CreateTaskRequestParametersExtraParams) Validate() error {
	return dara.Validate(s)
}

type CreateTaskRequestParametersIdentityRecognition struct {
	IdentityContents  []*CreateTaskRequestParametersIdentityRecognitionIdentityContents `json:"IdentityContents,omitempty" xml:"IdentityContents,omitempty" type:"Repeated"`
	SceneIntroduction *string                                                           `json:"SceneIntroduction,omitempty" xml:"SceneIntroduction,omitempty"`
}

func (s CreateTaskRequestParametersIdentityRecognition) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequestParametersIdentityRecognition) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestParametersIdentityRecognition) GetIdentityContents() []*CreateTaskRequestParametersIdentityRecognitionIdentityContents {
	return s.IdentityContents
}

func (s *CreateTaskRequestParametersIdentityRecognition) GetSceneIntroduction() *string {
	return s.SceneIntroduction
}

func (s *CreateTaskRequestParametersIdentityRecognition) SetIdentityContents(v []*CreateTaskRequestParametersIdentityRecognitionIdentityContents) *CreateTaskRequestParametersIdentityRecognition {
	s.IdentityContents = v
	return s
}

func (s *CreateTaskRequestParametersIdentityRecognition) SetSceneIntroduction(v string) *CreateTaskRequestParametersIdentityRecognition {
	s.SceneIntroduction = &v
	return s
}

func (s *CreateTaskRequestParametersIdentityRecognition) Validate() error {
	return dara.Validate(s)
}

type CreateTaskRequestParametersIdentityRecognitionIdentityContents struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateTaskRequestParametersIdentityRecognitionIdentityContents) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequestParametersIdentityRecognitionIdentityContents) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestParametersIdentityRecognitionIdentityContents) GetDescription() *string {
	return s.Description
}

func (s *CreateTaskRequestParametersIdentityRecognitionIdentityContents) GetName() *string {
	return s.Name
}

func (s *CreateTaskRequestParametersIdentityRecognitionIdentityContents) SetDescription(v string) *CreateTaskRequestParametersIdentityRecognitionIdentityContents {
	s.Description = &v
	return s
}

func (s *CreateTaskRequestParametersIdentityRecognitionIdentityContents) SetName(v string) *CreateTaskRequestParametersIdentityRecognitionIdentityContents {
	s.Name = &v
	return s
}

func (s *CreateTaskRequestParametersIdentityRecognitionIdentityContents) Validate() error {
	return dara.Validate(s)
}

type CreateTaskRequestParametersMeetingAssistance struct {
	Types []*string `json:"Types,omitempty" xml:"Types,omitempty" type:"Repeated"`
}

func (s CreateTaskRequestParametersMeetingAssistance) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequestParametersMeetingAssistance) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestParametersMeetingAssistance) GetTypes() []*string {
	return s.Types
}

func (s *CreateTaskRequestParametersMeetingAssistance) SetTypes(v []*string) *CreateTaskRequestParametersMeetingAssistance {
	s.Types = v
	return s
}

func (s *CreateTaskRequestParametersMeetingAssistance) Validate() error {
	return dara.Validate(s)
}

type CreateTaskRequestParametersServiceInspection struct {
	InspectionContents     []*CreateTaskRequestParametersServiceInspectionInspectionContents `json:"InspectionContents,omitempty" xml:"InspectionContents,omitempty" type:"Repeated"`
	InspectionIntroduction *string                                                           `json:"InspectionIntroduction,omitempty" xml:"InspectionIntroduction,omitempty"`
	SceneIntroduction      *string                                                           `json:"SceneIntroduction,omitempty" xml:"SceneIntroduction,omitempty"`
	SpeakerMap             map[string]interface{}                                            `json:"SpeakerMap,omitempty" xml:"SpeakerMap,omitempty"`
}

func (s CreateTaskRequestParametersServiceInspection) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequestParametersServiceInspection) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestParametersServiceInspection) GetInspectionContents() []*CreateTaskRequestParametersServiceInspectionInspectionContents {
	return s.InspectionContents
}

func (s *CreateTaskRequestParametersServiceInspection) GetInspectionIntroduction() *string {
	return s.InspectionIntroduction
}

func (s *CreateTaskRequestParametersServiceInspection) GetSceneIntroduction() *string {
	return s.SceneIntroduction
}

func (s *CreateTaskRequestParametersServiceInspection) GetSpeakerMap() map[string]interface{} {
	return s.SpeakerMap
}

func (s *CreateTaskRequestParametersServiceInspection) SetInspectionContents(v []*CreateTaskRequestParametersServiceInspectionInspectionContents) *CreateTaskRequestParametersServiceInspection {
	s.InspectionContents = v
	return s
}

func (s *CreateTaskRequestParametersServiceInspection) SetInspectionIntroduction(v string) *CreateTaskRequestParametersServiceInspection {
	s.InspectionIntroduction = &v
	return s
}

func (s *CreateTaskRequestParametersServiceInspection) SetSceneIntroduction(v string) *CreateTaskRequestParametersServiceInspection {
	s.SceneIntroduction = &v
	return s
}

func (s *CreateTaskRequestParametersServiceInspection) SetSpeakerMap(v map[string]interface{}) *CreateTaskRequestParametersServiceInspection {
	s.SpeakerMap = v
	return s
}

func (s *CreateTaskRequestParametersServiceInspection) Validate() error {
	return dara.Validate(s)
}

type CreateTaskRequestParametersServiceInspectionInspectionContents struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Title   *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateTaskRequestParametersServiceInspectionInspectionContents) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequestParametersServiceInspectionInspectionContents) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestParametersServiceInspectionInspectionContents) GetContent() *string {
	return s.Content
}

func (s *CreateTaskRequestParametersServiceInspectionInspectionContents) GetTitle() *string {
	return s.Title
}

func (s *CreateTaskRequestParametersServiceInspectionInspectionContents) SetContent(v string) *CreateTaskRequestParametersServiceInspectionInspectionContents {
	s.Content = &v
	return s
}

func (s *CreateTaskRequestParametersServiceInspectionInspectionContents) SetTitle(v string) *CreateTaskRequestParametersServiceInspectionInspectionContents {
	s.Title = &v
	return s
}

func (s *CreateTaskRequestParametersServiceInspectionInspectionContents) Validate() error {
	return dara.Validate(s)
}

type CreateTaskRequestParametersSummarization struct {
	// example:
	//
	// Paragraph
	Types []*string `json:"Types,omitempty" xml:"Types,omitempty" type:"Repeated"`
}

func (s CreateTaskRequestParametersSummarization) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequestParametersSummarization) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestParametersSummarization) GetTypes() []*string {
	return s.Types
}

func (s *CreateTaskRequestParametersSummarization) SetTypes(v []*string) *CreateTaskRequestParametersSummarization {
	s.Types = v
	return s
}

func (s *CreateTaskRequestParametersSummarization) Validate() error {
	return dara.Validate(s)
}

type CreateTaskRequestParametersTranscoding struct {
	// example:
	//
	// false
	SpectrumEnabled *bool `json:"SpectrumEnabled,omitempty" xml:"SpectrumEnabled,omitempty"`
	// example:
	//
	// mp3
	TargetAudioFormat *string `json:"TargetAudioFormat,omitempty" xml:"TargetAudioFormat,omitempty"`
	// example:
	//
	// mp4
	TargetVideoFormat *string `json:"TargetVideoFormat,omitempty" xml:"TargetVideoFormat,omitempty"`
	// example:
	//
	// false
	VideoThumbnailEnabled *bool `json:"VideoThumbnailEnabled,omitempty" xml:"VideoThumbnailEnabled,omitempty"`
}

func (s CreateTaskRequestParametersTranscoding) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequestParametersTranscoding) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestParametersTranscoding) GetSpectrumEnabled() *bool {
	return s.SpectrumEnabled
}

func (s *CreateTaskRequestParametersTranscoding) GetTargetAudioFormat() *string {
	return s.TargetAudioFormat
}

func (s *CreateTaskRequestParametersTranscoding) GetTargetVideoFormat() *string {
	return s.TargetVideoFormat
}

func (s *CreateTaskRequestParametersTranscoding) GetVideoThumbnailEnabled() *bool {
	return s.VideoThumbnailEnabled
}

func (s *CreateTaskRequestParametersTranscoding) SetSpectrumEnabled(v bool) *CreateTaskRequestParametersTranscoding {
	s.SpectrumEnabled = &v
	return s
}

func (s *CreateTaskRequestParametersTranscoding) SetTargetAudioFormat(v string) *CreateTaskRequestParametersTranscoding {
	s.TargetAudioFormat = &v
	return s
}

func (s *CreateTaskRequestParametersTranscoding) SetTargetVideoFormat(v string) *CreateTaskRequestParametersTranscoding {
	s.TargetVideoFormat = &v
	return s
}

func (s *CreateTaskRequestParametersTranscoding) SetVideoThumbnailEnabled(v bool) *CreateTaskRequestParametersTranscoding {
	s.VideoThumbnailEnabled = &v
	return s
}

func (s *CreateTaskRequestParametersTranscoding) Validate() error {
	return dara.Validate(s)
}

type CreateTaskRequestParametersTranscription struct {
	AdditionalStreamOutputLevel *int32 `json:"AdditionalStreamOutputLevel,omitempty" xml:"AdditionalStreamOutputLevel,omitempty"`
	// example:
	//
	// false
	AudioEventDetectionEnabled *bool                                                `json:"AudioEventDetectionEnabled,omitempty" xml:"AudioEventDetectionEnabled,omitempty"`
	Diarization                *CreateTaskRequestParametersTranscriptionDiarization `json:"Diarization,omitempty" xml:"Diarization,omitempty" type:"Struct"`
	// example:
	//
	// false
	DiarizationEnabled         *bool   `json:"DiarizationEnabled,omitempty" xml:"DiarizationEnabled,omitempty"`
	Model                      *string `json:"Model,omitempty" xml:"Model,omitempty"`
	OutputLevel                *int32  `json:"OutputLevel,omitempty" xml:"OutputLevel,omitempty"`
	PhraseId                   *string `json:"PhraseId,omitempty" xml:"PhraseId,omitempty"`
	ProfanityFilterEnabled     *bool   `json:"ProfanityFilterEnabled,omitempty" xml:"ProfanityFilterEnabled,omitempty"`
	RealtimeDiarizationEnabled *bool   `json:"RealtimeDiarizationEnabled,omitempty" xml:"RealtimeDiarizationEnabled,omitempty"`
}

func (s CreateTaskRequestParametersTranscription) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequestParametersTranscription) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestParametersTranscription) GetAdditionalStreamOutputLevel() *int32 {
	return s.AdditionalStreamOutputLevel
}

func (s *CreateTaskRequestParametersTranscription) GetAudioEventDetectionEnabled() *bool {
	return s.AudioEventDetectionEnabled
}

func (s *CreateTaskRequestParametersTranscription) GetDiarization() *CreateTaskRequestParametersTranscriptionDiarization {
	return s.Diarization
}

func (s *CreateTaskRequestParametersTranscription) GetDiarizationEnabled() *bool {
	return s.DiarizationEnabled
}

func (s *CreateTaskRequestParametersTranscription) GetModel() *string {
	return s.Model
}

func (s *CreateTaskRequestParametersTranscription) GetOutputLevel() *int32 {
	return s.OutputLevel
}

func (s *CreateTaskRequestParametersTranscription) GetPhraseId() *string {
	return s.PhraseId
}

func (s *CreateTaskRequestParametersTranscription) GetProfanityFilterEnabled() *bool {
	return s.ProfanityFilterEnabled
}

func (s *CreateTaskRequestParametersTranscription) GetRealtimeDiarizationEnabled() *bool {
	return s.RealtimeDiarizationEnabled
}

func (s *CreateTaskRequestParametersTranscription) SetAdditionalStreamOutputLevel(v int32) *CreateTaskRequestParametersTranscription {
	s.AdditionalStreamOutputLevel = &v
	return s
}

func (s *CreateTaskRequestParametersTranscription) SetAudioEventDetectionEnabled(v bool) *CreateTaskRequestParametersTranscription {
	s.AudioEventDetectionEnabled = &v
	return s
}

func (s *CreateTaskRequestParametersTranscription) SetDiarization(v *CreateTaskRequestParametersTranscriptionDiarization) *CreateTaskRequestParametersTranscription {
	s.Diarization = v
	return s
}

func (s *CreateTaskRequestParametersTranscription) SetDiarizationEnabled(v bool) *CreateTaskRequestParametersTranscription {
	s.DiarizationEnabled = &v
	return s
}

func (s *CreateTaskRequestParametersTranscription) SetModel(v string) *CreateTaskRequestParametersTranscription {
	s.Model = &v
	return s
}

func (s *CreateTaskRequestParametersTranscription) SetOutputLevel(v int32) *CreateTaskRequestParametersTranscription {
	s.OutputLevel = &v
	return s
}

func (s *CreateTaskRequestParametersTranscription) SetPhraseId(v string) *CreateTaskRequestParametersTranscription {
	s.PhraseId = &v
	return s
}

func (s *CreateTaskRequestParametersTranscription) SetProfanityFilterEnabled(v bool) *CreateTaskRequestParametersTranscription {
	s.ProfanityFilterEnabled = &v
	return s
}

func (s *CreateTaskRequestParametersTranscription) SetRealtimeDiarizationEnabled(v bool) *CreateTaskRequestParametersTranscription {
	s.RealtimeDiarizationEnabled = &v
	return s
}

func (s *CreateTaskRequestParametersTranscription) Validate() error {
	return dara.Validate(s)
}

type CreateTaskRequestParametersTranscriptionDiarization struct {
	// example:
	//
	// 2
	SpeakerCount *int32 `json:"SpeakerCount,omitempty" xml:"SpeakerCount,omitempty"`
}

func (s CreateTaskRequestParametersTranscriptionDiarization) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequestParametersTranscriptionDiarization) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestParametersTranscriptionDiarization) GetSpeakerCount() *int32 {
	return s.SpeakerCount
}

func (s *CreateTaskRequestParametersTranscriptionDiarization) SetSpeakerCount(v int32) *CreateTaskRequestParametersTranscriptionDiarization {
	s.SpeakerCount = &v
	return s
}

func (s *CreateTaskRequestParametersTranscriptionDiarization) Validate() error {
	return dara.Validate(s)
}

type CreateTaskRequestParametersTranslation struct {
	AdditionalStreamOutputLevel *int32    `json:"AdditionalStreamOutputLevel,omitempty" xml:"AdditionalStreamOutputLevel,omitempty"`
	OutputLevel                 *int32    `json:"OutputLevel,omitempty" xml:"OutputLevel,omitempty"`
	TargetLanguages             []*string `json:"TargetLanguages,omitempty" xml:"TargetLanguages,omitempty" type:"Repeated"`
}

func (s CreateTaskRequestParametersTranslation) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequestParametersTranslation) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestParametersTranslation) GetAdditionalStreamOutputLevel() *int32 {
	return s.AdditionalStreamOutputLevel
}

func (s *CreateTaskRequestParametersTranslation) GetOutputLevel() *int32 {
	return s.OutputLevel
}

func (s *CreateTaskRequestParametersTranslation) GetTargetLanguages() []*string {
	return s.TargetLanguages
}

func (s *CreateTaskRequestParametersTranslation) SetAdditionalStreamOutputLevel(v int32) *CreateTaskRequestParametersTranslation {
	s.AdditionalStreamOutputLevel = &v
	return s
}

func (s *CreateTaskRequestParametersTranslation) SetOutputLevel(v int32) *CreateTaskRequestParametersTranslation {
	s.OutputLevel = &v
	return s
}

func (s *CreateTaskRequestParametersTranslation) SetTargetLanguages(v []*string) *CreateTaskRequestParametersTranslation {
	s.TargetLanguages = v
	return s
}

func (s *CreateTaskRequestParametersTranslation) Validate() error {
	return dara.Validate(s)
}
