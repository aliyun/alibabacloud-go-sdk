// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRestartClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ModifyRestartClusterRequest
	GetClusterId() *string
	SetComponents(v string) *ModifyRestartClusterRequest
	GetComponents() *string
	SetOwnerId(v int64) *ModifyRestartClusterRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyRestartClusterRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyRestartClusterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyRestartClusterRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *ModifyRestartClusterRequest
	GetZoneId() *string
}

type ModifyRestartClusterRequest struct {
	// This parameter is required.
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Components *string `json:"Components,omitempty" xml:"Components,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ModifyRestartClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyRestartClusterRequest) GoString() string {
	return s.String()
}

func (s *ModifyRestartClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyRestartClusterRequest) GetComponents() *string {
	return s.Components
}

func (s *ModifyRestartClusterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyRestartClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyRestartClusterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyRestartClusterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyRestartClusterRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *ModifyRestartClusterRequest) SetClusterId(v string) *ModifyRestartClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyRestartClusterRequest) SetComponents(v string) *ModifyRestartClusterRequest {
	s.Components = &v
	return s
}

func (s *ModifyRestartClusterRequest) SetOwnerId(v int64) *ModifyRestartClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyRestartClusterRequest) SetRegionId(v string) *ModifyRestartClusterRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyRestartClusterRequest) SetResourceOwnerAccount(v string) *ModifyRestartClusterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyRestartClusterRequest) SetResourceOwnerId(v int64) *ModifyRestartClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyRestartClusterRequest) SetZoneId(v string) *ModifyRestartClusterRequest {
	s.ZoneId = &v
	return s
}

func (s *ModifyRestartClusterRequest) Validate() error {
	return dara.Validate(s)
}
