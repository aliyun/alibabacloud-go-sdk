// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIntlFixedPriceDomainOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *CreateIntlFixedPriceDomainOrderRequest
	GetAutoPay() *bool
	SetContactId(v int64) *CreateIntlFixedPriceDomainOrderRequest
	GetContactId() *int64
	SetDomain(v string) *CreateIntlFixedPriceDomainOrderRequest
	GetDomain() *string
	SetExpectedPrice(v int64) *CreateIntlFixedPriceDomainOrderRequest
	GetExpectedPrice() *int64
}

type CreateIntlFixedPriceDomainOrderRequest struct {
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// example:
	//
	// 13350500
	ContactId *int64 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// example:
	//
	// appp16.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// 58.00
	ExpectedPrice *int64 `json:"ExpectedPrice,omitempty" xml:"ExpectedPrice,omitempty"`
}

func (s CreateIntlFixedPriceDomainOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIntlFixedPriceDomainOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateIntlFixedPriceDomainOrderRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *CreateIntlFixedPriceDomainOrderRequest) GetContactId() *int64 {
	return s.ContactId
}

func (s *CreateIntlFixedPriceDomainOrderRequest) GetDomain() *string {
	return s.Domain
}

func (s *CreateIntlFixedPriceDomainOrderRequest) GetExpectedPrice() *int64 {
	return s.ExpectedPrice
}

func (s *CreateIntlFixedPriceDomainOrderRequest) SetAutoPay(v bool) *CreateIntlFixedPriceDomainOrderRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateIntlFixedPriceDomainOrderRequest) SetContactId(v int64) *CreateIntlFixedPriceDomainOrderRequest {
	s.ContactId = &v
	return s
}

func (s *CreateIntlFixedPriceDomainOrderRequest) SetDomain(v string) *CreateIntlFixedPriceDomainOrderRequest {
	s.Domain = &v
	return s
}

func (s *CreateIntlFixedPriceDomainOrderRequest) SetExpectedPrice(v int64) *CreateIntlFixedPriceDomainOrderRequest {
	s.ExpectedPrice = &v
	return s
}

func (s *CreateIntlFixedPriceDomainOrderRequest) Validate() error {
	return dara.Validate(s)
}
