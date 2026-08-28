// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTranslationTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTranslationTaskResponseBody
	GetCode() *string
	SetData(v *GetTranslationTaskResponseBodyData) *GetTranslationTaskResponseBody
	GetData() *GetTranslationTaskResponseBodyData
	SetMessage(v string) *GetTranslationTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTranslationTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTranslationTaskResponseBody
	GetSuccess() *bool
}

type GetTranslationTaskResponseBody struct {
	// The return code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data.
	Data *GetTranslationTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The return message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EDD51FD8-93E0-5161-BCA6-38A8393F26D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTranslationTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTranslationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetTranslationTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTranslationTaskResponseBody) GetData() *GetTranslationTaskResponseBodyData {
	return s.Data
}

func (s *GetTranslationTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTranslationTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTranslationTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTranslationTaskResponseBody) SetCode(v string) *GetTranslationTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetTranslationTaskResponseBody) SetData(v *GetTranslationTaskResponseBodyData) *GetTranslationTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetTranslationTaskResponseBody) SetMessage(v string) *GetTranslationTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetTranslationTaskResponseBody) SetRequestId(v string) *GetTranslationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTranslationTaskResponseBody) SetSuccess(v bool) *GetTranslationTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetTranslationTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTranslationTaskResponseBodyData struct {
	// The translation task ID of a previously submitted task. This parameter is passed in when resubmitting a translation task.
	//
	// example:
	//
	// f9c35b0453b
	BaseTaskId *string `json:"BaseTaskId,omitempty" xml:"BaseTaskId,omitempty"`
	// The translation configuration.
	Config *GetTranslationTaskResponseBodyDataConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// The credits consumed by this task.
	//
	// example:
	//
	// 2.5510
	CostCredits *float64 `json:"CostCredits,omitempty" xml:"CostCredits,omitempty"`
	// The time consumed, in milliseconds.
	//
	// example:
	//
	// 43
	CostTime *int64 `json:"CostTime,omitempty" xml:"CostTime,omitempty"`
	// The error message when the task fails.
	//
	// example:
	//
	// error
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The estimated credits to be consumed.
	//
	// example:
	//
	// 2.5510
	EstimatedCostCredits *float64 `json:"EstimatedCostCredits,omitempty" xml:"EstimatedCostCredits,omitempty"`
	// The estimated translation time, in **seconds**.
	//
	// example:
	//
	// 40000
	EstimatedTime *int64 `json:"EstimatedTime,omitempty" xml:"EstimatedTime,omitempty"`
	// The terms used in this task.
	ExtractedTerms []*GetTranslationTaskResponseBodyDataExtractedTerms `json:"ExtractedTerms,omitempty" xml:"ExtractedTerms,omitempty" type:"Repeated"`
	// The parsed file format.
	//
	// example:
	//
	// PPTX
	FileFormat *string `json:"FileFormat,omitempty" xml:"FileFormat,omitempty"`
	// The file name.
	//
	// example:
	//
	// translated_a_file.pptx
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The task completion time, expressed as a 13-digit timestamp.
	//
	// example:
	//
	// 1774147442
	FinishedAt *string `json:"FinishedAt,omitempty" xml:"FinishedAt,omitempty"`
	// The languages that support font modification and the corresponding font lists. The key of the map identifies the language type. Currently supported languages include English, French, Indonesian, and Japanese.
	Fonts map[string][]*string `json:"Fonts,omitempty" xml:"Fonts,omitempty"`
	// The organization ID.
	//
	// example:
	//
	// org_e5b74*****9c94209
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// The source file address.
	//
	// example:
	//
	// translated_a_file.pptx
	OriginalFileName *string `json:"OriginalFileName,omitempty" xml:"OriginalFileName,omitempty"`
	// The page count of the uploaded file.
	//
	// example:
	//
	// 0
	PageCount *int64 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	// The progress, expressed as a percentage number.
	//
	// example:
	//
	// 90
	Progress *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The task status.
	//
	// example:
	//
	// PROCESSING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The translation task ID.
	//
	// example:
	//
	// f9c35b0453b
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task type.
	//
	// - DOCUMENT: Document type.
	//
	// example:
	//
	// DOCUMENT
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// The word count of the uploaded document.
	//
	// example:
	//
	// 1600
	WordCount *int64 `json:"WordCount,omitempty" xml:"WordCount,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// c2b898f******c985c
	WorkSpaceId *string `json:"WorkSpaceId,omitempty" xml:"WorkSpaceId,omitempty"`
}

func (s GetTranslationTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTranslationTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTranslationTaskResponseBodyData) GetBaseTaskId() *string {
	return s.BaseTaskId
}

func (s *GetTranslationTaskResponseBodyData) GetConfig() *GetTranslationTaskResponseBodyDataConfig {
	return s.Config
}

func (s *GetTranslationTaskResponseBodyData) GetCostCredits() *float64 {
	return s.CostCredits
}

func (s *GetTranslationTaskResponseBodyData) GetCostTime() *int64 {
	return s.CostTime
}

func (s *GetTranslationTaskResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetTranslationTaskResponseBodyData) GetEstimatedCostCredits() *float64 {
	return s.EstimatedCostCredits
}

func (s *GetTranslationTaskResponseBodyData) GetEstimatedTime() *int64 {
	return s.EstimatedTime
}

func (s *GetTranslationTaskResponseBodyData) GetExtractedTerms() []*GetTranslationTaskResponseBodyDataExtractedTerms {
	return s.ExtractedTerms
}

func (s *GetTranslationTaskResponseBodyData) GetFileFormat() *string {
	return s.FileFormat
}

func (s *GetTranslationTaskResponseBodyData) GetFileName() *string {
	return s.FileName
}

func (s *GetTranslationTaskResponseBodyData) GetFinishedAt() *string {
	return s.FinishedAt
}

func (s *GetTranslationTaskResponseBodyData) GetFonts() map[string][]*string {
	return s.Fonts
}

func (s *GetTranslationTaskResponseBodyData) GetOrgId() *string {
	return s.OrgId
}

func (s *GetTranslationTaskResponseBodyData) GetOriginalFileName() *string {
	return s.OriginalFileName
}

func (s *GetTranslationTaskResponseBodyData) GetPageCount() *int64 {
	return s.PageCount
}

func (s *GetTranslationTaskResponseBodyData) GetProgress() *int32 {
	return s.Progress
}

func (s *GetTranslationTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetTranslationTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTranslationTaskResponseBodyData) GetTaskType() *string {
	return s.TaskType
}

func (s *GetTranslationTaskResponseBodyData) GetWordCount() *int64 {
	return s.WordCount
}

func (s *GetTranslationTaskResponseBodyData) GetWorkSpaceId() *string {
	return s.WorkSpaceId
}

func (s *GetTranslationTaskResponseBodyData) SetBaseTaskId(v string) *GetTranslationTaskResponseBodyData {
	s.BaseTaskId = &v
	return s
}

func (s *GetTranslationTaskResponseBodyData) SetConfig(v *GetTranslationTaskResponseBodyDataConfig) *GetTranslationTaskResponseBodyData {
	s.Config = v
	return s
}

func (s *GetTranslationTaskResponseBodyData) SetCostCredits(v float64) *GetTranslationTaskResponseBodyData {
	s.CostCredits = &v
	return s
}

func (s *GetTranslationTaskResponseBodyData) SetCostTime(v int64) *GetTranslationTaskResponseBodyData {
	s.CostTime = &v
	return s
}

func (s *GetTranslationTaskResponseBodyData) SetErrorMessage(v string) *GetTranslationTaskResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetTranslationTaskResponseBodyData) SetEstimatedCostCredits(v float64) *GetTranslationTaskResponseBodyData {
	s.EstimatedCostCredits = &v
	return s
}

func (s *GetTranslationTaskResponseBodyData) SetEstimatedTime(v int64) *GetTranslationTaskResponseBodyData {
	s.EstimatedTime = &v
	return s
}

func (s *GetTranslationTaskResponseBodyData) SetExtractedTerms(v []*GetTranslationTaskResponseBodyDataExtractedTerms) *GetTranslationTaskResponseBodyData {
	s.ExtractedTerms = v
	return s
}

func (s *GetTranslationTaskResponseBodyData) SetFileFormat(v string) *GetTranslationTaskResponseBodyData {
	s.FileFormat = &v
	return s
}

func (s *GetTranslationTaskResponseBodyData) SetFileName(v string) *GetTranslationTaskResponseBodyData {
	s.FileName = &v
	return s
}

func (s *GetTranslationTaskResponseBodyData) SetFinishedAt(v string) *GetTranslationTaskResponseBodyData {
	s.FinishedAt = &v
	return s
}

func (s *GetTranslationTaskResponseBodyData) SetFonts(v map[string][]*string) *GetTranslationTaskResponseBodyData {
	s.Fonts = v
	return s
}

func (s *GetTranslationTaskResponseBodyData) SetOrgId(v string) *GetTranslationTaskResponseBodyData {
	s.OrgId = &v
	return s
}

func (s *GetTranslationTaskResponseBodyData) SetOriginalFileName(v string) *GetTranslationTaskResponseBodyData {
	s.OriginalFileName = &v
	return s
}

func (s *GetTranslationTaskResponseBodyData) SetPageCount(v int64) *GetTranslationTaskResponseBodyData {
	s.PageCount = &v
	return s
}

func (s *GetTranslationTaskResponseBodyData) SetProgress(v int32) *GetTranslationTaskResponseBodyData {
	s.Progress = &v
	return s
}

func (s *GetTranslationTaskResponseBodyData) SetStatus(v string) *GetTranslationTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetTranslationTaskResponseBodyData) SetTaskId(v string) *GetTranslationTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetTranslationTaskResponseBodyData) SetTaskType(v string) *GetTranslationTaskResponseBodyData {
	s.TaskType = &v
	return s
}

func (s *GetTranslationTaskResponseBodyData) SetWordCount(v int64) *GetTranslationTaskResponseBodyData {
	s.WordCount = &v
	return s
}

func (s *GetTranslationTaskResponseBodyData) SetWorkSpaceId(v string) *GetTranslationTaskResponseBodyData {
	s.WorkSpaceId = &v
	return s
}

func (s *GetTranslationTaskResponseBodyData) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	if s.ExtractedTerms != nil {
		for _, item := range s.ExtractedTerms {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTranslationTaskResponseBodyDataConfig struct {
	// The security level.
	//
	// - public: Standard confidentiality.
	//
	// example:
	//
	// public
	SecurityLevel *string `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	// The source file language.
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// The translation style. This parameter takes effect only when the translation file is a PPT file.
	//
	// - normal: Standard. The original information is fully preserved.
	//
	// - minimal: More concise information with a more visually appealing layout.
	//
	// example:
	//
	// minimal
	Style *string `json:"Style,omitempty" xml:"Style,omitempty"`
	// The target language.
	//
	// example:
	//
	// en
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	// The translation template.
	//
	// - common: General-purpose.
	//
	// example:
	//
	// common
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s GetTranslationTaskResponseBodyDataConfig) String() string {
	return dara.Prettify(s)
}

func (s GetTranslationTaskResponseBodyDataConfig) GoString() string {
	return s.String()
}

func (s *GetTranslationTaskResponseBodyDataConfig) GetSecurityLevel() *string {
	return s.SecurityLevel
}

func (s *GetTranslationTaskResponseBodyDataConfig) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *GetTranslationTaskResponseBodyDataConfig) GetStyle() *string {
	return s.Style
}

func (s *GetTranslationTaskResponseBodyDataConfig) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *GetTranslationTaskResponseBodyDataConfig) GetTemplate() *string {
	return s.Template
}

func (s *GetTranslationTaskResponseBodyDataConfig) SetSecurityLevel(v string) *GetTranslationTaskResponseBodyDataConfig {
	s.SecurityLevel = &v
	return s
}

func (s *GetTranslationTaskResponseBodyDataConfig) SetSourceLanguage(v string) *GetTranslationTaskResponseBodyDataConfig {
	s.SourceLanguage = &v
	return s
}

func (s *GetTranslationTaskResponseBodyDataConfig) SetStyle(v string) *GetTranslationTaskResponseBodyDataConfig {
	s.Style = &v
	return s
}

func (s *GetTranslationTaskResponseBodyDataConfig) SetTargetLanguage(v string) *GetTranslationTaskResponseBodyDataConfig {
	s.TargetLanguage = &v
	return s
}

func (s *GetTranslationTaskResponseBodyDataConfig) SetTemplate(v string) *GetTranslationTaskResponseBodyDataConfig {
	s.Template = &v
	return s
}

func (s *GetTranslationTaskResponseBodyDataConfig) Validate() error {
	return dara.Validate(s)
}

type GetTranslationTaskResponseBodyDataExtractedTerms struct {
	// The source term.
	//
	// example:
	//
	// puppy
	SourceTerm *string `json:"SourceTerm,omitempty" xml:"SourceTerm,omitempty"`
	// The translated term.
	//
	// example:
	//
	// dog
	TargetTerm *string `json:"TargetTerm,omitempty" xml:"TargetTerm,omitempty"`
}

func (s GetTranslationTaskResponseBodyDataExtractedTerms) String() string {
	return dara.Prettify(s)
}

func (s GetTranslationTaskResponseBodyDataExtractedTerms) GoString() string {
	return s.String()
}

func (s *GetTranslationTaskResponseBodyDataExtractedTerms) GetSourceTerm() *string {
	return s.SourceTerm
}

func (s *GetTranslationTaskResponseBodyDataExtractedTerms) GetTargetTerm() *string {
	return s.TargetTerm
}

func (s *GetTranslationTaskResponseBodyDataExtractedTerms) SetSourceTerm(v string) *GetTranslationTaskResponseBodyDataExtractedTerms {
	s.SourceTerm = &v
	return s
}

func (s *GetTranslationTaskResponseBodyDataExtractedTerms) SetTargetTerm(v string) *GetTranslationTaskResponseBodyDataExtractedTerms {
	s.TargetTerm = &v
	return s
}

func (s *GetTranslationTaskResponseBodyDataExtractedTerms) Validate() error {
	return dara.Validate(s)
}
