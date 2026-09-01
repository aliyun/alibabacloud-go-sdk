// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunNotifyComponentWithMessageCenterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *RunNotifyComponentWithMessageCenterResponseBody
	GetData() *string
	SetPage(v *RunNotifyComponentWithMessageCenterResponseBodyPage) *RunNotifyComponentWithMessageCenterResponseBody
	GetPage() *RunNotifyComponentWithMessageCenterResponseBodyPage
	SetRequestId(v string) *RunNotifyComponentWithMessageCenterResponseBody
	GetRequestId() *string
}

type RunNotifyComponentWithMessageCenterResponseBody struct {
	// The data returned.
	//
	// example:
	//
	// {}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The pagination information.
	Page *RunNotifyComponentWithMessageCenterResponseBodyPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// The unique ID generated for the request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// E7698CFB-4E1C-5840-8EC9-691B86729E94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunNotifyComponentWithMessageCenterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunNotifyComponentWithMessageCenterResponseBody) GoString() string {
	return s.String()
}

func (s *RunNotifyComponentWithMessageCenterResponseBody) GetData() *string {
	return s.Data
}

func (s *RunNotifyComponentWithMessageCenterResponseBody) GetPage() *RunNotifyComponentWithMessageCenterResponseBodyPage {
	return s.Page
}

func (s *RunNotifyComponentWithMessageCenterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunNotifyComponentWithMessageCenterResponseBody) SetData(v string) *RunNotifyComponentWithMessageCenterResponseBody {
	s.Data = &v
	return s
}

func (s *RunNotifyComponentWithMessageCenterResponseBody) SetPage(v *RunNotifyComponentWithMessageCenterResponseBodyPage) *RunNotifyComponentWithMessageCenterResponseBody {
	s.Page = v
	return s
}

func (s *RunNotifyComponentWithMessageCenterResponseBody) SetRequestId(v string) *RunNotifyComponentWithMessageCenterResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunNotifyComponentWithMessageCenterResponseBody) Validate() error {
	if s.Page != nil {
		if err := s.Page.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunNotifyComponentWithMessageCenterResponseBodyPage struct {
	// The page number of the returned page.
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
	// The total number of entries.
	//
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s RunNotifyComponentWithMessageCenterResponseBodyPage) String() string {
	return dara.Prettify(s)
}

func (s RunNotifyComponentWithMessageCenterResponseBodyPage) GoString() string {
	return s.String()
}

func (s *RunNotifyComponentWithMessageCenterResponseBodyPage) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *RunNotifyComponentWithMessageCenterResponseBodyPage) GetPageSize() *int32 {
	return s.PageSize
}

func (s *RunNotifyComponentWithMessageCenterResponseBodyPage) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *RunNotifyComponentWithMessageCenterResponseBodyPage) SetPageNumber(v int32) *RunNotifyComponentWithMessageCenterResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *RunNotifyComponentWithMessageCenterResponseBodyPage) SetPageSize(v int32) *RunNotifyComponentWithMessageCenterResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *RunNotifyComponentWithMessageCenterResponseBodyPage) SetTotalCount(v int32) *RunNotifyComponentWithMessageCenterResponseBodyPage {
	s.TotalCount = &v
	return s
}

func (s *RunNotifyComponentWithMessageCenterResponseBodyPage) Validate() error {
	return dara.Validate(s)
}
