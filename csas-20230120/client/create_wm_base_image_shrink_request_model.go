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
	SetComment(v string) *CreateWmBaseImageShrinkRequest
	GetComment() *string
}

type CreateWmBaseImageShrinkRequest struct {
	// The height of the watermark image, in pixels. Valid values: 100 to 5000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1080
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The image watermark control parameters.
	ImageControlShrink *string `json:"ImageControl,omitempty" xml:"ImageControl,omitempty"`
	// The opacity of the watermark image. Valid values: 1 to 255. A larger value indicates lower transparency.
	//
	// This parameter is required.
	//
	// example:
	//
	// 255
	Opacity *int32 `json:"Opacity,omitempty" xml:"Opacity,omitempty"`
	// The scaling ratio of the watermark image.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Scale *int32 `json:"Scale,omitempty" xml:"Scale,omitempty"`
	// The width of the watermark image, in pixels. Valid values: 100 to 5000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1920
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// The watermark information in Base64-encoded string format. The length is 1 to 300 characters. If this parameter is set, the WmInfoUint parameter cannot be set.
	//
	// example:
	//
	// aGVsbG8gc2F*****
	WmInfoBytesB64 *string `json:"WmInfoBytesB64,omitempty" xml:"WmInfoBytesB64,omitempty"`
	// The bit width of the watermark information capacity. Default value: 32. This parameter must be consistent between embedding and extraction. For example, if the SDK used for embedding is 40-bit, set this parameter to 40 during extraction as well.
	//
	// example:
	//
	// 32
	WmInfoSize *int64 `json:"WmInfoSize,omitempty" xml:"WmInfoSize,omitempty"`
	// The watermark information in decimal number format. If this parameter is set, WmInfoBytesB64 cannot be set.
	//
	// The valid range depends on the WmInfoSize parameter:
	//
	// - If WmInfoSize is **32**, the valid range is 1 to 4294967295.
	//
	// - If WmInfoSize is **40**, the valid range is 1 to 1099511627775.
	//
	// - If WmInfoSize is **64**, the valid range is 1 to 18446744073709551615.
	//
	// example:
	//
	// 12*****
	WmInfoUint *string `json:"WmInfoUint,omitempty" xml:"WmInfoUint,omitempty"`
	// The watermark type. Valid values:
	//
	// - **PureWebappInvisible**: web watermark.
	//
	// - **PureAppInvisible**: App watermark.
	//
	// - **PureScreenInvisible**: screen watermark.
	//
	// - **AigcWebappInvisible**: AIGC web watermark.
	//
	// - **AigcAppInvisible**: AIGC App watermark.
	//
	// - **AigcScreenInvisible**: AIGC screen watermark.
	//
	// This parameter is required.
	//
	// example:
	//
	// PureWebappInvisible
	WmType *string `json:"WmType,omitempty" xml:"WmType,omitempty"`
	// The remarks.
	//
	// example:
	//
	// Remarks
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
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

func (s *CreateWmBaseImageShrinkRequest) GetComment() *string {
	return s.Comment
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

func (s *CreateWmBaseImageShrinkRequest) SetComment(v string) *CreateWmBaseImageShrinkRequest {
	s.Comment = &v
	return s
}

func (s *CreateWmBaseImageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
