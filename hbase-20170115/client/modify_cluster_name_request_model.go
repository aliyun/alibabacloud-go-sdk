// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ModifyClusterNameRequest
	GetClusterId() *string
	SetClusterName(v string) *ModifyClusterNameRequest
	GetClusterName() *string
	SetOwnerId(v int64) *ModifyClusterNameRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyClusterNameRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyClusterNameRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyClusterNameRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *ModifyClusterNameRequest
	GetZoneId() *string
}

type ModifyClusterNameRequest struct {
	// This parameter is required.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ModifyClusterNameRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterNameRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterNameRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyClusterNameRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *ModifyClusterNameRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyClusterNameRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyClusterNameRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyClusterNameRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyClusterNameRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *ModifyClusterNameRequest) SetClusterId(v string) *ModifyClusterNameRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyClusterNameRequest) SetClusterName(v string) *ModifyClusterNameRequest {
	s.ClusterName = &v
	return s
}

func (s *ModifyClusterNameRequest) SetOwnerId(v int64) *ModifyClusterNameRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyClusterNameRequest) SetRegionId(v string) *ModifyClusterNameRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyClusterNameRequest) SetResourceOwnerAccount(v string) *ModifyClusterNameRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyClusterNameRequest) SetResourceOwnerId(v int64) *ModifyClusterNameRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyClusterNameRequest) SetZoneId(v string) *ModifyClusterNameRequest {
	s.ZoneId = &v
	return s
}

func (s *ModifyClusterNameRequest) Validate() error {
	return dara.Validate(s)
}
