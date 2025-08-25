// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iAbstractEcommerceVideoAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDuration(v float32) *AbstractEcommerceVideoAdvanceRequest
	GetDuration() *float32
	SetHeight(v int32) *AbstractEcommerceVideoAdvanceRequest
	GetHeight() *int32
	SetVideoUrlObject(v io.Reader) *AbstractEcommerceVideoAdvanceRequest
	GetVideoUrlObject() io.Reader
	SetWidth(v int32) *AbstractEcommerceVideoAdvanceRequest
	GetWidth() *int32
}

type AbstractEcommerceVideoAdvanceRequest struct {
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
	VideoUrlObject io.Reader `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
	// example:
	//
	// 480
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s AbstractEcommerceVideoAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s AbstractEcommerceVideoAdvanceRequest) GoString() string {
	return s.String()
}

func (s *AbstractEcommerceVideoAdvanceRequest) GetDuration() *float32 {
	return s.Duration
}

func (s *AbstractEcommerceVideoAdvanceRequest) GetHeight() *int32 {
	return s.Height
}

func (s *AbstractEcommerceVideoAdvanceRequest) GetVideoUrlObject() io.Reader {
	return s.VideoUrlObject
}

func (s *AbstractEcommerceVideoAdvanceRequest) GetWidth() *int32 {
	return s.Width
}

func (s *AbstractEcommerceVideoAdvanceRequest) SetDuration(v float32) *AbstractEcommerceVideoAdvanceRequest {
	s.Duration = &v
	return s
}

func (s *AbstractEcommerceVideoAdvanceRequest) SetHeight(v int32) *AbstractEcommerceVideoAdvanceRequest {
	s.Height = &v
	return s
}

func (s *AbstractEcommerceVideoAdvanceRequest) SetVideoUrlObject(v io.Reader) *AbstractEcommerceVideoAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

func (s *AbstractEcommerceVideoAdvanceRequest) SetWidth(v int32) *AbstractEcommerceVideoAdvanceRequest {
	s.Width = &v
	return s
}

func (s *AbstractEcommerceVideoAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
