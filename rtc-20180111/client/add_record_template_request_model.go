// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRecordTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *AddRecordTemplateRequest
	GetAppId() *string
	SetBackgroundColor(v int32) *AddRecordTemplateRequest
	GetBackgroundColor() *int32
	SetBackgrounds(v []*AddRecordTemplateRequestBackgrounds) *AddRecordTemplateRequest
	GetBackgrounds() []*AddRecordTemplateRequestBackgrounds
	SetClockWidgets(v []*AddRecordTemplateRequestClockWidgets) *AddRecordTemplateRequest
	GetClockWidgets() []*AddRecordTemplateRequestClockWidgets
	SetDelayStopTime(v int32) *AddRecordTemplateRequest
	GetDelayStopTime() *int32
	SetEnableM3u8DateTime(v bool) *AddRecordTemplateRequest
	GetEnableM3u8DateTime() *bool
	SetFileSplitInterval(v int32) *AddRecordTemplateRequest
	GetFileSplitInterval() *int32
	SetFormats(v []*string) *AddRecordTemplateRequest
	GetFormats() []*string
	SetHttpCallbackUrl(v string) *AddRecordTemplateRequest
	GetHttpCallbackUrl() *string
	SetLayoutIds(v []*int64) *AddRecordTemplateRequest
	GetLayoutIds() []*int64
	SetMediaEncode(v int32) *AddRecordTemplateRequest
	GetMediaEncode() *int32
	SetMnsQueue(v string) *AddRecordTemplateRequest
	GetMnsQueue() *string
	SetName(v string) *AddRecordTemplateRequest
	GetName() *string
	SetOssBucket(v string) *AddRecordTemplateRequest
	GetOssBucket() *string
	SetOssEndpoint(v string) *AddRecordTemplateRequest
	GetOssEndpoint() *string
	SetOssFilePrefix(v string) *AddRecordTemplateRequest
	GetOssFilePrefix() *string
	SetOwnerId(v int64) *AddRecordTemplateRequest
	GetOwnerId() *int64
	SetTaskProfile(v string) *AddRecordTemplateRequest
	GetTaskProfile() *string
	SetWatermarks(v []*AddRecordTemplateRequestWatermarks) *AddRecordTemplateRequest
	GetWatermarks() []*AddRecordTemplateRequestWatermarks
}

type AddRecordTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 0
	BackgroundColor *int32                                  `json:"BackgroundColor,omitempty" xml:"BackgroundColor,omitempty"`
	Backgrounds     []*AddRecordTemplateRequestBackgrounds  `json:"Backgrounds,omitempty" xml:"Backgrounds,omitempty" type:"Repeated"`
	ClockWidgets    []*AddRecordTemplateRequestClockWidgets `json:"ClockWidgets,omitempty" xml:"ClockWidgets,omitempty" type:"Repeated"`
	// example:
	//
	// 180
	DelayStopTime *int32 `json:"DelayStopTime,omitempty" xml:"DelayStopTime,omitempty"`
	// example:
	//
	// false
	EnableM3u8DateTime *bool `json:"EnableM3u8DateTime,omitempty" xml:"EnableM3u8DateTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1800
	FileSplitInterval *int32 `json:"FileSplitInterval,omitempty" xml:"FileSplitInterval,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mp4
	Formats []*string `json:"Formats,omitempty" xml:"Formats,omitempty" type:"Repeated"`
	// example:
	//
	// http://example.com/callback
	HttpCallbackUrl *string `json:"HttpCallbackUrl,omitempty" xml:"HttpCallbackUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	LayoutIds []*int64 `json:"LayoutIds,omitempty" xml:"LayoutIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	MediaEncode *int32 `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	// example:
	//
	// record-callback-queue
	MnsQueue *string `json:"MnsQueue,omitempty" xml:"MnsQueue,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rtc-record-oss
	OssBucket   *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// record/{AppId}/{ChannelId_TaskId}/{EscapedStartTime}_{EscapedEndTime}
	OssFilePrefix *string `json:"OssFilePrefix,omitempty" xml:"OssFilePrefix,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4IN_1080P
	TaskProfile *string                               `json:"TaskProfile,omitempty" xml:"TaskProfile,omitempty"`
	Watermarks  []*AddRecordTemplateRequestWatermarks `json:"Watermarks,omitempty" xml:"Watermarks,omitempty" type:"Repeated"`
}

func (s AddRecordTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s AddRecordTemplateRequest) GoString() string {
	return s.String()
}

func (s *AddRecordTemplateRequest) GetAppId() *string {
	return s.AppId
}

func (s *AddRecordTemplateRequest) GetBackgroundColor() *int32 {
	return s.BackgroundColor
}

func (s *AddRecordTemplateRequest) GetBackgrounds() []*AddRecordTemplateRequestBackgrounds {
	return s.Backgrounds
}

func (s *AddRecordTemplateRequest) GetClockWidgets() []*AddRecordTemplateRequestClockWidgets {
	return s.ClockWidgets
}

func (s *AddRecordTemplateRequest) GetDelayStopTime() *int32 {
	return s.DelayStopTime
}

func (s *AddRecordTemplateRequest) GetEnableM3u8DateTime() *bool {
	return s.EnableM3u8DateTime
}

func (s *AddRecordTemplateRequest) GetFileSplitInterval() *int32 {
	return s.FileSplitInterval
}

func (s *AddRecordTemplateRequest) GetFormats() []*string {
	return s.Formats
}

func (s *AddRecordTemplateRequest) GetHttpCallbackUrl() *string {
	return s.HttpCallbackUrl
}

func (s *AddRecordTemplateRequest) GetLayoutIds() []*int64 {
	return s.LayoutIds
}

func (s *AddRecordTemplateRequest) GetMediaEncode() *int32 {
	return s.MediaEncode
}

func (s *AddRecordTemplateRequest) GetMnsQueue() *string {
	return s.MnsQueue
}

func (s *AddRecordTemplateRequest) GetName() *string {
	return s.Name
}

func (s *AddRecordTemplateRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *AddRecordTemplateRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *AddRecordTemplateRequest) GetOssFilePrefix() *string {
	return s.OssFilePrefix
}

func (s *AddRecordTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddRecordTemplateRequest) GetTaskProfile() *string {
	return s.TaskProfile
}

func (s *AddRecordTemplateRequest) GetWatermarks() []*AddRecordTemplateRequestWatermarks {
	return s.Watermarks
}

func (s *AddRecordTemplateRequest) SetAppId(v string) *AddRecordTemplateRequest {
	s.AppId = &v
	return s
}

func (s *AddRecordTemplateRequest) SetBackgroundColor(v int32) *AddRecordTemplateRequest {
	s.BackgroundColor = &v
	return s
}

func (s *AddRecordTemplateRequest) SetBackgrounds(v []*AddRecordTemplateRequestBackgrounds) *AddRecordTemplateRequest {
	s.Backgrounds = v
	return s
}

func (s *AddRecordTemplateRequest) SetClockWidgets(v []*AddRecordTemplateRequestClockWidgets) *AddRecordTemplateRequest {
	s.ClockWidgets = v
	return s
}

func (s *AddRecordTemplateRequest) SetDelayStopTime(v int32) *AddRecordTemplateRequest {
	s.DelayStopTime = &v
	return s
}

func (s *AddRecordTemplateRequest) SetEnableM3u8DateTime(v bool) *AddRecordTemplateRequest {
	s.EnableM3u8DateTime = &v
	return s
}

func (s *AddRecordTemplateRequest) SetFileSplitInterval(v int32) *AddRecordTemplateRequest {
	s.FileSplitInterval = &v
	return s
}

func (s *AddRecordTemplateRequest) SetFormats(v []*string) *AddRecordTemplateRequest {
	s.Formats = v
	return s
}

func (s *AddRecordTemplateRequest) SetHttpCallbackUrl(v string) *AddRecordTemplateRequest {
	s.HttpCallbackUrl = &v
	return s
}

func (s *AddRecordTemplateRequest) SetLayoutIds(v []*int64) *AddRecordTemplateRequest {
	s.LayoutIds = v
	return s
}

func (s *AddRecordTemplateRequest) SetMediaEncode(v int32) *AddRecordTemplateRequest {
	s.MediaEncode = &v
	return s
}

func (s *AddRecordTemplateRequest) SetMnsQueue(v string) *AddRecordTemplateRequest {
	s.MnsQueue = &v
	return s
}

func (s *AddRecordTemplateRequest) SetName(v string) *AddRecordTemplateRequest {
	s.Name = &v
	return s
}

func (s *AddRecordTemplateRequest) SetOssBucket(v string) *AddRecordTemplateRequest {
	s.OssBucket = &v
	return s
}

func (s *AddRecordTemplateRequest) SetOssEndpoint(v string) *AddRecordTemplateRequest {
	s.OssEndpoint = &v
	return s
}

func (s *AddRecordTemplateRequest) SetOssFilePrefix(v string) *AddRecordTemplateRequest {
	s.OssFilePrefix = &v
	return s
}

func (s *AddRecordTemplateRequest) SetOwnerId(v int64) *AddRecordTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *AddRecordTemplateRequest) SetTaskProfile(v string) *AddRecordTemplateRequest {
	s.TaskProfile = &v
	return s
}

func (s *AddRecordTemplateRequest) SetWatermarks(v []*AddRecordTemplateRequestWatermarks) *AddRecordTemplateRequest {
	s.Watermarks = v
	return s
}

func (s *AddRecordTemplateRequest) Validate() error {
	return dara.Validate(s)
}

type AddRecordTemplateRequestBackgrounds struct {
	// example:
	//
	// 0
	Display *int32 `json:"Display,omitempty" xml:"Display,omitempty"`
	// example:
	//
	// 0.2456
	Height *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// https://www.example.com/image.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// example:
	//
	// 0.2456
	Width *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// example:
	//
	// 0.7576
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 0.7576
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	// example:
	//
	// 0
	ZOrder *int32 `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s AddRecordTemplateRequestBackgrounds) String() string {
	return dara.Prettify(s)
}

func (s AddRecordTemplateRequestBackgrounds) GoString() string {
	return s.String()
}

func (s *AddRecordTemplateRequestBackgrounds) GetDisplay() *int32 {
	return s.Display
}

func (s *AddRecordTemplateRequestBackgrounds) GetHeight() *float32 {
	return s.Height
}

func (s *AddRecordTemplateRequestBackgrounds) GetUrl() *string {
	return s.Url
}

func (s *AddRecordTemplateRequestBackgrounds) GetWidth() *float32 {
	return s.Width
}

func (s *AddRecordTemplateRequestBackgrounds) GetX() *float32 {
	return s.X
}

func (s *AddRecordTemplateRequestBackgrounds) GetY() *float32 {
	return s.Y
}

func (s *AddRecordTemplateRequestBackgrounds) GetZOrder() *int32 {
	return s.ZOrder
}

func (s *AddRecordTemplateRequestBackgrounds) SetDisplay(v int32) *AddRecordTemplateRequestBackgrounds {
	s.Display = &v
	return s
}

func (s *AddRecordTemplateRequestBackgrounds) SetHeight(v float32) *AddRecordTemplateRequestBackgrounds {
	s.Height = &v
	return s
}

func (s *AddRecordTemplateRequestBackgrounds) SetUrl(v string) *AddRecordTemplateRequestBackgrounds {
	s.Url = &v
	return s
}

func (s *AddRecordTemplateRequestBackgrounds) SetWidth(v float32) *AddRecordTemplateRequestBackgrounds {
	s.Width = &v
	return s
}

func (s *AddRecordTemplateRequestBackgrounds) SetX(v float32) *AddRecordTemplateRequestBackgrounds {
	s.X = &v
	return s
}

func (s *AddRecordTemplateRequestBackgrounds) SetY(v float32) *AddRecordTemplateRequestBackgrounds {
	s.Y = &v
	return s
}

func (s *AddRecordTemplateRequestBackgrounds) SetZOrder(v int32) *AddRecordTemplateRequestBackgrounds {
	s.ZOrder = &v
	return s
}

func (s *AddRecordTemplateRequestBackgrounds) Validate() error {
	return dara.Validate(s)
}

type AddRecordTemplateRequestClockWidgets struct {
	// example:
	//
	// 0
	FontColor *int32 `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	// example:
	//
	// 1
	FontSize *int32 `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	// example:
	//
	// 0
	FontType *int32 `json:"FontType,omitempty" xml:"FontType,omitempty"`
	// example:
	//
	// 0.7576
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 0.7576
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	// example:
	//
	// 0
	ZOrder *int32 `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s AddRecordTemplateRequestClockWidgets) String() string {
	return dara.Prettify(s)
}

func (s AddRecordTemplateRequestClockWidgets) GoString() string {
	return s.String()
}

func (s *AddRecordTemplateRequestClockWidgets) GetFontColor() *int32 {
	return s.FontColor
}

func (s *AddRecordTemplateRequestClockWidgets) GetFontSize() *int32 {
	return s.FontSize
}

func (s *AddRecordTemplateRequestClockWidgets) GetFontType() *int32 {
	return s.FontType
}

func (s *AddRecordTemplateRequestClockWidgets) GetX() *float32 {
	return s.X
}

func (s *AddRecordTemplateRequestClockWidgets) GetY() *float32 {
	return s.Y
}

func (s *AddRecordTemplateRequestClockWidgets) GetZOrder() *int32 {
	return s.ZOrder
}

func (s *AddRecordTemplateRequestClockWidgets) SetFontColor(v int32) *AddRecordTemplateRequestClockWidgets {
	s.FontColor = &v
	return s
}

func (s *AddRecordTemplateRequestClockWidgets) SetFontSize(v int32) *AddRecordTemplateRequestClockWidgets {
	s.FontSize = &v
	return s
}

func (s *AddRecordTemplateRequestClockWidgets) SetFontType(v int32) *AddRecordTemplateRequestClockWidgets {
	s.FontType = &v
	return s
}

func (s *AddRecordTemplateRequestClockWidgets) SetX(v float32) *AddRecordTemplateRequestClockWidgets {
	s.X = &v
	return s
}

func (s *AddRecordTemplateRequestClockWidgets) SetY(v float32) *AddRecordTemplateRequestClockWidgets {
	s.Y = &v
	return s
}

func (s *AddRecordTemplateRequestClockWidgets) SetZOrder(v int32) *AddRecordTemplateRequestClockWidgets {
	s.ZOrder = &v
	return s
}

func (s *AddRecordTemplateRequestClockWidgets) Validate() error {
	return dara.Validate(s)
}

type AddRecordTemplateRequestWatermarks struct {
	// example:
	//
	// 0
	Alpha *float32 `json:"Alpha,omitempty" xml:"Alpha,omitempty"`
	// example:
	//
	// 0
	Display *int32 `json:"Display,omitempty" xml:"Display,omitempty"`
	// example:
	//
	// 0.2456
	Height *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// https://www.example.com/image.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// example:
	//
	// 0.2456
	Width *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// example:
	//
	// 0.7576
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 0.7576
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	// example:
	//
	// 0
	ZOrder *int32 `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s AddRecordTemplateRequestWatermarks) String() string {
	return dara.Prettify(s)
}

func (s AddRecordTemplateRequestWatermarks) GoString() string {
	return s.String()
}

func (s *AddRecordTemplateRequestWatermarks) GetAlpha() *float32 {
	return s.Alpha
}

func (s *AddRecordTemplateRequestWatermarks) GetDisplay() *int32 {
	return s.Display
}

func (s *AddRecordTemplateRequestWatermarks) GetHeight() *float32 {
	return s.Height
}

func (s *AddRecordTemplateRequestWatermarks) GetUrl() *string {
	return s.Url
}

func (s *AddRecordTemplateRequestWatermarks) GetWidth() *float32 {
	return s.Width
}

func (s *AddRecordTemplateRequestWatermarks) GetX() *float32 {
	return s.X
}

func (s *AddRecordTemplateRequestWatermarks) GetY() *float32 {
	return s.Y
}

func (s *AddRecordTemplateRequestWatermarks) GetZOrder() *int32 {
	return s.ZOrder
}

func (s *AddRecordTemplateRequestWatermarks) SetAlpha(v float32) *AddRecordTemplateRequestWatermarks {
	s.Alpha = &v
	return s
}

func (s *AddRecordTemplateRequestWatermarks) SetDisplay(v int32) *AddRecordTemplateRequestWatermarks {
	s.Display = &v
	return s
}

func (s *AddRecordTemplateRequestWatermarks) SetHeight(v float32) *AddRecordTemplateRequestWatermarks {
	s.Height = &v
	return s
}

func (s *AddRecordTemplateRequestWatermarks) SetUrl(v string) *AddRecordTemplateRequestWatermarks {
	s.Url = &v
	return s
}

func (s *AddRecordTemplateRequestWatermarks) SetWidth(v float32) *AddRecordTemplateRequestWatermarks {
	s.Width = &v
	return s
}

func (s *AddRecordTemplateRequestWatermarks) SetX(v float32) *AddRecordTemplateRequestWatermarks {
	s.X = &v
	return s
}

func (s *AddRecordTemplateRequestWatermarks) SetY(v float32) *AddRecordTemplateRequestWatermarks {
	s.Y = &v
	return s
}

func (s *AddRecordTemplateRequestWatermarks) SetZOrder(v int32) *AddRecordTemplateRequestWatermarks {
	s.ZOrder = &v
	return s
}

func (s *AddRecordTemplateRequestWatermarks) Validate() error {
	return dara.Validate(s)
}
