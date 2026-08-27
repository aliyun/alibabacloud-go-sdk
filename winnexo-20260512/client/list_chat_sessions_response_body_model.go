// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatSessionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListChatSessionsResponseBody
	GetCode() *string
	SetHasMore(v bool) *ListChatSessionsResponseBody
	GetHasMore() *bool
	SetMessage(v string) *ListChatSessionsResponseBody
	GetMessage() *string
	SetPage(v int32) *ListChatSessionsResponseBody
	GetPage() *int32
	SetPageSize(v string) *ListChatSessionsResponseBody
	GetPageSize() *string
	SetRequestId(v string) *ListChatSessionsResponseBody
	GetRequestId() *string
	SetSessions(v []interface{}) *ListChatSessionsResponseBody
	GetSessions() []interface{}
	SetTenantId(v string) *ListChatSessionsResponseBody
	GetTenantId() *string
	SetTotal(v int32) *ListChatSessionsResponseBody
	GetTotal() *int32
}

type ListChatSessionsResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Indicates whether there is a next page.
	//
	// example:
	//
	// true
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// The status code description.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 20
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The list of sessions.
	//
	// example:
	//
	// {}
	Sessions []interface{} `json:"sessions,omitempty" xml:"sessions,omitempty" type:"Repeated"`
	// The effective tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 1159
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListChatSessionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListChatSessionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListChatSessionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListChatSessionsResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *ListChatSessionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListChatSessionsResponseBody) GetPage() *int32 {
	return s.Page
}

func (s *ListChatSessionsResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *ListChatSessionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListChatSessionsResponseBody) GetSessions() []interface{} {
	return s.Sessions
}

func (s *ListChatSessionsResponseBody) GetTenantId() *string {
	return s.TenantId
}

func (s *ListChatSessionsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListChatSessionsResponseBody) SetCode(v string) *ListChatSessionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListChatSessionsResponseBody) SetHasMore(v bool) *ListChatSessionsResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListChatSessionsResponseBody) SetMessage(v string) *ListChatSessionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListChatSessionsResponseBody) SetPage(v int32) *ListChatSessionsResponseBody {
	s.Page = &v
	return s
}

func (s *ListChatSessionsResponseBody) SetPageSize(v string) *ListChatSessionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListChatSessionsResponseBody) SetRequestId(v string) *ListChatSessionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChatSessionsResponseBody) SetSessions(v []interface{}) *ListChatSessionsResponseBody {
	s.Sessions = v
	return s
}

func (s *ListChatSessionsResponseBody) SetTenantId(v string) *ListChatSessionsResponseBody {
	s.TenantId = &v
	return s
}

func (s *ListChatSessionsResponseBody) SetTotal(v int32) *ListChatSessionsResponseBody {
	s.Total = &v
	return s
}

func (s *ListChatSessionsResponseBody) Validate() error {
	return dara.Validate(s)
}
