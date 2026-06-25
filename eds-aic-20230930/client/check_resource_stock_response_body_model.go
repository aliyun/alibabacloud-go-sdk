// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckResourceStockResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckResourceStockResponseBody
	GetRequestId() *string
	SetResourceStockModels(v []*CheckResourceStockResponseBodyResourceStockModels) *CheckResourceStockResponseBody
	GetResourceStockModels() []*CheckResourceStockResponseBodyResourceStockModels
}

type CheckResourceStockResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 805D8FB6-512A-531C-9E4D-2A807D3C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource inventory details.
	ResourceStockModels []*CheckResourceStockResponseBodyResourceStockModels `json:"ResourceStockModels,omitempty" xml:"ResourceStockModels,omitempty" type:"Repeated"`
}

func (s CheckResourceStockResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckResourceStockResponseBody) GoString() string {
	return s.String()
}

func (s *CheckResourceStockResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckResourceStockResponseBody) GetResourceStockModels() []*CheckResourceStockResponseBodyResourceStockModels {
	return s.ResourceStockModels
}

func (s *CheckResourceStockResponseBody) SetRequestId(v string) *CheckResourceStockResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckResourceStockResponseBody) SetResourceStockModels(v []*CheckResourceStockResponseBodyResourceStockModels) *CheckResourceStockResponseBody {
	s.ResourceStockModels = v
	return s
}

func (s *CheckResourceStockResponseBody) Validate() error {
	if s.ResourceStockModels != nil {
		for _, item := range s.ResourceStockModels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CheckResourceStockResponseBodyResourceStockModels struct {
	// The ID of the region.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The inventory status of the instance group.
	//
	// example:
	//
	// Available
	StockStatus *string `json:"StockStatus,omitempty" xml:"StockStatus,omitempty"`
	// The ID of the zone.
	//
	// example:
	//
	// cn-shanghai-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CheckResourceStockResponseBodyResourceStockModels) String() string {
	return dara.Prettify(s)
}

func (s CheckResourceStockResponseBodyResourceStockModels) GoString() string {
	return s.String()
}

func (s *CheckResourceStockResponseBodyResourceStockModels) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckResourceStockResponseBodyResourceStockModels) GetStockStatus() *string {
	return s.StockStatus
}

func (s *CheckResourceStockResponseBodyResourceStockModels) GetZoneId() *string {
	return s.ZoneId
}

func (s *CheckResourceStockResponseBodyResourceStockModels) SetRegionId(v string) *CheckResourceStockResponseBodyResourceStockModels {
	s.RegionId = &v
	return s
}

func (s *CheckResourceStockResponseBodyResourceStockModels) SetStockStatus(v string) *CheckResourceStockResponseBodyResourceStockModels {
	s.StockStatus = &v
	return s
}

func (s *CheckResourceStockResponseBodyResourceStockModels) SetZoneId(v string) *CheckResourceStockResponseBodyResourceStockModels {
	s.ZoneId = &v
	return s
}

func (s *CheckResourceStockResponseBodyResourceStockModels) Validate() error {
	return dara.Validate(s)
}
