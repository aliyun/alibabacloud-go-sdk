// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterEndpointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeDBClusterEndpointsResponseBodyItems) *DescribeDBClusterEndpointsResponseBody
	GetItems() []*DescribeDBClusterEndpointsResponseBodyItems
	SetRequestId(v string) *DescribeDBClusterEndpointsResponseBody
	GetRequestId() *string
}

type DescribeDBClusterEndpointsResponseBody struct {
	// A list of cluster endpoints.
	Items []*DescribeDBClusterEndpointsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 2DC120BF-6EBA-4C63-BE99-B09F9E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterEndpointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterEndpointsResponseBody) GetItems() []*DescribeDBClusterEndpointsResponseBodyItems {
	return s.Items
}

func (s *DescribeDBClusterEndpointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterEndpointsResponseBody) SetItems(v []*DescribeDBClusterEndpointsResponseBodyItems) *DescribeDBClusterEndpointsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBody) SetRequestId(v string) *DescribeDBClusterEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBClusterEndpointsResponseBodyItems struct {
	// The connection addresses for the endpoint.
	AddressItems []*DescribeDBClusterEndpointsResponseBodyItemsAddressItems `json:"AddressItems,omitempty" xml:"AddressItems,omitempty" type:"Repeated"`
	// Indicates whether new nodes are automatically added to the default cluster endpoint. Valid values:
	//
	// - **Enable**
	//
	// - **Disable**
	//
	// example:
	//
	// Enable
	AutoAddNewNodes *string `json:"AutoAddNewNodes,omitempty" xml:"AutoAddNewNodes,omitempty"`
	// The connection string.
	//
	// example:
	//
	// ********.rwlb.polardb-pg-public.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// pc-bp1s826a1up******
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The endpoint name.
	//
	// example:
	//
	// test
	DBEndpointDescription *string `json:"DBEndpointDescription,omitempty" xml:"DBEndpointDescription,omitempty"`
	// The endpoint ID.
	//
	// example:
	//
	// pe-*************
	DBEndpointId *string `json:"DBEndpointId,omitempty" xml:"DBEndpointId,omitempty"`
	// The advanced settings for the cluster endpoint.
	//
	// - **DistributedTransaction**: The transaction splitting status. Valid values:
	//
	//   - **on**: enabled
	//
	//   - **off**: disabled
	//
	// - **ConsistLevel**: The consistency level. Valid values:
	//
	//   - **0**: eventual consistency
	//
	//   - **1**: session consistency
	//
	//   - **2**: global consistency
	//
	// - **LoadBalanceStrategy**: The load balancing policy. The value is fixed to **load**, which indicates load-based scheduling.
	//
	// - **MasterAcceptReads**: Indicates whether the primary node accepts read requests. Valid values:
	//
	//   - **on**: The primary node accepts read requests.
	//
	//   - **off**: The primary node does not accept read requests.
	//
	// example:
	//
	// {\\"DistributedTransaction\\":\\"off\\",\\"ConsistLevel\\":\\"0\\",\\"LoadBalanceStrategy\\":\\"load\\",\\"MasterAcceptReads\\":\\"on\\"}
	EndpointConfig *string `json:"EndpointConfig,omitempty" xml:"EndpointConfig,omitempty"`
	// The type of the cluster endpoint. Valid values:
	//
	// - **Cluster**: the default cluster endpoint.
	//
	// - **Primary**: the primary endpoint.
	//
	// - **Custom**: a custom cluster endpoint.
	//
	// example:
	//
	// Custom
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The network type. Valid values:
	//
	// - **Public**: public network
	//
	// - **Private**: private network
	//
	// example:
	//
	// Private
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The role of each node in the endpoint. The primary node has the **Writer*	- role. Read-only nodes have numbered roles, such as **Reader1**, **Reader2**, and so on.
	//
	// > This parameter is supported only by PolarDB for PostgreSQL clusters and PolarDB for PostgreSQL (compatible with Oracle) clusters.
	//
	// example:
	//
	// Reader1
	NodeWithRoles *string `json:"NodeWithRoles,omitempty" xml:"NodeWithRoles,omitempty"`
	// The list of nodes configured for the endpoint.
	//
	// example:
	//
	// pi-***************,pi-***************
	Nodes *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
	// The policy for handling global consistency read timeouts. Valid values:
	//
	// - **0**: Redirects the request to the primary node.
	//
	// - **2**: Downgrades the request. If a global consistency read times out, the system automatically downgrades the query to a non-consistent read, and the client does not receive an error.
	//
	// example:
	//
	// 0
	PolarSccTimeoutAction *string `json:"PolarSccTimeoutAction,omitempty" xml:"PolarSccTimeoutAction,omitempty"`
	// The timeout period for global consistency reads.
	//
	// example:
	//
	// 100
	PolarSccWaitTimeout *string `json:"PolarSccWaitTimeout,omitempty" xml:"PolarSccWaitTimeout,omitempty"`
	// The port.
	//
	// example:
	//
	// 1521
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol type.
	//
	// example:
	//
	// http
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The read/write mode. Valid values:
	//
	// - **ReadWrite**: read and write (automatic read/write splitting).
	//
	// - **ReadOnly**: read-only.
	//
	// example:
	//
	// ReadOnly
	ReadWriteMode *string `json:"ReadWriteMode,omitempty" xml:"ReadWriteMode,omitempty"`
	// Indicates whether global consistency (high-performance mode) is enabled for the node. Valid values:
	//
	// - **on**: enabled
	//
	// - **off**: disabled
	//
	// example:
	//
	// on
	SccMode *string `json:"SccMode,omitempty" xml:"SccMode,omitempty"`
	// The service name.
	//
	// example:
	//
	// test-name
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-***************
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
}

func (s DescribeDBClusterEndpointsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterEndpointsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) GetAddressItems() []*DescribeDBClusterEndpointsResponseBodyItemsAddressItems {
	return s.AddressItems
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) GetAutoAddNewNodes() *string {
	return s.AutoAddNewNodes
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) GetDBEndpointDescription() *string {
	return s.DBEndpointDescription
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) GetDBEndpointId() *string {
	return s.DBEndpointId
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) GetEndpointConfig() *string {
	return s.EndpointConfig
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) GetEndpointType() *string {
	return s.EndpointType
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) GetNetType() *string {
	return s.NetType
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) GetNodeWithRoles() *string {
	return s.NodeWithRoles
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) GetNodes() *string {
	return s.Nodes
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) GetPolarSccTimeoutAction() *string {
	return s.PolarSccTimeoutAction
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) GetPolarSccWaitTimeout() *string {
	return s.PolarSccWaitTimeout
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) GetPort() *string {
	return s.Port
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) GetReadWriteMode() *string {
	return s.ReadWriteMode
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) GetSccMode() *string {
	return s.SccMode
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) GetServiceName() *string {
	return s.ServiceName
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) SetAddressItems(v []*DescribeDBClusterEndpointsResponseBodyItemsAddressItems) *DescribeDBClusterEndpointsResponseBodyItems {
	s.AddressItems = v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) SetAutoAddNewNodes(v string) *DescribeDBClusterEndpointsResponseBodyItems {
	s.AutoAddNewNodes = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) SetConnectionString(v string) *DescribeDBClusterEndpointsResponseBodyItems {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) SetDBClusterId(v string) *DescribeDBClusterEndpointsResponseBodyItems {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) SetDBEndpointDescription(v string) *DescribeDBClusterEndpointsResponseBodyItems {
	s.DBEndpointDescription = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) SetDBEndpointId(v string) *DescribeDBClusterEndpointsResponseBodyItems {
	s.DBEndpointId = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) SetEndpointConfig(v string) *DescribeDBClusterEndpointsResponseBodyItems {
	s.EndpointConfig = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) SetEndpointType(v string) *DescribeDBClusterEndpointsResponseBodyItems {
	s.EndpointType = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) SetNetType(v string) *DescribeDBClusterEndpointsResponseBodyItems {
	s.NetType = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) SetNodeWithRoles(v string) *DescribeDBClusterEndpointsResponseBodyItems {
	s.NodeWithRoles = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) SetNodes(v string) *DescribeDBClusterEndpointsResponseBodyItems {
	s.Nodes = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) SetPolarSccTimeoutAction(v string) *DescribeDBClusterEndpointsResponseBodyItems {
	s.PolarSccTimeoutAction = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) SetPolarSccWaitTimeout(v string) *DescribeDBClusterEndpointsResponseBodyItems {
	s.PolarSccWaitTimeout = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) SetPort(v string) *DescribeDBClusterEndpointsResponseBodyItems {
	s.Port = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) SetProtocol(v string) *DescribeDBClusterEndpointsResponseBodyItems {
	s.Protocol = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) SetReadWriteMode(v string) *DescribeDBClusterEndpointsResponseBodyItems {
	s.ReadWriteMode = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) SetSccMode(v string) *DescribeDBClusterEndpointsResponseBodyItems {
	s.SccMode = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) SetServiceName(v string) *DescribeDBClusterEndpointsResponseBodyItems {
	s.ServiceName = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) SetVPCId(v string) *DescribeDBClusterEndpointsResponseBodyItems {
	s.VPCId = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) Validate() error {
	if s.AddressItems != nil {
		for _, item := range s.AddressItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBClusterEndpointsResponseBodyItemsAddressItems struct {
	// The connection string.
	//
	// example:
	//
	// ********.rwlb.polardb-pg-public.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// Indicates whether the endpoint is a dashboard endpoint for a PolarDB Search node.
	//
	// - **True**: Yes
	//
	// - **False**: No
	//
	// example:
	//
	// True
	DashboardUsed *bool `json:"DashboardUsed,omitempty" xml:"DashboardUsed,omitempty"`
	// The IP address.
	//
	// example:
	//
	// 192.***.***.***
	IPAddress *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	// The network type. Valid values:
	//
	// - **Public**: public network
	//
	// - **Private**: private network
	//
	// <props="china">
	//
	// - **Inner**: classic network
	//
	//
	//
	// <props="china">
	//
	// Only PolarDB for MySQL clusters support the classic network type.
	//
	// example:
	//
	// Private
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The port.
	//
	// example:
	//
	// 1521
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The connection string for the private domain name.
	//
	// example:
	//
	// ***.***.**.com
	PrivateZoneConnectionString *string `json:"PrivateZoneConnectionString,omitempty" xml:"PrivateZoneConnectionString,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-***************
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-************
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC instance ID.
	//
	// > This parameter is returned only for PolarDB for MySQL clusters.
	//
	// example:
	//
	// pe-*************
	VpcInstanceId *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
}

func (s DescribeDBClusterEndpointsResponseBodyItemsAddressItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterEndpointsResponseBodyItemsAddressItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterEndpointsResponseBodyItemsAddressItems) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeDBClusterEndpointsResponseBodyItemsAddressItems) GetDashboardUsed() *bool {
	return s.DashboardUsed
}

func (s *DescribeDBClusterEndpointsResponseBodyItemsAddressItems) GetIPAddress() *string {
	return s.IPAddress
}

func (s *DescribeDBClusterEndpointsResponseBodyItemsAddressItems) GetNetType() *string {
	return s.NetType
}

func (s *DescribeDBClusterEndpointsResponseBodyItemsAddressItems) GetPort() *string {
	return s.Port
}

func (s *DescribeDBClusterEndpointsResponseBodyItemsAddressItems) GetPrivateZoneConnectionString() *string {
	return s.PrivateZoneConnectionString
}

func (s *DescribeDBClusterEndpointsResponseBodyItemsAddressItems) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeDBClusterEndpointsResponseBodyItemsAddressItems) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBClusterEndpointsResponseBodyItemsAddressItems) GetVpcInstanceId() *string {
	return s.VpcInstanceId
}

func (s *DescribeDBClusterEndpointsResponseBodyItemsAddressItems) SetConnectionString(v string) *DescribeDBClusterEndpointsResponseBodyItemsAddressItems {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItemsAddressItems) SetDashboardUsed(v bool) *DescribeDBClusterEndpointsResponseBodyItemsAddressItems {
	s.DashboardUsed = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItemsAddressItems) SetIPAddress(v string) *DescribeDBClusterEndpointsResponseBodyItemsAddressItems {
	s.IPAddress = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItemsAddressItems) SetNetType(v string) *DescribeDBClusterEndpointsResponseBodyItemsAddressItems {
	s.NetType = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItemsAddressItems) SetPort(v string) *DescribeDBClusterEndpointsResponseBodyItemsAddressItems {
	s.Port = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItemsAddressItems) SetPrivateZoneConnectionString(v string) *DescribeDBClusterEndpointsResponseBodyItemsAddressItems {
	s.PrivateZoneConnectionString = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItemsAddressItems) SetVPCId(v string) *DescribeDBClusterEndpointsResponseBodyItemsAddressItems {
	s.VPCId = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItemsAddressItems) SetVSwitchId(v string) *DescribeDBClusterEndpointsResponseBodyItemsAddressItems {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItemsAddressItems) SetVpcInstanceId(v string) *DescribeDBClusterEndpointsResponseBodyItemsAddressItems {
	s.VpcInstanceId = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItemsAddressItems) Validate() error {
	return dara.Validate(s)
}
