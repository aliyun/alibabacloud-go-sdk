// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCashCouponsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryCashCouponsResponseBody
	GetCode() *string
	SetData(v *QueryCashCouponsResponseBodyData) *QueryCashCouponsResponseBody
	GetData() *QueryCashCouponsResponseBodyData
	SetMessage(v string) *QueryCashCouponsResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryCashCouponsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryCashCouponsResponseBody
	GetSuccess() *bool
}

type QueryCashCouponsResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *QueryCashCouponsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 147B566E-DB4C-4E43-BDBB-5DB1D9D268DB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryCashCouponsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCashCouponsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCashCouponsResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryCashCouponsResponseBody) GetData() *QueryCashCouponsResponseBodyData {
	return s.Data
}

func (s *QueryCashCouponsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryCashCouponsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCashCouponsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryCashCouponsResponseBody) SetCode(v string) *QueryCashCouponsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCashCouponsResponseBody) SetData(v *QueryCashCouponsResponseBodyData) *QueryCashCouponsResponseBody {
	s.Data = v
	return s
}

func (s *QueryCashCouponsResponseBody) SetMessage(v string) *QueryCashCouponsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCashCouponsResponseBody) SetRequestId(v string) *QueryCashCouponsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCashCouponsResponseBody) SetSuccess(v bool) *QueryCashCouponsResponseBody {
	s.Success = &v
	return s
}

func (s *QueryCashCouponsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryCashCouponsResponseBodyData struct {
	CashCoupon []*QueryCashCouponsResponseBodyDataCashCoupon `json:"CashCoupon,omitempty" xml:"CashCoupon,omitempty" type:"Repeated"`
}

func (s QueryCashCouponsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryCashCouponsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCashCouponsResponseBodyData) GetCashCoupon() []*QueryCashCouponsResponseBodyDataCashCoupon {
	return s.CashCoupon
}

func (s *QueryCashCouponsResponseBodyData) SetCashCoupon(v []*QueryCashCouponsResponseBodyDataCashCoupon) *QueryCashCouponsResponseBodyData {
	s.CashCoupon = v
	return s
}

func (s *QueryCashCouponsResponseBodyData) Validate() error {
	if s.CashCoupon != nil {
		for _, item := range s.CashCoupon {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryCashCouponsResponseBodyDataCashCoupon struct {
	// The service to which the voucher is applicable.
	//
	// example:
	//
	// All Alibaba Cloud services
	ApplicableProducts *string `json:"ApplicableProducts,omitempty" xml:"ApplicableProducts,omitempty"`
	// The scenario to which the voucher is applicable.
	//
	// example:
	//
	// Pay for the pay-as-you-go bills of Alibaba Cloud services or purchase an instance of an Alibaba Cloud service
	ApplicableScenarios *string `json:"ApplicableScenarios,omitempty" xml:"ApplicableScenarios,omitempty"`
	// The remaining quota of the voucher.
	//
	// example:
	//
	// 100.00
	Balance *string `json:"Balance,omitempty" xml:"Balance,omitempty"`
	// The ID of the voucher.
	//
	// example:
	//
	// 34534253254325
	CashCouponId *int64 `json:"CashCouponId,omitempty" xml:"CashCouponId,omitempty"`
	// The code of the voucher.
	//
	// example:
	//
	// Q-b1485def8f04a
	CashCouponNo *string `json:"CashCouponNo,omitempty" xml:"CashCouponNo,omitempty"`
	// The description of the voucher.
	//
	// example:
	//
	// This voucher is used for testing product function
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the voucher took effect.
	//
	// example:
	//
	// 2018-08-02T15:15:50Z
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The time when the voucher expired.
	//
	// example:
	//
	// 2019-01-29T15:15:50Z
	ExpiryTime *string `json:"ExpiryTime,omitempty" xml:"ExpiryTime,omitempty"`
	// The time when the voucher was released.
	//
	// example:
	//
	// 2018-08-02T15:15:50Z
	GrantedTime *string `json:"GrantedTime,omitempty" xml:"GrantedTime,omitempty"`
	// The denomination of the voucher.
	//
	// example:
	//
	// 100.00
	NominalValue *string `json:"NominalValue,omitempty" xml:"NominalValue,omitempty"`
	// The state of the voucher. Valid values:
	//
	// 	- Available: The voucher is valid.
	//
	// 	- Expired: The voucher has expired.
	//
	// 	- Cancelled: The voucher is canceled.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryCashCouponsResponseBodyDataCashCoupon) String() string {
	return dara.Prettify(s)
}

func (s QueryCashCouponsResponseBodyDataCashCoupon) GoString() string {
	return s.String()
}

func (s *QueryCashCouponsResponseBodyDataCashCoupon) GetApplicableProducts() *string {
	return s.ApplicableProducts
}

func (s *QueryCashCouponsResponseBodyDataCashCoupon) GetApplicableScenarios() *string {
	return s.ApplicableScenarios
}

func (s *QueryCashCouponsResponseBodyDataCashCoupon) GetBalance() *string {
	return s.Balance
}

func (s *QueryCashCouponsResponseBodyDataCashCoupon) GetCashCouponId() *int64 {
	return s.CashCouponId
}

func (s *QueryCashCouponsResponseBodyDataCashCoupon) GetCashCouponNo() *string {
	return s.CashCouponNo
}

func (s *QueryCashCouponsResponseBodyDataCashCoupon) GetDescription() *string {
	return s.Description
}

func (s *QueryCashCouponsResponseBodyDataCashCoupon) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *QueryCashCouponsResponseBodyDataCashCoupon) GetExpiryTime() *string {
	return s.ExpiryTime
}

func (s *QueryCashCouponsResponseBodyDataCashCoupon) GetGrantedTime() *string {
	return s.GrantedTime
}

func (s *QueryCashCouponsResponseBodyDataCashCoupon) GetNominalValue() *string {
	return s.NominalValue
}

func (s *QueryCashCouponsResponseBodyDataCashCoupon) GetStatus() *string {
	return s.Status
}

func (s *QueryCashCouponsResponseBodyDataCashCoupon) SetApplicableProducts(v string) *QueryCashCouponsResponseBodyDataCashCoupon {
	s.ApplicableProducts = &v
	return s
}

func (s *QueryCashCouponsResponseBodyDataCashCoupon) SetApplicableScenarios(v string) *QueryCashCouponsResponseBodyDataCashCoupon {
	s.ApplicableScenarios = &v
	return s
}

func (s *QueryCashCouponsResponseBodyDataCashCoupon) SetBalance(v string) *QueryCashCouponsResponseBodyDataCashCoupon {
	s.Balance = &v
	return s
}

func (s *QueryCashCouponsResponseBodyDataCashCoupon) SetCashCouponId(v int64) *QueryCashCouponsResponseBodyDataCashCoupon {
	s.CashCouponId = &v
	return s
}

func (s *QueryCashCouponsResponseBodyDataCashCoupon) SetCashCouponNo(v string) *QueryCashCouponsResponseBodyDataCashCoupon {
	s.CashCouponNo = &v
	return s
}

func (s *QueryCashCouponsResponseBodyDataCashCoupon) SetDescription(v string) *QueryCashCouponsResponseBodyDataCashCoupon {
	s.Description = &v
	return s
}

func (s *QueryCashCouponsResponseBodyDataCashCoupon) SetEffectiveTime(v string) *QueryCashCouponsResponseBodyDataCashCoupon {
	s.EffectiveTime = &v
	return s
}

func (s *QueryCashCouponsResponseBodyDataCashCoupon) SetExpiryTime(v string) *QueryCashCouponsResponseBodyDataCashCoupon {
	s.ExpiryTime = &v
	return s
}

func (s *QueryCashCouponsResponseBodyDataCashCoupon) SetGrantedTime(v string) *QueryCashCouponsResponseBodyDataCashCoupon {
	s.GrantedTime = &v
	return s
}

func (s *QueryCashCouponsResponseBodyDataCashCoupon) SetNominalValue(v string) *QueryCashCouponsResponseBodyDataCashCoupon {
	s.NominalValue = &v
	return s
}

func (s *QueryCashCouponsResponseBodyDataCashCoupon) SetStatus(v string) *QueryCashCouponsResponseBodyDataCashCoupon {
	s.Status = &v
	return s
}

func (s *QueryCashCouponsResponseBodyDataCashCoupon) Validate() error {
	return dara.Validate(s)
}
