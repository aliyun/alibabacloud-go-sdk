// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageCategory(v string) *ListImagesRequest
	GetImageCategory() *string
	SetImageIds(v []*string) *ListImagesRequest
	GetImageIds() []*string
	SetImageNames(v []*string) *ListImagesRequest
	GetImageNames() []*string
	SetImageType(v string) *ListImagesRequest
	GetImageType() *string
	SetMode(v string) *ListImagesRequest
	GetMode() *string
	SetPageNumber(v int64) *ListImagesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListImagesRequest
	GetPageSize() *int64
}

type ListImagesRequest struct {
	ImageCategory *string   `json:"ImageCategory,omitempty" xml:"ImageCategory,omitempty"`
	ImageIds      []*string `json:"ImageIds,omitempty" xml:"ImageIds,omitempty" type:"Repeated"`
	ImageNames    []*string `json:"ImageNames,omitempty" xml:"ImageNames,omitempty" type:"Repeated"`
	ImageType     *string   `json:"ImageType,omitempty" xml:"ImageType,omitempty"`
	Mode          *string   `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListImagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListImagesRequest) GoString() string {
	return s.String()
}

func (s *ListImagesRequest) GetImageCategory() *string {
	return s.ImageCategory
}

func (s *ListImagesRequest) GetImageIds() []*string {
	return s.ImageIds
}

func (s *ListImagesRequest) GetImageNames() []*string {
	return s.ImageNames
}

func (s *ListImagesRequest) GetImageType() *string {
	return s.ImageType
}

func (s *ListImagesRequest) GetMode() *string {
	return s.Mode
}

func (s *ListImagesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListImagesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListImagesRequest) SetImageCategory(v string) *ListImagesRequest {
	s.ImageCategory = &v
	return s
}

func (s *ListImagesRequest) SetImageIds(v []*string) *ListImagesRequest {
	s.ImageIds = v
	return s
}

func (s *ListImagesRequest) SetImageNames(v []*string) *ListImagesRequest {
	s.ImageNames = v
	return s
}

func (s *ListImagesRequest) SetImageType(v string) *ListImagesRequest {
	s.ImageType = &v
	return s
}

func (s *ListImagesRequest) SetMode(v string) *ListImagesRequest {
	s.Mode = &v
	return s
}

func (s *ListImagesRequest) SetPageNumber(v int64) *ListImagesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListImagesRequest) SetPageSize(v int64) *ListImagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListImagesRequest) Validate() error {
	return dara.Validate(s)
}
