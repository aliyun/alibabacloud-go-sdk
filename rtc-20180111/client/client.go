// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

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
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 76dasgb****
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddRecordTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *AddRecordTemplateResponse) SetStatusCode(v int32) *AddRecordTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *AddRecordTemplateResponse) SetBody(v *AddRecordTemplateResponseBody) *AddRecordTemplateResponse {
	s.Body = v
	return s
}

type CreateAppStreamingOutTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	StreamingOutTemplate *CreateAppStreamingOutTemplateRequestStreamingOutTemplate `json:"StreamingOutTemplate,omitempty" xml:"StreamingOutTemplate,omitempty" type:"Struct"`
}

func (s CreateAppStreamingOutTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppStreamingOutTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateAppStreamingOutTemplateRequest) SetAppId(v string) *CreateAppStreamingOutTemplateRequest {
	s.AppId = &v
	return s
}

func (s *CreateAppStreamingOutTemplateRequest) SetStreamingOutTemplate(v *CreateAppStreamingOutTemplateRequestStreamingOutTemplate) *CreateAppStreamingOutTemplateRequest {
	s.StreamingOutTemplate = v
	return s
}

type CreateAppStreamingOutTemplateRequestStreamingOutTemplate struct {
	// This parameter is required.
	LayoutIds []*string `json:"LayoutIds,omitempty" xml:"LayoutIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	MediaEncode *int32 `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 模版
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateAppStreamingOutTemplateRequestStreamingOutTemplate) String() string {
	return tea.Prettify(s)
}

func (s CreateAppStreamingOutTemplateRequestStreamingOutTemplate) GoString() string {
	return s.String()
}

func (s *CreateAppStreamingOutTemplateRequestStreamingOutTemplate) SetLayoutIds(v []*string) *CreateAppStreamingOutTemplateRequestStreamingOutTemplate {
	s.LayoutIds = v
	return s
}

func (s *CreateAppStreamingOutTemplateRequestStreamingOutTemplate) SetMediaEncode(v int32) *CreateAppStreamingOutTemplateRequestStreamingOutTemplate {
	s.MediaEncode = &v
	return s
}

func (s *CreateAppStreamingOutTemplateRequestStreamingOutTemplate) SetName(v string) *CreateAppStreamingOutTemplateRequestStreamingOutTemplate {
	s.Name = &v
	return s
}

type CreateAppStreamingOutTemplateShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	StreamingOutTemplateShrink *string `json:"StreamingOutTemplate,omitempty" xml:"StreamingOutTemplate,omitempty"`
}

func (s CreateAppStreamingOutTemplateShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppStreamingOutTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAppStreamingOutTemplateShrinkRequest) SetAppId(v string) *CreateAppStreamingOutTemplateShrinkRequest {
	s.AppId = &v
	return s
}

func (s *CreateAppStreamingOutTemplateShrinkRequest) SetStreamingOutTemplateShrink(v string) *CreateAppStreamingOutTemplateShrinkRequest {
	s.StreamingOutTemplateShrink = &v
	return s
}

type CreateAppStreamingOutTemplateResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 30D41049-D02D-1C21-86AE-B3E5FD805C27
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// bc5v****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateAppStreamingOutTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppStreamingOutTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppStreamingOutTemplateResponseBody) SetRequestId(v string) *CreateAppStreamingOutTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppStreamingOutTemplateResponseBody) SetTemplateId(v string) *CreateAppStreamingOutTemplateResponseBody {
	s.TemplateId = &v
	return s
}

type CreateAppStreamingOutTemplateResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppStreamingOutTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppStreamingOutTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppStreamingOutTemplateResponse) GoString() string {
	return s.String()
}

func (s *CreateAppStreamingOutTemplateResponse) SetHeaders(v map[string]*string) *CreateAppStreamingOutTemplateResponse {
	s.Headers = v
	return s
}

func (s *CreateAppStreamingOutTemplateResponse) SetStatusCode(v int32) *CreateAppStreamingOutTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppStreamingOutTemplateResponse) SetBody(v *CreateAppStreamingOutTemplateResponseBody) *CreateAppStreamingOutTemplateResponse {
	s.Body = v
	return s
}

type CreateAutoLiveStreamRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// http://example.com/callback
	CallBack          *string   `json:"CallBack,omitempty" xml:"CallBack,omitempty"`
	ChannelIdPrefixes []*string `json:"ChannelIdPrefixes,omitempty" xml:"ChannelIdPrefixes,omitempty" type:"Repeated"`
	ChannelIds        []*string `json:"ChannelIds,omitempty" xml:"ChannelIds,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MediaEncode *int32 `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	OwnerId     *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rtmp://${domain}/${app}/${stream}
	PlayDomain *string `json:"PlayDomain,omitempty" xml:"PlayDomain,omitempty"`
	// example:
	//
	// testRule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
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
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 12
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAutoLiveStreamRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateAutoLiveStreamRuleResponse) SetStatusCode(v int32) *CreateAutoLiveStreamRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAutoLiveStreamRuleResponse) SetBody(v *CreateAutoLiveStreamRuleResponseBody) *CreateAutoLiveStreamRuleResponse {
	s.Body = v
	return s
}

type CreateEventSubscribeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 9qb1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://****.com/callback
	CallbackUrl *string `json:"CallbackUrl,omitempty" xml:"CallbackUrl,omitempty"`
	// example:
	//
	// 123333
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ChannelEvent
	Events []*string `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// example:
	//
	// false
	NeedCallbackAuth *bool  `json:"NeedCallbackAuth,omitempty" xml:"NeedCallbackAuth,omitempty"`
	OwnerId          *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	Role *int64 `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// user1
	Users []*string `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
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

func (s *CreateEventSubscribeRequest) SetNeedCallbackAuth(v bool) *CreateEventSubscribeRequest {
	s.NeedCallbackAuth = &v
	return s
}

func (s *CreateEventSubscribeRequest) SetOwnerId(v int64) *CreateEventSubscribeRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateEventSubscribeRequest) SetRole(v int64) *CreateEventSubscribeRequest {
	s.Role = &v
	return s
}

func (s *CreateEventSubscribeRequest) SetUsers(v []*string) *CreateEventSubscribeRequest {
	s.Users = v
	return s
}

type CreateEventSubscribeResponseBody struct {
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// ad53276431c****
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEventSubscribeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateEventSubscribeResponse) SetStatusCode(v int32) *CreateEventSubscribeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEventSubscribeResponse) SetBody(v *CreateEventSubscribeResponseBody) *CreateEventSubscribeResponse {
	s.Body = v
	return s
}

type CreateMPULayoutRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	AudioMixCount *int32 `json:"AudioMixCount,omitempty" xml:"AudioMixCount,omitempty"`
	// example:
	//
	// LayoutName
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	Panes []*CreateMPULayoutRequestPanes `json:"Panes,omitempty" xml:"Panes,omitempty" type:"Repeated"`
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
	// example:
	//
	// 0.25
	Height *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 1
	MajorPane *int32 `json:"MajorPane,omitempty" xml:"MajorPane,omitempty"`
	// example:
	//
	// 0
	PaneId *int32 `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
	// example:
	//
	// 0.25
	Width *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// example:
	//
	// 0.25
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 0.25
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	// example:
	//
	// 0
	ZOrder *int32 `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
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
	// example:
	//
	// 2
	LayoutId *int64 `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMPULayoutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *CreateMPULayoutResponse) SetStatusCode(v int32) *CreateMPULayoutResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMPULayoutResponse) SetBody(v *CreateMPULayoutResponseBody) *CreateMPULayoutResponse {
	s.Body = v
	return s
}

type DeleteAppStreamingOutTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// wv7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	StreamingOutTemplate *DeleteAppStreamingOutTemplateRequestStreamingOutTemplate `json:"StreamingOutTemplate,omitempty" xml:"StreamingOutTemplate,omitempty" type:"Struct"`
}

func (s DeleteAppStreamingOutTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppStreamingOutTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppStreamingOutTemplateRequest) SetAppId(v string) *DeleteAppStreamingOutTemplateRequest {
	s.AppId = &v
	return s
}

func (s *DeleteAppStreamingOutTemplateRequest) SetStreamingOutTemplate(v *DeleteAppStreamingOutTemplateRequestStreamingOutTemplate) *DeleteAppStreamingOutTemplateRequest {
	s.StreamingOutTemplate = v
	return s
}

type DeleteAppStreamingOutTemplateRequestStreamingOutTemplate struct {
	// This parameter is required.
	//
	// example:
	//
	// xd4c****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteAppStreamingOutTemplateRequestStreamingOutTemplate) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppStreamingOutTemplateRequestStreamingOutTemplate) GoString() string {
	return s.String()
}

func (s *DeleteAppStreamingOutTemplateRequestStreamingOutTemplate) SetTemplateId(v string) *DeleteAppStreamingOutTemplateRequestStreamingOutTemplate {
	s.TemplateId = &v
	return s
}

type DeleteAppStreamingOutTemplateShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// wv7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	StreamingOutTemplateShrink *string `json:"StreamingOutTemplate,omitempty" xml:"StreamingOutTemplate,omitempty"`
}

func (s DeleteAppStreamingOutTemplateShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppStreamingOutTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppStreamingOutTemplateShrinkRequest) SetAppId(v string) *DeleteAppStreamingOutTemplateShrinkRequest {
	s.AppId = &v
	return s
}

func (s *DeleteAppStreamingOutTemplateShrinkRequest) SetStreamingOutTemplateShrink(v string) *DeleteAppStreamingOutTemplateShrinkRequest {
	s.StreamingOutTemplateShrink = &v
	return s
}

type DeleteAppStreamingOutTemplateResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 30D41049-D02D-1C21-86AE-B3E5FD825C17
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAppStreamingOutTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppStreamingOutTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppStreamingOutTemplateResponseBody) SetRequestId(v string) *DeleteAppStreamingOutTemplateResponseBody {
	s.RequestId = &v
	return s
}

type DeleteAppStreamingOutTemplateResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppStreamingOutTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppStreamingOutTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppStreamingOutTemplateResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppStreamingOutTemplateResponse) SetHeaders(v map[string]*string) *DeleteAppStreamingOutTemplateResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppStreamingOutTemplateResponse) SetStatusCode(v int32) *DeleteAppStreamingOutTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppStreamingOutTemplateResponse) SetBody(v *DeleteAppStreamingOutTemplateResponseBody) *DeleteAppStreamingOutTemplateResponse {
	s.Body = v
	return s
}

type DeleteAutoLiveStreamRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
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
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAutoLiveStreamRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteAutoLiveStreamRuleResponse) SetStatusCode(v int32) *DeleteAutoLiveStreamRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAutoLiveStreamRuleResponse) SetBody(v *DeleteAutoLiveStreamRuleResponseBody) *DeleteAutoLiveStreamRuleResponse {
	s.Body = v
	return s
}

type DeleteChannelRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DeleteChannelRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteChannelRequest) GoString() string {
	return s.String()
}

func (s *DeleteChannelRequest) SetAppId(v string) *DeleteChannelRequest {
	s.AppId = &v
	return s
}

func (s *DeleteChannelRequest) SetChannelId(v string) *DeleteChannelRequest {
	s.ChannelId = &v
	return s
}

func (s *DeleteChannelRequest) SetOwnerId(v int64) *DeleteChannelRequest {
	s.OwnerId = &v
	return s
}

type DeleteChannelResponseBody struct {
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CF8
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteChannelResponse) SetStatusCode(v int32) *DeleteChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteChannelResponse) SetBody(v *DeleteChannelResponseBody) *DeleteChannelResponse {
	s.Body = v
	return s
}

type DeleteEventSubscribeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 9qb1****
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ad53276431c****
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
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEventSubscribeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteEventSubscribeResponse) SetStatusCode(v int32) *DeleteEventSubscribeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEventSubscribeResponse) SetBody(v *DeleteEventSubscribeResponseBody) *DeleteEventSubscribeResponse {
	s.Body = v
	return s
}

type DeleteMPULayoutRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	LayoutId *int64 `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	OwnerId  *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMPULayoutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteMPULayoutResponse) SetStatusCode(v int32) *DeleteMPULayoutResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMPULayoutResponse) SetBody(v *DeleteMPULayoutResponseBody) *DeleteMPULayoutResponse {
	s.Body = v
	return s
}

type DeleteRecordTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 1
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 76dasgb****
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
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRecordTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteRecordTemplateResponse) SetStatusCode(v int32) *DeleteRecordTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRecordTemplateResponse) SetBody(v *DeleteRecordTemplateResponseBody) *DeleteRecordTemplateResponse {
	s.Body = v
	return s
}

type DescribeAppKeyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 0cho****
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeAppKeyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppKeyRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppKeyRequest) SetAppId(v string) *DescribeAppKeyRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppKeyRequest) SetOwnerId(v int64) *DescribeAppKeyRequest {
	s.OwnerId = &v
	return s
}

type DescribeAppKeyResponseBody struct {
	// AppKey。
	//
	// example:
	//
	// ba133b2cee4ab9be424674892c33****
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// 154EF5DE-3D08-1F2C-A482-281F78D74B7C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAppKeyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppKeyResponseBody) SetAppKey(v string) *DescribeAppKeyResponseBody {
	s.AppKey = &v
	return s
}

func (s *DescribeAppKeyResponseBody) SetRequestId(v string) *DescribeAppKeyResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAppKeyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppKeyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppKeyResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppKeyResponse) SetHeaders(v map[string]*string) *DescribeAppKeyResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppKeyResponse) SetStatusCode(v int32) *DescribeAppKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppKeyResponse) SetBody(v *DescribeAppKeyResponseBody) *DescribeAppKeyResponse {
	s.Body = v
	return s
}

type DescribeAppStreamingOutTemplatesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ioeh****
	AppId     *string                                           `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Condition *DescribeAppStreamingOutTemplatesRequestCondition `json:"Condition,omitempty" xml:"Condition,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeAppStreamingOutTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppStreamingOutTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppStreamingOutTemplatesRequest) SetAppId(v string) *DescribeAppStreamingOutTemplatesRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppStreamingOutTemplatesRequest) SetCondition(v *DescribeAppStreamingOutTemplatesRequestCondition) *DescribeAppStreamingOutTemplatesRequest {
	s.Condition = v
	return s
}

func (s *DescribeAppStreamingOutTemplatesRequest) SetPageNum(v int32) *DescribeAppStreamingOutTemplatesRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeAppStreamingOutTemplatesRequest) SetPageSize(v int32) *DescribeAppStreamingOutTemplatesRequest {
	s.PageSize = &v
	return s
}

type DescribeAppStreamingOutTemplatesRequestCondition struct {
	// example:
	//
	// 测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Bj6D****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeAppStreamingOutTemplatesRequestCondition) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppStreamingOutTemplatesRequestCondition) GoString() string {
	return s.String()
}

func (s *DescribeAppStreamingOutTemplatesRequestCondition) SetName(v string) *DescribeAppStreamingOutTemplatesRequestCondition {
	s.Name = &v
	return s
}

func (s *DescribeAppStreamingOutTemplatesRequestCondition) SetTemplateId(v string) *DescribeAppStreamingOutTemplatesRequestCondition {
	s.TemplateId = &v
	return s
}

type DescribeAppStreamingOutTemplatesShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ioeh****
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	ConditionShrink *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeAppStreamingOutTemplatesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppStreamingOutTemplatesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppStreamingOutTemplatesShrinkRequest) SetAppId(v string) *DescribeAppStreamingOutTemplatesShrinkRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppStreamingOutTemplatesShrinkRequest) SetConditionShrink(v string) *DescribeAppStreamingOutTemplatesShrinkRequest {
	s.ConditionShrink = &v
	return s
}

func (s *DescribeAppStreamingOutTemplatesShrinkRequest) SetPageNum(v int32) *DescribeAppStreamingOutTemplatesShrinkRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeAppStreamingOutTemplatesShrinkRequest) SetPageSize(v int32) *DescribeAppStreamingOutTemplatesShrinkRequest {
	s.PageSize = &v
	return s
}

type DescribeAppStreamingOutTemplatesResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// B0A2FCBC-43A4-428F-BC1D-3F4F85837F76
	RequestId *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Templates []*DescribeAppStreamingOutTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// example:
	//
	// 1
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeAppStreamingOutTemplatesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppStreamingOutTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppStreamingOutTemplatesResponseBody) SetRequestId(v string) *DescribeAppStreamingOutTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppStreamingOutTemplatesResponseBody) SetTemplates(v []*DescribeAppStreamingOutTemplatesResponseBodyTemplates) *DescribeAppStreamingOutTemplatesResponseBody {
	s.Templates = v
	return s
}

func (s *DescribeAppStreamingOutTemplatesResponseBody) SetTotalNum(v int64) *DescribeAppStreamingOutTemplatesResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeAppStreamingOutTemplatesResponseBody) SetTotalPage(v int64) *DescribeAppStreamingOutTemplatesResponseBody {
	s.TotalPage = &v
	return s
}

type DescribeAppStreamingOutTemplatesResponseBodyTemplates struct {
	// example:
	//
	// 2020-09-04T06:22:15Z
	CreateTime *string   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	LayoutIds  []*string `json:"LayoutIds,omitempty" xml:"LayoutIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	MediaEncode *int32 `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	// example:
	//
	// 模版名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Bj6D****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeAppStreamingOutTemplatesResponseBodyTemplates) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppStreamingOutTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *DescribeAppStreamingOutTemplatesResponseBodyTemplates) SetCreateTime(v string) *DescribeAppStreamingOutTemplatesResponseBodyTemplates {
	s.CreateTime = &v
	return s
}

func (s *DescribeAppStreamingOutTemplatesResponseBodyTemplates) SetLayoutIds(v []*string) *DescribeAppStreamingOutTemplatesResponseBodyTemplates {
	s.LayoutIds = v
	return s
}

func (s *DescribeAppStreamingOutTemplatesResponseBodyTemplates) SetMediaEncode(v int32) *DescribeAppStreamingOutTemplatesResponseBodyTemplates {
	s.MediaEncode = &v
	return s
}

func (s *DescribeAppStreamingOutTemplatesResponseBodyTemplates) SetName(v string) *DescribeAppStreamingOutTemplatesResponseBodyTemplates {
	s.Name = &v
	return s
}

func (s *DescribeAppStreamingOutTemplatesResponseBodyTemplates) SetTemplateId(v string) *DescribeAppStreamingOutTemplatesResponseBodyTemplates {
	s.TemplateId = &v
	return s
}

type DescribeAppStreamingOutTemplatesResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppStreamingOutTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppStreamingOutTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppStreamingOutTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppStreamingOutTemplatesResponse) SetHeaders(v map[string]*string) *DescribeAppStreamingOutTemplatesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppStreamingOutTemplatesResponse) SetStatusCode(v int32) *DescribeAppStreamingOutTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppStreamingOutTemplatesResponse) SetBody(v *DescribeAppStreamingOutTemplatesResponseBody) *DescribeAppStreamingOutTemplatesResponse {
	s.Body = v
	return s
}

type DescribeAppsRequest struct {
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 3.0
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// example:
	//
	// asc
	Order   *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppsRequest) SetAppId(v string) *DescribeAppsRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppsRequest) SetAppVersion(v string) *DescribeAppsRequest {
	s.AppVersion = &v
	return s
}

func (s *DescribeAppsRequest) SetOrder(v string) *DescribeAppsRequest {
	s.Order = &v
	return s
}

func (s *DescribeAppsRequest) SetOwnerId(v int64) *DescribeAppsRequest {
	s.OwnerId = &v
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

func (s *DescribeAppsRequest) SetStatus(v string) *DescribeAppsRequest {
	s.Status = &v
	return s
}

type DescribeAppsResponseBody struct {
	AppList *DescribeAppsResponseBodyAppList `json:"AppList,omitempty" xml:"AppList,omitempty" type:"Struct"`
	// example:
	//
	// 6159ba01-6687-4fb2-a831-f0cd8d188648
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// example:
	//
	// 1
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBody) SetAppList(v *DescribeAppsResponseBodyAppList) *DescribeAppsResponseBody {
	s.AppList = v
	return s
}

func (s *DescribeAppsResponseBody) SetRequestId(v string) *DescribeAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppsResponseBody) SetTotalNum(v int32) *DescribeAppsResponseBody {
	s.TotalNum = &v
	return s
}

func (s *DescribeAppsResponseBody) SetTotalPage(v int32) *DescribeAppsResponseBody {
	s.TotalPage = &v
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
	// example:
	//
	// rgf1****"
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// Default AppName
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// universal
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// paybyduration
	BillType *string `json:"BillType,omitempty" xml:"BillType,omitempty"`
	// example:
	//
	// 2020-01-09T02:02:29Z
	CreateTime   *string                                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ServiceAreas *DescribeAppsResponseBodyAppListAppServiceAreas `json:"ServiceAreas,omitempty" xml:"ServiceAreas,omitempty" type:"Struct"`
	// example:
	//
	// 1
	Status  *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeAppsResponseBodyAppListApp) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponseBodyAppListApp) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyAppListApp) SetAppId(v string) *DescribeAppsResponseBodyAppListApp {
	s.AppId = &v
	return s
}

func (s *DescribeAppsResponseBodyAppListApp) SetAppName(v string) *DescribeAppsResponseBodyAppListApp {
	s.AppName = &v
	return s
}

func (s *DescribeAppsResponseBodyAppListApp) SetAppType(v string) *DescribeAppsResponseBodyAppListApp {
	s.AppType = &v
	return s
}

func (s *DescribeAppsResponseBodyAppListApp) SetBillType(v string) *DescribeAppsResponseBodyAppListApp {
	s.BillType = &v
	return s
}

func (s *DescribeAppsResponseBodyAppListApp) SetCreateTime(v string) *DescribeAppsResponseBodyAppListApp {
	s.CreateTime = &v
	return s
}

func (s *DescribeAppsResponseBodyAppListApp) SetServiceAreas(v *DescribeAppsResponseBodyAppListAppServiceAreas) *DescribeAppsResponseBodyAppListApp {
	s.ServiceAreas = v
	return s
}

func (s *DescribeAppsResponseBodyAppListApp) SetStatus(v int32) *DescribeAppsResponseBodyAppListApp {
	s.Status = &v
	return s
}

func (s *DescribeAppsResponseBodyAppListApp) SetVersion(v string) *DescribeAppsResponseBodyAppListApp {
	s.Version = &v
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeAppsResponse) SetStatusCode(v int32) *DescribeAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppsResponse) SetBody(v *DescribeAppsResponseBody) *DescribeAppsResponse {
	s.Body = v
	return s
}

type DescribeAutoLiveStreamRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eo85****
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
	// example:
	//
	// 069BCB66-CD80-11E8-A82B-A70F78BBDC00
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
	// example:
	//
	// http://example.com/callBack
	CallBack          *string   `json:"CallBack,omitempty" xml:"CallBack,omitempty"`
	ChannelIdPrefixes []*string `json:"ChannelIdPrefixes,omitempty" xml:"ChannelIdPrefixes,omitempty" type:"Repeated"`
	ChannelIds        []*string `json:"ChannelIds,omitempty" xml:"ChannelIds,omitempty" type:"Repeated"`
	// example:
	//
	// 2021-08-19T02:53:07Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 20
	MediaEncode *int32 `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	// example:
	//
	// rtmp://${domain}/${app}/${stream}
	PlayDomain *string `json:"PlayDomain,omitempty" xml:"PlayDomain,omitempty"`
	// example:
	//
	// 12
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// testRule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// disable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAutoLiveStreamRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeAutoLiveStreamRuleResponse) SetStatusCode(v int32) *DescribeAutoLiveStreamRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAutoLiveStreamRuleResponse) SetBody(v *DescribeAutoLiveStreamRuleResponseBody) *DescribeAutoLiveStreamRuleResponse {
	s.Body = v
	return s
}

type DescribeCallRequest struct {
	// APP ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// testappid
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1614936817
	CreatedTs *int64 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	// example:
	//
	// 1614936817
	DestroyedTs *int64 `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	// example:
	//
	// USER_DURATION_STAT
	ExtDataType *string `json:"ExtDataType,omitempty" xml:"ExtDataType,omitempty"`
	// example:
	//
	// false
	QueryExpInfo *bool `json:"QueryExpInfo,omitempty" xml:"QueryExpInfo,omitempty"`
}

func (s DescribeCallRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCallRequest) GoString() string {
	return s.String()
}

func (s *DescribeCallRequest) SetAppId(v string) *DescribeCallRequest {
	s.AppId = &v
	return s
}

func (s *DescribeCallRequest) SetChannelId(v string) *DescribeCallRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeCallRequest) SetCreatedTs(v int64) *DescribeCallRequest {
	s.CreatedTs = &v
	return s
}

func (s *DescribeCallRequest) SetDestroyedTs(v int64) *DescribeCallRequest {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeCallRequest) SetExtDataType(v string) *DescribeCallRequest {
	s.ExtDataType = &v
	return s
}

func (s *DescribeCallRequest) SetQueryExpInfo(v bool) *DescribeCallRequest {
	s.QueryExpInfo = &v
	return s
}

type DescribeCallResponseBody struct {
	CallInfo *DescribeCallResponseBodyCallInfo `json:"CallInfo,omitempty" xml:"CallInfo,omitempty" type:"Struct"`
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId      *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserDetailList []*DescribeCallResponseBodyUserDetailList `json:"UserDetailList,omitempty" xml:"UserDetailList,omitempty" type:"Repeated"`
}

func (s DescribeCallResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCallResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCallResponseBody) SetCallInfo(v *DescribeCallResponseBodyCallInfo) *DescribeCallResponseBody {
	s.CallInfo = v
	return s
}

func (s *DescribeCallResponseBody) SetRequestId(v string) *DescribeCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCallResponseBody) SetUserDetailList(v []*DescribeCallResponseBodyUserDetailList) *DescribeCallResponseBody {
	s.UserDetailList = v
	return s
}

type DescribeCallResponseBodyCallInfo struct {
	// App ID。
	//
	// example:
	//
	// xxxxxxxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// IN
	CallStatus *string `json:"CallStatus,omitempty" xml:"CallStatus,omitempty"`
	// example:
	//
	// 123456
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// 1615860711
	CreatedTs *int64 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	// example:
	//
	// 1615860811
	DestroyedTs *int64 `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	// example:
	//
	// 100
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s DescribeCallResponseBodyCallInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeCallResponseBodyCallInfo) GoString() string {
	return s.String()
}

func (s *DescribeCallResponseBodyCallInfo) SetAppId(v string) *DescribeCallResponseBodyCallInfo {
	s.AppId = &v
	return s
}

func (s *DescribeCallResponseBodyCallInfo) SetCallStatus(v string) *DescribeCallResponseBodyCallInfo {
	s.CallStatus = &v
	return s
}

func (s *DescribeCallResponseBodyCallInfo) SetChannelId(v string) *DescribeCallResponseBodyCallInfo {
	s.ChannelId = &v
	return s
}

func (s *DescribeCallResponseBodyCallInfo) SetCreatedTs(v int64) *DescribeCallResponseBodyCallInfo {
	s.CreatedTs = &v
	return s
}

func (s *DescribeCallResponseBodyCallInfo) SetDestroyedTs(v int64) *DescribeCallResponseBodyCallInfo {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeCallResponseBodyCallInfo) SetDuration(v int64) *DescribeCallResponseBodyCallInfo {
	s.Duration = &v
	return s
}

type DescribeCallResponseBodyUserDetailList struct {
	// example:
	//
	// GOOD
	CallExp *string `json:"CallExp,omitempty" xml:"CallExp,omitempty"`
	// example:
	//
	// 1614936817
	CreatedTs *int64 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	// example:
	//
	// 1614936817
	DestroyedTs       *int64                                                   `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	DurMetricStatData *DescribeCallResponseBodyUserDetailListDurMetricStatData `json:"DurMetricStatData,omitempty" xml:"DurMetricStatData,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 浙江省-杭州市
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// example:
	//
	// 4G
	Network     *string   `json:"Network,omitempty" xml:"Network,omitempty"`
	NetworkList []*string `json:"NetworkList,omitempty" xml:"NetworkList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	OnlineDuration *int64                                                 `json:"OnlineDuration,omitempty" xml:"OnlineDuration,omitempty"`
	OnlinePeriods  []*DescribeCallResponseBodyUserDetailListOnlinePeriods `json:"OnlinePeriods,omitempty" xml:"OnlinePeriods,omitempty" type:"Repeated"`
	// example:
	//
	// iOS
	Os     *string   `json:"Os,omitempty" xml:"Os,omitempty"`
	OsList []*string `json:"OsList,omitempty" xml:"OsList,omitempty" type:"Repeated"`
	Roles  []*string `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// example:
	//
	// 1.0.0
	SdkVersion     *string   `json:"SdkVersion,omitempty" xml:"SdkVersion,omitempty"`
	SdkVersionList []*string `json:"SdkVersionList,omitempty" xml:"SdkVersionList,omitempty" type:"Repeated"`
	// example:
	//
	// testuserid
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeCallResponseBodyUserDetailList) String() string {
	return tea.Prettify(s)
}

func (s DescribeCallResponseBodyUserDetailList) GoString() string {
	return s.String()
}

func (s *DescribeCallResponseBodyUserDetailList) SetCallExp(v string) *DescribeCallResponseBodyUserDetailList {
	s.CallExp = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailList) SetCreatedTs(v int64) *DescribeCallResponseBodyUserDetailList {
	s.CreatedTs = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailList) SetDestroyedTs(v int64) *DescribeCallResponseBodyUserDetailList {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailList) SetDurMetricStatData(v *DescribeCallResponseBodyUserDetailListDurMetricStatData) *DescribeCallResponseBodyUserDetailList {
	s.DurMetricStatData = v
	return s
}

func (s *DescribeCallResponseBodyUserDetailList) SetDuration(v int64) *DescribeCallResponseBodyUserDetailList {
	s.Duration = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailList) SetLocation(v string) *DescribeCallResponseBodyUserDetailList {
	s.Location = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailList) SetNetwork(v string) *DescribeCallResponseBodyUserDetailList {
	s.Network = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailList) SetNetworkList(v []*string) *DescribeCallResponseBodyUserDetailList {
	s.NetworkList = v
	return s
}

func (s *DescribeCallResponseBodyUserDetailList) SetOnlineDuration(v int64) *DescribeCallResponseBodyUserDetailList {
	s.OnlineDuration = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailList) SetOnlinePeriods(v []*DescribeCallResponseBodyUserDetailListOnlinePeriods) *DescribeCallResponseBodyUserDetailList {
	s.OnlinePeriods = v
	return s
}

func (s *DescribeCallResponseBodyUserDetailList) SetOs(v string) *DescribeCallResponseBodyUserDetailList {
	s.Os = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailList) SetOsList(v []*string) *DescribeCallResponseBodyUserDetailList {
	s.OsList = v
	return s
}

func (s *DescribeCallResponseBodyUserDetailList) SetRoles(v []*string) *DescribeCallResponseBodyUserDetailList {
	s.Roles = v
	return s
}

func (s *DescribeCallResponseBodyUserDetailList) SetSdkVersion(v string) *DescribeCallResponseBodyUserDetailList {
	s.SdkVersion = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailList) SetSdkVersionList(v []*string) *DescribeCallResponseBodyUserDetailList {
	s.SdkVersionList = v
	return s
}

func (s *DescribeCallResponseBodyUserDetailList) SetUserId(v string) *DescribeCallResponseBodyUserDetailList {
	s.UserId = &v
	return s
}

type DescribeCallResponseBodyUserDetailListDurMetricStatData struct {
	// example:
	//
	// 0
	PubAudio *int64 `json:"PubAudio,omitempty" xml:"PubAudio,omitempty"`
	// example:
	//
	// 0
	PubVideo1080 *int64 `json:"PubVideo1080,omitempty" xml:"PubVideo1080,omitempty"`
	// example:
	//
	// 0
	PubVideo360 *int64 `json:"PubVideo360,omitempty" xml:"PubVideo360,omitempty"`
	// example:
	//
	// 0
	PubVideo720 *int64 `json:"PubVideo720,omitempty" xml:"PubVideo720,omitempty"`
	// example:
	//
	// 0
	PubVideoScreenShare *int64 `json:"PubVideoScreenShare,omitempty" xml:"PubVideoScreenShare,omitempty"`
	// example:
	//
	// 0
	SubAudio *int64 `json:"SubAudio,omitempty" xml:"SubAudio,omitempty"`
	// example:
	//
	// 0
	SubVideo1080 *int64 `json:"SubVideo1080,omitempty" xml:"SubVideo1080,omitempty"`
	// example:
	//
	// 0
	SubVideo360 *int64 `json:"SubVideo360,omitempty" xml:"SubVideo360,omitempty"`
	// example:
	//
	// 0
	SubVideo720 *int64 `json:"SubVideo720,omitempty" xml:"SubVideo720,omitempty"`
	// example:
	//
	// 0
	SubVideoScreenShare *int64 `json:"SubVideoScreenShare,omitempty" xml:"SubVideoScreenShare,omitempty"`
}

func (s DescribeCallResponseBodyUserDetailListDurMetricStatData) String() string {
	return tea.Prettify(s)
}

func (s DescribeCallResponseBodyUserDetailListDurMetricStatData) GoString() string {
	return s.String()
}

func (s *DescribeCallResponseBodyUserDetailListDurMetricStatData) SetPubAudio(v int64) *DescribeCallResponseBodyUserDetailListDurMetricStatData {
	s.PubAudio = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailListDurMetricStatData) SetPubVideo1080(v int64) *DescribeCallResponseBodyUserDetailListDurMetricStatData {
	s.PubVideo1080 = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailListDurMetricStatData) SetPubVideo360(v int64) *DescribeCallResponseBodyUserDetailListDurMetricStatData {
	s.PubVideo360 = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailListDurMetricStatData) SetPubVideo720(v int64) *DescribeCallResponseBodyUserDetailListDurMetricStatData {
	s.PubVideo720 = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailListDurMetricStatData) SetPubVideoScreenShare(v int64) *DescribeCallResponseBodyUserDetailListDurMetricStatData {
	s.PubVideoScreenShare = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailListDurMetricStatData) SetSubAudio(v int64) *DescribeCallResponseBodyUserDetailListDurMetricStatData {
	s.SubAudio = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailListDurMetricStatData) SetSubVideo1080(v int64) *DescribeCallResponseBodyUserDetailListDurMetricStatData {
	s.SubVideo1080 = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailListDurMetricStatData) SetSubVideo360(v int64) *DescribeCallResponseBodyUserDetailListDurMetricStatData {
	s.SubVideo360 = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailListDurMetricStatData) SetSubVideo720(v int64) *DescribeCallResponseBodyUserDetailListDurMetricStatData {
	s.SubVideo720 = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailListDurMetricStatData) SetSubVideoScreenShare(v int64) *DescribeCallResponseBodyUserDetailListDurMetricStatData {
	s.SubVideoScreenShare = &v
	return s
}

type DescribeCallResponseBodyUserDetailListOnlinePeriods struct {
	// example:
	//
	// 1614936817
	JoinTs *int64 `json:"JoinTs,omitempty" xml:"JoinTs,omitempty"`
	// example:
	//
	// 1614936817
	LeaveTs *int64 `json:"LeaveTs,omitempty" xml:"LeaveTs,omitempty"`
}

func (s DescribeCallResponseBodyUserDetailListOnlinePeriods) String() string {
	return tea.Prettify(s)
}

func (s DescribeCallResponseBodyUserDetailListOnlinePeriods) GoString() string {
	return s.String()
}

func (s *DescribeCallResponseBodyUserDetailListOnlinePeriods) SetJoinTs(v int64) *DescribeCallResponseBodyUserDetailListOnlinePeriods {
	s.JoinTs = &v
	return s
}

func (s *DescribeCallResponseBodyUserDetailListOnlinePeriods) SetLeaveTs(v int64) *DescribeCallResponseBodyUserDetailListOnlinePeriods {
	s.LeaveTs = &v
	return s
}

type DescribeCallResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCallResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCallResponse) GoString() string {
	return s.String()
}

func (s *DescribeCallResponse) SetHeaders(v map[string]*string) *DescribeCallResponse {
	s.Headers = v
	return s
}

func (s *DescribeCallResponse) SetStatusCode(v int32) *DescribeCallResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCallResponse) SetBody(v *DescribeCallResponseBody) *DescribeCallResponse {
	s.Body = v
	return s
}

type DescribeCallListRequest struct {
	// APP ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// testappid
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// OUT
	CallStatus *string `json:"CallStatus,omitempty" xml:"CallStatus,omitempty"`
	// example:
	//
	// 311
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1615892596
	EndTs *int64 `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	// example:
	//
	// BAD_EXP_USER_COUNT_DESC
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// ALL
	QueryMode *string `json:"QueryMode,omitempty" xml:"QueryMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1615806196
	StartTs *int64 `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
	// example:
	//
	// c906531af5f9****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeCallListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCallListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCallListRequest) SetAppId(v string) *DescribeCallListRequest {
	s.AppId = &v
	return s
}

func (s *DescribeCallListRequest) SetCallStatus(v string) *DescribeCallListRequest {
	s.CallStatus = &v
	return s
}

func (s *DescribeCallListRequest) SetChannelId(v string) *DescribeCallListRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeCallListRequest) SetEndTs(v int64) *DescribeCallListRequest {
	s.EndTs = &v
	return s
}

func (s *DescribeCallListRequest) SetOrderBy(v string) *DescribeCallListRequest {
	s.OrderBy = &v
	return s
}

func (s *DescribeCallListRequest) SetPageNo(v int32) *DescribeCallListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeCallListRequest) SetPageSize(v int32) *DescribeCallListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCallListRequest) SetQueryMode(v string) *DescribeCallListRequest {
	s.QueryMode = &v
	return s
}

func (s *DescribeCallListRequest) SetStartTs(v int64) *DescribeCallListRequest {
	s.StartTs = &v
	return s
}

func (s *DescribeCallListRequest) SetUserId(v string) *DescribeCallListRequest {
	s.UserId = &v
	return s
}

type DescribeCallListResponseBody struct {
	CallList []*DescribeCallListResponseBodyCallList `json:"CallList,omitempty" xml:"CallList,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 20
	TotalCnt *int32 `json:"TotalCnt,omitempty" xml:"TotalCnt,omitempty"`
}

func (s DescribeCallListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCallListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCallListResponseBody) SetCallList(v []*DescribeCallListResponseBodyCallList) *DescribeCallListResponseBody {
	s.CallList = v
	return s
}

func (s *DescribeCallListResponseBody) SetPageNo(v int32) *DescribeCallListResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribeCallListResponseBody) SetPageSize(v int32) *DescribeCallListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCallListResponseBody) SetRequestId(v string) *DescribeCallListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCallListResponseBody) SetTotalCnt(v int32) *DescribeCallListResponseBody {
	s.TotalCnt = &v
	return s
}

type DescribeCallListResponseBodyCallList struct {
	// App ID。
	//
	// example:
	//
	// 9qb1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 0
	BadExpUserCnt *int32 `json:"BadExpUserCnt,omitempty" xml:"BadExpUserCnt,omitempty"`
	// example:
	//
	// OUT
	CallStatus *string `json:"CallStatus,omitempty" xml:"CallStatus,omitempty"`
	// example:
	//
	// 904
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// 1614936817
	CreatedTs *int64 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	// example:
	//
	// 1614936817
	DestroyedTs *int64 `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	// example:
	//
	// 10
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 5
	UserCnt *int32 `json:"UserCnt,omitempty" xml:"UserCnt,omitempty"`
}

func (s DescribeCallListResponseBodyCallList) String() string {
	return tea.Prettify(s)
}

func (s DescribeCallListResponseBodyCallList) GoString() string {
	return s.String()
}

func (s *DescribeCallListResponseBodyCallList) SetAppId(v string) *DescribeCallListResponseBodyCallList {
	s.AppId = &v
	return s
}

func (s *DescribeCallListResponseBodyCallList) SetBadExpUserCnt(v int32) *DescribeCallListResponseBodyCallList {
	s.BadExpUserCnt = &v
	return s
}

func (s *DescribeCallListResponseBodyCallList) SetCallStatus(v string) *DescribeCallListResponseBodyCallList {
	s.CallStatus = &v
	return s
}

func (s *DescribeCallListResponseBodyCallList) SetChannelId(v string) *DescribeCallListResponseBodyCallList {
	s.ChannelId = &v
	return s
}

func (s *DescribeCallListResponseBodyCallList) SetCreatedTs(v int64) *DescribeCallListResponseBodyCallList {
	s.CreatedTs = &v
	return s
}

func (s *DescribeCallListResponseBodyCallList) SetDestroyedTs(v int64) *DescribeCallListResponseBodyCallList {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeCallListResponseBodyCallList) SetDuration(v int64) *DescribeCallListResponseBodyCallList {
	s.Duration = &v
	return s
}

func (s *DescribeCallListResponseBodyCallList) SetUserCnt(v int32) *DescribeCallListResponseBodyCallList {
	s.UserCnt = &v
	return s
}

type DescribeCallListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCallListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCallListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCallListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCallListResponse) SetHeaders(v map[string]*string) *DescribeCallListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCallListResponse) SetStatusCode(v int32) *DescribeCallListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCallListResponse) SetBody(v *DescribeCallListResponseBody) *DescribeCallListResponse {
	s.Body = v
	return s
}

type DescribeChannelRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
}

func (s DescribeChannelRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelRequest) GoString() string {
	return s.String()
}

func (s *DescribeChannelRequest) SetAppId(v string) *DescribeChannelRequest {
	s.AppId = &v
	return s
}

func (s *DescribeChannelRequest) SetChannelId(v string) *DescribeChannelRequest {
	s.ChannelId = &v
	return s
}

type DescribeChannelResponseBody struct {
	// channel
	Channel      *DescribeChannelResponseBodyChannel `json:"Channel,omitempty" xml:"Channel,omitempty" type:"Struct"`
	ChannelExist *bool                               `json:"ChannelExist,omitempty" xml:"ChannelExist,omitempty"`
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeChannelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChannelResponseBody) SetChannel(v *DescribeChannelResponseBodyChannel) *DescribeChannelResponseBody {
	s.Channel = v
	return s
}

func (s *DescribeChannelResponseBody) SetChannelExist(v bool) *DescribeChannelResponseBody {
	s.ChannelExist = &v
	return s
}

func (s *DescribeChannelResponseBody) SetRequestId(v string) *DescribeChannelResponseBody {
	s.RequestId = &v
	return s
}

type DescribeChannelResponseBodyChannel struct {
	// example:
	//
	// testid
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// 1557909133
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeChannelResponseBodyChannel) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelResponseBodyChannel) GoString() string {
	return s.String()
}

func (s *DescribeChannelResponseBodyChannel) SetChannelId(v string) *DescribeChannelResponseBodyChannel {
	s.ChannelId = &v
	return s
}

func (s *DescribeChannelResponseBodyChannel) SetStartTime(v int64) *DescribeChannelResponseBodyChannel {
	s.StartTime = &v
	return s
}

type DescribeChannelResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeChannelResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelResponse) GoString() string {
	return s.String()
}

func (s *DescribeChannelResponse) SetHeaders(v map[string]*string) *DescribeChannelResponse {
	s.Headers = v
	return s
}

func (s *DescribeChannelResponse) SetStatusCode(v int32) *DescribeChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeChannelResponse) SetBody(v *DescribeChannelResponseBody) *DescribeChannelResponse {
	s.Body = v
	return s
}

type DescribeChannelAllUsersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
}

func (s DescribeChannelAllUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelAllUsersRequest) GoString() string {
	return s.String()
}

func (s *DescribeChannelAllUsersRequest) SetAppId(v string) *DescribeChannelAllUsersRequest {
	s.AppId = &v
	return s
}

func (s *DescribeChannelAllUsersRequest) SetChannelId(v string) *DescribeChannelAllUsersRequest {
	s.ChannelId = &v
	return s
}

type DescribeChannelAllUsersResponseBody struct {
	ChannelExist *bool `json:"ChannelExist,omitempty" xml:"ChannelExist,omitempty"`
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CF8
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Users     []*DescribeChannelAllUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s DescribeChannelAllUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelAllUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChannelAllUsersResponseBody) SetChannelExist(v bool) *DescribeChannelAllUsersResponseBody {
	s.ChannelExist = &v
	return s
}

func (s *DescribeChannelAllUsersResponseBody) SetRequestId(v string) *DescribeChannelAllUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeChannelAllUsersResponseBody) SetUsers(v []*DescribeChannelAllUsersResponseBodyUsers) *DescribeChannelAllUsersResponseBody {
	s.Users = v
	return s
}

type DescribeChannelAllUsersResponseBodyUsers struct {
	// example:
	//
	// 1811****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeChannelAllUsersResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelAllUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *DescribeChannelAllUsersResponseBodyUsers) SetUserId(v string) *DescribeChannelAllUsersResponseBodyUsers {
	s.UserId = &v
	return s
}

type DescribeChannelAllUsersResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeChannelAllUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeChannelAllUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelAllUsersResponse) GoString() string {
	return s.String()
}

func (s *DescribeChannelAllUsersResponse) SetHeaders(v map[string]*string) *DescribeChannelAllUsersResponse {
	s.Headers = v
	return s
}

func (s *DescribeChannelAllUsersResponse) SetStatusCode(v int32) *DescribeChannelAllUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeChannelAllUsersResponse) SetBody(v *DescribeChannelAllUsersResponseBody) *DescribeChannelAllUsersResponse {
	s.Body = v
	return s
}

type DescribeChannelAreaDistributionStatDataRequest struct {
	// APP ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// testappid
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1614936817
	CreatedTs *int64 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	// example:
	//
	// 1614936817
	DestroyedTs *int64  `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	ParentArea  *string `json:"ParentArea,omitempty" xml:"ParentArea,omitempty"`
}

func (s DescribeChannelAreaDistributionStatDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelAreaDistributionStatDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeChannelAreaDistributionStatDataRequest) SetAppId(v string) *DescribeChannelAreaDistributionStatDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeChannelAreaDistributionStatDataRequest) SetChannelId(v string) *DescribeChannelAreaDistributionStatDataRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeChannelAreaDistributionStatDataRequest) SetCreatedTs(v int64) *DescribeChannelAreaDistributionStatDataRequest {
	s.CreatedTs = &v
	return s
}

func (s *DescribeChannelAreaDistributionStatDataRequest) SetDestroyedTs(v int64) *DescribeChannelAreaDistributionStatDataRequest {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeChannelAreaDistributionStatDataRequest) SetParentArea(v string) *DescribeChannelAreaDistributionStatDataRequest {
	s.ParentArea = &v
	return s
}

type DescribeChannelAreaDistributionStatDataResponseBody struct {
	AreaStatList []*DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList `json:"AreaStatList,omitempty" xml:"AreaStatList,omitempty" type:"Repeated"`
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeChannelAreaDistributionStatDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelAreaDistributionStatDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChannelAreaDistributionStatDataResponseBody) SetAreaStatList(v []*DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList) *DescribeChannelAreaDistributionStatDataResponseBody {
	s.AreaStatList = v
	return s
}

func (s *DescribeChannelAreaDistributionStatDataResponseBody) SetRequestId(v string) *DescribeChannelAreaDistributionStatDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList struct {
	// example:
	//
	// 浙江省
	AreaName *string `json:"AreaName,omitempty" xml:"AreaName,omitempty"`
	// example:
	//
	// 1
	CallUserCount *int32 `json:"CallUserCount,omitempty" xml:"CallUserCount,omitempty"`
	// example:
	//
	// 0.9999
	HighQualityTransmissionRate *string `json:"HighQualityTransmissionRate,omitempty" xml:"HighQualityTransmissionRate,omitempty"`
	// example:
	//
	// 1
	PubUserCount *int32 `json:"PubUserCount,omitempty" xml:"PubUserCount,omitempty"`
	// example:
	//
	// 1
	SubUserCount *int32 `json:"SubUserCount,omitempty" xml:"SubUserCount,omitempty"`
}

func (s DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList) GoString() string {
	return s.String()
}

func (s *DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList) SetAreaName(v string) *DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList {
	s.AreaName = &v
	return s
}

func (s *DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList) SetCallUserCount(v int32) *DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList {
	s.CallUserCount = &v
	return s
}

func (s *DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList) SetHighQualityTransmissionRate(v string) *DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList {
	s.HighQualityTransmissionRate = &v
	return s
}

func (s *DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList) SetPubUserCount(v int32) *DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList {
	s.PubUserCount = &v
	return s
}

func (s *DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList) SetSubUserCount(v int32) *DescribeChannelAreaDistributionStatDataResponseBodyAreaStatList {
	s.SubUserCount = &v
	return s
}

type DescribeChannelAreaDistributionStatDataResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeChannelAreaDistributionStatDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeChannelAreaDistributionStatDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelAreaDistributionStatDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeChannelAreaDistributionStatDataResponse) SetHeaders(v map[string]*string) *DescribeChannelAreaDistributionStatDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeChannelAreaDistributionStatDataResponse) SetStatusCode(v int32) *DescribeChannelAreaDistributionStatDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeChannelAreaDistributionStatDataResponse) SetBody(v *DescribeChannelAreaDistributionStatDataResponseBody) *DescribeChannelAreaDistributionStatDataResponse {
	s.Body = v
	return s
}

type DescribeChannelDistributionStatDataRequest struct {
	// APP ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// testappid
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1614936817
	CreatedTs *int64 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	// example:
	//
	// 1614936817
	DestroyedTs *int64 `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OS
	StatDim *string `json:"StatDim,omitempty" xml:"StatDim,omitempty"`
}

func (s DescribeChannelDistributionStatDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelDistributionStatDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeChannelDistributionStatDataRequest) SetAppId(v string) *DescribeChannelDistributionStatDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeChannelDistributionStatDataRequest) SetChannelId(v string) *DescribeChannelDistributionStatDataRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeChannelDistributionStatDataRequest) SetCreatedTs(v int64) *DescribeChannelDistributionStatDataRequest {
	s.CreatedTs = &v
	return s
}

func (s *DescribeChannelDistributionStatDataRequest) SetDestroyedTs(v int64) *DescribeChannelDistributionStatDataRequest {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeChannelDistributionStatDataRequest) SetStatDim(v string) *DescribeChannelDistributionStatDataRequest {
	s.StatDim = &v
	return s
}

type DescribeChannelDistributionStatDataResponseBody struct {
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StatList  []*DescribeChannelDistributionStatDataResponseBodyStatList `json:"StatList,omitempty" xml:"StatList,omitempty" type:"Repeated"`
}

func (s DescribeChannelDistributionStatDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelDistributionStatDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChannelDistributionStatDataResponseBody) SetRequestId(v string) *DescribeChannelDistributionStatDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeChannelDistributionStatDataResponseBody) SetStatList(v []*DescribeChannelDistributionStatDataResponseBodyStatList) *DescribeChannelDistributionStatDataResponseBody {
	s.StatList = v
	return s
}

type DescribeChannelDistributionStatDataResponseBodyStatList struct {
	// example:
	//
	// 1
	CallUserCount *int32 `json:"CallUserCount,omitempty" xml:"CallUserCount,omitempty"`
	// example:
	//
	// 1.0000
	CallUserRatio *string `json:"CallUserRatio,omitempty" xml:"CallUserRatio,omitempty"`
	// example:
	//
	// OS
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeChannelDistributionStatDataResponseBodyStatList) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelDistributionStatDataResponseBodyStatList) GoString() string {
	return s.String()
}

func (s *DescribeChannelDistributionStatDataResponseBodyStatList) SetCallUserCount(v int32) *DescribeChannelDistributionStatDataResponseBodyStatList {
	s.CallUserCount = &v
	return s
}

func (s *DescribeChannelDistributionStatDataResponseBodyStatList) SetCallUserRatio(v string) *DescribeChannelDistributionStatDataResponseBodyStatList {
	s.CallUserRatio = &v
	return s
}

func (s *DescribeChannelDistributionStatDataResponseBodyStatList) SetName(v string) *DescribeChannelDistributionStatDataResponseBodyStatList {
	s.Name = &v
	return s
}

type DescribeChannelDistributionStatDataResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeChannelDistributionStatDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeChannelDistributionStatDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelDistributionStatDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeChannelDistributionStatDataResponse) SetHeaders(v map[string]*string) *DescribeChannelDistributionStatDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeChannelDistributionStatDataResponse) SetStatusCode(v int32) *DescribeChannelDistributionStatDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeChannelDistributionStatDataResponse) SetBody(v *DescribeChannelDistributionStatDataResponseBody) *DescribeChannelDistributionStatDataResponse {
	s.Body = v
	return s
}

type DescribeChannelOverallDataRequest struct {
	// APP ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// testappid
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1615893133
	CreatedTs *int64 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	// example:
	//
	// 1615893757
	DestroyedTs *int64 `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
}

func (s DescribeChannelOverallDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelOverallDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeChannelOverallDataRequest) SetAppId(v string) *DescribeChannelOverallDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeChannelOverallDataRequest) SetChannelId(v string) *DescribeChannelOverallDataRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeChannelOverallDataRequest) SetCreatedTs(v int64) *DescribeChannelOverallDataRequest {
	s.CreatedTs = &v
	return s
}

func (s *DescribeChannelOverallDataRequest) SetDestroyedTs(v int64) *DescribeChannelOverallDataRequest {
	s.DestroyedTs = &v
	return s
}

type DescribeChannelOverallDataResponseBody struct {
	CallInfo    *DescribeChannelOverallDataResponseBodyCallInfo      `json:"CallInfo,omitempty" xml:"CallInfo,omitempty" type:"Struct"`
	MetricDatas []*DescribeChannelOverallDataResponseBodyMetricDatas `json:"MetricDatas,omitempty" xml:"MetricDatas,omitempty" type:"Repeated"`
	OverallData *DescribeChannelOverallDataResponseBodyOverallData   `json:"OverallData,omitempty" xml:"OverallData,omitempty" type:"Struct"`
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeChannelOverallDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelOverallDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChannelOverallDataResponseBody) SetCallInfo(v *DescribeChannelOverallDataResponseBodyCallInfo) *DescribeChannelOverallDataResponseBody {
	s.CallInfo = v
	return s
}

func (s *DescribeChannelOverallDataResponseBody) SetMetricDatas(v []*DescribeChannelOverallDataResponseBodyMetricDatas) *DescribeChannelOverallDataResponseBody {
	s.MetricDatas = v
	return s
}

func (s *DescribeChannelOverallDataResponseBody) SetOverallData(v *DescribeChannelOverallDataResponseBodyOverallData) *DescribeChannelOverallDataResponseBody {
	s.OverallData = v
	return s
}

func (s *DescribeChannelOverallDataResponseBody) SetRequestId(v string) *DescribeChannelOverallDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeChannelOverallDataResponseBodyCallInfo struct {
	// example:
	//
	// rjdhtnqy
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// IN
	CallStatus *string `json:"CallStatus,omitempty" xml:"CallStatus,omitempty"`
	// example:
	//
	// 123456
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// 1615860711
	CreatedTs *int64 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	// example:
	//
	// 1615860811
	DestroyedTs *int64 `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	// example:
	//
	// 100
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s DescribeChannelOverallDataResponseBodyCallInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelOverallDataResponseBodyCallInfo) GoString() string {
	return s.String()
}

func (s *DescribeChannelOverallDataResponseBodyCallInfo) SetAppId(v string) *DescribeChannelOverallDataResponseBodyCallInfo {
	s.AppId = &v
	return s
}

func (s *DescribeChannelOverallDataResponseBodyCallInfo) SetCallStatus(v string) *DescribeChannelOverallDataResponseBodyCallInfo {
	s.CallStatus = &v
	return s
}

func (s *DescribeChannelOverallDataResponseBodyCallInfo) SetChannelId(v string) *DescribeChannelOverallDataResponseBodyCallInfo {
	s.ChannelId = &v
	return s
}

func (s *DescribeChannelOverallDataResponseBodyCallInfo) SetCreatedTs(v int64) *DescribeChannelOverallDataResponseBodyCallInfo {
	s.CreatedTs = &v
	return s
}

func (s *DescribeChannelOverallDataResponseBodyCallInfo) SetDestroyedTs(v int64) *DescribeChannelOverallDataResponseBodyCallInfo {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeChannelOverallDataResponseBodyCallInfo) SetDuration(v int64) *DescribeChannelOverallDataResponseBodyCallInfo {
	s.Duration = &v
	return s
}

type DescribeChannelOverallDataResponseBodyMetricDatas struct {
	Nodes []*DescribeChannelOverallDataResponseBodyMetricDatasNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// example:
	//
	// CALL_QUALITY
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeChannelOverallDataResponseBodyMetricDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelOverallDataResponseBodyMetricDatas) GoString() string {
	return s.String()
}

func (s *DescribeChannelOverallDataResponseBodyMetricDatas) SetNodes(v []*DescribeChannelOverallDataResponseBodyMetricDatasNodes) *DescribeChannelOverallDataResponseBodyMetricDatas {
	s.Nodes = v
	return s
}

func (s *DescribeChannelOverallDataResponseBodyMetricDatas) SetType(v string) *DescribeChannelOverallDataResponseBodyMetricDatas {
	s.Type = &v
	return s
}

type DescribeChannelOverallDataResponseBodyMetricDatasNodes struct {
	Ext map[string]interface{} `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// example:
	//
	// 1612418625
	X *string `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 123
	Y *string `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DescribeChannelOverallDataResponseBodyMetricDatasNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelOverallDataResponseBodyMetricDatasNodes) GoString() string {
	return s.String()
}

func (s *DescribeChannelOverallDataResponseBodyMetricDatasNodes) SetExt(v map[string]interface{}) *DescribeChannelOverallDataResponseBodyMetricDatasNodes {
	s.Ext = v
	return s
}

func (s *DescribeChannelOverallDataResponseBodyMetricDatasNodes) SetX(v string) *DescribeChannelOverallDataResponseBodyMetricDatasNodes {
	s.X = &v
	return s
}

func (s *DescribeChannelOverallDataResponseBodyMetricDatasNodes) SetY(v string) *DescribeChannelOverallDataResponseBodyMetricDatasNodes {
	s.Y = &v
	return s
}

type DescribeChannelOverallDataResponseBodyOverallData struct {
	// example:
	//
	// 0.5
	ConnAvgTime *float32 `json:"ConnAvgTime,omitempty" xml:"ConnAvgTime,omitempty"`
	// example:
	//
	// 0.91
	FiveSecJoinRate *float32 `json:"FiveSecJoinRate,omitempty" xml:"FiveSecJoinRate,omitempty"`
	// example:
	//
	// 0.02
	TotalAudioStuckRate *float32 `json:"TotalAudioStuckRate,omitempty" xml:"TotalAudioStuckRate,omitempty"`
	// example:
	//
	// 0.02
	TotalVideoStuckRate *float32 `json:"TotalVideoStuckRate,omitempty" xml:"TotalVideoStuckRate,omitempty"`
	// example:
	//
	// 0.02
	TotalVideoVagueRate *float32 `json:"TotalVideoVagueRate,omitempty" xml:"TotalVideoVagueRate,omitempty"`
}

func (s DescribeChannelOverallDataResponseBodyOverallData) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelOverallDataResponseBodyOverallData) GoString() string {
	return s.String()
}

func (s *DescribeChannelOverallDataResponseBodyOverallData) SetConnAvgTime(v float32) *DescribeChannelOverallDataResponseBodyOverallData {
	s.ConnAvgTime = &v
	return s
}

func (s *DescribeChannelOverallDataResponseBodyOverallData) SetFiveSecJoinRate(v float32) *DescribeChannelOverallDataResponseBodyOverallData {
	s.FiveSecJoinRate = &v
	return s
}

func (s *DescribeChannelOverallDataResponseBodyOverallData) SetTotalAudioStuckRate(v float32) *DescribeChannelOverallDataResponseBodyOverallData {
	s.TotalAudioStuckRate = &v
	return s
}

func (s *DescribeChannelOverallDataResponseBodyOverallData) SetTotalVideoStuckRate(v float32) *DescribeChannelOverallDataResponseBodyOverallData {
	s.TotalVideoStuckRate = &v
	return s
}

func (s *DescribeChannelOverallDataResponseBodyOverallData) SetTotalVideoVagueRate(v float32) *DescribeChannelOverallDataResponseBodyOverallData {
	s.TotalVideoVagueRate = &v
	return s
}

type DescribeChannelOverallDataResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeChannelOverallDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeChannelOverallDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelOverallDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeChannelOverallDataResponse) SetHeaders(v map[string]*string) *DescribeChannelOverallDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeChannelOverallDataResponse) SetStatusCode(v int32) *DescribeChannelOverallDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeChannelOverallDataResponse) SetBody(v *DescribeChannelOverallDataResponseBody) *DescribeChannelOverallDataResponse {
	s.Body = v
	return s
}

type DescribeChannelParticipantsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// aec****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testId
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// asc
	Order   *string `json:"Order,omitempty" xml:"Order,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	// example:
	//
	// 6159ba01-6687-4fb2-a831-f0cd8d188648
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1557909133
	Timestamp *int32 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// example:
	//
	// 3
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// example:
	//
	// 1
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeChannelParticipantsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeChannelParticipantsResponse) SetStatusCode(v int32) *DescribeChannelParticipantsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeChannelParticipantsResponse) SetBody(v *DescribeChannelParticipantsResponseBody) *DescribeChannelParticipantsResponse {
	s.Body = v
	return s
}

type DescribeChannelTopPubUserListRequest struct {
	// APP ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// testappid
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1615893133
	CreatedTs *int64 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	// example:
	//
	// 1615893757
	DestroyedTs *int64 `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
}

func (s DescribeChannelTopPubUserListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelTopPubUserListRequest) GoString() string {
	return s.String()
}

func (s *DescribeChannelTopPubUserListRequest) SetAppId(v string) *DescribeChannelTopPubUserListRequest {
	s.AppId = &v
	return s
}

func (s *DescribeChannelTopPubUserListRequest) SetChannelId(v string) *DescribeChannelTopPubUserListRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeChannelTopPubUserListRequest) SetCreatedTs(v int64) *DescribeChannelTopPubUserListRequest {
	s.CreatedTs = &v
	return s
}

func (s *DescribeChannelTopPubUserListRequest) SetDestroyedTs(v int64) *DescribeChannelTopPubUserListRequest {
	s.DestroyedTs = &v
	return s
}

type DescribeChannelTopPubUserListResponseBody struct {
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId            *string                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TopPubUserDetailList []*DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList `json:"TopPubUserDetailList,omitempty" xml:"TopPubUserDetailList,omitempty" type:"Repeated"`
}

func (s DescribeChannelTopPubUserListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelTopPubUserListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChannelTopPubUserListResponseBody) SetRequestId(v string) *DescribeChannelTopPubUserListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeChannelTopPubUserListResponseBody) SetTopPubUserDetailList(v []*DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList) *DescribeChannelTopPubUserListResponseBody {
	s.TopPubUserDetailList = v
	return s
}

type DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList struct {
	// example:
	//
	// 1615893327
	CreatedTs *int64 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	// example:
	//
	// 1615893442
	DestroyedTs *int64 `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	// example:
	//
	// 0
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 浙江省-杭州市
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// example:
	//
	// 0
	OnlineDuration *int64                                                                        `json:"OnlineDuration,omitempty" xml:"OnlineDuration,omitempty"`
	OnlinePeriods  []*DescribeChannelTopPubUserListResponseBodyTopPubUserDetailListOnlinePeriods `json:"OnlinePeriods,omitempty" xml:"OnlinePeriods,omitempty" type:"Repeated"`
	// example:
	//
	// testuserid
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList) GoString() string {
	return s.String()
}

func (s *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList) SetCreatedTs(v int64) *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList {
	s.CreatedTs = &v
	return s
}

func (s *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList) SetDestroyedTs(v int64) *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList) SetDuration(v int64) *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList {
	s.Duration = &v
	return s
}

func (s *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList) SetLocation(v string) *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList {
	s.Location = &v
	return s
}

func (s *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList) SetOnlineDuration(v int64) *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList {
	s.OnlineDuration = &v
	return s
}

func (s *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList) SetOnlinePeriods(v []*DescribeChannelTopPubUserListResponseBodyTopPubUserDetailListOnlinePeriods) *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList {
	s.OnlinePeriods = v
	return s
}

func (s *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList) SetUserId(v string) *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailList {
	s.UserId = &v
	return s
}

type DescribeChannelTopPubUserListResponseBodyTopPubUserDetailListOnlinePeriods struct {
	// example:
	//
	// 1615893327
	JoinTs *int64 `json:"JoinTs,omitempty" xml:"JoinTs,omitempty"`
	// example:
	//
	// 1615893442
	LeaveTs *int64 `json:"LeaveTs,omitempty" xml:"LeaveTs,omitempty"`
}

func (s DescribeChannelTopPubUserListResponseBodyTopPubUserDetailListOnlinePeriods) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelTopPubUserListResponseBodyTopPubUserDetailListOnlinePeriods) GoString() string {
	return s.String()
}

func (s *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailListOnlinePeriods) SetJoinTs(v int64) *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailListOnlinePeriods {
	s.JoinTs = &v
	return s
}

func (s *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailListOnlinePeriods) SetLeaveTs(v int64) *DescribeChannelTopPubUserListResponseBodyTopPubUserDetailListOnlinePeriods {
	s.LeaveTs = &v
	return s
}

type DescribeChannelTopPubUserListResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeChannelTopPubUserListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeChannelTopPubUserListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelTopPubUserListResponse) GoString() string {
	return s.String()
}

func (s *DescribeChannelTopPubUserListResponse) SetHeaders(v map[string]*string) *DescribeChannelTopPubUserListResponse {
	s.Headers = v
	return s
}

func (s *DescribeChannelTopPubUserListResponse) SetStatusCode(v int32) *DescribeChannelTopPubUserListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeChannelTopPubUserListResponse) SetBody(v *DescribeChannelTopPubUserListResponseBody) *DescribeChannelTopPubUserListResponse {
	s.Body = v
	return s
}

type DescribeChannelUserRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1811****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeChannelUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelUserRequest) GoString() string {
	return s.String()
}

func (s *DescribeChannelUserRequest) SetAppId(v string) *DescribeChannelUserRequest {
	s.AppId = &v
	return s
}

func (s *DescribeChannelUserRequest) SetChannelId(v string) *DescribeChannelUserRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeChannelUserRequest) SetUserId(v string) *DescribeChannelUserRequest {
	s.UserId = &v
	return s
}

type DescribeChannelUserResponseBody struct {
	ChannelExist *bool `json:"ChannelExist,omitempty" xml:"ChannelExist,omitempty"`
	InChannel    *bool `json:"InChannel,omitempty" xml:"InChannel,omitempty"`
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CF8
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Sessions  []*DescribeChannelUserResponseBodySessions `json:"Sessions,omitempty" xml:"Sessions,omitempty" type:"Repeated"`
}

func (s DescribeChannelUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelUserResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChannelUserResponseBody) SetChannelExist(v bool) *DescribeChannelUserResponseBody {
	s.ChannelExist = &v
	return s
}

func (s *DescribeChannelUserResponseBody) SetInChannel(v bool) *DescribeChannelUserResponseBody {
	s.InChannel = &v
	return s
}

func (s *DescribeChannelUserResponseBody) SetRequestId(v string) *DescribeChannelUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeChannelUserResponseBody) SetSessions(v []*DescribeChannelUserResponseBodySessions) *DescribeChannelUserResponseBody {
	s.Sessions = v
	return s
}

type DescribeChannelUserResponseBodySessions struct {
	// example:
	//
	// 1557909133
	Joined *int64 `json:"Joined,omitempty" xml:"Joined,omitempty"`
	// example:
	//
	// xa744sxx8rtobgj****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 1811****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeChannelUserResponseBodySessions) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelUserResponseBodySessions) GoString() string {
	return s.String()
}

func (s *DescribeChannelUserResponseBodySessions) SetJoined(v int64) *DescribeChannelUserResponseBodySessions {
	s.Joined = &v
	return s
}

func (s *DescribeChannelUserResponseBodySessions) SetSessionId(v string) *DescribeChannelUserResponseBodySessions {
	s.SessionId = &v
	return s
}

func (s *DescribeChannelUserResponseBodySessions) SetUserId(v string) *DescribeChannelUserResponseBodySessions {
	s.UserId = &v
	return s
}

type DescribeChannelUserResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeChannelUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeChannelUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelUserResponse) GoString() string {
	return s.String()
}

func (s *DescribeChannelUserResponse) SetHeaders(v map[string]*string) *DescribeChannelUserResponse {
	s.Headers = v
	return s
}

func (s *DescribeChannelUserResponse) SetStatusCode(v int32) *DescribeChannelUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeChannelUserResponse) SetBody(v *DescribeChannelUserResponseBody) *DescribeChannelUserResponse {
	s.Body = v
	return s
}

type DescribeChannelUserMetricsRequest struct {
	// APP ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// testappid
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1614936817
	CreatedTs *int64 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	// example:
	//
	// 1614936817
	DestroyedTs *int64 `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
}

func (s DescribeChannelUserMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelUserMetricsRequest) GoString() string {
	return s.String()
}

func (s *DescribeChannelUserMetricsRequest) SetAppId(v string) *DescribeChannelUserMetricsRequest {
	s.AppId = &v
	return s
}

func (s *DescribeChannelUserMetricsRequest) SetChannelId(v string) *DescribeChannelUserMetricsRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeChannelUserMetricsRequest) SetCreatedTs(v int64) *DescribeChannelUserMetricsRequest {
	s.CreatedTs = &v
	return s
}

func (s *DescribeChannelUserMetricsRequest) SetDestroyedTs(v int64) *DescribeChannelUserMetricsRequest {
	s.DestroyedTs = &v
	return s
}

type DescribeChannelUserMetricsResponseBody struct {
	MetricDatas []*DescribeChannelUserMetricsResponseBodyMetricDatas `json:"MetricDatas,omitempty" xml:"MetricDatas,omitempty" type:"Repeated"`
	OverallData *DescribeChannelUserMetricsResponseBodyOverallData   `json:"OverallData,omitempty" xml:"OverallData,omitempty" type:"Struct"`
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeChannelUserMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelUserMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeChannelUserMetricsResponseBody) SetMetricDatas(v []*DescribeChannelUserMetricsResponseBodyMetricDatas) *DescribeChannelUserMetricsResponseBody {
	s.MetricDatas = v
	return s
}

func (s *DescribeChannelUserMetricsResponseBody) SetOverallData(v *DescribeChannelUserMetricsResponseBodyOverallData) *DescribeChannelUserMetricsResponseBody {
	s.OverallData = v
	return s
}

func (s *DescribeChannelUserMetricsResponseBody) SetRequestId(v string) *DescribeChannelUserMetricsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeChannelUserMetricsResponseBodyMetricDatas struct {
	Nodes []*DescribeChannelUserMetricsResponseBodyMetricDatasNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// example:
	//
	// ALL_NUM
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeChannelUserMetricsResponseBodyMetricDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelUserMetricsResponseBodyMetricDatas) GoString() string {
	return s.String()
}

func (s *DescribeChannelUserMetricsResponseBodyMetricDatas) SetNodes(v []*DescribeChannelUserMetricsResponseBodyMetricDatasNodes) *DescribeChannelUserMetricsResponseBodyMetricDatas {
	s.Nodes = v
	return s
}

func (s *DescribeChannelUserMetricsResponseBodyMetricDatas) SetType(v string) *DescribeChannelUserMetricsResponseBodyMetricDatas {
	s.Type = &v
	return s
}

type DescribeChannelUserMetricsResponseBodyMetricDatasNodes struct {
	Ext map[string]interface{} `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// example:
	//
	// 1612418625
	X *string `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 123
	Y *string `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DescribeChannelUserMetricsResponseBodyMetricDatasNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelUserMetricsResponseBodyMetricDatasNodes) GoString() string {
	return s.String()
}

func (s *DescribeChannelUserMetricsResponseBodyMetricDatasNodes) SetExt(v map[string]interface{}) *DescribeChannelUserMetricsResponseBodyMetricDatasNodes {
	s.Ext = v
	return s
}

func (s *DescribeChannelUserMetricsResponseBodyMetricDatasNodes) SetX(v string) *DescribeChannelUserMetricsResponseBodyMetricDatasNodes {
	s.X = &v
	return s
}

func (s *DescribeChannelUserMetricsResponseBodyMetricDatasNodes) SetY(v string) *DescribeChannelUserMetricsResponseBodyMetricDatasNodes {
	s.Y = &v
	return s
}

type DescribeChannelUserMetricsResponseBodyOverallData struct {
	// example:
	//
	// 0
	TotalBadExpNum *int64 `json:"TotalBadExpNum,omitempty" xml:"TotalBadExpNum,omitempty"`
	// example:
	//
	// 0
	TotalJoinFailNum *int64 `json:"TotalJoinFailNum,omitempty" xml:"TotalJoinFailNum,omitempty"`
	// example:
	//
	// 1
	TotalPubUserNum *int64 `json:"TotalPubUserNum,omitempty" xml:"TotalPubUserNum,omitempty"`
	// example:
	//
	// 3
	TotalSubUserNum *int64 `json:"TotalSubUserNum,omitempty" xml:"TotalSubUserNum,omitempty"`
	// example:
	//
	// 5
	TotalUserNum *int64 `json:"TotalUserNum,omitempty" xml:"TotalUserNum,omitempty"`
}

func (s DescribeChannelUserMetricsResponseBodyOverallData) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelUserMetricsResponseBodyOverallData) GoString() string {
	return s.String()
}

func (s *DescribeChannelUserMetricsResponseBodyOverallData) SetTotalBadExpNum(v int64) *DescribeChannelUserMetricsResponseBodyOverallData {
	s.TotalBadExpNum = &v
	return s
}

func (s *DescribeChannelUserMetricsResponseBodyOverallData) SetTotalJoinFailNum(v int64) *DescribeChannelUserMetricsResponseBodyOverallData {
	s.TotalJoinFailNum = &v
	return s
}

func (s *DescribeChannelUserMetricsResponseBodyOverallData) SetTotalPubUserNum(v int64) *DescribeChannelUserMetricsResponseBodyOverallData {
	s.TotalPubUserNum = &v
	return s
}

func (s *DescribeChannelUserMetricsResponseBodyOverallData) SetTotalSubUserNum(v int64) *DescribeChannelUserMetricsResponseBodyOverallData {
	s.TotalSubUserNum = &v
	return s
}

func (s *DescribeChannelUserMetricsResponseBodyOverallData) SetTotalUserNum(v int64) *DescribeChannelUserMetricsResponseBodyOverallData {
	s.TotalUserNum = &v
	return s
}

type DescribeChannelUserMetricsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeChannelUserMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeChannelUserMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeChannelUserMetricsResponse) GoString() string {
	return s.String()
}

func (s *DescribeChannelUserMetricsResponse) SetHeaders(v map[string]*string) *DescribeChannelUserMetricsResponse {
	s.Headers = v
	return s
}

func (s *DescribeChannelUserMetricsResponse) SetStatusCode(v int32) *DescribeChannelUserMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeChannelUserMetricsResponse) SetBody(v *DescribeChannelUserMetricsResponseBody) *DescribeChannelUserMetricsResponse {
	s.Body = v
	return s
}

type DescribeChannelUsersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a2hz****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testId
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
	// example:
	//
	// 1
	ChannelProfile *int32 `json:"ChannelProfile,omitempty" xml:"ChannelProfile,omitempty"`
	// example:
	//
	// 100
	CommTotalNum        *int32    `json:"CommTotalNum,omitempty" xml:"CommTotalNum,omitempty"`
	InteractiveUserList []*string `json:"InteractiveUserList,omitempty" xml:"InteractiveUserList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	InteractiveUserNum *int32 `json:"InteractiveUserNum,omitempty" xml:"InteractiveUserNum,omitempty"`
	// example:
	//
	// true
	IsChannelExist *bool     `json:"IsChannelExist,omitempty" xml:"IsChannelExist,omitempty"`
	LiveUserList   []*string `json:"LiveUserList,omitempty" xml:"LiveUserList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	LiveUserNum *int32 `json:"LiveUserNum,omitempty" xml:"LiveUserNum,omitempty"`
	// example:
	//
	// 6159ba01-6687-4fb2-a831-f0cd8d188648
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1557909133
	Timestamp *int32    `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	UserList  []*string `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeChannelUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeChannelUsersResponse) SetStatusCode(v int32) *DescribeChannelUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeChannelUsersResponse) SetBody(v *DescribeChannelUsersResponseBody) *DescribeChannelUsersResponse {
	s.Body = v
	return s
}

type DescribeEndPointEventListRequest struct {
	// APP ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// testappid
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1614936817
	CreatedTs *int64 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	// example:
	//
	// 1614936817
	DestroyedTs *int64 `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testuserid1,testuserid2
	UserIdList *string `json:"UserIdList,omitempty" xml:"UserIdList,omitempty"`
}

func (s DescribeEndPointEventListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndPointEventListRequest) GoString() string {
	return s.String()
}

func (s *DescribeEndPointEventListRequest) SetAppId(v string) *DescribeEndPointEventListRequest {
	s.AppId = &v
	return s
}

func (s *DescribeEndPointEventListRequest) SetChannelId(v string) *DescribeEndPointEventListRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeEndPointEventListRequest) SetCreatedTs(v int64) *DescribeEndPointEventListRequest {
	s.CreatedTs = &v
	return s
}

func (s *DescribeEndPointEventListRequest) SetDestroyedTs(v int64) *DescribeEndPointEventListRequest {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeEndPointEventListRequest) SetUserIdList(v string) *DescribeEndPointEventListRequest {
	s.UserIdList = &v
	return s
}

type DescribeEndPointEventListResponseBody struct {
	Nodes []*DescribeEndPointEventListResponseBodyNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEndPointEventListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndPointEventListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEndPointEventListResponseBody) SetNodes(v []*DescribeEndPointEventListResponseBodyNodes) *DescribeEndPointEventListResponseBody {
	s.Nodes = v
	return s
}

func (s *DescribeEndPointEventListResponseBody) SetRequestId(v string) *DescribeEndPointEventListResponseBody {
	s.RequestId = &v
	return s
}

type DescribeEndPointEventListResponseBodyNodes struct {
	EventDataItems []*DescribeEndPointEventListResponseBodyNodesEventDataItems `json:"EventDataItems,omitempty" xml:"EventDataItems,omitempty" type:"Repeated"`
	// example:
	//
	// testuserid
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeEndPointEventListResponseBodyNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndPointEventListResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *DescribeEndPointEventListResponseBodyNodes) SetEventDataItems(v []*DescribeEndPointEventListResponseBodyNodesEventDataItems) *DescribeEndPointEventListResponseBodyNodes {
	s.EventDataItems = v
	return s
}

func (s *DescribeEndPointEventListResponseBodyNodes) SetUserId(v string) *DescribeEndPointEventListResponseBodyNodes {
	s.UserId = &v
	return s
}

type DescribeEndPointEventListResponseBodyNodesEventDataItems struct {
	EventList []*DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList `json:"EventList,omitempty" xml:"EventList,omitempty" type:"Repeated"`
	// example:
	//
	// 1614936817
	Ts *int64 `json:"Ts,omitempty" xml:"Ts,omitempty"`
}

func (s DescribeEndPointEventListResponseBodyNodesEventDataItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndPointEventListResponseBodyNodesEventDataItems) GoString() string {
	return s.String()
}

func (s *DescribeEndPointEventListResponseBodyNodesEventDataItems) SetEventList(v []*DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList) *DescribeEndPointEventListResponseBodyNodesEventDataItems {
	s.EventList = v
	return s
}

func (s *DescribeEndPointEventListResponseBodyNodesEventDataItems) SetTs(v int64) *DescribeEndPointEventListResponseBodyNodesEventDataItems {
	s.Ts = &v
	return s
}

type DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList struct {
	// example:
	//
	// 开始发布
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// example:
	//
	// USER
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// example:
	//
	// 1614936817
	Ts *int64 `json:"Ts,omitempty" xml:"Ts,omitempty"`
	// example:
	//
	// 1614936817123
	TsInMs *string `json:"TsInMs,omitempty" xml:"TsInMs,omitempty"`
}

func (s DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList) GoString() string {
	return s.String()
}

func (s *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList) SetEventName(v string) *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList {
	s.EventName = &v
	return s
}

func (s *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList) SetEventType(v string) *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList {
	s.EventType = &v
	return s
}

func (s *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList) SetTs(v int64) *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList {
	s.Ts = &v
	return s
}

func (s *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList) SetTsInMs(v string) *DescribeEndPointEventListResponseBodyNodesEventDataItemsEventList {
	s.TsInMs = &v
	return s
}

type DescribeEndPointEventListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEndPointEventListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEndPointEventListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndPointEventListResponse) GoString() string {
	return s.String()
}

func (s *DescribeEndPointEventListResponse) SetHeaders(v map[string]*string) *DescribeEndPointEventListResponse {
	s.Headers = v
	return s
}

func (s *DescribeEndPointEventListResponse) SetStatusCode(v int32) *DescribeEndPointEventListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEndPointEventListResponse) SetBody(v *DescribeEndPointEventListResponseBody) *DescribeEndPointEventListResponse {
	s.Body = v
	return s
}

type DescribeEndPointMetricDataRequest struct {
	// APP ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// testappid
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1614936817
	CreatedTs *int64 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	// example:
	//
	// 1614936817
	DestroyedTs *int64 `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// VIDEO_STUCK_CAMERA
	Metrics *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	// example:
	//
	// testcall1,testcall2
	PubCallIdList *string `json:"PubCallIdList,omitempty" xml:"PubCallIdList,omitempty"`
	// example:
	//
	// testuserid
	PubUserId *string `json:"PubUserId,omitempty" xml:"PubUserId,omitempty"`
	// example:
	//
	// testuserid
	SubUserId *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s DescribeEndPointMetricDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndPointMetricDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeEndPointMetricDataRequest) SetAppId(v string) *DescribeEndPointMetricDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeEndPointMetricDataRequest) SetChannelId(v string) *DescribeEndPointMetricDataRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeEndPointMetricDataRequest) SetCreatedTs(v int64) *DescribeEndPointMetricDataRequest {
	s.CreatedTs = &v
	return s
}

func (s *DescribeEndPointMetricDataRequest) SetDestroyedTs(v int64) *DescribeEndPointMetricDataRequest {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeEndPointMetricDataRequest) SetMetrics(v string) *DescribeEndPointMetricDataRequest {
	s.Metrics = &v
	return s
}

func (s *DescribeEndPointMetricDataRequest) SetPubCallIdList(v string) *DescribeEndPointMetricDataRequest {
	s.PubCallIdList = &v
	return s
}

func (s *DescribeEndPointMetricDataRequest) SetPubUserId(v string) *DescribeEndPointMetricDataRequest {
	s.PubUserId = &v
	return s
}

func (s *DescribeEndPointMetricDataRequest) SetSubUserId(v string) *DescribeEndPointMetricDataRequest {
	s.SubUserId = &v
	return s
}

type DescribeEndPointMetricDataResponseBody struct {
	PubMetrics []*DescribeEndPointMetricDataResponseBodyPubMetrics `json:"PubMetrics,omitempty" xml:"PubMetrics,omitempty" type:"Repeated"`
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId  *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubMetrics []*DescribeEndPointMetricDataResponseBodySubMetrics `json:"SubMetrics,omitempty" xml:"SubMetrics,omitempty" type:"Repeated"`
}

func (s DescribeEndPointMetricDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndPointMetricDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEndPointMetricDataResponseBody) SetPubMetrics(v []*DescribeEndPointMetricDataResponseBodyPubMetrics) *DescribeEndPointMetricDataResponseBody {
	s.PubMetrics = v
	return s
}

func (s *DescribeEndPointMetricDataResponseBody) SetRequestId(v string) *DescribeEndPointMetricDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEndPointMetricDataResponseBody) SetSubMetrics(v []*DescribeEndPointMetricDataResponseBodySubMetrics) *DescribeEndPointMetricDataResponseBody {
	s.SubMetrics = v
	return s
}

type DescribeEndPointMetricDataResponseBodyPubMetrics struct {
	Nodes []*DescribeEndPointMetricDataResponseBodyPubMetricsNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// example:
	//
	// VIDEO_STUCK_CAMERA
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// testuserid
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeEndPointMetricDataResponseBodyPubMetrics) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndPointMetricDataResponseBodyPubMetrics) GoString() string {
	return s.String()
}

func (s *DescribeEndPointMetricDataResponseBodyPubMetrics) SetNodes(v []*DescribeEndPointMetricDataResponseBodyPubMetricsNodes) *DescribeEndPointMetricDataResponseBodyPubMetrics {
	s.Nodes = v
	return s
}

func (s *DescribeEndPointMetricDataResponseBodyPubMetrics) SetType(v string) *DescribeEndPointMetricDataResponseBodyPubMetrics {
	s.Type = &v
	return s
}

func (s *DescribeEndPointMetricDataResponseBodyPubMetrics) SetUserId(v string) *DescribeEndPointMetricDataResponseBodyPubMetrics {
	s.UserId = &v
	return s
}

type DescribeEndPointMetricDataResponseBodyPubMetricsNodes struct {
	Ext map[string]interface{} `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// example:
	//
	// 1548670257
	X *string `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 230100
	Y *string `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DescribeEndPointMetricDataResponseBodyPubMetricsNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndPointMetricDataResponseBodyPubMetricsNodes) GoString() string {
	return s.String()
}

func (s *DescribeEndPointMetricDataResponseBodyPubMetricsNodes) SetExt(v map[string]interface{}) *DescribeEndPointMetricDataResponseBodyPubMetricsNodes {
	s.Ext = v
	return s
}

func (s *DescribeEndPointMetricDataResponseBodyPubMetricsNodes) SetX(v string) *DescribeEndPointMetricDataResponseBodyPubMetricsNodes {
	s.X = &v
	return s
}

func (s *DescribeEndPointMetricDataResponseBodyPubMetricsNodes) SetY(v string) *DescribeEndPointMetricDataResponseBodyPubMetricsNodes {
	s.Y = &v
	return s
}

type DescribeEndPointMetricDataResponseBodySubMetrics struct {
	Nodes []*DescribeEndPointMetricDataResponseBodySubMetricsNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// example:
	//
	// VIDEO_STUCK_CAMERA
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// testuserid
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeEndPointMetricDataResponseBodySubMetrics) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndPointMetricDataResponseBodySubMetrics) GoString() string {
	return s.String()
}

func (s *DescribeEndPointMetricDataResponseBodySubMetrics) SetNodes(v []*DescribeEndPointMetricDataResponseBodySubMetricsNodes) *DescribeEndPointMetricDataResponseBodySubMetrics {
	s.Nodes = v
	return s
}

func (s *DescribeEndPointMetricDataResponseBodySubMetrics) SetType(v string) *DescribeEndPointMetricDataResponseBodySubMetrics {
	s.Type = &v
	return s
}

func (s *DescribeEndPointMetricDataResponseBodySubMetrics) SetUserId(v string) *DescribeEndPointMetricDataResponseBodySubMetrics {
	s.UserId = &v
	return s
}

type DescribeEndPointMetricDataResponseBodySubMetricsNodes struct {
	Ext map[string]interface{} `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// example:
	//
	// 1548670257
	X *string `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 230100
	Y *string `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DescribeEndPointMetricDataResponseBodySubMetricsNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndPointMetricDataResponseBodySubMetricsNodes) GoString() string {
	return s.String()
}

func (s *DescribeEndPointMetricDataResponseBodySubMetricsNodes) SetExt(v map[string]interface{}) *DescribeEndPointMetricDataResponseBodySubMetricsNodes {
	s.Ext = v
	return s
}

func (s *DescribeEndPointMetricDataResponseBodySubMetricsNodes) SetX(v string) *DescribeEndPointMetricDataResponseBodySubMetricsNodes {
	s.X = &v
	return s
}

func (s *DescribeEndPointMetricDataResponseBodySubMetricsNodes) SetY(v string) *DescribeEndPointMetricDataResponseBodySubMetricsNodes {
	s.Y = &v
	return s
}

type DescribeEndPointMetricDataResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEndPointMetricDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEndPointMetricDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndPointMetricDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeEndPointMetricDataResponse) SetHeaders(v map[string]*string) *DescribeEndPointMetricDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeEndPointMetricDataResponse) SetStatusCode(v int32) *DescribeEndPointMetricDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEndPointMetricDataResponse) SetBody(v *DescribeEndPointMetricDataResponseBody) *DescribeEndPointMetricDataResponse {
	s.Body = v
	return s
}

type DescribeFaultDiagnosisFactorDistributionStatRequest struct {
	// APP ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// 0rbd****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1615892596****
	EndTs *int64 `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1615892596
	StartTs *int64 `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
}

func (s DescribeFaultDiagnosisFactorDistributionStatRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisFactorDistributionStatRequest) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisFactorDistributionStatRequest) SetAppId(v string) *DescribeFaultDiagnosisFactorDistributionStatRequest {
	s.AppId = &v
	return s
}

func (s *DescribeFaultDiagnosisFactorDistributionStatRequest) SetEndTs(v int64) *DescribeFaultDiagnosisFactorDistributionStatRequest {
	s.EndTs = &v
	return s
}

func (s *DescribeFaultDiagnosisFactorDistributionStatRequest) SetStartTs(v int64) *DescribeFaultDiagnosisFactorDistributionStatRequest {
	s.StartTs = &v
	return s
}

type DescribeFaultDiagnosisFactorDistributionStatResponseBody struct {
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StatList  []*DescribeFaultDiagnosisFactorDistributionStatResponseBodyStatList `json:"StatList,omitempty" xml:"StatList,omitempty" type:"Repeated"`
}

func (s DescribeFaultDiagnosisFactorDistributionStatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisFactorDistributionStatResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisFactorDistributionStatResponseBody) SetRequestId(v string) *DescribeFaultDiagnosisFactorDistributionStatResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFaultDiagnosisFactorDistributionStatResponseBody) SetStatList(v []*DescribeFaultDiagnosisFactorDistributionStatResponseBodyStatList) *DescribeFaultDiagnosisFactorDistributionStatResponseBody {
	s.StatList = v
	return s
}

type DescribeFaultDiagnosisFactorDistributionStatResponseBodyStatList struct {
	// example:
	//
	// 1
	FactorId *string `json:"FactorId,omitempty" xml:"FactorId,omitempty"`
	// example:
	//
	// 100
	UserCount *int32 `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
	// example:
	//
	// 0.9239
	UserRatio *float32 `json:"UserRatio,omitempty" xml:"UserRatio,omitempty"`
}

func (s DescribeFaultDiagnosisFactorDistributionStatResponseBodyStatList) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisFactorDistributionStatResponseBodyStatList) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisFactorDistributionStatResponseBodyStatList) SetFactorId(v string) *DescribeFaultDiagnosisFactorDistributionStatResponseBodyStatList {
	s.FactorId = &v
	return s
}

func (s *DescribeFaultDiagnosisFactorDistributionStatResponseBodyStatList) SetUserCount(v int32) *DescribeFaultDiagnosisFactorDistributionStatResponseBodyStatList {
	s.UserCount = &v
	return s
}

func (s *DescribeFaultDiagnosisFactorDistributionStatResponseBodyStatList) SetUserRatio(v float32) *DescribeFaultDiagnosisFactorDistributionStatResponseBodyStatList {
	s.UserRatio = &v
	return s
}

type DescribeFaultDiagnosisFactorDistributionStatResponse struct {
	Headers    map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFaultDiagnosisFactorDistributionStatResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFaultDiagnosisFactorDistributionStatResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisFactorDistributionStatResponse) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisFactorDistributionStatResponse) SetHeaders(v map[string]*string) *DescribeFaultDiagnosisFactorDistributionStatResponse {
	s.Headers = v
	return s
}

func (s *DescribeFaultDiagnosisFactorDistributionStatResponse) SetStatusCode(v int32) *DescribeFaultDiagnosisFactorDistributionStatResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFaultDiagnosisFactorDistributionStatResponse) SetBody(v *DescribeFaultDiagnosisFactorDistributionStatResponseBody) *DescribeFaultDiagnosisFactorDistributionStatResponse {
	s.Body = v
	return s
}

type DescribeFaultDiagnosisOverallDataRequest struct {
	// APP ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 4eah****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1615910399
	EndTs *int64 `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1615824000
	StartTs *int64 `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// JOIN_SLOW_USER
	StatDim *string `json:"StatDim,omitempty" xml:"StatDim,omitempty"`
}

func (s DescribeFaultDiagnosisOverallDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisOverallDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisOverallDataRequest) SetAppId(v string) *DescribeFaultDiagnosisOverallDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeFaultDiagnosisOverallDataRequest) SetEndTs(v int64) *DescribeFaultDiagnosisOverallDataRequest {
	s.EndTs = &v
	return s
}

func (s *DescribeFaultDiagnosisOverallDataRequest) SetStartTs(v int64) *DescribeFaultDiagnosisOverallDataRequest {
	s.StartTs = &v
	return s
}

func (s *DescribeFaultDiagnosisOverallDataRequest) SetStatDim(v string) *DescribeFaultDiagnosisOverallDataRequest {
	s.StatDim = &v
	return s
}

type DescribeFaultDiagnosisOverallDataResponseBody struct {
	MetricData  *DescribeFaultDiagnosisOverallDataResponseBodyMetricData  `json:"MetricData,omitempty" xml:"MetricData,omitempty" type:"Struct"`
	OverallData *DescribeFaultDiagnosisOverallDataResponseBodyOverallData `json:"OverallData,omitempty" xml:"OverallData,omitempty" type:"Struct"`
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFaultDiagnosisOverallDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisOverallDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisOverallDataResponseBody) SetMetricData(v *DescribeFaultDiagnosisOverallDataResponseBodyMetricData) *DescribeFaultDiagnosisOverallDataResponseBody {
	s.MetricData = v
	return s
}

func (s *DescribeFaultDiagnosisOverallDataResponseBody) SetOverallData(v *DescribeFaultDiagnosisOverallDataResponseBodyOverallData) *DescribeFaultDiagnosisOverallDataResponseBody {
	s.OverallData = v
	return s
}

func (s *DescribeFaultDiagnosisOverallDataResponseBody) SetRequestId(v string) *DescribeFaultDiagnosisOverallDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeFaultDiagnosisOverallDataResponseBodyMetricData struct {
	Nodes []*DescribeFaultDiagnosisOverallDataResponseBodyMetricDataNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
}

func (s DescribeFaultDiagnosisOverallDataResponseBodyMetricData) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisOverallDataResponseBodyMetricData) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisOverallDataResponseBodyMetricData) SetNodes(v []*DescribeFaultDiagnosisOverallDataResponseBodyMetricDataNodes) *DescribeFaultDiagnosisOverallDataResponseBodyMetricData {
	s.Nodes = v
	return s
}

type DescribeFaultDiagnosisOverallDataResponseBodyMetricDataNodes struct {
	Ext map[string]interface{} `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// example:
	//
	// 1615824000
	X *string `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 1
	Y *string `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DescribeFaultDiagnosisOverallDataResponseBodyMetricDataNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisOverallDataResponseBodyMetricDataNodes) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisOverallDataResponseBodyMetricDataNodes) SetExt(v map[string]interface{}) *DescribeFaultDiagnosisOverallDataResponseBodyMetricDataNodes {
	s.Ext = v
	return s
}

func (s *DescribeFaultDiagnosisOverallDataResponseBodyMetricDataNodes) SetX(v string) *DescribeFaultDiagnosisOverallDataResponseBodyMetricDataNodes {
	s.X = &v
	return s
}

func (s *DescribeFaultDiagnosisOverallDataResponseBodyMetricDataNodes) SetY(v string) *DescribeFaultDiagnosisOverallDataResponseBodyMetricDataNodes {
	s.Y = &v
	return s
}

type DescribeFaultDiagnosisOverallDataResponseBodyOverallData struct {
	// example:
	//
	// 20
	FaultUserCount *int32 `json:"FaultUserCount,omitempty" xml:"FaultUserCount,omitempty"`
	// example:
	//
	// 0.1
	FaultUserRatio *float32 `json:"FaultUserRatio,omitempty" xml:"FaultUserRatio,omitempty"`
	// example:
	//
	// 40
	TotalUserCount *int32 `json:"TotalUserCount,omitempty" xml:"TotalUserCount,omitempty"`
}

func (s DescribeFaultDiagnosisOverallDataResponseBodyOverallData) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisOverallDataResponseBodyOverallData) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisOverallDataResponseBodyOverallData) SetFaultUserCount(v int32) *DescribeFaultDiagnosisOverallDataResponseBodyOverallData {
	s.FaultUserCount = &v
	return s
}

func (s *DescribeFaultDiagnosisOverallDataResponseBodyOverallData) SetFaultUserRatio(v float32) *DescribeFaultDiagnosisOverallDataResponseBodyOverallData {
	s.FaultUserRatio = &v
	return s
}

func (s *DescribeFaultDiagnosisOverallDataResponseBodyOverallData) SetTotalUserCount(v int32) *DescribeFaultDiagnosisOverallDataResponseBodyOverallData {
	s.TotalUserCount = &v
	return s
}

type DescribeFaultDiagnosisOverallDataResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFaultDiagnosisOverallDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFaultDiagnosisOverallDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisOverallDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisOverallDataResponse) SetHeaders(v map[string]*string) *DescribeFaultDiagnosisOverallDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeFaultDiagnosisOverallDataResponse) SetStatusCode(v int32) *DescribeFaultDiagnosisOverallDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFaultDiagnosisOverallDataResponse) SetBody(v *DescribeFaultDiagnosisOverallDataResponseBody) *DescribeFaultDiagnosisOverallDataResponse {
	s.Body = v
	return s
}

type DescribeFaultDiagnosisUserDetailRequest struct {
	// APP ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// 0rbd****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 311
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1615892596
	CreatedTs *int64 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// JOIN_SLOW
	FaultType *string `json:"FaultType,omitempty" xml:"FaultType,omitempty"`
	// example:
	//
	// true
	QueryCallUserInfo *bool `json:"QueryCallUserInfo,omitempty" xml:"QueryCallUserInfo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c906531af5f9****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeFaultDiagnosisUserDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailRequest) SetAppId(v string) *DescribeFaultDiagnosisUserDetailRequest {
	s.AppId = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailRequest) SetChannelId(v string) *DescribeFaultDiagnosisUserDetailRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailRequest) SetCreatedTs(v int64) *DescribeFaultDiagnosisUserDetailRequest {
	s.CreatedTs = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailRequest) SetFaultType(v string) *DescribeFaultDiagnosisUserDetailRequest {
	s.FaultType = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailRequest) SetQueryCallUserInfo(v bool) *DescribeFaultDiagnosisUserDetailRequest {
	s.QueryCallUserInfo = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailRequest) SetUserId(v string) *DescribeFaultDiagnosisUserDetailRequest {
	s.UserId = &v
	return s
}

type DescribeFaultDiagnosisUserDetailResponseBody struct {
	CallInfo         *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo        `json:"CallInfo,omitempty" xml:"CallInfo,omitempty" type:"Struct"`
	FactorList       []*DescribeFaultDiagnosisUserDetailResponseBodyFactorList    `json:"FactorList,omitempty" xml:"FactorList,omitempty" type:"Repeated"`
	FaultMetricData  *DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricData `json:"FaultMetricData,omitempty" xml:"FaultMetricData,omitempty" type:"Struct"`
	NetworkOperators []*string                                                    `json:"NetworkOperators,omitempty" xml:"NetworkOperators,omitempty" type:"Repeated"`
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId  *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserDetail *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail `json:"UserDetail,omitempty" xml:"UserDetail,omitempty" type:"Struct"`
}

func (s DescribeFaultDiagnosisUserDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailResponseBody) SetCallInfo(v *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo) *DescribeFaultDiagnosisUserDetailResponseBody {
	s.CallInfo = v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBody) SetFactorList(v []*DescribeFaultDiagnosisUserDetailResponseBodyFactorList) *DescribeFaultDiagnosisUserDetailResponseBody {
	s.FactorList = v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBody) SetFaultMetricData(v *DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricData) *DescribeFaultDiagnosisUserDetailResponseBody {
	s.FaultMetricData = v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBody) SetNetworkOperators(v []*string) *DescribeFaultDiagnosisUserDetailResponseBody {
	s.NetworkOperators = v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBody) SetRequestId(v string) *DescribeFaultDiagnosisUserDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBody) SetUserDetail(v *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) *DescribeFaultDiagnosisUserDetailResponseBody {
	s.UserDetail = v
	return s
}

type DescribeFaultDiagnosisUserDetailResponseBodyCallInfo struct {
	// App ID。
	//
	// example:
	//
	// 0rbd****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// IN
	CallStatus *string `json:"CallStatus,omitempty" xml:"CallStatus,omitempty"`
	// example:
	//
	// 311
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// 1620957905
	CreatedTs *int64 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	// example:
	//
	// 1620958150
	DestroyedTs *int64 `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	// example:
	//
	// 100
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyCallInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyCallInfo) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo) SetAppId(v string) *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo {
	s.AppId = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo) SetCallStatus(v string) *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo {
	s.CallStatus = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo) SetChannelId(v string) *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo {
	s.ChannelId = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo) SetCreatedTs(v int64) *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo {
	s.CreatedTs = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo) SetDestroyedTs(v int64) *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo) SetDuration(v int64) *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo {
	s.Duration = &v
	return s
}

type DescribeFaultDiagnosisUserDetailResponseBodyFactorList struct {
	// example:
	//
	// 1
	FactorId *string `json:"FactorId,omitempty" xml:"FactorId,omitempty"`
	// example:
	//
	// LOCAL
	FaultSource        *string                                                                     `json:"FaultSource,omitempty" xml:"FaultSource,omitempty"`
	RelatedEventDatas  []*DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatas  `json:"RelatedEventDatas,omitempty" xml:"RelatedEventDatas,omitempty" type:"Repeated"`
	RelatedMetricDatas []*DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas `json:"RelatedMetricDatas,omitempty" xml:"RelatedMetricDatas,omitempty" type:"Repeated"`
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFactorList) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFactorList) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorList) SetFactorId(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorList {
	s.FactorId = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorList) SetFaultSource(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorList {
	s.FaultSource = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorList) SetRelatedEventDatas(v []*DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatas) *DescribeFaultDiagnosisUserDetailResponseBodyFactorList {
	s.RelatedEventDatas = v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorList) SetRelatedMetricDatas(v []*DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas) *DescribeFaultDiagnosisUserDetailResponseBodyFactorList {
	s.RelatedMetricDatas = v
	return s
}

type DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatas struct {
	EventDataItems []*DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItems `json:"EventDataItems,omitempty" xml:"EventDataItems,omitempty" type:"Repeated"`
	// example:
	//
	// SENDER
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// 0a497933****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatas) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatas) SetEventDataItems(v []*DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItems) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatas {
	s.EventDataItems = v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatas) SetRole(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatas {
	s.Role = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatas) SetUserId(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatas {
	s.UserId = &v
	return s
}

type DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItems struct {
	EventList []*DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList `json:"EventList,omitempty" xml:"EventList,omitempty" type:"Repeated"`
	// example:
	//
	// 1614936817
	Ts *int64 `json:"Ts,omitempty" xml:"Ts,omitempty"`
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItems) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItems) SetEventList(v []*DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItems {
	s.EventList = v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItems) SetTs(v int64) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItems {
	s.Ts = &v
	return s
}

type DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList struct {
	// example:
	//
	// 开始发布
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// example:
	//
	// USER
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// example:
	//
	// 1614936817
	Ts *int64 `json:"Ts,omitempty" xml:"Ts,omitempty"`
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList) SetEventName(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList {
	s.EventName = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList) SetEventType(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList {
	s.EventType = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList) SetTs(v int64) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList {
	s.Ts = &v
	return s
}

type DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas struct {
	Nodes []*DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatasNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// example:
	//
	// SENDER
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// AUDIO_STUCK
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 0a497933****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas) SetNodes(v []*DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatasNodes) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas {
	s.Nodes = v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas) SetRole(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas {
	s.Role = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas) SetType(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas {
	s.Type = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas) SetUserId(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas {
	s.UserId = &v
	return s
}

type DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatasNodes struct {
	Ext map[string]interface{} `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// example:
	//
	// 1615892596
	X *string `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 20
	Y *string `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatasNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatasNodes) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatasNodes) SetExt(v map[string]interface{}) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatasNodes {
	s.Ext = v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatasNodes) SetX(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatasNodes {
	s.X = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatasNodes) SetY(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatasNodes {
	s.Y = &v
	return s
}

type DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricData struct {
	Nodes []*DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricDataNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricData) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricData) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricData) SetNodes(v []*DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricDataNodes) *DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricData {
	s.Nodes = v
	return s
}

type DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricDataNodes struct {
	// example:
	//
	// 1620957900
	X *string `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 0.4540
	Y *string `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricDataNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricDataNodes) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricDataNodes) SetX(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricDataNodes {
	s.X = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricDataNodes) SetY(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricDataNodes {
	s.Y = &v
	return s
}

type DescribeFaultDiagnosisUserDetailResponseBodyUserDetail struct {
	// example:
	//
	// 1620957919
	CreatedTs *int64 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	// example:
	//
	// 1620958150
	DestroyedTs *int64 `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	// example:
	//
	// 231
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 浙江省-杭州市
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// example:
	//
	// 4G
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// example:
	//
	// 231
	OnlineDuration *int64                                                                 `json:"OnlineDuration,omitempty" xml:"OnlineDuration,omitempty"`
	OnlinePeriods  []*DescribeFaultDiagnosisUserDetailResponseBodyUserDetailOnlinePeriods `json:"OnlinePeriods,omitempty" xml:"OnlinePeriods,omitempty" type:"Repeated"`
	// example:
	//
	// iOS
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// example:
	//
	// 1.0.0
	SdkVersion *string `json:"SdkVersion,omitempty" xml:"SdkVersion,omitempty"`
	// example:
	//
	// 0a497933****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) SetCreatedTs(v int64) *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail {
	s.CreatedTs = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) SetDestroyedTs(v int64) *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) SetDuration(v int64) *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail {
	s.Duration = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) SetLocation(v string) *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail {
	s.Location = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) SetNetwork(v string) *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail {
	s.Network = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) SetOnlineDuration(v int64) *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail {
	s.OnlineDuration = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) SetOnlinePeriods(v []*DescribeFaultDiagnosisUserDetailResponseBodyUserDetailOnlinePeriods) *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail {
	s.OnlinePeriods = v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) SetOs(v string) *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail {
	s.Os = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) SetSdkVersion(v string) *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail {
	s.SdkVersion = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) SetUserId(v string) *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail {
	s.UserId = &v
	return s
}

type DescribeFaultDiagnosisUserDetailResponseBodyUserDetailOnlinePeriods struct {
	// example:
	//
	// 1620957919
	JoinTs *int64 `json:"JoinTs,omitempty" xml:"JoinTs,omitempty"`
	// example:
	//
	// 1620958150
	LeaveTs *int64 `json:"LeaveTs,omitempty" xml:"LeaveTs,omitempty"`
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyUserDetailOnlinePeriods) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyUserDetailOnlinePeriods) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetailOnlinePeriods) SetJoinTs(v int64) *DescribeFaultDiagnosisUserDetailResponseBodyUserDetailOnlinePeriods {
	s.JoinTs = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetailOnlinePeriods) SetLeaveTs(v int64) *DescribeFaultDiagnosisUserDetailResponseBodyUserDetailOnlinePeriods {
	s.LeaveTs = &v
	return s
}

type DescribeFaultDiagnosisUserDetailResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFaultDiagnosisUserDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFaultDiagnosisUserDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailResponse) SetHeaders(v map[string]*string) *DescribeFaultDiagnosisUserDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponse) SetStatusCode(v int32) *DescribeFaultDiagnosisUserDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponse) SetBody(v *DescribeFaultDiagnosisUserDetailResponseBody) *DescribeFaultDiagnosisUserDetailResponse {
	s.Body = v
	return s
}

type DescribeFaultDiagnosisUserListRequest struct {
	// APP ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// 0rbd****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 311
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1615892596
	EndTs *int64 `json:"EndTs,omitempty" xml:"EndTs,omitempty"`
	// example:
	//
	// JOIN_SLOW,AUDIO_STUCK
	FaultTypes *string `json:"FaultTypes,omitempty" xml:"FaultTypes,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1615806196
	StartTs *int64 `json:"StartTs,omitempty" xml:"StartTs,omitempty"`
	// example:
	//
	// c906531af5f9****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeFaultDiagnosisUserListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisUserListRequest) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserListRequest) SetAppId(v string) *DescribeFaultDiagnosisUserListRequest {
	s.AppId = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListRequest) SetChannelId(v string) *DescribeFaultDiagnosisUserListRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListRequest) SetEndTs(v int64) *DescribeFaultDiagnosisUserListRequest {
	s.EndTs = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListRequest) SetFaultTypes(v string) *DescribeFaultDiagnosisUserListRequest {
	s.FaultTypes = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListRequest) SetPageNo(v int32) *DescribeFaultDiagnosisUserListRequest {
	s.PageNo = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListRequest) SetPageSize(v int32) *DescribeFaultDiagnosisUserListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListRequest) SetStartTs(v int64) *DescribeFaultDiagnosisUserListRequest {
	s.StartTs = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListRequest) SetUserId(v string) *DescribeFaultDiagnosisUserListRequest {
	s.UserId = &v
	return s
}

type DescribeFaultDiagnosisUserListResponseBody struct {
	// example:
	//
	// 2
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 20
	TotalCnt *int32                                                `json:"TotalCnt,omitempty" xml:"TotalCnt,omitempty"`
	UserList []*DescribeFaultDiagnosisUserListResponseBodyUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
}

func (s DescribeFaultDiagnosisUserListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisUserListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserListResponseBody) SetPageNo(v int32) *DescribeFaultDiagnosisUserListResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListResponseBody) SetPageSize(v int32) *DescribeFaultDiagnosisUserListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListResponseBody) SetRequestId(v string) *DescribeFaultDiagnosisUserListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListResponseBody) SetTotalCnt(v int32) *DescribeFaultDiagnosisUserListResponseBody {
	s.TotalCnt = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListResponseBody) SetUserList(v []*DescribeFaultDiagnosisUserListResponseBodyUserList) *DescribeFaultDiagnosisUserListResponseBody {
	s.UserList = v
	return s
}

type DescribeFaultDiagnosisUserListResponseBodyUserList struct {
	// example:
	//
	// 1614936817
	ChannelCreatedTs *int64 `json:"ChannelCreatedTs,omitempty" xml:"ChannelCreatedTs,omitempty"`
	// example:
	//
	// 904
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// 1614936817
	CreatedTs *int64 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	// example:
	//
	// 1614936817
	DestroyedTs *int64                                                         `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	FaultList   []*DescribeFaultDiagnosisUserListResponseBodyUserListFaultList `json:"FaultList,omitempty" xml:"FaultList,omitempty" type:"Repeated"`
	// example:
	//
	// 123456
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeFaultDiagnosisUserListResponseBodyUserList) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisUserListResponseBodyUserList) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserListResponseBodyUserList) SetChannelCreatedTs(v int64) *DescribeFaultDiagnosisUserListResponseBodyUserList {
	s.ChannelCreatedTs = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListResponseBodyUserList) SetChannelId(v string) *DescribeFaultDiagnosisUserListResponseBodyUserList {
	s.ChannelId = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListResponseBodyUserList) SetCreatedTs(v int64) *DescribeFaultDiagnosisUserListResponseBodyUserList {
	s.CreatedTs = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListResponseBodyUserList) SetDestroyedTs(v int64) *DescribeFaultDiagnosisUserListResponseBodyUserList {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListResponseBodyUserList) SetFaultList(v []*DescribeFaultDiagnosisUserListResponseBodyUserListFaultList) *DescribeFaultDiagnosisUserListResponseBodyUserList {
	s.FaultList = v
	return s
}

func (s *DescribeFaultDiagnosisUserListResponseBodyUserList) SetUserId(v string) *DescribeFaultDiagnosisUserListResponseBodyUserList {
	s.UserId = &v
	return s
}

type DescribeFaultDiagnosisUserListResponseBodyUserListFaultList struct {
	// example:
	//
	// JOIN_SLOW
	FaultType *string `json:"FaultType,omitempty" xml:"FaultType,omitempty"`
}

func (s DescribeFaultDiagnosisUserListResponseBodyUserListFaultList) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisUserListResponseBodyUserListFaultList) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserListResponseBodyUserListFaultList) SetFaultType(v string) *DescribeFaultDiagnosisUserListResponseBodyUserListFaultList {
	s.FaultType = &v
	return s
}

type DescribeFaultDiagnosisUserListResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFaultDiagnosisUserListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFaultDiagnosisUserListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFaultDiagnosisUserListResponse) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserListResponse) SetHeaders(v map[string]*string) *DescribeFaultDiagnosisUserListResponse {
	s.Headers = v
	return s
}

func (s *DescribeFaultDiagnosisUserListResponse) SetStatusCode(v int32) *DescribeFaultDiagnosisUserListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFaultDiagnosisUserListResponse) SetBody(v *DescribeFaultDiagnosisUserListResponseBody) *DescribeFaultDiagnosisUserListResponse {
	s.Body = v
	return s
}

type DescribeMPULayoutInfoListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 2
	LayoutId *int64 `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	// example:
	//
	// LayoutName
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	Layouts *DescribeMPULayoutInfoListResponseBodyLayouts `json:"Layouts,omitempty" xml:"Layouts,omitempty" type:"Struct"`
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// example:
	//
	// 1
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
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
	// example:
	//
	// 3
	AudioMixCount *int32 `json:"AudioMixCount,omitempty" xml:"AudioMixCount,omitempty"`
	// example:
	//
	// 2
	LayoutId *int64 `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	// example:
	//
	// LayoutName
	Name  *string                                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	Panes *DescribeMPULayoutInfoListResponseBodyLayoutsLayoutPanes `json:"Panes,omitempty" xml:"Panes,omitempty" type:"Struct"`
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
	// example:
	//
	// 0.5
	Height *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 0
	MajorPane *int32 `json:"MajorPane,omitempty" xml:"MajorPane,omitempty"`
	// example:
	//
	// 0
	PaneId *int32 `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
	// example:
	//
	// 0.5
	Width *float32 `json:"Width,omitempty" xml:"Width,omitempty"`
	// example:
	//
	// 0.5
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 0.5
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	// example:
	//
	// 0
	ZOrder *int32 `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMPULayoutInfoListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeMPULayoutInfoListResponse) SetStatusCode(v int32) *DescribeMPULayoutInfoListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMPULayoutInfoListResponse) SetBody(v *DescribeMPULayoutInfoListResponseBody) *DescribeMPULayoutInfoListResponse {
	s.Body = v
	return s
}

type DescribePubUserListBySubUserRequest struct {
	// APP ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// testappid
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1614936817
	CreatedTs *int64 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	// example:
	//
	// 1614936817
	DestroyedTs *int64 `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testuserid
	SubUserId *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s DescribePubUserListBySubUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePubUserListBySubUserRequest) GoString() string {
	return s.String()
}

func (s *DescribePubUserListBySubUserRequest) SetAppId(v string) *DescribePubUserListBySubUserRequest {
	s.AppId = &v
	return s
}

func (s *DescribePubUserListBySubUserRequest) SetChannelId(v string) *DescribePubUserListBySubUserRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribePubUserListBySubUserRequest) SetCreatedTs(v int64) *DescribePubUserListBySubUserRequest {
	s.CreatedTs = &v
	return s
}

func (s *DescribePubUserListBySubUserRequest) SetDestroyedTs(v int64) *DescribePubUserListBySubUserRequest {
	s.DestroyedTs = &v
	return s
}

func (s *DescribePubUserListBySubUserRequest) SetSubUserId(v string) *DescribePubUserListBySubUserRequest {
	s.SubUserId = &v
	return s
}

type DescribePubUserListBySubUserResponseBody struct {
	// example:
	//
	// IN
	CallStatus        *string                                                      `json:"CallStatus,omitempty" xml:"CallStatus,omitempty"`
	PubUserDetailList []*DescribePubUserListBySubUserResponseBodyPubUserDetailList `json:"PubUserDetailList,omitempty" xml:"PubUserDetailList,omitempty" type:"Repeated"`
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId     *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubUserDetail *DescribePubUserListBySubUserResponseBodySubUserDetail `json:"SubUserDetail,omitempty" xml:"SubUserDetail,omitempty" type:"Struct"`
}

func (s DescribePubUserListBySubUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePubUserListBySubUserResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePubUserListBySubUserResponseBody) SetCallStatus(v string) *DescribePubUserListBySubUserResponseBody {
	s.CallStatus = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBody) SetPubUserDetailList(v []*DescribePubUserListBySubUserResponseBodyPubUserDetailList) *DescribePubUserListBySubUserResponseBody {
	s.PubUserDetailList = v
	return s
}

func (s *DescribePubUserListBySubUserResponseBody) SetRequestId(v string) *DescribePubUserListBySubUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBody) SetSubUserDetail(v *DescribePubUserListBySubUserResponseBodySubUserDetail) *DescribePubUserListBySubUserResponseBody {
	s.SubUserDetail = v
	return s
}

type DescribePubUserListBySubUserResponseBodyPubUserDetailList struct {
	CallIdList []*string `json:"CallIdList,omitempty" xml:"CallIdList,omitempty" type:"Repeated"`
	// example:
	//
	// NATIVE
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// example:
	//
	// 1614936817
	CreatedTs *int64 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	// example:
	//
	// 1614936817
	DestroyedTs *int64 `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	// example:
	//
	// 0
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 浙江省-杭州市
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// example:
	//
	// 4G
	Network     *string   `json:"Network,omitempty" xml:"Network,omitempty"`
	NetworkList []*string `json:"NetworkList,omitempty" xml:"NetworkList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	OnlineDuration *int64                                                                    `json:"OnlineDuration,omitempty" xml:"OnlineDuration,omitempty"`
	OnlinePeriods  []*DescribePubUserListBySubUserResponseBodyPubUserDetailListOnlinePeriods `json:"OnlinePeriods,omitempty" xml:"OnlinePeriods,omitempty" type:"Repeated"`
	// example:
	//
	// iOS
	Os     *string   `json:"Os,omitempty" xml:"Os,omitempty"`
	OsList []*string `json:"OsList,omitempty" xml:"OsList,omitempty" type:"Repeated"`
	Roles  []*string `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// example:
	//
	// 1.0.0
	SdkVersion     *string   `json:"SdkVersion,omitempty" xml:"SdkVersion,omitempty"`
	SdkVersionList []*string `json:"SdkVersionList,omitempty" xml:"SdkVersionList,omitempty" type:"Repeated"`
	// example:
	//
	// testuserid
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 旁路转推
	UserIdAlias *string `json:"UserIdAlias,omitempty" xml:"UserIdAlias,omitempty"`
}

func (s DescribePubUserListBySubUserResponseBodyPubUserDetailList) String() string {
	return tea.Prettify(s)
}

func (s DescribePubUserListBySubUserResponseBodyPubUserDetailList) GoString() string {
	return s.String()
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetCallIdList(v []*string) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.CallIdList = v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetClientType(v string) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.ClientType = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetCreatedTs(v int64) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.CreatedTs = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetDestroyedTs(v int64) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.DestroyedTs = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetDuration(v int64) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.Duration = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetLocation(v string) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.Location = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetNetwork(v string) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.Network = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetNetworkList(v []*string) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.NetworkList = v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetOnlineDuration(v int64) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.OnlineDuration = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetOnlinePeriods(v []*DescribePubUserListBySubUserResponseBodyPubUserDetailListOnlinePeriods) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.OnlinePeriods = v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetOs(v string) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.Os = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetOsList(v []*string) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.OsList = v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetRoles(v []*string) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.Roles = v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetSdkVersion(v string) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.SdkVersion = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetSdkVersionList(v []*string) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.SdkVersionList = v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetUserId(v string) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.UserId = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailList) SetUserIdAlias(v string) *DescribePubUserListBySubUserResponseBodyPubUserDetailList {
	s.UserIdAlias = &v
	return s
}

type DescribePubUserListBySubUserResponseBodyPubUserDetailListOnlinePeriods struct {
	// example:
	//
	// 1614936817
	JoinTs *int64 `json:"JoinTs,omitempty" xml:"JoinTs,omitempty"`
	// example:
	//
	// 1614936817
	LeaveTs *int64 `json:"LeaveTs,omitempty" xml:"LeaveTs,omitempty"`
}

func (s DescribePubUserListBySubUserResponseBodyPubUserDetailListOnlinePeriods) String() string {
	return tea.Prettify(s)
}

func (s DescribePubUserListBySubUserResponseBodyPubUserDetailListOnlinePeriods) GoString() string {
	return s.String()
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailListOnlinePeriods) SetJoinTs(v int64) *DescribePubUserListBySubUserResponseBodyPubUserDetailListOnlinePeriods {
	s.JoinTs = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodyPubUserDetailListOnlinePeriods) SetLeaveTs(v int64) *DescribePubUserListBySubUserResponseBodyPubUserDetailListOnlinePeriods {
	s.LeaveTs = &v
	return s
}

type DescribePubUserListBySubUserResponseBodySubUserDetail struct {
	// example:
	//
	// NATIVE
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// example:
	//
	// 1614936817
	CreatedTs *int64 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	// example:
	//
	// 1614936817
	DestroyedTs *int64 `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	// example:
	//
	// 0
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 浙江省-杭州市
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// example:
	//
	// 4G
	Network     *string   `json:"Network,omitempty" xml:"Network,omitempty"`
	NetworkList []*string `json:"NetworkList,omitempty" xml:"NetworkList,omitempty" type:"Repeated"`
	// example:
	//
	// 0
	OnlineDuration *int64                                                                `json:"OnlineDuration,omitempty" xml:"OnlineDuration,omitempty"`
	OnlinePeriods  []*DescribePubUserListBySubUserResponseBodySubUserDetailOnlinePeriods `json:"OnlinePeriods,omitempty" xml:"OnlinePeriods,omitempty" type:"Repeated"`
	// example:
	//
	// iOS
	Os     *string   `json:"Os,omitempty" xml:"Os,omitempty"`
	OsList []*string `json:"OsList,omitempty" xml:"OsList,omitempty" type:"Repeated"`
	Roles  []*string `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// example:
	//
	// 1.0.0
	SdkVersion     *string   `json:"SdkVersion,omitempty" xml:"SdkVersion,omitempty"`
	SdkVersionList []*string `json:"SdkVersionList,omitempty" xml:"SdkVersionList,omitempty" type:"Repeated"`
	// example:
	//
	// testuserid
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// 旁路转推
	UserIdAlias *string `json:"UserIdAlias,omitempty" xml:"UserIdAlias,omitempty"`
}

func (s DescribePubUserListBySubUserResponseBodySubUserDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribePubUserListBySubUserResponseBodySubUserDetail) GoString() string {
	return s.String()
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetClientType(v string) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.ClientType = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetCreatedTs(v int64) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.CreatedTs = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetDestroyedTs(v int64) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.DestroyedTs = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetDuration(v int64) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.Duration = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetLocation(v string) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.Location = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetNetwork(v string) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.Network = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetNetworkList(v []*string) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.NetworkList = v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetOnlineDuration(v int64) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.OnlineDuration = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetOnlinePeriods(v []*DescribePubUserListBySubUserResponseBodySubUserDetailOnlinePeriods) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.OnlinePeriods = v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetOs(v string) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.Os = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetOsList(v []*string) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.OsList = v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetRoles(v []*string) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.Roles = v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetSdkVersion(v string) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.SdkVersion = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetSdkVersionList(v []*string) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.SdkVersionList = v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetUserId(v string) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.UserId = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetail) SetUserIdAlias(v string) *DescribePubUserListBySubUserResponseBodySubUserDetail {
	s.UserIdAlias = &v
	return s
}

type DescribePubUserListBySubUserResponseBodySubUserDetailOnlinePeriods struct {
	// example:
	//
	// 1614936817
	JoinTs *int64 `json:"JoinTs,omitempty" xml:"JoinTs,omitempty"`
	// example:
	//
	// 1614936817
	LeaveTs *int64 `json:"LeaveTs,omitempty" xml:"LeaveTs,omitempty"`
}

func (s DescribePubUserListBySubUserResponseBodySubUserDetailOnlinePeriods) String() string {
	return tea.Prettify(s)
}

func (s DescribePubUserListBySubUserResponseBodySubUserDetailOnlinePeriods) GoString() string {
	return s.String()
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetailOnlinePeriods) SetJoinTs(v int64) *DescribePubUserListBySubUserResponseBodySubUserDetailOnlinePeriods {
	s.JoinTs = &v
	return s
}

func (s *DescribePubUserListBySubUserResponseBodySubUserDetailOnlinePeriods) SetLeaveTs(v int64) *DescribePubUserListBySubUserResponseBodySubUserDetailOnlinePeriods {
	s.LeaveTs = &v
	return s
}

type DescribePubUserListBySubUserResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePubUserListBySubUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePubUserListBySubUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePubUserListBySubUserResponse) GoString() string {
	return s.String()
}

func (s *DescribePubUserListBySubUserResponse) SetHeaders(v map[string]*string) *DescribePubUserListBySubUserResponse {
	s.Headers = v
	return s
}

func (s *DescribePubUserListBySubUserResponse) SetStatusCode(v int32) *DescribePubUserListBySubUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePubUserListBySubUserResponse) SetBody(v *DescribePubUserListBySubUserResponseBody) *DescribePubUserListBySubUserResponse {
	s.Body = v
	return s
}

type DescribeQoeMetricDataRequest struct {
	// APP ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// testappid
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1614936817
	CreatedTs *int64 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	// example:
	//
	// 1614936817
	DestroyedTs *int64 `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testuserid
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeQoeMetricDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeQoeMetricDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeQoeMetricDataRequest) SetAppId(v string) *DescribeQoeMetricDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeQoeMetricDataRequest) SetChannelId(v string) *DescribeQoeMetricDataRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeQoeMetricDataRequest) SetCreatedTs(v int64) *DescribeQoeMetricDataRequest {
	s.CreatedTs = &v
	return s
}

func (s *DescribeQoeMetricDataRequest) SetDestroyedTs(v int64) *DescribeQoeMetricDataRequest {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeQoeMetricDataRequest) SetUserId(v string) *DescribeQoeMetricDataRequest {
	s.UserId = &v
	return s
}

type DescribeQoeMetricDataResponseBody struct {
	AudioData []*DescribeQoeMetricDataResponseBodyAudioData `json:"AudioData,omitempty" xml:"AudioData,omitempty" type:"Repeated"`
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VideoData []*DescribeQoeMetricDataResponseBodyVideoData `json:"VideoData,omitempty" xml:"VideoData,omitempty" type:"Repeated"`
}

func (s DescribeQoeMetricDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeQoeMetricDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeQoeMetricDataResponseBody) SetAudioData(v []*DescribeQoeMetricDataResponseBodyAudioData) *DescribeQoeMetricDataResponseBody {
	s.AudioData = v
	return s
}

func (s *DescribeQoeMetricDataResponseBody) SetRequestId(v string) *DescribeQoeMetricDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeQoeMetricDataResponseBody) SetVideoData(v []*DescribeQoeMetricDataResponseBodyVideoData) *DescribeQoeMetricDataResponseBody {
	s.VideoData = v
	return s
}

type DescribeQoeMetricDataResponseBodyAudioData struct {
	Nodes []*DescribeQoeMetricDataResponseBodyAudioDataNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// example:
	//
	// AUDIO
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// testuserid
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeQoeMetricDataResponseBodyAudioData) String() string {
	return tea.Prettify(s)
}

func (s DescribeQoeMetricDataResponseBodyAudioData) GoString() string {
	return s.String()
}

func (s *DescribeQoeMetricDataResponseBodyAudioData) SetNodes(v []*DescribeQoeMetricDataResponseBodyAudioDataNodes) *DescribeQoeMetricDataResponseBodyAudioData {
	s.Nodes = v
	return s
}

func (s *DescribeQoeMetricDataResponseBodyAudioData) SetType(v string) *DescribeQoeMetricDataResponseBodyAudioData {
	s.Type = &v
	return s
}

func (s *DescribeQoeMetricDataResponseBodyAudioData) SetUserId(v string) *DescribeQoeMetricDataResponseBodyAudioData {
	s.UserId = &v
	return s
}

type DescribeQoeMetricDataResponseBodyAudioDataNodes struct {
	// example:
	//
	// 1548670256
	X *string `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 123
	Y *string `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DescribeQoeMetricDataResponseBodyAudioDataNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeQoeMetricDataResponseBodyAudioDataNodes) GoString() string {
	return s.String()
}

func (s *DescribeQoeMetricDataResponseBodyAudioDataNodes) SetX(v string) *DescribeQoeMetricDataResponseBodyAudioDataNodes {
	s.X = &v
	return s
}

func (s *DescribeQoeMetricDataResponseBodyAudioDataNodes) SetY(v string) *DescribeQoeMetricDataResponseBodyAudioDataNodes {
	s.Y = &v
	return s
}

type DescribeQoeMetricDataResponseBodyVideoData struct {
	Nodes []*DescribeQoeMetricDataResponseBodyVideoDataNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// example:
	//
	// VIDEO_CAMERA
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// testuserid
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeQoeMetricDataResponseBodyVideoData) String() string {
	return tea.Prettify(s)
}

func (s DescribeQoeMetricDataResponseBodyVideoData) GoString() string {
	return s.String()
}

func (s *DescribeQoeMetricDataResponseBodyVideoData) SetNodes(v []*DescribeQoeMetricDataResponseBodyVideoDataNodes) *DescribeQoeMetricDataResponseBodyVideoData {
	s.Nodes = v
	return s
}

func (s *DescribeQoeMetricDataResponseBodyVideoData) SetType(v string) *DescribeQoeMetricDataResponseBodyVideoData {
	s.Type = &v
	return s
}

func (s *DescribeQoeMetricDataResponseBodyVideoData) SetUserId(v string) *DescribeQoeMetricDataResponseBodyVideoData {
	s.UserId = &v
	return s
}

type DescribeQoeMetricDataResponseBodyVideoDataNodes struct {
	// example:
	//
	// 1548670256
	X *string `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 123
	Y *string `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DescribeQoeMetricDataResponseBodyVideoDataNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeQoeMetricDataResponseBodyVideoDataNodes) GoString() string {
	return s.String()
}

func (s *DescribeQoeMetricDataResponseBodyVideoDataNodes) SetX(v string) *DescribeQoeMetricDataResponseBodyVideoDataNodes {
	s.X = &v
	return s
}

func (s *DescribeQoeMetricDataResponseBodyVideoDataNodes) SetY(v string) *DescribeQoeMetricDataResponseBodyVideoDataNodes {
	s.Y = &v
	return s
}

type DescribeQoeMetricDataResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeQoeMetricDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeQoeMetricDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeQoeMetricDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeQoeMetricDataResponse) SetHeaders(v map[string]*string) *DescribeQoeMetricDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeQoeMetricDataResponse) SetStatusCode(v int32) *DescribeQoeMetricDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeQoeMetricDataResponse) SetBody(v *DescribeQoeMetricDataResponseBody) *DescribeQoeMetricDataResponse {
	s.Body = v
	return s
}

type DescribeQualityAreaDistributionStatDataRequest struct {
	// APP ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 0rbd****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1615910399
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 中国
	ParentArea *string `json:"ParentArea,omitempty" xml:"ParentArea,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1615824000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeQualityAreaDistributionStatDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeQualityAreaDistributionStatDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeQualityAreaDistributionStatDataRequest) SetAppId(v string) *DescribeQualityAreaDistributionStatDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataRequest) SetEndDate(v int64) *DescribeQualityAreaDistributionStatDataRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataRequest) SetParentArea(v string) *DescribeQualityAreaDistributionStatDataRequest {
	s.ParentArea = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataRequest) SetStartDate(v int64) *DescribeQualityAreaDistributionStatDataRequest {
	s.StartDate = &v
	return s
}

type DescribeQualityAreaDistributionStatDataResponseBody struct {
	QualityStatDataList []*DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList `json:"QualityStatDataList,omitempty" xml:"QualityStatDataList,omitempty" type:"Repeated"`
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeQualityAreaDistributionStatDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeQualityAreaDistributionStatDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeQualityAreaDistributionStatDataResponseBody) SetQualityStatDataList(v []*DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) *DescribeQualityAreaDistributionStatDataResponseBody {
	s.QualityStatDataList = v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponseBody) SetRequestId(v string) *DescribeQualityAreaDistributionStatDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList struct {
	// example:
	//
	// 347
	AudioDelay *int64 `json:"AudioDelay,omitempty" xml:"AudioDelay,omitempty"`
	// example:
	//
	// 0.9933
	AudioHighQualityTransmissionRate *string `json:"AudioHighQualityTransmissionRate,omitempty" xml:"AudioHighQualityTransmissionRate,omitempty"`
	// example:
	//
	// 0.0021
	AudioStuckRate *string `json:"AudioStuckRate,omitempty" xml:"AudioStuckRate,omitempty"`
	// example:
	//
	// 0.6654
	CallDurationRatio *string `json:"CallDurationRatio,omitempty" xml:"CallDurationRatio,omitempty"`
	// example:
	//
	// 0.9338
	JoinChannelSucFiveSecRate *string `json:"JoinChannelSucFiveSecRate,omitempty" xml:"JoinChannelSucFiveSecRate,omitempty"`
	// example:
	//
	// 0.9356
	JoinChannelSucRate *string `json:"JoinChannelSucRate,omitempty" xml:"JoinChannelSucRate,omitempty"`
	// example:
	//
	// 中国_浙江省
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 291
	VideoDelay *int64 `json:"VideoDelay,omitempty" xml:"VideoDelay,omitempty"`
	// example:
	//
	// 1363
	VideoFirstPicDuration *int64 `json:"VideoFirstPicDuration,omitempty" xml:"VideoFirstPicDuration,omitempty"`
	// example:
	//
	// 0.9967
	VideoHighQualityTransmissionRate *string `json:"VideoHighQualityTransmissionRate,omitempty" xml:"VideoHighQualityTransmissionRate,omitempty"`
	// example:
	//
	// 0.0058
	VideoStuckRate *string `json:"VideoStuckRate,omitempty" xml:"VideoStuckRate,omitempty"`
}

func (s DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) String() string {
	return tea.Prettify(s)
}

func (s DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) GoString() string {
	return s.String()
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) SetAudioDelay(v int64) *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList {
	s.AudioDelay = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) SetAudioHighQualityTransmissionRate(v string) *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList {
	s.AudioHighQualityTransmissionRate = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) SetAudioStuckRate(v string) *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList {
	s.AudioStuckRate = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) SetCallDurationRatio(v string) *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList {
	s.CallDurationRatio = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) SetJoinChannelSucFiveSecRate(v string) *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList {
	s.JoinChannelSucFiveSecRate = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) SetJoinChannelSucRate(v string) *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList {
	s.JoinChannelSucRate = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) SetName(v string) *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList {
	s.Name = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) SetVideoDelay(v int64) *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList {
	s.VideoDelay = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) SetVideoFirstPicDuration(v int64) *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList {
	s.VideoFirstPicDuration = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) SetVideoHighQualityTransmissionRate(v string) *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList {
	s.VideoHighQualityTransmissionRate = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList) SetVideoStuckRate(v string) *DescribeQualityAreaDistributionStatDataResponseBodyQualityStatDataList {
	s.VideoStuckRate = &v
	return s
}

type DescribeQualityAreaDistributionStatDataResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeQualityAreaDistributionStatDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeQualityAreaDistributionStatDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeQualityAreaDistributionStatDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeQualityAreaDistributionStatDataResponse) SetHeaders(v map[string]*string) *DescribeQualityAreaDistributionStatDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponse) SetStatusCode(v int32) *DescribeQualityAreaDistributionStatDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeQualityAreaDistributionStatDataResponse) SetBody(v *DescribeQualityAreaDistributionStatDataResponseBody) *DescribeQualityAreaDistributionStatDataResponse {
	s.Body = v
	return s
}

type DescribeQualityDistributionStatDataRequest struct {
	// APP ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 0rbd****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1615910399
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1615824000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CHANNEL_ONLINE
	StatDim *string `json:"StatDim,omitempty" xml:"StatDim,omitempty"`
}

func (s DescribeQualityDistributionStatDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeQualityDistributionStatDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeQualityDistributionStatDataRequest) SetAppId(v string) *DescribeQualityDistributionStatDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeQualityDistributionStatDataRequest) SetEndDate(v int64) *DescribeQualityDistributionStatDataRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeQualityDistributionStatDataRequest) SetStartDate(v int64) *DescribeQualityDistributionStatDataRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeQualityDistributionStatDataRequest) SetStatDim(v string) *DescribeQualityDistributionStatDataRequest {
	s.StatDim = &v
	return s
}

type DescribeQualityDistributionStatDataResponseBody struct {
	QualityStatDataList []*DescribeQualityDistributionStatDataResponseBodyQualityStatDataList `json:"QualityStatDataList,omitempty" xml:"QualityStatDataList,omitempty" type:"Repeated"`
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeQualityDistributionStatDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeQualityDistributionStatDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeQualityDistributionStatDataResponseBody) SetQualityStatDataList(v []*DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) *DescribeQualityDistributionStatDataResponseBody {
	s.QualityStatDataList = v
	return s
}

func (s *DescribeQualityDistributionStatDataResponseBody) SetRequestId(v string) *DescribeQualityDistributionStatDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeQualityDistributionStatDataResponseBodyQualityStatDataList struct {
	// example:
	//
	// 554
	AudioDelay *int64 `json:"AudioDelay,omitempty" xml:"AudioDelay,omitempty"`
	// example:
	//
	// 0.9953
	AudioHighQualityTransmissionRate *string `json:"AudioHighQualityTransmissionRate,omitempty" xml:"AudioHighQualityTransmissionRate,omitempty"`
	// example:
	//
	// 0.0014
	AudioStuckRate *string `json:"AudioStuckRate,omitempty" xml:"AudioStuckRate,omitempty"`
	// example:
	//
	// 1.0000
	CallDurationRatio *string `json:"CallDurationRatio,omitempty" xml:"CallDurationRatio,omitempty"`
	// example:
	//
	// 0.9560
	JoinChannelSucFiveSecRate *string `json:"JoinChannelSucFiveSecRate,omitempty" xml:"JoinChannelSucFiveSecRate,omitempty"`
	// example:
	//
	// 0.9575
	JoinChannelSucRate *string `json:"JoinChannelSucRate,omitempty" xml:"JoinChannelSucRate,omitempty"`
	// example:
	//
	// ONE_TO_FIVE
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 517
	VideoDelay *int64 `json:"VideoDelay,omitempty" xml:"VideoDelay,omitempty"`
	// example:
	//
	// 1299
	VideoFirstPicDuration *int64 `json:"VideoFirstPicDuration,omitempty" xml:"VideoFirstPicDuration,omitempty"`
	// example:
	//
	// 0.9981
	VideoHighQualityTransmissionRate *string `json:"VideoHighQualityTransmissionRate,omitempty" xml:"VideoHighQualityTransmissionRate,omitempty"`
	// example:
	//
	// 0.0264
	VideoStuckRate *string `json:"VideoStuckRate,omitempty" xml:"VideoStuckRate,omitempty"`
}

func (s DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) String() string {
	return tea.Prettify(s)
}

func (s DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) GoString() string {
	return s.String()
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) SetAudioDelay(v int64) *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList {
	s.AudioDelay = &v
	return s
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) SetAudioHighQualityTransmissionRate(v string) *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList {
	s.AudioHighQualityTransmissionRate = &v
	return s
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) SetAudioStuckRate(v string) *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList {
	s.AudioStuckRate = &v
	return s
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) SetCallDurationRatio(v string) *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList {
	s.CallDurationRatio = &v
	return s
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) SetJoinChannelSucFiveSecRate(v string) *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList {
	s.JoinChannelSucFiveSecRate = &v
	return s
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) SetJoinChannelSucRate(v string) *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList {
	s.JoinChannelSucRate = &v
	return s
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) SetName(v string) *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList {
	s.Name = &v
	return s
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) SetVideoDelay(v int64) *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList {
	s.VideoDelay = &v
	return s
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) SetVideoFirstPicDuration(v int64) *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList {
	s.VideoFirstPicDuration = &v
	return s
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) SetVideoHighQualityTransmissionRate(v string) *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList {
	s.VideoHighQualityTransmissionRate = &v
	return s
}

func (s *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList) SetVideoStuckRate(v string) *DescribeQualityDistributionStatDataResponseBodyQualityStatDataList {
	s.VideoStuckRate = &v
	return s
}

type DescribeQualityDistributionStatDataResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeQualityDistributionStatDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeQualityDistributionStatDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeQualityDistributionStatDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeQualityDistributionStatDataResponse) SetHeaders(v map[string]*string) *DescribeQualityDistributionStatDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeQualityDistributionStatDataResponse) SetStatusCode(v int32) *DescribeQualityDistributionStatDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeQualityDistributionStatDataResponse) SetBody(v *DescribeQualityDistributionStatDataResponseBody) *DescribeQualityDistributionStatDataResponse {
	s.Body = v
	return s
}

type DescribeQualityOsSdkVersionDistributionStatDataRequest struct {
	// APP ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 0rbd****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1615910399
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1615824000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeQualityOsSdkVersionDistributionStatDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeQualityOsSdkVersionDistributionStatDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataRequest) SetAppId(v string) *DescribeQualityOsSdkVersionDistributionStatDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataRequest) SetEndDate(v int64) *DescribeQualityOsSdkVersionDistributionStatDataRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataRequest) SetStartDate(v int64) *DescribeQualityOsSdkVersionDistributionStatDataRequest {
	s.StartDate = &v
	return s
}

type DescribeQualityOsSdkVersionDistributionStatDataResponseBody struct {
	QualityOsSdkVersionStatDataList []*DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList `json:"QualityOsSdkVersionStatDataList,omitempty" xml:"QualityOsSdkVersionStatDataList,omitempty" type:"Repeated"`
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeQualityOsSdkVersionDistributionStatDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeQualityOsSdkVersionDistributionStatDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBody) SetQualityOsSdkVersionStatDataList(v []*DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) *DescribeQualityOsSdkVersionDistributionStatDataResponseBody {
	s.QualityOsSdkVersionStatDataList = v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBody) SetRequestId(v string) *DescribeQualityOsSdkVersionDistributionStatDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList struct {
	// example:
	//
	// 248
	AudioDelay *int64 `json:"AudioDelay,omitempty" xml:"AudioDelay,omitempty"`
	// example:
	//
	// 0.9987
	AudioHighQualityTransmissionRate *string `json:"AudioHighQualityTransmissionRate,omitempty" xml:"AudioHighQualityTransmissionRate,omitempty"`
	// example:
	//
	// 0.0011
	AudioStuckRate *string `json:"AudioStuckRate,omitempty" xml:"AudioStuckRate,omitempty"`
	// example:
	//
	// 0.0984
	CallDurationRatio *string `json:"CallDurationRatio,omitempty" xml:"CallDurationRatio,omitempty"`
	// example:
	//
	// 0.8276
	JoinChannelSucFiveSecRate *string `json:"JoinChannelSucFiveSecRate,omitempty" xml:"JoinChannelSucFiveSecRate,omitempty"`
	// example:
	//
	// 0.8276
	JoinChannelSucRate *string `json:"JoinChannelSucRate,omitempty" xml:"JoinChannelSucRate,omitempty"`
	// example:
	//
	// 2.1.0.210316.dev--release/rtcsdk_v2.1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// macOS
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// example:
	//
	// 333
	VideoDelay *int64 `json:"VideoDelay,omitempty" xml:"VideoDelay,omitempty"`
	// example:
	//
	// 5643
	VideoFirstPicDuration *int64 `json:"VideoFirstPicDuration,omitempty" xml:"VideoFirstPicDuration,omitempty"`
	// example:
	//
	// 0.9997
	VideoHighQualityTransmissionRate *string `json:"VideoHighQualityTransmissionRate,omitempty" xml:"VideoHighQualityTransmissionRate,omitempty"`
	// example:
	//
	// 0.0054
	VideoStuckRate *string `json:"VideoStuckRate,omitempty" xml:"VideoStuckRate,omitempty"`
}

func (s DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) String() string {
	return tea.Prettify(s)
}

func (s DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) GoString() string {
	return s.String()
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) SetAudioDelay(v int64) *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList {
	s.AudioDelay = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) SetAudioHighQualityTransmissionRate(v string) *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList {
	s.AudioHighQualityTransmissionRate = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) SetAudioStuckRate(v string) *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList {
	s.AudioStuckRate = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) SetCallDurationRatio(v string) *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList {
	s.CallDurationRatio = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) SetJoinChannelSucFiveSecRate(v string) *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList {
	s.JoinChannelSucFiveSecRate = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) SetJoinChannelSucRate(v string) *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList {
	s.JoinChannelSucRate = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) SetName(v string) *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList {
	s.Name = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) SetOs(v string) *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList {
	s.Os = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) SetVideoDelay(v int64) *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList {
	s.VideoDelay = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) SetVideoFirstPicDuration(v int64) *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList {
	s.VideoFirstPicDuration = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) SetVideoHighQualityTransmissionRate(v string) *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList {
	s.VideoHighQualityTransmissionRate = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList) SetVideoStuckRate(v string) *DescribeQualityOsSdkVersionDistributionStatDataResponseBodyQualityOsSdkVersionStatDataList {
	s.VideoStuckRate = &v
	return s
}

type DescribeQualityOsSdkVersionDistributionStatDataResponse struct {
	Headers    map[string]*string                                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeQualityOsSdkVersionDistributionStatDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeQualityOsSdkVersionDistributionStatDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeQualityOsSdkVersionDistributionStatDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponse) SetHeaders(v map[string]*string) *DescribeQualityOsSdkVersionDistributionStatDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponse) SetStatusCode(v int32) *DescribeQualityOsSdkVersionDistributionStatDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeQualityOsSdkVersionDistributionStatDataResponse) SetBody(v *DescribeQualityOsSdkVersionDistributionStatDataResponseBody) *DescribeQualityOsSdkVersionDistributionStatDataResponse {
	s.Body = v
	return s
}

type DescribeQualityOverallDataRequest struct {
	// APP ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 0rbd****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1615910399
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1615824000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// JOIN_CHANNEL_SUC_RATE
	Types *string `json:"Types,omitempty" xml:"Types,omitempty"`
}

func (s DescribeQualityOverallDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeQualityOverallDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeQualityOverallDataRequest) SetAppId(v string) *DescribeQualityOverallDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeQualityOverallDataRequest) SetEndDate(v int64) *DescribeQualityOverallDataRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeQualityOverallDataRequest) SetStartDate(v int64) *DescribeQualityOverallDataRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeQualityOverallDataRequest) SetTypes(v string) *DescribeQualityOverallDataRequest {
	s.Types = &v
	return s
}

type DescribeQualityOverallDataResponseBody struct {
	QualityOverallData []*DescribeQualityOverallDataResponseBodyQualityOverallData `json:"QualityOverallData,omitempty" xml:"QualityOverallData,omitempty" type:"Repeated"`
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeQualityOverallDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeQualityOverallDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeQualityOverallDataResponseBody) SetQualityOverallData(v []*DescribeQualityOverallDataResponseBodyQualityOverallData) *DescribeQualityOverallDataResponseBody {
	s.QualityOverallData = v
	return s
}

func (s *DescribeQualityOverallDataResponseBody) SetRequestId(v string) *DescribeQualityOverallDataResponseBody {
	s.RequestId = &v
	return s
}

type DescribeQualityOverallDataResponseBodyQualityOverallData struct {
	// example:
	//
	// 0.9376
	Average *string                                                          `json:"Average,omitempty" xml:"Average,omitempty"`
	Nodes   []*DescribeQualityOverallDataResponseBodyQualityOverallDataNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// example:
	//
	// JOIN_CHANNEL_SUC_RATE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeQualityOverallDataResponseBodyQualityOverallData) String() string {
	return tea.Prettify(s)
}

func (s DescribeQualityOverallDataResponseBodyQualityOverallData) GoString() string {
	return s.String()
}

func (s *DescribeQualityOverallDataResponseBodyQualityOverallData) SetAverage(v string) *DescribeQualityOverallDataResponseBodyQualityOverallData {
	s.Average = &v
	return s
}

func (s *DescribeQualityOverallDataResponseBodyQualityOverallData) SetNodes(v []*DescribeQualityOverallDataResponseBodyQualityOverallDataNodes) *DescribeQualityOverallDataResponseBodyQualityOverallData {
	s.Nodes = v
	return s
}

func (s *DescribeQualityOverallDataResponseBodyQualityOverallData) SetType(v string) *DescribeQualityOverallDataResponseBodyQualityOverallData {
	s.Type = &v
	return s
}

type DescribeQualityOverallDataResponseBodyQualityOverallDataNodes struct {
	// example:
	//
	// 1615831200
	X *string `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 1.0000
	Y *string `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DescribeQualityOverallDataResponseBodyQualityOverallDataNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeQualityOverallDataResponseBodyQualityOverallDataNodes) GoString() string {
	return s.String()
}

func (s *DescribeQualityOverallDataResponseBodyQualityOverallDataNodes) SetX(v string) *DescribeQualityOverallDataResponseBodyQualityOverallDataNodes {
	s.X = &v
	return s
}

func (s *DescribeQualityOverallDataResponseBodyQualityOverallDataNodes) SetY(v string) *DescribeQualityOverallDataResponseBodyQualityOverallDataNodes {
	s.Y = &v
	return s
}

type DescribeQualityOverallDataResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeQualityOverallDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeQualityOverallDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeQualityOverallDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeQualityOverallDataResponse) SetHeaders(v map[string]*string) *DescribeQualityOverallDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeQualityOverallDataResponse) SetStatusCode(v int32) *DescribeQualityOverallDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeQualityOverallDataResponse) SetBody(v *DescribeQualityOverallDataResponseBody) *DescribeQualityOverallDataResponse {
	s.Body = v
	return s
}

type DescribeRecordFilesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// yourChannelId
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// 2020-11-02T17:36:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 2020-11-01T17:36:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// yourTaskId
	TaskIds []*string `json:"TaskIds,omitempty" xml:"TaskIds,omitempty" type:"Repeated"`
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
	// example:
	//
	// 760bad53276431c499e30dc36f6b****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// example:
	//
	// 1
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
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
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// yourChannelId
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// 2020-10-02T17:36:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1800
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 2020-11-01T17:36:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 2020-11-02T17:36:00Z
	StopTime *string `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
	// example:
	//
	// yourTaskId
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Url    *string `json:"Url,omitempty" xml:"Url,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRecordFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeRecordFilesResponse) SetStatusCode(v int32) *DescribeRecordFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRecordFilesResponse) SetBody(v *DescribeRecordFilesResponseBody) *DescribeRecordFilesResponse {
	s.Body = v
	return s
}

type DescribeRecordTemplatesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 1
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 76dasgb****
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRecordTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeRecordTemplatesResponse) SetStatusCode(v int32) *DescribeRecordTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRecordTemplatesResponse) SetBody(v *DescribeRecordTemplatesResponseBody) *DescribeRecordTemplatesResponse {
	s.Body = v
	return s
}

type DescribeRtcChannelListRequest struct {
	// example:
	//
	// aoe****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// testChannel
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn
	ServiceArea *string `json:"ServiceArea,omitempty" xml:"ServiceArea,omitempty"`
	// example:
	//
	// desc
	SortType *string `json:"SortType,omitempty" xml:"SortType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2018-01-29T00:00:00Z
	TimePoint *string `json:"TimePoint,omitempty" xml:"TimePoint,omitempty"`
	// example:
	//
	// testUser
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeRtcChannelListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelListRequest) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelListRequest) SetAppId(v string) *DescribeRtcChannelListRequest {
	s.AppId = &v
	return s
}

func (s *DescribeRtcChannelListRequest) SetChannelId(v string) *DescribeRtcChannelListRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeRtcChannelListRequest) SetOwnerId(v int64) *DescribeRtcChannelListRequest {
	s.OwnerId = &v
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

func (s *DescribeRtcChannelListRequest) SetServiceArea(v string) *DescribeRtcChannelListRequest {
	s.ServiceArea = &v
	return s
}

func (s *DescribeRtcChannelListRequest) SetSortType(v string) *DescribeRtcChannelListRequest {
	s.SortType = &v
	return s
}

func (s *DescribeRtcChannelListRequest) SetTimePoint(v string) *DescribeRtcChannelListRequest {
	s.TimePoint = &v
	return s
}

func (s *DescribeRtcChannelListRequest) SetUserId(v string) *DescribeRtcChannelListRequest {
	s.UserId = &v
	return s
}

type DescribeRtcChannelListResponseBody struct {
	ChannelList *DescribeRtcChannelListResponseBodyChannelList `json:"ChannelList,omitempty" xml:"ChannelList,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 100
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1000
	TotalCnt *int64 `json:"TotalCnt,omitempty" xml:"TotalCnt,omitempty"`
}

func (s DescribeRtcChannelListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelListResponseBody) SetChannelList(v *DescribeRtcChannelListResponseBodyChannelList) *DescribeRtcChannelListResponseBody {
	s.ChannelList = v
	return s
}

func (s *DescribeRtcChannelListResponseBody) SetPageNo(v int64) *DescribeRtcChannelListResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribeRtcChannelListResponseBody) SetPageSize(v int64) *DescribeRtcChannelListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRtcChannelListResponseBody) SetRequestId(v string) *DescribeRtcChannelListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRtcChannelListResponseBody) SetTotalCnt(v int64) *DescribeRtcChannelListResponseBody {
	s.TotalCnt = &v
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
	CallArea *DescribeRtcChannelListResponseBodyChannelListChannelListCallArea `json:"CallArea,omitempty" xml:"CallArea,omitempty" type:"Struct"`
	// example:
	//
	// testChannel
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// 2018-01-29T02:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 2018-01-29T01:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 2
	TotalUserCnt *int64 `json:"TotalUserCnt,omitempty" xml:"TotalUserCnt,omitempty"`
}

func (s DescribeRtcChannelListResponseBodyChannelListChannelList) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelListResponseBodyChannelListChannelList) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelListResponseBodyChannelListChannelList) SetCallArea(v *DescribeRtcChannelListResponseBodyChannelListChannelListCallArea) *DescribeRtcChannelListResponseBodyChannelListChannelList {
	s.CallArea = v
	return s
}

func (s *DescribeRtcChannelListResponseBodyChannelListChannelList) SetChannelId(v string) *DescribeRtcChannelListResponseBodyChannelListChannelList {
	s.ChannelId = &v
	return s
}

func (s *DescribeRtcChannelListResponseBodyChannelListChannelList) SetEndTime(v string) *DescribeRtcChannelListResponseBodyChannelListChannelList {
	s.EndTime = &v
	return s
}

func (s *DescribeRtcChannelListResponseBodyChannelListChannelList) SetStartTime(v string) *DescribeRtcChannelListResponseBodyChannelListChannelList {
	s.StartTime = &v
	return s
}

func (s *DescribeRtcChannelListResponseBodyChannelListChannelList) SetTotalUserCnt(v int64) *DescribeRtcChannelListResponseBodyChannelListChannelList {
	s.TotalUserCnt = &v
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRtcChannelListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeRtcChannelListResponse) SetStatusCode(v int32) *DescribeRtcChannelListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRtcChannelListResponse) SetBody(v *DescribeRtcChannelListResponseBody) *DescribeRtcChannelListResponse {
	s.Body = v
	return s
}

type DescribeRtcChannelMetricRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// aoe****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testId
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2018-01-29T00:00:00Z
	TimePoint *string `json:"TimePoint,omitempty" xml:"TimePoint,omitempty"`
}

func (s DescribeRtcChannelMetricRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelMetricRequest) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelMetricRequest) SetAppId(v string) *DescribeRtcChannelMetricRequest {
	s.AppId = &v
	return s
}

func (s *DescribeRtcChannelMetricRequest) SetChannelId(v string) *DescribeRtcChannelMetricRequest {
	s.ChannelId = &v
	return s
}

func (s *DescribeRtcChannelMetricRequest) SetOwnerId(v int64) *DescribeRtcChannelMetricRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRtcChannelMetricRequest) SetTimePoint(v string) *DescribeRtcChannelMetricRequest {
	s.TimePoint = &v
	return s
}

type DescribeRtcChannelMetricResponseBody struct {
	ChannelMetricInfo *DescribeRtcChannelMetricResponseBodyChannelMetricInfo `json:"ChannelMetricInfo,omitempty" xml:"ChannelMetricInfo,omitempty" type:"Struct"`
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRtcChannelMetricResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelMetricResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelMetricResponseBody) SetChannelMetricInfo(v *DescribeRtcChannelMetricResponseBodyChannelMetricInfo) *DescribeRtcChannelMetricResponseBody {
	s.ChannelMetricInfo = v
	return s
}

func (s *DescribeRtcChannelMetricResponseBody) SetRequestId(v string) *DescribeRtcChannelMetricResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRtcChannelMetricResponseBodyChannelMetricInfo struct {
	ChannelMetric *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric `json:"ChannelMetric,omitempty" xml:"ChannelMetric,omitempty" type:"Struct"`
	Duration      *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDuration      `json:"Duration,omitempty" xml:"Duration,omitempty" type:"Struct"`
}

func (s DescribeRtcChannelMetricResponseBodyChannelMetricInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelMetricResponseBodyChannelMetricInfo) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfo) SetChannelMetric(v *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric) *DescribeRtcChannelMetricResponseBodyChannelMetricInfo {
	s.ChannelMetric = v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfo) SetDuration(v *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDuration) *DescribeRtcChannelMetricResponseBodyChannelMetricInfo {
	s.Duration = v
	return s
}

type DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric struct {
	// example:
	//
	// example_channel
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// 2019-06-06T18:57:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 10
	PubUserCount *int32 `json:"PubUserCount,omitempty" xml:"PubUserCount,omitempty"`
	// example:
	//
	// 2019-06-06T17:57:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 25
	SubUserCount *int32 `json:"SubUserCount,omitempty" xml:"SubUserCount,omitempty"`
	// example:
	//
	// 30
	UserCount *int32 `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
}

func (s DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric) SetChannelId(v string) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric {
	s.ChannelId = &v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric) SetEndTime(v string) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric {
	s.EndTime = &v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric) SetPubUserCount(v int32) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric {
	s.PubUserCount = &v
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

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric) SetUserCount(v int32) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoChannelMetric {
	s.UserCount = &v
	return s
}

type DescribeRtcChannelMetricResponseBodyChannelMetricInfoDuration struct {
	PubDuration *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration `json:"PubDuration,omitempty" xml:"PubDuration,omitempty" type:"Struct"`
	SubDuration *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration `json:"SubDuration,omitempty" xml:"SubDuration,omitempty" type:"Struct"`
}

func (s DescribeRtcChannelMetricResponseBodyChannelMetricInfoDuration) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelMetricResponseBodyChannelMetricInfoDuration) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDuration) SetPubDuration(v *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDuration {
	s.PubDuration = v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDuration) SetSubDuration(v *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDuration {
	s.SubDuration = v
	return s
}

type DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration struct {
	// example:
	//
	// 100
	Audio *int32 `json:"Audio,omitempty" xml:"Audio,omitempty"`
	// example:
	//
	// 100
	Content *int32 `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 100
	Video1080 *int32 `json:"Video1080,omitempty" xml:"Video1080,omitempty"`
	// example:
	//
	// 100
	Video360 *int32 `json:"Video360,omitempty" xml:"Video360,omitempty"`
	// example:
	//
	// 100
	Video720 *int32 `json:"Video720,omitempty" xml:"Video720,omitempty"`
}

func (s DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration) SetAudio(v int32) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration {
	s.Audio = &v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration) SetContent(v int32) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration {
	s.Content = &v
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

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration) SetVideo720(v int32) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationPubDuration {
	s.Video720 = &v
	return s
}

type DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration struct {
	// example:
	//
	// 100
	Audio *int32 `json:"Audio,omitempty" xml:"Audio,omitempty"`
	// example:
	//
	// 100
	Content *int32 `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 100
	Video1080 *int32 `json:"Video1080,omitempty" xml:"Video1080,omitempty"`
	// example:
	//
	// 100
	Video360 *int32 `json:"Video360,omitempty" xml:"Video360,omitempty"`
	// example:
	//
	// 100
	Video720 *int32 `json:"Video720,omitempty" xml:"Video720,omitempty"`
}

func (s DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration) GoString() string {
	return s.String()
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration) SetAudio(v int32) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration {
	s.Audio = &v
	return s
}

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration) SetContent(v int32) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration {
	s.Content = &v
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

func (s *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration) SetVideo720(v int32) *DescribeRtcChannelMetricResponseBodyChannelMetricInfoDurationSubDuration {
	s.Video720 = &v
	return s
}

type DescribeRtcChannelMetricResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRtcChannelMetricResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeRtcChannelMetricResponse) SetStatusCode(v int32) *DescribeRtcChannelMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRtcChannelMetricResponse) SetBody(v *DescribeRtcChannelMetricResponseBody) *DescribeRtcChannelMetricResponse {
	s.Body = v
	return s
}

type DescribeRtcDurationDataRequest struct {
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 2020-02-04T07:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 3600
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// CN
	ServiceArea *string `json:"ServiceArea,omitempty" xml:"ServiceArea,omitempty"`
	// example:
	//
	// 2020-02-04T05:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeRtcDurationDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcDurationDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeRtcDurationDataRequest) SetAppId(v string) *DescribeRtcDurationDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeRtcDurationDataRequest) SetEndTime(v string) *DescribeRtcDurationDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRtcDurationDataRequest) SetInterval(v string) *DescribeRtcDurationDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeRtcDurationDataRequest) SetOwnerId(v int64) *DescribeRtcDurationDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRtcDurationDataRequest) SetServiceArea(v string) *DescribeRtcDurationDataRequest {
	s.ServiceArea = &v
	return s
}

func (s *DescribeRtcDurationDataRequest) SetStartTime(v string) *DescribeRtcDurationDataRequest {
	s.StartTime = &v
	return s
}

type DescribeRtcDurationDataResponseBody struct {
	DurationDataPerInterval *DescribeRtcDurationDataResponseBodyDurationDataPerInterval `json:"DurationDataPerInterval,omitempty" xml:"DurationDataPerInterval,omitempty" type:"Struct"`
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRtcDurationDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcDurationDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRtcDurationDataResponseBody) SetDurationDataPerInterval(v *DescribeRtcDurationDataResponseBodyDurationDataPerInterval) *DescribeRtcDurationDataResponseBody {
	s.DurationDataPerInterval = v
	return s
}

func (s *DescribeRtcDurationDataResponseBody) SetRequestId(v string) *DescribeRtcDurationDataResponseBody {
	s.RequestId = &v
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
	// example:
	//
	// 200
	AudioDuration *int64 `json:"AudioDuration,omitempty" xml:"AudioDuration,omitempty"`
	// example:
	//
	// 200
	ContentDuration *int64 `json:"ContentDuration,omitempty" xml:"ContentDuration,omitempty"`
	// example:
	//
	// 2020-02-04T05:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// example:
	//
	// 1000
	TotalDuration *int64 `json:"TotalDuration,omitempty" xml:"TotalDuration,omitempty"`
	// example:
	//
	// 300
	V1080Duration *int64 `json:"V1080Duration,omitempty" xml:"V1080Duration,omitempty"`
	// example:
	//
	// 300
	V360Duration *int64 `json:"V360Duration,omitempty" xml:"V360Duration,omitempty"`
	// example:
	//
	// 200
	V720Duration *int64 `json:"V720Duration,omitempty" xml:"V720Duration,omitempty"`
}

func (s DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule) GoString() string {
	return s.String()
}

func (s *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule) SetAudioDuration(v int64) *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule {
	s.AudioDuration = &v
	return s
}

func (s *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule) SetContentDuration(v int64) *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule {
	s.ContentDuration = &v
	return s
}

func (s *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule) SetTimeStamp(v string) *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule {
	s.TimeStamp = &v
	return s
}

func (s *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule) SetTotalDuration(v int64) *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule {
	s.TotalDuration = &v
	return s
}

func (s *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule) SetV1080Duration(v int64) *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule {
	s.V1080Duration = &v
	return s
}

func (s *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule) SetV360Duration(v int64) *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule {
	s.V360Duration = &v
	return s
}

func (s *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule) SetV720Duration(v int64) *DescribeRtcDurationDataResponseBodyDurationDataPerIntervalDurationModule {
	s.V720Duration = &v
	return s
}

type DescribeRtcDurationDataResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRtcDurationDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeRtcDurationDataResponse) SetStatusCode(v int32) *DescribeRtcDurationDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRtcDurationDataResponse) SetBody(v *DescribeRtcDurationDataResponseBody) *DescribeRtcDurationDataResponse {
	s.Body = v
	return s
}

type DescribeRtcPeakChannelCntDataRequest struct {
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 2018-01-29T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 3600
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// CN
	ServiceArea *string `json:"ServiceArea,omitempty" xml:"ServiceArea,omitempty"`
	// example:
	//
	// 2018-01-29T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeRtcPeakChannelCntDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcPeakChannelCntDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeRtcPeakChannelCntDataRequest) SetAppId(v string) *DescribeRtcPeakChannelCntDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeRtcPeakChannelCntDataRequest) SetEndTime(v string) *DescribeRtcPeakChannelCntDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRtcPeakChannelCntDataRequest) SetInterval(v string) *DescribeRtcPeakChannelCntDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeRtcPeakChannelCntDataRequest) SetOwnerId(v int64) *DescribeRtcPeakChannelCntDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRtcPeakChannelCntDataRequest) SetServiceArea(v string) *DescribeRtcPeakChannelCntDataRequest {
	s.ServiceArea = &v
	return s
}

func (s *DescribeRtcPeakChannelCntDataRequest) SetStartTime(v string) *DescribeRtcPeakChannelCntDataRequest {
	s.StartTime = &v
	return s
}

type DescribeRtcPeakChannelCntDataResponseBody struct {
	PeakChannelCntDataPerInterval *DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerInterval `json:"PeakChannelCntDataPerInterval,omitempty" xml:"PeakChannelCntDataPerInterval,omitempty" type:"Struct"`
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// example:
	//
	// 10
	ActiveChannelPeak *int64 `json:"ActiveChannelPeak,omitempty" xml:"ActiveChannelPeak,omitempty"`
	// example:
	//
	// 2018-01-29T00:01:00Z
	ActiveChannelPeakTime *string `json:"ActiveChannelPeakTime,omitempty" xml:"ActiveChannelPeakTime,omitempty"`
	// example:
	//
	// 2018-01-29T00:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerIntervalPeakChannelCntModule) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerIntervalPeakChannelCntModule) GoString() string {
	return s.String()
}

func (s *DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerIntervalPeakChannelCntModule) SetActiveChannelPeak(v int64) *DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerIntervalPeakChannelCntModule {
	s.ActiveChannelPeak = &v
	return s
}

func (s *DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerIntervalPeakChannelCntModule) SetActiveChannelPeakTime(v string) *DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerIntervalPeakChannelCntModule {
	s.ActiveChannelPeakTime = &v
	return s
}

func (s *DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerIntervalPeakChannelCntModule) SetTimeStamp(v string) *DescribeRtcPeakChannelCntDataResponseBodyPeakChannelCntDataPerIntervalPeakChannelCntModule {
	s.TimeStamp = &v
	return s
}

type DescribeRtcPeakChannelCntDataResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRtcPeakChannelCntDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeRtcPeakChannelCntDataResponse) SetStatusCode(v int32) *DescribeRtcPeakChannelCntDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRtcPeakChannelCntDataResponse) SetBody(v *DescribeRtcPeakChannelCntDataResponseBody) *DescribeRtcPeakChannelCntDataResponse {
	s.Body = v
	return s
}

type DescribeRtcUserCntDataRequest struct {
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 2018-01-29T01:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 3600
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// CN
	ServiceArea *string `json:"ServiceArea,omitempty" xml:"ServiceArea,omitempty"`
	// example:
	//
	// 2018-01-29T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeRtcUserCntDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRtcUserCntDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeRtcUserCntDataRequest) SetAppId(v string) *DescribeRtcUserCntDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeRtcUserCntDataRequest) SetEndTime(v string) *DescribeRtcUserCntDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRtcUserCntDataRequest) SetInterval(v string) *DescribeRtcUserCntDataRequest {
	s.Interval = &v
	return s
}

func (s *DescribeRtcUserCntDataRequest) SetOwnerId(v int64) *DescribeRtcUserCntDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRtcUserCntDataRequest) SetServiceArea(v string) *DescribeRtcUserCntDataRequest {
	s.ServiceArea = &v
	return s
}

func (s *DescribeRtcUserCntDataRequest) SetStartTime(v string) *DescribeRtcUserCntDataRequest {
	s.StartTime = &v
	return s
}

type DescribeRtcUserCntDataResponseBody struct {
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
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
	// example:
	//
	// 10
	ActiveUserCnt *int64 `json:"ActiveUserCnt,omitempty" xml:"ActiveUserCnt,omitempty"`
	// example:
	//
	// 2018-01-29T00:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRtcUserCntDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeRtcUserCntDataResponse) SetStatusCode(v int32) *DescribeRtcUserCntDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRtcUserCntDataResponse) SetBody(v *DescribeRtcUserCntDataResponseBody) *DescribeRtcUserCntDataResponse {
	s.Body = v
	return s
}

type DescribeUsageAreaDistributionStatDataRequest struct {
	// APP ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 0rbd****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1615910399
	EndDate    *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	ParentArea *string `json:"ParentArea,omitempty" xml:"ParentArea,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1615824000
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeUsageAreaDistributionStatDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsageAreaDistributionStatDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeUsageAreaDistributionStatDataRequest) SetAppId(v string) *DescribeUsageAreaDistributionStatDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeUsageAreaDistributionStatDataRequest) SetEndDate(v string) *DescribeUsageAreaDistributionStatDataRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeUsageAreaDistributionStatDataRequest) SetParentArea(v string) *DescribeUsageAreaDistributionStatDataRequest {
	s.ParentArea = &v
	return s
}

func (s *DescribeUsageAreaDistributionStatDataRequest) SetStartDate(v string) *DescribeUsageAreaDistributionStatDataRequest {
	s.StartDate = &v
	return s
}

type DescribeUsageAreaDistributionStatDataResponseBody struct {
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId         *string                                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UsageAreaStatList []*DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList `json:"UsageAreaStatList,omitempty" xml:"UsageAreaStatList,omitempty" type:"Repeated"`
}

func (s DescribeUsageAreaDistributionStatDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsageAreaDistributionStatDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUsageAreaDistributionStatDataResponseBody) SetRequestId(v string) *DescribeUsageAreaDistributionStatDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUsageAreaDistributionStatDataResponseBody) SetUsageAreaStatList(v []*DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList) *DescribeUsageAreaDistributionStatDataResponseBody {
	s.UsageAreaStatList = v
	return s
}

type DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList struct {
	// example:
	//
	// 45
	AudioCallDuration *int32 `json:"AudioCallDuration,omitempty" xml:"AudioCallDuration,omitempty"`
	// example:
	//
	// 中国
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 4821
	TotalCallDuration *int32 `json:"TotalCallDuration,omitempty" xml:"TotalCallDuration,omitempty"`
	// example:
	//
	// 4776
	VideoCallDuration *int32 `json:"VideoCallDuration,omitempty" xml:"VideoCallDuration,omitempty"`
}

func (s DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList) GoString() string {
	return s.String()
}

func (s *DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList) SetAudioCallDuration(v int32) *DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList {
	s.AudioCallDuration = &v
	return s
}

func (s *DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList) SetName(v string) *DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList {
	s.Name = &v
	return s
}

func (s *DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList) SetTotalCallDuration(v int32) *DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList {
	s.TotalCallDuration = &v
	return s
}

func (s *DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList) SetVideoCallDuration(v int32) *DescribeUsageAreaDistributionStatDataResponseBodyUsageAreaStatList {
	s.VideoCallDuration = &v
	return s
}

type DescribeUsageAreaDistributionStatDataResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUsageAreaDistributionStatDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUsageAreaDistributionStatDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsageAreaDistributionStatDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeUsageAreaDistributionStatDataResponse) SetHeaders(v map[string]*string) *DescribeUsageAreaDistributionStatDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeUsageAreaDistributionStatDataResponse) SetStatusCode(v int32) *DescribeUsageAreaDistributionStatDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUsageAreaDistributionStatDataResponse) SetBody(v *DescribeUsageAreaDistributionStatDataResponseBody) *DescribeUsageAreaDistributionStatDataResponse {
	s.Body = v
	return s
}

type DescribeUsageDistributionStatDataRequest struct {
	// APP ID
	//
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1615910399
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1615824000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CHANNEL_ONLINE
	StatDim *string `json:"StatDim,omitempty" xml:"StatDim,omitempty"`
}

func (s DescribeUsageDistributionStatDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsageDistributionStatDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeUsageDistributionStatDataRequest) SetAppId(v string) *DescribeUsageDistributionStatDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeUsageDistributionStatDataRequest) SetEndDate(v int64) *DescribeUsageDistributionStatDataRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeUsageDistributionStatDataRequest) SetStartDate(v int64) *DescribeUsageDistributionStatDataRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeUsageDistributionStatDataRequest) SetStatDim(v string) *DescribeUsageDistributionStatDataRequest {
	s.StatDim = &v
	return s
}

type DescribeUsageDistributionStatDataResponseBody struct {
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId     *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UsageStatList []*DescribeUsageDistributionStatDataResponseBodyUsageStatList `json:"UsageStatList,omitempty" xml:"UsageStatList,omitempty" type:"Repeated"`
}

func (s DescribeUsageDistributionStatDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsageDistributionStatDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUsageDistributionStatDataResponseBody) SetRequestId(v string) *DescribeUsageDistributionStatDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUsageDistributionStatDataResponseBody) SetUsageStatList(v []*DescribeUsageDistributionStatDataResponseBodyUsageStatList) *DescribeUsageDistributionStatDataResponseBody {
	s.UsageStatList = v
	return s
}

type DescribeUsageDistributionStatDataResponseBodyUsageStatList struct {
	// example:
	//
	// 51
	AudioCallDuration *int64 `json:"AudioCallDuration,omitempty" xml:"AudioCallDuration,omitempty"`
	// example:
	//
	// 0.9782
	CallDurationRatio *string `json:"CallDurationRatio,omitempty" xml:"CallDurationRatio,omitempty"`
	// example:
	//
	// ONE_TO_FIVE
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 10636
	TotalCallDuration *int64 `json:"TotalCallDuration,omitempty" xml:"TotalCallDuration,omitempty"`
	// example:
	//
	// 10585
	VideoCallDuration *int64 `json:"VideoCallDuration,omitempty" xml:"VideoCallDuration,omitempty"`
}

func (s DescribeUsageDistributionStatDataResponseBodyUsageStatList) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsageDistributionStatDataResponseBodyUsageStatList) GoString() string {
	return s.String()
}

func (s *DescribeUsageDistributionStatDataResponseBodyUsageStatList) SetAudioCallDuration(v int64) *DescribeUsageDistributionStatDataResponseBodyUsageStatList {
	s.AudioCallDuration = &v
	return s
}

func (s *DescribeUsageDistributionStatDataResponseBodyUsageStatList) SetCallDurationRatio(v string) *DescribeUsageDistributionStatDataResponseBodyUsageStatList {
	s.CallDurationRatio = &v
	return s
}

func (s *DescribeUsageDistributionStatDataResponseBodyUsageStatList) SetName(v string) *DescribeUsageDistributionStatDataResponseBodyUsageStatList {
	s.Name = &v
	return s
}

func (s *DescribeUsageDistributionStatDataResponseBodyUsageStatList) SetTotalCallDuration(v int64) *DescribeUsageDistributionStatDataResponseBodyUsageStatList {
	s.TotalCallDuration = &v
	return s
}

func (s *DescribeUsageDistributionStatDataResponseBodyUsageStatList) SetVideoCallDuration(v int64) *DescribeUsageDistributionStatDataResponseBodyUsageStatList {
	s.VideoCallDuration = &v
	return s
}

type DescribeUsageDistributionStatDataResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUsageDistributionStatDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUsageDistributionStatDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsageDistributionStatDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeUsageDistributionStatDataResponse) SetHeaders(v map[string]*string) *DescribeUsageDistributionStatDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeUsageDistributionStatDataResponse) SetStatusCode(v int32) *DescribeUsageDistributionStatDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUsageDistributionStatDataResponse) SetBody(v *DescribeUsageDistributionStatDataResponseBody) *DescribeUsageDistributionStatDataResponse {
	s.Body = v
	return s
}

type DescribeUsageOsSdkVersionDistributionStatDataRequest struct {
	// APP ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 0rbd****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1615910399
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1615824000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
}

func (s DescribeUsageOsSdkVersionDistributionStatDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsageOsSdkVersionDistributionStatDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataRequest) SetAppId(v string) *DescribeUsageOsSdkVersionDistributionStatDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataRequest) SetEndDate(v int64) *DescribeUsageOsSdkVersionDistributionStatDataRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataRequest) SetStartDate(v int64) *DescribeUsageOsSdkVersionDistributionStatDataRequest {
	s.StartDate = &v
	return s
}

type DescribeUsageOsSdkVersionDistributionStatDataResponseBody struct {
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId                 *string                                                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UsageOsSdkVersionStatList []*DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList `json:"UsageOsSdkVersionStatList,omitempty" xml:"UsageOsSdkVersionStatList,omitempty" type:"Repeated"`
}

func (s DescribeUsageOsSdkVersionDistributionStatDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsageOsSdkVersionDistributionStatDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponseBody) SetRequestId(v string) *DescribeUsageOsSdkVersionDistributionStatDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponseBody) SetUsageOsSdkVersionStatList(v []*DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList) *DescribeUsageOsSdkVersionDistributionStatDataResponseBody {
	s.UsageOsSdkVersionStatList = v
	return s
}

type DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList struct {
	// example:
	//
	// 3
	AudioCallDuration *int64 `json:"AudioCallDuration,omitempty" xml:"AudioCallDuration,omitempty"`
	// example:
	//
	// 0.0984
	CallDurationRatio *string `json:"CallDurationRatio,omitempty" xml:"CallDurationRatio,omitempty"`
	// example:
	//
	// 1.0.0
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// macOS
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// example:
	//
	// 476
	TotalCallDuration *int64 `json:"TotalCallDuration,omitempty" xml:"TotalCallDuration,omitempty"`
	// example:
	//
	// 473
	VideoCallDuration *int64 `json:"VideoCallDuration,omitempty" xml:"VideoCallDuration,omitempty"`
}

func (s DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList) GoString() string {
	return s.String()
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList) SetAudioCallDuration(v int64) *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList {
	s.AudioCallDuration = &v
	return s
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList) SetCallDurationRatio(v string) *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList {
	s.CallDurationRatio = &v
	return s
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList) SetName(v string) *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList {
	s.Name = &v
	return s
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList) SetOs(v string) *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList {
	s.Os = &v
	return s
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList) SetTotalCallDuration(v int64) *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList {
	s.TotalCallDuration = &v
	return s
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList) SetVideoCallDuration(v int64) *DescribeUsageOsSdkVersionDistributionStatDataResponseBodyUsageOsSdkVersionStatList {
	s.VideoCallDuration = &v
	return s
}

type DescribeUsageOsSdkVersionDistributionStatDataResponse struct {
	Headers    map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUsageOsSdkVersionDistributionStatDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUsageOsSdkVersionDistributionStatDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsageOsSdkVersionDistributionStatDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponse) SetHeaders(v map[string]*string) *DescribeUsageOsSdkVersionDistributionStatDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponse) SetStatusCode(v int32) *DescribeUsageOsSdkVersionDistributionStatDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponse) SetBody(v *DescribeUsageOsSdkVersionDistributionStatDataResponseBody) *DescribeUsageOsSdkVersionDistributionStatDataResponse {
	s.Body = v
	return s
}

type DescribeUsageOverallDataRequest struct {
	// APP ID
	//
	// This parameter is required.
	//
	// example:
	//
	// a2hz****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1615910399
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1615824000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ONLINE_USER_PEAK
	Types *string `json:"Types,omitempty" xml:"Types,omitempty"`
}

func (s DescribeUsageOverallDataRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsageOverallDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeUsageOverallDataRequest) SetAppId(v string) *DescribeUsageOverallDataRequest {
	s.AppId = &v
	return s
}

func (s *DescribeUsageOverallDataRequest) SetEndDate(v int64) *DescribeUsageOverallDataRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeUsageOverallDataRequest) SetStartDate(v int64) *DescribeUsageOverallDataRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeUsageOverallDataRequest) SetTypes(v string) *DescribeUsageOverallDataRequest {
	s.Types = &v
	return s
}

type DescribeUsageOverallDataResponseBody struct {
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId        *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UsageOverallData []*DescribeUsageOverallDataResponseBodyUsageOverallData `json:"UsageOverallData,omitempty" xml:"UsageOverallData,omitempty" type:"Repeated"`
}

func (s DescribeUsageOverallDataResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsageOverallDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUsageOverallDataResponseBody) SetRequestId(v string) *DescribeUsageOverallDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUsageOverallDataResponseBody) SetUsageOverallData(v []*DescribeUsageOverallDataResponseBodyUsageOverallData) *DescribeUsageOverallDataResponseBody {
	s.UsageOverallData = v
	return s
}

type DescribeUsageOverallDataResponseBodyUsageOverallData struct {
	Nodes []*DescribeUsageOverallDataResponseBodyUsageOverallDataNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// example:
	//
	// ONLINE_USER_PEAK
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeUsageOverallDataResponseBodyUsageOverallData) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsageOverallDataResponseBodyUsageOverallData) GoString() string {
	return s.String()
}

func (s *DescribeUsageOverallDataResponseBodyUsageOverallData) SetNodes(v []*DescribeUsageOverallDataResponseBodyUsageOverallDataNodes) *DescribeUsageOverallDataResponseBodyUsageOverallData {
	s.Nodes = v
	return s
}

func (s *DescribeUsageOverallDataResponseBodyUsageOverallData) SetType(v string) *DescribeUsageOverallDataResponseBodyUsageOverallData {
	s.Type = &v
	return s
}

type DescribeUsageOverallDataResponseBodyUsageOverallDataNodes struct {
	// example:
	//
	// 1615824000
	X *string `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 1
	Y *string `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DescribeUsageOverallDataResponseBodyUsageOverallDataNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsageOverallDataResponseBodyUsageOverallDataNodes) GoString() string {
	return s.String()
}

func (s *DescribeUsageOverallDataResponseBodyUsageOverallDataNodes) SetX(v string) *DescribeUsageOverallDataResponseBodyUsageOverallDataNodes {
	s.X = &v
	return s
}

func (s *DescribeUsageOverallDataResponseBodyUsageOverallDataNodes) SetY(v string) *DescribeUsageOverallDataResponseBodyUsageOverallDataNodes {
	s.Y = &v
	return s
}

type DescribeUsageOverallDataResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUsageOverallDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUsageOverallDataResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUsageOverallDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeUsageOverallDataResponse) SetHeaders(v map[string]*string) *DescribeUsageOverallDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeUsageOverallDataResponse) SetStatusCode(v int32) *DescribeUsageOverallDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUsageOverallDataResponse) SetBody(v *DescribeUsageOverallDataResponseBody) *DescribeUsageOverallDataResponse {
	s.Body = v
	return s
}

type DescribeUserInfoInChannelRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4eah****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testId
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// example:
	//
	// true
	IsChannelExist *bool `json:"IsChannelExist,omitempty" xml:"IsChannelExist,omitempty"`
	// example:
	//
	// true
	IsInChannel *bool                                            `json:"IsInChannel,omitempty" xml:"IsInChannel,omitempty"`
	Property    []*DescribeUserInfoInChannelResponseBodyProperty `json:"Property,omitempty" xml:"Property,omitempty" type:"Repeated"`
	// example:
	//
	// 6159ba01-6687-4fb2-a831-f0cd8d188648
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1557909133
	Timestamp *int32 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
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
	// example:
	//
	// 1557909133
	Join *int32 `json:"Join,omitempty" xml:"Join,omitempty"`
	// example:
	//
	// 1
	Role *int32 `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// xa744sxx8rtobgj****
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserInfoInChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DescribeUserInfoInChannelResponse) SetStatusCode(v int32) *DescribeUserInfoInChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserInfoInChannelResponse) SetBody(v *DescribeUserInfoInChannelResponseBody) *DescribeUserInfoInChannelResponse {
	s.Body = v
	return s
}

type DisableAutoLiveStreamRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
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
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableAutoLiveStreamRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DisableAutoLiveStreamRuleResponse) SetStatusCode(v int32) *DisableAutoLiveStreamRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableAutoLiveStreamRuleResponse) SetBody(v *DisableAutoLiveStreamRuleResponseBody) *DisableAutoLiveStreamRuleResponse {
	s.Body = v
	return s
}

type EnableAutoLiveStreamRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
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
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableAutoLiveStreamRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *EnableAutoLiveStreamRuleResponse) SetStatusCode(v int32) *EnableAutoLiveStreamRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableAutoLiveStreamRuleResponse) SetBody(v *EnableAutoLiveStreamRuleResponseBody) *EnableAutoLiveStreamRuleResponse {
	s.Body = v
	return s
}

type GetMPUTaskStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// yourTaskId
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMPUTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *GetMPUTaskStatusResponse) SetStatusCode(v int32) *GetMPUTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMPUTaskStatusResponse) SetBody(v *GetMPUTaskStatusResponseBody) *GetMPUTaskStatusResponse {
	s.Body = v
	return s
}

type ModifyAppRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ioeh****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// defaultName
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s ModifyAppRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppRequest) SetAppId(v string) *ModifyAppRequest {
	s.AppId = &v
	return s
}

func (s *ModifyAppRequest) SetAppName(v string) *ModifyAppRequest {
	s.AppName = &v
	return s
}

func (s *ModifyAppRequest) SetOwnerId(v int64) *ModifyAppRequest {
	s.OwnerId = &v
	return s
}

type ModifyAppResponseBody struct {
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ModifyAppResponse) SetStatusCode(v int32) *ModifyAppResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAppResponse) SetBody(v *ModifyAppResponseBody) *ModifyAppResponse {
	s.Body = v
	return s
}

type ModifyAppStreamingOutTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// wv7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	StreamingOutTemplate *ModifyAppStreamingOutTemplateRequestStreamingOutTemplate `json:"StreamingOutTemplate,omitempty" xml:"StreamingOutTemplate,omitempty" type:"Struct"`
}

func (s ModifyAppStreamingOutTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppStreamingOutTemplateRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppStreamingOutTemplateRequest) SetAppId(v string) *ModifyAppStreamingOutTemplateRequest {
	s.AppId = &v
	return s
}

func (s *ModifyAppStreamingOutTemplateRequest) SetStreamingOutTemplate(v *ModifyAppStreamingOutTemplateRequestStreamingOutTemplate) *ModifyAppStreamingOutTemplateRequest {
	s.StreamingOutTemplate = v
	return s
}

type ModifyAppStreamingOutTemplateRequestStreamingOutTemplate struct {
	// This parameter is required.
	LayoutIds []*string `json:"LayoutIds,omitempty" xml:"LayoutIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	MediaEncode *int32 `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 模版
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ModifyAppStreamingOutTemplateRequestStreamingOutTemplate) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppStreamingOutTemplateRequestStreamingOutTemplate) GoString() string {
	return s.String()
}

func (s *ModifyAppStreamingOutTemplateRequestStreamingOutTemplate) SetLayoutIds(v []*string) *ModifyAppStreamingOutTemplateRequestStreamingOutTemplate {
	s.LayoutIds = v
	return s
}

func (s *ModifyAppStreamingOutTemplateRequestStreamingOutTemplate) SetMediaEncode(v int32) *ModifyAppStreamingOutTemplateRequestStreamingOutTemplate {
	s.MediaEncode = &v
	return s
}

func (s *ModifyAppStreamingOutTemplateRequestStreamingOutTemplate) SetName(v string) *ModifyAppStreamingOutTemplateRequestStreamingOutTemplate {
	s.Name = &v
	return s
}

func (s *ModifyAppStreamingOutTemplateRequestStreamingOutTemplate) SetTemplateId(v string) *ModifyAppStreamingOutTemplateRequestStreamingOutTemplate {
	s.TemplateId = &v
	return s
}

type ModifyAppStreamingOutTemplateShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// wv7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	StreamingOutTemplateShrink *string `json:"StreamingOutTemplate,omitempty" xml:"StreamingOutTemplate,omitempty"`
}

func (s ModifyAppStreamingOutTemplateShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppStreamingOutTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppStreamingOutTemplateShrinkRequest) SetAppId(v string) *ModifyAppStreamingOutTemplateShrinkRequest {
	s.AppId = &v
	return s
}

func (s *ModifyAppStreamingOutTemplateShrinkRequest) SetStreamingOutTemplateShrink(v string) *ModifyAppStreamingOutTemplateShrinkRequest {
	s.StreamingOutTemplateShrink = &v
	return s
}

type ModifyAppStreamingOutTemplateResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 30D41049-D02D-1C21-86AE-B3E5FD805C27
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// ac7N****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ModifyAppStreamingOutTemplateResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppStreamingOutTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAppStreamingOutTemplateResponseBody) SetRequestId(v string) *ModifyAppStreamingOutTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAppStreamingOutTemplateResponseBody) SetTemplateId(v string) *ModifyAppStreamingOutTemplateResponseBody {
	s.TemplateId = &v
	return s
}

type ModifyAppStreamingOutTemplateResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAppStreamingOutTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAppStreamingOutTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppStreamingOutTemplateResponse) GoString() string {
	return s.String()
}

func (s *ModifyAppStreamingOutTemplateResponse) SetHeaders(v map[string]*string) *ModifyAppStreamingOutTemplateResponse {
	s.Headers = v
	return s
}

func (s *ModifyAppStreamingOutTemplateResponse) SetStatusCode(v int32) *ModifyAppStreamingOutTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAppStreamingOutTemplateResponse) SetBody(v *ModifyAppStreamingOutTemplateResponseBody) *ModifyAppStreamingOutTemplateResponse {
	s.Body = v
	return s
}

type ModifyMPULayoutRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 3
	AudioMixCount *int32 `json:"AudioMixCount,omitempty" xml:"AudioMixCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10117
	LayoutId *int64 `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	// example:
	//
	// LayoutName
	Name    *string                        `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64                         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Panes   []*ModifyMPULayoutRequestPanes `json:"Panes,omitempty" xml:"Panes,omitempty" type:"Repeated"`
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
	// example:
	//
	// 0.2456
	Height *float32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 0
	MajorPane *int32 `json:"MajorPane,omitempty" xml:"MajorPane,omitempty"`
	// example:
	//
	// 0
	PaneId *int32 `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
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
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyMPULayoutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ModifyMPULayoutResponse) SetStatusCode(v int32) *ModifyMPULayoutResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMPULayoutResponse) SetBody(v *ModifyMPULayoutResponseBody) *ModifyMPULayoutResponse {
	s.Body = v
	return s
}

type RemoveTerminalsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// yourChannelId
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1811xxxx
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
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
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
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1811****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// Success
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveTerminalsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *RemoveTerminalsResponse) SetStatusCode(v int32) *RemoveTerminalsResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveTerminalsResponse) SetBody(v *RemoveTerminalsResponseBody) *RemoveTerminalsResponse {
	s.Body = v
	return s
}

type RemoveUsersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
	Users []*RemoveUsersRequestUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s RemoveUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveUsersRequest) GoString() string {
	return s.String()
}

func (s *RemoveUsersRequest) SetAppId(v string) *RemoveUsersRequest {
	s.AppId = &v
	return s
}

func (s *RemoveUsersRequest) SetChannelId(v string) *RemoveUsersRequest {
	s.ChannelId = &v
	return s
}

func (s *RemoveUsersRequest) SetUsers(v []*RemoveUsersRequestUsers) *RemoveUsersRequest {
	s.Users = v
	return s
}

type RemoveUsersRequestUsers struct {
	// This parameter is required.
	//
	// example:
	//
	// 1811****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s RemoveUsersRequestUsers) String() string {
	return tea.Prettify(s)
}

func (s RemoveUsersRequestUsers) GoString() string {
	return s.String()
}

func (s *RemoveUsersRequestUsers) SetUserId(v string) *RemoveUsersRequestUsers {
	s.UserId = &v
	return s
}

type RemoveUsersResponseBody struct {
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CF8
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Users     []*RemoveUsersResponseBodyUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s RemoveUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveUsersResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveUsersResponseBody) SetRequestId(v string) *RemoveUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveUsersResponseBody) SetUsers(v []*RemoveUsersResponseBodyUsers) *RemoveUsersResponseBody {
	s.Users = v
	return s
}

type RemoveUsersResponseBodyUsers struct {
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1811****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s RemoveUsersResponseBodyUsers) String() string {
	return tea.Prettify(s)
}

func (s RemoveUsersResponseBodyUsers) GoString() string {
	return s.String()
}

func (s *RemoveUsersResponseBodyUsers) SetCode(v int32) *RemoveUsersResponseBodyUsers {
	s.Code = &v
	return s
}

func (s *RemoveUsersResponseBodyUsers) SetMessage(v string) *RemoveUsersResponseBodyUsers {
	s.Message = &v
	return s
}

func (s *RemoveUsersResponseBodyUsers) SetUserId(v string) *RemoveUsersResponseBodyUsers {
	s.UserId = &v
	return s
}

type RemoveUsersResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveUsersResponse) GoString() string {
	return s.String()
}

func (s *RemoveUsersResponse) SetHeaders(v map[string]*string) *RemoveUsersResponse {
	s.Headers = v
	return s
}

func (s *RemoveUsersResponse) SetStatusCode(v int32) *RemoveUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveUsersResponse) SetBody(v *RemoveUsersResponseBody) *RemoveUsersResponse {
	s.Body = v
	return s
}

type StartCloudRecordRequest struct {
	// appId
	//
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
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
	CropMode *int32                           `json:"CropMode,omitempty" xml:"CropMode,omitempty"`
	Images   []*StartCloudRecordRequestImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// panes
	Panes []*StartCloudRecordRequestPanes `json:"Panes,omitempty" xml:"Panes,omitempty" type:"Repeated"`
	// storageConfig
	//
	// This parameter is required.
	StorageConfig *StartCloudRecordRequestStorageConfig `json:"StorageConfig,omitempty" xml:"StorageConfig,omitempty" type:"Struct"`
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
	return tea.Prettify(s)
}

func (s StartCloudRecordRequest) GoString() string {
	return s.String()
}

func (s *StartCloudRecordRequest) SetAppId(v string) *StartCloudRecordRequest {
	s.AppId = &v
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

func (s *StartCloudRecordRequest) SetPanes(v []*StartCloudRecordRequestPanes) *StartCloudRecordRequest {
	s.Panes = v
	return s
}

func (s *StartCloudRecordRequest) SetStorageConfig(v *StartCloudRecordRequestStorageConfig) *StartCloudRecordRequest {
	s.StorageConfig = v
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
	return tea.Prettify(s)
}

func (s StartCloudRecordRequestClockWidgets) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s StartCloudRecordRequestClockWidgetsBoxColor) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s StartCloudRecordRequestClockWidgetsFontColor) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s StartCloudRecordRequestImages) GoString() string {
	return s.String()
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

type StartCloudRecordRequestPanes struct {
	Images []*StartCloudRecordRequestPanesImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
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
	PaneId *int32 `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
	// source
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// sourceType
	//
	// example:
	//
	// video
	SourceType *string                              `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Texts      []*StartCloudRecordRequestPanesTexts `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
}

func (s StartCloudRecordRequestPanes) String() string {
	return tea.Prettify(s)
}

func (s StartCloudRecordRequestPanes) GoString() string {
	return s.String()
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

type StartCloudRecordRequestPanesImages struct {
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
	return tea.Prettify(s)
}

func (s StartCloudRecordRequestPanesImages) GoString() string {
	return s.String()
}

func (s *StartCloudRecordRequestPanesImages) SetAlpha(v float64) *StartCloudRecordRequestPanesImages {
	s.Alpha = &v
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
	return tea.Prettify(s)
}

func (s StartCloudRecordRequestPanesTexts) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s StartCloudRecordRequestPanesTextsBoxColor) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s StartCloudRecordRequestPanesTextsFontColor) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s StartCloudRecordRequestStorageConfig) GoString() string {
	return s.String()
}

func (s *StartCloudRecordRequestStorageConfig) SetAccessKey(v string) *StartCloudRecordRequestStorageConfig {
	s.AccessKey = &v
	return s
}

func (s *StartCloudRecordRequestStorageConfig) SetBucket(v string) *StartCloudRecordRequestStorageConfig {
	s.Bucket = &v
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
	return tea.Prettify(s)
}

func (s StartCloudRecordRequestTexts) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s StartCloudRecordRequestTextsBoxColor) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s StartCloudRecordRequestTextsFontColor) GoString() string {
	return s.String()
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

type StartCloudRecordResponseBody struct {
	// requestId
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// taskId
	//
	// example:
	//
	// 123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StartCloudRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartCloudRecordResponseBody) GoString() string {
	return s.String()
}

func (s *StartCloudRecordResponseBody) SetRequestId(v string) *StartCloudRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartCloudRecordResponseBody) SetTaskId(v string) *StartCloudRecordResponseBody {
	s.TaskId = &v
	return s
}

type StartCloudRecordResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartCloudRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartCloudRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s StartCloudRecordResponse) GoString() string {
	return s.String()
}

func (s *StartCloudRecordResponse) SetHeaders(v map[string]*string) *StartCloudRecordResponse {
	s.Headers = v
	return s
}

func (s *StartCloudRecordResponse) SetStatusCode(v int32) *StartCloudRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *StartCloudRecordResponse) SetBody(v *StartCloudRecordResponseBody) *StartCloudRecordResponse {
	s.Body = v
	return s
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
	return tea.Prettify(s)
}

func (s StartMPUTaskRequestClockWidgets) GoString() string {
	return s.String()
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

type StartMPUTaskRequestEnhancedParam struct {
	// example:
	//
	// false
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
	return tea.Prettify(s)
}

func (s StartMPUTaskRequestUserPanesTexts) GoString() string {
	return s.String()
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
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartMPUTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *StartMPUTaskResponse) SetStatusCode(v int32) *StartMPUTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StartMPUTaskResponse) SetBody(v *StartMPUTaskResponseBody) *StartMPUTaskResponse {
	s.Body = v
	return s
}

type StartRecordTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// yourChannelId
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// 1
	CropMode *int64 `json:"CropMode,omitempty" xml:"CropMode,omitempty"`
	// example:
	//
	// 1111
	LayoutIds []*int64 `json:"LayoutIds,omitempty" xml:"LayoutIds,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MediaEncode *int32 `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	// example:
	//
	// 1
	MixMode *int32 `json:"MixMode,omitempty" xml:"MixMode,omitempty"`
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// 1
	SubSpecAudioUsers []*string `json:"SubSpecAudioUsers,omitempty" xml:"SubSpecAudioUsers,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	SubSpecCameraUsers []*string `json:"SubSpecCameraUsers,omitempty" xml:"SubSpecCameraUsers,omitempty" type:"Repeated"`
	// example:
	//
	// 1
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
	// 4IN_1080P
	TaskProfile *string `json:"TaskProfile,omitempty" xml:"TaskProfile,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 76dasgb****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// 1
	UnsubSpecAudioUsers []*string `json:"UnsubSpecAudioUsers,omitempty" xml:"UnsubSpecAudioUsers,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	UnsubSpecCameraUsers []*string `json:"UnsubSpecCameraUsers,omitempty" xml:"UnsubSpecCameraUsers,omitempty" type:"Repeated"`
	// example:
	//
	// 1
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
	Images []*StartRecordTaskRequestUserPanesImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	PaneId *int32 `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
	// example:
	//
	// camera
	SourceType *string                                 `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Texts      []*StartRecordTaskRequestUserPanesTexts `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
	// example:
	//
	// TestId
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// example:
	//
	// 1
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
	// 0.7576
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	// example:
	//
	// 0
	ZOrder *int32 `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
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
	// example:
	//
	// 760bad53276431c499e30dc36f6b****
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartRecordTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *StartRecordTaskResponse) SetStatusCode(v int32) *StartRecordTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StartRecordTaskResponse) SetBody(v *StartRecordTaskResponseBody) *StartRecordTaskResponse {
	s.Body = v
	return s
}

type StartStreamingOutRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId    *string                                 `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ClockWidgets []*StartStreamingOutRequestClockWidgets `json:"ClockWidgets,omitempty" xml:"ClockWidgets,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	CropMode *int32                            `json:"CropMode,omitempty" xml:"CropMode,omitempty"`
	Images   []*StartStreamingOutRequestImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	Panes    []*StartStreamingOutRequestPanes  `json:"Panes,omitempty" xml:"Panes,omitempty" type:"Repeated"`
	// example:
	//
	// 123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 567
	TemplateId *string                          `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Texts      []*StartStreamingOutRequestTexts `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// rtmp://example.com/live/stream
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s StartStreamingOutRequest) String() string {
	return tea.Prettify(s)
}

func (s StartStreamingOutRequest) GoString() string {
	return s.String()
}

func (s *StartStreamingOutRequest) SetAppId(v string) *StartStreamingOutRequest {
	s.AppId = &v
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

func (s *StartStreamingOutRequest) SetPanes(v []*StartStreamingOutRequestPanes) *StartStreamingOutRequest {
	s.Panes = v
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

type StartStreamingOutRequestClockWidgets struct {
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
	BoxBorderw *int32                                        `json:"BoxBorderw,omitempty" xml:"BoxBorderw,omitempty"`
	BoxColor   *StartStreamingOutRequestClockWidgetsBoxColor `json:"BoxColor,omitempty" xml:"BoxColor,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Font      *int32                                         `json:"Font,omitempty" xml:"Font,omitempty"`
	FontColor *StartStreamingOutRequestClockWidgetsFontColor `json:"FontColor,omitempty" xml:"FontColor,omitempty" type:"Struct"`
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

func (s StartStreamingOutRequestClockWidgets) String() string {
	return tea.Prettify(s)
}

func (s StartStreamingOutRequestClockWidgets) GoString() string {
	return s.String()
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

type StartStreamingOutRequestClockWidgetsBoxColor struct {
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

func (s StartStreamingOutRequestClockWidgetsBoxColor) String() string {
	return tea.Prettify(s)
}

func (s StartStreamingOutRequestClockWidgetsBoxColor) GoString() string {
	return s.String()
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

type StartStreamingOutRequestClockWidgetsFontColor struct {
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

func (s StartStreamingOutRequestClockWidgetsFontColor) String() string {
	return tea.Prettify(s)
}

func (s StartStreamingOutRequestClockWidgetsFontColor) GoString() string {
	return s.String()
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

type StartStreamingOutRequestImages struct {
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
	// 0.2
	Y *float64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s StartStreamingOutRequestImages) String() string {
	return tea.Prettify(s)
}

func (s StartStreamingOutRequestImages) GoString() string {
	return s.String()
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

type StartStreamingOutRequestPanes struct {
	Images []*StartStreamingOutRequestPanesImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	PaneCropMode *int32 `json:"PaneCropMode,omitempty" xml:"PaneCropMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	PaneId *string `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
	// example:
	//
	// 1811****
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// Video
	SourceType *string                               `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Texts      []*StartStreamingOutRequestPanesTexts `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
}

func (s StartStreamingOutRequestPanes) String() string {
	return tea.Prettify(s)
}

func (s StartStreamingOutRequestPanes) GoString() string {
	return s.String()
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

type StartStreamingOutRequestPanesImages struct {
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

func (s StartStreamingOutRequestPanesImages) String() string {
	return tea.Prettify(s)
}

func (s StartStreamingOutRequestPanesImages) GoString() string {
	return s.String()
}

func (s *StartStreamingOutRequestPanesImages) SetAlpha(v float64) *StartStreamingOutRequestPanesImages {
	s.Alpha = &v
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

type StartStreamingOutRequestPanesTexts struct {
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
	BoxBorderw *int32                                      `json:"BoxBorderw,omitempty" xml:"BoxBorderw,omitempty"`
	BoxColor   *StartStreamingOutRequestPanesTextsBoxColor `json:"BoxColor,omitempty" xml:"BoxColor,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Font      *int32                                       `json:"Font,omitempty" xml:"Font,omitempty"`
	FontColor *StartStreamingOutRequestPanesTextsFontColor `json:"FontColor,omitempty" xml:"FontColor,omitempty" type:"Struct"`
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

func (s StartStreamingOutRequestPanesTexts) String() string {
	return tea.Prettify(s)
}

func (s StartStreamingOutRequestPanesTexts) GoString() string {
	return s.String()
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

type StartStreamingOutRequestPanesTextsBoxColor struct {
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

func (s StartStreamingOutRequestPanesTextsBoxColor) String() string {
	return tea.Prettify(s)
}

func (s StartStreamingOutRequestPanesTextsBoxColor) GoString() string {
	return s.String()
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

type StartStreamingOutRequestPanesTextsFontColor struct {
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

func (s StartStreamingOutRequestPanesTextsFontColor) String() string {
	return tea.Prettify(s)
}

func (s StartStreamingOutRequestPanesTextsFontColor) GoString() string {
	return s.String()
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

type StartStreamingOutRequestTexts struct {
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
	BoxBorderw *int32                                 `json:"BoxBorderw,omitempty" xml:"BoxBorderw,omitempty"`
	BoxColor   *StartStreamingOutRequestTextsBoxColor `json:"BoxColor,omitempty" xml:"BoxColor,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Font      *int32                                  `json:"Font,omitempty" xml:"Font,omitempty"`
	FontColor *StartStreamingOutRequestTextsFontColor `json:"FontColor,omitempty" xml:"FontColor,omitempty" type:"Struct"`
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

func (s StartStreamingOutRequestTexts) String() string {
	return tea.Prettify(s)
}

func (s StartStreamingOutRequestTexts) GoString() string {
	return s.String()
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

type StartStreamingOutRequestTextsBoxColor struct {
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

func (s StartStreamingOutRequestTextsBoxColor) String() string {
	return tea.Prettify(s)
}

func (s StartStreamingOutRequestTextsBoxColor) GoString() string {
	return s.String()
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

type StartStreamingOutRequestTextsFontColor struct {
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

func (s StartStreamingOutRequestTextsFontColor) String() string {
	return tea.Prettify(s)
}

func (s StartStreamingOutRequestTextsFontColor) GoString() string {
	return s.String()
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

type StartStreamingOutResponseBody struct {
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StartStreamingOutResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartStreamingOutResponseBody) GoString() string {
	return s.String()
}

func (s *StartStreamingOutResponseBody) SetRequestId(v string) *StartStreamingOutResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartStreamingOutResponseBody) SetTaskId(v string) *StartStreamingOutResponseBody {
	s.TaskId = &v
	return s
}

type StartStreamingOutResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartStreamingOutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartStreamingOutResponse) String() string {
	return tea.Prettify(s)
}

func (s StartStreamingOutResponse) GoString() string {
	return s.String()
}

func (s *StartStreamingOutResponse) SetHeaders(v map[string]*string) *StartStreamingOutResponse {
	s.Headers = v
	return s
}

func (s *StartStreamingOutResponse) SetStatusCode(v int32) *StartStreamingOutResponse {
	s.StatusCode = &v
	return s
}

func (s *StartStreamingOutResponse) SetBody(v *StartStreamingOutResponseBody) *StartStreamingOutResponse {
	s.Body = v
	return s
}

type StopChannelRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
}

func (s StopChannelRequest) String() string {
	return tea.Prettify(s)
}

func (s StopChannelRequest) GoString() string {
	return s.String()
}

func (s *StopChannelRequest) SetAppId(v string) *StopChannelRequest {
	s.AppId = &v
	return s
}

func (s *StopChannelRequest) SetChannelId(v string) *StopChannelRequest {
	s.ChannelId = &v
	return s
}

type StopChannelResponseBody struct {
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopChannelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopChannelResponseBody) GoString() string {
	return s.String()
}

func (s *StopChannelResponseBody) SetRequestId(v string) *StopChannelResponseBody {
	s.RequestId = &v
	return s
}

type StopChannelResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopChannelResponse) String() string {
	return tea.Prettify(s)
}

func (s StopChannelResponse) GoString() string {
	return s.String()
}

func (s *StopChannelResponse) SetHeaders(v map[string]*string) *StopChannelResponse {
	s.Headers = v
	return s
}

func (s *StopChannelResponse) SetStatusCode(v int32) *StopChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *StopChannelResponse) SetBody(v *StopChannelResponseBody) *StopChannelResponse {
	s.Body = v
	return s
}

type StopCloudRecordRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopCloudRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s StopCloudRecordRequest) GoString() string {
	return s.String()
}

func (s *StopCloudRecordRequest) SetAppId(v string) *StopCloudRecordRequest {
	s.AppId = &v
	return s
}

func (s *StopCloudRecordRequest) SetChannelId(v string) *StopCloudRecordRequest {
	s.ChannelId = &v
	return s
}

func (s *StopCloudRecordRequest) SetTaskId(v string) *StopCloudRecordRequest {
	s.TaskId = &v
	return s
}

type StopCloudRecordResponseBody struct {
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopCloudRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopCloudRecordResponseBody) GoString() string {
	return s.String()
}

func (s *StopCloudRecordResponseBody) SetRequestId(v string) *StopCloudRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopCloudRecordResponseBody) SetTaskId(v string) *StopCloudRecordResponseBody {
	s.TaskId = &v
	return s
}

type StopCloudRecordResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopCloudRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopCloudRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s StopCloudRecordResponse) GoString() string {
	return s.String()
}

func (s *StopCloudRecordResponse) SetHeaders(v map[string]*string) *StopCloudRecordResponse {
	s.Headers = v
	return s
}

func (s *StopCloudRecordResponse) SetStatusCode(v int32) *StopCloudRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *StopCloudRecordResponse) SetBody(v *StopCloudRecordResponseBody) *StopCloudRecordResponse {
	s.Body = v
	return s
}

type StopMPUTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// yourTaskId
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopMPUTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *StopMPUTaskResponse) SetStatusCode(v int32) *StopMPUTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StopMPUTaskResponse) SetBody(v *StopMPUTaskResponseBody) *StopMPUTaskResponse {
	s.Body = v
	return s
}

type StopRecordTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// yourTaskId
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopRecordTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *StopRecordTaskResponse) SetStatusCode(v int32) *StopRecordTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StopRecordTaskResponse) SetBody(v *StopRecordTaskResponseBody) *StopRecordTaskResponse {
	s.Body = v
	return s
}

type StopStreamingOutRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopStreamingOutRequest) String() string {
	return tea.Prettify(s)
}

func (s StopStreamingOutRequest) GoString() string {
	return s.String()
}

func (s *StopStreamingOutRequest) SetAppId(v string) *StopStreamingOutRequest {
	s.AppId = &v
	return s
}

func (s *StopStreamingOutRequest) SetChannelId(v string) *StopStreamingOutRequest {
	s.ChannelId = &v
	return s
}

func (s *StopStreamingOutRequest) SetTaskId(v string) *StopStreamingOutRequest {
	s.TaskId = &v
	return s
}

type StopStreamingOutResponseBody struct {
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopStreamingOutResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopStreamingOutResponseBody) GoString() string {
	return s.String()
}

func (s *StopStreamingOutResponseBody) SetRequestId(v string) *StopStreamingOutResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopStreamingOutResponseBody) SetTaskId(v string) *StopStreamingOutResponseBody {
	s.TaskId = &v
	return s
}

type StopStreamingOutResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopStreamingOutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopStreamingOutResponse) String() string {
	return tea.Prettify(s)
}

func (s StopStreamingOutResponse) GoString() string {
	return s.String()
}

func (s *StopStreamingOutResponse) SetHeaders(v map[string]*string) *StopStreamingOutResponse {
	s.Headers = v
	return s
}

func (s *StopStreamingOutResponse) SetStatusCode(v int32) *StopStreamingOutResponse {
	s.StatusCode = &v
	return s
}

func (s *StopStreamingOutResponse) SetBody(v *StopStreamingOutResponseBody) *StopStreamingOutResponse {
	s.Body = v
	return s
}

type UpdateAutoLiveStreamRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// http://example.com/callback
	CallBack          *string   `json:"CallBack,omitempty" xml:"CallBack,omitempty"`
	ChannelIdPrefixes []*string `json:"ChannelIdPrefixes,omitempty" xml:"ChannelIdPrefixes,omitempty" type:"Repeated"`
	ChannelIds        []*string `json:"ChannelIds,omitempty" xml:"ChannelIds,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MediaEncode *int32 `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	OwnerId     *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rtmp://${domain}/${app}/${stream}
	PlayDomain *string `json:"PlayDomain,omitempty" xml:"PlayDomain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12
	RuleId *int32 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// testRule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
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
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

type UpdateAutoLiveStreamRuleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAutoLiveStreamRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UpdateAutoLiveStreamRuleResponse) SetStatusCode(v int32) *UpdateAutoLiveStreamRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAutoLiveStreamRuleResponse) SetBody(v *UpdateAutoLiveStreamRuleResponseBody) *UpdateAutoLiveStreamRuleResponse {
	s.Body = v
	return s
}

type UpdateCloudRecordRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId    *string                                 `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ClockWidgets []*UpdateCloudRecordRequestClockWidgets `json:"ClockWidgets,omitempty" xml:"ClockWidgets,omitempty" type:"Repeated"`
	Images       []*UpdateCloudRecordRequestImages       `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	Panes        []*UpdateCloudRecordRequestPanes        `json:"Panes,omitempty" xml:"Panes,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 567
	TemplateId *string                          `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Texts      []*UpdateCloudRecordRequestTexts `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
}

func (s UpdateCloudRecordRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCloudRecordRequest) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordRequest) SetAppId(v string) *UpdateCloudRecordRequest {
	s.AppId = &v
	return s
}

func (s *UpdateCloudRecordRequest) SetChannelId(v string) *UpdateCloudRecordRequest {
	s.ChannelId = &v
	return s
}

func (s *UpdateCloudRecordRequest) SetClockWidgets(v []*UpdateCloudRecordRequestClockWidgets) *UpdateCloudRecordRequest {
	s.ClockWidgets = v
	return s
}

func (s *UpdateCloudRecordRequest) SetImages(v []*UpdateCloudRecordRequestImages) *UpdateCloudRecordRequest {
	s.Images = v
	return s
}

func (s *UpdateCloudRecordRequest) SetPanes(v []*UpdateCloudRecordRequestPanes) *UpdateCloudRecordRequest {
	s.Panes = v
	return s
}

func (s *UpdateCloudRecordRequest) SetTaskId(v string) *UpdateCloudRecordRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateCloudRecordRequest) SetTemplateId(v string) *UpdateCloudRecordRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateCloudRecordRequest) SetTexts(v []*UpdateCloudRecordRequestTexts) *UpdateCloudRecordRequest {
	s.Texts = v
	return s
}

type UpdateCloudRecordRequestClockWidgets struct {
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
	BoxBorderw *int32                                        `json:"BoxBorderw,omitempty" xml:"BoxBorderw,omitempty"`
	BoxColor   *UpdateCloudRecordRequestClockWidgetsBoxColor `json:"BoxColor,omitempty" xml:"BoxColor,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Font      *int32                                         `json:"Font,omitempty" xml:"Font,omitempty"`
	FontColor *UpdateCloudRecordRequestClockWidgetsFontColor `json:"FontColor,omitempty" xml:"FontColor,omitempty" type:"Struct"`
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

func (s UpdateCloudRecordRequestClockWidgets) String() string {
	return tea.Prettify(s)
}

func (s UpdateCloudRecordRequestClockWidgets) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordRequestClockWidgets) SetAlpha(v float64) *UpdateCloudRecordRequestClockWidgets {
	s.Alpha = &v
	return s
}

func (s *UpdateCloudRecordRequestClockWidgets) SetBoxAlpha(v float64) *UpdateCloudRecordRequestClockWidgets {
	s.BoxAlpha = &v
	return s
}

func (s *UpdateCloudRecordRequestClockWidgets) SetBoxBorderw(v int32) *UpdateCloudRecordRequestClockWidgets {
	s.BoxBorderw = &v
	return s
}

func (s *UpdateCloudRecordRequestClockWidgets) SetBoxColor(v *UpdateCloudRecordRequestClockWidgetsBoxColor) *UpdateCloudRecordRequestClockWidgets {
	s.BoxColor = v
	return s
}

func (s *UpdateCloudRecordRequestClockWidgets) SetFont(v int32) *UpdateCloudRecordRequestClockWidgets {
	s.Font = &v
	return s
}

func (s *UpdateCloudRecordRequestClockWidgets) SetFontColor(v *UpdateCloudRecordRequestClockWidgetsFontColor) *UpdateCloudRecordRequestClockWidgets {
	s.FontColor = v
	return s
}

func (s *UpdateCloudRecordRequestClockWidgets) SetFontSize(v int32) *UpdateCloudRecordRequestClockWidgets {
	s.FontSize = &v
	return s
}

func (s *UpdateCloudRecordRequestClockWidgets) SetHasBox(v bool) *UpdateCloudRecordRequestClockWidgets {
	s.HasBox = &v
	return s
}

func (s *UpdateCloudRecordRequestClockWidgets) SetLayer(v int32) *UpdateCloudRecordRequestClockWidgets {
	s.Layer = &v
	return s
}

func (s *UpdateCloudRecordRequestClockWidgets) SetX(v float64) *UpdateCloudRecordRequestClockWidgets {
	s.X = &v
	return s
}

func (s *UpdateCloudRecordRequestClockWidgets) SetY(v float64) *UpdateCloudRecordRequestClockWidgets {
	s.Y = &v
	return s
}

func (s *UpdateCloudRecordRequestClockWidgets) SetZone(v int32) *UpdateCloudRecordRequestClockWidgets {
	s.Zone = &v
	return s
}

type UpdateCloudRecordRequestClockWidgetsBoxColor struct {
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

func (s UpdateCloudRecordRequestClockWidgetsBoxColor) String() string {
	return tea.Prettify(s)
}

func (s UpdateCloudRecordRequestClockWidgetsBoxColor) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordRequestClockWidgetsBoxColor) SetB(v int32) *UpdateCloudRecordRequestClockWidgetsBoxColor {
	s.B = &v
	return s
}

func (s *UpdateCloudRecordRequestClockWidgetsBoxColor) SetG(v int32) *UpdateCloudRecordRequestClockWidgetsBoxColor {
	s.G = &v
	return s
}

func (s *UpdateCloudRecordRequestClockWidgetsBoxColor) SetR(v int32) *UpdateCloudRecordRequestClockWidgetsBoxColor {
	s.R = &v
	return s
}

type UpdateCloudRecordRequestClockWidgetsFontColor struct {
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

func (s UpdateCloudRecordRequestClockWidgetsFontColor) String() string {
	return tea.Prettify(s)
}

func (s UpdateCloudRecordRequestClockWidgetsFontColor) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordRequestClockWidgetsFontColor) SetB(v int32) *UpdateCloudRecordRequestClockWidgetsFontColor {
	s.B = &v
	return s
}

func (s *UpdateCloudRecordRequestClockWidgetsFontColor) SetG(v int32) *UpdateCloudRecordRequestClockWidgetsFontColor {
	s.G = &v
	return s
}

func (s *UpdateCloudRecordRequestClockWidgetsFontColor) SetR(v int32) *UpdateCloudRecordRequestClockWidgetsFontColor {
	s.R = &v
	return s
}

type UpdateCloudRecordRequestImages struct {
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

func (s UpdateCloudRecordRequestImages) String() string {
	return tea.Prettify(s)
}

func (s UpdateCloudRecordRequestImages) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordRequestImages) SetAlpha(v float64) *UpdateCloudRecordRequestImages {
	s.Alpha = &v
	return s
}

func (s *UpdateCloudRecordRequestImages) SetHeight(v float64) *UpdateCloudRecordRequestImages {
	s.Height = &v
	return s
}

func (s *UpdateCloudRecordRequestImages) SetImageCropMode(v int32) *UpdateCloudRecordRequestImages {
	s.ImageCropMode = &v
	return s
}

func (s *UpdateCloudRecordRequestImages) SetLayer(v int32) *UpdateCloudRecordRequestImages {
	s.Layer = &v
	return s
}

func (s *UpdateCloudRecordRequestImages) SetUrl(v string) *UpdateCloudRecordRequestImages {
	s.Url = &v
	return s
}

func (s *UpdateCloudRecordRequestImages) SetWidth(v float64) *UpdateCloudRecordRequestImages {
	s.Width = &v
	return s
}

func (s *UpdateCloudRecordRequestImages) SetX(v float64) *UpdateCloudRecordRequestImages {
	s.X = &v
	return s
}

func (s *UpdateCloudRecordRequestImages) SetY(v float64) *UpdateCloudRecordRequestImages {
	s.Y = &v
	return s
}

type UpdateCloudRecordRequestPanes struct {
	Images []*UpdateCloudRecordRequestPanesImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	PaneCropMode *int32 `json:"PaneCropMode,omitempty" xml:"PaneCropMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	PaneId *int32 `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
	// example:
	//
	// 22
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// video
	SourceType *string                               `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Texts      []*UpdateCloudRecordRequestPanesTexts `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
}

func (s UpdateCloudRecordRequestPanes) String() string {
	return tea.Prettify(s)
}

func (s UpdateCloudRecordRequestPanes) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordRequestPanes) SetImages(v []*UpdateCloudRecordRequestPanesImages) *UpdateCloudRecordRequestPanes {
	s.Images = v
	return s
}

func (s *UpdateCloudRecordRequestPanes) SetPaneCropMode(v int32) *UpdateCloudRecordRequestPanes {
	s.PaneCropMode = &v
	return s
}

func (s *UpdateCloudRecordRequestPanes) SetPaneId(v int32) *UpdateCloudRecordRequestPanes {
	s.PaneId = &v
	return s
}

func (s *UpdateCloudRecordRequestPanes) SetSource(v string) *UpdateCloudRecordRequestPanes {
	s.Source = &v
	return s
}

func (s *UpdateCloudRecordRequestPanes) SetSourceType(v string) *UpdateCloudRecordRequestPanes {
	s.SourceType = &v
	return s
}

func (s *UpdateCloudRecordRequestPanes) SetTexts(v []*UpdateCloudRecordRequestPanesTexts) *UpdateCloudRecordRequestPanes {
	s.Texts = v
	return s
}

type UpdateCloudRecordRequestPanesImages struct {
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
	// 0.2
	Y *float64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s UpdateCloudRecordRequestPanesImages) String() string {
	return tea.Prettify(s)
}

func (s UpdateCloudRecordRequestPanesImages) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordRequestPanesImages) SetAlpha(v float64) *UpdateCloudRecordRequestPanesImages {
	s.Alpha = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesImages) SetHeight(v float64) *UpdateCloudRecordRequestPanesImages {
	s.Height = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesImages) SetLayer(v int32) *UpdateCloudRecordRequestPanesImages {
	s.Layer = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesImages) SetPaneImageCropMode(v int32) *UpdateCloudRecordRequestPanesImages {
	s.PaneImageCropMode = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesImages) SetUrl(v string) *UpdateCloudRecordRequestPanesImages {
	s.Url = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesImages) SetWidth(v float64) *UpdateCloudRecordRequestPanesImages {
	s.Width = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesImages) SetX(v float64) *UpdateCloudRecordRequestPanesImages {
	s.X = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesImages) SetY(v float64) *UpdateCloudRecordRequestPanesImages {
	s.Y = &v
	return s
}

type UpdateCloudRecordRequestPanesTexts struct {
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
	BoxBorderw *int32                                      `json:"BoxBorderw,omitempty" xml:"BoxBorderw,omitempty"`
	BoxColor   *UpdateCloudRecordRequestPanesTextsBoxColor `json:"BoxColor,omitempty" xml:"BoxColor,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Font      *int32                                       `json:"Font,omitempty" xml:"Font,omitempty"`
	FontColor *UpdateCloudRecordRequestPanesTextsFontColor `json:"FontColor,omitempty" xml:"FontColor,omitempty" type:"Struct"`
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

func (s UpdateCloudRecordRequestPanesTexts) String() string {
	return tea.Prettify(s)
}

func (s UpdateCloudRecordRequestPanesTexts) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordRequestPanesTexts) SetAlpha(v float64) *UpdateCloudRecordRequestPanesTexts {
	s.Alpha = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesTexts) SetBoxAlpha(v float64) *UpdateCloudRecordRequestPanesTexts {
	s.BoxAlpha = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesTexts) SetBoxBorderw(v int32) *UpdateCloudRecordRequestPanesTexts {
	s.BoxBorderw = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesTexts) SetBoxColor(v *UpdateCloudRecordRequestPanesTextsBoxColor) *UpdateCloudRecordRequestPanesTexts {
	s.BoxColor = v
	return s
}

func (s *UpdateCloudRecordRequestPanesTexts) SetFont(v int32) *UpdateCloudRecordRequestPanesTexts {
	s.Font = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesTexts) SetFontColor(v *UpdateCloudRecordRequestPanesTextsFontColor) *UpdateCloudRecordRequestPanesTexts {
	s.FontColor = v
	return s
}

func (s *UpdateCloudRecordRequestPanesTexts) SetFontSize(v int32) *UpdateCloudRecordRequestPanesTexts {
	s.FontSize = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesTexts) SetHasBox(v bool) *UpdateCloudRecordRequestPanesTexts {
	s.HasBox = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesTexts) SetLayer(v int32) *UpdateCloudRecordRequestPanesTexts {
	s.Layer = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesTexts) SetTexture(v string) *UpdateCloudRecordRequestPanesTexts {
	s.Texture = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesTexts) SetX(v float64) *UpdateCloudRecordRequestPanesTexts {
	s.X = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesTexts) SetY(v float64) *UpdateCloudRecordRequestPanesTexts {
	s.Y = &v
	return s
}

type UpdateCloudRecordRequestPanesTextsBoxColor struct {
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

func (s UpdateCloudRecordRequestPanesTextsBoxColor) String() string {
	return tea.Prettify(s)
}

func (s UpdateCloudRecordRequestPanesTextsBoxColor) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordRequestPanesTextsBoxColor) SetB(v int32) *UpdateCloudRecordRequestPanesTextsBoxColor {
	s.B = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesTextsBoxColor) SetG(v int32) *UpdateCloudRecordRequestPanesTextsBoxColor {
	s.G = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesTextsBoxColor) SetR(v int32) *UpdateCloudRecordRequestPanesTextsBoxColor {
	s.R = &v
	return s
}

type UpdateCloudRecordRequestPanesTextsFontColor struct {
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

func (s UpdateCloudRecordRequestPanesTextsFontColor) String() string {
	return tea.Prettify(s)
}

func (s UpdateCloudRecordRequestPanesTextsFontColor) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordRequestPanesTextsFontColor) SetB(v int32) *UpdateCloudRecordRequestPanesTextsFontColor {
	s.B = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesTextsFontColor) SetG(v int32) *UpdateCloudRecordRequestPanesTextsFontColor {
	s.G = &v
	return s
}

func (s *UpdateCloudRecordRequestPanesTextsFontColor) SetR(v int32) *UpdateCloudRecordRequestPanesTextsFontColor {
	s.R = &v
	return s
}

type UpdateCloudRecordRequestTexts struct {
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
	BoxBorderw *int32                                 `json:"BoxBorderw,omitempty" xml:"BoxBorderw,omitempty"`
	BoxColor   *UpdateCloudRecordRequestTextsBoxColor `json:"BoxColor,omitempty" xml:"BoxColor,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Font      *int32                                  `json:"Font,omitempty" xml:"Font,omitempty"`
	FontColor *UpdateCloudRecordRequestTextsFontColor `json:"FontColor,omitempty" xml:"FontColor,omitempty" type:"Struct"`
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

func (s UpdateCloudRecordRequestTexts) String() string {
	return tea.Prettify(s)
}

func (s UpdateCloudRecordRequestTexts) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordRequestTexts) SetAlpha(v float64) *UpdateCloudRecordRequestTexts {
	s.Alpha = &v
	return s
}

func (s *UpdateCloudRecordRequestTexts) SetBoxAlpha(v float64) *UpdateCloudRecordRequestTexts {
	s.BoxAlpha = &v
	return s
}

func (s *UpdateCloudRecordRequestTexts) SetBoxBorderw(v int32) *UpdateCloudRecordRequestTexts {
	s.BoxBorderw = &v
	return s
}

func (s *UpdateCloudRecordRequestTexts) SetBoxColor(v *UpdateCloudRecordRequestTextsBoxColor) *UpdateCloudRecordRequestTexts {
	s.BoxColor = v
	return s
}

func (s *UpdateCloudRecordRequestTexts) SetFont(v int32) *UpdateCloudRecordRequestTexts {
	s.Font = &v
	return s
}

func (s *UpdateCloudRecordRequestTexts) SetFontColor(v *UpdateCloudRecordRequestTextsFontColor) *UpdateCloudRecordRequestTexts {
	s.FontColor = v
	return s
}

func (s *UpdateCloudRecordRequestTexts) SetFontSize(v int32) *UpdateCloudRecordRequestTexts {
	s.FontSize = &v
	return s
}

func (s *UpdateCloudRecordRequestTexts) SetHasBox(v bool) *UpdateCloudRecordRequestTexts {
	s.HasBox = &v
	return s
}

func (s *UpdateCloudRecordRequestTexts) SetLayer(v int32) *UpdateCloudRecordRequestTexts {
	s.Layer = &v
	return s
}

func (s *UpdateCloudRecordRequestTexts) SetTexture(v string) *UpdateCloudRecordRequestTexts {
	s.Texture = &v
	return s
}

func (s *UpdateCloudRecordRequestTexts) SetX(v float64) *UpdateCloudRecordRequestTexts {
	s.X = &v
	return s
}

func (s *UpdateCloudRecordRequestTexts) SetY(v float64) *UpdateCloudRecordRequestTexts {
	s.Y = &v
	return s
}

type UpdateCloudRecordRequestTextsBoxColor struct {
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

func (s UpdateCloudRecordRequestTextsBoxColor) String() string {
	return tea.Prettify(s)
}

func (s UpdateCloudRecordRequestTextsBoxColor) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordRequestTextsBoxColor) SetB(v int32) *UpdateCloudRecordRequestTextsBoxColor {
	s.B = &v
	return s
}

func (s *UpdateCloudRecordRequestTextsBoxColor) SetG(v int32) *UpdateCloudRecordRequestTextsBoxColor {
	s.G = &v
	return s
}

func (s *UpdateCloudRecordRequestTextsBoxColor) SetR(v int32) *UpdateCloudRecordRequestTextsBoxColor {
	s.R = &v
	return s
}

type UpdateCloudRecordRequestTextsFontColor struct {
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

func (s UpdateCloudRecordRequestTextsFontColor) String() string {
	return tea.Prettify(s)
}

func (s UpdateCloudRecordRequestTextsFontColor) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordRequestTextsFontColor) SetB(v int32) *UpdateCloudRecordRequestTextsFontColor {
	s.B = &v
	return s
}

func (s *UpdateCloudRecordRequestTextsFontColor) SetG(v int32) *UpdateCloudRecordRequestTextsFontColor {
	s.G = &v
	return s
}

func (s *UpdateCloudRecordRequestTextsFontColor) SetR(v int32) *UpdateCloudRecordRequestTextsFontColor {
	s.R = &v
	return s
}

type UpdateCloudRecordResponseBody struct {
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateCloudRecordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCloudRecordResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordResponseBody) SetRequestId(v string) *UpdateCloudRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCloudRecordResponseBody) SetTaskId(v string) *UpdateCloudRecordResponseBody {
	s.TaskId = &v
	return s
}

type UpdateCloudRecordResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCloudRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCloudRecordResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCloudRecordResponse) GoString() string {
	return s.String()
}

func (s *UpdateCloudRecordResponse) SetHeaders(v map[string]*string) *UpdateCloudRecordResponse {
	s.Headers = v
	return s
}

func (s *UpdateCloudRecordResponse) SetStatusCode(v int32) *UpdateCloudRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCloudRecordResponse) SetBody(v *UpdateCloudRecordResponseBody) *UpdateCloudRecordResponse {
	s.Body = v
	return s
}

type UpdateMPUTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 0
	BackgroundColor *int32                              `json:"BackgroundColor,omitempty" xml:"BackgroundColor,omitempty"`
	Backgrounds     []*UpdateMPUTaskRequestBackgrounds  `json:"Backgrounds,omitempty" xml:"Backgrounds,omitempty" type:"Repeated"`
	ClockWidgets    []*UpdateMPUTaskRequestClockWidgets `json:"ClockWidgets,omitempty" xml:"ClockWidgets,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	CropMode  *int32   `json:"CropMode,omitempty" xml:"CropMode,omitempty"`
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
	// camera
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// 0
	StreamType              *int32    `json:"StreamType,omitempty" xml:"StreamType,omitempty"`
	SubSpecAudioUsers       []*string `json:"SubSpecAudioUsers,omitempty" xml:"SubSpecAudioUsers,omitempty" type:"Repeated"`
	SubSpecCameraUsers      []*string `json:"SubSpecCameraUsers,omitempty" xml:"SubSpecCameraUsers,omitempty" type:"Repeated"`
	SubSpecShareScreenUsers []*string `json:"SubSpecShareScreenUsers,omitempty" xml:"SubSpecShareScreenUsers,omitempty" type:"Repeated"`
	SubSpecUsers            []*string `json:"SubSpecUsers,omitempty" xml:"SubSpecUsers,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// testId
	TaskId                    *string                           `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	UnsubSpecAudioUsers       []*string                         `json:"UnsubSpecAudioUsers,omitempty" xml:"UnsubSpecAudioUsers,omitempty" type:"Repeated"`
	UnsubSpecCameraUsers      []*string                         `json:"UnsubSpecCameraUsers,omitempty" xml:"UnsubSpecCameraUsers,omitempty" type:"Repeated"`
	UnsubSpecShareScreenUsers []*string                         `json:"UnsubSpecShareScreenUsers,omitempty" xml:"UnsubSpecShareScreenUsers,omitempty" type:"Repeated"`
	UserPanes                 []*UpdateMPUTaskRequestUserPanes  `json:"UserPanes,omitempty" xml:"UserPanes,omitempty" type:"Repeated"`
	Watermarks                []*UpdateMPUTaskRequestWatermarks `json:"Watermarks,omitempty" xml:"Watermarks,omitempty" type:"Repeated"`
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

func (s UpdateMPUTaskRequestClockWidgets) String() string {
	return tea.Prettify(s)
}

func (s UpdateMPUTaskRequestClockWidgets) GoString() string {
	return s.String()
}

func (s *UpdateMPUTaskRequestClockWidgets) SetAlpha(v float32) *UpdateMPUTaskRequestClockWidgets {
	s.Alpha = &v
	return s
}

func (s *UpdateMPUTaskRequestClockWidgets) SetBorderColor(v int64) *UpdateMPUTaskRequestClockWidgets {
	s.BorderColor = &v
	return s
}

func (s *UpdateMPUTaskRequestClockWidgets) SetBorderWidth(v int32) *UpdateMPUTaskRequestClockWidgets {
	s.BorderWidth = &v
	return s
}

func (s *UpdateMPUTaskRequestClockWidgets) SetBox(v bool) *UpdateMPUTaskRequestClockWidgets {
	s.Box = &v
	return s
}

func (s *UpdateMPUTaskRequestClockWidgets) SetBoxBorderWidth(v int32) *UpdateMPUTaskRequestClockWidgets {
	s.BoxBorderWidth = &v
	return s
}

func (s *UpdateMPUTaskRequestClockWidgets) SetBoxColor(v int64) *UpdateMPUTaskRequestClockWidgets {
	s.BoxColor = &v
	return s
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
	Images []*UpdateMPUTaskRequestUserPanesImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
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
	SourceType *string                               `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Texts      []*UpdateMPUTaskRequestUserPanesTexts `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
	// example:
	//
	// TestUserID
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// 0.7576
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	// example:
	//
	// 0
	ZOrder *int32 `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
}

func (s UpdateMPUTaskRequestUserPanesTexts) String() string {
	return tea.Prettify(s)
}

func (s UpdateMPUTaskRequestUserPanesTexts) GoString() string {
	return s.String()
}

func (s *UpdateMPUTaskRequestUserPanesTexts) SetAlpha(v float32) *UpdateMPUTaskRequestUserPanesTexts {
	s.Alpha = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanesTexts) SetBorderColor(v int64) *UpdateMPUTaskRequestUserPanesTexts {
	s.BorderColor = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanesTexts) SetBorderWidth(v int32) *UpdateMPUTaskRequestUserPanesTexts {
	s.BorderWidth = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanesTexts) SetBox(v bool) *UpdateMPUTaskRequestUserPanesTexts {
	s.Box = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanesTexts) SetBoxBorderWidth(v int32) *UpdateMPUTaskRequestUserPanesTexts {
	s.BoxBorderWidth = &v
	return s
}

func (s *UpdateMPUTaskRequestUserPanesTexts) SetBoxColor(v int64) *UpdateMPUTaskRequestUserPanesTexts {
	s.BoxColor = &v
	return s
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
	// example:
	//
	// 0
	Alpha *float32 `json:"Alpha,omitempty" xml:"Alpha,omitempty"`
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
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMPUTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UpdateMPUTaskResponse) SetStatusCode(v int32) *UpdateMPUTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMPUTaskResponse) SetBody(v *UpdateMPUTaskResponseBody) *UpdateMPUTaskResponse {
	s.Body = v
	return s
}

type UpdateRecordTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// yourChannelId
	ChannelId               *string   `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	CropMode                *int64    `json:"CropMode,omitempty" xml:"CropMode,omitempty"`
	LayoutIds               []*int64  `json:"LayoutIds,omitempty" xml:"LayoutIds,omitempty" type:"Repeated"`
	MediaEncode             *int64    `json:"MediaEncode,omitempty" xml:"MediaEncode,omitempty"`
	OwnerId                 *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskProfile *string `json:"TaskProfile,omitempty" xml:"TaskProfile,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 76dasgb****
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

func (s *UpdateRecordTaskRequest) SetCropMode(v int64) *UpdateRecordTaskRequest {
	s.CropMode = &v
	return s
}

func (s *UpdateRecordTaskRequest) SetLayoutIds(v []*int64) *UpdateRecordTaskRequest {
	s.LayoutIds = v
	return s
}

func (s *UpdateRecordTaskRequest) SetMediaEncode(v int64) *UpdateRecordTaskRequest {
	s.MediaEncode = &v
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

func (s *UpdateRecordTaskRequest) SetTaskProfile(v string) *UpdateRecordTaskRequest {
	s.TaskProfile = &v
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
	Images []*UpdateRecordTaskRequestUserPanesImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PaneId *int32 `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
	// example:
	//
	// camera
	SourceType *string                                  `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Texts      []*UpdateRecordTaskRequestUserPanesTexts `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
	// example:
	//
	// TestId
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
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
	// 0.7576
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
	// example:
	//
	// 0
	ZOrder *int32 `json:"ZOrder,omitempty" xml:"ZOrder,omitempty"`
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
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRecordTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UpdateRecordTaskResponse) SetStatusCode(v int32) *UpdateRecordTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRecordTaskResponse) SetBody(v *UpdateRecordTaskResponseBody) *UpdateRecordTaskResponse {
	s.Body = v
	return s
}

type UpdateRecordTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// yourAppId
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 0
	BackgroundColor *int32                                     `json:"BackgroundColor,omitempty" xml:"BackgroundColor,omitempty"`
	Backgrounds     []*UpdateRecordTemplateRequestBackgrounds  `json:"Backgrounds,omitempty" xml:"Backgrounds,omitempty" type:"Repeated"`
	ClockWidgets    []*UpdateRecordTemplateRequestClockWidgets `json:"ClockWidgets,omitempty" xml:"ClockWidgets,omitempty" type:"Repeated"`
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
	// 1111
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
	// rtc-record-pre
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
	TaskProfile *string `json:"TaskProfile,omitempty" xml:"TaskProfile,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 76dasgb****
	TemplateId *string                                  `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Watermarks []*UpdateRecordTemplateRequestWatermarks `json:"Watermarks,omitempty" xml:"Watermarks,omitempty" type:"Repeated"`
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

type UpdateRecordTemplateRequestBackgrounds struct {
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
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 76dasgb****
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRecordTemplateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UpdateRecordTemplateResponse) SetStatusCode(v int32) *UpdateRecordTemplateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRecordTemplateResponse) SetBody(v *UpdateRecordTemplateResponseBody) *UpdateRecordTemplateResponse {
	s.Body = v
	return s
}

type UpdateStreamingOutRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// eo85****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testid
	ChannelId    *string                                  `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ClockWidgets []*UpdateStreamingOutRequestClockWidgets `json:"ClockWidgets,omitempty" xml:"ClockWidgets,omitempty" type:"Repeated"`
	Images       []*UpdateStreamingOutRequestImages       `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	Panes        []*UpdateStreamingOutRequestPanes        `json:"Panes,omitempty" xml:"Panes,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 567
	TemplateId *string                           `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	Texts      []*UpdateStreamingOutRequestTexts `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
}

func (s UpdateStreamingOutRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateStreamingOutRequest) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutRequest) SetAppId(v string) *UpdateStreamingOutRequest {
	s.AppId = &v
	return s
}

func (s *UpdateStreamingOutRequest) SetChannelId(v string) *UpdateStreamingOutRequest {
	s.ChannelId = &v
	return s
}

func (s *UpdateStreamingOutRequest) SetClockWidgets(v []*UpdateStreamingOutRequestClockWidgets) *UpdateStreamingOutRequest {
	s.ClockWidgets = v
	return s
}

func (s *UpdateStreamingOutRequest) SetImages(v []*UpdateStreamingOutRequestImages) *UpdateStreamingOutRequest {
	s.Images = v
	return s
}

func (s *UpdateStreamingOutRequest) SetPanes(v []*UpdateStreamingOutRequestPanes) *UpdateStreamingOutRequest {
	s.Panes = v
	return s
}

func (s *UpdateStreamingOutRequest) SetTaskId(v string) *UpdateStreamingOutRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateStreamingOutRequest) SetTemplateId(v string) *UpdateStreamingOutRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateStreamingOutRequest) SetTexts(v []*UpdateStreamingOutRequestTexts) *UpdateStreamingOutRequest {
	s.Texts = v
	return s
}

type UpdateStreamingOutRequestClockWidgets struct {
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
	BoxBorderw *int32                                         `json:"BoxBorderw,omitempty" xml:"BoxBorderw,omitempty"`
	BoxColor   *UpdateStreamingOutRequestClockWidgetsBoxColor `json:"BoxColor,omitempty" xml:"BoxColor,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Font      *int32                                          `json:"Font,omitempty" xml:"Font,omitempty"`
	FontColor *UpdateStreamingOutRequestClockWidgetsFontColor `json:"FontColor,omitempty" xml:"FontColor,omitempty" type:"Struct"`
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

func (s UpdateStreamingOutRequestClockWidgets) String() string {
	return tea.Prettify(s)
}

func (s UpdateStreamingOutRequestClockWidgets) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutRequestClockWidgets) SetAlpha(v float64) *UpdateStreamingOutRequestClockWidgets {
	s.Alpha = &v
	return s
}

func (s *UpdateStreamingOutRequestClockWidgets) SetBoxAlpha(v float64) *UpdateStreamingOutRequestClockWidgets {
	s.BoxAlpha = &v
	return s
}

func (s *UpdateStreamingOutRequestClockWidgets) SetBoxBorderw(v int32) *UpdateStreamingOutRequestClockWidgets {
	s.BoxBorderw = &v
	return s
}

func (s *UpdateStreamingOutRequestClockWidgets) SetBoxColor(v *UpdateStreamingOutRequestClockWidgetsBoxColor) *UpdateStreamingOutRequestClockWidgets {
	s.BoxColor = v
	return s
}

func (s *UpdateStreamingOutRequestClockWidgets) SetFont(v int32) *UpdateStreamingOutRequestClockWidgets {
	s.Font = &v
	return s
}

func (s *UpdateStreamingOutRequestClockWidgets) SetFontColor(v *UpdateStreamingOutRequestClockWidgetsFontColor) *UpdateStreamingOutRequestClockWidgets {
	s.FontColor = v
	return s
}

func (s *UpdateStreamingOutRequestClockWidgets) SetFontSize(v int32) *UpdateStreamingOutRequestClockWidgets {
	s.FontSize = &v
	return s
}

func (s *UpdateStreamingOutRequestClockWidgets) SetHasBox(v bool) *UpdateStreamingOutRequestClockWidgets {
	s.HasBox = &v
	return s
}

func (s *UpdateStreamingOutRequestClockWidgets) SetLayer(v int32) *UpdateStreamingOutRequestClockWidgets {
	s.Layer = &v
	return s
}

func (s *UpdateStreamingOutRequestClockWidgets) SetX(v float64) *UpdateStreamingOutRequestClockWidgets {
	s.X = &v
	return s
}

func (s *UpdateStreamingOutRequestClockWidgets) SetY(v float64) *UpdateStreamingOutRequestClockWidgets {
	s.Y = &v
	return s
}

func (s *UpdateStreamingOutRequestClockWidgets) SetZone(v int32) *UpdateStreamingOutRequestClockWidgets {
	s.Zone = &v
	return s
}

type UpdateStreamingOutRequestClockWidgetsBoxColor struct {
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

func (s UpdateStreamingOutRequestClockWidgetsBoxColor) String() string {
	return tea.Prettify(s)
}

func (s UpdateStreamingOutRequestClockWidgetsBoxColor) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutRequestClockWidgetsBoxColor) SetB(v int32) *UpdateStreamingOutRequestClockWidgetsBoxColor {
	s.B = &v
	return s
}

func (s *UpdateStreamingOutRequestClockWidgetsBoxColor) SetG(v int32) *UpdateStreamingOutRequestClockWidgetsBoxColor {
	s.G = &v
	return s
}

func (s *UpdateStreamingOutRequestClockWidgetsBoxColor) SetR(v int32) *UpdateStreamingOutRequestClockWidgetsBoxColor {
	s.R = &v
	return s
}

type UpdateStreamingOutRequestClockWidgetsFontColor struct {
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

func (s UpdateStreamingOutRequestClockWidgetsFontColor) String() string {
	return tea.Prettify(s)
}

func (s UpdateStreamingOutRequestClockWidgetsFontColor) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutRequestClockWidgetsFontColor) SetB(v int32) *UpdateStreamingOutRequestClockWidgetsFontColor {
	s.B = &v
	return s
}

func (s *UpdateStreamingOutRequestClockWidgetsFontColor) SetG(v int32) *UpdateStreamingOutRequestClockWidgetsFontColor {
	s.G = &v
	return s
}

func (s *UpdateStreamingOutRequestClockWidgetsFontColor) SetR(v int32) *UpdateStreamingOutRequestClockWidgetsFontColor {
	s.R = &v
	return s
}

type UpdateStreamingOutRequestImages struct {
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

func (s UpdateStreamingOutRequestImages) String() string {
	return tea.Prettify(s)
}

func (s UpdateStreamingOutRequestImages) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutRequestImages) SetAlpha(v float64) *UpdateStreamingOutRequestImages {
	s.Alpha = &v
	return s
}

func (s *UpdateStreamingOutRequestImages) SetHeight(v float64) *UpdateStreamingOutRequestImages {
	s.Height = &v
	return s
}

func (s *UpdateStreamingOutRequestImages) SetImageCropMode(v int32) *UpdateStreamingOutRequestImages {
	s.ImageCropMode = &v
	return s
}

func (s *UpdateStreamingOutRequestImages) SetLayer(v int32) *UpdateStreamingOutRequestImages {
	s.Layer = &v
	return s
}

func (s *UpdateStreamingOutRequestImages) SetUrl(v string) *UpdateStreamingOutRequestImages {
	s.Url = &v
	return s
}

func (s *UpdateStreamingOutRequestImages) SetWidth(v float64) *UpdateStreamingOutRequestImages {
	s.Width = &v
	return s
}

func (s *UpdateStreamingOutRequestImages) SetX(v float64) *UpdateStreamingOutRequestImages {
	s.X = &v
	return s
}

func (s *UpdateStreamingOutRequestImages) SetY(v float64) *UpdateStreamingOutRequestImages {
	s.Y = &v
	return s
}

type UpdateStreamingOutRequestPanes struct {
	Images []*UpdateStreamingOutRequestPanesImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	PaneCropMode *int32 `json:"PaneCropMode,omitempty" xml:"PaneCropMode,omitempty"`
	// example:
	//
	// 1
	PaneId *int32 `json:"PaneId,omitempty" xml:"PaneId,omitempty"`
	// example:
	//
	// 22
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// video
	SourceType *string                                `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Texts      []*UpdateStreamingOutRequestPanesTexts `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
}

func (s UpdateStreamingOutRequestPanes) String() string {
	return tea.Prettify(s)
}

func (s UpdateStreamingOutRequestPanes) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutRequestPanes) SetImages(v []*UpdateStreamingOutRequestPanesImages) *UpdateStreamingOutRequestPanes {
	s.Images = v
	return s
}

func (s *UpdateStreamingOutRequestPanes) SetPaneCropMode(v int32) *UpdateStreamingOutRequestPanes {
	s.PaneCropMode = &v
	return s
}

func (s *UpdateStreamingOutRequestPanes) SetPaneId(v int32) *UpdateStreamingOutRequestPanes {
	s.PaneId = &v
	return s
}

func (s *UpdateStreamingOutRequestPanes) SetSource(v string) *UpdateStreamingOutRequestPanes {
	s.Source = &v
	return s
}

func (s *UpdateStreamingOutRequestPanes) SetSourceType(v string) *UpdateStreamingOutRequestPanes {
	s.SourceType = &v
	return s
}

func (s *UpdateStreamingOutRequestPanes) SetTexts(v []*UpdateStreamingOutRequestPanesTexts) *UpdateStreamingOutRequestPanes {
	s.Texts = v
	return s
}

type UpdateStreamingOutRequestPanesImages struct {
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

func (s UpdateStreamingOutRequestPanesImages) String() string {
	return tea.Prettify(s)
}

func (s UpdateStreamingOutRequestPanesImages) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutRequestPanesImages) SetAlpha(v float64) *UpdateStreamingOutRequestPanesImages {
	s.Alpha = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesImages) SetHeight(v float64) *UpdateStreamingOutRequestPanesImages {
	s.Height = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesImages) SetLayer(v int32) *UpdateStreamingOutRequestPanesImages {
	s.Layer = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesImages) SetPaneImageCropMode(v int32) *UpdateStreamingOutRequestPanesImages {
	s.PaneImageCropMode = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesImages) SetUrl(v string) *UpdateStreamingOutRequestPanesImages {
	s.Url = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesImages) SetWidth(v float64) *UpdateStreamingOutRequestPanesImages {
	s.Width = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesImages) SetX(v float64) *UpdateStreamingOutRequestPanesImages {
	s.X = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesImages) SetY(v float64) *UpdateStreamingOutRequestPanesImages {
	s.Y = &v
	return s
}

type UpdateStreamingOutRequestPanesTexts struct {
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
	BoxColor   *UpdateStreamingOutRequestPanesTextsBoxColor `json:"BoxColor,omitempty" xml:"BoxColor,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Font      *int32                                        `json:"Font,omitempty" xml:"Font,omitempty"`
	FontColor *UpdateStreamingOutRequestPanesTextsFontColor `json:"FontColor,omitempty" xml:"FontColor,omitempty" type:"Struct"`
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

func (s UpdateStreamingOutRequestPanesTexts) String() string {
	return tea.Prettify(s)
}

func (s UpdateStreamingOutRequestPanesTexts) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutRequestPanesTexts) SetAlpha(v float64) *UpdateStreamingOutRequestPanesTexts {
	s.Alpha = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesTexts) SetBoxAlpha(v float64) *UpdateStreamingOutRequestPanesTexts {
	s.BoxAlpha = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesTexts) SetBoxBorderw(v int32) *UpdateStreamingOutRequestPanesTexts {
	s.BoxBorderw = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesTexts) SetBoxColor(v *UpdateStreamingOutRequestPanesTextsBoxColor) *UpdateStreamingOutRequestPanesTexts {
	s.BoxColor = v
	return s
}

func (s *UpdateStreamingOutRequestPanesTexts) SetFont(v int32) *UpdateStreamingOutRequestPanesTexts {
	s.Font = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesTexts) SetFontColor(v *UpdateStreamingOutRequestPanesTextsFontColor) *UpdateStreamingOutRequestPanesTexts {
	s.FontColor = v
	return s
}

func (s *UpdateStreamingOutRequestPanesTexts) SetFontSize(v int32) *UpdateStreamingOutRequestPanesTexts {
	s.FontSize = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesTexts) SetHasBox(v bool) *UpdateStreamingOutRequestPanesTexts {
	s.HasBox = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesTexts) SetLayer(v int32) *UpdateStreamingOutRequestPanesTexts {
	s.Layer = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesTexts) SetTexture(v string) *UpdateStreamingOutRequestPanesTexts {
	s.Texture = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesTexts) SetX(v float64) *UpdateStreamingOutRequestPanesTexts {
	s.X = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesTexts) SetY(v float64) *UpdateStreamingOutRequestPanesTexts {
	s.Y = &v
	return s
}

type UpdateStreamingOutRequestPanesTextsBoxColor struct {
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

func (s UpdateStreamingOutRequestPanesTextsBoxColor) String() string {
	return tea.Prettify(s)
}

func (s UpdateStreamingOutRequestPanesTextsBoxColor) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutRequestPanesTextsBoxColor) SetB(v int32) *UpdateStreamingOutRequestPanesTextsBoxColor {
	s.B = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesTextsBoxColor) SetG(v int32) *UpdateStreamingOutRequestPanesTextsBoxColor {
	s.G = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesTextsBoxColor) SetR(v int32) *UpdateStreamingOutRequestPanesTextsBoxColor {
	s.R = &v
	return s
}

type UpdateStreamingOutRequestPanesTextsFontColor struct {
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

func (s UpdateStreamingOutRequestPanesTextsFontColor) String() string {
	return tea.Prettify(s)
}

func (s UpdateStreamingOutRequestPanesTextsFontColor) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutRequestPanesTextsFontColor) SetB(v int32) *UpdateStreamingOutRequestPanesTextsFontColor {
	s.B = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesTextsFontColor) SetG(v int32) *UpdateStreamingOutRequestPanesTextsFontColor {
	s.G = &v
	return s
}

func (s *UpdateStreamingOutRequestPanesTextsFontColor) SetR(v int32) *UpdateStreamingOutRequestPanesTextsFontColor {
	s.R = &v
	return s
}

type UpdateStreamingOutRequestTexts struct {
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
	BoxBorderw *int32                                  `json:"BoxBorderw,omitempty" xml:"BoxBorderw,omitempty"`
	BoxColor   *UpdateStreamingOutRequestTextsBoxColor `json:"BoxColor,omitempty" xml:"BoxColor,omitempty" type:"Struct"`
	// example:
	//
	// 0
	Font      *int32                                   `json:"Font,omitempty" xml:"Font,omitempty"`
	FontColor *UpdateStreamingOutRequestTextsFontColor `json:"FontColor,omitempty" xml:"FontColor,omitempty" type:"Struct"`
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

func (s UpdateStreamingOutRequestTexts) String() string {
	return tea.Prettify(s)
}

func (s UpdateStreamingOutRequestTexts) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutRequestTexts) SetAlpha(v float64) *UpdateStreamingOutRequestTexts {
	s.Alpha = &v
	return s
}

func (s *UpdateStreamingOutRequestTexts) SetBoxAlpha(v float64) *UpdateStreamingOutRequestTexts {
	s.BoxAlpha = &v
	return s
}

func (s *UpdateStreamingOutRequestTexts) SetBoxBorderw(v int32) *UpdateStreamingOutRequestTexts {
	s.BoxBorderw = &v
	return s
}

func (s *UpdateStreamingOutRequestTexts) SetBoxColor(v *UpdateStreamingOutRequestTextsBoxColor) *UpdateStreamingOutRequestTexts {
	s.BoxColor = v
	return s
}

func (s *UpdateStreamingOutRequestTexts) SetFont(v int32) *UpdateStreamingOutRequestTexts {
	s.Font = &v
	return s
}

func (s *UpdateStreamingOutRequestTexts) SetFontColor(v *UpdateStreamingOutRequestTextsFontColor) *UpdateStreamingOutRequestTexts {
	s.FontColor = v
	return s
}

func (s *UpdateStreamingOutRequestTexts) SetFontSize(v int32) *UpdateStreamingOutRequestTexts {
	s.FontSize = &v
	return s
}

func (s *UpdateStreamingOutRequestTexts) SetHasBox(v bool) *UpdateStreamingOutRequestTexts {
	s.HasBox = &v
	return s
}

func (s *UpdateStreamingOutRequestTexts) SetLayer(v int32) *UpdateStreamingOutRequestTexts {
	s.Layer = &v
	return s
}

func (s *UpdateStreamingOutRequestTexts) SetTexture(v string) *UpdateStreamingOutRequestTexts {
	s.Texture = &v
	return s
}

func (s *UpdateStreamingOutRequestTexts) SetX(v float64) *UpdateStreamingOutRequestTexts {
	s.X = &v
	return s
}

func (s *UpdateStreamingOutRequestTexts) SetY(v float64) *UpdateStreamingOutRequestTexts {
	s.Y = &v
	return s
}

type UpdateStreamingOutRequestTextsBoxColor struct {
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

func (s UpdateStreamingOutRequestTextsBoxColor) String() string {
	return tea.Prettify(s)
}

func (s UpdateStreamingOutRequestTextsBoxColor) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutRequestTextsBoxColor) SetB(v int32) *UpdateStreamingOutRequestTextsBoxColor {
	s.B = &v
	return s
}

func (s *UpdateStreamingOutRequestTextsBoxColor) SetG(v int32) *UpdateStreamingOutRequestTextsBoxColor {
	s.G = &v
	return s
}

func (s *UpdateStreamingOutRequestTextsBoxColor) SetR(v int32) *UpdateStreamingOutRequestTextsBoxColor {
	s.R = &v
	return s
}

type UpdateStreamingOutRequestTextsFontColor struct {
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

func (s UpdateStreamingOutRequestTextsFontColor) String() string {
	return tea.Prettify(s)
}

func (s UpdateStreamingOutRequestTextsFontColor) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutRequestTextsFontColor) SetB(v int32) *UpdateStreamingOutRequestTextsFontColor {
	s.B = &v
	return s
}

func (s *UpdateStreamingOutRequestTextsFontColor) SetG(v int32) *UpdateStreamingOutRequestTextsFontColor {
	s.G = &v
	return s
}

func (s *UpdateStreamingOutRequestTextsFontColor) SetR(v int32) *UpdateStreamingOutRequestTextsFontColor {
	s.R = &v
	return s
}

type UpdateStreamingOutResponseBody struct {
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateStreamingOutResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateStreamingOutResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutResponseBody) SetRequestId(v string) *UpdateStreamingOutResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateStreamingOutResponseBody) SetTaskId(v string) *UpdateStreamingOutResponseBody {
	s.TaskId = &v
	return s
}

type UpdateStreamingOutResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateStreamingOutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateStreamingOutResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateStreamingOutResponse) GoString() string {
	return s.String()
}

func (s *UpdateStreamingOutResponse) SetHeaders(v map[string]*string) *UpdateStreamingOutResponse {
	s.Headers = v
	return s
}

func (s *UpdateStreamingOutResponse) SetStatusCode(v int32) *UpdateStreamingOutResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateStreamingOutResponse) SetBody(v *UpdateStreamingOutResponseBody) *UpdateStreamingOutResponse {
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

// @param request - AddRecordTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddRecordTemplateResponse
func (client *Client) AddRecordTemplateWithOptions(request *AddRecordTemplateRequest, runtime *util.RuntimeOptions) (_result *AddRecordTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.BackgroundColor)) {
		query["BackgroundColor"] = request.BackgroundColor
	}

	if !tea.BoolValue(util.IsUnset(request.Backgrounds)) {
		query["Backgrounds"] = request.Backgrounds
	}

	if !tea.BoolValue(util.IsUnset(request.ClockWidgets)) {
		query["ClockWidgets"] = request.ClockWidgets
	}

	if !tea.BoolValue(util.IsUnset(request.DelayStopTime)) {
		query["DelayStopTime"] = request.DelayStopTime
	}

	if !tea.BoolValue(util.IsUnset(request.EnableM3u8DateTime)) {
		query["EnableM3u8DateTime"] = request.EnableM3u8DateTime
	}

	if !tea.BoolValue(util.IsUnset(request.FileSplitInterval)) {
		query["FileSplitInterval"] = request.FileSplitInterval
	}

	if !tea.BoolValue(util.IsUnset(request.Formats)) {
		query["Formats"] = request.Formats
	}

	if !tea.BoolValue(util.IsUnset(request.HttpCallbackUrl)) {
		query["HttpCallbackUrl"] = request.HttpCallbackUrl
	}

	if !tea.BoolValue(util.IsUnset(request.LayoutIds)) {
		query["LayoutIds"] = request.LayoutIds
	}

	if !tea.BoolValue(util.IsUnset(request.MediaEncode)) {
		query["MediaEncode"] = request.MediaEncode
	}

	if !tea.BoolValue(util.IsUnset(request.MnsQueue)) {
		query["MnsQueue"] = request.MnsQueue
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OssBucket)) {
		query["OssBucket"] = request.OssBucket
	}

	if !tea.BoolValue(util.IsUnset(request.OssEndpoint)) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	if !tea.BoolValue(util.IsUnset(request.OssFilePrefix)) {
		query["OssFilePrefix"] = request.OssFilePrefix
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskProfile)) {
		query["TaskProfile"] = request.TaskProfile
	}

	if !tea.BoolValue(util.IsUnset(request.Watermarks)) {
		query["Watermarks"] = request.Watermarks
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddRecordTemplate"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &AddRecordTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - AddRecordTemplateRequest
//
// @return AddRecordTemplateResponse
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

// Summary:
//
// 创建应用推流模版
//
// @param tmpReq - CreateAppStreamingOutTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppStreamingOutTemplateResponse
func (client *Client) CreateAppStreamingOutTemplateWithOptions(tmpReq *CreateAppStreamingOutTemplateRequest, runtime *util.RuntimeOptions) (_result *CreateAppStreamingOutTemplateResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateAppStreamingOutTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.StreamingOutTemplate)) {
		request.StreamingOutTemplateShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StreamingOutTemplate, tea.String("StreamingOutTemplate"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.StreamingOutTemplateShrink)) {
		query["StreamingOutTemplate"] = request.StreamingOutTemplateShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAppStreamingOutTemplate"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAppStreamingOutTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建应用推流模版
//
// @param request - CreateAppStreamingOutTemplateRequest
//
// @return CreateAppStreamingOutTemplateResponse
func (client *Client) CreateAppStreamingOutTemplate(request *CreateAppStreamingOutTemplateRequest) (_result *CreateAppStreamingOutTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAppStreamingOutTemplateResponse{}
	_body, _err := client.CreateAppStreamingOutTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateAutoLiveStreamRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAutoLiveStreamRuleResponse
func (client *Client) CreateAutoLiveStreamRuleWithOptions(request *CreateAutoLiveStreamRuleRequest, runtime *util.RuntimeOptions) (_result *CreateAutoLiveStreamRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.CallBack)) {
		query["CallBack"] = request.CallBack
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelIdPrefixes)) {
		query["ChannelIdPrefixes"] = request.ChannelIdPrefixes
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelIds)) {
		query["ChannelIds"] = request.ChannelIds
	}

	if !tea.BoolValue(util.IsUnset(request.MediaEncode)) {
		query["MediaEncode"] = request.MediaEncode
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PlayDomain)) {
		query["PlayDomain"] = request.PlayDomain
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAutoLiveStreamRule"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAutoLiveStreamRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateAutoLiveStreamRuleRequest
//
// @return CreateAutoLiveStreamRuleResponse
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

// @param request - CreateEventSubscribeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEventSubscribeResponse
func (client *Client) CreateEventSubscribeWithOptions(request *CreateEventSubscribeRequest, runtime *util.RuntimeOptions) (_result *CreateEventSubscribeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.CallbackUrl)) {
		query["CallbackUrl"] = request.CallbackUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Events)) {
		query["Events"] = request.Events
	}

	if !tea.BoolValue(util.IsUnset(request.NeedCallbackAuth)) {
		query["NeedCallbackAuth"] = request.NeedCallbackAuth
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Role)) {
		query["Role"] = request.Role
	}

	if !tea.BoolValue(util.IsUnset(request.Users)) {
		query["Users"] = request.Users
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateEventSubscribe"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateEventSubscribeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateEventSubscribeRequest
//
// @return CreateEventSubscribeResponse
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

// @param request - CreateMPULayoutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateMPULayoutResponse
func (client *Client) CreateMPULayoutWithOptions(request *CreateMPULayoutRequest, runtime *util.RuntimeOptions) (_result *CreateMPULayoutResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AudioMixCount)) {
		query["AudioMixCount"] = request.AudioMixCount
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Panes)) {
		query["Panes"] = request.Panes
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMPULayout"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMPULayoutResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateMPULayoutRequest
//
// @return CreateMPULayoutResponse
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

// Summary:
//
// 删除应用推流模版
//
// @param tmpReq - DeleteAppStreamingOutTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAppStreamingOutTemplateResponse
func (client *Client) DeleteAppStreamingOutTemplateWithOptions(tmpReq *DeleteAppStreamingOutTemplateRequest, runtime *util.RuntimeOptions) (_result *DeleteAppStreamingOutTemplateResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteAppStreamingOutTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.StreamingOutTemplate)) {
		request.StreamingOutTemplateShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StreamingOutTemplate, tea.String("StreamingOutTemplate"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.StreamingOutTemplateShrink)) {
		query["StreamingOutTemplate"] = request.StreamingOutTemplateShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAppStreamingOutTemplate"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAppStreamingOutTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除应用推流模版
//
// @param request - DeleteAppStreamingOutTemplateRequest
//
// @return DeleteAppStreamingOutTemplateResponse
func (client *Client) DeleteAppStreamingOutTemplate(request *DeleteAppStreamingOutTemplateRequest) (_result *DeleteAppStreamingOutTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAppStreamingOutTemplateResponse{}
	_body, _err := client.DeleteAppStreamingOutTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DeleteAutoLiveStreamRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteAutoLiveStreamRuleResponse
func (client *Client) DeleteAutoLiveStreamRuleWithOptions(request *DeleteAutoLiveStreamRuleRequest, runtime *util.RuntimeOptions) (_result *DeleteAutoLiveStreamRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteAutoLiveStreamRule"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteAutoLiveStreamRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteAutoLiveStreamRuleRequest
//
// @return DeleteAutoLiveStreamRuleResponse
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

// @param request - DeleteChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteChannelResponse
func (client *Client) DeleteChannelWithOptions(request *DeleteChannelRequest, runtime *util.RuntimeOptions) (_result *DeleteChannelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteChannel"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteChannelRequest
//
// @return DeleteChannelResponse
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

// @param request - DeleteEventSubscribeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteEventSubscribeResponse
func (client *Client) DeleteEventSubscribeWithOptions(request *DeleteEventSubscribeRequest, runtime *util.RuntimeOptions) (_result *DeleteEventSubscribeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SubscribeId)) {
		query["SubscribeId"] = request.SubscribeId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteEventSubscribe"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteEventSubscribeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteEventSubscribeRequest
//
// @return DeleteEventSubscribeResponse
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

// @param request - DeleteMPULayoutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteMPULayoutResponse
func (client *Client) DeleteMPULayoutWithOptions(request *DeleteMPULayoutRequest, runtime *util.RuntimeOptions) (_result *DeleteMPULayoutResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.LayoutId)) {
		query["LayoutId"] = request.LayoutId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMPULayout"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMPULayoutResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteMPULayoutRequest
//
// @return DeleteMPULayoutResponse
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

// @param request - DeleteRecordTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRecordTemplateResponse
func (client *Client) DeleteRecordTemplateWithOptions(request *DeleteRecordTemplateRequest, runtime *util.RuntimeOptions) (_result *DeleteRecordTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRecordTemplate"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRecordTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DeleteRecordTemplateRequest
//
// @return DeleteRecordTemplateResponse
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

// Summary:
//
// 查看AppKey
//
// @param request - DescribeAppKeyRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppKeyResponse
func (client *Client) DescribeAppKeyWithOptions(request *DescribeAppKeyRequest, runtime *util.RuntimeOptions) (_result *DescribeAppKeyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAppKey"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAppKeyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查看AppKey
//
// @param request - DescribeAppKeyRequest
//
// @return DescribeAppKeyResponse
func (client *Client) DescribeAppKey(request *DescribeAppKeyRequest) (_result *DescribeAppKeyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAppKeyResponse{}
	_body, _err := client.DescribeAppKeyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 应用推流模版列表
//
// @param tmpReq - DescribeAppStreamingOutTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppStreamingOutTemplatesResponse
func (client *Client) DescribeAppStreamingOutTemplatesWithOptions(tmpReq *DescribeAppStreamingOutTemplatesRequest, runtime *util.RuntimeOptions) (_result *DescribeAppStreamingOutTemplatesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DescribeAppStreamingOutTemplatesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Condition)) {
		request.ConditionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Condition, tea.String("Condition"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ConditionShrink)) {
		query["Condition"] = request.ConditionShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAppStreamingOutTemplates"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAppStreamingOutTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 应用推流模版列表
//
// @param request - DescribeAppStreamingOutTemplatesRequest
//
// @return DescribeAppStreamingOutTemplatesResponse
func (client *Client) DescribeAppStreamingOutTemplates(request *DescribeAppStreamingOutTemplatesRequest) (_result *DescribeAppStreamingOutTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAppStreamingOutTemplatesResponse{}
	_body, _err := client.DescribeAppStreamingOutTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// App列表
//
// @param request - DescribeAppsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppsResponse
func (client *Client) DescribeAppsWithOptions(request *DescribeAppsRequest, runtime *util.RuntimeOptions) (_result *DescribeAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppVersion)) {
		query["AppVersion"] = request.AppVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApps"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// App列表
//
// @param request - DescribeAppsRequest
//
// @return DescribeAppsResponse
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

// @param request - DescribeAutoLiveStreamRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAutoLiveStreamRuleResponse
func (client *Client) DescribeAutoLiveStreamRuleWithOptions(request *DescribeAutoLiveStreamRuleRequest, runtime *util.RuntimeOptions) (_result *DescribeAutoLiveStreamRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAutoLiveStreamRule"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAutoLiveStreamRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeAutoLiveStreamRuleRequest
//
// @return DescribeAutoLiveStreamRuleResponse
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

// Summary:
//
// 调用DescribeCall获取单次通信详情。
//
// @param request - DescribeCallRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCallResponse
func (client *Client) DescribeCallWithOptions(request *DescribeCallRequest, runtime *util.RuntimeOptions) (_result *DescribeCallResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.CreatedTs)) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !tea.BoolValue(util.IsUnset(request.DestroyedTs)) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	if !tea.BoolValue(util.IsUnset(request.ExtDataType)) {
		query["ExtDataType"] = request.ExtDataType
	}

	if !tea.BoolValue(util.IsUnset(request.QueryExpInfo)) {
		query["QueryExpInfo"] = request.QueryExpInfo
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCall"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCallResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用DescribeCall获取单次通信详情。
//
// @param request - DescribeCallRequest
//
// @return DescribeCallResponse
func (client *Client) DescribeCall(request *DescribeCallRequest) (_result *DescribeCallResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCallResponse{}
	_body, _err := client.DescribeCallWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 调用DescribeCallList分页查询时间范围内创建的通信信息。
//
// @param request - DescribeCallListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeCallListResponse
func (client *Client) DescribeCallListWithOptions(request *DescribeCallListRequest, runtime *util.RuntimeOptions) (_result *DescribeCallListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.CallStatus)) {
		query["CallStatus"] = request.CallStatus
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTs)) {
		query["EndTs"] = request.EndTs
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		query["OrderBy"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryMode)) {
		query["QueryMode"] = request.QueryMode
	}

	if !tea.BoolValue(util.IsUnset(request.StartTs)) {
		query["StartTs"] = request.StartTs
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCallList"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCallListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用DescribeCallList分页查询时间范围内创建的通信信息。
//
// @param request - DescribeCallListRequest
//
// @return DescribeCallListResponse
func (client *Client) DescribeCallList(request *DescribeCallListRequest) (_result *DescribeCallListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCallListResponse{}
	_body, _err := client.DescribeCallListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// DescribeChannel
//
// @param request - DescribeChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelResponse
func (client *Client) DescribeChannelWithOptions(request *DescribeChannelRequest, runtime *util.RuntimeOptions) (_result *DescribeChannelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeChannel"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// DescribeChannel
//
// @param request - DescribeChannelRequest
//
// @return DescribeChannelResponse
func (client *Client) DescribeChannel(request *DescribeChannelRequest) (_result *DescribeChannelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeChannelResponse{}
	_body, _err := client.DescribeChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询频道的所有参会者
//
// @param request - DescribeChannelAllUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelAllUsersResponse
func (client *Client) DescribeChannelAllUsersWithOptions(request *DescribeChannelAllUsersRequest, runtime *util.RuntimeOptions) (_result *DescribeChannelAllUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeChannelAllUsers"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeChannelAllUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询频道的所有参会者
//
// @param request - DescribeChannelAllUsersRequest
//
// @return DescribeChannelAllUsersResponse
func (client *Client) DescribeChannelAllUsers(request *DescribeChannelAllUsersRequest) (_result *DescribeChannelAllUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeChannelAllUsersResponse{}
	_body, _err := client.DescribeChannelAllUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 调用DescribeChannelAreaDistributionStatData获取频道地区分布统计数据。
//
// @param request - DescribeChannelAreaDistributionStatDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelAreaDistributionStatDataResponse
func (client *Client) DescribeChannelAreaDistributionStatDataWithOptions(request *DescribeChannelAreaDistributionStatDataRequest, runtime *util.RuntimeOptions) (_result *DescribeChannelAreaDistributionStatDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.CreatedTs)) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !tea.BoolValue(util.IsUnset(request.DestroyedTs)) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	if !tea.BoolValue(util.IsUnset(request.ParentArea)) {
		query["ParentArea"] = request.ParentArea
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeChannelAreaDistributionStatData"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeChannelAreaDistributionStatDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用DescribeChannelAreaDistributionStatData获取频道地区分布统计数据。
//
// @param request - DescribeChannelAreaDistributionStatDataRequest
//
// @return DescribeChannelAreaDistributionStatDataResponse
func (client *Client) DescribeChannelAreaDistributionStatData(request *DescribeChannelAreaDistributionStatDataRequest) (_result *DescribeChannelAreaDistributionStatDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeChannelAreaDistributionStatDataResponse{}
	_body, _err := client.DescribeChannelAreaDistributionStatDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 调用DescribeChannelDistributionStatData获取频道分布统计数据。
//
// @param request - DescribeChannelDistributionStatDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelDistributionStatDataResponse
func (client *Client) DescribeChannelDistributionStatDataWithOptions(request *DescribeChannelDistributionStatDataRequest, runtime *util.RuntimeOptions) (_result *DescribeChannelDistributionStatDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.CreatedTs)) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !tea.BoolValue(util.IsUnset(request.DestroyedTs)) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	if !tea.BoolValue(util.IsUnset(request.StatDim)) {
		query["StatDim"] = request.StatDim
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeChannelDistributionStatData"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeChannelDistributionStatDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用DescribeChannelDistributionStatData获取频道分布统计数据。
//
// @param request - DescribeChannelDistributionStatDataRequest
//
// @return DescribeChannelDistributionStatDataResponse
func (client *Client) DescribeChannelDistributionStatData(request *DescribeChannelDistributionStatDataRequest) (_result *DescribeChannelDistributionStatDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeChannelDistributionStatDataResponse{}
	_body, _err := client.DescribeChannelDistributionStatDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 调用DescribeChannelOverallData查询频道概览数据。
//
// @param request - DescribeChannelOverallDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelOverallDataResponse
func (client *Client) DescribeChannelOverallDataWithOptions(request *DescribeChannelOverallDataRequest, runtime *util.RuntimeOptions) (_result *DescribeChannelOverallDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.CreatedTs)) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !tea.BoolValue(util.IsUnset(request.DestroyedTs)) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeChannelOverallData"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeChannelOverallDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用DescribeChannelOverallData查询频道概览数据。
//
// @param request - DescribeChannelOverallDataRequest
//
// @return DescribeChannelOverallDataResponse
func (client *Client) DescribeChannelOverallData(request *DescribeChannelOverallDataRequest) (_result *DescribeChannelOverallDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeChannelOverallDataResponse{}
	_body, _err := client.DescribeChannelOverallDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeChannelParticipantsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelParticipantsResponse
func (client *Client) DescribeChannelParticipantsWithOptions(request *DescribeChannelParticipantsRequest, runtime *util.RuntimeOptions) (_result *DescribeChannelParticipantsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeChannelParticipants"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeChannelParticipantsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeChannelParticipantsRequest
//
// @return DescribeChannelParticipantsResponse
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

// Summary:
//
// 调用DescribeChannelTopPubUserList获取频道内发布端的用户列表（按用户在线时长降序）。
//
// @param request - DescribeChannelTopPubUserListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelTopPubUserListResponse
func (client *Client) DescribeChannelTopPubUserListWithOptions(request *DescribeChannelTopPubUserListRequest, runtime *util.RuntimeOptions) (_result *DescribeChannelTopPubUserListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.CreatedTs)) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !tea.BoolValue(util.IsUnset(request.DestroyedTs)) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeChannelTopPubUserList"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeChannelTopPubUserListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用DescribeChannelTopPubUserList获取频道内发布端的用户列表（按用户在线时长降序）。
//
// @param request - DescribeChannelTopPubUserListRequest
//
// @return DescribeChannelTopPubUserListResponse
func (client *Client) DescribeChannelTopPubUserList(request *DescribeChannelTopPubUserListRequest) (_result *DescribeChannelTopPubUserListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeChannelTopPubUserListResponse{}
	_body, _err := client.DescribeChannelTopPubUserListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// DescribeChannelUser
//
// @param request - DescribeChannelUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelUserResponse
func (client *Client) DescribeChannelUserWithOptions(request *DescribeChannelUserRequest, runtime *util.RuntimeOptions) (_result *DescribeChannelUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeChannelUser"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeChannelUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// DescribeChannelUser
//
// @param request - DescribeChannelUserRequest
//
// @return DescribeChannelUserResponse
func (client *Client) DescribeChannelUser(request *DescribeChannelUserRequest) (_result *DescribeChannelUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeChannelUserResponse{}
	_body, _err := client.DescribeChannelUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 调用DescribeChannelUserMetrics查询频道总览中的用户数据。
//
// @param request - DescribeChannelUserMetricsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelUserMetricsResponse
func (client *Client) DescribeChannelUserMetricsWithOptions(request *DescribeChannelUserMetricsRequest, runtime *util.RuntimeOptions) (_result *DescribeChannelUserMetricsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.CreatedTs)) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !tea.BoolValue(util.IsUnset(request.DestroyedTs)) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeChannelUserMetrics"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeChannelUserMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用DescribeChannelUserMetrics查询频道总览中的用户数据。
//
// @param request - DescribeChannelUserMetricsRequest
//
// @return DescribeChannelUserMetricsResponse
func (client *Client) DescribeChannelUserMetrics(request *DescribeChannelUserMetricsRequest) (_result *DescribeChannelUserMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeChannelUserMetricsResponse{}
	_body, _err := client.DescribeChannelUserMetricsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeChannelUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeChannelUsersResponse
func (client *Client) DescribeChannelUsersWithOptions(request *DescribeChannelUsersRequest, runtime *util.RuntimeOptions) (_result *DescribeChannelUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeChannelUsers"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeChannelUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeChannelUsersRequest
//
// @return DescribeChannelUsersResponse
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

// Summary:
//
// 调用DescribeEndPointEventList获取端对端用户事件列表。
//
// @param request - DescribeEndPointEventListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEndPointEventListResponse
func (client *Client) DescribeEndPointEventListWithOptions(request *DescribeEndPointEventListRequest, runtime *util.RuntimeOptions) (_result *DescribeEndPointEventListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.CreatedTs)) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !tea.BoolValue(util.IsUnset(request.DestroyedTs)) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	if !tea.BoolValue(util.IsUnset(request.UserIdList)) {
		query["UserIdList"] = request.UserIdList
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEndPointEventList"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEndPointEventListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用DescribeEndPointEventList获取端对端用户事件列表。
//
// @param request - DescribeEndPointEventListRequest
//
// @return DescribeEndPointEventListResponse
func (client *Client) DescribeEndPointEventList(request *DescribeEndPointEventListRequest) (_result *DescribeEndPointEventListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEndPointEventListResponse{}
	_body, _err := client.DescribeEndPointEventListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 调用DescribeEndPointMetricData获取端对端指标数据。
//
// @param request - DescribeEndPointMetricDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeEndPointMetricDataResponse
func (client *Client) DescribeEndPointMetricDataWithOptions(request *DescribeEndPointMetricDataRequest, runtime *util.RuntimeOptions) (_result *DescribeEndPointMetricDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.CreatedTs)) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !tea.BoolValue(util.IsUnset(request.DestroyedTs)) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	if !tea.BoolValue(util.IsUnset(request.Metrics)) {
		query["Metrics"] = request.Metrics
	}

	if !tea.BoolValue(util.IsUnset(request.PubCallIdList)) {
		query["PubCallIdList"] = request.PubCallIdList
	}

	if !tea.BoolValue(util.IsUnset(request.PubUserId)) {
		query["PubUserId"] = request.PubUserId
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		query["SubUserId"] = request.SubUserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeEndPointMetricData"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEndPointMetricDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用DescribeEndPointMetricData获取端对端指标数据。
//
// @param request - DescribeEndPointMetricDataRequest
//
// @return DescribeEndPointMetricDataResponse
func (client *Client) DescribeEndPointMetricData(request *DescribeEndPointMetricDataRequest) (_result *DescribeEndPointMetricDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEndPointMetricDataResponse{}
	_body, _err := client.DescribeEndPointMetricDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取异常诊断影响因素分布
//
// @param request - DescribeFaultDiagnosisFactorDistributionStatRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFaultDiagnosisFactorDistributionStatResponse
func (client *Client) DescribeFaultDiagnosisFactorDistributionStatWithOptions(request *DescribeFaultDiagnosisFactorDistributionStatRequest, runtime *util.RuntimeOptions) (_result *DescribeFaultDiagnosisFactorDistributionStatResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTs)) {
		query["EndTs"] = request.EndTs
	}

	if !tea.BoolValue(util.IsUnset(request.StartTs)) {
		query["StartTs"] = request.StartTs
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFaultDiagnosisFactorDistributionStat"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFaultDiagnosisFactorDistributionStatResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取异常诊断影响因素分布
//
// @param request - DescribeFaultDiagnosisFactorDistributionStatRequest
//
// @return DescribeFaultDiagnosisFactorDistributionStatResponse
func (client *Client) DescribeFaultDiagnosisFactorDistributionStat(request *DescribeFaultDiagnosisFactorDistributionStatRequest) (_result *DescribeFaultDiagnosisFactorDistributionStatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFaultDiagnosisFactorDistributionStatResponse{}
	_body, _err := client.DescribeFaultDiagnosisFactorDistributionStatWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取异常诊断总览数据
//
// @param request - DescribeFaultDiagnosisOverallDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFaultDiagnosisOverallDataResponse
func (client *Client) DescribeFaultDiagnosisOverallDataWithOptions(request *DescribeFaultDiagnosisOverallDataRequest, runtime *util.RuntimeOptions) (_result *DescribeFaultDiagnosisOverallDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTs)) {
		query["EndTs"] = request.EndTs
	}

	if !tea.BoolValue(util.IsUnset(request.StartTs)) {
		query["StartTs"] = request.StartTs
	}

	if !tea.BoolValue(util.IsUnset(request.StatDim)) {
		query["StatDim"] = request.StatDim
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFaultDiagnosisOverallData"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFaultDiagnosisOverallDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取异常诊断总览数据
//
// @param request - DescribeFaultDiagnosisOverallDataRequest
//
// @return DescribeFaultDiagnosisOverallDataResponse
func (client *Client) DescribeFaultDiagnosisOverallData(request *DescribeFaultDiagnosisOverallDataRequest) (_result *DescribeFaultDiagnosisOverallDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFaultDiagnosisOverallDataResponse{}
	_body, _err := client.DescribeFaultDiagnosisOverallDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取异常诊断用户详情
//
// @param request - DescribeFaultDiagnosisUserDetailRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFaultDiagnosisUserDetailResponse
func (client *Client) DescribeFaultDiagnosisUserDetailWithOptions(request *DescribeFaultDiagnosisUserDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeFaultDiagnosisUserDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.CreatedTs)) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !tea.BoolValue(util.IsUnset(request.FaultType)) {
		query["FaultType"] = request.FaultType
	}

	if !tea.BoolValue(util.IsUnset(request.QueryCallUserInfo)) {
		query["QueryCallUserInfo"] = request.QueryCallUserInfo
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFaultDiagnosisUserDetail"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFaultDiagnosisUserDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取异常诊断用户详情
//
// @param request - DescribeFaultDiagnosisUserDetailRequest
//
// @return DescribeFaultDiagnosisUserDetailResponse
func (client *Client) DescribeFaultDiagnosisUserDetail(request *DescribeFaultDiagnosisUserDetailRequest) (_result *DescribeFaultDiagnosisUserDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFaultDiagnosisUserDetailResponse{}
	_body, _err := client.DescribeFaultDiagnosisUserDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取异常诊断用户明细列表
//
// @param request - DescribeFaultDiagnosisUserListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFaultDiagnosisUserListResponse
func (client *Client) DescribeFaultDiagnosisUserListWithOptions(request *DescribeFaultDiagnosisUserListRequest, runtime *util.RuntimeOptions) (_result *DescribeFaultDiagnosisUserListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTs)) {
		query["EndTs"] = request.EndTs
	}

	if !tea.BoolValue(util.IsUnset(request.FaultTypes)) {
		query["FaultTypes"] = request.FaultTypes
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTs)) {
		query["StartTs"] = request.StartTs
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFaultDiagnosisUserList"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFaultDiagnosisUserListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取异常诊断用户明细列表
//
// @param request - DescribeFaultDiagnosisUserListRequest
//
// @return DescribeFaultDiagnosisUserListResponse
func (client *Client) DescribeFaultDiagnosisUserList(request *DescribeFaultDiagnosisUserListRequest) (_result *DescribeFaultDiagnosisUserListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeFaultDiagnosisUserListResponse{}
	_body, _err := client.DescribeFaultDiagnosisUserListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeMPULayoutInfoListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeMPULayoutInfoListResponse
func (client *Client) DescribeMPULayoutInfoListWithOptions(request *DescribeMPULayoutInfoListRequest, runtime *util.RuntimeOptions) (_result *DescribeMPULayoutInfoListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.LayoutId)) {
		query["LayoutId"] = request.LayoutId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMPULayoutInfoList"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMPULayoutInfoListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeMPULayoutInfoListRequest
//
// @return DescribeMPULayoutInfoListResponse
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

// Summary:
//
// 调用DescribePubUserListBySubUser根据订阅端获取通信中发布端用户列表。
//
// @param request - DescribePubUserListBySubUserRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribePubUserListBySubUserResponse
func (client *Client) DescribePubUserListBySubUserWithOptions(request *DescribePubUserListBySubUserRequest, runtime *util.RuntimeOptions) (_result *DescribePubUserListBySubUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.CreatedTs)) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !tea.BoolValue(util.IsUnset(request.DestroyedTs)) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	if !tea.BoolValue(util.IsUnset(request.SubUserId)) {
		query["SubUserId"] = request.SubUserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribePubUserListBySubUser"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribePubUserListBySubUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用DescribePubUserListBySubUser根据订阅端获取通信中发布端用户列表。
//
// @param request - DescribePubUserListBySubUserRequest
//
// @return DescribePubUserListBySubUserResponse
func (client *Client) DescribePubUserListBySubUser(request *DescribePubUserListBySubUserRequest) (_result *DescribePubUserListBySubUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePubUserListBySubUserResponse{}
	_body, _err := client.DescribePubUserListBySubUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 调用DescribeQoeMetricData获取单次通信中用户的下行体验质量指标。
//
// @param request - DescribeQoeMetricDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeQoeMetricDataResponse
func (client *Client) DescribeQoeMetricDataWithOptions(request *DescribeQoeMetricDataRequest, runtime *util.RuntimeOptions) (_result *DescribeQoeMetricDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.CreatedTs)) {
		query["CreatedTs"] = request.CreatedTs
	}

	if !tea.BoolValue(util.IsUnset(request.DestroyedTs)) {
		query["DestroyedTs"] = request.DestroyedTs
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeQoeMetricData"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeQoeMetricDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 调用DescribeQoeMetricData获取单次通信中用户的下行体验质量指标。
//
// @param request - DescribeQoeMetricDataRequest
//
// @return DescribeQoeMetricDataResponse
func (client *Client) DescribeQoeMetricData(request *DescribeQoeMetricDataRequest) (_result *DescribeQoeMetricDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeQoeMetricDataResponse{}
	_body, _err := client.DescribeQoeMetricDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取质量统计区域分布统计数据
//
// @param request - DescribeQualityAreaDistributionStatDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeQualityAreaDistributionStatDataResponse
func (client *Client) DescribeQualityAreaDistributionStatDataWithOptions(request *DescribeQualityAreaDistributionStatDataRequest, runtime *util.RuntimeOptions) (_result *DescribeQualityAreaDistributionStatDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.ParentArea)) {
		query["ParentArea"] = request.ParentArea
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeQualityAreaDistributionStatData"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeQualityAreaDistributionStatDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取质量统计区域分布统计数据
//
// @param request - DescribeQualityAreaDistributionStatDataRequest
//
// @return DescribeQualityAreaDistributionStatDataResponse
func (client *Client) DescribeQualityAreaDistributionStatData(request *DescribeQualityAreaDistributionStatDataRequest) (_result *DescribeQualityAreaDistributionStatDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeQualityAreaDistributionStatDataResponse{}
	_body, _err := client.DescribeQualityAreaDistributionStatDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取质量统计分布数据
//
// @param request - DescribeQualityDistributionStatDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeQualityDistributionStatDataResponse
func (client *Client) DescribeQualityDistributionStatDataWithOptions(request *DescribeQualityDistributionStatDataRequest, runtime *util.RuntimeOptions) (_result *DescribeQualityDistributionStatDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.StatDim)) {
		query["StatDim"] = request.StatDim
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeQualityDistributionStatData"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeQualityDistributionStatDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取质量统计分布数据
//
// @param request - DescribeQualityDistributionStatDataRequest
//
// @return DescribeQualityDistributionStatDataResponse
func (client *Client) DescribeQualityDistributionStatData(request *DescribeQualityDistributionStatDataRequest) (_result *DescribeQualityDistributionStatDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeQualityDistributionStatDataResponse{}
	_body, _err := client.DescribeQualityDistributionStatDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取质量统计各操作系统下SDK版本分布数据
//
// @param request - DescribeQualityOsSdkVersionDistributionStatDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeQualityOsSdkVersionDistributionStatDataResponse
func (client *Client) DescribeQualityOsSdkVersionDistributionStatDataWithOptions(request *DescribeQualityOsSdkVersionDistributionStatDataRequest, runtime *util.RuntimeOptions) (_result *DescribeQualityOsSdkVersionDistributionStatDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeQualityOsSdkVersionDistributionStatData"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeQualityOsSdkVersionDistributionStatDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取质量统计各操作系统下SDK版本分布数据
//
// @param request - DescribeQualityOsSdkVersionDistributionStatDataRequest
//
// @return DescribeQualityOsSdkVersionDistributionStatDataResponse
func (client *Client) DescribeQualityOsSdkVersionDistributionStatData(request *DescribeQualityOsSdkVersionDistributionStatDataRequest) (_result *DescribeQualityOsSdkVersionDistributionStatDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeQualityOsSdkVersionDistributionStatDataResponse{}
	_body, _err := client.DescribeQualityOsSdkVersionDistributionStatDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取质量统计概览数据
//
// @param request - DescribeQualityOverallDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeQualityOverallDataResponse
func (client *Client) DescribeQualityOverallDataWithOptions(request *DescribeQualityOverallDataRequest, runtime *util.RuntimeOptions) (_result *DescribeQualityOverallDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.Types)) {
		query["Types"] = request.Types
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeQualityOverallData"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeQualityOverallDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取质量统计概览数据
//
// @param request - DescribeQualityOverallDataRequest
//
// @return DescribeQualityOverallDataResponse
func (client *Client) DescribeQualityOverallData(request *DescribeQualityOverallDataRequest) (_result *DescribeQualityOverallDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeQualityOverallDataResponse{}
	_body, _err := client.DescribeQualityOverallDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeRecordFilesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRecordFilesResponse
func (client *Client) DescribeRecordFilesWithOptions(request *DescribeRecordFilesRequest, runtime *util.RuntimeOptions) (_result *DescribeRecordFilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TaskIds)) {
		query["TaskIds"] = request.TaskIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRecordFiles"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRecordFilesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRecordFilesRequest
//
// @return DescribeRecordFilesResponse
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

// @param request - DescribeRecordTemplatesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRecordTemplatesResponse
func (client *Client) DescribeRecordTemplatesWithOptions(request *DescribeRecordTemplatesRequest, runtime *util.RuntimeOptions) (_result *DescribeRecordTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateIds)) {
		query["TemplateIds"] = request.TemplateIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRecordTemplates"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRecordTemplatesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRecordTemplatesRequest
//
// @return DescribeRecordTemplatesResponse
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

// @param request - DescribeRtcChannelListRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRtcChannelListResponse
func (client *Client) DescribeRtcChannelListWithOptions(request *DescribeRtcChannelListRequest, runtime *util.RuntimeOptions) (_result *DescribeRtcChannelListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNo)) {
		query["PageNo"] = request.PageNo
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceArea)) {
		query["ServiceArea"] = request.ServiceArea
	}

	if !tea.BoolValue(util.IsUnset(request.SortType)) {
		query["SortType"] = request.SortType
	}

	if !tea.BoolValue(util.IsUnset(request.TimePoint)) {
		query["TimePoint"] = request.TimePoint
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRtcChannelList"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRtcChannelListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRtcChannelListRequest
//
// @return DescribeRtcChannelListResponse
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

// @param request - DescribeRtcChannelMetricRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRtcChannelMetricResponse
func (client *Client) DescribeRtcChannelMetricWithOptions(request *DescribeRtcChannelMetricRequest, runtime *util.RuntimeOptions) (_result *DescribeRtcChannelMetricResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TimePoint)) {
		query["TimePoint"] = request.TimePoint
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRtcChannelMetric"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRtcChannelMetricResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRtcChannelMetricRequest
//
// @return DescribeRtcChannelMetricResponse
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

// @param request - DescribeRtcDurationDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRtcDurationDataResponse
func (client *Client) DescribeRtcDurationDataWithOptions(request *DescribeRtcDurationDataRequest, runtime *util.RuntimeOptions) (_result *DescribeRtcDurationDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceArea)) {
		query["ServiceArea"] = request.ServiceArea
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRtcDurationData"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRtcDurationDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRtcDurationDataRequest
//
// @return DescribeRtcDurationDataResponse
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

// @param request - DescribeRtcPeakChannelCntDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRtcPeakChannelCntDataResponse
func (client *Client) DescribeRtcPeakChannelCntDataWithOptions(request *DescribeRtcPeakChannelCntDataRequest, runtime *util.RuntimeOptions) (_result *DescribeRtcPeakChannelCntDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceArea)) {
		query["ServiceArea"] = request.ServiceArea
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRtcPeakChannelCntData"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRtcPeakChannelCntDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRtcPeakChannelCntDataRequest
//
// @return DescribeRtcPeakChannelCntDataResponse
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

// @param request - DescribeRtcUserCntDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRtcUserCntDataResponse
func (client *Client) DescribeRtcUserCntDataWithOptions(request *DescribeRtcUserCntDataRequest, runtime *util.RuntimeOptions) (_result *DescribeRtcUserCntDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Interval)) {
		query["Interval"] = request.Interval
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceArea)) {
		query["ServiceArea"] = request.ServiceArea
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRtcUserCntData"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRtcUserCntDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeRtcUserCntDataRequest
//
// @return DescribeRtcUserCntDataResponse
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

// Summary:
//
// 获取用量统计地域分布数据
//
// @param request - DescribeUsageAreaDistributionStatDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUsageAreaDistributionStatDataResponse
func (client *Client) DescribeUsageAreaDistributionStatDataWithOptions(request *DescribeUsageAreaDistributionStatDataRequest, runtime *util.RuntimeOptions) (_result *DescribeUsageAreaDistributionStatDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.ParentArea)) {
		query["ParentArea"] = request.ParentArea
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUsageAreaDistributionStatData"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUsageAreaDistributionStatDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用量统计地域分布数据
//
// @param request - DescribeUsageAreaDistributionStatDataRequest
//
// @return DescribeUsageAreaDistributionStatDataResponse
func (client *Client) DescribeUsageAreaDistributionStatData(request *DescribeUsageAreaDistributionStatDataRequest) (_result *DescribeUsageAreaDistributionStatDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUsageAreaDistributionStatDataResponse{}
	_body, _err := client.DescribeUsageAreaDistributionStatDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用量统计分布数据
//
// @param request - DescribeUsageDistributionStatDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUsageDistributionStatDataResponse
func (client *Client) DescribeUsageDistributionStatDataWithOptions(request *DescribeUsageDistributionStatDataRequest, runtime *util.RuntimeOptions) (_result *DescribeUsageDistributionStatDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.StatDim)) {
		query["StatDim"] = request.StatDim
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUsageDistributionStatData"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUsageDistributionStatDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用量统计分布数据
//
// @param request - DescribeUsageDistributionStatDataRequest
//
// @return DescribeUsageDistributionStatDataResponse
func (client *Client) DescribeUsageDistributionStatData(request *DescribeUsageDistributionStatDataRequest) (_result *DescribeUsageDistributionStatDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUsageDistributionStatDataResponse{}
	_body, _err := client.DescribeUsageDistributionStatDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用量统计各操作系统下SDK版本分布数据
//
// @param request - DescribeUsageOsSdkVersionDistributionStatDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUsageOsSdkVersionDistributionStatDataResponse
func (client *Client) DescribeUsageOsSdkVersionDistributionStatDataWithOptions(request *DescribeUsageOsSdkVersionDistributionStatDataRequest, runtime *util.RuntimeOptions) (_result *DescribeUsageOsSdkVersionDistributionStatDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUsageOsSdkVersionDistributionStatData"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUsageOsSdkVersionDistributionStatDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用量统计各操作系统下SDK版本分布数据
//
// @param request - DescribeUsageOsSdkVersionDistributionStatDataRequest
//
// @return DescribeUsageOsSdkVersionDistributionStatDataResponse
func (client *Client) DescribeUsageOsSdkVersionDistributionStatData(request *DescribeUsageOsSdkVersionDistributionStatDataRequest) (_result *DescribeUsageOsSdkVersionDistributionStatDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUsageOsSdkVersionDistributionStatDataResponse{}
	_body, _err := client.DescribeUsageOsSdkVersionDistributionStatDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用量统计概览数据
//
// @param request - DescribeUsageOverallDataRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUsageOverallDataResponse
func (client *Client) DescribeUsageOverallDataWithOptions(request *DescribeUsageOverallDataRequest, runtime *util.RuntimeOptions) (_result *DescribeUsageOverallDataResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.EndDate)) {
		query["EndDate"] = request.EndDate
	}

	if !tea.BoolValue(util.IsUnset(request.StartDate)) {
		query["StartDate"] = request.StartDate
	}

	if !tea.BoolValue(util.IsUnset(request.Types)) {
		query["Types"] = request.Types
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUsageOverallData"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUsageOverallDataResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用量统计概览数据
//
// @param request - DescribeUsageOverallDataRequest
//
// @return DescribeUsageOverallDataResponse
func (client *Client) DescribeUsageOverallData(request *DescribeUsageOverallDataRequest) (_result *DescribeUsageOverallDataResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUsageOverallDataResponse{}
	_body, _err := client.DescribeUsageOverallDataWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - DescribeUserInfoInChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserInfoInChannelResponse
func (client *Client) DescribeUserInfoInChannelWithOptions(request *DescribeUserInfoInChannelRequest, runtime *util.RuntimeOptions) (_result *DescribeUserInfoInChannelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		query["UserId"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUserInfoInChannel"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUserInfoInChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DescribeUserInfoInChannelRequest
//
// @return DescribeUserInfoInChannelResponse
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

// @param request - DisableAutoLiveStreamRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableAutoLiveStreamRuleResponse
func (client *Client) DisableAutoLiveStreamRuleWithOptions(request *DisableAutoLiveStreamRuleRequest, runtime *util.RuntimeOptions) (_result *DisableAutoLiveStreamRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DisableAutoLiveStreamRule"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableAutoLiveStreamRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - DisableAutoLiveStreamRuleRequest
//
// @return DisableAutoLiveStreamRuleResponse
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

// @param request - EnableAutoLiveStreamRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableAutoLiveStreamRuleResponse
func (client *Client) EnableAutoLiveStreamRuleWithOptions(request *EnableAutoLiveStreamRuleRequest, runtime *util.RuntimeOptions) (_result *EnableAutoLiveStreamRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("EnableAutoLiveStreamRule"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableAutoLiveStreamRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - EnableAutoLiveStreamRuleRequest
//
// @return EnableAutoLiveStreamRuleResponse
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

// @param request - GetMPUTaskStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetMPUTaskStatusResponse
func (client *Client) GetMPUTaskStatusWithOptions(request *GetMPUTaskStatusRequest, runtime *util.RuntimeOptions) (_result *GetMPUTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetMPUTaskStatus"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetMPUTaskStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - GetMPUTaskStatusRequest
//
// @return GetMPUTaskStatusResponse
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

// Summary:
//
// 修改App信息
//
// @param request - ModifyAppRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppResponse
func (client *Client) ModifyAppWithOptions(request *ModifyAppRequest, runtime *util.RuntimeOptions) (_result *ModifyAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AppName)) {
		query["AppName"] = request.AppName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyApp"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 修改App信息
//
// @param request - ModifyAppRequest
//
// @return ModifyAppResponse
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

// Summary:
//
// 更新应用推流模版
//
// @param tmpReq - ModifyAppStreamingOutTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppStreamingOutTemplateResponse
func (client *Client) ModifyAppStreamingOutTemplateWithOptions(tmpReq *ModifyAppStreamingOutTemplateRequest, runtime *util.RuntimeOptions) (_result *ModifyAppStreamingOutTemplateResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ModifyAppStreamingOutTemplateShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.StreamingOutTemplate)) {
		request.StreamingOutTemplateShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StreamingOutTemplate, tea.String("StreamingOutTemplate"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.StreamingOutTemplateShrink)) {
		query["StreamingOutTemplate"] = request.StreamingOutTemplateShrink
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAppStreamingOutTemplate"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAppStreamingOutTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新应用推流模版
//
// @param request - ModifyAppStreamingOutTemplateRequest
//
// @return ModifyAppStreamingOutTemplateResponse
func (client *Client) ModifyAppStreamingOutTemplate(request *ModifyAppStreamingOutTemplateRequest) (_result *ModifyAppStreamingOutTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAppStreamingOutTemplateResponse{}
	_body, _err := client.ModifyAppStreamingOutTemplateWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - ModifyMPULayoutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyMPULayoutResponse
func (client *Client) ModifyMPULayoutWithOptions(request *ModifyMPULayoutRequest, runtime *util.RuntimeOptions) (_result *ModifyMPULayoutResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.AudioMixCount)) {
		query["AudioMixCount"] = request.AudioMixCount
	}

	if !tea.BoolValue(util.IsUnset(request.LayoutId)) {
		query["LayoutId"] = request.LayoutId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Panes)) {
		query["Panes"] = request.Panes
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyMPULayout"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyMPULayoutResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - ModifyMPULayoutRequest
//
// @return ModifyMPULayoutResponse
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

// @param request - RemoveTerminalsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveTerminalsResponse
func (client *Client) RemoveTerminalsWithOptions(request *RemoveTerminalsRequest, runtime *util.RuntimeOptions) (_result *RemoveTerminalsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TerminalIds)) {
		query["TerminalIds"] = request.TerminalIds
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveTerminals"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveTerminalsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - RemoveTerminalsRequest
//
// @return RemoveTerminalsResponse
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

// Summary:
//
// RemoveUsers
//
// @param request - RemoveUsersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUsersResponse
func (client *Client) RemoveUsersWithOptions(request *RemoveUsersRequest, runtime *util.RuntimeOptions) (_result *RemoveUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.Users)) {
		query["Users"] = request.Users
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveUsers"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// RemoveUsers
//
// @param request - RemoveUsersRequest
//
// @return RemoveUsersResponse
func (client *Client) RemoveUsers(request *RemoveUsersRequest) (_result *RemoveUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveUsersResponse{}
	_body, _err := client.RemoveUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// StartCloudRecord
//
// @param request - StartCloudRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartCloudRecordResponse
func (client *Client) StartCloudRecordWithOptions(request *StartCloudRecordRequest, runtime *util.RuntimeOptions) (_result *StartCloudRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.ClockWidgets)) {
		query["ClockWidgets"] = request.ClockWidgets
	}

	if !tea.BoolValue(util.IsUnset(request.CropMode)) {
		query["CropMode"] = request.CropMode
	}

	if !tea.BoolValue(util.IsUnset(request.Images)) {
		query["Images"] = request.Images
	}

	if !tea.BoolValue(util.IsUnset(request.Panes)) {
		query["Panes"] = request.Panes
	}

	if !tea.BoolValue(util.IsUnset(request.StorageConfig)) {
		query["StorageConfig"] = request.StorageConfig
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.Texts)) {
		query["Texts"] = request.Texts
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartCloudRecord"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartCloudRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// StartCloudRecord
//
// @param request - StartCloudRecordRequest
//
// @return StartCloudRecordResponse
func (client *Client) StartCloudRecord(request *StartCloudRecordRequest) (_result *StartCloudRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartCloudRecordResponse{}
	_body, _err := client.StartCloudRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - StartMPUTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartMPUTaskResponse
func (client *Client) StartMPUTaskWithOptions(request *StartMPUTaskRequest, runtime *util.RuntimeOptions) (_result *StartMPUTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.BackgroundColor)) {
		query["BackgroundColor"] = request.BackgroundColor
	}

	if !tea.BoolValue(util.IsUnset(request.Backgrounds)) {
		query["Backgrounds"] = request.Backgrounds
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.ClockWidgets)) {
		query["ClockWidgets"] = request.ClockWidgets
	}

	if !tea.BoolValue(util.IsUnset(request.CropMode)) {
		query["CropMode"] = request.CropMode
	}

	if !tea.BoolValue(util.IsUnset(request.LayoutIds)) {
		query["LayoutIds"] = request.LayoutIds
	}

	if !tea.BoolValue(util.IsUnset(request.MediaEncode)) {
		query["MediaEncode"] = request.MediaEncode
	}

	if !tea.BoolValue(util.IsUnset(request.MixMode)) {
		query["MixMode"] = request.MixMode
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PayloadType)) {
		query["PayloadType"] = request.PayloadType
	}

	if !tea.BoolValue(util.IsUnset(request.ReportVad)) {
		query["ReportVad"] = request.ReportVad
	}

	if !tea.BoolValue(util.IsUnset(request.RtpExtInfo)) {
		query["RtpExtInfo"] = request.RtpExtInfo
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.StreamType)) {
		query["StreamType"] = request.StreamType
	}

	if !tea.BoolValue(util.IsUnset(request.StreamURL)) {
		query["StreamURL"] = request.StreamURL
	}

	if !tea.BoolValue(util.IsUnset(request.SubSpecAudioUsers)) {
		query["SubSpecAudioUsers"] = request.SubSpecAudioUsers
	}

	if !tea.BoolValue(util.IsUnset(request.SubSpecCameraUsers)) {
		query["SubSpecCameraUsers"] = request.SubSpecCameraUsers
	}

	if !tea.BoolValue(util.IsUnset(request.SubSpecShareScreenUsers)) {
		query["SubSpecShareScreenUsers"] = request.SubSpecShareScreenUsers
	}

	if !tea.BoolValue(util.IsUnset(request.SubSpecUsers)) {
		query["SubSpecUsers"] = request.SubSpecUsers
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		query["TaskType"] = request.TaskType
	}

	if !tea.BoolValue(util.IsUnset(request.TimeStampRef)) {
		query["TimeStampRef"] = request.TimeStampRef
	}

	if !tea.BoolValue(util.IsUnset(request.UnsubSpecAudioUsers)) {
		query["UnsubSpecAudioUsers"] = request.UnsubSpecAudioUsers
	}

	if !tea.BoolValue(util.IsUnset(request.UnsubSpecCameraUsers)) {
		query["UnsubSpecCameraUsers"] = request.UnsubSpecCameraUsers
	}

	if !tea.BoolValue(util.IsUnset(request.UnsubSpecShareScreenUsers)) {
		query["UnsubSpecShareScreenUsers"] = request.UnsubSpecShareScreenUsers
	}

	if !tea.BoolValue(util.IsUnset(request.UserPanes)) {
		query["UserPanes"] = request.UserPanes
	}

	if !tea.BoolValue(util.IsUnset(request.VadInterval)) {
		query["VadInterval"] = request.VadInterval
	}

	if !tea.BoolValue(util.IsUnset(request.Watermarks)) {
		query["Watermarks"] = request.Watermarks
	}

	body := map[string]interface{}{}
	bodyFlat := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EnhancedParam)) {
		bodyFlat["EnhancedParam"] = request.EnhancedParam
	}

	body = tea.ToMap(body,
		openapiutil.Query(bodyFlat))
	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StartMPUTask"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartMPUTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - StartMPUTaskRequest
//
// @return StartMPUTaskResponse
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

// @param request - StartRecordTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartRecordTaskResponse
func (client *Client) StartRecordTaskWithOptions(request *StartRecordTaskRequest, runtime *util.RuntimeOptions) (_result *StartRecordTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.CropMode)) {
		query["CropMode"] = request.CropMode
	}

	if !tea.BoolValue(util.IsUnset(request.LayoutIds)) {
		query["LayoutIds"] = request.LayoutIds
	}

	if !tea.BoolValue(util.IsUnset(request.MediaEncode)) {
		query["MediaEncode"] = request.MediaEncode
	}

	if !tea.BoolValue(util.IsUnset(request.MixMode)) {
		query["MixMode"] = request.MixMode
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.StreamType)) {
		query["StreamType"] = request.StreamType
	}

	if !tea.BoolValue(util.IsUnset(request.SubSpecAudioUsers)) {
		query["SubSpecAudioUsers"] = request.SubSpecAudioUsers
	}

	if !tea.BoolValue(util.IsUnset(request.SubSpecCameraUsers)) {
		query["SubSpecCameraUsers"] = request.SubSpecCameraUsers
	}

	if !tea.BoolValue(util.IsUnset(request.SubSpecShareScreenUsers)) {
		query["SubSpecShareScreenUsers"] = request.SubSpecShareScreenUsers
	}

	if !tea.BoolValue(util.IsUnset(request.SubSpecUsers)) {
		query["SubSpecUsers"] = request.SubSpecUsers
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskProfile)) {
		query["TaskProfile"] = request.TaskProfile
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.UnsubSpecAudioUsers)) {
		query["UnsubSpecAudioUsers"] = request.UnsubSpecAudioUsers
	}

	if !tea.BoolValue(util.IsUnset(request.UnsubSpecCameraUsers)) {
		query["UnsubSpecCameraUsers"] = request.UnsubSpecCameraUsers
	}

	if !tea.BoolValue(util.IsUnset(request.UnsubSpecShareScreenUsers)) {
		query["UnsubSpecShareScreenUsers"] = request.UnsubSpecShareScreenUsers
	}

	if !tea.BoolValue(util.IsUnset(request.UserPanes)) {
		query["UserPanes"] = request.UserPanes
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartRecordTask"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartRecordTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - StartRecordTaskRequest
//
// @return StartRecordTaskResponse
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

// Summary:
//
// StartStreamingOut
//
// @param request - StartStreamingOutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartStreamingOutResponse
func (client *Client) StartStreamingOutWithOptions(request *StartStreamingOutRequest, runtime *util.RuntimeOptions) (_result *StartStreamingOutResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.ClockWidgets)) {
		query["ClockWidgets"] = request.ClockWidgets
	}

	if !tea.BoolValue(util.IsUnset(request.CropMode)) {
		query["CropMode"] = request.CropMode
	}

	if !tea.BoolValue(util.IsUnset(request.Images)) {
		query["Images"] = request.Images
	}

	if !tea.BoolValue(util.IsUnset(request.Panes)) {
		query["Panes"] = request.Panes
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.Texts)) {
		query["Texts"] = request.Texts
	}

	if !tea.BoolValue(util.IsUnset(request.Url)) {
		query["Url"] = request.Url
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartStreamingOut"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartStreamingOutResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// StartStreamingOut
//
// @param request - StartStreamingOutRequest
//
// @return StartStreamingOutResponse
func (client *Client) StartStreamingOut(request *StartStreamingOutRequest) (_result *StartStreamingOutResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartStreamingOutResponse{}
	_body, _err := client.StartStreamingOutWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除频道
//
// @param request - StopChannelRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopChannelResponse
func (client *Client) StopChannelWithOptions(request *StopChannelRequest, runtime *util.RuntimeOptions) (_result *StopChannelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopChannel"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopChannelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除频道
//
// @param request - StopChannelRequest
//
// @return StopChannelResponse
func (client *Client) StopChannel(request *StopChannelRequest) (_result *StopChannelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopChannelResponse{}
	_body, _err := client.StopChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// StopCloudRecord
//
// @param request - StopCloudRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopCloudRecordResponse
func (client *Client) StopCloudRecordWithOptions(request *StopCloudRecordRequest, runtime *util.RuntimeOptions) (_result *StopCloudRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopCloudRecord"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopCloudRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// StopCloudRecord
//
// @param request - StopCloudRecordRequest
//
// @return StopCloudRecordResponse
func (client *Client) StopCloudRecord(request *StopCloudRecordRequest) (_result *StopCloudRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopCloudRecordResponse{}
	_body, _err := client.StopCloudRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - StopMPUTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopMPUTaskResponse
func (client *Client) StopMPUTaskWithOptions(request *StopMPUTaskRequest, runtime *util.RuntimeOptions) (_result *StopMPUTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopMPUTask"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopMPUTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - StopMPUTaskRequest
//
// @return StopMPUTaskResponse
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

// @param request - StopRecordTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopRecordTaskResponse
func (client *Client) StopRecordTaskWithOptions(request *StopRecordTaskRequest, runtime *util.RuntimeOptions) (_result *StopRecordTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopRecordTask"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopRecordTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - StopRecordTaskRequest
//
// @return StopRecordTaskResponse
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

// Summary:
//
// StopStreamingOut
//
// @param request - StopStreamingOutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopStreamingOutResponse
func (client *Client) StopStreamingOutWithOptions(request *StopStreamingOutRequest, runtime *util.RuntimeOptions) (_result *StopStreamingOutResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopStreamingOut"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopStreamingOutResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// StopStreamingOut
//
// @param request - StopStreamingOutRequest
//
// @return StopStreamingOutResponse
func (client *Client) StopStreamingOut(request *StopStreamingOutRequest) (_result *StopStreamingOutResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopStreamingOutResponse{}
	_body, _err := client.StopStreamingOutWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateAutoLiveStreamRuleRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateAutoLiveStreamRuleResponse
func (client *Client) UpdateAutoLiveStreamRuleWithOptions(request *UpdateAutoLiveStreamRuleRequest, runtime *util.RuntimeOptions) (_result *UpdateAutoLiveStreamRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.CallBack)) {
		query["CallBack"] = request.CallBack
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelIdPrefixes)) {
		query["ChannelIdPrefixes"] = request.ChannelIdPrefixes
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelIds)) {
		query["ChannelIds"] = request.ChannelIds
	}

	if !tea.BoolValue(util.IsUnset(request.MediaEncode)) {
		query["MediaEncode"] = request.MediaEncode
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PlayDomain)) {
		query["PlayDomain"] = request.PlayDomain
	}

	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["RuleId"] = request.RuleId
	}

	if !tea.BoolValue(util.IsUnset(request.RuleName)) {
		query["RuleName"] = request.RuleName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateAutoLiveStreamRule"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateAutoLiveStreamRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateAutoLiveStreamRuleRequest
//
// @return UpdateAutoLiveStreamRuleResponse
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

// Summary:
//
// 更新云端录制任务
//
// @param request - UpdateCloudRecordRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateCloudRecordResponse
func (client *Client) UpdateCloudRecordWithOptions(request *UpdateCloudRecordRequest, runtime *util.RuntimeOptions) (_result *UpdateCloudRecordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.ClockWidgets)) {
		query["ClockWidgets"] = request.ClockWidgets
	}

	if !tea.BoolValue(util.IsUnset(request.Images)) {
		query["Images"] = request.Images
	}

	if !tea.BoolValue(util.IsUnset(request.Panes)) {
		query["Panes"] = request.Panes
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.Texts)) {
		query["Texts"] = request.Texts
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCloudRecord"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCloudRecordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新云端录制任务
//
// @param request - UpdateCloudRecordRequest
//
// @return UpdateCloudRecordResponse
func (client *Client) UpdateCloudRecord(request *UpdateCloudRecordRequest) (_result *UpdateCloudRecordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateCloudRecordResponse{}
	_body, _err := client.UpdateCloudRecordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - UpdateMPUTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateMPUTaskResponse
func (client *Client) UpdateMPUTaskWithOptions(request *UpdateMPUTaskRequest, runtime *util.RuntimeOptions) (_result *UpdateMPUTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.BackgroundColor)) {
		query["BackgroundColor"] = request.BackgroundColor
	}

	if !tea.BoolValue(util.IsUnset(request.Backgrounds)) {
		query["Backgrounds"] = request.Backgrounds
	}

	if !tea.BoolValue(util.IsUnset(request.ClockWidgets)) {
		query["ClockWidgets"] = request.ClockWidgets
	}

	if !tea.BoolValue(util.IsUnset(request.CropMode)) {
		query["CropMode"] = request.CropMode
	}

	if !tea.BoolValue(util.IsUnset(request.LayoutIds)) {
		query["LayoutIds"] = request.LayoutIds
	}

	if !tea.BoolValue(util.IsUnset(request.MediaEncode)) {
		query["MediaEncode"] = request.MediaEncode
	}

	if !tea.BoolValue(util.IsUnset(request.MixMode)) {
		query["MixMode"] = request.MixMode
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.StreamType)) {
		query["StreamType"] = request.StreamType
	}

	if !tea.BoolValue(util.IsUnset(request.SubSpecAudioUsers)) {
		query["SubSpecAudioUsers"] = request.SubSpecAudioUsers
	}

	if !tea.BoolValue(util.IsUnset(request.SubSpecCameraUsers)) {
		query["SubSpecCameraUsers"] = request.SubSpecCameraUsers
	}

	if !tea.BoolValue(util.IsUnset(request.SubSpecShareScreenUsers)) {
		query["SubSpecShareScreenUsers"] = request.SubSpecShareScreenUsers
	}

	if !tea.BoolValue(util.IsUnset(request.SubSpecUsers)) {
		query["SubSpecUsers"] = request.SubSpecUsers
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.UnsubSpecAudioUsers)) {
		query["UnsubSpecAudioUsers"] = request.UnsubSpecAudioUsers
	}

	if !tea.BoolValue(util.IsUnset(request.UnsubSpecCameraUsers)) {
		query["UnsubSpecCameraUsers"] = request.UnsubSpecCameraUsers
	}

	if !tea.BoolValue(util.IsUnset(request.UnsubSpecShareScreenUsers)) {
		query["UnsubSpecShareScreenUsers"] = request.UnsubSpecShareScreenUsers
	}

	if !tea.BoolValue(util.IsUnset(request.UserPanes)) {
		query["UserPanes"] = request.UserPanes
	}

	if !tea.BoolValue(util.IsUnset(request.Watermarks)) {
		query["Watermarks"] = request.Watermarks
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateMPUTask"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateMPUTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateMPUTaskRequest
//
// @return UpdateMPUTaskResponse
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

// @param request - UpdateRecordTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRecordTaskResponse
func (client *Client) UpdateRecordTaskWithOptions(request *UpdateRecordTaskRequest, runtime *util.RuntimeOptions) (_result *UpdateRecordTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.CropMode)) {
		query["CropMode"] = request.CropMode
	}

	if !tea.BoolValue(util.IsUnset(request.LayoutIds)) {
		query["LayoutIds"] = request.LayoutIds
	}

	if !tea.BoolValue(util.IsUnset(request.MediaEncode)) {
		query["MediaEncode"] = request.MediaEncode
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SubSpecAudioUsers)) {
		query["SubSpecAudioUsers"] = request.SubSpecAudioUsers
	}

	if !tea.BoolValue(util.IsUnset(request.SubSpecCameraUsers)) {
		query["SubSpecCameraUsers"] = request.SubSpecCameraUsers
	}

	if !tea.BoolValue(util.IsUnset(request.SubSpecShareScreenUsers)) {
		query["SubSpecShareScreenUsers"] = request.SubSpecShareScreenUsers
	}

	if !tea.BoolValue(util.IsUnset(request.SubSpecUsers)) {
		query["SubSpecUsers"] = request.SubSpecUsers
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskProfile)) {
		query["TaskProfile"] = request.TaskProfile
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.UnsubSpecAudioUsers)) {
		query["UnsubSpecAudioUsers"] = request.UnsubSpecAudioUsers
	}

	if !tea.BoolValue(util.IsUnset(request.UnsubSpecCameraUsers)) {
		query["UnsubSpecCameraUsers"] = request.UnsubSpecCameraUsers
	}

	if !tea.BoolValue(util.IsUnset(request.UnsubSpecShareScreenUsers)) {
		query["UnsubSpecShareScreenUsers"] = request.UnsubSpecShareScreenUsers
	}

	if !tea.BoolValue(util.IsUnset(request.UserPanes)) {
		query["UserPanes"] = request.UserPanes
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRecordTask"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRecordTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateRecordTaskRequest
//
// @return UpdateRecordTaskResponse
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

// @param request - UpdateRecordTemplateRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRecordTemplateResponse
func (client *Client) UpdateRecordTemplateWithOptions(request *UpdateRecordTemplateRequest, runtime *util.RuntimeOptions) (_result *UpdateRecordTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.BackgroundColor)) {
		query["BackgroundColor"] = request.BackgroundColor
	}

	if !tea.BoolValue(util.IsUnset(request.Backgrounds)) {
		query["Backgrounds"] = request.Backgrounds
	}

	if !tea.BoolValue(util.IsUnset(request.ClockWidgets)) {
		query["ClockWidgets"] = request.ClockWidgets
	}

	if !tea.BoolValue(util.IsUnset(request.DelayStopTime)) {
		query["DelayStopTime"] = request.DelayStopTime
	}

	if !tea.BoolValue(util.IsUnset(request.EnableM3u8DateTime)) {
		query["EnableM3u8DateTime"] = request.EnableM3u8DateTime
	}

	if !tea.BoolValue(util.IsUnset(request.FileSplitInterval)) {
		query["FileSplitInterval"] = request.FileSplitInterval
	}

	if !tea.BoolValue(util.IsUnset(request.Formats)) {
		query["Formats"] = request.Formats
	}

	if !tea.BoolValue(util.IsUnset(request.HttpCallbackUrl)) {
		query["HttpCallbackUrl"] = request.HttpCallbackUrl
	}

	if !tea.BoolValue(util.IsUnset(request.LayoutIds)) {
		query["LayoutIds"] = request.LayoutIds
	}

	if !tea.BoolValue(util.IsUnset(request.MediaEncode)) {
		query["MediaEncode"] = request.MediaEncode
	}

	if !tea.BoolValue(util.IsUnset(request.MnsQueue)) {
		query["MnsQueue"] = request.MnsQueue
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OssBucket)) {
		query["OssBucket"] = request.OssBucket
	}

	if !tea.BoolValue(util.IsUnset(request.OssEndpoint)) {
		query["OssEndpoint"] = request.OssEndpoint
	}

	if !tea.BoolValue(util.IsUnset(request.OssFilePrefix)) {
		query["OssFilePrefix"] = request.OssFilePrefix
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskProfile)) {
		query["TaskProfile"] = request.TaskProfile
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.Watermarks)) {
		query["Watermarks"] = request.Watermarks
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRecordTemplate"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRecordTemplateResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - UpdateRecordTemplateRequest
//
// @return UpdateRecordTemplateResponse
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

// Summary:
//
// 更新旁路推流任务
//
// @param request - UpdateStreamingOutRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateStreamingOutResponse
func (client *Client) UpdateStreamingOutWithOptions(request *UpdateStreamingOutRequest, runtime *util.RuntimeOptions) (_result *UpdateStreamingOutResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppId)) {
		query["AppId"] = request.AppId
	}

	if !tea.BoolValue(util.IsUnset(request.ChannelId)) {
		query["ChannelId"] = request.ChannelId
	}

	if !tea.BoolValue(util.IsUnset(request.ClockWidgets)) {
		query["ClockWidgets"] = request.ClockWidgets
	}

	if !tea.BoolValue(util.IsUnset(request.Images)) {
		query["Images"] = request.Images
	}

	if !tea.BoolValue(util.IsUnset(request.Panes)) {
		query["Panes"] = request.Panes
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		query["TemplateId"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.Texts)) {
		query["Texts"] = request.Texts
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateStreamingOut"),
		Version:     tea.String("2018-01-11"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateStreamingOutResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新旁路推流任务
//
// @param request - UpdateStreamingOutRequest
//
// @return UpdateStreamingOutResponse
func (client *Client) UpdateStreamingOut(request *UpdateStreamingOutRequest) (_result *UpdateStreamingOutResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateStreamingOutResponse{}
	_body, _err := client.UpdateStreamingOutWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
