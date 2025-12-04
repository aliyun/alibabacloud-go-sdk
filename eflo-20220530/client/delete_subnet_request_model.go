// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSubnetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeleteSubnetRequest
	GetRegionId() *string
	SetSubnetId(v string) *DeleteSubnetRequest
	GetSubnetId() *string
	SetVpdId(v string) *DeleteSubnetRequest
	GetVpdId() *string
	SetZoneId(v string) *DeleteSubnetRequest
	GetZoneId() *string
}

type DeleteSubnetRequest struct {
	// The region ID of the disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Lingjun subnet ID
	//
	// This parameter is required.
	//
	// example:
	//
	// subnet-f3zfzmnc
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	// Lingjun CIDR block ID
	//
	// This parameter is required.
	//
	// example:
	//
	// vpd-iv2zm1qf
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// Zone
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DeleteSubnetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSubnetRequest) GoString() string {
	return s.String()
}

func (s *DeleteSubnetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteSubnetRequest) GetSubnetId() *string {
	return s.SubnetId
}

func (s *DeleteSubnetRequest) GetVpdId() *string {
	return s.VpdId
}

func (s *DeleteSubnetRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DeleteSubnetRequest) SetRegionId(v string) *DeleteSubnetRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSubnetRequest) SetSubnetId(v string) *DeleteSubnetRequest {
	s.SubnetId = &v
	return s
}

func (s *DeleteSubnetRequest) SetVpdId(v string) *DeleteSubnetRequest {
	s.VpdId = &v
	return s
}

func (s *DeleteSubnetRequest) SetZoneId(v string) *DeleteSubnetRequest {
	s.ZoneId = &v
	return s
}

func (s *DeleteSubnetRequest) Validate() error {
	return dara.Validate(s)
}
