// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterEndpointsZonalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeDBClusterEndpointsZonalResponseBodyItems) *DescribeDBClusterEndpointsZonalResponseBody
	GetItems() []*DescribeDBClusterEndpointsZonalResponseBodyItems
	SetRequestId(v string) *DescribeDBClusterEndpointsZonalResponseBody
	GetRequestId() *string
}

type DescribeDBClusterEndpointsZonalResponseBody struct {
	// The cluster endpoint details.
	Items []*DescribeDBClusterEndpointsZonalResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 2DC120BF-6EBA-4C63-BE99-B09F9E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBClusterEndpointsZonalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterEndpointsZonalResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterEndpointsZonalResponseBody) GetItems() []*DescribeDBClusterEndpointsZonalResponseBodyItems {
	return s.Items
}

func (s *DescribeDBClusterEndpointsZonalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterEndpointsZonalResponseBody) SetItems(v []*DescribeDBClusterEndpointsZonalResponseBodyItems) *DescribeDBClusterEndpointsZonalResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBClusterEndpointsZonalResponseBody) SetRequestId(v string) *DescribeDBClusterEndpointsZonalResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterEndpointsZonalResponseBody) Validate() error {
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

type DescribeDBClusterEndpointsZonalResponseBodyItems struct {
	// The connection string information.
	AddressItems []*DescribeDBClusterEndpointsZonalResponseBodyItemsAddressItems `json:"AddressItems,omitempty" xml:"AddressItems,omitempty" type:"Repeated"`
	// Specifies whether new nodes are automatically added to the default cluster endpoint. Valid values:
	//
	// - Enable.
	//
	// - Disable.
	//
	// example:
	//
	// Enable
	AutoAddNewNodes *string `json:"AutoAddNewNodes,omitempty" xml:"AutoAddNewNodes,omitempty"`
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
	// The advanced configurations of the cluster endpoint.
	//
	// 	- **DistributedTransaction**: The transaction splitting status. Valid values:
	//
	//     	- **on**: Transaction splitting is enabled.
	//
	//     	- **off**: Transaction splitting is disabled.
	//
	// 	- **ConsistLevel**: The session consistency level. Valid values:
	//
	//     	- **0**: Eventual consistency.
	//
	//     	- **1**: Session consistency.
	//
	//     	- **2**: Global consistency.
	//
	// 	- **LoadBalanceStrategy**: The load balancing policy. Automatic scheduling based on load. The value is fixed as **load**.
	//
	// 	- **MasterAcceptReads**: Specifies whether the primary database accepts reads. Valid values:
	//
	//     	- **on**: The primary database accepts reads.
	//
	//     	- **off**: Offload reads from primary node.
	//
	// example:
	//
	// {\\"DistributedTransaction\\":\\"off\\",\\"ConsistLevel\\":\\"0\\",\\"LoadBalanceStrategy\\":\\"load\\",\\"MasterAcceptReads\\":\\"on\\"}
	EndpointConfig *string `json:"EndpointConfig,omitempty" xml:"EndpointConfig,omitempty"`
	// The cluster endpoint type. Valid values:
	//
	// - Cluster: Default cluster endpoint.
	//
	// - Primary: Primary endpoint.
	//
	// - Custom: Custom cluster endpoint.
	//
	// example:
	//
	// Custom
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The role names of the nodes under the target endpoint. The primary node is named Writer. Because multiple read-only nodes can be added to an endpoint, a number is appended to the role name of each read-only node for distinction, such as Reader1 and Reader2.
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
	// The global consistency timeout policy. Valid values:
	//
	// - 0: Sends the request to the primary node.
	//
	// - 2: Downgrades upon timeout. If a global consistency read times out, the query is automatically downgraded to regular requests, and the client does not receive an error message.
	//
	// example:
	//
	// 0
	PolarSccTimeoutAction *string `json:"PolarSccTimeoutAction,omitempty" xml:"PolarSccTimeoutAction,omitempty"`
	// The global consistency timeout period.
	//
	// example:
	//
	// 100
	PolarSccWaitTimeout *string `json:"PolarSccWaitTimeout,omitempty" xml:"PolarSccWaitTimeout,omitempty"`
	// The read/write mode. Valid values:
	//
	// - ReadWrite: Read and write (automatic read/write splitting).
	//
	// - ReadOnly: Read-only.
	//
	// example:
	//
	// ReadOnly
	ReadWriteMode *string `json:"ReadWriteMode,omitempty" xml:"ReadWriteMode,omitempty"`
	// Specifies whether the global consistency (high-performance mode) feature is enabled for the node. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	SccMode *string `json:"SccMode,omitempty" xml:"SccMode,omitempty"`
}

func (s DescribeDBClusterEndpointsZonalResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterEndpointsZonalResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItems) GetAddressItems() []*DescribeDBClusterEndpointsZonalResponseBodyItemsAddressItems {
	return s.AddressItems
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItems) GetAutoAddNewNodes() *string {
	return s.AutoAddNewNodes
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItems) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItems) GetDBEndpointDescription() *string {
	return s.DBEndpointDescription
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItems) GetDBEndpointId() *string {
	return s.DBEndpointId
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItems) GetEndpointConfig() *string {
	return s.EndpointConfig
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItems) GetEndpointType() *string {
	return s.EndpointType
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItems) GetNodeWithRoles() *string {
	return s.NodeWithRoles
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItems) GetNodes() *string {
	return s.Nodes
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItems) GetPolarSccTimeoutAction() *string {
	return s.PolarSccTimeoutAction
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItems) GetPolarSccWaitTimeout() *string {
	return s.PolarSccWaitTimeout
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItems) GetReadWriteMode() *string {
	return s.ReadWriteMode
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItems) GetSccMode() *string {
	return s.SccMode
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItems) SetAddressItems(v []*DescribeDBClusterEndpointsZonalResponseBodyItemsAddressItems) *DescribeDBClusterEndpointsZonalResponseBodyItems {
	s.AddressItems = v
	return s
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItems) SetAutoAddNewNodes(v string) *DescribeDBClusterEndpointsZonalResponseBodyItems {
	s.AutoAddNewNodes = &v
	return s
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItems) SetDBClusterId(v string) *DescribeDBClusterEndpointsZonalResponseBodyItems {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItems) SetDBEndpointDescription(v string) *DescribeDBClusterEndpointsZonalResponseBodyItems {
	s.DBEndpointDescription = &v
	return s
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItems) SetDBEndpointId(v string) *DescribeDBClusterEndpointsZonalResponseBodyItems {
	s.DBEndpointId = &v
	return s
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItems) SetEndpointConfig(v string) *DescribeDBClusterEndpointsZonalResponseBodyItems {
	s.EndpointConfig = &v
	return s
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItems) SetEndpointType(v string) *DescribeDBClusterEndpointsZonalResponseBodyItems {
	s.EndpointType = &v
	return s
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItems) SetNodeWithRoles(v string) *DescribeDBClusterEndpointsZonalResponseBodyItems {
	s.NodeWithRoles = &v
	return s
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItems) SetNodes(v string) *DescribeDBClusterEndpointsZonalResponseBodyItems {
	s.Nodes = &v
	return s
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItems) SetPolarSccTimeoutAction(v string) *DescribeDBClusterEndpointsZonalResponseBodyItems {
	s.PolarSccTimeoutAction = &v
	return s
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItems) SetPolarSccWaitTimeout(v string) *DescribeDBClusterEndpointsZonalResponseBodyItems {
	s.PolarSccWaitTimeout = &v
	return s
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItems) SetReadWriteMode(v string) *DescribeDBClusterEndpointsZonalResponseBodyItems {
	s.ReadWriteMode = &v
	return s
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItems) SetSccMode(v string) *DescribeDBClusterEndpointsZonalResponseBodyItems {
	s.SccMode = &v
	return s
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItems) Validate() error {
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

type DescribeDBClusterEndpointsZonalResponseBodyItemsAddressItems struct {
	// The connection string.
	//
	// example:
	//
	// ********.rwlb.polardb-pg-public.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// Specifies whether the address is the dashboard address of a PolarDB Search node. Valid values:
	//
	// - Ture: Yes.
	//
	// - False: No.
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
	// - Public: Internet.
	//
	// - Private: Private network.
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
	// The bound private domain name.
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
	// The VPC-connected instance ID.
	//
	// example:
	//
	// pe-*************
	VpcInstanceId *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
}

func (s DescribeDBClusterEndpointsZonalResponseBodyItemsAddressItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterEndpointsZonalResponseBodyItemsAddressItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItemsAddressItems) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItemsAddressItems) GetDashboardUsed() *bool {
	return s.DashboardUsed
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItemsAddressItems) GetIPAddress() *string {
	return s.IPAddress
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItemsAddressItems) GetNetType() *string {
	return s.NetType
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItemsAddressItems) GetPort() *string {
	return s.Port
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItemsAddressItems) GetPrivateZoneConnectionString() *string {
	return s.PrivateZoneConnectionString
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItemsAddressItems) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItemsAddressItems) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItemsAddressItems) GetVpcInstanceId() *string {
	return s.VpcInstanceId
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItemsAddressItems) SetConnectionString(v string) *DescribeDBClusterEndpointsZonalResponseBodyItemsAddressItems {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItemsAddressItems) SetDashboardUsed(v bool) *DescribeDBClusterEndpointsZonalResponseBodyItemsAddressItems {
	s.DashboardUsed = &v
	return s
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItemsAddressItems) SetIPAddress(v string) *DescribeDBClusterEndpointsZonalResponseBodyItemsAddressItems {
	s.IPAddress = &v
	return s
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItemsAddressItems) SetNetType(v string) *DescribeDBClusterEndpointsZonalResponseBodyItemsAddressItems {
	s.NetType = &v
	return s
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItemsAddressItems) SetPort(v string) *DescribeDBClusterEndpointsZonalResponseBodyItemsAddressItems {
	s.Port = &v
	return s
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItemsAddressItems) SetPrivateZoneConnectionString(v string) *DescribeDBClusterEndpointsZonalResponseBodyItemsAddressItems {
	s.PrivateZoneConnectionString = &v
	return s
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItemsAddressItems) SetVPCId(v string) *DescribeDBClusterEndpointsZonalResponseBodyItemsAddressItems {
	s.VPCId = &v
	return s
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItemsAddressItems) SetVSwitchId(v string) *DescribeDBClusterEndpointsZonalResponseBodyItemsAddressItems {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItemsAddressItems) SetVpcInstanceId(v string) *DescribeDBClusterEndpointsZonalResponseBodyItemsAddressItems {
	s.VpcInstanceId = &v
	return s
}

func (s *DescribeDBClusterEndpointsZonalResponseBodyItemsAddressItems) Validate() error {
	return dara.Validate(s)
}
