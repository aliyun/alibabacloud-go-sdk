// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalRegionIdsShrink(v string) *GetImageShrinkRequest
	GetAdditionalRegionIdsShrink() *string
	SetImageCategory(v string) *GetImageShrinkRequest
	GetImageCategory() *string
	SetImageId(v string) *GetImageShrinkRequest
	GetImageId() *string
	SetImageType(v string) *GetImageShrinkRequest
	GetImageType() *string
}

type GetImageShrinkRequest struct {
	AdditionalRegionIdsShrink *string `json:"AdditionalRegionIds,omitempty" xml:"AdditionalRegionIds,omitempty"`
	// The source of the image. Valid values:
	//
	// 	- Public: public images provided by Alibaba Cloud.
	//
	// 	- Custom: the custom image that you added.
	//
	// example:
	//
	// Custom
	ImageCategory *string `json:"ImageCategory,omitempty" xml:"ImageCategory,omitempty"`
	// The image ID.
	//
	// example:
	//
	// m-2ze74g5mvy4pjg*****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The type of the images. Valid values:
	//
	// 	- VM: virtual machine image.
	//
	// 	- Container: the container image.
	//
	// Default value: VM
	//
	// example:
	//
	// VM
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
}

func (s GetImageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetImageShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetImageShrinkRequest) GetAdditionalRegionIdsShrink() *string {
	return s.AdditionalRegionIdsShrink
}

func (s *GetImageShrinkRequest) GetImageCategory() *string {
	return s.ImageCategory
}

func (s *GetImageShrinkRequest) GetImageId() *string {
	return s.ImageId
}

func (s *GetImageShrinkRequest) GetImageType() *string {
	return s.ImageType
}

func (s *GetImageShrinkRequest) SetAdditionalRegionIdsShrink(v string) *GetImageShrinkRequest {
	s.AdditionalRegionIdsShrink = &v
	return s
}

func (s *GetImageShrinkRequest) SetImageCategory(v string) *GetImageShrinkRequest {
	s.ImageCategory = &v
	return s
}

func (s *GetImageShrinkRequest) SetImageId(v string) *GetImageShrinkRequest {
	s.ImageId = &v
	return s
}

func (s *GetImageShrinkRequest) SetImageType(v string) *GetImageShrinkRequest {
	s.ImageType = &v
	return s
}

func (s *GetImageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
