// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeShardingNetworkAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCompatibleConnections(v *DescribeShardingNetworkAddressResponseBodyCompatibleConnections) *DescribeShardingNetworkAddressResponseBody
	GetCompatibleConnections() *DescribeShardingNetworkAddressResponseBodyCompatibleConnections
	SetNetworkAddresses(v *DescribeShardingNetworkAddressResponseBodyNetworkAddresses) *DescribeShardingNetworkAddressResponseBody
	GetNetworkAddresses() *DescribeShardingNetworkAddressResponseBodyNetworkAddresses
	SetRequestId(v string) *DescribeShardingNetworkAddressResponseBody
	GetRequestId() *string
}

type DescribeShardingNetworkAddressResponseBody struct {
	// The endpoints of DynamoDB-compatible instances.
	CompatibleConnections *DescribeShardingNetworkAddressResponseBodyCompatibleConnections `json:"CompatibleConnections,omitempty" xml:"CompatibleConnections,omitempty" type:"Struct"`
	// The endpoints of the ApsaraDB for MongoDB sharded cluster instance.
	NetworkAddresses *DescribeShardingNetworkAddressResponseBodyNetworkAddresses `json:"NetworkAddresses,omitempty" xml:"NetworkAddresses,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 18D8AAFD-6BEB-420F-8164-810CB0C0AA39
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeShardingNetworkAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeShardingNetworkAddressResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeShardingNetworkAddressResponseBody) GetCompatibleConnections() *DescribeShardingNetworkAddressResponseBodyCompatibleConnections {
	return s.CompatibleConnections
}

func (s *DescribeShardingNetworkAddressResponseBody) GetNetworkAddresses() *DescribeShardingNetworkAddressResponseBodyNetworkAddresses {
	return s.NetworkAddresses
}

func (s *DescribeShardingNetworkAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeShardingNetworkAddressResponseBody) SetCompatibleConnections(v *DescribeShardingNetworkAddressResponseBodyCompatibleConnections) *DescribeShardingNetworkAddressResponseBody {
	s.CompatibleConnections = v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBody) SetNetworkAddresses(v *DescribeShardingNetworkAddressResponseBodyNetworkAddresses) *DescribeShardingNetworkAddressResponseBody {
	s.NetworkAddresses = v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBody) SetRequestId(v string) *DescribeShardingNetworkAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBody) Validate() error {
	if s.CompatibleConnections != nil {
		if err := s.CompatibleConnections.Validate(); err != nil {
			return err
		}
	}
	if s.NetworkAddresses != nil {
		if err := s.NetworkAddresses.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeShardingNetworkAddressResponseBodyCompatibleConnections struct {
	CompatibleConnection []*DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection `json:"CompatibleConnection,omitempty" xml:"CompatibleConnection,omitempty" type:"Repeated"`
}

func (s DescribeShardingNetworkAddressResponseBodyCompatibleConnections) String() string {
	return dara.Prettify(s)
}

func (s DescribeShardingNetworkAddressResponseBodyCompatibleConnections) GoString() string {
	return s.String()
}

func (s *DescribeShardingNetworkAddressResponseBodyCompatibleConnections) GetCompatibleConnection() []*DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection {
	return s.CompatibleConnection
}

func (s *DescribeShardingNetworkAddressResponseBodyCompatibleConnections) SetCompatibleConnection(v []*DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection) *DescribeShardingNetworkAddressResponseBodyCompatibleConnections {
	s.CompatibleConnection = v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBodyCompatibleConnections) Validate() error {
	if s.CompatibleConnection != nil {
		for _, item := range s.CompatibleConnection {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection struct {
	// The remaining duration of the classic network endpoint. Unit: seconds.
	//
	// example:
	//
	// 2591963
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The IP address of the instance.
	//
	// example:
	//
	// 10.140.xxx.xx
	IPAddress *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	// The endpoint of the instance.
	//
	// example:
	//
	// dds-bpxxxxxxxxxxxxxx.mongodb.rds.aliyuncs.com
	NetworkAddress *string `json:"NetworkAddress,omitempty" xml:"NetworkAddress,omitempty"`
	// The network type of the instance.
	//
	// 	- **VPC**: virtual private cloud
	//
	// 	- **Classic**: classic network
	//
	// 	- **Public**: the Internet
	//
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The port that is used to connect to the instance.
	//
	// example:
	//
	// 3717
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The VPC ID of the instance.
	//
	// >  This parameter is returned when the network type is **VPC**.
	//
	// example:
	//
	// vpc-bpxxxxxxxx
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The ID of the vSwitch in the Virtual Private Cloud (VPC).
	//
	// >  This parameter is returned when the network type is **VPC**.
	//
	// example:
	//
	// vsw-bpxxxxxxxx
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection) String() string {
	return dara.Prettify(s)
}

func (s DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection) GoString() string {
	return s.String()
}

func (s *DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection) GetIPAddress() *string {
	return s.IPAddress
}

func (s *DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection) GetNetworkAddress() *string {
	return s.NetworkAddress
}

func (s *DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection) GetPort() *string {
	return s.Port
}

func (s *DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection) GetVswitchId() *string {
	return s.VswitchId
}

func (s *DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection) SetExpiredTime(v string) *DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection) SetIPAddress(v string) *DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection {
	s.IPAddress = &v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection) SetNetworkAddress(v string) *DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection {
	s.NetworkAddress = &v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection) SetNetworkType(v string) *DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection {
	s.NetworkType = &v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection) SetPort(v string) *DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection {
	s.Port = &v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection) SetVPCId(v string) *DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection {
	s.VPCId = &v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection) SetVswitchId(v string) *DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection {
	s.VswitchId = &v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBodyCompatibleConnectionsCompatibleConnection) Validate() error {
	return dara.Validate(s)
}

type DescribeShardingNetworkAddressResponseBodyNetworkAddresses struct {
	NetworkAddress []*DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress `json:"NetworkAddress,omitempty" xml:"NetworkAddress,omitempty" type:"Repeated"`
}

func (s DescribeShardingNetworkAddressResponseBodyNetworkAddresses) String() string {
	return dara.Prettify(s)
}

func (s DescribeShardingNetworkAddressResponseBodyNetworkAddresses) GoString() string {
	return s.String()
}

func (s *DescribeShardingNetworkAddressResponseBodyNetworkAddresses) GetNetworkAddress() []*DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress {
	return s.NetworkAddress
}

func (s *DescribeShardingNetworkAddressResponseBodyNetworkAddresses) SetNetworkAddress(v []*DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) *DescribeShardingNetworkAddressResponseBodyNetworkAddresses {
	s.NetworkAddress = v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBodyNetworkAddresses) Validate() error {
	if s.NetworkAddress != nil {
		for _, item := range s.NetworkAddress {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress struct {
	// The public endpoint type. Valid values:
	//
	// 	- **SRV**
	//
	// 	- **Normal**
	//
	// example:
	//
	// SRV
	ConnectionType *string `json:"ConnectionType,omitempty" xml:"ConnectionType,omitempty"`
	// The remaining duration of the classic network endpoint. Unit: seconds.
	//
	// example:
	//
	// 2591963
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The IP address of the instance.
	//
	// example:
	//
	// 10.140.xxx.xx
	IPAddress *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	// The connection string of the instance.
	//
	// example:
	//
	// s-bpxxxxxxxx.mongodb.rds.aliyuncs.com
	NetworkAddress *string `json:"NetworkAddress,omitempty" xml:"NetworkAddress,omitempty"`
	// The network type of the instance.
	//
	// 	- **VPC**: virtual private cloud
	//
	// 	- **Classic**: classic network
	//
	// 	- **Public**: the Internet
	//
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The ID of the mongos node.
	//
	// example:
	//
	// s-bpxxxxxxxx
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The type of the node. Valid values:
	//
	// 	- **mongos**: mongos node
	//
	// 	- **shard**: shard node
	//
	// 	- **configserver**: Configserver node
	//
	// example:
	//
	// mongos
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The port that is used to connect to the instance.
	//
	// example:
	//
	// 3717
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The role of the node. Valid values:
	//
	// 	- Primary
	//
	// 	- Secondary
	//
	// example:
	//
	// Primary
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// Txt record which can be used to store MongoDB-related meta data, such as version, configuration parameters and etc. With the combination of txt record and other technology, for example SRV record, the MongoDB client can complete the complex service discovery and configuration passing.
	//
	// example:
	//
	// mongo.example.com. IN TXT "config=replicaSet=myReplicaSet"
	TxtRecord *string `json:"TxtRecord,omitempty" xml:"TxtRecord,omitempty"`
	// The VPC ID of the instance.
	//
	// >  This parameter is returned when the network type is **VPC**.
	//
	// example:
	//
	// vpc-bpxxxxxxxx
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The ID of the vSwitch in the VPC.
	//
	// >  This parameter is returned when the network type is **VPC**.
	//
	// example:
	//
	// vsw-bpxxxxxxxx
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) String() string {
	return dara.Prettify(s)
}

func (s DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) GoString() string {
	return s.String()
}

func (s *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) GetConnectionType() *string {
	return s.ConnectionType
}

func (s *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) GetIPAddress() *string {
	return s.IPAddress
}

func (s *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) GetNetworkAddress() *string {
	return s.NetworkAddress
}

func (s *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) GetNodeType() *string {
	return s.NodeType
}

func (s *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) GetPort() *string {
	return s.Port
}

func (s *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) GetRole() *string {
	return s.Role
}

func (s *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) GetTxtRecord() *string {
	return s.TxtRecord
}

func (s *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) GetVswitchId() *string {
	return s.VswitchId
}

func (s *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) SetConnectionType(v string) *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress {
	s.ConnectionType = &v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) SetExpiredTime(v string) *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) SetIPAddress(v string) *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress {
	s.IPAddress = &v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) SetNetworkAddress(v string) *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress {
	s.NetworkAddress = &v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) SetNetworkType(v string) *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress {
	s.NetworkType = &v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) SetNodeId(v string) *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress {
	s.NodeId = &v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) SetNodeType(v string) *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress {
	s.NodeType = &v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) SetPort(v string) *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress {
	s.Port = &v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) SetRole(v string) *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress {
	s.Role = &v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) SetTxtRecord(v string) *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress {
	s.TxtRecord = &v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) SetVPCId(v string) *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress {
	s.VPCId = &v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) SetVswitchId(v string) *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress {
	s.VswitchId = &v
	return s
}

func (s *DescribeShardingNetworkAddressResponseBodyNetworkAddressesNetworkAddress) Validate() error {
	return dara.Validate(s)
}
