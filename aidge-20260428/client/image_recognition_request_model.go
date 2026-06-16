// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageRecognitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageUrl(v string) *ImageRecognitionRequest
	GetImageUrl() *string
	SetNonObjectDetectElements(v []*int32) *ImageRecognitionRequest
	GetNonObjectDetectElements() []*int32
	SetObjectDetectElements(v []*int32) *ImageRecognitionRequest
	GetObjectDetectElements() []*int32
	SetReturnBorderPixel(v int32) *ImageRecognitionRequest
	GetReturnBorderPixel() *int32
	SetReturnCharacter(v int32) *ImageRecognitionRequest
	GetReturnCharacter() *int32
	SetReturnCharacterProp(v int32) *ImageRecognitionRequest
	GetReturnCharacterProp() *int32
	SetReturnProductNum(v int32) *ImageRecognitionRequest
	GetReturnProductNum() *int32
	SetReturnProductProp(v int32) *ImageRecognitionRequest
	GetReturnProductProp() *int32
}

type ImageRecognitionRequest struct {
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
	NonObjectDetectElements []*int32 `json:"NonObjectDetectElements,omitempty" xml:"NonObjectDetectElements,omitempty" type:"Repeated"`
	// The list of subject element types to detect. Valid values: 1 (product subject), 2 (model), 3 (text), and 4 (logo).
	//
	// example:
	//
	// [1,2,3,4]
	ObjectDetectElements []*int32 `json:"ObjectDetectElements,omitempty" xml:"ObjectDetectElements,omitempty" type:"Repeated"`
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

func (s ImageRecognitionRequest) String() string {
	return dara.Prettify(s)
}

func (s ImageRecognitionRequest) GoString() string {
	return s.String()
}

func (s *ImageRecognitionRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ImageRecognitionRequest) GetNonObjectDetectElements() []*int32 {
	return s.NonObjectDetectElements
}

func (s *ImageRecognitionRequest) GetObjectDetectElements() []*int32 {
	return s.ObjectDetectElements
}

func (s *ImageRecognitionRequest) GetReturnBorderPixel() *int32 {
	return s.ReturnBorderPixel
}

func (s *ImageRecognitionRequest) GetReturnCharacter() *int32 {
	return s.ReturnCharacter
}

func (s *ImageRecognitionRequest) GetReturnCharacterProp() *int32 {
	return s.ReturnCharacterProp
}

func (s *ImageRecognitionRequest) GetReturnProductNum() *int32 {
	return s.ReturnProductNum
}

func (s *ImageRecognitionRequest) GetReturnProductProp() *int32 {
	return s.ReturnProductProp
}

func (s *ImageRecognitionRequest) SetImageUrl(v string) *ImageRecognitionRequest {
	s.ImageUrl = &v
	return s
}

func (s *ImageRecognitionRequest) SetNonObjectDetectElements(v []*int32) *ImageRecognitionRequest {
	s.NonObjectDetectElements = v
	return s
}

func (s *ImageRecognitionRequest) SetObjectDetectElements(v []*int32) *ImageRecognitionRequest {
	s.ObjectDetectElements = v
	return s
}

func (s *ImageRecognitionRequest) SetReturnBorderPixel(v int32) *ImageRecognitionRequest {
	s.ReturnBorderPixel = &v
	return s
}

func (s *ImageRecognitionRequest) SetReturnCharacter(v int32) *ImageRecognitionRequest {
	s.ReturnCharacter = &v
	return s
}

func (s *ImageRecognitionRequest) SetReturnCharacterProp(v int32) *ImageRecognitionRequest {
	s.ReturnCharacterProp = &v
	return s
}

func (s *ImageRecognitionRequest) SetReturnProductNum(v int32) *ImageRecognitionRequest {
	s.ReturnProductNum = &v
	return s
}

func (s *ImageRecognitionRequest) SetReturnProductProp(v int32) *ImageRecognitionRequest {
	s.ReturnProductProp = &v
	return s
}

func (s *ImageRecognitionRequest) Validate() error {
	return dara.Validate(s)
}
