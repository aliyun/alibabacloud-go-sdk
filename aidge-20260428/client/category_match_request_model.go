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
	// Silk, new spring arrival
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The product attributes and feature descriptions, such as the material of the product. Enter the product attribute name and attribute value.
	//
	// example:
	//
	// Material:Polyester,Target user:Universal,Brand:AMASON PET
	ItemSpec *string `json:"ItemSpec,omitempty" xml:"ItemSpec,omitempty"`
	// The product SKU title.
	//
	// example:
	//
	// 0:0:Color:Black (without tracker);0:1:Color:Smart tracking (without collar);0:2:Color:Range tracking (without collar);0:3:Color:Blue (without tracker);0:4:Color:Orange (without tracker);0:5:Color:Pink (without tracker);0:6:Color:Red (without tracker);0:7:Color:Yellow (without tracker);0:8:Color:Purple (without tracker)
	Sku *string `json:"Sku,omitempty" xml:"Sku,omitempty"`
	// The product category.
	//
	// This parameter is required.
	//
	// example:
	//
	// Clothing
	SourceCategory *string `json:"SourceCategory,omitempty" xml:"SourceCategory,omitempty"`
	// The source platform.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1688
	SourcePlatform *string `json:"SourcePlatform,omitempty" xml:"SourcePlatform,omitempty"`
	// The target platform. Currently, only TEMU is supported.
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
	// Women\\"s new spring clothing
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
