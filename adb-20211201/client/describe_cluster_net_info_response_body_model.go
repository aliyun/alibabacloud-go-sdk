// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterNetInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterNetworkType(v string) *DescribeClusterNetInfoResponseBody
	GetClusterNetworkType() *string
	SetItems(v *DescribeClusterNetInfoResponseBodyItems) *DescribeClusterNetInfoResponseBody
	GetItems() *DescribeClusterNetInfoResponseBodyItems
	SetRequestId(v string) *DescribeClusterNetInfoResponseBody
	GetRequestId() *string
}

type DescribeClusterNetInfoResponseBody struct {
	// The network type of the cluster. Only the Virtual Private Cloud (VPC) network type is supported. **VPC*	- is returned.
	//
	// example:
	//
	// VPC
	ClusterNetworkType *string `json:"ClusterNetworkType,omitempty" xml:"ClusterNetworkType,omitempty"`
	// The queried network information about the cluster.
	Items *DescribeClusterNetInfoResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 69A29B65-CD0C-52B1-BE42-8B454569747F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClusterNetInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNetInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterNetInfoResponseBody) GetClusterNetworkType() *string {
	return s.ClusterNetworkType
}

func (s *DescribeClusterNetInfoResponseBody) GetItems() *DescribeClusterNetInfoResponseBodyItems {
	return s.Items
}

func (s *DescribeClusterNetInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClusterNetInfoResponseBody) SetClusterNetworkType(v string) *DescribeClusterNetInfoResponseBody {
	s.ClusterNetworkType = &v
	return s
}

func (s *DescribeClusterNetInfoResponseBody) SetItems(v *DescribeClusterNetInfoResponseBodyItems) *DescribeClusterNetInfoResponseBody {
	s.Items = v
	return s
}

func (s *DescribeClusterNetInfoResponseBody) SetRequestId(v string) *DescribeClusterNetInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterNetInfoResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClusterNetInfoResponseBodyItems struct {
	Address []*DescribeClusterNetInfoResponseBodyItemsAddress `json:"Address,omitempty" xml:"Address,omitempty" type:"Repeated"`
}

func (s DescribeClusterNetInfoResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNetInfoResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeClusterNetInfoResponseBodyItems) GetAddress() []*DescribeClusterNetInfoResponseBodyItemsAddress {
	return s.Address
}

func (s *DescribeClusterNetInfoResponseBodyItems) SetAddress(v []*DescribeClusterNetInfoResponseBodyItemsAddress) *DescribeClusterNetInfoResponseBodyItems {
	s.Address = v
	return s
}

func (s *DescribeClusterNetInfoResponseBodyItems) Validate() error {
	if s.Address != nil {
		for _, item := range s.Address {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeClusterNetInfoResponseBodyItemsAddress struct {
	// The endpoint of the cluster.
	//
	// 	- If NetType is set to VPC, the VPC endpoint of the cluster is returned.
	//
	// 	- If NetType is set to Public, the public endpoint of the cluster is returned.
	//
	// example:
	//
	// amv-wz9dqvn0o7****.ads.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The prefix of the endpoint.
	//
	// 	- If NetType is set to VPC, the prefix of the VPC endpoint is returned.
	//
	// 	- If NetType is set to Public, the prefix of the public endpoint is returned.
	//
	// example:
	//
	// amv-wz9dqvn0o7****
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// The IP address of the endpoint.
	//
	// 	- If NetType is set to VPC, the private IP address of the cluster is returned.
	//
	// 	- If NetType is set to Public, the public IP address of the cluster is returned.
	//
	// example:
	//
	// 192.168.xx.xx
	IPAddress *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	// The network type of the cluster. Valid values:
	//
	// 	- **Public**: Internet.
	//
	// 	- **VPC**: VPC.
	//
	// example:
	//
	// VPC
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The port number that is used to connect to the cluster. **3306*	- is returned.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ports.
	Ports *DescribeClusterNetInfoResponseBodyItemsAddressPorts `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Struct"`
	// The VPC ID.
	//
	// >  If NetType is set to Public, an empty string is returned.
	//
	// example:
	//
	// vpc-8vbhucmd5b****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID of the cluster.
	//
	// >  If NetType is set to Public, an empty string is returned.
	//
	// example:
	//
	// vsw-bp1syh8vvw8yec****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeClusterNetInfoResponseBodyItemsAddress) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNetInfoResponseBodyItemsAddress) GoString() string {
	return s.String()
}

func (s *DescribeClusterNetInfoResponseBodyItemsAddress) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeClusterNetInfoResponseBodyItemsAddress) GetConnectionStringPrefix() *string {
	return s.ConnectionStringPrefix
}

func (s *DescribeClusterNetInfoResponseBodyItemsAddress) GetIPAddress() *string {
	return s.IPAddress
}

func (s *DescribeClusterNetInfoResponseBodyItemsAddress) GetNetType() *string {
	return s.NetType
}

func (s *DescribeClusterNetInfoResponseBodyItemsAddress) GetPort() *string {
	return s.Port
}

func (s *DescribeClusterNetInfoResponseBodyItemsAddress) GetPorts() *DescribeClusterNetInfoResponseBodyItemsAddressPorts {
	return s.Ports
}

func (s *DescribeClusterNetInfoResponseBodyItemsAddress) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeClusterNetInfoResponseBodyItemsAddress) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeClusterNetInfoResponseBodyItemsAddress) SetConnectionString(v string) *DescribeClusterNetInfoResponseBodyItemsAddress {
	s.ConnectionString = &v
	return s
}

func (s *DescribeClusterNetInfoResponseBodyItemsAddress) SetConnectionStringPrefix(v string) *DescribeClusterNetInfoResponseBodyItemsAddress {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *DescribeClusterNetInfoResponseBodyItemsAddress) SetIPAddress(v string) *DescribeClusterNetInfoResponseBodyItemsAddress {
	s.IPAddress = &v
	return s
}

func (s *DescribeClusterNetInfoResponseBodyItemsAddress) SetNetType(v string) *DescribeClusterNetInfoResponseBodyItemsAddress {
	s.NetType = &v
	return s
}

func (s *DescribeClusterNetInfoResponseBodyItemsAddress) SetPort(v string) *DescribeClusterNetInfoResponseBodyItemsAddress {
	s.Port = &v
	return s
}

func (s *DescribeClusterNetInfoResponseBodyItemsAddress) SetPorts(v *DescribeClusterNetInfoResponseBodyItemsAddressPorts) *DescribeClusterNetInfoResponseBodyItemsAddress {
	s.Ports = v
	return s
}

func (s *DescribeClusterNetInfoResponseBodyItemsAddress) SetVPCId(v string) *DescribeClusterNetInfoResponseBodyItemsAddress {
	s.VPCId = &v
	return s
}

func (s *DescribeClusterNetInfoResponseBodyItemsAddress) SetVSwitchId(v string) *DescribeClusterNetInfoResponseBodyItemsAddress {
	s.VSwitchId = &v
	return s
}

func (s *DescribeClusterNetInfoResponseBodyItemsAddress) Validate() error {
	if s.Ports != nil {
		if err := s.Ports.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClusterNetInfoResponseBodyItemsAddressPorts struct {
	Ports []*DescribeClusterNetInfoResponseBodyItemsAddressPortsPorts `json:"ports,omitempty" xml:"ports,omitempty" type:"Repeated"`
}

func (s DescribeClusterNetInfoResponseBodyItemsAddressPorts) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNetInfoResponseBodyItemsAddressPorts) GoString() string {
	return s.String()
}

func (s *DescribeClusterNetInfoResponseBodyItemsAddressPorts) GetPorts() []*DescribeClusterNetInfoResponseBodyItemsAddressPortsPorts {
	return s.Ports
}

func (s *DescribeClusterNetInfoResponseBodyItemsAddressPorts) SetPorts(v []*DescribeClusterNetInfoResponseBodyItemsAddressPortsPorts) *DescribeClusterNetInfoResponseBodyItemsAddressPorts {
	s.Ports = v
	return s
}

func (s *DescribeClusterNetInfoResponseBodyItemsAddressPorts) Validate() error {
	if s.Ports != nil {
		for _, item := range s.Ports {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeClusterNetInfoResponseBodyItemsAddressPortsPorts struct {
	// The port.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The type of the protocol. Valid values:
	//
	// 	- **tcp**
	//
	// 	- **http**
	//
	// 	- **https**
	//
	// 	- **mysql**
	//
	// example:
	//
	// mysql
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s DescribeClusterNetInfoResponseBodyItemsAddressPortsPorts) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNetInfoResponseBodyItemsAddressPortsPorts) GoString() string {
	return s.String()
}

func (s *DescribeClusterNetInfoResponseBodyItemsAddressPortsPorts) GetPort() *string {
	return s.Port
}

func (s *DescribeClusterNetInfoResponseBodyItemsAddressPortsPorts) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeClusterNetInfoResponseBodyItemsAddressPortsPorts) SetPort(v string) *DescribeClusterNetInfoResponseBodyItemsAddressPortsPorts {
	s.Port = &v
	return s
}

func (s *DescribeClusterNetInfoResponseBodyItemsAddressPortsPorts) SetProtocol(v string) *DescribeClusterNetInfoResponseBodyItemsAddressPortsPorts {
	s.Protocol = &v
	return s
}

func (s *DescribeClusterNetInfoResponseBodyItemsAddressPortsPorts) Validate() error {
	return dara.Validate(s)
}
