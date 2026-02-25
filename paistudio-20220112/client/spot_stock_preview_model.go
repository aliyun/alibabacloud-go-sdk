// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSpotStockPreview interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableQuantity(v int32) *SpotStockPreview
	GetAvailableQuantity() *int32
	SetClusterId(v string) *SpotStockPreview
	GetClusterId() *string
	SetInstanceType(v string) *SpotStockPreview
	GetInstanceType() *string
	SetSpotDiscount(v float32) *SpotStockPreview
	GetSpotDiscount() *float32
	SetStockStatus(v string) *SpotStockPreview
	GetStockStatus() *string
}

type SpotStockPreview struct {
	AvailableQuantity *int32  `json:"AvailableQuantity,omitempty" xml:"AvailableQuantity,omitempty"`
	ClusterId         *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
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

func (s *SpotStockPreview) GetAvailableQuantity() *int32 {
	return s.AvailableQuantity
}

func (s *SpotStockPreview) GetClusterId() *string {
	return s.ClusterId
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

func (s *SpotStockPreview) SetAvailableQuantity(v int32) *SpotStockPreview {
	s.AvailableQuantity = &v
	return s
}

func (s *SpotStockPreview) SetClusterId(v string) *SpotStockPreview {
	s.ClusterId = &v
	return s
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
