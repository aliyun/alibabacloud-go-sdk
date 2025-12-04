// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterNetInfoItemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterNetworkType(v string) *DescribeDBClusterNetInfoItemsResponseBody
	GetClusterNetworkType() *string
	SetEnableSLB(v bool) *DescribeDBClusterNetInfoItemsResponseBody
	GetEnableSLB() *bool
	SetNetInfoItems(v *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItems) *DescribeDBClusterNetInfoItemsResponseBody
	GetNetInfoItems() *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItems
	SetRequestId(v string) *DescribeDBClusterNetInfoItemsResponseBody
	GetRequestId() *string
}

type DescribeDBClusterNetInfoItemsResponseBody struct {
	// The network type of the cluster. Only VPC is supported.
	//
	// example:
	//
	// VPC
	ClusterNetworkType *string `json:"ClusterNetworkType,omitempty" xml:"ClusterNetworkType,omitempty"`
	// Indicates whether Server Load Balancer (SLB) is activated in the VPC. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	EnableSLB *bool `json:"EnableSLB,omitempty" xml:"EnableSLB,omitempty"`
	// The network information about the cluster.
	NetInfoItems *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItems `json:"NetInfoItems,omitempty" xml:"NetInfoItems,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 9A23C87D-87DF-4DA0-A50E-CB13F4F7923D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterNetInfoItemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterNetInfoItemsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNetInfoItemsResponseBody) GetClusterNetworkType() *string {
	return s.ClusterNetworkType
}

func (s *DescribeDBClusterNetInfoItemsResponseBody) GetEnableSLB() *bool {
	return s.EnableSLB
}

func (s *DescribeDBClusterNetInfoItemsResponseBody) GetNetInfoItems() *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItems {
	return s.NetInfoItems
}

func (s *DescribeDBClusterNetInfoItemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterNetInfoItemsResponseBody) SetClusterNetworkType(v string) *DescribeDBClusterNetInfoItemsResponseBody {
	s.ClusterNetworkType = &v
	return s
}

func (s *DescribeDBClusterNetInfoItemsResponseBody) SetEnableSLB(v bool) *DescribeDBClusterNetInfoItemsResponseBody {
	s.EnableSLB = &v
	return s
}

func (s *DescribeDBClusterNetInfoItemsResponseBody) SetNetInfoItems(v *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItems) *DescribeDBClusterNetInfoItemsResponseBody {
	s.NetInfoItems = v
	return s
}

func (s *DescribeDBClusterNetInfoItemsResponseBody) SetRequestId(v string) *DescribeDBClusterNetInfoItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterNetInfoItemsResponseBody) Validate() error {
	if s.NetInfoItems != nil {
		if err := s.NetInfoItems.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBClusterNetInfoItemsResponseBodyNetInfoItems struct {
	NetInfoItem []*DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem `json:"NetInfoItem,omitempty" xml:"NetInfoItem,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterNetInfoItemsResponseBodyNetInfoItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterNetInfoItemsResponseBodyNetInfoItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItems) GetNetInfoItem() []*DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem {
	return s.NetInfoItem
}

func (s *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItems) SetNetInfoItem(v []*DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem) *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItems {
	s.NetInfoItem = v
	return s
}

func (s *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItems) Validate() error {
	if s.NetInfoItem != nil {
		for _, item := range s.NetInfoItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem struct {
	// The endpoint that is used to connect to the database.
	//
	// example:
	//
	// cc-bp1554t789i8e****.clickhouse.ads.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The HTTP port number.
	//
	// example:
	//
	// 8123
	HttpPort *string `json:"HttpPort,omitempty" xml:"HttpPort,omitempty"`
	// The HTTPS port number.
	//
	// example:
	//
	// 8443
	HttpsPort *string `json:"HttpsPort,omitempty" xml:"HttpsPort,omitempty"`
	// The IP address.
	//
	// example:
	//
	// 10.255.234.251
	IPAddress *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	// The port number that is used in Java Database Connectivity (JDBC).
	//
	// example:
	//
	// 3306
	JdbcPort *string `json:"JdbcPort,omitempty" xml:"JdbcPort,omitempty"`
	// The port of the MySQL instance.
	//
	// example:
	//
	// 9004
	MySQLPort *string `json:"MySQLPort,omitempty" xml:"MySQLPort,omitempty"`
	// The network type of the endpoint. Valid values:
	//
	// 	- Public: public endpoint
	//
	// 	- VPC: VPC
	//
	// example:
	//
	// VPC
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The vSwitch ID.
	//
	// >  If the value of the NetType parameter is set to Public, an empty string is returned.
	//
	// example:
	//
	// vsw-bp1gzt31twhlo0sa5****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The virtual private cloud (VPC) ID.
	//
	// >  If the value of the NetType parameter is set to Public, an empty string is returned.
	//
	// example:
	//
	// vpc-bp175iuvg8nxqraf2****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem) GetHttpPort() *string {
	return s.HttpPort
}

func (s *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem) GetHttpsPort() *string {
	return s.HttpsPort
}

func (s *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem) GetIPAddress() *string {
	return s.IPAddress
}

func (s *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem) GetJdbcPort() *string {
	return s.JdbcPort
}

func (s *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem) GetMySQLPort() *string {
	return s.MySQLPort
}

func (s *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem) GetNetType() *string {
	return s.NetType
}

func (s *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem) SetConnectionString(v string) *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem) SetHttpPort(v string) *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem {
	s.HttpPort = &v
	return s
}

func (s *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem) SetHttpsPort(v string) *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem {
	s.HttpsPort = &v
	return s
}

func (s *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem) SetIPAddress(v string) *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem {
	s.IPAddress = &v
	return s
}

func (s *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem) SetJdbcPort(v string) *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem {
	s.JdbcPort = &v
	return s
}

func (s *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem) SetMySQLPort(v string) *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem {
	s.MySQLPort = &v
	return s
}

func (s *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem) SetNetType(v string) *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem {
	s.NetType = &v
	return s
}

func (s *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem) SetVSwitchId(v string) *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem) SetVpcId(v string) *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem {
	s.VpcId = &v
	return s
}

func (s *DescribeDBClusterNetInfoItemsResponseBodyNetInfoItemsNetInfoItem) Validate() error {
	return dara.Validate(s)
}
