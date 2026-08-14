// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTaskInfoResponseBody
	GetCode() *string
	SetData(v *GetTaskInfoResponseBodyData) *GetTaskInfoResponseBody
	GetData() *GetTaskInfoResponseBodyData
	SetMessage(v string) *GetTaskInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTaskInfoResponseBody
	GetRequestId() *string
}

type GetTaskInfoResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Returned object.
	Data *GetTaskInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Status description.
	//
	// example:
	//
	// Success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID, used only for joint debugging.
	//
	// example:
	//
	// 35124E1C-AE99-5D6C-A52E-BD689D8D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTaskInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTaskInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTaskInfoResponseBody) GetData() *GetTaskInfoResponseBodyData {
	return s.Data
}

func (s *GetTaskInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTaskInfoResponseBody) GetRequestId() *string {
	return s.RequestId
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

func (s *GetTaskInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTaskInfoResponseBodyData struct {
	// Error code
	//
	// example:
	//
	// TSC.AudioFormat
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Error message
	//
	// example:
	//
	// Audio format invalid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// URL link to the MP3 conversion result
	//
	// example:
	//
	// http://xxxx.com/tingwu/output/1738248324/094e964bf0e04e39/094e964bf0e04e39_20231222101008.mp3?Expires=1706064016
	OutputMp3Path *string `json:"OutputMp3Path,omitempty" xml:"OutputMp3Path,omitempty"`
	// URL link to the MP4 conversion result
	//
	// example:
	//
	// http://xxxx.com/tingwu/output/1738248324/094e964bf0e04e39/094e964bf0e04e39_20231222101008.mp4?Expires=1706064016
	OutputMp4Path *string `json:"OutputMp4Path,omitempty" xml:"OutputMp4Path,omitempty"`
	// URL link to the audio waveform graph
	//
	// example:
	//
	// http://xxxx.com/tingwu/output/1738248324/094e964bf0e04e39/094e964bf0e04e39_20231222101008.spectrum?Expires=1706064016
	OutputSpectrumPath *string `json:"OutputSpectrumPath,omitempty" xml:"OutputSpectrumPath,omitempty"`
	// URL link to the video thumbnail
	//
	// example:
	//
	// http://xxxx.com/tingwu/output/1738248324/094e964bf0e04e39/094e964bf0e04e39_20231222101008.png?Expires=1706064016
	OutputThumbnailPath *string `json:"OutputThumbnailPath,omitempty" xml:"OutputThumbnailPath,omitempty"`
	// A collection of results from various algorithm processing tasks. The result is returned as an HTTP link, which the user can use to parse the native result.
	Result *GetTaskInfoResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Job ID.
	//
	// example:
	//
	// c5394c6ee0fb474899d42215a3925c7e
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The custom ID set by the user when creating the job.
	//
	// example:
	//
	// task_tingwu_123
	TaskKey *string `json:"TaskKey,omitempty" xml:"TaskKey,omitempty"`
	// Task Status.
	//
	// - ONGOING: The job is in progress.
	//
	// - COMPLETED: The job is completed.
	//
	// - FAILED: The job has failed.
	//
	// - INVALID: The job is invalid.
	//
	// example:
	//
	// COMPLETED
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s GetTaskInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTaskInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTaskInfoResponseBodyData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetTaskInfoResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetTaskInfoResponseBodyData) GetOutputMp3Path() *string {
	return s.OutputMp3Path
}

func (s *GetTaskInfoResponseBodyData) GetOutputMp4Path() *string {
	return s.OutputMp4Path
}

func (s *GetTaskInfoResponseBodyData) GetOutputSpectrumPath() *string {
	return s.OutputSpectrumPath
}

func (s *GetTaskInfoResponseBodyData) GetOutputThumbnailPath() *string {
	return s.OutputThumbnailPath
}

func (s *GetTaskInfoResponseBodyData) GetResult() *GetTaskInfoResponseBodyDataResult {
	return s.Result
}

func (s *GetTaskInfoResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTaskInfoResponseBodyData) GetTaskKey() *string {
	return s.TaskKey
}

func (s *GetTaskInfoResponseBodyData) GetTaskStatus() *string {
	return s.TaskStatus
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

func (s *GetTaskInfoResponseBodyData) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTaskInfoResponseBodyDataResult struct {
	// Link to the result of the Auto Chapters feature.
	//
	// example:
	//
	// http://xxxx.com/tingwu/output/1738248324/094e964bf0e04e39/094e964bf0e04e39_AutoChapters_20231222101215.json?Expires=1706064016
	AutoChapters *string `json:"AutoChapters,omitempty" xml:"AutoChapters,omitempty"`
	// URL link to the result of conversation content extraction
	//
	// example:
	//
	// http://xxxx.com/tingwu/output/1738248324/094e964bf0e04e39/094e964bf0e04e39_ ContentExtraction_20231222101215.json?Expires=1706064016
	ContentExtraction *string `json:"ContentExtraction,omitempty" xml:"ContentExtraction,omitempty"`
	// Link to the result of the Custom prompt
	//
	// example:
	//
	// http://xxxx.com/tingwu/output/1738248324/094e964bf0e04e39/094e964bf0e04e39_ CustomPrompt_20231222101215.json?Expires=1706064016
	CustomPrompt *string `json:"CustomPrompt,omitempty" xml:"CustomPrompt,omitempty"`
	// The URL link to the identity recognition result.
	//
	// example:
	//
	// http://xxxx.com/tingwu/output/1738248324/094e964bf0e04e39/094e964bf0e04e39_ IdentityRecognition_20231222101215.json?Expires=1706064016
	IdentityRecognition *string `json:"IdentityRecognition,omitempty" xml:"IdentityRecognition,omitempty"`
	// URL link to the result of Intelligent Meeting Summary
	//
	// example:
	//
	// http://xxxx.com/tingwu/output/1738248324/094e964bf0e04e39/094e964bf0e04e39_MeetingAssistance_20231222101112.json?Expires=1706064016
	MeetingAssistance *string `json:"MeetingAssistance,omitempty" xml:"MeetingAssistance,omitempty"`
	// URL link to the result of video PPT extraction and summarization
	//
	// example:
	//
	// http://xxxx.com/tingwu/output/1738248324/094e964bf0e04e39/094e964bf0e04e39_PptExtraction_20231222101215.json?Expires=1706064016
	PptExtraction *string `json:"PptExtraction,omitempty" xml:"PptExtraction,omitempty"`
	// Link to the result of service inspection
	//
	// example:
	//
	// http://xxxx.com/tingwu/output/1738248324/094e964bf0e04e39/094e964bf0e04e39_ ServiceInspection_20231222101215.json?Expires=1706064016
	ServiceInspection *string `json:"ServiceInspection,omitempty" xml:"ServiceInspection,omitempty"`
	// Link to the result of LLM-based summarization.
	//
	// example:
	//
	// http://xxxx.com/tingwu/output/1738248324/094e964bf0e04e39/094e964bf0e04e39_Summarization_20231222101215.json?Expires=1706064016
	Summarization *string `json:"Summarization,omitempty" xml:"Summarization,omitempty"`
	// Link to the result of spoken-to-written text conversion
	//
	// example:
	//
	// http://xxxx.com/tingwu/output/1738248324/094e964bf0e04e39/094e964bf0e04e39_TextPolish_20231222101215.json?Expires=1706064016
	TextPolish *string `json:"TextPolish,omitempty" xml:"TextPolish,omitempty"`
	// Link to the result of speech transcription.
	//
	// example:
	//
	// http://xxxx.com/tingwu/output/1738248324/094e964bf0e04e39/094e964bf0e04e39_Transcription_20231222101008.json?Expires=1706064016
	Transcription *string `json:"Transcription,omitempty" xml:"Transcription,omitempty"`
	// URL link to the result of text translation
	//
	// example:
	//
	// http://xxxx.com/tingwu/output/1738248324/094e964bf0e04e39/094e964bf0e04e39_Translation_20231222101215.json?Expires=1706064016
	Translation *string `json:"Translation,omitempty" xml:"Translation,omitempty"`
}

func (s GetTaskInfoResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s GetTaskInfoResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *GetTaskInfoResponseBodyDataResult) GetAutoChapters() *string {
	return s.AutoChapters
}

func (s *GetTaskInfoResponseBodyDataResult) GetContentExtraction() *string {
	return s.ContentExtraction
}

func (s *GetTaskInfoResponseBodyDataResult) GetCustomPrompt() *string {
	return s.CustomPrompt
}

func (s *GetTaskInfoResponseBodyDataResult) GetIdentityRecognition() *string {
	return s.IdentityRecognition
}

func (s *GetTaskInfoResponseBodyDataResult) GetMeetingAssistance() *string {
	return s.MeetingAssistance
}

func (s *GetTaskInfoResponseBodyDataResult) GetPptExtraction() *string {
	return s.PptExtraction
}

func (s *GetTaskInfoResponseBodyDataResult) GetServiceInspection() *string {
	return s.ServiceInspection
}

func (s *GetTaskInfoResponseBodyDataResult) GetSummarization() *string {
	return s.Summarization
}

func (s *GetTaskInfoResponseBodyDataResult) GetTextPolish() *string {
	return s.TextPolish
}

func (s *GetTaskInfoResponseBodyDataResult) GetTranscription() *string {
	return s.Transcription
}

func (s *GetTaskInfoResponseBodyDataResult) GetTranslation() *string {
	return s.Translation
}

func (s *GetTaskInfoResponseBodyDataResult) SetAutoChapters(v string) *GetTaskInfoResponseBodyDataResult {
	s.AutoChapters = &v
	return s
}

func (s *GetTaskInfoResponseBodyDataResult) SetContentExtraction(v string) *GetTaskInfoResponseBodyDataResult {
	s.ContentExtraction = &v
	return s
}

func (s *GetTaskInfoResponseBodyDataResult) SetCustomPrompt(v string) *GetTaskInfoResponseBodyDataResult {
	s.CustomPrompt = &v
	return s
}

func (s *GetTaskInfoResponseBodyDataResult) SetIdentityRecognition(v string) *GetTaskInfoResponseBodyDataResult {
	s.IdentityRecognition = &v
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

func (s *GetTaskInfoResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}
