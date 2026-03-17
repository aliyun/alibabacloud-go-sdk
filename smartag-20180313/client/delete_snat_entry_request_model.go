// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSnatEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteSnatEntryRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DeleteSnatEntryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteSnatEntryRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteSnatEntryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteSnatEntryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteSnatEntryRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DeleteSnatEntryRequest
	GetSmartAGId() *string
}

type DeleteSnatEntryRequest struct {
	// The ID of the SNAT entry.
	//
	// This parameter is required.
	//
	// example:
	//
	// snat-djngdheb*******
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region to which the SAG instance belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-hfbd*******
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s DeleteSnatEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSnatEntryRequest) GoString() string {
	return s.String()
}

func (s *DeleteSnatEntryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteSnatEntryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteSnatEntryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteSnatEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteSnatEntryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteSnatEntryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteSnatEntryRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DeleteSnatEntryRequest) SetInstanceId(v string) *DeleteSnatEntryRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteSnatEntryRequest) SetOwnerAccount(v string) *DeleteSnatEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteSnatEntryRequest) SetOwnerId(v int64) *DeleteSnatEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSnatEntryRequest) SetRegionId(v string) *DeleteSnatEntryRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSnatEntryRequest) SetResourceOwnerAccount(v string) *DeleteSnatEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteSnatEntryRequest) SetResourceOwnerId(v int64) *DeleteSnatEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteSnatEntryRequest) SetSmartAGId(v string) *DeleteSnatEntryRequest {
	s.SmartAGId = &v
	return s
}

func (s *DeleteSnatEntryRequest) Validate() error {
	return dara.Validate(s)
}
