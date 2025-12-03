// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHasRootPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ModifyHasRootPasswordRequest
	GetClusterId() *string
	SetHasPassword(v string) *ModifyHasRootPasswordRequest
	GetHasPassword() *string
	SetOwnerId(v int64) *ModifyHasRootPasswordRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyHasRootPasswordRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyHasRootPasswordRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyHasRootPasswordRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *ModifyHasRootPasswordRequest
	GetZoneId() *string
}

type ModifyHasRootPasswordRequest struct {
	// This parameter is required.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	HasPassword *string `json:"HasPassword,omitempty" xml:"HasPassword,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ModifyHasRootPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyHasRootPasswordRequest) GoString() string {
	return s.String()
}

func (s *ModifyHasRootPasswordRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyHasRootPasswordRequest) GetHasPassword() *string {
	return s.HasPassword
}

func (s *ModifyHasRootPasswordRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyHasRootPasswordRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyHasRootPasswordRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyHasRootPasswordRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyHasRootPasswordRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *ModifyHasRootPasswordRequest) SetClusterId(v string) *ModifyHasRootPasswordRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyHasRootPasswordRequest) SetHasPassword(v string) *ModifyHasRootPasswordRequest {
	s.HasPassword = &v
	return s
}

func (s *ModifyHasRootPasswordRequest) SetOwnerId(v int64) *ModifyHasRootPasswordRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyHasRootPasswordRequest) SetRegionId(v string) *ModifyHasRootPasswordRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyHasRootPasswordRequest) SetResourceOwnerAccount(v string) *ModifyHasRootPasswordRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyHasRootPasswordRequest) SetResourceOwnerId(v int64) *ModifyHasRootPasswordRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyHasRootPasswordRequest) SetZoneId(v string) *ModifyHasRootPasswordRequest {
	s.ZoneId = &v
	return s
}

func (s *ModifyHasRootPasswordRequest) Validate() error {
	return dara.Validate(s)
}
