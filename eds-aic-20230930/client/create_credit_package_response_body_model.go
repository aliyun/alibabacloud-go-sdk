// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCreditPackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreditPackageId(v string) *CreateCreditPackageResponseBody
	GetCreditPackageId() *string
	SetCreditPackageIds(v []*string) *CreateCreditPackageResponseBody
	GetCreditPackageIds() []*string
	SetEffectiveTime(v string) *CreateCreditPackageResponseBody
	GetEffectiveTime() *string
	SetExpiredTime(v string) *CreateCreditPackageResponseBody
	GetExpiredTime() *string
	SetOrderId(v string) *CreateCreditPackageResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateCreditPackageResponseBody
	GetRequestId() *string
}

type CreateCreditPackageResponseBody struct {
	// The ID of the credit booster pack.
	//
	// example:
	//
	// crp-bt7e2t4anbq50****
	CreditPackageId  *string   `json:"CreditPackageId,omitempty" xml:"CreditPackageId,omitempty"`
	CreditPackageIds []*string `json:"CreditPackageIds,omitempty" xml:"CreditPackageIds,omitempty" type:"Repeated"`
	// The effective period of the credit booster pack.
	//
	// example:
	//
	// 2026-04-30 00:00:00
	EffectiveTime *string `json:"EffectiveTime,omitempty" xml:"EffectiveTime,omitempty"`
	// The time when the credit booster pack expires.
	//
	// example:
	//
	// 2026-10-30 00:00:00
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 223684716098****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F07A1DA1-E1EB-5CCA-8EED-12F85D32****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCreditPackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCreditPackageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCreditPackageResponseBody) GetCreditPackageId() *string {
	return s.CreditPackageId
}

func (s *CreateCreditPackageResponseBody) GetCreditPackageIds() []*string {
	return s.CreditPackageIds
}

func (s *CreateCreditPackageResponseBody) GetEffectiveTime() *string {
	return s.EffectiveTime
}

func (s *CreateCreditPackageResponseBody) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *CreateCreditPackageResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateCreditPackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCreditPackageResponseBody) SetCreditPackageId(v string) *CreateCreditPackageResponseBody {
	s.CreditPackageId = &v
	return s
}

func (s *CreateCreditPackageResponseBody) SetCreditPackageIds(v []*string) *CreateCreditPackageResponseBody {
	s.CreditPackageIds = v
	return s
}

func (s *CreateCreditPackageResponseBody) SetEffectiveTime(v string) *CreateCreditPackageResponseBody {
	s.EffectiveTime = &v
	return s
}

func (s *CreateCreditPackageResponseBody) SetExpiredTime(v string) *CreateCreditPackageResponseBody {
	s.ExpiredTime = &v
	return s
}

func (s *CreateCreditPackageResponseBody) SetOrderId(v string) *CreateCreditPackageResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateCreditPackageResponseBody) SetRequestId(v string) *CreateCreditPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCreditPackageResponseBody) Validate() error {
	return dara.Validate(s)
}
