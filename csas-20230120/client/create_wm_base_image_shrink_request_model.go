// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWmBaseImageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHeight(v int32) *CreateWmBaseImageShrinkRequest
	GetHeight() *int32
	SetImageControlShrink(v string) *CreateWmBaseImageShrinkRequest
	GetImageControlShrink() *string
	SetOpacity(v int32) *CreateWmBaseImageShrinkRequest
	GetOpacity() *int32
	SetScale(v int32) *CreateWmBaseImageShrinkRequest
	GetScale() *int32
	SetWidth(v int32) *CreateWmBaseImageShrinkRequest
	GetWidth() *int32
	SetWmInfoBytesB64(v string) *CreateWmBaseImageShrinkRequest
	GetWmInfoBytesB64() *string
	SetWmInfoSize(v int64) *CreateWmBaseImageShrinkRequest
	GetWmInfoSize() *int64
	SetWmInfoUint(v string) *CreateWmBaseImageShrinkRequest
	GetWmInfoUint() *string
	SetWmType(v string) *CreateWmBaseImageShrinkRequest
	GetWmType() *string
}

type CreateWmBaseImageShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1080
	Height             *int32  `json:"Height,omitempty" xml:"Height,omitempty"`
	ImageControlShrink *string `json:"ImageControl,omitempty" xml:"ImageControl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 255
	Opacity *int32 `json:"Opacity,omitempty" xml:"Opacity,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Scale *int32 `json:"Scale,omitempty" xml:"Scale,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1920
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// example:
	//
	// aGVsbG8gc2F*****
	WmInfoBytesB64 *string `json:"WmInfoBytesB64,omitempty" xml:"WmInfoBytesB64,omitempty"`
	// example:
	//
	// 32
	WmInfoSize *int64 `json:"WmInfoSize,omitempty" xml:"WmInfoSize,omitempty"`
	// example:
	//
	// 12*****
	WmInfoUint *string `json:"WmInfoUint,omitempty" xml:"WmInfoUint,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PureWebappInvisible
	WmType *string `json:"WmType,omitempty" xml:"WmType,omitempty"`
}

func (s CreateWmBaseImageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWmBaseImageShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateWmBaseImageShrinkRequest) GetHeight() *int32 {
	return s.Height
}

func (s *CreateWmBaseImageShrinkRequest) GetImageControlShrink() *string {
	return s.ImageControlShrink
}

func (s *CreateWmBaseImageShrinkRequest) GetOpacity() *int32 {
	return s.Opacity
}

func (s *CreateWmBaseImageShrinkRequest) GetScale() *int32 {
	return s.Scale
}

func (s *CreateWmBaseImageShrinkRequest) GetWidth() *int32 {
	return s.Width
}

func (s *CreateWmBaseImageShrinkRequest) GetWmInfoBytesB64() *string {
	return s.WmInfoBytesB64
}

func (s *CreateWmBaseImageShrinkRequest) GetWmInfoSize() *int64 {
	return s.WmInfoSize
}

func (s *CreateWmBaseImageShrinkRequest) GetWmInfoUint() *string {
	return s.WmInfoUint
}

func (s *CreateWmBaseImageShrinkRequest) GetWmType() *string {
	return s.WmType
}

func (s *CreateWmBaseImageShrinkRequest) SetHeight(v int32) *CreateWmBaseImageShrinkRequest {
	s.Height = &v
	return s
}

func (s *CreateWmBaseImageShrinkRequest) SetImageControlShrink(v string) *CreateWmBaseImageShrinkRequest {
	s.ImageControlShrink = &v
	return s
}

func (s *CreateWmBaseImageShrinkRequest) SetOpacity(v int32) *CreateWmBaseImageShrinkRequest {
	s.Opacity = &v
	return s
}

func (s *CreateWmBaseImageShrinkRequest) SetScale(v int32) *CreateWmBaseImageShrinkRequest {
	s.Scale = &v
	return s
}

func (s *CreateWmBaseImageShrinkRequest) SetWidth(v int32) *CreateWmBaseImageShrinkRequest {
	s.Width = &v
	return s
}

func (s *CreateWmBaseImageShrinkRequest) SetWmInfoBytesB64(v string) *CreateWmBaseImageShrinkRequest {
	s.WmInfoBytesB64 = &v
	return s
}

func (s *CreateWmBaseImageShrinkRequest) SetWmInfoSize(v int64) *CreateWmBaseImageShrinkRequest {
	s.WmInfoSize = &v
	return s
}

func (s *CreateWmBaseImageShrinkRequest) SetWmInfoUint(v string) *CreateWmBaseImageShrinkRequest {
	s.WmInfoUint = &v
	return s
}

func (s *CreateWmBaseImageShrinkRequest) SetWmType(v string) *CreateWmBaseImageShrinkRequest {
	s.WmType = &v
	return s
}

func (s *CreateWmBaseImageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
