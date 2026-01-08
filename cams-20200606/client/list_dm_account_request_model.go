// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDmAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListDmAccountRequest
	GetKeyword() *string
	SetOwnerId(v int64) *ListDmAccountRequest
	GetOwnerId() *int64
	SetPageIndex(v int64) *ListDmAccountRequest
	GetPageIndex() *int64
	SetPageSize(v int64) *ListDmAccountRequest
	GetPageSize() *int64
	SetResourceOwnerAccount(v string) *ListDmAccountRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListDmAccountRequest
	GetResourceOwnerId() *int64
	SetSendType(v string) *ListDmAccountRequest
	GetSendType() *string
}

type ListDmAccountRequest struct {
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
	// example:
	//
	// batch
	SendType *string `json:"SendType,omitempty" xml:"SendType,omitempty"`
}

func (s ListDmAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDmAccountRequest) GoString() string {
	return s.String()
}

func (s *ListDmAccountRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListDmAccountRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListDmAccountRequest) GetPageIndex() *int64 {
	return s.PageIndex
}

func (s *ListDmAccountRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListDmAccountRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListDmAccountRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListDmAccountRequest) GetSendType() *string {
	return s.SendType
}

func (s *ListDmAccountRequest) SetKeyword(v string) *ListDmAccountRequest {
	s.Keyword = &v
	return s
}

func (s *ListDmAccountRequest) SetOwnerId(v int64) *ListDmAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *ListDmAccountRequest) SetPageIndex(v int64) *ListDmAccountRequest {
	s.PageIndex = &v
	return s
}

func (s *ListDmAccountRequest) SetPageSize(v int64) *ListDmAccountRequest {
	s.PageSize = &v
	return s
}

func (s *ListDmAccountRequest) SetResourceOwnerAccount(v string) *ListDmAccountRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListDmAccountRequest) SetResourceOwnerId(v int64) *ListDmAccountRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListDmAccountRequest) SetSendType(v string) *ListDmAccountRequest {
	s.SendType = &v
	return s
}

func (s *ListDmAccountRequest) Validate() error {
	return dara.Validate(s)
}
