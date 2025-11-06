// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDomainRealTimePriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrency(v string) *QueryDomainRealTimePriceRequest
	GetCurrency() *string
	SetDomainItem(v []*QueryDomainRealTimePriceRequestDomainItem) *QueryDomainRealTimePriceRequest
	GetDomainItem() []*QueryDomainRealTimePriceRequestDomainItem
}

type QueryDomainRealTimePriceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// This parameter is required.
	DomainItem []*QueryDomainRealTimePriceRequestDomainItem `json:"DomainItem,omitempty" xml:"DomainItem,omitempty" type:"Repeated"`
}

func (s QueryDomainRealTimePriceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainRealTimePriceRequest) GoString() string {
	return s.String()
}

func (s *QueryDomainRealTimePriceRequest) GetCurrency() *string {
	return s.Currency
}

func (s *QueryDomainRealTimePriceRequest) GetDomainItem() []*QueryDomainRealTimePriceRequestDomainItem {
	return s.DomainItem
}

func (s *QueryDomainRealTimePriceRequest) SetCurrency(v string) *QueryDomainRealTimePriceRequest {
	s.Currency = &v
	return s
}

func (s *QueryDomainRealTimePriceRequest) SetDomainItem(v []*QueryDomainRealTimePriceRequestDomainItem) *QueryDomainRealTimePriceRequest {
	s.DomainItem = v
	return s
}

func (s *QueryDomainRealTimePriceRequest) Validate() error {
	if s.DomainItem != nil {
		for _, item := range s.DomainItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryDomainRealTimePriceRequestDomainItem struct {
	// This parameter is required.
	//
	// example:
	//
	// renew
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// com
	Suffix *string `json:"Suffix,omitempty" xml:"Suffix,omitempty"`
}

func (s QueryDomainRealTimePriceRequestDomainItem) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainRealTimePriceRequestDomainItem) GoString() string {
	return s.String()
}

func (s *QueryDomainRealTimePriceRequestDomainItem) GetAction() *string {
	return s.Action
}

func (s *QueryDomainRealTimePriceRequestDomainItem) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryDomainRealTimePriceRequestDomainItem) GetPeriod() *int32 {
	return s.Period
}

func (s *QueryDomainRealTimePriceRequestDomainItem) GetSuffix() *string {
	return s.Suffix
}

func (s *QueryDomainRealTimePriceRequestDomainItem) SetAction(v string) *QueryDomainRealTimePriceRequestDomainItem {
	s.Action = &v
	return s
}

func (s *QueryDomainRealTimePriceRequestDomainItem) SetDomainName(v string) *QueryDomainRealTimePriceRequestDomainItem {
	s.DomainName = &v
	return s
}

func (s *QueryDomainRealTimePriceRequestDomainItem) SetPeriod(v int32) *QueryDomainRealTimePriceRequestDomainItem {
	s.Period = &v
	return s
}

func (s *QueryDomainRealTimePriceRequestDomainItem) SetSuffix(v string) *QueryDomainRealTimePriceRequestDomainItem {
	s.Suffix = &v
	return s
}

func (s *QueryDomainRealTimePriceRequestDomainItem) Validate() error {
	return dara.Validate(s)
}
