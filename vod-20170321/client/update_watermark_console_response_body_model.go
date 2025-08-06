// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWatermarkConsoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateWatermarkConsoleResponseBody
	GetRequestId() *string
	SetWatermark(v *UpdateWatermarkConsoleResponseBodyWatermark) *UpdateWatermarkConsoleResponseBody
	GetWatermark() *UpdateWatermarkConsoleResponseBodyWatermark
	SetWatermarkInfo(v *UpdateWatermarkConsoleResponseBodyWatermarkInfo) *UpdateWatermarkConsoleResponseBody
	GetWatermarkInfo() *UpdateWatermarkConsoleResponseBodyWatermarkInfo
}

type UpdateWatermarkConsoleResponseBody struct {
	RequestId     *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Watermark     *UpdateWatermarkConsoleResponseBodyWatermark     `json:"Watermark,omitempty" xml:"Watermark,omitempty" type:"Struct"`
	WatermarkInfo *UpdateWatermarkConsoleResponseBodyWatermarkInfo `json:"WatermarkInfo,omitempty" xml:"WatermarkInfo,omitempty" type:"Struct"`
}

func (s UpdateWatermarkConsoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWatermarkConsoleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWatermarkConsoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWatermarkConsoleResponseBody) GetWatermark() *UpdateWatermarkConsoleResponseBodyWatermark {
	return s.Watermark
}

func (s *UpdateWatermarkConsoleResponseBody) GetWatermarkInfo() *UpdateWatermarkConsoleResponseBodyWatermarkInfo {
	return s.WatermarkInfo
}

func (s *UpdateWatermarkConsoleResponseBody) SetRequestId(v string) *UpdateWatermarkConsoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWatermarkConsoleResponseBody) SetWatermark(v *UpdateWatermarkConsoleResponseBodyWatermark) *UpdateWatermarkConsoleResponseBody {
	s.Watermark = v
	return s
}

func (s *UpdateWatermarkConsoleResponseBody) SetWatermarkInfo(v *UpdateWatermarkConsoleResponseBodyWatermarkInfo) *UpdateWatermarkConsoleResponseBody {
	s.WatermarkInfo = v
	return s
}

func (s *UpdateWatermarkConsoleResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateWatermarkConsoleResponseBodyWatermark struct {
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

func (s UpdateWatermarkConsoleResponseBodyWatermark) String() string {
	return dara.Prettify(s)
}

func (s UpdateWatermarkConsoleResponseBodyWatermark) GoString() string {
	return s.String()
}

func (s *UpdateWatermarkConsoleResponseBodyWatermark) GetCreateTime() *int32 {
	return s.CreateTime
}

func (s *UpdateWatermarkConsoleResponseBodyWatermark) GetHeight() *string {
	return s.Height
}

func (s *UpdateWatermarkConsoleResponseBodyWatermark) GetHorizontalOffset() *string {
	return s.HorizontalOffset
}

func (s *UpdateWatermarkConsoleResponseBodyWatermark) GetIsDefault() *string {
	return s.IsDefault
}

func (s *UpdateWatermarkConsoleResponseBodyWatermark) GetName() *string {
	return s.Name
}

func (s *UpdateWatermarkConsoleResponseBodyWatermark) GetPosition() *string {
	return s.Position
}

func (s *UpdateWatermarkConsoleResponseBodyWatermark) GetScreenMode() *string {
	return s.ScreenMode
}

func (s *UpdateWatermarkConsoleResponseBodyWatermark) GetUrl() *string {
	return s.Url
}

func (s *UpdateWatermarkConsoleResponseBodyWatermark) GetVerticalOffset() *string {
	return s.VerticalOffset
}

func (s *UpdateWatermarkConsoleResponseBodyWatermark) GetVideoHeight() *int32 {
	return s.VideoHeight
}

func (s *UpdateWatermarkConsoleResponseBodyWatermark) GetVideoWidth() *int32 {
	return s.VideoWidth
}

func (s *UpdateWatermarkConsoleResponseBodyWatermark) GetWatermarkId() *string {
	return s.WatermarkId
}

func (s *UpdateWatermarkConsoleResponseBodyWatermark) GetWidth() *string {
	return s.Width
}

func (s *UpdateWatermarkConsoleResponseBodyWatermark) SetCreateTime(v int32) *UpdateWatermarkConsoleResponseBodyWatermark {
	s.CreateTime = &v
	return s
}

func (s *UpdateWatermarkConsoleResponseBodyWatermark) SetHeight(v string) *UpdateWatermarkConsoleResponseBodyWatermark {
	s.Height = &v
	return s
}

func (s *UpdateWatermarkConsoleResponseBodyWatermark) SetHorizontalOffset(v string) *UpdateWatermarkConsoleResponseBodyWatermark {
	s.HorizontalOffset = &v
	return s
}

func (s *UpdateWatermarkConsoleResponseBodyWatermark) SetIsDefault(v string) *UpdateWatermarkConsoleResponseBodyWatermark {
	s.IsDefault = &v
	return s
}

func (s *UpdateWatermarkConsoleResponseBodyWatermark) SetName(v string) *UpdateWatermarkConsoleResponseBodyWatermark {
	s.Name = &v
	return s
}

func (s *UpdateWatermarkConsoleResponseBodyWatermark) SetPosition(v string) *UpdateWatermarkConsoleResponseBodyWatermark {
	s.Position = &v
	return s
}

func (s *UpdateWatermarkConsoleResponseBodyWatermark) SetScreenMode(v string) *UpdateWatermarkConsoleResponseBodyWatermark {
	s.ScreenMode = &v
	return s
}

func (s *UpdateWatermarkConsoleResponseBodyWatermark) SetUrl(v string) *UpdateWatermarkConsoleResponseBodyWatermark {
	s.Url = &v
	return s
}

func (s *UpdateWatermarkConsoleResponseBodyWatermark) SetVerticalOffset(v string) *UpdateWatermarkConsoleResponseBodyWatermark {
	s.VerticalOffset = &v
	return s
}

func (s *UpdateWatermarkConsoleResponseBodyWatermark) SetVideoHeight(v int32) *UpdateWatermarkConsoleResponseBodyWatermark {
	s.VideoHeight = &v
	return s
}

func (s *UpdateWatermarkConsoleResponseBodyWatermark) SetVideoWidth(v int32) *UpdateWatermarkConsoleResponseBodyWatermark {
	s.VideoWidth = &v
	return s
}

func (s *UpdateWatermarkConsoleResponseBodyWatermark) SetWatermarkId(v string) *UpdateWatermarkConsoleResponseBodyWatermark {
	s.WatermarkId = &v
	return s
}

func (s *UpdateWatermarkConsoleResponseBodyWatermark) SetWidth(v string) *UpdateWatermarkConsoleResponseBodyWatermark {
	s.Width = &v
	return s
}

func (s *UpdateWatermarkConsoleResponseBodyWatermark) Validate() error {
	return dara.Validate(s)
}

type UpdateWatermarkConsoleResponseBodyWatermarkInfo struct {
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

func (s UpdateWatermarkConsoleResponseBodyWatermarkInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdateWatermarkConsoleResponseBodyWatermarkInfo) GoString() string {
	return s.String()
}

func (s *UpdateWatermarkConsoleResponseBodyWatermarkInfo) GetCreationTime() *string {
	return s.CreationTime
}

func (s *UpdateWatermarkConsoleResponseBodyWatermarkInfo) GetFileUrl() *string {
	return s.FileUrl
}

func (s *UpdateWatermarkConsoleResponseBodyWatermarkInfo) GetIsDefault() *string {
	return s.IsDefault
}

func (s *UpdateWatermarkConsoleResponseBodyWatermarkInfo) GetName() *string {
	return s.Name
}

func (s *UpdateWatermarkConsoleResponseBodyWatermarkInfo) GetScreenHeight() *int32 {
	return s.ScreenHeight
}

func (s *UpdateWatermarkConsoleResponseBodyWatermarkInfo) GetScreenWidth() *int32 {
	return s.ScreenWidth
}

func (s *UpdateWatermarkConsoleResponseBodyWatermarkInfo) GetType() *string {
	return s.Type
}

func (s *UpdateWatermarkConsoleResponseBodyWatermarkInfo) GetWatermarkConfig() *string {
	return s.WatermarkConfig
}

func (s *UpdateWatermarkConsoleResponseBodyWatermarkInfo) GetWatermarkId() *string {
	return s.WatermarkId
}

func (s *UpdateWatermarkConsoleResponseBodyWatermarkInfo) SetCreationTime(v string) *UpdateWatermarkConsoleResponseBodyWatermarkInfo {
	s.CreationTime = &v
	return s
}

func (s *UpdateWatermarkConsoleResponseBodyWatermarkInfo) SetFileUrl(v string) *UpdateWatermarkConsoleResponseBodyWatermarkInfo {
	s.FileUrl = &v
	return s
}

func (s *UpdateWatermarkConsoleResponseBodyWatermarkInfo) SetIsDefault(v string) *UpdateWatermarkConsoleResponseBodyWatermarkInfo {
	s.IsDefault = &v
	return s
}

func (s *UpdateWatermarkConsoleResponseBodyWatermarkInfo) SetName(v string) *UpdateWatermarkConsoleResponseBodyWatermarkInfo {
	s.Name = &v
	return s
}

func (s *UpdateWatermarkConsoleResponseBodyWatermarkInfo) SetScreenHeight(v int32) *UpdateWatermarkConsoleResponseBodyWatermarkInfo {
	s.ScreenHeight = &v
	return s
}

func (s *UpdateWatermarkConsoleResponseBodyWatermarkInfo) SetScreenWidth(v int32) *UpdateWatermarkConsoleResponseBodyWatermarkInfo {
	s.ScreenWidth = &v
	return s
}

func (s *UpdateWatermarkConsoleResponseBodyWatermarkInfo) SetType(v string) *UpdateWatermarkConsoleResponseBodyWatermarkInfo {
	s.Type = &v
	return s
}

func (s *UpdateWatermarkConsoleResponseBodyWatermarkInfo) SetWatermarkConfig(v string) *UpdateWatermarkConsoleResponseBodyWatermarkInfo {
	s.WatermarkConfig = &v
	return s
}

func (s *UpdateWatermarkConsoleResponseBodyWatermarkInfo) SetWatermarkId(v string) *UpdateWatermarkConsoleResponseBodyWatermarkInfo {
	s.WatermarkId = &v
	return s
}

func (s *UpdateWatermarkConsoleResponseBodyWatermarkInfo) Validate() error {
	return dara.Validate(s)
}
