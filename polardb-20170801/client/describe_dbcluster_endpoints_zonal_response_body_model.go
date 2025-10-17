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
	Items []*DescribeDBClusterEndpointsZonalResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
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
	AddressItems []*DescribeDBClusterEndpointsZonalResponseBodyItemsAddressItems `json:"AddressItems,omitempty" xml:"AddressItems,omitempty" type:"Repeated"`
	// example:
	//
	// Enable
	AutoAddNewNodes *string `json:"AutoAddNewNodes,omitempty" xml:"AutoAddNewNodes,omitempty"`
	// example:
	//
	// pc-bp1s826a1up******
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// test
	DBEndpointDescription *string `json:"DBEndpointDescription,omitempty" xml:"DBEndpointDescription,omitempty"`
	// example:
	//
	// pe-*************
	DBEndpointId *string `json:"DBEndpointId,omitempty" xml:"DBEndpointId,omitempty"`
	// example:
	//
	// {\\"DistributedTransaction\\":\\"off\\",\\"ConsistLevel\\":\\"0\\",\\"LoadBalanceStrategy\\":\\"load\\",\\"MasterAcceptReads\\":\\"on\\"}
	EndpointConfig *string `json:"EndpointConfig,omitempty" xml:"EndpointConfig,omitempty"`
	// example:
	//
	// Custom
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// example:
	//
	// Reader1
	NodeWithRoles *string `json:"NodeWithRoles,omitempty" xml:"NodeWithRoles,omitempty"`
	// example:
	//
	// pi-***************,pi-***************
	Nodes *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
	// example:
	//
	// 0
	PolarSccTimeoutAction *string `json:"PolarSccTimeoutAction,omitempty" xml:"PolarSccTimeoutAction,omitempty"`
	// example:
	//
	// 100
	PolarSccWaitTimeout *string `json:"PolarSccWaitTimeout,omitempty" xml:"PolarSccWaitTimeout,omitempty"`
	// example:
	//
	// ReadOnly
	ReadWriteMode *string `json:"ReadWriteMode,omitempty" xml:"ReadWriteMode,omitempty"`
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
	// example:
	//
	// ********.rwlb.polardb-pg-public.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// example:
	//
	// True
	DashboardUsed *bool `json:"DashboardUsed,omitempty" xml:"DashboardUsed,omitempty"`
	// example:
	//
	// 192.***.***.***
	IPAddress *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty"`
	// example:
	//
	// Private
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// example:
	//
	// 1521
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// ***.***.**.com
	PrivateZoneConnectionString *string `json:"PrivateZoneConnectionString,omitempty" xml:"PrivateZoneConnectionString,omitempty"`
	// example:
	//
	// vpc-***************
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// example:
	//
	// vsw-************
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
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
