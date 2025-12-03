// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterNetTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ModifyClusterNetTypeRequest
	GetClusterId() *string
	SetNetType(v string) *ModifyClusterNetTypeRequest
	GetNetType() *string
	SetOwnerId(v int64) *ModifyClusterNetTypeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyClusterNetTypeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyClusterNetTypeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyClusterNetTypeRequest
	GetResourceOwnerId() *int64
	SetVSwitchId(v string) *ModifyClusterNetTypeRequest
	GetVSwitchId() *string
	SetVpcId(v string) *ModifyClusterNetTypeRequest
	GetVpcId() *string
	SetZoneId(v string) *ModifyClusterNetTypeRequest
	GetZoneId() *string
}

type ModifyClusterNetTypeRequest struct {
	// This parameter is required.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	VSwitchId            *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId                *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ModifyClusterNetTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterNetTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterNetTypeRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyClusterNetTypeRequest) GetNetType() *string {
	return s.NetType
}

func (s *ModifyClusterNetTypeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyClusterNetTypeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyClusterNetTypeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyClusterNetTypeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyClusterNetTypeRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ModifyClusterNetTypeRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ModifyClusterNetTypeRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *ModifyClusterNetTypeRequest) SetClusterId(v string) *ModifyClusterNetTypeRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyClusterNetTypeRequest) SetNetType(v string) *ModifyClusterNetTypeRequest {
	s.NetType = &v
	return s
}

func (s *ModifyClusterNetTypeRequest) SetOwnerId(v int64) *ModifyClusterNetTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyClusterNetTypeRequest) SetRegionId(v string) *ModifyClusterNetTypeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyClusterNetTypeRequest) SetResourceOwnerAccount(v string) *ModifyClusterNetTypeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyClusterNetTypeRequest) SetResourceOwnerId(v int64) *ModifyClusterNetTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyClusterNetTypeRequest) SetVSwitchId(v string) *ModifyClusterNetTypeRequest {
	s.VSwitchId = &v
	return s
}

func (s *ModifyClusterNetTypeRequest) SetVpcId(v string) *ModifyClusterNetTypeRequest {
	s.VpcId = &v
	return s
}

func (s *ModifyClusterNetTypeRequest) SetZoneId(v string) *ModifyClusterNetTypeRequest {
	s.ZoneId = &v
	return s
}

func (s *ModifyClusterNetTypeRequest) Validate() error {
	return dara.Validate(s)
}
