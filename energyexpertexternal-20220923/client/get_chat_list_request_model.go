// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetChatListRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *GetChatListRequest
	GetPageSize() *int32
	SetSessionId(v string) *GetChatListRequest
	GetSessionId() *string
}

type GetChatListRequest struct {
	// Pagination parameter, page number, starting from 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Q&A session ID, used to record multiple Q&As for the same user.
	//
	// This parameter is required.
	//
	// example:
	//
	// bfce2248-1546-4298-8bcf-70ac26e69646
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
}

func (s GetChatListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetChatListRequest) GoString() string {
	return s.String()
}

func (s *GetChatListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetChatListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetChatListRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *GetChatListRequest) SetCurrentPage(v int32) *GetChatListRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetChatListRequest) SetPageSize(v int32) *GetChatListRequest {
	s.PageSize = &v
	return s
}

func (s *GetChatListRequest) SetSessionId(v string) *GetChatListRequest {
	s.SessionId = &v
	return s
}

func (s *GetChatListRequest) Validate() error {
	return dara.Validate(s)
}
