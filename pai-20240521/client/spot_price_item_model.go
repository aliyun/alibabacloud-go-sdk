// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSpotPriceItem interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceType(v string) *SpotPriceItem
	GetInstanceType() *string
	SetSpotDiscount(v float32) *SpotPriceItem
	GetSpotDiscount() *float32
	SetTimestamp(v string) *SpotPriceItem
	GetTimestamp() *string
	SetZoneId(v string) *SpotPriceItem
	GetZoneId() *string
}

type SpotPriceItem struct {
	// example:
	//
	// ml.gu8xf.8xlarge-gu108
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// 0.1
	SpotDiscount *float32 `json:"SpotDiscount,omitempty" xml:"SpotDiscount,omitempty"`
	// example:
	//
	// 2024-01-17T06:00:00Z
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// example:
	//
	// cn-wulanchabu-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s SpotPriceItem) String() string {
	return dara.Prettify(s)
}

func (s SpotPriceItem) GoString() string {
	return s.String()
}

func (s *SpotPriceItem) GetInstanceType() *string {
	return s.InstanceType
}

func (s *SpotPriceItem) GetSpotDiscount() *float32 {
	return s.SpotDiscount
}

func (s *SpotPriceItem) GetTimestamp() *string {
	return s.Timestamp
}

func (s *SpotPriceItem) GetZoneId() *string {
	return s.ZoneId
}

func (s *SpotPriceItem) SetInstanceType(v string) *SpotPriceItem {
	s.InstanceType = &v
	return s
}

func (s *SpotPriceItem) SetSpotDiscount(v float32) *SpotPriceItem {
	s.SpotDiscount = &v
	return s
}

func (s *SpotPriceItem) SetTimestamp(v string) *SpotPriceItem {
	s.Timestamp = &v
	return s
}

func (s *SpotPriceItem) SetZoneId(v string) *SpotPriceItem {
	s.ZoneId = &v
	return s
}

func (s *SpotPriceItem) Validate() error {
	return dara.Validate(s)
}
