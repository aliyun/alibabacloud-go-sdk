// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSpotStockPreview interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceType(v string) *SpotStockPreview
	GetInstanceType() *string
	SetSpotDiscount(v float32) *SpotStockPreview
	GetSpotDiscount() *float32
	SetStockStatus(v string) *SpotStockPreview
	GetStockStatus() *string
}

type SpotStockPreview struct {
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
	// WithStock
	StockStatus *string `json:"StockStatus,omitempty" xml:"StockStatus,omitempty"`
}

func (s SpotStockPreview) String() string {
	return dara.Prettify(s)
}

func (s SpotStockPreview) GoString() string {
	return s.String()
}

func (s *SpotStockPreview) GetInstanceType() *string {
	return s.InstanceType
}

func (s *SpotStockPreview) GetSpotDiscount() *float32 {
	return s.SpotDiscount
}

func (s *SpotStockPreview) GetStockStatus() *string {
	return s.StockStatus
}

func (s *SpotStockPreview) SetInstanceType(v string) *SpotStockPreview {
	s.InstanceType = &v
	return s
}

func (s *SpotStockPreview) SetSpotDiscount(v float32) *SpotStockPreview {
	s.SpotDiscount = &v
	return s
}

func (s *SpotStockPreview) SetStockStatus(v string) *SpotStockPreview {
	s.StockStatus = &v
	return s
}

func (s *SpotStockPreview) Validate() error {
	return dara.Validate(s)
}
