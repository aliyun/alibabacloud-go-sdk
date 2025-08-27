// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightCreateOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncCreateOrderKey(v string) *IntlFlightCreateOrderRequest
	GetAsyncCreateOrderKey() *string
	SetAsyncCreateOrderMode(v bool) *IntlFlightCreateOrderRequest
	GetAsyncCreateOrderMode() *bool
	SetBtripUserId(v string) *IntlFlightCreateOrderRequest
	GetBtripUserId() *string
	SetBuyerName(v string) *IntlFlightCreateOrderRequest
	GetBuyerName() *string
	SetContactInfo(v *IntlFlightCreateOrderRequestContactInfo) *IntlFlightCreateOrderRequest
	GetContactInfo() *IntlFlightCreateOrderRequestContactInfo
	SetExtraInfo(v map[string]*string) *IntlFlightCreateOrderRequest
	GetExtraInfo() map[string]*string
	SetIsvName(v string) *IntlFlightCreateOrderRequest
	GetIsvName() *string
	SetOrderPrice(v int64) *IntlFlightCreateOrderRequest
	GetOrderPrice() *int64
	SetOtaItemId(v string) *IntlFlightCreateOrderRequest
	GetOtaItemId() *string
	SetOutOrderId(v string) *IntlFlightCreateOrderRequest
	GetOutOrderId() *string
	SetPassengerList(v []*IntlFlightCreateOrderRequestPassengerList) *IntlFlightCreateOrderRequest
	GetPassengerList() []*IntlFlightCreateOrderRequestPassengerList
	SetRenderKey(v string) *IntlFlightCreateOrderRequest
	GetRenderKey() *string
}

type IntlFlightCreateOrderRequest struct {
	AsyncCreateOrderKey  *string `json:"async_create_order_key,omitempty" xml:"async_create_order_key,omitempty"`
	AsyncCreateOrderMode *bool   `json:"async_create_order_mode,omitempty" xml:"async_create_order_mode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10001
	BtripUserId *string `json:"btrip_user_id,omitempty" xml:"btrip_user_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ZHANG/SAN
	BuyerName *string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
	// This parameter is required.
	ContactInfo *IntlFlightCreateOrderRequestContactInfo `json:"contact_info,omitempty" xml:"contact_info,omitempty" type:"Struct"`
	ExtraInfo   map[string]*string                       `json:"extra_info,omitempty" xml:"extra_info,omitempty"`
	IsvName     *string                                  `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
	OrderPrice  *int64                                   `json:"order_price,omitempty" xml:"order_price,omitempty"`
	// This parameter is required.
	OtaItemId  *string `json:"ota_item_id,omitempty" xml:"ota_item_id,omitempty"`
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// This parameter is required.
	PassengerList []*IntlFlightCreateOrderRequestPassengerList `json:"passenger_list,omitempty" xml:"passenger_list,omitempty" type:"Repeated"`
	RenderKey     *string                                      `json:"render_key,omitempty" xml:"render_key,omitempty"`
}

func (s IntlFlightCreateOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightCreateOrderRequest) GoString() string {
	return s.String()
}

func (s *IntlFlightCreateOrderRequest) GetAsyncCreateOrderKey() *string {
	return s.AsyncCreateOrderKey
}

func (s *IntlFlightCreateOrderRequest) GetAsyncCreateOrderMode() *bool {
	return s.AsyncCreateOrderMode
}

func (s *IntlFlightCreateOrderRequest) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *IntlFlightCreateOrderRequest) GetBuyerName() *string {
	return s.BuyerName
}

func (s *IntlFlightCreateOrderRequest) GetContactInfo() *IntlFlightCreateOrderRequestContactInfo {
	return s.ContactInfo
}

func (s *IntlFlightCreateOrderRequest) GetExtraInfo() map[string]*string {
	return s.ExtraInfo
}

func (s *IntlFlightCreateOrderRequest) GetIsvName() *string {
	return s.IsvName
}

func (s *IntlFlightCreateOrderRequest) GetOrderPrice() *int64 {
	return s.OrderPrice
}

func (s *IntlFlightCreateOrderRequest) GetOtaItemId() *string {
	return s.OtaItemId
}

func (s *IntlFlightCreateOrderRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *IntlFlightCreateOrderRequest) GetPassengerList() []*IntlFlightCreateOrderRequestPassengerList {
	return s.PassengerList
}

func (s *IntlFlightCreateOrderRequest) GetRenderKey() *string {
	return s.RenderKey
}

func (s *IntlFlightCreateOrderRequest) SetAsyncCreateOrderKey(v string) *IntlFlightCreateOrderRequest {
	s.AsyncCreateOrderKey = &v
	return s
}

func (s *IntlFlightCreateOrderRequest) SetAsyncCreateOrderMode(v bool) *IntlFlightCreateOrderRequest {
	s.AsyncCreateOrderMode = &v
	return s
}

func (s *IntlFlightCreateOrderRequest) SetBtripUserId(v string) *IntlFlightCreateOrderRequest {
	s.BtripUserId = &v
	return s
}

func (s *IntlFlightCreateOrderRequest) SetBuyerName(v string) *IntlFlightCreateOrderRequest {
	s.BuyerName = &v
	return s
}

func (s *IntlFlightCreateOrderRequest) SetContactInfo(v *IntlFlightCreateOrderRequestContactInfo) *IntlFlightCreateOrderRequest {
	s.ContactInfo = v
	return s
}

func (s *IntlFlightCreateOrderRequest) SetExtraInfo(v map[string]*string) *IntlFlightCreateOrderRequest {
	s.ExtraInfo = v
	return s
}

func (s *IntlFlightCreateOrderRequest) SetIsvName(v string) *IntlFlightCreateOrderRequest {
	s.IsvName = &v
	return s
}

func (s *IntlFlightCreateOrderRequest) SetOrderPrice(v int64) *IntlFlightCreateOrderRequest {
	s.OrderPrice = &v
	return s
}

func (s *IntlFlightCreateOrderRequest) SetOtaItemId(v string) *IntlFlightCreateOrderRequest {
	s.OtaItemId = &v
	return s
}

func (s *IntlFlightCreateOrderRequest) SetOutOrderId(v string) *IntlFlightCreateOrderRequest {
	s.OutOrderId = &v
	return s
}

func (s *IntlFlightCreateOrderRequest) SetPassengerList(v []*IntlFlightCreateOrderRequestPassengerList) *IntlFlightCreateOrderRequest {
	s.PassengerList = v
	return s
}

func (s *IntlFlightCreateOrderRequest) SetRenderKey(v string) *IntlFlightCreateOrderRequest {
	s.RenderKey = &v
	return s
}

func (s *IntlFlightCreateOrderRequest) Validate() error {
	return dara.Validate(s)
}

type IntlFlightCreateOrderRequestContactInfo struct {
	// This parameter is required.
	ContactEmail *string `json:"contact_email,omitempty" xml:"contact_email,omitempty"`
	// This parameter is required.
	ContactName *string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// This parameter is required.
	ContactPhone *string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
}

func (s IntlFlightCreateOrderRequestContactInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightCreateOrderRequestContactInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightCreateOrderRequestContactInfo) GetContactEmail() *string {
	return s.ContactEmail
}

func (s *IntlFlightCreateOrderRequestContactInfo) GetContactName() *string {
	return s.ContactName
}

func (s *IntlFlightCreateOrderRequestContactInfo) GetContactPhone() *string {
	return s.ContactPhone
}

func (s *IntlFlightCreateOrderRequestContactInfo) SetContactEmail(v string) *IntlFlightCreateOrderRequestContactInfo {
	s.ContactEmail = &v
	return s
}

func (s *IntlFlightCreateOrderRequestContactInfo) SetContactName(v string) *IntlFlightCreateOrderRequestContactInfo {
	s.ContactName = &v
	return s
}

func (s *IntlFlightCreateOrderRequestContactInfo) SetContactPhone(v string) *IntlFlightCreateOrderRequestContactInfo {
	s.ContactPhone = &v
	return s
}

func (s *IntlFlightCreateOrderRequestContactInfo) Validate() error {
	return dara.Validate(s)
}

type IntlFlightCreateOrderRequestPassengerList struct {
	// This parameter is required.
	//
	// example:
	//
	// 1998-12-28
	Birthday *string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// This parameter is required.
	CertInfo *IntlFlightCreateOrderRequestPassengerListCertInfo `json:"cert_info,omitempty" xml:"cert_info,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// ZHANG/SAN
	FullName *string `json:"full_name,omitempty" xml:"full_name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Gender *int32 `json:"gender,omitempty" xml:"gender,omitempty"`
	// example:
	//
	// L5000924
	JobNo *string `json:"job_no,omitempty" xml:"job_no,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 中国大陆
	Nationality *string `json:"nationality,omitempty" xml:"nationality,omitempty"`
	// This parameter is required.
	NationalityCode *string `json:"nationality_code,omitempty" xml:"nationality_code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 13100008888
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	Type *int32 `json:"type,omitempty" xml:"type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12292812036903456
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// example:
	//
	// 0
	UserType *int32 `json:"user_type,omitempty" xml:"user_type,omitempty"`
}

func (s IntlFlightCreateOrderRequestPassengerList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightCreateOrderRequestPassengerList) GoString() string {
	return s.String()
}

func (s *IntlFlightCreateOrderRequestPassengerList) GetBirthday() *string {
	return s.Birthday
}

func (s *IntlFlightCreateOrderRequestPassengerList) GetCertInfo() *IntlFlightCreateOrderRequestPassengerListCertInfo {
	return s.CertInfo
}

func (s *IntlFlightCreateOrderRequestPassengerList) GetFullName() *string {
	return s.FullName
}

func (s *IntlFlightCreateOrderRequestPassengerList) GetGender() *int32 {
	return s.Gender
}

func (s *IntlFlightCreateOrderRequestPassengerList) GetJobNo() *string {
	return s.JobNo
}

func (s *IntlFlightCreateOrderRequestPassengerList) GetNationality() *string {
	return s.Nationality
}

func (s *IntlFlightCreateOrderRequestPassengerList) GetNationalityCode() *string {
	return s.NationalityCode
}

func (s *IntlFlightCreateOrderRequestPassengerList) GetPhone() *string {
	return s.Phone
}

func (s *IntlFlightCreateOrderRequestPassengerList) GetType() *int32 {
	return s.Type
}

func (s *IntlFlightCreateOrderRequestPassengerList) GetUserId() *string {
	return s.UserId
}

func (s *IntlFlightCreateOrderRequestPassengerList) GetUserType() *int32 {
	return s.UserType
}

func (s *IntlFlightCreateOrderRequestPassengerList) SetBirthday(v string) *IntlFlightCreateOrderRequestPassengerList {
	s.Birthday = &v
	return s
}

func (s *IntlFlightCreateOrderRequestPassengerList) SetCertInfo(v *IntlFlightCreateOrderRequestPassengerListCertInfo) *IntlFlightCreateOrderRequestPassengerList {
	s.CertInfo = v
	return s
}

func (s *IntlFlightCreateOrderRequestPassengerList) SetFullName(v string) *IntlFlightCreateOrderRequestPassengerList {
	s.FullName = &v
	return s
}

func (s *IntlFlightCreateOrderRequestPassengerList) SetGender(v int32) *IntlFlightCreateOrderRequestPassengerList {
	s.Gender = &v
	return s
}

func (s *IntlFlightCreateOrderRequestPassengerList) SetJobNo(v string) *IntlFlightCreateOrderRequestPassengerList {
	s.JobNo = &v
	return s
}

func (s *IntlFlightCreateOrderRequestPassengerList) SetNationality(v string) *IntlFlightCreateOrderRequestPassengerList {
	s.Nationality = &v
	return s
}

func (s *IntlFlightCreateOrderRequestPassengerList) SetNationalityCode(v string) *IntlFlightCreateOrderRequestPassengerList {
	s.NationalityCode = &v
	return s
}

func (s *IntlFlightCreateOrderRequestPassengerList) SetPhone(v string) *IntlFlightCreateOrderRequestPassengerList {
	s.Phone = &v
	return s
}

func (s *IntlFlightCreateOrderRequestPassengerList) SetType(v int32) *IntlFlightCreateOrderRequestPassengerList {
	s.Type = &v
	return s
}

func (s *IntlFlightCreateOrderRequestPassengerList) SetUserId(v string) *IntlFlightCreateOrderRequestPassengerList {
	s.UserId = &v
	return s
}

func (s *IntlFlightCreateOrderRequestPassengerList) SetUserType(v int32) *IntlFlightCreateOrderRequestPassengerList {
	s.UserType = &v
	return s
}

func (s *IntlFlightCreateOrderRequestPassengerList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightCreateOrderRequestPassengerListCertInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// E1234567
	CertNo *string `json:"cert_no,omitempty" xml:"cert_no,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CertType *int32 `json:"cert_type,omitempty" xml:"cert_type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2029-12-31
	CertValidDate *string `json:"cert_valid_date,omitempty" xml:"cert_valid_date,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 中国大陆
	IssuePlace *string `json:"issue_place,omitempty" xml:"issue_place,omitempty"`
}

func (s IntlFlightCreateOrderRequestPassengerListCertInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightCreateOrderRequestPassengerListCertInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightCreateOrderRequestPassengerListCertInfo) GetCertNo() *string {
	return s.CertNo
}

func (s *IntlFlightCreateOrderRequestPassengerListCertInfo) GetCertType() *int32 {
	return s.CertType
}

func (s *IntlFlightCreateOrderRequestPassengerListCertInfo) GetCertValidDate() *string {
	return s.CertValidDate
}

func (s *IntlFlightCreateOrderRequestPassengerListCertInfo) GetIssuePlace() *string {
	return s.IssuePlace
}

func (s *IntlFlightCreateOrderRequestPassengerListCertInfo) SetCertNo(v string) *IntlFlightCreateOrderRequestPassengerListCertInfo {
	s.CertNo = &v
	return s
}

func (s *IntlFlightCreateOrderRequestPassengerListCertInfo) SetCertType(v int32) *IntlFlightCreateOrderRequestPassengerListCertInfo {
	s.CertType = &v
	return s
}

func (s *IntlFlightCreateOrderRequestPassengerListCertInfo) SetCertValidDate(v string) *IntlFlightCreateOrderRequestPassengerListCertInfo {
	s.CertValidDate = &v
	return s
}

func (s *IntlFlightCreateOrderRequestPassengerListCertInfo) SetIssuePlace(v string) *IntlFlightCreateOrderRequestPassengerListCertInfo {
	s.IssuePlace = &v
	return s
}

func (s *IntlFlightCreateOrderRequestPassengerListCertInfo) Validate() error {
	return dara.Validate(s)
}
