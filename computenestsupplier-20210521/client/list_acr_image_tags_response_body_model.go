// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAcrImageTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImages(v []*ListAcrImageTagsResponseBodyImages) *ListAcrImageTagsResponseBody
	GetImages() []*ListAcrImageTagsResponseBodyImages
	SetMaxResults(v int32) *ListAcrImageTagsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListAcrImageTagsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAcrImageTagsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAcrImageTagsResponseBody
	GetTotalCount() *int32
}

type ListAcrImageTagsResponseBody struct {
	// The list of images.
	Images []*ListAcrImageTagsResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// ey14..
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FEF343B9-1A15-5789-BE88-7B36190F5BF6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAcrImageTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAcrImageTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAcrImageTagsResponseBody) GetImages() []*ListAcrImageTagsResponseBodyImages {
	return s.Images
}

func (s *ListAcrImageTagsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAcrImageTagsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAcrImageTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAcrImageTagsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAcrImageTagsResponseBody) SetImages(v []*ListAcrImageTagsResponseBodyImages) *ListAcrImageTagsResponseBody {
	s.Images = v
	return s
}

func (s *ListAcrImageTagsResponseBody) SetMaxResults(v int32) *ListAcrImageTagsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAcrImageTagsResponseBody) SetNextToken(v string) *ListAcrImageTagsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAcrImageTagsResponseBody) SetRequestId(v string) *ListAcrImageTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAcrImageTagsResponseBody) SetTotalCount(v int32) *ListAcrImageTagsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAcrImageTagsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAcrImageTagsResponseBodyImages struct {
	// The time when the image was created.
	//
	// example:
	//
	// 2021-05-20T00:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The image size. Unit: bytes.
	//
	// example:
	//
	// 188394616
	ImageSize *string `json:"ImageSize,omitempty" xml:"ImageSize,omitempty"`
	// The time when the image was modified.
	//
	// example:
	//
	// 2021-05-20T00:00:00Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The image version.
	//
	// example:
	//
	// 5.7.2
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListAcrImageTagsResponseBodyImages) String() string {
	return dara.Prettify(s)
}

func (s ListAcrImageTagsResponseBodyImages) GoString() string {
	return s.String()
}

func (s *ListAcrImageTagsResponseBodyImages) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAcrImageTagsResponseBodyImages) GetImageSize() *string {
	return s.ImageSize
}

func (s *ListAcrImageTagsResponseBodyImages) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *ListAcrImageTagsResponseBodyImages) GetTag() *string {
	return s.Tag
}

func (s *ListAcrImageTagsResponseBodyImages) SetCreateTime(v string) *ListAcrImageTagsResponseBodyImages {
	s.CreateTime = &v
	return s
}

func (s *ListAcrImageTagsResponseBodyImages) SetImageSize(v string) *ListAcrImageTagsResponseBodyImages {
	s.ImageSize = &v
	return s
}

func (s *ListAcrImageTagsResponseBodyImages) SetModifiedTime(v string) *ListAcrImageTagsResponseBodyImages {
	s.ModifiedTime = &v
	return s
}

func (s *ListAcrImageTagsResponseBodyImages) SetTag(v string) *ListAcrImageTagsResponseBodyImages {
	s.Tag = &v
	return s
}

func (s *ListAcrImageTagsResponseBodyImages) Validate() error {
	return dara.Validate(s)
}
