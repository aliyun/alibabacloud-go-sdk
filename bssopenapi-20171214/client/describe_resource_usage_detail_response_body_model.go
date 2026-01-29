// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceUsageDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeResourceUsageDetailResponseBody
	GetCode() *string
	SetData(v *DescribeResourceUsageDetailResponseBodyData) *DescribeResourceUsageDetailResponseBody
	GetData() *DescribeResourceUsageDetailResponseBodyData
	SetMessage(v string) *DescribeResourceUsageDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeResourceUsageDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeResourceUsageDetailResponseBody
	GetSuccess() *bool
}

type DescribeResourceUsageDetailResponseBody struct {
	// The response code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *DescribeResourceUsageDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 79EE7556-0CFD-44EB-9CD6-B3B526E3A85F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeResourceUsageDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceUsageDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceUsageDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeResourceUsageDetailResponseBody) GetData() *DescribeResourceUsageDetailResponseBodyData {
	return s.Data
}

func (s *DescribeResourceUsageDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeResourceUsageDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeResourceUsageDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeResourceUsageDetailResponseBody) SetCode(v string) *DescribeResourceUsageDetailResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeResourceUsageDetailResponseBody) SetData(v *DescribeResourceUsageDetailResponseBodyData) *DescribeResourceUsageDetailResponseBody {
	s.Data = v
	return s
}

func (s *DescribeResourceUsageDetailResponseBody) SetMessage(v string) *DescribeResourceUsageDetailResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeResourceUsageDetailResponseBody) SetRequestId(v string) *DescribeResourceUsageDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceUsageDetailResponseBody) SetSuccess(v bool) *DescribeResourceUsageDetailResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeResourceUsageDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeResourceUsageDetailResponseBodyData struct {
	// The data entries.
	Items []*DescribeResourceUsageDetailResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 200
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
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

func (s DescribeResourceUsageDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceUsageDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeResourceUsageDetailResponseBodyData) GetItems() []*DescribeResourceUsageDetailResponseBodyDataItems {
	return s.Items
}

func (s *DescribeResourceUsageDetailResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeResourceUsageDetailResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeResourceUsageDetailResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeResourceUsageDetailResponseBodyData) SetItems(v []*DescribeResourceUsageDetailResponseBodyDataItems) *DescribeResourceUsageDetailResponseBodyData {
	s.Items = v
	return s
}

func (s *DescribeResourceUsageDetailResponseBodyData) SetMaxResults(v int32) *DescribeResourceUsageDetailResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *DescribeResourceUsageDetailResponseBodyData) SetNextToken(v string) *DescribeResourceUsageDetailResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *DescribeResourceUsageDetailResponseBodyData) SetTotalCount(v int32) *DescribeResourceUsageDetailResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeResourceUsageDetailResponseBodyData) Validate() error {
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

type DescribeResourceUsageDetailResponseBodyDataItems struct {
	// The unit that is used to measure the resources that are deducted.
	CapacityUnit *string `json:"CapacityUnit,omitempty" xml:"CapacityUnit,omitempty"`
	// The type of the currency.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The amount of the deducted resources.
	//
	// example:
	//
	// 1
	DeductQuantity *float32 `json:"DeductQuantity,omitempty" xml:"DeductQuantity,omitempty"`
	// The end of the time range in which the usage details were queried.
	//
	// example:
	//
	// 2021-04-01 01:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The operating system.
	//
	// example:
	//
	// linux
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// The instance type.
	//
	// example:
	//
	// ecs.t5-lc2m1.nano
	InstanceSpec *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty"`
	// The equivalent of pay-as-you-go costs.
	//
	// example:
	//
	// 0.06
	PostpaidCost *string `json:"PostpaidCost,omitempty" xml:"PostpaidCost,omitempty"`
	// The potential net savings.
	//
	// example:
	//
	// 0.13
	PotentialSavedCost *string `json:"PotentialSavedCost,omitempty" xml:"PotentialSavedCost,omitempty"`
	// The number of deduction plans.
	//
	// example:
	//
	// 2
	Quantity *int64 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// The region.
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The code of the region.
	//
	// example:
	//
	// cn-hangzhou-dg-a01
	RegionNo *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty"`
	// The fee of the deduction plan.
	//
	// example:
	//
	// 0
	ReservationCost *string `json:"ReservationCost,omitempty" xml:"ReservationCost,omitempty"`
	// The ID of the deduction plan.
	//
	// example:
	//
	// ecsri-bp147nnfz21225k9mpix00
	ResourceInstanceId *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
	// The net savings.
	//
	// example:
	//
	// 0.06
	SavedCost *string `json:"SavedCost,omitempty" xml:"SavedCost,omitempty"`
	// The beginning of the time range in which the usage details were queried.
	//
	// example:
	//
	// 2021-04-01 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the deduction plan.
	//
	// example:
	//
	// Valid
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the status.
	StatusName *string `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
	// The total capacity of the deduction plan.
	//
	// example:
	//
	// 2
	TotalQuantity *float32 `json:"TotalQuantity,omitempty" xml:"TotalQuantity,omitempty"`
	// The usage rate of the deduction plan.
	//
	// example:
	//
	// 0.5
	UsagePercentage *float32 `json:"UsagePercentage,omitempty" xml:"UsagePercentage,omitempty"`
	// The account ID.
	//
	// example:
	//
	// 123745698925000
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username of the account.
	//
	// example:
	//
	// test@aliyun.com
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The zone.
	//
	// example:
	//
	// cn-hangzhou-i
	Zone *string `json:"Zone,omitempty" xml:"Zone,omitempty"`
	// The code of the zone.
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s DescribeResourceUsageDetailResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceUsageDetailResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) GetCapacityUnit() *string {
	return s.CapacityUnit
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) GetCurrency() *string {
	return s.Currency
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) GetDeductQuantity() *float32 {
	return s.DeductQuantity
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) GetImageType() *string {
	return s.ImageType
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) GetInstanceSpec() *string {
	return s.InstanceSpec
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) GetPostpaidCost() *string {
	return s.PostpaidCost
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) GetPotentialSavedCost() *string {
	return s.PotentialSavedCost
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) GetQuantity() *int64 {
	return s.Quantity
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) GetRegion() *string {
	return s.Region
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) GetRegionNo() *string {
	return s.RegionNo
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) GetReservationCost() *string {
	return s.ReservationCost
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) GetResourceInstanceId() *string {
	return s.ResourceInstanceId
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) GetSavedCost() *string {
	return s.SavedCost
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) GetStatusName() *string {
	return s.StatusName
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) GetTotalQuantity() *float32 {
	return s.TotalQuantity
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) GetUsagePercentage() *float32 {
	return s.UsagePercentage
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) GetUserId() *string {
	return s.UserId
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) GetUserName() *string {
	return s.UserName
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) GetZone() *string {
	return s.Zone
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) GetZoneName() *string {
	return s.ZoneName
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) SetCapacityUnit(v string) *DescribeResourceUsageDetailResponseBodyDataItems {
	s.CapacityUnit = &v
	return s
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) SetCurrency(v string) *DescribeResourceUsageDetailResponseBodyDataItems {
	s.Currency = &v
	return s
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) SetDeductQuantity(v float32) *DescribeResourceUsageDetailResponseBodyDataItems {
	s.DeductQuantity = &v
	return s
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) SetEndTime(v string) *DescribeResourceUsageDetailResponseBodyDataItems {
	s.EndTime = &v
	return s
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) SetImageType(v string) *DescribeResourceUsageDetailResponseBodyDataItems {
	s.ImageType = &v
	return s
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) SetInstanceSpec(v string) *DescribeResourceUsageDetailResponseBodyDataItems {
	s.InstanceSpec = &v
	return s
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) SetPostpaidCost(v string) *DescribeResourceUsageDetailResponseBodyDataItems {
	s.PostpaidCost = &v
	return s
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) SetPotentialSavedCost(v string) *DescribeResourceUsageDetailResponseBodyDataItems {
	s.PotentialSavedCost = &v
	return s
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) SetQuantity(v int64) *DescribeResourceUsageDetailResponseBodyDataItems {
	s.Quantity = &v
	return s
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) SetRegion(v string) *DescribeResourceUsageDetailResponseBodyDataItems {
	s.Region = &v
	return s
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) SetRegionNo(v string) *DescribeResourceUsageDetailResponseBodyDataItems {
	s.RegionNo = &v
	return s
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) SetReservationCost(v string) *DescribeResourceUsageDetailResponseBodyDataItems {
	s.ReservationCost = &v
	return s
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) SetResourceInstanceId(v string) *DescribeResourceUsageDetailResponseBodyDataItems {
	s.ResourceInstanceId = &v
	return s
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) SetSavedCost(v string) *DescribeResourceUsageDetailResponseBodyDataItems {
	s.SavedCost = &v
	return s
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) SetStartTime(v string) *DescribeResourceUsageDetailResponseBodyDataItems {
	s.StartTime = &v
	return s
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) SetStatus(v string) *DescribeResourceUsageDetailResponseBodyDataItems {
	s.Status = &v
	return s
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) SetStatusName(v string) *DescribeResourceUsageDetailResponseBodyDataItems {
	s.StatusName = &v
	return s
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) SetTotalQuantity(v float32) *DescribeResourceUsageDetailResponseBodyDataItems {
	s.TotalQuantity = &v
	return s
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) SetUsagePercentage(v float32) *DescribeResourceUsageDetailResponseBodyDataItems {
	s.UsagePercentage = &v
	return s
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) SetUserId(v string) *DescribeResourceUsageDetailResponseBodyDataItems {
	s.UserId = &v
	return s
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) SetUserName(v string) *DescribeResourceUsageDetailResponseBodyDataItems {
	s.UserName = &v
	return s
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) SetZone(v string) *DescribeResourceUsageDetailResponseBodyDataItems {
	s.Zone = &v
	return s
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) SetZoneName(v string) *DescribeResourceUsageDetailResponseBodyDataItems {
	s.ZoneName = &v
	return s
}

func (s *DescribeResourceUsageDetailResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
