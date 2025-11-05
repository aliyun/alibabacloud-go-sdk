// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCouponTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateCouponTemplateResponseBody
	GetCode() *string
	SetData(v *CreateCouponTemplateResponseBodyData) *CreateCouponTemplateResponseBody
	GetData() *CreateCouponTemplateResponseBodyData
	SetMessage(v string) *CreateCouponTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateCouponTemplateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateCouponTemplateResponseBody
	GetSuccess() *bool
}

type CreateCouponTemplateResponseBody struct {
	// example:
	//
	// 200
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateCouponTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2103a30617045934095083027d88c5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateCouponTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCouponTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCouponTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateCouponTemplateResponseBody) GetData() *CreateCouponTemplateResponseBodyData {
	return s.Data
}

func (s *CreateCouponTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateCouponTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCouponTemplateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateCouponTemplateResponseBody) SetCode(v string) *CreateCouponTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCouponTemplateResponseBody) SetData(v *CreateCouponTemplateResponseBodyData) *CreateCouponTemplateResponseBody {
	s.Data = v
	return s
}

func (s *CreateCouponTemplateResponseBody) SetMessage(v string) *CreateCouponTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCouponTemplateResponseBody) SetRequestId(v string) *CreateCouponTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCouponTemplateResponseBody) SetSuccess(v bool) *CreateCouponTemplateResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCouponTemplateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCouponTemplateResponseBodyData struct {
	// example:
	//
	// Custom
	ApplicableProducts *string `json:"ApplicableProducts,omitempty" xml:"ApplicableProducts,omitempty"`
	// example:
	//
	// Partner
	CostBearer *string `json:"CostBearer,omitempty" xml:"CostBearer,omitempty"`
	// example:
	//
	// 111111
	CouponTemplateID *int64 `json:"CouponTemplateID,omitempty" xml:"CouponTemplateID,omitempty"`
	// example:
	//
	// 2024-04-02 16:15:31
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2024-01-01
	Expireddate *string   `json:"Expireddate,omitempty" xml:"Expireddate,omitempty"`
	ProductType []*string `json:"ProductType,omitempty" xml:"ProductType,omitempty" type:"Repeated"`
	// example:
	//
	// APPROVED
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// example:
	//
	// 2024-01-01
	Vailddate *string `json:"Vailddate,omitempty" xml:"Vailddate,omitempty"`
	// example:
	//
	// 1
	Vaildperioddays *string `json:"Vaildperioddays,omitempty" xml:"Vaildperioddays,omitempty"`
	// example:
	//
	// Validity Duration
	ValidUntil *string `json:"ValidUntil,omitempty" xml:"ValidUntil,omitempty"`
	// example:
	//
	// 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateCouponTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateCouponTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateCouponTemplateResponseBodyData) GetApplicableProducts() *string {
	return s.ApplicableProducts
}

func (s *CreateCouponTemplateResponseBodyData) GetCostBearer() *string {
	return s.CostBearer
}

func (s *CreateCouponTemplateResponseBodyData) GetCouponTemplateID() *int64 {
	return s.CouponTemplateID
}

func (s *CreateCouponTemplateResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateCouponTemplateResponseBodyData) GetExpireddate() *string {
	return s.Expireddate
}

func (s *CreateCouponTemplateResponseBodyData) GetProductType() []*string {
	return s.ProductType
}

func (s *CreateCouponTemplateResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateCouponTemplateResponseBodyData) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateCouponTemplateResponseBodyData) GetVailddate() *string {
	return s.Vailddate
}

func (s *CreateCouponTemplateResponseBodyData) GetVaildperioddays() *string {
	return s.Vaildperioddays
}

func (s *CreateCouponTemplateResponseBodyData) GetValidUntil() *string {
	return s.ValidUntil
}

func (s *CreateCouponTemplateResponseBodyData) GetValue() *string {
	return s.Value
}

func (s *CreateCouponTemplateResponseBodyData) SetApplicableProducts(v string) *CreateCouponTemplateResponseBodyData {
	s.ApplicableProducts = &v
	return s
}

func (s *CreateCouponTemplateResponseBodyData) SetCostBearer(v string) *CreateCouponTemplateResponseBodyData {
	s.CostBearer = &v
	return s
}

func (s *CreateCouponTemplateResponseBodyData) SetCouponTemplateID(v int64) *CreateCouponTemplateResponseBodyData {
	s.CouponTemplateID = &v
	return s
}

func (s *CreateCouponTemplateResponseBodyData) SetCreateTime(v string) *CreateCouponTemplateResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateCouponTemplateResponseBodyData) SetExpireddate(v string) *CreateCouponTemplateResponseBodyData {
	s.Expireddate = &v
	return s
}

func (s *CreateCouponTemplateResponseBodyData) SetProductType(v []*string) *CreateCouponTemplateResponseBodyData {
	s.ProductType = v
	return s
}

func (s *CreateCouponTemplateResponseBodyData) SetStatus(v string) *CreateCouponTemplateResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateCouponTemplateResponseBodyData) SetTemplateName(v string) *CreateCouponTemplateResponseBodyData {
	s.TemplateName = &v
	return s
}

func (s *CreateCouponTemplateResponseBodyData) SetVailddate(v string) *CreateCouponTemplateResponseBodyData {
	s.Vailddate = &v
	return s
}

func (s *CreateCouponTemplateResponseBodyData) SetVaildperioddays(v string) *CreateCouponTemplateResponseBodyData {
	s.Vaildperioddays = &v
	return s
}

func (s *CreateCouponTemplateResponseBodyData) SetValidUntil(v string) *CreateCouponTemplateResponseBodyData {
	s.ValidUntil = &v
	return s
}

func (s *CreateCouponTemplateResponseBodyData) SetValue(v string) *CreateCouponTemplateResponseBodyData {
	s.Value = &v
	return s
}

func (s *CreateCouponTemplateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
