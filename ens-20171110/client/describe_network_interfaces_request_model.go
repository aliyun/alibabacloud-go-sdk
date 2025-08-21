// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkInterfacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnsRegionId(v string) *DescribeNetworkInterfacesRequest
	GetEnsRegionId() *string
	SetEnsRegionIds(v []*string) *DescribeNetworkInterfacesRequest
	GetEnsRegionIds() []*string
	SetInstanceId(v string) *DescribeNetworkInterfacesRequest
	GetInstanceId() *string
	SetIpv6Address(v []*string) *DescribeNetworkInterfacesRequest
	GetIpv6Address() []*string
	SetNetworkId(v string) *DescribeNetworkInterfacesRequest
	GetNetworkId() *string
	SetNetworkInterfaceId(v string) *DescribeNetworkInterfacesRequest
	GetNetworkInterfaceId() *string
	SetNetworkInterfaceIds(v []*string) *DescribeNetworkInterfacesRequest
	GetNetworkInterfaceIds() []*string
	SetNetworkInterfaceName(v string) *DescribeNetworkInterfacesRequest
	GetNetworkInterfaceName() *string
	SetPageNumber(v string) *DescribeNetworkInterfacesRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeNetworkInterfacesRequest
	GetPageSize() *string
	SetPrimaryIpAddress(v string) *DescribeNetworkInterfacesRequest
	GetPrimaryIpAddress() *string
	SetSecurityGroupId(v string) *DescribeNetworkInterfacesRequest
	GetSecurityGroupId() *string
	SetStatus(v string) *DescribeNetworkInterfacesRequest
	GetStatus() *string
	SetType(v string) *DescribeNetworkInterfacesRequest
	GetType() *string
	SetVSwitchId(v string) *DescribeNetworkInterfacesRequest
	GetVSwitchId() *string
}

type DescribeNetworkInterfacesRequest struct {
	// The region ID of the instance.
	//
	// example:
	//
	// cn-tianjin-cmcc
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The IDs of edge nodes. N indicates the number of edge node IDs that you can specify at the same time. Valid values of N: 1 to 100.
	EnsRegionIds []*string `json:"EnsRegionIds,omitempty" xml:"EnsRegionIds,omitempty" type:"Repeated"`
	// The ID of the instance.
	//
	// example:
	//
	// i-5t7z99n32gplriv
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// IPv6 addresses N of the ENI. You can specify multiple IPv6 addresses. Valid values of N: 1 to 100.
	Ipv6Address []*string `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty" type:"Repeated"`
	// The ID of the network.
	//
	// example:
	//
	// n-2zeuphj08tt7q3brd****
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// The ID of the ENI.
	//
	// example:
	//
	// eni-58z57orgmt6d1****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The IDs of the elastic network interfaces (ENIs). N indicates the number of ENI IDs that you can specify at the same time. Valid values of N: 1 to 100.
	NetworkInterfaceIds []*string `json:"NetworkInterfaceIds,omitempty" xml:"NetworkInterfaceIds,omitempty" type:"Repeated"`
	// The name of the ENI.
	//
	// example:
	//
	// test-01
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitempty" xml:"NetworkInterfaceName,omitempty"`
	// The page number of the returned page. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: 100. Default value: 50.
	//
	// example:
	//
	// 50
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The primary IP address of the ENI.
	//
	// example:
	//
	// ***
	PrimaryIpAddress *string `json:"PrimaryIpAddress,omitempty" xml:"PrimaryIpAddress,omitempty"`
	// The ID of the security group to which the secondary ENI belongs. To query the details of secondary ENIs based on the ID of a security group, specify this parameter.
	//
	// example:
	//
	// sg-5p1fg655nh68xyz9i***
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The status of the ENI. Valid values:
	//
	// 	- Available: The ENI is available.
	//
	// 	- Attaching: The ENI is being attached to an instance.
	//
	// 	- InUse: The ENI is attached to an instance.
	//
	// 	- Detaching: The ENI is being detached from an instance.
	//
	// 	- Deleting: The ENI is being deleted.
	//
	// This parameter is empty by default, which indicates that ENIs in all states are queried.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the ENI. Valid values:
	//
	// 	- Primary: the primary ENI.
	//
	// 	- Secondary: the secondary ENI.
	//
	// This parameter is empty by default, which indicates that both primary and secondary ENIs are queried.
	//
	// example:
	//
	// Secondary
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-12345
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeNetworkInterfacesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfacesRequest) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacesRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeNetworkInterfacesRequest) GetEnsRegionIds() []*string {
	return s.EnsRegionIds
}

func (s *DescribeNetworkInterfacesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeNetworkInterfacesRequest) GetIpv6Address() []*string {
	return s.Ipv6Address
}

func (s *DescribeNetworkInterfacesRequest) GetNetworkId() *string {
	return s.NetworkId
}

func (s *DescribeNetworkInterfacesRequest) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *DescribeNetworkInterfacesRequest) GetNetworkInterfaceIds() []*string {
	return s.NetworkInterfaceIds
}

func (s *DescribeNetworkInterfacesRequest) GetNetworkInterfaceName() *string {
	return s.NetworkInterfaceName
}

func (s *DescribeNetworkInterfacesRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeNetworkInterfacesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeNetworkInterfacesRequest) GetPrimaryIpAddress() *string {
	return s.PrimaryIpAddress
}

func (s *DescribeNetworkInterfacesRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeNetworkInterfacesRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeNetworkInterfacesRequest) GetType() *string {
	return s.Type
}

func (s *DescribeNetworkInterfacesRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeNetworkInterfacesRequest) SetEnsRegionId(v string) *DescribeNetworkInterfacesRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetEnsRegionIds(v []*string) *DescribeNetworkInterfacesRequest {
	s.EnsRegionIds = v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetInstanceId(v string) *DescribeNetworkInterfacesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetIpv6Address(v []*string) *DescribeNetworkInterfacesRequest {
	s.Ipv6Address = v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetNetworkId(v string) *DescribeNetworkInterfacesRequest {
	s.NetworkId = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetNetworkInterfaceId(v string) *DescribeNetworkInterfacesRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetNetworkInterfaceIds(v []*string) *DescribeNetworkInterfacesRequest {
	s.NetworkInterfaceIds = v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetNetworkInterfaceName(v string) *DescribeNetworkInterfacesRequest {
	s.NetworkInterfaceName = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetPageNumber(v string) *DescribeNetworkInterfacesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetPageSize(v string) *DescribeNetworkInterfacesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetPrimaryIpAddress(v string) *DescribeNetworkInterfacesRequest {
	s.PrimaryIpAddress = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetSecurityGroupId(v string) *DescribeNetworkInterfacesRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetStatus(v string) *DescribeNetworkInterfacesRequest {
	s.Status = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetType(v string) *DescribeNetworkInterfacesRequest {
	s.Type = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) SetVSwitchId(v string) *DescribeNetworkInterfacesRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeNetworkInterfacesRequest) Validate() error {
	return dara.Validate(s)
}
