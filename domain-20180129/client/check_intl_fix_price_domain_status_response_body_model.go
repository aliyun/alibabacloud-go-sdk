// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckIntlFixPriceDomainStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModule(v *CheckIntlFixPriceDomainStatusResponseBodyModule) *CheckIntlFixPriceDomainStatusResponseBody
	GetModule() *CheckIntlFixPriceDomainStatusResponseBodyModule
	SetRequestId(v string) *CheckIntlFixPriceDomainStatusResponseBody
	GetRequestId() *string
}

type CheckIntlFixPriceDomainStatusResponseBody struct {
	// The returned object.
	Module *CheckIntlFixPriceDomainStatusResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 40F46D3D-F4F3-4CCB-AC30-2DD20E32E528
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckIntlFixPriceDomainStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckIntlFixPriceDomainStatusResponseBody) GoString() string {
	return s.String()
}

func (s *CheckIntlFixPriceDomainStatusResponseBody) GetModule() *CheckIntlFixPriceDomainStatusResponseBodyModule {
	return s.Module
}

func (s *CheckIntlFixPriceDomainStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckIntlFixPriceDomainStatusResponseBody) SetModule(v *CheckIntlFixPriceDomainStatusResponseBodyModule) *CheckIntlFixPriceDomainStatusResponseBody {
	s.Module = v
	return s
}

func (s *CheckIntlFixPriceDomainStatusResponseBody) SetRequestId(v string) *CheckIntlFixPriceDomainStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckIntlFixPriceDomainStatusResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CheckIntlFixPriceDomainStatusResponseBodyModule struct {
	// The currency. Valid values:
	//
	// - RMB: Chinese Yuan.
	//
	// - USD: US Dollar.
	//
	// example:
	//
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The expiration date of the domain name. After this date, the domain name requires renewal.
	//
	// example:
	//
	// 1567353497
	DeadDate *int64 `json:"DeadDate,omitempty" xml:"DeadDate,omitempty"`
	// The domain name.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The sale deadline of the domain name. After this time, the domain name is no longer available for sale.
	//
	// example:
	//
	// 1567353497
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Indicates whether the domain name is a premium domain name. Valid values:
	//
	// - true: The domain name is a premium domain name.
	//
	// - false: The domain name is not a premium domain name.
	//
	// example:
	//
	// true
	Premium *bool `json:"Premium,omitempty" xml:"Premium,omitempty"`
	// The price.
	//
	// example:
	//
	// 20.00
	Price *int64 `json:"Price,omitempty" xml:"Price,omitempty"`
	// The registration date of the domain name.
	//
	// example:
	//
	// 1566353497
	RegDate *int64 `json:"RegDate,omitempty" xml:"RegDate,omitempty"`
}

func (s CheckIntlFixPriceDomainStatusResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s CheckIntlFixPriceDomainStatusResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CheckIntlFixPriceDomainStatusResponseBodyModule) GetCurrency() *string {
	return s.Currency
}

func (s *CheckIntlFixPriceDomainStatusResponseBodyModule) GetDeadDate() *int64 {
	return s.DeadDate
}

func (s *CheckIntlFixPriceDomainStatusResponseBodyModule) GetDomain() *string {
	return s.Domain
}

func (s *CheckIntlFixPriceDomainStatusResponseBodyModule) GetEndTime() *int64 {
	return s.EndTime
}

func (s *CheckIntlFixPriceDomainStatusResponseBodyModule) GetPremium() *bool {
	return s.Premium
}

func (s *CheckIntlFixPriceDomainStatusResponseBodyModule) GetPrice() *int64 {
	return s.Price
}

func (s *CheckIntlFixPriceDomainStatusResponseBodyModule) GetRegDate() *int64 {
	return s.RegDate
}

func (s *CheckIntlFixPriceDomainStatusResponseBodyModule) SetCurrency(v string) *CheckIntlFixPriceDomainStatusResponseBodyModule {
	s.Currency = &v
	return s
}

func (s *CheckIntlFixPriceDomainStatusResponseBodyModule) SetDeadDate(v int64) *CheckIntlFixPriceDomainStatusResponseBodyModule {
	s.DeadDate = &v
	return s
}

func (s *CheckIntlFixPriceDomainStatusResponseBodyModule) SetDomain(v string) *CheckIntlFixPriceDomainStatusResponseBodyModule {
	s.Domain = &v
	return s
}

func (s *CheckIntlFixPriceDomainStatusResponseBodyModule) SetEndTime(v int64) *CheckIntlFixPriceDomainStatusResponseBodyModule {
	s.EndTime = &v
	return s
}

func (s *CheckIntlFixPriceDomainStatusResponseBodyModule) SetPremium(v bool) *CheckIntlFixPriceDomainStatusResponseBodyModule {
	s.Premium = &v
	return s
}

func (s *CheckIntlFixPriceDomainStatusResponseBodyModule) SetPrice(v int64) *CheckIntlFixPriceDomainStatusResponseBodyModule {
	s.Price = &v
	return s
}

func (s *CheckIntlFixPriceDomainStatusResponseBodyModule) SetRegDate(v int64) *CheckIntlFixPriceDomainStatusResponseBodyModule {
	s.RegDate = &v
	return s
}

func (s *CheckIntlFixPriceDomainStatusResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
