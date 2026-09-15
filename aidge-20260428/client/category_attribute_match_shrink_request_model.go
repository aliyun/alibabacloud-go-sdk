// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCategoryAttributeMatchShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CategoryAttributeMatchShrinkRequest
	GetDescription() *string
	SetImageUrlShrink(v string) *CategoryAttributeMatchShrinkRequest
	GetImageUrlShrink() *string
	SetItemSpec(v string) *CategoryAttributeMatchShrinkRequest
	GetItemSpec() *string
	SetSku(v string) *CategoryAttributeMatchShrinkRequest
	GetSku() *string
	SetSourceCategory(v string) *CategoryAttributeMatchShrinkRequest
	GetSourceCategory() *string
	SetSourcePlatform(v string) *CategoryAttributeMatchShrinkRequest
	GetSourcePlatform() *string
	SetTargetPlatform(v string) *CategoryAttributeMatchShrinkRequest
	GetTargetPlatform() *string
	SetTitle(v string) *CategoryAttributeMatchShrinkRequest
	GetTitle() *string
}

type CategoryAttributeMatchShrinkRequest struct {
	// The product details.
	//
	// example:
	//
	// Silk, new spring style
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The product image URLs. A maximum of 10 images are supported.
	//
	// example:
	//
	// https://frametour-assets.oss-cn-shanghai.aliyuncs.com/user-faces/viid_face/dd0dd06c-9351-4e5f-bc70-24166a754d7f.jpg
	ImageUrlShrink *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The product attributes that describe the product characteristics.
	//
	// This parameter is required.
	//
	// example:
	//
	// For example, the material of the product. Enter the attribute name and attribute value
	ItemSpec *string `json:"ItemSpec,omitempty" xml:"ItemSpec,omitempty"`
	// The product SKU title.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0:0:Color:Black (without locator);0:1:Color:Smart Locator (without collar);0:2:Color:Range Locator (without collar);0:3:Color:Blue (without locator);0:4:Color:Orange (without locator);0:5:Color:Pink (without locator);0:6:Color:Red (without locator);0:7:Color:Yellow (without locator);0:8:Color:Purple (without locator)
	Sku *string `json:"Sku,omitempty" xml:"Sku,omitempty"`
	// The product category on the source platform.
	//
	// This parameter is required.
	//
	// example:
	//
	// Clothing
	SourceCategory *string `json:"SourceCategory,omitempty" xml:"SourceCategory,omitempty"`
	// The source platform from which the product originates.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1688
	SourcePlatform *string `json:"SourcePlatform,omitempty" xml:"SourcePlatform,omitempty"`
	// The target listing platform. Currently, only temu is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// temu
	TargetPlatform *string `json:"TargetPlatform,omitempty" xml:"TargetPlatform,omitempty"`
	// The product title.
	//
	// This parameter is required.
	//
	// example:
	//
	// Women\\"s New Spring Clothing
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CategoryAttributeMatchShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CategoryAttributeMatchShrinkRequest) GoString() string {
	return s.String()
}

func (s *CategoryAttributeMatchShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CategoryAttributeMatchShrinkRequest) GetImageUrlShrink() *string {
	return s.ImageUrlShrink
}

func (s *CategoryAttributeMatchShrinkRequest) GetItemSpec() *string {
	return s.ItemSpec
}

func (s *CategoryAttributeMatchShrinkRequest) GetSku() *string {
	return s.Sku
}

func (s *CategoryAttributeMatchShrinkRequest) GetSourceCategory() *string {
	return s.SourceCategory
}

func (s *CategoryAttributeMatchShrinkRequest) GetSourcePlatform() *string {
	return s.SourcePlatform
}

func (s *CategoryAttributeMatchShrinkRequest) GetTargetPlatform() *string {
	return s.TargetPlatform
}

func (s *CategoryAttributeMatchShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *CategoryAttributeMatchShrinkRequest) SetDescription(v string) *CategoryAttributeMatchShrinkRequest {
	s.Description = &v
	return s
}

func (s *CategoryAttributeMatchShrinkRequest) SetImageUrlShrink(v string) *CategoryAttributeMatchShrinkRequest {
	s.ImageUrlShrink = &v
	return s
}

func (s *CategoryAttributeMatchShrinkRequest) SetItemSpec(v string) *CategoryAttributeMatchShrinkRequest {
	s.ItemSpec = &v
	return s
}

func (s *CategoryAttributeMatchShrinkRequest) SetSku(v string) *CategoryAttributeMatchShrinkRequest {
	s.Sku = &v
	return s
}

func (s *CategoryAttributeMatchShrinkRequest) SetSourceCategory(v string) *CategoryAttributeMatchShrinkRequest {
	s.SourceCategory = &v
	return s
}

func (s *CategoryAttributeMatchShrinkRequest) SetSourcePlatform(v string) *CategoryAttributeMatchShrinkRequest {
	s.SourcePlatform = &v
	return s
}

func (s *CategoryAttributeMatchShrinkRequest) SetTargetPlatform(v string) *CategoryAttributeMatchShrinkRequest {
	s.TargetPlatform = &v
	return s
}

func (s *CategoryAttributeMatchShrinkRequest) SetTitle(v string) *CategoryAttributeMatchShrinkRequest {
	s.Title = &v
	return s
}

func (s *CategoryAttributeMatchShrinkRequest) Validate() error {
	return dara.Validate(s)
}
