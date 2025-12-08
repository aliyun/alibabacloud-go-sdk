// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iFaceBeautyAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURLObject(v io.Reader) *FaceBeautyAdvanceRequest
	GetImageURLObject() io.Reader
	SetSharp(v float32) *FaceBeautyAdvanceRequest
	GetSharp() *float32
	SetSmooth(v float32) *FaceBeautyAdvanceRequest
	GetSmooth() *float32
	SetWhite(v float32) *FaceBeautyAdvanceRequest
	GetWhite() *float32
}

type FaceBeautyAdvanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/FaceBeauty/FaceBeauty4.png
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Sharp *float32 `json:"Sharp,omitempty" xml:"Sharp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.1
	Smooth *float32 `json:"Smooth,omitempty" xml:"Smooth,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.3
	White *float32 `json:"White,omitempty" xml:"White,omitempty"`
}

func (s FaceBeautyAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s FaceBeautyAdvanceRequest) GoString() string {
	return s.String()
}

func (s *FaceBeautyAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *FaceBeautyAdvanceRequest) GetSharp() *float32 {
	return s.Sharp
}

func (s *FaceBeautyAdvanceRequest) GetSmooth() *float32 {
	return s.Smooth
}

func (s *FaceBeautyAdvanceRequest) GetWhite() *float32 {
	return s.White
}

func (s *FaceBeautyAdvanceRequest) SetImageURLObject(v io.Reader) *FaceBeautyAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *FaceBeautyAdvanceRequest) SetSharp(v float32) *FaceBeautyAdvanceRequest {
	s.Sharp = &v
	return s
}

func (s *FaceBeautyAdvanceRequest) SetSmooth(v float32) *FaceBeautyAdvanceRequest {
	s.Smooth = &v
	return s
}

func (s *FaceBeautyAdvanceRequest) SetWhite(v float32) *FaceBeautyAdvanceRequest {
	s.White = &v
	return s
}

func (s *FaceBeautyAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
