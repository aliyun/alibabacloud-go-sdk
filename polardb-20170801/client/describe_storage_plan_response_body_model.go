// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStoragePlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeStoragePlanResponseBodyItems) *DescribeStoragePlanResponseBody
	GetItems() []*DescribeStoragePlanResponseBodyItems
	SetPageNumber(v int64) *DescribeStoragePlanResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeStoragePlanResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeStoragePlanResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int64) *DescribeStoragePlanResponseBody
	GetTotalRecordCount() *int64
}

type DescribeStoragePlanResponseBody struct {
	// The list of storage plans.
	Items []*DescribeStoragePlanResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number. The value must be a positive integer that does not exceed the maximum value of the Integer data type.
	//
	// example:
	//
	// 10
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// - **30**
	//
	// - **50**
	//
	// - **100**
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D9384D30-5C9A-11EA-BFC0-F754D2Bxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 10
	TotalRecordCount *int64 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeStoragePlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeStoragePlanResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStoragePlanResponseBody) GetItems() []*DescribeStoragePlanResponseBodyItems {
	return s.Items
}

func (s *DescribeStoragePlanResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeStoragePlanResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeStoragePlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeStoragePlanResponseBody) GetTotalRecordCount() *int64 {
	return s.TotalRecordCount
}

func (s *DescribeStoragePlanResponseBody) SetItems(v []*DescribeStoragePlanResponseBodyItems) *DescribeStoragePlanResponseBody {
	s.Items = v
	return s
}

func (s *DescribeStoragePlanResponseBody) SetPageNumber(v int64) *DescribeStoragePlanResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeStoragePlanResponseBody) SetPageSize(v int64) *DescribeStoragePlanResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeStoragePlanResponseBody) SetRequestId(v string) *DescribeStoragePlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStoragePlanResponseBody) SetTotalRecordCount(v int64) *DescribeStoragePlanResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeStoragePlanResponseBody) Validate() error {
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

type DescribeStoragePlanResponseBodyItems struct {
	// The Alibaba Cloud account ID.
	//
	// example:
	//
	// 18170112322*****
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The commodity code.
	//
	// example:
	//
	// polardb_package
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The expiration time of the storage plan.
	//
	// example:
	//
	// 2021-10-23 00:00:00
	EndTimes *string `json:"EndTimes,omitempty" xml:"EndTimes,omitempty"`
	// The capacity unit of the storage plan.
	//
	// example:
	//
	// GB
	InitCapaCityViewUnit *string `json:"InitCapaCityViewUnit,omitempty" xml:"InitCapaCityViewUnit,omitempty"`
	// The total capacity of the storage plan.
	//
	// example:
	//
	// 500.000
	InitCapacityViewValue *string `json:"InitCapacityViewValue,omitempty" xml:"InitCapacityViewValue,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 500072****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The periodic capacity unit of the storage plan.
	//
	// example:
	//
	// GB
	PeriodCapaCityViewUnit *string `json:"PeriodCapaCityViewUnit,omitempty" xml:"PeriodCapaCityViewUnit,omitempty"`
	// The periodic capacity of the storage plan. The periodic capacity gradually decreases as the storage plan is used.
	//
	// example:
	//
	// 400.000
	PeriodCapacityViewValue *string `json:"PeriodCapacityViewValue,omitempty" xml:"PeriodCapacityViewValue,omitempty"`
	// The period duration, which indicates that the capacity of the storage plan is measured and consumed on an hourly basis.
	//
	// example:
	//
	// hour
	PeriodTime *string `json:"PeriodTime,omitempty" xml:"PeriodTime,omitempty"`
	// The product name.
	//
	// example:
	//
	// polardb
	ProdCode *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	// The purchase time of the storage plan.
	//
	// example:
	//
	// 2021-09-22 12:03:27
	PurchaseTimes *string `json:"PurchaseTimes,omitempty" xml:"PurchaseTimes,omitempty"`
	// The effective period of the storage plan.
	//
	// example:
	//
	// 2021-09-22 12:03:25
	StartTimes *string `json:"StartTimes,omitempty" xml:"StartTimes,omitempty"`
	// The status of the storage plan. Valid values:
	//
	// - **Valid**: The storage plan is valid.
	//
	// - **Invalid**: The storage plan is invalid.
	//
	// example:
	//
	// Valid
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the storage plan. Valid values:
	//
	// - **Overseas**: Hong Kong (China) and regions outside the Chinese mainland.
	//
	// - **Mainland**: The Chinese mainland.
	//
	// - **Financial**: Alibaba Finance Cloud.
	//
	// example:
	//
	// Overseas
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The type of the resource plan. Valid values:
	//
	//  - **FPT_polardb_package_absolute_1540888383**: The Chinese mainland.
	//
	// - **FPT_polardb_package_absolute_1540890081**: Hong Kong (China) and regions outside the Chinese mainland.
	//
	// - **FPT_polardb_package_absolute_1567151499**: Alibaba Finance Cloud.
	//
	// example:
	//
	// FPT_polardb_package_absolute_1540890081
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s DescribeStoragePlanResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeStoragePlanResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeStoragePlanResponseBodyItems) GetAliUid() *string {
	return s.AliUid
}

func (s *DescribeStoragePlanResponseBodyItems) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeStoragePlanResponseBodyItems) GetEndTimes() *string {
	return s.EndTimes
}

func (s *DescribeStoragePlanResponseBodyItems) GetInitCapaCityViewUnit() *string {
	return s.InitCapaCityViewUnit
}

func (s *DescribeStoragePlanResponseBodyItems) GetInitCapacityViewValue() *string {
	return s.InitCapacityViewValue
}

func (s *DescribeStoragePlanResponseBodyItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeStoragePlanResponseBodyItems) GetPeriodCapaCityViewUnit() *string {
	return s.PeriodCapaCityViewUnit
}

func (s *DescribeStoragePlanResponseBodyItems) GetPeriodCapacityViewValue() *string {
	return s.PeriodCapacityViewValue
}

func (s *DescribeStoragePlanResponseBodyItems) GetPeriodTime() *string {
	return s.PeriodTime
}

func (s *DescribeStoragePlanResponseBodyItems) GetProdCode() *string {
	return s.ProdCode
}

func (s *DescribeStoragePlanResponseBodyItems) GetPurchaseTimes() *string {
	return s.PurchaseTimes
}

func (s *DescribeStoragePlanResponseBodyItems) GetStartTimes() *string {
	return s.StartTimes
}

func (s *DescribeStoragePlanResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeStoragePlanResponseBodyItems) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeStoragePlanResponseBodyItems) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DescribeStoragePlanResponseBodyItems) SetAliUid(v string) *DescribeStoragePlanResponseBodyItems {
	s.AliUid = &v
	return s
}

func (s *DescribeStoragePlanResponseBodyItems) SetCommodityCode(v string) *DescribeStoragePlanResponseBodyItems {
	s.CommodityCode = &v
	return s
}

func (s *DescribeStoragePlanResponseBodyItems) SetEndTimes(v string) *DescribeStoragePlanResponseBodyItems {
	s.EndTimes = &v
	return s
}

func (s *DescribeStoragePlanResponseBodyItems) SetInitCapaCityViewUnit(v string) *DescribeStoragePlanResponseBodyItems {
	s.InitCapaCityViewUnit = &v
	return s
}

func (s *DescribeStoragePlanResponseBodyItems) SetInitCapacityViewValue(v string) *DescribeStoragePlanResponseBodyItems {
	s.InitCapacityViewValue = &v
	return s
}

func (s *DescribeStoragePlanResponseBodyItems) SetInstanceId(v string) *DescribeStoragePlanResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *DescribeStoragePlanResponseBodyItems) SetPeriodCapaCityViewUnit(v string) *DescribeStoragePlanResponseBodyItems {
	s.PeriodCapaCityViewUnit = &v
	return s
}

func (s *DescribeStoragePlanResponseBodyItems) SetPeriodCapacityViewValue(v string) *DescribeStoragePlanResponseBodyItems {
	s.PeriodCapacityViewValue = &v
	return s
}

func (s *DescribeStoragePlanResponseBodyItems) SetPeriodTime(v string) *DescribeStoragePlanResponseBodyItems {
	s.PeriodTime = &v
	return s
}

func (s *DescribeStoragePlanResponseBodyItems) SetProdCode(v string) *DescribeStoragePlanResponseBodyItems {
	s.ProdCode = &v
	return s
}

func (s *DescribeStoragePlanResponseBodyItems) SetPurchaseTimes(v string) *DescribeStoragePlanResponseBodyItems {
	s.PurchaseTimes = &v
	return s
}

func (s *DescribeStoragePlanResponseBodyItems) SetStartTimes(v string) *DescribeStoragePlanResponseBodyItems {
	s.StartTimes = &v
	return s
}

func (s *DescribeStoragePlanResponseBodyItems) SetStatus(v string) *DescribeStoragePlanResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeStoragePlanResponseBodyItems) SetStorageType(v string) *DescribeStoragePlanResponseBodyItems {
	s.StorageType = &v
	return s
}

func (s *DescribeStoragePlanResponseBodyItems) SetTemplateName(v string) *DescribeStoragePlanResponseBodyItems {
	s.TemplateName = &v
	return s
}

func (s *DescribeStoragePlanResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
