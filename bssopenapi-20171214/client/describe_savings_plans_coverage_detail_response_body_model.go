// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSavingsPlansCoverageDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSavingsPlansCoverageDetailResponseBody
	GetCode() *string
	SetData(v *DescribeSavingsPlansCoverageDetailResponseBodyData) *DescribeSavingsPlansCoverageDetailResponseBody
	GetData() *DescribeSavingsPlansCoverageDetailResponseBodyData
	SetMessage(v string) *DescribeSavingsPlansCoverageDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSavingsPlansCoverageDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeSavingsPlansCoverageDetailResponseBody
	GetSuccess() *bool
}

type DescribeSavingsPlansCoverageDetailResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The return data.
	Data *DescribeSavingsPlansCoverageDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s DescribeSavingsPlansCoverageDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlansCoverageDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlansCoverageDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSavingsPlansCoverageDetailResponseBody) GetData() *DescribeSavingsPlansCoverageDetailResponseBodyData {
	return s.Data
}

func (s *DescribeSavingsPlansCoverageDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSavingsPlansCoverageDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSavingsPlansCoverageDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeSavingsPlansCoverageDetailResponseBody) SetCode(v string) *DescribeSavingsPlansCoverageDetailResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailResponseBody) SetData(v *DescribeSavingsPlansCoverageDetailResponseBodyData) *DescribeSavingsPlansCoverageDetailResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailResponseBody) SetMessage(v string) *DescribeSavingsPlansCoverageDetailResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailResponseBody) SetRequestId(v string) *DescribeSavingsPlansCoverageDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailResponseBody) SetSuccess(v bool) *DescribeSavingsPlansCoverageDetailResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSavingsPlansCoverageDetailResponseBodyData struct {
	// The data entries.
	Items []*DescribeSavingsPlansCoverageDetailResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The token of the next page.
	//
	// example:
	//
	// eyJwYWdlTnVtIjoyLCJwYWdlU2l6ZSI6MTB9
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 400
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSavingsPlansCoverageDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlansCoverageDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlansCoverageDetailResponseBodyData) GetItems() []*DescribeSavingsPlansCoverageDetailResponseBodyDataItems {
	return s.Items
}

func (s *DescribeSavingsPlansCoverageDetailResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSavingsPlansCoverageDetailResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSavingsPlansCoverageDetailResponseBodyData) SetItems(v []*DescribeSavingsPlansCoverageDetailResponseBodyDataItems) *DescribeSavingsPlansCoverageDetailResponseBodyData {
	s.Items = v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailResponseBodyData) SetNextToken(v string) *DescribeSavingsPlansCoverageDetailResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailResponseBodyData) SetTotalCount(v int32) *DescribeSavingsPlansCoverageDetailResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeSavingsPlansCoverageDetailResponseBodyDataItems struct {
	// The coverage.
	//
	// example:
	//
	// 0.9
	CoveragePercentage *float32 `json:"CoveragePercentage,omitempty" xml:"CoveragePercentage,omitempty"`
	// The currency.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The deducted amount.
	//
	// example:
	//
	// 100
	DeductAmount *float32 `json:"DeductAmount,omitempty" xml:"DeductAmount,omitempty"`
	// The end time.
	//
	// example:
	//
	// 2021-05-05 15:00:00
	EndPeriod *string `json:"EndPeriod,omitempty" xml:"EndPeriod,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// spn-a1fhs54c243hP22
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The specifications.
	//
	// example:
	//
	// 7th_generation_X86_group
	InstanceSpec *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	// example:
	//
	// 1906589291020438
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The pay-as-you-go cost.
	//
	// example:
	//
	// 200
	PostpaidCost *float32 `json:"PostpaidCost,omitempty" xml:"PostpaidCost,omitempty"`
	// The region.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The start time.
	//
	// example:
	//
	// 2021-05-05 14:00:00
	StartPeriod *string `json:"StartPeriod,omitempty" xml:"StartPeriod,omitempty"`
	// The total expenditure.
	//
	// example:
	//
	// 100
	TotalAmount *float32 `json:"TotalAmount,omitempty" xml:"TotalAmount,omitempty"`
	// The ID of the account.
	//
	// example:
	//
	// 2831685687844416
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username of the account.
	//
	// example:
	//
	// tester1@test.aliyun.com
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeSavingsPlansCoverageDetailResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeSavingsPlansCoverageDetailResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *DescribeSavingsPlansCoverageDetailResponseBodyDataItems) GetCoveragePercentage() *float32 {
	return s.CoveragePercentage
}

func (s *DescribeSavingsPlansCoverageDetailResponseBodyDataItems) GetCurrency() *string {
	return s.Currency
}

func (s *DescribeSavingsPlansCoverageDetailResponseBodyDataItems) GetDeductAmount() *float32 {
	return s.DeductAmount
}

func (s *DescribeSavingsPlansCoverageDetailResponseBodyDataItems) GetEndPeriod() *string {
	return s.EndPeriod
}

func (s *DescribeSavingsPlansCoverageDetailResponseBodyDataItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSavingsPlansCoverageDetailResponseBodyDataItems) GetInstanceSpec() *string {
	return s.InstanceSpec
}

func (s *DescribeSavingsPlansCoverageDetailResponseBodyDataItems) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSavingsPlansCoverageDetailResponseBodyDataItems) GetPostpaidCost() *float32 {
	return s.PostpaidCost
}

func (s *DescribeSavingsPlansCoverageDetailResponseBodyDataItems) GetRegion() *string {
	return s.Region
}

func (s *DescribeSavingsPlansCoverageDetailResponseBodyDataItems) GetStartPeriod() *string {
	return s.StartPeriod
}

func (s *DescribeSavingsPlansCoverageDetailResponseBodyDataItems) GetTotalAmount() *float32 {
	return s.TotalAmount
}

func (s *DescribeSavingsPlansCoverageDetailResponseBodyDataItems) GetUserId() *int64 {
	return s.UserId
}

func (s *DescribeSavingsPlansCoverageDetailResponseBodyDataItems) GetUserName() *string {
	return s.UserName
}

func (s *DescribeSavingsPlansCoverageDetailResponseBodyDataItems) SetCoveragePercentage(v float32) *DescribeSavingsPlansCoverageDetailResponseBodyDataItems {
	s.CoveragePercentage = &v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailResponseBodyDataItems) SetCurrency(v string) *DescribeSavingsPlansCoverageDetailResponseBodyDataItems {
	s.Currency = &v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailResponseBodyDataItems) SetDeductAmount(v float32) *DescribeSavingsPlansCoverageDetailResponseBodyDataItems {
	s.DeductAmount = &v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailResponseBodyDataItems) SetEndPeriod(v string) *DescribeSavingsPlansCoverageDetailResponseBodyDataItems {
	s.EndPeriod = &v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailResponseBodyDataItems) SetInstanceId(v string) *DescribeSavingsPlansCoverageDetailResponseBodyDataItems {
	s.InstanceId = &v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailResponseBodyDataItems) SetInstanceSpec(v string) *DescribeSavingsPlansCoverageDetailResponseBodyDataItems {
	s.InstanceSpec = &v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailResponseBodyDataItems) SetOwnerId(v int64) *DescribeSavingsPlansCoverageDetailResponseBodyDataItems {
	s.OwnerId = &v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailResponseBodyDataItems) SetPostpaidCost(v float32) *DescribeSavingsPlansCoverageDetailResponseBodyDataItems {
	s.PostpaidCost = &v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailResponseBodyDataItems) SetRegion(v string) *DescribeSavingsPlansCoverageDetailResponseBodyDataItems {
	s.Region = &v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailResponseBodyDataItems) SetStartPeriod(v string) *DescribeSavingsPlansCoverageDetailResponseBodyDataItems {
	s.StartPeriod = &v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailResponseBodyDataItems) SetTotalAmount(v float32) *DescribeSavingsPlansCoverageDetailResponseBodyDataItems {
	s.TotalAmount = &v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailResponseBodyDataItems) SetUserId(v int64) *DescribeSavingsPlansCoverageDetailResponseBodyDataItems {
	s.UserId = &v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailResponseBodyDataItems) SetUserName(v string) *DescribeSavingsPlansCoverageDetailResponseBodyDataItems {
	s.UserName = &v
	return s
}

func (s *DescribeSavingsPlansCoverageDetailResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
