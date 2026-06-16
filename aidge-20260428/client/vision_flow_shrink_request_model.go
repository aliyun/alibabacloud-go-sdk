// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVisionFlowShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAbilityShrink(v string) *VisionFlowShrinkRequest
	GetAbilityShrink() *string
	SetBackGroundType(v string) *VisionFlowShrinkRequest
	GetBackGroundType() *string
	SetGlossary(v string) *VisionFlowShrinkRequest
	GetGlossary() *string
	SetImageUrl(v string) *VisionFlowShrinkRequest
	GetImageUrl() *string
	SetIncludingProductArea(v bool) *VisionFlowShrinkRequest
	GetIncludingProductArea() *bool
	SetIsFilter(v bool) *VisionFlowShrinkRequest
	GetIsFilter() *bool
	SetMask(v string) *VisionFlowShrinkRequest
	GetMask() *string
	SetNonobjectDetectElementsShrink(v string) *VisionFlowShrinkRequest
	GetNonobjectDetectElementsShrink() *string
	SetNonobjectRemoveElementsShrink(v string) *VisionFlowShrinkRequest
	GetNonobjectRemoveElementsShrink() *string
	SetObjectDetectElementsShrink(v string) *VisionFlowShrinkRequest
	GetObjectDetectElementsShrink() *string
	SetObjectRemoveElementsShrink(v string) *VisionFlowShrinkRequest
	GetObjectRemoveElementsShrink() *string
	SetSourceLanguage(v string) *VisionFlowShrinkRequest
	GetSourceLanguage() *string
	SetTargetHeight(v int32) *VisionFlowShrinkRequest
	GetTargetHeight() *int32
	SetTargetLanguage(v string) *VisionFlowShrinkRequest
	GetTargetLanguage() *string
	SetTargetWidth(v int32) *VisionFlowShrinkRequest
	GetTargetWidth() *int32
	SetTranslatingBrandInTheProduct(v bool) *VisionFlowShrinkRequest
	GetTranslatingBrandInTheProduct() *bool
	SetUpscaleFactor(v int32) *VisionFlowShrinkRequest
	GetUpscaleFactor() *int32
}

type VisionFlowShrinkRequest struct {
	// The AI capabilities to apply (1 = intelligent element detection, 2 = intelligent matting, 3 = intelligent removal, 4 = Image Translation Pro, 5 = intelligent cropping, 6 = HD upscaling). Multiple selections allowed.
	//
	// This parameter is required.
	//
	// example:
	//
	// [1,2,3,4]
	AbilityShrink *string `json:"Ability,omitempty" xml:"Ability,omitempty"`
	// The background type of the returned image. Valid values: WHITE_BACKGROUND (white background) and TRANSPARENT (transparent background). Required when the intelligent matting capability is selected.
	//
	// example:
	//
	// WHITE_BACKGROUND
	BackGroundType *string `json:"BackGroundType,omitempty" xml:"BackGroundType,omitempty"`
	// The intervention glossary ID. Optional. Create a glossary separately in the console and provide its ID. If left empty, translation results are not modified.
	//
	// example:
	//
	// glossary_1
	Glossary *string `json:"Glossary,omitempty" xml:"Glossary,omitempty"`
	// The URL of the image to process. Required. The resolution must be greater than 256 × 256, the long side must not exceed 1920 pixels, and the short side must not exceed 1080 pixels. The file size must not exceed 5 MB. Supported formats: png, jpeg, jpg, bmp, and webp.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://ae01.alicdn.com/kf/S342f0070dc9f4be09a6cbed34e90dc8fs.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// Specifies whether to translate text on the image subject. Optional. Default value: false. Helps protect embedded information such as product names from being translated.
	//
	// example:
	//
	// false
	IncludingProductArea *bool `json:"IncludingProductArea,omitempty" xml:"IncludingProductArea,omitempty"`
	// Specifies whether images with the detected elements proceed to subsequent processing. A value of true indicates that images containing the elements proceed to subsequent processing. A value of false indicates that they do not. Required when the intelligent element detection capability is selected.
	//
	// example:
	//
	// true
	IsFilter *bool `json:"IsFilter,omitempty" xml:"IsFilter,omitempty"`
	// The specific removal area in RLE format. Optional. If provided, this parameter takes priority and the ObjectRemoveElements and NonobjectRemoveElements parameters are ignored.
	//
	// example:
	//
	// 474556 160 475356 160
	Mask *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	// The elements to detect on the non-subject area of the image (1 = watermark, 2 = logo, 3 = text, 4 = text-bearing color block). Multiple selections allowed. When the intelligent element detection capability is selected, at least one of NonobjectDetectElements and ObjectDetectElements is required.
	//
	// example:
	//
	// [1,2,3]
	NonobjectDetectElementsShrink *string `json:"NonobjectDetectElements,omitempty" xml:"NonobjectDetectElements,omitempty"`
	// The elements to remove from the non-subject area of the image (1 = transparent text block, 2 = specific name, 3 = text, 4 = overlay patch). Multiple selections allowed. When the intelligent removal capability is selected, at least one of NonobjectRemoveElements and ObjectRemoveElements is required.
	//
	// example:
	//
	// [1,2,4]
	NonobjectRemoveElementsShrink *string `json:"NonobjectRemoveElements,omitempty" xml:"NonobjectRemoveElements,omitempty"`
	// The elements to detect on the image subject (1 = watermark, 2 = logo, 3 = text, 4 = text-bearing color block). Multiple selections allowed. When the intelligent element detection capability is selected, at least one of ObjectDetectElements and NonobjectDetectElements is required.
	//
	// example:
	//
	// [1,2,3,4]
	ObjectDetectElementsShrink *string `json:"ObjectDetectElements,omitempty" xml:"ObjectDetectElements,omitempty"`
	// The elements to remove from the image subject (1 = transparent text block, 2 = specific name, 3 = text, 4 = overlay patch). Multiple selections allowed. When the intelligent removal capability is selected, at least one of ObjectRemoveElements and NonobjectRemoveElements is required.
	//
	// example:
	//
	// [1,2,4]
	ObjectRemoveElementsShrink *string `json:"ObjectRemoveElements,omitempty" xml:"ObjectRemoveElements,omitempty"`
	// The source language code. Optional. For supported language pairs, see the supported translation language pairs list.
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// The desired height of the cropped image, in pixels. Valid values: 100 to 5000. Required when the intelligent cropping capability is selected.
	//
	// example:
	//
	// 800
	TargetHeight *int32 `json:"TargetHeight,omitempty" xml:"TargetHeight,omitempty"`
	// The target language code. Optional. For supported language pairs, see the supported translation language pairs list.
	//
	// example:
	//
	// en
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	// The desired width of the cropped image, in pixels. Valid values: 100 to 5000. Required when the intelligent cropping capability is selected.
	//
	// example:
	//
	// 800
	TargetWidth *int32 `json:"TargetWidth,omitempty" xml:"TargetWidth,omitempty"`
	// Specifies whether to translate brand names on the image. Optional. Default value: false. Helps protect brand name information from being translated.
	//
	// example:
	//
	// false
	TranslatingBrandInTheProduct *bool `json:"TranslatingBrandInTheProduct,omitempty" xml:"TranslatingBrandInTheProduct,omitempty"`
	// The image upscaling factor. Optional. Default value: 2. Valid values: 2 to 4. Required when the HD upscaling capability is selected.
	//
	// example:
	//
	// 2
	UpscaleFactor *int32 `json:"UpscaleFactor,omitempty" xml:"UpscaleFactor,omitempty"`
}

func (s VisionFlowShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s VisionFlowShrinkRequest) GoString() string {
	return s.String()
}

func (s *VisionFlowShrinkRequest) GetAbilityShrink() *string {
	return s.AbilityShrink
}

func (s *VisionFlowShrinkRequest) GetBackGroundType() *string {
	return s.BackGroundType
}

func (s *VisionFlowShrinkRequest) GetGlossary() *string {
	return s.Glossary
}

func (s *VisionFlowShrinkRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *VisionFlowShrinkRequest) GetIncludingProductArea() *bool {
	return s.IncludingProductArea
}

func (s *VisionFlowShrinkRequest) GetIsFilter() *bool {
	return s.IsFilter
}

func (s *VisionFlowShrinkRequest) GetMask() *string {
	return s.Mask
}

func (s *VisionFlowShrinkRequest) GetNonobjectDetectElementsShrink() *string {
	return s.NonobjectDetectElementsShrink
}

func (s *VisionFlowShrinkRequest) GetNonobjectRemoveElementsShrink() *string {
	return s.NonobjectRemoveElementsShrink
}

func (s *VisionFlowShrinkRequest) GetObjectDetectElementsShrink() *string {
	return s.ObjectDetectElementsShrink
}

func (s *VisionFlowShrinkRequest) GetObjectRemoveElementsShrink() *string {
	return s.ObjectRemoveElementsShrink
}

func (s *VisionFlowShrinkRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *VisionFlowShrinkRequest) GetTargetHeight() *int32 {
	return s.TargetHeight
}

func (s *VisionFlowShrinkRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *VisionFlowShrinkRequest) GetTargetWidth() *int32 {
	return s.TargetWidth
}

func (s *VisionFlowShrinkRequest) GetTranslatingBrandInTheProduct() *bool {
	return s.TranslatingBrandInTheProduct
}

func (s *VisionFlowShrinkRequest) GetUpscaleFactor() *int32 {
	return s.UpscaleFactor
}

func (s *VisionFlowShrinkRequest) SetAbilityShrink(v string) *VisionFlowShrinkRequest {
	s.AbilityShrink = &v
	return s
}

func (s *VisionFlowShrinkRequest) SetBackGroundType(v string) *VisionFlowShrinkRequest {
	s.BackGroundType = &v
	return s
}

func (s *VisionFlowShrinkRequest) SetGlossary(v string) *VisionFlowShrinkRequest {
	s.Glossary = &v
	return s
}

func (s *VisionFlowShrinkRequest) SetImageUrl(v string) *VisionFlowShrinkRequest {
	s.ImageUrl = &v
	return s
}

func (s *VisionFlowShrinkRequest) SetIncludingProductArea(v bool) *VisionFlowShrinkRequest {
	s.IncludingProductArea = &v
	return s
}

func (s *VisionFlowShrinkRequest) SetIsFilter(v bool) *VisionFlowShrinkRequest {
	s.IsFilter = &v
	return s
}

func (s *VisionFlowShrinkRequest) SetMask(v string) *VisionFlowShrinkRequest {
	s.Mask = &v
	return s
}

func (s *VisionFlowShrinkRequest) SetNonobjectDetectElementsShrink(v string) *VisionFlowShrinkRequest {
	s.NonobjectDetectElementsShrink = &v
	return s
}

func (s *VisionFlowShrinkRequest) SetNonobjectRemoveElementsShrink(v string) *VisionFlowShrinkRequest {
	s.NonobjectRemoveElementsShrink = &v
	return s
}

func (s *VisionFlowShrinkRequest) SetObjectDetectElementsShrink(v string) *VisionFlowShrinkRequest {
	s.ObjectDetectElementsShrink = &v
	return s
}

func (s *VisionFlowShrinkRequest) SetObjectRemoveElementsShrink(v string) *VisionFlowShrinkRequest {
	s.ObjectRemoveElementsShrink = &v
	return s
}

func (s *VisionFlowShrinkRequest) SetSourceLanguage(v string) *VisionFlowShrinkRequest {
	s.SourceLanguage = &v
	return s
}

func (s *VisionFlowShrinkRequest) SetTargetHeight(v int32) *VisionFlowShrinkRequest {
	s.TargetHeight = &v
	return s
}

func (s *VisionFlowShrinkRequest) SetTargetLanguage(v string) *VisionFlowShrinkRequest {
	s.TargetLanguage = &v
	return s
}

func (s *VisionFlowShrinkRequest) SetTargetWidth(v int32) *VisionFlowShrinkRequest {
	s.TargetWidth = &v
	return s
}

func (s *VisionFlowShrinkRequest) SetTranslatingBrandInTheProduct(v bool) *VisionFlowShrinkRequest {
	s.TranslatingBrandInTheProduct = &v
	return s
}

func (s *VisionFlowShrinkRequest) SetUpscaleFactor(v int32) *VisionFlowShrinkRequest {
	s.UpscaleFactor = &v
	return s
}

func (s *VisionFlowShrinkRequest) Validate() error {
	return dara.Validate(s)
}
