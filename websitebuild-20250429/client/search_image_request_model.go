// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColorHex(v string) *SearchImageRequest
	GetColorHex() *string
	SetHasPerson(v bool) *SearchImageRequest
	GetHasPerson() *bool
	SetImageCategory(v string) *SearchImageRequest
	GetImageCategory() *string
	SetImageRatio(v string) *SearchImageRequest
	GetImageRatio() *string
	SetMaxHeight(v int32) *SearchImageRequest
	GetMaxHeight() *int32
	SetMaxResults(v int32) *SearchImageRequest
	GetMaxResults() *int32
	SetMaxWidth(v int32) *SearchImageRequest
	GetMaxWidth() *int32
	SetMinHeight(v int32) *SearchImageRequest
	GetMinHeight() *int32
	SetMinWidth(v int32) *SearchImageRequest
	GetMinWidth() *int32
	SetNextToken(v string) *SearchImageRequest
	GetNextToken() *string
	SetOssKey(v string) *SearchImageRequest
	GetOssKey() *string
	SetSize(v int32) *SearchImageRequest
	GetSize() *int32
	SetStart(v int32) *SearchImageRequest
	GetStart() *int32
	SetTags(v []*string) *SearchImageRequest
	GetTags() []*string
	SetText(v string) *SearchImageRequest
	GetText() *string
}

type SearchImageRequest struct {
	// The color.
	//
	// example:
	//
	// #B0B0B0
	ColorHex *string `json:"ColorHex,omitempty" xml:"ColorHex,omitempty"`
	// Specifies whether the image contains a person.
	//
	// example:
	//
	// false
	HasPerson *bool `json:"HasPerson,omitempty" xml:"HasPerson,omitempty"`
	// The image category. Valid values:
	//
	// - normal: Illustrations and article images.
	//
	// - banner: Background images and carousel images.
	//
	// - goods: Product or service images.
	//
	// example:
	//
	// WindowsWithMssqlStdLicense
	ImageCategory *string `json:"ImageCategory,omitempty" xml:"ImageCategory,omitempty"`
	// The aspect ratio of the image. Valid values:
	//
	// - "16:9"
	//
	// - "4:3"
	//
	// - "2:1"
	//
	// - "1:1"
	//
	// - "3:4"
	//
	// - "9:16"
	//
	// example:
	//
	// 16:9
	ImageRatio *string `json:"ImageRatio,omitempty" xml:"ImageRatio,omitempty"`
	// The maximum height of the image.
	//
	// example:
	//
	// 4000
	MaxHeight *int32 `json:"MaxHeight,omitempty" xml:"MaxHeight,omitempty"`
	// The number of entries per page for a paged query. Maximum value: 100. Default value: 20.
	//
	// example:
	//
	// 500
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The maximum width of the image, inclusive.
	//
	// example:
	//
	// 4000
	MaxWidth *int32 `json:"MaxWidth,omitempty" xml:"MaxWidth,omitempty"`
	// The minimum height of the image.
	//
	// example:
	//
	// 500
	MinHeight *int32 `json:"MinHeight,omitempty" xml:"MinHeight,omitempty"`
	// The minimum width of the image, inclusive.
	//
	// example:
	//
	// 500
	MinWidth *int32 `json:"MinWidth,omitempty" xml:"MinWidth,omitempty"`
	// The query token. Set the value to the NextToken value returned from the previous call to this operation. You do not need to set this parameter for the first call. If NextToken is set, the PageSize and PageNumber request parameters become invalid, and the TotalCount value in the response is invalid.
	//
	// example:
	//
	// FFh3Xqm+JgZ/U9Jyb7wdVr9LWk80Tghn5UZjbcWEVEderBcbVF+Y6PS0i8PpCL4PQZ3e0C9oEH0Asd4tJEuGtkl2WuKdiWZpEwadNydQdJPFM=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The OSS key.
	//
	// example:
	//
	// backend/detection/objects/r-0008ujvfksltf5j45n81/task-000hckiuwyau0gwn17vq.jpg
	OssKey *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	// The number of results to return. Default value: 10.
	//
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The starting position of the returned results. Valid values: 0 to 499. Default value: 0.
	//
	// example:
	//
	// 0
	Start *int32 `json:"Start,omitempty" xml:"Start,omitempty"`
	// The tags.
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The descriptive text used to search for images. A maximum of 512 characters are supported.
	//
	// example:
	//
	// Bathroom hardware products, featuring a brushed stainless steel surface with an overall modern silver-gray tone, presenting a minimalist and textured style. The products include accessories such as showerheads, faucets, and towel racks, all of which are rust-proof and corrosion-resistant, making them suitable for humid environments. The joints between components are precisely designed with hidden screw structures to enhance overall aesthetics. The faucets and showerheads support multiple water flow modes, and the handles operate smoothly, conforming to ergonomic design. All hardware fittings have passed the national water-saving certification, support both cold and hot water connections, and are wall-mounted
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s SearchImageRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchImageRequest) GoString() string {
	return s.String()
}

func (s *SearchImageRequest) GetColorHex() *string {
	return s.ColorHex
}

func (s *SearchImageRequest) GetHasPerson() *bool {
	return s.HasPerson
}

func (s *SearchImageRequest) GetImageCategory() *string {
	return s.ImageCategory
}

func (s *SearchImageRequest) GetImageRatio() *string {
	return s.ImageRatio
}

func (s *SearchImageRequest) GetMaxHeight() *int32 {
	return s.MaxHeight
}

func (s *SearchImageRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *SearchImageRequest) GetMaxWidth() *int32 {
	return s.MaxWidth
}

func (s *SearchImageRequest) GetMinHeight() *int32 {
	return s.MinHeight
}

func (s *SearchImageRequest) GetMinWidth() *int32 {
	return s.MinWidth
}

func (s *SearchImageRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *SearchImageRequest) GetOssKey() *string {
	return s.OssKey
}

func (s *SearchImageRequest) GetSize() *int32 {
	return s.Size
}

func (s *SearchImageRequest) GetStart() *int32 {
	return s.Start
}

func (s *SearchImageRequest) GetTags() []*string {
	return s.Tags
}

func (s *SearchImageRequest) GetText() *string {
	return s.Text
}

func (s *SearchImageRequest) SetColorHex(v string) *SearchImageRequest {
	s.ColorHex = &v
	return s
}

func (s *SearchImageRequest) SetHasPerson(v bool) *SearchImageRequest {
	s.HasPerson = &v
	return s
}

func (s *SearchImageRequest) SetImageCategory(v string) *SearchImageRequest {
	s.ImageCategory = &v
	return s
}

func (s *SearchImageRequest) SetImageRatio(v string) *SearchImageRequest {
	s.ImageRatio = &v
	return s
}

func (s *SearchImageRequest) SetMaxHeight(v int32) *SearchImageRequest {
	s.MaxHeight = &v
	return s
}

func (s *SearchImageRequest) SetMaxResults(v int32) *SearchImageRequest {
	s.MaxResults = &v
	return s
}

func (s *SearchImageRequest) SetMaxWidth(v int32) *SearchImageRequest {
	s.MaxWidth = &v
	return s
}

func (s *SearchImageRequest) SetMinHeight(v int32) *SearchImageRequest {
	s.MinHeight = &v
	return s
}

func (s *SearchImageRequest) SetMinWidth(v int32) *SearchImageRequest {
	s.MinWidth = &v
	return s
}

func (s *SearchImageRequest) SetNextToken(v string) *SearchImageRequest {
	s.NextToken = &v
	return s
}

func (s *SearchImageRequest) SetOssKey(v string) *SearchImageRequest {
	s.OssKey = &v
	return s
}

func (s *SearchImageRequest) SetSize(v int32) *SearchImageRequest {
	s.Size = &v
	return s
}

func (s *SearchImageRequest) SetStart(v int32) *SearchImageRequest {
	s.Start = &v
	return s
}

func (s *SearchImageRequest) SetTags(v []*string) *SearchImageRequest {
	s.Tags = v
	return s
}

func (s *SearchImageRequest) SetText(v string) *SearchImageRequest {
	s.Text = &v
	return s
}

func (s *SearchImageRequest) Validate() error {
	return dara.Validate(s)
}
