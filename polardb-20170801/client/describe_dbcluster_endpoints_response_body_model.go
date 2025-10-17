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
	// The information about the endpoints.
	Items []*DescribeDBClusterEndpointsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The ID of the request.
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
	// The details of the endpoint.
	AddressItems []*DescribeDBClusterEndpointsResponseBodyItemsAddressItems `json:"AddressItems,omitempty" xml:"AddressItems,omitempty" type:"Repeated"`
	// Indicates whether new nodes are automatically associated with the default cluster endpoint. Valid values:
	//
	// 	- **Enable**
	//
	// 	- **Disable**
	//
	// example:
	//
	// Enable
	AutoAddNewNodes *string `json:"AutoAddNewNodes,omitempty" xml:"AutoAddNewNodes,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// pc-bp1s826a1up******
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the endpoint.
	//
	// example:
	//
	// test
	DBEndpointDescription *string `json:"DBEndpointDescription,omitempty" xml:"DBEndpointDescription,omitempty"`
	// The ID of the endpoint.
	//
	// example:
	//
	// pe-*************
	DBEndpointId *string `json:"DBEndpointId,omitempty" xml:"DBEndpointId,omitempty"`
	// The advanced configurations of the endpoint.
	//
	// 	- **DistributedTransaction**: indicates whether transaction splitting is enabled. Valid values:
	//
	//     	- **on**
	//
	//     	- **off**
	//
	// 	- **ConsistLevel**: the consistency level of sessions. Valid values:
	//
	//     	- **0**: eventual consistency.
	//
	//     	- **1**: session consistency.
	//
	//     	- **2**: global consistency.
	//
	// 	- **LoadBalanceStrategy**: the load balancing policy that automatically schedules loads. Only **load*	- may be returned.
	//
	// 	- **MasterAcceptReads**: indicates whether the primary node processes read requests. Valid values:
	//
	//     	- **on**
	//
	//     	- **off**
	//
	// example:
	//
	// {\\"DistributedTransaction\\":\\"off\\",\\"ConsistLevel\\":\\"0\\",\\"LoadBalanceStrategy\\":\\"load\\",\\"MasterAcceptReads\\":\\"on\\"}
	EndpointConfig *string `json:"EndpointConfig,omitempty" xml:"EndpointConfig,omitempty"`
	// The type of the endpoint. Valid values:
	//
	// 	- **Cluster**: the default endpoint.
	//
	// 	- **Primary**: the primary endpoint.
	//
	// 	- **Custom**: a custom cluster endpoint.
	//
	// example:
	//
	// Custom
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The role name of each node in the endpoint. The role name of the primary node is **Writer**. Multiple read-only nodes can be associated with an endpoint. Therefore, the role name of each read-only node is suffixed with a number. For example, you can use **Reader1*	- and **Reader2*	- as the role names.
	//
	// >  This parameter is valid only for PolarDB for PostgreSQL clusters and PolarDB for PostgreSQL (Compatible with Oracle)) clusters.
	//
	// example:
	//
	// Reader1
	NodeWithRoles *string `json:"NodeWithRoles,omitempty" xml:"NodeWithRoles,omitempty"`
	// The nodes in the endpoint.
	//
	// example:
	//
	// pi-***************,pi-***************
	Nodes *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
	// The global consistency timeout policy. Valid values:
	//
	// 	- **0**: sends the request to the primary node.
	//
	// 	- **2**: downgrades the consistency level of a query to inconsistent read when a global consistent read in the query times out. No error message is returned to the client.
	//
	// example:
	//
	// 0
	PolarSccTimeoutAction *string `json:"PolarSccTimeoutAction,omitempty" xml:"PolarSccTimeoutAction,omitempty"`
	// Global consistency timeout.
	//
	// example:
	//
	// 100
	PolarSccWaitTimeout *string `json:"PolarSccWaitTimeout,omitempty" xml:"PolarSccWaitTimeout,omitempty"`
	// The read/write mode. Valid values:
	//
	// 	- **ReadWrite**: handles read and write requests. Automatic read/write splitting is enabled.
	//
	// 	- **ReadOnly**: handles read-only requests.
	//
	// example:
	//
	// ReadOnly
	ReadWriteMode *string `json:"ReadWriteMode,omitempty" xml:"ReadWriteMode,omitempty"`
	// Indicates whether the global consistency (high-performance mode) feature is enabled for the node. Valid values:
	//
	// 	- **on**: enabled.
	//
	// 	- **off**: disabled
	//
	// example:
	//
	// on
	SccMode *string `json:"SccMode,omitempty" xml:"SccMode,omitempty"`
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

func (s *DescribeDBClusterEndpointsResponseBodyItems) GetReadWriteMode() *string {
	return s.ReadWriteMode
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) GetSccMode() *string {
	return s.SccMode
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) SetAddressItems(v []*DescribeDBClusterEndpointsResponseBodyItemsAddressItems) *DescribeDBClusterEndpointsResponseBodyItems {
	s.AddressItems = v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) SetAutoAddNewNodes(v string) *DescribeDBClusterEndpointsResponseBodyItems {
	s.AutoAddNewNodes = &v
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

func (s *DescribeDBClusterEndpointsResponseBodyItems) SetReadWriteMode(v string) *DescribeDBClusterEndpointsResponseBodyItems {
	s.ReadWriteMode = &v
	return s
}

func (s *DescribeDBClusterEndpointsResponseBodyItems) SetSccMode(v string) *DescribeDBClusterEndpointsResponseBodyItems {
	s.SccMode = &v
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
	// The endpoint.
	//
	// example:
	//
	// ********.rwlb.polardb-pg-public.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// Whether it is the dashboard endpoint of the PolarDB search node.
	//
	// 	- Ture
	//
	// 	- False
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
	// The network type of the endpoint. Valid values:
	//
	// 	- **Public**
	//
	// 	- **Private**
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
	// The private domain name that is bound to the endpoint.
	//
	// example:
	//
	// ***.***.**.com
	PrivateZoneConnectionString *string `json:"PrivateZoneConnectionString,omitempty" xml:"PrivateZoneConnectionString,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-***************
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-************
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC) instance.
	//
	// > This parameter is returned for only PolarDB for MySQL clusters.
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
