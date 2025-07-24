// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSpotBidPrice interface {
	dara.Model
	String() string
	GoString() string
	SetBidPrice(v float64) *SpotBidPrice
	GetBidPrice() *float64
	SetInstanceType(v string) *SpotBidPrice
	GetInstanceType() *string
}

type SpotBidPrice struct {
	// 实例的每小时最高出价。支持最大3位小数，参数SpotStrategy=SpotWithPriceLimit时，该参数生效。
	//
	// example:
	//
	// 1000.0
	BidPrice *float64 `json:"BidPrice,omitempty" xml:"BidPrice,omitempty"`
	// 实例类型。
	//
	// example:
	//
	// ecs.g7.2xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s SpotBidPrice) String() string {
	return dara.Prettify(s)
}

func (s SpotBidPrice) GoString() string {
	return s.String()
}

func (s *SpotBidPrice) GetBidPrice() *float64 {
	return s.BidPrice
}

func (s *SpotBidPrice) GetInstanceType() *string {
	return s.InstanceType
}

func (s *SpotBidPrice) SetBidPrice(v float64) *SpotBidPrice {
	s.BidPrice = &v
	return s
}

func (s *SpotBidPrice) SetInstanceType(v string) *SpotBidPrice {
	s.InstanceType = &v
	return s
}

func (s *SpotBidPrice) Validate() error {
	return dara.Validate(s)
}
