// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEffectiveImageByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizRegionId(v string) *GetEffectiveImageByIdRequest
	GetBizRegionId() *string
	SetContentLocale(v string) *GetEffectiveImageByIdRequest
	GetContentLocale() *string
	SetIgnoreFilter(v bool) *GetEffectiveImageByIdRequest
	GetIgnoreFilter() *bool
	SetImageId(v string) *GetEffectiveImageByIdRequest
	GetImageId() *string
	SetImageType(v string) *GetEffectiveImageByIdRequest
	GetImageType() *string
	SetProductType(v string) *GetEffectiveImageByIdRequest
	GetProductType() *string
	SetShowHidden(v string) *GetEffectiveImageByIdRequest
	GetShowHidden() *string
}

type GetEffectiveImageByIdRequest struct {
	// The ID of the business region to query. If this parameter is not specified, regions are not filtered by this parameter. The returned regions are still constrained by the product and the visible scope of the account.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The content language of the community image document in BCP 47 format, such as zh-CN or en-US. Default value: zh-CN.
	//
	// example:
	//
	// zh-CN
	ContentLocale *string `json:"ContentLocale,omitempty" xml:"ContentLocale,omitempty"`
	// Specifies whether to skip filtering by the operating system and adaptation status of applications. Default value: false. A value of false retains only applications that match the image operating system and have completed adaptation. A value of true skips these two filters. Application type filtering is still controlled by the API operation and the ShowHidden parameter.
	//
	// example:
	//
	// false
	IgnoreFilter *bool `json:"IgnoreFilter,omitempty" xml:"IgnoreFilter,omitempty"`
	// The image ID. Required when querying effective versions.
	//
	// example:
	//
	// imgc-bp1example****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image type. Valid values:
	//
	// - System: System image.
	//
	// - User: Custom image.
	//
	// - Shared: Shared image.
	//
	// - Community: Community image.
	//
	// example:
	//
	// User
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// The product type, such as CloudDesktop, CloudApp, or WuyingServer.
	//
	// example:
	//
	// CloudDesktop
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The additional types of virtual applications to include, in addition to the user and system applications returned by default. Valid values:
	//
	// - VIRTUAL_DESKTOP_APP: Desktop virtual applications.
	//
	// - ALL_VIRTUAL_APP: All virtual applications.
	//
	// example:
	//
	// VIRTUAL_DESKTOP_APP
	ShowHidden *string `json:"ShowHidden,omitempty" xml:"ShowHidden,omitempty"`
}

func (s GetEffectiveImageByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEffectiveImageByIdRequest) GoString() string {
	return s.String()
}

func (s *GetEffectiveImageByIdRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *GetEffectiveImageByIdRequest) GetContentLocale() *string {
	return s.ContentLocale
}

func (s *GetEffectiveImageByIdRequest) GetIgnoreFilter() *bool {
	return s.IgnoreFilter
}

func (s *GetEffectiveImageByIdRequest) GetImageId() *string {
	return s.ImageId
}

func (s *GetEffectiveImageByIdRequest) GetImageType() *string {
	return s.ImageType
}

func (s *GetEffectiveImageByIdRequest) GetProductType() *string {
	return s.ProductType
}

func (s *GetEffectiveImageByIdRequest) GetShowHidden() *string {
	return s.ShowHidden
}

func (s *GetEffectiveImageByIdRequest) SetBizRegionId(v string) *GetEffectiveImageByIdRequest {
	s.BizRegionId = &v
	return s
}

func (s *GetEffectiveImageByIdRequest) SetContentLocale(v string) *GetEffectiveImageByIdRequest {
	s.ContentLocale = &v
	return s
}

func (s *GetEffectiveImageByIdRequest) SetIgnoreFilter(v bool) *GetEffectiveImageByIdRequest {
	s.IgnoreFilter = &v
	return s
}

func (s *GetEffectiveImageByIdRequest) SetImageId(v string) *GetEffectiveImageByIdRequest {
	s.ImageId = &v
	return s
}

func (s *GetEffectiveImageByIdRequest) SetImageType(v string) *GetEffectiveImageByIdRequest {
	s.ImageType = &v
	return s
}

func (s *GetEffectiveImageByIdRequest) SetProductType(v string) *GetEffectiveImageByIdRequest {
	s.ProductType = &v
	return s
}

func (s *GetEffectiveImageByIdRequest) SetShowHidden(v string) *GetEffectiveImageByIdRequest {
	s.ShowHidden = &v
	return s
}

func (s *GetEffectiveImageByIdRequest) Validate() error {
	return dara.Validate(s)
}
