// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiMarketAttributesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v string) *DescribeApiMarketAttributesResponseBody
	GetApiId() *string
	SetMarketChargingMode(v string) *DescribeApiMarketAttributesResponseBody
	GetMarketChargingMode() *string
	SetNeedCharging(v string) *DescribeApiMarketAttributesResponseBody
	GetNeedCharging() *string
	SetRequestId(v string) *DescribeApiMarketAttributesResponseBody
	GetRequestId() *string
}

type DescribeApiMarketAttributesResponseBody struct {
	// The ID of the API.
	//
	// example:
	//
	// 6318cd8f6a304cac9318dea8d9a78f7a
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The billing method used by the Alibaba Cloud Marketplace.
	//
	// example:
	//
	// PREPAID_BY_USAGE
	MarketChargingMode *string `json:"MarketChargingMode,omitempty" xml:"MarketChargingMode,omitempty"`
	// Indicates whether fees are charged.
	//
	// example:
	//
	// true
	NeedCharging *string `json:"NeedCharging,omitempty" xml:"NeedCharging,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 545D4E52-4F77-5EC4-BB7E-7344CEC7B5E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeApiMarketAttributesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiMarketAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiMarketAttributesResponseBody) GetApiId() *string {
	return s.ApiId
}

func (s *DescribeApiMarketAttributesResponseBody) GetMarketChargingMode() *string {
	return s.MarketChargingMode
}

func (s *DescribeApiMarketAttributesResponseBody) GetNeedCharging() *string {
	return s.NeedCharging
}

func (s *DescribeApiMarketAttributesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApiMarketAttributesResponseBody) SetApiId(v string) *DescribeApiMarketAttributesResponseBody {
	s.ApiId = &v
	return s
}

func (s *DescribeApiMarketAttributesResponseBody) SetMarketChargingMode(v string) *DescribeApiMarketAttributesResponseBody {
	s.MarketChargingMode = &v
	return s
}

func (s *DescribeApiMarketAttributesResponseBody) SetNeedCharging(v string) *DescribeApiMarketAttributesResponseBody {
	s.NeedCharging = &v
	return s
}

func (s *DescribeApiMarketAttributesResponseBody) SetRequestId(v string) *DescribeApiMarketAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiMarketAttributesResponseBody) Validate() error {
	return dara.Validate(s)
}
