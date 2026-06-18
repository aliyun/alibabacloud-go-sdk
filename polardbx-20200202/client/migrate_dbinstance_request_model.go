// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateDBInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *MigrateDBInstanceRequest
	GetDBInstanceName() *string
	SetPrimaryZoneId(v string) *MigrateDBInstanceRequest
	GetPrimaryZoneId() *string
	SetRegionId(v string) *MigrateDBInstanceRequest
	GetRegionId() *string
	SetSecondaryZoneId(v string) *MigrateDBInstanceRequest
	GetSecondaryZoneId() *string
	SetSwitchMode(v string) *MigrateDBInstanceRequest
	GetSwitchMode() *string
	SetTertiaryZoneId(v string) *MigrateDBInstanceRequest
	GetTertiaryZoneId() *string
	SetTopologyType(v string) *MigrateDBInstanceRequest
	GetTopologyType() *string
	SetVpcId(v string) *MigrateDBInstanceRequest
	GetVpcId() *string
	SetVswitchId(v string) *MigrateDBInstanceRequest
	GetVswitchId() *string
}

type MigrateDBInstanceRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-xxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The zone ID of the primary zone for a multi-zone instance. **This parameter is required if you want to create a multi-zone instance.**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-zhangjiakou-b
	PrimaryZoneId *string `json:"PrimaryZoneId,omitempty" xml:"PrimaryZoneId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The secondary zone ID.
	//
	// > This parameter cannot be set to the same value as ZoneId.
	//
	// example:
	//
	// cn-beijing-l
	SecondaryZoneId *string `json:"SecondaryZoneId,omitempty" xml:"SecondaryZoneId,omitempty"`
	// The switchover mode. Valid values:
	//
	// - 0: immediately switches over.
	//
	// - 1: switches over within the O&M window.
	//
	// example:
	//
	// 0
	SwitchMode *string `json:"SwitchMode,omitempty" xml:"SwitchMode,omitempty"`
	// The zone ID for Three-zone deployment.
	//
	// example:
	//
	// cn-beijing-h
	TertiaryZoneId *string `json:"TertiaryZoneId,omitempty" xml:"TertiaryZoneId,omitempty"`
	// The topology type. Valid values:
	//
	// - **3azones**: three-zone deployment.
	//
	// - **1azone**: single-zone deployment.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3azones
	TopologyType *string `json:"TopologyType,omitempty" xml:"TopologyType,omitempty"`
	// The ID of the virtual private cloud (VPC) where the access endpoint resides.
	//
	// example:
	//
	// vpc-****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The vSwitch ID. This parameter is required when you create a DRDS instance of the VPC network type.
	//
	// example:
	//
	// vsw-****
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s MigrateDBInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s MigrateDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *MigrateDBInstanceRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *MigrateDBInstanceRequest) GetPrimaryZoneId() *string {
	return s.PrimaryZoneId
}

func (s *MigrateDBInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *MigrateDBInstanceRequest) GetSecondaryZoneId() *string {
	return s.SecondaryZoneId
}

func (s *MigrateDBInstanceRequest) GetSwitchMode() *string {
	return s.SwitchMode
}

func (s *MigrateDBInstanceRequest) GetTertiaryZoneId() *string {
	return s.TertiaryZoneId
}

func (s *MigrateDBInstanceRequest) GetTopologyType() *string {
	return s.TopologyType
}

func (s *MigrateDBInstanceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *MigrateDBInstanceRequest) GetVswitchId() *string {
	return s.VswitchId
}

func (s *MigrateDBInstanceRequest) SetDBInstanceName(v string) *MigrateDBInstanceRequest {
	s.DBInstanceName = &v
	return s
}

func (s *MigrateDBInstanceRequest) SetPrimaryZoneId(v string) *MigrateDBInstanceRequest {
	s.PrimaryZoneId = &v
	return s
}

func (s *MigrateDBInstanceRequest) SetRegionId(v string) *MigrateDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *MigrateDBInstanceRequest) SetSecondaryZoneId(v string) *MigrateDBInstanceRequest {
	s.SecondaryZoneId = &v
	return s
}

func (s *MigrateDBInstanceRequest) SetSwitchMode(v string) *MigrateDBInstanceRequest {
	s.SwitchMode = &v
	return s
}

func (s *MigrateDBInstanceRequest) SetTertiaryZoneId(v string) *MigrateDBInstanceRequest {
	s.TertiaryZoneId = &v
	return s
}

func (s *MigrateDBInstanceRequest) SetTopologyType(v string) *MigrateDBInstanceRequest {
	s.TopologyType = &v
	return s
}

func (s *MigrateDBInstanceRequest) SetVpcId(v string) *MigrateDBInstanceRequest {
	s.VpcId = &v
	return s
}

func (s *MigrateDBInstanceRequest) SetVswitchId(v string) *MigrateDBInstanceRequest {
	s.VswitchId = &v
	return s
}

func (s *MigrateDBInstanceRequest) Validate() error {
	return dara.Validate(s)
}
