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
	// The AppKey of the project that you created in the console.
	//
	// example:
	//
	// JV1sRTisRMi****
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// The basic input parameters for creating a task. The required parameters vary based on the task type.
	//
	// - For an offline task (`type="offline"`), you must specify the `SourceLanguage` and `FileUrl` parameters.
	//
	// - For a real-time task (`type="realtime"`), you must also specify the `SourceLanguage`, `Format`, and `SampleRate` parameters.
	Input *CreateTaskRequestInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// Algorithm-related parameters for customizing task processing.
	Parameters *CreateTaskRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Struct"`
	// The operation to perform. Valid values:
	//
	// - **start**: Creates a task. This is the default value and does not typically need to be set.
	//
	// - **stop**: Stops a real-time recording task. This value is used only for real-time tasks. To end the recording, set this parameter to `stop`.
	//
	// example:
	//
	// stop
	Operation *string `json:"operation,omitempty" xml:"operation,omitempty"`
	// The type of the task. Valid values:
	//
	// - **offline**: An offline task, such as an offline transcription.
	//
	// - **realtime**: A real-time task, such as a real-time recording.
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
	// Multi-channel audio or video processing mode.
	AudioChannelMode *string `json:"AudioChannelMode,omitempty" xml:"AudioChannelMode,omitempty"`
	// The HTTP or HTTPS URL of the source audio or video file. This parameter is required when you create an offline transcription task.
	//
	// example:
	//
	// http://xxx.com/zzz/1.wav
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// The encoding format of the audio stream data. This parameter is required when you create a real-time recording task. The following values are supported:
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
	// Preferred languages. This applies only when SourceLanguage is multilingual. It restricts the output language of the model.
	LanguageHints []*string `json:"LanguageHints,omitempty" xml:"LanguageHints,omitempty" type:"Repeated"`
	// Specifies whether to enable multi-channel audio stream recognition. This parameter applies only to real-time recording scenarios. The default value is `false`.
	//
	// example:
	//
	// false
	MultipleStreamsEnabled *bool `json:"MultipleStreamsEnabled,omitempty" xml:"MultipleStreamsEnabled,omitempty"`
	// After you configure OSS settings in the console, specify an OSS path to save results directly to your OSS bucket.
	OutputPath *string `json:"OutputPath,omitempty" xml:"OutputPath,omitempty"`
	// Specifies whether to enable callbacks. To receive callbacks, you must configure the callback type and URL in the console and set this parameter to `true`.
	//
	// example:
	//
	// false
	ProgressiveCallbacksEnabled *bool `json:"ProgressiveCallbacksEnabled,omitempty" xml:"ProgressiveCallbacksEnabled,omitempty"`
	// The sample rate of the audio stream data. This parameter is required when you create a real-time recording task. The supported values are 8000 and 16000.
	//
	// - **8000**: Suitable for telephony and customer service scenarios.
	//
	// - **16000**: Suitable for real-time meeting audio capture scenarios.
	//
	// example:
	//
	// 16000
	SampleRate *int32 `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
	// The language model for speech transcription. The following values are supported:
	//
	// - **cn**: Chinese
	//
	// - **en**: English
	//
	// - **fspk**: Chinese-English code-switching
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
	// The task ID that is returned when you create a real-time recording. This ID is required to stop the recording. Specify this parameter only when stopping a real-time recording.
	//
	// example:
	//
	// 9922c84c087044eda18659c128b56c84
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// A custom identifier that you can set for the task.
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
	// Specifies whether to generate a chapter summary, which includes chapter titles and summaries for each chapter.
	//
	// example:
	//
	// true
	AutoChaptersEnabled *bool `json:"AutoChaptersEnabled,omitempty" xml:"AutoChaptersEnabled,omitempty"`
	// Conversation content extraction parameters.
	ContentExtraction        *CreateTaskRequestParametersContentExtraction `json:"ContentExtraction,omitempty" xml:"ContentExtraction,omitempty" type:"Struct"`
	ContentExtractionEnabled *bool                                         `json:"ContentExtractionEnabled,omitempty" xml:"ContentExtractionEnabled,omitempty"`
	// Parameters to control the custom prompt feature.
	CustomPrompt *CreateTaskRequestParametersCustomPrompt `json:"CustomPrompt,omitempty" xml:"CustomPrompt,omitempty" type:"Struct"`
	// Specifies whether to enable the custom prompt feature.
	//
	// example:
	//
	// false
	CustomPromptEnabled *bool `json:"CustomPromptEnabled,omitempty" xml:"CustomPromptEnabled,omitempty"`
	// Extended parameters for advanced use cases. You do not typically need to configure these parameters.
	ExtraParams *CreateTaskRequestParametersExtraParams `json:"ExtraParams,omitempty" xml:"ExtraParams,omitempty" type:"Struct"`
	// Identity recognition parameters.
	IdentityRecognition *CreateTaskRequestParametersIdentityRecognition `json:"IdentityRecognition,omitempty" xml:"IdentityRecognition,omitempty" type:"Struct"`
	// Enable identity recognition.
	IdentityRecognitionEnabled *bool   `json:"IdentityRecognitionEnabled,omitempty" xml:"IdentityRecognitionEnabled,omitempty"`
	LlmOutputLanguage          *string `json:"LlmOutputLanguage,omitempty" xml:"LlmOutputLanguage,omitempty"`
	// Parameters for the intelligent minutes feature, which supports processing for action items, keywords, and key points. If `MeetingAssistanceEnabled` is set to `true` but you do not specify this object, all analysis types are enabled by default.
	MeetingAssistance *CreateTaskRequestParametersMeetingAssistance `json:"MeetingAssistance,omitempty" xml:"MeetingAssistance,omitempty" type:"Struct"`
	// Specifies whether to generate intelligent minutes, which include keywords, key points, and action items.
	//
	// example:
	//
	// false
	MeetingAssistanceEnabled *bool   `json:"MeetingAssistanceEnabled,omitempty" xml:"MeetingAssistanceEnabled,omitempty"`
	Model                    *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// Specifies whether to enable PPT extraction. If enabled, the service extracts slides from the video file and generates corresponding summaries. This feature applies only to offline transcription tasks with a video source file and has no effect on other task types.
	//
	// example:
	//
	// false
	PptExtractionEnabled *bool `json:"PptExtractionEnabled,omitempty" xml:"PptExtractionEnabled,omitempty"`
	// Service quality inspection parameters.
	ServiceInspection *CreateTaskRequestParametersServiceInspection `json:"ServiceInspection,omitempty" xml:"ServiceInspection,omitempty" type:"Struct"`
	// Enable service quality inspection. Default is false.
	ServiceInspectionEnabled *bool `json:"ServiceInspectionEnabled,omitempty" xml:"ServiceInspectionEnabled,omitempty"`
	// Parameters for the summarization feature.
	Summarization *CreateTaskRequestParametersSummarization `json:"Summarization,omitempty" xml:"Summarization,omitempty" type:"Struct"`
	// Specifies whether to enable the summarization feature, which can generate results such as a full-text summary and a speaker summary.
	//
	// example:
	//
	// false
	SummarizationEnabled *bool `json:"SummarizationEnabled,omitempty" xml:"SummarizationEnabled,omitempty"`
	// Specifies whether to enable the spoken-to-written conversion feature.
	//
	// example:
	//
	// false
	TextPolishEnabled *bool `json:"TextPolishEnabled,omitempty" xml:"TextPolishEnabled,omitempty"`
	// Parameters for transcoding source audio/video files or audio streams.
	Transcoding *CreateTaskRequestParametersTranscoding `json:"Transcoding,omitempty" xml:"Transcoding,omitempty" type:"Struct"`
	// Parameters to control the speech transcription process.
	Transcription *CreateTaskRequestParametersTranscription `json:"Transcription,omitempty" xml:"Transcription,omitempty" type:"Struct"`
	// Parameters to control the translation feature.
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
	// List of content extraction dimensions. Each dimension includes a name and definition.
	ExtractionContents []*CreateTaskRequestParametersContentExtractionExtractionContents `json:"ExtractionContents,omitempty" xml:"ExtractionContents,omitempty" type:"Repeated"`
	// Description of the conversation scenario for content extraction.
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
	// Definition of the content extraction dimension.
	Content  *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Identity *string `json:"Identity,omitempty" xml:"Identity,omitempty"`
	// Name of the content extraction dimension.
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
	// A list of custom prompt parameters.
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
	// The model to use for the prompt.
	//
	// example:
	//
	// tingwu-turbo
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// A custom name for the prompt, used to identify the corresponding output.
	//
	// This parameter is required.
	//
	// example:
	//
	// summary-demo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The content of the custom prompt.
	//
	// This parameter is required.
	//
	// example:
	//
	// 总结一下下面的对话内容:{Transcription}
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// Specifies the format for the `{Transcription}` tag.
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
	// Full-text summary format.
	FullTextSummaryFormat *string `json:"FullTextSummaryFormat,omitempty" xml:"FullTextSummaryFormat,omitempty"`
	// Maximum number of keywords.
	MaxKeywords *int32 `json:"MaxKeywords,omitempty" xml:"MaxKeywords,omitempty"`
	// Specifies whether to enable nfix. You do not typically need to configure this parameter.
	//
	// example:
	//
	// true
	NfixEnabled              *bool                                                        `json:"NfixEnabled,omitempty" xml:"NfixEnabled,omitempty"`
	OcrAuxiliaryEnabled      *bool                                                        `json:"OcrAuxiliaryEnabled,omitempty" xml:"OcrAuxiliaryEnabled,omitempty"`
	TranslateLlmSceneEnabled *bool                                                        `json:"TranslateLlmSceneEnabled,omitempty" xml:"TranslateLlmSceneEnabled,omitempty"`
	TranslationHotwordMap    *CreateTaskRequestParametersExtraParamsTranslationHotwordMap `json:"TranslationHotwordMap,omitempty" xml:"TranslationHotwordMap,omitempty" type:"Struct"`
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
	// List of identities, including identity name and description.
	IdentityContents []*CreateTaskRequestParametersIdentityRecognitionIdentityContents `json:"IdentityContents,omitempty" xml:"IdentityContents,omitempty" type:"Repeated"`
	// Description of the scenario for identity recognition.
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
	// Identity description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Identity name.
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
	// The types of analysis to perform when the intelligent minutes feature is enabled. Supported values: `Actions` (action items) and `KeyInformation` (key information, including keywords and key points).
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
	// List of inspection dimensions for service quality inspection. Each dimension includes a name and definition, which tells the Large Language Model how to evaluate whether the dimension is met.
	InspectionContents []*CreateTaskRequestParametersServiceInspectionInspectionContents `json:"InspectionContents,omitempty" xml:"InspectionContents,omitempty" type:"Repeated"`
	// Description of the inspection goals and focus areas for service quality inspection.
	InspectionIntroduction *string `json:"InspectionIntroduction,omitempty" xml:"InspectionIntroduction,omitempty"`
	// Description of the conversation scenario for service quality inspection.
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
	// Definition of the inspection dimension.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Name of the inspection dimension.
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
	// The types of summaries to generate. This parameter is required when summarization is enabled. Supported types include `Paragraph` (full-text summary), `Conversational` (speaker summary), and `QuestionsAnswering` (Q\\&A summary).
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
	// Specifies whether to generate and save an audio waveform from the source audio/video file or audio stream. This parameter is optional for offline transcription and real-time recording tasks.
	//
	// example:
	//
	// false
	SpectrumEnabled *bool `json:"SpectrumEnabled,omitempty" xml:"SpectrumEnabled,omitempty"`
	// Specifies the target format for the transcoded audio. Set to `mp3` to transcode the source audio into MP3 format for storage. This parameter is optional for offline transcription and real-time recording tasks.
	//
	// example:
	//
	// mp3
	TargetAudioFormat *string `json:"TargetAudioFormat,omitempty" xml:"TargetAudioFormat,omitempty"`
	// Specifies the target format for the transcoded video. Set to `mp4` to transcode the source video into MP4 format for storage. This parameter applies only to offline transcription tasks with a video source file.
	//
	// example:
	//
	// mp4
	TargetVideoFormat *string `json:"TargetVideoFormat,omitempty" xml:"TargetVideoFormat,omitempty"`
	// Specifies whether to extract and save video thumbnails from the source video file. This parameter applies only to offline transcription tasks with a video source file.
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
	// Specifies the level of detail for speech transcription results for the active speaker in a real-time recording scenario.
	//
	// - **1**: Returns results only when a complete sentence is recognized.
	//
	// - **2**: Returns both intermediate and final results as they are recognized.
	//
	// This parameter applies only to real-time recordings when `MultipleStreamsEnabled` is set to `true`.
	//
	// example:
	//
	// 1
	AdditionalStreamOutputLevel *int32 `json:"AdditionalStreamOutputLevel,omitempty" xml:"AdditionalStreamOutputLevel,omitempty"`
	// Specifies whether to enable sound event detection, which identifies non-speech events in the audio, such as music.
	//
	// example:
	//
	// false
	AudioEventDetectionEnabled *bool `json:"AudioEventDetectionEnabled,omitempty" xml:"AudioEventDetectionEnabled,omitempty"`
	// Parameters for the speaker diarization feature.
	Diarization *CreateTaskRequestParametersTranscriptionDiarization `json:"Diarization,omitempty" xml:"Diarization,omitempty" type:"Struct"`
	// Specifies whether to enable speaker diarization.
	//
	// example:
	//
	// false
	DiarizationEnabled *bool `json:"DiarizationEnabled,omitempty" xml:"DiarizationEnabled,omitempty"`
	DisfluencyEnabled  *bool `json:"DisfluencyEnabled,omitempty" xml:"DisfluencyEnabled,omitempty"`
	// Set the speech transcription model to improve accuracy for specific domains.
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// Specifies the level of detail for the speech transcription results. Default value: `1`.
	//
	// - **1**: Returns results only when a complete sentence is recognized.
	//
	// - **2**: Returns both intermediate and final results as they are recognized.
	//
	// example:
	//
	// 2
	OutputLevel *int32 `json:"OutputLevel,omitempty" xml:"OutputLevel,omitempty"`
	// The ID of the hotword list.
	//
	// example:
	//
	// ce9c2a34b6d847bf92a77d0a196f****
	PhraseId *string `json:"PhraseId,omitempty" xml:"PhraseId,omitempty"`
	// Enable sensitive word filtering during speech transcription. Enabled by default.
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
	// Specifies the number of speakers to identify.
	//
	// If this parameter is not set, speakers are not differentiated in the transcript.
	//
	// Set the value to `0` to identify an unknown number of speakers.
	//
	// Set the value to `2` to identify two speakers.
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
	// Specifies the level of detail for real-time translation results for the active speaker.
	//
	// - **1**: Returns results only for complete sentences.
	//
	// - **2**: Returns both intermediate and final results.
	//
	// This parameter applies only to real-time recordings when `MultipleStreamsEnabled` is set to `true`.
	//
	// example:
	//
	// 1
	AdditionalStreamOutputLevel *int32 `json:"AdditionalStreamOutputLevel,omitempty" xml:"AdditionalStreamOutputLevel,omitempty"`
	// Specifies the level of detail for real-time translation results. Default value: `1`.
	//
	// - **1**: Returns results only for complete sentences.
	//
	// - **2**: Returns both intermediate and final results.
	//
	// This parameter applies only to real-time recordings.
	//
	// example:
	//
	// 2
	OutputLevel *int32 `json:"OutputLevel,omitempty" xml:"OutputLevel,omitempty"`
	// The target languages for translation. This parameter is required if translation is enabled. Supported languages include Chinese, English, and Japanese.
	TargetLanguages          []*string `json:"TargetLanguages,omitempty" xml:"TargetLanguages,omitempty" type:"Repeated"`
	TranslateLlmSceneEnabled *bool     `json:"TranslateLlmSceneEnabled,omitempty" xml:"TranslateLlmSceneEnabled,omitempty"`
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
