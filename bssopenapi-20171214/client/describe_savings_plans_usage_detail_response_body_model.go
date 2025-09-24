// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSavingsPlansUsageDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSavingsPlansUsageDetailResponseBody
	GetCode() *string
	SetData(v *DescribeSavingsPlansUsageDetailResponseBodyData) *DescribeSavingsPlansUsageDetailResponseBody
	GetData() *DescribeSavingsPlansUsageDetailResponseBodyData
	SetMessage(v string) *DescribeSavingsPlansUsageDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSavingsPlansUsageDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeSavingsPlansUsageDetailResponseBody
	GetSuccess() *bool
}

type DescribeSavingsPlansUsageDetailResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The return data.
	Data *DescribeSavingsPlansUsageDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 79EE7556-0CFD-44EB-9CD6-B3B526E3A85F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSavingsPlansUsageDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlansUsageDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlansUsageDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSavingsPlansUsageDetailResponseBody) GetData() *DescribeSavingsPlansUsageDetailResponseBodyData {
	return s.Data
}

func (s *DescribeSavingsPlansUsageDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSavingsPlansUsageDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSavingsPlansUsageDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeSavingsPlansUsageDetailResponseBody) SetCode(v string) *DescribeSavingsPlansUsageDetailResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailResponseBody) SetData(v *DescribeSavingsPlansUsageDetailResponseBodyData) *DescribeSavingsPlansUsageDetailResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSavingsPlansUsageDetailResponseBody) SetMessage(v string) *DescribeSavingsPlansUsageDetailResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailResponseBody) SetRequestId(v string) *DescribeSavingsPlansUsageDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailResponseBody) SetSuccess(v bool) *DescribeSavingsPlansUsageDetailResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSavingsPlansUsageDetailResponseBodyData struct {
	// The data entries.
	Items []*DescribeSavingsPlansUsageDetailResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The token of the next page.
	//
	// example:
	//
	// eyJwYWdlTnVtIjoyLCJwYWdlU2ldhsSI6MTB9
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1000
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSavingsPlansUsageDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlansUsageDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlansUsageDetailResponseBodyData) GetItems() []*DescribeSavingsPlansUsageDetailResponseBodyDataItems {
	return s.Items
}

func (s *DescribeSavingsPlansUsageDetailResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSavingsPlansUsageDetailResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSavingsPlansUsageDetailResponseBodyData) SetItems(v []*DescribeSavingsPlansUsageDetailResponseBodyDataItems) *DescribeSavingsPlansUsageDetailResponseBodyData {
	s.Items = v
	return s
}

func (s *DescribeSavingsPlansUsageDetailResponseBodyData) SetNextToken(v string) *DescribeSavingsPlansUsageDetailResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailResponseBodyData) SetTotalCount(v int32) *DescribeSavingsPlansUsageDetailResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeSavingsPlansUsageDetailResponseBodyDataItems struct {
	// The currency.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The used amount of the savings plan.
	//
	// example:
	//
	// 3.94
	DeductValue *float32 `json:"DeductValue,omitempty" xml:"DeductValue,omitempty"`
	// The end time.
	//
	// example:
	//
	// 2021-08-09 00:00:00
	EndPeriod *string `json:"EndPeriod,omitempty" xml:"EndPeriod,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// spn-a1fhs54c243hP22
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The total amount of the savings plan.
	//
	// example:
	//
	// 29.84
	PoolValue *float32 `json:"PoolValue,omitempty" xml:"PoolValue,omitempty"`
	// The pay-as-you-go cost.
	//
	// example:
	//
	// 4.2
	PostpaidCost *float32 `json:"PostpaidCost,omitempty" xml:"PostpaidCost,omitempty"`
	// The amount that is saved.
	//
	// example:
	//
	// 0.08
	SavedCost *float32 `json:"SavedCost,omitempty" xml:"SavedCost,omitempty"`
	// The start time.
	//
	// example:
	//
	// 2021-08-01 00:00:00
	StartPeriod *string `json:"StartPeriod,omitempty" xml:"StartPeriod,omitempty"`
	// The status of the instance.
	//
	// A value of -1 indicates that the payment is overdue. A value of 1 indicates that the instance is active.
	//
	// example:
	//
	// -1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the savings plan. Valid values: universal and ECS compute.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The usage.
	//
	// example:
	//
	// 0.9
	UsagePercentage *float32 `json:"UsagePercentage,omitempty" xml:"UsagePercentage,omitempty"`
	// The ID of the account.
	//
	// example:
	//
	// 123745698925000
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username of the account.
	//
	// example:
	//
	// test13@test.aliyun.com
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeSavingsPlansUsageDetailResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlansUsageDetailResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlansUsageDetailResponseBodyDataItems) GetCurrency() *string {
	return s.Currency
}

func (s *DescribeSavingsPlansUsageDetailResponseBodyDataItems) GetDeductValue() *float32 {
	return s.DeductValue
}

func (s *DescribeSavingsPlansUsageDetailResponseBodyDataItems) GetEndPeriod() *string {
	return s.EndPeriod
}

func (s *DescribeSavingsPlansUsageDetailResponseBodyDataItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSavingsPlansUsageDetailResponseBodyDataItems) GetPoolValue() *float32 {
	return s.PoolValue
}

func (s *DescribeSavingsPlansUsageDetailResponseBodyDataItems) GetPostpaidCost() *float32 {
	return s.PostpaidCost
}

func (s *DescribeSavingsPlansUsageDetailResponseBodyDataItems) GetSavedCost() *float32 {
	return s.SavedCost
}

func (s *DescribeSavingsPlansUsageDetailResponseBodyDataItems) GetStartPeriod() *string {
	return s.StartPeriod
}

func (s *DescribeSavingsPlansUsageDetailResponseBodyDataItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeSavingsPlansUsageDetailResponseBodyDataItems) GetType() *string {
	return s.Type
}

func (s *DescribeSavingsPlansUsageDetailResponseBodyDataItems) GetUsagePercentage() *float32 {
	return s.UsagePercentage
}

func (s *DescribeSavingsPlansUsageDetailResponseBodyDataItems) GetUserId() *int64 {
	return s.UserId
}

func (s *DescribeSavingsPlansUsageDetailResponseBodyDataItems) GetUserName() *string {
	return s.UserName
}

func (s *DescribeSavingsPlansUsageDetailResponseBodyDataItems) SetCurrency(v string) *DescribeSavingsPlansUsageDetailResponseBodyDataItems {
	s.Currency = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailResponseBodyDataItems) SetDeductValue(v float32) *DescribeSavingsPlansUsageDetailResponseBodyDataItems {
	s.DeductValue = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailResponseBodyDataItems) SetEndPeriod(v string) *DescribeSavingsPlansUsageDetailResponseBodyDataItems {
	s.EndPeriod = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailResponseBodyDataItems) SetInstanceId(v string) *DescribeSavingsPlansUsageDetailResponseBodyDataItems {
	s.InstanceId = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailResponseBodyDataItems) SetPoolValue(v float32) *DescribeSavingsPlansUsageDetailResponseBodyDataItems {
	s.PoolValue = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailResponseBodyDataItems) SetPostpaidCost(v float32) *DescribeSavingsPlansUsageDetailResponseBodyDataItems {
	s.PostpaidCost = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailResponseBodyDataItems) SetSavedCost(v float32) *DescribeSavingsPlansUsageDetailResponseBodyDataItems {
	s.SavedCost = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailResponseBodyDataItems) SetStartPeriod(v string) *DescribeSavingsPlansUsageDetailResponseBodyDataItems {
	s.StartPeriod = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailResponseBodyDataItems) SetStatus(v string) *DescribeSavingsPlansUsageDetailResponseBodyDataItems {
	s.Status = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailResponseBodyDataItems) SetType(v string) *DescribeSavingsPlansUsageDetailResponseBodyDataItems {
	s.Type = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailResponseBodyDataItems) SetUsagePercentage(v float32) *DescribeSavingsPlansUsageDetailResponseBodyDataItems {
	s.UsagePercentage = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailResponseBodyDataItems) SetUserId(v int64) *DescribeSavingsPlansUsageDetailResponseBodyDataItems {
	s.UserId = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailResponseBodyDataItems) SetUserName(v string) *DescribeSavingsPlansUsageDetailResponseBodyDataItems {
	s.UserName = &v
	return s
}

func (s *DescribeSavingsPlansUsageDetailResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
