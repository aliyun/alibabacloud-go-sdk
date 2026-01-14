// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBasicAcceleratorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerators(v []*ListBasicAcceleratorsResponseBodyAccelerators) *ListBasicAcceleratorsResponseBody
	GetAccelerators() []*ListBasicAcceleratorsResponseBodyAccelerators
	SetPageNumber(v int32) *ListBasicAcceleratorsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListBasicAcceleratorsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListBasicAcceleratorsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListBasicAcceleratorsResponseBody
	GetTotalCount() *int32
}

type ListBasicAcceleratorsResponseBody struct {
	// The information about basic GA instances.
	Accelerators []*ListBasicAcceleratorsResponseBodyAccelerators `json:"Accelerators,omitempty" xml:"Accelerators,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of basic GA instances returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBasicAcceleratorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBasicAcceleratorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBasicAcceleratorsResponseBody) GetAccelerators() []*ListBasicAcceleratorsResponseBodyAccelerators {
	return s.Accelerators
}

func (s *ListBasicAcceleratorsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBasicAcceleratorsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBasicAcceleratorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBasicAcceleratorsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListBasicAcceleratorsResponseBody) SetAccelerators(v []*ListBasicAcceleratorsResponseBodyAccelerators) *ListBasicAcceleratorsResponseBody {
	s.Accelerators = v
	return s
}

func (s *ListBasicAcceleratorsResponseBody) SetPageNumber(v int32) *ListBasicAcceleratorsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBody) SetPageSize(v int32) *ListBasicAcceleratorsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBody) SetRequestId(v string) *ListBasicAcceleratorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBody) SetTotalCount(v int32) *ListBasicAcceleratorsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBody) Validate() error {
	if s.Accelerators != nil {
		for _, item := range s.Accelerators {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListBasicAcceleratorsResponseBodyAccelerators struct {
	// The ID of the basic GA instance.
	//
	// example:
	//
	// ga-bp17frjjh0udz4qz****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The bandwidth billing method.
	//
	// 	- **BandwidthPackage**: billed based on bandwidth plans.
	//
	// 	- **CDT**: billed through Cloud Data Transfer (CDT) and based on data transfer.
	//
	// 	- **CDT95**: billed through CDT and based on the 95th percentile bandwidth. This bandwidth billing method is available only for users that are included in the whitelist.
	//
	// example:
	//
	// BandwidthPackage
	BandwidthBillingType *string `json:"BandwidthBillingType,omitempty" xml:"BandwidthBillingType,omitempty"`
	// Details about the basic bandwidth plan that is associated with the basic GA instance.
	BasicBandwidthPackage *ListBasicAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage `json:"BasicBandwidthPackage,omitempty" xml:"BasicBandwidthPackage,omitempty" type:"Struct"`
	// The ID of the endpoint group that is associated with the basic GA instance.
	//
	// example:
	//
	// epg-bp1dmlohjjz4kqaun****
	BasicEndpointGroupId *string `json:"BasicEndpointGroupId,omitempty" xml:"BasicEndpointGroupId,omitempty"`
	// The ID of the acceleration region where the basic GA instance is deployed.
	//
	// example:
	//
	// ips-bp11ilwqjdkjeg9r7****
	BasicIpSetId *string `json:"BasicIpSetId,omitempty" xml:"BasicIpSetId,omitempty"`
	// The timestamp that indicates when the basic GA instance was created.
	//
	// The time follows the UNIX time format. It is the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1637734547
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether cross-border acceleration is enabled for the GA instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	CrossBorderStatus *bool `json:"CrossBorderStatus,omitempty" xml:"CrossBorderStatus,omitempty"`
	// Details about the cross-region acceleration bandwidth plan that is associated with the GA instance.
	//
	// This parameter is returned only when you call this operation on the International site (alibabacloud.com).
	CrossDomainBandwidthPackage *ListBasicAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage `json:"CrossDomainBandwidthPackage,omitempty" xml:"CrossDomainBandwidthPackage,omitempty" type:"Struct"`
	// The description of the basic GA instance.
	//
	// example:
	//
	// BasicAccelerator
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The timestamp that indicates when the basic GA instance expires.
	//
	// The time follows the UNIX time format. It is the number of seconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1640326547
	ExpiredTime *int64 `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The billing method of the basic GA instance. Only **PREPAY*	- is returned, which indicates the subscription billing method.
	//
	// example:
	//
	// PREPAY
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The name of the basic GA instance.
	//
	// example:
	//
	// BasicAccelerator
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region where the basic GA instance is deployed.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the basic GA instance belongs.
	//
	// example:
	//
	// rg-aekzrnd67gq****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The state of the basic GA instance.
	//
	// 	- **init**: The basic GA instance is being initialized.
	//
	// 	- **active**: The basic GA instance is available.
	//
	// 	- **configuring**: The basic GA instance is being configured.
	//
	// 	- **binding**: The basic GA instance is being associated.
	//
	// 	- **unbinding**: The GA instance is being disassociated.
	//
	// 	- **deleting**: The basic GA instance is being deleted.
	//
	// 	- **finacialLocked**: The basic GA instance is locked due to overdue payments.
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The tags of the basic GA instance.
	Tags []*ListBasicAcceleratorsResponseBodyAcceleratorsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// This parameter is invalid.
	//
	// example:
	//
	// None
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListBasicAcceleratorsResponseBodyAccelerators) String() string {
	return dara.Prettify(s)
}

func (s ListBasicAcceleratorsResponseBodyAccelerators) GoString() string {
	return s.String()
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) GetBandwidthBillingType() *string {
	return s.BandwidthBillingType
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) GetBasicBandwidthPackage() *ListBasicAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage {
	return s.BasicBandwidthPackage
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) GetBasicEndpointGroupId() *string {
	return s.BasicEndpointGroupId
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) GetBasicIpSetId() *string {
	return s.BasicIpSetId
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) GetCrossBorderStatus() *bool {
	return s.CrossBorderStatus
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) GetCrossDomainBandwidthPackage() *ListBasicAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage {
	return s.CrossDomainBandwidthPackage
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) GetDescription() *string {
	return s.Description
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) GetExpiredTime() *int64 {
	return s.ExpiredTime
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) GetName() *string {
	return s.Name
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) GetRegionId() *string {
	return s.RegionId
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) GetState() *string {
	return s.State
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) GetTags() []*ListBasicAcceleratorsResponseBodyAcceleratorsTags {
	return s.Tags
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) GetType() *string {
	return s.Type
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) SetAcceleratorId(v string) *ListBasicAcceleratorsResponseBodyAccelerators {
	s.AcceleratorId = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) SetBandwidthBillingType(v string) *ListBasicAcceleratorsResponseBodyAccelerators {
	s.BandwidthBillingType = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) SetBasicBandwidthPackage(v *ListBasicAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage) *ListBasicAcceleratorsResponseBodyAccelerators {
	s.BasicBandwidthPackage = v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) SetBasicEndpointGroupId(v string) *ListBasicAcceleratorsResponseBodyAccelerators {
	s.BasicEndpointGroupId = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) SetBasicIpSetId(v string) *ListBasicAcceleratorsResponseBodyAccelerators {
	s.BasicIpSetId = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) SetCreateTime(v int64) *ListBasicAcceleratorsResponseBodyAccelerators {
	s.CreateTime = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) SetCrossBorderStatus(v bool) *ListBasicAcceleratorsResponseBodyAccelerators {
	s.CrossBorderStatus = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) SetCrossDomainBandwidthPackage(v *ListBasicAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage) *ListBasicAcceleratorsResponseBodyAccelerators {
	s.CrossDomainBandwidthPackage = v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) SetDescription(v string) *ListBasicAcceleratorsResponseBodyAccelerators {
	s.Description = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) SetExpiredTime(v int64) *ListBasicAcceleratorsResponseBodyAccelerators {
	s.ExpiredTime = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) SetInstanceChargeType(v string) *ListBasicAcceleratorsResponseBodyAccelerators {
	s.InstanceChargeType = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) SetName(v string) *ListBasicAcceleratorsResponseBodyAccelerators {
	s.Name = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) SetRegionId(v string) *ListBasicAcceleratorsResponseBodyAccelerators {
	s.RegionId = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) SetResourceGroupId(v string) *ListBasicAcceleratorsResponseBodyAccelerators {
	s.ResourceGroupId = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) SetState(v string) *ListBasicAcceleratorsResponseBodyAccelerators {
	s.State = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) SetTags(v []*ListBasicAcceleratorsResponseBodyAcceleratorsTags) *ListBasicAcceleratorsResponseBodyAccelerators {
	s.Tags = v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) SetType(v string) *ListBasicAcceleratorsResponseBodyAccelerators {
	s.Type = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAccelerators) Validate() error {
	if s.BasicBandwidthPackage != nil {
		if err := s.BasicBandwidthPackage.Validate(); err != nil {
			return err
		}
	}
	if s.CrossDomainBandwidthPackage != nil {
		if err := s.CrossDomainBandwidthPackage.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListBasicAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage struct {
	// The bandwidth value of the basic bandwidth plan. Unit: Mbit/s.
	//
	// example:
	//
	// 2
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The type of the bandwidth that is provided by the basic bandwidth plan.
	//
	// 	- **Basic**: basic
	//
	// 	- **Enhanced**: enhanced
	//
	// 	- **Advanced**: premium
	//
	// example:
	//
	// Basic
	BandwidthType *string `json:"BandwidthType,omitempty" xml:"BandwidthType,omitempty"`
	// The ID of the basic bandwidth plan.
	//
	// example:
	//
	// gbwp-bp1d8xk8bg139j0fw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListBasicAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage) String() string {
	return dara.Prettify(s)
}

func (s ListBasicAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage) GoString() string {
	return s.String()
}

func (s *ListBasicAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *ListBasicAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage) GetBandwidthType() *string {
	return s.BandwidthType
}

func (s *ListBasicAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListBasicAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage) SetBandwidth(v int32) *ListBasicAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage {
	s.Bandwidth = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage) SetBandwidthType(v string) *ListBasicAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage {
	s.BandwidthType = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage) SetInstanceId(v string) *ListBasicAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage {
	s.InstanceId = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAcceleratorsBasicBandwidthPackage) Validate() error {
	return dara.Validate(s)
}

type ListBasicAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage struct {
	// The bandwidth value of the cross-region acceleration bandwidth plan. Unit: Mbit/s.
	//
	// example:
	//
	// 2
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The ID of the cross-region acceleration bandwidth plan.
	//
	// example:
	//
	// gbwp-bp1d8xk8bg139j0fw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListBasicAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage) String() string {
	return dara.Prettify(s)
}

func (s ListBasicAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage) GoString() string {
	return s.String()
}

func (s *ListBasicAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *ListBasicAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListBasicAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage) SetBandwidth(v int32) *ListBasicAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage {
	s.Bandwidth = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage) SetInstanceId(v string) *ListBasicAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage {
	s.InstanceId = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAcceleratorsCrossDomainBandwidthPackage) Validate() error {
	return dara.Validate(s)
}

type ListBasicAcceleratorsResponseBodyAcceleratorsTags struct {
	// The tag key of the basic GA instance.
	//
	// example:
	//
	// Keytest
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the basic GA instance.
	//
	// example:
	//
	// Valuetest
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListBasicAcceleratorsResponseBodyAcceleratorsTags) String() string {
	return dara.Prettify(s)
}

func (s ListBasicAcceleratorsResponseBodyAcceleratorsTags) GoString() string {
	return s.String()
}

func (s *ListBasicAcceleratorsResponseBodyAcceleratorsTags) GetKey() *string {
	return s.Key
}

func (s *ListBasicAcceleratorsResponseBodyAcceleratorsTags) GetValue() *string {
	return s.Value
}

func (s *ListBasicAcceleratorsResponseBodyAcceleratorsTags) SetKey(v string) *ListBasicAcceleratorsResponseBodyAcceleratorsTags {
	s.Key = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAcceleratorsTags) SetValue(v string) *ListBasicAcceleratorsResponseBodyAcceleratorsTags {
	s.Value = &v
	return s
}

func (s *ListBasicAcceleratorsResponseBodyAcceleratorsTags) Validate() error {
	return dara.Validate(s)
}
