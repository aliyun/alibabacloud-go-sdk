// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomEntitiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *ListCustomEntitiesRequest
	GetAlgorithm() *string
	SetCustomGroupId(v string) *ListCustomEntitiesRequest
	GetCustomGroupId() *string
	SetOwnerAccount(v string) *ListCustomEntitiesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListCustomEntitiesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *ListCustomEntitiesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCustomEntitiesRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *ListCustomEntitiesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListCustomEntitiesRequest
	GetResourceOwnerId() *int64
}

type ListCustomEntitiesRequest struct {
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

func (s ListCustomEntitiesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCustomEntitiesRequest) GoString() string {
	return s.String()
}

func (s *ListCustomEntitiesRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *ListCustomEntitiesRequest) GetCustomGroupId() *string {
	return s.CustomGroupId
}

func (s *ListCustomEntitiesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListCustomEntitiesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListCustomEntitiesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCustomEntitiesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomEntitiesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListCustomEntitiesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListCustomEntitiesRequest) SetAlgorithm(v string) *ListCustomEntitiesRequest {
	s.Algorithm = &v
	return s
}

func (s *ListCustomEntitiesRequest) SetCustomGroupId(v string) *ListCustomEntitiesRequest {
	s.CustomGroupId = &v
	return s
}

func (s *ListCustomEntitiesRequest) SetOwnerAccount(v string) *ListCustomEntitiesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListCustomEntitiesRequest) SetOwnerId(v int64) *ListCustomEntitiesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListCustomEntitiesRequest) SetPageNumber(v int32) *ListCustomEntitiesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCustomEntitiesRequest) SetPageSize(v int32) *ListCustomEntitiesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCustomEntitiesRequest) SetResourceOwnerAccount(v string) *ListCustomEntitiesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListCustomEntitiesRequest) SetResourceOwnerId(v int64) *ListCustomEntitiesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListCustomEntitiesRequest) Validate() error {
	return dara.Validate(s)
}
