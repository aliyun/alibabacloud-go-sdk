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
	// The product description.
	//
	// This parameter is required.
	//
	// example:
	//
	// Silk, Spring New Arrival
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The product attributes that describe the characteristics of the product, such as material. Provide the attribute names and values.
	//
	// example:
	//
	// Material:Polyester,Target Audience:General,Brand:AMASON PET
	ItemSpec *string `json:"ItemSpec,omitempty" xml:"ItemSpec,omitempty"`
	// The SKU title of the product.
	//
	// example:
	//
	// 0:0:Color:Black(without locator);0:1:Color:Smart Locator(without collar);0:2:Color:Range Locator(without collar);0:3:Color:Blue(without locator);0:4:Color:Orange(without locator);0:5:Color:Pink(without locator);0:6:Color:Red(without locator);0:7:Color:Yellow(without locator);0:8:Color:Purple(without locator)
	Sku *string `json:"Sku,omitempty" xml:"Sku,omitempty"`
	// The product category on the source platform.
	//
	// This parameter is required.
	//
	// example:
	//
	// Clothing
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
	// Women\\"s Spring New Arrival Clothing
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
