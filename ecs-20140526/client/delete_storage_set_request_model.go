// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStorageSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteStorageSetRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *DeleteStorageSetRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteStorageSetRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteStorageSetRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteStorageSetRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteStorageSetRequest
	GetResourceOwnerId() *int64
	SetStorageSetId(v string) *DeleteStorageSetRequest
	GetStorageSetId() *string
}

type DeleteStorageSetRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	StorageSetId *string `json:"StorageSetId,omitempty" xml:"StorageSetId,omitempty"`
}

func (s DeleteStorageSetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteStorageSetRequest) GoString() string {
	return s.String()
}

func (s *DeleteStorageSetRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteStorageSetRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteStorageSetRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteStorageSetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteStorageSetRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteStorageSetRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteStorageSetRequest) GetStorageSetId() *string {
	return s.StorageSetId
}

func (s *DeleteStorageSetRequest) SetClientToken(v string) *DeleteStorageSetRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteStorageSetRequest) SetOwnerAccount(v string) *DeleteStorageSetRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteStorageSetRequest) SetOwnerId(v int64) *DeleteStorageSetRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteStorageSetRequest) SetRegionId(v string) *DeleteStorageSetRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteStorageSetRequest) SetResourceOwnerAccount(v string) *DeleteStorageSetRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteStorageSetRequest) SetResourceOwnerId(v int64) *DeleteStorageSetRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteStorageSetRequest) SetStorageSetId(v string) *DeleteStorageSetRequest {
	s.StorageSetId = &v
	return s
}

func (s *DeleteStorageSetRequest) Validate() error {
	return dara.Validate(s)
}
