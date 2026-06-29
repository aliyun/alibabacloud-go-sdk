// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubscriptionSharedPackagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSubscriptionSharedPackagesResponseBody
	GetCode() *string
	SetData(v *ListSubscriptionSharedPackagesResponseBodyData) *ListSubscriptionSharedPackagesResponseBody
	GetData() *ListSubscriptionSharedPackagesResponseBodyData
	SetMessage(v string) *ListSubscriptionSharedPackagesResponseBody
	GetMessage() *string
	SetSuccess(v bool) *ListSubscriptionSharedPackagesResponseBody
	GetSuccess() *bool
}

type ListSubscriptionSharedPackagesResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data.
	Data *ListSubscriptionSharedPackagesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the API call is successful. Valid values:
	//
	// - true: The call is successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSubscriptionSharedPackagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSubscriptionSharedPackagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubscriptionSharedPackagesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSubscriptionSharedPackagesResponseBody) GetData() *ListSubscriptionSharedPackagesResponseBodyData {
	return s.Data
}

func (s *ListSubscriptionSharedPackagesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSubscriptionSharedPackagesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSubscriptionSharedPackagesResponseBody) SetCode(v string) *ListSubscriptionSharedPackagesResponseBody {
	s.Code = &v
	return s
}

func (s *ListSubscriptionSharedPackagesResponseBody) SetData(v *ListSubscriptionSharedPackagesResponseBodyData) *ListSubscriptionSharedPackagesResponseBody {
	s.Data = v
	return s
}

func (s *ListSubscriptionSharedPackagesResponseBody) SetMessage(v string) *ListSubscriptionSharedPackagesResponseBody {
	s.Message = &v
	return s
}

func (s *ListSubscriptionSharedPackagesResponseBody) SetSuccess(v bool) *ListSubscriptionSharedPackagesResponseBody {
	s.Success = &v
	return s
}

func (s *ListSubscriptionSharedPackagesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSubscriptionSharedPackagesResponseBodyData struct {
	// The data entries.
	Items []*ListSubscriptionSharedPackagesResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number. The value must be greater than 0 and cannot exceed the maximum value of the Integer data type. Default value: 1.
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

func (s ListSubscriptionSharedPackagesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSubscriptionSharedPackagesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSubscriptionSharedPackagesResponseBodyData) GetItems() []*ListSubscriptionSharedPackagesResponseBodyDataItems {
	return s.Items
}

func (s *ListSubscriptionSharedPackagesResponseBodyData) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListSubscriptionSharedPackagesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSubscriptionSharedPackagesResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListSubscriptionSharedPackagesResponseBodyData) SetItems(v []*ListSubscriptionSharedPackagesResponseBodyDataItems) *ListSubscriptionSharedPackagesResponseBodyData {
	s.Items = v
	return s
}

func (s *ListSubscriptionSharedPackagesResponseBodyData) SetPageNo(v int32) *ListSubscriptionSharedPackagesResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *ListSubscriptionSharedPackagesResponseBodyData) SetPageSize(v int32) *ListSubscriptionSharedPackagesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListSubscriptionSharedPackagesResponseBodyData) SetTotal(v int32) *ListSubscriptionSharedPackagesResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListSubscriptionSharedPackagesResponseBodyData) Validate() error {
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

type ListSubscriptionSharedPackagesResponseBodyDataItems struct {
	// The currently active equity instances.
	EquityList []*ListSubscriptionSharedPackagesResponseBodyDataItemsEquityList `json:"EquityList,omitempty" xml:"EquityList,omitempty" type:"Repeated"`
	// The instance code of the seat.
	//
	// example:
	//
	// subs-1234567
	InstanceCode *string `json:"InstanceCode,omitempty" xml:"InstanceCode,omitempty"`
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

func (s ListSubscriptionSharedPackagesResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListSubscriptionSharedPackagesResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListSubscriptionSharedPackagesResponseBodyDataItems) GetEquityList() []*ListSubscriptionSharedPackagesResponseBodyDataItemsEquityList {
	return s.EquityList
}

func (s *ListSubscriptionSharedPackagesResponseBodyDataItems) GetInstanceCode() *string {
	return s.InstanceCode
}

func (s *ListSubscriptionSharedPackagesResponseBodyDataItems) GetStatus() *string {
	return s.Status
}

func (s *ListSubscriptionSharedPackagesResponseBodyDataItems) SetEquityList(v []*ListSubscriptionSharedPackagesResponseBodyDataItemsEquityList) *ListSubscriptionSharedPackagesResponseBodyDataItems {
	s.EquityList = v
	return s
}

func (s *ListSubscriptionSharedPackagesResponseBodyDataItems) SetInstanceCode(v string) *ListSubscriptionSharedPackagesResponseBodyDataItems {
	s.InstanceCode = &v
	return s
}

func (s *ListSubscriptionSharedPackagesResponseBodyDataItems) SetStatus(v string) *ListSubscriptionSharedPackagesResponseBodyDataItems {
	s.Status = &v
	return s
}

func (s *ListSubscriptionSharedPackagesResponseBodyDataItems) Validate() error {
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

type ListSubscriptionSharedPackagesResponseBodyDataItemsEquityList struct {
	// The end time of the current cycle, in milliseconds.
	//
	// example:
	//
	// 1756310400
	CycleEndTime *int64 `json:"CycleEndTime,omitempty" xml:"CycleEndTime,omitempty"`
	// The equity code (subscription code). This is not required for consumption in the credits scenario.
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
	// The sequential version of the current cycle.
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

func (s ListSubscriptionSharedPackagesResponseBodyDataItemsEquityList) String() string {
	return dara.Prettify(s)
}

func (s ListSubscriptionSharedPackagesResponseBodyDataItemsEquityList) GoString() string {
	return s.String()
}

func (s *ListSubscriptionSharedPackagesResponseBodyDataItemsEquityList) GetCycleEndTime() *int64 {
	return s.CycleEndTime
}

func (s *ListSubscriptionSharedPackagesResponseBodyDataItemsEquityList) GetCycleInstanceId() *string {
	return s.CycleInstanceId
}

func (s *ListSubscriptionSharedPackagesResponseBodyDataItemsEquityList) GetCycleStartTime() *int64 {
	return s.CycleStartTime
}

func (s *ListSubscriptionSharedPackagesResponseBodyDataItemsEquityList) GetCycleSurplusValue() *float64 {
	return s.CycleSurplusValue
}

func (s *ListSubscriptionSharedPackagesResponseBodyDataItemsEquityList) GetCycleTotalValue() *float64 {
	return s.CycleTotalValue
}

func (s *ListSubscriptionSharedPackagesResponseBodyDataItemsEquityList) GetCycleVersion() *int64 {
	return s.CycleVersion
}

func (s *ListSubscriptionSharedPackagesResponseBodyDataItemsEquityList) GetEquityType() *string {
	return s.EquityType
}

func (s *ListSubscriptionSharedPackagesResponseBodyDataItemsEquityList) SetCycleEndTime(v int64) *ListSubscriptionSharedPackagesResponseBodyDataItemsEquityList {
	s.CycleEndTime = &v
	return s
}

func (s *ListSubscriptionSharedPackagesResponseBodyDataItemsEquityList) SetCycleInstanceId(v string) *ListSubscriptionSharedPackagesResponseBodyDataItemsEquityList {
	s.CycleInstanceId = &v
	return s
}

func (s *ListSubscriptionSharedPackagesResponseBodyDataItemsEquityList) SetCycleStartTime(v int64) *ListSubscriptionSharedPackagesResponseBodyDataItemsEquityList {
	s.CycleStartTime = &v
	return s
}

func (s *ListSubscriptionSharedPackagesResponseBodyDataItemsEquityList) SetCycleSurplusValue(v float64) *ListSubscriptionSharedPackagesResponseBodyDataItemsEquityList {
	s.CycleSurplusValue = &v
	return s
}

func (s *ListSubscriptionSharedPackagesResponseBodyDataItemsEquityList) SetCycleTotalValue(v float64) *ListSubscriptionSharedPackagesResponseBodyDataItemsEquityList {
	s.CycleTotalValue = &v
	return s
}

func (s *ListSubscriptionSharedPackagesResponseBodyDataItemsEquityList) SetCycleVersion(v int64) *ListSubscriptionSharedPackagesResponseBodyDataItemsEquityList {
	s.CycleVersion = &v
	return s
}

func (s *ListSubscriptionSharedPackagesResponseBodyDataItemsEquityList) SetEquityType(v string) *ListSubscriptionSharedPackagesResponseBodyDataItemsEquityList {
	s.EquityType = &v
	return s
}

func (s *ListSubscriptionSharedPackagesResponseBodyDataItemsEquityList) Validate() error {
	return dara.Validate(s)
}
