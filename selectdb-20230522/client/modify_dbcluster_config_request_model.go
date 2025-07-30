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
	// This parameter is required.
	//
	// example:
	//
	// be.conf
	ConfigKey *string `json:"ConfigKey,omitempty" xml:"ConfigKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213c8yvv09-be
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// selectdb-xxxd8a5h60y
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Specifies whether to perform parallel operations on the cluster node.
	//
	// example:
	//
	// false
	ParallelOperation *bool `json:"ParallelOperation,omitempty" xml:"ParallelOperation,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [{\\"name\\":\\"param1\\",\\"value\\":\\"1234577777\\"},{\\"name\\":\\"param2\\",\\"value\\":\\"${yyyyMMdd}\\"}]
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
