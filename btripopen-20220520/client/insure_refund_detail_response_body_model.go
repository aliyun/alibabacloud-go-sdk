// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsureRefundDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InsureRefundDetailResponseBody
	GetCode() *string
	SetMessage(v string) *InsureRefundDetailResponseBody
	GetMessage() *string
	SetModule(v *InsureRefundDetailResponseBodyModule) *InsureRefundDetailResponseBody
	GetModule() *InsureRefundDetailResponseBodyModule
	SetRequestId(v string) *InsureRefundDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InsureRefundDetailResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *InsureRefundDetailResponseBody
	GetTraceId() *string
}

type InsureRefundDetailResponseBody struct {
	Code      *string                               `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string                               `json:"message,omitempty" xml:"message,omitempty"`
	Module    *InsureRefundDetailResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	RequestId *string                               `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool                                 `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string                               `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s InsureRefundDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InsureRefundDetailResponseBody) GoString() string {
	return s.String()
}

func (s *InsureRefundDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *InsureRefundDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InsureRefundDetailResponseBody) GetModule() *InsureRefundDetailResponseBodyModule {
	return s.Module
}

func (s *InsureRefundDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InsureRefundDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InsureRefundDetailResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *InsureRefundDetailResponseBody) SetCode(v string) *InsureRefundDetailResponseBody {
	s.Code = &v
	return s
}

func (s *InsureRefundDetailResponseBody) SetMessage(v string) *InsureRefundDetailResponseBody {
	s.Message = &v
	return s
}

func (s *InsureRefundDetailResponseBody) SetModule(v *InsureRefundDetailResponseBodyModule) *InsureRefundDetailResponseBody {
	s.Module = v
	return s
}

func (s *InsureRefundDetailResponseBody) SetRequestId(v string) *InsureRefundDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsureRefundDetailResponseBody) SetSuccess(v bool) *InsureRefundDetailResponseBody {
	s.Success = &v
	return s
}

func (s *InsureRefundDetailResponseBody) SetTraceId(v string) *InsureRefundDetailResponseBody {
	s.TraceId = &v
	return s
}

func (s *InsureRefundDetailResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InsureRefundDetailResponseBodyModule struct {
	ApplyId            *string                                                   `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	GmtCreate          *string                                                   `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	GmtModified        *string                                                   `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	InsOrderId         *string                                                   `json:"ins_order_id,omitempty" xml:"ins_order_id,omitempty"`
	InsureOrder        *InsureRefundDetailResponseBodyModuleInsureOrder          `json:"insure_order,omitempty" xml:"insure_order,omitempty" type:"Struct"`
	OutApplyId         *string                                                   `json:"out_apply_id,omitempty" xml:"out_apply_id,omitempty"`
	SubOrderRefundList []*InsureRefundDetailResponseBodyModuleSubOrderRefundList `json:"sub_order_refund_list,omitempty" xml:"sub_order_refund_list,omitempty" type:"Repeated"`
}

func (s InsureRefundDetailResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s InsureRefundDetailResponseBodyModule) GoString() string {
	return s.String()
}

func (s *InsureRefundDetailResponseBodyModule) GetApplyId() *string {
	return s.ApplyId
}

func (s *InsureRefundDetailResponseBodyModule) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *InsureRefundDetailResponseBodyModule) GetGmtModified() *string {
	return s.GmtModified
}

func (s *InsureRefundDetailResponseBodyModule) GetInsOrderId() *string {
	return s.InsOrderId
}

func (s *InsureRefundDetailResponseBodyModule) GetInsureOrder() *InsureRefundDetailResponseBodyModuleInsureOrder {
	return s.InsureOrder
}

func (s *InsureRefundDetailResponseBodyModule) GetOutApplyId() *string {
	return s.OutApplyId
}

func (s *InsureRefundDetailResponseBodyModule) GetSubOrderRefundList() []*InsureRefundDetailResponseBodyModuleSubOrderRefundList {
	return s.SubOrderRefundList
}

func (s *InsureRefundDetailResponseBodyModule) SetApplyId(v string) *InsureRefundDetailResponseBodyModule {
	s.ApplyId = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModule) SetGmtCreate(v string) *InsureRefundDetailResponseBodyModule {
	s.GmtCreate = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModule) SetGmtModified(v string) *InsureRefundDetailResponseBodyModule {
	s.GmtModified = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModule) SetInsOrderId(v string) *InsureRefundDetailResponseBodyModule {
	s.InsOrderId = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModule) SetInsureOrder(v *InsureRefundDetailResponseBodyModuleInsureOrder) *InsureRefundDetailResponseBodyModule {
	s.InsureOrder = v
	return s
}

func (s *InsureRefundDetailResponseBodyModule) SetOutApplyId(v string) *InsureRefundDetailResponseBodyModule {
	s.OutApplyId = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModule) SetSubOrderRefundList(v []*InsureRefundDetailResponseBodyModuleSubOrderRefundList) *InsureRefundDetailResponseBodyModule {
	s.SubOrderRefundList = v
	return s
}

func (s *InsureRefundDetailResponseBodyModule) Validate() error {
	if s.InsureOrder != nil {
		if err := s.InsureOrder.Validate(); err != nil {
			return err
		}
	}
	if s.SubOrderRefundList != nil {
		for _, item := range s.SubOrderRefundList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type InsureRefundDetailResponseBodyModuleInsureOrder struct {
	Applicant     *InsureRefundDetailResponseBodyModuleInsureOrderApplicant `json:"applicant,omitempty" xml:"applicant,omitempty" type:"Struct"`
	BizOrderId    *string                                                   `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
	BizType       *int32                                                    `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	CloseTime     *string                                                   `json:"close_time,omitempty" xml:"close_time,omitempty"`
	InsOrderId    *string                                                   `json:"ins_order_id,omitempty" xml:"ins_order_id,omitempty"`
	OutInsOrderId *string                                                   `json:"out_ins_order_id,omitempty" xml:"out_ins_order_id,omitempty"`
	PayTime       *string                                                   `json:"pay_time,omitempty" xml:"pay_time,omitempty"`
	Price         *int64                                                    `json:"price,omitempty" xml:"price,omitempty"`
	SettleType    *int32                                                    `json:"settle_type,omitempty" xml:"settle_type,omitempty"`
	Status        *string                                                   `json:"status,omitempty" xml:"status,omitempty"`
}

func (s InsureRefundDetailResponseBodyModuleInsureOrder) String() string {
	return dara.Prettify(s)
}

func (s InsureRefundDetailResponseBodyModuleInsureOrder) GoString() string {
	return s.String()
}

func (s *InsureRefundDetailResponseBodyModuleInsureOrder) GetApplicant() *InsureRefundDetailResponseBodyModuleInsureOrderApplicant {
	return s.Applicant
}

func (s *InsureRefundDetailResponseBodyModuleInsureOrder) GetBizOrderId() *string {
	return s.BizOrderId
}

func (s *InsureRefundDetailResponseBodyModuleInsureOrder) GetBizType() *int32 {
	return s.BizType
}

func (s *InsureRefundDetailResponseBodyModuleInsureOrder) GetCloseTime() *string {
	return s.CloseTime
}

func (s *InsureRefundDetailResponseBodyModuleInsureOrder) GetInsOrderId() *string {
	return s.InsOrderId
}

func (s *InsureRefundDetailResponseBodyModuleInsureOrder) GetOutInsOrderId() *string {
	return s.OutInsOrderId
}

func (s *InsureRefundDetailResponseBodyModuleInsureOrder) GetPayTime() *string {
	return s.PayTime
}

func (s *InsureRefundDetailResponseBodyModuleInsureOrder) GetPrice() *int64 {
	return s.Price
}

func (s *InsureRefundDetailResponseBodyModuleInsureOrder) GetSettleType() *int32 {
	return s.SettleType
}

func (s *InsureRefundDetailResponseBodyModuleInsureOrder) GetStatus() *string {
	return s.Status
}

func (s *InsureRefundDetailResponseBodyModuleInsureOrder) SetApplicant(v *InsureRefundDetailResponseBodyModuleInsureOrderApplicant) *InsureRefundDetailResponseBodyModuleInsureOrder {
	s.Applicant = v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleInsureOrder) SetBizOrderId(v string) *InsureRefundDetailResponseBodyModuleInsureOrder {
	s.BizOrderId = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleInsureOrder) SetBizType(v int32) *InsureRefundDetailResponseBodyModuleInsureOrder {
	s.BizType = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleInsureOrder) SetCloseTime(v string) *InsureRefundDetailResponseBodyModuleInsureOrder {
	s.CloseTime = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleInsureOrder) SetInsOrderId(v string) *InsureRefundDetailResponseBodyModuleInsureOrder {
	s.InsOrderId = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleInsureOrder) SetOutInsOrderId(v string) *InsureRefundDetailResponseBodyModuleInsureOrder {
	s.OutInsOrderId = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleInsureOrder) SetPayTime(v string) *InsureRefundDetailResponseBodyModuleInsureOrder {
	s.PayTime = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleInsureOrder) SetPrice(v int64) *InsureRefundDetailResponseBodyModuleInsureOrder {
	s.Price = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleInsureOrder) SetSettleType(v int32) *InsureRefundDetailResponseBodyModuleInsureOrder {
	s.SettleType = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleInsureOrder) SetStatus(v string) *InsureRefundDetailResponseBodyModuleInsureOrder {
	s.Status = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleInsureOrder) Validate() error {
	if s.Applicant != nil {
		if err := s.Applicant.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InsureRefundDetailResponseBodyModuleInsureOrderApplicant struct {
	CertName *string `json:"cert_name,omitempty" xml:"cert_name,omitempty"`
	CertNo   *string `json:"cert_no,omitempty" xml:"cert_no,omitempty"`
	CertType *string `json:"cert_type,omitempty" xml:"cert_type,omitempty"`
	Phone    *string `json:"phone,omitempty" xml:"phone,omitempty"`
}

func (s InsureRefundDetailResponseBodyModuleInsureOrderApplicant) String() string {
	return dara.Prettify(s)
}

func (s InsureRefundDetailResponseBodyModuleInsureOrderApplicant) GoString() string {
	return s.String()
}

func (s *InsureRefundDetailResponseBodyModuleInsureOrderApplicant) GetCertName() *string {
	return s.CertName
}

func (s *InsureRefundDetailResponseBodyModuleInsureOrderApplicant) GetCertNo() *string {
	return s.CertNo
}

func (s *InsureRefundDetailResponseBodyModuleInsureOrderApplicant) GetCertType() *string {
	return s.CertType
}

func (s *InsureRefundDetailResponseBodyModuleInsureOrderApplicant) GetPhone() *string {
	return s.Phone
}

func (s *InsureRefundDetailResponseBodyModuleInsureOrderApplicant) SetCertName(v string) *InsureRefundDetailResponseBodyModuleInsureOrderApplicant {
	s.CertName = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleInsureOrderApplicant) SetCertNo(v string) *InsureRefundDetailResponseBodyModuleInsureOrderApplicant {
	s.CertNo = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleInsureOrderApplicant) SetCertType(v string) *InsureRefundDetailResponseBodyModuleInsureOrderApplicant {
	s.CertType = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleInsureOrderApplicant) SetPhone(v string) *InsureRefundDetailResponseBodyModuleInsureOrderApplicant {
	s.Phone = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleInsureOrderApplicant) Validate() error {
	return dara.Validate(s)
}

type InsureRefundDetailResponseBodyModuleSubOrderRefundList struct {
	EffectiveEndTime   *string                                                              `json:"effective_end_time,omitempty" xml:"effective_end_time,omitempty"`
	EffectiveStartTime *string                                                              `json:"effective_start_time,omitempty" xml:"effective_start_time,omitempty"`
	InsureSegment      *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsureSegment `json:"insure_segment,omitempty" xml:"insure_segment,omitempty" type:"Struct"`
	InsureTime         *string                                                              `json:"insure_time,omitempty" xml:"insure_time,omitempty"`
	Insured            *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsured       `json:"insured,omitempty" xml:"insured,omitempty" type:"Struct"`
	OutSubInsOrderId   *string                                                              `json:"out_sub_ins_order_id,omitempty" xml:"out_sub_ins_order_id,omitempty"`
	PolicyNo           *string                                                              `json:"policy_no,omitempty" xml:"policy_no,omitempty"`
	PolicyRefundNo     *string                                                              `json:"policy_refund_no,omitempty" xml:"policy_refund_no,omitempty"`
	Price              *int64                                                               `json:"price,omitempty" xml:"price,omitempty"`
	ProductName        *string                                                              `json:"product_name,omitempty" xml:"product_name,omitempty"`
	ProductNo          *string                                                              `json:"product_no,omitempty" xml:"product_no,omitempty"`
	RefundStatus       *string                                                              `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	RefundTime         *string                                                              `json:"refund_time,omitempty" xml:"refund_time,omitempty"`
	Status             *string                                                              `json:"status,omitempty" xml:"status,omitempty"`
	SubInsOrderId      *string                                                              `json:"sub_ins_order_id,omitempty" xml:"sub_ins_order_id,omitempty"`
}

func (s InsureRefundDetailResponseBodyModuleSubOrderRefundList) String() string {
	return dara.Prettify(s)
}

func (s InsureRefundDetailResponseBodyModuleSubOrderRefundList) GoString() string {
	return s.String()
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundList) GetEffectiveEndTime() *string {
	return s.EffectiveEndTime
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundList) GetEffectiveStartTime() *string {
	return s.EffectiveStartTime
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundList) GetInsureSegment() *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsureSegment {
	return s.InsureSegment
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundList) GetInsureTime() *string {
	return s.InsureTime
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundList) GetInsured() *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsured {
	return s.Insured
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundList) GetOutSubInsOrderId() *string {
	return s.OutSubInsOrderId
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundList) GetPolicyNo() *string {
	return s.PolicyNo
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundList) GetPolicyRefundNo() *string {
	return s.PolicyRefundNo
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundList) GetPrice() *int64 {
	return s.Price
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundList) GetProductName() *string {
	return s.ProductName
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundList) GetProductNo() *string {
	return s.ProductNo
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundList) GetRefundStatus() *string {
	return s.RefundStatus
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundList) GetRefundTime() *string {
	return s.RefundTime
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundList) GetStatus() *string {
	return s.Status
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundList) GetSubInsOrderId() *string {
	return s.SubInsOrderId
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundList) SetEffectiveEndTime(v string) *InsureRefundDetailResponseBodyModuleSubOrderRefundList {
	s.EffectiveEndTime = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundList) SetEffectiveStartTime(v string) *InsureRefundDetailResponseBodyModuleSubOrderRefundList {
	s.EffectiveStartTime = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundList) SetInsureSegment(v *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsureSegment) *InsureRefundDetailResponseBodyModuleSubOrderRefundList {
	s.InsureSegment = v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundList) SetInsureTime(v string) *InsureRefundDetailResponseBodyModuleSubOrderRefundList {
	s.InsureTime = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundList) SetInsured(v *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsured) *InsureRefundDetailResponseBodyModuleSubOrderRefundList {
	s.Insured = v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundList) SetOutSubInsOrderId(v string) *InsureRefundDetailResponseBodyModuleSubOrderRefundList {
	s.OutSubInsOrderId = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundList) SetPolicyNo(v string) *InsureRefundDetailResponseBodyModuleSubOrderRefundList {
	s.PolicyNo = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundList) SetPolicyRefundNo(v string) *InsureRefundDetailResponseBodyModuleSubOrderRefundList {
	s.PolicyRefundNo = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundList) SetPrice(v int64) *InsureRefundDetailResponseBodyModuleSubOrderRefundList {
	s.Price = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundList) SetProductName(v string) *InsureRefundDetailResponseBodyModuleSubOrderRefundList {
	s.ProductName = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundList) SetProductNo(v string) *InsureRefundDetailResponseBodyModuleSubOrderRefundList {
	s.ProductNo = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundList) SetRefundStatus(v string) *InsureRefundDetailResponseBodyModuleSubOrderRefundList {
	s.RefundStatus = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundList) SetRefundTime(v string) *InsureRefundDetailResponseBodyModuleSubOrderRefundList {
	s.RefundTime = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundList) SetStatus(v string) *InsureRefundDetailResponseBodyModuleSubOrderRefundList {
	s.Status = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundList) SetSubInsOrderId(v string) *InsureRefundDetailResponseBodyModuleSubOrderRefundList {
	s.SubInsOrderId = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundList) Validate() error {
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

type InsureRefundDetailResponseBodyModuleSubOrderRefundListInsureSegment struct {
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

func (s InsureRefundDetailResponseBodyModuleSubOrderRefundListInsureSegment) String() string {
	return dara.Prettify(s)
}

func (s InsureRefundDetailResponseBodyModuleSubOrderRefundListInsureSegment) GoString() string {
	return s.String()
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsureSegment) GetArrAirportCode() *string {
	return s.ArrAirportCode
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsureSegment) GetArrCity() *string {
	return s.ArrCity
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsureSegment) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsureSegment) GetArrTime() *string {
	return s.ArrTime
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsureSegment) GetDepAirportCode() *string {
	return s.DepAirportCode
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsureSegment) GetDepCity() *string {
	return s.DepCity
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsureSegment) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsureSegment) GetDepTime() *string {
	return s.DepTime
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsureSegment) GetFlightNo() *string {
	return s.FlightNo
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsureSegment) SetArrAirportCode(v string) *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsureSegment {
	s.ArrAirportCode = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsureSegment) SetArrCity(v string) *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsureSegment {
	s.ArrCity = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsureSegment) SetArrCityCode(v string) *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsureSegment {
	s.ArrCityCode = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsureSegment) SetArrTime(v string) *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsureSegment {
	s.ArrTime = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsureSegment) SetDepAirportCode(v string) *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsureSegment {
	s.DepAirportCode = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsureSegment) SetDepCity(v string) *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsureSegment {
	s.DepCity = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsureSegment) SetDepCityCode(v string) *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsureSegment {
	s.DepCityCode = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsureSegment) SetDepTime(v string) *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsureSegment {
	s.DepTime = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsureSegment) SetFlightNo(v string) *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsureSegment {
	s.FlightNo = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsureSegment) Validate() error {
	return dara.Validate(s)
}

type InsureRefundDetailResponseBodyModuleSubOrderRefundListInsured struct {
	Birthday    *string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	BtripUserId *string `json:"btrip_user_id,omitempty" xml:"btrip_user_id,omitempty"`
	CertName    *string `json:"cert_name,omitempty" xml:"cert_name,omitempty"`
	CertNo      *string `json:"cert_no,omitempty" xml:"cert_no,omitempty"`
	CertType    *string `json:"cert_type,omitempty" xml:"cert_type,omitempty"`
	Gender      *string `json:"gender,omitempty" xml:"gender,omitempty"`
	Phone       *string `json:"phone,omitempty" xml:"phone,omitempty"`
}

func (s InsureRefundDetailResponseBodyModuleSubOrderRefundListInsured) String() string {
	return dara.Prettify(s)
}

func (s InsureRefundDetailResponseBodyModuleSubOrderRefundListInsured) GoString() string {
	return s.String()
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsured) GetBirthday() *string {
	return s.Birthday
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsured) GetBtripUserId() *string {
	return s.BtripUserId
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsured) GetCertName() *string {
	return s.CertName
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsured) GetCertNo() *string {
	return s.CertNo
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsured) GetCertType() *string {
	return s.CertType
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsured) GetGender() *string {
	return s.Gender
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsured) GetPhone() *string {
	return s.Phone
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsured) SetBirthday(v string) *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsured {
	s.Birthday = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsured) SetBtripUserId(v string) *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsured {
	s.BtripUserId = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsured) SetCertName(v string) *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsured {
	s.CertName = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsured) SetCertNo(v string) *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsured {
	s.CertNo = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsured) SetCertType(v string) *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsured {
	s.CertType = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsured) SetGender(v string) *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsured {
	s.Gender = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsured) SetPhone(v string) *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsured {
	s.Phone = &v
	return s
}

func (s *InsureRefundDetailResponseBodyModuleSubOrderRefundListInsured) Validate() error {
	return dara.Validate(s)
}
