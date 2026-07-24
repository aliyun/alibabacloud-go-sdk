// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOrderDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OrderDetailResponseBody
	GetRequestId() *string
	SetData(v *OrderDetailResponseBodyData) *OrderDetailResponseBody
	GetData() *OrderDetailResponseBodyData
	SetErrorCode(v string) *OrderDetailResponseBody
	GetErrorCode() *string
	SetErrorData(v interface{}) *OrderDetailResponseBody
	GetErrorData() interface{}
	SetErrorMsg(v string) *OrderDetailResponseBody
	GetErrorMsg() *string
	SetStatus(v int32) *OrderDetailResponseBody
	GetStatus() *int32
	SetSuccess(v bool) *OrderDetailResponseBody
	GetSuccess() *bool
}

type OrderDetailResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The data returned for a successful request.
	Data *OrderDetailResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The business error code.
	//
	// example:
	//
	// null
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// The data returned with the error.
	//
	// example:
	//
	// null
	ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
	// The error message.
	//
	// example:
	//
	// null
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// The HTTP status code. The value is always 200 for successful requests.
	//
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s OrderDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OrderDetailResponseBody) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OrderDetailResponseBody) GetData() *OrderDetailResponseBodyData {
	return s.Data
}

func (s *OrderDetailResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *OrderDetailResponseBody) GetErrorData() interface{} {
	return s.ErrorData
}

func (s *OrderDetailResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *OrderDetailResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *OrderDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OrderDetailResponseBody) SetRequestId(v string) *OrderDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *OrderDetailResponseBody) SetData(v *OrderDetailResponseBodyData) *OrderDetailResponseBody {
	s.Data = v
	return s
}

func (s *OrderDetailResponseBody) SetErrorCode(v string) *OrderDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *OrderDetailResponseBody) SetErrorData(v interface{}) *OrderDetailResponseBody {
	s.ErrorData = v
	return s
}

func (s *OrderDetailResponseBody) SetErrorMsg(v string) *OrderDetailResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *OrderDetailResponseBody) SetStatus(v int32) *OrderDetailResponseBody {
	s.Status = &v
	return s
}

func (s *OrderDetailResponseBody) SetSuccess(v bool) *OrderDetailResponseBody {
	s.Success = &v
	return s
}

func (s *OrderDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OrderDetailResponseBodyData struct {
	// The ancillary product fulfillment details.
	AncillaryItemDetailList []*OrderDetailResponseBodyDataAncillaryItemDetailList `json:"ancillary_item_detail_list,omitempty" xml:"ancillary_item_detail_list,omitempty" type:"Repeated"`
	// The mapping of passenger types to baggage rules.
	BaggageAllowanceMap map[string]*DataBaggageAllowanceMapValue `json:"baggage_allowance_map,omitempty" xml:"baggage_allowance_map,omitempty"`
	// The booking time (order creation time). The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1677210784000
	BookTime *int64 `json:"book_time,omitempty" xml:"book_time,omitempty"`
	// The flight ticket fulfillment details.
	FlightItemDetailList []*OrderDetailResponseBodyDataFlightItemDetailList `json:"flight_item_detail_list,omitempty" xml:"flight_item_detail_list,omitempty" type:"Repeated"`
	// The order number.
	//
	// example:
	//
	// 4966***617111
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
	// The order status. Valid values:
	//
	// - 1: Booking in progress.
	//
	// - 2: Booking succeeded.
	//
	// - 3: Order paid.
	//
	// - 4: Order succeeded.
	//
	// - 5: Order closed.
	//
	// example:
	//
	// 4
	OrderStatus *int32 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// The external order number.
	//
	// example:
	//
	// x091-2023-0220-j-0001
	OutOrderNum *string `json:"out_order_num,omitempty" xml:"out_order_num,omitempty"`
	// The passenger list.
	PassengerList []*OrderDetailResponseBodyDataPassengerList `json:"passenger_list,omitempty" xml:"passenger_list,omitempty" type:"Repeated"`
	// The payment status. Valid values:
	//
	// - 0: Initialized.
	//
	// - 1: Created.
	//
	// - 2: Payment succeeded.
	//
	// - 4: Transaction closed.
	//
	// example:
	//
	// 2
	PayStatus *int32 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// The payment time. The value is a 13-digit timestamp.
	//
	// example:
	//
	// 1677210788000
	PayTime *int64 `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// The discount amount. Unit: CNY.
	//
	// example:
	//
	// 10
	PromotionPrice *float64 `json:"promotion_price,omitempty" xml:"promotion_price,omitempty"`
	// The actual payment amount. Unit: CNY.
	//
	// example:
	//
	// 3000
	RealPayPrice *float64 `json:"real_pay_price,omitempty" xml:"real_pay_price,omitempty"`
	// The mapping of passenger types to refund and change rules.
	RefundChangeRuleMap map[string]*DataRefundChangeRuleMapValue `json:"refund_change_rule_map,omitempty" xml:"refund_change_rule_map,omitempty"`
	// The buyer nickname.
	//
	// example:
	//
	// nick
	SessionNick *string `json:"session_nick,omitempty" xml:"session_nick,omitempty"`
	// The flight information.
	Solution *OrderDetailResponseBodyDataSolution `json:"solution,omitempty" xml:"solution,omitempty" type:"Struct"`
	// The ticketing time. The value is a 13-digit timestamp. This parameter has a value only after ticketing is complete.
	//
	// example:
	//
	// 1677210786000
	SucceedTime *int64 `json:"succeed_time,omitempty" xml:"succeed_time,omitempty"`
	// The total order price. Unit: CNY.
	//
	// example:
	//
	// 3000
	TotalPrice *float64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// The transaction number.
	//
	// example:
	//
	// hkduendkd-2023-dj0
	TransactionNo *string `json:"transaction_no,omitempty" xml:"transaction_no,omitempty"`
}

func (s OrderDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s OrderDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBodyData) GetAncillaryItemDetailList() []*OrderDetailResponseBodyDataAncillaryItemDetailList {
	return s.AncillaryItemDetailList
}

func (s *OrderDetailResponseBodyData) GetBaggageAllowanceMap() map[string]*DataBaggageAllowanceMapValue {
	return s.BaggageAllowanceMap
}

func (s *OrderDetailResponseBodyData) GetBookTime() *int64 {
	return s.BookTime
}

func (s *OrderDetailResponseBodyData) GetFlightItemDetailList() []*OrderDetailResponseBodyDataFlightItemDetailList {
	return s.FlightItemDetailList
}

func (s *OrderDetailResponseBodyData) GetOrderNum() *int64 {
	return s.OrderNum
}

func (s *OrderDetailResponseBodyData) GetOrderStatus() *int32 {
	return s.OrderStatus
}

func (s *OrderDetailResponseBodyData) GetOutOrderNum() *string {
	return s.OutOrderNum
}

func (s *OrderDetailResponseBodyData) GetPassengerList() []*OrderDetailResponseBodyDataPassengerList {
	return s.PassengerList
}

func (s *OrderDetailResponseBodyData) GetPayStatus() *int32 {
	return s.PayStatus
}

func (s *OrderDetailResponseBodyData) GetPayTime() *int64 {
	return s.PayTime
}

func (s *OrderDetailResponseBodyData) GetPromotionPrice() *float64 {
	return s.PromotionPrice
}

func (s *OrderDetailResponseBodyData) GetRealPayPrice() *float64 {
	return s.RealPayPrice
}

func (s *OrderDetailResponseBodyData) GetRefundChangeRuleMap() map[string]*DataRefundChangeRuleMapValue {
	return s.RefundChangeRuleMap
}

func (s *OrderDetailResponseBodyData) GetSessionNick() *string {
	return s.SessionNick
}

func (s *OrderDetailResponseBodyData) GetSolution() *OrderDetailResponseBodyDataSolution {
	return s.Solution
}

func (s *OrderDetailResponseBodyData) GetSucceedTime() *int64 {
	return s.SucceedTime
}

func (s *OrderDetailResponseBodyData) GetTotalPrice() *float64 {
	return s.TotalPrice
}

func (s *OrderDetailResponseBodyData) GetTransactionNo() *string {
	return s.TransactionNo
}

func (s *OrderDetailResponseBodyData) SetAncillaryItemDetailList(v []*OrderDetailResponseBodyDataAncillaryItemDetailList) *OrderDetailResponseBodyData {
	s.AncillaryItemDetailList = v
	return s
}

func (s *OrderDetailResponseBodyData) SetBaggageAllowanceMap(v map[string]*DataBaggageAllowanceMapValue) *OrderDetailResponseBodyData {
	s.BaggageAllowanceMap = v
	return s
}

func (s *OrderDetailResponseBodyData) SetBookTime(v int64) *OrderDetailResponseBodyData {
	s.BookTime = &v
	return s
}

func (s *OrderDetailResponseBodyData) SetFlightItemDetailList(v []*OrderDetailResponseBodyDataFlightItemDetailList) *OrderDetailResponseBodyData {
	s.FlightItemDetailList = v
	return s
}

func (s *OrderDetailResponseBodyData) SetOrderNum(v int64) *OrderDetailResponseBodyData {
	s.OrderNum = &v
	return s
}

func (s *OrderDetailResponseBodyData) SetOrderStatus(v int32) *OrderDetailResponseBodyData {
	s.OrderStatus = &v
	return s
}

func (s *OrderDetailResponseBodyData) SetOutOrderNum(v string) *OrderDetailResponseBodyData {
	s.OutOrderNum = &v
	return s
}

func (s *OrderDetailResponseBodyData) SetPassengerList(v []*OrderDetailResponseBodyDataPassengerList) *OrderDetailResponseBodyData {
	s.PassengerList = v
	return s
}

func (s *OrderDetailResponseBodyData) SetPayStatus(v int32) *OrderDetailResponseBodyData {
	s.PayStatus = &v
	return s
}

func (s *OrderDetailResponseBodyData) SetPayTime(v int64) *OrderDetailResponseBodyData {
	s.PayTime = &v
	return s
}

func (s *OrderDetailResponseBodyData) SetPromotionPrice(v float64) *OrderDetailResponseBodyData {
	s.PromotionPrice = &v
	return s
}

func (s *OrderDetailResponseBodyData) SetRealPayPrice(v float64) *OrderDetailResponseBodyData {
	s.RealPayPrice = &v
	return s
}

func (s *OrderDetailResponseBodyData) SetRefundChangeRuleMap(v map[string]*DataRefundChangeRuleMapValue) *OrderDetailResponseBodyData {
	s.RefundChangeRuleMap = v
	return s
}

func (s *OrderDetailResponseBodyData) SetSessionNick(v string) *OrderDetailResponseBodyData {
	s.SessionNick = &v
	return s
}

func (s *OrderDetailResponseBodyData) SetSolution(v *OrderDetailResponseBodyDataSolution) *OrderDetailResponseBodyData {
	s.Solution = v
	return s
}

func (s *OrderDetailResponseBodyData) SetSucceedTime(v int64) *OrderDetailResponseBodyData {
	s.SucceedTime = &v
	return s
}

func (s *OrderDetailResponseBodyData) SetTotalPrice(v float64) *OrderDetailResponseBodyData {
	s.TotalPrice = &v
	return s
}

func (s *OrderDetailResponseBodyData) SetTransactionNo(v string) *OrderDetailResponseBodyData {
	s.TransactionNo = &v
	return s
}

func (s *OrderDetailResponseBodyData) Validate() error {
	if s.AncillaryItemDetailList != nil {
		for _, item := range s.AncillaryItemDetailList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FlightItemDetailList != nil {
		for _, item := range s.FlightItemDetailList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PassengerList != nil {
		for _, item := range s.PassengerList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Solution != nil {
		if err := s.Solution.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OrderDetailResponseBodyDataAncillaryItemDetailList struct {
	// The ancillary product details.
	Ancillary *OrderDetailResponseBodyDataAncillaryItemDetailListAncillary `json:"ancillary,omitempty" xml:"ancillary,omitempty" type:"Struct"`
	// The passenger information.
	Passenger *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger `json:"passenger,omitempty" xml:"passenger,omitempty" type:"Struct"`
	// The segment IDs to which the ancillary product applies.
	SegmentIdList []*string `json:"segment_id_list,omitempty" xml:"segment_id_list,omitempty" type:"Repeated"`
}

func (s OrderDetailResponseBodyDataAncillaryItemDetailList) String() string {
	return dara.Prettify(s)
}

func (s OrderDetailResponseBodyDataAncillaryItemDetailList) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailList) GetAncillary() *OrderDetailResponseBodyDataAncillaryItemDetailListAncillary {
	return s.Ancillary
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailList) GetPassenger() *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger {
	return s.Passenger
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailList) GetSegmentIdList() []*string {
	return s.SegmentIdList
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailList) SetAncillary(v *OrderDetailResponseBodyDataAncillaryItemDetailListAncillary) *OrderDetailResponseBodyDataAncillaryItemDetailList {
	s.Ancillary = v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailList) SetPassenger(v *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger) *OrderDetailResponseBodyDataAncillaryItemDetailList {
	s.Passenger = v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailList) SetSegmentIdList(v []*string) *OrderDetailResponseBodyDataAncillaryItemDetailList {
	s.SegmentIdList = v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailList) Validate() error {
	if s.Ancillary != nil {
		if err := s.Ancillary.Validate(); err != nil {
			return err
		}
	}
	if s.Passenger != nil {
		if err := s.Passenger.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OrderDetailResponseBodyDataAncillaryItemDetailListAncillary struct {
	// The ancillary product ID.
	//
	// example:
	//
	// MDY2NTAxLCJleHAiOjE2NxNzM3MDEsIm5ix
	AncillaryId *string `json:"ancillary_id,omitempty" xml:"ancillary_id,omitempty"`
	// The ancillary product type. Currently supported value: 4 (paid baggage).
	//
	// example:
	//
	// 4
	AncillaryType *int32 `json:"ancillary_type,omitempty" xml:"ancillary_type,omitempty"`
	// The baggage ancillary details.
	BaggageAncillary *OrderDetailResponseBodyDataAncillaryItemDetailListAncillaryBaggageAncillary `json:"baggage_ancillary,omitempty" xml:"baggage_ancillary,omitempty" type:"Struct"`
}

func (s OrderDetailResponseBodyDataAncillaryItemDetailListAncillary) String() string {
	return dara.Prettify(s)
}

func (s OrderDetailResponseBodyDataAncillaryItemDetailListAncillary) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListAncillary) GetAncillaryId() *string {
	return s.AncillaryId
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListAncillary) GetAncillaryType() *int32 {
	return s.AncillaryType
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListAncillary) GetBaggageAncillary() *OrderDetailResponseBodyDataAncillaryItemDetailListAncillaryBaggageAncillary {
	return s.BaggageAncillary
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListAncillary) SetAncillaryId(v string) *OrderDetailResponseBodyDataAncillaryItemDetailListAncillary {
	s.AncillaryId = &v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListAncillary) SetAncillaryType(v int32) *OrderDetailResponseBodyDataAncillaryItemDetailListAncillary {
	s.AncillaryType = &v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListAncillary) SetBaggageAncillary(v *OrderDetailResponseBodyDataAncillaryItemDetailListAncillaryBaggageAncillary) *OrderDetailResponseBodyDataAncillaryItemDetailListAncillary {
	s.BaggageAncillary = v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListAncillary) Validate() error {
	if s.BaggageAncillary != nil {
		if err := s.BaggageAncillary.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OrderDetailResponseBodyDataAncillaryItemDetailListAncillaryBaggageAncillary struct {
	// The number of baggage pieces. Valid values: 3, 2, 1, 0, and -2. A value of -2 indicates weight-based calculation.
	//
	// example:
	//
	// 0
	BaggageAmount *int32 `json:"baggage_amount,omitempty" xml:"baggage_amount,omitempty"`
	// The baggage weight, ranging from 0 to 50. When isAllWeight is set to true, this value represents the total weight of all pieces.
	//
	// example:
	//
	// 0
	BaggageWeight *int32 `json:"baggage_weight,omitempty" xml:"baggage_weight,omitempty"`
	// The baggage weight unit.
	//
	// example:
	//
	// KG
	BaggageWeightUnit *string `json:"baggage_weight_unit,omitempty" xml:"baggage_weight_unit,omitempty"`
	// Indicates whether the weight represents the total weight of all baggage pieces.
	IsAllWeight *bool `json:"is_all_weight,omitempty" xml:"is_all_weight,omitempty"`
	// The total price.
	//
	// example:
	//
	// 10.0
	Price *float64 `json:"price,omitempty" xml:"price,omitempty"`
}

func (s OrderDetailResponseBodyDataAncillaryItemDetailListAncillaryBaggageAncillary) String() string {
	return dara.Prettify(s)
}

func (s OrderDetailResponseBodyDataAncillaryItemDetailListAncillaryBaggageAncillary) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListAncillaryBaggageAncillary) GetBaggageAmount() *int32 {
	return s.BaggageAmount
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListAncillaryBaggageAncillary) GetBaggageWeight() *int32 {
	return s.BaggageWeight
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListAncillaryBaggageAncillary) GetBaggageWeightUnit() *string {
	return s.BaggageWeightUnit
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListAncillaryBaggageAncillary) GetIsAllWeight() *bool {
	return s.IsAllWeight
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListAncillaryBaggageAncillary) GetPrice() *float64 {
	return s.Price
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListAncillaryBaggageAncillary) SetBaggageAmount(v int32) *OrderDetailResponseBodyDataAncillaryItemDetailListAncillaryBaggageAncillary {
	s.BaggageAmount = &v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListAncillaryBaggageAncillary) SetBaggageWeight(v int32) *OrderDetailResponseBodyDataAncillaryItemDetailListAncillaryBaggageAncillary {
	s.BaggageWeight = &v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListAncillaryBaggageAncillary) SetBaggageWeightUnit(v string) *OrderDetailResponseBodyDataAncillaryItemDetailListAncillaryBaggageAncillary {
	s.BaggageWeightUnit = &v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListAncillaryBaggageAncillary) SetIsAllWeight(v bool) *OrderDetailResponseBodyDataAncillaryItemDetailListAncillaryBaggageAncillary {
	s.IsAllWeight = &v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListAncillaryBaggageAncillary) SetPrice(v float64) *OrderDetailResponseBodyDataAncillaryItemDetailListAncillaryBaggageAncillary {
	s.Price = &v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListAncillaryBaggageAncillary) Validate() error {
	return dara.Validate(s)
}

type OrderDetailResponseBodyDataAncillaryItemDetailListPassenger struct {
	// The date of birth in yyyyMMdd format.
	//
	// example:
	//
	// 20020301
	Birthday *string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// The credential information.
	Credential *OrderDetailResponseBodyDataAncillaryItemDetailListPassengerCredential `json:"credential,omitempty" xml:"credential,omitempty" type:"Struct"`
	// The first name.
	//
	// example:
	//
	// SAN
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// The gender. Valid values:
	//
	// - 0: MALE.
	//
	// - 1: FEMALE.
	//
	// example:
	//
	// 0
	Gender *int32 `json:"gender,omitempty" xml:"gender,omitempty"`
	// The last name.
	//
	// example:
	//
	// ZHANG
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	// The country code of the mobile phone number.
	//
	// example:
	//
	// 86
	MobileCountryCode *string `json:"mobile_country_code,omitempty" xml:"mobile_country_code,omitempty"`
	// The mobile phone number.
	//
	// example:
	//
	// 183******96
	MobilePhoneNumber *string `json:"mobile_phone_number,omitempty" xml:"mobile_phone_number,omitempty"`
	// The two-letter nationality code.
	//
	// example:
	//
	// CN
	Nationality *string `json:"nationality,omitempty" xml:"nationality,omitempty"`
	// The passenger type. Valid values:
	//
	// - 0: adult.
	//
	// - 1: child.
	//
	// - 8: infant.
	//
	// example:
	//
	// 0
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s OrderDetailResponseBodyDataAncillaryItemDetailListPassenger) String() string {
	return dara.Prettify(s)
}

func (s OrderDetailResponseBodyDataAncillaryItemDetailListPassenger) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger) GetBirthday() *string {
	return s.Birthday
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger) GetCredential() *OrderDetailResponseBodyDataAncillaryItemDetailListPassengerCredential {
	return s.Credential
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger) GetFirstName() *string {
	return s.FirstName
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger) GetGender() *int32 {
	return s.Gender
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger) GetLastName() *string {
	return s.LastName
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger) GetMobileCountryCode() *string {
	return s.MobileCountryCode
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger) GetMobilePhoneNumber() *string {
	return s.MobilePhoneNumber
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger) GetNationality() *string {
	return s.Nationality
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger) GetType() *int32 {
	return s.Type
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger) SetBirthday(v string) *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger {
	s.Birthday = &v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger) SetCredential(v *OrderDetailResponseBodyDataAncillaryItemDetailListPassengerCredential) *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger {
	s.Credential = v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger) SetFirstName(v string) *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger {
	s.FirstName = &v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger) SetGender(v int32) *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger {
	s.Gender = &v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger) SetLastName(v string) *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger {
	s.LastName = &v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger) SetMobileCountryCode(v string) *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger {
	s.MobileCountryCode = &v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger) SetMobilePhoneNumber(v string) *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger {
	s.MobilePhoneNumber = &v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger) SetNationality(v string) *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger {
	s.Nationality = &v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger) SetType(v int32) *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger {
	s.Type = &v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassenger) Validate() error {
	if s.Credential != nil {
		if err := s.Credential.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OrderDetailResponseBodyDataAncillaryItemDetailListPassengerCredential struct {
	// The place of issuance, represented as a two-letter code.
	//
	// example:
	//
	// CN
	CertIssuePlace *string `json:"cert_issue_place,omitempty" xml:"cert_issue_place,omitempty"`
	// The credential number.
	//
	// example:
	//
	// E1***5674
	CredentialNum *string `json:"credential_num,omitempty" xml:"credential_num,omitempty"`
	// The credential type. Valid values:
	//
	// - 0: ID card.
	//
	// - 1: passport.
	//
	// - 4: Home Return Permit.
	//
	// - 5: Mainland Travel Permit for Taiwan Residents.
	//
	// - 6: Exit-Entry Permit for Hong Kong and Macao Residents.
	//
	// - 12: Taiwan Travel Permit.
	//
	// - 19: no credential.
	//
	// example:
	//
	// 1
	CredentialType *int32 `json:"credential_type,omitempty" xml:"credential_type,omitempty"`
	// The credential expiration date.
	//
	// example:
	//
	// 20290101
	ExpireDate *string `json:"expire_date,omitempty" xml:"expire_date,omitempty"`
}

func (s OrderDetailResponseBodyDataAncillaryItemDetailListPassengerCredential) String() string {
	return dara.Prettify(s)
}

func (s OrderDetailResponseBodyDataAncillaryItemDetailListPassengerCredential) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassengerCredential) GetCertIssuePlace() *string {
	return s.CertIssuePlace
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassengerCredential) GetCredentialNum() *string {
	return s.CredentialNum
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassengerCredential) GetCredentialType() *int32 {
	return s.CredentialType
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassengerCredential) GetExpireDate() *string {
	return s.ExpireDate
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassengerCredential) SetCertIssuePlace(v string) *OrderDetailResponseBodyDataAncillaryItemDetailListPassengerCredential {
	s.CertIssuePlace = &v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassengerCredential) SetCredentialNum(v string) *OrderDetailResponseBodyDataAncillaryItemDetailListPassengerCredential {
	s.CredentialNum = &v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassengerCredential) SetCredentialType(v int32) *OrderDetailResponseBodyDataAncillaryItemDetailListPassengerCredential {
	s.CredentialType = &v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassengerCredential) SetExpireDate(v string) *OrderDetailResponseBodyDataAncillaryItemDetailListPassengerCredential {
	s.ExpireDate = &v
	return s
}

func (s *OrderDetailResponseBodyDataAncillaryItemDetailListPassengerCredential) Validate() error {
	return dara.Validate(s)
}

type OrderDetailResponseBodyDataFlightItemDetailList struct {
	// The list of bPnr values.
	BPnrList []*string `json:"b_pnr_list,omitempty" xml:"b_pnr_list,omitempty" type:"Repeated"`
	// The list of cPnr values.
	CPnrList []*string `json:"c_pnr_list,omitempty" xml:"c_pnr_list,omitempty" type:"Repeated"`
	// The passenger price information.
	FlightPrice *OrderDetailResponseBodyDataFlightItemDetailListFlightPrice `json:"flight_price,omitempty" xml:"flight_price,omitempty" type:"Struct"`
	// The list of segment-cabin information.
	FlightSegmentCabinRelation []*OrderDetailResponseBodyDataFlightItemDetailListFlightSegmentCabinRelation `json:"flight_segment_cabin_relation,omitempty" xml:"flight_segment_cabin_relation,omitempty" type:"Repeated"`
	// The passenger information.
	Passenger *OrderDetailResponseBodyDataFlightItemDetailListPassenger `json:"passenger,omitempty" xml:"passenger,omitempty" type:"Struct"`
	// The ticketing airline. Multiple ticketing airlines may be concatenated.
	//
	// example:
	//
	// HO
	TicketAirLine *string `json:"ticket_air_line,omitempty" xml:"ticket_air_line,omitempty"`
	// The list of ticket numbers.
	TicketNos []*string `json:"ticket_nos,omitempty" xml:"ticket_nos,omitempty" type:"Repeated"`
}

func (s OrderDetailResponseBodyDataFlightItemDetailList) String() string {
	return dara.Prettify(s)
}

func (s OrderDetailResponseBodyDataFlightItemDetailList) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBodyDataFlightItemDetailList) GetBPnrList() []*string {
	return s.BPnrList
}

func (s *OrderDetailResponseBodyDataFlightItemDetailList) GetCPnrList() []*string {
	return s.CPnrList
}

func (s *OrderDetailResponseBodyDataFlightItemDetailList) GetFlightPrice() *OrderDetailResponseBodyDataFlightItemDetailListFlightPrice {
	return s.FlightPrice
}

func (s *OrderDetailResponseBodyDataFlightItemDetailList) GetFlightSegmentCabinRelation() []*OrderDetailResponseBodyDataFlightItemDetailListFlightSegmentCabinRelation {
	return s.FlightSegmentCabinRelation
}

func (s *OrderDetailResponseBodyDataFlightItemDetailList) GetPassenger() *OrderDetailResponseBodyDataFlightItemDetailListPassenger {
	return s.Passenger
}

func (s *OrderDetailResponseBodyDataFlightItemDetailList) GetTicketAirLine() *string {
	return s.TicketAirLine
}

func (s *OrderDetailResponseBodyDataFlightItemDetailList) GetTicketNos() []*string {
	return s.TicketNos
}

func (s *OrderDetailResponseBodyDataFlightItemDetailList) SetBPnrList(v []*string) *OrderDetailResponseBodyDataFlightItemDetailList {
	s.BPnrList = v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailList) SetCPnrList(v []*string) *OrderDetailResponseBodyDataFlightItemDetailList {
	s.CPnrList = v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailList) SetFlightPrice(v *OrderDetailResponseBodyDataFlightItemDetailListFlightPrice) *OrderDetailResponseBodyDataFlightItemDetailList {
	s.FlightPrice = v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailList) SetFlightSegmentCabinRelation(v []*OrderDetailResponseBodyDataFlightItemDetailListFlightSegmentCabinRelation) *OrderDetailResponseBodyDataFlightItemDetailList {
	s.FlightSegmentCabinRelation = v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailList) SetPassenger(v *OrderDetailResponseBodyDataFlightItemDetailListPassenger) *OrderDetailResponseBodyDataFlightItemDetailList {
	s.Passenger = v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailList) SetTicketAirLine(v string) *OrderDetailResponseBodyDataFlightItemDetailList {
	s.TicketAirLine = &v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailList) SetTicketNos(v []*string) *OrderDetailResponseBodyDataFlightItemDetailList {
	s.TicketNos = v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailList) Validate() error {
	if s.FlightPrice != nil {
		if err := s.FlightPrice.Validate(); err != nil {
			return err
		}
	}
	if s.FlightSegmentCabinRelation != nil {
		for _, item := range s.FlightSegmentCabinRelation {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Passenger != nil {
		if err := s.Passenger.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OrderDetailResponseBodyDataFlightItemDetailListFlightPrice struct {
	// The selling price. Unit: CNY.
	//
	// example:
	//
	// 300
	SellPrice *float64 `json:"sell_price,omitempty" xml:"sell_price,omitempty"`
	// The tax.
	//
	// example:
	//
	// 10
	Tax *float64 `json:"tax,omitempty" xml:"tax,omitempty"`
}

func (s OrderDetailResponseBodyDataFlightItemDetailListFlightPrice) String() string {
	return dara.Prettify(s)
}

func (s OrderDetailResponseBodyDataFlightItemDetailListFlightPrice) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListFlightPrice) GetSellPrice() *float64 {
	return s.SellPrice
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListFlightPrice) GetTax() *float64 {
	return s.Tax
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListFlightPrice) SetSellPrice(v float64) *OrderDetailResponseBodyDataFlightItemDetailListFlightPrice {
	s.SellPrice = &v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListFlightPrice) SetTax(v float64) *OrderDetailResponseBodyDataFlightItemDetailListFlightPrice {
	s.Tax = &v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListFlightPrice) Validate() error {
	return dara.Validate(s)
}

type OrderDetailResponseBodyDataFlightItemDetailListFlightSegmentCabinRelation struct {
	// The cabin.
	//
	// example:
	//
	// V
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// The cabin class.
	//
	// example:
	//
	// Y
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// The cabin class description.
	//
	// example:
	//
	// economy class
	CabinClassName *string `json:"cabin_class_name,omitempty" xml:"cabin_class_name,omitempty"`
	// The number of available tickets.
	//
	// example:
	//
	// A
	CabinQuantity *string `json:"cabin_quantity,omitempty" xml:"cabin_quantity,omitempty"`
	// The segment ID.
	//
	// example:
	//
	// HO1295-PVG-MFM-20230310
	SegmentId *string `json:"segment_id,omitempty" xml:"segment_id,omitempty"`
}

func (s OrderDetailResponseBodyDataFlightItemDetailListFlightSegmentCabinRelation) String() string {
	return dara.Prettify(s)
}

func (s OrderDetailResponseBodyDataFlightItemDetailListFlightSegmentCabinRelation) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListFlightSegmentCabinRelation) GetCabin() *string {
	return s.Cabin
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListFlightSegmentCabinRelation) GetCabinClass() *string {
	return s.CabinClass
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListFlightSegmentCabinRelation) GetCabinClassName() *string {
	return s.CabinClassName
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListFlightSegmentCabinRelation) GetCabinQuantity() *string {
	return s.CabinQuantity
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListFlightSegmentCabinRelation) GetSegmentId() *string {
	return s.SegmentId
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListFlightSegmentCabinRelation) SetCabin(v string) *OrderDetailResponseBodyDataFlightItemDetailListFlightSegmentCabinRelation {
	s.Cabin = &v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListFlightSegmentCabinRelation) SetCabinClass(v string) *OrderDetailResponseBodyDataFlightItemDetailListFlightSegmentCabinRelation {
	s.CabinClass = &v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListFlightSegmentCabinRelation) SetCabinClassName(v string) *OrderDetailResponseBodyDataFlightItemDetailListFlightSegmentCabinRelation {
	s.CabinClassName = &v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListFlightSegmentCabinRelation) SetCabinQuantity(v string) *OrderDetailResponseBodyDataFlightItemDetailListFlightSegmentCabinRelation {
	s.CabinQuantity = &v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListFlightSegmentCabinRelation) SetSegmentId(v string) *OrderDetailResponseBodyDataFlightItemDetailListFlightSegmentCabinRelation {
	s.SegmentId = &v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListFlightSegmentCabinRelation) Validate() error {
	return dara.Validate(s)
}

type OrderDetailResponseBodyDataFlightItemDetailListPassenger struct {
	// The date of birth in yyyyMMdd format.
	//
	// example:
	//
	// 20020301
	Birthday *string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// The credential information.
	Credential *OrderDetailResponseBodyDataFlightItemDetailListPassengerCredential `json:"credential,omitempty" xml:"credential,omitempty" type:"Struct"`
	// The first name.
	//
	// example:
	//
	// SAN
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// The gender. Valid values:
	//
	// - 0: MALE.
	//
	// - 1: FEMALE.
	//
	// example:
	//
	// 0
	Gender *int32 `json:"gender,omitempty" xml:"gender,omitempty"`
	// The last name.
	//
	// example:
	//
	// ZHANG
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	// The country code of the mobile phone number.
	//
	// example:
	//
	// 86
	MobileCountryCode *string `json:"mobile_country_code,omitempty" xml:"mobile_country_code,omitempty"`
	// The mobile phone number.
	//
	// example:
	//
	// 183******96
	MobilePhoneNumber *string `json:"mobile_phone_number,omitempty" xml:"mobile_phone_number,omitempty"`
	// The two-letter nationality code.
	//
	// example:
	//
	// CN
	Nationality *string `json:"nationality,omitempty" xml:"nationality,omitempty"`
	// The passenger type. Valid values:
	//
	// - 0: adult.
	//
	// - 1: child.
	//
	// - 8: infant.
	//
	// example:
	//
	// 0
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s OrderDetailResponseBodyDataFlightItemDetailListPassenger) String() string {
	return dara.Prettify(s)
}

func (s OrderDetailResponseBodyDataFlightItemDetailListPassenger) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassenger) GetBirthday() *string {
	return s.Birthday
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassenger) GetCredential() *OrderDetailResponseBodyDataFlightItemDetailListPassengerCredential {
	return s.Credential
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassenger) GetFirstName() *string {
	return s.FirstName
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassenger) GetGender() *int32 {
	return s.Gender
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassenger) GetLastName() *string {
	return s.LastName
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassenger) GetMobileCountryCode() *string {
	return s.MobileCountryCode
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassenger) GetMobilePhoneNumber() *string {
	return s.MobilePhoneNumber
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassenger) GetNationality() *string {
	return s.Nationality
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassenger) GetType() *int32 {
	return s.Type
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassenger) SetBirthday(v string) *OrderDetailResponseBodyDataFlightItemDetailListPassenger {
	s.Birthday = &v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassenger) SetCredential(v *OrderDetailResponseBodyDataFlightItemDetailListPassengerCredential) *OrderDetailResponseBodyDataFlightItemDetailListPassenger {
	s.Credential = v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassenger) SetFirstName(v string) *OrderDetailResponseBodyDataFlightItemDetailListPassenger {
	s.FirstName = &v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassenger) SetGender(v int32) *OrderDetailResponseBodyDataFlightItemDetailListPassenger {
	s.Gender = &v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassenger) SetLastName(v string) *OrderDetailResponseBodyDataFlightItemDetailListPassenger {
	s.LastName = &v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassenger) SetMobileCountryCode(v string) *OrderDetailResponseBodyDataFlightItemDetailListPassenger {
	s.MobileCountryCode = &v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassenger) SetMobilePhoneNumber(v string) *OrderDetailResponseBodyDataFlightItemDetailListPassenger {
	s.MobilePhoneNumber = &v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassenger) SetNationality(v string) *OrderDetailResponseBodyDataFlightItemDetailListPassenger {
	s.Nationality = &v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassenger) SetType(v int32) *OrderDetailResponseBodyDataFlightItemDetailListPassenger {
	s.Type = &v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassenger) Validate() error {
	if s.Credential != nil {
		if err := s.Credential.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OrderDetailResponseBodyDataFlightItemDetailListPassengerCredential struct {
	// The place of issuance, represented as a two-letter code.
	//
	// example:
	//
	// CN
	CertIssuePlace *string `json:"cert_issue_place,omitempty" xml:"cert_issue_place,omitempty"`
	// The credential number.
	//
	// example:
	//
	// E1***5674
	CredentialNum *string `json:"credential_num,omitempty" xml:"credential_num,omitempty"`
	// The credential type. Valid values:
	//
	// - 0: ID card.
	//
	// - 1: passport.
	//
	// - 4: Home Return Permit.
	//
	// - 5: Mainland Travel Permit for Taiwan Residents.
	//
	// - 6: Exit-Entry Permit for Hong Kong and Macao Residents.
	//
	// - 12: Taiwan Travel Permit.
	//
	// - 19: no credential.
	//
	// example:
	//
	// 1
	CredentialType *int32 `json:"credential_type,omitempty" xml:"credential_type,omitempty"`
	// The credential expiration date.
	//
	// example:
	//
	// 20290101
	ExpireDate *string `json:"expire_date,omitempty" xml:"expire_date,omitempty"`
}

func (s OrderDetailResponseBodyDataFlightItemDetailListPassengerCredential) String() string {
	return dara.Prettify(s)
}

func (s OrderDetailResponseBodyDataFlightItemDetailListPassengerCredential) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassengerCredential) GetCertIssuePlace() *string {
	return s.CertIssuePlace
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassengerCredential) GetCredentialNum() *string {
	return s.CredentialNum
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassengerCredential) GetCredentialType() *int32 {
	return s.CredentialType
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassengerCredential) GetExpireDate() *string {
	return s.ExpireDate
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassengerCredential) SetCertIssuePlace(v string) *OrderDetailResponseBodyDataFlightItemDetailListPassengerCredential {
	s.CertIssuePlace = &v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassengerCredential) SetCredentialNum(v string) *OrderDetailResponseBodyDataFlightItemDetailListPassengerCredential {
	s.CredentialNum = &v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassengerCredential) SetCredentialType(v int32) *OrderDetailResponseBodyDataFlightItemDetailListPassengerCredential {
	s.CredentialType = &v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassengerCredential) SetExpireDate(v string) *OrderDetailResponseBodyDataFlightItemDetailListPassengerCredential {
	s.ExpireDate = &v
	return s
}

func (s *OrderDetailResponseBodyDataFlightItemDetailListPassengerCredential) Validate() error {
	return dara.Validate(s)
}

type OrderDetailResponseBodyDataPassengerList struct {
	// The date of birth in yyyyMMdd format.
	//
	// example:
	//
	// 20020301
	Birthday *string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// The credential information.
	Credential *OrderDetailResponseBodyDataPassengerListCredential `json:"credential,omitempty" xml:"credential,omitempty" type:"Struct"`
	// The first name.
	//
	// example:
	//
	// SAN
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// The gender. Valid values:
	//
	// - 0: MALE.
	//
	// - 1: FEMALE.
	//
	// example:
	//
	// 0
	Gender *int32 `json:"gender,omitempty" xml:"gender,omitempty"`
	// The last name.
	//
	// example:
	//
	// ZHANG
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
	// The country code of the mobile phone number.
	//
	// example:
	//
	// 86
	MobileCountryCode *string `json:"mobile_country_code,omitempty" xml:"mobile_country_code,omitempty"`
	// The mobile phone number.
	//
	// example:
	//
	// 183******96
	MobilePhoneNumber *string `json:"mobile_phone_number,omitempty" xml:"mobile_phone_number,omitempty"`
	// The two-letter nationality code.
	//
	// example:
	//
	// CN
	Nationality *string `json:"nationality,omitempty" xml:"nationality,omitempty"`
	// The passenger type. Valid values:
	//
	// - 0: adult.
	//
	// - 1: child.
	//
	// - 8: infant.
	//
	// example:
	//
	// 0
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s OrderDetailResponseBodyDataPassengerList) String() string {
	return dara.Prettify(s)
}

func (s OrderDetailResponseBodyDataPassengerList) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBodyDataPassengerList) GetBirthday() *string {
	return s.Birthday
}

func (s *OrderDetailResponseBodyDataPassengerList) GetCredential() *OrderDetailResponseBodyDataPassengerListCredential {
	return s.Credential
}

func (s *OrderDetailResponseBodyDataPassengerList) GetFirstName() *string {
	return s.FirstName
}

func (s *OrderDetailResponseBodyDataPassengerList) GetGender() *int32 {
	return s.Gender
}

func (s *OrderDetailResponseBodyDataPassengerList) GetLastName() *string {
	return s.LastName
}

func (s *OrderDetailResponseBodyDataPassengerList) GetMobileCountryCode() *string {
	return s.MobileCountryCode
}

func (s *OrderDetailResponseBodyDataPassengerList) GetMobilePhoneNumber() *string {
	return s.MobilePhoneNumber
}

func (s *OrderDetailResponseBodyDataPassengerList) GetNationality() *string {
	return s.Nationality
}

func (s *OrderDetailResponseBodyDataPassengerList) GetType() *int32 {
	return s.Type
}

func (s *OrderDetailResponseBodyDataPassengerList) SetBirthday(v string) *OrderDetailResponseBodyDataPassengerList {
	s.Birthday = &v
	return s
}

func (s *OrderDetailResponseBodyDataPassengerList) SetCredential(v *OrderDetailResponseBodyDataPassengerListCredential) *OrderDetailResponseBodyDataPassengerList {
	s.Credential = v
	return s
}

func (s *OrderDetailResponseBodyDataPassengerList) SetFirstName(v string) *OrderDetailResponseBodyDataPassengerList {
	s.FirstName = &v
	return s
}

func (s *OrderDetailResponseBodyDataPassengerList) SetGender(v int32) *OrderDetailResponseBodyDataPassengerList {
	s.Gender = &v
	return s
}

func (s *OrderDetailResponseBodyDataPassengerList) SetLastName(v string) *OrderDetailResponseBodyDataPassengerList {
	s.LastName = &v
	return s
}

func (s *OrderDetailResponseBodyDataPassengerList) SetMobileCountryCode(v string) *OrderDetailResponseBodyDataPassengerList {
	s.MobileCountryCode = &v
	return s
}

func (s *OrderDetailResponseBodyDataPassengerList) SetMobilePhoneNumber(v string) *OrderDetailResponseBodyDataPassengerList {
	s.MobilePhoneNumber = &v
	return s
}

func (s *OrderDetailResponseBodyDataPassengerList) SetNationality(v string) *OrderDetailResponseBodyDataPassengerList {
	s.Nationality = &v
	return s
}

func (s *OrderDetailResponseBodyDataPassengerList) SetType(v int32) *OrderDetailResponseBodyDataPassengerList {
	s.Type = &v
	return s
}

func (s *OrderDetailResponseBodyDataPassengerList) Validate() error {
	if s.Credential != nil {
		if err := s.Credential.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OrderDetailResponseBodyDataPassengerListCredential struct {
	// The place of issuance, represented as a two-letter code.
	//
	// example:
	//
	// CN
	CertIssuePlace *string `json:"cert_issue_place,omitempty" xml:"cert_issue_place,omitempty"`
	// The credential number.
	//
	// example:
	//
	// E1***5674
	CredentialNum *string `json:"credential_num,omitempty" xml:"credential_num,omitempty"`
	// The credential type. Valid values:
	//
	// - 0: ID card.
	//
	// - 1: passport.
	//
	// - 4: Home Return Permit.
	//
	// - 5: Mainland Travel Permit for Taiwan Residents.
	//
	// - 6: Exit-Entry Permit for Hong Kong and Macao Residents.
	//
	// - 12: Taiwan Travel Permit.
	//
	// - 19: no credential.
	//
	// example:
	//
	// 1
	CredentialType *int32 `json:"credential_type,omitempty" xml:"credential_type,omitempty"`
	// The credential expiration date.
	//
	// example:
	//
	// 20290101
	ExpireDate *string `json:"expire_date,omitempty" xml:"expire_date,omitempty"`
}

func (s OrderDetailResponseBodyDataPassengerListCredential) String() string {
	return dara.Prettify(s)
}

func (s OrderDetailResponseBodyDataPassengerListCredential) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBodyDataPassengerListCredential) GetCertIssuePlace() *string {
	return s.CertIssuePlace
}

func (s *OrderDetailResponseBodyDataPassengerListCredential) GetCredentialNum() *string {
	return s.CredentialNum
}

func (s *OrderDetailResponseBodyDataPassengerListCredential) GetCredentialType() *int32 {
	return s.CredentialType
}

func (s *OrderDetailResponseBodyDataPassengerListCredential) GetExpireDate() *string {
	return s.ExpireDate
}

func (s *OrderDetailResponseBodyDataPassengerListCredential) SetCertIssuePlace(v string) *OrderDetailResponseBodyDataPassengerListCredential {
	s.CertIssuePlace = &v
	return s
}

func (s *OrderDetailResponseBodyDataPassengerListCredential) SetCredentialNum(v string) *OrderDetailResponseBodyDataPassengerListCredential {
	s.CredentialNum = &v
	return s
}

func (s *OrderDetailResponseBodyDataPassengerListCredential) SetCredentialType(v int32) *OrderDetailResponseBodyDataPassengerListCredential {
	s.CredentialType = &v
	return s
}

func (s *OrderDetailResponseBodyDataPassengerListCredential) SetExpireDate(v string) *OrderDetailResponseBodyDataPassengerListCredential {
	s.ExpireDate = &v
	return s
}

func (s *OrderDetailResponseBodyDataPassengerListCredential) Validate() error {
	return dara.Validate(s)
}

type OrderDetailResponseBodyDataSolution struct {
	// The unit price for an adult.
	//
	// example:
	//
	// 300
	AdultPrice *float64 `json:"adult_price,omitempty" xml:"adult_price,omitempty"`
	// The tax for an adult.
	//
	// example:
	//
	// 30
	AdultTax *float64 `json:"adult_tax,omitempty" xml:"adult_tax,omitempty"`
	// The unit price for a child.
	//
	// example:
	//
	// 200
	ChildPrice *float64 `json:"child_price,omitempty" xml:"child_price,omitempty"`
	// The tax for a child.
	//
	// example:
	//
	// 20
	ChildTax *float64 `json:"child_tax,omitempty" xml:"child_tax,omitempty"`
	// The unit price for an infant.
	//
	// example:
	//
	// 100
	InfantPrice *float64 `json:"infant_price,omitempty" xml:"infant_price,omitempty"`
	// The tax for an infant.
	//
	// example:
	//
	// 10
	InfantTax *float64 `json:"infant_tax,omitempty" xml:"infant_tax,omitempty"`
	// The journey list.
	JourneyList []*OrderDetailResponseBodyDataSolutionJourneyList `json:"journey_list,omitempty" xml:"journey_list,omitempty" type:"Repeated"`
	// The baggage through-check rules.
	SegmentBaggageCheckInInfoList []*OrderDetailResponseBodyDataSolutionSegmentBaggageCheckInInfoList `json:"segment_baggage_check_in_info_list,omitempty" xml:"segment_baggage_check_in_info_list,omitempty" type:"Repeated"`
	// The complimentary baggage rules.
	SegmentBaggageMappingList []*OrderDetailResponseBodyDataSolutionSegmentBaggageMappingList `json:"segment_baggage_mapping_list,omitempty" xml:"segment_baggage_mapping_list,omitempty" type:"Repeated"`
	// The refund and change rules.
	SegmentRefundChangeRuleMappingList []*OrderDetailResponseBodyDataSolutionSegmentRefundChangeRuleMappingList `json:"segment_refund_change_rule_mapping_list,omitempty" xml:"segment_refund_change_rule_mapping_list,omitempty" type:"Repeated"`
	SolutionAttribute                  *OrderDetailResponseBodyDataSolutionSolutionAttribute                    `json:"solution_attribute,omitempty" xml:"solution_attribute,omitempty" type:"Struct"`
	// solution_id
	//
	// example:
	//
	// eJwz8DeySEo0NjQ01TU3TUxxx
	SolutionId *string `json:"solution_id,omitempty" xml:"solution_id,omitempty"`
}

func (s OrderDetailResponseBodyDataSolution) String() string {
	return dara.Prettify(s)
}

func (s OrderDetailResponseBodyDataSolution) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBodyDataSolution) GetAdultPrice() *float64 {
	return s.AdultPrice
}

func (s *OrderDetailResponseBodyDataSolution) GetAdultTax() *float64 {
	return s.AdultTax
}

func (s *OrderDetailResponseBodyDataSolution) GetChildPrice() *float64 {
	return s.ChildPrice
}

func (s *OrderDetailResponseBodyDataSolution) GetChildTax() *float64 {
	return s.ChildTax
}

func (s *OrderDetailResponseBodyDataSolution) GetInfantPrice() *float64 {
	return s.InfantPrice
}

func (s *OrderDetailResponseBodyDataSolution) GetInfantTax() *float64 {
	return s.InfantTax
}

func (s *OrderDetailResponseBodyDataSolution) GetJourneyList() []*OrderDetailResponseBodyDataSolutionJourneyList {
	return s.JourneyList
}

func (s *OrderDetailResponseBodyDataSolution) GetSegmentBaggageCheckInInfoList() []*OrderDetailResponseBodyDataSolutionSegmentBaggageCheckInInfoList {
	return s.SegmentBaggageCheckInInfoList
}

func (s *OrderDetailResponseBodyDataSolution) GetSegmentBaggageMappingList() []*OrderDetailResponseBodyDataSolutionSegmentBaggageMappingList {
	return s.SegmentBaggageMappingList
}

func (s *OrderDetailResponseBodyDataSolution) GetSegmentRefundChangeRuleMappingList() []*OrderDetailResponseBodyDataSolutionSegmentRefundChangeRuleMappingList {
	return s.SegmentRefundChangeRuleMappingList
}

func (s *OrderDetailResponseBodyDataSolution) GetSolutionAttribute() *OrderDetailResponseBodyDataSolutionSolutionAttribute {
	return s.SolutionAttribute
}

func (s *OrderDetailResponseBodyDataSolution) GetSolutionId() *string {
	return s.SolutionId
}

func (s *OrderDetailResponseBodyDataSolution) SetAdultPrice(v float64) *OrderDetailResponseBodyDataSolution {
	s.AdultPrice = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolution) SetAdultTax(v float64) *OrderDetailResponseBodyDataSolution {
	s.AdultTax = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolution) SetChildPrice(v float64) *OrderDetailResponseBodyDataSolution {
	s.ChildPrice = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolution) SetChildTax(v float64) *OrderDetailResponseBodyDataSolution {
	s.ChildTax = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolution) SetInfantPrice(v float64) *OrderDetailResponseBodyDataSolution {
	s.InfantPrice = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolution) SetInfantTax(v float64) *OrderDetailResponseBodyDataSolution {
	s.InfantTax = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolution) SetJourneyList(v []*OrderDetailResponseBodyDataSolutionJourneyList) *OrderDetailResponseBodyDataSolution {
	s.JourneyList = v
	return s
}

func (s *OrderDetailResponseBodyDataSolution) SetSegmentBaggageCheckInInfoList(v []*OrderDetailResponseBodyDataSolutionSegmentBaggageCheckInInfoList) *OrderDetailResponseBodyDataSolution {
	s.SegmentBaggageCheckInInfoList = v
	return s
}

func (s *OrderDetailResponseBodyDataSolution) SetSegmentBaggageMappingList(v []*OrderDetailResponseBodyDataSolutionSegmentBaggageMappingList) *OrderDetailResponseBodyDataSolution {
	s.SegmentBaggageMappingList = v
	return s
}

func (s *OrderDetailResponseBodyDataSolution) SetSegmentRefundChangeRuleMappingList(v []*OrderDetailResponseBodyDataSolutionSegmentRefundChangeRuleMappingList) *OrderDetailResponseBodyDataSolution {
	s.SegmentRefundChangeRuleMappingList = v
	return s
}

func (s *OrderDetailResponseBodyDataSolution) SetSolutionAttribute(v *OrderDetailResponseBodyDataSolutionSolutionAttribute) *OrderDetailResponseBodyDataSolution {
	s.SolutionAttribute = v
	return s
}

func (s *OrderDetailResponseBodyDataSolution) SetSolutionId(v string) *OrderDetailResponseBodyDataSolution {
	s.SolutionId = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolution) Validate() error {
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
	if s.SolutionAttribute != nil {
		if err := s.SolutionAttribute.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OrderDetailResponseBodyDataSolutionJourneyList struct {
	// The segment information.
	SegmentList []*OrderDetailResponseBodyDataSolutionJourneyListSegmentList `json:"segment_list,omitempty" xml:"segment_list,omitempty" type:"Repeated"`
	// The number of transfers.
	//
	// example:
	//
	// 0
	TransferCount *int32 `json:"transfer_count,omitempty" xml:"transfer_count,omitempty"`
}

func (s OrderDetailResponseBodyDataSolutionJourneyList) String() string {
	return dara.Prettify(s)
}

func (s OrderDetailResponseBodyDataSolutionJourneyList) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBodyDataSolutionJourneyList) GetSegmentList() []*OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	return s.SegmentList
}

func (s *OrderDetailResponseBodyDataSolutionJourneyList) GetTransferCount() *int32 {
	return s.TransferCount
}

func (s *OrderDetailResponseBodyDataSolutionJourneyList) SetSegmentList(v []*OrderDetailResponseBodyDataSolutionJourneyListSegmentList) *OrderDetailResponseBodyDataSolutionJourneyList {
	s.SegmentList = v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyList) SetTransferCount(v int32) *OrderDetailResponseBodyDataSolutionJourneyList {
	s.TransferCount = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyList) Validate() error {
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

type OrderDetailResponseBodyDataSolutionJourneyListSegmentList struct {
	// The three-letter IATA code of the arrival airport (uppercase).
	//
	// example:
	//
	// MFM
	ArrivalAirport *string `json:"arrival_airport,omitempty" xml:"arrival_airport,omitempty"`
	// The three-letter IATA code of the arrival city (uppercase).
	//
	// example:
	//
	// MFM
	ArrivalCity *string `json:"arrival_city,omitempty" xml:"arrival_city,omitempty"`
	// The arrival terminal.
	//
	// example:
	//
	// T1
	ArrivalTerminal *string `json:"arrival_terminal,omitempty" xml:"arrival_terminal,omitempty"`
	// The arrival date and time in string format (yyyy-MM-dd HH:mm:ss).
	//
	// example:
	//
	// 2023-03-10 10:40:00
	ArrivalTime *string `json:"arrival_time,omitempty" xml:"arrival_time,omitempty"`
	// The number of remaining seats.
	//
	// example:
	//
	// 7
	Availability *string `json:"availability,omitempty" xml:"availability,omitempty"`
	// The cabin.
	//
	// example:
	//
	// V
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// The cabin class.
	//
	// example:
	//
	// Y
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// Indicates whether the flight is a codeshare flight.
	//
	// example:
	//
	// false
	CodeShare *bool `json:"code_share,omitempty" xml:"code_share,omitempty"`
	// The three-letter IATA code of the departure airport (uppercase).
	//
	// example:
	//
	// PVG
	DepartureAirport *string `json:"departure_airport,omitempty" xml:"departure_airport,omitempty"`
	// The three-letter IATA code of the departure city (uppercase).
	//
	// example:
	//
	// SHA
	DepartureCity *string `json:"departure_city,omitempty" xml:"departure_city,omitempty"`
	// The departure terminal.
	//
	// example:
	//
	// T2
	DepartureTerminal *string `json:"departure_terminal,omitempty" xml:"departure_terminal,omitempty"`
	// The departure date and time in string format (yyyy-MM-dd HH:mm:ss).
	//
	// example:
	//
	// 2023-03-10 07:55:00
	DepartureTime *string `json:"departure_time,omitempty" xml:"departure_time,omitempty"`
	// The aircraft type.
	//
	// example:
	//
	// 32Q
	EquipType *string `json:"equip_type,omitempty" xml:"equip_type,omitempty"`
	// The flight duration. Unit: minutes.
	//
	// example:
	//
	// 165
	FlightDuration *int32 `json:"flight_duration,omitempty" xml:"flight_duration,omitempty"`
	// The marketing airline code (for example, HO).
	//
	// example:
	//
	// HO
	MarketingAirline *string `json:"marketing_airline,omitempty" xml:"marketing_airline,omitempty"`
	// The marketing flight number (for example, HO1295).
	//
	// example:
	//
	// HO1295
	MarketingFlightNo *string `json:"marketing_flight_no,omitempty" xml:"marketing_flight_no,omitempty"`
	// The numeric marketing flight number (for example, 1295).
	//
	// example:
	//
	// 1295
	MarketingFlightNoInt *int32 `json:"marketing_flight_no_int,omitempty" xml:"marketing_flight_no_int,omitempty"`
	// The operating airline code (for example, CX).
	//
	// example:
	//
	// HO
	OperatingAirline *string `json:"operating_airline,omitempty" xml:"operating_airline,omitempty"`
	// The operating flight number (for example, CX601).
	//
	// example:
	//
	// HO1295
	OperatingFlightNo *string `json:"operating_flight_no,omitempty" xml:"operating_flight_no,omitempty"`
	// The segment ID. Format: flight number + departure airport + arrival airport + departure date (MMdd).
	//
	// example:
	//
	// HO1295-PVG-MFM-20230310
	SegmentId *string `json:"segment_id,omitempty" xml:"segment_id,omitempty"`
	// The list of stopover cities. This parameter has a value when stopQuantity is greater than 0. Multiple cities are separated by commas.
	//
	// example:
	//
	// SEL,HKG
	StopCityList *string `json:"stop_city_list,omitempty" xml:"stop_city_list,omitempty"`
	// The number of stopover cities.
	//
	// example:
	//
	// 0
	StopQuantity *int32 `json:"stop_quantity,omitempty" xml:"stop_quantity,omitempty"`
}

func (s OrderDetailResponseBodyDataSolutionJourneyListSegmentList) String() string {
	return dara.Prettify(s)
}

func (s OrderDetailResponseBodyDataSolutionJourneyListSegmentList) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) GetArrivalAirport() *string {
	return s.ArrivalAirport
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) GetArrivalCity() *string {
	return s.ArrivalCity
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) GetArrivalTerminal() *string {
	return s.ArrivalTerminal
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) GetArrivalTime() *string {
	return s.ArrivalTime
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) GetAvailability() *string {
	return s.Availability
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) GetCabin() *string {
	return s.Cabin
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) GetCabinClass() *string {
	return s.CabinClass
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) GetCodeShare() *bool {
	return s.CodeShare
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) GetDepartureAirport() *string {
	return s.DepartureAirport
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) GetDepartureCity() *string {
	return s.DepartureCity
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) GetDepartureTerminal() *string {
	return s.DepartureTerminal
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) GetDepartureTime() *string {
	return s.DepartureTime
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) GetEquipType() *string {
	return s.EquipType
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) GetFlightDuration() *int32 {
	return s.FlightDuration
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) GetMarketingAirline() *string {
	return s.MarketingAirline
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) GetMarketingFlightNo() *string {
	return s.MarketingFlightNo
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) GetMarketingFlightNoInt() *int32 {
	return s.MarketingFlightNoInt
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) GetOperatingAirline() *string {
	return s.OperatingAirline
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) GetOperatingFlightNo() *string {
	return s.OperatingFlightNo
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) GetSegmentId() *string {
	return s.SegmentId
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) GetStopCityList() *string {
	return s.StopCityList
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) GetStopQuantity() *int32 {
	return s.StopQuantity
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetArrivalAirport(v string) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.ArrivalAirport = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetArrivalCity(v string) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.ArrivalCity = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetArrivalTerminal(v string) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.ArrivalTerminal = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetArrivalTime(v string) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.ArrivalTime = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetAvailability(v string) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.Availability = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetCabin(v string) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.Cabin = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetCabinClass(v string) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.CabinClass = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetCodeShare(v bool) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.CodeShare = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetDepartureAirport(v string) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.DepartureAirport = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetDepartureCity(v string) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.DepartureCity = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetDepartureTerminal(v string) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.DepartureTerminal = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetDepartureTime(v string) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.DepartureTime = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetEquipType(v string) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.EquipType = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetFlightDuration(v int32) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.FlightDuration = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetMarketingAirline(v string) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.MarketingAirline = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetMarketingFlightNo(v string) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.MarketingFlightNo = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetMarketingFlightNoInt(v int32) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.MarketingFlightNoInt = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetOperatingAirline(v string) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.OperatingAirline = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetOperatingFlightNo(v string) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.OperatingFlightNo = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetSegmentId(v string) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.SegmentId = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetStopCityList(v string) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.StopCityList = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) SetStopQuantity(v int32) *OrderDetailResponseBodyDataSolutionJourneyListSegmentList {
	s.StopQuantity = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionJourneyListSegmentList) Validate() error {
	return dara.Validate(s)
}

type OrderDetailResponseBodyDataSolutionSegmentBaggageCheckInInfoList struct {
	// The baggage through-check rule type. Valid values:
	//
	// - 1: baggage is checked through between segments.
	//
	// - 2: baggage must be rechecked between segments.
	//
	// - 3: baggage is checked through at stopover cities.
	//
	// - 4: baggage must be rechecked at stopover cities.
	//
	// example:
	//
	// 1
	LuggageDirectInfoType *int32 `json:"luggage_direct_info_type,omitempty" xml:"luggage_direct_info_type,omitempty"`
	// The list of segment IDs. These segments share the same baggage through-check rule.
	SegmentIdList []*string `json:"segment_id_list,omitempty" xml:"segment_id_list,omitempty" type:"Repeated"`
}

func (s OrderDetailResponseBodyDataSolutionSegmentBaggageCheckInInfoList) String() string {
	return dara.Prettify(s)
}

func (s OrderDetailResponseBodyDataSolutionSegmentBaggageCheckInInfoList) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBodyDataSolutionSegmentBaggageCheckInInfoList) GetLuggageDirectInfoType() *int32 {
	return s.LuggageDirectInfoType
}

func (s *OrderDetailResponseBodyDataSolutionSegmentBaggageCheckInInfoList) GetSegmentIdList() []*string {
	return s.SegmentIdList
}

func (s *OrderDetailResponseBodyDataSolutionSegmentBaggageCheckInInfoList) SetLuggageDirectInfoType(v int32) *OrderDetailResponseBodyDataSolutionSegmentBaggageCheckInInfoList {
	s.LuggageDirectInfoType = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionSegmentBaggageCheckInInfoList) SetSegmentIdList(v []*string) *OrderDetailResponseBodyDataSolutionSegmentBaggageCheckInInfoList {
	s.SegmentIdList = v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionSegmentBaggageCheckInInfoList) Validate() error {
	return dara.Validate(s)
}

type OrderDetailResponseBodyDataSolutionSegmentBaggageMappingList struct {
	// The mapping of passenger types to complimentary baggage allowances.
	PassengerBaggageAllowanceMapping map[string]*DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue `json:"passenger_baggage_allowance_mapping,omitempty" xml:"passenger_baggage_allowance_mapping,omitempty"`
	// The list of segment IDs. These segments share the same complimentary baggage rule.
	SegmentIdList []*string `json:"segment_id_list,omitempty" xml:"segment_id_list,omitempty" type:"Repeated"`
}

func (s OrderDetailResponseBodyDataSolutionSegmentBaggageMappingList) String() string {
	return dara.Prettify(s)
}

func (s OrderDetailResponseBodyDataSolutionSegmentBaggageMappingList) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBodyDataSolutionSegmentBaggageMappingList) GetPassengerBaggageAllowanceMapping() map[string]*DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue {
	return s.PassengerBaggageAllowanceMapping
}

func (s *OrderDetailResponseBodyDataSolutionSegmentBaggageMappingList) GetSegmentIdList() []*string {
	return s.SegmentIdList
}

func (s *OrderDetailResponseBodyDataSolutionSegmentBaggageMappingList) SetPassengerBaggageAllowanceMapping(v map[string]*DataSolutionSegmentBaggageMappingListPassengerBaggageAllowanceMappingValue) *OrderDetailResponseBodyDataSolutionSegmentBaggageMappingList {
	s.PassengerBaggageAllowanceMapping = v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionSegmentBaggageMappingList) SetSegmentIdList(v []*string) *OrderDetailResponseBodyDataSolutionSegmentBaggageMappingList {
	s.SegmentIdList = v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionSegmentBaggageMappingList) Validate() error {
	return dara.Validate(s)
}

type OrderDetailResponseBodyDataSolutionSegmentRefundChangeRuleMappingList struct {
	// The mapping of passenger types to refund and change rules.
	RefundChangeRuleMap map[string]*DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue `json:"refund_change_rule_map,omitempty" xml:"refund_change_rule_map,omitempty"`
	// The list of segment IDs. These segments share the same refund and change rule.
	SegmentIdList []*string `json:"segment_id_list,omitempty" xml:"segment_id_list,omitempty" type:"Repeated"`
}

func (s OrderDetailResponseBodyDataSolutionSegmentRefundChangeRuleMappingList) String() string {
	return dara.Prettify(s)
}

func (s OrderDetailResponseBodyDataSolutionSegmentRefundChangeRuleMappingList) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBodyDataSolutionSegmentRefundChangeRuleMappingList) GetRefundChangeRuleMap() map[string]*DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue {
	return s.RefundChangeRuleMap
}

func (s *OrderDetailResponseBodyDataSolutionSegmentRefundChangeRuleMappingList) GetSegmentIdList() []*string {
	return s.SegmentIdList
}

func (s *OrderDetailResponseBodyDataSolutionSegmentRefundChangeRuleMappingList) SetRefundChangeRuleMap(v map[string]*DataSolutionSegmentRefundChangeRuleMappingListRefundChangeRuleMapValue) *OrderDetailResponseBodyDataSolutionSegmentRefundChangeRuleMappingList {
	s.RefundChangeRuleMap = v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionSegmentRefundChangeRuleMappingList) SetSegmentIdList(v []*string) *OrderDetailResponseBodyDataSolutionSegmentRefundChangeRuleMappingList {
	s.SegmentIdList = v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionSegmentRefundChangeRuleMappingList) Validate() error {
	return dara.Validate(s)
}

type OrderDetailResponseBodyDataSolutionSolutionAttribute struct {
	IssueTimeInfo    *OrderDetailResponseBodyDataSolutionSolutionAttributeIssueTimeInfo `json:"issue_time_info,omitempty" xml:"issue_time_info,omitempty" type:"Struct"`
	SupplySourceType *string                                                            `json:"supply_source_type,omitempty" xml:"supply_source_type,omitempty"`
}

func (s OrderDetailResponseBodyDataSolutionSolutionAttribute) String() string {
	return dara.Prettify(s)
}

func (s OrderDetailResponseBodyDataSolutionSolutionAttribute) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBodyDataSolutionSolutionAttribute) GetIssueTimeInfo() *OrderDetailResponseBodyDataSolutionSolutionAttributeIssueTimeInfo {
	return s.IssueTimeInfo
}

func (s *OrderDetailResponseBodyDataSolutionSolutionAttribute) GetSupplySourceType() *string {
	return s.SupplySourceType
}

func (s *OrderDetailResponseBodyDataSolutionSolutionAttribute) SetIssueTimeInfo(v *OrderDetailResponseBodyDataSolutionSolutionAttributeIssueTimeInfo) *OrderDetailResponseBodyDataSolutionSolutionAttribute {
	s.IssueTimeInfo = v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionSolutionAttribute) SetSupplySourceType(v string) *OrderDetailResponseBodyDataSolutionSolutionAttribute {
	s.SupplySourceType = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionSolutionAttribute) Validate() error {
	if s.IssueTimeInfo != nil {
		if err := s.IssueTimeInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OrderDetailResponseBodyDataSolutionSolutionAttributeIssueTimeInfo struct {
	IssueTicketType *int32 `json:"issue_ticket_type,omitempty" xml:"issue_ticket_type,omitempty"`
	IssueTimeLimit  *int32 `json:"issue_time_limit,omitempty" xml:"issue_time_limit,omitempty"`
}

func (s OrderDetailResponseBodyDataSolutionSolutionAttributeIssueTimeInfo) String() string {
	return dara.Prettify(s)
}

func (s OrderDetailResponseBodyDataSolutionSolutionAttributeIssueTimeInfo) GoString() string {
	return s.String()
}

func (s *OrderDetailResponseBodyDataSolutionSolutionAttributeIssueTimeInfo) GetIssueTicketType() *int32 {
	return s.IssueTicketType
}

func (s *OrderDetailResponseBodyDataSolutionSolutionAttributeIssueTimeInfo) GetIssueTimeLimit() *int32 {
	return s.IssueTimeLimit
}

func (s *OrderDetailResponseBodyDataSolutionSolutionAttributeIssueTimeInfo) SetIssueTicketType(v int32) *OrderDetailResponseBodyDataSolutionSolutionAttributeIssueTimeInfo {
	s.IssueTicketType = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionSolutionAttributeIssueTimeInfo) SetIssueTimeLimit(v int32) *OrderDetailResponseBodyDataSolutionSolutionAttributeIssueTimeInfo {
	s.IssueTimeLimit = &v
	return s
}

func (s *OrderDetailResponseBodyDataSolutionSolutionAttributeIssueTimeInfo) Validate() error {
	return dara.Validate(s)
}
