// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageRecognitionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageUrl(v string) *ImageRecognitionShrinkRequest
	GetImageUrl() *string
	SetNonObjectDetectElementsShrink(v string) *ImageRecognitionShrinkRequest
	GetNonObjectDetectElementsShrink() *string
	SetObjectDetectElementsShrink(v string) *ImageRecognitionShrinkRequest
	GetObjectDetectElementsShrink() *string
	SetReturnBorderPixel(v int32) *ImageRecognitionShrinkRequest
	GetReturnBorderPixel() *int32
	SetReturnCharacter(v int32) *ImageRecognitionShrinkRequest
	GetReturnCharacter() *int32
	SetReturnCharacterProp(v int32) *ImageRecognitionShrinkRequest
	GetReturnCharacterProp() *int32
	SetReturnProductNum(v int32) *ImageRecognitionShrinkRequest
	GetReturnProductNum() *int32
	SetReturnProductProp(v int32) *ImageRecognitionShrinkRequest
	GetReturnProductProp() *int32
}

type ImageRecognitionShrinkRequest struct {
	// The URL of the image to recognize.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://fb-es.mrvcdn.com/kf/E6332bde0101849968532c4f08dac757cZ.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The list of non-subject element types to detect. Valid values: 1 (background), 2 (border), 3 (watermark), and 4 (collage).
	//
	// example:
	//
	// [1,2,3,4]
	NonObjectDetectElementsShrink *string `json:"NonObjectDetectElements,omitempty" xml:"NonObjectDetectElements,omitempty"`
	// The list of subject element types to detect. Valid values: 1 (product subject), 2 (model), 3 (text), and 4 (logo).
	//
	// example:
	//
	// [1,2,3,4]
	ObjectDetectElementsShrink *string `json:"ObjectDetectElements,omitempty" xml:"ObjectDetectElements,omitempty"`
	// Specifies whether to return border pixel information. Valid values: 1 (return) and 0 (do not return).
	//
	// example:
	//
	// 1
	ReturnBorderPixel *int32 `json:"ReturnBorderPixel,omitempty" xml:"ReturnBorderPixel,omitempty"`
	// Specifies whether to return text information. Valid values: 1 (return) and 0 (do not return).
	//
	// example:
	//
	// 1
	ReturnCharacter *int32 `json:"ReturnCharacter,omitempty" xml:"ReturnCharacter,omitempty"`
	// Specifies whether to return text property information. Valid values: 1 (return) and 0 (do not return).
	//
	// example:
	//
	// 1
	ReturnCharacterProp *int32 `json:"ReturnCharacterProp,omitempty" xml:"ReturnCharacterProp,omitempty"`
	// Specifies whether to return the product count. Valid values: 1 (return) and 0 (do not return).
	//
	// example:
	//
	// 1
	ReturnProductNum *int32 `json:"ReturnProductNum,omitempty" xml:"ReturnProductNum,omitempty"`
	// Specifies whether to return product property information. Valid values: 1 (return) and 0 (do not return).
	//
	// example:
	//
	// 1
	ReturnProductProp *int32 `json:"ReturnProductProp,omitempty" xml:"ReturnProductProp,omitempty"`
}

func (s ImageRecognitionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ImageRecognitionShrinkRequest) GoString() string {
	return s.String()
}

func (s *ImageRecognitionShrinkRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ImageRecognitionShrinkRequest) GetNonObjectDetectElementsShrink() *string {
	return s.NonObjectDetectElementsShrink
}

func (s *ImageRecognitionShrinkRequest) GetObjectDetectElementsShrink() *string {
	return s.ObjectDetectElementsShrink
}

func (s *ImageRecognitionShrinkRequest) GetReturnBorderPixel() *int32 {
	return s.ReturnBorderPixel
}

func (s *ImageRecognitionShrinkRequest) GetReturnCharacter() *int32 {
	return s.ReturnCharacter
}

func (s *ImageRecognitionShrinkRequest) GetReturnCharacterProp() *int32 {
	return s.ReturnCharacterProp
}

func (s *ImageRecognitionShrinkRequest) GetReturnProductNum() *int32 {
	return s.ReturnProductNum
}

func (s *ImageRecognitionShrinkRequest) GetReturnProductProp() *int32 {
	return s.ReturnProductProp
}

func (s *ImageRecognitionShrinkRequest) SetImageUrl(v string) *ImageRecognitionShrinkRequest {
	s.ImageUrl = &v
	return s
}

func (s *ImageRecognitionShrinkRequest) SetNonObjectDetectElementsShrink(v string) *ImageRecognitionShrinkRequest {
	s.NonObjectDetectElementsShrink = &v
	return s
}

func (s *ImageRecognitionShrinkRequest) SetObjectDetectElementsShrink(v string) *ImageRecognitionShrinkRequest {
	s.ObjectDetectElementsShrink = &v
	return s
}

func (s *ImageRecognitionShrinkRequest) SetReturnBorderPixel(v int32) *ImageRecognitionShrinkRequest {
	s.ReturnBorderPixel = &v
	return s
}

func (s *ImageRecognitionShrinkRequest) SetReturnCharacter(v int32) *ImageRecognitionShrinkRequest {
	s.ReturnCharacter = &v
	return s
}

func (s *ImageRecognitionShrinkRequest) SetReturnCharacterProp(v int32) *ImageRecognitionShrinkRequest {
	s.ReturnCharacterProp = &v
	return s
}

func (s *ImageRecognitionShrinkRequest) SetReturnProductNum(v int32) *ImageRecognitionShrinkRequest {
	s.ReturnProductNum = &v
	return s
}

func (s *ImageRecognitionShrinkRequest) SetReturnProductProp(v int32) *ImageRecognitionShrinkRequest {
	s.ReturnProductProp = &v
	return s
}

func (s *ImageRecognitionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
