// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSpotSpec interface {
	dara.Model
	String() string
	GoString() string
	SetSpotDiscountLimit(v float32) *SpotSpec
	GetSpotDiscountLimit() *float32
	SetSpotPriceLimit(v float32) *SpotSpec
	GetSpotPriceLimit() *float32
	SetSpotStrategy(v string) *SpotSpec
	GetSpotStrategy() *string
}

type SpotSpec struct {
	// The maximum discount. Specify only one of SpotDiscountLimit and SpotPriceLimit.
	//
	// example:
	//
	// 0.8
	SpotDiscountLimit *float32 `json:"SpotDiscountLimit,omitempty" xml:"SpotDiscountLimit,omitempty"`
	// The maximum price. Unit: CNY/minute. Specify only one of SpotDiscountLimit and SpotPriceLimit.
	//
	// example:
	//
	// 0.4744
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// The spot policy. Valid values:
	//
	// 	- SpotWithPriceLimit
	//
	// example:
	//
	// SpotWithPriceLimit
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
}

func (s SpotSpec) String() string {
	return dara.Prettify(s)
}

func (s SpotSpec) GoString() string {
	return s.String()
}

func (s *SpotSpec) GetSpotDiscountLimit() *float32 {
	return s.SpotDiscountLimit
}

func (s *SpotSpec) GetSpotPriceLimit() *float32 {
	return s.SpotPriceLimit
}

func (s *SpotSpec) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *SpotSpec) SetSpotDiscountLimit(v float32) *SpotSpec {
	s.SpotDiscountLimit = &v
	return s
}

func (s *SpotSpec) SetSpotPriceLimit(v float32) *SpotSpec {
	s.SpotPriceLimit = &v
	return s
}

func (s *SpotSpec) SetSpotStrategy(v string) *SpotSpec {
	s.SpotStrategy = &v
	return s
}

func (s *SpotSpec) Validate() error {
	return dara.Validate(s)
}
