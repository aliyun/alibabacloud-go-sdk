// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserResourcePackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeUserResourcePackageResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeUserResourcePackageResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeUserResourcePackageResponseBody
	GetRequestId() *string
	SetResourcePackageInfos(v []*DescribeUserResourcePackageResponseBodyResourcePackageInfos) *DescribeUserResourcePackageResponseBody
	GetResourcePackageInfos() []*DescribeUserResourcePackageResponseBodyResourcePackageInfos
	SetTotalCount(v int32) *DescribeUserResourcePackageResponseBody
	GetTotalCount() *int32
}

type DescribeUserResourcePackageResponseBody struct {
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-A198-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The array of ResourcePackageInfo objects.
	ResourcePackageInfos []*DescribeUserResourcePackageResponseBodyResourcePackageInfos `json:"ResourcePackageInfos,omitempty" xml:"ResourcePackageInfos,omitempty" type:"Repeated"`
	// The total number of records.
	//
	// example:
	//
	// 68
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeUserResourcePackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserResourcePackageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserResourcePackageResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeUserResourcePackageResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeUserResourcePackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserResourcePackageResponseBody) GetResourcePackageInfos() []*DescribeUserResourcePackageResponseBodyResourcePackageInfos {
	return s.ResourcePackageInfos
}

func (s *DescribeUserResourcePackageResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeUserResourcePackageResponseBody) SetPageNumber(v int32) *DescribeUserResourcePackageResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeUserResourcePackageResponseBody) SetPageSize(v int32) *DescribeUserResourcePackageResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeUserResourcePackageResponseBody) SetRequestId(v string) *DescribeUserResourcePackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserResourcePackageResponseBody) SetResourcePackageInfos(v []*DescribeUserResourcePackageResponseBodyResourcePackageInfos) *DescribeUserResourcePackageResponseBody {
	s.ResourcePackageInfos = v
	return s
}

func (s *DescribeUserResourcePackageResponseBody) SetTotalCount(v int32) *DescribeUserResourcePackageResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeUserResourcePackageResponseBody) Validate() error {
	if s.ResourcePackageInfos != nil {
		for _, item := range s.ResourcePackageInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUserResourcePackageResponseBodyResourcePackageInfos struct {
	// The commodity code of the resource plan.
	//
	// example:
	//
	// dcdnpaybag
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The current remaining capacity of the instance.
	//
	// - Unit for traffic plans: Byte.
	//
	// - Unit for request plans: count.
	//
	// example:
	//
	// 53661095687
	CurrCapacity *string `json:"CurrCapacity,omitempty" xml:"CurrCapacity,omitempty"`
	// The base unit of the current remaining capacity of the instance.
	//
	// example:
	//
	// Byte
	CurrCapacityBaseUnit *string `json:"CurrCapacityBaseUnit,omitempty" xml:"CurrCapacityBaseUnit,omitempty"`
	// The display unit of the current remaining capacity of the instance.
	//
	// example:
	//
	// GB
	CurrCapacityShowUnit *string `json:"CurrCapacityShowUnit,omitempty" xml:"CurrCapacityShowUnit,omitempty"`
	// The display value of the current remaining capacity of the instance.
	//
	// example:
	//
	// 49.975789
	CurrCapacityShowValue *string `json:"CurrCapacityShowValue,omitempty" xml:"CurrCapacityShowValue,omitempty"`
	// The name of the resource plan.
	//
	// example:
	//
	// Data Transfer Plan in Asia Pacific 1
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The expiration time in UTC. Format: YYYY-MM-DDTHH:mm:ssZ.
	//
	// example:
	//
	// 2024-12-02T15:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The total capacity of the resource plan.
	//
	// - Unit for traffic plans: Byte.
	//
	// - Unit for request plans: count.
	//
	// example:
	//
	// 107374182400
	InitCapacity *string `json:"InitCapacity,omitempty" xml:"InitCapacity,omitempty"`
	// The base unit of the total capacity of the resource plan.
	//
	// example:
	//
	// Byte
	InitCapacityBaseUnit *string `json:"InitCapacityBaseUnit,omitempty" xml:"InitCapacityBaseUnit,omitempty"`
	// The display unit of the total capacity of the resource plan.
	//
	// example:
	//
	// GB
	InitCapacityShowUnit *string `json:"InitCapacityShowUnit,omitempty" xml:"InitCapacityShowUnit,omitempty"`
	// The display value of the total capacity of the resource plan.
	//
	// example:
	//
	// 100.000000
	InitCapacityShowValue *string `json:"InitCapacityShowValue,omitempty" xml:"InitCapacityShowValue,omitempty"`
	// The instance ID of the resource plan.
	//
	// example:
	//
	// ****_ResourcePack-cn-****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region.
	//
	// example:
	//
	// CN
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The effective period in UTC. Format: YYYY-MM-DDTHH:mm:ssZ.
	//
	// example:
	//
	// 2024-03-20T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the resource plan.
	//
	// example:
	//
	// valid
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The template name.
	//
	// example:
	//
	// FPT_dcdnpaybag_deadlineAcc_****
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s DescribeUserResourcePackageResponseBodyResourcePackageInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserResourcePackageResponseBodyResourcePackageInfos) GoString() string {
	return s.String()
}

func (s *DescribeUserResourcePackageResponseBodyResourcePackageInfos) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeUserResourcePackageResponseBodyResourcePackageInfos) GetCurrCapacity() *string {
	return s.CurrCapacity
}

func (s *DescribeUserResourcePackageResponseBodyResourcePackageInfos) GetCurrCapacityBaseUnit() *string {
	return s.CurrCapacityBaseUnit
}

func (s *DescribeUserResourcePackageResponseBodyResourcePackageInfos) GetCurrCapacityShowUnit() *string {
	return s.CurrCapacityShowUnit
}

func (s *DescribeUserResourcePackageResponseBodyResourcePackageInfos) GetCurrCapacityShowValue() *string {
	return s.CurrCapacityShowValue
}

func (s *DescribeUserResourcePackageResponseBodyResourcePackageInfos) GetDisplayName() *string {
	return s.DisplayName
}

func (s *DescribeUserResourcePackageResponseBodyResourcePackageInfos) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeUserResourcePackageResponseBodyResourcePackageInfos) GetInitCapacity() *string {
	return s.InitCapacity
}

func (s *DescribeUserResourcePackageResponseBodyResourcePackageInfos) GetInitCapacityBaseUnit() *string {
	return s.InitCapacityBaseUnit
}

func (s *DescribeUserResourcePackageResponseBodyResourcePackageInfos) GetInitCapacityShowUnit() *string {
	return s.InitCapacityShowUnit
}

func (s *DescribeUserResourcePackageResponseBodyResourcePackageInfos) GetInitCapacityShowValue() *string {
	return s.InitCapacityShowValue
}

func (s *DescribeUserResourcePackageResponseBodyResourcePackageInfos) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeUserResourcePackageResponseBodyResourcePackageInfos) GetRegion() *string {
	return s.Region
}

func (s *DescribeUserResourcePackageResponseBodyResourcePackageInfos) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeUserResourcePackageResponseBodyResourcePackageInfos) GetStatus() *string {
	return s.Status
}

func (s *DescribeUserResourcePackageResponseBodyResourcePackageInfos) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DescribeUserResourcePackageResponseBodyResourcePackageInfos) SetCommodityCode(v string) *DescribeUserResourcePackageResponseBodyResourcePackageInfos {
	s.CommodityCode = &v
	return s
}

func (s *DescribeUserResourcePackageResponseBodyResourcePackageInfos) SetCurrCapacity(v string) *DescribeUserResourcePackageResponseBodyResourcePackageInfos {
	s.CurrCapacity = &v
	return s
}

func (s *DescribeUserResourcePackageResponseBodyResourcePackageInfos) SetCurrCapacityBaseUnit(v string) *DescribeUserResourcePackageResponseBodyResourcePackageInfos {
	s.CurrCapacityBaseUnit = &v
	return s
}

func (s *DescribeUserResourcePackageResponseBodyResourcePackageInfos) SetCurrCapacityShowUnit(v string) *DescribeUserResourcePackageResponseBodyResourcePackageInfos {
	s.CurrCapacityShowUnit = &v
	return s
}

func (s *DescribeUserResourcePackageResponseBodyResourcePackageInfos) SetCurrCapacityShowValue(v string) *DescribeUserResourcePackageResponseBodyResourcePackageInfos {
	s.CurrCapacityShowValue = &v
	return s
}

func (s *DescribeUserResourcePackageResponseBodyResourcePackageInfos) SetDisplayName(v string) *DescribeUserResourcePackageResponseBodyResourcePackageInfos {
	s.DisplayName = &v
	return s
}

func (s *DescribeUserResourcePackageResponseBodyResourcePackageInfos) SetEndTime(v string) *DescribeUserResourcePackageResponseBodyResourcePackageInfos {
	s.EndTime = &v
	return s
}

func (s *DescribeUserResourcePackageResponseBodyResourcePackageInfos) SetInitCapacity(v string) *DescribeUserResourcePackageResponseBodyResourcePackageInfos {
	s.InitCapacity = &v
	return s
}

func (s *DescribeUserResourcePackageResponseBodyResourcePackageInfos) SetInitCapacityBaseUnit(v string) *DescribeUserResourcePackageResponseBodyResourcePackageInfos {
	s.InitCapacityBaseUnit = &v
	return s
}

func (s *DescribeUserResourcePackageResponseBodyResourcePackageInfos) SetInitCapacityShowUnit(v string) *DescribeUserResourcePackageResponseBodyResourcePackageInfos {
	s.InitCapacityShowUnit = &v
	return s
}

func (s *DescribeUserResourcePackageResponseBodyResourcePackageInfos) SetInitCapacityShowValue(v string) *DescribeUserResourcePackageResponseBodyResourcePackageInfos {
	s.InitCapacityShowValue = &v
	return s
}

func (s *DescribeUserResourcePackageResponseBodyResourcePackageInfos) SetInstanceId(v string) *DescribeUserResourcePackageResponseBodyResourcePackageInfos {
	s.InstanceId = &v
	return s
}

func (s *DescribeUserResourcePackageResponseBodyResourcePackageInfos) SetRegion(v string) *DescribeUserResourcePackageResponseBodyResourcePackageInfos {
	s.Region = &v
	return s
}

func (s *DescribeUserResourcePackageResponseBodyResourcePackageInfos) SetStartTime(v string) *DescribeUserResourcePackageResponseBodyResourcePackageInfos {
	s.StartTime = &v
	return s
}

func (s *DescribeUserResourcePackageResponseBodyResourcePackageInfos) SetStatus(v string) *DescribeUserResourcePackageResponseBodyResourcePackageInfos {
	s.Status = &v
	return s
}

func (s *DescribeUserResourcePackageResponseBodyResourcePackageInfos) SetTemplateName(v string) *DescribeUserResourcePackageResponseBodyResourcePackageInfos {
	s.TemplateName = &v
	return s
}

func (s *DescribeUserResourcePackageResponseBodyResourcePackageInfos) Validate() error {
	return dara.Validate(s)
}
