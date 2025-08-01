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
	// This parameter is required.
	//
	// example:
	//
	// 1080
	Height       *int32                                `json:"Height,omitempty" xml:"Height,omitempty"`
	ImageControl *CreateWmBaseImageRequestImageControl `json:"ImageControl,omitempty" xml:"ImageControl,omitempty" type:"Struct"`
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
	WmType  *string `json:"WmType,omitempty" xml:"WmType,omitempty"`
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
	return dara.Validate(s)
}

type CreateWmBaseImageRequestImageControl struct {
	LogoVisibleControl *CreateWmBaseImageRequestImageControlLogoVisibleControl `json:"LogoVisibleControl,omitempty" xml:"LogoVisibleControl,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type CreateWmBaseImageRequestImageControlLogoVisibleControl struct {
	Angle      *int64                                                        `json:"Angle,omitempty" xml:"Angle,omitempty"`
	Enhance    *bool                                                         `json:"Enhance,omitempty" xml:"Enhance,omitempty"`
	LogoBase64 *string                                                       `json:"LogoBase64,omitempty" xml:"LogoBase64,omitempty"`
	Margin     *CreateWmBaseImageRequestImageControlLogoVisibleControlMargin `json:"Margin,omitempty" xml:"Margin,omitempty" type:"Struct"`
	Mode       *string                                                       `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Opacity    *int32                                                        `json:"Opacity,omitempty" xml:"Opacity,omitempty"`
	PosAx      *float32                                                      `json:"PosAx,omitempty" xml:"PosAx,omitempty"`
	PosAy      *float32                                                      `json:"PosAy,omitempty" xml:"PosAy,omitempty"`
	PosX       *int64                                                        `json:"PosX,omitempty" xml:"PosX,omitempty"`
	PosY       *int64                                                        `json:"PosY,omitempty" xml:"PosY,omitempty"`
	SpaceX     *int64                                                        `json:"SpaceX,omitempty" xml:"SpaceX,omitempty"`
	SpaceY     *int64                                                        `json:"SpaceY,omitempty" xml:"SpaceY,omitempty"`
	Visible    *bool                                                         `json:"Visible,omitempty" xml:"Visible,omitempty"`
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
	return dara.Validate(s)
}

type CreateWmBaseImageRequestImageControlLogoVisibleControlMargin struct {
	Bottom *float32 `json:"Bottom,omitempty" xml:"Bottom,omitempty"`
	Left   *float32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Right  *float32 `json:"Right,omitempty" xml:"Right,omitempty"`
	Top    *float32 `json:"Top,omitempty" xml:"Top,omitempty"`
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
	Angle       *int64                                                        `json:"Angle,omitempty" xml:"Angle,omitempty"`
	FontColor   *string                                                       `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	FontSize    *int64                                                        `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	Margin      *CreateWmBaseImageRequestImageControlTextVisibleControlMargin `json:"Margin,omitempty" xml:"Margin,omitempty" type:"Struct"`
	Mode        *string                                                       `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Opacity     *int32                                                        `json:"Opacity,omitempty" xml:"Opacity,omitempty"`
	PosAx       *float32                                                      `json:"PosAx,omitempty" xml:"PosAx,omitempty"`
	PosAy       *float32                                                      `json:"PosAy,omitempty" xml:"PosAy,omitempty"`
	PosX        *int64                                                        `json:"PosX,omitempty" xml:"PosX,omitempty"`
	PosY        *int64                                                        `json:"PosY,omitempty" xml:"PosY,omitempty"`
	SpaceX      *int64                                                        `json:"SpaceX,omitempty" xml:"SpaceX,omitempty"`
	SpaceY      *int64                                                        `json:"SpaceY,omitempty" xml:"SpaceY,omitempty"`
	Visible     *bool                                                         `json:"Visible,omitempty" xml:"Visible,omitempty"`
	VisibleText *string                                                       `json:"VisibleText,omitempty" xml:"VisibleText,omitempty"`
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
	return dara.Validate(s)
}

type CreateWmBaseImageRequestImageControlTextVisibleControlMargin struct {
	Bottom *float32 `json:"Bottom,omitempty" xml:"Bottom,omitempty"`
	Left   *float32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Right  *float32 `json:"Right,omitempty" xml:"Right,omitempty"`
	Top    *float32 `json:"Top,omitempty" xml:"Top,omitempty"`
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
