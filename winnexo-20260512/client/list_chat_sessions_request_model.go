// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatSessionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDigitalEmployeeName(v string) *ListChatSessionsRequest
	GetDigitalEmployeeName() *string
	SetKeyword(v string) *ListChatSessionsRequest
	GetKeyword() *string
	SetPage(v int32) *ListChatSessionsRequest
	GetPage() *int32
	SetPageSize(v int32) *ListChatSessionsRequest
	GetPageSize() *int32
	SetTenantId(v string) *ListChatSessionsRequest
	GetTenantId() *string
}

type ListChatSessionsRequest struct {
	// The list of digital employee names. A single string can be passed for backward compatibility with the legacy format.
	//
	// example:
	//
	// string_value
	DigitalEmployeeName *string `json:"digitalEmployeeName,omitempty" xml:"digitalEmployeeName,omitempty"`
	// The rule name keyword for fuzzy match.
	//
	// example:
	//
	// SampleKeyword
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The page number. Default value: 1. Pages start from page 1.
	//
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// The maximum number of data records to read in this request.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The tenant ID to which the task belongs.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListChatSessionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListChatSessionsRequest) GoString() string {
	return s.String()
}

func (s *ListChatSessionsRequest) GetDigitalEmployeeName() *string {
	return s.DigitalEmployeeName
}

func (s *ListChatSessionsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListChatSessionsRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListChatSessionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListChatSessionsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListChatSessionsRequest) SetDigitalEmployeeName(v string) *ListChatSessionsRequest {
	s.DigitalEmployeeName = &v
	return s
}

func (s *ListChatSessionsRequest) SetKeyword(v string) *ListChatSessionsRequest {
	s.Keyword = &v
	return s
}

func (s *ListChatSessionsRequest) SetPage(v int32) *ListChatSessionsRequest {
	s.Page = &v
	return s
}

func (s *ListChatSessionsRequest) SetPageSize(v int32) *ListChatSessionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListChatSessionsRequest) SetTenantId(v string) *ListChatSessionsRequest {
	s.TenantId = &v
	return s
}

func (s *ListChatSessionsRequest) Validate() error {
	return dara.Validate(s)
}
