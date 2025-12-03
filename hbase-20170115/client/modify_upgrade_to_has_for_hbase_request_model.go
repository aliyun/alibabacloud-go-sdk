// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUpgradeToHasForHbaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyUpgradeToHasForHbaseRequest
	GetClientToken() *string
	SetClusterId(v string) *ModifyUpgradeToHasForHbaseRequest
	GetClusterId() *string
	SetHasPassword(v string) *ModifyUpgradeToHasForHbaseRequest
	GetHasPassword() *string
	SetOwnerId(v int64) *ModifyUpgradeToHasForHbaseRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyUpgradeToHasForHbaseRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyUpgradeToHasForHbaseRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyUpgradeToHasForHbaseRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *ModifyUpgradeToHasForHbaseRequest
	GetZoneId() *string
}

type ModifyUpgradeToHasForHbaseRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	HasPassword          *string `json:"HasPassword,omitempty" xml:"HasPassword,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ModifyUpgradeToHasForHbaseRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyUpgradeToHasForHbaseRequest) GoString() string {
	return s.String()
}

func (s *ModifyUpgradeToHasForHbaseRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyUpgradeToHasForHbaseRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyUpgradeToHasForHbaseRequest) GetHasPassword() *string {
	return s.HasPassword
}

func (s *ModifyUpgradeToHasForHbaseRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyUpgradeToHasForHbaseRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyUpgradeToHasForHbaseRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyUpgradeToHasForHbaseRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyUpgradeToHasForHbaseRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *ModifyUpgradeToHasForHbaseRequest) SetClientToken(v string) *ModifyUpgradeToHasForHbaseRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyUpgradeToHasForHbaseRequest) SetClusterId(v string) *ModifyUpgradeToHasForHbaseRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyUpgradeToHasForHbaseRequest) SetHasPassword(v string) *ModifyUpgradeToHasForHbaseRequest {
	s.HasPassword = &v
	return s
}

func (s *ModifyUpgradeToHasForHbaseRequest) SetOwnerId(v int64) *ModifyUpgradeToHasForHbaseRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyUpgradeToHasForHbaseRequest) SetRegionId(v string) *ModifyUpgradeToHasForHbaseRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyUpgradeToHasForHbaseRequest) SetResourceOwnerAccount(v string) *ModifyUpgradeToHasForHbaseRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyUpgradeToHasForHbaseRequest) SetResourceOwnerId(v int64) *ModifyUpgradeToHasForHbaseRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyUpgradeToHasForHbaseRequest) SetZoneId(v string) *ModifyUpgradeToHasForHbaseRequest {
	s.ZoneId = &v
	return s
}

func (s *ModifyUpgradeToHasForHbaseRequest) Validate() error {
	return dara.Validate(s)
}
