// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTicketTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetTicketTemplateResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *GetTicketTemplateResponseBody
	GetCode() *int32
	SetData(v map[string]interface{}) *GetTicketTemplateResponseBody
	GetData() map[string]interface{}
	SetHttpStatusCode(v int32) *GetTicketTemplateResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetTicketTemplateResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *GetTicketTemplateResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *GetTicketTemplateResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetTicketTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTicketTemplateResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *GetTicketTemplateResponseBody
	GetTotalCount() *int64
}

type GetTicketTemplateResponseBody struct {
	AccessDeniedDetail *string                `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *int32                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	HttpStatusCode     *int32                 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber         *int32                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int32                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId          *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount         *int64                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetTicketTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTicketTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetTicketTemplateResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetTicketTemplateResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetTicketTemplateResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *GetTicketTemplateResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetTicketTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTicketTemplateResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetTicketTemplateResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetTicketTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTicketTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTicketTemplateResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetTicketTemplateResponseBody) SetAccessDeniedDetail(v string) *GetTicketTemplateResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetTicketTemplateResponseBody) SetCode(v int32) *GetTicketTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *GetTicketTemplateResponseBody) SetData(v map[string]interface{}) *GetTicketTemplateResponseBody {
	s.Data = v
	return s
}

func (s *GetTicketTemplateResponseBody) SetHttpStatusCode(v int32) *GetTicketTemplateResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTicketTemplateResponseBody) SetMessage(v string) *GetTicketTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *GetTicketTemplateResponseBody) SetPageNumber(v int32) *GetTicketTemplateResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetTicketTemplateResponseBody) SetPageSize(v int32) *GetTicketTemplateResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetTicketTemplateResponseBody) SetRequestId(v string) *GetTicketTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTicketTemplateResponseBody) SetSuccess(v bool) *GetTicketTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *GetTicketTemplateResponseBody) SetTotalCount(v int64) *GetTicketTemplateResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetTicketTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
