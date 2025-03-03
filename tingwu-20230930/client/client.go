// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

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
	return tea.Prettify(s)
}

func (s CreateTaskRequest) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s CreateTaskRequestInput) GoString() string {
	return s.String()
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

type CreateTaskRequestParameters struct {
	// example:
	//
	// true
	AutoChaptersEnabled      *bool                                         `json:"AutoChaptersEnabled,omitempty" xml:"AutoChaptersEnabled,omitempty"`
	ContentExtraction        *CreateTaskRequestParametersContentExtraction `json:"ContentExtraction,omitempty" xml:"ContentExtraction,omitempty" type:"Struct"`
	ContentExtractionEnabled *bool                                         `json:"ContentExtractionEnabled,omitempty" xml:"ContentExtractionEnabled,omitempty"`
	CustomPrompt             *CreateTaskRequestParametersCustomPrompt      `json:"CustomPrompt,omitempty" xml:"CustomPrompt,omitempty" type:"Struct"`
	CustomPromptEnabled      *bool                                         `json:"CustomPromptEnabled,omitempty" xml:"CustomPromptEnabled,omitempty"`
	ExtraParams              *CreateTaskRequestParametersExtraParams       `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty" type:"Struct"`
	MeetingAssistance        *CreateTaskRequestParametersMeetingAssistance `json:"MeetingAssistance,omitempty" xml:"MeetingAssistance,omitempty" type:"Struct"`
	// example:
	//
	// false
	MeetingAssistanceEnabled *bool                                         `json:"MeetingAssistanceEnabled,omitempty" xml:"MeetingAssistanceEnabled,omitempty"`
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
	return tea.Prettify(s)
}

func (s CreateTaskRequestParameters) GoString() string {
	return s.String()
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

func (s *CreateTaskRequestParameters) SetMeetingAssistance(v *CreateTaskRequestParametersMeetingAssistance) *CreateTaskRequestParameters {
	s.MeetingAssistance = v
	return s
}

func (s *CreateTaskRequestParameters) SetMeetingAssistanceEnabled(v bool) *CreateTaskRequestParameters {
	s.MeetingAssistanceEnabled = &v
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

type CreateTaskRequestParametersContentExtraction struct {
	ExtractionContents []*CreateTaskRequestParametersContentExtractionExtractionContents `json:"ExtractionContents,omitempty" xml:"ExtractionContents,omitempty" type:"Repeated"`
	SceneIntroduction  *string                                                           `json:"SceneIntroduction,omitempty" xml:"SceneIntroduction,omitempty"`
	SpeakerMap         map[string]interface{}                                            `json:"SpeakerMap,omitempty" xml:"SpeakerMap,omitempty"`
}

func (s CreateTaskRequestParametersContentExtraction) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskRequestParametersContentExtraction) GoString() string {
	return s.String()
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

type CreateTaskRequestParametersContentExtractionExtractionContents struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Title   *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateTaskRequestParametersContentExtractionExtractionContents) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskRequestParametersContentExtractionExtractionContents) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestParametersContentExtractionExtractionContents) SetContent(v string) *CreateTaskRequestParametersContentExtractionExtractionContents {
	s.Content = &v
	return s
}

func (s *CreateTaskRequestParametersContentExtractionExtractionContents) SetTitle(v string) *CreateTaskRequestParametersContentExtractionExtractionContents {
	s.Title = &v
	return s
}

type CreateTaskRequestParametersCustomPrompt struct {
	Contents []*CreateTaskRequestParametersCustomPromptContents `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
}

func (s CreateTaskRequestParametersCustomPrompt) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskRequestParametersCustomPrompt) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestParametersCustomPrompt) SetContents(v []*CreateTaskRequestParametersCustomPromptContents) *CreateTaskRequestParametersCustomPrompt {
	s.Contents = v
	return s
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
	return tea.Prettify(s)
}

func (s CreateTaskRequestParametersCustomPromptContents) GoString() string {
	return s.String()
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

type CreateTaskRequestParametersExtraParams struct {
	DomainEducationEnabled   *bool  `json:"DomainEducationEnabled,omitempty" xml:"DomainEducationEnabled,omitempty"`
	MaxKeywords              *int32 `json:"MaxKeywords,omitempty" xml:"MaxKeywords,omitempty"`
	NfixEnabled              *bool  `json:"NfixEnabled,omitempty" xml:"NfixEnabled,omitempty"`
	OcrAuxiliaryEnabled      *bool  `json:"OcrAuxiliaryEnabled,omitempty" xml:"OcrAuxiliaryEnabled,omitempty"`
	TranslateLlmSceneEnabled *bool  `json:"TranslateLlmSceneEnabled,omitempty" xml:"TranslateLlmSceneEnabled,omitempty"`
}

func (s CreateTaskRequestParametersExtraParams) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskRequestParametersExtraParams) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestParametersExtraParams) SetDomainEducationEnabled(v bool) *CreateTaskRequestParametersExtraParams {
	s.DomainEducationEnabled = &v
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

type CreateTaskRequestParametersMeetingAssistance struct {
	Types []*string `json:"Types,omitempty" xml:"Types,omitempty" type:"Repeated"`
}

func (s CreateTaskRequestParametersMeetingAssistance) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskRequestParametersMeetingAssistance) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestParametersMeetingAssistance) SetTypes(v []*string) *CreateTaskRequestParametersMeetingAssistance {
	s.Types = v
	return s
}

type CreateTaskRequestParametersServiceInspection struct {
	InspectionContents     []*CreateTaskRequestParametersServiceInspectionInspectionContents `json:"InspectionContents,omitempty" xml:"InspectionContents,omitempty" type:"Repeated"`
	InspectionIntroduction *string                                                           `json:"InspectionIntroduction,omitempty" xml:"InspectionIntroduction,omitempty"`
	SceneIntroduction      *string                                                           `json:"SceneIntroduction,omitempty" xml:"SceneIntroduction,omitempty"`
	SpeakerMap             map[string]interface{}                                            `json:"SpeakerMap,omitempty" xml:"SpeakerMap,omitempty"`
}

func (s CreateTaskRequestParametersServiceInspection) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskRequestParametersServiceInspection) GoString() string {
	return s.String()
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

type CreateTaskRequestParametersServiceInspectionInspectionContents struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Title   *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateTaskRequestParametersServiceInspectionInspectionContents) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskRequestParametersServiceInspectionInspectionContents) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestParametersServiceInspectionInspectionContents) SetContent(v string) *CreateTaskRequestParametersServiceInspectionInspectionContents {
	s.Content = &v
	return s
}

func (s *CreateTaskRequestParametersServiceInspectionInspectionContents) SetTitle(v string) *CreateTaskRequestParametersServiceInspectionInspectionContents {
	s.Title = &v
	return s
}

type CreateTaskRequestParametersSummarization struct {
	// example:
	//
	// Paragraph
	Types []*string `json:"Types,omitempty" xml:"Types,omitempty" type:"Repeated"`
}

func (s CreateTaskRequestParametersSummarization) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskRequestParametersSummarization) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestParametersSummarization) SetTypes(v []*string) *CreateTaskRequestParametersSummarization {
	s.Types = v
	return s
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
	return tea.Prettify(s)
}

func (s CreateTaskRequestParametersTranscoding) GoString() string {
	return s.String()
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
	RealtimeDiarizationEnabled *bool   `json:"RealtimeDiarizationEnabled,omitempty" xml:"RealtimeDiarizationEnabled,omitempty"`
}

func (s CreateTaskRequestParametersTranscription) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskRequestParametersTranscription) GoString() string {
	return s.String()
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

func (s *CreateTaskRequestParametersTranscription) SetRealtimeDiarizationEnabled(v bool) *CreateTaskRequestParametersTranscription {
	s.RealtimeDiarizationEnabled = &v
	return s
}

type CreateTaskRequestParametersTranscriptionDiarization struct {
	// example:
	//
	// 2
	SpeakerCount *int32 `json:"SpeakerCount,omitempty" xml:"SpeakerCount,omitempty"`
}

func (s CreateTaskRequestParametersTranscriptionDiarization) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskRequestParametersTranscriptionDiarization) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestParametersTranscriptionDiarization) SetSpeakerCount(v int32) *CreateTaskRequestParametersTranscriptionDiarization {
	s.SpeakerCount = &v
	return s
}

type CreateTaskRequestParametersTranslation struct {
	AdditionalStreamOutputLevel *int32    `json:"AdditionalStreamOutputLevel,omitempty" xml:"AdditionalStreamOutputLevel,omitempty"`
	OutputLevel                 *int32    `json:"OutputLevel,omitempty" xml:"OutputLevel,omitempty"`
	TargetLanguages             []*string `json:"TargetLanguages,omitempty" xml:"TargetLanguages,omitempty" type:"Repeated"`
}

func (s CreateTaskRequestParametersTranslation) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskRequestParametersTranslation) GoString() string {
	return s.String()
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

type CreateTaskResponseBody struct {
	// example:
	//
	// 0
	Code *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 35124E1C-AE99-5D6C-A52E-BD689D8D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTaskResponseBody) SetCode(v string) *CreateTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTaskResponseBody) SetData(v *CreateTaskResponseBodyData) *CreateTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateTaskResponseBody) SetMessage(v string) *CreateTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTaskResponseBody) SetRequestId(v string) *CreateTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateTaskResponseBodyData struct {
	MeetingJoinUrl *string `json:"MeetingJoinUrl,omitempty" xml:"MeetingJoinUrl,omitempty"`
	// example:
	//
	// c5394c6ee0fb474899d42215a3925c7e
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// task_tingwu_123
	TaskKey    *string `json:"TaskKey,omitempty" xml:"TaskKey,omitempty"`
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s CreateTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateTaskResponseBodyData) SetMeetingJoinUrl(v string) *CreateTaskResponseBodyData {
	s.MeetingJoinUrl = &v
	return s
}

func (s *CreateTaskResponseBodyData) SetTaskId(v string) *CreateTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *CreateTaskResponseBodyData) SetTaskKey(v string) *CreateTaskResponseBodyData {
	s.TaskKey = &v
	return s
}

func (s *CreateTaskResponseBodyData) SetTaskStatus(v string) *CreateTaskResponseBodyData {
	s.TaskStatus = &v
	return s
}

type CreateTaskResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateTaskResponse) SetHeaders(v map[string]*string) *CreateTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateTaskResponse) SetStatusCode(v int32) *CreateTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTaskResponse) SetBody(v *CreateTaskResponseBody) *CreateTaskResponse {
	s.Body = v
	return s
}

type CreateTranscriptionPhrasesRequest struct {
	// example:
	//
	// custom fruit phrases list
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// fruit_phrase
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"苹果":3,"西瓜":3}
	WordWeights *string `json:"WordWeights,omitempty" xml:"WordWeights,omitempty"`
}

func (s CreateTranscriptionPhrasesRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTranscriptionPhrasesRequest) GoString() string {
	return s.String()
}

func (s *CreateTranscriptionPhrasesRequest) SetDescription(v string) *CreateTranscriptionPhrasesRequest {
	s.Description = &v
	return s
}

func (s *CreateTranscriptionPhrasesRequest) SetName(v string) *CreateTranscriptionPhrasesRequest {
	s.Name = &v
	return s
}

func (s *CreateTranscriptionPhrasesRequest) SetWordWeights(v string) *CreateTranscriptionPhrasesRequest {
	s.WordWeights = &v
	return s
}

type CreateTranscriptionPhrasesResponseBody struct {
	// example:
	//
	// 0
	Code *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateTranscriptionPhrasesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 35124E1C-AE99-5D6C-A52E-BD689D8D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTranscriptionPhrasesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTranscriptionPhrasesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTranscriptionPhrasesResponseBody) SetCode(v string) *CreateTranscriptionPhrasesResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTranscriptionPhrasesResponseBody) SetData(v *CreateTranscriptionPhrasesResponseBodyData) *CreateTranscriptionPhrasesResponseBody {
	s.Data = v
	return s
}

func (s *CreateTranscriptionPhrasesResponseBody) SetMessage(v string) *CreateTranscriptionPhrasesResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTranscriptionPhrasesResponseBody) SetRequestId(v string) *CreateTranscriptionPhrasesResponseBody {
	s.RequestId = &v
	return s
}

type CreateTranscriptionPhrasesResponseBodyData struct {
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// success
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// a93b91141c0f422fa114af203f8b****
	PhraseId *string `json:"PhraseId,omitempty" xml:"PhraseId,omitempty"`
	// example:
	//
	// SUCCEEDED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateTranscriptionPhrasesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateTranscriptionPhrasesResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateTranscriptionPhrasesResponseBodyData) SetErrorCode(v string) *CreateTranscriptionPhrasesResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *CreateTranscriptionPhrasesResponseBodyData) SetErrorMessage(v string) *CreateTranscriptionPhrasesResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *CreateTranscriptionPhrasesResponseBodyData) SetPhraseId(v string) *CreateTranscriptionPhrasesResponseBodyData {
	s.PhraseId = &v
	return s
}

func (s *CreateTranscriptionPhrasesResponseBodyData) SetStatus(v string) *CreateTranscriptionPhrasesResponseBodyData {
	s.Status = &v
	return s
}

type CreateTranscriptionPhrasesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTranscriptionPhrasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTranscriptionPhrasesResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTranscriptionPhrasesResponse) GoString() string {
	return s.String()
}

func (s *CreateTranscriptionPhrasesResponse) SetHeaders(v map[string]*string) *CreateTranscriptionPhrasesResponse {
	s.Headers = v
	return s
}

func (s *CreateTranscriptionPhrasesResponse) SetStatusCode(v int32) *CreateTranscriptionPhrasesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTranscriptionPhrasesResponse) SetBody(v *CreateTranscriptionPhrasesResponseBody) *CreateTranscriptionPhrasesResponse {
	s.Body = v
	return s
}

type DeleteTranscriptionPhrasesResponseBody struct {
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// success
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// SUCCEEDED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteTranscriptionPhrasesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTranscriptionPhrasesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTranscriptionPhrasesResponseBody) SetErrorCode(v string) *DeleteTranscriptionPhrasesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteTranscriptionPhrasesResponseBody) SetErrorMessage(v string) *DeleteTranscriptionPhrasesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteTranscriptionPhrasesResponseBody) SetStatus(v string) *DeleteTranscriptionPhrasesResponseBody {
	s.Status = &v
	return s
}

type DeleteTranscriptionPhrasesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTranscriptionPhrasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTranscriptionPhrasesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTranscriptionPhrasesResponse) GoString() string {
	return s.String()
}

func (s *DeleteTranscriptionPhrasesResponse) SetHeaders(v map[string]*string) *DeleteTranscriptionPhrasesResponse {
	s.Headers = v
	return s
}

func (s *DeleteTranscriptionPhrasesResponse) SetStatusCode(v int32) *DeleteTranscriptionPhrasesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTranscriptionPhrasesResponse) SetBody(v *DeleteTranscriptionPhrasesResponseBody) *DeleteTranscriptionPhrasesResponse {
	s.Body = v
	return s
}

type GetTaskInfoResponseBody struct {
	// example:
	//
	// 0
	Code *string                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetTaskInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 35124E1C-AE99-5D6C-A52E-BD689D8D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTaskInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskInfoResponseBody) SetCode(v string) *GetTaskInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetTaskInfoResponseBody) SetData(v *GetTaskInfoResponseBodyData) *GetTaskInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetTaskInfoResponseBody) SetMessage(v string) *GetTaskInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetTaskInfoResponseBody) SetRequestId(v string) *GetTaskInfoResponseBody {
	s.RequestId = &v
	return s
}

type GetTaskInfoResponseBodyData struct {
	ErrorCode           *string                            `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMessage        *string                            `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	OutputMp3Path       *string                            `json:"OutputMp3Path,omitempty" xml:"OutputMp3Path,omitempty"`
	OutputMp4Path       *string                            `json:"OutputMp4Path,omitempty" xml:"OutputMp4Path,omitempty"`
	OutputSpectrumPath  *string                            `json:"OutputSpectrumPath,omitempty" xml:"OutputSpectrumPath,omitempty"`
	OutputThumbnailPath *string                            `json:"OutputThumbnailPath,omitempty" xml:"OutputThumbnailPath,omitempty"`
	Result              *GetTaskInfoResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// c5394c6ee0fb474899d42215a3925c7e
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// task_tingwu_123
	TaskKey *string `json:"TaskKey,omitempty" xml:"TaskKey,omitempty"`
	// example:
	//
	// COMPLETE
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s GetTaskInfoResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTaskInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTaskInfoResponseBodyData) SetErrorCode(v string) *GetTaskInfoResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *GetTaskInfoResponseBodyData) SetErrorMessage(v string) *GetTaskInfoResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetTaskInfoResponseBodyData) SetOutputMp3Path(v string) *GetTaskInfoResponseBodyData {
	s.OutputMp3Path = &v
	return s
}

func (s *GetTaskInfoResponseBodyData) SetOutputMp4Path(v string) *GetTaskInfoResponseBodyData {
	s.OutputMp4Path = &v
	return s
}

func (s *GetTaskInfoResponseBodyData) SetOutputSpectrumPath(v string) *GetTaskInfoResponseBodyData {
	s.OutputSpectrumPath = &v
	return s
}

func (s *GetTaskInfoResponseBodyData) SetOutputThumbnailPath(v string) *GetTaskInfoResponseBodyData {
	s.OutputThumbnailPath = &v
	return s
}

func (s *GetTaskInfoResponseBodyData) SetResult(v *GetTaskInfoResponseBodyDataResult) *GetTaskInfoResponseBodyData {
	s.Result = v
	return s
}

func (s *GetTaskInfoResponseBodyData) SetTaskId(v string) *GetTaskInfoResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetTaskInfoResponseBodyData) SetTaskKey(v string) *GetTaskInfoResponseBodyData {
	s.TaskKey = &v
	return s
}

func (s *GetTaskInfoResponseBodyData) SetTaskStatus(v string) *GetTaskInfoResponseBodyData {
	s.TaskStatus = &v
	return s
}

type GetTaskInfoResponseBodyDataResult struct {
	AutoChapters      *string `json:"AutoChapters,omitempty" xml:"AutoChapters,omitempty"`
	CustomPrompt      *string `json:"CustomPrompt,omitempty" xml:"CustomPrompt,omitempty"`
	MeetingAssistance *string `json:"MeetingAssistance,omitempty" xml:"MeetingAssistance,omitempty"`
	PptExtraction     *string `json:"PptExtraction,omitempty" xml:"PptExtraction,omitempty"`
	ServiceInspection *string `json:"ServiceInspection,omitempty" xml:"ServiceInspection,omitempty"`
	Summarization     *string `json:"Summarization,omitempty" xml:"Summarization,omitempty"`
	TextPolish        *string `json:"TextPolish,omitempty" xml:"TextPolish,omitempty"`
	Transcription     *string `json:"Transcription,omitempty" xml:"Transcription,omitempty"`
	Translation       *string `json:"Translation,omitempty" xml:"Translation,omitempty"`
}

func (s GetTaskInfoResponseBodyDataResult) String() string {
	return tea.Prettify(s)
}

func (s GetTaskInfoResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *GetTaskInfoResponseBodyDataResult) SetAutoChapters(v string) *GetTaskInfoResponseBodyDataResult {
	s.AutoChapters = &v
	return s
}

func (s *GetTaskInfoResponseBodyDataResult) SetCustomPrompt(v string) *GetTaskInfoResponseBodyDataResult {
	s.CustomPrompt = &v
	return s
}

func (s *GetTaskInfoResponseBodyDataResult) SetMeetingAssistance(v string) *GetTaskInfoResponseBodyDataResult {
	s.MeetingAssistance = &v
	return s
}

func (s *GetTaskInfoResponseBodyDataResult) SetPptExtraction(v string) *GetTaskInfoResponseBodyDataResult {
	s.PptExtraction = &v
	return s
}

func (s *GetTaskInfoResponseBodyDataResult) SetServiceInspection(v string) *GetTaskInfoResponseBodyDataResult {
	s.ServiceInspection = &v
	return s
}

func (s *GetTaskInfoResponseBodyDataResult) SetSummarization(v string) *GetTaskInfoResponseBodyDataResult {
	s.Summarization = &v
	return s
}

func (s *GetTaskInfoResponseBodyDataResult) SetTextPolish(v string) *GetTaskInfoResponseBodyDataResult {
	s.TextPolish = &v
	return s
}

func (s *GetTaskInfoResponseBodyDataResult) SetTranscription(v string) *GetTaskInfoResponseBodyDataResult {
	s.Transcription = &v
	return s
}

func (s *GetTaskInfoResponseBodyDataResult) SetTranslation(v string) *GetTaskInfoResponseBodyDataResult {
	s.Translation = &v
	return s
}

type GetTaskInfoResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskInfoResponse) GoString() string {
	return s.String()
}

func (s *GetTaskInfoResponse) SetHeaders(v map[string]*string) *GetTaskInfoResponse {
	s.Headers = v
	return s
}

func (s *GetTaskInfoResponse) SetStatusCode(v int32) *GetTaskInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskInfoResponse) SetBody(v *GetTaskInfoResponseBody) *GetTaskInfoResponse {
	s.Body = v
	return s
}

type GetTranscriptionPhrasesResponseBody struct {
	// example:
	//
	// 0
	Code *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetTranscriptionPhrasesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 35124E1C-AE99-5D6C-A52E-BD689D8D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTranscriptionPhrasesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTranscriptionPhrasesResponseBody) GoString() string {
	return s.String()
}

func (s *GetTranscriptionPhrasesResponseBody) SetCode(v string) *GetTranscriptionPhrasesResponseBody {
	s.Code = &v
	return s
}

func (s *GetTranscriptionPhrasesResponseBody) SetData(v *GetTranscriptionPhrasesResponseBodyData) *GetTranscriptionPhrasesResponseBody {
	s.Data = v
	return s
}

func (s *GetTranscriptionPhrasesResponseBody) SetMessage(v string) *GetTranscriptionPhrasesResponseBody {
	s.Message = &v
	return s
}

func (s *GetTranscriptionPhrasesResponseBody) SetRequestId(v string) *GetTranscriptionPhrasesResponseBody {
	s.RequestId = &v
	return s
}

type GetTranscriptionPhrasesResponseBodyData struct {
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// success
	ErrorMessage *string                                           `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Phrases      []*GetTranscriptionPhrasesResponseBodyDataPhrases `json:"Phrases,omitempty" xml:"Phrases,omitempty" type:"Repeated"`
	// example:
	//
	// SUCCEEDED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetTranscriptionPhrasesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetTranscriptionPhrasesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTranscriptionPhrasesResponseBodyData) SetErrorCode(v string) *GetTranscriptionPhrasesResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *GetTranscriptionPhrasesResponseBodyData) SetErrorMessage(v string) *GetTranscriptionPhrasesResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetTranscriptionPhrasesResponseBodyData) SetPhrases(v []*GetTranscriptionPhrasesResponseBodyDataPhrases) *GetTranscriptionPhrasesResponseBodyData {
	s.Phrases = v
	return s
}

func (s *GetTranscriptionPhrasesResponseBodyData) SetStatus(v string) *GetTranscriptionPhrasesResponseBodyData {
	s.Status = &v
	return s
}

type GetTranscriptionPhrasesResponseBodyDataPhrases struct {
	// example:
	//
	// custom fruit phrases list
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// fruit_phrase
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// a93b91141c0f422fa114af203f8b****
	PhraseId *string `json:"PhraseId,omitempty" xml:"PhraseId,omitempty"`
	// example:
	//
	// {"苹果":3,"西瓜":3}
	WordWeights *string `json:"WordWeights,omitempty" xml:"WordWeights,omitempty"`
}

func (s GetTranscriptionPhrasesResponseBodyDataPhrases) String() string {
	return tea.Prettify(s)
}

func (s GetTranscriptionPhrasesResponseBodyDataPhrases) GoString() string {
	return s.String()
}

func (s *GetTranscriptionPhrasesResponseBodyDataPhrases) SetDescription(v string) *GetTranscriptionPhrasesResponseBodyDataPhrases {
	s.Description = &v
	return s
}

func (s *GetTranscriptionPhrasesResponseBodyDataPhrases) SetName(v string) *GetTranscriptionPhrasesResponseBodyDataPhrases {
	s.Name = &v
	return s
}

func (s *GetTranscriptionPhrasesResponseBodyDataPhrases) SetPhraseId(v string) *GetTranscriptionPhrasesResponseBodyDataPhrases {
	s.PhraseId = &v
	return s
}

func (s *GetTranscriptionPhrasesResponseBodyDataPhrases) SetWordWeights(v string) *GetTranscriptionPhrasesResponseBodyDataPhrases {
	s.WordWeights = &v
	return s
}

type GetTranscriptionPhrasesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTranscriptionPhrasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTranscriptionPhrasesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTranscriptionPhrasesResponse) GoString() string {
	return s.String()
}

func (s *GetTranscriptionPhrasesResponse) SetHeaders(v map[string]*string) *GetTranscriptionPhrasesResponse {
	s.Headers = v
	return s
}

func (s *GetTranscriptionPhrasesResponse) SetStatusCode(v int32) *GetTranscriptionPhrasesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTranscriptionPhrasesResponse) SetBody(v *GetTranscriptionPhrasesResponseBody) *GetTranscriptionPhrasesResponse {
	s.Body = v
	return s
}

type ListTranscriptionPhrasesResponseBody struct {
	// example:
	//
	// 0
	Code *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListTranscriptionPhrasesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 35124E1C-AE99-5D6C-A52E-BD689D8D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTranscriptionPhrasesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTranscriptionPhrasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTranscriptionPhrasesResponseBody) SetCode(v string) *ListTranscriptionPhrasesResponseBody {
	s.Code = &v
	return s
}

func (s *ListTranscriptionPhrasesResponseBody) SetData(v *ListTranscriptionPhrasesResponseBodyData) *ListTranscriptionPhrasesResponseBody {
	s.Data = v
	return s
}

func (s *ListTranscriptionPhrasesResponseBody) SetMessage(v string) *ListTranscriptionPhrasesResponseBody {
	s.Message = &v
	return s
}

func (s *ListTranscriptionPhrasesResponseBody) SetRequestId(v string) *ListTranscriptionPhrasesResponseBody {
	s.RequestId = &v
	return s
}

type ListTranscriptionPhrasesResponseBodyData struct {
	// example:
	//
	// PHS.Exceed
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// The num of the phrase exceeds the upper limit.
	ErrorMessage *string                                            `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Phrases      []*ListTranscriptionPhrasesResponseBodyDataPhrases `json:"Phrases,omitempty" xml:"Phrases,omitempty" type:"Repeated"`
	// example:
	//
	// SUCCEEDED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListTranscriptionPhrasesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListTranscriptionPhrasesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTranscriptionPhrasesResponseBodyData) SetErrorCode(v string) *ListTranscriptionPhrasesResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *ListTranscriptionPhrasesResponseBodyData) SetErrorMessage(v string) *ListTranscriptionPhrasesResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *ListTranscriptionPhrasesResponseBodyData) SetPhrases(v []*ListTranscriptionPhrasesResponseBodyDataPhrases) *ListTranscriptionPhrasesResponseBodyData {
	s.Phrases = v
	return s
}

func (s *ListTranscriptionPhrasesResponseBodyData) SetStatus(v string) *ListTranscriptionPhrasesResponseBodyData {
	s.Status = &v
	return s
}

type ListTranscriptionPhrasesResponseBodyDataPhrases struct {
	// example:
	//
	// custom fruit phrases list
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// fruit_phrase
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// a93b91141c0f422fa114af203f8b****
	PhraseId *string `json:"PhraseId,omitempty" xml:"PhraseId,omitempty"`
}

func (s ListTranscriptionPhrasesResponseBodyDataPhrases) String() string {
	return tea.Prettify(s)
}

func (s ListTranscriptionPhrasesResponseBodyDataPhrases) GoString() string {
	return s.String()
}

func (s *ListTranscriptionPhrasesResponseBodyDataPhrases) SetDescription(v string) *ListTranscriptionPhrasesResponseBodyDataPhrases {
	s.Description = &v
	return s
}

func (s *ListTranscriptionPhrasesResponseBodyDataPhrases) SetName(v string) *ListTranscriptionPhrasesResponseBodyDataPhrases {
	s.Name = &v
	return s
}

func (s *ListTranscriptionPhrasesResponseBodyDataPhrases) SetPhraseId(v string) *ListTranscriptionPhrasesResponseBodyDataPhrases {
	s.PhraseId = &v
	return s
}

type ListTranscriptionPhrasesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTranscriptionPhrasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTranscriptionPhrasesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTranscriptionPhrasesResponse) GoString() string {
	return s.String()
}

func (s *ListTranscriptionPhrasesResponse) SetHeaders(v map[string]*string) *ListTranscriptionPhrasesResponse {
	s.Headers = v
	return s
}

func (s *ListTranscriptionPhrasesResponse) SetStatusCode(v int32) *ListTranscriptionPhrasesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTranscriptionPhrasesResponse) SetBody(v *ListTranscriptionPhrasesResponseBody) *ListTranscriptionPhrasesResponse {
	s.Body = v
	return s
}

type UpdateTranscriptionPhrasesRequest struct {
	// example:
	//
	// custom fruit phrases list
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// fruit_phrase
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"苹果":3,"西瓜":3}
	WordWeights *string `json:"WordWeights,omitempty" xml:"WordWeights,omitempty"`
}

func (s UpdateTranscriptionPhrasesRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTranscriptionPhrasesRequest) GoString() string {
	return s.String()
}

func (s *UpdateTranscriptionPhrasesRequest) SetDescription(v string) *UpdateTranscriptionPhrasesRequest {
	s.Description = &v
	return s
}

func (s *UpdateTranscriptionPhrasesRequest) SetName(v string) *UpdateTranscriptionPhrasesRequest {
	s.Name = &v
	return s
}

func (s *UpdateTranscriptionPhrasesRequest) SetWordWeights(v string) *UpdateTranscriptionPhrasesRequest {
	s.WordWeights = &v
	return s
}

type UpdateTranscriptionPhrasesResponseBody struct {
	// example:
	//
	// 0
	Code *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *UpdateTranscriptionPhrasesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 35124E1C-AE99-5D6C-A52E-BD689D8D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTranscriptionPhrasesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTranscriptionPhrasesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTranscriptionPhrasesResponseBody) SetCode(v string) *UpdateTranscriptionPhrasesResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTranscriptionPhrasesResponseBody) SetData(v *UpdateTranscriptionPhrasesResponseBodyData) *UpdateTranscriptionPhrasesResponseBody {
	s.Data = v
	return s
}

func (s *UpdateTranscriptionPhrasesResponseBody) SetMessage(v string) *UpdateTranscriptionPhrasesResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTranscriptionPhrasesResponseBody) SetRequestId(v string) *UpdateTranscriptionPhrasesResponseBody {
	s.RequestId = &v
	return s
}

type UpdateTranscriptionPhrasesResponseBodyData struct {
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// success
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// SUCCEEDED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateTranscriptionPhrasesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateTranscriptionPhrasesResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateTranscriptionPhrasesResponseBodyData) SetErrorCode(v string) *UpdateTranscriptionPhrasesResponseBodyData {
	s.ErrorCode = &v
	return s
}

func (s *UpdateTranscriptionPhrasesResponseBodyData) SetErrorMessage(v string) *UpdateTranscriptionPhrasesResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateTranscriptionPhrasesResponseBodyData) SetStatus(v string) *UpdateTranscriptionPhrasesResponseBodyData {
	s.Status = &v
	return s
}

type UpdateTranscriptionPhrasesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTranscriptionPhrasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTranscriptionPhrasesResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTranscriptionPhrasesResponse) GoString() string {
	return s.String()
}

func (s *UpdateTranscriptionPhrasesResponse) SetHeaders(v map[string]*string) *UpdateTranscriptionPhrasesResponse {
	s.Headers = v
	return s
}

func (s *UpdateTranscriptionPhrasesResponse) SetStatusCode(v int32) *UpdateTranscriptionPhrasesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTranscriptionPhrasesResponse) SetBody(v *UpdateTranscriptionPhrasesResponseBody) *UpdateTranscriptionPhrasesResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("tingwu"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
// 创建听悟任务
//
// @param request - CreateTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTaskResponse
func (client *Client) CreateTaskWithOptions(request *CreateTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Operation)) {
		query["operation"] = request.Operation
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppKey)) {
		body["AppKey"] = request.AppKey
	}

	if !tea.BoolValue(util.IsUnset(request.Input)) {
		body["Input"] = request.Input
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		body["Parameters"] = request.Parameters
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTask"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/tingwu/v2/tasks"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateTaskResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateTaskResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建听悟任务
//
// @param request - CreateTaskRequest
//
// @return CreateTaskResponse
func (client *Client) CreateTask(request *CreateTaskRequest) (_result *CreateTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTaskResponse{}
	_body, _err := client.CreateTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建热词词表
//
// @param request - CreateTranscriptionPhrasesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateTranscriptionPhrasesResponse
func (client *Client) CreateTranscriptionPhrasesWithOptions(request *CreateTranscriptionPhrasesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateTranscriptionPhrasesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.WordWeights)) {
		body["WordWeights"] = request.WordWeights
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTranscriptionPhrases"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/tingwu/v2/resources/phrases"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateTranscriptionPhrasesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateTranscriptionPhrasesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建热词词表
//
// @param request - CreateTranscriptionPhrasesRequest
//
// @return CreateTranscriptionPhrasesResponse
func (client *Client) CreateTranscriptionPhrases(request *CreateTranscriptionPhrasesRequest) (_result *CreateTranscriptionPhrasesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTranscriptionPhrasesResponse{}
	_body, _err := client.CreateTranscriptionPhrasesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除词表
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteTranscriptionPhrasesResponse
func (client *Client) DeleteTranscriptionPhrasesWithOptions(PhraseId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteTranscriptionPhrasesResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTranscriptionPhrases"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/tingwu/v2/resources/phrases/" + tea.StringValue(openapiutil.GetEncodeParam(PhraseId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &DeleteTranscriptionPhrasesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &DeleteTranscriptionPhrasesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 删除词表
//
// @return DeleteTranscriptionPhrasesResponse
func (client *Client) DeleteTranscriptionPhrases(PhraseId *string) (_result *DeleteTranscriptionPhrasesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTranscriptionPhrasesResponse{}
	_body, _err := client.DeleteTranscriptionPhrasesWithOptions(PhraseId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询听悟任务信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskInfoResponse
func (client *Client) GetTaskInfoWithOptions(TaskId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTaskInfoResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetTaskInfo"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/tingwu/v2/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(TaskId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetTaskInfoResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetTaskInfoResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查询听悟任务信息
//
// @return GetTaskInfoResponse
func (client *Client) GetTaskInfo(TaskId *string) (_result *GetTaskInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTaskInfoResponse{}
	_body, _err := client.GetTaskInfoWithOptions(TaskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询热词词表信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTranscriptionPhrasesResponse
func (client *Client) GetTranscriptionPhrasesWithOptions(PhraseId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTranscriptionPhrasesResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetTranscriptionPhrases"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/tingwu/v2/resources/phrases/" + tea.StringValue(openapiutil.GetEncodeParam(PhraseId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetTranscriptionPhrasesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetTranscriptionPhrasesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查询热词词表信息
//
// @return GetTranscriptionPhrasesResponse
func (client *Client) GetTranscriptionPhrases(PhraseId *string) (_result *GetTranscriptionPhrasesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTranscriptionPhrasesResponse{}
	_body, _err := client.GetTranscriptionPhrasesWithOptions(PhraseId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列举用户所有热词词表信息
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTranscriptionPhrasesResponse
func (client *Client) ListTranscriptionPhrasesWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTranscriptionPhrasesResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListTranscriptionPhrases"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/tingwu/v2/resources/phrases"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListTranscriptionPhrasesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListTranscriptionPhrasesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 列举用户所有热词词表信息
//
// @return ListTranscriptionPhrasesResponse
func (client *Client) ListTranscriptionPhrases() (_result *ListTranscriptionPhrasesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTranscriptionPhrasesResponse{}
	_body, _err := client.ListTranscriptionPhrasesWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新热词词表
//
// @param request - UpdateTranscriptionPhrasesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateTranscriptionPhrasesResponse
func (client *Client) UpdateTranscriptionPhrasesWithOptions(PhraseId *string, request *UpdateTranscriptionPhrasesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateTranscriptionPhrasesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.WordWeights)) {
		body["WordWeights"] = request.WordWeights
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTranscriptionPhrases"),
		Version:     tea.String("2023-09-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/openapi/tingwu/v2/resources/phrases/" + tea.StringValue(openapiutil.GetEncodeParam(PhraseId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateTranscriptionPhrasesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateTranscriptionPhrasesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 更新热词词表
//
// @param request - UpdateTranscriptionPhrasesRequest
//
// @return UpdateTranscriptionPhrasesResponse
func (client *Client) UpdateTranscriptionPhrases(PhraseId *string, request *UpdateTranscriptionPhrasesRequest) (_result *UpdateTranscriptionPhrasesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateTranscriptionPhrasesResponse{}
	_body, _err := client.UpdateTranscriptionPhrasesWithOptions(PhraseId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
