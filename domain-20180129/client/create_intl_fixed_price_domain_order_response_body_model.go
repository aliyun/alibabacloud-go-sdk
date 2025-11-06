// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIntlFixedPriceDomainOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModule(v *CreateIntlFixedPriceDomainOrderResponseBodyModule) *CreateIntlFixedPriceDomainOrderResponseBody
	GetModule() *CreateIntlFixedPriceDomainOrderResponseBodyModule
	SetRequestId(v string) *CreateIntlFixedPriceDomainOrderResponseBody
	GetRequestId() *string
}

type CreateIntlFixedPriceDomainOrderResponseBody struct {
	Module *CreateIntlFixedPriceDomainOrderResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// example:
	//
	// E879DC07-38EE-4408-9F33-73B30CD965CD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIntlFixedPriceDomainOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIntlFixedPriceDomainOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIntlFixedPriceDomainOrderResponseBody) GetModule() *CreateIntlFixedPriceDomainOrderResponseBodyModule {
	return s.Module
}

func (s *CreateIntlFixedPriceDomainOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIntlFixedPriceDomainOrderResponseBody) SetModule(v *CreateIntlFixedPriceDomainOrderResponseBodyModule) *CreateIntlFixedPriceDomainOrderResponseBody {
	s.Module = v
	return s
}

func (s *CreateIntlFixedPriceDomainOrderResponseBody) SetRequestId(v string) *CreateIntlFixedPriceDomainOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIntlFixedPriceDomainOrderResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateIntlFixedPriceDomainOrderResponseBodyModule struct {
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// 31199295f2074ce895645d386cb2****
	OrderNo *string `json:"OrderNo,omitempty" xml:"OrderNo,omitempty"`
	// example:
	//
	// 100.00
	PayPrice *int64 `json:"PayPrice,omitempty" xml:"PayPrice,omitempty"`
	// example:
	//
	// https://
	PayUrl *string `json:"PayUrl,omitempty" xml:"PayUrl,omitempty"`
}

func (s CreateIntlFixedPriceDomainOrderResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s CreateIntlFixedPriceDomainOrderResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CreateIntlFixedPriceDomainOrderResponseBodyModule) GetDomain() *string {
	return s.Domain
}

func (s *CreateIntlFixedPriceDomainOrderResponseBodyModule) GetOrderNo() *string {
	return s.OrderNo
}

func (s *CreateIntlFixedPriceDomainOrderResponseBodyModule) GetPayPrice() *int64 {
	return s.PayPrice
}

func (s *CreateIntlFixedPriceDomainOrderResponseBodyModule) GetPayUrl() *string {
	return s.PayUrl
}

func (s *CreateIntlFixedPriceDomainOrderResponseBodyModule) SetDomain(v string) *CreateIntlFixedPriceDomainOrderResponseBodyModule {
	s.Domain = &v
	return s
}

func (s *CreateIntlFixedPriceDomainOrderResponseBodyModule) SetOrderNo(v string) *CreateIntlFixedPriceDomainOrderResponseBodyModule {
	s.OrderNo = &v
	return s
}

func (s *CreateIntlFixedPriceDomainOrderResponseBodyModule) SetPayPrice(v int64) *CreateIntlFixedPriceDomainOrderResponseBodyModule {
	s.PayPrice = &v
	return s
}

func (s *CreateIntlFixedPriceDomainOrderResponseBodyModule) SetPayUrl(v string) *CreateIntlFixedPriceDomainOrderResponseBodyModule {
	s.PayUrl = &v
	return s
}

func (s *CreateIntlFixedPriceDomainOrderResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
