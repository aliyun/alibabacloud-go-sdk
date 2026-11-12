// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageIds(v string) *ListImagesRequest
	GetImageIds() *string
	SetImageType(v string) *ListImagesRequest
	GetImageType() *string
	SetRegionId(v string) *ListImagesRequest
	GetRegionId() *string
}

type ListImagesRequest struct {
	// The image IDs. The value can be a JSON array that consists of up to 50 image IDs. Format: `["xxx", "yyy", … "zzz"]`. Separate multiple image IDs with commas (,).
	//
	// example:
	//
	// ["fe9c66133a9d4688872869726b52****", "794c230fd3e64ea19f83f4d7a0ad****"]
	ImageIds *string `json:"ImageIds,omitempty" xml:"ImageIds,omitempty"`
	// The type of the images. Valid values:
	//
	// - system: OS images
	//
	// - app: application images
	//
	// - custom: custom images
	//
	// example:
	//
	// system
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// The region ID of the images. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListImagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListImagesRequest) GoString() string {
	return s.String()
}

func (s *ListImagesRequest) GetImageIds() *string {
	return s.ImageIds
}

func (s *ListImagesRequest) GetImageType() *string {
	return s.ImageType
}

func (s *ListImagesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListImagesRequest) SetImageIds(v string) *ListImagesRequest {
	s.ImageIds = &v
	return s
}

func (s *ListImagesRequest) SetImageType(v string) *ListImagesRequest {
	s.ImageType = &v
	return s
}

func (s *ListImagesRequest) SetRegionId(v string) *ListImagesRequest {
	s.RegionId = &v
	return s
}

func (s *ListImagesRequest) Validate() error {
	return dara.Validate(s)
}
