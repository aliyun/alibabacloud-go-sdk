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
	SetApplicationReason(v string) *IssueCouponForCustomerRequest
	GetApplicationReason() *string
	SetCouponTemplateId(v int64) *IssueCouponForCustomerRequest
	GetCouponTemplateId() *int64
	SetIsUseBenefit(v bool) *IssueCouponForCustomerRequest
	GetIsUseBenefit() *bool
	SetUidlist(v string) *IssueCouponForCustomerRequest
	GetUidlist() *string
}

type IssueCouponForCustomerRequest struct {
	// The language parameter. Default value: none. Valid values:
	//
	// - zh-CN: Chinese.
	//
	// - en-US: English.
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The reason for applying for the coupon. The reason can be up to 500 characters in length. Example: Apply for a coupon for Customer XXX.
	//
	// example:
	//
	// Apply for a coupon for Customer XXX, who is XXXX
	ApplicationReason *string `json:"ApplicationReason,omitempty" xml:"ApplicationReason,omitempty"`
	// The ID of the coupon template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5075915
	CouponTemplateId *int64 `json:"CouponTemplateId,omitempty" xml:"CouponTemplateId,omitempty"`
	// Deprecated
	//
	// Specifies whether to use the POC privilege quota. Valid values:
	//
	// - true: Uses the POC privilege quota. If all submitted accounts are marked as approval-exempt, the current coupon application skips the approval process, and the POC privilege quota in your partner benefit package is deducted.
	//
	// - false: Does not use the POC privilege quota. Regardless of the approval-exempt status of the submitted accounts, your POC coupon benefit quota remains unchanged. In this case, the application follows the standard coupon application approval process.
	//
	// example:
	//
	// true
	IsUseBenefit *bool `json:"IsUseBenefit,omitempty" xml:"IsUseBenefit,omitempty"`
	// The list of UIDs. Separate multiple UIDs with commas (,). Example: 111,2222.
	//
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

func (s *IssueCouponForCustomerRequest) GetApplicationReason() *string {
	return s.ApplicationReason
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

func (s *IssueCouponForCustomerRequest) SetApplicationReason(v string) *IssueCouponForCustomerRequest {
	s.ApplicationReason = &v
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
