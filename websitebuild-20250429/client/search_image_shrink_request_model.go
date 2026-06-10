// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchImageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColorHex(v string) *SearchImageShrinkRequest
	GetColorHex() *string
	SetHasPerson(v bool) *SearchImageShrinkRequest
	GetHasPerson() *bool
	SetImageCategory(v string) *SearchImageShrinkRequest
	GetImageCategory() *string
	SetImageRatio(v string) *SearchImageShrinkRequest
	GetImageRatio() *string
	SetMaxHeight(v int32) *SearchImageShrinkRequest
	GetMaxHeight() *int32
	SetMaxResults(v int32) *SearchImageShrinkRequest
	GetMaxResults() *int32
	SetMaxWidth(v int32) *SearchImageShrinkRequest
	GetMaxWidth() *int32
	SetMinHeight(v int32) *SearchImageShrinkRequest
	GetMinHeight() *int32
	SetMinWidth(v int32) *SearchImageShrinkRequest
	GetMinWidth() *int32
	SetNextToken(v string) *SearchImageShrinkRequest
	GetNextToken() *string
	SetOssKey(v string) *SearchImageShrinkRequest
	GetOssKey() *string
	SetSize(v int32) *SearchImageShrinkRequest
	GetSize() *int32
	SetStart(v int32) *SearchImageShrinkRequest
	GetStart() *int32
	SetTagsShrink(v string) *SearchImageShrinkRequest
	GetTagsShrink() *string
	SetText(v string) *SearchImageShrinkRequest
	GetText() *string
}

type SearchImageShrinkRequest struct {
	// Color
	//
	// example:
	//
	// #B0B0B0
	ColorHex *string `json:"ColorHex,omitempty" xml:"ColorHex,omitempty"`
	// Indicates whether the image contains a person.
	//
	// example:
	//
	// false
	HasPerson *bool `json:"HasPerson,omitempty" xml:"HasPerson,omitempty"`
	// Image category. Valid values:
	//
	// - normal: Illustrations or article images.
	//
	// - banner: Background images or image carousels.
	//
	// - goods: Product or service images.
	//
	// example:
	//
	// WindowsWithMssqlStdLicense
	ImageCategory *string `json:"ImageCategory,omitempty" xml:"ImageCategory,omitempty"`
	// Image aspect ratio, including:
	//
	// "16:9"
	//
	// "4:3"
	//
	// "2:1"
	//
	// "1:1"
	//
	// "3:4"
	//
	// "9:16"
	//
	// example:
	//
	// 16:9
	ImageRatio *string `json:"ImageRatio,omitempty" xml:"ImageRatio,omitempty"`
	// Maximum image height.
	//
	// example:
	//
	// 4000
	MaxHeight *int32 `json:"MaxHeight,omitempty" xml:"MaxHeight,omitempty"`
	// Number of items per page in a paged query. Maximum value is 100. Default value is 20.
	//
	// example:
	//
	// 500
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Maximum image width (inclusive).
	//
	// example:
	//
	// 4000
	MaxWidth *int32 `json:"MaxWidth,omitempty" xml:"MaxWidth,omitempty"`
	// Minimum image height
	//
	// example:
	//
	// 500
	MinHeight *int32 `json:"MinHeight,omitempty" xml:"MinHeight,omitempty"`
	// Minimum image width (inclusive).
	//
	// example:
	//
	// 500
	MinWidth *int32 `json:"MinWidth,omitempty" xml:"MinWidth,omitempty"`
	// Query credential (Token). Set this parameter to the NextToken value returned in the previous API call. You do not need to set this parameter for the initial API call. If NextToken is specified, the request parameters PageSize and PageNumber become invalid, and the TotalCount in the returned data is also invalid.
	//
	// example:
	//
	// FFh3Xqm+JgZ/U9Jyb7wdVr9LWk80Tghn5UZjbcWEVEderBcbVF+Y6PS0i8PpCL4PQZ3e0C9oEH0Asd4tJEuGtkl2WuKdiWZpEwadNydQdJPFM=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Osskey。
	//
	// example:
	//
	// backend/detection/objects/r-0008ujvfksltf5j45n81/task-000hckiuwyau0gwn17vq.jpg
	OssKey *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	// Number of results to return. Default value is 10.
	//
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// Starting position of the return result. Valid values: 0 to 499. Default value is 0.
	//
	// example:
	//
	// 0
	Start *int32 `json:"Start,omitempty" xml:"Start,omitempty"`
	// Tags.
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// Description text for searching images.
	//
	// > Supports up to 512 characters.
	//
	// example:
	//
	// 卫浴五金产品，表面采用拉丝不锈钢材质，整体色调为现代银灰色，呈现简约且富有质感的风格。产品包括淋浴花洒、水龙头及毛巾架等配件，均具备防锈、耐腐蚀性能，适用于潮湿环境。各部件连接处设计精密，配有隐藏式螺丝结构，提升整体美观度。水龙头和花洒喷头支持多模式出水，把手操作顺滑，符合人体工学设计。所有五金件通过国家节水认证，支持冷热水接入，安装方式为壁挂式。
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s SearchImageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchImageShrinkRequest) GoString() string {
	return s.String()
}

func (s *SearchImageShrinkRequest) GetColorHex() *string {
	return s.ColorHex
}

func (s *SearchImageShrinkRequest) GetHasPerson() *bool {
	return s.HasPerson
}

func (s *SearchImageShrinkRequest) GetImageCategory() *string {
	return s.ImageCategory
}

func (s *SearchImageShrinkRequest) GetImageRatio() *string {
	return s.ImageRatio
}

func (s *SearchImageShrinkRequest) GetMaxHeight() *int32 {
	return s.MaxHeight
}

func (s *SearchImageShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *SearchImageShrinkRequest) GetMaxWidth() *int32 {
	return s.MaxWidth
}

func (s *SearchImageShrinkRequest) GetMinHeight() *int32 {
	return s.MinHeight
}

func (s *SearchImageShrinkRequest) GetMinWidth() *int32 {
	return s.MinWidth
}

func (s *SearchImageShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *SearchImageShrinkRequest) GetOssKey() *string {
	return s.OssKey
}

func (s *SearchImageShrinkRequest) GetSize() *int32 {
	return s.Size
}

func (s *SearchImageShrinkRequest) GetStart() *int32 {
	return s.Start
}

func (s *SearchImageShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *SearchImageShrinkRequest) GetText() *string {
	return s.Text
}

func (s *SearchImageShrinkRequest) SetColorHex(v string) *SearchImageShrinkRequest {
	s.ColorHex = &v
	return s
}

func (s *SearchImageShrinkRequest) SetHasPerson(v bool) *SearchImageShrinkRequest {
	s.HasPerson = &v
	return s
}

func (s *SearchImageShrinkRequest) SetImageCategory(v string) *SearchImageShrinkRequest {
	s.ImageCategory = &v
	return s
}

func (s *SearchImageShrinkRequest) SetImageRatio(v string) *SearchImageShrinkRequest {
	s.ImageRatio = &v
	return s
}

func (s *SearchImageShrinkRequest) SetMaxHeight(v int32) *SearchImageShrinkRequest {
	s.MaxHeight = &v
	return s
}

func (s *SearchImageShrinkRequest) SetMaxResults(v int32) *SearchImageShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *SearchImageShrinkRequest) SetMaxWidth(v int32) *SearchImageShrinkRequest {
	s.MaxWidth = &v
	return s
}

func (s *SearchImageShrinkRequest) SetMinHeight(v int32) *SearchImageShrinkRequest {
	s.MinHeight = &v
	return s
}

func (s *SearchImageShrinkRequest) SetMinWidth(v int32) *SearchImageShrinkRequest {
	s.MinWidth = &v
	return s
}

func (s *SearchImageShrinkRequest) SetNextToken(v string) *SearchImageShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *SearchImageShrinkRequest) SetOssKey(v string) *SearchImageShrinkRequest {
	s.OssKey = &v
	return s
}

func (s *SearchImageShrinkRequest) SetSize(v int32) *SearchImageShrinkRequest {
	s.Size = &v
	return s
}

func (s *SearchImageShrinkRequest) SetStart(v int32) *SearchImageShrinkRequest {
	s.Start = &v
	return s
}

func (s *SearchImageShrinkRequest) SetTagsShrink(v string) *SearchImageShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *SearchImageShrinkRequest) SetText(v string) *SearchImageShrinkRequest {
	s.Text = &v
	return s
}

func (s *SearchImageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
