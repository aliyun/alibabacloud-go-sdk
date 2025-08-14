// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnroutePrivateZoneInCenToVpcRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessRegionId(v string) *UnroutePrivateZoneInCenToVpcRequest
	GetAccessRegionId() *string
	SetCenId(v string) *UnroutePrivateZoneInCenToVpcRequest
	GetCenId() *string
	SetOwnerAccount(v string) *UnroutePrivateZoneInCenToVpcRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UnroutePrivateZoneInCenToVpcRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *UnroutePrivateZoneInCenToVpcRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UnroutePrivateZoneInCenToVpcRequest
	GetResourceOwnerId() *int64
}

type UnroutePrivateZoneInCenToVpcRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	AccessRegionId *string `json:"AccessRegionId,omitempty" xml:"AccessRegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cen-7qthudw0ll6jmc****
	CenId                *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UnroutePrivateZoneInCenToVpcRequest) String() string {
	return dara.Prettify(s)
}

func (s UnroutePrivateZoneInCenToVpcRequest) GoString() string {
	return s.String()
}

func (s *UnroutePrivateZoneInCenToVpcRequest) GetAccessRegionId() *string {
	return s.AccessRegionId
}

func (s *UnroutePrivateZoneInCenToVpcRequest) GetCenId() *string {
	return s.CenId
}

func (s *UnroutePrivateZoneInCenToVpcRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UnroutePrivateZoneInCenToVpcRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnroutePrivateZoneInCenToVpcRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UnroutePrivateZoneInCenToVpcRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UnroutePrivateZoneInCenToVpcRequest) SetAccessRegionId(v string) *UnroutePrivateZoneInCenToVpcRequest {
	s.AccessRegionId = &v
	return s
}

func (s *UnroutePrivateZoneInCenToVpcRequest) SetCenId(v string) *UnroutePrivateZoneInCenToVpcRequest {
	s.CenId = &v
	return s
}

func (s *UnroutePrivateZoneInCenToVpcRequest) SetOwnerAccount(v string) *UnroutePrivateZoneInCenToVpcRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UnroutePrivateZoneInCenToVpcRequest) SetOwnerId(v int64) *UnroutePrivateZoneInCenToVpcRequest {
	s.OwnerId = &v
	return s
}

func (s *UnroutePrivateZoneInCenToVpcRequest) SetResourceOwnerAccount(v string) *UnroutePrivateZoneInCenToVpcRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnroutePrivateZoneInCenToVpcRequest) SetResourceOwnerId(v int64) *UnroutePrivateZoneInCenToVpcRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnroutePrivateZoneInCenToVpcRequest) Validate() error {
	return dara.Validate(s)
}
