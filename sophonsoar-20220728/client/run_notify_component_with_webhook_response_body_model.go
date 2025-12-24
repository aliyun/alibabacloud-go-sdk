// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunNotifyComponentWithWebhookResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *RunNotifyComponentWithWebhookResponseBody
	GetData() *string
	SetPage(v *RunNotifyComponentWithWebhookResponseBodyPage) *RunNotifyComponentWithWebhookResponseBody
	GetPage() *RunNotifyComponentWithWebhookResponseBodyPage
	SetRequestId(v string) *RunNotifyComponentWithWebhookResponseBody
	GetRequestId() *string
}

type RunNotifyComponentWithWebhookResponseBody struct {
	// The data returned.
	//
	// example:
	//
	// {}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The pagination information.
	Page *RunNotifyComponentWithWebhookResponseBodyPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// E7698CFB-****-5840-8EC9-691B86729E94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunNotifyComponentWithWebhookResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunNotifyComponentWithWebhookResponseBody) GoString() string {
	return s.String()
}

func (s *RunNotifyComponentWithWebhookResponseBody) GetData() *string {
	return s.Data
}

func (s *RunNotifyComponentWithWebhookResponseBody) GetPage() *RunNotifyComponentWithWebhookResponseBodyPage {
	return s.Page
}

func (s *RunNotifyComponentWithWebhookResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunNotifyComponentWithWebhookResponseBody) SetData(v string) *RunNotifyComponentWithWebhookResponseBody {
	s.Data = &v
	return s
}

func (s *RunNotifyComponentWithWebhookResponseBody) SetPage(v *RunNotifyComponentWithWebhookResponseBodyPage) *RunNotifyComponentWithWebhookResponseBody {
	s.Page = v
	return s
}

func (s *RunNotifyComponentWithWebhookResponseBody) SetRequestId(v string) *RunNotifyComponentWithWebhookResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunNotifyComponentWithWebhookResponseBody) Validate() error {
	if s.Page != nil {
		if err := s.Page.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunNotifyComponentWithWebhookResponseBodyPage struct {
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

func (s RunNotifyComponentWithWebhookResponseBodyPage) String() string {
	return dara.Prettify(s)
}

func (s RunNotifyComponentWithWebhookResponseBodyPage) GoString() string {
	return s.String()
}

func (s *RunNotifyComponentWithWebhookResponseBodyPage) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *RunNotifyComponentWithWebhookResponseBodyPage) GetPageSize() *int32 {
	return s.PageSize
}

func (s *RunNotifyComponentWithWebhookResponseBodyPage) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *RunNotifyComponentWithWebhookResponseBodyPage) SetPageNumber(v int32) *RunNotifyComponentWithWebhookResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *RunNotifyComponentWithWebhookResponseBodyPage) SetPageSize(v int32) *RunNotifyComponentWithWebhookResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *RunNotifyComponentWithWebhookResponseBodyPage) SetTotalCount(v int32) *RunNotifyComponentWithWebhookResponseBodyPage {
	s.TotalCount = &v
	return s
}

func (s *RunNotifyComponentWithWebhookResponseBodyPage) Validate() error {
	return dara.Validate(s)
}
