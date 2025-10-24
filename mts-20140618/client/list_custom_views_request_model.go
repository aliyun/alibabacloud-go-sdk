// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomViewsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *ListCustomViewsRequest
	GetAlgorithm() *string
	SetCustomEntityId(v string) *ListCustomViewsRequest
	GetCustomEntityId() *string
	SetCustomGroupId(v string) *ListCustomViewsRequest
	GetCustomGroupId() *string
	SetOwnerAccount(v string) *ListCustomViewsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListCustomViewsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *ListCustomViewsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCustomViewsRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *ListCustomViewsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListCustomViewsRequest
	GetResourceOwnerId() *int64
}

type ListCustomViewsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// landmark
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CustomEntityId *string `json:"CustomEntityId,omitempty" xml:"CustomEntityId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CustomGroupId *string `json:"CustomGroupId,omitempty" xml:"CustomGroupId,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s ListCustomViewsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCustomViewsRequest) GoString() string {
	return s.String()
}

func (s *ListCustomViewsRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *ListCustomViewsRequest) GetCustomEntityId() *string {
	return s.CustomEntityId
}

func (s *ListCustomViewsRequest) GetCustomGroupId() *string {
	return s.CustomGroupId
}

func (s *ListCustomViewsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListCustomViewsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListCustomViewsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCustomViewsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomViewsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListCustomViewsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListCustomViewsRequest) SetAlgorithm(v string) *ListCustomViewsRequest {
	s.Algorithm = &v
	return s
}

func (s *ListCustomViewsRequest) SetCustomEntityId(v string) *ListCustomViewsRequest {
	s.CustomEntityId = &v
	return s
}

func (s *ListCustomViewsRequest) SetCustomGroupId(v string) *ListCustomViewsRequest {
	s.CustomGroupId = &v
	return s
}

func (s *ListCustomViewsRequest) SetOwnerAccount(v string) *ListCustomViewsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListCustomViewsRequest) SetOwnerId(v int64) *ListCustomViewsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListCustomViewsRequest) SetPageNumber(v int32) *ListCustomViewsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCustomViewsRequest) SetPageSize(v int32) *ListCustomViewsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCustomViewsRequest) SetResourceOwnerAccount(v string) *ListCustomViewsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListCustomViewsRequest) SetResourceOwnerId(v int64) *ListCustomViewsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListCustomViewsRequest) Validate() error {
	return dara.Validate(s)
}
