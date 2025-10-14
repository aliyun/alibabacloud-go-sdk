// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefundDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RefundDetailResponseBody
	GetRequestId() *string
	SetData(v *RefundDetailResponseBodyData) *RefundDetailResponseBody
	GetData() *RefundDetailResponseBodyData
	SetErrorCode(v string) *RefundDetailResponseBody
	GetErrorCode() *string
	SetErrorData(v interface{}) *RefundDetailResponseBody
	GetErrorData() interface{}
	SetErrorMsg(v string) *RefundDetailResponseBody
	GetErrorMsg() *string
	SetStatus(v int32) *RefundDetailResponseBody
	GetStatus() *int32
	SetSuccess(v bool) *RefundDetailResponseBody
	GetSuccess() *bool
}

type RefundDetailResponseBody struct {
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// data
	Data *RefundDetailResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s RefundDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefundDetailResponseBody) GoString() string {
	return s.String()
}

func (s *RefundDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefundDetailResponseBody) GetData() *RefundDetailResponseBodyData {
	return s.Data
}

func (s *RefundDetailResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RefundDetailResponseBody) GetErrorData() interface{} {
	return s.ErrorData
}

func (s *RefundDetailResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *RefundDetailResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *RefundDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RefundDetailResponseBody) SetRequestId(v string) *RefundDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefundDetailResponseBody) SetData(v *RefundDetailResponseBodyData) *RefundDetailResponseBody {
	s.Data = v
	return s
}

func (s *RefundDetailResponseBody) SetErrorCode(v string) *RefundDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RefundDetailResponseBody) SetErrorData(v interface{}) *RefundDetailResponseBody {
	s.ErrorData = v
	return s
}

func (s *RefundDetailResponseBody) SetErrorMsg(v string) *RefundDetailResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *RefundDetailResponseBody) SetStatus(v int32) *RefundDetailResponseBody {
	s.Status = &v
	return s
}

func (s *RefundDetailResponseBody) SetSuccess(v bool) *RefundDetailResponseBody {
	s.Success = &v
	return s
}

func (s *RefundDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RefundDetailResponseBodyData struct {
	// whether it is a supplementary refund order (if the refund amount is not enough, you can use RefundApply to create a supplementary refund order)
	//
	// example:
	//
	// false
	ContainMultiRefund *bool `json:"contain_multi_refund,omitempty" xml:"contain_multi_refund,omitempty"`
	// supplementary refund orders
	MultiRefundDetails []*RefundDetailResponseBodyDataMultiRefundDetails `json:"multi_refund_details,omitempty" xml:"multi_refund_details,omitempty" type:"Repeated"`
	// order number that returned by Book
	//
	// example:
	//
	// 4966***617111
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
	// refund details by passenger dimension
	PassengerRefundDetails []*RefundDetailResponseBodyDataPassengerRefundDetails `json:"passenger_refund_details,omitempty" xml:"passenger_refund_details,omitempty" type:"Repeated"`
	// refund completed time(timestamp)
	//
	// example:
	//
	// 1677229005000
	PaySuccessUtcTime *int64 `json:"pay_success_utc_time,omitempty" xml:"pay_success_utc_time,omitempty"`
	// URLs for refund attachments
	//
	// example:
	//
	// [zzz,yyy]
	RefundAttachmentUrls []*string `json:"refund_attachment_urls,omitempty" xml:"refund_attachment_urls,omitempty" type:"Repeated"`
	// refunded journey
	RefundJourneys []*RefundDetailResponseBodyDataRefundJourneys `json:"refund_journeys,omitempty" xml:"refund_journeys,omitempty" type:"Repeated"`
	// refund order number that returned by RefundApply
	//
	// example:
	//
	// 4966***617654
	RefundOrderNum *int64 `json:"refund_order_num,omitempty" xml:"refund_order_num,omitempty"`
	// reason for refund
	//
	// example:
	//
	// desc reason
	RefundReason *string `json:"refund_reason,omitempty" xml:"refund_reason,omitempty"`
	// refund type
	//
	// 2: voluntary application
	//
	// 5: flight delay or cancellation, flight schedule change, etc., due to airline reasons
	//
	// 6: health reasons with a certificate from a secondary class A hospital or above
	//
	// 7: non-voluntary confirmed guidance
	//
	// 100: non-voluntary non-confirmed guidance
	//
	// example:
	//
	// 5
	RefundType *int32 `json:"refund_type,omitempty" xml:"refund_type,omitempty"`
	// reason for refund rejection
	//
	// example:
	//
	// refuse reason
	RefuseReason *string `json:"refuse_reason,omitempty" xml:"refuse_reason,omitempty"`
	// refund order status
	//
	// 0: refund application
	//
	// 1: refund in progress
	//
	// 2: refund failed
	//
	// 3: refund succeeded
	//
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// transaction number
	//
	// example:
	//
	// 1677229005000
	TransactionNo *string `json:"transaction_no,omitempty" xml:"transaction_no,omitempty"`
	// refund order created time(timestamp)
	//
	// example:
	//
	// 1677229002000
	UtcCreateTime *int64 `json:"utc_create_time,omitempty" xml:"utc_create_time,omitempty"`
}

func (s RefundDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RefundDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *RefundDetailResponseBodyData) GetContainMultiRefund() *bool {
	return s.ContainMultiRefund
}

func (s *RefundDetailResponseBodyData) GetMultiRefundDetails() []*RefundDetailResponseBodyDataMultiRefundDetails {
	return s.MultiRefundDetails
}

func (s *RefundDetailResponseBodyData) GetOrderNum() *int64 {
	return s.OrderNum
}

func (s *RefundDetailResponseBodyData) GetPassengerRefundDetails() []*RefundDetailResponseBodyDataPassengerRefundDetails {
	return s.PassengerRefundDetails
}

func (s *RefundDetailResponseBodyData) GetPaySuccessUtcTime() *int64 {
	return s.PaySuccessUtcTime
}

func (s *RefundDetailResponseBodyData) GetRefundAttachmentUrls() []*string {
	return s.RefundAttachmentUrls
}

func (s *RefundDetailResponseBodyData) GetRefundJourneys() []*RefundDetailResponseBodyDataRefundJourneys {
	return s.RefundJourneys
}

func (s *RefundDetailResponseBodyData) GetRefundOrderNum() *int64 {
	return s.RefundOrderNum
}

func (s *RefundDetailResponseBodyData) GetRefundReason() *string {
	return s.RefundReason
}

func (s *RefundDetailResponseBodyData) GetRefundType() *int32 {
	return s.RefundType
}

func (s *RefundDetailResponseBodyData) GetRefuseReason() *string {
	return s.RefuseReason
}

func (s *RefundDetailResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *RefundDetailResponseBodyData) GetTransactionNo() *string {
	return s.TransactionNo
}

func (s *RefundDetailResponseBodyData) GetUtcCreateTime() *int64 {
	return s.UtcCreateTime
}

func (s *RefundDetailResponseBodyData) SetContainMultiRefund(v bool) *RefundDetailResponseBodyData {
	s.ContainMultiRefund = &v
	return s
}

func (s *RefundDetailResponseBodyData) SetMultiRefundDetails(v []*RefundDetailResponseBodyDataMultiRefundDetails) *RefundDetailResponseBodyData {
	s.MultiRefundDetails = v
	return s
}

func (s *RefundDetailResponseBodyData) SetOrderNum(v int64) *RefundDetailResponseBodyData {
	s.OrderNum = &v
	return s
}

func (s *RefundDetailResponseBodyData) SetPassengerRefundDetails(v []*RefundDetailResponseBodyDataPassengerRefundDetails) *RefundDetailResponseBodyData {
	s.PassengerRefundDetails = v
	return s
}

func (s *RefundDetailResponseBodyData) SetPaySuccessUtcTime(v int64) *RefundDetailResponseBodyData {
	s.PaySuccessUtcTime = &v
	return s
}

func (s *RefundDetailResponseBodyData) SetRefundAttachmentUrls(v []*string) *RefundDetailResponseBodyData {
	s.RefundAttachmentUrls = v
	return s
}

func (s *RefundDetailResponseBodyData) SetRefundJourneys(v []*RefundDetailResponseBodyDataRefundJourneys) *RefundDetailResponseBodyData {
	s.RefundJourneys = v
	return s
}

func (s *RefundDetailResponseBodyData) SetRefundOrderNum(v int64) *RefundDetailResponseBodyData {
	s.RefundOrderNum = &v
	return s
}

func (s *RefundDetailResponseBodyData) SetRefundReason(v string) *RefundDetailResponseBodyData {
	s.RefundReason = &v
	return s
}

func (s *RefundDetailResponseBodyData) SetRefundType(v int32) *RefundDetailResponseBodyData {
	s.RefundType = &v
	return s
}

func (s *RefundDetailResponseBodyData) SetRefuseReason(v string) *RefundDetailResponseBodyData {
	s.RefuseReason = &v
	return s
}

func (s *RefundDetailResponseBodyData) SetStatus(v int32) *RefundDetailResponseBodyData {
	s.Status = &v
	return s
}

func (s *RefundDetailResponseBodyData) SetTransactionNo(v string) *RefundDetailResponseBodyData {
	s.TransactionNo = &v
	return s
}

func (s *RefundDetailResponseBodyData) SetUtcCreateTime(v int64) *RefundDetailResponseBodyData {
	s.UtcCreateTime = &v
	return s
}

func (s *RefundDetailResponseBodyData) Validate() error {
	if s.MultiRefundDetails != nil {
		for _, item := range s.MultiRefundDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PassengerRefundDetails != nil {
		for _, item := range s.PassengerRefundDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RefundJourneys != nil {
		for _, item := range s.RefundJourneys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RefundDetailResponseBodyDataMultiRefundDetails struct {
	// supplementary refund order number
	//
	// example:
	//
	// 498843***6950
	MultiRefundOrderNum *int64 `json:"multi_refund_order_num,omitempty" xml:"multi_refund_order_num,omitempty"`
	// transaction number of the supplementary refund order
	//
	// example:
	//
	// 498843***6950
	MultiRefundTransactionNo *string `json:"multi_refund_transaction_no,omitempty" xml:"multi_refund_transaction_no,omitempty"`
	// supplementary refund details in passenger dimension
	PassengerMultiRefundDetails []*RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetails `json:"passenger_multi_refund_details,omitempty" xml:"passenger_multi_refund_details,omitempty" type:"Repeated"`
}

func (s RefundDetailResponseBodyDataMultiRefundDetails) String() string {
	return dara.Prettify(s)
}

func (s RefundDetailResponseBodyDataMultiRefundDetails) GoString() string {
	return s.String()
}

func (s *RefundDetailResponseBodyDataMultiRefundDetails) GetMultiRefundOrderNum() *int64 {
	return s.MultiRefundOrderNum
}

func (s *RefundDetailResponseBodyDataMultiRefundDetails) GetMultiRefundTransactionNo() *string {
	return s.MultiRefundTransactionNo
}

func (s *RefundDetailResponseBodyDataMultiRefundDetails) GetPassengerMultiRefundDetails() []*RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetails {
	return s.PassengerMultiRefundDetails
}

func (s *RefundDetailResponseBodyDataMultiRefundDetails) SetMultiRefundOrderNum(v int64) *RefundDetailResponseBodyDataMultiRefundDetails {
	s.MultiRefundOrderNum = &v
	return s
}

func (s *RefundDetailResponseBodyDataMultiRefundDetails) SetMultiRefundTransactionNo(v string) *RefundDetailResponseBodyDataMultiRefundDetails {
	s.MultiRefundTransactionNo = &v
	return s
}

func (s *RefundDetailResponseBodyDataMultiRefundDetails) SetPassengerMultiRefundDetails(v []*RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetails) *RefundDetailResponseBodyDataMultiRefundDetails {
	s.PassengerMultiRefundDetails = v
	return s
}

func (s *RefundDetailResponseBodyDataMultiRefundDetails) Validate() error {
	if s.PassengerMultiRefundDetails != nil {
		for _, item := range s.PassengerMultiRefundDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetails struct {
	// amount of the supplementary refund for the change order
	//
	// example:
	//
	// 30
	ChangeOrderRefundFee *float64 `json:"change_order_refund_fee,omitempty" xml:"change_order_refund_fee,omitempty"`
	// amount of the supplementary refund for the original order
	//
	// example:
	//
	// 30
	OriginalOrderRefundFee *float64 `json:"original_order_refund_fee,omitempty" xml:"original_order_refund_fee,omitempty"`
	// passenger for the refund
	Passenger *RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetailsPassenger `json:"passenger,omitempty" xml:"passenger,omitempty" type:"Struct"`
}

func (s RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetails) String() string {
	return dara.Prettify(s)
}

func (s RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetails) GoString() string {
	return s.String()
}

func (s *RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetails) GetChangeOrderRefundFee() *float64 {
	return s.ChangeOrderRefundFee
}

func (s *RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetails) GetOriginalOrderRefundFee() *float64 {
	return s.OriginalOrderRefundFee
}

func (s *RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetails) GetPassenger() *RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetailsPassenger {
	return s.Passenger
}

func (s *RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetails) SetChangeOrderRefundFee(v float64) *RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetails {
	s.ChangeOrderRefundFee = &v
	return s
}

func (s *RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetails) SetOriginalOrderRefundFee(v float64) *RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetails {
	s.OriginalOrderRefundFee = &v
	return s
}

func (s *RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetails) SetPassenger(v *RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetailsPassenger) *RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetails {
	s.Passenger = v
	return s
}

func (s *RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetails) Validate() error {
	if s.Passenger != nil {
		if err := s.Passenger.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetailsPassenger struct {
	// credential number
	//
	// example:
	//
	// 411***********4411
	Document *string `json:"document,omitempty" xml:"document,omitempty"`
	// first name
	//
	// example:
	//
	// SAN
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// last name
	//
	// example:
	//
	// ZHANG
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
}

func (s RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetailsPassenger) String() string {
	return dara.Prettify(s)
}

func (s RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetailsPassenger) GoString() string {
	return s.String()
}

func (s *RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetailsPassenger) GetDocument() *string {
	return s.Document
}

func (s *RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetailsPassenger) GetFirstName() *string {
	return s.FirstName
}

func (s *RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetailsPassenger) GetLastName() *string {
	return s.LastName
}

func (s *RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetailsPassenger) SetDocument(v string) *RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetailsPassenger {
	s.Document = &v
	return s
}

func (s *RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetailsPassenger) SetFirstName(v string) *RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetailsPassenger {
	s.FirstName = &v
	return s
}

func (s *RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetailsPassenger) SetLastName(v string) *RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetailsPassenger {
	s.LastName = &v
	return s
}

func (s *RefundDetailResponseBodyDataMultiRefundDetailsPassengerMultiRefundDetailsPassenger) Validate() error {
	return dara.Validate(s)
}

type RefundDetailResponseBodyDataPassengerRefundDetails struct {
	// information of the passenger applying for a refund
	Passenger *RefundDetailResponseBodyDataPassengerRefundDetailsPassenger `json:"passenger,omitempty" xml:"passenger,omitempty" type:"Struct"`
	// details of the refund fee
	RefundFee *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee `json:"refund_fee,omitempty" xml:"refund_fee,omitempty" type:"Struct"`
}

func (s RefundDetailResponseBodyDataPassengerRefundDetails) String() string {
	return dara.Prettify(s)
}

func (s RefundDetailResponseBodyDataPassengerRefundDetails) GoString() string {
	return s.String()
}

func (s *RefundDetailResponseBodyDataPassengerRefundDetails) GetPassenger() *RefundDetailResponseBodyDataPassengerRefundDetailsPassenger {
	return s.Passenger
}

func (s *RefundDetailResponseBodyDataPassengerRefundDetails) GetRefundFee() *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee {
	return s.RefundFee
}

func (s *RefundDetailResponseBodyDataPassengerRefundDetails) SetPassenger(v *RefundDetailResponseBodyDataPassengerRefundDetailsPassenger) *RefundDetailResponseBodyDataPassengerRefundDetails {
	s.Passenger = v
	return s
}

func (s *RefundDetailResponseBodyDataPassengerRefundDetails) SetRefundFee(v *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee) *RefundDetailResponseBodyDataPassengerRefundDetails {
	s.RefundFee = v
	return s
}

func (s *RefundDetailResponseBodyDataPassengerRefundDetails) Validate() error {
	if s.Passenger != nil {
		if err := s.Passenger.Validate(); err != nil {
			return err
		}
	}
	if s.RefundFee != nil {
		if err := s.RefundFee.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RefundDetailResponseBodyDataPassengerRefundDetailsPassenger struct {
	// credential number
	//
	// example:
	//
	// 411***********4411
	Document *string `json:"document,omitempty" xml:"document,omitempty"`
	// first name
	//
	// example:
	//
	// SAN
	FirstName *string `json:"first_name,omitempty" xml:"first_name,omitempty"`
	// first name
	//
	// example:
	//
	// ZHANG
	LastName *string `json:"last_name,omitempty" xml:"last_name,omitempty"`
}

func (s RefundDetailResponseBodyDataPassengerRefundDetailsPassenger) String() string {
	return dara.Prettify(s)
}

func (s RefundDetailResponseBodyDataPassengerRefundDetailsPassenger) GoString() string {
	return s.String()
}

func (s *RefundDetailResponseBodyDataPassengerRefundDetailsPassenger) GetDocument() *string {
	return s.Document
}

func (s *RefundDetailResponseBodyDataPassengerRefundDetailsPassenger) GetFirstName() *string {
	return s.FirstName
}

func (s *RefundDetailResponseBodyDataPassengerRefundDetailsPassenger) GetLastName() *string {
	return s.LastName
}

func (s *RefundDetailResponseBodyDataPassengerRefundDetailsPassenger) SetDocument(v string) *RefundDetailResponseBodyDataPassengerRefundDetailsPassenger {
	s.Document = &v
	return s
}

func (s *RefundDetailResponseBodyDataPassengerRefundDetailsPassenger) SetFirstName(v string) *RefundDetailResponseBodyDataPassengerRefundDetailsPassenger {
	s.FirstName = &v
	return s
}

func (s *RefundDetailResponseBodyDataPassengerRefundDetailsPassenger) SetLastName(v string) *RefundDetailResponseBodyDataPassengerRefundDetailsPassenger {
	s.LastName = &v
	return s
}

func (s *RefundDetailResponseBodyDataPassengerRefundDetailsPassenger) Validate() error {
	return dara.Validate(s)
}

type RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee struct {
	// total price of the used flight tickets
	//
	// example:
	//
	// 30
	AlreadyUsedTotalFee *float64 `json:"already_used_total_fee,omitempty" xml:"already_used_total_fee,omitempty"`
	// amount refunded to the user after change (only available when has change order)
	//
	// example:
	//
	// 30
	ModifyRefundToBuyerMoney *float64 `json:"modify_refund_to_buyer_money,omitempty" xml:"modify_refund_to_buyer_money,omitempty"`
	// non-refundable change service fee
	//
	// example:
	//
	// 30
	NonRefundableChangeServiceFee *float64 `json:"non_refundable_change_service_fee,omitempty" xml:"non_refundable_change_service_fee,omitempty"`
	// non-refundable upgrade fee
	//
	// example:
	//
	// 30
	NonRefundableChangeUpgradeFee *float64 `json:"non_refundable_change_upgrade_fee,omitempty" xml:"non_refundable_change_upgrade_fee,omitempty"`
	// non-refundable tax amount, i.e., tax refund fee
	//
	// example:
	//
	// 30
	NonRefundableTaxFee *float64 `json:"non_refundable_tax_fee,omitempty" xml:"non_refundable_tax_fee,omitempty"`
	// non-refundable ticket amount, i.e., ticket refund fee
	//
	// example:
	//
	// 30
	NonRefundableTicketFee *float64 `json:"non_refundable_ticket_fee,omitempty" xml:"non_refundable_ticket_fee,omitempty"`
	// amount refundable to the user from the original ticket (fare + tax - non_refundable_ticket_fee - non_refundable_tax_fee - already_used_total_fee - discount)
	//
	// example:
	//
	// 30
	RefundToBuyerMoney *float64 `json:"refund_to_buyer_money,omitempty" xml:"refund_to_buyer_money,omitempty"`
}

func (s RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee) String() string {
	return dara.Prettify(s)
}

func (s RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee) GoString() string {
	return s.String()
}

func (s *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee) GetAlreadyUsedTotalFee() *float64 {
	return s.AlreadyUsedTotalFee
}

func (s *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee) GetModifyRefundToBuyerMoney() *float64 {
	return s.ModifyRefundToBuyerMoney
}

func (s *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee) GetNonRefundableChangeServiceFee() *float64 {
	return s.NonRefundableChangeServiceFee
}

func (s *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee) GetNonRefundableChangeUpgradeFee() *float64 {
	return s.NonRefundableChangeUpgradeFee
}

func (s *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee) GetNonRefundableTaxFee() *float64 {
	return s.NonRefundableTaxFee
}

func (s *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee) GetNonRefundableTicketFee() *float64 {
	return s.NonRefundableTicketFee
}

func (s *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee) GetRefundToBuyerMoney() *float64 {
	return s.RefundToBuyerMoney
}

func (s *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee) SetAlreadyUsedTotalFee(v float64) *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee {
	s.AlreadyUsedTotalFee = &v
	return s
}

func (s *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee) SetModifyRefundToBuyerMoney(v float64) *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee {
	s.ModifyRefundToBuyerMoney = &v
	return s
}

func (s *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee) SetNonRefundableChangeServiceFee(v float64) *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee {
	s.NonRefundableChangeServiceFee = &v
	return s
}

func (s *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee) SetNonRefundableChangeUpgradeFee(v float64) *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee {
	s.NonRefundableChangeUpgradeFee = &v
	return s
}

func (s *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee) SetNonRefundableTaxFee(v float64) *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee {
	s.NonRefundableTaxFee = &v
	return s
}

func (s *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee) SetNonRefundableTicketFee(v float64) *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee {
	s.NonRefundableTicketFee = &v
	return s
}

func (s *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee) SetRefundToBuyerMoney(v float64) *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee {
	s.RefundToBuyerMoney = &v
	return s
}

func (s *RefundDetailResponseBodyDataPassengerRefundDetailsRefundFee) Validate() error {
	return dara.Validate(s)
}

type RefundDetailResponseBodyDataRefundJourneys struct {
	// segment list
	SegmentList []*RefundDetailResponseBodyDataRefundJourneysSegmentList `json:"segment_list,omitempty" xml:"segment_list,omitempty" type:"Repeated"`
	// number of transfer
	//
	// example:
	//
	// 0
	TransferCount *int32 `json:"transfer_count,omitempty" xml:"transfer_count,omitempty"`
}

func (s RefundDetailResponseBodyDataRefundJourneys) String() string {
	return dara.Prettify(s)
}

func (s RefundDetailResponseBodyDataRefundJourneys) GoString() string {
	return s.String()
}

func (s *RefundDetailResponseBodyDataRefundJourneys) GetSegmentList() []*RefundDetailResponseBodyDataRefundJourneysSegmentList {
	return s.SegmentList
}

func (s *RefundDetailResponseBodyDataRefundJourneys) GetTransferCount() *int32 {
	return s.TransferCount
}

func (s *RefundDetailResponseBodyDataRefundJourneys) SetSegmentList(v []*RefundDetailResponseBodyDataRefundJourneysSegmentList) *RefundDetailResponseBodyDataRefundJourneys {
	s.SegmentList = v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneys) SetTransferCount(v int32) *RefundDetailResponseBodyDataRefundJourneys {
	s.TransferCount = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneys) Validate() error {
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

type RefundDetailResponseBodyDataRefundJourneysSegmentList struct {
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
	// field deprecated
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
	// Y
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
	// marketing flight no. (eg: KA5809)
	//
	// example:
	//
	// HO1295
	MarketingFlightNo *string `json:"marketing_flight_no,omitempty" xml:"marketing_flight_no,omitempty"`
	// marketing flight no. (eg: 5809)
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
	// operating flight no. (eg: CX601)
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
	// stopover city list when stop_quantity > 0 , use “,” for seperation use
	//
	// example:
	//
	// SEL,HKG
	StopCityList *string `json:"stop_city_list,omitempty" xml:"stop_city_list,omitempty"`
	// number of stopover
	//
	// example:
	//
	// 0
	StopQuantity *int32 `json:"stop_quantity,omitempty" xml:"stop_quantity,omitempty"`
}

func (s RefundDetailResponseBodyDataRefundJourneysSegmentList) String() string {
	return dara.Prettify(s)
}

func (s RefundDetailResponseBodyDataRefundJourneysSegmentList) GoString() string {
	return s.String()
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) GetArrivalAirport() *string {
	return s.ArrivalAirport
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) GetArrivalCity() *string {
	return s.ArrivalCity
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) GetArrivalTerminal() *string {
	return s.ArrivalTerminal
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) GetArrivalTime() *string {
	return s.ArrivalTime
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) GetAvailability() *string {
	return s.Availability
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) GetCabin() *string {
	return s.Cabin
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) GetCabinClass() *string {
	return s.CabinClass
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) GetCodeShare() *bool {
	return s.CodeShare
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) GetDepartureAirport() *string {
	return s.DepartureAirport
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) GetDepartureCity() *string {
	return s.DepartureCity
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) GetDepartureTerminal() *string {
	return s.DepartureTerminal
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) GetDepartureTime() *string {
	return s.DepartureTime
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) GetEquipType() *string {
	return s.EquipType
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) GetFlightDuration() *int32 {
	return s.FlightDuration
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) GetMarketingAirline() *string {
	return s.MarketingAirline
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) GetMarketingFlightNo() *string {
	return s.MarketingFlightNo
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) GetMarketingFlightNoInt() *int32 {
	return s.MarketingFlightNoInt
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) GetOperatingAirline() *string {
	return s.OperatingAirline
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) GetOperatingFlightNo() *string {
	return s.OperatingFlightNo
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) GetSegmentId() *string {
	return s.SegmentId
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) GetStopCityList() *string {
	return s.StopCityList
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) GetStopQuantity() *int32 {
	return s.StopQuantity
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetArrivalAirport(v string) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.ArrivalAirport = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetArrivalCity(v string) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.ArrivalCity = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetArrivalTerminal(v string) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.ArrivalTerminal = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetArrivalTime(v string) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.ArrivalTime = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetAvailability(v string) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.Availability = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetCabin(v string) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.Cabin = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetCabinClass(v string) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.CabinClass = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetCodeShare(v bool) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.CodeShare = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetDepartureAirport(v string) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.DepartureAirport = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetDepartureCity(v string) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.DepartureCity = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetDepartureTerminal(v string) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.DepartureTerminal = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetDepartureTime(v string) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.DepartureTime = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetEquipType(v string) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.EquipType = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetFlightDuration(v int32) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.FlightDuration = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetMarketingAirline(v string) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.MarketingAirline = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetMarketingFlightNo(v string) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.MarketingFlightNo = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetMarketingFlightNoInt(v int32) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.MarketingFlightNoInt = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetOperatingAirline(v string) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.OperatingAirline = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetOperatingFlightNo(v string) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.OperatingFlightNo = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetSegmentId(v string) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.SegmentId = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetStopCityList(v string) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.StopCityList = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) SetStopQuantity(v int32) *RefundDetailResponseBodyDataRefundJourneysSegmentList {
	s.StopQuantity = &v
	return s
}

func (s *RefundDetailResponseBodyDataRefundJourneysSegmentList) Validate() error {
	return dara.Validate(s)
}
