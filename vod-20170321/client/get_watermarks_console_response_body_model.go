// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWatermarksConsoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetWatermarksConsoleResponseBody
	GetRequestId() *string
	SetWatermarkInfos(v *GetWatermarksConsoleResponseBodyWatermarkInfos) *GetWatermarksConsoleResponseBody
	GetWatermarkInfos() *GetWatermarksConsoleResponseBodyWatermarkInfos
	SetWatermarks(v *GetWatermarksConsoleResponseBodyWatermarks) *GetWatermarksConsoleResponseBody
	GetWatermarks() *GetWatermarksConsoleResponseBodyWatermarks
}

type GetWatermarksConsoleResponseBody struct {
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	WatermarkInfos *GetWatermarksConsoleResponseBodyWatermarkInfos `json:"WatermarkInfos,omitempty" xml:"WatermarkInfos,omitempty" type:"Struct"`
	Watermarks     *GetWatermarksConsoleResponseBodyWatermarks     `json:"Watermarks,omitempty" xml:"Watermarks,omitempty" type:"Struct"`
}

func (s GetWatermarksConsoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWatermarksConsoleResponseBody) GoString() string {
	return s.String()
}

func (s *GetWatermarksConsoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWatermarksConsoleResponseBody) GetWatermarkInfos() *GetWatermarksConsoleResponseBodyWatermarkInfos {
	return s.WatermarkInfos
}

func (s *GetWatermarksConsoleResponseBody) GetWatermarks() *GetWatermarksConsoleResponseBodyWatermarks {
	return s.Watermarks
}

func (s *GetWatermarksConsoleResponseBody) SetRequestId(v string) *GetWatermarksConsoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWatermarksConsoleResponseBody) SetWatermarkInfos(v *GetWatermarksConsoleResponseBodyWatermarkInfos) *GetWatermarksConsoleResponseBody {
	s.WatermarkInfos = v
	return s
}

func (s *GetWatermarksConsoleResponseBody) SetWatermarks(v *GetWatermarksConsoleResponseBodyWatermarks) *GetWatermarksConsoleResponseBody {
	s.Watermarks = v
	return s
}

func (s *GetWatermarksConsoleResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetWatermarksConsoleResponseBodyWatermarkInfos struct {
	WatermarkInfo []*GetWatermarksConsoleResponseBodyWatermarkInfosWatermarkInfo `json:"WatermarkInfo,omitempty" xml:"WatermarkInfo,omitempty" type:"Repeated"`
}

func (s GetWatermarksConsoleResponseBodyWatermarkInfos) String() string {
	return dara.Prettify(s)
}

func (s GetWatermarksConsoleResponseBodyWatermarkInfos) GoString() string {
	return s.String()
}

func (s *GetWatermarksConsoleResponseBodyWatermarkInfos) GetWatermarkInfo() []*GetWatermarksConsoleResponseBodyWatermarkInfosWatermarkInfo {
	return s.WatermarkInfo
}

func (s *GetWatermarksConsoleResponseBodyWatermarkInfos) SetWatermarkInfo(v []*GetWatermarksConsoleResponseBodyWatermarkInfosWatermarkInfo) *GetWatermarksConsoleResponseBodyWatermarkInfos {
	s.WatermarkInfo = v
	return s
}

func (s *GetWatermarksConsoleResponseBodyWatermarkInfos) Validate() error {
	return dara.Validate(s)
}

type GetWatermarksConsoleResponseBodyWatermarkInfosWatermarkInfo struct {
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

func (s GetWatermarksConsoleResponseBodyWatermarkInfosWatermarkInfo) String() string {
	return dara.Prettify(s)
}

func (s GetWatermarksConsoleResponseBodyWatermarkInfosWatermarkInfo) GoString() string {
	return s.String()
}

func (s *GetWatermarksConsoleResponseBodyWatermarkInfosWatermarkInfo) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetWatermarksConsoleResponseBodyWatermarkInfosWatermarkInfo) GetFileUrl() *string {
	return s.FileUrl
}

func (s *GetWatermarksConsoleResponseBodyWatermarkInfosWatermarkInfo) GetIsDefault() *string {
	return s.IsDefault
}

func (s *GetWatermarksConsoleResponseBodyWatermarkInfosWatermarkInfo) GetName() *string {
	return s.Name
}

func (s *GetWatermarksConsoleResponseBodyWatermarkInfosWatermarkInfo) GetScreenHeight() *int32 {
	return s.ScreenHeight
}

func (s *GetWatermarksConsoleResponseBodyWatermarkInfosWatermarkInfo) GetScreenWidth() *int32 {
	return s.ScreenWidth
}

func (s *GetWatermarksConsoleResponseBodyWatermarkInfosWatermarkInfo) GetType() *string {
	return s.Type
}

func (s *GetWatermarksConsoleResponseBodyWatermarkInfosWatermarkInfo) GetWatermarkConfig() *string {
	return s.WatermarkConfig
}

func (s *GetWatermarksConsoleResponseBodyWatermarkInfosWatermarkInfo) GetWatermarkId() *string {
	return s.WatermarkId
}

func (s *GetWatermarksConsoleResponseBodyWatermarkInfosWatermarkInfo) SetCreationTime(v string) *GetWatermarksConsoleResponseBodyWatermarkInfosWatermarkInfo {
	s.CreationTime = &v
	return s
}

func (s *GetWatermarksConsoleResponseBodyWatermarkInfosWatermarkInfo) SetFileUrl(v string) *GetWatermarksConsoleResponseBodyWatermarkInfosWatermarkInfo {
	s.FileUrl = &v
	return s
}

func (s *GetWatermarksConsoleResponseBodyWatermarkInfosWatermarkInfo) SetIsDefault(v string) *GetWatermarksConsoleResponseBodyWatermarkInfosWatermarkInfo {
	s.IsDefault = &v
	return s
}

func (s *GetWatermarksConsoleResponseBodyWatermarkInfosWatermarkInfo) SetName(v string) *GetWatermarksConsoleResponseBodyWatermarkInfosWatermarkInfo {
	s.Name = &v
	return s
}

func (s *GetWatermarksConsoleResponseBodyWatermarkInfosWatermarkInfo) SetScreenHeight(v int32) *GetWatermarksConsoleResponseBodyWatermarkInfosWatermarkInfo {
	s.ScreenHeight = &v
	return s
}

func (s *GetWatermarksConsoleResponseBodyWatermarkInfosWatermarkInfo) SetScreenWidth(v int32) *GetWatermarksConsoleResponseBodyWatermarkInfosWatermarkInfo {
	s.ScreenWidth = &v
	return s
}

func (s *GetWatermarksConsoleResponseBodyWatermarkInfosWatermarkInfo) SetType(v string) *GetWatermarksConsoleResponseBodyWatermarkInfosWatermarkInfo {
	s.Type = &v
	return s
}

func (s *GetWatermarksConsoleResponseBodyWatermarkInfosWatermarkInfo) SetWatermarkConfig(v string) *GetWatermarksConsoleResponseBodyWatermarkInfosWatermarkInfo {
	s.WatermarkConfig = &v
	return s
}

func (s *GetWatermarksConsoleResponseBodyWatermarkInfosWatermarkInfo) SetWatermarkId(v string) *GetWatermarksConsoleResponseBodyWatermarkInfosWatermarkInfo {
	s.WatermarkId = &v
	return s
}

func (s *GetWatermarksConsoleResponseBodyWatermarkInfosWatermarkInfo) Validate() error {
	return dara.Validate(s)
}

type GetWatermarksConsoleResponseBodyWatermarks struct {
	Watermark []*GetWatermarksConsoleResponseBodyWatermarksWatermark `json:"Watermark,omitempty" xml:"Watermark,omitempty" type:"Repeated"`
}

func (s GetWatermarksConsoleResponseBodyWatermarks) String() string {
	return dara.Prettify(s)
}

func (s GetWatermarksConsoleResponseBodyWatermarks) GoString() string {
	return s.String()
}

func (s *GetWatermarksConsoleResponseBodyWatermarks) GetWatermark() []*GetWatermarksConsoleResponseBodyWatermarksWatermark {
	return s.Watermark
}

func (s *GetWatermarksConsoleResponseBodyWatermarks) SetWatermark(v []*GetWatermarksConsoleResponseBodyWatermarksWatermark) *GetWatermarksConsoleResponseBodyWatermarks {
	s.Watermark = v
	return s
}

func (s *GetWatermarksConsoleResponseBodyWatermarks) Validate() error {
	return dara.Validate(s)
}

type GetWatermarksConsoleResponseBodyWatermarksWatermark struct {
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

func (s GetWatermarksConsoleResponseBodyWatermarksWatermark) String() string {
	return dara.Prettify(s)
}

func (s GetWatermarksConsoleResponseBodyWatermarksWatermark) GoString() string {
	return s.String()
}

func (s *GetWatermarksConsoleResponseBodyWatermarksWatermark) GetActive() *string {
	return s.Active
}

func (s *GetWatermarksConsoleResponseBodyWatermarksWatermark) GetCreateTime() *int32 {
	return s.CreateTime
}

func (s *GetWatermarksConsoleResponseBodyWatermarksWatermark) GetHeight() *string {
	return s.Height
}

func (s *GetWatermarksConsoleResponseBodyWatermarksWatermark) GetHorizontalOffset() *string {
	return s.HorizontalOffset
}

func (s *GetWatermarksConsoleResponseBodyWatermarksWatermark) GetIsDefault() *string {
	return s.IsDefault
}

func (s *GetWatermarksConsoleResponseBodyWatermarksWatermark) GetName() *string {
	return s.Name
}

func (s *GetWatermarksConsoleResponseBodyWatermarksWatermark) GetPosition() *string {
	return s.Position
}

func (s *GetWatermarksConsoleResponseBodyWatermarksWatermark) GetScreenMode() *string {
	return s.ScreenMode
}

func (s *GetWatermarksConsoleResponseBodyWatermarksWatermark) GetUrl() *string {
	return s.Url
}

func (s *GetWatermarksConsoleResponseBodyWatermarksWatermark) GetVerticalOffset() *string {
	return s.VerticalOffset
}

func (s *GetWatermarksConsoleResponseBodyWatermarksWatermark) GetVideoHeight() *int32 {
	return s.VideoHeight
}

func (s *GetWatermarksConsoleResponseBodyWatermarksWatermark) GetVideoWidth() *int32 {
	return s.VideoWidth
}

func (s *GetWatermarksConsoleResponseBodyWatermarksWatermark) GetWatermarkId() *string {
	return s.WatermarkId
}

func (s *GetWatermarksConsoleResponseBodyWatermarksWatermark) GetWidth() *string {
	return s.Width
}

func (s *GetWatermarksConsoleResponseBodyWatermarksWatermark) SetActive(v string) *GetWatermarksConsoleResponseBodyWatermarksWatermark {
	s.Active = &v
	return s
}

func (s *GetWatermarksConsoleResponseBodyWatermarksWatermark) SetCreateTime(v int32) *GetWatermarksConsoleResponseBodyWatermarksWatermark {
	s.CreateTime = &v
	return s
}

func (s *GetWatermarksConsoleResponseBodyWatermarksWatermark) SetHeight(v string) *GetWatermarksConsoleResponseBodyWatermarksWatermark {
	s.Height = &v
	return s
}

func (s *GetWatermarksConsoleResponseBodyWatermarksWatermark) SetHorizontalOffset(v string) *GetWatermarksConsoleResponseBodyWatermarksWatermark {
	s.HorizontalOffset = &v
	return s
}

func (s *GetWatermarksConsoleResponseBodyWatermarksWatermark) SetIsDefault(v string) *GetWatermarksConsoleResponseBodyWatermarksWatermark {
	s.IsDefault = &v
	return s
}

func (s *GetWatermarksConsoleResponseBodyWatermarksWatermark) SetName(v string) *GetWatermarksConsoleResponseBodyWatermarksWatermark {
	s.Name = &v
	return s
}

func (s *GetWatermarksConsoleResponseBodyWatermarksWatermark) SetPosition(v string) *GetWatermarksConsoleResponseBodyWatermarksWatermark {
	s.Position = &v
	return s
}

func (s *GetWatermarksConsoleResponseBodyWatermarksWatermark) SetScreenMode(v string) *GetWatermarksConsoleResponseBodyWatermarksWatermark {
	s.ScreenMode = &v
	return s
}

func (s *GetWatermarksConsoleResponseBodyWatermarksWatermark) SetUrl(v string) *GetWatermarksConsoleResponseBodyWatermarksWatermark {
	s.Url = &v
	return s
}

func (s *GetWatermarksConsoleResponseBodyWatermarksWatermark) SetVerticalOffset(v string) *GetWatermarksConsoleResponseBodyWatermarksWatermark {
	s.VerticalOffset = &v
	return s
}

func (s *GetWatermarksConsoleResponseBodyWatermarksWatermark) SetVideoHeight(v int32) *GetWatermarksConsoleResponseBodyWatermarksWatermark {
	s.VideoHeight = &v
	return s
}

func (s *GetWatermarksConsoleResponseBodyWatermarksWatermark) SetVideoWidth(v int32) *GetWatermarksConsoleResponseBodyWatermarksWatermark {
	s.VideoWidth = &v
	return s
}

func (s *GetWatermarksConsoleResponseBodyWatermarksWatermark) SetWatermarkId(v string) *GetWatermarksConsoleResponseBodyWatermarksWatermark {
	s.WatermarkId = &v
	return s
}

func (s *GetWatermarksConsoleResponseBodyWatermarksWatermark) SetWidth(v string) *GetWatermarksConsoleResponseBodyWatermarksWatermark {
	s.Width = &v
	return s
}

func (s *GetWatermarksConsoleResponseBodyWatermarksWatermark) Validate() error {
	return dara.Validate(s)
}
