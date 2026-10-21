// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageRemoveAsyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *ImageRemoveAsyncRequest
	GetBizId() *string
	SetImageUrl(v string) *ImageRemoveAsyncRequest
	GetImageUrl() *string
	SetMask(v string) *ImageRemoveAsyncRequest
	GetMask() *string
	SetNonObjectRemoveElements(v []*int32) *ImageRemoveAsyncRequest
	GetNonObjectRemoveElements() []*int32
	SetObjectRemoveElements(v []*int32) *ImageRemoveAsyncRequest
	GetObjectRemoveElements() []*int32
	SetPosition(v string) *ImageRemoveAsyncRequest
	GetPosition() *string
	SetUserImage(v []*string) *ImageRemoveAsyncRequest
	GetUserImage() []*string
	SetUserText(v []*string) *ImageRemoveAsyncRequest
	GetUserText() []*string
}

type ImageRemoveAsyncRequest struct {
	// The business order ID of the caller. This ID is recorded when the request is accepted and returned as-is when you query the result of the asynchronous task.
	//
	// example:
	//
	// order-20260928-001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The URL of the original image. The image must be in JPG, JPEG, PNG, BMP, or WEBP format. The resolution must be between 256 × 256 and 3000 × 3000 pixels, and the file size cannot exceed 10 MB.
	//
	// This parameter is required.
	//
	// example:
	//
	// "https://ae01.alicdn.com/kf/S342f0070dc9f4be09a6cbed34e90dc8fs.jpg"
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The specific area to erase. This parameter must be used in conjunction with the editor. The input format is Run-Length Encoding (RLE).
	//
	// If this parameter is specified, it takes precedence and the remove_non_product_area_elements and remove_product_area_elements parameters are ignored. This parameter is optional, but you must specify at least one of the following parameters: ObjectRemoveElements, NonObjectRemoveElements, Mask, Position, UserText, and UserImage.
	//
	// If multiple parameters are specified, the priority is as follows: UserImage > UserText > Position > Mask > ObjectRemoveElements = NonObjectRemoveElements.
	//
	// example:
	//
	// "474556 160 475356 160 476156 160 476956 160 477756 160 478556 160 479356 160 480156 160 480956 160 481756 160 482556 160 483356 160 484156 160 484956 160 485756 160 486556 160 487356 160 488156 160 488956 160 489756 160 490556 160 491356 160 492156 160"
	Mask *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	// The list of element types to remove from the image background.
	NonObjectRemoveElements []*int32 `json:"NonObjectRemoveElements,omitempty" xml:"NonObjectRemoveElements,omitempty" type:"Repeated"`
	// The types of elements to remove from the image subject. Valid values:
	//
	// - 1: Transparent text blocks.
	//
	// - 2: Specific names.
	//
	// - 3: Text.
	//
	// - 4: Promotional stickers.
	//
	// You can select multiple element types. This parameter is optional, but you must specify at least one of the following parameters: ObjectRemoveElements, NonObjectRemoveElements, Mask, Position, UserText, and UserImage.
	//
	// If multiple parameters are specified, the priority is as follows: UserImage > UserText > Position > Mask > ObjectRemoveElements = NonObjectRemoveElements.
	//
	// For more information about each type, see the product documentation.
	//
	// Image subject: the core product area in the image.
	ObjectRemoveElements []*int32 `json:"ObjectRemoveElements,omitempty" xml:"ObjectRemoveElements,omitempty" type:"Repeated"`
	// The specific area to erase. This parameter must be used in conjunction with the editor. The input format is a four-point coordinate array, such as [xx,yy,zz,dd]. This parameter is optional, but you must specify at least one of the following parameters: ObjectRemoveElements, NonObjectRemoveElements, Mask, Position, UserText, and UserImage.
	//
	// If multiple parameters are specified, the priority is as follows: UserImage > UserText > Position > Mask > ObjectRemoveElements = NonObjectRemoveElements.
	//
	// example:
	//
	// [10,20,100,200]
	Position *string `json:"Position,omitempty" xml:"Position,omitempty"`
	// The list of URLs for user-specified image elements to erase.
	UserImage []*string `json:"UserImage,omitempty" xml:"UserImage,omitempty" type:"Repeated"`
	// The list of user-specified text strings to erase.
	UserText []*string `json:"UserText,omitempty" xml:"UserText,omitempty" type:"Repeated"`
}

func (s ImageRemoveAsyncRequest) String() string {
	return dara.Prettify(s)
}

func (s ImageRemoveAsyncRequest) GoString() string {
	return s.String()
}

func (s *ImageRemoveAsyncRequest) GetBizId() *string {
	return s.BizId
}

func (s *ImageRemoveAsyncRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ImageRemoveAsyncRequest) GetMask() *string {
	return s.Mask
}

func (s *ImageRemoveAsyncRequest) GetNonObjectRemoveElements() []*int32 {
	return s.NonObjectRemoveElements
}

func (s *ImageRemoveAsyncRequest) GetObjectRemoveElements() []*int32 {
	return s.ObjectRemoveElements
}

func (s *ImageRemoveAsyncRequest) GetPosition() *string {
	return s.Position
}

func (s *ImageRemoveAsyncRequest) GetUserImage() []*string {
	return s.UserImage
}

func (s *ImageRemoveAsyncRequest) GetUserText() []*string {
	return s.UserText
}

func (s *ImageRemoveAsyncRequest) SetBizId(v string) *ImageRemoveAsyncRequest {
	s.BizId = &v
	return s
}

func (s *ImageRemoveAsyncRequest) SetImageUrl(v string) *ImageRemoveAsyncRequest {
	s.ImageUrl = &v
	return s
}

func (s *ImageRemoveAsyncRequest) SetMask(v string) *ImageRemoveAsyncRequest {
	s.Mask = &v
	return s
}

func (s *ImageRemoveAsyncRequest) SetNonObjectRemoveElements(v []*int32) *ImageRemoveAsyncRequest {
	s.NonObjectRemoveElements = v
	return s
}

func (s *ImageRemoveAsyncRequest) SetObjectRemoveElements(v []*int32) *ImageRemoveAsyncRequest {
	s.ObjectRemoveElements = v
	return s
}

func (s *ImageRemoveAsyncRequest) SetPosition(v string) *ImageRemoveAsyncRequest {
	s.Position = &v
	return s
}

func (s *ImageRemoveAsyncRequest) SetUserImage(v []*string) *ImageRemoveAsyncRequest {
	s.UserImage = v
	return s
}

func (s *ImageRemoveAsyncRequest) SetUserText(v []*string) *ImageRemoveAsyncRequest {
	s.UserText = v
	return s
}

func (s *ImageRemoveAsyncRequest) Validate() error {
	return dara.Validate(s)
}
