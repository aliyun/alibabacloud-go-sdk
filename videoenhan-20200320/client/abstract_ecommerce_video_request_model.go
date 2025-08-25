// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbstractEcommerceVideoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDuration(v float32) *AbstractEcommerceVideoRequest
	GetDuration() *float32
	SetHeight(v int32) *AbstractEcommerceVideoRequest
	GetHeight() *int32
	SetVideoUrl(v string) *AbstractEcommerceVideoRequest
	GetVideoUrl() *string
	SetWidth(v int32) *AbstractEcommerceVideoRequest
	GetWidth() *int32
}

type AbstractEcommerceVideoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 5
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 480
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoenhan/AbstractEcommerceVideo/AbstractEcommerceVideo1.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
	// example:
	//
	// 480
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s AbstractEcommerceVideoRequest) String() string {
	return dara.Prettify(s)
}

func (s AbstractEcommerceVideoRequest) GoString() string {
	return s.String()
}

func (s *AbstractEcommerceVideoRequest) GetDuration() *float32 {
	return s.Duration
}

func (s *AbstractEcommerceVideoRequest) GetHeight() *int32 {
	return s.Height
}

func (s *AbstractEcommerceVideoRequest) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *AbstractEcommerceVideoRequest) GetWidth() *int32 {
	return s.Width
}

func (s *AbstractEcommerceVideoRequest) SetDuration(v float32) *AbstractEcommerceVideoRequest {
	s.Duration = &v
	return s
}

func (s *AbstractEcommerceVideoRequest) SetHeight(v int32) *AbstractEcommerceVideoRequest {
	s.Height = &v
	return s
}

func (s *AbstractEcommerceVideoRequest) SetVideoUrl(v string) *AbstractEcommerceVideoRequest {
	s.VideoUrl = &v
	return s
}

func (s *AbstractEcommerceVideoRequest) SetWidth(v int32) *AbstractEcommerceVideoRequest {
	s.Width = &v
	return s
}

func (s *AbstractEcommerceVideoRequest) Validate() error {
	return dara.Validate(s)
}
