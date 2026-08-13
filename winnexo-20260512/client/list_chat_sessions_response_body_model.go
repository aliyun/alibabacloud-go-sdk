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
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 是否有更多数据
	//
	// example:
	//
	// true
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	Page    *int32  `json:"page,omitempty" xml:"page,omitempty"`
	// 每页条数
	//
	// example:
	//
	// 20
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 响应数据负载
	//
	// example:
	//
	// {}
	Sessions []interface{} `json:"sessions,omitempty" xml:"sessions,omitempty" type:"Repeated"`
	// 租户ID
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	Total    *int32  `json:"total,omitempty" xml:"total,omitempty"`
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
