// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIssueCouponForCustomerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *IssueCouponForCustomerRequest
	GetAcceptLanguage() *string
	SetCouponTemplateId(v int64) *IssueCouponForCustomerRequest
	GetCouponTemplateId() *int64
	SetIsUseBenefit(v bool) *IssueCouponForCustomerRequest
	GetIsUseBenefit() *bool
	SetUidlist(v string) *IssueCouponForCustomerRequest
	GetUidlist() *string
}

type IssueCouponForCustomerRequest struct {
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5075915
	CouponTemplateId *int64 `json:"CouponTemplateId,omitempty" xml:"CouponTemplateId,omitempty"`
	IsUseBenefit     *bool  `json:"IsUseBenefit,omitempty" xml:"IsUseBenefit,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 111,2222
	Uidlist *string `json:"Uidlist,omitempty" xml:"Uidlist,omitempty"`
}

func (s IssueCouponForCustomerRequest) String() string {
	return dara.Prettify(s)
}

func (s IssueCouponForCustomerRequest) GoString() string {
	return s.String()
}

func (s *IssueCouponForCustomerRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *IssueCouponForCustomerRequest) GetCouponTemplateId() *int64 {
	return s.CouponTemplateId
}

func (s *IssueCouponForCustomerRequest) GetIsUseBenefit() *bool {
	return s.IsUseBenefit
}

func (s *IssueCouponForCustomerRequest) GetUidlist() *string {
	return s.Uidlist
}

func (s *IssueCouponForCustomerRequest) SetAcceptLanguage(v string) *IssueCouponForCustomerRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *IssueCouponForCustomerRequest) SetCouponTemplateId(v int64) *IssueCouponForCustomerRequest {
	s.CouponTemplateId = &v
	return s
}

func (s *IssueCouponForCustomerRequest) SetIsUseBenefit(v bool) *IssueCouponForCustomerRequest {
	s.IsUseBenefit = &v
	return s
}

func (s *IssueCouponForCustomerRequest) SetUidlist(v string) *IssueCouponForCustomerRequest {
	s.Uidlist = &v
	return s
}

func (s *IssueCouponForCustomerRequest) Validate() error {
	return dara.Validate(s)
}
