// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartMPUTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *StartMPUTaskRequest
	GetAppId() *string
	SetBackgroundColor(v int32) *StartMPUTaskRequest
	GetBackgroundColor() *int32
	SetBackgrounds(v []*StartMPUTaskRequestBackgrounds) *StartMPUTaskRequest
	GetBackgrounds() []*StartMPUTaskRequestBackgrounds
	SetChannelId(v string) *StartMPUTaskRequest
	GetChannelId() *string
	SetClockWidgets(v []*StartMPUTaskRequestClockWidgets) *StartMPUTaskRequest
	GetClockWidgets() []*StartMPUTaskRequestClockWidgets
	SetCropMode(v int32) *StartMPUTaskRequest
	GetCropMode() *int32
	SetEnhancedParam(v *StartMPUTaskRequestEnhancedParam) *StartMPUTaskRequest
	GetEnhancedParam() *StartMPUTaskRequestEnhancedParam
	SetLayoutIds(v []*int64) *StartMPUTaskRequest
	GetLayoutIds() []*int64
	SetMediaEncode(v int32) *StartMPUTaskRequest
	GetMediaEncode() *int32
	SetMixMode(v int32) *StartMPUTaskRequest
	GetMixMode() *int32
	SetOwnerId(v int64) *StartMPUTaskRequest
	GetOwnerId() *int64
	SetPayloadType(v int32) *StartMPUTaskRequest
	GetPayloadType() *int32
	SetReportVad(v int32) *StartMPUTaskRequest
	GetReportVad() *int32
	SetRtpExtInfo(v int32) *StartMPUTaskRequest
	GetRtpExtInfo() *int32
	SetSourceType(v string) *StartMPUTaskRequest
	GetSourceType() *string
	SetStreamType(v int32) *StartMPUTaskRequest
	GetStreamType() *int32
	SetStreamURL(v string) *StartMPUTaskRequest
	GetStreamURL() *string
	SetSubSpecAudioUsers(v []*string) *StartMPUTaskRequest
	GetSubSpecAudioUsers() []*string
	SetSubSpecCameraUsers(v []*string) *StartMPUTaskRequest
	GetSubSpecCameraUsers() []*string
	SetSubSpecShareScreenUsers(v []*string) *StartMPUTaskRequest
	GetSubSpecShareScreenUsers() []*string
	SetSubSpecUsers(v []*string) *StartMPUTaskRequest
	GetSubSpecUsers() []*string
	SetTaskId(v string) *StartMPUTaskRequest
	GetTaskId() *string
	SetTaskType(v int32) *StartMPUTaskRequest
	GetTaskType() *int32
	SetTimeStampRef(v int64) *StartMPUTaskRequest
	GetTimeStampRef() *int64
	SetUnsubSpecAudioUsers(v []*string) *StartMPUTaskRequest
	GetUnsubSpecAudioUsers() []*string
	SetUnsubSpecCameraUsers(v []*string) *StartMPUTaskRequest
	GetUnsubSpecCameraUsers() []*string
	SetUnsubSpecShareScreenUsers(v []*string) *StartMPUTaskRequest
	GetUnsubSpecShareScreenUsers() []*string
	SetUserPanes(v []*StartMPUTaskRequestUserPanes) *StartMPUTaskRequest
	GetUserPanes() []*StartMPUTaskRequestUserPanes
	SetVadInterval(v int64) *StartMPUTaskRequest
	GetVadInterval() *int64
	SetWatermarks(v []*StartMPUTaskRequestWatermarks) *StartMPUTaskRequest
	GetWatermarks() []*StartMPUTaskRequestWatermarks
}

type StartMPUTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 0
	BackgroundColor *int32                            `json:"BackgroundColor,omitempty" xml:"BackgroundColor,omitempty"`
	Backgrounds     []*StartMPUTaskRequestBackgrounds `json:"Backgrounds,omitempty" xml:"Backgrounds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// yourChannelId
	ChannelId    *string                            `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ClockWidgets []*StartMPUTaskRequestClockWidgets `json:"ClockWidgets,omitempty" xml:"ClockWidgets,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	CropMode      *int32                            `json:"CropMode,omitempty" xml:"CropMode,omitempty"`
	EnhancedParam *StartMPUTaskRequestEnhancedParam `json:"EnhancedParam,omitempty" xml:"EnhancedParam,omitempty" type:"Struct"`
	// example:
	//
	// 1
	LayoutIds []*int64 `json:"LayoutIds,omitempty" xml:"LayoutIds,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	MediaEncode *int32 `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	// example:
	//
	// 0
	MixMode *int32 `json:"MixMode,omitempty" xml:"MixMode,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 0
	PayloadType *int32 `json:"PayloadType,omitempty" xml:"PayloadType,omitempty"`
	// example:
	//
	// 0
	ReportVad *int32 `json:"ReportVad,omitempty" xml:"ReportVad,omitempty"`
	// example:
	//
	// 0
	RtpExtInfo *int32 `json:"RtpExtInfo,omitempty" xml:"RtpExtInfo,omitempty"`
	// example:
	//
	// camera
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// 0
	StreamType *int32 `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	// example:
	//
	// rtmp://example.com/live/stream
	StreamURL *string `json:"StreamURL,omitempty" xml:"StreamURL,omitempty"`
	// example:
	//
	// audioUserID
	SubSpecAudioUsers       []*string `json:"SubSpecAudioUsers,omitempty" xml:"SubSpecAudioUsers,omitempty" type:"Repeated"`
	SubSpecCameraUsers      []*string `json:"SubSpecCameraUsers,omitempty" xml:"SubSpecCameraUsers,omitempty" type:"Repeated"`
	SubSpecShareScreenUsers []*string `json:"SubSpecShareScreenUsers,omitempty" xml:"SubSpecShareScreenUsers,omitempty" type:"Repeated"`
	// example:
	//
	// userID
	SubSpecUsers []*string `json:"SubSpecUsers,omitempty" xml:"SubSpecUsers,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// yourTaskId
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 0
	TaskType *int32 `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// example:
	//
	// 15273582735
	TimeStampRef              *int64                          `json:"TimeStampRef,omitempty" xml:"TimeStampRef,omitempty"`
	UnsubSpecAudioUsers       []*string                       `json:"UnsubSpecAudioUsers,omitempty" xml:"UnsubSpecAudioUsers,omitempty" type:"Repeated"`
	UnsubSpecCameraUsers      []*string                       `json:"UnsubSpecCameraUsers,omitempty" xml:"UnsubSpecCameraUsers,omitempty" type:"Repeated"`
	UnsubSpecShareScreenUsers []*string                       `json:"UnsubSpecShareScreenUsers,omitempty" xml:"UnsubSpecShareScreenUsers,omitempty" type:"Repeated"`
	UserPanes                 []*StartMPUTaskRequestUserPanes `json:"UserPanes,omitempty" xml:"UserPanes,omitempty" type:"Repeated"`
	// example:
	//
	// 86400
	VadInterval *int64                           `json:"VadInterval,omitempty" xml:"VadInterval,omitempty"`
	Watermarks  []*StartMPUTaskRequestWatermarks `json:"Watermarks,omitempty" xml:"Watermarks,omitempty" type:"Repeated"`
}

func (s StartMPUTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s StartMPUTaskRequest) GoString() string {
	return s.String()
}

func (s *StartMPUTaskRequest) GetAppId() *string {
	return s.AppId
}

func (s *StartMPUTaskRequest) GetBackgroundColor() *int32 {
	return s.BackgroundColor
}

func (s *StartMPUTaskRequest) GetBackgrounds() []*StartMPUTaskRequestBackgrounds {
	return s.Backgrounds
}

func (s *StartMPUTaskRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *StartMPUTaskRequest) GetClockWidgets() []*StartMPUTaskRequestClockWidgets {
	return s.ClockWidgets
}

func (s *StartMPUTaskRequest) GetCropMode() *int32 {
	return s.CropMode
}

func (s *StartMPUTaskRequest) GetEnhancedParam() *StartMPUTaskRequestEnhancedParam {
	return s.EnhancedParam
}

func (s *StartMPUTaskRequest) GetLayoutIds() []*int64 {
	return s.LayoutIds
}

func (s *StartMPUTaskRequest) GetMediaEncode() *int32 {
	return s.MediaEncode
}

func (s *StartMPUTaskRequest) GetMixMode() *int32 {
	return s.MixMode
}

func (s *StartMPUTaskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StartMPUTaskRequest) GetPayloadType() *int32 {
	return s.PayloadType
}

func (s *StartMPUTaskRequest) GetReportVad() *int32 {
	return s.ReportVad
}

func (s *StartMPUTaskRequest) GetRtpExtInfo() *int32 {
	return s.RtpExtInfo
}

func (s *StartMPUTaskRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *StartMPUTaskRequest) GetStreamType() *int32 {
	return s.StreamType
}

func (s *StartMPUTaskRequest) GetStreamURL() *string {
	return s.StreamURL
}

func (s *StartMPUTaskRequest) GetSubSpecAudioUsers() []*string {
	return s.SubSpecAudioUsers
}

func (s *StartMPUTaskRequest) GetSubSpecCameraUsers() []*string {
	return s.SubSpecCameraUsers
}

func (s *StartMPUTaskRequest) GetSubSpecShareScreenUsers() []*string {
	return s.SubSpecShareScreenUsers
}

func (s *StartMPUTaskRequest) GetSubSpecUsers() []*string {
	return s.SubSpecUsers
}

func (s *StartMPUTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *StartMPUTaskRequest) GetTaskType() *int32 {
	return s.TaskType
}

func (s *StartMPUTaskRequest) GetTimeStampRef() *int64 {
	return s.TimeStampRef
}

func (s *StartMPUTaskRequest) GetUnsubSpecAudioUsers() []*string {
	return s.UnsubSpecAudioUsers
}

func (s *StartMPUTaskRequest) GetUnsubSpecCameraUsers() []*string {
	return s.UnsubSpecCameraUsers
}

func (s *StartMPUTaskRequest) GetUnsubSpecShareScreenUsers() []*string {
	return s.UnsubSpecShareScreenUsers
}

func (s *StartMPUTaskRequest) GetUserPanes() []*StartMPUTaskRequestUserPanes {
	return s.UserPanes
}

func (s *StartMPUTaskRequest) GetVadInterval() *int64 {
	return s.VadInterval
}

func (s *StartMPUTaskRequest) GetWatermarks() []*StartMPUTaskRequestWatermarks {
	return s.Watermarks
}

func (s *StartMPUTaskRequest) SetAppId(v string) *StartMPUTaskRequest {
	s.AppId = &v
	return s
}

func (s *StartMPUTaskRequest) SetBackgroundColor(v int32) *StartMPUTaskRequest {
	s.BackgroundColor = &v
	return s
}

func (s *StartMPUTaskRequest) SetBackgrounds(v []*StartMPUTaskRequestBackgrounds) *StartMPUTaskRequest {
	s.Backgrounds = v
	return s
}

func (s *StartMPUTaskRequest) SetChannelId(v string) *StartMPUTaskRequest {
	s.ChannelId = &v
	return s
}

func (s *StartMPUTaskRequest) SetClockWidgets(v []*StartMPUTaskRequestClockWidgets) *StartMPUTaskRequest {
	s.ClockWidgets = v
	return s
}

func (s *StartMPUTaskRequest) SetCropMode(v int32) *StartMPUTaskRequest {
	s.CropMode = &v
	return s
}

func (s *StartMPUTaskRequest) SetEnhancedParam(v *StartMPUTaskRequestEnhancedParam) *StartMPUTaskRequest {
	s.EnhancedParam = v
	return s
}

func (s *StartMPUTaskRequest) SetLayoutIds(v []*int64) *StartMPUTaskRequest {
	s.LayoutIds = v
	return s
}

func (s *StartMPUTaskRequest) SetMediaEncode(v int32) *StartMPUTaskRequest {
	s.MediaEncode = &v
	return s
}

func (s *StartMPUTaskRequest) SetMixMode(v int32) *StartMPUTaskRequest {
	s.MixMode = &v
	return s
}

func (s *StartMPUTaskRequest) SetOwnerId(v int64) *StartMPUTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *StartMPUTaskRequest) SetPayloadType(v int32) *StartMPUTaskRequest {
	s.PayloadType = &v
	return s
}

func (s *StartMPUTaskRequest) SetReportVad(v int32) *StartMPUTaskRequest {
	s.ReportVad = &v
	return s
}

func (s *StartMPUTaskRequest) SetRtpExtInfo(v int32) *StartMPUTaskRequest {
	s.RtpExtInfo = &v
	return s
}

func (s *StartMPUTaskRequest) SetSourceType(v string) *StartMPUTaskRequest {
	s.SourceType = &v
	return s
}

func (s *StartMPUTaskRequest) SetStreamType(v int32) *StartMPUTaskRequest {
	s.StreamType = &v
	return s
}

func (s *StartMPUTaskRequest) SetStreamURL(v string) *StartMPUTaskRequest {
	s.StreamURL = &v
	return s
}

func (s *StartMPUTaskRequest) SetSubSpecAudioUsers(v []*string) *StartMPUTaskRequest {
	s.SubSpecAudioUsers = v
	return s
}

func (s *StartMPUTaskRequest) SetSubSpecCameraUsers(v []*string) *StartMPUTaskRequest {
	s.SubSpecCameraUsers = v
	return s
}

func (s *StartMPUTaskRequest) SetSubSpecShareScreenUsers(v []*string) *StartMPUTaskRequest {
	s.SubSpecShareScreenUsers = v
	return s
}

func (s *StartMPUTaskRequest) SetSubSpecUsers(v []*string) *StartMPUTaskRequest {
	s.SubSpecUsers = v
	return s
}

func (s *StartMPUTaskRequest) SetTaskId(v string) *StartMPUTaskRequest {
	s.TaskId = &v
	return s
}

func (s *StartMPUTaskRequest) SetTaskType(v int32) *StartMPUTaskRequest {
	s.TaskType = &v
	return s
}

func (s *StartMPUTaskRequest) SetTimeStampRef(v int64) *StartMPUTaskRequest {
	s.TimeStampRef = &v
	return s
}

func (s *StartMPUTaskRequest) SetUnsubSpecAudioUsers(v []*string) *StartMPUTaskRequest {
	s.UnsubSpecAudioUsers = v
	return s
}

func (s *StartMPUTaskRequest) SetUnsubSpecCameraUsers(v []*string) *StartMPUTaskRequest {
	s.UnsubSpecCameraUsers = v
	return s
}

func (s *StartMPUTaskRequest) SetUnsubSpecShareScreenUsers(v []*string) *StartMPUTaskRequest {
	s.UnsubSpecShareScreenUsers = v
	return s
}

func (s *StartMPUTaskRequest) SetUserPanes(v []*StartMPUTaskRequestUserPanes) *StartMPUTaskRequest {
	s.UserPanes = v
	return s
}

func (s *StartMPUTaskRequest) SetVadInterval(v int64) *StartMPUTaskRequest {
	s.VadInterval = &v
	return s
}

func (s *StartMPUTaskRequest) SetWatermarks(v []*StartMPUTaskRequestWatermarks) *StartMPUTaskRequest {
	s.Watermarks = v
	return s
}

func (s *StartMPUTaskRequest) Validate() error {
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
	if s.EnhancedParam != nil {
		if err := s.EnhancedParam.Validate(); err != nil {
			return err
		}
	}
	if s.UserPanes != nil {
		for _, item := range s.UserPanes {
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

type StartMPUTaskRequestBackgrounds struct {
	// example:
	//
	// 1
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

func (s StartMPUTaskRequestBackgrounds) String() string {
	return dara.Prettify(s)
}

func (s StartMPUTaskRequestBackgrounds) GoString() string {
	return s.String()
}

func (s *StartMPUTaskRequestBackgrounds) GetDisplay() *int32 {
	return s.Display
}

func (s *StartMPUTaskRequestBackgrounds) GetHeight() *float32 {
	return s.Height
}

func (s *StartMPUTaskRequestBackgrounds) GetUrl() *string {
	return s.Url
}

func (s *StartMPUTaskRequestBackgrounds) GetWidth() *float32 {
	return s.Width
}

func (s *StartMPUTaskRequestBackgrounds) GetX() *float32 {
	return s.X
}

func (s *StartMPUTaskRequestBackgrounds) GetY() *float32 {
	return s.Y
}

func (s *StartMPUTaskRequestBackgrounds) GetZOrder() *int32 {
	return s.ZOrder
}

func (s *StartMPUTaskRequestBackgrounds) SetDisplay(v int32) *StartMPUTaskRequestBackgrounds {
	s.Display = &v
	return s
}

func (s *StartMPUTaskRequestBackgrounds) SetHeight(v float32) *StartMPUTaskRequestBackgrounds {
	s.Height = &v
	return s
}

func (s *StartMPUTaskRequestBackgrounds) SetUrl(v string) *StartMPUTaskRequestBackgrounds {
	s.Url = &v
	return s
}

func (s *StartMPUTaskRequestBackgrounds) SetWidth(v float32) *StartMPUTaskRequestBackgrounds {
	s.Width = &v
	return s
}

func (s *StartMPUTaskRequestBackgrounds) SetX(v float32) *StartMPUTaskRequestBackgrounds {
	s.X = &v
	return s
}

func (s *StartMPUTaskRequestBackgrounds) SetY(v float32) *StartMPUTaskRequestBackgrounds {
	s.Y = &v
	return s
}

func (s *StartMPUTaskRequestBackgrounds) SetZOrder(v int32) *StartMPUTaskRequestBackgrounds {
	s.ZOrder = &v
	return s
}

func (s *StartMPUTaskRequestBackgrounds) Validate() error {
	return dara.Validate(s)
}

type StartMPUTaskRequestClockWidgets struct {
	// example:
	//
	// 0
	Alpha *float32 `json:"Alpha,omitempty" xml:"Alpha,omitempty"`
	// example:
	//
	// 0
	BorderColor *int64 `json:"BorderColor,omitempty" xml:"BorderColor,omitempty"`
	// example:
	//
	// 1
	BorderWidth *int32 `json:"BorderWidth,omitempty" xml:"BorderWidth,omitempty"`
	// example:
	//
	// false
	Box *bool `json:"Box,omitempty" xml:"Box,omitempty"`
	// example:
	//
	// 0
	BoxBorderWidth *int32 `json:"BoxBorderWidth,omitempty" xml:"BoxBorderWidth,omitempty"`
	// example:
	//
	// 0
	BoxColor *int64 `json:"BoxColor,omitempty" xml:"BoxColor,omitempty"`
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

func (s StartMPUTaskRequestClockWidgets) String() string {
	return dara.Prettify(s)
}

func (s StartMPUTaskRequestClockWidgets) GoString() string {
	return s.String()
}

func (s *StartMPUTaskRequestClockWidgets) GetAlpha() *float32 {
	return s.Alpha
}

func (s *StartMPUTaskRequestClockWidgets) GetBorderColor() *int64 {
	return s.BorderColor
}

func (s *StartMPUTaskRequestClockWidgets) GetBorderWidth() *int32 {
	return s.BorderWidth
}

func (s *StartMPUTaskRequestClockWidgets) GetBox() *bool {
	return s.Box
}

func (s *StartMPUTaskRequestClockWidgets) GetBoxBorderWidth() *int32 {
	return s.BoxBorderWidth
}

func (s *StartMPUTaskRequestClockWidgets) GetBoxColor() *int64 {
	return s.BoxColor
}

func (s *StartMPUTaskRequestClockWidgets) GetFontColor() *int32 {
	return s.FontColor
}

func (s *StartMPUTaskRequestClockWidgets) GetFontSize() *int32 {
	return s.FontSize
}

func (s *StartMPUTaskRequestClockWidgets) GetFontType() *int32 {
	return s.FontType
}

func (s *StartMPUTaskRequestClockWidgets) GetX() *float32 {
	return s.X
}

func (s *StartMPUTaskRequestClockWidgets) GetY() *float32 {
	return s.Y
}

func (s *StartMPUTaskRequestClockWidgets) GetZOrder() *int32 {
	return s.ZOrder
}

func (s *StartMPUTaskRequestClockWidgets) SetAlpha(v float32) *StartMPUTaskRequestClockWidgets {
	s.Alpha = &v
	return s
}

func (s *StartMPUTaskRequestClockWidgets) SetBorderColor(v int64) *StartMPUTaskRequestClockWidgets {
	s.BorderColor = &v
	return s
}

func (s *StartMPUTaskRequestClockWidgets) SetBorderWidth(v int32) *StartMPUTaskRequestClockWidgets {
	s.BorderWidth = &v
	return s
}

func (s *StartMPUTaskRequestClockWidgets) SetBox(v bool) *StartMPUTaskRequestClockWidgets {
	s.Box = &v
	return s
}

func (s *StartMPUTaskRequestClockWidgets) SetBoxBorderWidth(v int32) *StartMPUTaskRequestClockWidgets {
	s.BoxBorderWidth = &v
	return s
}

func (s *StartMPUTaskRequestClockWidgets) SetBoxColor(v int64) *StartMPUTaskRequestClockWidgets {
	s.BoxColor = &v
	return s
}

func (s *StartMPUTaskRequestClockWidgets) SetFontColor(v int32) *StartMPUTaskRequestClockWidgets {
	s.FontColor = &v
	return s
}

func (s *StartMPUTaskRequestClockWidgets) SetFontSize(v int32) *StartMPUTaskRequestClockWidgets {
	s.FontSize = &v
	return s
}

func (s *StartMPUTaskRequestClockWidgets) SetFontType(v int32) *StartMPUTaskRequestClockWidgets {
	s.FontType = &v
	return s
}

func (s *StartMPUTaskRequestClockWidgets) SetX(v float32) *StartMPUTaskRequestClockWidgets {
	s.X = &v
	return s
}

func (s *StartMPUTaskRequestClockWidgets) SetY(v float32) *StartMPUTaskRequestClockWidgets {
	s.Y = &v
	return s
}

func (s *StartMPUTaskRequestClockWidgets) SetZOrder(v int32) *StartMPUTaskRequestClockWidgets {
	s.ZOrder = &v
	return s
}

func (s *StartMPUTaskRequestClockWidgets) Validate() error {
	return dara.Validate(s)
}

type StartMPUTaskRequestEnhancedParam struct {
	// example:
	//
	// false
	EnablePortraitSegmentation *bool `json:"EnablePortraitSegmentation,omitempty" xml:"EnablePortraitSegmentation,omitempty"`
}

func (s StartMPUTaskRequestEnhancedParam) String() string {
	return dara.Prettify(s)
}

func (s StartMPUTaskRequestEnhancedParam) GoString() string {
	return s.String()
}

func (s *StartMPUTaskRequestEnhancedParam) GetEnablePortraitSegmentation() *bool {
	return s.EnablePortraitSegmentation
}

func (s *StartMPUTaskRequestEnhancedParam) SetEnablePortraitSegmentation(v bool) *StartMPUTaskRequestEnhancedParam {
	s.EnablePortraitSegmentation = &v
	return s
}

func (s *StartMPUTaskRequestEnhancedParam) Validate() error {
	return dara.Validate(s)
}

type StartMPUTaskRequestUserPanes struct {
	Images []*StartMPUTaskRequestUserPanesImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	PaneId *int32 `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
	// example:
	//
	// 0
	SegmentType *int32 `json:"SegmentType,omitempty" xml:"SegmentType,omitempty"`
	// example:
	//
	// camera
	SourceType *string                              `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Texts      []*StartMPUTaskRequestUserPanesTexts `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
	// example:
	//
	// TestId
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s StartMPUTaskRequestUserPanes) String() string {
	return dara.Prettify(s)
}

func (s StartMPUTaskRequestUserPanes) GoString() string {
	return s.String()
}

func (s *StartMPUTaskRequestUserPanes) GetImages() []*StartMPUTaskRequestUserPanesImages {
	return s.Images
}

func (s *StartMPUTaskRequestUserPanes) GetPaneId() *int32 {
	return s.PaneId
}

func (s *StartMPUTaskRequestUserPanes) GetSegmentType() *int32 {
	return s.SegmentType
}

func (s *StartMPUTaskRequestUserPanes) GetSourceType() *string {
	return s.SourceType
}

func (s *StartMPUTaskRequestUserPanes) GetTexts() []*StartMPUTaskRequestUserPanesTexts {
	return s.Texts
}

func (s *StartMPUTaskRequestUserPanes) GetUserId() *string {
	return s.UserId
}

func (s *StartMPUTaskRequestUserPanes) SetImages(v []*StartMPUTaskRequestUserPanesImages) *StartMPUTaskRequestUserPanes {
	s.Images = v
	return s
}

func (s *StartMPUTaskRequestUserPanes) SetPaneId(v int32) *StartMPUTaskRequestUserPanes {
	s.PaneId = &v
	return s
}

func (s *StartMPUTaskRequestUserPanes) SetSegmentType(v int32) *StartMPUTaskRequestUserPanes {
	s.SegmentType = &v
	return s
}

func (s *StartMPUTaskRequestUserPanes) SetSourceType(v string) *StartMPUTaskRequestUserPanes {
	s.SourceType = &v
	return s
}

func (s *StartMPUTaskRequestUserPanes) SetTexts(v []*StartMPUTaskRequestUserPanesTexts) *StartMPUTaskRequestUserPanes {
	s.Texts = v
	return s
}

func (s *StartMPUTaskRequestUserPanes) SetUserId(v string) *StartMPUTaskRequestUserPanes {
	s.UserId = &v
	return s
}

func (s *StartMPUTaskRequestUserPanes) Validate() error {
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

type StartMPUTaskRequestUserPanesImages struct {
	// example:
	//
	// 1
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

func (s StartMPUTaskRequestUserPanesImages) String() string {
	return dara.Prettify(s)
}

func (s StartMPUTaskRequestUserPanesImages) GoString() string {
	return s.String()
}

func (s *StartMPUTaskRequestUserPanesImages) GetDisplay() *int32 {
	return s.Display
}

func (s *StartMPUTaskRequestUserPanesImages) GetHeight() *float32 {
	return s.Height
}

func (s *StartMPUTaskRequestUserPanesImages) GetUrl() *string {
	return s.Url
}

func (s *StartMPUTaskRequestUserPanesImages) GetWidth() *float32 {
	return s.Width
}

func (s *StartMPUTaskRequestUserPanesImages) GetX() *float32 {
	return s.X
}

func (s *StartMPUTaskRequestUserPanesImages) GetY() *float32 {
	return s.Y
}

func (s *StartMPUTaskRequestUserPanesImages) GetZOrder() *int32 {
	return s.ZOrder
}

func (s *StartMPUTaskRequestUserPanesImages) SetDisplay(v int32) *StartMPUTaskRequestUserPanesImages {
	s.Display = &v
	return s
}

func (s *StartMPUTaskRequestUserPanesImages) SetHeight(v float32) *StartMPUTaskRequestUserPanesImages {
	s.Height = &v
	return s
}

func (s *StartMPUTaskRequestUserPanesImages) SetUrl(v string) *StartMPUTaskRequestUserPanesImages {
	s.Url = &v
	return s
}

func (s *StartMPUTaskRequestUserPanesImages) SetWidth(v float32) *StartMPUTaskRequestUserPanesImages {
	s.Width = &v
	return s
}

func (s *StartMPUTaskRequestUserPanesImages) SetX(v float32) *StartMPUTaskRequestUserPanesImages {
	s.X = &v
	return s
}

func (s *StartMPUTaskRequestUserPanesImages) SetY(v float32) *StartMPUTaskRequestUserPanesImages {
	s.Y = &v
	return s
}

func (s *StartMPUTaskRequestUserPanesImages) SetZOrder(v int32) *StartMPUTaskRequestUserPanesImages {
	s.ZOrder = &v
	return s
}

func (s *StartMPUTaskRequestUserPanesImages) Validate() error {
	return dara.Validate(s)
}

type StartMPUTaskRequestUserPanesTexts struct {
	// example:
	//
	// 0
	Alpha *float32 `json:"Alpha,omitempty" xml:"Alpha,omitempty"`
	// example:
	//
	// 0
	BorderColor *int64 `json:"BorderColor,omitempty" xml:"BorderColor,omitempty"`
	// example:
	//
	// 1
	BorderWidth *int32 `json:"BorderWidth,omitempty" xml:"BorderWidth,omitempty"`
	// example:
	//
	// false
	Box *bool `json:"Box,omitempty" xml:"Box,omitempty"`
	// example:
	//
	// 0
	BoxBorderWidth *int32 `json:"BoxBorderWidth,omitempty" xml:"BoxBorderWidth,omitempty"`
	// example:
	//
	// 0
	BoxColor *int64 `json:"BoxColor,omitempty" xml:"BoxColor,omitempty"`
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
	// text
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// example:
	//
	// 0.7576
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 0.2456
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	// example:
	//
	// 0
	ZOrder *int32 `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s StartMPUTaskRequestUserPanesTexts) String() string {
	return dara.Prettify(s)
}

func (s StartMPUTaskRequestUserPanesTexts) GoString() string {
	return s.String()
}

func (s *StartMPUTaskRequestUserPanesTexts) GetAlpha() *float32 {
	return s.Alpha
}

func (s *StartMPUTaskRequestUserPanesTexts) GetBorderColor() *int64 {
	return s.BorderColor
}

func (s *StartMPUTaskRequestUserPanesTexts) GetBorderWidth() *int32 {
	return s.BorderWidth
}

func (s *StartMPUTaskRequestUserPanesTexts) GetBox() *bool {
	return s.Box
}

func (s *StartMPUTaskRequestUserPanesTexts) GetBoxBorderWidth() *int32 {
	return s.BoxBorderWidth
}

func (s *StartMPUTaskRequestUserPanesTexts) GetBoxColor() *int64 {
	return s.BoxColor
}

func (s *StartMPUTaskRequestUserPanesTexts) GetFontColor() *int32 {
	return s.FontColor
}

func (s *StartMPUTaskRequestUserPanesTexts) GetFontSize() *int32 {
	return s.FontSize
}

func (s *StartMPUTaskRequestUserPanesTexts) GetFontType() *int32 {
	return s.FontType
}

func (s *StartMPUTaskRequestUserPanesTexts) GetText() *string {
	return s.Text
}

func (s *StartMPUTaskRequestUserPanesTexts) GetX() *float32 {
	return s.X
}

func (s *StartMPUTaskRequestUserPanesTexts) GetY() *float32 {
	return s.Y
}

func (s *StartMPUTaskRequestUserPanesTexts) GetZOrder() *int32 {
	return s.ZOrder
}

func (s *StartMPUTaskRequestUserPanesTexts) SetAlpha(v float32) *StartMPUTaskRequestUserPanesTexts {
	s.Alpha = &v
	return s
}

func (s *StartMPUTaskRequestUserPanesTexts) SetBorderColor(v int64) *StartMPUTaskRequestUserPanesTexts {
	s.BorderColor = &v
	return s
}

func (s *StartMPUTaskRequestUserPanesTexts) SetBorderWidth(v int32) *StartMPUTaskRequestUserPanesTexts {
	s.BorderWidth = &v
	return s
}

func (s *StartMPUTaskRequestUserPanesTexts) SetBox(v bool) *StartMPUTaskRequestUserPanesTexts {
	s.Box = &v
	return s
}

func (s *StartMPUTaskRequestUserPanesTexts) SetBoxBorderWidth(v int32) *StartMPUTaskRequestUserPanesTexts {
	s.BoxBorderWidth = &v
	return s
}

func (s *StartMPUTaskRequestUserPanesTexts) SetBoxColor(v int64) *StartMPUTaskRequestUserPanesTexts {
	s.BoxColor = &v
	return s
}

func (s *StartMPUTaskRequestUserPanesTexts) SetFontColor(v int32) *StartMPUTaskRequestUserPanesTexts {
	s.FontColor = &v
	return s
}

func (s *StartMPUTaskRequestUserPanesTexts) SetFontSize(v int32) *StartMPUTaskRequestUserPanesTexts {
	s.FontSize = &v
	return s
}

func (s *StartMPUTaskRequestUserPanesTexts) SetFontType(v int32) *StartMPUTaskRequestUserPanesTexts {
	s.FontType = &v
	return s
}

func (s *StartMPUTaskRequestUserPanesTexts) SetText(v string) *StartMPUTaskRequestUserPanesTexts {
	s.Text = &v
	return s
}

func (s *StartMPUTaskRequestUserPanesTexts) SetX(v float32) *StartMPUTaskRequestUserPanesTexts {
	s.X = &v
	return s
}

func (s *StartMPUTaskRequestUserPanesTexts) SetY(v float32) *StartMPUTaskRequestUserPanesTexts {
	s.Y = &v
	return s
}

func (s *StartMPUTaskRequestUserPanesTexts) SetZOrder(v int32) *StartMPUTaskRequestUserPanesTexts {
	s.ZOrder = &v
	return s
}

func (s *StartMPUTaskRequestUserPanesTexts) Validate() error {
	return dara.Validate(s)
}

type StartMPUTaskRequestWatermarks struct {
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

func (s StartMPUTaskRequestWatermarks) String() string {
	return dara.Prettify(s)
}

func (s StartMPUTaskRequestWatermarks) GoString() string {
	return s.String()
}

func (s *StartMPUTaskRequestWatermarks) GetAlpha() *float32 {
	return s.Alpha
}

func (s *StartMPUTaskRequestWatermarks) GetDisplay() *int32 {
	return s.Display
}

func (s *StartMPUTaskRequestWatermarks) GetHeight() *float32 {
	return s.Height
}

func (s *StartMPUTaskRequestWatermarks) GetUrl() *string {
	return s.Url
}

func (s *StartMPUTaskRequestWatermarks) GetWidth() *float32 {
	return s.Width
}

func (s *StartMPUTaskRequestWatermarks) GetX() *float32 {
	return s.X
}

func (s *StartMPUTaskRequestWatermarks) GetY() *float32 {
	return s.Y
}

func (s *StartMPUTaskRequestWatermarks) GetZOrder() *int32 {
	return s.ZOrder
}

func (s *StartMPUTaskRequestWatermarks) SetAlpha(v float32) *StartMPUTaskRequestWatermarks {
	s.Alpha = &v
	return s
}

func (s *StartMPUTaskRequestWatermarks) SetDisplay(v int32) *StartMPUTaskRequestWatermarks {
	s.Display = &v
	return s
}

func (s *StartMPUTaskRequestWatermarks) SetHeight(v float32) *StartMPUTaskRequestWatermarks {
	s.Height = &v
	return s
}

func (s *StartMPUTaskRequestWatermarks) SetUrl(v string) *StartMPUTaskRequestWatermarks {
	s.Url = &v
	return s
}

func (s *StartMPUTaskRequestWatermarks) SetWidth(v float32) *StartMPUTaskRequestWatermarks {
	s.Width = &v
	return s
}

func (s *StartMPUTaskRequestWatermarks) SetX(v float32) *StartMPUTaskRequestWatermarks {
	s.X = &v
	return s
}

func (s *StartMPUTaskRequestWatermarks) SetY(v float32) *StartMPUTaskRequestWatermarks {
	s.Y = &v
	return s
}

func (s *StartMPUTaskRequestWatermarks) SetZOrder(v int32) *StartMPUTaskRequestWatermarks {
	s.ZOrder = &v
	return s
}

func (s *StartMPUTaskRequestWatermarks) Validate() error {
	return dara.Validate(s)
}
