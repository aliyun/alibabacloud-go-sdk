// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterServiceConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ModifyClusterServiceConfigRequest
	GetClusterId() *string
	SetName(v string) *ModifyClusterServiceConfigRequest
	GetName() *string
	SetOwnerId(v int64) *ModifyClusterServiceConfigRequest
	GetOwnerId() *int64
	SetParameters(v string) *ModifyClusterServiceConfigRequest
	GetParameters() *string
	SetRegionId(v string) *ModifyClusterServiceConfigRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyClusterServiceConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyClusterServiceConfigRequest
	GetResourceOwnerId() *int64
	SetRestart(v bool) *ModifyClusterServiceConfigRequest
	GetRestart() *bool
	SetValue(v string) *ModifyClusterServiceConfigRequest
	GetValue() *string
	SetZoneId(v string) *ModifyClusterServiceConfigRequest
	GetZoneId() *string
}

type ModifyClusterServiceConfigRequest struct {
	// This parameter is required.
	ClusterId  *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Restart              *bool   `json:"Restart,omitempty" xml:"Restart,omitempty"`
	Value                *string `json:"Value,omitempty" xml:"Value,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ModifyClusterServiceConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterServiceConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterServiceConfigRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyClusterServiceConfigRequest) GetName() *string {
	return s.Name
}

func (s *ModifyClusterServiceConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyClusterServiceConfigRequest) GetParameters() *string {
	return s.Parameters
}

func (s *ModifyClusterServiceConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyClusterServiceConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyClusterServiceConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyClusterServiceConfigRequest) GetRestart() *bool {
	return s.Restart
}

func (s *ModifyClusterServiceConfigRequest) GetValue() *string {
	return s.Value
}

func (s *ModifyClusterServiceConfigRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *ModifyClusterServiceConfigRequest) SetClusterId(v string) *ModifyClusterServiceConfigRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyClusterServiceConfigRequest) SetName(v string) *ModifyClusterServiceConfigRequest {
	s.Name = &v
	return s
}

func (s *ModifyClusterServiceConfigRequest) SetOwnerId(v int64) *ModifyClusterServiceConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyClusterServiceConfigRequest) SetParameters(v string) *ModifyClusterServiceConfigRequest {
	s.Parameters = &v
	return s
}

func (s *ModifyClusterServiceConfigRequest) SetRegionId(v string) *ModifyClusterServiceConfigRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyClusterServiceConfigRequest) SetResourceOwnerAccount(v string) *ModifyClusterServiceConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyClusterServiceConfigRequest) SetResourceOwnerId(v int64) *ModifyClusterServiceConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyClusterServiceConfigRequest) SetRestart(v bool) *ModifyClusterServiceConfigRequest {
	s.Restart = &v
	return s
}

func (s *ModifyClusterServiceConfigRequest) SetValue(v string) *ModifyClusterServiceConfigRequest {
	s.Value = &v
	return s
}

func (s *ModifyClusterServiceConfigRequest) SetZoneId(v string) *ModifyClusterServiceConfigRequest {
	s.ZoneId = &v
	return s
}

func (s *ModifyClusterServiceConfigRequest) Validate() error {
	return dara.Validate(s)
}
