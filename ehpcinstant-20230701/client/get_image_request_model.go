// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalRegionIds(v []*string) *GetImageRequest
	GetAdditionalRegionIds() []*string
	SetImageCategory(v string) *GetImageRequest
	GetImageCategory() *string
	SetImageId(v string) *GetImageRequest
	GetImageId() *string
	SetImageType(v string) *GetImageRequest
	GetImageType() *string
}

type GetImageRequest struct {
	AdditionalRegionIds []*string `json:"AdditionalRegionIds,omitempty" xml:"AdditionalRegionIds,omitempty" type:"Repeated"`
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

func (s GetImageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetImageRequest) GoString() string {
	return s.String()
}

func (s *GetImageRequest) GetAdditionalRegionIds() []*string {
	return s.AdditionalRegionIds
}

func (s *GetImageRequest) GetImageCategory() *string {
	return s.ImageCategory
}

func (s *GetImageRequest) GetImageId() *string {
	return s.ImageId
}

func (s *GetImageRequest) GetImageType() *string {
	return s.ImageType
}

func (s *GetImageRequest) SetAdditionalRegionIds(v []*string) *GetImageRequest {
	s.AdditionalRegionIds = v
	return s
}

func (s *GetImageRequest) SetImageCategory(v string) *GetImageRequest {
	s.ImageCategory = &v
	return s
}

func (s *GetImageRequest) SetImageId(v string) *GetImageRequest {
	s.ImageId = &v
	return s
}

func (s *GetImageRequest) SetImageType(v string) *GetImageRequest {
	s.ImageType = &v
	return s
}

func (s *GetImageRequest) Validate() error {
	return dara.Validate(s)
}
