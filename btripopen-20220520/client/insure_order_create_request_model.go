// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsureOrderCreateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicant(v *InsureOrderCreateRequestApplicant) *InsureOrderCreateRequest
	GetApplicant() *InsureOrderCreateRequestApplicant
	SetBtripUserId(v string) *InsureOrderCreateRequest
	GetBtripUserId() *string
	SetBuyerName(v string) *InsureOrderCreateRequest
	GetBuyerName() *string
	SetInsPersonAndSegmentList(v []*InsureOrderCreateRequestInsPersonAndSegmentList) *InsureOrderCreateRequest
	GetInsPersonAndSegmentList() []*InsureOrderCreateRequestInsPersonAndSegmentList
	SetIsvName(v string) *InsureOrderCreateRequest
	GetIsvName() *string
	SetOutInsOrderId(v string) *InsureOrderCreateRequest
	GetOutInsOrderId() *string
	SetOutOrderId(v string) *InsureOrderCreateRequest
	GetOutOrderId() *string
	SetOutSubOrderId(v string) *InsureOrderCreateRequest
	GetOutSubOrderId() *string
	SetSupplierCode(v string) *InsureOrderCreateRequest
	GetSupplierCode() *string
}

type InsureOrderCreateRequest struct {
	// This parameter is required.
	Applicant *InsureOrderCreateRequestApplicant `json:"applicant,omitempty" xml:"applicant,omitempty" type:"Struct"`
	// example:
	//
	// 20202109390122
	BtripUserId *string `json:"btrip_user_id,omitempty" xml:"btrip_user_id,omitempty"`
	BuyerName   *string `json:"buyer_name,omitempty" xml:"buyer_name,omitempty"`
	// This parameter is required.
	InsPersonAndSegmentList []*InsureOrderCreateRequestInsPersonAndSegmentList `json:"ins_person_and_segment_list,omitempty" xml:"ins_person_and_segment_list,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// PostalSavingsBank
	IsvName *string `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
	// example:
	//
	// 100000320302020
	OutInsOrderId *string `json:"out_ins_order_id,omitempty" xml:"out_ins_order_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 202310101026030
	OutOrderId *string `json:"out_order_id,omitempty" xml:"out_order_id,omitempty"`
	// example:
	//
	// 1020030003332000
	OutSubOrderId *string `json:"out_sub_order_id,omitempty" xml:"out_sub_order_id,omitempty"`
	// example:
	//
	// fliggy
	SupplierCode *string `json:"supplier_code,omitempty" xml:"supplier_code,omitempty"`
}

func (s InsureOrderCreateRequest) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderCreateRequest) GoString() string {
	return s.String()
}

func (s *InsureOrderCreateRequest) GetApplicant() *InsureOrderCreateRequestApplicant {
	return s.Applicant
}

func (s *InsureOrderCreateRequest) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *InsureOrderCreateRequest) GetBuyerName() *string {
	return s.BuyerName
}

func (s *InsureOrderCreateRequest) GetInsPersonAndSegmentList() []*InsureOrderCreateRequestInsPersonAndSegmentList {
	return s.InsPersonAndSegmentList
}

func (s *InsureOrderCreateRequest) GetIsvName() *string {
	return s.IsvName
}

func (s *InsureOrderCreateRequest) GetOutInsOrderId() *string {
	return s.OutInsOrderId
}

func (s *InsureOrderCreateRequest) GetOutOrderId() *string {
	return s.OutOrderId
}

func (s *InsureOrderCreateRequest) GetOutSubOrderId() *string {
	return s.OutSubOrderId
}

func (s *InsureOrderCreateRequest) GetSupplierCode() *string {
	return s.SupplierCode
}

func (s *InsureOrderCreateRequest) SetApplicant(v *InsureOrderCreateRequestApplicant) *InsureOrderCreateRequest {
	s.Applicant = v
	return s
}

func (s *InsureOrderCreateRequest) SetBtripUserId(v string) *InsureOrderCreateRequest {
	s.BtripUserId = &v
	return s
}

func (s *InsureOrderCreateRequest) SetBuyerName(v string) *InsureOrderCreateRequest {
	s.BuyerName = &v
	return s
}

func (s *InsureOrderCreateRequest) SetInsPersonAndSegmentList(v []*InsureOrderCreateRequestInsPersonAndSegmentList) *InsureOrderCreateRequest {
	s.InsPersonAndSegmentList = v
	return s
}

func (s *InsureOrderCreateRequest) SetIsvName(v string) *InsureOrderCreateRequest {
	s.IsvName = &v
	return s
}

func (s *InsureOrderCreateRequest) SetOutInsOrderId(v string) *InsureOrderCreateRequest {
	s.OutInsOrderId = &v
	return s
}

func (s *InsureOrderCreateRequest) SetOutOrderId(v string) *InsureOrderCreateRequest {
	s.OutOrderId = &v
	return s
}

func (s *InsureOrderCreateRequest) SetOutSubOrderId(v string) *InsureOrderCreateRequest {
	s.OutSubOrderId = &v
	return s
}

func (s *InsureOrderCreateRequest) SetSupplierCode(v string) *InsureOrderCreateRequest {
	s.SupplierCode = &v
	return s
}

func (s *InsureOrderCreateRequest) Validate() error {
	return dara.Validate(s)
}

type InsureOrderCreateRequestApplicant struct {
	// example:
	//
	// 2000-11-01
	Birthday *string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// example:
	//
	// 20202109390122
	BtripUserId *string `json:"btrip_user_id,omitempty" xml:"btrip_user_id,omitempty"`
	CertName    *string `json:"cert_name,omitempty" xml:"cert_name,omitempty"`
	// example:
	//
	// 110102200011018872
	CertNo *string `json:"cert_no,omitempty" xml:"cert_no,omitempty"`
	// example:
	//
	// 100
	CertType *string `json:"cert_type,omitempty" xml:"cert_type,omitempty"`
	// example:
	//
	// F
	Gender *string `json:"gender,omitempty" xml:"gender,omitempty"`
	// example:
	//
	// 1000000
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
}

func (s InsureOrderCreateRequestApplicant) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderCreateRequestApplicant) GoString() string {
	return s.String()
}

func (s *InsureOrderCreateRequestApplicant) GetBirthday() *string {
	return s.Birthday
}

func (s *InsureOrderCreateRequestApplicant) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *InsureOrderCreateRequestApplicant) GetCertName() *string {
	return s.CertName
}

func (s *InsureOrderCreateRequestApplicant) GetCertNo() *string {
	return s.CertNo
}

func (s *InsureOrderCreateRequestApplicant) GetCertType() *string {
	return s.CertType
}

func (s *InsureOrderCreateRequestApplicant) GetGender() *string {
	return s.Gender
}

func (s *InsureOrderCreateRequestApplicant) GetPhone() *string {
	return s.Phone
}

func (s *InsureOrderCreateRequestApplicant) SetBirthday(v string) *InsureOrderCreateRequestApplicant {
	s.Birthday = &v
	return s
}

func (s *InsureOrderCreateRequestApplicant) SetBtripUserId(v string) *InsureOrderCreateRequestApplicant {
	s.BtripUserId = &v
	return s
}

func (s *InsureOrderCreateRequestApplicant) SetCertName(v string) *InsureOrderCreateRequestApplicant {
	s.CertName = &v
	return s
}

func (s *InsureOrderCreateRequestApplicant) SetCertNo(v string) *InsureOrderCreateRequestApplicant {
	s.CertNo = &v
	return s
}

func (s *InsureOrderCreateRequestApplicant) SetCertType(v string) *InsureOrderCreateRequestApplicant {
	s.CertType = &v
	return s
}

func (s *InsureOrderCreateRequestApplicant) SetGender(v string) *InsureOrderCreateRequestApplicant {
	s.Gender = &v
	return s
}

func (s *InsureOrderCreateRequestApplicant) SetPhone(v string) *InsureOrderCreateRequestApplicant {
	s.Phone = &v
	return s
}

func (s *InsureOrderCreateRequestApplicant) Validate() error {
	return dara.Validate(s)
}

type InsureOrderCreateRequestInsPersonAndSegmentList struct {
	InsureSegment *InsureOrderCreateRequestInsPersonAndSegmentListInsureSegment `json:"insure_segment,omitempty" xml:"insure_segment,omitempty" type:"Struct"`
	Insured       *InsureOrderCreateRequestInsPersonAndSegmentListInsured       `json:"insured,omitempty" xml:"insured,omitempty" type:"Struct"`
	// example:
	//
	// 12399992002002010
	OutSubInsOrderId *string `json:"out_sub_ins_order_id,omitempty" xml:"out_sub_ins_order_id,omitempty"`
}

func (s InsureOrderCreateRequestInsPersonAndSegmentList) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderCreateRequestInsPersonAndSegmentList) GoString() string {
	return s.String()
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentList) GetInsureSegment() *InsureOrderCreateRequestInsPersonAndSegmentListInsureSegment {
	return s.InsureSegment
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentList) GetInsured() *InsureOrderCreateRequestInsPersonAndSegmentListInsured {
	return s.Insured
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentList) GetOutSubInsOrderId() *string {
	return s.OutSubInsOrderId
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentList) SetInsureSegment(v *InsureOrderCreateRequestInsPersonAndSegmentListInsureSegment) *InsureOrderCreateRequestInsPersonAndSegmentList {
	s.InsureSegment = v
	return s
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentList) SetInsured(v *InsureOrderCreateRequestInsPersonAndSegmentListInsured) *InsureOrderCreateRequestInsPersonAndSegmentList {
	s.Insured = v
	return s
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentList) SetOutSubInsOrderId(v string) *InsureOrderCreateRequestInsPersonAndSegmentList {
	s.OutSubInsOrderId = &v
	return s
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentList) Validate() error {
	return dara.Validate(s)
}

type InsureOrderCreateRequestInsPersonAndSegmentListInsureSegment struct {
	// example:
	//
	// YNT
	ArrAirportCode *string `json:"arr_airport_code,omitempty" xml:"arr_airport_code,omitempty"`
	// example:
	//
	// CKG
	ArrCity *string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// example:
	//
	// HGH
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// example:
	//
	// 2023-10-31 13:10:00
	ArrTime *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	// example:
	//
	// HGH
	DepAirportCode *string `json:"dep_airport_code,omitempty" xml:"dep_airport_code,omitempty"`
	// example:
	//
	// HGH
	DepCity *string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	// example:
	//
	// CAN
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	// example:
	//
	// 2023-10-31 10:55:00
	DepTime *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	// example:
	//
	// ZH9891
	FlightNo *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
}

func (s InsureOrderCreateRequestInsPersonAndSegmentListInsureSegment) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderCreateRequestInsPersonAndSegmentListInsureSegment) GoString() string {
	return s.String()
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentListInsureSegment) GetArrAirportCode() *string {
	return s.ArrAirportCode
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentListInsureSegment) GetArrCity() *string {
	return s.ArrCity
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentListInsureSegment) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentListInsureSegment) GetArrTime() *string {
	return s.ArrTime
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentListInsureSegment) GetDepAirportCode() *string {
	return s.DepAirportCode
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentListInsureSegment) GetDepCity() *string {
	return s.DepCity
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentListInsureSegment) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentListInsureSegment) GetDepTime() *string {
	return s.DepTime
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentListInsureSegment) GetFlightNo() *string {
	return s.FlightNo
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentListInsureSegment) SetArrAirportCode(v string) *InsureOrderCreateRequestInsPersonAndSegmentListInsureSegment {
	s.ArrAirportCode = &v
	return s
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentListInsureSegment) SetArrCity(v string) *InsureOrderCreateRequestInsPersonAndSegmentListInsureSegment {
	s.ArrCity = &v
	return s
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentListInsureSegment) SetArrCityCode(v string) *InsureOrderCreateRequestInsPersonAndSegmentListInsureSegment {
	s.ArrCityCode = &v
	return s
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentListInsureSegment) SetArrTime(v string) *InsureOrderCreateRequestInsPersonAndSegmentListInsureSegment {
	s.ArrTime = &v
	return s
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentListInsureSegment) SetDepAirportCode(v string) *InsureOrderCreateRequestInsPersonAndSegmentListInsureSegment {
	s.DepAirportCode = &v
	return s
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentListInsureSegment) SetDepCity(v string) *InsureOrderCreateRequestInsPersonAndSegmentListInsureSegment {
	s.DepCity = &v
	return s
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentListInsureSegment) SetDepCityCode(v string) *InsureOrderCreateRequestInsPersonAndSegmentListInsureSegment {
	s.DepCityCode = &v
	return s
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentListInsureSegment) SetDepTime(v string) *InsureOrderCreateRequestInsPersonAndSegmentListInsureSegment {
	s.DepTime = &v
	return s
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentListInsureSegment) SetFlightNo(v string) *InsureOrderCreateRequestInsPersonAndSegmentListInsureSegment {
	s.FlightNo = &v
	return s
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentListInsureSegment) Validate() error {
	return dara.Validate(s)
}

type InsureOrderCreateRequestInsPersonAndSegmentListInsured struct {
	// example:
	//
	// 2000-11-01
	Birthday *string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// example:
	//
	// 20202109390122
	BtripUserId *string `json:"btrip_user_id,omitempty" xml:"btrip_user_id,omitempty"`
	CertName    *string `json:"cert_name,omitempty" xml:"cert_name,omitempty"`
	// example:
	//
	// 110102200011018872
	CertNo *string `json:"cert_no,omitempty" xml:"cert_no,omitempty"`
	// example:
	//
	// 100
	CertType *string `json:"cert_type,omitempty" xml:"cert_type,omitempty"`
	// example:
	//
	// F
	Gender *string `json:"gender,omitempty" xml:"gender,omitempty"`
	// example:
	//
	// 1000000
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
}

func (s InsureOrderCreateRequestInsPersonAndSegmentListInsured) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderCreateRequestInsPersonAndSegmentListInsured) GoString() string {
	return s.String()
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentListInsured) GetBirthday() *string {
	return s.Birthday
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentListInsured) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentListInsured) GetCertName() *string {
	return s.CertName
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentListInsured) GetCertNo() *string {
	return s.CertNo
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentListInsured) GetCertType() *string {
	return s.CertType
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentListInsured) GetGender() *string {
	return s.Gender
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentListInsured) GetPhone() *string {
	return s.Phone
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentListInsured) SetBirthday(v string) *InsureOrderCreateRequestInsPersonAndSegmentListInsured {
	s.Birthday = &v
	return s
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentListInsured) SetBtripUserId(v string) *InsureOrderCreateRequestInsPersonAndSegmentListInsured {
	s.BtripUserId = &v
	return s
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentListInsured) SetCertName(v string) *InsureOrderCreateRequestInsPersonAndSegmentListInsured {
	s.CertName = &v
	return s
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentListInsured) SetCertNo(v string) *InsureOrderCreateRequestInsPersonAndSegmentListInsured {
	s.CertNo = &v
	return s
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentListInsured) SetCertType(v string) *InsureOrderCreateRequestInsPersonAndSegmentListInsured {
	s.CertType = &v
	return s
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentListInsured) SetGender(v string) *InsureOrderCreateRequestInsPersonAndSegmentListInsured {
	s.Gender = &v
	return s
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentListInsured) SetPhone(v string) *InsureOrderCreateRequestInsPersonAndSegmentListInsured {
	s.Phone = &v
	return s
}

func (s *InsureOrderCreateRequestInsPersonAndSegmentListInsured) Validate() error {
	return dara.Validate(s)
}
