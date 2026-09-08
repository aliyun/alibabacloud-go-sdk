// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalSpotPriceItem interface {
	dara.Model
	String() string
	GoString() string
	SetEffectiveAt(v string) *GlobalSpotPriceItem
	GetEffectiveAt() *string
	SetInstanceType(v string) *GlobalSpotPriceItem
	GetInstanceType() *string
	SetSpotDiscount(v string) *GlobalSpotPriceItem
	GetSpotDiscount() *string
}

type GlobalSpotPriceItem struct {
	// The effective period.
	//
	// example:
	//
	// 2026-08-20T06:45:00Z
	EffectiveAt *string `json:"effectiveAt,omitempty" xml:"effectiveAt,omitempty"`
	// The instance type.
	//
	// example:
	//
	// ml.gp7vf.16.40xlarge
	InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	// The current market price.
	//
	// example:
	//
	// 0.1
	SpotDiscount *string `json:"spotDiscount,omitempty" xml:"spotDiscount,omitempty"`
}

func (s GlobalSpotPriceItem) String() string {
	return dara.Prettify(s)
}

func (s GlobalSpotPriceItem) GoString() string {
	return s.String()
}

func (s *GlobalSpotPriceItem) GetEffectiveAt() *string {
	return s.EffectiveAt
}

func (s *GlobalSpotPriceItem) GetInstanceType() *string {
	return s.InstanceType
}

func (s *GlobalSpotPriceItem) GetSpotDiscount() *string {
	return s.SpotDiscount
}

func (s *GlobalSpotPriceItem) SetEffectiveAt(v string) *GlobalSpotPriceItem {
	s.EffectiveAt = &v
	return s
}

func (s *GlobalSpotPriceItem) SetInstanceType(v string) *GlobalSpotPriceItem {
	s.InstanceType = &v
	return s
}

func (s *GlobalSpotPriceItem) SetSpotDiscount(v string) *GlobalSpotPriceItem {
	s.SpotDiscount = &v
	return s
}

func (s *GlobalSpotPriceItem) Validate() error {
	return dara.Validate(s)
}
