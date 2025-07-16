// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnOrderCommodityCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderCommodityCode(v string) *DescribeCdnOrderCommodityCodeResponseBody
	GetOrderCommodityCode() *string
	SetRequestId(v string) *DescribeCdnOrderCommodityCodeResponseBody
	GetRequestId() *string
}

type DescribeCdnOrderCommodityCodeResponseBody struct {
	// The commodity code that includes the organization unit.
	//
	// example:
	//
	// xxx
	OrderCommodityCode *string `json:"OrderCommodityCode,omitempty" xml:"OrderCommodityCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// BFFCDFAD-DACC-484E-9BE6-0AF3B3A0DD23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdnOrderCommodityCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnOrderCommodityCodeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnOrderCommodityCodeResponseBody) GetOrderCommodityCode() *string {
	return s.OrderCommodityCode
}

func (s *DescribeCdnOrderCommodityCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnOrderCommodityCodeResponseBody) SetOrderCommodityCode(v string) *DescribeCdnOrderCommodityCodeResponseBody {
	s.OrderCommodityCode = &v
	return s
}

func (s *DescribeCdnOrderCommodityCodeResponseBody) SetRequestId(v string) *DescribeCdnOrderCommodityCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnOrderCommodityCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
