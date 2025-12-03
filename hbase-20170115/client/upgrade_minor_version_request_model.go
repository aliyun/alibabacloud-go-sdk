// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeMinorVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpgradeMinorVersionRequest
	GetClientToken() *string
	SetClusterId(v string) *UpgradeMinorVersionRequest
	GetClusterId() *string
	SetComponents(v string) *UpgradeMinorVersionRequest
	GetComponents() *string
	SetOwnerId(v int64) *UpgradeMinorVersionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpgradeMinorVersionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UpgradeMinorVersionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpgradeMinorVersionRequest
	GetResourceOwnerId() *int64
	SetUpgradeVersion(v string) *UpgradeMinorVersionRequest
	GetUpgradeVersion() *string
	SetZoneId(v string) *UpgradeMinorVersionRequest
	GetZoneId() *string
}

type UpgradeMinorVersionRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Components           *string `json:"Components,omitempty" xml:"Components,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UpgradeVersion       *string `json:"UpgradeVersion,omitempty" xml:"UpgradeVersion,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s UpgradeMinorVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeMinorVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeMinorVersionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpgradeMinorVersionRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpgradeMinorVersionRequest) GetComponents() *string {
	return s.Components
}

func (s *UpgradeMinorVersionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpgradeMinorVersionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpgradeMinorVersionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpgradeMinorVersionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpgradeMinorVersionRequest) GetUpgradeVersion() *string {
	return s.UpgradeVersion
}

func (s *UpgradeMinorVersionRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *UpgradeMinorVersionRequest) SetClientToken(v string) *UpgradeMinorVersionRequest {
	s.ClientToken = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetClusterId(v string) *UpgradeMinorVersionRequest {
	s.ClusterId = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetComponents(v string) *UpgradeMinorVersionRequest {
	s.Components = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetOwnerId(v int64) *UpgradeMinorVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetRegionId(v string) *UpgradeMinorVersionRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetResourceOwnerAccount(v string) *UpgradeMinorVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetResourceOwnerId(v int64) *UpgradeMinorVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetUpgradeVersion(v string) *UpgradeMinorVersionRequest {
	s.UpgradeVersion = &v
	return s
}

func (s *UpgradeMinorVersionRequest) SetZoneId(v string) *UpgradeMinorVersionRequest {
	s.ZoneId = &v
	return s
}

func (s *UpgradeMinorVersionRequest) Validate() error {
	return dara.Validate(s)
}
