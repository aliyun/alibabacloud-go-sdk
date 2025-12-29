// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomQAResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *ListCustomQAResponseBody
	GetMessage() *string
	SetPage(v *ListCustomQAResponseBodyPage) *ListCustomQAResponseBody
	GetPage() *ListCustomQAResponseBodyPage
	SetRequestId(v string) *ListCustomQAResponseBody
	GetRequestId() *string
	SetResult(v []*ListCustomQAResponseBodyResult) *ListCustomQAResponseBody
	GetResult() []*ListCustomQAResponseBodyResult
	SetStatusCode(v int32) *ListCustomQAResponseBody
	GetStatusCode() *int32
}

type ListCustomQAResponseBody struct {
	// example:
	//
	// success
	Message *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	Page    *ListCustomQAResponseBodyPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// example:
	//
	// 0EC7***726E
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListCustomQAResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s ListCustomQAResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCustomQAResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomQAResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCustomQAResponseBody) GetPage() *ListCustomQAResponseBodyPage {
	return s.Page
}

func (s *ListCustomQAResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCustomQAResponseBody) GetResult() []*ListCustomQAResponseBodyResult {
	return s.Result
}

func (s *ListCustomQAResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCustomQAResponseBody) SetMessage(v string) *ListCustomQAResponseBody {
	s.Message = &v
	return s
}

func (s *ListCustomQAResponseBody) SetPage(v *ListCustomQAResponseBodyPage) *ListCustomQAResponseBody {
	s.Page = v
	return s
}

func (s *ListCustomQAResponseBody) SetRequestId(v string) *ListCustomQAResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomQAResponseBody) SetResult(v []*ListCustomQAResponseBodyResult) *ListCustomQAResponseBody {
	s.Result = v
	return s
}

func (s *ListCustomQAResponseBody) SetStatusCode(v int32) *ListCustomQAResponseBody {
	s.StatusCode = &v
	return s
}

func (s *ListCustomQAResponseBody) Validate() error {
	if s.Page != nil {
		if err := s.Page.Validate(); err != nil {
			return err
		}
	}
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCustomQAResponseBodyPage struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 21
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListCustomQAResponseBodyPage) String() string {
	return dara.Prettify(s)
}

func (s ListCustomQAResponseBodyPage) GoString() string {
	return s.String()
}

func (s *ListCustomQAResponseBodyPage) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCustomQAResponseBodyPage) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomQAResponseBodyPage) GetTotal() *int32 {
	return s.Total
}

func (s *ListCustomQAResponseBodyPage) SetPageNumber(v int32) *ListCustomQAResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *ListCustomQAResponseBodyPage) SetPageSize(v int32) *ListCustomQAResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *ListCustomQAResponseBodyPage) SetTotal(v int32) *ListCustomQAResponseBodyPage {
	s.Total = &v
	return s
}

func (s *ListCustomQAResponseBodyPage) Validate() error {
	return dara.Validate(s)
}

type ListCustomQAResponseBodyResult struct {
	// example:
	//
	// 22;11
	Answers *string `json:"Answers,omitempty" xml:"Answers,omitempty"`
	// example:
	//
	// 2023-01-10 10:01:59
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 111
	CustomQAId *string `json:"CustomQAId,omitempty" xml:"CustomQAId,omitempty"`
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// 22;11
	KeyWords *string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty"`
	// example:
	//
	// ***
	MajorQuestion *string `json:"MajorQuestion,omitempty" xml:"MajorQuestion,omitempty"`
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 22;11
	SupplementaryQuestion *string `json:"SupplementaryQuestion,omitempty" xml:"SupplementaryQuestion,omitempty"`
	// example:
	//
	// 2023-01-10 10:01:59
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListCustomQAResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListCustomQAResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListCustomQAResponseBodyResult) GetAnswers() *string {
	return s.Answers
}

func (s *ListCustomQAResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListCustomQAResponseBodyResult) GetCustomQAId() *string {
	return s.CustomQAId
}

func (s *ListCustomQAResponseBodyResult) GetHotelId() *string {
	return s.HotelId
}

func (s *ListCustomQAResponseBodyResult) GetKeyWords() *string {
	return s.KeyWords
}

func (s *ListCustomQAResponseBodyResult) GetMajorQuestion() *string {
	return s.MajorQuestion
}

func (s *ListCustomQAResponseBodyResult) GetStatus() *int32 {
	return s.Status
}

func (s *ListCustomQAResponseBodyResult) GetSupplementaryQuestion() *string {
	return s.SupplementaryQuestion
}

func (s *ListCustomQAResponseBodyResult) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListCustomQAResponseBodyResult) SetAnswers(v string) *ListCustomQAResponseBodyResult {
	s.Answers = &v
	return s
}

func (s *ListCustomQAResponseBodyResult) SetCreateTime(v string) *ListCustomQAResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListCustomQAResponseBodyResult) SetCustomQAId(v string) *ListCustomQAResponseBodyResult {
	s.CustomQAId = &v
	return s
}

func (s *ListCustomQAResponseBodyResult) SetHotelId(v string) *ListCustomQAResponseBodyResult {
	s.HotelId = &v
	return s
}

func (s *ListCustomQAResponseBodyResult) SetKeyWords(v string) *ListCustomQAResponseBodyResult {
	s.KeyWords = &v
	return s
}

func (s *ListCustomQAResponseBodyResult) SetMajorQuestion(v string) *ListCustomQAResponseBodyResult {
	s.MajorQuestion = &v
	return s
}

func (s *ListCustomQAResponseBodyResult) SetStatus(v int32) *ListCustomQAResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListCustomQAResponseBodyResult) SetSupplementaryQuestion(v string) *ListCustomQAResponseBodyResult {
	s.SupplementaryQuestion = &v
	return s
}

func (s *ListCustomQAResponseBodyResult) SetUpdateTime(v string) *ListCustomQAResponseBodyResult {
	s.UpdateTime = &v
	return s
}

func (s *ListCustomQAResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
