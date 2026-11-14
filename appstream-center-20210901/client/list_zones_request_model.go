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
	// The region ID. Required. Specifies the region for which you want to query available zones. All returned zones are located in this region.
	//
	// The value must be a region ID supported by App Streaming. You can call [ListRegions](~~ListRegions~~) to query the supported region IDs. If you specify an unsupported region, the error code InvalidParameter.ValueInvalid is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The operating system type. Required. Specifies the operating system used by the resource. This parameter and ProductType together determine the available zones. The value is case-insensitive. Use one of the following values:
	//
	// Valid values:
	//
	// - Windows: Windows operating system
	//
	// - Linux: Linux operating system
	//
	// - Android: Android operating system
	//
	// This parameter is required.
	//
	// example:
	//
	// Windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The product type. Required. Specifies the product for which you want to query available zones. The list of available zones is returned based on the available resources of the product in the specified region. The value is case-insensitive. Use one of the following values:
	//
	// Valid values:
	//
	// - CloudApp: App Streaming
	//
	// - CloudBrowser: Cloud Browser
	//
	// - WuyingServer: Enterprise Edition workstation
	//
	// - WuyingWorkstation: Personal Edition Linggou container workstation
	//
	// - WuyingWorkstationTeam: Team Edition Linggou container workstation
	//
	// - WuyingWorkstationBusiness: Dedicated Edition Linggou container workstation
	//
	// - AndroidCloud: Cloud Phone
	//
	// - AIAgent: AgentBay (AI agent)
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
