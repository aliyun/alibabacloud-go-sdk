// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageRemoveAsyncShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageUrl(v string) *ImageRemoveAsyncShrinkRequest
	GetImageUrl() *string
	SetMask(v string) *ImageRemoveAsyncShrinkRequest
	GetMask() *string
	SetNonObjectRemoveElementsShrink(v string) *ImageRemoveAsyncShrinkRequest
	GetNonObjectRemoveElementsShrink() *string
	SetObjectRemoveElementsShrink(v string) *ImageRemoveAsyncShrinkRequest
	GetObjectRemoveElementsShrink() *string
	SetPosition(v string) *ImageRemoveAsyncShrinkRequest
	GetPosition() *string
	SetUserImageShrink(v string) *ImageRemoveAsyncShrinkRequest
	GetUserImageShrink() *string
	SetUserTextShrink(v string) *ImageRemoveAsyncShrinkRequest
	GetUserTextShrink() *string
}

type ImageRemoveAsyncShrinkRequest struct {
	// The URL of the original image. The image must be in JPG, JPEG, PNG, BMP, or WEBP format. The resolution must be between 256 × 256 and 3000 × 3000 pixels, and the file size must not exceed 10 MB.
	//
	// This parameter is required.
	//
	// example:
	//
	// "https://ae01.alicdn.com/kf/S342f0070dc9f4be09a6cbed34e90dc8fs.jpg"
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The specific area to erase. This parameter must be used in conjunction with the editor, and the input format is RLE. If this parameter is specified, it takes precedence and the remove_non_product_area_elements and remove_product_area_elements parameters are ignored. This parameter is not required, but you must specify at least one of the following parameters: ObjectRemoveElements, NonObjectRemoveElements, Mask, Position, UserText, and UserImage. If multiple parameters are specified, the priority is as follows: UserImage > UserText > Position > Mask > ObjectRemoveElements = NonObjectRemoveElements.
	//
	// example:
	//
	// "474556 160 475356 160 476156 160 476956 160 477756 160 478556 160 479356 160 480156 160 480956 160 481756 160 482556 160 483356 160 484156 160 484956 160 485756 160 486556 160 487356 160 488156 160 488956 160 489756 160 490556 160 491356 160 492156 160"
	Mask *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	// The list of element types to remove from the non-subject area of the image.
	NonObjectRemoveElementsShrink *string `json:"NonObjectRemoveElements,omitempty" xml:"NonObjectRemoveElements,omitempty"`
	// The types of elements to remove from the image subject. Valid values: 1 (transparent text blocks), 2 (specific names), 3 (text), and 4 (watermarks). You can select multiple element types. This parameter is not required, but you must specify at least one of the following parameters: ObjectRemoveElements, NonObjectRemoveElements, Mask, Position, UserText, and UserImage. If multiple parameters are specified, the priority is as follows: UserImage > UserText > Position > Mask > ObjectRemoveElements = NonObjectRemoveElements. For more information about each type, see the product documentation. Image subject: the core product area in the image.
	ObjectRemoveElementsShrink *string `json:"ObjectRemoveElements,omitempty" xml:"ObjectRemoveElements,omitempty"`
	// The specific area to erase. This parameter must be used in conjunction with the editor, and the input format is four-point coordinates [xx,yy,zz,dd]. This parameter is not required, but you must specify at least one of the following parameters: ObjectRemoveElements, NonObjectRemoveElements, Mask, Position, UserText, and UserImage. If multiple parameters are specified, the priority is as follows: UserImage > UserText > Position > Mask > ObjectRemoveElements = NonObjectRemoveElements.
	//
	// example:
	//
	// [10,20,100,200]
	Position *string `json:"Position,omitempty" xml:"Position,omitempty"`
	// The list of URLs for user-specified image elements to erase.
	UserImageShrink *string `json:"UserImage,omitempty" xml:"UserImage,omitempty"`
	// The list of user-specified texts to erase.
	UserTextShrink *string `json:"UserText,omitempty" xml:"UserText,omitempty"`
}

func (s ImageRemoveAsyncShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ImageRemoveAsyncShrinkRequest) GoString() string {
	return s.String()
}

func (s *ImageRemoveAsyncShrinkRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ImageRemoveAsyncShrinkRequest) GetMask() *string {
	return s.Mask
}

func (s *ImageRemoveAsyncShrinkRequest) GetNonObjectRemoveElementsShrink() *string {
	return s.NonObjectRemoveElementsShrink
}

func (s *ImageRemoveAsyncShrinkRequest) GetObjectRemoveElementsShrink() *string {
	return s.ObjectRemoveElementsShrink
}

func (s *ImageRemoveAsyncShrinkRequest) GetPosition() *string {
	return s.Position
}

func (s *ImageRemoveAsyncShrinkRequest) GetUserImageShrink() *string {
	return s.UserImageShrink
}

func (s *ImageRemoveAsyncShrinkRequest) GetUserTextShrink() *string {
	return s.UserTextShrink
}

func (s *ImageRemoveAsyncShrinkRequest) SetImageUrl(v string) *ImageRemoveAsyncShrinkRequest {
	s.ImageUrl = &v
	return s
}

func (s *ImageRemoveAsyncShrinkRequest) SetMask(v string) *ImageRemoveAsyncShrinkRequest {
	s.Mask = &v
	return s
}

func (s *ImageRemoveAsyncShrinkRequest) SetNonObjectRemoveElementsShrink(v string) *ImageRemoveAsyncShrinkRequest {
	s.NonObjectRemoveElementsShrink = &v
	return s
}

func (s *ImageRemoveAsyncShrinkRequest) SetObjectRemoveElementsShrink(v string) *ImageRemoveAsyncShrinkRequest {
	s.ObjectRemoveElementsShrink = &v
	return s
}

func (s *ImageRemoveAsyncShrinkRequest) SetPosition(v string) *ImageRemoveAsyncShrinkRequest {
	s.Position = &v
	return s
}

func (s *ImageRemoveAsyncShrinkRequest) SetUserImageShrink(v string) *ImageRemoveAsyncShrinkRequest {
	s.UserImageShrink = &v
	return s
}

func (s *ImageRemoveAsyncShrinkRequest) SetUserTextShrink(v string) *ImageRemoveAsyncShrinkRequest {
	s.UserTextShrink = &v
	return s
}

func (s *ImageRemoveAsyncShrinkRequest) Validate() error {
	return dara.Validate(s)
}
