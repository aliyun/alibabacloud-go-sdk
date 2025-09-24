// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySavingsPlansDeductLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QuerySavingsPlansDeductLogResponseBody
	GetCode() *string
	SetData(v *QuerySavingsPlansDeductLogResponseBodyData) *QuerySavingsPlansDeductLogResponseBody
	GetData() *QuerySavingsPlansDeductLogResponseBodyData
	SetMessage(v string) *QuerySavingsPlansDeductLogResponseBody
	GetMessage() *string
	SetRequestId(v string) *QuerySavingsPlansDeductLogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QuerySavingsPlansDeductLogResponseBody
	GetSuccess() *bool
}

type QuerySavingsPlansDeductLogResponseBody struct {
	// The error code.
	//
	// example:
	//
	// PARAM_ERROR
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The return data.
	Data *QuerySavingsPlansDeductLogResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
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

func (s QuerySavingsPlansDeductLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QuerySavingsPlansDeductLogResponseBody) GoString() string {
	return s.String()
}

func (s *QuerySavingsPlansDeductLogResponseBody) GetCode() *string {
	return s.Code
}

func (s *QuerySavingsPlansDeductLogResponseBody) GetData() *QuerySavingsPlansDeductLogResponseBodyData {
	return s.Data
}

func (s *QuerySavingsPlansDeductLogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QuerySavingsPlansDeductLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QuerySavingsPlansDeductLogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QuerySavingsPlansDeductLogResponseBody) SetCode(v string) *QuerySavingsPlansDeductLogResponseBody {
	s.Code = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBody) SetData(v *QuerySavingsPlansDeductLogResponseBodyData) *QuerySavingsPlansDeductLogResponseBody {
	s.Data = v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBody) SetMessage(v string) *QuerySavingsPlansDeductLogResponseBody {
	s.Message = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBody) SetRequestId(v string) *QuerySavingsPlansDeductLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBody) SetSuccess(v bool) *QuerySavingsPlansDeductLogResponseBody {
	s.Success = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBody) Validate() error {
	return dara.Validate(s)
}

type QuerySavingsPlansDeductLogResponseBodyData struct {
	// The deduction details.
	Items []*QuerySavingsPlansDeductLogResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
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
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QuerySavingsPlansDeductLogResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QuerySavingsPlansDeductLogResponseBodyData) GoString() string {
	return s.String()
}

func (s *QuerySavingsPlansDeductLogResponseBodyData) GetItems() []*QuerySavingsPlansDeductLogResponseBodyDataItems {
	return s.Items
}

func (s *QuerySavingsPlansDeductLogResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QuerySavingsPlansDeductLogResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QuerySavingsPlansDeductLogResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QuerySavingsPlansDeductLogResponseBodyData) SetItems(v []*QuerySavingsPlansDeductLogResponseBodyDataItems) *QuerySavingsPlansDeductLogResponseBodyData {
	s.Items = v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBodyData) SetPageNum(v int32) *QuerySavingsPlansDeductLogResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBodyData) SetPageSize(v int32) *QuerySavingsPlansDeductLogResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBodyData) SetTotalCount(v int32) *QuerySavingsPlansDeductLogResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type QuerySavingsPlansDeductLogResponseBodyDataItems struct {
	// The billable item for which the fee is deducted.
	BillModule           *string `json:"BillModule,omitempty" xml:"BillModule,omitempty"`
	BillingCycle         *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	BillingOfficialPrice *string `json:"BillingOfficialPrice,omitempty" xml:"BillingOfficialPrice,omitempty"`
	// The service for which the fee is deducted.
	DeductCommodity *string `json:"DeductCommodity,omitempty" xml:"DeductCommodity,omitempty"`
	// The deducted amount.
	//
	// example:
	//
	// 0.52
	DeductFee *string `json:"DeductFee,omitempty" xml:"DeductFee,omitempty"`
	// The ID of the instance for which the fee is deducted.
	//
	// example:
	//
	// i-XXXXXXXXX
	DeductInstanceId *string `json:"DeductInstanceId,omitempty" xml:"DeductInstanceId,omitempty"`
	// The deduction rate.
	//
	// example:
	//
	// 1.0
	DeductRate            *string `json:"DeductRate,omitempty" xml:"DeductRate,omitempty"`
	DeductedOfficialPrice *string `json:"DeductedOfficialPrice,omitempty" xml:"DeductedOfficialPrice,omitempty"`
	// The discount used for the current deduction.
	//
	// example:
	//
	// 0.069
	DiscountRate *string `json:"DiscountRate,omitempty" xml:"DiscountRate,omitempty"`
	// The end of the billing cycle for which the fee is deducted.
	//
	// example:
	//
	// 2020-12-01 01:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the savings plan instance.
	//
	// example:
	//
	// spn-xxxxxxx
	InstanceId         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceSpec       *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	InstanceTypeFamily *string `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
	// example:
	//
	// 1906589291020438
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Region  *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The type of the savings plan. Valid values:
	//
	// 	- universal: general-purpose
	//
	// 	- ecs: ECS compute
	//
	// example:
	//
	// ecs
	SavingsType *string `json:"SavingsType,omitempty" xml:"SavingsType,omitempty"`
	// The beginning of the billing cycle for which the fee is deducted. The time is in the format of yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2020-12-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// 1906589291020438
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QuerySavingsPlansDeductLogResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s QuerySavingsPlansDeductLogResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) GetBillModule() *string {
	return s.BillModule
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) GetBillingOfficialPrice() *string {
	return s.BillingOfficialPrice
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) GetDeductCommodity() *string {
	return s.DeductCommodity
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) GetDeductFee() *string {
	return s.DeductFee
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) GetDeductInstanceId() *string {
	return s.DeductInstanceId
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) GetDeductRate() *string {
	return s.DeductRate
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) GetDeductedOfficialPrice() *string {
	return s.DeductedOfficialPrice
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) GetDiscountRate() *string {
	return s.DiscountRate
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) GetEndTime() *string {
	return s.EndTime
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) GetInstanceSpec() *string {
	return s.InstanceSpec
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) GetInstanceTypeFamily() *string {
	return s.InstanceTypeFamily
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) GetRegion() *string {
	return s.Region
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) GetSavingsType() *string {
	return s.SavingsType
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) GetStartTime() *string {
	return s.StartTime
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) GetUserId() *int64 {
	return s.UserId
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) SetBillModule(v string) *QuerySavingsPlansDeductLogResponseBodyDataItems {
	s.BillModule = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) SetBillingCycle(v string) *QuerySavingsPlansDeductLogResponseBodyDataItems {
	s.BillingCycle = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) SetBillingOfficialPrice(v string) *QuerySavingsPlansDeductLogResponseBodyDataItems {
	s.BillingOfficialPrice = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) SetDeductCommodity(v string) *QuerySavingsPlansDeductLogResponseBodyDataItems {
	s.DeductCommodity = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) SetDeductFee(v string) *QuerySavingsPlansDeductLogResponseBodyDataItems {
	s.DeductFee = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) SetDeductInstanceId(v string) *QuerySavingsPlansDeductLogResponseBodyDataItems {
	s.DeductInstanceId = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) SetDeductRate(v string) *QuerySavingsPlansDeductLogResponseBodyDataItems {
	s.DeductRate = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) SetDeductedOfficialPrice(v string) *QuerySavingsPlansDeductLogResponseBodyDataItems {
	s.DeductedOfficialPrice = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) SetDiscountRate(v string) *QuerySavingsPlansDeductLogResponseBodyDataItems {
	s.DiscountRate = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) SetEndTime(v string) *QuerySavingsPlansDeductLogResponseBodyDataItems {
	s.EndTime = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) SetInstanceId(v string) *QuerySavingsPlansDeductLogResponseBodyDataItems {
	s.InstanceId = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) SetInstanceSpec(v string) *QuerySavingsPlansDeductLogResponseBodyDataItems {
	s.InstanceSpec = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) SetInstanceTypeFamily(v string) *QuerySavingsPlansDeductLogResponseBodyDataItems {
	s.InstanceTypeFamily = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) SetOwnerId(v int64) *QuerySavingsPlansDeductLogResponseBodyDataItems {
	s.OwnerId = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) SetRegion(v string) *QuerySavingsPlansDeductLogResponseBodyDataItems {
	s.Region = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) SetSavingsType(v string) *QuerySavingsPlansDeductLogResponseBodyDataItems {
	s.SavingsType = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) SetStartTime(v string) *QuerySavingsPlansDeductLogResponseBodyDataItems {
	s.StartTime = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) SetUserId(v int64) *QuerySavingsPlansDeductLogResponseBodyDataItems {
	s.UserId = &v
	return s
}

func (s *QuerySavingsPlansDeductLogResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
