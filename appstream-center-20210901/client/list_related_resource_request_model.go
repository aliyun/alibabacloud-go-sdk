// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRelatedResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *ListRelatedResourceRequest
	GetImageId() *string
	SetPageNumber(v int32) *ListRelatedResourceRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRelatedResourceRequest
	GetPageSize() *int32
}

type ListRelatedResourceRequest struct {
	// The image ID.
	//
	// example:
	//
	// imgc-bp1example****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of delivery groups returned per page, which is the same as the PageSize request parameter.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListRelatedResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRelatedResourceRequest) GoString() string {
	return s.String()
}

func (s *ListRelatedResourceRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ListRelatedResourceRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRelatedResourceRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRelatedResourceRequest) SetImageId(v string) *ListRelatedResourceRequest {
	s.ImageId = &v
	return s
}

func (s *ListRelatedResourceRequest) SetPageNumber(v int32) *ListRelatedResourceRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRelatedResourceRequest) SetPageSize(v int32) *ListRelatedResourceRequest {
	s.PageSize = &v
	return s
}

func (s *ListRelatedResourceRequest) Validate() error {
	return dara.Validate(s)
}
