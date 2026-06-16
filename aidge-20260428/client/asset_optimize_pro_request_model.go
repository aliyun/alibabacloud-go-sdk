// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssetOptimizeProRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColumnNameList(v []*string) *AssetOptimizeProRequest
	GetColumnNameList() []*string
	SetGlossary(v string) *AssetOptimizeProRequest
	GetGlossary() *string
	SetIncludingProductArea(v bool) *AssetOptimizeProRequest
	GetIncludingProductArea() *bool
	SetLanguageModel(v string) *AssetOptimizeProRequest
	GetLanguageModel() *string
	SetNeedTrans(v bool) *AssetOptimizeProRequest
	GetNeedTrans() *bool
	SetProductUrl(v string) *AssetOptimizeProRequest
	GetProductUrl() *string
	SetSourceLanguage(v string) *AssetOptimizeProRequest
	GetSourceLanguage() *string
	SetSourcePlatform(v string) *AssetOptimizeProRequest
	GetSourcePlatform() *string
	SetTargetLanguage(v string) *AssetOptimizeProRequest
	GetTargetLanguage() *string
	SetTargetPlatform(v string) *AssetOptimizeProRequest
	GetTargetPlatform() *string
	SetThreshold(v float64) *AssetOptimizeProRequest
	GetThreshold() *float64
	SetTranslatingBrandInTheProduct(v bool) *AssetOptimizeProRequest
	GetTranslatingBrandInTheProduct() *bool
}

type AssetOptimizeProRequest struct {
	// The list of column names to recognize in size chart images. Optional.
	//
	// example:
	//
	// ["胸围","腰围","臀围"]
	ColumnNameList []*string `json:"ColumnNameList,omitempty" xml:"ColumnNameList,omitempty" type:"Repeated"`
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

func (s AssetOptimizeProRequest) String() string {
	return dara.Prettify(s)
}

func (s AssetOptimizeProRequest) GoString() string {
	return s.String()
}

func (s *AssetOptimizeProRequest) GetColumnNameList() []*string {
	return s.ColumnNameList
}

func (s *AssetOptimizeProRequest) GetGlossary() *string {
	return s.Glossary
}

func (s *AssetOptimizeProRequest) GetIncludingProductArea() *bool {
	return s.IncludingProductArea
}

func (s *AssetOptimizeProRequest) GetLanguageModel() *string {
	return s.LanguageModel
}

func (s *AssetOptimizeProRequest) GetNeedTrans() *bool {
	return s.NeedTrans
}

func (s *AssetOptimizeProRequest) GetProductUrl() *string {
	return s.ProductUrl
}

func (s *AssetOptimizeProRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *AssetOptimizeProRequest) GetSourcePlatform() *string {
	return s.SourcePlatform
}

func (s *AssetOptimizeProRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *AssetOptimizeProRequest) GetTargetPlatform() *string {
	return s.TargetPlatform
}

func (s *AssetOptimizeProRequest) GetThreshold() *float64 {
	return s.Threshold
}

func (s *AssetOptimizeProRequest) GetTranslatingBrandInTheProduct() *bool {
	return s.TranslatingBrandInTheProduct
}

func (s *AssetOptimizeProRequest) SetColumnNameList(v []*string) *AssetOptimizeProRequest {
	s.ColumnNameList = v
	return s
}

func (s *AssetOptimizeProRequest) SetGlossary(v string) *AssetOptimizeProRequest {
	s.Glossary = &v
	return s
}

func (s *AssetOptimizeProRequest) SetIncludingProductArea(v bool) *AssetOptimizeProRequest {
	s.IncludingProductArea = &v
	return s
}

func (s *AssetOptimizeProRequest) SetLanguageModel(v string) *AssetOptimizeProRequest {
	s.LanguageModel = &v
	return s
}

func (s *AssetOptimizeProRequest) SetNeedTrans(v bool) *AssetOptimizeProRequest {
	s.NeedTrans = &v
	return s
}

func (s *AssetOptimizeProRequest) SetProductUrl(v string) *AssetOptimizeProRequest {
	s.ProductUrl = &v
	return s
}

func (s *AssetOptimizeProRequest) SetSourceLanguage(v string) *AssetOptimizeProRequest {
	s.SourceLanguage = &v
	return s
}

func (s *AssetOptimizeProRequest) SetSourcePlatform(v string) *AssetOptimizeProRequest {
	s.SourcePlatform = &v
	return s
}

func (s *AssetOptimizeProRequest) SetTargetLanguage(v string) *AssetOptimizeProRequest {
	s.TargetLanguage = &v
	return s
}

func (s *AssetOptimizeProRequest) SetTargetPlatform(v string) *AssetOptimizeProRequest {
	s.TargetPlatform = &v
	return s
}

func (s *AssetOptimizeProRequest) SetThreshold(v float64) *AssetOptimizeProRequest {
	s.Threshold = &v
	return s
}

func (s *AssetOptimizeProRequest) SetTranslatingBrandInTheProduct(v bool) *AssetOptimizeProRequest {
	s.TranslatingBrandInTheProduct = &v
	return s
}

func (s *AssetOptimizeProRequest) Validate() error {
	return dara.Validate(s)
}
