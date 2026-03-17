// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRoamClientUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOriginRegionId(v string) *RoamClientUserRequest
	GetOriginRegionId() *string
	SetOriginSmartAGId(v string) *RoamClientUserRequest
	GetOriginSmartAGId() *string
	SetOwnerAccount(v string) *RoamClientUserRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RoamClientUserRequest
	GetOwnerId() *int64
	SetRegionId(v string) *RoamClientUserRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *RoamClientUserRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RoamClientUserRequest
	GetResourceOwnerId() *int64
	SetTargetSmartAGId(v string) *RoamClientUserRequest
	GetTargetSmartAGId() *string
	SetUserName(v string) *RoamClientUserRequest
	GetUserName() *string
}

type RoamClientUserRequest struct {
	// The region ID of the SAG app instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ap-southeast-1
	OriginRegionId *string `json:"OriginRegionId,omitempty" xml:"OriginRegionId,omitempty"`
	// The ID of the source SAG app instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-m9uhqekwnqcnyy****
	OriginSmartAGId *string `json:"OriginSmartAGId,omitempty" xml:"OriginSmartAGId,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the SAG app instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the destination SAG app instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-ghwa10ko6ndwug****
	TargetSmartAGId *string `json:"TargetSmartAGId,omitempty" xml:"TargetSmartAGId,omitempty"`
	// The usernames of the client for which you want to enable roaming.
	//
	// This parameter is required.
	//
	// example:
	//
	// nametest
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s RoamClientUserRequest) String() string {
	return dara.Prettify(s)
}

func (s RoamClientUserRequest) GoString() string {
	return s.String()
}

func (s *RoamClientUserRequest) GetOriginRegionId() *string {
	return s.OriginRegionId
}

func (s *RoamClientUserRequest) GetOriginSmartAGId() *string {
	return s.OriginSmartAGId
}

func (s *RoamClientUserRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RoamClientUserRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RoamClientUserRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RoamClientUserRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RoamClientUserRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RoamClientUserRequest) GetTargetSmartAGId() *string {
	return s.TargetSmartAGId
}

func (s *RoamClientUserRequest) GetUserName() *string {
	return s.UserName
}

func (s *RoamClientUserRequest) SetOriginRegionId(v string) *RoamClientUserRequest {
	s.OriginRegionId = &v
	return s
}

func (s *RoamClientUserRequest) SetOriginSmartAGId(v string) *RoamClientUserRequest {
	s.OriginSmartAGId = &v
	return s
}

func (s *RoamClientUserRequest) SetOwnerAccount(v string) *RoamClientUserRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RoamClientUserRequest) SetOwnerId(v int64) *RoamClientUserRequest {
	s.OwnerId = &v
	return s
}

func (s *RoamClientUserRequest) SetRegionId(v string) *RoamClientUserRequest {
	s.RegionId = &v
	return s
}

func (s *RoamClientUserRequest) SetResourceOwnerAccount(v string) *RoamClientUserRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RoamClientUserRequest) SetResourceOwnerId(v int64) *RoamClientUserRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RoamClientUserRequest) SetTargetSmartAGId(v string) *RoamClientUserRequest {
	s.TargetSmartAGId = &v
	return s
}

func (s *RoamClientUserRequest) SetUserName(v string) *RoamClientUserRequest {
	s.UserName = &v
	return s
}

func (s *RoamClientUserRequest) Validate() error {
	return dara.Validate(s)
}
