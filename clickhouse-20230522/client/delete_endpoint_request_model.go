// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComputingGroupId(v string) *DeleteEndpointRequest
	GetComputingGroupId() *string
	SetConnectionString(v string) *DeleteEndpointRequest
	GetConnectionString() *string
	SetDBInstanceId(v string) *DeleteEndpointRequest
	GetDBInstanceId() *string
	SetDBInstanceNetType(v string) *DeleteEndpointRequest
	GetDBInstanceNetType() *string
	SetRegionId(v string) *DeleteEndpointRequest
	GetRegionId() *string
}

type DeleteEndpointRequest struct {
	ComputingGroupId *string `json:"ComputingGroupId,omitempty" xml:"ComputingGroupId,omitempty"`
	// The prefix of the endpoint, which indicates the prefix of the value of the ConnectionString parameter.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****-clickhouse.clickhouseserver.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId      *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DBInstanceNetType *string `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEndpointRequest) GoString() string {
	return s.String()
}

func (s *DeleteEndpointRequest) GetComputingGroupId() *string {
	return s.ComputingGroupId
}

func (s *DeleteEndpointRequest) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DeleteEndpointRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteEndpointRequest) GetDBInstanceNetType() *string {
	return s.DBInstanceNetType
}

func (s *DeleteEndpointRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteEndpointRequest) SetComputingGroupId(v string) *DeleteEndpointRequest {
	s.ComputingGroupId = &v
	return s
}

func (s *DeleteEndpointRequest) SetConnectionString(v string) *DeleteEndpointRequest {
	s.ConnectionString = &v
	return s
}

func (s *DeleteEndpointRequest) SetDBInstanceId(v string) *DeleteEndpointRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteEndpointRequest) SetDBInstanceNetType(v string) *DeleteEndpointRequest {
	s.DBInstanceNetType = &v
	return s
}

func (s *DeleteEndpointRequest) SetRegionId(v string) *DeleteEndpointRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteEndpointRequest) Validate() error {
	return dara.Validate(s)
}
