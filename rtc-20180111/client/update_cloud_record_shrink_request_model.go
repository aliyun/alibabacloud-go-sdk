// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudRecordShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateCloudRecordShrinkRequest
	GetAppId() *string
	SetBackgrounds(v []*UpdateCloudRecordShrinkRequestBackgrounds) *UpdateCloudRecordShrinkRequest
	GetBackgrounds() []*UpdateCloudRecordShrinkRequestBackgrounds
	SetChannelId(v string) *UpdateCloudRecordShrinkRequest
	GetChannelId() *string
	SetClockWidgets(v []*UpdateCloudRecordShrinkRequestClockWidgets) *UpdateCloudRecordShrinkRequest
	GetClockWidgets() []*UpdateCloudRecordShrinkRequestClockWidgets
	SetImages(v []*UpdateCloudRecordShrinkRequestImages) *UpdateCloudRecordShrinkRequest
	GetImages() []*UpdateCloudRecordShrinkRequestImages
	SetLayoutSpecifiedUsersShrink(v string) *UpdateCloudRecordShrinkRequest
	GetLayoutSpecifiedUsersShrink() *string
	SetPanes(v []*UpdateCloudRecordShrinkRequestPanes) *UpdateCloudRecordShrinkRequest
	GetPanes() []*UpdateCloudRecordShrinkRequestPanes
	SetTaskId(v string) *UpdateCloudRecordShrinkRequest
	GetTaskId() *string
	SetTemplateId(v string) *UpdateCloudRecordShrinkRequest
	GetTemplateId() *string
	SetTexts(v []*UpdateCloudRecordShrinkRequestTexts) *UpdateCloudRecordShrinkRequest
	GetTexts() []*UpdateCloudRecordShrinkRequestTexts
}

type UpdateCloudRecordShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId       *string                                      `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Backgrounds []*UpdateCloudRecordShrinkRequestBackgrounds `json:"Backgrounds,omitempty" xml:"Backgrounds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId                  *string                                       `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ClockWidgets               []*UpdateCloudRecordShrinkRequestClockWidgets `json:"ClockWidgets,omitempty" xml:"ClockWidgets,omitempty" type:"Repeated"`
	Images                     []*UpdateCloudRecordShrinkRequestImages       `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	LayoutSpecifiedUsersShrink *string                                       `json:"LayoutSpecifiedUsers,omitempty" xml:"LayoutSpecifiedUsers,omitempty"`
	Panes                      []*UpdateCloudRecordShrinkRequestPanes        `json:"Panes,omitempty" xml:"Panes,omitempty" type:"Repeated"`
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
	TemplateId *string                                `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Texts      []*UpdateCloudRecordShrinkRequestTexts `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
}

func (s UpdateCloudRecordShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudRecordShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateCloudRecordShrinkRequest) GetBackgrounds() []*UpdateCloudRecordShrinkRequestBackgrounds {
	return s.Backgrounds
}

func (s *UpdateCloudRecordShrinkRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *UpdateCloudRecordShrinkRequest) GetClockWidgets() []*UpdateCloudRecordShrinkRequestClockWidgets {
	return s.ClockWidgets
}

func (s *UpdateCloudRecordShrinkRequest) GetImages() []*UpdateCloudRecordShrinkRequestImages {
	return s.Images
}

func (s *UpdateCloudRecordShrinkRequest) GetLayoutSpecifiedUsersShrink() *string {
	return s.LayoutSpecifiedUsersShrink
}

func (s *UpdateCloudRecordShrinkRequest) GetPanes() []*UpdateCloudRecordShrinkRequestPanes {
	return s.Panes
}

func (s *UpdateCloudRecordShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateCloudRecordShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateCloudRecordShrinkRequest) GetTexts() []*UpdateCloudRecordShrinkRequestTexts {
	return s.Texts
}

func (s *UpdateCloudRecordShrinkRequest) SetAppId(v string) *UpdateCloudRecordShrinkRequest {
	s.AppId = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequest) SetBackgrounds(v []*UpdateCloudRecordShrinkRequestBackgrounds) *UpdateCloudRecordShrinkRequest {
	s.Backgrounds = v
	return s
}

func (s *UpdateCloudRecordShrinkRequest) SetChannelId(v string) *UpdateCloudRecordShrinkRequest {
	s.ChannelId = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequest) SetClockWidgets(v []*UpdateCloudRecordShrinkRequestClockWidgets) *UpdateCloudRecordShrinkRequest {
	s.ClockWidgets = v
	return s
}

func (s *UpdateCloudRecordShrinkRequest) SetImages(v []*UpdateCloudRecordShrinkRequestImages) *UpdateCloudRecordShrinkRequest {
	s.Images = v
	return s
}

func (s *UpdateCloudRecordShrinkRequest) SetLayoutSpecifiedUsersShrink(v string) *UpdateCloudRecordShrinkRequest {
	s.LayoutSpecifiedUsersShrink = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequest) SetPanes(v []*UpdateCloudRecordShrinkRequestPanes) *UpdateCloudRecordShrinkRequest {
	s.Panes = v
	return s
}

func (s *UpdateCloudRecordShrinkRequest) SetTaskId(v string) *UpdateCloudRecordShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequest) SetTemplateId(v string) *UpdateCloudRecordShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequest) SetTexts(v []*UpdateCloudRecordShrinkRequestTexts) *UpdateCloudRecordShrinkRequest {
	s.Texts = v
	return s
}

func (s *UpdateCloudRecordShrinkRequest) Validate() error {
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

type UpdateCloudRecordShrinkRequestBackgrounds struct {
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

func (s UpdateCloudRecordShrinkRequestBackgrounds) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudRecordShrinkRequestBackgrounds) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordShrinkRequestBackgrounds) GetAlpha() *float64 {
	return s.Alpha
}

func (s *UpdateCloudRecordShrinkRequestBackgrounds) GetBackgroundCropMode() *int32 {
	return s.BackgroundCropMode
}

func (s *UpdateCloudRecordShrinkRequestBackgrounds) GetHeight() *float64 {
	return s.Height
}

func (s *UpdateCloudRecordShrinkRequestBackgrounds) GetLayer() *int32 {
	return s.Layer
}

func (s *UpdateCloudRecordShrinkRequestBackgrounds) GetUrl() *string {
	return s.Url
}

func (s *UpdateCloudRecordShrinkRequestBackgrounds) GetWidth() *float64 {
	return s.Width
}

func (s *UpdateCloudRecordShrinkRequestBackgrounds) GetX() *float64 {
	return s.X
}

func (s *UpdateCloudRecordShrinkRequestBackgrounds) GetY() *float64 {
	return s.Y
}

func (s *UpdateCloudRecordShrinkRequestBackgrounds) SetAlpha(v float64) *UpdateCloudRecordShrinkRequestBackgrounds {
	s.Alpha = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestBackgrounds) SetBackgroundCropMode(v int32) *UpdateCloudRecordShrinkRequestBackgrounds {
	s.BackgroundCropMode = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestBackgrounds) SetHeight(v float64) *UpdateCloudRecordShrinkRequestBackgrounds {
	s.Height = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestBackgrounds) SetLayer(v int32) *UpdateCloudRecordShrinkRequestBackgrounds {
	s.Layer = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestBackgrounds) SetUrl(v string) *UpdateCloudRecordShrinkRequestBackgrounds {
	s.Url = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestBackgrounds) SetWidth(v float64) *UpdateCloudRecordShrinkRequestBackgrounds {
	s.Width = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestBackgrounds) SetX(v float64) *UpdateCloudRecordShrinkRequestBackgrounds {
	s.X = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestBackgrounds) SetY(v float64) *UpdateCloudRecordShrinkRequestBackgrounds {
	s.Y = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestBackgrounds) Validate() error {
	return dara.Validate(s)
}

type UpdateCloudRecordShrinkRequestClockWidgets struct {
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
	BoxColor   *UpdateCloudRecordShrinkRequestClockWidgetsBoxColor `json:"BoxColor,omitempty" xml:"BoxColor,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Font      *int32                                               `json:"Font,omitempty" xml:"Font,omitempty"`
	FontColor *UpdateCloudRecordShrinkRequestClockWidgetsFontColor `json:"FontColor,omitempty" xml:"FontColor,omitempty" type:"Struct"`
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

func (s UpdateCloudRecordShrinkRequestClockWidgets) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudRecordShrinkRequestClockWidgets) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordShrinkRequestClockWidgets) GetAlpha() *float64 {
	return s.Alpha
}

func (s *UpdateCloudRecordShrinkRequestClockWidgets) GetBoxAlpha() *float64 {
	return s.BoxAlpha
}

func (s *UpdateCloudRecordShrinkRequestClockWidgets) GetBoxBorderw() *int32 {
	return s.BoxBorderw
}

func (s *UpdateCloudRecordShrinkRequestClockWidgets) GetBoxColor() *UpdateCloudRecordShrinkRequestClockWidgetsBoxColor {
	return s.BoxColor
}

func (s *UpdateCloudRecordShrinkRequestClockWidgets) GetFont() *int32 {
	return s.Font
}

func (s *UpdateCloudRecordShrinkRequestClockWidgets) GetFontColor() *UpdateCloudRecordShrinkRequestClockWidgetsFontColor {
	return s.FontColor
}

func (s *UpdateCloudRecordShrinkRequestClockWidgets) GetFontSize() *int32 {
	return s.FontSize
}

func (s *UpdateCloudRecordShrinkRequestClockWidgets) GetHasBox() *bool {
	return s.HasBox
}

func (s *UpdateCloudRecordShrinkRequestClockWidgets) GetLayer() *int32 {
	return s.Layer
}

func (s *UpdateCloudRecordShrinkRequestClockWidgets) GetX() *float64 {
	return s.X
}

func (s *UpdateCloudRecordShrinkRequestClockWidgets) GetY() *float64 {
	return s.Y
}

func (s *UpdateCloudRecordShrinkRequestClockWidgets) GetZone() *int32 {
	return s.Zone
}

func (s *UpdateCloudRecordShrinkRequestClockWidgets) SetAlpha(v float64) *UpdateCloudRecordShrinkRequestClockWidgets {
	s.Alpha = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestClockWidgets) SetBoxAlpha(v float64) *UpdateCloudRecordShrinkRequestClockWidgets {
	s.BoxAlpha = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestClockWidgets) SetBoxBorderw(v int32) *UpdateCloudRecordShrinkRequestClockWidgets {
	s.BoxBorderw = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestClockWidgets) SetBoxColor(v *UpdateCloudRecordShrinkRequestClockWidgetsBoxColor) *UpdateCloudRecordShrinkRequestClockWidgets {
	s.BoxColor = v
	return s
}

func (s *UpdateCloudRecordShrinkRequestClockWidgets) SetFont(v int32) *UpdateCloudRecordShrinkRequestClockWidgets {
	s.Font = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestClockWidgets) SetFontColor(v *UpdateCloudRecordShrinkRequestClockWidgetsFontColor) *UpdateCloudRecordShrinkRequestClockWidgets {
	s.FontColor = v
	return s
}

func (s *UpdateCloudRecordShrinkRequestClockWidgets) SetFontSize(v int32) *UpdateCloudRecordShrinkRequestClockWidgets {
	s.FontSize = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestClockWidgets) SetHasBox(v bool) *UpdateCloudRecordShrinkRequestClockWidgets {
	s.HasBox = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestClockWidgets) SetLayer(v int32) *UpdateCloudRecordShrinkRequestClockWidgets {
	s.Layer = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestClockWidgets) SetX(v float64) *UpdateCloudRecordShrinkRequestClockWidgets {
	s.X = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestClockWidgets) SetY(v float64) *UpdateCloudRecordShrinkRequestClockWidgets {
	s.Y = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestClockWidgets) SetZone(v int32) *UpdateCloudRecordShrinkRequestClockWidgets {
	s.Zone = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestClockWidgets) Validate() error {
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

type UpdateCloudRecordShrinkRequestClockWidgetsBoxColor struct {
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

func (s UpdateCloudRecordShrinkRequestClockWidgetsBoxColor) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudRecordShrinkRequestClockWidgetsBoxColor) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordShrinkRequestClockWidgetsBoxColor) GetB() *int32 {
	return s.B
}

func (s *UpdateCloudRecordShrinkRequestClockWidgetsBoxColor) GetG() *int32 {
	return s.G
}

func (s *UpdateCloudRecordShrinkRequestClockWidgetsBoxColor) GetR() *int32 {
	return s.R
}

func (s *UpdateCloudRecordShrinkRequestClockWidgetsBoxColor) SetB(v int32) *UpdateCloudRecordShrinkRequestClockWidgetsBoxColor {
	s.B = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestClockWidgetsBoxColor) SetG(v int32) *UpdateCloudRecordShrinkRequestClockWidgetsBoxColor {
	s.G = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestClockWidgetsBoxColor) SetR(v int32) *UpdateCloudRecordShrinkRequestClockWidgetsBoxColor {
	s.R = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestClockWidgetsBoxColor) Validate() error {
	return dara.Validate(s)
}

type UpdateCloudRecordShrinkRequestClockWidgetsFontColor struct {
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

func (s UpdateCloudRecordShrinkRequestClockWidgetsFontColor) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudRecordShrinkRequestClockWidgetsFontColor) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordShrinkRequestClockWidgetsFontColor) GetB() *int32 {
	return s.B
}

func (s *UpdateCloudRecordShrinkRequestClockWidgetsFontColor) GetG() *int32 {
	return s.G
}

func (s *UpdateCloudRecordShrinkRequestClockWidgetsFontColor) GetR() *int32 {
	return s.R
}

func (s *UpdateCloudRecordShrinkRequestClockWidgetsFontColor) SetB(v int32) *UpdateCloudRecordShrinkRequestClockWidgetsFontColor {
	s.B = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestClockWidgetsFontColor) SetG(v int32) *UpdateCloudRecordShrinkRequestClockWidgetsFontColor {
	s.G = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestClockWidgetsFontColor) SetR(v int32) *UpdateCloudRecordShrinkRequestClockWidgetsFontColor {
	s.R = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestClockWidgetsFontColor) Validate() error {
	return dara.Validate(s)
}

type UpdateCloudRecordShrinkRequestImages struct {
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

func (s UpdateCloudRecordShrinkRequestImages) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudRecordShrinkRequestImages) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordShrinkRequestImages) GetAlpha() *float64 {
	return s.Alpha
}

func (s *UpdateCloudRecordShrinkRequestImages) GetHeight() *float64 {
	return s.Height
}

func (s *UpdateCloudRecordShrinkRequestImages) GetImageCropMode() *int32 {
	return s.ImageCropMode
}

func (s *UpdateCloudRecordShrinkRequestImages) GetLayer() *int32 {
	return s.Layer
}

func (s *UpdateCloudRecordShrinkRequestImages) GetUrl() *string {
	return s.Url
}

func (s *UpdateCloudRecordShrinkRequestImages) GetWidth() *float64 {
	return s.Width
}

func (s *UpdateCloudRecordShrinkRequestImages) GetX() *float64 {
	return s.X
}

func (s *UpdateCloudRecordShrinkRequestImages) GetY() *float64 {
	return s.Y
}

func (s *UpdateCloudRecordShrinkRequestImages) SetAlpha(v float64) *UpdateCloudRecordShrinkRequestImages {
	s.Alpha = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestImages) SetHeight(v float64) *UpdateCloudRecordShrinkRequestImages {
	s.Height = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestImages) SetImageCropMode(v int32) *UpdateCloudRecordShrinkRequestImages {
	s.ImageCropMode = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestImages) SetLayer(v int32) *UpdateCloudRecordShrinkRequestImages {
	s.Layer = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestImages) SetUrl(v string) *UpdateCloudRecordShrinkRequestImages {
	s.Url = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestImages) SetWidth(v float64) *UpdateCloudRecordShrinkRequestImages {
	s.Width = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestImages) SetX(v float64) *UpdateCloudRecordShrinkRequestImages {
	s.X = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestImages) SetY(v float64) *UpdateCloudRecordShrinkRequestImages {
	s.Y = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestImages) Validate() error {
	return dara.Validate(s)
}

type UpdateCloudRecordShrinkRequestPanes struct {
	Backgrounds []*UpdateCloudRecordShrinkRequestPanesBackgrounds `json:"Backgrounds,omitempty" xml:"Backgrounds,omitempty" type:"Repeated"`
	Images      []*UpdateCloudRecordShrinkRequestPanesImages      `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	PaneCropMode *int32 `json:"PaneCropMode,omitempty" xml:"PaneCropMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	PaneId                    *int32 `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
	ReservePaneForOfflineUser *bool  `json:"ReservePaneForOfflineUser,omitempty" xml:"ReservePaneForOfflineUser,omitempty"`
	// example:
	//
	// 22
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// video
	SourceType *string                                     `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Texts      []*UpdateCloudRecordShrinkRequestPanesTexts `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
	// example:
	//
	// cameraFirst
	VideoOrder *string                                        `json:"VideoOrder,omitempty" xml:"VideoOrder,omitempty"`
	Whiteboard *UpdateCloudRecordShrinkRequestPanesWhiteboard `json:"Whiteboard,omitempty" xml:"Whiteboard,omitempty" type:"Struct"`
}

func (s UpdateCloudRecordShrinkRequestPanes) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudRecordShrinkRequestPanes) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordShrinkRequestPanes) GetBackgrounds() []*UpdateCloudRecordShrinkRequestPanesBackgrounds {
	return s.Backgrounds
}

func (s *UpdateCloudRecordShrinkRequestPanes) GetImages() []*UpdateCloudRecordShrinkRequestPanesImages {
	return s.Images
}

func (s *UpdateCloudRecordShrinkRequestPanes) GetPaneCropMode() *int32 {
	return s.PaneCropMode
}

func (s *UpdateCloudRecordShrinkRequestPanes) GetPaneId() *int32 {
	return s.PaneId
}

func (s *UpdateCloudRecordShrinkRequestPanes) GetReservePaneForOfflineUser() *bool {
	return s.ReservePaneForOfflineUser
}

func (s *UpdateCloudRecordShrinkRequestPanes) GetSource() *string {
	return s.Source
}

func (s *UpdateCloudRecordShrinkRequestPanes) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateCloudRecordShrinkRequestPanes) GetTexts() []*UpdateCloudRecordShrinkRequestPanesTexts {
	return s.Texts
}

func (s *UpdateCloudRecordShrinkRequestPanes) GetVideoOrder() *string {
	return s.VideoOrder
}

func (s *UpdateCloudRecordShrinkRequestPanes) GetWhiteboard() *UpdateCloudRecordShrinkRequestPanesWhiteboard {
	return s.Whiteboard
}

func (s *UpdateCloudRecordShrinkRequestPanes) SetBackgrounds(v []*UpdateCloudRecordShrinkRequestPanesBackgrounds) *UpdateCloudRecordShrinkRequestPanes {
	s.Backgrounds = v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanes) SetImages(v []*UpdateCloudRecordShrinkRequestPanesImages) *UpdateCloudRecordShrinkRequestPanes {
	s.Images = v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanes) SetPaneCropMode(v int32) *UpdateCloudRecordShrinkRequestPanes {
	s.PaneCropMode = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanes) SetPaneId(v int32) *UpdateCloudRecordShrinkRequestPanes {
	s.PaneId = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanes) SetReservePaneForOfflineUser(v bool) *UpdateCloudRecordShrinkRequestPanes {
	s.ReservePaneForOfflineUser = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanes) SetSource(v string) *UpdateCloudRecordShrinkRequestPanes {
	s.Source = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanes) SetSourceType(v string) *UpdateCloudRecordShrinkRequestPanes {
	s.SourceType = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanes) SetTexts(v []*UpdateCloudRecordShrinkRequestPanesTexts) *UpdateCloudRecordShrinkRequestPanes {
	s.Texts = v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanes) SetVideoOrder(v string) *UpdateCloudRecordShrinkRequestPanes {
	s.VideoOrder = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanes) SetWhiteboard(v *UpdateCloudRecordShrinkRequestPanesWhiteboard) *UpdateCloudRecordShrinkRequestPanes {
	s.Whiteboard = v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanes) Validate() error {
	if s.Backgrounds != nil {
		for _, item := range s.Backgrounds {
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
	if s.Texts != nil {
		for _, item := range s.Texts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Whiteboard != nil {
		if err := s.Whiteboard.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateCloudRecordShrinkRequestPanesBackgrounds struct {
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

func (s UpdateCloudRecordShrinkRequestPanesBackgrounds) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudRecordShrinkRequestPanesBackgrounds) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordShrinkRequestPanesBackgrounds) GetAlpha() *float64 {
	return s.Alpha
}

func (s *UpdateCloudRecordShrinkRequestPanesBackgrounds) GetDisplay() *string {
	return s.Display
}

func (s *UpdateCloudRecordShrinkRequestPanesBackgrounds) GetHeight() *float64 {
	return s.Height
}

func (s *UpdateCloudRecordShrinkRequestPanesBackgrounds) GetLayer() *int32 {
	return s.Layer
}

func (s *UpdateCloudRecordShrinkRequestPanesBackgrounds) GetPaneBackgroundCropMode() *int32 {
	return s.PaneBackgroundCropMode
}

func (s *UpdateCloudRecordShrinkRequestPanesBackgrounds) GetUrl() *string {
	return s.Url
}

func (s *UpdateCloudRecordShrinkRequestPanesBackgrounds) GetWidth() *float64 {
	return s.Width
}

func (s *UpdateCloudRecordShrinkRequestPanesBackgrounds) GetX() *float64 {
	return s.X
}

func (s *UpdateCloudRecordShrinkRequestPanesBackgrounds) GetY() *float64 {
	return s.Y
}

func (s *UpdateCloudRecordShrinkRequestPanesBackgrounds) SetAlpha(v float64) *UpdateCloudRecordShrinkRequestPanesBackgrounds {
	s.Alpha = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanesBackgrounds) SetDisplay(v string) *UpdateCloudRecordShrinkRequestPanesBackgrounds {
	s.Display = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanesBackgrounds) SetHeight(v float64) *UpdateCloudRecordShrinkRequestPanesBackgrounds {
	s.Height = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanesBackgrounds) SetLayer(v int32) *UpdateCloudRecordShrinkRequestPanesBackgrounds {
	s.Layer = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanesBackgrounds) SetPaneBackgroundCropMode(v int32) *UpdateCloudRecordShrinkRequestPanesBackgrounds {
	s.PaneBackgroundCropMode = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanesBackgrounds) SetUrl(v string) *UpdateCloudRecordShrinkRequestPanesBackgrounds {
	s.Url = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanesBackgrounds) SetWidth(v float64) *UpdateCloudRecordShrinkRequestPanesBackgrounds {
	s.Width = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanesBackgrounds) SetX(v float64) *UpdateCloudRecordShrinkRequestPanesBackgrounds {
	s.X = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanesBackgrounds) SetY(v float64) *UpdateCloudRecordShrinkRequestPanesBackgrounds {
	s.Y = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanesBackgrounds) Validate() error {
	return dara.Validate(s)
}

type UpdateCloudRecordShrinkRequestPanesImages struct {
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
	// 0.2
	Y *float64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s UpdateCloudRecordShrinkRequestPanesImages) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudRecordShrinkRequestPanesImages) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordShrinkRequestPanesImages) GetAlpha() *float64 {
	return s.Alpha
}

func (s *UpdateCloudRecordShrinkRequestPanesImages) GetDisplay() *string {
	return s.Display
}

func (s *UpdateCloudRecordShrinkRequestPanesImages) GetHeight() *float64 {
	return s.Height
}

func (s *UpdateCloudRecordShrinkRequestPanesImages) GetLayer() *int32 {
	return s.Layer
}

func (s *UpdateCloudRecordShrinkRequestPanesImages) GetPaneImageCropMode() *int32 {
	return s.PaneImageCropMode
}

func (s *UpdateCloudRecordShrinkRequestPanesImages) GetUrl() *string {
	return s.Url
}

func (s *UpdateCloudRecordShrinkRequestPanesImages) GetWidth() *float64 {
	return s.Width
}

func (s *UpdateCloudRecordShrinkRequestPanesImages) GetX() *float64 {
	return s.X
}

func (s *UpdateCloudRecordShrinkRequestPanesImages) GetY() *float64 {
	return s.Y
}

func (s *UpdateCloudRecordShrinkRequestPanesImages) SetAlpha(v float64) *UpdateCloudRecordShrinkRequestPanesImages {
	s.Alpha = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanesImages) SetDisplay(v string) *UpdateCloudRecordShrinkRequestPanesImages {
	s.Display = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanesImages) SetHeight(v float64) *UpdateCloudRecordShrinkRequestPanesImages {
	s.Height = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanesImages) SetLayer(v int32) *UpdateCloudRecordShrinkRequestPanesImages {
	s.Layer = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanesImages) SetPaneImageCropMode(v int32) *UpdateCloudRecordShrinkRequestPanesImages {
	s.PaneImageCropMode = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanesImages) SetUrl(v string) *UpdateCloudRecordShrinkRequestPanesImages {
	s.Url = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanesImages) SetWidth(v float64) *UpdateCloudRecordShrinkRequestPanesImages {
	s.Width = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanesImages) SetX(v float64) *UpdateCloudRecordShrinkRequestPanesImages {
	s.X = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanesImages) SetY(v float64) *UpdateCloudRecordShrinkRequestPanesImages {
	s.Y = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanesImages) Validate() error {
	return dara.Validate(s)
}

type UpdateCloudRecordShrinkRequestPanesTexts struct {
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
	BoxColor   *UpdateCloudRecordShrinkRequestPanesTextsBoxColor `json:"BoxColor,omitempty" xml:"BoxColor,omitempty" type:"Struct"`
	// example:
	//
	// backup
	Display *string `json:"Display,omitempty" xml:"Display,omitempty"`
	// example:
	//
	// 0
	Font      *int32                                             `json:"Font,omitempty" xml:"Font,omitempty"`
	FontColor *UpdateCloudRecordShrinkRequestPanesTextsFontColor `json:"FontColor,omitempty" xml:"FontColor,omitempty" type:"Struct"`
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

func (s UpdateCloudRecordShrinkRequestPanesTexts) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudRecordShrinkRequestPanesTexts) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordShrinkRequestPanesTexts) GetAlpha() *float64 {
	return s.Alpha
}

func (s *UpdateCloudRecordShrinkRequestPanesTexts) GetBoxAlpha() *float64 {
	return s.BoxAlpha
}

func (s *UpdateCloudRecordShrinkRequestPanesTexts) GetBoxBorderw() *int32 {
	return s.BoxBorderw
}

func (s *UpdateCloudRecordShrinkRequestPanesTexts) GetBoxColor() *UpdateCloudRecordShrinkRequestPanesTextsBoxColor {
	return s.BoxColor
}

func (s *UpdateCloudRecordShrinkRequestPanesTexts) GetDisplay() *string {
	return s.Display
}

func (s *UpdateCloudRecordShrinkRequestPanesTexts) GetFont() *int32 {
	return s.Font
}

func (s *UpdateCloudRecordShrinkRequestPanesTexts) GetFontColor() *UpdateCloudRecordShrinkRequestPanesTextsFontColor {
	return s.FontColor
}

func (s *UpdateCloudRecordShrinkRequestPanesTexts) GetFontSize() *int32 {
	return s.FontSize
}

func (s *UpdateCloudRecordShrinkRequestPanesTexts) GetHasBox() *bool {
	return s.HasBox
}

func (s *UpdateCloudRecordShrinkRequestPanesTexts) GetLayer() *int32 {
	return s.Layer
}

func (s *UpdateCloudRecordShrinkRequestPanesTexts) GetTexture() *string {
	return s.Texture
}

func (s *UpdateCloudRecordShrinkRequestPanesTexts) GetX() *float64 {
	return s.X
}

func (s *UpdateCloudRecordShrinkRequestPanesTexts) GetY() *float64 {
	return s.Y
}

func (s *UpdateCloudRecordShrinkRequestPanesTexts) SetAlpha(v float64) *UpdateCloudRecordShrinkRequestPanesTexts {
	s.Alpha = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanesTexts) SetBoxAlpha(v float64) *UpdateCloudRecordShrinkRequestPanesTexts {
	s.BoxAlpha = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanesTexts) SetBoxBorderw(v int32) *UpdateCloudRecordShrinkRequestPanesTexts {
	s.BoxBorderw = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanesTexts) SetBoxColor(v *UpdateCloudRecordShrinkRequestPanesTextsBoxColor) *UpdateCloudRecordShrinkRequestPanesTexts {
	s.BoxColor = v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanesTexts) SetDisplay(v string) *UpdateCloudRecordShrinkRequestPanesTexts {
	s.Display = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanesTexts) SetFont(v int32) *UpdateCloudRecordShrinkRequestPanesTexts {
	s.Font = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanesTexts) SetFontColor(v *UpdateCloudRecordShrinkRequestPanesTextsFontColor) *UpdateCloudRecordShrinkRequestPanesTexts {
	s.FontColor = v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanesTexts) SetFontSize(v int32) *UpdateCloudRecordShrinkRequestPanesTexts {
	s.FontSize = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanesTexts) SetHasBox(v bool) *UpdateCloudRecordShrinkRequestPanesTexts {
	s.HasBox = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanesTexts) SetLayer(v int32) *UpdateCloudRecordShrinkRequestPanesTexts {
	s.Layer = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanesTexts) SetTexture(v string) *UpdateCloudRecordShrinkRequestPanesTexts {
	s.Texture = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanesTexts) SetX(v float64) *UpdateCloudRecordShrinkRequestPanesTexts {
	s.X = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanesTexts) SetY(v float64) *UpdateCloudRecordShrinkRequestPanesTexts {
	s.Y = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanesTexts) Validate() error {
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

type UpdateCloudRecordShrinkRequestPanesTextsBoxColor struct {
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

func (s UpdateCloudRecordShrinkRequestPanesTextsBoxColor) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudRecordShrinkRequestPanesTextsBoxColor) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordShrinkRequestPanesTextsBoxColor) GetB() *int32 {
	return s.B
}

func (s *UpdateCloudRecordShrinkRequestPanesTextsBoxColor) GetG() *int32 {
	return s.G
}

func (s *UpdateCloudRecordShrinkRequestPanesTextsBoxColor) GetR() *int32 {
	return s.R
}

func (s *UpdateCloudRecordShrinkRequestPanesTextsBoxColor) SetB(v int32) *UpdateCloudRecordShrinkRequestPanesTextsBoxColor {
	s.B = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanesTextsBoxColor) SetG(v int32) *UpdateCloudRecordShrinkRequestPanesTextsBoxColor {
	s.G = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanesTextsBoxColor) SetR(v int32) *UpdateCloudRecordShrinkRequestPanesTextsBoxColor {
	s.R = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanesTextsBoxColor) Validate() error {
	return dara.Validate(s)
}

type UpdateCloudRecordShrinkRequestPanesTextsFontColor struct {
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

func (s UpdateCloudRecordShrinkRequestPanesTextsFontColor) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudRecordShrinkRequestPanesTextsFontColor) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordShrinkRequestPanesTextsFontColor) GetB() *int32 {
	return s.B
}

func (s *UpdateCloudRecordShrinkRequestPanesTextsFontColor) GetG() *int32 {
	return s.G
}

func (s *UpdateCloudRecordShrinkRequestPanesTextsFontColor) GetR() *int32 {
	return s.R
}

func (s *UpdateCloudRecordShrinkRequestPanesTextsFontColor) SetB(v int32) *UpdateCloudRecordShrinkRequestPanesTextsFontColor {
	s.B = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanesTextsFontColor) SetG(v int32) *UpdateCloudRecordShrinkRequestPanesTextsFontColor {
	s.G = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanesTextsFontColor) SetR(v int32) *UpdateCloudRecordShrinkRequestPanesTextsFontColor {
	s.R = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanesTextsFontColor) Validate() error {
	return dara.Validate(s)
}

type UpdateCloudRecordShrinkRequestPanesWhiteboard struct {
	// example:
	//
	// default
	WhiteboardId *string `json:"WhiteboardId,omitempty" xml:"WhiteboardId,omitempty"`
}

func (s UpdateCloudRecordShrinkRequestPanesWhiteboard) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudRecordShrinkRequestPanesWhiteboard) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordShrinkRequestPanesWhiteboard) GetWhiteboardId() *string {
	return s.WhiteboardId
}

func (s *UpdateCloudRecordShrinkRequestPanesWhiteboard) SetWhiteboardId(v string) *UpdateCloudRecordShrinkRequestPanesWhiteboard {
	s.WhiteboardId = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestPanesWhiteboard) Validate() error {
	return dara.Validate(s)
}

type UpdateCloudRecordShrinkRequestTexts struct {
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
	BoxColor   *UpdateCloudRecordShrinkRequestTextsBoxColor `json:"BoxColor,omitempty" xml:"BoxColor,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Font      *int32                                        `json:"Font,omitempty" xml:"Font,omitempty"`
	FontColor *UpdateCloudRecordShrinkRequestTextsFontColor `json:"FontColor,omitempty" xml:"FontColor,omitempty" type:"Struct"`
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

func (s UpdateCloudRecordShrinkRequestTexts) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudRecordShrinkRequestTexts) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordShrinkRequestTexts) GetAlpha() *float64 {
	return s.Alpha
}

func (s *UpdateCloudRecordShrinkRequestTexts) GetBoxAlpha() *float64 {
	return s.BoxAlpha
}

func (s *UpdateCloudRecordShrinkRequestTexts) GetBoxBorderw() *int32 {
	return s.BoxBorderw
}

func (s *UpdateCloudRecordShrinkRequestTexts) GetBoxColor() *UpdateCloudRecordShrinkRequestTextsBoxColor {
	return s.BoxColor
}

func (s *UpdateCloudRecordShrinkRequestTexts) GetFont() *int32 {
	return s.Font
}

func (s *UpdateCloudRecordShrinkRequestTexts) GetFontColor() *UpdateCloudRecordShrinkRequestTextsFontColor {
	return s.FontColor
}

func (s *UpdateCloudRecordShrinkRequestTexts) GetFontSize() *int32 {
	return s.FontSize
}

func (s *UpdateCloudRecordShrinkRequestTexts) GetHasBox() *bool {
	return s.HasBox
}

func (s *UpdateCloudRecordShrinkRequestTexts) GetLayer() *int32 {
	return s.Layer
}

func (s *UpdateCloudRecordShrinkRequestTexts) GetTexture() *string {
	return s.Texture
}

func (s *UpdateCloudRecordShrinkRequestTexts) GetX() *float64 {
	return s.X
}

func (s *UpdateCloudRecordShrinkRequestTexts) GetY() *float64 {
	return s.Y
}

func (s *UpdateCloudRecordShrinkRequestTexts) SetAlpha(v float64) *UpdateCloudRecordShrinkRequestTexts {
	s.Alpha = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestTexts) SetBoxAlpha(v float64) *UpdateCloudRecordShrinkRequestTexts {
	s.BoxAlpha = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestTexts) SetBoxBorderw(v int32) *UpdateCloudRecordShrinkRequestTexts {
	s.BoxBorderw = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestTexts) SetBoxColor(v *UpdateCloudRecordShrinkRequestTextsBoxColor) *UpdateCloudRecordShrinkRequestTexts {
	s.BoxColor = v
	return s
}

func (s *UpdateCloudRecordShrinkRequestTexts) SetFont(v int32) *UpdateCloudRecordShrinkRequestTexts {
	s.Font = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestTexts) SetFontColor(v *UpdateCloudRecordShrinkRequestTextsFontColor) *UpdateCloudRecordShrinkRequestTexts {
	s.FontColor = v
	return s
}

func (s *UpdateCloudRecordShrinkRequestTexts) SetFontSize(v int32) *UpdateCloudRecordShrinkRequestTexts {
	s.FontSize = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestTexts) SetHasBox(v bool) *UpdateCloudRecordShrinkRequestTexts {
	s.HasBox = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestTexts) SetLayer(v int32) *UpdateCloudRecordShrinkRequestTexts {
	s.Layer = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestTexts) SetTexture(v string) *UpdateCloudRecordShrinkRequestTexts {
	s.Texture = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestTexts) SetX(v float64) *UpdateCloudRecordShrinkRequestTexts {
	s.X = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestTexts) SetY(v float64) *UpdateCloudRecordShrinkRequestTexts {
	s.Y = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestTexts) Validate() error {
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

type UpdateCloudRecordShrinkRequestTextsBoxColor struct {
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

func (s UpdateCloudRecordShrinkRequestTextsBoxColor) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudRecordShrinkRequestTextsBoxColor) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordShrinkRequestTextsBoxColor) GetB() *int32 {
	return s.B
}

func (s *UpdateCloudRecordShrinkRequestTextsBoxColor) GetG() *int32 {
	return s.G
}

func (s *UpdateCloudRecordShrinkRequestTextsBoxColor) GetR() *int32 {
	return s.R
}

func (s *UpdateCloudRecordShrinkRequestTextsBoxColor) SetB(v int32) *UpdateCloudRecordShrinkRequestTextsBoxColor {
	s.B = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestTextsBoxColor) SetG(v int32) *UpdateCloudRecordShrinkRequestTextsBoxColor {
	s.G = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestTextsBoxColor) SetR(v int32) *UpdateCloudRecordShrinkRequestTextsBoxColor {
	s.R = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestTextsBoxColor) Validate() error {
	return dara.Validate(s)
}

type UpdateCloudRecordShrinkRequestTextsFontColor struct {
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

func (s UpdateCloudRecordShrinkRequestTextsFontColor) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudRecordShrinkRequestTextsFontColor) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordShrinkRequestTextsFontColor) GetB() *int32 {
	return s.B
}

func (s *UpdateCloudRecordShrinkRequestTextsFontColor) GetG() *int32 {
	return s.G
}

func (s *UpdateCloudRecordShrinkRequestTextsFontColor) GetR() *int32 {
	return s.R
}

func (s *UpdateCloudRecordShrinkRequestTextsFontColor) SetB(v int32) *UpdateCloudRecordShrinkRequestTextsFontColor {
	s.B = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestTextsFontColor) SetG(v int32) *UpdateCloudRecordShrinkRequestTextsFontColor {
	s.G = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestTextsFontColor) SetR(v int32) *UpdateCloudRecordShrinkRequestTextsFontColor {
	s.R = &v
	return s
}

func (s *UpdateCloudRecordShrinkRequestTextsFontColor) Validate() error {
	return dara.Validate(s)
}
