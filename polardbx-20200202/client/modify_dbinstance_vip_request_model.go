// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceVipRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *ModifyDBInstanceVipRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *ModifyDBInstanceVipRequest
	GetRegionId() *string
	SetVSwitchId(v string) *ModifyDBInstanceVipRequest
	GetVSwitchId() *string
	SetVpcId(v string) *ModifyDBInstanceVipRequest
	GetVpcId() *string
}

type ModifyDBInstanceVipRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vsw-*********
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vpc-********
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ModifyDBInstanceVipRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceVipRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceVipRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifyDBInstanceVipRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDBInstanceVipRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ModifyDBInstanceVipRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ModifyDBInstanceVipRequest) SetDBInstanceName(v string) *ModifyDBInstanceVipRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyDBInstanceVipRequest) SetRegionId(v string) *ModifyDBInstanceVipRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBInstanceVipRequest) SetVSwitchId(v string) *ModifyDBInstanceVipRequest {
	s.VSwitchId = &v
	return s
}

func (s *ModifyDBInstanceVipRequest) SetVpcId(v string) *ModifyDBInstanceVipRequest {
	s.VpcId = &v
	return s
}

func (s *ModifyDBInstanceVipRequest) Validate() error {
	return dara.Validate(s)
}
