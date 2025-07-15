// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIPv6TranslatorBandwidthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *ModifyIPv6TranslatorBandwidthResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ModifyIPv6TranslatorBandwidthResponseBody
	GetRequestId() *string
}

type ModifyIPv6TranslatorBandwidthResponseBody struct {
	// The order ID.
	//
	// example:
	//
	// 202304500950739
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EF8198EE-8FC9-49C2-A22E-xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyIPv6TranslatorBandwidthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyIPv6TranslatorBandwidthResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyIPv6TranslatorBandwidthResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ModifyIPv6TranslatorBandwidthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyIPv6TranslatorBandwidthResponseBody) SetOrderId(v string) *ModifyIPv6TranslatorBandwidthResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyIPv6TranslatorBandwidthResponseBody) SetRequestId(v string) *ModifyIPv6TranslatorBandwidthResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyIPv6TranslatorBandwidthResponseBody) Validate() error {
	return dara.Validate(s)
}
