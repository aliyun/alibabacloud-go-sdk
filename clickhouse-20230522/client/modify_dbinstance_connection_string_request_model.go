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
	ComputingGroupId *string `json:"ComputingGroupId,omitempty" xml:"ComputingGroupId,omitempty"`
	// The endpoint of the cluster.
	//
	// example:
	//
	// cc-2ze34****-clickhouse..clickhouseserver.pre.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The prefix of the endpoint that is used to connect to the database.
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
	DBInstanceId      *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceNetType *string `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	// 	- The database ports that you want to disable. Separate multiple ports with commas (,).
	//
	// 	- This parameter is supported only for clusters whose minor engine version is 24.10.1.11098_1 or later.
	//
	//     **
	//
	//     **Note*	- If you create a cluster whose minor engine version is earlier than 24.10.1.11098_1 and you update the minor engine version to 24.10.1.11098_1 or later, the cluster still does not support this parameter.
	//
	// example:
	//
	// 9001,8123
	DisablePorts *string `json:"DisablePorts,omitempty" xml:"DisablePorts,omitempty"`
	// The region ID.
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
