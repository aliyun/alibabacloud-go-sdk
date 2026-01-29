// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceCoverageDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeResourceCoverageDetailResponseBody
	GetCode() *string
	SetData(v *DescribeResourceCoverageDetailResponseBodyData) *DescribeResourceCoverageDetailResponseBody
	GetData() *DescribeResourceCoverageDetailResponseBodyData
	SetMessage(v string) *DescribeResourceCoverageDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeResourceCoverageDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeResourceCoverageDetailResponseBody
	GetSuccess() *bool
}

type DescribeResourceCoverageDetailResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *DescribeResourceCoverageDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s DescribeResourceCoverageDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceCoverageDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceCoverageDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeResourceCoverageDetailResponseBody) GetData() *DescribeResourceCoverageDetailResponseBodyData {
	return s.Data
}

func (s *DescribeResourceCoverageDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeResourceCoverageDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeResourceCoverageDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeResourceCoverageDetailResponseBody) SetCode(v string) *DescribeResourceCoverageDetailResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeResourceCoverageDetailResponseBody) SetData(v *DescribeResourceCoverageDetailResponseBodyData) *DescribeResourceCoverageDetailResponseBody {
	s.Data = v
	return s
}

func (s *DescribeResourceCoverageDetailResponseBody) SetMessage(v string) *DescribeResourceCoverageDetailResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeResourceCoverageDetailResponseBody) SetRequestId(v string) *DescribeResourceCoverageDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceCoverageDetailResponseBody) SetSuccess(v bool) *DescribeResourceCoverageDetailResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeResourceCoverageDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeResourceCoverageDetailResponseBodyData struct {
	// The data entries.
	Items []*DescribeResourceCoverageDetailResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 200
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token of the next page.
	//
	// example:
	//
	// eyJwYWdlTnVtIjoyLCJwYWdlU2l6ZSI6MTB9
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10000
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeResourceCoverageDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceCoverageDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeResourceCoverageDetailResponseBodyData) GetItems() []*DescribeResourceCoverageDetailResponseBodyDataItems {
	return s.Items
}

func (s *DescribeResourceCoverageDetailResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeResourceCoverageDetailResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeResourceCoverageDetailResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeResourceCoverageDetailResponseBodyData) SetItems(v []*DescribeResourceCoverageDetailResponseBodyDataItems) *DescribeResourceCoverageDetailResponseBodyData {
	s.Items = v
	return s
}

func (s *DescribeResourceCoverageDetailResponseBodyData) SetMaxResults(v int32) *DescribeResourceCoverageDetailResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *DescribeResourceCoverageDetailResponseBodyData) SetNextToken(v string) *DescribeResourceCoverageDetailResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *DescribeResourceCoverageDetailResponseBodyData) SetTotalCount(v int32) *DescribeResourceCoverageDetailResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeResourceCoverageDetailResponseBodyData) Validate() error {
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

type DescribeResourceCoverageDetailResponseBodyDataItems struct {
	// The unit that is used to measure the resources deducted from deduction plans.
	CapacityUnit *string `json:"CapacityUnit,omitempty" xml:"CapacityUnit,omitempty"`
	// The code of the service.
	//
	// example:
	//
	// ecs
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The name and billing method of the service.
	CommodityName *string `json:"CommodityName,omitempty" xml:"CommodityName,omitempty"`
	// The coverage rate of a deduction plan.
	//
	// example:
	//
	// 1
	CoveragePercentage *float32 `json:"CoveragePercentage,omitempty" xml:"CoveragePercentage,omitempty"`
	// The currency in which deduction plans were priced.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The amount of the resources deducted from a deduction plan.
	//
	// example:
	//
	// 1
	DeductQuantity *float32 `json:"DeductQuantity,omitempty" xml:"DeductQuantity,omitempty"`
	// The end of the time range in which the coverage details were queried.
	//
	// example:
	//
	// 2021-04-01 01:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of a pay-as-you-go instance.
	//
	// example:
	//
	// i-bp1d9x623987rlj0dx4xx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The specifications of a deduction plan.
	//
	// example:
	//
	// ecs.t5-lc2m1.nano
	InstanceSpec *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	// The amount of the bill.
	//
	// example:
	//
	// 0
	PaymentAmount *float32 `json:"PaymentAmount,omitempty" xml:"PaymentAmount,omitempty"`
	// The code of the service.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The name of the service.
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The region.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The code of the region.
	//
	// example:
	//
	// cn-hangzhou-dg-a01
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The beginning of the time range in which the coverage details were queried.
	//
	// example:
	//
	// 2021-04-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The total amount of resources consumed.
	//
	// example:
	//
	// 1
	TotalQuantity *float32 `json:"TotalQuantity,omitempty" xml:"TotalQuantity,omitempty"`
	// The ID of the account.
	//
	// example:
	//
	// 123745698925000
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username of the account.
	//
	// example:
	//
	// test13@test.aliyun.com
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The code of the zone.
	//
	// example:
	//
	// cn-hangzhou-i
	Zone *string `json:"Zone,omitempty" xml:"Zone,omitempty"`
	// The zone.
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s DescribeResourceCoverageDetailResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceCoverageDetailResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) GetCapacityUnit() *string {
	return s.CapacityUnit
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) GetCommodityName() *string {
	return s.CommodityName
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) GetCoveragePercentage() *float32 {
	return s.CoveragePercentage
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) GetCurrency() *string {
	return s.Currency
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) GetDeductQuantity() *float32 {
	return s.DeductQuantity
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) GetInstanceSpec() *string {
	return s.InstanceSpec
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) GetPaymentAmount() *float32 {
	return s.PaymentAmount
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) GetProductName() *string {
	return s.ProductName
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) GetRegion() *string {
	return s.Region
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) GetTotalQuantity() *float32 {
	return s.TotalQuantity
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) GetUserId() *string {
	return s.UserId
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) GetUserName() *string {
	return s.UserName
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) GetZone() *string {
	return s.Zone
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) GetZoneName() *string {
	return s.ZoneName
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) SetCapacityUnit(v string) *DescribeResourceCoverageDetailResponseBodyDataItems {
	s.CapacityUnit = &v
	return s
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) SetCommodityCode(v string) *DescribeResourceCoverageDetailResponseBodyDataItems {
	s.CommodityCode = &v
	return s
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) SetCommodityName(v string) *DescribeResourceCoverageDetailResponseBodyDataItems {
	s.CommodityName = &v
	return s
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) SetCoveragePercentage(v float32) *DescribeResourceCoverageDetailResponseBodyDataItems {
	s.CoveragePercentage = &v
	return s
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) SetCurrency(v string) *DescribeResourceCoverageDetailResponseBodyDataItems {
	s.Currency = &v
	return s
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) SetDeductQuantity(v float32) *DescribeResourceCoverageDetailResponseBodyDataItems {
	s.DeductQuantity = &v
	return s
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) SetEndTime(v string) *DescribeResourceCoverageDetailResponseBodyDataItems {
	s.EndTime = &v
	return s
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) SetInstanceId(v string) *DescribeResourceCoverageDetailResponseBodyDataItems {
	s.InstanceId = &v
	return s
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) SetInstanceSpec(v string) *DescribeResourceCoverageDetailResponseBodyDataItems {
	s.InstanceSpec = &v
	return s
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) SetPaymentAmount(v float32) *DescribeResourceCoverageDetailResponseBodyDataItems {
	s.PaymentAmount = &v
	return s
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) SetProductCode(v string) *DescribeResourceCoverageDetailResponseBodyDataItems {
	s.ProductCode = &v
	return s
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) SetProductName(v string) *DescribeResourceCoverageDetailResponseBodyDataItems {
	s.ProductName = &v
	return s
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) SetRegion(v string) *DescribeResourceCoverageDetailResponseBodyDataItems {
	s.Region = &v
	return s
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) SetRegionNo(v string) *DescribeResourceCoverageDetailResponseBodyDataItems {
	s.RegionNo = &v
	return s
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) SetStartTime(v string) *DescribeResourceCoverageDetailResponseBodyDataItems {
	s.StartTime = &v
	return s
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) SetTotalQuantity(v float32) *DescribeResourceCoverageDetailResponseBodyDataItems {
	s.TotalQuantity = &v
	return s
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) SetUserId(v string) *DescribeResourceCoverageDetailResponseBodyDataItems {
	s.UserId = &v
	return s
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) SetUserName(v string) *DescribeResourceCoverageDetailResponseBodyDataItems {
	s.UserName = &v
	return s
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) SetZone(v string) *DescribeResourceCoverageDetailResponseBodyDataItems {
	s.Zone = &v
	return s
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) SetZoneName(v string) *DescribeResourceCoverageDetailResponseBodyDataItems {
	s.ZoneName = &v
	return s
}

func (s *DescribeResourceCoverageDetailResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
