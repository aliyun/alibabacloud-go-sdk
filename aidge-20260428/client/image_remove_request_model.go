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
	// The URL of the image to process. This parameter is mutually exclusive with ImageBase64. You must specify one of them.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://img.alicdn.com/bao/uploaded/i2/xxx.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The specific erasure region in RLE format. If this parameter is specified, it takes priority and the remove parameters are ignored.
	//
	// example:
	//
	// null
	Mask *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	// The elements to remove from the non-subject area of the image. Valid values:
	//
	// - 1: transparent text block
	//
	// - 2: specific name
	//
	// - 3: text
	//
	// - 4: image blemish
	//
	// You can specify multiple values.
	//
	// example:
	//
	// [1,2]
	NonObjectRemoveElements []*int32 `json:"NonObjectRemoveElements,omitempty" xml:"NonObjectRemoveElements,omitempty" type:"Repeated"`
	// The elements to remove from the image subject area. Valid values:
	//
	// - 1: transparent text block
	//
	// - 2: specific name
	//
	// - 3: text
	//
	// - 4: image blemish
	//
	// You can specify multiple values.
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
