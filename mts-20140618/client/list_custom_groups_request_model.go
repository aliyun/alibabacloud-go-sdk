// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *ListCustomGroupsRequest
	GetAlgorithm() *string
	SetOwnerAccount(v string) *ListCustomGroupsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListCustomGroupsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *ListCustomGroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCustomGroupsRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *ListCustomGroupsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListCustomGroupsRequest
	GetResourceOwnerId() *int64
}

type ListCustomGroupsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// landmark
	Algorithm    *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListCustomGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCustomGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListCustomGroupsRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *ListCustomGroupsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListCustomGroupsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListCustomGroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCustomGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomGroupsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListCustomGroupsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListCustomGroupsRequest) SetAlgorithm(v string) *ListCustomGroupsRequest {
	s.Algorithm = &v
	return s
}

func (s *ListCustomGroupsRequest) SetOwnerAccount(v string) *ListCustomGroupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListCustomGroupsRequest) SetOwnerId(v int64) *ListCustomGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListCustomGroupsRequest) SetPageNumber(v int32) *ListCustomGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCustomGroupsRequest) SetPageSize(v int32) *ListCustomGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCustomGroupsRequest) SetResourceOwnerAccount(v string) *ListCustomGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListCustomGroupsRequest) SetResourceOwnerId(v int64) *ListCustomGroupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListCustomGroupsRequest) Validate() error {
	return dara.Validate(s)
}
