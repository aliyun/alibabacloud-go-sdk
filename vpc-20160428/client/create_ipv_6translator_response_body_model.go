// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIPv6TranslatorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIpv6TranslatorId(v string) *CreateIPv6TranslatorResponseBody
	GetIpv6TranslatorId() *string
	SetName(v string) *CreateIPv6TranslatorResponseBody
	GetName() *string
	SetOrderId(v int64) *CreateIPv6TranslatorResponseBody
	GetOrderId() *int64
	SetRequestId(v string) *CreateIPv6TranslatorResponseBody
	GetRequestId() *string
	SetSpec(v string) *CreateIPv6TranslatorResponseBody
	GetSpec() *string
}

type CreateIPv6TranslatorResponseBody struct {
	// The ID of the IPv6 Translation Service instance.
	//
	// example:
	//
	// ipv6trans-bp1i8ahxut1xxxx
	Ipv6TranslatorId *string `json:"Ipv6TranslatorId,omitempty" xml:"Ipv6TranslatorId,omitempty"`
	// The name of the IPv6 Translation Service instance.
	//
	// example:
	//
	// test_nat64gw
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 202303300940739
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AE05898-06E5-4782-xxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The specification of the IPv6 Translation Service instance.
	//
	// example:
	//
	// small
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s CreateIPv6TranslatorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIPv6TranslatorResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIPv6TranslatorResponseBody) GetIpv6TranslatorId() *string {
	return s.Ipv6TranslatorId
}

func (s *CreateIPv6TranslatorResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateIPv6TranslatorResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *CreateIPv6TranslatorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIPv6TranslatorResponseBody) GetSpec() *string {
	return s.Spec
}

func (s *CreateIPv6TranslatorResponseBody) SetIpv6TranslatorId(v string) *CreateIPv6TranslatorResponseBody {
	s.Ipv6TranslatorId = &v
	return s
}

func (s *CreateIPv6TranslatorResponseBody) SetName(v string) *CreateIPv6TranslatorResponseBody {
	s.Name = &v
	return s
}

func (s *CreateIPv6TranslatorResponseBody) SetOrderId(v int64) *CreateIPv6TranslatorResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateIPv6TranslatorResponseBody) SetRequestId(v string) *CreateIPv6TranslatorResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIPv6TranslatorResponseBody) SetSpec(v string) *CreateIPv6TranslatorResponseBody {
	s.Spec = &v
	return s
}

func (s *CreateIPv6TranslatorResponseBody) Validate() error {
	return dara.Validate(s)
}
