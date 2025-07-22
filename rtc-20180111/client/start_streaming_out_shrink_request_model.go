// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartStreamingOutShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnnotation(v string) *StartStreamingOutShrinkRequest
	GetAnnotation() *string
	SetAppId(v string) *StartStreamingOutShrinkRequest
	GetAppId() *string
	SetBackgrounds(v []*StartStreamingOutShrinkRequestBackgrounds) *StartStreamingOutShrinkRequest
	GetBackgrounds() []*StartStreamingOutShrinkRequestBackgrounds
	SetBgColor(v *StartStreamingOutShrinkRequestBgColor) *StartStreamingOutShrinkRequest
	GetBgColor() *StartStreamingOutShrinkRequestBgColor
	SetChannelId(v string) *StartStreamingOutShrinkRequest
	GetChannelId() *string
	SetClockWidgets(v []*StartStreamingOutShrinkRequestClockWidgets) *StartStreamingOutShrinkRequest
	GetClockWidgets() []*StartStreamingOutShrinkRequestClockWidgets
	SetCropMode(v int32) *StartStreamingOutShrinkRequest
	GetCropMode() *int32
	SetImages(v []*StartStreamingOutShrinkRequestImages) *StartStreamingOutShrinkRequest
	GetImages() []*StartStreamingOutShrinkRequestImages
	SetLayoutSpecifiedUsersShrink(v string) *StartStreamingOutShrinkRequest
	GetLayoutSpecifiedUsersShrink() *string
	SetPanes(v []*StartStreamingOutShrinkRequestPanes) *StartStreamingOutShrinkRequest
	GetPanes() []*StartStreamingOutShrinkRequestPanes
	SetRegionColor(v *StartStreamingOutShrinkRequestRegionColor) *StartStreamingOutShrinkRequest
	GetRegionColor() *StartStreamingOutShrinkRequestRegionColor
	SetReservePaneForNoCameraUser(v bool) *StartStreamingOutShrinkRequest
	GetReservePaneForNoCameraUser() *bool
	SetShowDefaultBackgroundOnMute(v bool) *StartStreamingOutShrinkRequest
	GetShowDefaultBackgroundOnMute() *bool
	SetSpecMixedUserList(v []*string) *StartStreamingOutShrinkRequest
	GetSpecMixedUserList() []*string
	SetStartWithoutChannel(v bool) *StartStreamingOutShrinkRequest
	GetStartWithoutChannel() *bool
	SetStartWithoutChannelWaitTime(v int32) *StartStreamingOutShrinkRequest
	GetStartWithoutChannelWaitTime() *int32
	SetSubHighResolutionStream(v bool) *StartStreamingOutShrinkRequest
	GetSubHighResolutionStream() *bool
	SetTaskId(v string) *StartStreamingOutShrinkRequest
	GetTaskId() *string
	SetTemplateId(v string) *StartStreamingOutShrinkRequest
	GetTemplateId() *string
	SetTexts(v []*StartStreamingOutShrinkRequestTexts) *StartStreamingOutShrinkRequest
	GetTexts() []*StartStreamingOutShrinkRequestTexts
	SetUrl(v string) *StartStreamingOutShrinkRequest
	GetUrl() *string
}

type StartStreamingOutShrinkRequest struct {
	// example:
	//
	// disable
	Annotation *string `json:"Annotation,omitempty" xml:"Annotation,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId       *string                                      `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Backgrounds []*StartStreamingOutShrinkRequestBackgrounds `json:"Backgrounds,omitempty" xml:"Backgrounds,omitempty" type:"Repeated"`
	BgColor     *StartStreamingOutShrinkRequestBgColor       `json:"BgColor,omitempty" xml:"BgColor,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId    *string                                       `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ClockWidgets []*StartStreamingOutShrinkRequestClockWidgets `json:"ClockWidgets,omitempty" xml:"ClockWidgets,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	CropMode                    *int32                                     `json:"CropMode,omitempty" xml:"CropMode,omitempty"`
	Images                      []*StartStreamingOutShrinkRequestImages    `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	LayoutSpecifiedUsersShrink  *string                                    `json:"LayoutSpecifiedUsers,omitempty" xml:"LayoutSpecifiedUsers,omitempty"`
	Panes                       []*StartStreamingOutShrinkRequestPanes     `json:"Panes,omitempty" xml:"Panes,omitempty" type:"Repeated"`
	RegionColor                 *StartStreamingOutShrinkRequestRegionColor `json:"RegionColor,omitempty" xml:"RegionColor,omitempty" type:"Struct"`
	ReservePaneForNoCameraUser  *bool                                      `json:"ReservePaneForNoCameraUser,omitempty" xml:"ReservePaneForNoCameraUser,omitempty"`
	ShowDefaultBackgroundOnMute *bool                                      `json:"ShowDefaultBackgroundOnMute,omitempty" xml:"ShowDefaultBackgroundOnMute,omitempty"`
	SpecMixedUserList           []*string                                  `json:"SpecMixedUserList,omitempty" xml:"SpecMixedUserList,omitempty" type:"Repeated"`
	StartWithoutChannel         *bool                                      `json:"StartWithoutChannel,omitempty" xml:"StartWithoutChannel,omitempty"`
	// example:
	//
	// 30
	StartWithoutChannelWaitTime *int32 `json:"StartWithoutChannelWaitTime,omitempty" xml:"StartWithoutChannelWaitTime,omitempty"`
	SubHighResolutionStream     *bool  `json:"SubHighResolutionStream,omitempty" xml:"SubHighResolutionStream,omitempty"`
	// example:
	//
	// 123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 567
	TemplateId *string                                `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Texts      []*StartStreamingOutShrinkRequestTexts `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// rtmp://example.com/live/stream
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s StartStreamingOutShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StartStreamingOutShrinkRequest) GoString() string {
	return s.String()
}

func (s *StartStreamingOutShrinkRequest) GetAnnotation() *string {
	return s.Annotation
}

func (s *StartStreamingOutShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *StartStreamingOutShrinkRequest) GetBackgrounds() []*StartStreamingOutShrinkRequestBackgrounds {
	return s.Backgrounds
}

func (s *StartStreamingOutShrinkRequest) GetBgColor() *StartStreamingOutShrinkRequestBgColor {
	return s.BgColor
}

func (s *StartStreamingOutShrinkRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *StartStreamingOutShrinkRequest) GetClockWidgets() []*StartStreamingOutShrinkRequestClockWidgets {
	return s.ClockWidgets
}

func (s *StartStreamingOutShrinkRequest) GetCropMode() *int32 {
	return s.CropMode
}

func (s *StartStreamingOutShrinkRequest) GetImages() []*StartStreamingOutShrinkRequestImages {
	return s.Images
}

func (s *StartStreamingOutShrinkRequest) GetLayoutSpecifiedUsersShrink() *string {
	return s.LayoutSpecifiedUsersShrink
}

func (s *StartStreamingOutShrinkRequest) GetPanes() []*StartStreamingOutShrinkRequestPanes {
	return s.Panes
}

func (s *StartStreamingOutShrinkRequest) GetRegionColor() *StartStreamingOutShrinkRequestRegionColor {
	return s.RegionColor
}

func (s *StartStreamingOutShrinkRequest) GetReservePaneForNoCameraUser() *bool {
	return s.ReservePaneForNoCameraUser
}

func (s *StartStreamingOutShrinkRequest) GetShowDefaultBackgroundOnMute() *bool {
	return s.ShowDefaultBackgroundOnMute
}

func (s *StartStreamingOutShrinkRequest) GetSpecMixedUserList() []*string {
	return s.SpecMixedUserList
}

func (s *StartStreamingOutShrinkRequest) GetStartWithoutChannel() *bool {
	return s.StartWithoutChannel
}

func (s *StartStreamingOutShrinkRequest) GetStartWithoutChannelWaitTime() *int32 {
	return s.StartWithoutChannelWaitTime
}

func (s *StartStreamingOutShrinkRequest) GetSubHighResolutionStream() *bool {
	return s.SubHighResolutionStream
}

func (s *StartStreamingOutShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *StartStreamingOutShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *StartStreamingOutShrinkRequest) GetTexts() []*StartStreamingOutShrinkRequestTexts {
	return s.Texts
}

func (s *StartStreamingOutShrinkRequest) GetUrl() *string {
	return s.Url
}

func (s *StartStreamingOutShrinkRequest) SetAnnotation(v string) *StartStreamingOutShrinkRequest {
	s.Annotation = &v
	return s
}

func (s *StartStreamingOutShrinkRequest) SetAppId(v string) *StartStreamingOutShrinkRequest {
	s.AppId = &v
	return s
}

func (s *StartStreamingOutShrinkRequest) SetBackgrounds(v []*StartStreamingOutShrinkRequestBackgrounds) *StartStreamingOutShrinkRequest {
	s.Backgrounds = v
	return s
}

func (s *StartStreamingOutShrinkRequest) SetBgColor(v *StartStreamingOutShrinkRequestBgColor) *StartStreamingOutShrinkRequest {
	s.BgColor = v
	return s
}

func (s *StartStreamingOutShrinkRequest) SetChannelId(v string) *StartStreamingOutShrinkRequest {
	s.ChannelId = &v
	return s
}

func (s *StartStreamingOutShrinkRequest) SetClockWidgets(v []*StartStreamingOutShrinkRequestClockWidgets) *StartStreamingOutShrinkRequest {
	s.ClockWidgets = v
	return s
}

func (s *StartStreamingOutShrinkRequest) SetCropMode(v int32) *StartStreamingOutShrinkRequest {
	s.CropMode = &v
	return s
}

func (s *StartStreamingOutShrinkRequest) SetImages(v []*StartStreamingOutShrinkRequestImages) *StartStreamingOutShrinkRequest {
	s.Images = v
	return s
}

func (s *StartStreamingOutShrinkRequest) SetLayoutSpecifiedUsersShrink(v string) *StartStreamingOutShrinkRequest {
	s.LayoutSpecifiedUsersShrink = &v
	return s
}

func (s *StartStreamingOutShrinkRequest) SetPanes(v []*StartStreamingOutShrinkRequestPanes) *StartStreamingOutShrinkRequest {
	s.Panes = v
	return s
}

func (s *StartStreamingOutShrinkRequest) SetRegionColor(v *StartStreamingOutShrinkRequestRegionColor) *StartStreamingOutShrinkRequest {
	s.RegionColor = v
	return s
}

func (s *StartStreamingOutShrinkRequest) SetReservePaneForNoCameraUser(v bool) *StartStreamingOutShrinkRequest {
	s.ReservePaneForNoCameraUser = &v
	return s
}

func (s *StartStreamingOutShrinkRequest) SetShowDefaultBackgroundOnMute(v bool) *StartStreamingOutShrinkRequest {
	s.ShowDefaultBackgroundOnMute = &v
	return s
}

func (s *StartStreamingOutShrinkRequest) SetSpecMixedUserList(v []*string) *StartStreamingOutShrinkRequest {
	s.SpecMixedUserList = v
	return s
}

func (s *StartStreamingOutShrinkRequest) SetStartWithoutChannel(v bool) *StartStreamingOutShrinkRequest {
	s.StartWithoutChannel = &v
	return s
}

func (s *StartStreamingOutShrinkRequest) SetStartWithoutChannelWaitTime(v int32) *StartStreamingOutShrinkRequest {
	s.StartWithoutChannelWaitTime = &v
	return s
}

func (s *StartStreamingOutShrinkRequest) SetSubHighResolutionStream(v bool) *StartStreamingOutShrinkRequest {
	s.SubHighResolutionStream = &v
	return s
}

func (s *StartStreamingOutShrinkRequest) SetTaskId(v string) *StartStreamingOutShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *StartStreamingOutShrinkRequest) SetTemplateId(v string) *StartStreamingOutShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *StartStreamingOutShrinkRequest) SetTexts(v []*StartStreamingOutShrinkRequestTexts) *StartStreamingOutShrinkRequest {
	s.Texts = v
	return s
}

func (s *StartStreamingOutShrinkRequest) SetUrl(v string) *StartStreamingOutShrinkRequest {
	s.Url = &v
	return s
}

func (s *StartStreamingOutShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type StartStreamingOutShrinkRequestBackgrounds struct {
	// example:
	//
	// 0.9
	Alpha *float64 `json:"Alpha,omitempty" xml:"Alpha,omitempty"`
	// example:
	//
	// 2
	BackgroundCropMode *int32 `json:"BackgroundCropMode,omitempty" xml:"BackgroundCropMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	Height *float64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 0
	Layer *int32 `json:"Layer,omitempty" xml:"Layer,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://aliyun.com/123.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	Width *float64 `json:"Width,omitempty" xml:"Width,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	X *float64 `json:"X,omitempty" xml:"X,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	Y *float64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s StartStreamingOutShrinkRequestBackgrounds) String() string {
	return dara.Prettify(s)
}

func (s StartStreamingOutShrinkRequestBackgrounds) GoString() string {
	return s.String()
}

func (s *StartStreamingOutShrinkRequestBackgrounds) GetAlpha() *float64 {
	return s.Alpha
}

func (s *StartStreamingOutShrinkRequestBackgrounds) GetBackgroundCropMode() *int32 {
	return s.BackgroundCropMode
}

func (s *StartStreamingOutShrinkRequestBackgrounds) GetHeight() *float64 {
	return s.Height
}

func (s *StartStreamingOutShrinkRequestBackgrounds) GetLayer() *int32 {
	return s.Layer
}

func (s *StartStreamingOutShrinkRequestBackgrounds) GetUrl() *string {
	return s.Url
}

func (s *StartStreamingOutShrinkRequestBackgrounds) GetWidth() *float64 {
	return s.Width
}

func (s *StartStreamingOutShrinkRequestBackgrounds) GetX() *float64 {
	return s.X
}

func (s *StartStreamingOutShrinkRequestBackgrounds) GetY() *float64 {
	return s.Y
}

func (s *StartStreamingOutShrinkRequestBackgrounds) SetAlpha(v float64) *StartStreamingOutShrinkRequestBackgrounds {
	s.Alpha = &v
	return s
}

func (s *StartStreamingOutShrinkRequestBackgrounds) SetBackgroundCropMode(v int32) *StartStreamingOutShrinkRequestBackgrounds {
	s.BackgroundCropMode = &v
	return s
}

func (s *StartStreamingOutShrinkRequestBackgrounds) SetHeight(v float64) *StartStreamingOutShrinkRequestBackgrounds {
	s.Height = &v
	return s
}

func (s *StartStreamingOutShrinkRequestBackgrounds) SetLayer(v int32) *StartStreamingOutShrinkRequestBackgrounds {
	s.Layer = &v
	return s
}

func (s *StartStreamingOutShrinkRequestBackgrounds) SetUrl(v string) *StartStreamingOutShrinkRequestBackgrounds {
	s.Url = &v
	return s
}

func (s *StartStreamingOutShrinkRequestBackgrounds) SetWidth(v float64) *StartStreamingOutShrinkRequestBackgrounds {
	s.Width = &v
	return s
}

func (s *StartStreamingOutShrinkRequestBackgrounds) SetX(v float64) *StartStreamingOutShrinkRequestBackgrounds {
	s.X = &v
	return s
}

func (s *StartStreamingOutShrinkRequestBackgrounds) SetY(v float64) *StartStreamingOutShrinkRequestBackgrounds {
	s.Y = &v
	return s
}

func (s *StartStreamingOutShrinkRequestBackgrounds) Validate() error {
	return dara.Validate(s)
}

type StartStreamingOutShrinkRequestBgColor struct {
	// example:
	//
	// 255
	B *int32 `json:"B,omitempty" xml:"B,omitempty"`
	// example:
	//
	// 255
	G *int32 `json:"G,omitempty" xml:"G,omitempty"`
	// example:
	//
	// 255
	R *int32 `json:"R,omitempty" xml:"R,omitempty"`
}

func (s StartStreamingOutShrinkRequestBgColor) String() string {
	return dara.Prettify(s)
}

func (s StartStreamingOutShrinkRequestBgColor) GoString() string {
	return s.String()
}

func (s *StartStreamingOutShrinkRequestBgColor) GetB() *int32 {
	return s.B
}

func (s *StartStreamingOutShrinkRequestBgColor) GetG() *int32 {
	return s.G
}

func (s *StartStreamingOutShrinkRequestBgColor) GetR() *int32 {
	return s.R
}

func (s *StartStreamingOutShrinkRequestBgColor) SetB(v int32) *StartStreamingOutShrinkRequestBgColor {
	s.B = &v
	return s
}

func (s *StartStreamingOutShrinkRequestBgColor) SetG(v int32) *StartStreamingOutShrinkRequestBgColor {
	s.G = &v
	return s
}

func (s *StartStreamingOutShrinkRequestBgColor) SetR(v int32) *StartStreamingOutShrinkRequestBgColor {
	s.R = &v
	return s
}

func (s *StartStreamingOutShrinkRequestBgColor) Validate() error {
	return dara.Validate(s)
}

type StartStreamingOutShrinkRequestClockWidgets struct {
	// example:
	//
	// 0.9
	Alpha *float64 `json:"Alpha,omitempty" xml:"Alpha,omitempty"`
	// example:
	//
	// 0.6
	BoxAlpha *float64 `json:"BoxAlpha,omitempty" xml:"BoxAlpha,omitempty"`
	// example:
	//
	// 5
	BoxBorderw *int32                                              `json:"BoxBorderw,omitempty" xml:"BoxBorderw,omitempty"`
	BoxColor   *StartStreamingOutShrinkRequestClockWidgetsBoxColor `json:"BoxColor,omitempty" xml:"BoxColor,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Font      *int32                                               `json:"Font,omitempty" xml:"Font,omitempty"`
	FontColor *StartStreamingOutShrinkRequestClockWidgetsFontColor `json:"FontColor,omitempty" xml:"FontColor,omitempty" type:"Struct"`
	// example:
	//
	// 30
	FontSize *int32 `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	HasBox   *bool  `json:"HasBox,omitempty" xml:"HasBox,omitempty"`
	// example:
	//
	// 0
	Layer *int32 `json:"Layer,omitempty" xml:"Layer,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	X *float64 `json:"X,omitempty" xml:"X,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	Y *float64 `json:"Y,omitempty" xml:"Y,omitempty"`
	// example:
	//
	// 8
	Zone *int32 `json:"Zone,omitempty" xml:"Zone,omitempty"`
}

func (s StartStreamingOutShrinkRequestClockWidgets) String() string {
	return dara.Prettify(s)
}

func (s StartStreamingOutShrinkRequestClockWidgets) GoString() string {
	return s.String()
}

func (s *StartStreamingOutShrinkRequestClockWidgets) GetAlpha() *float64 {
	return s.Alpha
}

func (s *StartStreamingOutShrinkRequestClockWidgets) GetBoxAlpha() *float64 {
	return s.BoxAlpha
}

func (s *StartStreamingOutShrinkRequestClockWidgets) GetBoxBorderw() *int32 {
	return s.BoxBorderw
}

func (s *StartStreamingOutShrinkRequestClockWidgets) GetBoxColor() *StartStreamingOutShrinkRequestClockWidgetsBoxColor {
	return s.BoxColor
}

func (s *StartStreamingOutShrinkRequestClockWidgets) GetFont() *int32 {
	return s.Font
}

func (s *StartStreamingOutShrinkRequestClockWidgets) GetFontColor() *StartStreamingOutShrinkRequestClockWidgetsFontColor {
	return s.FontColor
}

func (s *StartStreamingOutShrinkRequestClockWidgets) GetFontSize() *int32 {
	return s.FontSize
}

func (s *StartStreamingOutShrinkRequestClockWidgets) GetHasBox() *bool {
	return s.HasBox
}

func (s *StartStreamingOutShrinkRequestClockWidgets) GetLayer() *int32 {
	return s.Layer
}

func (s *StartStreamingOutShrinkRequestClockWidgets) GetX() *float64 {
	return s.X
}

func (s *StartStreamingOutShrinkRequestClockWidgets) GetY() *float64 {
	return s.Y
}

func (s *StartStreamingOutShrinkRequestClockWidgets) GetZone() *int32 {
	return s.Zone
}

func (s *StartStreamingOutShrinkRequestClockWidgets) SetAlpha(v float64) *StartStreamingOutShrinkRequestClockWidgets {
	s.Alpha = &v
	return s
}

func (s *StartStreamingOutShrinkRequestClockWidgets) SetBoxAlpha(v float64) *StartStreamingOutShrinkRequestClockWidgets {
	s.BoxAlpha = &v
	return s
}

func (s *StartStreamingOutShrinkRequestClockWidgets) SetBoxBorderw(v int32) *StartStreamingOutShrinkRequestClockWidgets {
	s.BoxBorderw = &v
	return s
}

func (s *StartStreamingOutShrinkRequestClockWidgets) SetBoxColor(v *StartStreamingOutShrinkRequestClockWidgetsBoxColor) *StartStreamingOutShrinkRequestClockWidgets {
	s.BoxColor = v
	return s
}

func (s *StartStreamingOutShrinkRequestClockWidgets) SetFont(v int32) *StartStreamingOutShrinkRequestClockWidgets {
	s.Font = &v
	return s
}

func (s *StartStreamingOutShrinkRequestClockWidgets) SetFontColor(v *StartStreamingOutShrinkRequestClockWidgetsFontColor) *StartStreamingOutShrinkRequestClockWidgets {
	s.FontColor = v
	return s
}

func (s *StartStreamingOutShrinkRequestClockWidgets) SetFontSize(v int32) *StartStreamingOutShrinkRequestClockWidgets {
	s.FontSize = &v
	return s
}

func (s *StartStreamingOutShrinkRequestClockWidgets) SetHasBox(v bool) *StartStreamingOutShrinkRequestClockWidgets {
	s.HasBox = &v
	return s
}

func (s *StartStreamingOutShrinkRequestClockWidgets) SetLayer(v int32) *StartStreamingOutShrinkRequestClockWidgets {
	s.Layer = &v
	return s
}

func (s *StartStreamingOutShrinkRequestClockWidgets) SetX(v float64) *StartStreamingOutShrinkRequestClockWidgets {
	s.X = &v
	return s
}

func (s *StartStreamingOutShrinkRequestClockWidgets) SetY(v float64) *StartStreamingOutShrinkRequestClockWidgets {
	s.Y = &v
	return s
}

func (s *StartStreamingOutShrinkRequestClockWidgets) SetZone(v int32) *StartStreamingOutShrinkRequestClockWidgets {
	s.Zone = &v
	return s
}

func (s *StartStreamingOutShrinkRequestClockWidgets) Validate() error {
	return dara.Validate(s)
}

type StartStreamingOutShrinkRequestClockWidgetsBoxColor struct {
	// example:
	//
	// 255
	B *int32 `json:"B,omitempty" xml:"B,omitempty"`
	// example:
	//
	// 255
	G *int32 `json:"G,omitempty" xml:"G,omitempty"`
	// example:
	//
	// 255
	R *int32 `json:"R,omitempty" xml:"R,omitempty"`
}

func (s StartStreamingOutShrinkRequestClockWidgetsBoxColor) String() string {
	return dara.Prettify(s)
}

func (s StartStreamingOutShrinkRequestClockWidgetsBoxColor) GoString() string {
	return s.String()
}

func (s *StartStreamingOutShrinkRequestClockWidgetsBoxColor) GetB() *int32 {
	return s.B
}

func (s *StartStreamingOutShrinkRequestClockWidgetsBoxColor) GetG() *int32 {
	return s.G
}

func (s *StartStreamingOutShrinkRequestClockWidgetsBoxColor) GetR() *int32 {
	return s.R
}

func (s *StartStreamingOutShrinkRequestClockWidgetsBoxColor) SetB(v int32) *StartStreamingOutShrinkRequestClockWidgetsBoxColor {
	s.B = &v
	return s
}

func (s *StartStreamingOutShrinkRequestClockWidgetsBoxColor) SetG(v int32) *StartStreamingOutShrinkRequestClockWidgetsBoxColor {
	s.G = &v
	return s
}

func (s *StartStreamingOutShrinkRequestClockWidgetsBoxColor) SetR(v int32) *StartStreamingOutShrinkRequestClockWidgetsBoxColor {
	s.R = &v
	return s
}

func (s *StartStreamingOutShrinkRequestClockWidgetsBoxColor) Validate() error {
	return dara.Validate(s)
}

type StartStreamingOutShrinkRequestClockWidgetsFontColor struct {
	// example:
	//
	// 255
	B *int32 `json:"B,omitempty" xml:"B,omitempty"`
	// example:
	//
	// 255
	G *int32 `json:"G,omitempty" xml:"G,omitempty"`
	// example:
	//
	// 255
	R *int32 `json:"R,omitempty" xml:"R,omitempty"`
}

func (s StartStreamingOutShrinkRequestClockWidgetsFontColor) String() string {
	return dara.Prettify(s)
}

func (s StartStreamingOutShrinkRequestClockWidgetsFontColor) GoString() string {
	return s.String()
}

func (s *StartStreamingOutShrinkRequestClockWidgetsFontColor) GetB() *int32 {
	return s.B
}

func (s *StartStreamingOutShrinkRequestClockWidgetsFontColor) GetG() *int32 {
	return s.G
}

func (s *StartStreamingOutShrinkRequestClockWidgetsFontColor) GetR() *int32 {
	return s.R
}

func (s *StartStreamingOutShrinkRequestClockWidgetsFontColor) SetB(v int32) *StartStreamingOutShrinkRequestClockWidgetsFontColor {
	s.B = &v
	return s
}

func (s *StartStreamingOutShrinkRequestClockWidgetsFontColor) SetG(v int32) *StartStreamingOutShrinkRequestClockWidgetsFontColor {
	s.G = &v
	return s
}

func (s *StartStreamingOutShrinkRequestClockWidgetsFontColor) SetR(v int32) *StartStreamingOutShrinkRequestClockWidgetsFontColor {
	s.R = &v
	return s
}

func (s *StartStreamingOutShrinkRequestClockWidgetsFontColor) Validate() error {
	return dara.Validate(s)
}

type StartStreamingOutShrinkRequestImages struct {
	// example:
	//
	// 0.9
	Alpha *float64 `json:"Alpha,omitempty" xml:"Alpha,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	Height *float64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 2
	ImageCropMode *int32 `json:"ImageCropMode,omitempty" xml:"ImageCropMode,omitempty"`
	// example:
	//
	// 0
	Layer *int32 `json:"Layer,omitempty" xml:"Layer,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://aliyun.com/123xxx.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	Width *float64 `json:"Width,omitempty" xml:"Width,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	X *float64 `json:"X,omitempty" xml:"X,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	Y *float64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s StartStreamingOutShrinkRequestImages) String() string {
	return dara.Prettify(s)
}

func (s StartStreamingOutShrinkRequestImages) GoString() string {
	return s.String()
}

func (s *StartStreamingOutShrinkRequestImages) GetAlpha() *float64 {
	return s.Alpha
}

func (s *StartStreamingOutShrinkRequestImages) GetHeight() *float64 {
	return s.Height
}

func (s *StartStreamingOutShrinkRequestImages) GetImageCropMode() *int32 {
	return s.ImageCropMode
}

func (s *StartStreamingOutShrinkRequestImages) GetLayer() *int32 {
	return s.Layer
}

func (s *StartStreamingOutShrinkRequestImages) GetUrl() *string {
	return s.Url
}

func (s *StartStreamingOutShrinkRequestImages) GetWidth() *float64 {
	return s.Width
}

func (s *StartStreamingOutShrinkRequestImages) GetX() *float64 {
	return s.X
}

func (s *StartStreamingOutShrinkRequestImages) GetY() *float64 {
	return s.Y
}

func (s *StartStreamingOutShrinkRequestImages) SetAlpha(v float64) *StartStreamingOutShrinkRequestImages {
	s.Alpha = &v
	return s
}

func (s *StartStreamingOutShrinkRequestImages) SetHeight(v float64) *StartStreamingOutShrinkRequestImages {
	s.Height = &v
	return s
}

func (s *StartStreamingOutShrinkRequestImages) SetImageCropMode(v int32) *StartStreamingOutShrinkRequestImages {
	s.ImageCropMode = &v
	return s
}

func (s *StartStreamingOutShrinkRequestImages) SetLayer(v int32) *StartStreamingOutShrinkRequestImages {
	s.Layer = &v
	return s
}

func (s *StartStreamingOutShrinkRequestImages) SetUrl(v string) *StartStreamingOutShrinkRequestImages {
	s.Url = &v
	return s
}

func (s *StartStreamingOutShrinkRequestImages) SetWidth(v float64) *StartStreamingOutShrinkRequestImages {
	s.Width = &v
	return s
}

func (s *StartStreamingOutShrinkRequestImages) SetX(v float64) *StartStreamingOutShrinkRequestImages {
	s.X = &v
	return s
}

func (s *StartStreamingOutShrinkRequestImages) SetY(v float64) *StartStreamingOutShrinkRequestImages {
	s.Y = &v
	return s
}

func (s *StartStreamingOutShrinkRequestImages) Validate() error {
	return dara.Validate(s)
}

type StartStreamingOutShrinkRequestPanes struct {
	Backgrounds []*StartStreamingOutShrinkRequestPanesBackgrounds `json:"Backgrounds,omitempty" xml:"Backgrounds,omitempty" type:"Repeated"`
	Images      []*StartStreamingOutShrinkRequestPanesImages      `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	PaneCropMode *int32 `json:"PaneCropMode,omitempty" xml:"PaneCropMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	PaneId                    *string `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
	ReservePaneForOfflineUser *bool   `json:"ReservePaneForOfflineUser,omitempty" xml:"ReservePaneForOfflineUser,omitempty"`
	// example:
	//
	// 1811****
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// Video
	SourceType *string                                     `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Texts      []*StartStreamingOutShrinkRequestPanesTexts `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
	// example:
	//
	// cameraFirst
	VideoOrder *string                                        `json:"VideoOrder,omitempty" xml:"VideoOrder,omitempty"`
	Whiteboard *StartStreamingOutShrinkRequestPanesWhiteboard `json:"Whiteboard,omitempty" xml:"Whiteboard,omitempty" type:"Struct"`
}

func (s StartStreamingOutShrinkRequestPanes) String() string {
	return dara.Prettify(s)
}

func (s StartStreamingOutShrinkRequestPanes) GoString() string {
	return s.String()
}

func (s *StartStreamingOutShrinkRequestPanes) GetBackgrounds() []*StartStreamingOutShrinkRequestPanesBackgrounds {
	return s.Backgrounds
}

func (s *StartStreamingOutShrinkRequestPanes) GetImages() []*StartStreamingOutShrinkRequestPanesImages {
	return s.Images
}

func (s *StartStreamingOutShrinkRequestPanes) GetPaneCropMode() *int32 {
	return s.PaneCropMode
}

func (s *StartStreamingOutShrinkRequestPanes) GetPaneId() *string {
	return s.PaneId
}

func (s *StartStreamingOutShrinkRequestPanes) GetReservePaneForOfflineUser() *bool {
	return s.ReservePaneForOfflineUser
}

func (s *StartStreamingOutShrinkRequestPanes) GetSource() *string {
	return s.Source
}

func (s *StartStreamingOutShrinkRequestPanes) GetSourceType() *string {
	return s.SourceType
}

func (s *StartStreamingOutShrinkRequestPanes) GetTexts() []*StartStreamingOutShrinkRequestPanesTexts {
	return s.Texts
}

func (s *StartStreamingOutShrinkRequestPanes) GetVideoOrder() *string {
	return s.VideoOrder
}

func (s *StartStreamingOutShrinkRequestPanes) GetWhiteboard() *StartStreamingOutShrinkRequestPanesWhiteboard {
	return s.Whiteboard
}

func (s *StartStreamingOutShrinkRequestPanes) SetBackgrounds(v []*StartStreamingOutShrinkRequestPanesBackgrounds) *StartStreamingOutShrinkRequestPanes {
	s.Backgrounds = v
	return s
}

func (s *StartStreamingOutShrinkRequestPanes) SetImages(v []*StartStreamingOutShrinkRequestPanesImages) *StartStreamingOutShrinkRequestPanes {
	s.Images = v
	return s
}

func (s *StartStreamingOutShrinkRequestPanes) SetPaneCropMode(v int32) *StartStreamingOutShrinkRequestPanes {
	s.PaneCropMode = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanes) SetPaneId(v string) *StartStreamingOutShrinkRequestPanes {
	s.PaneId = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanes) SetReservePaneForOfflineUser(v bool) *StartStreamingOutShrinkRequestPanes {
	s.ReservePaneForOfflineUser = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanes) SetSource(v string) *StartStreamingOutShrinkRequestPanes {
	s.Source = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanes) SetSourceType(v string) *StartStreamingOutShrinkRequestPanes {
	s.SourceType = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanes) SetTexts(v []*StartStreamingOutShrinkRequestPanesTexts) *StartStreamingOutShrinkRequestPanes {
	s.Texts = v
	return s
}

func (s *StartStreamingOutShrinkRequestPanes) SetVideoOrder(v string) *StartStreamingOutShrinkRequestPanes {
	s.VideoOrder = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanes) SetWhiteboard(v *StartStreamingOutShrinkRequestPanesWhiteboard) *StartStreamingOutShrinkRequestPanes {
	s.Whiteboard = v
	return s
}

func (s *StartStreamingOutShrinkRequestPanes) Validate() error {
	return dara.Validate(s)
}

type StartStreamingOutShrinkRequestPanesBackgrounds struct {
	// example:
	//
	// 0.9
	Alpha *float64 `json:"Alpha,omitempty" xml:"Alpha,omitempty"`
	// example:
	//
	// backup
	Display *string `json:"Display,omitempty" xml:"Display,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	Height *float64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 0
	Layer *int32 `json:"Layer,omitempty" xml:"Layer,omitempty"`
	// example:
	//
	// 2
	PaneBackgroundCropMode *int32 `json:"PaneBackgroundCropMode,omitempty" xml:"PaneBackgroundCropMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://aliyun.com/123xx.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	Width *float64 `json:"Width,omitempty" xml:"Width,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	X *float64 `json:"X,omitempty" xml:"X,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	Y *float64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s StartStreamingOutShrinkRequestPanesBackgrounds) String() string {
	return dara.Prettify(s)
}

func (s StartStreamingOutShrinkRequestPanesBackgrounds) GoString() string {
	return s.String()
}

func (s *StartStreamingOutShrinkRequestPanesBackgrounds) GetAlpha() *float64 {
	return s.Alpha
}

func (s *StartStreamingOutShrinkRequestPanesBackgrounds) GetDisplay() *string {
	return s.Display
}

func (s *StartStreamingOutShrinkRequestPanesBackgrounds) GetHeight() *float64 {
	return s.Height
}

func (s *StartStreamingOutShrinkRequestPanesBackgrounds) GetLayer() *int32 {
	return s.Layer
}

func (s *StartStreamingOutShrinkRequestPanesBackgrounds) GetPaneBackgroundCropMode() *int32 {
	return s.PaneBackgroundCropMode
}

func (s *StartStreamingOutShrinkRequestPanesBackgrounds) GetUrl() *string {
	return s.Url
}

func (s *StartStreamingOutShrinkRequestPanesBackgrounds) GetWidth() *float64 {
	return s.Width
}

func (s *StartStreamingOutShrinkRequestPanesBackgrounds) GetX() *float64 {
	return s.X
}

func (s *StartStreamingOutShrinkRequestPanesBackgrounds) GetY() *float64 {
	return s.Y
}

func (s *StartStreamingOutShrinkRequestPanesBackgrounds) SetAlpha(v float64) *StartStreamingOutShrinkRequestPanesBackgrounds {
	s.Alpha = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanesBackgrounds) SetDisplay(v string) *StartStreamingOutShrinkRequestPanesBackgrounds {
	s.Display = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanesBackgrounds) SetHeight(v float64) *StartStreamingOutShrinkRequestPanesBackgrounds {
	s.Height = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanesBackgrounds) SetLayer(v int32) *StartStreamingOutShrinkRequestPanesBackgrounds {
	s.Layer = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanesBackgrounds) SetPaneBackgroundCropMode(v int32) *StartStreamingOutShrinkRequestPanesBackgrounds {
	s.PaneBackgroundCropMode = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanesBackgrounds) SetUrl(v string) *StartStreamingOutShrinkRequestPanesBackgrounds {
	s.Url = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanesBackgrounds) SetWidth(v float64) *StartStreamingOutShrinkRequestPanesBackgrounds {
	s.Width = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanesBackgrounds) SetX(v float64) *StartStreamingOutShrinkRequestPanesBackgrounds {
	s.X = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanesBackgrounds) SetY(v float64) *StartStreamingOutShrinkRequestPanesBackgrounds {
	s.Y = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanesBackgrounds) Validate() error {
	return dara.Validate(s)
}

type StartStreamingOutShrinkRequestPanesImages struct {
	// example:
	//
	// 0.9
	Alpha *float64 `json:"Alpha,omitempty" xml:"Alpha,omitempty"`
	// example:
	//
	// backup
	Display *string `json:"Display,omitempty" xml:"Display,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	Height *float64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 0
	Layer *int32 `json:"Layer,omitempty" xml:"Layer,omitempty"`
	// example:
	//
	// 2
	PaneImageCropMode *int32 `json:"PaneImageCropMode,omitempty" xml:"PaneImageCropMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://aliyun.com/123xxx.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	Width *float64 `json:"Width,omitempty" xml:"Width,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	X *float64 `json:"X,omitempty" xml:"X,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.3
	Y *float64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s StartStreamingOutShrinkRequestPanesImages) String() string {
	return dara.Prettify(s)
}

func (s StartStreamingOutShrinkRequestPanesImages) GoString() string {
	return s.String()
}

func (s *StartStreamingOutShrinkRequestPanesImages) GetAlpha() *float64 {
	return s.Alpha
}

func (s *StartStreamingOutShrinkRequestPanesImages) GetDisplay() *string {
	return s.Display
}

func (s *StartStreamingOutShrinkRequestPanesImages) GetHeight() *float64 {
	return s.Height
}

func (s *StartStreamingOutShrinkRequestPanesImages) GetLayer() *int32 {
	return s.Layer
}

func (s *StartStreamingOutShrinkRequestPanesImages) GetPaneImageCropMode() *int32 {
	return s.PaneImageCropMode
}

func (s *StartStreamingOutShrinkRequestPanesImages) GetUrl() *string {
	return s.Url
}

func (s *StartStreamingOutShrinkRequestPanesImages) GetWidth() *float64 {
	return s.Width
}

func (s *StartStreamingOutShrinkRequestPanesImages) GetX() *float64 {
	return s.X
}

func (s *StartStreamingOutShrinkRequestPanesImages) GetY() *float64 {
	return s.Y
}

func (s *StartStreamingOutShrinkRequestPanesImages) SetAlpha(v float64) *StartStreamingOutShrinkRequestPanesImages {
	s.Alpha = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanesImages) SetDisplay(v string) *StartStreamingOutShrinkRequestPanesImages {
	s.Display = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanesImages) SetHeight(v float64) *StartStreamingOutShrinkRequestPanesImages {
	s.Height = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanesImages) SetLayer(v int32) *StartStreamingOutShrinkRequestPanesImages {
	s.Layer = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanesImages) SetPaneImageCropMode(v int32) *StartStreamingOutShrinkRequestPanesImages {
	s.PaneImageCropMode = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanesImages) SetUrl(v string) *StartStreamingOutShrinkRequestPanesImages {
	s.Url = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanesImages) SetWidth(v float64) *StartStreamingOutShrinkRequestPanesImages {
	s.Width = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanesImages) SetX(v float64) *StartStreamingOutShrinkRequestPanesImages {
	s.X = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanesImages) SetY(v float64) *StartStreamingOutShrinkRequestPanesImages {
	s.Y = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanesImages) Validate() error {
	return dara.Validate(s)
}

type StartStreamingOutShrinkRequestPanesTexts struct {
	// example:
	//
	// 0.9
	Alpha *float64 `json:"Alpha,omitempty" xml:"Alpha,omitempty"`
	// example:
	//
	// 0.6
	BoxAlpha *float64 `json:"BoxAlpha,omitempty" xml:"BoxAlpha,omitempty"`
	// example:
	//
	// 5
	BoxBorderw *int32                                            `json:"BoxBorderw,omitempty" xml:"BoxBorderw,omitempty"`
	BoxColor   *StartStreamingOutShrinkRequestPanesTextsBoxColor `json:"BoxColor,omitempty" xml:"BoxColor,omitempty" type:"Struct"`
	// example:
	//
	// backup
	Display *string `json:"Display,omitempty" xml:"Display,omitempty"`
	// example:
	//
	// 0
	Font      *int32                                             `json:"Font,omitempty" xml:"Font,omitempty"`
	FontColor *StartStreamingOutShrinkRequestPanesTextsFontColor `json:"FontColor,omitempty" xml:"FontColor,omitempty" type:"Struct"`
	// example:
	//
	// 36
	FontSize *int32 `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	HasBox   *bool  `json:"HasBox,omitempty" xml:"HasBox,omitempty"`
	// example:
	//
	// 0
	Layer *int32 `json:"Layer,omitempty" xml:"Layer,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 文字水印
	Texture *string `json:"Texture,omitempty" xml:"Texture,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	X *float64 `json:"X,omitempty" xml:"X,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	Y *float64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s StartStreamingOutShrinkRequestPanesTexts) String() string {
	return dara.Prettify(s)
}

func (s StartStreamingOutShrinkRequestPanesTexts) GoString() string {
	return s.String()
}

func (s *StartStreamingOutShrinkRequestPanesTexts) GetAlpha() *float64 {
	return s.Alpha
}

func (s *StartStreamingOutShrinkRequestPanesTexts) GetBoxAlpha() *float64 {
	return s.BoxAlpha
}

func (s *StartStreamingOutShrinkRequestPanesTexts) GetBoxBorderw() *int32 {
	return s.BoxBorderw
}

func (s *StartStreamingOutShrinkRequestPanesTexts) GetBoxColor() *StartStreamingOutShrinkRequestPanesTextsBoxColor {
	return s.BoxColor
}

func (s *StartStreamingOutShrinkRequestPanesTexts) GetDisplay() *string {
	return s.Display
}

func (s *StartStreamingOutShrinkRequestPanesTexts) GetFont() *int32 {
	return s.Font
}

func (s *StartStreamingOutShrinkRequestPanesTexts) GetFontColor() *StartStreamingOutShrinkRequestPanesTextsFontColor {
	return s.FontColor
}

func (s *StartStreamingOutShrinkRequestPanesTexts) GetFontSize() *int32 {
	return s.FontSize
}

func (s *StartStreamingOutShrinkRequestPanesTexts) GetHasBox() *bool {
	return s.HasBox
}

func (s *StartStreamingOutShrinkRequestPanesTexts) GetLayer() *int32 {
	return s.Layer
}

func (s *StartStreamingOutShrinkRequestPanesTexts) GetTexture() *string {
	return s.Texture
}

func (s *StartStreamingOutShrinkRequestPanesTexts) GetX() *float64 {
	return s.X
}

func (s *StartStreamingOutShrinkRequestPanesTexts) GetY() *float64 {
	return s.Y
}

func (s *StartStreamingOutShrinkRequestPanesTexts) SetAlpha(v float64) *StartStreamingOutShrinkRequestPanesTexts {
	s.Alpha = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanesTexts) SetBoxAlpha(v float64) *StartStreamingOutShrinkRequestPanesTexts {
	s.BoxAlpha = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanesTexts) SetBoxBorderw(v int32) *StartStreamingOutShrinkRequestPanesTexts {
	s.BoxBorderw = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanesTexts) SetBoxColor(v *StartStreamingOutShrinkRequestPanesTextsBoxColor) *StartStreamingOutShrinkRequestPanesTexts {
	s.BoxColor = v
	return s
}

func (s *StartStreamingOutShrinkRequestPanesTexts) SetDisplay(v string) *StartStreamingOutShrinkRequestPanesTexts {
	s.Display = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanesTexts) SetFont(v int32) *StartStreamingOutShrinkRequestPanesTexts {
	s.Font = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanesTexts) SetFontColor(v *StartStreamingOutShrinkRequestPanesTextsFontColor) *StartStreamingOutShrinkRequestPanesTexts {
	s.FontColor = v
	return s
}

func (s *StartStreamingOutShrinkRequestPanesTexts) SetFontSize(v int32) *StartStreamingOutShrinkRequestPanesTexts {
	s.FontSize = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanesTexts) SetHasBox(v bool) *StartStreamingOutShrinkRequestPanesTexts {
	s.HasBox = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanesTexts) SetLayer(v int32) *StartStreamingOutShrinkRequestPanesTexts {
	s.Layer = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanesTexts) SetTexture(v string) *StartStreamingOutShrinkRequestPanesTexts {
	s.Texture = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanesTexts) SetX(v float64) *StartStreamingOutShrinkRequestPanesTexts {
	s.X = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanesTexts) SetY(v float64) *StartStreamingOutShrinkRequestPanesTexts {
	s.Y = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanesTexts) Validate() error {
	return dara.Validate(s)
}

type StartStreamingOutShrinkRequestPanesTextsBoxColor struct {
	// example:
	//
	// 255
	B *int32 `json:"B,omitempty" xml:"B,omitempty"`
	// example:
	//
	// 255
	G *int32 `json:"G,omitempty" xml:"G,omitempty"`
	// example:
	//
	// 255
	R *int32 `json:"R,omitempty" xml:"R,omitempty"`
}

func (s StartStreamingOutShrinkRequestPanesTextsBoxColor) String() string {
	return dara.Prettify(s)
}

func (s StartStreamingOutShrinkRequestPanesTextsBoxColor) GoString() string {
	return s.String()
}

func (s *StartStreamingOutShrinkRequestPanesTextsBoxColor) GetB() *int32 {
	return s.B
}

func (s *StartStreamingOutShrinkRequestPanesTextsBoxColor) GetG() *int32 {
	return s.G
}

func (s *StartStreamingOutShrinkRequestPanesTextsBoxColor) GetR() *int32 {
	return s.R
}

func (s *StartStreamingOutShrinkRequestPanesTextsBoxColor) SetB(v int32) *StartStreamingOutShrinkRequestPanesTextsBoxColor {
	s.B = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanesTextsBoxColor) SetG(v int32) *StartStreamingOutShrinkRequestPanesTextsBoxColor {
	s.G = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanesTextsBoxColor) SetR(v int32) *StartStreamingOutShrinkRequestPanesTextsBoxColor {
	s.R = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanesTextsBoxColor) Validate() error {
	return dara.Validate(s)
}

type StartStreamingOutShrinkRequestPanesTextsFontColor struct {
	// example:
	//
	// 255
	B *int32 `json:"B,omitempty" xml:"B,omitempty"`
	// example:
	//
	// 255
	G *int32 `json:"G,omitempty" xml:"G,omitempty"`
	// example:
	//
	// 255
	R *int32 `json:"R,omitempty" xml:"R,omitempty"`
}

func (s StartStreamingOutShrinkRequestPanesTextsFontColor) String() string {
	return dara.Prettify(s)
}

func (s StartStreamingOutShrinkRequestPanesTextsFontColor) GoString() string {
	return s.String()
}

func (s *StartStreamingOutShrinkRequestPanesTextsFontColor) GetB() *int32 {
	return s.B
}

func (s *StartStreamingOutShrinkRequestPanesTextsFontColor) GetG() *int32 {
	return s.G
}

func (s *StartStreamingOutShrinkRequestPanesTextsFontColor) GetR() *int32 {
	return s.R
}

func (s *StartStreamingOutShrinkRequestPanesTextsFontColor) SetB(v int32) *StartStreamingOutShrinkRequestPanesTextsFontColor {
	s.B = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanesTextsFontColor) SetG(v int32) *StartStreamingOutShrinkRequestPanesTextsFontColor {
	s.G = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanesTextsFontColor) SetR(v int32) *StartStreamingOutShrinkRequestPanesTextsFontColor {
	s.R = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanesTextsFontColor) Validate() error {
	return dara.Validate(s)
}

type StartStreamingOutShrinkRequestPanesWhiteboard struct {
	// example:
	//
	// default
	WhiteboardId *string `json:"WhiteboardId,omitempty" xml:"WhiteboardId,omitempty"`
}

func (s StartStreamingOutShrinkRequestPanesWhiteboard) String() string {
	return dara.Prettify(s)
}

func (s StartStreamingOutShrinkRequestPanesWhiteboard) GoString() string {
	return s.String()
}

func (s *StartStreamingOutShrinkRequestPanesWhiteboard) GetWhiteboardId() *string {
	return s.WhiteboardId
}

func (s *StartStreamingOutShrinkRequestPanesWhiteboard) SetWhiteboardId(v string) *StartStreamingOutShrinkRequestPanesWhiteboard {
	s.WhiteboardId = &v
	return s
}

func (s *StartStreamingOutShrinkRequestPanesWhiteboard) Validate() error {
	return dara.Validate(s)
}

type StartStreamingOutShrinkRequestRegionColor struct {
	// example:
	//
	// 255
	B *int32 `json:"B,omitempty" xml:"B,omitempty"`
	// example:
	//
	// 255
	G *int32 `json:"G,omitempty" xml:"G,omitempty"`
	// example:
	//
	// 255
	R *int32 `json:"R,omitempty" xml:"R,omitempty"`
}

func (s StartStreamingOutShrinkRequestRegionColor) String() string {
	return dara.Prettify(s)
}

func (s StartStreamingOutShrinkRequestRegionColor) GoString() string {
	return s.String()
}

func (s *StartStreamingOutShrinkRequestRegionColor) GetB() *int32 {
	return s.B
}

func (s *StartStreamingOutShrinkRequestRegionColor) GetG() *int32 {
	return s.G
}

func (s *StartStreamingOutShrinkRequestRegionColor) GetR() *int32 {
	return s.R
}

func (s *StartStreamingOutShrinkRequestRegionColor) SetB(v int32) *StartStreamingOutShrinkRequestRegionColor {
	s.B = &v
	return s
}

func (s *StartStreamingOutShrinkRequestRegionColor) SetG(v int32) *StartStreamingOutShrinkRequestRegionColor {
	s.G = &v
	return s
}

func (s *StartStreamingOutShrinkRequestRegionColor) SetR(v int32) *StartStreamingOutShrinkRequestRegionColor {
	s.R = &v
	return s
}

func (s *StartStreamingOutShrinkRequestRegionColor) Validate() error {
	return dara.Validate(s)
}

type StartStreamingOutShrinkRequestTexts struct {
	// example:
	//
	// 0.1
	Alpha *float64 `json:"Alpha,omitempty" xml:"Alpha,omitempty"`
	// example:
	//
	// 0.6
	BoxAlpha *float64 `json:"BoxAlpha,omitempty" xml:"BoxAlpha,omitempty"`
	// example:
	//
	// 5
	BoxBorderw *int32                                       `json:"BoxBorderw,omitempty" xml:"BoxBorderw,omitempty"`
	BoxColor   *StartStreamingOutShrinkRequestTextsBoxColor `json:"BoxColor,omitempty" xml:"BoxColor,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Font      *int32                                        `json:"Font,omitempty" xml:"Font,omitempty"`
	FontColor *StartStreamingOutShrinkRequestTextsFontColor `json:"FontColor,omitempty" xml:"FontColor,omitempty" type:"Struct"`
	// example:
	//
	// 36
	FontSize *int32 `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	HasBox   *bool  `json:"HasBox,omitempty" xml:"HasBox,omitempty"`
	// example:
	//
	// 0
	Layer *int32 `json:"Layer,omitempty" xml:"Layer,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 文字水印
	Texture *string `json:"Texture,omitempty" xml:"Texture,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	X *float64 `json:"X,omitempty" xml:"X,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	Y *float64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s StartStreamingOutShrinkRequestTexts) String() string {
	return dara.Prettify(s)
}

func (s StartStreamingOutShrinkRequestTexts) GoString() string {
	return s.String()
}

func (s *StartStreamingOutShrinkRequestTexts) GetAlpha() *float64 {
	return s.Alpha
}

func (s *StartStreamingOutShrinkRequestTexts) GetBoxAlpha() *float64 {
	return s.BoxAlpha
}

func (s *StartStreamingOutShrinkRequestTexts) GetBoxBorderw() *int32 {
	return s.BoxBorderw
}

func (s *StartStreamingOutShrinkRequestTexts) GetBoxColor() *StartStreamingOutShrinkRequestTextsBoxColor {
	return s.BoxColor
}

func (s *StartStreamingOutShrinkRequestTexts) GetFont() *int32 {
	return s.Font
}

func (s *StartStreamingOutShrinkRequestTexts) GetFontColor() *StartStreamingOutShrinkRequestTextsFontColor {
	return s.FontColor
}

func (s *StartStreamingOutShrinkRequestTexts) GetFontSize() *int32 {
	return s.FontSize
}

func (s *StartStreamingOutShrinkRequestTexts) GetHasBox() *bool {
	return s.HasBox
}

func (s *StartStreamingOutShrinkRequestTexts) GetLayer() *int32 {
	return s.Layer
}

func (s *StartStreamingOutShrinkRequestTexts) GetTexture() *string {
	return s.Texture
}

func (s *StartStreamingOutShrinkRequestTexts) GetX() *float64 {
	return s.X
}

func (s *StartStreamingOutShrinkRequestTexts) GetY() *float64 {
	return s.Y
}

func (s *StartStreamingOutShrinkRequestTexts) SetAlpha(v float64) *StartStreamingOutShrinkRequestTexts {
	s.Alpha = &v
	return s
}

func (s *StartStreamingOutShrinkRequestTexts) SetBoxAlpha(v float64) *StartStreamingOutShrinkRequestTexts {
	s.BoxAlpha = &v
	return s
}

func (s *StartStreamingOutShrinkRequestTexts) SetBoxBorderw(v int32) *StartStreamingOutShrinkRequestTexts {
	s.BoxBorderw = &v
	return s
}

func (s *StartStreamingOutShrinkRequestTexts) SetBoxColor(v *StartStreamingOutShrinkRequestTextsBoxColor) *StartStreamingOutShrinkRequestTexts {
	s.BoxColor = v
	return s
}

func (s *StartStreamingOutShrinkRequestTexts) SetFont(v int32) *StartStreamingOutShrinkRequestTexts {
	s.Font = &v
	return s
}

func (s *StartStreamingOutShrinkRequestTexts) SetFontColor(v *StartStreamingOutShrinkRequestTextsFontColor) *StartStreamingOutShrinkRequestTexts {
	s.FontColor = v
	return s
}

func (s *StartStreamingOutShrinkRequestTexts) SetFontSize(v int32) *StartStreamingOutShrinkRequestTexts {
	s.FontSize = &v
	return s
}

func (s *StartStreamingOutShrinkRequestTexts) SetHasBox(v bool) *StartStreamingOutShrinkRequestTexts {
	s.HasBox = &v
	return s
}

func (s *StartStreamingOutShrinkRequestTexts) SetLayer(v int32) *StartStreamingOutShrinkRequestTexts {
	s.Layer = &v
	return s
}

func (s *StartStreamingOutShrinkRequestTexts) SetTexture(v string) *StartStreamingOutShrinkRequestTexts {
	s.Texture = &v
	return s
}

func (s *StartStreamingOutShrinkRequestTexts) SetX(v float64) *StartStreamingOutShrinkRequestTexts {
	s.X = &v
	return s
}

func (s *StartStreamingOutShrinkRequestTexts) SetY(v float64) *StartStreamingOutShrinkRequestTexts {
	s.Y = &v
	return s
}

func (s *StartStreamingOutShrinkRequestTexts) Validate() error {
	return dara.Validate(s)
}

type StartStreamingOutShrinkRequestTextsBoxColor struct {
	// example:
	//
	// 255
	B *int32 `json:"B,omitempty" xml:"B,omitempty"`
	// example:
	//
	// 255
	G *int32 `json:"G,omitempty" xml:"G,omitempty"`
	// example:
	//
	// 255
	R *int32 `json:"R,omitempty" xml:"R,omitempty"`
}

func (s StartStreamingOutShrinkRequestTextsBoxColor) String() string {
	return dara.Prettify(s)
}

func (s StartStreamingOutShrinkRequestTextsBoxColor) GoString() string {
	return s.String()
}

func (s *StartStreamingOutShrinkRequestTextsBoxColor) GetB() *int32 {
	return s.B
}

func (s *StartStreamingOutShrinkRequestTextsBoxColor) GetG() *int32 {
	return s.G
}

func (s *StartStreamingOutShrinkRequestTextsBoxColor) GetR() *int32 {
	return s.R
}

func (s *StartStreamingOutShrinkRequestTextsBoxColor) SetB(v int32) *StartStreamingOutShrinkRequestTextsBoxColor {
	s.B = &v
	return s
}

func (s *StartStreamingOutShrinkRequestTextsBoxColor) SetG(v int32) *StartStreamingOutShrinkRequestTextsBoxColor {
	s.G = &v
	return s
}

func (s *StartStreamingOutShrinkRequestTextsBoxColor) SetR(v int32) *StartStreamingOutShrinkRequestTextsBoxColor {
	s.R = &v
	return s
}

func (s *StartStreamingOutShrinkRequestTextsBoxColor) Validate() error {
	return dara.Validate(s)
}

type StartStreamingOutShrinkRequestTextsFontColor struct {
	// example:
	//
	// 255
	B *int32 `json:"B,omitempty" xml:"B,omitempty"`
	// example:
	//
	// 255
	G *int32 `json:"G,omitempty" xml:"G,omitempty"`
	// example:
	//
	// 255
	R *int32 `json:"R,omitempty" xml:"R,omitempty"`
}

func (s StartStreamingOutShrinkRequestTextsFontColor) String() string {
	return dara.Prettify(s)
}

func (s StartStreamingOutShrinkRequestTextsFontColor) GoString() string {
	return s.String()
}

func (s *StartStreamingOutShrinkRequestTextsFontColor) GetB() *int32 {
	return s.B
}

func (s *StartStreamingOutShrinkRequestTextsFontColor) GetG() *int32 {
	return s.G
}

func (s *StartStreamingOutShrinkRequestTextsFontColor) GetR() *int32 {
	return s.R
}

func (s *StartStreamingOutShrinkRequestTextsFontColor) SetB(v int32) *StartStreamingOutShrinkRequestTextsFontColor {
	s.B = &v
	return s
}

func (s *StartStreamingOutShrinkRequestTextsFontColor) SetG(v int32) *StartStreamingOutShrinkRequestTextsFontColor {
	s.G = &v
	return s
}

func (s *StartStreamingOutShrinkRequestTextsFontColor) SetR(v int32) *StartStreamingOutShrinkRequestTextsFontColor {
	s.R = &v
	return s
}

func (s *StartStreamingOutShrinkRequestTextsFontColor) Validate() error {
	return dara.Validate(s)
}
