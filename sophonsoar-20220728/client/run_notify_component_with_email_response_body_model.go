// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunNotifyComponentWithEmailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *RunNotifyComponentWithEmailResponseBody
	GetData() *string
	SetPage(v *RunNotifyComponentWithEmailResponseBodyPage) *RunNotifyComponentWithEmailResponseBody
	GetPage() *RunNotifyComponentWithEmailResponseBodyPage
	SetRequestId(v string) *RunNotifyComponentWithEmailResponseBody
	GetRequestId() *string
}

type RunNotifyComponentWithEmailResponseBody struct {
	// The data returned.
	//
	// example:
	//
	// {}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The pagination information.
	Page *RunNotifyComponentWithEmailResponseBodyPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// D4CC979E-3D5B-5A6A-BC87-C93C9E861C7B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunNotifyComponentWithEmailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunNotifyComponentWithEmailResponseBody) GoString() string {
	return s.String()
}

func (s *RunNotifyComponentWithEmailResponseBody) GetData() *string {
	return s.Data
}

func (s *RunNotifyComponentWithEmailResponseBody) GetPage() *RunNotifyComponentWithEmailResponseBodyPage {
	return s.Page
}

func (s *RunNotifyComponentWithEmailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunNotifyComponentWithEmailResponseBody) SetData(v string) *RunNotifyComponentWithEmailResponseBody {
	s.Data = &v
	return s
}

func (s *RunNotifyComponentWithEmailResponseBody) SetPage(v *RunNotifyComponentWithEmailResponseBodyPage) *RunNotifyComponentWithEmailResponseBody {
	s.Page = v
	return s
}

func (s *RunNotifyComponentWithEmailResponseBody) SetRequestId(v string) *RunNotifyComponentWithEmailResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunNotifyComponentWithEmailResponseBody) Validate() error {
	if s.Page != nil {
		if err := s.Page.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunNotifyComponentWithEmailResponseBodyPage struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s RunNotifyComponentWithEmailResponseBodyPage) String() string {
	return dara.Prettify(s)
}

func (s RunNotifyComponentWithEmailResponseBodyPage) GoString() string {
	return s.String()
}

func (s *RunNotifyComponentWithEmailResponseBodyPage) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *RunNotifyComponentWithEmailResponseBodyPage) GetPageSize() *int32 {
	return s.PageSize
}

func (s *RunNotifyComponentWithEmailResponseBodyPage) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *RunNotifyComponentWithEmailResponseBodyPage) SetPageNumber(v int32) *RunNotifyComponentWithEmailResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *RunNotifyComponentWithEmailResponseBodyPage) SetPageSize(v int32) *RunNotifyComponentWithEmailResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *RunNotifyComponentWithEmailResponseBodyPage) SetTotalCount(v int32) *RunNotifyComponentWithEmailResponseBodyPage {
	s.TotalCount = &v
	return s
}

func (s *RunNotifyComponentWithEmailResponseBodyPage) Validate() error {
	return dara.Validate(s)
}
