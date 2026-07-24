// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOrderListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OrderListResponseBody
	GetRequestId() *string
	SetData(v *OrderListResponseBodyData) *OrderListResponseBody
	GetData() *OrderListResponseBodyData
	SetErrorCode(v string) *OrderListResponseBody
	GetErrorCode() *string
	SetErrorData(v interface{}) *OrderListResponseBody
	GetErrorData() interface{}
	SetErrorMsg(v string) *OrderListResponseBody
	GetErrorMsg() *string
	SetStatus(v int32) *OrderListResponseBody
	GetStatus() *int32
	SetSuccess(v bool) *OrderListResponseBody
	GetSuccess() *bool
}

type OrderListResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The data returned for a successful request.
	Data *OrderListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s OrderListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OrderListResponseBody) GoString() string {
	return s.String()
}

func (s *OrderListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OrderListResponseBody) GetData() *OrderListResponseBodyData {
	return s.Data
}

func (s *OrderListResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *OrderListResponseBody) GetErrorData() interface{} {
	return s.ErrorData
}

func (s *OrderListResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *OrderListResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *OrderListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OrderListResponseBody) SetRequestId(v string) *OrderListResponseBody {
	s.RequestId = &v
	return s
}

func (s *OrderListResponseBody) SetData(v *OrderListResponseBodyData) *OrderListResponseBody {
	s.Data = v
	return s
}

func (s *OrderListResponseBody) SetErrorCode(v string) *OrderListResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *OrderListResponseBody) SetErrorData(v interface{}) *OrderListResponseBody {
	s.ErrorData = v
	return s
}

func (s *OrderListResponseBody) SetErrorMsg(v string) *OrderListResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *OrderListResponseBody) SetStatus(v int32) *OrderListResponseBody {
	s.Status = &v
	return s
}

func (s *OrderListResponseBody) SetSuccess(v bool) *OrderListResponseBody {
	s.Success = &v
	return s
}

func (s *OrderListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OrderListResponseBodyData struct {
	// The data list.
	List []*OrderListResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// The pagination information.
	Pagination *OrderListResponseBodyDataPagination `json:"pagination,omitempty" xml:"pagination,omitempty" type:"Struct"`
}

func (s OrderListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s OrderListResponseBodyData) GoString() string {
	return s.String()
}

func (s *OrderListResponseBodyData) GetList() []*OrderListResponseBodyDataList {
	return s.List
}

func (s *OrderListResponseBodyData) GetPagination() *OrderListResponseBodyDataPagination {
	return s.Pagination
}

func (s *OrderListResponseBodyData) SetList(v []*OrderListResponseBodyDataList) *OrderListResponseBodyData {
	s.List = v
	return s
}

func (s *OrderListResponseBodyData) SetPagination(v *OrderListResponseBodyDataPagination) *OrderListResponseBodyData {
	s.Pagination = v
	return s
}

func (s *OrderListResponseBodyData) Validate() error {
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

type OrderListResponseBodyDataList struct {
	// The booking time (order creation time). The value is a 13-digit UNIX timestamp.
	//
	// example:
	//
	// 1677210784000
	BookTime *int64 `json:"book_time,omitempty" xml:"book_time,omitempty"`
	// The order number.
	//
	// example:
	//
	// 4966***617111
	OrderNum *int64 `json:"order_num,omitempty" xml:"order_num,omitempty"`
	// The order status. Valid values:
	//
	// - 2: order creation succeeded.
	//
	// - 3: order paid.
	//
	// - 4: order succeeded.
	//
	// - 5: order closed.
	//
	// example:
	//
	// 4
	OrderStatus *string `json:"order_status,omitempty" xml:"order_status,omitempty"`
	// The external order number.
	//
	// example:
	//
	// x091-2023-0220-j-0001
	OutOrderNum *string `json:"out_order_num,omitempty" xml:"out_order_num,omitempty"`
	// The passenger list.
	PassengerList []*OrderListResponseBodyDataListPassengerList `json:"passenger_list,omitempty" xml:"passenger_list,omitempty" type:"Repeated"`
	// The payment status. Valid values:
	//
	// - 0: initialized.
	//
	// - 1: creation succeeded.
	//
	// - 2: payment succeeded.
	//
	// - 4: transaction closed.
	//
	// example:
	//
	// 2
	PayStatus *string `json:"pay_status,omitempty" xml:"pay_status,omitempty"`
	// The payment time. The value is a 13-digit UNIX timestamp.
	//
	// example:
	//
	// 1677210788000
	PayTime *int64 `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	// The discount amount. Unit: yuan.
	//
	// example:
	//
	// 10
	PromotionPrice *float64 `json:"promotion_price,omitempty" xml:"promotion_price,omitempty"`
	// The actual payment amount. Unit: yuan.
	//
	// example:
	//
	// 3000
	RealPayPrice *float64 `json:"real_pay_price,omitempty" xml:"real_pay_price,omitempty"`
	// The buyer nickname.
	//
	// example:
	//
	// nick
	SessionNick *string `json:"session_nick,omitempty" xml:"session_nick,omitempty"`
	// The ticketing time. The value is a 13-digit UNIX timestamp.
	//
	// example:
	//
	// 1677210786000
	SucceedTime *int64 `json:"succeed_time,omitempty" xml:"succeed_time,omitempty"`
	// The total price of the order. Unit: yuan.
	//
	// example:
	//
	// 3000
	TotalPrice *float64 `json:"total_price,omitempty" xml:"total_price,omitempty"`
	// The transaction serial number.
	//
	// example:
	//
	// hkduendkd-2023-dj0
	TransactionNo *string `json:"transaction_no,omitempty" xml:"transaction_no,omitempty"`
}

func (s OrderListResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s OrderListResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *OrderListResponseBodyDataList) GetBookTime() *int64 {
	return s.BookTime
}

func (s *OrderListResponseBodyDataList) GetOrderNum() *int64 {
	return s.OrderNum
}

func (s *OrderListResponseBodyDataList) GetOrderStatus() *string {
	return s.OrderStatus
}

func (s *OrderListResponseBodyDataList) GetOutOrderNum() *string {
	return s.OutOrderNum
}

func (s *OrderListResponseBodyDataList) GetPassengerList() []*OrderListResponseBodyDataListPassengerList {
	return s.PassengerList
}

func (s *OrderListResponseBodyDataList) GetPayStatus() *string {
	return s.PayStatus
}

func (s *OrderListResponseBodyDataList) GetPayTime() *int64 {
	return s.PayTime
}

func (s *OrderListResponseBodyDataList) GetPromotionPrice() *float64 {
	return s.PromotionPrice
}

func (s *OrderListResponseBodyDataList) GetRealPayPrice() *float64 {
	return s.RealPayPrice
}

func (s *OrderListResponseBodyDataList) GetSessionNick() *string {
	return s.SessionNick
}

func (s *OrderListResponseBodyDataList) GetSucceedTime() *int64 {
	return s.SucceedTime
}

func (s *OrderListResponseBodyDataList) GetTotalPrice() *float64 {
	return s.TotalPrice
}

func (s *OrderListResponseBodyDataList) GetTransactionNo() *string {
	return s.TransactionNo
}

func (s *OrderListResponseBodyDataList) SetBookTime(v int64) *OrderListResponseBodyDataList {
	s.BookTime = &v
	return s
}

func (s *OrderListResponseBodyDataList) SetOrderNum(v int64) *OrderListResponseBodyDataList {
	s.OrderNum = &v
	return s
}

func (s *OrderListResponseBodyDataList) SetOrderStatus(v string) *OrderListResponseBodyDataList {
	s.OrderStatus = &v
	return s
}

func (s *OrderListResponseBodyDataList) SetOutOrderNum(v string) *OrderListResponseBodyDataList {
	s.OutOrderNum = &v
	return s
}

func (s *OrderListResponseBodyDataList) SetPassengerList(v []*OrderListResponseBodyDataListPassengerList) *OrderListResponseBodyDataList {
	s.PassengerList = v
	return s
}

func (s *OrderListResponseBodyDataList) SetPayStatus(v string) *OrderListResponseBodyDataList {
	s.PayStatus = &v
	return s
}

func (s *OrderListResponseBodyDataList) SetPayTime(v int64) *OrderListResponseBodyDataList {
	s.PayTime = &v
	return s
}

func (s *OrderListResponseBodyDataList) SetPromotionPrice(v float64) *OrderListResponseBodyDataList {
	s.PromotionPrice = &v
	return s
}

func (s *OrderListResponseBodyDataList) SetRealPayPrice(v float64) *OrderListResponseBodyDataList {
	s.RealPayPrice = &v
	return s
}

func (s *OrderListResponseBodyDataList) SetSessionNick(v string) *OrderListResponseBodyDataList {
	s.SessionNick = &v
	return s
}

func (s *OrderListResponseBodyDataList) SetSucceedTime(v int64) *OrderListResponseBodyDataList {
	s.SucceedTime = &v
	return s
}

func (s *OrderListResponseBodyDataList) SetTotalPrice(v float64) *OrderListResponseBodyDataList {
	s.TotalPrice = &v
	return s
}

func (s *OrderListResponseBodyDataList) SetTransactionNo(v string) *OrderListResponseBodyDataList {
	s.TransactionNo = &v
	return s
}

func (s *OrderListResponseBodyDataList) Validate() error {
	if s.PassengerList != nil {
		for _, item := range s.PassengerList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type OrderListResponseBodyDataListPassengerList struct {
	// The date of birth in the yyyyMMdd format.
	//
	// example:
	//
	// 20020301
	Birthday *string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// The credential information.
	Credential *OrderListResponseBodyDataListPassengerListCredential `json:"credential,omitempty" xml:"credential,omitempty" type:"Struct"`
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

func (s OrderListResponseBodyDataListPassengerList) String() string {
	return dara.Prettify(s)
}

func (s OrderListResponseBodyDataListPassengerList) GoString() string {
	return s.String()
}

func (s *OrderListResponseBodyDataListPassengerList) GetBirthday() *string {
	return s.Birthday
}

func (s *OrderListResponseBodyDataListPassengerList) GetCredential() *OrderListResponseBodyDataListPassengerListCredential {
	return s.Credential
}

func (s *OrderListResponseBodyDataListPassengerList) GetFirstName() *string {
	return s.FirstName
}

func (s *OrderListResponseBodyDataListPassengerList) GetGender() *int32 {
	return s.Gender
}

func (s *OrderListResponseBodyDataListPassengerList) GetLastName() *string {
	return s.LastName
}

func (s *OrderListResponseBodyDataListPassengerList) GetMobileCountryCode() *string {
	return s.MobileCountryCode
}

func (s *OrderListResponseBodyDataListPassengerList) GetMobilePhoneNumber() *string {
	return s.MobilePhoneNumber
}

func (s *OrderListResponseBodyDataListPassengerList) GetNationality() *string {
	return s.Nationality
}

func (s *OrderListResponseBodyDataListPassengerList) GetType() *int32 {
	return s.Type
}

func (s *OrderListResponseBodyDataListPassengerList) SetBirthday(v string) *OrderListResponseBodyDataListPassengerList {
	s.Birthday = &v
	return s
}

func (s *OrderListResponseBodyDataListPassengerList) SetCredential(v *OrderListResponseBodyDataListPassengerListCredential) *OrderListResponseBodyDataListPassengerList {
	s.Credential = v
	return s
}

func (s *OrderListResponseBodyDataListPassengerList) SetFirstName(v string) *OrderListResponseBodyDataListPassengerList {
	s.FirstName = &v
	return s
}

func (s *OrderListResponseBodyDataListPassengerList) SetGender(v int32) *OrderListResponseBodyDataListPassengerList {
	s.Gender = &v
	return s
}

func (s *OrderListResponseBodyDataListPassengerList) SetLastName(v string) *OrderListResponseBodyDataListPassengerList {
	s.LastName = &v
	return s
}

func (s *OrderListResponseBodyDataListPassengerList) SetMobileCountryCode(v string) *OrderListResponseBodyDataListPassengerList {
	s.MobileCountryCode = &v
	return s
}

func (s *OrderListResponseBodyDataListPassengerList) SetMobilePhoneNumber(v string) *OrderListResponseBodyDataListPassengerList {
	s.MobilePhoneNumber = &v
	return s
}

func (s *OrderListResponseBodyDataListPassengerList) SetNationality(v string) *OrderListResponseBodyDataListPassengerList {
	s.Nationality = &v
	return s
}

func (s *OrderListResponseBodyDataListPassengerList) SetType(v int32) *OrderListResponseBodyDataListPassengerList {
	s.Type = &v
	return s
}

func (s *OrderListResponseBodyDataListPassengerList) Validate() error {
	if s.Credential != nil {
		if err := s.Credential.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OrderListResponseBodyDataListPassengerListCredential struct {
	// The place of issue, represented as a two-letter code.
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
	// - 2: student ID.
	//
	// - 3: military ID.
	//
	// - 4: Home Return Permit.
	//
	// - 5: Taiwan Compatriot Permit.
	//
	// - 6: Hong Kong and Macao Travel Permit.
	//
	// - 7: international seafarer certificate.
	//
	// - 8: Foreigner Permanent Residence Card.
	//
	// - 10: police officer ID.
	//
	// - 11: soldier ID.
	//
	// - 12: Taiwan Travel Permit.
	//
	// - 13: Taiwan Entry Permit.
	//
	// - 14: household register.
	//
	// - 15: birth certificate.
	//
	// - 16: driver license.
	//
	// - 17: Hong Kong and Macao Resident Residence Permit.
	//
	// - 18: Taiwan Resident Residence Permit.
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

func (s OrderListResponseBodyDataListPassengerListCredential) String() string {
	return dara.Prettify(s)
}

func (s OrderListResponseBodyDataListPassengerListCredential) GoString() string {
	return s.String()
}

func (s *OrderListResponseBodyDataListPassengerListCredential) GetCertIssuePlace() *string {
	return s.CertIssuePlace
}

func (s *OrderListResponseBodyDataListPassengerListCredential) GetCredentialNum() *string {
	return s.CredentialNum
}

func (s *OrderListResponseBodyDataListPassengerListCredential) GetCredentialType() *int32 {
	return s.CredentialType
}

func (s *OrderListResponseBodyDataListPassengerListCredential) GetExpireDate() *string {
	return s.ExpireDate
}

func (s *OrderListResponseBodyDataListPassengerListCredential) SetCertIssuePlace(v string) *OrderListResponseBodyDataListPassengerListCredential {
	s.CertIssuePlace = &v
	return s
}

func (s *OrderListResponseBodyDataListPassengerListCredential) SetCredentialNum(v string) *OrderListResponseBodyDataListPassengerListCredential {
	s.CredentialNum = &v
	return s
}

func (s *OrderListResponseBodyDataListPassengerListCredential) SetCredentialType(v int32) *OrderListResponseBodyDataListPassengerListCredential {
	s.CredentialType = &v
	return s
}

func (s *OrderListResponseBodyDataListPassengerListCredential) SetExpireDate(v string) *OrderListResponseBodyDataListPassengerListCredential {
	s.ExpireDate = &v
	return s
}

func (s *OrderListResponseBodyDataListPassengerListCredential) Validate() error {
	return dara.Validate(s)
}

type OrderListResponseBodyDataPagination struct {
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

func (s OrderListResponseBodyDataPagination) String() string {
	return dara.Prettify(s)
}

func (s OrderListResponseBodyDataPagination) GoString() string {
	return s.String()
}

func (s *OrderListResponseBodyDataPagination) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *OrderListResponseBodyDataPagination) GetPageSize() *int32 {
	return s.PageSize
}

func (s *OrderListResponseBodyDataPagination) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *OrderListResponseBodyDataPagination) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *OrderListResponseBodyDataPagination) SetCurrentPage(v int32) *OrderListResponseBodyDataPagination {
	s.CurrentPage = &v
	return s
}

func (s *OrderListResponseBodyDataPagination) SetPageSize(v int32) *OrderListResponseBodyDataPagination {
	s.PageSize = &v
	return s
}

func (s *OrderListResponseBodyDataPagination) SetTotalCount(v int32) *OrderListResponseBodyDataPagination {
	s.TotalCount = &v
	return s
}

func (s *OrderListResponseBodyDataPagination) SetTotalPage(v int32) *OrderListResponseBodyDataPagination {
	s.TotalPage = &v
	return s
}

func (s *OrderListResponseBodyDataPagination) Validate() error {
	return dara.Validate(s)
}
