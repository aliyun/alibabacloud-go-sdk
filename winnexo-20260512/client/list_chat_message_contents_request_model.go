// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatMessageContentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxCharsPerAttachment(v int32) *ListChatMessageContentsRequest
	GetMaxCharsPerAttachment() *int32
	SetPage(v int32) *ListChatMessageContentsRequest
	GetPage() *int32
	SetPageSize(v int32) *ListChatMessageContentsRequest
	GetPageSize() *int32
	SetSourceId(v string) *ListChatMessageContentsRequest
	GetSourceId() *string
	SetTenantId(v string) *ListChatMessageContentsRequest
	GetTenantId() *string
}

type ListChatMessageContentsRequest struct {
	// The maximum length of the attachment content.
	//
	// example:
	//
	// 100000
	MaxCharsPerAttachment *int32 `json:"maxCharsPerAttachment,omitempty" xml:"maxCharsPerAttachment,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The ID of the data source.
	//
	// This parameter is required.
	//
	// example:
	//
	// source-chat-1
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The ID of the effective tenant.
	//
	// example:
	//
	// 10001
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ListChatMessageContentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListChatMessageContentsRequest) GoString() string {
	return s.String()
}

func (s *ListChatMessageContentsRequest) GetMaxCharsPerAttachment() *int32 {
	return s.MaxCharsPerAttachment
}

func (s *ListChatMessageContentsRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListChatMessageContentsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListChatMessageContentsRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *ListChatMessageContentsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListChatMessageContentsRequest) SetMaxCharsPerAttachment(v int32) *ListChatMessageContentsRequest {
	s.MaxCharsPerAttachment = &v
	return s
}

func (s *ListChatMessageContentsRequest) SetPage(v int32) *ListChatMessageContentsRequest {
	s.Page = &v
	return s
}

func (s *ListChatMessageContentsRequest) SetPageSize(v int32) *ListChatMessageContentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListChatMessageContentsRequest) SetSourceId(v string) *ListChatMessageContentsRequest {
	s.SourceId = &v
	return s
}

func (s *ListChatMessageContentsRequest) SetTenantId(v string) *ListChatMessageContentsRequest {
	s.TenantId = &v
	return s
}

func (s *ListChatMessageContentsRequest) Validate() error {
	return dara.Validate(s)
}
