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
	// example:
	//
	// C292B80E-5175-4BA4-8CC292B80E-5175-4BA4-8C1E-2ABEC4D7C2FE1E-2ABEC4D7****
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Templates []*DescribeRecordTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
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
	// example:
	//
	// 0
	BackgroundColor *int32                                                      `json:"BackgroundColor,omitempty" xml:"BackgroundColor,omitempty"`
	Backgrounds     []*DescribeRecordTemplatesResponseBodyTemplatesBackgrounds  `json:"Backgrounds,omitempty" xml:"Backgrounds,omitempty" type:"Repeated"`
	ClockWidgets    []*DescribeRecordTemplatesResponseBodyTemplatesClockWidgets `json:"ClockWidgets,omitempty" xml:"ClockWidgets,omitempty" type:"Repeated"`
	// example:
	//
	// 2020-09-04T06:22:15Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 180
	DelayStopTime *int32 `json:"DelayStopTime,omitempty" xml:"DelayStopTime,omitempty"`
	// example:
	//
	// false
	EnableM3u8DateTime *bool `json:"EnableM3u8DateTime,omitempty" xml:"EnableM3u8DateTime,omitempty"`
	// example:
	//
	// 1800
	FileSplitInterval *int32    `json:"FileSplitInterval,omitempty" xml:"FileSplitInterval,omitempty"`
	Formats           []*string `json:"Formats,omitempty" xml:"Formats,omitempty" type:"Repeated"`
	// example:
	//
	// http://example.com/callback
	HttpCallbackUrl *string  `json:"HttpCallbackUrl,omitempty" xml:"HttpCallbackUrl,omitempty"`
	LayoutIds       []*int64 `json:"LayoutIds,omitempty" xml:"LayoutIds,omitempty" type:"Repeated"`
	// example:
	//
	// 50
	MediaEncode *int32 `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	// example:
	//
	// record-callback-queue
	MnsQueue *string `json:"MnsQueue,omitempty" xml:"MnsQueue,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// rtc-record-oss
	OssBucket *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	// example:
	//
	// record/pre/{AppId}/{ChannelId_TaskId}/{EscapedStartTime}_{EscapedEndTime}
	OssFilePrefix *string `json:"OssFilePrefix,omitempty" xml:"OssFilePrefix,omitempty"`
	// example:
	//
	// 4IN_1080P
	TaskProfile *string `json:"TaskProfile,omitempty" xml:"TaskProfile,omitempty"`
	// example:
	//
	// 1ca698e2-57fa-4314-8e11-00d950d4****
	TemplateId *string                                                   `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
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
