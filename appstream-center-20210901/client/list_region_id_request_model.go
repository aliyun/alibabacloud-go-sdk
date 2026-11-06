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
	// The product type for querying available image regions. If this parameter is omitted, the query defaults to CloudApp. Valid values:
	//
	// - **CloudGame**: cloud gaming
	//
	// - **CloudApp**: WUYING Cloud Application
	//
	// - **CloudBrowser**: secure browser
	//
	// - **CloudDesktop**: WUYING Cloud Desktop
	//
	// - **ImageBuilder**: image creation resources
	//
	// - **ImageValidate**: image validation resources
	//
	// - **AndroidCloud**: Android cloud
	//
	// - **GpuRender**: GPU rendering
	//
	// - **AndroidGpuRender**: Android cloud GPU rendering
	//
	// - **WuyingServer**: enterprise edition workstation
	//
	// - **GpuProxy**: GPU remote rendering proxy
	//
	// - **GpuCache**: GPU remote rendering cache
	//
	// - **AIAgent**: AI agent service
	//
	// - **HostedDesktop**: hosted desktop
	//
	// - **WuyingWorkstation**: personal edition container workstation
	//
	// - **WuyingWorkstationBusiness**: business edition container workstation
	//
	// - **WuyingWorkstationTeam**: Linggou container workstation team edition
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
