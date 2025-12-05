// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConversationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLastId(v string) *GetConversationsRequest
	GetLastId() *string
	SetLimit(v string) *GetConversationsRequest
	GetLimit() *string
	SetPinned(v string) *GetConversationsRequest
	GetPinned() *string
	SetSortBy(v string) *GetConversationsRequest
	GetSortBy() *string
}

type GetConversationsRequest struct {
	// example:
	//
	// 77be60cd-237b-4ca9-9c46-48b663cb****
	LastId *string `json:"LastId,omitempty" xml:"LastId,omitempty"`
	// example:
	//
	// 10
	Limit *string `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// example:
	//
	// true
	Pinned *string `json:"Pinned,omitempty" xml:"Pinned,omitempty"`
	// example:
	//
	// CreatedAt
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s GetConversationsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetConversationsRequest) GoString() string {
	return s.String()
}

func (s *GetConversationsRequest) GetLastId() *string {
	return s.LastId
}

func (s *GetConversationsRequest) GetLimit() *string {
	return s.Limit
}

func (s *GetConversationsRequest) GetPinned() *string {
	return s.Pinned
}

func (s *GetConversationsRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *GetConversationsRequest) SetLastId(v string) *GetConversationsRequest {
	s.LastId = &v
	return s
}

func (s *GetConversationsRequest) SetLimit(v string) *GetConversationsRequest {
	s.Limit = &v
	return s
}

func (s *GetConversationsRequest) SetPinned(v string) *GetConversationsRequest {
	s.Pinned = &v
	return s
}

func (s *GetConversationsRequest) SetSortBy(v string) *GetConversationsRequest {
	s.SortBy = &v
	return s
}

func (s *GetConversationsRequest) Validate() error {
	return dara.Validate(s)
}
