// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBandwidthPackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerators(v []*string) *DescribeBandwidthPackageResponseBody
	GetAccelerators() []*string
	SetBandwidth(v int32) *DescribeBandwidthPackageResponseBody
	GetBandwidth() *int32
	SetBandwidthPackageId(v string) *DescribeBandwidthPackageResponseBody
	GetBandwidthPackageId() *string
	SetBandwidthType(v string) *DescribeBandwidthPackageResponseBody
	GetBandwidthType() *string
	SetBillingType(v string) *DescribeBandwidthPackageResponseBody
	GetBillingType() *string
	SetCbnGeographicRegionIdA(v string) *DescribeBandwidthPackageResponseBody
	GetCbnGeographicRegionIdA() *string
	SetCbnGeographicRegionIdB(v string) *DescribeBandwidthPackageResponseBody
	GetCbnGeographicRegionIdB() *string
	SetChargeType(v string) *DescribeBandwidthPackageResponseBody
	GetChargeType() *string
	SetCreateTime(v string) *DescribeBandwidthPackageResponseBody
	GetCreateTime() *string
	SetDescription(v string) *DescribeBandwidthPackageResponseBody
	GetDescription() *string
	SetExpiredTime(v string) *DescribeBandwidthPackageResponseBody
	GetExpiredTime() *string
	SetName(v string) *DescribeBandwidthPackageResponseBody
	GetName() *string
	SetRatio(v int32) *DescribeBandwidthPackageResponseBody
	GetRatio() *int32
	SetRegionId(v string) *DescribeBandwidthPackageResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeBandwidthPackageResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *DescribeBandwidthPackageResponseBody
	GetResourceGroupId() *string
	SetState(v string) *DescribeBandwidthPackageResponseBody
	GetState() *string
	SetTags(v []*DescribeBandwidthPackageResponseBodyTags) *DescribeBandwidthPackageResponseBody
	GetTags() []*DescribeBandwidthPackageResponseBodyTags
	SetType(v string) *DescribeBandwidthPackageResponseBody
	GetType() *string
}

type DescribeBandwidthPackageResponseBody struct {
	// The instance ID of the Alibaba Cloud Global Accelerator (GA) instance attached to the bandwidth plan.
	Accelerators []*string `json:"Accelerators,omitempty" xml:"Accelerators,omitempty" type:"Repeated"`
	// The bandwidth value of the bandwidth plan. Unit: Mbit/s.
	//
	// example:
	//
	// 2
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The bandwidth plan ID.
	//
	// example:
	//
	// gbwp-bp1sgzldyj6b4q7cx****
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The bandwidth type. Valid values:
	//
	// - **Basic**: standard acceleration bandwidth.
	//
	// - **Enhanced**: enhanced acceleration bandwidth.
	//
	// - **Advanced**: premium acceleration bandwidth.
	//
	// example:
	//
	// Basic
	BandwidthType *string `json:"BandwidthType,omitempty" xml:"BandwidthType,omitempty"`
	// The billable methods for the pay-as-you-go billing method. Valid values:
	//
	// - **PayByTraffic**: pay-by-data-transfer.
	//
	// - **PayBY95**: pay-by-95th-percentile.
	//
	// example:
	//
	// PayByTraffic
	BillingType *string `json:"BillingType,omitempty" xml:"BillingType,omitempty"`
	// The interconnected area A of the cross-border acceleration bandwidth plan. The value is returned only as **China-mainland*	- (the Chinese mainland).
	//
	// This parameter is returned only on the Alibaba Cloud International Website (www.alibabacloud.com).
	//
	// example:
	//
	// China-mainland
	CbnGeographicRegionIdA *string `json:"CbnGeographicRegionIdA,omitempty" xml:"CbnGeographicRegionIdA,omitempty"`
	// The interconnected area B of the cross-border acceleration bandwidth plan. The value is returned only as **Global**.
	//
	// This parameter is returned only on the Chinese site (Chinese mainland).
	//
	// example:
	//
	// Global
	CbnGeographicRegionIdB *string `json:"CbnGeographicRegionIdB,omitempty" xml:"CbnGeographicRegionIdB,omitempty"`
	// The billing method. Valid values:
	//
	// - **PREPAY*	- (default): subscription.
	//
	// - **POSTPAY**: pay-as-you-go.
	//
	// example:
	//
	// PREPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The timestamp when the bandwidth plan was created.
	//
	// example:
	//
	// 1578966918000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the bandwidth plan.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The timestamp when the bandwidth plan expires.
	//
	// example:
	//
	// 1578966918000
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The name of the bandwidth plan.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The minimum percentage for the pay-by-95th-percentile metering method. Valid values: **30*	- to **100**.
	//
	// example:
	//
	// 30
	Ratio *int32 `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	// The region ID of the Alibaba Cloud Global Accelerator (GA) instance. Set the value to **cn-hangzhou**.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4B6DBBB0-2D01-4C6A-A384-4129266E6B78
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfnvueepcihjiq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the bandwidth plan. Valid values:
	//
	// - **init**: initialization.
	//
	// - **active**: active.
	//
	// - **binded**: attached.
	//
	// - **binding**: being attached.
	//
	// - **unbinding**: being disassociated.
	//
	// - **updating**: being updated.
	//
	// - **finacialLocked**: locked due to overdue payment.
	//
	// - **locked**: locked.
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The resource tags.
	Tags []*DescribeBandwidthPackageResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The type of the bandwidth plan. Valid values:
	//
	// - **Basic**: basic bandwidth plan.
	//
	// - **CrossDomain**: cross-border acceleration bandwidth plan.
	//
	// Only **Basic*	- is returned on the Alibaba Cloud China Website (www.aliyun.com).
	//
	// example:
	//
	// Basic
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeBandwidthPackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBandwidthPackageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBandwidthPackageResponseBody) GetAccelerators() []*string {
	return s.Accelerators
}

func (s *DescribeBandwidthPackageResponseBody) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeBandwidthPackageResponseBody) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *DescribeBandwidthPackageResponseBody) GetBandwidthType() *string {
	return s.BandwidthType
}

func (s *DescribeBandwidthPackageResponseBody) GetBillingType() *string {
	return s.BillingType
}

func (s *DescribeBandwidthPackageResponseBody) GetCbnGeographicRegionIdA() *string {
	return s.CbnGeographicRegionIdA
}

func (s *DescribeBandwidthPackageResponseBody) GetCbnGeographicRegionIdB() *string {
	return s.CbnGeographicRegionIdB
}

func (s *DescribeBandwidthPackageResponseBody) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeBandwidthPackageResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeBandwidthPackageResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeBandwidthPackageResponseBody) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeBandwidthPackageResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeBandwidthPackageResponseBody) GetRatio() *int32 {
	return s.Ratio
}

func (s *DescribeBandwidthPackageResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeBandwidthPackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBandwidthPackageResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeBandwidthPackageResponseBody) GetState() *string {
	return s.State
}

func (s *DescribeBandwidthPackageResponseBody) GetTags() []*DescribeBandwidthPackageResponseBodyTags {
	return s.Tags
}

func (s *DescribeBandwidthPackageResponseBody) GetType() *string {
	return s.Type
}

func (s *DescribeBandwidthPackageResponseBody) SetAccelerators(v []*string) *DescribeBandwidthPackageResponseBody {
	s.Accelerators = v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetBandwidth(v int32) *DescribeBandwidthPackageResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetBandwidthPackageId(v string) *DescribeBandwidthPackageResponseBody {
	s.BandwidthPackageId = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetBandwidthType(v string) *DescribeBandwidthPackageResponseBody {
	s.BandwidthType = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetBillingType(v string) *DescribeBandwidthPackageResponseBody {
	s.BillingType = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetCbnGeographicRegionIdA(v string) *DescribeBandwidthPackageResponseBody {
	s.CbnGeographicRegionIdA = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetCbnGeographicRegionIdB(v string) *DescribeBandwidthPackageResponseBody {
	s.CbnGeographicRegionIdB = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetChargeType(v string) *DescribeBandwidthPackageResponseBody {
	s.ChargeType = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetCreateTime(v string) *DescribeBandwidthPackageResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetDescription(v string) *DescribeBandwidthPackageResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetExpiredTime(v string) *DescribeBandwidthPackageResponseBody {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetName(v string) *DescribeBandwidthPackageResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetRatio(v int32) *DescribeBandwidthPackageResponseBody {
	s.Ratio = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetRegionId(v string) *DescribeBandwidthPackageResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetRequestId(v string) *DescribeBandwidthPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetResourceGroupId(v string) *DescribeBandwidthPackageResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetState(v string) *DescribeBandwidthPackageResponseBody {
	s.State = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetTags(v []*DescribeBandwidthPackageResponseBodyTags) *DescribeBandwidthPackageResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) SetType(v string) *DescribeBandwidthPackageResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBody) Validate() error {
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

type DescribeBandwidthPackageResponseBodyTags struct {
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

func (s DescribeBandwidthPackageResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeBandwidthPackageResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeBandwidthPackageResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *DescribeBandwidthPackageResponseBodyTags) GetValue() *string {
	return s.Value
}

func (s *DescribeBandwidthPackageResponseBodyTags) SetKey(v string) *DescribeBandwidthPackageResponseBodyTags {
	s.Key = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBodyTags) SetValue(v string) *DescribeBandwidthPackageResponseBodyTags {
	s.Value = &v
	return s
}

func (s *DescribeBandwidthPackageResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
