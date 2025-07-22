// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStreamingOutShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateStreamingOutShrinkRequest
	GetAppId() *string
	SetBackgrounds(v []*UpdateStreamingOutShrinkRequestBackgrounds) *UpdateStreamingOutShrinkRequest
	GetBackgrounds() []*UpdateStreamingOutShrinkRequestBackgrounds
	SetBgColor(v *UpdateStreamingOutShrinkRequestBgColor) *UpdateStreamingOutShrinkRequest
	GetBgColor() *UpdateStreamingOutShrinkRequestBgColor
	SetChannelId(v string) *UpdateStreamingOutShrinkRequest
	GetChannelId() *string
	SetClockWidgets(v []*UpdateStreamingOutShrinkRequestClockWidgets) *UpdateStreamingOutShrinkRequest
	GetClockWidgets() []*UpdateStreamingOutShrinkRequestClockWidgets
	SetCropMode(v int32) *UpdateStreamingOutShrinkRequest
	GetCropMode() *int32
	SetImages(v []*UpdateStreamingOutShrinkRequestImages) *UpdateStreamingOutShrinkRequest
	GetImages() []*UpdateStreamingOutShrinkRequestImages
	SetLayoutSpecifiedUsersShrink(v string) *UpdateStreamingOutShrinkRequest
	GetLayoutSpecifiedUsersShrink() *string
	SetPanes(v []*UpdateStreamingOutShrinkRequestPanes) *UpdateStreamingOutShrinkRequest
	GetPanes() []*UpdateStreamingOutShrinkRequestPanes
	SetRegionColor(v *UpdateStreamingOutShrinkRequestRegionColor) *UpdateStreamingOutShrinkRequest
	GetRegionColor() *UpdateStreamingOutShrinkRequestRegionColor
	SetSpecMixedUserList(v []*string) *UpdateStreamingOutShrinkRequest
	GetSpecMixedUserList() []*string
	SetTaskId(v string) *UpdateStreamingOutShrinkRequest
	GetTaskId() *string
	SetTemplateId(v string) *UpdateStreamingOutShrinkRequest
	GetTemplateId() *string
	SetTexts(v []*UpdateStreamingOutShrinkRequestTexts) *UpdateStreamingOutShrinkRequest
	GetTexts() []*UpdateStreamingOutShrinkRequestTexts
}

type UpdateStreamingOutShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId       *string                                       `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Backgrounds []*UpdateStreamingOutShrinkRequestBackgrounds `json:"Backgrounds,omitempty" xml:"Backgrounds,omitempty" type:"Repeated"`
	BgColor     *UpdateStreamingOutShrinkRequestBgColor       `json:"BgColor,omitempty" xml:"BgColor,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId    *string                                        `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ClockWidgets []*UpdateStreamingOutShrinkRequestClockWidgets `json:"ClockWidgets,omitempty" xml:"ClockWidgets,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	CropMode                   *int32                                      `json:"CropMode,omitempty" xml:"CropMode,omitempty"`
	Images                     []*UpdateStreamingOutShrinkRequestImages    `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	LayoutSpecifiedUsersShrink *string                                     `json:"LayoutSpecifiedUsers,omitempty" xml:"LayoutSpecifiedUsers,omitempty"`
	Panes                      []*UpdateStreamingOutShrinkRequestPanes     `json:"Panes,omitempty" xml:"Panes,omitempty" type:"Repeated"`
	RegionColor                *UpdateStreamingOutShrinkRequestRegionColor `json:"RegionColor,omitempty" xml:"RegionColor,omitempty" type:"Struct"`
	SpecMixedUserList          []*string                                   `json:"SpecMixedUserList,omitempty" xml:"SpecMixedUserList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 567
	TemplateId *string                                 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Texts      []*UpdateStreamingOutShrinkRequestTexts `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
}

func (s UpdateStreamingOutShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamingOutShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateStreamingOutShrinkRequest) GetBackgrounds() []*UpdateStreamingOutShrinkRequestBackgrounds {
	return s.Backgrounds
}

func (s *UpdateStreamingOutShrinkRequest) GetBgColor() *UpdateStreamingOutShrinkRequestBgColor {
	return s.BgColor
}

func (s *UpdateStreamingOutShrinkRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *UpdateStreamingOutShrinkRequest) GetClockWidgets() []*UpdateStreamingOutShrinkRequestClockWidgets {
	return s.ClockWidgets
}

func (s *UpdateStreamingOutShrinkRequest) GetCropMode() *int32 {
	return s.CropMode
}

func (s *UpdateStreamingOutShrinkRequest) GetImages() []*UpdateStreamingOutShrinkRequestImages {
	return s.Images
}

func (s *UpdateStreamingOutShrinkRequest) GetLayoutSpecifiedUsersShrink() *string {
	return s.LayoutSpecifiedUsersShrink
}

func (s *UpdateStreamingOutShrinkRequest) GetPanes() []*UpdateStreamingOutShrinkRequestPanes {
	return s.Panes
}

func (s *UpdateStreamingOutShrinkRequest) GetRegionColor() *UpdateStreamingOutShrinkRequestRegionColor {
	return s.RegionColor
}

func (s *UpdateStreamingOutShrinkRequest) GetSpecMixedUserList() []*string {
	return s.SpecMixedUserList
}

func (s *UpdateStreamingOutShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateStreamingOutShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateStreamingOutShrinkRequest) GetTexts() []*UpdateStreamingOutShrinkRequestTexts {
	return s.Texts
}

func (s *UpdateStreamingOutShrinkRequest) SetAppId(v string) *UpdateStreamingOutShrinkRequest {
	s.AppId = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequest) SetBackgrounds(v []*UpdateStreamingOutShrinkRequestBackgrounds) *UpdateStreamingOutShrinkRequest {
	s.Backgrounds = v
	return s
}

func (s *UpdateStreamingOutShrinkRequest) SetBgColor(v *UpdateStreamingOutShrinkRequestBgColor) *UpdateStreamingOutShrinkRequest {
	s.BgColor = v
	return s
}

func (s *UpdateStreamingOutShrinkRequest) SetChannelId(v string) *UpdateStreamingOutShrinkRequest {
	s.ChannelId = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequest) SetClockWidgets(v []*UpdateStreamingOutShrinkRequestClockWidgets) *UpdateStreamingOutShrinkRequest {
	s.ClockWidgets = v
	return s
}

func (s *UpdateStreamingOutShrinkRequest) SetCropMode(v int32) *UpdateStreamingOutShrinkRequest {
	s.CropMode = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequest) SetImages(v []*UpdateStreamingOutShrinkRequestImages) *UpdateStreamingOutShrinkRequest {
	s.Images = v
	return s
}

func (s *UpdateStreamingOutShrinkRequest) SetLayoutSpecifiedUsersShrink(v string) *UpdateStreamingOutShrinkRequest {
	s.LayoutSpecifiedUsersShrink = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequest) SetPanes(v []*UpdateStreamingOutShrinkRequestPanes) *UpdateStreamingOutShrinkRequest {
	s.Panes = v
	return s
}

func (s *UpdateStreamingOutShrinkRequest) SetRegionColor(v *UpdateStreamingOutShrinkRequestRegionColor) *UpdateStreamingOutShrinkRequest {
	s.RegionColor = v
	return s
}

func (s *UpdateStreamingOutShrinkRequest) SetSpecMixedUserList(v []*string) *UpdateStreamingOutShrinkRequest {
	s.SpecMixedUserList = v
	return s
}

func (s *UpdateStreamingOutShrinkRequest) SetTaskId(v string) *UpdateStreamingOutShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequest) SetTemplateId(v string) *UpdateStreamingOutShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequest) SetTexts(v []*UpdateStreamingOutShrinkRequestTexts) *UpdateStreamingOutShrinkRequest {
	s.Texts = v
	return s
}

func (s *UpdateStreamingOutShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateStreamingOutShrinkRequestBackgrounds struct {
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

func (s UpdateStreamingOutShrinkRequestBackgrounds) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamingOutShrinkRequestBackgrounds) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutShrinkRequestBackgrounds) GetAlpha() *float64 {
	return s.Alpha
}

func (s *UpdateStreamingOutShrinkRequestBackgrounds) GetBackgroundCropMode() *int32 {
	return s.BackgroundCropMode
}

func (s *UpdateStreamingOutShrinkRequestBackgrounds) GetHeight() *float64 {
	return s.Height
}

func (s *UpdateStreamingOutShrinkRequestBackgrounds) GetLayer() *int32 {
	return s.Layer
}

func (s *UpdateStreamingOutShrinkRequestBackgrounds) GetUrl() *string {
	return s.Url
}

func (s *UpdateStreamingOutShrinkRequestBackgrounds) GetWidth() *float64 {
	return s.Width
}

func (s *UpdateStreamingOutShrinkRequestBackgrounds) GetX() *float64 {
	return s.X
}

func (s *UpdateStreamingOutShrinkRequestBackgrounds) GetY() *float64 {
	return s.Y
}

func (s *UpdateStreamingOutShrinkRequestBackgrounds) SetAlpha(v float64) *UpdateStreamingOutShrinkRequestBackgrounds {
	s.Alpha = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestBackgrounds) SetBackgroundCropMode(v int32) *UpdateStreamingOutShrinkRequestBackgrounds {
	s.BackgroundCropMode = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestBackgrounds) SetHeight(v float64) *UpdateStreamingOutShrinkRequestBackgrounds {
	s.Height = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestBackgrounds) SetLayer(v int32) *UpdateStreamingOutShrinkRequestBackgrounds {
	s.Layer = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestBackgrounds) SetUrl(v string) *UpdateStreamingOutShrinkRequestBackgrounds {
	s.Url = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestBackgrounds) SetWidth(v float64) *UpdateStreamingOutShrinkRequestBackgrounds {
	s.Width = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestBackgrounds) SetX(v float64) *UpdateStreamingOutShrinkRequestBackgrounds {
	s.X = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestBackgrounds) SetY(v float64) *UpdateStreamingOutShrinkRequestBackgrounds {
	s.Y = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestBackgrounds) Validate() error {
	return dara.Validate(s)
}

type UpdateStreamingOutShrinkRequestBgColor struct {
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

func (s UpdateStreamingOutShrinkRequestBgColor) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamingOutShrinkRequestBgColor) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutShrinkRequestBgColor) GetB() *int32 {
	return s.B
}

func (s *UpdateStreamingOutShrinkRequestBgColor) GetG() *int32 {
	return s.G
}

func (s *UpdateStreamingOutShrinkRequestBgColor) GetR() *int32 {
	return s.R
}

func (s *UpdateStreamingOutShrinkRequestBgColor) SetB(v int32) *UpdateStreamingOutShrinkRequestBgColor {
	s.B = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestBgColor) SetG(v int32) *UpdateStreamingOutShrinkRequestBgColor {
	s.G = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestBgColor) SetR(v int32) *UpdateStreamingOutShrinkRequestBgColor {
	s.R = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestBgColor) Validate() error {
	return dara.Validate(s)
}

type UpdateStreamingOutShrinkRequestClockWidgets struct {
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
	BoxBorderw *int32                                               `json:"BoxBorderw,omitempty" xml:"BoxBorderw,omitempty"`
	BoxColor   *UpdateStreamingOutShrinkRequestClockWidgetsBoxColor `json:"BoxColor,omitempty" xml:"BoxColor,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Font      *int32                                                `json:"Font,omitempty" xml:"Font,omitempty"`
	FontColor *UpdateStreamingOutShrinkRequestClockWidgetsFontColor `json:"FontColor,omitempty" xml:"FontColor,omitempty" type:"Struct"`
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

func (s UpdateStreamingOutShrinkRequestClockWidgets) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamingOutShrinkRequestClockWidgets) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutShrinkRequestClockWidgets) GetAlpha() *float64 {
	return s.Alpha
}

func (s *UpdateStreamingOutShrinkRequestClockWidgets) GetBoxAlpha() *float64 {
	return s.BoxAlpha
}

func (s *UpdateStreamingOutShrinkRequestClockWidgets) GetBoxBorderw() *int32 {
	return s.BoxBorderw
}

func (s *UpdateStreamingOutShrinkRequestClockWidgets) GetBoxColor() *UpdateStreamingOutShrinkRequestClockWidgetsBoxColor {
	return s.BoxColor
}

func (s *UpdateStreamingOutShrinkRequestClockWidgets) GetFont() *int32 {
	return s.Font
}

func (s *UpdateStreamingOutShrinkRequestClockWidgets) GetFontColor() *UpdateStreamingOutShrinkRequestClockWidgetsFontColor {
	return s.FontColor
}

func (s *UpdateStreamingOutShrinkRequestClockWidgets) GetFontSize() *int32 {
	return s.FontSize
}

func (s *UpdateStreamingOutShrinkRequestClockWidgets) GetHasBox() *bool {
	return s.HasBox
}

func (s *UpdateStreamingOutShrinkRequestClockWidgets) GetLayer() *int32 {
	return s.Layer
}

func (s *UpdateStreamingOutShrinkRequestClockWidgets) GetX() *float64 {
	return s.X
}

func (s *UpdateStreamingOutShrinkRequestClockWidgets) GetY() *float64 {
	return s.Y
}

func (s *UpdateStreamingOutShrinkRequestClockWidgets) GetZone() *int32 {
	return s.Zone
}

func (s *UpdateStreamingOutShrinkRequestClockWidgets) SetAlpha(v float64) *UpdateStreamingOutShrinkRequestClockWidgets {
	s.Alpha = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestClockWidgets) SetBoxAlpha(v float64) *UpdateStreamingOutShrinkRequestClockWidgets {
	s.BoxAlpha = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestClockWidgets) SetBoxBorderw(v int32) *UpdateStreamingOutShrinkRequestClockWidgets {
	s.BoxBorderw = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestClockWidgets) SetBoxColor(v *UpdateStreamingOutShrinkRequestClockWidgetsBoxColor) *UpdateStreamingOutShrinkRequestClockWidgets {
	s.BoxColor = v
	return s
}

func (s *UpdateStreamingOutShrinkRequestClockWidgets) SetFont(v int32) *UpdateStreamingOutShrinkRequestClockWidgets {
	s.Font = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestClockWidgets) SetFontColor(v *UpdateStreamingOutShrinkRequestClockWidgetsFontColor) *UpdateStreamingOutShrinkRequestClockWidgets {
	s.FontColor = v
	return s
}

func (s *UpdateStreamingOutShrinkRequestClockWidgets) SetFontSize(v int32) *UpdateStreamingOutShrinkRequestClockWidgets {
	s.FontSize = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestClockWidgets) SetHasBox(v bool) *UpdateStreamingOutShrinkRequestClockWidgets {
	s.HasBox = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestClockWidgets) SetLayer(v int32) *UpdateStreamingOutShrinkRequestClockWidgets {
	s.Layer = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestClockWidgets) SetX(v float64) *UpdateStreamingOutShrinkRequestClockWidgets {
	s.X = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestClockWidgets) SetY(v float64) *UpdateStreamingOutShrinkRequestClockWidgets {
	s.Y = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestClockWidgets) SetZone(v int32) *UpdateStreamingOutShrinkRequestClockWidgets {
	s.Zone = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestClockWidgets) Validate() error {
	return dara.Validate(s)
}

type UpdateStreamingOutShrinkRequestClockWidgetsBoxColor struct {
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

func (s UpdateStreamingOutShrinkRequestClockWidgetsBoxColor) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamingOutShrinkRequestClockWidgetsBoxColor) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutShrinkRequestClockWidgetsBoxColor) GetB() *int32 {
	return s.B
}

func (s *UpdateStreamingOutShrinkRequestClockWidgetsBoxColor) GetG() *int32 {
	return s.G
}

func (s *UpdateStreamingOutShrinkRequestClockWidgetsBoxColor) GetR() *int32 {
	return s.R
}

func (s *UpdateStreamingOutShrinkRequestClockWidgetsBoxColor) SetB(v int32) *UpdateStreamingOutShrinkRequestClockWidgetsBoxColor {
	s.B = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestClockWidgetsBoxColor) SetG(v int32) *UpdateStreamingOutShrinkRequestClockWidgetsBoxColor {
	s.G = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestClockWidgetsBoxColor) SetR(v int32) *UpdateStreamingOutShrinkRequestClockWidgetsBoxColor {
	s.R = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestClockWidgetsBoxColor) Validate() error {
	return dara.Validate(s)
}

type UpdateStreamingOutShrinkRequestClockWidgetsFontColor struct {
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

func (s UpdateStreamingOutShrinkRequestClockWidgetsFontColor) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamingOutShrinkRequestClockWidgetsFontColor) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutShrinkRequestClockWidgetsFontColor) GetB() *int32 {
	return s.B
}

func (s *UpdateStreamingOutShrinkRequestClockWidgetsFontColor) GetG() *int32 {
	return s.G
}

func (s *UpdateStreamingOutShrinkRequestClockWidgetsFontColor) GetR() *int32 {
	return s.R
}

func (s *UpdateStreamingOutShrinkRequestClockWidgetsFontColor) SetB(v int32) *UpdateStreamingOutShrinkRequestClockWidgetsFontColor {
	s.B = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestClockWidgetsFontColor) SetG(v int32) *UpdateStreamingOutShrinkRequestClockWidgetsFontColor {
	s.G = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestClockWidgetsFontColor) SetR(v int32) *UpdateStreamingOutShrinkRequestClockWidgetsFontColor {
	s.R = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestClockWidgetsFontColor) Validate() error {
	return dara.Validate(s)
}

type UpdateStreamingOutShrinkRequestImages struct {
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

func (s UpdateStreamingOutShrinkRequestImages) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamingOutShrinkRequestImages) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutShrinkRequestImages) GetAlpha() *float64 {
	return s.Alpha
}

func (s *UpdateStreamingOutShrinkRequestImages) GetHeight() *float64 {
	return s.Height
}

func (s *UpdateStreamingOutShrinkRequestImages) GetImageCropMode() *int32 {
	return s.ImageCropMode
}

func (s *UpdateStreamingOutShrinkRequestImages) GetLayer() *int32 {
	return s.Layer
}

func (s *UpdateStreamingOutShrinkRequestImages) GetUrl() *string {
	return s.Url
}

func (s *UpdateStreamingOutShrinkRequestImages) GetWidth() *float64 {
	return s.Width
}

func (s *UpdateStreamingOutShrinkRequestImages) GetX() *float64 {
	return s.X
}

func (s *UpdateStreamingOutShrinkRequestImages) GetY() *float64 {
	return s.Y
}

func (s *UpdateStreamingOutShrinkRequestImages) SetAlpha(v float64) *UpdateStreamingOutShrinkRequestImages {
	s.Alpha = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestImages) SetHeight(v float64) *UpdateStreamingOutShrinkRequestImages {
	s.Height = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestImages) SetImageCropMode(v int32) *UpdateStreamingOutShrinkRequestImages {
	s.ImageCropMode = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestImages) SetLayer(v int32) *UpdateStreamingOutShrinkRequestImages {
	s.Layer = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestImages) SetUrl(v string) *UpdateStreamingOutShrinkRequestImages {
	s.Url = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestImages) SetWidth(v float64) *UpdateStreamingOutShrinkRequestImages {
	s.Width = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestImages) SetX(v float64) *UpdateStreamingOutShrinkRequestImages {
	s.X = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestImages) SetY(v float64) *UpdateStreamingOutShrinkRequestImages {
	s.Y = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestImages) Validate() error {
	return dara.Validate(s)
}

type UpdateStreamingOutShrinkRequestPanes struct {
	Backgrounds []*UpdateStreamingOutShrinkRequestPanesBackgrounds `json:"Backgrounds,omitempty" xml:"Backgrounds,omitempty" type:"Repeated"`
	Images      []*UpdateStreamingOutShrinkRequestPanesImages      `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	PaneCropMode *int32 `json:"PaneCropMode,omitempty" xml:"PaneCropMode,omitempty"`
	// example:
	//
	// 1
	PaneId                    *int32 `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
	ReservePaneForOfflineUser *bool  `json:"ReservePaneForOfflineUser,omitempty" xml:"ReservePaneForOfflineUser,omitempty"`
	// example:
	//
	// 22
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// video
	SourceType *string                                      `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Texts      []*UpdateStreamingOutShrinkRequestPanesTexts `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
	// example:
	//
	// cameraFirst
	VideoOrder *string                                         `json:"VideoOrder,omitempty" xml:"VideoOrder,omitempty"`
	Whiteboard *UpdateStreamingOutShrinkRequestPanesWhiteboard `json:"Whiteboard,omitempty" xml:"Whiteboard,omitempty" type:"Struct"`
}

func (s UpdateStreamingOutShrinkRequestPanes) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamingOutShrinkRequestPanes) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutShrinkRequestPanes) GetBackgrounds() []*UpdateStreamingOutShrinkRequestPanesBackgrounds {
	return s.Backgrounds
}

func (s *UpdateStreamingOutShrinkRequestPanes) GetImages() []*UpdateStreamingOutShrinkRequestPanesImages {
	return s.Images
}

func (s *UpdateStreamingOutShrinkRequestPanes) GetPaneCropMode() *int32 {
	return s.PaneCropMode
}

func (s *UpdateStreamingOutShrinkRequestPanes) GetPaneId() *int32 {
	return s.PaneId
}

func (s *UpdateStreamingOutShrinkRequestPanes) GetReservePaneForOfflineUser() *bool {
	return s.ReservePaneForOfflineUser
}

func (s *UpdateStreamingOutShrinkRequestPanes) GetSource() *string {
	return s.Source
}

func (s *UpdateStreamingOutShrinkRequestPanes) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateStreamingOutShrinkRequestPanes) GetTexts() []*UpdateStreamingOutShrinkRequestPanesTexts {
	return s.Texts
}

func (s *UpdateStreamingOutShrinkRequestPanes) GetVideoOrder() *string {
	return s.VideoOrder
}

func (s *UpdateStreamingOutShrinkRequestPanes) GetWhiteboard() *UpdateStreamingOutShrinkRequestPanesWhiteboard {
	return s.Whiteboard
}

func (s *UpdateStreamingOutShrinkRequestPanes) SetBackgrounds(v []*UpdateStreamingOutShrinkRequestPanesBackgrounds) *UpdateStreamingOutShrinkRequestPanes {
	s.Backgrounds = v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanes) SetImages(v []*UpdateStreamingOutShrinkRequestPanesImages) *UpdateStreamingOutShrinkRequestPanes {
	s.Images = v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanes) SetPaneCropMode(v int32) *UpdateStreamingOutShrinkRequestPanes {
	s.PaneCropMode = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanes) SetPaneId(v int32) *UpdateStreamingOutShrinkRequestPanes {
	s.PaneId = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanes) SetReservePaneForOfflineUser(v bool) *UpdateStreamingOutShrinkRequestPanes {
	s.ReservePaneForOfflineUser = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanes) SetSource(v string) *UpdateStreamingOutShrinkRequestPanes {
	s.Source = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanes) SetSourceType(v string) *UpdateStreamingOutShrinkRequestPanes {
	s.SourceType = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanes) SetTexts(v []*UpdateStreamingOutShrinkRequestPanesTexts) *UpdateStreamingOutShrinkRequestPanes {
	s.Texts = v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanes) SetVideoOrder(v string) *UpdateStreamingOutShrinkRequestPanes {
	s.VideoOrder = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanes) SetWhiteboard(v *UpdateStreamingOutShrinkRequestPanesWhiteboard) *UpdateStreamingOutShrinkRequestPanes {
	s.Whiteboard = v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanes) Validate() error {
	return dara.Validate(s)
}

type UpdateStreamingOutShrinkRequestPanesBackgrounds struct {
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

func (s UpdateStreamingOutShrinkRequestPanesBackgrounds) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamingOutShrinkRequestPanesBackgrounds) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutShrinkRequestPanesBackgrounds) GetAlpha() *float64 {
	return s.Alpha
}

func (s *UpdateStreamingOutShrinkRequestPanesBackgrounds) GetDisplay() *string {
	return s.Display
}

func (s *UpdateStreamingOutShrinkRequestPanesBackgrounds) GetHeight() *float64 {
	return s.Height
}

func (s *UpdateStreamingOutShrinkRequestPanesBackgrounds) GetLayer() *int32 {
	return s.Layer
}

func (s *UpdateStreamingOutShrinkRequestPanesBackgrounds) GetPaneBackgroundCropMode() *int32 {
	return s.PaneBackgroundCropMode
}

func (s *UpdateStreamingOutShrinkRequestPanesBackgrounds) GetUrl() *string {
	return s.Url
}

func (s *UpdateStreamingOutShrinkRequestPanesBackgrounds) GetWidth() *float64 {
	return s.Width
}

func (s *UpdateStreamingOutShrinkRequestPanesBackgrounds) GetX() *float64 {
	return s.X
}

func (s *UpdateStreamingOutShrinkRequestPanesBackgrounds) GetY() *float64 {
	return s.Y
}

func (s *UpdateStreamingOutShrinkRequestPanesBackgrounds) SetAlpha(v float64) *UpdateStreamingOutShrinkRequestPanesBackgrounds {
	s.Alpha = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanesBackgrounds) SetDisplay(v string) *UpdateStreamingOutShrinkRequestPanesBackgrounds {
	s.Display = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanesBackgrounds) SetHeight(v float64) *UpdateStreamingOutShrinkRequestPanesBackgrounds {
	s.Height = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanesBackgrounds) SetLayer(v int32) *UpdateStreamingOutShrinkRequestPanesBackgrounds {
	s.Layer = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanesBackgrounds) SetPaneBackgroundCropMode(v int32) *UpdateStreamingOutShrinkRequestPanesBackgrounds {
	s.PaneBackgroundCropMode = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanesBackgrounds) SetUrl(v string) *UpdateStreamingOutShrinkRequestPanesBackgrounds {
	s.Url = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanesBackgrounds) SetWidth(v float64) *UpdateStreamingOutShrinkRequestPanesBackgrounds {
	s.Width = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanesBackgrounds) SetX(v float64) *UpdateStreamingOutShrinkRequestPanesBackgrounds {
	s.X = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanesBackgrounds) SetY(v float64) *UpdateStreamingOutShrinkRequestPanesBackgrounds {
	s.Y = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanesBackgrounds) Validate() error {
	return dara.Validate(s)
}

type UpdateStreamingOutShrinkRequestPanesImages struct {
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

func (s UpdateStreamingOutShrinkRequestPanesImages) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamingOutShrinkRequestPanesImages) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutShrinkRequestPanesImages) GetAlpha() *float64 {
	return s.Alpha
}

func (s *UpdateStreamingOutShrinkRequestPanesImages) GetDisplay() *string {
	return s.Display
}

func (s *UpdateStreamingOutShrinkRequestPanesImages) GetHeight() *float64 {
	return s.Height
}

func (s *UpdateStreamingOutShrinkRequestPanesImages) GetLayer() *int32 {
	return s.Layer
}

func (s *UpdateStreamingOutShrinkRequestPanesImages) GetPaneImageCropMode() *int32 {
	return s.PaneImageCropMode
}

func (s *UpdateStreamingOutShrinkRequestPanesImages) GetUrl() *string {
	return s.Url
}

func (s *UpdateStreamingOutShrinkRequestPanesImages) GetWidth() *float64 {
	return s.Width
}

func (s *UpdateStreamingOutShrinkRequestPanesImages) GetX() *float64 {
	return s.X
}

func (s *UpdateStreamingOutShrinkRequestPanesImages) GetY() *float64 {
	return s.Y
}

func (s *UpdateStreamingOutShrinkRequestPanesImages) SetAlpha(v float64) *UpdateStreamingOutShrinkRequestPanesImages {
	s.Alpha = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanesImages) SetDisplay(v string) *UpdateStreamingOutShrinkRequestPanesImages {
	s.Display = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanesImages) SetHeight(v float64) *UpdateStreamingOutShrinkRequestPanesImages {
	s.Height = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanesImages) SetLayer(v int32) *UpdateStreamingOutShrinkRequestPanesImages {
	s.Layer = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanesImages) SetPaneImageCropMode(v int32) *UpdateStreamingOutShrinkRequestPanesImages {
	s.PaneImageCropMode = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanesImages) SetUrl(v string) *UpdateStreamingOutShrinkRequestPanesImages {
	s.Url = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanesImages) SetWidth(v float64) *UpdateStreamingOutShrinkRequestPanesImages {
	s.Width = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanesImages) SetX(v float64) *UpdateStreamingOutShrinkRequestPanesImages {
	s.X = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanesImages) SetY(v float64) *UpdateStreamingOutShrinkRequestPanesImages {
	s.Y = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanesImages) Validate() error {
	return dara.Validate(s)
}

type UpdateStreamingOutShrinkRequestPanesTexts struct {
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
	BoxBorderw *int32                                             `json:"BoxBorderw,omitempty" xml:"BoxBorderw,omitempty"`
	BoxColor   *UpdateStreamingOutShrinkRequestPanesTextsBoxColor `json:"BoxColor,omitempty" xml:"BoxColor,omitempty" type:"Struct"`
	// example:
	//
	// backup
	Display *string `json:"Display,omitempty" xml:"Display,omitempty"`
	// example:
	//
	// 0
	Font      *int32                                              `json:"Font,omitempty" xml:"Font,omitempty"`
	FontColor *UpdateStreamingOutShrinkRequestPanesTextsFontColor `json:"FontColor,omitempty" xml:"FontColor,omitempty" type:"Struct"`
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

func (s UpdateStreamingOutShrinkRequestPanesTexts) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamingOutShrinkRequestPanesTexts) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutShrinkRequestPanesTexts) GetAlpha() *float64 {
	return s.Alpha
}

func (s *UpdateStreamingOutShrinkRequestPanesTexts) GetBoxAlpha() *float64 {
	return s.BoxAlpha
}

func (s *UpdateStreamingOutShrinkRequestPanesTexts) GetBoxBorderw() *int32 {
	return s.BoxBorderw
}

func (s *UpdateStreamingOutShrinkRequestPanesTexts) GetBoxColor() *UpdateStreamingOutShrinkRequestPanesTextsBoxColor {
	return s.BoxColor
}

func (s *UpdateStreamingOutShrinkRequestPanesTexts) GetDisplay() *string {
	return s.Display
}

func (s *UpdateStreamingOutShrinkRequestPanesTexts) GetFont() *int32 {
	return s.Font
}

func (s *UpdateStreamingOutShrinkRequestPanesTexts) GetFontColor() *UpdateStreamingOutShrinkRequestPanesTextsFontColor {
	return s.FontColor
}

func (s *UpdateStreamingOutShrinkRequestPanesTexts) GetFontSize() *int32 {
	return s.FontSize
}

func (s *UpdateStreamingOutShrinkRequestPanesTexts) GetHasBox() *bool {
	return s.HasBox
}

func (s *UpdateStreamingOutShrinkRequestPanesTexts) GetLayer() *int32 {
	return s.Layer
}

func (s *UpdateStreamingOutShrinkRequestPanesTexts) GetTexture() *string {
	return s.Texture
}

func (s *UpdateStreamingOutShrinkRequestPanesTexts) GetX() *float64 {
	return s.X
}

func (s *UpdateStreamingOutShrinkRequestPanesTexts) GetY() *float64 {
	return s.Y
}

func (s *UpdateStreamingOutShrinkRequestPanesTexts) SetAlpha(v float64) *UpdateStreamingOutShrinkRequestPanesTexts {
	s.Alpha = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanesTexts) SetBoxAlpha(v float64) *UpdateStreamingOutShrinkRequestPanesTexts {
	s.BoxAlpha = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanesTexts) SetBoxBorderw(v int32) *UpdateStreamingOutShrinkRequestPanesTexts {
	s.BoxBorderw = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanesTexts) SetBoxColor(v *UpdateStreamingOutShrinkRequestPanesTextsBoxColor) *UpdateStreamingOutShrinkRequestPanesTexts {
	s.BoxColor = v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanesTexts) SetDisplay(v string) *UpdateStreamingOutShrinkRequestPanesTexts {
	s.Display = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanesTexts) SetFont(v int32) *UpdateStreamingOutShrinkRequestPanesTexts {
	s.Font = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanesTexts) SetFontColor(v *UpdateStreamingOutShrinkRequestPanesTextsFontColor) *UpdateStreamingOutShrinkRequestPanesTexts {
	s.FontColor = v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanesTexts) SetFontSize(v int32) *UpdateStreamingOutShrinkRequestPanesTexts {
	s.FontSize = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanesTexts) SetHasBox(v bool) *UpdateStreamingOutShrinkRequestPanesTexts {
	s.HasBox = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanesTexts) SetLayer(v int32) *UpdateStreamingOutShrinkRequestPanesTexts {
	s.Layer = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanesTexts) SetTexture(v string) *UpdateStreamingOutShrinkRequestPanesTexts {
	s.Texture = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanesTexts) SetX(v float64) *UpdateStreamingOutShrinkRequestPanesTexts {
	s.X = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanesTexts) SetY(v float64) *UpdateStreamingOutShrinkRequestPanesTexts {
	s.Y = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanesTexts) Validate() error {
	return dara.Validate(s)
}

type UpdateStreamingOutShrinkRequestPanesTextsBoxColor struct {
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

func (s UpdateStreamingOutShrinkRequestPanesTextsBoxColor) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamingOutShrinkRequestPanesTextsBoxColor) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutShrinkRequestPanesTextsBoxColor) GetB() *int32 {
	return s.B
}

func (s *UpdateStreamingOutShrinkRequestPanesTextsBoxColor) GetG() *int32 {
	return s.G
}

func (s *UpdateStreamingOutShrinkRequestPanesTextsBoxColor) GetR() *int32 {
	return s.R
}

func (s *UpdateStreamingOutShrinkRequestPanesTextsBoxColor) SetB(v int32) *UpdateStreamingOutShrinkRequestPanesTextsBoxColor {
	s.B = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanesTextsBoxColor) SetG(v int32) *UpdateStreamingOutShrinkRequestPanesTextsBoxColor {
	s.G = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanesTextsBoxColor) SetR(v int32) *UpdateStreamingOutShrinkRequestPanesTextsBoxColor {
	s.R = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanesTextsBoxColor) Validate() error {
	return dara.Validate(s)
}

type UpdateStreamingOutShrinkRequestPanesTextsFontColor struct {
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

func (s UpdateStreamingOutShrinkRequestPanesTextsFontColor) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamingOutShrinkRequestPanesTextsFontColor) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutShrinkRequestPanesTextsFontColor) GetB() *int32 {
	return s.B
}

func (s *UpdateStreamingOutShrinkRequestPanesTextsFontColor) GetG() *int32 {
	return s.G
}

func (s *UpdateStreamingOutShrinkRequestPanesTextsFontColor) GetR() *int32 {
	return s.R
}

func (s *UpdateStreamingOutShrinkRequestPanesTextsFontColor) SetB(v int32) *UpdateStreamingOutShrinkRequestPanesTextsFontColor {
	s.B = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanesTextsFontColor) SetG(v int32) *UpdateStreamingOutShrinkRequestPanesTextsFontColor {
	s.G = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanesTextsFontColor) SetR(v int32) *UpdateStreamingOutShrinkRequestPanesTextsFontColor {
	s.R = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanesTextsFontColor) Validate() error {
	return dara.Validate(s)
}

type UpdateStreamingOutShrinkRequestPanesWhiteboard struct {
	// example:
	//
	// default
	WhiteboardId *string `json:"WhiteboardId,omitempty" xml:"WhiteboardId,omitempty"`
}

func (s UpdateStreamingOutShrinkRequestPanesWhiteboard) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamingOutShrinkRequestPanesWhiteboard) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutShrinkRequestPanesWhiteboard) GetWhiteboardId() *string {
	return s.WhiteboardId
}

func (s *UpdateStreamingOutShrinkRequestPanesWhiteboard) SetWhiteboardId(v string) *UpdateStreamingOutShrinkRequestPanesWhiteboard {
	s.WhiteboardId = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestPanesWhiteboard) Validate() error {
	return dara.Validate(s)
}

type UpdateStreamingOutShrinkRequestRegionColor struct {
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

func (s UpdateStreamingOutShrinkRequestRegionColor) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamingOutShrinkRequestRegionColor) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutShrinkRequestRegionColor) GetB() *int32 {
	return s.B
}

func (s *UpdateStreamingOutShrinkRequestRegionColor) GetG() *int32 {
	return s.G
}

func (s *UpdateStreamingOutShrinkRequestRegionColor) GetR() *int32 {
	return s.R
}

func (s *UpdateStreamingOutShrinkRequestRegionColor) SetB(v int32) *UpdateStreamingOutShrinkRequestRegionColor {
	s.B = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestRegionColor) SetG(v int32) *UpdateStreamingOutShrinkRequestRegionColor {
	s.G = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestRegionColor) SetR(v int32) *UpdateStreamingOutShrinkRequestRegionColor {
	s.R = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestRegionColor) Validate() error {
	return dara.Validate(s)
}

type UpdateStreamingOutShrinkRequestTexts struct {
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
	BoxBorderw *int32                                        `json:"BoxBorderw,omitempty" xml:"BoxBorderw,omitempty"`
	BoxColor   *UpdateStreamingOutShrinkRequestTextsBoxColor `json:"BoxColor,omitempty" xml:"BoxColor,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Font      *int32                                         `json:"Font,omitempty" xml:"Font,omitempty"`
	FontColor *UpdateStreamingOutShrinkRequestTextsFontColor `json:"FontColor,omitempty" xml:"FontColor,omitempty" type:"Struct"`
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

func (s UpdateStreamingOutShrinkRequestTexts) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamingOutShrinkRequestTexts) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutShrinkRequestTexts) GetAlpha() *float64 {
	return s.Alpha
}

func (s *UpdateStreamingOutShrinkRequestTexts) GetBoxAlpha() *float64 {
	return s.BoxAlpha
}

func (s *UpdateStreamingOutShrinkRequestTexts) GetBoxBorderw() *int32 {
	return s.BoxBorderw
}

func (s *UpdateStreamingOutShrinkRequestTexts) GetBoxColor() *UpdateStreamingOutShrinkRequestTextsBoxColor {
	return s.BoxColor
}

func (s *UpdateStreamingOutShrinkRequestTexts) GetFont() *int32 {
	return s.Font
}

func (s *UpdateStreamingOutShrinkRequestTexts) GetFontColor() *UpdateStreamingOutShrinkRequestTextsFontColor {
	return s.FontColor
}

func (s *UpdateStreamingOutShrinkRequestTexts) GetFontSize() *int32 {
	return s.FontSize
}

func (s *UpdateStreamingOutShrinkRequestTexts) GetHasBox() *bool {
	return s.HasBox
}

func (s *UpdateStreamingOutShrinkRequestTexts) GetLayer() *int32 {
	return s.Layer
}

func (s *UpdateStreamingOutShrinkRequestTexts) GetTexture() *string {
	return s.Texture
}

func (s *UpdateStreamingOutShrinkRequestTexts) GetX() *float64 {
	return s.X
}

func (s *UpdateStreamingOutShrinkRequestTexts) GetY() *float64 {
	return s.Y
}

func (s *UpdateStreamingOutShrinkRequestTexts) SetAlpha(v float64) *UpdateStreamingOutShrinkRequestTexts {
	s.Alpha = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestTexts) SetBoxAlpha(v float64) *UpdateStreamingOutShrinkRequestTexts {
	s.BoxAlpha = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestTexts) SetBoxBorderw(v int32) *UpdateStreamingOutShrinkRequestTexts {
	s.BoxBorderw = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestTexts) SetBoxColor(v *UpdateStreamingOutShrinkRequestTextsBoxColor) *UpdateStreamingOutShrinkRequestTexts {
	s.BoxColor = v
	return s
}

func (s *UpdateStreamingOutShrinkRequestTexts) SetFont(v int32) *UpdateStreamingOutShrinkRequestTexts {
	s.Font = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestTexts) SetFontColor(v *UpdateStreamingOutShrinkRequestTextsFontColor) *UpdateStreamingOutShrinkRequestTexts {
	s.FontColor = v
	return s
}

func (s *UpdateStreamingOutShrinkRequestTexts) SetFontSize(v int32) *UpdateStreamingOutShrinkRequestTexts {
	s.FontSize = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestTexts) SetHasBox(v bool) *UpdateStreamingOutShrinkRequestTexts {
	s.HasBox = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestTexts) SetLayer(v int32) *UpdateStreamingOutShrinkRequestTexts {
	s.Layer = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestTexts) SetTexture(v string) *UpdateStreamingOutShrinkRequestTexts {
	s.Texture = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestTexts) SetX(v float64) *UpdateStreamingOutShrinkRequestTexts {
	s.X = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestTexts) SetY(v float64) *UpdateStreamingOutShrinkRequestTexts {
	s.Y = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestTexts) Validate() error {
	return dara.Validate(s)
}

type UpdateStreamingOutShrinkRequestTextsBoxColor struct {
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

func (s UpdateStreamingOutShrinkRequestTextsBoxColor) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamingOutShrinkRequestTextsBoxColor) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutShrinkRequestTextsBoxColor) GetB() *int32 {
	return s.B
}

func (s *UpdateStreamingOutShrinkRequestTextsBoxColor) GetG() *int32 {
	return s.G
}

func (s *UpdateStreamingOutShrinkRequestTextsBoxColor) GetR() *int32 {
	return s.R
}

func (s *UpdateStreamingOutShrinkRequestTextsBoxColor) SetB(v int32) *UpdateStreamingOutShrinkRequestTextsBoxColor {
	s.B = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestTextsBoxColor) SetG(v int32) *UpdateStreamingOutShrinkRequestTextsBoxColor {
	s.G = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestTextsBoxColor) SetR(v int32) *UpdateStreamingOutShrinkRequestTextsBoxColor {
	s.R = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestTextsBoxColor) Validate() error {
	return dara.Validate(s)
}

type UpdateStreamingOutShrinkRequestTextsFontColor struct {
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

func (s UpdateStreamingOutShrinkRequestTextsFontColor) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamingOutShrinkRequestTextsFontColor) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutShrinkRequestTextsFontColor) GetB() *int32 {
	return s.B
}

func (s *UpdateStreamingOutShrinkRequestTextsFontColor) GetG() *int32 {
	return s.G
}

func (s *UpdateStreamingOutShrinkRequestTextsFontColor) GetR() *int32 {
	return s.R
}

func (s *UpdateStreamingOutShrinkRequestTextsFontColor) SetB(v int32) *UpdateStreamingOutShrinkRequestTextsFontColor {
	s.B = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestTextsFontColor) SetG(v int32) *UpdateStreamingOutShrinkRequestTextsFontColor {
	s.G = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestTextsFontColor) SetR(v int32) *UpdateStreamingOutShrinkRequestTextsFontColor {
	s.R = &v
	return s
}

func (s *UpdateStreamingOutShrinkRequestTextsFontColor) Validate() error {
	return dara.Validate(s)
}
