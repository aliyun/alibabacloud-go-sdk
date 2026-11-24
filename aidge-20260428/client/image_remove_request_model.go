// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageRemoveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageUrl(v string) *ImageRemoveRequest
	GetImageUrl() *string
	SetMask(v string) *ImageRemoveRequest
	GetMask() *string
	SetNonObjectRemoveElements(v []*int32) *ImageRemoveRequest
	GetNonObjectRemoveElements() []*int32
	SetObjectRemoveElements(v []*int32) *ImageRemoveRequest
	GetObjectRemoveElements() []*int32
	SetPosition(v string) *ImageRemoveRequest
	GetPosition() *string
	SetUserImage(v []*string) *ImageRemoveRequest
	GetUserImage() []*string
	SetUserText(v []*string) *ImageRemoveRequest
	GetUserText() []*string
}

type ImageRemoveRequest struct {
	// The URL of the original image. The image must be in JPG, JPEG, PNG, BMP, or WEBP format. The resolution must be between 256 × 256 and 3000 × 3000 pixels, and the file size cannot exceed 10 MB.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://img.alicdn.com/bao/uploaded/i2/xxx.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The specific removal area. This parameter must be used in conjunction with the editor. The input format is RLE.
	//
	// If this parameter is specified, it takes precedence and the remove_non_product_area_elements and remove_product_area_elements parameters are ignored. This parameter is optional, but you must specify at least one of the following parameters: ObjectRemoveElements, NonObjectRemoveElements, Mask, Position, UserText, and UserImage.
	//
	// If multiple parameters are specified, the priority order is: UserImage > UserText > Position > Mask > ObjectRemoveElements = NonObjectRemoveElements.
	//
	// example:
	//
	// "474556 160 475356 160 476156 160 476956 160 477756 160 478556 160 479356 160 480156 160 480956 160 481756 160 482556 160 483356 160 484156 160 484956 160 485756 160 486556 160 487356 160 488156 160 488956 160 489756 160 490556 160 491356 160 492156 160"
	Mask *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	// The elements to remove from the image background. Valid values: 1 (transparent text blocks), 2 (specific names), 3 (text), and 4 (promotional stickers). You can select multiple values. This parameter is optional, but you must specify at least one of the following parameters: ObjectRemoveElements, NonObjectRemoveElements, Mask, Position, UserText, and UserImage.
	//
	// If multiple parameters are specified, the priority order is: UserImage > UserText > Position > Mask > ObjectRemoveElements = NonObjectRemoveElements.
	//
	// For more information about each type, see the product documentation.
	//
	// example:
	//
	// [1,2]
	NonObjectRemoveElements []*int32 `json:"NonObjectRemoveElements,omitempty" xml:"NonObjectRemoveElements,omitempty" type:"Repeated"`
	// The elements to remove from the image subject. Valid values: 1 (transparent text blocks), 2 (specific names), 3 (text), and 4 (promotional stickers). You can select multiple values. This parameter is optional, but you must specify at least one of the following parameters: ObjectRemoveElements, NonObjectRemoveElements, Mask, Position, UserText, and UserImage.
	//
	// If multiple parameters are specified, the priority order is: UserImage > UserText > Position > Mask > ObjectRemoveElements = NonObjectRemoveElements.
	//
	// For more information about each type, see the product documentation.
	//
	// Image subject: the core product area in the image.
	//
	// example:
	//
	// [1,2]
	ObjectRemoveElements []*int32 `json:"ObjectRemoveElements,omitempty" xml:"ObjectRemoveElements,omitempty" type:"Repeated"`
	// The specific removal area. This parameter must be used in conjunction with the editor. The input format is four-point coordinates [xx,yy,zz,dd]. This parameter is optional, but you must specify at least one of the following parameters: ObjectRemoveElements, NonObjectRemoveElements, Mask, Position, UserText, and UserImage.
	//
	// If multiple parameters are specified, the priority order is: UserImage > UserText > Position > Mask > ObjectRemoveElements = NonObjectRemoveElements.
	//
	// example:
	//
	// [10,10,100,100]
	Position *string `json:"Position,omitempty" xml:"Position,omitempty"`
	// The URLs of user-specified image elements to remove. Multiple image URLs are supported. The input format is ["https://ae01.alicdn.com/kf/S342f0070dc9f4be09a6cbed34e90dc8fs.jpg","https://ae01.alicdn.com/kf/S342f0070dc9f4be09a6cbed34e90dc8fs.jpg"]. This parameter is optional, but you must specify at least one of the following parameters: ObjectRemoveElements, NonObjectRemoveElements, Mask, Position, UserText, and UserImage.
	//
	// If multiple parameters are specified, the priority order is: UserImage > UserText > Position > Mask > ObjectRemoveElements = NonObjectRemoveElements.
	//
	// example:
	//
	// ["https://img.alicdn.com/bao/uploaded/i2/xxx.jpg"]
	UserImage []*string `json:"UserImage,omitempty" xml:"UserImage,omitempty" type:"Repeated"`
	// The user-specified text to remove. Multiple text strings are supported. The input format is ["xx","yy"]. This parameter is optional, but you must specify at least one of the following parameters: ObjectRemoveElements, NonObjectRemoveElements, Mask, Position, UserText, and UserImage.
	//
	// If multiple parameters are specified, the priority order is: UserImage > UserText > Position > Mask > ObjectRemoveElements = NonObjectRemoveElements.
	//
	// example:
	//
	// ["xx","yy"]
	UserText []*string `json:"UserText,omitempty" xml:"UserText,omitempty" type:"Repeated"`
}

func (s ImageRemoveRequest) String() string {
	return dara.Prettify(s)
}

func (s ImageRemoveRequest) GoString() string {
	return s.String()
}

func (s *ImageRemoveRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ImageRemoveRequest) GetMask() *string {
	return s.Mask
}

func (s *ImageRemoveRequest) GetNonObjectRemoveElements() []*int32 {
	return s.NonObjectRemoveElements
}

func (s *ImageRemoveRequest) GetObjectRemoveElements() []*int32 {
	return s.ObjectRemoveElements
}

func (s *ImageRemoveRequest) GetPosition() *string {
	return s.Position
}

func (s *ImageRemoveRequest) GetUserImage() []*string {
	return s.UserImage
}

func (s *ImageRemoveRequest) GetUserText() []*string {
	return s.UserText
}

func (s *ImageRemoveRequest) SetImageUrl(v string) *ImageRemoveRequest {
	s.ImageUrl = &v
	return s
}

func (s *ImageRemoveRequest) SetMask(v string) *ImageRemoveRequest {
	s.Mask = &v
	return s
}

func (s *ImageRemoveRequest) SetNonObjectRemoveElements(v []*int32) *ImageRemoveRequest {
	s.NonObjectRemoveElements = v
	return s
}

func (s *ImageRemoveRequest) SetObjectRemoveElements(v []*int32) *ImageRemoveRequest {
	s.ObjectRemoveElements = v
	return s
}

func (s *ImageRemoveRequest) SetPosition(v string) *ImageRemoveRequest {
	s.Position = &v
	return s
}

func (s *ImageRemoveRequest) SetUserImage(v []*string) *ImageRemoveRequest {
	s.UserImage = v
	return s
}

func (s *ImageRemoveRequest) SetUserText(v []*string) *ImageRemoveRequest {
	s.UserText = v
	return s
}

func (s *ImageRemoveRequest) Validate() error {
	return dara.Validate(s)
}
