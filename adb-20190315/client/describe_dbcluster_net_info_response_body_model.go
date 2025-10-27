// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterNetInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterNetworkType(v string) *DescribeDBClusterNetInfoResponseBody
	GetClusterNetworkType() *string
	SetItems(v *DescribeDBClusterNetInfoResponseBodyItems) *DescribeDBClusterNetInfoResponseBody
	GetItems() *DescribeDBClusterNetInfoResponseBodyItems
	SetRequestId(v string) *DescribeDBClusterNetInfoResponseBody
	GetRequestId() *string
}

type DescribeDBClusterNetInfoResponseBody struct {
	// The network type of the cluster.
	//
	// example:
	//
	// VPC
	ClusterNetworkType *string `json:"ClusterNetworkType,omitempty" xml:"ClusterNetworkType,omitempty"`
	// The queried network information about the cluster.
	Items *DescribeDBClusterNetInfoResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterNetInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterNetInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNetInfoResponseBody) GetClusterNetworkType() *string {
	return s.ClusterNetworkType
}

func (s *DescribeDBClusterNetInfoResponseBody) GetItems() *DescribeDBClusterNetInfoResponseBodyItems {
	return s.Items
}

func (s *DescribeDBClusterNetInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterNetInfoResponseBody) SetClusterNetworkType(v string) *DescribeDBClusterNetInfoResponseBody {
	s.ClusterNetworkType = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBody) SetItems(v *DescribeDBClusterNetInfoResponseBodyItems) *DescribeDBClusterNetInfoResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBody) SetRequestId(v string) *DescribeDBClusterNetInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBClusterNetInfoResponseBodyItems struct {
	Address []*DescribeDBClusterNetInfoResponseBodyItemsAddress `json:"Address,omitempty" xml:"Address,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterNetInfoResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterNetInfoResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNetInfoResponseBodyItems) GetAddress() []*DescribeDBClusterNetInfoResponseBodyItemsAddress {
	return s.Address
}

func (s *DescribeDBClusterNetInfoResponseBodyItems) SetAddress(v []*DescribeDBClusterNetInfoResponseBodyItemsAddress) *DescribeDBClusterNetInfoResponseBodyItems {
	s.Address = v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyItems) Validate() error {
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

type DescribeDBClusterNetInfoResponseBodyItemsAddress struct {
	// The endpoint of the cluster.
	//
	// example:
	//
	// am-bpxxxxxxxx.ads.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The prefix of the cluster endpoint.
	//
	// example:
	//
	// am-bpxxxxxxxx89k51380
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// The IP address.
	//
	// example:
	//
	// 192.168.x.x
	IPAddress *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	// The network type of the cluster. Valid values:
	//
	// 	- **Public**: public endpoint.
	//
	// 	- **VPC**: Virtual Private Cloud (VPC) endpoint.
	//
	// example:
	//
	// VPC
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The port number that is used to connect to the cluster.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The VPC ID.
	//
	// >  If NetType is set to Public, an empty string is returned for this parameter.
	//
	// example:
	//
	// vpc-xxxxxxxxx
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID.
	//
	// >  If NetType is set to Public, an empty string is returned for this parameter.
	//
	// example:
	//
	// vsw-xxxxxx
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeDBClusterNetInfoResponseBodyItemsAddress) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterNetInfoResponseBodyItemsAddress) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNetInfoResponseBodyItemsAddress) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeDBClusterNetInfoResponseBodyItemsAddress) GetConnectionStringPrefix() *string {
	return s.ConnectionStringPrefix
}

func (s *DescribeDBClusterNetInfoResponseBodyItemsAddress) GetIPAddress() *string {
	return s.IPAddress
}

func (s *DescribeDBClusterNetInfoResponseBodyItemsAddress) GetNetType() *string {
	return s.NetType
}

func (s *DescribeDBClusterNetInfoResponseBodyItemsAddress) GetPort() *string {
	return s.Port
}

func (s *DescribeDBClusterNetInfoResponseBodyItemsAddress) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeDBClusterNetInfoResponseBodyItemsAddress) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBClusterNetInfoResponseBodyItemsAddress) SetConnectionString(v string) *DescribeDBClusterNetInfoResponseBodyItemsAddress {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyItemsAddress) SetConnectionStringPrefix(v string) *DescribeDBClusterNetInfoResponseBodyItemsAddress {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyItemsAddress) SetIPAddress(v string) *DescribeDBClusterNetInfoResponseBodyItemsAddress {
	s.IPAddress = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyItemsAddress) SetNetType(v string) *DescribeDBClusterNetInfoResponseBodyItemsAddress {
	s.NetType = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyItemsAddress) SetPort(v string) *DescribeDBClusterNetInfoResponseBodyItemsAddress {
	s.Port = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyItemsAddress) SetVPCId(v string) *DescribeDBClusterNetInfoResponseBodyItemsAddress {
	s.VPCId = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyItemsAddress) SetVSwitchId(v string) *DescribeDBClusterNetInfoResponseBodyItemsAddress {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBClusterNetInfoResponseBodyItemsAddress) Validate() error {
	return dara.Validate(s)
}
