// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyViewLayoutRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ModifyViewLayoutRequest
	GetAppId() *string
	SetBackgrounds(v []*ModifyViewLayoutRequestBackgrounds) *ModifyViewLayoutRequest
	GetBackgrounds() []*ModifyViewLayoutRequestBackgrounds
	SetChannelId(v string) *ModifyViewLayoutRequest
	GetChannelId() *string
	SetClockWidgets(v []*ModifyViewLayoutRequestClockWidgets) *ModifyViewLayoutRequest
	GetClockWidgets() []*ModifyViewLayoutRequestClockWidgets
	SetImages(v []*ModifyViewLayoutRequestImages) *ModifyViewLayoutRequest
	GetImages() []*ModifyViewLayoutRequestImages
	SetLayoutSpecifiedUsers(v *ModifyViewLayoutRequestLayoutSpecifiedUsers) *ModifyViewLayoutRequest
	GetLayoutSpecifiedUsers() *ModifyViewLayoutRequestLayoutSpecifiedUsers
	SetPanes(v []*ModifyViewLayoutRequestPanes) *ModifyViewLayoutRequest
	GetPanes() []*ModifyViewLayoutRequestPanes
	SetTaskId(v string) *ModifyViewLayoutRequest
	GetTaskId() *string
	SetTemplateId(v string) *ModifyViewLayoutRequest
	GetTemplateId() *string
	SetTexts(v []*ModifyViewLayoutRequestTexts) *ModifyViewLayoutRequest
	GetTexts() []*ModifyViewLayoutRequestTexts
}

type ModifyViewLayoutRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId       *string                               `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Backgrounds []*ModifyViewLayoutRequestBackgrounds `json:"Backgrounds,omitempty" xml:"Backgrounds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId            *string                                      `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ClockWidgets         []*ModifyViewLayoutRequestClockWidgets       `json:"ClockWidgets,omitempty" xml:"ClockWidgets,omitempty" type:"Repeated"`
	Images               []*ModifyViewLayoutRequestImages             `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	LayoutSpecifiedUsers *ModifyViewLayoutRequestLayoutSpecifiedUsers `json:"LayoutSpecifiedUsers,omitempty" xml:"LayoutSpecifiedUsers,omitempty" type:"Struct"`
	Panes                []*ModifyViewLayoutRequestPanes              `json:"Panes,omitempty" xml:"Panes,omitempty" type:"Repeated"`
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
	TemplateId *string                         `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Texts      []*ModifyViewLayoutRequestTexts `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
}

func (s ModifyViewLayoutRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyViewLayoutRequest) GoString() string {
	return s.String()
}

func (s *ModifyViewLayoutRequest) GetAppId() *string {
	return s.AppId
}

func (s *ModifyViewLayoutRequest) GetBackgrounds() []*ModifyViewLayoutRequestBackgrounds {
	return s.Backgrounds
}

func (s *ModifyViewLayoutRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *ModifyViewLayoutRequest) GetClockWidgets() []*ModifyViewLayoutRequestClockWidgets {
	return s.ClockWidgets
}

func (s *ModifyViewLayoutRequest) GetImages() []*ModifyViewLayoutRequestImages {
	return s.Images
}

func (s *ModifyViewLayoutRequest) GetLayoutSpecifiedUsers() *ModifyViewLayoutRequestLayoutSpecifiedUsers {
	return s.LayoutSpecifiedUsers
}

func (s *ModifyViewLayoutRequest) GetPanes() []*ModifyViewLayoutRequestPanes {
	return s.Panes
}

func (s *ModifyViewLayoutRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ModifyViewLayoutRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ModifyViewLayoutRequest) GetTexts() []*ModifyViewLayoutRequestTexts {
	return s.Texts
}

func (s *ModifyViewLayoutRequest) SetAppId(v string) *ModifyViewLayoutRequest {
	s.AppId = &v
	return s
}

func (s *ModifyViewLayoutRequest) SetBackgrounds(v []*ModifyViewLayoutRequestBackgrounds) *ModifyViewLayoutRequest {
	s.Backgrounds = v
	return s
}

func (s *ModifyViewLayoutRequest) SetChannelId(v string) *ModifyViewLayoutRequest {
	s.ChannelId = &v
	return s
}

func (s *ModifyViewLayoutRequest) SetClockWidgets(v []*ModifyViewLayoutRequestClockWidgets) *ModifyViewLayoutRequest {
	s.ClockWidgets = v
	return s
}

func (s *ModifyViewLayoutRequest) SetImages(v []*ModifyViewLayoutRequestImages) *ModifyViewLayoutRequest {
	s.Images = v
	return s
}

func (s *ModifyViewLayoutRequest) SetLayoutSpecifiedUsers(v *ModifyViewLayoutRequestLayoutSpecifiedUsers) *ModifyViewLayoutRequest {
	s.LayoutSpecifiedUsers = v
	return s
}

func (s *ModifyViewLayoutRequest) SetPanes(v []*ModifyViewLayoutRequestPanes) *ModifyViewLayoutRequest {
	s.Panes = v
	return s
}

func (s *ModifyViewLayoutRequest) SetTaskId(v string) *ModifyViewLayoutRequest {
	s.TaskId = &v
	return s
}

func (s *ModifyViewLayoutRequest) SetTemplateId(v string) *ModifyViewLayoutRequest {
	s.TemplateId = &v
	return s
}

func (s *ModifyViewLayoutRequest) SetTexts(v []*ModifyViewLayoutRequestTexts) *ModifyViewLayoutRequest {
	s.Texts = v
	return s
}

func (s *ModifyViewLayoutRequest) Validate() error {
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
	if s.LayoutSpecifiedUsers != nil {
		if err := s.LayoutSpecifiedUsers.Validate(); err != nil {
			return err
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

type ModifyViewLayoutRequestBackgrounds struct {
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

func (s ModifyViewLayoutRequestBackgrounds) String() string {
	return dara.Prettify(s)
}

func (s ModifyViewLayoutRequestBackgrounds) GoString() string {
	return s.String()
}

func (s *ModifyViewLayoutRequestBackgrounds) GetAlpha() *float64 {
	return s.Alpha
}

func (s *ModifyViewLayoutRequestBackgrounds) GetBackgroundCropMode() *int32 {
	return s.BackgroundCropMode
}

func (s *ModifyViewLayoutRequestBackgrounds) GetHeight() *float64 {
	return s.Height
}

func (s *ModifyViewLayoutRequestBackgrounds) GetLayer() *int32 {
	return s.Layer
}

func (s *ModifyViewLayoutRequestBackgrounds) GetUrl() *string {
	return s.Url
}

func (s *ModifyViewLayoutRequestBackgrounds) GetWidth() *float64 {
	return s.Width
}

func (s *ModifyViewLayoutRequestBackgrounds) GetX() *float64 {
	return s.X
}

func (s *ModifyViewLayoutRequestBackgrounds) GetY() *float64 {
	return s.Y
}

func (s *ModifyViewLayoutRequestBackgrounds) SetAlpha(v float64) *ModifyViewLayoutRequestBackgrounds {
	s.Alpha = &v
	return s
}

func (s *ModifyViewLayoutRequestBackgrounds) SetBackgroundCropMode(v int32) *ModifyViewLayoutRequestBackgrounds {
	s.BackgroundCropMode = &v
	return s
}

func (s *ModifyViewLayoutRequestBackgrounds) SetHeight(v float64) *ModifyViewLayoutRequestBackgrounds {
	s.Height = &v
	return s
}

func (s *ModifyViewLayoutRequestBackgrounds) SetLayer(v int32) *ModifyViewLayoutRequestBackgrounds {
	s.Layer = &v
	return s
}

func (s *ModifyViewLayoutRequestBackgrounds) SetUrl(v string) *ModifyViewLayoutRequestBackgrounds {
	s.Url = &v
	return s
}

func (s *ModifyViewLayoutRequestBackgrounds) SetWidth(v float64) *ModifyViewLayoutRequestBackgrounds {
	s.Width = &v
	return s
}

func (s *ModifyViewLayoutRequestBackgrounds) SetX(v float64) *ModifyViewLayoutRequestBackgrounds {
	s.X = &v
	return s
}

func (s *ModifyViewLayoutRequestBackgrounds) SetY(v float64) *ModifyViewLayoutRequestBackgrounds {
	s.Y = &v
	return s
}

func (s *ModifyViewLayoutRequestBackgrounds) Validate() error {
	return dara.Validate(s)
}

type ModifyViewLayoutRequestClockWidgets struct {
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
	BoxBorderw *int32                                       `json:"BoxBorderw,omitempty" xml:"BoxBorderw,omitempty"`
	BoxColor   *ModifyViewLayoutRequestClockWidgetsBoxColor `json:"BoxColor,omitempty" xml:"BoxColor,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Font      *int32                                        `json:"Font,omitempty" xml:"Font,omitempty"`
	FontColor *ModifyViewLayoutRequestClockWidgetsFontColor `json:"FontColor,omitempty" xml:"FontColor,omitempty" type:"Struct"`
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

func (s ModifyViewLayoutRequestClockWidgets) String() string {
	return dara.Prettify(s)
}

func (s ModifyViewLayoutRequestClockWidgets) GoString() string {
	return s.String()
}

func (s *ModifyViewLayoutRequestClockWidgets) GetAlpha() *float64 {
	return s.Alpha
}

func (s *ModifyViewLayoutRequestClockWidgets) GetBoxAlpha() *float64 {
	return s.BoxAlpha
}

func (s *ModifyViewLayoutRequestClockWidgets) GetBoxBorderw() *int32 {
	return s.BoxBorderw
}

func (s *ModifyViewLayoutRequestClockWidgets) GetBoxColor() *ModifyViewLayoutRequestClockWidgetsBoxColor {
	return s.BoxColor
}

func (s *ModifyViewLayoutRequestClockWidgets) GetFont() *int32 {
	return s.Font
}

func (s *ModifyViewLayoutRequestClockWidgets) GetFontColor() *ModifyViewLayoutRequestClockWidgetsFontColor {
	return s.FontColor
}

func (s *ModifyViewLayoutRequestClockWidgets) GetFontSize() *int32 {
	return s.FontSize
}

func (s *ModifyViewLayoutRequestClockWidgets) GetHasBox() *bool {
	return s.HasBox
}

func (s *ModifyViewLayoutRequestClockWidgets) GetLayer() *int32 {
	return s.Layer
}

func (s *ModifyViewLayoutRequestClockWidgets) GetX() *float64 {
	return s.X
}

func (s *ModifyViewLayoutRequestClockWidgets) GetY() *float64 {
	return s.Y
}

func (s *ModifyViewLayoutRequestClockWidgets) GetZone() *int32 {
	return s.Zone
}

func (s *ModifyViewLayoutRequestClockWidgets) SetAlpha(v float64) *ModifyViewLayoutRequestClockWidgets {
	s.Alpha = &v
	return s
}

func (s *ModifyViewLayoutRequestClockWidgets) SetBoxAlpha(v float64) *ModifyViewLayoutRequestClockWidgets {
	s.BoxAlpha = &v
	return s
}

func (s *ModifyViewLayoutRequestClockWidgets) SetBoxBorderw(v int32) *ModifyViewLayoutRequestClockWidgets {
	s.BoxBorderw = &v
	return s
}

func (s *ModifyViewLayoutRequestClockWidgets) SetBoxColor(v *ModifyViewLayoutRequestClockWidgetsBoxColor) *ModifyViewLayoutRequestClockWidgets {
	s.BoxColor = v
	return s
}

func (s *ModifyViewLayoutRequestClockWidgets) SetFont(v int32) *ModifyViewLayoutRequestClockWidgets {
	s.Font = &v
	return s
}

func (s *ModifyViewLayoutRequestClockWidgets) SetFontColor(v *ModifyViewLayoutRequestClockWidgetsFontColor) *ModifyViewLayoutRequestClockWidgets {
	s.FontColor = v
	return s
}

func (s *ModifyViewLayoutRequestClockWidgets) SetFontSize(v int32) *ModifyViewLayoutRequestClockWidgets {
	s.FontSize = &v
	return s
}

func (s *ModifyViewLayoutRequestClockWidgets) SetHasBox(v bool) *ModifyViewLayoutRequestClockWidgets {
	s.HasBox = &v
	return s
}

func (s *ModifyViewLayoutRequestClockWidgets) SetLayer(v int32) *ModifyViewLayoutRequestClockWidgets {
	s.Layer = &v
	return s
}

func (s *ModifyViewLayoutRequestClockWidgets) SetX(v float64) *ModifyViewLayoutRequestClockWidgets {
	s.X = &v
	return s
}

func (s *ModifyViewLayoutRequestClockWidgets) SetY(v float64) *ModifyViewLayoutRequestClockWidgets {
	s.Y = &v
	return s
}

func (s *ModifyViewLayoutRequestClockWidgets) SetZone(v int32) *ModifyViewLayoutRequestClockWidgets {
	s.Zone = &v
	return s
}

func (s *ModifyViewLayoutRequestClockWidgets) Validate() error {
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

type ModifyViewLayoutRequestClockWidgetsBoxColor struct {
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

func (s ModifyViewLayoutRequestClockWidgetsBoxColor) String() string {
	return dara.Prettify(s)
}

func (s ModifyViewLayoutRequestClockWidgetsBoxColor) GoString() string {
	return s.String()
}

func (s *ModifyViewLayoutRequestClockWidgetsBoxColor) GetB() *int32 {
	return s.B
}

func (s *ModifyViewLayoutRequestClockWidgetsBoxColor) GetG() *int32 {
	return s.G
}

func (s *ModifyViewLayoutRequestClockWidgetsBoxColor) GetR() *int32 {
	return s.R
}

func (s *ModifyViewLayoutRequestClockWidgetsBoxColor) SetB(v int32) *ModifyViewLayoutRequestClockWidgetsBoxColor {
	s.B = &v
	return s
}

func (s *ModifyViewLayoutRequestClockWidgetsBoxColor) SetG(v int32) *ModifyViewLayoutRequestClockWidgetsBoxColor {
	s.G = &v
	return s
}

func (s *ModifyViewLayoutRequestClockWidgetsBoxColor) SetR(v int32) *ModifyViewLayoutRequestClockWidgetsBoxColor {
	s.R = &v
	return s
}

func (s *ModifyViewLayoutRequestClockWidgetsBoxColor) Validate() error {
	return dara.Validate(s)
}

type ModifyViewLayoutRequestClockWidgetsFontColor struct {
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

func (s ModifyViewLayoutRequestClockWidgetsFontColor) String() string {
	return dara.Prettify(s)
}

func (s ModifyViewLayoutRequestClockWidgetsFontColor) GoString() string {
	return s.String()
}

func (s *ModifyViewLayoutRequestClockWidgetsFontColor) GetB() *int32 {
	return s.B
}

func (s *ModifyViewLayoutRequestClockWidgetsFontColor) GetG() *int32 {
	return s.G
}

func (s *ModifyViewLayoutRequestClockWidgetsFontColor) GetR() *int32 {
	return s.R
}

func (s *ModifyViewLayoutRequestClockWidgetsFontColor) SetB(v int32) *ModifyViewLayoutRequestClockWidgetsFontColor {
	s.B = &v
	return s
}

func (s *ModifyViewLayoutRequestClockWidgetsFontColor) SetG(v int32) *ModifyViewLayoutRequestClockWidgetsFontColor {
	s.G = &v
	return s
}

func (s *ModifyViewLayoutRequestClockWidgetsFontColor) SetR(v int32) *ModifyViewLayoutRequestClockWidgetsFontColor {
	s.R = &v
	return s
}

func (s *ModifyViewLayoutRequestClockWidgetsFontColor) Validate() error {
	return dara.Validate(s)
}

type ModifyViewLayoutRequestImages struct {
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

func (s ModifyViewLayoutRequestImages) String() string {
	return dara.Prettify(s)
}

func (s ModifyViewLayoutRequestImages) GoString() string {
	return s.String()
}

func (s *ModifyViewLayoutRequestImages) GetAlpha() *float64 {
	return s.Alpha
}

func (s *ModifyViewLayoutRequestImages) GetHeight() *float64 {
	return s.Height
}

func (s *ModifyViewLayoutRequestImages) GetImageCropMode() *int32 {
	return s.ImageCropMode
}

func (s *ModifyViewLayoutRequestImages) GetLayer() *int32 {
	return s.Layer
}

func (s *ModifyViewLayoutRequestImages) GetUrl() *string {
	return s.Url
}

func (s *ModifyViewLayoutRequestImages) GetWidth() *float64 {
	return s.Width
}

func (s *ModifyViewLayoutRequestImages) GetX() *float64 {
	return s.X
}

func (s *ModifyViewLayoutRequestImages) GetY() *float64 {
	return s.Y
}

func (s *ModifyViewLayoutRequestImages) SetAlpha(v float64) *ModifyViewLayoutRequestImages {
	s.Alpha = &v
	return s
}

func (s *ModifyViewLayoutRequestImages) SetHeight(v float64) *ModifyViewLayoutRequestImages {
	s.Height = &v
	return s
}

func (s *ModifyViewLayoutRequestImages) SetImageCropMode(v int32) *ModifyViewLayoutRequestImages {
	s.ImageCropMode = &v
	return s
}

func (s *ModifyViewLayoutRequestImages) SetLayer(v int32) *ModifyViewLayoutRequestImages {
	s.Layer = &v
	return s
}

func (s *ModifyViewLayoutRequestImages) SetUrl(v string) *ModifyViewLayoutRequestImages {
	s.Url = &v
	return s
}

func (s *ModifyViewLayoutRequestImages) SetWidth(v float64) *ModifyViewLayoutRequestImages {
	s.Width = &v
	return s
}

func (s *ModifyViewLayoutRequestImages) SetX(v float64) *ModifyViewLayoutRequestImages {
	s.X = &v
	return s
}

func (s *ModifyViewLayoutRequestImages) SetY(v float64) *ModifyViewLayoutRequestImages {
	s.Y = &v
	return s
}

func (s *ModifyViewLayoutRequestImages) Validate() error {
	return dara.Validate(s)
}

type ModifyViewLayoutRequestLayoutSpecifiedUsers struct {
	// This parameter is required.
	Ids []*string `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// white
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyViewLayoutRequestLayoutSpecifiedUsers) String() string {
	return dara.Prettify(s)
}

func (s ModifyViewLayoutRequestLayoutSpecifiedUsers) GoString() string {
	return s.String()
}

func (s *ModifyViewLayoutRequestLayoutSpecifiedUsers) GetIds() []*string {
	return s.Ids
}

func (s *ModifyViewLayoutRequestLayoutSpecifiedUsers) GetType() *string {
	return s.Type
}

func (s *ModifyViewLayoutRequestLayoutSpecifiedUsers) SetIds(v []*string) *ModifyViewLayoutRequestLayoutSpecifiedUsers {
	s.Ids = v
	return s
}

func (s *ModifyViewLayoutRequestLayoutSpecifiedUsers) SetType(v string) *ModifyViewLayoutRequestLayoutSpecifiedUsers {
	s.Type = &v
	return s
}

func (s *ModifyViewLayoutRequestLayoutSpecifiedUsers) Validate() error {
	return dara.Validate(s)
}

type ModifyViewLayoutRequestPanes struct {
	Images []*ModifyViewLayoutRequestPanesImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
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
	SourceType *string                              `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Texts      []*ModifyViewLayoutRequestPanesTexts `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
}

func (s ModifyViewLayoutRequestPanes) String() string {
	return dara.Prettify(s)
}

func (s ModifyViewLayoutRequestPanes) GoString() string {
	return s.String()
}

func (s *ModifyViewLayoutRequestPanes) GetImages() []*ModifyViewLayoutRequestPanesImages {
	return s.Images
}

func (s *ModifyViewLayoutRequestPanes) GetPaneCropMode() *int32 {
	return s.PaneCropMode
}

func (s *ModifyViewLayoutRequestPanes) GetPaneId() *string {
	return s.PaneId
}

func (s *ModifyViewLayoutRequestPanes) GetSource() *string {
	return s.Source
}

func (s *ModifyViewLayoutRequestPanes) GetSourceType() *string {
	return s.SourceType
}

func (s *ModifyViewLayoutRequestPanes) GetTexts() []*ModifyViewLayoutRequestPanesTexts {
	return s.Texts
}

func (s *ModifyViewLayoutRequestPanes) SetImages(v []*ModifyViewLayoutRequestPanesImages) *ModifyViewLayoutRequestPanes {
	s.Images = v
	return s
}

func (s *ModifyViewLayoutRequestPanes) SetPaneCropMode(v int32) *ModifyViewLayoutRequestPanes {
	s.PaneCropMode = &v
	return s
}

func (s *ModifyViewLayoutRequestPanes) SetPaneId(v string) *ModifyViewLayoutRequestPanes {
	s.PaneId = &v
	return s
}

func (s *ModifyViewLayoutRequestPanes) SetSource(v string) *ModifyViewLayoutRequestPanes {
	s.Source = &v
	return s
}

func (s *ModifyViewLayoutRequestPanes) SetSourceType(v string) *ModifyViewLayoutRequestPanes {
	s.SourceType = &v
	return s
}

func (s *ModifyViewLayoutRequestPanes) SetTexts(v []*ModifyViewLayoutRequestPanesTexts) *ModifyViewLayoutRequestPanes {
	s.Texts = v
	return s
}

func (s *ModifyViewLayoutRequestPanes) Validate() error {
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

type ModifyViewLayoutRequestPanesImages struct {
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

func (s ModifyViewLayoutRequestPanesImages) String() string {
	return dara.Prettify(s)
}

func (s ModifyViewLayoutRequestPanesImages) GoString() string {
	return s.String()
}

func (s *ModifyViewLayoutRequestPanesImages) GetAlpha() *float64 {
	return s.Alpha
}

func (s *ModifyViewLayoutRequestPanesImages) GetHeight() *float64 {
	return s.Height
}

func (s *ModifyViewLayoutRequestPanesImages) GetLayer() *int32 {
	return s.Layer
}

func (s *ModifyViewLayoutRequestPanesImages) GetPaneImageCropMode() *int32 {
	return s.PaneImageCropMode
}

func (s *ModifyViewLayoutRequestPanesImages) GetUrl() *string {
	return s.Url
}

func (s *ModifyViewLayoutRequestPanesImages) GetWidth() *float64 {
	return s.Width
}

func (s *ModifyViewLayoutRequestPanesImages) GetX() *float64 {
	return s.X
}

func (s *ModifyViewLayoutRequestPanesImages) GetY() *float64 {
	return s.Y
}

func (s *ModifyViewLayoutRequestPanesImages) SetAlpha(v float64) *ModifyViewLayoutRequestPanesImages {
	s.Alpha = &v
	return s
}

func (s *ModifyViewLayoutRequestPanesImages) SetHeight(v float64) *ModifyViewLayoutRequestPanesImages {
	s.Height = &v
	return s
}

func (s *ModifyViewLayoutRequestPanesImages) SetLayer(v int32) *ModifyViewLayoutRequestPanesImages {
	s.Layer = &v
	return s
}

func (s *ModifyViewLayoutRequestPanesImages) SetPaneImageCropMode(v int32) *ModifyViewLayoutRequestPanesImages {
	s.PaneImageCropMode = &v
	return s
}

func (s *ModifyViewLayoutRequestPanesImages) SetUrl(v string) *ModifyViewLayoutRequestPanesImages {
	s.Url = &v
	return s
}

func (s *ModifyViewLayoutRequestPanesImages) SetWidth(v float64) *ModifyViewLayoutRequestPanesImages {
	s.Width = &v
	return s
}

func (s *ModifyViewLayoutRequestPanesImages) SetX(v float64) *ModifyViewLayoutRequestPanesImages {
	s.X = &v
	return s
}

func (s *ModifyViewLayoutRequestPanesImages) SetY(v float64) *ModifyViewLayoutRequestPanesImages {
	s.Y = &v
	return s
}

func (s *ModifyViewLayoutRequestPanesImages) Validate() error {
	return dara.Validate(s)
}

type ModifyViewLayoutRequestPanesTexts struct {
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
	BoxBorderw *int32                                     `json:"BoxBorderw,omitempty" xml:"BoxBorderw,omitempty"`
	BoxColor   *ModifyViewLayoutRequestPanesTextsBoxColor `json:"BoxColor,omitempty" xml:"BoxColor,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Font      *int32                                      `json:"Font,omitempty" xml:"Font,omitempty"`
	FontColor *ModifyViewLayoutRequestPanesTextsFontColor `json:"FontColor,omitempty" xml:"FontColor,omitempty" type:"Struct"`
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

func (s ModifyViewLayoutRequestPanesTexts) String() string {
	return dara.Prettify(s)
}

func (s ModifyViewLayoutRequestPanesTexts) GoString() string {
	return s.String()
}

func (s *ModifyViewLayoutRequestPanesTexts) GetAlpha() *float64 {
	return s.Alpha
}

func (s *ModifyViewLayoutRequestPanesTexts) GetBoxAlpha() *float64 {
	return s.BoxAlpha
}

func (s *ModifyViewLayoutRequestPanesTexts) GetBoxBorderw() *int32 {
	return s.BoxBorderw
}

func (s *ModifyViewLayoutRequestPanesTexts) GetBoxColor() *ModifyViewLayoutRequestPanesTextsBoxColor {
	return s.BoxColor
}

func (s *ModifyViewLayoutRequestPanesTexts) GetFont() *int32 {
	return s.Font
}

func (s *ModifyViewLayoutRequestPanesTexts) GetFontColor() *ModifyViewLayoutRequestPanesTextsFontColor {
	return s.FontColor
}

func (s *ModifyViewLayoutRequestPanesTexts) GetFontSize() *int32 {
	return s.FontSize
}

func (s *ModifyViewLayoutRequestPanesTexts) GetHasBox() *bool {
	return s.HasBox
}

func (s *ModifyViewLayoutRequestPanesTexts) GetLayer() *int32 {
	return s.Layer
}

func (s *ModifyViewLayoutRequestPanesTexts) GetTexture() *string {
	return s.Texture
}

func (s *ModifyViewLayoutRequestPanesTexts) GetX() *float64 {
	return s.X
}

func (s *ModifyViewLayoutRequestPanesTexts) GetY() *float64 {
	return s.Y
}

func (s *ModifyViewLayoutRequestPanesTexts) SetAlpha(v float64) *ModifyViewLayoutRequestPanesTexts {
	s.Alpha = &v
	return s
}

func (s *ModifyViewLayoutRequestPanesTexts) SetBoxAlpha(v float64) *ModifyViewLayoutRequestPanesTexts {
	s.BoxAlpha = &v
	return s
}

func (s *ModifyViewLayoutRequestPanesTexts) SetBoxBorderw(v int32) *ModifyViewLayoutRequestPanesTexts {
	s.BoxBorderw = &v
	return s
}

func (s *ModifyViewLayoutRequestPanesTexts) SetBoxColor(v *ModifyViewLayoutRequestPanesTextsBoxColor) *ModifyViewLayoutRequestPanesTexts {
	s.BoxColor = v
	return s
}

func (s *ModifyViewLayoutRequestPanesTexts) SetFont(v int32) *ModifyViewLayoutRequestPanesTexts {
	s.Font = &v
	return s
}

func (s *ModifyViewLayoutRequestPanesTexts) SetFontColor(v *ModifyViewLayoutRequestPanesTextsFontColor) *ModifyViewLayoutRequestPanesTexts {
	s.FontColor = v
	return s
}

func (s *ModifyViewLayoutRequestPanesTexts) SetFontSize(v int32) *ModifyViewLayoutRequestPanesTexts {
	s.FontSize = &v
	return s
}

func (s *ModifyViewLayoutRequestPanesTexts) SetHasBox(v bool) *ModifyViewLayoutRequestPanesTexts {
	s.HasBox = &v
	return s
}

func (s *ModifyViewLayoutRequestPanesTexts) SetLayer(v int32) *ModifyViewLayoutRequestPanesTexts {
	s.Layer = &v
	return s
}

func (s *ModifyViewLayoutRequestPanesTexts) SetTexture(v string) *ModifyViewLayoutRequestPanesTexts {
	s.Texture = &v
	return s
}

func (s *ModifyViewLayoutRequestPanesTexts) SetX(v float64) *ModifyViewLayoutRequestPanesTexts {
	s.X = &v
	return s
}

func (s *ModifyViewLayoutRequestPanesTexts) SetY(v float64) *ModifyViewLayoutRequestPanesTexts {
	s.Y = &v
	return s
}

func (s *ModifyViewLayoutRequestPanesTexts) Validate() error {
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

type ModifyViewLayoutRequestPanesTextsBoxColor struct {
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

func (s ModifyViewLayoutRequestPanesTextsBoxColor) String() string {
	return dara.Prettify(s)
}

func (s ModifyViewLayoutRequestPanesTextsBoxColor) GoString() string {
	return s.String()
}

func (s *ModifyViewLayoutRequestPanesTextsBoxColor) GetB() *int32 {
	return s.B
}

func (s *ModifyViewLayoutRequestPanesTextsBoxColor) GetG() *int32 {
	return s.G
}

func (s *ModifyViewLayoutRequestPanesTextsBoxColor) GetR() *int32 {
	return s.R
}

func (s *ModifyViewLayoutRequestPanesTextsBoxColor) SetB(v int32) *ModifyViewLayoutRequestPanesTextsBoxColor {
	s.B = &v
	return s
}

func (s *ModifyViewLayoutRequestPanesTextsBoxColor) SetG(v int32) *ModifyViewLayoutRequestPanesTextsBoxColor {
	s.G = &v
	return s
}

func (s *ModifyViewLayoutRequestPanesTextsBoxColor) SetR(v int32) *ModifyViewLayoutRequestPanesTextsBoxColor {
	s.R = &v
	return s
}

func (s *ModifyViewLayoutRequestPanesTextsBoxColor) Validate() error {
	return dara.Validate(s)
}

type ModifyViewLayoutRequestPanesTextsFontColor struct {
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

func (s ModifyViewLayoutRequestPanesTextsFontColor) String() string {
	return dara.Prettify(s)
}

func (s ModifyViewLayoutRequestPanesTextsFontColor) GoString() string {
	return s.String()
}

func (s *ModifyViewLayoutRequestPanesTextsFontColor) GetB() *int32 {
	return s.B
}

func (s *ModifyViewLayoutRequestPanesTextsFontColor) GetG() *int32 {
	return s.G
}

func (s *ModifyViewLayoutRequestPanesTextsFontColor) GetR() *int32 {
	return s.R
}

func (s *ModifyViewLayoutRequestPanesTextsFontColor) SetB(v int32) *ModifyViewLayoutRequestPanesTextsFontColor {
	s.B = &v
	return s
}

func (s *ModifyViewLayoutRequestPanesTextsFontColor) SetG(v int32) *ModifyViewLayoutRequestPanesTextsFontColor {
	s.G = &v
	return s
}

func (s *ModifyViewLayoutRequestPanesTextsFontColor) SetR(v int32) *ModifyViewLayoutRequestPanesTextsFontColor {
	s.R = &v
	return s
}

func (s *ModifyViewLayoutRequestPanesTextsFontColor) Validate() error {
	return dara.Validate(s)
}

type ModifyViewLayoutRequestTexts struct {
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
	BoxBorderw *int32                                `json:"BoxBorderw,omitempty" xml:"BoxBorderw,omitempty"`
	BoxColor   *ModifyViewLayoutRequestTextsBoxColor `json:"BoxColor,omitempty" xml:"BoxColor,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Font      *int32                                 `json:"Font,omitempty" xml:"Font,omitempty"`
	FontColor *ModifyViewLayoutRequestTextsFontColor `json:"FontColor,omitempty" xml:"FontColor,omitempty" type:"Struct"`
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

func (s ModifyViewLayoutRequestTexts) String() string {
	return dara.Prettify(s)
}

func (s ModifyViewLayoutRequestTexts) GoString() string {
	return s.String()
}

func (s *ModifyViewLayoutRequestTexts) GetAlpha() *float64 {
	return s.Alpha
}

func (s *ModifyViewLayoutRequestTexts) GetBoxAlpha() *float64 {
	return s.BoxAlpha
}

func (s *ModifyViewLayoutRequestTexts) GetBoxBorderw() *int32 {
	return s.BoxBorderw
}

func (s *ModifyViewLayoutRequestTexts) GetBoxColor() *ModifyViewLayoutRequestTextsBoxColor {
	return s.BoxColor
}

func (s *ModifyViewLayoutRequestTexts) GetFont() *int32 {
	return s.Font
}

func (s *ModifyViewLayoutRequestTexts) GetFontColor() *ModifyViewLayoutRequestTextsFontColor {
	return s.FontColor
}

func (s *ModifyViewLayoutRequestTexts) GetFontSize() *int32 {
	return s.FontSize
}

func (s *ModifyViewLayoutRequestTexts) GetHasBox() *bool {
	return s.HasBox
}

func (s *ModifyViewLayoutRequestTexts) GetLayer() *int32 {
	return s.Layer
}

func (s *ModifyViewLayoutRequestTexts) GetTexture() *string {
	return s.Texture
}

func (s *ModifyViewLayoutRequestTexts) GetX() *float64 {
	return s.X
}

func (s *ModifyViewLayoutRequestTexts) GetY() *float64 {
	return s.Y
}

func (s *ModifyViewLayoutRequestTexts) SetAlpha(v float64) *ModifyViewLayoutRequestTexts {
	s.Alpha = &v
	return s
}

func (s *ModifyViewLayoutRequestTexts) SetBoxAlpha(v float64) *ModifyViewLayoutRequestTexts {
	s.BoxAlpha = &v
	return s
}

func (s *ModifyViewLayoutRequestTexts) SetBoxBorderw(v int32) *ModifyViewLayoutRequestTexts {
	s.BoxBorderw = &v
	return s
}

func (s *ModifyViewLayoutRequestTexts) SetBoxColor(v *ModifyViewLayoutRequestTextsBoxColor) *ModifyViewLayoutRequestTexts {
	s.BoxColor = v
	return s
}

func (s *ModifyViewLayoutRequestTexts) SetFont(v int32) *ModifyViewLayoutRequestTexts {
	s.Font = &v
	return s
}

func (s *ModifyViewLayoutRequestTexts) SetFontColor(v *ModifyViewLayoutRequestTextsFontColor) *ModifyViewLayoutRequestTexts {
	s.FontColor = v
	return s
}

func (s *ModifyViewLayoutRequestTexts) SetFontSize(v int32) *ModifyViewLayoutRequestTexts {
	s.FontSize = &v
	return s
}

func (s *ModifyViewLayoutRequestTexts) SetHasBox(v bool) *ModifyViewLayoutRequestTexts {
	s.HasBox = &v
	return s
}

func (s *ModifyViewLayoutRequestTexts) SetLayer(v int32) *ModifyViewLayoutRequestTexts {
	s.Layer = &v
	return s
}

func (s *ModifyViewLayoutRequestTexts) SetTexture(v string) *ModifyViewLayoutRequestTexts {
	s.Texture = &v
	return s
}

func (s *ModifyViewLayoutRequestTexts) SetX(v float64) *ModifyViewLayoutRequestTexts {
	s.X = &v
	return s
}

func (s *ModifyViewLayoutRequestTexts) SetY(v float64) *ModifyViewLayoutRequestTexts {
	s.Y = &v
	return s
}

func (s *ModifyViewLayoutRequestTexts) Validate() error {
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

type ModifyViewLayoutRequestTextsBoxColor struct {
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

func (s ModifyViewLayoutRequestTextsBoxColor) String() string {
	return dara.Prettify(s)
}

func (s ModifyViewLayoutRequestTextsBoxColor) GoString() string {
	return s.String()
}

func (s *ModifyViewLayoutRequestTextsBoxColor) GetB() *int32 {
	return s.B
}

func (s *ModifyViewLayoutRequestTextsBoxColor) GetG() *int32 {
	return s.G
}

func (s *ModifyViewLayoutRequestTextsBoxColor) GetR() *int32 {
	return s.R
}

func (s *ModifyViewLayoutRequestTextsBoxColor) SetB(v int32) *ModifyViewLayoutRequestTextsBoxColor {
	s.B = &v
	return s
}

func (s *ModifyViewLayoutRequestTextsBoxColor) SetG(v int32) *ModifyViewLayoutRequestTextsBoxColor {
	s.G = &v
	return s
}

func (s *ModifyViewLayoutRequestTextsBoxColor) SetR(v int32) *ModifyViewLayoutRequestTextsBoxColor {
	s.R = &v
	return s
}

func (s *ModifyViewLayoutRequestTextsBoxColor) Validate() error {
	return dara.Validate(s)
}

type ModifyViewLayoutRequestTextsFontColor struct {
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

func (s ModifyViewLayoutRequestTextsFontColor) String() string {
	return dara.Prettify(s)
}

func (s ModifyViewLayoutRequestTextsFontColor) GoString() string {
	return s.String()
}

func (s *ModifyViewLayoutRequestTextsFontColor) GetB() *int32 {
	return s.B
}

func (s *ModifyViewLayoutRequestTextsFontColor) GetG() *int32 {
	return s.G
}

func (s *ModifyViewLayoutRequestTextsFontColor) GetR() *int32 {
	return s.R
}

func (s *ModifyViewLayoutRequestTextsFontColor) SetB(v int32) *ModifyViewLayoutRequestTextsFontColor {
	s.B = &v
	return s
}

func (s *ModifyViewLayoutRequestTextsFontColor) SetG(v int32) *ModifyViewLayoutRequestTextsFontColor {
	s.G = &v
	return s
}

func (s *ModifyViewLayoutRequestTextsFontColor) SetR(v int32) *ModifyViewLayoutRequestTextsFontColor {
	s.R = &v
	return s
}

func (s *ModifyViewLayoutRequestTextsFontColor) Validate() error {
	return dara.Validate(s)
}
