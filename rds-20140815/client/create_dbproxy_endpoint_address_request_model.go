// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBProxyEndpointAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionStringPrefix(v string) *CreateDBProxyEndpointAddressRequest
	GetConnectionStringPrefix() *string
	SetDBInstanceId(v string) *CreateDBProxyEndpointAddressRequest
	GetDBInstanceId() *string
	SetDBProxyConnectStringNetType(v string) *CreateDBProxyEndpointAddressRequest
	GetDBProxyConnectStringNetType() *string
	SetDBProxyEndpointId(v string) *CreateDBProxyEndpointAddressRequest
	GetDBProxyEndpointId() *string
	SetDBProxyEngineType(v string) *CreateDBProxyEndpointAddressRequest
	GetDBProxyEngineType() *string
	SetDBProxyNewConnectStringPort(v string) *CreateDBProxyEndpointAddressRequest
	GetDBProxyNewConnectStringPort() *string
	SetRegionId(v string) *CreateDBProxyEndpointAddressRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateDBProxyEndpointAddressRequest
	GetResourceGroupId() *string
	SetVPCId(v string) *CreateDBProxyEndpointAddressRequest
	GetVPCId() *string
	SetVSwitchId(v string) *CreateDBProxyEndpointAddressRequest
	GetVSwitchId() *string
}

type CreateDBProxyEndpointAddressRequest struct {
	// The prefix of the proxy endpoint Enter a custom prefix.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1234
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-t4n3axxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The network type of the proxy endpoint. Valid values:
	//
	// 	- **Public**: Internet
	//
	// 	- **VPC**: Virtual Private Cloud (VPC)
	//
	// 	- **Classic**: classic network
	//
	// Default value: **Classic**
	//
	// This parameter is required.
	//
	// example:
	//
	// Public
	DBProxyConnectStringNetType *string `json:"DBProxyConnectStringNetType,omitempty" xml:"DBProxyConnectStringNetType,omitempty"`
	// The proxy endpoint ID. You can call the DescribeDBProxyEndpoint operation to query the proxy endpoint ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ta9um4xxxxx
	DBProxyEndpointId *string `json:"DBProxyEndpointId,omitempty" xml:"DBProxyEndpointId,omitempty"`
	// A reserved parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// normal
	DBProxyEngineType *string `json:"DBProxyEngineType,omitempty" xml:"DBProxyEngineType,omitempty"`
	// The port number that is associated with the proxy endpoint.
	//
	// 	- If the instance runs MySQL, the default value is **3306**.
	//
	// 	- If the instance runs PostgreSQL, the default value is **5432**.
	//
	// example:
	//
	// 3306
	DBProxyNewConnectStringPort *string `json:"DBProxyNewConnectStringPort,omitempty" xml:"DBProxyNewConnectStringPort,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The ID of the VPC to which the proxy endpoint belongs. You can call the DescribeDBInstanceAttribute operation to query the information.
	//
	// >  This parameter must be specified when **DBProxyConnectStringNetType*	- is set to **VPC**.
	//
	// example:
	//
	// vpc-bpxxxxxx
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The ID of the vSwitch that is associated with the specified VPC. You can call the DescribeDBInstanceAttribute operation to query the vSwitch ID.
	//
	// >  This parameter must be specified when **DBProxyConnectStringNetType*	- is set to **VPC**.
	//
	// example:
	//
	// vsw-bpxxxxxx
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateDBProxyEndpointAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBProxyEndpointAddressRequest) GoString() string {
	return s.String()
}

func (s *CreateDBProxyEndpointAddressRequest) GetConnectionStringPrefix() *string {
	return s.ConnectionStringPrefix
}

func (s *CreateDBProxyEndpointAddressRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateDBProxyEndpointAddressRequest) GetDBProxyConnectStringNetType() *string {
	return s.DBProxyConnectStringNetType
}

func (s *CreateDBProxyEndpointAddressRequest) GetDBProxyEndpointId() *string {
	return s.DBProxyEndpointId
}

func (s *CreateDBProxyEndpointAddressRequest) GetDBProxyEngineType() *string {
	return s.DBProxyEngineType
}

func (s *CreateDBProxyEndpointAddressRequest) GetDBProxyNewConnectStringPort() *string {
	return s.DBProxyNewConnectStringPort
}

func (s *CreateDBProxyEndpointAddressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDBProxyEndpointAddressRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDBProxyEndpointAddressRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *CreateDBProxyEndpointAddressRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateDBProxyEndpointAddressRequest) SetConnectionStringPrefix(v string) *CreateDBProxyEndpointAddressRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *CreateDBProxyEndpointAddressRequest) SetDBInstanceId(v string) *CreateDBProxyEndpointAddressRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDBProxyEndpointAddressRequest) SetDBProxyConnectStringNetType(v string) *CreateDBProxyEndpointAddressRequest {
	s.DBProxyConnectStringNetType = &v
	return s
}

func (s *CreateDBProxyEndpointAddressRequest) SetDBProxyEndpointId(v string) *CreateDBProxyEndpointAddressRequest {
	s.DBProxyEndpointId = &v
	return s
}

func (s *CreateDBProxyEndpointAddressRequest) SetDBProxyEngineType(v string) *CreateDBProxyEndpointAddressRequest {
	s.DBProxyEngineType = &v
	return s
}

func (s *CreateDBProxyEndpointAddressRequest) SetDBProxyNewConnectStringPort(v string) *CreateDBProxyEndpointAddressRequest {
	s.DBProxyNewConnectStringPort = &v
	return s
}

func (s *CreateDBProxyEndpointAddressRequest) SetRegionId(v string) *CreateDBProxyEndpointAddressRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBProxyEndpointAddressRequest) SetResourceGroupId(v string) *CreateDBProxyEndpointAddressRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDBProxyEndpointAddressRequest) SetVPCId(v string) *CreateDBProxyEndpointAddressRequest {
	s.VPCId = &v
	return s
}

func (s *CreateDBProxyEndpointAddressRequest) SetVSwitchId(v string) *CreateDBProxyEndpointAddressRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDBProxyEndpointAddressRequest) Validate() error {
	return dara.Validate(s)
}
