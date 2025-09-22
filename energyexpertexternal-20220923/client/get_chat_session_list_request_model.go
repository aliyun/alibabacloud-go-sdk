// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatSessionListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *GetChatSessionListRequest
	GetCurrentPage() *int32
	SetName(v string) *GetChatSessionListRequest
	GetName() *string
	SetPageSize(v int32) *GetChatSessionListRequest
	GetPageSize() *int32
	SetUserId(v string) *GetChatSessionListRequest
	GetUserId() *string
}

type GetChatSessionListRequest struct {
	// Pagination parameter, page number, default is 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Session name.
	//
	// example:
	//
	// oklabs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Page size, default is 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The unique identifier of the user. If not provided, the SDK calling account will be used as the user ID by default.
	//
	// example:
	//
	// 12222
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetChatSessionListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetChatSessionListRequest) GoString() string {
	return s.String()
}

func (s *GetChatSessionListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetChatSessionListRequest) GetName() *string {
	return s.Name
}

func (s *GetChatSessionListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetChatSessionListRequest) GetUserId() *string {
	return s.UserId
}

func (s *GetChatSessionListRequest) SetCurrentPage(v int32) *GetChatSessionListRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetChatSessionListRequest) SetName(v string) *GetChatSessionListRequest {
	s.Name = &v
	return s
}

func (s *GetChatSessionListRequest) SetPageSize(v int32) *GetChatSessionListRequest {
	s.PageSize = &v
	return s
}

func (s *GetChatSessionListRequest) SetUserId(v string) *GetChatSessionListRequest {
	s.UserId = &v
	return s
}

func (s *GetChatSessionListRequest) Validate() error {
	return dara.Validate(s)
}
