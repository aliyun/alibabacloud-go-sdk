// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommodity(v string) *DescribePriceRequest
	GetCommodity() *string
	SetOrderType(v string) *DescribePriceRequest
	GetOrderType() *string
}

type DescribePriceRequest struct {
	// example:
	//
	// {\\"duration\\": 1, \\"productCode\\": \\"cmjz000325\\", \\"quantity\\": 1, \\"pricingCycle\\": \\"Year\\", \\"skuCode\\": \\"jichuban\\"}
	Commodity *string `json:"Commodity,omitempty" xml:"Commodity,omitempty"`
	// example:
	//
	// DOWNGRADE
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
}

func (s DescribePriceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceRequest) GoString() string {
	return s.String()
}

func (s *DescribePriceRequest) GetCommodity() *string {
	return s.Commodity
}

func (s *DescribePriceRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribePriceRequest) SetCommodity(v string) *DescribePriceRequest {
	s.Commodity = &v
	return s
}

func (s *DescribePriceRequest) SetOrderType(v string) *DescribePriceRequest {
	s.OrderType = &v
	return s
}

func (s *DescribePriceRequest) Validate() error {
	return dara.Validate(s)
}
