// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigKey(v string) *ModifyDBClusterConfigRequest
	GetConfigKey() *string
	SetDBClusterId(v string) *ModifyDBClusterConfigRequest
	GetDBClusterId() *string
	SetDBInstanceId(v string) *ModifyDBClusterConfigRequest
	GetDBInstanceId() *string
	SetParallelOperation(v bool) *ModifyDBClusterConfigRequest
	GetParallelOperation() *bool
	SetParameters(v string) *ModifyDBClusterConfigRequest
	GetParameters() *string
	SetRegionId(v string) *ModifyDBClusterConfigRequest
	GetRegionId() *string
	SetSwitchTimeMode(v string) *ModifyDBClusterConfigRequest
	GetSwitchTimeMode() *string
}

type ModifyDBClusterConfigRequest struct {
	// The configuration file to modify. Set this parameter to be.conf for compute clusters or fe.conf for FE clusters.
	//
	// This parameter is required.
	//
	// example:
	//
	// be.conf
	ConfigKey *string `json:"ConfigKey,omitempty" xml:"ConfigKey,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213c8*****-be
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213c8*****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Specifies whether to perform operations on cluster nodes in parallel.
	//
	// example:
	//
	// false
	ParallelOperation *bool `json:"ParallelOperation,omitempty" xml:"ParallelOperation,omitempty"`
	// The JSON string of parameters and parameter values.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"param1_name":"param1_value","param2_name":"param2_value"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The upgrade mode. If this parameter is not specified, the upgrade takes effect immediately. Set this parameter to 1 to perform the upgrade during the maintenance window.
	//
	// example:
	//
	// 1
	SwitchTimeMode *string `json:"SwitchTimeMode,omitempty" xml:"SwitchTimeMode,omitempty"`
}

func (s ModifyDBClusterConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterConfigRequest) GetConfigKey() *string {
	return s.ConfigKey
}

func (s *ModifyDBClusterConfigRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterConfigRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBClusterConfigRequest) GetParallelOperation() *bool {
	return s.ParallelOperation
}

func (s *ModifyDBClusterConfigRequest) GetParameters() *string {
	return s.Parameters
}

func (s *ModifyDBClusterConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDBClusterConfigRequest) GetSwitchTimeMode() *string {
	return s.SwitchTimeMode
}

func (s *ModifyDBClusterConfigRequest) SetConfigKey(v string) *ModifyDBClusterConfigRequest {
	s.ConfigKey = &v
	return s
}

func (s *ModifyDBClusterConfigRequest) SetDBClusterId(v string) *ModifyDBClusterConfigRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterConfigRequest) SetDBInstanceId(v string) *ModifyDBClusterConfigRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBClusterConfigRequest) SetParallelOperation(v bool) *ModifyDBClusterConfigRequest {
	s.ParallelOperation = &v
	return s
}

func (s *ModifyDBClusterConfigRequest) SetParameters(v string) *ModifyDBClusterConfigRequest {
	s.Parameters = &v
	return s
}

func (s *ModifyDBClusterConfigRequest) SetRegionId(v string) *ModifyDBClusterConfigRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBClusterConfigRequest) SetSwitchTimeMode(v string) *ModifyDBClusterConfigRequest {
	s.SwitchTimeMode = &v
	return s
}

func (s *ModifyDBClusterConfigRequest) Validate() error {
	return dara.Validate(s)
}
