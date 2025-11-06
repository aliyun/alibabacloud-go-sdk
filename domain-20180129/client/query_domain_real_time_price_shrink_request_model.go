// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDomainRealTimePriceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrency(v string) *QueryDomainRealTimePriceShrinkRequest
	GetCurrency() *string
	SetDomainItemShrink(v string) *QueryDomainRealTimePriceShrinkRequest
	GetDomainItemShrink() *string
}

type QueryDomainRealTimePriceShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// This parameter is required.
	DomainItemShrink *string `json:"DomainItem,omitempty" xml:"DomainItem,omitempty"`
}

func (s QueryDomainRealTimePriceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainRealTimePriceShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryDomainRealTimePriceShrinkRequest) GetCurrency() *string {
	return s.Currency
}

func (s *QueryDomainRealTimePriceShrinkRequest) GetDomainItemShrink() *string {
	return s.DomainItemShrink
}

func (s *QueryDomainRealTimePriceShrinkRequest) SetCurrency(v string) *QueryDomainRealTimePriceShrinkRequest {
	s.Currency = &v
	return s
}

func (s *QueryDomainRealTimePriceShrinkRequest) SetDomainItemShrink(v string) *QueryDomainRealTimePriceShrinkRequest {
	s.DomainItemShrink = &v
	return s
}

func (s *QueryDomainRealTimePriceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
