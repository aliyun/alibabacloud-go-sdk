// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightInventoryPriceCheckRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBtripUserId(v string) *IntlFlightInventoryPriceCheckRequest
	GetBtripUserId() *string
	SetBuyerName(v string) *IntlFlightInventoryPriceCheckRequest
	GetBuyerName() *string
	SetIsvName(v string) *IntlFlightInventoryPriceCheckRequest
	GetIsvName() *string
	SetOrderPrice(v int64) *IntlFlightInventoryPriceCheckRequest
	GetOrderPrice() *int64
	SetOtaItemId(v string) *IntlFlightInventoryPriceCheckRequest
	GetOtaItemId() *string
	SetPassengerList(v []*IntlFlightInventoryPriceCheckRequestPassengerList) *IntlFlightInventoryPriceCheckRequest
	GetPassengerList() []*IntlFlightInventoryPriceCheckRequestPassengerList
}

type IntlFlightInventoryPriceCheckRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 10001
	BtripUserId *string `json:"btrip_user_id,omitempty" xml:"btrip_user_id,omitempty"`
	// example:
	//
	// ZHANG/SAN
	BuyerName *string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
	// example:
	//
	// ZJTD
	IsvName *string `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
	// example:
	//
	// 102000
	OrderPrice *int64 `json:"order_price,omitempty" xml:"order_price,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 22111acaf9ea47c09ed0db6abc45be2d_0
	OtaItemId *string `json:"ota_item_id,omitempty" xml:"ota_item_id,omitempty"`
	// This parameter is required.
	PassengerList []*IntlFlightInventoryPriceCheckRequestPassengerList `json:"passenger_list,omitempty" xml:"passenger_list,omitempty" type:"Repeated"`
}

func (s IntlFlightInventoryPriceCheckRequest) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightInventoryPriceCheckRequest) GoString() string {
	return s.String()
}

func (s *IntlFlightInventoryPriceCheckRequest) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *IntlFlightInventoryPriceCheckRequest) GetBuyerName() *string {
	return s.BuyerName
}

func (s *IntlFlightInventoryPriceCheckRequest) GetIsvName() *string {
	return s.IsvName
}

func (s *IntlFlightInventoryPriceCheckRequest) GetOrderPrice() *int64 {
	return s.OrderPrice
}

func (s *IntlFlightInventoryPriceCheckRequest) GetOtaItemId() *string {
	return s.OtaItemId
}

func (s *IntlFlightInventoryPriceCheckRequest) GetPassengerList() []*IntlFlightInventoryPriceCheckRequestPassengerList {
	return s.PassengerList
}

func (s *IntlFlightInventoryPriceCheckRequest) SetBtripUserId(v string) *IntlFlightInventoryPriceCheckRequest {
	s.BtripUserId = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckRequest) SetBuyerName(v string) *IntlFlightInventoryPriceCheckRequest {
	s.BuyerName = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckRequest) SetIsvName(v string) *IntlFlightInventoryPriceCheckRequest {
	s.IsvName = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckRequest) SetOrderPrice(v int64) *IntlFlightInventoryPriceCheckRequest {
	s.OrderPrice = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckRequest) SetOtaItemId(v string) *IntlFlightInventoryPriceCheckRequest {
	s.OtaItemId = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckRequest) SetPassengerList(v []*IntlFlightInventoryPriceCheckRequestPassengerList) *IntlFlightInventoryPriceCheckRequest {
	s.PassengerList = v
	return s
}

func (s *IntlFlightInventoryPriceCheckRequest) Validate() error {
	return dara.Validate(s)
}

type IntlFlightInventoryPriceCheckRequestPassengerList struct {
	// This parameter is required.
	//
	// example:
	//
	// 2020-01-01
	Birthday *string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// This parameter is required.
	CertInfo *IntlFlightInventoryPriceCheckRequestPassengerListCertInfo `json:"cert_info,omitempty" xml:"cert_info,omitempty" type:"Struct"`
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
	// example:
	//
	// CN
	NationalityCode *string `json:"nationality_code,omitempty" xml:"nationality_code,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 18012341234
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
	// 10001
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// example:
	//
	// 0
	UserType *int32 `json:"user_type,omitempty" xml:"user_type,omitempty"`
}

func (s IntlFlightInventoryPriceCheckRequestPassengerList) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightInventoryPriceCheckRequestPassengerList) GoString() string {
	return s.String()
}

func (s *IntlFlightInventoryPriceCheckRequestPassengerList) GetBirthday() *string {
	return s.Birthday
}

func (s *IntlFlightInventoryPriceCheckRequestPassengerList) GetCertInfo() *IntlFlightInventoryPriceCheckRequestPassengerListCertInfo {
	return s.CertInfo
}

func (s *IntlFlightInventoryPriceCheckRequestPassengerList) GetFullName() *string {
	return s.FullName
}

func (s *IntlFlightInventoryPriceCheckRequestPassengerList) GetGender() *int32 {
	return s.Gender
}

func (s *IntlFlightInventoryPriceCheckRequestPassengerList) GetJobNo() *string {
	return s.JobNo
}

func (s *IntlFlightInventoryPriceCheckRequestPassengerList) GetNationality() *string {
	return s.Nationality
}

func (s *IntlFlightInventoryPriceCheckRequestPassengerList) GetNationalityCode() *string {
	return s.NationalityCode
}

func (s *IntlFlightInventoryPriceCheckRequestPassengerList) GetPhone() *string {
	return s.Phone
}

func (s *IntlFlightInventoryPriceCheckRequestPassengerList) GetType() *int32 {
	return s.Type
}

func (s *IntlFlightInventoryPriceCheckRequestPassengerList) GetUserId() *string {
	return s.UserId
}

func (s *IntlFlightInventoryPriceCheckRequestPassengerList) GetUserType() *int32 {
	return s.UserType
}

func (s *IntlFlightInventoryPriceCheckRequestPassengerList) SetBirthday(v string) *IntlFlightInventoryPriceCheckRequestPassengerList {
	s.Birthday = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckRequestPassengerList) SetCertInfo(v *IntlFlightInventoryPriceCheckRequestPassengerListCertInfo) *IntlFlightInventoryPriceCheckRequestPassengerList {
	s.CertInfo = v
	return s
}

func (s *IntlFlightInventoryPriceCheckRequestPassengerList) SetFullName(v string) *IntlFlightInventoryPriceCheckRequestPassengerList {
	s.FullName = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckRequestPassengerList) SetGender(v int32) *IntlFlightInventoryPriceCheckRequestPassengerList {
	s.Gender = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckRequestPassengerList) SetJobNo(v string) *IntlFlightInventoryPriceCheckRequestPassengerList {
	s.JobNo = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckRequestPassengerList) SetNationality(v string) *IntlFlightInventoryPriceCheckRequestPassengerList {
	s.Nationality = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckRequestPassengerList) SetNationalityCode(v string) *IntlFlightInventoryPriceCheckRequestPassengerList {
	s.NationalityCode = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckRequestPassengerList) SetPhone(v string) *IntlFlightInventoryPriceCheckRequestPassengerList {
	s.Phone = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckRequestPassengerList) SetType(v int32) *IntlFlightInventoryPriceCheckRequestPassengerList {
	s.Type = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckRequestPassengerList) SetUserId(v string) *IntlFlightInventoryPriceCheckRequestPassengerList {
	s.UserId = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckRequestPassengerList) SetUserType(v int32) *IntlFlightInventoryPriceCheckRequestPassengerList {
	s.UserType = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckRequestPassengerList) Validate() error {
	return dara.Validate(s)
}

type IntlFlightInventoryPriceCheckRequestPassengerListCertInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// H123456
	CertNo *string `json:"cert_no,omitempty" xml:"cert_no,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CertType *int32 `json:"cert_type,omitempty" xml:"cert_type,omitempty"`
	// example:
	//
	// 2033-01-09
	CertValidDate *string `json:"cert_valid_date,omitempty" xml:"cert_valid_date,omitempty"`
	// example:
	//
	// 中国大陆
	IssuePlace *string `json:"issue_place,omitempty" xml:"issue_place,omitempty"`
}

func (s IntlFlightInventoryPriceCheckRequestPassengerListCertInfo) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightInventoryPriceCheckRequestPassengerListCertInfo) GoString() string {
	return s.String()
}

func (s *IntlFlightInventoryPriceCheckRequestPassengerListCertInfo) GetCertNo() *string {
	return s.CertNo
}

func (s *IntlFlightInventoryPriceCheckRequestPassengerListCertInfo) GetCertType() *int32 {
	return s.CertType
}

func (s *IntlFlightInventoryPriceCheckRequestPassengerListCertInfo) GetCertValidDate() *string {
	return s.CertValidDate
}

func (s *IntlFlightInventoryPriceCheckRequestPassengerListCertInfo) GetIssuePlace() *string {
	return s.IssuePlace
}

func (s *IntlFlightInventoryPriceCheckRequestPassengerListCertInfo) SetCertNo(v string) *IntlFlightInventoryPriceCheckRequestPassengerListCertInfo {
	s.CertNo = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckRequestPassengerListCertInfo) SetCertType(v int32) *IntlFlightInventoryPriceCheckRequestPassengerListCertInfo {
	s.CertType = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckRequestPassengerListCertInfo) SetCertValidDate(v string) *IntlFlightInventoryPriceCheckRequestPassengerListCertInfo {
	s.CertValidDate = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckRequestPassengerListCertInfo) SetIssuePlace(v string) *IntlFlightInventoryPriceCheckRequestPassengerListCertInfo {
	s.IssuePlace = &v
	return s
}

func (s *IntlFlightInventoryPriceCheckRequestPassengerListCertInfo) Validate() error {
	return dara.Validate(s)
}
