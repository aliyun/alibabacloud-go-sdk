// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceDeploymentModeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyDBInstanceDeploymentModeRequest
	GetDBInstanceId() *string
	SetDeployMode(v string) *ModifyDBInstanceDeploymentModeRequest
	GetDeployMode() *string
	SetStandbyVSwitchId(v string) *ModifyDBInstanceDeploymentModeRequest
	GetStandbyVSwitchId() *string
	SetStandbyZoneId(v string) *ModifyDBInstanceDeploymentModeRequest
	GetStandbyZoneId() *string
}

type ModifyDBInstanceDeploymentModeRequest struct {
	// The cluster ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the IDs of all AnalyticDB for PostgreSQL instances in the specified region.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The deployment mode. Valid values:
	//
	// 	- multiple: Multi-zone development.
	//
	// 	- single: Single-zone deployment.
	//
	// This parameter is required.
	//
	// example:
	//
	// multiple
	DeployMode *string `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	// The vSwitch ID of the secondary zone.
	//
	// >
	//
	// 	- This parameter must be specified only when DeployMode is set to multiple.
	//
	// 	- The vSwitch must be deployed in the zone that is specified by the StandbyZoneId parameter.
	//
	// example:
	//
	// vsw-bp1cpq8mr64paltkb****
	StandbyVSwitchId *string `json:"StandbyVSwitchId,omitempty" xml:"StandbyVSwitchId,omitempty"`
	// The ID of the secondary zone.
	//
	// >
	//
	// 	- This parameter must be specified only when DeployMode is set to multiple.
	//
	// 	- You can call the [DescribeRegions](https://help.aliyun.com/document_detail/86912.html) operation to query the available zone list.
	//
	// 	- The ID of the secondary zone must be different from the ID of the primary zone.
	//
	// example:
	//
	// cn-hangzhou-j
	StandbyZoneId *string `json:"StandbyZoneId,omitempty" xml:"StandbyZoneId,omitempty"`
}

func (s ModifyDBInstanceDeploymentModeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceDeploymentModeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceDeploymentModeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceDeploymentModeRequest) GetDeployMode() *string {
	return s.DeployMode
}

func (s *ModifyDBInstanceDeploymentModeRequest) GetStandbyVSwitchId() *string {
	return s.StandbyVSwitchId
}

func (s *ModifyDBInstanceDeploymentModeRequest) GetStandbyZoneId() *string {
	return s.StandbyZoneId
}

func (s *ModifyDBInstanceDeploymentModeRequest) SetDBInstanceId(v string) *ModifyDBInstanceDeploymentModeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceDeploymentModeRequest) SetDeployMode(v string) *ModifyDBInstanceDeploymentModeRequest {
	s.DeployMode = &v
	return s
}

func (s *ModifyDBInstanceDeploymentModeRequest) SetStandbyVSwitchId(v string) *ModifyDBInstanceDeploymentModeRequest {
	s.StandbyVSwitchId = &v
	return s
}

func (s *ModifyDBInstanceDeploymentModeRequest) SetStandbyZoneId(v string) *ModifyDBInstanceDeploymentModeRequest {
	s.StandbyZoneId = &v
	return s
}

func (s *ModifyDBInstanceDeploymentModeRequest) Validate() error {
	return dara.Validate(s)
}
