// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRenderingImagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *ListRenderingImagesRequest
	GetImageId() *string
	SetPageNumber(v int32) *ListRenderingImagesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRenderingImagesRequest
	GetPageSize() *int32
}

type ListRenderingImagesRequest struct {
	// The cloud application service instance ID.
	//
	// example:
	//
	// m-9timxhrrgopkec8ju
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The page number. The value starts from 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for a paged query.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListRenderingImagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRenderingImagesRequest) GoString() string {
	return s.String()
}

func (s *ListRenderingImagesRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ListRenderingImagesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRenderingImagesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRenderingImagesRequest) SetImageId(v string) *ListRenderingImagesRequest {
	s.ImageId = &v
	return s
}

func (s *ListRenderingImagesRequest) SetPageNumber(v int32) *ListRenderingImagesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRenderingImagesRequest) SetPageSize(v int32) *ListRenderingImagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListRenderingImagesRequest) Validate() error {
	return dara.Validate(s)
}
