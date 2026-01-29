// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySavingsPlansInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QuerySavingsPlansInstanceResponseBody
	GetCode() *string
	SetData(v *QuerySavingsPlansInstanceResponseBodyData) *QuerySavingsPlansInstanceResponseBody
	GetData() *QuerySavingsPlansInstanceResponseBodyData
	SetMessage(v string) *QuerySavingsPlansInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *QuerySavingsPlansInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QuerySavingsPlansInstanceResponseBody
	GetSuccess() *bool
}

type QuerySavingsPlansInstanceResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The return data.
	Data *QuerySavingsPlansInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 61293E7A-3406-4447-8620-EC88B0AA66AD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QuerySavingsPlansInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySavingsPlansInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySavingsPlansInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuerySavingsPlansInstanceResponseBody) GetData() *QuerySavingsPlansInstanceResponseBodyData {
	return s.Data
}

func (s *QuerySavingsPlansInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuerySavingsPlansInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySavingsPlansInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QuerySavingsPlansInstanceResponseBody) SetCode(v string) *QuerySavingsPlansInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBody) SetData(v *QuerySavingsPlansInstanceResponseBodyData) *QuerySavingsPlansInstanceResponseBody {
	s.Data = v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBody) SetMessage(v string) *QuerySavingsPlansInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBody) SetRequestId(v string) *QuerySavingsPlansInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBody) SetSuccess(v bool) *QuerySavingsPlansInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QuerySavingsPlansInstanceResponseBodyData struct {
	// The details about the instances.
	Items []*QuerySavingsPlansInstanceResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QuerySavingsPlansInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QuerySavingsPlansInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySavingsPlansInstanceResponseBodyData) GetItems() []*QuerySavingsPlansInstanceResponseBodyDataItems {
	return s.Items
}

func (s *QuerySavingsPlansInstanceResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QuerySavingsPlansInstanceResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QuerySavingsPlansInstanceResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QuerySavingsPlansInstanceResponseBodyData) SetItems(v []*QuerySavingsPlansInstanceResponseBodyDataItems) *QuerySavingsPlansInstanceResponseBodyData {
	s.Items = v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyData) SetPageNum(v int32) *QuerySavingsPlansInstanceResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyData) SetPageSize(v int32) *QuerySavingsPlansInstanceResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyData) SetTotalCount(v int32) *QuerySavingsPlansInstanceResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyData) Validate() error {
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

type QuerySavingsPlansInstanceResponseBodyDataItems struct {
	// The allocation status. Valid values:
	//
	// 	- unallocated
	//
	// 	- allocated
	//
	// 	- beAllocated
	//
	// example:
	//
	// unallocated
	AllocationStatus *string `json:"AllocationStatus,omitempty" xml:"AllocationStatus,omitempty"`
	// example:
	//
	// savingplan_common_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The currency. Valid values: CNY and USD.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// 100.0
	CurrentPoolValue *string `json:"CurrentPoolValue,omitempty" xml:"CurrentPoolValue,omitempty"`
	// example:
	//
	// 1:Year
	Cycle *string `json:"Cycle,omitempty" xml:"Cycle,omitempty"`
	// example:
	//
	// HOUR
	DeductCycleType *string `json:"DeductCycleType,omitempty" xml:"DeductCycleType,omitempty"`
	// The time when the instance expires. The time is in the format of yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2021-01-01 00:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1663135741039
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The instance family information. For an instance of the Elastic Compute Service (ECS) compute type, the value indicates the ECS instance family or the ECS instance family package.
	//
	// example:
	//
	// ecs.g6
	InstanceFamily *string `json:"InstanceFamily,omitempty" xml:"InstanceFamily,omitempty"`
	// The ID of the savings plan instance.
	//
	// example:
	//
	// spn-xxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 100.0
	LastBillTotalUsage *string `json:"LastBillTotalUsage,omitempty" xml:"LastBillTotalUsage,omitempty"`
	// example:
	//
	// 0.8
	LastBillUtilization *string `json:"LastBillUtilization,omitempty" xml:"LastBillUtilization,omitempty"`
	// The payment type. Valid values:
	//
	// 	- total: All Upfront
	//
	// 	- half: Partial Upfront
	//
	// 	- zero: No Upfront
	//
	// example:
	//
	// total
	PayMode *string `json:"PayMode,omitempty" xml:"PayMode,omitempty"`
	// The commitment.
	//
	// example:
	//
	// 1.00
	PoolValue *string `json:"PoolValue,omitempty" xml:"PoolValue,omitempty"`
	// The prepaid amount.
	//
	// example:
	//
	// 8760.00
	PrepayFee *string `json:"PrepayFee,omitempty" xml:"PrepayFee,omitempty"`
	// The region.
	//
	// example:
	//
	// cn-zhangjiakou-na62-a01
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// 100.0
	RestPoolValue *string `json:"RestPoolValue,omitempty" xml:"RestPoolValue,omitempty"`
	// The type of the savings plan. Valid values:
	//
	// 	- universal: general-purpose
	//
	// 	- ecs: ECS compute
	//
	// example:
	//
	// ECS
	SavingsType *string `json:"SavingsType,omitempty" xml:"SavingsType,omitempty"`
	// The time when the instance takes effect. The time is in the format of yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2020-01-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 1663135741039
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	// The status of the instance. Valid values:
	//
	// 	- NORMAL: normal
	//
	// 	- LIMIT: stopped due to overdue payment
	//
	// 	- RELEASE: released
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The details about the tags.
	Tags []*QuerySavingsPlansInstanceResponseBodyDataItemsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The total amount that is saved.
	//
	// example:
	//
	// 20.00
	TotalSave *string `json:"TotalSave,omitempty" xml:"TotalSave,omitempty"`
	// The total usage.
	//
	// example:
	//
	// 1.0
	Utilization *string `json:"Utilization,omitempty" xml:"Utilization,omitempty"`
}

func (s QuerySavingsPlansInstanceResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s QuerySavingsPlansInstanceResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) GetAllocationStatus() *string {
	return s.AllocationStatus
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) GetCurrency() *string {
	return s.Currency
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) GetCurrentPoolValue() *string {
	return s.CurrentPoolValue
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) GetCycle() *string {
	return s.Cycle
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) GetDeductCycleType() *string {
	return s.DeductCycleType
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) GetEndTime() *string {
	return s.EndTime
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) GetEndTimestamp() *int64 {
	return s.EndTimestamp
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) GetInstanceFamily() *string {
	return s.InstanceFamily
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) GetLastBillTotalUsage() *string {
	return s.LastBillTotalUsage
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) GetLastBillUtilization() *string {
	return s.LastBillUtilization
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) GetPayMode() *string {
	return s.PayMode
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) GetPoolValue() *string {
	return s.PoolValue
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) GetPrepayFee() *string {
	return s.PrepayFee
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) GetRegion() *string {
	return s.Region
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) GetRestPoolValue() *string {
	return s.RestPoolValue
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) GetSavingsType() *string {
	return s.SavingsType
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) GetStartTime() *string {
	return s.StartTime
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) GetStartTimestamp() *int64 {
	return s.StartTimestamp
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) GetStatus() *string {
	return s.Status
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) GetTags() []*QuerySavingsPlansInstanceResponseBodyDataItemsTags {
	return s.Tags
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) GetTotalSave() *string {
	return s.TotalSave
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) GetUtilization() *string {
	return s.Utilization
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) SetAllocationStatus(v string) *QuerySavingsPlansInstanceResponseBodyDataItems {
	s.AllocationStatus = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) SetCommodityCode(v string) *QuerySavingsPlansInstanceResponseBodyDataItems {
	s.CommodityCode = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) SetCurrency(v string) *QuerySavingsPlansInstanceResponseBodyDataItems {
	s.Currency = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) SetCurrentPoolValue(v string) *QuerySavingsPlansInstanceResponseBodyDataItems {
	s.CurrentPoolValue = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) SetCycle(v string) *QuerySavingsPlansInstanceResponseBodyDataItems {
	s.Cycle = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) SetDeductCycleType(v string) *QuerySavingsPlansInstanceResponseBodyDataItems {
	s.DeductCycleType = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) SetEndTime(v string) *QuerySavingsPlansInstanceResponseBodyDataItems {
	s.EndTime = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) SetEndTimestamp(v int64) *QuerySavingsPlansInstanceResponseBodyDataItems {
	s.EndTimestamp = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) SetInstanceFamily(v string) *QuerySavingsPlansInstanceResponseBodyDataItems {
	s.InstanceFamily = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) SetInstanceId(v string) *QuerySavingsPlansInstanceResponseBodyDataItems {
	s.InstanceId = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) SetLastBillTotalUsage(v string) *QuerySavingsPlansInstanceResponseBodyDataItems {
	s.LastBillTotalUsage = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) SetLastBillUtilization(v string) *QuerySavingsPlansInstanceResponseBodyDataItems {
	s.LastBillUtilization = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) SetPayMode(v string) *QuerySavingsPlansInstanceResponseBodyDataItems {
	s.PayMode = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) SetPoolValue(v string) *QuerySavingsPlansInstanceResponseBodyDataItems {
	s.PoolValue = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) SetPrepayFee(v string) *QuerySavingsPlansInstanceResponseBodyDataItems {
	s.PrepayFee = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) SetRegion(v string) *QuerySavingsPlansInstanceResponseBodyDataItems {
	s.Region = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) SetRestPoolValue(v string) *QuerySavingsPlansInstanceResponseBodyDataItems {
	s.RestPoolValue = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) SetSavingsType(v string) *QuerySavingsPlansInstanceResponseBodyDataItems {
	s.SavingsType = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) SetStartTime(v string) *QuerySavingsPlansInstanceResponseBodyDataItems {
	s.StartTime = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) SetStartTimestamp(v int64) *QuerySavingsPlansInstanceResponseBodyDataItems {
	s.StartTimestamp = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) SetStatus(v string) *QuerySavingsPlansInstanceResponseBodyDataItems {
	s.Status = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) SetTags(v []*QuerySavingsPlansInstanceResponseBodyDataItemsTags) *QuerySavingsPlansInstanceResponseBodyDataItems {
	s.Tags = v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) SetTotalSave(v string) *QuerySavingsPlansInstanceResponseBodyDataItems {
	s.TotalSave = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) SetUtilization(v string) *QuerySavingsPlansInstanceResponseBodyDataItems {
	s.Utilization = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItems) Validate() error {
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

type QuerySavingsPlansInstanceResponseBodyDataItemsTags struct {
	// The key of the tag.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QuerySavingsPlansInstanceResponseBodyDataItemsTags) String() string {
	return dara.Prettify(s)
}

func (s QuerySavingsPlansInstanceResponseBodyDataItemsTags) GoString() string {
	return s.String()
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItemsTags) GetKey() *string {
	return s.Key
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItemsTags) GetValue() *string {
	return s.Value
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItemsTags) SetKey(v string) *QuerySavingsPlansInstanceResponseBodyDataItemsTags {
	s.Key = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItemsTags) SetValue(v string) *QuerySavingsPlansInstanceResponseBodyDataItemsTags {
	s.Value = &v
	return s
}

func (s *QuerySavingsPlansInstanceResponseBodyDataItemsTags) Validate() error {
	return dara.Validate(s)
}
