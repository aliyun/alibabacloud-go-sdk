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
	OwnerId           *int64                                  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog           *string                                 `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId             *string                                 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Name              *string                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	TaskProfile       *string                                 `json:"TaskProfile,omitempty" xml:"TaskProfile,omitempty"`
	MediaEncode       *int32                                  `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	BackgroundColor   *int32                                  `json:"BackgroundColor,omitempty" xml:"BackgroundColor,omitempty"`
	OssBucket         *string                                 `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssFilePrefix     *string                                 `json:"OssFilePrefix,omitempty" xml:"OssFilePrefix,omitempty"`
	FileSplitInterval *int32                                  `json:"FileSplitInterval,omitempty" xml:"FileSplitInterval,omitempty"`
	DelayStopTime     *int32                                  `json:"DelayStopTime,omitempty" xml:"DelayStopTime,omitempty"`
	MnsQueue          *string                                 `json:"MnsQueue,omitempty" xml:"MnsQueue,omitempty"`
	HttpCallbackUrl   *string                                 `json:"HttpCallbackUrl,omitempty" xml:"HttpCallbackUrl,omitempty"`
	LayoutIds         []*int                                  `json:"LayoutIds,omitempty" xml:"LayoutIds,omitempty" type:"Repeated"`
	Formats           []*string                               `json:"Formats,omitempty" xml:"Formats,omitempty" type:"Repeated"`
	Backgrounds       []*AddRecordTemplateRequestBackgrounds  `json:"Backgrounds,omitempty" xml:"Backgrounds,omitempty" type:"Repeated"`
	Watermarks        []*AddRecordTemplateRequestWatermarks   `json:"Watermarks,omitempty" xml:"Watermarks,omitempty" type:"Repeated"`
	ClockWidgets      []*AddRecordTemplateRequestClockWidgets `json:"ClockWidgets,omitempty" xml:"ClockWidgets,omitempty" type:"Repeated"`
}

func (s AddRecordTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s AddRecordTemplateRequest) GoString() string {
	return s.String()
}

func (s *AddRecordTemplateRequest) SetOwnerId(v int64) *AddRecordTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *AddRecordTemplateRequest) SetShowLog(v string) *AddRecordTemplateRequest {
	s.ShowLog = &v
	return s
}

func (s *AddRecordTemplateRequest) SetAppId(v string) *AddRecordTemplateRequest {
	s.AppId = &v
	return s
}

func (s *AddRecordTemplateRequest) SetName(v string) *AddRecordTemplateRequest {
	s.Name = &v
	return s
}

func (s *AddRecordTemplateRequest) SetTaskProfile(v string) *AddRecordTemplateRequest {
	s.TaskProfile = &v
	return s
}

func (s *AddRecordTemplateRequest) SetMediaEncode(v int32) *AddRecordTemplateRequest {
	s.MediaEncode = &v
	return s
}

func (s *AddRecordTemplateRequest) SetBackgroundColor(v int32) *AddRecordTemplateRequest {
	s.BackgroundColor = &v
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

func (s *AddRecordTemplateRequest) SetFileSplitInterval(v int32) *AddRecordTemplateRequest {
	s.FileSplitInterval = &v
	return s
}

func (s *AddRecordTemplateRequest) SetDelayStopTime(v int32) *AddRecordTemplateRequest {
	s.DelayStopTime = &v
	return s
}

func (s *AddRecordTemplateRequest) SetMnsQueue(v string) *AddRecordTemplateRequest {
	s.MnsQueue = &v
	return s
}

func (s *AddRecordTemplateRequest) SetHttpCallbackUrl(v string) *AddRecordTemplateRequest {
	s.HttpCallbackUrl = &v
	return s
}

func (s *AddRecordTemplateRequest) SetLayoutIds(v []*int) *AddRecordTemplateRequest {
	s.LayoutIds = v
	return s
}

func (s *AddRecordTemplateRequest) SetFormats(v []*string) *AddRecordTemplateRequest {
	s.Formats = v
	return s
}

func (s *AddRecordTemplateRequest) SetBackgrounds(v []*AddRecordTemplateRequestBackgrounds) *AddRecordTemplateRequest {
	s.Backgrounds = v
	return s
}

func (s *AddRecordTemplateRequest) SetWatermarks(v []*AddRecordTemplateRequestWatermarks) *AddRecordTemplateRequest {
	s.Watermarks = v
	return s
}

func (s *AddRecordTemplateRequest) SetClockWidgets(v []*AddRecordTemplateRequestClockWidgets) *AddRecordTemplateRequest {
	s.ClockWidgets = v
	return s
}

type AddRecordTemplateRequestBackgrounds struct {
	Width   *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height  *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Y       *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	Url     *string  `json:"Url,omitempty" xml:"Url,omitempty"`
	Display *int32   `json:"Display,omitempty" xml:"Display,omitempty"`
	ZOrder  *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
	X       *float32 `json:"X,omitempty" xml:"X,omitempty"`
}

func (s AddRecordTemplateRequestBackgrounds) String() string {
	return tea.Prettify(s)
}

func (s AddRecordTemplateRequestBackgrounds) GoString() string {
	return s.String()
}

func (s *AddRecordTemplateRequestBackgrounds) SetWidth(v float32) *AddRecordTemplateRequestBackgrounds {
	s.Width = &v
	return s
}

func (s *AddRecordTemplateRequestBackgrounds) SetHeight(v float32) *AddRecordTemplateRequestBackgrounds {
	s.Height = &v
	return s
}

func (s *AddRecordTemplateRequestBackgrounds) SetY(v float32) *AddRecordTemplateRequestBackgrounds {
	s.Y = &v
	return s
}

func (s *AddRecordTemplateRequestBackgrounds) SetUrl(v string) *AddRecordTemplateRequestBackgrounds {
	s.Url = &v
	return s
}

func (s *AddRecordTemplateRequestBackgrounds) SetDisplay(v int32) *AddRecordTemplateRequestBackgrounds {
	s.Display = &v
	return s
}

func (s *AddRecordTemplateRequestBackgrounds) SetZOrder(v int32) *AddRecordTemplateRequestBackgrounds {
	s.ZOrder = &v
	return s
}

func (s *AddRecordTemplateRequestBackgrounds) SetX(v float32) *AddRecordTemplateRequestBackgrounds {
	s.X = &v
	return s
}

type AddRecordTemplateRequestWatermarks struct {
	Alpha   *float32 `json:"Alpha,omitempty" xml:"Alpha,omitempty"`
	Width   *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height  *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Y       *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	Url     *string  `json:"Url,omitempty" xml:"Url,omitempty"`
	Display *int32   `json:"Display,omitempty" xml:"Display,omitempty"`
	ZOrder  *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
	X       *float32 `json:"X,omitempty" xml:"X,omitempty"`
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

func (s *AddRecordTemplateRequestWatermarks) SetWidth(v float32) *AddRecordTemplateRequestWatermarks {
	s.Width = &v
	return s
}

func (s *AddRecordTemplateRequestWatermarks) SetHeight(v float32) *AddRecordTemplateRequestWatermarks {
	s.Height = &v
	return s
}

func (s *AddRecordTemplateRequestWatermarks) SetY(v float32) *AddRecordTemplateRequestWatermarks {
	s.Y = &v
	return s
}

func (s *AddRecordTemplateRequestWatermarks) SetUrl(v string) *AddRecordTemplateRequestWatermarks {
	s.Url = &v
	return s
}

func (s *AddRecordTemplateRequestWatermarks) SetDisplay(v int32) *AddRecordTemplateRequestWatermarks {
	s.Display = &v
	return s
}

func (s *AddRecordTemplateRequestWatermarks) SetZOrder(v int32) *AddRecordTemplateRequestWatermarks {
	s.ZOrder = &v
	return s
}

func (s *AddRecordTemplateRequestWatermarks) SetX(v float32) *AddRecordTemplateRequestWatermarks {
	s.X = &v
	return s
}

type AddRecordTemplateRequestClockWidgets struct {
	FontType  *int32   `json:"FontType,omitempty" xml:"FontType,omitempty"`
	FontColor *int32   `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	Y         *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	ZOrder    *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
	X         *float32 `json:"X,omitempty" xml:"X,omitempty"`
	FontSize  *int32   `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
}

func (s AddRecordTemplateRequestClockWidgets) String() string {
	return tea.Prettify(s)
}

func (s AddRecordTemplateRequestClockWidgets) GoString() string {
	return s.String()
}

func (s *AddRecordTemplateRequestClockWidgets) SetFontType(v int32) *AddRecordTemplateRequestClockWidgets {
	s.FontType = &v
	return s
}

func (s *AddRecordTemplateRequestClockWidgets) SetFontColor(v int32) *AddRecordTemplateRequestClockWidgets {
	s.FontColor = &v
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

func (s *AddRecordTemplateRequestClockWidgets) SetX(v float32) *AddRecordTemplateRequestClockWidgets {
	s.X = &v
	return s
}

func (s *AddRecordTemplateRequestClockWidgets) SetFontSize(v int32) *AddRecordTemplateRequestClockWidgets {
	s.FontSize = &v
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
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog    *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	PlayDomain *string `json:"PlayDomain,omitempty" xml:"PlayDomain,omitempty"`
	RuleName   *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	CallBack   *string `json:"CallBack,omitempty" xml:"CallBack,omitempty"`
}

func (s CreateAutoLiveStreamRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAutoLiveStreamRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateAutoLiveStreamRuleRequest) SetOwnerId(v int64) *CreateAutoLiveStreamRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAutoLiveStreamRuleRequest) SetShowLog(v string) *CreateAutoLiveStreamRuleRequest {
	s.ShowLog = &v
	return s
}

func (s *CreateAutoLiveStreamRuleRequest) SetAppId(v string) *CreateAutoLiveStreamRuleRequest {
	s.AppId = &v
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

func (s *CreateAutoLiveStreamRuleRequest) SetCallBack(v string) *CreateAutoLiveStreamRuleRequest {
	s.CallBack = &v
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

type CreateChannelRequest struct {
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog   *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
}

func (s CreateChannelRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateChannelRequest) GoString() string {
	return s.String()
}

func (s *CreateChannelRequest) SetOwnerId(v int64) *CreateChannelRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateChannelRequest) SetShowLog(v string) *CreateChannelRequest {
	s.ShowLog = &v
	return s
}

func (s *CreateChannelRequest) SetAppId(v string) *CreateChannelRequest {
	s.AppId = &v
	return s
}

func (s *CreateChannelRequest) SetChannelId(v string) *CreateChannelRequest {
	s.ChannelId = &v
	return s
}

type CreateChannelResponseBody struct {
	Nonce      *string `json:"Nonce,omitempty" xml:"Nonce,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ChannelKey *string `json:"ChannelKey,omitempty" xml:"ChannelKey,omitempty"`
	Timestamp  *int32  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s CreateChannelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateChannelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateChannelResponseBody) SetNonce(v string) *CreateChannelResponseBody {
	s.Nonce = &v
	return s
}

func (s *CreateChannelResponseBody) SetRequestId(v string) *CreateChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateChannelResponseBody) SetChannelKey(v string) *CreateChannelResponseBody {
	s.ChannelKey = &v
	return s
}

func (s *CreateChannelResponseBody) SetTimestamp(v int32) *CreateChannelResponseBody {
	s.Timestamp = &v
	return s
}

type CreateChannelResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateChannelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateChannelResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateChannelResponse) GoString() string {
	return s.String()
}

func (s *CreateChannelResponse) SetHeaders(v map[string]*string) *CreateChannelResponse {
	s.Headers = v
	return s
}

func (s *CreateChannelResponse) SetBody(v *CreateChannelResponseBody) *CreateChannelResponse {
	s.Body = v
	return s
}

type CreateConferenceRequest struct {
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog        *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ConferenceName *string `json:"ConferenceName,omitempty" xml:"ConferenceName,omitempty"`
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
	RemindNotice   *int32  `json:"RemindNotice,omitempty" xml:"RemindNotice,omitempty"`
}

func (s CreateConferenceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateConferenceRequest) GoString() string {
	return s.String()
}

func (s *CreateConferenceRequest) SetOwnerId(v int64) *CreateConferenceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateConferenceRequest) SetShowLog(v string) *CreateConferenceRequest {
	s.ShowLog = &v
	return s
}

func (s *CreateConferenceRequest) SetAppId(v string) *CreateConferenceRequest {
	s.AppId = &v
	return s
}

func (s *CreateConferenceRequest) SetConferenceName(v string) *CreateConferenceRequest {
	s.ConferenceName = &v
	return s
}

func (s *CreateConferenceRequest) SetClientToken(v string) *CreateConferenceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateConferenceRequest) SetStartTime(v string) *CreateConferenceRequest {
	s.StartTime = &v
	return s
}

func (s *CreateConferenceRequest) SetType(v string) *CreateConferenceRequest {
	s.Type = &v
	return s
}

func (s *CreateConferenceRequest) SetRemindNotice(v int32) *CreateConferenceRequest {
	s.RemindNotice = &v
	return s
}

type CreateConferenceResponseBody struct {
	AuthInfo     *CreateConferenceResponseBodyAuthInfo `json:"AuthInfo,omitempty" xml:"AuthInfo,omitempty" type:"Struct"`
	RequestId    *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ConferenceId *string                               `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
}

func (s CreateConferenceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateConferenceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConferenceResponseBody) SetAuthInfo(v *CreateConferenceResponseBodyAuthInfo) *CreateConferenceResponseBody {
	s.AuthInfo = v
	return s
}

func (s *CreateConferenceResponseBody) SetRequestId(v string) *CreateConferenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateConferenceResponseBody) SetConferenceId(v string) *CreateConferenceResponseBody {
	s.ConferenceId = &v
	return s
}

type CreateConferenceResponseBodyAuthInfo struct {
	Key       *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Nonce     *string `json:"Nonce,omitempty" xml:"Nonce,omitempty"`
	Timestamp *int32  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s CreateConferenceResponseBodyAuthInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateConferenceResponseBodyAuthInfo) GoString() string {
	return s.String()
}

func (s *CreateConferenceResponseBodyAuthInfo) SetKey(v string) *CreateConferenceResponseBodyAuthInfo {
	s.Key = &v
	return s
}

func (s *CreateConferenceResponseBodyAuthInfo) SetNonce(v string) *CreateConferenceResponseBodyAuthInfo {
	s.Nonce = &v
	return s
}

func (s *CreateConferenceResponseBodyAuthInfo) SetTimestamp(v int32) *CreateConferenceResponseBodyAuthInfo {
	s.Timestamp = &v
	return s
}

type CreateConferenceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateConferenceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateConferenceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateConferenceResponse) GoString() string {
	return s.String()
}

func (s *CreateConferenceResponse) SetHeaders(v map[string]*string) *CreateConferenceResponse {
	s.Headers = v
	return s
}

func (s *CreateConferenceResponse) SetBody(v *CreateConferenceResponseBody) *CreateConferenceResponse {
	s.Body = v
	return s
}

type CreateEventSubscribeRequest struct {
	OwnerId     *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog     *string   `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId       *string   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId   *string   `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	CallbackUrl *string   `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	ClientToken *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Users       []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
	Events      []*string `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
}

func (s CreateEventSubscribeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEventSubscribeRequest) GoString() string {
	return s.String()
}

func (s *CreateEventSubscribeRequest) SetOwnerId(v int64) *CreateEventSubscribeRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateEventSubscribeRequest) SetShowLog(v string) *CreateEventSubscribeRequest {
	s.ShowLog = &v
	return s
}

func (s *CreateEventSubscribeRequest) SetAppId(v string) *CreateEventSubscribeRequest {
	s.AppId = &v
	return s
}

func (s *CreateEventSubscribeRequest) SetChannelId(v string) *CreateEventSubscribeRequest {
	s.ChannelId = &v
	return s
}

func (s *CreateEventSubscribeRequest) SetCallbackUrl(v string) *CreateEventSubscribeRequest {
	s.CallbackUrl = &v
	return s
}

func (s *CreateEventSubscribeRequest) SetClientToken(v string) *CreateEventSubscribeRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateEventSubscribeRequest) SetUsers(v []*string) *CreateEventSubscribeRequest {
	s.Users = v
	return s
}

func (s *CreateEventSubscribeRequest) SetEvents(v []*string) *CreateEventSubscribeRequest {
	s.Events = v
	return s
}

type CreateEventSubscribeResponseBody struct {
	SubscribeId *string `json:"SubscribeId,omitempty" xml:"SubscribeId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEventSubscribeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEventSubscribeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEventSubscribeResponseBody) SetSubscribeId(v string) *CreateEventSubscribeResponseBody {
	s.SubscribeId = &v
	return s
}

func (s *CreateEventSubscribeResponseBody) SetRequestId(v string) *CreateEventSubscribeResponseBody {
	s.RequestId = &v
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
	OwnerId       *int64                         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog       *string                        `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId         *string                        `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Name          *string                        `json:"Name,omitempty" xml:"Name,omitempty"`
	AudioMixCount *int32                         `json:"AudioMixCount,omitempty" xml:"AudioMixCount,omitempty"`
	Panes         []*CreateMPULayoutRequestPanes `json:"Panes,omitempty" xml:"Panes,omitempty" type:"Repeated"`
}

func (s CreateMPULayoutRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMPULayoutRequest) GoString() string {
	return s.String()
}

func (s *CreateMPULayoutRequest) SetOwnerId(v int64) *CreateMPULayoutRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateMPULayoutRequest) SetShowLog(v string) *CreateMPULayoutRequest {
	s.ShowLog = &v
	return s
}

func (s *CreateMPULayoutRequest) SetAppId(v string) *CreateMPULayoutRequest {
	s.AppId = &v
	return s
}

func (s *CreateMPULayoutRequest) SetName(v string) *CreateMPULayoutRequest {
	s.Name = &v
	return s
}

func (s *CreateMPULayoutRequest) SetAudioMixCount(v int32) *CreateMPULayoutRequest {
	s.AudioMixCount = &v
	return s
}

func (s *CreateMPULayoutRequest) SetPanes(v []*CreateMPULayoutRequestPanes) *CreateMPULayoutRequest {
	s.Panes = v
	return s
}

type CreateMPULayoutRequestPanes struct {
	MajorPane *int32   `json:"MajorPane,omitempty" xml:"MajorPane,omitempty"`
	Width     *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height    *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Y         *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	PaneId    *int32   `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
	ZOrder    *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
	X         *float32 `json:"X,omitempty" xml:"X,omitempty"`
}

func (s CreateMPULayoutRequestPanes) String() string {
	return tea.Prettify(s)
}

func (s CreateMPULayoutRequestPanes) GoString() string {
	return s.String()
}

func (s *CreateMPULayoutRequestPanes) SetMajorPane(v int32) *CreateMPULayoutRequestPanes {
	s.MajorPane = &v
	return s
}

func (s *CreateMPULayoutRequestPanes) SetWidth(v float32) *CreateMPULayoutRequestPanes {
	s.Width = &v
	return s
}

func (s *CreateMPULayoutRequestPanes) SetHeight(v float32) *CreateMPULayoutRequestPanes {
	s.Height = &v
	return s
}

func (s *CreateMPULayoutRequestPanes) SetY(v float32) *CreateMPULayoutRequestPanes {
	s.Y = &v
	return s
}

func (s *CreateMPULayoutRequestPanes) SetPaneId(v int32) *CreateMPULayoutRequestPanes {
	s.PaneId = &v
	return s
}

func (s *CreateMPULayoutRequestPanes) SetZOrder(v int32) *CreateMPULayoutRequestPanes {
	s.ZOrder = &v
	return s
}

func (s *CreateMPULayoutRequestPanes) SetX(v float32) *CreateMPULayoutRequestPanes {
	s.X = &v
	return s
}

type CreateMPULayoutResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	LayoutId  *int64  `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
}

func (s CreateMPULayoutResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMPULayoutResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMPULayoutResponseBody) SetRequestId(v string) *CreateMPULayoutResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMPULayoutResponseBody) SetLayoutId(v int64) *CreateMPULayoutResponseBody {
	s.LayoutId = &v
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

type CreateMPURuleRequest struct {
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog         *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelPrefix   *string `json:"ChannelPrefix,omitempty" xml:"ChannelPrefix,omitempty"`
	MediaEncode     *int32  `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	BackgroundColor *int32  `json:"BackgroundColor,omitempty" xml:"BackgroundColor,omitempty"`
	CropMode        *int32  `json:"CropMode,omitempty" xml:"CropMode,omitempty"`
	TaskProfile     *string `json:"TaskProfile,omitempty" xml:"TaskProfile,omitempty"`
	PlayDomain      *string `json:"PlayDomain,omitempty" xml:"PlayDomain,omitempty"`
	CallBack        *string `json:"CallBack,omitempty" xml:"CallBack,omitempty"`
	LayoutIds       []*int  `json:"LayoutIds,omitempty" xml:"LayoutIds,omitempty" type:"Repeated"`
}

func (s CreateMPURuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMPURuleRequest) GoString() string {
	return s.String()
}

func (s *CreateMPURuleRequest) SetOwnerId(v int64) *CreateMPURuleRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateMPURuleRequest) SetShowLog(v string) *CreateMPURuleRequest {
	s.ShowLog = &v
	return s
}

func (s *CreateMPURuleRequest) SetAppId(v string) *CreateMPURuleRequest {
	s.AppId = &v
	return s
}

func (s *CreateMPURuleRequest) SetChannelPrefix(v string) *CreateMPURuleRequest {
	s.ChannelPrefix = &v
	return s
}

func (s *CreateMPURuleRequest) SetMediaEncode(v int32) *CreateMPURuleRequest {
	s.MediaEncode = &v
	return s
}

func (s *CreateMPURuleRequest) SetBackgroundColor(v int32) *CreateMPURuleRequest {
	s.BackgroundColor = &v
	return s
}

func (s *CreateMPURuleRequest) SetCropMode(v int32) *CreateMPURuleRequest {
	s.CropMode = &v
	return s
}

func (s *CreateMPURuleRequest) SetTaskProfile(v string) *CreateMPURuleRequest {
	s.TaskProfile = &v
	return s
}

func (s *CreateMPURuleRequest) SetPlayDomain(v string) *CreateMPURuleRequest {
	s.PlayDomain = &v
	return s
}

func (s *CreateMPURuleRequest) SetCallBack(v string) *CreateMPURuleRequest {
	s.CallBack = &v
	return s
}

func (s *CreateMPURuleRequest) SetLayoutIds(v []*int) *CreateMPURuleRequest {
	s.LayoutIds = v
	return s
}

type CreateMPURuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RuleId    *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s CreateMPURuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMPURuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMPURuleResponseBody) SetRequestId(v string) *CreateMPURuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMPURuleResponseBody) SetRuleId(v int64) *CreateMPURuleResponseBody {
	s.RuleId = &v
	return s
}

type CreateMPURuleResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateMPURuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMPURuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMPURuleResponse) GoString() string {
	return s.String()
}

func (s *CreateMPURuleResponse) SetHeaders(v map[string]*string) *CreateMPURuleResponse {
	s.Headers = v
	return s
}

func (s *CreateMPURuleResponse) SetBody(v *CreateMPURuleResponseBody) *CreateMPURuleResponse {
	s.Body = v
	return s
}

type CreateServiceLinkedRoleForRtcRequest struct {
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
}

func (s CreateServiceLinkedRoleForRtcRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleForRtcRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleForRtcRequest) SetOwnerId(v int64) *CreateServiceLinkedRoleForRtcRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateServiceLinkedRoleForRtcRequest) SetShowLog(v string) *CreateServiceLinkedRoleForRtcRequest {
	s.ShowLog = &v
	return s
}

type CreateServiceLinkedRoleForRtcResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateServiceLinkedRoleForRtcResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleForRtcResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleForRtcResponseBody) SetRequestId(v string) *CreateServiceLinkedRoleForRtcResponseBody {
	s.RequestId = &v
	return s
}

type CreateServiceLinkedRoleForRtcResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateServiceLinkedRoleForRtcResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateServiceLinkedRoleForRtcResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleForRtcResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleForRtcResponse) SetHeaders(v map[string]*string) *CreateServiceLinkedRoleForRtcResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceLinkedRoleForRtcResponse) SetBody(v *CreateServiceLinkedRoleForRtcResponseBody) *CreateServiceLinkedRoleForRtcResponse {
	s.Body = v
	return s
}

type CreateSubscribeRequest struct {
	OwnerId     *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog     *string   `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId       *string   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId   *string   `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	CallbackUrl *string   `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	ClientToken *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Users       []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
	Events      []*string `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
}

func (s CreateSubscribeRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSubscribeRequest) GoString() string {
	return s.String()
}

func (s *CreateSubscribeRequest) SetOwnerId(v int64) *CreateSubscribeRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSubscribeRequest) SetShowLog(v string) *CreateSubscribeRequest {
	s.ShowLog = &v
	return s
}

func (s *CreateSubscribeRequest) SetAppId(v string) *CreateSubscribeRequest {
	s.AppId = &v
	return s
}

func (s *CreateSubscribeRequest) SetChannelId(v string) *CreateSubscribeRequest {
	s.ChannelId = &v
	return s
}

func (s *CreateSubscribeRequest) SetCallbackUrl(v string) *CreateSubscribeRequest {
	s.CallbackUrl = &v
	return s
}

func (s *CreateSubscribeRequest) SetClientToken(v string) *CreateSubscribeRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSubscribeRequest) SetUsers(v []*string) *CreateSubscribeRequest {
	s.Users = v
	return s
}

func (s *CreateSubscribeRequest) SetEvents(v []*string) *CreateSubscribeRequest {
	s.Events = v
	return s
}

type CreateSubscribeResponseBody struct {
	SubscribeId *string `json:"SubscribeId,omitempty" xml:"SubscribeId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSubscribeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSubscribeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSubscribeResponseBody) SetSubscribeId(v string) *CreateSubscribeResponseBody {
	s.SubscribeId = &v
	return s
}

func (s *CreateSubscribeResponseBody) SetRequestId(v string) *CreateSubscribeResponseBody {
	s.RequestId = &v
	return s
}

type CreateSubscribeResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSubscribeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSubscribeResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSubscribeResponse) GoString() string {
	return s.String()
}

func (s *CreateSubscribeResponse) SetHeaders(v map[string]*string) *CreateSubscribeResponse {
	s.Headers = v
	return s
}

func (s *CreateSubscribeResponse) SetBody(v *CreateSubscribeResponseBody) *CreateSubscribeResponse {
	s.Body = v
	return s
}

type DeleteAutoLiveStreamRuleRequest struct {
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RuleId  *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteAutoLiveStreamRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAutoLiveStreamRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteAutoLiveStreamRuleRequest) SetOwnerId(v int64) *DeleteAutoLiveStreamRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteAutoLiveStreamRuleRequest) SetShowLog(v string) *DeleteAutoLiveStreamRuleRequest {
	s.ShowLog = &v
	return s
}

func (s *DeleteAutoLiveStreamRuleRequest) SetAppId(v string) *DeleteAutoLiveStreamRuleRequest {
	s.AppId = &v
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

type DeleteChannelRequest struct {
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog   *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
}

func (s DeleteChannelRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteChannelRequest) GoString() string {
	return s.String()
}

func (s *DeleteChannelRequest) SetOwnerId(v int64) *DeleteChannelRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteChannelRequest) SetShowLog(v string) *DeleteChannelRequest {
	s.ShowLog = &v
	return s
}

func (s *DeleteChannelRequest) SetAppId(v string) *DeleteChannelRequest {
	s.AppId = &v
	return s
}

func (s *DeleteChannelRequest) SetChannelId(v string) *DeleteChannelRequest {
	s.ChannelId = &v
	return s
}

type DeleteChannelResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteChannelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteChannelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteChannelResponseBody) SetRequestId(v string) *DeleteChannelResponseBody {
	s.RequestId = &v
	return s
}

type DeleteChannelResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteChannelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteChannelResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteChannelResponse) GoString() string {
	return s.String()
}

func (s *DeleteChannelResponse) SetHeaders(v map[string]*string) *DeleteChannelResponse {
	s.Headers = v
	return s
}

func (s *DeleteChannelResponse) SetBody(v *DeleteChannelResponseBody) *DeleteChannelResponse {
	s.Body = v
	return s
}

type DeleteConferenceRequest struct {
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog      *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ConferenceId *string `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
}

func (s DeleteConferenceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteConferenceRequest) GoString() string {
	return s.String()
}

func (s *DeleteConferenceRequest) SetOwnerId(v int64) *DeleteConferenceRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteConferenceRequest) SetShowLog(v string) *DeleteConferenceRequest {
	s.ShowLog = &v
	return s
}

func (s *DeleteConferenceRequest) SetAppId(v string) *DeleteConferenceRequest {
	s.AppId = &v
	return s
}

func (s *DeleteConferenceRequest) SetConferenceId(v string) *DeleteConferenceRequest {
	s.ConferenceId = &v
	return s
}

type DeleteConferenceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteConferenceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteConferenceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConferenceResponseBody) SetRequestId(v string) *DeleteConferenceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteConferenceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteConferenceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteConferenceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteConferenceResponse) GoString() string {
	return s.String()
}

func (s *DeleteConferenceResponse) SetHeaders(v map[string]*string) *DeleteConferenceResponse {
	s.Headers = v
	return s
}

func (s *DeleteConferenceResponse) SetBody(v *DeleteConferenceResponseBody) *DeleteConferenceResponse {
	s.Body = v
	return s
}

type DeleteEventSubscribeRequest struct {
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog     *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	SubscribeId *string `json:"SubscribeId,omitempty" xml:"SubscribeId,omitempty"`
}

func (s DeleteEventSubscribeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventSubscribeRequest) GoString() string {
	return s.String()
}

func (s *DeleteEventSubscribeRequest) SetOwnerId(v int64) *DeleteEventSubscribeRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteEventSubscribeRequest) SetShowLog(v string) *DeleteEventSubscribeRequest {
	s.ShowLog = &v
	return s
}

func (s *DeleteEventSubscribeRequest) SetAppId(v string) *DeleteEventSubscribeRequest {
	s.AppId = &v
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
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog  *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	LayoutId *int64  `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
}

func (s DeleteMPULayoutRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMPULayoutRequest) GoString() string {
	return s.String()
}

func (s *DeleteMPULayoutRequest) SetOwnerId(v int64) *DeleteMPULayoutRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteMPULayoutRequest) SetShowLog(v string) *DeleteMPULayoutRequest {
	s.ShowLog = &v
	return s
}

func (s *DeleteMPULayoutRequest) SetAppId(v string) *DeleteMPULayoutRequest {
	s.AppId = &v
	return s
}

func (s *DeleteMPULayoutRequest) SetLayoutId(v int64) *DeleteMPULayoutRequest {
	s.LayoutId = &v
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

type DeleteMPURuleRequest struct {
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RuleId  *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteMPURuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMPURuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteMPURuleRequest) SetOwnerId(v int64) *DeleteMPURuleRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteMPURuleRequest) SetShowLog(v string) *DeleteMPURuleRequest {
	s.ShowLog = &v
	return s
}

func (s *DeleteMPURuleRequest) SetAppId(v string) *DeleteMPURuleRequest {
	s.AppId = &v
	return s
}

func (s *DeleteMPURuleRequest) SetRuleId(v int64) *DeleteMPURuleRequest {
	s.RuleId = &v
	return s
}

type DeleteMPURuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMPURuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMPURuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMPURuleResponseBody) SetRequestId(v string) *DeleteMPURuleResponseBody {
	s.RequestId = &v
	return s
}

type DeleteMPURuleResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteMPURuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMPURuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMPURuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteMPURuleResponse) SetHeaders(v map[string]*string) *DeleteMPURuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteMPURuleResponse) SetBody(v *DeleteMPURuleResponseBody) *DeleteMPURuleResponse {
	s.Body = v
	return s
}

type DeleteRecordTemplateRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog    *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteRecordTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRecordTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteRecordTemplateRequest) SetOwnerId(v int64) *DeleteRecordTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteRecordTemplateRequest) SetShowLog(v string) *DeleteRecordTemplateRequest {
	s.ShowLog = &v
	return s
}

func (s *DeleteRecordTemplateRequest) SetAppId(v string) *DeleteRecordTemplateRequest {
	s.AppId = &v
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

type DeleteSubscribeRequest struct {
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog     *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	SubscribeId *string `json:"SubscribeId,omitempty" xml:"SubscribeId,omitempty"`
}

func (s DeleteSubscribeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubscribeRequest) GoString() string {
	return s.String()
}

func (s *DeleteSubscribeRequest) SetOwnerId(v int64) *DeleteSubscribeRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSubscribeRequest) SetShowLog(v string) *DeleteSubscribeRequest {
	s.ShowLog = &v
	return s
}

func (s *DeleteSubscribeRequest) SetAppId(v string) *DeleteSubscribeRequest {
	s.AppId = &v
	return s
}

func (s *DeleteSubscribeRequest) SetSubscribeId(v string) *DeleteSubscribeRequest {
	s.SubscribeId = &v
	return s
}

type DeleteSubscribeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSubscribeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubscribeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSubscribeResponseBody) SetRequestId(v string) *DeleteSubscribeResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSubscribeResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSubscribeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSubscribeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubscribeResponse) GoString() string {
	return s.String()
}

func (s *DeleteSubscribeResponse) SetHeaders(v map[string]*string) *DeleteSubscribeResponse {
	s.Headers = v
	return s
}

func (s *DeleteSubscribeResponse) SetBody(v *DeleteSubscribeResponseBody) *DeleteSubscribeResponse {
	s.Body = v
	return s
}

type DescribeAppsRequest struct {
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog  *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Order    *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNum  *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppsRequest) SetOwnerId(v int64) *DescribeAppsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAppsRequest) SetShowLog(v string) *DescribeAppsRequest {
	s.ShowLog = &v
	return s
}

func (s *DescribeAppsRequest) SetAppId(v string) *DescribeAppsRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppsRequest) SetStatus(v string) *DescribeAppsRequest {
	s.Status = &v
	return s
}

func (s *DescribeAppsRequest) SetOrder(v string) *DescribeAppsRequest {
	s.Order = &v
	return s
}

func (s *DescribeAppsRequest) SetPageNum(v int32) *DescribeAppsRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeAppsRequest) SetPageSize(v int32) *DescribeAppsRequest {
	s.PageSize = &v
	return s
}

type DescribeAppsResponseBody struct {
	TotalNum  *int32                           `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	TotalPage *int32                           `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AppList   *DescribeAppsResponseBodyAppList `json:"AppList,omitempty" xml:"AppList,omitempty" type:"Struct"`
}

func (s DescribeAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBody) SetTotalNum(v int32) *DescribeAppsResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeAppsResponseBody) SetTotalPage(v int32) *DescribeAppsResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeAppsResponseBody) SetRequestId(v string) *DescribeAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppsResponseBody) SetAppList(v *DescribeAppsResponseBodyAppList) *DescribeAppsResponseBody {
	s.AppList = v
	return s
}

type DescribeAppsResponseBodyAppList struct {
	App []*DescribeAppsResponseBodyAppListApp `json:"App,omitempty" xml:"App,omitempty" type:"Repeated"`
}

func (s DescribeAppsResponseBodyAppList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponseBodyAppList) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyAppList) SetApp(v []*DescribeAppsResponseBodyAppListApp) *DescribeAppsResponseBodyAppList {
	s.App = v
	return s
}

type DescribeAppsResponseBodyAppListApp struct {
	Status       *int32                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	AppName      *string                                         `json:"AppName,omitempty" xml:"AppName,omitempty"`
	ServiceAreas *DescribeAppsResponseBodyAppListAppServiceAreas `json:"ServiceAreas,omitempty" xml:"ServiceAreas,omitempty" type:"Struct"`
	AppId        *string                                         `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CreateTime   *string                                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	BillType     *string                                         `json:"BillType,omitempty" xml:"BillType,omitempty"`
	AppType      *string                                         `json:"AppType,omitempty" xml:"AppType,omitempty"`
}

func (s DescribeAppsResponseBodyAppListApp) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponseBodyAppListApp) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyAppListApp) SetStatus(v int32) *DescribeAppsResponseBodyAppListApp {
	s.Status = &v
	return s
}

func (s *DescribeAppsResponseBodyAppListApp) SetAppName(v string) *DescribeAppsResponseBodyAppListApp {
	s.AppName = &v
	return s
}

func (s *DescribeAppsResponseBodyAppListApp) SetServiceAreas(v *DescribeAppsResponseBodyAppListAppServiceAreas) *DescribeAppsResponseBodyAppListApp {
	s.ServiceAreas = v
	return s
}

func (s *DescribeAppsResponseBodyAppListApp) SetAppId(v string) *DescribeAppsResponseBodyAppListApp {
	s.AppId = &v
	return s
}

func (s *DescribeAppsResponseBodyAppListApp) SetCreateTime(v string) *DescribeAppsResponseBodyAppListApp {
	s.CreateTime = &v
	return s
}

func (s *DescribeAppsResponseBodyAppListApp) SetBillType(v string) *DescribeAppsResponseBodyAppListApp {
	s.BillType = &v
	return s
}

func (s *DescribeAppsResponseBodyAppListApp) SetAppType(v string) *DescribeAppsResponseBodyAppListApp {
	s.AppType = &v
	return s
}

type DescribeAppsResponseBodyAppListAppServiceAreas struct {
	ServiceArea []*string `json:"ServiceArea,omitempty" xml:"ServiceArea,omitempty" type:"Repeated"`
}

func (s DescribeAppsResponseBodyAppListAppServiceAreas) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponseBodyAppListAppServiceAreas) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyAppListAppServiceAreas) SetServiceArea(v []*string) *DescribeAppsResponseBodyAppListAppServiceAreas {
	s.ServiceArea = v
	return s
}

type DescribeAppsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeAppsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponse) SetHeaders(v map[string]*string) *DescribeAppsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppsResponse) SetBody(v *DescribeAppsResponseBody) *DescribeAppsResponse {
	s.Body = v
	return s
}

type DescribeAutoLiveStreamRuleRequest struct {
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DescribeAutoLiveStreamRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAutoLiveStreamRuleRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutoLiveStreamRuleRequest) SetOwnerId(v int64) *DescribeAutoLiveStreamRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAutoLiveStreamRuleRequest) SetShowLog(v string) *DescribeAutoLiveStreamRuleRequest {
	s.ShowLog = &v
	return s
}

func (s *DescribeAutoLiveStreamRuleRequest) SetAppId(v string) *DescribeAutoLiveStreamRuleRequest {
	s.AppId = &v
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
	CallBack   *string `json:"CallBack,omitempty" xml:"CallBack,omitempty"`
	PlayDomain *string `json:"PlayDomain,omitempty" xml:"PlayDomain,omitempty"`
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	RuleName   *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	RuleId     *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
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

func (s *DescribeAutoLiveStreamRuleResponseBodyRules) SetPlayDomain(v string) *DescribeAutoLiveStreamRuleResponseBodyRules {
	s.PlayDomain = &v
	return s
}

func (s *DescribeAutoLiveStreamRuleResponseBodyRules) SetCreateTime(v string) *DescribeAutoLiveStreamRuleResponseBodyRules {
	s.CreateTime = &v
	return s
}

func (s *DescribeAutoLiveStreamRuleResponseBodyRules) SetRuleName(v string) *DescribeAutoLiveStreamRuleResponseBodyRules {
	s.RuleName = &v
	return s
}

func (s *DescribeAutoLiveStreamRuleResponseBodyRules) SetRuleId(v int64) *DescribeAutoLiveStreamRuleResponseBodyRules {
	s.RuleId = &v
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
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog   *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	Order     *string `json:"Order,omitempty" xml:"Order,omitempty"`
	PageNum   *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeChannelParticipantsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelParticipantsRequest) GoString() string {
	return s.String()
}

func (s *DescribeChannelParticipantsRequest) SetOwnerId(v int64) *DescribeChannelParticipantsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeChannelParticipantsRequest) SetShowLog(v string) *DescribeChannelParticipantsRequest {
	s.ShowLog = &v
	return s
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

func (s *DescribeChannelParticipantsRequest) SetPageNum(v int32) *DescribeChannelParticipantsRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeChannelParticipantsRequest) SetPageSize(v int32) *DescribeChannelParticipantsRequest {
	s.PageSize = &v
	return s
}

type DescribeChannelParticipantsResponseBody struct {
	TotalNum  *int32                                           `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	TotalPage *int32                                           `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Timestamp *int32                                           `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	UserList  *DescribeChannelParticipantsResponseBodyUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Struct"`
}

func (s DescribeChannelParticipantsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelParticipantsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChannelParticipantsResponseBody) SetTotalNum(v int32) *DescribeChannelParticipantsResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeChannelParticipantsResponseBody) SetTotalPage(v int32) *DescribeChannelParticipantsResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeChannelParticipantsResponseBody) SetRequestId(v string) *DescribeChannelParticipantsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeChannelParticipantsResponseBody) SetTimestamp(v int32) *DescribeChannelParticipantsResponseBody {
	s.Timestamp = &v
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
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog   *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
}

func (s DescribeChannelUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelUsersRequest) GoString() string {
	return s.String()
}

func (s *DescribeChannelUsersRequest) SetOwnerId(v int64) *DescribeChannelUsersRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeChannelUsersRequest) SetShowLog(v string) *DescribeChannelUsersRequest {
	s.ShowLog = &v
	return s
}

func (s *DescribeChannelUsersRequest) SetAppId(v string) *DescribeChannelUsersRequest {
	s.AppId = &v
	return s
}

func (s *DescribeChannelUsersRequest) SetChannelId(v string) *DescribeChannelUsersRequest {
	s.ChannelId = &v
	return s
}

type DescribeChannelUsersResponseBody struct {
	RequestId           *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	InteractiveUserList []*string `json:"InteractiveUserList,omitempty" xml:"InteractiveUserList,omitempty" type:"Repeated"`
	LiveUserNum         *int32    `json:"LiveUserNum,omitempty" xml:"LiveUserNum,omitempty"`
	ChannelProfile      *int32    `json:"ChannelProfile,omitempty" xml:"ChannelProfile,omitempty"`
	InteractiveUserNum  *int32    `json:"InteractiveUserNum,omitempty" xml:"InteractiveUserNum,omitempty"`
	IsChannelExist      *bool     `json:"IsChannelExist,omitempty" xml:"IsChannelExist,omitempty"`
	Timestamp           *int32    `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	UserList            []*string `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
	LiveUserList        []*string `json:"LiveUserList,omitempty" xml:"LiveUserList,omitempty" type:"Repeated"`
	CommTotalNum        *int32    `json:"CommTotalNum,omitempty" xml:"CommTotalNum,omitempty"`
}

func (s DescribeChannelUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChannelUsersResponseBody) SetRequestId(v string) *DescribeChannelUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeChannelUsersResponseBody) SetInteractiveUserList(v []*string) *DescribeChannelUsersResponseBody {
	s.InteractiveUserList = v
	return s
}

func (s *DescribeChannelUsersResponseBody) SetLiveUserNum(v int32) *DescribeChannelUsersResponseBody {
	s.LiveUserNum = &v
	return s
}

func (s *DescribeChannelUsersResponseBody) SetChannelProfile(v int32) *DescribeChannelUsersResponseBody {
	s.ChannelProfile = &v
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

func (s *DescribeChannelUsersResponseBody) SetTimestamp(v int32) *DescribeChannelUsersResponseBody {
	s.Timestamp = &v
	return s
}

func (s *DescribeChannelUsersResponseBody) SetUserList(v []*string) *DescribeChannelUsersResponseBody {
	s.UserList = v
	return s
}

func (s *DescribeChannelUsersResponseBody) SetLiveUserList(v []*string) *DescribeChannelUsersResponseBody {
	s.LiveUserList = v
	return s
}

func (s *DescribeChannelUsersResponseBody) SetCommTotalNum(v int32) *DescribeChannelUsersResponseBody {
	s.CommTotalNum = &v
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

type DescribeConferenceAuthInfoRequest struct {
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog      *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId        *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ConferenceId *string `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
}

func (s DescribeConferenceAuthInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeConferenceAuthInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeConferenceAuthInfoRequest) SetOwnerId(v int64) *DescribeConferenceAuthInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeConferenceAuthInfoRequest) SetShowLog(v string) *DescribeConferenceAuthInfoRequest {
	s.ShowLog = &v
	return s
}

func (s *DescribeConferenceAuthInfoRequest) SetAppId(v string) *DescribeConferenceAuthInfoRequest {
	s.AppId = &v
	return s
}

func (s *DescribeConferenceAuthInfoRequest) SetConferenceId(v string) *DescribeConferenceAuthInfoRequest {
	s.ConferenceId = &v
	return s
}

type DescribeConferenceAuthInfoResponseBody struct {
	AuthInfo     *DescribeConferenceAuthInfoResponseBodyAuthInfo `json:"AuthInfo,omitempty" xml:"AuthInfo,omitempty" type:"Struct"`
	RequestId    *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ConferenceId *string                                         `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
}

func (s DescribeConferenceAuthInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeConferenceAuthInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConferenceAuthInfoResponseBody) SetAuthInfo(v *DescribeConferenceAuthInfoResponseBodyAuthInfo) *DescribeConferenceAuthInfoResponseBody {
	s.AuthInfo = v
	return s
}

func (s *DescribeConferenceAuthInfoResponseBody) SetRequestId(v string) *DescribeConferenceAuthInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConferenceAuthInfoResponseBody) SetConferenceId(v string) *DescribeConferenceAuthInfoResponseBody {
	s.ConferenceId = &v
	return s
}

type DescribeConferenceAuthInfoResponseBodyAuthInfo struct {
	Key       *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Nonce     *string `json:"Nonce,omitempty" xml:"Nonce,omitempty"`
	Timestamp *int32  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeConferenceAuthInfoResponseBodyAuthInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeConferenceAuthInfoResponseBodyAuthInfo) GoString() string {
	return s.String()
}

func (s *DescribeConferenceAuthInfoResponseBodyAuthInfo) SetKey(v string) *DescribeConferenceAuthInfoResponseBodyAuthInfo {
	s.Key = &v
	return s
}

func (s *DescribeConferenceAuthInfoResponseBodyAuthInfo) SetNonce(v string) *DescribeConferenceAuthInfoResponseBodyAuthInfo {
	s.Nonce = &v
	return s
}

func (s *DescribeConferenceAuthInfoResponseBodyAuthInfo) SetTimestamp(v int32) *DescribeConferenceAuthInfoResponseBodyAuthInfo {
	s.Timestamp = &v
	return s
}

type DescribeConferenceAuthInfoResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeConferenceAuthInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeConferenceAuthInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeConferenceAuthInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeConferenceAuthInfoResponse) SetHeaders(v map[string]*string) *DescribeConferenceAuthInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeConferenceAuthInfoResponse) SetBody(v *DescribeConferenceAuthInfoResponseBody) *DescribeConferenceAuthInfoResponse {
	s.Body = v
	return s
}

type DescribeMPULayoutInfoRequest struct {
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog  *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	LayoutId *int64  `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
}

func (s DescribeMPULayoutInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPULayoutInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeMPULayoutInfoRequest) SetOwnerId(v int64) *DescribeMPULayoutInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMPULayoutInfoRequest) SetShowLog(v string) *DescribeMPULayoutInfoRequest {
	s.ShowLog = &v
	return s
}

func (s *DescribeMPULayoutInfoRequest) SetAppId(v string) *DescribeMPULayoutInfoRequest {
	s.AppId = &v
	return s
}

func (s *DescribeMPULayoutInfoRequest) SetLayoutId(v int64) *DescribeMPULayoutInfoRequest {
	s.LayoutId = &v
	return s
}

type DescribeMPULayoutInfoResponseBody struct {
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Layout    *DescribeMPULayoutInfoResponseBodyLayout `json:"Layout,omitempty" xml:"Layout,omitempty" type:"Struct"`
}

func (s DescribeMPULayoutInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPULayoutInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMPULayoutInfoResponseBody) SetRequestId(v string) *DescribeMPULayoutInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMPULayoutInfoResponseBody) SetLayout(v *DescribeMPULayoutInfoResponseBodyLayout) *DescribeMPULayoutInfoResponseBody {
	s.Layout = v
	return s
}

type DescribeMPULayoutInfoResponseBodyLayout struct {
	LayoutId      *int64                                        `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	Panes         *DescribeMPULayoutInfoResponseBodyLayoutPanes `json:"Panes,omitempty" xml:"Panes,omitempty" type:"Struct"`
	Name          *string                                       `json:"Name,omitempty" xml:"Name,omitempty"`
	AudioMixCount *int32                                        `json:"AudioMixCount,omitempty" xml:"AudioMixCount,omitempty"`
}

func (s DescribeMPULayoutInfoResponseBodyLayout) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPULayoutInfoResponseBodyLayout) GoString() string {
	return s.String()
}

func (s *DescribeMPULayoutInfoResponseBodyLayout) SetLayoutId(v int64) *DescribeMPULayoutInfoResponseBodyLayout {
	s.LayoutId = &v
	return s
}

func (s *DescribeMPULayoutInfoResponseBodyLayout) SetPanes(v *DescribeMPULayoutInfoResponseBodyLayoutPanes) *DescribeMPULayoutInfoResponseBodyLayout {
	s.Panes = v
	return s
}

func (s *DescribeMPULayoutInfoResponseBodyLayout) SetName(v string) *DescribeMPULayoutInfoResponseBodyLayout {
	s.Name = &v
	return s
}

func (s *DescribeMPULayoutInfoResponseBodyLayout) SetAudioMixCount(v int32) *DescribeMPULayoutInfoResponseBodyLayout {
	s.AudioMixCount = &v
	return s
}

type DescribeMPULayoutInfoResponseBodyLayoutPanes struct {
	Panes []*DescribeMPULayoutInfoResponseBodyLayoutPanesPanes `json:"Panes,omitempty" xml:"Panes,omitempty" type:"Repeated"`
}

func (s DescribeMPULayoutInfoResponseBodyLayoutPanes) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPULayoutInfoResponseBodyLayoutPanes) GoString() string {
	return s.String()
}

func (s *DescribeMPULayoutInfoResponseBodyLayoutPanes) SetPanes(v []*DescribeMPULayoutInfoResponseBodyLayoutPanesPanes) *DescribeMPULayoutInfoResponseBodyLayoutPanes {
	s.Panes = v
	return s
}

type DescribeMPULayoutInfoResponseBodyLayoutPanesPanes struct {
	MajorPane *int32   `json:"MajorPane,omitempty" xml:"MajorPane,omitempty"`
	Width     *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height    *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Y         *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	PaneId    *int32   `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
	ZOrder    *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
	X         *float32 `json:"X,omitempty" xml:"X,omitempty"`
}

func (s DescribeMPULayoutInfoResponseBodyLayoutPanesPanes) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPULayoutInfoResponseBodyLayoutPanesPanes) GoString() string {
	return s.String()
}

func (s *DescribeMPULayoutInfoResponseBodyLayoutPanesPanes) SetMajorPane(v int32) *DescribeMPULayoutInfoResponseBodyLayoutPanesPanes {
	s.MajorPane = &v
	return s
}

func (s *DescribeMPULayoutInfoResponseBodyLayoutPanesPanes) SetWidth(v float32) *DescribeMPULayoutInfoResponseBodyLayoutPanesPanes {
	s.Width = &v
	return s
}

func (s *DescribeMPULayoutInfoResponseBodyLayoutPanesPanes) SetHeight(v float32) *DescribeMPULayoutInfoResponseBodyLayoutPanesPanes {
	s.Height = &v
	return s
}

func (s *DescribeMPULayoutInfoResponseBodyLayoutPanesPanes) SetY(v float32) *DescribeMPULayoutInfoResponseBodyLayoutPanesPanes {
	s.Y = &v
	return s
}

func (s *DescribeMPULayoutInfoResponseBodyLayoutPanesPanes) SetPaneId(v int32) *DescribeMPULayoutInfoResponseBodyLayoutPanesPanes {
	s.PaneId = &v
	return s
}

func (s *DescribeMPULayoutInfoResponseBodyLayoutPanesPanes) SetZOrder(v int32) *DescribeMPULayoutInfoResponseBodyLayoutPanesPanes {
	s.ZOrder = &v
	return s
}

func (s *DescribeMPULayoutInfoResponseBodyLayoutPanesPanes) SetX(v float32) *DescribeMPULayoutInfoResponseBodyLayoutPanesPanes {
	s.X = &v
	return s
}

type DescribeMPULayoutInfoResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMPULayoutInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMPULayoutInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPULayoutInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeMPULayoutInfoResponse) SetHeaders(v map[string]*string) *DescribeMPULayoutInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeMPULayoutInfoResponse) SetBody(v *DescribeMPULayoutInfoResponseBody) *DescribeMPULayoutInfoResponse {
	s.Body = v
	return s
}

type DescribeMPULayoutInfoListRequest struct {
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog  *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	LayoutId *int64  `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PageNum  *int64  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeMPULayoutInfoListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPULayoutInfoListRequest) GoString() string {
	return s.String()
}

func (s *DescribeMPULayoutInfoListRequest) SetOwnerId(v int64) *DescribeMPULayoutInfoListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMPULayoutInfoListRequest) SetShowLog(v string) *DescribeMPULayoutInfoListRequest {
	s.ShowLog = &v
	return s
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

func (s *DescribeMPULayoutInfoListRequest) SetPageNum(v int64) *DescribeMPULayoutInfoListRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeMPULayoutInfoListRequest) SetPageSize(v int64) *DescribeMPULayoutInfoListRequest {
	s.PageSize = &v
	return s
}

type DescribeMPULayoutInfoListResponseBody struct {
	TotalNum  *int64                                        `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	TotalPage *int64                                        `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Layouts   *DescribeMPULayoutInfoListResponseBodyLayouts `json:"Layouts,omitempty" xml:"Layouts,omitempty" type:"Struct"`
}

func (s DescribeMPULayoutInfoListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPULayoutInfoListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMPULayoutInfoListResponseBody) SetTotalNum(v int64) *DescribeMPULayoutInfoListResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBody) SetTotalPage(v int64) *DescribeMPULayoutInfoListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBody) SetRequestId(v string) *DescribeMPULayoutInfoListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBody) SetLayouts(v *DescribeMPULayoutInfoListResponseBodyLayouts) *DescribeMPULayoutInfoListResponseBody {
	s.Layouts = v
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
	LayoutId      *int64                                                   `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	Panes         *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanes `json:"Panes,omitempty" xml:"Panes,omitempty" type:"Struct"`
	Name          *string                                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	AudioMixCount *int32                                                   `json:"AudioMixCount,omitempty" xml:"AudioMixCount,omitempty"`
}

func (s DescribeMPULayoutInfoListResponseBodyLayoutsLayout) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPULayoutInfoListResponseBodyLayoutsLayout) GoString() string {
	return s.String()
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayout) SetLayoutId(v int64) *DescribeMPULayoutInfoListResponseBodyLayoutsLayout {
	s.LayoutId = &v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayout) SetPanes(v *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanes) *DescribeMPULayoutInfoListResponseBodyLayoutsLayout {
	s.Panes = v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayout) SetName(v string) *DescribeMPULayoutInfoListResponseBodyLayoutsLayout {
	s.Name = &v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayout) SetAudioMixCount(v int32) *DescribeMPULayoutInfoListResponseBodyLayoutsLayout {
	s.AudioMixCount = &v
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
	MajorPane *int32   `json:"MajorPane,omitempty" xml:"MajorPane,omitempty"`
	Width     *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height    *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Y         *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	PaneId    *int32   `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
	ZOrder    *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
	X         *float32 `json:"X,omitempty" xml:"X,omitempty"`
}

func (s DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes) GoString() string {
	return s.String()
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes) SetMajorPane(v int32) *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes {
	s.MajorPane = &v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes) SetWidth(v float32) *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes {
	s.Width = &v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes) SetHeight(v float32) *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes {
	s.Height = &v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes) SetY(v float32) *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes {
	s.Y = &v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes) SetPaneId(v int32) *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes {
	s.PaneId = &v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes) SetZOrder(v int32) *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes {
	s.ZOrder = &v
	return s
}

func (s *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes) SetX(v float32) *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanesPanes {
	s.X = &v
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

type DescribeMPULayoutListRequest struct {
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DescribeMPULayoutListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPULayoutListRequest) GoString() string {
	return s.String()
}

func (s *DescribeMPULayoutListRequest) SetOwnerId(v int64) *DescribeMPULayoutListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMPULayoutListRequest) SetShowLog(v string) *DescribeMPULayoutListRequest {
	s.ShowLog = &v
	return s
}

func (s *DescribeMPULayoutListRequest) SetAppId(v string) *DescribeMPULayoutListRequest {
	s.AppId = &v
	return s
}

type DescribeMPULayoutListResponseBody struct {
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	LayoutIds *DescribeMPULayoutListResponseBodyLayoutIds `json:"LayoutIds,omitempty" xml:"LayoutIds,omitempty" type:"Struct"`
}

func (s DescribeMPULayoutListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPULayoutListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMPULayoutListResponseBody) SetRequestId(v string) *DescribeMPULayoutListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMPULayoutListResponseBody) SetLayoutIds(v *DescribeMPULayoutListResponseBodyLayoutIds) *DescribeMPULayoutListResponseBody {
	s.LayoutIds = v
	return s
}

type DescribeMPULayoutListResponseBodyLayoutIds struct {
	LayoutId []*string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty" type:"Repeated"`
}

func (s DescribeMPULayoutListResponseBodyLayoutIds) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPULayoutListResponseBodyLayoutIds) GoString() string {
	return s.String()
}

func (s *DescribeMPULayoutListResponseBodyLayoutIds) SetLayoutId(v []*string) *DescribeMPULayoutListResponseBodyLayoutIds {
	s.LayoutId = v
	return s
}

type DescribeMPULayoutListResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMPULayoutListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMPULayoutListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPULayoutListResponse) GoString() string {
	return s.String()
}

func (s *DescribeMPULayoutListResponse) SetHeaders(v map[string]*string) *DescribeMPULayoutListResponse {
	s.Headers = v
	return s
}

func (s *DescribeMPULayoutListResponse) SetBody(v *DescribeMPULayoutListResponseBody) *DescribeMPULayoutListResponse {
	s.Body = v
	return s
}

type DescribeMPURuleRequest struct {
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DescribeMPURuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPURuleRequest) GoString() string {
	return s.String()
}

func (s *DescribeMPURuleRequest) SetOwnerId(v int64) *DescribeMPURuleRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMPURuleRequest) SetShowLog(v string) *DescribeMPURuleRequest {
	s.ShowLog = &v
	return s
}

func (s *DescribeMPURuleRequest) SetAppId(v string) *DescribeMPURuleRequest {
	s.AppId = &v
	return s
}

type DescribeMPURuleResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rules     []*DescribeMPURuleResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeMPURuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPURuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMPURuleResponseBody) SetRequestId(v string) *DescribeMPURuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMPURuleResponseBody) SetRules(v []*DescribeMPURuleResponseBodyRules) *DescribeMPURuleResponseBody {
	s.Rules = v
	return s
}

type DescribeMPURuleResponseBodyRules struct {
	MediaEncode    *int32    `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	CropMode       *int32    `json:"CropMode,omitempty" xml:"CropMode,omitempty"`
	CallBack       *string   `json:"CallBack,omitempty" xml:"CallBack,omitempty"`
	PlayDomain     *string   `json:"PlayDomain,omitempty" xml:"PlayDomain,omitempty"`
	ChannelPrefix  *string   `json:"ChannelPrefix,omitempty" xml:"ChannelPrefix,omitempty"`
	BackgroudColor *int32    `json:"BackgroudColor,omitempty" xml:"BackgroudColor,omitempty"`
	IsEnable       *int32    `json:"IsEnable,omitempty" xml:"IsEnable,omitempty"`
	LayoutIds      []*string `json:"LayoutIds,omitempty" xml:"LayoutIds,omitempty" type:"Repeated"`
	TaskProfile    *string   `json:"TaskProfile,omitempty" xml:"TaskProfile,omitempty"`
	RuleId         *int64    `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeMPURuleResponseBodyRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPURuleResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribeMPURuleResponseBodyRules) SetMediaEncode(v int32) *DescribeMPURuleResponseBodyRules {
	s.MediaEncode = &v
	return s
}

func (s *DescribeMPURuleResponseBodyRules) SetCropMode(v int32) *DescribeMPURuleResponseBodyRules {
	s.CropMode = &v
	return s
}

func (s *DescribeMPURuleResponseBodyRules) SetCallBack(v string) *DescribeMPURuleResponseBodyRules {
	s.CallBack = &v
	return s
}

func (s *DescribeMPURuleResponseBodyRules) SetPlayDomain(v string) *DescribeMPURuleResponseBodyRules {
	s.PlayDomain = &v
	return s
}

func (s *DescribeMPURuleResponseBodyRules) SetChannelPrefix(v string) *DescribeMPURuleResponseBodyRules {
	s.ChannelPrefix = &v
	return s
}

func (s *DescribeMPURuleResponseBodyRules) SetBackgroudColor(v int32) *DescribeMPURuleResponseBodyRules {
	s.BackgroudColor = &v
	return s
}

func (s *DescribeMPURuleResponseBodyRules) SetIsEnable(v int32) *DescribeMPURuleResponseBodyRules {
	s.IsEnable = &v
	return s
}

func (s *DescribeMPURuleResponseBodyRules) SetLayoutIds(v []*string) *DescribeMPURuleResponseBodyRules {
	s.LayoutIds = v
	return s
}

func (s *DescribeMPURuleResponseBodyRules) SetTaskProfile(v string) *DescribeMPURuleResponseBodyRules {
	s.TaskProfile = &v
	return s
}

func (s *DescribeMPURuleResponseBodyRules) SetRuleId(v int64) *DescribeMPURuleResponseBodyRules {
	s.RuleId = &v
	return s
}

type DescribeMPURuleResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMPURuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMPURuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMPURuleResponse) GoString() string {
	return s.String()
}

func (s *DescribeMPURuleResponse) SetHeaders(v map[string]*string) *DescribeMPURuleResponse {
	s.Headers = v
	return s
}

func (s *DescribeMPURuleResponse) SetBody(v *DescribeMPURuleResponseBody) *DescribeMPURuleResponse {
	s.Body = v
	return s
}

type DescribeRecordFilesRequest struct {
	OwnerId   *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog   *string   `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId     *string   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId *string   `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	PageSize  *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNum   *int32    `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	StartTime *string   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *string   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	TaskIds   []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
}

func (s DescribeRecordFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordFilesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecordFilesRequest) SetOwnerId(v int64) *DescribeRecordFilesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRecordFilesRequest) SetShowLog(v string) *DescribeRecordFilesRequest {
	s.ShowLog = &v
	return s
}

func (s *DescribeRecordFilesRequest) SetAppId(v string) *DescribeRecordFilesRequest {
	s.AppId = &v
	return s
}

func (s *DescribeRecordFilesRequest) SetChannelId(v string) *DescribeRecordFilesRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeRecordFilesRequest) SetPageSize(v int32) *DescribeRecordFilesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRecordFilesRequest) SetPageNum(v int32) *DescribeRecordFilesRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeRecordFilesRequest) SetStartTime(v string) *DescribeRecordFilesRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRecordFilesRequest) SetEndTime(v string) *DescribeRecordFilesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRecordFilesRequest) SetTaskIds(v []*string) *DescribeRecordFilesRequest {
	s.TaskIds = v
	return s
}

type DescribeRecordFilesResponseBody struct {
	TotalNum    *int64                                        `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	TotalPage   *int64                                        `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RecordFiles []*DescribeRecordFilesResponseBodyRecordFiles `json:"RecordFiles,omitempty" xml:"RecordFiles,omitempty" type:"Repeated"`
}

func (s DescribeRecordFilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordFilesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecordFilesResponseBody) SetTotalNum(v int64) *DescribeRecordFilesResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeRecordFilesResponseBody) SetTotalPage(v int64) *DescribeRecordFilesResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeRecordFilesResponseBody) SetRequestId(v string) *DescribeRecordFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecordFilesResponseBody) SetRecordFiles(v []*DescribeRecordFilesResponseBodyRecordFiles) *DescribeRecordFilesResponseBody {
	s.RecordFiles = v
	return s
}

type DescribeRecordFilesResponseBodyRecordFiles struct {
	StartTime  *string  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	CreateTime *string  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	AppId      *string  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId  *string  `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	Url        *string  `json:"Url,omitempty" xml:"Url,omitempty"`
	Duration   *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	TaskId     *string  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	StopTime   *string  `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
}

func (s DescribeRecordFilesResponseBodyRecordFiles) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordFilesResponseBodyRecordFiles) GoString() string {
	return s.String()
}

func (s *DescribeRecordFilesResponseBodyRecordFiles) SetStartTime(v string) *DescribeRecordFilesResponseBodyRecordFiles {
	s.StartTime = &v
	return s
}

func (s *DescribeRecordFilesResponseBodyRecordFiles) SetCreateTime(v string) *DescribeRecordFilesResponseBodyRecordFiles {
	s.CreateTime = &v
	return s
}

func (s *DescribeRecordFilesResponseBodyRecordFiles) SetAppId(v string) *DescribeRecordFilesResponseBodyRecordFiles {
	s.AppId = &v
	return s
}

func (s *DescribeRecordFilesResponseBodyRecordFiles) SetChannelId(v string) *DescribeRecordFilesResponseBodyRecordFiles {
	s.ChannelId = &v
	return s
}

func (s *DescribeRecordFilesResponseBodyRecordFiles) SetUrl(v string) *DescribeRecordFilesResponseBodyRecordFiles {
	s.Url = &v
	return s
}

func (s *DescribeRecordFilesResponseBodyRecordFiles) SetDuration(v float32) *DescribeRecordFilesResponseBodyRecordFiles {
	s.Duration = &v
	return s
}

func (s *DescribeRecordFilesResponseBodyRecordFiles) SetTaskId(v string) *DescribeRecordFilesResponseBodyRecordFiles {
	s.TaskId = &v
	return s
}

func (s *DescribeRecordFilesResponseBodyRecordFiles) SetStopTime(v string) *DescribeRecordFilesResponseBodyRecordFiles {
	s.StopTime = &v
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
	OwnerId   *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog   *string   `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId     *string   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId *string   `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	Status    *string   `json:"Status,omitempty" xml:"Status,omitempty"`
	StartTime *string   `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *string   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageSize  *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNum   *int32    `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	TaskIds   []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
}

func (s DescribeRecordTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecordTasksRequest) SetOwnerId(v int64) *DescribeRecordTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRecordTasksRequest) SetShowLog(v string) *DescribeRecordTasksRequest {
	s.ShowLog = &v
	return s
}

func (s *DescribeRecordTasksRequest) SetAppId(v string) *DescribeRecordTasksRequest {
	s.AppId = &v
	return s
}

func (s *DescribeRecordTasksRequest) SetChannelId(v string) *DescribeRecordTasksRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeRecordTasksRequest) SetStatus(v string) *DescribeRecordTasksRequest {
	s.Status = &v
	return s
}

func (s *DescribeRecordTasksRequest) SetStartTime(v string) *DescribeRecordTasksRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRecordTasksRequest) SetEndTime(v string) *DescribeRecordTasksRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRecordTasksRequest) SetPageSize(v int32) *DescribeRecordTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRecordTasksRequest) SetPageNum(v int32) *DescribeRecordTasksRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeRecordTasksRequest) SetTaskIds(v []*string) *DescribeRecordTasksRequest {
	s.TaskIds = v
	return s
}

type DescribeRecordTasksResponseBody struct {
	RecordTasks []*DescribeRecordTasksResponseBodyRecordTasks `json:"RecordTasks,omitempty" xml:"RecordTasks,omitempty" type:"Repeated"`
	TotalNum    *int64                                        `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	TotalPage   *int64                                        `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *DescribeRecordTasksResponseBody) SetTotalNum(v int64) *DescribeRecordTasksResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeRecordTasksResponseBody) SetTotalPage(v int64) *DescribeRecordTasksResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeRecordTasksResponseBody) SetRequestId(v string) *DescribeRecordTasksResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRecordTasksResponseBodyRecordTasks struct {
	Status       *int32                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	SubSpecUsers []*string                                              `json:"SubSpecUsers,omitempty" xml:"SubSpecUsers,omitempty" type:"Repeated"`
	UserPanes    []*DescribeRecordTasksResponseBodyRecordTasksUserPanes `json:"UserPanes,omitempty" xml:"UserPanes,omitempty" type:"Repeated"`
	CreateTime   *string                                                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	AppId        *string                                                `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId    *string                                                `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	TaskId       *string                                                `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TemplateId   *string                                                `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeRecordTasksResponseBodyRecordTasks) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordTasksResponseBodyRecordTasks) GoString() string {
	return s.String()
}

func (s *DescribeRecordTasksResponseBodyRecordTasks) SetStatus(v int32) *DescribeRecordTasksResponseBodyRecordTasks {
	s.Status = &v
	return s
}

func (s *DescribeRecordTasksResponseBodyRecordTasks) SetSubSpecUsers(v []*string) *DescribeRecordTasksResponseBodyRecordTasks {
	s.SubSpecUsers = v
	return s
}

func (s *DescribeRecordTasksResponseBodyRecordTasks) SetUserPanes(v []*DescribeRecordTasksResponseBodyRecordTasksUserPanes) *DescribeRecordTasksResponseBodyRecordTasks {
	s.UserPanes = v
	return s
}

func (s *DescribeRecordTasksResponseBodyRecordTasks) SetCreateTime(v string) *DescribeRecordTasksResponseBodyRecordTasks {
	s.CreateTime = &v
	return s
}

func (s *DescribeRecordTasksResponseBodyRecordTasks) SetAppId(v string) *DescribeRecordTasksResponseBodyRecordTasks {
	s.AppId = &v
	return s
}

func (s *DescribeRecordTasksResponseBodyRecordTasks) SetChannelId(v string) *DescribeRecordTasksResponseBodyRecordTasks {
	s.ChannelId = &v
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

type DescribeRecordTasksResponseBodyRecordTasksUserPanes struct {
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	PaneId *int32  `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s DescribeRecordTasksResponseBodyRecordTasksUserPanes) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordTasksResponseBodyRecordTasksUserPanes) GoString() string {
	return s.String()
}

func (s *DescribeRecordTasksResponseBodyRecordTasksUserPanes) SetUserId(v string) *DescribeRecordTasksResponseBodyRecordTasksUserPanes {
	s.UserId = &v
	return s
}

func (s *DescribeRecordTasksResponseBodyRecordTasksUserPanes) SetPaneId(v int32) *DescribeRecordTasksResponseBodyRecordTasksUserPanes {
	s.PaneId = &v
	return s
}

func (s *DescribeRecordTasksResponseBodyRecordTasksUserPanes) SetSource(v string) *DescribeRecordTasksResponseBodyRecordTasksUserPanes {
	s.Source = &v
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
	OwnerId     *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog     *string   `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId       *string   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	PageSize    *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNum     *int32    `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	TemplateIds []*string `json:"TemplateIds,omitempty" xml:"TemplateIds,omitempty" type:"Repeated"`
}

func (s DescribeRecordTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecordTemplatesRequest) SetOwnerId(v int64) *DescribeRecordTemplatesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRecordTemplatesRequest) SetShowLog(v string) *DescribeRecordTemplatesRequest {
	s.ShowLog = &v
	return s
}

func (s *DescribeRecordTemplatesRequest) SetAppId(v string) *DescribeRecordTemplatesRequest {
	s.AppId = &v
	return s
}

func (s *DescribeRecordTemplatesRequest) SetPageSize(v int32) *DescribeRecordTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRecordTemplatesRequest) SetPageNum(v int32) *DescribeRecordTemplatesRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeRecordTemplatesRequest) SetTemplateIds(v []*string) *DescribeRecordTemplatesRequest {
	s.TemplateIds = v
	return s
}

type DescribeRecordTemplatesResponseBody struct {
	TotalNum  *int64                                          `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	TotalPage *int64                                          `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Templates []*DescribeRecordTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
}

func (s DescribeRecordTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecordTemplatesResponseBody) SetTotalNum(v int64) *DescribeRecordTemplatesResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBody) SetTotalPage(v int64) *DescribeRecordTemplatesResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBody) SetRequestId(v string) *DescribeRecordTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBody) SetTemplates(v []*DescribeRecordTemplatesResponseBodyTemplates) *DescribeRecordTemplatesResponseBody {
	s.Templates = v
	return s
}

type DescribeRecordTemplatesResponseBodyTemplates struct {
	MnsQueue          *string                                                     `json:"MnsQueue,omitempty" xml:"MnsQueue,omitempty"`
	OssFilePrefix     *string                                                     `json:"OssFilePrefix,omitempty" xml:"OssFilePrefix,omitempty"`
	CreateTime        *string                                                     `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ClockWidgets      []*DescribeRecordTemplatesResponseBodyTemplatesClockWidgets `json:"ClockWidgets,omitempty" xml:"ClockWidgets,omitempty" type:"Repeated"`
	OssBucket         *string                                                     `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	DelayStopTime     *string                                                     `json:"DelayStopTime,omitempty" xml:"DelayStopTime,omitempty"`
	LayoutIds         []*int32                                                    `json:"LayoutIds,omitempty" xml:"LayoutIds,omitempty" type:"Repeated"`
	MediaEncode       *int32                                                      `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	FileSplitInterval *int32                                                      `json:"FileSplitInterval,omitempty" xml:"FileSplitInterval,omitempty"`
	HttpCallbackUrl   *string                                                     `json:"HttpCallbackUrl,omitempty" xml:"HttpCallbackUrl,omitempty"`
	Formats           []*string                                                   `json:"Formats,omitempty" xml:"Formats,omitempty" type:"Repeated"`
	BackgroundColor   *int32                                                      `json:"BackgroundColor,omitempty" xml:"BackgroundColor,omitempty"`
	Backgrounds       []*DescribeRecordTemplatesResponseBodyTemplatesBackgrounds  `json:"Backgrounds,omitempty" xml:"Backgrounds,omitempty" type:"Repeated"`
	Watermarks        []*DescribeRecordTemplatesResponseBodyTemplatesWatermarks   `json:"Watermarks,omitempty" xml:"Watermarks,omitempty" type:"Repeated"`
	Name              *string                                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	TemplateId        *string                                                     `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TaskProfile       *string                                                     `json:"TaskProfile,omitempty" xml:"TaskProfile,omitempty"`
}

func (s DescribeRecordTemplatesResponseBodyTemplates) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) SetMnsQueue(v string) *DescribeRecordTemplatesResponseBodyTemplates {
	s.MnsQueue = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) SetOssFilePrefix(v string) *DescribeRecordTemplatesResponseBodyTemplates {
	s.OssFilePrefix = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) SetCreateTime(v string) *DescribeRecordTemplatesResponseBodyTemplates {
	s.CreateTime = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) SetClockWidgets(v []*DescribeRecordTemplatesResponseBodyTemplatesClockWidgets) *DescribeRecordTemplatesResponseBodyTemplates {
	s.ClockWidgets = v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) SetOssBucket(v string) *DescribeRecordTemplatesResponseBodyTemplates {
	s.OssBucket = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) SetDelayStopTime(v string) *DescribeRecordTemplatesResponseBodyTemplates {
	s.DelayStopTime = &v
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

func (s *DescribeRecordTemplatesResponseBodyTemplates) SetFileSplitInterval(v int32) *DescribeRecordTemplatesResponseBodyTemplates {
	s.FileSplitInterval = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) SetHttpCallbackUrl(v string) *DescribeRecordTemplatesResponseBodyTemplates {
	s.HttpCallbackUrl = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) SetFormats(v []*string) *DescribeRecordTemplatesResponseBodyTemplates {
	s.Formats = v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) SetBackgroundColor(v int32) *DescribeRecordTemplatesResponseBodyTemplates {
	s.BackgroundColor = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) SetBackgrounds(v []*DescribeRecordTemplatesResponseBodyTemplatesBackgrounds) *DescribeRecordTemplatesResponseBodyTemplates {
	s.Backgrounds = v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) SetWatermarks(v []*DescribeRecordTemplatesResponseBodyTemplatesWatermarks) *DescribeRecordTemplatesResponseBodyTemplates {
	s.Watermarks = v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) SetName(v string) *DescribeRecordTemplatesResponseBodyTemplates {
	s.Name = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) SetTemplateId(v string) *DescribeRecordTemplatesResponseBodyTemplates {
	s.TemplateId = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplates) SetTaskProfile(v string) *DescribeRecordTemplatesResponseBodyTemplates {
	s.TaskProfile = &v
	return s
}

type DescribeRecordTemplatesResponseBodyTemplatesClockWidgets struct {
	FontType  *int32   `json:"FontType,omitempty" xml:"FontType,omitempty"`
	FontColor *int32   `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	Y         *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	ZOrder    *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
	X         *float32 `json:"X,omitempty" xml:"X,omitempty"`
	FontSize  *int32   `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
}

func (s DescribeRecordTemplatesResponseBodyTemplatesClockWidgets) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordTemplatesResponseBodyTemplatesClockWidgets) GoString() string {
	return s.String()
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesClockWidgets) SetFontType(v int32) *DescribeRecordTemplatesResponseBodyTemplatesClockWidgets {
	s.FontType = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesClockWidgets) SetFontColor(v int32) *DescribeRecordTemplatesResponseBodyTemplatesClockWidgets {
	s.FontColor = &v
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

func (s *DescribeRecordTemplatesResponseBodyTemplatesClockWidgets) SetX(v float32) *DescribeRecordTemplatesResponseBodyTemplatesClockWidgets {
	s.X = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesClockWidgets) SetFontSize(v int32) *DescribeRecordTemplatesResponseBodyTemplatesClockWidgets {
	s.FontSize = &v
	return s
}

type DescribeRecordTemplatesResponseBodyTemplatesBackgrounds struct {
	Width   *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height  *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Y       *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	Url     *string  `json:"Url,omitempty" xml:"Url,omitempty"`
	Display *int32   `json:"Display,omitempty" xml:"Display,omitempty"`
	ZOrder  *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
	X       *float32 `json:"X,omitempty" xml:"X,omitempty"`
}

func (s DescribeRecordTemplatesResponseBodyTemplatesBackgrounds) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecordTemplatesResponseBodyTemplatesBackgrounds) GoString() string {
	return s.String()
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesBackgrounds) SetWidth(v float32) *DescribeRecordTemplatesResponseBodyTemplatesBackgrounds {
	s.Width = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesBackgrounds) SetHeight(v float32) *DescribeRecordTemplatesResponseBodyTemplatesBackgrounds {
	s.Height = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesBackgrounds) SetY(v float32) *DescribeRecordTemplatesResponseBodyTemplatesBackgrounds {
	s.Y = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesBackgrounds) SetUrl(v string) *DescribeRecordTemplatesResponseBodyTemplatesBackgrounds {
	s.Url = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesBackgrounds) SetDisplay(v int32) *DescribeRecordTemplatesResponseBodyTemplatesBackgrounds {
	s.Display = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesBackgrounds) SetZOrder(v int32) *DescribeRecordTemplatesResponseBodyTemplatesBackgrounds {
	s.ZOrder = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesBackgrounds) SetX(v float32) *DescribeRecordTemplatesResponseBodyTemplatesBackgrounds {
	s.X = &v
	return s
}

type DescribeRecordTemplatesResponseBodyTemplatesWatermarks struct {
	Alpha   *float32 `json:"Alpha,omitempty" xml:"Alpha,omitempty"`
	Width   *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height  *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Y       *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	Url     *string  `json:"Url,omitempty" xml:"Url,omitempty"`
	Display *int32   `json:"Display,omitempty" xml:"Display,omitempty"`
	ZOrder  *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
	X       *float32 `json:"X,omitempty" xml:"X,omitempty"`
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

func (s *DescribeRecordTemplatesResponseBodyTemplatesWatermarks) SetWidth(v float32) *DescribeRecordTemplatesResponseBodyTemplatesWatermarks {
	s.Width = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesWatermarks) SetHeight(v float32) *DescribeRecordTemplatesResponseBodyTemplatesWatermarks {
	s.Height = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesWatermarks) SetY(v float32) *DescribeRecordTemplatesResponseBodyTemplatesWatermarks {
	s.Y = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesWatermarks) SetUrl(v string) *DescribeRecordTemplatesResponseBodyTemplatesWatermarks {
	s.Url = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesWatermarks) SetDisplay(v int32) *DescribeRecordTemplatesResponseBodyTemplatesWatermarks {
	s.Display = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesWatermarks) SetZOrder(v int32) *DescribeRecordTemplatesResponseBodyTemplatesWatermarks {
	s.ZOrder = &v
	return s
}

func (s *DescribeRecordTemplatesResponseBodyTemplatesWatermarks) SetX(v float32) *DescribeRecordTemplatesResponseBodyTemplatesWatermarks {
	s.X = &v
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

type DescribeRTCAppKeyRequest struct {
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DescribeRTCAppKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRTCAppKeyRequest) GoString() string {
	return s.String()
}

func (s *DescribeRTCAppKeyRequest) SetOwnerId(v int64) *DescribeRTCAppKeyRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRTCAppKeyRequest) SetShowLog(v string) *DescribeRTCAppKeyRequest {
	s.ShowLog = &v
	return s
}

func (s *DescribeRTCAppKeyRequest) SetAppId(v string) *DescribeRTCAppKeyRequest {
	s.AppId = &v
	return s
}

type DescribeRTCAppKeyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	AppKey    *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
}

func (s DescribeRTCAppKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRTCAppKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRTCAppKeyResponseBody) SetRequestId(v string) *DescribeRTCAppKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRTCAppKeyResponseBody) SetAppKey(v string) *DescribeRTCAppKeyResponseBody {
	s.AppKey = &v
	return s
}

type DescribeRTCAppKeyResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRTCAppKeyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRTCAppKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRTCAppKeyResponse) GoString() string {
	return s.String()
}

func (s *DescribeRTCAppKeyResponse) SetHeaders(v map[string]*string) *DescribeRTCAppKeyResponse {
	s.Headers = v
	return s
}

func (s *DescribeRTCAppKeyResponse) SetBody(v *DescribeRTCAppKeyResponseBody) *DescribeRTCAppKeyResponse {
	s.Body = v
	return s
}

type DescribeRtcChannelCntDataRequest struct {
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog     *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ServiceArea *string `json:"ServiceArea,omitempty" xml:"ServiceArea,omitempty"`
	Interval    *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s DescribeRtcChannelCntDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelCntDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelCntDataRequest) SetOwnerId(v int64) *DescribeRtcChannelCntDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRtcChannelCntDataRequest) SetShowLog(v string) *DescribeRtcChannelCntDataRequest {
	s.ShowLog = &v
	return s
}

func (s *DescribeRtcChannelCntDataRequest) SetStartTime(v string) *DescribeRtcChannelCntDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRtcChannelCntDataRequest) SetEndTime(v string) *DescribeRtcChannelCntDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRtcChannelCntDataRequest) SetAppId(v string) *DescribeRtcChannelCntDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeRtcChannelCntDataRequest) SetServiceArea(v string) *DescribeRtcChannelCntDataRequest {
	s.ServiceArea = &v
	return s
}

func (s *DescribeRtcChannelCntDataRequest) SetInterval(v string) *DescribeRtcChannelCntDataRequest {
	s.Interval = &v
	return s
}

type DescribeRtcChannelCntDataResponseBody struct {
	RequestId                 *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ChannelCntDataPerInterval *DescribeRtcChannelCntDataResponseBodyChannelCntDataPerInterval `json:"ChannelCntDataPerInterval,omitempty" xml:"ChannelCntDataPerInterval,omitempty" type:"Struct"`
}

func (s DescribeRtcChannelCntDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelCntDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelCntDataResponseBody) SetRequestId(v string) *DescribeRtcChannelCntDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRtcChannelCntDataResponseBody) SetChannelCntDataPerInterval(v *DescribeRtcChannelCntDataResponseBodyChannelCntDataPerInterval) *DescribeRtcChannelCntDataResponseBody {
	s.ChannelCntDataPerInterval = v
	return s
}

type DescribeRtcChannelCntDataResponseBodyChannelCntDataPerInterval struct {
	ChannelCntModule []*DescribeRtcChannelCntDataResponseBodyChannelCntDataPerIntervalChannelCntModule `json:"ChannelCntModule,omitempty" xml:"ChannelCntModule,omitempty" type:"Repeated"`
}

func (s DescribeRtcChannelCntDataResponseBodyChannelCntDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelCntDataResponseBodyChannelCntDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelCntDataResponseBodyChannelCntDataPerInterval) SetChannelCntModule(v []*DescribeRtcChannelCntDataResponseBodyChannelCntDataPerIntervalChannelCntModule) *DescribeRtcChannelCntDataResponseBodyChannelCntDataPerInterval {
	s.ChannelCntModule = v
	return s
}

type DescribeRtcChannelCntDataResponseBodyChannelCntDataPerIntervalChannelCntModule struct {
	ActiveChannelCnt *int64  `json:"ActiveChannelCnt,omitempty" xml:"ActiveChannelCnt,omitempty"`
	TimeStamp        *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeRtcChannelCntDataResponseBodyChannelCntDataPerIntervalChannelCntModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelCntDataResponseBodyChannelCntDataPerIntervalChannelCntModule) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelCntDataResponseBodyChannelCntDataPerIntervalChannelCntModule) SetActiveChannelCnt(v int64) *DescribeRtcChannelCntDataResponseBodyChannelCntDataPerIntervalChannelCntModule {
	s.ActiveChannelCnt = &v
	return s
}

func (s *DescribeRtcChannelCntDataResponseBodyChannelCntDataPerIntervalChannelCntModule) SetTimeStamp(v string) *DescribeRtcChannelCntDataResponseBodyChannelCntDataPerIntervalChannelCntModule {
	s.TimeStamp = &v
	return s
}

type DescribeRtcChannelCntDataResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRtcChannelCntDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRtcChannelCntDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelCntDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelCntDataResponse) SetHeaders(v map[string]*string) *DescribeRtcChannelCntDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeRtcChannelCntDataResponse) SetBody(v *DescribeRtcChannelCntDataResponseBody) *DescribeRtcChannelCntDataResponse {
	s.Body = v
	return s
}

type DescribeRtcChannelDetailRequest struct {
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog   *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeRtcChannelDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelDetailRequest) SetOwnerId(v int64) *DescribeRtcChannelDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRtcChannelDetailRequest) SetShowLog(v string) *DescribeRtcChannelDetailRequest {
	s.ShowLog = &v
	return s
}

func (s *DescribeRtcChannelDetailRequest) SetAppId(v string) *DescribeRtcChannelDetailRequest {
	s.AppId = &v
	return s
}

func (s *DescribeRtcChannelDetailRequest) SetChannelId(v string) *DescribeRtcChannelDetailRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeRtcChannelDetailRequest) SetStartTime(v string) *DescribeRtcChannelDetailRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRtcChannelDetailRequest) SetEndTime(v string) *DescribeRtcChannelDetailRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRtcChannelDetailRequest) SetPageNo(v int32) *DescribeRtcChannelDetailRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeRtcChannelDetailRequest) SetPageSize(v int32) *DescribeRtcChannelDetailRequest {
	s.PageSize = &v
	return s
}

type DescribeRtcChannelDetailResponseBody struct {
	TotalCnt    *int64                                             `json:"TotalCnt,omitempty" xml:"TotalCnt,omitempty"`
	PageSize    *int64                                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNo      *int64                                             `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	ChannelInfo []*DescribeRtcChannelDetailResponseBodyChannelInfo `json:"ChannelInfo,omitempty" xml:"ChannelInfo,omitempty" type:"Repeated"`
	ChannelId   *string                                            `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
}

func (s DescribeRtcChannelDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelDetailResponseBody) SetTotalCnt(v int64) *DescribeRtcChannelDetailResponseBody {
	s.TotalCnt = &v
	return s
}

func (s *DescribeRtcChannelDetailResponseBody) SetPageSize(v int64) *DescribeRtcChannelDetailResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRtcChannelDetailResponseBody) SetRequestId(v string) *DescribeRtcChannelDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRtcChannelDetailResponseBody) SetPageNo(v int64) *DescribeRtcChannelDetailResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribeRtcChannelDetailResponseBody) SetChannelInfo(v []*DescribeRtcChannelDetailResponseBodyChannelInfo) *DescribeRtcChannelDetailResponseBody {
	s.ChannelInfo = v
	return s
}

func (s *DescribeRtcChannelDetailResponseBody) SetChannelId(v string) *DescribeRtcChannelDetailResponseBody {
	s.ChannelId = &v
	return s
}

type DescribeRtcChannelDetailResponseBodyChannelInfo struct {
	Sid        *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	OS         *string `json:"OS,omitempty" xml:"OS,omitempty"`
	LeaveTime  *string `json:"LeaveTime,omitempty" xml:"LeaveTime,omitempty"`
	JoinTime   *string `json:"JoinTime,omitempty" xml:"JoinTime,omitempty"`
	Platform   *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	SdkVersion *string `json:"SdkVersion,omitempty" xml:"SdkVersion,omitempty"`
	Uid        *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s DescribeRtcChannelDetailResponseBodyChannelInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelDetailResponseBodyChannelInfo) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelDetailResponseBodyChannelInfo) SetSid(v string) *DescribeRtcChannelDetailResponseBodyChannelInfo {
	s.Sid = &v
	return s
}

func (s *DescribeRtcChannelDetailResponseBodyChannelInfo) SetDeviceType(v string) *DescribeRtcChannelDetailResponseBodyChannelInfo {
	s.DeviceType = &v
	return s
}

func (s *DescribeRtcChannelDetailResponseBodyChannelInfo) SetOS(v string) *DescribeRtcChannelDetailResponseBodyChannelInfo {
	s.OS = &v
	return s
}

func (s *DescribeRtcChannelDetailResponseBodyChannelInfo) SetLeaveTime(v string) *DescribeRtcChannelDetailResponseBodyChannelInfo {
	s.LeaveTime = &v
	return s
}

func (s *DescribeRtcChannelDetailResponseBodyChannelInfo) SetJoinTime(v string) *DescribeRtcChannelDetailResponseBodyChannelInfo {
	s.JoinTime = &v
	return s
}

func (s *DescribeRtcChannelDetailResponseBodyChannelInfo) SetPlatform(v string) *DescribeRtcChannelDetailResponseBodyChannelInfo {
	s.Platform = &v
	return s
}

func (s *DescribeRtcChannelDetailResponseBodyChannelInfo) SetSdkVersion(v string) *DescribeRtcChannelDetailResponseBodyChannelInfo {
	s.SdkVersion = &v
	return s
}

func (s *DescribeRtcChannelDetailResponseBodyChannelInfo) SetUid(v string) *DescribeRtcChannelDetailResponseBodyChannelInfo {
	s.Uid = &v
	return s
}

type DescribeRtcChannelDetailResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRtcChannelDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRtcChannelDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelDetailResponse) SetHeaders(v map[string]*string) *DescribeRtcChannelDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeRtcChannelDetailResponse) SetBody(v *DescribeRtcChannelDetailResponseBody) *DescribeRtcChannelDetailResponse {
	s.Body = v
	return s
}

type DescribeRtcChannelListRequest struct {
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog     *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	TimePoint   *string `json:"TimePoint,omitempty" xml:"TimePoint,omitempty"`
	SortType    *string `json:"SortType,omitempty" xml:"SortType,omitempty"`
	ServiceArea *string `json:"ServiceArea,omitempty" xml:"ServiceArea,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	ChannelId   *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	PageNo      *int64  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize    *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeRtcChannelListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelListRequest) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelListRequest) SetOwnerId(v int64) *DescribeRtcChannelListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRtcChannelListRequest) SetShowLog(v string) *DescribeRtcChannelListRequest {
	s.ShowLog = &v
	return s
}

func (s *DescribeRtcChannelListRequest) SetAppId(v string) *DescribeRtcChannelListRequest {
	s.AppId = &v
	return s
}

func (s *DescribeRtcChannelListRequest) SetTimePoint(v string) *DescribeRtcChannelListRequest {
	s.TimePoint = &v
	return s
}

func (s *DescribeRtcChannelListRequest) SetSortType(v string) *DescribeRtcChannelListRequest {
	s.SortType = &v
	return s
}

func (s *DescribeRtcChannelListRequest) SetServiceArea(v string) *DescribeRtcChannelListRequest {
	s.ServiceArea = &v
	return s
}

func (s *DescribeRtcChannelListRequest) SetUserId(v string) *DescribeRtcChannelListRequest {
	s.UserId = &v
	return s
}

func (s *DescribeRtcChannelListRequest) SetChannelId(v string) *DescribeRtcChannelListRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeRtcChannelListRequest) SetPageNo(v int64) *DescribeRtcChannelListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeRtcChannelListRequest) SetPageSize(v int64) *DescribeRtcChannelListRequest {
	s.PageSize = &v
	return s
}

type DescribeRtcChannelListResponseBody struct {
	TotalCnt    *int64                                         `json:"TotalCnt,omitempty" xml:"TotalCnt,omitempty"`
	RequestId   *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize    *int64                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNo      *int64                                         `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	ChannelList *DescribeRtcChannelListResponseBodyChannelList `json:"ChannelList,omitempty" xml:"ChannelList,omitempty" type:"Struct"`
}

func (s DescribeRtcChannelListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelListResponseBody) SetTotalCnt(v int64) *DescribeRtcChannelListResponseBody {
	s.TotalCnt = &v
	return s
}

func (s *DescribeRtcChannelListResponseBody) SetRequestId(v string) *DescribeRtcChannelListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRtcChannelListResponseBody) SetPageSize(v int64) *DescribeRtcChannelListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRtcChannelListResponseBody) SetPageNo(v int64) *DescribeRtcChannelListResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribeRtcChannelListResponseBody) SetChannelList(v *DescribeRtcChannelListResponseBodyChannelList) *DescribeRtcChannelListResponseBody {
	s.ChannelList = v
	return s
}

type DescribeRtcChannelListResponseBodyChannelList struct {
	ChannelList []*DescribeRtcChannelListResponseBodyChannelListChannelList `json:"ChannelList,omitempty" xml:"ChannelList,omitempty" type:"Repeated"`
}

func (s DescribeRtcChannelListResponseBodyChannelList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelListResponseBodyChannelList) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelListResponseBodyChannelList) SetChannelList(v []*DescribeRtcChannelListResponseBodyChannelListChannelList) *DescribeRtcChannelListResponseBodyChannelList {
	s.ChannelList = v
	return s
}

type DescribeRtcChannelListResponseBodyChannelListChannelList struct {
	EndTime      *string                                                           `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	TotalUserCnt *int64                                                            `json:"TotalUserCnt,omitempty" xml:"TotalUserCnt,omitempty"`
	StartTime    *string                                                           `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	CallArea     *DescribeRtcChannelListResponseBodyChannelListChannelListCallArea `json:"CallArea,omitempty" xml:"CallArea,omitempty" type:"Struct"`
	ChannelId    *string                                                           `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
}

func (s DescribeRtcChannelListResponseBodyChannelListChannelList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelListResponseBodyChannelListChannelList) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelListResponseBodyChannelListChannelList) SetEndTime(v string) *DescribeRtcChannelListResponseBodyChannelListChannelList {
	s.EndTime = &v
	return s
}

func (s *DescribeRtcChannelListResponseBodyChannelListChannelList) SetTotalUserCnt(v int64) *DescribeRtcChannelListResponseBodyChannelListChannelList {
	s.TotalUserCnt = &v
	return s
}

func (s *DescribeRtcChannelListResponseBodyChannelListChannelList) SetStartTime(v string) *DescribeRtcChannelListResponseBodyChannelListChannelList {
	s.StartTime = &v
	return s
}

func (s *DescribeRtcChannelListResponseBodyChannelListChannelList) SetCallArea(v *DescribeRtcChannelListResponseBodyChannelListChannelListCallArea) *DescribeRtcChannelListResponseBodyChannelListChannelList {
	s.CallArea = v
	return s
}

func (s *DescribeRtcChannelListResponseBodyChannelListChannelList) SetChannelId(v string) *DescribeRtcChannelListResponseBodyChannelListChannelList {
	s.ChannelId = &v
	return s
}

type DescribeRtcChannelListResponseBodyChannelListChannelListCallArea struct {
	CallArea []*string `json:"CallArea,omitempty" xml:"CallArea,omitempty" type:"Repeated"`
}

func (s DescribeRtcChannelListResponseBodyChannelListChannelListCallArea) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelListResponseBodyChannelListChannelListCallArea) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelListResponseBodyChannelListChannelListCallArea) SetCallArea(v []*string) *DescribeRtcChannelListResponseBodyChannelListChannelListCallArea {
	s.CallArea = v
	return s
}

type DescribeRtcChannelListResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRtcChannelListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRtcChannelListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelListResponse) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelListResponse) SetHeaders(v map[string]*string) *DescribeRtcChannelListResponse {
	s.Headers = v
	return s
}

func (s *DescribeRtcChannelListResponse) SetBody(v *DescribeRtcChannelListResponseBody) *DescribeRtcChannelListResponse {
	s.Body = v
	return s
}

type DescribeRtcChannelMetricRequest struct {
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog   *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	TimePoint *string `json:"TimePoint,omitempty" xml:"TimePoint,omitempty"`
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
}

func (s DescribeRtcChannelMetricRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelMetricRequest) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelMetricRequest) SetOwnerId(v int64) *DescribeRtcChannelMetricRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRtcChannelMetricRequest) SetShowLog(v string) *DescribeRtcChannelMetricRequest {
	s.ShowLog = &v
	return s
}

func (s *DescribeRtcChannelMetricRequest) SetTimePoint(v string) *DescribeRtcChannelMetricRequest {
	s.TimePoint = &v
	return s
}

func (s *DescribeRtcChannelMetricRequest) SetAppId(v string) *DescribeRtcChannelMetricRequest {
	s.AppId = &v
	return s
}

func (s *DescribeRtcChannelMetricRequest) SetChannelId(v string) *DescribeRtcChannelMetricRequest {
	s.ChannelId = &v
	return s
}

type DescribeRtcChannelMetricResponseBody struct {
	RequestId         *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ChannelMetricInfo *DescribeRtcChannelMetricResponseBodyChannelMetricInfo `json:"ChannelMetricInfo,omitempty" xml:"ChannelMetricInfo,omitempty" type:"Struct"`
}

func (s DescribeRtcChannelMetricResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelMetricResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelMetricResponseBody) SetRequestId(v string) *DescribeRtcChannelMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRtcChannelMetricResponseBody) SetChannelMetricInfo(v *DescribeRtcChannelMetricResponseBodyChannelMetricInfo) *DescribeRtcChannelMetricResponseBody {
	s.ChannelMetricInfo = v
	return s
}

type DescribeRtcChannelMetricResponseBodyChannelMetricInfo struct {
	Duration      *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDuration      `json:"Duration,omitempty" xml:"Duration,omitempty" type:"Struct"`
	ChannelMetric *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric `json:"ChannelMetric,omitempty" xml:"ChannelMetric,omitempty" type:"Struct"`
}

func (s DescribeRtcChannelMetricResponseBodyChannelMetricInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelMetricResponseBodyChannelMetricInfo) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfo) SetDuration(v *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDuration) *DescribeRtcChannelMetricResponseBodyChannelMetricInfo {
	s.Duration = v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfo) SetChannelMetric(v *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric) *DescribeRtcChannelMetricResponseBodyChannelMetricInfo {
	s.ChannelMetric = v
	return s
}

type DescribeRtcChannelMetricResponseBodyChannelMetricInfoDuration struct {
	SubDuration *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration `json:"SubDuration,omitempty" xml:"SubDuration,omitempty" type:"Struct"`
	PubDuration *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration `json:"PubDuration,omitempty" xml:"PubDuration,omitempty" type:"Struct"`
}

func (s DescribeRtcChannelMetricResponseBodyChannelMetricInfoDuration) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelMetricResponseBodyChannelMetricInfoDuration) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDuration) SetSubDuration(v *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDuration {
	s.SubDuration = v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDuration) SetPubDuration(v *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDuration {
	s.PubDuration = v
	return s
}

type DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration struct {
	Video720  *int32 `json:"Video720,omitempty" xml:"Video720,omitempty"`
	Video1080 *int32 `json:"Video1080,omitempty" xml:"Video1080,omitempty"`
	Video360  *int32 `json:"Video360,omitempty" xml:"Video360,omitempty"`
	Content   *int32 `json:"Content,omitempty" xml:"Content,omitempty"`
	Audio     *int32 `json:"Audio,omitempty" xml:"Audio,omitempty"`
}

func (s DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration) SetVideo720(v int32) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration {
	s.Video720 = &v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration) SetVideo1080(v int32) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration {
	s.Video1080 = &v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration) SetVideo360(v int32) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration {
	s.Video360 = &v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration) SetContent(v int32) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration {
	s.Content = &v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration) SetAudio(v int32) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration {
	s.Audio = &v
	return s
}

type DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration struct {
	Video720  *int32 `json:"Video720,omitempty" xml:"Video720,omitempty"`
	Video1080 *int32 `json:"Video1080,omitempty" xml:"Video1080,omitempty"`
	Video360  *int32 `json:"Video360,omitempty" xml:"Video360,omitempty"`
	Content   *int32 `json:"Content,omitempty" xml:"Content,omitempty"`
	Audio     *int32 `json:"Audio,omitempty" xml:"Audio,omitempty"`
}

func (s DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration) SetVideo720(v int32) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration {
	s.Video720 = &v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration) SetVideo1080(v int32) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration {
	s.Video1080 = &v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration) SetVideo360(v int32) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration {
	s.Video360 = &v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration) SetContent(v int32) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration {
	s.Content = &v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration) SetAudio(v int32) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration {
	s.Audio = &v
	return s
}

type DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric struct {
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	SubUserCount *int32  `json:"SubUserCount,omitempty" xml:"SubUserCount,omitempty"`
	ChannelId    *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	UserCount    *int32  `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
	PubUserCount *int32  `json:"PubUserCount,omitempty" xml:"PubUserCount,omitempty"`
}

func (s DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric) SetEndTime(v string) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric {
	s.EndTime = &v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric) SetStartTime(v string) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric {
	s.StartTime = &v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric) SetSubUserCount(v int32) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric {
	s.SubUserCount = &v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric) SetChannelId(v string) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric {
	s.ChannelId = &v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric) SetUserCount(v int32) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric {
	s.UserCount = &v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric) SetPubUserCount(v int32) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric {
	s.PubUserCount = &v
	return s
}

type DescribeRtcChannelMetricResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRtcChannelMetricResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRtcChannelMetricResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelMetricResponse) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelMetricResponse) SetHeaders(v map[string]*string) *DescribeRtcChannelMetricResponse {
	s.Headers = v
	return s
}

func (s *DescribeRtcChannelMetricResponse) SetBody(v *DescribeRtcChannelMetricResponseBody) *DescribeRtcChannelMetricResponse {
	s.Body = v
	return s
}

type DescribeRtcChannelMetricsRequest struct {
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog   *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	PubUid    *string `json:"PubUid,omitempty" xml:"PubUid,omitempty"`
	SubUid    *string `json:"SubUid,omitempty" xml:"SubUid,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeRtcChannelMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelMetricsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelMetricsRequest) SetOwnerId(v int64) *DescribeRtcChannelMetricsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRtcChannelMetricsRequest) SetShowLog(v string) *DescribeRtcChannelMetricsRequest {
	s.ShowLog = &v
	return s
}

func (s *DescribeRtcChannelMetricsRequest) SetAppId(v string) *DescribeRtcChannelMetricsRequest {
	s.AppId = &v
	return s
}

func (s *DescribeRtcChannelMetricsRequest) SetChannelId(v string) *DescribeRtcChannelMetricsRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeRtcChannelMetricsRequest) SetPubUid(v string) *DescribeRtcChannelMetricsRequest {
	s.PubUid = &v
	return s
}

func (s *DescribeRtcChannelMetricsRequest) SetSubUid(v string) *DescribeRtcChannelMetricsRequest {
	s.SubUid = &v
	return s
}

func (s *DescribeRtcChannelMetricsRequest) SetStartTime(v string) *DescribeRtcChannelMetricsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRtcChannelMetricsRequest) SetEndTime(v string) *DescribeRtcChannelMetricsRequest {
	s.EndTime = &v
	return s
}

type DescribeRtcChannelMetricsResponseBody struct {
	Metrics   []*DescribeRtcChannelMetricsResponseBodyMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRtcChannelMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelMetricsResponseBody) SetMetrics(v []*DescribeRtcChannelMetricsResponseBodyMetrics) *DescribeRtcChannelMetricsResponseBody {
	s.Metrics = v
	return s
}

func (s *DescribeRtcChannelMetricsResponseBody) SetRequestId(v string) *DescribeRtcChannelMetricsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRtcChannelMetricsResponseBodyMetrics struct {
	Mid *string   `json:"Mid,omitempty" xml:"Mid,omitempty"`
	KVs []*string `json:"KVs,omitempty" xml:"KVs,omitempty" type:"Repeated"`
	Uid *string   `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s DescribeRtcChannelMetricsResponseBodyMetrics) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelMetricsResponseBodyMetrics) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelMetricsResponseBodyMetrics) SetMid(v string) *DescribeRtcChannelMetricsResponseBodyMetrics {
	s.Mid = &v
	return s
}

func (s *DescribeRtcChannelMetricsResponseBodyMetrics) SetKVs(v []*string) *DescribeRtcChannelMetricsResponseBodyMetrics {
	s.KVs = v
	return s
}

func (s *DescribeRtcChannelMetricsResponseBodyMetrics) SetUid(v string) *DescribeRtcChannelMetricsResponseBodyMetrics {
	s.Uid = &v
	return s
}

type DescribeRtcChannelMetricsResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRtcChannelMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRtcChannelMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelMetricsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelMetricsResponse) SetHeaders(v map[string]*string) *DescribeRtcChannelMetricsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRtcChannelMetricsResponse) SetBody(v *DescribeRtcChannelMetricsResponseBody) *DescribeRtcChannelMetricsResponse {
	s.Body = v
	return s
}

type DescribeRtcChannelsRequest struct {
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog   *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	PageNo    *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeRtcChannelsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelsRequest) SetOwnerId(v int64) *DescribeRtcChannelsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRtcChannelsRequest) SetShowLog(v string) *DescribeRtcChannelsRequest {
	s.ShowLog = &v
	return s
}

func (s *DescribeRtcChannelsRequest) SetAppId(v string) *DescribeRtcChannelsRequest {
	s.AppId = &v
	return s
}

func (s *DescribeRtcChannelsRequest) SetStartTime(v string) *DescribeRtcChannelsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRtcChannelsRequest) SetEndTime(v string) *DescribeRtcChannelsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRtcChannelsRequest) SetChannelId(v string) *DescribeRtcChannelsRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeRtcChannelsRequest) SetPageNo(v int32) *DescribeRtcChannelsRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeRtcChannelsRequest) SetPageSize(v int32) *DescribeRtcChannelsRequest {
	s.PageSize = &v
	return s
}

type DescribeRtcChannelsResponseBody struct {
	TotalCnt  *int64                                     `json:"TotalCnt,omitempty" xml:"TotalCnt,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize  *int64                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNo    *int64                                     `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	Channels  []*DescribeRtcChannelsResponseBodyChannels `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Repeated"`
}

func (s DescribeRtcChannelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelsResponseBody) SetTotalCnt(v int64) *DescribeRtcChannelsResponseBody {
	s.TotalCnt = &v
	return s
}

func (s *DescribeRtcChannelsResponseBody) SetRequestId(v string) *DescribeRtcChannelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRtcChannelsResponseBody) SetPageSize(v int64) *DescribeRtcChannelsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRtcChannelsResponseBody) SetPageNo(v int64) *DescribeRtcChannelsResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribeRtcChannelsResponseBody) SetChannels(v []*DescribeRtcChannelsResponseBodyChannels) *DescribeRtcChannelsResponseBody {
	s.Channels = v
	return s
}

type DescribeRtcChannelsResponseBodyChannels struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	Finished  *bool   `json:"Finished,omitempty" xml:"Finished,omitempty"`
}

func (s DescribeRtcChannelsResponseBodyChannels) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelsResponseBodyChannels) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelsResponseBodyChannels) SetEndTime(v string) *DescribeRtcChannelsResponseBodyChannels {
	s.EndTime = &v
	return s
}

func (s *DescribeRtcChannelsResponseBodyChannels) SetStartTime(v string) *DescribeRtcChannelsResponseBodyChannels {
	s.StartTime = &v
	return s
}

func (s *DescribeRtcChannelsResponseBodyChannels) SetChannelId(v string) *DescribeRtcChannelsResponseBodyChannels {
	s.ChannelId = &v
	return s
}

func (s *DescribeRtcChannelsResponseBodyChannels) SetFinished(v bool) *DescribeRtcChannelsResponseBodyChannels {
	s.Finished = &v
	return s
}

type DescribeRtcChannelsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRtcChannelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRtcChannelsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelsResponse) SetHeaders(v map[string]*string) *DescribeRtcChannelsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRtcChannelsResponse) SetBody(v *DescribeRtcChannelsResponseBody) *DescribeRtcChannelsResponse {
	s.Body = v
	return s
}

type DescribeRtcChannelUserListRequest struct {
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog   *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	PageNo    *int64  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize  *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TimePoint *string `json:"TimePoint,omitempty" xml:"TimePoint,omitempty"`
}

func (s DescribeRtcChannelUserListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelUserListRequest) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelUserListRequest) SetOwnerId(v int64) *DescribeRtcChannelUserListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRtcChannelUserListRequest) SetShowLog(v string) *DescribeRtcChannelUserListRequest {
	s.ShowLog = &v
	return s
}

func (s *DescribeRtcChannelUserListRequest) SetAppId(v string) *DescribeRtcChannelUserListRequest {
	s.AppId = &v
	return s
}

func (s *DescribeRtcChannelUserListRequest) SetChannelId(v string) *DescribeRtcChannelUserListRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeRtcChannelUserListRequest) SetPageNo(v int64) *DescribeRtcChannelUserListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeRtcChannelUserListRequest) SetPageSize(v int64) *DescribeRtcChannelUserListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRtcChannelUserListRequest) SetTimePoint(v string) *DescribeRtcChannelUserListRequest {
	s.TimePoint = &v
	return s
}

type DescribeRtcChannelUserListResponseBody struct {
	TotalCnt  *int64                                          `json:"TotalCnt,omitempty" xml:"TotalCnt,omitempty"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageSize  *int64                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNo    *int64                                          `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	UserList  *DescribeRtcChannelUserListResponseBodyUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Struct"`
}

func (s DescribeRtcChannelUserListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelUserListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelUserListResponseBody) SetTotalCnt(v int64) *DescribeRtcChannelUserListResponseBody {
	s.TotalCnt = &v
	return s
}

func (s *DescribeRtcChannelUserListResponseBody) SetRequestId(v string) *DescribeRtcChannelUserListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRtcChannelUserListResponseBody) SetPageSize(v int64) *DescribeRtcChannelUserListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRtcChannelUserListResponseBody) SetPageNo(v int64) *DescribeRtcChannelUserListResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribeRtcChannelUserListResponseBody) SetUserList(v *DescribeRtcChannelUserListResponseBodyUserList) *DescribeRtcChannelUserListResponseBody {
	s.UserList = v
	return s
}

type DescribeRtcChannelUserListResponseBodyUserList struct {
	UserList []*DescribeRtcChannelUserListResponseBodyUserListUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
}

func (s DescribeRtcChannelUserListResponseBodyUserList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelUserListResponseBodyUserList) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelUserListResponseBodyUserList) SetUserList(v []*DescribeRtcChannelUserListResponseBodyUserListUserList) *DescribeRtcChannelUserListResponseBodyUserList {
	s.UserList = v
	return s
}

type DescribeRtcChannelUserListResponseBodyUserListUserList struct {
	SubVideo720  *int32  `json:"SubVideo720,omitempty" xml:"SubVideo720,omitempty"`
	SubVideo1080 *int32  `json:"SubVideo1080,omitempty" xml:"SubVideo1080,omitempty"`
	SubContent   *int32  `json:"SubContent,omitempty" xml:"SubContent,omitempty"`
	UserId       *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	PubVideo360  *int32  `json:"PubVideo360,omitempty" xml:"PubVideo360,omitempty"`
	SubVideo360  *int32  `json:"SubVideo360,omitempty" xml:"SubVideo360,omitempty"`
	ServiceArea  *string `json:"ServiceArea,omitempty" xml:"ServiceArea,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	PubContent   *int32  `json:"PubContent,omitempty" xml:"PubContent,omitempty"`
	ChannelId    *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	PubVideo1080 *int32  `json:"PubVideo1080,omitempty" xml:"PubVideo1080,omitempty"`
	PubAudio     *int32  `json:"PubAudio,omitempty" xml:"PubAudio,omitempty"`
	PubVideo720  *int32  `json:"PubVideo720,omitempty" xml:"PubVideo720,omitempty"`
	SubAudio     *int32  `json:"SubAudio,omitempty" xml:"SubAudio,omitempty"`
}

func (s DescribeRtcChannelUserListResponseBodyUserListUserList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelUserListResponseBodyUserListUserList) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelUserListResponseBodyUserListUserList) SetSubVideo720(v int32) *DescribeRtcChannelUserListResponseBodyUserListUserList {
	s.SubVideo720 = &v
	return s
}

func (s *DescribeRtcChannelUserListResponseBodyUserListUserList) SetSubVideo1080(v int32) *DescribeRtcChannelUserListResponseBodyUserListUserList {
	s.SubVideo1080 = &v
	return s
}

func (s *DescribeRtcChannelUserListResponseBodyUserListUserList) SetSubContent(v int32) *DescribeRtcChannelUserListResponseBodyUserListUserList {
	s.SubContent = &v
	return s
}

func (s *DescribeRtcChannelUserListResponseBodyUserListUserList) SetUserId(v string) *DescribeRtcChannelUserListResponseBodyUserListUserList {
	s.UserId = &v
	return s
}

func (s *DescribeRtcChannelUserListResponseBodyUserListUserList) SetPubVideo360(v int32) *DescribeRtcChannelUserListResponseBodyUserListUserList {
	s.PubVideo360 = &v
	return s
}

func (s *DescribeRtcChannelUserListResponseBodyUserListUserList) SetSubVideo360(v int32) *DescribeRtcChannelUserListResponseBodyUserListUserList {
	s.SubVideo360 = &v
	return s
}

func (s *DescribeRtcChannelUserListResponseBodyUserListUserList) SetServiceArea(v string) *DescribeRtcChannelUserListResponseBodyUserListUserList {
	s.ServiceArea = &v
	return s
}

func (s *DescribeRtcChannelUserListResponseBodyUserListUserList) SetEndTime(v string) *DescribeRtcChannelUserListResponseBodyUserListUserList {
	s.EndTime = &v
	return s
}

func (s *DescribeRtcChannelUserListResponseBodyUserListUserList) SetStartTime(v string) *DescribeRtcChannelUserListResponseBodyUserListUserList {
	s.StartTime = &v
	return s
}

func (s *DescribeRtcChannelUserListResponseBodyUserListUserList) SetPubContent(v int32) *DescribeRtcChannelUserListResponseBodyUserListUserList {
	s.PubContent = &v
	return s
}

func (s *DescribeRtcChannelUserListResponseBodyUserListUserList) SetChannelId(v string) *DescribeRtcChannelUserListResponseBodyUserListUserList {
	s.ChannelId = &v
	return s
}

func (s *DescribeRtcChannelUserListResponseBodyUserListUserList) SetPubVideo1080(v int32) *DescribeRtcChannelUserListResponseBodyUserListUserList {
	s.PubVideo1080 = &v
	return s
}

func (s *DescribeRtcChannelUserListResponseBodyUserListUserList) SetPubAudio(v int32) *DescribeRtcChannelUserListResponseBodyUserListUserList {
	s.PubAudio = &v
	return s
}

func (s *DescribeRtcChannelUserListResponseBodyUserListUserList) SetPubVideo720(v int32) *DescribeRtcChannelUserListResponseBodyUserListUserList {
	s.PubVideo720 = &v
	return s
}

func (s *DescribeRtcChannelUserListResponseBodyUserListUserList) SetSubAudio(v int32) *DescribeRtcChannelUserListResponseBodyUserListUserList {
	s.SubAudio = &v
	return s
}

type DescribeRtcChannelUserListResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRtcChannelUserListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRtcChannelUserListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelUserListResponse) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelUserListResponse) SetHeaders(v map[string]*string) *DescribeRtcChannelUserListResponse {
	s.Headers = v
	return s
}

func (s *DescribeRtcChannelUserListResponse) SetBody(v *DescribeRtcChannelUserListResponseBody) *DescribeRtcChannelUserListResponse {
	s.Body = v
	return s
}

type DescribeRtcDurationDataRequest struct {
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog     *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ServiceArea *string `json:"ServiceArea,omitempty" xml:"ServiceArea,omitempty"`
	Interval    *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s DescribeRtcDurationDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcDurationDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeRtcDurationDataRequest) SetOwnerId(v int64) *DescribeRtcDurationDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRtcDurationDataRequest) SetShowLog(v string) *DescribeRtcDurationDataRequest {
	s.ShowLog = &v
	return s
}

func (s *DescribeRtcDurationDataRequest) SetStartTime(v string) *DescribeRtcDurationDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRtcDurationDataRequest) SetEndTime(v string) *DescribeRtcDurationDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRtcDurationDataRequest) SetAppId(v string) *DescribeRtcDurationDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeRtcDurationDataRequest) SetServiceArea(v string) *DescribeRtcDurationDataRequest {
	s.ServiceArea = &v
	return s
}

func (s *DescribeRtcDurationDataRequest) SetInterval(v string) *DescribeRtcDurationDataRequest {
	s.Interval = &v
	return s
}

type DescribeRtcDurationDataResponseBody struct {
	RequestId               *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DurationDataPerInterval *DescribeRtcDurationDataResponseBodyDurationDataPerInterval `json:"DurationDataPerInterval,omitempty" xml:"DurationDataPerInterval,omitempty" type:"Struct"`
}

func (s DescribeRtcDurationDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcDurationDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRtcDurationDataResponseBody) SetRequestId(v string) *DescribeRtcDurationDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRtcDurationDataResponseBody) SetDurationDataPerInterval(v *DescribeRtcDurationDataResponseBodyDurationDataPerInterval) *DescribeRtcDurationDataResponseBody {
	s.DurationDataPerInterval = v
	return s
}

type DescribeRtcDurationDataResponseBodyDurationDataPerInterval struct {
	DurationModule []*DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule `json:"DurationModule,omitempty" xml:"DurationModule,omitempty" type:"Repeated"`
}

func (s DescribeRtcDurationDataResponseBodyDurationDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcDurationDataResponseBodyDurationDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeRtcDurationDataResponseBodyDurationDataPerInterval) SetDurationModule(v []*DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule) *DescribeRtcDurationDataResponseBodyDurationDataPerInterval {
	s.DurationModule = v
	return s
}

type DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule struct {
	ContentDuration *int64  `json:"ContentDuration,omitempty" xml:"ContentDuration,omitempty"`
	V720Duration    *int64  `json:"V720Duration,omitempty" xml:"V720Duration,omitempty"`
	V360Duration    *int64  `json:"V360Duration,omitempty" xml:"V360Duration,omitempty"`
	AudioDuration   *int64  `json:"AudioDuration,omitempty" xml:"AudioDuration,omitempty"`
	TimeStamp       *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	V1080Duration   *int64  `json:"V1080Duration,omitempty" xml:"V1080Duration,omitempty"`
	TotalDuration   *int64  `json:"TotalDuration,omitempty" xml:"TotalDuration,omitempty"`
}

func (s DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule) GoString() string {
	return s.String()
}

func (s *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule) SetContentDuration(v int64) *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule {
	s.ContentDuration = &v
	return s
}

func (s *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule) SetV720Duration(v int64) *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule {
	s.V720Duration = &v
	return s
}

func (s *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule) SetV360Duration(v int64) *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule {
	s.V360Duration = &v
	return s
}

func (s *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule) SetAudioDuration(v int64) *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule {
	s.AudioDuration = &v
	return s
}

func (s *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule) SetTimeStamp(v string) *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule) SetV1080Duration(v int64) *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule {
	s.V1080Duration = &v
	return s
}

func (s *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule) SetTotalDuration(v int64) *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule {
	s.TotalDuration = &v
	return s
}

type DescribeRtcDurationDataResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRtcDurationDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRtcDurationDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcDurationDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeRtcDurationDataResponse) SetHeaders(v map[string]*string) *DescribeRtcDurationDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeRtcDurationDataResponse) SetBody(v *DescribeRtcDurationDataResponseBody) *DescribeRtcDurationDataResponse {
	s.Body = v
	return s
}

type DescribeRtcPeakChannelCntDataRequest struct {
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog     *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ServiceArea *string `json:"ServiceArea,omitempty" xml:"ServiceArea,omitempty"`
	Interval    *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s DescribeRtcPeakChannelCntDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcPeakChannelCntDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeRtcPeakChannelCntDataRequest) SetOwnerId(v int64) *DescribeRtcPeakChannelCntDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRtcPeakChannelCntDataRequest) SetShowLog(v string) *DescribeRtcPeakChannelCntDataRequest {
	s.ShowLog = &v
	return s
}

func (s *DescribeRtcPeakChannelCntDataRequest) SetStartTime(v string) *DescribeRtcPeakChannelCntDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRtcPeakChannelCntDataRequest) SetEndTime(v string) *DescribeRtcPeakChannelCntDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRtcPeakChannelCntDataRequest) SetAppId(v string) *DescribeRtcPeakChannelCntDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeRtcPeakChannelCntDataRequest) SetServiceArea(v string) *DescribeRtcPeakChannelCntDataRequest {
	s.ServiceArea = &v
	return s
}

func (s *DescribeRtcPeakChannelCntDataRequest) SetInterval(v string) *DescribeRtcPeakChannelCntDataRequest {
	s.Interval = &v
	return s
}

type DescribeRtcPeakChannelCntDataResponseBody struct {
	PeakChannelCntDataPerInterval *DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerInterval `json:"PeakChannelCntDataPerInterval,omitempty" xml:"PeakChannelCntDataPerInterval,omitempty" type:"Struct"`
	RequestId                     *string                                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRtcPeakChannelCntDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcPeakChannelCntDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRtcPeakChannelCntDataResponseBody) SetPeakChannelCntDataPerInterval(v *DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerInterval) *DescribeRtcPeakChannelCntDataResponseBody {
	s.PeakChannelCntDataPerInterval = v
	return s
}

func (s *DescribeRtcPeakChannelCntDataResponseBody) SetRequestId(v string) *DescribeRtcPeakChannelCntDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerInterval struct {
	PeakChannelCntModule []*DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerIntervalPeakChannelCntModule `json:"PeakChannelCntModule,omitempty" xml:"PeakChannelCntModule,omitempty" type:"Repeated"`
}

func (s DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerInterval) SetPeakChannelCntModule(v []*DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerIntervalPeakChannelCntModule) *DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerInterval {
	s.PeakChannelCntModule = v
	return s
}

type DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerIntervalPeakChannelCntModule struct {
	ActiveChannelPeakTime *string `json:"ActiveChannelPeakTime,omitempty" xml:"ActiveChannelPeakTime,omitempty"`
	TimeStamp             *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	ActiveChannelPeak     *int64  `json:"ActiveChannelPeak,omitempty" xml:"ActiveChannelPeak,omitempty"`
}

func (s DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerIntervalPeakChannelCntModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerIntervalPeakChannelCntModule) GoString() string {
	return s.String()
}

func (s *DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerIntervalPeakChannelCntModule) SetActiveChannelPeakTime(v string) *DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerIntervalPeakChannelCntModule {
	s.ActiveChannelPeakTime = &v
	return s
}

func (s *DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerIntervalPeakChannelCntModule) SetTimeStamp(v string) *DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerIntervalPeakChannelCntModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerIntervalPeakChannelCntModule) SetActiveChannelPeak(v int64) *DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerIntervalPeakChannelCntModule {
	s.ActiveChannelPeak = &v
	return s
}

type DescribeRtcPeakChannelCntDataResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRtcPeakChannelCntDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRtcPeakChannelCntDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcPeakChannelCntDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeRtcPeakChannelCntDataResponse) SetHeaders(v map[string]*string) *DescribeRtcPeakChannelCntDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeRtcPeakChannelCntDataResponse) SetBody(v *DescribeRtcPeakChannelCntDataResponseBody) *DescribeRtcPeakChannelCntDataResponse {
	s.Body = v
	return s
}

type DescribeRtcPeakUserCntDataRequest struct {
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog     *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ServiceArea *string `json:"ServiceArea,omitempty" xml:"ServiceArea,omitempty"`
	Interval    *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s DescribeRtcPeakUserCntDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcPeakUserCntDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeRtcPeakUserCntDataRequest) SetOwnerId(v int64) *DescribeRtcPeakUserCntDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRtcPeakUserCntDataRequest) SetShowLog(v string) *DescribeRtcPeakUserCntDataRequest {
	s.ShowLog = &v
	return s
}

func (s *DescribeRtcPeakUserCntDataRequest) SetStartTime(v string) *DescribeRtcPeakUserCntDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRtcPeakUserCntDataRequest) SetEndTime(v string) *DescribeRtcPeakUserCntDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRtcPeakUserCntDataRequest) SetAppId(v string) *DescribeRtcPeakUserCntDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeRtcPeakUserCntDataRequest) SetServiceArea(v string) *DescribeRtcPeakUserCntDataRequest {
	s.ServiceArea = &v
	return s
}

func (s *DescribeRtcPeakUserCntDataRequest) SetInterval(v string) *DescribeRtcPeakUserCntDataRequest {
	s.Interval = &v
	return s
}

type DescribeRtcPeakUserCntDataResponseBody struct {
	PeakUserCntDataPerInterval *DescribeRtcPeakUserCntDataResponseBodyPeakUserCntDataPerInterval `json:"PeakUserCntDataPerInterval,omitempty" xml:"PeakUserCntDataPerInterval,omitempty" type:"Struct"`
	RequestId                  *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRtcPeakUserCntDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcPeakUserCntDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRtcPeakUserCntDataResponseBody) SetPeakUserCntDataPerInterval(v *DescribeRtcPeakUserCntDataResponseBodyPeakUserCntDataPerInterval) *DescribeRtcPeakUserCntDataResponseBody {
	s.PeakUserCntDataPerInterval = v
	return s
}

func (s *DescribeRtcPeakUserCntDataResponseBody) SetRequestId(v string) *DescribeRtcPeakUserCntDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRtcPeakUserCntDataResponseBodyPeakUserCntDataPerInterval struct {
	PeakUserCntModule []*DescribeRtcPeakUserCntDataResponseBodyPeakUserCntDataPerIntervalPeakUserCntModule `json:"PeakUserCntModule,omitempty" xml:"PeakUserCntModule,omitempty" type:"Repeated"`
}

func (s DescribeRtcPeakUserCntDataResponseBodyPeakUserCntDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcPeakUserCntDataResponseBodyPeakUserCntDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeRtcPeakUserCntDataResponseBodyPeakUserCntDataPerInterval) SetPeakUserCntModule(v []*DescribeRtcPeakUserCntDataResponseBodyPeakUserCntDataPerIntervalPeakUserCntModule) *DescribeRtcPeakUserCntDataResponseBodyPeakUserCntDataPerInterval {
	s.PeakUserCntModule = v
	return s
}

type DescribeRtcPeakUserCntDataResponseBodyPeakUserCntDataPerIntervalPeakUserCntModule struct {
	ActiveUserPeakTime *string `json:"ActiveUserPeakTime,omitempty" xml:"ActiveUserPeakTime,omitempty"`
	TimeStamp          *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	ActiveUserPeak     *int64  `json:"ActiveUserPeak,omitempty" xml:"ActiveUserPeak,omitempty"`
}

func (s DescribeRtcPeakUserCntDataResponseBodyPeakUserCntDataPerIntervalPeakUserCntModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcPeakUserCntDataResponseBodyPeakUserCntDataPerIntervalPeakUserCntModule) GoString() string {
	return s.String()
}

func (s *DescribeRtcPeakUserCntDataResponseBodyPeakUserCntDataPerIntervalPeakUserCntModule) SetActiveUserPeakTime(v string) *DescribeRtcPeakUserCntDataResponseBodyPeakUserCntDataPerIntervalPeakUserCntModule {
	s.ActiveUserPeakTime = &v
	return s
}

func (s *DescribeRtcPeakUserCntDataResponseBodyPeakUserCntDataPerIntervalPeakUserCntModule) SetTimeStamp(v string) *DescribeRtcPeakUserCntDataResponseBodyPeakUserCntDataPerIntervalPeakUserCntModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeRtcPeakUserCntDataResponseBodyPeakUserCntDataPerIntervalPeakUserCntModule) SetActiveUserPeak(v int64) *DescribeRtcPeakUserCntDataResponseBodyPeakUserCntDataPerIntervalPeakUserCntModule {
	s.ActiveUserPeak = &v
	return s
}

type DescribeRtcPeakUserCntDataResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRtcPeakUserCntDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRtcPeakUserCntDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcPeakUserCntDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeRtcPeakUserCntDataResponse) SetHeaders(v map[string]*string) *DescribeRtcPeakUserCntDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeRtcPeakUserCntDataResponse) SetBody(v *DescribeRtcPeakUserCntDataResponseBody) *DescribeRtcPeakUserCntDataResponse {
	s.Body = v
	return s
}

type DescribeRtcScaleRequest struct {
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog   *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeRtcScaleRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcScaleRequest) GoString() string {
	return s.String()
}

func (s *DescribeRtcScaleRequest) SetOwnerId(v int64) *DescribeRtcScaleRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRtcScaleRequest) SetShowLog(v string) *DescribeRtcScaleRequest {
	s.ShowLog = &v
	return s
}

func (s *DescribeRtcScaleRequest) SetAppId(v string) *DescribeRtcScaleRequest {
	s.AppId = &v
	return s
}

func (s *DescribeRtcScaleRequest) SetStartTime(v string) *DescribeRtcScaleRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRtcScaleRequest) SetEndTime(v string) *DescribeRtcScaleRequest {
	s.EndTime = &v
	return s
}

type DescribeRtcScaleResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Scale     []*DescribeRtcScaleResponseBodyScale `json:"Scale,omitempty" xml:"Scale,omitempty" type:"Repeated"`
}

func (s DescribeRtcScaleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcScaleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRtcScaleResponseBody) SetRequestId(v string) *DescribeRtcScaleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRtcScaleResponseBody) SetScale(v []*DescribeRtcScaleResponseBodyScale) *DescribeRtcScaleResponseBody {
	s.Scale = v
	return s
}

type DescribeRtcScaleResponseBodyScale struct {
	SessionCount  *int64  `json:"SessionCount,omitempty" xml:"SessionCount,omitempty"`
	Time          *string `json:"Time,omitempty" xml:"Time,omitempty"`
	ChannelCount  *int64  `json:"ChannelCount,omitempty" xml:"ChannelCount,omitempty"`
	AudioDuration *int64  `json:"AudioDuration,omitempty" xml:"AudioDuration,omitempty"`
	UserCount     *int64  `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
	VideoDuration *int64  `json:"VideoDuration,omitempty" xml:"VideoDuration,omitempty"`
}

func (s DescribeRtcScaleResponseBodyScale) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcScaleResponseBodyScale) GoString() string {
	return s.String()
}

func (s *DescribeRtcScaleResponseBodyScale) SetSessionCount(v int64) *DescribeRtcScaleResponseBodyScale {
	s.SessionCount = &v
	return s
}

func (s *DescribeRtcScaleResponseBodyScale) SetTime(v string) *DescribeRtcScaleResponseBodyScale {
	s.Time = &v
	return s
}

func (s *DescribeRtcScaleResponseBodyScale) SetChannelCount(v int64) *DescribeRtcScaleResponseBodyScale {
	s.ChannelCount = &v
	return s
}

func (s *DescribeRtcScaleResponseBodyScale) SetAudioDuration(v int64) *DescribeRtcScaleResponseBodyScale {
	s.AudioDuration = &v
	return s
}

func (s *DescribeRtcScaleResponseBodyScale) SetUserCount(v int64) *DescribeRtcScaleResponseBodyScale {
	s.UserCount = &v
	return s
}

func (s *DescribeRtcScaleResponseBodyScale) SetVideoDuration(v int64) *DescribeRtcScaleResponseBodyScale {
	s.VideoDuration = &v
	return s
}

type DescribeRtcScaleResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRtcScaleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRtcScaleResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcScaleResponse) GoString() string {
	return s.String()
}

func (s *DescribeRtcScaleResponse) SetHeaders(v map[string]*string) *DescribeRtcScaleResponse {
	s.Headers = v
	return s
}

func (s *DescribeRtcScaleResponse) SetBody(v *DescribeRtcScaleResponseBody) *DescribeRtcScaleResponse {
	s.Body = v
	return s
}

type DescribeRtcScaleDetailRequest struct {
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog   *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeRtcScaleDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcScaleDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeRtcScaleDetailRequest) SetOwnerId(v int64) *DescribeRtcScaleDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRtcScaleDetailRequest) SetShowLog(v string) *DescribeRtcScaleDetailRequest {
	s.ShowLog = &v
	return s
}

func (s *DescribeRtcScaleDetailRequest) SetAppId(v string) *DescribeRtcScaleDetailRequest {
	s.AppId = &v
	return s
}

func (s *DescribeRtcScaleDetailRequest) SetStartTime(v string) *DescribeRtcScaleDetailRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRtcScaleDetailRequest) SetEndTime(v string) *DescribeRtcScaleDetailRequest {
	s.EndTime = &v
	return s
}

type DescribeRtcScaleDetailResponseBody struct {
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Scale     []*DescribeRtcScaleDetailResponseBodyScale `json:"Scale,omitempty" xml:"Scale,omitempty" type:"Repeated"`
}

func (s DescribeRtcScaleDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcScaleDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRtcScaleDetailResponseBody) SetRequestId(v string) *DescribeRtcScaleDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRtcScaleDetailResponseBody) SetScale(v []*DescribeRtcScaleDetailResponseBodyScale) *DescribeRtcScaleDetailResponseBody {
	s.Scale = v
	return s
}

type DescribeRtcScaleDetailResponseBodyScale struct {
	CC *int64  `json:"CC,omitempty" xml:"CC,omitempty"`
	TS *string `json:"TS,omitempty" xml:"TS,omitempty"`
	UC *int64  `json:"UC,omitempty" xml:"UC,omitempty"`
}

func (s DescribeRtcScaleDetailResponseBodyScale) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcScaleDetailResponseBodyScale) GoString() string {
	return s.String()
}

func (s *DescribeRtcScaleDetailResponseBodyScale) SetCC(v int64) *DescribeRtcScaleDetailResponseBodyScale {
	s.CC = &v
	return s
}

func (s *DescribeRtcScaleDetailResponseBodyScale) SetTS(v string) *DescribeRtcScaleDetailResponseBodyScale {
	s.TS = &v
	return s
}

func (s *DescribeRtcScaleDetailResponseBodyScale) SetUC(v int64) *DescribeRtcScaleDetailResponseBodyScale {
	s.UC = &v
	return s
}

type DescribeRtcScaleDetailResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRtcScaleDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRtcScaleDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcScaleDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeRtcScaleDetailResponse) SetHeaders(v map[string]*string) *DescribeRtcScaleDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeRtcScaleDetailResponse) SetBody(v *DescribeRtcScaleDetailResponseBody) *DescribeRtcScaleDetailResponse {
	s.Body = v
	return s
}

type DescribeRtcUserCntDataRequest struct {
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog     *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime     *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ServiceArea *string `json:"ServiceArea,omitempty" xml:"ServiceArea,omitempty"`
	Interval    *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
}

func (s DescribeRtcUserCntDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcUserCntDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeRtcUserCntDataRequest) SetOwnerId(v int64) *DescribeRtcUserCntDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRtcUserCntDataRequest) SetShowLog(v string) *DescribeRtcUserCntDataRequest {
	s.ShowLog = &v
	return s
}

func (s *DescribeRtcUserCntDataRequest) SetStartTime(v string) *DescribeRtcUserCntDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRtcUserCntDataRequest) SetEndTime(v string) *DescribeRtcUserCntDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRtcUserCntDataRequest) SetAppId(v string) *DescribeRtcUserCntDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeRtcUserCntDataRequest) SetServiceArea(v string) *DescribeRtcUserCntDataRequest {
	s.ServiceArea = &v
	return s
}

func (s *DescribeRtcUserCntDataRequest) SetInterval(v string) *DescribeRtcUserCntDataRequest {
	s.Interval = &v
	return s
}

type DescribeRtcUserCntDataResponseBody struct {
	RequestId              *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserCntDataPerInterval *DescribeRtcUserCntDataResponseBodyUserCntDataPerInterval `json:"UserCntDataPerInterval,omitempty" xml:"UserCntDataPerInterval,omitempty" type:"Struct"`
}

func (s DescribeRtcUserCntDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcUserCntDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRtcUserCntDataResponseBody) SetRequestId(v string) *DescribeRtcUserCntDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRtcUserCntDataResponseBody) SetUserCntDataPerInterval(v *DescribeRtcUserCntDataResponseBodyUserCntDataPerInterval) *DescribeRtcUserCntDataResponseBody {
	s.UserCntDataPerInterval = v
	return s
}

type DescribeRtcUserCntDataResponseBodyUserCntDataPerInterval struct {
	UserCntModule []*DescribeRtcUserCntDataResponseBodyUserCntDataPerIntervalUserCntModule `json:"UserCntModule,omitempty" xml:"UserCntModule,omitempty" type:"Repeated"`
}

func (s DescribeRtcUserCntDataResponseBodyUserCntDataPerInterval) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcUserCntDataResponseBodyUserCntDataPerInterval) GoString() string {
	return s.String()
}

func (s *DescribeRtcUserCntDataResponseBodyUserCntDataPerInterval) SetUserCntModule(v []*DescribeRtcUserCntDataResponseBodyUserCntDataPerIntervalUserCntModule) *DescribeRtcUserCntDataResponseBodyUserCntDataPerInterval {
	s.UserCntModule = v
	return s
}

type DescribeRtcUserCntDataResponseBodyUserCntDataPerIntervalUserCntModule struct {
	ActiveUserCnt *int64  `json:"ActiveUserCnt,omitempty" xml:"ActiveUserCnt,omitempty"`
	TimeStamp     *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeRtcUserCntDataResponseBodyUserCntDataPerIntervalUserCntModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcUserCntDataResponseBodyUserCntDataPerIntervalUserCntModule) GoString() string {
	return s.String()
}

func (s *DescribeRtcUserCntDataResponseBodyUserCntDataPerIntervalUserCntModule) SetActiveUserCnt(v int64) *DescribeRtcUserCntDataResponseBodyUserCntDataPerIntervalUserCntModule {
	s.ActiveUserCnt = &v
	return s
}

func (s *DescribeRtcUserCntDataResponseBodyUserCntDataPerIntervalUserCntModule) SetTimeStamp(v string) *DescribeRtcUserCntDataResponseBodyUserCntDataPerIntervalUserCntModule {
	s.TimeStamp = &v
	return s
}

type DescribeRtcUserCntDataResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRtcUserCntDataResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRtcUserCntDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcUserCntDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeRtcUserCntDataResponse) SetHeaders(v map[string]*string) *DescribeRtcUserCntDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeRtcUserCntDataResponse) SetBody(v *DescribeRtcUserCntDataResponseBody) *DescribeRtcUserCntDataResponse {
	s.Body = v
	return s
}

type DescribeRtcUserEventsRequest struct {
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog   *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Uid       *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeRtcUserEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcUserEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRtcUserEventsRequest) SetOwnerId(v int64) *DescribeRtcUserEventsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRtcUserEventsRequest) SetShowLog(v string) *DescribeRtcUserEventsRequest {
	s.ShowLog = &v
	return s
}

func (s *DescribeRtcUserEventsRequest) SetAppId(v string) *DescribeRtcUserEventsRequest {
	s.AppId = &v
	return s
}

func (s *DescribeRtcUserEventsRequest) SetUid(v string) *DescribeRtcUserEventsRequest {
	s.Uid = &v
	return s
}

func (s *DescribeRtcUserEventsRequest) SetChannelId(v string) *DescribeRtcUserEventsRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeRtcUserEventsRequest) SetStartTime(v string) *DescribeRtcUserEventsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRtcUserEventsRequest) SetEndTime(v string) *DescribeRtcUserEventsRequest {
	s.EndTime = &v
	return s
}

type DescribeRtcUserEventsResponseBody struct {
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Events    []*DescribeRtcUserEventsResponseBodyEvents `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
}

func (s DescribeRtcUserEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcUserEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRtcUserEventsResponseBody) SetRequestId(v string) *DescribeRtcUserEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRtcUserEventsResponseBody) SetEvents(v []*DescribeRtcUserEventsResponseBodyEvents) *DescribeRtcUserEventsResponseBody {
	s.Events = v
	return s
}

type DescribeRtcUserEventsResponseBodyEvents struct {
	EventId   *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	EventTime *int64  `json:"EventTime,omitempty" xml:"EventTime,omitempty"`
	Category  *string `json:"Category,omitempty" xml:"Category,omitempty"`
}

func (s DescribeRtcUserEventsResponseBodyEvents) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcUserEventsResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *DescribeRtcUserEventsResponseBodyEvents) SetEventId(v string) *DescribeRtcUserEventsResponseBodyEvents {
	s.EventId = &v
	return s
}

func (s *DescribeRtcUserEventsResponseBodyEvents) SetEventTime(v int64) *DescribeRtcUserEventsResponseBodyEvents {
	s.EventTime = &v
	return s
}

func (s *DescribeRtcUserEventsResponseBodyEvents) SetCategory(v string) *DescribeRtcUserEventsResponseBodyEvents {
	s.Category = &v
	return s
}

type DescribeRtcUserEventsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRtcUserEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRtcUserEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcUserEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRtcUserEventsResponse) SetHeaders(v map[string]*string) *DescribeRtcUserEventsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRtcUserEventsResponse) SetBody(v *DescribeRtcUserEventsResponseBody) *DescribeRtcUserEventsResponse {
	s.Body = v
	return s
}

type DescribeRtcUserListRequest struct {
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog   *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	PubUser   *string `json:"PubUser,omitempty" xml:"PubUser,omitempty"`
	SubUser   *string `json:"SubUser,omitempty" xml:"SubUser,omitempty"`
}

func (s DescribeRtcUserListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcUserListRequest) GoString() string {
	return s.String()
}

func (s *DescribeRtcUserListRequest) SetOwnerId(v int64) *DescribeRtcUserListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRtcUserListRequest) SetShowLog(v string) *DescribeRtcUserListRequest {
	s.ShowLog = &v
	return s
}

func (s *DescribeRtcUserListRequest) SetStartTime(v string) *DescribeRtcUserListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRtcUserListRequest) SetEndTime(v string) *DescribeRtcUserListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRtcUserListRequest) SetAppId(v string) *DescribeRtcUserListRequest {
	s.AppId = &v
	return s
}

func (s *DescribeRtcUserListRequest) SetChannelId(v string) *DescribeRtcUserListRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeRtcUserListRequest) SetPubUser(v string) *DescribeRtcUserListRequest {
	s.PubUser = &v
	return s
}

func (s *DescribeRtcUserListRequest) SetSubUser(v string) *DescribeRtcUserListRequest {
	s.SubUser = &v
	return s
}

type DescribeRtcUserListResponseBody struct {
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserList  *DescribeRtcUserListResponseBodyUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Struct"`
}

func (s DescribeRtcUserListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcUserListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRtcUserListResponseBody) SetRequestId(v string) *DescribeRtcUserListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRtcUserListResponseBody) SetUserList(v *DescribeRtcUserListResponseBodyUserList) *DescribeRtcUserListResponseBody {
	s.UserList = v
	return s
}

type DescribeRtcUserListResponseBodyUserList struct {
	UserList []*DescribeRtcUserListResponseBodyUserListUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
}

func (s DescribeRtcUserListResponseBodyUserList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcUserListResponseBodyUserList) GoString() string {
	return s.String()
}

func (s *DescribeRtcUserListResponseBodyUserList) SetUserList(v []*DescribeRtcUserListResponseBodyUserListUserList) *DescribeRtcUserListResponseBodyUserList {
	s.UserList = v
	return s
}

type DescribeRtcUserListResponseBodyUserListUserList struct {
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s DescribeRtcUserListResponseBodyUserListUserList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcUserListResponseBodyUserListUserList) GoString() string {
	return s.String()
}

func (s *DescribeRtcUserListResponseBodyUserListUserList) SetUser(v string) *DescribeRtcUserListResponseBodyUserListUserList {
	s.User = &v
	return s
}

type DescribeRtcUserListResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeRtcUserListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRtcUserListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcUserListResponse) GoString() string {
	return s.String()
}

func (s *DescribeRtcUserListResponse) SetHeaders(v map[string]*string) *DescribeRtcUserListResponse {
	s.Headers = v
	return s
}

func (s *DescribeRtcUserListResponse) SetBody(v *DescribeRtcUserListResponseBody) *DescribeRtcUserListResponse {
	s.Body = v
	return s
}

type DescribeUserInfoInChannelRequest struct {
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog   *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	UserId    *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeUserInfoInChannelRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserInfoInChannelRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserInfoInChannelRequest) SetOwnerId(v int64) *DescribeUserInfoInChannelRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUserInfoInChannelRequest) SetShowLog(v string) *DescribeUserInfoInChannelRequest {
	s.ShowLog = &v
	return s
}

func (s *DescribeUserInfoInChannelRequest) SetAppId(v string) *DescribeUserInfoInChannelRequest {
	s.AppId = &v
	return s
}

func (s *DescribeUserInfoInChannelRequest) SetChannelId(v string) *DescribeUserInfoInChannelRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeUserInfoInChannelRequest) SetUserId(v string) *DescribeUserInfoInChannelRequest {
	s.UserId = &v
	return s
}

type DescribeUserInfoInChannelResponseBody struct {
	RequestId      *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	IsInChannel    *bool                                            `json:"IsInChannel,omitempty" xml:"IsInChannel,omitempty"`
	Timestamp      *int32                                           `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	IsChannelExist *bool                                            `json:"IsChannelExist,omitempty" xml:"IsChannelExist,omitempty"`
	Property       []*DescribeUserInfoInChannelResponseBodyProperty `json:"Property,omitempty" xml:"Property,omitempty" type:"Repeated"`
}

func (s DescribeUserInfoInChannelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserInfoInChannelResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserInfoInChannelResponseBody) SetRequestId(v string) *DescribeUserInfoInChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserInfoInChannelResponseBody) SetIsInChannel(v bool) *DescribeUserInfoInChannelResponseBody {
	s.IsInChannel = &v
	return s
}

func (s *DescribeUserInfoInChannelResponseBody) SetTimestamp(v int32) *DescribeUserInfoInChannelResponseBody {
	s.Timestamp = &v
	return s
}

func (s *DescribeUserInfoInChannelResponseBody) SetIsChannelExist(v bool) *DescribeUserInfoInChannelResponseBody {
	s.IsChannelExist = &v
	return s
}

func (s *DescribeUserInfoInChannelResponseBody) SetProperty(v []*DescribeUserInfoInChannelResponseBodyProperty) *DescribeUserInfoInChannelResponseBody {
	s.Property = v
	return s
}

type DescribeUserInfoInChannelResponseBodyProperty struct {
	Session *string `json:"Session,omitempty" xml:"Session,omitempty"`
	Join    *int32  `json:"Join,omitempty" xml:"Join,omitempty"`
	Role    *int32  `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s DescribeUserInfoInChannelResponseBodyProperty) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserInfoInChannelResponseBodyProperty) GoString() string {
	return s.String()
}

func (s *DescribeUserInfoInChannelResponseBodyProperty) SetSession(v string) *DescribeUserInfoInChannelResponseBodyProperty {
	s.Session = &v
	return s
}

func (s *DescribeUserInfoInChannelResponseBodyProperty) SetJoin(v int32) *DescribeUserInfoInChannelResponseBodyProperty {
	s.Join = &v
	return s
}

func (s *DescribeUserInfoInChannelResponseBodyProperty) SetRole(v int32) *DescribeUserInfoInChannelResponseBodyProperty {
	s.Role = &v
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

type DisableMPURuleRequest struct {
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RuleId  *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DisableMPURuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DisableMPURuleRequest) GoString() string {
	return s.String()
}

func (s *DisableMPURuleRequest) SetOwnerId(v int64) *DisableMPURuleRequest {
	s.OwnerId = &v
	return s
}

func (s *DisableMPURuleRequest) SetShowLog(v string) *DisableMPURuleRequest {
	s.ShowLog = &v
	return s
}

func (s *DisableMPURuleRequest) SetAppId(v string) *DisableMPURuleRequest {
	s.AppId = &v
	return s
}

func (s *DisableMPURuleRequest) SetRuleId(v int64) *DisableMPURuleRequest {
	s.RuleId = &v
	return s
}

type DisableMPURuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableMPURuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableMPURuleResponseBody) GoString() string {
	return s.String()
}

func (s *DisableMPURuleResponseBody) SetRequestId(v string) *DisableMPURuleResponseBody {
	s.RequestId = &v
	return s
}

type DisableMPURuleResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DisableMPURuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableMPURuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableMPURuleResponse) GoString() string {
	return s.String()
}

func (s *DisableMPURuleResponse) SetHeaders(v map[string]*string) *DisableMPURuleResponse {
	s.Headers = v
	return s
}

func (s *DisableMPURuleResponse) SetBody(v *DisableMPURuleResponseBody) *DisableMPURuleResponse {
	s.Body = v
	return s
}

type EnableMPURuleRequest struct {
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RuleId  *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s EnableMPURuleRequest) String() string {
	return tea.Prettify(s)
}

func (s EnableMPURuleRequest) GoString() string {
	return s.String()
}

func (s *EnableMPURuleRequest) SetOwnerId(v int64) *EnableMPURuleRequest {
	s.OwnerId = &v
	return s
}

func (s *EnableMPURuleRequest) SetShowLog(v string) *EnableMPURuleRequest {
	s.ShowLog = &v
	return s
}

func (s *EnableMPURuleRequest) SetAppId(v string) *EnableMPURuleRequest {
	s.AppId = &v
	return s
}

func (s *EnableMPURuleRequest) SetRuleId(v int64) *EnableMPURuleRequest {
	s.RuleId = &v
	return s
}

type EnableMPURuleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s EnableMPURuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableMPURuleResponseBody) GoString() string {
	return s.String()
}

func (s *EnableMPURuleResponseBody) SetRequestId(v string) *EnableMPURuleResponseBody {
	s.RequestId = &v
	return s
}

type EnableMPURuleResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EnableMPURuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableMPURuleResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableMPURuleResponse) GoString() string {
	return s.String()
}

func (s *EnableMPURuleResponse) SetHeaders(v map[string]*string) *EnableMPURuleResponse {
	s.Headers = v
	return s
}

func (s *EnableMPURuleResponse) SetBody(v *EnableMPURuleResponseBody) *EnableMPURuleResponse {
	s.Body = v
	return s
}

type GetMPUTaskStatusRequest struct {
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	TaskId  *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetMPUTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetMPUTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *GetMPUTaskStatusRequest) SetOwnerId(v int64) *GetMPUTaskStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *GetMPUTaskStatusRequest) SetShowLog(v string) *GetMPUTaskStatusRequest {
	s.ShowLog = &v
	return s
}

func (s *GetMPUTaskStatusRequest) SetAppId(v string) *GetMPUTaskStatusRequest {
	s.AppId = &v
	return s
}

func (s *GetMPUTaskStatusRequest) SetTaskId(v string) *GetMPUTaskStatusRequest {
	s.TaskId = &v
	return s
}

type GetMPUTaskStatusResponseBody struct {
	Status    *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMPUTaskStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetMPUTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetMPUTaskStatusResponseBody) SetStatus(v int32) *GetMPUTaskStatusResponseBody {
	s.Status = &v
	return s
}

func (s *GetMPUTaskStatusResponseBody) SetRequestId(v string) *GetMPUTaskStatusResponseBody {
	s.RequestId = &v
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

type ModifyAppRequest struct {
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s ModifyAppRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppRequest) SetOwnerId(v int64) *ModifyAppRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAppRequest) SetShowLog(v string) *ModifyAppRequest {
	s.ShowLog = &v
	return s
}

func (s *ModifyAppRequest) SetAppId(v string) *ModifyAppRequest {
	s.AppId = &v
	return s
}

func (s *ModifyAppRequest) SetAppName(v string) *ModifyAppRequest {
	s.AppName = &v
	return s
}

type ModifyAppResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAppResponseBody) SetRequestId(v string) *ModifyAppResponseBody {
	s.RequestId = &v
	return s
}

type ModifyAppResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAppResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppResponse) GoString() string {
	return s.String()
}

func (s *ModifyAppResponse) SetHeaders(v map[string]*string) *ModifyAppResponse {
	s.Headers = v
	return s
}

func (s *ModifyAppResponse) SetBody(v *ModifyAppResponseBody) *ModifyAppResponse {
	s.Body = v
	return s
}

type ModifyConferenceRequest struct {
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog        *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ConferenceId   *string `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	ConferenceName *string `json:"ConferenceName,omitempty" xml:"ConferenceName,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
	RemindNotice   *int32  `json:"RemindNotice,omitempty" xml:"RemindNotice,omitempty"`
}

func (s ModifyConferenceRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyConferenceRequest) GoString() string {
	return s.String()
}

func (s *ModifyConferenceRequest) SetOwnerId(v int64) *ModifyConferenceRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyConferenceRequest) SetShowLog(v string) *ModifyConferenceRequest {
	s.ShowLog = &v
	return s
}

func (s *ModifyConferenceRequest) SetAppId(v string) *ModifyConferenceRequest {
	s.AppId = &v
	return s
}

func (s *ModifyConferenceRequest) SetConferenceId(v string) *ModifyConferenceRequest {
	s.ConferenceId = &v
	return s
}

func (s *ModifyConferenceRequest) SetConferenceName(v string) *ModifyConferenceRequest {
	s.ConferenceName = &v
	return s
}

func (s *ModifyConferenceRequest) SetStartTime(v string) *ModifyConferenceRequest {
	s.StartTime = &v
	return s
}

func (s *ModifyConferenceRequest) SetType(v string) *ModifyConferenceRequest {
	s.Type = &v
	return s
}

func (s *ModifyConferenceRequest) SetRemindNotice(v int32) *ModifyConferenceRequest {
	s.RemindNotice = &v
	return s
}

type ModifyConferenceResponseBody struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ConferenceId *string `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
}

func (s ModifyConferenceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyConferenceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyConferenceResponseBody) SetRequestId(v string) *ModifyConferenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyConferenceResponseBody) SetConferenceId(v string) *ModifyConferenceResponseBody {
	s.ConferenceId = &v
	return s
}

type ModifyConferenceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyConferenceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyConferenceResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyConferenceResponse) GoString() string {
	return s.String()
}

func (s *ModifyConferenceResponse) SetHeaders(v map[string]*string) *ModifyConferenceResponse {
	s.Headers = v
	return s
}

func (s *ModifyConferenceResponse) SetBody(v *ModifyConferenceResponseBody) *ModifyConferenceResponse {
	s.Body = v
	return s
}

type ModifyMPULayoutRequest struct {
	OwnerId       *int64                         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog       *string                        `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId         *string                        `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Name          *string                        `json:"Name,omitempty" xml:"Name,omitempty"`
	LayoutId      *int64                         `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	AudioMixCount *int32                         `json:"AudioMixCount,omitempty" xml:"AudioMixCount,omitempty"`
	Panes         []*ModifyMPULayoutRequestPanes `json:"Panes,omitempty" xml:"Panes,omitempty" type:"Repeated"`
}

func (s ModifyMPULayoutRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyMPULayoutRequest) GoString() string {
	return s.String()
}

func (s *ModifyMPULayoutRequest) SetOwnerId(v int64) *ModifyMPULayoutRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyMPULayoutRequest) SetShowLog(v string) *ModifyMPULayoutRequest {
	s.ShowLog = &v
	return s
}

func (s *ModifyMPULayoutRequest) SetAppId(v string) *ModifyMPULayoutRequest {
	s.AppId = &v
	return s
}

func (s *ModifyMPULayoutRequest) SetName(v string) *ModifyMPULayoutRequest {
	s.Name = &v
	return s
}

func (s *ModifyMPULayoutRequest) SetLayoutId(v int64) *ModifyMPULayoutRequest {
	s.LayoutId = &v
	return s
}

func (s *ModifyMPULayoutRequest) SetAudioMixCount(v int32) *ModifyMPULayoutRequest {
	s.AudioMixCount = &v
	return s
}

func (s *ModifyMPULayoutRequest) SetPanes(v []*ModifyMPULayoutRequestPanes) *ModifyMPULayoutRequest {
	s.Panes = v
	return s
}

type ModifyMPULayoutRequestPanes struct {
	MajorPane *int32   `json:"MajorPane,omitempty" xml:"MajorPane,omitempty"`
	Width     *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height    *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Y         *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	PaneId    *int32   `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
	ZOrder    *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
	X         *float32 `json:"X,omitempty" xml:"X,omitempty"`
}

func (s ModifyMPULayoutRequestPanes) String() string {
	return tea.Prettify(s)
}

func (s ModifyMPULayoutRequestPanes) GoString() string {
	return s.String()
}

func (s *ModifyMPULayoutRequestPanes) SetMajorPane(v int32) *ModifyMPULayoutRequestPanes {
	s.MajorPane = &v
	return s
}

func (s *ModifyMPULayoutRequestPanes) SetWidth(v float32) *ModifyMPULayoutRequestPanes {
	s.Width = &v
	return s
}

func (s *ModifyMPULayoutRequestPanes) SetHeight(v float32) *ModifyMPULayoutRequestPanes {
	s.Height = &v
	return s
}

func (s *ModifyMPULayoutRequestPanes) SetY(v float32) *ModifyMPULayoutRequestPanes {
	s.Y = &v
	return s
}

func (s *ModifyMPULayoutRequestPanes) SetPaneId(v int32) *ModifyMPULayoutRequestPanes {
	s.PaneId = &v
	return s
}

func (s *ModifyMPULayoutRequestPanes) SetZOrder(v int32) *ModifyMPULayoutRequestPanes {
	s.ZOrder = &v
	return s
}

func (s *ModifyMPULayoutRequestPanes) SetX(v float32) *ModifyMPULayoutRequestPanes {
	s.X = &v
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

type MuteAudioRequest struct {
	OwnerId        *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog        *string   `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId          *string   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ConferenceId   *string   `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	ParticipantIds []*string `json:"ParticipantIds,omitempty" xml:"ParticipantIds,omitempty" type:"Repeated"`
}

func (s MuteAudioRequest) String() string {
	return tea.Prettify(s)
}

func (s MuteAudioRequest) GoString() string {
	return s.String()
}

func (s *MuteAudioRequest) SetOwnerId(v int64) *MuteAudioRequest {
	s.OwnerId = &v
	return s
}

func (s *MuteAudioRequest) SetShowLog(v string) *MuteAudioRequest {
	s.ShowLog = &v
	return s
}

func (s *MuteAudioRequest) SetAppId(v string) *MuteAudioRequest {
	s.AppId = &v
	return s
}

func (s *MuteAudioRequest) SetConferenceId(v string) *MuteAudioRequest {
	s.ConferenceId = &v
	return s
}

func (s *MuteAudioRequest) SetParticipantIds(v []*string) *MuteAudioRequest {
	s.ParticipantIds = v
	return s
}

type MuteAudioResponseBody struct {
	RequestId    *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ConferenceId *string                            `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	Participants *MuteAudioResponseBodyParticipants `json:"Participants,omitempty" xml:"Participants,omitempty" type:"Struct"`
}

func (s MuteAudioResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MuteAudioResponseBody) GoString() string {
	return s.String()
}

func (s *MuteAudioResponseBody) SetRequestId(v string) *MuteAudioResponseBody {
	s.RequestId = &v
	return s
}

func (s *MuteAudioResponseBody) SetConferenceId(v string) *MuteAudioResponseBody {
	s.ConferenceId = &v
	return s
}

func (s *MuteAudioResponseBody) SetParticipants(v *MuteAudioResponseBodyParticipants) *MuteAudioResponseBody {
	s.Participants = v
	return s
}

type MuteAudioResponseBodyParticipants struct {
	Participant []*MuteAudioResponseBodyParticipantsParticipant `json:"Participant,omitempty" xml:"Participant,omitempty" type:"Repeated"`
}

func (s MuteAudioResponseBodyParticipants) String() string {
	return tea.Prettify(s)
}

func (s MuteAudioResponseBodyParticipants) GoString() string {
	return s.String()
}

func (s *MuteAudioResponseBodyParticipants) SetParticipant(v []*MuteAudioResponseBodyParticipantsParticipant) *MuteAudioResponseBodyParticipants {
	s.Participant = v
	return s
}

type MuteAudioResponseBodyParticipantsParticipant struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s MuteAudioResponseBodyParticipantsParticipant) String() string {
	return tea.Prettify(s)
}

func (s MuteAudioResponseBodyParticipantsParticipant) GoString() string {
	return s.String()
}

func (s *MuteAudioResponseBodyParticipantsParticipant) SetCode(v string) *MuteAudioResponseBodyParticipantsParticipant {
	s.Code = &v
	return s
}

func (s *MuteAudioResponseBodyParticipantsParticipant) SetMessage(v string) *MuteAudioResponseBodyParticipantsParticipant {
	s.Message = &v
	return s
}

func (s *MuteAudioResponseBodyParticipantsParticipant) SetId(v string) *MuteAudioResponseBodyParticipantsParticipant {
	s.Id = &v
	return s
}

type MuteAudioResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MuteAudioResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MuteAudioResponse) String() string {
	return tea.Prettify(s)
}

func (s MuteAudioResponse) GoString() string {
	return s.String()
}

func (s *MuteAudioResponse) SetHeaders(v map[string]*string) *MuteAudioResponse {
	s.Headers = v
	return s
}

func (s *MuteAudioResponse) SetBody(v *MuteAudioResponseBody) *MuteAudioResponse {
	s.Body = v
	return s
}

type MuteAudioAllRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog       *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ConferenceId  *string `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	ParticipantId *string `json:"ParticipantId,omitempty" xml:"ParticipantId,omitempty"`
}

func (s MuteAudioAllRequest) String() string {
	return tea.Prettify(s)
}

func (s MuteAudioAllRequest) GoString() string {
	return s.String()
}

func (s *MuteAudioAllRequest) SetOwnerId(v int64) *MuteAudioAllRequest {
	s.OwnerId = &v
	return s
}

func (s *MuteAudioAllRequest) SetShowLog(v string) *MuteAudioAllRequest {
	s.ShowLog = &v
	return s
}

func (s *MuteAudioAllRequest) SetAppId(v string) *MuteAudioAllRequest {
	s.AppId = &v
	return s
}

func (s *MuteAudioAllRequest) SetConferenceId(v string) *MuteAudioAllRequest {
	s.ConferenceId = &v
	return s
}

func (s *MuteAudioAllRequest) SetParticipantId(v string) *MuteAudioAllRequest {
	s.ParticipantId = &v
	return s
}

type MuteAudioAllResponseBody struct {
	RequestId    *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ConferenceId *string                               `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	Participants *MuteAudioAllResponseBodyParticipants `json:"Participants,omitempty" xml:"Participants,omitempty" type:"Struct"`
}

func (s MuteAudioAllResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MuteAudioAllResponseBody) GoString() string {
	return s.String()
}

func (s *MuteAudioAllResponseBody) SetRequestId(v string) *MuteAudioAllResponseBody {
	s.RequestId = &v
	return s
}

func (s *MuteAudioAllResponseBody) SetConferenceId(v string) *MuteAudioAllResponseBody {
	s.ConferenceId = &v
	return s
}

func (s *MuteAudioAllResponseBody) SetParticipants(v *MuteAudioAllResponseBodyParticipants) *MuteAudioAllResponseBody {
	s.Participants = v
	return s
}

type MuteAudioAllResponseBodyParticipants struct {
	Participant []*MuteAudioAllResponseBodyParticipantsParticipant `json:"Participant,omitempty" xml:"Participant,omitempty" type:"Repeated"`
}

func (s MuteAudioAllResponseBodyParticipants) String() string {
	return tea.Prettify(s)
}

func (s MuteAudioAllResponseBodyParticipants) GoString() string {
	return s.String()
}

func (s *MuteAudioAllResponseBodyParticipants) SetParticipant(v []*MuteAudioAllResponseBodyParticipantsParticipant) *MuteAudioAllResponseBodyParticipants {
	s.Participant = v
	return s
}

type MuteAudioAllResponseBodyParticipantsParticipant struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s MuteAudioAllResponseBodyParticipantsParticipant) String() string {
	return tea.Prettify(s)
}

func (s MuteAudioAllResponseBodyParticipantsParticipant) GoString() string {
	return s.String()
}

func (s *MuteAudioAllResponseBodyParticipantsParticipant) SetCode(v string) *MuteAudioAllResponseBodyParticipantsParticipant {
	s.Code = &v
	return s
}

func (s *MuteAudioAllResponseBodyParticipantsParticipant) SetMessage(v string) *MuteAudioAllResponseBodyParticipantsParticipant {
	s.Message = &v
	return s
}

func (s *MuteAudioAllResponseBodyParticipantsParticipant) SetId(v string) *MuteAudioAllResponseBodyParticipantsParticipant {
	s.Id = &v
	return s
}

type MuteAudioAllResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MuteAudioAllResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MuteAudioAllResponse) String() string {
	return tea.Prettify(s)
}

func (s MuteAudioAllResponse) GoString() string {
	return s.String()
}

func (s *MuteAudioAllResponse) SetHeaders(v map[string]*string) *MuteAudioAllResponse {
	s.Headers = v
	return s
}

func (s *MuteAudioAllResponse) SetBody(v *MuteAudioAllResponseBody) *MuteAudioAllResponse {
	s.Body = v
	return s
}

type ReceiveNotifyRequest struct {
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog     *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	TraceId     *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
	BizId       *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	Event       *string `json:"Event,omitempty" xml:"Event,omitempty"`
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s ReceiveNotifyRequest) String() string {
	return tea.Prettify(s)
}

func (s ReceiveNotifyRequest) GoString() string {
	return s.String()
}

func (s *ReceiveNotifyRequest) SetOwnerId(v int64) *ReceiveNotifyRequest {
	s.OwnerId = &v
	return s
}

func (s *ReceiveNotifyRequest) SetShowLog(v string) *ReceiveNotifyRequest {
	s.ShowLog = &v
	return s
}

func (s *ReceiveNotifyRequest) SetTraceId(v string) *ReceiveNotifyRequest {
	s.TraceId = &v
	return s
}

func (s *ReceiveNotifyRequest) SetBizId(v string) *ReceiveNotifyRequest {
	s.BizId = &v
	return s
}

func (s *ReceiveNotifyRequest) SetEvent(v string) *ReceiveNotifyRequest {
	s.Event = &v
	return s
}

func (s *ReceiveNotifyRequest) SetContentType(v string) *ReceiveNotifyRequest {
	s.ContentType = &v
	return s
}

func (s *ReceiveNotifyRequest) SetContent(v string) *ReceiveNotifyRequest {
	s.Content = &v
	return s
}

type ReceiveNotifyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TraceId   *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ReceiveNotifyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReceiveNotifyResponseBody) GoString() string {
	return s.String()
}

func (s *ReceiveNotifyResponseBody) SetRequestId(v string) *ReceiveNotifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReceiveNotifyResponseBody) SetTraceId(v string) *ReceiveNotifyResponseBody {
	s.TraceId = &v
	return s
}

type ReceiveNotifyResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReceiveNotifyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReceiveNotifyResponse) String() string {
	return tea.Prettify(s)
}

func (s ReceiveNotifyResponse) GoString() string {
	return s.String()
}

func (s *ReceiveNotifyResponse) SetHeaders(v map[string]*string) *ReceiveNotifyResponse {
	s.Headers = v
	return s
}

func (s *ReceiveNotifyResponse) SetBody(v *ReceiveNotifyResponseBody) *ReceiveNotifyResponse {
	s.Body = v
	return s
}

type RemoveParticipantsRequest struct {
	OwnerId        *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog        *string   `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId          *string   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ConferenceId   *string   `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	ParticipantIds []*string `json:"ParticipantIds,omitempty" xml:"ParticipantIds,omitempty" type:"Repeated"`
}

func (s RemoveParticipantsRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveParticipantsRequest) GoString() string {
	return s.String()
}

func (s *RemoveParticipantsRequest) SetOwnerId(v int64) *RemoveParticipantsRequest {
	s.OwnerId = &v
	return s
}

func (s *RemoveParticipantsRequest) SetShowLog(v string) *RemoveParticipantsRequest {
	s.ShowLog = &v
	return s
}

func (s *RemoveParticipantsRequest) SetAppId(v string) *RemoveParticipantsRequest {
	s.AppId = &v
	return s
}

func (s *RemoveParticipantsRequest) SetConferenceId(v string) *RemoveParticipantsRequest {
	s.ConferenceId = &v
	return s
}

func (s *RemoveParticipantsRequest) SetParticipantIds(v []*string) *RemoveParticipantsRequest {
	s.ParticipantIds = v
	return s
}

type RemoveParticipantsResponseBody struct {
	RequestId    *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ConferenceId *string                                     `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	Participants *RemoveParticipantsResponseBodyParticipants `json:"Participants,omitempty" xml:"Participants,omitempty" type:"Struct"`
}

func (s RemoveParticipantsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveParticipantsResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveParticipantsResponseBody) SetRequestId(v string) *RemoveParticipantsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveParticipantsResponseBody) SetConferenceId(v string) *RemoveParticipantsResponseBody {
	s.ConferenceId = &v
	return s
}

func (s *RemoveParticipantsResponseBody) SetParticipants(v *RemoveParticipantsResponseBodyParticipants) *RemoveParticipantsResponseBody {
	s.Participants = v
	return s
}

type RemoveParticipantsResponseBodyParticipants struct {
	Participant []*RemoveParticipantsResponseBodyParticipantsParticipant `json:"Participant,omitempty" xml:"Participant,omitempty" type:"Repeated"`
}

func (s RemoveParticipantsResponseBodyParticipants) String() string {
	return tea.Prettify(s)
}

func (s RemoveParticipantsResponseBodyParticipants) GoString() string {
	return s.String()
}

func (s *RemoveParticipantsResponseBodyParticipants) SetParticipant(v []*RemoveParticipantsResponseBodyParticipantsParticipant) *RemoveParticipantsResponseBodyParticipants {
	s.Participant = v
	return s
}

type RemoveParticipantsResponseBodyParticipantsParticipant struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s RemoveParticipantsResponseBodyParticipantsParticipant) String() string {
	return tea.Prettify(s)
}

func (s RemoveParticipantsResponseBodyParticipantsParticipant) GoString() string {
	return s.String()
}

func (s *RemoveParticipantsResponseBodyParticipantsParticipant) SetCode(v string) *RemoveParticipantsResponseBodyParticipantsParticipant {
	s.Code = &v
	return s
}

func (s *RemoveParticipantsResponseBodyParticipantsParticipant) SetMessage(v string) *RemoveParticipantsResponseBodyParticipantsParticipant {
	s.Message = &v
	return s
}

func (s *RemoveParticipantsResponseBodyParticipantsParticipant) SetId(v string) *RemoveParticipantsResponseBodyParticipantsParticipant {
	s.Id = &v
	return s
}

type RemoveParticipantsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveParticipantsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveParticipantsResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveParticipantsResponse) GoString() string {
	return s.String()
}

func (s *RemoveParticipantsResponse) SetHeaders(v map[string]*string) *RemoveParticipantsResponse {
	s.Headers = v
	return s
}

func (s *RemoveParticipantsResponse) SetBody(v *RemoveParticipantsResponseBody) *RemoveParticipantsResponse {
	s.Body = v
	return s
}

type RemoveTerminalsRequest struct {
	OwnerId     *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog     *string   `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId       *string   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId   *string   `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	TerminalIds []*string `json:"TerminalIds,omitempty" xml:"TerminalIds,omitempty" type:"Repeated"`
}

func (s RemoveTerminalsRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveTerminalsRequest) GoString() string {
	return s.String()
}

func (s *RemoveTerminalsRequest) SetOwnerId(v int64) *RemoveTerminalsRequest {
	s.OwnerId = &v
	return s
}

func (s *RemoveTerminalsRequest) SetShowLog(v string) *RemoveTerminalsRequest {
	s.ShowLog = &v
	return s
}

func (s *RemoveTerminalsRequest) SetAppId(v string) *RemoveTerminalsRequest {
	s.AppId = &v
	return s
}

func (s *RemoveTerminalsRequest) SetChannelId(v string) *RemoveTerminalsRequest {
	s.ChannelId = &v
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
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
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

func (s *RemoveTerminalsResponseBodyTerminalsTerminal) SetMessage(v string) *RemoveTerminalsResponseBodyTerminalsTerminal {
	s.Message = &v
	return s
}

func (s *RemoveTerminalsResponseBodyTerminalsTerminal) SetId(v string) *RemoveTerminalsResponseBodyTerminalsTerminal {
	s.Id = &v
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

type SetChannelPropertyRequest struct {
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog    *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId  *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	MaxUserNum *int32  `json:"MaxUserNum,omitempty" xml:"MaxUserNum,omitempty"`
	StartTime  *int32  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Duration   *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Priority   *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Topics     *string `json:"Topics,omitempty" xml:"Topics,omitempty"`
}

func (s SetChannelPropertyRequest) String() string {
	return tea.Prettify(s)
}

func (s SetChannelPropertyRequest) GoString() string {
	return s.String()
}

func (s *SetChannelPropertyRequest) SetOwnerId(v int64) *SetChannelPropertyRequest {
	s.OwnerId = &v
	return s
}

func (s *SetChannelPropertyRequest) SetShowLog(v string) *SetChannelPropertyRequest {
	s.ShowLog = &v
	return s
}

func (s *SetChannelPropertyRequest) SetAppId(v string) *SetChannelPropertyRequest {
	s.AppId = &v
	return s
}

func (s *SetChannelPropertyRequest) SetChannelId(v string) *SetChannelPropertyRequest {
	s.ChannelId = &v
	return s
}

func (s *SetChannelPropertyRequest) SetMaxUserNum(v int32) *SetChannelPropertyRequest {
	s.MaxUserNum = &v
	return s
}

func (s *SetChannelPropertyRequest) SetStartTime(v int32) *SetChannelPropertyRequest {
	s.StartTime = &v
	return s
}

func (s *SetChannelPropertyRequest) SetDuration(v int32) *SetChannelPropertyRequest {
	s.Duration = &v
	return s
}

func (s *SetChannelPropertyRequest) SetPriority(v string) *SetChannelPropertyRequest {
	s.Priority = &v
	return s
}

func (s *SetChannelPropertyRequest) SetTopics(v string) *SetChannelPropertyRequest {
	s.Topics = &v
	return s
}

type SetChannelPropertyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetChannelPropertyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetChannelPropertyResponseBody) GoString() string {
	return s.String()
}

func (s *SetChannelPropertyResponseBody) SetRequestId(v string) *SetChannelPropertyResponseBody {
	s.RequestId = &v
	return s
}

type SetChannelPropertyResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetChannelPropertyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetChannelPropertyResponse) String() string {
	return tea.Prettify(s)
}

func (s SetChannelPropertyResponse) GoString() string {
	return s.String()
}

func (s *SetChannelPropertyResponse) SetHeaders(v map[string]*string) *SetChannelPropertyResponse {
	s.Headers = v
	return s
}

func (s *SetChannelPropertyResponse) SetBody(v *SetChannelPropertyResponseBody) *SetChannelPropertyResponse {
	s.Body = v
	return s
}

type StartMPUTaskRequest struct {
	OwnerId           *int64                             `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog           *string                            `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId             *string                            `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId         *string                            `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	TaskId            *string                            `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskProfile       *string                            `json:"TaskProfile,omitempty" xml:"TaskProfile,omitempty"`
	TaskMode          *int32                             `json:"TaskMode,omitempty" xml:"TaskMode,omitempty"`
	MixMode           *int32                             `json:"MixMode,omitempty" xml:"MixMode,omitempty"`
	CropMode          *int32                             `json:"CropMode,omitempty" xml:"CropMode,omitempty"`
	MediaEncode       *int32                             `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	SourceType        *string                            `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	StreamType        *int32                             `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	BackgroundColor   *int32                             `json:"BackgroundColor,omitempty" xml:"BackgroundColor,omitempty"`
	StreamURL         *string                            `json:"StreamURL,omitempty" xml:"StreamURL,omitempty"`
	PayloadType       *int32                             `json:"PayloadType,omitempty" xml:"PayloadType,omitempty"`
	ReportVad         *int32                             `json:"ReportVad,omitempty" xml:"ReportVad,omitempty"`
	RtpExtInfo        *int32                             `json:"RtpExtInfo,omitempty" xml:"RtpExtInfo,omitempty"`
	TimeStampRef      *int64                             `json:"TimeStampRef,omitempty" xml:"TimeStampRef,omitempty"`
	VadInterval       *int64                             `json:"VadInterval,omitempty" xml:"VadInterval,omitempty"`
	SubSpecUsers      []*string                          `json:"SubSpecUsers,omitempty" xml:"SubSpecUsers,omitempty" type:"Repeated"`
	SubSpecAudioUsers []*string                          `json:"SubSpecAudioUsers,omitempty" xml:"SubSpecAudioUsers,omitempty" type:"Repeated"`
	LayoutIds         []*int                             `json:"LayoutIds,omitempty" xml:"LayoutIds,omitempty" type:"Repeated"`
	UserPanes         []*StartMPUTaskRequestUserPanes    `json:"UserPanes,omitempty" xml:"UserPanes,omitempty" type:"Repeated"`
	Backgrounds       []*StartMPUTaskRequestBackgrounds  `json:"Backgrounds,omitempty" xml:"Backgrounds,omitempty" type:"Repeated"`
	Watermarks        []*StartMPUTaskRequestWatermarks   `json:"Watermarks,omitempty" xml:"Watermarks,omitempty" type:"Repeated"`
	ClockWidgets      []*StartMPUTaskRequestClockWidgets `json:"ClockWidgets,omitempty" xml:"ClockWidgets,omitempty" type:"Repeated"`
}

func (s StartMPUTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s StartMPUTaskRequest) GoString() string {
	return s.String()
}

func (s *StartMPUTaskRequest) SetOwnerId(v int64) *StartMPUTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *StartMPUTaskRequest) SetShowLog(v string) *StartMPUTaskRequest {
	s.ShowLog = &v
	return s
}

func (s *StartMPUTaskRequest) SetAppId(v string) *StartMPUTaskRequest {
	s.AppId = &v
	return s
}

func (s *StartMPUTaskRequest) SetChannelId(v string) *StartMPUTaskRequest {
	s.ChannelId = &v
	return s
}

func (s *StartMPUTaskRequest) SetTaskId(v string) *StartMPUTaskRequest {
	s.TaskId = &v
	return s
}

func (s *StartMPUTaskRequest) SetTaskProfile(v string) *StartMPUTaskRequest {
	s.TaskProfile = &v
	return s
}

func (s *StartMPUTaskRequest) SetTaskMode(v int32) *StartMPUTaskRequest {
	s.TaskMode = &v
	return s
}

func (s *StartMPUTaskRequest) SetMixMode(v int32) *StartMPUTaskRequest {
	s.MixMode = &v
	return s
}

func (s *StartMPUTaskRequest) SetCropMode(v int32) *StartMPUTaskRequest {
	s.CropMode = &v
	return s
}

func (s *StartMPUTaskRequest) SetMediaEncode(v int32) *StartMPUTaskRequest {
	s.MediaEncode = &v
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

func (s *StartMPUTaskRequest) SetBackgroundColor(v int32) *StartMPUTaskRequest {
	s.BackgroundColor = &v
	return s
}

func (s *StartMPUTaskRequest) SetStreamURL(v string) *StartMPUTaskRequest {
	s.StreamURL = &v
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

func (s *StartMPUTaskRequest) SetTimeStampRef(v int64) *StartMPUTaskRequest {
	s.TimeStampRef = &v
	return s
}

func (s *StartMPUTaskRequest) SetVadInterval(v int64) *StartMPUTaskRequest {
	s.VadInterval = &v
	return s
}

func (s *StartMPUTaskRequest) SetSubSpecUsers(v []*string) *StartMPUTaskRequest {
	s.SubSpecUsers = v
	return s
}

func (s *StartMPUTaskRequest) SetSubSpecAudioUsers(v []*string) *StartMPUTaskRequest {
	s.SubSpecAudioUsers = v
	return s
}

func (s *StartMPUTaskRequest) SetLayoutIds(v []*int) *StartMPUTaskRequest {
	s.LayoutIds = v
	return s
}

func (s *StartMPUTaskRequest) SetUserPanes(v []*StartMPUTaskRequestUserPanes) *StartMPUTaskRequest {
	s.UserPanes = v
	return s
}

func (s *StartMPUTaskRequest) SetBackgrounds(v []*StartMPUTaskRequestBackgrounds) *StartMPUTaskRequest {
	s.Backgrounds = v
	return s
}

func (s *StartMPUTaskRequest) SetWatermarks(v []*StartMPUTaskRequestWatermarks) *StartMPUTaskRequest {
	s.Watermarks = v
	return s
}

func (s *StartMPUTaskRequest) SetClockWidgets(v []*StartMPUTaskRequestClockWidgets) *StartMPUTaskRequest {
	s.ClockWidgets = v
	return s
}

type StartMPUTaskRequestUserPanes struct {
	Images      []*StartMPUTaskRequestUserPanesImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	SegmentType *int32                                `json:"SegmentType,omitempty" xml:"SegmentType,omitempty"`
	UserId      *string                               `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Texts       []*StartMPUTaskRequestUserPanesTexts  `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
	SourceType  *string                               `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	PaneId      *int32                                `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
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

func (s *StartMPUTaskRequestUserPanes) SetSegmentType(v int32) *StartMPUTaskRequestUserPanes {
	s.SegmentType = &v
	return s
}

func (s *StartMPUTaskRequestUserPanes) SetUserId(v string) *StartMPUTaskRequestUserPanes {
	s.UserId = &v
	return s
}

func (s *StartMPUTaskRequestUserPanes) SetTexts(v []*StartMPUTaskRequestUserPanesTexts) *StartMPUTaskRequestUserPanes {
	s.Texts = v
	return s
}

func (s *StartMPUTaskRequestUserPanes) SetSourceType(v string) *StartMPUTaskRequestUserPanes {
	s.SourceType = &v
	return s
}

func (s *StartMPUTaskRequestUserPanes) SetPaneId(v int32) *StartMPUTaskRequestUserPanes {
	s.PaneId = &v
	return s
}

type StartMPUTaskRequestUserPanesImages struct {
	Width   *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height  *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Y       *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	Url     *string  `json:"Url,omitempty" xml:"Url,omitempty"`
	Display *int32   `json:"Display,omitempty" xml:"Display,omitempty"`
	ZOrder  *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
	X       *float32 `json:"X,omitempty" xml:"X,omitempty"`
}

func (s StartMPUTaskRequestUserPanesImages) String() string {
	return tea.Prettify(s)
}

func (s StartMPUTaskRequestUserPanesImages) GoString() string {
	return s.String()
}

func (s *StartMPUTaskRequestUserPanesImages) SetWidth(v float32) *StartMPUTaskRequestUserPanesImages {
	s.Width = &v
	return s
}

func (s *StartMPUTaskRequestUserPanesImages) SetHeight(v float32) *StartMPUTaskRequestUserPanesImages {
	s.Height = &v
	return s
}

func (s *StartMPUTaskRequestUserPanesImages) SetY(v float32) *StartMPUTaskRequestUserPanesImages {
	s.Y = &v
	return s
}

func (s *StartMPUTaskRequestUserPanesImages) SetUrl(v string) *StartMPUTaskRequestUserPanesImages {
	s.Url = &v
	return s
}

func (s *StartMPUTaskRequestUserPanesImages) SetDisplay(v int32) *StartMPUTaskRequestUserPanesImages {
	s.Display = &v
	return s
}

func (s *StartMPUTaskRequestUserPanesImages) SetZOrder(v int32) *StartMPUTaskRequestUserPanesImages {
	s.ZOrder = &v
	return s
}

func (s *StartMPUTaskRequestUserPanesImages) SetX(v float32) *StartMPUTaskRequestUserPanesImages {
	s.X = &v
	return s
}

type StartMPUTaskRequestUserPanesTexts struct {
	FontType  *int32   `json:"FontType,omitempty" xml:"FontType,omitempty"`
	FontColor *int32   `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	Y         *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	Text      *string  `json:"Text,omitempty" xml:"Text,omitempty"`
	ZOrder    *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
	FontSize  *int32   `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	X         *float32 `json:"X,omitempty" xml:"X,omitempty"`
}

func (s StartMPUTaskRequestUserPanesTexts) String() string {
	return tea.Prettify(s)
}

func (s StartMPUTaskRequestUserPanesTexts) GoString() string {
	return s.String()
}

func (s *StartMPUTaskRequestUserPanesTexts) SetFontType(v int32) *StartMPUTaskRequestUserPanesTexts {
	s.FontType = &v
	return s
}

func (s *StartMPUTaskRequestUserPanesTexts) SetFontColor(v int32) *StartMPUTaskRequestUserPanesTexts {
	s.FontColor = &v
	return s
}

func (s *StartMPUTaskRequestUserPanesTexts) SetY(v float32) *StartMPUTaskRequestUserPanesTexts {
	s.Y = &v
	return s
}

func (s *StartMPUTaskRequestUserPanesTexts) SetText(v string) *StartMPUTaskRequestUserPanesTexts {
	s.Text = &v
	return s
}

func (s *StartMPUTaskRequestUserPanesTexts) SetZOrder(v int32) *StartMPUTaskRequestUserPanesTexts {
	s.ZOrder = &v
	return s
}

func (s *StartMPUTaskRequestUserPanesTexts) SetFontSize(v int32) *StartMPUTaskRequestUserPanesTexts {
	s.FontSize = &v
	return s
}

func (s *StartMPUTaskRequestUserPanesTexts) SetX(v float32) *StartMPUTaskRequestUserPanesTexts {
	s.X = &v
	return s
}

type StartMPUTaskRequestBackgrounds struct {
	Width   *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height  *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Y       *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	Url     *string  `json:"Url,omitempty" xml:"Url,omitempty"`
	Display *int32   `json:"Display,omitempty" xml:"Display,omitempty"`
	ZOrder  *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
	X       *float32 `json:"X,omitempty" xml:"X,omitempty"`
}

func (s StartMPUTaskRequestBackgrounds) String() string {
	return tea.Prettify(s)
}

func (s StartMPUTaskRequestBackgrounds) GoString() string {
	return s.String()
}

func (s *StartMPUTaskRequestBackgrounds) SetWidth(v float32) *StartMPUTaskRequestBackgrounds {
	s.Width = &v
	return s
}

func (s *StartMPUTaskRequestBackgrounds) SetHeight(v float32) *StartMPUTaskRequestBackgrounds {
	s.Height = &v
	return s
}

func (s *StartMPUTaskRequestBackgrounds) SetY(v float32) *StartMPUTaskRequestBackgrounds {
	s.Y = &v
	return s
}

func (s *StartMPUTaskRequestBackgrounds) SetUrl(v string) *StartMPUTaskRequestBackgrounds {
	s.Url = &v
	return s
}

func (s *StartMPUTaskRequestBackgrounds) SetDisplay(v int32) *StartMPUTaskRequestBackgrounds {
	s.Display = &v
	return s
}

func (s *StartMPUTaskRequestBackgrounds) SetZOrder(v int32) *StartMPUTaskRequestBackgrounds {
	s.ZOrder = &v
	return s
}

func (s *StartMPUTaskRequestBackgrounds) SetX(v float32) *StartMPUTaskRequestBackgrounds {
	s.X = &v
	return s
}

type StartMPUTaskRequestWatermarks struct {
	Alpha   *float32 `json:"Alpha,omitempty" xml:"Alpha,omitempty"`
	Width   *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height  *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Y       *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	Url     *string  `json:"Url,omitempty" xml:"Url,omitempty"`
	Display *int32   `json:"Display,omitempty" xml:"Display,omitempty"`
	ZOrder  *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
	X       *float32 `json:"X,omitempty" xml:"X,omitempty"`
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

func (s *StartMPUTaskRequestWatermarks) SetWidth(v float32) *StartMPUTaskRequestWatermarks {
	s.Width = &v
	return s
}

func (s *StartMPUTaskRequestWatermarks) SetHeight(v float32) *StartMPUTaskRequestWatermarks {
	s.Height = &v
	return s
}

func (s *StartMPUTaskRequestWatermarks) SetY(v float32) *StartMPUTaskRequestWatermarks {
	s.Y = &v
	return s
}

func (s *StartMPUTaskRequestWatermarks) SetUrl(v string) *StartMPUTaskRequestWatermarks {
	s.Url = &v
	return s
}

func (s *StartMPUTaskRequestWatermarks) SetDisplay(v int32) *StartMPUTaskRequestWatermarks {
	s.Display = &v
	return s
}

func (s *StartMPUTaskRequestWatermarks) SetZOrder(v int32) *StartMPUTaskRequestWatermarks {
	s.ZOrder = &v
	return s
}

func (s *StartMPUTaskRequestWatermarks) SetX(v float32) *StartMPUTaskRequestWatermarks {
	s.X = &v
	return s
}

type StartMPUTaskRequestClockWidgets struct {
	FontType  *int32   `json:"FontType,omitempty" xml:"FontType,omitempty"`
	FontColor *int32   `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	Y         *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	ZOrder    *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
	X         *float32 `json:"X,omitempty" xml:"X,omitempty"`
	FontSize  *int32   `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
}

func (s StartMPUTaskRequestClockWidgets) String() string {
	return tea.Prettify(s)
}

func (s StartMPUTaskRequestClockWidgets) GoString() string {
	return s.String()
}

func (s *StartMPUTaskRequestClockWidgets) SetFontType(v int32) *StartMPUTaskRequestClockWidgets {
	s.FontType = &v
	return s
}

func (s *StartMPUTaskRequestClockWidgets) SetFontColor(v int32) *StartMPUTaskRequestClockWidgets {
	s.FontColor = &v
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

func (s *StartMPUTaskRequestClockWidgets) SetX(v float32) *StartMPUTaskRequestClockWidgets {
	s.X = &v
	return s
}

func (s *StartMPUTaskRequestClockWidgets) SetFontSize(v int32) *StartMPUTaskRequestClockWidgets {
	s.FontSize = &v
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
	OwnerId      *int64                             `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog      *string                            `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId        *string                            `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId    *string                            `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	TaskId       *string                            `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskProfile  *string                            `json:"TaskProfile,omitempty" xml:"TaskProfile,omitempty"`
	MediaEncode  *int32                             `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	TemplateId   *string                            `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	SubSpecUsers []*string                          `json:"SubSpecUsers,omitempty" xml:"SubSpecUsers,omitempty" type:"Repeated"`
	UserPanes    []*StartRecordTaskRequestUserPanes `json:"UserPanes,omitempty" xml:"UserPanes,omitempty" type:"Repeated"`
}

func (s StartRecordTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s StartRecordTaskRequest) GoString() string {
	return s.String()
}

func (s *StartRecordTaskRequest) SetOwnerId(v int64) *StartRecordTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *StartRecordTaskRequest) SetShowLog(v string) *StartRecordTaskRequest {
	s.ShowLog = &v
	return s
}

func (s *StartRecordTaskRequest) SetAppId(v string) *StartRecordTaskRequest {
	s.AppId = &v
	return s
}

func (s *StartRecordTaskRequest) SetChannelId(v string) *StartRecordTaskRequest {
	s.ChannelId = &v
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

func (s *StartRecordTaskRequest) SetMediaEncode(v int32) *StartRecordTaskRequest {
	s.MediaEncode = &v
	return s
}

func (s *StartRecordTaskRequest) SetTemplateId(v string) *StartRecordTaskRequest {
	s.TemplateId = &v
	return s
}

func (s *StartRecordTaskRequest) SetSubSpecUsers(v []*string) *StartRecordTaskRequest {
	s.SubSpecUsers = v
	return s
}

func (s *StartRecordTaskRequest) SetUserPanes(v []*StartRecordTaskRequestUserPanes) *StartRecordTaskRequest {
	s.UserPanes = v
	return s
}

type StartRecordTaskRequestUserPanes struct {
	Images     []*StartRecordTaskRequestUserPanesImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	UserId     *string                                  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Texts      []*StartRecordTaskRequestUserPanesTexts  `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
	SourceType *string                                  `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	PaneId     *int32                                   `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
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

func (s *StartRecordTaskRequestUserPanes) SetUserId(v string) *StartRecordTaskRequestUserPanes {
	s.UserId = &v
	return s
}

func (s *StartRecordTaskRequestUserPanes) SetTexts(v []*StartRecordTaskRequestUserPanesTexts) *StartRecordTaskRequestUserPanes {
	s.Texts = v
	return s
}

func (s *StartRecordTaskRequestUserPanes) SetSourceType(v string) *StartRecordTaskRequestUserPanes {
	s.SourceType = &v
	return s
}

func (s *StartRecordTaskRequestUserPanes) SetPaneId(v int32) *StartRecordTaskRequestUserPanes {
	s.PaneId = &v
	return s
}

type StartRecordTaskRequestUserPanesImages struct {
	Width   *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height  *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Y       *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	Url     *string  `json:"Url,omitempty" xml:"Url,omitempty"`
	Display *int32   `json:"Display,omitempty" xml:"Display,omitempty"`
	ZOrder  *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
	X       *float32 `json:"X,omitempty" xml:"X,omitempty"`
}

func (s StartRecordTaskRequestUserPanesImages) String() string {
	return tea.Prettify(s)
}

func (s StartRecordTaskRequestUserPanesImages) GoString() string {
	return s.String()
}

func (s *StartRecordTaskRequestUserPanesImages) SetWidth(v float32) *StartRecordTaskRequestUserPanesImages {
	s.Width = &v
	return s
}

func (s *StartRecordTaskRequestUserPanesImages) SetHeight(v float32) *StartRecordTaskRequestUserPanesImages {
	s.Height = &v
	return s
}

func (s *StartRecordTaskRequestUserPanesImages) SetY(v float32) *StartRecordTaskRequestUserPanesImages {
	s.Y = &v
	return s
}

func (s *StartRecordTaskRequestUserPanesImages) SetUrl(v string) *StartRecordTaskRequestUserPanesImages {
	s.Url = &v
	return s
}

func (s *StartRecordTaskRequestUserPanesImages) SetDisplay(v int32) *StartRecordTaskRequestUserPanesImages {
	s.Display = &v
	return s
}

func (s *StartRecordTaskRequestUserPanesImages) SetZOrder(v int32) *StartRecordTaskRequestUserPanesImages {
	s.ZOrder = &v
	return s
}

func (s *StartRecordTaskRequestUserPanesImages) SetX(v float32) *StartRecordTaskRequestUserPanesImages {
	s.X = &v
	return s
}

type StartRecordTaskRequestUserPanesTexts struct {
	FontType  *int32   `json:"FontType,omitempty" xml:"FontType,omitempty"`
	FontColor *int32   `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	Y         *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	Text      *string  `json:"Text,omitempty" xml:"Text,omitempty"`
	ZOrder    *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
	FontSize  *int32   `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	X         *float32 `json:"X,omitempty" xml:"X,omitempty"`
}

func (s StartRecordTaskRequestUserPanesTexts) String() string {
	return tea.Prettify(s)
}

func (s StartRecordTaskRequestUserPanesTexts) GoString() string {
	return s.String()
}

func (s *StartRecordTaskRequestUserPanesTexts) SetFontType(v int32) *StartRecordTaskRequestUserPanesTexts {
	s.FontType = &v
	return s
}

func (s *StartRecordTaskRequestUserPanesTexts) SetFontColor(v int32) *StartRecordTaskRequestUserPanesTexts {
	s.FontColor = &v
	return s
}

func (s *StartRecordTaskRequestUserPanesTexts) SetY(v float32) *StartRecordTaskRequestUserPanesTexts {
	s.Y = &v
	return s
}

func (s *StartRecordTaskRequestUserPanesTexts) SetText(v string) *StartRecordTaskRequestUserPanesTexts {
	s.Text = &v
	return s
}

func (s *StartRecordTaskRequestUserPanesTexts) SetZOrder(v int32) *StartRecordTaskRequestUserPanesTexts {
	s.ZOrder = &v
	return s
}

func (s *StartRecordTaskRequestUserPanesTexts) SetFontSize(v int32) *StartRecordTaskRequestUserPanesTexts {
	s.FontSize = &v
	return s
}

func (s *StartRecordTaskRequestUserPanesTexts) SetX(v float32) *StartRecordTaskRequestUserPanesTexts {
	s.X = &v
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
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog   *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	UserId    *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s StopChannelUserPublishRequest) String() string {
	return tea.Prettify(s)
}

func (s StopChannelUserPublishRequest) GoString() string {
	return s.String()
}

func (s *StopChannelUserPublishRequest) SetOwnerId(v int64) *StopChannelUserPublishRequest {
	s.OwnerId = &v
	return s
}

func (s *StopChannelUserPublishRequest) SetShowLog(v string) *StopChannelUserPublishRequest {
	s.ShowLog = &v
	return s
}

func (s *StopChannelUserPublishRequest) SetAppId(v string) *StopChannelUserPublishRequest {
	s.AppId = &v
	return s
}

func (s *StopChannelUserPublishRequest) SetChannelId(v string) *StopChannelUserPublishRequest {
	s.ChannelId = &v
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
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	TaskId  *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopMPUTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s StopMPUTaskRequest) GoString() string {
	return s.String()
}

func (s *StopMPUTaskRequest) SetOwnerId(v int64) *StopMPUTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *StopMPUTaskRequest) SetShowLog(v string) *StopMPUTaskRequest {
	s.ShowLog = &v
	return s
}

func (s *StopMPUTaskRequest) SetAppId(v string) *StopMPUTaskRequest {
	s.AppId = &v
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
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	TaskId  *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopRecordTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s StopRecordTaskRequest) GoString() string {
	return s.String()
}

func (s *StopRecordTaskRequest) SetOwnerId(v int64) *StopRecordTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *StopRecordTaskRequest) SetShowLog(v string) *StopRecordTaskRequest {
	s.ShowLog = &v
	return s
}

func (s *StopRecordTaskRequest) SetAppId(v string) *StopRecordTaskRequest {
	s.AppId = &v
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

type UnmuteAudioRequest struct {
	OwnerId        *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog        *string   `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId          *string   `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ConferenceId   *string   `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	ParticipantIds []*string `json:"ParticipantIds,omitempty" xml:"ParticipantIds,omitempty" type:"Repeated"`
}

func (s UnmuteAudioRequest) String() string {
	return tea.Prettify(s)
}

func (s UnmuteAudioRequest) GoString() string {
	return s.String()
}

func (s *UnmuteAudioRequest) SetOwnerId(v int64) *UnmuteAudioRequest {
	s.OwnerId = &v
	return s
}

func (s *UnmuteAudioRequest) SetShowLog(v string) *UnmuteAudioRequest {
	s.ShowLog = &v
	return s
}

func (s *UnmuteAudioRequest) SetAppId(v string) *UnmuteAudioRequest {
	s.AppId = &v
	return s
}

func (s *UnmuteAudioRequest) SetConferenceId(v string) *UnmuteAudioRequest {
	s.ConferenceId = &v
	return s
}

func (s *UnmuteAudioRequest) SetParticipantIds(v []*string) *UnmuteAudioRequest {
	s.ParticipantIds = v
	return s
}

type UnmuteAudioResponseBody struct {
	RequestId    *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ConferenceId *string                              `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	Participants *UnmuteAudioResponseBodyParticipants `json:"Participants,omitempty" xml:"Participants,omitempty" type:"Struct"`
}

func (s UnmuteAudioResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnmuteAudioResponseBody) GoString() string {
	return s.String()
}

func (s *UnmuteAudioResponseBody) SetRequestId(v string) *UnmuteAudioResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnmuteAudioResponseBody) SetConferenceId(v string) *UnmuteAudioResponseBody {
	s.ConferenceId = &v
	return s
}

func (s *UnmuteAudioResponseBody) SetParticipants(v *UnmuteAudioResponseBodyParticipants) *UnmuteAudioResponseBody {
	s.Participants = v
	return s
}

type UnmuteAudioResponseBodyParticipants struct {
	Participant []*UnmuteAudioResponseBodyParticipantsParticipant `json:"Participant,omitempty" xml:"Participant,omitempty" type:"Repeated"`
}

func (s UnmuteAudioResponseBodyParticipants) String() string {
	return tea.Prettify(s)
}

func (s UnmuteAudioResponseBodyParticipants) GoString() string {
	return s.String()
}

func (s *UnmuteAudioResponseBodyParticipants) SetParticipant(v []*UnmuteAudioResponseBodyParticipantsParticipant) *UnmuteAudioResponseBodyParticipants {
	s.Participant = v
	return s
}

type UnmuteAudioResponseBodyParticipantsParticipant struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UnmuteAudioResponseBodyParticipantsParticipant) String() string {
	return tea.Prettify(s)
}

func (s UnmuteAudioResponseBodyParticipantsParticipant) GoString() string {
	return s.String()
}

func (s *UnmuteAudioResponseBodyParticipantsParticipant) SetCode(v string) *UnmuteAudioResponseBodyParticipantsParticipant {
	s.Code = &v
	return s
}

func (s *UnmuteAudioResponseBodyParticipantsParticipant) SetMessage(v string) *UnmuteAudioResponseBodyParticipantsParticipant {
	s.Message = &v
	return s
}

func (s *UnmuteAudioResponseBodyParticipantsParticipant) SetId(v string) *UnmuteAudioResponseBodyParticipantsParticipant {
	s.Id = &v
	return s
}

type UnmuteAudioResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnmuteAudioResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnmuteAudioResponse) String() string {
	return tea.Prettify(s)
}

func (s UnmuteAudioResponse) GoString() string {
	return s.String()
}

func (s *UnmuteAudioResponse) SetHeaders(v map[string]*string) *UnmuteAudioResponse {
	s.Headers = v
	return s
}

func (s *UnmuteAudioResponse) SetBody(v *UnmuteAudioResponseBody) *UnmuteAudioResponse {
	s.Body = v
	return s
}

type UnmuteAudioAllRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog       *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ConferenceId  *string `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	ParticipantId *string `json:"ParticipantId,omitempty" xml:"ParticipantId,omitempty"`
}

func (s UnmuteAudioAllRequest) String() string {
	return tea.Prettify(s)
}

func (s UnmuteAudioAllRequest) GoString() string {
	return s.String()
}

func (s *UnmuteAudioAllRequest) SetOwnerId(v int64) *UnmuteAudioAllRequest {
	s.OwnerId = &v
	return s
}

func (s *UnmuteAudioAllRequest) SetShowLog(v string) *UnmuteAudioAllRequest {
	s.ShowLog = &v
	return s
}

func (s *UnmuteAudioAllRequest) SetAppId(v string) *UnmuteAudioAllRequest {
	s.AppId = &v
	return s
}

func (s *UnmuteAudioAllRequest) SetConferenceId(v string) *UnmuteAudioAllRequest {
	s.ConferenceId = &v
	return s
}

func (s *UnmuteAudioAllRequest) SetParticipantId(v string) *UnmuteAudioAllRequest {
	s.ParticipantId = &v
	return s
}

type UnmuteAudioAllResponseBody struct {
	RequestId    *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ConferenceId *string                                 `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	Participants *UnmuteAudioAllResponseBodyParticipants `json:"Participants,omitempty" xml:"Participants,omitempty" type:"Struct"`
}

func (s UnmuteAudioAllResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnmuteAudioAllResponseBody) GoString() string {
	return s.String()
}

func (s *UnmuteAudioAllResponseBody) SetRequestId(v string) *UnmuteAudioAllResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnmuteAudioAllResponseBody) SetConferenceId(v string) *UnmuteAudioAllResponseBody {
	s.ConferenceId = &v
	return s
}

func (s *UnmuteAudioAllResponseBody) SetParticipants(v *UnmuteAudioAllResponseBodyParticipants) *UnmuteAudioAllResponseBody {
	s.Participants = v
	return s
}

type UnmuteAudioAllResponseBodyParticipants struct {
	Participant []*UnmuteAudioAllResponseBodyParticipantsParticipant `json:"Participant,omitempty" xml:"Participant,omitempty" type:"Repeated"`
}

func (s UnmuteAudioAllResponseBodyParticipants) String() string {
	return tea.Prettify(s)
}

func (s UnmuteAudioAllResponseBodyParticipants) GoString() string {
	return s.String()
}

func (s *UnmuteAudioAllResponseBodyParticipants) SetParticipant(v []*UnmuteAudioAllResponseBodyParticipantsParticipant) *UnmuteAudioAllResponseBodyParticipants {
	s.Participant = v
	return s
}

type UnmuteAudioAllResponseBodyParticipantsParticipant struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UnmuteAudioAllResponseBodyParticipantsParticipant) String() string {
	return tea.Prettify(s)
}

func (s UnmuteAudioAllResponseBodyParticipantsParticipant) GoString() string {
	return s.String()
}

func (s *UnmuteAudioAllResponseBodyParticipantsParticipant) SetCode(v string) *UnmuteAudioAllResponseBodyParticipantsParticipant {
	s.Code = &v
	return s
}

func (s *UnmuteAudioAllResponseBodyParticipantsParticipant) SetMessage(v string) *UnmuteAudioAllResponseBodyParticipantsParticipant {
	s.Message = &v
	return s
}

func (s *UnmuteAudioAllResponseBodyParticipantsParticipant) SetId(v string) *UnmuteAudioAllResponseBodyParticipantsParticipant {
	s.Id = &v
	return s
}

type UnmuteAudioAllResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UnmuteAudioAllResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnmuteAudioAllResponse) String() string {
	return tea.Prettify(s)
}

func (s UnmuteAudioAllResponse) GoString() string {
	return s.String()
}

func (s *UnmuteAudioAllResponse) SetHeaders(v map[string]*string) *UnmuteAudioAllResponse {
	s.Headers = v
	return s
}

func (s *UnmuteAudioAllResponse) SetBody(v *UnmuteAudioAllResponseBody) *UnmuteAudioAllResponse {
	s.Body = v
	return s
}

type UpdateChannelRequest struct {
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog   *string `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	Nonce     *string `json:"Nonce,omitempty" xml:"Nonce,omitempty"`
}

func (s UpdateChannelRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateChannelRequest) GoString() string {
	return s.String()
}

func (s *UpdateChannelRequest) SetOwnerId(v int64) *UpdateChannelRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateChannelRequest) SetShowLog(v string) *UpdateChannelRequest {
	s.ShowLog = &v
	return s
}

func (s *UpdateChannelRequest) SetAppId(v string) *UpdateChannelRequest {
	s.AppId = &v
	return s
}

func (s *UpdateChannelRequest) SetChannelId(v string) *UpdateChannelRequest {
	s.ChannelId = &v
	return s
}

func (s *UpdateChannelRequest) SetNonce(v string) *UpdateChannelRequest {
	s.Nonce = &v
	return s
}

type UpdateChannelResponseBody struct {
	Nonce     *string `json:"Nonce,omitempty" xml:"Nonce,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Timestamp *int32  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s UpdateChannelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateChannelResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateChannelResponseBody) SetNonce(v string) *UpdateChannelResponseBody {
	s.Nonce = &v
	return s
}

func (s *UpdateChannelResponseBody) SetRequestId(v string) *UpdateChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateChannelResponseBody) SetTimestamp(v int32) *UpdateChannelResponseBody {
	s.Timestamp = &v
	return s
}

type UpdateChannelResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateChannelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateChannelResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateChannelResponse) GoString() string {
	return s.String()
}

func (s *UpdateChannelResponse) SetHeaders(v map[string]*string) *UpdateChannelResponse {
	s.Headers = v
	return s
}

func (s *UpdateChannelResponse) SetBody(v *UpdateChannelResponseBody) *UpdateChannelResponse {
	s.Body = v
	return s
}

type UpdateMPULayoutRequest struct {
	OwnerId         *int64                                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog         *string                               `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId           *string                               `json:"AppId,omitempty" xml:"AppId,omitempty"`
	TaskId          *string                               `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	CropMode        *int32                                `json:"CropMode,omitempty" xml:"CropMode,omitempty"`
	BackgroundColor *int32                                `json:"BackgroundColor,omitempty" xml:"BackgroundColor,omitempty"`
	LayoutIds       []*int                                `json:"LayoutIds,omitempty" xml:"LayoutIds,omitempty" type:"Repeated"`
	SubSpecUsers    []*string                             `json:"SubSpecUsers,omitempty" xml:"SubSpecUsers,omitempty" type:"Repeated"`
	UserPanes       []*UpdateMPULayoutRequestUserPanes    `json:"UserPanes,omitempty" xml:"UserPanes,omitempty" type:"Repeated"`
	Backgrounds     []*UpdateMPULayoutRequestBackgrounds  `json:"Backgrounds,omitempty" xml:"Backgrounds,omitempty" type:"Repeated"`
	Watermarks      []*UpdateMPULayoutRequestWatermarks   `json:"Watermarks,omitempty" xml:"Watermarks,omitempty" type:"Repeated"`
	ClockWidgets    []*UpdateMPULayoutRequestClockWidgets `json:"ClockWidgets,omitempty" xml:"ClockWidgets,omitempty" type:"Repeated"`
}

func (s UpdateMPULayoutRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateMPULayoutRequest) GoString() string {
	return s.String()
}

func (s *UpdateMPULayoutRequest) SetOwnerId(v int64) *UpdateMPULayoutRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateMPULayoutRequest) SetShowLog(v string) *UpdateMPULayoutRequest {
	s.ShowLog = &v
	return s
}

func (s *UpdateMPULayoutRequest) SetAppId(v string) *UpdateMPULayoutRequest {
	s.AppId = &v
	return s
}

func (s *UpdateMPULayoutRequest) SetTaskId(v string) *UpdateMPULayoutRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateMPULayoutRequest) SetCropMode(v int32) *UpdateMPULayoutRequest {
	s.CropMode = &v
	return s
}

func (s *UpdateMPULayoutRequest) SetBackgroundColor(v int32) *UpdateMPULayoutRequest {
	s.BackgroundColor = &v
	return s
}

func (s *UpdateMPULayoutRequest) SetLayoutIds(v []*int) *UpdateMPULayoutRequest {
	s.LayoutIds = v
	return s
}

func (s *UpdateMPULayoutRequest) SetSubSpecUsers(v []*string) *UpdateMPULayoutRequest {
	s.SubSpecUsers = v
	return s
}

func (s *UpdateMPULayoutRequest) SetUserPanes(v []*UpdateMPULayoutRequestUserPanes) *UpdateMPULayoutRequest {
	s.UserPanes = v
	return s
}

func (s *UpdateMPULayoutRequest) SetBackgrounds(v []*UpdateMPULayoutRequestBackgrounds) *UpdateMPULayoutRequest {
	s.Backgrounds = v
	return s
}

func (s *UpdateMPULayoutRequest) SetWatermarks(v []*UpdateMPULayoutRequestWatermarks) *UpdateMPULayoutRequest {
	s.Watermarks = v
	return s
}

func (s *UpdateMPULayoutRequest) SetClockWidgets(v []*UpdateMPULayoutRequestClockWidgets) *UpdateMPULayoutRequest {
	s.ClockWidgets = v
	return s
}

type UpdateMPULayoutRequestUserPanes struct {
	Images      []*UpdateMPULayoutRequestUserPanesImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	SegmentType *int32                                   `json:"SegmentType,omitempty" xml:"SegmentType,omitempty"`
	UserId      *string                                  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Texts       []*UpdateMPULayoutRequestUserPanesTexts  `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
	SourceType  *string                                  `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	PaneId      *int32                                   `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
}

func (s UpdateMPULayoutRequestUserPanes) String() string {
	return tea.Prettify(s)
}

func (s UpdateMPULayoutRequestUserPanes) GoString() string {
	return s.String()
}

func (s *UpdateMPULayoutRequestUserPanes) SetImages(v []*UpdateMPULayoutRequestUserPanesImages) *UpdateMPULayoutRequestUserPanes {
	s.Images = v
	return s
}

func (s *UpdateMPULayoutRequestUserPanes) SetSegmentType(v int32) *UpdateMPULayoutRequestUserPanes {
	s.SegmentType = &v
	return s
}

func (s *UpdateMPULayoutRequestUserPanes) SetUserId(v string) *UpdateMPULayoutRequestUserPanes {
	s.UserId = &v
	return s
}

func (s *UpdateMPULayoutRequestUserPanes) SetTexts(v []*UpdateMPULayoutRequestUserPanesTexts) *UpdateMPULayoutRequestUserPanes {
	s.Texts = v
	return s
}

func (s *UpdateMPULayoutRequestUserPanes) SetSourceType(v string) *UpdateMPULayoutRequestUserPanes {
	s.SourceType = &v
	return s
}

func (s *UpdateMPULayoutRequestUserPanes) SetPaneId(v int32) *UpdateMPULayoutRequestUserPanes {
	s.PaneId = &v
	return s
}

type UpdateMPULayoutRequestUserPanesImages struct {
	Width   *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height  *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Y       *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	Url     *string  `json:"Url,omitempty" xml:"Url,omitempty"`
	Display *int32   `json:"Display,omitempty" xml:"Display,omitempty"`
	ZOrder  *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
	X       *float32 `json:"X,omitempty" xml:"X,omitempty"`
}

func (s UpdateMPULayoutRequestUserPanesImages) String() string {
	return tea.Prettify(s)
}

func (s UpdateMPULayoutRequestUserPanesImages) GoString() string {
	return s.String()
}

func (s *UpdateMPULayoutRequestUserPanesImages) SetWidth(v float32) *UpdateMPULayoutRequestUserPanesImages {
	s.Width = &v
	return s
}

func (s *UpdateMPULayoutRequestUserPanesImages) SetHeight(v float32) *UpdateMPULayoutRequestUserPanesImages {
	s.Height = &v
	return s
}

func (s *UpdateMPULayoutRequestUserPanesImages) SetY(v float32) *UpdateMPULayoutRequestUserPanesImages {
	s.Y = &v
	return s
}

func (s *UpdateMPULayoutRequestUserPanesImages) SetUrl(v string) *UpdateMPULayoutRequestUserPanesImages {
	s.Url = &v
	return s
}

func (s *UpdateMPULayoutRequestUserPanesImages) SetDisplay(v int32) *UpdateMPULayoutRequestUserPanesImages {
	s.Display = &v
	return s
}

func (s *UpdateMPULayoutRequestUserPanesImages) SetZOrder(v int32) *UpdateMPULayoutRequestUserPanesImages {
	s.ZOrder = &v
	return s
}

func (s *UpdateMPULayoutRequestUserPanesImages) SetX(v float32) *UpdateMPULayoutRequestUserPanesImages {
	s.X = &v
	return s
}

type UpdateMPULayoutRequestUserPanesTexts struct {
	FontType  *int32   `json:"FontType,omitempty" xml:"FontType,omitempty"`
	FontColor *int32   `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	Y         *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	Text      *string  `json:"Text,omitempty" xml:"Text,omitempty"`
	ZOrder    *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
	FontSize  *int32   `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	X         *float32 `json:"X,omitempty" xml:"X,omitempty"`
}

func (s UpdateMPULayoutRequestUserPanesTexts) String() string {
	return tea.Prettify(s)
}

func (s UpdateMPULayoutRequestUserPanesTexts) GoString() string {
	return s.String()
}

func (s *UpdateMPULayoutRequestUserPanesTexts) SetFontType(v int32) *UpdateMPULayoutRequestUserPanesTexts {
	s.FontType = &v
	return s
}

func (s *UpdateMPULayoutRequestUserPanesTexts) SetFontColor(v int32) *UpdateMPULayoutRequestUserPanesTexts {
	s.FontColor = &v
	return s
}

func (s *UpdateMPULayoutRequestUserPanesTexts) SetY(v float32) *UpdateMPULayoutRequestUserPanesTexts {
	s.Y = &v
	return s
}

func (s *UpdateMPULayoutRequestUserPanesTexts) SetText(v string) *UpdateMPULayoutRequestUserPanesTexts {
	s.Text = &v
	return s
}

func (s *UpdateMPULayoutRequestUserPanesTexts) SetZOrder(v int32) *UpdateMPULayoutRequestUserPanesTexts {
	s.ZOrder = &v
	return s
}

func (s *UpdateMPULayoutRequestUserPanesTexts) SetFontSize(v int32) *UpdateMPULayoutRequestUserPanesTexts {
	s.FontSize = &v
	return s
}

func (s *UpdateMPULayoutRequestUserPanesTexts) SetX(v float32) *UpdateMPULayoutRequestUserPanesTexts {
	s.X = &v
	return s
}

type UpdateMPULayoutRequestBackgrounds struct {
	Width   *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height  *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Y       *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	Url     *string  `json:"Url,omitempty" xml:"Url,omitempty"`
	Display *int32   `json:"Display,omitempty" xml:"Display,omitempty"`
	ZOrder  *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
	X       *float32 `json:"X,omitempty" xml:"X,omitempty"`
}

func (s UpdateMPULayoutRequestBackgrounds) String() string {
	return tea.Prettify(s)
}

func (s UpdateMPULayoutRequestBackgrounds) GoString() string {
	return s.String()
}

func (s *UpdateMPULayoutRequestBackgrounds) SetWidth(v float32) *UpdateMPULayoutRequestBackgrounds {
	s.Width = &v
	return s
}

func (s *UpdateMPULayoutRequestBackgrounds) SetHeight(v float32) *UpdateMPULayoutRequestBackgrounds {
	s.Height = &v
	return s
}

func (s *UpdateMPULayoutRequestBackgrounds) SetY(v float32) *UpdateMPULayoutRequestBackgrounds {
	s.Y = &v
	return s
}

func (s *UpdateMPULayoutRequestBackgrounds) SetUrl(v string) *UpdateMPULayoutRequestBackgrounds {
	s.Url = &v
	return s
}

func (s *UpdateMPULayoutRequestBackgrounds) SetDisplay(v int32) *UpdateMPULayoutRequestBackgrounds {
	s.Display = &v
	return s
}

func (s *UpdateMPULayoutRequestBackgrounds) SetZOrder(v int32) *UpdateMPULayoutRequestBackgrounds {
	s.ZOrder = &v
	return s
}

func (s *UpdateMPULayoutRequestBackgrounds) SetX(v float32) *UpdateMPULayoutRequestBackgrounds {
	s.X = &v
	return s
}

type UpdateMPULayoutRequestWatermarks struct {
	Alpha   *float32 `json:"Alpha,omitempty" xml:"Alpha,omitempty"`
	Width   *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height  *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Y       *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	Url     *string  `json:"Url,omitempty" xml:"Url,omitempty"`
	Display *int32   `json:"Display,omitempty" xml:"Display,omitempty"`
	ZOrder  *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
	X       *float32 `json:"X,omitempty" xml:"X,omitempty"`
}

func (s UpdateMPULayoutRequestWatermarks) String() string {
	return tea.Prettify(s)
}

func (s UpdateMPULayoutRequestWatermarks) GoString() string {
	return s.String()
}

func (s *UpdateMPULayoutRequestWatermarks) SetAlpha(v float32) *UpdateMPULayoutRequestWatermarks {
	s.Alpha = &v
	return s
}

func (s *UpdateMPULayoutRequestWatermarks) SetWidth(v float32) *UpdateMPULayoutRequestWatermarks {
	s.Width = &v
	return s
}

func (s *UpdateMPULayoutRequestWatermarks) SetHeight(v float32) *UpdateMPULayoutRequestWatermarks {
	s.Height = &v
	return s
}

func (s *UpdateMPULayoutRequestWatermarks) SetY(v float32) *UpdateMPULayoutRequestWatermarks {
	s.Y = &v
	return s
}

func (s *UpdateMPULayoutRequestWatermarks) SetUrl(v string) *UpdateMPULayoutRequestWatermarks {
	s.Url = &v
	return s
}

func (s *UpdateMPULayoutRequestWatermarks) SetDisplay(v int32) *UpdateMPULayoutRequestWatermarks {
	s.Display = &v
	return s
}

func (s *UpdateMPULayoutRequestWatermarks) SetZOrder(v int32) *UpdateMPULayoutRequestWatermarks {
	s.ZOrder = &v
	return s
}

func (s *UpdateMPULayoutRequestWatermarks) SetX(v float32) *UpdateMPULayoutRequestWatermarks {
	s.X = &v
	return s
}

type UpdateMPULayoutRequestClockWidgets struct {
	FontType  *int32   `json:"FontType,omitempty" xml:"FontType,omitempty"`
	FontColor *int32   `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	Y         *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	ZOrder    *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
	X         *float32 `json:"X,omitempty" xml:"X,omitempty"`
	FontSize  *int32   `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
}

func (s UpdateMPULayoutRequestClockWidgets) String() string {
	return tea.Prettify(s)
}

func (s UpdateMPULayoutRequestClockWidgets) GoString() string {
	return s.String()
}

func (s *UpdateMPULayoutRequestClockWidgets) SetFontType(v int32) *UpdateMPULayoutRequestClockWidgets {
	s.FontType = &v
	return s
}

func (s *UpdateMPULayoutRequestClockWidgets) SetFontColor(v int32) *UpdateMPULayoutRequestClockWidgets {
	s.FontColor = &v
	return s
}

func (s *UpdateMPULayoutRequestClockWidgets) SetY(v float32) *UpdateMPULayoutRequestClockWidgets {
	s.Y = &v
	return s
}

func (s *UpdateMPULayoutRequestClockWidgets) SetZOrder(v int32) *UpdateMPULayoutRequestClockWidgets {
	s.ZOrder = &v
	return s
}

func (s *UpdateMPULayoutRequestClockWidgets) SetX(v float32) *UpdateMPULayoutRequestClockWidgets {
	s.X = &v
	return s
}

func (s *UpdateMPULayoutRequestClockWidgets) SetFontSize(v int32) *UpdateMPULayoutRequestClockWidgets {
	s.FontSize = &v
	return s
}

type UpdateMPULayoutResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateMPULayoutResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateMPULayoutResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMPULayoutResponseBody) SetRequestId(v string) *UpdateMPULayoutResponseBody {
	s.RequestId = &v
	return s
}

type UpdateMPULayoutResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateMPULayoutResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateMPULayoutResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateMPULayoutResponse) GoString() string {
	return s.String()
}

func (s *UpdateMPULayoutResponse) SetHeaders(v map[string]*string) *UpdateMPULayoutResponse {
	s.Headers = v
	return s
}

func (s *UpdateMPULayoutResponse) SetBody(v *UpdateMPULayoutResponseBody) *UpdateMPULayoutResponse {
	s.Body = v
	return s
}

type UpdateRecordTaskRequest struct {
	OwnerId      *int64                              `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog      *string                             `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId        *string                             `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ChannelId    *string                             `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	TaskId       *string                             `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TemplateId   *string                             `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	SubSpecUsers []*string                           `json:"SubSpecUsers,omitempty" xml:"SubSpecUsers,omitempty" type:"Repeated"`
	UserPanes    []*UpdateRecordTaskRequestUserPanes `json:"UserPanes,omitempty" xml:"UserPanes,omitempty" type:"Repeated"`
}

func (s UpdateRecordTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecordTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateRecordTaskRequest) SetOwnerId(v int64) *UpdateRecordTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateRecordTaskRequest) SetShowLog(v string) *UpdateRecordTaskRequest {
	s.ShowLog = &v
	return s
}

func (s *UpdateRecordTaskRequest) SetAppId(v string) *UpdateRecordTaskRequest {
	s.AppId = &v
	return s
}

func (s *UpdateRecordTaskRequest) SetChannelId(v string) *UpdateRecordTaskRequest {
	s.ChannelId = &v
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

func (s *UpdateRecordTaskRequest) SetSubSpecUsers(v []*string) *UpdateRecordTaskRequest {
	s.SubSpecUsers = v
	return s
}

func (s *UpdateRecordTaskRequest) SetUserPanes(v []*UpdateRecordTaskRequestUserPanes) *UpdateRecordTaskRequest {
	s.UserPanes = v
	return s
}

type UpdateRecordTaskRequestUserPanes struct {
	Images     []*UpdateRecordTaskRequestUserPanesImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	UserId     *string                                   `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Texts      []*UpdateRecordTaskRequestUserPanesTexts  `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
	SourceType *string                                   `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	PaneId     *int32                                    `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
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

func (s *UpdateRecordTaskRequestUserPanes) SetUserId(v string) *UpdateRecordTaskRequestUserPanes {
	s.UserId = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanes) SetTexts(v []*UpdateRecordTaskRequestUserPanesTexts) *UpdateRecordTaskRequestUserPanes {
	s.Texts = v
	return s
}

func (s *UpdateRecordTaskRequestUserPanes) SetSourceType(v string) *UpdateRecordTaskRequestUserPanes {
	s.SourceType = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanes) SetPaneId(v int32) *UpdateRecordTaskRequestUserPanes {
	s.PaneId = &v
	return s
}

type UpdateRecordTaskRequestUserPanesImages struct {
	Width   *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height  *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Y       *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	Url     *string  `json:"Url,omitempty" xml:"Url,omitempty"`
	Display *int32   `json:"Display,omitempty" xml:"Display,omitempty"`
	ZOrder  *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
	X       *float32 `json:"X,omitempty" xml:"X,omitempty"`
}

func (s UpdateRecordTaskRequestUserPanesImages) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecordTaskRequestUserPanesImages) GoString() string {
	return s.String()
}

func (s *UpdateRecordTaskRequestUserPanesImages) SetWidth(v float32) *UpdateRecordTaskRequestUserPanesImages {
	s.Width = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanesImages) SetHeight(v float32) *UpdateRecordTaskRequestUserPanesImages {
	s.Height = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanesImages) SetY(v float32) *UpdateRecordTaskRequestUserPanesImages {
	s.Y = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanesImages) SetUrl(v string) *UpdateRecordTaskRequestUserPanesImages {
	s.Url = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanesImages) SetDisplay(v int32) *UpdateRecordTaskRequestUserPanesImages {
	s.Display = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanesImages) SetZOrder(v int32) *UpdateRecordTaskRequestUserPanesImages {
	s.ZOrder = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanesImages) SetX(v float32) *UpdateRecordTaskRequestUserPanesImages {
	s.X = &v
	return s
}

type UpdateRecordTaskRequestUserPanesTexts struct {
	FontType  *int32   `json:"FontType,omitempty" xml:"FontType,omitempty"`
	FontColor *int32   `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	Y         *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	Text      *string  `json:"Text,omitempty" xml:"Text,omitempty"`
	ZOrder    *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
	FontSize  *int32   `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	X         *float32 `json:"X,omitempty" xml:"X,omitempty"`
}

func (s UpdateRecordTaskRequestUserPanesTexts) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecordTaskRequestUserPanesTexts) GoString() string {
	return s.String()
}

func (s *UpdateRecordTaskRequestUserPanesTexts) SetFontType(v int32) *UpdateRecordTaskRequestUserPanesTexts {
	s.FontType = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanesTexts) SetFontColor(v int32) *UpdateRecordTaskRequestUserPanesTexts {
	s.FontColor = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanesTexts) SetY(v float32) *UpdateRecordTaskRequestUserPanesTexts {
	s.Y = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanesTexts) SetText(v string) *UpdateRecordTaskRequestUserPanesTexts {
	s.Text = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanesTexts) SetZOrder(v int32) *UpdateRecordTaskRequestUserPanesTexts {
	s.ZOrder = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanesTexts) SetFontSize(v int32) *UpdateRecordTaskRequestUserPanesTexts {
	s.FontSize = &v
	return s
}

func (s *UpdateRecordTaskRequestUserPanesTexts) SetX(v float32) *UpdateRecordTaskRequestUserPanesTexts {
	s.X = &v
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
	OwnerId           *int64                                     `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ShowLog           *string                                    `json:"ShowLog,omitempty" xml:"ShowLog,omitempty"`
	AppId             *string                                    `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Name              *string                                    `json:"Name,omitempty" xml:"Name,omitempty"`
	TemplateId        *string                                    `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TaskProfile       *string                                    `json:"TaskProfile,omitempty" xml:"TaskProfile,omitempty"`
	MediaEncode       *int32                                     `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	BackgroundColor   *int32                                     `json:"BackgroundColor,omitempty" xml:"BackgroundColor,omitempty"`
	OssBucket         *string                                    `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssFilePrefix     *string                                    `json:"OssFilePrefix,omitempty" xml:"OssFilePrefix,omitempty"`
	MnsQueue          *string                                    `json:"MnsQueue,omitempty" xml:"MnsQueue,omitempty"`
	HttpCallbackUrl   *string                                    `json:"HttpCallbackUrl,omitempty" xml:"HttpCallbackUrl,omitempty"`
	FileSplitInterval *int32                                     `json:"FileSplitInterval,omitempty" xml:"FileSplitInterval,omitempty"`
	DelayStopTime     *int32                                     `json:"DelayStopTime,omitempty" xml:"DelayStopTime,omitempty"`
	LayoutIds         []*int                                     `json:"LayoutIds,omitempty" xml:"LayoutIds,omitempty" type:"Repeated"`
	Formats           []*string                                  `json:"Formats,omitempty" xml:"Formats,omitempty" type:"Repeated"`
	Backgrounds       []*UpdateRecordTemplateRequestBackgrounds  `json:"Backgrounds,omitempty" xml:"Backgrounds,omitempty" type:"Repeated"`
	Watermarks        []*UpdateRecordTemplateRequestWatermarks   `json:"Watermarks,omitempty" xml:"Watermarks,omitempty" type:"Repeated"`
	ClockWidgets      []*UpdateRecordTemplateRequestClockWidgets `json:"ClockWidgets,omitempty" xml:"ClockWidgets,omitempty" type:"Repeated"`
}

func (s UpdateRecordTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecordTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateRecordTemplateRequest) SetOwnerId(v int64) *UpdateRecordTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateRecordTemplateRequest) SetShowLog(v string) *UpdateRecordTemplateRequest {
	s.ShowLog = &v
	return s
}

func (s *UpdateRecordTemplateRequest) SetAppId(v string) *UpdateRecordTemplateRequest {
	s.AppId = &v
	return s
}

func (s *UpdateRecordTemplateRequest) SetName(v string) *UpdateRecordTemplateRequest {
	s.Name = &v
	return s
}

func (s *UpdateRecordTemplateRequest) SetTemplateId(v string) *UpdateRecordTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateRecordTemplateRequest) SetTaskProfile(v string) *UpdateRecordTemplateRequest {
	s.TaskProfile = &v
	return s
}

func (s *UpdateRecordTemplateRequest) SetMediaEncode(v int32) *UpdateRecordTemplateRequest {
	s.MediaEncode = &v
	return s
}

func (s *UpdateRecordTemplateRequest) SetBackgroundColor(v int32) *UpdateRecordTemplateRequest {
	s.BackgroundColor = &v
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

func (s *UpdateRecordTemplateRequest) SetMnsQueue(v string) *UpdateRecordTemplateRequest {
	s.MnsQueue = &v
	return s
}

func (s *UpdateRecordTemplateRequest) SetHttpCallbackUrl(v string) *UpdateRecordTemplateRequest {
	s.HttpCallbackUrl = &v
	return s
}

func (s *UpdateRecordTemplateRequest) SetFileSplitInterval(v int32) *UpdateRecordTemplateRequest {
	s.FileSplitInterval = &v
	return s
}

func (s *UpdateRecordTemplateRequest) SetDelayStopTime(v int32) *UpdateRecordTemplateRequest {
	s.DelayStopTime = &v
	return s
}

func (s *UpdateRecordTemplateRequest) SetLayoutIds(v []*int) *UpdateRecordTemplateRequest {
	s.LayoutIds = v
	return s
}

func (s *UpdateRecordTemplateRequest) SetFormats(v []*string) *UpdateRecordTemplateRequest {
	s.Formats = v
	return s
}

func (s *UpdateRecordTemplateRequest) SetBackgrounds(v []*UpdateRecordTemplateRequestBackgrounds) *UpdateRecordTemplateRequest {
	s.Backgrounds = v
	return s
}

func (s *UpdateRecordTemplateRequest) SetWatermarks(v []*UpdateRecordTemplateRequestWatermarks) *UpdateRecordTemplateRequest {
	s.Watermarks = v
	return s
}

func (s *UpdateRecordTemplateRequest) SetClockWidgets(v []*UpdateRecordTemplateRequestClockWidgets) *UpdateRecordTemplateRequest {
	s.ClockWidgets = v
	return s
}

type UpdateRecordTemplateRequestBackgrounds struct {
	Width   *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height  *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Y       *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	Url     *string  `json:"Url,omitempty" xml:"Url,omitempty"`
	Display *int32   `json:"Display,omitempty" xml:"Display,omitempty"`
	ZOrder  *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
	X       *float32 `json:"X,omitempty" xml:"X,omitempty"`
}

func (s UpdateRecordTemplateRequestBackgrounds) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecordTemplateRequestBackgrounds) GoString() string {
	return s.String()
}

func (s *UpdateRecordTemplateRequestBackgrounds) SetWidth(v float32) *UpdateRecordTemplateRequestBackgrounds {
	s.Width = &v
	return s
}

func (s *UpdateRecordTemplateRequestBackgrounds) SetHeight(v float32) *UpdateRecordTemplateRequestBackgrounds {
	s.Height = &v
	return s
}

func (s *UpdateRecordTemplateRequestBackgrounds) SetY(v float32) *UpdateRecordTemplateRequestBackgrounds {
	s.Y = &v
	return s
}

func (s *UpdateRecordTemplateRequestBackgrounds) SetUrl(v string) *UpdateRecordTemplateRequestBackgrounds {
	s.Url = &v
	return s
}

func (s *UpdateRecordTemplateRequestBackgrounds) SetDisplay(v int32) *UpdateRecordTemplateRequestBackgrounds {
	s.Display = &v
	return s
}

func (s *UpdateRecordTemplateRequestBackgrounds) SetZOrder(v int32) *UpdateRecordTemplateRequestBackgrounds {
	s.ZOrder = &v
	return s
}

func (s *UpdateRecordTemplateRequestBackgrounds) SetX(v float32) *UpdateRecordTemplateRequestBackgrounds {
	s.X = &v
	return s
}

type UpdateRecordTemplateRequestWatermarks struct {
	Alpha   *float32 `json:"Alpha,omitempty" xml:"Alpha,omitempty"`
	Width   *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	Height  *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	Y       *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	Url     *string  `json:"Url,omitempty" xml:"Url,omitempty"`
	Display *int32   `json:"Display,omitempty" xml:"Display,omitempty"`
	ZOrder  *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
	X       *float32 `json:"X,omitempty" xml:"X,omitempty"`
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

func (s *UpdateRecordTemplateRequestWatermarks) SetWidth(v float32) *UpdateRecordTemplateRequestWatermarks {
	s.Width = &v
	return s
}

func (s *UpdateRecordTemplateRequestWatermarks) SetHeight(v float32) *UpdateRecordTemplateRequestWatermarks {
	s.Height = &v
	return s
}

func (s *UpdateRecordTemplateRequestWatermarks) SetY(v float32) *UpdateRecordTemplateRequestWatermarks {
	s.Y = &v
	return s
}

func (s *UpdateRecordTemplateRequestWatermarks) SetUrl(v string) *UpdateRecordTemplateRequestWatermarks {
	s.Url = &v
	return s
}

func (s *UpdateRecordTemplateRequestWatermarks) SetDisplay(v int32) *UpdateRecordTemplateRequestWatermarks {
	s.Display = &v
	return s
}

func (s *UpdateRecordTemplateRequestWatermarks) SetZOrder(v int32) *UpdateRecordTemplateRequestWatermarks {
	s.ZOrder = &v
	return s
}

func (s *UpdateRecordTemplateRequestWatermarks) SetX(v float32) *UpdateRecordTemplateRequestWatermarks {
	s.X = &v
	return s
}

type UpdateRecordTemplateRequestClockWidgets struct {
	FontType  *int32   `json:"FontType,omitempty" xml:"FontType,omitempty"`
	FontColor *int32   `json:"FontColor,omitempty" xml:"FontColor,omitempty"`
	Y         *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	ZOrder    *int32   `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
	X         *float32 `json:"X,omitempty" xml:"X,omitempty"`
	FontSize  *int32   `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
}

func (s UpdateRecordTemplateRequestClockWidgets) String() string {
	return tea.Prettify(s)
}

func (s UpdateRecordTemplateRequestClockWidgets) GoString() string {
	return s.String()
}

func (s *UpdateRecordTemplateRequestClockWidgets) SetFontType(v int32) *UpdateRecordTemplateRequestClockWidgets {
	s.FontType = &v
	return s
}

func (s *UpdateRecordTemplateRequestClockWidgets) SetFontColor(v int32) *UpdateRecordTemplateRequestClockWidgets {
	s.FontColor = &v
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

func (s *UpdateRecordTemplateRequestClockWidgets) SetX(v float32) *UpdateRecordTemplateRequestClockWidgets {
	s.X = &v
	return s
}

func (s *UpdateRecordTemplateRequestClockWidgets) SetFontSize(v int32) *UpdateRecordTemplateRequestClockWidgets {
	s.FontSize = &v
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

func (client *Client) CreateChannelWithOptions(request *CreateChannelRequest, runtime *util.RuntimeOptions) (_result *CreateChannelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateChannelResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateChannel"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateChannel(request *CreateChannelRequest) (_result *CreateChannelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateChannelResponse{}
	_body, _err := client.CreateChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateConferenceWithOptions(request *CreateConferenceRequest, runtime *util.RuntimeOptions) (_result *CreateConferenceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateConferenceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateConference"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateConference(request *CreateConferenceRequest) (_result *CreateConferenceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateConferenceResponse{}
	_body, _err := client.CreateConferenceWithOptions(request, runtime)
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

func (client *Client) CreateMPURuleWithOptions(request *CreateMPURuleRequest, runtime *util.RuntimeOptions) (_result *CreateMPURuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateMPURuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateMPURule"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMPURule(request *CreateMPURuleRequest) (_result *CreateMPURuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMPURuleResponse{}
	_body, _err := client.CreateMPURuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateServiceLinkedRoleForRtcWithOptions(request *CreateServiceLinkedRoleForRtcRequest, runtime *util.RuntimeOptions) (_result *CreateServiceLinkedRoleForRtcResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateServiceLinkedRoleForRtcResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateServiceLinkedRoleForRtc"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateServiceLinkedRoleForRtc(request *CreateServiceLinkedRoleForRtcRequest) (_result *CreateServiceLinkedRoleForRtcResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServiceLinkedRoleForRtcResponse{}
	_body, _err := client.CreateServiceLinkedRoleForRtcWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSubscribeWithOptions(request *CreateSubscribeRequest, runtime *util.RuntimeOptions) (_result *CreateSubscribeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateSubscribeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateSubscribe"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSubscribe(request *CreateSubscribeRequest) (_result *CreateSubscribeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSubscribeResponse{}
	_body, _err := client.CreateSubscribeWithOptions(request, runtime)
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

func (client *Client) DeleteChannelWithOptions(request *DeleteChannelRequest, runtime *util.RuntimeOptions) (_result *DeleteChannelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteChannelResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteChannel"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteChannel(request *DeleteChannelRequest) (_result *DeleteChannelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteChannelResponse{}
	_body, _err := client.DeleteChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteConferenceWithOptions(request *DeleteConferenceRequest, runtime *util.RuntimeOptions) (_result *DeleteConferenceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteConferenceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteConference"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteConference(request *DeleteConferenceRequest) (_result *DeleteConferenceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteConferenceResponse{}
	_body, _err := client.DeleteConferenceWithOptions(request, runtime)
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

func (client *Client) DeleteMPURuleWithOptions(request *DeleteMPURuleRequest, runtime *util.RuntimeOptions) (_result *DeleteMPURuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteMPURuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteMPURule"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMPURule(request *DeleteMPURuleRequest) (_result *DeleteMPURuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMPURuleResponse{}
	_body, _err := client.DeleteMPURuleWithOptions(request, runtime)
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

func (client *Client) DeleteSubscribeWithOptions(request *DeleteSubscribeRequest, runtime *util.RuntimeOptions) (_result *DeleteSubscribeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteSubscribeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteSubscribe"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSubscribe(request *DeleteSubscribeRequest) (_result *DeleteSubscribeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSubscribeResponse{}
	_body, _err := client.DeleteSubscribeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAppsWithOptions(request *DescribeAppsRequest, runtime *util.RuntimeOptions) (_result *DescribeAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeAppsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeApps"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApps(request *DescribeAppsRequest) (_result *DescribeAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAppsResponse{}
	_body, _err := client.DescribeAppsWithOptions(request, runtime)
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

func (client *Client) DescribeConferenceAuthInfoWithOptions(request *DescribeConferenceAuthInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeConferenceAuthInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeConferenceAuthInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeConferenceAuthInfo"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeConferenceAuthInfo(request *DescribeConferenceAuthInfoRequest) (_result *DescribeConferenceAuthInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeConferenceAuthInfoResponse{}
	_body, _err := client.DescribeConferenceAuthInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMPULayoutInfoWithOptions(request *DescribeMPULayoutInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeMPULayoutInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMPULayoutInfoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMPULayoutInfo"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMPULayoutInfo(request *DescribeMPULayoutInfoRequest) (_result *DescribeMPULayoutInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMPULayoutInfoResponse{}
	_body, _err := client.DescribeMPULayoutInfoWithOptions(request, runtime)
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

func (client *Client) DescribeMPULayoutListWithOptions(request *DescribeMPULayoutListRequest, runtime *util.RuntimeOptions) (_result *DescribeMPULayoutListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMPULayoutListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMPULayoutList"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMPULayoutList(request *DescribeMPULayoutListRequest) (_result *DescribeMPULayoutListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMPULayoutListResponse{}
	_body, _err := client.DescribeMPULayoutListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMPURuleWithOptions(request *DescribeMPURuleRequest, runtime *util.RuntimeOptions) (_result *DescribeMPURuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMPURuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMPURule"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMPURule(request *DescribeMPURuleRequest) (_result *DescribeMPURuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMPURuleResponse{}
	_body, _err := client.DescribeMPURuleWithOptions(request, runtime)
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

func (client *Client) DescribeRTCAppKeyWithOptions(request *DescribeRTCAppKeyRequest, runtime *util.RuntimeOptions) (_result *DescribeRTCAppKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRTCAppKeyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRTCAppKey"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRTCAppKey(request *DescribeRTCAppKeyRequest) (_result *DescribeRTCAppKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRTCAppKeyResponse{}
	_body, _err := client.DescribeRTCAppKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRtcChannelCntDataWithOptions(request *DescribeRtcChannelCntDataRequest, runtime *util.RuntimeOptions) (_result *DescribeRtcChannelCntDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRtcChannelCntDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRtcChannelCntData"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRtcChannelCntData(request *DescribeRtcChannelCntDataRequest) (_result *DescribeRtcChannelCntDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRtcChannelCntDataResponse{}
	_body, _err := client.DescribeRtcChannelCntDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRtcChannelDetailWithOptions(request *DescribeRtcChannelDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeRtcChannelDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRtcChannelDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRtcChannelDetail"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRtcChannelDetail(request *DescribeRtcChannelDetailRequest) (_result *DescribeRtcChannelDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRtcChannelDetailResponse{}
	_body, _err := client.DescribeRtcChannelDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRtcChannelListWithOptions(request *DescribeRtcChannelListRequest, runtime *util.RuntimeOptions) (_result *DescribeRtcChannelListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRtcChannelListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRtcChannelList"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRtcChannelList(request *DescribeRtcChannelListRequest) (_result *DescribeRtcChannelListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRtcChannelListResponse{}
	_body, _err := client.DescribeRtcChannelListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRtcChannelMetricWithOptions(request *DescribeRtcChannelMetricRequest, runtime *util.RuntimeOptions) (_result *DescribeRtcChannelMetricResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRtcChannelMetricResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRtcChannelMetric"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRtcChannelMetric(request *DescribeRtcChannelMetricRequest) (_result *DescribeRtcChannelMetricResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRtcChannelMetricResponse{}
	_body, _err := client.DescribeRtcChannelMetricWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRtcChannelMetricsWithOptions(request *DescribeRtcChannelMetricsRequest, runtime *util.RuntimeOptions) (_result *DescribeRtcChannelMetricsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRtcChannelMetricsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRtcChannelMetrics"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRtcChannelMetrics(request *DescribeRtcChannelMetricsRequest) (_result *DescribeRtcChannelMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRtcChannelMetricsResponse{}
	_body, _err := client.DescribeRtcChannelMetricsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRtcChannelsWithOptions(request *DescribeRtcChannelsRequest, runtime *util.RuntimeOptions) (_result *DescribeRtcChannelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRtcChannelsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRtcChannels"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRtcChannels(request *DescribeRtcChannelsRequest) (_result *DescribeRtcChannelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRtcChannelsResponse{}
	_body, _err := client.DescribeRtcChannelsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRtcChannelUserListWithOptions(request *DescribeRtcChannelUserListRequest, runtime *util.RuntimeOptions) (_result *DescribeRtcChannelUserListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRtcChannelUserListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRtcChannelUserList"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRtcChannelUserList(request *DescribeRtcChannelUserListRequest) (_result *DescribeRtcChannelUserListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRtcChannelUserListResponse{}
	_body, _err := client.DescribeRtcChannelUserListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRtcDurationDataWithOptions(request *DescribeRtcDurationDataRequest, runtime *util.RuntimeOptions) (_result *DescribeRtcDurationDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRtcDurationDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRtcDurationData"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRtcDurationData(request *DescribeRtcDurationDataRequest) (_result *DescribeRtcDurationDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRtcDurationDataResponse{}
	_body, _err := client.DescribeRtcDurationDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRtcPeakChannelCntDataWithOptions(request *DescribeRtcPeakChannelCntDataRequest, runtime *util.RuntimeOptions) (_result *DescribeRtcPeakChannelCntDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRtcPeakChannelCntDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRtcPeakChannelCntData"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRtcPeakChannelCntData(request *DescribeRtcPeakChannelCntDataRequest) (_result *DescribeRtcPeakChannelCntDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRtcPeakChannelCntDataResponse{}
	_body, _err := client.DescribeRtcPeakChannelCntDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRtcPeakUserCntDataWithOptions(request *DescribeRtcPeakUserCntDataRequest, runtime *util.RuntimeOptions) (_result *DescribeRtcPeakUserCntDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRtcPeakUserCntDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRtcPeakUserCntData"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRtcPeakUserCntData(request *DescribeRtcPeakUserCntDataRequest) (_result *DescribeRtcPeakUserCntDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRtcPeakUserCntDataResponse{}
	_body, _err := client.DescribeRtcPeakUserCntDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRtcScaleWithOptions(request *DescribeRtcScaleRequest, runtime *util.RuntimeOptions) (_result *DescribeRtcScaleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRtcScaleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRtcScale"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRtcScale(request *DescribeRtcScaleRequest) (_result *DescribeRtcScaleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRtcScaleResponse{}
	_body, _err := client.DescribeRtcScaleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRtcScaleDetailWithOptions(request *DescribeRtcScaleDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeRtcScaleDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRtcScaleDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRtcScaleDetail"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRtcScaleDetail(request *DescribeRtcScaleDetailRequest) (_result *DescribeRtcScaleDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRtcScaleDetailResponse{}
	_body, _err := client.DescribeRtcScaleDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRtcUserCntDataWithOptions(request *DescribeRtcUserCntDataRequest, runtime *util.RuntimeOptions) (_result *DescribeRtcUserCntDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRtcUserCntDataResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRtcUserCntData"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRtcUserCntData(request *DescribeRtcUserCntDataRequest) (_result *DescribeRtcUserCntDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRtcUserCntDataResponse{}
	_body, _err := client.DescribeRtcUserCntDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRtcUserEventsWithOptions(request *DescribeRtcUserEventsRequest, runtime *util.RuntimeOptions) (_result *DescribeRtcUserEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRtcUserEventsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRtcUserEvents"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRtcUserEvents(request *DescribeRtcUserEventsRequest) (_result *DescribeRtcUserEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRtcUserEventsResponse{}
	_body, _err := client.DescribeRtcUserEventsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRtcUserListWithOptions(request *DescribeRtcUserListRequest, runtime *util.RuntimeOptions) (_result *DescribeRtcUserListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeRtcUserListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeRtcUserList"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRtcUserList(request *DescribeRtcUserListRequest) (_result *DescribeRtcUserListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRtcUserListResponse{}
	_body, _err := client.DescribeRtcUserListWithOptions(request, runtime)
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

func (client *Client) DisableMPURuleWithOptions(request *DisableMPURuleRequest, runtime *util.RuntimeOptions) (_result *DisableMPURuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DisableMPURuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DisableMPURule"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableMPURule(request *DisableMPURuleRequest) (_result *DisableMPURuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DisableMPURuleResponse{}
	_body, _err := client.DisableMPURuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableMPURuleWithOptions(request *EnableMPURuleRequest, runtime *util.RuntimeOptions) (_result *EnableMPURuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EnableMPURuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EnableMPURule"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableMPURule(request *EnableMPURuleRequest) (_result *EnableMPURuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EnableMPURuleResponse{}
	_body, _err := client.EnableMPURuleWithOptions(request, runtime)
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

func (client *Client) ModifyAppWithOptions(request *ModifyAppRequest, runtime *util.RuntimeOptions) (_result *ModifyAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyAppResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyApp"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyApp(request *ModifyAppRequest) (_result *ModifyAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAppResponse{}
	_body, _err := client.ModifyAppWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyConferenceWithOptions(request *ModifyConferenceRequest, runtime *util.RuntimeOptions) (_result *ModifyConferenceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyConferenceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyConference"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyConference(request *ModifyConferenceRequest) (_result *ModifyConferenceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyConferenceResponse{}
	_body, _err := client.ModifyConferenceWithOptions(request, runtime)
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

func (client *Client) MuteAudioWithOptions(request *MuteAudioRequest, runtime *util.RuntimeOptions) (_result *MuteAudioResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &MuteAudioResponse{}
	_body, _err := client.DoRPCRequest(tea.String("MuteAudio"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MuteAudio(request *MuteAudioRequest) (_result *MuteAudioResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MuteAudioResponse{}
	_body, _err := client.MuteAudioWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MuteAudioAllWithOptions(request *MuteAudioAllRequest, runtime *util.RuntimeOptions) (_result *MuteAudioAllResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &MuteAudioAllResponse{}
	_body, _err := client.DoRPCRequest(tea.String("MuteAudioAll"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MuteAudioAll(request *MuteAudioAllRequest) (_result *MuteAudioAllResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MuteAudioAllResponse{}
	_body, _err := client.MuteAudioAllWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReceiveNotifyWithOptions(request *ReceiveNotifyRequest, runtime *util.RuntimeOptions) (_result *ReceiveNotifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReceiveNotifyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReceiveNotify"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReceiveNotify(request *ReceiveNotifyRequest) (_result *ReceiveNotifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReceiveNotifyResponse{}
	_body, _err := client.ReceiveNotifyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveParticipantsWithOptions(request *RemoveParticipantsRequest, runtime *util.RuntimeOptions) (_result *RemoveParticipantsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveParticipantsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveParticipants"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveParticipants(request *RemoveParticipantsRequest) (_result *RemoveParticipantsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveParticipantsResponse{}
	_body, _err := client.RemoveParticipantsWithOptions(request, runtime)
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

func (client *Client) SetChannelPropertyWithOptions(request *SetChannelPropertyRequest, runtime *util.RuntimeOptions) (_result *SetChannelPropertyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetChannelPropertyResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetChannelProperty"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetChannelProperty(request *SetChannelPropertyRequest) (_result *SetChannelPropertyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetChannelPropertyResponse{}
	_body, _err := client.SetChannelPropertyWithOptions(request, runtime)
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

func (client *Client) UnmuteAudioWithOptions(request *UnmuteAudioRequest, runtime *util.RuntimeOptions) (_result *UnmuteAudioResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UnmuteAudioResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UnmuteAudio"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnmuteAudio(request *UnmuteAudioRequest) (_result *UnmuteAudioResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnmuteAudioResponse{}
	_body, _err := client.UnmuteAudioWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnmuteAudioAllWithOptions(request *UnmuteAudioAllRequest, runtime *util.RuntimeOptions) (_result *UnmuteAudioAllResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UnmuteAudioAllResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UnmuteAudioAll"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnmuteAudioAll(request *UnmuteAudioAllRequest) (_result *UnmuteAudioAllResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UnmuteAudioAllResponse{}
	_body, _err := client.UnmuteAudioAllWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateChannelWithOptions(request *UpdateChannelRequest, runtime *util.RuntimeOptions) (_result *UpdateChannelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateChannelResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateChannel"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateChannel(request *UpdateChannelRequest) (_result *UpdateChannelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateChannelResponse{}
	_body, _err := client.UpdateChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateMPULayoutWithOptions(request *UpdateMPULayoutRequest, runtime *util.RuntimeOptions) (_result *UpdateMPULayoutResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpdateMPULayoutResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpdateMPULayout"), tea.String("2018-01-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateMPULayout(request *UpdateMPULayoutRequest) (_result *UpdateMPULayoutResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateMPULayoutResponse{}
	_body, _err := client.UpdateMPULayoutWithOptions(request, runtime)
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
