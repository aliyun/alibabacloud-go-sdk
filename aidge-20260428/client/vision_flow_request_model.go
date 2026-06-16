// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVisionFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAbility(v []*int32) *VisionFlowRequest
	GetAbility() []*int32
	SetBackGroundType(v string) *VisionFlowRequest
	GetBackGroundType() *string
	SetGlossary(v string) *VisionFlowRequest
	GetGlossary() *string
	SetImageUrl(v string) *VisionFlowRequest
	GetImageUrl() *string
	SetIncludingProductArea(v bool) *VisionFlowRequest
	GetIncludingProductArea() *bool
	SetIsFilter(v bool) *VisionFlowRequest
	GetIsFilter() *bool
	SetMask(v string) *VisionFlowRequest
	GetMask() *string
	SetNonobjectDetectElements(v []*int32) *VisionFlowRequest
	GetNonobjectDetectElements() []*int32
	SetNonobjectRemoveElements(v []*int32) *VisionFlowRequest
	GetNonobjectRemoveElements() []*int32
	SetObjectDetectElements(v []*int32) *VisionFlowRequest
	GetObjectDetectElements() []*int32
	SetObjectRemoveElements(v []*int32) *VisionFlowRequest
	GetObjectRemoveElements() []*int32
	SetSourceLanguage(v string) *VisionFlowRequest
	GetSourceLanguage() *string
	SetTargetHeight(v int32) *VisionFlowRequest
	GetTargetHeight() *int32
	SetTargetLanguage(v string) *VisionFlowRequest
	GetTargetLanguage() *string
	SetTargetWidth(v int32) *VisionFlowRequest
	GetTargetWidth() *int32
	SetTranslatingBrandInTheProduct(v bool) *VisionFlowRequest
	GetTranslatingBrandInTheProduct() *bool
	SetUpscaleFactor(v int32) *VisionFlowRequest
	GetUpscaleFactor() *int32
}

type VisionFlowRequest struct {
	// The AI capabilities to apply (1 = intelligent element detection, 2 = intelligent matting, 3 = intelligent removal, 4 = Image Translation Pro, 5 = intelligent cropping, 6 = HD upscaling). Multiple selections allowed.
	//
	// This parameter is required.
	//
	// example:
	//
	// [1,2,3,4]
	Ability []*int32 `json:"Ability,omitempty" xml:"Ability,omitempty" type:"Repeated"`
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
	NonobjectDetectElements []*int32 `json:"NonobjectDetectElements,omitempty" xml:"NonobjectDetectElements,omitempty" type:"Repeated"`
	// The elements to remove from the non-subject area of the image (1 = transparent text block, 2 = specific name, 3 = text, 4 = overlay patch). Multiple selections allowed. When the intelligent removal capability is selected, at least one of NonobjectRemoveElements and ObjectRemoveElements is required.
	//
	// example:
	//
	// [1,2,4]
	NonobjectRemoveElements []*int32 `json:"NonobjectRemoveElements,omitempty" xml:"NonobjectRemoveElements,omitempty" type:"Repeated"`
	// The elements to detect on the image subject (1 = watermark, 2 = logo, 3 = text, 4 = text-bearing color block). Multiple selections allowed. When the intelligent element detection capability is selected, at least one of ObjectDetectElements and NonobjectDetectElements is required.
	//
	// example:
	//
	// [1,2,3,4]
	ObjectDetectElements []*int32 `json:"ObjectDetectElements,omitempty" xml:"ObjectDetectElements,omitempty" type:"Repeated"`
	// The elements to remove from the image subject (1 = transparent text block, 2 = specific name, 3 = text, 4 = overlay patch). Multiple selections allowed. When the intelligent removal capability is selected, at least one of ObjectRemoveElements and NonobjectRemoveElements is required.
	//
	// example:
	//
	// [1,2,4]
	ObjectRemoveElements []*int32 `json:"ObjectRemoveElements,omitempty" xml:"ObjectRemoveElements,omitempty" type:"Repeated"`
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

func (s VisionFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s VisionFlowRequest) GoString() string {
	return s.String()
}

func (s *VisionFlowRequest) GetAbility() []*int32 {
	return s.Ability
}

func (s *VisionFlowRequest) GetBackGroundType() *string {
	return s.BackGroundType
}

func (s *VisionFlowRequest) GetGlossary() *string {
	return s.Glossary
}

func (s *VisionFlowRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *VisionFlowRequest) GetIncludingProductArea() *bool {
	return s.IncludingProductArea
}

func (s *VisionFlowRequest) GetIsFilter() *bool {
	return s.IsFilter
}

func (s *VisionFlowRequest) GetMask() *string {
	return s.Mask
}

func (s *VisionFlowRequest) GetNonobjectDetectElements() []*int32 {
	return s.NonobjectDetectElements
}

func (s *VisionFlowRequest) GetNonobjectRemoveElements() []*int32 {
	return s.NonobjectRemoveElements
}

func (s *VisionFlowRequest) GetObjectDetectElements() []*int32 {
	return s.ObjectDetectElements
}

func (s *VisionFlowRequest) GetObjectRemoveElements() []*int32 {
	return s.ObjectRemoveElements
}

func (s *VisionFlowRequest) GetSourceLanguage() *string {
	return s.SourceLanguage
}

func (s *VisionFlowRequest) GetTargetHeight() *int32 {
	return s.TargetHeight
}

func (s *VisionFlowRequest) GetTargetLanguage() *string {
	return s.TargetLanguage
}

func (s *VisionFlowRequest) GetTargetWidth() *int32 {
	return s.TargetWidth
}

func (s *VisionFlowRequest) GetTranslatingBrandInTheProduct() *bool {
	return s.TranslatingBrandInTheProduct
}

func (s *VisionFlowRequest) GetUpscaleFactor() *int32 {
	return s.UpscaleFactor
}

func (s *VisionFlowRequest) SetAbility(v []*int32) *VisionFlowRequest {
	s.Ability = v
	return s
}

func (s *VisionFlowRequest) SetBackGroundType(v string) *VisionFlowRequest {
	s.BackGroundType = &v
	return s
}

func (s *VisionFlowRequest) SetGlossary(v string) *VisionFlowRequest {
	s.Glossary = &v
	return s
}

func (s *VisionFlowRequest) SetImageUrl(v string) *VisionFlowRequest {
	s.ImageUrl = &v
	return s
}

func (s *VisionFlowRequest) SetIncludingProductArea(v bool) *VisionFlowRequest {
	s.IncludingProductArea = &v
	return s
}

func (s *VisionFlowRequest) SetIsFilter(v bool) *VisionFlowRequest {
	s.IsFilter = &v
	return s
}

func (s *VisionFlowRequest) SetMask(v string) *VisionFlowRequest {
	s.Mask = &v
	return s
}

func (s *VisionFlowRequest) SetNonobjectDetectElements(v []*int32) *VisionFlowRequest {
	s.NonobjectDetectElements = v
	return s
}

func (s *VisionFlowRequest) SetNonobjectRemoveElements(v []*int32) *VisionFlowRequest {
	s.NonobjectRemoveElements = v
	return s
}

func (s *VisionFlowRequest) SetObjectDetectElements(v []*int32) *VisionFlowRequest {
	s.ObjectDetectElements = v
	return s
}

func (s *VisionFlowRequest) SetObjectRemoveElements(v []*int32) *VisionFlowRequest {
	s.ObjectRemoveElements = v
	return s
}

func (s *VisionFlowRequest) SetSourceLanguage(v string) *VisionFlowRequest {
	s.SourceLanguage = &v
	return s
}

func (s *VisionFlowRequest) SetTargetHeight(v int32) *VisionFlowRequest {
	s.TargetHeight = &v
	return s
}

func (s *VisionFlowRequest) SetTargetLanguage(v string) *VisionFlowRequest {
	s.TargetLanguage = &v
	return s
}

func (s *VisionFlowRequest) SetTargetWidth(v int32) *VisionFlowRequest {
	s.TargetWidth = &v
	return s
}

func (s *VisionFlowRequest) SetTranslatingBrandInTheProduct(v bool) *VisionFlowRequest {
	s.TranslatingBrandInTheProduct = &v
	return s
}

func (s *VisionFlowRequest) SetUpscaleFactor(v int32) *VisionFlowRequest {
	s.UpscaleFactor = &v
	return s
}

func (s *VisionFlowRequest) Validate() error {
	return dara.Validate(s)
}
