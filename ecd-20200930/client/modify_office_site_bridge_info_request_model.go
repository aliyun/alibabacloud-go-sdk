// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOfficeSiteBridgeInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBridgeId(v string) *ModifyOfficeSiteBridgeInfoRequest
	GetBridgeId() *string
	SetBridgeLevel(v string) *ModifyOfficeSiteBridgeInfoRequest
	GetBridgeLevel() *string
	SetBridgeType(v string) *ModifyOfficeSiteBridgeInfoRequest
	GetBridgeType() *string
	SetEnableBridge(v bool) *ModifyOfficeSiteBridgeInfoRequest
	GetEnableBridge() *bool
	SetLicense(v string) *ModifyOfficeSiteBridgeInfoRequest
	GetLicense() *string
	SetOfficeSiteId(v string) *ModifyOfficeSiteBridgeInfoRequest
	GetOfficeSiteId() *string
	SetRegionId(v string) *ModifyOfficeSiteBridgeInfoRequest
	GetRegionId() *string
}

type ModifyOfficeSiteBridgeInfoRequest struct {
	// The virtual bridge ID.
	//
	// example:
	//
	// vb-fsifhaiushfishf***
	BridgeId *string `json:"BridgeId,omitempty" xml:"BridgeId,omitempty"`
	// The virtual bridge specifications.
	//
	// example:
	//
	// vb.pro
	BridgeLevel *string `json:"BridgeLevel,omitempty" xml:"BridgeLevel,omitempty"`
	// The third-party plugin type of the virtual bridge.
	//
	// example:
	//
	// unsr
	BridgeType *string `json:"BridgeType,omitempty" xml:"BridgeType,omitempty"`
	// Specifies whether to enable the bridge.
	//
	// example:
	//
	// true
	EnableBridge *bool `json:"EnableBridge,omitempty" xml:"EnableBridge,omitempty"`
	// The activation code object.
	//
	// example:
	//
	// ab5b76f4c0bf4a5abd06ea23991a47afa1a4bb4acb4e4204882b40795f946e74
	License *string `json:"License,omitempty" xml:"License,omitempty"`
	// The office network ID.
	//
	// example:
	//
	// cn-hangzhou+dir-387822****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The region ID. You can call [DescribeRegions](~~DescribeRegions~~) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyOfficeSiteBridgeInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfficeSiteBridgeInfoRequest) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteBridgeInfoRequest) GetBridgeId() *string {
	return s.BridgeId
}

func (s *ModifyOfficeSiteBridgeInfoRequest) GetBridgeLevel() *string {
	return s.BridgeLevel
}

func (s *ModifyOfficeSiteBridgeInfoRequest) GetBridgeType() *string {
	return s.BridgeType
}

func (s *ModifyOfficeSiteBridgeInfoRequest) GetEnableBridge() *bool {
	return s.EnableBridge
}

func (s *ModifyOfficeSiteBridgeInfoRequest) GetLicense() *string {
	return s.License
}

func (s *ModifyOfficeSiteBridgeInfoRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *ModifyOfficeSiteBridgeInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyOfficeSiteBridgeInfoRequest) SetBridgeId(v string) *ModifyOfficeSiteBridgeInfoRequest {
	s.BridgeId = &v
	return s
}

func (s *ModifyOfficeSiteBridgeInfoRequest) SetBridgeLevel(v string) *ModifyOfficeSiteBridgeInfoRequest {
	s.BridgeLevel = &v
	return s
}

func (s *ModifyOfficeSiteBridgeInfoRequest) SetBridgeType(v string) *ModifyOfficeSiteBridgeInfoRequest {
	s.BridgeType = &v
	return s
}

func (s *ModifyOfficeSiteBridgeInfoRequest) SetEnableBridge(v bool) *ModifyOfficeSiteBridgeInfoRequest {
	s.EnableBridge = &v
	return s
}

func (s *ModifyOfficeSiteBridgeInfoRequest) SetLicense(v string) *ModifyOfficeSiteBridgeInfoRequest {
	s.License = &v
	return s
}

func (s *ModifyOfficeSiteBridgeInfoRequest) SetOfficeSiteId(v string) *ModifyOfficeSiteBridgeInfoRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ModifyOfficeSiteBridgeInfoRequest) SetRegionId(v string) *ModifyOfficeSiteBridgeInfoRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyOfficeSiteBridgeInfoRequest) Validate() error {
	return dara.Validate(s)
}
