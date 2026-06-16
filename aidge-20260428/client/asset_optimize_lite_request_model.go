// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssetOptimizeLiteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGlossary(v string) *AssetOptimizeLiteRequest
	GetGlossary() *string
	SetIncludingProductArea(v bool) *AssetOptimizeLiteRequest
	GetIncludingProductArea() *bool
	SetNeedTrans(v bool) *AssetOptimizeLiteRequest
	GetNeedTrans() *bool
	SetProductUrl(v string) *AssetOptimizeLiteRequest
	GetProductUrl() *string
	SetSourceLanguage(v string) *AssetOptimizeLiteRequest
	GetSourceLanguage() *string
	SetSourcePlatform(v string) *AssetOptimizeLiteRequest
	GetSourcePlatform() *string
	SetTargetLanguage(v string) *AssetOptimizeLiteRequest
	GetTargetLanguage() *string
	SetTargetPlatform(v string) *AssetOptimizeLiteRequest
	GetTargetPlatform() *string
	SetTranslatingBrandInTheProduct(v bool) *AssetOptimizeLiteRequest
	GetTranslatingBrandInTheProduct() *bool
}

type AssetOptimizeLiteRequest struct {
	// Custom glossary
	//
	// example:
	//
	// 测试
	Glossary *string `json:"Glossary,omitempty" xml:"Glossary,omitempty"`
	// Whether to include product region translation
	//
	// example:
	//
	// false
	IncludingProductArea *bool `json:"IncludingProductArea,omitempty" xml:"IncludingProductArea,omitempty"`
	// Whether translation is required
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	NeedTrans *bool `json:"NeedTrans,omitempty" xml:"NeedTrans,omitempty"`
	// Product Link URL
	//
	// This parameter is required.
	//
	// example:
	//
	// https://detail.1688.com/offer/631960323816.html
	ProductUrl *string `json:"ProductUrl,omitempty" xml:"ProductUrl,omitempty"`
	// Source language code, e.g., zh
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// Source platform, e.g., 1688
	//
	// This parameter is required.
	//
	// example:
	//
	// 1688
	SourcePlatform *string `json:"SourcePlatform,omitempty" xml:"SourcePlatform,omitempty"`
	// Target language code, e.g., en
	//
	// example:
	//
	// en
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	// Target platform, e.g., temu
	//
	// This parameter is required.
	//
	// example:
	//
	// temu
	TargetPlatform *string `json:"TargetPlatform,omitempty" xml:"TargetPlatform,omitempty"`
	// Whether to translate brand names in images, default false
	//
	// example:
	//
	// false
	TranslatingBrandInTheProduct *bool `json:"TranslatingBrandInTheProduct,omitempty" xml:"TranslatingBrandInTheProduct,omitempty"`
}

func (s AssetOptimizeLiteRequest) String() string {
	return dara.Prettify(s)
}

func (s AssetOptimizeLiteRequest) GoString() string {
	return s.String()
}

func (s *AssetOptimizeLiteRequest) GetGlossary() *string {
	return s.Glossary
}

func (s *AssetOptimizeLiteRequest) GetIncludingProductArea() *bool {
	return s.IncludingProductArea
}

func (s *AssetOptimizeLiteRequest) GetNeedTrans() *bool {
	return s.NeedTrans
}

func (s *AssetOptimizeLiteRequest) GetProductUrl() *string {
	return s.ProductUrl
}

func (s *AssetOptimizeLiteRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *AssetOptimizeLiteRequest) GetSourcePlatform() *string {
	return s.SourcePlatform
}

func (s *AssetOptimizeLiteRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *AssetOptimizeLiteRequest) GetTargetPlatform() *string {
	return s.TargetPlatform
}

func (s *AssetOptimizeLiteRequest) GetTranslatingBrandInTheProduct() *bool {
	return s.TranslatingBrandInTheProduct
}

func (s *AssetOptimizeLiteRequest) SetGlossary(v string) *AssetOptimizeLiteRequest {
	s.Glossary = &v
	return s
}

func (s *AssetOptimizeLiteRequest) SetIncludingProductArea(v bool) *AssetOptimizeLiteRequest {
	s.IncludingProductArea = &v
	return s
}

func (s *AssetOptimizeLiteRequest) SetNeedTrans(v bool) *AssetOptimizeLiteRequest {
	s.NeedTrans = &v
	return s
}

func (s *AssetOptimizeLiteRequest) SetProductUrl(v string) *AssetOptimizeLiteRequest {
	s.ProductUrl = &v
	return s
}

func (s *AssetOptimizeLiteRequest) SetSourceLanguage(v string) *AssetOptimizeLiteRequest {
	s.SourceLanguage = &v
	return s
}

func (s *AssetOptimizeLiteRequest) SetSourcePlatform(v string) *AssetOptimizeLiteRequest {
	s.SourcePlatform = &v
	return s
}

func (s *AssetOptimizeLiteRequest) SetTargetLanguage(v string) *AssetOptimizeLiteRequest {
	s.TargetLanguage = &v
	return s
}

func (s *AssetOptimizeLiteRequest) SetTargetPlatform(v string) *AssetOptimizeLiteRequest {
	s.TargetPlatform = &v
	return s
}

func (s *AssetOptimizeLiteRequest) SetTranslatingBrandInTheProduct(v bool) *AssetOptimizeLiteRequest {
	s.TranslatingBrandInTheProduct = &v
	return s
}

func (s *AssetOptimizeLiteRequest) Validate() error {
	return dara.Validate(s)
}
