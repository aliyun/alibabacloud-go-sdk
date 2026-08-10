// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceConnectionStringRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComputingGroupId(v string) *ModifyDBInstanceConnectionStringRequest
	GetComputingGroupId() *string
	SetConnectionString(v string) *ModifyDBInstanceConnectionStringRequest
	GetConnectionString() *string
	SetConnectionStringPrefix(v string) *ModifyDBInstanceConnectionStringRequest
	GetConnectionStringPrefix() *string
	SetDBInstanceId(v string) *ModifyDBInstanceConnectionStringRequest
	GetDBInstanceId() *string
	SetDBInstanceNetType(v string) *ModifyDBInstanceConnectionStringRequest
	GetDBInstanceNetType() *string
	SetDisablePorts(v string) *ModifyDBInstanceConnectionStringRequest
	GetDisablePorts() *string
	SetRegionId(v string) *ModifyDBInstanceConnectionStringRequest
	GetRegionId() *string
}

type ModifyDBInstanceConnectionStringRequest struct {
	// The compute group ID.
	//
	// example:
	//
	// cc-2ze34****-clickhouse
	ComputingGroupId *string `json:"ComputingGroupId,omitempty" xml:"ComputingGroupId,omitempty"`
	// The endpoint.
	//
	// example:
	//
	// cc-2ze34****-clickhouse..clickhouseserver.pre.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The prefix of the database endpoint.
	//
	// example:
	//
	// cc-****-clickhouse
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-xxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The network type. Valid values:
	//
	// - Vpc: VPC network.
	//
	// - Public: public network.
	//
	// example:
	//
	// Public
	DBInstanceNetType *string `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	// - Disables specified database ports. You can specify multiple ports separated by commas (,).
	//
	// - Only clusters with a kernel version of 24.10.1.11098_1 or later support this parameter.
	//
	// 	Notice: If the cluster was created with a version earlier than 24.10.1.11098_1 and later upgraded to version 24.10.1.11098_1 or later, this parameter is not supported.</notice>
	//
	// example:
	//
	// 9001,8123
	DisablePorts *string `json:"DisablePorts,omitempty" xml:"DisablePorts,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDBInstanceConnectionStringRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceConnectionStringRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConnectionStringRequest) GetComputingGroupId() *string {
	return s.ComputingGroupId
}

func (s *ModifyDBInstanceConnectionStringRequest) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *ModifyDBInstanceConnectionStringRequest) GetConnectionStringPrefix() *string {
	return s.ConnectionStringPrefix
}

func (s *ModifyDBInstanceConnectionStringRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceConnectionStringRequest) GetDBInstanceNetType() *string {
	return s.DBInstanceNetType
}

func (s *ModifyDBInstanceConnectionStringRequest) GetDisablePorts() *string {
	return s.DisablePorts
}

func (s *ModifyDBInstanceConnectionStringRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDBInstanceConnectionStringRequest) SetComputingGroupId(v string) *ModifyDBInstanceConnectionStringRequest {
	s.ComputingGroupId = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetConnectionString(v string) *ModifyDBInstanceConnectionStringRequest {
	s.ConnectionString = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetConnectionStringPrefix(v string) *ModifyDBInstanceConnectionStringRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetDBInstanceId(v string) *ModifyDBInstanceConnectionStringRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetDBInstanceNetType(v string) *ModifyDBInstanceConnectionStringRequest {
	s.DBInstanceNetType = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetDisablePorts(v string) *ModifyDBInstanceConnectionStringRequest {
	s.DisablePorts = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetRegionId(v string) *ModifyDBInstanceConnectionStringRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) Validate() error {
	return dara.Validate(s)
}
