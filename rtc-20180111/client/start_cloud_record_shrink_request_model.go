// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCloudRecordShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnnotation(v string) *StartCloudRecordShrinkRequest
	GetAnnotation() *string
	SetAppId(v string) *StartCloudRecordShrinkRequest
	GetAppId() *string
	SetBackgrounds(v []*StartCloudRecordShrinkRequestBackgrounds) *StartCloudRecordShrinkRequest
	GetBackgrounds() []*StartCloudRecordShrinkRequestBackgrounds
	SetBgColor(v *StartCloudRecordShrinkRequestBgColor) *StartCloudRecordShrinkRequest
	GetBgColor() *StartCloudRecordShrinkRequestBgColor
	SetChannelId(v string) *StartCloudRecordShrinkRequest
	GetChannelId() *string
	SetClockWidgets(v []*StartCloudRecordShrinkRequestClockWidgets) *StartCloudRecordShrinkRequest
	GetClockWidgets() []*StartCloudRecordShrinkRequestClockWidgets
	SetCropMode(v int32) *StartCloudRecordShrinkRequest
	GetCropMode() *int32
	SetImages(v []*StartCloudRecordShrinkRequestImages) *StartCloudRecordShrinkRequest
	GetImages() []*StartCloudRecordShrinkRequestImages
	SetLayoutSpecifiedUsersShrink(v string) *StartCloudRecordShrinkRequest
	GetLayoutSpecifiedUsersShrink() *string
	SetPanes(v []*StartCloudRecordShrinkRequestPanes) *StartCloudRecordShrinkRequest
	GetPanes() []*StartCloudRecordShrinkRequestPanes
	SetRecordMode(v int32) *StartCloudRecordShrinkRequest
	GetRecordMode() *int32
	SetRegionColor(v *StartCloudRecordShrinkRequestRegionColor) *StartCloudRecordShrinkRequest
	GetRegionColor() *StartCloudRecordShrinkRequestRegionColor
	SetReservePaneForNoCameraUser(v bool) *StartCloudRecordShrinkRequest
	GetReservePaneForNoCameraUser() *bool
	SetShowDefaultBackgroundOnMute(v bool) *StartCloudRecordShrinkRequest
	GetShowDefaultBackgroundOnMute() *bool
	SetSingleStreamingRecordShrink(v string) *StartCloudRecordShrinkRequest
	GetSingleStreamingRecordShrink() *string
	SetStartWithoutChannel(v bool) *StartCloudRecordShrinkRequest
	GetStartWithoutChannel() *bool
	SetStartWithoutChannelWaitTime(v int32) *StartCloudRecordShrinkRequest
	GetStartWithoutChannelWaitTime() *int32
	SetStorageConfig(v *StartCloudRecordShrinkRequestStorageConfig) *StartCloudRecordShrinkRequest
	GetStorageConfig() *StartCloudRecordShrinkRequestStorageConfig
	SetSubHighResolutionStream(v bool) *StartCloudRecordShrinkRequest
	GetSubHighResolutionStream() *bool
	SetTaskId(v string) *StartCloudRecordShrinkRequest
	GetTaskId() *string
	SetTemplateId(v string) *StartCloudRecordShrinkRequest
	GetTemplateId() *string
	SetTexts(v []*StartCloudRecordShrinkRequestTexts) *StartCloudRecordShrinkRequest
	GetTexts() []*StartCloudRecordShrinkRequestTexts
}

type StartCloudRecordShrinkRequest struct {
	// example:
	//
	// disable
	Annotation *string `json:"Annotation,omitempty" xml:"Annotation,omitempty"`
	// appId
	//
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId       *string                                     `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Backgrounds []*StartCloudRecordShrinkRequestBackgrounds `json:"Backgrounds,omitempty" xml:"Backgrounds,omitempty" type:"Repeated"`
	BgColor     *StartCloudRecordShrinkRequestBgColor       `json:"BgColor,omitempty" xml:"BgColor,omitempty" type:"Struct"`
	// channelName
	//
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId    *string                                      `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ClockWidgets []*StartCloudRecordShrinkRequestClockWidgets `json:"ClockWidgets,omitempty" xml:"ClockWidgets,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	CropMode                   *int32                                 `json:"CropMode,omitempty" xml:"CropMode,omitempty"`
	Images                     []*StartCloudRecordShrinkRequestImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	LayoutSpecifiedUsersShrink *string                                `json:"LayoutSpecifiedUsers,omitempty" xml:"LayoutSpecifiedUsers,omitempty"`
	// panes
	Panes                       []*StartCloudRecordShrinkRequestPanes     `json:"Panes,omitempty" xml:"Panes,omitempty" type:"Repeated"`
	RecordMode                  *int32                                    `json:"RecordMode,omitempty" xml:"RecordMode,omitempty"`
	RegionColor                 *StartCloudRecordShrinkRequestRegionColor `json:"RegionColor,omitempty" xml:"RegionColor,omitempty" type:"Struct"`
	ReservePaneForNoCameraUser  *bool                                     `json:"ReservePaneForNoCameraUser,omitempty" xml:"ReservePaneForNoCameraUser,omitempty"`
	ShowDefaultBackgroundOnMute *bool                                     `json:"ShowDefaultBackgroundOnMute,omitempty" xml:"ShowDefaultBackgroundOnMute,omitempty"`
	SingleStreamingRecordShrink *string                                   `json:"SingleStreamingRecord,omitempty" xml:"SingleStreamingRecord,omitempty"`
	StartWithoutChannel         *bool                                     `json:"StartWithoutChannel,omitempty" xml:"StartWithoutChannel,omitempty"`
	// example:
	//
	// 30
	StartWithoutChannelWaitTime *int32 `json:"StartWithoutChannelWaitTime,omitempty" xml:"StartWithoutChannelWaitTime,omitempty"`
	// storageConfig
	//
	// This parameter is required.
	StorageConfig           *StartCloudRecordShrinkRequestStorageConfig `json:"StorageConfig,omitempty" xml:"StorageConfig,omitempty" type:"Struct"`
	SubHighResolutionStream *bool                                       `json:"SubHighResolutionStream,omitempty" xml:"SubHighResolutionStream,omitempty"`
	// taskId
	//
	// example:
	//
	// 123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// templateId
	//
	// This parameter is required.
	//
	// example:
	//
	// 567
	TemplateId *string                               `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Texts      []*StartCloudRecordShrinkRequestTexts `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
}

func (s StartCloudRecordShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordShrinkRequest) GoString() string {
	return s.String()
}

func (s *StartCloudRecordShrinkRequest) GetAnnotation() *string {
	return s.Annotation
}

func (s *StartCloudRecordShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *StartCloudRecordShrinkRequest) GetBackgrounds() []*StartCloudRecordShrinkRequestBackgrounds {
	return s.Backgrounds
}

func (s *StartCloudRecordShrinkRequest) GetBgColor() *StartCloudRecordShrinkRequestBgColor {
	return s.BgColor
}

func (s *StartCloudRecordShrinkRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *StartCloudRecordShrinkRequest) GetClockWidgets() []*StartCloudRecordShrinkRequestClockWidgets {
	return s.ClockWidgets
}

func (s *StartCloudRecordShrinkRequest) GetCropMode() *int32 {
	return s.CropMode
}

func (s *StartCloudRecordShrinkRequest) GetImages() []*StartCloudRecordShrinkRequestImages {
	return s.Images
}

func (s *StartCloudRecordShrinkRequest) GetLayoutSpecifiedUsersShrink() *string {
	return s.LayoutSpecifiedUsersShrink
}

func (s *StartCloudRecordShrinkRequest) GetPanes() []*StartCloudRecordShrinkRequestPanes {
	return s.Panes
}

func (s *StartCloudRecordShrinkRequest) GetRecordMode() *int32 {
	return s.RecordMode
}

func (s *StartCloudRecordShrinkRequest) GetRegionColor() *StartCloudRecordShrinkRequestRegionColor {
	return s.RegionColor
}

func (s *StartCloudRecordShrinkRequest) GetReservePaneForNoCameraUser() *bool {
	return s.ReservePaneForNoCameraUser
}

func (s *StartCloudRecordShrinkRequest) GetShowDefaultBackgroundOnMute() *bool {
	return s.ShowDefaultBackgroundOnMute
}

func (s *StartCloudRecordShrinkRequest) GetSingleStreamingRecordShrink() *string {
	return s.SingleStreamingRecordShrink
}

func (s *StartCloudRecordShrinkRequest) GetStartWithoutChannel() *bool {
	return s.StartWithoutChannel
}

func (s *StartCloudRecordShrinkRequest) GetStartWithoutChannelWaitTime() *int32 {
	return s.StartWithoutChannelWaitTime
}

func (s *StartCloudRecordShrinkRequest) GetStorageConfig() *StartCloudRecordShrinkRequestStorageConfig {
	return s.StorageConfig
}

func (s *StartCloudRecordShrinkRequest) GetSubHighResolutionStream() *bool {
	return s.SubHighResolutionStream
}

func (s *StartCloudRecordShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *StartCloudRecordShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *StartCloudRecordShrinkRequest) GetTexts() []*StartCloudRecordShrinkRequestTexts {
	return s.Texts
}

func (s *StartCloudRecordShrinkRequest) SetAnnotation(v string) *StartCloudRecordShrinkRequest {
	s.Annotation = &v
	return s
}

func (s *StartCloudRecordShrinkRequest) SetAppId(v string) *StartCloudRecordShrinkRequest {
	s.AppId = &v
	return s
}

func (s *StartCloudRecordShrinkRequest) SetBackgrounds(v []*StartCloudRecordShrinkRequestBackgrounds) *StartCloudRecordShrinkRequest {
	s.Backgrounds = v
	return s
}

func (s *StartCloudRecordShrinkRequest) SetBgColor(v *StartCloudRecordShrinkRequestBgColor) *StartCloudRecordShrinkRequest {
	s.BgColor = v
	return s
}

func (s *StartCloudRecordShrinkRequest) SetChannelId(v string) *StartCloudRecordShrinkRequest {
	s.ChannelId = &v
	return s
}

func (s *StartCloudRecordShrinkRequest) SetClockWidgets(v []*StartCloudRecordShrinkRequestClockWidgets) *StartCloudRecordShrinkRequest {
	s.ClockWidgets = v
	return s
}

func (s *StartCloudRecordShrinkRequest) SetCropMode(v int32) *StartCloudRecordShrinkRequest {
	s.CropMode = &v
	return s
}

func (s *StartCloudRecordShrinkRequest) SetImages(v []*StartCloudRecordShrinkRequestImages) *StartCloudRecordShrinkRequest {
	s.Images = v
	return s
}

func (s *StartCloudRecordShrinkRequest) SetLayoutSpecifiedUsersShrink(v string) *StartCloudRecordShrinkRequest {
	s.LayoutSpecifiedUsersShrink = &v
	return s
}

func (s *StartCloudRecordShrinkRequest) SetPanes(v []*StartCloudRecordShrinkRequestPanes) *StartCloudRecordShrinkRequest {
	s.Panes = v
	return s
}

func (s *StartCloudRecordShrinkRequest) SetRecordMode(v int32) *StartCloudRecordShrinkRequest {
	s.RecordMode = &v
	return s
}

func (s *StartCloudRecordShrinkRequest) SetRegionColor(v *StartCloudRecordShrinkRequestRegionColor) *StartCloudRecordShrinkRequest {
	s.RegionColor = v
	return s
}

func (s *StartCloudRecordShrinkRequest) SetReservePaneForNoCameraUser(v bool) *StartCloudRecordShrinkRequest {
	s.ReservePaneForNoCameraUser = &v
	return s
}

func (s *StartCloudRecordShrinkRequest) SetShowDefaultBackgroundOnMute(v bool) *StartCloudRecordShrinkRequest {
	s.ShowDefaultBackgroundOnMute = &v
	return s
}

func (s *StartCloudRecordShrinkRequest) SetSingleStreamingRecordShrink(v string) *StartCloudRecordShrinkRequest {
	s.SingleStreamingRecordShrink = &v
	return s
}

func (s *StartCloudRecordShrinkRequest) SetStartWithoutChannel(v bool) *StartCloudRecordShrinkRequest {
	s.StartWithoutChannel = &v
	return s
}

func (s *StartCloudRecordShrinkRequest) SetStartWithoutChannelWaitTime(v int32) *StartCloudRecordShrinkRequest {
	s.StartWithoutChannelWaitTime = &v
	return s
}

func (s *StartCloudRecordShrinkRequest) SetStorageConfig(v *StartCloudRecordShrinkRequestStorageConfig) *StartCloudRecordShrinkRequest {
	s.StorageConfig = v
	return s
}

func (s *StartCloudRecordShrinkRequest) SetSubHighResolutionStream(v bool) *StartCloudRecordShrinkRequest {
	s.SubHighResolutionStream = &v
	return s
}

func (s *StartCloudRecordShrinkRequest) SetTaskId(v string) *StartCloudRecordShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *StartCloudRecordShrinkRequest) SetTemplateId(v string) *StartCloudRecordShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *StartCloudRecordShrinkRequest) SetTexts(v []*StartCloudRecordShrinkRequestTexts) *StartCloudRecordShrinkRequest {
	s.Texts = v
	return s
}

func (s *StartCloudRecordShrinkRequest) Validate() error {
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
	if s.StorageConfig != nil {
		if err := s.StorageConfig.Validate(); err != nil {
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

type StartCloudRecordShrinkRequestBackgrounds struct {
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

func (s StartCloudRecordShrinkRequestBackgrounds) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordShrinkRequestBackgrounds) GoString() string {
	return s.String()
}

func (s *StartCloudRecordShrinkRequestBackgrounds) GetAlpha() *float64 {
	return s.Alpha
}

func (s *StartCloudRecordShrinkRequestBackgrounds) GetBackgroundCropMode() *int32 {
	return s.BackgroundCropMode
}

func (s *StartCloudRecordShrinkRequestBackgrounds) GetHeight() *float64 {
	return s.Height
}

func (s *StartCloudRecordShrinkRequestBackgrounds) GetLayer() *int32 {
	return s.Layer
}

func (s *StartCloudRecordShrinkRequestBackgrounds) GetUrl() *string {
	return s.Url
}

func (s *StartCloudRecordShrinkRequestBackgrounds) GetWidth() *float64 {
	return s.Width
}

func (s *StartCloudRecordShrinkRequestBackgrounds) GetX() *float64 {
	return s.X
}

func (s *StartCloudRecordShrinkRequestBackgrounds) GetY() *float64 {
	return s.Y
}

func (s *StartCloudRecordShrinkRequestBackgrounds) SetAlpha(v float64) *StartCloudRecordShrinkRequestBackgrounds {
	s.Alpha = &v
	return s
}

func (s *StartCloudRecordShrinkRequestBackgrounds) SetBackgroundCropMode(v int32) *StartCloudRecordShrinkRequestBackgrounds {
	s.BackgroundCropMode = &v
	return s
}

func (s *StartCloudRecordShrinkRequestBackgrounds) SetHeight(v float64) *StartCloudRecordShrinkRequestBackgrounds {
	s.Height = &v
	return s
}

func (s *StartCloudRecordShrinkRequestBackgrounds) SetLayer(v int32) *StartCloudRecordShrinkRequestBackgrounds {
	s.Layer = &v
	return s
}

func (s *StartCloudRecordShrinkRequestBackgrounds) SetUrl(v string) *StartCloudRecordShrinkRequestBackgrounds {
	s.Url = &v
	return s
}

func (s *StartCloudRecordShrinkRequestBackgrounds) SetWidth(v float64) *StartCloudRecordShrinkRequestBackgrounds {
	s.Width = &v
	return s
}

func (s *StartCloudRecordShrinkRequestBackgrounds) SetX(v float64) *StartCloudRecordShrinkRequestBackgrounds {
	s.X = &v
	return s
}

func (s *StartCloudRecordShrinkRequestBackgrounds) SetY(v float64) *StartCloudRecordShrinkRequestBackgrounds {
	s.Y = &v
	return s
}

func (s *StartCloudRecordShrinkRequestBackgrounds) Validate() error {
	return dara.Validate(s)
}

type StartCloudRecordShrinkRequestBgColor struct {
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

func (s StartCloudRecordShrinkRequestBgColor) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordShrinkRequestBgColor) GoString() string {
	return s.String()
}

func (s *StartCloudRecordShrinkRequestBgColor) GetB() *int32 {
	return s.B
}

func (s *StartCloudRecordShrinkRequestBgColor) GetG() *int32 {
	return s.G
}

func (s *StartCloudRecordShrinkRequestBgColor) GetR() *int32 {
	return s.R
}

func (s *StartCloudRecordShrinkRequestBgColor) SetB(v int32) *StartCloudRecordShrinkRequestBgColor {
	s.B = &v
	return s
}

func (s *StartCloudRecordShrinkRequestBgColor) SetG(v int32) *StartCloudRecordShrinkRequestBgColor {
	s.G = &v
	return s
}

func (s *StartCloudRecordShrinkRequestBgColor) SetR(v int32) *StartCloudRecordShrinkRequestBgColor {
	s.R = &v
	return s
}

func (s *StartCloudRecordShrinkRequestBgColor) Validate() error {
	return dara.Validate(s)
}

type StartCloudRecordShrinkRequestClockWidgets struct {
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
	BoxColor   *StartCloudRecordShrinkRequestClockWidgetsBoxColor `json:"BoxColor,omitempty" xml:"BoxColor,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Font      *int32                                              `json:"Font,omitempty" xml:"Font,omitempty"`
	FontColor *StartCloudRecordShrinkRequestClockWidgetsFontColor `json:"FontColor,omitempty" xml:"FontColor,omitempty" type:"Struct"`
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

func (s StartCloudRecordShrinkRequestClockWidgets) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordShrinkRequestClockWidgets) GoString() string {
	return s.String()
}

func (s *StartCloudRecordShrinkRequestClockWidgets) GetAlpha() *float64 {
	return s.Alpha
}

func (s *StartCloudRecordShrinkRequestClockWidgets) GetBoxAlpha() *float64 {
	return s.BoxAlpha
}

func (s *StartCloudRecordShrinkRequestClockWidgets) GetBoxBorderw() *int32 {
	return s.BoxBorderw
}

func (s *StartCloudRecordShrinkRequestClockWidgets) GetBoxColor() *StartCloudRecordShrinkRequestClockWidgetsBoxColor {
	return s.BoxColor
}

func (s *StartCloudRecordShrinkRequestClockWidgets) GetFont() *int32 {
	return s.Font
}

func (s *StartCloudRecordShrinkRequestClockWidgets) GetFontColor() *StartCloudRecordShrinkRequestClockWidgetsFontColor {
	return s.FontColor
}

func (s *StartCloudRecordShrinkRequestClockWidgets) GetFontSize() *int32 {
	return s.FontSize
}

func (s *StartCloudRecordShrinkRequestClockWidgets) GetHasBox() *bool {
	return s.HasBox
}

func (s *StartCloudRecordShrinkRequestClockWidgets) GetLayer() *int32 {
	return s.Layer
}

func (s *StartCloudRecordShrinkRequestClockWidgets) GetX() *float64 {
	return s.X
}

func (s *StartCloudRecordShrinkRequestClockWidgets) GetY() *float64 {
	return s.Y
}

func (s *StartCloudRecordShrinkRequestClockWidgets) GetZone() *int32 {
	return s.Zone
}

func (s *StartCloudRecordShrinkRequestClockWidgets) SetAlpha(v float64) *StartCloudRecordShrinkRequestClockWidgets {
	s.Alpha = &v
	return s
}

func (s *StartCloudRecordShrinkRequestClockWidgets) SetBoxAlpha(v float64) *StartCloudRecordShrinkRequestClockWidgets {
	s.BoxAlpha = &v
	return s
}

func (s *StartCloudRecordShrinkRequestClockWidgets) SetBoxBorderw(v int32) *StartCloudRecordShrinkRequestClockWidgets {
	s.BoxBorderw = &v
	return s
}

func (s *StartCloudRecordShrinkRequestClockWidgets) SetBoxColor(v *StartCloudRecordShrinkRequestClockWidgetsBoxColor) *StartCloudRecordShrinkRequestClockWidgets {
	s.BoxColor = v
	return s
}

func (s *StartCloudRecordShrinkRequestClockWidgets) SetFont(v int32) *StartCloudRecordShrinkRequestClockWidgets {
	s.Font = &v
	return s
}

func (s *StartCloudRecordShrinkRequestClockWidgets) SetFontColor(v *StartCloudRecordShrinkRequestClockWidgetsFontColor) *StartCloudRecordShrinkRequestClockWidgets {
	s.FontColor = v
	return s
}

func (s *StartCloudRecordShrinkRequestClockWidgets) SetFontSize(v int32) *StartCloudRecordShrinkRequestClockWidgets {
	s.FontSize = &v
	return s
}

func (s *StartCloudRecordShrinkRequestClockWidgets) SetHasBox(v bool) *StartCloudRecordShrinkRequestClockWidgets {
	s.HasBox = &v
	return s
}

func (s *StartCloudRecordShrinkRequestClockWidgets) SetLayer(v int32) *StartCloudRecordShrinkRequestClockWidgets {
	s.Layer = &v
	return s
}

func (s *StartCloudRecordShrinkRequestClockWidgets) SetX(v float64) *StartCloudRecordShrinkRequestClockWidgets {
	s.X = &v
	return s
}

func (s *StartCloudRecordShrinkRequestClockWidgets) SetY(v float64) *StartCloudRecordShrinkRequestClockWidgets {
	s.Y = &v
	return s
}

func (s *StartCloudRecordShrinkRequestClockWidgets) SetZone(v int32) *StartCloudRecordShrinkRequestClockWidgets {
	s.Zone = &v
	return s
}

func (s *StartCloudRecordShrinkRequestClockWidgets) Validate() error {
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

type StartCloudRecordShrinkRequestClockWidgetsBoxColor struct {
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

func (s StartCloudRecordShrinkRequestClockWidgetsBoxColor) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordShrinkRequestClockWidgetsBoxColor) GoString() string {
	return s.String()
}

func (s *StartCloudRecordShrinkRequestClockWidgetsBoxColor) GetB() *int32 {
	return s.B
}

func (s *StartCloudRecordShrinkRequestClockWidgetsBoxColor) GetG() *int32 {
	return s.G
}

func (s *StartCloudRecordShrinkRequestClockWidgetsBoxColor) GetR() *int32 {
	return s.R
}

func (s *StartCloudRecordShrinkRequestClockWidgetsBoxColor) SetB(v int32) *StartCloudRecordShrinkRequestClockWidgetsBoxColor {
	s.B = &v
	return s
}

func (s *StartCloudRecordShrinkRequestClockWidgetsBoxColor) SetG(v int32) *StartCloudRecordShrinkRequestClockWidgetsBoxColor {
	s.G = &v
	return s
}

func (s *StartCloudRecordShrinkRequestClockWidgetsBoxColor) SetR(v int32) *StartCloudRecordShrinkRequestClockWidgetsBoxColor {
	s.R = &v
	return s
}

func (s *StartCloudRecordShrinkRequestClockWidgetsBoxColor) Validate() error {
	return dara.Validate(s)
}

type StartCloudRecordShrinkRequestClockWidgetsFontColor struct {
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

func (s StartCloudRecordShrinkRequestClockWidgetsFontColor) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordShrinkRequestClockWidgetsFontColor) GoString() string {
	return s.String()
}

func (s *StartCloudRecordShrinkRequestClockWidgetsFontColor) GetB() *int32 {
	return s.B
}

func (s *StartCloudRecordShrinkRequestClockWidgetsFontColor) GetG() *int32 {
	return s.G
}

func (s *StartCloudRecordShrinkRequestClockWidgetsFontColor) GetR() *int32 {
	return s.R
}

func (s *StartCloudRecordShrinkRequestClockWidgetsFontColor) SetB(v int32) *StartCloudRecordShrinkRequestClockWidgetsFontColor {
	s.B = &v
	return s
}

func (s *StartCloudRecordShrinkRequestClockWidgetsFontColor) SetG(v int32) *StartCloudRecordShrinkRequestClockWidgetsFontColor {
	s.G = &v
	return s
}

func (s *StartCloudRecordShrinkRequestClockWidgetsFontColor) SetR(v int32) *StartCloudRecordShrinkRequestClockWidgetsFontColor {
	s.R = &v
	return s
}

func (s *StartCloudRecordShrinkRequestClockWidgetsFontColor) Validate() error {
	return dara.Validate(s)
}

type StartCloudRecordShrinkRequestImages struct {
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
	// 0.3
	Y *float64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s StartCloudRecordShrinkRequestImages) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordShrinkRequestImages) GoString() string {
	return s.String()
}

func (s *StartCloudRecordShrinkRequestImages) GetAlpha() *float64 {
	return s.Alpha
}

func (s *StartCloudRecordShrinkRequestImages) GetHeight() *float64 {
	return s.Height
}

func (s *StartCloudRecordShrinkRequestImages) GetImageCropMode() *int32 {
	return s.ImageCropMode
}

func (s *StartCloudRecordShrinkRequestImages) GetLayer() *int32 {
	return s.Layer
}

func (s *StartCloudRecordShrinkRequestImages) GetUrl() *string {
	return s.Url
}

func (s *StartCloudRecordShrinkRequestImages) GetWidth() *float64 {
	return s.Width
}

func (s *StartCloudRecordShrinkRequestImages) GetX() *float64 {
	return s.X
}

func (s *StartCloudRecordShrinkRequestImages) GetY() *float64 {
	return s.Y
}

func (s *StartCloudRecordShrinkRequestImages) SetAlpha(v float64) *StartCloudRecordShrinkRequestImages {
	s.Alpha = &v
	return s
}

func (s *StartCloudRecordShrinkRequestImages) SetHeight(v float64) *StartCloudRecordShrinkRequestImages {
	s.Height = &v
	return s
}

func (s *StartCloudRecordShrinkRequestImages) SetImageCropMode(v int32) *StartCloudRecordShrinkRequestImages {
	s.ImageCropMode = &v
	return s
}

func (s *StartCloudRecordShrinkRequestImages) SetLayer(v int32) *StartCloudRecordShrinkRequestImages {
	s.Layer = &v
	return s
}

func (s *StartCloudRecordShrinkRequestImages) SetUrl(v string) *StartCloudRecordShrinkRequestImages {
	s.Url = &v
	return s
}

func (s *StartCloudRecordShrinkRequestImages) SetWidth(v float64) *StartCloudRecordShrinkRequestImages {
	s.Width = &v
	return s
}

func (s *StartCloudRecordShrinkRequestImages) SetX(v float64) *StartCloudRecordShrinkRequestImages {
	s.X = &v
	return s
}

func (s *StartCloudRecordShrinkRequestImages) SetY(v float64) *StartCloudRecordShrinkRequestImages {
	s.Y = &v
	return s
}

func (s *StartCloudRecordShrinkRequestImages) Validate() error {
	return dara.Validate(s)
}

type StartCloudRecordShrinkRequestPanes struct {
	Backgrounds []*StartCloudRecordShrinkRequestPanesBackgrounds `json:"Backgrounds,omitempty" xml:"Backgrounds,omitempty" type:"Repeated"`
	Images      []*StartCloudRecordShrinkRequestPanesImages      `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// example:
	//
	// 3
	PaneCropMode *int32 `json:"PaneCropMode,omitempty" xml:"PaneCropMode,omitempty"`
	// paneId
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	PaneId                    *int32 `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
	ReservePaneForOfflineUser *bool  `json:"ReservePaneForOfflineUser,omitempty" xml:"ReservePaneForOfflineUser,omitempty"`
	// source
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// sourceType
	//
	// example:
	//
	// video
	SourceType *string                                    `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Texts      []*StartCloudRecordShrinkRequestPanesTexts `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
	// example:
	//
	// cameraFirst
	VideoOrder *string                                       `json:"VideoOrder,omitempty" xml:"VideoOrder,omitempty"`
	Whiteboard *StartCloudRecordShrinkRequestPanesWhiteboard `json:"Whiteboard,omitempty" xml:"Whiteboard,omitempty" type:"Struct"`
}

func (s StartCloudRecordShrinkRequestPanes) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordShrinkRequestPanes) GoString() string {
	return s.String()
}

func (s *StartCloudRecordShrinkRequestPanes) GetBackgrounds() []*StartCloudRecordShrinkRequestPanesBackgrounds {
	return s.Backgrounds
}

func (s *StartCloudRecordShrinkRequestPanes) GetImages() []*StartCloudRecordShrinkRequestPanesImages {
	return s.Images
}

func (s *StartCloudRecordShrinkRequestPanes) GetPaneCropMode() *int32 {
	return s.PaneCropMode
}

func (s *StartCloudRecordShrinkRequestPanes) GetPaneId() *int32 {
	return s.PaneId
}

func (s *StartCloudRecordShrinkRequestPanes) GetReservePaneForOfflineUser() *bool {
	return s.ReservePaneForOfflineUser
}

func (s *StartCloudRecordShrinkRequestPanes) GetSource() *string {
	return s.Source
}

func (s *StartCloudRecordShrinkRequestPanes) GetSourceType() *string {
	return s.SourceType
}

func (s *StartCloudRecordShrinkRequestPanes) GetTexts() []*StartCloudRecordShrinkRequestPanesTexts {
	return s.Texts
}

func (s *StartCloudRecordShrinkRequestPanes) GetVideoOrder() *string {
	return s.VideoOrder
}

func (s *StartCloudRecordShrinkRequestPanes) GetWhiteboard() *StartCloudRecordShrinkRequestPanesWhiteboard {
	return s.Whiteboard
}

func (s *StartCloudRecordShrinkRequestPanes) SetBackgrounds(v []*StartCloudRecordShrinkRequestPanesBackgrounds) *StartCloudRecordShrinkRequestPanes {
	s.Backgrounds = v
	return s
}

func (s *StartCloudRecordShrinkRequestPanes) SetImages(v []*StartCloudRecordShrinkRequestPanesImages) *StartCloudRecordShrinkRequestPanes {
	s.Images = v
	return s
}

func (s *StartCloudRecordShrinkRequestPanes) SetPaneCropMode(v int32) *StartCloudRecordShrinkRequestPanes {
	s.PaneCropMode = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanes) SetPaneId(v int32) *StartCloudRecordShrinkRequestPanes {
	s.PaneId = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanes) SetReservePaneForOfflineUser(v bool) *StartCloudRecordShrinkRequestPanes {
	s.ReservePaneForOfflineUser = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanes) SetSource(v string) *StartCloudRecordShrinkRequestPanes {
	s.Source = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanes) SetSourceType(v string) *StartCloudRecordShrinkRequestPanes {
	s.SourceType = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanes) SetTexts(v []*StartCloudRecordShrinkRequestPanesTexts) *StartCloudRecordShrinkRequestPanes {
	s.Texts = v
	return s
}

func (s *StartCloudRecordShrinkRequestPanes) SetVideoOrder(v string) *StartCloudRecordShrinkRequestPanes {
	s.VideoOrder = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanes) SetWhiteboard(v *StartCloudRecordShrinkRequestPanesWhiteboard) *StartCloudRecordShrinkRequestPanes {
	s.Whiteboard = v
	return s
}

func (s *StartCloudRecordShrinkRequestPanes) Validate() error {
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

type StartCloudRecordShrinkRequestPanesBackgrounds struct {
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

func (s StartCloudRecordShrinkRequestPanesBackgrounds) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordShrinkRequestPanesBackgrounds) GoString() string {
	return s.String()
}

func (s *StartCloudRecordShrinkRequestPanesBackgrounds) GetAlpha() *float64 {
	return s.Alpha
}

func (s *StartCloudRecordShrinkRequestPanesBackgrounds) GetDisplay() *string {
	return s.Display
}

func (s *StartCloudRecordShrinkRequestPanesBackgrounds) GetHeight() *float64 {
	return s.Height
}

func (s *StartCloudRecordShrinkRequestPanesBackgrounds) GetLayer() *int32 {
	return s.Layer
}

func (s *StartCloudRecordShrinkRequestPanesBackgrounds) GetPaneBackgroundCropMode() *int32 {
	return s.PaneBackgroundCropMode
}

func (s *StartCloudRecordShrinkRequestPanesBackgrounds) GetUrl() *string {
	return s.Url
}

func (s *StartCloudRecordShrinkRequestPanesBackgrounds) GetWidth() *float64 {
	return s.Width
}

func (s *StartCloudRecordShrinkRequestPanesBackgrounds) GetX() *float64 {
	return s.X
}

func (s *StartCloudRecordShrinkRequestPanesBackgrounds) GetY() *float64 {
	return s.Y
}

func (s *StartCloudRecordShrinkRequestPanesBackgrounds) SetAlpha(v float64) *StartCloudRecordShrinkRequestPanesBackgrounds {
	s.Alpha = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanesBackgrounds) SetDisplay(v string) *StartCloudRecordShrinkRequestPanesBackgrounds {
	s.Display = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanesBackgrounds) SetHeight(v float64) *StartCloudRecordShrinkRequestPanesBackgrounds {
	s.Height = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanesBackgrounds) SetLayer(v int32) *StartCloudRecordShrinkRequestPanesBackgrounds {
	s.Layer = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanesBackgrounds) SetPaneBackgroundCropMode(v int32) *StartCloudRecordShrinkRequestPanesBackgrounds {
	s.PaneBackgroundCropMode = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanesBackgrounds) SetUrl(v string) *StartCloudRecordShrinkRequestPanesBackgrounds {
	s.Url = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanesBackgrounds) SetWidth(v float64) *StartCloudRecordShrinkRequestPanesBackgrounds {
	s.Width = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanesBackgrounds) SetX(v float64) *StartCloudRecordShrinkRequestPanesBackgrounds {
	s.X = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanesBackgrounds) SetY(v float64) *StartCloudRecordShrinkRequestPanesBackgrounds {
	s.Y = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanesBackgrounds) Validate() error {
	return dara.Validate(s)
}

type StartCloudRecordShrinkRequestPanesImages struct {
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

func (s StartCloudRecordShrinkRequestPanesImages) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordShrinkRequestPanesImages) GoString() string {
	return s.String()
}

func (s *StartCloudRecordShrinkRequestPanesImages) GetAlpha() *float64 {
	return s.Alpha
}

func (s *StartCloudRecordShrinkRequestPanesImages) GetDisplay() *string {
	return s.Display
}

func (s *StartCloudRecordShrinkRequestPanesImages) GetHeight() *float64 {
	return s.Height
}

func (s *StartCloudRecordShrinkRequestPanesImages) GetLayer() *int32 {
	return s.Layer
}

func (s *StartCloudRecordShrinkRequestPanesImages) GetPaneImageCropMode() *int32 {
	return s.PaneImageCropMode
}

func (s *StartCloudRecordShrinkRequestPanesImages) GetUrl() *string {
	return s.Url
}

func (s *StartCloudRecordShrinkRequestPanesImages) GetWidth() *float64 {
	return s.Width
}

func (s *StartCloudRecordShrinkRequestPanesImages) GetX() *float64 {
	return s.X
}

func (s *StartCloudRecordShrinkRequestPanesImages) GetY() *float64 {
	return s.Y
}

func (s *StartCloudRecordShrinkRequestPanesImages) SetAlpha(v float64) *StartCloudRecordShrinkRequestPanesImages {
	s.Alpha = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanesImages) SetDisplay(v string) *StartCloudRecordShrinkRequestPanesImages {
	s.Display = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanesImages) SetHeight(v float64) *StartCloudRecordShrinkRequestPanesImages {
	s.Height = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanesImages) SetLayer(v int32) *StartCloudRecordShrinkRequestPanesImages {
	s.Layer = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanesImages) SetPaneImageCropMode(v int32) *StartCloudRecordShrinkRequestPanesImages {
	s.PaneImageCropMode = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanesImages) SetUrl(v string) *StartCloudRecordShrinkRequestPanesImages {
	s.Url = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanesImages) SetWidth(v float64) *StartCloudRecordShrinkRequestPanesImages {
	s.Width = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanesImages) SetX(v float64) *StartCloudRecordShrinkRequestPanesImages {
	s.X = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanesImages) SetY(v float64) *StartCloudRecordShrinkRequestPanesImages {
	s.Y = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanesImages) Validate() error {
	return dara.Validate(s)
}

type StartCloudRecordShrinkRequestPanesTexts struct {
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
	BoxColor   *StartCloudRecordShrinkRequestPanesTextsBoxColor `json:"BoxColor,omitempty" xml:"BoxColor,omitempty" type:"Struct"`
	// example:
	//
	// backup
	Display *string `json:"Display,omitempty" xml:"Display,omitempty"`
	// example:
	//
	// 0
	Font      *int32                                            `json:"Font,omitempty" xml:"Font,omitempty"`
	FontColor *StartCloudRecordShrinkRequestPanesTextsFontColor `json:"FontColor,omitempty" xml:"FontColor,omitempty" type:"Struct"`
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

func (s StartCloudRecordShrinkRequestPanesTexts) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordShrinkRequestPanesTexts) GoString() string {
	return s.String()
}

func (s *StartCloudRecordShrinkRequestPanesTexts) GetAlpha() *float64 {
	return s.Alpha
}

func (s *StartCloudRecordShrinkRequestPanesTexts) GetBoxAlpha() *float64 {
	return s.BoxAlpha
}

func (s *StartCloudRecordShrinkRequestPanesTexts) GetBoxBorderw() *int32 {
	return s.BoxBorderw
}

func (s *StartCloudRecordShrinkRequestPanesTexts) GetBoxColor() *StartCloudRecordShrinkRequestPanesTextsBoxColor {
	return s.BoxColor
}

func (s *StartCloudRecordShrinkRequestPanesTexts) GetDisplay() *string {
	return s.Display
}

func (s *StartCloudRecordShrinkRequestPanesTexts) GetFont() *int32 {
	return s.Font
}

func (s *StartCloudRecordShrinkRequestPanesTexts) GetFontColor() *StartCloudRecordShrinkRequestPanesTextsFontColor {
	return s.FontColor
}

func (s *StartCloudRecordShrinkRequestPanesTexts) GetFontSize() *int32 {
	return s.FontSize
}

func (s *StartCloudRecordShrinkRequestPanesTexts) GetHasBox() *bool {
	return s.HasBox
}

func (s *StartCloudRecordShrinkRequestPanesTexts) GetLayer() *int32 {
	return s.Layer
}

func (s *StartCloudRecordShrinkRequestPanesTexts) GetTexture() *string {
	return s.Texture
}

func (s *StartCloudRecordShrinkRequestPanesTexts) GetX() *float64 {
	return s.X
}

func (s *StartCloudRecordShrinkRequestPanesTexts) GetY() *float64 {
	return s.Y
}

func (s *StartCloudRecordShrinkRequestPanesTexts) SetAlpha(v float64) *StartCloudRecordShrinkRequestPanesTexts {
	s.Alpha = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanesTexts) SetBoxAlpha(v float64) *StartCloudRecordShrinkRequestPanesTexts {
	s.BoxAlpha = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanesTexts) SetBoxBorderw(v int32) *StartCloudRecordShrinkRequestPanesTexts {
	s.BoxBorderw = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanesTexts) SetBoxColor(v *StartCloudRecordShrinkRequestPanesTextsBoxColor) *StartCloudRecordShrinkRequestPanesTexts {
	s.BoxColor = v
	return s
}

func (s *StartCloudRecordShrinkRequestPanesTexts) SetDisplay(v string) *StartCloudRecordShrinkRequestPanesTexts {
	s.Display = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanesTexts) SetFont(v int32) *StartCloudRecordShrinkRequestPanesTexts {
	s.Font = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanesTexts) SetFontColor(v *StartCloudRecordShrinkRequestPanesTextsFontColor) *StartCloudRecordShrinkRequestPanesTexts {
	s.FontColor = v
	return s
}

func (s *StartCloudRecordShrinkRequestPanesTexts) SetFontSize(v int32) *StartCloudRecordShrinkRequestPanesTexts {
	s.FontSize = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanesTexts) SetHasBox(v bool) *StartCloudRecordShrinkRequestPanesTexts {
	s.HasBox = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanesTexts) SetLayer(v int32) *StartCloudRecordShrinkRequestPanesTexts {
	s.Layer = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanesTexts) SetTexture(v string) *StartCloudRecordShrinkRequestPanesTexts {
	s.Texture = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanesTexts) SetX(v float64) *StartCloudRecordShrinkRequestPanesTexts {
	s.X = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanesTexts) SetY(v float64) *StartCloudRecordShrinkRequestPanesTexts {
	s.Y = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanesTexts) Validate() error {
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

type StartCloudRecordShrinkRequestPanesTextsBoxColor struct {
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

func (s StartCloudRecordShrinkRequestPanesTextsBoxColor) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordShrinkRequestPanesTextsBoxColor) GoString() string {
	return s.String()
}

func (s *StartCloudRecordShrinkRequestPanesTextsBoxColor) GetB() *int32 {
	return s.B
}

func (s *StartCloudRecordShrinkRequestPanesTextsBoxColor) GetG() *int32 {
	return s.G
}

func (s *StartCloudRecordShrinkRequestPanesTextsBoxColor) GetR() *int32 {
	return s.R
}

func (s *StartCloudRecordShrinkRequestPanesTextsBoxColor) SetB(v int32) *StartCloudRecordShrinkRequestPanesTextsBoxColor {
	s.B = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanesTextsBoxColor) SetG(v int32) *StartCloudRecordShrinkRequestPanesTextsBoxColor {
	s.G = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanesTextsBoxColor) SetR(v int32) *StartCloudRecordShrinkRequestPanesTextsBoxColor {
	s.R = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanesTextsBoxColor) Validate() error {
	return dara.Validate(s)
}

type StartCloudRecordShrinkRequestPanesTextsFontColor struct {
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

func (s StartCloudRecordShrinkRequestPanesTextsFontColor) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordShrinkRequestPanesTextsFontColor) GoString() string {
	return s.String()
}

func (s *StartCloudRecordShrinkRequestPanesTextsFontColor) GetB() *int32 {
	return s.B
}

func (s *StartCloudRecordShrinkRequestPanesTextsFontColor) GetG() *int32 {
	return s.G
}

func (s *StartCloudRecordShrinkRequestPanesTextsFontColor) GetR() *int32 {
	return s.R
}

func (s *StartCloudRecordShrinkRequestPanesTextsFontColor) SetB(v int32) *StartCloudRecordShrinkRequestPanesTextsFontColor {
	s.B = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanesTextsFontColor) SetG(v int32) *StartCloudRecordShrinkRequestPanesTextsFontColor {
	s.G = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanesTextsFontColor) SetR(v int32) *StartCloudRecordShrinkRequestPanesTextsFontColor {
	s.R = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanesTextsFontColor) Validate() error {
	return dara.Validate(s)
}

type StartCloudRecordShrinkRequestPanesWhiteboard struct {
	// example:
	//
	// default
	WhiteboardId *string `json:"WhiteboardId,omitempty" xml:"WhiteboardId,omitempty"`
}

func (s StartCloudRecordShrinkRequestPanesWhiteboard) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordShrinkRequestPanesWhiteboard) GoString() string {
	return s.String()
}

func (s *StartCloudRecordShrinkRequestPanesWhiteboard) GetWhiteboardId() *string {
	return s.WhiteboardId
}

func (s *StartCloudRecordShrinkRequestPanesWhiteboard) SetWhiteboardId(v string) *StartCloudRecordShrinkRequestPanesWhiteboard {
	s.WhiteboardId = &v
	return s
}

func (s *StartCloudRecordShrinkRequestPanesWhiteboard) Validate() error {
	return dara.Validate(s)
}

type StartCloudRecordShrinkRequestRegionColor struct {
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

func (s StartCloudRecordShrinkRequestRegionColor) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordShrinkRequestRegionColor) GoString() string {
	return s.String()
}

func (s *StartCloudRecordShrinkRequestRegionColor) GetB() *int32 {
	return s.B
}

func (s *StartCloudRecordShrinkRequestRegionColor) GetG() *int32 {
	return s.G
}

func (s *StartCloudRecordShrinkRequestRegionColor) GetR() *int32 {
	return s.R
}

func (s *StartCloudRecordShrinkRequestRegionColor) SetB(v int32) *StartCloudRecordShrinkRequestRegionColor {
	s.B = &v
	return s
}

func (s *StartCloudRecordShrinkRequestRegionColor) SetG(v int32) *StartCloudRecordShrinkRequestRegionColor {
	s.G = &v
	return s
}

func (s *StartCloudRecordShrinkRequestRegionColor) SetR(v int32) *StartCloudRecordShrinkRequestRegionColor {
	s.R = &v
	return s
}

func (s *StartCloudRecordShrinkRequestRegionColor) Validate() error {
	return dara.Validate(s)
}

type StartCloudRecordShrinkRequestStorageConfig struct {
	// accessKey
	//
	// This parameter is required.
	//
	// example:
	//
	// LTAX***
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// bucket
	//
	// This parameter is required.
	//
	// example:
	//
	// test-bucket-for-recording
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// example:
	//
	// https://aliyuns.dalian.oss.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// region
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Region *int32 `json:"Region,omitempty" xml:"Region,omitempty"`
	// secretKey
	//
	// This parameter is required.
	//
	// example:
	//
	// APb6qWYEzKtYxE***
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// vendor
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Vendor *int32 `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s StartCloudRecordShrinkRequestStorageConfig) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordShrinkRequestStorageConfig) GoString() string {
	return s.String()
}

func (s *StartCloudRecordShrinkRequestStorageConfig) GetAccessKey() *string {
	return s.AccessKey
}

func (s *StartCloudRecordShrinkRequestStorageConfig) GetBucket() *string {
	return s.Bucket
}

func (s *StartCloudRecordShrinkRequestStorageConfig) GetEndpoint() *string {
	return s.Endpoint
}

func (s *StartCloudRecordShrinkRequestStorageConfig) GetRegion() *int32 {
	return s.Region
}

func (s *StartCloudRecordShrinkRequestStorageConfig) GetSecretKey() *string {
	return s.SecretKey
}

func (s *StartCloudRecordShrinkRequestStorageConfig) GetVendor() *int32 {
	return s.Vendor
}

func (s *StartCloudRecordShrinkRequestStorageConfig) SetAccessKey(v string) *StartCloudRecordShrinkRequestStorageConfig {
	s.AccessKey = &v
	return s
}

func (s *StartCloudRecordShrinkRequestStorageConfig) SetBucket(v string) *StartCloudRecordShrinkRequestStorageConfig {
	s.Bucket = &v
	return s
}

func (s *StartCloudRecordShrinkRequestStorageConfig) SetEndpoint(v string) *StartCloudRecordShrinkRequestStorageConfig {
	s.Endpoint = &v
	return s
}

func (s *StartCloudRecordShrinkRequestStorageConfig) SetRegion(v int32) *StartCloudRecordShrinkRequestStorageConfig {
	s.Region = &v
	return s
}

func (s *StartCloudRecordShrinkRequestStorageConfig) SetSecretKey(v string) *StartCloudRecordShrinkRequestStorageConfig {
	s.SecretKey = &v
	return s
}

func (s *StartCloudRecordShrinkRequestStorageConfig) SetVendor(v int32) *StartCloudRecordShrinkRequestStorageConfig {
	s.Vendor = &v
	return s
}

func (s *StartCloudRecordShrinkRequestStorageConfig) Validate() error {
	return dara.Validate(s)
}

type StartCloudRecordShrinkRequestTexts struct {
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
	BoxColor   *StartCloudRecordShrinkRequestTextsBoxColor `json:"BoxColor,omitempty" xml:"BoxColor,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Font      *int32                                       `json:"Font,omitempty" xml:"Font,omitempty"`
	FontColor *StartCloudRecordShrinkRequestTextsFontColor `json:"FontColor,omitempty" xml:"FontColor,omitempty" type:"Struct"`
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

func (s StartCloudRecordShrinkRequestTexts) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordShrinkRequestTexts) GoString() string {
	return s.String()
}

func (s *StartCloudRecordShrinkRequestTexts) GetAlpha() *float64 {
	return s.Alpha
}

func (s *StartCloudRecordShrinkRequestTexts) GetBoxAlpha() *float64 {
	return s.BoxAlpha
}

func (s *StartCloudRecordShrinkRequestTexts) GetBoxBorderw() *int32 {
	return s.BoxBorderw
}

func (s *StartCloudRecordShrinkRequestTexts) GetBoxColor() *StartCloudRecordShrinkRequestTextsBoxColor {
	return s.BoxColor
}

func (s *StartCloudRecordShrinkRequestTexts) GetFont() *int32 {
	return s.Font
}

func (s *StartCloudRecordShrinkRequestTexts) GetFontColor() *StartCloudRecordShrinkRequestTextsFontColor {
	return s.FontColor
}

func (s *StartCloudRecordShrinkRequestTexts) GetFontSize() *int32 {
	return s.FontSize
}

func (s *StartCloudRecordShrinkRequestTexts) GetHasBox() *bool {
	return s.HasBox
}

func (s *StartCloudRecordShrinkRequestTexts) GetLayer() *int32 {
	return s.Layer
}

func (s *StartCloudRecordShrinkRequestTexts) GetTexture() *string {
	return s.Texture
}

func (s *StartCloudRecordShrinkRequestTexts) GetX() *float64 {
	return s.X
}

func (s *StartCloudRecordShrinkRequestTexts) GetY() *float64 {
	return s.Y
}

func (s *StartCloudRecordShrinkRequestTexts) SetAlpha(v float64) *StartCloudRecordShrinkRequestTexts {
	s.Alpha = &v
	return s
}

func (s *StartCloudRecordShrinkRequestTexts) SetBoxAlpha(v float64) *StartCloudRecordShrinkRequestTexts {
	s.BoxAlpha = &v
	return s
}

func (s *StartCloudRecordShrinkRequestTexts) SetBoxBorderw(v int32) *StartCloudRecordShrinkRequestTexts {
	s.BoxBorderw = &v
	return s
}

func (s *StartCloudRecordShrinkRequestTexts) SetBoxColor(v *StartCloudRecordShrinkRequestTextsBoxColor) *StartCloudRecordShrinkRequestTexts {
	s.BoxColor = v
	return s
}

func (s *StartCloudRecordShrinkRequestTexts) SetFont(v int32) *StartCloudRecordShrinkRequestTexts {
	s.Font = &v
	return s
}

func (s *StartCloudRecordShrinkRequestTexts) SetFontColor(v *StartCloudRecordShrinkRequestTextsFontColor) *StartCloudRecordShrinkRequestTexts {
	s.FontColor = v
	return s
}

func (s *StartCloudRecordShrinkRequestTexts) SetFontSize(v int32) *StartCloudRecordShrinkRequestTexts {
	s.FontSize = &v
	return s
}

func (s *StartCloudRecordShrinkRequestTexts) SetHasBox(v bool) *StartCloudRecordShrinkRequestTexts {
	s.HasBox = &v
	return s
}

func (s *StartCloudRecordShrinkRequestTexts) SetLayer(v int32) *StartCloudRecordShrinkRequestTexts {
	s.Layer = &v
	return s
}

func (s *StartCloudRecordShrinkRequestTexts) SetTexture(v string) *StartCloudRecordShrinkRequestTexts {
	s.Texture = &v
	return s
}

func (s *StartCloudRecordShrinkRequestTexts) SetX(v float64) *StartCloudRecordShrinkRequestTexts {
	s.X = &v
	return s
}

func (s *StartCloudRecordShrinkRequestTexts) SetY(v float64) *StartCloudRecordShrinkRequestTexts {
	s.Y = &v
	return s
}

func (s *StartCloudRecordShrinkRequestTexts) Validate() error {
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

type StartCloudRecordShrinkRequestTextsBoxColor struct {
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

func (s StartCloudRecordShrinkRequestTextsBoxColor) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordShrinkRequestTextsBoxColor) GoString() string {
	return s.String()
}

func (s *StartCloudRecordShrinkRequestTextsBoxColor) GetB() *int32 {
	return s.B
}

func (s *StartCloudRecordShrinkRequestTextsBoxColor) GetG() *int32 {
	return s.G
}

func (s *StartCloudRecordShrinkRequestTextsBoxColor) GetR() *int32 {
	return s.R
}

func (s *StartCloudRecordShrinkRequestTextsBoxColor) SetB(v int32) *StartCloudRecordShrinkRequestTextsBoxColor {
	s.B = &v
	return s
}

func (s *StartCloudRecordShrinkRequestTextsBoxColor) SetG(v int32) *StartCloudRecordShrinkRequestTextsBoxColor {
	s.G = &v
	return s
}

func (s *StartCloudRecordShrinkRequestTextsBoxColor) SetR(v int32) *StartCloudRecordShrinkRequestTextsBoxColor {
	s.R = &v
	return s
}

func (s *StartCloudRecordShrinkRequestTextsBoxColor) Validate() error {
	return dara.Validate(s)
}

type StartCloudRecordShrinkRequestTextsFontColor struct {
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

func (s StartCloudRecordShrinkRequestTextsFontColor) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordShrinkRequestTextsFontColor) GoString() string {
	return s.String()
}

func (s *StartCloudRecordShrinkRequestTextsFontColor) GetB() *int32 {
	return s.B
}

func (s *StartCloudRecordShrinkRequestTextsFontColor) GetG() *int32 {
	return s.G
}

func (s *StartCloudRecordShrinkRequestTextsFontColor) GetR() *int32 {
	return s.R
}

func (s *StartCloudRecordShrinkRequestTextsFontColor) SetB(v int32) *StartCloudRecordShrinkRequestTextsFontColor {
	s.B = &v
	return s
}

func (s *StartCloudRecordShrinkRequestTextsFontColor) SetG(v int32) *StartCloudRecordShrinkRequestTextsFontColor {
	s.G = &v
	return s
}

func (s *StartCloudRecordShrinkRequestTextsFontColor) SetR(v int32) *StartCloudRecordShrinkRequestTextsFontColor {
	s.R = &v
	return s
}

func (s *StartCloudRecordShrinkRequestTextsFontColor) Validate() error {
	return dara.Validate(s)
}
