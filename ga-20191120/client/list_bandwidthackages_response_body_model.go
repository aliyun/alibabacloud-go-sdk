// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBandwidthackagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidthPackages(v []*ListBandwidthackagesResponseBodyBandwidthPackages) *ListBandwidthackagesResponseBody
	GetBandwidthPackages() []*ListBandwidthackagesResponseBodyBandwidthPackages
	SetPageNumber(v int32) *ListBandwidthackagesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListBandwidthackagesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListBandwidthackagesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListBandwidthackagesResponseBody
	GetTotalCount() *int32
}

type ListBandwidthackagesResponseBody struct {
	// The details of the bandwidth plans.
	BandwidthPackages []*ListBandwidthackagesResponseBodyBandwidthPackages `json:"BandwidthPackages,omitempty" xml:"BandwidthPackages,omitempty" type:"Repeated"`
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

func (s ListBandwidthackagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBandwidthackagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListBandwidthackagesResponseBody) GetBandwidthPackages() []*ListBandwidthackagesResponseBodyBandwidthPackages {
	return s.BandwidthPackages
}

func (s *ListBandwidthackagesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBandwidthackagesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBandwidthackagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBandwidthackagesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListBandwidthackagesResponseBody) SetBandwidthPackages(v []*ListBandwidthackagesResponseBodyBandwidthPackages) *ListBandwidthackagesResponseBody {
	s.BandwidthPackages = v
	return s
}

func (s *ListBandwidthackagesResponseBody) SetPageNumber(v int32) *ListBandwidthackagesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListBandwidthackagesResponseBody) SetPageSize(v int32) *ListBandwidthackagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListBandwidthackagesResponseBody) SetRequestId(v string) *ListBandwidthackagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBandwidthackagesResponseBody) SetTotalCount(v int32) *ListBandwidthackagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListBandwidthackagesResponseBody) Validate() error {
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

type ListBandwidthackagesResponseBodyBandwidthPackages struct {
	// The IDs of the GA instances that are associated with the bandwidth plans.
	Accelerators []*string `json:"Accelerators,omitempty" xml:"Accelerators,omitempty" type:"Repeated"`
	// The bandwidth of the bandwidth plan. Unit: Mbit/s.
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
	// The billing method of the bandwidth plan. Valid values:
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
	// The name of the GA instance.
	//
	// example:
	//
	// Accelerator
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DE77A7F3-3B74-41C0-A5BC-CAFD188C28B6
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzzwgr7vz2liy
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
	// 	- **locked:*	- The bandwidth plan is locked.
	//
	// example:
	//
	// active
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The tag of the bandwidth plan.
	Tags []*ListBandwidthackagesResponseBodyBandwidthPackagesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListBandwidthackagesResponseBodyBandwidthPackages) String() string {
	return dara.Prettify(s)
}

func (s ListBandwidthackagesResponseBodyBandwidthPackages) GoString() string {
	return s.String()
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) GetAccelerators() []*string {
	return s.Accelerators
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) GetBandwidthPackageId() *string {
	return s.BandwidthPackageId
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) GetDescription() *string {
	return s.Description
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) GetName() *string {
	return s.Name
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) GetRegionId() *string {
	return s.RegionId
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) GetState() *string {
	return s.State
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) GetTags() []*ListBandwidthackagesResponseBodyBandwidthPackagesTags {
	return s.Tags
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) SetAccelerators(v []*string) *ListBandwidthackagesResponseBodyBandwidthPackages {
	s.Accelerators = v
	return s
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) SetBandwidth(v int32) *ListBandwidthackagesResponseBodyBandwidthPackages {
	s.Bandwidth = &v
	return s
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) SetBandwidthPackageId(v string) *ListBandwidthackagesResponseBodyBandwidthPackages {
	s.BandwidthPackageId = &v
	return s
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) SetChargeType(v string) *ListBandwidthackagesResponseBodyBandwidthPackages {
	s.ChargeType = &v
	return s
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) SetCreateTime(v string) *ListBandwidthackagesResponseBodyBandwidthPackages {
	s.CreateTime = &v
	return s
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) SetDescription(v string) *ListBandwidthackagesResponseBodyBandwidthPackages {
	s.Description = &v
	return s
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) SetExpiredTime(v string) *ListBandwidthackagesResponseBodyBandwidthPackages {
	s.ExpiredTime = &v
	return s
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) SetName(v string) *ListBandwidthackagesResponseBodyBandwidthPackages {
	s.Name = &v
	return s
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) SetRegionId(v string) *ListBandwidthackagesResponseBodyBandwidthPackages {
	s.RegionId = &v
	return s
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) SetResourceGroupId(v string) *ListBandwidthackagesResponseBodyBandwidthPackages {
	s.ResourceGroupId = &v
	return s
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) SetState(v string) *ListBandwidthackagesResponseBodyBandwidthPackages {
	s.State = &v
	return s
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) SetTags(v []*ListBandwidthackagesResponseBodyBandwidthPackagesTags) *ListBandwidthackagesResponseBodyBandwidthPackages {
	s.Tags = v
	return s
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackages) Validate() error {
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

type ListBandwidthackagesResponseBodyBandwidthPackagesTags struct {
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

func (s ListBandwidthackagesResponseBodyBandwidthPackagesTags) String() string {
	return dara.Prettify(s)
}

func (s ListBandwidthackagesResponseBodyBandwidthPackagesTags) GoString() string {
	return s.String()
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackagesTags) GetKey() *string {
	return s.Key
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackagesTags) GetValue() *string {
	return s.Value
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackagesTags) SetKey(v string) *ListBandwidthackagesResponseBodyBandwidthPackagesTags {
	s.Key = &v
	return s
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackagesTags) SetValue(v string) *ListBandwidthackagesResponseBodyBandwidthPackagesTags {
	s.Value = &v
	return s
}

func (s *ListBandwidthackagesResponseBodyBandwidthPackagesTags) Validate() error {
	return dara.Validate(s)
}
