// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComputingGroupId(v string) *CreateEndpointRequest
	GetComputingGroupId() *string
	SetConnectionPrefix(v string) *CreateEndpointRequest
	GetConnectionPrefix() *string
	SetDBInstanceId(v string) *CreateEndpointRequest
	GetDBInstanceId() *string
	SetDBInstanceNetType(v string) *CreateEndpointRequest
	GetDBInstanceNetType() *string
	SetRegionId(v string) *CreateEndpointRequest
	GetRegionId() *string
}

type CreateEndpointRequest struct {
	ComputingGroupId *string `json:"ComputingGroupId,omitempty" xml:"ComputingGroupId,omitempty"`
	// The prefix of the new endpoint. The prefix of the ConnectionString parameter.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****-clickhouse.clickhouseserver.rds.aliyuncs.com
	ConnectionPrefix *string `json:"ConnectionPrefix,omitempty" xml:"ConnectionPrefix,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The network type.
	//
	// Valid values:
	//
	// 	- Public
	//
	// example:
	//
	// Public
	DBInstanceNetType *string `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEndpointRequest) GoString() string {
	return s.String()
}

func (s *CreateEndpointRequest) GetComputingGroupId() *string {
	return s.ComputingGroupId
}

func (s *CreateEndpointRequest) GetConnectionPrefix() *string {
	return s.ConnectionPrefix
}

func (s *CreateEndpointRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateEndpointRequest) GetDBInstanceNetType() *string {
	return s.DBInstanceNetType
}

func (s *CreateEndpointRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateEndpointRequest) SetComputingGroupId(v string) *CreateEndpointRequest {
	s.ComputingGroupId = &v
	return s
}

func (s *CreateEndpointRequest) SetConnectionPrefix(v string) *CreateEndpointRequest {
	s.ConnectionPrefix = &v
	return s
}

func (s *CreateEndpointRequest) SetDBInstanceId(v string) *CreateEndpointRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateEndpointRequest) SetDBInstanceNetType(v string) *CreateEndpointRequest {
	s.DBInstanceNetType = &v
	return s
}

func (s *CreateEndpointRequest) SetRegionId(v string) *CreateEndpointRequest {
	s.RegionId = &v
	return s
}

func (s *CreateEndpointRequest) Validate() error {
	return dara.Validate(s)
}
