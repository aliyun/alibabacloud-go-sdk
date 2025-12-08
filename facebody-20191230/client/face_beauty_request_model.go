// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFaceBeautyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageURL(v string) *FaceBeautyRequest
	GetImageURL() *string
	SetSharp(v float32) *FaceBeautyRequest
	GetSharp() *float32
	SetSmooth(v float32) *FaceBeautyRequest
	GetSmooth() *float32
	SetWhite(v float32) *FaceBeautyRequest
	GetWhite() *float32
}

type FaceBeautyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/facebody/FaceBeauty/FaceBeauty4.png
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
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

func (s FaceBeautyRequest) String() string {
	return dara.Prettify(s)
}

func (s FaceBeautyRequest) GoString() string {
	return s.String()
}

func (s *FaceBeautyRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *FaceBeautyRequest) GetSharp() *float32 {
	return s.Sharp
}

func (s *FaceBeautyRequest) GetSmooth() *float32 {
	return s.Smooth
}

func (s *FaceBeautyRequest) GetWhite() *float32 {
	return s.White
}

func (s *FaceBeautyRequest) SetImageURL(v string) *FaceBeautyRequest {
	s.ImageURL = &v
	return s
}

func (s *FaceBeautyRequest) SetSharp(v float32) *FaceBeautyRequest {
	s.Sharp = &v
	return s
}

func (s *FaceBeautyRequest) SetSmooth(v float32) *FaceBeautyRequest {
	s.Smooth = &v
	return s
}

func (s *FaceBeautyRequest) SetWhite(v float32) *FaceBeautyRequest {
	s.White = &v
	return s
}

func (s *FaceBeautyRequest) Validate() error {
	return dara.Validate(s)
}
