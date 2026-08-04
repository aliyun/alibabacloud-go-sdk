// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSpotStockItem interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceType(v string) *SpotStockItem
	GetInstanceType() *string
	SetInternalInfo(v []*SpotStockInternalInfo) *SpotStockItem
	GetInternalInfo() []*SpotStockInternalInfo
	SetStockStatus(v string) *SpotStockItem
	GetStockStatus() *string
}

type SpotStockItem struct {
	// The instance type.
	InstanceType *string `json:"instanceType,omitempty" xml:"instanceType,omitempty"`
	// Internal information about the stock of the spot instance type.
	InternalInfo []*SpotStockInternalInfo `json:"internalInfo,omitempty" xml:"internalInfo,omitempty" type:"Repeated"`
	// The stock status of the instance type. Valid values are `Available` and `SoldOut`.
	StockStatus *string `json:"stockStatus,omitempty" xml:"stockStatus,omitempty"`
}

func (s SpotStockItem) String() string {
	return dara.Prettify(s)
}

func (s SpotStockItem) GoString() string {
	return s.String()
}

func (s *SpotStockItem) GetInstanceType() *string {
	return s.InstanceType
}

func (s *SpotStockItem) GetInternalInfo() []*SpotStockInternalInfo {
	return s.InternalInfo
}

func (s *SpotStockItem) GetStockStatus() *string {
	return s.StockStatus
}

func (s *SpotStockItem) SetInstanceType(v string) *SpotStockItem {
	s.InstanceType = &v
	return s
}

func (s *SpotStockItem) SetInternalInfo(v []*SpotStockInternalInfo) *SpotStockItem {
	s.InternalInfo = v
	return s
}

func (s *SpotStockItem) SetStockStatus(v string) *SpotStockItem {
	s.StockStatus = &v
	return s
}

func (s *SpotStockItem) Validate() error {
	if s.InternalInfo != nil {
		for _, item := range s.InternalInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
