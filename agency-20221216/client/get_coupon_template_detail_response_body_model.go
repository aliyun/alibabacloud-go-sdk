// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCouponTemplateDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCouponTemplateDetailResponseBody
	GetCode() *string
	SetData(v *GetCouponTemplateDetailResponseBodyData) *GetCouponTemplateDetailResponseBody
	GetData() *GetCouponTemplateDetailResponseBodyData
	SetMessage(v string) *GetCouponTemplateDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCouponTemplateDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCouponTemplateDetailResponseBody
	GetSuccess() *bool
}

type GetCouponTemplateDetailResponseBody struct {
	// example:
	//
	// 200
	Code    *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data    *GetCouponTemplateDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCouponTemplateDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCouponTemplateDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetCouponTemplateDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCouponTemplateDetailResponseBody) GetData() *GetCouponTemplateDetailResponseBodyData {
	return s.Data
}

func (s *GetCouponTemplateDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCouponTemplateDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCouponTemplateDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCouponTemplateDetailResponseBody) SetCode(v string) *GetCouponTemplateDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetCouponTemplateDetailResponseBody) SetData(v *GetCouponTemplateDetailResponseBodyData) *GetCouponTemplateDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetCouponTemplateDetailResponseBody) SetMessage(v string) *GetCouponTemplateDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetCouponTemplateDetailResponseBody) SetRequestId(v string) *GetCouponTemplateDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCouponTemplateDetailResponseBody) SetSuccess(v bool) *GetCouponTemplateDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetCouponTemplateDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCouponTemplateDetailResponseBodyData struct {
	// example:
	//
	// UNIVERSAL
	ApplicableProducts *string `json:"ApplicableProducts,omitempty" xml:"ApplicableProducts,omitempty"`
	// example:
	//
	// aliyun_poc
	CostBearer        *string `json:"CostBearer,omitempty" xml:"CostBearer,omitempty"`
	CouponDescription *string `json:"CouponDescription,omitempty" xml:"CouponDescription,omitempty"`
	// example:
	//
	// 2024-11-21 18:18:22
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// 100
	Denomination *float64 `json:"Denomination,omitempty" xml:"Denomination,omitempty"`
	// example:
	//
	// 1
	LimitPerPerson *int32 `json:"LimitPerPerson,omitempty" xml:"LimitPerPerson,omitempty"`
	// example:
	//
	// ALL,BILLING
	PurchaseType         *string `json:"PurchaseType,omitempty" xml:"PurchaseType,omitempty"`
	ReasonForApplication *string `json:"ReasonForApplication,omitempty" xml:"ReasonForApplication,omitempty"`
	// example:
	//
	// APPROVED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1576
	TemplateId   *int64  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// 100
	ValidUntil *string `json:"ValidUntil,omitempty" xml:"ValidUntil,omitempty"`
	// example:
	//
	// 0
	ValidUntilType *string `json:"ValidUntilType,omitempty" xml:"ValidUntilType,omitempty"`
}

func (s GetCouponTemplateDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCouponTemplateDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCouponTemplateDetailResponseBodyData) GetApplicableProducts() *string {
	return s.ApplicableProducts
}

func (s *GetCouponTemplateDetailResponseBodyData) GetCostBearer() *string {
	return s.CostBearer
}

func (s *GetCouponTemplateDetailResponseBodyData) GetCouponDescription() *string {
	return s.CouponDescription
}

func (s *GetCouponTemplateDetailResponseBodyData) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *GetCouponTemplateDetailResponseBodyData) GetDenomination() *float64 {
	return s.Denomination
}

func (s *GetCouponTemplateDetailResponseBodyData) GetLimitPerPerson() *int32 {
	return s.LimitPerPerson
}

func (s *GetCouponTemplateDetailResponseBodyData) GetPurchaseType() *string {
	return s.PurchaseType
}

func (s *GetCouponTemplateDetailResponseBodyData) GetReasonForApplication() *string {
	return s.ReasonForApplication
}

func (s *GetCouponTemplateDetailResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetCouponTemplateDetailResponseBodyData) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *GetCouponTemplateDetailResponseBodyData) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetCouponTemplateDetailResponseBodyData) GetValidUntil() *string {
	return s.ValidUntil
}

func (s *GetCouponTemplateDetailResponseBodyData) GetValidUntilType() *string {
	return s.ValidUntilType
}

func (s *GetCouponTemplateDetailResponseBodyData) SetApplicableProducts(v string) *GetCouponTemplateDetailResponseBodyData {
	s.ApplicableProducts = &v
	return s
}

func (s *GetCouponTemplateDetailResponseBodyData) SetCostBearer(v string) *GetCouponTemplateDetailResponseBodyData {
	s.CostBearer = &v
	return s
}

func (s *GetCouponTemplateDetailResponseBodyData) SetCouponDescription(v string) *GetCouponTemplateDetailResponseBodyData {
	s.CouponDescription = &v
	return s
}

func (s *GetCouponTemplateDetailResponseBodyData) SetCreatedTime(v string) *GetCouponTemplateDetailResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *GetCouponTemplateDetailResponseBodyData) SetDenomination(v float64) *GetCouponTemplateDetailResponseBodyData {
	s.Denomination = &v
	return s
}

func (s *GetCouponTemplateDetailResponseBodyData) SetLimitPerPerson(v int32) *GetCouponTemplateDetailResponseBodyData {
	s.LimitPerPerson = &v
	return s
}

func (s *GetCouponTemplateDetailResponseBodyData) SetPurchaseType(v string) *GetCouponTemplateDetailResponseBodyData {
	s.PurchaseType = &v
	return s
}

func (s *GetCouponTemplateDetailResponseBodyData) SetReasonForApplication(v string) *GetCouponTemplateDetailResponseBodyData {
	s.ReasonForApplication = &v
	return s
}

func (s *GetCouponTemplateDetailResponseBodyData) SetStatus(v string) *GetCouponTemplateDetailResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetCouponTemplateDetailResponseBodyData) SetTemplateId(v int64) *GetCouponTemplateDetailResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *GetCouponTemplateDetailResponseBodyData) SetTemplateName(v string) *GetCouponTemplateDetailResponseBodyData {
	s.TemplateName = &v
	return s
}

func (s *GetCouponTemplateDetailResponseBodyData) SetValidUntil(v string) *GetCouponTemplateDetailResponseBodyData {
	s.ValidUntil = &v
	return s
}

func (s *GetCouponTemplateDetailResponseBodyData) SetValidUntilType(v string) *GetCouponTemplateDetailResponseBodyData {
	s.ValidUntilType = &v
	return s
}

func (s *GetCouponTemplateDetailResponseBodyData) Validate() error {
	return dara.Validate(s)
}
