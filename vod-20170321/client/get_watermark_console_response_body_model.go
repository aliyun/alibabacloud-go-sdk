// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWatermarkConsoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetWatermarkConsoleResponseBody
	GetRequestId() *string
	SetWatermark(v *GetWatermarkConsoleResponseBodyWatermark) *GetWatermarkConsoleResponseBody
	GetWatermark() *GetWatermarkConsoleResponseBodyWatermark
	SetWatermarkInfo(v *GetWatermarkConsoleResponseBodyWatermarkInfo) *GetWatermarkConsoleResponseBody
	GetWatermarkInfo() *GetWatermarkConsoleResponseBodyWatermarkInfo
}

type GetWatermarkConsoleResponseBody struct {
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Watermark     *GetWatermarkConsoleResponseBodyWatermark     `json:"Watermark,omitempty" xml:"Watermark,omitempty" type:"Struct"`
	WatermarkInfo *GetWatermarkConsoleResponseBodyWatermarkInfo `json:"WatermarkInfo,omitempty" xml:"WatermarkInfo,omitempty" type:"Struct"`
}

func (s GetWatermarkConsoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWatermarkConsoleResponseBody) GoString() string {
	return s.String()
}

func (s *GetWatermarkConsoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWatermarkConsoleResponseBody) GetWatermark() *GetWatermarkConsoleResponseBodyWatermark {
	return s.Watermark
}

func (s *GetWatermarkConsoleResponseBody) GetWatermarkInfo() *GetWatermarkConsoleResponseBodyWatermarkInfo {
	return s.WatermarkInfo
}

func (s *GetWatermarkConsoleResponseBody) SetRequestId(v string) *GetWatermarkConsoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWatermarkConsoleResponseBody) SetWatermark(v *GetWatermarkConsoleResponseBodyWatermark) *GetWatermarkConsoleResponseBody {
	s.Watermark = v
	return s
}

func (s *GetWatermarkConsoleResponseBody) SetWatermarkInfo(v *GetWatermarkConsoleResponseBodyWatermarkInfo) *GetWatermarkConsoleResponseBody {
	s.WatermarkInfo = v
	return s
}

func (s *GetWatermarkConsoleResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetWatermarkConsoleResponseBodyWatermark struct {
	Active           *string `json:"Active,omitempty" xml:"Active,omitempty"`
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

func (s GetWatermarkConsoleResponseBodyWatermark) String() string {
	return dara.Prettify(s)
}

func (s GetWatermarkConsoleResponseBodyWatermark) GoString() string {
	return s.String()
}

func (s *GetWatermarkConsoleResponseBodyWatermark) GetActive() *string {
	return s.Active
}

func (s *GetWatermarkConsoleResponseBodyWatermark) GetCreateTime() *int32 {
	return s.CreateTime
}

func (s *GetWatermarkConsoleResponseBodyWatermark) GetHeight() *string {
	return s.Height
}

func (s *GetWatermarkConsoleResponseBodyWatermark) GetHorizontalOffset() *string {
	return s.HorizontalOffset
}

func (s *GetWatermarkConsoleResponseBodyWatermark) GetIsDefault() *string {
	return s.IsDefault
}

func (s *GetWatermarkConsoleResponseBodyWatermark) GetName() *string {
	return s.Name
}

func (s *GetWatermarkConsoleResponseBodyWatermark) GetPosition() *string {
	return s.Position
}

func (s *GetWatermarkConsoleResponseBodyWatermark) GetScreenMode() *string {
	return s.ScreenMode
}

func (s *GetWatermarkConsoleResponseBodyWatermark) GetUrl() *string {
	return s.Url
}

func (s *GetWatermarkConsoleResponseBodyWatermark) GetVerticalOffset() *string {
	return s.VerticalOffset
}

func (s *GetWatermarkConsoleResponseBodyWatermark) GetVideoHeight() *int32 {
	return s.VideoHeight
}

func (s *GetWatermarkConsoleResponseBodyWatermark) GetVideoWidth() *int32 {
	return s.VideoWidth
}

func (s *GetWatermarkConsoleResponseBodyWatermark) GetWatermarkId() *string {
	return s.WatermarkId
}

func (s *GetWatermarkConsoleResponseBodyWatermark) GetWidth() *string {
	return s.Width
}

func (s *GetWatermarkConsoleResponseBodyWatermark) SetActive(v string) *GetWatermarkConsoleResponseBodyWatermark {
	s.Active = &v
	return s
}

func (s *GetWatermarkConsoleResponseBodyWatermark) SetCreateTime(v int32) *GetWatermarkConsoleResponseBodyWatermark {
	s.CreateTime = &v
	return s
}

func (s *GetWatermarkConsoleResponseBodyWatermark) SetHeight(v string) *GetWatermarkConsoleResponseBodyWatermark {
	s.Height = &v
	return s
}

func (s *GetWatermarkConsoleResponseBodyWatermark) SetHorizontalOffset(v string) *GetWatermarkConsoleResponseBodyWatermark {
	s.HorizontalOffset = &v
	return s
}

func (s *GetWatermarkConsoleResponseBodyWatermark) SetIsDefault(v string) *GetWatermarkConsoleResponseBodyWatermark {
	s.IsDefault = &v
	return s
}

func (s *GetWatermarkConsoleResponseBodyWatermark) SetName(v string) *GetWatermarkConsoleResponseBodyWatermark {
	s.Name = &v
	return s
}

func (s *GetWatermarkConsoleResponseBodyWatermark) SetPosition(v string) *GetWatermarkConsoleResponseBodyWatermark {
	s.Position = &v
	return s
}

func (s *GetWatermarkConsoleResponseBodyWatermark) SetScreenMode(v string) *GetWatermarkConsoleResponseBodyWatermark {
	s.ScreenMode = &v
	return s
}

func (s *GetWatermarkConsoleResponseBodyWatermark) SetUrl(v string) *GetWatermarkConsoleResponseBodyWatermark {
	s.Url = &v
	return s
}

func (s *GetWatermarkConsoleResponseBodyWatermark) SetVerticalOffset(v string) *GetWatermarkConsoleResponseBodyWatermark {
	s.VerticalOffset = &v
	return s
}

func (s *GetWatermarkConsoleResponseBodyWatermark) SetVideoHeight(v int32) *GetWatermarkConsoleResponseBodyWatermark {
	s.VideoHeight = &v
	return s
}

func (s *GetWatermarkConsoleResponseBodyWatermark) SetVideoWidth(v int32) *GetWatermarkConsoleResponseBodyWatermark {
	s.VideoWidth = &v
	return s
}

func (s *GetWatermarkConsoleResponseBodyWatermark) SetWatermarkId(v string) *GetWatermarkConsoleResponseBodyWatermark {
	s.WatermarkId = &v
	return s
}

func (s *GetWatermarkConsoleResponseBodyWatermark) SetWidth(v string) *GetWatermarkConsoleResponseBodyWatermark {
	s.Width = &v
	return s
}

func (s *GetWatermarkConsoleResponseBodyWatermark) Validate() error {
	return dara.Validate(s)
}

type GetWatermarkConsoleResponseBodyWatermarkInfo struct {
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

func (s GetWatermarkConsoleResponseBodyWatermarkInfo) String() string {
	return dara.Prettify(s)
}

func (s GetWatermarkConsoleResponseBodyWatermarkInfo) GoString() string {
	return s.String()
}

func (s *GetWatermarkConsoleResponseBodyWatermarkInfo) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetWatermarkConsoleResponseBodyWatermarkInfo) GetFileUrl() *string {
	return s.FileUrl
}

func (s *GetWatermarkConsoleResponseBodyWatermarkInfo) GetIsDefault() *string {
	return s.IsDefault
}

func (s *GetWatermarkConsoleResponseBodyWatermarkInfo) GetName() *string {
	return s.Name
}

func (s *GetWatermarkConsoleResponseBodyWatermarkInfo) GetScreenHeight() *int32 {
	return s.ScreenHeight
}

func (s *GetWatermarkConsoleResponseBodyWatermarkInfo) GetScreenWidth() *int32 {
	return s.ScreenWidth
}

func (s *GetWatermarkConsoleResponseBodyWatermarkInfo) GetType() *string {
	return s.Type
}

func (s *GetWatermarkConsoleResponseBodyWatermarkInfo) GetWatermarkConfig() *string {
	return s.WatermarkConfig
}

func (s *GetWatermarkConsoleResponseBodyWatermarkInfo) GetWatermarkId() *string {
	return s.WatermarkId
}

func (s *GetWatermarkConsoleResponseBodyWatermarkInfo) SetCreationTime(v string) *GetWatermarkConsoleResponseBodyWatermarkInfo {
	s.CreationTime = &v
	return s
}

func (s *GetWatermarkConsoleResponseBodyWatermarkInfo) SetFileUrl(v string) *GetWatermarkConsoleResponseBodyWatermarkInfo {
	s.FileUrl = &v
	return s
}

func (s *GetWatermarkConsoleResponseBodyWatermarkInfo) SetIsDefault(v string) *GetWatermarkConsoleResponseBodyWatermarkInfo {
	s.IsDefault = &v
	return s
}

func (s *GetWatermarkConsoleResponseBodyWatermarkInfo) SetName(v string) *GetWatermarkConsoleResponseBodyWatermarkInfo {
	s.Name = &v
	return s
}

func (s *GetWatermarkConsoleResponseBodyWatermarkInfo) SetScreenHeight(v int32) *GetWatermarkConsoleResponseBodyWatermarkInfo {
	s.ScreenHeight = &v
	return s
}

func (s *GetWatermarkConsoleResponseBodyWatermarkInfo) SetScreenWidth(v int32) *GetWatermarkConsoleResponseBodyWatermarkInfo {
	s.ScreenWidth = &v
	return s
}

func (s *GetWatermarkConsoleResponseBodyWatermarkInfo) SetType(v string) *GetWatermarkConsoleResponseBodyWatermarkInfo {
	s.Type = &v
	return s
}

func (s *GetWatermarkConsoleResponseBodyWatermarkInfo) SetWatermarkConfig(v string) *GetWatermarkConsoleResponseBodyWatermarkInfo {
	s.WatermarkConfig = &v
	return s
}

func (s *GetWatermarkConsoleResponseBodyWatermarkInfo) SetWatermarkId(v string) *GetWatermarkConsoleResponseBodyWatermarkInfo {
	s.WatermarkId = &v
	return s
}

func (s *GetWatermarkConsoleResponseBodyWatermarkInfo) Validate() error {
	return dara.Validate(s)
}
