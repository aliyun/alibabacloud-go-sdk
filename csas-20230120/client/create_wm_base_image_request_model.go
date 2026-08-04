// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWmBaseImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHeight(v int32) *CreateWmBaseImageRequest
	GetHeight() *int32
	SetImageControl(v *CreateWmBaseImageRequestImageControl) *CreateWmBaseImageRequest
	GetImageControl() *CreateWmBaseImageRequestImageControl
	SetOpacity(v int32) *CreateWmBaseImageRequest
	GetOpacity() *int32
	SetScale(v int32) *CreateWmBaseImageRequest
	GetScale() *int32
	SetWidth(v int32) *CreateWmBaseImageRequest
	GetWidth() *int32
	SetWmInfoBytesB64(v string) *CreateWmBaseImageRequest
	GetWmInfoBytesB64() *string
	SetWmInfoSize(v int64) *CreateWmBaseImageRequest
	GetWmInfoSize() *int64
	SetWmInfoUint(v string) *CreateWmBaseImageRequest
	GetWmInfoUint() *string
	SetWmType(v string) *CreateWmBaseImageRequest
	GetWmType() *string
	SetComment(v string) *CreateWmBaseImageRequest
	GetComment() *string
}

type CreateWmBaseImageRequest struct {
	// Height of the watermark image, in pixels. Valid values: 100 to 5000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1080
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// Image watermark control parameters.
	ImageControl *CreateWmBaseImageRequestImageControl `json:"ImageControl,omitempty" xml:"ImageControl,omitempty" type:"Struct"`
	// Opacity of the watermark image. Valid values: 1 to 255. Higher values mean lower transparency.
	//
	// This parameter is required.
	//
	// example:
	//
	// 255
	Opacity *int32 `json:"Opacity,omitempty" xml:"Opacity,omitempty"`
	// Scaling factor of the watermark image.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Scale *int32 `json:"Scale,omitempty" xml:"Scale,omitempty"`
	// Width of the watermark image, in pixels. Valid values: 100 to 5000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1920
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// Base64-encoded watermark information. Length: 1 to 300 characters. Do not set this parameter if you set WmInfoUint.
	//
	// example:
	//
	// aGVsbG8gc2F*****
	WmInfoBytesB64 *string `json:"WmInfoBytesB64,omitempty" xml:"WmInfoBytesB64,omitempty"`
	// Bit width of the watermark information. Default value: 32. This value must be the same during embedding and extraction. For example, if you use a 40-bit SDK to embed the watermark, set this value to 40 when extracting it.
	//
	// example:
	//
	// 32
	WmInfoSize *int64 `json:"WmInfoSize,omitempty" xml:"WmInfoSize,omitempty"`
	// Decimal-form watermark information. Do not set this parameter if you set WmInfoBytesB64.
	//
	// The valid range depends on the WmInfoSize value:
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
	// Watermark type. Valid values:
	//
	// - **PureWebappInvisible**: Web watermark.
	//
	// - **PureAppInvisible**: App watermark.
	//
	// - **PureScreenInvisible**: Screen watermark.
	//
	// - **AigcWebappInvisible**: AIGC web watermark.
	//
	// - **AigcAppInvisible**: AIGC app watermark.
	//
	// - **AigcScreenInvisible**: AIGC screen watermark.
	//
	// This parameter is required.
	//
	// example:
	//
	// PureWebappInvisible
	WmType *string `json:"WmType,omitempty" xml:"WmType,omitempty"`
	// Comments.
	//
	// example:
	//
	// 备注
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
}

func (s CreateWmBaseImageRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWmBaseImageRequest) GoString() string {
	return s.String()
}

func (s *CreateWmBaseImageRequest) GetHeight() *int32 {
	return s.Height
}

func (s *CreateWmBaseImageRequest) GetImageControl() *CreateWmBaseImageRequestImageControl {
	return s.ImageControl
}

func (s *CreateWmBaseImageRequest) GetOpacity() *int32 {
	return s.Opacity
}

func (s *CreateWmBaseImageRequest) GetScale() *int32 {
	return s.Scale
}

func (s *CreateWmBaseImageRequest) GetWidth() *int32 {
	return s.Width
}

func (s *CreateWmBaseImageRequest) GetWmInfoBytesB64() *string {
	return s.WmInfoBytesB64
}

func (s *CreateWmBaseImageRequest) GetWmInfoSize() *int64 {
	return s.WmInfoSize
}

func (s *CreateWmBaseImageRequest) GetWmInfoUint() *string {
	return s.WmInfoUint
}

func (s *CreateWmBaseImageRequest) GetWmType() *string {
	return s.WmType
}

func (s *CreateWmBaseImageRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateWmBaseImageRequest) SetHeight(v int32) *CreateWmBaseImageRequest {
	s.Height = &v
	return s
}

func (s *CreateWmBaseImageRequest) SetImageControl(v *CreateWmBaseImageRequestImageControl) *CreateWmBaseImageRequest {
	s.ImageControl = v
	return s
}

func (s *CreateWmBaseImageRequest) SetOpacity(v int32) *CreateWmBaseImageRequest {
	s.Opacity = &v
	return s
}

func (s *CreateWmBaseImageRequest) SetScale(v int32) *CreateWmBaseImageRequest {
	s.Scale = &v
	return s
}

func (s *CreateWmBaseImageRequest) SetWidth(v int32) *CreateWmBaseImageRequest {
	s.Width = &v
	return s
}

func (s *CreateWmBaseImageRequest) SetWmInfoBytesB64(v string) *CreateWmBaseImageRequest {
	s.WmInfoBytesB64 = &v
	return s
}

func (s *CreateWmBaseImageRequest) SetWmInfoSize(v int64) *CreateWmBaseImageRequest {
	s.WmInfoSize = &v
	return s
}

func (s *CreateWmBaseImageRequest) SetWmInfoUint(v string) *CreateWmBaseImageRequest {
	s.WmInfoUint = &v
	return s
}

func (s *CreateWmBaseImageRequest) SetWmType(v string) *CreateWmBaseImageRequest {
	s.WmType = &v
	return s
}

func (s *CreateWmBaseImageRequest) SetComment(v string) *CreateWmBaseImageRequest {
	s.Comment = &v
	return s
}

func (s *CreateWmBaseImageRequest) Validate() error {
	if s.ImageControl != nil {
		if err := s.ImageControl.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWmBaseImageRequestImageControl struct {
	// Logo watermark control parameters.
	LogoVisibleControl *CreateWmBaseImageRequestImageControlLogoVisibleControl `json:"LogoVisibleControl,omitempty" xml:"LogoVisibleControl,omitempty" type:"Struct"`
	// Text watermark control parameters.
	TextVisibleControl *CreateWmBaseImageRequestImageControlTextVisibleControl `json:"TextVisibleControl,omitempty" xml:"TextVisibleControl,omitempty" type:"Struct"`
}

func (s CreateWmBaseImageRequestImageControl) String() string {
	return dara.Prettify(s)
}

func (s CreateWmBaseImageRequestImageControl) GoString() string {
	return s.String()
}

func (s *CreateWmBaseImageRequestImageControl) GetLogoVisibleControl() *CreateWmBaseImageRequestImageControlLogoVisibleControl {
	return s.LogoVisibleControl
}

func (s *CreateWmBaseImageRequestImageControl) GetTextVisibleControl() *CreateWmBaseImageRequestImageControlTextVisibleControl {
	return s.TextVisibleControl
}

func (s *CreateWmBaseImageRequestImageControl) SetLogoVisibleControl(v *CreateWmBaseImageRequestImageControlLogoVisibleControl) *CreateWmBaseImageRequestImageControl {
	s.LogoVisibleControl = v
	return s
}

func (s *CreateWmBaseImageRequestImageControl) SetTextVisibleControl(v *CreateWmBaseImageRequestImageControlTextVisibleControl) *CreateWmBaseImageRequestImageControl {
	s.TextVisibleControl = v
	return s
}

func (s *CreateWmBaseImageRequestImageControl) Validate() error {
	if s.LogoVisibleControl != nil {
		if err := s.LogoVisibleControl.Validate(); err != nil {
			return err
		}
	}
	if s.TextVisibleControl != nil {
		if err := s.TextVisibleControl.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWmBaseImageRequestImageControlLogoVisibleControl struct {
	// The clockwise rotation angle of the logo watermark text. Valid values: 1 to 360.
	//
	// example:
	//
	// 30
	Angle *int64 `json:"Angle,omitempty" xml:"Angle,omitempty"`
	// Enable enhanced visible logo watermark. If enabled, the logo is converted into a watermark and added to the image.
	Enhance *bool `json:"Enhance,omitempty" xml:"Enhance,omitempty"`
	// Base64-encoded logo watermark. The logo file must be a PNG image encoded in Base64 format.
	//
	// example:
	//
	// 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
	LogoBase64 *string `json:"LogoBase64,omitempty" xml:"LogoBase64,omitempty"`
	// Applies when Mode is top-left, top-right, bottom-left, or bottom-right. Specifies the margin.
	Margin *CreateWmBaseImageRequestImageControlLogoVisibleControlMargin `json:"Margin,omitempty" xml:"Margin,omitempty" type:"Struct"`
	// The display mode for the logo watermark. Valid values:
	//
	// - **pos**: Fixed position pattern.
	//
	// - **repeat**: Tile pattern.
	//
	// example:
	//
	// pos
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// Logo watermark opacity. Valid values: 1 to 255. A higher value means lower transparency.
	//
	// example:
	//
	// 255
	Opacity *int32 `json:"Opacity,omitempty" xml:"Opacity,omitempty"`
	// The horizontal anchor point for the logo watermark. Valid values: 0 to 1. When (PosAx, PosAy) is (0, 0), the watermark anchors to the top-left corner of the text. When the value is 0.5, the watermark anchors to the centroid of the text. When (PosAx, PosAy) is (1, 1), the watermark anchors to the bottom-right corner of the text.
	//
	// example:
	//
	// 0
	PosAx *float32 `json:"PosAx,omitempty" xml:"PosAx,omitempty"`
	// The vertical anchor point for the logo watermark. The value ranges from 0 to 1. The coordinates (PosAx, PosAy) define the anchor point on the watermark. For example, (0,0) represents the top-left corner, (0.5, 0.5) represents the centroid, and (1,1) represents the bottom-right corner.
	//
	// example:
	//
	// 0
	PosAy *float32 `json:"PosAy,omitempty" xml:"PosAy,omitempty"`
	// Takes effect when Mode is set to pos. This parameter controls the horizontal position of a visible watermark, using pixel coordinates with the top-left corner as the origin.
	//
	// example:
	//
	// 0
	PosX *int64 `json:"PosX,omitempty" xml:"PosX,omitempty"`
	// Specifies the vertical position of the visible watermark in pixels. The top-left corner is the origin. This parameter is valid only when Mode is set to pos.
	//
	// example:
	//
	// 0
	PosY *int64 `json:"PosY,omitempty" xml:"PosY,omitempty"`
	// This parameter takes effect when Mode is set to repeat. It specifies the horizontal pitch for the tiled visible watermark.
	//
	// example:
	//
	// 30
	SpaceX *int64 `json:"SpaceX,omitempty" xml:"SpaceX,omitempty"`
	// Applies only when Mode is set to repeat. Controls the vertical pitch between tiled visible watermarks.
	//
	// example:
	//
	// 30
	SpaceY *int64 `json:"SpaceY,omitempty" xml:"SpaceY,omitempty"`
	// Visibility:
	//
	// **true**: Display
	//
	// **false**: Hide
	//
	// example:
	//
	// true
	Visible *bool `json:"Visible,omitempty" xml:"Visible,omitempty"`
}

func (s CreateWmBaseImageRequestImageControlLogoVisibleControl) String() string {
	return dara.Prettify(s)
}

func (s CreateWmBaseImageRequestImageControlLogoVisibleControl) GoString() string {
	return s.String()
}

func (s *CreateWmBaseImageRequestImageControlLogoVisibleControl) GetAngle() *int64 {
	return s.Angle
}

func (s *CreateWmBaseImageRequestImageControlLogoVisibleControl) GetEnhance() *bool {
	return s.Enhance
}

func (s *CreateWmBaseImageRequestImageControlLogoVisibleControl) GetLogoBase64() *string {
	return s.LogoBase64
}

func (s *CreateWmBaseImageRequestImageControlLogoVisibleControl) GetMargin() *CreateWmBaseImageRequestImageControlLogoVisibleControlMargin {
	return s.Margin
}

func (s *CreateWmBaseImageRequestImageControlLogoVisibleControl) GetMode() *string {
	return s.Mode
}

func (s *CreateWmBaseImageRequestImageControlLogoVisibleControl) GetOpacity() *int32 {
	return s.Opacity
}

func (s *CreateWmBaseImageRequestImageControlLogoVisibleControl) GetPosAx() *float32 {
	return s.PosAx
}

func (s *CreateWmBaseImageRequestImageControlLogoVisibleControl) GetPosAy() *float32 {
	return s.PosAy
}

func (s *CreateWmBaseImageRequestImageControlLogoVisibleControl) GetPosX() *int64 {
	return s.PosX
}

func (s *CreateWmBaseImageRequestImageControlLogoVisibleControl) GetPosY() *int64 {
	return s.PosY
}

func (s *CreateWmBaseImageRequestImageControlLogoVisibleControl) GetSpaceX() *int64 {
	return s.SpaceX
}

func (s *CreateWmBaseImageRequestImageControlLogoVisibleControl) GetSpaceY() *int64 {
	return s.SpaceY
}

func (s *CreateWmBaseImageRequestImageControlLogoVisibleControl) GetVisible() *bool {
	return s.Visible
}

func (s *CreateWmBaseImageRequestImageControlLogoVisibleControl) SetAngle(v int64) *CreateWmBaseImageRequestImageControlLogoVisibleControl {
	s.Angle = &v
	return s
}

func (s *CreateWmBaseImageRequestImageControlLogoVisibleControl) SetEnhance(v bool) *CreateWmBaseImageRequestImageControlLogoVisibleControl {
	s.Enhance = &v
	return s
}

func (s *CreateWmBaseImageRequestImageControlLogoVisibleControl) SetLogoBase64(v string) *CreateWmBaseImageRequestImageControlLogoVisibleControl {
	s.LogoBase64 = &v
	return s
}

func (s *CreateWmBaseImageRequestImageControlLogoVisibleControl) SetMargin(v *CreateWmBaseImageRequestImageControlLogoVisibleControlMargin) *CreateWmBaseImageRequestImageControlLogoVisibleControl {
	s.Margin = v
	return s
}

func (s *CreateWmBaseImageRequestImageControlLogoVisibleControl) SetMode(v string) *CreateWmBaseImageRequestImageControlLogoVisibleControl {
	s.Mode = &v
	return s
}

func (s *CreateWmBaseImageRequestImageControlLogoVisibleControl) SetOpacity(v int32) *CreateWmBaseImageRequestImageControlLogoVisibleControl {
	s.Opacity = &v
	return s
}

func (s *CreateWmBaseImageRequestImageControlLogoVisibleControl) SetPosAx(v float32) *CreateWmBaseImageRequestImageControlLogoVisibleControl {
	s.PosAx = &v
	return s
}

func (s *CreateWmBaseImageRequestImageControlLogoVisibleControl) SetPosAy(v float32) *CreateWmBaseImageRequestImageControlLogoVisibleControl {
	s.PosAy = &v
	return s
}

func (s *CreateWmBaseImageRequestImageControlLogoVisibleControl) SetPosX(v int64) *CreateWmBaseImageRequestImageControlLogoVisibleControl {
	s.PosX = &v
	return s
}

func (s *CreateWmBaseImageRequestImageControlLogoVisibleControl) SetPosY(v int64) *CreateWmBaseImageRequestImageControlLogoVisibleControl {
	s.PosY = &v
	return s
}

func (s *CreateWmBaseImageRequestImageControlLogoVisibleControl) SetSpaceX(v int64) *CreateWmBaseImageRequestImageControlLogoVisibleControl {
	s.SpaceX = &v
	return s
}

func (s *CreateWmBaseImageRequestImageControlLogoVisibleControl) SetSpaceY(v int64) *CreateWmBaseImageRequestImageControlLogoVisibleControl {
	s.SpaceY = &v
	return s
}

func (s *CreateWmBaseImageRequestImageControlLogoVisibleControl) SetVisible(v bool) *CreateWmBaseImageRequestImageControlLogoVisibleControl {
	s.Visible = &v
	return s
}

func (s *CreateWmBaseImageRequestImageControlLogoVisibleControl) Validate() error {
	if s.Margin != nil {
		if err := s.Margin.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWmBaseImageRequestImageControlLogoVisibleControlMargin struct {
	// Applies when Mode is bottom-left or bottom-right. The bottom margin.
	//
	// example:
	//
	// 0
	Bottom *float32 `json:"Bottom,omitempty" xml:"Bottom,omitempty"`
	// Applies when Mode is top-left or bottom-left. The left margin.
	//
	// example:
	//
	// 0
	Left *float32 `json:"Left,omitempty" xml:"Left,omitempty"`
	// The right margin. This parameter is valid only when Mode is set to top-right or bottom-right.
	//
	// example:
	//
	// 0
	Right *float32 `json:"Right,omitempty" xml:"Right,omitempty"`
	// Applies when Mode is top-left or top-right. The top margin.
	//
	// example:
	//
	// 0
	Top *float32 `json:"Top,omitempty" xml:"Top,omitempty"`
}

func (s CreateWmBaseImageRequestImageControlLogoVisibleControlMargin) String() string {
	return dara.Prettify(s)
}

func (s CreateWmBaseImageRequestImageControlLogoVisibleControlMargin) GoString() string {
	return s.String()
}

func (s *CreateWmBaseImageRequestImageControlLogoVisibleControlMargin) GetBottom() *float32 {
	return s.Bottom
}

func (s *CreateWmBaseImageRequestImageControlLogoVisibleControlMargin) GetLeft() *float32 {
	return s.Left
}

func (s *CreateWmBaseImageRequestImageControlLogoVisibleControlMargin) GetRight() *float32 {
	return s.Right
}

func (s *CreateWmBaseImageRequestImageControlLogoVisibleControlMargin) GetTop() *float32 {
	return s.Top
}

func (s *CreateWmBaseImageRequestImageControlLogoVisibleControlMargin) SetBottom(v float32) *CreateWmBaseImageRequestImageControlLogoVisibleControlMargin {
	s.Bottom = &v
	return s
}

func (s *CreateWmBaseImageRequestImageControlLogoVisibleControlMargin) SetLeft(v float32) *CreateWmBaseImageRequestImageControlLogoVisibleControlMargin {
	s.Left = &v
	return s
}

func (s *CreateWmBaseImageRequestImageControlLogoVisibleControlMargin) SetRight(v float32) *CreateWmBaseImageRequestImageControlLogoVisibleControlMargin {
	s.Right = &v
	return s
}

func (s *CreateWmBaseImageRequestImageControlLogoVisibleControlMargin) SetTop(v float32) *CreateWmBaseImageRequestImageControlLogoVisibleControlMargin {
	s.Top = &v
	return s
}

func (s *CreateWmBaseImageRequestImageControlLogoVisibleControlMargin) Validate() error {
	return dara.Validate(s)
}

type CreateWmBaseImageRequestImageControlTextVisibleControl struct {
	// Rotation angle of the text watermark, in degrees clockwise. Valid values: 0 to 360.
	//
	// example:
	//
	// 30
	Angle *int64 `json:"Angle,omitempty" xml:"Angle,omitempty"`
	// Text watermark color. Format: RGB, such as 0xFFFFFF or #FFFFFF. For example, 0x000000 or #000000 means black.
	//
	// example:
	//
	// #FF0000
	FontColor *string `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	// Font size of the text watermark. Larger values produce larger fonts.
	//
	// example:
	//
	// 30
	FontSize *int64 `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	// Takes effect only when Mode is top-left, top-right, bottom-left, or bottom-right. Margin settings.
	Margin *CreateWmBaseImageRequestImageControlTextVisibleControlMargin `json:"Margin,omitempty" xml:"Margin,omitempty" type:"Struct"`
	// Text watermark display mode. Valid values:
	//
	// - **pos**: Fixed position mode.
	//
	// - **repeat**: Tiled mode.
	//
	// example:
	//
	// pos
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// Text watermark opacity. Valid values: 1 to 255. Higher values mean less transparency.
	//
	// example:
	//
	// 255
	Opacity *int32 `json:"Opacity,omitempty" xml:"Opacity,omitempty"`
	// Horizontal anchor point for the text watermark. Valid values: 0 to 1. When (PosAx, PosAy) is (0, 0), the top-left corner of the text is used as the anchor point. When it is 0.5, the center point is used. When it is (1, 1), the bottom-right corner is used.
	//
	// example:
	//
	// 0
	PosAx *float32 `json:"PosAx,omitempty" xml:"PosAx,omitempty"`
	// Vertical anchor point for the text watermark. Valid values: 0 to 1. When (PosAx, PosAy) is (0, 0), the top-left corner of the text is used as the anchor point. When it is 0.5, the center point is used. When it is (1, 1), the bottom-right corner is used.
	//
	// example:
	//
	// 0
	PosAy *float32 `json:"PosAy,omitempty" xml:"PosAy,omitempty"`
	// Takes effect only when Mode is pos. Horizontal position of the text watermark, measured in pixels from the top-left corner.
	//
	// example:
	//
	// 0
	PosX *int64 `json:"PosX,omitempty" xml:"PosX,omitempty"`
	// Takes effect only when Mode is pos. Vertical position of the text watermark, measured in pixels from the top-left corner.
	//
	// example:
	//
	// 0
	PosY *int64 `json:"PosY,omitempty" xml:"PosY,omitempty"`
	// Takes effect only when Mode is repeat. Horizontal spacing between repeated text watermarks.
	//
	// example:
	//
	// 30
	SpaceX *int64 `json:"SpaceX,omitempty" xml:"SpaceX,omitempty"`
	// Takes effect only when Mode is repeat. Vertical spacing between repeated text watermarks.
	//
	// example:
	//
	// 0
	SpaceY *int64 `json:"SpaceY,omitempty" xml:"SpaceY,omitempty"`
	// Visibility setting:
	//
	// true: Show the watermark.
	//
	// false: Hide the watermark.
	//
	// example:
	//
	// true
	Visible *bool `json:"Visible,omitempty" xml:"Visible,omitempty"`
	// Text watermark content. Format: UTF-8 string.
	//
	// example:
	//
	// 水印文本
	VisibleText *string `json:"VisibleText,omitempty" xml:"VisibleText,omitempty"`
}

func (s CreateWmBaseImageRequestImageControlTextVisibleControl) String() string {
	return dara.Prettify(s)
}

func (s CreateWmBaseImageRequestImageControlTextVisibleControl) GoString() string {
	return s.String()
}

func (s *CreateWmBaseImageRequestImageControlTextVisibleControl) GetAngle() *int64 {
	return s.Angle
}

func (s *CreateWmBaseImageRequestImageControlTextVisibleControl) GetFontColor() *string {
	return s.FontColor
}

func (s *CreateWmBaseImageRequestImageControlTextVisibleControl) GetFontSize() *int64 {
	return s.FontSize
}

func (s *CreateWmBaseImageRequestImageControlTextVisibleControl) GetMargin() *CreateWmBaseImageRequestImageControlTextVisibleControlMargin {
	return s.Margin
}

func (s *CreateWmBaseImageRequestImageControlTextVisibleControl) GetMode() *string {
	return s.Mode
}

func (s *CreateWmBaseImageRequestImageControlTextVisibleControl) GetOpacity() *int32 {
	return s.Opacity
}

func (s *CreateWmBaseImageRequestImageControlTextVisibleControl) GetPosAx() *float32 {
	return s.PosAx
}

func (s *CreateWmBaseImageRequestImageControlTextVisibleControl) GetPosAy() *float32 {
	return s.PosAy
}

func (s *CreateWmBaseImageRequestImageControlTextVisibleControl) GetPosX() *int64 {
	return s.PosX
}

func (s *CreateWmBaseImageRequestImageControlTextVisibleControl) GetPosY() *int64 {
	return s.PosY
}

func (s *CreateWmBaseImageRequestImageControlTextVisibleControl) GetSpaceX() *int64 {
	return s.SpaceX
}

func (s *CreateWmBaseImageRequestImageControlTextVisibleControl) GetSpaceY() *int64 {
	return s.SpaceY
}

func (s *CreateWmBaseImageRequestImageControlTextVisibleControl) GetVisible() *bool {
	return s.Visible
}

func (s *CreateWmBaseImageRequestImageControlTextVisibleControl) GetVisibleText() *string {
	return s.VisibleText
}

func (s *CreateWmBaseImageRequestImageControlTextVisibleControl) SetAngle(v int64) *CreateWmBaseImageRequestImageControlTextVisibleControl {
	s.Angle = &v
	return s
}

func (s *CreateWmBaseImageRequestImageControlTextVisibleControl) SetFontColor(v string) *CreateWmBaseImageRequestImageControlTextVisibleControl {
	s.FontColor = &v
	return s
}

func (s *CreateWmBaseImageRequestImageControlTextVisibleControl) SetFontSize(v int64) *CreateWmBaseImageRequestImageControlTextVisibleControl {
	s.FontSize = &v
	return s
}

func (s *CreateWmBaseImageRequestImageControlTextVisibleControl) SetMargin(v *CreateWmBaseImageRequestImageControlTextVisibleControlMargin) *CreateWmBaseImageRequestImageControlTextVisibleControl {
	s.Margin = v
	return s
}

func (s *CreateWmBaseImageRequestImageControlTextVisibleControl) SetMode(v string) *CreateWmBaseImageRequestImageControlTextVisibleControl {
	s.Mode = &v
	return s
}

func (s *CreateWmBaseImageRequestImageControlTextVisibleControl) SetOpacity(v int32) *CreateWmBaseImageRequestImageControlTextVisibleControl {
	s.Opacity = &v
	return s
}

func (s *CreateWmBaseImageRequestImageControlTextVisibleControl) SetPosAx(v float32) *CreateWmBaseImageRequestImageControlTextVisibleControl {
	s.PosAx = &v
	return s
}

func (s *CreateWmBaseImageRequestImageControlTextVisibleControl) SetPosAy(v float32) *CreateWmBaseImageRequestImageControlTextVisibleControl {
	s.PosAy = &v
	return s
}

func (s *CreateWmBaseImageRequestImageControlTextVisibleControl) SetPosX(v int64) *CreateWmBaseImageRequestImageControlTextVisibleControl {
	s.PosX = &v
	return s
}

func (s *CreateWmBaseImageRequestImageControlTextVisibleControl) SetPosY(v int64) *CreateWmBaseImageRequestImageControlTextVisibleControl {
	s.PosY = &v
	return s
}

func (s *CreateWmBaseImageRequestImageControlTextVisibleControl) SetSpaceX(v int64) *CreateWmBaseImageRequestImageControlTextVisibleControl {
	s.SpaceX = &v
	return s
}

func (s *CreateWmBaseImageRequestImageControlTextVisibleControl) SetSpaceY(v int64) *CreateWmBaseImageRequestImageControlTextVisibleControl {
	s.SpaceY = &v
	return s
}

func (s *CreateWmBaseImageRequestImageControlTextVisibleControl) SetVisible(v bool) *CreateWmBaseImageRequestImageControlTextVisibleControl {
	s.Visible = &v
	return s
}

func (s *CreateWmBaseImageRequestImageControlTextVisibleControl) SetVisibleText(v string) *CreateWmBaseImageRequestImageControlTextVisibleControl {
	s.VisibleText = &v
	return s
}

func (s *CreateWmBaseImageRequestImageControlTextVisibleControl) Validate() error {
	if s.Margin != nil {
		if err := s.Margin.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWmBaseImageRequestImageControlTextVisibleControlMargin struct {
	// Takes effect only when Mode is bottom-left or bottom-right. Bottom margin.
	//
	// example:
	//
	// 0
	Bottom *float32 `json:"Bottom,omitempty" xml:"Bottom,omitempty"`
	// Takes effect only when Mode is top-left or bottom-left. Left margin.
	//
	// example:
	//
	// 0
	Left *float32 `json:"Left,omitempty" xml:"Left,omitempty"`
	// Takes effect only when Mode is top-right or bottom-right. Right margin.
	//
	// example:
	//
	// 0
	Right *float32 `json:"Right,omitempty" xml:"Right,omitempty"`
	// Takes effect only when Mode is top-left or top-right. Top margin.
	//
	// example:
	//
	// 0
	Top *float32 `json:"Top,omitempty" xml:"Top,omitempty"`
}

func (s CreateWmBaseImageRequestImageControlTextVisibleControlMargin) String() string {
	return dara.Prettify(s)
}

func (s CreateWmBaseImageRequestImageControlTextVisibleControlMargin) GoString() string {
	return s.String()
}

func (s *CreateWmBaseImageRequestImageControlTextVisibleControlMargin) GetBottom() *float32 {
	return s.Bottom
}

func (s *CreateWmBaseImageRequestImageControlTextVisibleControlMargin) GetLeft() *float32 {
	return s.Left
}

func (s *CreateWmBaseImageRequestImageControlTextVisibleControlMargin) GetRight() *float32 {
	return s.Right
}

func (s *CreateWmBaseImageRequestImageControlTextVisibleControlMargin) GetTop() *float32 {
	return s.Top
}

func (s *CreateWmBaseImageRequestImageControlTextVisibleControlMargin) SetBottom(v float32) *CreateWmBaseImageRequestImageControlTextVisibleControlMargin {
	s.Bottom = &v
	return s
}

func (s *CreateWmBaseImageRequestImageControlTextVisibleControlMargin) SetLeft(v float32) *CreateWmBaseImageRequestImageControlTextVisibleControlMargin {
	s.Left = &v
	return s
}

func (s *CreateWmBaseImageRequestImageControlTextVisibleControlMargin) SetRight(v float32) *CreateWmBaseImageRequestImageControlTextVisibleControlMargin {
	s.Right = &v
	return s
}

func (s *CreateWmBaseImageRequestImageControlTextVisibleControlMargin) SetTop(v float32) *CreateWmBaseImageRequestImageControlTextVisibleControlMargin {
	s.Top = &v
	return s
}

func (s *CreateWmBaseImageRequestImageControlTextVisibleControlMargin) Validate() error {
	return dara.Validate(s)
}
