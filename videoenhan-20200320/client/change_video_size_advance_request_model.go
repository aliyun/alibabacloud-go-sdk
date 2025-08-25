// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iChangeVideoSizeAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetB(v int32) *ChangeVideoSizeAdvanceRequest
	GetB() *int32
	SetCropType(v string) *ChangeVideoSizeAdvanceRequest
	GetCropType() *string
	SetFillType(v string) *ChangeVideoSizeAdvanceRequest
	GetFillType() *string
	SetG(v int32) *ChangeVideoSizeAdvanceRequest
	GetG() *int32
	SetHeight(v int32) *ChangeVideoSizeAdvanceRequest
	GetHeight() *int32
	SetR(v int32) *ChangeVideoSizeAdvanceRequest
	GetR() *int32
	SetTightness(v float32) *ChangeVideoSizeAdvanceRequest
	GetTightness() *float32
	SetVideoUrlObject(v io.Reader) *ChangeVideoSizeAdvanceRequest
	GetVideoUrlObject() io.Reader
	SetWidth(v int32) *ChangeVideoSizeAdvanceRequest
	GetWidth() *int32
}

type ChangeVideoSizeAdvanceRequest struct {
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
	VideoUrlObject io.Reader `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 600
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s ChangeVideoSizeAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeVideoSizeAdvanceRequest) GoString() string {
	return s.String()
}

func (s *ChangeVideoSizeAdvanceRequest) GetB() *int32 {
	return s.B
}

func (s *ChangeVideoSizeAdvanceRequest) GetCropType() *string {
	return s.CropType
}

func (s *ChangeVideoSizeAdvanceRequest) GetFillType() *string {
	return s.FillType
}

func (s *ChangeVideoSizeAdvanceRequest) GetG() *int32 {
	return s.G
}

func (s *ChangeVideoSizeAdvanceRequest) GetHeight() *int32 {
	return s.Height
}

func (s *ChangeVideoSizeAdvanceRequest) GetR() *int32 {
	return s.R
}

func (s *ChangeVideoSizeAdvanceRequest) GetTightness() *float32 {
	return s.Tightness
}

func (s *ChangeVideoSizeAdvanceRequest) GetVideoUrlObject() io.Reader {
	return s.VideoUrlObject
}

func (s *ChangeVideoSizeAdvanceRequest) GetWidth() *int32 {
	return s.Width
}

func (s *ChangeVideoSizeAdvanceRequest) SetB(v int32) *ChangeVideoSizeAdvanceRequest {
	s.B = &v
	return s
}

func (s *ChangeVideoSizeAdvanceRequest) SetCropType(v string) *ChangeVideoSizeAdvanceRequest {
	s.CropType = &v
	return s
}

func (s *ChangeVideoSizeAdvanceRequest) SetFillType(v string) *ChangeVideoSizeAdvanceRequest {
	s.FillType = &v
	return s
}

func (s *ChangeVideoSizeAdvanceRequest) SetG(v int32) *ChangeVideoSizeAdvanceRequest {
	s.G = &v
	return s
}

func (s *ChangeVideoSizeAdvanceRequest) SetHeight(v int32) *ChangeVideoSizeAdvanceRequest {
	s.Height = &v
	return s
}

func (s *ChangeVideoSizeAdvanceRequest) SetR(v int32) *ChangeVideoSizeAdvanceRequest {
	s.R = &v
	return s
}

func (s *ChangeVideoSizeAdvanceRequest) SetTightness(v float32) *ChangeVideoSizeAdvanceRequest {
	s.Tightness = &v
	return s
}

func (s *ChangeVideoSizeAdvanceRequest) SetVideoUrlObject(v io.Reader) *ChangeVideoSizeAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

func (s *ChangeVideoSizeAdvanceRequest) SetWidth(v int32) *ChangeVideoSizeAdvanceRequest {
	s.Width = &v
	return s
}

func (s *ChangeVideoSizeAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
