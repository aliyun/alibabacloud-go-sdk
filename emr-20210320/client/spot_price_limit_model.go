// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSpotPriceLimit interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceType(v string) *SpotPriceLimit
	GetInstanceType() *string
	SetPriceLimit(v float64) *SpotPriceLimit
	GetPriceLimit() *float64
}

type SpotPriceLimit struct {
	InstanceType *string  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	PriceLimit   *float64 `json:"PriceLimit,omitempty" xml:"PriceLimit,omitempty"`
}

func (s SpotPriceLimit) String() string {
	return dara.Prettify(s)
}

func (s SpotPriceLimit) GoString() string {
	return s.String()
}

func (s *SpotPriceLimit) GetInstanceType() *string {
	return s.InstanceType
}

func (s *SpotPriceLimit) GetPriceLimit() *float64 {
	return s.PriceLimit
}

func (s *SpotPriceLimit) SetInstanceType(v string) *SpotPriceLimit {
	s.InstanceType = &v
	return s
}

func (s *SpotPriceLimit) SetPriceLimit(v float64) *SpotPriceLimit {
	s.PriceLimit = &v
	return s
}

func (s *SpotPriceLimit) Validate() error {
	return dara.Validate(s)
}
