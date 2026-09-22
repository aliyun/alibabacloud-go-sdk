// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizRegionId(v string) *ListZonesRequest
	GetBizRegionId() *string
	SetOsType(v string) *ListZonesRequest
	GetOsType() *string
	SetProductType(v string) *ListZonesRequest
	GetProductType() *string
}

type ListZonesRequest struct {
	// The region ID. Required. Specifies the region for which to query available zones. All returned zones are within this region.
	//
	// The value must be a region ID supported by WUYING Cloud Application. Call [ListRegions](~~ListRegions~~) to obtain the supported region IDs. If an unsupported region is specified, the error code `InvalidParameter.ValueInvalid` is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The operating system type. Required. Specifies the operating system used by the resource. This parameter, together with `ProductType`, determines the available zones. The value is case-insensitive. Use the following recommended values.
	//
	// Valid values:
	//
	// - `Windows`: Windows operating system.
	//
	// - `Linux`: Linux operating system.
	//
	// - `Android`: Android operating system.
	//
	// This parameter is required.
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The product type. Required. Specifies the product for which to query available zones. The zone list is returned based on the available resources of this product in the specified region. The value is case-insensitive. Use the following recommended values.
	//
	// Valid values:
	//
	// - `CloudApp`: WUYING Cloud Application.
	//
	// - `CloudBrowser`: Cloud Browser.
	//
	// - `WuyingServer`: Enterprise Edition Workstation.
	//
	// - `WuyingWorkstation`: Personal Edition Lingou Container Workstation.
	//
	// - `WuyingWorkstationTeam`: Lingou Team Edition Container Workstation.
	//
	// - `WuyingWorkstationBusiness`: Lingou Dedicated Edition Container Workstation.
	//
	// - `AndroidCloud`: Cloud Phone.
	//
	// - `AIAgent`: AgentBay (AI agent).
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s ListZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListZonesRequest) GoString() string {
	return s.String()
}

func (s *ListZonesRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *ListZonesRequest) GetOsType() *string {
	return s.OsType
}

func (s *ListZonesRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ListZonesRequest) SetBizRegionId(v string) *ListZonesRequest {
	s.BizRegionId = &v
	return s
}

func (s *ListZonesRequest) SetOsType(v string) *ListZonesRequest {
	s.OsType = &v
	return s
}

func (s *ListZonesRequest) SetProductType(v string) *ListZonesRequest {
	s.ProductType = &v
	return s
}

func (s *ListZonesRequest) Validate() error {
	return dara.Validate(s)
}
