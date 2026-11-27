// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartStreamingOutRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnnotation(v string) *StartStreamingOutRequest
	GetAnnotation() *string
	SetAppId(v string) *StartStreamingOutRequest
	GetAppId() *string
	SetBackgrounds(v []*StartStreamingOutRequestBackgrounds) *StartStreamingOutRequest
	GetBackgrounds() []*StartStreamingOutRequestBackgrounds
	SetBgColor(v *StartStreamingOutRequestBgColor) *StartStreamingOutRequest
	GetBgColor() *StartStreamingOutRequestBgColor
	SetChannelId(v string) *StartStreamingOutRequest
	GetChannelId() *string
	SetClockWidgets(v []*StartStreamingOutRequestClockWidgets) *StartStreamingOutRequest
	GetClockWidgets() []*StartStreamingOutRequestClockWidgets
	SetCropMode(v int32) *StartStreamingOutRequest
	GetCropMode() *int32
	SetImages(v []*StartStreamingOutRequestImages) *StartStreamingOutRequest
	GetImages() []*StartStreamingOutRequestImages
	SetLayoutSpecifiedUsers(v *StartStreamingOutRequestLayoutSpecifiedUsers) *StartStreamingOutRequest
	GetLayoutSpecifiedUsers() *StartStreamingOutRequestLayoutSpecifiedUsers
	SetPanes(v []*StartStreamingOutRequestPanes) *StartStreamingOutRequest
	GetPanes() []*StartStreamingOutRequestPanes
	SetRegionColor(v *StartStreamingOutRequestRegionColor) *StartStreamingOutRequest
	GetRegionColor() *StartStreamingOutRequestRegionColor
	SetReservePaneForNoCameraUser(v bool) *StartStreamingOutRequest
	GetReservePaneForNoCameraUser() *bool
	SetShowDefaultBackgroundOnMute(v bool) *StartStreamingOutRequest
	GetShowDefaultBackgroundOnMute() *bool
	SetSpecMixedUserList(v []*string) *StartStreamingOutRequest
	GetSpecMixedUserList() []*string
	SetStartWithoutChannel(v bool) *StartStreamingOutRequest
	GetStartWithoutChannel() *bool
	SetStartWithoutChannelWaitTime(v int32) *StartStreamingOutRequest
	GetStartWithoutChannelWaitTime() *int32
	SetSubHighResolutionStream(v bool) *StartStreamingOutRequest
	GetSubHighResolutionStream() *bool
	SetTaskId(v string) *StartStreamingOutRequest
	GetTaskId() *string
	SetTemplateId(v string) *StartStreamingOutRequest
	GetTemplateId() *string
	SetTexts(v []*StartStreamingOutRequestTexts) *StartStreamingOutRequest
	GetTexts() []*StartStreamingOutRequestTexts
	SetUrl(v string) *StartStreamingOutRequest
	GetUrl() *string
}

type StartStreamingOutRequest struct {
	// The annotation switch.
	//
	// example:
	//
	// disable
	Annotation *string `json:"Annotation,omitempty" xml:"Annotation,omitempty"`
	// The application ID. You can create and query the ID in the console. Only a single ID is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The background image.
	Backgrounds []*StartStreamingOutRequestBackgrounds `json:"Backgrounds,omitempty" xml:"Backgrounds,omitempty" type:"Repeated"`
	// The global background color.
	BgColor *StartStreamingOutRequestBgColor `json:"BgColor,omitempty" xml:"BgColor,omitempty" type:"Struct"`
	// The ID of an existing channel. Only a single ID is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The global clock watermark.
	ClockWidgets []*StartStreamingOutRequestClockWidgets `json:"ClockWidgets,omitempty" xml:"ClockWidgets,omitempty" type:"Repeated"`
	// The crop mode. Valid values:
	//
	// - **1**: Crop the image while maintaining the aspect ratio.
	//
	// - **2**: Maintain the aspect ratio and add borders (default).
	//
	// - **3**: Stretch the image to fill the pane.
	//
	// example:
	//
	// 2
	CropMode *int32 `json:"CropMode,omitempty" xml:"CropMode,omitempty"`
	// The global image watermark.
	//
	// example:
	//
	// 30
	Images []*StartStreamingOutRequestImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// The specified user list for layout generation, including Blacklist/Whitelist.
	//
	// Whitelist: specifies the users whose video streams are subscribed to in the channel.
	//
	// 	- After the settings are applied, automatic padding is performed in sequence based on this list.
	//
	// 	- When switching between multiple layouts in a template, only users in the whitelist are considered.
	//
	// Blacklist: specifies the users whose video streams are not subscribed to in the channel.
	//
	// 	- When switching between multiple layouts in a template, only users not in the blacklist are considered.
	LayoutSpecifiedUsers *StartStreamingOutRequestLayoutSpecifiedUsers `json:"LayoutSpecifiedUsers,omitempty" xml:"LayoutSpecifiedUsers,omitempty" type:"Struct"`
	// The pane settings.
	//
	// example:
	//
	// eo85****
	Panes []*StartStreamingOutRequestPanes `json:"Panes,omitempty" xml:"Panes,omitempty" type:"Repeated"`
	// The global pane background color.
	RegionColor *StartStreamingOutRequestRegionColor `json:"RegionColor,omitempty" xml:"RegionColor,omitempty" type:"Struct"`
	// Specifies whether to reserve a pane placeholder for users who do not push camera streams. Default value: true. A value of true indicates that a pane is reserved when a user does not push a camera stream. A value of false indicates that no pane is reserved.
	ReservePaneForNoCameraUser *bool `json:"ReservePaneForNoCameraUser,omitempty" xml:"ReservePaneForNoCameraUser,omitempty"`
	// Specifies whether to display the default background image of a pane when the user is muted. Valid values:
	//
	// - true: Displays the default background image (default).
	//
	// - false: Does not display the default background image, but the configured local pane background image is still displayed.
	ShowDefaultBackgroundOnMute *bool `json:"ShowDefaultBackgroundOnMute,omitempty" xml:"ShowDefaultBackgroundOnMute,omitempty"`
	// The list of specified users for audio mixing.
	SpecMixedUserList []*string `json:"SpecMixedUserList,omitempty" xml:"SpecMixedUserList,omitempty" type:"Repeated"`
	// Specifies whether to allow starting the task when the channel does not exist. Valid values:
	//
	// - true: Yes.
	//
	// - false: No (default).
	StartWithoutChannel *bool `json:"StartWithoutChannel,omitempty" xml:"StartWithoutChannel,omitempty"`
	// The wait time for the task when the channel is not started, in seconds. This parameter is valid only when the StartWithoutChannel parameter is set to true. Valid values: 30 to 120.
	//
	// example:
	//
	// 30
	StartWithoutChannelWaitTime *int32 `json:"StartWithoutChannelWaitTime,omitempty" xml:"StartWithoutChannelWaitTime,omitempty"`
	// Specifies whether to subscribe to high resolution streams. Valid values:
	//
	// - false: Subscribes to streams of different resolutions based on the pane area (default).
	//
	// - true: Uses the highest resolution stream pushed by the client for padding by default.
	SubHighResolutionStream *bool `json:"SubHighResolutionStream,omitempty" xml:"SubHighResolutionStream,omitempty"`
	// The task ID of the bypass streaming task. Only a single ID is supported. The ID can contain uppercase and lowercase letters, digits, underscores (_), and hyphens (-), and can be up to 55 bytes in length. If you do not specify this parameter, the system automatically generates a task ID.
	//
	// example:
	//
	// 123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 567
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The global text watermark.
	Texts []*StartStreamingOutRequestTexts `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
	// The live stream ingest URL. Only a single URL is supported. This URL cannot be used by multiple tasks at the same time.
	//
	// This parameter is required.
	//
	// example:
	//
	// rtmp://example.com/live/stream
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s StartStreamingOutRequest) String() string {
	return dara.Prettify(s)
}

func (s StartStreamingOutRequest) GoString() string {
	return s.String()
}

func (s *StartStreamingOutRequest) GetAnnotation() *string {
	return s.Annotation
}

func (s *StartStreamingOutRequest) GetAppId() *string {
	return s.AppId
}

func (s *StartStreamingOutRequest) GetBackgrounds() []*StartStreamingOutRequestBackgrounds {
	return s.Backgrounds
}

func (s *StartStreamingOutRequest) GetBgColor() *StartStreamingOutRequestBgColor {
	return s.BgColor
}

func (s *StartStreamingOutRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *StartStreamingOutRequest) GetClockWidgets() []*StartStreamingOutRequestClockWidgets {
	return s.ClockWidgets
}

func (s *StartStreamingOutRequest) GetCropMode() *int32 {
	return s.CropMode
}

func (s *StartStreamingOutRequest) GetImages() []*StartStreamingOutRequestImages {
	return s.Images
}

func (s *StartStreamingOutRequest) GetLayoutSpecifiedUsers() *StartStreamingOutRequestLayoutSpecifiedUsers {
	return s.LayoutSpecifiedUsers
}

func (s *StartStreamingOutRequest) GetPanes() []*StartStreamingOutRequestPanes {
	return s.Panes
}

func (s *StartStreamingOutRequest) GetRegionColor() *StartStreamingOutRequestRegionColor {
	return s.RegionColor
}

func (s *StartStreamingOutRequest) GetReservePaneForNoCameraUser() *bool {
	return s.ReservePaneForNoCameraUser
}

func (s *StartStreamingOutRequest) GetShowDefaultBackgroundOnMute() *bool {
	return s.ShowDefaultBackgroundOnMute
}

func (s *StartStreamingOutRequest) GetSpecMixedUserList() []*string {
	return s.SpecMixedUserList
}

func (s *StartStreamingOutRequest) GetStartWithoutChannel() *bool {
	return s.StartWithoutChannel
}

func (s *StartStreamingOutRequest) GetStartWithoutChannelWaitTime() *int32 {
	return s.StartWithoutChannelWaitTime
}

func (s *StartStreamingOutRequest) GetSubHighResolutionStream() *bool {
	return s.SubHighResolutionStream
}

func (s *StartStreamingOutRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *StartStreamingOutRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *StartStreamingOutRequest) GetTexts() []*StartStreamingOutRequestTexts {
	return s.Texts
}

func (s *StartStreamingOutRequest) GetUrl() *string {
	return s.Url
}

func (s *StartStreamingOutRequest) SetAnnotation(v string) *StartStreamingOutRequest {
	s.Annotation = &v
	return s
}

func (s *StartStreamingOutRequest) SetAppId(v string) *StartStreamingOutRequest {
	s.AppId = &v
	return s
}

func (s *StartStreamingOutRequest) SetBackgrounds(v []*StartStreamingOutRequestBackgrounds) *StartStreamingOutRequest {
	s.Backgrounds = v
	return s
}

func (s *StartStreamingOutRequest) SetBgColor(v *StartStreamingOutRequestBgColor) *StartStreamingOutRequest {
	s.BgColor = v
	return s
}

func (s *StartStreamingOutRequest) SetChannelId(v string) *StartStreamingOutRequest {
	s.ChannelId = &v
	return s
}

func (s *StartStreamingOutRequest) SetClockWidgets(v []*StartStreamingOutRequestClockWidgets) *StartStreamingOutRequest {
	s.ClockWidgets = v
	return s
}

func (s *StartStreamingOutRequest) SetCropMode(v int32) *StartStreamingOutRequest {
	s.CropMode = &v
	return s
}

func (s *StartStreamingOutRequest) SetImages(v []*StartStreamingOutRequestImages) *StartStreamingOutRequest {
	s.Images = v
	return s
}

func (s *StartStreamingOutRequest) SetLayoutSpecifiedUsers(v *StartStreamingOutRequestLayoutSpecifiedUsers) *StartStreamingOutRequest {
	s.LayoutSpecifiedUsers = v
	return s
}

func (s *StartStreamingOutRequest) SetPanes(v []*StartStreamingOutRequestPanes) *StartStreamingOutRequest {
	s.Panes = v
	return s
}

func (s *StartStreamingOutRequest) SetRegionColor(v *StartStreamingOutRequestRegionColor) *StartStreamingOutRequest {
	s.RegionColor = v
	return s
}

func (s *StartStreamingOutRequest) SetReservePaneForNoCameraUser(v bool) *StartStreamingOutRequest {
	s.ReservePaneForNoCameraUser = &v
	return s
}

func (s *StartStreamingOutRequest) SetShowDefaultBackgroundOnMute(v bool) *StartStreamingOutRequest {
	s.ShowDefaultBackgroundOnMute = &v
	return s
}

func (s *StartStreamingOutRequest) SetSpecMixedUserList(v []*string) *StartStreamingOutRequest {
	s.SpecMixedUserList = v
	return s
}

func (s *StartStreamingOutRequest) SetStartWithoutChannel(v bool) *StartStreamingOutRequest {
	s.StartWithoutChannel = &v
	return s
}

func (s *StartStreamingOutRequest) SetStartWithoutChannelWaitTime(v int32) *StartStreamingOutRequest {
	s.StartWithoutChannelWaitTime = &v
	return s
}

func (s *StartStreamingOutRequest) SetSubHighResolutionStream(v bool) *StartStreamingOutRequest {
	s.SubHighResolutionStream = &v
	return s
}

func (s *StartStreamingOutRequest) SetTaskId(v string) *StartStreamingOutRequest {
	s.TaskId = &v
	return s
}

func (s *StartStreamingOutRequest) SetTemplateId(v string) *StartStreamingOutRequest {
	s.TemplateId = &v
	return s
}

func (s *StartStreamingOutRequest) SetTexts(v []*StartStreamingOutRequestTexts) *StartStreamingOutRequest {
	s.Texts = v
	return s
}

func (s *StartStreamingOutRequest) SetUrl(v string) *StartStreamingOutRequest {
	s.Url = &v
	return s
}

func (s *StartStreamingOutRequest) Validate() error {
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

type StartStreamingOutRequestBackgrounds struct {
	// The transparency.
	//
	// example:
	//
	// 0.9
	Alpha *float64 `json:"Alpha,omitempty" xml:"Alpha,omitempty"`
	// The crop mode. Valid values:
	//
	// - **1**: Crop while maintaining the aspect ratio.
	//
	// - **2**: Maintain the aspect ratio and add borders.
	//
	// - **3**: Stretch to fill.
	//
	// example:
	//
	// 2
	BackgroundCropMode *int32 `json:"BackgroundCropMode,omitempty" xml:"BackgroundCropMode,omitempty"`
	// The height.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	Height *float64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The z-axis layer.
	//
	// example:
	//
	// 0
	Layer *int32 `json:"Layer,omitempty" xml:"Layer,omitempty"`
	// The image URL.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://aliyun.com/123.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The width.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	Width *float64 `json:"Width,omitempty" xml:"Width,omitempty"`
	// The x-coordinate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	X *float64 `json:"X,omitempty" xml:"X,omitempty"`
	// The y-coordinate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	Y *float64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s StartStreamingOutRequestBackgrounds) String() string {
	return dara.Prettify(s)
}

func (s StartStreamingOutRequestBackgrounds) GoString() string {
	return s.String()
}

func (s *StartStreamingOutRequestBackgrounds) GetAlpha() *float64 {
	return s.Alpha
}

func (s *StartStreamingOutRequestBackgrounds) GetBackgroundCropMode() *int32 {
	return s.BackgroundCropMode
}

func (s *StartStreamingOutRequestBackgrounds) GetHeight() *float64 {
	return s.Height
}

func (s *StartStreamingOutRequestBackgrounds) GetLayer() *int32 {
	return s.Layer
}

func (s *StartStreamingOutRequestBackgrounds) GetUrl() *string {
	return s.Url
}

func (s *StartStreamingOutRequestBackgrounds) GetWidth() *float64 {
	return s.Width
}

func (s *StartStreamingOutRequestBackgrounds) GetX() *float64 {
	return s.X
}

func (s *StartStreamingOutRequestBackgrounds) GetY() *float64 {
	return s.Y
}

func (s *StartStreamingOutRequestBackgrounds) SetAlpha(v float64) *StartStreamingOutRequestBackgrounds {
	s.Alpha = &v
	return s
}

func (s *StartStreamingOutRequestBackgrounds) SetBackgroundCropMode(v int32) *StartStreamingOutRequestBackgrounds {
	s.BackgroundCropMode = &v
	return s
}

func (s *StartStreamingOutRequestBackgrounds) SetHeight(v float64) *StartStreamingOutRequestBackgrounds {
	s.Height = &v
	return s
}

func (s *StartStreamingOutRequestBackgrounds) SetLayer(v int32) *StartStreamingOutRequestBackgrounds {
	s.Layer = &v
	return s
}

func (s *StartStreamingOutRequestBackgrounds) SetUrl(v string) *StartStreamingOutRequestBackgrounds {
	s.Url = &v
	return s
}

func (s *StartStreamingOutRequestBackgrounds) SetWidth(v float64) *StartStreamingOutRequestBackgrounds {
	s.Width = &v
	return s
}

func (s *StartStreamingOutRequestBackgrounds) SetX(v float64) *StartStreamingOutRequestBackgrounds {
	s.X = &v
	return s
}

func (s *StartStreamingOutRequestBackgrounds) SetY(v float64) *StartStreamingOutRequestBackgrounds {
	s.Y = &v
	return s
}

func (s *StartStreamingOutRequestBackgrounds) Validate() error {
	return dara.Validate(s)
}

type StartStreamingOutRequestBgColor struct {
	// The B value.
	//
	// example:
	//
	// 255
	B *int32 `json:"B,omitempty" xml:"B,omitempty"`
	// The G value.
	//
	// example:
	//
	// 255
	G *int32 `json:"G,omitempty" xml:"G,omitempty"`
	// The R value.
	//
	// example:
	//
	// 255
	R *int32 `json:"R,omitempty" xml:"R,omitempty"`
}

func (s StartStreamingOutRequestBgColor) String() string {
	return dara.Prettify(s)
}

func (s StartStreamingOutRequestBgColor) GoString() string {
	return s.String()
}

func (s *StartStreamingOutRequestBgColor) GetB() *int32 {
	return s.B
}

func (s *StartStreamingOutRequestBgColor) GetG() *int32 {
	return s.G
}

func (s *StartStreamingOutRequestBgColor) GetR() *int32 {
	return s.R
}

func (s *StartStreamingOutRequestBgColor) SetB(v int32) *StartStreamingOutRequestBgColor {
	s.B = &v
	return s
}

func (s *StartStreamingOutRequestBgColor) SetG(v int32) *StartStreamingOutRequestBgColor {
	s.G = &v
	return s
}

func (s *StartStreamingOutRequestBgColor) SetR(v int32) *StartStreamingOutRequestBgColor {
	s.R = &v
	return s
}

func (s *StartStreamingOutRequestBgColor) Validate() error {
	return dara.Validate(s)
}

type StartStreamingOutRequestClockWidgets struct {
	// The transparency.
	//
	// example:
	//
	// 0.9
	Alpha *float64 `json:"Alpha,omitempty" xml:"Alpha,omitempty"`
	// The transparency of the rectangle box.
	//
	// example:
	//
	// 0.6
	BoxAlpha *float64 `json:"BoxAlpha,omitempty" xml:"BoxAlpha,omitempty"`
	// The border width of the rectangle box.
	//
	// example:
	//
	// 5
	BoxBorderw *int32 `json:"BoxBorderw,omitempty" xml:"BoxBorderw,omitempty"`
	// The color of the rectangle box.
	BoxColor *StartStreamingOutRequestClockWidgetsBoxColor `json:"BoxColor,omitempty" xml:"BoxColor,omitempty" type:"Struct"`
	// The font. Valid values:
	//
	// - **0**: Alibaba PuHuiTi Regular
	//
	// - **1**: Alibaba PuHuiTi Light
	//
	// - **2**: Alibaba PuHuiTi Medium
	//
	// - **3**: Alibaba PuHuiTi Bold
	//
	// - **4**: Alibaba PuHuiTi Heavy
	//
	// - **5**: Alimama ShuHeiTi
	//
	// - **6**: DingTalk JinBuTi
	//
	// - **7**: Alimama DaoLiTi
	//
	// - **8**: Alimama DongFangDaKai
	//
	// - **9**: Alimama FangYuanTi
	//
	// example:
	//
	// 0
	Font *int32 `json:"Font,omitempty" xml:"Font,omitempty"`
	// The font color.
	FontColor *StartStreamingOutRequestClockWidgetsFontColor `json:"FontColor,omitempty" xml:"FontColor,omitempty" type:"Struct"`
	// The font size of the clock.
	//
	// example:
	//
	// 30
	FontSize *int32 `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	// Specifies whether a rectangle box is required.
	HasBox *bool `json:"HasBox,omitempty" xml:"HasBox,omitempty"`
	// The z-axis layer.
	//
	// example:
	//
	// 0
	Layer *int32 `json:"Layer,omitempty" xml:"Layer,omitempty"`
	// The x-coordinate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	X *float64 `json:"X,omitempty" xml:"X,omitempty"`
	// The y-coordinate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	Y *float64 `json:"Y,omitempty" xml:"Y,omitempty"`
	// The time zone.
	//
	// example:
	//
	// 8
	Zone *int32 `json:"Zone,omitempty" xml:"Zone,omitempty"`
}

func (s StartStreamingOutRequestClockWidgets) String() string {
	return dara.Prettify(s)
}

func (s StartStreamingOutRequestClockWidgets) GoString() string {
	return s.String()
}

func (s *StartStreamingOutRequestClockWidgets) GetAlpha() *float64 {
	return s.Alpha
}

func (s *StartStreamingOutRequestClockWidgets) GetBoxAlpha() *float64 {
	return s.BoxAlpha
}

func (s *StartStreamingOutRequestClockWidgets) GetBoxBorderw() *int32 {
	return s.BoxBorderw
}

func (s *StartStreamingOutRequestClockWidgets) GetBoxColor() *StartStreamingOutRequestClockWidgetsBoxColor {
	return s.BoxColor
}

func (s *StartStreamingOutRequestClockWidgets) GetFont() *int32 {
	return s.Font
}

func (s *StartStreamingOutRequestClockWidgets) GetFontColor() *StartStreamingOutRequestClockWidgetsFontColor {
	return s.FontColor
}

func (s *StartStreamingOutRequestClockWidgets) GetFontSize() *int32 {
	return s.FontSize
}

func (s *StartStreamingOutRequestClockWidgets) GetHasBox() *bool {
	return s.HasBox
}

func (s *StartStreamingOutRequestClockWidgets) GetLayer() *int32 {
	return s.Layer
}

func (s *StartStreamingOutRequestClockWidgets) GetX() *float64 {
	return s.X
}

func (s *StartStreamingOutRequestClockWidgets) GetY() *float64 {
	return s.Y
}

func (s *StartStreamingOutRequestClockWidgets) GetZone() *int32 {
	return s.Zone
}

func (s *StartStreamingOutRequestClockWidgets) SetAlpha(v float64) *StartStreamingOutRequestClockWidgets {
	s.Alpha = &v
	return s
}

func (s *StartStreamingOutRequestClockWidgets) SetBoxAlpha(v float64) *StartStreamingOutRequestClockWidgets {
	s.BoxAlpha = &v
	return s
}

func (s *StartStreamingOutRequestClockWidgets) SetBoxBorderw(v int32) *StartStreamingOutRequestClockWidgets {
	s.BoxBorderw = &v
	return s
}

func (s *StartStreamingOutRequestClockWidgets) SetBoxColor(v *StartStreamingOutRequestClockWidgetsBoxColor) *StartStreamingOutRequestClockWidgets {
	s.BoxColor = v
	return s
}

func (s *StartStreamingOutRequestClockWidgets) SetFont(v int32) *StartStreamingOutRequestClockWidgets {
	s.Font = &v
	return s
}

func (s *StartStreamingOutRequestClockWidgets) SetFontColor(v *StartStreamingOutRequestClockWidgetsFontColor) *StartStreamingOutRequestClockWidgets {
	s.FontColor = v
	return s
}

func (s *StartStreamingOutRequestClockWidgets) SetFontSize(v int32) *StartStreamingOutRequestClockWidgets {
	s.FontSize = &v
	return s
}

func (s *StartStreamingOutRequestClockWidgets) SetHasBox(v bool) *StartStreamingOutRequestClockWidgets {
	s.HasBox = &v
	return s
}

func (s *StartStreamingOutRequestClockWidgets) SetLayer(v int32) *StartStreamingOutRequestClockWidgets {
	s.Layer = &v
	return s
}

func (s *StartStreamingOutRequestClockWidgets) SetX(v float64) *StartStreamingOutRequestClockWidgets {
	s.X = &v
	return s
}

func (s *StartStreamingOutRequestClockWidgets) SetY(v float64) *StartStreamingOutRequestClockWidgets {
	s.Y = &v
	return s
}

func (s *StartStreamingOutRequestClockWidgets) SetZone(v int32) *StartStreamingOutRequestClockWidgets {
	s.Zone = &v
	return s
}

func (s *StartStreamingOutRequestClockWidgets) Validate() error {
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

type StartStreamingOutRequestClockWidgetsBoxColor struct {
	// The B value.
	//
	// example:
	//
	// 255
	B *int32 `json:"B,omitempty" xml:"B,omitempty"`
	// The G value.
	//
	// example:
	//
	// 255
	G *int32 `json:"G,omitempty" xml:"G,omitempty"`
	// The R value.
	//
	// example:
	//
	// 255
	R *int32 `json:"R,omitempty" xml:"R,omitempty"`
}

func (s StartStreamingOutRequestClockWidgetsBoxColor) String() string {
	return dara.Prettify(s)
}

func (s StartStreamingOutRequestClockWidgetsBoxColor) GoString() string {
	return s.String()
}

func (s *StartStreamingOutRequestClockWidgetsBoxColor) GetB() *int32 {
	return s.B
}

func (s *StartStreamingOutRequestClockWidgetsBoxColor) GetG() *int32 {
	return s.G
}

func (s *StartStreamingOutRequestClockWidgetsBoxColor) GetR() *int32 {
	return s.R
}

func (s *StartStreamingOutRequestClockWidgetsBoxColor) SetB(v int32) *StartStreamingOutRequestClockWidgetsBoxColor {
	s.B = &v
	return s
}

func (s *StartStreamingOutRequestClockWidgetsBoxColor) SetG(v int32) *StartStreamingOutRequestClockWidgetsBoxColor {
	s.G = &v
	return s
}

func (s *StartStreamingOutRequestClockWidgetsBoxColor) SetR(v int32) *StartStreamingOutRequestClockWidgetsBoxColor {
	s.R = &v
	return s
}

func (s *StartStreamingOutRequestClockWidgetsBoxColor) Validate() error {
	return dara.Validate(s)
}

type StartStreamingOutRequestClockWidgetsFontColor struct {
	// The B value.
	//
	// example:
	//
	// 255
	B *int32 `json:"B,omitempty" xml:"B,omitempty"`
	// The G value.
	//
	// example:
	//
	// 255
	G *int32 `json:"G,omitempty" xml:"G,omitempty"`
	// The R value.
	//
	// example:
	//
	// 255
	R *int32 `json:"R,omitempty" xml:"R,omitempty"`
}

func (s StartStreamingOutRequestClockWidgetsFontColor) String() string {
	return dara.Prettify(s)
}

func (s StartStreamingOutRequestClockWidgetsFontColor) GoString() string {
	return s.String()
}

func (s *StartStreamingOutRequestClockWidgetsFontColor) GetB() *int32 {
	return s.B
}

func (s *StartStreamingOutRequestClockWidgetsFontColor) GetG() *int32 {
	return s.G
}

func (s *StartStreamingOutRequestClockWidgetsFontColor) GetR() *int32 {
	return s.R
}

func (s *StartStreamingOutRequestClockWidgetsFontColor) SetB(v int32) *StartStreamingOutRequestClockWidgetsFontColor {
	s.B = &v
	return s
}

func (s *StartStreamingOutRequestClockWidgetsFontColor) SetG(v int32) *StartStreamingOutRequestClockWidgetsFontColor {
	s.G = &v
	return s
}

func (s *StartStreamingOutRequestClockWidgetsFontColor) SetR(v int32) *StartStreamingOutRequestClockWidgetsFontColor {
	s.R = &v
	return s
}

func (s *StartStreamingOutRequestClockWidgetsFontColor) Validate() error {
	return dara.Validate(s)
}

type StartStreamingOutRequestImages struct {
	// The transparency.
	//
	// example:
	//
	// 0.9
	Alpha *float64 `json:"Alpha,omitempty" xml:"Alpha,omitempty"`
	// The height.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	Height *float64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The crop mode. Valid values:
	//
	// - **1**: Crop while maintaining the aspect ratio.
	//
	// - **2**: Maintain the aspect ratio and add borders.
	//
	// - **3**: Stretch to fill.
	//
	// example:
	//
	// 2
	ImageCropMode *int32 `json:"ImageCropMode,omitempty" xml:"ImageCropMode,omitempty"`
	// The z-axis layer.
	//
	// example:
	//
	// 0
	Layer *int32 `json:"Layer,omitempty" xml:"Layer,omitempty"`
	// The image URL. The image content cannot be updated during stream ingest.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://aliyun.com/123xxx.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The width.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	Width *float64 `json:"Width,omitempty" xml:"Width,omitempty"`
	// The x-coordinate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	X *float64 `json:"X,omitempty" xml:"X,omitempty"`
	// The y-coordinate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	Y *float64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s StartStreamingOutRequestImages) String() string {
	return dara.Prettify(s)
}

func (s StartStreamingOutRequestImages) GoString() string {
	return s.String()
}

func (s *StartStreamingOutRequestImages) GetAlpha() *float64 {
	return s.Alpha
}

func (s *StartStreamingOutRequestImages) GetHeight() *float64 {
	return s.Height
}

func (s *StartStreamingOutRequestImages) GetImageCropMode() *int32 {
	return s.ImageCropMode
}

func (s *StartStreamingOutRequestImages) GetLayer() *int32 {
	return s.Layer
}

func (s *StartStreamingOutRequestImages) GetUrl() *string {
	return s.Url
}

func (s *StartStreamingOutRequestImages) GetWidth() *float64 {
	return s.Width
}

func (s *StartStreamingOutRequestImages) GetX() *float64 {
	return s.X
}

func (s *StartStreamingOutRequestImages) GetY() *float64 {
	return s.Y
}

func (s *StartStreamingOutRequestImages) SetAlpha(v float64) *StartStreamingOutRequestImages {
	s.Alpha = &v
	return s
}

func (s *StartStreamingOutRequestImages) SetHeight(v float64) *StartStreamingOutRequestImages {
	s.Height = &v
	return s
}

func (s *StartStreamingOutRequestImages) SetImageCropMode(v int32) *StartStreamingOutRequestImages {
	s.ImageCropMode = &v
	return s
}

func (s *StartStreamingOutRequestImages) SetLayer(v int32) *StartStreamingOutRequestImages {
	s.Layer = &v
	return s
}

func (s *StartStreamingOutRequestImages) SetUrl(v string) *StartStreamingOutRequestImages {
	s.Url = &v
	return s
}

func (s *StartStreamingOutRequestImages) SetWidth(v float64) *StartStreamingOutRequestImages {
	s.Width = &v
	return s
}

func (s *StartStreamingOutRequestImages) SetX(v float64) *StartStreamingOutRequestImages {
	s.X = &v
	return s
}

func (s *StartStreamingOutRequestImages) SetY(v float64) *StartStreamingOutRequestImages {
	s.Y = &v
	return s
}

func (s *StartStreamingOutRequestImages) Validate() error {
	return dara.Validate(s)
}

type StartStreamingOutRequestLayoutSpecifiedUsers struct {
	// The list of specified user IDs.
	//
	// This parameter is required.
	Ids []*string `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	// The type of the user list. Valid values:
	//
	// - black: Blacklist.
	//
	// - white: Whitelist.
	//
	// This parameter is required.
	//
	// example:
	//
	// white
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s StartStreamingOutRequestLayoutSpecifiedUsers) String() string {
	return dara.Prettify(s)
}

func (s StartStreamingOutRequestLayoutSpecifiedUsers) GoString() string {
	return s.String()
}

func (s *StartStreamingOutRequestLayoutSpecifiedUsers) GetIds() []*string {
	return s.Ids
}

func (s *StartStreamingOutRequestLayoutSpecifiedUsers) GetType() *string {
	return s.Type
}

func (s *StartStreamingOutRequestLayoutSpecifiedUsers) SetIds(v []*string) *StartStreamingOutRequestLayoutSpecifiedUsers {
	s.Ids = v
	return s
}

func (s *StartStreamingOutRequestLayoutSpecifiedUsers) SetType(v string) *StartStreamingOutRequestLayoutSpecifiedUsers {
	s.Type = &v
	return s
}

func (s *StartStreamingOutRequestLayoutSpecifiedUsers) Validate() error {
	return dara.Validate(s)
}

type StartStreamingOutRequestPanes struct {
	// The local background images.
	Backgrounds []*StartStreamingOutRequestPanesBackgrounds `json:"Backgrounds,omitempty" xml:"Backgrounds,omitempty" type:"Repeated"`
	// The local image watermarks for the pane.
	Images []*StartStreamingOutRequestPanesImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// The crop mode. Valid values:
	//
	// - **1**: Crop while maintaining the aspect ratio.
	//
	// - **2**: Maintain the aspect ratio and add borders.
	//
	// - **3**: Stretch to fill.
	//
	// example:
	//
	// 2
	PaneCropMode *int32 `json:"PaneCropMode,omitempty" xml:"PaneCropMode,omitempty"`
	// The pane ID. The value starts from 0.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	PaneId *string `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
	// Specifies whether to reserve the pane for users who have not joined the meeting.
	//
	// Default value: true. This value indicates that when the specific user or wildcard user specified for the grid is offline, the pane is forcibly reserved.
	//
	// A value of false indicates that when the specific user or wildcard user specified for the grid is offline, video streams are displayed based on the automatic fill order.
	ReservePaneForOfflineUser *bool `json:"ReservePaneForOfflineUser,omitempty" xml:"ReservePaneForOfflineUser,omitempty"`
	// The media source of the pane. You can specify a userId. To specify a specific user type, you can use wildcard configurations. The following types are supported:
	//
	// - **$ActiveSpeaker**: The current speaker in the channel.
	//
	// - **$ScreenShare**: The user who is sharing their screen. If multiple users are sharing their screens in the channel, the panes are filled in chronological order.
	//
	// example:
	//
	// 22
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The media type of the pane. Valid values:
	//
	// - **video**: The user camera or screen sharing stream. This is the default value.
	//
	// - **camera**: The camera stream.
	//
	// - **screen**: The screen sharing stream.
	//
	// - **whiteboard**: The whiteboard.
	//
	// example:
	//
	// video
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The local text watermarks for the pane.
	Texts []*StartStreamingOutRequestPanesTexts `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
	// The fill order when SourceType is set to video. Valid values:
	//
	// - screenFirst: Screen sharing takes precedence. This is the default value.
	//
	// - cameraFirst: The camera takes precedence.
	//
	// example:
	//
	// cameraFirst
	VideoOrder *string `json:"VideoOrder,omitempty" xml:"VideoOrder,omitempty"`
	// The whiteboard configurations.
	Whiteboard *StartStreamingOutRequestPanesWhiteboard `json:"Whiteboard,omitempty" xml:"Whiteboard,omitempty" type:"Struct"`
}

func (s StartStreamingOutRequestPanes) String() string {
	return dara.Prettify(s)
}

func (s StartStreamingOutRequestPanes) GoString() string {
	return s.String()
}

func (s *StartStreamingOutRequestPanes) GetBackgrounds() []*StartStreamingOutRequestPanesBackgrounds {
	return s.Backgrounds
}

func (s *StartStreamingOutRequestPanes) GetImages() []*StartStreamingOutRequestPanesImages {
	return s.Images
}

func (s *StartStreamingOutRequestPanes) GetPaneCropMode() *int32 {
	return s.PaneCropMode
}

func (s *StartStreamingOutRequestPanes) GetPaneId() *string {
	return s.PaneId
}

func (s *StartStreamingOutRequestPanes) GetReservePaneForOfflineUser() *bool {
	return s.ReservePaneForOfflineUser
}

func (s *StartStreamingOutRequestPanes) GetSource() *string {
	return s.Source
}

func (s *StartStreamingOutRequestPanes) GetSourceType() *string {
	return s.SourceType
}

func (s *StartStreamingOutRequestPanes) GetTexts() []*StartStreamingOutRequestPanesTexts {
	return s.Texts
}

func (s *StartStreamingOutRequestPanes) GetVideoOrder() *string {
	return s.VideoOrder
}

func (s *StartStreamingOutRequestPanes) GetWhiteboard() *StartStreamingOutRequestPanesWhiteboard {
	return s.Whiteboard
}

func (s *StartStreamingOutRequestPanes) SetBackgrounds(v []*StartStreamingOutRequestPanesBackgrounds) *StartStreamingOutRequestPanes {
	s.Backgrounds = v
	return s
}

func (s *StartStreamingOutRequestPanes) SetImages(v []*StartStreamingOutRequestPanesImages) *StartStreamingOutRequestPanes {
	s.Images = v
	return s
}

func (s *StartStreamingOutRequestPanes) SetPaneCropMode(v int32) *StartStreamingOutRequestPanes {
	s.PaneCropMode = &v
	return s
}

func (s *StartStreamingOutRequestPanes) SetPaneId(v string) *StartStreamingOutRequestPanes {
	s.PaneId = &v
	return s
}

func (s *StartStreamingOutRequestPanes) SetReservePaneForOfflineUser(v bool) *StartStreamingOutRequestPanes {
	s.ReservePaneForOfflineUser = &v
	return s
}

func (s *StartStreamingOutRequestPanes) SetSource(v string) *StartStreamingOutRequestPanes {
	s.Source = &v
	return s
}

func (s *StartStreamingOutRequestPanes) SetSourceType(v string) *StartStreamingOutRequestPanes {
	s.SourceType = &v
	return s
}

func (s *StartStreamingOutRequestPanes) SetTexts(v []*StartStreamingOutRequestPanesTexts) *StartStreamingOutRequestPanes {
	s.Texts = v
	return s
}

func (s *StartStreamingOutRequestPanes) SetVideoOrder(v string) *StartStreamingOutRequestPanes {
	s.VideoOrder = &v
	return s
}

func (s *StartStreamingOutRequestPanes) SetWhiteboard(v *StartStreamingOutRequestPanesWhiteboard) *StartStreamingOutRequestPanes {
	s.Whiteboard = v
	return s
}

func (s *StartStreamingOutRequestPanes) Validate() error {
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

type StartStreamingOutRequestPanesBackgrounds struct {
	// The transparency.
	//
	// example:
	//
	// 0.9
	Alpha *float64 `json:"Alpha,omitempty" xml:"Alpha,omitempty"`
	// The display condition. Valid values:
	//
	// - streamOnly: Displayed when the current user has a video stream. This is the default value.
	//
	// - backup: Displayed when the current user does not have a video stream.
	//
	// - always: Always displayed.
	//
	// example:
	//
	// backup
	Display *string `json:"Display,omitempty" xml:"Display,omitempty"`
	// The height.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	Height *float64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The z-axis layer.
	//
	// example:
	//
	// 0
	Layer *int32 `json:"Layer,omitempty" xml:"Layer,omitempty"`
	// The crop mode. Valid values:
	//
	// - **1**: Crop while maintaining the aspect ratio.
	//
	// - **2**: Maintain the aspect ratio and add borders.
	//
	// - **3**: Stretch to fill.
	//
	// example:
	//
	// 2
	PaneBackgroundCropMode *int32 `json:"PaneBackgroundCropMode,omitempty" xml:"PaneBackgroundCropMode,omitempty"`
	// The image URL.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://aliyun.com/123xx.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The width.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.3
	Width *float64 `json:"Width,omitempty" xml:"Width,omitempty"`
	// The x-coordinate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	X *float64 `json:"X,omitempty" xml:"X,omitempty"`
	// The y-coordinate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	Y *float64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s StartStreamingOutRequestPanesBackgrounds) String() string {
	return dara.Prettify(s)
}

func (s StartStreamingOutRequestPanesBackgrounds) GoString() string {
	return s.String()
}

func (s *StartStreamingOutRequestPanesBackgrounds) GetAlpha() *float64 {
	return s.Alpha
}

func (s *StartStreamingOutRequestPanesBackgrounds) GetDisplay() *string {
	return s.Display
}

func (s *StartStreamingOutRequestPanesBackgrounds) GetHeight() *float64 {
	return s.Height
}

func (s *StartStreamingOutRequestPanesBackgrounds) GetLayer() *int32 {
	return s.Layer
}

func (s *StartStreamingOutRequestPanesBackgrounds) GetPaneBackgroundCropMode() *int32 {
	return s.PaneBackgroundCropMode
}

func (s *StartStreamingOutRequestPanesBackgrounds) GetUrl() *string {
	return s.Url
}

func (s *StartStreamingOutRequestPanesBackgrounds) GetWidth() *float64 {
	return s.Width
}

func (s *StartStreamingOutRequestPanesBackgrounds) GetX() *float64 {
	return s.X
}

func (s *StartStreamingOutRequestPanesBackgrounds) GetY() *float64 {
	return s.Y
}

func (s *StartStreamingOutRequestPanesBackgrounds) SetAlpha(v float64) *StartStreamingOutRequestPanesBackgrounds {
	s.Alpha = &v
	return s
}

func (s *StartStreamingOutRequestPanesBackgrounds) SetDisplay(v string) *StartStreamingOutRequestPanesBackgrounds {
	s.Display = &v
	return s
}

func (s *StartStreamingOutRequestPanesBackgrounds) SetHeight(v float64) *StartStreamingOutRequestPanesBackgrounds {
	s.Height = &v
	return s
}

func (s *StartStreamingOutRequestPanesBackgrounds) SetLayer(v int32) *StartStreamingOutRequestPanesBackgrounds {
	s.Layer = &v
	return s
}

func (s *StartStreamingOutRequestPanesBackgrounds) SetPaneBackgroundCropMode(v int32) *StartStreamingOutRequestPanesBackgrounds {
	s.PaneBackgroundCropMode = &v
	return s
}

func (s *StartStreamingOutRequestPanesBackgrounds) SetUrl(v string) *StartStreamingOutRequestPanesBackgrounds {
	s.Url = &v
	return s
}

func (s *StartStreamingOutRequestPanesBackgrounds) SetWidth(v float64) *StartStreamingOutRequestPanesBackgrounds {
	s.Width = &v
	return s
}

func (s *StartStreamingOutRequestPanesBackgrounds) SetX(v float64) *StartStreamingOutRequestPanesBackgrounds {
	s.X = &v
	return s
}

func (s *StartStreamingOutRequestPanesBackgrounds) SetY(v float64) *StartStreamingOutRequestPanesBackgrounds {
	s.Y = &v
	return s
}

func (s *StartStreamingOutRequestPanesBackgrounds) Validate() error {
	return dara.Validate(s)
}

type StartStreamingOutRequestPanesImages struct {
	// The transparency.
	//
	// example:
	//
	// 0.9
	Alpha *float64 `json:"Alpha,omitempty" xml:"Alpha,omitempty"`
	// The condition for displaying the watermark. Valid values:
	//
	// - streamOnly: Displayed when the current user has a video stream. This is the default value.
	//
	// - backup: Displayed when the current user does not have a video stream.
	//
	// - always: Always displayed.
	//
	// example:
	//
	// backup
	Display *string `json:"Display,omitempty" xml:"Display,omitempty"`
	// The height.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	Height *float64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The z-axis layer.
	//
	// example:
	//
	// 0
	Layer *int32 `json:"Layer,omitempty" xml:"Layer,omitempty"`
	// The crop mode. Valid values:
	//
	// - **1**: Crop while maintaining the aspect ratio.
	//
	// - **2**: Maintain the aspect ratio and add borders.
	//
	// - **3**: Stretch to fill.
	//
	// example:
	//
	// 2
	PaneImageCropMode *int32 `json:"PaneImageCropMode,omitempty" xml:"PaneImageCropMode,omitempty"`
	// The corresponding image content cannot be updated while streaming.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://aliyun.com/123xxx.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The width.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	Width *float64 `json:"Width,omitempty" xml:"Width,omitempty"`
	// The x-coordinate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	X *float64 `json:"X,omitempty" xml:"X,omitempty"`
	// The y-coordinate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.3
	Y *float64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s StartStreamingOutRequestPanesImages) String() string {
	return dara.Prettify(s)
}

func (s StartStreamingOutRequestPanesImages) GoString() string {
	return s.String()
}

func (s *StartStreamingOutRequestPanesImages) GetAlpha() *float64 {
	return s.Alpha
}

func (s *StartStreamingOutRequestPanesImages) GetDisplay() *string {
	return s.Display
}

func (s *StartStreamingOutRequestPanesImages) GetHeight() *float64 {
	return s.Height
}

func (s *StartStreamingOutRequestPanesImages) GetLayer() *int32 {
	return s.Layer
}

func (s *StartStreamingOutRequestPanesImages) GetPaneImageCropMode() *int32 {
	return s.PaneImageCropMode
}

func (s *StartStreamingOutRequestPanesImages) GetUrl() *string {
	return s.Url
}

func (s *StartStreamingOutRequestPanesImages) GetWidth() *float64 {
	return s.Width
}

func (s *StartStreamingOutRequestPanesImages) GetX() *float64 {
	return s.X
}

func (s *StartStreamingOutRequestPanesImages) GetY() *float64 {
	return s.Y
}

func (s *StartStreamingOutRequestPanesImages) SetAlpha(v float64) *StartStreamingOutRequestPanesImages {
	s.Alpha = &v
	return s
}

func (s *StartStreamingOutRequestPanesImages) SetDisplay(v string) *StartStreamingOutRequestPanesImages {
	s.Display = &v
	return s
}

func (s *StartStreamingOutRequestPanesImages) SetHeight(v float64) *StartStreamingOutRequestPanesImages {
	s.Height = &v
	return s
}

func (s *StartStreamingOutRequestPanesImages) SetLayer(v int32) *StartStreamingOutRequestPanesImages {
	s.Layer = &v
	return s
}

func (s *StartStreamingOutRequestPanesImages) SetPaneImageCropMode(v int32) *StartStreamingOutRequestPanesImages {
	s.PaneImageCropMode = &v
	return s
}

func (s *StartStreamingOutRequestPanesImages) SetUrl(v string) *StartStreamingOutRequestPanesImages {
	s.Url = &v
	return s
}

func (s *StartStreamingOutRequestPanesImages) SetWidth(v float64) *StartStreamingOutRequestPanesImages {
	s.Width = &v
	return s
}

func (s *StartStreamingOutRequestPanesImages) SetX(v float64) *StartStreamingOutRequestPanesImages {
	s.X = &v
	return s
}

func (s *StartStreamingOutRequestPanesImages) SetY(v float64) *StartStreamingOutRequestPanesImages {
	s.Y = &v
	return s
}

func (s *StartStreamingOutRequestPanesImages) Validate() error {
	return dara.Validate(s)
}

type StartStreamingOutRequestPanesTexts struct {
	// The watermark transparency.
	//
	// example:
	//
	// 0.9
	Alpha *float64 `json:"Alpha,omitempty" xml:"Alpha,omitempty"`
	// The transparency of the rectangle box.
	//
	// example:
	//
	// 0.6
	BoxAlpha *float64 `json:"BoxAlpha,omitempty" xml:"BoxAlpha,omitempty"`
	// The border width of the rectangle box.
	//
	// example:
	//
	// 5
	BoxBorderw *int32 `json:"BoxBorderw,omitempty" xml:"BoxBorderw,omitempty"`
	// The color of the rectangle box.
	BoxColor *StartStreamingOutRequestPanesTextsBoxColor `json:"BoxColor,omitempty" xml:"BoxColor,omitempty" type:"Struct"`
	// The condition for displaying the watermark. Valid values:
	//
	// - streamOnly: Displayed when the current user has a video stream. This is the default value.
	//
	// - backup: Displayed when the current user does not have a video stream.
	//
	// - always: Always displayed.
	//
	// example:
	//
	// backup
	Display *string `json:"Display,omitempty" xml:"Display,omitempty"`
	// The font. Valid values:
	//
	// - **0**: Alibaba PuHuiTi Regular
	//
	// - **1**: Alibaba PuHuiTi Light
	//
	// - **2**: Alibaba PuHuiTi Medium
	//
	// - **3**: Alibaba PuHuiTi Bold
	//
	// - **4**: Alibaba PuHuiTi Heavy
	//
	// - **5**: Alimama ShuHeiTi
	//
	// - **6**: DingTalk JinBuTi
	//
	// - **7**: Alimama DaoLiTi
	//
	// - **8**: Alimama DongFangDaKai
	//
	// - **9**: Alimama FangYuanTi
	//
	// example:
	//
	// 0
	Font *int32 `json:"Font,omitempty" xml:"Font,omitempty"`
	// The font color.
	FontColor *StartStreamingOutRequestPanesTextsFontColor `json:"FontColor,omitempty" xml:"FontColor,omitempty" type:"Struct"`
	// The font size.
	//
	// example:
	//
	// 36
	FontSize *int32 `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	// Specifies whether a rectangle box is required.
	HasBox *bool `json:"HasBox,omitempty" xml:"HasBox,omitempty"`
	// The z-axis layer.
	//
	// example:
	//
	// 0
	Layer *int32 `json:"Layer,omitempty" xml:"Layer,omitempty"`
	// The watermark content.
	//
	// This parameter is required.
	//
	// example:
	//
	// Text watermark
	Texture *string `json:"Texture,omitempty" xml:"Texture,omitempty"`
	// The x-coordinate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	X *float64 `json:"X,omitempty" xml:"X,omitempty"`
	// The y-coordinate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	Y *float64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s StartStreamingOutRequestPanesTexts) String() string {
	return dara.Prettify(s)
}

func (s StartStreamingOutRequestPanesTexts) GoString() string {
	return s.String()
}

func (s *StartStreamingOutRequestPanesTexts) GetAlpha() *float64 {
	return s.Alpha
}

func (s *StartStreamingOutRequestPanesTexts) GetBoxAlpha() *float64 {
	return s.BoxAlpha
}

func (s *StartStreamingOutRequestPanesTexts) GetBoxBorderw() *int32 {
	return s.BoxBorderw
}

func (s *StartStreamingOutRequestPanesTexts) GetBoxColor() *StartStreamingOutRequestPanesTextsBoxColor {
	return s.BoxColor
}

func (s *StartStreamingOutRequestPanesTexts) GetDisplay() *string {
	return s.Display
}

func (s *StartStreamingOutRequestPanesTexts) GetFont() *int32 {
	return s.Font
}

func (s *StartStreamingOutRequestPanesTexts) GetFontColor() *StartStreamingOutRequestPanesTextsFontColor {
	return s.FontColor
}

func (s *StartStreamingOutRequestPanesTexts) GetFontSize() *int32 {
	return s.FontSize
}

func (s *StartStreamingOutRequestPanesTexts) GetHasBox() *bool {
	return s.HasBox
}

func (s *StartStreamingOutRequestPanesTexts) GetLayer() *int32 {
	return s.Layer
}

func (s *StartStreamingOutRequestPanesTexts) GetTexture() *string {
	return s.Texture
}

func (s *StartStreamingOutRequestPanesTexts) GetX() *float64 {
	return s.X
}

func (s *StartStreamingOutRequestPanesTexts) GetY() *float64 {
	return s.Y
}

func (s *StartStreamingOutRequestPanesTexts) SetAlpha(v float64) *StartStreamingOutRequestPanesTexts {
	s.Alpha = &v
	return s
}

func (s *StartStreamingOutRequestPanesTexts) SetBoxAlpha(v float64) *StartStreamingOutRequestPanesTexts {
	s.BoxAlpha = &v
	return s
}

func (s *StartStreamingOutRequestPanesTexts) SetBoxBorderw(v int32) *StartStreamingOutRequestPanesTexts {
	s.BoxBorderw = &v
	return s
}

func (s *StartStreamingOutRequestPanesTexts) SetBoxColor(v *StartStreamingOutRequestPanesTextsBoxColor) *StartStreamingOutRequestPanesTexts {
	s.BoxColor = v
	return s
}

func (s *StartStreamingOutRequestPanesTexts) SetDisplay(v string) *StartStreamingOutRequestPanesTexts {
	s.Display = &v
	return s
}

func (s *StartStreamingOutRequestPanesTexts) SetFont(v int32) *StartStreamingOutRequestPanesTexts {
	s.Font = &v
	return s
}

func (s *StartStreamingOutRequestPanesTexts) SetFontColor(v *StartStreamingOutRequestPanesTextsFontColor) *StartStreamingOutRequestPanesTexts {
	s.FontColor = v
	return s
}

func (s *StartStreamingOutRequestPanesTexts) SetFontSize(v int32) *StartStreamingOutRequestPanesTexts {
	s.FontSize = &v
	return s
}

func (s *StartStreamingOutRequestPanesTexts) SetHasBox(v bool) *StartStreamingOutRequestPanesTexts {
	s.HasBox = &v
	return s
}

func (s *StartStreamingOutRequestPanesTexts) SetLayer(v int32) *StartStreamingOutRequestPanesTexts {
	s.Layer = &v
	return s
}

func (s *StartStreamingOutRequestPanesTexts) SetTexture(v string) *StartStreamingOutRequestPanesTexts {
	s.Texture = &v
	return s
}

func (s *StartStreamingOutRequestPanesTexts) SetX(v float64) *StartStreamingOutRequestPanesTexts {
	s.X = &v
	return s
}

func (s *StartStreamingOutRequestPanesTexts) SetY(v float64) *StartStreamingOutRequestPanesTexts {
	s.Y = &v
	return s
}

func (s *StartStreamingOutRequestPanesTexts) Validate() error {
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

type StartStreamingOutRequestPanesTextsBoxColor struct {
	// The B value.
	//
	// example:
	//
	// 255
	B *int32 `json:"B,omitempty" xml:"B,omitempty"`
	// The G value.
	//
	// example:
	//
	// 255
	G *int32 `json:"G,omitempty" xml:"G,omitempty"`
	// The R value.
	//
	// example:
	//
	// 255
	R *int32 `json:"R,omitempty" xml:"R,omitempty"`
}

func (s StartStreamingOutRequestPanesTextsBoxColor) String() string {
	return dara.Prettify(s)
}

func (s StartStreamingOutRequestPanesTextsBoxColor) GoString() string {
	return s.String()
}

func (s *StartStreamingOutRequestPanesTextsBoxColor) GetB() *int32 {
	return s.B
}

func (s *StartStreamingOutRequestPanesTextsBoxColor) GetG() *int32 {
	return s.G
}

func (s *StartStreamingOutRequestPanesTextsBoxColor) GetR() *int32 {
	return s.R
}

func (s *StartStreamingOutRequestPanesTextsBoxColor) SetB(v int32) *StartStreamingOutRequestPanesTextsBoxColor {
	s.B = &v
	return s
}

func (s *StartStreamingOutRequestPanesTextsBoxColor) SetG(v int32) *StartStreamingOutRequestPanesTextsBoxColor {
	s.G = &v
	return s
}

func (s *StartStreamingOutRequestPanesTextsBoxColor) SetR(v int32) *StartStreamingOutRequestPanesTextsBoxColor {
	s.R = &v
	return s
}

func (s *StartStreamingOutRequestPanesTextsBoxColor) Validate() error {
	return dara.Validate(s)
}

type StartStreamingOutRequestPanesTextsFontColor struct {
	// The B value.
	//
	// example:
	//
	// 255
	B *int32 `json:"B,omitempty" xml:"B,omitempty"`
	// The G value.
	//
	// example:
	//
	// 255
	G *int32 `json:"G,omitempty" xml:"G,omitempty"`
	// The R value.
	//
	// example:
	//
	// 255
	R *int32 `json:"R,omitempty" xml:"R,omitempty"`
}

func (s StartStreamingOutRequestPanesTextsFontColor) String() string {
	return dara.Prettify(s)
}

func (s StartStreamingOutRequestPanesTextsFontColor) GoString() string {
	return s.String()
}

func (s *StartStreamingOutRequestPanesTextsFontColor) GetB() *int32 {
	return s.B
}

func (s *StartStreamingOutRequestPanesTextsFontColor) GetG() *int32 {
	return s.G
}

func (s *StartStreamingOutRequestPanesTextsFontColor) GetR() *int32 {
	return s.R
}

func (s *StartStreamingOutRequestPanesTextsFontColor) SetB(v int32) *StartStreamingOutRequestPanesTextsFontColor {
	s.B = &v
	return s
}

func (s *StartStreamingOutRequestPanesTextsFontColor) SetG(v int32) *StartStreamingOutRequestPanesTextsFontColor {
	s.G = &v
	return s
}

func (s *StartStreamingOutRequestPanesTextsFontColor) SetR(v int32) *StartStreamingOutRequestPanesTextsFontColor {
	s.R = &v
	return s
}

func (s *StartStreamingOutRequestPanesTextsFontColor) Validate() error {
	return dara.Validate(s)
}

type StartStreamingOutRequestPanesWhiteboard struct {
	// The whiteboard ID. Default value: default.
	//
	// example:
	//
	// default
	WhiteboardId *string `json:"WhiteboardId,omitempty" xml:"WhiteboardId,omitempty"`
}

func (s StartStreamingOutRequestPanesWhiteboard) String() string {
	return dara.Prettify(s)
}

func (s StartStreamingOutRequestPanesWhiteboard) GoString() string {
	return s.String()
}

func (s *StartStreamingOutRequestPanesWhiteboard) GetWhiteboardId() *string {
	return s.WhiteboardId
}

func (s *StartStreamingOutRequestPanesWhiteboard) SetWhiteboardId(v string) *StartStreamingOutRequestPanesWhiteboard {
	s.WhiteboardId = &v
	return s
}

func (s *StartStreamingOutRequestPanesWhiteboard) Validate() error {
	return dara.Validate(s)
}

type StartStreamingOutRequestRegionColor struct {
	// The B value.
	//
	// example:
	//
	// 255
	B *int32 `json:"B,omitempty" xml:"B,omitempty"`
	// The G value.
	//
	// example:
	//
	// 255
	G *int32 `json:"G,omitempty" xml:"G,omitempty"`
	// The R value.
	//
	// example:
	//
	// 255
	R *int32 `json:"R,omitempty" xml:"R,omitempty"`
}

func (s StartStreamingOutRequestRegionColor) String() string {
	return dara.Prettify(s)
}

func (s StartStreamingOutRequestRegionColor) GoString() string {
	return s.String()
}

func (s *StartStreamingOutRequestRegionColor) GetB() *int32 {
	return s.B
}

func (s *StartStreamingOutRequestRegionColor) GetG() *int32 {
	return s.G
}

func (s *StartStreamingOutRequestRegionColor) GetR() *int32 {
	return s.R
}

func (s *StartStreamingOutRequestRegionColor) SetB(v int32) *StartStreamingOutRequestRegionColor {
	s.B = &v
	return s
}

func (s *StartStreamingOutRequestRegionColor) SetG(v int32) *StartStreamingOutRequestRegionColor {
	s.G = &v
	return s
}

func (s *StartStreamingOutRequestRegionColor) SetR(v int32) *StartStreamingOutRequestRegionColor {
	s.R = &v
	return s
}

func (s *StartStreamingOutRequestRegionColor) Validate() error {
	return dara.Validate(s)
}

type StartStreamingOutRequestTexts struct {
	// The watermark transparency.
	//
	// example:
	//
	// 0.1
	Alpha *float64 `json:"Alpha,omitempty" xml:"Alpha,omitempty"`
	// The transparency of the rectangle box.
	//
	// example:
	//
	// 0.6
	BoxAlpha *float64 `json:"BoxAlpha,omitempty" xml:"BoxAlpha,omitempty"`
	// The border width of the rectangle box.
	//
	// example:
	//
	// 5
	BoxBorderw *int32 `json:"BoxBorderw,omitempty" xml:"BoxBorderw,omitempty"`
	// The color of the rectangle box.
	BoxColor *StartStreamingOutRequestTextsBoxColor `json:"BoxColor,omitempty" xml:"BoxColor,omitempty" type:"Struct"`
	// The font. Valid values:
	//
	// - **0**: Alibaba PuHuiTi Regular
	//
	// - **1**: Alibaba PuHuiTi Light
	//
	// - **2**: Alibaba PuHuiTi Medium
	//
	// - **3**: Alibaba PuHuiTi Bold
	//
	// - **4**: Alibaba PuHuiTi Heavy
	//
	// - **5**: Alimama ShuHeiTi
	//
	// - **6**: DingTalk JinBuTi
	//
	// - **7**: Alimama DaoLiTi
	//
	// - **8**: Alimama DongFangDaKai
	//
	// - **9**: Alimama FangYuanTi
	//
	// example:
	//
	// 0
	Font *int32 `json:"Font,omitempty" xml:"Font,omitempty"`
	// The font color.
	FontColor *StartStreamingOutRequestTextsFontColor `json:"FontColor,omitempty" xml:"FontColor,omitempty" type:"Struct"`
	// The font size.
	//
	// example:
	//
	// 36
	FontSize *int32 `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	// Specifies whether a rectangle box is required.
	HasBox *bool `json:"HasBox,omitempty" xml:"HasBox,omitempty"`
	// The z-axis layer.
	//
	// example:
	//
	// 0
	Layer *int32 `json:"Layer,omitempty" xml:"Layer,omitempty"`
	// The watermark content.
	//
	// This parameter is required.
	//
	// example:
	//
	// Text watermark
	Texture *string `json:"Texture,omitempty" xml:"Texture,omitempty"`
	// The x-coordinate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	X *float64 `json:"X,omitempty" xml:"X,omitempty"`
	// The y-coordinate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.2
	Y *float64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s StartStreamingOutRequestTexts) String() string {
	return dara.Prettify(s)
}

func (s StartStreamingOutRequestTexts) GoString() string {
	return s.String()
}

func (s *StartStreamingOutRequestTexts) GetAlpha() *float64 {
	return s.Alpha
}

func (s *StartStreamingOutRequestTexts) GetBoxAlpha() *float64 {
	return s.BoxAlpha
}

func (s *StartStreamingOutRequestTexts) GetBoxBorderw() *int32 {
	return s.BoxBorderw
}

func (s *StartStreamingOutRequestTexts) GetBoxColor() *StartStreamingOutRequestTextsBoxColor {
	return s.BoxColor
}

func (s *StartStreamingOutRequestTexts) GetFont() *int32 {
	return s.Font
}

func (s *StartStreamingOutRequestTexts) GetFontColor() *StartStreamingOutRequestTextsFontColor {
	return s.FontColor
}

func (s *StartStreamingOutRequestTexts) GetFontSize() *int32 {
	return s.FontSize
}

func (s *StartStreamingOutRequestTexts) GetHasBox() *bool {
	return s.HasBox
}

func (s *StartStreamingOutRequestTexts) GetLayer() *int32 {
	return s.Layer
}

func (s *StartStreamingOutRequestTexts) GetTexture() *string {
	return s.Texture
}

func (s *StartStreamingOutRequestTexts) GetX() *float64 {
	return s.X
}

func (s *StartStreamingOutRequestTexts) GetY() *float64 {
	return s.Y
}

func (s *StartStreamingOutRequestTexts) SetAlpha(v float64) *StartStreamingOutRequestTexts {
	s.Alpha = &v
	return s
}

func (s *StartStreamingOutRequestTexts) SetBoxAlpha(v float64) *StartStreamingOutRequestTexts {
	s.BoxAlpha = &v
	return s
}

func (s *StartStreamingOutRequestTexts) SetBoxBorderw(v int32) *StartStreamingOutRequestTexts {
	s.BoxBorderw = &v
	return s
}

func (s *StartStreamingOutRequestTexts) SetBoxColor(v *StartStreamingOutRequestTextsBoxColor) *StartStreamingOutRequestTexts {
	s.BoxColor = v
	return s
}

func (s *StartStreamingOutRequestTexts) SetFont(v int32) *StartStreamingOutRequestTexts {
	s.Font = &v
	return s
}

func (s *StartStreamingOutRequestTexts) SetFontColor(v *StartStreamingOutRequestTextsFontColor) *StartStreamingOutRequestTexts {
	s.FontColor = v
	return s
}

func (s *StartStreamingOutRequestTexts) SetFontSize(v int32) *StartStreamingOutRequestTexts {
	s.FontSize = &v
	return s
}

func (s *StartStreamingOutRequestTexts) SetHasBox(v bool) *StartStreamingOutRequestTexts {
	s.HasBox = &v
	return s
}

func (s *StartStreamingOutRequestTexts) SetLayer(v int32) *StartStreamingOutRequestTexts {
	s.Layer = &v
	return s
}

func (s *StartStreamingOutRequestTexts) SetTexture(v string) *StartStreamingOutRequestTexts {
	s.Texture = &v
	return s
}

func (s *StartStreamingOutRequestTexts) SetX(v float64) *StartStreamingOutRequestTexts {
	s.X = &v
	return s
}

func (s *StartStreamingOutRequestTexts) SetY(v float64) *StartStreamingOutRequestTexts {
	s.Y = &v
	return s
}

func (s *StartStreamingOutRequestTexts) Validate() error {
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

type StartStreamingOutRequestTextsBoxColor struct {
	// The B value.
	//
	// example:
	//
	// 255
	B *int32 `json:"B,omitempty" xml:"B,omitempty"`
	// The G value.
	//
	// example:
	//
	// 255
	G *int32 `json:"G,omitempty" xml:"G,omitempty"`
	// The R value.
	//
	// example:
	//
	// 255
	R *int32 `json:"R,omitempty" xml:"R,omitempty"`
}

func (s StartStreamingOutRequestTextsBoxColor) String() string {
	return dara.Prettify(s)
}

func (s StartStreamingOutRequestTextsBoxColor) GoString() string {
	return s.String()
}

func (s *StartStreamingOutRequestTextsBoxColor) GetB() *int32 {
	return s.B
}

func (s *StartStreamingOutRequestTextsBoxColor) GetG() *int32 {
	return s.G
}

func (s *StartStreamingOutRequestTextsBoxColor) GetR() *int32 {
	return s.R
}

func (s *StartStreamingOutRequestTextsBoxColor) SetB(v int32) *StartStreamingOutRequestTextsBoxColor {
	s.B = &v
	return s
}

func (s *StartStreamingOutRequestTextsBoxColor) SetG(v int32) *StartStreamingOutRequestTextsBoxColor {
	s.G = &v
	return s
}

func (s *StartStreamingOutRequestTextsBoxColor) SetR(v int32) *StartStreamingOutRequestTextsBoxColor {
	s.R = &v
	return s
}

func (s *StartStreamingOutRequestTextsBoxColor) Validate() error {
	return dara.Validate(s)
}

type StartStreamingOutRequestTextsFontColor struct {
	// The B value.
	//
	// example:
	//
	// 255
	B *int32 `json:"B,omitempty" xml:"B,omitempty"`
	// The G value.
	//
	// example:
	//
	// 255
	G *int32 `json:"G,omitempty" xml:"G,omitempty"`
	// The R value.
	//
	// example:
	//
	// 255
	R *int32 `json:"R,omitempty" xml:"R,omitempty"`
}

func (s StartStreamingOutRequestTextsFontColor) String() string {
	return dara.Prettify(s)
}

func (s StartStreamingOutRequestTextsFontColor) GoString() string {
	return s.String()
}

func (s *StartStreamingOutRequestTextsFontColor) GetB() *int32 {
	return s.B
}

func (s *StartStreamingOutRequestTextsFontColor) GetG() *int32 {
	return s.G
}

func (s *StartStreamingOutRequestTextsFontColor) GetR() *int32 {
	return s.R
}

func (s *StartStreamingOutRequestTextsFontColor) SetB(v int32) *StartStreamingOutRequestTextsFontColor {
	s.B = &v
	return s
}

func (s *StartStreamingOutRequestTextsFontColor) SetG(v int32) *StartStreamingOutRequestTextsFontColor {
	s.G = &v
	return s
}

func (s *StartStreamingOutRequestTextsFontColor) SetR(v int32) *StartStreamingOutRequestTextsFontColor {
	s.R = &v
	return s
}

func (s *StartStreamingOutRequestTextsFontColor) Validate() error {
	return dara.Validate(s)
}
