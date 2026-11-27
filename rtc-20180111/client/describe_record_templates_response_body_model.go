// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecordTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRecordTemplatesResponseBody
	GetRequestId() *string
	SetTemplates(v []*DescribeRecordTemplatesResponseBodyTemplates) *DescribeRecordTemplatesResponseBody
	GetTemplates() []*DescribeRecordTemplatesResponseBodyTemplates
	SetTotalNum(v int64) *DescribeRecordTemplatesResponseBody
	GetTotalNum() *int64
	SetTotalPage(v int64) *DescribeRecordTemplatesResponseBody
	GetTotalPage() *int64
}

type DescribeRecordTemplatesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C292B80E-5175-4BA4-8CC292B80E-5175-4BA4-8C1E-2ABEC4D7C2FE1E-2ABEC4D7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The recording configuration templates.
	Templates []*DescribeRecordTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 1
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeRecordTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecordTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRecordTemplatesResponseBody) GetTemplates() []*DescribeRecordTemplatesResponseBodyTemplates {
	return s.Templates
}

func (s *DescribeRecordTemplatesResponseBody) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *DescribeRecordTemplatesResponseBody) GetTotalPage() *int64 {
	return s.TotalPage
}

func (s *DescribeRecordTemplatesResponseBody) SetRequestId(v string) *DescribeRecordTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBody) SetTemplates(v []*DescribeRecordTemplatesResponseBodyTemplates) *DescribeRecordTemplatesResponseBody {
	s.Templates = v
	return s
}

func (s *DescribeRecordTemplatesResponseBody) SetTotalNum(v int64) *DescribeRecordTemplatesResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBody) SetTotalPage(v int64) *DescribeRecordTemplatesResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBody) Validate() error {
	if s.Templates != nil {
		for _, item := range s.Templates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRecordTemplatesResponseBodyTemplates struct {
	// The RGB value of the background color. Default value: **0*	- (black). The calculation formula is R + G × 256 + B × 65536. Valid values for R (red), G (green), and B (blue): **0*	- to **255**.
	//
	// example:
	//
	// 0
	BackgroundColor *int32 `json:"BackgroundColor,omitempty" xml:"BackgroundColor,omitempty"`
	// The list of background image information.
	Backgrounds []*DescribeRecordTemplatesResponseBodyTemplatesBackgrounds `json:"Backgrounds,omitempty" xml:"Backgrounds,omitempty" type:"Repeated"`
	// The list of clock widgets.
	ClockWidgets []*DescribeRecordTemplatesResponseBodyTemplatesClockWidgets `json:"ClockWidgets,omitempty" xml:"ClockWidgets,omitempty" type:"Repeated"`
	// The creation time. The time is displayed in the <i>yyyy-MM-dd</i>T<i>HH:mm:ss</i>Z format (UTC time).
	//
	// example:
	//
	// 2020-09-04T06:22:15Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The delay before stopping the recording. Unit: seconds.
	//
	// example:
	//
	// 180
	DelayStopTime *int32 `json:"DelayStopTime,omitempty" xml:"DelayStopTime,omitempty"`
	// Specifies whether to add the timestamp of the TS file to the recorded M3U8 file.
	//
	// example:
	//
	// false
	EnableM3u8DateTime *bool `json:"EnableM3u8DateTime,omitempty" xml:"EnableM3u8DateTime,omitempty"`
	// The file splitting duration.
	//
	// example:
	//
	// 1800
	FileSplitInterval *int32 `json:"FileSplitInterval,omitempty" xml:"FileSplitInterval,omitempty"`
	// The list of recording file formats. Valid formats include M3U8, FLV, and MP4.
	Formats []*string `json:"Formats,omitempty" xml:"Formats,omitempty" type:"Repeated"`
	// The HTTP callback URL for recording events.
	//
	// example:
	//
	// http://example.com/callback
	HttpCallbackUrl *string `json:"HttpCallbackUrl,omitempty" xml:"HttpCallbackUrl,omitempty"`
	// The list of mixed-stream layout IDs used by the recording template.
	LayoutIds []*int64 `json:"LayoutIds,omitempty" xml:"LayoutIds,omitempty" type:"Repeated"`
	// The encoding options.
	//
	// example:
	//
	// 50
	MediaEncode *int32 `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	// The name of the Message Service (MNS) queue for recording event callbacks.
	//
	// example:
	//
	// record-callback-queue
	MnsQueue *string `json:"MnsQueue,omitempty" xml:"MnsQueue,omitempty"`
	// The name of the recording configuration template.
	//
	// example:
	//
	// Template-demo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the Object Storage Service (OSS) bucket where recording files are stored.
	//
	// example:
	//
	// rtc-record-oss
	OssBucket *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	// The naming rule for recording files.
	//
	// example:
	//
	// record/pre/{AppId}/{ChannelId_TaskId}/{EscapedStartTime}_{EscapedEndTime}
	OssFilePrefix *string `json:"OssFilePrefix,omitempty" xml:"OssFilePrefix,omitempty"`
	// The task specifications.
	//
	// example:
	//
	// 4IN_1080P
	TaskProfile *string `json:"TaskProfile,omitempty" xml:"TaskProfile,omitempty"`
	// The ID of the recording configuration template.
	//
	// example:
	//
	// 1ca698e2-57fa-4314-8e11-00d950d4****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The list of watermark information.
	Watermarks []*DescribeRecordTemplatesResponseBodyTemplatesWatermarks `json:"Watermarks,omitempty" xml:"Watermarks,omitempty" type:"Repeated"`
}

func (s DescribeRecordTemplatesResponseBodyTemplates) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) GetBackgroundColor() *int32 {
	return s.BackgroundColor
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) GetBackgrounds() []*DescribeRecordTemplatesResponseBodyTemplatesBackgrounds {
	return s.Backgrounds
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) GetClockWidgets() []*DescribeRecordTemplatesResponseBodyTemplatesClockWidgets {
	return s.ClockWidgets
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) GetDelayStopTime() *int32 {
	return s.DelayStopTime
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) GetEnableM3u8DateTime() *bool {
	return s.EnableM3u8DateTime
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) GetFileSplitInterval() *int32 {
	return s.FileSplitInterval
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) GetFormats() []*string {
	return s.Formats
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) GetHttpCallbackUrl() *string {
	return s.HttpCallbackUrl
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) GetLayoutIds() []*int64 {
	return s.LayoutIds
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) GetMediaEncode() *int32 {
	return s.MediaEncode
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) GetMnsQueue() *string {
	return s.MnsQueue
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) GetName() *string {
	return s.Name
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) GetOssBucket() *string {
	return s.OssBucket
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) GetOssFilePrefix() *string {
	return s.OssFilePrefix
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) GetTaskProfile() *string {
	return s.TaskProfile
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) GetWatermarks() []*DescribeRecordTemplatesResponseBodyTemplatesWatermarks {
	return s.Watermarks
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) SetBackgroundColor(v int32) *DescribeRecordTemplatesResponseBodyTemplates {
	s.BackgroundColor = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) SetBackgrounds(v []*DescribeRecordTemplatesResponseBodyTemplatesBackgrounds) *DescribeRecordTemplatesResponseBodyTemplates {
	s.Backgrounds = v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) SetClockWidgets(v []*DescribeRecordTemplatesResponseBodyTemplatesClockWidgets) *DescribeRecordTemplatesResponseBodyTemplates {
	s.ClockWidgets = v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) SetCreateTime(v string) *DescribeRecordTemplatesResponseBodyTemplates {
	s.CreateTime = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) SetDelayStopTime(v int32) *DescribeRecordTemplatesResponseBodyTemplates {
	s.DelayStopTime = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) SetEnableM3u8DateTime(v bool) *DescribeRecordTemplatesResponseBodyTemplates {
	s.EnableM3u8DateTime = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) SetFileSplitInterval(v int32) *DescribeRecordTemplatesResponseBodyTemplates {
	s.FileSplitInterval = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) SetFormats(v []*string) *DescribeRecordTemplatesResponseBodyTemplates {
	s.Formats = v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) SetHttpCallbackUrl(v string) *DescribeRecordTemplatesResponseBodyTemplates {
	s.HttpCallbackUrl = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) SetLayoutIds(v []*int64) *DescribeRecordTemplatesResponseBodyTemplates {
	s.LayoutIds = v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) SetMediaEncode(v int32) *DescribeRecordTemplatesResponseBodyTemplates {
	s.MediaEncode = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) SetMnsQueue(v string) *DescribeRecordTemplatesResponseBodyTemplates {
	s.MnsQueue = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) SetName(v string) *DescribeRecordTemplatesResponseBodyTemplates {
	s.Name = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) SetOssBucket(v string) *DescribeRecordTemplatesResponseBodyTemplates {
	s.OssBucket = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) SetOssFilePrefix(v string) *DescribeRecordTemplatesResponseBodyTemplates {
	s.OssFilePrefix = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) SetTaskProfile(v string) *DescribeRecordTemplatesResponseBodyTemplates {
	s.TaskProfile = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) SetTemplateId(v string) *DescribeRecordTemplatesResponseBodyTemplates {
	s.TemplateId = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) SetWatermarks(v []*DescribeRecordTemplatesResponseBodyTemplatesWatermarks) *DescribeRecordTemplatesResponseBodyTemplates {
	s.Watermarks = v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) Validate() error {
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

type DescribeRecordTemplatesResponseBodyTemplatesBackgrounds struct {
	// Specifies whether to display the background image. Valid values:
	//
	// - **0*	- (None): Not displayed.
	//
	// - **1*	- (Always): Always displayed.
	//
	// example:
	//
	// 0
	Display *int32 `json:"Display,omitempty" xml:"Display,omitempty"`
	// The height of the background image pane. The value is a normalized percentage.
	//
	// example:
	//
	// 0.2456
	Height *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The HTTP or HTTPS URL of the background image.
	//
	// example:
	//
	// https://www.example.com/image.jpg
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
	// The width of the background image pane. The value is a normalized percentage.
	//
	// example:
	//
	// 0.2456
	Width *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// The X-coordinate of the background image. The value is a normalized percentage.
	//
	// example:
	//
	// 0.7576
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// The Y-coordinate of the background image. The value is a normalized percentage.
	//
	// example:
	//
	// 0.7576
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	// The stacking order of the background image. A value of 0 indicates the bottom layer. A value of 1 indicates the layer above the bottom layer, and so on.
	//
	// example:
	//
	// 0
	ZOrder *int32 `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s DescribeRecordTemplatesResponseBodyTemplatesBackgrounds) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordTemplatesResponseBodyTemplatesBackgrounds) GoString() string {
	return s.String()
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesBackgrounds) GetDisplay() *int32 {
	return s.Display
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesBackgrounds) GetHeight() *float32 {
	return s.Height
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesBackgrounds) GetUrl() *string {
	return s.Url
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesBackgrounds) GetWidth() *float32 {
	return s.Width
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesBackgrounds) GetX() *float32 {
	return s.X
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesBackgrounds) GetY() *float32 {
	return s.Y
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesBackgrounds) GetZOrder() *int32 {
	return s.ZOrder
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesBackgrounds) SetDisplay(v int32) *DescribeRecordTemplatesResponseBodyTemplatesBackgrounds {
	s.Display = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesBackgrounds) SetHeight(v float32) *DescribeRecordTemplatesResponseBodyTemplatesBackgrounds {
	s.Height = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesBackgrounds) SetUrl(v string) *DescribeRecordTemplatesResponseBodyTemplatesBackgrounds {
	s.Url = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesBackgrounds) SetWidth(v float32) *DescribeRecordTemplatesResponseBodyTemplatesBackgrounds {
	s.Width = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesBackgrounds) SetX(v float32) *DescribeRecordTemplatesResponseBodyTemplatesBackgrounds {
	s.X = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesBackgrounds) SetY(v float32) *DescribeRecordTemplatesResponseBodyTemplatesBackgrounds {
	s.Y = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesBackgrounds) SetZOrder(v int32) *DescribeRecordTemplatesResponseBodyTemplatesBackgrounds {
	s.ZOrder = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesBackgrounds) Validate() error {
	return dara.Validate(s)
}

type DescribeRecordTemplatesResponseBodyTemplatesClockWidgets struct {
	// The text color of the clock (RGB).
	//
	// The calculation formula is `R + G × 256 + B × 65536`. Valid values for R (red), G (green), and B (blue): **0*	- to **255**.
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
	// The X-coordinate of the clock. The value is a normalized percentage.
	//
	// example:
	//
	// 0.7576
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// The Y-coordinate of the clock. The value is a normalized percentage.
	//
	// example:
	//
	// 0.7576
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	// The stacking order of the clock. A value of 0 indicates the bottom layer. A value of 1 indicates the layer above the bottom layer, and so on.
	//
	// example:
	//
	// 0
	ZOrder *int32 `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s DescribeRecordTemplatesResponseBodyTemplatesClockWidgets) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordTemplatesResponseBodyTemplatesClockWidgets) GoString() string {
	return s.String()
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesClockWidgets) GetFontColor() *int32 {
	return s.FontColor
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesClockWidgets) GetFontSize() *int32 {
	return s.FontSize
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesClockWidgets) GetFontType() *int32 {
	return s.FontType
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesClockWidgets) GetX() *float32 {
	return s.X
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesClockWidgets) GetY() *float32 {
	return s.Y
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesClockWidgets) GetZOrder() *int32 {
	return s.ZOrder
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesClockWidgets) SetFontColor(v int32) *DescribeRecordTemplatesResponseBodyTemplatesClockWidgets {
	s.FontColor = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesClockWidgets) SetFontSize(v int32) *DescribeRecordTemplatesResponseBodyTemplatesClockWidgets {
	s.FontSize = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesClockWidgets) SetFontType(v int32) *DescribeRecordTemplatesResponseBodyTemplatesClockWidgets {
	s.FontType = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesClockWidgets) SetX(v float32) *DescribeRecordTemplatesResponseBodyTemplatesClockWidgets {
	s.X = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesClockWidgets) SetY(v float32) *DescribeRecordTemplatesResponseBodyTemplatesClockWidgets {
	s.Y = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesClockWidgets) SetZOrder(v int32) *DescribeRecordTemplatesResponseBodyTemplatesClockWidgets {
	s.ZOrder = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesClockWidgets) Validate() error {
	return dara.Validate(s)
}

type DescribeRecordTemplatesResponseBodyTemplatesWatermarks struct {
	// The transparency of the watermark. The value is a normalized percentage. A value of 0 indicates transparent, and a value of 1 indicates completely opaque.
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
	// The height of the watermark pane. The value is a normalized percentage.
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
	// The width of the watermark pane. The value is a normalized percentage.
	//
	// example:
	//
	// 0.2456
	Width *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// The X-coordinate of the watermark. The value is a normalized percentage.
	//
	// example:
	//
	// 0.7576
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// The Y-coordinate of the watermark. The value is a normalized percentage.
	//
	// example:
	//
	// 0.7576
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	// The stacking order of the watermark. A value of 0 indicates the bottom layer. A value of 1 indicates the layer above the bottom layer, and so on.
	//
	// example:
	//
	// 0
	ZOrder *int32 `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s DescribeRecordTemplatesResponseBodyTemplatesWatermarks) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordTemplatesResponseBodyTemplatesWatermarks) GoString() string {
	return s.String()
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesWatermarks) GetAlpha() *float32 {
	return s.Alpha
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesWatermarks) GetDisplay() *int32 {
	return s.Display
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesWatermarks) GetHeight() *float32 {
	return s.Height
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesWatermarks) GetUrl() *string {
	return s.Url
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesWatermarks) GetWidth() *float32 {
	return s.Width
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesWatermarks) GetX() *float32 {
	return s.X
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesWatermarks) GetY() *float32 {
	return s.Y
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesWatermarks) GetZOrder() *int32 {
	return s.ZOrder
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesWatermarks) SetAlpha(v float32) *DescribeRecordTemplatesResponseBodyTemplatesWatermarks {
	s.Alpha = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesWatermarks) SetDisplay(v int32) *DescribeRecordTemplatesResponseBodyTemplatesWatermarks {
	s.Display = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesWatermarks) SetHeight(v float32) *DescribeRecordTemplatesResponseBodyTemplatesWatermarks {
	s.Height = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesWatermarks) SetUrl(v string) *DescribeRecordTemplatesResponseBodyTemplatesWatermarks {
	s.Url = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesWatermarks) SetWidth(v float32) *DescribeRecordTemplatesResponseBodyTemplatesWatermarks {
	s.Width = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesWatermarks) SetX(v float32) *DescribeRecordTemplatesResponseBodyTemplatesWatermarks {
	s.X = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesWatermarks) SetY(v float32) *DescribeRecordTemplatesResponseBodyTemplatesWatermarks {
	s.Y = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesWatermarks) SetZOrder(v int32) *DescribeRecordTemplatesResponseBodyTemplatesWatermarks {
	s.ZOrder = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesWatermarks) Validate() error {
	return dara.Validate(s)
}
