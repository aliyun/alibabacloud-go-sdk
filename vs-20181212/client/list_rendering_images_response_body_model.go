// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRenderingImagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImages(v []*ListRenderingImagesResponseBodyImages) *ListRenderingImagesResponseBody
	GetImages() []*ListRenderingImagesResponseBodyImages
	SetPageNumber(v string) *ListRenderingImagesResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *ListRenderingImagesResponseBody
	GetPageSize() *string
	SetRequestId(v string) *ListRenderingImagesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListRenderingImagesResponseBody
	GetTotalCount() *int64
}

type ListRenderingImagesResponseBody struct {
	// The session list.
	Images []*ListRenderingImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of matched sessions.
	//
	// example:
	//
	// 8
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRenderingImagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRenderingImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRenderingImagesResponseBody) GetImages() []*ListRenderingImagesResponseBodyImages {
	return s.Images
}

func (s *ListRenderingImagesResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListRenderingImagesResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *ListRenderingImagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRenderingImagesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListRenderingImagesResponseBody) SetImages(v []*ListRenderingImagesResponseBodyImages) *ListRenderingImagesResponseBody {
	s.Images = v
	return s
}

func (s *ListRenderingImagesResponseBody) SetPageNumber(v string) *ListRenderingImagesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListRenderingImagesResponseBody) SetPageSize(v string) *ListRenderingImagesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRenderingImagesResponseBody) SetRequestId(v string) *ListRenderingImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRenderingImagesResponseBody) SetTotalCount(v int64) *ListRenderingImagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRenderingImagesResponseBody) Validate() error {
	if s.Images != nil {
		for _, item := range s.Images {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRenderingImagesResponseBodyImages struct {
	// The image description.
	//
	// example:
	//
	// this is test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The image ID.
	//
	// example:
	//
	// m-0nd0nxyl7zn220n2y
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The publish time.
	//
	// example:
	//
	// 2026-09-17T00:22:53+08:00
	PublishTime *string `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
}

func (s ListRenderingImagesResponseBodyImages) String() string {
	return dara.Prettify(s)
}

func (s ListRenderingImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *ListRenderingImagesResponseBodyImages) GetDescription() *string {
	return s.Description
}

func (s *ListRenderingImagesResponseBodyImages) GetImageId() *string {
	return s.ImageId
}

func (s *ListRenderingImagesResponseBodyImages) GetName() *string {
	return s.Name
}

func (s *ListRenderingImagesResponseBodyImages) GetPublishTime() *string {
	return s.PublishTime
}

func (s *ListRenderingImagesResponseBodyImages) SetDescription(v string) *ListRenderingImagesResponseBodyImages {
	s.Description = &v
	return s
}

func (s *ListRenderingImagesResponseBodyImages) SetImageId(v string) *ListRenderingImagesResponseBodyImages {
	s.ImageId = &v
	return s
}

func (s *ListRenderingImagesResponseBodyImages) SetName(v string) *ListRenderingImagesResponseBodyImages {
	s.Name = &v
	return s
}

func (s *ListRenderingImagesResponseBodyImages) SetPublishTime(v string) *ListRenderingImagesResponseBodyImages {
	s.PublishTime = &v
	return s
}

func (s *ListRenderingImagesResponseBodyImages) Validate() error {
	return dara.Validate(s)
}
