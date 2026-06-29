// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubscriptionSeatDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSubscriptionSeatDetailsResponseBody
	GetCode() *string
	SetData(v *GetSubscriptionSeatDetailsResponseBodyData) *GetSubscriptionSeatDetailsResponseBody
	GetData() *GetSubscriptionSeatDetailsResponseBodyData
	SetMessage(v string) *GetSubscriptionSeatDetailsResponseBody
	GetMessage() *string
	SetSuccess(v bool) *GetSubscriptionSeatDetailsResponseBody
	GetSuccess() *bool
}

type GetSubscriptionSeatDetailsResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data.
	Data *GetSubscriptionSeatDetailsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// - true: Successful.
	//
	// - false: Failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSubscriptionSeatDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSubscriptionSeatDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *GetSubscriptionSeatDetailsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSubscriptionSeatDetailsResponseBody) GetData() *GetSubscriptionSeatDetailsResponseBodyData {
	return s.Data
}

func (s *GetSubscriptionSeatDetailsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSubscriptionSeatDetailsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSubscriptionSeatDetailsResponseBody) SetCode(v string) *GetSubscriptionSeatDetailsResponseBody {
	s.Code = &v
	return s
}

func (s *GetSubscriptionSeatDetailsResponseBody) SetData(v *GetSubscriptionSeatDetailsResponseBodyData) *GetSubscriptionSeatDetailsResponseBody {
	s.Data = v
	return s
}

func (s *GetSubscriptionSeatDetailsResponseBody) SetMessage(v string) *GetSubscriptionSeatDetailsResponseBody {
	s.Message = &v
	return s
}

func (s *GetSubscriptionSeatDetailsResponseBody) SetSuccess(v bool) *GetSubscriptionSeatDetailsResponseBody {
	s.Success = &v
	return s
}

func (s *GetSubscriptionSeatDetailsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSubscriptionSeatDetailsResponseBodyData struct {
	// The data items.
	Items []*GetSubscriptionSeatDetailsResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number. The value is greater than 0 and does not exceed the maximum value of the Integer data type.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of seats.
	//
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSubscriptionSeatDetailsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSubscriptionSeatDetailsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSubscriptionSeatDetailsResponseBodyData) GetItems() []*GetSubscriptionSeatDetailsResponseBodyDataItems {
	return s.Items
}

func (s *GetSubscriptionSeatDetailsResponseBodyData) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetSubscriptionSeatDetailsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSubscriptionSeatDetailsResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *GetSubscriptionSeatDetailsResponseBodyData) SetItems(v []*GetSubscriptionSeatDetailsResponseBodyDataItems) *GetSubscriptionSeatDetailsResponseBodyData {
	s.Items = v
	return s
}

func (s *GetSubscriptionSeatDetailsResponseBodyData) SetPageNo(v int32) *GetSubscriptionSeatDetailsResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetSubscriptionSeatDetailsResponseBodyData) SetPageSize(v int32) *GetSubscriptionSeatDetailsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetSubscriptionSeatDetailsResponseBodyData) SetTotal(v int32) *GetSubscriptionSeatDetailsResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetSubscriptionSeatDetailsResponseBodyData) Validate() error {
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

type GetSubscriptionSeatDetailsResponseBodyDataItems struct {
	// The mailbox of the member accounts.
	//
	// example:
	//
	// test@email.com
	AccountEmail *string `json:"AccountEmail,omitempty" xml:"AccountEmail,omitempty"`
	// The ID of the attached member accounts.
	//
	// example:
	//
	// acc_123456789
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The name of the member accounts.
	//
	// example:
	//
	// ExampleName
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The assignment status. Valid values:
	//
	// - ASSIGNED
	//
	// - UNASSIGNED
	//
	// example:
	//
	// ASSIGNED
	AssignedStatus *string `json:"AssignedStatus,omitempty" xml:"AssignedStatus,omitempty"`
	// The expiration time of the seat.
	//
	// example:
	//
	// 1781422733
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The currently active equity instances. For TokenPlan products, this list contains only one active equity instance.
	EquityList []*GetSubscriptionSeatDetailsResponseBodyDataItemsEquityList `json:"EquityList,omitempty" xml:"EquityList,omitempty" type:"Repeated"`
	// The instance code of the seat.
	//
	// example:
	//
	// subs-1234567
	InstanceCode *string `json:"InstanceCode,omitempty" xml:"InstanceCode,omitempty"`
	// The seat ID.
	//
	// example:
	//
	// seat_123456
	SeatId *string `json:"SeatId,omitempty" xml:"SeatId,omitempty"`
	// The seat type. Valid values:
	//
	// - standard: standard seat.
	//
	// - pro: pro seat.
	//
	// - max: premium seat.
	//
	// example:
	//
	// standard
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
	// The start time of the seat.
	//
	// example:
	//
	// 1781422733
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The seat status. Valid values:
	//
	// - CREATING: being created.
	//
	// - NORMAL: active.
	//
	// - LIMIT: restricted due to overdue payment.
	//
	// - RELEASE: released upon expiration.
	//
	// - STOP: stopped upon expiration.
	//
	// - REFUNDED: refunded.
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetSubscriptionSeatDetailsResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s GetSubscriptionSeatDetailsResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *GetSubscriptionSeatDetailsResponseBodyDataItems) GetAccountEmail() *string {
	return s.AccountEmail
}

func (s *GetSubscriptionSeatDetailsResponseBodyDataItems) GetAccountId() *string {
	return s.AccountId
}

func (s *GetSubscriptionSeatDetailsResponseBodyDataItems) GetAccountName() *string {
	return s.AccountName
}

func (s *GetSubscriptionSeatDetailsResponseBodyDataItems) GetAssignedStatus() *string {
	return s.AssignedStatus
}

func (s *GetSubscriptionSeatDetailsResponseBodyDataItems) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetSubscriptionSeatDetailsResponseBodyDataItems) GetEquityList() []*GetSubscriptionSeatDetailsResponseBodyDataItemsEquityList {
	return s.EquityList
}

func (s *GetSubscriptionSeatDetailsResponseBodyDataItems) GetInstanceCode() *string {
	return s.InstanceCode
}

func (s *GetSubscriptionSeatDetailsResponseBodyDataItems) GetSeatId() *string {
	return s.SeatId
}

func (s *GetSubscriptionSeatDetailsResponseBodyDataItems) GetSpecType() *string {
	return s.SpecType
}

func (s *GetSubscriptionSeatDetailsResponseBodyDataItems) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetSubscriptionSeatDetailsResponseBodyDataItems) GetStatus() *string {
	return s.Status
}

func (s *GetSubscriptionSeatDetailsResponseBodyDataItems) SetAccountEmail(v string) *GetSubscriptionSeatDetailsResponseBodyDataItems {
	s.AccountEmail = &v
	return s
}

func (s *GetSubscriptionSeatDetailsResponseBodyDataItems) SetAccountId(v string) *GetSubscriptionSeatDetailsResponseBodyDataItems {
	s.AccountId = &v
	return s
}

func (s *GetSubscriptionSeatDetailsResponseBodyDataItems) SetAccountName(v string) *GetSubscriptionSeatDetailsResponseBodyDataItems {
	s.AccountName = &v
	return s
}

func (s *GetSubscriptionSeatDetailsResponseBodyDataItems) SetAssignedStatus(v string) *GetSubscriptionSeatDetailsResponseBodyDataItems {
	s.AssignedStatus = &v
	return s
}

func (s *GetSubscriptionSeatDetailsResponseBodyDataItems) SetEndTime(v int64) *GetSubscriptionSeatDetailsResponseBodyDataItems {
	s.EndTime = &v
	return s
}

func (s *GetSubscriptionSeatDetailsResponseBodyDataItems) SetEquityList(v []*GetSubscriptionSeatDetailsResponseBodyDataItemsEquityList) *GetSubscriptionSeatDetailsResponseBodyDataItems {
	s.EquityList = v
	return s
}

func (s *GetSubscriptionSeatDetailsResponseBodyDataItems) SetInstanceCode(v string) *GetSubscriptionSeatDetailsResponseBodyDataItems {
	s.InstanceCode = &v
	return s
}

func (s *GetSubscriptionSeatDetailsResponseBodyDataItems) SetSeatId(v string) *GetSubscriptionSeatDetailsResponseBodyDataItems {
	s.SeatId = &v
	return s
}

func (s *GetSubscriptionSeatDetailsResponseBodyDataItems) SetSpecType(v string) *GetSubscriptionSeatDetailsResponseBodyDataItems {
	s.SpecType = &v
	return s
}

func (s *GetSubscriptionSeatDetailsResponseBodyDataItems) SetStartTime(v int64) *GetSubscriptionSeatDetailsResponseBodyDataItems {
	s.StartTime = &v
	return s
}

func (s *GetSubscriptionSeatDetailsResponseBodyDataItems) SetStatus(v string) *GetSubscriptionSeatDetailsResponseBodyDataItems {
	s.Status = &v
	return s
}

func (s *GetSubscriptionSeatDetailsResponseBodyDataItems) Validate() error {
	if s.EquityList != nil {
		for _, item := range s.EquityList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSubscriptionSeatDetailsResponseBodyDataItemsEquityList struct {
	// The end time of the current cycle, in milliseconds.
	//
	// example:
	//
	// 1755273600
	CycleEndTime *int64 `json:"CycleEndTime,omitempty" xml:"CycleEndTime,omitempty"`
	// The equity code (subscription code). This does not need to be consumed in CREDITS scenarios.
	//
	// example:
	//
	// 123456
	CycleInstanceId *string `json:"CycleInstanceId,omitempty" xml:"CycleInstanceId,omitempty"`
	// The start time of the current cycle, in milliseconds.
	//
	// example:
	//
	// 1775232000
	CycleStartTime *int64 `json:"CycleStartTime,omitempty" xml:"CycleStartTime,omitempty"`
	// The remaining quota of the current cycle.
	//
	// example:
	//
	// 40.00000000
	CycleSurplusValue *float64 `json:"CycleSurplusValue,omitempty" xml:"CycleSurplusValue,omitempty"`
	// The total quota of the current cycle.
	//
	// example:
	//
	// 100.00000000
	CycleTotalValue *float64 `json:"CycleTotalValue,omitempty" xml:"CycleTotalValue,omitempty"`
	// The time-series version of the current cycle.
	//
	// example:
	//
	// 1
	CycleVersion *int64 `json:"CycleVersion,omitempty" xml:"CycleVersion,omitempty"`
	// The equity type, such as CREDITS, SPN, or resource plan.
	//
	// example:
	//
	// CREDITS
	EquityType *string `json:"EquityType,omitempty" xml:"EquityType,omitempty"`
}

func (s GetSubscriptionSeatDetailsResponseBodyDataItemsEquityList) String() string {
	return dara.Prettify(s)
}

func (s GetSubscriptionSeatDetailsResponseBodyDataItemsEquityList) GoString() string {
	return s.String()
}

func (s *GetSubscriptionSeatDetailsResponseBodyDataItemsEquityList) GetCycleEndTime() *int64 {
	return s.CycleEndTime
}

func (s *GetSubscriptionSeatDetailsResponseBodyDataItemsEquityList) GetCycleInstanceId() *string {
	return s.CycleInstanceId
}

func (s *GetSubscriptionSeatDetailsResponseBodyDataItemsEquityList) GetCycleStartTime() *int64 {
	return s.CycleStartTime
}

func (s *GetSubscriptionSeatDetailsResponseBodyDataItemsEquityList) GetCycleSurplusValue() *float64 {
	return s.CycleSurplusValue
}

func (s *GetSubscriptionSeatDetailsResponseBodyDataItemsEquityList) GetCycleTotalValue() *float64 {
	return s.CycleTotalValue
}

func (s *GetSubscriptionSeatDetailsResponseBodyDataItemsEquityList) GetCycleVersion() *int64 {
	return s.CycleVersion
}

func (s *GetSubscriptionSeatDetailsResponseBodyDataItemsEquityList) GetEquityType() *string {
	return s.EquityType
}

func (s *GetSubscriptionSeatDetailsResponseBodyDataItemsEquityList) SetCycleEndTime(v int64) *GetSubscriptionSeatDetailsResponseBodyDataItemsEquityList {
	s.CycleEndTime = &v
	return s
}

func (s *GetSubscriptionSeatDetailsResponseBodyDataItemsEquityList) SetCycleInstanceId(v string) *GetSubscriptionSeatDetailsResponseBodyDataItemsEquityList {
	s.CycleInstanceId = &v
	return s
}

func (s *GetSubscriptionSeatDetailsResponseBodyDataItemsEquityList) SetCycleStartTime(v int64) *GetSubscriptionSeatDetailsResponseBodyDataItemsEquityList {
	s.CycleStartTime = &v
	return s
}

func (s *GetSubscriptionSeatDetailsResponseBodyDataItemsEquityList) SetCycleSurplusValue(v float64) *GetSubscriptionSeatDetailsResponseBodyDataItemsEquityList {
	s.CycleSurplusValue = &v
	return s
}

func (s *GetSubscriptionSeatDetailsResponseBodyDataItemsEquityList) SetCycleTotalValue(v float64) *GetSubscriptionSeatDetailsResponseBodyDataItemsEquityList {
	s.CycleTotalValue = &v
	return s
}

func (s *GetSubscriptionSeatDetailsResponseBodyDataItemsEquityList) SetCycleVersion(v int64) *GetSubscriptionSeatDetailsResponseBodyDataItemsEquityList {
	s.CycleVersion = &v
	return s
}

func (s *GetSubscriptionSeatDetailsResponseBodyDataItemsEquityList) SetEquityType(v string) *GetSubscriptionSeatDetailsResponseBodyDataItemsEquityList {
	s.EquityType = &v
	return s
}

func (s *GetSubscriptionSeatDetailsResponseBodyDataItemsEquityList) Validate() error {
	return dara.Validate(s)
}
