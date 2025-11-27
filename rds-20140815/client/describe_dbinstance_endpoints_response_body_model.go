// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceEndpointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeDBInstanceEndpointsResponseBodyData) *DescribeDBInstanceEndpointsResponseBody
	GetData() *DescribeDBInstanceEndpointsResponseBodyData
	SetRequestId(v string) *DescribeDBInstanceEndpointsResponseBody
	GetRequestId() *string
}

type DescribeDBInstanceEndpointsResponseBody struct {
	// The data returned.
	Data *DescribeDBInstanceEndpointsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 777C4593-8053-427B-****105593277CAB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceEndpointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceEndpointsResponseBody) GetData() *DescribeDBInstanceEndpointsResponseBodyData {
	return s.Data
}

func (s *DescribeDBInstanceEndpointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceEndpointsResponseBody) SetData(v *DescribeDBInstanceEndpointsResponseBodyData) *DescribeDBInstanceEndpointsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDBInstanceEndpointsResponseBody) SetRequestId(v string) *DescribeDBInstanceEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceEndpointsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstanceEndpointsResponseBodyData struct {
	// The information of the endpoints of the instance.
	DBInstanceEndpoints *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpoints `json:"DBInstanceEndpoints,omitempty" xml:"DBInstanceEndpoints,omitempty" type:"Struct"`
	// The name of the instance.
	//
	// example:
	//
	// rm-u****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The version of the IP protocol. Valid values:
	//
	// 	- **ipv4**
	//
	// 	- **ipv6**
	//
	// example:
	//
	// ipv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
}

func (s DescribeDBInstanceEndpointsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceEndpointsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceEndpointsResponseBodyData) GetDBInstanceEndpoints() *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpoints {
	return s.DBInstanceEndpoints
}

func (s *DescribeDBInstanceEndpointsResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeDBInstanceEndpointsResponseBodyData) GetIpVersion() *string {
	return s.IpVersion
}

func (s *DescribeDBInstanceEndpointsResponseBodyData) SetDBInstanceEndpoints(v *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpoints) *DescribeDBInstanceEndpointsResponseBodyData {
	s.DBInstanceEndpoints = v
	return s
}

func (s *DescribeDBInstanceEndpointsResponseBodyData) SetDBInstanceName(v string) *DescribeDBInstanceEndpointsResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDBInstanceEndpointsResponseBodyData) SetIpVersion(v string) *DescribeDBInstanceEndpointsResponseBodyData {
	s.IpVersion = &v
	return s
}

func (s *DescribeDBInstanceEndpointsResponseBodyData) Validate() error {
	if s.DBInstanceEndpoints != nil {
		if err := s.DBInstanceEndpoints.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpoints struct {
	DBInstanceEndpoint []*DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpoint `json:"DBInstanceEndpoint,omitempty" xml:"DBInstanceEndpoint,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpoints) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpoints) GetDBInstanceEndpoint() []*DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpoint {
	return s.DBInstanceEndpoint
}

func (s *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpoints) SetDBInstanceEndpoint(v []*DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpoint) *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpoints {
	s.DBInstanceEndpoint = v
	return s
}

func (s *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpoints) Validate() error {
	if s.DBInstanceEndpoint != nil {
		for _, item := range s.DBInstanceEndpoint {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpoint struct {
	// The information about the endpoint.
	AddressItems *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointAddressItems `json:"AddressItems,omitempty" xml:"AddressItems,omitempty" type:"Struct"`
	// The user-defined description of the endpoint.
	//
	// example:
	//
	// for readonly business
	EndpointDescription *string `json:"EndpointDescription,omitempty" xml:"EndpointDescription,omitempty"`
	// The endpoint ID of the instance.
	//
	// example:
	//
	// ep-****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The type of the endpoint. Valid values:
	//
	// 	- **Primary**: the read/write endpoint of the instance
	//
	// 	- **Readonly**: the read-only endpoint of the instance
	//
	// example:
	//
	// Readonly
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The information about the node that is configured for the endpoint.
	NodeItems *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointNodeItems `json:"NodeItems,omitempty" xml:"NodeItems,omitempty" type:"Struct"`
}

func (s DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpoint) GetAddressItems() *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointAddressItems {
	return s.AddressItems
}

func (s *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpoint) GetEndpointDescription() *string {
	return s.EndpointDescription
}

func (s *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpoint) GetEndpointId() *string {
	return s.EndpointId
}

func (s *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpoint) GetEndpointType() *string {
	return s.EndpointType
}

func (s *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpoint) GetNodeItems() *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointNodeItems {
	return s.NodeItems
}

func (s *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpoint) SetAddressItems(v *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointAddressItems) *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpoint {
	s.AddressItems = v
	return s
}

func (s *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpoint) SetEndpointDescription(v string) *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpoint {
	s.EndpointDescription = &v
	return s
}

func (s *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpoint) SetEndpointId(v string) *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpoint {
	s.EndpointId = &v
	return s
}

func (s *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpoint) SetEndpointType(v string) *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpoint {
	s.EndpointType = &v
	return s
}

func (s *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpoint) SetNodeItems(v *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointNodeItems) *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpoint {
	s.NodeItems = v
	return s
}

func (s *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpoint) Validate() error {
	if s.AddressItems != nil {
		if err := s.AddressItems.Validate(); err != nil {
			return err
		}
	}
	if s.NodeItems != nil {
		if err := s.NodeItems.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointAddressItems struct {
	AddressItem []*DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointAddressItemsAddressItem `json:"AddressItem,omitempty" xml:"AddressItem,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointAddressItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointAddressItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointAddressItems) GetAddressItem() []*DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointAddressItemsAddressItem {
	return s.AddressItem
}

func (s *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointAddressItems) SetAddressItem(v []*DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointAddressItemsAddressItem) *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointAddressItems {
	s.AddressItem = v
	return s
}

func (s *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointAddressItems) Validate() error {
	if s.AddressItem != nil {
		for _, item := range s.AddressItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointAddressItemsAddressItem struct {
	// The endpoints of the instance.
	//
	// example:
	//
	// rm****.mysql.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The IP address.
	//
	// example:
	//
	// 10.71.XX.XX
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The type of the IP address. Valid values:
	//
	// 	- **Public**: Internet
	//
	// 	- **Private**: internal network
	//
	// example:
	//
	// Private
	IpType *string `json:"IpType,omitempty" xml:"IpType,omitempty"`
	// The port number of the endpoint.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp12u14ecz****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-bp17xdic25d****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointAddressItemsAddressItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointAddressItemsAddressItem) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointAddressItemsAddressItem) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointAddressItemsAddressItem) GetIpAddress() *string {
	return s.IpAddress
}

func (s *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointAddressItemsAddressItem) GetIpType() *string {
	return s.IpType
}

func (s *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointAddressItemsAddressItem) GetPort() *string {
	return s.Port
}

func (s *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointAddressItemsAddressItem) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointAddressItemsAddressItem) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointAddressItemsAddressItem) SetConnectionString(v string) *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointAddressItemsAddressItem {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointAddressItemsAddressItem) SetIpAddress(v string) *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointAddressItemsAddressItem {
	s.IpAddress = &v
	return s
}

func (s *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointAddressItemsAddressItem) SetIpType(v string) *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointAddressItemsAddressItem {
	s.IpType = &v
	return s
}

func (s *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointAddressItemsAddressItem) SetPort(v string) *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointAddressItemsAddressItem {
	s.Port = &v
	return s
}

func (s *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointAddressItemsAddressItem) SetVSwitchId(v string) *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointAddressItemsAddressItem {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointAddressItemsAddressItem) SetVpcId(v string) *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointAddressItemsAddressItem {
	s.VpcId = &v
	return s
}

func (s *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointAddressItemsAddressItem) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointNodeItems struct {
	NodeItem []*DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointNodeItemsNodeItem `json:"NodeItem,omitempty" xml:"NodeItem,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointNodeItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointNodeItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointNodeItems) GetNodeItem() []*DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointNodeItemsNodeItem {
	return s.NodeItem
}

func (s *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointNodeItems) SetNodeItem(v []*DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointNodeItemsNodeItem) *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointNodeItems {
	s.NodeItem = v
	return s
}

func (s *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointNodeItems) Validate() error {
	if s.NodeItem != nil {
		for _, item := range s.NodeItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointNodeItemsNodeItem struct {
	// The instance ID.
	//
	// example:
	//
	// rm-u****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The node ID.
	//
	// example:
	//
	// rn-****13p6tum4289h
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The weight of the node. Read requests are distributed based on the weight.
	//
	// Valid values: 0 to 100.
	//
	// example:
	//
	// 50
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointNodeItemsNodeItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointNodeItemsNodeItem) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointNodeItemsNodeItem) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointNodeItemsNodeItem) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointNodeItemsNodeItem) GetWeight() *int32 {
	return s.Weight
}

func (s *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointNodeItemsNodeItem) SetDBInstanceId(v string) *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointNodeItemsNodeItem {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointNodeItemsNodeItem) SetNodeId(v string) *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointNodeItemsNodeItem {
	s.NodeId = &v
	return s
}

func (s *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointNodeItemsNodeItem) SetWeight(v int32) *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointNodeItemsNodeItem {
	s.Weight = &v
	return s
}

func (s *DescribeDBInstanceEndpointsResponseBodyDataDBInstanceEndpointsDBInstanceEndpointNodeItemsNodeItem) Validate() error {
	return dara.Validate(s)
}
