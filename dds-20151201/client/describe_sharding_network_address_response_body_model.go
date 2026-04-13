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
	SetConnectionStringSuffix(v string) *DescribeShardingNetworkAddressResponseBody
	GetConnectionStringSuffix() *string
	SetNetworkAddresses(v *DescribeShardingNetworkAddressResponseBodyNetworkAddresses) *DescribeShardingNetworkAddressResponseBody
	GetNetworkAddresses() *DescribeShardingNetworkAddressResponseBodyNetworkAddresses
	SetRequestId(v string) *DescribeShardingNetworkAddressResponseBody
	GetRequestId() *string
}

type DescribeShardingNetworkAddressResponseBody struct {
	CompatibleConnections  *DescribeShardingNetworkAddressResponseBodyCompatibleConnections `json:"CompatibleConnections,omitempty" xml:"CompatibleConnections,omitempty" type:"Struct"`
	ConnectionStringSuffix *string                                                          `json:"ConnectionStringSuffix,omitempty" xml:"ConnectionStringSuffix,omitempty"`
	NetworkAddresses       *DescribeShardingNetworkAddressResponseBodyNetworkAddresses      `json:"NetworkAddresses,omitempty" xml:"NetworkAddresses,omitempty" type:"Struct"`
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

func (s *DescribeShardingNetworkAddressResponseBody) GetConnectionStringSuffix() *string {
	return s.ConnectionStringSuffix
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

func (s *DescribeShardingNetworkAddressResponseBody) SetConnectionStringSuffix(v string) *DescribeShardingNetworkAddressResponseBody {
	s.ConnectionStringSuffix = &v
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
	ExpiredTime    *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	IPAddress      *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	NetworkAddress *string `json:"NetworkAddress,omitempty" xml:"NetworkAddress,omitempty"`
	NetworkType    *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Port           *string `json:"Port,omitempty" xml:"Port,omitempty"`
	VPCId          *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VswitchId      *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
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
	ConnectionType *string `json:"ConnectionType,omitempty" xml:"ConnectionType,omitempty"`
	ExpiredTime    *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	IPAddress      *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	NetworkAddress *string `json:"NetworkAddress,omitempty" xml:"NetworkAddress,omitempty"`
	NetworkType    *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	NodeId         *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeType       *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	Port           *string `json:"Port,omitempty" xml:"Port,omitempty"`
	Role           *string `json:"Role,omitempty" xml:"Role,omitempty"`
	TxtRecord      *string `json:"TxtRecord,omitempty" xml:"TxtRecord,omitempty"`
	VPCId          *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VswitchId      *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
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
