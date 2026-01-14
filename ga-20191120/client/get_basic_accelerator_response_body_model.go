// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBasicAcceleratorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *GetBasicAcceleratorResponseBody
	GetAcceleratorId() *string
	SetBandwidthBillingType(v string) *GetBasicAcceleratorResponseBody
	GetBandwidthBillingType() *string
	SetBasicBandwidthPackage(v *GetBasicAcceleratorResponseBodyBasicBandwidthPackage) *GetBasicAcceleratorResponseBody
	GetBasicBandwidthPackage() *GetBasicAcceleratorResponseBodyBasicBandwidthPackage
	SetBasicEndpointGroupId(v string) *GetBasicAcceleratorResponseBody
	GetBasicEndpointGroupId() *string
	SetBasicIpSetId(v string) *GetBasicAcceleratorResponseBody
	GetBasicIpSetId() *string
	SetCenId(v string) *GetBasicAcceleratorResponseBody
	GetCenId() *string
	SetCreateTime(v int64) *GetBasicAcceleratorResponseBody
	GetCreateTime() *int64
	SetCrossBorderStatus(v bool) *GetBasicAcceleratorResponseBody
	GetCrossBorderStatus() *bool
	SetCrossDomainBandwidthPackage(v *GetBasicAcceleratorResponseBodyCrossDomainBandwidthPackage) *GetBasicAcceleratorResponseBody
	GetCrossDomainBandwidthPackage() *GetBasicAcceleratorResponseBodyCrossDomainBandwidthPackage
	SetCrossPrivateState(v string) *GetBasicAcceleratorResponseBody
	GetCrossPrivateState() *string
	SetDescription(v string) *GetBasicAcceleratorResponseBody
	GetDescription() *string
	SetExpiredTime(v int64) *GetBasicAcceleratorResponseBody
	GetExpiredTime() *int64
	SetInstanceChargeType(v string) *GetBasicAcceleratorResponseBody
	GetInstanceChargeType() *string
	SetName(v string) *GetBasicAcceleratorResponseBody
	GetName() *string
	SetRegionId(v string) *GetBasicAcceleratorResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetBasicAcceleratorResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetBasicAcceleratorResponseBody
	GetResourceGroupId() *string
	SetState(v string) *GetBasicAcceleratorResponseBody
	GetState() *string
	SetTags(v []*GetBasicAcceleratorResponseBodyTags) *GetBasicAcceleratorResponseBody
	GetTags() []*GetBasicAcceleratorResponseBodyTags
}

type GetBasicAcceleratorResponseBody struct {
	// The ID of the basic GA instance.
	//
	// example:
	//
	// ga-bp17frjjh0udz4qz****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The bandwidth metering method.
	//
	// 	- **BandwidthPackage**: billed based on bandwidth plans.
	//
	// 	- **CDT**: billed by Cloud Data Transfer (CDT) and based on data transfer.
	//
	// 	- **CDT95**: billed by CDT and based on the 95th percentile bandwidth. This bandwidth metering method is available only to users that are included in the whitelist.
	//
	// example:
	//
	// CDT
	BandwidthBillingType *string `json:"BandwidthBillingType,omitempty" xml:"BandwidthBillingType,omitempty"`
	// The details about the basic bandwidth plan that is associated with the basic GA instance.
	BasicBandwidthPackage *GetBasicAcceleratorResponseBodyBasicBandwidthPackage `json:"BasicBandwidthPackage,omitempty" xml:"BasicBandwidthPackage,omitempty" type:"Struct"`
	// The ID of the endpoint group.
	//
	// example:
	//
	// epg-bp1dmlohjjz4kqaun****
	BasicEndpointGroupId *string `json:"BasicEndpointGroupId,omitempty" xml:"BasicEndpointGroupId,omitempty"`
	// The ID of the acceleration region.
	//
	// example:
	//
	// ips-bp11ilwqjdkjeg9r7****
	BasicIpSetId *string `json:"BasicIpSetId,omitempty" xml:"BasicIpSetId,omitempty"`
	// The ID of the Cloud Enterprise Network (CEN) instance to which the basic GA instance is attached.
	//
	// example:
	//
	// cen-hjkduu767hc****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The timestamp that indicates when the basic GA instance is created.
	//
	// example:
	//
	// 1637734547
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether cross-border acceleration is enabled for the basic GA instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	CrossBorderStatus *bool `json:"CrossBorderStatus,omitempty" xml:"CrossBorderStatus,omitempty"`
	// The details about the cross-border acceleration bandwidth plan that is associated with the GA instance.
	//
	// This array is returned only for GA instances that are created on the international site (alibabacloud.com).
	CrossDomainBandwidthPackage *GetBasicAcceleratorResponseBodyCrossDomainBandwidthPackage `json:"CrossDomainBandwidthPackage,omitempty" xml:"CrossDomainBandwidthPackage,omitempty" type:"Struct"`
	// Indicates whether cross-border acceleration is enabled.
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// false
	CrossPrivateState *string `json:"CrossPrivateState,omitempty" xml:"CrossPrivateState,omitempty"`
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
	// The ID of the request.
	//
	// example:
	//
	// F591955F-5CB5-4CCE-A75D-17CF2085CE22
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the basic GA instance belongs.
	//
	// example:
	//
	// rg-aekzrnd67gq****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the basic GA instance.
	//
	// 	- **init**: The GA instance is being initialized.
	//
	// 	- **active**: The GA instance is available.
	//
	// 	- **configuring**: The GA instance is being configured.
	//
	// 	- **binding**: The GA instance is being associated.
	//
	// 	- **unbinding**: The GA instance is being disassociated.
	//
	// 	- **deleting**: The GA instance is being deleted.
	//
	// 	- **finacialLocked**: The GA instance is locked due to overdue payments.
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The tags of the basic GA instance.
	Tags []*GetBasicAcceleratorResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s GetBasicAcceleratorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBasicAcceleratorResponseBody) GoString() string {
	return s.String()
}

func (s *GetBasicAcceleratorResponseBody) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *GetBasicAcceleratorResponseBody) GetBandwidthBillingType() *string {
	return s.BandwidthBillingType
}

func (s *GetBasicAcceleratorResponseBody) GetBasicBandwidthPackage() *GetBasicAcceleratorResponseBodyBasicBandwidthPackage {
	return s.BasicBandwidthPackage
}

func (s *GetBasicAcceleratorResponseBody) GetBasicEndpointGroupId() *string {
	return s.BasicEndpointGroupId
}

func (s *GetBasicAcceleratorResponseBody) GetBasicIpSetId() *string {
	return s.BasicIpSetId
}

func (s *GetBasicAcceleratorResponseBody) GetCenId() *string {
	return s.CenId
}

func (s *GetBasicAcceleratorResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetBasicAcceleratorResponseBody) GetCrossBorderStatus() *bool {
	return s.CrossBorderStatus
}

func (s *GetBasicAcceleratorResponseBody) GetCrossDomainBandwidthPackage() *GetBasicAcceleratorResponseBodyCrossDomainBandwidthPackage {
	return s.CrossDomainBandwidthPackage
}

func (s *GetBasicAcceleratorResponseBody) GetCrossPrivateState() *string {
	return s.CrossPrivateState
}

func (s *GetBasicAcceleratorResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetBasicAcceleratorResponseBody) GetExpiredTime() *int64 {
	return s.ExpiredTime
}

func (s *GetBasicAcceleratorResponseBody) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *GetBasicAcceleratorResponseBody) GetName() *string {
	return s.Name
}

func (s *GetBasicAcceleratorResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetBasicAcceleratorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBasicAcceleratorResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetBasicAcceleratorResponseBody) GetState() *string {
	return s.State
}

func (s *GetBasicAcceleratorResponseBody) GetTags() []*GetBasicAcceleratorResponseBodyTags {
	return s.Tags
}

func (s *GetBasicAcceleratorResponseBody) SetAcceleratorId(v string) *GetBasicAcceleratorResponseBody {
	s.AcceleratorId = &v
	return s
}

func (s *GetBasicAcceleratorResponseBody) SetBandwidthBillingType(v string) *GetBasicAcceleratorResponseBody {
	s.BandwidthBillingType = &v
	return s
}

func (s *GetBasicAcceleratorResponseBody) SetBasicBandwidthPackage(v *GetBasicAcceleratorResponseBodyBasicBandwidthPackage) *GetBasicAcceleratorResponseBody {
	s.BasicBandwidthPackage = v
	return s
}

func (s *GetBasicAcceleratorResponseBody) SetBasicEndpointGroupId(v string) *GetBasicAcceleratorResponseBody {
	s.BasicEndpointGroupId = &v
	return s
}

func (s *GetBasicAcceleratorResponseBody) SetBasicIpSetId(v string) *GetBasicAcceleratorResponseBody {
	s.BasicIpSetId = &v
	return s
}

func (s *GetBasicAcceleratorResponseBody) SetCenId(v string) *GetBasicAcceleratorResponseBody {
	s.CenId = &v
	return s
}

func (s *GetBasicAcceleratorResponseBody) SetCreateTime(v int64) *GetBasicAcceleratorResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetBasicAcceleratorResponseBody) SetCrossBorderStatus(v bool) *GetBasicAcceleratorResponseBody {
	s.CrossBorderStatus = &v
	return s
}

func (s *GetBasicAcceleratorResponseBody) SetCrossDomainBandwidthPackage(v *GetBasicAcceleratorResponseBodyCrossDomainBandwidthPackage) *GetBasicAcceleratorResponseBody {
	s.CrossDomainBandwidthPackage = v
	return s
}

func (s *GetBasicAcceleratorResponseBody) SetCrossPrivateState(v string) *GetBasicAcceleratorResponseBody {
	s.CrossPrivateState = &v
	return s
}

func (s *GetBasicAcceleratorResponseBody) SetDescription(v string) *GetBasicAcceleratorResponseBody {
	s.Description = &v
	return s
}

func (s *GetBasicAcceleratorResponseBody) SetExpiredTime(v int64) *GetBasicAcceleratorResponseBody {
	s.ExpiredTime = &v
	return s
}

func (s *GetBasicAcceleratorResponseBody) SetInstanceChargeType(v string) *GetBasicAcceleratorResponseBody {
	s.InstanceChargeType = &v
	return s
}

func (s *GetBasicAcceleratorResponseBody) SetName(v string) *GetBasicAcceleratorResponseBody {
	s.Name = &v
	return s
}

func (s *GetBasicAcceleratorResponseBody) SetRegionId(v string) *GetBasicAcceleratorResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetBasicAcceleratorResponseBody) SetRequestId(v string) *GetBasicAcceleratorResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBasicAcceleratorResponseBody) SetResourceGroupId(v string) *GetBasicAcceleratorResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetBasicAcceleratorResponseBody) SetState(v string) *GetBasicAcceleratorResponseBody {
	s.State = &v
	return s
}

func (s *GetBasicAcceleratorResponseBody) SetTags(v []*GetBasicAcceleratorResponseBodyTags) *GetBasicAcceleratorResponseBody {
	s.Tags = v
	return s
}

func (s *GetBasicAcceleratorResponseBody) Validate() error {
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

type GetBasicAcceleratorResponseBodyBasicBandwidthPackage struct {
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

func (s GetBasicAcceleratorResponseBodyBasicBandwidthPackage) String() string {
	return dara.Prettify(s)
}

func (s GetBasicAcceleratorResponseBodyBasicBandwidthPackage) GoString() string {
	return s.String()
}

func (s *GetBasicAcceleratorResponseBodyBasicBandwidthPackage) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *GetBasicAcceleratorResponseBodyBasicBandwidthPackage) GetBandwidthType() *string {
	return s.BandwidthType
}

func (s *GetBasicAcceleratorResponseBodyBasicBandwidthPackage) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetBasicAcceleratorResponseBodyBasicBandwidthPackage) SetBandwidth(v int32) *GetBasicAcceleratorResponseBodyBasicBandwidthPackage {
	s.Bandwidth = &v
	return s
}

func (s *GetBasicAcceleratorResponseBodyBasicBandwidthPackage) SetBandwidthType(v string) *GetBasicAcceleratorResponseBodyBasicBandwidthPackage {
	s.BandwidthType = &v
	return s
}

func (s *GetBasicAcceleratorResponseBodyBasicBandwidthPackage) SetInstanceId(v string) *GetBasicAcceleratorResponseBodyBasicBandwidthPackage {
	s.InstanceId = &v
	return s
}

func (s *GetBasicAcceleratorResponseBodyBasicBandwidthPackage) Validate() error {
	return dara.Validate(s)
}

type GetBasicAcceleratorResponseBodyCrossDomainBandwidthPackage struct {
	// The bandwidth value of the cross-border acceleration bandwidth plan. Unit: Mbit/s.
	//
	// example:
	//
	// 2
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The ID of the cross-border acceleration bandwidth plan.
	//
	// example:
	//
	// gbwp-bp1d8xk8bg139j0fw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetBasicAcceleratorResponseBodyCrossDomainBandwidthPackage) String() string {
	return dara.Prettify(s)
}

func (s GetBasicAcceleratorResponseBodyCrossDomainBandwidthPackage) GoString() string {
	return s.String()
}

func (s *GetBasicAcceleratorResponseBodyCrossDomainBandwidthPackage) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *GetBasicAcceleratorResponseBodyCrossDomainBandwidthPackage) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetBasicAcceleratorResponseBodyCrossDomainBandwidthPackage) SetBandwidth(v int32) *GetBasicAcceleratorResponseBodyCrossDomainBandwidthPackage {
	s.Bandwidth = &v
	return s
}

func (s *GetBasicAcceleratorResponseBodyCrossDomainBandwidthPackage) SetInstanceId(v string) *GetBasicAcceleratorResponseBodyCrossDomainBandwidthPackage {
	s.InstanceId = &v
	return s
}

func (s *GetBasicAcceleratorResponseBodyCrossDomainBandwidthPackage) Validate() error {
	return dara.Validate(s)
}

type GetBasicAcceleratorResponseBodyTags struct {
	// The tag key.
	//
	// example:
	//
	// tag-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// tag-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetBasicAcceleratorResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s GetBasicAcceleratorResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetBasicAcceleratorResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *GetBasicAcceleratorResponseBodyTags) GetValue() *string {
	return s.Value
}

func (s *GetBasicAcceleratorResponseBodyTags) SetKey(v string) *GetBasicAcceleratorResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetBasicAcceleratorResponseBodyTags) SetValue(v string) *GetBasicAcceleratorResponseBodyTags {
	s.Value = &v
	return s
}

func (s *GetBasicAcceleratorResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
