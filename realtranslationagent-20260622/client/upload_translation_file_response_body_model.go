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
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *UploadTranslationFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
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
	CreditBreakdown      *string              `json:"CreditBreakdown,omitempty" xml:"CreditBreakdown,omitempty"`
	CreditsAvailable     *bool                `json:"CreditsAvailable,omitempty" xml:"CreditsAvailable,omitempty"`
	DetectedLang         *string              `json:"DetectedLang,omitempty" xml:"DetectedLang,omitempty"`
	EstimatedCostCredits *float64             `json:"EstimatedCostCredits,omitempty" xml:"EstimatedCostCredits,omitempty"`
	EstimatedTime        *int64               `json:"EstimatedTime,omitempty" xml:"EstimatedTime,omitempty"`
	Fonts                map[string][]*string `json:"Fonts,omitempty" xml:"Fonts,omitempty"`
	PageCount            *int64               `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	SensitiveDetected    *bool                `json:"SensitiveDetected,omitempty" xml:"SensitiveDetected,omitempty"`
	SensitiveTags        []*string            `json:"SensitiveTags,omitempty" xml:"SensitiveTags,omitempty" type:"Repeated"`
	TaskId               *string              `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	WordCount            *int64               `json:"WordCount,omitempty" xml:"WordCount,omitempty"`
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
