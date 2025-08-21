// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHaVipsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHaVips(v []*DescribeHaVipsResponseBodyHaVips) *DescribeHaVipsResponseBody
	GetHaVips() []*DescribeHaVipsResponseBodyHaVips
	SetPageNumber(v string) *DescribeHaVipsResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *DescribeHaVipsResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeHaVipsResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeHaVipsResponseBody
	GetTotalCount() *string
}

type DescribeHaVipsResponseBody struct {
	// Details of the HAVIPs.
	HaVips []*DescribeHaVipsResponseBodyHaVips `json:"HaVips,omitempty" xml:"HaVips,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AAE90880-4970-4D81-A534-A6C0F3631F74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 49
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeHaVipsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHaVipsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHaVipsResponseBody) GetHaVips() []*DescribeHaVipsResponseBodyHaVips {
	return s.HaVips
}

func (s *DescribeHaVipsResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeHaVipsResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeHaVipsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHaVipsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeHaVipsResponseBody) SetHaVips(v []*DescribeHaVipsResponseBodyHaVips) *DescribeHaVipsResponseBody {
	s.HaVips = v
	return s
}

func (s *DescribeHaVipsResponseBody) SetPageNumber(v string) *DescribeHaVipsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeHaVipsResponseBody) SetPageSize(v string) *DescribeHaVipsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeHaVipsResponseBody) SetRequestId(v string) *DescribeHaVipsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHaVipsResponseBody) SetTotalCount(v string) *DescribeHaVipsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeHaVipsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeHaVipsResponseBodyHaVips struct {
	// The elastic IP addresses (EIPs) that are associated with the HAVIP.
	AssociatedEipAddresses []*DescribeHaVipsResponseBodyHaVipsAssociatedEipAddresses `json:"AssociatedEipAddresses,omitempty" xml:"AssociatedEipAddresses,omitempty" type:"Repeated"`
	// The information about instances that are associated with the HAVIP.
	AssociatedInstances []*DescribeHaVipsResponseBodyHaVipsAssociatedInstances `json:"AssociatedInstances,omitempty" xml:"AssociatedInstances,omitempty" type:"Repeated"`
	// The time when the HAVIP was created.
	//
	// example:
	//
	// 2023-03-29T11:17:38Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the HAVIP.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-guiyang-14
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The ID of the HAVIP.
	//
	// example:
	//
	// havip-52y28****
	HaVipId *string `json:"HaVipId,omitempty" xml:"HaVipId,omitempty"`
	// The IP address of the HAVIP.
	//
	// example:
	//
	// 192.XX.XX.5
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The name of the HAVIP.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the network.
	//
	// example:
	//
	// n-5wtkyrk****
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// The status of the HAVIP. Valid values:
	//
	// 	- Creating
	//
	// 	- Available
	//
	// 	- InUse
	//
	// 	- Deleting
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-5yc8d****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeHaVipsResponseBodyHaVips) String() string {
	return dara.Prettify(s)
}

func (s DescribeHaVipsResponseBodyHaVips) GoString() string {
	return s.String()
}

func (s *DescribeHaVipsResponseBodyHaVips) GetAssociatedEipAddresses() []*DescribeHaVipsResponseBodyHaVipsAssociatedEipAddresses {
	return s.AssociatedEipAddresses
}

func (s *DescribeHaVipsResponseBodyHaVips) GetAssociatedInstances() []*DescribeHaVipsResponseBodyHaVipsAssociatedInstances {
	return s.AssociatedInstances
}

func (s *DescribeHaVipsResponseBodyHaVips) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeHaVipsResponseBodyHaVips) GetDescription() *string {
	return s.Description
}

func (s *DescribeHaVipsResponseBodyHaVips) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeHaVipsResponseBodyHaVips) GetHaVipId() *string {
	return s.HaVipId
}

func (s *DescribeHaVipsResponseBodyHaVips) GetIpAddress() *string {
	return s.IpAddress
}

func (s *DescribeHaVipsResponseBodyHaVips) GetName() *string {
	return s.Name
}

func (s *DescribeHaVipsResponseBodyHaVips) GetNetworkId() *string {
	return s.NetworkId
}

func (s *DescribeHaVipsResponseBodyHaVips) GetStatus() *string {
	return s.Status
}

func (s *DescribeHaVipsResponseBodyHaVips) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeHaVipsResponseBodyHaVips) SetAssociatedEipAddresses(v []*DescribeHaVipsResponseBodyHaVipsAssociatedEipAddresses) *DescribeHaVipsResponseBodyHaVips {
	s.AssociatedEipAddresses = v
	return s
}

func (s *DescribeHaVipsResponseBodyHaVips) SetAssociatedInstances(v []*DescribeHaVipsResponseBodyHaVipsAssociatedInstances) *DescribeHaVipsResponseBodyHaVips {
	s.AssociatedInstances = v
	return s
}

func (s *DescribeHaVipsResponseBodyHaVips) SetCreationTime(v string) *DescribeHaVipsResponseBodyHaVips {
	s.CreationTime = &v
	return s
}

func (s *DescribeHaVipsResponseBodyHaVips) SetDescription(v string) *DescribeHaVipsResponseBodyHaVips {
	s.Description = &v
	return s
}

func (s *DescribeHaVipsResponseBodyHaVips) SetEnsRegionId(v string) *DescribeHaVipsResponseBodyHaVips {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeHaVipsResponseBodyHaVips) SetHaVipId(v string) *DescribeHaVipsResponseBodyHaVips {
	s.HaVipId = &v
	return s
}

func (s *DescribeHaVipsResponseBodyHaVips) SetIpAddress(v string) *DescribeHaVipsResponseBodyHaVips {
	s.IpAddress = &v
	return s
}

func (s *DescribeHaVipsResponseBodyHaVips) SetName(v string) *DescribeHaVipsResponseBodyHaVips {
	s.Name = &v
	return s
}

func (s *DescribeHaVipsResponseBodyHaVips) SetNetworkId(v string) *DescribeHaVipsResponseBodyHaVips {
	s.NetworkId = &v
	return s
}

func (s *DescribeHaVipsResponseBodyHaVips) SetStatus(v string) *DescribeHaVipsResponseBodyHaVips {
	s.Status = &v
	return s
}

func (s *DescribeHaVipsResponseBodyHaVips) SetVSwitchId(v string) *DescribeHaVipsResponseBodyHaVips {
	s.VSwitchId = &v
	return s
}

func (s *DescribeHaVipsResponseBodyHaVips) Validate() error {
	return dara.Validate(s)
}

type DescribeHaVipsResponseBodyHaVipsAssociatedEipAddresses struct {
	// The EIP.
	//
	// example:
	//
	// 47.XX.XX.40
	Eip *string `json:"Eip,omitempty" xml:"Eip,omitempty"`
	// The ID of the EIP.
	//
	// example:
	//
	// eip-5p1wz****
	EipId *string `json:"EipId,omitempty" xml:"EipId,omitempty"`
}

func (s DescribeHaVipsResponseBodyHaVipsAssociatedEipAddresses) String() string {
	return dara.Prettify(s)
}

func (s DescribeHaVipsResponseBodyHaVipsAssociatedEipAddresses) GoString() string {
	return s.String()
}

func (s *DescribeHaVipsResponseBodyHaVipsAssociatedEipAddresses) GetEip() *string {
	return s.Eip
}

func (s *DescribeHaVipsResponseBodyHaVipsAssociatedEipAddresses) GetEipId() *string {
	return s.EipId
}

func (s *DescribeHaVipsResponseBodyHaVipsAssociatedEipAddresses) SetEip(v string) *DescribeHaVipsResponseBodyHaVipsAssociatedEipAddresses {
	s.Eip = &v
	return s
}

func (s *DescribeHaVipsResponseBodyHaVipsAssociatedEipAddresses) SetEipId(v string) *DescribeHaVipsResponseBodyHaVipsAssociatedEipAddresses {
	s.EipId = &v
	return s
}

func (s *DescribeHaVipsResponseBodyHaVipsAssociatedEipAddresses) Validate() error {
	return dara.Validate(s)
}

type DescribeHaVipsResponseBodyHaVipsAssociatedInstances struct {
	// The time when the instance was created.
	//
	// example:
	//
	// 2023-01-05T07:09:28Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-51p****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the instance that is associated with the HAVIP. Valid values:
	//
	// 	- EnsInstance: ENS instance
	//
	// 	- NetworkInterface: elastic network interface (ENI)
	//
	// example:
	//
	// EnsInstance
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The private IP address of the instance that is associated with the HAVIP. Valid values:
	//
	// example:
	//
	// 192.XX.XX.9
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The association status of the HAVIP. Valid values:
	//
	// 	- Associating
	//
	// 	- InUse
	//
	// 	- Unassociating
	//
	// example:
	//
	// InUse
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeHaVipsResponseBodyHaVipsAssociatedInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeHaVipsResponseBodyHaVipsAssociatedInstances) GoString() string {
	return s.String()
}

func (s *DescribeHaVipsResponseBodyHaVipsAssociatedInstances) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeHaVipsResponseBodyHaVipsAssociatedInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeHaVipsResponseBodyHaVipsAssociatedInstances) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeHaVipsResponseBodyHaVipsAssociatedInstances) GetIpAddress() *string {
	return s.IpAddress
}

func (s *DescribeHaVipsResponseBodyHaVipsAssociatedInstances) GetStatus() *string {
	return s.Status
}

func (s *DescribeHaVipsResponseBodyHaVipsAssociatedInstances) SetCreationTime(v string) *DescribeHaVipsResponseBodyHaVipsAssociatedInstances {
	s.CreationTime = &v
	return s
}

func (s *DescribeHaVipsResponseBodyHaVipsAssociatedInstances) SetInstanceId(v string) *DescribeHaVipsResponseBodyHaVipsAssociatedInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeHaVipsResponseBodyHaVipsAssociatedInstances) SetInstanceType(v string) *DescribeHaVipsResponseBodyHaVipsAssociatedInstances {
	s.InstanceType = &v
	return s
}

func (s *DescribeHaVipsResponseBodyHaVipsAssociatedInstances) SetIpAddress(v string) *DescribeHaVipsResponseBodyHaVipsAssociatedInstances {
	s.IpAddress = &v
	return s
}

func (s *DescribeHaVipsResponseBodyHaVipsAssociatedInstances) SetStatus(v string) *DescribeHaVipsResponseBodyHaVipsAssociatedInstances {
	s.Status = &v
	return s
}

func (s *DescribeHaVipsResponseBodyHaVipsAssociatedInstances) Validate() error {
	return dara.Validate(s)
}
