// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsureOrderDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InsureOrderDetailResponseBody
	GetCode() *string
	SetMessage(v string) *InsureOrderDetailResponseBody
	GetMessage() *string
	SetModule(v *InsureOrderDetailResponseBodyModule) *InsureOrderDetailResponseBody
	GetModule() *InsureOrderDetailResponseBodyModule
	SetRequestId(v string) *InsureOrderDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InsureOrderDetailResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *InsureOrderDetailResponseBody
	GetTraceId() *string
}

type InsureOrderDetailResponseBody struct {
	Code      *string                              `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                              `json:"message,omitempty" xml:"message,omitempty"`
	Module    *InsureOrderDetailResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                              `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                                `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                              `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s InsureOrderDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderDetailResponseBody) GoString() string {
	return s.String()
}

func (s *InsureOrderDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *InsureOrderDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InsureOrderDetailResponseBody) GetModule() *InsureOrderDetailResponseBodyModule {
	return s.Module
}

func (s *InsureOrderDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InsureOrderDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InsureOrderDetailResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *InsureOrderDetailResponseBody) SetCode(v string) *InsureOrderDetailResponseBody {
	s.Code = &v
	return s
}

func (s *InsureOrderDetailResponseBody) SetMessage(v string) *InsureOrderDetailResponseBody {
	s.Message = &v
	return s
}

func (s *InsureOrderDetailResponseBody) SetModule(v *InsureOrderDetailResponseBodyModule) *InsureOrderDetailResponseBody {
	s.Module = v
	return s
}

func (s *InsureOrderDetailResponseBody) SetRequestId(v string) *InsureOrderDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsureOrderDetailResponseBody) SetSuccess(v bool) *InsureOrderDetailResponseBody {
	s.Success = &v
	return s
}

func (s *InsureOrderDetailResponseBody) SetTraceId(v string) *InsureOrderDetailResponseBody {
	s.TraceId = &v
	return s
}

func (s *InsureOrderDetailResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InsureOrderDetailResponseBodyModule struct {
	Applicant             *InsureOrderDetailResponseBodyModuleApplicant               `json:"applicant,omitempty" xml:"applicant,omitempty" type:"Struct"`
	InsOrderId            *string                                                     `json:"ins_order_id,omitempty" xml:"ins_order_id,omitempty"`
	InsureOrderDetailList []*InsureOrderDetailResponseBodyModuleInsureOrderDetailList `json:"insure_order_detail_list,omitempty" xml:"insure_order_detail_list,omitempty" type:"Repeated"`
	Status                *string                                                     `json:"status,omitempty" xml:"status,omitempty"`
}

func (s InsureOrderDetailResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderDetailResponseBodyModule) GoString() string {
	return s.String()
}

func (s *InsureOrderDetailResponseBodyModule) GetApplicant() *InsureOrderDetailResponseBodyModuleApplicant {
	return s.Applicant
}

func (s *InsureOrderDetailResponseBodyModule) GetInsOrderId() *string {
	return s.InsOrderId
}

func (s *InsureOrderDetailResponseBodyModule) GetInsureOrderDetailList() []*InsureOrderDetailResponseBodyModuleInsureOrderDetailList {
	return s.InsureOrderDetailList
}

func (s *InsureOrderDetailResponseBodyModule) GetStatus() *string {
	return s.Status
}

func (s *InsureOrderDetailResponseBodyModule) SetApplicant(v *InsureOrderDetailResponseBodyModuleApplicant) *InsureOrderDetailResponseBodyModule {
	s.Applicant = v
	return s
}

func (s *InsureOrderDetailResponseBodyModule) SetInsOrderId(v string) *InsureOrderDetailResponseBodyModule {
	s.InsOrderId = &v
	return s
}

func (s *InsureOrderDetailResponseBodyModule) SetInsureOrderDetailList(v []*InsureOrderDetailResponseBodyModuleInsureOrderDetailList) *InsureOrderDetailResponseBodyModule {
	s.InsureOrderDetailList = v
	return s
}

func (s *InsureOrderDetailResponseBodyModule) SetStatus(v string) *InsureOrderDetailResponseBodyModule {
	s.Status = &v
	return s
}

func (s *InsureOrderDetailResponseBodyModule) Validate() error {
	if s.Applicant != nil {
		if err := s.Applicant.Validate(); err != nil {
			return err
		}
	}
	if s.InsureOrderDetailList != nil {
		for _, item := range s.InsureOrderDetailList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type InsureOrderDetailResponseBodyModuleApplicant struct {
	CertName *string `json:"cert_name,omitempty" xml:"cert_name,omitempty"`
	CertNo   *string `json:"cert_no,omitempty" xml:"cert_no,omitempty"`
	CertType *string `json:"cert_type,omitempty" xml:"cert_type,omitempty"`
	Phone    *string `json:"phone,omitempty" xml:"phone,omitempty"`
}

func (s InsureOrderDetailResponseBodyModuleApplicant) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderDetailResponseBodyModuleApplicant) GoString() string {
	return s.String()
}

func (s *InsureOrderDetailResponseBodyModuleApplicant) GetCertName() *string {
	return s.CertName
}

func (s *InsureOrderDetailResponseBodyModuleApplicant) GetCertNo() *string {
	return s.CertNo
}

func (s *InsureOrderDetailResponseBodyModuleApplicant) GetCertType() *string {
	return s.CertType
}

func (s *InsureOrderDetailResponseBodyModuleApplicant) GetPhone() *string {
	return s.Phone
}

func (s *InsureOrderDetailResponseBodyModuleApplicant) SetCertName(v string) *InsureOrderDetailResponseBodyModuleApplicant {
	s.CertName = &v
	return s
}

func (s *InsureOrderDetailResponseBodyModuleApplicant) SetCertNo(v string) *InsureOrderDetailResponseBodyModuleApplicant {
	s.CertNo = &v
	return s
}

func (s *InsureOrderDetailResponseBodyModuleApplicant) SetCertType(v string) *InsureOrderDetailResponseBodyModuleApplicant {
	s.CertType = &v
	return s
}

func (s *InsureOrderDetailResponseBodyModuleApplicant) SetPhone(v string) *InsureOrderDetailResponseBodyModuleApplicant {
	s.Phone = &v
	return s
}

func (s *InsureOrderDetailResponseBodyModuleApplicant) Validate() error {
	return dara.Validate(s)
}

type InsureOrderDetailResponseBodyModuleInsureOrderDetailList struct {
	EffectiveEndTime   *string                                                                `json:"effective_end_time,omitempty" xml:"effective_end_time,omitempty"`
	EffectiveStartTime *string                                                                `json:"effective_start_time,omitempty" xml:"effective_start_time,omitempty"`
	InsureSegment      *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsureSegment `json:"insure_segment,omitempty" xml:"insure_segment,omitempty" type:"Struct"`
	InsureTime         *string                                                                `json:"insure_time,omitempty" xml:"insure_time,omitempty"`
	Insured            *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsured       `json:"insured,omitempty" xml:"insured,omitempty" type:"Struct"`
	OutSubInsOrderId   *string                                                                `json:"out_sub_ins_order_id,omitempty" xml:"out_sub_ins_order_id,omitempty"`
	PolicyNo           *string                                                                `json:"policy_no,omitempty" xml:"policy_no,omitempty"`
	Price              *int64                                                                 `json:"price,omitempty" xml:"price,omitempty"`
	ProductName        *string                                                                `json:"product_name,omitempty" xml:"product_name,omitempty"`
	ProductNo          *string                                                                `json:"product_no,omitempty" xml:"product_no,omitempty"`
	Status             *string                                                                `json:"status,omitempty" xml:"status,omitempty"`
	SubInsOrderId      *string                                                                `json:"sub_ins_order_id,omitempty" xml:"sub_ins_order_id,omitempty"`
}

func (s InsureOrderDetailResponseBodyModuleInsureOrderDetailList) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderDetailResponseBodyModuleInsureOrderDetailList) GoString() string {
	return s.String()
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailList) GetEffectiveEndTime() *string {
	return s.EffectiveEndTime
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailList) GetEffectiveStartTime() *string {
	return s.EffectiveStartTime
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailList) GetInsureSegment() *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsureSegment {
	return s.InsureSegment
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailList) GetInsureTime() *string {
	return s.InsureTime
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailList) GetInsured() *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsured {
	return s.Insured
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailList) GetOutSubInsOrderId() *string {
	return s.OutSubInsOrderId
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailList) GetPolicyNo() *string {
	return s.PolicyNo
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailList) GetPrice() *int64 {
	return s.Price
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailList) GetProductName() *string {
	return s.ProductName
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailList) GetProductNo() *string {
	return s.ProductNo
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailList) GetStatus() *string {
	return s.Status
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailList) GetSubInsOrderId() *string {
	return s.SubInsOrderId
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailList) SetEffectiveEndTime(v string) *InsureOrderDetailResponseBodyModuleInsureOrderDetailList {
	s.EffectiveEndTime = &v
	return s
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailList) SetEffectiveStartTime(v string) *InsureOrderDetailResponseBodyModuleInsureOrderDetailList {
	s.EffectiveStartTime = &v
	return s
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailList) SetInsureSegment(v *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsureSegment) *InsureOrderDetailResponseBodyModuleInsureOrderDetailList {
	s.InsureSegment = v
	return s
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailList) SetInsureTime(v string) *InsureOrderDetailResponseBodyModuleInsureOrderDetailList {
	s.InsureTime = &v
	return s
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailList) SetInsured(v *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsured) *InsureOrderDetailResponseBodyModuleInsureOrderDetailList {
	s.Insured = v
	return s
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailList) SetOutSubInsOrderId(v string) *InsureOrderDetailResponseBodyModuleInsureOrderDetailList {
	s.OutSubInsOrderId = &v
	return s
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailList) SetPolicyNo(v string) *InsureOrderDetailResponseBodyModuleInsureOrderDetailList {
	s.PolicyNo = &v
	return s
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailList) SetPrice(v int64) *InsureOrderDetailResponseBodyModuleInsureOrderDetailList {
	s.Price = &v
	return s
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailList) SetProductName(v string) *InsureOrderDetailResponseBodyModuleInsureOrderDetailList {
	s.ProductName = &v
	return s
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailList) SetProductNo(v string) *InsureOrderDetailResponseBodyModuleInsureOrderDetailList {
	s.ProductNo = &v
	return s
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailList) SetStatus(v string) *InsureOrderDetailResponseBodyModuleInsureOrderDetailList {
	s.Status = &v
	return s
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailList) SetSubInsOrderId(v string) *InsureOrderDetailResponseBodyModuleInsureOrderDetailList {
	s.SubInsOrderId = &v
	return s
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailList) Validate() error {
	if s.InsureSegment != nil {
		if err := s.InsureSegment.Validate(); err != nil {
			return err
		}
	}
	if s.Insured != nil {
		if err := s.Insured.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsureSegment struct {
	ArrAirportCode *string `json:"arr_airport_code,omitempty" xml:"arr_airport_code,omitempty"`
	ArrCity        *string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	ArrCityCode    *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	ArrTime        *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	DepAirportCode *string `json:"dep_airport_code,omitempty" xml:"dep_airport_code,omitempty"`
	DepCity        *string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	DepCityCode    *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	DepTime        *string `json:"dep_time,omitempty" xml:"dep_time,omitempty"`
	FlightNo       *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
}

func (s InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsureSegment) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsureSegment) GoString() string {
	return s.String()
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsureSegment) GetArrAirportCode() *string {
	return s.ArrAirportCode
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsureSegment) GetArrCity() *string {
	return s.ArrCity
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsureSegment) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsureSegment) GetArrTime() *string {
	return s.ArrTime
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsureSegment) GetDepAirportCode() *string {
	return s.DepAirportCode
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsureSegment) GetDepCity() *string {
	return s.DepCity
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsureSegment) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsureSegment) GetDepTime() *string {
	return s.DepTime
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsureSegment) GetFlightNo() *string {
	return s.FlightNo
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsureSegment) SetArrAirportCode(v string) *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsureSegment {
	s.ArrAirportCode = &v
	return s
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsureSegment) SetArrCity(v string) *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsureSegment {
	s.ArrCity = &v
	return s
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsureSegment) SetArrCityCode(v string) *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsureSegment {
	s.ArrCityCode = &v
	return s
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsureSegment) SetArrTime(v string) *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsureSegment {
	s.ArrTime = &v
	return s
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsureSegment) SetDepAirportCode(v string) *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsureSegment {
	s.DepAirportCode = &v
	return s
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsureSegment) SetDepCity(v string) *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsureSegment {
	s.DepCity = &v
	return s
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsureSegment) SetDepCityCode(v string) *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsureSegment {
	s.DepCityCode = &v
	return s
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsureSegment) SetDepTime(v string) *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsureSegment {
	s.DepTime = &v
	return s
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsureSegment) SetFlightNo(v string) *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsureSegment {
	s.FlightNo = &v
	return s
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsureSegment) Validate() error {
	return dara.Validate(s)
}

type InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsured struct {
	Birthday    *string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	BtripUserId *string `json:"btrip_user_id,omitempty" xml:"btrip_user_id,omitempty"`
	CertName    *string `json:"cert_name,omitempty" xml:"cert_name,omitempty"`
	CertNo      *string `json:"cert_no,omitempty" xml:"cert_no,omitempty"`
	CertType    *string `json:"cert_type,omitempty" xml:"cert_type,omitempty"`
	Gender      *string `json:"gender,omitempty" xml:"gender,omitempty"`
	Phone       *string `json:"phone,omitempty" xml:"phone,omitempty"`
}

func (s InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsured) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsured) GoString() string {
	return s.String()
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsured) GetBirthday() *string {
	return s.Birthday
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsured) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsured) GetCertName() *string {
	return s.CertName
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsured) GetCertNo() *string {
	return s.CertNo
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsured) GetCertType() *string {
	return s.CertType
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsured) GetGender() *string {
	return s.Gender
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsured) GetPhone() *string {
	return s.Phone
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsured) SetBirthday(v string) *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsured {
	s.Birthday = &v
	return s
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsured) SetBtripUserId(v string) *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsured {
	s.BtripUserId = &v
	return s
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsured) SetCertName(v string) *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsured {
	s.CertName = &v
	return s
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsured) SetCertNo(v string) *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsured {
	s.CertNo = &v
	return s
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsured) SetCertType(v string) *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsured {
	s.CertType = &v
	return s
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsured) SetGender(v string) *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsured {
	s.Gender = &v
	return s
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsured) SetPhone(v string) *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsured {
	s.Phone = &v
	return s
}

func (s *InsureOrderDetailResponseBodyModuleInsureOrderDetailListInsured) Validate() error {
	return dara.Validate(s)
}
