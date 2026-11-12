// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizRegionId(v string) *GetImageByIdRequest
	GetBizRegionId() *string
	SetIgnoreFilter(v bool) *GetImageByIdRequest
	GetIgnoreFilter() *bool
	SetImageId(v string) *GetImageByIdRequest
	GetImageId() *string
	SetProductType(v string) *GetImageByIdRequest
	GetProductType() *string
	SetShowHidden(v string) *GetImageByIdRequest
	GetShowHidden() *string
	SetVersionId(v string) *GetImageByIdRequest
	GetVersionId() *string
}

type GetImageByIdRequest struct {
	// The ID of the business region to query. If you do not specify this parameter, regions are not filtered by this parameter. The returned regions are still subject to the visibility scope of the product and the account.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// Specifies whether to skip the operating system and adaptation status filtering for adapted applications. Default value: false. If you set this parameter to false, only applications whose operating system matches the image and whose adaptation is complete are retained. If you set this parameter to true, these two filters are skipped. Application type filtering is still controlled by the API operation and the ShowHidden parameter.
	IgnoreFilter *bool `json:"IgnoreFilter,omitempty" xml:"IgnoreFilter,omitempty"`
	// The image ID. You must specify at least one of ImageId and VersionId. If you specify only this parameter, the online version is queried.
	//
	// example:
	//
	// imgc-bp1example****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The product type. Examples: CloudDesktop (Cloud Desktop), CloudApp (Cloud App), and WuyingServer (WUYING Workspace Enterprise Edition workstation).
	//
	// example:
	//
	// CloudDesktop
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The types of virtual applications to include in addition to the default user applications and system applications. Valid values:
	//
	// - VIRTUAL_DESKTOP_APP: desktop virtual applications.
	//
	// - ALL_VIRTUAL_APP: all virtual applications.
	//
	// example:
	//
	// VIRTUAL_DESKTOP_APP
	ShowHidden *string `json:"ShowHidden,omitempty" xml:"ShowHidden,omitempty"`
	// The image version ID. You can specify this parameter alone to query a specific version. If you specify both this parameter and ImageId, the version must belong to the specified image.
	//
	// example:
	//
	// iv-xxx
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetImageByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetImageByIdRequest) GoString() string {
	return s.String()
}

func (s *GetImageByIdRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *GetImageByIdRequest) GetIgnoreFilter() *bool {
	return s.IgnoreFilter
}

func (s *GetImageByIdRequest) GetImageId() *string {
	return s.ImageId
}

func (s *GetImageByIdRequest) GetProductType() *string {
	return s.ProductType
}

func (s *GetImageByIdRequest) GetShowHidden() *string {
	return s.ShowHidden
}

func (s *GetImageByIdRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *GetImageByIdRequest) SetBizRegionId(v string) *GetImageByIdRequest {
	s.BizRegionId = &v
	return s
}

func (s *GetImageByIdRequest) SetIgnoreFilter(v bool) *GetImageByIdRequest {
	s.IgnoreFilter = &v
	return s
}

func (s *GetImageByIdRequest) SetImageId(v string) *GetImageByIdRequest {
	s.ImageId = &v
	return s
}

func (s *GetImageByIdRequest) SetProductType(v string) *GetImageByIdRequest {
	s.ProductType = &v
	return s
}

func (s *GetImageByIdRequest) SetShowHidden(v string) *GetImageByIdRequest {
	s.ShowHidden = &v
	return s
}

func (s *GetImageByIdRequest) SetVersionId(v string) *GetImageByIdRequest {
	s.VersionId = &v
	return s
}

func (s *GetImageByIdRequest) Validate() error {
	return dara.Validate(s)
}
