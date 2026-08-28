// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadTranslationFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UploadTranslationFileResponseBody
	GetCode() *string
	SetData(v *UploadTranslationFileResponseBodyData) *UploadTranslationFileResponseBody
	GetData() *UploadTranslationFileResponseBodyData
	SetMessage(v string) *UploadTranslationFileResponseBody
	GetMessage() *string
	SetRequestId(v string) *UploadTranslationFileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UploadTranslationFileResponseBody
	GetSuccess() *bool
}

type UploadTranslationFileResponseBody struct {
	// The return code.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data.
	Data *UploadTranslationFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F60AC23C-27A7-5376-9A68-0B6EF2D4F9E4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UploadTranslationFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadTranslationFileResponseBody) GoString() string {
	return s.String()
}

func (s *UploadTranslationFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *UploadTranslationFileResponseBody) GetData() *UploadTranslationFileResponseBodyData {
	return s.Data
}

func (s *UploadTranslationFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UploadTranslationFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadTranslationFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UploadTranslationFileResponseBody) SetCode(v string) *UploadTranslationFileResponseBody {
	s.Code = &v
	return s
}

func (s *UploadTranslationFileResponseBody) SetData(v *UploadTranslationFileResponseBodyData) *UploadTranslationFileResponseBody {
	s.Data = v
	return s
}

func (s *UploadTranslationFileResponseBody) SetMessage(v string) *UploadTranslationFileResponseBody {
	s.Message = &v
	return s
}

func (s *UploadTranslationFileResponseBody) SetRequestId(v string) *UploadTranslationFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadTranslationFileResponseBody) SetSuccess(v bool) *UploadTranslationFileResponseBody {
	s.Success = &v
	return s
}

func (s *UploadTranslationFileResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UploadTranslationFileResponseBodyData struct {
	// The description of the estimated Credits billing and time consumption for the task.
	//
	// example:
	//
	// docx: 0.002 creadits/word × 1000 words (original 2 characters rounded up to nearest thousand) / 1000 = 0.002
	CreditBreakdown *string `json:"CreditBreakdown,omitempty" xml:"CreditBreakdown,omitempty"`
	// Indicates whether the available Credits are sufficient for this translation task.
	//
	// - The estimate may be affected by various factors and may deviate. The actual result is based on the task submission operation.
	//
	// example:
	//
	// True
	CreditsAvailable *bool `json:"CreditsAvailable,omitempty" xml:"CreditsAvailable,omitempty"`
	// The detected language type.
	//
	// example:
	//
	// zh
	DetectedLang *string `json:"DetectedLang,omitempty" xml:"DetectedLang,omitempty"`
	// The estimated Credits consumption.
	//
	// example:
	//
	// 3.0021
	EstimatedCostCredits *float64 `json:"EstimatedCostCredits,omitempty" xml:"EstimatedCostCredits,omitempty"`
	// The estimated translation time in **seconds**.
	//
	// example:
	//
	// 60000
	EstimatedTime *int64 `json:"EstimatedTime,omitempty" xml:"EstimatedTime,omitempty"`
	// The languages that support font modification and the corresponding font lists. The key of the map identifies the language type.
	//
	// - Currently supported fonts include: English, French, Indonesian, and Japanese.
	Fonts map[string][]*string `json:"Fonts,omitempty" xml:"Fonts,omitempty"`
	// The page count of the uploaded file.
	//
	// example:
	//
	// 10
	PageCount *int64 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	// Indicates whether sensitive information was detected.
	//
	// example:
	//
	// True
	SensitiveDetected *bool `json:"SensitiveDetected,omitempty" xml:"SensitiveDetected,omitempty"`
	// The keywords that matched sensitive information.
	SensitiveTags []*string `json:"SensitiveTags,omitempty" xml:"SensitiveTags,omitempty" type:"Repeated"`
	// The translation task ID, used for subsequent task submission.
	//
	// example:
	//
	// f9c35b0453b
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The word count of the uploaded file.
	//
	// example:
	//
	// 2000
	WordCount *int64 `json:"WordCount,omitempty" xml:"WordCount,omitempty"`
}

func (s UploadTranslationFileResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UploadTranslationFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *UploadTranslationFileResponseBodyData) GetCreditBreakdown() *string {
	return s.CreditBreakdown
}

func (s *UploadTranslationFileResponseBodyData) GetCreditsAvailable() *bool {
	return s.CreditsAvailable
}

func (s *UploadTranslationFileResponseBodyData) GetDetectedLang() *string {
	return s.DetectedLang
}

func (s *UploadTranslationFileResponseBodyData) GetEstimatedCostCredits() *float64 {
	return s.EstimatedCostCredits
}

func (s *UploadTranslationFileResponseBodyData) GetEstimatedTime() *int64 {
	return s.EstimatedTime
}

func (s *UploadTranslationFileResponseBodyData) GetFonts() map[string][]*string {
	return s.Fonts
}

func (s *UploadTranslationFileResponseBodyData) GetPageCount() *int64 {
	return s.PageCount
}

func (s *UploadTranslationFileResponseBodyData) GetSensitiveDetected() *bool {
	return s.SensitiveDetected
}

func (s *UploadTranslationFileResponseBodyData) GetSensitiveTags() []*string {
	return s.SensitiveTags
}

func (s *UploadTranslationFileResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *UploadTranslationFileResponseBodyData) GetWordCount() *int64 {
	return s.WordCount
}

func (s *UploadTranslationFileResponseBodyData) SetCreditBreakdown(v string) *UploadTranslationFileResponseBodyData {
	s.CreditBreakdown = &v
	return s
}

func (s *UploadTranslationFileResponseBodyData) SetCreditsAvailable(v bool) *UploadTranslationFileResponseBodyData {
	s.CreditsAvailable = &v
	return s
}

func (s *UploadTranslationFileResponseBodyData) SetDetectedLang(v string) *UploadTranslationFileResponseBodyData {
	s.DetectedLang = &v
	return s
}

func (s *UploadTranslationFileResponseBodyData) SetEstimatedCostCredits(v float64) *UploadTranslationFileResponseBodyData {
	s.EstimatedCostCredits = &v
	return s
}

func (s *UploadTranslationFileResponseBodyData) SetEstimatedTime(v int64) *UploadTranslationFileResponseBodyData {
	s.EstimatedTime = &v
	return s
}

func (s *UploadTranslationFileResponseBodyData) SetFonts(v map[string][]*string) *UploadTranslationFileResponseBodyData {
	s.Fonts = v
	return s
}

func (s *UploadTranslationFileResponseBodyData) SetPageCount(v int64) *UploadTranslationFileResponseBodyData {
	s.PageCount = &v
	return s
}

func (s *UploadTranslationFileResponseBodyData) SetSensitiveDetected(v bool) *UploadTranslationFileResponseBodyData {
	s.SensitiveDetected = &v
	return s
}

func (s *UploadTranslationFileResponseBodyData) SetSensitiveTags(v []*string) *UploadTranslationFileResponseBodyData {
	s.SensitiveTags = v
	return s
}

func (s *UploadTranslationFileResponseBodyData) SetTaskId(v string) *UploadTranslationFileResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *UploadTranslationFileResponseBodyData) SetWordCount(v int64) *UploadTranslationFileResponseBodyData {
	s.WordCount = &v
	return s
}

func (s *UploadTranslationFileResponseBodyData) Validate() error {
	return dara.Validate(s)
}
