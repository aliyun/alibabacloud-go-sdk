// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateCloudRecordRequest
	GetAppId() *string
	SetBackgrounds(v []*UpdateCloudRecordRequestBackgrounds) *UpdateCloudRecordRequest
	GetBackgrounds() []*UpdateCloudRecordRequestBackgrounds
	SetChannelId(v string) *UpdateCloudRecordRequest
	GetChannelId() *string
	SetClockWidgets(v []*UpdateCloudRecordRequestClockWidgets) *UpdateCloudRecordRequest
	GetClockWidgets() []*UpdateCloudRecordRequestClockWidgets
	SetImages(v []*UpdateCloudRecordRequestImages) *UpdateCloudRecordRequest
	GetImages() []*UpdateCloudRecordRequestImages
	SetLayoutSpecifiedUsers(v *UpdateCloudRecordRequestLayoutSpecifiedUsers) *UpdateCloudRecordRequest
	GetLayoutSpecifiedUsers() *UpdateCloudRecordRequestLayoutSpecifiedUsers
	SetPanes(v []*UpdateCloudRecordRequestPanes) *UpdateCloudRecordRequest
	GetPanes() []*UpdateCloudRecordRequestPanes
	SetTaskId(v string) *UpdateCloudRecordRequest
	GetTaskId() *string
	SetTemplateId(v string) *UpdateCloudRecordRequest
	GetTemplateId() *string
	SetTexts(v []*UpdateCloudRecordRequestTexts) *UpdateCloudRecordRequest
	GetTexts() []*UpdateCloudRecordRequestTexts
}

type UpdateCloudRecordRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId       *string                                `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Backgrounds []*UpdateCloudRecordRequestBackgrounds `json:"Backgrounds,omitempty" xml:"Backgrounds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId            *string                                       `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ClockWidgets         []*UpdateCloudRecordRequestClockWidgets       `json:"ClockWidgets,omitempty" xml:"ClockWidgets,omitempty" type:"Repeated"`
	Images               []*UpdateCloudRecordRequestImages             `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	LayoutSpecifiedUsers *UpdateCloudRecordRequestLayoutSpecifiedUsers `json:"LayoutSpecifiedUsers,omitempty" xml:"LayoutSpecifiedUsers,omitempty" type:"Struct"`
	Panes                []*UpdateCloudRecordRequestPanes              `json:"Panes,omitempty" xml:"Panes,omitempty" type:"Repeated"`
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
	TemplateId *string                          `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Texts      []*UpdateCloudRecordRequestTexts `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
}

func (s UpdateCloudRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudRecordRequest) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateCloudRecordRequest) GetBackgrounds() []*UpdateCloudRecordRequestBackgrounds {
	return s.Backgrounds
}

func (s *UpdateCloudRecordRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *UpdateCloudRecordRequest) GetClockWidgets() []*UpdateCloudRecordRequestClockWidgets {
	return s.ClockWidgets
}

func (s *UpdateCloudRecordRequest) GetImages() []*UpdateCloudRecordRequestImages {
	return s.Images
}

func (s *UpdateCloudRecordRequest) GetLayoutSpecifiedUsers() *UpdateCloudRecordRequestLayoutSpecifiedUsers {
	return s.LayoutSpecifiedUsers
}

func (s *UpdateCloudRecordRequest) GetPanes() []*UpdateCloudRecordRequestPanes {
	return s.Panes
}

func (s *UpdateCloudRecordRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateCloudRecordRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateCloudRecordRequest) GetTexts() []*UpdateCloudRecordRequestTexts {
	return s.Texts
}

func (s *UpdateCloudRecordRequest) SetAppId(v string) *UpdateCloudRecordRequest {
	s.AppId = &v
	return s
}

func (s *UpdateCloudRecordRequest) SetBackgrounds(v []*UpdateCloudRecordRequestBackgrounds) *UpdateCloudRecordRequest {
	s.Backgrounds = v
	return s
}

func (s *UpdateCloudRecordRequest) SetChannelId(v string) *UpdateCloudRecordRequest {
	s.ChannelId = &v
	return s
}

func (s *UpdateCloudRecordRequest) SetClockWidgets(v []*UpdateCloudRecordRequestClockWidgets) *UpdateCloudRecordRequest {
	s.ClockWidgets = v
	return s
}

func (s *UpdateCloudRecordRequest) SetImages(v []*UpdateCloudRecordRequestImages) *UpdateCloudRecordRequest {
	s.Images = v
	return s
}

func (s *UpdateCloudRecordRequest) SetLayoutSpecifiedUsers(v *UpdateCloudRecordRequestLayoutSpecifiedUsers) *UpdateCloudRecordRequest {
	s.LayoutSpecifiedUsers = v
	return s
}

func (s *UpdateCloudRecordRequest) SetPanes(v []*UpdateCloudRecordRequestPanes) *UpdateCloudRecordRequest {
	s.Panes = v
	return s
}

func (s *UpdateCloudRecordRequest) SetTaskId(v string) *UpdateCloudRecordRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateCloudRecordRequest) SetTemplateId(v string) *UpdateCloudRecordRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateCloudRecordRequest) SetTexts(v []*UpdateCloudRecordRequestTexts) *UpdateCloudRecordRequest {
	s.Texts = v
	return s
}

func (s *UpdateCloudRecordRequest) Validate() error {
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

type UpdateCloudRecordRequestBackgrounds struct {
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

func (s UpdateCloudRecordRequestBackgrounds) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudRecordRequestBackgrounds) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordRequestBackgrounds) GetAlpha() *float64 {
	return s.Alpha
}

func (s *UpdateCloudRecordRequestBackgrounds) GetBackgroundCropMode() *int32 {
	return s.BackgroundCropMode
}

func (s *UpdateCloudRecordRequestBackgrounds) GetHeight() *float64 {
	return s.Height
}

func (s *UpdateCloudRecordRequestBackgrounds) GetLayer() *int32 {
	return s.Layer
}

func (s *UpdateCloudRecordRequestBackgrounds) GetUrl() *string {
	return s.Url
}

func (s *UpdateCloudRecordRequestBackgrounds) GetWidth() *float64 {
	return s.Width
}

func (s *UpdateCloudRecordRequestBackgrounds) GetX() *float64 {
	return s.X
}

func (s *UpdateCloudRecordRequestBackgrounds) GetY() *float64 {
	return s.Y
}

func (s *UpdateCloudRecordRequestBackgrounds) SetAlpha(v float64) *UpdateCloudRecordRequestBackgrounds {
	s.Alpha = &v
	return s
}

func (s *UpdateCloudRecordRequestBackgrounds) SetBackgroundCropMode(v int32) *UpdateCloudRecordRequestBackgrounds {
	s.BackgroundCropMode = &v
	return s
}

func (s *UpdateCloudRecordRequestBackgrounds) SetHeight(v float64) *UpdateCloudRecordRequestBackgrounds {
	s.Height = &v
	return s
}

func (s *UpdateCloudRecordRequestBackgrounds) SetLayer(v int32) *UpdateCloudRecordRequestBackgrounds {
	s.Layer = &v
	return s
}

func (s *UpdateCloudRecordRequestBackgrounds) SetUrl(v string) *UpdateCloudRecordRequestBackgrounds {
	s.Url = &v
	return s
}

func (s *UpdateCloudRecordRequestBackgrounds) SetWidth(v float64) *UpdateCloudRecordRequestBackgrounds {
	s.Width = &v
	return s
}

func (s *UpdateCloudRecordRequestBackgrounds) SetX(v float64) *UpdateCloudRecordRequestBackgrounds {
	s.X = &v
	return s
}

func (s *UpdateCloudRecordRequestBackgrounds) SetY(v float64) *UpdateCloudRecordRequestBackgrounds {
	s.Y = &v
	return s
}

func (s *UpdateCloudRecordRequestBackgrounds) Validate() error {
	return dara.Validate(s)
}

type UpdateCloudRecordRequestClockWidgets struct {
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
	BoxBorderw *int32                                        `json:"BoxBorderw,omitempty" xml:"BoxBorderw,omitempty"`
	BoxColor   *UpdateCloudRecordRequestClockWidgetsBoxColor `json:"BoxColor,omitempty" xml:"BoxColor,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Font      *int32                                         `json:"Font,omitempty" xml:"Font,omitempty"`
	FontColor *UpdateCloudRecordRequestClockWidgetsFontColor `json:"FontColor,omitempty" xml:"FontColor,omitempty" type:"Struct"`
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

func (s UpdateCloudRecordRequestClockWidgets) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudRecordRequestClockWidgets) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordRequestClockWidgets) GetAlpha() *float64 {
	return s.Alpha
}

func (s *UpdateCloudRecordRequestClockWidgets) GetBoxAlpha() *float64 {
	return s.BoxAlpha
}

func (s *UpdateCloudRecordRequestClockWidgets) GetBoxBorderw() *int32 {
	return s.BoxBorderw
}

func (s *UpdateCloudRecordRequestClockWidgets) GetBoxColor() *UpdateCloudRecordRequestClockWidgetsBoxColor {
	return s.BoxColor
}

func (s *UpdateCloudRecordRequestClockWidgets) GetFont() *int32 {
	return s.Font
}

func (s *UpdateCloudRecordRequestClockWidgets) GetFontColor() *UpdateCloudRecordRequestClockWidgetsFontColor {
	return s.FontColor
}

func (s *UpdateCloudRecordRequestClockWidgets) GetFontSize() *int32 {
	return s.FontSize
}

func (s *UpdateCloudRecordRequestClockWidgets) GetHasBox() *bool {
	return s.HasBox
}

func (s *UpdateCloudRecordRequestClockWidgets) GetLayer() *int32 {
	return s.Layer
}

func (s *UpdateCloudRecordRequestClockWidgets) GetX() *float64 {
	return s.X
}

func (s *UpdateCloudRecordRequestClockWidgets) GetY() *float64 {
	return s.Y
}

func (s *UpdateCloudRecordRequestClockWidgets) GetZone() *int32 {
	return s.Zone
}

func (s *UpdateCloudRecordRequestClockWidgets) SetAlpha(v float64) *UpdateCloudRecordRequestClockWidgets {
	s.Alpha = &v
	return s
}

func (s *UpdateCloudRecordRequestClockWidgets) SetBoxAlpha(v float64) *UpdateCloudRecordRequestClockWidgets {
	s.BoxAlpha = &v
	return s
}

func (s *UpdateCloudRecordRequestClockWidgets) SetBoxBorderw(v int32) *UpdateCloudRecordRequestClockWidgets {
	s.BoxBorderw = &v
	return s
}

func (s *UpdateCloudRecordRequestClockWidgets) SetBoxColor(v *UpdateCloudRecordRequestClockWidgetsBoxColor) *UpdateCloudRecordRequestClockWidgets {
	s.BoxColor = v
	return s
}

func (s *UpdateCloudRecordRequestClockWidgets) SetFont(v int32) *UpdateCloudRecordRequestClockWidgets {
	s.Font = &v
	return s
}

func (s *UpdateCloudRecordRequestClockWidgets) SetFontColor(v *UpdateCloudRecordRequestClockWidgetsFontColor) *UpdateCloudRecordRequestClockWidgets {
	s.FontColor = v
	return s
}

func (s *UpdateCloudRecordRequestClockWidgets) SetFontSize(v int32) *UpdateCloudRecordRequestClockWidgets {
	s.FontSize = &v
	return s
}

func (s *UpdateCloudRecordRequestClockWidgets) SetHasBox(v bool) *UpdateCloudRecordRequestClockWidgets {
	s.HasBox = &v
	return s
}

func (s *UpdateCloudRecordRequestClockWidgets) SetLayer(v int32) *UpdateCloudRecordRequestClockWidgets {
	s.Layer = &v
	return s
}

func (s *UpdateCloudRecordRequestClockWidgets) SetX(v float64) *UpdateCloudRecordRequestClockWidgets {
	s.X = &v
	return s
}

func (s *UpdateCloudRecordRequestClockWidgets) SetY(v float64) *UpdateCloudRecordRequestClockWidgets {
	s.Y = &v
	return s
}

func (s *UpdateCloudRecordRequestClockWidgets) SetZone(v int32) *UpdateCloudRecordRequestClockWidgets {
	s.Zone = &v
	return s
}

func (s *UpdateCloudRecordRequestClockWidgets) Validate() error {
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

type UpdateCloudRecordRequestClockWidgetsBoxColor struct {
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

func (s UpdateCloudRecordRequestClockWidgetsBoxColor) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudRecordRequestClockWidgetsBoxColor) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordRequestClockWidgetsBoxColor) GetB() *int32 {
	return s.B
}

func (s *UpdateCloudRecordRequestClockWidgetsBoxColor) GetG() *int32 {
	return s.G
}

func (s *UpdateCloudRecordRequestClockWidgetsBoxColor) GetR() *int32 {
	return s.R
}

func (s *UpdateCloudRecordRequestClockWidgetsBoxColor) SetB(v int32) *UpdateCloudRecordRequestClockWidgetsBoxColor {
	s.B = &v
	return s
}

func (s *UpdateCloudRecordRequestClockWidgetsBoxColor) SetG(v int32) *UpdateCloudRecordRequestClockWidgetsBoxColor {
	s.G = &v
	return s
}

func (s *UpdateCloudRecordRequestClockWidgetsBoxColor) SetR(v int32) *UpdateCloudRecordRequestClockWidgetsBoxColor {
	s.R = &v
	return s
}

func (s *UpdateCloudRecordRequestClockWidgetsBoxColor) Validate() error {
	return dara.Validate(s)
}

type UpdateCloudRecordRequestClockWidgetsFontColor struct {
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

func (s UpdateCloudRecordRequestClockWidgetsFontColor) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudRecordRequestClockWidgetsFontColor) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordRequestClockWidgetsFontColor) GetB() *int32 {
	return s.B
}

func (s *UpdateCloudRecordRequestClockWidgetsFontColor) GetG() *int32 {
	return s.G
}

func (s *UpdateCloudRecordRequestClockWidgetsFontColor) GetR() *int32 {
	return s.R
}

func (s *UpdateCloudRecordRequestClockWidgetsFontColor) SetB(v int32) *UpdateCloudRecordRequestClockWidgetsFontColor {
	s.B = &v
	return s
}

func (s *UpdateCloudRecordRequestClockWidgetsFontColor) SetG(v int32) *UpdateCloudRecordRequestClockWidgetsFontColor {
	s.G = &v
	return s
}

func (s *UpdateCloudRecordRequestClockWidgetsFontColor) SetR(v int32) *UpdateCloudRecordRequestClockWidgetsFontColor {
	s.R = &v
	return s
}

func (s *UpdateCloudRecordRequestClockWidgetsFontColor) Validate() error {
	return dara.Validate(s)
}

type UpdateCloudRecordRequestImages struct {
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

func (s UpdateCloudRecordRequestImages) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudRecordRequestImages) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordRequestImages) GetAlpha() *float64 {
	return s.Alpha
}

func (s *UpdateCloudRecordRequestImages) GetHeight() *float64 {
	return s.Height
}

func (s *UpdateCloudRecordRequestImages) GetImageCropMode() *int32 {
	return s.ImageCropMode
}

func (s *UpdateCloudRecordRequestImages) GetLayer() *int32 {
	return s.Layer
}

func (s *UpdateCloudRecordRequestImages) GetUrl() *string {
	return s.Url
}

func (s *UpdateCloudRecordRequestImages) GetWidth() *float64 {
	return s.Width
}

func (s *UpdateCloudRecordRequestImages) GetX() *float64 {
	return s.X
}

func (s *UpdateCloudRecordRequestImages) GetY() *float64 {
	return s.Y
}

func (s *UpdateCloudRecordRequestImages) SetAlpha(v float64) *UpdateCloudRecordRequestImages {
	s.Alpha = &v
	return s
}

func (s *UpdateCloudRecordRequestImages) SetHeight(v float64) *UpdateCloudRecordRequestImages {
	s.Height = &v
	return s
}

func (s *UpdateCloudRecordRequestImages) SetImageCropMode(v int32) *UpdateCloudRecordRequestImages {
	s.ImageCropMode = &v
	return s
}

func (s *UpdateCloudRecordRequestImages) SetLayer(v int32) *UpdateCloudRecordRequestImages {
	s.Layer = &v
	return s
}

func (s *UpdateCloudRecordRequestImages) SetUrl(v string) *UpdateCloudRecordRequestImages {
	s.Url = &v
	return s
}

func (s *UpdateCloudRecordRequestImages) SetWidth(v float64) *UpdateCloudRecordRequestImages {
	s.Width = &v
	return s
}

func (s *UpdateCloudRecordRequestImages) SetX(v float64) *UpdateCloudRecordRequestImages {
	s.X = &v
	return s
}

func (s *UpdateCloudRecordRequestImages) SetY(v float64) *UpdateCloudRecordRequestImages {
	s.Y = &v
	return s
}

func (s *UpdateCloudRecordRequestImages) Validate() error {
	return dara.Validate(s)
}

type UpdateCloudRecordRequestLayoutSpecifiedUsers struct {
	// This parameter is required.
	Ids []*string `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateCloudRecordRequestLayoutSpecifiedUsers) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudRecordRequestLayoutSpecifiedUsers) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordRequestLayoutSpecifiedUsers) GetIds() []*string {
	return s.Ids
}

func (s *UpdateCloudRecordRequestLayoutSpecifiedUsers) GetType() *string {
	return s.Type
}

func (s *UpdateCloudRecordRequestLayoutSpecifiedUsers) SetIds(v []*string) *UpdateCloudRecordRequestLayoutSpecifiedUsers {
	s.Ids = v
	return s
}

func (s *UpdateCloudRecordRequestLayoutSpecifiedUsers) SetType(v string) *UpdateCloudRecordRequestLayoutSpecifiedUsers {
	s.Type = &v
	return s
}

func (s *UpdateCloudRecordRequestLayoutSpecifiedUsers) Validate() error {
	return dara.Validate(s)
}

type UpdateCloudRecordRequestPanes struct {
	Backgrounds []*UpdateCloudRecordRequestPanesBackgrounds `json:"Backgrounds,omitempty" xml:"Backgrounds,omitempty" type:"Repeated"`
	Images      []*UpdateCloudRecordRequestPanesImages      `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
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
	SourceType *string                               `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Texts      []*UpdateCloudRecordRequestPanesTexts `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
	// example:
	//
	// cameraFirst
	VideoOrder *string                                  `json:"VideoOrder,omitempty" xml:"VideoOrder,omitempty"`
	Whiteboard *UpdateCloudRecordRequestPanesWhiteboard `json:"Whiteboard,omitempty" xml:"Whiteboard,omitempty" type:"Struct"`
}

func (s UpdateCloudRecordRequestPanes) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudRecordRequestPanes) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordRequestPanes) GetBackgrounds() []*UpdateCloudRecordRequestPanesBackgrounds {
	return s.Backgrounds
}

func (s *UpdateCloudRecordRequestPanes) GetImages() []*UpdateCloudRecordRequestPanesImages {
	return s.Images
}

func (s *UpdateCloudRecordRequestPanes) GetPaneCropMode() *int32 {
	return s.PaneCropMode
}

func (s *UpdateCloudRecordRequestPanes) GetPaneId() *int32 {
	return s.PaneId
}

func (s *UpdateCloudRecordRequestPanes) GetReservePaneForOfflineUser() *bool {
	return s.ReservePaneForOfflineUser
}

func (s *UpdateCloudRecordRequestPanes) GetSource() *string {
	return s.Source
}

func (s *UpdateCloudRecordRequestPanes) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateCloudRecordRequestPanes) GetTexts() []*UpdateCloudRecordRequestPanesTexts {
	return s.Texts
}

func (s *UpdateCloudRecordRequestPanes) GetVideoOrder() *string {
	return s.VideoOrder
}

func (s *UpdateCloudRecordRequestPanes) GetWhiteboard() *UpdateCloudRecordRequestPanesWhiteboard {
	return s.Whiteboard
}

func (s *UpdateCloudRecordRequestPanes) SetBackgrounds(v []*UpdateCloudRecordRequestPanesBackgrounds) *UpdateCloudRecordRequestPanes {
	s.Backgrounds = v
	return s
}

func (s *UpdateCloudRecordRequestPanes) SetImages(v []*UpdateCloudRecordRequestPanesImages) *UpdateCloudRecordRequestPanes {
	s.Images = v
	return s
}

func (s *UpdateCloudRecordRequestPanes) SetPaneCropMode(v int32) *UpdateCloudRecordRequestPanes {
	s.PaneCropMode = &v
	return s
}

func (s *UpdateCloudRecordRequestPanes) SetPaneId(v int32) *UpdateCloudRecordRequestPanes {
	s.PaneId = &v
	return s
}

func (s *UpdateCloudRecordRequestPanes) SetReservePaneForOfflineUser(v bool) *UpdateCloudRecordRequestPanes {
	s.ReservePaneForOfflineUser = &v
	return s
}

func (s *UpdateCloudRecordRequestPanes) SetSource(v string) *UpdateCloudRecordRequestPanes {
	s.Source = &v
	return s
}

func (s *UpdateCloudRecordRequestPanes) SetSourceType(v string) *UpdateCloudRecordRequestPanes {
	s.SourceType = &v
	return s
}

func (s *UpdateCloudRecordRequestPanes) SetTexts(v []*UpdateCloudRecordRequestPanesTexts) *UpdateCloudRecordRequestPanes {
	s.Texts = v
	return s
}

func (s *UpdateCloudRecordRequestPanes) SetVideoOrder(v string) *UpdateCloudRecordRequestPanes {
	s.VideoOrder = &v
	return s
}

func (s *UpdateCloudRecordRequestPanes) SetWhiteboard(v *UpdateCloudRecordRequestPanesWhiteboard) *UpdateCloudRecordRequestPanes {
	s.Whiteboard = v
	return s
}

func (s *UpdateCloudRecordRequestPanes) Validate() error {
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

type UpdateCloudRecordRequestPanesBackgrounds struct {
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

func (s UpdateCloudRecordRequestPanesBackgrounds) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudRecordRequestPanesBackgrounds) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordRequestPanesBackgrounds) GetAlpha() *float64 {
	return s.Alpha
}

func (s *UpdateCloudRecordRequestPanesBackgrounds) GetDisplay() *string {
	return s.Display
}

func (s *UpdateCloudRecordRequestPanesBackgrounds) GetHeight() *float64 {
	return s.Height
}

func (s *UpdateCloudRecordRequestPanesBackgrounds) GetLayer() *int32 {
	return s.Layer
}

func (s *UpdateCloudRecordRequestPanesBackgrounds) GetPaneBackgroundCropMode() *int32 {
	return s.PaneBackgroundCropMode
}

func (s *UpdateCloudRecordRequestPanesBackgrounds) GetUrl() *string {
	return s.Url
}

func (s *UpdateCloudRecordRequestPanesBackgrounds) GetWidth() *float64 {
	return s.Width
}

func (s *UpdateCloudRecordRequestPanesBackgrounds) GetX() *float64 {
	return s.X
}

func (s *UpdateCloudRecordRequestPanesBackgrounds) GetY() *float64 {
	return s.Y
}

func (s *UpdateCloudRecordRequestPanesBackgrounds) SetAlpha(v float64) *UpdateCloudRecordRequestPanesBackgrounds {
	s.Alpha = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesBackgrounds) SetDisplay(v string) *UpdateCloudRecordRequestPanesBackgrounds {
	s.Display = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesBackgrounds) SetHeight(v float64) *UpdateCloudRecordRequestPanesBackgrounds {
	s.Height = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesBackgrounds) SetLayer(v int32) *UpdateCloudRecordRequestPanesBackgrounds {
	s.Layer = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesBackgrounds) SetPaneBackgroundCropMode(v int32) *UpdateCloudRecordRequestPanesBackgrounds {
	s.PaneBackgroundCropMode = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesBackgrounds) SetUrl(v string) *UpdateCloudRecordRequestPanesBackgrounds {
	s.Url = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesBackgrounds) SetWidth(v float64) *UpdateCloudRecordRequestPanesBackgrounds {
	s.Width = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesBackgrounds) SetX(v float64) *UpdateCloudRecordRequestPanesBackgrounds {
	s.X = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesBackgrounds) SetY(v float64) *UpdateCloudRecordRequestPanesBackgrounds {
	s.Y = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesBackgrounds) Validate() error {
	return dara.Validate(s)
}

type UpdateCloudRecordRequestPanesImages struct {
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

func (s UpdateCloudRecordRequestPanesImages) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudRecordRequestPanesImages) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordRequestPanesImages) GetAlpha() *float64 {
	return s.Alpha
}

func (s *UpdateCloudRecordRequestPanesImages) GetDisplay() *string {
	return s.Display
}

func (s *UpdateCloudRecordRequestPanesImages) GetHeight() *float64 {
	return s.Height
}

func (s *UpdateCloudRecordRequestPanesImages) GetLayer() *int32 {
	return s.Layer
}

func (s *UpdateCloudRecordRequestPanesImages) GetPaneImageCropMode() *int32 {
	return s.PaneImageCropMode
}

func (s *UpdateCloudRecordRequestPanesImages) GetUrl() *string {
	return s.Url
}

func (s *UpdateCloudRecordRequestPanesImages) GetWidth() *float64 {
	return s.Width
}

func (s *UpdateCloudRecordRequestPanesImages) GetX() *float64 {
	return s.X
}

func (s *UpdateCloudRecordRequestPanesImages) GetY() *float64 {
	return s.Y
}

func (s *UpdateCloudRecordRequestPanesImages) SetAlpha(v float64) *UpdateCloudRecordRequestPanesImages {
	s.Alpha = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesImages) SetDisplay(v string) *UpdateCloudRecordRequestPanesImages {
	s.Display = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesImages) SetHeight(v float64) *UpdateCloudRecordRequestPanesImages {
	s.Height = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesImages) SetLayer(v int32) *UpdateCloudRecordRequestPanesImages {
	s.Layer = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesImages) SetPaneImageCropMode(v int32) *UpdateCloudRecordRequestPanesImages {
	s.PaneImageCropMode = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesImages) SetUrl(v string) *UpdateCloudRecordRequestPanesImages {
	s.Url = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesImages) SetWidth(v float64) *UpdateCloudRecordRequestPanesImages {
	s.Width = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesImages) SetX(v float64) *UpdateCloudRecordRequestPanesImages {
	s.X = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesImages) SetY(v float64) *UpdateCloudRecordRequestPanesImages {
	s.Y = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesImages) Validate() error {
	return dara.Validate(s)
}

type UpdateCloudRecordRequestPanesTexts struct {
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
	BoxBorderw *int32                                      `json:"BoxBorderw,omitempty" xml:"BoxBorderw,omitempty"`
	BoxColor   *UpdateCloudRecordRequestPanesTextsBoxColor `json:"BoxColor,omitempty" xml:"BoxColor,omitempty" type:"Struct"`
	// example:
	//
	// backup
	Display *string `json:"Display,omitempty" xml:"Display,omitempty"`
	// example:
	//
	// 0
	Font      *int32                                       `json:"Font,omitempty" xml:"Font,omitempty"`
	FontColor *UpdateCloudRecordRequestPanesTextsFontColor `json:"FontColor,omitempty" xml:"FontColor,omitempty" type:"Struct"`
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

func (s UpdateCloudRecordRequestPanesTexts) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudRecordRequestPanesTexts) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordRequestPanesTexts) GetAlpha() *float64 {
	return s.Alpha
}

func (s *UpdateCloudRecordRequestPanesTexts) GetBoxAlpha() *float64 {
	return s.BoxAlpha
}

func (s *UpdateCloudRecordRequestPanesTexts) GetBoxBorderw() *int32 {
	return s.BoxBorderw
}

func (s *UpdateCloudRecordRequestPanesTexts) GetBoxColor() *UpdateCloudRecordRequestPanesTextsBoxColor {
	return s.BoxColor
}

func (s *UpdateCloudRecordRequestPanesTexts) GetDisplay() *string {
	return s.Display
}

func (s *UpdateCloudRecordRequestPanesTexts) GetFont() *int32 {
	return s.Font
}

func (s *UpdateCloudRecordRequestPanesTexts) GetFontColor() *UpdateCloudRecordRequestPanesTextsFontColor {
	return s.FontColor
}

func (s *UpdateCloudRecordRequestPanesTexts) GetFontSize() *int32 {
	return s.FontSize
}

func (s *UpdateCloudRecordRequestPanesTexts) GetHasBox() *bool {
	return s.HasBox
}

func (s *UpdateCloudRecordRequestPanesTexts) GetLayer() *int32 {
	return s.Layer
}

func (s *UpdateCloudRecordRequestPanesTexts) GetTexture() *string {
	return s.Texture
}

func (s *UpdateCloudRecordRequestPanesTexts) GetX() *float64 {
	return s.X
}

func (s *UpdateCloudRecordRequestPanesTexts) GetY() *float64 {
	return s.Y
}

func (s *UpdateCloudRecordRequestPanesTexts) SetAlpha(v float64) *UpdateCloudRecordRequestPanesTexts {
	s.Alpha = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesTexts) SetBoxAlpha(v float64) *UpdateCloudRecordRequestPanesTexts {
	s.BoxAlpha = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesTexts) SetBoxBorderw(v int32) *UpdateCloudRecordRequestPanesTexts {
	s.BoxBorderw = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesTexts) SetBoxColor(v *UpdateCloudRecordRequestPanesTextsBoxColor) *UpdateCloudRecordRequestPanesTexts {
	s.BoxColor = v
	return s
}

func (s *UpdateCloudRecordRequestPanesTexts) SetDisplay(v string) *UpdateCloudRecordRequestPanesTexts {
	s.Display = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesTexts) SetFont(v int32) *UpdateCloudRecordRequestPanesTexts {
	s.Font = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesTexts) SetFontColor(v *UpdateCloudRecordRequestPanesTextsFontColor) *UpdateCloudRecordRequestPanesTexts {
	s.FontColor = v
	return s
}

func (s *UpdateCloudRecordRequestPanesTexts) SetFontSize(v int32) *UpdateCloudRecordRequestPanesTexts {
	s.FontSize = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesTexts) SetHasBox(v bool) *UpdateCloudRecordRequestPanesTexts {
	s.HasBox = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesTexts) SetLayer(v int32) *UpdateCloudRecordRequestPanesTexts {
	s.Layer = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesTexts) SetTexture(v string) *UpdateCloudRecordRequestPanesTexts {
	s.Texture = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesTexts) SetX(v float64) *UpdateCloudRecordRequestPanesTexts {
	s.X = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesTexts) SetY(v float64) *UpdateCloudRecordRequestPanesTexts {
	s.Y = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesTexts) Validate() error {
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

type UpdateCloudRecordRequestPanesTextsBoxColor struct {
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

func (s UpdateCloudRecordRequestPanesTextsBoxColor) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudRecordRequestPanesTextsBoxColor) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordRequestPanesTextsBoxColor) GetB() *int32 {
	return s.B
}

func (s *UpdateCloudRecordRequestPanesTextsBoxColor) GetG() *int32 {
	return s.G
}

func (s *UpdateCloudRecordRequestPanesTextsBoxColor) GetR() *int32 {
	return s.R
}

func (s *UpdateCloudRecordRequestPanesTextsBoxColor) SetB(v int32) *UpdateCloudRecordRequestPanesTextsBoxColor {
	s.B = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesTextsBoxColor) SetG(v int32) *UpdateCloudRecordRequestPanesTextsBoxColor {
	s.G = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesTextsBoxColor) SetR(v int32) *UpdateCloudRecordRequestPanesTextsBoxColor {
	s.R = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesTextsBoxColor) Validate() error {
	return dara.Validate(s)
}

type UpdateCloudRecordRequestPanesTextsFontColor struct {
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

func (s UpdateCloudRecordRequestPanesTextsFontColor) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudRecordRequestPanesTextsFontColor) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordRequestPanesTextsFontColor) GetB() *int32 {
	return s.B
}

func (s *UpdateCloudRecordRequestPanesTextsFontColor) GetG() *int32 {
	return s.G
}

func (s *UpdateCloudRecordRequestPanesTextsFontColor) GetR() *int32 {
	return s.R
}

func (s *UpdateCloudRecordRequestPanesTextsFontColor) SetB(v int32) *UpdateCloudRecordRequestPanesTextsFontColor {
	s.B = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesTextsFontColor) SetG(v int32) *UpdateCloudRecordRequestPanesTextsFontColor {
	s.G = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesTextsFontColor) SetR(v int32) *UpdateCloudRecordRequestPanesTextsFontColor {
	s.R = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesTextsFontColor) Validate() error {
	return dara.Validate(s)
}

type UpdateCloudRecordRequestPanesWhiteboard struct {
	// example:
	//
	// default
	WhiteboardId *string `json:"WhiteboardId,omitempty" xml:"WhiteboardId,omitempty"`
}

func (s UpdateCloudRecordRequestPanesWhiteboard) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudRecordRequestPanesWhiteboard) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordRequestPanesWhiteboard) GetWhiteboardId() *string {
	return s.WhiteboardId
}

func (s *UpdateCloudRecordRequestPanesWhiteboard) SetWhiteboardId(v string) *UpdateCloudRecordRequestPanesWhiteboard {
	s.WhiteboardId = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesWhiteboard) Validate() error {
	return dara.Validate(s)
}

type UpdateCloudRecordRequestTexts struct {
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
	BoxBorderw *int32                                 `json:"BoxBorderw,omitempty" xml:"BoxBorderw,omitempty"`
	BoxColor   *UpdateCloudRecordRequestTextsBoxColor `json:"BoxColor,omitempty" xml:"BoxColor,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Font      *int32                                  `json:"Font,omitempty" xml:"Font,omitempty"`
	FontColor *UpdateCloudRecordRequestTextsFontColor `json:"FontColor,omitempty" xml:"FontColor,omitempty" type:"Struct"`
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

func (s UpdateCloudRecordRequestTexts) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudRecordRequestTexts) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordRequestTexts) GetAlpha() *float64 {
	return s.Alpha
}

func (s *UpdateCloudRecordRequestTexts) GetBoxAlpha() *float64 {
	return s.BoxAlpha
}

func (s *UpdateCloudRecordRequestTexts) GetBoxBorderw() *int32 {
	return s.BoxBorderw
}

func (s *UpdateCloudRecordRequestTexts) GetBoxColor() *UpdateCloudRecordRequestTextsBoxColor {
	return s.BoxColor
}

func (s *UpdateCloudRecordRequestTexts) GetFont() *int32 {
	return s.Font
}

func (s *UpdateCloudRecordRequestTexts) GetFontColor() *UpdateCloudRecordRequestTextsFontColor {
	return s.FontColor
}

func (s *UpdateCloudRecordRequestTexts) GetFontSize() *int32 {
	return s.FontSize
}

func (s *UpdateCloudRecordRequestTexts) GetHasBox() *bool {
	return s.HasBox
}

func (s *UpdateCloudRecordRequestTexts) GetLayer() *int32 {
	return s.Layer
}

func (s *UpdateCloudRecordRequestTexts) GetTexture() *string {
	return s.Texture
}

func (s *UpdateCloudRecordRequestTexts) GetX() *float64 {
	return s.X
}

func (s *UpdateCloudRecordRequestTexts) GetY() *float64 {
	return s.Y
}

func (s *UpdateCloudRecordRequestTexts) SetAlpha(v float64) *UpdateCloudRecordRequestTexts {
	s.Alpha = &v
	return s
}

func (s *UpdateCloudRecordRequestTexts) SetBoxAlpha(v float64) *UpdateCloudRecordRequestTexts {
	s.BoxAlpha = &v
	return s
}

func (s *UpdateCloudRecordRequestTexts) SetBoxBorderw(v int32) *UpdateCloudRecordRequestTexts {
	s.BoxBorderw = &v
	return s
}

func (s *UpdateCloudRecordRequestTexts) SetBoxColor(v *UpdateCloudRecordRequestTextsBoxColor) *UpdateCloudRecordRequestTexts {
	s.BoxColor = v
	return s
}

func (s *UpdateCloudRecordRequestTexts) SetFont(v int32) *UpdateCloudRecordRequestTexts {
	s.Font = &v
	return s
}

func (s *UpdateCloudRecordRequestTexts) SetFontColor(v *UpdateCloudRecordRequestTextsFontColor) *UpdateCloudRecordRequestTexts {
	s.FontColor = v
	return s
}

func (s *UpdateCloudRecordRequestTexts) SetFontSize(v int32) *UpdateCloudRecordRequestTexts {
	s.FontSize = &v
	return s
}

func (s *UpdateCloudRecordRequestTexts) SetHasBox(v bool) *UpdateCloudRecordRequestTexts {
	s.HasBox = &v
	return s
}

func (s *UpdateCloudRecordRequestTexts) SetLayer(v int32) *UpdateCloudRecordRequestTexts {
	s.Layer = &v
	return s
}

func (s *UpdateCloudRecordRequestTexts) SetTexture(v string) *UpdateCloudRecordRequestTexts {
	s.Texture = &v
	return s
}

func (s *UpdateCloudRecordRequestTexts) SetX(v float64) *UpdateCloudRecordRequestTexts {
	s.X = &v
	return s
}

func (s *UpdateCloudRecordRequestTexts) SetY(v float64) *UpdateCloudRecordRequestTexts {
	s.Y = &v
	return s
}

func (s *UpdateCloudRecordRequestTexts) Validate() error {
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

type UpdateCloudRecordRequestTextsBoxColor struct {
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

func (s UpdateCloudRecordRequestTextsBoxColor) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudRecordRequestTextsBoxColor) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordRequestTextsBoxColor) GetB() *int32 {
	return s.B
}

func (s *UpdateCloudRecordRequestTextsBoxColor) GetG() *int32 {
	return s.G
}

func (s *UpdateCloudRecordRequestTextsBoxColor) GetR() *int32 {
	return s.R
}

func (s *UpdateCloudRecordRequestTextsBoxColor) SetB(v int32) *UpdateCloudRecordRequestTextsBoxColor {
	s.B = &v
	return s
}

func (s *UpdateCloudRecordRequestTextsBoxColor) SetG(v int32) *UpdateCloudRecordRequestTextsBoxColor {
	s.G = &v
	return s
}

func (s *UpdateCloudRecordRequestTextsBoxColor) SetR(v int32) *UpdateCloudRecordRequestTextsBoxColor {
	s.R = &v
	return s
}

func (s *UpdateCloudRecordRequestTextsBoxColor) Validate() error {
	return dara.Validate(s)
}

type UpdateCloudRecordRequestTextsFontColor struct {
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

func (s UpdateCloudRecordRequestTextsFontColor) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudRecordRequestTextsFontColor) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordRequestTextsFontColor) GetB() *int32 {
	return s.B
}

func (s *UpdateCloudRecordRequestTextsFontColor) GetG() *int32 {
	return s.G
}

func (s *UpdateCloudRecordRequestTextsFontColor) GetR() *int32 {
	return s.R
}

func (s *UpdateCloudRecordRequestTextsFontColor) SetB(v int32) *UpdateCloudRecordRequestTextsFontColor {
	s.B = &v
	return s
}

func (s *UpdateCloudRecordRequestTextsFontColor) SetG(v int32) *UpdateCloudRecordRequestTextsFontColor {
	s.G = &v
	return s
}

func (s *UpdateCloudRecordRequestTextsFontColor) SetR(v int32) *UpdateCloudRecordRequestTextsFontColor {
	s.R = &v
	return s
}

func (s *UpdateCloudRecordRequestTextsFontColor) Validate() error {
	return dara.Validate(s)
}
