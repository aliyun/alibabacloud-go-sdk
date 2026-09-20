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
	// The AppKey of the project created in the console.
	//
	// example:
	//
	// JV1sRTisRMi****
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// The basic parameters set when creating a task. The required parameters vary depending on the task type.
	//
	// - When type=offline (offline task), you must set the SourceLanguage and FileUrl parameters.
	//
	// - When type=realtime (real-time meeting task), you must additionally set the SourceLanguage, Format, and SampleRate parameters.
	Input *CreateTaskRequestInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The algorithm-related parameters set when creating a task. You can set these as needed.
	Parameters *CreateTaskRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Struct"`
	// The operation. Valid values:
	//
	// - start: creates a task. This is the default value. In most cases, you do not need to explicitly set this parameter.
	//
	// - stop: stops a real-time meeting task. This value is used in real-time meeting scenarios. After a meeting ends, set this parameter to stop and trigger the call.
	//
	// > Note: When ending a real-time recording, you must set this parameter to stop.
	//
	// example:
	//
	// stop
	Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
	// The task type. Valid values:
	//
	// - **offline**: offline task, such as offline transcription.
	//
	// - **realtime**: real-time task, such as creating a real-time recording.
	//
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
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	if s.Parameters != nil {
		if err := s.Parameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTaskRequestInput struct {
	// The multi-channel audio and video processing mode.
	AudioChannelMode *string `json:"AudioChannelMode,omitempty" xml:"AudioChannelMode,omitempty"`
	// The HTTP or HTTPS URL of the original audio or video file. This parameter is required when you create an offline transcription task.
	//
	// example:
	//
	// http://xxx.com/zzz/1.wav
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// The encoding format of the audio stream data when you create a real-time meeting, such as pcm. Valid values:
	//
	// - **pcm**
	//
	// - **opus**
	//
	// - **aac**
	//
	// - **speex**
	//
	// - **mp3**
	//
	// example:
	//
	// pcm
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The preferred languages. This parameter takes effect only when SourceLanguage is set to "multilingual". It restricts the output languages of the model.
	LanguageHints []*string `json:"LanguageHints,omitempty" xml:"LanguageHints,omitempty" type:"Repeated"`
	// Specifies whether to enable multi-channel audio stream recognition. This parameter needs to be set only in real-time recording scenarios. Default value: false.
	//
	// example:
	//
	// false
	MultipleStreamsEnabled *bool `json:"MultipleStreamsEnabled,omitempty" xml:"MultipleStreamsEnabled,omitempty"`
	// After configuring OSS information in the console, you can specify an OSS write path to save results directly to your custom OSS bucket.
	OutputPath *string `json:"OutputPath,omitempty" xml:"OutputPath,omitempty"`
	// Specifies whether to enable the callback feature.
	//
	// To enable the callback feature, configure the callback type and address in the console, and set this parameter to true when creating a task.
	//
	// example:
	//
	// false
	ProgressiveCallbacksEnabled *bool `json:"ProgressiveCallbacksEnabled,omitempty" xml:"ProgressiveCallbacksEnabled,omitempty"`
	// The sample rate of the audio stream data when you create a real-time meeting. Valid values: 8000 and 16000.
	//
	// - **8000**: telephone customer service scenarios.
	//
	// - **16000**: real-time meeting audio capture scenarios.
	//
	// example:
	//
	// 16000
	SampleRate *int32 `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	// The language model used for audio transcription. Valid values:
	//
	// - **cn**: Chinese
	//
	// - **en**: English
	//
	// - **fspk**: Chinese-English free speaking
	//
	// - **ja**: Japanese
	//
	// - **yue**: Cantonese
	//
	// This parameter is required.
	//
	// example:
	//
	// cn
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// The TaskId returned when you create a real-time recording. You can use this ID to end the real-time recording. Set this parameter only when ending a real-time recording. Do not set it at other times.
	//
	// example:
	//
	// 9922c84c087044eda18659c128b56c84
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The custom identifier set by the user to associate with this task.
	//
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
	// Specifies whether to enable the chapter overview feature. When enabled, chapter titles and chapter summaries are generated.
	//
	// example:
	//
	// true
	AutoChaptersEnabled *bool `json:"AutoChaptersEnabled,omitempty" xml:"AutoChaptersEnabled,omitempty"`
	// The conversation content extraction parameter object.
	ContentExtraction *CreateTaskRequestParametersContentExtraction `json:"ContentExtraction,omitempty" xml:"ContentExtraction,omitempty" type:"Struct"`
	// The business user ID.
	ContentExtractionEnabled *bool `json:"ContentExtractionEnabled,omitempty" xml:"ContentExtractionEnabled,omitempty"`
	// The custom prompt control parameter object.
	CustomPrompt *CreateTaskRequestParametersCustomPrompt `json:"CustomPrompt,omitempty" xml:"CustomPrompt,omitempty" type:"Struct"`
	// Specifies whether to enable the custom prompt feature. When enabled, you can enter a personalized custom prompt.
	//
	// example:
	//
	// false
	CustomPromptEnabled *bool `json:"CustomPromptEnabled,omitempty" xml:"CustomPromptEnabled,omitempty"`
	// The extra parameters. In most cases, you do not need to set this parameter.
	ExtraParams *CreateTaskRequestParametersExtraParams `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty" type:"Struct"`
	// The identity recognition parameter object.
	IdentityRecognition *CreateTaskRequestParametersIdentityRecognition `json:"IdentityRecognition,omitempty" xml:"IdentityRecognition,omitempty" type:"Struct"`
	// Specifies whether to enable the identity recognition feature.
	IdentityRecognitionEnabled *bool   `json:"IdentityRecognitionEnabled,omitempty" xml:"IdentityRecognitionEnabled,omitempty"`
	LlmOutputLanguage          *string `json:"LlmOutputLanguage,omitempty" xml:"LlmOutputLanguage,omitempty"`
	// The control parameters for the intelligent meeting notes feature, which supports algorithm processing for action items, keywords, and key content. If you enable MeetingAssistanceEnabled but do not specify algorithm types through MeetingAssistance, all types are called and returned by default.
	MeetingAssistance *CreateTaskRequestParametersMeetingAssistance `json:"MeetingAssistance,omitempty" xml:"MeetingAssistance,omitempty" type:"Struct"`
	// Specifies whether to enable the intelligent meeting notes feature. When enabled, results such as keywords, key content, and action items are generated.
	//
	// example:
	//
	// false
	MeetingAssistanceEnabled *bool   `json:"MeetingAssistanceEnabled,omitempty" xml:"MeetingAssistanceEnabled,omitempty"`
	Model                    *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// Specifies whether to enable PPT extraction and PPT summarization. When enabled, PPT frames are extracted from the video file and corresponding summaries are generated. Enable this parameter only for offline transcription when the source file is a video file. Results cannot be generated in real-time recording scenarios or offline transcription scenarios where the source file is audio only.
	//
	// example:
	//
	// false
	PptExtractionEnabled *bool `json:"PptExtractionEnabled,omitempty" xml:"PptExtractionEnabled,omitempty"`
	// The service inspection parameter object.
	ServiceInspection *CreateTaskRequestParametersServiceInspection `json:"ServiceInspection,omitempty" xml:"ServiceInspection,omitempty" type:"Struct"`
	// Specifies whether to enable the service inspection feature. Default value: false.
	ServiceInspectionEnabled *bool `json:"ServiceInspectionEnabled,omitempty" xml:"ServiceInspectionEnabled,omitempty"`
	// The summarization control parameters.
	Summarization *CreateTaskRequestParametersSummarization `json:"Summarization,omitempty" xml:"Summarization,omitempty" type:"Struct"`
	// Specifies whether to enable the summarization feature. When enabled, results such as full-text summaries and speaker summaries can be generated.
	//
	// example:
	//
	// false
	SummarizationEnabled *bool `json:"SummarizationEnabled,omitempty" xml:"SummarizationEnabled,omitempty"`
	// Specifies whether to enable the spoken-to-written text conversion feature.
	//
	// example:
	//
	// false
	TextPolishEnabled *bool `json:"TextPolishEnabled,omitempty" xml:"TextPolishEnabled,omitempty"`
	// The audio/video or audio stream transcoding module.
	Transcoding *CreateTaskRequestParametersTranscoding `json:"Transcoding,omitempty" xml:"Transcoding,omitempty" type:"Struct"`
	// The speech transcription control parameters.
	Transcription *CreateTaskRequestParametersTranscription `json:"Transcription,omitempty" xml:"Transcription,omitempty" type:"Struct"`
	// The translation control parameters.
	Translation *CreateTaskRequestParametersTranslation `json:"Translation,omitempty" xml:"Translation,omitempty" type:"Struct"`
	// Specifies whether to enable the translation feature.
	//
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
	if s.AutoChapters != nil {
		if err := s.AutoChapters.Validate(); err != nil {
			return err
		}
	}
	if s.ContentExtraction != nil {
		if err := s.ContentExtraction.Validate(); err != nil {
			return err
		}
	}
	if s.CustomPrompt != nil {
		if err := s.CustomPrompt.Validate(); err != nil {
			return err
		}
	}
	if s.ExtraParams != nil {
		if err := s.ExtraParams.Validate(); err != nil {
			return err
		}
	}
	if s.IdentityRecognition != nil {
		if err := s.IdentityRecognition.Validate(); err != nil {
			return err
		}
	}
	if s.MeetingAssistance != nil {
		if err := s.MeetingAssistance.Validate(); err != nil {
			return err
		}
	}
	if s.ServiceInspection != nil {
		if err := s.ServiceInspection.Validate(); err != nil {
			return err
		}
	}
	if s.Summarization != nil {
		if err := s.Summarization.Validate(); err != nil {
			return err
		}
	}
	if s.Transcoding != nil {
		if err := s.Transcoding.Validate(); err != nil {
			return err
		}
	}
	if s.Transcription != nil {
		if err := s.Transcription.Validate(); err != nil {
			return err
		}
	}
	if s.Translation != nil {
		if err := s.Translation.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	// The list of extraction dimensions for conversation content extraction, including the name and definition of each extraction item.
	ExtractionContents []*CreateTaskRequestParametersContentExtractionExtractionContents `json:"ExtractionContents,omitempty" xml:"ExtractionContents,omitempty" type:"Repeated"`
	// The scene description for conversation content extraction.
	SceneIntroduction *string                `json:"SceneIntroduction,omitempty" xml:"SceneIntroduction,omitempty"`
	SpeakerMap        map[string]interface{} `json:"SpeakerMap,omitempty" xml:"SpeakerMap,omitempty"`
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
	if s.ExtractionContents != nil {
		for _, item := range s.ExtractionContents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateTaskRequestParametersContentExtractionExtractionContents struct {
	// The extraction dimension definition for conversation content extraction.
	Content  *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Identity *string `json:"Identity,omitempty" xml:"Identity,omitempty"`
	// The extraction dimension name for conversation content extraction.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
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
	// The list of custom prompt parameters.
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
	if s.Contents != nil {
		for _, item := range s.Contents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateTaskRequestParametersCustomPromptContents struct {
	// The model specified for the prompt.
	//
	// example:
	//
	// tingwu-turbo
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The custom name of the prompt, used to match output results.
	//
	// This parameter is required.
	//
	// example:
	//
	// summary-demo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The custom content of the prompt.
	//
	// This parameter is required.
	//
	// example:
	//
	// Summarize the following conversation:{Transcription}
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// The format of the {Transcription} tag.
	//
	// example:
	//
	// default
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
	DomainEducationEnabled *bool `json:"DomainEducationEnabled,omitempty" xml:"DomainEducationEnabled,omitempty"`
	// The return format of the full-text summary.
	FullTextSummaryFormat *string `json:"FullTextSummaryFormat,omitempty" xml:"FullTextSummaryFormat,omitempty"`
	// The number of keywords to extract.
	MaxKeywords *int32 `json:"MaxKeywords,omitempty" xml:"MaxKeywords,omitempty"`
	// Specifies whether to enable Nfix. In most cases, you do not need to set this parameter.
	//
	// example:
	//
	// true
	NfixEnabled              *bool `json:"NfixEnabled,omitempty" xml:"NfixEnabled,omitempty"`
	OcrAuxiliaryEnabled      *bool `json:"OcrAuxiliaryEnabled,omitempty" xml:"OcrAuxiliaryEnabled,omitempty"`
	TranslateLlmSceneEnabled *bool `json:"TranslateLlmSceneEnabled,omitempty" xml:"TranslateLlmSceneEnabled,omitempty"`
	// The translation hotword configuration.
	TranslationHotwordMap *CreateTaskRequestParametersExtraParamsTranslationHotwordMap `json:"TranslationHotwordMap,omitempty" xml:"TranslationHotwordMap,omitempty" type:"Struct"`
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

func (s *CreateTaskRequestParametersExtraParams) GetTranslationHotwordMap() *CreateTaskRequestParametersExtraParamsTranslationHotwordMap {
	return s.TranslationHotwordMap
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

func (s *CreateTaskRequestParametersExtraParams) SetTranslationHotwordMap(v *CreateTaskRequestParametersExtraParamsTranslationHotwordMap) *CreateTaskRequestParametersExtraParams {
	s.TranslationHotwordMap = v
	return s
}

func (s *CreateTaskRequestParametersExtraParams) Validate() error {
	if s.TranslationHotwordMap != nil {
		if err := s.TranslationHotwordMap.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTaskRequestParametersExtraParamsTranslationHotwordMap struct {
	// The business scenario type.
	BizType   *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	BizUserId *string `json:"bizUserId,omitempty" xml:"bizUserId,omitempty"`
}

func (s CreateTaskRequestParametersExtraParamsTranslationHotwordMap) String() string {
	return dara.Prettify(s)
}

func (s CreateTaskRequestParametersExtraParamsTranslationHotwordMap) GoString() string {
	return s.String()
}

func (s *CreateTaskRequestParametersExtraParamsTranslationHotwordMap) GetBizType() *string {
	return s.BizType
}

func (s *CreateTaskRequestParametersExtraParamsTranslationHotwordMap) GetBizUserId() *string {
	return s.BizUserId
}

func (s *CreateTaskRequestParametersExtraParamsTranslationHotwordMap) SetBizType(v string) *CreateTaskRequestParametersExtraParamsTranslationHotwordMap {
	s.BizType = &v
	return s
}

func (s *CreateTaskRequestParametersExtraParamsTranslationHotwordMap) SetBizUserId(v string) *CreateTaskRequestParametersExtraParamsTranslationHotwordMap {
	s.BizUserId = &v
	return s
}

func (s *CreateTaskRequestParametersExtraParamsTranslationHotwordMap) Validate() error {
	return dara.Validate(s)
}

type CreateTaskRequestParametersIdentityRecognition struct {
	// The list of identity contents for identity recognition, including the identity name and description.
	IdentityContents []*CreateTaskRequestParametersIdentityRecognitionIdentityContents `json:"IdentityContents,omitempty" xml:"IdentityContents,omitempty" type:"Repeated"`
	// The scene description for identity recognition.
	SceneIntroduction *string `json:"SceneIntroduction,omitempty" xml:"SceneIntroduction,omitempty"`
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
	if s.IdentityContents != nil {
		for _, item := range s.IdentityContents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateTaskRequestParametersIdentityRecognitionIdentityContents struct {
	// The identity description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The identity name.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// When the intelligent meeting notes feature is enabled, pass in the expected feature parameter types. Supported types: action items (Actions) and key information (KeyInformation). Key information includes keywords and key content (key sentences).
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
	// The list of inspection dimensions for service inspection, including the dimension name and definition. The definition specifies the criteria that the large language model uses to determine whether a dimension is matched.
	InspectionContents []*CreateTaskRequestParametersServiceInspectionInspectionContents `json:"InspectionContents,omitempty" xml:"InspectionContents,omitempty" type:"Repeated"`
	// The description of the inspection target and focus for service inspection.
	InspectionIntroduction *string `json:"InspectionIntroduction,omitempty" xml:"InspectionIntroduction,omitempty"`
	// The conversation scene description for service inspection.
	SceneIntroduction *string                `json:"SceneIntroduction,omitempty" xml:"SceneIntroduction,omitempty"`
	SpeakerMap        map[string]interface{} `json:"SpeakerMap,omitempty" xml:"SpeakerMap,omitempty"`
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
	if s.InspectionContents != nil {
		for _, item := range s.InspectionContents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateTaskRequestParametersServiceInspectionInspectionContents struct {
	// The inspection dimension definition for service inspection.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The inspection dimension name for service inspection.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
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
	// When the summarization feature is enabled, pass in the expected summarization types. Supported types: full-text summary (Paragraph), speaker summary (Conversational), and Q&A review summary (QuestionsAnswering).
	//
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
	// Specifies whether to generate an audio waveform from the original audio/video file or audio stream and save it. Currently, only MP3 format is supported. This parameter is optional when creating offline file transcription or real-time meetings.
	//
	// example:
	//
	// false
	SpectrumEnabled *bool `json:"SpectrumEnabled,omitempty" xml:"SpectrumEnabled,omitempty"`
	// Specifies whether to convert the original audio/video file or audio stream to MP3 format for storage. Currently, only MP3 format is supported. This parameter is optional when creating offline file transcription or real-time meetings.
	//
	// example:
	//
	// mp3
	TargetAudioFormat *string `json:"TargetAudioFormat,omitempty" xml:"TargetAudioFormat,omitempty"`
	// Specifies whether to convert the original video file to MP4 format for storage. Currently, only MP4 format is supported. This parameter is meaningful only when creating offline file transcription and the original file is in video format. Typically, you do not need to set this parameter.
	//
	// example:
	//
	// mp4
	TargetVideoFormat *string `json:"TargetVideoFormat,omitempty" xml:"TargetVideoFormat,omitempty"`
	// Specifies whether to extract video thumbnails from the original video file and save them. This parameter is meaningful only when creating offline file transcription and the original file is in video format. Typically, you do not need to set this parameter.
	//
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
	// Sets the output level for speech recognition results of the active speaker in real-time recording scenarios.
	//
	// - **1**: Returns results when a complete sentence is recognized.
	//
	// - **2**: Returns results for both intermediate results and complete sentences.
	//
	// Set this parameter as needed only in real-time recording scenarios when MultipleStreamsEnabled is set to true. This parameter does not need to be set for offline transcription scenarios.
	//
	// example:
	//
	// 1
	AdditionalStreamOutputLevel *int32 `json:"AdditionalStreamOutputLevel,omitempty" xml:"AdditionalStreamOutputLevel,omitempty"`
	// Specifies whether to enable audio event detection during speech transcription to determine whether events such as music exist in the audio.
	//
	// example:
	//
	// false
	AudioEventDetectionEnabled *bool `json:"AudioEventDetectionEnabled,omitempty" xml:"AudioEventDetectionEnabled,omitempty"`
	// The speaker diarization parameters.
	Diarization *CreateTaskRequestParametersTranscriptionDiarization `json:"Diarization,omitempty" xml:"Diarization,omitempty" type:"Struct"`
	// Specifies whether to enable speaker diarization.
	//
	// example:
	//
	// false
	DiarizationEnabled *bool `json:"DiarizationEnabled,omitempty" xml:"DiarizationEnabled,omitempty"`
	// Specifies whether to enable disfluency removal during speech transcription. Enabled by default.
	DisfluencyEnabled *bool `json:"DisfluencyEnabled,omitempty" xml:"DisfluencyEnabled,omitempty"`
	// Sets the speech transcription model to improve transcription accuracy in specific domains.
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// Sets the output level for speech recognition results. Default value: 1.
	//
	// - **1**: Returns results when a complete sentence is recognized.
	//
	// - **2**: Returns results for both intermediate results and complete sentences.
	//
	// example:
	//
	// 2
	OutputLevel *int32                 `json:"OutputLevel,omitempty" xml:"OutputLevel,omitempty"`
	Phrase      map[string]interface{} `json:"Phrase,omitempty" xml:"Phrase,omitempty"`
	// The vocabulary ID of the hot words.
	//
	// example:
	//
	// ce9c2a34b6d847bf92a77d0a196f****
	PhraseId *string `json:"PhraseId,omitempty" xml:"PhraseId,omitempty"`
	// Specifies whether to enable profanity filtering during speech transcription. Enabled by default.
	ProfanityFilterEnabled     *bool `json:"ProfanityFilterEnabled,omitempty" xml:"ProfanityFilterEnabled,omitempty"`
	RealtimeDiarizationEnabled *bool `json:"RealtimeDiarizationEnabled,omitempty" xml:"RealtimeDiarizationEnabled,omitempty"`
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

func (s *CreateTaskRequestParametersTranscription) GetDisfluencyEnabled() *bool {
	return s.DisfluencyEnabled
}

func (s *CreateTaskRequestParametersTranscription) GetModel() *string {
	return s.Model
}

func (s *CreateTaskRequestParametersTranscription) GetOutputLevel() *int32 {
	return s.OutputLevel
}

func (s *CreateTaskRequestParametersTranscription) GetPhrase() map[string]interface{} {
	return s.Phrase
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

func (s *CreateTaskRequestParametersTranscription) SetDisfluencyEnabled(v bool) *CreateTaskRequestParametersTranscription {
	s.DisfluencyEnabled = &v
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

func (s *CreateTaskRequestParametersTranscription) SetPhrase(v map[string]interface{}) *CreateTaskRequestParametersTranscription {
	s.Phrase = v
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
	if s.Diarization != nil {
		if err := s.Diarization.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateTaskRequestParametersTranscriptionDiarization struct {
	// Sets the speaker diarization parameter.
	//
	// If not set: speaker role differentiation is not used.
	//
	// 0: the number of speakers is undetermined.
	//
	// 2: the number of speakers is 2.
	//
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
	// Sets the output level for translation results of the active speaker in real-time recording scenarios.
	//
	// - **1**: Returns results when a complete sentence is recognized.
	//
	// - **2**: Returns results for both intermediate results and complete sentences.
	//
	// Set this parameter as needed only in real-time recording scenarios when MultipleStreamsEnabled is set to true. This parameter does not need to be set for offline transcription scenarios.
	//
	// example:
	//
	// 1
	AdditionalStreamOutputLevel *int32 `json:"AdditionalStreamOutputLevel,omitempty" xml:"AdditionalStreamOutputLevel,omitempty"`
	// Sets the output level for real-time translation results. Default value: 1.
	//
	// - **1**: Returns results when a complete sentence is recognized.
	//
	// - **2**: Returns results for both intermediate results and complete sentences.
	//
	// Set this parameter as needed only in real-time recording scenarios. This parameter does not need to be set for offline transcription scenarios.
	//
	// example:
	//
	// 2
	OutputLevel *int32 `json:"OutputLevel,omitempty" xml:"OutputLevel,omitempty"`
	// The target languages to set when the translation feature is enabled. Chinese, English, and Japanese are supported.
	TargetLanguages []*string `json:"TargetLanguages,omitempty" xml:"TargetLanguages,omitempty" type:"Repeated"`
	// Specifies whether to use large language model-based translation. Default value: false.
	TranslateLlmSceneEnabled *bool `json:"TranslateLlmSceneEnabled,omitempty" xml:"TranslateLlmSceneEnabled,omitempty"`
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

func (s *CreateTaskRequestParametersTranslation) GetTranslateLlmSceneEnabled() *bool {
	return s.TranslateLlmSceneEnabled
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

func (s *CreateTaskRequestParametersTranslation) SetTranslateLlmSceneEnabled(v bool) *CreateTaskRequestParametersTranslation {
	s.TranslateLlmSceneEnabled = &v
	return s
}

func (s *CreateTaskRequestParametersTranslation) Validate() error {
	return dara.Validate(s)
}
