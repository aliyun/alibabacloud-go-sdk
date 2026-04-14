// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWhatAppTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListWhatAppTemplateResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListWhatAppTemplateResponseBody
	GetCode() *string
	SetData(v string) *ListWhatAppTemplateResponseBody
	GetData() *string
	SetMessage(v string) *ListWhatAppTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListWhatAppTemplateResponseBody
	GetRequestId() *string
	SetTotal(v int32) *ListWhatAppTemplateResponseBody
	GetTotal() *int32
}

type ListWhatAppTemplateResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// True
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4AAC06A8-1B23-5616-AAAE-A7E7D7E48ACB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListWhatAppTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWhatAppTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ListWhatAppTemplateResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListWhatAppTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListWhatAppTemplateResponseBody) GetData() *string {
	return s.Data
}

func (s *ListWhatAppTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListWhatAppTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWhatAppTemplateResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListWhatAppTemplateResponseBody) SetAccessDeniedDetail(v string) *ListWhatAppTemplateResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListWhatAppTemplateResponseBody) SetCode(v string) *ListWhatAppTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *ListWhatAppTemplateResponseBody) SetData(v string) *ListWhatAppTemplateResponseBody {
	s.Data = &v
	return s
}

func (s *ListWhatAppTemplateResponseBody) SetMessage(v string) *ListWhatAppTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *ListWhatAppTemplateResponseBody) SetRequestId(v string) *ListWhatAppTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWhatAppTemplateResponseBody) SetTotal(v int32) *ListWhatAppTemplateResponseBody {
	s.Total = &v
	return s
}

func (s *ListWhatAppTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
