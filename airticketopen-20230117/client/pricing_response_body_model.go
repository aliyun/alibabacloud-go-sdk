// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPricingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PricingResponseBody
	GetRequestId() *string
	SetData(v *PricingResponseBodyData) *PricingResponseBody
	GetData() *PricingResponseBodyData
	SetErrorCode(v string) *PricingResponseBody
	GetErrorCode() *string
	SetErrorData(v interface{}) *PricingResponseBody
	GetErrorData() interface{}
	SetErrorMsg(v string) *PricingResponseBody
	GetErrorMsg() *string
	SetStatus(v int32) *PricingResponseBody
	GetStatus() *int32
	SetSuccess(v bool) *PricingResponseBody
	GetSuccess() *bool
}

type PricingResponseBody struct {
	// request ID
	//
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// data
	Data *PricingResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// error code
	//
	// example:
	//
	// null
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// error data
	//
	// example:
	//
	// null
	ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
	// error message
	//
	// example:
	//
	// null
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// http reqeust has been processed successfully，status code is 200
	//
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// true represents success, false represents failure
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PricingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PricingResponseBody) GoString() string {
	return s.String()
}

func (s *PricingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PricingResponseBody) GetData() *PricingResponseBodyData {
	return s.Data
}

func (s *PricingResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *PricingResponseBody) GetErrorData() interface{} {
	return s.ErrorData
}

func (s *PricingResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *PricingResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *PricingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PricingResponseBody) SetRequestId(v string) *PricingResponseBody {
	s.RequestId = &v
	return s
}

func (s *PricingResponseBody) SetData(v *PricingResponseBodyData) *PricingResponseBody {
	s.Data = v
	return s
}

func (s *PricingResponseBody) SetErrorCode(v string) *PricingResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *PricingResponseBody) SetErrorData(v interface{}) *PricingResponseBody {
	s.ErrorData = v
	return s
}

func (s *PricingResponseBody) SetErrorMsg(v string) *PricingResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *PricingResponseBody) SetStatus(v int32) *PricingResponseBody {
	s.Status = &v
	return s
}

func (s *PricingResponseBody) SetSuccess(v bool) *PricingResponseBody {
	s.Success = &v
	return s
}

func (s *PricingResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PricingResponseBodyData struct {
	// price information after the price change
	ChangedPriceInfo *PricingResponseBodyDataChangedPriceInfo `json:"changed_price_info,omitempty" xml:"changed_price_info,omitempty" type:"Struct"`
	// whether the price has changed
	//
	// example:
	//
	// true
	IsChanged *bool `json:"is_changed,omitempty" xml:"is_changed,omitempty"`
	// the price information before the change, only available when is_changed = true
	OriginalPriceInfo *PricingResponseBodyDataOriginalPriceInfo `json:"original_price_info,omitempty" xml:"original_price_info,omitempty" type:"Struct"`
	// remaining seats: A indicates more than 9, 0-9 represents the specific number
	//
	// example:
	//
	// A
	RemainSeats *string `json:"remain_seats,omitempty" xml:"remain_seats,omitempty"`
	// the solution represented by the solution_id in request
	Solution *PricingResponseBodyDataSolution `json:"solution,omitempty" xml:"solution,omitempty" type:"Struct"`
}

func (s PricingResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PricingResponseBodyData) GoString() string {
	return s.String()
}

func (s *PricingResponseBodyData) GetChangedPriceInfo() *PricingResponseBodyDataChangedPriceInfo {
	return s.ChangedPriceInfo
}

func (s *PricingResponseBodyData) GetIsChanged() *bool {
	return s.IsChanged
}

func (s *PricingResponseBodyData) GetOriginalPriceInfo() *PricingResponseBodyDataOriginalPriceInfo {
	return s.OriginalPriceInfo
}

func (s *PricingResponseBodyData) GetRemainSeats() *string {
	return s.RemainSeats
}

func (s *PricingResponseBodyData) GetSolution() *PricingResponseBodyDataSolution {
	return s.Solution
}

func (s *PricingResponseBodyData) SetChangedPriceInfo(v *PricingResponseBodyDataChangedPriceInfo) *PricingResponseBodyData {
	s.ChangedPriceInfo = v
	return s
}

func (s *PricingResponseBodyData) SetIsChanged(v bool) *PricingResponseBodyData {
	s.IsChanged = &v
	return s
}

func (s *PricingResponseBodyData) SetOriginalPriceInfo(v *PricingResponseBodyDataOriginalPriceInfo) *PricingResponseBodyData {
	s.OriginalPriceInfo = v
	return s
}

func (s *PricingResponseBodyData) SetRemainSeats(v string) *PricingResponseBodyData {
	s.RemainSeats = &v
	return s
}

func (s *PricingResponseBodyData) SetSolution(v *PricingResponseBodyDataSolution) *PricingResponseBodyData {
	s.Solution = v
	return s
}

func (s *PricingResponseBodyData) Validate() error {
	if s.ChangedPriceInfo != nil {
		if err := s.ChangedPriceInfo.Validate(); err != nil {
			return err
		}
	}
	if s.OriginalPriceInfo != nil {
		if err := s.OriginalPriceInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Solution != nil {
		if err := s.Solution.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PricingResponseBodyDataChangedPriceInfo struct {
	// adult fare
	//
	// example:
	//
	// 100
	AdultPrice *float64 `json:"adult_price,omitempty" xml:"adult_price,omitempty"`
	// adult tax
	//
	// example:
	//
	// 10
	AdultTax *float64 `json:"adult_tax,omitempty" xml:"adult_tax,omitempty"`
	// child fare
	//
	// example:
	//
	// 100
	ChildPrice *float64 `json:"child_price,omitempty" xml:"child_price,omitempty"`
	// child tax
	//
	// example:
	//
	// 10
	ChildTax *float64 `json:"child_tax,omitempty" xml:"child_tax,omitempty"`
	// infant fare
	//
	// example:
	//
	// 100
	InfantPrice *float64 `json:"infant_price,omitempty" xml:"infant_price,omitempty"`
	// infant tax
	//
	// example:
	//
	// 10
	InfantTax *float64 `json:"infant_tax,omitempty" xml:"infant_tax,omitempty"`
}

func (s PricingResponseBodyDataChangedPriceInfo) String() string {
	return dara.Prettify(s)
}

func (s PricingResponseBodyDataChangedPriceInfo) GoString() string {
	return s.String()
}

func (s *PricingResponseBodyDataChangedPriceInfo) GetAdultPrice() *float64 {
	return s.AdultPrice
}

func (s *PricingResponseBodyDataChangedPriceInfo) GetAdultTax() *float64 {
	return s.AdultTax
}

func (s *PricingResponseBodyDataChangedPriceInfo) GetChildPrice() *float64 {
	return s.ChildPrice
}

func (s *PricingResponseBodyDataChangedPriceInfo) GetChildTax() *float64 {
	return s.ChildTax
}

func (s *PricingResponseBodyDataChangedPriceInfo) GetInfantPrice() *float64 {
	return s.InfantPrice
}

func (s *PricingResponseBodyDataChangedPriceInfo) GetInfantTax() *float64 {
	return s.InfantTax
}

func (s *PricingResponseBodyDataChangedPriceInfo) SetAdultPrice(v float64) *PricingResponseBodyDataChangedPriceInfo {
	s.AdultPrice = &v
	return s
}

func (s *PricingResponseBodyDataChangedPriceInfo) SetAdultTax(v float64) *PricingResponseBodyDataChangedPriceInfo {
	s.AdultTax = &v
	return s
}

func (s *PricingResponseBodyDataChangedPriceInfo) SetChildPrice(v float64) *PricingResponseBodyDataChangedPriceInfo {
	s.ChildPrice = &v
	return s
}

func (s *PricingResponseBodyDataChangedPriceInfo) SetChildTax(v float64) *PricingResponseBodyDataChangedPriceInfo {
	s.ChildTax = &v
	return s
}

func (s *PricingResponseBodyDataChangedPriceInfo) SetInfantPrice(v float64) *PricingResponseBodyDataChangedPriceInfo {
	s.InfantPrice = &v
	return s
}

func (s *PricingResponseBodyDataChangedPriceInfo) SetInfantTax(v float64) *PricingResponseBodyDataChangedPriceInfo {
	s.InfantTax = &v
	return s
}

func (s *PricingResponseBodyDataChangedPriceInfo) Validate() error {
	return dara.Validate(s)
}

type PricingResponseBodyDataOriginalPriceInfo struct {
	// adult fare
	//
	// example:
	//
	// 200
	AdultPrice *float64 `json:"adult_price,omitempty" xml:"adult_price,omitempty"`
	// adult tax
	//
	// example:
	//
	// 20
	AdultTax *float64 `json:"adult_tax,omitempty" xml:"adult_tax,omitempty"`
	// child fare
	//
	// example:
	//
	// 200
	ChildPrice *float64 `json:"child_price,omitempty" xml:"child_price,omitempty"`
	// child tax
	//
	// example:
	//
	// 20
	ChildTax *float64 `json:"child_tax,omitempty" xml:"child_tax,omitempty"`
	// infant fare
	//
	// example:
	//
	// 200
	InfantPrice *float64 `json:"infant_price,omitempty" xml:"infant_price,omitempty"`
	// infant tax
	//
	// example:
	//
	// 20
	InfantTax *float64 `json:"infant_tax,omitempty" xml:"infant_tax,omitempty"`
}

func (s PricingResponseBodyDataOriginalPriceInfo) String() string {
	return dara.Prettify(s)
}

func (s PricingResponseBodyDataOriginalPriceInfo) GoString() string {
	return s.String()
}

func (s *PricingResponseBodyDataOriginalPriceInfo) GetAdultPrice() *float64 {
	return s.AdultPrice
}

func (s *PricingResponseBodyDataOriginalPriceInfo) GetAdultTax() *float64 {
	return s.AdultTax
}

func (s *PricingResponseBodyDataOriginalPriceInfo) GetChildPrice() *float64 {
	return s.ChildPrice
}

func (s *PricingResponseBodyDataOriginalPriceInfo) GetChildTax() *float64 {
	return s.ChildTax
}

func (s *PricingResponseBodyDataOriginalPriceInfo) GetInfantPrice() *float64 {
	return s.InfantPrice
}

func (s *PricingResponseBodyDataOriginalPriceInfo) GetInfantTax() *float64 {
	return s.InfantTax
}

func (s *PricingResponseBodyDataOriginalPriceInfo) SetAdultPrice(v float64) *PricingResponseBodyDataOriginalPriceInfo {
	s.AdultPrice = &v
	return s
}

func (s *PricingResponseBodyDataOriginalPriceInfo) SetAdultTax(v float64) *PricingResponseBodyDataOriginalPriceInfo {
	s.AdultTax = &v
	return s
}

func (s *PricingResponseBodyDataOriginalPriceInfo) SetChildPrice(v float64) *PricingResponseBodyDataOriginalPriceInfo {
	s.ChildPrice = &v
	return s
}

func (s *PricingResponseBodyDataOriginalPriceInfo) SetChildTax(v float64) *PricingResponseBodyDataOriginalPriceInfo {
	s.ChildTax = &v
	return s
}

func (s *PricingResponseBodyDataOriginalPriceInfo) SetInfantPrice(v float64) *PricingResponseBodyDataOriginalPriceInfo {
	s.InfantPrice = &v
	return s
}

func (s *PricingResponseBodyDataOriginalPriceInfo) SetInfantTax(v float64) *PricingResponseBodyDataOriginalPriceInfo {
	s.InfantTax = &v
	return s
}

func (s *PricingResponseBodyDataOriginalPriceInfo) Validate() error {
	return dara.Validate(s)
}

type PricingResponseBodyDataSolution struct {
	// adult fare
	//
	// example:
	//
	// 300
	AdultPrice *float64 `json:"adult_price,omitempty" xml:"adult_price,omitempty"`
	// adult tax
	//
	// example:
	//
	// 30
	AdultTax *float64 `json:"adult_tax,omitempty" xml:"adult_tax,omitempty"`
	// child fare
	//
	// example:
	//
	// 200
	ChildPrice *float64 `json:"child_price,omitempty" xml:"child_price,omitempty"`
	// child tax
	//
	// example:
	//
	// 20
	ChildTax *float64 `json:"child_tax,omitempty" xml:"child_tax,omitempty"`
	// infant fare
	//
	// example:
	//
	// 200
	InfantPrice *float64 `json:"infant_price,omitempty" xml:"infant_price,omitempty"`
	// infant tax
	//
	// example:
	//
	// 10
	InfantTax *float64 `json:"infant_tax,omitempty" xml:"infant_tax,omitempty"`
	// journey list
	JourneyList []*PricingResponseBodyDataSolutionJourneyList `json:"journey_list,omitempty" xml:"journey_list,omitempty" type:"Repeated"`
	// product type description
	//
	// example:
	//
	// description
	ProductTypeDescription *string `json:"product_type_description,omitempty" xml:"product_type_description,omitempty"`
	// refund coupon description
	//
	// example:
	//
	// description
	RefundTicketCouponDescription *string `json:"refund_ticket_coupon_description,omitempty" xml:"refund_ticket_coupon_description,omitempty"`
	// through check-in baggage policy
	SegmentBaggageCheckInInfoList []*PricingResponseBodyDataSolutionSegmentBaggageCheckInInfoList `json:"segment_baggage_check_in_info_list,omitempty" xml:"segment_baggage_check_in_info_list,omitempty" type:"Repeated"`
	// baggage rule list
	SegmentBaggageMappingList []*PricingResponseBodyDataSolutionSegmentBaggageMappingList `json:"segment_baggage_mapping_list,omitempty" xml:"segment_baggage_mapping_list,omitempty" type:"Repeated"`
	// change and refund policy
	SegmentRefundChangeRuleMappingList []*PricingResponseBodyDataSolutionSegmentRefundChangeRuleMappingList `json:"segment_refund_change_rule_mapping_list,omitempty" xml:"segment_refund_change_rule_mapping_list,omitempty" type:"Repeated"`
	// solution_id, equals to solution_id in request
	//
	// example:
	//
	// eJwz8DeySEo0NjQ01TU3TU7TNTFINNO1SE5O0jVKM0hKNjEwTElLNYwz0A32cNT1dfPVNTIwMjYwNjRQ8/A3NLI01Q0Ic0cRBwBVFxJJ
	SolutionId *string `json:"solution_id,omitempty" xml:"solution_id,omitempty"`
}

func (s PricingResponseBodyDataSolution) String() string {
	return dara.Prettify(s)
}

func (s PricingResponseBodyDataSolution) GoString() string {
	return s.String()
}

func (s *PricingResponseBodyDataSolution) GetAdultPrice() *float64 {
	return s.AdultPrice
}

func (s *PricingResponseBodyDataSolution) GetAdultTax() *float64 {
	return s.AdultTax
}

func (s *PricingResponseBodyDataSolution) GetChildPrice() *float64 {
	return s.ChildPrice
}

func (s *PricingResponseBodyDataSolution) GetChildTax() *float64 {
	return s.ChildTax
}

func (s *PricingResponseBodyDataSolution) GetInfantPrice() *float64 {
	return s.InfantPrice
}

func (s *PricingResponseBodyDataSolution) GetInfantTax() *float64 {
	return s.InfantTax
}

func (s *PricingResponseBodyDataSolution) GetJourneyList() []*PricingResponseBodyDataSolutionJourneyList {
	return s.JourneyList
}

func (s *PricingResponseBodyDataSolution) GetProductTypeDescription() *string {
	return s.ProductTypeDescription
}

func (s *PricingResponseBodyDataSolution) GetRefundTicketCouponDescription() *string {
	return s.RefundTicketCouponDescription
}

func (s *PricingResponseBodyDataSolution) GetSegmentBaggageCheckInInfoList() []*PricingResponseBodyDataSolutionSegmentBaggageCheckInInfoList {
	return s.SegmentBaggageCheckInInfoList
}

func (s *PricingResponseBodyDataSolution) GetSegmentBaggageMappingList() []*PricingResponseBodyDataSolutionSegmentBaggageMappingList {
	return s.SegmentBaggageMappingList
}

func (s *PricingResponseBodyDataSolution) GetSegmentRefundChangeRuleMappingList() []*PricingResponseBodyDataSolutionSegmentRefundChangeRuleMappingList {
	return s.SegmentRefundChangeRuleMappingList
}

func (s *PricingResponseBodyDataSolution) GetSolutionId() *string {
	return s.SolutionId
}

func (s *PricingResponseBodyDataSolution) SetAdultPrice(v float64) *PricingResponseBodyDataSolution {
	s.AdultPrice = &v
	return s
}

func (s *PricingResponseBodyDataSolution) SetAdultTax(v float64) *PricingResponseBodyDataSolution {
	s.AdultTax = &v
	return s
}

func (s *PricingResponseBodyDataSolution) SetChildPrice(v float64) *PricingResponseBodyDataSolution {
	s.ChildPrice = &v
	return s
}

func (s *PricingResponseBodyDataSolution) SetChildTax(v float64) *PricingResponseBodyDataSolution {
	s.ChildTax = &v
	return s
}

func (s *PricingResponseBodyDataSolution) SetInfantPrice(v float64) *PricingResponseBodyDataSolution {
	s.InfantPrice = &v
	return s
}

func (s *PricingResponseBodyDataSolution) SetInfantTax(v float64) *PricingResponseBodyDataSolution {
	s.InfantTax = &v
	return s
}

func (s *PricingResponseBodyDataSolution) SetJourneyList(v []*PricingResponseBodyDataSolutionJourneyList) *PricingResponseBodyDataSolution {
	s.JourneyList = v
	return s
}

func (s *PricingResponseBodyDataSolution) SetProductTypeDescription(v string) *PricingResponseBodyDataSolution {
	s.ProductTypeDescription = &v
	return s
}

func (s *PricingResponseBodyDataSolution) SetRefundTicketCouponDescription(v string) *PricingResponseBodyDataSolution {
	s.RefundTicketCouponDescription = &v
	return s
}

func (s *PricingResponseBodyDataSolution) SetSegmentBaggageCheckInInfoList(v []*PricingResponseBodyDataSolutionSegmentBaggageCheckInInfoList) *PricingResponseBodyDataSolution {
	s.SegmentBaggageCheckInInfoList = v
	return s
}

func (s *PricingResponseBodyDataSolution) SetSegmentBaggageMappingList(v []*PricingResponseBodyDataSolutionSegmentBaggageMappingList) *PricingResponseBodyDataSolution {
	s.SegmentBaggageMappingList = v
	return s
}

func (s *PricingResponseBodyDataSolution) SetSegmentRefundChangeRuleMappingList(v []*PricingResponseBodyDataSolutionSegmentRefundChangeRuleMappingList) *PricingResponseBodyDataSolution {
	s.SegmentRefundChangeRuleMappingList = v
	return s
}

func (s *PricingResponseBodyDataSolution) SetSolutionId(v string) *PricingResponseBodyDataSolution {
	s.SolutionId = &v
	return s
}

func (s *PricingResponseBodyDataSolution) Validate() error {
	if s.JourneyList != nil {
		for _, item := range s.JourneyList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SegmentBaggageCheckInInfoList != nil {
		for _, item := range s.SegmentBaggageCheckInInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SegmentBaggageMappingList != nil {
		for _, item := range s.SegmentBaggageMappingList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SegmentRefundChangeRuleMappingList != nil {
		for _, item := range s.SegmentRefundChangeRuleMappingList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PricingResponseBodyDataSolutionJourneyList struct {
	// segment list
	SegmentList []*PricingResponseBodyDataSolutionJourneyListSegmentList `json:"segment_list,omitempty" xml:"segment_list,omitempty" type:"Repeated"`
	// number of transfers
	//
	// example:
	//
	// 0
	TransferCount *int32 `json:"transfer_count,omitempty" xml:"transfer_count,omitempty"`
}

func (s PricingResponseBodyDataSolutionJourneyList) String() string {
	return dara.Prettify(s)
}

func (s PricingResponseBodyDataSolutionJourneyList) GoString() string {
	return s.String()
}

func (s *PricingResponseBodyDataSolutionJourneyList) GetSegmentList() []*PricingResponseBodyDataSolutionJourneyListSegmentList {
	return s.SegmentList
}

func (s *PricingResponseBodyDataSolutionJourneyList) GetTransferCount() *int32 {
	return s.TransferCount
}

func (s *PricingResponseBodyDataSolutionJourneyList) SetSegmentList(v []*PricingResponseBodyDataSolutionJourneyListSegmentList) *PricingResponseBodyDataSolutionJourneyList {
	s.SegmentList = v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyList) SetTransferCount(v int32) *PricingResponseBodyDataSolutionJourneyList {
	s.TransferCount = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyList) Validate() error {
	if s.SegmentList != nil {
		for _, item := range s.SegmentList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PricingResponseBodyDataSolutionJourneyListSegmentList struct {
	// arrival airport code (capitalized)
	//
	// example:
	//
	// MFM
	ArrivalAirport *string `json:"arrival_airport,omitempty" xml:"arrival_airport,omitempty"`
	// arrival city code (capitalized)
	//
	// example:
	//
	// MFM
	ArrivalCity *string `json:"arrival_city,omitempty" xml:"arrival_city,omitempty"`
	// arrival terminal
	//
	// example:
	//
	// T1
	ArrivalTerminal *string `json:"arrival_terminal,omitempty" xml:"arrival_terminal,omitempty"`
	// arrival time in string format (yyyy-MM-dd HH:mm:ss)
	//
	// example:
	//
	// 2023-03-10 10:40:00
	ArrivalTime *string `json:"arrival_time,omitempty" xml:"arrival_time,omitempty"`
	// available seats (for reference only)
	//
	// example:
	//
	// 7
	Availability *string `json:"availability,omitempty" xml:"availability,omitempty"`
	// RBD
	//
	// example:
	//
	// V
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// cabin class
	//
	// example:
	//
	// ALL_CABIN
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// code share or not
	//
	// example:
	//
	// false
	CodeShare *bool `json:"code_share,omitempty" xml:"code_share,omitempty"`
	// departure airport code (capitalized)
	//
	// example:
	//
	// PVG
	DepartureAirport *string `json:"departure_airport,omitempty" xml:"departure_airport,omitempty"`
	// departure city code (capitalized)
	//
	// example:
	//
	// SHA
	DepartureCity *string `json:"departure_city,omitempty" xml:"departure_city,omitempty"`
	// departure terminal
	//
	// example:
	//
	// T2
	DepartureTerminal *string `json:"departure_terminal,omitempty" xml:"departure_terminal,omitempty"`
	// departure time in string format (yyyy-MM-dd HH:mm:ss)
	//
	// example:
	//
	// 2023-03-10 07:55:00
	DepartureTime *string `json:"departure_time,omitempty" xml:"departure_time,omitempty"`
	// equipment type
	//
	// example:
	//
	// 32Q
	EquipType *string `json:"equip_type,omitempty" xml:"equip_type,omitempty"`
	// flight time, unit: minute
	//
	// example:
	//
	// 165
	FlightDuration *int32 `json:"flight_duration,omitempty" xml:"flight_duration,omitempty"`
	// marketing airline code (eg: KA)
	//
	// example:
	//
	// HO
	MarketingAirline *string `json:"marketing_airline,omitempty" xml:"marketing_airline,omitempty"`
	// marketing airline flight no. (eg: KA5809)
	//
	// example:
	//
	// HO1295
	MarketingFlightNo *string `json:"marketing_flight_no,omitempty" xml:"marketing_flight_no,omitempty"`
	// marketing airline flight no. (eg: 5809)
	//
	// example:
	//
	// 1295
	MarketingFlightNoInt *int32 `json:"marketing_flight_no_int,omitempty" xml:"marketing_flight_no_int,omitempty"`
	// operating airline code (eg: CX)
	//
	// example:
	//
	// HO
	OperatingAirline *string `json:"operating_airline,omitempty" xml:"operating_airline,omitempty"`
	// operating airline flight no. (eg: CX601)
	//
	// example:
	//
	// HO1295
	OperatingFlightNo *string `json:"operating_flight_no,omitempty" xml:"operating_flight_no,omitempty"`
	// segment ID format: flight no.+departure airport[IATA airport code]+arrival airport[IATA airport code]+departure time(MMdd)
	//
	// example:
	//
	// HO1295-PVG-MFM-20230310
	SegmentId *string `json:"segment_id,omitempty" xml:"segment_id,omitempty"`
	// stop city list. when stop_quantity > 1 , use “,” for seperation
	//
	// example:
	//
	// MFM,PVG
	StopCityList *string `json:"stop_city_list,omitempty" xml:"stop_city_list,omitempty"`
	// number of stops
	//
	// example:
	//
	// 0
	StopQuantity *int32 `json:"stop_quantity,omitempty" xml:"stop_quantity,omitempty"`
}

func (s PricingResponseBodyDataSolutionJourneyListSegmentList) String() string {
	return dara.Prettify(s)
}

func (s PricingResponseBodyDataSolutionJourneyListSegmentList) GoString() string {
	return s.String()
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) GetArrivalAirport() *string {
	return s.ArrivalAirport
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) GetArrivalCity() *string {
	return s.ArrivalCity
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) GetArrivalTerminal() *string {
	return s.ArrivalTerminal
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) GetArrivalTime() *string {
	return s.ArrivalTime
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) GetAvailability() *string {
	return s.Availability
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) GetCabin() *string {
	return s.Cabin
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) GetCabinClass() *string {
	return s.CabinClass
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) GetCodeShare() *bool {
	return s.CodeShare
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) GetDepartureAirport() *string {
	return s.DepartureAirport
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) GetDepartureCity() *string {
	return s.DepartureCity
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) GetDepartureTerminal() *string {
	return s.DepartureTerminal
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) GetDepartureTime() *string {
	return s.DepartureTime
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) GetEquipType() *string {
	return s.EquipType
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) GetFlightDuration() *int32 {
	return s.FlightDuration
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) GetMarketingAirline() *string {
	return s.MarketingAirline
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) GetMarketingFlightNo() *string {
	return s.MarketingFlightNo
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) GetMarketingFlightNoInt() *int32 {
	return s.MarketingFlightNoInt
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) GetOperatingAirline() *string {
	return s.OperatingAirline
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) GetOperatingFlightNo() *string {
	return s.OperatingFlightNo
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) GetSegmentId() *string {
	return s.SegmentId
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) GetStopCityList() *string {
	return s.StopCityList
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) GetStopQuantity() *int32 {
	return s.StopQuantity
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetArrivalAirport(v string) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.ArrivalAirport = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetArrivalCity(v string) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.ArrivalCity = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetArrivalTerminal(v string) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.ArrivalTerminal = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetArrivalTime(v string) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.ArrivalTime = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetAvailability(v string) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.Availability = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetCabin(v string) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.Cabin = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetCabinClass(v string) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.CabinClass = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetCodeShare(v bool) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.CodeShare = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetDepartureAirport(v string) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.DepartureAirport = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetDepartureCity(v string) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.DepartureCity = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetDepartureTerminal(v string) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.DepartureTerminal = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetDepartureTime(v string) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.DepartureTime = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetEquipType(v string) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.EquipType = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetFlightDuration(v int32) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.FlightDuration = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetMarketingAirline(v string) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.MarketingAirline = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetMarketingFlightNo(v string) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.MarketingFlightNo = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetMarketingFlightNoInt(v int32) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.MarketingFlightNoInt = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetOperatingAirline(v string) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.OperatingAirline = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetOperatingFlightNo(v string) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.OperatingFlightNo = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetSegmentId(v string) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.SegmentId = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetStopCityList(v string) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.StopCityList = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) SetStopQuantity(v int32) *PricingResponseBodyDataSolutionJourneyListSegmentList {
	s.StopQuantity = &v
	return s
}

func (s *PricingResponseBodyDataSolutionJourneyListSegmentList) Validate() error {
	return dara.Validate(s)
}

type PricingResponseBodyDataSolutionSegmentBaggageCheckInInfoList struct {
	// through check-in baggage policy type
	//
	// 1. baggage through check-in between segments
	//
	// 2. baggage re-check-in needed between segments
	//
	// 3. baggage through check-in at stop city ( applies for stop flight )
	//
	// 4. baggage re-checkin needed at stop city ( applies for stop flight )
	//
	// example:
	//
	// 1
	LuggageDirectInfoType *int32 `json:"luggage_direct_info_type,omitempty" xml:"luggage_direct_info_type,omitempty"`
	// segment id list. all the listed segment ids share the same baggage through check-in policy
	SegmentIdList []*string `json:"segment_id_list,omitempty" xml:"segment_id_list,omitempty" type:"Repeated"`
}

func (s PricingResponseBodyDataSolutionSegmentBaggageCheckInInfoList) String() string {
	return dara.Prettify(s)
}

func (s PricingResponseBodyDataSolutionSegmentBaggageCheckInInfoList) GoString() string {
	return s.String()
}

func (s *PricingResponseBodyDataSolutionSegmentBaggageCheckInInfoList) GetLuggageDirectInfoType() *int32 {
	return s.LuggageDirectInfoType
}

func (s *PricingResponseBodyDataSolutionSegmentBaggageCheckInInfoList) GetSegmentIdList() []*string {
	return s.SegmentIdList
}

func (s *PricingResponseBodyDataSolutionSegmentBaggageCheckInInfoList) SetLuggageDirectInfoType(v int32) *PricingResponseBodyDataSolutionSegmentBaggageCheckInInfoList {
	s.LuggageDirectInfoType = &v
	return s
}

func (s *PricingResponseBodyDataSolutionSegmentBaggageCheckInInfoList) SetSegmentIdList(v []*string) *PricingResponseBodyDataSolutionSegmentBaggageCheckInInfoList {
	s.SegmentIdList = v
	return s
}

func (s *PricingResponseBodyDataSolutionSegmentBaggageCheckInInfoList) Validate() error {
	return dara.Validate(s)
}

type PricingResponseBodyDataSolutionSegmentBaggageMappingList struct {
	// baggage rule mapping, key is passenger type, value is baggage allowance details
	PassengerBaggageAllowanceMapping map[string]*DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue `json:"passenger_baggage_allowance_mapping,omitempty" xml:"passenger_baggage_allowance_mapping,omitempty"`
	// segment id list all the listed segment id share the same baggage rule
	SegmentIdList []*string `json:"segment_id_list,omitempty" xml:"segment_id_list,omitempty" type:"Repeated"`
}

func (s PricingResponseBodyDataSolutionSegmentBaggageMappingList) String() string {
	return dara.Prettify(s)
}

func (s PricingResponseBodyDataSolutionSegmentBaggageMappingList) GoString() string {
	return s.String()
}

func (s *PricingResponseBodyDataSolutionSegmentBaggageMappingList) GetPassengerBaggageAllowanceMapping() map[string]*DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	return s.PassengerBaggageAllowanceMapping
}

func (s *PricingResponseBodyDataSolutionSegmentBaggageMappingList) GetSegmentIdList() []*string {
	return s.SegmentIdList
}

func (s *PricingResponseBodyDataSolutionSegmentBaggageMappingList) SetPassengerBaggageAllowanceMapping(v map[string]*DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) *PricingResponseBodyDataSolutionSegmentBaggageMappingList {
	s.PassengerBaggageAllowanceMapping = v
	return s
}

func (s *PricingResponseBodyDataSolutionSegmentBaggageMappingList) SetSegmentIdList(v []*string) *PricingResponseBodyDataSolutionSegmentBaggageMappingList {
	s.SegmentIdList = v
	return s
}

func (s *PricingResponseBodyDataSolutionSegmentBaggageMappingList) Validate() error {
	return dara.Validate(s)
}

type PricingResponseBodyDataSolutionSegmentRefundChangeRuleMappingList struct {
	// change and refund policy mapping, key is passenger type, value is change and refund policy detail
	RefundChangeRuleMap map[string]*DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue `json:"refund_change_rule_map,omitempty" xml:"refund_change_rule_map,omitempty"`
	// segment id list. all the listed segment ids share the same change and refund policy
	SegmentIdList []*string `json:"segment_id_list,omitempty" xml:"segment_id_list,omitempty" type:"Repeated"`
}

func (s PricingResponseBodyDataSolutionSegmentRefundChangeRuleMappingList) String() string {
	return dara.Prettify(s)
}

func (s PricingResponseBodyDataSolutionSegmentRefundChangeRuleMappingList) GoString() string {
	return s.String()
}

func (s *PricingResponseBodyDataSolutionSegmentRefundChangeRuleMappingList) GetRefundChangeRuleMap() map[string]*DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue {
	return s.RefundChangeRuleMap
}

func (s *PricingResponseBodyDataSolutionSegmentRefundChangeRuleMappingList) GetSegmentIdList() []*string {
	return s.SegmentIdList
}

func (s *PricingResponseBodyDataSolutionSegmentRefundChangeRuleMappingList) SetRefundChangeRuleMap(v map[string]*DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue) *PricingResponseBodyDataSolutionSegmentRefundChangeRuleMappingList {
	s.RefundChangeRuleMap = v
	return s
}

func (s *PricingResponseBodyDataSolutionSegmentRefundChangeRuleMappingList) SetSegmentIdList(v []*string) *PricingResponseBodyDataSolutionSegmentRefundChangeRuleMappingList {
	s.SegmentIdList = v
	return s
}

func (s *PricingResponseBodyDataSolutionSegmentRefundChangeRuleMappingList) Validate() error {
	return dara.Validate(s)
}
