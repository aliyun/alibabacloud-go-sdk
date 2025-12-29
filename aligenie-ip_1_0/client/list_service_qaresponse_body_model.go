// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceQAResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *ListServiceQAResponseBody
	GetMessage() *string
	SetPage(v *ListServiceQAResponseBodyPage) *ListServiceQAResponseBody
	GetPage() *ListServiceQAResponseBodyPage
	SetRequestId(v string) *ListServiceQAResponseBody
	GetRequestId() *string
	SetResult(v []*ListServiceQAResponseBodyResult) *ListServiceQAResponseBody
	GetResult() []*ListServiceQAResponseBodyResult
	SetStatusCode(v int32) *ListServiceQAResponseBody
	GetStatusCode() *int32
}

type ListServiceQAResponseBody struct {
	// example:
	//
	// success
	Message *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	Page    *ListServiceQAResponseBodyPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// example:
	//
	// 0EC7***726E
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*ListServiceQAResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s ListServiceQAResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServiceQAResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceQAResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListServiceQAResponseBody) GetPage() *ListServiceQAResponseBodyPage {
	return s.Page
}

func (s *ListServiceQAResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServiceQAResponseBody) GetResult() []*ListServiceQAResponseBodyResult {
	return s.Result
}

func (s *ListServiceQAResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServiceQAResponseBody) SetMessage(v string) *ListServiceQAResponseBody {
	s.Message = &v
	return s
}

func (s *ListServiceQAResponseBody) SetPage(v *ListServiceQAResponseBodyPage) *ListServiceQAResponseBody {
	s.Page = v
	return s
}

func (s *ListServiceQAResponseBody) SetRequestId(v string) *ListServiceQAResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceQAResponseBody) SetResult(v []*ListServiceQAResponseBodyResult) *ListServiceQAResponseBody {
	s.Result = v
	return s
}

func (s *ListServiceQAResponseBody) SetStatusCode(v int32) *ListServiceQAResponseBody {
	s.StatusCode = &v
	return s
}

func (s *ListServiceQAResponseBody) Validate() error {
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

type ListServiceQAResponseBodyPage struct {
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
	// 12
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListServiceQAResponseBodyPage) String() string {
	return dara.Prettify(s)
}

func (s ListServiceQAResponseBodyPage) GoString() string {
	return s.String()
}

func (s *ListServiceQAResponseBodyPage) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListServiceQAResponseBodyPage) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListServiceQAResponseBodyPage) GetTotal() *int32 {
	return s.Total
}

func (s *ListServiceQAResponseBodyPage) SetPageNumber(v int32) *ListServiceQAResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *ListServiceQAResponseBodyPage) SetPageSize(v int32) *ListServiceQAResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *ListServiceQAResponseBodyPage) SetTotal(v int32) *ListServiceQAResponseBodyPage {
	s.Total = &v
	return s
}

func (s *ListServiceQAResponseBodyPage) Validate() error {
	return dara.Validate(s)
}

type ListServiceQAResponseBodyResult struct {
	// example:
	//
	// true
	Active *bool `json:"Active,omitempty" xml:"Active,omitempty"`
	// example:
	//
	// ***
	Answer *string `json:"Answer,omitempty" xml:"Answer,omitempty"`
	// example:
	//
	// 2022-07-27 14:06:27
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Question    *string `json:"Question,omitempty" xml:"Question,omitempty"`
	// example:
	//
	// 1
	ServiceQAId *int64  `json:"ServiceQAId,omitempty" xml:"ServiceQAId,omitempty"`
	Templates   *string `json:"Templates,omitempty" xml:"Templates,omitempty"`
}

func (s ListServiceQAResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListServiceQAResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListServiceQAResponseBodyResult) GetActive() *bool {
	return s.Active
}

func (s *ListServiceQAResponseBodyResult) GetAnswer() *string {
	return s.Answer
}

func (s *ListServiceQAResponseBodyResult) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListServiceQAResponseBodyResult) GetQuestion() *string {
	return s.Question
}

func (s *ListServiceQAResponseBodyResult) GetServiceQAId() *int64 {
	return s.ServiceQAId
}

func (s *ListServiceQAResponseBodyResult) GetTemplates() *string {
	return s.Templates
}

func (s *ListServiceQAResponseBodyResult) SetActive(v bool) *ListServiceQAResponseBodyResult {
	s.Active = &v
	return s
}

func (s *ListServiceQAResponseBodyResult) SetAnswer(v string) *ListServiceQAResponseBodyResult {
	s.Answer = &v
	return s
}

func (s *ListServiceQAResponseBodyResult) SetGmtModified(v string) *ListServiceQAResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *ListServiceQAResponseBodyResult) SetQuestion(v string) *ListServiceQAResponseBodyResult {
	s.Question = &v
	return s
}

func (s *ListServiceQAResponseBodyResult) SetServiceQAId(v int64) *ListServiceQAResponseBodyResult {
	s.ServiceQAId = &v
	return s
}

func (s *ListServiceQAResponseBodyResult) SetTemplates(v string) *ListServiceQAResponseBodyResult {
	s.Templates = &v
	return s
}

func (s *ListServiceQAResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
