// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterSecurityIpListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ModifyClusterSecurityIpListRequest
	GetClusterId() *string
	SetOwnerId(v int64) *ModifyClusterSecurityIpListRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyClusterSecurityIpListRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyClusterSecurityIpListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyClusterSecurityIpListRequest
	GetResourceOwnerId() *int64
	SetSecurityIpList(v string) *ModifyClusterSecurityIpListRequest
	GetSecurityIpList() *string
	SetZoneId(v string) *ModifyClusterSecurityIpListRequest
	GetZoneId() *string
}

type ModifyClusterSecurityIpListRequest struct {
	// This parameter is required.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	SecurityIpList *string `json:"SecurityIpList,omitempty" xml:"SecurityIpList,omitempty"`
	ZoneId         *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ModifyClusterSecurityIpListRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterSecurityIpListRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterSecurityIpListRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyClusterSecurityIpListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyClusterSecurityIpListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyClusterSecurityIpListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyClusterSecurityIpListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyClusterSecurityIpListRequest) GetSecurityIpList() *string {
	return s.SecurityIpList
}

func (s *ModifyClusterSecurityIpListRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *ModifyClusterSecurityIpListRequest) SetClusterId(v string) *ModifyClusterSecurityIpListRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyClusterSecurityIpListRequest) SetOwnerId(v int64) *ModifyClusterSecurityIpListRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyClusterSecurityIpListRequest) SetRegionId(v string) *ModifyClusterSecurityIpListRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyClusterSecurityIpListRequest) SetResourceOwnerAccount(v string) *ModifyClusterSecurityIpListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyClusterSecurityIpListRequest) SetResourceOwnerId(v int64) *ModifyClusterSecurityIpListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyClusterSecurityIpListRequest) SetSecurityIpList(v string) *ModifyClusterSecurityIpListRequest {
	s.SecurityIpList = &v
	return s
}

func (s *ModifyClusterSecurityIpListRequest) SetZoneId(v string) *ModifyClusterSecurityIpListRequest {
	s.ZoneId = &v
	return s
}

func (s *ModifyClusterSecurityIpListRequest) Validate() error {
	return dara.Validate(s)
}
