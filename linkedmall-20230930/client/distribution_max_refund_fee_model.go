// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDistributionMaxRefundFee interface {
	dara.Model
	String() string
	GoString() string
	SetMaxRefundFee(v int64) *DistributionMaxRefundFee
	GetMaxRefundFee() *int64
	SetMinRefundFee(v int64) *DistributionMaxRefundFee
	GetMinRefundFee() *int64
}

type DistributionMaxRefundFee struct {
	// example:
	//
	// 100
	MaxRefundFee *int64 `json:"maxRefundFee,omitempty" xml:"maxRefundFee,omitempty"`
	// example:
	//
	// 1
	MinRefundFee *int64 `json:"minRefundFee,omitempty" xml:"minRefundFee,omitempty"`
}

func (s DistributionMaxRefundFee) String() string {
	return dara.Prettify(s)
}

func (s DistributionMaxRefundFee) GoString() string {
	return s.String()
}

func (s *DistributionMaxRefundFee) GetMaxRefundFee() *int64 {
	return s.MaxRefundFee
}

func (s *DistributionMaxRefundFee) GetMinRefundFee() *int64 {
	return s.MinRefundFee
}

func (s *DistributionMaxRefundFee) SetMaxRefundFee(v int64) *DistributionMaxRefundFee {
	s.MaxRefundFee = &v
	return s
}

func (s *DistributionMaxRefundFee) SetMinRefundFee(v int64) *DistributionMaxRefundFee {
	s.MinRefundFee = &v
	return s
}

func (s *DistributionMaxRefundFee) Validate() error {
	return dara.Validate(s)
}
