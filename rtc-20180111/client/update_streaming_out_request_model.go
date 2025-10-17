// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStreamingOutRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateStreamingOutRequest
	GetAppId() *string
	SetBackgrounds(v []*UpdateStreamingOutRequestBackgrounds) *UpdateStreamingOutRequest
	GetBackgrounds() []*UpdateStreamingOutRequestBackgrounds
	SetBgColor(v *UpdateStreamingOutRequestBgColor) *UpdateStreamingOutRequest
	GetBgColor() *UpdateStreamingOutRequestBgColor
	SetChannelId(v string) *UpdateStreamingOutRequest
	GetChannelId() *string
	SetClockWidgets(v []*UpdateStreamingOutRequestClockWidgets) *UpdateStreamingOutRequest
	GetClockWidgets() []*UpdateStreamingOutRequestClockWidgets
	SetCropMode(v int32) *UpdateStreamingOutRequest
	GetCropMode() *int32
	SetImages(v []*UpdateStreamingOutRequestImages) *UpdateStreamingOutRequest
	GetImages() []*UpdateStreamingOutRequestImages
	SetLayoutSpecifiedUsers(v *UpdateStreamingOutRequestLayoutSpecifiedUsers) *UpdateStreamingOutRequest
	GetLayoutSpecifiedUsers() *UpdateStreamingOutRequestLayoutSpecifiedUsers
	SetPanes(v []*UpdateStreamingOutRequestPanes) *UpdateStreamingOutRequest
	GetPanes() []*UpdateStreamingOutRequestPanes
	SetRegionColor(v *UpdateStreamingOutRequestRegionColor) *UpdateStreamingOutRequest
	GetRegionColor() *UpdateStreamingOutRequestRegionColor
	SetSpecMixedUserList(v []*string) *UpdateStreamingOutRequest
	GetSpecMixedUserList() []*string
	SetTaskId(v string) *UpdateStreamingOutRequest
	GetTaskId() *string
	SetTemplateId(v string) *UpdateStreamingOutRequest
	GetTemplateId() *string
	SetTexts(v []*UpdateStreamingOutRequestTexts) *UpdateStreamingOutRequest
	GetTexts() []*UpdateStreamingOutRequestTexts
}

type UpdateStreamingOutRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId       *string                                 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Backgrounds []*UpdateStreamingOutRequestBackgrounds `json:"Backgrounds,omitempty" xml:"Backgrounds,omitempty" type:"Repeated"`
	BgColor     *UpdateStreamingOutRequestBgColor       `json:"BgColor,omitempty" xml:"BgColor,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId    *string                                  `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ClockWidgets []*UpdateStreamingOutRequestClockWidgets `json:"ClockWidgets,omitempty" xml:"ClockWidgets,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	CropMode             *int32                                         `json:"CropMode,omitempty" xml:"CropMode,omitempty"`
	Images               []*UpdateStreamingOutRequestImages             `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	LayoutSpecifiedUsers *UpdateStreamingOutRequestLayoutSpecifiedUsers `json:"LayoutSpecifiedUsers,omitempty" xml:"LayoutSpecifiedUsers,omitempty" type:"Struct"`
	Panes                []*UpdateStreamingOutRequestPanes              `json:"Panes,omitempty" xml:"Panes,omitempty" type:"Repeated"`
	RegionColor          *UpdateStreamingOutRequestRegionColor          `json:"RegionColor,omitempty" xml:"RegionColor,omitempty" type:"Struct"`
	SpecMixedUserList    []*string                                      `json:"SpecMixedUserList,omitempty" xml:"SpecMixedUserList,omitempty" type:"Repeated"`
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
	TemplateId *string                           `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Texts      []*UpdateStreamingOutRequestTexts `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
}

func (s UpdateStreamingOutRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamingOutRequest) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateStreamingOutRequest) GetBackgrounds() []*UpdateStreamingOutRequestBackgrounds {
	return s.Backgrounds
}

func (s *UpdateStreamingOutRequest) GetBgColor() *UpdateStreamingOutRequestBgColor {
	return s.BgColor
}

func (s *UpdateStreamingOutRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *UpdateStreamingOutRequest) GetClockWidgets() []*UpdateStreamingOutRequestClockWidgets {
	return s.ClockWidgets
}

func (s *UpdateStreamingOutRequest) GetCropMode() *int32 {
	return s.CropMode
}

func (s *UpdateStreamingOutRequest) GetImages() []*UpdateStreamingOutRequestImages {
	return s.Images
}

func (s *UpdateStreamingOutRequest) GetLayoutSpecifiedUsers() *UpdateStreamingOutRequestLayoutSpecifiedUsers {
	return s.LayoutSpecifiedUsers
}

func (s *UpdateStreamingOutRequest) GetPanes() []*UpdateStreamingOutRequestPanes {
	return s.Panes
}

func (s *UpdateStreamingOutRequest) GetRegionColor() *UpdateStreamingOutRequestRegionColor {
	return s.RegionColor
}

func (s *UpdateStreamingOutRequest) GetSpecMixedUserList() []*string {
	return s.SpecMixedUserList
}

func (s *UpdateStreamingOutRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateStreamingOutRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateStreamingOutRequest) GetTexts() []*UpdateStreamingOutRequestTexts {
	return s.Texts
}

func (s *UpdateStreamingOutRequest) SetAppId(v string) *UpdateStreamingOutRequest {
	s.AppId = &v
	return s
}

func (s *UpdateStreamingOutRequest) SetBackgrounds(v []*UpdateStreamingOutRequestBackgrounds) *UpdateStreamingOutRequest {
	s.Backgrounds = v
	return s
}

func (s *UpdateStreamingOutRequest) SetBgColor(v *UpdateStreamingOutRequestBgColor) *UpdateStreamingOutRequest {
	s.BgColor = v
	return s
}

func (s *UpdateStreamingOutRequest) SetChannelId(v string) *UpdateStreamingOutRequest {
	s.ChannelId = &v
	return s
}

func (s *UpdateStreamingOutRequest) SetClockWidgets(v []*UpdateStreamingOutRequestClockWidgets) *UpdateStreamingOutRequest {
	s.ClockWidgets = v
	return s
}

func (s *UpdateStreamingOutRequest) SetCropMode(v int32) *UpdateStreamingOutRequest {
	s.CropMode = &v
	return s
}

func (s *UpdateStreamingOutRequest) SetImages(v []*UpdateStreamingOutRequestImages) *UpdateStreamingOutRequest {
	s.Images = v
	return s
}

func (s *UpdateStreamingOutRequest) SetLayoutSpecifiedUsers(v *UpdateStreamingOutRequestLayoutSpecifiedUsers) *UpdateStreamingOutRequest {
	s.LayoutSpecifiedUsers = v
	return s
}

func (s *UpdateStreamingOutRequest) SetPanes(v []*UpdateStreamingOutRequestPanes) *UpdateStreamingOutRequest {
	s.Panes = v
	return s
}

func (s *UpdateStreamingOutRequest) SetRegionColor(v *UpdateStreamingOutRequestRegionColor) *UpdateStreamingOutRequest {
	s.RegionColor = v
	return s
}

func (s *UpdateStreamingOutRequest) SetSpecMixedUserList(v []*string) *UpdateStreamingOutRequest {
	s.SpecMixedUserList = v
	return s
}

func (s *UpdateStreamingOutRequest) SetTaskId(v string) *UpdateStreamingOutRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateStreamingOutRequest) SetTemplateId(v string) *UpdateStreamingOutRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateStreamingOutRequest) SetTexts(v []*UpdateStreamingOutRequestTexts) *UpdateStreamingOutRequest {
	s.Texts = v
	return s
}

func (s *UpdateStreamingOutRequest) Validate() error {
	if s.Backgrounds != nil {
		for _, item := range s.Backgrounds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.BgColor != nil {
		if err := s.BgColor.Validate(); err != nil {
			return err
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
	if s.RegionColor != nil {
		if err := s.RegionColor.Validate(); err != nil {
			return err
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

type UpdateStreamingOutRequestBackgrounds struct {
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

func (s UpdateStreamingOutRequestBackgrounds) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamingOutRequestBackgrounds) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutRequestBackgrounds) GetAlpha() *float64 {
	return s.Alpha
}

func (s *UpdateStreamingOutRequestBackgrounds) GetBackgroundCropMode() *int32 {
	return s.BackgroundCropMode
}

func (s *UpdateStreamingOutRequestBackgrounds) GetHeight() *float64 {
	return s.Height
}

func (s *UpdateStreamingOutRequestBackgrounds) GetLayer() *int32 {
	return s.Layer
}

func (s *UpdateStreamingOutRequestBackgrounds) GetUrl() *string {
	return s.Url
}

func (s *UpdateStreamingOutRequestBackgrounds) GetWidth() *float64 {
	return s.Width
}

func (s *UpdateStreamingOutRequestBackgrounds) GetX() *float64 {
	return s.X
}

func (s *UpdateStreamingOutRequestBackgrounds) GetY() *float64 {
	return s.Y
}

func (s *UpdateStreamingOutRequestBackgrounds) SetAlpha(v float64) *UpdateStreamingOutRequestBackgrounds {
	s.Alpha = &v
	return s
}

func (s *UpdateStreamingOutRequestBackgrounds) SetBackgroundCropMode(v int32) *UpdateStreamingOutRequestBackgrounds {
	s.BackgroundCropMode = &v
	return s
}

func (s *UpdateStreamingOutRequestBackgrounds) SetHeight(v float64) *UpdateStreamingOutRequestBackgrounds {
	s.Height = &v
	return s
}

func (s *UpdateStreamingOutRequestBackgrounds) SetLayer(v int32) *UpdateStreamingOutRequestBackgrounds {
	s.Layer = &v
	return s
}

func (s *UpdateStreamingOutRequestBackgrounds) SetUrl(v string) *UpdateStreamingOutRequestBackgrounds {
	s.Url = &v
	return s
}

func (s *UpdateStreamingOutRequestBackgrounds) SetWidth(v float64) *UpdateStreamingOutRequestBackgrounds {
	s.Width = &v
	return s
}

func (s *UpdateStreamingOutRequestBackgrounds) SetX(v float64) *UpdateStreamingOutRequestBackgrounds {
	s.X = &v
	return s
}

func (s *UpdateStreamingOutRequestBackgrounds) SetY(v float64) *UpdateStreamingOutRequestBackgrounds {
	s.Y = &v
	return s
}

func (s *UpdateStreamingOutRequestBackgrounds) Validate() error {
	return dara.Validate(s)
}

type UpdateStreamingOutRequestBgColor struct {
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

func (s UpdateStreamingOutRequestBgColor) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamingOutRequestBgColor) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutRequestBgColor) GetB() *int32 {
	return s.B
}

func (s *UpdateStreamingOutRequestBgColor) GetG() *int32 {
	return s.G
}

func (s *UpdateStreamingOutRequestBgColor) GetR() *int32 {
	return s.R
}

func (s *UpdateStreamingOutRequestBgColor) SetB(v int32) *UpdateStreamingOutRequestBgColor {
	s.B = &v
	return s
}

func (s *UpdateStreamingOutRequestBgColor) SetG(v int32) *UpdateStreamingOutRequestBgColor {
	s.G = &v
	return s
}

func (s *UpdateStreamingOutRequestBgColor) SetR(v int32) *UpdateStreamingOutRequestBgColor {
	s.R = &v
	return s
}

func (s *UpdateStreamingOutRequestBgColor) Validate() error {
	return dara.Validate(s)
}

type UpdateStreamingOutRequestClockWidgets struct {
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
	BoxBorderw *int32                                         `json:"BoxBorderw,omitempty" xml:"BoxBorderw,omitempty"`
	BoxColor   *UpdateStreamingOutRequestClockWidgetsBoxColor `json:"BoxColor,omitempty" xml:"BoxColor,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Font      *int32                                          `json:"Font,omitempty" xml:"Font,omitempty"`
	FontColor *UpdateStreamingOutRequestClockWidgetsFontColor `json:"FontColor,omitempty" xml:"FontColor,omitempty" type:"Struct"`
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

func (s UpdateStreamingOutRequestClockWidgets) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamingOutRequestClockWidgets) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutRequestClockWidgets) GetAlpha() *float64 {
	return s.Alpha
}

func (s *UpdateStreamingOutRequestClockWidgets) GetBoxAlpha() *float64 {
	return s.BoxAlpha
}

func (s *UpdateStreamingOutRequestClockWidgets) GetBoxBorderw() *int32 {
	return s.BoxBorderw
}

func (s *UpdateStreamingOutRequestClockWidgets) GetBoxColor() *UpdateStreamingOutRequestClockWidgetsBoxColor {
	return s.BoxColor
}

func (s *UpdateStreamingOutRequestClockWidgets) GetFont() *int32 {
	return s.Font
}

func (s *UpdateStreamingOutRequestClockWidgets) GetFontColor() *UpdateStreamingOutRequestClockWidgetsFontColor {
	return s.FontColor
}

func (s *UpdateStreamingOutRequestClockWidgets) GetFontSize() *int32 {
	return s.FontSize
}

func (s *UpdateStreamingOutRequestClockWidgets) GetHasBox() *bool {
	return s.HasBox
}

func (s *UpdateStreamingOutRequestClockWidgets) GetLayer() *int32 {
	return s.Layer
}

func (s *UpdateStreamingOutRequestClockWidgets) GetX() *float64 {
	return s.X
}

func (s *UpdateStreamingOutRequestClockWidgets) GetY() *float64 {
	return s.Y
}

func (s *UpdateStreamingOutRequestClockWidgets) GetZone() *int32 {
	return s.Zone
}

func (s *UpdateStreamingOutRequestClockWidgets) SetAlpha(v float64) *UpdateStreamingOutRequestClockWidgets {
	s.Alpha = &v
	return s
}

func (s *UpdateStreamingOutRequestClockWidgets) SetBoxAlpha(v float64) *UpdateStreamingOutRequestClockWidgets {
	s.BoxAlpha = &v
	return s
}

func (s *UpdateStreamingOutRequestClockWidgets) SetBoxBorderw(v int32) *UpdateStreamingOutRequestClockWidgets {
	s.BoxBorderw = &v
	return s
}

func (s *UpdateStreamingOutRequestClockWidgets) SetBoxColor(v *UpdateStreamingOutRequestClockWidgetsBoxColor) *UpdateStreamingOutRequestClockWidgets {
	s.BoxColor = v
	return s
}

func (s *UpdateStreamingOutRequestClockWidgets) SetFont(v int32) *UpdateStreamingOutRequestClockWidgets {
	s.Font = &v
	return s
}

func (s *UpdateStreamingOutRequestClockWidgets) SetFontColor(v *UpdateStreamingOutRequestClockWidgetsFontColor) *UpdateStreamingOutRequestClockWidgets {
	s.FontColor = v
	return s
}

func (s *UpdateStreamingOutRequestClockWidgets) SetFontSize(v int32) *UpdateStreamingOutRequestClockWidgets {
	s.FontSize = &v
	return s
}

func (s *UpdateStreamingOutRequestClockWidgets) SetHasBox(v bool) *UpdateStreamingOutRequestClockWidgets {
	s.HasBox = &v
	return s
}

func (s *UpdateStreamingOutRequestClockWidgets) SetLayer(v int32) *UpdateStreamingOutRequestClockWidgets {
	s.Layer = &v
	return s
}

func (s *UpdateStreamingOutRequestClockWidgets) SetX(v float64) *UpdateStreamingOutRequestClockWidgets {
	s.X = &v
	return s
}

func (s *UpdateStreamingOutRequestClockWidgets) SetY(v float64) *UpdateStreamingOutRequestClockWidgets {
	s.Y = &v
	return s
}

func (s *UpdateStreamingOutRequestClockWidgets) SetZone(v int32) *UpdateStreamingOutRequestClockWidgets {
	s.Zone = &v
	return s
}

func (s *UpdateStreamingOutRequestClockWidgets) Validate() error {
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

type UpdateStreamingOutRequestClockWidgetsBoxColor struct {
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

func (s UpdateStreamingOutRequestClockWidgetsBoxColor) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamingOutRequestClockWidgetsBoxColor) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutRequestClockWidgetsBoxColor) GetB() *int32 {
	return s.B
}

func (s *UpdateStreamingOutRequestClockWidgetsBoxColor) GetG() *int32 {
	return s.G
}

func (s *UpdateStreamingOutRequestClockWidgetsBoxColor) GetR() *int32 {
	return s.R
}

func (s *UpdateStreamingOutRequestClockWidgetsBoxColor) SetB(v int32) *UpdateStreamingOutRequestClockWidgetsBoxColor {
	s.B = &v
	return s
}

func (s *UpdateStreamingOutRequestClockWidgetsBoxColor) SetG(v int32) *UpdateStreamingOutRequestClockWidgetsBoxColor {
	s.G = &v
	return s
}

func (s *UpdateStreamingOutRequestClockWidgetsBoxColor) SetR(v int32) *UpdateStreamingOutRequestClockWidgetsBoxColor {
	s.R = &v
	return s
}

func (s *UpdateStreamingOutRequestClockWidgetsBoxColor) Validate() error {
	return dara.Validate(s)
}

type UpdateStreamingOutRequestClockWidgetsFontColor struct {
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

func (s UpdateStreamingOutRequestClockWidgetsFontColor) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamingOutRequestClockWidgetsFontColor) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutRequestClockWidgetsFontColor) GetB() *int32 {
	return s.B
}

func (s *UpdateStreamingOutRequestClockWidgetsFontColor) GetG() *int32 {
	return s.G
}

func (s *UpdateStreamingOutRequestClockWidgetsFontColor) GetR() *int32 {
	return s.R
}

func (s *UpdateStreamingOutRequestClockWidgetsFontColor) SetB(v int32) *UpdateStreamingOutRequestClockWidgetsFontColor {
	s.B = &v
	return s
}

func (s *UpdateStreamingOutRequestClockWidgetsFontColor) SetG(v int32) *UpdateStreamingOutRequestClockWidgetsFontColor {
	s.G = &v
	return s
}

func (s *UpdateStreamingOutRequestClockWidgetsFontColor) SetR(v int32) *UpdateStreamingOutRequestClockWidgetsFontColor {
	s.R = &v
	return s
}

func (s *UpdateStreamingOutRequestClockWidgetsFontColor) Validate() error {
	return dara.Validate(s)
}

type UpdateStreamingOutRequestImages struct {
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

func (s UpdateStreamingOutRequestImages) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamingOutRequestImages) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutRequestImages) GetAlpha() *float64 {
	return s.Alpha
}

func (s *UpdateStreamingOutRequestImages) GetHeight() *float64 {
	return s.Height
}

func (s *UpdateStreamingOutRequestImages) GetImageCropMode() *int32 {
	return s.ImageCropMode
}

func (s *UpdateStreamingOutRequestImages) GetLayer() *int32 {
	return s.Layer
}

func (s *UpdateStreamingOutRequestImages) GetUrl() *string {
	return s.Url
}

func (s *UpdateStreamingOutRequestImages) GetWidth() *float64 {
	return s.Width
}

func (s *UpdateStreamingOutRequestImages) GetX() *float64 {
	return s.X
}

func (s *UpdateStreamingOutRequestImages) GetY() *float64 {
	return s.Y
}

func (s *UpdateStreamingOutRequestImages) SetAlpha(v float64) *UpdateStreamingOutRequestImages {
	s.Alpha = &v
	return s
}

func (s *UpdateStreamingOutRequestImages) SetHeight(v float64) *UpdateStreamingOutRequestImages {
	s.Height = &v
	return s
}

func (s *UpdateStreamingOutRequestImages) SetImageCropMode(v int32) *UpdateStreamingOutRequestImages {
	s.ImageCropMode = &v
	return s
}

func (s *UpdateStreamingOutRequestImages) SetLayer(v int32) *UpdateStreamingOutRequestImages {
	s.Layer = &v
	return s
}

func (s *UpdateStreamingOutRequestImages) SetUrl(v string) *UpdateStreamingOutRequestImages {
	s.Url = &v
	return s
}

func (s *UpdateStreamingOutRequestImages) SetWidth(v float64) *UpdateStreamingOutRequestImages {
	s.Width = &v
	return s
}

func (s *UpdateStreamingOutRequestImages) SetX(v float64) *UpdateStreamingOutRequestImages {
	s.X = &v
	return s
}

func (s *UpdateStreamingOutRequestImages) SetY(v float64) *UpdateStreamingOutRequestImages {
	s.Y = &v
	return s
}

func (s *UpdateStreamingOutRequestImages) Validate() error {
	return dara.Validate(s)
}

type UpdateStreamingOutRequestLayoutSpecifiedUsers struct {
	// This parameter is required.
	Ids []*string `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateStreamingOutRequestLayoutSpecifiedUsers) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamingOutRequestLayoutSpecifiedUsers) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutRequestLayoutSpecifiedUsers) GetIds() []*string {
	return s.Ids
}

func (s *UpdateStreamingOutRequestLayoutSpecifiedUsers) GetType() *string {
	return s.Type
}

func (s *UpdateStreamingOutRequestLayoutSpecifiedUsers) SetIds(v []*string) *UpdateStreamingOutRequestLayoutSpecifiedUsers {
	s.Ids = v
	return s
}

func (s *UpdateStreamingOutRequestLayoutSpecifiedUsers) SetType(v string) *UpdateStreamingOutRequestLayoutSpecifiedUsers {
	s.Type = &v
	return s
}

func (s *UpdateStreamingOutRequestLayoutSpecifiedUsers) Validate() error {
	return dara.Validate(s)
}

type UpdateStreamingOutRequestPanes struct {
	Backgrounds []*UpdateStreamingOutRequestPanesBackgrounds `json:"Backgrounds,omitempty" xml:"Backgrounds,omitempty" type:"Repeated"`
	Images      []*UpdateStreamingOutRequestPanesImages      `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
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
	SourceType *string                                `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Texts      []*UpdateStreamingOutRequestPanesTexts `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
	// example:
	//
	// cameraFirst
	VideoOrder *string                                   `json:"VideoOrder,omitempty" xml:"VideoOrder,omitempty"`
	Whiteboard *UpdateStreamingOutRequestPanesWhiteboard `json:"Whiteboard,omitempty" xml:"Whiteboard,omitempty" type:"Struct"`
}

func (s UpdateStreamingOutRequestPanes) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamingOutRequestPanes) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutRequestPanes) GetBackgrounds() []*UpdateStreamingOutRequestPanesBackgrounds {
	return s.Backgrounds
}

func (s *UpdateStreamingOutRequestPanes) GetImages() []*UpdateStreamingOutRequestPanesImages {
	return s.Images
}

func (s *UpdateStreamingOutRequestPanes) GetPaneCropMode() *int32 {
	return s.PaneCropMode
}

func (s *UpdateStreamingOutRequestPanes) GetPaneId() *int32 {
	return s.PaneId
}

func (s *UpdateStreamingOutRequestPanes) GetReservePaneForOfflineUser() *bool {
	return s.ReservePaneForOfflineUser
}

func (s *UpdateStreamingOutRequestPanes) GetSource() *string {
	return s.Source
}

func (s *UpdateStreamingOutRequestPanes) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateStreamingOutRequestPanes) GetTexts() []*UpdateStreamingOutRequestPanesTexts {
	return s.Texts
}

func (s *UpdateStreamingOutRequestPanes) GetVideoOrder() *string {
	return s.VideoOrder
}

func (s *UpdateStreamingOutRequestPanes) GetWhiteboard() *UpdateStreamingOutRequestPanesWhiteboard {
	return s.Whiteboard
}

func (s *UpdateStreamingOutRequestPanes) SetBackgrounds(v []*UpdateStreamingOutRequestPanesBackgrounds) *UpdateStreamingOutRequestPanes {
	s.Backgrounds = v
	return s
}

func (s *UpdateStreamingOutRequestPanes) SetImages(v []*UpdateStreamingOutRequestPanesImages) *UpdateStreamingOutRequestPanes {
	s.Images = v
	return s
}

func (s *UpdateStreamingOutRequestPanes) SetPaneCropMode(v int32) *UpdateStreamingOutRequestPanes {
	s.PaneCropMode = &v
	return s
}

func (s *UpdateStreamingOutRequestPanes) SetPaneId(v int32) *UpdateStreamingOutRequestPanes {
	s.PaneId = &v
	return s
}

func (s *UpdateStreamingOutRequestPanes) SetReservePaneForOfflineUser(v bool) *UpdateStreamingOutRequestPanes {
	s.ReservePaneForOfflineUser = &v
	return s
}

func (s *UpdateStreamingOutRequestPanes) SetSource(v string) *UpdateStreamingOutRequestPanes {
	s.Source = &v
	return s
}

func (s *UpdateStreamingOutRequestPanes) SetSourceType(v string) *UpdateStreamingOutRequestPanes {
	s.SourceType = &v
	return s
}

func (s *UpdateStreamingOutRequestPanes) SetTexts(v []*UpdateStreamingOutRequestPanesTexts) *UpdateStreamingOutRequestPanes {
	s.Texts = v
	return s
}

func (s *UpdateStreamingOutRequestPanes) SetVideoOrder(v string) *UpdateStreamingOutRequestPanes {
	s.VideoOrder = &v
	return s
}

func (s *UpdateStreamingOutRequestPanes) SetWhiteboard(v *UpdateStreamingOutRequestPanesWhiteboard) *UpdateStreamingOutRequestPanes {
	s.Whiteboard = v
	return s
}

func (s *UpdateStreamingOutRequestPanes) Validate() error {
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

type UpdateStreamingOutRequestPanesBackgrounds struct {
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

func (s UpdateStreamingOutRequestPanesBackgrounds) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamingOutRequestPanesBackgrounds) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutRequestPanesBackgrounds) GetAlpha() *float64 {
	return s.Alpha
}

func (s *UpdateStreamingOutRequestPanesBackgrounds) GetDisplay() *string {
	return s.Display
}

func (s *UpdateStreamingOutRequestPanesBackgrounds) GetHeight() *float64 {
	return s.Height
}

func (s *UpdateStreamingOutRequestPanesBackgrounds) GetLayer() *int32 {
	return s.Layer
}

func (s *UpdateStreamingOutRequestPanesBackgrounds) GetPaneBackgroundCropMode() *int32 {
	return s.PaneBackgroundCropMode
}

func (s *UpdateStreamingOutRequestPanesBackgrounds) GetUrl() *string {
	return s.Url
}

func (s *UpdateStreamingOutRequestPanesBackgrounds) GetWidth() *float64 {
	return s.Width
}

func (s *UpdateStreamingOutRequestPanesBackgrounds) GetX() *float64 {
	return s.X
}

func (s *UpdateStreamingOutRequestPanesBackgrounds) GetY() *float64 {
	return s.Y
}

func (s *UpdateStreamingOutRequestPanesBackgrounds) SetAlpha(v float64) *UpdateStreamingOutRequestPanesBackgrounds {
	s.Alpha = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesBackgrounds) SetDisplay(v string) *UpdateStreamingOutRequestPanesBackgrounds {
	s.Display = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesBackgrounds) SetHeight(v float64) *UpdateStreamingOutRequestPanesBackgrounds {
	s.Height = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesBackgrounds) SetLayer(v int32) *UpdateStreamingOutRequestPanesBackgrounds {
	s.Layer = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesBackgrounds) SetPaneBackgroundCropMode(v int32) *UpdateStreamingOutRequestPanesBackgrounds {
	s.PaneBackgroundCropMode = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesBackgrounds) SetUrl(v string) *UpdateStreamingOutRequestPanesBackgrounds {
	s.Url = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesBackgrounds) SetWidth(v float64) *UpdateStreamingOutRequestPanesBackgrounds {
	s.Width = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesBackgrounds) SetX(v float64) *UpdateStreamingOutRequestPanesBackgrounds {
	s.X = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesBackgrounds) SetY(v float64) *UpdateStreamingOutRequestPanesBackgrounds {
	s.Y = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesBackgrounds) Validate() error {
	return dara.Validate(s)
}

type UpdateStreamingOutRequestPanesImages struct {
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

func (s UpdateStreamingOutRequestPanesImages) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamingOutRequestPanesImages) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutRequestPanesImages) GetAlpha() *float64 {
	return s.Alpha
}

func (s *UpdateStreamingOutRequestPanesImages) GetDisplay() *string {
	return s.Display
}

func (s *UpdateStreamingOutRequestPanesImages) GetHeight() *float64 {
	return s.Height
}

func (s *UpdateStreamingOutRequestPanesImages) GetLayer() *int32 {
	return s.Layer
}

func (s *UpdateStreamingOutRequestPanesImages) GetPaneImageCropMode() *int32 {
	return s.PaneImageCropMode
}

func (s *UpdateStreamingOutRequestPanesImages) GetUrl() *string {
	return s.Url
}

func (s *UpdateStreamingOutRequestPanesImages) GetWidth() *float64 {
	return s.Width
}

func (s *UpdateStreamingOutRequestPanesImages) GetX() *float64 {
	return s.X
}

func (s *UpdateStreamingOutRequestPanesImages) GetY() *float64 {
	return s.Y
}

func (s *UpdateStreamingOutRequestPanesImages) SetAlpha(v float64) *UpdateStreamingOutRequestPanesImages {
	s.Alpha = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesImages) SetDisplay(v string) *UpdateStreamingOutRequestPanesImages {
	s.Display = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesImages) SetHeight(v float64) *UpdateStreamingOutRequestPanesImages {
	s.Height = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesImages) SetLayer(v int32) *UpdateStreamingOutRequestPanesImages {
	s.Layer = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesImages) SetPaneImageCropMode(v int32) *UpdateStreamingOutRequestPanesImages {
	s.PaneImageCropMode = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesImages) SetUrl(v string) *UpdateStreamingOutRequestPanesImages {
	s.Url = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesImages) SetWidth(v float64) *UpdateStreamingOutRequestPanesImages {
	s.Width = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesImages) SetX(v float64) *UpdateStreamingOutRequestPanesImages {
	s.X = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesImages) SetY(v float64) *UpdateStreamingOutRequestPanesImages {
	s.Y = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesImages) Validate() error {
	return dara.Validate(s)
}

type UpdateStreamingOutRequestPanesTexts struct {
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
	BoxColor   *UpdateStreamingOutRequestPanesTextsBoxColor `json:"BoxColor,omitempty" xml:"BoxColor,omitempty" type:"Struct"`
	// example:
	//
	// backup
	Display *string `json:"Display,omitempty" xml:"Display,omitempty"`
	// example:
	//
	// 0
	Font      *int32                                        `json:"Font,omitempty" xml:"Font,omitempty"`
	FontColor *UpdateStreamingOutRequestPanesTextsFontColor `json:"FontColor,omitempty" xml:"FontColor,omitempty" type:"Struct"`
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

func (s UpdateStreamingOutRequestPanesTexts) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamingOutRequestPanesTexts) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutRequestPanesTexts) GetAlpha() *float64 {
	return s.Alpha
}

func (s *UpdateStreamingOutRequestPanesTexts) GetBoxAlpha() *float64 {
	return s.BoxAlpha
}

func (s *UpdateStreamingOutRequestPanesTexts) GetBoxBorderw() *int32 {
	return s.BoxBorderw
}

func (s *UpdateStreamingOutRequestPanesTexts) GetBoxColor() *UpdateStreamingOutRequestPanesTextsBoxColor {
	return s.BoxColor
}

func (s *UpdateStreamingOutRequestPanesTexts) GetDisplay() *string {
	return s.Display
}

func (s *UpdateStreamingOutRequestPanesTexts) GetFont() *int32 {
	return s.Font
}

func (s *UpdateStreamingOutRequestPanesTexts) GetFontColor() *UpdateStreamingOutRequestPanesTextsFontColor {
	return s.FontColor
}

func (s *UpdateStreamingOutRequestPanesTexts) GetFontSize() *int32 {
	return s.FontSize
}

func (s *UpdateStreamingOutRequestPanesTexts) GetHasBox() *bool {
	return s.HasBox
}

func (s *UpdateStreamingOutRequestPanesTexts) GetLayer() *int32 {
	return s.Layer
}

func (s *UpdateStreamingOutRequestPanesTexts) GetTexture() *string {
	return s.Texture
}

func (s *UpdateStreamingOutRequestPanesTexts) GetX() *float64 {
	return s.X
}

func (s *UpdateStreamingOutRequestPanesTexts) GetY() *float64 {
	return s.Y
}

func (s *UpdateStreamingOutRequestPanesTexts) SetAlpha(v float64) *UpdateStreamingOutRequestPanesTexts {
	s.Alpha = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesTexts) SetBoxAlpha(v float64) *UpdateStreamingOutRequestPanesTexts {
	s.BoxAlpha = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesTexts) SetBoxBorderw(v int32) *UpdateStreamingOutRequestPanesTexts {
	s.BoxBorderw = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesTexts) SetBoxColor(v *UpdateStreamingOutRequestPanesTextsBoxColor) *UpdateStreamingOutRequestPanesTexts {
	s.BoxColor = v
	return s
}

func (s *UpdateStreamingOutRequestPanesTexts) SetDisplay(v string) *UpdateStreamingOutRequestPanesTexts {
	s.Display = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesTexts) SetFont(v int32) *UpdateStreamingOutRequestPanesTexts {
	s.Font = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesTexts) SetFontColor(v *UpdateStreamingOutRequestPanesTextsFontColor) *UpdateStreamingOutRequestPanesTexts {
	s.FontColor = v
	return s
}

func (s *UpdateStreamingOutRequestPanesTexts) SetFontSize(v int32) *UpdateStreamingOutRequestPanesTexts {
	s.FontSize = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesTexts) SetHasBox(v bool) *UpdateStreamingOutRequestPanesTexts {
	s.HasBox = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesTexts) SetLayer(v int32) *UpdateStreamingOutRequestPanesTexts {
	s.Layer = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesTexts) SetTexture(v string) *UpdateStreamingOutRequestPanesTexts {
	s.Texture = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesTexts) SetX(v float64) *UpdateStreamingOutRequestPanesTexts {
	s.X = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesTexts) SetY(v float64) *UpdateStreamingOutRequestPanesTexts {
	s.Y = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesTexts) Validate() error {
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

type UpdateStreamingOutRequestPanesTextsBoxColor struct {
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

func (s UpdateStreamingOutRequestPanesTextsBoxColor) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamingOutRequestPanesTextsBoxColor) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutRequestPanesTextsBoxColor) GetB() *int32 {
	return s.B
}

func (s *UpdateStreamingOutRequestPanesTextsBoxColor) GetG() *int32 {
	return s.G
}

func (s *UpdateStreamingOutRequestPanesTextsBoxColor) GetR() *int32 {
	return s.R
}

func (s *UpdateStreamingOutRequestPanesTextsBoxColor) SetB(v int32) *UpdateStreamingOutRequestPanesTextsBoxColor {
	s.B = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesTextsBoxColor) SetG(v int32) *UpdateStreamingOutRequestPanesTextsBoxColor {
	s.G = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesTextsBoxColor) SetR(v int32) *UpdateStreamingOutRequestPanesTextsBoxColor {
	s.R = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesTextsBoxColor) Validate() error {
	return dara.Validate(s)
}

type UpdateStreamingOutRequestPanesTextsFontColor struct {
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

func (s UpdateStreamingOutRequestPanesTextsFontColor) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamingOutRequestPanesTextsFontColor) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutRequestPanesTextsFontColor) GetB() *int32 {
	return s.B
}

func (s *UpdateStreamingOutRequestPanesTextsFontColor) GetG() *int32 {
	return s.G
}

func (s *UpdateStreamingOutRequestPanesTextsFontColor) GetR() *int32 {
	return s.R
}

func (s *UpdateStreamingOutRequestPanesTextsFontColor) SetB(v int32) *UpdateStreamingOutRequestPanesTextsFontColor {
	s.B = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesTextsFontColor) SetG(v int32) *UpdateStreamingOutRequestPanesTextsFontColor {
	s.G = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesTextsFontColor) SetR(v int32) *UpdateStreamingOutRequestPanesTextsFontColor {
	s.R = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesTextsFontColor) Validate() error {
	return dara.Validate(s)
}

type UpdateStreamingOutRequestPanesWhiteboard struct {
	// example:
	//
	// default
	WhiteboardId *string `json:"WhiteboardId,omitempty" xml:"WhiteboardId,omitempty"`
}

func (s UpdateStreamingOutRequestPanesWhiteboard) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamingOutRequestPanesWhiteboard) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutRequestPanesWhiteboard) GetWhiteboardId() *string {
	return s.WhiteboardId
}

func (s *UpdateStreamingOutRequestPanesWhiteboard) SetWhiteboardId(v string) *UpdateStreamingOutRequestPanesWhiteboard {
	s.WhiteboardId = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesWhiteboard) Validate() error {
	return dara.Validate(s)
}

type UpdateStreamingOutRequestRegionColor struct {
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

func (s UpdateStreamingOutRequestRegionColor) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamingOutRequestRegionColor) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutRequestRegionColor) GetB() *int32 {
	return s.B
}

func (s *UpdateStreamingOutRequestRegionColor) GetG() *int32 {
	return s.G
}

func (s *UpdateStreamingOutRequestRegionColor) GetR() *int32 {
	return s.R
}

func (s *UpdateStreamingOutRequestRegionColor) SetB(v int32) *UpdateStreamingOutRequestRegionColor {
	s.B = &v
	return s
}

func (s *UpdateStreamingOutRequestRegionColor) SetG(v int32) *UpdateStreamingOutRequestRegionColor {
	s.G = &v
	return s
}

func (s *UpdateStreamingOutRequestRegionColor) SetR(v int32) *UpdateStreamingOutRequestRegionColor {
	s.R = &v
	return s
}

func (s *UpdateStreamingOutRequestRegionColor) Validate() error {
	return dara.Validate(s)
}

type UpdateStreamingOutRequestTexts struct {
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
	BoxBorderw *int32                                  `json:"BoxBorderw,omitempty" xml:"BoxBorderw,omitempty"`
	BoxColor   *UpdateStreamingOutRequestTextsBoxColor `json:"BoxColor,omitempty" xml:"BoxColor,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Font      *int32                                   `json:"Font,omitempty" xml:"Font,omitempty"`
	FontColor *UpdateStreamingOutRequestTextsFontColor `json:"FontColor,omitempty" xml:"FontColor,omitempty" type:"Struct"`
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

func (s UpdateStreamingOutRequestTexts) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamingOutRequestTexts) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutRequestTexts) GetAlpha() *float64 {
	return s.Alpha
}

func (s *UpdateStreamingOutRequestTexts) GetBoxAlpha() *float64 {
	return s.BoxAlpha
}

func (s *UpdateStreamingOutRequestTexts) GetBoxBorderw() *int32 {
	return s.BoxBorderw
}

func (s *UpdateStreamingOutRequestTexts) GetBoxColor() *UpdateStreamingOutRequestTextsBoxColor {
	return s.BoxColor
}

func (s *UpdateStreamingOutRequestTexts) GetFont() *int32 {
	return s.Font
}

func (s *UpdateStreamingOutRequestTexts) GetFontColor() *UpdateStreamingOutRequestTextsFontColor {
	return s.FontColor
}

func (s *UpdateStreamingOutRequestTexts) GetFontSize() *int32 {
	return s.FontSize
}

func (s *UpdateStreamingOutRequestTexts) GetHasBox() *bool {
	return s.HasBox
}

func (s *UpdateStreamingOutRequestTexts) GetLayer() *int32 {
	return s.Layer
}

func (s *UpdateStreamingOutRequestTexts) GetTexture() *string {
	return s.Texture
}

func (s *UpdateStreamingOutRequestTexts) GetX() *float64 {
	return s.X
}

func (s *UpdateStreamingOutRequestTexts) GetY() *float64 {
	return s.Y
}

func (s *UpdateStreamingOutRequestTexts) SetAlpha(v float64) *UpdateStreamingOutRequestTexts {
	s.Alpha = &v
	return s
}

func (s *UpdateStreamingOutRequestTexts) SetBoxAlpha(v float64) *UpdateStreamingOutRequestTexts {
	s.BoxAlpha = &v
	return s
}

func (s *UpdateStreamingOutRequestTexts) SetBoxBorderw(v int32) *UpdateStreamingOutRequestTexts {
	s.BoxBorderw = &v
	return s
}

func (s *UpdateStreamingOutRequestTexts) SetBoxColor(v *UpdateStreamingOutRequestTextsBoxColor) *UpdateStreamingOutRequestTexts {
	s.BoxColor = v
	return s
}

func (s *UpdateStreamingOutRequestTexts) SetFont(v int32) *UpdateStreamingOutRequestTexts {
	s.Font = &v
	return s
}

func (s *UpdateStreamingOutRequestTexts) SetFontColor(v *UpdateStreamingOutRequestTextsFontColor) *UpdateStreamingOutRequestTexts {
	s.FontColor = v
	return s
}

func (s *UpdateStreamingOutRequestTexts) SetFontSize(v int32) *UpdateStreamingOutRequestTexts {
	s.FontSize = &v
	return s
}

func (s *UpdateStreamingOutRequestTexts) SetHasBox(v bool) *UpdateStreamingOutRequestTexts {
	s.HasBox = &v
	return s
}

func (s *UpdateStreamingOutRequestTexts) SetLayer(v int32) *UpdateStreamingOutRequestTexts {
	s.Layer = &v
	return s
}

func (s *UpdateStreamingOutRequestTexts) SetTexture(v string) *UpdateStreamingOutRequestTexts {
	s.Texture = &v
	return s
}

func (s *UpdateStreamingOutRequestTexts) SetX(v float64) *UpdateStreamingOutRequestTexts {
	s.X = &v
	return s
}

func (s *UpdateStreamingOutRequestTexts) SetY(v float64) *UpdateStreamingOutRequestTexts {
	s.Y = &v
	return s
}

func (s *UpdateStreamingOutRequestTexts) Validate() error {
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

type UpdateStreamingOutRequestTextsBoxColor struct {
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

func (s UpdateStreamingOutRequestTextsBoxColor) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamingOutRequestTextsBoxColor) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutRequestTextsBoxColor) GetB() *int32 {
	return s.B
}

func (s *UpdateStreamingOutRequestTextsBoxColor) GetG() *int32 {
	return s.G
}

func (s *UpdateStreamingOutRequestTextsBoxColor) GetR() *int32 {
	return s.R
}

func (s *UpdateStreamingOutRequestTextsBoxColor) SetB(v int32) *UpdateStreamingOutRequestTextsBoxColor {
	s.B = &v
	return s
}

func (s *UpdateStreamingOutRequestTextsBoxColor) SetG(v int32) *UpdateStreamingOutRequestTextsBoxColor {
	s.G = &v
	return s
}

func (s *UpdateStreamingOutRequestTextsBoxColor) SetR(v int32) *UpdateStreamingOutRequestTextsBoxColor {
	s.R = &v
	return s
}

func (s *UpdateStreamingOutRequestTextsBoxColor) Validate() error {
	return dara.Validate(s)
}

type UpdateStreamingOutRequestTextsFontColor struct {
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

func (s UpdateStreamingOutRequestTextsFontColor) String() string {
	return dara.Prettify(s)
}

func (s UpdateStreamingOutRequestTextsFontColor) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutRequestTextsFontColor) GetB() *int32 {
	return s.B
}

func (s *UpdateStreamingOutRequestTextsFontColor) GetG() *int32 {
	return s.G
}

func (s *UpdateStreamingOutRequestTextsFontColor) GetR() *int32 {
	return s.R
}

func (s *UpdateStreamingOutRequestTextsFontColor) SetB(v int32) *UpdateStreamingOutRequestTextsFontColor {
	s.B = &v
	return s
}

func (s *UpdateStreamingOutRequestTextsFontColor) SetG(v int32) *UpdateStreamingOutRequestTextsFontColor {
	s.G = &v
	return s
}

func (s *UpdateStreamingOutRequestTextsFontColor) SetR(v int32) *UpdateStreamingOutRequestTextsFontColor {
	s.R = &v
	return s
}

func (s *UpdateStreamingOutRequestTextsFontColor) Validate() error {
	return dara.Validate(s)
}
