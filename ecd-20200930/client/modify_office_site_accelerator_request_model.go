// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOfficeSiteAcceleratorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerateRegion(v []*ModifyOfficeSiteAcceleratorRequestAccelerateRegion) *ModifyOfficeSiteAcceleratorRequest
	GetAccelerateRegion() []*ModifyOfficeSiteAcceleratorRequestAccelerateRegion
	SetOfficeSiteId(v string) *ModifyOfficeSiteAcceleratorRequest
	GetOfficeSiteId() *string
	SetRegionId(v string) *ModifyOfficeSiteAcceleratorRequest
	GetRegionId() *string
}

type ModifyOfficeSiteAcceleratorRequest struct {
	AccelerateRegion []*ModifyOfficeSiteAcceleratorRequestAccelerateRegion `json:"AccelerateRegion,omitempty" xml:"AccelerateRegion,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou+dir-363353****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyOfficeSiteAcceleratorRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfficeSiteAcceleratorRequest) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteAcceleratorRequest) GetAccelerateRegion() []*ModifyOfficeSiteAcceleratorRequestAccelerateRegion {
	return s.AccelerateRegion
}

func (s *ModifyOfficeSiteAcceleratorRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *ModifyOfficeSiteAcceleratorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyOfficeSiteAcceleratorRequest) SetAccelerateRegion(v []*ModifyOfficeSiteAcceleratorRequestAccelerateRegion) *ModifyOfficeSiteAcceleratorRequest {
	s.AccelerateRegion = v
	return s
}

func (s *ModifyOfficeSiteAcceleratorRequest) SetOfficeSiteId(v string) *ModifyOfficeSiteAcceleratorRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ModifyOfficeSiteAcceleratorRequest) SetRegionId(v string) *ModifyOfficeSiteAcceleratorRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyOfficeSiteAcceleratorRequest) Validate() error {
	if s.AccelerateRegion != nil {
		for _, item := range s.AccelerateRegion {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyOfficeSiteAcceleratorRequestAccelerateRegion struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	AccelerateRegionId *string `json:"AccelerateRegionId,omitempty" xml:"AccelerateRegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 50
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// BGP
	IspType *string `json:"IspType,omitempty" xml:"IspType,omitempty"`
}

func (s ModifyOfficeSiteAcceleratorRequestAccelerateRegion) String() string {
	return dara.Prettify(s)
}

func (s ModifyOfficeSiteAcceleratorRequestAccelerateRegion) GoString() string {
	return s.String()
}

func (s *ModifyOfficeSiteAcceleratorRequestAccelerateRegion) GetAccelerateRegionId() *string {
	return s.AccelerateRegionId
}

func (s *ModifyOfficeSiteAcceleratorRequestAccelerateRegion) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *ModifyOfficeSiteAcceleratorRequestAccelerateRegion) GetIpVersion() *string {
	return s.IpVersion
}

func (s *ModifyOfficeSiteAcceleratorRequestAccelerateRegion) GetIspType() *string {
	return s.IspType
}

func (s *ModifyOfficeSiteAcceleratorRequestAccelerateRegion) SetAccelerateRegionId(v string) *ModifyOfficeSiteAcceleratorRequestAccelerateRegion {
	s.AccelerateRegionId = &v
	return s
}

func (s *ModifyOfficeSiteAcceleratorRequestAccelerateRegion) SetBandwidth(v int32) *ModifyOfficeSiteAcceleratorRequestAccelerateRegion {
	s.Bandwidth = &v
	return s
}

func (s *ModifyOfficeSiteAcceleratorRequestAccelerateRegion) SetIpVersion(v string) *ModifyOfficeSiteAcceleratorRequestAccelerateRegion {
	s.IpVersion = &v
	return s
}

func (s *ModifyOfficeSiteAcceleratorRequestAccelerateRegion) SetIspType(v string) *ModifyOfficeSiteAcceleratorRequestAccelerateRegion {
	s.IspType = &v
	return s
}

func (s *ModifyOfficeSiteAcceleratorRequestAccelerateRegion) Validate() error {
	return dara.Validate(s)
}
