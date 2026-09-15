// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitProductMatchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBrandName(v string) *SubmitProductMatchRequest
	GetBrandName() *string
	SetCategory(v string) *SubmitProductMatchRequest
	GetCategory() *string
	SetImageUrl(v string) *SubmitProductMatchRequest
	GetImageUrl() *string
	SetItemId(v string) *SubmitProductMatchRequest
	GetItemId() *string
	SetProductUrl(v string) *SubmitProductMatchRequest
	GetProductUrl() *string
	SetShopName(v string) *SubmitProductMatchRequest
	GetShopName() *string
	SetTitle(v string) *SubmitProductMatchRequest
	GetTitle() *string
}

type SubmitProductMatchRequest struct {
	// The product brand. If this value is not specified, the system attempts to extract the brand from the shop name.
	//
	// example:
	//
	// FILA
	BrandName *string `json:"BrandName,omitempty" xml:"BrandName,omitempty"`
	// The Miaojie product category. Currently used for extension and auditing purposes.
	//
	// example:
	//
	// Children\\"s Shoes
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The HTTP or HTTPS URL of the product main image.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/items/228516909/main.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The Miaojie product ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 228516909
	ItemId *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	// The HTTP or HTTPS URL of the product detail page.
	//
	// example:
	//
	// https://example.com/items/228516909
	ProductUrl *string `json:"ProductUrl,omitempty" xml:"ProductUrl,omitempty"`
	// The shop name. This value is also used as the extraction source when the brand name is missing.
	//
	// This parameter is required.
	//
	// example:
	//
	// FILA斐乐官方旗舰店
	ShopName *string `json:"ShopName,omitempty" xml:"ShopName,omitempty"`
	// The product title.
	//
	// This parameter is required.
	//
	// example:
	//
	// FILA Kids Training Shoes
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s SubmitProductMatchRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitProductMatchRequest) GoString() string {
	return s.String()
}

func (s *SubmitProductMatchRequest) GetBrandName() *string {
	return s.BrandName
}

func (s *SubmitProductMatchRequest) GetCategory() *string {
	return s.Category
}

func (s *SubmitProductMatchRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *SubmitProductMatchRequest) GetItemId() *string {
	return s.ItemId
}

func (s *SubmitProductMatchRequest) GetProductUrl() *string {
	return s.ProductUrl
}

func (s *SubmitProductMatchRequest) GetShopName() *string {
	return s.ShopName
}

func (s *SubmitProductMatchRequest) GetTitle() *string {
	return s.Title
}

func (s *SubmitProductMatchRequest) SetBrandName(v string) *SubmitProductMatchRequest {
	s.BrandName = &v
	return s
}

func (s *SubmitProductMatchRequest) SetCategory(v string) *SubmitProductMatchRequest {
	s.Category = &v
	return s
}

func (s *SubmitProductMatchRequest) SetImageUrl(v string) *SubmitProductMatchRequest {
	s.ImageUrl = &v
	return s
}

func (s *SubmitProductMatchRequest) SetItemId(v string) *SubmitProductMatchRequest {
	s.ItemId = &v
	return s
}

func (s *SubmitProductMatchRequest) SetProductUrl(v string) *SubmitProductMatchRequest {
	s.ProductUrl = &v
	return s
}

func (s *SubmitProductMatchRequest) SetShopName(v string) *SubmitProductMatchRequest {
	s.ShopName = &v
	return s
}

func (s *SubmitProductMatchRequest) SetTitle(v string) *SubmitProductMatchRequest {
	s.Title = &v
	return s
}

func (s *SubmitProductMatchRequest) Validate() error {
	return dara.Validate(s)
}
