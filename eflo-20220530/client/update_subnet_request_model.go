// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSubnetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *UpdateSubnetRequest
	GetRegionId() *string
	SetSubnetId(v string) *UpdateSubnetRequest
	GetSubnetId() *string
	SetSubnetName(v string) *UpdateSubnetRequest
	GetSubnetName() *string
	SetVpdId(v string) *UpdateSubnetRequest
	GetVpdId() *string
	SetZoneId(v string) *UpdateSubnetRequest
	GetZoneId() *string
}

type UpdateSubnetRequest struct {
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The subnet instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// subnet-f3zfzmnc
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	// The new name for the subnet instance.
	//
	// example:
	//
	// subnet-1
	SubnetName *string `json:"SubnetName,omitempty" xml:"SubnetName,omitempty"`
	// The ID of the VPD to which the subnet belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpd-aof7dat1
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// The zone ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s UpdateSubnetRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSubnetRequest) GoString() string {
	return s.String()
}

func (s *UpdateSubnetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateSubnetRequest) GetSubnetId() *string {
	return s.SubnetId
}

func (s *UpdateSubnetRequest) GetSubnetName() *string {
	return s.SubnetName
}

func (s *UpdateSubnetRequest) GetVpdId() *string {
	return s.VpdId
}

func (s *UpdateSubnetRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *UpdateSubnetRequest) SetRegionId(v string) *UpdateSubnetRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateSubnetRequest) SetSubnetId(v string) *UpdateSubnetRequest {
	s.SubnetId = &v
	return s
}

func (s *UpdateSubnetRequest) SetSubnetName(v string) *UpdateSubnetRequest {
	s.SubnetName = &v
	return s
}

func (s *UpdateSubnetRequest) SetVpdId(v string) *UpdateSubnetRequest {
	s.VpdId = &v
	return s
}

func (s *UpdateSubnetRequest) SetZoneId(v string) *UpdateSubnetRequest {
	s.ZoneId = &v
	return s
}

func (s *UpdateSubnetRequest) Validate() error {
	return dara.Validate(s)
}
