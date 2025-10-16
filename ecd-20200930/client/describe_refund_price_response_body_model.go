// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRefundPriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPriceInfo(v *DescribeRefundPriceResponseBodyPriceInfo) *DescribeRefundPriceResponseBody
	GetPriceInfo() *DescribeRefundPriceResponseBodyPriceInfo
	SetRequestId(v string) *DescribeRefundPriceResponseBody
	GetRequestId() *string
}

type DescribeRefundPriceResponseBody struct {
	// The price details.
	PriceInfo *DescribeRefundPriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRefundPriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRefundPriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRefundPriceResponseBody) GetPriceInfo() *DescribeRefundPriceResponseBodyPriceInfo {
	return s.PriceInfo
}

func (s *DescribeRefundPriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRefundPriceResponseBody) SetPriceInfo(v *DescribeRefundPriceResponseBodyPriceInfo) *DescribeRefundPriceResponseBody {
	s.PriceInfo = v
	return s
}

func (s *DescribeRefundPriceResponseBody) SetRequestId(v string) *DescribeRefundPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRefundPriceResponseBody) Validate() error {
	if s.PriceInfo != nil {
		if err := s.PriceInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRefundPriceResponseBodyPriceInfo struct {
	// The unit of currency (USD).
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The amount of the refund.
	//
	// example:
	//
	// 3990.75
	RefundFee *float32 `json:"RefundFee,omitempty" xml:"RefundFee,omitempty"`
}

func (s DescribeRefundPriceResponseBodyPriceInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeRefundPriceResponseBodyPriceInfo) GoString() string {
	return s.String()
}

func (s *DescribeRefundPriceResponseBodyPriceInfo) GetCurrency() *string {
	return s.Currency
}

func (s *DescribeRefundPriceResponseBodyPriceInfo) GetRefundFee() *float32 {
	return s.RefundFee
}

func (s *DescribeRefundPriceResponseBodyPriceInfo) SetCurrency(v string) *DescribeRefundPriceResponseBodyPriceInfo {
	s.Currency = &v
	return s
}

func (s *DescribeRefundPriceResponseBodyPriceInfo) SetRefundFee(v float32) *DescribeRefundPriceResponseBodyPriceInfo {
	s.RefundFee = &v
	return s
}

func (s *DescribeRefundPriceResponseBodyPriceInfo) Validate() error {
	return dara.Validate(s)
}
