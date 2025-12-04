// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchVccConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *SwitchVccConnectionRequest
	GetCenId() *string
	SetConnectionType(v string) *SwitchVccConnectionRequest
	GetConnectionType() *string
	SetRegionId(v string) *SwitchVccConnectionRequest
	GetRegionId() *string
	SetVSwitchId(v string) *SwitchVccConnectionRequest
	GetVSwitchId() *string
	SetVccId(v string) *SwitchVccConnectionRequest
	GetVccId() *string
	SetVpcId(v string) *SwitchVccConnectionRequest
	GetVpcId() *string
}

type SwitchVccConnectionRequest struct {
	// CEN
	//
	// example:
	//
	// cen-bkiw0x1347roek****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// Connection type, CENTR/VPC
	//
	// example:
	//
	// CENTR
	ConnectionType *string `json:"ConnectionType,omitempty" xml:"ConnectionType,omitempty"`
	// Region ID
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// vSwitch ID
	//
	// example:
	//
	// vsw-t4nahb0pxckgktxfv****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// Cloud Connect Network (CCN) ID
	//
	// This parameter is required.
	//
	// example:
	//
	// vcc-cn-zvp2w22****
	VccId *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
	// VPC ID
	//
	// example:
	//
	// vpc-uf6aa4ddo97frj22t****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s SwitchVccConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchVccConnectionRequest) GoString() string {
	return s.String()
}

func (s *SwitchVccConnectionRequest) GetCenId() *string {
	return s.CenId
}

func (s *SwitchVccConnectionRequest) GetConnectionType() *string {
	return s.ConnectionType
}

func (s *SwitchVccConnectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SwitchVccConnectionRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *SwitchVccConnectionRequest) GetVccId() *string {
	return s.VccId
}

func (s *SwitchVccConnectionRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *SwitchVccConnectionRequest) SetCenId(v string) *SwitchVccConnectionRequest {
	s.CenId = &v
	return s
}

func (s *SwitchVccConnectionRequest) SetConnectionType(v string) *SwitchVccConnectionRequest {
	s.ConnectionType = &v
	return s
}

func (s *SwitchVccConnectionRequest) SetRegionId(v string) *SwitchVccConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *SwitchVccConnectionRequest) SetVSwitchId(v string) *SwitchVccConnectionRequest {
	s.VSwitchId = &v
	return s
}

func (s *SwitchVccConnectionRequest) SetVccId(v string) *SwitchVccConnectionRequest {
	s.VccId = &v
	return s
}

func (s *SwitchVccConnectionRequest) SetVpcId(v string) *SwitchVccConnectionRequest {
	s.VpcId = &v
	return s
}

func (s *SwitchVccConnectionRequest) Validate() error {
	return dara.Validate(s)
}
