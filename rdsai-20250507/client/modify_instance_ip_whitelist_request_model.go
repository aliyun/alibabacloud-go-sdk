// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceIpWhitelistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyInstanceIpWhitelistRequest
	GetClientToken() *string
	SetGroupName(v string) *ModifyInstanceIpWhitelistRequest
	GetGroupName() *string
	SetInstanceName(v string) *ModifyInstanceIpWhitelistRequest
	GetInstanceName() *string
	SetIpWhitelist(v string) *ModifyInstanceIpWhitelistRequest
	GetIpWhitelist() *string
	SetModifyMode(v string) *ModifyInstanceIpWhitelistRequest
	GetModifyMode() *string
	SetRegionId(v string) *ModifyInstanceIpWhitelistRequest
	GetRegionId() *string
}

type ModifyInstanceIpWhitelistRequest struct {
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// default
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// 10.23.XXX.XXX
	IpWhitelist *string `json:"IpWhitelist,omitempty" xml:"IpWhitelist,omitempty"`
	// example:
	//
	// Cover
	ModifyMode *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyInstanceIpWhitelistRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceIpWhitelistRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceIpWhitelistRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyInstanceIpWhitelistRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ModifyInstanceIpWhitelistRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyInstanceIpWhitelistRequest) GetIpWhitelist() *string {
	return s.IpWhitelist
}

func (s *ModifyInstanceIpWhitelistRequest) GetModifyMode() *string {
	return s.ModifyMode
}

func (s *ModifyInstanceIpWhitelistRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInstanceIpWhitelistRequest) SetClientToken(v string) *ModifyInstanceIpWhitelistRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyInstanceIpWhitelistRequest) SetGroupName(v string) *ModifyInstanceIpWhitelistRequest {
	s.GroupName = &v
	return s
}

func (s *ModifyInstanceIpWhitelistRequest) SetInstanceName(v string) *ModifyInstanceIpWhitelistRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyInstanceIpWhitelistRequest) SetIpWhitelist(v string) *ModifyInstanceIpWhitelistRequest {
	s.IpWhitelist = &v
	return s
}

func (s *ModifyInstanceIpWhitelistRequest) SetModifyMode(v string) *ModifyInstanceIpWhitelistRequest {
	s.ModifyMode = &v
	return s
}

func (s *ModifyInstanceIpWhitelistRequest) SetRegionId(v string) *ModifyInstanceIpWhitelistRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceIpWhitelistRequest) Validate() error {
	return dara.Validate(s)
}
