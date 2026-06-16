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
}

type ImageRemoveRequest struct {
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
	NonObjectRemoveElements []*int32 `json:"NonObjectRemoveElements,omitempty" xml:"NonObjectRemoveElements,omitempty" type:"Repeated"`
	// Elements to remove from the image subject (1=transparent text blocks; 2=specific names; 3=text; 4=blemishes). Multiple selections allowed
	//
	// This parameter is required.
	//
	// example:
	//
	// [1,2]
	ObjectRemoveElements []*int32 `json:"ObjectRemoveElements,omitempty" xml:"ObjectRemoveElements,omitempty" type:"Repeated"`
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

func (s *ImageRemoveRequest) Validate() error {
	return dara.Validate(s)
}
