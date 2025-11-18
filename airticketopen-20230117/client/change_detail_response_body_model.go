// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ChangeDetailResponseBody
	GetRequestId() *string
	SetData(v *ChangeDetailResponseBodyData) *ChangeDetailResponseBody
	GetData() *ChangeDetailResponseBodyData
	SetErrorCode(v string) *ChangeDetailResponseBody
	GetErrorCode() *string
	SetErrorData(v interface{}) *ChangeDetailResponseBody
	GetErrorData() interface{}
	SetErrorMsg(v string) *ChangeDetailResponseBody
	GetErrorMsg() *string
	SetStatus(v int32) *ChangeDetailResponseBody
	GetStatus() *int32
	SetSuccess(v bool) *ChangeDetailResponseBody
	GetSuccess() *bool
}

type ChangeDetailResponseBody struct {
	// Request RequestId
	//
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Correctly processed return data
	Data *ChangeDetailResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// error code
	//
	// example:
	//
	// null
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// Data carried in error handling
	//
	// example:
	//
	// null
	ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
	// Error message
	//
	// example:
	//
	// null
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// HTTP request successful, status value is 200
	//
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// Whether it is successful
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ChangeDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeDetailResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeDetailResponseBody) GetData() *ChangeDetailResponseBodyData {
	return s.Data
}

func (s *ChangeDetailResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ChangeDetailResponseBody) GetErrorData() interface{} {
	return s.ErrorData
}

func (s *ChangeDetailResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ChangeDetailResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *ChangeDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChangeDetailResponseBody) SetRequestId(v string) *ChangeDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeDetailResponseBody) SetData(v *ChangeDetailResponseBodyData) *ChangeDetailResponseBody {
	s.Data = v
	return s
}

func (s *ChangeDetailResponseBody) SetErrorCode(v string) *ChangeDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ChangeDetailResponseBody) SetErrorData(v interface{}) *ChangeDetailResponseBody {
	s.ErrorData = v
	return s
}

func (s *ChangeDetailResponseBody) SetErrorMsg(v string) *ChangeDetailResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ChangeDetailResponseBody) SetStatus(v int32) *ChangeDetailResponseBody {
	s.Status = &v
	return s
}

func (s *ChangeDetailResponseBody) SetSuccess(v bool) *ChangeDetailResponseBody {
	s.Success = &v
	return s
}

func (s *ChangeDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChangeDetailResponseBodyData struct {
	// Change fee details, per passenger
	ChangeFeeDetails []*ChangeDetailResponseBodyDataChangeFeeDetails `json:"change_fee_details,omitempty" xml:"change_fee_details,omitempty" type:"Repeated"`
	// Change order number
	//
	// example:
	//
	// 4988430***950
	ChangeOrderNum *int64 `json:"change_order_num,omitempty" xml:"change_order_num,omitempty"`
	// List of passengers for the change
	ChangePassengers []*ChangeDetailResponseBodyDataChangePassengers `json:"change_passengers,omitempty" xml:"change_passengers,omitempty" type:"Repeated"`
	// Change reason type.
	//
	// 0: Voluntary change;
	//
	// 1: Involuntary change, due to flight delay or cancellation, schedule changes, or other airline reasons;
	//
	// 2: Involuntary change, due to health reasons with a medical report
	//
	// example:
	//
	// 1
	ChangeReasonType *int32 `json:"change_reason_type,omitempty" xml:"change_reason_type,omitempty"`
	// New journeys
	ChangedJourneys []*ChangeDetailResponseBodyDataChangedJourneys `json:"changed_journeys,omitempty" xml:"changed_journeys,omitempty" type:"Repeated"`
	// Reason for closing the change order
	//
	// example:
	//
	// reason desc
	CloseReason *string `json:"close_reason,omitempty" xml:"close_reason,omitempty"`
	// The time when the order was closed, in UTC timestamp
	//
	// example:
	//
	// 1677415244000
	CloseUtcTime *int64 `json:"close_utc_time,omitempty" xml:"close_utc_time,omitempty"`
	// Contact information for the change request
	Contact *ChangeDetailResponseBodyDataContact `json:"contact,omitempty" xml:"contact,omitempty" type:"Struct"`
	// Creation time of the change order, UTC timestamp
	//
	// example:
	//
	// 1677415276000
	CreateUtcTime *int64 `json:"create_utc_time,omitempty" xml:"create_utc_time,omitempty"`
	// Latest payment time for the buyer, UTC timestamp
	//
	// example:
	//
	// 1677415278000
	LastConfirmUtcTime *int64 `json:"last_confirm_utc_time,omitempty" xml:"last_confirm_utc_time,omitempty"`
	// The itinerary of the last change
	LastJourneys []*ChangeDetailResponseBodyDataLastJourneys `json:"last_journeys,omitempty" xml:"last_journeys,omitempty" type:"Repeated"`
	// Ticketing Order number
	//
	// example:
	//
	// 5988430***541
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
	// Change order status 0: Initial state; 1: Pending payment; 2: Payment successful; 3: Change successful; 4: Change closed
	//
	// example:
	//
	// 2
	OrderStatus *int32 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// Original journeys
	OriginalJourneys []*ChangeDetailResponseBodyDataOriginalJourneys `json:"original_journeys,omitempty" xml:"original_journeys,omitempty" type:"Repeated"`
	// Payment status 0: initial state; 1: pending payment; 2: payment successful; 4: successfully closed paid order; 5: successfully closed unpaid order
	//
	// example:
	//
	// 2
	PayStatus *int32 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// The time when the buyer successfully paid, in UTC timestamp
	//
	// example:
	//
	// 1677415255000
	PaySuccessUtcTime *int64 `json:"pay_success_utc_time,omitempty" xml:"pay_success_utc_time,omitempty"`
	// Total payment amount for the change order
	//
	// example:
	//
	// 300
	TotalAmount *float64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// Transaction serial number
	//
	// example:
	//
	// hkduendkd-2023-dj0
	TransactionNo *string `json:"transaction_no,omitempty" xml:"transaction_no,omitempty"`
}

func (s ChangeDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ChangeDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *ChangeDetailResponseBodyData) GetChangeFeeDetails() []*ChangeDetailResponseBodyDataChangeFeeDetails {
	return s.ChangeFeeDetails
}

func (s *ChangeDetailResponseBodyData) GetChangeOrderNum() *int64 {
	return s.ChangeOrderNum
}

func (s *ChangeDetailResponseBodyData) GetChangePassengers() []*ChangeDetailResponseBodyDataChangePassengers {
	return s.ChangePassengers
}

func (s *ChangeDetailResponseBodyData) GetChangeReasonType() *int32 {
	return s.ChangeReasonType
}

func (s *ChangeDetailResponseBodyData) GetChangedJourneys() []*ChangeDetailResponseBodyDataChangedJourneys {
	return s.ChangedJourneys
}

func (s *ChangeDetailResponseBodyData) GetCloseReason() *string {
	return s.CloseReason
}

func (s *ChangeDetailResponseBodyData) GetCloseUtcTime() *int64 {
	return s.CloseUtcTime
}

func (s *ChangeDetailResponseBodyData) GetContact() *ChangeDetailResponseBodyDataContact {
	return s.Contact
}

func (s *ChangeDetailResponseBodyData) GetCreateUtcTime() *int64 {
	return s.CreateUtcTime
}

func (s *ChangeDetailResponseBodyData) GetLastConfirmUtcTime() *int64 {
	return s.LastConfirmUtcTime
}

func (s *ChangeDetailResponseBodyData) GetLastJourneys() []*ChangeDetailResponseBodyDataLastJourneys {
	return s.LastJourneys
}

func (s *ChangeDetailResponseBodyData) GetOrderNum() *int64 {
	return s.OrderNum
}

func (s *ChangeDetailResponseBodyData) GetOrderStatus() *int32 {
	return s.OrderStatus
}

func (s *ChangeDetailResponseBodyData) GetOriginalJourneys() []*ChangeDetailResponseBodyDataOriginalJourneys {
	return s.OriginalJourneys
}

func (s *ChangeDetailResponseBodyData) GetPayStatus() *int32 {
	return s.PayStatus
}

func (s *ChangeDetailResponseBodyData) GetPaySuccessUtcTime() *int64 {
	return s.PaySuccessUtcTime
}

func (s *ChangeDetailResponseBodyData) GetTotalAmount() *float64 {
	return s.TotalAmount
}

func (s *ChangeDetailResponseBodyData) GetTransactionNo() *string {
	return s.TransactionNo
}

func (s *ChangeDetailResponseBodyData) SetChangeFeeDetails(v []*ChangeDetailResponseBodyDataChangeFeeDetails) *ChangeDetailResponseBodyData {
	s.ChangeFeeDetails = v
	return s
}

func (s *ChangeDetailResponseBodyData) SetChangeOrderNum(v int64) *ChangeDetailResponseBodyData {
	s.ChangeOrderNum = &v
	return s
}

func (s *ChangeDetailResponseBodyData) SetChangePassengers(v []*ChangeDetailResponseBodyDataChangePassengers) *ChangeDetailResponseBodyData {
	s.ChangePassengers = v
	return s
}

func (s *ChangeDetailResponseBodyData) SetChangeReasonType(v int32) *ChangeDetailResponseBodyData {
	s.ChangeReasonType = &v
	return s
}

func (s *ChangeDetailResponseBodyData) SetChangedJourneys(v []*ChangeDetailResponseBodyDataChangedJourneys) *ChangeDetailResponseBodyData {
	s.ChangedJourneys = v
	return s
}

func (s *ChangeDetailResponseBodyData) SetCloseReason(v string) *ChangeDetailResponseBodyData {
	s.CloseReason = &v
	return s
}

func (s *ChangeDetailResponseBodyData) SetCloseUtcTime(v int64) *ChangeDetailResponseBodyData {
	s.CloseUtcTime = &v
	return s
}

func (s *ChangeDetailResponseBodyData) SetContact(v *ChangeDetailResponseBodyDataContact) *ChangeDetailResponseBodyData {
	s.Contact = v
	return s
}

func (s *ChangeDetailResponseBodyData) SetCreateUtcTime(v int64) *ChangeDetailResponseBodyData {
	s.CreateUtcTime = &v
	return s
}

func (s *ChangeDetailResponseBodyData) SetLastConfirmUtcTime(v int64) *ChangeDetailResponseBodyData {
	s.LastConfirmUtcTime = &v
	return s
}

func (s *ChangeDetailResponseBodyData) SetLastJourneys(v []*ChangeDetailResponseBodyDataLastJourneys) *ChangeDetailResponseBodyData {
	s.LastJourneys = v
	return s
}

func (s *ChangeDetailResponseBodyData) SetOrderNum(v int64) *ChangeDetailResponseBodyData {
	s.OrderNum = &v
	return s
}

func (s *ChangeDetailResponseBodyData) SetOrderStatus(v int32) *ChangeDetailResponseBodyData {
	s.OrderStatus = &v
	return s
}

func (s *ChangeDetailResponseBodyData) SetOriginalJourneys(v []*ChangeDetailResponseBodyDataOriginalJourneys) *ChangeDetailResponseBodyData {
	s.OriginalJourneys = v
	return s
}

func (s *ChangeDetailResponseBodyData) SetPayStatus(v int32) *ChangeDetailResponseBodyData {
	s.PayStatus = &v
	return s
}

func (s *ChangeDetailResponseBodyData) SetPaySuccessUtcTime(v int64) *ChangeDetailResponseBodyData {
	s.PaySuccessUtcTime = &v
	return s
}

func (s *ChangeDetailResponseBodyData) SetTotalAmount(v float64) *ChangeDetailResponseBodyData {
	s.TotalAmount = &v
	return s
}

func (s *ChangeDetailResponseBodyData) SetTransactionNo(v string) *ChangeDetailResponseBodyData {
	s.TransactionNo = &v
	return s
}

func (s *ChangeDetailResponseBodyData) Validate() error {
	if s.ChangeFeeDetails != nil {
		for _, item := range s.ChangeFeeDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ChangePassengers != nil {
		for _, item := range s.ChangePassengers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ChangedJourneys != nil {
		for _, item := range s.ChangedJourneys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Contact != nil {
		if err := s.Contact.Validate(); err != nil {
			return err
		}
	}
	if s.LastJourneys != nil {
		for _, item := range s.LastJourneys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OriginalJourneys != nil {
		for _, item := range s.OriginalJourneys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ChangeDetailResponseBodyDataChangeFeeDetails struct {
	// Change fee details for the passenger
	ChangeFee *ChangeDetailResponseBodyDataChangeFeeDetailsChangeFee `json:"change_fee,omitempty" xml:"change_fee,omitempty" type:"Struct"`
	// Information of the passenger for the change
	Passenger *ChangeDetailResponseBodyDataChangeFeeDetailsPassenger `json:"passenger,omitempty" xml:"passenger,omitempty" type:"Struct"`
}

func (s ChangeDetailResponseBodyDataChangeFeeDetails) String() string {
	return dara.Prettify(s)
}

func (s ChangeDetailResponseBodyDataChangeFeeDetails) GoString() string {
	return s.String()
}

func (s *ChangeDetailResponseBodyDataChangeFeeDetails) GetChangeFee() *ChangeDetailResponseBodyDataChangeFeeDetailsChangeFee {
	return s.ChangeFee
}

func (s *ChangeDetailResponseBodyDataChangeFeeDetails) GetPassenger() *ChangeDetailResponseBodyDataChangeFeeDetailsPassenger {
	return s.Passenger
}

func (s *ChangeDetailResponseBodyDataChangeFeeDetails) SetChangeFee(v *ChangeDetailResponseBodyDataChangeFeeDetailsChangeFee) *ChangeDetailResponseBodyDataChangeFeeDetails {
	s.ChangeFee = v
	return s
}

func (s *ChangeDetailResponseBodyDataChangeFeeDetails) SetPassenger(v *ChangeDetailResponseBodyDataChangeFeeDetailsPassenger) *ChangeDetailResponseBodyDataChangeFeeDetails {
	s.Passenger = v
	return s
}

func (s *ChangeDetailResponseBodyDataChangeFeeDetails) Validate() error {
	if s.ChangeFee != nil {
		if err := s.ChangeFee.Validate(); err != nil {
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

type ChangeDetailResponseBodyDataChangeFeeDetailsChangeFee struct {
	// fare penalty
	//
	// example:
	//
	// 50
	ServiceFee *float64 `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
	// tax penalty
	//
	// example:
	//
	// 20
	TaxFee *float64 `json:"tax_fee,omitempty" xml:"tax_fee,omitempty"`
	// price difference
	//
	// example:
	//
	// 30
	UpgradeFee *float64 `json:"upgrade_fee,omitempty" xml:"upgrade_fee,omitempty"`
}

func (s ChangeDetailResponseBodyDataChangeFeeDetailsChangeFee) String() string {
	return dara.Prettify(s)
}

func (s ChangeDetailResponseBodyDataChangeFeeDetailsChangeFee) GoString() string {
	return s.String()
}

func (s *ChangeDetailResponseBodyDataChangeFeeDetailsChangeFee) GetServiceFee() *float64 {
	return s.ServiceFee
}

func (s *ChangeDetailResponseBodyDataChangeFeeDetailsChangeFee) GetTaxFee() *float64 {
	return s.TaxFee
}

func (s *ChangeDetailResponseBodyDataChangeFeeDetailsChangeFee) GetUpgradeFee() *float64 {
	return s.UpgradeFee
}

func (s *ChangeDetailResponseBodyDataChangeFeeDetailsChangeFee) SetServiceFee(v float64) *ChangeDetailResponseBodyDataChangeFeeDetailsChangeFee {
	s.ServiceFee = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangeFeeDetailsChangeFee) SetTaxFee(v float64) *ChangeDetailResponseBodyDataChangeFeeDetailsChangeFee {
	s.TaxFee = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangeFeeDetailsChangeFee) SetUpgradeFee(v float64) *ChangeDetailResponseBodyDataChangeFeeDetailsChangeFee {
	s.UpgradeFee = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangeFeeDetailsChangeFee) Validate() error {
	return dara.Validate(s)
}

type ChangeDetailResponseBodyDataChangeFeeDetailsPassenger struct {
	// Document number
	//
	// example:
	//
	// 411***********4411
	Document *string `json:"document,omitempty" xml:"document,omitempty"`
	// Passenger\\"s first name
	//
	// example:
	//
	// SAN
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// Passenger\\"s last name
	//
	// example:
	//
	// ZHANG
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
}

func (s ChangeDetailResponseBodyDataChangeFeeDetailsPassenger) String() string {
	return dara.Prettify(s)
}

func (s ChangeDetailResponseBodyDataChangeFeeDetailsPassenger) GoString() string {
	return s.String()
}

func (s *ChangeDetailResponseBodyDataChangeFeeDetailsPassenger) GetDocument() *string {
	return s.Document
}

func (s *ChangeDetailResponseBodyDataChangeFeeDetailsPassenger) GetFirstName() *string {
	return s.FirstName
}

func (s *ChangeDetailResponseBodyDataChangeFeeDetailsPassenger) GetLastName() *string {
	return s.LastName
}

func (s *ChangeDetailResponseBodyDataChangeFeeDetailsPassenger) SetDocument(v string) *ChangeDetailResponseBodyDataChangeFeeDetailsPassenger {
	s.Document = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangeFeeDetailsPassenger) SetFirstName(v string) *ChangeDetailResponseBodyDataChangeFeeDetailsPassenger {
	s.FirstName = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangeFeeDetailsPassenger) SetLastName(v string) *ChangeDetailResponseBodyDataChangeFeeDetailsPassenger {
	s.LastName = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangeFeeDetailsPassenger) Validate() error {
	return dara.Validate(s)
}

type ChangeDetailResponseBodyDataChangePassengers struct {
	// Document number
	//
	// example:
	//
	// 411***********4411
	Document *string `json:"document,omitempty" xml:"document,omitempty"`
	// Passenger first name
	//
	// example:
	//
	// SAN
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// Passenger last name
	//
	// example:
	//
	// ZHANG
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
}

func (s ChangeDetailResponseBodyDataChangePassengers) String() string {
	return dara.Prettify(s)
}

func (s ChangeDetailResponseBodyDataChangePassengers) GoString() string {
	return s.String()
}

func (s *ChangeDetailResponseBodyDataChangePassengers) GetDocument() *string {
	return s.Document
}

func (s *ChangeDetailResponseBodyDataChangePassengers) GetFirstName() *string {
	return s.FirstName
}

func (s *ChangeDetailResponseBodyDataChangePassengers) GetLastName() *string {
	return s.LastName
}

func (s *ChangeDetailResponseBodyDataChangePassengers) SetDocument(v string) *ChangeDetailResponseBodyDataChangePassengers {
	s.Document = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangePassengers) SetFirstName(v string) *ChangeDetailResponseBodyDataChangePassengers {
	s.FirstName = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangePassengers) SetLastName(v string) *ChangeDetailResponseBodyDataChangePassengers {
	s.LastName = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangePassengers) Validate() error {
	return dara.Validate(s)
}

type ChangeDetailResponseBodyDataChangedJourneys struct {
	// Segment information
	SegmentList []*ChangeDetailResponseBodyDataChangedJourneysSegmentList `json:"segment_list,omitempty" xml:"segment_list,omitempty" type:"Repeated"`
	// Number of transfers
	//
	// example:
	//
	// 0
	TransferCount *int32 `json:"transfer_count,omitempty" xml:"transfer_count,omitempty"`
}

func (s ChangeDetailResponseBodyDataChangedJourneys) String() string {
	return dara.Prettify(s)
}

func (s ChangeDetailResponseBodyDataChangedJourneys) GoString() string {
	return s.String()
}

func (s *ChangeDetailResponseBodyDataChangedJourneys) GetSegmentList() []*ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	return s.SegmentList
}

func (s *ChangeDetailResponseBodyDataChangedJourneys) GetTransferCount() *int32 {
	return s.TransferCount
}

func (s *ChangeDetailResponseBodyDataChangedJourneys) SetSegmentList(v []*ChangeDetailResponseBodyDataChangedJourneysSegmentList) *ChangeDetailResponseBodyDataChangedJourneys {
	s.SegmentList = v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneys) SetTransferCount(v int32) *ChangeDetailResponseBodyDataChangedJourneys {
	s.TransferCount = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneys) Validate() error {
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

type ChangeDetailResponseBodyDataChangedJourneysSegmentList struct {
	// Arrival airport three-letter code (uppercase)
	//
	// example:
	//
	// MFM
	ArrivalAirport *string `json:"arrival_airport,omitempty" xml:"arrival_airport,omitempty"`
	// Arrival city three-letter code (uppercase)
	//
	// example:
	//
	// MFM
	ArrivalCity *string `json:"arrival_city,omitempty" xml:"arrival_city,omitempty"`
	// Arrival terminal of the flight
	//
	// example:
	//
	// T1
	ArrivalTerminal *string `json:"arrival_terminal,omitempty" xml:"arrival_terminal,omitempty"`
	// Flight arrival date and time, in the format (yyyy-MM-dd HH:mm:ss)
	//
	// example:
	//
	// 2023-03-10 10:40:00
	ArrivalTime *string `json:"arrival_time,omitempty" xml:"arrival_time,omitempty"`
	// Number of available seats
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
	// service class ( compartment )
	//
	// example:
	//
	// Y
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// Whether it is a code-share flight
	//
	// example:
	//
	// false
	CodeShare *bool `json:"code_share,omitempty" xml:"code_share,omitempty"`
	// Departure airport three-letter code (uppercase)
	//
	// example:
	//
	// PVG
	DepartureAirport *string `json:"departure_airport,omitempty" xml:"departure_airport,omitempty"`
	// Departure city three-letter code (uppercase)
	//
	// example:
	//
	// SHA
	DepartureCity *string `json:"departure_city,omitempty" xml:"departure_city,omitempty"`
	// Departure terminal of the flight
	//
	// example:
	//
	// T2
	DepartureTerminal *string `json:"departure_terminal,omitempty" xml:"departure_terminal,omitempty"`
	// Flight departure date and time, in the format (yyyy-MM-dd HH:mm:ss)
	//
	// example:
	//
	// 2023-03-10 07:55:00
	DepartureTime *string `json:"departure_time,omitempty" xml:"departure_time,omitempty"`
	// Aircraft type
	//
	// example:
	//
	// 32Q
	EquipType *string `json:"equip_type,omitempty" xml:"equip_type,omitempty"`
	// Flight duration in minutes
	//
	// example:
	//
	// 165
	FlightDuration *int32 `json:"flight_duration,omitempty" xml:"flight_duration,omitempty"`
	// Market airline (e.g., HO)
	//
	// example:
	//
	// HO
	MarketingAirline *string `json:"marketing_airline,omitempty" xml:"marketing_airline,omitempty"`
	// Marketing flight number (e.g., HO1295)
	//
	// example:
	//
	// HO1295
	MarketingFlightNo *string `json:"marketing_flight_no,omitempty" xml:"marketing_flight_no,omitempty"`
	// Marketing flight number (e.g., 1295)
	//
	// example:
	//
	// 1295
	MarketingFlightNoInt *int32 `json:"marketing_flight_no_int,omitempty" xml:"marketing_flight_no_int,omitempty"`
	// Operating airline (e.g., CX)
	//
	// example:
	//
	// HO
	OperatingAirline *string `json:"operating_airline,omitempty" xml:"operating_airline,omitempty"`
	// Operating flight number (e.g., CX601)
	//
	// example:
	//
	// HO1295
	OperatingFlightNo *string `json:"operating_flight_no,omitempty" xml:"operating_flight_no,omitempty"`
	// Segment ID format: flight number + departure airport + arrival airport + departure date (MMdd)
	//
	// example:
	//
	// HO1295-PVG-MFM-20230310
	SegmentId *string `json:"segment_id,omitempty" xml:"segment_id,omitempty"`
	// List of stop cities, with values when stopQuantity > 0, separated by commas
	//
	// example:
	//
	// SEL,HKG
	StopCityList *string `json:"stop_city_list,omitempty" xml:"stop_city_list,omitempty"`
	// Number of stop cities
	//
	// example:
	//
	// 0
	StopQuantity *int32 `json:"stop_quantity,omitempty" xml:"stop_quantity,omitempty"`
}

func (s ChangeDetailResponseBodyDataChangedJourneysSegmentList) String() string {
	return dara.Prettify(s)
}

func (s ChangeDetailResponseBodyDataChangedJourneysSegmentList) GoString() string {
	return s.String()
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) GetArrivalAirport() *string {
	return s.ArrivalAirport
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) GetArrivalCity() *string {
	return s.ArrivalCity
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) GetArrivalTerminal() *string {
	return s.ArrivalTerminal
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) GetArrivalTime() *string {
	return s.ArrivalTime
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) GetAvailability() *string {
	return s.Availability
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) GetCabin() *string {
	return s.Cabin
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) GetCabinClass() *string {
	return s.CabinClass
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) GetCodeShare() *bool {
	return s.CodeShare
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) GetDepartureAirport() *string {
	return s.DepartureAirport
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) GetDepartureCity() *string {
	return s.DepartureCity
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) GetDepartureTerminal() *string {
	return s.DepartureTerminal
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) GetDepartureTime() *string {
	return s.DepartureTime
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) GetEquipType() *string {
	return s.EquipType
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) GetFlightDuration() *int32 {
	return s.FlightDuration
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) GetMarketingAirline() *string {
	return s.MarketingAirline
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) GetMarketingFlightNo() *string {
	return s.MarketingFlightNo
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) GetMarketingFlightNoInt() *int32 {
	return s.MarketingFlightNoInt
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) GetOperatingAirline() *string {
	return s.OperatingAirline
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) GetOperatingFlightNo() *string {
	return s.OperatingFlightNo
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) GetSegmentId() *string {
	return s.SegmentId
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) GetStopCityList() *string {
	return s.StopCityList
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) GetStopQuantity() *int32 {
	return s.StopQuantity
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetArrivalAirport(v string) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.ArrivalAirport = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetArrivalCity(v string) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.ArrivalCity = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetArrivalTerminal(v string) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.ArrivalTerminal = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetArrivalTime(v string) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.ArrivalTime = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetAvailability(v string) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.Availability = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetCabin(v string) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.Cabin = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetCabinClass(v string) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.CabinClass = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetCodeShare(v bool) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.CodeShare = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetDepartureAirport(v string) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.DepartureAirport = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetDepartureCity(v string) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.DepartureCity = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetDepartureTerminal(v string) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.DepartureTerminal = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetDepartureTime(v string) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.DepartureTime = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetEquipType(v string) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.EquipType = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetFlightDuration(v int32) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.FlightDuration = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetMarketingAirline(v string) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.MarketingAirline = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetMarketingFlightNo(v string) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.MarketingFlightNo = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetMarketingFlightNoInt(v int32) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.MarketingFlightNoInt = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetOperatingAirline(v string) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.OperatingAirline = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetOperatingFlightNo(v string) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.OperatingFlightNo = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetSegmentId(v string) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.SegmentId = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetStopCityList(v string) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.StopCityList = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) SetStopQuantity(v int32) *ChangeDetailResponseBodyDataChangedJourneysSegmentList {
	s.StopQuantity = &v
	return s
}

func (s *ChangeDetailResponseBodyDataChangedJourneysSegmentList) Validate() error {
	return dara.Validate(s)
}

type ChangeDetailResponseBodyDataContact struct {
	// Email address
	//
	// example:
	//
	// gao******@gmail.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// Country code
	//
	// example:
	//
	// 86
	MobileCountryCode *string `json:"mobile_country_code,omitempty" xml:"mobile_country_code,omitempty"`
	// Contact\\"s mobile phone number
	//
	// example:
	//
	// 183*****92
	MobilePhoneNum *string `json:"mobile_phone_num,omitempty" xml:"mobile_phone_num,omitempty"`
}

func (s ChangeDetailResponseBodyDataContact) String() string {
	return dara.Prettify(s)
}

func (s ChangeDetailResponseBodyDataContact) GoString() string {
	return s.String()
}

func (s *ChangeDetailResponseBodyDataContact) GetEmail() *string {
	return s.Email
}

func (s *ChangeDetailResponseBodyDataContact) GetMobileCountryCode() *string {
	return s.MobileCountryCode
}

func (s *ChangeDetailResponseBodyDataContact) GetMobilePhoneNum() *string {
	return s.MobilePhoneNum
}

func (s *ChangeDetailResponseBodyDataContact) SetEmail(v string) *ChangeDetailResponseBodyDataContact {
	s.Email = &v
	return s
}

func (s *ChangeDetailResponseBodyDataContact) SetMobileCountryCode(v string) *ChangeDetailResponseBodyDataContact {
	s.MobileCountryCode = &v
	return s
}

func (s *ChangeDetailResponseBodyDataContact) SetMobilePhoneNum(v string) *ChangeDetailResponseBodyDataContact {
	s.MobilePhoneNum = &v
	return s
}

func (s *ChangeDetailResponseBodyDataContact) Validate() error {
	return dara.Validate(s)
}

type ChangeDetailResponseBodyDataLastJourneys struct {
	// Segment information
	SegmentList []*ChangeDetailResponseBodyDataLastJourneysSegmentList `json:"segment_list,omitempty" xml:"segment_list,omitempty" type:"Repeated"`
	// Number of transfers
	//
	// example:
	//
	// 0
	TransferCount *int32 `json:"transfer_count,omitempty" xml:"transfer_count,omitempty"`
}

func (s ChangeDetailResponseBodyDataLastJourneys) String() string {
	return dara.Prettify(s)
}

func (s ChangeDetailResponseBodyDataLastJourneys) GoString() string {
	return s.String()
}

func (s *ChangeDetailResponseBodyDataLastJourneys) GetSegmentList() []*ChangeDetailResponseBodyDataLastJourneysSegmentList {
	return s.SegmentList
}

func (s *ChangeDetailResponseBodyDataLastJourneys) GetTransferCount() *int32 {
	return s.TransferCount
}

func (s *ChangeDetailResponseBodyDataLastJourneys) SetSegmentList(v []*ChangeDetailResponseBodyDataLastJourneysSegmentList) *ChangeDetailResponseBodyDataLastJourneys {
	s.SegmentList = v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneys) SetTransferCount(v int32) *ChangeDetailResponseBodyDataLastJourneys {
	s.TransferCount = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneys) Validate() error {
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

type ChangeDetailResponseBodyDataLastJourneysSegmentList struct {
	// Arrival airport three-letter code (uppercase)
	//
	// example:
	//
	// MFM
	ArrivalAirport *string `json:"arrival_airport,omitempty" xml:"arrival_airport,omitempty"`
	// Arrival city three-letter code (uppercase)
	//
	// example:
	//
	// MFM
	ArrivalCity *string `json:"arrival_city,omitempty" xml:"arrival_city,omitempty"`
	// Arrival terminal of the flight
	//
	// example:
	//
	// T1
	ArrivalTerminal *string `json:"arrival_terminal,omitempty" xml:"arrival_terminal,omitempty"`
	// Flight arrival date and time, in the format (yyyy-MM-dd HH:mm:ss)
	//
	// example:
	//
	// 2023-03-10 10:40:00
	ArrivalTime *string `json:"arrival_time,omitempty" xml:"arrival_time,omitempty"`
	// Number of available seats
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
	// service class ( compartment )
	//
	// example:
	//
	// Y
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// Whether it is a codeshare flight
	//
	// example:
	//
	// false
	CodeShare *bool `json:"code_share,omitempty" xml:"code_share,omitempty"`
	// Departure airport three-letter code (uppercase)
	//
	// example:
	//
	// PVG
	DepartureAirport *string `json:"departure_airport,omitempty" xml:"departure_airport,omitempty"`
	// Departure city three-letter code (uppercase)
	//
	// example:
	//
	// SHA
	DepartureCity *string `json:"departure_city,omitempty" xml:"departure_city,omitempty"`
	// Departure terminal of the flight
	//
	// example:
	//
	// T2
	DepartureTerminal *string `json:"departure_terminal,omitempty" xml:"departure_terminal,omitempty"`
	// Flight departure date and time, in the format (yyyy-MM-dd HH:mm:ss)
	//
	// example:
	//
	// 2023-03-10 07:55:00
	DepartureTime *string `json:"departure_time,omitempty" xml:"departure_time,omitempty"`
	// Aircraft type
	//
	// example:
	//
	// 32Q
	EquipType *string `json:"equip_type,omitempty" xml:"equip_type,omitempty"`
	// Flight duration in minutes
	//
	// example:
	//
	// 165
	FlightDuration *int32 `json:"flight_duration,omitempty" xml:"flight_duration,omitempty"`
	// Marketing airline (e.g., HO)
	//
	// example:
	//
	// HO
	MarketingAirline *string `json:"marketing_airline,omitempty" xml:"marketing_airline,omitempty"`
	// Marketing flight number (e.g., HO1295)
	//
	// example:
	//
	// HO1295
	MarketingFlightNo *string `json:"marketing_flight_no,omitempty" xml:"marketing_flight_no,omitempty"`
	// Marketing flight number (e.g., 1295)
	//
	// example:
	//
	// 1295
	MarketingFlightNoInt *int32 `json:"marketing_flight_no_int,omitempty" xml:"marketing_flight_no_int,omitempty"`
	// Operating airline (e.g., CX)
	//
	// example:
	//
	// HO
	OperatingAirline *string `json:"operating_airline,omitempty" xml:"operating_airline,omitempty"`
	// Operating flight number (e.g., CX601)
	//
	// example:
	//
	// HO1295
	OperatingFlightNo *string `json:"operating_flight_no,omitempty" xml:"operating_flight_no,omitempty"`
	// Segment ID format: flight number + departure airport + arrival airport + departure date (MMdd)
	//
	// example:
	//
	// HO1295-PVG-MFM-20230310
	SegmentId *string `json:"segment_id,omitempty" xml:"segment_id,omitempty"`
	// List of stop cities, with values when stopQuantity > 0, separated by commas
	//
	// example:
	//
	// SEL,HKG
	StopCityList *string `json:"stop_city_list,omitempty" xml:"stop_city_list,omitempty"`
	// Number of stop cities
	//
	// example:
	//
	// 0
	StopQuantity *int32 `json:"stop_quantity,omitempty" xml:"stop_quantity,omitempty"`
}

func (s ChangeDetailResponseBodyDataLastJourneysSegmentList) String() string {
	return dara.Prettify(s)
}

func (s ChangeDetailResponseBodyDataLastJourneysSegmentList) GoString() string {
	return s.String()
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) GetArrivalAirport() *string {
	return s.ArrivalAirport
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) GetArrivalCity() *string {
	return s.ArrivalCity
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) GetArrivalTerminal() *string {
	return s.ArrivalTerminal
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) GetArrivalTime() *string {
	return s.ArrivalTime
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) GetAvailability() *string {
	return s.Availability
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) GetCabin() *string {
	return s.Cabin
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) GetCabinClass() *string {
	return s.CabinClass
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) GetCodeShare() *bool {
	return s.CodeShare
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) GetDepartureAirport() *string {
	return s.DepartureAirport
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) GetDepartureCity() *string {
	return s.DepartureCity
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) GetDepartureTerminal() *string {
	return s.DepartureTerminal
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) GetDepartureTime() *string {
	return s.DepartureTime
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) GetEquipType() *string {
	return s.EquipType
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) GetFlightDuration() *int32 {
	return s.FlightDuration
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) GetMarketingAirline() *string {
	return s.MarketingAirline
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) GetMarketingFlightNo() *string {
	return s.MarketingFlightNo
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) GetMarketingFlightNoInt() *int32 {
	return s.MarketingFlightNoInt
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) GetOperatingAirline() *string {
	return s.OperatingAirline
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) GetOperatingFlightNo() *string {
	return s.OperatingFlightNo
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) GetSegmentId() *string {
	return s.SegmentId
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) GetStopCityList() *string {
	return s.StopCityList
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) GetStopQuantity() *int32 {
	return s.StopQuantity
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetArrivalAirport(v string) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.ArrivalAirport = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetArrivalCity(v string) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.ArrivalCity = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetArrivalTerminal(v string) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.ArrivalTerminal = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetArrivalTime(v string) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.ArrivalTime = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetAvailability(v string) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.Availability = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetCabin(v string) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.Cabin = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetCabinClass(v string) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.CabinClass = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetCodeShare(v bool) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.CodeShare = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetDepartureAirport(v string) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.DepartureAirport = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetDepartureCity(v string) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.DepartureCity = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetDepartureTerminal(v string) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.DepartureTerminal = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetDepartureTime(v string) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.DepartureTime = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetEquipType(v string) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.EquipType = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetFlightDuration(v int32) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.FlightDuration = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetMarketingAirline(v string) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.MarketingAirline = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetMarketingFlightNo(v string) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.MarketingFlightNo = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetMarketingFlightNoInt(v int32) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.MarketingFlightNoInt = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetOperatingAirline(v string) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.OperatingAirline = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetOperatingFlightNo(v string) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.OperatingFlightNo = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetSegmentId(v string) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.SegmentId = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetStopCityList(v string) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.StopCityList = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) SetStopQuantity(v int32) *ChangeDetailResponseBodyDataLastJourneysSegmentList {
	s.StopQuantity = &v
	return s
}

func (s *ChangeDetailResponseBodyDataLastJourneysSegmentList) Validate() error {
	return dara.Validate(s)
}

type ChangeDetailResponseBodyDataOriginalJourneys struct {
	// Segment information
	SegmentList []*ChangeDetailResponseBodyDataOriginalJourneysSegmentList `json:"segment_list,omitempty" xml:"segment_list,omitempty" type:"Repeated"`
	// Number of transfers
	//
	// example:
	//
	// 0
	TransferCount *int32 `json:"transfer_count,omitempty" xml:"transfer_count,omitempty"`
}

func (s ChangeDetailResponseBodyDataOriginalJourneys) String() string {
	return dara.Prettify(s)
}

func (s ChangeDetailResponseBodyDataOriginalJourneys) GoString() string {
	return s.String()
}

func (s *ChangeDetailResponseBodyDataOriginalJourneys) GetSegmentList() []*ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	return s.SegmentList
}

func (s *ChangeDetailResponseBodyDataOriginalJourneys) GetTransferCount() *int32 {
	return s.TransferCount
}

func (s *ChangeDetailResponseBodyDataOriginalJourneys) SetSegmentList(v []*ChangeDetailResponseBodyDataOriginalJourneysSegmentList) *ChangeDetailResponseBodyDataOriginalJourneys {
	s.SegmentList = v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneys) SetTransferCount(v int32) *ChangeDetailResponseBodyDataOriginalJourneys {
	s.TransferCount = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneys) Validate() error {
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

type ChangeDetailResponseBodyDataOriginalJourneysSegmentList struct {
	// Arrival airport three-letter code (uppercase)
	//
	// example:
	//
	// MFM
	ArrivalAirport *string `json:"arrival_airport,omitempty" xml:"arrival_airport,omitempty"`
	// Arrival city three-letter code (uppercase)
	//
	// example:
	//
	// MFM
	ArrivalCity *string `json:"arrival_city,omitempty" xml:"arrival_city,omitempty"`
	// Arrival terminal of the flight
	//
	// example:
	//
	// T1
	ArrivalTerminal *string `json:"arrival_terminal,omitempty" xml:"arrival_terminal,omitempty"`
	// Flight arrival date and time in string format (yyyy-MM-dd HH:mm:ss)
	//
	// example:
	//
	// 2023-03-10 10:40:00
	ArrivalTime *string `json:"arrival_time,omitempty" xml:"arrival_time,omitempty"`
	// Number of available seats
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
	// service class ( compartment )
	//
	// example:
	//
	// Y
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// Whether it is a codeshare flight
	//
	// example:
	//
	// false
	CodeShare *bool `json:"code_share,omitempty" xml:"code_share,omitempty"`
	// Departure airport three-letter code (uppercase)
	//
	// example:
	//
	// PVG
	DepartureAirport *string `json:"departure_airport,omitempty" xml:"departure_airport,omitempty"`
	// Departure city three-letter code (uppercase)
	//
	// example:
	//
	// SHA
	DepartureCity *string `json:"departure_city,omitempty" xml:"departure_city,omitempty"`
	// Departure terminal of the flight
	//
	// example:
	//
	// T2
	DepartureTerminal *string `json:"departure_terminal,omitempty" xml:"departure_terminal,omitempty"`
	// Flight departure date and time in string format (yyyy-MM-dd HH:mm:ss)
	//
	// example:
	//
	// 2023-03-10 07:55:00
	DepartureTime *string `json:"departure_time,omitempty" xml:"departure_time,omitempty"`
	// Aircraft type
	//
	// example:
	//
	// 32Q
	EquipType *string `json:"equip_type,omitempty" xml:"equip_type,omitempty"`
	// Flight duration in minutes
	//
	// example:
	//
	// 165
	FlightDuration *int32 `json:"flight_duration,omitempty" xml:"flight_duration,omitempty"`
	// Marketing airline (e.g., HO)
	//
	// example:
	//
	// HO
	MarketingAirline *string `json:"marketing_airline,omitempty" xml:"marketing_airline,omitempty"`
	// Marketing flight number (e.g., HO1295)
	//
	// example:
	//
	// HO1295
	MarketingFlightNo *string `json:"marketing_flight_no,omitempty" xml:"marketing_flight_no,omitempty"`
	// Marketing airline\\"s numeric flight number (e.g., 1295)
	//
	// example:
	//
	// 1295
	MarketingFlightNoInt *int32 `json:"marketing_flight_no_int,omitempty" xml:"marketing_flight_no_int,omitempty"`
	// Operating airline (e.g., CX)
	//
	// example:
	//
	// HO
	OperatingAirline *string `json:"operating_airline,omitempty" xml:"operating_airline,omitempty"`
	// Operating airline\\"s flight number (e.g., CX601)
	//
	// example:
	//
	// HO1295
	OperatingFlightNo *string `json:"operating_flight_no,omitempty" xml:"operating_flight_no,omitempty"`
	// Segment ID format: flight number + departure airport + arrival airport + departure date (yyyyMMdd)
	//
	// example:
	//
	// HO1295-PVG-MFM-20230310
	SegmentId *string `json:"segment_id,omitempty" xml:"segment_id,omitempty"`
	// List of stop cities, with values when stopQuantity > 0, separated by commas
	//
	// example:
	//
	// SEL,HKG
	StopCityList *string `json:"stop_city_list,omitempty" xml:"stop_city_list,omitempty"`
	// Number of stop cities
	//
	// example:
	//
	// 0
	StopQuantity *int32 `json:"stop_quantity,omitempty" xml:"stop_quantity,omitempty"`
}

func (s ChangeDetailResponseBodyDataOriginalJourneysSegmentList) String() string {
	return dara.Prettify(s)
}

func (s ChangeDetailResponseBodyDataOriginalJourneysSegmentList) GoString() string {
	return s.String()
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) GetArrivalAirport() *string {
	return s.ArrivalAirport
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) GetArrivalCity() *string {
	return s.ArrivalCity
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) GetArrivalTerminal() *string {
	return s.ArrivalTerminal
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) GetArrivalTime() *string {
	return s.ArrivalTime
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) GetAvailability() *string {
	return s.Availability
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) GetCabin() *string {
	return s.Cabin
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) GetCabinClass() *string {
	return s.CabinClass
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) GetCodeShare() *bool {
	return s.CodeShare
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) GetDepartureAirport() *string {
	return s.DepartureAirport
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) GetDepartureCity() *string {
	return s.DepartureCity
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) GetDepartureTerminal() *string {
	return s.DepartureTerminal
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) GetDepartureTime() *string {
	return s.DepartureTime
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) GetEquipType() *string {
	return s.EquipType
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) GetFlightDuration() *int32 {
	return s.FlightDuration
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) GetMarketingAirline() *string {
	return s.MarketingAirline
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) GetMarketingFlightNo() *string {
	return s.MarketingFlightNo
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) GetMarketingFlightNoInt() *int32 {
	return s.MarketingFlightNoInt
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) GetOperatingAirline() *string {
	return s.OperatingAirline
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) GetOperatingFlightNo() *string {
	return s.OperatingFlightNo
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) GetSegmentId() *string {
	return s.SegmentId
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) GetStopCityList() *string {
	return s.StopCityList
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) GetStopQuantity() *int32 {
	return s.StopQuantity
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetArrivalAirport(v string) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.ArrivalAirport = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetArrivalCity(v string) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.ArrivalCity = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetArrivalTerminal(v string) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.ArrivalTerminal = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetArrivalTime(v string) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.ArrivalTime = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetAvailability(v string) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.Availability = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetCabin(v string) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.Cabin = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetCabinClass(v string) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.CabinClass = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetCodeShare(v bool) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.CodeShare = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetDepartureAirport(v string) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.DepartureAirport = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetDepartureCity(v string) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.DepartureCity = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetDepartureTerminal(v string) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.DepartureTerminal = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetDepartureTime(v string) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.DepartureTime = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetEquipType(v string) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.EquipType = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetFlightDuration(v int32) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.FlightDuration = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetMarketingAirline(v string) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.MarketingAirline = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetMarketingFlightNo(v string) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.MarketingFlightNo = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetMarketingFlightNoInt(v int32) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.MarketingFlightNoInt = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetOperatingAirline(v string) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.OperatingAirline = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetOperatingFlightNo(v string) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.OperatingFlightNo = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetSegmentId(v string) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.SegmentId = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetStopCityList(v string) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.StopCityList = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) SetStopQuantity(v int32) *ChangeDetailResponseBodyDataOriginalJourneysSegmentList {
	s.StopQuantity = &v
	return s
}

func (s *ChangeDetailResponseBodyDataOriginalJourneysSegmentList) Validate() error {
	return dara.Validate(s)
}
