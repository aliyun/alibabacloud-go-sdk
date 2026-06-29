// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubscriptionStatsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSubscriptionStatsResponseBody
	GetCode() *string
	SetData(v *GetSubscriptionStatsResponseBodyData) *GetSubscriptionStatsResponseBody
	GetData() *GetSubscriptionStatsResponseBodyData
	SetMessage(v string) *GetSubscriptionStatsResponseBody
	GetMessage() *string
	SetSuccess(v bool) *GetSubscriptionStatsResponseBody
	GetSuccess() *bool
}

type GetSubscriptionStatsResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data.
	Data *GetSubscriptionStatsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the API call is successful. Valid values:
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

func (s GetSubscriptionStatsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSubscriptionStatsResponseBody) GoString() string {
	return s.String()
}

func (s *GetSubscriptionStatsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSubscriptionStatsResponseBody) GetData() *GetSubscriptionStatsResponseBodyData {
	return s.Data
}

func (s *GetSubscriptionStatsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSubscriptionStatsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSubscriptionStatsResponseBody) SetCode(v string) *GetSubscriptionStatsResponseBody {
	s.Code = &v
	return s
}

func (s *GetSubscriptionStatsResponseBody) SetData(v *GetSubscriptionStatsResponseBodyData) *GetSubscriptionStatsResponseBody {
	s.Data = v
	return s
}

func (s *GetSubscriptionStatsResponseBody) SetMessage(v string) *GetSubscriptionStatsResponseBody {
	s.Message = &v
	return s
}

func (s *GetSubscriptionStatsResponseBody) SetSuccess(v bool) *GetSubscriptionStatsResponseBody {
	s.Success = &v
	return s
}

func (s *GetSubscriptionStatsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSubscriptionStatsResponseBodyData struct {
	// The list of seat information, grouped by specType.
	Items []*GetSubscriptionStatsResponseBodyDataItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The subscription end time, in milliseconds.
	//
	// example:
	//
	// 1781422734
	SubscriptionEndTime *int64 `json:"SubscriptionEndTime,omitempty" xml:"SubscriptionEndTime,omitempty"`
	// The subscription start time, in milliseconds.
	//
	// example:
	//
	// 1781422733
	SubscriptionStartTime *int64 `json:"SubscriptionStartTime,omitempty" xml:"SubscriptionStartTime,omitempty"`
}

func (s GetSubscriptionStatsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSubscriptionStatsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSubscriptionStatsResponseBodyData) GetItems() []*GetSubscriptionStatsResponseBodyDataItems {
	return s.Items
}

func (s *GetSubscriptionStatsResponseBodyData) GetSubscriptionEndTime() *int64 {
	return s.SubscriptionEndTime
}

func (s *GetSubscriptionStatsResponseBodyData) GetSubscriptionStartTime() *int64 {
	return s.SubscriptionStartTime
}

func (s *GetSubscriptionStatsResponseBodyData) SetItems(v []*GetSubscriptionStatsResponseBodyDataItems) *GetSubscriptionStatsResponseBodyData {
	s.Items = v
	return s
}

func (s *GetSubscriptionStatsResponseBodyData) SetSubscriptionEndTime(v int64) *GetSubscriptionStatsResponseBodyData {
	s.SubscriptionEndTime = &v
	return s
}

func (s *GetSubscriptionStatsResponseBodyData) SetSubscriptionStartTime(v int64) *GetSubscriptionStatsResponseBodyData {
	s.SubscriptionStartTime = &v
	return s
}

func (s *GetSubscriptionStatsResponseBodyData) Validate() error {
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

type GetSubscriptionStatsResponseBodyDataItems struct {
	// The number of assigned seats.
	//
	// example:
	//
	// 1
	AssignedSeats *int64 `json:"AssignedSeats,omitempty" xml:"AssignedSeats,omitempty"`
	// The total credits quota for the seat.
	//
	// example:
	//
	// 100
	SeatCredits *float64 `json:"SeatCredits,omitempty" xml:"SeatCredits,omitempty"`
	// The refresh time of the current cycle, in milliseconds.
	//
	// example:
	//
	// 1781422734
	SeatRefreshTime *int64 `json:"SeatRefreshTime,omitempty" xml:"SeatRefreshTime,omitempty"`
	// The remaining credits for the current cycle.
	//
	// example:
	//
	// 40
	SeatRemainingCredits *float64 `json:"SeatRemainingCredits,omitempty" xml:"SeatRemainingCredits,omitempty"`
	// The seat type (specType). Valid values:
	//
	// - standard: Standard seat.
	//
	// - pro: Pro seat.
	//
	// - max: Premium seat.
	//
	// example:
	//
	// standard
	SeatType *string `json:"SeatType,omitempty" xml:"SeatType,omitempty"`
	// The total number of seats.
	//
	// example:
	//
	// 2
	TotalSeats *int64 `json:"TotalSeats,omitempty" xml:"TotalSeats,omitempty"`
}

func (s GetSubscriptionStatsResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s GetSubscriptionStatsResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *GetSubscriptionStatsResponseBodyDataItems) GetAssignedSeats() *int64 {
	return s.AssignedSeats
}

func (s *GetSubscriptionStatsResponseBodyDataItems) GetSeatCredits() *float64 {
	return s.SeatCredits
}

func (s *GetSubscriptionStatsResponseBodyDataItems) GetSeatRefreshTime() *int64 {
	return s.SeatRefreshTime
}

func (s *GetSubscriptionStatsResponseBodyDataItems) GetSeatRemainingCredits() *float64 {
	return s.SeatRemainingCredits
}

func (s *GetSubscriptionStatsResponseBodyDataItems) GetSeatType() *string {
	return s.SeatType
}

func (s *GetSubscriptionStatsResponseBodyDataItems) GetTotalSeats() *int64 {
	return s.TotalSeats
}

func (s *GetSubscriptionStatsResponseBodyDataItems) SetAssignedSeats(v int64) *GetSubscriptionStatsResponseBodyDataItems {
	s.AssignedSeats = &v
	return s
}

func (s *GetSubscriptionStatsResponseBodyDataItems) SetSeatCredits(v float64) *GetSubscriptionStatsResponseBodyDataItems {
	s.SeatCredits = &v
	return s
}

func (s *GetSubscriptionStatsResponseBodyDataItems) SetSeatRefreshTime(v int64) *GetSubscriptionStatsResponseBodyDataItems {
	s.SeatRefreshTime = &v
	return s
}

func (s *GetSubscriptionStatsResponseBodyDataItems) SetSeatRemainingCredits(v float64) *GetSubscriptionStatsResponseBodyDataItems {
	s.SeatRemainingCredits = &v
	return s
}

func (s *GetSubscriptionStatsResponseBodyDataItems) SetSeatType(v string) *GetSubscriptionStatsResponseBodyDataItems {
	s.SeatType = &v
	return s
}

func (s *GetSubscriptionStatsResponseBodyDataItems) SetTotalSeats(v int64) *GetSubscriptionStatsResponseBodyDataItems {
	s.TotalSeats = &v
	return s
}

func (s *GetSubscriptionStatsResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
