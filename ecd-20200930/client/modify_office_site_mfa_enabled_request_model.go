// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOfficeSiteMfaEnabledRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMfaEnabled(v bool) *ModifyOfficeSiteMfaEnabledRequest
	GetMfaEnabled() *bool
	SetOfficeSiteId(v string) *ModifyOfficeSiteMfaEnabledRequest
	GetOfficeSiteId() *string
	SetRegionId(v string) *ModifyOfficeSiteMfaEnabledRequest
	GetRegionId() *string
}

type ModifyOfficeSiteMfaEnabledRequest struct {
	// Specifies whether to enable MFA.
	//
	// Valid values:
	//
	// 	- true
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- false
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	MfaEnabled *bool `json:"MfaEnabled,omitempty" xml:"MfaEnabled,omitempty"`
	// The office network ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+dir-363353****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyOfficeSiteMfaEnabledRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfficeSiteMfaEnabledRequest) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteMfaEnabledRequest) GetMfaEnabled() *bool {
	return s.MfaEnabled
}

func (s *ModifyOfficeSiteMfaEnabledRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *ModifyOfficeSiteMfaEnabledRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyOfficeSiteMfaEnabledRequest) SetMfaEnabled(v bool) *ModifyOfficeSiteMfaEnabledRequest {
	s.MfaEnabled = &v
	return s
}

func (s *ModifyOfficeSiteMfaEnabledRequest) SetOfficeSiteId(v string) *ModifyOfficeSiteMfaEnabledRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ModifyOfficeSiteMfaEnabledRequest) SetRegionId(v string) *ModifyOfficeSiteMfaEnabledRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyOfficeSiteMfaEnabledRequest) Validate() error {
	return dara.Validate(s)
}
