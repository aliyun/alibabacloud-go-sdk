// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyViewLayoutShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ModifyViewLayoutShrinkRequest
	GetAppId() *string
	SetBackgrounds(v []*ModifyViewLayoutShrinkRequestBackgrounds) *ModifyViewLayoutShrinkRequest
	GetBackgrounds() []*ModifyViewLayoutShrinkRequestBackgrounds
	SetChannelId(v string) *ModifyViewLayoutShrinkRequest
	GetChannelId() *string
	SetClockWidgets(v []*ModifyViewLayoutShrinkRequestClockWidgets) *ModifyViewLayoutShrinkRequest
	GetClockWidgets() []*ModifyViewLayoutShrinkRequestClockWidgets
	SetImages(v []*ModifyViewLayoutShrinkRequestImages) *ModifyViewLayoutShrinkRequest
	GetImages() []*ModifyViewLayoutShrinkRequestImages
	SetLayoutSpecifiedUsersShrink(v string) *ModifyViewLayoutShrinkRequest
	GetLayoutSpecifiedUsersShrink() *string
	SetPanes(v []*ModifyViewLayoutShrinkRequestPanes) *ModifyViewLayoutShrinkRequest
	GetPanes() []*ModifyViewLayoutShrinkRequestPanes
	SetTaskId(v string) *ModifyViewLayoutShrinkRequest
	GetTaskId() *string
	SetTemplateId(v string) *ModifyViewLayoutShrinkRequest
	GetTemplateId() *string
	SetTexts(v []*ModifyViewLayoutShrinkRequestTexts) *ModifyViewLayoutShrinkRequest
	GetTexts() []*ModifyViewLayoutShrinkRequestTexts
}

type ModifyViewLayoutShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId       *string                                     `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Backgrounds []*ModifyViewLayoutShrinkRequestBackgrounds `json:"Backgrounds,omitempty" xml:"Backgrounds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId                  *string                                      `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ClockWidgets               []*ModifyViewLayoutShrinkRequestClockWidgets `json:"ClockWidgets,omitempty" xml:"ClockWidgets,omitempty" type:"Repeated"`
	Images                     []*ModifyViewLayoutShrinkRequestImages       `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	LayoutSpecifiedUsersShrink *string                                      `json:"LayoutSpecifiedUsers,omitempty" xml:"LayoutSpecifiedUsers,omitempty"`
	Panes                      []*ModifyViewLayoutShrinkRequestPanes        `json:"Panes,omitempty" xml:"Panes,omitempty" type:"Repeated"`
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
	TemplateId *string                               `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Texts      []*ModifyViewLayoutShrinkRequestTexts `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
}

func (s ModifyViewLayoutShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyViewLayoutShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyViewLayoutShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *ModifyViewLayoutShrinkRequest) GetBackgrounds() []*ModifyViewLayoutShrinkRequestBackgrounds {
	return s.Backgrounds
}

func (s *ModifyViewLayoutShrinkRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *ModifyViewLayoutShrinkRequest) GetClockWidgets() []*ModifyViewLayoutShrinkRequestClockWidgets {
	return s.ClockWidgets
}

func (s *ModifyViewLayoutShrinkRequest) GetImages() []*ModifyViewLayoutShrinkRequestImages {
	return s.Images
}

func (s *ModifyViewLayoutShrinkRequest) GetLayoutSpecifiedUsersShrink() *string {
	return s.LayoutSpecifiedUsersShrink
}

func (s *ModifyViewLayoutShrinkRequest) GetPanes() []*ModifyViewLayoutShrinkRequestPanes {
	return s.Panes
}

func (s *ModifyViewLayoutShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ModifyViewLayoutShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ModifyViewLayoutShrinkRequest) GetTexts() []*ModifyViewLayoutShrinkRequestTexts {
	return s.Texts
}

func (s *ModifyViewLayoutShrinkRequest) SetAppId(v string) *ModifyViewLayoutShrinkRequest {
	s.AppId = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequest) SetBackgrounds(v []*ModifyViewLayoutShrinkRequestBackgrounds) *ModifyViewLayoutShrinkRequest {
	s.Backgrounds = v
	return s
}

func (s *ModifyViewLayoutShrinkRequest) SetChannelId(v string) *ModifyViewLayoutShrinkRequest {
	s.ChannelId = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequest) SetClockWidgets(v []*ModifyViewLayoutShrinkRequestClockWidgets) *ModifyViewLayoutShrinkRequest {
	s.ClockWidgets = v
	return s
}

func (s *ModifyViewLayoutShrinkRequest) SetImages(v []*ModifyViewLayoutShrinkRequestImages) *ModifyViewLayoutShrinkRequest {
	s.Images = v
	return s
}

func (s *ModifyViewLayoutShrinkRequest) SetLayoutSpecifiedUsersShrink(v string) *ModifyViewLayoutShrinkRequest {
	s.LayoutSpecifiedUsersShrink = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequest) SetPanes(v []*ModifyViewLayoutShrinkRequestPanes) *ModifyViewLayoutShrinkRequest {
	s.Panes = v
	return s
}

func (s *ModifyViewLayoutShrinkRequest) SetTaskId(v string) *ModifyViewLayoutShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequest) SetTemplateId(v string) *ModifyViewLayoutShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequest) SetTexts(v []*ModifyViewLayoutShrinkRequestTexts) *ModifyViewLayoutShrinkRequest {
	s.Texts = v
	return s
}

func (s *ModifyViewLayoutShrinkRequest) Validate() error {
	if s.Backgrounds != nil {
		for _, item := range s.Backgrounds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ClockWidgets != nil {
		for _, item := range s.ClockWidgets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Images != nil {
		for _, item := range s.Images {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Panes != nil {
		for _, item := range s.Panes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Texts != nil {
		for _, item := range s.Texts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyViewLayoutShrinkRequestBackgrounds struct {
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

func (s ModifyViewLayoutShrinkRequestBackgrounds) String() string {
	return dara.Prettify(s)
}

func (s ModifyViewLayoutShrinkRequestBackgrounds) GoString() string {
	return s.String()
}

func (s *ModifyViewLayoutShrinkRequestBackgrounds) GetAlpha() *float64 {
	return s.Alpha
}

func (s *ModifyViewLayoutShrinkRequestBackgrounds) GetBackgroundCropMode() *int32 {
	return s.BackgroundCropMode
}

func (s *ModifyViewLayoutShrinkRequestBackgrounds) GetHeight() *float64 {
	return s.Height
}

func (s *ModifyViewLayoutShrinkRequestBackgrounds) GetLayer() *int32 {
	return s.Layer
}

func (s *ModifyViewLayoutShrinkRequestBackgrounds) GetUrl() *string {
	return s.Url
}

func (s *ModifyViewLayoutShrinkRequestBackgrounds) GetWidth() *float64 {
	return s.Width
}

func (s *ModifyViewLayoutShrinkRequestBackgrounds) GetX() *float64 {
	return s.X
}

func (s *ModifyViewLayoutShrinkRequestBackgrounds) GetY() *float64 {
	return s.Y
}

func (s *ModifyViewLayoutShrinkRequestBackgrounds) SetAlpha(v float64) *ModifyViewLayoutShrinkRequestBackgrounds {
	s.Alpha = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestBackgrounds) SetBackgroundCropMode(v int32) *ModifyViewLayoutShrinkRequestBackgrounds {
	s.BackgroundCropMode = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestBackgrounds) SetHeight(v float64) *ModifyViewLayoutShrinkRequestBackgrounds {
	s.Height = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestBackgrounds) SetLayer(v int32) *ModifyViewLayoutShrinkRequestBackgrounds {
	s.Layer = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestBackgrounds) SetUrl(v string) *ModifyViewLayoutShrinkRequestBackgrounds {
	s.Url = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestBackgrounds) SetWidth(v float64) *ModifyViewLayoutShrinkRequestBackgrounds {
	s.Width = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestBackgrounds) SetX(v float64) *ModifyViewLayoutShrinkRequestBackgrounds {
	s.X = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestBackgrounds) SetY(v float64) *ModifyViewLayoutShrinkRequestBackgrounds {
	s.Y = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestBackgrounds) Validate() error {
	return dara.Validate(s)
}

type ModifyViewLayoutShrinkRequestClockWidgets struct {
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
	BoxColor   *ModifyViewLayoutShrinkRequestClockWidgetsBoxColor `json:"BoxColor,omitempty" xml:"BoxColor,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Font      *int32                                              `json:"Font,omitempty" xml:"Font,omitempty"`
	FontColor *ModifyViewLayoutShrinkRequestClockWidgetsFontColor `json:"FontColor,omitempty" xml:"FontColor,omitempty" type:"Struct"`
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

func (s ModifyViewLayoutShrinkRequestClockWidgets) String() string {
	return dara.Prettify(s)
}

func (s ModifyViewLayoutShrinkRequestClockWidgets) GoString() string {
	return s.String()
}

func (s *ModifyViewLayoutShrinkRequestClockWidgets) GetAlpha() *float64 {
	return s.Alpha
}

func (s *ModifyViewLayoutShrinkRequestClockWidgets) GetBoxAlpha() *float64 {
	return s.BoxAlpha
}

func (s *ModifyViewLayoutShrinkRequestClockWidgets) GetBoxBorderw() *int32 {
	return s.BoxBorderw
}

func (s *ModifyViewLayoutShrinkRequestClockWidgets) GetBoxColor() *ModifyViewLayoutShrinkRequestClockWidgetsBoxColor {
	return s.BoxColor
}

func (s *ModifyViewLayoutShrinkRequestClockWidgets) GetFont() *int32 {
	return s.Font
}

func (s *ModifyViewLayoutShrinkRequestClockWidgets) GetFontColor() *ModifyViewLayoutShrinkRequestClockWidgetsFontColor {
	return s.FontColor
}

func (s *ModifyViewLayoutShrinkRequestClockWidgets) GetFontSize() *int32 {
	return s.FontSize
}

func (s *ModifyViewLayoutShrinkRequestClockWidgets) GetHasBox() *bool {
	return s.HasBox
}

func (s *ModifyViewLayoutShrinkRequestClockWidgets) GetLayer() *int32 {
	return s.Layer
}

func (s *ModifyViewLayoutShrinkRequestClockWidgets) GetX() *float64 {
	return s.X
}

func (s *ModifyViewLayoutShrinkRequestClockWidgets) GetY() *float64 {
	return s.Y
}

func (s *ModifyViewLayoutShrinkRequestClockWidgets) GetZone() *int32 {
	return s.Zone
}

func (s *ModifyViewLayoutShrinkRequestClockWidgets) SetAlpha(v float64) *ModifyViewLayoutShrinkRequestClockWidgets {
	s.Alpha = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestClockWidgets) SetBoxAlpha(v float64) *ModifyViewLayoutShrinkRequestClockWidgets {
	s.BoxAlpha = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestClockWidgets) SetBoxBorderw(v int32) *ModifyViewLayoutShrinkRequestClockWidgets {
	s.BoxBorderw = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestClockWidgets) SetBoxColor(v *ModifyViewLayoutShrinkRequestClockWidgetsBoxColor) *ModifyViewLayoutShrinkRequestClockWidgets {
	s.BoxColor = v
	return s
}

func (s *ModifyViewLayoutShrinkRequestClockWidgets) SetFont(v int32) *ModifyViewLayoutShrinkRequestClockWidgets {
	s.Font = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestClockWidgets) SetFontColor(v *ModifyViewLayoutShrinkRequestClockWidgetsFontColor) *ModifyViewLayoutShrinkRequestClockWidgets {
	s.FontColor = v
	return s
}

func (s *ModifyViewLayoutShrinkRequestClockWidgets) SetFontSize(v int32) *ModifyViewLayoutShrinkRequestClockWidgets {
	s.FontSize = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestClockWidgets) SetHasBox(v bool) *ModifyViewLayoutShrinkRequestClockWidgets {
	s.HasBox = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestClockWidgets) SetLayer(v int32) *ModifyViewLayoutShrinkRequestClockWidgets {
	s.Layer = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestClockWidgets) SetX(v float64) *ModifyViewLayoutShrinkRequestClockWidgets {
	s.X = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestClockWidgets) SetY(v float64) *ModifyViewLayoutShrinkRequestClockWidgets {
	s.Y = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestClockWidgets) SetZone(v int32) *ModifyViewLayoutShrinkRequestClockWidgets {
	s.Zone = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestClockWidgets) Validate() error {
	if s.BoxColor != nil {
		if err := s.BoxColor.Validate(); err != nil {
			return err
		}
	}
	if s.FontColor != nil {
		if err := s.FontColor.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyViewLayoutShrinkRequestClockWidgetsBoxColor struct {
	// B。
	//
	// example:
	//
	// 255
	B *int32 `json:"B,omitempty" xml:"B,omitempty"`
	// G。
	//
	// example:
	//
	// 255
	G *int32 `json:"G,omitempty" xml:"G,omitempty"`
	// R。
	//
	// example:
	//
	// 255
	R *int32 `json:"R,omitempty" xml:"R,omitempty"`
}

func (s ModifyViewLayoutShrinkRequestClockWidgetsBoxColor) String() string {
	return dara.Prettify(s)
}

func (s ModifyViewLayoutShrinkRequestClockWidgetsBoxColor) GoString() string {
	return s.String()
}

func (s *ModifyViewLayoutShrinkRequestClockWidgetsBoxColor) GetB() *int32 {
	return s.B
}

func (s *ModifyViewLayoutShrinkRequestClockWidgetsBoxColor) GetG() *int32 {
	return s.G
}

func (s *ModifyViewLayoutShrinkRequestClockWidgetsBoxColor) GetR() *int32 {
	return s.R
}

func (s *ModifyViewLayoutShrinkRequestClockWidgetsBoxColor) SetB(v int32) *ModifyViewLayoutShrinkRequestClockWidgetsBoxColor {
	s.B = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestClockWidgetsBoxColor) SetG(v int32) *ModifyViewLayoutShrinkRequestClockWidgetsBoxColor {
	s.G = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestClockWidgetsBoxColor) SetR(v int32) *ModifyViewLayoutShrinkRequestClockWidgetsBoxColor {
	s.R = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestClockWidgetsBoxColor) Validate() error {
	return dara.Validate(s)
}

type ModifyViewLayoutShrinkRequestClockWidgetsFontColor struct {
	// B。
	//
	// example:
	//
	// 255
	B *int32 `json:"B,omitempty" xml:"B,omitempty"`
	// G。
	//
	// example:
	//
	// 255
	G *int32 `json:"G,omitempty" xml:"G,omitempty"`
	// R。
	//
	// example:
	//
	// 255
	R *int32 `json:"R,omitempty" xml:"R,omitempty"`
}

func (s ModifyViewLayoutShrinkRequestClockWidgetsFontColor) String() string {
	return dara.Prettify(s)
}

func (s ModifyViewLayoutShrinkRequestClockWidgetsFontColor) GoString() string {
	return s.String()
}

func (s *ModifyViewLayoutShrinkRequestClockWidgetsFontColor) GetB() *int32 {
	return s.B
}

func (s *ModifyViewLayoutShrinkRequestClockWidgetsFontColor) GetG() *int32 {
	return s.G
}

func (s *ModifyViewLayoutShrinkRequestClockWidgetsFontColor) GetR() *int32 {
	return s.R
}

func (s *ModifyViewLayoutShrinkRequestClockWidgetsFontColor) SetB(v int32) *ModifyViewLayoutShrinkRequestClockWidgetsFontColor {
	s.B = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestClockWidgetsFontColor) SetG(v int32) *ModifyViewLayoutShrinkRequestClockWidgetsFontColor {
	s.G = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestClockWidgetsFontColor) SetR(v int32) *ModifyViewLayoutShrinkRequestClockWidgetsFontColor {
	s.R = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestClockWidgetsFontColor) Validate() error {
	return dara.Validate(s)
}

type ModifyViewLayoutShrinkRequestImages struct {
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

func (s ModifyViewLayoutShrinkRequestImages) String() string {
	return dara.Prettify(s)
}

func (s ModifyViewLayoutShrinkRequestImages) GoString() string {
	return s.String()
}

func (s *ModifyViewLayoutShrinkRequestImages) GetAlpha() *float64 {
	return s.Alpha
}

func (s *ModifyViewLayoutShrinkRequestImages) GetHeight() *float64 {
	return s.Height
}

func (s *ModifyViewLayoutShrinkRequestImages) GetImageCropMode() *int32 {
	return s.ImageCropMode
}

func (s *ModifyViewLayoutShrinkRequestImages) GetLayer() *int32 {
	return s.Layer
}

func (s *ModifyViewLayoutShrinkRequestImages) GetUrl() *string {
	return s.Url
}

func (s *ModifyViewLayoutShrinkRequestImages) GetWidth() *float64 {
	return s.Width
}

func (s *ModifyViewLayoutShrinkRequestImages) GetX() *float64 {
	return s.X
}

func (s *ModifyViewLayoutShrinkRequestImages) GetY() *float64 {
	return s.Y
}

func (s *ModifyViewLayoutShrinkRequestImages) SetAlpha(v float64) *ModifyViewLayoutShrinkRequestImages {
	s.Alpha = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestImages) SetHeight(v float64) *ModifyViewLayoutShrinkRequestImages {
	s.Height = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestImages) SetImageCropMode(v int32) *ModifyViewLayoutShrinkRequestImages {
	s.ImageCropMode = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestImages) SetLayer(v int32) *ModifyViewLayoutShrinkRequestImages {
	s.Layer = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestImages) SetUrl(v string) *ModifyViewLayoutShrinkRequestImages {
	s.Url = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestImages) SetWidth(v float64) *ModifyViewLayoutShrinkRequestImages {
	s.Width = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestImages) SetX(v float64) *ModifyViewLayoutShrinkRequestImages {
	s.X = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestImages) SetY(v float64) *ModifyViewLayoutShrinkRequestImages {
	s.Y = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestImages) Validate() error {
	return dara.Validate(s)
}

type ModifyViewLayoutShrinkRequestPanes struct {
	Images []*ModifyViewLayoutShrinkRequestPanesImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	PaneCropMode *int32 `json:"PaneCropMode,omitempty" xml:"PaneCropMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	PaneId *string `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
	// example:
	//
	// 22
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// video
	SourceType *string                                    `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Texts      []*ModifyViewLayoutShrinkRequestPanesTexts `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
}

func (s ModifyViewLayoutShrinkRequestPanes) String() string {
	return dara.Prettify(s)
}

func (s ModifyViewLayoutShrinkRequestPanes) GoString() string {
	return s.String()
}

func (s *ModifyViewLayoutShrinkRequestPanes) GetImages() []*ModifyViewLayoutShrinkRequestPanesImages {
	return s.Images
}

func (s *ModifyViewLayoutShrinkRequestPanes) GetPaneCropMode() *int32 {
	return s.PaneCropMode
}

func (s *ModifyViewLayoutShrinkRequestPanes) GetPaneId() *string {
	return s.PaneId
}

func (s *ModifyViewLayoutShrinkRequestPanes) GetSource() *string {
	return s.Source
}

func (s *ModifyViewLayoutShrinkRequestPanes) GetSourceType() *string {
	return s.SourceType
}

func (s *ModifyViewLayoutShrinkRequestPanes) GetTexts() []*ModifyViewLayoutShrinkRequestPanesTexts {
	return s.Texts
}

func (s *ModifyViewLayoutShrinkRequestPanes) SetImages(v []*ModifyViewLayoutShrinkRequestPanesImages) *ModifyViewLayoutShrinkRequestPanes {
	s.Images = v
	return s
}

func (s *ModifyViewLayoutShrinkRequestPanes) SetPaneCropMode(v int32) *ModifyViewLayoutShrinkRequestPanes {
	s.PaneCropMode = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestPanes) SetPaneId(v string) *ModifyViewLayoutShrinkRequestPanes {
	s.PaneId = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestPanes) SetSource(v string) *ModifyViewLayoutShrinkRequestPanes {
	s.Source = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestPanes) SetSourceType(v string) *ModifyViewLayoutShrinkRequestPanes {
	s.SourceType = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestPanes) SetTexts(v []*ModifyViewLayoutShrinkRequestPanesTexts) *ModifyViewLayoutShrinkRequestPanes {
	s.Texts = v
	return s
}

func (s *ModifyViewLayoutShrinkRequestPanes) Validate() error {
	if s.Images != nil {
		for _, item := range s.Images {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Texts != nil {
		for _, item := range s.Texts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyViewLayoutShrinkRequestPanesImages struct {
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

func (s ModifyViewLayoutShrinkRequestPanesImages) String() string {
	return dara.Prettify(s)
}

func (s ModifyViewLayoutShrinkRequestPanesImages) GoString() string {
	return s.String()
}

func (s *ModifyViewLayoutShrinkRequestPanesImages) GetAlpha() *float64 {
	return s.Alpha
}

func (s *ModifyViewLayoutShrinkRequestPanesImages) GetHeight() *float64 {
	return s.Height
}

func (s *ModifyViewLayoutShrinkRequestPanesImages) GetLayer() *int32 {
	return s.Layer
}

func (s *ModifyViewLayoutShrinkRequestPanesImages) GetPaneImageCropMode() *int32 {
	return s.PaneImageCropMode
}

func (s *ModifyViewLayoutShrinkRequestPanesImages) GetUrl() *string {
	return s.Url
}

func (s *ModifyViewLayoutShrinkRequestPanesImages) GetWidth() *float64 {
	return s.Width
}

func (s *ModifyViewLayoutShrinkRequestPanesImages) GetX() *float64 {
	return s.X
}

func (s *ModifyViewLayoutShrinkRequestPanesImages) GetY() *float64 {
	return s.Y
}

func (s *ModifyViewLayoutShrinkRequestPanesImages) SetAlpha(v float64) *ModifyViewLayoutShrinkRequestPanesImages {
	s.Alpha = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestPanesImages) SetHeight(v float64) *ModifyViewLayoutShrinkRequestPanesImages {
	s.Height = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestPanesImages) SetLayer(v int32) *ModifyViewLayoutShrinkRequestPanesImages {
	s.Layer = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestPanesImages) SetPaneImageCropMode(v int32) *ModifyViewLayoutShrinkRequestPanesImages {
	s.PaneImageCropMode = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestPanesImages) SetUrl(v string) *ModifyViewLayoutShrinkRequestPanesImages {
	s.Url = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestPanesImages) SetWidth(v float64) *ModifyViewLayoutShrinkRequestPanesImages {
	s.Width = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestPanesImages) SetX(v float64) *ModifyViewLayoutShrinkRequestPanesImages {
	s.X = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestPanesImages) SetY(v float64) *ModifyViewLayoutShrinkRequestPanesImages {
	s.Y = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestPanesImages) Validate() error {
	return dara.Validate(s)
}

type ModifyViewLayoutShrinkRequestPanesTexts struct {
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
	BoxBorderw *int32                                           `json:"BoxBorderw,omitempty" xml:"BoxBorderw,omitempty"`
	BoxColor   *ModifyViewLayoutShrinkRequestPanesTextsBoxColor `json:"BoxColor,omitempty" xml:"BoxColor,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Font      *int32                                            `json:"Font,omitempty" xml:"Font,omitempty"`
	FontColor *ModifyViewLayoutShrinkRequestPanesTextsFontColor `json:"FontColor,omitempty" xml:"FontColor,omitempty" type:"Struct"`
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

func (s ModifyViewLayoutShrinkRequestPanesTexts) String() string {
	return dara.Prettify(s)
}

func (s ModifyViewLayoutShrinkRequestPanesTexts) GoString() string {
	return s.String()
}

func (s *ModifyViewLayoutShrinkRequestPanesTexts) GetAlpha() *float64 {
	return s.Alpha
}

func (s *ModifyViewLayoutShrinkRequestPanesTexts) GetBoxAlpha() *float64 {
	return s.BoxAlpha
}

func (s *ModifyViewLayoutShrinkRequestPanesTexts) GetBoxBorderw() *int32 {
	return s.BoxBorderw
}

func (s *ModifyViewLayoutShrinkRequestPanesTexts) GetBoxColor() *ModifyViewLayoutShrinkRequestPanesTextsBoxColor {
	return s.BoxColor
}

func (s *ModifyViewLayoutShrinkRequestPanesTexts) GetFont() *int32 {
	return s.Font
}

func (s *ModifyViewLayoutShrinkRequestPanesTexts) GetFontColor() *ModifyViewLayoutShrinkRequestPanesTextsFontColor {
	return s.FontColor
}

func (s *ModifyViewLayoutShrinkRequestPanesTexts) GetFontSize() *int32 {
	return s.FontSize
}

func (s *ModifyViewLayoutShrinkRequestPanesTexts) GetHasBox() *bool {
	return s.HasBox
}

func (s *ModifyViewLayoutShrinkRequestPanesTexts) GetLayer() *int32 {
	return s.Layer
}

func (s *ModifyViewLayoutShrinkRequestPanesTexts) GetTexture() *string {
	return s.Texture
}

func (s *ModifyViewLayoutShrinkRequestPanesTexts) GetX() *float64 {
	return s.X
}

func (s *ModifyViewLayoutShrinkRequestPanesTexts) GetY() *float64 {
	return s.Y
}

func (s *ModifyViewLayoutShrinkRequestPanesTexts) SetAlpha(v float64) *ModifyViewLayoutShrinkRequestPanesTexts {
	s.Alpha = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestPanesTexts) SetBoxAlpha(v float64) *ModifyViewLayoutShrinkRequestPanesTexts {
	s.BoxAlpha = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestPanesTexts) SetBoxBorderw(v int32) *ModifyViewLayoutShrinkRequestPanesTexts {
	s.BoxBorderw = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestPanesTexts) SetBoxColor(v *ModifyViewLayoutShrinkRequestPanesTextsBoxColor) *ModifyViewLayoutShrinkRequestPanesTexts {
	s.BoxColor = v
	return s
}

func (s *ModifyViewLayoutShrinkRequestPanesTexts) SetFont(v int32) *ModifyViewLayoutShrinkRequestPanesTexts {
	s.Font = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestPanesTexts) SetFontColor(v *ModifyViewLayoutShrinkRequestPanesTextsFontColor) *ModifyViewLayoutShrinkRequestPanesTexts {
	s.FontColor = v
	return s
}

func (s *ModifyViewLayoutShrinkRequestPanesTexts) SetFontSize(v int32) *ModifyViewLayoutShrinkRequestPanesTexts {
	s.FontSize = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestPanesTexts) SetHasBox(v bool) *ModifyViewLayoutShrinkRequestPanesTexts {
	s.HasBox = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestPanesTexts) SetLayer(v int32) *ModifyViewLayoutShrinkRequestPanesTexts {
	s.Layer = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestPanesTexts) SetTexture(v string) *ModifyViewLayoutShrinkRequestPanesTexts {
	s.Texture = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestPanesTexts) SetX(v float64) *ModifyViewLayoutShrinkRequestPanesTexts {
	s.X = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestPanesTexts) SetY(v float64) *ModifyViewLayoutShrinkRequestPanesTexts {
	s.Y = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestPanesTexts) Validate() error {
	if s.BoxColor != nil {
		if err := s.BoxColor.Validate(); err != nil {
			return err
		}
	}
	if s.FontColor != nil {
		if err := s.FontColor.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyViewLayoutShrinkRequestPanesTextsBoxColor struct {
	// B。
	//
	// example:
	//
	// 255
	B *int32 `json:"B,omitempty" xml:"B,omitempty"`
	// G。
	//
	// example:
	//
	// 255
	G *int32 `json:"G,omitempty" xml:"G,omitempty"`
	// R。
	//
	// example:
	//
	// 255
	R *int32 `json:"R,omitempty" xml:"R,omitempty"`
}

func (s ModifyViewLayoutShrinkRequestPanesTextsBoxColor) String() string {
	return dara.Prettify(s)
}

func (s ModifyViewLayoutShrinkRequestPanesTextsBoxColor) GoString() string {
	return s.String()
}

func (s *ModifyViewLayoutShrinkRequestPanesTextsBoxColor) GetB() *int32 {
	return s.B
}

func (s *ModifyViewLayoutShrinkRequestPanesTextsBoxColor) GetG() *int32 {
	return s.G
}

func (s *ModifyViewLayoutShrinkRequestPanesTextsBoxColor) GetR() *int32 {
	return s.R
}

func (s *ModifyViewLayoutShrinkRequestPanesTextsBoxColor) SetB(v int32) *ModifyViewLayoutShrinkRequestPanesTextsBoxColor {
	s.B = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestPanesTextsBoxColor) SetG(v int32) *ModifyViewLayoutShrinkRequestPanesTextsBoxColor {
	s.G = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestPanesTextsBoxColor) SetR(v int32) *ModifyViewLayoutShrinkRequestPanesTextsBoxColor {
	s.R = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestPanesTextsBoxColor) Validate() error {
	return dara.Validate(s)
}

type ModifyViewLayoutShrinkRequestPanesTextsFontColor struct {
	// B。
	//
	// example:
	//
	// 255
	B *int32 `json:"B,omitempty" xml:"B,omitempty"`
	// G。
	//
	// example:
	//
	// 255
	G *int32 `json:"G,omitempty" xml:"G,omitempty"`
	// R。
	//
	// example:
	//
	// 255
	R *int32 `json:"R,omitempty" xml:"R,omitempty"`
}

func (s ModifyViewLayoutShrinkRequestPanesTextsFontColor) String() string {
	return dara.Prettify(s)
}

func (s ModifyViewLayoutShrinkRequestPanesTextsFontColor) GoString() string {
	return s.String()
}

func (s *ModifyViewLayoutShrinkRequestPanesTextsFontColor) GetB() *int32 {
	return s.B
}

func (s *ModifyViewLayoutShrinkRequestPanesTextsFontColor) GetG() *int32 {
	return s.G
}

func (s *ModifyViewLayoutShrinkRequestPanesTextsFontColor) GetR() *int32 {
	return s.R
}

func (s *ModifyViewLayoutShrinkRequestPanesTextsFontColor) SetB(v int32) *ModifyViewLayoutShrinkRequestPanesTextsFontColor {
	s.B = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestPanesTextsFontColor) SetG(v int32) *ModifyViewLayoutShrinkRequestPanesTextsFontColor {
	s.G = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestPanesTextsFontColor) SetR(v int32) *ModifyViewLayoutShrinkRequestPanesTextsFontColor {
	s.R = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestPanesTextsFontColor) Validate() error {
	return dara.Validate(s)
}

type ModifyViewLayoutShrinkRequestTexts struct {
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
	BoxBorderw *int32                                      `json:"BoxBorderw,omitempty" xml:"BoxBorderw,omitempty"`
	BoxColor   *ModifyViewLayoutShrinkRequestTextsBoxColor `json:"BoxColor,omitempty" xml:"BoxColor,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Font      *int32                                       `json:"Font,omitempty" xml:"Font,omitempty"`
	FontColor *ModifyViewLayoutShrinkRequestTextsFontColor `json:"FontColor,omitempty" xml:"FontColor,omitempty" type:"Struct"`
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

func (s ModifyViewLayoutShrinkRequestTexts) String() string {
	return dara.Prettify(s)
}

func (s ModifyViewLayoutShrinkRequestTexts) GoString() string {
	return s.String()
}

func (s *ModifyViewLayoutShrinkRequestTexts) GetAlpha() *float64 {
	return s.Alpha
}

func (s *ModifyViewLayoutShrinkRequestTexts) GetBoxAlpha() *float64 {
	return s.BoxAlpha
}

func (s *ModifyViewLayoutShrinkRequestTexts) GetBoxBorderw() *int32 {
	return s.BoxBorderw
}

func (s *ModifyViewLayoutShrinkRequestTexts) GetBoxColor() *ModifyViewLayoutShrinkRequestTextsBoxColor {
	return s.BoxColor
}

func (s *ModifyViewLayoutShrinkRequestTexts) GetFont() *int32 {
	return s.Font
}

func (s *ModifyViewLayoutShrinkRequestTexts) GetFontColor() *ModifyViewLayoutShrinkRequestTextsFontColor {
	return s.FontColor
}

func (s *ModifyViewLayoutShrinkRequestTexts) GetFontSize() *int32 {
	return s.FontSize
}

func (s *ModifyViewLayoutShrinkRequestTexts) GetHasBox() *bool {
	return s.HasBox
}

func (s *ModifyViewLayoutShrinkRequestTexts) GetLayer() *int32 {
	return s.Layer
}

func (s *ModifyViewLayoutShrinkRequestTexts) GetTexture() *string {
	return s.Texture
}

func (s *ModifyViewLayoutShrinkRequestTexts) GetX() *float64 {
	return s.X
}

func (s *ModifyViewLayoutShrinkRequestTexts) GetY() *float64 {
	return s.Y
}

func (s *ModifyViewLayoutShrinkRequestTexts) SetAlpha(v float64) *ModifyViewLayoutShrinkRequestTexts {
	s.Alpha = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestTexts) SetBoxAlpha(v float64) *ModifyViewLayoutShrinkRequestTexts {
	s.BoxAlpha = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestTexts) SetBoxBorderw(v int32) *ModifyViewLayoutShrinkRequestTexts {
	s.BoxBorderw = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestTexts) SetBoxColor(v *ModifyViewLayoutShrinkRequestTextsBoxColor) *ModifyViewLayoutShrinkRequestTexts {
	s.BoxColor = v
	return s
}

func (s *ModifyViewLayoutShrinkRequestTexts) SetFont(v int32) *ModifyViewLayoutShrinkRequestTexts {
	s.Font = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestTexts) SetFontColor(v *ModifyViewLayoutShrinkRequestTextsFontColor) *ModifyViewLayoutShrinkRequestTexts {
	s.FontColor = v
	return s
}

func (s *ModifyViewLayoutShrinkRequestTexts) SetFontSize(v int32) *ModifyViewLayoutShrinkRequestTexts {
	s.FontSize = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestTexts) SetHasBox(v bool) *ModifyViewLayoutShrinkRequestTexts {
	s.HasBox = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestTexts) SetLayer(v int32) *ModifyViewLayoutShrinkRequestTexts {
	s.Layer = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestTexts) SetTexture(v string) *ModifyViewLayoutShrinkRequestTexts {
	s.Texture = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestTexts) SetX(v float64) *ModifyViewLayoutShrinkRequestTexts {
	s.X = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestTexts) SetY(v float64) *ModifyViewLayoutShrinkRequestTexts {
	s.Y = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestTexts) Validate() error {
	if s.BoxColor != nil {
		if err := s.BoxColor.Validate(); err != nil {
			return err
		}
	}
	if s.FontColor != nil {
		if err := s.FontColor.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyViewLayoutShrinkRequestTextsBoxColor struct {
	// B。
	//
	// example:
	//
	// 255
	B *int32 `json:"B,omitempty" xml:"B,omitempty"`
	// G。
	//
	// example:
	//
	// 255
	G *int32 `json:"G,omitempty" xml:"G,omitempty"`
	// R。
	//
	// example:
	//
	// 255
	R *int32 `json:"R,omitempty" xml:"R,omitempty"`
}

func (s ModifyViewLayoutShrinkRequestTextsBoxColor) String() string {
	return dara.Prettify(s)
}

func (s ModifyViewLayoutShrinkRequestTextsBoxColor) GoString() string {
	return s.String()
}

func (s *ModifyViewLayoutShrinkRequestTextsBoxColor) GetB() *int32 {
	return s.B
}

func (s *ModifyViewLayoutShrinkRequestTextsBoxColor) GetG() *int32 {
	return s.G
}

func (s *ModifyViewLayoutShrinkRequestTextsBoxColor) GetR() *int32 {
	return s.R
}

func (s *ModifyViewLayoutShrinkRequestTextsBoxColor) SetB(v int32) *ModifyViewLayoutShrinkRequestTextsBoxColor {
	s.B = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestTextsBoxColor) SetG(v int32) *ModifyViewLayoutShrinkRequestTextsBoxColor {
	s.G = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestTextsBoxColor) SetR(v int32) *ModifyViewLayoutShrinkRequestTextsBoxColor {
	s.R = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestTextsBoxColor) Validate() error {
	return dara.Validate(s)
}

type ModifyViewLayoutShrinkRequestTextsFontColor struct {
	// B。
	//
	// example:
	//
	// 255
	B *int32 `json:"B,omitempty" xml:"B,omitempty"`
	// G。
	//
	// example:
	//
	// 255
	G *int32 `json:"G,omitempty" xml:"G,omitempty"`
	// R。
	//
	// example:
	//
	// 255
	R *int32 `json:"R,omitempty" xml:"R,omitempty"`
}

func (s ModifyViewLayoutShrinkRequestTextsFontColor) String() string {
	return dara.Prettify(s)
}

func (s ModifyViewLayoutShrinkRequestTextsFontColor) GoString() string {
	return s.String()
}

func (s *ModifyViewLayoutShrinkRequestTextsFontColor) GetB() *int32 {
	return s.B
}

func (s *ModifyViewLayoutShrinkRequestTextsFontColor) GetG() *int32 {
	return s.G
}

func (s *ModifyViewLayoutShrinkRequestTextsFontColor) GetR() *int32 {
	return s.R
}

func (s *ModifyViewLayoutShrinkRequestTextsFontColor) SetB(v int32) *ModifyViewLayoutShrinkRequestTextsFontColor {
	s.B = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestTextsFontColor) SetG(v int32) *ModifyViewLayoutShrinkRequestTextsFontColor {
	s.G = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestTextsFontColor) SetR(v int32) *ModifyViewLayoutShrinkRequestTextsFontColor {
	s.R = &v
	return s
}

func (s *ModifyViewLayoutShrinkRequestTextsFontColor) Validate() error {
	return dara.Validate(s)
}
