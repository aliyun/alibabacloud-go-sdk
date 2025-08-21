// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWmEmbedTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCsvControl(v *CreateWmEmbedTaskRequestCsvControl) *CreateWmEmbedTaskRequest
	GetCsvControl() *CreateWmEmbedTaskRequestCsvControl
	SetDocumentControl(v *CreateWmEmbedTaskRequestDocumentControl) *CreateWmEmbedTaskRequest
	GetDocumentControl() *CreateWmEmbedTaskRequestDocumentControl
	SetFileUrl(v string) *CreateWmEmbedTaskRequest
	GetFileUrl() *string
	SetFilename(v string) *CreateWmEmbedTaskRequest
	GetFilename() *string
	SetImageControl(v *CreateWmEmbedTaskRequestImageControl) *CreateWmEmbedTaskRequest
	GetImageControl() *CreateWmEmbedTaskRequestImageControl
	SetImageEmbedJpegQuality(v int64) *CreateWmEmbedTaskRequest
	GetImageEmbedJpegQuality() *int64
	SetImageEmbedLevel(v int64) *CreateWmEmbedTaskRequest
	GetImageEmbedLevel() *int64
	SetVideoBitrate(v string) *CreateWmEmbedTaskRequest
	GetVideoBitrate() *string
	SetVideoIsLong(v bool) *CreateWmEmbedTaskRequest
	GetVideoIsLong() *bool
	SetWmInfoBytesB64(v string) *CreateWmEmbedTaskRequest
	GetWmInfoBytesB64() *string
	SetWmInfoSize(v int64) *CreateWmEmbedTaskRequest
	GetWmInfoSize() *int64
	SetWmInfoUint(v string) *CreateWmEmbedTaskRequest
	GetWmInfoUint() *string
	SetWmType(v string) *CreateWmEmbedTaskRequest
	GetWmType() *string
}

type CreateWmEmbedTaskRequest struct {
	CsvControl      *CreateWmEmbedTaskRequestCsvControl      `json:"CsvControl,omitempty" xml:"CsvControl,omitempty" type:"Struct"`
	DocumentControl *CreateWmEmbedTaskRequestDocumentControl `json:"DocumentControl,omitempty" xml:"DocumentControl,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/abc****.pdf
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abc****.pdf
	Filename     *string                               `json:"Filename,omitempty" xml:"Filename,omitempty"`
	ImageControl *CreateWmEmbedTaskRequestImageControl `json:"ImageControl,omitempty" xml:"ImageControl,omitempty" type:"Struct"`
	// example:
	//
	// 95
	ImageEmbedJpegQuality *int64 `json:"ImageEmbedJpegQuality,omitempty" xml:"ImageEmbedJpegQuality,omitempty"`
	// example:
	//
	// 2
	ImageEmbedLevel *int64 `json:"ImageEmbedLevel,omitempty" xml:"ImageEmbedLevel,omitempty"`
	// example:
	//
	// 3000k
	VideoBitrate *string `json:"VideoBitrate,omitempty" xml:"VideoBitrate,omitempty"`
	// example:
	//
	// false
	VideoIsLong *bool `json:"VideoIsLong,omitempty" xml:"VideoIsLong,omitempty"`
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
	// 123***
	WmInfoUint *string `json:"WmInfoUint,omitempty" xml:"WmInfoUint,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PureDocument
	WmType *string `json:"WmType,omitempty" xml:"WmType,omitempty"`
}

func (s CreateWmEmbedTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWmEmbedTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskRequest) GetCsvControl() *CreateWmEmbedTaskRequestCsvControl {
	return s.CsvControl
}

func (s *CreateWmEmbedTaskRequest) GetDocumentControl() *CreateWmEmbedTaskRequestDocumentControl {
	return s.DocumentControl
}

func (s *CreateWmEmbedTaskRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *CreateWmEmbedTaskRequest) GetFilename() *string {
	return s.Filename
}

func (s *CreateWmEmbedTaskRequest) GetImageControl() *CreateWmEmbedTaskRequestImageControl {
	return s.ImageControl
}

func (s *CreateWmEmbedTaskRequest) GetImageEmbedJpegQuality() *int64 {
	return s.ImageEmbedJpegQuality
}

func (s *CreateWmEmbedTaskRequest) GetImageEmbedLevel() *int64 {
	return s.ImageEmbedLevel
}

func (s *CreateWmEmbedTaskRequest) GetVideoBitrate() *string {
	return s.VideoBitrate
}

func (s *CreateWmEmbedTaskRequest) GetVideoIsLong() *bool {
	return s.VideoIsLong
}

func (s *CreateWmEmbedTaskRequest) GetWmInfoBytesB64() *string {
	return s.WmInfoBytesB64
}

func (s *CreateWmEmbedTaskRequest) GetWmInfoSize() *int64 {
	return s.WmInfoSize
}

func (s *CreateWmEmbedTaskRequest) GetWmInfoUint() *string {
	return s.WmInfoUint
}

func (s *CreateWmEmbedTaskRequest) GetWmType() *string {
	return s.WmType
}

func (s *CreateWmEmbedTaskRequest) SetCsvControl(v *CreateWmEmbedTaskRequestCsvControl) *CreateWmEmbedTaskRequest {
	s.CsvControl = v
	return s
}

func (s *CreateWmEmbedTaskRequest) SetDocumentControl(v *CreateWmEmbedTaskRequestDocumentControl) *CreateWmEmbedTaskRequest {
	s.DocumentControl = v
	return s
}

func (s *CreateWmEmbedTaskRequest) SetFileUrl(v string) *CreateWmEmbedTaskRequest {
	s.FileUrl = &v
	return s
}

func (s *CreateWmEmbedTaskRequest) SetFilename(v string) *CreateWmEmbedTaskRequest {
	s.Filename = &v
	return s
}

func (s *CreateWmEmbedTaskRequest) SetImageControl(v *CreateWmEmbedTaskRequestImageControl) *CreateWmEmbedTaskRequest {
	s.ImageControl = v
	return s
}

func (s *CreateWmEmbedTaskRequest) SetImageEmbedJpegQuality(v int64) *CreateWmEmbedTaskRequest {
	s.ImageEmbedJpegQuality = &v
	return s
}

func (s *CreateWmEmbedTaskRequest) SetImageEmbedLevel(v int64) *CreateWmEmbedTaskRequest {
	s.ImageEmbedLevel = &v
	return s
}

func (s *CreateWmEmbedTaskRequest) SetVideoBitrate(v string) *CreateWmEmbedTaskRequest {
	s.VideoBitrate = &v
	return s
}

func (s *CreateWmEmbedTaskRequest) SetVideoIsLong(v bool) *CreateWmEmbedTaskRequest {
	s.VideoIsLong = &v
	return s
}

func (s *CreateWmEmbedTaskRequest) SetWmInfoBytesB64(v string) *CreateWmEmbedTaskRequest {
	s.WmInfoBytesB64 = &v
	return s
}

func (s *CreateWmEmbedTaskRequest) SetWmInfoSize(v int64) *CreateWmEmbedTaskRequest {
	s.WmInfoSize = &v
	return s
}

func (s *CreateWmEmbedTaskRequest) SetWmInfoUint(v string) *CreateWmEmbedTaskRequest {
	s.WmInfoUint = &v
	return s
}

func (s *CreateWmEmbedTaskRequest) SetWmType(v string) *CreateWmEmbedTaskRequest {
	s.WmType = &v
	return s
}

func (s *CreateWmEmbedTaskRequest) Validate() error {
	return dara.Validate(s)
}

type CreateWmEmbedTaskRequestCsvControl struct {
	EmbedBitsNumberInEachTime *int64  `json:"EmbedBitsNumberInEachTime,omitempty" xml:"EmbedBitsNumberInEachTime,omitempty"`
	EmbedColumn               *int64  `json:"EmbedColumn,omitempty" xml:"EmbedColumn,omitempty"`
	EmbedDensity              *string `json:"EmbedDensity,omitempty" xml:"EmbedDensity,omitempty"`
	EmbedPrecision            *int64  `json:"EmbedPrecision,omitempty" xml:"EmbedPrecision,omitempty"`
	EmbedTimePosition         *string `json:"EmbedTimePosition,omitempty" xml:"EmbedTimePosition,omitempty"`
	Method                    *string `json:"Method,omitempty" xml:"Method,omitempty"`
	TimeFormat                *string `json:"TimeFormat,omitempty" xml:"TimeFormat,omitempty"`
}

func (s CreateWmEmbedTaskRequestCsvControl) String() string {
	return dara.Prettify(s)
}

func (s CreateWmEmbedTaskRequestCsvControl) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskRequestCsvControl) GetEmbedBitsNumberInEachTime() *int64 {
	return s.EmbedBitsNumberInEachTime
}

func (s *CreateWmEmbedTaskRequestCsvControl) GetEmbedColumn() *int64 {
	return s.EmbedColumn
}

func (s *CreateWmEmbedTaskRequestCsvControl) GetEmbedDensity() *string {
	return s.EmbedDensity
}

func (s *CreateWmEmbedTaskRequestCsvControl) GetEmbedPrecision() *int64 {
	return s.EmbedPrecision
}

func (s *CreateWmEmbedTaskRequestCsvControl) GetEmbedTimePosition() *string {
	return s.EmbedTimePosition
}

func (s *CreateWmEmbedTaskRequestCsvControl) GetMethod() *string {
	return s.Method
}

func (s *CreateWmEmbedTaskRequestCsvControl) GetTimeFormat() *string {
	return s.TimeFormat
}

func (s *CreateWmEmbedTaskRequestCsvControl) SetEmbedBitsNumberInEachTime(v int64) *CreateWmEmbedTaskRequestCsvControl {
	s.EmbedBitsNumberInEachTime = &v
	return s
}

func (s *CreateWmEmbedTaskRequestCsvControl) SetEmbedColumn(v int64) *CreateWmEmbedTaskRequestCsvControl {
	s.EmbedColumn = &v
	return s
}

func (s *CreateWmEmbedTaskRequestCsvControl) SetEmbedDensity(v string) *CreateWmEmbedTaskRequestCsvControl {
	s.EmbedDensity = &v
	return s
}

func (s *CreateWmEmbedTaskRequestCsvControl) SetEmbedPrecision(v int64) *CreateWmEmbedTaskRequestCsvControl {
	s.EmbedPrecision = &v
	return s
}

func (s *CreateWmEmbedTaskRequestCsvControl) SetEmbedTimePosition(v string) *CreateWmEmbedTaskRequestCsvControl {
	s.EmbedTimePosition = &v
	return s
}

func (s *CreateWmEmbedTaskRequestCsvControl) SetMethod(v string) *CreateWmEmbedTaskRequestCsvControl {
	s.Method = &v
	return s
}

func (s *CreateWmEmbedTaskRequestCsvControl) SetTimeFormat(v string) *CreateWmEmbedTaskRequestCsvControl {
	s.TimeFormat = &v
	return s
}

func (s *CreateWmEmbedTaskRequestCsvControl) Validate() error {
	return dara.Validate(s)
}

type CreateWmEmbedTaskRequestDocumentControl struct {
	BackgroundControl *CreateWmEmbedTaskRequestDocumentControlBackgroundControl `json:"BackgroundControl,omitempty" xml:"BackgroundControl,omitempty" type:"Struct"`
	// example:
	//
	// true
	InvisibleAntiAllCopy *bool `json:"InvisibleAntiAllCopy,omitempty" xml:"InvisibleAntiAllCopy,omitempty"`
	// example:
	//
	// true
	InvisibleAntiTextCopy *bool `json:"InvisibleAntiTextCopy,omitempty" xml:"InvisibleAntiTextCopy,omitempty"`
}

func (s CreateWmEmbedTaskRequestDocumentControl) String() string {
	return dara.Prettify(s)
}

func (s CreateWmEmbedTaskRequestDocumentControl) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskRequestDocumentControl) GetBackgroundControl() *CreateWmEmbedTaskRequestDocumentControlBackgroundControl {
	return s.BackgroundControl
}

func (s *CreateWmEmbedTaskRequestDocumentControl) GetInvisibleAntiAllCopy() *bool {
	return s.InvisibleAntiAllCopy
}

func (s *CreateWmEmbedTaskRequestDocumentControl) GetInvisibleAntiTextCopy() *bool {
	return s.InvisibleAntiTextCopy
}

func (s *CreateWmEmbedTaskRequestDocumentControl) SetBackgroundControl(v *CreateWmEmbedTaskRequestDocumentControlBackgroundControl) *CreateWmEmbedTaskRequestDocumentControl {
	s.BackgroundControl = v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControl) SetInvisibleAntiAllCopy(v bool) *CreateWmEmbedTaskRequestDocumentControl {
	s.InvisibleAntiAllCopy = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControl) SetInvisibleAntiTextCopy(v bool) *CreateWmEmbedTaskRequestDocumentControl {
	s.InvisibleAntiTextCopy = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControl) Validate() error {
	return dara.Validate(s)
}

type CreateWmEmbedTaskRequestDocumentControlBackgroundControl struct {
	// example:
	//
	// true
	BgAddInvisible *bool `json:"BgAddInvisible,omitempty" xml:"BgAddInvisible,omitempty"`
	// example:
	//
	// true
	BgAddVisible       *bool                                                                       `json:"BgAddVisible,omitempty" xml:"BgAddVisible,omitempty"`
	BgInvisibleControl *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgInvisibleControl `json:"BgInvisibleControl,omitempty" xml:"BgInvisibleControl,omitempty" type:"Struct"`
	BgVisibleControl   *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl   `json:"BgVisibleControl,omitempty" xml:"BgVisibleControl,omitempty" type:"Struct"`
}

func (s CreateWmEmbedTaskRequestDocumentControlBackgroundControl) String() string {
	return dara.Prettify(s)
}

func (s CreateWmEmbedTaskRequestDocumentControlBackgroundControl) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControl) GetBgAddInvisible() *bool {
	return s.BgAddInvisible
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControl) GetBgAddVisible() *bool {
	return s.BgAddVisible
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControl) GetBgInvisibleControl() *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgInvisibleControl {
	return s.BgInvisibleControl
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControl) GetBgVisibleControl() *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl {
	return s.BgVisibleControl
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControl) SetBgAddInvisible(v bool) *CreateWmEmbedTaskRequestDocumentControlBackgroundControl {
	s.BgAddInvisible = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControl) SetBgAddVisible(v bool) *CreateWmEmbedTaskRequestDocumentControlBackgroundControl {
	s.BgAddVisible = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControl) SetBgInvisibleControl(v *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgInvisibleControl) *CreateWmEmbedTaskRequestDocumentControlBackgroundControl {
	s.BgInvisibleControl = v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControl) SetBgVisibleControl(v *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) *CreateWmEmbedTaskRequestDocumentControlBackgroundControl {
	s.BgVisibleControl = v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControl) Validate() error {
	return dara.Validate(s)
}

type CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgInvisibleControl struct {
	// example:
	//
	// 10
	Opacity *int64 `json:"Opacity,omitempty" xml:"Opacity,omitempty"`
}

func (s CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgInvisibleControl) String() string {
	return dara.Prettify(s)
}

func (s CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgInvisibleControl) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgInvisibleControl) GetOpacity() *int64 {
	return s.Opacity
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgInvisibleControl) SetOpacity(v int64) *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgInvisibleControl {
	s.Opacity = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgInvisibleControl) Validate() error {
	return dara.Validate(s)
}

type CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl struct {
	// example:
	//
	// 30
	Angle *int64 `json:"Angle,omitempty" xml:"Angle,omitempty"`
	// example:
	//
	// 0x000000
	FontColor *string `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	// example:
	//
	// 30
	FontSize *int64 `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	// example:
	//
	// 3
	HorizontalNumber *int64 `json:"HorizontalNumber,omitempty" xml:"HorizontalNumber,omitempty"`
	// example:
	//
	// pos
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// example:
	//
	// 100
	Opacity *int64 `json:"Opacity,omitempty" xml:"Opacity,omitempty"`
	// example:
	//
	// 0.5
	PosX *string `json:"PosX,omitempty" xml:"PosX,omitempty"`
	// example:
	//
	// 0.5
	PosY *string `json:"PosY,omitempty" xml:"PosY,omitempty"`
	// example:
	//
	// 3
	VerticalNumber *int64 `json:"VerticalNumber,omitempty" xml:"VerticalNumber,omitempty"`
	// example:
	//
	// hello ****
	VisibleText *string `json:"VisibleText,omitempty" xml:"VisibleText,omitempty"`
}

func (s CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) String() string {
	return dara.Prettify(s)
}

func (s CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) GetAngle() *int64 {
	return s.Angle
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) GetFontColor() *string {
	return s.FontColor
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) GetFontSize() *int64 {
	return s.FontSize
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) GetHorizontalNumber() *int64 {
	return s.HorizontalNumber
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) GetMode() *string {
	return s.Mode
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) GetOpacity() *int64 {
	return s.Opacity
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) GetPosX() *string {
	return s.PosX
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) GetPosY() *string {
	return s.PosY
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) GetVerticalNumber() *int64 {
	return s.VerticalNumber
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) GetVisibleText() *string {
	return s.VisibleText
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) SetAngle(v int64) *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl {
	s.Angle = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) SetFontColor(v string) *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl {
	s.FontColor = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) SetFontSize(v int64) *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl {
	s.FontSize = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) SetHorizontalNumber(v int64) *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl {
	s.HorizontalNumber = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) SetMode(v string) *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl {
	s.Mode = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) SetOpacity(v int64) *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl {
	s.Opacity = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) SetPosX(v string) *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl {
	s.PosX = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) SetPosY(v string) *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl {
	s.PosY = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) SetVerticalNumber(v int64) *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl {
	s.VerticalNumber = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) SetVisibleText(v string) *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl {
	s.VisibleText = &v
	return s
}

func (s *CreateWmEmbedTaskRequestDocumentControlBackgroundControlBgVisibleControl) Validate() error {
	return dara.Validate(s)
}

type CreateWmEmbedTaskRequestImageControl struct {
	LogoVisibleControl *CreateWmEmbedTaskRequestImageControlLogoVisibleControl `json:"LogoVisibleControl,omitempty" xml:"LogoVisibleControl,omitempty" type:"Struct"`
	MetadataControl    *CreateWmEmbedTaskRequestImageControlMetadataControl    `json:"MetadataControl,omitempty" xml:"MetadataControl,omitempty" type:"Struct"`
	TextVisibleControl *CreateWmEmbedTaskRequestImageControlTextVisibleControl `json:"TextVisibleControl,omitempty" xml:"TextVisibleControl,omitempty" type:"Struct"`
}

func (s CreateWmEmbedTaskRequestImageControl) String() string {
	return dara.Prettify(s)
}

func (s CreateWmEmbedTaskRequestImageControl) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskRequestImageControl) GetLogoVisibleControl() *CreateWmEmbedTaskRequestImageControlLogoVisibleControl {
	return s.LogoVisibleControl
}

func (s *CreateWmEmbedTaskRequestImageControl) GetMetadataControl() *CreateWmEmbedTaskRequestImageControlMetadataControl {
	return s.MetadataControl
}

func (s *CreateWmEmbedTaskRequestImageControl) GetTextVisibleControl() *CreateWmEmbedTaskRequestImageControlTextVisibleControl {
	return s.TextVisibleControl
}

func (s *CreateWmEmbedTaskRequestImageControl) SetLogoVisibleControl(v *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) *CreateWmEmbedTaskRequestImageControl {
	s.LogoVisibleControl = v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControl) SetMetadataControl(v *CreateWmEmbedTaskRequestImageControlMetadataControl) *CreateWmEmbedTaskRequestImageControl {
	s.MetadataControl = v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControl) SetTextVisibleControl(v *CreateWmEmbedTaskRequestImageControlTextVisibleControl) *CreateWmEmbedTaskRequestImageControl {
	s.TextVisibleControl = v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControl) Validate() error {
	return dara.Validate(s)
}

type CreateWmEmbedTaskRequestImageControlLogoVisibleControl struct {
	Angle      *int64                                                        `json:"Angle,omitempty" xml:"Angle,omitempty"`
	LogoBase64 *string                                                       `json:"LogoBase64,omitempty" xml:"LogoBase64,omitempty"`
	Margin     *CreateWmEmbedTaskRequestImageControlLogoVisibleControlMargin `json:"Margin,omitempty" xml:"Margin,omitempty" type:"Struct"`
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

func (s CreateWmEmbedTaskRequestImageControlLogoVisibleControl) String() string {
	return dara.Prettify(s)
}

func (s CreateWmEmbedTaskRequestImageControlLogoVisibleControl) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) GetAngle() *int64 {
	return s.Angle
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) GetLogoBase64() *string {
	return s.LogoBase64
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) GetMargin() *CreateWmEmbedTaskRequestImageControlLogoVisibleControlMargin {
	return s.Margin
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) GetMode() *string {
	return s.Mode
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) GetOpacity() *int32 {
	return s.Opacity
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) GetPosAx() *float32 {
	return s.PosAx
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) GetPosAy() *float32 {
	return s.PosAy
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) GetPosX() *int64 {
	return s.PosX
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) GetPosY() *int64 {
	return s.PosY
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) GetSpaceX() *int64 {
	return s.SpaceX
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) GetSpaceY() *int64 {
	return s.SpaceY
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) GetVisible() *bool {
	return s.Visible
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) SetAngle(v int64) *CreateWmEmbedTaskRequestImageControlLogoVisibleControl {
	s.Angle = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) SetLogoBase64(v string) *CreateWmEmbedTaskRequestImageControlLogoVisibleControl {
	s.LogoBase64 = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) SetMargin(v *CreateWmEmbedTaskRequestImageControlLogoVisibleControlMargin) *CreateWmEmbedTaskRequestImageControlLogoVisibleControl {
	s.Margin = v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) SetMode(v string) *CreateWmEmbedTaskRequestImageControlLogoVisibleControl {
	s.Mode = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) SetOpacity(v int32) *CreateWmEmbedTaskRequestImageControlLogoVisibleControl {
	s.Opacity = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) SetPosAx(v float32) *CreateWmEmbedTaskRequestImageControlLogoVisibleControl {
	s.PosAx = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) SetPosAy(v float32) *CreateWmEmbedTaskRequestImageControlLogoVisibleControl {
	s.PosAy = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) SetPosX(v int64) *CreateWmEmbedTaskRequestImageControlLogoVisibleControl {
	s.PosX = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) SetPosY(v int64) *CreateWmEmbedTaskRequestImageControlLogoVisibleControl {
	s.PosY = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) SetSpaceX(v int64) *CreateWmEmbedTaskRequestImageControlLogoVisibleControl {
	s.SpaceX = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) SetSpaceY(v int64) *CreateWmEmbedTaskRequestImageControlLogoVisibleControl {
	s.SpaceY = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) SetVisible(v bool) *CreateWmEmbedTaskRequestImageControlLogoVisibleControl {
	s.Visible = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControl) Validate() error {
	return dara.Validate(s)
}

type CreateWmEmbedTaskRequestImageControlLogoVisibleControlMargin struct {
	Bottom *float32 `json:"Bottom,omitempty" xml:"Bottom,omitempty"`
	Left   *float32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Right  *float32 `json:"Right,omitempty" xml:"Right,omitempty"`
	Top    *float32 `json:"Top,omitempty" xml:"Top,omitempty"`
}

func (s CreateWmEmbedTaskRequestImageControlLogoVisibleControlMargin) String() string {
	return dara.Prettify(s)
}

func (s CreateWmEmbedTaskRequestImageControlLogoVisibleControlMargin) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControlMargin) GetBottom() *float32 {
	return s.Bottom
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControlMargin) GetLeft() *float32 {
	return s.Left
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControlMargin) GetRight() *float32 {
	return s.Right
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControlMargin) GetTop() *float32 {
	return s.Top
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControlMargin) SetBottom(v float32) *CreateWmEmbedTaskRequestImageControlLogoVisibleControlMargin {
	s.Bottom = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControlMargin) SetLeft(v float32) *CreateWmEmbedTaskRequestImageControlLogoVisibleControlMargin {
	s.Left = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControlMargin) SetRight(v float32) *CreateWmEmbedTaskRequestImageControlLogoVisibleControlMargin {
	s.Right = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControlMargin) SetTop(v float32) *CreateWmEmbedTaskRequestImageControlLogoVisibleControlMargin {
	s.Top = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlLogoVisibleControlMargin) Validate() error {
	return dara.Validate(s)
}

type CreateWmEmbedTaskRequestImageControlMetadataControl struct {
	Enable      *bool   `json:"Enable,omitempty" xml:"Enable,omitempty"`
	XmpKvBase64 *string `json:"XmpKvBase64,omitempty" xml:"XmpKvBase64,omitempty"`
}

func (s CreateWmEmbedTaskRequestImageControlMetadataControl) String() string {
	return dara.Prettify(s)
}

func (s CreateWmEmbedTaskRequestImageControlMetadataControl) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskRequestImageControlMetadataControl) GetEnable() *bool {
	return s.Enable
}

func (s *CreateWmEmbedTaskRequestImageControlMetadataControl) GetXmpKvBase64() *string {
	return s.XmpKvBase64
}

func (s *CreateWmEmbedTaskRequestImageControlMetadataControl) SetEnable(v bool) *CreateWmEmbedTaskRequestImageControlMetadataControl {
	s.Enable = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlMetadataControl) SetXmpKvBase64(v string) *CreateWmEmbedTaskRequestImageControlMetadataControl {
	s.XmpKvBase64 = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlMetadataControl) Validate() error {
	return dara.Validate(s)
}

type CreateWmEmbedTaskRequestImageControlTextVisibleControl struct {
	Angle       *int64                                                        `json:"Angle,omitempty" xml:"Angle,omitempty"`
	FontColor   *string                                                       `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	FontSize    *int64                                                        `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	Margin      *CreateWmEmbedTaskRequestImageControlTextVisibleControlMargin `json:"Margin,omitempty" xml:"Margin,omitempty" type:"Struct"`
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

func (s CreateWmEmbedTaskRequestImageControlTextVisibleControl) String() string {
	return dara.Prettify(s)
}

func (s CreateWmEmbedTaskRequestImageControlTextVisibleControl) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) GetAngle() *int64 {
	return s.Angle
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) GetFontColor() *string {
	return s.FontColor
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) GetFontSize() *int64 {
	return s.FontSize
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) GetMargin() *CreateWmEmbedTaskRequestImageControlTextVisibleControlMargin {
	return s.Margin
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) GetMode() *string {
	return s.Mode
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) GetOpacity() *int32 {
	return s.Opacity
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) GetPosAx() *float32 {
	return s.PosAx
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) GetPosAy() *float32 {
	return s.PosAy
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) GetPosX() *int64 {
	return s.PosX
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) GetPosY() *int64 {
	return s.PosY
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) GetSpaceX() *int64 {
	return s.SpaceX
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) GetSpaceY() *int64 {
	return s.SpaceY
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) GetVisible() *bool {
	return s.Visible
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) GetVisibleText() *string {
	return s.VisibleText
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) SetAngle(v int64) *CreateWmEmbedTaskRequestImageControlTextVisibleControl {
	s.Angle = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) SetFontColor(v string) *CreateWmEmbedTaskRequestImageControlTextVisibleControl {
	s.FontColor = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) SetFontSize(v int64) *CreateWmEmbedTaskRequestImageControlTextVisibleControl {
	s.FontSize = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) SetMargin(v *CreateWmEmbedTaskRequestImageControlTextVisibleControlMargin) *CreateWmEmbedTaskRequestImageControlTextVisibleControl {
	s.Margin = v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) SetMode(v string) *CreateWmEmbedTaskRequestImageControlTextVisibleControl {
	s.Mode = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) SetOpacity(v int32) *CreateWmEmbedTaskRequestImageControlTextVisibleControl {
	s.Opacity = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) SetPosAx(v float32) *CreateWmEmbedTaskRequestImageControlTextVisibleControl {
	s.PosAx = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) SetPosAy(v float32) *CreateWmEmbedTaskRequestImageControlTextVisibleControl {
	s.PosAy = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) SetPosX(v int64) *CreateWmEmbedTaskRequestImageControlTextVisibleControl {
	s.PosX = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) SetPosY(v int64) *CreateWmEmbedTaskRequestImageControlTextVisibleControl {
	s.PosY = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) SetSpaceX(v int64) *CreateWmEmbedTaskRequestImageControlTextVisibleControl {
	s.SpaceX = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) SetSpaceY(v int64) *CreateWmEmbedTaskRequestImageControlTextVisibleControl {
	s.SpaceY = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) SetVisible(v bool) *CreateWmEmbedTaskRequestImageControlTextVisibleControl {
	s.Visible = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) SetVisibleText(v string) *CreateWmEmbedTaskRequestImageControlTextVisibleControl {
	s.VisibleText = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControl) Validate() error {
	return dara.Validate(s)
}

type CreateWmEmbedTaskRequestImageControlTextVisibleControlMargin struct {
	Bottom *float32 `json:"Bottom,omitempty" xml:"Bottom,omitempty"`
	Left   *float32 `json:"Left,omitempty" xml:"Left,omitempty"`
	Right  *float32 `json:"Right,omitempty" xml:"Right,omitempty"`
	Top    *float32 `json:"Top,omitempty" xml:"Top,omitempty"`
}

func (s CreateWmEmbedTaskRequestImageControlTextVisibleControlMargin) String() string {
	return dara.Prettify(s)
}

func (s CreateWmEmbedTaskRequestImageControlTextVisibleControlMargin) GoString() string {
	return s.String()
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControlMargin) GetBottom() *float32 {
	return s.Bottom
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControlMargin) GetLeft() *float32 {
	return s.Left
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControlMargin) GetRight() *float32 {
	return s.Right
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControlMargin) GetTop() *float32 {
	return s.Top
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControlMargin) SetBottom(v float32) *CreateWmEmbedTaskRequestImageControlTextVisibleControlMargin {
	s.Bottom = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControlMargin) SetLeft(v float32) *CreateWmEmbedTaskRequestImageControlTextVisibleControlMargin {
	s.Left = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControlMargin) SetRight(v float32) *CreateWmEmbedTaskRequestImageControlTextVisibleControlMargin {
	s.Right = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControlMargin) SetTop(v float32) *CreateWmEmbedTaskRequestImageControlTextVisibleControlMargin {
	s.Top = &v
	return s
}

func (s *CreateWmEmbedTaskRequestImageControlTextVisibleControlMargin) Validate() error {
	return dara.Validate(s)
}
