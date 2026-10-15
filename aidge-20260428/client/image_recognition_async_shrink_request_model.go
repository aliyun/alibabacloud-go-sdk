// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageRecognitionAsyncShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageUrl(v string) *ImageRecognitionAsyncShrinkRequest
	GetImageUrl() *string
	SetNonObjectDetectElementsShrink(v string) *ImageRecognitionAsyncShrinkRequest
	GetNonObjectDetectElementsShrink() *string
	SetObjectDetectElementsShrink(v string) *ImageRecognitionAsyncShrinkRequest
	GetObjectDetectElementsShrink() *string
	SetReturnBorderPixel(v int32) *ImageRecognitionAsyncShrinkRequest
	GetReturnBorderPixel() *int32
	SetReturnCharacter(v int32) *ImageRecognitionAsyncShrinkRequest
	GetReturnCharacter() *int32
	SetReturnCharacterProp(v int32) *ImageRecognitionAsyncShrinkRequest
	GetReturnCharacterProp() *int32
	SetReturnProductNum(v int32) *ImageRecognitionAsyncShrinkRequest
	GetReturnProductNum() *int32
	SetReturnProductProp(v int32) *ImageRecognitionAsyncShrinkRequest
	GetReturnProductProp() *int32
}

type ImageRecognitionAsyncShrinkRequest struct {
	// The URL of the image. The image must be in JPG, JPEG, PNG, BMP, or WEBP format. The resolution must be between 256 × 256 and 3000 × 3000 pixels, and the file size cannot exceed 10 MB.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://ae01.alicdn.com/kf/S342f0070dc9f4be09a6cbed34e90dc8fs.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// Specifies the types of elements to be detected outside the image subject. Valid values:
	//
	// - 1: Watermark.
	//
	// - 2: Logo.
	//
	// - 3: Text.
	//
	// - 4: Text-containing color block.
	//
	// You can specify multiple element types. This parameter and ObjectDetectElements cannot be empty at the same time.
	NonObjectDetectElementsShrink *string `json:"NonObjectDetectElements,omitempty" xml:"NonObjectDetectElements,omitempty"`
	// Specifies the types of elements to be detected on the image subject. Valid values:
	//
	// - 1: Watermark.
	//
	// - 2: Logo.
	//
	// - 3: Text.
	//
	// - 4: Text-containing color block.
	//
	// You can specify multiple element types.
	ObjectDetectElementsShrink *string `json:"ObjectDetectElements,omitempty" xml:"ObjectDetectElements,omitempty"`
	// Specifies whether to return the marginalization distance of the subject. Valid values:
	//
	// - 1: Yes.
	//
	// - 0: No.
	//
	// Default value: 0.
	//
	// example:
	//
	// 1
	ReturnBorderPixel *int32 `json:"ReturnBorderPixel,omitempty" xml:"ReturnBorderPixel,omitempty"`
	// Specifies whether to return the OCR results of the text in the recognized area. Valid values:
	//
	// - 1: Yes.
	//
	// - 0: No.
	//
	// Default value: 0.
	//
	// example:
	//
	// 1
	ReturnCharacter *int32 `json:"ReturnCharacter,omitempty" xml:"ReturnCharacter,omitempty"`
	// Specifies whether to return the text proportion. Valid values:
	//
	// - 1: Yes.
	//
	// - 0: No.
	//
	// Default value: 0.
	//
	// This parameter takes effect only when the following two conditions are met:
	//
	// 1. ReturnCharacter is set to 1.
	//
	// 2. ObjectDetectElements or NonObjectDetectElements is set to 3, which means text recognition is required.
	//
	// example:
	//
	// 1
	ReturnCharacterProp *int32 `json:"ReturnCharacterProp,omitempty" xml:"ReturnCharacterProp,omitempty"`
	// Specifies whether to return the number of subjects. Valid values:
	//
	// - 1: Yes.
	//
	// - 0: No.
	//
	// Default value: 0.
	//
	// example:
	//
	// 1
	ReturnProductNum *int32 `json:"ReturnProductNum,omitempty" xml:"ReturnProductNum,omitempty"`
	// Specifies whether to return the area proportion of the subject. Valid values:
	//
	// - 1: Yes.
	//
	// - 0: No.
	//
	// Default value: 0.
	//
	// example:
	//
	// 1
	ReturnProductProp *int32 `json:"ReturnProductProp,omitempty" xml:"ReturnProductProp,omitempty"`
}

func (s ImageRecognitionAsyncShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ImageRecognitionAsyncShrinkRequest) GoString() string {
	return s.String()
}

func (s *ImageRecognitionAsyncShrinkRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ImageRecognitionAsyncShrinkRequest) GetNonObjectDetectElementsShrink() *string {
	return s.NonObjectDetectElementsShrink
}

func (s *ImageRecognitionAsyncShrinkRequest) GetObjectDetectElementsShrink() *string {
	return s.ObjectDetectElementsShrink
}

func (s *ImageRecognitionAsyncShrinkRequest) GetReturnBorderPixel() *int32 {
	return s.ReturnBorderPixel
}

func (s *ImageRecognitionAsyncShrinkRequest) GetReturnCharacter() *int32 {
	return s.ReturnCharacter
}

func (s *ImageRecognitionAsyncShrinkRequest) GetReturnCharacterProp() *int32 {
	return s.ReturnCharacterProp
}

func (s *ImageRecognitionAsyncShrinkRequest) GetReturnProductNum() *int32 {
	return s.ReturnProductNum
}

func (s *ImageRecognitionAsyncShrinkRequest) GetReturnProductProp() *int32 {
	return s.ReturnProductProp
}

func (s *ImageRecognitionAsyncShrinkRequest) SetImageUrl(v string) *ImageRecognitionAsyncShrinkRequest {
	s.ImageUrl = &v
	return s
}

func (s *ImageRecognitionAsyncShrinkRequest) SetNonObjectDetectElementsShrink(v string) *ImageRecognitionAsyncShrinkRequest {
	s.NonObjectDetectElementsShrink = &v
	return s
}

func (s *ImageRecognitionAsyncShrinkRequest) SetObjectDetectElementsShrink(v string) *ImageRecognitionAsyncShrinkRequest {
	s.ObjectDetectElementsShrink = &v
	return s
}

func (s *ImageRecognitionAsyncShrinkRequest) SetReturnBorderPixel(v int32) *ImageRecognitionAsyncShrinkRequest {
	s.ReturnBorderPixel = &v
	return s
}

func (s *ImageRecognitionAsyncShrinkRequest) SetReturnCharacter(v int32) *ImageRecognitionAsyncShrinkRequest {
	s.ReturnCharacter = &v
	return s
}

func (s *ImageRecognitionAsyncShrinkRequest) SetReturnCharacterProp(v int32) *ImageRecognitionAsyncShrinkRequest {
	s.ReturnCharacterProp = &v
	return s
}

func (s *ImageRecognitionAsyncShrinkRequest) SetReturnProductNum(v int32) *ImageRecognitionAsyncShrinkRequest {
	s.ReturnProductNum = &v
	return s
}

func (s *ImageRecognitionAsyncShrinkRequest) SetReturnProductProp(v int32) *ImageRecognitionAsyncShrinkRequest {
	s.ReturnProductProp = &v
	return s
}

func (s *ImageRecognitionAsyncShrinkRequest) Validate() error {
	return dara.Validate(s)
}
