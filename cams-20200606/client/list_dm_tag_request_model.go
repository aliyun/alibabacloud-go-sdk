// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDmTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListDmTagRequest
	GetKeyword() *string
	SetOwnerId(v int64) *ListDmTagRequest
	GetOwnerId() *int64
	SetPageIndex(v int64) *ListDmTagRequest
	GetPageIndex() *int64
	SetPageSize(v int64) *ListDmTagRequest
	GetPageSize() *int64
	SetResourceOwnerAccount(v string) *ListDmTagRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListDmTagRequest
	GetResourceOwnerId() *int64
}

type ListDmTagRequest struct {
	// example:
	//
	// a
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageIndex *int64 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize             *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListDmTagRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDmTagRequest) GoString() string {
	return s.String()
}

func (s *ListDmTagRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListDmTagRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListDmTagRequest) GetPageIndex() *int64 {
	return s.PageIndex
}

func (s *ListDmTagRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListDmTagRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListDmTagRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListDmTagRequest) SetKeyword(v string) *ListDmTagRequest {
	s.Keyword = &v
	return s
}

func (s *ListDmTagRequest) SetOwnerId(v int64) *ListDmTagRequest {
	s.OwnerId = &v
	return s
}

func (s *ListDmTagRequest) SetPageIndex(v int64) *ListDmTagRequest {
	s.PageIndex = &v
	return s
}

func (s *ListDmTagRequest) SetPageSize(v int64) *ListDmTagRequest {
	s.PageSize = &v
	return s
}

func (s *ListDmTagRequest) SetResourceOwnerAccount(v string) *ListDmTagRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListDmTagRequest) SetResourceOwnerId(v int64) *ListDmTagRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListDmTagRequest) Validate() error {
	return dara.Validate(s)
}
