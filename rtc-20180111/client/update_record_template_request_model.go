// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecordTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateRecordTemplateRequest
	GetAppId() *string
	SetBackgroundColor(v int32) *UpdateRecordTemplateRequest
	GetBackgroundColor() *int32
	SetBackgrounds(v []*UpdateRecordTemplateRequestBackgrounds) *UpdateRecordTemplateRequest
	GetBackgrounds() []*UpdateRecordTemplateRequestBackgrounds
	SetClockWidgets(v []*UpdateRecordTemplateRequestClockWidgets) *UpdateRecordTemplateRequest
	GetClockWidgets() []*UpdateRecordTemplateRequestClockWidgets
	SetDelayStopTime(v int32) *UpdateRecordTemplateRequest
	GetDelayStopTime() *int32
	SetEnableM3u8DateTime(v bool) *UpdateRecordTemplateRequest
	GetEnableM3u8DateTime() *bool
	SetFileSplitInterval(v int32) *UpdateRecordTemplateRequest
	GetFileSplitInterval() *int32
	SetFormats(v []*string) *UpdateRecordTemplateRequest
	GetFormats() []*string
	SetHttpCallbackUrl(v string) *UpdateRecordTemplateRequest
	GetHttpCallbackUrl() *string
	SetLayoutIds(v []*int64) *UpdateRecordTemplateRequest
	GetLayoutIds() []*int64
	SetMediaEncode(v int32) *UpdateRecordTemplateRequest
	GetMediaEncode() *int32
	SetMnsQueue(v string) *UpdateRecordTemplateRequest
	GetMnsQueue() *string
	SetName(v string) *UpdateRecordTemplateRequest
	GetName() *string
	SetOssBucket(v string) *UpdateRecordTemplateRequest
	GetOssBucket() *string
	SetOssEndpoint(v string) *UpdateRecordTemplateRequest
	GetOssEndpoint() *string
	SetOssFilePrefix(v string) *UpdateRecordTemplateRequest
	GetOssFilePrefix() *string
	SetOwnerId(v int64) *UpdateRecordTemplateRequest
	GetOwnerId() *int64
	SetTaskProfile(v string) *UpdateRecordTemplateRequest
	GetTaskProfile() *string
	SetTemplateId(v string) *UpdateRecordTemplateRequest
	GetTemplateId() *string
	SetWatermarks(v []*UpdateRecordTemplateRequestWatermarks) *UpdateRecordTemplateRequest
	GetWatermarks() []*UpdateRecordTemplateRequestWatermarks
}

type UpdateRecordTemplateRequest struct {
	// The application ID. You can create and query the ID in the console. Only a single ID is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The background color in RGB format. Default value: **0*	- (black). The formula is `R + G × 256 + B × 65536`. Valid values for R (red), G (green), and B (blue): **0*	- to **255**.
	//
	// example:
	//
	// 0
	BackgroundColor *int32 `json:"BackgroundColor,omitempty" xml:"BackgroundColor,omitempty"`
	// The background image information.
	Backgrounds []*UpdateRecordTemplateRequestBackgrounds `json:"Backgrounds,omitempty" xml:"Backgrounds,omitempty" type:"Repeated"`
	// The clock information.
	ClockWidgets []*UpdateRecordTemplateRequestClockWidgets `json:"ClockWidgets,omitempty" xml:"ClockWidgets,omitempty" type:"Repeated"`
	// The delay before stopping the recording. Unit: seconds. Default value: **180**.
	//
	// example:
	//
	// 180
	DelayStopTime *int32 `json:"DelayStopTime,omitempty" xml:"DelayStopTime,omitempty"`
	// Specifies whether to add the timestamp of TS files to the recorded M3U8 file.
	//
	// example:
	//
	// false
	EnableM3u8DateTime *bool `json:"EnableM3u8DateTime,omitempty" xml:"EnableM3u8DateTime,omitempty"`
	// The interval at which recording files are split. The value must be greater than or equal to 1800. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1800
	FileSplitInterval *int32 `json:"FileSplitInterval,omitempty" xml:"FileSplitInterval,omitempty"`
	// The recording file formats.
	//
	// This parameter is required.
	//
	// example:
	//
	// mp4
	Formats []*string `json:"Formats,omitempty" xml:"Formats,omitempty" type:"Repeated"`
	// The HTTP callback URL for recording events.
	//
	// example:
	//
	// http://example.com/callback
	HttpCallbackUrl *string `json:"HttpCallbackUrl,omitempty" xml:"HttpCallbackUrl,omitempty"`
	// The layout IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1111
	LayoutIds []*int64 `json:"LayoutIds,omitempty" xml:"LayoutIds,omitempty" type:"Repeated"`
	// The encoding options.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	MediaEncode *int32 `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	// The MSMQ for recording event callbacks.
	//
	// example:
	//
	// record-callback-queue
	MnsQueue *string `json:"MnsQueue,omitempty" xml:"MnsQueue,omitempty"`
	// The name of the recording configuration template.
	//
	// This parameter is required.
	//
	// example:
	//
	// Recording template
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The OSS bucket where the recording files are stored.
	//
	// This parameter is required.
	//
	// example:
	//
	// rtc-record-pre
	OssBucket *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	// The OSS endpoint of the region where the recording files are stored.
	//
	// example:
	//
	// oss-cn-shanghai.aliyuncs.com
	OssEndpoint *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	// The naming rule for recording files.
	//
	// This parameter is required.
	//
	// example:
	//
	// record/{AppId}/{ChannelId_TaskId}/{EscapedStartTime}_{EscapedEndTime}
	OssFilePrefix *string `json:"OssFilePrefix,omitempty" xml:"OssFilePrefix,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The billing configuration of the task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4IN_1080P
	TaskProfile *string `json:"TaskProfile,omitempty" xml:"TaskProfile,omitempty"`
	// The configuration template ID. To obtain the template ID, call [AddRecordTemplate](https://help.aliyun.com/document_detail/193625.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 76dasgb****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The watermark information.
	Watermarks []*UpdateRecordTemplateRequestWatermarks `json:"Watermarks,omitempty" xml:"Watermarks,omitempty" type:"Repeated"`
}

func (s UpdateRecordTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecordTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateRecordTemplateRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateRecordTemplateRequest) GetBackgroundColor() *int32 {
	return s.BackgroundColor
}

func (s *UpdateRecordTemplateRequest) GetBackgrounds() []*UpdateRecordTemplateRequestBackgrounds {
	return s.Backgrounds
}

func (s *UpdateRecordTemplateRequest) GetClockWidgets() []*UpdateRecordTemplateRequestClockWidgets {
	return s.ClockWidgets
}

func (s *UpdateRecordTemplateRequest) GetDelayStopTime() *int32 {
	return s.DelayStopTime
}

func (s *UpdateRecordTemplateRequest) GetEnableM3u8DateTime() *bool {
	return s.EnableM3u8DateTime
}

func (s *UpdateRecordTemplateRequest) GetFileSplitInterval() *int32 {
	return s.FileSplitInterval
}

func (s *UpdateRecordTemplateRequest) GetFormats() []*string {
	return s.Formats
}

func (s *UpdateRecordTemplateRequest) GetHttpCallbackUrl() *string {
	return s.HttpCallbackUrl
}

func (s *UpdateRecordTemplateRequest) GetLayoutIds() []*int64 {
	return s.LayoutIds
}

func (s *UpdateRecordTemplateRequest) GetMediaEncode() *int32 {
	return s.MediaEncode
}

func (s *UpdateRecordTemplateRequest) GetMnsQueue() *string {
	return s.MnsQueue
}

func (s *UpdateRecordTemplateRequest) GetName() *string {
	return s.Name
}

func (s *UpdateRecordTemplateRequest) GetOssBucket() *string {
	return s.OssBucket
}

func (s *UpdateRecordTemplateRequest) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *UpdateRecordTemplateRequest) GetOssFilePrefix() *string {
	return s.OssFilePrefix
}

func (s *UpdateRecordTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateRecordTemplateRequest) GetTaskProfile() *string {
	return s.TaskProfile
}

func (s *UpdateRecordTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateRecordTemplateRequest) GetWatermarks() []*UpdateRecordTemplateRequestWatermarks {
	return s.Watermarks
}

func (s *UpdateRecordTemplateRequest) SetAppId(v string) *UpdateRecordTemplateRequest {
	s.AppId = &v
	return s
}

func (s *UpdateRecordTemplateRequest) SetBackgroundColor(v int32) *UpdateRecordTemplateRequest {
	s.BackgroundColor = &v
	return s
}

func (s *UpdateRecordTemplateRequest) SetBackgrounds(v []*UpdateRecordTemplateRequestBackgrounds) *UpdateRecordTemplateRequest {
	s.Backgrounds = v
	return s
}

func (s *UpdateRecordTemplateRequest) SetClockWidgets(v []*UpdateRecordTemplateRequestClockWidgets) *UpdateRecordTemplateRequest {
	s.ClockWidgets = v
	return s
}

func (s *UpdateRecordTemplateRequest) SetDelayStopTime(v int32) *UpdateRecordTemplateRequest {
	s.DelayStopTime = &v
	return s
}

func (s *UpdateRecordTemplateRequest) SetEnableM3u8DateTime(v bool) *UpdateRecordTemplateRequest {
	s.EnableM3u8DateTime = &v
	return s
}

func (s *UpdateRecordTemplateRequest) SetFileSplitInterval(v int32) *UpdateRecordTemplateRequest {
	s.FileSplitInterval = &v
	return s
}

func (s *UpdateRecordTemplateRequest) SetFormats(v []*string) *UpdateRecordTemplateRequest {
	s.Formats = v
	return s
}

func (s *UpdateRecordTemplateRequest) SetHttpCallbackUrl(v string) *UpdateRecordTemplateRequest {
	s.HttpCallbackUrl = &v
	return s
}

func (s *UpdateRecordTemplateRequest) SetLayoutIds(v []*int64) *UpdateRecordTemplateRequest {
	s.LayoutIds = v
	return s
}

func (s *UpdateRecordTemplateRequest) SetMediaEncode(v int32) *UpdateRecordTemplateRequest {
	s.MediaEncode = &v
	return s
}

func (s *UpdateRecordTemplateRequest) SetMnsQueue(v string) *UpdateRecordTemplateRequest {
	s.MnsQueue = &v
	return s
}

func (s *UpdateRecordTemplateRequest) SetName(v string) *UpdateRecordTemplateRequest {
	s.Name = &v
	return s
}

func (s *UpdateRecordTemplateRequest) SetOssBucket(v string) *UpdateRecordTemplateRequest {
	s.OssBucket = &v
	return s
}

func (s *UpdateRecordTemplateRequest) SetOssEndpoint(v string) *UpdateRecordTemplateRequest {
	s.OssEndpoint = &v
	return s
}

func (s *UpdateRecordTemplateRequest) SetOssFilePrefix(v string) *UpdateRecordTemplateRequest {
	s.OssFilePrefix = &v
	return s
}

func (s *UpdateRecordTemplateRequest) SetOwnerId(v int64) *UpdateRecordTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateRecordTemplateRequest) SetTaskProfile(v string) *UpdateRecordTemplateRequest {
	s.TaskProfile = &v
	return s
}

func (s *UpdateRecordTemplateRequest) SetTemplateId(v string) *UpdateRecordTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateRecordTemplateRequest) SetWatermarks(v []*UpdateRecordTemplateRequestWatermarks) *UpdateRecordTemplateRequest {
	s.Watermarks = v
	return s
}

func (s *UpdateRecordTemplateRequest) Validate() error {
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
	if s.Watermarks != nil {
		for _, item := range s.Watermarks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateRecordTemplateRequestBackgrounds struct {
	// Specifies whether to display the image. Valid values:
	//
	// - **0*	- (None): Not displayed.
	//
	// - **1*	- (Always): Always displayed.
	//
	// example:
	//
	// 0
	Display *int32 `json:"Display,omitempty" xml:"Display,omitempty"`
	// The height of the pane, as a normalized percentage.
	//
	// example:
	//
	// 0.2456
	Height *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The HTTP or HTTPS URL of the image.
	//
	// example:
	//
	// https://www.example.com/image.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The width of the pane, as a normalized percentage.
	//
	// example:
	//
	// 0.2456
	Width *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// The X-coordinate, as a normalized percentage.
	//
	// example:
	//
	// 0.7576
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// The Y-coordinate, as a normalized percentage.
	//
	// example:
	//
	// 0.7576
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	// The Z-order. A value of 0 indicates the bottom layer. A value of 1 indicates the layer above 0, and so on.
	//
	// example:
	//
	// 0
	ZOrder *int32 `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s UpdateRecordTemplateRequestBackgrounds) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecordTemplateRequestBackgrounds) GoString() string {
	return s.String()
}

func (s *UpdateRecordTemplateRequestBackgrounds) GetDisplay() *int32 {
	return s.Display
}

func (s *UpdateRecordTemplateRequestBackgrounds) GetHeight() *float32 {
	return s.Height
}

func (s *UpdateRecordTemplateRequestBackgrounds) GetUrl() *string {
	return s.Url
}

func (s *UpdateRecordTemplateRequestBackgrounds) GetWidth() *float32 {
	return s.Width
}

func (s *UpdateRecordTemplateRequestBackgrounds) GetX() *float32 {
	return s.X
}

func (s *UpdateRecordTemplateRequestBackgrounds) GetY() *float32 {
	return s.Y
}

func (s *UpdateRecordTemplateRequestBackgrounds) GetZOrder() *int32 {
	return s.ZOrder
}

func (s *UpdateRecordTemplateRequestBackgrounds) SetDisplay(v int32) *UpdateRecordTemplateRequestBackgrounds {
	s.Display = &v
	return s
}

func (s *UpdateRecordTemplateRequestBackgrounds) SetHeight(v float32) *UpdateRecordTemplateRequestBackgrounds {
	s.Height = &v
	return s
}

func (s *UpdateRecordTemplateRequestBackgrounds) SetUrl(v string) *UpdateRecordTemplateRequestBackgrounds {
	s.Url = &v
	return s
}

func (s *UpdateRecordTemplateRequestBackgrounds) SetWidth(v float32) *UpdateRecordTemplateRequestBackgrounds {
	s.Width = &v
	return s
}

func (s *UpdateRecordTemplateRequestBackgrounds) SetX(v float32) *UpdateRecordTemplateRequestBackgrounds {
	s.X = &v
	return s
}

func (s *UpdateRecordTemplateRequestBackgrounds) SetY(v float32) *UpdateRecordTemplateRequestBackgrounds {
	s.Y = &v
	return s
}

func (s *UpdateRecordTemplateRequestBackgrounds) SetZOrder(v int32) *UpdateRecordTemplateRequestBackgrounds {
	s.ZOrder = &v
	return s
}

func (s *UpdateRecordTemplateRequestBackgrounds) Validate() error {
	return dara.Validate(s)
}

type UpdateRecordTemplateRequestClockWidgets struct {
	// The text color of the clock in RGB format.
	//
	// The formula is `R + G × 256 + B × 65536`. Valid values for R (red), G (green), and B (blue): **0*	- to **255**.
	//
	// example:
	//
	// 0
	FontColor *int32 `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	// The font size of the clock. Valid values: **(0, 72]**.
	//
	// example:
	//
	// 1
	FontSize *int32 `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	// The font type of the clock. Valid values:
	//
	// - **0*	- (default): NOTO_SERIF_CJKSC_REGULAR
	//
	// - **1**: ALIBABA_PUHUITI_REGULAR
	//
	// - **2**: ALIBABA_PUHUITI_BOLD
	//
	// - **3**: ALIBABA_PUHUITI_Heavy
	//
	// - **4**: ALIBABA_PUHUITI_LIGHT
	//
	// - **5**: ALIBABA_PUHUITI_MEDIUM
	//
	// example:
	//
	// 0
	FontType *int32 `json:"FontType,omitempty" xml:"FontType,omitempty"`
	// The X-coordinate of the clock, as a normalized percentage.
	//
	// example:
	//
	// 0.7576
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// The Y-coordinate of the clock, as a normalized percentage.
	//
	// example:
	//
	// 0.7576
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	// The Z-order of the clock. A value of 0 indicates the bottom layer. A value of 1 indicates the layer above 0, and so on.
	//
	// example:
	//
	// 0
	ZOrder *int32 `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s UpdateRecordTemplateRequestClockWidgets) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecordTemplateRequestClockWidgets) GoString() string {
	return s.String()
}

func (s *UpdateRecordTemplateRequestClockWidgets) GetFontColor() *int32 {
	return s.FontColor
}

func (s *UpdateRecordTemplateRequestClockWidgets) GetFontSize() *int32 {
	return s.FontSize
}

func (s *UpdateRecordTemplateRequestClockWidgets) GetFontType() *int32 {
	return s.FontType
}

func (s *UpdateRecordTemplateRequestClockWidgets) GetX() *float32 {
	return s.X
}

func (s *UpdateRecordTemplateRequestClockWidgets) GetY() *float32 {
	return s.Y
}

func (s *UpdateRecordTemplateRequestClockWidgets) GetZOrder() *int32 {
	return s.ZOrder
}

func (s *UpdateRecordTemplateRequestClockWidgets) SetFontColor(v int32) *UpdateRecordTemplateRequestClockWidgets {
	s.FontColor = &v
	return s
}

func (s *UpdateRecordTemplateRequestClockWidgets) SetFontSize(v int32) *UpdateRecordTemplateRequestClockWidgets {
	s.FontSize = &v
	return s
}

func (s *UpdateRecordTemplateRequestClockWidgets) SetFontType(v int32) *UpdateRecordTemplateRequestClockWidgets {
	s.FontType = &v
	return s
}

func (s *UpdateRecordTemplateRequestClockWidgets) SetX(v float32) *UpdateRecordTemplateRequestClockWidgets {
	s.X = &v
	return s
}

func (s *UpdateRecordTemplateRequestClockWidgets) SetY(v float32) *UpdateRecordTemplateRequestClockWidgets {
	s.Y = &v
	return s
}

func (s *UpdateRecordTemplateRequestClockWidgets) SetZOrder(v int32) *UpdateRecordTemplateRequestClockWidgets {
	s.ZOrder = &v
	return s
}

func (s *UpdateRecordTemplateRequestClockWidgets) Validate() error {
	return dara.Validate(s)
}

type UpdateRecordTemplateRequestWatermarks struct {
	// The transparency. A value of 0.0 indicates fully transparent, and a value of 1.0 indicates fully opaque.
	//
	// example:
	//
	// 0
	Alpha *float32 `json:"Alpha,omitempty" xml:"Alpha,omitempty"`
	// Specifies whether to display the watermark. Valid values:
	//
	// - **0*	- (None): Not displayed.
	//
	// - **1*	- (Always): Always displayed.
	//
	// example:
	//
	// 0
	Display *int32 `json:"Display,omitempty" xml:"Display,omitempty"`
	// The height of the pane, as a normalized percentage.
	//
	// example:
	//
	// 0.2456
	Height *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The HTTP or HTTPS URL of the watermark.
	//
	// example:
	//
	// https://www.example.com/image.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The width of the pane, as a normalized percentage.
	//
	// example:
	//
	// 0.2456
	Width *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// The X-coordinate, as a normalized percentage.
	//
	// example:
	//
	// 0.7576
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// The Y-coordinate, as a normalized percentage.
	//
	// example:
	//
	// 0.7576
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	// The Z-order. A value of 0 indicates the bottom layer. A value of 1 indicates the layer above 0, and so on.
	//
	// example:
	//
	// 0
	ZOrder *int32 `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s UpdateRecordTemplateRequestWatermarks) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecordTemplateRequestWatermarks) GoString() string {
	return s.String()
}

func (s *UpdateRecordTemplateRequestWatermarks) GetAlpha() *float32 {
	return s.Alpha
}

func (s *UpdateRecordTemplateRequestWatermarks) GetDisplay() *int32 {
	return s.Display
}

func (s *UpdateRecordTemplateRequestWatermarks) GetHeight() *float32 {
	return s.Height
}

func (s *UpdateRecordTemplateRequestWatermarks) GetUrl() *string {
	return s.Url
}

func (s *UpdateRecordTemplateRequestWatermarks) GetWidth() *float32 {
	return s.Width
}

func (s *UpdateRecordTemplateRequestWatermarks) GetX() *float32 {
	return s.X
}

func (s *UpdateRecordTemplateRequestWatermarks) GetY() *float32 {
	return s.Y
}

func (s *UpdateRecordTemplateRequestWatermarks) GetZOrder() *int32 {
	return s.ZOrder
}

func (s *UpdateRecordTemplateRequestWatermarks) SetAlpha(v float32) *UpdateRecordTemplateRequestWatermarks {
	s.Alpha = &v
	return s
}

func (s *UpdateRecordTemplateRequestWatermarks) SetDisplay(v int32) *UpdateRecordTemplateRequestWatermarks {
	s.Display = &v
	return s
}

func (s *UpdateRecordTemplateRequestWatermarks) SetHeight(v float32) *UpdateRecordTemplateRequestWatermarks {
	s.Height = &v
	return s
}

func (s *UpdateRecordTemplateRequestWatermarks) SetUrl(v string) *UpdateRecordTemplateRequestWatermarks {
	s.Url = &v
	return s
}

func (s *UpdateRecordTemplateRequestWatermarks) SetWidth(v float32) *UpdateRecordTemplateRequestWatermarks {
	s.Width = &v
	return s
}

func (s *UpdateRecordTemplateRequestWatermarks) SetX(v float32) *UpdateRecordTemplateRequestWatermarks {
	s.X = &v
	return s
}

func (s *UpdateRecordTemplateRequestWatermarks) SetY(v float32) *UpdateRecordTemplateRequestWatermarks {
	s.Y = &v
	return s
}

func (s *UpdateRecordTemplateRequestWatermarks) SetZOrder(v int32) *UpdateRecordTemplateRequestWatermarks {
	s.ZOrder = &v
	return s
}

func (s *UpdateRecordTemplateRequestWatermarks) Validate() error {
	return dara.Validate(s)
}
