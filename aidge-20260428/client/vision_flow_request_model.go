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
	// The AI capabilities to be processed. Valid values: 1 for intelligent element recognition, 2 for intelligent cutout, 3 for intelligent erasure, 4 for image translation Pro, 5 for intelligent cropping, and 6 for high-definition upscaling. You can specify multiple values.
	//
	// This parameter is required.
	//
	// example:
	//
	// [1,2,3,4]
	Ability []*int32 `json:"Ability,omitempty" xml:"Ability,omitempty" type:"Repeated"`
	// The background type of the returned image. Valid values: WHITE_BACKGROUND (white background) and TRANSPARENT (transparent background). Required if you select the intelligent cutout capability.
	//
	// example:
	//
	// WHITE_BACKGROUND
	BackGroundType *string `json:"BackGroundType,omitempty" xml:"BackGroundType,omitempty"`
	// The ID of the intervention glossary. Optional. You must create the glossary in the console and provide its ID. If this parameter is empty, the translation results are not modified.
	//
	// example:
	//
	// glossary_1
	Glossary *string `json:"Glossary,omitempty" xml:"Glossary,omitempty"`
	// The URL of the image to be processed. Required. The resolution must be greater than 256×256, the long edge must not exceed 1920 pixels, and the short edge must not exceed 1080 pixels. The maximum size is 5 MB. Supported formats include PNG, JPEG, JPG, BMP, and WEBP.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://ae01.alicdn.com/kf/S342f0070dc9f4be09a6cbed34e90dc8fs.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// Specifies whether to translate text on the image subject. Optional. Default value: false. This parameter helps protect embedded information such as product names from being translated.
	//
	// example:
	//
	// false
	IncludingProductArea *bool `json:"IncludingProductArea,omitempty" xml:"IncludingProductArea,omitempty"`
	// Specifies whether to include images with the specified elements in subsequent processing. A value of true indicates that images with the elements are included. A value of false indicates that images with the elements are excluded. Required if you select the intelligent element recognition capability.
	//
	// example:
	//
	// true
	IsFilter *bool `json:"IsFilter,omitempty" xml:"IsFilter,omitempty"`
	// The specific area to remove in Run-Length Encoding (RLE) format. Optional. If specified, this parameter takes precedence, and the ObjectRemoveElements and NonobjectRemoveElements parameters are ignored.
	//
	// example:
	//
	// 474556 160 475356 160
	Mask *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	// The elements to detect on the non-subject area of the image. Valid values: 1 for watermark, 2 for logo, 3 for text, and 4 for text block. You can specify multiple values. If you select the intelligent element recognition capability, you must specify at least one of this parameter and ObjectDetectElements.
	//
	// example:
	//
	// [1,2,3]
	NonobjectDetectElements []*int32 `json:"NonobjectDetectElements,omitempty" xml:"NonobjectDetectElements,omitempty" type:"Repeated"`
	// The elements to remove from the non-subject area of the image. Valid values: 1 for transparent text block, 2 for specific name, 3 for text, and 4 for cluttered text overlay. You can specify multiple values. If you select the intelligent erasure capability, you must specify at least one of this parameter and ObjectRemoveElements.
	//
	// example:
	//
	// [1,2,4]
	NonobjectRemoveElements []*int32 `json:"NonobjectRemoveElements,omitempty" xml:"NonobjectRemoveElements,omitempty" type:"Repeated"`
	// The elements to detect on the image subject. Valid values: 1 for watermark, 2 for logo, 3 for text, and 4 for text block. You can specify multiple values. If you select the intelligent element recognition capability, you must specify at least one of this parameter and NonobjectDetectElements.
	//
	// example:
	//
	// [1,2,3,4]
	ObjectDetectElements []*int32 `json:"ObjectDetectElements,omitempty" xml:"ObjectDetectElements,omitempty" type:"Repeated"`
	// The elements to remove from the image subject. Valid values: 1 for transparent text block, 2 for specific name, 3 for text, and 4 for cluttered text overlay. You can specify multiple values. If you select the intelligent erasure capability, you must specify at least one of this parameter and NonobjectRemoveElements.
	//
	// example:
	//
	// [1,2,4]
	ObjectRemoveElements []*int32 `json:"ObjectRemoveElements,omitempty" xml:"ObjectRemoveElements,omitempty" type:"Repeated"`
	// The source language code. Optional. For supported language pairs, see the list of supported translation languages.
	//
	// example:
	//
	// zh
	SourceLanguage *string `json:"SourceLanguage,omitempty" xml:"SourceLanguage,omitempty"`
	// The expected height of the cropped image in pixels. Valid values: 100 to 5000. Required if you select the intelligent cropping capability.
	//
	// example:
	//
	// 800
	TargetHeight *int32 `json:"TargetHeight,omitempty" xml:"TargetHeight,omitempty"`
	// The target language code. Optional. For supported language pairs, see the list of supported translation languages.
	//
	// example:
	//
	// en
	TargetLanguage *string `json:"TargetLanguage,omitempty" xml:"TargetLanguage,omitempty"`
	// The expected width of the cropped image in pixels. Valid values: 100 to 5000. Required if you select the intelligent cropping capability.
	//
	// example:
	//
	// 800
	TargetWidth *int32 `json:"TargetWidth,omitempty" xml:"TargetWidth,omitempty"`
	// Specifies whether to translate brand names in the image. Optional. Default value: false. This parameter helps protect brand name information from being translated.
	//
	// example:
	//
	// false
	TranslatingBrandInTheProduct *bool `json:"TranslatingBrandInTheProduct,omitempty" xml:"TranslatingBrandInTheProduct,omitempty"`
	// The image upscaling factor. Optional. Default value: 2. Valid values: 2 to 4. Required if you select the high-definition upscaling capability.
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
