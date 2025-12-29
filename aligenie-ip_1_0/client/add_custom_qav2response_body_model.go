// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomQAV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *AddCustomQAV2ResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddCustomQAV2ResponseBody
	GetRequestId() *string
	SetResult(v *AddCustomQAV2ResponseBodyResult) *AddCustomQAV2ResponseBody
	GetResult() *AddCustomQAV2ResponseBodyResult
	SetStatusCode(v int32) *AddCustomQAV2ResponseBody
	GetStatusCode() *int32
}

type AddCustomQAV2ResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// FAFCD152-4791-5F2F-B0BE-2DC06FD4F05B
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *AddCustomQAV2ResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s AddCustomQAV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddCustomQAV2ResponseBody) GoString() string {
	return s.String()
}

func (s *AddCustomQAV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddCustomQAV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddCustomQAV2ResponseBody) GetResult() *AddCustomQAV2ResponseBodyResult {
	return s.Result
}

func (s *AddCustomQAV2ResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddCustomQAV2ResponseBody) SetMessage(v string) *AddCustomQAV2ResponseBody {
	s.Message = &v
	return s
}

func (s *AddCustomQAV2ResponseBody) SetRequestId(v string) *AddCustomQAV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCustomQAV2ResponseBody) SetResult(v *AddCustomQAV2ResponseBodyResult) *AddCustomQAV2ResponseBody {
	s.Result = v
	return s
}

func (s *AddCustomQAV2ResponseBody) SetStatusCode(v int32) *AddCustomQAV2ResponseBody {
	s.StatusCode = &v
	return s
}

func (s *AddCustomQAV2ResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddCustomQAV2ResponseBodyResult struct {
	Answers *string `json:"Answers,omitempty" xml:"Answers,omitempty"`
	// example:
	//
	// 2023-01-10 10:01:59
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// a7***83
	HotelId  *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	KeyWords *string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty"`
	// example:
	//
	// 8xxx9
	LastOperator  *string `json:"LastOperator,omitempty" xml:"LastOperator,omitempty"`
	MajorQuestion *string `json:"MajorQuestion,omitempty" xml:"MajorQuestion,omitempty"`
	// qaID
	//
	// example:
	//
	// 1
	QaId *string `json:"QaId,omitempty" xml:"QaId,omitempty"`
	// example:
	//
	// 0
	Status                *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	SupplementaryQuestion *string `json:"SupplementaryQuestion,omitempty" xml:"SupplementaryQuestion,omitempty"`
	// example:
	//
	// 2023-01-10 10:01:59
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s AddCustomQAV2ResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s AddCustomQAV2ResponseBodyResult) GoString() string {
	return s.String()
}

func (s *AddCustomQAV2ResponseBodyResult) GetAnswers() *string {
	return s.Answers
}

func (s *AddCustomQAV2ResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *AddCustomQAV2ResponseBodyResult) GetHotelId() *string {
	return s.HotelId
}

func (s *AddCustomQAV2ResponseBodyResult) GetKeyWords() *string {
	return s.KeyWords
}

func (s *AddCustomQAV2ResponseBodyResult) GetLastOperator() *string {
	return s.LastOperator
}

func (s *AddCustomQAV2ResponseBodyResult) GetMajorQuestion() *string {
	return s.MajorQuestion
}

func (s *AddCustomQAV2ResponseBodyResult) GetQaId() *string {
	return s.QaId
}

func (s *AddCustomQAV2ResponseBodyResult) GetStatus() *int32 {
	return s.Status
}

func (s *AddCustomQAV2ResponseBodyResult) GetSupplementaryQuestion() *string {
	return s.SupplementaryQuestion
}

func (s *AddCustomQAV2ResponseBodyResult) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *AddCustomQAV2ResponseBodyResult) SetAnswers(v string) *AddCustomQAV2ResponseBodyResult {
	s.Answers = &v
	return s
}

func (s *AddCustomQAV2ResponseBodyResult) SetCreateTime(v string) *AddCustomQAV2ResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *AddCustomQAV2ResponseBodyResult) SetHotelId(v string) *AddCustomQAV2ResponseBodyResult {
	s.HotelId = &v
	return s
}

func (s *AddCustomQAV2ResponseBodyResult) SetKeyWords(v string) *AddCustomQAV2ResponseBodyResult {
	s.KeyWords = &v
	return s
}

func (s *AddCustomQAV2ResponseBodyResult) SetLastOperator(v string) *AddCustomQAV2ResponseBodyResult {
	s.LastOperator = &v
	return s
}

func (s *AddCustomQAV2ResponseBodyResult) SetMajorQuestion(v string) *AddCustomQAV2ResponseBodyResult {
	s.MajorQuestion = &v
	return s
}

func (s *AddCustomQAV2ResponseBodyResult) SetQaId(v string) *AddCustomQAV2ResponseBodyResult {
	s.QaId = &v
	return s
}

func (s *AddCustomQAV2ResponseBodyResult) SetStatus(v int32) *AddCustomQAV2ResponseBodyResult {
	s.Status = &v
	return s
}

func (s *AddCustomQAV2ResponseBodyResult) SetSupplementaryQuestion(v string) *AddCustomQAV2ResponseBodyResult {
	s.SupplementaryQuestion = &v
	return s
}

func (s *AddCustomQAV2ResponseBodyResult) SetUpdateTime(v string) *AddCustomQAV2ResponseBodyResult {
	s.UpdateTime = &v
	return s
}

func (s *AddCustomQAV2ResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
