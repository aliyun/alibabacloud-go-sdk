// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddRecordTemplateRequest struct {
	AppId              *string                                 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	BackgroundColor    *int32                                  `json:"BackgroundColor,omitempty" xml:"BackgroundColor,omitempty"`
	Backgrounds        []*AddRecordTemplateRequestBackgrounds  `json:"Backgrounds,omitempty" xml:"Backgrounds,omitempty" type:"Repeated"`
	ClockWidgets       []*AddRecordTemplateRequestClockWidgets `json:"ClockWidgets,omitempty" xml:"ClockWidgets,omitempty" type:"Repeated"`
	DelayStopTime      *int32                                  `json:"DelayStopTime,omitempty" xml:"DelayStopTime,omitempty"`
	EnableM3u8DateTime *bool                                   `json:"EnableM3u8DateTime,omitempty" xml:"EnableM3u8DateTime,omitempty"`
	FileSplitInterval  *int32                                  `json:"FileSplitInterval,omitempty" xml:"FileSplitInterval,omitempty"`
	Formats            []*string                               `json:"Formats,omitempty" xml:"Formats,omitempty" type:"Repeated"`
	HttpCallbackUrl    *string                                 `json:"HttpCallbackUrl,omitempty" xml:"HttpCallbackUrl,omitempty"`
	LayoutIds          []*int64                                `json:"LayoutIds,omitempty" xml:"LayoutIds,omitempty" type:"Repeated"`
	MediaEncode        *int32                                  `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	MnsQueue           *string                                 `json:"MnsQueue,omitempty" xml:"MnsQueue,omitempty"`
	Name               *string                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	OssBucket          *string                                 `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssFilePrefix      *string                                 `json:"OssFilePrefix,omitempty" xml:"OssFilePrefix,omitempty"`
	OwnerId            *int64                                  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	TaskProfile        *string                                 `json:"TaskProfile,omitempty" xml:"TaskProfile,omitempty"`
	Watermarks         []*AddRecordTemplateRequestWatermarks   `json:"Watermarks,omitempty" xml:"Watermarks,omitempty" type:"Repeated"`
}

func (s AddRecordTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s AddRecordTemplateRequest) GoString() string {
	return s.String()
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

type AddRecordTemplateRequestBackgrounds struct {
	Display *int32   `json:"Display,omitempty" xml:"Display,omitempty"`
	Height  *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Url     *string  `json:"Url,omitempty" xml:"Url,omitempty"`
	Width   *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	X       *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y       *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	ZOrder  *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s AddRecordTemplateRequestBackgrounds) String() string {
	return tea.Prettify(s)
}

func (s AddRecordTemplateRequestBackgrounds) GoString() string {
	return s.String()
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

type AddRecordTemplateRequestClockWidgets struct {
	FontColor *int32   `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	FontSize  *int32   `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	FontType  *int32   `json:"FontType,omitempty" xml:"FontType,omitempty"`
	X         *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y         *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	ZOrder    *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s AddRecordTemplateRequestClockWidgets) String() string {
	return tea.Prettify(s)
}

func (s AddRecordTemplateRequestClockWidgets) GoString() string {
	return s.String()
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

type AddRecordTemplateRequestWatermarks struct {
	Alpha   *float32 `json:"Alpha,omitempty" xml:"Alpha,omitempty"`
	Display *int32   `json:"Display,omitempty" xml:"Display,omitempty"`
	Height  *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Url     *string  `json:"Url,omitempty" xml:"Url,omitempty"`
	Width   *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	X       *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y       *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	ZOrder  *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s AddRecordTemplateRequestWatermarks) String() string {
	return tea.Prettify(s)
}

func (s AddRecordTemplateRequestWatermarks) GoString() string {
	return s.String()
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

type AddRecordTemplateResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s AddRecordTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddRecordTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *AddRecordTemplateResponseBody) SetRequestId(v string) *AddRecordTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddRecordTemplateResponseBody) SetTemplateId(v string) *AddRecordTemplateResponseBody {
	s.TemplateId = &v
	return s
}

type AddRecordTemplateResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddRecordTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddRecordTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s AddRecordTemplateResponse) GoString() string {
	return s.String()
}

func (s *AddRecordTemplateResponse) SetHeaders(v map[string]*string) *AddRecordTemplateResponse {
	s.Headers = v
	return s
}

func (s *AddRecordTemplateResponse) SetBody(v *AddRecordTemplateResponseBody) *AddRecordTemplateResponse {
	s.Body = v
	return s
}

type CreateAutoLiveStreamRuleRequest struct {
	AppId             *string   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CallBack          *string   `json:"CallBack,omitempty" xml:"CallBack,omitempty"`
	ChannelIdPrefixes []*string `json:"ChannelIdPrefixes,omitempty" xml:"ChannelIdPrefixes,omitempty" type:"Repeated"`
	ChannelIds        []*string `json:"ChannelIds,omitempty" xml:"ChannelIds,omitempty" type:"Repeated"`
	MediaEncode       *int32    `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	OwnerId           *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PlayDomain        *string   `json:"PlayDomain,omitempty" xml:"PlayDomain,omitempty"`
	RuleName          *string   `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s CreateAutoLiveStreamRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAutoLiveStreamRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateAutoLiveStreamRuleRequest) SetAppId(v string) *CreateAutoLiveStreamRuleRequest {
	s.AppId = &v
	return s
}

func (s *CreateAutoLiveStreamRuleRequest) SetCallBack(v string) *CreateAutoLiveStreamRuleRequest {
	s.CallBack = &v
	return s
}

func (s *CreateAutoLiveStreamRuleRequest) SetChannelIdPrefixes(v []*string) *CreateAutoLiveStreamRuleRequest {
	s.ChannelIdPrefixes = v
	return s
}

func (s *CreateAutoLiveStreamRuleRequest) SetChannelIds(v []*string) *CreateAutoLiveStreamRuleRequest {
	s.ChannelIds = v
	return s
}

func (s *CreateAutoLiveStreamRuleRequest) SetMediaEncode(v int32) *CreateAutoLiveStreamRuleRequest {
	s.MediaEncode = &v
	return s
}

func (s *CreateAutoLiveStreamRuleRequest) SetOwnerId(v int64) *CreateAutoLiveStreamRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAutoLiveStreamRuleRequest) SetPlayDomain(v string) *CreateAutoLiveStreamRuleRequest {
	s.PlayDomain = &v
	return s
}

func (s *CreateAutoLiveStreamRuleRequest) SetRuleName(v string) *CreateAutoLiveStreamRuleRequest {
	s.RuleName = &v
	return s
}

type CreateAutoLiveStreamRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RuleId    *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s CreateAutoLiveStreamRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAutoLiveStreamRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAutoLiveStreamRuleResponseBody) SetRequestId(v string) *CreateAutoLiveStreamRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAutoLiveStreamRuleResponseBody) SetRuleId(v int64) *CreateAutoLiveStreamRuleResponseBody {
	s.RuleId = &v
	return s
}

type CreateAutoLiveStreamRuleResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAutoLiveStreamRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAutoLiveStreamRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAutoLiveStreamRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateAutoLiveStreamRuleResponse) SetHeaders(v map[string]*string) *CreateAutoLiveStreamRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateAutoLiveStreamRuleResponse) SetBody(v *CreateAutoLiveStreamRuleResponseBody) *CreateAutoLiveStreamRuleResponse {
	s.Body = v
	return s
}

type CreateEventSubscribeRequest struct {
	AppId       *string   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CallbackUrl *string   `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	ChannelId   *string   `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ClientToken *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Events      []*string `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	OwnerId     *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Users       []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s CreateEventSubscribeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEventSubscribeRequest) GoString() string {
	return s.String()
}

func (s *CreateEventSubscribeRequest) SetAppId(v string) *CreateEventSubscribeRequest {
	s.AppId = &v
	return s
}

func (s *CreateEventSubscribeRequest) SetCallbackUrl(v string) *CreateEventSubscribeRequest {
	s.CallbackUrl = &v
	return s
}

func (s *CreateEventSubscribeRequest) SetChannelId(v string) *CreateEventSubscribeRequest {
	s.ChannelId = &v
	return s
}

func (s *CreateEventSubscribeRequest) SetClientToken(v string) *CreateEventSubscribeRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEventSubscribeRequest) SetEvents(v []*string) *CreateEventSubscribeRequest {
	s.Events = v
	return s
}

func (s *CreateEventSubscribeRequest) SetOwnerId(v int64) *CreateEventSubscribeRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateEventSubscribeRequest) SetUsers(v []*string) *CreateEventSubscribeRequest {
	s.Users = v
	return s
}

type CreateEventSubscribeResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubscribeId *string `json:"SubscribeId,omitempty" xml:"SubscribeId,omitempty"`
}

func (s CreateEventSubscribeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEventSubscribeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEventSubscribeResponseBody) SetRequestId(v string) *CreateEventSubscribeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEventSubscribeResponseBody) SetSubscribeId(v string) *CreateEventSubscribeResponseBody {
	s.SubscribeId = &v
	return s
}

type CreateEventSubscribeResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateEventSubscribeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateEventSubscribeResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEventSubscribeResponse) GoString() string {
	return s.String()
}

func (s *CreateEventSubscribeResponse) SetHeaders(v map[string]*string) *CreateEventSubscribeResponse {
	s.Headers = v
	return s
}

func (s *CreateEventSubscribeResponse) SetBody(v *CreateEventSubscribeResponseBody) *CreateEventSubscribeResponse {
	s.Body = v
	return s
}

type CreateMPULayoutRequest struct {
	AppId         *string                        `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AudioMixCount *int32                         `json:"AudioMixCount,omitempty" xml:"AudioMixCount,omitempty"`
	Name          *string                        `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId       *int64                         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Panes         []*CreateMPULayoutRequestPanes `json:"Panes,omitempty" xml:"Panes,omitempty" type:"Repeated"`
}

func (s CreateMPULayoutRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMPULayoutRequest) GoString() string {
	return s.String()
}

func (s *CreateMPULayoutRequest) SetAppId(v string) *CreateMPULayoutRequest {
	s.AppId = &v
	return s
}

func (s *CreateMPULayoutRequest) SetAudioMixCount(v int32) *CreateMPULayoutRequest {
	s.AudioMixCount = &v
	return s
}

func (s *CreateMPULayoutRequest) SetName(v string) *CreateMPULayoutRequest {
	s.Name = &v
	return s
}

func (s *CreateMPULayoutRequest) SetOwnerId(v int64) *CreateMPULayoutRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateMPULayoutRequest) SetPanes(v []*CreateMPULayoutRequestPanes) *CreateMPULayoutRequest {
	s.Panes = v
	return s
}

type CreateMPULayoutRequestPanes struct {
	Height    *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	MajorPane *int32   `json:"MajorPane,omitempty" xml:"MajorPane,omitempty"`
	PaneId    *int32   `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
	Width     *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	X         *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y         *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	ZOrder    *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s CreateMPULayoutRequestPanes) String() string {
	return tea.Prettify(s)
}

func (s CreateMPULayoutRequestPanes) GoString() string {
	return s.String()
}

func (s *CreateMPULayoutRequestPanes) SetHeight(v float32) *CreateMPULayoutRequestPanes {
	s.Height = &v
	return s
}

func (s *CreateMPULayoutRequestPanes) SetMajorPane(v int32) *CreateMPULayoutRequestPanes {
	s.MajorPane = &v
	return s
}

func (s *CreateMPULayoutRequestPanes) SetPaneId(v int32) *CreateMPULayoutRequestPanes {
	s.PaneId = &v
	return s
}

func (s *CreateMPULayoutRequestPanes) SetWidth(v float32) *CreateMPULayoutRequestPanes {
	s.Width = &v
	return s
}

func (s *CreateMPULayoutRequestPanes) SetX(v float32) *CreateMPULayoutRequestPanes {
	s.X = &v
	return s
}

func (s *CreateMPULayoutRequestPanes) SetY(v float32) *CreateMPULayoutRequestPanes {
	s.Y = &v
	return s
}

func (s *CreateMPULayoutRequestPanes) SetZOrder(v int32) *CreateMPULayoutRequestPanes {
	s.ZOrder = &v
	return s
}

type CreateMPULayoutResponseBody struct {
	LayoutId  *int64  `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMPULayoutResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMPULayoutResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMPULayoutResponseBody) SetLayoutId(v int64) *CreateMPULayoutResponseBody {
	s.LayoutId = &v
	return s
}

func (s *CreateMPULayoutResponseBody) SetRequestId(v string) *CreateMPULayoutResponseBody {
	s.RequestId = &v
	return s
}

type CreateMPULayoutResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateMPULayoutResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMPULayoutResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMPULayoutResponse) GoString() string {
	return s.String()
}

func (s *CreateMPULayoutResponse) SetHeaders(v map[string]*string) *CreateMPULayoutResponse {
	s.Headers = v
	return s
}

func (s *CreateMPULayoutResponse) SetBody(v *CreateMPULayoutResponseBody) *CreateMPULayoutResponse {
	s.Body = v
	return s
}

type CreateRecordIndexFileRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId   *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	EndTime     *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OssBucket   *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssEndpoint *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	OssObject   *string `json:"OssObject,omitempty" xml:"OssObject,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateRecordIndexFileRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRecordIndexFileRequest) GoString() string {
	return s.String()
}

func (s *CreateRecordIndexFileRequest) SetAppId(v string) *CreateRecordIndexFileRequest {
	s.AppId = &v
	return s
}

func (s *CreateRecordIndexFileRequest) SetChannelId(v string) *CreateRecordIndexFileRequest {
	s.ChannelId = &v
	return s
}

func (s *CreateRecordIndexFileRequest) SetEndTime(v string) *CreateRecordIndexFileRequest {
	s.EndTime = &v
	return s
}

func (s *CreateRecordIndexFileRequest) SetOssBucket(v string) *CreateRecordIndexFileRequest {
	s.OssBucket = &v
	return s
}

func (s *CreateRecordIndexFileRequest) SetOssEndpoint(v string) *CreateRecordIndexFileRequest {
	s.OssEndpoint = &v
	return s
}

func (s *CreateRecordIndexFileRequest) SetOssObject(v string) *CreateRecordIndexFileRequest {
	s.OssObject = &v
	return s
}

func (s *CreateRecordIndexFileRequest) SetOwnerId(v int64) *CreateRecordIndexFileRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateRecordIndexFileRequest) SetStartTime(v string) *CreateRecordIndexFileRequest {
	s.StartTime = &v
	return s
}

func (s *CreateRecordIndexFileRequest) SetTaskId(v string) *CreateRecordIndexFileRequest {
	s.TaskId = &v
	return s
}

type CreateRecordIndexFileResponseBody struct {
	Duration  *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Height    *int32   `json:"Height,omitempty" xml:"Height,omitempty"`
	RequestId *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Url       *string  `json:"Url,omitempty" xml:"Url,omitempty"`
	Width     *int32   `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s CreateRecordIndexFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRecordIndexFileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRecordIndexFileResponseBody) SetDuration(v float32) *CreateRecordIndexFileResponseBody {
	s.Duration = &v
	return s
}

func (s *CreateRecordIndexFileResponseBody) SetHeight(v int32) *CreateRecordIndexFileResponseBody {
	s.Height = &v
	return s
}

func (s *CreateRecordIndexFileResponseBody) SetRequestId(v string) *CreateRecordIndexFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRecordIndexFileResponseBody) SetUrl(v string) *CreateRecordIndexFileResponseBody {
	s.Url = &v
	return s
}

func (s *CreateRecordIndexFileResponseBody) SetWidth(v int32) *CreateRecordIndexFileResponseBody {
	s.Width = &v
	return s
}

type CreateRecordIndexFileResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateRecordIndexFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRecordIndexFileResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRecordIndexFileResponse) GoString() string {
	return s.String()
}

func (s *CreateRecordIndexFileResponse) SetHeaders(v map[string]*string) *CreateRecordIndexFileResponse {
	s.Headers = v
	return s
}

func (s *CreateRecordIndexFileResponse) SetBody(v *CreateRecordIndexFileResponseBody) *CreateRecordIndexFileResponse {
	s.Body = v
	return s
}

type DeleteAutoLiveStreamRuleRequest struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RuleId  *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteAutoLiveStreamRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAutoLiveStreamRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteAutoLiveStreamRuleRequest) SetAppId(v string) *DeleteAutoLiveStreamRuleRequest {
	s.AppId = &v
	return s
}

func (s *DeleteAutoLiveStreamRuleRequest) SetOwnerId(v int64) *DeleteAutoLiveStreamRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteAutoLiveStreamRuleRequest) SetRuleId(v int64) *DeleteAutoLiveStreamRuleRequest {
	s.RuleId = &v
	return s
}

type DeleteAutoLiveStreamRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAutoLiveStreamRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAutoLiveStreamRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAutoLiveStreamRuleResponseBody) SetRequestId(v string) *DeleteAutoLiveStreamRuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAutoLiveStreamRuleResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAutoLiveStreamRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAutoLiveStreamRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAutoLiveStreamRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteAutoLiveStreamRuleResponse) SetHeaders(v map[string]*string) *DeleteAutoLiveStreamRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteAutoLiveStreamRuleResponse) SetBody(v *DeleteAutoLiveStreamRuleResponseBody) *DeleteAutoLiveStreamRuleResponse {
	s.Body = v
	return s
}

type DeleteEventSubscribeRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SubscribeId *string `json:"SubscribeId,omitempty" xml:"SubscribeId,omitempty"`
}

func (s DeleteEventSubscribeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventSubscribeRequest) GoString() string {
	return s.String()
}

func (s *DeleteEventSubscribeRequest) SetAppId(v string) *DeleteEventSubscribeRequest {
	s.AppId = &v
	return s
}

func (s *DeleteEventSubscribeRequest) SetOwnerId(v int64) *DeleteEventSubscribeRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteEventSubscribeRequest) SetSubscribeId(v string) *DeleteEventSubscribeRequest {
	s.SubscribeId = &v
	return s
}

type DeleteEventSubscribeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEventSubscribeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventSubscribeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEventSubscribeResponseBody) SetRequestId(v string) *DeleteEventSubscribeResponseBody {
	s.RequestId = &v
	return s
}

type DeleteEventSubscribeResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteEventSubscribeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteEventSubscribeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventSubscribeResponse) GoString() string {
	return s.String()
}

func (s *DeleteEventSubscribeResponse) SetHeaders(v map[string]*string) *DeleteEventSubscribeResponse {
	s.Headers = v
	return s
}

func (s *DeleteEventSubscribeResponse) SetBody(v *DeleteEventSubscribeResponseBody) *DeleteEventSubscribeResponse {
	s.Body = v
	return s
}

type DeleteMPULayoutRequest struct {
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	LayoutId *int64  `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DeleteMPULayoutRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMPULayoutRequest) GoString() string {
	return s.String()
}

func (s *DeleteMPULayoutRequest) SetAppId(v string) *DeleteMPULayoutRequest {
	s.AppId = &v
	return s
}

func (s *DeleteMPULayoutRequest) SetLayoutId(v int64) *DeleteMPULayoutRequest {
	s.LayoutId = &v
	return s
}

func (s *DeleteMPULayoutRequest) SetOwnerId(v int64) *DeleteMPULayoutRequest {
	s.OwnerId = &v
	return s
}

type DeleteMPULayoutResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMPULayoutResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMPULayoutResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMPULayoutResponseBody) SetRequestId(v string) *DeleteMPULayoutResponseBody {
	s.RequestId = &v
	return s
}

type DeleteMPULayoutResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteMPULayoutResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMPULayoutResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMPULayoutResponse) GoString() string {
	return s.String()
}

func (s *DeleteMPULayoutResponse) SetHeaders(v map[string]*string) *DeleteMPULayoutResponse {
	s.Headers = v
	return s
}

func (s *DeleteMPULayoutResponse) SetBody(v *DeleteMPULayoutResponseBody) *DeleteMPULayoutResponse {
	s.Body = v
	return s
}

type DeleteRecordTemplateRequest struct {
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteRecordTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRecordTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteRecordTemplateRequest) SetAppId(v string) *DeleteRecordTemplateRequest {
	s.AppId = &v
	return s
}

func (s *DeleteRecordTemplateRequest) SetOwnerId(v int64) *DeleteRecordTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteRecordTemplateRequest) SetTemplateId(v string) *DeleteRecordTemplateRequest {
	s.TemplateId = &v
	return s
}

type DeleteRecordTemplateResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRecordTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRecordTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRecordTemplateResponseBody) SetRequestId(v string) *DeleteRecordTemplateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteRecordTemplateResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteRecordTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRecordTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRecordTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteRecordTemplateResponse) SetHeaders(v map[string]*string) *DeleteRecordTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteRecordTemplateResponse) SetBody(v *DeleteRecordTemplateResponseBody) *DeleteRecordTemplateResponse {
	s.Body = v
	return s
}

type DescribeAutoLiveStreamRuleRequest struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeAutoLiveStreamRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoLiveStreamRuleRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutoLiveStreamRuleRequest) SetAppId(v string) *DescribeAutoLiveStreamRuleRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAutoLiveStreamRuleRequest) SetOwnerId(v int64) *DescribeAutoLiveStreamRuleRequest {
	s.OwnerId = &v
	return s
}

type DescribeAutoLiveStreamRuleResponseBody struct {
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rules     []*DescribeAutoLiveStreamRuleResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeAutoLiveStreamRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoLiveStreamRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutoLiveStreamRuleResponseBody) SetRequestId(v string) *DescribeAutoLiveStreamRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutoLiveStreamRuleResponseBody) SetRules(v []*DescribeAutoLiveStreamRuleResponseBodyRules) *DescribeAutoLiveStreamRuleResponseBody {
	s.Rules = v
	return s
}

type DescribeAutoLiveStreamRuleResponseBodyRules struct {
	CallBack          *string   `json:"CallBack,omitempty" xml:"CallBack,omitempty"`
	ChannelIdPrefixes []*string `json:"ChannelIdPrefixes,omitempty" xml:"ChannelIdPrefixes,omitempty" type:"Repeated"`
	ChannelIds        []*string `json:"ChannelIds,omitempty" xml:"ChannelIds,omitempty" type:"Repeated"`
	CreateTime        *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	MediaEncode       *int32    `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	PlayDomain        *string   `json:"PlayDomain,omitempty" xml:"PlayDomain,omitempty"`
	RuleId            *int64    `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName          *string   `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Status            *string   `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeAutoLiveStreamRuleResponseBodyRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoLiveStreamRuleResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribeAutoLiveStreamRuleResponseBodyRules) SetCallBack(v string) *DescribeAutoLiveStreamRuleResponseBodyRules {
	s.CallBack = &v
	return s
}

func (s *DescribeAutoLiveStreamRuleResponseBodyRules) SetChannelIdPrefixes(v []*string) *DescribeAutoLiveStreamRuleResponseBodyRules {
	s.ChannelIdPrefixes = v
	return s
}

func (s *DescribeAutoLiveStreamRuleResponseBodyRules) SetChannelIds(v []*string) *DescribeAutoLiveStreamRuleResponseBodyRules {
	s.ChannelIds = v
	return s
}

func (s *DescribeAutoLiveStreamRuleResponseBodyRules) SetCreateTime(v string) *DescribeAutoLiveStreamRuleResponseBodyRules {
	s.CreateTime = &v
	return s
}

func (s *DescribeAutoLiveStreamRuleResponseBodyRules) SetMediaEncode(v int32) *DescribeAutoLiveStreamRuleResponseBodyRules {
	s.MediaEncode = &v
	return s
}

func (s *DescribeAutoLiveStreamRuleResponseBodyRules) SetPlayDomain(v string) *DescribeAutoLiveStreamRuleResponseBodyRules {
	s.PlayDomain = &v
	return s
}

func (s *DescribeAutoLiveStreamRuleResponseBodyRules) SetRuleId(v int64) *DescribeAutoLiveStreamRuleResponseBodyRules {
	s.RuleId = &v
	return s
}

func (s *DescribeAutoLiveStreamRuleResponseBodyRules) SetRuleName(v string) *DescribeAutoLiveStreamRuleResponseBodyRules {
	s.RuleName = &v
	return s
}

func (s *DescribeAutoLiveStreamRuleResponseBodyRules) SetStatus(v string) *DescribeAutoLiveStreamRuleResponseBodyRules {
	s.Status = &v
	return s
}

type DescribeAutoLiveStreamRuleResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAutoLiveStreamRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAutoLiveStreamRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoLiveStreamRuleResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutoLiveStreamRuleResponse) SetHeaders(v map[string]*string) *DescribeAutoLiveStreamRuleResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutoLiveStreamRuleResponse) SetBody(v *DescribeAutoLiveStreamRuleResponseBody) *DescribeAutoLiveStreamRuleResponse {
	s.Body = v
	return s
}

type DescribeChannelParticipantsRequest struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	Order     *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNum   *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeChannelParticipantsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelParticipantsRequest) GoString() string {
	return s.String()
}

func (s *DescribeChannelParticipantsRequest) SetAppId(v string) *DescribeChannelParticipantsRequest {
	s.AppId = &v
	return s
}

func (s *DescribeChannelParticipantsRequest) SetChannelId(v string) *DescribeChannelParticipantsRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeChannelParticipantsRequest) SetOrder(v string) *DescribeChannelParticipantsRequest {
	s.Order = &v
	return s
}

func (s *DescribeChannelParticipantsRequest) SetOwnerId(v int64) *DescribeChannelParticipantsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeChannelParticipantsRequest) SetPageNum(v int32) *DescribeChannelParticipantsRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeChannelParticipantsRequest) SetPageSize(v int32) *DescribeChannelParticipantsRequest {
	s.PageSize = &v
	return s
}

type DescribeChannelParticipantsResponseBody struct {
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Timestamp *int32                                           `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	TotalNum  *int32                                           `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	TotalPage *int32                                           `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	UserList  *DescribeChannelParticipantsResponseBodyUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Struct"`
}

func (s DescribeChannelParticipantsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelParticipantsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChannelParticipantsResponseBody) SetRequestId(v string) *DescribeChannelParticipantsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeChannelParticipantsResponseBody) SetTimestamp(v int32) *DescribeChannelParticipantsResponseBody {
	s.Timestamp = &v
	return s
}

func (s *DescribeChannelParticipantsResponseBody) SetTotalNum(v int32) *DescribeChannelParticipantsResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeChannelParticipantsResponseBody) SetTotalPage(v int32) *DescribeChannelParticipantsResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeChannelParticipantsResponseBody) SetUserList(v *DescribeChannelParticipantsResponseBodyUserList) *DescribeChannelParticipantsResponseBody {
	s.UserList = v
	return s
}

type DescribeChannelParticipantsResponseBodyUserList struct {
	User []*string `json:"User,omitempty" xml:"User,omitempty" type:"Repeated"`
}

func (s DescribeChannelParticipantsResponseBodyUserList) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelParticipantsResponseBodyUserList) GoString() string {
	return s.String()
}

func (s *DescribeChannelParticipantsResponseBodyUserList) SetUser(v []*string) *DescribeChannelParticipantsResponseBodyUserList {
	s.User = v
	return s
}

type DescribeChannelParticipantsResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeChannelParticipantsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeChannelParticipantsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelParticipantsResponse) GoString() string {
	return s.String()
}

func (s *DescribeChannelParticipantsResponse) SetHeaders(v map[string]*string) *DescribeChannelParticipantsResponse {
	s.Headers = v
	return s
}

func (s *DescribeChannelParticipantsResponse) SetBody(v *DescribeChannelParticipantsResponseBody) *DescribeChannelParticipantsResponse {
	s.Body = v
	return s
}

type DescribeChannelUsersRequest struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeChannelUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelUsersRequest) GoString() string {
	return s.String()
}

func (s *DescribeChannelUsersRequest) SetAppId(v string) *DescribeChannelUsersRequest {
	s.AppId = &v
	return s
}

func (s *DescribeChannelUsersRequest) SetChannelId(v string) *DescribeChannelUsersRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeChannelUsersRequest) SetOwnerId(v int64) *DescribeChannelUsersRequest {
	s.OwnerId = &v
	return s
}

type DescribeChannelUsersResponseBody struct {
	ChannelProfile      *int32    `json:"ChannelProfile,omitempty" xml:"ChannelProfile,omitempty"`
	CommTotalNum        *int32    `json:"CommTotalNum,omitempty" xml:"CommTotalNum,omitempty"`
	InteractiveUserList []*string `json:"InteractiveUserList,omitempty" xml:"InteractiveUserList,omitempty" type:"Repeated"`
	InteractiveUserNum  *int32    `json:"InteractiveUserNum,omitempty" xml:"InteractiveUserNum,omitempty"`
	IsChannelExist      *bool     `json:"IsChannelExist,omitempty" xml:"IsChannelExist,omitempty"`
	LiveUserList        []*string `json:"LiveUserList,omitempty" xml:"LiveUserList,omitempty" type:"Repeated"`
	LiveUserNum         *int32    `json:"LiveUserNum,omitempty" xml:"LiveUserNum,omitempty"`
	RequestId           *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Timestamp           *int32    `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	UserList            []*string `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
}

func (s DescribeChannelUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChannelUsersResponseBody) SetChannelProfile(v int32) *DescribeChannelUsersResponseBody {
	s.ChannelProfile = &v
	return s
}

func (s *DescribeChannelUsersResponseBody) SetCommTotalNum(v int32) *DescribeChannelUsersResponseBody {
	s.CommTotalNum = &v
	return s
}

func (s *DescribeChannelUsersResponseBody) SetInteractiveUserList(v []*string) *DescribeChannelUsersResponseBody {
	s.InteractiveUserList = v
	return s
}

func (s *DescribeChannelUsersResponseBody) SetInteractiveUserNum(v int32) *DescribeChannelUsersResponseBody {
	s.InteractiveUserNum = &v
	return s
}

func (s *DescribeChannelUsersResponseBody) SetIsChannelExist(v bool) *DescribeChannelUsersResponseBody {
	s.IsChannelExist = &v
	return s
}

func (s *DescribeChannelUsersResponseBody) SetLiveUserList(v []*string) *DescribeChannelUsersResponseBody {
	s.LiveUserList = v
	return s
}

func (s *DescribeChannelUsersResponseBody) SetLiveUserNum(v int32) *DescribeChannelUsersResponseBody {
	s.LiveUserNum = &v
	return s
}

func (s *DescribeChannelUsersResponseBody) SetRequestId(v string) *DescribeChannelUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeChannelUsersResponseBody) SetTimestamp(v int32) *DescribeChannelUsersResponseBody {
	s.Timestamp = &v
	return s
}

func (s *DescribeChannelUsersResponseBody) SetUserList(v []*string) *DescribeChannelUsersResponseBody {
	s.UserList = v
	return s
}

type DescribeChannelUsersResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeChannelUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeChannelUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelUsersResponse) GoString() string {
	return s.String()
}

func (s *DescribeChannelUsersResponse) SetHeaders(v map[string]*string) *DescribeChannelUsersResponse {
	s.Headers = v
	return s
}

func (s *DescribeChannelUsersResponse) SetBody(v *DescribeChannelUsersResponseBody) *DescribeChannelUsersResponse {
	s.Body = v
	return s
}

type DescribeMPULayoutInfoListRequest struct {
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	LayoutId *int64  `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNum  *int64  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeMPULayoutInfoListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPULayoutInfoListRequest) GoString() string {
	return s.String()
}

func (s *DescribeMPULayoutInfoListRequest) SetAppId(v string) *DescribeMPULayoutInfoListRequest {
	s.AppId = &v
	return s
}

func (s *DescribeMPULayoutInfoListRequest) SetLayoutId(v int64) *DescribeMPULayoutInfoListRequest {
	s.LayoutId = &v
	return s
}

func (s *DescribeMPULayoutInfoListRequest) SetName(v string) *DescribeMPULayoutInfoListRequest {
	s.Name = &v
	return s
}

func (s *DescribeMPULayoutInfoListRequest) SetOwnerId(v int64) *DescribeMPULayoutInfoListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMPULayoutInfoListRequest) SetPageNum(v int64) *DescribeMPULayoutInfoListRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeMPULayoutInfoListRequest) SetPageSize(v int64) *DescribeMPULayoutInfoListRequest {
	s.PageSize = &v
	return s
}

type DescribeMPULayoutInfoListResponseBody struct {
	Layouts   *DescribeMPULayoutInfoListResponseBodyLayouts `json:"Layouts,omitempty" xml:"Layouts,omitempty" type:"Struct"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalNum  *int64                                        `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	TotalPage *int64                                        `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeMPULayoutInfoListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPULayoutInfoListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMPULayoutInfoListResponseBody) SetLayouts(v *DescribeMPULayoutInfoListResponseBodyLayouts) *DescribeMPULayoutInfoListResponseBody {
	s.Layouts = v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBody) SetRequestId(v string) *DescribeMPULayoutInfoListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBody) SetTotalNum(v int64) *DescribeMPULayoutInfoListResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBody) SetTotalPage(v int64) *DescribeMPULayoutInfoListResponseBody {
	s.TotalPage = &v
	return s
}

type DescribeMPULayoutInfoListResponseBodyLayouts struct {
	Layout []*DescribeMPULayoutInfoListResponseBodyLayoutsLayout `json:"Layout,omitempty" xml:"Layout,omitempty" type:"Repeated"`
}

func (s DescribeMPULayoutInfoListResponseBodyLayouts) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPULayoutInfoListResponseBodyLayouts) GoString() string {
	return s.String()
}

func (s *DescribeMPULayoutInfoListResponseBodyLayouts) SetLayout(v []*DescribeMPULayoutInfoListResponseBodyLayoutsLayout) *DescribeMPULayoutInfoListResponseBodyLayouts {
	s.Layout = v
	return s
}

type DescribeMPULayoutInfoListResponseBodyLayoutsLayout struct {
	AudioMixCount *int32                                                   `json:"AudioMixCount,omitempty" xml:"AudioMixCount,omitempty"`
	LayoutId      *int64                                                   `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	Name          *string                                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	Panes         *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanes `json:"Panes,omitempty" xml:"Panes,omitempty" type:"Struct"`
}

func (s DescribeMPULayoutInfoListResponseBodyLayoutsLayout) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPULayoutInfoListResponseBodyLayoutsLayout) GoString() string {
	return s.String()
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayout) SetAudioMixCount(v int32) *DescribeMPULayoutInfoListResponseBodyLayoutsLayout {
	s.AudioMixCount = &v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayout) SetLayoutId(v int64) *DescribeMPULayoutInfoListResponseBodyLayoutsLayout {
	s.LayoutId = &v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayout) SetName(v string) *DescribeMPULayoutInfoListResponseBodyLayoutsLayout {
	s.Name = &v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayout) SetPanes(v *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanes) *DescribeMPULayoutInfoListResponseBodyLayoutsLayout {
	s.Panes = v
	return s
}

type DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanes struct {
	Panes []*DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes `json:"Panes,omitempty" xml:"Panes,omitempty" type:"Repeated"`
}

func (s DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanes) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanes) GoString() string {
	return s.String()
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanes) SetPanes(v []*DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes) *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanes {
	s.Panes = v
	return s
}

type DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes struct {
	Height    *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	MajorPane *int32   `json:"MajorPane,omitempty" xml:"MajorPane,omitempty"`
	PaneId    *int32   `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
	Width     *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	X         *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y         *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	ZOrder    *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes) GoString() string {
	return s.String()
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes) SetHeight(v float32) *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes {
	s.Height = &v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes) SetMajorPane(v int32) *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes {
	s.MajorPane = &v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes) SetPaneId(v int32) *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes {
	s.PaneId = &v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes) SetWidth(v float32) *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes {
	s.Width = &v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes) SetX(v float32) *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes {
	s.X = &v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes) SetY(v float32) *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes {
	s.Y = &v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes) SetZOrder(v int32) *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes {
	s.ZOrder = &v
	return s
}

type DescribeMPULayoutInfoListResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMPULayoutInfoListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMPULayoutInfoListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPULayoutInfoListResponse) GoString() string {
	return s.String()
}

func (s *DescribeMPULayoutInfoListResponse) SetHeaders(v map[string]*string) *DescribeMPULayoutInfoListResponse {
	s.Headers = v
	return s
}

func (s *DescribeMPULayoutInfoListResponse) SetBody(v *DescribeMPULayoutInfoListResponseBody) *DescribeMPULayoutInfoListResponse {
	s.Body = v
	return s
}

type DescribeRecordFilesRequest struct {
	AppId     *string   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId *string   `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	EndTime   *string   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId   *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNum   *int32    `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize  *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime *string   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TaskIds   []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
}

func (s DescribeRecordFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordFilesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecordFilesRequest) SetAppId(v string) *DescribeRecordFilesRequest {
	s.AppId = &v
	return s
}

func (s *DescribeRecordFilesRequest) SetChannelId(v string) *DescribeRecordFilesRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeRecordFilesRequest) SetEndTime(v string) *DescribeRecordFilesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRecordFilesRequest) SetOwnerId(v int64) *DescribeRecordFilesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRecordFilesRequest) SetPageNum(v int32) *DescribeRecordFilesRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeRecordFilesRequest) SetPageSize(v int32) *DescribeRecordFilesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRecordFilesRequest) SetStartTime(v string) *DescribeRecordFilesRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRecordFilesRequest) SetTaskIds(v []*string) *DescribeRecordFilesRequest {
	s.TaskIds = v
	return s
}

type DescribeRecordFilesResponseBody struct {
	RecordFiles []*DescribeRecordFilesResponseBodyRecordFiles `json:"RecordFiles,omitempty" xml:"RecordFiles,omitempty" type:"Repeated"`
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalNum    *int64                                        `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	TotalPage   *int64                                        `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeRecordFilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordFilesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecordFilesResponseBody) SetRecordFiles(v []*DescribeRecordFilesResponseBodyRecordFiles) *DescribeRecordFilesResponseBody {
	s.RecordFiles = v
	return s
}

func (s *DescribeRecordFilesResponseBody) SetRequestId(v string) *DescribeRecordFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecordFilesResponseBody) SetTotalNum(v int64) *DescribeRecordFilesResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeRecordFilesResponseBody) SetTotalPage(v int64) *DescribeRecordFilesResponseBody {
	s.TotalPage = &v
	return s
}

type DescribeRecordFilesResponseBodyRecordFiles struct {
	AppId      *string  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId  *string  `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	CreateTime *string  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Duration   *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	StartTime  *string  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StopTime   *string  `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
	TaskId     *string  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Url        *string  `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s DescribeRecordFilesResponseBodyRecordFiles) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordFilesResponseBodyRecordFiles) GoString() string {
	return s.String()
}

func (s *DescribeRecordFilesResponseBodyRecordFiles) SetAppId(v string) *DescribeRecordFilesResponseBodyRecordFiles {
	s.AppId = &v
	return s
}

func (s *DescribeRecordFilesResponseBodyRecordFiles) SetChannelId(v string) *DescribeRecordFilesResponseBodyRecordFiles {
	s.ChannelId = &v
	return s
}

func (s *DescribeRecordFilesResponseBodyRecordFiles) SetCreateTime(v string) *DescribeRecordFilesResponseBodyRecordFiles {
	s.CreateTime = &v
	return s
}

func (s *DescribeRecordFilesResponseBodyRecordFiles) SetDuration(v float32) *DescribeRecordFilesResponseBodyRecordFiles {
	s.Duration = &v
	return s
}

func (s *DescribeRecordFilesResponseBodyRecordFiles) SetStartTime(v string) *DescribeRecordFilesResponseBodyRecordFiles {
	s.StartTime = &v
	return s
}

func (s *DescribeRecordFilesResponseBodyRecordFiles) SetStopTime(v string) *DescribeRecordFilesResponseBodyRecordFiles {
	s.StopTime = &v
	return s
}

func (s *DescribeRecordFilesResponseBodyRecordFiles) SetTaskId(v string) *DescribeRecordFilesResponseBodyRecordFiles {
	s.TaskId = &v
	return s
}

func (s *DescribeRecordFilesResponseBodyRecordFiles) SetUrl(v string) *DescribeRecordFilesResponseBodyRecordFiles {
	s.Url = &v
	return s
}

type DescribeRecordFilesResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRecordFilesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRecordFilesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordFilesResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecordFilesResponse) SetHeaders(v map[string]*string) *DescribeRecordFilesResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecordFilesResponse) SetBody(v *DescribeRecordFilesResponseBody) *DescribeRecordFilesResponse {
	s.Body = v
	return s
}

type DescribeRecordTasksRequest struct {
	AppId     *string   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId *string   `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	EndTime   *string   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId   *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNum   *int32    `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize  *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime *string   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status    *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskIds   []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
}

func (s DescribeRecordTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecordTasksRequest) SetAppId(v string) *DescribeRecordTasksRequest {
	s.AppId = &v
	return s
}

func (s *DescribeRecordTasksRequest) SetChannelId(v string) *DescribeRecordTasksRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeRecordTasksRequest) SetEndTime(v string) *DescribeRecordTasksRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRecordTasksRequest) SetOwnerId(v int64) *DescribeRecordTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRecordTasksRequest) SetPageNum(v int32) *DescribeRecordTasksRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeRecordTasksRequest) SetPageSize(v int32) *DescribeRecordTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRecordTasksRequest) SetStartTime(v string) *DescribeRecordTasksRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRecordTasksRequest) SetStatus(v string) *DescribeRecordTasksRequest {
	s.Status = &v
	return s
}

func (s *DescribeRecordTasksRequest) SetTaskIds(v []*string) *DescribeRecordTasksRequest {
	s.TaskIds = v
	return s
}

type DescribeRecordTasksResponseBody struct {
	RecordTasks []*DescribeRecordTasksResponseBodyRecordTasks `json:"RecordTasks,omitempty" xml:"RecordTasks,omitempty" type:"Repeated"`
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalNum    *int64                                        `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	TotalPage   *int64                                        `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeRecordTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecordTasksResponseBody) SetRecordTasks(v []*DescribeRecordTasksResponseBodyRecordTasks) *DescribeRecordTasksResponseBody {
	s.RecordTasks = v
	return s
}

func (s *DescribeRecordTasksResponseBody) SetRequestId(v string) *DescribeRecordTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecordTasksResponseBody) SetTotalNum(v int64) *DescribeRecordTasksResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeRecordTasksResponseBody) SetTotalPage(v int64) *DescribeRecordTasksResponseBody {
	s.TotalPage = &v
	return s
}

type DescribeRecordTasksResponseBodyRecordTasks struct {
	AppId        *string                                                `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId    *string                                                `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	CreateTime   *string                                                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Status       *int32                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	SubSpecUsers []*string                                              `json:"SubSpecUsers,omitempty" xml:"SubSpecUsers,omitempty" type:"Repeated"`
	TaskId       *string                                                `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TemplateId   *string                                                `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	UserPanes    []*DescribeRecordTasksResponseBodyRecordTasksUserPanes `json:"UserPanes,omitempty" xml:"UserPanes,omitempty" type:"Repeated"`
}

func (s DescribeRecordTasksResponseBodyRecordTasks) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordTasksResponseBodyRecordTasks) GoString() string {
	return s.String()
}

func (s *DescribeRecordTasksResponseBodyRecordTasks) SetAppId(v string) *DescribeRecordTasksResponseBodyRecordTasks {
	s.AppId = &v
	return s
}

func (s *DescribeRecordTasksResponseBodyRecordTasks) SetChannelId(v string) *DescribeRecordTasksResponseBodyRecordTasks {
	s.ChannelId = &v
	return s
}

func (s *DescribeRecordTasksResponseBodyRecordTasks) SetCreateTime(v string) *DescribeRecordTasksResponseBodyRecordTasks {
	s.CreateTime = &v
	return s
}

func (s *DescribeRecordTasksResponseBodyRecordTasks) SetStatus(v int32) *DescribeRecordTasksResponseBodyRecordTasks {
	s.Status = &v
	return s
}

func (s *DescribeRecordTasksResponseBodyRecordTasks) SetSubSpecUsers(v []*string) *DescribeRecordTasksResponseBodyRecordTasks {
	s.SubSpecUsers = v
	return s
}

func (s *DescribeRecordTasksResponseBodyRecordTasks) SetTaskId(v string) *DescribeRecordTasksResponseBodyRecordTasks {
	s.TaskId = &v
	return s
}

func (s *DescribeRecordTasksResponseBodyRecordTasks) SetTemplateId(v string) *DescribeRecordTasksResponseBodyRecordTasks {
	s.TemplateId = &v
	return s
}

func (s *DescribeRecordTasksResponseBodyRecordTasks) SetUserPanes(v []*DescribeRecordTasksResponseBodyRecordTasksUserPanes) *DescribeRecordTasksResponseBodyRecordTasks {
	s.UserPanes = v
	return s
}

type DescribeRecordTasksResponseBodyRecordTasksUserPanes struct {
	PaneId *int32  `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeRecordTasksResponseBodyRecordTasksUserPanes) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordTasksResponseBodyRecordTasksUserPanes) GoString() string {
	return s.String()
}

func (s *DescribeRecordTasksResponseBodyRecordTasksUserPanes) SetPaneId(v int32) *DescribeRecordTasksResponseBodyRecordTasksUserPanes {
	s.PaneId = &v
	return s
}

func (s *DescribeRecordTasksResponseBodyRecordTasksUserPanes) SetSource(v string) *DescribeRecordTasksResponseBodyRecordTasksUserPanes {
	s.Source = &v
	return s
}

func (s *DescribeRecordTasksResponseBodyRecordTasksUserPanes) SetUserId(v string) *DescribeRecordTasksResponseBodyRecordTasksUserPanes {
	s.UserId = &v
	return s
}

type DescribeRecordTasksResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRecordTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRecordTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecordTasksResponse) SetHeaders(v map[string]*string) *DescribeRecordTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecordTasksResponse) SetBody(v *DescribeRecordTasksResponseBody) *DescribeRecordTasksResponse {
	s.Body = v
	return s
}

type DescribeRecordTemplatesRequest struct {
	AppId       *string   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OwnerId     *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNum     *int32    `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize    *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TemplateIds []*string `json:"TemplateIds,omitempty" xml:"TemplateIds,omitempty" type:"Repeated"`
}

func (s DescribeRecordTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecordTemplatesRequest) SetAppId(v string) *DescribeRecordTemplatesRequest {
	s.AppId = &v
	return s
}

func (s *DescribeRecordTemplatesRequest) SetOwnerId(v int64) *DescribeRecordTemplatesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRecordTemplatesRequest) SetPageNum(v int32) *DescribeRecordTemplatesRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeRecordTemplatesRequest) SetPageSize(v int32) *DescribeRecordTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRecordTemplatesRequest) SetTemplateIds(v []*string) *DescribeRecordTemplatesRequest {
	s.TemplateIds = v
	return s
}

type DescribeRecordTemplatesResponseBody struct {
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Templates []*DescribeRecordTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
	TotalNum  *int64                                          `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	TotalPage *int64                                          `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeRecordTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordTemplatesResponseBody) GoString() string {
	return s.String()
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

type DescribeRecordTemplatesResponseBodyTemplates struct {
	BackgroundColor    *int32                                                      `json:"BackgroundColor,omitempty" xml:"BackgroundColor,omitempty"`
	Backgrounds        []*DescribeRecordTemplatesResponseBodyTemplatesBackgrounds  `json:"Backgrounds,omitempty" xml:"Backgrounds,omitempty" type:"Repeated"`
	ClockWidgets       []*DescribeRecordTemplatesResponseBodyTemplatesClockWidgets `json:"ClockWidgets,omitempty" xml:"ClockWidgets,omitempty" type:"Repeated"`
	CreateTime         *string                                                     `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DelayStopTime      *int32                                                      `json:"DelayStopTime,omitempty" xml:"DelayStopTime,omitempty"`
	EnableM3u8DateTime *bool                                                       `json:"EnableM3u8DateTime,omitempty" xml:"EnableM3u8DateTime,omitempty"`
	FileSplitInterval  *int32                                                      `json:"FileSplitInterval,omitempty" xml:"FileSplitInterval,omitempty"`
	Formats            []*string                                                   `json:"Formats,omitempty" xml:"Formats,omitempty" type:"Repeated"`
	HttpCallbackUrl    *string                                                     `json:"HttpCallbackUrl,omitempty" xml:"HttpCallbackUrl,omitempty"`
	LayoutIds          []*int32                                                    `json:"LayoutIds,omitempty" xml:"LayoutIds,omitempty" type:"Repeated"`
	MediaEncode        *int32                                                      `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	MnsQueue           *string                                                     `json:"MnsQueue,omitempty" xml:"MnsQueue,omitempty"`
	Name               *string                                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	OssBucket          *string                                                     `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssFilePrefix      *string                                                     `json:"OssFilePrefix,omitempty" xml:"OssFilePrefix,omitempty"`
	TaskProfile        *string                                                     `json:"TaskProfile,omitempty" xml:"TaskProfile,omitempty"`
	TemplateId         *string                                                     `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Watermarks         []*DescribeRecordTemplatesResponseBodyTemplatesWatermarks   `json:"Watermarks,omitempty" xml:"Watermarks,omitempty" type:"Repeated"`
}

func (s DescribeRecordTemplatesResponseBodyTemplates) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
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

func (s *DescribeRecordTemplatesResponseBodyTemplates) SetLayoutIds(v []*int32) *DescribeRecordTemplatesResponseBodyTemplates {
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

type DescribeRecordTemplatesResponseBodyTemplatesBackgrounds struct {
	Display *int32   `json:"Display,omitempty" xml:"Display,omitempty"`
	Height  *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Url     *string  `json:"Url,omitempty" xml:"Url,omitempty"`
	Width   *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	X       *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y       *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	ZOrder  *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s DescribeRecordTemplatesResponseBodyTemplatesBackgrounds) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordTemplatesResponseBodyTemplatesBackgrounds) GoString() string {
	return s.String()
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

type DescribeRecordTemplatesResponseBodyTemplatesClockWidgets struct {
	FontColor *int32   `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	FontSize  *int32   `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	FontType  *int32   `json:"FontType,omitempty" xml:"FontType,omitempty"`
	X         *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y         *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	ZOrder    *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s DescribeRecordTemplatesResponseBodyTemplatesClockWidgets) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordTemplatesResponseBodyTemplatesClockWidgets) GoString() string {
	return s.String()
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

type DescribeRecordTemplatesResponseBodyTemplatesWatermarks struct {
	Alpha   *float32 `json:"Alpha,omitempty" xml:"Alpha,omitempty"`
	Display *int32   `json:"Display,omitempty" xml:"Display,omitempty"`
	Height  *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Url     *string  `json:"Url,omitempty" xml:"Url,omitempty"`
	Width   *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	X       *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y       *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	ZOrder  *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s DescribeRecordTemplatesResponseBodyTemplatesWatermarks) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordTemplatesResponseBodyTemplatesWatermarks) GoString() string {
	return s.String()
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

type DescribeRecordTemplatesResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRecordTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRecordTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecordTemplatesResponse) SetHeaders(v map[string]*string) *DescribeRecordTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecordTemplatesResponse) SetBody(v *DescribeRecordTemplatesResponseBody) *DescribeRecordTemplatesResponse {
	s.Body = v
	return s
}

type DescribeUserInfoInChannelRequest struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	UserId    *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeUserInfoInChannelRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserInfoInChannelRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserInfoInChannelRequest) SetAppId(v string) *DescribeUserInfoInChannelRequest {
	s.AppId = &v
	return s
}

func (s *DescribeUserInfoInChannelRequest) SetChannelId(v string) *DescribeUserInfoInChannelRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeUserInfoInChannelRequest) SetOwnerId(v int64) *DescribeUserInfoInChannelRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUserInfoInChannelRequest) SetUserId(v string) *DescribeUserInfoInChannelRequest {
	s.UserId = &v
	return s
}

type DescribeUserInfoInChannelResponseBody struct {
	IsChannelExist *bool                                            `json:"IsChannelExist,omitempty" xml:"IsChannelExist,omitempty"`
	IsInChannel    *bool                                            `json:"IsInChannel,omitempty" xml:"IsInChannel,omitempty"`
	Property       []*DescribeUserInfoInChannelResponseBodyProperty `json:"Property,omitempty" xml:"Property,omitempty" type:"Repeated"`
	RequestId      *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Timestamp      *int32                                           `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeUserInfoInChannelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserInfoInChannelResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserInfoInChannelResponseBody) SetIsChannelExist(v bool) *DescribeUserInfoInChannelResponseBody {
	s.IsChannelExist = &v
	return s
}

func (s *DescribeUserInfoInChannelResponseBody) SetIsInChannel(v bool) *DescribeUserInfoInChannelResponseBody {
	s.IsInChannel = &v
	return s
}

func (s *DescribeUserInfoInChannelResponseBody) SetProperty(v []*DescribeUserInfoInChannelResponseBodyProperty) *DescribeUserInfoInChannelResponseBody {
	s.Property = v
	return s
}

func (s *DescribeUserInfoInChannelResponseBody) SetRequestId(v string) *DescribeUserInfoInChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserInfoInChannelResponseBody) SetTimestamp(v int32) *DescribeUserInfoInChannelResponseBody {
	s.Timestamp = &v
	return s
}

type DescribeUserInfoInChannelResponseBodyProperty struct {
	Join    *int32  `json:"Join,omitempty" xml:"Join,omitempty"`
	Role    *int32  `json:"Role,omitempty" xml:"Role,omitempty"`
	Session *string `json:"Session,omitempty" xml:"Session,omitempty"`
}

func (s DescribeUserInfoInChannelResponseBodyProperty) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserInfoInChannelResponseBodyProperty) GoString() string {
	return s.String()
}

func (s *DescribeUserInfoInChannelResponseBodyProperty) SetJoin(v int32) *DescribeUserInfoInChannelResponseBodyProperty {
	s.Join = &v
	return s
}

func (s *DescribeUserInfoInChannelResponseBodyProperty) SetRole(v int32) *DescribeUserInfoInChannelResponseBodyProperty {
	s.Role = &v
	return s
}

func (s *DescribeUserInfoInChannelResponseBodyProperty) SetSession(v string) *DescribeUserInfoInChannelResponseBodyProperty {
	s.Session = &v
	return s
}

type DescribeUserInfoInChannelResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeUserInfoInChannelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserInfoInChannelResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserInfoInChannelResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserInfoInChannelResponse) SetHeaders(v map[string]*string) *DescribeUserInfoInChannelResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserInfoInChannelResponse) SetBody(v *DescribeUserInfoInChannelResponseBody) *DescribeUserInfoInChannelResponse {
	s.Body = v
	return s
}

type DisableAutoLiveStreamRuleRequest struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RuleId  *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DisableAutoLiveStreamRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableAutoLiveStreamRuleRequest) GoString() string {
	return s.String()
}

func (s *DisableAutoLiveStreamRuleRequest) SetAppId(v string) *DisableAutoLiveStreamRuleRequest {
	s.AppId = &v
	return s
}

func (s *DisableAutoLiveStreamRuleRequest) SetOwnerId(v int64) *DisableAutoLiveStreamRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *DisableAutoLiveStreamRuleRequest) SetRuleId(v int64) *DisableAutoLiveStreamRuleRequest {
	s.RuleId = &v
	return s
}

type DisableAutoLiveStreamRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableAutoLiveStreamRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableAutoLiveStreamRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DisableAutoLiveStreamRuleResponseBody) SetRequestId(v string) *DisableAutoLiveStreamRuleResponseBody {
	s.RequestId = &v
	return s
}

type DisableAutoLiveStreamRuleResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableAutoLiveStreamRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableAutoLiveStreamRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableAutoLiveStreamRuleResponse) GoString() string {
	return s.String()
}

func (s *DisableAutoLiveStreamRuleResponse) SetHeaders(v map[string]*string) *DisableAutoLiveStreamRuleResponse {
	s.Headers = v
	return s
}

func (s *DisableAutoLiveStreamRuleResponse) SetBody(v *DisableAutoLiveStreamRuleResponseBody) *DisableAutoLiveStreamRuleResponse {
	s.Body = v
	return s
}

type EnableAutoLiveStreamRuleRequest struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RuleId  *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s EnableAutoLiveStreamRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableAutoLiveStreamRuleRequest) GoString() string {
	return s.String()
}

func (s *EnableAutoLiveStreamRuleRequest) SetAppId(v string) *EnableAutoLiveStreamRuleRequest {
	s.AppId = &v
	return s
}

func (s *EnableAutoLiveStreamRuleRequest) SetOwnerId(v int64) *EnableAutoLiveStreamRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *EnableAutoLiveStreamRuleRequest) SetRuleId(v int64) *EnableAutoLiveStreamRuleRequest {
	s.RuleId = &v
	return s
}

type EnableAutoLiveStreamRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableAutoLiveStreamRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableAutoLiveStreamRuleResponseBody) GoString() string {
	return s.String()
}

func (s *EnableAutoLiveStreamRuleResponseBody) SetRequestId(v string) *EnableAutoLiveStreamRuleResponseBody {
	s.RequestId = &v
	return s
}

type EnableAutoLiveStreamRuleResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableAutoLiveStreamRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableAutoLiveStreamRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableAutoLiveStreamRuleResponse) GoString() string {
	return s.String()
}

func (s *EnableAutoLiveStreamRuleResponse) SetHeaders(v map[string]*string) *EnableAutoLiveStreamRuleResponse {
	s.Headers = v
	return s
}

func (s *EnableAutoLiveStreamRuleResponse) SetBody(v *EnableAutoLiveStreamRuleResponseBody) *EnableAutoLiveStreamRuleResponse {
	s.Body = v
	return s
}

type GetMPUTaskStatusRequest struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	TaskId  *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetMPUTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMPUTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *GetMPUTaskStatusRequest) SetAppId(v string) *GetMPUTaskStatusRequest {
	s.AppId = &v
	return s
}

func (s *GetMPUTaskStatusRequest) SetOwnerId(v int64) *GetMPUTaskStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *GetMPUTaskStatusRequest) SetTaskId(v string) *GetMPUTaskStatusRequest {
	s.TaskId = &v
	return s
}

type GetMPUTaskStatusResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetMPUTaskStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMPUTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetMPUTaskStatusResponseBody) SetRequestId(v string) *GetMPUTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMPUTaskStatusResponseBody) SetStatus(v int32) *GetMPUTaskStatusResponseBody {
	s.Status = &v
	return s
}

type GetMPUTaskStatusResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetMPUTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetMPUTaskStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetMPUTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *GetMPUTaskStatusResponse) SetHeaders(v map[string]*string) *GetMPUTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *GetMPUTaskStatusResponse) SetBody(v *GetMPUTaskStatusResponseBody) *GetMPUTaskStatusResponse {
	s.Body = v
	return s
}

type ModifyMPULayoutRequest struct {
	AppId         *string                        `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AudioMixCount *int32                         `json:"AudioMixCount,omitempty" xml:"AudioMixCount,omitempty"`
	LayoutId      *int64                         `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	Name          *string                        `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId       *int64                         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Panes         []*ModifyMPULayoutRequestPanes `json:"Panes,omitempty" xml:"Panes,omitempty" type:"Repeated"`
}

func (s ModifyMPULayoutRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyMPULayoutRequest) GoString() string {
	return s.String()
}

func (s *ModifyMPULayoutRequest) SetAppId(v string) *ModifyMPULayoutRequest {
	s.AppId = &v
	return s
}

func (s *ModifyMPULayoutRequest) SetAudioMixCount(v int32) *ModifyMPULayoutRequest {
	s.AudioMixCount = &v
	return s
}

func (s *ModifyMPULayoutRequest) SetLayoutId(v int64) *ModifyMPULayoutRequest {
	s.LayoutId = &v
	return s
}

func (s *ModifyMPULayoutRequest) SetName(v string) *ModifyMPULayoutRequest {
	s.Name = &v
	return s
}

func (s *ModifyMPULayoutRequest) SetOwnerId(v int64) *ModifyMPULayoutRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyMPULayoutRequest) SetPanes(v []*ModifyMPULayoutRequestPanes) *ModifyMPULayoutRequest {
	s.Panes = v
	return s
}

type ModifyMPULayoutRequestPanes struct {
	Height    *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	MajorPane *int32   `json:"MajorPane,omitempty" xml:"MajorPane,omitempty"`
	PaneId    *int32   `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
	Width     *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	X         *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y         *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	ZOrder    *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s ModifyMPULayoutRequestPanes) String() string {
	return tea.Prettify(s)
}

func (s ModifyMPULayoutRequestPanes) GoString() string {
	return s.String()
}

func (s *ModifyMPULayoutRequestPanes) SetHeight(v float32) *ModifyMPULayoutRequestPanes {
	s.Height = &v
	return s
}

func (s *ModifyMPULayoutRequestPanes) SetMajorPane(v int32) *ModifyMPULayoutRequestPanes {
	s.MajorPane = &v
	return s
}

func (s *ModifyMPULayoutRequestPanes) SetPaneId(v int32) *ModifyMPULayoutRequestPanes {
	s.PaneId = &v
	return s
}

func (s *ModifyMPULayoutRequestPanes) SetWidth(v float32) *ModifyMPULayoutRequestPanes {
	s.Width = &v
	return s
}

func (s *ModifyMPULayoutRequestPanes) SetX(v float32) *ModifyMPULayoutRequestPanes {
	s.X = &v
	return s
}

func (s *ModifyMPULayoutRequestPanes) SetY(v float32) *ModifyMPULayoutRequestPanes {
	s.Y = &v
	return s
}

func (s *ModifyMPULayoutRequestPanes) SetZOrder(v int32) *ModifyMPULayoutRequestPanes {
	s.ZOrder = &v
	return s
}

type ModifyMPULayoutResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyMPULayoutResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyMPULayoutResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMPULayoutResponseBody) SetRequestId(v string) *ModifyMPULayoutResponseBody {
	s.RequestId = &v
	return s
}

type ModifyMPULayoutResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyMPULayoutResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyMPULayoutResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyMPULayoutResponse) GoString() string {
	return s.String()
}

func (s *ModifyMPULayoutResponse) SetHeaders(v map[string]*string) *ModifyMPULayoutResponse {
	s.Headers = v
	return s
}

func (s *ModifyMPULayoutResponse) SetBody(v *ModifyMPULayoutResponseBody) *ModifyMPULayoutResponse {
	s.Body = v
	return s
}

type RemoveTerminalsRequest struct {
	AppId       *string   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId   *string   `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	OwnerId     *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	TerminalIds []*string `json:"TerminalIds,omitempty" xml:"TerminalIds,omitempty" type:"Repeated"`
}

func (s RemoveTerminalsRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveTerminalsRequest) GoString() string {
	return s.String()
}

func (s *RemoveTerminalsRequest) SetAppId(v string) *RemoveTerminalsRequest {
	s.AppId = &v
	return s
}

func (s *RemoveTerminalsRequest) SetChannelId(v string) *RemoveTerminalsRequest {
	s.ChannelId = &v
	return s
}

func (s *RemoveTerminalsRequest) SetOwnerId(v int64) *RemoveTerminalsRequest {
	s.OwnerId = &v
	return s
}

func (s *RemoveTerminalsRequest) SetTerminalIds(v []*string) *RemoveTerminalsRequest {
	s.TerminalIds = v
	return s
}

type RemoveTerminalsResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Terminals *RemoveTerminalsResponseBodyTerminals `json:"Terminals,omitempty" xml:"Terminals,omitempty" type:"Struct"`
}

func (s RemoveTerminalsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveTerminalsResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveTerminalsResponseBody) SetRequestId(v string) *RemoveTerminalsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveTerminalsResponseBody) SetTerminals(v *RemoveTerminalsResponseBodyTerminals) *RemoveTerminalsResponseBody {
	s.Terminals = v
	return s
}

type RemoveTerminalsResponseBodyTerminals struct {
	Terminal []*RemoveTerminalsResponseBodyTerminalsTerminal `json:"Terminal,omitempty" xml:"Terminal,omitempty" type:"Repeated"`
}

func (s RemoveTerminalsResponseBodyTerminals) String() string {
	return tea.Prettify(s)
}

func (s RemoveTerminalsResponseBodyTerminals) GoString() string {
	return s.String()
}

func (s *RemoveTerminalsResponseBodyTerminals) SetTerminal(v []*RemoveTerminalsResponseBodyTerminalsTerminal) *RemoveTerminalsResponseBodyTerminals {
	s.Terminal = v
	return s
}

type RemoveTerminalsResponseBodyTerminalsTerminal struct {
	Code    *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RemoveTerminalsResponseBodyTerminalsTerminal) String() string {
	return tea.Prettify(s)
}

func (s RemoveTerminalsResponseBodyTerminalsTerminal) GoString() string {
	return s.String()
}

func (s *RemoveTerminalsResponseBodyTerminalsTerminal) SetCode(v int32) *RemoveTerminalsResponseBodyTerminalsTerminal {
	s.Code = &v
	return s
}

func (s *RemoveTerminalsResponseBodyTerminalsTerminal) SetId(v string) *RemoveTerminalsResponseBodyTerminalsTerminal {
	s.Id = &v
	return s
}

func (s *RemoveTerminalsResponseBodyTerminalsTerminal) SetMessage(v string) *RemoveTerminalsResponseBodyTerminalsTerminal {
	s.Message = &v
	return s
}

type RemoveTerminalsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveTerminalsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveTerminalsResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveTerminalsResponse) GoString() string {
	return s.String()
}

func (s *RemoveTerminalsResponse) SetHeaders(v map[string]*string) *RemoveTerminalsResponse {
	s.Headers = v
	return s
}

func (s *RemoveTerminalsResponse) SetBody(v *RemoveTerminalsResponseBody) *RemoveTerminalsResponse {
	s.Body = v
	return s
}

type StartMPUTaskRequest struct {
	AppId                     *string                            `json:"AppId,omitempty" xml:"AppId,omitempty"`
	BackgroundColor           *int32                             `json:"BackgroundColor,omitempty" xml:"BackgroundColor,omitempty"`
	Backgrounds               []*StartMPUTaskRequestBackgrounds  `json:"Backgrounds,omitempty" xml:"Backgrounds,omitempty" type:"Repeated"`
	ChannelId                 *string                            `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ClockWidgets              []*StartMPUTaskRequestClockWidgets `json:"ClockWidgets,omitempty" xml:"ClockWidgets,omitempty" type:"Repeated"`
	CropMode                  *int32                             `json:"CropMode,omitempty" xml:"CropMode,omitempty"`
	EnhancedParam             *StartMPUTaskRequestEnhancedParam  `json:"EnhancedParam,omitempty" xml:"EnhancedParam,omitempty" type:"Struct"`
	LayoutIds                 []*int64                           `json:"LayoutIds,omitempty" xml:"LayoutIds,omitempty" type:"Repeated"`
	MediaEncode               *int32                             `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	MixMode                   *int32                             `json:"MixMode,omitempty" xml:"MixMode,omitempty"`
	OwnerId                   *int64                             `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PayloadType               *int32                             `json:"PayloadType,omitempty" xml:"PayloadType,omitempty"`
	ReportVad                 *int32                             `json:"ReportVad,omitempty" xml:"ReportVad,omitempty"`
	RtpExtInfo                *int32                             `json:"RtpExtInfo,omitempty" xml:"RtpExtInfo,omitempty"`
	SourceType                *string                            `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	StreamType                *int32                             `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	StreamURL                 *string                            `json:"StreamURL,omitempty" xml:"StreamURL,omitempty"`
	SubSpecAudioUsers         []*string                          `json:"SubSpecAudioUsers,omitempty" xml:"SubSpecAudioUsers,omitempty" type:"Repeated"`
	SubSpecCameraUsers        []*string                          `json:"SubSpecCameraUsers,omitempty" xml:"SubSpecCameraUsers,omitempty" type:"Repeated"`
	SubSpecShareScreenUsers   []*string                          `json:"SubSpecShareScreenUsers,omitempty" xml:"SubSpecShareScreenUsers,omitempty" type:"Repeated"`
	SubSpecUsers              []*string                          `json:"SubSpecUsers,omitempty" xml:"SubSpecUsers,omitempty" type:"Repeated"`
	TaskId                    *string                            `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskType                  *int32                             `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TimeStampRef              *int64                             `json:"TimeStampRef,omitempty" xml:"TimeStampRef,omitempty"`
	UnsubSpecAudioUsers       []*string                          `json:"UnsubSpecAudioUsers,omitempty" xml:"UnsubSpecAudioUsers,omitempty" type:"Repeated"`
	UnsubSpecCameraUsers      []*string                          `json:"UnsubSpecCameraUsers,omitempty" xml:"UnsubSpecCameraUsers,omitempty" type:"Repeated"`
	UnsubSpecShareScreenUsers []*string                          `json:"UnsubSpecShareScreenUsers,omitempty" xml:"UnsubSpecShareScreenUsers,omitempty" type:"Repeated"`
	UserPanes                 []*StartMPUTaskRequestUserPanes    `json:"UserPanes,omitempty" xml:"UserPanes,omitempty" type:"Repeated"`
	VadInterval               *int64                             `json:"VadInterval,omitempty" xml:"VadInterval,omitempty"`
	Watermarks                []*StartMPUTaskRequestWatermarks   `json:"Watermarks,omitempty" xml:"Watermarks,omitempty" type:"Repeated"`
}

func (s StartMPUTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s StartMPUTaskRequest) GoString() string {
	return s.String()
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

type StartMPUTaskRequestBackgrounds struct {
	Display *int32   `json:"Display,omitempty" xml:"Display,omitempty"`
	Height  *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Url     *string  `json:"Url,omitempty" xml:"Url,omitempty"`
	Width   *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	X       *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y       *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	ZOrder  *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s StartMPUTaskRequestBackgrounds) String() string {
	return tea.Prettify(s)
}

func (s StartMPUTaskRequestBackgrounds) GoString() string {
	return s.String()
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

type StartMPUTaskRequestClockWidgets struct {
	FontColor *int32   `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	FontSize  *int32   `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	FontType  *int32   `json:"FontType,omitempty" xml:"FontType,omitempty"`
	X         *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y         *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	ZOrder    *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s StartMPUTaskRequestClockWidgets) String() string {
	return tea.Prettify(s)
}

func (s StartMPUTaskRequestClockWidgets) GoString() string {
	return s.String()
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

type StartMPUTaskRequestEnhancedParam struct {
	EnablePortraitSegmentation *bool `json:"EnablePortraitSegmentation,omitempty" xml:"EnablePortraitSegmentation,omitempty"`
}

func (s StartMPUTaskRequestEnhancedParam) String() string {
	return tea.Prettify(s)
}

func (s StartMPUTaskRequestEnhancedParam) GoString() string {
	return s.String()
}

func (s *StartMPUTaskRequestEnhancedParam) SetEnablePortraitSegmentation(v bool) *StartMPUTaskRequestEnhancedParam {
	s.EnablePortraitSegmentation = &v
	return s
}

type StartMPUTaskRequestUserPanes struct {
	Images      []*StartMPUTaskRequestUserPanesImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	PaneId      *int32                                `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
	SegmentType *int32                                `json:"SegmentType,omitempty" xml:"SegmentType,omitempty"`
	SourceType  *string                               `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Texts       []*StartMPUTaskRequestUserPanesTexts  `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
	UserId      *string                               `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s StartMPUTaskRequestUserPanes) String() string {
	return tea.Prettify(s)
}

func (s StartMPUTaskRequestUserPanes) GoString() string {
	return s.String()
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

type StartMPUTaskRequestUserPanesImages struct {
	Display *int32   `json:"Display,omitempty" xml:"Display,omitempty"`
	Height  *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Url     *string  `json:"Url,omitempty" xml:"Url,omitempty"`
	Width   *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	X       *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y       *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	ZOrder  *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s StartMPUTaskRequestUserPanesImages) String() string {
	return tea.Prettify(s)
}

func (s StartMPUTaskRequestUserPanesImages) GoString() string {
	return s.String()
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

type StartMPUTaskRequestUserPanesTexts struct {
	FontColor *int32   `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	FontSize  *int32   `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	FontType  *int32   `json:"FontType,omitempty" xml:"FontType,omitempty"`
	Text      *string  `json:"Text,omitempty" xml:"Text,omitempty"`
	X         *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y         *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	ZOrder    *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s StartMPUTaskRequestUserPanesTexts) String() string {
	return tea.Prettify(s)
}

func (s StartMPUTaskRequestUserPanesTexts) GoString() string {
	return s.String()
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

type StartMPUTaskRequestWatermarks struct {
	Alpha   *float32 `json:"Alpha,omitempty" xml:"Alpha,omitempty"`
	Display *int32   `json:"Display,omitempty" xml:"Display,omitempty"`
	Height  *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Url     *string  `json:"Url,omitempty" xml:"Url,omitempty"`
	Width   *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	X       *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y       *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	ZOrder  *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s StartMPUTaskRequestWatermarks) String() string {
	return tea.Prettify(s)
}

func (s StartMPUTaskRequestWatermarks) GoString() string {
	return s.String()
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

type StartMPUTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartMPUTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartMPUTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StartMPUTaskResponseBody) SetRequestId(v string) *StartMPUTaskResponseBody {
	s.RequestId = &v
	return s
}

type StartMPUTaskResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartMPUTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartMPUTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s StartMPUTaskResponse) GoString() string {
	return s.String()
}

func (s *StartMPUTaskResponse) SetHeaders(v map[string]*string) *StartMPUTaskResponse {
	s.Headers = v
	return s
}

func (s *StartMPUTaskResponse) SetBody(v *StartMPUTaskResponseBody) *StartMPUTaskResponse {
	s.Body = v
	return s
}

type StartRecordTaskRequest struct {
	AppId                     *string                            `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId                 *string                            `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	CropMode                  *int64                             `json:"CropMode,omitempty" xml:"CropMode,omitempty"`
	LayoutIds                 []*int64                           `json:"LayoutIds,omitempty" xml:"LayoutIds,omitempty" type:"Repeated"`
	MediaEncode               *int32                             `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	MixMode                   *int32                             `json:"MixMode,omitempty" xml:"MixMode,omitempty"`
	OwnerId                   *int64                             `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SourceType                *string                            `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	StreamType                *int32                             `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	SubSpecAudioUsers         []*string                          `json:"SubSpecAudioUsers,omitempty" xml:"SubSpecAudioUsers,omitempty" type:"Repeated"`
	SubSpecCameraUsers        []*string                          `json:"SubSpecCameraUsers,omitempty" xml:"SubSpecCameraUsers,omitempty" type:"Repeated"`
	SubSpecShareScreenUsers   []*string                          `json:"SubSpecShareScreenUsers,omitempty" xml:"SubSpecShareScreenUsers,omitempty" type:"Repeated"`
	SubSpecUsers              []*string                          `json:"SubSpecUsers,omitempty" xml:"SubSpecUsers,omitempty" type:"Repeated"`
	TaskId                    *string                            `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskProfile               *string                            `json:"TaskProfile,omitempty" xml:"TaskProfile,omitempty"`
	TemplateId                *string                            `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	UnsubSpecAudioUsers       []*string                          `json:"UnsubSpecAudioUsers,omitempty" xml:"UnsubSpecAudioUsers,omitempty" type:"Repeated"`
	UnsubSpecCameraUsers      []*string                          `json:"UnsubSpecCameraUsers,omitempty" xml:"UnsubSpecCameraUsers,omitempty" type:"Repeated"`
	UnsubSpecShareScreenUsers []*string                          `json:"UnsubSpecShareScreenUsers,omitempty" xml:"UnsubSpecShareScreenUsers,omitempty" type:"Repeated"`
	UserPanes                 []*StartRecordTaskRequestUserPanes `json:"UserPanes,omitempty" xml:"UserPanes,omitempty" type:"Repeated"`
}

func (s StartRecordTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s StartRecordTaskRequest) GoString() string {
	return s.String()
}

func (s *StartRecordTaskRequest) SetAppId(v string) *StartRecordTaskRequest {
	s.AppId = &v
	return s
}

func (s *StartRecordTaskRequest) SetChannelId(v string) *StartRecordTaskRequest {
	s.ChannelId = &v
	return s
}

func (s *StartRecordTaskRequest) SetCropMode(v int64) *StartRecordTaskRequest {
	s.CropMode = &v
	return s
}

func (s *StartRecordTaskRequest) SetLayoutIds(v []*int64) *StartRecordTaskRequest {
	s.LayoutIds = v
	return s
}

func (s *StartRecordTaskRequest) SetMediaEncode(v int32) *StartRecordTaskRequest {
	s.MediaEncode = &v
	return s
}

func (s *StartRecordTaskRequest) SetMixMode(v int32) *StartRecordTaskRequest {
	s.MixMode = &v
	return s
}

func (s *StartRecordTaskRequest) SetOwnerId(v int64) *StartRecordTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *StartRecordTaskRequest) SetSourceType(v string) *StartRecordTaskRequest {
	s.SourceType = &v
	return s
}

func (s *StartRecordTaskRequest) SetStreamType(v int32) *StartRecordTaskRequest {
	s.StreamType = &v
	return s
}

func (s *StartRecordTaskRequest) SetSubSpecAudioUsers(v []*string) *StartRecordTaskRequest {
	s.SubSpecAudioUsers = v
	return s
}

func (s *StartRecordTaskRequest) SetSubSpecCameraUsers(v []*string) *StartRecordTaskRequest {
	s.SubSpecCameraUsers = v
	return s
}

func (s *StartRecordTaskRequest) SetSubSpecShareScreenUsers(v []*string) *StartRecordTaskRequest {
	s.SubSpecShareScreenUsers = v
	return s
}

func (s *StartRecordTaskRequest) SetSubSpecUsers(v []*string) *StartRecordTaskRequest {
	s.SubSpecUsers = v
	return s
}

func (s *StartRecordTaskRequest) SetTaskId(v string) *StartRecordTaskRequest {
	s.TaskId = &v
	return s
}

func (s *StartRecordTaskRequest) SetTaskProfile(v string) *StartRecordTaskRequest {
	s.TaskProfile = &v
	return s
}

func (s *StartRecordTaskRequest) SetTemplateId(v string) *StartRecordTaskRequest {
	s.TemplateId = &v
	return s
}

func (s *StartRecordTaskRequest) SetUnsubSpecAudioUsers(v []*string) *StartRecordTaskRequest {
	s.UnsubSpecAudioUsers = v
	return s
}

func (s *StartRecordTaskRequest) SetUnsubSpecCameraUsers(v []*string) *StartRecordTaskRequest {
	s.UnsubSpecCameraUsers = v
	return s
}

func (s *StartRecordTaskRequest) SetUnsubSpecShareScreenUsers(v []*string) *StartRecordTaskRequest {
	s.UnsubSpecShareScreenUsers = v
	return s
}

func (s *StartRecordTaskRequest) SetUserPanes(v []*StartRecordTaskRequestUserPanes) *StartRecordTaskRequest {
	s.UserPanes = v
	return s
}

type StartRecordTaskRequestUserPanes struct {
	Images     []*StartRecordTaskRequestUserPanesImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	PaneId     *int32                                   `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
	SourceType *string                                  `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Texts      []*StartRecordTaskRequestUserPanesTexts  `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
	UserId     *string                                  `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s StartRecordTaskRequestUserPanes) String() string {
	return tea.Prettify(s)
}

func (s StartRecordTaskRequestUserPanes) GoString() string {
	return s.String()
}

func (s *StartRecordTaskRequestUserPanes) SetImages(v []*StartRecordTaskRequestUserPanesImages) *StartRecordTaskRequestUserPanes {
	s.Images = v
	return s
}

func (s *StartRecordTaskRequestUserPanes) SetPaneId(v int32) *StartRecordTaskRequestUserPanes {
	s.PaneId = &v
	return s
}

func (s *StartRecordTaskRequestUserPanes) SetSourceType(v string) *StartRecordTaskRequestUserPanes {
	s.SourceType = &v
	return s
}

func (s *StartRecordTaskRequestUserPanes) SetTexts(v []*StartRecordTaskRequestUserPanesTexts) *StartRecordTaskRequestUserPanes {
	s.Texts = v
	return s
}

func (s *StartRecordTaskRequestUserPanes) SetUserId(v string) *StartRecordTaskRequestUserPanes {
	s.UserId = &v
	return s
}

type StartRecordTaskRequestUserPanesImages struct {
	Display *int32   `json:"Display,omitempty" xml:"Display,omitempty"`
	Height  *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Url     *string  `json:"Url,omitempty" xml:"Url,omitempty"`
	Width   *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	X       *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y       *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	ZOrder  *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s StartRecordTaskRequestUserPanesImages) String() string {
	return tea.Prettify(s)
}

func (s StartRecordTaskRequestUserPanesImages) GoString() string {
	return s.String()
}

func (s *StartRecordTaskRequestUserPanesImages) SetDisplay(v int32) *StartRecordTaskRequestUserPanesImages {
	s.Display = &v
	return s
}

func (s *StartRecordTaskRequestUserPanesImages) SetHeight(v float32) *StartRecordTaskRequestUserPanesImages {
	s.Height = &v
	return s
}

func (s *StartRecordTaskRequestUserPanesImages) SetUrl(v string) *StartRecordTaskRequestUserPanesImages {
	s.Url = &v
	return s
}

func (s *StartRecordTaskRequestUserPanesImages) SetWidth(v float32) *StartRecordTaskRequestUserPanesImages {
	s.Width = &v
	return s
}

func (s *StartRecordTaskRequestUserPanesImages) SetX(v float32) *StartRecordTaskRequestUserPanesImages {
	s.X = &v
	return s
}

func (s *StartRecordTaskRequestUserPanesImages) SetY(v float32) *StartRecordTaskRequestUserPanesImages {
	s.Y = &v
	return s
}

func (s *StartRecordTaskRequestUserPanesImages) SetZOrder(v int32) *StartRecordTaskRequestUserPanesImages {
	s.ZOrder = &v
	return s
}

type StartRecordTaskRequestUserPanesTexts struct {
	FontColor *int32   `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	FontSize  *int32   `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	FontType  *int32   `json:"FontType,omitempty" xml:"FontType,omitempty"`
	Text      *string  `json:"Text,omitempty" xml:"Text,omitempty"`
	X         *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y         *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	ZOrder    *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s StartRecordTaskRequestUserPanesTexts) String() string {
	return tea.Prettify(s)
}

func (s StartRecordTaskRequestUserPanesTexts) GoString() string {
	return s.String()
}

func (s *StartRecordTaskRequestUserPanesTexts) SetFontColor(v int32) *StartRecordTaskRequestUserPanesTexts {
	s.FontColor = &v
	return s
}

func (s *StartRecordTaskRequestUserPanesTexts) SetFontSize(v int32) *StartRecordTaskRequestUserPanesTexts {
	s.FontSize = &v
	return s
}

func (s *StartRecordTaskRequestUserPanesTexts) SetFontType(v int32) *StartRecordTaskRequestUserPanesTexts {
	s.FontType = &v
	return s
}

func (s *StartRecordTaskRequestUserPanesTexts) SetText(v string) *StartRecordTaskRequestUserPanesTexts {
	s.Text = &v
	return s
}

func (s *StartRecordTaskRequestUserPanesTexts) SetX(v float32) *StartRecordTaskRequestUserPanesTexts {
	s.X = &v
	return s
}

func (s *StartRecordTaskRequestUserPanesTexts) SetY(v float32) *StartRecordTaskRequestUserPanesTexts {
	s.Y = &v
	return s
}

func (s *StartRecordTaskRequestUserPanesTexts) SetZOrder(v int32) *StartRecordTaskRequestUserPanesTexts {
	s.ZOrder = &v
	return s
}

type StartRecordTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartRecordTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartRecordTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StartRecordTaskResponseBody) SetRequestId(v string) *StartRecordTaskResponseBody {
	s.RequestId = &v
	return s
}

type StartRecordTaskResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartRecordTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartRecordTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s StartRecordTaskResponse) GoString() string {
	return s.String()
}

func (s *StartRecordTaskResponse) SetHeaders(v map[string]*string) *StartRecordTaskResponse {
	s.Headers = v
	return s
}

func (s *StartRecordTaskResponse) SetBody(v *StartRecordTaskResponseBody) *StartRecordTaskResponse {
	s.Body = v
	return s
}

type StopChannelUserPublishRequest struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	UserId    *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s StopChannelUserPublishRequest) String() string {
	return tea.Prettify(s)
}

func (s StopChannelUserPublishRequest) GoString() string {
	return s.String()
}

func (s *StopChannelUserPublishRequest) SetAppId(v string) *StopChannelUserPublishRequest {
	s.AppId = &v
	return s
}

func (s *StopChannelUserPublishRequest) SetChannelId(v string) *StopChannelUserPublishRequest {
	s.ChannelId = &v
	return s
}

func (s *StopChannelUserPublishRequest) SetOwnerId(v int64) *StopChannelUserPublishRequest {
	s.OwnerId = &v
	return s
}

func (s *StopChannelUserPublishRequest) SetUserId(v string) *StopChannelUserPublishRequest {
	s.UserId = &v
	return s
}

type StopChannelUserPublishResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopChannelUserPublishResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopChannelUserPublishResponseBody) GoString() string {
	return s.String()
}

func (s *StopChannelUserPublishResponseBody) SetRequestId(v string) *StopChannelUserPublishResponseBody {
	s.RequestId = &v
	return s
}

type StopChannelUserPublishResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopChannelUserPublishResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopChannelUserPublishResponse) String() string {
	return tea.Prettify(s)
}

func (s StopChannelUserPublishResponse) GoString() string {
	return s.String()
}

func (s *StopChannelUserPublishResponse) SetHeaders(v map[string]*string) *StopChannelUserPublishResponse {
	s.Headers = v
	return s
}

func (s *StopChannelUserPublishResponse) SetBody(v *StopChannelUserPublishResponseBody) *StopChannelUserPublishResponse {
	s.Body = v
	return s
}

type StopMPUTaskRequest struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	TaskId  *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopMPUTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s StopMPUTaskRequest) GoString() string {
	return s.String()
}

func (s *StopMPUTaskRequest) SetAppId(v string) *StopMPUTaskRequest {
	s.AppId = &v
	return s
}

func (s *StopMPUTaskRequest) SetOwnerId(v int64) *StopMPUTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *StopMPUTaskRequest) SetTaskId(v string) *StopMPUTaskRequest {
	s.TaskId = &v
	return s
}

type StopMPUTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopMPUTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopMPUTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopMPUTaskResponseBody) SetRequestId(v string) *StopMPUTaskResponseBody {
	s.RequestId = &v
	return s
}

type StopMPUTaskResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopMPUTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopMPUTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s StopMPUTaskResponse) GoString() string {
	return s.String()
}

func (s *StopMPUTaskResponse) SetHeaders(v map[string]*string) *StopMPUTaskResponse {
	s.Headers = v
	return s
}

func (s *StopMPUTaskResponse) SetBody(v *StopMPUTaskResponseBody) *StopMPUTaskResponse {
	s.Body = v
	return s
}

type StopRecordTaskRequest struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	TaskId  *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopRecordTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s StopRecordTaskRequest) GoString() string {
	return s.String()
}

func (s *StopRecordTaskRequest) SetAppId(v string) *StopRecordTaskRequest {
	s.AppId = &v
	return s
}

func (s *StopRecordTaskRequest) SetOwnerId(v int64) *StopRecordTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *StopRecordTaskRequest) SetTaskId(v string) *StopRecordTaskRequest {
	s.TaskId = &v
	return s
}

type StopRecordTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopRecordTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopRecordTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopRecordTaskResponseBody) SetRequestId(v string) *StopRecordTaskResponseBody {
	s.RequestId = &v
	return s
}

type StopRecordTaskResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopRecordTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopRecordTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s StopRecordTaskResponse) GoString() string {
	return s.String()
}

func (s *StopRecordTaskResponse) SetHeaders(v map[string]*string) *StopRecordTaskResponse {
	s.Headers = v
	return s
}

func (s *StopRecordTaskResponse) SetBody(v *StopRecordTaskResponseBody) *StopRecordTaskResponse {
	s.Body = v
	return s
}

type UpdateAutoLiveStreamRuleRequest struct {
	AppId             *string   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CallBack          *string   `json:"CallBack,omitempty" xml:"CallBack,omitempty"`
	ChannelIdPrefixes []*string `json:"ChannelIdPrefixes,omitempty" xml:"ChannelIdPrefixes,omitempty" type:"Repeated"`
	ChannelIds        []*string `json:"ChannelIds,omitempty" xml:"ChannelIds,omitempty" type:"Repeated"`
	MediaEncode       *int32    `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	OwnerId           *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PlayDomain        *string   `json:"PlayDomain,omitempty" xml:"PlayDomain,omitempty"`
	RuleId            *int32    `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	RuleName          *string   `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s UpdateAutoLiveStreamRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateAutoLiveStreamRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateAutoLiveStreamRuleRequest) SetAppId(v string) *UpdateAutoLiveStreamRuleRequest {
	s.AppId = &v
	return s
}

func (s *UpdateAutoLiveStreamRuleRequest) SetCallBack(v string) *UpdateAutoLiveStreamRuleRequest {
	s.CallBack = &v
	return s
}

func (s *UpdateAutoLiveStreamRuleRequest) SetChannelIdPrefixes(v []*string) *UpdateAutoLiveStreamRuleRequest {
	s.ChannelIdPrefixes = v
	return s
}

func (s *UpdateAutoLiveStreamRuleRequest) SetChannelIds(v []*string) *UpdateAutoLiveStreamRuleRequest {
	s.ChannelIds = v
	return s
}

func (s *UpdateAutoLiveStreamRuleRequest) SetMediaEncode(v int32) *UpdateAutoLiveStreamRuleRequest {
	s.MediaEncode = &v
	return s
}

func (s *UpdateAutoLiveStreamRuleRequest) SetOwnerId(v int64) *UpdateAutoLiveStreamRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateAutoLiveStreamRuleRequest) SetPlayDomain(v string) *UpdateAutoLiveStreamRuleRequest {
	s.PlayDomain = &v
	return s
}

func (s *UpdateAutoLiveStreamRuleRequest) SetRuleId(v int32) *UpdateAutoLiveStreamRuleRequest {
	s.RuleId = &v
	return s
}

func (s *UpdateAutoLiveStreamRuleRequest) SetRuleName(v string) *UpdateAutoLiveStreamRuleRequest {
	s.RuleName = &v
	return s
}

type UpdateAutoLiveStreamRuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RuleId    *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s UpdateAutoLiveStreamRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateAutoLiveStreamRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAutoLiveStreamRuleResponseBody) SetRequestId(v string) *UpdateAutoLiveStreamRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAutoLiveStreamRuleResponseBody) SetRuleId(v int64) *UpdateAutoLiveStreamRuleResponseBody {
	s.RuleId = &v
	return s
}

type UpdateAutoLiveStreamRuleResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateAutoLiveStreamRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateAutoLiveStreamRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateAutoLiveStreamRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateAutoLiveStreamRuleResponse) SetHeaders(v map[string]*string) *UpdateAutoLiveStreamRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateAutoLiveStreamRuleResponse) SetBody(v *UpdateAutoLiveStreamRuleResponseBody) *UpdateAutoLiveStreamRuleResponse {
	s.Body = v
	return s
}

type UpdateMPUTaskRequest struct {
	AppId                     *string                             `json:"AppId,omitempty" xml:"AppId,omitempty"`
	BackgroundColor           *int32                              `json:"BackgroundColor,omitempty" xml:"BackgroundColor,omitempty"`
	Backgrounds               []*UpdateMPUTaskRequestBackgrounds  `json:"Backgrounds,omitempty" xml:"Backgrounds,omitempty" type:"Repeated"`
	ClockWidgets              []*UpdateMPUTaskRequestClockWidgets `json:"ClockWidgets,omitempty" xml:"ClockWidgets,omitempty" type:"Repeated"`
	CropMode                  *int32                              `json:"CropMode,omitempty" xml:"CropMode,omitempty"`
	LayoutIds                 []*int64                            `json:"LayoutIds,omitempty" xml:"LayoutIds,omitempty" type:"Repeated"`
	MediaEncode               *int32                              `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	MixMode                   *int32                              `json:"MixMode,omitempty" xml:"MixMode,omitempty"`
	OwnerId                   *int64                              `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SourceType                *string                             `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	StreamType                *int32                              `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	SubSpecAudioUsers         []*string                           `json:"SubSpecAudioUsers,omitempty" xml:"SubSpecAudioUsers,omitempty" type:"Repeated"`
	SubSpecCameraUsers        []*string                           `json:"SubSpecCameraUsers,omitempty" xml:"SubSpecCameraUsers,omitempty" type:"Repeated"`
	SubSpecShareScreenUsers   []*string                           `json:"SubSpecShareScreenUsers,omitempty" xml:"SubSpecShareScreenUsers,omitempty" type:"Repeated"`
	SubSpecUsers              []*string                           `json:"SubSpecUsers,omitempty" xml:"SubSpecUsers,omitempty" type:"Repeated"`
	TaskId                    *string                             `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	UnsubSpecAudioUsers       []*string                           `json:"UnsubSpecAudioUsers,omitempty" xml:"UnsubSpecAudioUsers,omitempty" type:"Repeated"`
	UnsubSpecCameraUsers      []*string                           `json:"UnsubSpecCameraUsers,omitempty" xml:"UnsubSpecCameraUsers,omitempty" type:"Repeated"`
	UnsubSpecShareScreenUsers []*string                           `json:"UnsubSpecShareScreenUsers,omitempty" xml:"UnsubSpecShareScreenUsers,omitempty" type:"Repeated"`
	UserPanes                 []*UpdateMPUTaskRequestUserPanes    `json:"UserPanes,omitempty" xml:"UserPanes,omitempty" type:"Repeated"`
	Watermarks                []*UpdateMPUTaskRequestWatermarks   `json:"Watermarks,omitempty" xml:"Watermarks,omitempty" type:"Repeated"`
}

func (s UpdateMPUTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMPUTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateMPUTaskRequest) SetAppId(v string) *UpdateMPUTaskRequest {
	s.AppId = &v
	return s
}

func (s *UpdateMPUTaskRequest) SetBackgroundColor(v int32) *UpdateMPUTaskRequest {
	s.BackgroundColor = &v
	return s
}

func (s *UpdateMPUTaskRequest) SetBackgrounds(v []*UpdateMPUTaskRequestBackgrounds) *UpdateMPUTaskRequest {
	s.Backgrounds = v
	return s
}

func (s *UpdateMPUTaskRequest) SetClockWidgets(v []*UpdateMPUTaskRequestClockWidgets) *UpdateMPUTaskRequest {
	s.ClockWidgets = v
	return s
}

func (s *UpdateMPUTaskRequest) SetCropMode(v int32) *UpdateMPUTaskRequest {
	s.CropMode = &v
	return s
}

func (s *UpdateMPUTaskRequest) SetLayoutIds(v []*int64) *UpdateMPUTaskRequest {
	s.LayoutIds = v
	return s
}

func (s *UpdateMPUTaskRequest) SetMediaEncode(v int32) *UpdateMPUTaskRequest {
	s.MediaEncode = &v
	return s
}

func (s *UpdateMPUTaskRequest) SetMixMode(v int32) *UpdateMPUTaskRequest {
	s.MixMode = &v
	return s
}

func (s *UpdateMPUTaskRequest) SetOwnerId(v int64) *UpdateMPUTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateMPUTaskRequest) SetSourceType(v string) *UpdateMPUTaskRequest {
	s.SourceType = &v
	return s
}

func (s *UpdateMPUTaskRequest) SetStreamType(v int32) *UpdateMPUTaskRequest {
	s.StreamType = &v
	return s
}

func (s *UpdateMPUTaskRequest) SetSubSpecAudioUsers(v []*string) *UpdateMPUTaskRequest {
	s.SubSpecAudioUsers = v
	return s
}

func (s *UpdateMPUTaskRequest) SetSubSpecCameraUsers(v []*string) *UpdateMPUTaskRequest {
	s.SubSpecCameraUsers = v
	return s
}

func (s *UpdateMPUTaskRequest) SetSubSpecShareScreenUsers(v []*string) *UpdateMPUTaskRequest {
	s.SubSpecShareScreenUsers = v
	return s
}

func (s *UpdateMPUTaskRequest) SetSubSpecUsers(v []*string) *UpdateMPUTaskRequest {
	s.SubSpecUsers = v
	return s
}

func (s *UpdateMPUTaskRequest) SetTaskId(v string) *UpdateMPUTaskRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateMPUTaskRequest) SetUnsubSpecAudioUsers(v []*string) *UpdateMPUTaskRequest {
	s.UnsubSpecAudioUsers = v
	return s
}

func (s *UpdateMPUTaskRequest) SetUnsubSpecCameraUsers(v []*string) *UpdateMPUTaskRequest {
	s.UnsubSpecCameraUsers = v
	return s
}

func (s *UpdateMPUTaskRequest) SetUnsubSpecShareScreenUsers(v []*string) *UpdateMPUTaskRequest {
	s.UnsubSpecShareScreenUsers = v
	return s
}

func (s *UpdateMPUTaskRequest) SetUserPanes(v []*UpdateMPUTaskRequestUserPanes) *UpdateMPUTaskRequest {
	s.UserPanes = v
	return s
}

func (s *UpdateMPUTaskRequest) SetWatermarks(v []*UpdateMPUTaskRequestWatermarks) *UpdateMPUTaskRequest {
	s.Watermarks = v
	return s
}

type UpdateMPUTaskRequestBackgrounds struct {
	Display *int32   `json:"Display,omitempty" xml:"Display,omitempty"`
	Height  *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Url     *string  `json:"Url,omitempty" xml:"Url,omitempty"`
	Width   *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	X       *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y       *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	ZOrder  *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s UpdateMPUTaskRequestBackgrounds) String() string {
	return tea.Prettify(s)
}

func (s UpdateMPUTaskRequestBackgrounds) GoString() string {
	return s.String()
}

func (s *UpdateMPUTaskRequestBackgrounds) SetDisplay(v int32) *UpdateMPUTaskRequestBackgrounds {
	s.Display = &v
	return s
}

func (s *UpdateMPUTaskRequestBackgrounds) SetHeight(v float32) *UpdateMPUTaskRequestBackgrounds {
	s.Height = &v
	return s
}

func (s *UpdateMPUTaskRequestBackgrounds) SetUrl(v string) *UpdateMPUTaskRequestBackgrounds {
	s.Url = &v
	return s
}

func (s *UpdateMPUTaskRequestBackgrounds) SetWidth(v float32) *UpdateMPUTaskRequestBackgrounds {
	s.Width = &v
	return s
}

func (s *UpdateMPUTaskRequestBackgrounds) SetX(v float32) *UpdateMPUTaskRequestBackgrounds {
	s.X = &v
	return s
}

func (s *UpdateMPUTaskRequestBackgrounds) SetY(v float32) *UpdateMPUTaskRequestBackgrounds {
	s.Y = &v
	return s
}

func (s *UpdateMPUTaskRequestBackgrounds) SetZOrder(v int32) *UpdateMPUTaskRequestBackgrounds {
	s.ZOrder = &v
	return s
}

type UpdateMPUTaskRequestClockWidgets struct {
	FontColor *int32   `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	FontSize  *int32   `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	FontType  *int32   `json:"FontType,omitempty" xml:"FontType,omitempty"`
	X         *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y         *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	ZOrder    *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s UpdateMPUTaskRequestClockWidgets) String() string {
	return tea.Prettify(s)
}

func (s UpdateMPUTaskRequestClockWidgets) GoString() string {
	return s.String()
}

func (s *UpdateMPUTaskRequestClockWidgets) SetFontColor(v int32) *UpdateMPUTaskRequestClockWidgets {
	s.FontColor = &v
	return s
}

func (s *UpdateMPUTaskRequestClockWidgets) SetFontSize(v int32) *UpdateMPUTaskRequestClockWidgets {
	s.FontSize = &v
	return s
}

func (s *UpdateMPUTaskRequestClockWidgets) SetFontType(v int32) *UpdateMPUTaskRequestClockWidgets {
	s.FontType = &v
	return s
}

func (s *UpdateMPUTaskRequestClockWidgets) SetX(v float32) *UpdateMPUTaskRequestClockWidgets {
	s.X = &v
	return s
}

func (s *UpdateMPUTaskRequestClockWidgets) SetY(v float32) *UpdateMPUTaskRequestClockWidgets {
	s.Y = &v
	return s
}

func (s *UpdateMPUTaskRequestClockWidgets) SetZOrder(v int32) *UpdateMPUTaskRequestClockWidgets {
	s.ZOrder = &v
	return s
}

type UpdateMPUTaskRequestUserPanes struct {
	Images      []*UpdateMPUTaskRequestUserPanesImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	PaneId      *int32                                 `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
	SegmentType *int32                                 `json:"SegmentType,omitempty" xml:"SegmentType,omitempty"`
	SourceType  *string                                `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Texts       []*UpdateMPUTaskRequestUserPanesTexts  `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
	UserId      *string                                `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateMPUTaskRequestUserPanes) String() string {
	return tea.Prettify(s)
}

func (s UpdateMPUTaskRequestUserPanes) GoString() string {
	return s.String()
}

func (s *UpdateMPUTaskRequestUserPanes) SetImages(v []*UpdateMPUTaskRequestUserPanesImages) *UpdateMPUTaskRequestUserPanes {
	s.Images = v
	return s
}

func (s *UpdateMPUTaskRequestUserPanes) SetPaneId(v int32) *UpdateMPUTaskRequestUserPanes {
	s.PaneId = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanes) SetSegmentType(v int32) *UpdateMPUTaskRequestUserPanes {
	s.SegmentType = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanes) SetSourceType(v string) *UpdateMPUTaskRequestUserPanes {
	s.SourceType = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanes) SetTexts(v []*UpdateMPUTaskRequestUserPanesTexts) *UpdateMPUTaskRequestUserPanes {
	s.Texts = v
	return s
}

func (s *UpdateMPUTaskRequestUserPanes) SetUserId(v string) *UpdateMPUTaskRequestUserPanes {
	s.UserId = &v
	return s
}

type UpdateMPUTaskRequestUserPanesImages struct {
	Display *int32   `json:"Display,omitempty" xml:"Display,omitempty"`
	Height  *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Url     *string  `json:"Url,omitempty" xml:"Url,omitempty"`
	Width   *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	X       *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y       *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	ZOrder  *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s UpdateMPUTaskRequestUserPanesImages) String() string {
	return tea.Prettify(s)
}

func (s UpdateMPUTaskRequestUserPanesImages) GoString() string {
	return s.String()
}

func (s *UpdateMPUTaskRequestUserPanesImages) SetDisplay(v int32) *UpdateMPUTaskRequestUserPanesImages {
	s.Display = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanesImages) SetHeight(v float32) *UpdateMPUTaskRequestUserPanesImages {
	s.Height = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanesImages) SetUrl(v string) *UpdateMPUTaskRequestUserPanesImages {
	s.Url = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanesImages) SetWidth(v float32) *UpdateMPUTaskRequestUserPanesImages {
	s.Width = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanesImages) SetX(v float32) *UpdateMPUTaskRequestUserPanesImages {
	s.X = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanesImages) SetY(v float32) *UpdateMPUTaskRequestUserPanesImages {
	s.Y = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanesImages) SetZOrder(v int32) *UpdateMPUTaskRequestUserPanesImages {
	s.ZOrder = &v
	return s
}

type UpdateMPUTaskRequestUserPanesTexts struct {
	FontColor *int32   `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	FontSize  *int32   `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	FontType  *int32   `json:"FontType,omitempty" xml:"FontType,omitempty"`
	Text      *string  `json:"Text,omitempty" xml:"Text,omitempty"`
	X         *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y         *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	ZOrder    *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s UpdateMPUTaskRequestUserPanesTexts) String() string {
	return tea.Prettify(s)
}

func (s UpdateMPUTaskRequestUserPanesTexts) GoString() string {
	return s.String()
}

func (s *UpdateMPUTaskRequestUserPanesTexts) SetFontColor(v int32) *UpdateMPUTaskRequestUserPanesTexts {
	s.FontColor = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanesTexts) SetFontSize(v int32) *UpdateMPUTaskRequestUserPanesTexts {
	s.FontSize = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanesTexts) SetFontType(v int32) *UpdateMPUTaskRequestUserPanesTexts {
	s.FontType = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanesTexts) SetText(v string) *UpdateMPUTaskRequestUserPanesTexts {
	s.Text = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanesTexts) SetX(v float32) *UpdateMPUTaskRequestUserPanesTexts {
	s.X = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanesTexts) SetY(v float32) *UpdateMPUTaskRequestUserPanesTexts {
	s.Y = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanesTexts) SetZOrder(v int32) *UpdateMPUTaskRequestUserPanesTexts {
	s.ZOrder = &v
	return s
}

type UpdateMPUTaskRequestWatermarks struct {
	Alpha   *float32 `json:"Alpha,omitempty" xml:"Alpha,omitempty"`
	Display *int32   `json:"Display,omitempty" xml:"Display,omitempty"`
	Height  *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Url     *string  `json:"Url,omitempty" xml:"Url,omitempty"`
	Width   *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	X       *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y       *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	ZOrder  *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s UpdateMPUTaskRequestWatermarks) String() string {
	return tea.Prettify(s)
}

func (s UpdateMPUTaskRequestWatermarks) GoString() string {
	return s.String()
}

func (s *UpdateMPUTaskRequestWatermarks) SetAlpha(v float32) *UpdateMPUTaskRequestWatermarks {
	s.Alpha = &v
	return s
}

func (s *UpdateMPUTaskRequestWatermarks) SetDisplay(v int32) *UpdateMPUTaskRequestWatermarks {
	s.Display = &v
	return s
}

func (s *UpdateMPUTaskRequestWatermarks) SetHeight(v float32) *UpdateMPUTaskRequestWatermarks {
	s.Height = &v
	return s
}

func (s *UpdateMPUTaskRequestWatermarks) SetUrl(v string) *UpdateMPUTaskRequestWatermarks {
	s.Url = &v
	return s
}

func (s *UpdateMPUTaskRequestWatermarks) SetWidth(v float32) *UpdateMPUTaskRequestWatermarks {
	s.Width = &v
	return s
}

func (s *UpdateMPUTaskRequestWatermarks) SetX(v float32) *UpdateMPUTaskRequestWatermarks {
	s.X = &v
	return s
}

func (s *UpdateMPUTaskRequestWatermarks) SetY(v float32) *UpdateMPUTaskRequestWatermarks {
	s.Y = &v
	return s
}

func (s *UpdateMPUTaskRequestWatermarks) SetZOrder(v int32) *UpdateMPUTaskRequestWatermarks {
	s.ZOrder = &v
	return s
}

type UpdateMPUTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateMPUTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMPUTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMPUTaskResponseBody) SetRequestId(v string) *UpdateMPUTaskResponseBody {
	s.RequestId = &v
	return s
}

type UpdateMPUTaskResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateMPUTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateMPUTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMPUTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateMPUTaskResponse) SetHeaders(v map[string]*string) *UpdateMPUTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateMPUTaskResponse) SetBody(v *UpdateMPUTaskResponseBody) *UpdateMPUTaskResponse {
	s.Body = v
	return s
}

type UpdateRecordTaskRequest struct {
	AppId                     *string                             `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId                 *string                             `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	LayoutIds                 []*int64                            `json:"LayoutIds,omitempty" xml:"LayoutIds,omitempty" type:"Repeated"`
	OwnerId                   *int64                              `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SubSpecAudioUsers         []*string                           `json:"SubSpecAudioUsers,omitempty" xml:"SubSpecAudioUsers,omitempty" type:"Repeated"`
	SubSpecCameraUsers        []*string                           `json:"SubSpecCameraUsers,omitempty" xml:"SubSpecCameraUsers,omitempty" type:"Repeated"`
	SubSpecShareScreenUsers   []*string                           `json:"SubSpecShareScreenUsers,omitempty" xml:"SubSpecShareScreenUsers,omitempty" type:"Repeated"`
	SubSpecUsers              []*string                           `json:"SubSpecUsers,omitempty" xml:"SubSpecUsers,omitempty" type:"Repeated"`
	TaskId                    *string                             `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TemplateId                *string                             `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	UnsubSpecAudioUsers       []*string                           `json:"UnsubSpecAudioUsers,omitempty" xml:"UnsubSpecAudioUsers,omitempty" type:"Repeated"`
	UnsubSpecCameraUsers      []*string                           `json:"UnsubSpecCameraUsers,omitempty" xml:"UnsubSpecCameraUsers,omitempty" type:"Repeated"`
	UnsubSpecShareScreenUsers []*string                           `json:"UnsubSpecShareScreenUsers,omitempty" xml:"UnsubSpecShareScreenUsers,omitempty" type:"Repeated"`
	UserPanes                 []*UpdateRecordTaskRequestUserPanes `json:"UserPanes,omitempty" xml:"UserPanes,omitempty" type:"Repeated"`
}

func (s UpdateRecordTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecordTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateRecordTaskRequest) SetAppId(v string) *UpdateRecordTaskRequest {
	s.AppId = &v
	return s
}

func (s *UpdateRecordTaskRequest) SetChannelId(v string) *UpdateRecordTaskRequest {
	s.ChannelId = &v
	return s
}

func (s *UpdateRecordTaskRequest) SetLayoutIds(v []*int64) *UpdateRecordTaskRequest {
	s.LayoutIds = v
	return s
}

func (s *UpdateRecordTaskRequest) SetOwnerId(v int64) *UpdateRecordTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateRecordTaskRequest) SetSubSpecAudioUsers(v []*string) *UpdateRecordTaskRequest {
	s.SubSpecAudioUsers = v
	return s
}

func (s *UpdateRecordTaskRequest) SetSubSpecCameraUsers(v []*string) *UpdateRecordTaskRequest {
	s.SubSpecCameraUsers = v
	return s
}

func (s *UpdateRecordTaskRequest) SetSubSpecShareScreenUsers(v []*string) *UpdateRecordTaskRequest {
	s.SubSpecShareScreenUsers = v
	return s
}

func (s *UpdateRecordTaskRequest) SetSubSpecUsers(v []*string) *UpdateRecordTaskRequest {
	s.SubSpecUsers = v
	return s
}

func (s *UpdateRecordTaskRequest) SetTaskId(v string) *UpdateRecordTaskRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateRecordTaskRequest) SetTemplateId(v string) *UpdateRecordTaskRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateRecordTaskRequest) SetUnsubSpecAudioUsers(v []*string) *UpdateRecordTaskRequest {
	s.UnsubSpecAudioUsers = v
	return s
}

func (s *UpdateRecordTaskRequest) SetUnsubSpecCameraUsers(v []*string) *UpdateRecordTaskRequest {
	s.UnsubSpecCameraUsers = v
	return s
}

func (s *UpdateRecordTaskRequest) SetUnsubSpecShareScreenUsers(v []*string) *UpdateRecordTaskRequest {
	s.UnsubSpecShareScreenUsers = v
	return s
}

func (s *UpdateRecordTaskRequest) SetUserPanes(v []*UpdateRecordTaskRequestUserPanes) *UpdateRecordTaskRequest {
	s.UserPanes = v
	return s
}

type UpdateRecordTaskRequestUserPanes struct {
	Images     []*UpdateRecordTaskRequestUserPanesImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	PaneId     *int32                                    `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
	SourceType *string                                   `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Texts      []*UpdateRecordTaskRequestUserPanesTexts  `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
	UserId     *string                                   `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s UpdateRecordTaskRequestUserPanes) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecordTaskRequestUserPanes) GoString() string {
	return s.String()
}

func (s *UpdateRecordTaskRequestUserPanes) SetImages(v []*UpdateRecordTaskRequestUserPanesImages) *UpdateRecordTaskRequestUserPanes {
	s.Images = v
	return s
}

func (s *UpdateRecordTaskRequestUserPanes) SetPaneId(v int32) *UpdateRecordTaskRequestUserPanes {
	s.PaneId = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanes) SetSourceType(v string) *UpdateRecordTaskRequestUserPanes {
	s.SourceType = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanes) SetTexts(v []*UpdateRecordTaskRequestUserPanesTexts) *UpdateRecordTaskRequestUserPanes {
	s.Texts = v
	return s
}

func (s *UpdateRecordTaskRequestUserPanes) SetUserId(v string) *UpdateRecordTaskRequestUserPanes {
	s.UserId = &v
	return s
}

type UpdateRecordTaskRequestUserPanesImages struct {
	Display *int32   `json:"Display,omitempty" xml:"Display,omitempty"`
	Height  *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Url     *string  `json:"Url,omitempty" xml:"Url,omitempty"`
	Width   *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	X       *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y       *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	ZOrder  *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s UpdateRecordTaskRequestUserPanesImages) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecordTaskRequestUserPanesImages) GoString() string {
	return s.String()
}

func (s *UpdateRecordTaskRequestUserPanesImages) SetDisplay(v int32) *UpdateRecordTaskRequestUserPanesImages {
	s.Display = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanesImages) SetHeight(v float32) *UpdateRecordTaskRequestUserPanesImages {
	s.Height = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanesImages) SetUrl(v string) *UpdateRecordTaskRequestUserPanesImages {
	s.Url = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanesImages) SetWidth(v float32) *UpdateRecordTaskRequestUserPanesImages {
	s.Width = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanesImages) SetX(v float32) *UpdateRecordTaskRequestUserPanesImages {
	s.X = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanesImages) SetY(v float32) *UpdateRecordTaskRequestUserPanesImages {
	s.Y = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanesImages) SetZOrder(v int32) *UpdateRecordTaskRequestUserPanesImages {
	s.ZOrder = &v
	return s
}

type UpdateRecordTaskRequestUserPanesTexts struct {
	FontColor *int32   `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	FontSize  *int32   `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	FontType  *int32   `json:"FontType,omitempty" xml:"FontType,omitempty"`
	Text      *string  `json:"Text,omitempty" xml:"Text,omitempty"`
	X         *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y         *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	ZOrder    *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s UpdateRecordTaskRequestUserPanesTexts) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecordTaskRequestUserPanesTexts) GoString() string {
	return s.String()
}

func (s *UpdateRecordTaskRequestUserPanesTexts) SetFontColor(v int32) *UpdateRecordTaskRequestUserPanesTexts {
	s.FontColor = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanesTexts) SetFontSize(v int32) *UpdateRecordTaskRequestUserPanesTexts {
	s.FontSize = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanesTexts) SetFontType(v int32) *UpdateRecordTaskRequestUserPanesTexts {
	s.FontType = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanesTexts) SetText(v string) *UpdateRecordTaskRequestUserPanesTexts {
	s.Text = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanesTexts) SetX(v float32) *UpdateRecordTaskRequestUserPanesTexts {
	s.X = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanesTexts) SetY(v float32) *UpdateRecordTaskRequestUserPanesTexts {
	s.Y = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanesTexts) SetZOrder(v int32) *UpdateRecordTaskRequestUserPanesTexts {
	s.ZOrder = &v
	return s
}

type UpdateRecordTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRecordTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecordTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRecordTaskResponseBody) SetRequestId(v string) *UpdateRecordTaskResponseBody {
	s.RequestId = &v
	return s
}

type UpdateRecordTaskResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateRecordTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateRecordTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecordTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateRecordTaskResponse) SetHeaders(v map[string]*string) *UpdateRecordTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateRecordTaskResponse) SetBody(v *UpdateRecordTaskResponseBody) *UpdateRecordTaskResponse {
	s.Body = v
	return s
}

type UpdateRecordTemplateRequest struct {
	AppId              *string                                    `json:"AppId,omitempty" xml:"AppId,omitempty"`
	BackgroundColor    *int32                                     `json:"BackgroundColor,omitempty" xml:"BackgroundColor,omitempty"`
	Backgrounds        []*UpdateRecordTemplateRequestBackgrounds  `json:"Backgrounds,omitempty" xml:"Backgrounds,omitempty" type:"Repeated"`
	ClockWidgets       []*UpdateRecordTemplateRequestClockWidgets `json:"ClockWidgets,omitempty" xml:"ClockWidgets,omitempty" type:"Repeated"`
	DelayStopTime      *int32                                     `json:"DelayStopTime,omitempty" xml:"DelayStopTime,omitempty"`
	EnableM3u8DateTime *bool                                      `json:"EnableM3u8DateTime,omitempty" xml:"EnableM3u8DateTime,omitempty"`
	FileSplitInterval  *int32                                     `json:"FileSplitInterval,omitempty" xml:"FileSplitInterval,omitempty"`
	Formats            []*string                                  `json:"Formats,omitempty" xml:"Formats,omitempty" type:"Repeated"`
	HttpCallbackUrl    *string                                    `json:"HttpCallbackUrl,omitempty" xml:"HttpCallbackUrl,omitempty"`
	LayoutIds          []*int64                                   `json:"LayoutIds,omitempty" xml:"LayoutIds,omitempty" type:"Repeated"`
	MediaEncode        *int32                                     `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	MnsQueue           *string                                    `json:"MnsQueue,omitempty" xml:"MnsQueue,omitempty"`
	Name               *string                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	OssBucket          *string                                    `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssFilePrefix      *string                                    `json:"OssFilePrefix,omitempty" xml:"OssFilePrefix,omitempty"`
	OwnerId            *int64                                     `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	TaskProfile        *string                                    `json:"TaskProfile,omitempty" xml:"TaskProfile,omitempty"`
	TemplateId         *string                                    `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Watermarks         []*UpdateRecordTemplateRequestWatermarks   `json:"Watermarks,omitempty" xml:"Watermarks,omitempty" type:"Repeated"`
}

func (s UpdateRecordTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecordTemplateRequest) GoString() string {
	return s.String()
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

type UpdateRecordTemplateRequestBackgrounds struct {
	Display *int32   `json:"Display,omitempty" xml:"Display,omitempty"`
	Height  *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Url     *string  `json:"Url,omitempty" xml:"Url,omitempty"`
	Width   *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	X       *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y       *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	ZOrder  *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s UpdateRecordTemplateRequestBackgrounds) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecordTemplateRequestBackgrounds) GoString() string {
	return s.String()
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

type UpdateRecordTemplateRequestClockWidgets struct {
	FontColor *int32   `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	FontSize  *int32   `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	FontType  *int32   `json:"FontType,omitempty" xml:"FontType,omitempty"`
	X         *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y         *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	ZOrder    *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s UpdateRecordTemplateRequestClockWidgets) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecordTemplateRequestClockWidgets) GoString() string {
	return s.String()
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

type UpdateRecordTemplateRequestWatermarks struct {
	Alpha   *float32 `json:"Alpha,omitempty" xml:"Alpha,omitempty"`
	Display *int32   `json:"Display,omitempty" xml:"Display,omitempty"`
	Height  *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Url     *string  `json:"Url,omitempty" xml:"Url,omitempty"`
	Width   *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	X       *float32 `json:"X,omitempty" xml:"X,omitempty"`
	Y       *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	ZOrder  *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s UpdateRecordTemplateRequestWatermarks) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecordTemplateRequestWatermarks) GoString() string {
	return s.String()
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

type UpdateRecordTemplateResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s UpdateRecordTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecordTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRecordTemplateResponseBody) SetRequestId(v string) *UpdateRecordTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRecordTemplateResponseBody) SetTemplateId(v string) *UpdateRecordTemplateResponseBody {
	s.TemplateId = &v
	return s
}

type UpdateRecordTemplateResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateRecordTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateRecordTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecordTemplateResponse) GoString() string {
	return s.String()
}

func (s *UpdateRecordTemplateResponse) SetHeaders(v map[string]*string) *UpdateRecordTemplateResponse {
	s.Headers = v
	return s
}

func (s *UpdateRecordTemplateResponse) SetBody(v *UpdateRecordTemplateResponseBody) *UpdateRecordTemplateResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("central")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("rtc"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddRecordTemplateWithOptions(request *AddRecordTemplateRequest, runtime *util.RuntimeOptions) (_result *AddRecordTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddRecordTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddRecordTemplate"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddRecordTemplate(request *AddRecordTemplateRequest) (_result *AddRecordTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddRecordTemplateResponse{}
	_body, _err := client.AddRecordTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAutoLiveStreamRuleWithOptions(request *CreateAutoLiveStreamRuleRequest, runtime *util.RuntimeOptions) (_result *CreateAutoLiveStreamRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAutoLiveStreamRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAutoLiveStreamRule"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAutoLiveStreamRule(request *CreateAutoLiveStreamRuleRequest) (_result *CreateAutoLiveStreamRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAutoLiveStreamRuleResponse{}
	_body, _err := client.CreateAutoLiveStreamRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateEventSubscribeWithOptions(request *CreateEventSubscribeRequest, runtime *util.RuntimeOptions) (_result *CreateEventSubscribeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateEventSubscribeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateEventSubscribe"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateEventSubscribe(request *CreateEventSubscribeRequest) (_result *CreateEventSubscribeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateEventSubscribeResponse{}
	_body, _err := client.CreateEventSubscribeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMPULayoutWithOptions(request *CreateMPULayoutRequest, runtime *util.RuntimeOptions) (_result *CreateMPULayoutResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateMPULayoutResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateMPULayout"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMPULayout(request *CreateMPULayoutRequest) (_result *CreateMPULayoutResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMPULayoutResponse{}
	_body, _err := client.CreateMPULayoutWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRecordIndexFileWithOptions(request *CreateRecordIndexFileRequest, runtime *util.RuntimeOptions) (_result *CreateRecordIndexFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateRecordIndexFileResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateRecordIndexFile"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRecordIndexFile(request *CreateRecordIndexFileRequest) (_result *CreateRecordIndexFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateRecordIndexFileResponse{}
	_body, _err := client.CreateRecordIndexFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAutoLiveStreamRuleWithOptions(request *DeleteAutoLiveStreamRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteAutoLiveStreamRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAutoLiveStreamRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAutoLiveStreamRule"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAutoLiveStreamRule(request *DeleteAutoLiveStreamRuleRequest) (_result *DeleteAutoLiveStreamRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAutoLiveStreamRuleResponse{}
	_body, _err := client.DeleteAutoLiveStreamRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEventSubscribeWithOptions(request *DeleteEventSubscribeRequest, runtime *util.RuntimeOptions) (_result *DeleteEventSubscribeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteEventSubscribeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteEventSubscribe"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEventSubscribe(request *DeleteEventSubscribeRequest) (_result *DeleteEventSubscribeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteEventSubscribeResponse{}
	_body, _err := client.DeleteEventSubscribeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMPULayoutWithOptions(request *DeleteMPULayoutRequest, runtime *util.RuntimeOptions) (_result *DeleteMPULayoutResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteMPULayoutResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteMPULayout"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMPULayout(request *DeleteMPULayoutRequest) (_result *DeleteMPULayoutResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMPULayoutResponse{}
	_body, _err := client.DeleteMPULayoutWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRecordTemplateWithOptions(request *DeleteRecordTemplateRequest, runtime *util.RuntimeOptions) (_result *DeleteRecordTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteRecordTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteRecordTemplate"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRecordTemplate(request *DeleteRecordTemplateRequest) (_result *DeleteRecordTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteRecordTemplateResponse{}
	_body, _err := client.DeleteRecordTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAutoLiveStreamRuleWithOptions(request *DescribeAutoLiveStreamRuleRequest, runtime *util.RuntimeOptions) (_result *DescribeAutoLiveStreamRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAutoLiveStreamRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeAutoLiveStreamRule"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAutoLiveStreamRule(request *DescribeAutoLiveStreamRuleRequest) (_result *DescribeAutoLiveStreamRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAutoLiveStreamRuleResponse{}
	_body, _err := client.DescribeAutoLiveStreamRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeChannelParticipantsWithOptions(request *DescribeChannelParticipantsRequest, runtime *util.RuntimeOptions) (_result *DescribeChannelParticipantsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeChannelParticipantsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeChannelParticipants"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeChannelParticipants(request *DescribeChannelParticipantsRequest) (_result *DescribeChannelParticipantsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeChannelParticipantsResponse{}
	_body, _err := client.DescribeChannelParticipantsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeChannelUsersWithOptions(request *DescribeChannelUsersRequest, runtime *util.RuntimeOptions) (_result *DescribeChannelUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeChannelUsersResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeChannelUsers"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeChannelUsers(request *DescribeChannelUsersRequest) (_result *DescribeChannelUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeChannelUsersResponse{}
	_body, _err := client.DescribeChannelUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMPULayoutInfoListWithOptions(request *DescribeMPULayoutInfoListRequest, runtime *util.RuntimeOptions) (_result *DescribeMPULayoutInfoListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMPULayoutInfoListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMPULayoutInfoList"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMPULayoutInfoList(request *DescribeMPULayoutInfoListRequest) (_result *DescribeMPULayoutInfoListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMPULayoutInfoListResponse{}
	_body, _err := client.DescribeMPULayoutInfoListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRecordFilesWithOptions(request *DescribeRecordFilesRequest, runtime *util.RuntimeOptions) (_result *DescribeRecordFilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRecordFilesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRecordFiles"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRecordFiles(request *DescribeRecordFilesRequest) (_result *DescribeRecordFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRecordFilesResponse{}
	_body, _err := client.DescribeRecordFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRecordTasksWithOptions(request *DescribeRecordTasksRequest, runtime *util.RuntimeOptions) (_result *DescribeRecordTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRecordTasksResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRecordTasks"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRecordTasks(request *DescribeRecordTasksRequest) (_result *DescribeRecordTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRecordTasksResponse{}
	_body, _err := client.DescribeRecordTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRecordTemplatesWithOptions(request *DescribeRecordTemplatesRequest, runtime *util.RuntimeOptions) (_result *DescribeRecordTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRecordTemplatesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRecordTemplates"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRecordTemplates(request *DescribeRecordTemplatesRequest) (_result *DescribeRecordTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRecordTemplatesResponse{}
	_body, _err := client.DescribeRecordTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserInfoInChannelWithOptions(request *DescribeUserInfoInChannelRequest, runtime *util.RuntimeOptions) (_result *DescribeUserInfoInChannelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeUserInfoInChannelResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeUserInfoInChannel"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserInfoInChannel(request *DescribeUserInfoInChannelRequest) (_result *DescribeUserInfoInChannelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserInfoInChannelResponse{}
	_body, _err := client.DescribeUserInfoInChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableAutoLiveStreamRuleWithOptions(request *DisableAutoLiveStreamRuleRequest, runtime *util.RuntimeOptions) (_result *DisableAutoLiveStreamRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DisableAutoLiveStreamRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisableAutoLiveStreamRule"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableAutoLiveStreamRule(request *DisableAutoLiveStreamRuleRequest) (_result *DisableAutoLiveStreamRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableAutoLiveStreamRuleResponse{}
	_body, _err := client.DisableAutoLiveStreamRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableAutoLiveStreamRuleWithOptions(request *EnableAutoLiveStreamRuleRequest, runtime *util.RuntimeOptions) (_result *EnableAutoLiveStreamRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableAutoLiveStreamRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableAutoLiveStreamRule"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableAutoLiveStreamRule(request *EnableAutoLiveStreamRuleRequest) (_result *EnableAutoLiveStreamRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableAutoLiveStreamRuleResponse{}
	_body, _err := client.EnableAutoLiveStreamRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetMPUTaskStatusWithOptions(request *GetMPUTaskStatusRequest, runtime *util.RuntimeOptions) (_result *GetMPUTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetMPUTaskStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetMPUTaskStatus"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetMPUTaskStatus(request *GetMPUTaskStatusRequest) (_result *GetMPUTaskStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetMPUTaskStatusResponse{}
	_body, _err := client.GetMPUTaskStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyMPULayoutWithOptions(request *ModifyMPULayoutRequest, runtime *util.RuntimeOptions) (_result *ModifyMPULayoutResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyMPULayoutResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyMPULayout"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyMPULayout(request *ModifyMPULayoutRequest) (_result *ModifyMPULayoutResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyMPULayoutResponse{}
	_body, _err := client.ModifyMPULayoutWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveTerminalsWithOptions(request *RemoveTerminalsRequest, runtime *util.RuntimeOptions) (_result *RemoveTerminalsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveTerminalsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveTerminals"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveTerminals(request *RemoveTerminalsRequest) (_result *RemoveTerminalsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveTerminalsResponse{}
	_body, _err := client.RemoveTerminalsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartMPUTaskWithOptions(request *StartMPUTaskRequest, runtime *util.RuntimeOptions) (_result *StartMPUTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartMPUTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartMPUTask"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartMPUTask(request *StartMPUTaskRequest) (_result *StartMPUTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartMPUTaskResponse{}
	_body, _err := client.StartMPUTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartRecordTaskWithOptions(request *StartRecordTaskRequest, runtime *util.RuntimeOptions) (_result *StartRecordTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartRecordTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartRecordTask"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartRecordTask(request *StartRecordTaskRequest) (_result *StartRecordTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartRecordTaskResponse{}
	_body, _err := client.StartRecordTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopChannelUserPublishWithOptions(request *StopChannelUserPublishRequest, runtime *util.RuntimeOptions) (_result *StopChannelUserPublishResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StopChannelUserPublishResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopChannelUserPublish"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopChannelUserPublish(request *StopChannelUserPublishRequest) (_result *StopChannelUserPublishResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopChannelUserPublishResponse{}
	_body, _err := client.StopChannelUserPublishWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopMPUTaskWithOptions(request *StopMPUTaskRequest, runtime *util.RuntimeOptions) (_result *StopMPUTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StopMPUTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopMPUTask"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopMPUTask(request *StopMPUTaskRequest) (_result *StopMPUTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopMPUTaskResponse{}
	_body, _err := client.StopMPUTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopRecordTaskWithOptions(request *StopRecordTaskRequest, runtime *util.RuntimeOptions) (_result *StopRecordTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StopRecordTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopRecordTask"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopRecordTask(request *StopRecordTaskRequest) (_result *StopRecordTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopRecordTaskResponse{}
	_body, _err := client.StopRecordTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateAutoLiveStreamRuleWithOptions(request *UpdateAutoLiveStreamRuleRequest, runtime *util.RuntimeOptions) (_result *UpdateAutoLiveStreamRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateAutoLiveStreamRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateAutoLiveStreamRule"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateAutoLiveStreamRule(request *UpdateAutoLiveStreamRuleRequest) (_result *UpdateAutoLiveStreamRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateAutoLiveStreamRuleResponse{}
	_body, _err := client.UpdateAutoLiveStreamRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateMPUTaskWithOptions(request *UpdateMPUTaskRequest, runtime *util.RuntimeOptions) (_result *UpdateMPUTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateMPUTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateMPUTask"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateMPUTask(request *UpdateMPUTaskRequest) (_result *UpdateMPUTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateMPUTaskResponse{}
	_body, _err := client.UpdateMPUTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateRecordTaskWithOptions(request *UpdateRecordTaskRequest, runtime *util.RuntimeOptions) (_result *UpdateRecordTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateRecordTaskResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateRecordTask"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateRecordTask(request *UpdateRecordTaskRequest) (_result *UpdateRecordTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateRecordTaskResponse{}
	_body, _err := client.UpdateRecordTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateRecordTemplateWithOptions(request *UpdateRecordTemplateRequest, runtime *util.RuntimeOptions) (_result *UpdateRecordTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateRecordTemplateResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateRecordTemplate"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateRecordTemplate(request *UpdateRecordTemplateRequest) (_result *UpdateRecordTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateRecordTemplateResponse{}
	_body, _err := client.UpdateRecordTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
