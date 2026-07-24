// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeDetailListOfOrderNumResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ChangeDetailListOfOrderNumResponseBody
	GetRequestId() *string
	SetData(v *ChangeDetailListOfOrderNumResponseBodyData) *ChangeDetailListOfOrderNumResponseBody
	GetData() *ChangeDetailListOfOrderNumResponseBodyData
	SetErrorCode(v string) *ChangeDetailListOfOrderNumResponseBody
	GetErrorCode() *string
	SetErrorData(v interface{}) *ChangeDetailListOfOrderNumResponseBody
	GetErrorData() interface{}
	SetErrorMsg(v string) *ChangeDetailListOfOrderNumResponseBody
	GetErrorMsg() *string
	SetStatus(v int32) *ChangeDetailListOfOrderNumResponseBody
	GetStatus() *int32
	SetSuccess(v bool) *ChangeDetailListOfOrderNumResponseBody
	GetSuccess() *bool
}

type ChangeDetailListOfOrderNumResponseBody struct {
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The data returned for a successful request.
	Data *ChangeDetailListOfOrderNumResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s ChangeDetailListOfOrderNumResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeDetailListOfOrderNumResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfOrderNumResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeDetailListOfOrderNumResponseBody) GetData() *ChangeDetailListOfOrderNumResponseBodyData {
	return s.Data
}

func (s *ChangeDetailListOfOrderNumResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ChangeDetailListOfOrderNumResponseBody) GetErrorData() interface{} {
	return s.ErrorData
}

func (s *ChangeDetailListOfOrderNumResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ChangeDetailListOfOrderNumResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *ChangeDetailListOfOrderNumResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChangeDetailListOfOrderNumResponseBody) SetRequestId(v string) *ChangeDetailListOfOrderNumResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBody) SetData(v *ChangeDetailListOfOrderNumResponseBodyData) *ChangeDetailListOfOrderNumResponseBody {
	s.Data = v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBody) SetErrorCode(v string) *ChangeDetailListOfOrderNumResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBody) SetErrorData(v interface{}) *ChangeDetailListOfOrderNumResponseBody {
	s.ErrorData = v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBody) SetErrorMsg(v string) *ChangeDetailListOfOrderNumResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBody) SetStatus(v int32) *ChangeDetailListOfOrderNumResponseBody {
	s.Status = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBody) SetSuccess(v bool) *ChangeDetailListOfOrderNumResponseBody {
	s.Success = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChangeDetailListOfOrderNumResponseBodyData struct {
	// The data list.
	List []*ChangeDetailListOfOrderNumResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// The pagination information.
	Pagination *ChangeDetailListOfOrderNumResponseBodyDataPagination `json:"pagination,omitempty" xml:"pagination,omitempty" type:"Struct"`
}

func (s ChangeDetailListOfOrderNumResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ChangeDetailListOfOrderNumResponseBodyData) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfOrderNumResponseBodyData) GetList() []*ChangeDetailListOfOrderNumResponseBodyDataList {
	return s.List
}

func (s *ChangeDetailListOfOrderNumResponseBodyData) GetPagination() *ChangeDetailListOfOrderNumResponseBodyDataPagination {
	return s.Pagination
}

func (s *ChangeDetailListOfOrderNumResponseBodyData) SetList(v []*ChangeDetailListOfOrderNumResponseBodyDataList) *ChangeDetailListOfOrderNumResponseBodyData {
	s.List = v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyData) SetPagination(v *ChangeDetailListOfOrderNumResponseBodyDataPagination) *ChangeDetailListOfOrderNumResponseBodyData {
	s.Pagination = v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Pagination != nil {
		if err := s.Pagination.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChangeDetailListOfOrderNumResponseBodyDataList struct {
	// The change fee details at the passenger level.
	ChangeFeeDetails []*ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetails `json:"change_fee_details,omitempty" xml:"change_fee_details,omitempty" type:"Repeated"`
	// The change order number.
	//
	// example:
	//
	// 4988430***950
	ChangeOrderNum *int64 `json:"change_order_num,omitempty" xml:"change_order_num,omitempty"`
	// The list of passengers for the change order.
	ChangePassengers []*ChangeDetailListOfOrderNumResponseBodyDataListChangePassengers `json:"change_passengers,omitempty" xml:"change_passengers,omitempty" type:"Repeated"`
	// The change reason type. Valid values:
	//
	// - 0: voluntary date change
	//
	// - 1: flight schedule change or flight cancellation
	//
	// - 2: change due to epidemic.
	//
	// example:
	//
	// 1
	ChangeReasonType *int32 `json:"change_reason_type,omitempty" xml:"change_reason_type,omitempty"`
	// The journeys after the change.
	ChangedJourneys []*ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneys `json:"changed_journeys,omitempty" xml:"changed_journeys,omitempty" type:"Repeated"`
	// The reason for closing the change order.
	//
	// example:
	//
	// reason desc
	CloseReason *string `json:"close_reason,omitempty" xml:"close_reason,omitempty"`
	// The time when the order was closed, in UTC timestamp.
	//
	// example:
	//
	// 1677415244000
	CloseUtcTime *int64 `json:"close_utc_time,omitempty" xml:"close_utc_time,omitempty"`
	// The contact information for the change order.
	Contact *ChangeDetailListOfOrderNumResponseBodyDataListContact `json:"contact,omitempty" xml:"contact,omitempty" type:"Struct"`
	// The creation time of the change order, in UTC timestamp.
	//
	// example:
	//
	// 1677415276000
	CreateUtcTime *int64 `json:"create_utc_time,omitempty" xml:"create_utc_time,omitempty"`
	// The latest payment deadline for the buyer, in UTC timestamp.
	//
	// example:
	//
	// 1677415278000
	LastConfirmUtcTime *int64 `json:"last_confirm_utc_time,omitempty" xml:"last_confirm_utc_time,omitempty"`
	// The journeys from the previous change.
	LastJourneys []*ChangeDetailListOfOrderNumResponseBodyDataListLastJourneys `json:"last_journeys,omitempty" xml:"last_journeys,omitempty" type:"Repeated"`
	// The order number.
	//
	// example:
	//
	// 5988430***541
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
	// The change order status. Valid values:
	//
	// - 0: initial state
	//
	// - 1: pending payment
	//
	// - 2: payment successful
	//
	// - 3: change successful
	//
	// - 4: change closed.
	//
	// example:
	//
	// 2
	OrderStatus *int32 `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// The original journeys.
	OriginalJourneys []*ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneys `json:"original_journeys,omitempty" xml:"original_journeys,omitempty" type:"Repeated"`
	// The payment status. Valid values:
	//
	// - 0: initial state
	//
	// - 1: pending payment
	//
	// - 2: payment successful
	//
	// - 3: transaction transfer successful
	//
	// - 4: paid order closed successfully
	//
	// - 5: unpaid order closed successfully.
	//
	// example:
	//
	// 2
	PayStatus *int32 `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// The time when the buyer completed the payment, in UTC timestamp.
	//
	// example:
	//
	// 1677415255000
	PaySuccessUtcTime *int64 `json:"pay_success_utc_time,omitempty" xml:"pay_success_utc_time,omitempty"`
	// The total payment amount of the change order.
	//
	// example:
	//
	// 300
	TotalAmount *float64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// The transaction number.
	//
	// example:
	//
	// hkduendkd-2023-dj0
	TransactionNo *string `json:"transaction_no,omitempty" xml:"transaction_no,omitempty"`
}

func (s ChangeDetailListOfOrderNumResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ChangeDetailListOfOrderNumResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) GetChangeFeeDetails() []*ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetails {
	return s.ChangeFeeDetails
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) GetChangeOrderNum() *int64 {
	return s.ChangeOrderNum
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) GetChangePassengers() []*ChangeDetailListOfOrderNumResponseBodyDataListChangePassengers {
	return s.ChangePassengers
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) GetChangeReasonType() *int32 {
	return s.ChangeReasonType
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) GetChangedJourneys() []*ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneys {
	return s.ChangedJourneys
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) GetCloseReason() *string {
	return s.CloseReason
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) GetCloseUtcTime() *int64 {
	return s.CloseUtcTime
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) GetContact() *ChangeDetailListOfOrderNumResponseBodyDataListContact {
	return s.Contact
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) GetCreateUtcTime() *int64 {
	return s.CreateUtcTime
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) GetLastConfirmUtcTime() *int64 {
	return s.LastConfirmUtcTime
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) GetLastJourneys() []*ChangeDetailListOfOrderNumResponseBodyDataListLastJourneys {
	return s.LastJourneys
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) GetOrderNum() *int64 {
	return s.OrderNum
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) GetOrderStatus() *int32 {
	return s.OrderStatus
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) GetOriginalJourneys() []*ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneys {
	return s.OriginalJourneys
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) GetPayStatus() *int32 {
	return s.PayStatus
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) GetPaySuccessUtcTime() *int64 {
	return s.PaySuccessUtcTime
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) GetTotalAmount() *float64 {
	return s.TotalAmount
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) GetTransactionNo() *string {
	return s.TransactionNo
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) SetChangeFeeDetails(v []*ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetails) *ChangeDetailListOfOrderNumResponseBodyDataList {
	s.ChangeFeeDetails = v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) SetChangeOrderNum(v int64) *ChangeDetailListOfOrderNumResponseBodyDataList {
	s.ChangeOrderNum = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) SetChangePassengers(v []*ChangeDetailListOfOrderNumResponseBodyDataListChangePassengers) *ChangeDetailListOfOrderNumResponseBodyDataList {
	s.ChangePassengers = v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) SetChangeReasonType(v int32) *ChangeDetailListOfOrderNumResponseBodyDataList {
	s.ChangeReasonType = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) SetChangedJourneys(v []*ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneys) *ChangeDetailListOfOrderNumResponseBodyDataList {
	s.ChangedJourneys = v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) SetCloseReason(v string) *ChangeDetailListOfOrderNumResponseBodyDataList {
	s.CloseReason = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) SetCloseUtcTime(v int64) *ChangeDetailListOfOrderNumResponseBodyDataList {
	s.CloseUtcTime = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) SetContact(v *ChangeDetailListOfOrderNumResponseBodyDataListContact) *ChangeDetailListOfOrderNumResponseBodyDataList {
	s.Contact = v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) SetCreateUtcTime(v int64) *ChangeDetailListOfOrderNumResponseBodyDataList {
	s.CreateUtcTime = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) SetLastConfirmUtcTime(v int64) *ChangeDetailListOfOrderNumResponseBodyDataList {
	s.LastConfirmUtcTime = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) SetLastJourneys(v []*ChangeDetailListOfOrderNumResponseBodyDataListLastJourneys) *ChangeDetailListOfOrderNumResponseBodyDataList {
	s.LastJourneys = v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) SetOrderNum(v int64) *ChangeDetailListOfOrderNumResponseBodyDataList {
	s.OrderNum = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) SetOrderStatus(v int32) *ChangeDetailListOfOrderNumResponseBodyDataList {
	s.OrderStatus = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) SetOriginalJourneys(v []*ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneys) *ChangeDetailListOfOrderNumResponseBodyDataList {
	s.OriginalJourneys = v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) SetPayStatus(v int32) *ChangeDetailListOfOrderNumResponseBodyDataList {
	s.PayStatus = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) SetPaySuccessUtcTime(v int64) *ChangeDetailListOfOrderNumResponseBodyDataList {
	s.PaySuccessUtcTime = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) SetTotalAmount(v float64) *ChangeDetailListOfOrderNumResponseBodyDataList {
	s.TotalAmount = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) SetTransactionNo(v string) *ChangeDetailListOfOrderNumResponseBodyDataList {
	s.TransactionNo = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataList) Validate() error {
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

type ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetails struct {
	// The change fee details for the passenger.
	ChangeFee *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsChangeFee `json:"change_fee,omitempty" xml:"change_fee,omitempty" type:"Struct"`
	// The passenger information for the change.
	Passenger *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsPassenger `json:"passenger,omitempty" xml:"passenger,omitempty" type:"Struct"`
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetails) String() string {
	return dara.Prettify(s)
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetails) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetails) GetChangeFee() *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsChangeFee {
	return s.ChangeFee
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetails) GetPassenger() *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsPassenger {
	return s.Passenger
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetails) SetChangeFee(v *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsChangeFee) *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetails {
	s.ChangeFee = v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetails) SetPassenger(v *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsPassenger) *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetails {
	s.Passenger = v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetails) Validate() error {
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

type ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsChangeFee struct {
	// The service fee.
	//
	// example:
	//
	// 50
	ServiceFee     *float64 `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
	SuezServiceFee *float64 `json:"suez_service_fee,omitempty" xml:"suez_service_fee,omitempty"`
	// The change tax fee.
	//
	// example:
	//
	// 20
	TaxFee *float64 `json:"tax_fee,omitempty" xml:"tax_fee,omitempty"`
	// The upgrade fee.
	//
	// example:
	//
	// 30
	UpgradeFee *float64 `json:"upgrade_fee,omitempty" xml:"upgrade_fee,omitempty"`
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsChangeFee) String() string {
	return dara.Prettify(s)
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsChangeFee) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsChangeFee) GetServiceFee() *float64 {
	return s.ServiceFee
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsChangeFee) GetSuezServiceFee() *float64 {
	return s.SuezServiceFee
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsChangeFee) GetTaxFee() *float64 {
	return s.TaxFee
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsChangeFee) GetUpgradeFee() *float64 {
	return s.UpgradeFee
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsChangeFee) SetServiceFee(v float64) *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsChangeFee {
	s.ServiceFee = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsChangeFee) SetSuezServiceFee(v float64) *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsChangeFee {
	s.SuezServiceFee = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsChangeFee) SetTaxFee(v float64) *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsChangeFee {
	s.TaxFee = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsChangeFee) SetUpgradeFee(v float64) *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsChangeFee {
	s.UpgradeFee = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsChangeFee) Validate() error {
	return dara.Validate(s)
}

type ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsPassenger struct {
	// The document number.
	//
	// example:
	//
	// 411***********4411
	Document *string `json:"document,omitempty" xml:"document,omitempty"`
	// The first name of the passenger.
	//
	// example:
	//
	// SAN
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// The last name of the passenger.
	//
	// example:
	//
	// ZHANG
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsPassenger) String() string {
	return dara.Prettify(s)
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsPassenger) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsPassenger) GetDocument() *string {
	return s.Document
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsPassenger) GetFirstName() *string {
	return s.FirstName
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsPassenger) GetLastName() *string {
	return s.LastName
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsPassenger) SetDocument(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsPassenger {
	s.Document = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsPassenger) SetFirstName(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsPassenger {
	s.FirstName = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsPassenger) SetLastName(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsPassenger {
	s.LastName = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangeFeeDetailsPassenger) Validate() error {
	return dara.Validate(s)
}

type ChangeDetailListOfOrderNumResponseBodyDataListChangePassengers struct {
	// The document number.
	//
	// example:
	//
	// 411***********4411
	Document *string `json:"document,omitempty" xml:"document,omitempty"`
	// The first name of the passenger.
	//
	// example:
	//
	// SAN
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// The last name of the passenger.
	//
	// example:
	//
	// ZHANG
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListChangePassengers) String() string {
	return dara.Prettify(s)
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListChangePassengers) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangePassengers) GetDocument() *string {
	return s.Document
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangePassengers) GetFirstName() *string {
	return s.FirstName
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangePassengers) GetLastName() *string {
	return s.LastName
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangePassengers) SetDocument(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangePassengers {
	s.Document = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangePassengers) SetFirstName(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangePassengers {
	s.FirstName = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangePassengers) SetLastName(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangePassengers {
	s.LastName = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangePassengers) Validate() error {
	return dara.Validate(s)
}

type ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneys struct {
	// The segment information.
	SegmentList []*ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList `json:"segment_list,omitempty" xml:"segment_list,omitempty" type:"Repeated"`
	// The number of transfers.
	//
	// example:
	//
	// 0
	TransferCount *int32 `json:"transfer_count,omitempty" xml:"transfer_count,omitempty"`
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneys) String() string {
	return dara.Prettify(s)
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneys) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneys) GetSegmentList() []*ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	return s.SegmentList
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneys) GetTransferCount() *int32 {
	return s.TransferCount
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneys) SetSegmentList(v []*ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneys {
	s.SegmentList = v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneys) SetTransferCount(v int32) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneys {
	s.TransferCount = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneys) Validate() error {
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

type ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList struct {
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
	// The cabin code.
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
	// Indicates whether this is a codeshare flight.
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
	// The flight duration in minutes.
	//
	// example:
	//
	// 165
	FlightDuration *int32 `json:"flight_duration,omitempty" xml:"flight_duration,omitempty"`
	// The marketing airline code (such as HO).
	//
	// example:
	//
	// HO
	MarketingAirline *string `json:"marketing_airline,omitempty" xml:"marketing_airline,omitempty"`
	// The marketing flight number (such as HO1295).
	//
	// example:
	//
	// HO1295
	MarketingFlightNo *string `json:"marketing_flight_no,omitempty" xml:"marketing_flight_no,omitempty"`
	// The numeric marketing flight number (such as 1295).
	//
	// example:
	//
	// 1295
	MarketingFlightNoInt *int32 `json:"marketing_flight_no_int,omitempty" xml:"marketing_flight_no_int,omitempty"`
	// The operating airline code (such as CX).
	//
	// example:
	//
	// HO
	OperatingAirline *string `json:"operating_airline,omitempty" xml:"operating_airline,omitempty"`
	// The operating flight number (such as CX601).
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
	// The list of stopover cities. This field has a value when stopQuantity is greater than 0. Multiple cities are separated by commas.
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

func (s ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) String() string {
	return dara.Prettify(s)
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) GetArrivalAirport() *string {
	return s.ArrivalAirport
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) GetArrivalCity() *string {
	return s.ArrivalCity
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) GetArrivalTerminal() *string {
	return s.ArrivalTerminal
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) GetArrivalTime() *string {
	return s.ArrivalTime
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) GetAvailability() *string {
	return s.Availability
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) GetCabin() *string {
	return s.Cabin
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) GetCabinClass() *string {
	return s.CabinClass
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) GetCodeShare() *bool {
	return s.CodeShare
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) GetDepartureAirport() *string {
	return s.DepartureAirport
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) GetDepartureCity() *string {
	return s.DepartureCity
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) GetDepartureTerminal() *string {
	return s.DepartureTerminal
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) GetDepartureTime() *string {
	return s.DepartureTime
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) GetEquipType() *string {
	return s.EquipType
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) GetFlightDuration() *int32 {
	return s.FlightDuration
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) GetMarketingAirline() *string {
	return s.MarketingAirline
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) GetMarketingFlightNo() *string {
	return s.MarketingFlightNo
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) GetMarketingFlightNoInt() *int32 {
	return s.MarketingFlightNoInt
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) GetOperatingAirline() *string {
	return s.OperatingAirline
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) GetOperatingFlightNo() *string {
	return s.OperatingFlightNo
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) GetSegmentId() *string {
	return s.SegmentId
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) GetStopCityList() *string {
	return s.StopCityList
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) GetStopQuantity() *int32 {
	return s.StopQuantity
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetArrivalAirport(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.ArrivalAirport = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetArrivalCity(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.ArrivalCity = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetArrivalTerminal(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.ArrivalTerminal = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetArrivalTime(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.ArrivalTime = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetAvailability(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.Availability = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetCabin(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.Cabin = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetCabinClass(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.CabinClass = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetCodeShare(v bool) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.CodeShare = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetDepartureAirport(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.DepartureAirport = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetDepartureCity(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.DepartureCity = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetDepartureTerminal(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.DepartureTerminal = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetDepartureTime(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.DepartureTime = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetEquipType(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.EquipType = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetFlightDuration(v int32) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.FlightDuration = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetMarketingAirline(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.MarketingAirline = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetMarketingFlightNo(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.MarketingFlightNo = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetMarketingFlightNoInt(v int32) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.MarketingFlightNoInt = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetOperatingAirline(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.OperatingAirline = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetOperatingFlightNo(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.OperatingFlightNo = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetSegmentId(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.SegmentId = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetStopCityList(v string) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.StopCityList = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) SetStopQuantity(v int32) *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList {
	s.StopQuantity = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListChangedJourneysSegmentList) Validate() error {
	return dara.Validate(s)
}

type ChangeDetailListOfOrderNumResponseBodyDataListContact struct {
	// The email address.
	//
	// example:
	//
	// gao******@gmail.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// The country calling code.
	//
	// example:
	//
	// 86
	MobileCountryCode *string `json:"mobile_country_code,omitempty" xml:"mobile_country_code,omitempty"`
	// The mobile phone number of the contact.
	//
	// example:
	//
	// 183*****92
	MobilePhoneNum *string `json:"mobile_phone_num,omitempty" xml:"mobile_phone_num,omitempty"`
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListContact) String() string {
	return dara.Prettify(s)
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListContact) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListContact) GetEmail() *string {
	return s.Email
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListContact) GetMobileCountryCode() *string {
	return s.MobileCountryCode
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListContact) GetMobilePhoneNum() *string {
	return s.MobilePhoneNum
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListContact) SetEmail(v string) *ChangeDetailListOfOrderNumResponseBodyDataListContact {
	s.Email = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListContact) SetMobileCountryCode(v string) *ChangeDetailListOfOrderNumResponseBodyDataListContact {
	s.MobileCountryCode = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListContact) SetMobilePhoneNum(v string) *ChangeDetailListOfOrderNumResponseBodyDataListContact {
	s.MobilePhoneNum = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListContact) Validate() error {
	return dara.Validate(s)
}

type ChangeDetailListOfOrderNumResponseBodyDataListLastJourneys struct {
	// The segment information.
	SegmentList []*ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList `json:"segment_list,omitempty" xml:"segment_list,omitempty" type:"Repeated"`
	// The number of transfers.
	//
	// example:
	//
	// 0
	TransferCount *int32 `json:"transfer_count,omitempty" xml:"transfer_count,omitempty"`
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListLastJourneys) String() string {
	return dara.Prettify(s)
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListLastJourneys) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneys) GetSegmentList() []*ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	return s.SegmentList
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneys) GetTransferCount() *int32 {
	return s.TransferCount
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneys) SetSegmentList(v []*ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneys {
	s.SegmentList = v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneys) SetTransferCount(v int32) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneys {
	s.TransferCount = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneys) Validate() error {
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

type ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList struct {
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
	// The cabin code.
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
	// Indicates whether this is a codeshare flight.
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
	// The flight duration in minutes.
	//
	// example:
	//
	// 165
	FlightDuration *int32 `json:"flight_duration,omitempty" xml:"flight_duration,omitempty"`
	// The marketing airline code (such as HO).
	//
	// example:
	//
	// HO
	MarketingAirline *string `json:"marketing_airline,omitempty" xml:"marketing_airline,omitempty"`
	// The marketing flight number (such as HO1295).
	//
	// example:
	//
	// HO1295
	MarketingFlightNo *string `json:"marketing_flight_no,omitempty" xml:"marketing_flight_no,omitempty"`
	// The numeric marketing flight number (such as 1295).
	//
	// example:
	//
	// 1295
	MarketingFlightNoInt *int32 `json:"marketing_flight_no_int,omitempty" xml:"marketing_flight_no_int,omitempty"`
	// The operating airline code (such as CX).
	//
	// example:
	//
	// HO
	OperatingAirline *string `json:"operating_airline,omitempty" xml:"operating_airline,omitempty"`
	// The operating flight number (such as CX601).
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
	// The list of stopover cities. This field has a value when stopQuantity is greater than 0. Multiple cities are separated by commas.
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

func (s ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) String() string {
	return dara.Prettify(s)
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) GetArrivalAirport() *string {
	return s.ArrivalAirport
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) GetArrivalCity() *string {
	return s.ArrivalCity
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) GetArrivalTerminal() *string {
	return s.ArrivalTerminal
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) GetArrivalTime() *string {
	return s.ArrivalTime
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) GetAvailability() *string {
	return s.Availability
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) GetCabin() *string {
	return s.Cabin
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) GetCabinClass() *string {
	return s.CabinClass
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) GetCodeShare() *bool {
	return s.CodeShare
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) GetDepartureAirport() *string {
	return s.DepartureAirport
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) GetDepartureCity() *string {
	return s.DepartureCity
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) GetDepartureTerminal() *string {
	return s.DepartureTerminal
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) GetDepartureTime() *string {
	return s.DepartureTime
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) GetEquipType() *string {
	return s.EquipType
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) GetFlightDuration() *int32 {
	return s.FlightDuration
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) GetMarketingAirline() *string {
	return s.MarketingAirline
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) GetMarketingFlightNo() *string {
	return s.MarketingFlightNo
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) GetMarketingFlightNoInt() *int32 {
	return s.MarketingFlightNoInt
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) GetOperatingAirline() *string {
	return s.OperatingAirline
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) GetOperatingFlightNo() *string {
	return s.OperatingFlightNo
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) GetSegmentId() *string {
	return s.SegmentId
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) GetStopCityList() *string {
	return s.StopCityList
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) GetStopQuantity() *int32 {
	return s.StopQuantity
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetArrivalAirport(v string) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.ArrivalAirport = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetArrivalCity(v string) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.ArrivalCity = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetArrivalTerminal(v string) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.ArrivalTerminal = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetArrivalTime(v string) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.ArrivalTime = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetAvailability(v string) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.Availability = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetCabin(v string) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.Cabin = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetCabinClass(v string) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.CabinClass = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetCodeShare(v bool) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.CodeShare = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetDepartureAirport(v string) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.DepartureAirport = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetDepartureCity(v string) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.DepartureCity = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetDepartureTerminal(v string) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.DepartureTerminal = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetDepartureTime(v string) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.DepartureTime = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetEquipType(v string) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.EquipType = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetFlightDuration(v int32) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.FlightDuration = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetMarketingAirline(v string) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.MarketingAirline = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetMarketingFlightNo(v string) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.MarketingFlightNo = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetMarketingFlightNoInt(v int32) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.MarketingFlightNoInt = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetOperatingAirline(v string) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.OperatingAirline = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetOperatingFlightNo(v string) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.OperatingFlightNo = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetSegmentId(v string) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.SegmentId = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetStopCityList(v string) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.StopCityList = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) SetStopQuantity(v int32) *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList {
	s.StopQuantity = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListLastJourneysSegmentList) Validate() error {
	return dara.Validate(s)
}

type ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneys struct {
	// The segment information.
	SegmentList []*ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList `json:"segment_list,omitempty" xml:"segment_list,omitempty" type:"Repeated"`
	// The number of transfers.
	//
	// example:
	//
	// 0
	TransferCount *int32 `json:"transfer_count,omitempty" xml:"transfer_count,omitempty"`
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneys) String() string {
	return dara.Prettify(s)
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneys) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneys) GetSegmentList() []*ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	return s.SegmentList
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneys) GetTransferCount() *int32 {
	return s.TransferCount
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneys) SetSegmentList(v []*ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneys {
	s.SegmentList = v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneys) SetTransferCount(v int32) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneys {
	s.TransferCount = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneys) Validate() error {
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

type ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList struct {
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
	// The cabin code.
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
	// Indicates whether this is a codeshare flight.
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
	// The flight duration in minutes.
	//
	// example:
	//
	// 165
	FlightDuration *int32 `json:"flight_duration,omitempty" xml:"flight_duration,omitempty"`
	// The marketing airline code (such as HO).
	//
	// example:
	//
	// HO
	MarketingAirline *string `json:"marketing_airline,omitempty" xml:"marketing_airline,omitempty"`
	// The marketing flight number (such as HO1295).
	//
	// example:
	//
	// HO1295
	MarketingFlightNo *string `json:"marketing_flight_no,omitempty" xml:"marketing_flight_no,omitempty"`
	// The numeric marketing flight number (such as 1295).
	//
	// example:
	//
	// 1295
	MarketingFlightNoInt *int32 `json:"marketing_flight_no_int,omitempty" xml:"marketing_flight_no_int,omitempty"`
	// The operating airline code (such as CX).
	//
	// example:
	//
	// HO
	OperatingAirline *string `json:"operating_airline,omitempty" xml:"operating_airline,omitempty"`
	// The operating flight number (such as CX601).
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
	// The list of stopover cities. This field has a value when stopQuantity is greater than 0. Multiple cities are separated by commas.
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

func (s ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) String() string {
	return dara.Prettify(s)
}

func (s ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) GetArrivalAirport() *string {
	return s.ArrivalAirport
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) GetArrivalCity() *string {
	return s.ArrivalCity
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) GetArrivalTerminal() *string {
	return s.ArrivalTerminal
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) GetArrivalTime() *string {
	return s.ArrivalTime
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) GetAvailability() *string {
	return s.Availability
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) GetCabin() *string {
	return s.Cabin
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) GetCabinClass() *string {
	return s.CabinClass
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) GetCodeShare() *bool {
	return s.CodeShare
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) GetDepartureAirport() *string {
	return s.DepartureAirport
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) GetDepartureCity() *string {
	return s.DepartureCity
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) GetDepartureTerminal() *string {
	return s.DepartureTerminal
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) GetDepartureTime() *string {
	return s.DepartureTime
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) GetEquipType() *string {
	return s.EquipType
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) GetFlightDuration() *int32 {
	return s.FlightDuration
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) GetMarketingAirline() *string {
	return s.MarketingAirline
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) GetMarketingFlightNo() *string {
	return s.MarketingFlightNo
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) GetMarketingFlightNoInt() *int32 {
	return s.MarketingFlightNoInt
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) GetOperatingAirline() *string {
	return s.OperatingAirline
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) GetOperatingFlightNo() *string {
	return s.OperatingFlightNo
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) GetSegmentId() *string {
	return s.SegmentId
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) GetStopCityList() *string {
	return s.StopCityList
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) GetStopQuantity() *int32 {
	return s.StopQuantity
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetArrivalAirport(v string) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.ArrivalAirport = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetArrivalCity(v string) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.ArrivalCity = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetArrivalTerminal(v string) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.ArrivalTerminal = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetArrivalTime(v string) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.ArrivalTime = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetAvailability(v string) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.Availability = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetCabin(v string) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.Cabin = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetCabinClass(v string) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.CabinClass = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetCodeShare(v bool) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.CodeShare = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetDepartureAirport(v string) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.DepartureAirport = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetDepartureCity(v string) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.DepartureCity = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetDepartureTerminal(v string) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.DepartureTerminal = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetDepartureTime(v string) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.DepartureTime = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetEquipType(v string) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.EquipType = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetFlightDuration(v int32) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.FlightDuration = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetMarketingAirline(v string) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.MarketingAirline = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetMarketingFlightNo(v string) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.MarketingFlightNo = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetMarketingFlightNoInt(v int32) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.MarketingFlightNoInt = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetOperatingAirline(v string) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.OperatingAirline = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetOperatingFlightNo(v string) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.OperatingFlightNo = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetSegmentId(v string) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.SegmentId = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetStopCityList(v string) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.StopCityList = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) SetStopQuantity(v int32) *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList {
	s.StopQuantity = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataListOriginalJourneysSegmentList) Validate() error {
	return dara.Validate(s)
}

type ChangeDetailListOfOrderNumResponseBodyDataPagination struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// The number of records per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 1
	TotalPage *int32 `json:"total_page,omitempty" xml:"total_page,omitempty"`
}

func (s ChangeDetailListOfOrderNumResponseBodyDataPagination) String() string {
	return dara.Prettify(s)
}

func (s ChangeDetailListOfOrderNumResponseBodyDataPagination) GoString() string {
	return s.String()
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataPagination) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataPagination) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataPagination) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataPagination) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataPagination) SetCurrentPage(v int32) *ChangeDetailListOfOrderNumResponseBodyDataPagination {
	s.CurrentPage = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataPagination) SetPageSize(v int32) *ChangeDetailListOfOrderNumResponseBodyDataPagination {
	s.PageSize = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataPagination) SetTotalCount(v int32) *ChangeDetailListOfOrderNumResponseBodyDataPagination {
	s.TotalCount = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataPagination) SetTotalPage(v int32) *ChangeDetailListOfOrderNumResponseBodyDataPagination {
	s.TotalPage = &v
	return s
}

func (s *ChangeDetailListOfOrderNumResponseBodyDataPagination) Validate() error {
	return dara.Validate(s)
}
