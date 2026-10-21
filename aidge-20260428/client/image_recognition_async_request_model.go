// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageRecognitionAsyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *ImageRecognitionAsyncRequest
	GetBizId() *string
	SetImageUrl(v string) *ImageRecognitionAsyncRequest
	GetImageUrl() *string
	SetNonObjectDetectElements(v []*int32) *ImageRecognitionAsyncRequest
	GetNonObjectDetectElements() []*int32
	SetObjectDetectElements(v []*int32) *ImageRecognitionAsyncRequest
	GetObjectDetectElements() []*int32
	SetReturnBorderPixel(v int32) *ImageRecognitionAsyncRequest
	GetReturnBorderPixel() *int32
	SetReturnCharacter(v int32) *ImageRecognitionAsyncRequest
	GetReturnCharacter() *int32
	SetReturnCharacterProp(v int32) *ImageRecognitionAsyncRequest
	GetReturnCharacterProp() *int32
	SetReturnProductNum(v int32) *ImageRecognitionAsyncRequest
	GetReturnProductNum() *int32
	SetReturnProductProp(v int32) *ImageRecognitionAsyncRequest
	GetReturnProductProp() *int32
}

type ImageRecognitionAsyncRequest struct {
	// The business ID of the caller. It is recorded upon acceptance and returned as is when you query the result of the asynchronous task.
	//
	// example:
	//
	// order-20260928-001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The image URL. The image must be in JPG, JPEG, PNG, BMP, or WEBP format. The resolution must be between 256 × 256 and 3000 × 3000 pixels, and the file size cannot exceed 10 MB.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://ae01.alicdn.com/kf/S342f0070dc9f4be09a6cbed34e90dc8fs.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The types of elements to detect outside the image subject. Valid values:
	//
	// - 1: watermark
	//
	// - 2: logo
	//
	// - 3: text
	//
	// - 4: text-containing color block
	//
	// You can specify multiple values. This parameter and ObjectDetectElements cannot be empty at the same time.
	NonObjectDetectElements []*int32 `json:"NonObjectDetectElements,omitempty" xml:"NonObjectDetectElements,omitempty" type:"Repeated"`
	// The types of elements to detect on the image subject. Valid values:
	//
	// - 1: watermark
	//
	// - 2: logo
	//
	// - 3: text
	//
	// - 4: text-containing color block
	//
	// You can specify multiple values.
	ObjectDetectElements []*int32 `json:"ObjectDetectElements,omitempty" xml:"ObjectDetectElements,omitempty" type:"Repeated"`
	// Specifies whether to return the marginalization distance of the subject. Valid values:
	//
	// - 1: yes
	//
	// - 0: no
	//
	// Default value: 0.
	//
	// example:
	//
	// 1
	ReturnBorderPixel *int32 `json:"ReturnBorderPixel,omitempty" xml:"ReturnBorderPixel,omitempty"`
	// Specifies whether to return the OCR result of the text in the recognized area. Valid values:
	//
	// - 1: yes
	//
	// - 0: no
	//
	// Default value: 0.
	//
	// example:
	//
	// 1
	ReturnCharacter *int32 `json:"ReturnCharacter,omitempty" xml:"ReturnCharacter,omitempty"`
	// Specifies whether to return the text proportion. Valid values:
	//
	// - 1: yes
	//
	// - 0: no
	//
	// Default value: 0.
	//
	// This parameter takes effect only when the following two conditions are met:
	//
	// 1. ReturnCharacter is set to 1.
	//
	// 2. ObjectDetectElements or NonObjectDetectElements is set to 3, which means text needs to be recognized.
	//
	// example:
	//
	// 1
	ReturnCharacterProp *int32 `json:"ReturnCharacterProp,omitempty" xml:"ReturnCharacterProp,omitempty"`
	// Specifies whether to return the number of subjects. Valid values:
	//
	// - 1: yes
	//
	// - 0: no
	//
	// Default value: 0.
	//
	// example:
	//
	// 1
	ReturnProductNum *int32 `json:"ReturnProductNum,omitempty" xml:"ReturnProductNum,omitempty"`
	// Specifies whether to return the area proportion of the subject. Valid values:
	//
	// - 1: yes
	//
	// - 0: no
	//
	// Default value: 0.
	//
	// example:
	//
	// 1
	ReturnProductProp *int32 `json:"ReturnProductProp,omitempty" xml:"ReturnProductProp,omitempty"`
}

func (s ImageRecognitionAsyncRequest) String() string {
	return dara.Prettify(s)
}

func (s ImageRecognitionAsyncRequest) GoString() string {
	return s.String()
}

func (s *ImageRecognitionAsyncRequest) GetBizId() *string {
	return s.BizId
}

func (s *ImageRecognitionAsyncRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ImageRecognitionAsyncRequest) GetNonObjectDetectElements() []*int32 {
	return s.NonObjectDetectElements
}

func (s *ImageRecognitionAsyncRequest) GetObjectDetectElements() []*int32 {
	return s.ObjectDetectElements
}

func (s *ImageRecognitionAsyncRequest) GetReturnBorderPixel() *int32 {
	return s.ReturnBorderPixel
}

func (s *ImageRecognitionAsyncRequest) GetReturnCharacter() *int32 {
	return s.ReturnCharacter
}

func (s *ImageRecognitionAsyncRequest) GetReturnCharacterProp() *int32 {
	return s.ReturnCharacterProp
}

func (s *ImageRecognitionAsyncRequest) GetReturnProductNum() *int32 {
	return s.ReturnProductNum
}

func (s *ImageRecognitionAsyncRequest) GetReturnProductProp() *int32 {
	return s.ReturnProductProp
}

func (s *ImageRecognitionAsyncRequest) SetBizId(v string) *ImageRecognitionAsyncRequest {
	s.BizId = &v
	return s
}

func (s *ImageRecognitionAsyncRequest) SetImageUrl(v string) *ImageRecognitionAsyncRequest {
	s.ImageUrl = &v
	return s
}

func (s *ImageRecognitionAsyncRequest) SetNonObjectDetectElements(v []*int32) *ImageRecognitionAsyncRequest {
	s.NonObjectDetectElements = v
	return s
}

func (s *ImageRecognitionAsyncRequest) SetObjectDetectElements(v []*int32) *ImageRecognitionAsyncRequest {
	s.ObjectDetectElements = v
	return s
}

func (s *ImageRecognitionAsyncRequest) SetReturnBorderPixel(v int32) *ImageRecognitionAsyncRequest {
	s.ReturnBorderPixel = &v
	return s
}

func (s *ImageRecognitionAsyncRequest) SetReturnCharacter(v int32) *ImageRecognitionAsyncRequest {
	s.ReturnCharacter = &v
	return s
}

func (s *ImageRecognitionAsyncRequest) SetReturnCharacterProp(v int32) *ImageRecognitionAsyncRequest {
	s.ReturnCharacterProp = &v
	return s
}

func (s *ImageRecognitionAsyncRequest) SetReturnProductNum(v int32) *ImageRecognitionAsyncRequest {
	s.ReturnProductNum = &v
	return s
}

func (s *ImageRecognitionAsyncRequest) SetReturnProductProp(v int32) *ImageRecognitionAsyncRequest {
	s.ReturnProductProp = &v
	return s
}

func (s *ImageRecognitionAsyncRequest) Validate() error {
	return dara.Validate(s)
}
