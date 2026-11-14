// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRegionIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductType(v string) *ListRegionIdRequest
	GetProductType() *string
}

type ListRegionIdRequest struct {
	// The product type for querying available image regions. If this parameter is omitted, the query is performed based on CloudApp. Valid values:
	//
	// - **CloudGame**: cloud gaming
	//
	// - **CloudApp**: WUYING cloud application
	//
	// - **CloudBrowser**: secure browser
	//
	// - **CloudDesktop**: WUYING cloud desktop
	//
	// - **ImageBuilder**: image creation resource
	//
	// - **ImageValidate**: image validation resource
	//
	// - **AndroidCloud**: Android cloud
	//
	// - **GpuRender**: GPU rendering
	//
	// - **AndroidGpuRender**: Android cloud GPU rendering
	//
	// - **WuyingServer**: enterprise workstation
	//
	// - **GpuProxy**: GPU remote rendering proxy
	//
	// - **GpuCache**: GPU remote rendering cache
	//
	// - **AIAgent**: AI agent service
	//
	// - **HostedDesktop**: hosted computer
	//
	// - **WuyingWorkstation**: personal container workstation
	//
	// - **WuyingWorkstationBusiness**: business container workstation
	//
	// - **WuyingWorkstationTeam**: LingGou container workstation team edition
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s ListRegionIdRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRegionIdRequest) GoString() string {
	return s.String()
}

func (s *ListRegionIdRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ListRegionIdRequest) SetProductType(v string) *ListRegionIdRequest {
	s.ProductType = &v
	return s
}

func (s *ListRegionIdRequest) Validate() error {
	return dara.Validate(s)
}
