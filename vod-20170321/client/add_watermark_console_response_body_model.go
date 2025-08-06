// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWatermarkConsoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddWatermarkConsoleResponseBody
	GetRequestId() *string
	SetWatermark(v *AddWatermarkConsoleResponseBodyWatermark) *AddWatermarkConsoleResponseBody
	GetWatermark() *AddWatermarkConsoleResponseBodyWatermark
	SetWatermarkInfo(v *AddWatermarkConsoleResponseBodyWatermarkInfo) *AddWatermarkConsoleResponseBody
	GetWatermarkInfo() *AddWatermarkConsoleResponseBodyWatermarkInfo
}

type AddWatermarkConsoleResponseBody struct {
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Watermark     *AddWatermarkConsoleResponseBodyWatermark     `json:"Watermark,omitempty" xml:"Watermark,omitempty" type:"Struct"`
	WatermarkInfo *AddWatermarkConsoleResponseBodyWatermarkInfo `json:"WatermarkInfo,omitempty" xml:"WatermarkInfo,omitempty" type:"Struct"`
}

func (s AddWatermarkConsoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddWatermarkConsoleResponseBody) GoString() string {
	return s.String()
}

func (s *AddWatermarkConsoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddWatermarkConsoleResponseBody) GetWatermark() *AddWatermarkConsoleResponseBodyWatermark {
	return s.Watermark
}

func (s *AddWatermarkConsoleResponseBody) GetWatermarkInfo() *AddWatermarkConsoleResponseBodyWatermarkInfo {
	return s.WatermarkInfo
}

func (s *AddWatermarkConsoleResponseBody) SetRequestId(v string) *AddWatermarkConsoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddWatermarkConsoleResponseBody) SetWatermark(v *AddWatermarkConsoleResponseBodyWatermark) *AddWatermarkConsoleResponseBody {
	s.Watermark = v
	return s
}

func (s *AddWatermarkConsoleResponseBody) SetWatermarkInfo(v *AddWatermarkConsoleResponseBodyWatermarkInfo) *AddWatermarkConsoleResponseBody {
	s.WatermarkInfo = v
	return s
}

func (s *AddWatermarkConsoleResponseBody) Validate() error {
	return dara.Validate(s)
}

type AddWatermarkConsoleResponseBodyWatermark struct {
	CreateTime       *int32  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Height           *string `json:"Height,omitempty" xml:"Height,omitempty"`
	HorizontalOffset *string `json:"HorizontalOffset,omitempty" xml:"HorizontalOffset,omitempty"`
	IsDefault        *string `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Position         *string `json:"Position,omitempty" xml:"Position,omitempty"`
	ScreenMode       *string `json:"ScreenMode,omitempty" xml:"ScreenMode,omitempty"`
	Url              *string `json:"Url,omitempty" xml:"Url,omitempty"`
	VerticalOffset   *string `json:"VerticalOffset,omitempty" xml:"VerticalOffset,omitempty"`
	VideoHeight      *int32  `json:"VideoHeight,omitempty" xml:"VideoHeight,omitempty"`
	VideoWidth       *int32  `json:"VideoWidth,omitempty" xml:"VideoWidth,omitempty"`
	WatermarkId      *string `json:"WatermarkId,omitempty" xml:"WatermarkId,omitempty"`
	Width            *string `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s AddWatermarkConsoleResponseBodyWatermark) String() string {
	return dara.Prettify(s)
}

func (s AddWatermarkConsoleResponseBodyWatermark) GoString() string {
	return s.String()
}

func (s *AddWatermarkConsoleResponseBodyWatermark) GetCreateTime() *int32 {
	return s.CreateTime
}

func (s *AddWatermarkConsoleResponseBodyWatermark) GetHeight() *string {
	return s.Height
}

func (s *AddWatermarkConsoleResponseBodyWatermark) GetHorizontalOffset() *string {
	return s.HorizontalOffset
}

func (s *AddWatermarkConsoleResponseBodyWatermark) GetIsDefault() *string {
	return s.IsDefault
}

func (s *AddWatermarkConsoleResponseBodyWatermark) GetName() *string {
	return s.Name
}

func (s *AddWatermarkConsoleResponseBodyWatermark) GetPosition() *string {
	return s.Position
}

func (s *AddWatermarkConsoleResponseBodyWatermark) GetScreenMode() *string {
	return s.ScreenMode
}

func (s *AddWatermarkConsoleResponseBodyWatermark) GetUrl() *string {
	return s.Url
}

func (s *AddWatermarkConsoleResponseBodyWatermark) GetVerticalOffset() *string {
	return s.VerticalOffset
}

func (s *AddWatermarkConsoleResponseBodyWatermark) GetVideoHeight() *int32 {
	return s.VideoHeight
}

func (s *AddWatermarkConsoleResponseBodyWatermark) GetVideoWidth() *int32 {
	return s.VideoWidth
}

func (s *AddWatermarkConsoleResponseBodyWatermark) GetWatermarkId() *string {
	return s.WatermarkId
}

func (s *AddWatermarkConsoleResponseBodyWatermark) GetWidth() *string {
	return s.Width
}

func (s *AddWatermarkConsoleResponseBodyWatermark) SetCreateTime(v int32) *AddWatermarkConsoleResponseBodyWatermark {
	s.CreateTime = &v
	return s
}

func (s *AddWatermarkConsoleResponseBodyWatermark) SetHeight(v string) *AddWatermarkConsoleResponseBodyWatermark {
	s.Height = &v
	return s
}

func (s *AddWatermarkConsoleResponseBodyWatermark) SetHorizontalOffset(v string) *AddWatermarkConsoleResponseBodyWatermark {
	s.HorizontalOffset = &v
	return s
}

func (s *AddWatermarkConsoleResponseBodyWatermark) SetIsDefault(v string) *AddWatermarkConsoleResponseBodyWatermark {
	s.IsDefault = &v
	return s
}

func (s *AddWatermarkConsoleResponseBodyWatermark) SetName(v string) *AddWatermarkConsoleResponseBodyWatermark {
	s.Name = &v
	return s
}

func (s *AddWatermarkConsoleResponseBodyWatermark) SetPosition(v string) *AddWatermarkConsoleResponseBodyWatermark {
	s.Position = &v
	return s
}

func (s *AddWatermarkConsoleResponseBodyWatermark) SetScreenMode(v string) *AddWatermarkConsoleResponseBodyWatermark {
	s.ScreenMode = &v
	return s
}

func (s *AddWatermarkConsoleResponseBodyWatermark) SetUrl(v string) *AddWatermarkConsoleResponseBodyWatermark {
	s.Url = &v
	return s
}

func (s *AddWatermarkConsoleResponseBodyWatermark) SetVerticalOffset(v string) *AddWatermarkConsoleResponseBodyWatermark {
	s.VerticalOffset = &v
	return s
}

func (s *AddWatermarkConsoleResponseBodyWatermark) SetVideoHeight(v int32) *AddWatermarkConsoleResponseBodyWatermark {
	s.VideoHeight = &v
	return s
}

func (s *AddWatermarkConsoleResponseBodyWatermark) SetVideoWidth(v int32) *AddWatermarkConsoleResponseBodyWatermark {
	s.VideoWidth = &v
	return s
}

func (s *AddWatermarkConsoleResponseBodyWatermark) SetWatermarkId(v string) *AddWatermarkConsoleResponseBodyWatermark {
	s.WatermarkId = &v
	return s
}

func (s *AddWatermarkConsoleResponseBodyWatermark) SetWidth(v string) *AddWatermarkConsoleResponseBodyWatermark {
	s.Width = &v
	return s
}

func (s *AddWatermarkConsoleResponseBodyWatermark) Validate() error {
	return dara.Validate(s)
}

type AddWatermarkConsoleResponseBodyWatermarkInfo struct {
	CreationTime    *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	FileUrl         *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	IsDefault       *string `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ScreenHeight    *int32  `json:"ScreenHeight,omitempty" xml:"ScreenHeight,omitempty"`
	ScreenWidth     *int32  `json:"ScreenWidth,omitempty" xml:"ScreenWidth,omitempty"`
	Type            *string `json:"Type,omitempty" xml:"Type,omitempty"`
	WatermarkConfig *string `json:"WatermarkConfig,omitempty" xml:"WatermarkConfig,omitempty"`
	WatermarkId     *string `json:"WatermarkId,omitempty" xml:"WatermarkId,omitempty"`
}

func (s AddWatermarkConsoleResponseBodyWatermarkInfo) String() string {
	return dara.Prettify(s)
}

func (s AddWatermarkConsoleResponseBodyWatermarkInfo) GoString() string {
	return s.String()
}

func (s *AddWatermarkConsoleResponseBodyWatermarkInfo) GetCreationTime() *string {
	return s.CreationTime
}

func (s *AddWatermarkConsoleResponseBodyWatermarkInfo) GetFileUrl() *string {
	return s.FileUrl
}

func (s *AddWatermarkConsoleResponseBodyWatermarkInfo) GetIsDefault() *string {
	return s.IsDefault
}

func (s *AddWatermarkConsoleResponseBodyWatermarkInfo) GetName() *string {
	return s.Name
}

func (s *AddWatermarkConsoleResponseBodyWatermarkInfo) GetScreenHeight() *int32 {
	return s.ScreenHeight
}

func (s *AddWatermarkConsoleResponseBodyWatermarkInfo) GetScreenWidth() *int32 {
	return s.ScreenWidth
}

func (s *AddWatermarkConsoleResponseBodyWatermarkInfo) GetType() *string {
	return s.Type
}

func (s *AddWatermarkConsoleResponseBodyWatermarkInfo) GetWatermarkConfig() *string {
	return s.WatermarkConfig
}

func (s *AddWatermarkConsoleResponseBodyWatermarkInfo) GetWatermarkId() *string {
	return s.WatermarkId
}

func (s *AddWatermarkConsoleResponseBodyWatermarkInfo) SetCreationTime(v string) *AddWatermarkConsoleResponseBodyWatermarkInfo {
	s.CreationTime = &v
	return s
}

func (s *AddWatermarkConsoleResponseBodyWatermarkInfo) SetFileUrl(v string) *AddWatermarkConsoleResponseBodyWatermarkInfo {
	s.FileUrl = &v
	return s
}

func (s *AddWatermarkConsoleResponseBodyWatermarkInfo) SetIsDefault(v string) *AddWatermarkConsoleResponseBodyWatermarkInfo {
	s.IsDefault = &v
	return s
}

func (s *AddWatermarkConsoleResponseBodyWatermarkInfo) SetName(v string) *AddWatermarkConsoleResponseBodyWatermarkInfo {
	s.Name = &v
	return s
}

func (s *AddWatermarkConsoleResponseBodyWatermarkInfo) SetScreenHeight(v int32) *AddWatermarkConsoleResponseBodyWatermarkInfo {
	s.ScreenHeight = &v
	return s
}

func (s *AddWatermarkConsoleResponseBodyWatermarkInfo) SetScreenWidth(v int32) *AddWatermarkConsoleResponseBodyWatermarkInfo {
	s.ScreenWidth = &v
	return s
}

func (s *AddWatermarkConsoleResponseBodyWatermarkInfo) SetType(v string) *AddWatermarkConsoleResponseBodyWatermarkInfo {
	s.Type = &v
	return s
}

func (s *AddWatermarkConsoleResponseBodyWatermarkInfo) SetWatermarkConfig(v string) *AddWatermarkConsoleResponseBodyWatermarkInfo {
	s.WatermarkConfig = &v
	return s
}

func (s *AddWatermarkConsoleResponseBodyWatermarkInfo) SetWatermarkId(v string) *AddWatermarkConsoleResponseBodyWatermarkInfo {
	s.WatermarkId = &v
	return s
}

func (s *AddWatermarkConsoleResponseBodyWatermarkInfo) Validate() error {
	return dara.Validate(s)
}
