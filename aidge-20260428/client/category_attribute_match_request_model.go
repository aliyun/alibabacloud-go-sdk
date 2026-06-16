// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCategoryAttributeMatchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CategoryAttributeMatchRequest
	GetDescription() *string
	SetImageUrl(v []*string) *CategoryAttributeMatchRequest
	GetImageUrl() []*string
	SetItemSpec(v string) *CategoryAttributeMatchRequest
	GetItemSpec() *string
	SetSku(v string) *CategoryAttributeMatchRequest
	GetSku() *string
	SetSourceCategory(v string) *CategoryAttributeMatchRequest
	GetSourceCategory() *string
	SetSourcePlatform(v string) *CategoryAttributeMatchRequest
	GetSourcePlatform() *string
	SetTargetPlatform(v string) *CategoryAttributeMatchRequest
	GetTargetPlatform() *string
	SetTitle(v string) *CategoryAttributeMatchRequest
	GetTitle() *string
}

type CategoryAttributeMatchRequest struct {
	// The product details.
	//
	// example:
	//
	// 真丝，春季新款
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The product image URLs. A maximum of 10 images are supported.
	//
	// example:
	//
	// https://frametour-assets.oss-cn-shanghai.aliyuncs.com/user-faces/viid_face/dd0dd06c-9351-4e5f-bc70-24166a754d7f.jpg
	ImageUrl []*string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty" type:"Repeated"`
	// The product attributes that describe the product characteristics.
	//
	// This parameter is required.
	//
	// example:
	//
	// 例如商品的材质等。输入商品属性名称和属性内容
	ItemSpec *string `json:"ItemSpec,omitempty" xml:"ItemSpec,omitempty"`
	// The product SKU title.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0:0:颜色:黑(不含定位器);0:1:颜色:智能定位(不含项圈);0:2:颜色:范围定位(不含项圈);0:3:颜色:蓝(不含定位器);0:4:颜色:橙(不含定位器);0:5:颜色:粉(不含定位器);0:6:颜色:红(不含定位器);0:7:颜色:黄(不含定位器);0:8:颜色:紫(不含定位器)
	Sku *string `json:"Sku,omitempty" xml:"Sku,omitempty"`
	// The product category on the source platform.
	//
	// This parameter is required.
	//
	// example:
	//
	// 衣服
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
	// 女士春季新款衣服
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CategoryAttributeMatchRequest) String() string {
	return dara.Prettify(s)
}

func (s CategoryAttributeMatchRequest) GoString() string {
	return s.String()
}

func (s *CategoryAttributeMatchRequest) GetDescription() *string {
	return s.Description
}

func (s *CategoryAttributeMatchRequest) GetImageUrl() []*string {
	return s.ImageUrl
}

func (s *CategoryAttributeMatchRequest) GetItemSpec() *string {
	return s.ItemSpec
}

func (s *CategoryAttributeMatchRequest) GetSku() *string {
	return s.Sku
}

func (s *CategoryAttributeMatchRequest) GetSourceCategory() *string {
	return s.SourceCategory
}

func (s *CategoryAttributeMatchRequest) GetSourcePlatform() *string {
	return s.SourcePlatform
}

func (s *CategoryAttributeMatchRequest) GetTargetPlatform() *string {
	return s.TargetPlatform
}

func (s *CategoryAttributeMatchRequest) GetTitle() *string {
	return s.Title
}

func (s *CategoryAttributeMatchRequest) SetDescription(v string) *CategoryAttributeMatchRequest {
	s.Description = &v
	return s
}

func (s *CategoryAttributeMatchRequest) SetImageUrl(v []*string) *CategoryAttributeMatchRequest {
	s.ImageUrl = v
	return s
}

func (s *CategoryAttributeMatchRequest) SetItemSpec(v string) *CategoryAttributeMatchRequest {
	s.ItemSpec = &v
	return s
}

func (s *CategoryAttributeMatchRequest) SetSku(v string) *CategoryAttributeMatchRequest {
	s.Sku = &v
	return s
}

func (s *CategoryAttributeMatchRequest) SetSourceCategory(v string) *CategoryAttributeMatchRequest {
	s.SourceCategory = &v
	return s
}

func (s *CategoryAttributeMatchRequest) SetSourcePlatform(v string) *CategoryAttributeMatchRequest {
	s.SourcePlatform = &v
	return s
}

func (s *CategoryAttributeMatchRequest) SetTargetPlatform(v string) *CategoryAttributeMatchRequest {
	s.TargetPlatform = &v
	return s
}

func (s *CategoryAttributeMatchRequest) SetTitle(v string) *CategoryAttributeMatchRequest {
	s.Title = &v
	return s
}

func (s *CategoryAttributeMatchRequest) Validate() error {
	return dara.Validate(s)
}
