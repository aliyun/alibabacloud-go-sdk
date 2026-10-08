// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBProxyEndpointAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DeleteDBProxyEndpointAddressRequest
	GetDBInstanceId() *string
	SetDBProxyConnectStringNetType(v string) *DeleteDBProxyEndpointAddressRequest
	GetDBProxyConnectStringNetType() *string
	SetDBProxyEndpointId(v string) *DeleteDBProxyEndpointAddressRequest
	GetDBProxyEndpointId() *string
	SetDBProxyEngineType(v string) *DeleteDBProxyEndpointAddressRequest
	GetDBProxyEngineType() *string
	SetRegionId(v string) *DeleteDBProxyEndpointAddressRequest
	GetRegionId() *string
}

type DeleteDBProxyEndpointAddressRequest struct {
	// The instance ID. You can call DescribeDBInstances to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-t4n3a****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The network type of the database proxy endpoint to delete. Valid values:
	//
	// 	- **Public**: Internet
	//
	// 	- **VPC**: internal network (VPC)
	//
	// 	- **Classic**: internal network (classic network)
	//
	// Default value: **Classic**.
	//
	// > - You cannot delete the internal endpoint that is created by default.
	//
	// > - ApsaraDB RDS for PostgreSQL supports only **Public*	- and **VPC**.
	//
	// This parameter is required.
	//
	// example:
	//
	// Public
	DBProxyConnectStringNetType *string `json:"DBProxyConnectStringNetType,omitempty" xml:"DBProxyConnectStringNetType,omitempty"`
	// The ID of the database proxy endpoint. You can call DescribeDBProxyEndpoint to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ta9um4****
	DBProxyEndpointId *string `json:"DBProxyEndpointId,omitempty" xml:"DBProxyEndpointId,omitempty"`
	// A deprecated parameter. You do not need to configure this parameter.
	//
	// example:
	//
	// normal
	DBProxyEngineType *string `json:"DBProxyEngineType,omitempty" xml:"DBProxyEngineType,omitempty"`
	// The region ID. You can call DescribeRegions to query the available regions.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteDBProxyEndpointAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBProxyEndpointAddressRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBProxyEndpointAddressRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteDBProxyEndpointAddressRequest) GetDBProxyConnectStringNetType() *string {
	return s.DBProxyConnectStringNetType
}

func (s *DeleteDBProxyEndpointAddressRequest) GetDBProxyEndpointId() *string {
	return s.DBProxyEndpointId
}

func (s *DeleteDBProxyEndpointAddressRequest) GetDBProxyEngineType() *string {
	return s.DBProxyEngineType
}

func (s *DeleteDBProxyEndpointAddressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDBProxyEndpointAddressRequest) SetDBInstanceId(v string) *DeleteDBProxyEndpointAddressRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDBProxyEndpointAddressRequest) SetDBProxyConnectStringNetType(v string) *DeleteDBProxyEndpointAddressRequest {
	s.DBProxyConnectStringNetType = &v
	return s
}

func (s *DeleteDBProxyEndpointAddressRequest) SetDBProxyEndpointId(v string) *DeleteDBProxyEndpointAddressRequest {
	s.DBProxyEndpointId = &v
	return s
}

func (s *DeleteDBProxyEndpointAddressRequest) SetDBProxyEngineType(v string) *DeleteDBProxyEndpointAddressRequest {
	s.DBProxyEngineType = &v
	return s
}

func (s *DeleteDBProxyEndpointAddressRequest) SetRegionId(v string) *DeleteDBProxyEndpointAddressRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDBProxyEndpointAddressRequest) Validate() error {
	return dara.Validate(s)
}
