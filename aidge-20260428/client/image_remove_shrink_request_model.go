// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageRemoveShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageUrl(v string) *ImageRemoveShrinkRequest
	GetImageUrl() *string
	SetMask(v string) *ImageRemoveShrinkRequest
	GetMask() *string
	SetNonObjectRemoveElementsShrink(v string) *ImageRemoveShrinkRequest
	GetNonObjectRemoveElementsShrink() *string
	SetObjectRemoveElementsShrink(v string) *ImageRemoveShrinkRequest
	GetObjectRemoveElementsShrink() *string
}

type ImageRemoveShrinkRequest struct {
	// URL of the image to be processed (mutually exclusive with ImageBase64)
	//
	// This parameter is required.
	//
	// example:
	//
	// https://img.alicdn.com/bao/uploaded/i2/xxx.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// Specific removal area in RLE format. If provided, this takes priority and the remove parameters are ignored
	//
	// example:
	//
	// null
	Mask *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	// Elements to remove from the non-subject area of the image (1=transparent text blocks; 2=specific names; 3=text; 4=blemishes). Multiple selections allowed
	//
	// example:
	//
	// [1,2]
	NonObjectRemoveElementsShrink *string `json:"NonObjectRemoveElements,omitempty" xml:"NonObjectRemoveElements,omitempty"`
	// Elements to remove from the image subject (1=transparent text blocks; 2=specific names; 3=text; 4=blemishes). Multiple selections allowed
	//
	// This parameter is required.
	//
	// example:
	//
	// [1,2]
	ObjectRemoveElementsShrink *string `json:"ObjectRemoveElements,omitempty" xml:"ObjectRemoveElements,omitempty"`
}

func (s ImageRemoveShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ImageRemoveShrinkRequest) GoString() string {
	return s.String()
}

func (s *ImageRemoveShrinkRequest) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ImageRemoveShrinkRequest) GetMask() *string {
	return s.Mask
}

func (s *ImageRemoveShrinkRequest) GetNonObjectRemoveElementsShrink() *string {
	return s.NonObjectRemoveElementsShrink
}

func (s *ImageRemoveShrinkRequest) GetObjectRemoveElementsShrink() *string {
	return s.ObjectRemoveElementsShrink
}

func (s *ImageRemoveShrinkRequest) SetImageUrl(v string) *ImageRemoveShrinkRequest {
	s.ImageUrl = &v
	return s
}

func (s *ImageRemoveShrinkRequest) SetMask(v string) *ImageRemoveShrinkRequest {
	s.Mask = &v
	return s
}

func (s *ImageRemoveShrinkRequest) SetNonObjectRemoveElementsShrink(v string) *ImageRemoveShrinkRequest {
	s.NonObjectRemoveElementsShrink = &v
	return s
}

func (s *ImageRemoveShrinkRequest) SetObjectRemoveElementsShrink(v string) *ImageRemoveShrinkRequest {
	s.ObjectRemoveElementsShrink = &v
	return s
}

func (s *ImageRemoveShrinkRequest) Validate() error {
	return dara.Validate(s)
}
