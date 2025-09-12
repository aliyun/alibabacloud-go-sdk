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
	return dara.Validate(s)
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
	return dara.Validate(s)
}

type GetTaskInfoResponseBodyDataResult struct {
	AutoChapters        *string `json:"AutoChapters,omitempty" xml:"AutoChapters,omitempty"`
	ContentExtraction   *string `json:"ContentExtraction,omitempty" xml:"ContentExtraction,omitempty"`
	CustomPrompt        *string `json:"CustomPrompt,omitempty" xml:"CustomPrompt,omitempty"`
	IdentityRecognition *string `json:"IdentityRecognition,omitempty" xml:"IdentityRecognition,omitempty"`
	MeetingAssistance   *string `json:"MeetingAssistance,omitempty" xml:"MeetingAssistance,omitempty"`
	PptExtraction       *string `json:"PptExtraction,omitempty" xml:"PptExtraction,omitempty"`
	ServiceInspection   *string `json:"ServiceInspection,omitempty" xml:"ServiceInspection,omitempty"`
	Summarization       *string `json:"Summarization,omitempty" xml:"Summarization,omitempty"`
	TextPolish          *string `json:"TextPolish,omitempty" xml:"TextPolish,omitempty"`
	Transcription       *string `json:"Transcription,omitempty" xml:"Transcription,omitempty"`
	Translation         *string `json:"Translation,omitempty" xml:"Translation,omitempty"`
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
