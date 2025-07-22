// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCloudRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnnotation(v string) *StartCloudRecordRequest
	GetAnnotation() *string
	SetAppId(v string) *StartCloudRecordRequest
	GetAppId() *string
	SetBackgrounds(v []*StartCloudRecordRequestBackgrounds) *StartCloudRecordRequest
	GetBackgrounds() []*StartCloudRecordRequestBackgrounds
	SetBgColor(v *StartCloudRecordRequestBgColor) *StartCloudRecordRequest
	GetBgColor() *StartCloudRecordRequestBgColor
	SetChannelId(v string) *StartCloudRecordRequest
	GetChannelId() *string
	SetClockWidgets(v []*StartCloudRecordRequestClockWidgets) *StartCloudRecordRequest
	GetClockWidgets() []*StartCloudRecordRequestClockWidgets
	SetCropMode(v int32) *StartCloudRecordRequest
	GetCropMode() *int32
	SetImages(v []*StartCloudRecordRequestImages) *StartCloudRecordRequest
	GetImages() []*StartCloudRecordRequestImages
	SetLayoutSpecifiedUsers(v *StartCloudRecordRequestLayoutSpecifiedUsers) *StartCloudRecordRequest
	GetLayoutSpecifiedUsers() *StartCloudRecordRequestLayoutSpecifiedUsers
	SetPanes(v []*StartCloudRecordRequestPanes) *StartCloudRecordRequest
	GetPanes() []*StartCloudRecordRequestPanes
	SetRecordMode(v int32) *StartCloudRecordRequest
	GetRecordMode() *int32
	SetRegionColor(v *StartCloudRecordRequestRegionColor) *StartCloudRecordRequest
	GetRegionColor() *StartCloudRecordRequestRegionColor
	SetReservePaneForNoCameraUser(v bool) *StartCloudRecordRequest
	GetReservePaneForNoCameraUser() *bool
	SetShowDefaultBackgroundOnMute(v bool) *StartCloudRecordRequest
	GetShowDefaultBackgroundOnMute() *bool
	SetSingleStreamingRecord(v *StartCloudRecordRequestSingleStreamingRecord) *StartCloudRecordRequest
	GetSingleStreamingRecord() *StartCloudRecordRequestSingleStreamingRecord
	SetStartWithoutChannel(v bool) *StartCloudRecordRequest
	GetStartWithoutChannel() *bool
	SetStartWithoutChannelWaitTime(v int32) *StartCloudRecordRequest
	GetStartWithoutChannelWaitTime() *int32
	SetStorageConfig(v *StartCloudRecordRequestStorageConfig) *StartCloudRecordRequest
	GetStorageConfig() *StartCloudRecordRequestStorageConfig
	SetSubHighResolutionStream(v bool) *StartCloudRecordRequest
	GetSubHighResolutionStream() *bool
	SetTaskId(v string) *StartCloudRecordRequest
	GetTaskId() *string
	SetTemplateId(v string) *StartCloudRecordRequest
	GetTemplateId() *string
	SetTexts(v []*StartCloudRecordRequestTexts) *StartCloudRecordRequest
	GetTexts() []*StartCloudRecordRequestTexts
}

type StartCloudRecordRequest struct {
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
	AppId       *string                               `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Backgrounds []*StartCloudRecordRequestBackgrounds `json:"Backgrounds,omitempty" xml:"Backgrounds,omitempty" type:"Repeated"`
	BgColor     *StartCloudRecordRequestBgColor       `json:"BgColor,omitempty" xml:"BgColor,omitempty" type:"Struct"`
	// channelName
	//
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId    *string                                `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ClockWidgets []*StartCloudRecordRequestClockWidgets `json:"ClockWidgets,omitempty" xml:"ClockWidgets,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	CropMode             *int32                                       `json:"CropMode,omitempty" xml:"CropMode,omitempty"`
	Images               []*StartCloudRecordRequestImages             `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	LayoutSpecifiedUsers *StartCloudRecordRequestLayoutSpecifiedUsers `json:"LayoutSpecifiedUsers,omitempty" xml:"LayoutSpecifiedUsers,omitempty" type:"Struct"`
	// panes
	Panes                       []*StartCloudRecordRequestPanes               `json:"Panes,omitempty" xml:"Panes,omitempty" type:"Repeated"`
	RecordMode                  *int32                                        `json:"RecordMode,omitempty" xml:"RecordMode,omitempty"`
	RegionColor                 *StartCloudRecordRequestRegionColor           `json:"RegionColor,omitempty" xml:"RegionColor,omitempty" type:"Struct"`
	ReservePaneForNoCameraUser  *bool                                         `json:"ReservePaneForNoCameraUser,omitempty" xml:"ReservePaneForNoCameraUser,omitempty"`
	ShowDefaultBackgroundOnMute *bool                                         `json:"ShowDefaultBackgroundOnMute,omitempty" xml:"ShowDefaultBackgroundOnMute,omitempty"`
	SingleStreamingRecord       *StartCloudRecordRequestSingleStreamingRecord `json:"SingleStreamingRecord,omitempty" xml:"SingleStreamingRecord,omitempty" type:"Struct"`
	StartWithoutChannel         *bool                                         `json:"StartWithoutChannel,omitempty" xml:"StartWithoutChannel,omitempty"`
	// example:
	//
	// 30
	StartWithoutChannelWaitTime *int32 `json:"StartWithoutChannelWaitTime,omitempty" xml:"StartWithoutChannelWaitTime,omitempty"`
	// storageConfig
	//
	// This parameter is required.
	StorageConfig           *StartCloudRecordRequestStorageConfig `json:"StorageConfig,omitempty" xml:"StorageConfig,omitempty" type:"Struct"`
	SubHighResolutionStream *bool                                 `json:"SubHighResolutionStream,omitempty" xml:"SubHighResolutionStream,omitempty"`
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
	TemplateId *string                         `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Texts      []*StartCloudRecordRequestTexts `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
}

func (s StartCloudRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordRequest) GoString() string {
	return s.String()
}

func (s *StartCloudRecordRequest) GetAnnotation() *string {
	return s.Annotation
}

func (s *StartCloudRecordRequest) GetAppId() *string {
	return s.AppId
}

func (s *StartCloudRecordRequest) GetBackgrounds() []*StartCloudRecordRequestBackgrounds {
	return s.Backgrounds
}

func (s *StartCloudRecordRequest) GetBgColor() *StartCloudRecordRequestBgColor {
	return s.BgColor
}

func (s *StartCloudRecordRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *StartCloudRecordRequest) GetClockWidgets() []*StartCloudRecordRequestClockWidgets {
	return s.ClockWidgets
}

func (s *StartCloudRecordRequest) GetCropMode() *int32 {
	return s.CropMode
}

func (s *StartCloudRecordRequest) GetImages() []*StartCloudRecordRequestImages {
	return s.Images
}

func (s *StartCloudRecordRequest) GetLayoutSpecifiedUsers() *StartCloudRecordRequestLayoutSpecifiedUsers {
	return s.LayoutSpecifiedUsers
}

func (s *StartCloudRecordRequest) GetPanes() []*StartCloudRecordRequestPanes {
	return s.Panes
}

func (s *StartCloudRecordRequest) GetRecordMode() *int32 {
	return s.RecordMode
}

func (s *StartCloudRecordRequest) GetRegionColor() *StartCloudRecordRequestRegionColor {
	return s.RegionColor
}

func (s *StartCloudRecordRequest) GetReservePaneForNoCameraUser() *bool {
	return s.ReservePaneForNoCameraUser
}

func (s *StartCloudRecordRequest) GetShowDefaultBackgroundOnMute() *bool {
	return s.ShowDefaultBackgroundOnMute
}

func (s *StartCloudRecordRequest) GetSingleStreamingRecord() *StartCloudRecordRequestSingleStreamingRecord {
	return s.SingleStreamingRecord
}

func (s *StartCloudRecordRequest) GetStartWithoutChannel() *bool {
	return s.StartWithoutChannel
}

func (s *StartCloudRecordRequest) GetStartWithoutChannelWaitTime() *int32 {
	return s.StartWithoutChannelWaitTime
}

func (s *StartCloudRecordRequest) GetStorageConfig() *StartCloudRecordRequestStorageConfig {
	return s.StorageConfig
}

func (s *StartCloudRecordRequest) GetSubHighResolutionStream() *bool {
	return s.SubHighResolutionStream
}

func (s *StartCloudRecordRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *StartCloudRecordRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *StartCloudRecordRequest) GetTexts() []*StartCloudRecordRequestTexts {
	return s.Texts
}

func (s *StartCloudRecordRequest) SetAnnotation(v string) *StartCloudRecordRequest {
	s.Annotation = &v
	return s
}

func (s *StartCloudRecordRequest) SetAppId(v string) *StartCloudRecordRequest {
	s.AppId = &v
	return s
}

func (s *StartCloudRecordRequest) SetBackgrounds(v []*StartCloudRecordRequestBackgrounds) *StartCloudRecordRequest {
	s.Backgrounds = v
	return s
}

func (s *StartCloudRecordRequest) SetBgColor(v *StartCloudRecordRequestBgColor) *StartCloudRecordRequest {
	s.BgColor = v
	return s
}

func (s *StartCloudRecordRequest) SetChannelId(v string) *StartCloudRecordRequest {
	s.ChannelId = &v
	return s
}

func (s *StartCloudRecordRequest) SetClockWidgets(v []*StartCloudRecordRequestClockWidgets) *StartCloudRecordRequest {
	s.ClockWidgets = v
	return s
}

func (s *StartCloudRecordRequest) SetCropMode(v int32) *StartCloudRecordRequest {
	s.CropMode = &v
	return s
}

func (s *StartCloudRecordRequest) SetImages(v []*StartCloudRecordRequestImages) *StartCloudRecordRequest {
	s.Images = v
	return s
}

func (s *StartCloudRecordRequest) SetLayoutSpecifiedUsers(v *StartCloudRecordRequestLayoutSpecifiedUsers) *StartCloudRecordRequest {
	s.LayoutSpecifiedUsers = v
	return s
}

func (s *StartCloudRecordRequest) SetPanes(v []*StartCloudRecordRequestPanes) *StartCloudRecordRequest {
	s.Panes = v
	return s
}

func (s *StartCloudRecordRequest) SetRecordMode(v int32) *StartCloudRecordRequest {
	s.RecordMode = &v
	return s
}

func (s *StartCloudRecordRequest) SetRegionColor(v *StartCloudRecordRequestRegionColor) *StartCloudRecordRequest {
	s.RegionColor = v
	return s
}

func (s *StartCloudRecordRequest) SetReservePaneForNoCameraUser(v bool) *StartCloudRecordRequest {
	s.ReservePaneForNoCameraUser = &v
	return s
}

func (s *StartCloudRecordRequest) SetShowDefaultBackgroundOnMute(v bool) *StartCloudRecordRequest {
	s.ShowDefaultBackgroundOnMute = &v
	return s
}

func (s *StartCloudRecordRequest) SetSingleStreamingRecord(v *StartCloudRecordRequestSingleStreamingRecord) *StartCloudRecordRequest {
	s.SingleStreamingRecord = v
	return s
}

func (s *StartCloudRecordRequest) SetStartWithoutChannel(v bool) *StartCloudRecordRequest {
	s.StartWithoutChannel = &v
	return s
}

func (s *StartCloudRecordRequest) SetStartWithoutChannelWaitTime(v int32) *StartCloudRecordRequest {
	s.StartWithoutChannelWaitTime = &v
	return s
}

func (s *StartCloudRecordRequest) SetStorageConfig(v *StartCloudRecordRequestStorageConfig) *StartCloudRecordRequest {
	s.StorageConfig = v
	return s
}

func (s *StartCloudRecordRequest) SetSubHighResolutionStream(v bool) *StartCloudRecordRequest {
	s.SubHighResolutionStream = &v
	return s
}

func (s *StartCloudRecordRequest) SetTaskId(v string) *StartCloudRecordRequest {
	s.TaskId = &v
	return s
}

func (s *StartCloudRecordRequest) SetTemplateId(v string) *StartCloudRecordRequest {
	s.TemplateId = &v
	return s
}

func (s *StartCloudRecordRequest) SetTexts(v []*StartCloudRecordRequestTexts) *StartCloudRecordRequest {
	s.Texts = v
	return s
}

func (s *StartCloudRecordRequest) Validate() error {
	return dara.Validate(s)
}

type StartCloudRecordRequestBackgrounds struct {
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

func (s StartCloudRecordRequestBackgrounds) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordRequestBackgrounds) GoString() string {
	return s.String()
}

func (s *StartCloudRecordRequestBackgrounds) GetAlpha() *float64 {
	return s.Alpha
}

func (s *StartCloudRecordRequestBackgrounds) GetBackgroundCropMode() *int32 {
	return s.BackgroundCropMode
}

func (s *StartCloudRecordRequestBackgrounds) GetHeight() *float64 {
	return s.Height
}

func (s *StartCloudRecordRequestBackgrounds) GetLayer() *int32 {
	return s.Layer
}

func (s *StartCloudRecordRequestBackgrounds) GetUrl() *string {
	return s.Url
}

func (s *StartCloudRecordRequestBackgrounds) GetWidth() *float64 {
	return s.Width
}

func (s *StartCloudRecordRequestBackgrounds) GetX() *float64 {
	return s.X
}

func (s *StartCloudRecordRequestBackgrounds) GetY() *float64 {
	return s.Y
}

func (s *StartCloudRecordRequestBackgrounds) SetAlpha(v float64) *StartCloudRecordRequestBackgrounds {
	s.Alpha = &v
	return s
}

func (s *StartCloudRecordRequestBackgrounds) SetBackgroundCropMode(v int32) *StartCloudRecordRequestBackgrounds {
	s.BackgroundCropMode = &v
	return s
}

func (s *StartCloudRecordRequestBackgrounds) SetHeight(v float64) *StartCloudRecordRequestBackgrounds {
	s.Height = &v
	return s
}

func (s *StartCloudRecordRequestBackgrounds) SetLayer(v int32) *StartCloudRecordRequestBackgrounds {
	s.Layer = &v
	return s
}

func (s *StartCloudRecordRequestBackgrounds) SetUrl(v string) *StartCloudRecordRequestBackgrounds {
	s.Url = &v
	return s
}

func (s *StartCloudRecordRequestBackgrounds) SetWidth(v float64) *StartCloudRecordRequestBackgrounds {
	s.Width = &v
	return s
}

func (s *StartCloudRecordRequestBackgrounds) SetX(v float64) *StartCloudRecordRequestBackgrounds {
	s.X = &v
	return s
}

func (s *StartCloudRecordRequestBackgrounds) SetY(v float64) *StartCloudRecordRequestBackgrounds {
	s.Y = &v
	return s
}

func (s *StartCloudRecordRequestBackgrounds) Validate() error {
	return dara.Validate(s)
}

type StartCloudRecordRequestBgColor struct {
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

func (s StartCloudRecordRequestBgColor) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordRequestBgColor) GoString() string {
	return s.String()
}

func (s *StartCloudRecordRequestBgColor) GetB() *int32 {
	return s.B
}

func (s *StartCloudRecordRequestBgColor) GetG() *int32 {
	return s.G
}

func (s *StartCloudRecordRequestBgColor) GetR() *int32 {
	return s.R
}

func (s *StartCloudRecordRequestBgColor) SetB(v int32) *StartCloudRecordRequestBgColor {
	s.B = &v
	return s
}

func (s *StartCloudRecordRequestBgColor) SetG(v int32) *StartCloudRecordRequestBgColor {
	s.G = &v
	return s
}

func (s *StartCloudRecordRequestBgColor) SetR(v int32) *StartCloudRecordRequestBgColor {
	s.R = &v
	return s
}

func (s *StartCloudRecordRequestBgColor) Validate() error {
	return dara.Validate(s)
}

type StartCloudRecordRequestClockWidgets struct {
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
	BoxColor   *StartCloudRecordRequestClockWidgetsBoxColor `json:"BoxColor,omitempty" xml:"BoxColor,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Font      *int32                                        `json:"Font,omitempty" xml:"Font,omitempty"`
	FontColor *StartCloudRecordRequestClockWidgetsFontColor `json:"FontColor,omitempty" xml:"FontColor,omitempty" type:"Struct"`
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

func (s StartCloudRecordRequestClockWidgets) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordRequestClockWidgets) GoString() string {
	return s.String()
}

func (s *StartCloudRecordRequestClockWidgets) GetAlpha() *float64 {
	return s.Alpha
}

func (s *StartCloudRecordRequestClockWidgets) GetBoxAlpha() *float64 {
	return s.BoxAlpha
}

func (s *StartCloudRecordRequestClockWidgets) GetBoxBorderw() *int32 {
	return s.BoxBorderw
}

func (s *StartCloudRecordRequestClockWidgets) GetBoxColor() *StartCloudRecordRequestClockWidgetsBoxColor {
	return s.BoxColor
}

func (s *StartCloudRecordRequestClockWidgets) GetFont() *int32 {
	return s.Font
}

func (s *StartCloudRecordRequestClockWidgets) GetFontColor() *StartCloudRecordRequestClockWidgetsFontColor {
	return s.FontColor
}

func (s *StartCloudRecordRequestClockWidgets) GetFontSize() *int32 {
	return s.FontSize
}

func (s *StartCloudRecordRequestClockWidgets) GetHasBox() *bool {
	return s.HasBox
}

func (s *StartCloudRecordRequestClockWidgets) GetLayer() *int32 {
	return s.Layer
}

func (s *StartCloudRecordRequestClockWidgets) GetX() *float64 {
	return s.X
}

func (s *StartCloudRecordRequestClockWidgets) GetY() *float64 {
	return s.Y
}

func (s *StartCloudRecordRequestClockWidgets) GetZone() *int32 {
	return s.Zone
}

func (s *StartCloudRecordRequestClockWidgets) SetAlpha(v float64) *StartCloudRecordRequestClockWidgets {
	s.Alpha = &v
	return s
}

func (s *StartCloudRecordRequestClockWidgets) SetBoxAlpha(v float64) *StartCloudRecordRequestClockWidgets {
	s.BoxAlpha = &v
	return s
}

func (s *StartCloudRecordRequestClockWidgets) SetBoxBorderw(v int32) *StartCloudRecordRequestClockWidgets {
	s.BoxBorderw = &v
	return s
}

func (s *StartCloudRecordRequestClockWidgets) SetBoxColor(v *StartCloudRecordRequestClockWidgetsBoxColor) *StartCloudRecordRequestClockWidgets {
	s.BoxColor = v
	return s
}

func (s *StartCloudRecordRequestClockWidgets) SetFont(v int32) *StartCloudRecordRequestClockWidgets {
	s.Font = &v
	return s
}

func (s *StartCloudRecordRequestClockWidgets) SetFontColor(v *StartCloudRecordRequestClockWidgetsFontColor) *StartCloudRecordRequestClockWidgets {
	s.FontColor = v
	return s
}

func (s *StartCloudRecordRequestClockWidgets) SetFontSize(v int32) *StartCloudRecordRequestClockWidgets {
	s.FontSize = &v
	return s
}

func (s *StartCloudRecordRequestClockWidgets) SetHasBox(v bool) *StartCloudRecordRequestClockWidgets {
	s.HasBox = &v
	return s
}

func (s *StartCloudRecordRequestClockWidgets) SetLayer(v int32) *StartCloudRecordRequestClockWidgets {
	s.Layer = &v
	return s
}

func (s *StartCloudRecordRequestClockWidgets) SetX(v float64) *StartCloudRecordRequestClockWidgets {
	s.X = &v
	return s
}

func (s *StartCloudRecordRequestClockWidgets) SetY(v float64) *StartCloudRecordRequestClockWidgets {
	s.Y = &v
	return s
}

func (s *StartCloudRecordRequestClockWidgets) SetZone(v int32) *StartCloudRecordRequestClockWidgets {
	s.Zone = &v
	return s
}

func (s *StartCloudRecordRequestClockWidgets) Validate() error {
	return dara.Validate(s)
}

type StartCloudRecordRequestClockWidgetsBoxColor struct {
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

func (s StartCloudRecordRequestClockWidgetsBoxColor) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordRequestClockWidgetsBoxColor) GoString() string {
	return s.String()
}

func (s *StartCloudRecordRequestClockWidgetsBoxColor) GetB() *int32 {
	return s.B
}

func (s *StartCloudRecordRequestClockWidgetsBoxColor) GetG() *int32 {
	return s.G
}

func (s *StartCloudRecordRequestClockWidgetsBoxColor) GetR() *int32 {
	return s.R
}

func (s *StartCloudRecordRequestClockWidgetsBoxColor) SetB(v int32) *StartCloudRecordRequestClockWidgetsBoxColor {
	s.B = &v
	return s
}

func (s *StartCloudRecordRequestClockWidgetsBoxColor) SetG(v int32) *StartCloudRecordRequestClockWidgetsBoxColor {
	s.G = &v
	return s
}

func (s *StartCloudRecordRequestClockWidgetsBoxColor) SetR(v int32) *StartCloudRecordRequestClockWidgetsBoxColor {
	s.R = &v
	return s
}

func (s *StartCloudRecordRequestClockWidgetsBoxColor) Validate() error {
	return dara.Validate(s)
}

type StartCloudRecordRequestClockWidgetsFontColor struct {
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

func (s StartCloudRecordRequestClockWidgetsFontColor) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordRequestClockWidgetsFontColor) GoString() string {
	return s.String()
}

func (s *StartCloudRecordRequestClockWidgetsFontColor) GetB() *int32 {
	return s.B
}

func (s *StartCloudRecordRequestClockWidgetsFontColor) GetG() *int32 {
	return s.G
}

func (s *StartCloudRecordRequestClockWidgetsFontColor) GetR() *int32 {
	return s.R
}

func (s *StartCloudRecordRequestClockWidgetsFontColor) SetB(v int32) *StartCloudRecordRequestClockWidgetsFontColor {
	s.B = &v
	return s
}

func (s *StartCloudRecordRequestClockWidgetsFontColor) SetG(v int32) *StartCloudRecordRequestClockWidgetsFontColor {
	s.G = &v
	return s
}

func (s *StartCloudRecordRequestClockWidgetsFontColor) SetR(v int32) *StartCloudRecordRequestClockWidgetsFontColor {
	s.R = &v
	return s
}

func (s *StartCloudRecordRequestClockWidgetsFontColor) Validate() error {
	return dara.Validate(s)
}

type StartCloudRecordRequestImages struct {
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

func (s StartCloudRecordRequestImages) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordRequestImages) GoString() string {
	return s.String()
}

func (s *StartCloudRecordRequestImages) GetAlpha() *float64 {
	return s.Alpha
}

func (s *StartCloudRecordRequestImages) GetHeight() *float64 {
	return s.Height
}

func (s *StartCloudRecordRequestImages) GetImageCropMode() *int32 {
	return s.ImageCropMode
}

func (s *StartCloudRecordRequestImages) GetLayer() *int32 {
	return s.Layer
}

func (s *StartCloudRecordRequestImages) GetUrl() *string {
	return s.Url
}

func (s *StartCloudRecordRequestImages) GetWidth() *float64 {
	return s.Width
}

func (s *StartCloudRecordRequestImages) GetX() *float64 {
	return s.X
}

func (s *StartCloudRecordRequestImages) GetY() *float64 {
	return s.Y
}

func (s *StartCloudRecordRequestImages) SetAlpha(v float64) *StartCloudRecordRequestImages {
	s.Alpha = &v
	return s
}

func (s *StartCloudRecordRequestImages) SetHeight(v float64) *StartCloudRecordRequestImages {
	s.Height = &v
	return s
}

func (s *StartCloudRecordRequestImages) SetImageCropMode(v int32) *StartCloudRecordRequestImages {
	s.ImageCropMode = &v
	return s
}

func (s *StartCloudRecordRequestImages) SetLayer(v int32) *StartCloudRecordRequestImages {
	s.Layer = &v
	return s
}

func (s *StartCloudRecordRequestImages) SetUrl(v string) *StartCloudRecordRequestImages {
	s.Url = &v
	return s
}

func (s *StartCloudRecordRequestImages) SetWidth(v float64) *StartCloudRecordRequestImages {
	s.Width = &v
	return s
}

func (s *StartCloudRecordRequestImages) SetX(v float64) *StartCloudRecordRequestImages {
	s.X = &v
	return s
}

func (s *StartCloudRecordRequestImages) SetY(v float64) *StartCloudRecordRequestImages {
	s.Y = &v
	return s
}

func (s *StartCloudRecordRequestImages) Validate() error {
	return dara.Validate(s)
}

type StartCloudRecordRequestLayoutSpecifiedUsers struct {
	// This parameter is required.
	Ids []*string `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s StartCloudRecordRequestLayoutSpecifiedUsers) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordRequestLayoutSpecifiedUsers) GoString() string {
	return s.String()
}

func (s *StartCloudRecordRequestLayoutSpecifiedUsers) GetIds() []*string {
	return s.Ids
}

func (s *StartCloudRecordRequestLayoutSpecifiedUsers) GetType() *string {
	return s.Type
}

func (s *StartCloudRecordRequestLayoutSpecifiedUsers) SetIds(v []*string) *StartCloudRecordRequestLayoutSpecifiedUsers {
	s.Ids = v
	return s
}

func (s *StartCloudRecordRequestLayoutSpecifiedUsers) SetType(v string) *StartCloudRecordRequestLayoutSpecifiedUsers {
	s.Type = &v
	return s
}

func (s *StartCloudRecordRequestLayoutSpecifiedUsers) Validate() error {
	return dara.Validate(s)
}

type StartCloudRecordRequestPanes struct {
	Backgrounds []*StartCloudRecordRequestPanesBackgrounds `json:"Backgrounds,omitempty" xml:"Backgrounds,omitempty" type:"Repeated"`
	Images      []*StartCloudRecordRequestPanesImages      `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
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
	SourceType *string                              `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Texts      []*StartCloudRecordRequestPanesTexts `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
	// example:
	//
	// cameraFirst
	VideoOrder *string                                 `json:"VideoOrder,omitempty" xml:"VideoOrder,omitempty"`
	Whiteboard *StartCloudRecordRequestPanesWhiteboard `json:"Whiteboard,omitempty" xml:"Whiteboard,omitempty" type:"Struct"`
}

func (s StartCloudRecordRequestPanes) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordRequestPanes) GoString() string {
	return s.String()
}

func (s *StartCloudRecordRequestPanes) GetBackgrounds() []*StartCloudRecordRequestPanesBackgrounds {
	return s.Backgrounds
}

func (s *StartCloudRecordRequestPanes) GetImages() []*StartCloudRecordRequestPanesImages {
	return s.Images
}

func (s *StartCloudRecordRequestPanes) GetPaneCropMode() *int32 {
	return s.PaneCropMode
}

func (s *StartCloudRecordRequestPanes) GetPaneId() *int32 {
	return s.PaneId
}

func (s *StartCloudRecordRequestPanes) GetReservePaneForOfflineUser() *bool {
	return s.ReservePaneForOfflineUser
}

func (s *StartCloudRecordRequestPanes) GetSource() *string {
	return s.Source
}

func (s *StartCloudRecordRequestPanes) GetSourceType() *string {
	return s.SourceType
}

func (s *StartCloudRecordRequestPanes) GetTexts() []*StartCloudRecordRequestPanesTexts {
	return s.Texts
}

func (s *StartCloudRecordRequestPanes) GetVideoOrder() *string {
	return s.VideoOrder
}

func (s *StartCloudRecordRequestPanes) GetWhiteboard() *StartCloudRecordRequestPanesWhiteboard {
	return s.Whiteboard
}

func (s *StartCloudRecordRequestPanes) SetBackgrounds(v []*StartCloudRecordRequestPanesBackgrounds) *StartCloudRecordRequestPanes {
	s.Backgrounds = v
	return s
}

func (s *StartCloudRecordRequestPanes) SetImages(v []*StartCloudRecordRequestPanesImages) *StartCloudRecordRequestPanes {
	s.Images = v
	return s
}

func (s *StartCloudRecordRequestPanes) SetPaneCropMode(v int32) *StartCloudRecordRequestPanes {
	s.PaneCropMode = &v
	return s
}

func (s *StartCloudRecordRequestPanes) SetPaneId(v int32) *StartCloudRecordRequestPanes {
	s.PaneId = &v
	return s
}

func (s *StartCloudRecordRequestPanes) SetReservePaneForOfflineUser(v bool) *StartCloudRecordRequestPanes {
	s.ReservePaneForOfflineUser = &v
	return s
}

func (s *StartCloudRecordRequestPanes) SetSource(v string) *StartCloudRecordRequestPanes {
	s.Source = &v
	return s
}

func (s *StartCloudRecordRequestPanes) SetSourceType(v string) *StartCloudRecordRequestPanes {
	s.SourceType = &v
	return s
}

func (s *StartCloudRecordRequestPanes) SetTexts(v []*StartCloudRecordRequestPanesTexts) *StartCloudRecordRequestPanes {
	s.Texts = v
	return s
}

func (s *StartCloudRecordRequestPanes) SetVideoOrder(v string) *StartCloudRecordRequestPanes {
	s.VideoOrder = &v
	return s
}

func (s *StartCloudRecordRequestPanes) SetWhiteboard(v *StartCloudRecordRequestPanesWhiteboard) *StartCloudRecordRequestPanes {
	s.Whiteboard = v
	return s
}

func (s *StartCloudRecordRequestPanes) Validate() error {
	return dara.Validate(s)
}

type StartCloudRecordRequestPanesBackgrounds struct {
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

func (s StartCloudRecordRequestPanesBackgrounds) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordRequestPanesBackgrounds) GoString() string {
	return s.String()
}

func (s *StartCloudRecordRequestPanesBackgrounds) GetAlpha() *float64 {
	return s.Alpha
}

func (s *StartCloudRecordRequestPanesBackgrounds) GetDisplay() *string {
	return s.Display
}

func (s *StartCloudRecordRequestPanesBackgrounds) GetHeight() *float64 {
	return s.Height
}

func (s *StartCloudRecordRequestPanesBackgrounds) GetLayer() *int32 {
	return s.Layer
}

func (s *StartCloudRecordRequestPanesBackgrounds) GetPaneBackgroundCropMode() *int32 {
	return s.PaneBackgroundCropMode
}

func (s *StartCloudRecordRequestPanesBackgrounds) GetUrl() *string {
	return s.Url
}

func (s *StartCloudRecordRequestPanesBackgrounds) GetWidth() *float64 {
	return s.Width
}

func (s *StartCloudRecordRequestPanesBackgrounds) GetX() *float64 {
	return s.X
}

func (s *StartCloudRecordRequestPanesBackgrounds) GetY() *float64 {
	return s.Y
}

func (s *StartCloudRecordRequestPanesBackgrounds) SetAlpha(v float64) *StartCloudRecordRequestPanesBackgrounds {
	s.Alpha = &v
	return s
}

func (s *StartCloudRecordRequestPanesBackgrounds) SetDisplay(v string) *StartCloudRecordRequestPanesBackgrounds {
	s.Display = &v
	return s
}

func (s *StartCloudRecordRequestPanesBackgrounds) SetHeight(v float64) *StartCloudRecordRequestPanesBackgrounds {
	s.Height = &v
	return s
}

func (s *StartCloudRecordRequestPanesBackgrounds) SetLayer(v int32) *StartCloudRecordRequestPanesBackgrounds {
	s.Layer = &v
	return s
}

func (s *StartCloudRecordRequestPanesBackgrounds) SetPaneBackgroundCropMode(v int32) *StartCloudRecordRequestPanesBackgrounds {
	s.PaneBackgroundCropMode = &v
	return s
}

func (s *StartCloudRecordRequestPanesBackgrounds) SetUrl(v string) *StartCloudRecordRequestPanesBackgrounds {
	s.Url = &v
	return s
}

func (s *StartCloudRecordRequestPanesBackgrounds) SetWidth(v float64) *StartCloudRecordRequestPanesBackgrounds {
	s.Width = &v
	return s
}

func (s *StartCloudRecordRequestPanesBackgrounds) SetX(v float64) *StartCloudRecordRequestPanesBackgrounds {
	s.X = &v
	return s
}

func (s *StartCloudRecordRequestPanesBackgrounds) SetY(v float64) *StartCloudRecordRequestPanesBackgrounds {
	s.Y = &v
	return s
}

func (s *StartCloudRecordRequestPanesBackgrounds) Validate() error {
	return dara.Validate(s)
}

type StartCloudRecordRequestPanesImages struct {
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

func (s StartCloudRecordRequestPanesImages) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordRequestPanesImages) GoString() string {
	return s.String()
}

func (s *StartCloudRecordRequestPanesImages) GetAlpha() *float64 {
	return s.Alpha
}

func (s *StartCloudRecordRequestPanesImages) GetDisplay() *string {
	return s.Display
}

func (s *StartCloudRecordRequestPanesImages) GetHeight() *float64 {
	return s.Height
}

func (s *StartCloudRecordRequestPanesImages) GetLayer() *int32 {
	return s.Layer
}

func (s *StartCloudRecordRequestPanesImages) GetPaneImageCropMode() *int32 {
	return s.PaneImageCropMode
}

func (s *StartCloudRecordRequestPanesImages) GetUrl() *string {
	return s.Url
}

func (s *StartCloudRecordRequestPanesImages) GetWidth() *float64 {
	return s.Width
}

func (s *StartCloudRecordRequestPanesImages) GetX() *float64 {
	return s.X
}

func (s *StartCloudRecordRequestPanesImages) GetY() *float64 {
	return s.Y
}

func (s *StartCloudRecordRequestPanesImages) SetAlpha(v float64) *StartCloudRecordRequestPanesImages {
	s.Alpha = &v
	return s
}

func (s *StartCloudRecordRequestPanesImages) SetDisplay(v string) *StartCloudRecordRequestPanesImages {
	s.Display = &v
	return s
}

func (s *StartCloudRecordRequestPanesImages) SetHeight(v float64) *StartCloudRecordRequestPanesImages {
	s.Height = &v
	return s
}

func (s *StartCloudRecordRequestPanesImages) SetLayer(v int32) *StartCloudRecordRequestPanesImages {
	s.Layer = &v
	return s
}

func (s *StartCloudRecordRequestPanesImages) SetPaneImageCropMode(v int32) *StartCloudRecordRequestPanesImages {
	s.PaneImageCropMode = &v
	return s
}

func (s *StartCloudRecordRequestPanesImages) SetUrl(v string) *StartCloudRecordRequestPanesImages {
	s.Url = &v
	return s
}

func (s *StartCloudRecordRequestPanesImages) SetWidth(v float64) *StartCloudRecordRequestPanesImages {
	s.Width = &v
	return s
}

func (s *StartCloudRecordRequestPanesImages) SetX(v float64) *StartCloudRecordRequestPanesImages {
	s.X = &v
	return s
}

func (s *StartCloudRecordRequestPanesImages) SetY(v float64) *StartCloudRecordRequestPanesImages {
	s.Y = &v
	return s
}

func (s *StartCloudRecordRequestPanesImages) Validate() error {
	return dara.Validate(s)
}

type StartCloudRecordRequestPanesTexts struct {
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
	BoxColor   *StartCloudRecordRequestPanesTextsBoxColor `json:"BoxColor,omitempty" xml:"BoxColor,omitempty" type:"Struct"`
	// example:
	//
	// backup
	Display *string `json:"Display,omitempty" xml:"Display,omitempty"`
	// example:
	//
	// 0
	Font      *int32                                      `json:"Font,omitempty" xml:"Font,omitempty"`
	FontColor *StartCloudRecordRequestPanesTextsFontColor `json:"FontColor,omitempty" xml:"FontColor,omitempty" type:"Struct"`
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

func (s StartCloudRecordRequestPanesTexts) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordRequestPanesTexts) GoString() string {
	return s.String()
}

func (s *StartCloudRecordRequestPanesTexts) GetAlpha() *float64 {
	return s.Alpha
}

func (s *StartCloudRecordRequestPanesTexts) GetBoxAlpha() *float64 {
	return s.BoxAlpha
}

func (s *StartCloudRecordRequestPanesTexts) GetBoxBorderw() *int32 {
	return s.BoxBorderw
}

func (s *StartCloudRecordRequestPanesTexts) GetBoxColor() *StartCloudRecordRequestPanesTextsBoxColor {
	return s.BoxColor
}

func (s *StartCloudRecordRequestPanesTexts) GetDisplay() *string {
	return s.Display
}

func (s *StartCloudRecordRequestPanesTexts) GetFont() *int32 {
	return s.Font
}

func (s *StartCloudRecordRequestPanesTexts) GetFontColor() *StartCloudRecordRequestPanesTextsFontColor {
	return s.FontColor
}

func (s *StartCloudRecordRequestPanesTexts) GetFontSize() *int32 {
	return s.FontSize
}

func (s *StartCloudRecordRequestPanesTexts) GetHasBox() *bool {
	return s.HasBox
}

func (s *StartCloudRecordRequestPanesTexts) GetLayer() *int32 {
	return s.Layer
}

func (s *StartCloudRecordRequestPanesTexts) GetTexture() *string {
	return s.Texture
}

func (s *StartCloudRecordRequestPanesTexts) GetX() *float64 {
	return s.X
}

func (s *StartCloudRecordRequestPanesTexts) GetY() *float64 {
	return s.Y
}

func (s *StartCloudRecordRequestPanesTexts) SetAlpha(v float64) *StartCloudRecordRequestPanesTexts {
	s.Alpha = &v
	return s
}

func (s *StartCloudRecordRequestPanesTexts) SetBoxAlpha(v float64) *StartCloudRecordRequestPanesTexts {
	s.BoxAlpha = &v
	return s
}

func (s *StartCloudRecordRequestPanesTexts) SetBoxBorderw(v int32) *StartCloudRecordRequestPanesTexts {
	s.BoxBorderw = &v
	return s
}

func (s *StartCloudRecordRequestPanesTexts) SetBoxColor(v *StartCloudRecordRequestPanesTextsBoxColor) *StartCloudRecordRequestPanesTexts {
	s.BoxColor = v
	return s
}

func (s *StartCloudRecordRequestPanesTexts) SetDisplay(v string) *StartCloudRecordRequestPanesTexts {
	s.Display = &v
	return s
}

func (s *StartCloudRecordRequestPanesTexts) SetFont(v int32) *StartCloudRecordRequestPanesTexts {
	s.Font = &v
	return s
}

func (s *StartCloudRecordRequestPanesTexts) SetFontColor(v *StartCloudRecordRequestPanesTextsFontColor) *StartCloudRecordRequestPanesTexts {
	s.FontColor = v
	return s
}

func (s *StartCloudRecordRequestPanesTexts) SetFontSize(v int32) *StartCloudRecordRequestPanesTexts {
	s.FontSize = &v
	return s
}

func (s *StartCloudRecordRequestPanesTexts) SetHasBox(v bool) *StartCloudRecordRequestPanesTexts {
	s.HasBox = &v
	return s
}

func (s *StartCloudRecordRequestPanesTexts) SetLayer(v int32) *StartCloudRecordRequestPanesTexts {
	s.Layer = &v
	return s
}

func (s *StartCloudRecordRequestPanesTexts) SetTexture(v string) *StartCloudRecordRequestPanesTexts {
	s.Texture = &v
	return s
}

func (s *StartCloudRecordRequestPanesTexts) SetX(v float64) *StartCloudRecordRequestPanesTexts {
	s.X = &v
	return s
}

func (s *StartCloudRecordRequestPanesTexts) SetY(v float64) *StartCloudRecordRequestPanesTexts {
	s.Y = &v
	return s
}

func (s *StartCloudRecordRequestPanesTexts) Validate() error {
	return dara.Validate(s)
}

type StartCloudRecordRequestPanesTextsBoxColor struct {
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

func (s StartCloudRecordRequestPanesTextsBoxColor) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordRequestPanesTextsBoxColor) GoString() string {
	return s.String()
}

func (s *StartCloudRecordRequestPanesTextsBoxColor) GetB() *int32 {
	return s.B
}

func (s *StartCloudRecordRequestPanesTextsBoxColor) GetG() *int32 {
	return s.G
}

func (s *StartCloudRecordRequestPanesTextsBoxColor) GetR() *int32 {
	return s.R
}

func (s *StartCloudRecordRequestPanesTextsBoxColor) SetB(v int32) *StartCloudRecordRequestPanesTextsBoxColor {
	s.B = &v
	return s
}

func (s *StartCloudRecordRequestPanesTextsBoxColor) SetG(v int32) *StartCloudRecordRequestPanesTextsBoxColor {
	s.G = &v
	return s
}

func (s *StartCloudRecordRequestPanesTextsBoxColor) SetR(v int32) *StartCloudRecordRequestPanesTextsBoxColor {
	s.R = &v
	return s
}

func (s *StartCloudRecordRequestPanesTextsBoxColor) Validate() error {
	return dara.Validate(s)
}

type StartCloudRecordRequestPanesTextsFontColor struct {
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

func (s StartCloudRecordRequestPanesTextsFontColor) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordRequestPanesTextsFontColor) GoString() string {
	return s.String()
}

func (s *StartCloudRecordRequestPanesTextsFontColor) GetB() *int32 {
	return s.B
}

func (s *StartCloudRecordRequestPanesTextsFontColor) GetG() *int32 {
	return s.G
}

func (s *StartCloudRecordRequestPanesTextsFontColor) GetR() *int32 {
	return s.R
}

func (s *StartCloudRecordRequestPanesTextsFontColor) SetB(v int32) *StartCloudRecordRequestPanesTextsFontColor {
	s.B = &v
	return s
}

func (s *StartCloudRecordRequestPanesTextsFontColor) SetG(v int32) *StartCloudRecordRequestPanesTextsFontColor {
	s.G = &v
	return s
}

func (s *StartCloudRecordRequestPanesTextsFontColor) SetR(v int32) *StartCloudRecordRequestPanesTextsFontColor {
	s.R = &v
	return s
}

func (s *StartCloudRecordRequestPanesTextsFontColor) Validate() error {
	return dara.Validate(s)
}

type StartCloudRecordRequestPanesWhiteboard struct {
	// example:
	//
	// default
	WhiteboardId *string `json:"WhiteboardId,omitempty" xml:"WhiteboardId,omitempty"`
}

func (s StartCloudRecordRequestPanesWhiteboard) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordRequestPanesWhiteboard) GoString() string {
	return s.String()
}

func (s *StartCloudRecordRequestPanesWhiteboard) GetWhiteboardId() *string {
	return s.WhiteboardId
}

func (s *StartCloudRecordRequestPanesWhiteboard) SetWhiteboardId(v string) *StartCloudRecordRequestPanesWhiteboard {
	s.WhiteboardId = &v
	return s
}

func (s *StartCloudRecordRequestPanesWhiteboard) Validate() error {
	return dara.Validate(s)
}

type StartCloudRecordRequestRegionColor struct {
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

func (s StartCloudRecordRequestRegionColor) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordRequestRegionColor) GoString() string {
	return s.String()
}

func (s *StartCloudRecordRequestRegionColor) GetB() *int32 {
	return s.B
}

func (s *StartCloudRecordRequestRegionColor) GetG() *int32 {
	return s.G
}

func (s *StartCloudRecordRequestRegionColor) GetR() *int32 {
	return s.R
}

func (s *StartCloudRecordRequestRegionColor) SetB(v int32) *StartCloudRecordRequestRegionColor {
	s.B = &v
	return s
}

func (s *StartCloudRecordRequestRegionColor) SetG(v int32) *StartCloudRecordRequestRegionColor {
	s.G = &v
	return s
}

func (s *StartCloudRecordRequestRegionColor) SetR(v int32) *StartCloudRecordRequestRegionColor {
	s.R = &v
	return s
}

func (s *StartCloudRecordRequestRegionColor) Validate() error {
	return dara.Validate(s)
}

type StartCloudRecordRequestSingleStreamingRecord struct {
	// This parameter is required.
	SpecifiedStreams      []*StartCloudRecordRequestSingleStreamingRecordSpecifiedStreams    `json:"SpecifiedStreams,omitempty" xml:"SpecifiedStreams,omitempty" type:"Repeated"`
	TranscodingParameters *StartCloudRecordRequestSingleStreamingRecordTranscodingParameters `json:"TranscodingParameters,omitempty" xml:"TranscodingParameters,omitempty" type:"Struct"`
}

func (s StartCloudRecordRequestSingleStreamingRecord) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordRequestSingleStreamingRecord) GoString() string {
	return s.String()
}

func (s *StartCloudRecordRequestSingleStreamingRecord) GetSpecifiedStreams() []*StartCloudRecordRequestSingleStreamingRecordSpecifiedStreams {
	return s.SpecifiedStreams
}

func (s *StartCloudRecordRequestSingleStreamingRecord) GetTranscodingParameters() *StartCloudRecordRequestSingleStreamingRecordTranscodingParameters {
	return s.TranscodingParameters
}

func (s *StartCloudRecordRequestSingleStreamingRecord) SetSpecifiedStreams(v []*StartCloudRecordRequestSingleStreamingRecordSpecifiedStreams) *StartCloudRecordRequestSingleStreamingRecord {
	s.SpecifiedStreams = v
	return s
}

func (s *StartCloudRecordRequestSingleStreamingRecord) SetTranscodingParameters(v *StartCloudRecordRequestSingleStreamingRecordTranscodingParameters) *StartCloudRecordRequestSingleStreamingRecord {
	s.TranscodingParameters = v
	return s
}

func (s *StartCloudRecordRequestSingleStreamingRecord) Validate() error {
	return dara.Validate(s)
}

type StartCloudRecordRequestSingleStreamingRecordSpecifiedStreams struct {
	// This parameter is required.
	Ids []*string `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	// This parameter is required.
	StreamType *string `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	// example:
	//
	// white
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s StartCloudRecordRequestSingleStreamingRecordSpecifiedStreams) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordRequestSingleStreamingRecordSpecifiedStreams) GoString() string {
	return s.String()
}

func (s *StartCloudRecordRequestSingleStreamingRecordSpecifiedStreams) GetIds() []*string {
	return s.Ids
}

func (s *StartCloudRecordRequestSingleStreamingRecordSpecifiedStreams) GetStreamType() *string {
	return s.StreamType
}

func (s *StartCloudRecordRequestSingleStreamingRecordSpecifiedStreams) GetType() *string {
	return s.Type
}

func (s *StartCloudRecordRequestSingleStreamingRecordSpecifiedStreams) SetIds(v []*string) *StartCloudRecordRequestSingleStreamingRecordSpecifiedStreams {
	s.Ids = v
	return s
}

func (s *StartCloudRecordRequestSingleStreamingRecordSpecifiedStreams) SetStreamType(v string) *StartCloudRecordRequestSingleStreamingRecordSpecifiedStreams {
	s.StreamType = &v
	return s
}

func (s *StartCloudRecordRequestSingleStreamingRecordSpecifiedStreams) SetType(v string) *StartCloudRecordRequestSingleStreamingRecordSpecifiedStreams {
	s.Type = &v
	return s
}

func (s *StartCloudRecordRequestSingleStreamingRecordSpecifiedStreams) Validate() error {
	return dara.Validate(s)
}

type StartCloudRecordRequestSingleStreamingRecordTranscodingParameters struct {
	Audio     *StartCloudRecordRequestSingleStreamingRecordTranscodingParametersAudio `json:"Audio,omitempty" xml:"Audio,omitempty" type:"Struct"`
	Container *string                                                                 `json:"Container,omitempty" xml:"Container,omitempty"`
}

func (s StartCloudRecordRequestSingleStreamingRecordTranscodingParameters) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordRequestSingleStreamingRecordTranscodingParameters) GoString() string {
	return s.String()
}

func (s *StartCloudRecordRequestSingleStreamingRecordTranscodingParameters) GetAudio() *StartCloudRecordRequestSingleStreamingRecordTranscodingParametersAudio {
	return s.Audio
}

func (s *StartCloudRecordRequestSingleStreamingRecordTranscodingParameters) GetContainer() *string {
	return s.Container
}

func (s *StartCloudRecordRequestSingleStreamingRecordTranscodingParameters) SetAudio(v *StartCloudRecordRequestSingleStreamingRecordTranscodingParametersAudio) *StartCloudRecordRequestSingleStreamingRecordTranscodingParameters {
	s.Audio = v
	return s
}

func (s *StartCloudRecordRequestSingleStreamingRecordTranscodingParameters) SetContainer(v string) *StartCloudRecordRequestSingleStreamingRecordTranscodingParameters {
	s.Container = &v
	return s
}

func (s *StartCloudRecordRequestSingleStreamingRecordTranscodingParameters) Validate() error {
	return dara.Validate(s)
}

type StartCloudRecordRequestSingleStreamingRecordTranscodingParametersAudio struct {
	Bitrate    *int32  `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	Codec      *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	SampleRate *int32  `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
}

func (s StartCloudRecordRequestSingleStreamingRecordTranscodingParametersAudio) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordRequestSingleStreamingRecordTranscodingParametersAudio) GoString() string {
	return s.String()
}

func (s *StartCloudRecordRequestSingleStreamingRecordTranscodingParametersAudio) GetBitrate() *int32 {
	return s.Bitrate
}

func (s *StartCloudRecordRequestSingleStreamingRecordTranscodingParametersAudio) GetCodec() *string {
	return s.Codec
}

func (s *StartCloudRecordRequestSingleStreamingRecordTranscodingParametersAudio) GetSampleRate() *int32 {
	return s.SampleRate
}

func (s *StartCloudRecordRequestSingleStreamingRecordTranscodingParametersAudio) SetBitrate(v int32) *StartCloudRecordRequestSingleStreamingRecordTranscodingParametersAudio {
	s.Bitrate = &v
	return s
}

func (s *StartCloudRecordRequestSingleStreamingRecordTranscodingParametersAudio) SetCodec(v string) *StartCloudRecordRequestSingleStreamingRecordTranscodingParametersAudio {
	s.Codec = &v
	return s
}

func (s *StartCloudRecordRequestSingleStreamingRecordTranscodingParametersAudio) SetSampleRate(v int32) *StartCloudRecordRequestSingleStreamingRecordTranscodingParametersAudio {
	s.SampleRate = &v
	return s
}

func (s *StartCloudRecordRequestSingleStreamingRecordTranscodingParametersAudio) Validate() error {
	return dara.Validate(s)
}

type StartCloudRecordRequestStorageConfig struct {
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

func (s StartCloudRecordRequestStorageConfig) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordRequestStorageConfig) GoString() string {
	return s.String()
}

func (s *StartCloudRecordRequestStorageConfig) GetAccessKey() *string {
	return s.AccessKey
}

func (s *StartCloudRecordRequestStorageConfig) GetBucket() *string {
	return s.Bucket
}

func (s *StartCloudRecordRequestStorageConfig) GetEndpoint() *string {
	return s.Endpoint
}

func (s *StartCloudRecordRequestStorageConfig) GetRegion() *int32 {
	return s.Region
}

func (s *StartCloudRecordRequestStorageConfig) GetSecretKey() *string {
	return s.SecretKey
}

func (s *StartCloudRecordRequestStorageConfig) GetVendor() *int32 {
	return s.Vendor
}

func (s *StartCloudRecordRequestStorageConfig) SetAccessKey(v string) *StartCloudRecordRequestStorageConfig {
	s.AccessKey = &v
	return s
}

func (s *StartCloudRecordRequestStorageConfig) SetBucket(v string) *StartCloudRecordRequestStorageConfig {
	s.Bucket = &v
	return s
}

func (s *StartCloudRecordRequestStorageConfig) SetEndpoint(v string) *StartCloudRecordRequestStorageConfig {
	s.Endpoint = &v
	return s
}

func (s *StartCloudRecordRequestStorageConfig) SetRegion(v int32) *StartCloudRecordRequestStorageConfig {
	s.Region = &v
	return s
}

func (s *StartCloudRecordRequestStorageConfig) SetSecretKey(v string) *StartCloudRecordRequestStorageConfig {
	s.SecretKey = &v
	return s
}

func (s *StartCloudRecordRequestStorageConfig) SetVendor(v int32) *StartCloudRecordRequestStorageConfig {
	s.Vendor = &v
	return s
}

func (s *StartCloudRecordRequestStorageConfig) Validate() error {
	return dara.Validate(s)
}

type StartCloudRecordRequestTexts struct {
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
	BoxColor   *StartCloudRecordRequestTextsBoxColor `json:"BoxColor,omitempty" xml:"BoxColor,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Font      *int32                                 `json:"Font,omitempty" xml:"Font,omitempty"`
	FontColor *StartCloudRecordRequestTextsFontColor `json:"FontColor,omitempty" xml:"FontColor,omitempty" type:"Struct"`
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

func (s StartCloudRecordRequestTexts) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordRequestTexts) GoString() string {
	return s.String()
}

func (s *StartCloudRecordRequestTexts) GetAlpha() *float64 {
	return s.Alpha
}

func (s *StartCloudRecordRequestTexts) GetBoxAlpha() *float64 {
	return s.BoxAlpha
}

func (s *StartCloudRecordRequestTexts) GetBoxBorderw() *int32 {
	return s.BoxBorderw
}

func (s *StartCloudRecordRequestTexts) GetBoxColor() *StartCloudRecordRequestTextsBoxColor {
	return s.BoxColor
}

func (s *StartCloudRecordRequestTexts) GetFont() *int32 {
	return s.Font
}

func (s *StartCloudRecordRequestTexts) GetFontColor() *StartCloudRecordRequestTextsFontColor {
	return s.FontColor
}

func (s *StartCloudRecordRequestTexts) GetFontSize() *int32 {
	return s.FontSize
}

func (s *StartCloudRecordRequestTexts) GetHasBox() *bool {
	return s.HasBox
}

func (s *StartCloudRecordRequestTexts) GetLayer() *int32 {
	return s.Layer
}

func (s *StartCloudRecordRequestTexts) GetTexture() *string {
	return s.Texture
}

func (s *StartCloudRecordRequestTexts) GetX() *float64 {
	return s.X
}

func (s *StartCloudRecordRequestTexts) GetY() *float64 {
	return s.Y
}

func (s *StartCloudRecordRequestTexts) SetAlpha(v float64) *StartCloudRecordRequestTexts {
	s.Alpha = &v
	return s
}

func (s *StartCloudRecordRequestTexts) SetBoxAlpha(v float64) *StartCloudRecordRequestTexts {
	s.BoxAlpha = &v
	return s
}

func (s *StartCloudRecordRequestTexts) SetBoxBorderw(v int32) *StartCloudRecordRequestTexts {
	s.BoxBorderw = &v
	return s
}

func (s *StartCloudRecordRequestTexts) SetBoxColor(v *StartCloudRecordRequestTextsBoxColor) *StartCloudRecordRequestTexts {
	s.BoxColor = v
	return s
}

func (s *StartCloudRecordRequestTexts) SetFont(v int32) *StartCloudRecordRequestTexts {
	s.Font = &v
	return s
}

func (s *StartCloudRecordRequestTexts) SetFontColor(v *StartCloudRecordRequestTextsFontColor) *StartCloudRecordRequestTexts {
	s.FontColor = v
	return s
}

func (s *StartCloudRecordRequestTexts) SetFontSize(v int32) *StartCloudRecordRequestTexts {
	s.FontSize = &v
	return s
}

func (s *StartCloudRecordRequestTexts) SetHasBox(v bool) *StartCloudRecordRequestTexts {
	s.HasBox = &v
	return s
}

func (s *StartCloudRecordRequestTexts) SetLayer(v int32) *StartCloudRecordRequestTexts {
	s.Layer = &v
	return s
}

func (s *StartCloudRecordRequestTexts) SetTexture(v string) *StartCloudRecordRequestTexts {
	s.Texture = &v
	return s
}

func (s *StartCloudRecordRequestTexts) SetX(v float64) *StartCloudRecordRequestTexts {
	s.X = &v
	return s
}

func (s *StartCloudRecordRequestTexts) SetY(v float64) *StartCloudRecordRequestTexts {
	s.Y = &v
	return s
}

func (s *StartCloudRecordRequestTexts) Validate() error {
	return dara.Validate(s)
}

type StartCloudRecordRequestTextsBoxColor struct {
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

func (s StartCloudRecordRequestTextsBoxColor) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordRequestTextsBoxColor) GoString() string {
	return s.String()
}

func (s *StartCloudRecordRequestTextsBoxColor) GetB() *int32 {
	return s.B
}

func (s *StartCloudRecordRequestTextsBoxColor) GetG() *int32 {
	return s.G
}

func (s *StartCloudRecordRequestTextsBoxColor) GetR() *int32 {
	return s.R
}

func (s *StartCloudRecordRequestTextsBoxColor) SetB(v int32) *StartCloudRecordRequestTextsBoxColor {
	s.B = &v
	return s
}

func (s *StartCloudRecordRequestTextsBoxColor) SetG(v int32) *StartCloudRecordRequestTextsBoxColor {
	s.G = &v
	return s
}

func (s *StartCloudRecordRequestTextsBoxColor) SetR(v int32) *StartCloudRecordRequestTextsBoxColor {
	s.R = &v
	return s
}

func (s *StartCloudRecordRequestTextsBoxColor) Validate() error {
	return dara.Validate(s)
}

type StartCloudRecordRequestTextsFontColor struct {
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

func (s StartCloudRecordRequestTextsFontColor) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordRequestTextsFontColor) GoString() string {
	return s.String()
}

func (s *StartCloudRecordRequestTextsFontColor) GetB() *int32 {
	return s.B
}

func (s *StartCloudRecordRequestTextsFontColor) GetG() *int32 {
	return s.G
}

func (s *StartCloudRecordRequestTextsFontColor) GetR() *int32 {
	return s.R
}

func (s *StartCloudRecordRequestTextsFontColor) SetB(v int32) *StartCloudRecordRequestTextsFontColor {
	s.B = &v
	return s
}

func (s *StartCloudRecordRequestTextsFontColor) SetG(v int32) *StartCloudRecordRequestTextsFontColor {
	s.G = &v
	return s
}

func (s *StartCloudRecordRequestTextsFontColor) SetR(v int32) *StartCloudRecordRequestTextsFontColor {
	s.R = &v
	return s
}

func (s *StartCloudRecordRequestTextsFontColor) Validate() error {
	return dara.Validate(s)
}
