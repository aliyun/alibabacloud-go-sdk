// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSpotStockInternalInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableQuantity(v int32) *SpotStockInternalInfo
	GetAvailableQuantity() *int32
	SetClusterId(v string) *SpotStockInternalInfo
	GetClusterId() *string
	SetHpnZone(v string) *SpotStockInternalInfo
	GetHpnZone() *string
	SetTotalQuantity(v int32) *SpotStockInternalInfo
	GetTotalQuantity() *int32
}

type SpotStockInternalInfo struct {
	// The number of available Spot Instances.
	AvailableQuantity *int32 `json:"availableQuantity,omitempty" xml:"availableQuantity,omitempty"`
	// The ID of the cluster.
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	// The ID of the high-performance network (HPN) zone.
	HpnZone *string `json:"hpnZone,omitempty" xml:"hpnZone,omitempty"`
	// The total number of Spot Instances.
	TotalQuantity *int32 `json:"totalQuantity,omitempty" xml:"totalQuantity,omitempty"`
}

func (s SpotStockInternalInfo) String() string {
	return dara.Prettify(s)
}

func (s SpotStockInternalInfo) GoString() string {
	return s.String()
}

func (s *SpotStockInternalInfo) GetAvailableQuantity() *int32 {
	return s.AvailableQuantity
}

func (s *SpotStockInternalInfo) GetClusterId() *string {
	return s.ClusterId
}

func (s *SpotStockInternalInfo) GetHpnZone() *string {
	return s.HpnZone
}

func (s *SpotStockInternalInfo) GetTotalQuantity() *int32 {
	return s.TotalQuantity
}

func (s *SpotStockInternalInfo) SetAvailableQuantity(v int32) *SpotStockInternalInfo {
	s.AvailableQuantity = &v
	return s
}

func (s *SpotStockInternalInfo) SetClusterId(v string) *SpotStockInternalInfo {
	s.ClusterId = &v
	return s
}

func (s *SpotStockInternalInfo) SetHpnZone(v string) *SpotStockInternalInfo {
	s.HpnZone = &v
	return s
}

func (s *SpotStockInternalInfo) SetTotalQuantity(v int32) *SpotStockInternalInfo {
	s.TotalQuantity = &v
	return s
}

func (s *SpotStockInternalInfo) Validate() error {
	return dara.Validate(s)
}
