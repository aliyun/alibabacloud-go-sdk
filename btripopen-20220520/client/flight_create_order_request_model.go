// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightCreateOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArrAirportCode(v string) *FlightCreateOrderRequest
	GetArrAirportCode() *string
	SetArrCityCode(v string) *FlightCreateOrderRequest
	GetArrCityCode() *string
	SetAutoPay(v int32) *FlightCreateOrderRequest
	GetAutoPay() *int32
	SetBuyerName(v string) *FlightCreateOrderRequest
	GetBuyerName() *string
	SetBuyerUniqueKey(v string) *FlightCreateOrderRequest
	GetBuyerUniqueKey() *string
	SetContactInfo(v *FlightCreateOrderRequestContactInfo) *FlightCreateOrderRequest
	GetContactInfo() *FlightCreateOrderRequestContactInfo
	SetDepAirportCode(v string) *FlightCreateOrderRequest
	GetDepAirportCode() *string
	SetDepCityCode(v string) *FlightCreateOrderRequest
	GetDepCityCode() *string
	SetDepDate(v string) *FlightCreateOrderRequest
	GetDepDate() *string
	SetDisOrderId(v string) *FlightCreateOrderRequest
	GetDisOrderId() *string
	SetOrderAttr(v map[string]interface{}) *FlightCreateOrderRequest
	GetOrderAttr() map[string]interface{}
	SetOrderParams(v string) *FlightCreateOrderRequest
	GetOrderParams() *string
	SetOtaItemId(v string) *FlightCreateOrderRequest
	GetOtaItemId() *string
	SetPrice(v int64) *FlightCreateOrderRequest
	GetPrice() *int64
	SetReceiptAddress(v string) *FlightCreateOrderRequest
	GetReceiptAddress() *string
	SetReceiptTarget(v int32) *FlightCreateOrderRequest
	GetReceiptTarget() *int32
	SetReceiptTitle(v string) *FlightCreateOrderRequest
	GetReceiptTitle() *string
	SetTravelerInfoList(v []*FlightCreateOrderRequestTravelerInfoList) *FlightCreateOrderRequest
	GetTravelerInfoList() []*FlightCreateOrderRequestTravelerInfoList
	SetTripType(v int32) *FlightCreateOrderRequest
	GetTripType() *int32
}

type FlightCreateOrderRequest struct {
	ArrAirportCode *string `json:"arr_airport_code,omitempty" xml:"arr_airport_code,omitempty"`
	// This parameter is required.
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	AutoPay     *int32  `json:"auto_pay,omitempty" xml:"auto_pay,omitempty"`
	BuyerName   *string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
	// This parameter is required.
	BuyerUniqueKey *string `json:"buyer_unique_key,omitempty" xml:"buyer_unique_key,omitempty"`
	// This parameter is required.
	ContactInfo    *FlightCreateOrderRequestContactInfo `json:"contact_info,omitempty" xml:"contact_info,omitempty" type:"Struct"`
	DepAirportCode *string                              `json:"dep_airport_code,omitempty" xml:"dep_airport_code,omitempty"`
	// This parameter is required.
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// This parameter is required.
	DepDate *string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	// This parameter is required.
	DisOrderId *string                `json:"dis_order_id,omitempty" xml:"dis_order_id,omitempty"`
	OrderAttr  map[string]interface{} `json:"order_attr,omitempty" xml:"order_attr,omitempty"`
	// This parameter is required.
	OrderParams *string `json:"order_params,omitempty" xml:"order_params,omitempty"`
	// This parameter is required.
	OtaItemId *string `json:"ota_item_id,omitempty" xml:"ota_item_id,omitempty"`
	// This parameter is required.
	Price          *int64  `json:"price,omitempty" xml:"price,omitempty"`
	ReceiptAddress *string `json:"receipt_address,omitempty" xml:"receipt_address,omitempty"`
	ReceiptTarget  *int32  `json:"receipt_target,omitempty" xml:"receipt_target,omitempty"`
	ReceiptTitle   *string `json:"receipt_title,omitempty" xml:"receipt_title,omitempty"`
	// This parameter is required.
	TravelerInfoList []*FlightCreateOrderRequestTravelerInfoList `json:"traveler_info_list,omitempty" xml:"traveler_info_list,omitempty" type:"Repeated"`
	// This parameter is required.
	TripType *int32 `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
}

func (s FlightCreateOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s FlightCreateOrderRequest) GoString() string {
	return s.String()
}

func (s *FlightCreateOrderRequest) GetArrAirportCode() *string {
	return s.ArrAirportCode
}

func (s *FlightCreateOrderRequest) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *FlightCreateOrderRequest) GetAutoPay() *int32 {
	return s.AutoPay
}

func (s *FlightCreateOrderRequest) GetBuyerName() *string {
	return s.BuyerName
}

func (s *FlightCreateOrderRequest) GetBuyerUniqueKey() *string {
	return s.BuyerUniqueKey
}

func (s *FlightCreateOrderRequest) GetContactInfo() *FlightCreateOrderRequestContactInfo {
	return s.ContactInfo
}

func (s *FlightCreateOrderRequest) GetDepAirportCode() *string {
	return s.DepAirportCode
}

func (s *FlightCreateOrderRequest) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *FlightCreateOrderRequest) GetDepDate() *string {
	return s.DepDate
}

func (s *FlightCreateOrderRequest) GetDisOrderId() *string {
	return s.DisOrderId
}

func (s *FlightCreateOrderRequest) GetOrderAttr() map[string]interface{} {
	return s.OrderAttr
}

func (s *FlightCreateOrderRequest) GetOrderParams() *string {
	return s.OrderParams
}

func (s *FlightCreateOrderRequest) GetOtaItemId() *string {
	return s.OtaItemId
}

func (s *FlightCreateOrderRequest) GetPrice() *int64 {
	return s.Price
}

func (s *FlightCreateOrderRequest) GetReceiptAddress() *string {
	return s.ReceiptAddress
}

func (s *FlightCreateOrderRequest) GetReceiptTarget() *int32 {
	return s.ReceiptTarget
}

func (s *FlightCreateOrderRequest) GetReceiptTitle() *string {
	return s.ReceiptTitle
}

func (s *FlightCreateOrderRequest) GetTravelerInfoList() []*FlightCreateOrderRequestTravelerInfoList {
	return s.TravelerInfoList
}

func (s *FlightCreateOrderRequest) GetTripType() *int32 {
	return s.TripType
}

func (s *FlightCreateOrderRequest) SetArrAirportCode(v string) *FlightCreateOrderRequest {
	s.ArrAirportCode = &v
	return s
}

func (s *FlightCreateOrderRequest) SetArrCityCode(v string) *FlightCreateOrderRequest {
	s.ArrCityCode = &v
	return s
}

func (s *FlightCreateOrderRequest) SetAutoPay(v int32) *FlightCreateOrderRequest {
	s.AutoPay = &v
	return s
}

func (s *FlightCreateOrderRequest) SetBuyerName(v string) *FlightCreateOrderRequest {
	s.BuyerName = &v
	return s
}

func (s *FlightCreateOrderRequest) SetBuyerUniqueKey(v string) *FlightCreateOrderRequest {
	s.BuyerUniqueKey = &v
	return s
}

func (s *FlightCreateOrderRequest) SetContactInfo(v *FlightCreateOrderRequestContactInfo) *FlightCreateOrderRequest {
	s.ContactInfo = v
	return s
}

func (s *FlightCreateOrderRequest) SetDepAirportCode(v string) *FlightCreateOrderRequest {
	s.DepAirportCode = &v
	return s
}

func (s *FlightCreateOrderRequest) SetDepCityCode(v string) *FlightCreateOrderRequest {
	s.DepCityCode = &v
	return s
}

func (s *FlightCreateOrderRequest) SetDepDate(v string) *FlightCreateOrderRequest {
	s.DepDate = &v
	return s
}

func (s *FlightCreateOrderRequest) SetDisOrderId(v string) *FlightCreateOrderRequest {
	s.DisOrderId = &v
	return s
}

func (s *FlightCreateOrderRequest) SetOrderAttr(v map[string]interface{}) *FlightCreateOrderRequest {
	s.OrderAttr = v
	return s
}

func (s *FlightCreateOrderRequest) SetOrderParams(v string) *FlightCreateOrderRequest {
	s.OrderParams = &v
	return s
}

func (s *FlightCreateOrderRequest) SetOtaItemId(v string) *FlightCreateOrderRequest {
	s.OtaItemId = &v
	return s
}

func (s *FlightCreateOrderRequest) SetPrice(v int64) *FlightCreateOrderRequest {
	s.Price = &v
	return s
}

func (s *FlightCreateOrderRequest) SetReceiptAddress(v string) *FlightCreateOrderRequest {
	s.ReceiptAddress = &v
	return s
}

func (s *FlightCreateOrderRequest) SetReceiptTarget(v int32) *FlightCreateOrderRequest {
	s.ReceiptTarget = &v
	return s
}

func (s *FlightCreateOrderRequest) SetReceiptTitle(v string) *FlightCreateOrderRequest {
	s.ReceiptTitle = &v
	return s
}

func (s *FlightCreateOrderRequest) SetTravelerInfoList(v []*FlightCreateOrderRequestTravelerInfoList) *FlightCreateOrderRequest {
	s.TravelerInfoList = v
	return s
}

func (s *FlightCreateOrderRequest) SetTripType(v int32) *FlightCreateOrderRequest {
	s.TripType = &v
	return s
}

func (s *FlightCreateOrderRequest) Validate() error {
	if s.ContactInfo != nil {
		if err := s.ContactInfo.Validate(); err != nil {
			return err
		}
	}
	if s.TravelerInfoList != nil {
		for _, item := range s.TravelerInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type FlightCreateOrderRequestContactInfo struct {
	ContactEmail *string `json:"contact_email,omitempty" xml:"contact_email,omitempty"`
	// This parameter is required.
	ContactName *string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// This parameter is required.
	ContactPhone *string `json:"contact_phone,omitempty" xml:"contact_phone,omitempty"`
}

func (s FlightCreateOrderRequestContactInfo) String() string {
	return dara.Prettify(s)
}

func (s FlightCreateOrderRequestContactInfo) GoString() string {
	return s.String()
}

func (s *FlightCreateOrderRequestContactInfo) GetContactEmail() *string {
	return s.ContactEmail
}

func (s *FlightCreateOrderRequestContactInfo) GetContactName() *string {
	return s.ContactName
}

func (s *FlightCreateOrderRequestContactInfo) GetContactPhone() *string {
	return s.ContactPhone
}

func (s *FlightCreateOrderRequestContactInfo) SetContactEmail(v string) *FlightCreateOrderRequestContactInfo {
	s.ContactEmail = &v
	return s
}

func (s *FlightCreateOrderRequestContactInfo) SetContactName(v string) *FlightCreateOrderRequestContactInfo {
	s.ContactName = &v
	return s
}

func (s *FlightCreateOrderRequestContactInfo) SetContactPhone(v string) *FlightCreateOrderRequestContactInfo {
	s.ContactPhone = &v
	return s
}

func (s *FlightCreateOrderRequestContactInfo) Validate() error {
	return dara.Validate(s)
}

type FlightCreateOrderRequestTravelerInfoList struct {
	Birthday   *string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	CertNation *string `json:"cert_nation,omitempty" xml:"cert_nation,omitempty"`
	// This parameter is required.
	CertNo *string `json:"cert_no,omitempty" xml:"cert_no,omitempty"`
	// This parameter is required.
	CertType      *string `json:"cert_type,omitempty" xml:"cert_type,omitempty"`
	CertValidDate *string `json:"cert_valid_date,omitempty" xml:"cert_valid_date,omitempty"`
	// This parameter is required.
	Name            *string `json:"name,omitempty" xml:"name,omitempty"`
	Nationality     *string `json:"nationality,omitempty" xml:"nationality,omitempty"`
	NationalityCode *string `json:"nationality_code,omitempty" xml:"nationality_code,omitempty"`
	// This parameter is required.
	OutUserId *string `json:"out_user_id,omitempty" xml:"out_user_id,omitempty"`
	// This parameter is required.
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	Sex   *int32  `json:"sex,omitempty" xml:"sex,omitempty"`
	// This parameter is required.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FlightCreateOrderRequestTravelerInfoList) String() string {
	return dara.Prettify(s)
}

func (s FlightCreateOrderRequestTravelerInfoList) GoString() string {
	return s.String()
}

func (s *FlightCreateOrderRequestTravelerInfoList) GetBirthday() *string {
	return s.Birthday
}

func (s *FlightCreateOrderRequestTravelerInfoList) GetCertNation() *string {
	return s.CertNation
}

func (s *FlightCreateOrderRequestTravelerInfoList) GetCertNo() *string {
	return s.CertNo
}

func (s *FlightCreateOrderRequestTravelerInfoList) GetCertType() *string {
	return s.CertType
}

func (s *FlightCreateOrderRequestTravelerInfoList) GetCertValidDate() *string {
	return s.CertValidDate
}

func (s *FlightCreateOrderRequestTravelerInfoList) GetName() *string {
	return s.Name
}

func (s *FlightCreateOrderRequestTravelerInfoList) GetNationality() *string {
	return s.Nationality
}

func (s *FlightCreateOrderRequestTravelerInfoList) GetNationalityCode() *string {
	return s.NationalityCode
}

func (s *FlightCreateOrderRequestTravelerInfoList) GetOutUserId() *string {
	return s.OutUserId
}

func (s *FlightCreateOrderRequestTravelerInfoList) GetPhone() *string {
	return s.Phone
}

func (s *FlightCreateOrderRequestTravelerInfoList) GetSex() *int32 {
	return s.Sex
}

func (s *FlightCreateOrderRequestTravelerInfoList) GetType() *string {
	return s.Type
}

func (s *FlightCreateOrderRequestTravelerInfoList) SetBirthday(v string) *FlightCreateOrderRequestTravelerInfoList {
	s.Birthday = &v
	return s
}

func (s *FlightCreateOrderRequestTravelerInfoList) SetCertNation(v string) *FlightCreateOrderRequestTravelerInfoList {
	s.CertNation = &v
	return s
}

func (s *FlightCreateOrderRequestTravelerInfoList) SetCertNo(v string) *FlightCreateOrderRequestTravelerInfoList {
	s.CertNo = &v
	return s
}

func (s *FlightCreateOrderRequestTravelerInfoList) SetCertType(v string) *FlightCreateOrderRequestTravelerInfoList {
	s.CertType = &v
	return s
}

func (s *FlightCreateOrderRequestTravelerInfoList) SetCertValidDate(v string) *FlightCreateOrderRequestTravelerInfoList {
	s.CertValidDate = &v
	return s
}

func (s *FlightCreateOrderRequestTravelerInfoList) SetName(v string) *FlightCreateOrderRequestTravelerInfoList {
	s.Name = &v
	return s
}

func (s *FlightCreateOrderRequestTravelerInfoList) SetNationality(v string) *FlightCreateOrderRequestTravelerInfoList {
	s.Nationality = &v
	return s
}

func (s *FlightCreateOrderRequestTravelerInfoList) SetNationalityCode(v string) *FlightCreateOrderRequestTravelerInfoList {
	s.NationalityCode = &v
	return s
}

func (s *FlightCreateOrderRequestTravelerInfoList) SetOutUserId(v string) *FlightCreateOrderRequestTravelerInfoList {
	s.OutUserId = &v
	return s
}

func (s *FlightCreateOrderRequestTravelerInfoList) SetPhone(v string) *FlightCreateOrderRequestTravelerInfoList {
	s.Phone = &v
	return s
}

func (s *FlightCreateOrderRequestTravelerInfoList) SetSex(v int32) *FlightCreateOrderRequestTravelerInfoList {
	s.Sex = &v
	return s
}

func (s *FlightCreateOrderRequestTravelerInfoList) SetType(v string) *FlightCreateOrderRequestTravelerInfoList {
	s.Type = &v
	return s
}

func (s *FlightCreateOrderRequestTravelerInfoList) Validate() error {
	return dara.Validate(s)
}
