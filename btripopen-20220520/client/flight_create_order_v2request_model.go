// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightCreateOrderV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetAsyncCreateOrderKey(v string) *FlightCreateOrderV2Request
	GetAsyncCreateOrderKey() *string
	SetAsyncCreateOrderMode(v bool) *FlightCreateOrderV2Request
	GetAsyncCreateOrderMode() *bool
	SetBtripUserId(v string) *FlightCreateOrderV2Request
	GetBtripUserId() *string
	SetBuyerName(v string) *FlightCreateOrderV2Request
	GetBuyerName() *string
	SetContactInfo(v *FlightCreateOrderV2RequestContactInfo) *FlightCreateOrderV2Request
	GetContactInfo() *FlightCreateOrderV2RequestContactInfo
	SetIsvName(v string) *FlightCreateOrderV2Request
	GetIsvName() *string
	SetOtaItemId(v string) *FlightCreateOrderV2Request
	GetOtaItemId() *string
	SetOutOrderId(v string) *FlightCreateOrderV2Request
	GetOutOrderId() *string
	SetTotalPriceCent(v int64) *FlightCreateOrderV2Request
	GetTotalPriceCent() *int64
	SetTravelers(v []*FlightCreateOrderV2RequestTravelers) *FlightCreateOrderV2Request
	GetTravelers() []*FlightCreateOrderV2RequestTravelers
}

type FlightCreateOrderV2Request struct {
	AsyncCreateOrderKey *string `json:"async_create_order_key,omitempty" xml:"async_create_order_key,omitempty"`
	// example:
	//
	// false
	AsyncCreateOrderMode *bool   `json:"async_create_order_mode,omitempty" xml:"async_create_order_mode,omitempty"`
	BtripUserId          *string `json:"btrip_user_id,omitempty" xml:"btrip_user_id,omitempty"`
	BuyerName            *string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
	// This parameter is required.
	ContactInfo *FlightCreateOrderV2RequestContactInfo `json:"contact_info,omitempty" xml:"contact_info,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// cheshiapi
	IsvName *string `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7fb731deeb4510b86c17e8c8c25740_11
	OtaItemId *string `json:"ota_item_id,omitempty" xml:"ota_item_id,omitempty"`
	// This parameter is required.
	OutOrderId     *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	TotalPriceCent *int64  `json:"total_price_cent,omitempty" xml:"total_price_cent,omitempty"`
	// This parameter is required.
	Travelers []*FlightCreateOrderV2RequestTravelers `json:"travelers,omitempty" xml:"travelers,omitempty" type:"Repeated"`
}

func (s FlightCreateOrderV2Request) String() string {
	return dara.Prettify(s)
}

func (s FlightCreateOrderV2Request) GoString() string {
	return s.String()
}

func (s *FlightCreateOrderV2Request) GetAsyncCreateOrderKey() *string {
	return s.AsyncCreateOrderKey
}

func (s *FlightCreateOrderV2Request) GetAsyncCreateOrderMode() *bool {
	return s.AsyncCreateOrderMode
}

func (s *FlightCreateOrderV2Request) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *FlightCreateOrderV2Request) GetBuyerName() *string {
	return s.BuyerName
}

func (s *FlightCreateOrderV2Request) GetContactInfo() *FlightCreateOrderV2RequestContactInfo {
	return s.ContactInfo
}

func (s *FlightCreateOrderV2Request) GetIsvName() *string {
	return s.IsvName
}

func (s *FlightCreateOrderV2Request) GetOtaItemId() *string {
	return s.OtaItemId
}

func (s *FlightCreateOrderV2Request) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *FlightCreateOrderV2Request) GetTotalPriceCent() *int64 {
	return s.TotalPriceCent
}

func (s *FlightCreateOrderV2Request) GetTravelers() []*FlightCreateOrderV2RequestTravelers {
	return s.Travelers
}

func (s *FlightCreateOrderV2Request) SetAsyncCreateOrderKey(v string) *FlightCreateOrderV2Request {
	s.AsyncCreateOrderKey = &v
	return s
}

func (s *FlightCreateOrderV2Request) SetAsyncCreateOrderMode(v bool) *FlightCreateOrderV2Request {
	s.AsyncCreateOrderMode = &v
	return s
}

func (s *FlightCreateOrderV2Request) SetBtripUserId(v string) *FlightCreateOrderV2Request {
	s.BtripUserId = &v
	return s
}

func (s *FlightCreateOrderV2Request) SetBuyerName(v string) *FlightCreateOrderV2Request {
	s.BuyerName = &v
	return s
}

func (s *FlightCreateOrderV2Request) SetContactInfo(v *FlightCreateOrderV2RequestContactInfo) *FlightCreateOrderV2Request {
	s.ContactInfo = v
	return s
}

func (s *FlightCreateOrderV2Request) SetIsvName(v string) *FlightCreateOrderV2Request {
	s.IsvName = &v
	return s
}

func (s *FlightCreateOrderV2Request) SetOtaItemId(v string) *FlightCreateOrderV2Request {
	s.OtaItemId = &v
	return s
}

func (s *FlightCreateOrderV2Request) SetOutOrderId(v string) *FlightCreateOrderV2Request {
	s.OutOrderId = &v
	return s
}

func (s *FlightCreateOrderV2Request) SetTotalPriceCent(v int64) *FlightCreateOrderV2Request {
	s.TotalPriceCent = &v
	return s
}

func (s *FlightCreateOrderV2Request) SetTravelers(v []*FlightCreateOrderV2RequestTravelers) *FlightCreateOrderV2Request {
	s.Travelers = v
	return s
}

func (s *FlightCreateOrderV2Request) Validate() error {
	return dara.Validate(s)
}

type FlightCreateOrderV2RequestContactInfo struct {
	ContactEmail       *string `json:"contact_email,omitempty" xml:"contact_email,omitempty"`
	ContactName        *string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	ContactPhone       *string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
	SendMsgToPassenger *bool   `json:"send_msg_to_passenger,omitempty" xml:"send_msg_to_passenger,omitempty"`
}

func (s FlightCreateOrderV2RequestContactInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightCreateOrderV2RequestContactInfo) GoString() string {
	return s.String()
}

func (s *FlightCreateOrderV2RequestContactInfo) GetContactEmail() *string {
	return s.ContactEmail
}

func (s *FlightCreateOrderV2RequestContactInfo) GetContactName() *string {
	return s.ContactName
}

func (s *FlightCreateOrderV2RequestContactInfo) GetContactPhone() *string {
	return s.ContactPhone
}

func (s *FlightCreateOrderV2RequestContactInfo) GetSendMsgToPassenger() *bool {
	return s.SendMsgToPassenger
}

func (s *FlightCreateOrderV2RequestContactInfo) SetContactEmail(v string) *FlightCreateOrderV2RequestContactInfo {
	s.ContactEmail = &v
	return s
}

func (s *FlightCreateOrderV2RequestContactInfo) SetContactName(v string) *FlightCreateOrderV2RequestContactInfo {
	s.ContactName = &v
	return s
}

func (s *FlightCreateOrderV2RequestContactInfo) SetContactPhone(v string) *FlightCreateOrderV2RequestContactInfo {
	s.ContactPhone = &v
	return s
}

func (s *FlightCreateOrderV2RequestContactInfo) SetSendMsgToPassenger(v bool) *FlightCreateOrderV2RequestContactInfo {
	s.SendMsgToPassenger = &v
	return s
}

func (s *FlightCreateOrderV2RequestContactInfo) Validate() error {
	return dara.Validate(s)
}

type FlightCreateOrderV2RequestTravelers struct {
	// This parameter is required.
	//
	// example:
	//
	// 1991-01-22
	Birthday   *string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	CertNation *string `json:"cert_nation,omitempty" xml:"cert_nation,omitempty"`
	// This parameter is required.
	CertNo *string `json:"cert_no,omitempty" xml:"cert_no,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	CertType         *int32  `json:"cert_type,omitempty" xml:"cert_type,omitempty"`
	CertValidDate    *string `json:"cert_valid_date,omitempty" xml:"cert_valid_date,omitempty"`
	CostCenterName   *string `json:"cost_center_name,omitempty" xml:"cost_center_name,omitempty"`
	CostCenterNumber *string `json:"cost_center_number,omitempty" xml:"cost_center_number,omitempty"`
	DeptId           *string `json:"dept_id,omitempty" xml:"dept_id,omitempty"`
	DeptName         *string `json:"dept_name,omitempty" xml:"dept_name,omitempty"`
	// This parameter is required.
	Gender          *int32  `json:"gender,omitempty" xml:"gender,omitempty"`
	InvoiceTitle    *string `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	Nationality     *string `json:"nationality,omitempty" xml:"nationality,omitempty"`
	NationalityCode *string `json:"nationality_code,omitempty" xml:"nationality_code,omitempty"`
	// This parameter is required.
	PassengerName *string `json:"passenger_name,omitempty" xml:"passenger_name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	PassengerType *int32 `json:"passenger_type,omitempty" xml:"passenger_type,omitempty"`
	// This parameter is required.
	Phone        *string `json:"phone,omitempty" xml:"phone,omitempty"`
	ProjectCode  *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectTitle *string `json:"project_title,omitempty" xml:"project_title,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 18155711459129970552412
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserType *int32  `json:"user_type,omitempty" xml:"user_type,omitempty"`
}

func (s FlightCreateOrderV2RequestTravelers) String() string {
	return dara.Prettify(s)
}

func (s FlightCreateOrderV2RequestTravelers) GoString() string {
	return s.String()
}

func (s *FlightCreateOrderV2RequestTravelers) GetBirthday() *string {
	return s.Birthday
}

func (s *FlightCreateOrderV2RequestTravelers) GetCertNation() *string {
	return s.CertNation
}

func (s *FlightCreateOrderV2RequestTravelers) GetCertNo() *string {
	return s.CertNo
}

func (s *FlightCreateOrderV2RequestTravelers) GetCertType() *int32 {
	return s.CertType
}

func (s *FlightCreateOrderV2RequestTravelers) GetCertValidDate() *string {
	return s.CertValidDate
}

func (s *FlightCreateOrderV2RequestTravelers) GetCostCenterName() *string {
	return s.CostCenterName
}

func (s *FlightCreateOrderV2RequestTravelers) GetCostCenterNumber() *string {
	return s.CostCenterNumber
}

func (s *FlightCreateOrderV2RequestTravelers) GetDeptId() *string {
	return s.DeptId
}

func (s *FlightCreateOrderV2RequestTravelers) GetDeptName() *string {
	return s.DeptName
}

func (s *FlightCreateOrderV2RequestTravelers) GetGender() *int32 {
	return s.Gender
}

func (s *FlightCreateOrderV2RequestTravelers) GetInvoiceTitle() *string {
	return s.InvoiceTitle
}

func (s *FlightCreateOrderV2RequestTravelers) GetNationality() *string {
	return s.Nationality
}

func (s *FlightCreateOrderV2RequestTravelers) GetNationalityCode() *string {
	return s.NationalityCode
}

func (s *FlightCreateOrderV2RequestTravelers) GetPassengerName() *string {
	return s.PassengerName
}

func (s *FlightCreateOrderV2RequestTravelers) GetPassengerType() *int32 {
	return s.PassengerType
}

func (s *FlightCreateOrderV2RequestTravelers) GetPhone() *string {
	return s.Phone
}

func (s *FlightCreateOrderV2RequestTravelers) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *FlightCreateOrderV2RequestTravelers) GetProjectTitle() *string {
	return s.ProjectTitle
}

func (s *FlightCreateOrderV2RequestTravelers) GetUserId() *string {
	return s.UserId
}

func (s *FlightCreateOrderV2RequestTravelers) GetUserType() *int32 {
	return s.UserType
}

func (s *FlightCreateOrderV2RequestTravelers) SetBirthday(v string) *FlightCreateOrderV2RequestTravelers {
	s.Birthday = &v
	return s
}

func (s *FlightCreateOrderV2RequestTravelers) SetCertNation(v string) *FlightCreateOrderV2RequestTravelers {
	s.CertNation = &v
	return s
}

func (s *FlightCreateOrderV2RequestTravelers) SetCertNo(v string) *FlightCreateOrderV2RequestTravelers {
	s.CertNo = &v
	return s
}

func (s *FlightCreateOrderV2RequestTravelers) SetCertType(v int32) *FlightCreateOrderV2RequestTravelers {
	s.CertType = &v
	return s
}

func (s *FlightCreateOrderV2RequestTravelers) SetCertValidDate(v string) *FlightCreateOrderV2RequestTravelers {
	s.CertValidDate = &v
	return s
}

func (s *FlightCreateOrderV2RequestTravelers) SetCostCenterName(v string) *FlightCreateOrderV2RequestTravelers {
	s.CostCenterName = &v
	return s
}

func (s *FlightCreateOrderV2RequestTravelers) SetCostCenterNumber(v string) *FlightCreateOrderV2RequestTravelers {
	s.CostCenterNumber = &v
	return s
}

func (s *FlightCreateOrderV2RequestTravelers) SetDeptId(v string) *FlightCreateOrderV2RequestTravelers {
	s.DeptId = &v
	return s
}

func (s *FlightCreateOrderV2RequestTravelers) SetDeptName(v string) *FlightCreateOrderV2RequestTravelers {
	s.DeptName = &v
	return s
}

func (s *FlightCreateOrderV2RequestTravelers) SetGender(v int32) *FlightCreateOrderV2RequestTravelers {
	s.Gender = &v
	return s
}

func (s *FlightCreateOrderV2RequestTravelers) SetInvoiceTitle(v string) *FlightCreateOrderV2RequestTravelers {
	s.InvoiceTitle = &v
	return s
}

func (s *FlightCreateOrderV2RequestTravelers) SetNationality(v string) *FlightCreateOrderV2RequestTravelers {
	s.Nationality = &v
	return s
}

func (s *FlightCreateOrderV2RequestTravelers) SetNationalityCode(v string) *FlightCreateOrderV2RequestTravelers {
	s.NationalityCode = &v
	return s
}

func (s *FlightCreateOrderV2RequestTravelers) SetPassengerName(v string) *FlightCreateOrderV2RequestTravelers {
	s.PassengerName = &v
	return s
}

func (s *FlightCreateOrderV2RequestTravelers) SetPassengerType(v int32) *FlightCreateOrderV2RequestTravelers {
	s.PassengerType = &v
	return s
}

func (s *FlightCreateOrderV2RequestTravelers) SetPhone(v string) *FlightCreateOrderV2RequestTravelers {
	s.Phone = &v
	return s
}

func (s *FlightCreateOrderV2RequestTravelers) SetProjectCode(v string) *FlightCreateOrderV2RequestTravelers {
	s.ProjectCode = &v
	return s
}

func (s *FlightCreateOrderV2RequestTravelers) SetProjectTitle(v string) *FlightCreateOrderV2RequestTravelers {
	s.ProjectTitle = &v
	return s
}

func (s *FlightCreateOrderV2RequestTravelers) SetUserId(v string) *FlightCreateOrderV2RequestTravelers {
	s.UserId = &v
	return s
}

func (s *FlightCreateOrderV2RequestTravelers) SetUserType(v int32) *FlightCreateOrderV2RequestTravelers {
	s.UserType = &v
	return s
}

func (s *FlightCreateOrderV2RequestTravelers) Validate() error {
	return dara.Validate(s)
}
