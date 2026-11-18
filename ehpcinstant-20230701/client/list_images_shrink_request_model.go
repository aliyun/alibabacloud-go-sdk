// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImagesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageCategory(v string) *ListImagesShrinkRequest
	GetImageCategory() *string
	SetImageIdsShrink(v string) *ListImagesShrinkRequest
	GetImageIdsShrink() *string
	SetImageNamesShrink(v string) *ListImagesShrinkRequest
	GetImageNamesShrink() *string
	SetImageType(v string) *ListImagesShrinkRequest
	GetImageType() *string
	SetMode(v string) *ListImagesShrinkRequest
	GetMode() *string
	SetPageNumber(v int64) *ListImagesShrinkRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListImagesShrinkRequest
	GetPageSize() *int64
}

type ListImagesShrinkRequest struct {
	// The image source. Valid values:
	//
	// - Public: public images provided by Alibaba Cloud
	//
	// - Custom: custom images added by you
	//
	// example:
	//
	// Custom
	ImageCategory *string `json:"ImageCategory,omitempty" xml:"ImageCategory,omitempty"`
	// The array of image IDs.
	ImageIdsShrink *string `json:"ImageIds,omitempty" xml:"ImageIds,omitempty"`
	// The array of image names.
	ImageNamesShrink *string `json:"ImageNames,omitempty" xml:"ImageNames,omitempty"`
	// The image type. Valid values:
	//
	// - VM: virtual machine image
	//
	// - Container: container image
	//
	// Default value: VM.
	//
	// example:
	//
	// VM
	ImageType *string `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	// The query mode. Valid values:
	//
	// - List: Queries the list of all image versions.
	//
	// - Merge: Queries the list of the latest versions for each image.
	//
	// example:
	//
	// Merge
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The current page number.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Default value: 20.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListImagesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListImagesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListImagesShrinkRequest) GetImageCategory() *string {
	return s.ImageCategory
}

func (s *ListImagesShrinkRequest) GetImageIdsShrink() *string {
	return s.ImageIdsShrink
}

func (s *ListImagesShrinkRequest) GetImageNamesShrink() *string {
	return s.ImageNamesShrink
}

func (s *ListImagesShrinkRequest) GetImageType() *string {
	return s.ImageType
}

func (s *ListImagesShrinkRequest) GetMode() *string {
	return s.Mode
}

func (s *ListImagesShrinkRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListImagesShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListImagesShrinkRequest) SetImageCategory(v string) *ListImagesShrinkRequest {
	s.ImageCategory = &v
	return s
}

func (s *ListImagesShrinkRequest) SetImageIdsShrink(v string) *ListImagesShrinkRequest {
	s.ImageIdsShrink = &v
	return s
}

func (s *ListImagesShrinkRequest) SetImageNamesShrink(v string) *ListImagesShrinkRequest {
	s.ImageNamesShrink = &v
	return s
}

func (s *ListImagesShrinkRequest) SetImageType(v string) *ListImagesShrinkRequest {
	s.ImageType = &v
	return s
}

func (s *ListImagesShrinkRequest) SetMode(v string) *ListImagesShrinkRequest {
	s.Mode = &v
	return s
}

func (s *ListImagesShrinkRequest) SetPageNumber(v int64) *ListImagesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListImagesShrinkRequest) SetPageSize(v int64) *ListImagesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListImagesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
