// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeVideoSizeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetB(v int32) *ChangeVideoSizeRequest
	GetB() *int32
	SetCropType(v string) *ChangeVideoSizeRequest
	GetCropType() *string
	SetFillType(v string) *ChangeVideoSizeRequest
	GetFillType() *string
	SetG(v int32) *ChangeVideoSizeRequest
	GetG() *int32
	SetHeight(v int32) *ChangeVideoSizeRequest
	GetHeight() *int32
	SetR(v int32) *ChangeVideoSizeRequest
	GetR() *int32
	SetTightness(v float32) *ChangeVideoSizeRequest
	GetTightness() *float32
	SetVideoUrl(v string) *ChangeVideoSizeRequest
	GetVideoUrl() *string
	SetWidth(v int32) *ChangeVideoSizeRequest
	GetWidth() *int32
}

type ChangeVideoSizeRequest struct {
	// example:
	//
	// 0
	B *int32 `json:"B,omitempty" xml:"B,omitempty"`
	// example:
	//
	// smart
	CropType *string `json:"CropType,omitempty" xml:"CropType,omitempty"`
	// example:
	//
	// image
	FillType *string `json:"FillType,omitempty" xml:"FillType,omitempty"`
	// example:
	//
	// 0
	G *int32 `json:"G,omitempty" xml:"G,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1000
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 0
	R *int32 `json:"R,omitempty" xml:"R,omitempty"`
	// example:
	//
	// 0.5
	Tightness *float32 `json:"Tightness,omitempty" xml:"Tightness,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoenhan/ChangeVideoSize/ChangeVideoSize1.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 600
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s ChangeVideoSizeRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeVideoSizeRequest) GoString() string {
	return s.String()
}

func (s *ChangeVideoSizeRequest) GetB() *int32 {
	return s.B
}

func (s *ChangeVideoSizeRequest) GetCropType() *string {
	return s.CropType
}

func (s *ChangeVideoSizeRequest) GetFillType() *string {
	return s.FillType
}

func (s *ChangeVideoSizeRequest) GetG() *int32 {
	return s.G
}

func (s *ChangeVideoSizeRequest) GetHeight() *int32 {
	return s.Height
}

func (s *ChangeVideoSizeRequest) GetR() *int32 {
	return s.R
}

func (s *ChangeVideoSizeRequest) GetTightness() *float32 {
	return s.Tightness
}

func (s *ChangeVideoSizeRequest) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *ChangeVideoSizeRequest) GetWidth() *int32 {
	return s.Width
}

func (s *ChangeVideoSizeRequest) SetB(v int32) *ChangeVideoSizeRequest {
	s.B = &v
	return s
}

func (s *ChangeVideoSizeRequest) SetCropType(v string) *ChangeVideoSizeRequest {
	s.CropType = &v
	return s
}

func (s *ChangeVideoSizeRequest) SetFillType(v string) *ChangeVideoSizeRequest {
	s.FillType = &v
	return s
}

func (s *ChangeVideoSizeRequest) SetG(v int32) *ChangeVideoSizeRequest {
	s.G = &v
	return s
}

func (s *ChangeVideoSizeRequest) SetHeight(v int32) *ChangeVideoSizeRequest {
	s.Height = &v
	return s
}

func (s *ChangeVideoSizeRequest) SetR(v int32) *ChangeVideoSizeRequest {
	s.R = &v
	return s
}

func (s *ChangeVideoSizeRequest) SetTightness(v float32) *ChangeVideoSizeRequest {
	s.Tightness = &v
	return s
}

func (s *ChangeVideoSizeRequest) SetVideoUrl(v string) *ChangeVideoSizeRequest {
	s.VideoUrl = &v
	return s
}

func (s *ChangeVideoSizeRequest) SetWidth(v int32) *ChangeVideoSizeRequest {
	s.Width = &v
	return s
}

func (s *ChangeVideoSizeRequest) Validate() error {
	return dara.Validate(s)
}
