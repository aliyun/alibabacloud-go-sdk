// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCategoryMatchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CategoryMatchRequest
	GetDescription() *string
	SetItemSpec(v string) *CategoryMatchRequest
	GetItemSpec() *string
	SetSku(v string) *CategoryMatchRequest
	GetSku() *string
	SetSourceCategory(v string) *CategoryMatchRequest
	GetSourceCategory() *string
	SetSourcePlatform(v string) *CategoryMatchRequest
	GetSourcePlatform() *string
	SetTargetPlatform(v string) *CategoryMatchRequest
	GetTargetPlatform() *string
	SetTitle(v string) *CategoryMatchRequest
	GetTitle() *string
}

type CategoryMatchRequest struct {
	// The product details.
	//
	// This parameter is required.
	//
	// example:
	//
	// 真丝，春季新款
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The product attributes that describe the characteristics of the product, such as material. Specify the attribute names and attribute values.
	//
	// example:
	//
	// 材质:涤纶,适用对象:通用,品牌:艾马逊AMASON PET
	ItemSpec *string `json:"ItemSpec,omitempty" xml:"ItemSpec,omitempty"`
	// The SKU title of the product.
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
	// The source platform from which products are sourced.
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

func (s CategoryMatchRequest) String() string {
	return dara.Prettify(s)
}

func (s CategoryMatchRequest) GoString() string {
	return s.String()
}

func (s *CategoryMatchRequest) GetDescription() *string {
	return s.Description
}

func (s *CategoryMatchRequest) GetItemSpec() *string {
	return s.ItemSpec
}

func (s *CategoryMatchRequest) GetSku() *string {
	return s.Sku
}

func (s *CategoryMatchRequest) GetSourceCategory() *string {
	return s.SourceCategory
}

func (s *CategoryMatchRequest) GetSourcePlatform() *string {
	return s.SourcePlatform
}

func (s *CategoryMatchRequest) GetTargetPlatform() *string {
	return s.TargetPlatform
}

func (s *CategoryMatchRequest) GetTitle() *string {
	return s.Title
}

func (s *CategoryMatchRequest) SetDescription(v string) *CategoryMatchRequest {
	s.Description = &v
	return s
}

func (s *CategoryMatchRequest) SetItemSpec(v string) *CategoryMatchRequest {
	s.ItemSpec = &v
	return s
}

func (s *CategoryMatchRequest) SetSku(v string) *CategoryMatchRequest {
	s.Sku = &v
	return s
}

func (s *CategoryMatchRequest) SetSourceCategory(v string) *CategoryMatchRequest {
	s.SourceCategory = &v
	return s
}

func (s *CategoryMatchRequest) SetSourcePlatform(v string) *CategoryMatchRequest {
	s.SourcePlatform = &v
	return s
}

func (s *CategoryMatchRequest) SetTargetPlatform(v string) *CategoryMatchRequest {
	s.TargetPlatform = &v
	return s
}

func (s *CategoryMatchRequest) SetTitle(v string) *CategoryMatchRequest {
	s.Title = &v
	return s
}

func (s *CategoryMatchRequest) Validate() error {
	return dara.Validate(s)
}
