// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssetOptimizeProShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColumnNameListShrink(v string) *AssetOptimizeProShrinkRequest
	GetColumnNameListShrink() *string
	SetGlossary(v string) *AssetOptimizeProShrinkRequest
	GetGlossary() *string
	SetIncludingProductArea(v bool) *AssetOptimizeProShrinkRequest
	GetIncludingProductArea() *bool
	SetLanguageModel(v string) *AssetOptimizeProShrinkRequest
	GetLanguageModel() *string
	SetNeedTrans(v bool) *AssetOptimizeProShrinkRequest
	GetNeedTrans() *bool
	SetProductUrl(v string) *AssetOptimizeProShrinkRequest
	GetProductUrl() *string
	SetSourceLanguage(v string) *AssetOptimizeProShrinkRequest
	GetSourceLanguage() *string
	SetSourcePlatform(v string) *AssetOptimizeProShrinkRequest
	GetSourcePlatform() *string
	SetTargetLanguage(v string) *AssetOptimizeProShrinkRequest
	GetTargetLanguage() *string
	SetTargetPlatform(v string) *AssetOptimizeProShrinkRequest
	GetTargetPlatform() *string
	SetThreshold(v float64) *AssetOptimizeProShrinkRequest
	GetThreshold() *float64
	SetTranslatingBrandInTheProduct(v bool) *AssetOptimizeProShrinkRequest
	GetTranslatingBrandInTheProduct() *bool
}

type AssetOptimizeProShrinkRequest struct {
	// The list of column names to recognize in size chart images. Optional.
	//
	// example:
	//
	// ["胸围","腰围","臀围"]
	ColumnNameListShrink *string `json:"ColumnNameList,omitempty" xml:"ColumnNameList,omitempty"`
	// The glossary ID. Optional. Create a glossary in the console and provide its ID. If left empty, translation results are not modified by any glossary.
	//
	// example:
	//
	// glossary_1
	Glossary *string `json:"Glossary,omitempty" xml:"Glossary,omitempty"`
	// Specifies whether to translate text on the product subject area of images. Setting this to false helps protect embedded information such as product names from being translated. Default value: false.
	//
	// example:
	//
	// false
	IncludingProductArea *bool `json:"IncludingProductArea,omitempty" xml:"IncludingProductArea,omitempty"`
	// The output language format for size chart images. If not specified, the original format is used. Set to en for English output or cn for Chinese output.
	//
	// example:
	//
	// cn
	LanguageModel *string `json:"LanguageModel,omitempty" xml:"LanguageModel,omitempty"`
	// Specifies whether translation is required (true/false). If set to true, SourceLanguage and TargetLanguage are required.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	NeedTrans *bool `json:"NeedTrans,omitempty" xml:"NeedTrans,omitempty"`
	// The product URL. This parameter is required. Only 1688 product links are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://detail.1688.com/offer/771024907536.html
	ProductUrl *string `json:"ProductUrl,omitempty" xml:"ProductUrl,omitempty"`
	// The source language code. Optional. For supported language pairs, refer to the supported translation language list. This parameter is required if NeedTrans is set to true.
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// The source platform. Only 1688 is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1688
	SourcePlatform *string `json:"SourcePlatform,omitempty" xml:"SourcePlatform,omitempty"`
	// The target language code. Optional. For supported language pairs, refer to the supported translation language list. This parameter is required if NeedTrans is set to true.
	//
	// example:
	//
	// en
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	// The target listing platform. Only temu is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// temu
	TargetPlatform *string `json:"TargetPlatform,omitempty" xml:"TargetPlatform,omitempty"`
	// The confidence threshold for size chart detection. Default value: 0.4. A value of 0 treats all images as size charts. A value of 1 treats no images as size charts.
	//
	// example:
	//
	// 0.4
	Threshold *float64 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// Specifies whether to translate brand names on images. Optional. Default value: false. Setting this to false helps protect brand name information from being translated.
	//
	// example:
	//
	// false
	TranslatingBrandInTheProduct *bool `json:"TranslatingBrandInTheProduct,omitempty" xml:"TranslatingBrandInTheProduct,omitempty"`
}

func (s AssetOptimizeProShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AssetOptimizeProShrinkRequest) GoString() string {
	return s.String()
}

func (s *AssetOptimizeProShrinkRequest) GetColumnNameListShrink() *string {
	return s.ColumnNameListShrink
}

func (s *AssetOptimizeProShrinkRequest) GetGlossary() *string {
	return s.Glossary
}

func (s *AssetOptimizeProShrinkRequest) GetIncludingProductArea() *bool {
	return s.IncludingProductArea
}

func (s *AssetOptimizeProShrinkRequest) GetLanguageModel() *string {
	return s.LanguageModel
}

func (s *AssetOptimizeProShrinkRequest) GetNeedTrans() *bool {
	return s.NeedTrans
}

func (s *AssetOptimizeProShrinkRequest) GetProductUrl() *string {
	return s.ProductUrl
}

func (s *AssetOptimizeProShrinkRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *AssetOptimizeProShrinkRequest) GetSourcePlatform() *string {
	return s.SourcePlatform
}

func (s *AssetOptimizeProShrinkRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *AssetOptimizeProShrinkRequest) GetTargetPlatform() *string {
	return s.TargetPlatform
}

func (s *AssetOptimizeProShrinkRequest) GetThreshold() *float64 {
	return s.Threshold
}

func (s *AssetOptimizeProShrinkRequest) GetTranslatingBrandInTheProduct() *bool {
	return s.TranslatingBrandInTheProduct
}

func (s *AssetOptimizeProShrinkRequest) SetColumnNameListShrink(v string) *AssetOptimizeProShrinkRequest {
	s.ColumnNameListShrink = &v
	return s
}

func (s *AssetOptimizeProShrinkRequest) SetGlossary(v string) *AssetOptimizeProShrinkRequest {
	s.Glossary = &v
	return s
}

func (s *AssetOptimizeProShrinkRequest) SetIncludingProductArea(v bool) *AssetOptimizeProShrinkRequest {
	s.IncludingProductArea = &v
	return s
}

func (s *AssetOptimizeProShrinkRequest) SetLanguageModel(v string) *AssetOptimizeProShrinkRequest {
	s.LanguageModel = &v
	return s
}

func (s *AssetOptimizeProShrinkRequest) SetNeedTrans(v bool) *AssetOptimizeProShrinkRequest {
	s.NeedTrans = &v
	return s
}

func (s *AssetOptimizeProShrinkRequest) SetProductUrl(v string) *AssetOptimizeProShrinkRequest {
	s.ProductUrl = &v
	return s
}

func (s *AssetOptimizeProShrinkRequest) SetSourceLanguage(v string) *AssetOptimizeProShrinkRequest {
	s.SourceLanguage = &v
	return s
}

func (s *AssetOptimizeProShrinkRequest) SetSourcePlatform(v string) *AssetOptimizeProShrinkRequest {
	s.SourcePlatform = &v
	return s
}

func (s *AssetOptimizeProShrinkRequest) SetTargetLanguage(v string) *AssetOptimizeProShrinkRequest {
	s.TargetLanguage = &v
	return s
}

func (s *AssetOptimizeProShrinkRequest) SetTargetPlatform(v string) *AssetOptimizeProShrinkRequest {
	s.TargetPlatform = &v
	return s
}

func (s *AssetOptimizeProShrinkRequest) SetThreshold(v float64) *AssetOptimizeProShrinkRequest {
	s.Threshold = &v
	return s
}

func (s *AssetOptimizeProShrinkRequest) SetTranslatingBrandInTheProduct(v bool) *AssetOptimizeProShrinkRequest {
	s.TranslatingBrandInTheProduct = &v
	return s
}

func (s *AssetOptimizeProShrinkRequest) Validate() error {
	return dara.Validate(s)
}
