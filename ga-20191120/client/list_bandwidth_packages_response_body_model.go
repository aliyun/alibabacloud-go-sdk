// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBandwidthPackagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidthPackages(v []*ListBandwidthPackagesResponseBodyBandwidthPackages) *ListBandwidthPackagesResponseBody
	GetBandwidthPackages() []*ListBandwidthPackagesResponseBodyBandwidthPackages
	SetPageNumber(v int32) *ListBandwidthPackagesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListBandwidthPackagesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListBandwidthPackagesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListBandwidthPackagesResponseBody
	GetTotalCount() *int32
}

type ListBandwidthPackagesResponseBody struct {
	// The details of the bandwidth plans.
	BandwidthPackages []*ListBandwidthPackagesResponseBodyBandwidthPackages `json:"BandwidthPackages,omitempty" xml:"BandwidthPackages,omitempty" type:"Repeated"`
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
	// 4B6DBBB0-2D01-4C6A-A384-4129266E6B78
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBandwidthPackagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBandwidthPackagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListBandwidthPackagesResponseBody) GetBandwidthPackages() []*ListBandwidthPackagesResponseBodyBandwidthPackages {
	return s.BandwidthPackages
}

func (s *ListBandwidthPackagesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBandwidthPackagesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBandwidthPackagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBandwidthPackagesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListBandwidthPackagesResponseBody) SetBandwidthPackages(v []*ListBandwidthPackagesResponseBodyBandwidthPackages) *ListBandwidthPackagesResponseBody {
	s.BandwidthPackages = v
	return s
}

func (s *ListBandwidthPackagesResponseBody) SetPageNumber(v int32) *ListBandwidthPackagesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListBandwidthPackagesResponseBody) SetPageSize(v int32) *ListBandwidthPackagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListBandwidthPackagesResponseBody) SetRequestId(v string) *ListBandwidthPackagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBandwidthPackagesResponseBody) SetTotalCount(v int32) *ListBandwidthPackagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListBandwidthPackagesResponseBody) Validate() error {
	if s.BandwidthPackages != nil {
		for _, item := range s.BandwidthPackages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListBandwidthPackagesResponseBodyBandwidthPackages struct {
	// The IDs of the GA instances that are associated with the bandwidth plans.
	Accelerators []*string `json:"Accelerators,omitempty" xml:"Accelerators,omitempty" type:"Repeated"`
	// The bandwidth value of the bandwidth plan. Unit: Mbit/s.
	//
	// example:
	//
	// 2
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The ID of the bandwidth plan.
	//
	// example:
	//
	// gbwp-bp1sgzldyj6b4q7cx****
	BandwidthPackageId *string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty"`
	// The type of the bandwidth. Valid values:
	//
	// 	- **Basic**
	//
	// 	- **Enhanced**
	//
	// 	- **Advanced**
	//
	// example:
	//
	// Basic
	BandwidthType *string `json:"BandwidthType,omitempty" xml:"BandwidthType,omitempty"`
	// The metering method that is used when you use the pay-as-you-go billing method.
	//
	// 	- **PayByTraffic**: pay-by-data-transfer.
	//
	// 	- **PayBY95**: pay-by-95th-percentile.
	//
	// example:
	//
	// PayByTraffic
	BillingType *string `json:"BillingType,omitempty" xml:"BillingType,omitempty"`
	// Area A of the cross-region acceleration bandwidth plan. **China-mainland*	- is returned.
	//
	// This parameter is returned only if you call the operation on the international site (alibabacloud.com).
	//
	// example:
	//
	// China-mainland
	CbnGeographicRegionIdA *string `json:"CbnGeographicRegionIdA,omitempty" xml:"CbnGeographicRegionIdA,omitempty"`
	// Area B of the cross-region acceleration bandwidth plan. **Global*	- is returned.
	//
	// This parameter is returned only if you call the operation on the international site (alibabacloud.com).
	//
	// example:
	//
	// Global
	CbnGeographicRegionIdB *string `json:"CbnGeographicRegionIdB,omitempty" xml:"CbnGeographicRegionIdB,omitempty"`
	// The billing method of the bandwidth plan.
	//
	// 	- **PREPAY**: subscription. This is the default value.
	//
	// 	- **POSTPAY**: pay-as-you-go.
	//
	// example:
	//
	// PREPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The time when the bandwidth plan was created.
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
	// The expiration time of the bandwidth plan.
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
	// The percentage of the guaranteed minimum bandwidth if the pay-by-95th-percentile metering method is used.
	//
	// example:
	//
	// 30
	Ratio *int32 `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	// The ID of the region where the GA instance is deployed. **cn-hangzhou*	- is returned.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek2ry6mp2c****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the bandwidth plan. Valid values:
	//
	// 	- **init:*	- The bandwidth plan is being initialized.
	//
	// 	- **active:*	- The bandwidth plan is available.
	//
	// 	- **binded:*	- The bandwidth plan is associated with a GA instance.
	//
	// 	- **binding:*	- The bandwidth plan is being associated with a GA instance.
	//
	// 	- **unbinding:*	- The bandwidth plan is being disassociated from a GA instance.
	//
	// 	- **updating:*	- The bandwidth plan is being updated.
	//
	// 	- **finacialLocked:*	- The bandwidth plan is locked due to overdue payments.
	//
	// 	- **locked**: The bandwidth plan is locked.
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The tag of the bandwidth plan.
	Tags []*ListBandwidthPackagesResponseBodyBandwidthPackagesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The type of the bandwidth plan. Valid values:
	//
	// 	- **Basic:*	- a basic bandwidth plan.
	//
	// 	- **CrossDomain:*	- a cross-region acceleration bandwidth plan.
	//
	// If you call the operation on the China site (aliyun.com), **Basic*	- is returned.
	//
	// example:
	//
	// Basic
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListBandwidthPackagesResponseBodyBandwidthPackages) String() string {
	return dara.Prettify(s)
}

func (s ListBandwidthPackagesResponseBodyBandwidthPackages) GoString() string {
	return s.String()
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) GetAccelerators() []*string {
	return s.Accelerators
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) GetBandwidthType() *string {
	return s.BandwidthType
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) GetBillingType() *string {
	return s.BillingType
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) GetCbnGeographicRegionIdA() *string {
	return s.CbnGeographicRegionIdA
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) GetCbnGeographicRegionIdB() *string {
	return s.CbnGeographicRegionIdB
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) GetDescription() *string {
	return s.Description
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) GetName() *string {
	return s.Name
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) GetRatio() *int32 {
	return s.Ratio
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) GetRegionId() *string {
	return s.RegionId
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) GetState() *string {
	return s.State
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) GetTags() []*ListBandwidthPackagesResponseBodyBandwidthPackagesTags {
	return s.Tags
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) GetType() *string {
	return s.Type
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetAccelerators(v []*string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.Accelerators = v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetBandwidth(v int32) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.Bandwidth = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetBandwidthPackageId(v string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.BandwidthPackageId = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetBandwidthType(v string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.BandwidthType = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetBillingType(v string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.BillingType = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetCbnGeographicRegionIdA(v string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.CbnGeographicRegionIdA = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetCbnGeographicRegionIdB(v string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.CbnGeographicRegionIdB = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetChargeType(v string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.ChargeType = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetCreateTime(v string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.CreateTime = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetDescription(v string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.Description = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetExpiredTime(v string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.ExpiredTime = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetName(v string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.Name = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetRatio(v int32) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.Ratio = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetRegionId(v string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.RegionId = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetResourceGroupId(v string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.ResourceGroupId = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetState(v string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.State = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetTags(v []*ListBandwidthPackagesResponseBodyBandwidthPackagesTags) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.Tags = v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) SetType(v string) *ListBandwidthPackagesResponseBodyBandwidthPackages {
	s.Type = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackages) Validate() error {
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

type ListBandwidthPackagesResponseBodyBandwidthPackagesTags struct {
	// The tag key of the bandwidth plan.
	//
	// example:
	//
	// tag-key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the bandwidth plan.
	//
	// example:
	//
	// tag-value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListBandwidthPackagesResponseBodyBandwidthPackagesTags) String() string {
	return dara.Prettify(s)
}

func (s ListBandwidthPackagesResponseBodyBandwidthPackagesTags) GoString() string {
	return s.String()
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackagesTags) GetKey() *string {
	return s.Key
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackagesTags) GetValue() *string {
	return s.Value
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackagesTags) SetKey(v string) *ListBandwidthPackagesResponseBodyBandwidthPackagesTags {
	s.Key = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackagesTags) SetValue(v string) *ListBandwidthPackagesResponseBodyBandwidthPackagesTags {
	s.Value = &v
	return s
}

func (s *ListBandwidthPackagesResponseBodyBandwidthPackagesTags) Validate() error {
	return dara.Validate(s)
}
